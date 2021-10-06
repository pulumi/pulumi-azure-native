


package v20180820preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:healthcareapis/v20180820preview:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupServiceResult struct {
	Etag       *string                    `pulumi:"etag"`
	Id         string                     `pulumi:"id"`
	Identity   *ResourceResponseIdentity  `pulumi:"identity"`
	Kind       string                     `pulumi:"kind"`
	Location   string                     `pulumi:"location"`
	Name       string                     `pulumi:"name"`
	Properties ServicesPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string          `pulumi:"tags"`
	Type       string                     `pulumi:"type"`
}
