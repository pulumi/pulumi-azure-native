


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConsumerGroup(ctx *pulumi.Context, args *LookupConsumerGroupArgs, opts ...pulumi.InvokeOption) (*LookupConsumerGroupResult, error) {
	var rv LookupConsumerGroupResult
	err := ctx.Invoke("azure-native:eventhub/v20221001preview:getConsumerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsumerGroupArgs struct {
	ConsumerGroupName string `pulumi:"consumerGroupName"`
	EventHubName      string `pulumi:"eventHubName"`
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConsumerGroupResult struct {
	CreatedAt    string             `pulumi:"createdAt"`
	Id           string             `pulumi:"id"`
	Location     string             `pulumi:"location"`
	Name         string             `pulumi:"name"`
	SystemData   SystemDataResponse `pulumi:"systemData"`
	Type         string             `pulumi:"type"`
	UpdatedAt    string             `pulumi:"updatedAt"`
	UserMetadata *string            `pulumi:"userMetadata"`
}

func LookupConsumerGroupOutput(ctx *pulumi.Context, args LookupConsumerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupConsumerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConsumerGroupResult, error) {
			args := v.(LookupConsumerGroupArgs)
			r, err := LookupConsumerGroup(ctx, &args, opts...)
			var s LookupConsumerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConsumerGroupResultOutput)
}

type LookupConsumerGroupOutputArgs struct {
	ConsumerGroupName pulumi.StringInput `pulumi:"consumerGroupName"`
	EventHubName      pulumi.StringInput `pulumi:"eventHubName"`
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConsumerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsumerGroupArgs)(nil)).Elem()
}


type LookupConsumerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupConsumerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsumerGroupResult)(nil)).Elem()
}

func (o LookupConsumerGroupResultOutput) ToLookupConsumerGroupResultOutput() LookupConsumerGroupResultOutput {
	return o
}

func (o LookupConsumerGroupResultOutput) ToLookupConsumerGroupResultOutputWithContext(ctx context.Context) LookupConsumerGroupResultOutput {
	return o
}

func (o LookupConsumerGroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsumerGroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupConsumerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsumerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConsumerGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsumerGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupConsumerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsumerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConsumerGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConsumerGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConsumerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsumerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupConsumerGroupResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsumerGroupResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupConsumerGroupResultOutput) UserMetadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConsumerGroupResult) *string { return v.UserMetadata }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConsumerGroupResultOutput{})
}
