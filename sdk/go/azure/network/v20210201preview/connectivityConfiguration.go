


package v20210201preview

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
	DisplayName           pulumi.StringPtrOutput                   `pulumi:"displayName"`
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
	})
	opts = append(opts, aliases)
	var resource ConnectivityConfiguration
	err := ctx.RegisterResource("azure-native:network/v20210201preview:ConnectivityConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectivityConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectivityConfigurationState, opts ...pulumi.ResourceOption) (*ConnectivityConfiguration, error) {
	var resource ConnectivityConfiguration
	err := ctx.ReadResource("azure-native:network/v20210201preview:ConnectivityConfiguration", name, id, state, &resource, opts...)
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
	DisplayName           *string                 `pulumi:"displayName"`
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
	DisplayName           pulumi.StringPtrInput
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
	return reflect.TypeOf((*ConnectivityConfiguration)(nil))
}

func (i *ConnectivityConfiguration) ToConnectivityConfigurationOutput() ConnectivityConfigurationOutput {
	return i.ToConnectivityConfigurationOutputWithContext(context.Background())
}

func (i *ConnectivityConfiguration) ToConnectivityConfigurationOutputWithContext(ctx context.Context) ConnectivityConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectivityConfigurationOutput)
}

type ConnectivityConfigurationOutput struct{ *pulumi.OutputState }

func (ConnectivityConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityConfiguration)(nil))
}

func (o ConnectivityConfigurationOutput) ToConnectivityConfigurationOutput() ConnectivityConfigurationOutput {
	return o
}

func (o ConnectivityConfigurationOutput) ToConnectivityConfigurationOutputWithContext(ctx context.Context) ConnectivityConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectivityConfigurationOutput{})
}
