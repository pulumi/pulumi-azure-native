


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MicrosoftSecurityIncidentCreationAlertRule struct {
	pulumi.CustomResourceState

	AlertRuleTemplateName     pulumi.StringPtrOutput   `pulumi:"alertRuleTemplateName"`
	Description               pulumi.StringPtrOutput   `pulumi:"description"`
	DisplayName               pulumi.StringOutput      `pulumi:"displayName"`
	DisplayNamesExcludeFilter pulumi.StringArrayOutput `pulumi:"displayNamesExcludeFilter"`
	DisplayNamesFilter        pulumi.StringArrayOutput `pulumi:"displayNamesFilter"`
	Enabled                   pulumi.BoolOutput        `pulumi:"enabled"`
	Etag                      pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind                      pulumi.StringOutput      `pulumi:"kind"`
	LastModifiedUtc           pulumi.StringOutput      `pulumi:"lastModifiedUtc"`
	Name                      pulumi.StringOutput      `pulumi:"name"`
	ProductFilter             pulumi.StringOutput      `pulumi:"productFilter"`
	SeveritiesFilter          pulumi.StringArrayOutput `pulumi:"severitiesFilter"`
	SystemData                SystemDataResponseOutput `pulumi:"systemData"`
	Type                      pulumi.StringOutput      `pulumi:"type"`
}


func NewMicrosoftSecurityIncidentCreationAlertRule(ctx *pulumi.Context,
	name string, args *MicrosoftSecurityIncidentCreationAlertRuleArgs, opts ...pulumi.ResourceOption) (*MicrosoftSecurityIncidentCreationAlertRule, error) {
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
	if args.ProductFilter == nil {
		return nil, errors.New("invalid value for required argument 'ProductFilter'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("MicrosoftSecurityIncidentCreation")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource MicrosoftSecurityIncidentCreationAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20220101preview:MicrosoftSecurityIncidentCreationAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMicrosoftSecurityIncidentCreationAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MicrosoftSecurityIncidentCreationAlertRuleState, opts ...pulumi.ResourceOption) (*MicrosoftSecurityIncidentCreationAlertRule, error) {
	var resource MicrosoftSecurityIncidentCreationAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20220101preview:MicrosoftSecurityIncidentCreationAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type microsoftSecurityIncidentCreationAlertRuleState struct {
}

type MicrosoftSecurityIncidentCreationAlertRuleState struct {
}

func (MicrosoftSecurityIncidentCreationAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftSecurityIncidentCreationAlertRuleState)(nil)).Elem()
}

type microsoftSecurityIncidentCreationAlertRuleArgs struct {
	AlertRuleTemplateName     *string  `pulumi:"alertRuleTemplateName"`
	Description               *string  `pulumi:"description"`
	DisplayName               string   `pulumi:"displayName"`
	DisplayNamesExcludeFilter []string `pulumi:"displayNamesExcludeFilter"`
	DisplayNamesFilter        []string `pulumi:"displayNamesFilter"`
	Enabled                   bool     `pulumi:"enabled"`
	Kind                      string   `pulumi:"kind"`
	ProductFilter             string   `pulumi:"productFilter"`
	ResourceGroupName         string   `pulumi:"resourceGroupName"`
	RuleId                    *string  `pulumi:"ruleId"`
	SeveritiesFilter          []string `pulumi:"severitiesFilter"`
	WorkspaceName             string   `pulumi:"workspaceName"`
}


type MicrosoftSecurityIncidentCreationAlertRuleArgs struct {
	AlertRuleTemplateName     pulumi.StringPtrInput
	Description               pulumi.StringPtrInput
	DisplayName               pulumi.StringInput
	DisplayNamesExcludeFilter pulumi.StringArrayInput
	DisplayNamesFilter        pulumi.StringArrayInput
	Enabled                   pulumi.BoolInput
	Kind                      pulumi.StringInput
	ProductFilter             pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
	RuleId                    pulumi.StringPtrInput
	SeveritiesFilter          pulumi.StringArrayInput
	WorkspaceName             pulumi.StringInput
}

func (MicrosoftSecurityIncidentCreationAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftSecurityIncidentCreationAlertRuleArgs)(nil)).Elem()
}

type MicrosoftSecurityIncidentCreationAlertRuleInput interface {
	pulumi.Input

	ToMicrosoftSecurityIncidentCreationAlertRuleOutput() MicrosoftSecurityIncidentCreationAlertRuleOutput
	ToMicrosoftSecurityIncidentCreationAlertRuleOutputWithContext(ctx context.Context) MicrosoftSecurityIncidentCreationAlertRuleOutput
}

func (*MicrosoftSecurityIncidentCreationAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftSecurityIncidentCreationAlertRule)(nil)).Elem()
}

func (i *MicrosoftSecurityIncidentCreationAlertRule) ToMicrosoftSecurityIncidentCreationAlertRuleOutput() MicrosoftSecurityIncidentCreationAlertRuleOutput {
	return i.ToMicrosoftSecurityIncidentCreationAlertRuleOutputWithContext(context.Background())
}

func (i *MicrosoftSecurityIncidentCreationAlertRule) ToMicrosoftSecurityIncidentCreationAlertRuleOutputWithContext(ctx context.Context) MicrosoftSecurityIncidentCreationAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MicrosoftSecurityIncidentCreationAlertRuleOutput)
}

type MicrosoftSecurityIncidentCreationAlertRuleOutput struct{ *pulumi.OutputState }

func (MicrosoftSecurityIncidentCreationAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftSecurityIncidentCreationAlertRule)(nil)).Elem()
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) ToMicrosoftSecurityIncidentCreationAlertRuleOutput() MicrosoftSecurityIncidentCreationAlertRuleOutput {
	return o
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) ToMicrosoftSecurityIncidentCreationAlertRuleOutputWithContext(ctx context.Context) MicrosoftSecurityIncidentCreationAlertRuleOutput {
	return o
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringPtrOutput {
		return v.AlertRuleTemplateName
	}).(pulumi.StringPtrOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) DisplayNamesExcludeFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringArrayOutput {
		return v.DisplayNamesExcludeFilter
	}).(pulumi.StringArrayOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) DisplayNamesFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringArrayOutput {
		return v.DisplayNamesFilter
	}).(pulumi.StringArrayOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) ProductFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.ProductFilter }).(pulumi.StringOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) SeveritiesFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringArrayOutput {
		return v.SeveritiesFilter
	}).(pulumi.StringArrayOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MicrosoftSecurityIncidentCreationAlertRuleOutput{})
}
