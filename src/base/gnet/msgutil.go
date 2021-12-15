package gnet

import (
	"github.com/golang/protobuf/proto"
	"netgame/base/util"
	"netgame/command"
)

func PackMessage(cmd proto.Message) *command.Message {
	msg := new(command.Message)
	msg.Name = proto.MessageName(cmd)
	msg.Data, _ = proto.Marshal(cmd)
	msg.Type = util.BKDRHash(msg.Name)

	return msg
}
