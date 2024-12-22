package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Id     string `gorm:"type:uuid;default:gen_random_uuid()"`
	Title  string `gorm:"not null"`
	amount int    `gorm:"not null"`
	Date   string `gorm:"type:date;not null"`
	Tags   []Tag  `gorm:"many2many:transaction_tags;"`
}

type Tag struct {
	gorm.Model
	Id   string `gorm:"type:uuid;default:gen_random_uuid()"`
	Name string `gorm:"not null"`
}
