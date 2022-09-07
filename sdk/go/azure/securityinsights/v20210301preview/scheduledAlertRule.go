


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScheduledAlertRule struct {
	pulumi.CustomResourceState

	AlertDetailsOverride  AlertDetailsOverrideResponsePtrOutput  `pulumi:"alertDetailsOverride"`
	AlertRuleTemplateName pulumi.StringPtrOutput                 `pulumi:"alertRuleTemplateName"`
	CustomDetails         pulumi.StringMapOutput                 `pulumi:"customDetails"`
	Description           pulumi.StringPtrOutput                 `pulumi:"description"`
	DisplayName           pulumi.StringOutput                    `pulumi:"displayName"`
	Enabled               pulumi.BoolOutput                      `pulumi:"enabled"`
	EntityMappings        EntityMappingResponseArrayOutput       `pulumi:"entityMappings"`
	Etag                  pulumi.StringPtrOutput                 `pulumi:"etag"`
	EventGroupingSettings EventGroupingSettingsResponsePtrOutput `pulumi:"eventGroupingSettings"`
	IncidentConfiguration IncidentConfigurationResponsePtrOutput `pulumi:"incidentConfiguration"`
	Kind                  pulumi.StringOutput                    `pulumi:"kind"`
	LastModifiedUtc       pulumi.StringOutput                    `pulumi:"lastModifiedUtc"`
	Name                  pulumi.StringOutput                    `pulumi:"name"`
	Query                 pulumi.StringOutput                    `pulumi:"query"`
	QueryFrequency        pulumi.StringOutput                    `pulumi:"queryFrequency"`
	QueryPeriod           pulumi.StringOutput                    `pulumi:"queryPeriod"`
	Severity              pulumi.StringOutput                    `pulumi:"severity"`
	SuppressionDuration   pulumi.StringOutput                    `pulumi:"suppressionDuration"`
	SuppressionEnabled    pulumi.BoolOutput                      `pulumi:"suppressionEnabled"`
	SystemData            SystemDataResponseOutput               `pulumi:"systemData"`
	Tactics               pulumi.StringArrayOutput               `pulumi:"tactics"`
	TriggerOperator       pulumi.StringOutput                    `pulumi:"triggerOperator"`
	TriggerThreshold      pulumi.IntOutput                       `pulumi:"triggerThreshold"`
	Type                  pulumi.StringOutput                    `pulumi:"type"`
}


func NewScheduledAlertRule(ctx *pulumi.Context,
	name string, args *ScheduledAlertRuleArgs, opts ...pulumi.ResourceOption) (*ScheduledAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.QueryFrequency == nil {
		return nil, errors.New("invalid value for required argument 'QueryFrequency'")
	}
	if args.QueryPeriod == nil {
		return nil, errors.New("invalid value for required argument 'QueryPeriod'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.SuppressionDuration == nil {
		return nil, errors.New("invalid value for required argument 'SuppressionDuration'")
	}
	if args.SuppressionEnabled == nil {
		return nil, errors.New("invalid value for required argument 'SuppressionEnabled'")
	}
	if args.TriggerOperator == nil {
		return nil, errors.New("invalid value for required argument 'TriggerOperator'")
	}
	if args.TriggerThreshold == nil {
		return nil, errors.New("invalid value for required argument 'TriggerThreshold'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Scheduled")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:ScheduledAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:ScheduledAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScheduledAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledAlertRuleState, opts ...pulumi.ResourceOption) (*ScheduledAlertRule, error) {
	var resource ScheduledAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:ScheduledAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scheduledAlertRuleState struct {
}

type ScheduledAlertRuleState struct {
}

func (ScheduledAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledAlertRuleState)(nil)).Elem()
}

type scheduledAlertRuleArgs struct {
	AlertDetailsOverride                *AlertDetailsOverride  `pulumi:"alertDetailsOverride"`
	AlertRuleTemplateName               *string                `pulumi:"alertRuleTemplateName"`
	CustomDetails                       map[string]string      `pulumi:"customDetails"`
	Description                         *string                `pulumi:"description"`
	DisplayName                         string                 `pulumi:"displayName"`
	Enabled                             bool                   `pulumi:"enabled"`
	EntityMappings                      []EntityMapping        `pulumi:"entityMappings"`
	EventGroupingSettings               *EventGroupingSettings `pulumi:"eventGroupingSettings"`
	IncidentConfiguration               *IncidentConfiguration `pulumi:"incidentConfiguration"`
	Kind                                string                 `pulumi:"kind"`
	OperationalInsightsResourceProvider string                 `pulumi:"operationalInsightsResourceProvider"`
	Query                               string                 `pulumi:"query"`
	QueryFrequency                      string                 `pulumi:"queryFrequency"`
	QueryPeriod                         string                 `pulumi:"queryPeriod"`
	ResourceGroupName                   string                 `pulumi:"resourceGroupName"`
	RuleId                              *string                `pulumi:"ruleId"`
	Severity                            string                 `pulumi:"severity"`
	SuppressionDuration                 string                 `pulumi:"suppressionDuration"`
	SuppressionEnabled                  bool                   `pulumi:"suppressionEnabled"`
	Tactics                             []string               `pulumi:"tactics"`
	TriggerOperator                     TriggerOperator        `pulumi:"triggerOperator"`
	TriggerThreshold                    int                    `pulumi:"triggerThreshold"`
	WorkspaceName                       string                 `pulumi:"workspaceName"`
}


type ScheduledAlertRuleArgs struct {
	AlertDetailsOverride                AlertDetailsOverridePtrInput
	AlertRuleTemplateName               pulumi.StringPtrInput
	CustomDetails                       pulumi.StringMapInput
	Description                         pulumi.StringPtrInput
	DisplayName                         pulumi.StringInput
	Enabled                             pulumi.BoolInput
	EntityMappings                      EntityMappingArrayInput
	EventGroupingSettings               EventGroupingSettingsPtrInput
	IncidentConfiguration               IncidentConfigurationPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	Query                               pulumi.StringInput
	QueryFrequency                      pulumi.StringInput
	QueryPeriod                         pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	RuleId                              pulumi.StringPtrInput
	Severity                            pulumi.StringInput
	SuppressionDuration                 pulumi.StringInput
	SuppressionEnabled                  pulumi.BoolInput
	Tactics                             pulumi.StringArrayInput
	TriggerOperator                     TriggerOperatorInput
	TriggerThreshold                    pulumi.IntInput
	WorkspaceName                       pulumi.StringInput
}

func (ScheduledAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledAlertRuleArgs)(nil)).Elem()
}

type ScheduledAlertRuleInput interface {
	pulumi.Input

	ToScheduledAlertRuleOutput() ScheduledAlertRuleOutput
	ToScheduledAlertRuleOutputWithContext(ctx context.Context) ScheduledAlertRuleOutput
}

func (*ScheduledAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAlertRule)(nil)).Elem()
}

func (i *ScheduledAlertRule) ToScheduledAlertRuleOutput() ScheduledAlertRuleOutput {
	return i.ToScheduledAlertRuleOutputWithContext(context.Background())
}

func (i *ScheduledAlertRule) ToScheduledAlertRuleOutputWithContext(ctx context.Context) ScheduledAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledAlertRuleOutput)
}

type ScheduledAlertRuleOutput struct{ *pulumi.OutputState }

func (ScheduledAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAlertRule)(nil)).Elem()
}

func (o ScheduledAlertRuleOutput) ToScheduledAlertRuleOutput() ScheduledAlertRuleOutput {
	return o
}

func (o ScheduledAlertRuleOutput) ToScheduledAlertRuleOutputWithContext(ctx context.Context) ScheduledAlertRuleOutput {
	return o
}

func (o ScheduledAlertRuleOutput) AlertDetailsOverride() AlertDetailsOverrideResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) AlertDetailsOverrideResponsePtrOutput { return v.AlertDetailsOverride }).(AlertDetailsOverrideResponsePtrOutput)
}

func (o ScheduledAlertRuleOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringPtrOutput { return v.AlertRuleTemplateName }).(pulumi.StringPtrOutput)
}

func (o ScheduledAlertRuleOutput) CustomDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringMapOutput { return v.CustomDetails }).(pulumi.StringMapOutput)
}

func (o ScheduledAlertRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ScheduledAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ScheduledAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ScheduledAlertRuleOutput) EntityMappings() EntityMappingResponseArrayOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) EntityMappingResponseArrayOutput { return v.EntityMappings }).(EntityMappingResponseArrayOutput)
}

func (o ScheduledAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ScheduledAlertRuleOutput) EventGroupingSettings() EventGroupingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) EventGroupingSettingsResponsePtrOutput { return v.EventGroupingSettings }).(EventGroupingSettingsResponsePtrOutput)
}

func (o ScheduledAlertRuleOutput) IncidentConfiguration() IncidentConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) IncidentConfigurationResponsePtrOutput { return v.IncidentConfiguration }).(IncidentConfigurationResponsePtrOutput)
}

func (o ScheduledAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ScheduledAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o ScheduledAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScheduledAlertRuleOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

func (o ScheduledAlertRuleOutput) QueryFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.QueryFrequency }).(pulumi.StringOutput)
}

func (o ScheduledAlertRuleOutput) QueryPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.QueryPeriod }).(pulumi.StringOutput)
}

func (o ScheduledAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o ScheduledAlertRuleOutput) SuppressionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.SuppressionDuration }).(pulumi.StringOutput)
}

func (o ScheduledAlertRuleOutput) SuppressionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.BoolOutput { return v.SuppressionEnabled }).(pulumi.BoolOutput)
}

func (o ScheduledAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScheduledAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o ScheduledAlertRuleOutput) TriggerOperator() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.TriggerOperator }).(pulumi.StringOutput)
}

func (o ScheduledAlertRuleOutput) TriggerThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.IntOutput { return v.TriggerThreshold }).(pulumi.IntOutput)
}

func (o ScheduledAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledAlertRuleOutput{})
}
