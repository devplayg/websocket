package controllers

import (
	"github.com/astaxie/beego"
)

type WebsocketController struct {
	beego.Controller
}

func (this *WebsocketController) Signin() {
	this.Data["httpport"] = beego.AppConfig.String("httpport")
	this.TplName = "websocket.html"
}

func (this *WebsocketController) OpenSocket() {
	conn, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
		return
	}
	client := &Client{hub: hub, conn: conn, name: this.GetString("username"), send: make(chan []byte, 256)}
	this.SetSession("username", this.GetString("username"))
	client.hub.register <- client
	go client.writePump()
	go client.readPump()
	beego.Info("Registered: " + this.GetString("username"))
	this.ServeJSON()
}
