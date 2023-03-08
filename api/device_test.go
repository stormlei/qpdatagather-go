package api

import (
	"fmt"
	"testing"
)

func TestDataParse(t *testing.T) {
	payload := deviceCreatePayload{
		Type:    "heightWeight",
		Brand:   "lejia",
		Model:   "lj700",
		OriData: "VzA1NDQkSDE1OTAkUiQ=",
	}
	result := deviceTest(payload)
	fmt.Println(result)
}
