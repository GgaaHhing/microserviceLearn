package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Mobile   string `gorm:"index:id_mobile; unique; varchar(11) ;not null"`
	Password string `grom:"type:varchar(64) ;not null"`
	NickName string `gorm:"type:varchar(32)"`
	Gender   string `gorm:"varchar(6); default: male"`
	Salt     string `gorm:"type: varchar(16)"`
	Role     int    `gorm:"type:int;default:1;comment'1-普通用户，2-管理员'"`
}
