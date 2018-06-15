package handler

import (
	"net/http"
	"eth-contract/web/util"
	"eth-contract/web/model/response"
	"eth-contract/eth"
	util_main "eth-contract/util"
	"fmt"
)

func GetEthBalance(w http.ResponseWriter, r *http.Request) {
	respEthBlance := response.RespEth{}
	respEthBlance.Data.Balance = eth.GetBalance()
	respEthBlance.Success = "0"
	respEthBlance.Message = "pass"
	util.RespondJSON(w, http.StatusAccepted, respEthBlance)
}

func GetEthBalanceWeb(w http.ResponseWriter, r *http.Request){
	balance := eth.GetBalance()
	codeStr := util_main.GenerateQRCodeString(balance)
	//write html
	ls := "<!DOCTYPE html><head><meta charset='utf-8' /></head><body><img src='data:image/png;base64," + codeStr + "'/></body></html>"
	fmt.Fprintf(w, ls)
}
