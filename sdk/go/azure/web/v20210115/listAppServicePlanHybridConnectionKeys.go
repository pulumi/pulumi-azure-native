


package v20210115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAppServicePlanHybridConnectionKeys(ctx *pulumi.Context, args *ListAppServicePlanHybridConnectionKeysArgs, opts ...pulumi.InvokeOption) (*ListAppServicePlanHybridConnectionKeysResult, error) {
	var rv ListAppServicePlanHybridConnectionKeysResult
	err := ctx.Invoke("azure-native:web/v20210115:listAppServicePlanHybridConnectionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAppServicePlanHybridConnectionKeysArgs struct {
	Name              string `pulumi:"name"`
	NamespaceName     string `pulumi:"namespaceName"`
	RelayName         string `pulumi:"relayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAppServicePlanHybridConnectionKeysResult struct {
	Id           string  `pulumi:"id"`
	Kind         *string `pulumi:"kind"`
	Name         string  `pulumi:"name"`
	SendKeyName  string  `pulumi:"sendKeyName"`
	SendKeyValue string  `pulumi:"sendKeyValue"`
	Type         string  `pulumi:"type"`
}

func ListAppServicePlanHybridConnectionKeysOutput(ctx *pulumi.Context, args ListAppServicePlanHybridConnectionKeysOutputArgs, opts ...pulumi.InvokeOption) ListAppServicePlanHybridConnectionKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAppServicePlanHybridConnectionKeysResult, error) {
			args := v.(ListAppServicePlanHybridConnectionKeysArgs)
			r, err := ListAppServicePlanHybridConnectionKeys(ctx, &args, opts...)
			var s ListAppServicePlanHybridConnectionKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAppServicePlanHybridConnectionKeysResultOutput)
}

type ListAppServicePlanHybridConnectionKeysOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	RelayName         pulumi.StringInput `pulumi:"relayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAppServicePlanHybridConnectionKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAppServicePlanHybridConnectionKeysArgs)(nil)).Elem()
}


type ListAppServicePlanHybridConnectionKeysResultOutput struct{ *pulumi.OutputState }

func (ListAppServicePlanHybridConnectionKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAppServicePlanHybridConnectionKeysResult)(nil)).Elem()
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) ToListAppServicePlanHybridConnectionKeysResultOutput() ListAppServicePlanHybridConnectionKeysResultOutput {
	return o
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) ToListAppServicePlanHybridConnectionKeysResultOutputWithContext(ctx context.Context) ListAppServicePlanHybridConnectionKeysResultOutput {
	return o
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) SendKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) string { return v.SendKeyName }).(pulumi.StringOutput)
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) SendKeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) string { return v.SendKeyValue }).(pulumi.StringOutput)
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAppServicePlanHybridConnectionKeysResultOutput{})
}
