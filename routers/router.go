package routers

import (
	"github.com/astaxie/beego"
	"github.com/devplayg/websocket/controllers"
)

func init() {
	beego.Router("/", &controllers.WebsocketController{}, "get:Signin")
	beego.Router("/ws", &controllers.WebsocketController{}, "get:OpenSocket")
}
