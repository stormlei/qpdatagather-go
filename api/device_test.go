package api

import (
	"fmt"
	"testing"
)

func TestDataParse(t *testing.T) {
	payload := deviceCreatePayload{
		Type:    "验光仪",
		Brand:   "tianle",
		Model:   "rm9000",
		OriData: "UyBObz0wMDAwMCBSIEIgRVMtUi1SIC0wMS41MCAtMDAuMjUgICs5MiA1OSBFUy1SLUwgLTAwLjc1ICswMC4wMCAgKzAwIDU5IEU=",
	}
	result := deviceTest(payload)
	fmt.Println(result)
}
