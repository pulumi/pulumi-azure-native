


package v20200901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMonitoringConfig(ctx *pulumi.Context, args *LookupMonitoringConfigArgs, opts ...pulumi.InvokeOption) (*LookupMonitoringConfigResult, error) {
	var rv LookupMonitoringConfigResult
	err := ctx.Invoke("azure-native:databoxedge/v20200901preview:getMonitoringConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitoringConfigArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoleName          string `pulumi:"roleName"`
}


type LookupMonitoringConfigResult struct {
	Id                   string                        `pulumi:"id"`
	MetricConfigurations []MetricConfigurationResponse `pulumi:"metricConfigurations"`
	Name                 string                        `pulumi:"name"`
	Type                 string                        `pulumi:"type"`
}
