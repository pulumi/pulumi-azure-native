


package iotsecurity

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSite(ctx *pulumi.Context, args *LookupSiteArgs, opts ...pulumi.InvokeOption) (*LookupSiteResult, error) {
	var rv LookupSiteResult
	err := ctx.Invoke("azure-native:iotsecurity:getSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteArgs struct {
	Scope string `pulumi:"scope"`
}


type LookupSiteResult struct {
	DisplayName string             `pulumi:"displayName"`
	Id          string             `pulumi:"id"`
	Name        string             `pulumi:"name"`
	SystemData  SystemDataResponse `pulumi:"systemData"`
	Tags        map[string]string  `pulumi:"tags"`
	Type        string             `pulumi:"type"`
}
