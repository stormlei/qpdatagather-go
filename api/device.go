package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"qpdatagather/dataparser/diopter/hand/meiwo"
	"qpdatagather/dataparser/diopter/nidek"
	"qpdatagather/dataparser/diopter/tianle"
	"qpdatagather/dataparser/diopter/topcon"
	"qpdatagather/entity"
	"qpdatagather/enum"
	"qpdatagather/log"
	"qpdatagather/util"
	"qpdatagather/validator"
)

type deviceCreatePayload struct {
	AppId   int64      `json:"appId"`
	Project string     `json:"project"`
	Type    enum.Type  `json:"type" binding:"required"`
	Brand   enum.Brand `json:"brand" binding:"required"`
	Model   enum.Model `json:"model" binding:"required"`
	OriData string     `json:"oriData" binding:"required"`
}

var result any

type jsonData struct {
	Version int64  `json:"version"`
	Gzip    bool   `json:"gzip"`
	Data    string `json:"data"`
}

func dataParse(payload deviceCreatePayload) {
	var typeT = payload.Type
	var brand = payload.Brand
	var model = payload.Model
	var oriData = payload.OriData
	oriDataByteSlice, _ := base64.StdEncoding.DecodeString(oriData)
	var str = string(oriDataByteSlice)
	var version int64 = 0
	if isJSON(str) {
		jsonObj := &jsonData{}
		_ = json.Unmarshal([]byte(str), jsonObj)
		version = jsonObj.Version
		if version == 1 {
			var gzipT = jsonObj.Gzip
			var data = jsonObj.Data
			if gzipT {
				dataByteSlice, _ := base64.StdEncoding.DecodeString(data)
				oriDataByteSlice = unGzip(dataByteSlice)
			} else {
				oriDataByteSlice = []byte(data)
			}
		}
	}

	switch typeT {
	case enum.Optometry, "验光仪":
		switch brand {
		case enum.Nidek, "尼德克":
			switch model {
			case enum.Ark1:
				result = nidek.NidekDataParse(oriDataByteSlice)
			case enum.Ar330:
				result = nidek.Nidek330DataParse(oriDataByteSlice)
			case enum.RT5100:
				result = nidek.RT5100DataParse(oriDataByteSlice)
			}
		case enum.Topcon, "拓普康":
			switch model {
			case enum.RM8900:
				result = topcon.TopconDataParse(oriDataByteSlice)
			}
		case enum.Tianle, "天乐":
			switch model {
			case enum.KR9800:
				result = tianle.Kr9800DataParse(oriDataByteSlice)
			case enum.RM9000:
				result = tianle.Rm9000DataParse(oriDataByteSlice)
			}
		case enum.Meiwo, enum.Mediworks, "美沃":
			switch model {
			case enum.V100:
				result = meiwo.V100DataParse(oriDataByteSlice)
			}
		}
	case enum.Biometer, "生物测量仪":
	case enum.Tonometer, "眼压计":
	case enum.Cem, "角膜内皮镜":
	case enum.Pyrometer, "焦度计":
	}
}

func deviceParser(c *gin.Context) {
	var payload deviceCreatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		util.ResponseErrf(c, "请求错误：%s", validator.Translate(err))
		return
	}

	dataParse(payload)

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

func unGzip(byteSlice []byte) []byte {
	reader := bytes.NewReader(byteSlice)
	gzReader, err := gzip.NewReader(reader)
	if err != nil {
		log.Error(err) // Maybe panic here, depends on your error handling.
	}
	defer gzReader.Close()
	output, err := io.ReadAll(gzReader)
	if err != nil {
		log.Error(err)
	}
	return output
}

func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func deviceTest(payload deviceCreatePayload) any {
	dataParse(payload)

	device := &entity.Device{}
	payloadJson, _ := json.Marshal(payload)
	_ = json.Unmarshal(payloadJson, device)
	if result != nil {
		resultJson, _ := json.Marshal(result)
		device.ParData = string(resultJson)
		device.Status = 100
		return device.ParData
	} else {
		return nil
	}
}
