


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBandwidthSchedule(ctx *pulumi.Context, args *LookupBandwidthScheduleArgs, opts ...pulumi.InvokeOption) (*LookupBandwidthScheduleResult, error) {
	var rv LookupBandwidthScheduleResult
	err := ctx.Invoke("azure-native:databoxedge/v20210601preview:getBandwidthSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBandwidthScheduleArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBandwidthScheduleResult struct {
	Days       []string           `pulumi:"days"`
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	RateInMbps int                `pulumi:"rateInMbps"`
	Start      string             `pulumi:"start"`
	Stop       string             `pulumi:"stop"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}
