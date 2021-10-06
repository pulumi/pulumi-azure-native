


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTag(ctx *pulumi.Context, args *LookupTagArgs, opts ...pulumi.InvokeOption) (*LookupTagResult, error) {
	var rv LookupTagResult
	err := ctx.Invoke("azure-native:apimanagement/v20210401preview:getTag", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TagId             string `pulumi:"tagId"`
}


type LookupTagResult struct {
	DisplayName string `pulumi:"displayName"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	Type        string `pulumi:"type"`
}
