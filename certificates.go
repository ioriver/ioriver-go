package ioriver

import (
	"fmt"
)

type CertificateType string

const (
	MANAGED      CertificateType = "MANAGED"
	SELF_MANAGED CertificateType = "SELF_MANAGED"
	EXTERNAL     CertificateType = "EXTERNAL"
)

type CertificateStatus string

const (
	VALID   CertificateStatus = "VALID"
	FAILED  CertificateStatus = "FAILED"
	EXPIRED CertificateStatus = "EXPIRED"
	PENDING CertificateStatus = "PENDING"
)

type Certificate struct {
	Id               string            `json:"id,omitempty"`
	Name             string            `json:"name"`
	Type             CertificateType   `json:"type"`
	Cn               string            `json:"cn,omitempty"`
	NotValidAfter    string            `json:"not_valid_after,omitempty"`
	Certificate      string            `json:"certificate,omitempty"`
	PrivateKey       string            `json:"private_key,omitempty"`
	CertificateChain string            `json:"certificate_chain,omitempty"`
	Challenges       string            `json:"challenges,omitempty"`
	Status           CertificateStatus `json:"status,omitempty"`
}

const certBasePath = "certificates/"

func (client *IORiverClient) GetCertificate(id string) (*Certificate, error) {
	path := fmt.Sprintf("%s%s/", certBasePath, id)
	return Get[Certificate](client, path)
}

func (client *IORiverClient) ListCertificates() ([]Certificate, error) {
	return List[Certificate](client, certBasePath)
}

func (client *IORiverClient) CreateCertificate(cert Certificate) (*Certificate, error) {
	return Create[Certificate](client, certBasePath, cert)
}

func (client *IORiverClient) UpdateCertificate(cert Certificate) (*Certificate, error) {
	path := certBasePath + cert.Id + "/"
	return Update[Certificate](client, path, cert)
}

func (client *IORiverClient) DeleteCertificate(certId string) error {
	path := fmt.Sprintf("%s%s/", certBasePath, certId)
	return Delete(client, path)
}
