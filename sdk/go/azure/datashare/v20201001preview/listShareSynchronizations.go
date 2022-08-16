


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListShareSynchronizations(ctx *pulumi.Context, args *ListShareSynchronizationsArgs, opts ...pulumi.InvokeOption) (*ListShareSynchronizationsResult, error) {
	var rv ListShareSynchronizationsResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:listShareSynchronizations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListShareSynchronizationsArgs struct {
	AccountName       string  `pulumi:"accountName"`
	Filter            *string `pulumi:"filter"`
	Orderby           *string `pulumi:"orderby"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ShareName         string  `pulumi:"shareName"`
	SkipToken         *string `pulumi:"skipToken"`
}


type ListShareSynchronizationsResult struct {
	NextLink *string                        `pulumi:"nextLink"`
	Value    []ShareSynchronizationResponse `pulumi:"value"`
}

func ListShareSynchronizationsOutput(ctx *pulumi.Context, args ListShareSynchronizationsOutputArgs, opts ...pulumi.InvokeOption) ListShareSynchronizationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListShareSynchronizationsResult, error) {
			args := v.(ListShareSynchronizationsArgs)
			r, err := ListShareSynchronizations(ctx, &args, opts...)
			var s ListShareSynchronizationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListShareSynchronizationsResultOutput)
}

type ListShareSynchronizationsOutputArgs struct {
	AccountName       pulumi.StringInput    `pulumi:"accountName"`
	Filter            pulumi.StringPtrInput `pulumi:"filter"`
	Orderby           pulumi.StringPtrInput `pulumi:"orderby"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput    `pulumi:"shareName"`
	SkipToken         pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListShareSynchronizationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSynchronizationsArgs)(nil)).Elem()
}


type ListShareSynchronizationsResultOutput struct{ *pulumi.OutputState }

func (ListShareSynchronizationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSynchronizationsResult)(nil)).Elem()
}

func (o ListShareSynchronizationsResultOutput) ToListShareSynchronizationsResultOutput() ListShareSynchronizationsResultOutput {
	return o
}

func (o ListShareSynchronizationsResultOutput) ToListShareSynchronizationsResultOutputWithContext(ctx context.Context) ListShareSynchronizationsResultOutput {
	return o
}

func (o ListShareSynchronizationsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListShareSynchronizationsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListShareSynchronizationsResultOutput) Value() ShareSynchronizationResponseArrayOutput {
	return o.ApplyT(func(v ListShareSynchronizationsResult) []ShareSynchronizationResponse { return v.Value }).(ShareSynchronizationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListShareSynchronizationsResultOutput{})
}
