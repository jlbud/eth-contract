package testCase

import (
	"testing"
	"eth-contract/eth"
)

func Test_getBalance(t *testing.T) {
	balance := eth.GetBalance()
	t.Log("balance is " ,balance)
}