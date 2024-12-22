package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Id     string `gorm:"type:uuid;default:gen_random_uuid()"`
	Title  string `gorm:"not null"`
	Amount int    `gorm:"not null"`
	Date   string `gorm:"type:date;not null"`
	Tags   []Tag  `gorm:"many2many:transaction_tags;"`
}

type Tag struct {
	gorm.Model
	Id       string     `gorm:"type:uuid;default:gen_random_uuid()"`
	Name     string     `gorm:"unique;not null"`
	TagTerms []TagTerms `gorm:"foreignkey:TagID"`
}

type TagTerms struct {
	gorm.Model
	Id         string `gorm:"type:uuid;default:gen_random_uuid()"`
	Expression string `gorm:"not null"`
	TagID      string `gorm:"not null"`
	Tag        Tag
}
