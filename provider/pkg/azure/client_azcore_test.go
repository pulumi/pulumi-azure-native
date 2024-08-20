// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeLocationHeader(t *testing.T) {
	const host = "https://management.azure.com"
	const apiVersion = "2022-09-01"

	for testName, loc := range map[string]string{
		"RelativeUrl":               "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm",
		"RelativeUrlWithApiVersion": "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm?api-version=" + apiVersion,
		"FullUrl":                   "https://management.azure.com/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm",
		"FullUrlWithApiVersion":     "https://management.azure.com/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm?api-version=" + apiVersion,
	} {
		t.Run(testName, func(t *testing.T) {
			headers := http.Header{}
			headers.Add("Location", loc)
			normalizeLocationHeader(host, apiVersion, headers)
			assert.Equal(t,
				host+"/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm?api-version="+apiVersion,
				headers.Get("Location"))
		})
	}

	for testName, loc := range map[string]string{
		"MalformedUrl": "foo[:",
		"EmptyUrl":     "",
	} {
		t.Run(testName, func(t *testing.T) {
			headers := http.Header{}
			headers.Add("Location", loc)
			normalizeLocationHeader(host, apiVersion, headers)
			assert.Equal(t, loc, headers.Get("Location"))
		})
	}

	t.Run("NoLocationHeader", func(t *testing.T) {
		headers := http.Header{}
		normalizeLocationHeader(host, apiVersion, headers)
		assert.Empty(t, headers.Get("Location"))
	})
}
