


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiVersionSet(ctx *pulumi.Context, args *LookupApiVersionSetArgs, opts ...pulumi.InvokeOption) (*LookupApiVersionSetResult, error) {
	var rv LookupApiVersionSetResult
	err := ctx.Invoke("azure-native:apimanagement/v20210801:getApiVersionSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiVersionSetArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	VersionSetId      string `pulumi:"versionSetId"`
}


type LookupApiVersionSetResult struct {
	Description       *string `pulumi:"description"`
	DisplayName       string  `pulumi:"displayName"`
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	Type              string  `pulumi:"type"`
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	VersionQueryName  *string `pulumi:"versionQueryName"`
	VersioningScheme  string  `pulumi:"versioningScheme"`
}
