


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ThreatIntelligenceIndicator struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewThreatIntelligenceIndicator(ctx *pulumi.Context,
	name string, args *ThreatIntelligenceIndicatorArgs, opts ...pulumi.ResourceOption) (*ThreatIntelligenceIndicator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("indicator")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:ThreatIntelligenceIndicator"),
		},
	})
	opts = append(opts, aliases)
	var resource ThreatIntelligenceIndicator
	err := ctx.RegisterResource("azure-native:securityinsights/v20220901preview:ThreatIntelligenceIndicator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetThreatIntelligenceIndicator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThreatIntelligenceIndicatorState, opts ...pulumi.ResourceOption) (*ThreatIntelligenceIndicator, error) {
	var resource ThreatIntelligenceIndicator
	err := ctx.ReadResource("azure-native:securityinsights/v20220901preview:ThreatIntelligenceIndicator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type threatIntelligenceIndicatorState struct {
}

type ThreatIntelligenceIndicatorState struct {
}

func (ThreatIntelligenceIndicatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelligenceIndicatorState)(nil)).Elem()
}

type threatIntelligenceIndicatorArgs struct {
	Confidence                 *int                                     `pulumi:"confidence"`
	Created                    *string                                  `pulumi:"created"`
	CreatedByRef               *string                                  `pulumi:"createdByRef"`
	Defanged                   *bool                                    `pulumi:"defanged"`
	Description                *string                                  `pulumi:"description"`
	DisplayName                *string                                  `pulumi:"displayName"`
	Extensions                 interface{}                              `pulumi:"extensions"`
	ExternalId                 *string                                  `pulumi:"externalId"`
	ExternalLastUpdatedTimeUtc *string                                  `pulumi:"externalLastUpdatedTimeUtc"`
	ExternalReferences         []ThreatIntelligenceExternalReference    `pulumi:"externalReferences"`
	GranularMarkings           []ThreatIntelligenceGranularMarkingModel `pulumi:"granularMarkings"`
	IndicatorTypes             []string                                 `pulumi:"indicatorTypes"`
	KillChainPhases            []ThreatIntelligenceKillChainPhase       `pulumi:"killChainPhases"`
	Kind                       string                                   `pulumi:"kind"`
	Labels                     []string                                 `pulumi:"labels"`
	Language                   *string                                  `pulumi:"language"`
	LastUpdatedTimeUtc         *string                                  `pulumi:"lastUpdatedTimeUtc"`
	Modified                   *string                                  `pulumi:"modified"`
	Name                       *string                                  `pulumi:"name"`
	ObjectMarkingRefs          []string                                 `pulumi:"objectMarkingRefs"`
	ParsedPattern              []ThreatIntelligenceParsedPattern        `pulumi:"parsedPattern"`
	Pattern                    *string                                  `pulumi:"pattern"`
	PatternType                *string                                  `pulumi:"patternType"`
	PatternVersion             *string                                  `pulumi:"patternVersion"`
	ResourceGroupName          string                                   `pulumi:"resourceGroupName"`
	Revoked                    *bool                                    `pulumi:"revoked"`
	Source                     *string                                  `pulumi:"source"`
	ThreatIntelligenceTags     []string                                 `pulumi:"threatIntelligenceTags"`
	ThreatTypes                []string                                 `pulumi:"threatTypes"`
	ValidFrom                  *string                                  `pulumi:"validFrom"`
	ValidUntil                 *string                                  `pulumi:"validUntil"`
	WorkspaceName              string                                   `pulumi:"workspaceName"`
}


type ThreatIntelligenceIndicatorArgs struct {
	Confidence                 pulumi.IntPtrInput
	Created                    pulumi.StringPtrInput
	CreatedByRef               pulumi.StringPtrInput
	Defanged                   pulumi.BoolPtrInput
	Description                pulumi.StringPtrInput
	DisplayName                pulumi.StringPtrInput
	Extensions                 pulumi.Input
	ExternalId                 pulumi.StringPtrInput
	ExternalLastUpdatedTimeUtc pulumi.StringPtrInput
	ExternalReferences         ThreatIntelligenceExternalReferenceArrayInput
	GranularMarkings           ThreatIntelligenceGranularMarkingModelArrayInput
	IndicatorTypes             pulumi.StringArrayInput
	KillChainPhases            ThreatIntelligenceKillChainPhaseArrayInput
	Kind                       pulumi.StringInput
	Labels                     pulumi.StringArrayInput
	Language                   pulumi.StringPtrInput
	LastUpdatedTimeUtc         pulumi.StringPtrInput
	Modified                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	ObjectMarkingRefs          pulumi.StringArrayInput
	ParsedPattern              ThreatIntelligenceParsedPatternArrayInput
	Pattern                    pulumi.StringPtrInput
	PatternType                pulumi.StringPtrInput
	PatternVersion             pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	Revoked                    pulumi.BoolPtrInput
	Source                     pulumi.StringPtrInput
	ThreatIntelligenceTags     pulumi.StringArrayInput
	ThreatTypes                pulumi.StringArrayInput
	ValidFrom                  pulumi.StringPtrInput
	ValidUntil                 pulumi.StringPtrInput
	WorkspaceName              pulumi.StringInput
}

func (ThreatIntelligenceIndicatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelligenceIndicatorArgs)(nil)).Elem()
}

type ThreatIntelligenceIndicatorInput interface {
	pulumi.Input

	ToThreatIntelligenceIndicatorOutput() ThreatIntelligenceIndicatorOutput
	ToThreatIntelligenceIndicatorOutputWithContext(ctx context.Context) ThreatIntelligenceIndicatorOutput
}

func (*ThreatIntelligenceIndicator) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelligenceIndicator)(nil)).Elem()
}

func (i *ThreatIntelligenceIndicator) ToThreatIntelligenceIndicatorOutput() ThreatIntelligenceIndicatorOutput {
	return i.ToThreatIntelligenceIndicatorOutputWithContext(context.Background())
}

func (i *ThreatIntelligenceIndicator) ToThreatIntelligenceIndicatorOutputWithContext(ctx context.Context) ThreatIntelligenceIndicatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceIndicatorOutput)
}

type ThreatIntelligenceIndicatorOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceIndicatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelligenceIndicator)(nil)).Elem()
}

func (o ThreatIntelligenceIndicatorOutput) ToThreatIntelligenceIndicatorOutput() ThreatIntelligenceIndicatorOutput {
	return o
}

func (o ThreatIntelligenceIndicatorOutput) ToThreatIntelligenceIndicatorOutputWithContext(ctx context.Context) ThreatIntelligenceIndicatorOutput {
	return o
}

func (o ThreatIntelligenceIndicatorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThreatIntelligenceIndicator) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ThreatIntelligenceIndicatorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceIndicator) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ThreatIntelligenceIndicatorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceIndicator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ThreatIntelligenceIndicatorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ThreatIntelligenceIndicator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ThreatIntelligenceIndicatorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceIndicator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ThreatIntelligenceIndicatorOutput{})
}
