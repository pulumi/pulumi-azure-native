


package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagRule struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties MonitoringTagRulesPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                   `pulumi:"systemData"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewTagRule(ctx *pulumi.Context,
	name string, args *TagRuleArgs, opts ...pulumi.ResourceOption) (*TagRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:elastic/v20200701:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic:TagRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:elastic:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20200701preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:elastic/v20200701preview:TagRule"),
		},
	})
	opts = append(opts, aliases)
	var resource TagRule
	err := ctx.RegisterResource("azure-native:elastic/v20200701:TagRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTagRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagRuleState, opts ...pulumi.ResourceOption) (*TagRule, error) {
	var resource TagRule
	err := ctx.ReadResource("azure-native:elastic/v20200701:TagRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tagRuleState struct {
}

type TagRuleState struct {
}

func (TagRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagRuleState)(nil)).Elem()
}

type tagRuleArgs struct {
	MonitorName       string                        `pulumi:"monitorName"`
	Properties        *MonitoringTagRulesProperties `pulumi:"properties"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	RuleSetName       *string                       `pulumi:"ruleSetName"`
}


type TagRuleArgs struct {
	MonitorName       pulumi.StringInput
	Properties        MonitoringTagRulesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	RuleSetName       pulumi.StringPtrInput
}

func (TagRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagRuleArgs)(nil)).Elem()
}

type TagRuleInput interface {
	pulumi.Input

	ToTagRuleOutput() TagRuleOutput
	ToTagRuleOutputWithContext(ctx context.Context) TagRuleOutput
}

func (*TagRule) ElementType() reflect.Type {
	return reflect.TypeOf((*TagRule)(nil))
}

func (i *TagRule) ToTagRuleOutput() TagRuleOutput {
	return i.ToTagRuleOutputWithContext(context.Background())
}

func (i *TagRule) ToTagRuleOutputWithContext(ctx context.Context) TagRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagRuleOutput)
}

type TagRuleOutput struct{ *pulumi.OutputState }

func (TagRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagRule)(nil))
}

func (o TagRuleOutput) ToTagRuleOutput() TagRuleOutput {
	return o
}

func (o TagRuleOutput) ToTagRuleOutputWithContext(ctx context.Context) TagRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagRuleOutput{})
}
