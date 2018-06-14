package main

import "token-contract/web"

func main(){
	app := &web.App{}
	app.Initialize()
	app.Run(":8080")
}