


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Network struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties NetworkResourcePropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabricmesh:Network"),
		},
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180701preview:Network"),
		},
	})
	opts = append(opts, aliases)
	var resource Network
	err := ctx.RegisterResource("azure-native:servicefabricmesh/v20180901preview:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("azure-native:servicefabricmesh/v20180901preview:Network", name, id, state, &resource, opts...)
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
	Location            *string                   `pulumi:"location"`
	NetworkResourceName *string                   `pulumi:"networkResourceName"`
	Properties          NetworkResourceProperties `pulumi:"properties"`
	ResourceGroupName   string                    `pulumi:"resourceGroupName"`
	Tags                map[string]string         `pulumi:"tags"`
}


type NetworkArgs struct {
	Location            pulumi.StringPtrInput
	NetworkResourceName pulumi.StringPtrInput
	Properties          NetworkResourcePropertiesInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
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

func (o NetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkOutput) Properties() NetworkResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *Network) NetworkResourcePropertiesResponseOutput { return v.Properties }).(NetworkResourcePropertiesResponseOutput)
}

func (o NetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Network) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkOutput{})
}
