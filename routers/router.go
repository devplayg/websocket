package routers

import (
	"github.com/astaxie/beego"
	"github.com/iwondory/websocket/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/join", &controllers.WebsocketController{}, "post:Join")
	beego.Router("/ws", &controllers.WebsocketController{}, "get:OpenSocket")
}
