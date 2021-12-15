package main

import (
	"github.com/golang/protobuf/proto"
	"netgame/base/gnet"
	"netgame/base/log"
	"netgame/command"
)

type SuperClient struct {
	gnet.TCPClient

	msgHandler gnet.MessageHandler
}

func NewSuperClient() *SuperClient {
	client := &SuperClient{}
	client.Derived = client

	if !client.Connect(config.GetString("super_ip"), config.GetInt("super_port")) {
		return nil
	}

	client.init()

	return client
}

func (client *SuperClient) MsgParse(msg *command.Message) bool {

	log.Println(msg)

	client.msgHandler.Process(msg)

	return true
}

func (client *SuperClient) OnConnected() {

	msg := new(command.ReqServerVerify)
	msg.Info = service.GetServerInfo()
	client.SendCmd(msg)
}

func (client *SuperClient) init() {

	client.msgHandler.Reg(&command.NotifyRouteServerInit{}, client.onNotifyRouteServerInit)
	client.msgHandler.Reg(&command.NotifyRouteServerAdd{}, client.onNotifyRouteServerAdd)

	client.msgHandler.Reg(&command.RetServerVerify{}, client.onRetServerVerify)

}

func (client *SuperClient) onNotifyRouteServerInit(cmd proto.Message) {

	msg := cmd.(*command.NotifyRouteServerInit)
	routeManager.InitRouteList(msg.Serverlist)
}

func (client *SuperClient) onNotifyRouteServerAdd(cmd proto.Message) {
	msg := cmd.(*command.NotifyRouteServerAdd)
	routeManager.Add(msg.Info)
}

func (client *SuperClient) onRetServerVerify(cmd proto.Message) {
	log.Println("服务器校验完成，请求获取网关数据")
}
