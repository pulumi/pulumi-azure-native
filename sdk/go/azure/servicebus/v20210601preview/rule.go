


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Rule struct {
	pulumi.CustomResourceState

	Action            ActionResponsePtrOutput            `pulumi:"action"`
	CorrelationFilter CorrelationFilterResponsePtrOutput `pulumi:"correlationFilter"`
	FilterType        pulumi.StringPtrOutput             `pulumi:"filterType"`
	Name              pulumi.StringOutput                `pulumi:"name"`
	SqlFilter         SqlFilterResponsePtrOutput         `pulumi:"sqlFilter"`
	SystemData        SystemDataResponseOutput           `pulumi:"systemData"`
	Type              pulumi.StringOutput                `pulumi:"type"`
}


func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20210601preview:Rule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus:Rule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus:Rule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Rule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20170401:Rule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:Rule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20180101preview:Rule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:Rule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20210101preview:Rule"),
		},
	})
	opts = append(opts, aliases)
	var resource Rule
	err := ctx.RegisterResource("azure-native:servicebus/v20210601preview:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("azure-native:servicebus/v20210601preview:Rule", name, id, state, &resource, opts...)
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
	Action            *Action            `pulumi:"action"`
	CorrelationFilter *CorrelationFilter `pulumi:"correlationFilter"`
	FilterType        *FilterType        `pulumi:"filterType"`
	NamespaceName     string             `pulumi:"namespaceName"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	RuleName          *string            `pulumi:"ruleName"`
	SqlFilter         *SqlFilter         `pulumi:"sqlFilter"`
	SubscriptionName  string             `pulumi:"subscriptionName"`
	TopicName         string             `pulumi:"topicName"`
}


type RuleArgs struct {
	Action            ActionPtrInput
	CorrelationFilter CorrelationFilterPtrInput
	FilterType        FilterTypePtrInput
	NamespaceName     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	RuleName          pulumi.StringPtrInput
	SqlFilter         SqlFilterPtrInput
	SubscriptionName  pulumi.StringInput
	TopicName         pulumi.StringInput
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
	return reflect.TypeOf((*Rule)(nil))
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rule)(nil))
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterOutputType(RuleOutput{})
}
