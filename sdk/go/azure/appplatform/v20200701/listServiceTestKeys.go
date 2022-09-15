


package v20200701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListServiceTestKeys(ctx *pulumi.Context, args *ListServiceTestKeysArgs, opts ...pulumi.InvokeOption) (*ListServiceTestKeysResult, error) {
	var rv ListServiceTestKeysResult
	err := ctx.Invoke("azure-native:appplatform/v20200701:listServiceTestKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServiceTestKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListServiceTestKeysResult struct {
	Enabled               *bool   `pulumi:"enabled"`
	PrimaryKey            *string `pulumi:"primaryKey"`
	PrimaryTestEndpoint   *string `pulumi:"primaryTestEndpoint"`
	SecondaryKey          *string `pulumi:"secondaryKey"`
	SecondaryTestEndpoint *string `pulumi:"secondaryTestEndpoint"`
}

func ListServiceTestKeysOutput(ctx *pulumi.Context, args ListServiceTestKeysOutputArgs, opts ...pulumi.InvokeOption) ListServiceTestKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServiceTestKeysResult, error) {
			args := v.(ListServiceTestKeysArgs)
			r, err := ListServiceTestKeys(ctx, &args, opts...)
			var s ListServiceTestKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListServiceTestKeysResultOutput)
}

type ListServiceTestKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (ListServiceTestKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServiceTestKeysArgs)(nil)).Elem()
}


type ListServiceTestKeysResultOutput struct{ *pulumi.OutputState }

func (ListServiceTestKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServiceTestKeysResult)(nil)).Elem()
}

func (o ListServiceTestKeysResultOutput) ToListServiceTestKeysResultOutput() ListServiceTestKeysResultOutput {
	return o
}

func (o ListServiceTestKeysResultOutput) ToListServiceTestKeysResultOutputWithContext(ctx context.Context) ListServiceTestKeysResultOutput {
	return o
}

func (o ListServiceTestKeysResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListServiceTestKeysResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListServiceTestKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServiceTestKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListServiceTestKeysResultOutput) PrimaryTestEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServiceTestKeysResult) *string { return v.PrimaryTestEndpoint }).(pulumi.StringPtrOutput)
}

func (o ListServiceTestKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServiceTestKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func (o ListServiceTestKeysResultOutput) SecondaryTestEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServiceTestKeysResult) *string { return v.SecondaryTestEndpoint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServiceTestKeysResultOutput{})
}
