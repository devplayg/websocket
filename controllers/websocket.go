package controllers

import (
	"github.com/astaxie/beego"
)

type WebsocketController struct {
	beego.Controller
}

func (this *WebsocketController) Join() {
	this.Data["username"] = this.GetString("username")
	this.TplName = "chat.html"
}

func (this *WebsocketController) OpenSocket() {
}

func init() {
	beego.Info("Starting websocket..")
	hub := newHub()
	go hub.run()
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case _ = <-h.register:
			beego.Info("Registerd")
		case _ = <-h.unregister:
			beego.Info("Unregisterd")
		case _ = <-h.broadcast:
			beego.Info("Message")
		}
		beego.Info("wait")
	}
}
