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
var gatewayTaskManager *GatewayTaskManager
var gateUserManager *GateUserManager

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
	server.SetServerID(command.GetServerID(command.GatewayServer, config.GetInt("server_index")))

	ret := server.Bind("gatewayserver", "", config.GetInt("port"))
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
	gatewayTaskManager = NewGatewayTaskManager()
	gateUserManager = NewGateUserManager()

	timetick = NewTimeTick()

	return true
}

func (server *Service) Final() {

	routeManager.Close()
}

func (server *Service) NewTCPTask(conn net.Conn, port int) {

	task := NewGatewayTask()
	task.GoHandler(conn)
}
