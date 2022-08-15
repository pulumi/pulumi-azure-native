


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebPubSubKeys(ctx *pulumi.Context, args *ListWebPubSubKeysArgs, opts ...pulumi.InvokeOption) (*ListWebPubSubKeysResult, error) {
	var rv ListWebPubSubKeysResult
	err := ctx.Invoke("azure-native:webpubsub/v20210401preview:listWebPubSubKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebPubSubKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListWebPubSubKeysResult struct {
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}

func ListWebPubSubKeysOutput(ctx *pulumi.Context, args ListWebPubSubKeysOutputArgs, opts ...pulumi.InvokeOption) ListWebPubSubKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebPubSubKeysResult, error) {
			args := v.(ListWebPubSubKeysArgs)
			r, err := ListWebPubSubKeys(ctx, &args, opts...)
			var s ListWebPubSubKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebPubSubKeysResultOutput)
}

type ListWebPubSubKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListWebPubSubKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebPubSubKeysArgs)(nil)).Elem()
}


type ListWebPubSubKeysResultOutput struct{ *pulumi.OutputState }

func (ListWebPubSubKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebPubSubKeysResult)(nil)).Elem()
}

func (o ListWebPubSubKeysResultOutput) ToListWebPubSubKeysResultOutput() ListWebPubSubKeysResultOutput {
	return o
}

func (o ListWebPubSubKeysResultOutput) ToListWebPubSubKeysResultOutputWithContext(ctx context.Context) ListWebPubSubKeysResultOutput {
	return o
}

func (o ListWebPubSubKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebPubSubKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListWebPubSubKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebPubSubKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListWebPubSubKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebPubSubKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListWebPubSubKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebPubSubKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebPubSubKeysResultOutput{})
}
