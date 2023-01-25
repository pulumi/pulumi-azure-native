


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssociationsInterface struct {
	pulumi.CustomResourceState

	AssociationType   pulumi.StringOutput                `pulumi:"associationType"`
	Location          pulumi.StringOutput                `pulumi:"location"`
	Name              pulumi.StringOutput                `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                `pulumi:"provisioningState"`
	Subnet            AssociationSubnetResponsePtrOutput `pulumi:"subnet"`
	SystemData        SystemDataResponseOutput           `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput             `pulumi:"tags"`
	Type              pulumi.StringOutput                `pulumi:"type"`
}


func NewAssociationsInterface(ctx *pulumi.Context,
	name string, args *AssociationsInterfaceArgs, opts ...pulumi.ResourceOption) (*AssociationsInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssociationType == nil {
		return nil, errors.New("invalid value for required argument 'AssociationType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TrafficControllerName == nil {
		return nil, errors.New("invalid value for required argument 'TrafficControllerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicenetworking:AssociationsInterface"),
		},
	})
	opts = append(opts, aliases)
	var resource AssociationsInterface
	err := ctx.RegisterResource("azure-native:servicenetworking/v20221001preview:AssociationsInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAssociationsInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssociationsInterfaceState, opts ...pulumi.ResourceOption) (*AssociationsInterface, error) {
	var resource AssociationsInterface
	err := ctx.ReadResource("azure-native:servicenetworking/v20221001preview:AssociationsInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type associationsInterfaceState struct {
}

type AssociationsInterfaceState struct {
}

func (AssociationsInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*associationsInterfaceState)(nil)).Elem()
}

type associationsInterfaceArgs struct {
	AssociationName       *string            `pulumi:"associationName"`
	AssociationType       AssociationType    `pulumi:"associationType"`
	Location              *string            `pulumi:"location"`
	ResourceGroupName     string             `pulumi:"resourceGroupName"`
	Subnet                *AssociationSubnet `pulumi:"subnet"`
	Tags                  map[string]string  `pulumi:"tags"`
	TrafficControllerName string             `pulumi:"trafficControllerName"`
}


type AssociationsInterfaceArgs struct {
	AssociationName       pulumi.StringPtrInput
	AssociationType       AssociationTypeInput
	Location              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Subnet                AssociationSubnetPtrInput
	Tags                  pulumi.StringMapInput
	TrafficControllerName pulumi.StringInput
}

func (AssociationsInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*associationsInterfaceArgs)(nil)).Elem()
}

type AssociationsInterfaceInput interface {
	pulumi.Input

	ToAssociationsInterfaceOutput() AssociationsInterfaceOutput
	ToAssociationsInterfaceOutputWithContext(ctx context.Context) AssociationsInterfaceOutput
}

func (*AssociationsInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationsInterface)(nil)).Elem()
}

func (i *AssociationsInterface) ToAssociationsInterfaceOutput() AssociationsInterfaceOutput {
	return i.ToAssociationsInterfaceOutputWithContext(context.Background())
}

func (i *AssociationsInterface) ToAssociationsInterfaceOutputWithContext(ctx context.Context) AssociationsInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationsInterfaceOutput)
}

type AssociationsInterfaceOutput struct{ *pulumi.OutputState }

func (AssociationsInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationsInterface)(nil)).Elem()
}

func (o AssociationsInterfaceOutput) ToAssociationsInterfaceOutput() AssociationsInterfaceOutput {
	return o
}

func (o AssociationsInterfaceOutput) ToAssociationsInterfaceOutputWithContext(ctx context.Context) AssociationsInterfaceOutput {
	return o
}

func (o AssociationsInterfaceOutput) AssociationType() pulumi.StringOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringOutput { return v.AssociationType }).(pulumi.StringOutput)
}

func (o AssociationsInterfaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AssociationsInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AssociationsInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AssociationsInterfaceOutput) Subnet() AssociationSubnetResponsePtrOutput {
	return o.ApplyT(func(v *AssociationsInterface) AssociationSubnetResponsePtrOutput { return v.Subnet }).(AssociationSubnetResponsePtrOutput)
}

func (o AssociationsInterfaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AssociationsInterface) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AssociationsInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AssociationsInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssociationsInterfaceOutput{})
}
