package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"qpdatagather/dataparser/diopter/tianle"
	"qpdatagather/entity"
	"qpdatagather/log"
	"qpdatagather/util"
	"qpdatagather/validator"
)

type deviceCreatePayload struct {
	AppId   int64      `json:"appId"`
	Project string     `json:"project"`
	Type    util.Type  `json:"type" binding:"required"`
	Brand   util.Brand `json:"brand" binding:"required"`
	Model   util.Model `json:"model" binding:"required"`
	OriData string     `json:"oriData" binding:"required"`
}

func deviceParser(c *gin.Context) {
	var payload deviceCreatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		util.ResponseErrf(c, "请求错误：%s", validator.Translate(err))
		return
	}

	var typeT = payload.Type
	var brand = payload.Brand
	var model = payload.Model
	var oriData = payload.OriData
	oriDataByteSlice, _ := base64.StdEncoding.DecodeString(oriData)
	//var str = string(oriDataByteSince)
	//var version = 0
	//var imageUrl = ""

	var result any
	switch typeT {
	case util.Optometry, "验光仪":
		switch brand {
		case util.Tianle, "天乐":
			switch model {
			case util.KR9800:
				result = tianle.Kr9800Parse(oriDataByteSlice)
			}
		case util.Faliao, "法里奥":

		}
	case util.Biometer, "生物测量仪":
	case util.Tonometer, "眼压计":
	}

	device := &entity.Device{}
	payloadJson, _ := json.Marshal(payload)
	_ = json.Unmarshal(payloadJson, device)
	if result != nil {
		resultJson, _ := json.Marshal(result)
		device.ParData = string(resultJson)
		device.Status = 100
		log.Info(device.ToString())
		util.ResponseSuccess(c, result)
	} else {
		log.Info(device.ToString())
		util.ResponseErr(c, "解析失败")
	}
}
