package global

import (
	"net"
	"sync"
	"time"
)

type app struct {
	BuildTime time.Time
	Version   string

	Host string
	Port string

	locker sync.Mutex

	Con net.Conn
}

var App = &app{}
var buffServer = make([]byte, 1024)
var buffClient = make([]byte, 1024)
var clients = make(map[string]net.Conn)
var messages = make(chan string, 10)
