


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEventHub(ctx *pulumi.Context, args *LookupEventHubArgs, opts ...pulumi.InvokeOption) (*LookupEventHubResult, error) {
	var rv LookupEventHubResult
	err := ctx.Invoke("azure-native:eventhub/v20150801:getEventHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubArgs struct {
	EventHubName      string `pulumi:"eventHubName"`
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEventHubResult struct {
	CreatedAt              string   `pulumi:"createdAt"`
	Id                     string   `pulumi:"id"`
	Location               *string  `pulumi:"location"`
	MessageRetentionInDays *float64 `pulumi:"messageRetentionInDays"`
	Name                   string   `pulumi:"name"`
	PartitionCount         *float64 `pulumi:"partitionCount"`
	PartitionIds           []string `pulumi:"partitionIds"`
	Status                 *string  `pulumi:"status"`
	Type                   string   `pulumi:"type"`
	UpdatedAt              string   `pulumi:"updatedAt"`
}

func LookupEventHubOutput(ctx *pulumi.Context, args LookupEventHubOutputArgs, opts ...pulumi.InvokeOption) LookupEventHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventHubResult, error) {
			args := v.(LookupEventHubArgs)
			r, err := LookupEventHub(ctx, &args, opts...)
			var s LookupEventHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventHubResultOutput)
}

type LookupEventHubOutputArgs struct {
	EventHubName      pulumi.StringInput `pulumi:"eventHubName"`
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubArgs)(nil)).Elem()
}


type LookupEventHubResultOutput struct{ *pulumi.OutputState }

func (LookupEventHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubResult)(nil)).Elem()
}

func (o LookupEventHubResultOutput) ToLookupEventHubResultOutput() LookupEventHubResultOutput {
	return o
}

func (o LookupEventHubResultOutput) ToLookupEventHubResultOutputWithContext(ctx context.Context) LookupEventHubResultOutput {
	return o
}

func (o LookupEventHubResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupEventHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEventHubResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubResultOutput) MessageRetentionInDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEventHubResult) *float64 { return v.MessageRetentionInDays }).(pulumi.Float64PtrOutput)
}

func (o LookupEventHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEventHubResultOutput) PartitionCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEventHubResult) *float64 { return v.PartitionCount }).(pulumi.Float64PtrOutput)
}

func (o LookupEventHubResultOutput) PartitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventHubResult) []string { return v.PartitionIds }).(pulumi.StringArrayOutput)
}

func (o LookupEventHubResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupEventHubResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventHubResultOutput{})
}
