


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEffectiveVirtualNetworkByNetworkManager(ctx *pulumi.Context, args *ListEffectiveVirtualNetworkByNetworkManagerArgs, opts ...pulumi.InvokeOption) (*ListEffectiveVirtualNetworkByNetworkManagerResult, error) {
	var rv ListEffectiveVirtualNetworkByNetworkManagerResult
	err := ctx.Invoke("azure-native:network:listEffectiveVirtualNetworkByNetworkManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEffectiveVirtualNetworkByNetworkManagerArgs struct {
	ConditionalMembers *string `pulumi:"conditionalMembers"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
	Top                *int    `pulumi:"top"`
}


type ListEffectiveVirtualNetworkByNetworkManagerResult struct {
	SkipToken *string                           `pulumi:"skipToken"`
	Value     []EffectiveVirtualNetworkResponse `pulumi:"value"`
}

func ListEffectiveVirtualNetworkByNetworkManagerOutput(ctx *pulumi.Context, args ListEffectiveVirtualNetworkByNetworkManagerOutputArgs, opts ...pulumi.InvokeOption) ListEffectiveVirtualNetworkByNetworkManagerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEffectiveVirtualNetworkByNetworkManagerResult, error) {
			args := v.(ListEffectiveVirtualNetworkByNetworkManagerArgs)
			r, err := ListEffectiveVirtualNetworkByNetworkManager(ctx, &args, opts...)
			var s ListEffectiveVirtualNetworkByNetworkManagerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEffectiveVirtualNetworkByNetworkManagerResultOutput)
}

type ListEffectiveVirtualNetworkByNetworkManagerOutputArgs struct {
	ConditionalMembers pulumi.StringPtrInput `pulumi:"conditionalMembers"`
	NetworkManagerName pulumi.StringInput    `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput `pulumi:"skipToken"`
	Top                pulumi.IntPtrInput    `pulumi:"top"`
}

func (ListEffectiveVirtualNetworkByNetworkManagerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEffectiveVirtualNetworkByNetworkManagerArgs)(nil)).Elem()
}


type ListEffectiveVirtualNetworkByNetworkManagerResultOutput struct{ *pulumi.OutputState }

func (ListEffectiveVirtualNetworkByNetworkManagerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEffectiveVirtualNetworkByNetworkManagerResult)(nil)).Elem()
}

func (o ListEffectiveVirtualNetworkByNetworkManagerResultOutput) ToListEffectiveVirtualNetworkByNetworkManagerResultOutput() ListEffectiveVirtualNetworkByNetworkManagerResultOutput {
	return o
}

func (o ListEffectiveVirtualNetworkByNetworkManagerResultOutput) ToListEffectiveVirtualNetworkByNetworkManagerResultOutputWithContext(ctx context.Context) ListEffectiveVirtualNetworkByNetworkManagerResultOutput {
	return o
}

func (o ListEffectiveVirtualNetworkByNetworkManagerResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListEffectiveVirtualNetworkByNetworkManagerResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListEffectiveVirtualNetworkByNetworkManagerResultOutput) Value() EffectiveVirtualNetworkResponseArrayOutput {
	return o.ApplyT(func(v ListEffectiveVirtualNetworkByNetworkManagerResult) []EffectiveVirtualNetworkResponse {
		return v.Value
	}).(EffectiveVirtualNetworkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEffectiveVirtualNetworkByNetworkManagerResultOutput{})
}
