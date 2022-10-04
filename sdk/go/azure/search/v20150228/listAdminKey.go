


package v20150228

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAdminKey(ctx *pulumi.Context, args *ListAdminKeyArgs, opts ...pulumi.InvokeOption) (*ListAdminKeyResult, error) {
	var rv ListAdminKeyResult
	err := ctx.Invoke("azure-native:search/v20150228:listAdminKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAdminKeyArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListAdminKeyResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListAdminKeyOutput(ctx *pulumi.Context, args ListAdminKeyOutputArgs, opts ...pulumi.InvokeOption) ListAdminKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAdminKeyResult, error) {
			args := v.(ListAdminKeyArgs)
			r, err := ListAdminKey(ctx, &args, opts...)
			var s ListAdminKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAdminKeyResultOutput)
}

type ListAdminKeyOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (ListAdminKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAdminKeyArgs)(nil)).Elem()
}


type ListAdminKeyResultOutput struct{ *pulumi.OutputState }

func (ListAdminKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAdminKeyResult)(nil)).Elem()
}

func (o ListAdminKeyResultOutput) ToListAdminKeyResultOutput() ListAdminKeyResultOutput {
	return o
}

func (o ListAdminKeyResultOutput) ToListAdminKeyResultOutputWithContext(ctx context.Context) ListAdminKeyResultOutput {
	return o
}

func (o ListAdminKeyResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListAdminKeyResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o ListAdminKeyResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListAdminKeyResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAdminKeyResultOutput{})
}
