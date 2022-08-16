


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFlowLog(ctx *pulumi.Context, args *LookupFlowLogArgs, opts ...pulumi.InvokeOption) (*LookupFlowLogResult, error) {
	var rv LookupFlowLogResult
	err := ctx.Invoke("azure-native:network/v20200601:getFlowLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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


func (val *LookupFlowLogResult) Defaults() *LookupFlowLogResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Format = tmp.Format.Defaults()

	tmp.RetentionPolicy = tmp.RetentionPolicy.Defaults()

	return &tmp
}

func LookupFlowLogOutput(ctx *pulumi.Context, args LookupFlowLogOutputArgs, opts ...pulumi.InvokeOption) LookupFlowLogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFlowLogResult, error) {
			args := v.(LookupFlowLogArgs)
			r, err := LookupFlowLog(ctx, &args, opts...)
			var s LookupFlowLogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFlowLogResultOutput)
}

type LookupFlowLogOutputArgs struct {
	FlowLogName        pulumi.StringInput `pulumi:"flowLogName"`
	NetworkWatcherName pulumi.StringInput `pulumi:"networkWatcherName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFlowLogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlowLogArgs)(nil)).Elem()
}


type LookupFlowLogResultOutput struct{ *pulumi.OutputState }

func (LookupFlowLogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlowLogResult)(nil)).Elem()
}

func (o LookupFlowLogResultOutput) ToLookupFlowLogResultOutput() LookupFlowLogResultOutput {
	return o
}

func (o LookupFlowLogResultOutput) ToLookupFlowLogResultOutputWithContext(ctx context.Context) LookupFlowLogResultOutput {
	return o
}

func (o LookupFlowLogResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFlowLogResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFlowLogResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlowLogResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupFlowLogResultOutput) FlowAnalyticsConfiguration() TrafficAnalyticsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupFlowLogResult) *TrafficAnalyticsPropertiesResponse { return v.FlowAnalyticsConfiguration }).(TrafficAnalyticsPropertiesResponsePtrOutput)
}

func (o LookupFlowLogResultOutput) Format() FlowLogFormatParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupFlowLogResult) *FlowLogFormatParametersResponse { return v.Format }).(FlowLogFormatParametersResponsePtrOutput)
}

func (o LookupFlowLogResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowLogResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupFlowLogResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowLogResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupFlowLogResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlowLogResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFlowLogResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlowLogResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFlowLogResultOutput) RetentionPolicy() RetentionPolicyParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupFlowLogResult) *RetentionPolicyParametersResponse { return v.RetentionPolicy }).(RetentionPolicyParametersResponsePtrOutput)
}

func (o LookupFlowLogResultOutput) StorageId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlowLogResult) string { return v.StorageId }).(pulumi.StringOutput)
}

func (o LookupFlowLogResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFlowLogResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFlowLogResultOutput) TargetResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlowLogResult) string { return v.TargetResourceGuid }).(pulumi.StringOutput)
}

func (o LookupFlowLogResultOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlowLogResult) string { return v.TargetResourceId }).(pulumi.StringOutput)
}

func (o LookupFlowLogResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlowLogResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFlowLogResultOutput{})
}
