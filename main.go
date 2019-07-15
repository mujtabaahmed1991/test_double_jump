package main

// go:generate abigen --bin=contracts/build/DJTERC721Implementation.bin --abi=contracts/build/DJTERC721Implementation.abi --pkg=contracts --out=contracts/erc721.go --type DJTERC721Implementation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gravitational/configure"
	"github.com/test_double_jump/contracts"
)

type Config struct {
	Endpoint string `env:"GETH_ENDPOINT" cli:"geth_endpoint" yaml:"Geth_Endpoint"`
	KeyFile  string `env:"KEY_FILE" cli:"key_file" yaml:"key_file"`
}

type Key struct {
	Address string `json:"address"`
}

func readConfig() Config {
	var config Config
	configure.ParseEnv(&config)
	configure.ParseCommandLine(&config, os.Args[1:])
	if config.Endpoint == "" || config.KeyFile == "" {
		log.Fatal(`Usage go run main.go --geth_endpoint=<GETH> --token_hash=<ContractAddress> --key_file=<AccountKeyJSON>
		           or set env variable GETH_ENDPOINT, TOKEN_HASH`)
	}
	return config
}

func getConnection(endpoint, tokenHash string) (*ethclient.Client, *contracts.DJTERC721Implementation) {
	conn, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatalf("Filed to connect to ethereum client: %v", err)
	}
	token, err := contracts.NewDJTERC721Implementation(common.HexToAddress(tokenHash), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	return conn, token
}

func getHandler(endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()
		tokenId, err := strconv.Atoi(queryValues.Get("tokenId"))
		tokenHash := queryValues.Get("contractAddress")
		_, token := getConnection(endpoint, tokenHash)
		if err != nil {
			fmt.Fprintf(w, "Invalid tokenId:"+queryValues.Get("tokenId"))
			fmt.Println(queryValues.Get("tokenId"))
			log.Fatalf("Invalid tokenId")
		}
		address, err := token.OwnerOf(nil, big.NewInt(int64(tokenId)))
		if err != nil {
			fmt.Fprintf(w, `{"Error": "%v"}`, err)
		} else {
			fmt.Fprintf(w, `{"owner": "%s"}`, address.String())
		}
	}
}

func main() {
	cfg := readConfig()
	key, err := ioutil.ReadFile(cfg.KeyFile)
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}
	var address Key
	if err = json.Unmarshal(key, &address); err != nil {
		log.Fatalf("Invalid JSON: %s error: v%", string(key), err)
	}
	http.HandleFunc("/Get", getHandler(cfg.Endpoint))
	http.HandleFunc("/Add", func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()
		_, token := getConnection(cfg.Endpoint, queryValues.Get("contractAddress"))
		auth, err := bind.NewTransactor(strings.NewReader(string(key)), "")
		if err != nil {
			log.Fatalf("Failed to create authorized transactor: %v", err)
		}
		_, err = token.Mint(auth) //, common.HexToAddress("0x0d095f25c488c960341d40b63bf726a6fe88b48f"), common.HexToAddress("0x2cc9f14b021e871519a53b890f20b4b7cd0cfa69"), big.NewInt(1))
		if err != nil {
			log.Fatalf("Failed to request token transfer: %v", err)
		}
		token.Approve(auth, common.HexToAddress(address.Address), big.NewInt(1))
		fmt.Fprintf(w, "Added to token waiting for approval")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
