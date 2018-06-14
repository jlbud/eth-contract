package testCase

import (
	"testing"
	"token-contract/eth"
)

func Test_getBalance(t *testing.T) {
	balance := eth.GetBalance()
	t.Log("balance is " ,balance)
}