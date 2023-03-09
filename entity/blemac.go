package entity

import "gorm.io/gorm"

type BleMac struct {
	gorm.Model
	Mac string `gorm:"type:varchar(20);not null;default:'';comment:蓝牙mac地址;" json:"mac"`
}
