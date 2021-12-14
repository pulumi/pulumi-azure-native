


package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHyperVSite(ctx *pulumi.Context, args *LookupHyperVSiteArgs, opts ...pulumi.InvokeOption) (*LookupHyperVSiteResult, error) {
	var rv LookupHyperVSiteResult
	err := ctx.Invoke("azure-native:offazure/v20200101:getHyperVSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHyperVSiteArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SiteName          string `pulumi:"siteName"`
}


type LookupHyperVSiteResult struct {
	ETag       *string                `pulumi:"eTag"`
	Id         string                 `pulumi:"id"`
	Location   *string                `pulumi:"location"`
	Name       *string                `pulumi:"name"`
	Properties SitePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string      `pulumi:"tags"`
	Type       string                 `pulumi:"type"`
}
