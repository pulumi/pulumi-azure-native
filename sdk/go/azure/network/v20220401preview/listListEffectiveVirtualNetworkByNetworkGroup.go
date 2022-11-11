


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListListEffectiveVirtualNetworkByNetworkGroup(ctx *pulumi.Context, args *ListListEffectiveVirtualNetworkByNetworkGroupArgs, opts ...pulumi.InvokeOption) (*ListListEffectiveVirtualNetworkByNetworkGroupResult, error) {
	var rv ListListEffectiveVirtualNetworkByNetworkGroupResult
	err := ctx.Invoke("azure-native:network/v20220401preview:listListEffectiveVirtualNetworkByNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListEffectiveVirtualNetworkByNetworkGroupArgs struct {
	NetworkGroupName   string  `pulumi:"networkGroupName"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
}


type ListListEffectiveVirtualNetworkByNetworkGroupResult struct {
	SkipToken *string                           `pulumi:"skipToken"`
	Value     []EffectiveVirtualNetworkResponse `pulumi:"value"`
}

func ListListEffectiveVirtualNetworkByNetworkGroupOutput(ctx *pulumi.Context, args ListListEffectiveVirtualNetworkByNetworkGroupOutputArgs, opts ...pulumi.InvokeOption) ListListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListListEffectiveVirtualNetworkByNetworkGroupResult, error) {
			args := v.(ListListEffectiveVirtualNetworkByNetworkGroupArgs)
			r, err := ListListEffectiveVirtualNetworkByNetworkGroup(ctx, &args, opts...)
			var s ListListEffectiveVirtualNetworkByNetworkGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListListEffectiveVirtualNetworkByNetworkGroupResultOutput)
}

type ListListEffectiveVirtualNetworkByNetworkGroupOutputArgs struct {
	NetworkGroupName   pulumi.StringInput    `pulumi:"networkGroupName"`
	NetworkManagerName pulumi.StringInput    `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListListEffectiveVirtualNetworkByNetworkGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListEffectiveVirtualNetworkByNetworkGroupArgs)(nil)).Elem()
}


type ListListEffectiveVirtualNetworkByNetworkGroupResultOutput struct{ *pulumi.OutputState }

func (ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListEffectiveVirtualNetworkByNetworkGroupResult)(nil)).Elem()
}

func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) ToListListEffectiveVirtualNetworkByNetworkGroupResultOutput() ListListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return o
}

func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) ToListListEffectiveVirtualNetworkByNetworkGroupResultOutputWithContext(ctx context.Context) ListListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return o
}

func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListListEffectiveVirtualNetworkByNetworkGroupResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) Value() EffectiveVirtualNetworkResponseArrayOutput {
	return o.ApplyT(func(v ListListEffectiveVirtualNetworkByNetworkGroupResult) []EffectiveVirtualNetworkResponse {
		return v.Value
	}).(EffectiveVirtualNetworkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListListEffectiveVirtualNetworkByNetworkGroupResultOutput{})
}
