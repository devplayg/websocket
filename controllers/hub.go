package controllers

import (
	"github.com/astaxie/beego"
)

var hub *Hub

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
		case client := <-h.register:
			beego.Info("Registered: " + client.name)
			h.clients[client] = true
			//			client.hub.broadcast <- []byte(client.name + " has joined at the chat")
			for client := range h.clients {
				select {
				case client.send <- []byte(client.name + " has joined at the chat"):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}

	}
}

func init() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.AppName = "websocket"

	beego.Info("Starting websocket..")
	hub = newHub()
	go hub.run()
}
