


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNamedValue(ctx *pulumi.Context, args *ListNamedValueArgs, opts ...pulumi.InvokeOption) (*ListNamedValueResult, error) {
	var rv ListNamedValueResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:listNamedValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNamedValueArgs struct {
	NamedValueId      string `pulumi:"namedValueId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListNamedValueResult struct {
	Value *string `pulumi:"value"`
}

func ListNamedValueOutput(ctx *pulumi.Context, args ListNamedValueOutputArgs, opts ...pulumi.InvokeOption) ListNamedValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNamedValueResult, error) {
			args := v.(ListNamedValueArgs)
			r, err := ListNamedValue(ctx, &args, opts...)
			var s ListNamedValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNamedValueResultOutput)
}

type ListNamedValueOutputArgs struct {
	NamedValueId      pulumi.StringInput `pulumi:"namedValueId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (ListNamedValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamedValueArgs)(nil)).Elem()
}


type ListNamedValueResultOutput struct{ *pulumi.OutputState }

func (ListNamedValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamedValueResult)(nil)).Elem()
}

func (o ListNamedValueResultOutput) ToListNamedValueResultOutput() ListNamedValueResultOutput {
	return o
}

func (o ListNamedValueResultOutput) ToListNamedValueResultOutputWithContext(ctx context.Context) ListNamedValueResultOutput {
	return o
}

func (o ListNamedValueResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNamedValueResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNamedValueResultOutput{})
}
