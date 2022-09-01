


package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMediaServiceKeys(ctx *pulumi.Context, args *ListMediaServiceKeysArgs, opts ...pulumi.InvokeOption) (*ListMediaServiceKeysResult, error) {
	var rv ListMediaServiceKeysResult
	err := ctx.Invoke("azure-native:media:listMediaServiceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMediaServiceKeysArgs struct {
	MediaServiceName  string `pulumi:"mediaServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMediaServiceKeysResult struct {
	PrimaryAuthEndpoint   *string `pulumi:"primaryAuthEndpoint"`
	PrimaryKey            *string `pulumi:"primaryKey"`
	Scope                 *string `pulumi:"scope"`
	SecondaryAuthEndpoint *string `pulumi:"secondaryAuthEndpoint"`
	SecondaryKey          *string `pulumi:"secondaryKey"`
}

func ListMediaServiceKeysOutput(ctx *pulumi.Context, args ListMediaServiceKeysOutputArgs, opts ...pulumi.InvokeOption) ListMediaServiceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMediaServiceKeysResult, error) {
			args := v.(ListMediaServiceKeysArgs)
			r, err := ListMediaServiceKeys(ctx, &args, opts...)
			var s ListMediaServiceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMediaServiceKeysResultOutput)
}

type ListMediaServiceKeysOutputArgs struct {
	MediaServiceName  pulumi.StringInput `pulumi:"mediaServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMediaServiceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMediaServiceKeysArgs)(nil)).Elem()
}


type ListMediaServiceKeysResultOutput struct{ *pulumi.OutputState }

func (ListMediaServiceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMediaServiceKeysResult)(nil)).Elem()
}

func (o ListMediaServiceKeysResultOutput) ToListMediaServiceKeysResultOutput() ListMediaServiceKeysResultOutput {
	return o
}

func (o ListMediaServiceKeysResultOutput) ToListMediaServiceKeysResultOutputWithContext(ctx context.Context) ListMediaServiceKeysResultOutput {
	return o
}

func (o ListMediaServiceKeysResultOutput) PrimaryAuthEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMediaServiceKeysResult) *string { return v.PrimaryAuthEndpoint }).(pulumi.StringPtrOutput)
}

func (o ListMediaServiceKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMediaServiceKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListMediaServiceKeysResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMediaServiceKeysResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o ListMediaServiceKeysResultOutput) SecondaryAuthEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMediaServiceKeysResult) *string { return v.SecondaryAuthEndpoint }).(pulumi.StringPtrOutput)
}

func (o ListMediaServiceKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMediaServiceKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMediaServiceKeysResultOutput{})
}
