


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectionMonitor(ctx *pulumi.Context, args *LookupConnectionMonitorArgs, opts ...pulumi.InvokeOption) (*LookupConnectionMonitorResult, error) {
	var rv LookupConnectionMonitorResult
	err := ctx.Invoke("azure-native:network/v20200601:getConnectionMonitor", args, &rv, opts...)
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
	AutoStart                   *bool                                        `pulumi:"autoStart"`
	ConnectionMonitorType       string                                       `pulumi:"connectionMonitorType"`
	Destination                 *ConnectionMonitorDestinationResponse        `pulumi:"destination"`
	Endpoints                   []ConnectionMonitorEndpointResponse          `pulumi:"endpoints"`
	Etag                        string                                       `pulumi:"etag"`
	Id                          string                                       `pulumi:"id"`
	Location                    *string                                      `pulumi:"location"`
	MonitoringIntervalInSeconds *int                                         `pulumi:"monitoringIntervalInSeconds"`
	MonitoringStatus            string                                       `pulumi:"monitoringStatus"`
	Name                        string                                       `pulumi:"name"`
	Notes                       *string                                      `pulumi:"notes"`
	Outputs                     []ConnectionMonitorOutputResponse            `pulumi:"outputs"`
	ProvisioningState           string                                       `pulumi:"provisioningState"`
	Source                      *ConnectionMonitorSourceResponse             `pulumi:"source"`
	StartTime                   string                                       `pulumi:"startTime"`
	Tags                        map[string]string                            `pulumi:"tags"`
	TestConfigurations          []ConnectionMonitorTestConfigurationResponse `pulumi:"testConfigurations"`
	TestGroups                  []ConnectionMonitorTestGroupResponse         `pulumi:"testGroups"`
	Type                        string                                       `pulumi:"type"`
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
	if isZero(tmp.MonitoringIntervalInSeconds) {
		monitoringIntervalInSeconds_ := 60
		tmp.MonitoringIntervalInSeconds = &monitoringIntervalInSeconds_
	}
	return &tmp
}
