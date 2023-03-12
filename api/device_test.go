package api

import (
	"encoding/json"
	"fmt"
	"qpdatagather/entity"
	"testing"
)

func TestDataParse(t *testing.T) {
	payload := deviceCreatePayload{
		Type:    "bloodpressure",
		Brand:   "yuwell",
		Model:   "ye900",
		OriData: "FoYA/wf/B+UHCxANEBz/BxAAFoMA/wf/B+UHCxANEB3/BxAA",
	}
	result := deviceTest(payload)
	fmt.Println(result)
}

func deviceTest(payload deviceCreatePayload) any {
	dataParser(payload)

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
