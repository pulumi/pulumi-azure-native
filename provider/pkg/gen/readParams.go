// Copyright 2024, Pulumi Corporation.  All rights reserved.

package gen

// Map of "azure-native:foo:Bar" -> map of URL params and values, to be used in GET requests.
// When the spec defines optional query parameters for GET requests, the provider doesn't use them
// by default. Add them here to be used by default.
var defaultReadQueryParams = map[string]map[string]any{
	"azure-native:sqlvirtualmachine:SqlVirtualMachine": {
		"$expand": "*", // #3581: without this parameter, GET requests only return a subset of properties
	},
}
