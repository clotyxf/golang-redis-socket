package main

import (
	"config"
	"global"
	"http/controller"
	"log"
	"net/http"
	"route"
	"socket"
)

func main() {
	//读取配置
	configFile := "config/env.yml"
	config.YmlFileRead(configFile)
	var localhost string

	global.App.Host = config.YamlConfig.Get("listen.host").String()
	global.App.Port = config.YamlConfig.Get("listen.port").String()
	log.Printf("%v", "start socket connect")
	localhost = global.App.Host + ":" + global.App.Port

	hub := socket.NewHub()

	go socket.Dial(hub)

	go hub.Run()
	controller.RegisterRoutes(hub)
	//config read

	log.Printf("%v", localhost)

	s := "loadinig..."
	log.Printf("%v", s)

	log.Fatal(http.ListenAndServe(localhost, route.DefaultMux))
}
