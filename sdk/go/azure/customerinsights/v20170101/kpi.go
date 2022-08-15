


package v20170101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Kpi struct {
	pulumi.CustomResourceState

	Aliases                     KpiAliasResponseArrayOutput                       `pulumi:"aliases"`
	CalculationWindow           pulumi.StringOutput                               `pulumi:"calculationWindow"`
	CalculationWindowFieldName  pulumi.StringPtrOutput                            `pulumi:"calculationWindowFieldName"`
	Description                 pulumi.StringMapOutput                            `pulumi:"description"`
	DisplayName                 pulumi.StringMapOutput                            `pulumi:"displayName"`
	EntityType                  pulumi.StringOutput                               `pulumi:"entityType"`
	EntityTypeName              pulumi.StringOutput                               `pulumi:"entityTypeName"`
	Expression                  pulumi.StringOutput                               `pulumi:"expression"`
	Extracts                    KpiExtractResponseArrayOutput                     `pulumi:"extracts"`
	Filter                      pulumi.StringPtrOutput                            `pulumi:"filter"`
	Function                    pulumi.StringOutput                               `pulumi:"function"`
	GroupBy                     pulumi.StringArrayOutput                          `pulumi:"groupBy"`
	GroupByMetadata             KpiGroupByMetadataResponseArrayOutput             `pulumi:"groupByMetadata"`
	KpiName                     pulumi.StringOutput                               `pulumi:"kpiName"`
	Name                        pulumi.StringOutput                               `pulumi:"name"`
	ParticipantProfilesMetadata KpiParticipantProfilesMetadataResponseArrayOutput `pulumi:"participantProfilesMetadata"`
	ProvisioningState           pulumi.StringOutput                               `pulumi:"provisioningState"`
	TenantId                    pulumi.StringOutput                               `pulumi:"tenantId"`
	ThresHolds                  KpiThresholdsResponsePtrOutput                    `pulumi:"thresHolds"`
	Type                        pulumi.StringOutput                               `pulumi:"type"`
	Unit                        pulumi.StringPtrOutput                            `pulumi:"unit"`
}


func NewKpi(ctx *pulumi.Context,
	name string, args *KpiArgs, opts ...pulumi.ResourceOption) (*Kpi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CalculationWindow == nil {
		return nil, errors.New("invalid value for required argument 'CalculationWindow'")
	}
	if args.EntityType == nil {
		return nil, errors.New("invalid value for required argument 'EntityType'")
	}
	if args.EntityTypeName == nil {
		return nil, errors.New("invalid value for required argument 'EntityTypeName'")
	}
	if args.Expression == nil {
		return nil, errors.New("invalid value for required argument 'Expression'")
	}
	if args.Function == nil {
		return nil, errors.New("invalid value for required argument 'Function'")
	}
	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights:Kpi"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:Kpi"),
		},
	})
	opts = append(opts, aliases)
	var resource Kpi
	err := ctx.RegisterResource("azure-native:customerinsights/v20170101:Kpi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKpi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KpiState, opts ...pulumi.ResourceOption) (*Kpi, error) {
	var resource Kpi
	err := ctx.ReadResource("azure-native:customerinsights/v20170101:Kpi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kpiState struct {
}

type KpiState struct {
}

func (KpiState) ElementType() reflect.Type {
	return reflect.TypeOf((*kpiState)(nil)).Elem()
}

type kpiArgs struct {
	Aliases                    []KpiAlias             `pulumi:"aliases"`
	CalculationWindow          CalculationWindowTypes `pulumi:"calculationWindow"`
	CalculationWindowFieldName *string                `pulumi:"calculationWindowFieldName"`
	Description                map[string]string      `pulumi:"description"`
	DisplayName                map[string]string      `pulumi:"displayName"`
	EntityType                 EntityTypes            `pulumi:"entityType"`
	EntityTypeName             string                 `pulumi:"entityTypeName"`
	Expression                 string                 `pulumi:"expression"`
	Extracts                   []KpiExtract           `pulumi:"extracts"`
	Filter                     *string                `pulumi:"filter"`
	Function                   KpiFunctions           `pulumi:"function"`
	GroupBy                    []string               `pulumi:"groupBy"`
	HubName                    string                 `pulumi:"hubName"`
	KpiName                    *string                `pulumi:"kpiName"`
	ResourceGroupName          string                 `pulumi:"resourceGroupName"`
	ThresHolds                 *KpiThresholds         `pulumi:"thresHolds"`
	Unit                       *string                `pulumi:"unit"`
}


type KpiArgs struct {
	Aliases                    KpiAliasArrayInput
	CalculationWindow          CalculationWindowTypesInput
	CalculationWindowFieldName pulumi.StringPtrInput
	Description                pulumi.StringMapInput
	DisplayName                pulumi.StringMapInput
	EntityType                 EntityTypesInput
	EntityTypeName             pulumi.StringInput
	Expression                 pulumi.StringInput
	Extracts                   KpiExtractArrayInput
	Filter                     pulumi.StringPtrInput
	Function                   KpiFunctionsInput
	GroupBy                    pulumi.StringArrayInput
	HubName                    pulumi.StringInput
	KpiName                    pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	ThresHolds                 KpiThresholdsPtrInput
	Unit                       pulumi.StringPtrInput
}

func (KpiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kpiArgs)(nil)).Elem()
}

type KpiInput interface {
	pulumi.Input

	ToKpiOutput() KpiOutput
	ToKpiOutputWithContext(ctx context.Context) KpiOutput
}

func (*Kpi) ElementType() reflect.Type {
	return reflect.TypeOf((**Kpi)(nil)).Elem()
}

func (i *Kpi) ToKpiOutput() KpiOutput {
	return i.ToKpiOutputWithContext(context.Background())
}

func (i *Kpi) ToKpiOutputWithContext(ctx context.Context) KpiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiOutput)
}

type KpiOutput struct{ *pulumi.OutputState }

func (KpiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kpi)(nil)).Elem()
}

func (o KpiOutput) ToKpiOutput() KpiOutput {
	return o
}

func (o KpiOutput) ToKpiOutputWithContext(ctx context.Context) KpiOutput {
	return o
}

func (o KpiOutput) Aliases() KpiAliasResponseArrayOutput {
	return o.ApplyT(func(v *Kpi) KpiAliasResponseArrayOutput { return v.Aliases }).(KpiAliasResponseArrayOutput)
}

func (o KpiOutput) CalculationWindow() pulumi.StringOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringOutput { return v.CalculationWindow }).(pulumi.StringOutput)
}

func (o KpiOutput) CalculationWindowFieldName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringPtrOutput { return v.CalculationWindowFieldName }).(pulumi.StringPtrOutput)
}

func (o KpiOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringMapOutput { return v.Description }).(pulumi.StringMapOutput)
}

func (o KpiOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringMapOutput { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o KpiOutput) EntityType() pulumi.StringOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringOutput { return v.EntityType }).(pulumi.StringOutput)
}

func (o KpiOutput) EntityTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringOutput { return v.EntityTypeName }).(pulumi.StringOutput)
}

func (o KpiOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringOutput { return v.Expression }).(pulumi.StringOutput)
}

func (o KpiOutput) Extracts() KpiExtractResponseArrayOutput {
	return o.ApplyT(func(v *Kpi) KpiExtractResponseArrayOutput { return v.Extracts }).(KpiExtractResponseArrayOutput)
}

func (o KpiOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringPtrOutput { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o KpiOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringOutput { return v.Function }).(pulumi.StringOutput)
}

func (o KpiOutput) GroupBy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringArrayOutput { return v.GroupBy }).(pulumi.StringArrayOutput)
}

func (o KpiOutput) GroupByMetadata() KpiGroupByMetadataResponseArrayOutput {
	return o.ApplyT(func(v *Kpi) KpiGroupByMetadataResponseArrayOutput { return v.GroupByMetadata }).(KpiGroupByMetadataResponseArrayOutput)
}

func (o KpiOutput) KpiName() pulumi.StringOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringOutput { return v.KpiName }).(pulumi.StringOutput)
}

func (o KpiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KpiOutput) ParticipantProfilesMetadata() KpiParticipantProfilesMetadataResponseArrayOutput {
	return o.ApplyT(func(v *Kpi) KpiParticipantProfilesMetadataResponseArrayOutput { return v.ParticipantProfilesMetadata }).(KpiParticipantProfilesMetadataResponseArrayOutput)
}

func (o KpiOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o KpiOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o KpiOutput) ThresHolds() KpiThresholdsResponsePtrOutput {
	return o.ApplyT(func(v *Kpi) KpiThresholdsResponsePtrOutput { return v.ThresHolds }).(KpiThresholdsResponsePtrOutput)
}

func (o KpiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o KpiOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kpi) pulumi.StringPtrOutput { return v.Unit }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(KpiOutput{})
}
