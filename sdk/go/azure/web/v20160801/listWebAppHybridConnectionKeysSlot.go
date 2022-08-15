


package v20160801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppHybridConnectionKeysSlot(ctx *pulumi.Context, args *ListWebAppHybridConnectionKeysSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppHybridConnectionKeysSlotResult, error) {
	var rv ListWebAppHybridConnectionKeysSlotResult
	err := ctx.Invoke("azure-native:web/v20160801:listWebAppHybridConnectionKeysSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppHybridConnectionKeysSlotArgs struct {
	Name              string `pulumi:"name"`
	NamespaceName     string `pulumi:"namespaceName"`
	RelayName         string `pulumi:"relayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppHybridConnectionKeysSlotResult struct {
	Id           string  `pulumi:"id"`
	Kind         *string `pulumi:"kind"`
	Name         string  `pulumi:"name"`
	SendKeyName  string  `pulumi:"sendKeyName"`
	SendKeyValue string  `pulumi:"sendKeyValue"`
	Type         string  `pulumi:"type"`
}

func ListWebAppHybridConnectionKeysSlotOutput(ctx *pulumi.Context, args ListWebAppHybridConnectionKeysSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppHybridConnectionKeysSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppHybridConnectionKeysSlotResult, error) {
			args := v.(ListWebAppHybridConnectionKeysSlotArgs)
			r, err := ListWebAppHybridConnectionKeysSlot(ctx, &args, opts...)
			var s ListWebAppHybridConnectionKeysSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppHybridConnectionKeysSlotResultOutput)
}

type ListWebAppHybridConnectionKeysSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	RelayName         pulumi.StringInput `pulumi:"relayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppHybridConnectionKeysSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHybridConnectionKeysSlotArgs)(nil)).Elem()
}


type ListWebAppHybridConnectionKeysSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppHybridConnectionKeysSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHybridConnectionKeysSlotResult)(nil)).Elem()
}

func (o ListWebAppHybridConnectionKeysSlotResultOutput) ToListWebAppHybridConnectionKeysSlotResultOutput() ListWebAppHybridConnectionKeysSlotResultOutput {
	return o
}

func (o ListWebAppHybridConnectionKeysSlotResultOutput) ToListWebAppHybridConnectionKeysSlotResultOutputWithContext(ctx context.Context) ListWebAppHybridConnectionKeysSlotResultOutput {
	return o
}

func (o ListWebAppHybridConnectionKeysSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppHybridConnectionKeysSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppHybridConnectionKeysSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppHybridConnectionKeysSlotResultOutput) SendKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysSlotResult) string { return v.SendKeyName }).(pulumi.StringOutput)
}

func (o ListWebAppHybridConnectionKeysSlotResultOutput) SendKeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysSlotResult) string { return v.SendKeyValue }).(pulumi.StringOutput)
}

func (o ListWebAppHybridConnectionKeysSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppHybridConnectionKeysSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppHybridConnectionKeysSlotResultOutput{})
}
