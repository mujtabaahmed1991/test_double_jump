# Test Double JUMP

Implements ERC 721 smart contract protocol for ethereum block chain

## Build

```
1. go generate ./...
2. go build
```

## Enviornment Variable or CommandLine args

```
ENV:
    GETH_ENDPOINT: Ethereum node endpoint for example local geth <$GOPATH:/src/github.com/test_double_jump/chaindata/geth.ipc>
    KEY_FILE: Path for authentication json for acount
    
CommandLine:
    geth_endpoint: Ethereum node endpoint for example local geth <$GOPATH:/src/github.com/test_double_jump/chaindata/geth.ipc>
    key_file: Path for authentication json for acount
```

## Usage

```
go run main.go
```

### Rest endpoint
```
POST localhost:8080/Add/<ContractHash> 
```
Create tokens for smart contract

```
GET localhost:8080/Get/<ContractHash>/<TokenId>
```