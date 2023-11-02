package ioriver

import "fmt"

type Service struct {
	Id          string `json:"id,omitempty"`
	Account     string `json:"account,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Certificate string `json:"certificate"`
	ServiceUid  string `json:"service_uid,omitempty"`
	Cname       string `json:"cname,omitempty"`
}

const servicesBasePath = "services/"

func (client *IORiverClient) GetService(id string) (*Service, error) {
	path := fmt.Sprintf("%s%s/", servicesBasePath, id)
	return Get[Service](client, path)
}

func (client *IORiverClient) ListServices() ([]Service, error) {
	return List[Service](client, servicesBasePath)
}

func (client *IORiverClient) CreateService(newService Service) (*Service, error) {
	return Create[Service](client, servicesBasePath, newService)
}

func (client *IORiverClient) UpdateService(service Service) (*Service, error) {
	path := servicesBasePath + service.Id + "/"
	return Update[Service](client, path, service)
}

func (client *IORiverClient) DeleteService(serviceId string) error {
	path := fmt.Sprintf("%s%s/", servicesBasePath, serviceId)
	return Delete(client, path)
}
