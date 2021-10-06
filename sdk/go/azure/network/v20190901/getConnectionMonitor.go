


package v20190901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectionMonitor(ctx *pulumi.Context, args *LookupConnectionMonitorArgs, opts ...pulumi.InvokeOption) (*LookupConnectionMonitorResult, error) {
	var rv LookupConnectionMonitorResult
	err := ctx.Invoke("azure-native:network/v20190901:getConnectionMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
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
	ProvisioningState           string                               `pulumi:"provisioningState"`
	Source                      ConnectionMonitorSourceResponse      `pulumi:"source"`
	StartTime                   *string                              `pulumi:"startTime"`
	Tags                        map[string]string                    `pulumi:"tags"`
	Type                        string                               `pulumi:"type"`
}
