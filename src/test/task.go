package main

import (
	"netgame/base/gnet"
	"netgame/command"
)

type MyTask struct {
	gnet.TCPTask
}

func NewMyTask() *MyTask {
	task := &MyTask{}
	return task
}

func (task *MyTask) VerifyConn(msg *command.Message) bool {
	return true
}

func (task *MyTask) RecycleConn() bool {
	return true
}

func (task *MyTask) MsgParse(msg *command.Message) bool {
	return true
}
