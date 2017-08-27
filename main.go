package main

import (
	"github.com/astaxie/beego"
	_ "github.com/devplayg/websocket/routers"
)

func main() {
	beego.Run()
}
