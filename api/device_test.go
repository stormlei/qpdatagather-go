package api

import (
	"fmt"
	"testing"
)

func TestDataParse(t *testing.T) {
	payload := deviceCreatePayload{
		Type:    "生物测量仪",
		Brand:   "nidek",
		Model:   "al-scan",
		OriData: "AURybQJJRE5JREVLL0FSLTEXTk8zNTI5F0RBTUFSLzAzLzIwMjMuMTQ6NTEXVkQxMi4wMBdXRDQwF09MLTAzLjAwLTAwLjI1MTU3F09SLTAzLjUwLTAwLjAwMDAwF2RMLTAwLjI1LTAwLjAwLTA1F2RSLTAwLjI1KzAwLjI1KzU4FwFEUk0CSUROSURFSy9BUi0xF05PMzUyORdEQU1BUi8wMy8yMDIzLjE0OjUxF1ZEMTIuMDAXV0Q0MBdPTC0wMi43NS0wMC4yNTE2MhdPUi0wMy4yNS0wMC4yNTEyMjkgF1BENTg/Pz8/Pz8XBDMxQzE=",
	}
	result := deviceTest(payload)
	fmt.Println(result)
}
