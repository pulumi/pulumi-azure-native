


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BotConnection struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput                    `pulumi:"etag"`
	Kind       pulumi.StringPtrOutput                    `pulumi:"kind"`
	Location   pulumi.StringPtrOutput                    `pulumi:"location"`
	Name       pulumi.StringOutput                       `pulumi:"name"`
	Properties ConnectionSettingPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput                      `pulumi:"sku"`
	Tags       pulumi.StringMapOutput                    `pulumi:"tags"`
	Type       pulumi.StringOutput                       `pulumi:"type"`
	Zones      pulumi.StringArrayOutput                  `pulumi:"zones"`
}


func NewBotConnection(ctx *pulumi.Context,
	name string, args *BotConnectionArgs, opts ...pulumi.ResourceOption) (*BotConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToConnectionSettingPropertiesPtrOutput().ApplyT(func(v *ConnectionSettingProperties) *ConnectionSettingProperties { return v.Defaults() }).(ConnectionSettingPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:botservice:BotConnection"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20171201:BotConnection"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20180712:BotConnection"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20200602:BotConnection"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20210501preview:BotConnection"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20220615preview:BotConnection"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20220915:BotConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource BotConnection
	err := ctx.RegisterResource("azure-native:botservice/v20210301:BotConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBotConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotConnectionState, opts ...pulumi.ResourceOption) (*BotConnection, error) {
	var resource BotConnection
	err := ctx.ReadResource("azure-native:botservice/v20210301:BotConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type botConnectionState struct {
}

type BotConnectionState struct {
}

func (BotConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*botConnectionState)(nil)).Elem()
}

type botConnectionArgs struct {
	ConnectionName    *string                      `pulumi:"connectionName"`
	Kind              *string                      `pulumi:"kind"`
	Location          *string                      `pulumi:"location"`
	Properties        *ConnectionSettingProperties `pulumi:"properties"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	ResourceName      string                       `pulumi:"resourceName"`
	Sku               *Sku                         `pulumi:"sku"`
	Tags              map[string]string            `pulumi:"tags"`
}


type BotConnectionArgs struct {
	ConnectionName    pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        ConnectionSettingPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
}

func (BotConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*botConnectionArgs)(nil)).Elem()
}

type BotConnectionInput interface {
	pulumi.Input

	ToBotConnectionOutput() BotConnectionOutput
	ToBotConnectionOutputWithContext(ctx context.Context) BotConnectionOutput
}

func (*BotConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**BotConnection)(nil)).Elem()
}

func (i *BotConnection) ToBotConnectionOutput() BotConnectionOutput {
	return i.ToBotConnectionOutputWithContext(context.Background())
}

func (i *BotConnection) ToBotConnectionOutputWithContext(ctx context.Context) BotConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotConnectionOutput)
}

type BotConnectionOutput struct{ *pulumi.OutputState }

func (BotConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BotConnection)(nil)).Elem()
}

func (o BotConnectionOutput) ToBotConnectionOutput() BotConnectionOutput {
	return o
}

func (o BotConnectionOutput) ToBotConnectionOutputWithContext(ctx context.Context) BotConnectionOutput {
	return o
}

func (o BotConnectionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotConnection) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o BotConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o BotConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o BotConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BotConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BotConnectionOutput) Properties() ConnectionSettingPropertiesResponseOutput {
	return o.ApplyT(func(v *BotConnection) ConnectionSettingPropertiesResponseOutput { return v.Properties }).(ConnectionSettingPropertiesResponseOutput)
}

func (o BotConnectionOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *BotConnection) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o BotConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BotConnection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BotConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BotConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o BotConnectionOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BotConnection) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(BotConnectionOutput{})
}
