


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Rule struct {
	pulumi.CustomResourceState

	Actions                 pulumi.ArrayOutput       `pulumi:"actions"`
	Conditions              pulumi.ArrayOutput       `pulumi:"conditions"`
	DeploymentStatus        pulumi.StringOutput      `pulumi:"deploymentStatus"`
	MatchProcessingBehavior pulumi.StringPtrOutput   `pulumi:"matchProcessingBehavior"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	Order                   pulumi.IntOutput         `pulumi:"order"`
	ProvisioningState       pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData              SystemDataResponseOutput `pulumi:"systemData"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
}


func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Order == nil {
		return nil, errors.New("invalid value for required argument 'Order'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleSetName == nil {
		return nil, errors.New("invalid value for required argument 'RuleSetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:Rule"),
		},
	})
	opts = append(opts, aliases)
	var resource Rule
	err := ctx.RegisterResource("azure-native:cdn/v20200901:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("azure-native:cdn/v20200901:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ruleState struct {
}

type RuleState struct {
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	Actions                 []interface{} `pulumi:"actions"`
	Conditions              []interface{} `pulumi:"conditions"`
	MatchProcessingBehavior *string       `pulumi:"matchProcessingBehavior"`
	Order                   int           `pulumi:"order"`
	ProfileName             string        `pulumi:"profileName"`
	ResourceGroupName       string        `pulumi:"resourceGroupName"`
	RuleName                *string       `pulumi:"ruleName"`
	RuleSetName             string        `pulumi:"ruleSetName"`
}


type RuleArgs struct {
	Actions                 pulumi.ArrayInput
	Conditions              pulumi.ArrayInput
	MatchProcessingBehavior pulumi.StringPtrInput
	Order                   pulumi.IntInput
	ProfileName             pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	RuleName                pulumi.StringPtrInput
	RuleSetName             pulumi.StringInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RuleOutput{})
}
