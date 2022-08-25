


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkSecurityPerimeter struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput    `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	PerimeterGuid     pulumi.StringPtrOutput `pulumi:"perimeterGuid"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewNetworkSecurityPerimeter(ctx *pulumi.Context,
	name string, args *NetworkSecurityPerimeterArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityPerimeter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210201preview:NetworkSecurityPerimeter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301preview:NetworkSecurityPerimeter"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkSecurityPerimeter
	err := ctx.RegisterResource("azure-native:network:NetworkSecurityPerimeter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkSecurityPerimeter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityPerimeterState, opts ...pulumi.ResourceOption) (*NetworkSecurityPerimeter, error) {
	var resource NetworkSecurityPerimeter
	err := ctx.ReadResource("azure-native:network:NetworkSecurityPerimeter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkSecurityPerimeterState struct {
}

type NetworkSecurityPerimeterState struct {
}

func (NetworkSecurityPerimeterState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityPerimeterState)(nil)).Elem()
}

type networkSecurityPerimeterArgs struct {
	Id                           *string           `pulumi:"id"`
	Location                     *string           `pulumi:"location"`
	Name                         *string           `pulumi:"name"`
	NetworkSecurityPerimeterName *string           `pulumi:"networkSecurityPerimeterName"`
	PerimeterGuid                *string           `pulumi:"perimeterGuid"`
	ResourceGroupName            string            `pulumi:"resourceGroupName"`
	Tags                         map[string]string `pulumi:"tags"`
}


type NetworkSecurityPerimeterArgs struct {
	Id                           pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	NetworkSecurityPerimeterName pulumi.StringPtrInput
	PerimeterGuid                pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
}

func (NetworkSecurityPerimeterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityPerimeterArgs)(nil)).Elem()
}

type NetworkSecurityPerimeterInput interface {
	pulumi.Input

	ToNetworkSecurityPerimeterOutput() NetworkSecurityPerimeterOutput
	ToNetworkSecurityPerimeterOutputWithContext(ctx context.Context) NetworkSecurityPerimeterOutput
}

func (*NetworkSecurityPerimeter) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityPerimeter)(nil)).Elem()
}

func (i *NetworkSecurityPerimeter) ToNetworkSecurityPerimeterOutput() NetworkSecurityPerimeterOutput {
	return i.ToNetworkSecurityPerimeterOutputWithContext(context.Background())
}

func (i *NetworkSecurityPerimeter) ToNetworkSecurityPerimeterOutputWithContext(ctx context.Context) NetworkSecurityPerimeterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityPerimeterOutput)
}

type NetworkSecurityPerimeterOutput struct{ *pulumi.OutputState }

func (NetworkSecurityPerimeterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityPerimeter)(nil)).Elem()
}

func (o NetworkSecurityPerimeterOutput) ToNetworkSecurityPerimeterOutput() NetworkSecurityPerimeterOutput {
	return o
}

func (o NetworkSecurityPerimeterOutput) ToNetworkSecurityPerimeterOutputWithContext(ctx context.Context) NetworkSecurityPerimeterOutput {
	return o
}

func (o NetworkSecurityPerimeterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityPerimeter) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NetworkSecurityPerimeterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityPerimeter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkSecurityPerimeterOutput) PerimeterGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityPerimeter) pulumi.StringPtrOutput { return v.PerimeterGuid }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityPerimeterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityPerimeter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NetworkSecurityPerimeterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityPerimeter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityPerimeterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityPerimeter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkSecurityPerimeterOutput{})
}
