package ioriver

import (
	"fmt"
	"testing"
)

const (
	testLogDestinationBucket = "test-bucket"
	testLogDestinationPath   = "/test/path"
)

const serverLogDestinationData = `{
	"id":"%s",
	"service":"%s",
	"name":"test",
	"type":"S3",
	"s3_bucket": "%s",
	"s3_path": "%s"
}`

var expectedLogDestination = LogDestination{
	Id:       testObjectId,
	Service:  testServiceId,
	Name:     "test",
	Type:     "S3",
	S3Bucket: testLogDestinationBucket,
	S3Path:   testLogDestinationPath,
}

func TestListLogDestinations(t *testing.T) {
	path := fmt.Sprintf("/services/%s/log-destinations/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverLogDestinationData, testObjectId, testServiceId, testLogDestinationBucket, testLogDestinationPath))
	expected := []LogDestination{expectedLogDestination}
	RunServiceList[LogDestination](t, (*IORiverClient).ListLogDestinations, path, testServiceId, serverData, expected)
}

func TestGetLogDestination(t *testing.T) {
	path := fmt.Sprintf("/services/%s/log-destinations/%s/", testServiceId, testObjectId)
	serverData := fmt.Sprintf(serverLogDestinationData, testObjectId, testServiceId, testLogDestinationBucket, testLogDestinationPath)
	RunServiceGet[LogDestination](t, (*IORiverClient).GetLogDestination, path, testServiceId, testObjectId, serverData, &expectedLogDestination)
}

func TestCreateLogDestination(t *testing.T) {
	newLogDestination := LogDestination{
		Name:     "test",
		Service:  testServiceId,
		Type:     "S3",
		S3Bucket: testLogDestinationBucket,
		S3Path:   testLogDestinationPath,
	}

	path := fmt.Sprintf("/services/%s/log-destinations/", testServiceId)
	serverData := fmt.Sprintf(serverLogDestinationData, testObjectId, testServiceId, testLogDestinationBucket, testLogDestinationPath)
	RunCreate[LogDestination](t, (*IORiverClient).CreateLogDestination, path, newLogDestination, serverData, &expectedLogDestination)
}
