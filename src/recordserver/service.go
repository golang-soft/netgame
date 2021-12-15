package main

import (
	"gopkg.in/mgo.v2"
	"net"
	"netgame/base/db"
	"netgame/base/gnet"
	"netgame/base/log"
	"netgame/command"
)

var timetick *TimeTick
var superclient *SuperClient
var routeManager *RouteManager
var mongo *mgo.Session

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
	server.SetServerID(command.GetServerID(command.RecordServer, config.GetInt("server_index")))

	ret := server.Bind("recordserver", "", 0)
	if !ret {
		log.Println("bid port error,service run is error")
		return false
	}

	routeManager = NewRouteManager()

	superclient = NewSuperClient()
	if superclient == nil {
		log.Println("connect superserver is error")
		return false
	}

	mongo = db.NewMongo(config.GetString("mongdb"))
	if mongo == nil {
		log.Println("init mongo db error")
		return false
	}

	timetick = NewTimeTick()

	return true
}

func (server *Service) Final() {

}

func (server *Service) NewTCPTask(conn net.Conn, port int) {

}
