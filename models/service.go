package models

import "gorm.io/gorm"

type Servier struct {
	gorm.Model
	name        string
	description string
	containerId string // foreign key to container
}
