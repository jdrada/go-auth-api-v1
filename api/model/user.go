package model

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Email    string `gorm:"type:varchar(100);unique_index"`
    Password string `gorm:"size:100"`
}