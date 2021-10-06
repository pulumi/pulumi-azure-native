


package v20191101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:deploymentmanager/v20191101preview:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ServiceName         string `pulumi:"serviceName"`
	ServiceTopologyName string `pulumi:"serviceTopologyName"`
}


type LookupServiceResult struct {
	Id                   string            `pulumi:"id"`
	Location             string            `pulumi:"location"`
	Name                 string            `pulumi:"name"`
	Tags                 map[string]string `pulumi:"tags"`
	TargetLocation       string            `pulumi:"targetLocation"`
	TargetSubscriptionId string            `pulumi:"targetSubscriptionId"`
	Type                 string            `pulumi:"type"`
}
