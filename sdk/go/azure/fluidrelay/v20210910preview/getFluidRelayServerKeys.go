


package v20210910preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFluidRelayServerKeys(ctx *pulumi.Context, args *GetFluidRelayServerKeysArgs, opts ...pulumi.InvokeOption) (*GetFluidRelayServerKeysResult, error) {
	var rv GetFluidRelayServerKeysResult
	err := ctx.Invoke("azure-native:fluidrelay/v20210910preview:getFluidRelayServerKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetFluidRelayServerKeysArgs struct {
	FluidRelayServerName string `pulumi:"fluidRelayServerName"`
	ResourceGroup        string `pulumi:"resourceGroup"`
}


type GetFluidRelayServerKeysResult struct {
	Key1 string `pulumi:"key1"`
	Key2 string `pulumi:"key2"`
}

func GetFluidRelayServerKeysOutput(ctx *pulumi.Context, args GetFluidRelayServerKeysOutputArgs, opts ...pulumi.InvokeOption) GetFluidRelayServerKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFluidRelayServerKeysResult, error) {
			args := v.(GetFluidRelayServerKeysArgs)
			r, err := GetFluidRelayServerKeys(ctx, &args, opts...)
			var s GetFluidRelayServerKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFluidRelayServerKeysResultOutput)
}

type GetFluidRelayServerKeysOutputArgs struct {
	FluidRelayServerName pulumi.StringInput `pulumi:"fluidRelayServerName"`
	ResourceGroup        pulumi.StringInput `pulumi:"resourceGroup"`
}

func (GetFluidRelayServerKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFluidRelayServerKeysArgs)(nil)).Elem()
}


type GetFluidRelayServerKeysResultOutput struct{ *pulumi.OutputState }

func (GetFluidRelayServerKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFluidRelayServerKeysResult)(nil)).Elem()
}

func (o GetFluidRelayServerKeysResultOutput) ToGetFluidRelayServerKeysResultOutput() GetFluidRelayServerKeysResultOutput {
	return o
}

func (o GetFluidRelayServerKeysResultOutput) ToGetFluidRelayServerKeysResultOutputWithContext(ctx context.Context) GetFluidRelayServerKeysResultOutput {
	return o
}

func (o GetFluidRelayServerKeysResultOutput) Key1() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluidRelayServerKeysResult) string { return v.Key1 }).(pulumi.StringOutput)
}

func (o GetFluidRelayServerKeysResultOutput) Key2() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluidRelayServerKeysResult) string { return v.Key2 }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFluidRelayServerKeysResultOutput{})
}
