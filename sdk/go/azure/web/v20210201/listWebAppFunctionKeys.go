


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppFunctionKeys(ctx *pulumi.Context, args *ListWebAppFunctionKeysArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionKeysResult, error) {
	var rv ListWebAppFunctionKeysResult
	err := ctx.Invoke("azure-native:web/v20210201:listWebAppFunctionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionKeysArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppFunctionKeysResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}

func ListWebAppFunctionKeysOutput(ctx *pulumi.Context, args ListWebAppFunctionKeysOutputArgs, opts ...pulumi.InvokeOption) ListWebAppFunctionKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppFunctionKeysResult, error) {
			args := v.(ListWebAppFunctionKeysArgs)
			r, err := ListWebAppFunctionKeys(ctx, &args, opts...)
			var s ListWebAppFunctionKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppFunctionKeysResultOutput)
}

type ListWebAppFunctionKeysOutputArgs struct {
	FunctionName      pulumi.StringInput `pulumi:"functionName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppFunctionKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionKeysArgs)(nil)).Elem()
}


type ListWebAppFunctionKeysResultOutput struct{ *pulumi.OutputState }

func (ListWebAppFunctionKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionKeysResult)(nil)).Elem()
}

func (o ListWebAppFunctionKeysResultOutput) ToListWebAppFunctionKeysResultOutput() ListWebAppFunctionKeysResultOutput {
	return o
}

func (o ListWebAppFunctionKeysResultOutput) ToListWebAppFunctionKeysResultOutputWithContext(ctx context.Context) ListWebAppFunctionKeysResultOutput {
	return o
}

func (o ListWebAppFunctionKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppFunctionKeysResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppFunctionKeysResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppFunctionKeysResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListWebAppFunctionKeysResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppFunctionKeysResultOutput{})
}
