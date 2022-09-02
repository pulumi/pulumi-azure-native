


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DevBoxDefinition struct {
	pulumi.CustomResourceState

	ActiveImageReference        ImageReferenceResponseOutput              `pulumi:"activeImageReference"`
	ImageReference              ImageReferenceResponseOutput              `pulumi:"imageReference"`
	ImageValidationErrorDetails ImageValidationErrorDetailsResponseOutput `pulumi:"imageValidationErrorDetails"`
	ImageValidationStatus       pulumi.StringOutput                       `pulumi:"imageValidationStatus"`
	Location                    pulumi.StringOutput                       `pulumi:"location"`
	Name                        pulumi.StringOutput                       `pulumi:"name"`
	OsStorageType               pulumi.StringOutput                       `pulumi:"osStorageType"`
	ProvisioningState           pulumi.StringOutput                       `pulumi:"provisioningState"`
	Sku                         SkuResponseOutput                         `pulumi:"sku"`
	SystemData                  SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags                        pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                        pulumi.StringOutput                       `pulumi:"type"`
}


func NewDevBoxDefinition(ctx *pulumi.Context,
	name string, args *DevBoxDefinitionArgs, opts ...pulumi.ResourceOption) (*DevBoxDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.ImageReference == nil {
		return nil, errors.New("invalid value for required argument 'ImageReference'")
	}
	if args.OsStorageType == nil {
		return nil, errors.New("invalid value for required argument 'OsStorageType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:DevBoxDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource DevBoxDefinition
	err := ctx.RegisterResource("azure-native:devcenter/v20220901preview:DevBoxDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDevBoxDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevBoxDefinitionState, opts ...pulumi.ResourceOption) (*DevBoxDefinition, error) {
	var resource DevBoxDefinition
	err := ctx.ReadResource("azure-native:devcenter/v20220901preview:DevBoxDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type devBoxDefinitionState struct {
}

type DevBoxDefinitionState struct {
}

func (DevBoxDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*devBoxDefinitionState)(nil)).Elem()
}

type devBoxDefinitionArgs struct {
	DevBoxDefinitionName *string           `pulumi:"devBoxDefinitionName"`
	DevCenterName        string            `pulumi:"devCenterName"`
	ImageReference       ImageReference    `pulumi:"imageReference"`
	Location             *string           `pulumi:"location"`
	OsStorageType        string            `pulumi:"osStorageType"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	Sku                  Sku               `pulumi:"sku"`
	Tags                 map[string]string `pulumi:"tags"`
}


type DevBoxDefinitionArgs struct {
	DevBoxDefinitionName pulumi.StringPtrInput
	DevCenterName        pulumi.StringInput
	ImageReference       ImageReferenceInput
	Location             pulumi.StringPtrInput
	OsStorageType        pulumi.StringInput
	ResourceGroupName    pulumi.StringInput
	Sku                  SkuInput
	Tags                 pulumi.StringMapInput
}

func (DevBoxDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devBoxDefinitionArgs)(nil)).Elem()
}

type DevBoxDefinitionInput interface {
	pulumi.Input

	ToDevBoxDefinitionOutput() DevBoxDefinitionOutput
	ToDevBoxDefinitionOutputWithContext(ctx context.Context) DevBoxDefinitionOutput
}

func (*DevBoxDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**DevBoxDefinition)(nil)).Elem()
}

func (i *DevBoxDefinition) ToDevBoxDefinitionOutput() DevBoxDefinitionOutput {
	return i.ToDevBoxDefinitionOutputWithContext(context.Background())
}

func (i *DevBoxDefinition) ToDevBoxDefinitionOutputWithContext(ctx context.Context) DevBoxDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevBoxDefinitionOutput)
}

type DevBoxDefinitionOutput struct{ *pulumi.OutputState }

func (DevBoxDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevBoxDefinition)(nil)).Elem()
}

func (o DevBoxDefinitionOutput) ToDevBoxDefinitionOutput() DevBoxDefinitionOutput {
	return o
}

func (o DevBoxDefinitionOutput) ToDevBoxDefinitionOutputWithContext(ctx context.Context) DevBoxDefinitionOutput {
	return o
}

func (o DevBoxDefinitionOutput) ActiveImageReference() ImageReferenceResponseOutput {
	return o.ApplyT(func(v *DevBoxDefinition) ImageReferenceResponseOutput { return v.ActiveImageReference }).(ImageReferenceResponseOutput)
}

func (o DevBoxDefinitionOutput) ImageReference() ImageReferenceResponseOutput {
	return o.ApplyT(func(v *DevBoxDefinition) ImageReferenceResponseOutput { return v.ImageReference }).(ImageReferenceResponseOutput)
}

func (o DevBoxDefinitionOutput) ImageValidationErrorDetails() ImageValidationErrorDetailsResponseOutput {
	return o.ApplyT(func(v *DevBoxDefinition) ImageValidationErrorDetailsResponseOutput {
		return v.ImageValidationErrorDetails
	}).(ImageValidationErrorDetailsResponseOutput)
}

func (o DevBoxDefinitionOutput) ImageValidationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.ImageValidationStatus }).(pulumi.StringOutput)
}

func (o DevBoxDefinitionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DevBoxDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DevBoxDefinitionOutput) OsStorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.OsStorageType }).(pulumi.StringOutput)
}

func (o DevBoxDefinitionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DevBoxDefinitionOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *DevBoxDefinition) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o DevBoxDefinitionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DevBoxDefinition) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DevBoxDefinitionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DevBoxDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DevBoxDefinitionOutput{})
}
