


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubAccountTagRule struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties MonitoringTagRulesPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                   `pulumi:"systemData"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewSubAccountTagRule(ctx *pulumi.Context,
	name string, args *SubAccountTagRuleArgs, opts ...pulumi.ResourceOption) (*SubAccountTagRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubAccountName == nil {
		return nil, errors.New("invalid value for required argument 'SubAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logz/v20201001preview:SubAccountTagRule"),
		},
		{
			Type: pulumi.String("azure-native:logz:SubAccountTagRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:logz:SubAccountTagRule"),
		},
		{
			Type: pulumi.String("azure-native:logz/v20201001:SubAccountTagRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:logz/v20201001:SubAccountTagRule"),
		},
	})
	opts = append(opts, aliases)
	var resource SubAccountTagRule
	err := ctx.RegisterResource("azure-native:logz/v20201001preview:SubAccountTagRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubAccountTagRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubAccountTagRuleState, opts ...pulumi.ResourceOption) (*SubAccountTagRule, error) {
	var resource SubAccountTagRule
	err := ctx.ReadResource("azure-native:logz/v20201001preview:SubAccountTagRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subAccountTagRuleState struct {
}

type SubAccountTagRuleState struct {
}

func (SubAccountTagRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountTagRuleState)(nil)).Elem()
}

type subAccountTagRuleArgs struct {
	MonitorName       string                        `pulumi:"monitorName"`
	Properties        *MonitoringTagRulesProperties `pulumi:"properties"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	RuleSetName       *string                       `pulumi:"ruleSetName"`
	SubAccountName    string                        `pulumi:"subAccountName"`
}


type SubAccountTagRuleArgs struct {
	MonitorName       pulumi.StringInput
	Properties        MonitoringTagRulesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	RuleSetName       pulumi.StringPtrInput
	SubAccountName    pulumi.StringInput
}

func (SubAccountTagRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountTagRuleArgs)(nil)).Elem()
}

type SubAccountTagRuleInput interface {
	pulumi.Input

	ToSubAccountTagRuleOutput() SubAccountTagRuleOutput
	ToSubAccountTagRuleOutputWithContext(ctx context.Context) SubAccountTagRuleOutput
}

func (*SubAccountTagRule) ElementType() reflect.Type {
	return reflect.TypeOf((*SubAccountTagRule)(nil))
}

func (i *SubAccountTagRule) ToSubAccountTagRuleOutput() SubAccountTagRuleOutput {
	return i.ToSubAccountTagRuleOutputWithContext(context.Background())
}

func (i *SubAccountTagRule) ToSubAccountTagRuleOutputWithContext(ctx context.Context) SubAccountTagRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubAccountTagRuleOutput)
}

type SubAccountTagRuleOutput struct{ *pulumi.OutputState }

func (SubAccountTagRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubAccountTagRule)(nil))
}

func (o SubAccountTagRuleOutput) ToSubAccountTagRuleOutput() SubAccountTagRuleOutput {
	return o
}

func (o SubAccountTagRuleOutput) ToSubAccountTagRuleOutputWithContext(ctx context.Context) SubAccountTagRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubAccountTagRuleOutput{})
}
