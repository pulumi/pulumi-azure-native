


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRedisKeys(ctx *pulumi.Context, args *ListRedisKeysArgs, opts ...pulumi.InvokeOption) (*ListRedisKeysResult, error) {
	var rv ListRedisKeysResult
	err := ctx.Invoke("azure-native:cache/v20210601:listRedisKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRedisKeysArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListRedisKeysResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListRedisKeysOutput(ctx *pulumi.Context, args ListRedisKeysOutputArgs, opts ...pulumi.InvokeOption) ListRedisKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRedisKeysResult, error) {
			args := v.(ListRedisKeysArgs)
			r, err := ListRedisKeys(ctx, &args, opts...)
			var s ListRedisKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRedisKeysResultOutput)
}

type ListRedisKeysOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListRedisKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRedisKeysArgs)(nil)).Elem()
}


type ListRedisKeysResultOutput struct{ *pulumi.OutputState }

func (ListRedisKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRedisKeysResult)(nil)).Elem()
}

func (o ListRedisKeysResultOutput) ToListRedisKeysResultOutput() ListRedisKeysResultOutput {
	return o
}

func (o ListRedisKeysResultOutput) ToListRedisKeysResultOutputWithContext(ctx context.Context) ListRedisKeysResultOutput {
	return o
}

func (o ListRedisKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListRedisKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o ListRedisKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListRedisKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRedisKeysResultOutput{})
}
