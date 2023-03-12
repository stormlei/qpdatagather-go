package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	cachet "qpdatagather/cache"
	nidek2 "qpdatagather/dataparser/biometer/nidek"
	"qpdatagather/dataparser/biometer/suoer"
	"qpdatagather/dataparser/biometer/zeiss"
	"qpdatagather/dataparser/bloodpressure/yuwell"
	nidek4 "qpdatagather/dataparser/cem/nidek"
	faliao2 "qpdatagather/dataparser/diopter/faliao"
	"qpdatagather/dataparser/diopter/hand/meiwo"
	"qpdatagather/dataparser/diopter/nidek"
	"qpdatagather/dataparser/diopter/tianle"
	"qpdatagather/dataparser/diopter/topcon"
	"qpdatagather/dataparser/heightweight/lejia"
	"qpdatagather/dataparser/pyrometer/faliao"
	"qpdatagather/dataparser/pyrometer/jumu"
	nidek3 "qpdatagather/dataparser/tonometer/nidek"
	topcon2 "qpdatagather/dataparser/tonometer/topcon"
	"qpdatagather/dataparser/vitalcapacity/jihao"
	"qpdatagather/db"
	"qpdatagather/entity"
	"qpdatagather/enum"
	"qpdatagather/log"
	"qpdatagather/util"
	"qpdatagather/validator"
	"strings"
)

type deviceCreatePayload struct {
	Project string     `json:"project"`
	Type    enum.Type  `json:"type" binding:"required"`
	Brand   enum.Brand `json:"brand" binding:"required"`
	Model   enum.Model `json:"model" binding:"required"`
	OriData string     `json:"oriData" binding:"required"`
	Mac     string     `json:"mac" binding:"required"`
}

var result any

type jsonData struct {
	Version int64  `json:"version"`
	Gzip    bool   `json:"gzip"`
	Data    string `json:"data"`
}

func dataParser(payload deviceCreatePayload) {
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
			case enum.RM8900, enum.RM800, enum.KR800:
				result = topcon.TopconDataParse(oriDataByteSlice)
			case enum.CV5000:
				result = topcon.CV5000DataParse(oriDataByteSlice)
			}
		case enum.Faliao, "法里奥":
			switch model {
			case enum.FR8900:
				result = nil
			case enum.FR710, enum.FR900:
				result = faliao2.Fr710DataParse(oriDataByteSlice)
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
		switch brand {
		case enum.Nidek, "尼德克":
			switch model {
			case enum.ALScan:
				result = nidek2.ALScanV2DataParse(oriDataByteSlice)
			}
		case enum.Suoer, "索维":
			switch model {
			case enum.Sw9000:
				result = suoer.SW9000DataParse(oriDataByteSlice)
			}
		case enum.Zeiss, "蔡司":
			switch model {
			case enum.Iolmaster500:
				result = zeiss.IolMaster500DataParse(oriDataByteSlice)
			case enum.Iolmaster700:
				result = zeiss.IolMaster700DataParse(oriDataByteSlice)
			}
		}
	case enum.Tonometer, "眼压计":
		switch brand {
		case enum.Nidek, "尼德克":
			switch model {
			case enum.NT510:
				result = nidek3.NT5100DataParse(oriDataByteSlice)
			}
		case enum.Topcon, "拓普康":
			switch model {
			case enum.CT1:
				result = topcon2.CT1DataParse(oriDataByteSlice)
			case enum.CT800:
				result = topcon2.CT800DataParse(oriDataByteSlice)
			}
		}
	case enum.Cem, "角膜内皮镜":
		switch brand {
		case enum.Nidek, "尼德克":
			switch model {
			case enum.Cem530:
				result = nidek4.Cem530DataParse(oriDataByteSlice)
			}
		}
	case enum.Pyrometer, "焦度计":
		switch brand {
		case enum.Faliao, "法里奥":
			switch model {
			case enum.Fl800:
				result = faliao.FL800DataParse(oriDataByteSlice)
			}
		case enum.SuperVision, "巨目光学":
			switch model {
			case enum.LM260:
				result = jumu.LM260DataParse(oriDataByteSlice)
			}
		}
	case enum.HeightWeight, "身高体重秤":
		switch brand {
		case enum.Lejia, "乐佳":
			switch model {
			case enum.HWS7:
				result = lejia.HWS7DataParse(oriDataByteSlice)
			case enum.LJ700:
				result = lejia.LJ700DataParse(oriDataByteSlice)
			}
		}
	case enum.BloodPressure, "血压计":
		switch brand {
		case enum.Yuwell, "鱼跃":
			switch model {
			case enum.YE900:
				result = yuwell.YE900DataParse(oriDataByteSlice)
			}
		}
	case enum.VitalCapacity, "肺活量":
		switch brand {
		case enum.Jihao, "继豪":
			switch model {
			case enum.JH1663:
				result = jihao.JH1663DataParse(oriDataByteSlice)
			}
		}
	}
}

func deviceParse(c *gin.Context) {
	var payload deviceCreatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		util.ResponseErrf(c, "请求错误：%s", validator.Translate(err))
		return
	}

	if len(payload.Mac) != 17 && strings.Count(payload.Mac, ":") != 5 {
		util.ResponseErrf(c, "mac地址格式不正确")
		return
	}
	//1。验证数据库中有没有mac
	var macList []entity.BleMac
	db.GetDb().Where("mac = ?", payload.Mac).Find(&macList)
	if len(macList) == 0 {
		util.ResponseErr(c, "mac地址未通过认证")
		return
	}
	//2。验证缓存，3s过期
	mac, found := cachet.GetCache().Get("mac")
	if found {
		util.ResponseErrf(c, "mac地址%s未通过认证", mac) //3s内存在相同mac请求解析
		return
	}
	cachet.GetCache().Set("mac", payload.Mac, cachet.DefaultExpiration)

	dataParser(payload)

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

type macCreatePayload struct {
	Mac string `json:"mac" binding:"required"`
}

func bleMacCreate(c *gin.Context) {
	var payload macCreatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		util.ResponseErrf(c, "请求错误：%s", validator.Translate(err))
		return
	}

	if len(payload.Mac) != 17 && strings.Count(payload.Mac, ":") != 5 {
		util.ResponseErrf(c, "mac地址格式不正确")
		return
	}

	var macList []entity.BleMac
	db.GetDb().Where("mac = ?", payload.Mac).Find(&macList)
	if len(macList) > 0 {
		util.ResponseErr(c, "该mac地址已经存在")
		return
	}
	bleMac := &entity.BleMac{}
	bleMac.Mac = payload.Mac
	db.GetDb().Create(&bleMac)
	util.ResponseSuccess(c, bleMac)
}
