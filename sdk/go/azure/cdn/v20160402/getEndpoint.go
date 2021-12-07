


package v20160402

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("azure-native:cdn/v20160402:getEndpoint", args, &rv, opts...)
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
	ContentTypesToCompress     []string                    `pulumi:"contentTypesToCompress"`
	HostName                   string                      `pulumi:"hostName"`
	Id                         string                      `pulumi:"id"`
	IsCompressionEnabled       *bool                       `pulumi:"isCompressionEnabled"`
	IsHttpAllowed              *bool                       `pulumi:"isHttpAllowed"`
	IsHttpsAllowed             *bool                       `pulumi:"isHttpsAllowed"`
	Location                   string                      `pulumi:"location"`
	Name                       string                      `pulumi:"name"`
	OriginHostHeader           *string                     `pulumi:"originHostHeader"`
	OriginPath                 *string                     `pulumi:"originPath"`
	Origins                    []DeepCreatedOriginResponse `pulumi:"origins"`
	ProvisioningState          string                      `pulumi:"provisioningState"`
	QueryStringCachingBehavior *string                     `pulumi:"queryStringCachingBehavior"`
	ResourceState              string                      `pulumi:"resourceState"`
	Tags                       map[string]string           `pulumi:"tags"`
	Type                       string                      `pulumi:"type"`
}
