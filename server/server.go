package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qpdatagather/config"
	"qpdatagather/log"
)

func Init() {
	conf := config.GetConfig()

	//set gin mode
	if conf.GetString("profile") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := NewRouter(conf)

	if err := r.Run(fmt.Sprintf(":%s", conf.GetString("service.port"))); err != nil {
		log.Errorf("error start server: %v", err)
		return
	}
}
