package controllers

import (
	"github.com/astaxie/beego"
)

type WebsocketController struct {
	beego.Controller
}

func (this *WebsocketController) Get() {
	this.Data["httpport"] = beego.AppConfig.String("httpport")
	this.TplName = "websocket.html"
}

func (this *WebsocketController) OpenSocket() {
	conn, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	username := this.Ctx.GetCookie("username")
	beego.Info("### Username:" + username)

	if err != nil {
		beego.Error(err)
		return
	}
	client := &Client{hub: hub, conn: conn, name: username, send: make(chan []byte, 256)}
	client.hub.register <- client
	go client.writePump()
	go client.readPump()
	this.ServeJSON()
}
