package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/myzhan/boomer"
	"go.uber.org/zap"

	msgPb "wakutester/pb"

	"github.com/waku-org/go-waku/waku/v2/node"
	"github.com/waku-org/go-waku/waku/v2/payload"
	"github.com/waku-org/go-waku/waku/v2/protocol"
	"github.com/waku-org/go-waku/waku/v2/protocol/pb"
	"github.com/waku-org/go-waku/waku/v2/utils"
)

var PeerHostList = []string{
	// write your boot waku nodes p2p address
}

func GetRandomHost() string {
	rand.Seed(time.Now().Unix())

	n := rand.Int() % len(PeerHostList)

	return PeerHostList[n]
}

var ContentTopic = protocol.NewContentTopic("waku", 2, "default-waku", "proto")

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

var NodeList = []string{}

func GetRandomNodeID() string {
	if len(NodeList) == 0 {
		return ""
	}

	rand.Seed(time.Now().Unix())

	n := rand.Int() % len(NodeList)

	return NodeList[n]
}

type WakuSocket struct {
	WakuNode *node.WakuNode
	ctx      context.Context
}

type MessageData struct {
	ComeFrom     string
	FromSign     string
	Version      uint32
	Payload      []byte
	ContentTopic string
	Timestamp    uint64
	CipherSuite  string
	PayloadType  string
	MessageType  string
	MessageId    string
	NeedStore    bool
	NodeId       string
}

func GetMessageData(bytedata []byte) (*MessageData, error) {
	data := &msgPb.Web3CommonMessage{}

	if err := data.Unmarshal(bytedata); err != nil {
		log.Println(err)
		return nil, err
	}
	MessageType := data.MessageType

	if data.MessageType == "" {
		MessageType = "message"
	}

	retdata := &MessageData{
		ComeFrom:     data.ComeFrom,
		FromSign:     data.FromSign,
		ContentTopic: data.ContentTopic,
		Version:      data.Version,
		Payload:      data.Payload,
		PayloadType:  data.PayloadType,
		Timestamp:    data.Timestamp,
		MessageId:    data.MessageId,
		CipherSuite:  data.CipherSuite,
		NeedStore:    data.NeedStore,
		MessageType:  MessageType,
		NodeId:       data.NodeId,
	}
	return retdata, nil
}

func NewWakuNode() *WakuSocket {
	hostAddr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:0")
	key, err := randomHex(32)
	if err != nil {
		fmt.Errorf("Could not generate random key", zap.Error(err))
		return nil
	}
	prvKey, err := crypto.HexToECDSA(key)
	if err != nil {
		fmt.Errorf("Could not convert hex into ecdsa key", zap.Error(err))
		return nil
	}

	ctx := context.Background()

	wNode, err := node.New(
		node.WithPrivateKey(prvKey),
		node.WithHostAddress(hostAddr),
		node.WithWakuRelay(),
	)
	if err != nil {
		fmt.Errorf("Error creating wakunode", zap.Error(err))
		return nil
	}

	return &WakuSocket{
		WakuNode: wNode,

		ctx: ctx,
	}
}

func (w *WakuSocket) Start() error {
	if err := w.WakuNode.Start(w.ctx); err != nil {
		fmt.Errorf("Error starting wakunode", zap.Error(err))
		return err
	}

	return nil
}

func (w *WakuSocket) DialPeer(peer string) error {
	return w.WakuNode.DialPeer(w.ctx, peer)
}

func (w *WakuSocket) Stop() {
	w.WakuNode.Stop()
}

func write(ctx context.Context, wakuNode *WakuSocket) {
	topicID := GetRandomNodeID()
	if topicID == "" {
		return
	}

	var version uint32 = 0
	var timestamp int64 = utils.GetUnixEpoch(wakuNode.WakuNode.Timesource())

	p := new(payload.Payload)

	payloadByte := make([]byte, 1024)
	rand.Read(payloadByte)

	msgPbx := &msgPb.Web3CommonMessage{
		Timestamp:    uint64(MakeTimestamp()),
		Payload:      payloadByte,
		ContentTopic: topicID,
		ComeFrom:     wakuNode.WakuNode.ID(),
	}
	byteData, err := msgPbx.Marshal()
	if err != nil {
		return
	}

	p.Data = byteData
	p.Key = &payload.KeyInfo{Kind: payload.None}

	payload, err := p.Encode(version)
	if err != nil {
		fmt.Errorf("Error encoding the payload", zap.Error(err))
		return
	}

	msg := &pb.WakuMessage{
		Payload:      payload,
		Version:      version,
		ContentTopic: ContentTopic.String(),
		Timestamp:    timestamp,
	}

	_, err = wakuNode.WakuNode.Relay().Publish(ctx, msg)
	if err != nil {
		fmt.Errorf("Error sending a message", zap.Error(err))
	}
}

func writeLoop(ctx context.Context, wakuNode *WakuSocket) {
	for {
		// Controls the frequency of sending. Sends every second.
		time.Sleep(1 * time.Second)
		write(ctx, wakuNode)
	}
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func readLoop(ctx context.Context, wakuNode *WakuSocket) {
	sub, err := wakuNode.WakuNode.Relay().Subscribe(ctx)
	if err != nil {
		fmt.Errorf("Could not subscribe", zap.Error(err))
		return
	}

	for envelope := range sub.Ch {
		payload, err := payload.DecodePayload(envelope.Message(), &payload.KeyInfo{Kind: payload.None})
		if err != nil {
			fmt.Errorf("Error decoding payload", zap.Error(err))
			continue
		}

		msgData, err := GetMessageData(payload.Data)
		if err != nil {
			fmt.Errorf("err", zap.Error(err))
			continue
		}
		if msgData.ContentTopic != wakuNode.WakuNode.ID() {
			continue
		}

		diff_ts := uint64(MakeTimestamp()) - msgData.Timestamp

		boomer.RecordSuccess("success", "waku", int64(diff_ts), int64(1024))
	}
}

func wakutest() {
	wakuNode := NewWakuNode()

	err := wakuNode.Start()
	if err != nil {
		panic(err)
	}

	NodeList = append(NodeList, wakuNode.WakuNode.ID())

	randhost := GetRandomHost()

	wakuNode.DialPeer(randhost)

	go writeLoop(wakuNode.ctx, wakuNode)
	go readLoop(wakuNode.ctx, wakuNode)

	// Wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("\n\n\nReceived signal, shutting down...")

	// shut the node down
	wakuNode.Stop()
}

func main() {
	task := &boomer.Task{
		Name: "waku",
		// The weight is used to distribute goroutines over multiple tasks.
		Fn: wakutest,
	}

	boomer.Run(task)
}
