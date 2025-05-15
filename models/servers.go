package models

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	Name     string
	Ip       string
	Location string
  HostName string 
  SSHPort int `json:"sshPort" validate:"required,min=1,max=65535"` 
  Owner string
}
