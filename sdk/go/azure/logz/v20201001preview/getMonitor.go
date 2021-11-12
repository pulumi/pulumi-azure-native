


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMonitor(ctx *pulumi.Context, args *LookupMonitorArgs, opts ...pulumi.InvokeOption) (*LookupMonitorResult, error) {
	var rv LookupMonitorResult
	err := ctx.Invoke("azure-native:logz/v20201001preview:getMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitorArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupMonitorResult struct {
	Id         string                      `pulumi:"id"`
	Identity   *IdentityPropertiesResponse `pulumi:"identity"`
	Location   string                      `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties MonitorPropertiesResponse   `pulumi:"properties"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}
