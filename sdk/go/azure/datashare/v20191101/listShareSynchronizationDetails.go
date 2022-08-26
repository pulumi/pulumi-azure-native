


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListShareSynchronizationDetails(ctx *pulumi.Context, args *ListShareSynchronizationDetailsArgs, opts ...pulumi.InvokeOption) (*ListShareSynchronizationDetailsResult, error) {
	var rv ListShareSynchronizationDetailsResult
	err := ctx.Invoke("azure-native:datashare/v20191101:listShareSynchronizationDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListShareSynchronizationDetailsArgs struct {
	AccountName        string  `pulumi:"accountName"`
	ConsumerEmail      *string `pulumi:"consumerEmail"`
	ConsumerName       *string `pulumi:"consumerName"`
	ConsumerTenantName *string `pulumi:"consumerTenantName"`
	DurationMs         *int    `pulumi:"durationMs"`
	EndTime            *string `pulumi:"endTime"`
	Filter             *string `pulumi:"filter"`
	Message            *string `pulumi:"message"`
	Orderby            *string `pulumi:"orderby"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	ShareName          string  `pulumi:"shareName"`
	SkipToken          *string `pulumi:"skipToken"`
	StartTime          *string `pulumi:"startTime"`
	Status             *string `pulumi:"status"`
	SynchronizationId  *string `pulumi:"synchronizationId"`
}


type ListShareSynchronizationDetailsResult struct {
	NextLink *string                          `pulumi:"nextLink"`
	Value    []SynchronizationDetailsResponse `pulumi:"value"`
}

func ListShareSynchronizationDetailsOutput(ctx *pulumi.Context, args ListShareSynchronizationDetailsOutputArgs, opts ...pulumi.InvokeOption) ListShareSynchronizationDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListShareSynchronizationDetailsResult, error) {
			args := v.(ListShareSynchronizationDetailsArgs)
			r, err := ListShareSynchronizationDetails(ctx, &args, opts...)
			var s ListShareSynchronizationDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListShareSynchronizationDetailsResultOutput)
}

type ListShareSynchronizationDetailsOutputArgs struct {
	AccountName        pulumi.StringInput    `pulumi:"accountName"`
	ConsumerEmail      pulumi.StringPtrInput `pulumi:"consumerEmail"`
	ConsumerName       pulumi.StringPtrInput `pulumi:"consumerName"`
	ConsumerTenantName pulumi.StringPtrInput `pulumi:"consumerTenantName"`
	DurationMs         pulumi.IntPtrInput    `pulumi:"durationMs"`
	EndTime            pulumi.StringPtrInput `pulumi:"endTime"`
	Filter             pulumi.StringPtrInput `pulumi:"filter"`
	Message            pulumi.StringPtrInput `pulumi:"message"`
	Orderby            pulumi.StringPtrInput `pulumi:"orderby"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	ShareName          pulumi.StringInput    `pulumi:"shareName"`
	SkipToken          pulumi.StringPtrInput `pulumi:"skipToken"`
	StartTime          pulumi.StringPtrInput `pulumi:"startTime"`
	Status             pulumi.StringPtrInput `pulumi:"status"`
	SynchronizationId  pulumi.StringPtrInput `pulumi:"synchronizationId"`
}

func (ListShareSynchronizationDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSynchronizationDetailsArgs)(nil)).Elem()
}


type ListShareSynchronizationDetailsResultOutput struct{ *pulumi.OutputState }

func (ListShareSynchronizationDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSynchronizationDetailsResult)(nil)).Elem()
}

func (o ListShareSynchronizationDetailsResultOutput) ToListShareSynchronizationDetailsResultOutput() ListShareSynchronizationDetailsResultOutput {
	return o
}

func (o ListShareSynchronizationDetailsResultOutput) ToListShareSynchronizationDetailsResultOutputWithContext(ctx context.Context) ListShareSynchronizationDetailsResultOutput {
	return o
}

func (o ListShareSynchronizationDetailsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListShareSynchronizationDetailsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListShareSynchronizationDetailsResultOutput) Value() SynchronizationDetailsResponseArrayOutput {
	return o.ApplyT(func(v ListShareSynchronizationDetailsResult) []SynchronizationDetailsResponse { return v.Value }).(SynchronizationDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListShareSynchronizationDetailsResultOutput{})
}
