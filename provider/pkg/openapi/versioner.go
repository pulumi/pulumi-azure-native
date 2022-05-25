// Copyright 2021, Pulumi Corporation.  All rights reserved.

package openapi

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/segmentio/encoding/json"
)

// versioner checks whether the given combination of a provider, a resource path, and a versions
// are known in version metadata returned by Azure (retrieved via the Azure CLI, stored in a local JSON file).
type versioner struct {
	lookup map[string]map[string]codegen.StringSet
}

func newVersioner() (*versioner, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	jsonFile, err := os.Open(filepath.Join(dir, "/azure-provider-versions/provider_list.json"))
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var provs []prov
	err = json.Unmarshal(byteValue, &provs)
	if err != nil {
		return nil, err
	}

	result := map[string]map[string]codegen.StringSet{}

	for _, prov := range provs {
		namespace := strings.ToLower(prov.Namespace)
		if !strings.HasPrefix(namespace, "microsoft.") {
			continue
		}
		providerName := strings.TrimPrefix(namespace, "microsoft.")

		versions := map[string]codegen.StringSet{}
		allVersions := codegen.NewStringSet()
		for _, rt := range prov.ResourceTypes {
			set := codegen.NewStringSet()
			for _, v := range rt.ApiVersions {
				name := "v" + strings.ReplaceAll(v, "-", "")
				allVersions.Add(name)
				set.Add(name)
			}
			versions[strings.ToLower(rt.ResourceType)] = set
		}
		versions[""] = allVersions
		result[providerName] = versions
	}

	return &versioner{lookup: result}, nil
}

// A manually-maintained list of stable versions that we want to promote a later preview version to be used for
// the top-level resource. These versions are "known" to be outdated but no newer stable versions were released yet.
// However, there's no formal way to derive this from Open API specs, so we have to maintain this manual map.
var deprecatedProviderVersions = map[string][]string{
	"Authorization": {"v20150701"},
	"Sql":           {"v20140401", "v20150501"},
	// Task API definition is broken in these stable versions but fixed in later previews,
	// see https://github.com/pulumi/pulumi-azure-native/issues/736.
	"ContainerRegistry": {"v20180901", "v20190401"},
}

// A manually-maintained list of API versions to ignore while calculating the top-level resources.
var ignoredProviderVersions = map[string][]string{
	"Automanage": {"v20210430preview"}, // Conflict in configuration profile definition (enum vs resource) with the previous version.
	"Migrate":    {"v20200501"},        // Conflict in property types with the previous version.
	"NetApp":     {"v20210801"},        // Conflict in property counts with the previous version.
	"StorSimple": {"v20161001"},
}

// A manually-maintained list of versions that we want to use for the top-level resource. The primary goal is to
// avoid breaking changes within a single major version of the provider that could come with new API versions.
// We reset this map every time we release a new major version.
// Currently populated for 1.* series.
var cutoffProviderVersions = map[string]string{
	"aad":                           "v20210301",
	"alertsmanagement":              "v20200804preview",
	"apimanagement":                 "v20210401preview",
	"appconfiguration":              "v20200701preview",
	"appplatform":                   "v20210901preview",
	"authorization":                 "v20210301preview",
	"automanage":                    "v20200630preview",
	"autonomousdevelopmentplatform": "v20210201preview",
	"avs":                           "v20210101preview",
	"azureactivedirectory":          "v20190101preview",
	"azurearcdata":                  "v20210601preview",
	"azurestackhci":                 "v20210101preview",
	"batch":                         "v20210101",
	"cache":                         "v20210301",
	"cdn":                           "v20200901",
	"certificateregistration":       "v20201001",
	"cognitiveservices":             "v20170418",
	"compute":                       "v20210301",
	"confidentialledger":            "v20201201preview",
	"confluent":                     "v20200301",
	"consumption":                   "v20191001",
	"containerinstance":             "v20210301",
	"containerregistry":             "v20201101preview",
	"containerservice":              "v20210301",
	"costmanagement":                "v20200601",
	"databox":                       "v20201101",
	"databoxedge":                   "v20201201",
	"datamigration":                 "v20180715preview",
	"dataprotection":                "v20210101",
	"datashare":                     "v20200901",
	"dbformysql":                    "v20171201",
	"dbforpostgresql":               "v20171201",
	"desktopvirtualization":         "v20210201preview",
	"devices":                       "v20200831",
	"documentdb":                    "v20210315",
	"domainregistration":            "v20201001",
	"eventgrid":                     "v20210601preview",
	"eventhub":                      "v20180101preview",
	"extendedlocation":              "v20210315preview",
	"fluidrelay":                    "v20210312preview",
	"hardwaresecuritymodules":       "v20181031preview",
	"hdinsight":                     "v20180601preview",
	"healthbot":                     "v20201208",
	"hybridcompute":                 "v20210325preview",
	"hybridnetwork":                 "v20200101preview",
	"insights":                      "v20201020",
	"iotsecurity":                   "v20210201preview",
	"keyvault":                      "v20210601preview",
	"kubernetes":                    "v20210401preview",
	"kubernetesconfiguration":       "v20210301",
	"kusto":                         "v20210101",
	"labservices":                   "v20211001preview",
	"machinelearningservices":       "v20210101",
	"maintenance":                   "v20210401preview",
	"management":                    "v20200501",
	"maps":                          "v20200201preview",
	"media":                         "v20200501",
	"migrate":                       "v20210101",
	"mobilenetwork":                 "v202201101preview",
	"netapp":                        "v20201201",
	"operationalinsights":           "v20201001",
	"peering":                       "v20210101",
	"policyinsights":                "v20190701",
	"purview":                       "v20201201preview",
	"quantum":                       "v20191104preview",
	"recoveryservices":              "v20210201preview",
	"security":                      "v20200101preview",
	"securityinsights":              "v20210301preview",
	"servicebus":                    "v20170401", // This might be able to be lifted https://github.com/pulumi/pulumi-azure-native/pull/1592#issuecomment-1076447550
	"servicefabric":                 "v20200301",
	"servicelinker":                 "v20211101preview",
	"signalrservice":                "v20210401preview",
	"solutions":                     "v20190701",
	"sql":                           "v20201101preview",
	"sqlvirtualmachine":             "v20170301preview",
	"streamanalytics":               "v20160301",
	"storage":                       "v20210201",
	"storagecache":                  "v20210301",
	"storagesync":                   "v20200301",
	"storagepool":                   "v20200315preview",
	"subscription":                  "v20200901",
	"synapse":                       "v20210301",
	"videoanalyzer":                 "v20210501preview",
	"videoindexer":                  "v20211018preview",
	"virtualmachineimages":          "v20200214",
	"web":                           "v20201201",
	"webpubsub":                     "v20210401preview",
}

var lockedTypeVersions = map[string]string{
	"azure-native:authorization:getRoleManagementPolicyAssignment": "v20201001preview",
	"azure-native:authorization:RoleManagementPolicyAssignment":    "v20201001preview",

	"cache:FirewallRule":  "v20200601",
	"cache:LinkedServer":  "v20200601",
	"cache:PatchSchedule": "v20200601",
	"cache:Redis":         "v20200601",
	"cache:listRedisKeys": "v20200601",

	"compute:AvailabilitySet":                           "v20201201",
	"compute:DedicatedHost":                             "v20201201",
	"compute:DedicatedHostGroup":                        "v20201201",
	"compute:Image":                                     "v20201201",
	"compute:ProximityPlacementGroup":                   "v20201201",
	"compute:SshPublicKey":                              "v20201201",
	"compute:getLogAnalyticExportRequestRateByInterval": "v20201201",
	"compute:getLogAnalyticExportThrottledRequests":     "v20201201",

	"containerservice:Snapshot":    "v20210801",
	"containerservice:getSnapshot": "v20210801",

	"eventgrid:PartnerTopicEventSubscription": "v20200401preview",
	"eventgrid:SystemTopicEventSubscription":  "v20200401preview",

	"insights:Component": "v20150501",

	"machinelearningservices:ACIService":             "v20210101",
	"machinelearningservices:AKSService":             "v20210101",
	"machinelearningservices:EndpointVariant":        "v20210101",
	"machinelearningservices:MachineLearningService": "v20210101",

	"maintenance:MaintenanceConfiguration":    "v20200401",
	"maintenance:getMaintenanceConfiguration": "v20200401",

	"network:AdminRule":                                   "v20210201preview",
	"network:AdminRuleCollection":                         "v20210201preview",
	"network:ApplicationGateway":                          "v20201101",
	"network:ApplicationGatewayPrivateEndpointConnection": "v20201101",
	"network:ApplicationSecurityGroup":                    "v20201101",
	"network:AzureFirewall":                               "v20201101",
	"network:BastionHost":                                 "v20201101",
	"network:ConnectionMonitor":                           "v20201101",
	"network:ConnectivityConfiguration":                   "v20210201preview",
	"network:CustomIPPrefix":                              "v20201101",
	"network:DdosCustomPolicy":                            "v20201101",
	"network:DdosProtectionPlan":                          "v20201101",
	"network:DscpConfiguration":                           "v20201101",
	"network:ExpressRouteCircuit":                         "v20201101",
	"network:ExpressRouteCircuitAuthorization":            "v20201101",
	"network:ExpressRouteCircuitConnection":               "v20201101",
	"network:ExpressRouteCircuitPeering":                  "v20201101",
	"network:ExpressRouteConnection":                      "v20201101",
	"network:ExpressRouteCrossConnectionPeering":          "v20201101",
	"network:ExpressRouteGateway":                         "v20201101",
	"network:ExpressRoutePort":                            "v20201101",
	"network:FirewallPolicy":                              "v20201101",
	"network:FirewallPolicyRuleCollectionGroup":           "v20201101",
	"network:FlowLog":                                     "v20201101",
	"network:HubRouteTable":                               "v20201101",
	"network:HubVirtualNetworkConnection":                 "v20201101",
	"network:InboundNatRule":                              "v20201101",
	"network:IpAllocation":                                "v20201101",
	"network:IpGroup":                                     "v20201101",
	"network:LoadBalancer":                                "v20201101",
	"network:LoadBalancerBackendAddressPool":              "v20201101",
	"network:LocalNetworkGateway":                         "v20201101",
	"network:ManagementGroupNetworkManagerConnection":     "v20210501preview",
	"network:NatGateway":                                  "v20201101",
	"network:NatRule":                                     "v20201101",
	"network:NetworkGroup":                                "v20210201preview",
	"network:NetworkInterface":                            "v20201101",
	"network:NetworkInterfaceTapConfiguration":            "v20201101",
	"network:NetworkManager":                              "v20210201preview",
	"network:NetworkProfile":                              "v20201101",
	"network:NetworkSecurityGroup":                        "v20201101",
	"network:NetworkVirtualAppliance":                     "v20201101",
	"network:NetworkWatcher":                              "v20201101",
	"network:P2sVpnGateway":                               "v20201101",
	"network:PacketCapture":                               "v20201101",
	"network:PrivateDnsZoneGroup":                         "v20201101",
	"network:PrivateEndpoint":                             "v20201101",
	"network:PrivateLinkService":                          "v20201101",
	"network:PrivateLinkServicePrivateEndpointConnection": "v20201101",
	"network:PublicIPAddress":                             "v20201101",
	"network:PublicIPPrefix":                              "v20201101",
	"network:Route":                                       "v20201101",
	"network:RouteFilter":                                 "v20201101",
	"network:RouteFilterRule":                             "v20201101",
	"network:RouteTable":                                  "v20201101",
	"network:SecurityAdminConfiguration":                  "v20210201preview",
	"network:SecurityPartnerProvider":                     "v20201101",
	"network:SecurityRule":                                "v20201101",
	"network:SecurityUserConfiguration":                   "v20210201preview",
	"network:ServiceEndpointPolicy":                       "v20201101",
	"network:ServiceEndpointPolicyDefinition":             "v20201101",
	"network:Subnet":                                      "v20201101",
	"network:UserRule":                                    "v20210201preview",
	"network:UserRuleCollection":                          "v20210201preview",
	"network:VirtualApplianceSite":                        "v20201101",
	"network:VirtualHub":                                  "v20201101",
	"network:VirtualHubBgpConnection":                     "v20201101",
	"network:VirtualHubIpConfiguration":                   "v20201101",
	"network:VirtualHubRouteTableV2":                      "v20201101",
	"network:VirtualNetwork":                              "v20201101",
	"network:VirtualNetworkGateway":                       "v20201101",
	"network:VirtualNetworkGatewayConnection":             "v20201101",
	"network:VirtualNetworkGatewayNatRule":                "v20210301",
	"network:VirtualNetworkPeering":                       "v20201101",
	"network:VirtualNetworkTap":                           "v20201101",
	"network:VirtualWan":                                  "v20201101",
	"network:VpnConnection":                               "v20201101",
	"network:VpnGateway":                                  "v20201101",
	"network:VpnServerConfiguration":                      "v20201101",
	"network:VpnSite":                                     "v20201101",
	"network:WebApplicationFirewallPolicy":                "v20201101",

	"network:getActiveSessions":                                     "v20201101",
	"network:getApplicationGatewayBackendHealthOnDemand":            "v20201101",
	"network:getBastionShareableLink":                               "v20201101",
	"network:getP2sVpnGatewayP2sVpnConnectionHealth":                "v20201101",
	"network:getP2sVpnGatewayP2sVpnConnectionHealthDetailed":        "v20201101",
	"network:getVirtualNetworkGatewayAdvertisedRoutes":              "v20201101",
	"network:getVirtualNetworkGatewayBgpPeerStatus":                 "v20201101",
	"network:getVirtualNetworkGatewayLearnedRoutes":                 "v20201101",
	"network:getVirtualNetworkGatewayVpnclientConnectionHealth":     "v20201101",
	"network:getVirtualNetworkGatewayVpnclientIpsecParameters":      "v20201101",
	"network:listActiveConnectivityConfiguration":                   "v20210201preview",
	"network:listActiveConnectivityConfigurations":                  "v20210201preview",
	"network:listActiveSecurityAdminRule":                           "v20210201preview",
	"network:listActiveSecurityAdminRules":                          "v20210201preview",
	"network:listActiveSecurityUserRule":                            "v20210201preview",
	"network:listActiveSecurityUserRules":                           "v20210201preview",
	"network:listEffectiveConnectivityConfiguration":                "v20210201preview",
	"network:listEffectiveVirtualNetworkByNetworkGroup":             "v20210201preview",
	"network:listEffectiveVirtualNetworkByNetworkManager":           "v20210201preview",
	"network:listListEffectiveVirtualNetworkByNetworkGroup":         "v20210201preview",
	"network:listNetworkManagerDeploymentStatus":                    "v20210201preview",
	"network:listNetworkManagerEffectiveConnectivityConfigurations": "v20210201preview",
	"network:listNetworkManagerEffectiveSecurityAdminRule":          "v20210201preview",
	"network:listNetworkManagerEffectiveSecurityAdminRules":         "v20210201preview",

	"resources:Deployment":                       "v20210101",
	"resources:DeploymentAtManagementGroupScope": "v20210101",
	"resources:DeploymentAtScope":                "v20210101",
	"resources:DeploymentAtSubscriptionScope":    "v20210101",
	"resources:DeploymentAtTenantScope":          "v20210101",

	"servicebus:PrivateEndpointConnection":    "v20180101preview",
	"servicebus:getPrivateEndpointConnection": "v20180101preview",

	"sql:TransparentDataEncryption":    "v20140401",
	"sql:getTransparentDataEncryption": "v20140401",

	"web:WebAppSwiftVirtualNetworkConnection": "v20201001",
}

// A manually-maintained list of resources that were deprecated and have no direct successor. They "pollute"
// the top-level resources, including bringing old incompatible types, so we'd rather exclude them. Note that
// they aren't officially deprecated in the APIs.
var deprecatedResources = codegen.NewStringSet(
	"apimanagement:TenantPolicy",
	"batchai:FileServer",
	"consumption:BudgetByResourceGroupName",
	"containerregistry:BuildStep",
	"containerservice:ContainerService",
	"costmanagement:Budget",
	"costmanagement:ReportConfig",
	"costmanagement:ReportConfigByResourceGroupName",
	"datamigration:ServiceTask",
	"synapse:SqlDatabase",
	"web:CertificateCsr",
	"web:SiteInstanceDeployment",
	"web:SiteInstanceDeploymentSlot",
)

// calculateLatestVersions builds a map of latest versions per API paths from a map of all versions of a resource
// provider. The result is a map from a resource type name to resource specs.
func (c *versioner) calculateLatestVersions(provider string, versionMap ProviderVersions,
	invokes bool) (latestResources map[string]*ResourceSpec) {
	deprecatedVersions := codegen.NewStringSet()
	if v, ok := deprecatedProviderVersions[provider]; ok {
		deprecatedVersions = codegen.NewStringSet(v...)
	}
	ignoredVersions := codegen.NewStringSet()
	if v, ok := ignoredProviderVersions[provider]; ok {
		ignoredVersions = codegen.NewStringSet(v...)
	}

	var stables, previews []string
	for version := range versionMap {
		switch {
		case strings.Contains(version, "private"):
			// skip
		case ignoredVersions.Has(version):
			// skip
		case deprecatedVersions.Has(version):
			// Treat deprecated versions as preview for sorting purpose.
			previews = append(previews, version)
		case !IsPreview(version):
			stables = append(stables, version)
		default:
			previews = append(previews, version)
		}
	}
	// Sort the versions from earliest to latest previews, then from earliest to latest stable.
	sort.Strings(previews)
	sort.Strings(stables)
	versions := append(previews, stables...)
	availableVersions := codegen.NewStringSet(versions...)

	pathTypeNames := map[string]string{}
	latestResources = map[string]*ResourceSpec{}
	knownResources := codegen.NewStringSet()
	for _, version := range versions {
		items := versionMap[version]
		res := items.Resources
		if invokes {
			res = items.Invokes
		}

		for typeName, r := range res {
			fullTypeName := fmt.Sprintf("%s:%s", strings.ToLower(provider), typeName)
			if deprecatedResources.Has(fullTypeName) {
				continue
			}
			if lockedVersion, ok := lockedTypeVersions[fullTypeName]; ok && lockedVersion != version {
				// If we have a locked version for this resource, ignore all other versions.
				continue
			}

			normalizedPath := normalizePath(r.Path)
			if cutoffVersion, ok := cutoffProviderVersions[strings.ToLower(provider)]; ok && version > cutoffVersion {
				if _, exists := pathTypeNames[normalizedPath]; exists {
					// Ignore all the versions that are newer than the cutoff version for this provider,
					// given that a pre-cutoff version exists.
					continue
				}
			}

			knownVersions := c.knownVersions(provider, r.Path)
			isKnown := knownVersions.Has(version)
			if !isKnown && knownResources.Has(normalizedPath) {
				availableKnown := intersectStringSet(knownVersions, availableVersions)
				if availableKnown.Any() {
					continue
				}
			}
			if isKnown {
				knownResources.Add(normalizedPath)
			}

			previousTypeName := pathTypeNames[normalizedPath]
			if previousTypeName != "" && previousTypeName != typeName {
				delete(latestResources, previousTypeName)
			}

			pathTypeNames[normalizedPath] = typeName
			copyResource := *r
			latestResources[typeName] = &copyResource
		}
	}
	return latestResources
}

func (c *versioner) knownVersions(provider, path string) codegen.StringSet {
	providerLookup, ok := c.lookup[strings.ToLower(provider)]
	if !ok {
		return codegen.NewStringSet()
	}

	resourceTypes := c.armResourceTypes(path)
	for _, resourceType := range resourceTypes {
		if resourceVersions, ok := providerLookup[resourceType]; ok {
			return resourceVersions
		}
	}
	return codegen.NewStringSet()
}

func intersectStringSet(ss, other codegen.StringSet) codegen.StringSet {
	result := codegen.NewStringSet()
	for v := range ss {
		if other.Has(v) {
			result.Add(v)
		}
	}
	return result
}

// armResourceTypes returns a list of ARM resource type identifiers starting from the most specific one and ending
// with an empty name. We use this helper function to lookup resource types starting from the most specific one
// and then generalizing. We don't know at which level Microsoft defined the versions for this resource, so we should
// check them all until we find one.
// Example of a result: { "servers/databases/transparentDataEncryption", "servers/databases", "servers", "" }.
func (c *versioner) armResourceTypes(path string) []string {
	parts := strings.Split(strings.ToLower(path), "microsoft.")
	resourcePath := parts[len(parts)-1]
	segments := strings.Split(resourcePath, "/")
	var result, builder []string
	for _, segment := range segments[1:] {
		if !strings.HasPrefix(segment, "{") {
			builder = append(builder, segment)
			result = append([]string{strings.Join(builder, "/")}, result...)
		}
	}
	result = append(result, "")
	return result
}

// calculatePathVersions builds a map of all versions defined for an API paths from a map of all versions of a resource
// provider. The result is a map from a normalized path to a set of versions for that path.
func (c *versioner) calculatePathVersions(versionMap ProviderVersions) (pathVersions map[string]codegen.StringSet) {
	pathVersions = map[string]codegen.StringSet{}
	for version, items := range versionMap {
		for _, r := range items.Resources {
			normalizedPath := normalizePath(r.Path)
			versions, ok := pathVersions[normalizedPath]
			if !ok {
				versions = codegen.StringSet{}
				pathVersions[normalizedPath] = versions
			}
			versions.Add(version)
		}
	}
	return pathVersions
}

type prov struct {
	Namespace     string    `json:"namespace"`
	ResourceTypes []provRes `json:"resourceTypes"`
}

type provRes struct {
	ResourceType string   `json:"resourceType"`
	ApiVersions  []string `json:"apiVersions"`
}
