// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/version"
)

var providerName = "azure-nextgen"

func main() {
	provider.Serve(providerName, version.Version, pulumiSchema, azureApiResources)
}
