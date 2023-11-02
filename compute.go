package ioriver

import (
	"fmt"
)

type ComputeRoute struct {
	Id        string `json:"id,omitempty"`
	ComputeId string `json:"compute,omitempty"`
	Domain    string `json:"domain"`
	Path      string `json:"path"`
}

type Compute struct {
	Id           string         `json:"id,omitempty"`
	Service      string         `json:"service"`
	Name         string         `json:"name"`
	RequestCode  string         `json:"request_code,omitempty"`
	ResponseCode string         `json:"response_code,omitempty"`
	Routes       []ComputeRoute `json:"compute_routes"`
}

const computesBasePath = `services/%s/compute/`

func (client *IORiverClient) GetCompute(serviceId string, id string) (*Compute, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(computesBasePath, serviceId), id)
	return Get[Compute](client, path)
}

func (client *IORiverClient) ListComputes(serviceId string) ([]Compute, error) {
	path := fmt.Sprintf(computesBasePath, serviceId)
	return List[Compute](client, path)
}

func (client *IORiverClient) CreateCompute(compute Compute) (*Compute, error) {
	path := fmt.Sprintf(computesBasePath, compute.Service)
	return Create[Compute](client, path, compute)
}

func (client *IORiverClient) UpdateCompute(compute Compute) (*Compute, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(computesBasePath, compute.Service), compute.Id)
	return Update[Compute](client, path, compute)
}

func (client *IORiverClient) DeleteCompute(serviceId string, computeId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(computesBasePath, serviceId), computeId)
	return Delete(client, path)
}
