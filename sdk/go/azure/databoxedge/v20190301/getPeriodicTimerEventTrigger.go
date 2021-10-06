


package v20190301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPeriodicTimerEventTrigger(ctx *pulumi.Context, args *LookupPeriodicTimerEventTriggerArgs, opts ...pulumi.InvokeOption) (*LookupPeriodicTimerEventTriggerResult, error) {
	var rv LookupPeriodicTimerEventTriggerResult
	err := ctx.Invoke("azure-native:databoxedge/v20190301:getPeriodicTimerEventTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeriodicTimerEventTriggerArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPeriodicTimerEventTriggerResult struct {
	CustomContextTag *string                         `pulumi:"customContextTag"`
	Id               string                          `pulumi:"id"`
	Kind             string                          `pulumi:"kind"`
	Name             string                          `pulumi:"name"`
	SinkInfo         RoleSinkInfoResponse            `pulumi:"sinkInfo"`
	SourceInfo       PeriodicTimerSourceInfoResponse `pulumi:"sourceInfo"`
	Type             string                          `pulumi:"type"`
}
