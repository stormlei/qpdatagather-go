package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"qpdatagather/api"
)

func NewRouter(conf *viper.Viper) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routeGroup := router.Group("")
	{
		api.Routes(routeGroup)
	}

	return router
}
