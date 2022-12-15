


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppHybridConnectionSlot(ctx *pulumi.Context, args *LookupWebAppHybridConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppHybridConnectionSlotResult, error) {
	var rv LookupWebAppHybridConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20180201:getWebAppHybridConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppHybridConnectionSlotArgs struct {
	Name              string `pulumi:"name"`
	NamespaceName     string `pulumi:"namespaceName"`
	RelayName         string `pulumi:"relayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppHybridConnectionSlotResult struct {
	Hostname            *string `pulumi:"hostname"`
	Id                  string  `pulumi:"id"`
	Kind                *string `pulumi:"kind"`
	Name                string  `pulumi:"name"`
	Port                *int    `pulumi:"port"`
	RelayArmUri         *string `pulumi:"relayArmUri"`
	RelayName           *string `pulumi:"relayName"`
	SendKeyName         *string `pulumi:"sendKeyName"`
	SendKeyValue        *string `pulumi:"sendKeyValue"`
	ServiceBusNamespace *string `pulumi:"serviceBusNamespace"`
	ServiceBusSuffix    *string `pulumi:"serviceBusSuffix"`
	Type                string  `pulumi:"type"`
}

func LookupWebAppHybridConnectionSlotOutput(ctx *pulumi.Context, args LookupWebAppHybridConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppHybridConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppHybridConnectionSlotResult, error) {
			args := v.(LookupWebAppHybridConnectionSlotArgs)
			r, err := LookupWebAppHybridConnectionSlot(ctx, &args, opts...)
			var s LookupWebAppHybridConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppHybridConnectionSlotResultOutput)
}

type LookupWebAppHybridConnectionSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	RelayName         pulumi.StringInput `pulumi:"relayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppHybridConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppHybridConnectionSlotArgs)(nil)).Elem()
}


type LookupWebAppHybridConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppHybridConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppHybridConnectionSlotResult)(nil)).Elem()
}

func (o LookupWebAppHybridConnectionSlotResultOutput) ToLookupWebAppHybridConnectionSlotResultOutput() LookupWebAppHybridConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppHybridConnectionSlotResultOutput) ToLookupWebAppHybridConnectionSlotResultOutputWithContext(ctx context.Context) LookupWebAppHybridConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppHybridConnectionSlotResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) RelayArmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.RelayArmUri }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) RelayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.RelayName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) SendKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.SendKeyName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) SendKeyValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.SendKeyValue }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) ServiceBusSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.ServiceBusSuffix }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppHybridConnectionSlotResultOutput{})
}
