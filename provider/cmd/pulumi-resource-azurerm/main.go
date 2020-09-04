// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	"github.com/pulumi/pulumi-azurerm/provider/pkg/provider"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/version"
)

var providerName = "azurerm"

func main() {
	provider.Serve(providerName, version.Version, pulumiSchema, azureApiResources)
}
