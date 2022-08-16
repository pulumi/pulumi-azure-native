


package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppHybridConnectionKeys(ctx *pulumi.Context, args *ListWebAppHybridConnectionKeysArgs, opts ...pulumi.InvokeOption) (*ListWebAppHybridConnectionKeysResult, error) {
	var rv ListWebAppHybridConnectionKeysResult
	err := ctx.Invoke("azure-native:web/v20181101:listWebAppHybridConnectionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppHybridConnectionKeysArgs struct {
	Name              string `pulumi:"name"`
	NamespaceName     string `pulumi:"namespaceName"`
	RelayName         string `pulumi:"relayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppHybridConnectionKeysResult struct {
	Id           string  `pulumi:"id"`
	Kind         *string `pulumi:"kind"`
	Name         string  `pulumi:"name"`
	SendKeyName  string  `pulumi:"sendKeyName"`
	SendKeyValue string  `pulumi:"sendKeyValue"`
	Type         string  `pulumi:"type"`
}

func ListWebAppHybridConnectionKeysOutput(ctx *pulumi.Context, args ListWebAppHybridConnectionKeysOutputArgs, opts ...pulumi.InvokeOption) ListWebAppHybridConnectionKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppHybridConnectionKeysResult, error) {
			args := v.(ListWebAppHybridConnectionKeysArgs)
			r, err := ListWebAppHybridConnectionKeys(ctx, &args, opts...)
			var s ListWebAppHybridConnectionKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppHybridConnectionKeysResultOutput)
}

type ListWebAppHybridConnectionKeysOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	RelayName         pulumi.StringInput `pulumi:"relayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppHybridConnectionKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHybridConnectionKeysArgs)(nil)).Elem()
}


type ListWebAppHybridConnectionKeysResultOutput struct{ *pulumi.OutputState }

func (ListWebAppHybridConnectionKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHybridConnectionKeysResult)(nil)).Elem()
}

func (o ListWebAppHybridConnectionKeysResultOutput) ToListWebAppHybridConnectionKeysResultOutput() ListWebAppHybridConnectionKeysResultOutput {
	return o
}

func (o ListWebAppHybridConnectionKeysResultOutput) ToListWebAppHybridConnectionKeysResultOutputWithContext(ctx context.Context) ListWebAppHybridConnectionKeysResultOutput {
	return o
}

func (o ListWebAppHybridConnectionKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppHybridConnectionKeysResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppHybridConnectionKeysResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppHybridConnectionKeysResultOutput) SendKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysResult) string { return v.SendKeyName }).(pulumi.StringOutput)
}

func (o ListWebAppHybridConnectionKeysResultOutput) SendKeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysResult) string { return v.SendKeyValue }).(pulumi.StringOutput)
}

func (o ListWebAppHybridConnectionKeysResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppHybridConnectionKeysResultOutput{})
}
