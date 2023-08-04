// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Web3CommonMessage struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	ContentTopic         string   `protobuf:"bytes,2,opt,name=contentTopic,proto3" json:"contentTopic,omitempty"`
	Version              uint32   `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	ComeFrom             string   `protobuf:"bytes,4,opt,name=comeFrom,proto3" json:"comeFrom,omitempty"`
	FromSign             string   `protobuf:"bytes,5,opt,name=fromSign,proto3" json:"fromSign,omitempty"`
	PayloadType          string   `protobuf:"bytes,6,opt,name=payloadType,proto3" json:"payloadType,omitempty"`
	CipherSuite          string   `protobuf:"bytes,7,opt,name=cipherSuite,proto3" json:"cipherSuite,omitempty"`
	NeedStore            bool     `protobuf:"varint,8,opt,name=needStore,proto3" json:"needStore,omitempty"`
	Timestamp            uint64   `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MessageId            string   `protobuf:"bytes,10,opt,name=messageId,proto3" json:"messageId,omitempty"`
	MessageType          string   `protobuf:"bytes,11,opt,name=messageType,proto3" json:"messageType,omitempty"`
	NodeId               string   `protobuf:"bytes,12,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Web3CommonMessage) Reset()         { *m = Web3CommonMessage{} }
func (m *Web3CommonMessage) String() string { return proto.CompactTextString(m) }
func (*Web3CommonMessage) ProtoMessage()    {}
func (*Web3CommonMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *Web3CommonMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Web3CommonMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Web3CommonMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Web3CommonMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Web3CommonMessage.Merge(m, src)
}
func (m *Web3CommonMessage) XXX_Size() int {
	return m.Size()
}
func (m *Web3CommonMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Web3CommonMessage.DiscardUnknown(m)
}

var xxx_messageInfo_Web3CommonMessage proto.InternalMessageInfo

func (m *Web3CommonMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Web3CommonMessage) GetContentTopic() string {
	if m != nil {
		return m.ContentTopic
	}
	return ""
}

func (m *Web3CommonMessage) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Web3CommonMessage) GetComeFrom() string {
	if m != nil {
		return m.ComeFrom
	}
	return ""
}

func (m *Web3CommonMessage) GetFromSign() string {
	if m != nil {
		return m.FromSign
	}
	return ""
}

func (m *Web3CommonMessage) GetPayloadType() string {
	if m != nil {
		return m.PayloadType
	}
	return ""
}

func (m *Web3CommonMessage) GetCipherSuite() string {
	if m != nil {
		return m.CipherSuite
	}
	return ""
}

func (m *Web3CommonMessage) GetNeedStore() bool {
	if m != nil {
		return m.NeedStore
	}
	return false
}

func (m *Web3CommonMessage) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Web3CommonMessage) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *Web3CommonMessage) GetMessageType() string {
	if m != nil {
		return m.MessageType
	}
	return ""
}

func (m *Web3CommonMessage) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func init() {
	proto.RegisterType((*Web3CommonMessage)(nil), "pb.Web3CommonMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x86, 0xbf, 0xcc, 0x37, 0xce, 0x4c, 0x33, 0x1d, 0xd0, 0x2c, 0x24, 0x88, 0x94, 0x30, 0xab,
	0xae, 0xdc, 0xcc, 0x1d, 0x28, 0x08, 0x2e, 0xdc, 0xb4, 0x03, 0xae, 0xfb, 0x73, 0x1c, 0x03, 0x26,
	0x27, 0xa4, 0x51, 0x98, 0x3b, 0xf1, 0x92, 0x5c, 0x7a, 0x09, 0x52, 0x2f, 0xc3, 0x8d, 0x24, 0x6d,
	0x6d, 0x5d, 0xbe, 0xcf, 0x73, 0xde, 0x9e, 0x43, 0x43, 0x37, 0x0a, 0x9a, 0xa6, 0x38, 0xc0, 0x95,
	0xb1, 0xe8, 0x90, 0xcd, 0x4c, 0xb9, 0xfd, 0x9e, 0xd1, 0xb3, 0x07, 0x28, 0x77, 0x37, 0xa8, 0x14,
	0xea, 0xfb, 0xce, 0x33, 0x4e, 0x97, 0xa6, 0x38, 0x3e, 0x63, 0x51, 0x73, 0x22, 0x48, 0x1a, 0x67,
	0x43, 0x64, 0x5b, 0x1a, 0x57, 0xa8, 0x1d, 0x68, 0xb7, 0x47, 0x23, 0x2b, 0x3e, 0x13, 0x24, 0x8d,
	0xb2, 0x3f, 0xcc, 0xb7, 0x5f, 0xc1, 0x36, 0x12, 0x35, 0xff, 0x2f, 0x48, 0xba, 0xc9, 0x86, 0xc8,
	0x2e, 0xe8, 0xaa, 0x42, 0x05, 0xb7, 0x16, 0x15, 0x9f, 0x87, 0xe6, 0x6f, 0xf6, 0xee, 0xd1, 0xa2,
	0xca, 0xe5, 0x41, 0xf3, 0x93, 0xce, 0x0d, 0x99, 0x09, 0xba, 0xee, 0x0f, 0xd8, 0x1f, 0x0d, 0xf0,
	0x45, 0xd0, 0x53, 0xe4, 0x27, 0x2a, 0x69, 0x9e, 0xc0, 0xe6, 0x2f, 0xd2, 0x01, 0x5f, 0x76, 0x13,
	0x13, 0xc4, 0x2e, 0x69, 0xa4, 0x01, 0xea, 0xdc, 0xa1, 0x05, 0xbe, 0x12, 0x24, 0x5d, 0x65, 0x23,
	0xf0, 0xd6, 0x49, 0x05, 0x8d, 0x2b, 0x94, 0xe1, 0x91, 0x20, 0xe9, 0x3c, 0x1b, 0x81, 0xb7, 0xfd,
	0xaf, 0xbb, 0xab, 0x39, 0x0d, 0xdf, 0x1e, 0x81, 0xdf, 0xdd, 0x87, 0x70, 0xdd, 0xba, 0xdb, 0x3d,
	0x41, 0xec, 0x9c, 0x2e, 0x34, 0xd6, 0xbe, 0x1c, 0x07, 0xd9, 0xa7, 0xeb, 0xd3, 0xf7, 0x36, 0x21,
	0x1f, 0x6d, 0x42, 0x3e, 0xdb, 0x84, 0xbc, 0x7d, 0x25, 0xff, 0xca, 0x45, 0x78, 0x9a, 0xdd, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0xf5, 0xac, 0x9b, 0xab, 0x01, 0x00, 0x00,
}

func (m *Web3CommonMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Web3CommonMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Web3CommonMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.NodeId) > 0 {
		i -= len(m.NodeId)
		copy(dAtA[i:], m.NodeId)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.NodeId)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.MessageType) > 0 {
		i -= len(m.MessageType)
		copy(dAtA[i:], m.MessageType)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.MessageType)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.MessageId) > 0 {
		i -= len(m.MessageId)
		copy(dAtA[i:], m.MessageId)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.MessageId)))
		i--
		dAtA[i] = 0x52
	}
	if m.Timestamp != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x48
	}
	if m.NeedStore {
		i--
		if m.NeedStore {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.CipherSuite) > 0 {
		i -= len(m.CipherSuite)
		copy(dAtA[i:], m.CipherSuite)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.CipherSuite)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.PayloadType) > 0 {
		i -= len(m.PayloadType)
		copy(dAtA[i:], m.PayloadType)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.PayloadType)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.FromSign) > 0 {
		i -= len(m.FromSign)
		copy(dAtA[i:], m.FromSign)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.FromSign)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ComeFrom) > 0 {
		i -= len(m.ComeFrom)
		copy(dAtA[i:], m.ComeFrom)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.ComeFrom)))
		i--
		dAtA[i] = 0x22
	}
	if m.Version != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContentTopic) > 0 {
		i -= len(m.ContentTopic)
		copy(dAtA[i:], m.ContentTopic)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.ContentTopic)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Web3CommonMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.ContentTopic)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovMessage(uint64(m.Version))
	}
	l = len(m.ComeFrom)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.FromSign)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.PayloadType)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.CipherSuite)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.NeedStore {
		n += 2
	}
	if m.Timestamp != 0 {
		n += 1 + sovMessage(uint64(m.Timestamp))
	}
	l = len(m.MessageId)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.MessageType)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Web3CommonMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Web3CommonMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Web3CommonMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentTopic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentTopic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComeFrom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ComeFrom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromSign", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromSign = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayloadType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CipherSuite", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CipherSuite = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NeedStore", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NeedStore = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)