


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFileEventTrigger(ctx *pulumi.Context, args *LookupFileEventTriggerArgs, opts ...pulumi.InvokeOption) (*LookupFileEventTriggerResult, error) {
	var rv LookupFileEventTriggerResult
	err := ctx.Invoke("azure-native:databoxedge/v20220301:getFileEventTrigger", args, &rv, opts...)
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
	SystemData       SystemDataResponse     `pulumi:"systemData"`
	Type             string                 `pulumi:"type"`
}

func LookupFileEventTriggerOutput(ctx *pulumi.Context, args LookupFileEventTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupFileEventTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileEventTriggerResult, error) {
			args := v.(LookupFileEventTriggerArgs)
			r, err := LookupFileEventTrigger(ctx, &args, opts...)
			var s LookupFileEventTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileEventTriggerResultOutput)
}

type LookupFileEventTriggerOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFileEventTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileEventTriggerArgs)(nil)).Elem()
}


type LookupFileEventTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupFileEventTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileEventTriggerResult)(nil)).Elem()
}

func (o LookupFileEventTriggerResultOutput) ToLookupFileEventTriggerResultOutput() LookupFileEventTriggerResultOutput {
	return o
}

func (o LookupFileEventTriggerResultOutput) ToLookupFileEventTriggerResultOutputWithContext(ctx context.Context) LookupFileEventTriggerResultOutput {
	return o
}

func (o LookupFileEventTriggerResultOutput) CustomContextTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) *string { return v.CustomContextTag }).(pulumi.StringPtrOutput)
}

func (o LookupFileEventTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFileEventTriggerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupFileEventTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFileEventTriggerResultOutput) SinkInfo() RoleSinkInfoResponseOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) RoleSinkInfoResponse { return v.SinkInfo }).(RoleSinkInfoResponseOutput)
}

func (o LookupFileEventTriggerResultOutput) SourceInfo() FileSourceInfoResponseOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) FileSourceInfoResponse { return v.SourceInfo }).(FileSourceInfoResponseOutput)
}

func (o LookupFileEventTriggerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFileEventTriggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileEventTriggerResultOutput{})
}
