package response

type RespEth struct {
	RespModel
	Data RespEthBalance `json:"data"`
}

type RespEthBalance struct {
	Balance string `json:"balance"`
}
