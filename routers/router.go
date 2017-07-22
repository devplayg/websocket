package routers

import (
	"github.com/astaxie/beego"
	"github.com/iwondory/websocket/controllers"
)

func init() {
	beego.Router("/", &controllers.WebsocketController{}, "get:Signin")
	beego.Router("/ws", &controllers.WebsocketController{}, "get:OpenSocket")
}
