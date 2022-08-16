


package v20220421

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFluidRelayServerKeys(ctx *pulumi.Context, args *ListFluidRelayServerKeysArgs, opts ...pulumi.InvokeOption) (*ListFluidRelayServerKeysResult, error) {
	var rv ListFluidRelayServerKeysResult
	err := ctx.Invoke("azure-native:fluidrelay/v20220421:listFluidRelayServerKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFluidRelayServerKeysArgs struct {
	FluidRelayServerName string `pulumi:"fluidRelayServerName"`
	ResourceGroup        string `pulumi:"resourceGroup"`
}


type ListFluidRelayServerKeysResult struct {
	Key1 string `pulumi:"key1"`
	Key2 string `pulumi:"key2"`
}

func ListFluidRelayServerKeysOutput(ctx *pulumi.Context, args ListFluidRelayServerKeysOutputArgs, opts ...pulumi.InvokeOption) ListFluidRelayServerKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFluidRelayServerKeysResult, error) {
			args := v.(ListFluidRelayServerKeysArgs)
			r, err := ListFluidRelayServerKeys(ctx, &args, opts...)
			var s ListFluidRelayServerKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFluidRelayServerKeysResultOutput)
}

type ListFluidRelayServerKeysOutputArgs struct {
	FluidRelayServerName pulumi.StringInput `pulumi:"fluidRelayServerName"`
	ResourceGroup        pulumi.StringInput `pulumi:"resourceGroup"`
}

func (ListFluidRelayServerKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFluidRelayServerKeysArgs)(nil)).Elem()
}


type ListFluidRelayServerKeysResultOutput struct{ *pulumi.OutputState }

func (ListFluidRelayServerKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFluidRelayServerKeysResult)(nil)).Elem()
}

func (o ListFluidRelayServerKeysResultOutput) ToListFluidRelayServerKeysResultOutput() ListFluidRelayServerKeysResultOutput {
	return o
}

func (o ListFluidRelayServerKeysResultOutput) ToListFluidRelayServerKeysResultOutputWithContext(ctx context.Context) ListFluidRelayServerKeysResultOutput {
	return o
}

func (o ListFluidRelayServerKeysResultOutput) Key1() pulumi.StringOutput {
	return o.ApplyT(func(v ListFluidRelayServerKeysResult) string { return v.Key1 }).(pulumi.StringOutput)
}

func (o ListFluidRelayServerKeysResultOutput) Key2() pulumi.StringOutput {
	return o.ApplyT(func(v ListFluidRelayServerKeysResult) string { return v.Key2 }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFluidRelayServerKeysResultOutput{})
}
