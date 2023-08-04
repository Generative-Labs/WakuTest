# Waku test

## Dependency

- locust

```bash
pip install -r requirements.txt

cd src && go mod tidy
```


## waku test

```bash
# https://github.com/waku-org/go-waku

# Start a waku node first !important
```

**Write your waku node p2p address in `src/worker.go`**

```go
 // src/worker.go
var PeerHostList = []string{
    // write your boot waku nodes p2p address
}
```

### Master

```bash
locust --master -f dummy.py --master --master-bind-host=0.0.0.0 --master-bind-port=5557
```

### Worker


```bash
go run . --master-host=127.0.0.1 --master-port=5557
```
