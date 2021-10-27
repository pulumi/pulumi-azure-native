


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertsSuppressionRule struct {
	pulumi.CustomResourceState

	AlertType              pulumi.StringOutput                     `pulumi:"alertType"`
	Comment                pulumi.StringPtrOutput                  `pulumi:"comment"`
	ExpirationDateUtc      pulumi.StringPtrOutput                  `pulumi:"expirationDateUtc"`
	LastModifiedUtc        pulumi.StringOutput                     `pulumi:"lastModifiedUtc"`
	Name                   pulumi.StringOutput                     `pulumi:"name"`
	Reason                 pulumi.StringOutput                     `pulumi:"reason"`
	State                  pulumi.StringOutput                     `pulumi:"state"`
	SuppressionAlertsScope SuppressionAlertsScopeResponsePtrOutput `pulumi:"suppressionAlertsScope"`
	Type                   pulumi.StringOutput                     `pulumi:"type"`
}


func NewAlertsSuppressionRule(ctx *pulumi.Context,
	name string, args *AlertsSuppressionRuleArgs, opts ...pulumi.ResourceOption) (*AlertsSuppressionRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertType == nil {
		return nil, errors.New("invalid value for required argument 'AlertType'")
	}
	if args.Reason == nil {
		return nil, errors.New("invalid value for required argument 'Reason'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security/v20190101preview:AlertsSuppressionRule"),
		},
		{
			Type: pulumi.String("azure-native:security:AlertsSuppressionRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:security:AlertsSuppressionRule"),
		},
	})
	opts = append(opts, aliases)
	var resource AlertsSuppressionRule
	err := ctx.RegisterResource("azure-native:security/v20190101preview:AlertsSuppressionRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAlertsSuppressionRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertsSuppressionRuleState, opts ...pulumi.ResourceOption) (*AlertsSuppressionRule, error) {
	var resource AlertsSuppressionRule
	err := ctx.ReadResource("azure-native:security/v20190101preview:AlertsSuppressionRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type alertsSuppressionRuleState struct {
}

type AlertsSuppressionRuleState struct {
}

func (AlertsSuppressionRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertsSuppressionRuleState)(nil)).Elem()
}

type alertsSuppressionRuleArgs struct {
	AlertType                 string                  `pulumi:"alertType"`
	AlertsSuppressionRuleName *string                 `pulumi:"alertsSuppressionRuleName"`
	Comment                   *string                 `pulumi:"comment"`
	ExpirationDateUtc         *string                 `pulumi:"expirationDateUtc"`
	Reason                    string                  `pulumi:"reason"`
	State                     string                  `pulumi:"state"`
	SuppressionAlertsScope    *SuppressionAlertsScope `pulumi:"suppressionAlertsScope"`
}


type AlertsSuppressionRuleArgs struct {
	AlertType                 pulumi.StringInput
	AlertsSuppressionRuleName pulumi.StringPtrInput
	Comment                   pulumi.StringPtrInput
	ExpirationDateUtc         pulumi.StringPtrInput
	Reason                    pulumi.StringInput
	State                     pulumi.StringInput
	SuppressionAlertsScope    SuppressionAlertsScopePtrInput
}

func (AlertsSuppressionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertsSuppressionRuleArgs)(nil)).Elem()
}

type AlertsSuppressionRuleInput interface {
	pulumi.Input

	ToAlertsSuppressionRuleOutput() AlertsSuppressionRuleOutput
	ToAlertsSuppressionRuleOutputWithContext(ctx context.Context) AlertsSuppressionRuleOutput
}

func (*AlertsSuppressionRule) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertsSuppressionRule)(nil))
}

func (i *AlertsSuppressionRule) ToAlertsSuppressionRuleOutput() AlertsSuppressionRuleOutput {
	return i.ToAlertsSuppressionRuleOutputWithContext(context.Background())
}

func (i *AlertsSuppressionRule) ToAlertsSuppressionRuleOutputWithContext(ctx context.Context) AlertsSuppressionRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsSuppressionRuleOutput)
}

type AlertsSuppressionRuleOutput struct{ *pulumi.OutputState }

func (AlertsSuppressionRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertsSuppressionRule)(nil))
}

func (o AlertsSuppressionRuleOutput) ToAlertsSuppressionRuleOutput() AlertsSuppressionRuleOutput {
	return o
}

func (o AlertsSuppressionRuleOutput) ToAlertsSuppressionRuleOutputWithContext(ctx context.Context) AlertsSuppressionRuleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertsSuppressionRuleInput)(nil)).Elem(), &AlertsSuppressionRule{})
	pulumi.RegisterOutputType(AlertsSuppressionRuleOutput{})
}
