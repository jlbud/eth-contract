package eth

import (
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"math/big"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"eth-contract/eth/releaseToken"
)

var Key  = `{"address":"d31953ac085fe7bb5169365461d99d54abcf76c3","crypto":{"cipher":"aes-128-ctr","ciphertext":"fe47a2dae5aee66b95ae8ddf377f651877825b888c89e7417db028d78f2ebc17","cipherparams":{"iv":"b0a9b8f40c39661126c13e0344f93de1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"cd49d49c20a624fbcc85a287dbe9b85e774ba6d5a1445dc6d6b7b0d6960c39c0"},"mac":"5601f161cc87130ac35deca184f99f8b1f5baa324f1e56a468fcf32edd8e5f21"},"id":"f200e315-eaf8-4a65-9a3c-f8c553174e23","version":3}`

func translate() {
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial("/Users/liuwei/Documents/chain/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	token, err := releaseToken.NewMyToken(common.HexToAddress("0xb4532be4ec76a67f4bc6ed207b24155adbdbc4a5"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	// Create an authorized transactor and spend 1 unicorn
	auth, err := bind.NewTransactor(strings.NewReader(Key), "123456")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	tx, err := token.Transfer(auth, common.HexToAddress("0x822d83926bc0d775a8594e0bda8b2260cf015260"), big.NewInt(1))//给0*822账户转了387个代币
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
}