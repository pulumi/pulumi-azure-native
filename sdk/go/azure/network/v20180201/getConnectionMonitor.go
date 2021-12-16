


package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectionMonitor(ctx *pulumi.Context, args *LookupConnectionMonitorArgs, opts ...pulumi.InvokeOption) (*LookupConnectionMonitorResult, error) {
	var rv LookupConnectionMonitorResult
	err := ctx.Invoke("azure-native:network/v20180201:getConnectionMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectionMonitorArgs struct {
	ConnectionMonitorName string `pulumi:"connectionMonitorName"`
	NetworkWatcherName    string `pulumi:"networkWatcherName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupConnectionMonitorResult struct {
	AutoStart                   *bool                                `pulumi:"autoStart"`
	Destination                 ConnectionMonitorDestinationResponse `pulumi:"destination"`
	Etag                        *string                              `pulumi:"etag"`
	Id                          string                               `pulumi:"id"`
	Location                    *string                              `pulumi:"location"`
	MonitoringIntervalInSeconds *int                                 `pulumi:"monitoringIntervalInSeconds"`
	MonitoringStatus            *string                              `pulumi:"monitoringStatus"`
	Name                        string                               `pulumi:"name"`
	ProvisioningState           *string                              `pulumi:"provisioningState"`
	Source                      ConnectionMonitorSourceResponse      `pulumi:"source"`
	StartTime                   *string                              `pulumi:"startTime"`
	Tags                        map[string]string                    `pulumi:"tags"`
	Type                        string                               `pulumi:"type"`
}


func (val *LookupConnectionMonitorResult) Defaults() *LookupConnectionMonitorResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AutoStart) {
		autoStart_ := true
		tmp.AutoStart = &autoStart_
	}
	if isZero(tmp.Etag) {
		etag_ := "A unique read-only string that changes whenever the resource is updated."
		tmp.Etag = &etag_
	}
	if isZero(tmp.MonitoringIntervalInSeconds) {
		monitoringIntervalInSeconds_ := 60
		tmp.MonitoringIntervalInSeconds = &monitoringIntervalInSeconds_
	}
	return &tmp
}
