


package v20220101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectivityConfiguration struct {
	pulumi.CustomResourceState

	AppliesToGroups       ConnectivityGroupItemResponseArrayOutput `pulumi:"appliesToGroups"`
	ConnectivityTopology  pulumi.StringOutput                      `pulumi:"connectivityTopology"`
	DeleteExistingPeering pulumi.StringPtrOutput                   `pulumi:"deleteExistingPeering"`
	Description           pulumi.StringPtrOutput                   `pulumi:"description"`
	Etag                  pulumi.StringOutput                      `pulumi:"etag"`
	Hubs                  HubResponseArrayOutput                   `pulumi:"hubs"`
	IsGlobal              pulumi.StringPtrOutput                   `pulumi:"isGlobal"`
	Name                  pulumi.StringOutput                      `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput                      `pulumi:"provisioningState"`
	SystemData            SystemDataResponseOutput                 `pulumi:"systemData"`
	Type                  pulumi.StringOutput                      `pulumi:"type"`
}


func NewConnectivityConfiguration(ctx *pulumi.Context,
	name string, args *ConnectivityConfigurationArgs, opts ...pulumi.ResourceOption) (*ConnectivityConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppliesToGroups == nil {
		return nil, errors.New("invalid value for required argument 'AppliesToGroups'")
	}
	if args.ConnectivityTopology == nil {
		return nil, errors.New("invalid value for required argument 'ConnectivityTopology'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ConnectivityConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:ConnectivityConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:ConnectivityConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:ConnectivityConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:ConnectivityConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ConnectivityConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectivityConfiguration
	err := ctx.RegisterResource("azure-native:network/v20220101:ConnectivityConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectivityConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectivityConfigurationState, opts ...pulumi.ResourceOption) (*ConnectivityConfiguration, error) {
	var resource ConnectivityConfiguration
	err := ctx.ReadResource("azure-native:network/v20220101:ConnectivityConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectivityConfigurationState struct {
}

type ConnectivityConfigurationState struct {
}

func (ConnectivityConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectivityConfigurationState)(nil)).Elem()
}

type connectivityConfigurationArgs struct {
	AppliesToGroups       []ConnectivityGroupItem `pulumi:"appliesToGroups"`
	ConfigurationName     *string                 `pulumi:"configurationName"`
	ConnectivityTopology  string                  `pulumi:"connectivityTopology"`
	DeleteExistingPeering *string                 `pulumi:"deleteExistingPeering"`
	Description           *string                 `pulumi:"description"`
	Hubs                  []Hub                   `pulumi:"hubs"`
	IsGlobal              *string                 `pulumi:"isGlobal"`
	NetworkManagerName    string                  `pulumi:"networkManagerName"`
	ResourceGroupName     string                  `pulumi:"resourceGroupName"`
}


type ConnectivityConfigurationArgs struct {
	AppliesToGroups       ConnectivityGroupItemArrayInput
	ConfigurationName     pulumi.StringPtrInput
	ConnectivityTopology  pulumi.StringInput
	DeleteExistingPeering pulumi.StringPtrInput
	Description           pulumi.StringPtrInput
	Hubs                  HubArrayInput
	IsGlobal              pulumi.StringPtrInput
	NetworkManagerName    pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
}

func (ConnectivityConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectivityConfigurationArgs)(nil)).Elem()
}

type ConnectivityConfigurationInput interface {
	pulumi.Input

	ToConnectivityConfigurationOutput() ConnectivityConfigurationOutput
	ToConnectivityConfigurationOutputWithContext(ctx context.Context) ConnectivityConfigurationOutput
}

func (*ConnectivityConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectivityConfiguration)(nil)).Elem()
}

func (i *ConnectivityConfiguration) ToConnectivityConfigurationOutput() ConnectivityConfigurationOutput {
	return i.ToConnectivityConfigurationOutputWithContext(context.Background())
}

func (i *ConnectivityConfiguration) ToConnectivityConfigurationOutputWithContext(ctx context.Context) ConnectivityConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectivityConfigurationOutput)
}

type ConnectivityConfigurationOutput struct{ *pulumi.OutputState }

func (ConnectivityConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectivityConfiguration)(nil)).Elem()
}

func (o ConnectivityConfigurationOutput) ToConnectivityConfigurationOutput() ConnectivityConfigurationOutput {
	return o
}

func (o ConnectivityConfigurationOutput) ToConnectivityConfigurationOutputWithContext(ctx context.Context) ConnectivityConfigurationOutput {
	return o
}

func (o ConnectivityConfigurationOutput) AppliesToGroups() ConnectivityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) ConnectivityGroupItemResponseArrayOutput { return v.AppliesToGroups }).(ConnectivityGroupItemResponseArrayOutput)
}

func (o ConnectivityConfigurationOutput) ConnectivityTopology() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringOutput { return v.ConnectivityTopology }).(pulumi.StringOutput)
}

func (o ConnectivityConfigurationOutput) DeleteExistingPeering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringPtrOutput { return v.DeleteExistingPeering }).(pulumi.StringPtrOutput)
}

func (o ConnectivityConfigurationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectivityConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ConnectivityConfigurationOutput) Hubs() HubResponseArrayOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) HubResponseArrayOutput { return v.Hubs }).(HubResponseArrayOutput)
}

func (o ConnectivityConfigurationOutput) IsGlobal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringPtrOutput { return v.IsGlobal }).(pulumi.StringPtrOutput)
}

func (o ConnectivityConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectivityConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConnectivityConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConnectivityConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectivityConfigurationOutput{})
}
