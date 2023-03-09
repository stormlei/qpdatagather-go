package api

import "github.com/gin-gonic/gin"

func Routes(route *gin.RouterGroup) {
	route.POST("/deviceParser", deviceParser)

	deviceGroup := route.Group("/device")
	{
		deviceV1 := deviceGroup.Group("/v1")
		{
			deviceV1.POST("/deviceParser", deviceParser)
			deviceV1.POST("/bleMac", bleMacCreate)
		}
	}

}
