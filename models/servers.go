package models

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	name     string
	ip       string
	location string
}
