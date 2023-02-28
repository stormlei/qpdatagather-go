package entity

import (
	"fmt"
	"qpdatagather/util"
)

type Device struct {
	AppId   int64      `json:"appId"`
	Project string     `json:"project"`
	Type    util.Type  `json:"type"`
	Brand   util.Brand `json:"brand"`
	Model   util.Model `json:"model"`
	OriData string     `json:"oriData"`
	ParData string     `json:"parData"`
	Status  int64      `json:"status"`
}

func (d *Device) ToString() string {
	return fmt.Sprintf("Device(appId=%d, project=%s, type=%s, brand=%s, model=%s, oriData=%s, "+
		"parData=%s, status=%d)", d.AppId, d.Project, d.Type, d.Brand, d.Model, d.OriData, d.ParData, d.Status)
}
