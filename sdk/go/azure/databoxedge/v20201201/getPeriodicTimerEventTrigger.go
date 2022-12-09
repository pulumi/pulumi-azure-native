


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPeriodicTimerEventTrigger(ctx *pulumi.Context, args *LookupPeriodicTimerEventTriggerArgs, opts ...pulumi.InvokeOption) (*LookupPeriodicTimerEventTriggerResult, error) {
	var rv LookupPeriodicTimerEventTriggerResult
	err := ctx.Invoke("azure-native:databoxedge/v20201201:getPeriodicTimerEventTrigger", args, &rv, opts...)
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
	SystemData       SystemDataResponse              `pulumi:"systemData"`
	Type             string                          `pulumi:"type"`
}

func LookupPeriodicTimerEventTriggerOutput(ctx *pulumi.Context, args LookupPeriodicTimerEventTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupPeriodicTimerEventTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPeriodicTimerEventTriggerResult, error) {
			args := v.(LookupPeriodicTimerEventTriggerArgs)
			r, err := LookupPeriodicTimerEventTrigger(ctx, &args, opts...)
			var s LookupPeriodicTimerEventTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPeriodicTimerEventTriggerResultOutput)
}

type LookupPeriodicTimerEventTriggerOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPeriodicTimerEventTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeriodicTimerEventTriggerArgs)(nil)).Elem()
}


type LookupPeriodicTimerEventTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupPeriodicTimerEventTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeriodicTimerEventTriggerResult)(nil)).Elem()
}

func (o LookupPeriodicTimerEventTriggerResultOutput) ToLookupPeriodicTimerEventTriggerResultOutput() LookupPeriodicTimerEventTriggerResultOutput {
	return o
}

func (o LookupPeriodicTimerEventTriggerResultOutput) ToLookupPeriodicTimerEventTriggerResultOutputWithContext(ctx context.Context) LookupPeriodicTimerEventTriggerResultOutput {
	return o
}

func (o LookupPeriodicTimerEventTriggerResultOutput) CustomContextTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPeriodicTimerEventTriggerResult) *string { return v.CustomContextTag }).(pulumi.StringPtrOutput)
}

func (o LookupPeriodicTimerEventTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeriodicTimerEventTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPeriodicTimerEventTriggerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeriodicTimerEventTriggerResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupPeriodicTimerEventTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeriodicTimerEventTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPeriodicTimerEventTriggerResultOutput) SinkInfo() RoleSinkInfoResponseOutput {
	return o.ApplyT(func(v LookupPeriodicTimerEventTriggerResult) RoleSinkInfoResponse { return v.SinkInfo }).(RoleSinkInfoResponseOutput)
}

func (o LookupPeriodicTimerEventTriggerResultOutput) SourceInfo() PeriodicTimerSourceInfoResponseOutput {
	return o.ApplyT(func(v LookupPeriodicTimerEventTriggerResult) PeriodicTimerSourceInfoResponse { return v.SourceInfo }).(PeriodicTimerSourceInfoResponseOutput)
}

func (o LookupPeriodicTimerEventTriggerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPeriodicTimerEventTriggerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPeriodicTimerEventTriggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeriodicTimerEventTriggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPeriodicTimerEventTriggerResultOutput{})
}
