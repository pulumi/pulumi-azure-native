


package v20200201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorApiKeys(ctx *pulumi.Context, args *ListMonitorApiKeysArgs, opts ...pulumi.InvokeOption) (*ListMonitorApiKeysResult, error) {
	var rv ListMonitorApiKeysResult
	err := ctx.Invoke("azure-native:datadog/v20200201preview:listMonitorApiKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorApiKeysArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMonitorApiKeysResult struct {
	NextLink *string                 `pulumi:"nextLink"`
	Value    []DatadogApiKeyResponse `pulumi:"value"`
}
