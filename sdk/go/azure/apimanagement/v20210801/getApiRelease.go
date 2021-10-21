


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiRelease(ctx *pulumi.Context, args *LookupApiReleaseArgs, opts ...pulumi.InvokeOption) (*LookupApiReleaseResult, error) {
	var rv LookupApiReleaseResult
	err := ctx.Invoke("azure-native:apimanagement/v20210801:getApiRelease", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiReleaseArgs struct {
	ApiId             string `pulumi:"apiId"`
	ReleaseId         string `pulumi:"releaseId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiReleaseResult struct {
	ApiId           *string `pulumi:"apiId"`
	CreatedDateTime string  `pulumi:"createdDateTime"`
	Id              string  `pulumi:"id"`
	Name            string  `pulumi:"name"`
	Notes           *string `pulumi:"notes"`
	Type            string  `pulumi:"type"`
	UpdatedDateTime string  `pulumi:"updatedDateTime"`
}
