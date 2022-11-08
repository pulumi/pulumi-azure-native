


package v20211001

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
	SystemData          SystemDataResponseOutput                    `pulumi:"systemData"`
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
			Type: pulumi.String("azure-native:securityinsights:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:AutomationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource AutomationRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20211001:AutomationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAutomationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationRuleState, opts ...pulumi.ResourceOption) (*AutomationRule, error) {
	var resource AutomationRule
	err := ctx.ReadResource("azure-native:securityinsights/v20211001:AutomationRule", name, id, state, &resource, opts...)
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
	Actions           []interface{}                 `pulumi:"actions"`
	AutomationRuleId  *string                       `pulumi:"automationRuleId"`
	DisplayName       string                        `pulumi:"displayName"`
	Order             int                           `pulumi:"order"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	TriggeringLogic   AutomationRuleTriggeringLogic `pulumi:"triggeringLogic"`
	WorkspaceName     string                        `pulumi:"workspaceName"`
}


type AutomationRuleArgs struct {
	Actions           pulumi.ArrayInput
	AutomationRuleId  pulumi.StringPtrInput
	DisplayName       pulumi.StringInput
	Order             pulumi.IntInput
	ResourceGroupName pulumi.StringInput
	TriggeringLogic   AutomationRuleTriggeringLogicInput
	WorkspaceName     pulumi.StringInput
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

func (o AutomationRuleOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.ArrayOutput { return v.Actions }).(pulumi.ArrayOutput)
}

func (o AutomationRuleOutput) CreatedBy() ClientInfoResponseOutput {
	return o.ApplyT(func(v *AutomationRule) ClientInfoResponseOutput { return v.CreatedBy }).(ClientInfoResponseOutput)
}

func (o AutomationRuleOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o AutomationRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o AutomationRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AutomationRuleOutput) LastModifiedBy() ClientInfoResponseOutput {
	return o.ApplyT(func(v *AutomationRule) ClientInfoResponseOutput { return v.LastModifiedBy }).(ClientInfoResponseOutput)
}

func (o AutomationRuleOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o AutomationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AutomationRuleOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.IntOutput { return v.Order }).(pulumi.IntOutput)
}

func (o AutomationRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AutomationRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AutomationRuleOutput) TriggeringLogic() AutomationRuleTriggeringLogicResponseOutput {
	return o.ApplyT(func(v *AutomationRule) AutomationRuleTriggeringLogicResponseOutput { return v.TriggeringLogic }).(AutomationRuleTriggeringLogicResponseOutput)
}

func (o AutomationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AutomationRuleOutput{})
}
