package api

import (
	"fmt"
	"testing"
)

func TestDataParse(t *testing.T) {
	payload := deviceCreatePayload{
		Type:    "验光仪",
		Brand:   "美沃",
		Model:   "v100",
		OriData: "UE9TVCAvMTkyLjE2OC4xLjE6OTEwMCBIVFRQLzEuMQ0KQ29udGVudC1UeXBlOiBhcHBsaWNhdGlvbi94LXd3dy1mb3JtLXVybGVuY29kZWQNCmNoYXJzZXQ6IFVURi04DQpob3N0OiAxOTIuMTY4LjEuMTo5MTAwDQpjb250ZW50LWxlbmd0aDogMjQyDQpDb25uZWN0aW9uOiBjbG9zZQ0KDQpCYXJDb2RlPXVuZGVmaW5lZCZPcHRvbWVyeT17Ik1vZGFsaXR5VHlwZSI6Ik1XLVYxMDAiLCJTcGhlcmVMZWZ0IjoiMC4wIiwiU3BoZXJlUmlnaHQiOiIlMkIwLjMiLCJTcGhlcmVwZCI6IjU3IiwiQ3lsaW5kZXJMZWZ0IjoiLTAuMyIsIkN5bGluZGVyUmlnaHQiOiItMC41IiwiQmVpemh1TGVmdCI6IiIsIkF1dG9LdWJ1blJpZ2h0IjoiMiIsIkJlaXpodVJpZ2h0IjoiIiwiQXhpc0xlZnQiOiIyMiIsIkF4aXNSaWdodCI6IjMifQ==",
	}
	result := deviceTest(payload)
	fmt.Println(result)
}
