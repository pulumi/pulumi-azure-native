


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnergyService struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                   `pulumi:"location"`
	Name       pulumi.StringOutput                   `pulumi:"name"`
	Properties EnergyServicePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput              `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                `pulumi:"tags"`
	Type       pulumi.StringOutput                   `pulumi:"type"`
}


func NewEnergyService(ctx *pulumi.Context,
	name string, args *EnergyServiceArgs, opts ...pulumi.ResourceOption) (*EnergyService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:openenergyplatform:EnergyService"),
		},
		{
			Type: pulumi.String("azure-native:openenergyplatform/v20220404preview:EnergyService"),
		},
	})
	opts = append(opts, aliases)
	var resource EnergyService
	err := ctx.RegisterResource("azure-native:openenergyplatform/v20210601preview:EnergyService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnergyService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnergyServiceState, opts ...pulumi.ResourceOption) (*EnergyService, error) {
	var resource EnergyService
	err := ctx.ReadResource("azure-native:openenergyplatform/v20210601preview:EnergyService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type energyServiceState struct {
}

type EnergyServiceState struct {
}

func (EnergyServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*energyServiceState)(nil)).Elem()
}

type energyServiceArgs struct {
	Location          *string                  `pulumi:"location"`
	Properties        *EnergyServiceProperties `pulumi:"properties"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	ResourceName      *string                  `pulumi:"resourceName"`
	Tags              map[string]string        `pulumi:"tags"`
}


type EnergyServiceArgs struct {
	Location          pulumi.StringPtrInput
	Properties        EnergyServicePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (EnergyServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*energyServiceArgs)(nil)).Elem()
}

type EnergyServiceInput interface {
	pulumi.Input

	ToEnergyServiceOutput() EnergyServiceOutput
	ToEnergyServiceOutputWithContext(ctx context.Context) EnergyServiceOutput
}

func (*EnergyService) ElementType() reflect.Type {
	return reflect.TypeOf((**EnergyService)(nil)).Elem()
}

func (i *EnergyService) ToEnergyServiceOutput() EnergyServiceOutput {
	return i.ToEnergyServiceOutputWithContext(context.Background())
}

func (i *EnergyService) ToEnergyServiceOutputWithContext(ctx context.Context) EnergyServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnergyServiceOutput)
}

type EnergyServiceOutput struct{ *pulumi.OutputState }

func (EnergyServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnergyService)(nil)).Elem()
}

func (o EnergyServiceOutput) ToEnergyServiceOutput() EnergyServiceOutput {
	return o
}

func (o EnergyServiceOutput) ToEnergyServiceOutputWithContext(ctx context.Context) EnergyServiceOutput {
	return o
}

func (o EnergyServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EnergyService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o EnergyServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnergyService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EnergyServiceOutput) Properties() EnergyServicePropertiesResponseOutput {
	return o.ApplyT(func(v *EnergyService) EnergyServicePropertiesResponseOutput { return v.Properties }).(EnergyServicePropertiesResponseOutput)
}

func (o EnergyServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnergyService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EnergyServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnergyService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o EnergyServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnergyService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnergyServiceOutput{})
}
