


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppRelayServiceConnection struct {
	pulumi.CustomResourceState

	BiztalkUri               pulumi.StringPtrOutput `pulumi:"biztalkUri"`
	EntityConnectionString   pulumi.StringPtrOutput `pulumi:"entityConnectionString"`
	EntityName               pulumi.StringPtrOutput `pulumi:"entityName"`
	Hostname                 pulumi.StringPtrOutput `pulumi:"hostname"`
	Kind                     pulumi.StringPtrOutput `pulumi:"kind"`
	Name                     pulumi.StringOutput    `pulumi:"name"`
	Port                     pulumi.IntPtrOutput    `pulumi:"port"`
	ResourceConnectionString pulumi.StringPtrOutput `pulumi:"resourceConnectionString"`
	ResourceType             pulumi.StringPtrOutput `pulumi:"resourceType"`
	Type                     pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppRelayServiceConnection(ctx *pulumi.Context,
	name string, args *WebAppRelayServiceConnectionArgs, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppRelayServiceConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppRelayServiceConnection
	err := ctx.RegisterResource("azure-native:web/v20210101:WebAppRelayServiceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppRelayServiceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppRelayServiceConnectionState, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnection, error) {
	var resource WebAppRelayServiceConnection
	err := ctx.ReadResource("azure-native:web/v20210101:WebAppRelayServiceConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppRelayServiceConnectionState struct {
}

type WebAppRelayServiceConnectionState struct {
}

func (WebAppRelayServiceConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionState)(nil)).Elem()
}

type webAppRelayServiceConnectionArgs struct {
	BiztalkUri               *string `pulumi:"biztalkUri"`
	EntityConnectionString   *string `pulumi:"entityConnectionString"`
	EntityName               *string `pulumi:"entityName"`
	Hostname                 *string `pulumi:"hostname"`
	Kind                     *string `pulumi:"kind"`
	Name                     string  `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	ResourceGroupName        string  `pulumi:"resourceGroupName"`
	ResourceType             *string `pulumi:"resourceType"`
}


type WebAppRelayServiceConnectionArgs struct {
	BiztalkUri               pulumi.StringPtrInput
	EntityConnectionString   pulumi.StringPtrInput
	EntityName               pulumi.StringPtrInput
	Hostname                 pulumi.StringPtrInput
	Kind                     pulumi.StringPtrInput
	Name                     pulumi.StringInput
	Port                     pulumi.IntPtrInput
	ResourceConnectionString pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	ResourceType             pulumi.StringPtrInput
}

func (WebAppRelayServiceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionArgs)(nil)).Elem()
}

type WebAppRelayServiceConnectionInput interface {
	pulumi.Input

	ToWebAppRelayServiceConnectionOutput() WebAppRelayServiceConnectionOutput
	ToWebAppRelayServiceConnectionOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionOutput
}

func (*WebAppRelayServiceConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppRelayServiceConnection)(nil)).Elem()
}

func (i *WebAppRelayServiceConnection) ToWebAppRelayServiceConnectionOutput() WebAppRelayServiceConnectionOutput {
	return i.ToWebAppRelayServiceConnectionOutputWithContext(context.Background())
}

func (i *WebAppRelayServiceConnection) ToWebAppRelayServiceConnectionOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppRelayServiceConnectionOutput)
}

type WebAppRelayServiceConnectionOutput struct{ *pulumi.OutputState }

func (WebAppRelayServiceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppRelayServiceConnection)(nil)).Elem()
}

func (o WebAppRelayServiceConnectionOutput) ToWebAppRelayServiceConnectionOutput() WebAppRelayServiceConnectionOutput {
	return o
}

func (o WebAppRelayServiceConnectionOutput) ToWebAppRelayServiceConnectionOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionOutput {
	return o
}

func (o WebAppRelayServiceConnectionOutput) BiztalkUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.BiztalkUri }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) EntityConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.EntityConnectionString }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) EntityName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.EntityName }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppRelayServiceConnectionOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) ResourceConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.ResourceConnectionString }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppRelayServiceConnectionOutput{})
}
