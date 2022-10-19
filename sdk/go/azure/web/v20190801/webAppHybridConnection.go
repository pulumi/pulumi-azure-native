


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppHybridConnection struct {
	pulumi.CustomResourceState

	Hostname            pulumi.StringPtrOutput `pulumi:"hostname"`
	Kind                pulumi.StringPtrOutput `pulumi:"kind"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	Port                pulumi.IntPtrOutput    `pulumi:"port"`
	RelayArmUri         pulumi.StringPtrOutput `pulumi:"relayArmUri"`
	RelayName           pulumi.StringPtrOutput `pulumi:"relayName"`
	SendKeyName         pulumi.StringPtrOutput `pulumi:"sendKeyName"`
	SendKeyValue        pulumi.StringPtrOutput `pulumi:"sendKeyValue"`
	ServiceBusNamespace pulumi.StringPtrOutput `pulumi:"serviceBusNamespace"`
	ServiceBusSuffix    pulumi.StringPtrOutput `pulumi:"serviceBusSuffix"`
	Type                pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppHybridConnection(ctx *pulumi.Context,
	name string, args *WebAppHybridConnectionArgs, opts ...pulumi.ResourceOption) (*WebAppHybridConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppHybridConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppHybridConnection
	err := ctx.RegisterResource("azure-native:web/v20190801:WebAppHybridConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppHybridConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppHybridConnectionState, opts ...pulumi.ResourceOption) (*WebAppHybridConnection, error) {
	var resource WebAppHybridConnection
	err := ctx.ReadResource("azure-native:web/v20190801:WebAppHybridConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppHybridConnectionState struct {
}

type WebAppHybridConnectionState struct {
}

func (WebAppHybridConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHybridConnectionState)(nil)).Elem()
}

type webAppHybridConnectionArgs struct {
	Hostname            *string `pulumi:"hostname"`
	Kind                *string `pulumi:"kind"`
	Name                string  `pulumi:"name"`
	NamespaceName       string  `pulumi:"namespaceName"`
	Port                *int    `pulumi:"port"`
	RelayArmUri         *string `pulumi:"relayArmUri"`
	RelayName           *string `pulumi:"relayName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	SendKeyName         *string `pulumi:"sendKeyName"`
	SendKeyValue        *string `pulumi:"sendKeyValue"`
	ServiceBusNamespace *string `pulumi:"serviceBusNamespace"`
	ServiceBusSuffix    *string `pulumi:"serviceBusSuffix"`
}


type WebAppHybridConnectionArgs struct {
	Hostname            pulumi.StringPtrInput
	Kind                pulumi.StringPtrInput
	Name                pulumi.StringInput
	NamespaceName       pulumi.StringInput
	Port                pulumi.IntPtrInput
	RelayArmUri         pulumi.StringPtrInput
	RelayName           pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	SendKeyName         pulumi.StringPtrInput
	SendKeyValue        pulumi.StringPtrInput
	ServiceBusNamespace pulumi.StringPtrInput
	ServiceBusSuffix    pulumi.StringPtrInput
}

func (WebAppHybridConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHybridConnectionArgs)(nil)).Elem()
}

type WebAppHybridConnectionInput interface {
	pulumi.Input

	ToWebAppHybridConnectionOutput() WebAppHybridConnectionOutput
	ToWebAppHybridConnectionOutputWithContext(ctx context.Context) WebAppHybridConnectionOutput
}

func (*WebAppHybridConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHybridConnection)(nil)).Elem()
}

func (i *WebAppHybridConnection) ToWebAppHybridConnectionOutput() WebAppHybridConnectionOutput {
	return i.ToWebAppHybridConnectionOutputWithContext(context.Background())
}

func (i *WebAppHybridConnection) ToWebAppHybridConnectionOutputWithContext(ctx context.Context) WebAppHybridConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppHybridConnectionOutput)
}

type WebAppHybridConnectionOutput struct{ *pulumi.OutputState }

func (WebAppHybridConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHybridConnection)(nil)).Elem()
}

func (o WebAppHybridConnectionOutput) ToWebAppHybridConnectionOutput() WebAppHybridConnectionOutput {
	return o
}

func (o WebAppHybridConnectionOutput) ToWebAppHybridConnectionOutputWithContext(ctx context.Context) WebAppHybridConnectionOutput {
	return o
}

func (o WebAppHybridConnectionOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppHybridConnectionOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o WebAppHybridConnectionOutput) RelayArmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.RelayArmUri }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionOutput) RelayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.RelayName }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionOutput) SendKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.SendKeyName }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionOutput) SendKeyValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.SendKeyValue }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionOutput) ServiceBusSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.ServiceBusSuffix }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppHybridConnectionOutput{})
}
