


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFileEventTrigger(ctx *pulumi.Context, args *LookupFileEventTriggerArgs, opts ...pulumi.InvokeOption) (*LookupFileEventTriggerResult, error) {
	var rv LookupFileEventTriggerResult
	err := ctx.Invoke("azure-native:databoxedge/v20190801:getFileEventTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileEventTriggerArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFileEventTriggerResult struct {
	CustomContextTag *string                `pulumi:"customContextTag"`
	Id               string                 `pulumi:"id"`
	Kind             string                 `pulumi:"kind"`
	Name             string                 `pulumi:"name"`
	SinkInfo         RoleSinkInfoResponse   `pulumi:"sinkInfo"`
	SourceInfo       FileSourceInfoResponse `pulumi:"sourceInfo"`
	Type             string                 `pulumi:"type"`
}
