package configs

import (
	"context"
	"fmt"
	"log"

  "math"
  "math/big"

	"github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/common"
)

var goerli = "https://eth-goerli.g.alchemy.com/v2/iUbl0BOfRwQX_IvKKpk22O2xH6HL95lk"

func ConnectGoerli() *ethclient.Client {
	client, err := ethclient.DialContext(context.Background(), goerli)
	if err != nil {
		log.Fatal("Failed connect to goerli ethereum node ", err)
	}

	defer client.Close()
	fmt.Println("[NOWTR] connected to goerli node")
	return client
}

var GOERLI *ethclient.Client = ConnectGoerli()

func GetBalance(client *ethclient.Client,address string) *big.Float {
  addr := common.HexToAddress(address)
  balance, err := client.BalanceAt(context.Background(), addr, nil)
  if err != nil {
    log.Fatal(err)
  }

  // balance in eth value
  fbal := new(big.Float)
  fbal.SetString(balance.String())
  eth := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(18)))
  return eth
}