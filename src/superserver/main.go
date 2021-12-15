package main

import (
	"flag"
	"netgame/base/log"
	"netgame/base/util"
)

var config *util.Config

func main() {

	var logfile string

	flag.StringVar(&logfile, "logfile", "", "日志文件路径")
	flag.Parse()

	log.NewLog(logfile)

	var err error = nil
	config, err = util.NewConfig("config.json", "superserver")
	if err != nil {
		log.Errorln("读取config.json失败")
		return
	}

	server := NewService()
	server.Run()

}
