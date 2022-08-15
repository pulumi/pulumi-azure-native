


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListShareSubscriptionSynchronizationDetails(ctx *pulumi.Context, args *ListShareSubscriptionSynchronizationDetailsArgs, opts ...pulumi.InvokeOption) (*ListShareSubscriptionSynchronizationDetailsResult, error) {
	var rv ListShareSubscriptionSynchronizationDetailsResult
	err := ctx.Invoke("azure-native:datashare/v20210801:listShareSubscriptionSynchronizationDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListShareSubscriptionSynchronizationDetailsArgs struct {
	AccountName           string  `pulumi:"accountName"`
	Filter                *string `pulumi:"filter"`
	Orderby               *string `pulumi:"orderby"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	SkipToken             *string `pulumi:"skipToken"`
	SynchronizationId     string  `pulumi:"synchronizationId"`
}


type ListShareSubscriptionSynchronizationDetailsResult struct {
	NextLink *string                          `pulumi:"nextLink"`
	Value    []SynchronizationDetailsResponse `pulumi:"value"`
}

func ListShareSubscriptionSynchronizationDetailsOutput(ctx *pulumi.Context, args ListShareSubscriptionSynchronizationDetailsOutputArgs, opts ...pulumi.InvokeOption) ListShareSubscriptionSynchronizationDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListShareSubscriptionSynchronizationDetailsResult, error) {
			args := v.(ListShareSubscriptionSynchronizationDetailsArgs)
			r, err := ListShareSubscriptionSynchronizationDetails(ctx, &args, opts...)
			var s ListShareSubscriptionSynchronizationDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListShareSubscriptionSynchronizationDetailsResultOutput)
}

type ListShareSubscriptionSynchronizationDetailsOutputArgs struct {
	AccountName           pulumi.StringInput    `pulumi:"accountName"`
	Filter                pulumi.StringPtrInput `pulumi:"filter"`
	Orderby               pulumi.StringPtrInput `pulumi:"orderby"`
	ResourceGroupName     pulumi.StringInput    `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput    `pulumi:"shareSubscriptionName"`
	SkipToken             pulumi.StringPtrInput `pulumi:"skipToken"`
	SynchronizationId     pulumi.StringInput    `pulumi:"synchronizationId"`
}

func (ListShareSubscriptionSynchronizationDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSubscriptionSynchronizationDetailsArgs)(nil)).Elem()
}


type ListShareSubscriptionSynchronizationDetailsResultOutput struct{ *pulumi.OutputState }

func (ListShareSubscriptionSynchronizationDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSubscriptionSynchronizationDetailsResult)(nil)).Elem()
}

func (o ListShareSubscriptionSynchronizationDetailsResultOutput) ToListShareSubscriptionSynchronizationDetailsResultOutput() ListShareSubscriptionSynchronizationDetailsResultOutput {
	return o
}

func (o ListShareSubscriptionSynchronizationDetailsResultOutput) ToListShareSubscriptionSynchronizationDetailsResultOutputWithContext(ctx context.Context) ListShareSubscriptionSynchronizationDetailsResultOutput {
	return o
}

func (o ListShareSubscriptionSynchronizationDetailsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListShareSubscriptionSynchronizationDetailsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListShareSubscriptionSynchronizationDetailsResultOutput) Value() SynchronizationDetailsResponseArrayOutput {
	return o.ApplyT(func(v ListShareSubscriptionSynchronizationDetailsResult) []SynchronizationDetailsResponse {
		return v.Value
	}).(SynchronizationDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListShareSubscriptionSynchronizationDetailsResultOutput{})
}
