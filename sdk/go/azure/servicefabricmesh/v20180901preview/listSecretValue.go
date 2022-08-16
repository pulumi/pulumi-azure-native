


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSecretValue(ctx *pulumi.Context, args *ListSecretValueArgs, opts ...pulumi.InvokeOption) (*ListSecretValueResult, error) {
	var rv ListSecretValueResult
	err := ctx.Invoke("azure-native:servicefabricmesh/v20180901preview:listSecretValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSecretValueArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SecretResourceName      string `pulumi:"secretResourceName"`
	SecretValueResourceName string `pulumi:"secretValueResourceName"`
}


type ListSecretValueResult struct {
	Value *string `pulumi:"value"`
}

func ListSecretValueOutput(ctx *pulumi.Context, args ListSecretValueOutputArgs, opts ...pulumi.InvokeOption) ListSecretValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSecretValueResult, error) {
			args := v.(ListSecretValueArgs)
			r, err := ListSecretValue(ctx, &args, opts...)
			var s ListSecretValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSecretValueResultOutput)
}

type ListSecretValueOutputArgs struct {
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	SecretResourceName      pulumi.StringInput `pulumi:"secretResourceName"`
	SecretValueResourceName pulumi.StringInput `pulumi:"secretValueResourceName"`
}

func (ListSecretValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSecretValueArgs)(nil)).Elem()
}


type ListSecretValueResultOutput struct{ *pulumi.OutputState }

func (ListSecretValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSecretValueResult)(nil)).Elem()
}

func (o ListSecretValueResultOutput) ToListSecretValueResultOutput() ListSecretValueResultOutput {
	return o
}

func (o ListSecretValueResultOutput) ToListSecretValueResultOutputWithContext(ctx context.Context) ListSecretValueResultOutput {
	return o
}

func (o ListSecretValueResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSecretValueResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSecretValueResultOutput{})
}
