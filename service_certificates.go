package ioriver

import (
	"fmt"
)

type ServiceCertificate struct {
	Id          string `json:"id"`
	Service     string `json:"service"`
	Certificate string `json:"certificate"`
}

const serviceCertBasePath = "services/%s/certificates/"

func ListServiceCertificates(client *IORiverClient, serviceId string) ([]ServiceCertificate, error) {
	path := fmt.Sprintf(serviceCertBasePath, serviceId)
	return List[ServiceCertificate](client, path)
}

func AddServiceCertificate(client *IORiverClient, serviceId, certificateId string) (*ServiceCertificate, error) {
	path := fmt.Sprintf(serviceCertBasePath, serviceId)
	payload := ServiceCertificate{
		Service:     serviceId,
		Certificate: certificateId,
	}
	return Create[ServiceCertificate](client, path, payload)
}

func RemoveServiceCertificate(client *IORiverClient, serviceId, serviceCertId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(serviceCertBasePath, serviceId), serviceCertId)
	return Delete(client, path)
}

type replaceCertPayload struct {
	NewCertificateId string `json:"new_certificate_id"`
}

// The input ID parameters refer to different resource types.
func ReplaceServiceCertificate(client *IORiverClient, serviceId, newCertificateId, oldServiceCertId string) (*ServiceCertificate, error) {
	path := fmt.Sprintf("%s%s/replace/", fmt.Sprintf(serviceCertBasePath, serviceId), oldServiceCertId)
	payload := replaceCertPayload{
		NewCertificateId: newCertificateId,
	}
	return Update[ServiceCertificate](client, path, payload)
}
