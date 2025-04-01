package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// ModuleData represents one spec file entry like
//
//	Aad:
//	  tracking: "2022-12-01"
type ModuleData struct {
	Tracking string `yaml:"tracking,omitempty"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: program vA-spec.yaml vB-spec.yaml")
		os.Exit(1)
	}

	fileAPath := os.Args[1]
	fileBPath := os.Args[2]

	// Read and parse both YAML files
	servicesA, err := parseYamlFile(fileAPath)
	if err != nil {
		fmt.Printf("Error parsing file %s: %v\n", fileAPath, err)
		os.Exit(1)
	}

	servicesB, err := parseYamlFile(fileBPath)
	if err != nil {
		fmt.Printf("Error parsing file %s: %v\n", fileBPath, err)
		os.Exit(1)
	}

	// Compare tracking values
	compareTrackingValues(servicesA, servicesB)
}

func parseYamlFile(filePath string) (map[string]ModuleData, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var services map[string]ModuleData
	err = yaml.Unmarshal(data, &services)
	if err != nil {
		return nil, err
	}

	return services, nil
}

func compareTrackingValues(servicesA, servicesB map[string]ModuleData) {
	keys := make([]string, 0, len(servicesA))
	for key := range servicesA {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		serviceA := servicesA[key]
		fmt.Println(key)
		if serviceB, exists := servicesB[key]; exists {
			// Skip if either service doesn't have a tracking value
			if serviceA.Tracking == "" || serviceB.Tracking == "" {
				continue
			}

			// Compare tracking values (as strings)
			fmt.Printf("  comparing %q to %q in %s\n", serviceA.Tracking, serviceB.Tracking, key)
			if compareVersions(serviceB.Tracking, serviceA.Tracking) < 0 {
				fmt.Printf("%s: %q to %q\n", key, serviceA.Tracking, serviceB.Tracking)
			}
		}
	}
}

// compareVersions compares two version strings
// Returns -1 if version1 < version2, 0 if equal, 1 if version1 > version2
func compareVersions(version1, version2 string) int {
	// Handle preview versions with special suffix
	isPreview1 := strings.Contains(version1, "preview")
	isPreview2 := strings.Contains(version2, "preview")

	// Extract the base date part for comparison
	datePart1 := strings.Split(version1, "-preview")[0]
	datePart2 := strings.Split(version2, "-preview")[0]

	// Compare the date parts
	if datePart1 < datePart2 {
		return -1
	} else if datePart1 > datePart2 {
		return 1
	}

	// If date parts are equal, preview versions are considered less than non-preview
	if isPreview1 && !isPreview2 {
		return -1
	} else if !isPreview1 && isPreview2 {
		return 1
	}

	// Both are the same (both preview or both not preview with same date)
	return 0
}
