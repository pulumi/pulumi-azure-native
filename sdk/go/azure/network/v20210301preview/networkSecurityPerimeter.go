


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkSecurityPerimeter struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName       pulumi.StringPtrOutput `pulumi:"displayName"`
	Etag              pulumi.StringOutput    `pulumi:"etag"`
	Location          pulumi.StringPtrOutput `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
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
			Type: pulumi.String("azure-nextgen:network/v20210301preview:NetworkSecurityPerimeter"),
		},
		{
			Type: pulumi.String("azure-native:network:NetworkSecurityPerimeter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:NetworkSecurityPerimeter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:NetworkSecurityPerimeter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201preview:NetworkSecurityPerimeter"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkSecurityPerimeter
	err := ctx.RegisterResource("azure-native:network/v20210301preview:NetworkSecurityPerimeter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkSecurityPerimeter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityPerimeterState, opts ...pulumi.ResourceOption) (*NetworkSecurityPerimeter, error) {
	var resource NetworkSecurityPerimeter
	err := ctx.ReadResource("azure-native:network/v20210301preview:NetworkSecurityPerimeter", name, id, state, &resource, opts...)
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
	Description                  *string           `pulumi:"description"`
	DisplayName                  *string           `pulumi:"displayName"`
	Id                           *string           `pulumi:"id"`
	Location                     *string           `pulumi:"location"`
	Name                         *string           `pulumi:"name"`
	NetworkSecurityPerimeterName *string           `pulumi:"networkSecurityPerimeterName"`
	ResourceGroupName            string            `pulumi:"resourceGroupName"`
	Tags                         map[string]string `pulumi:"tags"`
}


type NetworkSecurityPerimeterArgs struct {
	Description                  pulumi.StringPtrInput
	DisplayName                  pulumi.StringPtrInput
	Id                           pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	NetworkSecurityPerimeterName pulumi.StringPtrInput
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
	return reflect.TypeOf((*NetworkSecurityPerimeter)(nil))
}

func (i *NetworkSecurityPerimeter) ToNetworkSecurityPerimeterOutput() NetworkSecurityPerimeterOutput {
	return i.ToNetworkSecurityPerimeterOutputWithContext(context.Background())
}

func (i *NetworkSecurityPerimeter) ToNetworkSecurityPerimeterOutputWithContext(ctx context.Context) NetworkSecurityPerimeterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityPerimeterOutput)
}

type NetworkSecurityPerimeterOutput struct{ *pulumi.OutputState }

func (NetworkSecurityPerimeterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityPerimeter)(nil))
}

func (o NetworkSecurityPerimeterOutput) ToNetworkSecurityPerimeterOutput() NetworkSecurityPerimeterOutput {
	return o
}

func (o NetworkSecurityPerimeterOutput) ToNetworkSecurityPerimeterOutputWithContext(ctx context.Context) NetworkSecurityPerimeterOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityPerimeterInput)(nil)).Elem(), &NetworkSecurityPerimeter{})
	pulumi.RegisterOutputType(NetworkSecurityPerimeterOutput{})
}
