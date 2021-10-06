


package storsimple

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBandwidthSetting(ctx *pulumi.Context, args *LookupBandwidthSettingArgs, opts ...pulumi.InvokeOption) (*LookupBandwidthSettingResult, error) {
	var rv LookupBandwidthSettingResult
	err := ctx.Invoke("azure-native:storsimple:getBandwidthSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBandwidthSettingArgs struct {
	BandwidthSettingName string `pulumi:"bandwidthSettingName"`
	ManagerName          string `pulumi:"managerName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupBandwidthSettingResult struct {
	Id          string                      `pulumi:"id"`
	Kind        *string                     `pulumi:"kind"`
	Name        string                      `pulumi:"name"`
	Schedules   []BandwidthScheduleResponse `pulumi:"schedules"`
	Type        string                      `pulumi:"type"`
	VolumeCount int                         `pulumi:"volumeCount"`
}
