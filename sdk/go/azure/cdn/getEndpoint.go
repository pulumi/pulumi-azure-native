


package cdn

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("azure-native:cdn:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEndpointResult struct {
	ContentTypesToCompress           []string                                                                    `pulumi:"contentTypesToCompress"`
	DefaultOriginGroup               *ResourceReferenceResponse                                                  `pulumi:"defaultOriginGroup"`
	DeliveryPolicy                   *EndpointPropertiesUpdateParametersResponseDeliveryPolicy                   `pulumi:"deliveryPolicy"`
	GeoFilters                       []GeoFilterResponse                                                         `pulumi:"geoFilters"`
	HostName                         string                                                                      `pulumi:"hostName"`
	Id                               string                                                                      `pulumi:"id"`
	IsCompressionEnabled             *bool                                                                       `pulumi:"isCompressionEnabled"`
	IsHttpAllowed                    *bool                                                                       `pulumi:"isHttpAllowed"`
	IsHttpsAllowed                   *bool                                                                       `pulumi:"isHttpsAllowed"`
	Location                         string                                                                      `pulumi:"location"`
	Name                             string                                                                      `pulumi:"name"`
	OptimizationType                 *string                                                                     `pulumi:"optimizationType"`
	OriginGroups                     []DeepCreatedOriginGroupResponse                                            `pulumi:"originGroups"`
	OriginHostHeader                 *string                                                                     `pulumi:"originHostHeader"`
	OriginPath                       *string                                                                     `pulumi:"originPath"`
	Origins                          []DeepCreatedOriginResponse                                                 `pulumi:"origins"`
	ProbePath                        *string                                                                     `pulumi:"probePath"`
	ProvisioningState                string                                                                      `pulumi:"provisioningState"`
	QueryStringCachingBehavior       *string                                                                     `pulumi:"queryStringCachingBehavior"`
	ResourceState                    string                                                                      `pulumi:"resourceState"`
	SystemData                       SystemDataResponse                                                          `pulumi:"systemData"`
	Tags                             map[string]string                                                           `pulumi:"tags"`
	Type                             string                                                                      `pulumi:"type"`
	UrlSigningKeys                   []UrlSigningKeyResponse                                                     `pulumi:"urlSigningKeys"`
	WebApplicationFirewallPolicyLink *EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink `pulumi:"webApplicationFirewallPolicyLink"`
}
