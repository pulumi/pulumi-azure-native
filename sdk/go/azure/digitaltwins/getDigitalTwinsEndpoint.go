


package digitaltwins

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDigitalTwinsEndpoint(ctx *pulumi.Context, args *LookupDigitalTwinsEndpointArgs, opts ...pulumi.InvokeOption) (*LookupDigitalTwinsEndpointResult, error) {
	var rv LookupDigitalTwinsEndpointResult
	err := ctx.Invoke("azure-native:digitaltwins:getDigitalTwinsEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDigitalTwinsEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupDigitalTwinsEndpointResult struct {
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}
