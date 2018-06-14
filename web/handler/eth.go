package handler

import (
	"net/http"
	"token-contract/web/util"
	"token-contract/web/model/response"
	"token-contract/eth"
	util_main "token-contract/util"
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
