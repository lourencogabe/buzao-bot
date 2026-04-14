package models

import "gorm.io/gorm"

type BusLine struct {
    gorm.Model
    Number      int `gorm:"column:NUMERO_LINHA"`
    Description string `gorm:"column:DESC_LINHA"`
    UrlUrbs     string `gorm:"column:URL_URBS"`
    UrlGazeta   string `gorm:"column:URL_GAZETA"`
    UrlMoovit   string `gorm:"column:URL_MOOVIT"`
}
