package main

import (
	"net"
	"netgame/base/gnet"
	"netgame/base/log"
	"netgame/command"
)

var timetick *TimeTick
var superclient *SuperClient
var routeManager *RouteManager
var gatewayManager *GatewayManager
var loginTaskManager *LoginTaskManager

type Service struct {
	gnet.NetService
}

func NewService() *Service {
	server := &Service{}
	server.Derived = server

	return server
}

func (server *Service) Init() bool {

	//初始化serverid
	server.SetServerID(command.GetServerID(command.LoginServer, config.GetInt("server_index")))

	ret := server.Bind("loginserver", "", config.GetInt("port"))
	if !ret {
		log.Println("bid port error,service run is error")
		return false
	}

	superclient = NewSuperClient()
	if superclient == nil {
		log.Println("connect superserver is error")
		return false
	}

	routeManager = NewRouteManager()
	loginTaskManager = NewLoginTaskManager()
	gatewayManager = NewGatewayManager()
	timetick = NewTimeTick()

	return true
}

func (server *Service) Final() {

}

func (server *Service) NewTCPTask(conn net.Conn, port int) {

	task := NewLoginTask()
	task.GoHandler(conn)
}
