package ioriver

import (
	"fmt"
	"testing"
)

const (
	testBehaviorId         = "8462bbdf-6a19-46f7-b4c1-6d2fe23727e4"
	testBehaviorName       = "test_behavior_name"
	testBehaviorActionId   = "8462bbdf-6a19-46f7-b4c1-6d2fe23727e5"
	testBehaviorActionType = "REDIRECT"
	testRedirectURL        = "https://www.mytest.com"
)

const serverBehaviorData = `{
	"id":"%s",
	"service":"%s",
	"name":"%s",
	"path_pattern":"%s",
	"behavior_actions": [
		{
			"id":"%s",
			"behavior":"%s",
			"type":"%s",
			"max_ttl":null,
			"response_header_name":null,
			"response_header_value":null,
			"cache_behavior_value":null,
			"redirect_url":"%s",
			"origin_cache_control_enabled":null,
			"pattern":null,
			"cookie":null,
			"auto_minify":null,
			"host_header":null
		}
	]
}
`

var expectedBehavior = Behavior{
	Id:          testBehaviorId,
	Service:     testServiceId,
	Name:        testBehaviorName,
	PathPattern: testPathPattern,
	Actions: []BehaviorAction{
		{
			Id:          testBehaviorActionId,
			Type:        testBehaviorActionType,
			RedirectURL: testRedirectURL,
		},
	},
}

func TestListBehaviors(t *testing.T) {
	path := fmt.Sprintf("/services/%s/behaviors/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverBehaviorData, testBehaviorId, testServiceId,
		testBehaviorName, testPathPattern, testBehaviorActionId, testBehaviorId, testBehaviorActionType,
		testRedirectURL))
	expected := []Behavior{expectedBehavior}
	RunServiceList[Behavior](t, (*IORiverClient).ListBehaviors, path, testServiceId, serverData, expected)
}

func TestGetBehavior(t *testing.T) {
	path := fmt.Sprintf("/services/%s/behaviors/%s/", testServiceId, testBehaviorId)
	serverData := fmt.Sprintf(serverBehaviorData, testBehaviorId, testServiceId, testBehaviorName, testPathPattern,
		testBehaviorActionId, testBehaviorId, testBehaviorActionType, testRedirectURL)
	RunServiceGet[Behavior](t, (*IORiverClient).GetBehavior, path,
		testServiceId, testBehaviorId, serverData, &expectedBehavior)
}

func TestCreateBehavior(t *testing.T) {
	newBehavior := Behavior{
		Name:        testBehaviorName,
		Service:     testServiceId,
		PathPattern: testPathPattern,
		Actions: []BehaviorAction{
			{
				Type:        testBehaviorActionType,
				RedirectURL: testRedirectURL,
			},
		},
	}

	path := fmt.Sprintf("/services/%s/behaviors/", testServiceId)
	serverData := fmt.Sprintf(serverBehaviorData, testBehaviorId, testServiceId, testBehaviorName, testPathPattern,
		testBehaviorActionId, testBehaviorId, testBehaviorActionType, testRedirectURL)
	RunCreate[Behavior](t, (*IORiverClient).CreateBehavior, path,
		newBehavior, serverData, &expectedBehavior)
}
