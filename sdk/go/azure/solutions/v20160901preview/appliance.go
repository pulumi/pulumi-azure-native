


package v20160901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Appliance struct {
	pulumi.CustomResourceState

	ApplianceDefinitionId  pulumi.StringPtrOutput    `pulumi:"applianceDefinitionId"`
	Identity               IdentityResponsePtrOutput `pulumi:"identity"`
	Kind                   pulumi.StringPtrOutput    `pulumi:"kind"`
	Location               pulumi.StringPtrOutput    `pulumi:"location"`
	ManagedBy              pulumi.StringPtrOutput    `pulumi:"managedBy"`
	ManagedResourceGroupId pulumi.StringOutput       `pulumi:"managedResourceGroupId"`
	Name                   pulumi.StringOutput       `pulumi:"name"`
	Outputs                pulumi.AnyOutput          `pulumi:"outputs"`
	Parameters             pulumi.AnyOutput          `pulumi:"parameters"`
	Plan                   PlanResponsePtrOutput     `pulumi:"plan"`
	ProvisioningState      pulumi.StringOutput       `pulumi:"provisioningState"`
	Sku                    SkuResponsePtrOutput      `pulumi:"sku"`
	Tags                   pulumi.StringMapOutput    `pulumi:"tags"`
	Type                   pulumi.StringOutput       `pulumi:"type"`
	UiDefinitionUri        pulumi.StringPtrOutput    `pulumi:"uiDefinitionUri"`
}


func NewAppliance(ctx *pulumi.Context,
	name string, args *ApplianceArgs, opts ...pulumi.ResourceOption) (*Appliance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagedResourceGroupId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:solutions/v20160901preview:Appliance"),
		},
		{
			Type: pulumi.String("azure-native:solutions:Appliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions:Appliance"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20170901:Appliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions/v20170901:Appliance"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180601:Appliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions/v20180601:Appliance"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20190701:Appliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions/v20190701:Appliance"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20200821preview:Appliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions/v20200821preview:Appliance"),
		},
	})
	opts = append(opts, aliases)
	var resource Appliance
	err := ctx.RegisterResource("azure-native:solutions/v20160901preview:Appliance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppliance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplianceState, opts ...pulumi.ResourceOption) (*Appliance, error) {
	var resource Appliance
	err := ctx.ReadResource("azure-native:solutions/v20160901preview:Appliance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applianceState struct {
}

type ApplianceState struct {
}

func (ApplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*applianceState)(nil)).Elem()
}

type applianceArgs struct {
	ApplianceDefinitionId  *string           `pulumi:"applianceDefinitionId"`
	ApplianceName          *string           `pulumi:"applianceName"`
	Identity               *Identity         `pulumi:"identity"`
	Kind                   *string           `pulumi:"kind"`
	Location               *string           `pulumi:"location"`
	ManagedBy              *string           `pulumi:"managedBy"`
	ManagedResourceGroupId string            `pulumi:"managedResourceGroupId"`
	Parameters             interface{}       `pulumi:"parameters"`
	Plan                   *Plan             `pulumi:"plan"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Sku                    *Sku              `pulumi:"sku"`
	Tags                   map[string]string `pulumi:"tags"`
	UiDefinitionUri        *string           `pulumi:"uiDefinitionUri"`
}


type ApplianceArgs struct {
	ApplianceDefinitionId  pulumi.StringPtrInput
	ApplianceName          pulumi.StringPtrInput
	Identity               IdentityPtrInput
	Kind                   pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	ManagedBy              pulumi.StringPtrInput
	ManagedResourceGroupId pulumi.StringInput
	Parameters             pulumi.Input
	Plan                   PlanPtrInput
	ResourceGroupName      pulumi.StringInput
	Sku                    SkuPtrInput
	Tags                   pulumi.StringMapInput
	UiDefinitionUri        pulumi.StringPtrInput
}

func (ApplianceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applianceArgs)(nil)).Elem()
}

type ApplianceInput interface {
	pulumi.Input

	ToApplianceOutput() ApplianceOutput
	ToApplianceOutputWithContext(ctx context.Context) ApplianceOutput
}

func (*Appliance) ElementType() reflect.Type {
	return reflect.TypeOf((*Appliance)(nil))
}

func (i *Appliance) ToApplianceOutput() ApplianceOutput {
	return i.ToApplianceOutputWithContext(context.Background())
}

func (i *Appliance) ToApplianceOutputWithContext(ctx context.Context) ApplianceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceOutput)
}

type ApplianceOutput struct{ *pulumi.OutputState }

func (ApplianceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Appliance)(nil))
}

func (o ApplianceOutput) ToApplianceOutput() ApplianceOutput {
	return o
}

func (o ApplianceOutput) ToApplianceOutputWithContext(ctx context.Context) ApplianceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplianceOutput{})
}
