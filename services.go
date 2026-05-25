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

func (client *IORiverClient) CreateServiceWithConfig(newService Service, serviceConfig ServiceConfig) (*Service, error) {
	path := servicesBasePath + "create_from_service_config/"
	req := ServiceCreateFromConfigRequest{
		Description:   newService.Description,
		CertificateID: newService.Certificate,
		ServiceConfig: serviceConfig.ConfigJSON,
	}

	resp, err := Create[ServiceCreateFromConfigResponse](client, path, req)
	if err != nil {
		return nil, err
	}

	service := Service{
		Id:          resp.ID,
		Account:     resp.Account,
		Name:        resp.Name,
		Description: resp.Description,
		Certificate: resp.Certificate,
		ServiceUid:  resp.ServiceUid,
		Cname:       resp.Cname,
	}

	return &service, nil
}

func (client *IORiverClient) UpdateService(service Service) (*Service, error) {
	path := servicesBasePath + service.Id + "/"
	return Update[Service](client, path, service)
}

func (client *IORiverClient) DeleteService(id string) error {
	path := fmt.Sprintf("%s%s/", servicesBasePath, id)
	return Delete(client, path)
}
