// Copyright 2022, Pulumi Corporation.  All rights reserved.

package gen

import (
	"testing"

	"github.com/segmentio/encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestExtractResponseNameId(t *testing.T) {
	// from connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2020-10-01-preview/examples/CreateResourcePool.json
	var createExample = []byte(`{
		"parameters": {},
		"responses": {
		  "200": {
			"body": {
			  "id": "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool",
			  "name": "HRPool"
			}
		  },
		  "201": {
			"body": {
			  "id": "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/ContosoAgent",
			  "name": "ContosoAgent"
			}
		  }
		}
	  }`)

	exampleJSON := parse(createExample, t)

	responseId, responseName := extractExampleResponseNameId(exampleJSON)
	assert.Equal(t, "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool", responseId)
	assert.Equal(t, "HRPool", responseName)
}

func TestExtractResponseNameIdInOrder(t *testing.T) {
	// lowest status code first regardless of order in the example
	var createExample = []byte(`{
		"parameters": {},
		"responses": {
		  "201": {
			"body": {
			  "id": "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/ContosoAgent",
			  "name": "ContosoAgent"
			}
		  },
		  "200": {
			"body": {
			  "id": "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool",
			  "name": "HRPool"
			}
		  }
		}
	  }`)

	exampleJSON := parse(createExample, t)

	responseId, responseName := extractExampleResponseNameId(exampleJSON)
	assert.Equal(t, "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool", responseId)
	assert.Equal(t, "HRPool", responseName)
}

func TestCanHaveOnlyId(t *testing.T) {
	var createExample = []byte(`{
		"parameters": {},
		"responses": {
		  "200": {
			"body": {
			  "id": "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool"
			}
		  },
		  "201": {
			"body": {
			  "name": "ContosoAgent"
			}
		  }
		}
	  }`)

	exampleJSON := parse(createExample, t)

	responseId, responseName := extractExampleResponseNameId(exampleJSON)
	assert.Equal(t, "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool", responseId)
	assert.Empty(t, responseName)
}

func TestPicksNameIdPairIfPresent(t *testing.T) {
	// lowest status code first regardless of order in the example
	var createExample = []byte(`{
		"parameters": {},
		"responses": {
		  "200": {
		    "body": {
			  "id": "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/ContosoAgent"
			}
		  },
		  "201": {
		    "body": {
			  "name": "ContosoAgent"
			}
		  },
		  "202": {
			"body": {
			  "id": "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool",
			  "name": "HRPool"
			}
		  }
		}
	  }`)

	exampleJSON := parse(createExample, t)

	responseId, responseName := extractExampleResponseNameId(exampleJSON)
	assert.Equal(t, "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool", responseId)
	assert.Equal(t, "HRPool", responseName)
}

func parse(jsonStr []byte, t *testing.T) map[string]interface{} {
	var exampleJSON map[string]interface{}
	if err := json.Unmarshal(jsonStr, &exampleJSON); err != nil {
		t.Fatalf("Could not parse example JSON: %v", err)
	}
	return exampleJSON
}
