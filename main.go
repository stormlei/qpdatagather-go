package main

import (
	"flag"
	"fmt"
	"os"
	"qpdatagather/config"
	"qpdatagather/log"
	"qpdatagather/server"
	"qpdatagather/validator"
)

func main() {
	profile := flag.String("p", "prod", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -p {mode}")
		os.Exit(1)
	}
	flag.Parse()

	//初始化配置
	config.Init(*profile)

	//add config values
	conf := config.GetConfig()
	conf.Set("profile", *profile)

	log.InitLogger(conf)
	defer log.Sync()

	log.Debugf("profile = %s", *profile)

	validator.Init()
	server.Init()
}
