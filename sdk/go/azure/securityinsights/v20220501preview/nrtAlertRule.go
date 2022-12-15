


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NrtAlertRule struct {
	pulumi.CustomResourceState

	AlertDetailsOverride  AlertDetailsOverrideResponsePtrOutput  `pulumi:"alertDetailsOverride"`
	AlertRuleTemplateName pulumi.StringPtrOutput                 `pulumi:"alertRuleTemplateName"`
	CustomDetails         pulumi.StringMapOutput                 `pulumi:"customDetails"`
	Description           pulumi.StringPtrOutput                 `pulumi:"description"`
	DisplayName           pulumi.StringOutput                    `pulumi:"displayName"`
	Enabled               pulumi.BoolOutput                      `pulumi:"enabled"`
	EntityMappings        EntityMappingResponseArrayOutput       `pulumi:"entityMappings"`
	Etag                  pulumi.StringPtrOutput                 `pulumi:"etag"`
	IncidentConfiguration IncidentConfigurationResponsePtrOutput `pulumi:"incidentConfiguration"`
	Kind                  pulumi.StringOutput                    `pulumi:"kind"`
	LastModifiedUtc       pulumi.StringOutput                    `pulumi:"lastModifiedUtc"`
	Name                  pulumi.StringOutput                    `pulumi:"name"`
	Query                 pulumi.StringOutput                    `pulumi:"query"`
	Severity              pulumi.StringOutput                    `pulumi:"severity"`
	SuppressionDuration   pulumi.StringOutput                    `pulumi:"suppressionDuration"`
	SuppressionEnabled    pulumi.BoolOutput                      `pulumi:"suppressionEnabled"`
	SystemData            SystemDataResponseOutput               `pulumi:"systemData"`
	Tactics               pulumi.StringArrayOutput               `pulumi:"tactics"`
	Techniques            pulumi.StringArrayOutput               `pulumi:"techniques"`
	TemplateVersion       pulumi.StringPtrOutput                 `pulumi:"templateVersion"`
	Type                  pulumi.StringOutput                    `pulumi:"type"`
}


func NewNrtAlertRule(ctx *pulumi.Context,
	name string, args *NrtAlertRuleArgs, opts ...pulumi.ResourceOption) (*NrtAlertRule, error) {
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
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
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
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("NRT")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:NrtAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NrtAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20220501preview:NrtAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNrtAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NrtAlertRuleState, opts ...pulumi.ResourceOption) (*NrtAlertRule, error) {
	var resource NrtAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20220501preview:NrtAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type nrtAlertRuleState struct {
}

type NrtAlertRuleState struct {
}

func (NrtAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*nrtAlertRuleState)(nil)).Elem()
}

type nrtAlertRuleArgs struct {
	AlertDetailsOverride  *AlertDetailsOverride  `pulumi:"alertDetailsOverride"`
	AlertRuleTemplateName *string                `pulumi:"alertRuleTemplateName"`
	CustomDetails         map[string]string      `pulumi:"customDetails"`
	Description           *string                `pulumi:"description"`
	DisplayName           string                 `pulumi:"displayName"`
	Enabled               bool                   `pulumi:"enabled"`
	EntityMappings        []EntityMapping        `pulumi:"entityMappings"`
	IncidentConfiguration *IncidentConfiguration `pulumi:"incidentConfiguration"`
	Kind                  string                 `pulumi:"kind"`
	Query                 string                 `pulumi:"query"`
	ResourceGroupName     string                 `pulumi:"resourceGroupName"`
	RuleId                *string                `pulumi:"ruleId"`
	Severity              string                 `pulumi:"severity"`
	SuppressionDuration   string                 `pulumi:"suppressionDuration"`
	SuppressionEnabled    bool                   `pulumi:"suppressionEnabled"`
	Tactics               []string               `pulumi:"tactics"`
	Techniques            []string               `pulumi:"techniques"`
	TemplateVersion       *string                `pulumi:"templateVersion"`
	WorkspaceName         string                 `pulumi:"workspaceName"`
}


type NrtAlertRuleArgs struct {
	AlertDetailsOverride  AlertDetailsOverridePtrInput
	AlertRuleTemplateName pulumi.StringPtrInput
	CustomDetails         pulumi.StringMapInput
	Description           pulumi.StringPtrInput
	DisplayName           pulumi.StringInput
	Enabled               pulumi.BoolInput
	EntityMappings        EntityMappingArrayInput
	IncidentConfiguration IncidentConfigurationPtrInput
	Kind                  pulumi.StringInput
	Query                 pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	RuleId                pulumi.StringPtrInput
	Severity              pulumi.StringInput
	SuppressionDuration   pulumi.StringInput
	SuppressionEnabled    pulumi.BoolInput
	Tactics               pulumi.StringArrayInput
	Techniques            pulumi.StringArrayInput
	TemplateVersion       pulumi.StringPtrInput
	WorkspaceName         pulumi.StringInput
}

func (NrtAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nrtAlertRuleArgs)(nil)).Elem()
}

type NrtAlertRuleInput interface {
	pulumi.Input

	ToNrtAlertRuleOutput() NrtAlertRuleOutput
	ToNrtAlertRuleOutputWithContext(ctx context.Context) NrtAlertRuleOutput
}

func (*NrtAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NrtAlertRule)(nil)).Elem()
}

func (i *NrtAlertRule) ToNrtAlertRuleOutput() NrtAlertRuleOutput {
	return i.ToNrtAlertRuleOutputWithContext(context.Background())
}

func (i *NrtAlertRule) ToNrtAlertRuleOutputWithContext(ctx context.Context) NrtAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NrtAlertRuleOutput)
}

type NrtAlertRuleOutput struct{ *pulumi.OutputState }

func (NrtAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NrtAlertRule)(nil)).Elem()
}

func (o NrtAlertRuleOutput) ToNrtAlertRuleOutput() NrtAlertRuleOutput {
	return o
}

func (o NrtAlertRuleOutput) ToNrtAlertRuleOutputWithContext(ctx context.Context) NrtAlertRuleOutput {
	return o
}

func (o NrtAlertRuleOutput) AlertDetailsOverride() AlertDetailsOverrideResponsePtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) AlertDetailsOverrideResponsePtrOutput { return v.AlertDetailsOverride }).(AlertDetailsOverrideResponsePtrOutput)
}

func (o NrtAlertRuleOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringPtrOutput { return v.AlertRuleTemplateName }).(pulumi.StringPtrOutput)
}

func (o NrtAlertRuleOutput) CustomDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringMapOutput { return v.CustomDetails }).(pulumi.StringMapOutput)
}

func (o NrtAlertRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NrtAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o NrtAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o NrtAlertRuleOutput) EntityMappings() EntityMappingResponseArrayOutput {
	return o.ApplyT(func(v *NrtAlertRule) EntityMappingResponseArrayOutput { return v.EntityMappings }).(EntityMappingResponseArrayOutput)
}

func (o NrtAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NrtAlertRuleOutput) IncidentConfiguration() IncidentConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) IncidentConfigurationResponsePtrOutput { return v.IncidentConfiguration }).(IncidentConfigurationResponsePtrOutput)
}

func (o NrtAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o NrtAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o NrtAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NrtAlertRuleOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

func (o NrtAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o NrtAlertRuleOutput) SuppressionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.SuppressionDuration }).(pulumi.StringOutput)
}

func (o NrtAlertRuleOutput) SuppressionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.BoolOutput { return v.SuppressionEnabled }).(pulumi.BoolOutput)
}

func (o NrtAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NrtAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o NrtAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o NrtAlertRuleOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringArrayOutput { return v.Techniques }).(pulumi.StringArrayOutput)
}

func (o NrtAlertRuleOutput) TemplateVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringPtrOutput { return v.TemplateVersion }).(pulumi.StringPtrOutput)
}

func (o NrtAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NrtAlertRuleOutput{})
}
