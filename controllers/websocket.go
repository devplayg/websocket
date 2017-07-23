package controllers

import (
	"github.com/astaxie/beego"
)

type WebsocketController struct {
	beego.Controller
}

func (this *WebsocketController) Signin() {
	this.Data["httpport"] = beego.AppConfig.String("httpport")
	var username string
	if this.GetSession("username") != nil {
		username = this.GetSession("username").(string)
	}
	this.Data["username"] = username
	this.TplName = "websocket.html"
}

func (this *WebsocketController) OpenSocket() {
	conn, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
		return
	}
	client := &Client{hub: hub, conn: conn, name: this.GetString("username"), send: make(chan []byte, 256), sessionId: this.Ctx.GetCookie(beego.BConfig.WebConfig.Session.SessionName)}
	this.SetSession("username", this.GetString("username"))
	client.hub.register <- client
	go client.writePump()
	go client.readPump()
	this.ServeJSON()
}
