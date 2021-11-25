


package v20170501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azure-native:network/v20170501:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupProfileResult struct {
	DnsConfig            *DnsConfigResponse     `pulumi:"dnsConfig"`
	Endpoints            []EndpointResponse     `pulumi:"endpoints"`
	Id                   string                 `pulumi:"id"`
	Location             *string                `pulumi:"location"`
	MonitorConfig        *MonitorConfigResponse `pulumi:"monitorConfig"`
	Name                 string                 `pulumi:"name"`
	ProfileStatus        *string                `pulumi:"profileStatus"`
	Tags                 map[string]string      `pulumi:"tags"`
	TrafficRoutingMethod *string                `pulumi:"trafficRoutingMethod"`
	Type                 string                 `pulumi:"type"`
}
