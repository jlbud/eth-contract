package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"math/big"
	"fmt"
	"time"
	"context"
	"bytes"
	"eth-contract/eth/releaseToken"
)

const key  = `{"address":"d31953ac085fe7bb5169365461d99d54abcf76c3","crypto":{"cipher":"aes-128-ctr","ciphertext":"fe47a2dae5aee66b95ae8ddf377f651877825b888c89e7417db028d78f2ebc17","cipherparams":{"iv":"b0a9b8f40c39661126c13e0344f93de1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"cd49d49c20a624fbcc85a287dbe9b85e774ba6d5a1445dc6d6b7b0d6960c39c0"},"mac":"5601f161cc87130ac35deca184f99f8b1f5baa324f1e56a468fcf32edd8e5f21"},"id":"f200e315-eaf8-4a65-9a3c-f8c553174e23","version":3}`

func deploy() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial("/Users/liuwei/Documents/chain/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy a new awesome contract for the binding demo
	address, tx, token, err := releaseToken.DeployMyToken(auth, conn, big.NewInt(9651), "Contracts in Go!!!", 0, "Go!")
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	startTime := time.Now()
	ctx := context.Background()
	addressAfterMined, err := bind.WaitDeployed(ctx, conn, tx)//等待事物结束，挖矿开始结束等
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
	}
	fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))
	if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
		log.Fatalf("mined address :%s,before mined address:%s", addressAfterMined, address)
	}
	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)

}