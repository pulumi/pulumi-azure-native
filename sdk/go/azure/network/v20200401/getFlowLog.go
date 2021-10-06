


package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFlowLog(ctx *pulumi.Context, args *LookupFlowLogArgs, opts ...pulumi.InvokeOption) (*LookupFlowLogResult, error) {
	var rv LookupFlowLogResult
	err := ctx.Invoke("azure-native:network/v20200401:getFlowLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFlowLogArgs struct {
	FlowLogName        string `pulumi:"flowLogName"`
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupFlowLogResult struct {
	Enabled                    *bool                               `pulumi:"enabled"`
	Etag                       string                              `pulumi:"etag"`
	FlowAnalyticsConfiguration *TrafficAnalyticsPropertiesResponse `pulumi:"flowAnalyticsConfiguration"`
	Format                     *FlowLogFormatParametersResponse    `pulumi:"format"`
	Id                         *string                             `pulumi:"id"`
	Location                   *string                             `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	RetentionPolicy            *RetentionPolicyParametersResponse  `pulumi:"retentionPolicy"`
	StorageId                  string                              `pulumi:"storageId"`
	Tags                       map[string]string                   `pulumi:"tags"`
	TargetResourceGuid         string                              `pulumi:"targetResourceGuid"`
	TargetResourceId           string                              `pulumi:"targetResourceId"`
	Type                       string                              `pulumi:"type"`
}
