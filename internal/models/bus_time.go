package models

import "gorm.io/gorm"

type BusTime struct {
	gorm.Model
	NumberLine int `gorm:"column:NUMERO_LINHA"` //chave externa
	stop string `gorm:"column:pontos"`
	day string `gorm:"column:dia"`
	time string `gorm:"column:horarios"`
}
