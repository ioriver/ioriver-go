package ioriver

import (
	"fmt"
	"strings"
	"testing"
)

const (
	testCertId         = "e6dbff1b-7287-41a2-b988-8c9f48c13780"
	testCertName       = "test-cert"
	testCertType       = "MANAGED"
	testCertChain      = "dummy_cert_chain_data"
	testNotValidAfter  = "2024-01-02T23:59:59Z"
	testCertChallenges = `[{'domain': 'www.test.com', 'name': '_acme-challenge.www.test.com.', 'value': '441c92d7-141d-4753-a469-0a706f85cb3d.www.test.com.ioriver-acme.com'}]`
	testCertStatus     = "PENDING"
)

const serverCertData = `{
	"id":"%s",
	"name":"%s",
	"type": "%s",
	"cn":"%s",
	"not_valid_after": "%s",
	"certificate_chain":"%s",
	"challenges":"%s",
	"status":"%s"
}
`

var cnStr string = fmt.Sprintf(`[\"%s.com\", \"www.%s.com\"]`, testDomainName, testDomainName)
var rem string = `\"`
var rep *strings.Replacer = strings.NewReplacer(rem, "\"")

var expectedCertificate = Certificate{
	Id:               testCertId,
	Name:             testCertName,
	Type:             testCertType,
	Cn:               rep.Replace(cnStr),
	NotValidAfter:    testNotValidAfter,
	CertificateChain: testCertChain,
	Challenges:       testCertChallenges,
	Status:           testCertStatus,
}

func TestListCertificates(t *testing.T) {
	path := "/certificates/"
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverCertData, testCertId, testCertName, testCertType,
		cnStr, testNotValidAfter, testCertChain, testCertChallenges, testCertStatus))
	expected := []Certificate{expectedCertificate}
	RunList[Certificate](t, (*IORiverClient).ListCertificates, path, serverData, expected)
}

func TestGetCertificate(t *testing.T) {
	path := fmt.Sprintf("/certificates/%s/", testCertId)
	serverData := fmt.Sprintf(serverCertData, testCertId, testCertName, testCertType, cnStr,
		testNotValidAfter, testCertChain, testCertChallenges, testCertStatus)
	RunGet[Certificate](t, (*IORiverClient).GetCertificate, path, testCertId, serverData, &expectedCertificate)
}

func TestCreateCertificate(t *testing.T) {
	newCert := Certificate{
		Id:               testCertId,
		Name:             testCertName,
		Type:             testCertType,
		Cn:               rep.Replace(cnStr),
		NotValidAfter:    testNotValidAfter,
		Certificate:      "test_cert_data",
		PrivateKey:       "test_private_key_data",
		CertificateChain: testCertChain,
		Challenges:       testCertChallenges,
		Status:           testCertStatus,
	}

	path := "/certificates/"
	serverData := fmt.Sprintf(serverCertData, testCertId, testCertName, testCertType, cnStr,
		testNotValidAfter, testCertChain, testCertChallenges, testCertStatus)
	RunCreate[Certificate](t, (*IORiverClient).CreateCertificate, path, newCert, serverData, &expectedCertificate)
}
