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
	"Databricks": {
		// computeMode has no x-ms-mutability extension in the spec, but its description is explicit:
		// "Required on create, cannot be changed." Force a replacement rather than attempting an
		// in-place update the ARM API doesn't support. See
		// https://github.com/pulumi/pulumi-azure-native/issues/4766.
		"Workspace": codegen.NewStringSet("computeMode"),
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

// caseInsensitiveDiffMap is a map of Module Name -> Resource Name -> input properties whose string
// values should be diffed case-insensitively. Only add an entry here after confirming that Azure
// treats the property's values as case-insensitive but echoes back a different casing than what was
// sent (verified against a live resource, not just the spec). See
// https://github.com/pulumi/pulumi-azure-native/issues/4772: ManagedCluster's networkProfile.
// loadBalancerProfile.backendPoolType resolves to "NodeIPConfiguration" in the spec/SDK enum, but
// the AKS RP always returns "nodeIPConfiguration", producing a spurious diff.
var caseInsensitiveDiffMap = map[openapi.ModuleName]map[string]codegen.StringSet{
	"ContainerService": {
		"ManagedCluster": codegen.NewStringSet("backendPoolType"),
	},
}

// noDefaultMap is a map of Module Name -> Resource Name -> properties whose spec-declared `default`
// should be dropped from the generated schema, so the SDKs leave the property out of the request
// body unless the user assigns it explicitly. Only add an entry here after confirming that Azure
// rejects (or misbehaves on) receiving the spec's own default value, i.e. that the spec documents a
// default the RP won't actually accept on the wire for every flavour of the resource.
var noDefaultMap = map[openapi.ModuleName]map[string]codegen.StringSet{
	"Web": {
		// siteConfig.http20ProxyFlag was added to Microsoft.Web/sites in API version 2024-11-01
		// with a spec default of 0, which the SDKs then send on every request. Azure Functions on
		// Azure Container Apps rejects the property outright ("Http20ProxyFlag is not supported for
		// Azure Functions on Azure Container apps"), so apps with kind
		// `functionapp,linux,container,azurecontainerapps` started failing to create as soon as the
		// provider's default Web API version moved past 2024-11-01. Users can't work around it
		// because an unset (null) input is indistinguishable from an omitted one. See
		// https://github.com/pulumi/pulumi-azure-native/issues/4782.
		// Both resources that embed SiteConfig are listed because the type is generated once and
		// whichever resource's pass reaches it first fixes its shape for both.
		//
		// reserved is a real, honest create-time input (x-ms-mutability: [create, read], not
		// readOnly) that marks a Linux app: Azure genuinely expects the caller to set it, in
		// tandem with `kind`, when creating a Linux app. But its *correct* value depends on
		// `kind` in a way the spec's flat `default: false` can't express, so baking that default
		// into the generated SDK (e.g. `Reserved = false` in WebAppArgs()'s constructor) means
		// every Linux app created without explicitly setting `reserved: true` silently sends the
		// wrong value, and any later diff against an already-Linux resource's recorded state
		// forces an unwanted replace. See https://github.com/pulumi/pulumi-azure-native/issues/4447.
		"WebApp":     codegen.NewStringSet("http20ProxyFlag", "reserved"),
		"WebAppSlot": codegen.NewStringSet("http20ProxyFlag", "reserved"),
	},
}

// notRequiredInputsMap is a map of Module Name -> Resource Name -> input properties that the
// OpenAPI spec marks as required, but that we keep optional in the generated SDK. Only add an
// entry here after directly verifying (e.g. via a raw ARM REST call or ARM template deployment)
// that the live Azure API accepts the property being omitted and backfills a sensible default
// server-side
var notRequiredInputsMap = map[openapi.ModuleName]map[string]codegen.StringSet{
	"Databricks": {
		// computeMode became a required input in API version 2026-01-01. Confirmed live against
		// Azure (raw ARM REST PUT, an ARM template deployment, and reading a workspace created via
		// an older API version that predates the property) that omitting it is accepted and Azure
		// backfills "Hybrid" server-side. See https://github.com/pulumi/pulumi-azure-native/issues/4766.
		"Workspace": codegen.NewStringSet("computeMode"),
	},
}

// EnumOverride describes a synthetic enum type to substitute for a property's plain string (or
// array-of-string) type. Values should match the last API version that had a real enum for the
// property, so existing SDK consumers referencing enum members by name keep compiling.
type EnumOverride struct {
	// TypeName is the enum's name, e.g. "ComplianceStandard" becomes the token module:TypeName.
	TypeName string
	// Description is the generated enum type's description.
	Description string
	// Values are the allowed enum values (used verbatim as both name and value, matching how this
	// provider already models most Azure spec-derived enums that lack per-value x-ms-enum names).
	Values []string
}

// enumOverrideMap is a map of Module Name -> Resource Name -> input property name -> a synthetic
// enum type to substitute for that property's type. Only add an entry here when Azure's spec used
// to declare a real enum for the property and later downgraded it to a plain string (or dropped
// the enum type entirely) -- this keeps re-specifying a previously-valid enum member from becoming
// a breaking SDK change. It only takes effect where the spec doesn't already declare a real enum
// for the property (see genTypeSpec in types.go), so it's a no-op for any API version that still
// has one.
var enumOverrideMap = map[openapi.ModuleName]map[string]map[string]EnumOverride{
	"Databricks": {
		"Workspace": {
			// complianceStandards had a real enum (ComplianceStandard) through API version
			// 2025-08-01-preview. Starting 2025-10-01-preview the spec dropped the enum's value
			// list and x-ms-enum extension (kept as a bare string), and by 2026-01-01 the named
			// definition was removed entirely -- see https://github.com/pulumi/pulumi-azure-native/issues/4766.
			// Values below are the full set from 2025-08-01-preview, the richest version that had them.
			"complianceStandards": {
				TypeName:    "ComplianceStandard",
				Description: "Compliance standard that can be associated with a workspace.",
				Values: []string{
					"NONE",
					"HIPAA",
					"PCI_DSS",
					"CYBER_ESSENTIAL_PLUS",
					"FEDRAMP_HIGH",
					"CANADA_PROTECTED_B",
					"IRAP_PROTECTED",
					"ISMAP",
					"HITRUST",
					"K_FSI",
				},
			},
		},
	},
}
