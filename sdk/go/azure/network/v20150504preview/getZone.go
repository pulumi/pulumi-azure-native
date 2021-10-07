


package v20150504preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupZone(ctx *pulumi.Context, args *LookupZoneArgs, opts ...pulumi.InvokeOption) (*LookupZoneResult, error) {
	var rv LookupZoneResult
	err := ctx.Invoke("azure-native:network/v20150504preview:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupZoneArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ZoneName          string `pulumi:"zoneName"`
}


type LookupZoneResult struct {
	Etag       *string                `pulumi:"etag"`
	Id         string                 `pulumi:"id"`
	Location   string                 `pulumi:"location"`
	Name       string                 `pulumi:"name"`
	Properties ZonePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string      `pulumi:"tags"`
	Type       string                 `pulumi:"type"`
}
