package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Id         string     `gorm:"type:uuid;default:gen_random_uuid()"`
	Title      string     `gorm:"not null" csv:"title"`
	Amount     int        `gorm:"not null"`
	Date       string     `gorm:"type:date;not null" csv:"date"`
	Categories []Category `gorm:"many2many:transaction_categories;"`
}

type Category struct {
	gorm.Model
	Id            string          `gorm:"type:uuid;default:gen_random_uuid()"`
	Name          string          `gorm:"unique;not null"`
	CategoryTerms []CategoryTerms `gorm:"foreignkey:CategoryID"`
}

type CategoryTerms struct {
	gorm.Model
	Id         string `gorm:"type:uuid;default:gen_random_uuid()"`
	Expression string `gorm:"not null"`
	CategoryID string `gorm:"not null"`
	Category   Category
}
