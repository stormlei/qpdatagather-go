package api

import "github.com/gin-gonic/gin"

func Routes(route *gin.RouterGroup) {
	//route.POST("/deviceParse", deviceParse)

	deviceGroup := route.Group("/device")
	{
		deviceV1 := deviceGroup.Group("/v1")
		{
			deviceV1.POST("/device-parse", deviceParse)
			deviceV1.POST("/ble-mac", bleMacCreate)
		}
	}

}
