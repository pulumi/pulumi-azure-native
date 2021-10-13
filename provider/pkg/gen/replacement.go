// Copyright 2021, Pulumi Corporation.  All rights reserved.

package gen

import "github.com/pulumi/pulumi/pkg/v3/codegen"

// forceNewMap is a map of Azure Resource Provider -> Resource Name -> Properties that cause replacements.
// API Versions are currently ignored.
var forceNewMap = map[string]map[string]codegen.StringSet{
	"Authorization": {
		"RoleAssignment": codegen.NewStringSet("principalId", "scope"),
	},
	"ContainerService": {
		"ManagedCluster": codegen.NewStringSet(
			// aadProfile
			"enableAzureRBAC",
			// agentPoolProfiles
			"availabilityZones",
			"enableEncryptionAtHost",
			"enableFIPS",
			"enableNodePublicIP",
			"kubeletConfig",
			"linuxOSConfig",
			"maxPods",
			"name",
			"nodeLabels",
			"nodePublicIPPrefixID",
			"nodeTaints",
			"osDiskSizeGB",
			"osSKU",
			"podSubnetID",
			"proximityPlacementGroupID",
			"type",
			"vnetSubnetID",
			"vmSize",
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
	"Storage": {
		"BlobContainer":  codegen.NewStringSet(), // no force-news
		"StorageAccount": codegen.NewStringSet("isHnsEnabled", "location"),
	},
	"Web": {
		"AppServicePlan": codegen.NewStringSet(), // covered by x-ms-mutability
		"WebApp":         codegen.NewStringSet("location"),
	},
}
