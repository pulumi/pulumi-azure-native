


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectedRegistry struct {
	pulumi.CustomResourceState

	Activation        ActivationPropertiesResponseOutput        `pulumi:"activation"`
	ClientTokenIds    pulumi.StringArrayOutput                  `pulumi:"clientTokenIds"`
	ConnectionState   pulumi.StringOutput                       `pulumi:"connectionState"`
	LastActivityTime  pulumi.StringOutput                       `pulumi:"lastActivityTime"`
	Logging           LoggingPropertiesResponsePtrOutput        `pulumi:"logging"`
	LoginServer       LoginServerPropertiesResponsePtrOutput    `pulumi:"loginServer"`
	Mode              pulumi.StringOutput                       `pulumi:"mode"`
	Name              pulumi.StringOutput                       `pulumi:"name"`
	NotificationsList pulumi.StringArrayOutput                  `pulumi:"notificationsList"`
	Parent            ParentPropertiesResponseOutput            `pulumi:"parent"`
	ProvisioningState pulumi.StringOutput                       `pulumi:"provisioningState"`
	StatusDetails     StatusDetailPropertiesResponseArrayOutput `pulumi:"statusDetails"`
	SystemData        SystemDataResponseOutput                  `pulumi:"systemData"`
	Type              pulumi.StringOutput                       `pulumi:"type"`
	Version           pulumi.StringOutput                       `pulumi:"version"`
}


func NewConnectedRegistry(ctx *pulumi.Context,
	name string, args *ConnectedRegistryArgs, opts ...pulumi.ResourceOption) (*ConnectedRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Logging != nil {
		args.Logging = args.Logging.ToLoggingPropertiesPtrOutput().ApplyT(func(v *LoggingProperties) *LoggingProperties { return v.Defaults() }).(LoggingPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20220201preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230101preview:ConnectedRegistry"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectedRegistry
	err := ctx.RegisterResource("azure-native:containerregistry/v20211201preview:ConnectedRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectedRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedRegistryState, opts ...pulumi.ResourceOption) (*ConnectedRegistry, error) {
	var resource ConnectedRegistry
	err := ctx.ReadResource("azure-native:containerregistry/v20211201preview:ConnectedRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectedRegistryState struct {
}

type ConnectedRegistryState struct {
}

func (ConnectedRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedRegistryState)(nil)).Elem()
}

type connectedRegistryArgs struct {
	ClientTokenIds        []string           `pulumi:"clientTokenIds"`
	ConnectedRegistryName *string            `pulumi:"connectedRegistryName"`
	Logging               *LoggingProperties `pulumi:"logging"`
	Mode                  string             `pulumi:"mode"`
	NotificationsList     []string           `pulumi:"notificationsList"`
	Parent                ParentProperties   `pulumi:"parent"`
	RegistryName          string             `pulumi:"registryName"`
	ResourceGroupName     string             `pulumi:"resourceGroupName"`
}


type ConnectedRegistryArgs struct {
	ClientTokenIds        pulumi.StringArrayInput
	ConnectedRegistryName pulumi.StringPtrInput
	Logging               LoggingPropertiesPtrInput
	Mode                  pulumi.StringInput
	NotificationsList     pulumi.StringArrayInput
	Parent                ParentPropertiesInput
	RegistryName          pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
}

func (ConnectedRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedRegistryArgs)(nil)).Elem()
}

type ConnectedRegistryInput interface {
	pulumi.Input

	ToConnectedRegistryOutput() ConnectedRegistryOutput
	ToConnectedRegistryOutputWithContext(ctx context.Context) ConnectedRegistryOutput
}

func (*ConnectedRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedRegistry)(nil)).Elem()
}

func (i *ConnectedRegistry) ToConnectedRegistryOutput() ConnectedRegistryOutput {
	return i.ToConnectedRegistryOutputWithContext(context.Background())
}

func (i *ConnectedRegistry) ToConnectedRegistryOutputWithContext(ctx context.Context) ConnectedRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedRegistryOutput)
}

type ConnectedRegistryOutput struct{ *pulumi.OutputState }

func (ConnectedRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedRegistry)(nil)).Elem()
}

func (o ConnectedRegistryOutput) ToConnectedRegistryOutput() ConnectedRegistryOutput {
	return o
}

func (o ConnectedRegistryOutput) ToConnectedRegistryOutputWithContext(ctx context.Context) ConnectedRegistryOutput {
	return o
}

func (o ConnectedRegistryOutput) Activation() ActivationPropertiesResponseOutput {
	return o.ApplyT(func(v *ConnectedRegistry) ActivationPropertiesResponseOutput { return v.Activation }).(ActivationPropertiesResponseOutput)
}

func (o ConnectedRegistryOutput) ClientTokenIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringArrayOutput { return v.ClientTokenIds }).(pulumi.StringArrayOutput)
}

func (o ConnectedRegistryOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.ConnectionState }).(pulumi.StringOutput)
}

func (o ConnectedRegistryOutput) LastActivityTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.LastActivityTime }).(pulumi.StringOutput)
}

func (o ConnectedRegistryOutput) Logging() LoggingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ConnectedRegistry) LoggingPropertiesResponsePtrOutput { return v.Logging }).(LoggingPropertiesResponsePtrOutput)
}

func (o ConnectedRegistryOutput) LoginServer() LoginServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ConnectedRegistry) LoginServerPropertiesResponsePtrOutput { return v.LoginServer }).(LoginServerPropertiesResponsePtrOutput)
}

func (o ConnectedRegistryOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

func (o ConnectedRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectedRegistryOutput) NotificationsList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringArrayOutput { return v.NotificationsList }).(pulumi.StringArrayOutput)
}

func (o ConnectedRegistryOutput) Parent() ParentPropertiesResponseOutput {
	return o.ApplyT(func(v *ConnectedRegistry) ParentPropertiesResponseOutput { return v.Parent }).(ParentPropertiesResponseOutput)
}

func (o ConnectedRegistryOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConnectedRegistryOutput) StatusDetails() StatusDetailPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *ConnectedRegistry) StatusDetailPropertiesResponseArrayOutput { return v.StatusDetails }).(StatusDetailPropertiesResponseArrayOutput)
}

func (o ConnectedRegistryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedRegistry) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConnectedRegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ConnectedRegistryOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedRegistryOutput{})
}
