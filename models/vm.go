package models

import "gorm.io/gorm"

type VM struct {
	gorm.Model
	name     string
	os       string
	serverId uint
}
