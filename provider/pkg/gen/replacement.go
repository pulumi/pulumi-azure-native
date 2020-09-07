package gen

import "github.com/pulumi/pulumi/pkg/v2/codegen"

// forceNewMap is a map of Azure Resource Provider -> Resource Name -> Properties that cause replacements.
// API Versions are currently ignored.
var forceNewMap map[string]map[string]codegen.StringSet

func init() {
	forceNewMap = map[string]map[string]codegen.StringSet{
		"Insights": {
			"Component": codegen.NewStringSet(), // covered by x-ms-mutability
		},
		"Network": {
			"PublicIPAddress":  codegen.NewStringSet("location", "publicIPAddressVersion", "publicIPPrefix", "sku"),
			"Subnet":  codegen.NewStringSet(), // no force-news
			"VirtualNetwork": codegen.NewStringSet("location"),
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
			"WebApp": codegen.NewStringSet("location"),
		},
	}
}
