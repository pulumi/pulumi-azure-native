package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

// ResourceInfo represents the common structure for each resource
type ResourceInfo struct {
	ApiVersion string `yaml:"ApiVersion"`
}

var caseChanges = map[string]string{
	"monitor":                                         "Monitor",
	"azureADMetric":                                   "AzureADMetric",
	"privateLinkForAzureAd":                           "PrivateLinkForAzureAd",
	"vNetPeering":                                     "VNetPeering",
	"agentPool":                                       "AgentPool",
	"storageSpaceRetrieve":                            "StorageSpaceRetrieve",
	"virtualNetworkRetrieve":                          "VirtualNetworkRetrieve",
	"guestDiagnosticsSetting":                         "GuestDiagnosticsSetting",
	"privateLinkServicesForEDMUpload":                 "PrivateLinkServicesForEDMUpload",
	"privateLinkServicesForM365ComplianceCenter":      "PrivateLinkServicesForM365ComplianceCenter",
	"privateLinkServicesForM365SecurityCenter":        "PrivateLinkServicesForM365SecurityCenter",
	"privateLinkServicesForMIPPolicySync":             "PrivateLinkServicesForMIPPolicySync",
	"privateLinkServicesForO365ManagementActivityAPI": "PrivateLinkServicesForO365ManagementActivityAPI",
	"privateLinkServicesForSCCPowershell":             "PrivateLinkServicesForSCCPowershell",
	"kustoPool":                                       "KustoPool",
}

// Config represents the root YAML structure
// Each service contains a map of resource names to their info
type Config map[string]map[string]*ResourceInfo

func main() {
	// Define command-line flag for version number
	baseVersion := flag.Int("v", 2, "Base version number N to compare vN.yaml with v(N+1).yaml")
	flag.Parse()

	// Construct file paths
	currentFile := filepath.Join(".", fmt.Sprintf("v%d.yaml", *baseVersion))
	nextFile := filepath.Join(".", fmt.Sprintf("v%d.yaml", *baseVersion+1))

	// Read YAML files
	currentData, err := os.ReadFile(currentFile)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", currentFile, err)
		return
	}

	nextData, err := os.ReadFile(nextFile)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", nextFile, err)
		return
	}

	// Parse YAML files
	var currentConfig Config
	var nextConfig Config
	if err := yaml.Unmarshal(currentData, &currentConfig); err != nil {
		fmt.Printf("Error parsing %s: %v\n", currentFile, err)
		return
	}
	if err := yaml.Unmarshal(nextData, &nextConfig); err != nil {
		fmt.Printf("Error parsing %s: %v\n", nextFile, err)
		return
	}

	// Collect all services
	services := make(map[string]bool)
	for service := range currentConfig {
		services[service] = true
	}
	for service := range nextConfig {
		services[service] = true
	}

	// Convert to sorted slice
	servicesList := make([]string, 0, len(services))
	for service := range services {
		servicesList = append(servicesList, service)
	}
	sort.Strings(servicesList)

	// Print markdown table
	fmt.Printf("|Service|Resource|REST version in v%d|REST version in v%d|\n", *baseVersion, *baseVersion+1)
	fmt.Println("|---|---|---|---|")

	// Print rows
	for _, service := range servicesList {
		// Collect all resources for this service
		resources := make(map[string]bool)
		if currentResources, ok := currentConfig[service]; ok {
			for resource := range currentResources {
				// #2366
				if newResource, ok := caseChanges[resource]; ok {
					resource = newResource
				}
				resources[resource] = true
			}
		}
		if nextResources, ok := nextConfig[service]; ok {
			for resource := range nextResources {
				resources[resource] = true
			}
		}

		// Convert to sorted slice
		resourcesList := make([]string, 0, len(resources))
		for resource := range resources {
			resourcesList = append(resourcesList, resource)
		}
		sort.Strings(resourcesList)

		// Print each resource
		for _, resource := range resourcesList {
			currentVersion := "not present"
			nextVersion := "not present"

			if currentResources, ok := currentConfig[service]; ok {
				if currentResource := currentResources[resource]; currentResource != nil {
					currentVersion = currentResource.ApiVersion
				}
			}
			if nextResources, ok := nextConfig[service]; ok {
				if nextResource := nextResources[resource]; nextResource != nil {
					nextVersion = nextResource.ApiVersion
				}
			}

			fmt.Printf("|%s|%s|%s|%s|\n", service, resource, currentVersion, nextVersion)
		}
	}
}
