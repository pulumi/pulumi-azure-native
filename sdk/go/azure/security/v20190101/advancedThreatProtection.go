


package v20190101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdvancedThreatProtection struct {
	pulumi.CustomResourceState

	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	Name      pulumi.StringOutput  `pulumi:"name"`
	Type      pulumi.StringOutput  `pulumi:"type"`
}


func NewAdvancedThreatProtection(ctx *pulumi.Context,
	name string, args *AdvancedThreatProtectionArgs, opts ...pulumi.ResourceOption) (*AdvancedThreatProtection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:AdvancedThreatProtection"),
		},
		{
			Type: pulumi.String("azure-native:security/v20170801preview:AdvancedThreatProtection"),
		},
	})
	opts = append(opts, aliases)
	var resource AdvancedThreatProtection
	err := ctx.RegisterResource("azure-native:security/v20190101:AdvancedThreatProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAdvancedThreatProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdvancedThreatProtectionState, opts ...pulumi.ResourceOption) (*AdvancedThreatProtection, error) {
	var resource AdvancedThreatProtection
	err := ctx.ReadResource("azure-native:security/v20190101:AdvancedThreatProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type advancedThreatProtectionState struct {
}

type AdvancedThreatProtectionState struct {
}

func (AdvancedThreatProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*advancedThreatProtectionState)(nil)).Elem()
}

type advancedThreatProtectionArgs struct {
	IsEnabled   *bool   `pulumi:"isEnabled"`
	ResourceId  string  `pulumi:"resourceId"`
	SettingName *string `pulumi:"settingName"`
}


type AdvancedThreatProtectionArgs struct {
	IsEnabled   pulumi.BoolPtrInput
	ResourceId  pulumi.StringInput
	SettingName pulumi.StringPtrInput
}

func (AdvancedThreatProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*advancedThreatProtectionArgs)(nil)).Elem()
}

type AdvancedThreatProtectionInput interface {
	pulumi.Input

	ToAdvancedThreatProtectionOutput() AdvancedThreatProtectionOutput
	ToAdvancedThreatProtectionOutputWithContext(ctx context.Context) AdvancedThreatProtectionOutput
}

func (*AdvancedThreatProtection) ElementType() reflect.Type {
	return reflect.TypeOf((**AdvancedThreatProtection)(nil)).Elem()
}

func (i *AdvancedThreatProtection) ToAdvancedThreatProtectionOutput() AdvancedThreatProtectionOutput {
	return i.ToAdvancedThreatProtectionOutputWithContext(context.Background())
}

func (i *AdvancedThreatProtection) ToAdvancedThreatProtectionOutputWithContext(ctx context.Context) AdvancedThreatProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedThreatProtectionOutput)
}

type AdvancedThreatProtectionOutput struct{ *pulumi.OutputState }

func (AdvancedThreatProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdvancedThreatProtection)(nil)).Elem()
}

func (o AdvancedThreatProtectionOutput) ToAdvancedThreatProtectionOutput() AdvancedThreatProtectionOutput {
	return o
}

func (o AdvancedThreatProtectionOutput) ToAdvancedThreatProtectionOutputWithContext(ctx context.Context) AdvancedThreatProtectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AdvancedThreatProtectionOutput{})
}
