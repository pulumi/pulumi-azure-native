


package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemTopic(ctx *pulumi.Context, args *LookupSystemTopicArgs, opts ...pulumi.InvokeOption) (*LookupSystemTopicResult, error) {
	var rv LookupSystemTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getSystemTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSystemTopicArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SystemTopicName   string `pulumi:"systemTopicName"`
}


type LookupSystemTopicResult struct {
	Id                string                `pulumi:"id"`
	Identity          *IdentityInfoResponse `pulumi:"identity"`
	Location          string                `pulumi:"location"`
	MetricResourceId  string                `pulumi:"metricResourceId"`
	Name              string                `pulumi:"name"`
	ProvisioningState string                `pulumi:"provisioningState"`
	Source            *string               `pulumi:"source"`
	SystemData        SystemDataResponse    `pulumi:"systemData"`
	Tags              map[string]string     `pulumi:"tags"`
	TopicType         *string               `pulumi:"topicType"`
	Type              string                `pulumi:"type"`
}

func LookupSystemTopicOutput(ctx *pulumi.Context, args LookupSystemTopicOutputArgs, opts ...pulumi.InvokeOption) LookupSystemTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemTopicResult, error) {
			args := v.(LookupSystemTopicArgs)
			r, err := LookupSystemTopic(ctx, &args, opts...)
			var s LookupSystemTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemTopicResultOutput)
}

type LookupSystemTopicOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SystemTopicName   pulumi.StringInput `pulumi:"systemTopicName"`
}

func (LookupSystemTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemTopicArgs)(nil)).Elem()
}


type LookupSystemTopicResultOutput struct{ *pulumi.OutputState }

func (LookupSystemTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemTopicResult)(nil)).Elem()
}

func (o LookupSystemTopicResultOutput) ToLookupSystemTopicResultOutput() LookupSystemTopicResultOutput {
	return o
}

func (o LookupSystemTopicResultOutput) ToLookupSystemTopicResultOutputWithContext(ctx context.Context) LookupSystemTopicResultOutput {
	return o
}

func (o LookupSystemTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemTopicResultOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) *IdentityInfoResponse { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

func (o LookupSystemTopicResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSystemTopicResultOutput) MetricResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.MetricResourceId }).(pulumi.StringOutput)
}

func (o LookupSystemTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemTopicResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSystemTopicResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o LookupSystemTopicResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSystemTopicResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSystemTopicResultOutput) TopicType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) *string { return v.TopicType }).(pulumi.StringPtrOutput)
}

func (o LookupSystemTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemTopicResultOutput{})
}
