


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEffectiveVirtualNetworkByNetworkGroup(ctx *pulumi.Context, args *ListEffectiveVirtualNetworkByNetworkGroupArgs, opts ...pulumi.InvokeOption) (*ListEffectiveVirtualNetworkByNetworkGroupResult, error) {
	var rv ListEffectiveVirtualNetworkByNetworkGroupResult
	err := ctx.Invoke("azure-native:network:listEffectiveVirtualNetworkByNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEffectiveVirtualNetworkByNetworkGroupArgs struct {
	NetworkGroupName   string  `pulumi:"networkGroupName"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
}


type ListEffectiveVirtualNetworkByNetworkGroupResult struct {
	SkipToken *string                           `pulumi:"skipToken"`
	Value     []EffectiveVirtualNetworkResponse `pulumi:"value"`
}

func ListEffectiveVirtualNetworkByNetworkGroupOutput(ctx *pulumi.Context, args ListEffectiveVirtualNetworkByNetworkGroupOutputArgs, opts ...pulumi.InvokeOption) ListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEffectiveVirtualNetworkByNetworkGroupResult, error) {
			args := v.(ListEffectiveVirtualNetworkByNetworkGroupArgs)
			r, err := ListEffectiveVirtualNetworkByNetworkGroup(ctx, &args, opts...)
			var s ListEffectiveVirtualNetworkByNetworkGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEffectiveVirtualNetworkByNetworkGroupResultOutput)
}

type ListEffectiveVirtualNetworkByNetworkGroupOutputArgs struct {
	NetworkGroupName   pulumi.StringInput    `pulumi:"networkGroupName"`
	NetworkManagerName pulumi.StringInput    `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListEffectiveVirtualNetworkByNetworkGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEffectiveVirtualNetworkByNetworkGroupArgs)(nil)).Elem()
}


type ListEffectiveVirtualNetworkByNetworkGroupResultOutput struct{ *pulumi.OutputState }

func (ListEffectiveVirtualNetworkByNetworkGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEffectiveVirtualNetworkByNetworkGroupResult)(nil)).Elem()
}

func (o ListEffectiveVirtualNetworkByNetworkGroupResultOutput) ToListEffectiveVirtualNetworkByNetworkGroupResultOutput() ListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return o
}

func (o ListEffectiveVirtualNetworkByNetworkGroupResultOutput) ToListEffectiveVirtualNetworkByNetworkGroupResultOutputWithContext(ctx context.Context) ListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return o
}

func (o ListEffectiveVirtualNetworkByNetworkGroupResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListEffectiveVirtualNetworkByNetworkGroupResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListEffectiveVirtualNetworkByNetworkGroupResultOutput) Value() EffectiveVirtualNetworkResponseArrayOutput {
	return o.ApplyT(func(v ListEffectiveVirtualNetworkByNetworkGroupResult) []EffectiveVirtualNetworkResponse {
		return v.Value
	}).(EffectiveVirtualNetworkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEffectiveVirtualNetworkByNetworkGroupResultOutput{})
}
