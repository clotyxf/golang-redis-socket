package socket

import (
	"bytes"
	"config"
	"log"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	conn redis.Conn
}

type Message struct {
	channel  string      "频道"
	data     interface{} "消息内容"
	uid      int         "通知用户"
	messType int         "消息类型"
}

type PMessage struct {
	channel string
	event   string
	data    []byte
}

var messInfo []byte

func Dial(h *Hub) {
	configFile := "config/env.yml"
	config.YmlFileRead(configFile)

	host := config.YamlConfig.Get("redis.host").String()
	port := config.YamlConfig.Get("redis.port").String()
	pwd := config.YamlConfig.Get("redis.pwd").String()

	localhost := host + ":" + port
	conn, err := redis.Dial("tcp", localhost)

	checkErr(err)

	_, err = conn.Do("AUTH", pwd)

	checkErr(err)

	rds := &Redis{conn: conn}

	go rds.Subscribe(h)
}

func (r *Redis) Subscribe(h *Hub) {
	psc := redis.PubSubConn{Conn: r.conn}

	psc.PSubscribe("*")

	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			log.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.PMessage:
			log.Printf("PMessage: %s %s %s\n", v.Pattern, v.Channel, v.Data)
			pm := new(PMessage)
			pm.channel = v.Channel
			pm.event = v.Pattern
			pm.data = v.Data
			h.broadcast <- pm
		case redis.Subscription:
			log.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			log.Println(v.Error())
		}
	}
}
