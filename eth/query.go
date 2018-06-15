package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"fmt"
	"eth-contract/eth/releaseToken"
)

func GetBalance() string {
	//Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial("/Users/liuwei/Documents/chain/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	token, err := releaseToken.NewMyToken(common.HexToAddress("0xb4532be4ec76a67f4bc6ed207b24155adbdbc4a5"), conn)//合约地址
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	contractName, err := token.Name(nil)
	if err != nil {
		log.Fatalf("query name err:%v", err)
	}
	fmt.Printf("MyToken Name is:%s\n", contractName)
	//balance, err := token.BalanceOf(nil, common.HexToAddress("0xd31953ac085fe7bb5169365461d99d54abcf76c3"))//合约发布到的账户地址，查询账户1的余额
	balance, err := token.BalanceOf(nil, common.HexToAddress("0x822d83926bc0d775a8594e0bda8b2260cf015260"))//合约发布到的账户地址，查询账户2的余额
	if err != nil {
		log.Fatalf("query balance error:%v", err)
	}
	fmt.Printf("0xd31953ac085fe7bb5169365461d99d54abcf76c3's balance is %s\n", balance)
	//todo 需要转换单位为wei
	return balance.String()
}