


package v20210808preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertProcessingRuleByName struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                         `pulumi:"location"`
	Name       pulumi.StringOutput                         `pulumi:"name"`
	Properties AlertProcessingRulePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                    `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                      `pulumi:"tags"`
	Type       pulumi.StringOutput                         `pulumi:"type"`
}


func NewAlertProcessingRuleByName(ctx *pulumi.Context,
	name string, args *AlertProcessingRuleByNameArgs, opts ...pulumi.ResourceOption) (*AlertProcessingRuleByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	propertiesApplier := func(v AlertProcessingRuleProperties) *AlertProcessingRuleProperties { return v.Defaults() }
	if args.Properties != nil {
		args.Properties = args.Properties.ToAlertProcessingRulePropertiesPtrOutput().Elem().ApplyT(propertiesApplier).(AlertProcessingRulePropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:alertsmanagement:AlertProcessingRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20181102privatepreview:AlertProcessingRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20190505preview:AlertProcessingRuleByName"),
		},
	})
	opts = append(opts, aliases)
	var resource AlertProcessingRuleByName
	err := ctx.RegisterResource("azure-native:alertsmanagement/v20210808preview:AlertProcessingRuleByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAlertProcessingRuleByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertProcessingRuleByNameState, opts ...pulumi.ResourceOption) (*AlertProcessingRuleByName, error) {
	var resource AlertProcessingRuleByName
	err := ctx.ReadResource("azure-native:alertsmanagement/v20210808preview:AlertProcessingRuleByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type alertProcessingRuleByNameState struct {
}

type AlertProcessingRuleByNameState struct {
}

func (AlertProcessingRuleByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertProcessingRuleByNameState)(nil)).Elem()
}

type alertProcessingRuleByNameArgs struct {
	AlertProcessingRuleName *string                        `pulumi:"alertProcessingRuleName"`
	Location                *string                        `pulumi:"location"`
	Properties              *AlertProcessingRuleProperties `pulumi:"properties"`
	ResourceGroupName       string                         `pulumi:"resourceGroupName"`
	Tags                    map[string]string              `pulumi:"tags"`
}


type AlertProcessingRuleByNameArgs struct {
	AlertProcessingRuleName pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	Properties              AlertProcessingRulePropertiesPtrInput
	ResourceGroupName       pulumi.StringInput
	Tags                    pulumi.StringMapInput
}

func (AlertProcessingRuleByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertProcessingRuleByNameArgs)(nil)).Elem()
}

type AlertProcessingRuleByNameInput interface {
	pulumi.Input

	ToAlertProcessingRuleByNameOutput() AlertProcessingRuleByNameOutput
	ToAlertProcessingRuleByNameOutputWithContext(ctx context.Context) AlertProcessingRuleByNameOutput
}

func (*AlertProcessingRuleByName) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertProcessingRuleByName)(nil)).Elem()
}

func (i *AlertProcessingRuleByName) ToAlertProcessingRuleByNameOutput() AlertProcessingRuleByNameOutput {
	return i.ToAlertProcessingRuleByNameOutputWithContext(context.Background())
}

func (i *AlertProcessingRuleByName) ToAlertProcessingRuleByNameOutputWithContext(ctx context.Context) AlertProcessingRuleByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertProcessingRuleByNameOutput)
}

type AlertProcessingRuleByNameOutput struct{ *pulumi.OutputState }

func (AlertProcessingRuleByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertProcessingRuleByName)(nil)).Elem()
}

func (o AlertProcessingRuleByNameOutput) ToAlertProcessingRuleByNameOutput() AlertProcessingRuleByNameOutput {
	return o
}

func (o AlertProcessingRuleByNameOutput) ToAlertProcessingRuleByNameOutputWithContext(ctx context.Context) AlertProcessingRuleByNameOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AlertProcessingRuleByNameOutput{})
}
