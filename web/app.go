package web

import (
	"log"
	"net/http"
	"eth-contract/web/handler"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) setRouters() {
	a.Get("/app/eth/balance", a.GetEthBalance)
	a.Get("/web/eth/balance_qrcode", a.GenerateQRCode)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) GetEthBalance(w http.ResponseWriter, r *http.Request) {
	handler.GetEthBalance(w, r)
}

func (a *App) GenerateQRCode(w http.ResponseWriter,r *http.Request){
	handler.GetEthBalanceWeb(w,r)
}
