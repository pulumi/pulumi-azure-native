


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Network struct {
	pulumi.CustomResourceState

	AddressPrefix     pulumi.StringOutput            `pulumi:"addressPrefix"`
	Description       pulumi.StringPtrOutput         `pulumi:"description"`
	IngressConfig     IngressConfigResponsePtrOutput `pulumi:"ingressConfig"`
	Location          pulumi.StringOutput            `pulumi:"location"`
	Name              pulumi.StringOutput            `pulumi:"name"`
	ProvisioningState pulumi.StringOutput            `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput         `pulumi:"tags"`
	Type              pulumi.StringOutput            `pulumi:"type"`
}


func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressPrefix == nil {
		return nil, errors.New("invalid value for required argument 'AddressPrefix'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabricmesh:Network"),
		},
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180901preview:Network"),
		},
	})
	opts = append(opts, aliases)
	var resource Network
	err := ctx.RegisterResource("azure-native:servicefabricmesh/v20180701preview:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("azure-native:servicefabricmesh/v20180701preview:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkState struct {
}

type NetworkState struct {
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	AddressPrefix     string            `pulumi:"addressPrefix"`
	Description       *string           `pulumi:"description"`
	IngressConfig     *IngressConfig    `pulumi:"ingressConfig"`
	Location          *string           `pulumi:"location"`
	NetworkName       *string           `pulumi:"networkName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type NetworkArgs struct {
	AddressPrefix     pulumi.StringInput
	Description       pulumi.StringPtrInput
	IngressConfig     IngressConfigPtrInput
	Location          pulumi.StringPtrInput
	NetworkName       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkArgs)(nil)).Elem()
}

type NetworkInput interface {
	pulumi.Input

	ToNetworkOutput() NetworkOutput
	ToNetworkOutputWithContext(ctx context.Context) NetworkOutput
}

func (*Network) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (i *Network) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i *Network) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
}

type NetworkOutput struct{ *pulumi.OutputState }

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkOutput{})
}
