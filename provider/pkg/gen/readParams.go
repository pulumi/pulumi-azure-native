// Copyright 2024, Pulumi Corporation.  All rights reserved.

package gen

// Map of "azure-native:foo:Bar" -> map of URL params and values, to be used in GET requests
var additionalReadParams = map[string]map[string]any{
	"azure-native:sqlvirtualmachine:SqlVirtualMachine": {
		"$expand": "*", // #3581: without this parameter, GET requests only return a subset of properties
	},
}
