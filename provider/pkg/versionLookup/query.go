package versionLookup

import (
	_ "embed"
	"encoding/json"
)

//go:embed v2-lock.json
var rawVersionLock []byte

// versionLock is a map from Azure provider name to resource name to API version.
// This is actually the openapi.DefaultVersionLock type but we don't want to import openapi here
// to avoid cycles.
var versionLock map[string]map[string]string

func init() {
	err := json.Unmarshal(rawVersionLock, &versionLock)
	if err != nil {
		panic(err)
	}
}

func GetDefaultApiVersionForResource(providerName, resourceName string) (string, bool) {
	if resources, ok := versionLock[providerName]; ok {
		if apiVersion, ok := resources[resourceName]; ok {
			return apiVersion, true
		}
	}
	return "", false
}
