package controller

import (
	"fuego_backend/models"

	"github.com/go-fuego/fuego"
	"gorm.io/gorm"
)

type Server struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Ip       string `json:"ip"`
	Location string `json:"location"`
  SSHPort int  `json:"sshPort"`
  HostName string `json:"hostName"`
  Owner string `json:"owner"` 

}
type ServerCreate struct {
	Name     string `json:"name"`
	Ip       string `json:"ip"`
	Location string `json:"location"`
  SSHPort int `json:"sshPort"`
  HostName string `json:"hostName"`
  Owner string `json:"owner"`
}

type ServerUpdate struct {
	Name     *string `json:"name"`
	Ip       *string `json:"ip"`
	Location *string `json:"location"`
  SSHPort *int `json:"sshPort"`
  HostName *string `json:"hostName"`
  Owner *string `json:"owner"`
}
type ServerService interface {
  GetServer(id string) (Server, error)
	GetAllServer() ([]Server, error)
	UpdateServer(id string, input ServerUpdate) (Server, error)
	DeleteServer(id string) (any, error)
	CreateServer(input ServerCreate) (Server, error)
}

type PersistentServerService struct {
	db *gorm.DB
}

// linking the interface with the DB instance
func NewPersistentServerService(db *gorm.DB) ServerService {
	return &PersistentServerService{db: db}
}

func (s *PersistentServerService) GetServer(id string) (Server, error) {
	var serverModel models.Server
	serverInstance := s.db.First(&serverModel, id)

	if serverInstance.Error != nil {
		return Server{}, serverInstance.Error
	}

  return Server{ID: serverModel.ID, Name: serverModel.Name, Ip: serverModel.Ip, Location: serverModel.Location, SSHPort: serverModel.SSHPort, HostName: serverModel.HostName, Owner: serverModel.Owner}, nil


}

func (s *PersistentServerService) GetAllServer() ([]Server, error) {
	var servers []Server
	var results []Server

	instances := s.db.Find(&servers)
	if instances.Error != nil {
		return results, instances.Error
	}

	for _, server := range servers {
    results = append(results, Server{ID: server.ID, Name: server.Name, Ip: server.Ip, Location: server.Location, SSHPort: server.SSHPort, HostName: server.HostName, Owner: server.Owner})
	}

	return results, nil
}

func (s *PersistentServerService) CreateServer(input ServerCreate) (Server, error) {
	instance := models.Server{Name: input.Name, Location: input.Location, Ip: input.Ip}
	s.db.Create(&instance)
  return Server{ID: instance.ID, Name: instance.Name, Location: instance.Location, Ip: instance.Ip, SSHPort: instance.SSHPort, HostName: instance.HostName, Owner: instance.Owner}, nil

}

func (s *PersistentServerService) UpdateServer(id string, input ServerUpdate) (Server, error) {
	var serverModel models.Server
	serverInstance := s.db.First(&serverModel, id)

	if serverInstance.Error != nil {
		return Server{}, serverInstance.Error
	}

	if input.Name != nil {
		serverModel.Name = *input.Name
	}
	if input.Ip != nil {
		serverModel.Ip = *input.Ip
	}
	if input.Location != nil {
		serverModel.Location = *input.Location
	}
  if input.SSHPort != nil {
    serverModel.SSHPort = *input.SSHPort
  }
  if input.HostName != nil {
    serverModel.HostName = *input.HostName
  }
  if input.Owner != nil {
    serverModel.Owner = *input.Owner
  }

	s.db.Save(&serverModel)
  return Server{ID: serverModel.ID, Name: serverModel.Name, Location: serverModel.Location, Ip: serverModel.Ip, SSHPort: serverModel.SSHPort, HostName: serverModel.HostName, Owner: serverModel.Owner}, nil

}
func (s *PersistentServerService) DeleteServer(id string) (any, error) {
	s.db.Delete(&models.Server{}, id)
	return "Resource deleted", nil
}

type ServerResources struct {
	ServerService ServerService
}

// linking handlers with fuego router
func (r ServerResources) GetAllServer(c fuego.ContextNoBody) ([]Server, error) {
	return r.ServerService.GetAllServer()
}

func (r ServerResources) GetServer(c fuego.ContextNoBody) (Server, error) {
	serverId := c.PathParam("id")
	return r.ServerService.GetServer(serverId)
}

func (r ServerResources) CreateServer(c fuego.ContextWithBody[ServerCreate]) (Server, error) {
	input, err := c.Body()
	if err != nil {
		return Server{}, err
	}
	return r.ServerService.CreateServer(input)
}

func (r ServerResources) UpdateServer(c fuego.ContextWithBody[ServerUpdate]) (Server, error) {
	input, error := c.Body()
	if error != nil {
		return Server{}, error
	}
	id := c.QueryParam("id")
	return r.ServerService.UpdateServer(id, input)
}

func (r ServerResources) DeleteServer(c fuego.ContextNoBody) (any, error) {
	id := c.QueryParam("id")
	return r.ServerService.DeleteServer(id)
}

func (r ServerResources) Routes(s *fuego.Server) {
	fuego.Get(s, "/servers", r.GetAllServer)
	fuego.Get(s, "/servers/:id", r.GetServer)
	fuego.Post(s, "/servers", r.CreateServer)
	fuego.Patch(s, "/servers/:id", r.UpdateServer)
	fuego.Delete(s, "/servers/:id", r.DeleteServer)
}
