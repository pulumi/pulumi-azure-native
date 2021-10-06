


package iotsecurity

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLocationSite(ctx *pulumi.Context, args *LookupLocationSiteArgs, opts ...pulumi.InvokeOption) (*LookupLocationSiteResult, error) {
	var rv LookupLocationSiteResult
	err := ctx.Invoke("azure-native:iotsecurity:getLocationSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocationSiteArgs struct {
	IotDefenderLocation string `pulumi:"iotDefenderLocation"`
	SiteName            string `pulumi:"siteName"`
}


type LookupLocationSiteResult struct {
	DisplayName string             `pulumi:"displayName"`
	Id          string             `pulumi:"id"`
	Name        string             `pulumi:"name"`
	SystemData  SystemDataResponse `pulumi:"systemData"`
	Tags        map[string]string  `pulumi:"tags"`
	Type        string             `pulumi:"type"`
}
