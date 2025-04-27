package models

import "gorm.io/gorm"

type Container struct {
	gorm.Model
	name  string
	image string
	vmid  string // foreign key to VM
}
