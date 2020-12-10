package gen

import (
	"fmt"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/debug"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/resources"
)

// FlattenParams takes the parameters specified in Azure API specs/ARM templates and
// flattens them to match the desired format for the Pulumi schema.
// resourceParams is a mapping of API parameter names to provider.AzureApiParameter
// and types is a mapping for the API type names to provider.AzureApiType.
// The latter two can be derived from the metadata generated during schema generation.
func FlattenParams(
	input map[string]interface{},
	resourceParams map[string]resources.AzureAPIParameter,
	types map[string]resources.AzureAPIType,
) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	converter := resources.SdkShapeConverter{Types: types}
	for k, v := range input {
		switch k {
		case "If-Match": // TODO: Not handled in schema
			continue
		case "api-version", "apiVersion", "subscriptionId":
			continue // No need to emit these since we auto inject them
		}

		paramMetadata, ok := resourceParams[k]
		if !ok {
			debug.Log("missing item '%s' from resource metadata for resource, skipping", k)
			continue
		}
		// body parameters are folded in
		if paramMetadata.Body != nil {
			inBody, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("expect body for param %s to be a map, received %T", k, v)
			}
			flattened := converter.BodyPropertiesToSDK(paramMetadata.Body.Properties, inBody)
			mergeMap(out, flattened)
			continue
		}

		// replace param name with value in SdkName if provided.
		if paramMetadata.Value != nil && paramMetadata.Value.SdkName != "" {
			k = paramMetadata.Value.SdkName
		}

		out[k] = v
	}
	return out, nil
}

func mergeMap(dst map[string]interface{}, src map[string]interface{}) {
	for k, v := range src {
		dst[k] = v
	}
}
