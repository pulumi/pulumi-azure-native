


package v20160701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListHybridConnectionKeys(ctx *pulumi.Context, args *ListHybridConnectionKeysArgs, opts ...pulumi.InvokeOption) (*ListHybridConnectionKeysResult, error) {
	var rv ListHybridConnectionKeysResult
	err := ctx.Invoke("azure-native:relay/v20160701:listHybridConnectionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListHybridConnectionKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	HybridConnectionName  string `pulumi:"hybridConnectionName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListHybridConnectionKeysResult struct {
	KeyName                   *string `pulumi:"keyName"`
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}

func ListHybridConnectionKeysOutput(ctx *pulumi.Context, args ListHybridConnectionKeysOutputArgs, opts ...pulumi.InvokeOption) ListHybridConnectionKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListHybridConnectionKeysResult, error) {
			args := v.(ListHybridConnectionKeysArgs)
			r, err := ListHybridConnectionKeys(ctx, &args, opts...)
			var s ListHybridConnectionKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListHybridConnectionKeysResultOutput)
}

type ListHybridConnectionKeysOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	HybridConnectionName  pulumi.StringInput `pulumi:"hybridConnectionName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListHybridConnectionKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListHybridConnectionKeysArgs)(nil)).Elem()
}


type ListHybridConnectionKeysResultOutput struct{ *pulumi.OutputState }

func (ListHybridConnectionKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListHybridConnectionKeysResult)(nil)).Elem()
}

func (o ListHybridConnectionKeysResultOutput) ToListHybridConnectionKeysResultOutput() ListHybridConnectionKeysResultOutput {
	return o
}

func (o ListHybridConnectionKeysResultOutput) ToListHybridConnectionKeysResultOutputWithContext(ctx context.Context) ListHybridConnectionKeysResultOutput {
	return o
}

func (o ListHybridConnectionKeysResultOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListHybridConnectionKeysResult) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o ListHybridConnectionKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListHybridConnectionKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListHybridConnectionKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListHybridConnectionKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListHybridConnectionKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListHybridConnectionKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListHybridConnectionKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListHybridConnectionKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListHybridConnectionKeysResultOutput{})
}
