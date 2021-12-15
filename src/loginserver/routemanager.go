package main

import (
	"github.com/golang/protobuf/proto"
	"netgame/base/gnet"
	"netgame/base/log"
	"netgame/command"
)

type RouteManager struct {
	gnet.RouteManager

	msgHandler gnet.MessageHandler
}

func NewRouteManager() *RouteManager {
	mgr := &RouteManager{}

	mgr.init()

	return mgr
}

func (mgr *RouteManager) GetServerInfo() *command.ServerInfo {
	return service.GetServerInfo()
}

func (mgr *RouteManager) MsgParse(msg *command.Message) bool {

	//	log.Println("route manager:", msg)

	mgr.msgHandler.Process(msg)

	return true
}

func (mgr *RouteManager) init() {

	mgr.Derived = mgr

	mgr.msgHandler.Reg(&command.UpdateGatewayOnline{}, mgr.onUpdateGatewayOnline)
}

func (mgr *RouteManager) onUpdateGatewayOnline(cmd proto.Message) {
	msg := cmd.(*command.UpdateGatewayOnline)

	gatewayManager.Update(msg.Id, msg.Online)

	log.Println("更新网关在线", msg.Id, msg.Online)
}
