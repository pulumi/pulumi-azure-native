// Copyright 2021, Pulumi Corporation.  All rights reserved.

package gen

import "github.com/pulumi/pulumi/pkg/v3/codegen"

// forceNewMap is a map of Azure Resource Provider -> Resource Name -> Properties that cause replacements.
// API Versions are currently ignored.
var forceNewMap = map[string]map[string]codegen.StringSet{
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
	"Insights": {
		"Component": codegen.NewStringSet(), // covered by x-ms-mutability
	},
	"Network": {
		"PublicIPAddress": codegen.NewStringSet("location", "publicIPAddressVersion", "publicIPPrefix", "sku"),
		"Subnet":          codegen.NewStringSet(), // no force-news
		"VirtualNetwork":  codegen.NewStringSet("location"),
	},
	"Resources": {
		"ResourceGroup": codegen.NewStringSet("location"),
	},
	"ServiceBus": {
		"Topic": codegen.NewStringSet("requiresDuplicateDetection", "requiresSession", "enablePartitioning"),
		"Queue": codegen.NewStringSet("requiresDuplicateDetection", "requiresSession", "enablePartitioning"),
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
