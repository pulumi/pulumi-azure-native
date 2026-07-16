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
			"nodeResourceGroup",
			"windowsProfile",
			// networkProfile (see https://github.com/pulumi/pulumi-azure-native/issues/4756):
			// only the sub-properties below force a replacement. The rest of networkProfile
			// (e.g. advancedNetworking, loadBalancerProfile, natGatewayProfile, loadBalancerSku)
			// can be updated in place per `az aks update`.
			// Create-only, no update path exists at all in the AKS API:
			"serviceCidr",
			"serviceCidrs",
			"podCidrs",
			"dnsServiceIP",
			"networkMode",
			// `az aks update` accepts these, but only along a specific one-way migration
			// (e.g. CNI -> CNI Overlay, azure -> cilium dataplane, single-stack -> dual-stack),
			// not arbitrary edits. Force-new to stay conservative and avoid
			// https://github.com/pulumi/pulumi-azure-native/issues/959.
			"networkPlugin",
			"networkPluginMode",
			"podCidr",
			"networkDataplane",
			"networkPolicy",
			"outboundType",
			"ipFamilies",
			// AgentPoolNetworkProfile is a type shared with the standalone AgentPool resource (see
			// below); nodePublicIPTags is create-only there too (no `az aks nodepool update`
			// equivalent, unlike allowedHostPorts/applicationSecurityGroups). Listed here as well
			// since the type is generated once and whichever resource's pass generates it first
			// determines its ForceNew flags for both.
			"nodePublicIPTags",
		),
		"AgentPool": codegen.NewStringSet(
			"gpuInstanceProfile",
			"vmSize",
			// networkProfile.nodePublicIPTags is create-only: no `az aks nodepool update` equivalent
			// exists (unlike allowedHostPorts/applicationSecurityGroups, which are freely updatable).
			"nodePublicIPTags",
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

// noForceNewMap overrides forceNewMap and x-ms-mutability to *not* force
// replacement. Appropriate when Azure previously forced replacement but no longer does.
var noForceNewMap = map[openapi.ModuleName]map[string]codegen.StringSet{
	"ServiceBus": {
		"Namespace": codegen.NewStringSet("zoneRedundant"), // https://github.com/pulumi/pulumi-azure-native/issues/4105
	},
}
