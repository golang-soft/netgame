package main

import (
	"net"
	"netgame/base/gnet"
	"netgame/base/log"
)

var serverManager *ServerManager

type Service struct {
	gnet.NetService
}

func NewService() *Service {
	server := &Service{}
	server.Derived = server

	return server
}

func (server *Service) Init() bool {

	ret := server.Bind("test", "", config.GetInt("port"))
	if !ret {
		log.Println("bid port error,service run is error")
		return false
	}

	serverManager = NewServerManager()

	return true
}

func (server *Service) Final() {

}

func (server *Service) NewTCPTask(conn net.Conn, port int) {

	task := NewServerTask()
	task.GoHandler(conn)
}
