package api

import (
	"fmt"
	"testing"
)

func TestDataParse(t *testing.T) {
	payload := deviceCreatePayload{
		Type:    "bloodpressure",
		Brand:   "yuwell",
		Model:   "ye900",
		OriData: "FmQAOwBIAOUHCxAGOhNWAAQA",
	}
	result := deviceTest(payload)
	fmt.Println(result)
}
