


package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSite(ctx *pulumi.Context, args *LookupSiteArgs, opts ...pulumi.InvokeOption) (*LookupSiteResult, error) {
	var rv LookupSiteResult
	err := ctx.Invoke("azure-native:offazure/v20200101:getSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SiteName          string `pulumi:"siteName"`
}


type LookupSiteResult struct {
	ETag       *string                `pulumi:"eTag"`
	Id         string                 `pulumi:"id"`
	Location   *string                `pulumi:"location"`
	Name       *string                `pulumi:"name"`
	Properties SitePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string      `pulumi:"tags"`
	Type       string                 `pulumi:"type"`
}
