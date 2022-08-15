


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppHybridConnection(ctx *pulumi.Context, args *LookupWebAppHybridConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppHybridConnectionResult, error) {
	var rv LookupWebAppHybridConnectionResult
	err := ctx.Invoke("azure-native:web/v20220301:getWebAppHybridConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppHybridConnectionArgs struct {
	Name              string `pulumi:"name"`
	NamespaceName     string `pulumi:"namespaceName"`
	RelayName         string `pulumi:"relayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppHybridConnectionResult struct {
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

func LookupWebAppHybridConnectionOutput(ctx *pulumi.Context, args LookupWebAppHybridConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppHybridConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppHybridConnectionResult, error) {
			args := v.(LookupWebAppHybridConnectionArgs)
			r, err := LookupWebAppHybridConnection(ctx, &args, opts...)
			var s LookupWebAppHybridConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppHybridConnectionResultOutput)
}

type LookupWebAppHybridConnectionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	RelayName         pulumi.StringInput `pulumi:"relayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppHybridConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppHybridConnectionArgs)(nil)).Elem()
}


type LookupWebAppHybridConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppHybridConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppHybridConnectionResult)(nil)).Elem()
}

func (o LookupWebAppHybridConnectionResultOutput) ToLookupWebAppHybridConnectionResultOutput() LookupWebAppHybridConnectionResultOutput {
	return o
}

func (o LookupWebAppHybridConnectionResultOutput) ToLookupWebAppHybridConnectionResultOutputWithContext(ctx context.Context) LookupWebAppHybridConnectionResultOutput {
	return o
}

func (o LookupWebAppHybridConnectionResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) RelayArmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) *string { return v.RelayArmUri }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) RelayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) *string { return v.RelayName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) SendKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) *string { return v.SendKeyName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) SendKeyValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) *string { return v.SendKeyValue }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) ServiceBusSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) *string { return v.ServiceBusSuffix }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHybridConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppHybridConnectionResultOutput{})
}
