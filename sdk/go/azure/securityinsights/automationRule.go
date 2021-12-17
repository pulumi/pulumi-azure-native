


package securityinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutomationRule struct {
	pulumi.CustomResourceState

	Actions             pulumi.ArrayOutput                          `pulumi:"actions"`
	CreatedBy           ClientInfoResponseOutput                    `pulumi:"createdBy"`
	CreatedTimeUtc      pulumi.StringOutput                         `pulumi:"createdTimeUtc"`
	DisplayName         pulumi.StringOutput                         `pulumi:"displayName"`
	Etag                pulumi.StringPtrOutput                      `pulumi:"etag"`
	LastModifiedBy      ClientInfoResponseOutput                    `pulumi:"lastModifiedBy"`
	LastModifiedTimeUtc pulumi.StringOutput                         `pulumi:"lastModifiedTimeUtc"`
	Name                pulumi.StringOutput                         `pulumi:"name"`
	Order               pulumi.IntOutput                            `pulumi:"order"`
	TriggeringLogic     AutomationRuleTriggeringLogicResponseOutput `pulumi:"triggeringLogic"`
	Type                pulumi.StringOutput                         `pulumi:"type"`
}


func NewAutomationRule(ctx *pulumi.Context,
	name string, args *AutomationRuleArgs, opts ...pulumi.ResourceOption) (*AutomationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.Order == nil {
		return nil, errors.New("invalid value for required argument 'Order'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TriggeringLogic == nil {
		return nil, errors.New("invalid value for required argument 'TriggeringLogic'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:AutomationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource AutomationRule
	err := ctx.RegisterResource("azure-native:securityinsights:AutomationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAutomationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationRuleState, opts ...pulumi.ResourceOption) (*AutomationRule, error) {
	var resource AutomationRule
	err := ctx.ReadResource("azure-native:securityinsights:AutomationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type automationRuleState struct {
}

type AutomationRuleState struct {
}

func (AutomationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationRuleState)(nil)).Elem()
}

type automationRuleArgs struct {
	Actions                             []interface{}                 `pulumi:"actions"`
	AutomationRuleId                    *string                       `pulumi:"automationRuleId"`
	DisplayName                         string                        `pulumi:"displayName"`
	Etag                                *string                       `pulumi:"etag"`
	OperationalInsightsResourceProvider string                        `pulumi:"operationalInsightsResourceProvider"`
	Order                               int                           `pulumi:"order"`
	ResourceGroupName                   string                        `pulumi:"resourceGroupName"`
	TriggeringLogic                     AutomationRuleTriggeringLogic `pulumi:"triggeringLogic"`
	WorkspaceName                       string                        `pulumi:"workspaceName"`
}


type AutomationRuleArgs struct {
	Actions                             pulumi.ArrayInput
	AutomationRuleId                    pulumi.StringPtrInput
	DisplayName                         pulumi.StringInput
	Etag                                pulumi.StringPtrInput
	OperationalInsightsResourceProvider pulumi.StringInput
	Order                               pulumi.IntInput
	ResourceGroupName                   pulumi.StringInput
	TriggeringLogic                     AutomationRuleTriggeringLogicInput
	WorkspaceName                       pulumi.StringInput
}

func (AutomationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationRuleArgs)(nil)).Elem()
}

type AutomationRuleInput interface {
	pulumi.Input

	ToAutomationRuleOutput() AutomationRuleOutput
	ToAutomationRuleOutputWithContext(ctx context.Context) AutomationRuleOutput
}

func (*AutomationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRule)(nil)).Elem()
}

func (i *AutomationRule) ToAutomationRuleOutput() AutomationRuleOutput {
	return i.ToAutomationRuleOutputWithContext(context.Background())
}

func (i *AutomationRule) ToAutomationRuleOutputWithContext(ctx context.Context) AutomationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleOutput)
}

type AutomationRuleOutput struct{ *pulumi.OutputState }

func (AutomationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRule)(nil)).Elem()
}

func (o AutomationRuleOutput) ToAutomationRuleOutput() AutomationRuleOutput {
	return o
}

func (o AutomationRuleOutput) ToAutomationRuleOutputWithContext(ctx context.Context) AutomationRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AutomationRuleOutput{})
}
