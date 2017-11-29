package controller

import (
	"config"
	"global"
	"log"
	"net/http"
	"route"
	"socket"
	//"golang.org/x/net/websocket"
)

type socketController struct{}

var hub *socket.Hub

func (this socketController) RegisterRoute() {
	route.HandleFunc("/socket", this.SocketClient)
}

func (socketController) SocketClient(w http.ResponseWriter, r *http.Request) {
	//读取配置
	configFile := "config/env.yml"
	config.YmlFileRead(configFile)
	global.App.Host = config.YamlConfig.Get("listen.host").String()
	global.App.Port = config.YamlConfig.Get("listen.port").String()
	log.Printf("%v", "start socket_client connect")
	log.Printf("socket_ip:%v", global.App.Host)
	log.Printf("socket_port:%v", global.App.Port)

	socket.ServeWs(hub, w, r)
}
