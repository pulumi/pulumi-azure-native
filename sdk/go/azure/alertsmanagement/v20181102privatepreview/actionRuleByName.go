


package v20181102privatepreview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionRuleByName struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                `pulumi:"location"`
	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties ActionRulePropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput             `pulumi:"tags"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewActionRuleByName(ctx *pulumi.Context,
	name string, args *ActionRuleByNameArgs, opts ...pulumi.ResourceOption) (*ActionRuleByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:alertsmanagement:ActionRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20190505preview:ActionRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20210808preview:ActionRuleByName"),
		},
	})
	opts = append(opts, aliases)
	var resource ActionRuleByName
	err := ctx.RegisterResource("azure-native:alertsmanagement/v20181102privatepreview:ActionRuleByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetActionRuleByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionRuleByNameState, opts ...pulumi.ResourceOption) (*ActionRuleByName, error) {
	var resource ActionRuleByName
	err := ctx.ReadResource("azure-native:alertsmanagement/v20181102privatepreview:ActionRuleByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type actionRuleByNameState struct {
}

type ActionRuleByNameState struct {
}

func (ActionRuleByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionRuleByNameState)(nil)).Elem()
}

type actionRuleByNameArgs struct {
	ActionRuleName *string               `pulumi:"actionRuleName"`
	Location       *string               `pulumi:"location"`
	Properties     *ActionRuleProperties `pulumi:"properties"`
	ResourceGroup  string                `pulumi:"resourceGroup"`
	Tags           map[string]string     `pulumi:"tags"`
}


type ActionRuleByNameArgs struct {
	ActionRuleName pulumi.StringPtrInput
	Location       pulumi.StringPtrInput
	Properties     ActionRulePropertiesPtrInput
	ResourceGroup  pulumi.StringInput
	Tags           pulumi.StringMapInput
}

func (ActionRuleByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionRuleByNameArgs)(nil)).Elem()
}

type ActionRuleByNameInput interface {
	pulumi.Input

	ToActionRuleByNameOutput() ActionRuleByNameOutput
	ToActionRuleByNameOutputWithContext(ctx context.Context) ActionRuleByNameOutput
}

func (*ActionRuleByName) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRuleByName)(nil)).Elem()
}

func (i *ActionRuleByName) ToActionRuleByNameOutput() ActionRuleByNameOutput {
	return i.ToActionRuleByNameOutputWithContext(context.Background())
}

func (i *ActionRuleByName) ToActionRuleByNameOutputWithContext(ctx context.Context) ActionRuleByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRuleByNameOutput)
}

type ActionRuleByNameOutput struct{ *pulumi.OutputState }

func (ActionRuleByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRuleByName)(nil)).Elem()
}

func (o ActionRuleByNameOutput) ToActionRuleByNameOutput() ActionRuleByNameOutput {
	return o
}

func (o ActionRuleByNameOutput) ToActionRuleByNameOutputWithContext(ctx context.Context) ActionRuleByNameOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ActionRuleByNameOutput{})
}
