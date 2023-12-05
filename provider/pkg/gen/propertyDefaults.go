package gen

import "strings"

// In some cases, we need to set default values for properties that are not given as inputs.
// One case is when the property was set and is then removed from the input, but Azure treats
// the omission as "no change" instead of removing the value.
// propertyDefaults returns a map of property names to default values for the given module
// and resource, for recording in metadata for later use.
func propertyDefaults(module, resourceName string) map[string]interface{} {
	if resourceName == "StorageAccount" && (module == "storage" || strings.HasPrefix(module, "storage/")) {
		return map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"bypass":              "AzureServices",
				"defaultAction":       "Allow",
				"ipRules":             []interface{}{},
				"virtualNetworkRules": []interface{}{},
			},
		}
	}
	return nil
}
