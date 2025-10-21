// Copyright 2021, Pulumi Corporation.  All rights reserved.

package gen

import (
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
)

// forceNewMap is a map of Module Name -> Resource Name -> Properties that cause replacements.
// API Versions are currently ignored.
var forceNewMap = map[openapi.ModuleName]map[string]codegen.StringSet{
	"Authorization": {
		"RoleAssignment": codegen.NewStringSet("principalId", "scope"),
	},
	"Automation": {
		"JobSchedule": codegen.NewStringSet("parameters"),
	},
	"Cdn": {
		"Profile": codegen.NewStringSet(
			"location",
			// sku
			"name",
		),
	},
	"ContainerService": {
		"ManagedCluster": codegen.NewStringSet(
			// cluster
			"diskEncryptionSetID",
			"dnsPrefix",
			"fqdnSubdomain",
			"linuxProfile",
			"location",
			"networkProfile",
			"nodeResourceGroup",
			"windowsProfile",
		),
		"AgentPool": codegen.NewStringSet(
			"gpuInstanceProfile",
			"vmSize",
		),
	},
	"DocumentDB": {
		"SqlResourceSqlRoleAssignment": codegen.NewStringSet("principalId", "scope"),
	},
	"Insights": {
		"Component": codegen.NewStringSet(), // covered by x-ms-mutability
	},
	"Network": {
		// https://github.com/pulumi/pulumi-azure-native/issues/3883
		// https://stackoverflow.com/questions/78877433/arm-template-is-it-possible-to-change-subnets-of-private-endpoint
		// https://learn.microsoft.com/en-us/answers/questions/1295954/how-to-migrate-a-keyvault-private-endpoint-to-a-ne
		"PrivateEndpoint": codegen.NewStringSet("subnet"),
		"PublicIPAddress": codegen.NewStringSet("location", "publicIPAddressVersion", "publicIPPrefix", "sku"),
		"Subnet":          codegen.NewStringSet(), // no force-news
		"VirtualNetwork":  codegen.NewStringSet("location"),
	},
	"Resources": {
		"ResourceGroup": codegen.NewStringSet("location"),
	},
	"ServiceBus": {
		"Topic":     codegen.NewStringSet("requiresDuplicateDetection", "requiresSession", "enablePartitioning"),
		"Queue":     codegen.NewStringSet("requiresDuplicateDetection", "requiresSession", "enablePartitioning"),
		"Namespace": codegen.NewStringSet(), // no force-news, allows zoneRedundant to be updated in-place
	},
	"Storage": {
		"BlobContainer":  codegen.NewStringSet(), // no force-news
		"StorageAccount": codegen.NewStringSet("isHnsEnabled", "location"),
	},
	"Web": {
		"AppServicePlan": codegen.NewStringSet(), // covered by x-ms-mutability
		"WebApp":         codegen.NewStringSet("location", "kind"),
	},
}
