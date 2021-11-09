


package v20211006preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("azure-native:hybridconnectivity/v20211006preview:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	EndpointName string `pulumi:"endpointName"`
	ResourceUri  string `pulumi:"resourceUri"`
}


type LookupEndpointResult struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	Id                 string  `pulumi:"id"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
	Name               string  `pulumi:"name"`
	ProvisioningState  string  `pulumi:"provisioningState"`
	ResourceId         *string `pulumi:"resourceId"`
	Type               string  `pulumi:"type"`
}
