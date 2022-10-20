


package v20140901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNamespaceKeys(ctx *pulumi.Context, args *ListNamespaceKeysArgs, opts ...pulumi.InvokeOption) (*ListNamespaceKeysResult, error) {
	var rv ListNamespaceKeysResult
	err := ctx.Invoke("azure-native:notificationhubs/v20140901:listNamespaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNamespaceKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListNamespaceKeysResult struct {
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
}

func ListNamespaceKeysOutput(ctx *pulumi.Context, args ListNamespaceKeysOutputArgs, opts ...pulumi.InvokeOption) ListNamespaceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNamespaceKeysResult, error) {
			args := v.(ListNamespaceKeysArgs)
			r, err := ListNamespaceKeys(ctx, &args, opts...)
			var s ListNamespaceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNamespaceKeysResultOutput)
}

type ListNamespaceKeysOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListNamespaceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamespaceKeysArgs)(nil)).Elem()
}


type ListNamespaceKeysResultOutput struct{ *pulumi.OutputState }

func (ListNamespaceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamespaceKeysResult)(nil)).Elem()
}

func (o ListNamespaceKeysResultOutput) ToListNamespaceKeysResultOutput() ListNamespaceKeysResultOutput {
	return o
}

func (o ListNamespaceKeysResultOutput) ToListNamespaceKeysResultOutputWithContext(ctx context.Context) ListNamespaceKeysResultOutput {
	return o
}

func (o ListNamespaceKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListNamespaceKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNamespaceKeysResultOutput{})
}
