


package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagerActivationKey(ctx *pulumi.Context, args *ListManagerActivationKeyArgs, opts ...pulumi.InvokeOption) (*ListManagerActivationKeyResult, error) {
	var rv ListManagerActivationKeyResult
	err := ctx.Invoke("azure-native:storsimple:listManagerActivationKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagerActivationKeyArgs struct {
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListManagerActivationKeyResult struct {
	ActivationKey string `pulumi:"activationKey"`
}

func ListManagerActivationKeyOutput(ctx *pulumi.Context, args ListManagerActivationKeyOutputArgs, opts ...pulumi.InvokeOption) ListManagerActivationKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagerActivationKeyResult, error) {
			args := v.(ListManagerActivationKeyArgs)
			r, err := ListManagerActivationKey(ctx, &args, opts...)
			var s ListManagerActivationKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagerActivationKeyResultOutput)
}

type ListManagerActivationKeyOutputArgs struct {
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListManagerActivationKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagerActivationKeyArgs)(nil)).Elem()
}


type ListManagerActivationKeyResultOutput struct{ *pulumi.OutputState }

func (ListManagerActivationKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagerActivationKeyResult)(nil)).Elem()
}

func (o ListManagerActivationKeyResultOutput) ToListManagerActivationKeyResultOutput() ListManagerActivationKeyResultOutput {
	return o
}

func (o ListManagerActivationKeyResultOutput) ToListManagerActivationKeyResultOutputWithContext(ctx context.Context) ListManagerActivationKeyResultOutput {
	return o
}

func (o ListManagerActivationKeyResultOutput) ActivationKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagerActivationKeyResult) string { return v.ActivationKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagerActivationKeyResultOutput{})
}
