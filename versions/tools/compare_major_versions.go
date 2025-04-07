// Copyright 2025, Pulumi Corporation.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

// ResourceInfo represents the common structure for each resource
type ResourceInfo struct {
	ApiVersion  string `yaml:"ApiVersion"`
	ResourceUri string `yaml:"ResourceUri"`
}

// Config represents the root YAML structure
// Each service contains a map of resource names to their info
type Config map[string]map[string]*ResourceInfo

type Resource struct {
	Module       string
	ResourceName string
	ApiVersion   string
}

type ResourceComparison struct {
	Previous *Resource `json:"previous,omitempty"`
	Next     *Resource `json:"next,omitempty"`
	Path     string
}

type ResourceComparisons []ResourceComparison

func (cc ResourceComparisons) toChanges() Changes {
	added := []ResourceComparison{}
	removed := []ResourceComparison{}
	changed := []ResourceComparison{}
	for _, c := range cc {
		if c.Previous != nil && c.Next == nil {
			removed = append(removed, c)
		}
		if c.Previous == nil && c.Next != nil {
			added = append(added, c)
		}
		if c.Previous != nil && c.Next != nil &&
			(c.Previous.Module != c.Next.Module || c.Previous.ResourceName != c.Next.ResourceName) {
			changed = append(changed, c)
		}
	}
	return Changes{added, removed, changed}
}

type Changes struct {
	Added, Removed, Changed ResourceComparisons
}

func mapConfigByPath(config Config) map[string]*Resource {
	resourceMap := make(map[string]*Resource)
	for service, resources := range config {
		for r, info := range resources {
			resourceMap[info.ResourceUri] = &Resource{service, r, info.ApiVersion}
		}
	}
	return resourceMap
}

func main() {
	// Define command-line flag for version number
	baseVersion := flag.Int("olderMajorVersion", 2, "Base version N to compare vN.yaml with v(N+1).yaml")
	mode := flag.String("out", "changes", "Output mode: `changes` (in JSON) or `table` for a Markdown table of all resources")
	flag.Parse()

	if *mode != "changes" && *mode != "table" {
		fmt.Printf("Invalid output mode: %s\n", *mode)
		flag.Usage()
		os.Exit(1)
	}

	// Construct file paths
	currentFile := filepath.Join("..", fmt.Sprintf("v%d.yaml", *baseVersion))
	nextFile := filepath.Join("..", fmt.Sprintf("v%d.yaml", *baseVersion+1))

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

	if *mode == "changes" {
		currentConfigByPath := mapConfigByPath(currentConfig)
		nextConfigByPath := mapConfigByPath(nextConfig)

		allPaths := make([]string, 0, len(currentConfigByPath)+len(nextConfigByPath))
		for path := range currentConfigByPath {
			allPaths = append(allPaths, path)
		}
		for path := range nextConfigByPath {
			allPaths = append(allPaths, path)
		}
		sort.Strings(allPaths)

		var comparisons ResourceComparisons = make([]ResourceComparison, 0, len(allPaths))
		for _, path := range allPaths {
			cur := currentConfigByPath[path]
			next := nextConfigByPath[path]
			comparisons = append(comparisons, ResourceComparison{cur, next, path})
		}

		changes := comparisons.toChanges()
		changesJson, err := json.MarshalIndent(changes, "", "  ")
		if err != nil {
			fmt.Printf("Error marshalling changes: %v\n", err)
			return
		}
		fmt.Println(string(changesJson))
	}

	if *mode == "table" {
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
		fmt.Printf("| Service | Resource | REST version in v%d | REST version in v%d |\n", *baseVersion, *baseVersion+1)
		fmt.Println("|---|---|---|---|")

		// Print rows
		for _, service := range servicesList {
			// Collect all resources for this service
			resources := make(map[string]bool)
			if currentResources, ok := currentConfig[service]; ok {
				for resource := range currentResources {
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

				fmt.Printf("| %s | %s | %s | %s |\n", service, resource, currentVersion, nextVersion)
			}
		}
	}
}
