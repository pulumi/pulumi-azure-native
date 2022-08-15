


package delegatednetwork

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DelegatedSubnetServiceDetails struct {
	pulumi.CustomResourceState

	ControllerDetails ControllerDetailsResponsePtrOutput `pulumi:"controllerDetails"`
	Location          pulumi.StringPtrOutput             `pulumi:"location"`
	Name              pulumi.StringOutput                `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                `pulumi:"provisioningState"`
	ResourceGuid      pulumi.StringOutput                `pulumi:"resourceGuid"`
	SubnetDetails     SubnetDetailsResponsePtrOutput     `pulumi:"subnetDetails"`
	Tags              pulumi.StringMapOutput             `pulumi:"tags"`
	Type              pulumi.StringOutput                `pulumi:"type"`
}


func NewDelegatedSubnetServiceDetails(ctx *pulumi.Context,
	name string, args *DelegatedSubnetServiceDetailsArgs, opts ...pulumi.ResourceOption) (*DelegatedSubnetServiceDetails, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:delegatednetwork/v20200808preview:DelegatedSubnetServiceDetails"),
		},
		{
			Type: pulumi.String("azure-native:delegatednetwork/v20210315:DelegatedSubnetServiceDetails"),
		},
	})
	opts = append(opts, aliases)
	var resource DelegatedSubnetServiceDetails
	err := ctx.RegisterResource("azure-native:delegatednetwork:DelegatedSubnetServiceDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDelegatedSubnetServiceDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DelegatedSubnetServiceDetailsState, opts ...pulumi.ResourceOption) (*DelegatedSubnetServiceDetails, error) {
	var resource DelegatedSubnetServiceDetails
	err := ctx.ReadResource("azure-native:delegatednetwork:DelegatedSubnetServiceDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type delegatedSubnetServiceDetailsState struct {
}

type DelegatedSubnetServiceDetailsState struct {
}

func (DelegatedSubnetServiceDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*delegatedSubnetServiceDetailsState)(nil)).Elem()
}

type delegatedSubnetServiceDetailsArgs struct {
	ControllerDetails *ControllerDetailsType `pulumi:"controllerDetails"`
	Location          *string                `pulumi:"location"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	ResourceName      *string                `pulumi:"resourceName"`
	SubnetDetails     *SubnetDetails         `pulumi:"subnetDetails"`
	Tags              map[string]string      `pulumi:"tags"`
}


type DelegatedSubnetServiceDetailsArgs struct {
	ControllerDetails ControllerDetailsTypePtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	SubnetDetails     SubnetDetailsPtrInput
	Tags              pulumi.StringMapInput
}

func (DelegatedSubnetServiceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*delegatedSubnetServiceDetailsArgs)(nil)).Elem()
}

type DelegatedSubnetServiceDetailsInput interface {
	pulumi.Input

	ToDelegatedSubnetServiceDetailsOutput() DelegatedSubnetServiceDetailsOutput
	ToDelegatedSubnetServiceDetailsOutputWithContext(ctx context.Context) DelegatedSubnetServiceDetailsOutput
}

func (*DelegatedSubnetServiceDetails) ElementType() reflect.Type {
	return reflect.TypeOf((**DelegatedSubnetServiceDetails)(nil)).Elem()
}

func (i *DelegatedSubnetServiceDetails) ToDelegatedSubnetServiceDetailsOutput() DelegatedSubnetServiceDetailsOutput {
	return i.ToDelegatedSubnetServiceDetailsOutputWithContext(context.Background())
}

func (i *DelegatedSubnetServiceDetails) ToDelegatedSubnetServiceDetailsOutputWithContext(ctx context.Context) DelegatedSubnetServiceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelegatedSubnetServiceDetailsOutput)
}

type DelegatedSubnetServiceDetailsOutput struct{ *pulumi.OutputState }

func (DelegatedSubnetServiceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DelegatedSubnetServiceDetails)(nil)).Elem()
}

func (o DelegatedSubnetServiceDetailsOutput) ToDelegatedSubnetServiceDetailsOutput() DelegatedSubnetServiceDetailsOutput {
	return o
}

func (o DelegatedSubnetServiceDetailsOutput) ToDelegatedSubnetServiceDetailsOutputWithContext(ctx context.Context) DelegatedSubnetServiceDetailsOutput {
	return o
}

func (o DelegatedSubnetServiceDetailsOutput) ControllerDetails() ControllerDetailsResponsePtrOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) ControllerDetailsResponsePtrOutput { return v.ControllerDetails }).(ControllerDetailsResponsePtrOutput)
}

func (o DelegatedSubnetServiceDetailsOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DelegatedSubnetServiceDetailsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DelegatedSubnetServiceDetailsOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DelegatedSubnetServiceDetailsOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o DelegatedSubnetServiceDetailsOutput) SubnetDetails() SubnetDetailsResponsePtrOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) SubnetDetailsResponsePtrOutput { return v.SubnetDetails }).(SubnetDetailsResponsePtrOutput)
}

func (o DelegatedSubnetServiceDetailsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DelegatedSubnetServiceDetailsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DelegatedSubnetServiceDetailsOutput{})
}
