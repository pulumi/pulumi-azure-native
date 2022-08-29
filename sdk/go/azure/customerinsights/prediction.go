


package customerinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Prediction struct {
	pulumi.CustomResourceState

	AutoAnalyze               pulumi.BoolOutput                               `pulumi:"autoAnalyze"`
	Description               pulumi.StringMapOutput                          `pulumi:"description"`
	DisplayName               pulumi.StringMapOutput                          `pulumi:"displayName"`
	Grades                    PredictionResponseGradesArrayOutput             `pulumi:"grades"`
	InvolvedInteractionTypes  pulumi.StringArrayOutput                        `pulumi:"involvedInteractionTypes"`
	InvolvedKpiTypes          pulumi.StringArrayOutput                        `pulumi:"involvedKpiTypes"`
	InvolvedRelationships     pulumi.StringArrayOutput                        `pulumi:"involvedRelationships"`
	Mappings                  PredictionResponseMappingsOutput                `pulumi:"mappings"`
	Name                      pulumi.StringOutput                             `pulumi:"name"`
	NegativeOutcomeExpression pulumi.StringOutput                             `pulumi:"negativeOutcomeExpression"`
	PositiveOutcomeExpression pulumi.StringOutput                             `pulumi:"positiveOutcomeExpression"`
	PredictionName            pulumi.StringPtrOutput                          `pulumi:"predictionName"`
	PrimaryProfileType        pulumi.StringOutput                             `pulumi:"primaryProfileType"`
	ProvisioningState         pulumi.StringOutput                             `pulumi:"provisioningState"`
	ScopeExpression           pulumi.StringOutput                             `pulumi:"scopeExpression"`
	ScoreLabel                pulumi.StringOutput                             `pulumi:"scoreLabel"`
	SystemGeneratedEntities   PredictionResponseSystemGeneratedEntitiesOutput `pulumi:"systemGeneratedEntities"`
	TenantId                  pulumi.StringOutput                             `pulumi:"tenantId"`
	Type                      pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrediction(ctx *pulumi.Context,
	name string, args *PredictionArgs, opts ...pulumi.ResourceOption) (*Prediction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoAnalyze == nil {
		return nil, errors.New("invalid value for required argument 'AutoAnalyze'")
	}
	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.Mappings == nil {
		return nil, errors.New("invalid value for required argument 'Mappings'")
	}
	if args.NegativeOutcomeExpression == nil {
		return nil, errors.New("invalid value for required argument 'NegativeOutcomeExpression'")
	}
	if args.PositiveOutcomeExpression == nil {
		return nil, errors.New("invalid value for required argument 'PositiveOutcomeExpression'")
	}
	if args.PrimaryProfileType == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryProfileType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScopeExpression == nil {
		return nil, errors.New("invalid value for required argument 'ScopeExpression'")
	}
	if args.ScoreLabel == nil {
		return nil, errors.New("invalid value for required argument 'ScoreLabel'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:Prediction"),
		},
	})
	opts = append(opts, aliases)
	var resource Prediction
	err := ctx.RegisterResource("azure-native:customerinsights:Prediction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrediction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PredictionState, opts ...pulumi.ResourceOption) (*Prediction, error) {
	var resource Prediction
	err := ctx.ReadResource("azure-native:customerinsights:Prediction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type predictionState struct {
}

type PredictionState struct {
}

func (PredictionState) ElementType() reflect.Type {
	return reflect.TypeOf((*predictionState)(nil)).Elem()
}

type predictionArgs struct {
	AutoAnalyze               bool               `pulumi:"autoAnalyze"`
	Description               map[string]string  `pulumi:"description"`
	DisplayName               map[string]string  `pulumi:"displayName"`
	Grades                    []PredictionGrades `pulumi:"grades"`
	HubName                   string             `pulumi:"hubName"`
	InvolvedInteractionTypes  []string           `pulumi:"involvedInteractionTypes"`
	InvolvedKpiTypes          []string           `pulumi:"involvedKpiTypes"`
	InvolvedRelationships     []string           `pulumi:"involvedRelationships"`
	Mappings                  PredictionMappings `pulumi:"mappings"`
	NegativeOutcomeExpression string             `pulumi:"negativeOutcomeExpression"`
	PositiveOutcomeExpression string             `pulumi:"positiveOutcomeExpression"`
	PredictionName            *string            `pulumi:"predictionName"`
	PrimaryProfileType        string             `pulumi:"primaryProfileType"`
	ResourceGroupName         string             `pulumi:"resourceGroupName"`
	ScopeExpression           string             `pulumi:"scopeExpression"`
	ScoreLabel                string             `pulumi:"scoreLabel"`
}


type PredictionArgs struct {
	AutoAnalyze               pulumi.BoolInput
	Description               pulumi.StringMapInput
	DisplayName               pulumi.StringMapInput
	Grades                    PredictionGradesArrayInput
	HubName                   pulumi.StringInput
	InvolvedInteractionTypes  pulumi.StringArrayInput
	InvolvedKpiTypes          pulumi.StringArrayInput
	InvolvedRelationships     pulumi.StringArrayInput
	Mappings                  PredictionMappingsInput
	NegativeOutcomeExpression pulumi.StringInput
	PositiveOutcomeExpression pulumi.StringInput
	PredictionName            pulumi.StringPtrInput
	PrimaryProfileType        pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
	ScopeExpression           pulumi.StringInput
	ScoreLabel                pulumi.StringInput
}

func (PredictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*predictionArgs)(nil)).Elem()
}

type PredictionInput interface {
	pulumi.Input

	ToPredictionOutput() PredictionOutput
	ToPredictionOutputWithContext(ctx context.Context) PredictionOutput
}

func (*Prediction) ElementType() reflect.Type {
	return reflect.TypeOf((**Prediction)(nil)).Elem()
}

func (i *Prediction) ToPredictionOutput() PredictionOutput {
	return i.ToPredictionOutputWithContext(context.Background())
}

func (i *Prediction) ToPredictionOutputWithContext(ctx context.Context) PredictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PredictionOutput)
}

type PredictionOutput struct{ *pulumi.OutputState }

func (PredictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Prediction)(nil)).Elem()
}

func (o PredictionOutput) ToPredictionOutput() PredictionOutput {
	return o
}

func (o PredictionOutput) ToPredictionOutputWithContext(ctx context.Context) PredictionOutput {
	return o
}

func (o PredictionOutput) AutoAnalyze() pulumi.BoolOutput {
	return o.ApplyT(func(v *Prediction) pulumi.BoolOutput { return v.AutoAnalyze }).(pulumi.BoolOutput)
}

func (o PredictionOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringMapOutput { return v.Description }).(pulumi.StringMapOutput)
}

func (o PredictionOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringMapOutput { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o PredictionOutput) Grades() PredictionResponseGradesArrayOutput {
	return o.ApplyT(func(v *Prediction) PredictionResponseGradesArrayOutput { return v.Grades }).(PredictionResponseGradesArrayOutput)
}

func (o PredictionOutput) InvolvedInteractionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringArrayOutput { return v.InvolvedInteractionTypes }).(pulumi.StringArrayOutput)
}

func (o PredictionOutput) InvolvedKpiTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringArrayOutput { return v.InvolvedKpiTypes }).(pulumi.StringArrayOutput)
}

func (o PredictionOutput) InvolvedRelationships() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringArrayOutput { return v.InvolvedRelationships }).(pulumi.StringArrayOutput)
}

func (o PredictionOutput) Mappings() PredictionResponseMappingsOutput {
	return o.ApplyT(func(v *Prediction) PredictionResponseMappingsOutput { return v.Mappings }).(PredictionResponseMappingsOutput)
}

func (o PredictionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PredictionOutput) NegativeOutcomeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.NegativeOutcomeExpression }).(pulumi.StringOutput)
}

func (o PredictionOutput) PositiveOutcomeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.PositiveOutcomeExpression }).(pulumi.StringOutput)
}

func (o PredictionOutput) PredictionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringPtrOutput { return v.PredictionName }).(pulumi.StringPtrOutput)
}

func (o PredictionOutput) PrimaryProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.PrimaryProfileType }).(pulumi.StringOutput)
}

func (o PredictionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PredictionOutput) ScopeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.ScopeExpression }).(pulumi.StringOutput)
}

func (o PredictionOutput) ScoreLabel() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.ScoreLabel }).(pulumi.StringOutput)
}

func (o PredictionOutput) SystemGeneratedEntities() PredictionResponseSystemGeneratedEntitiesOutput {
	return o.ApplyT(func(v *Prediction) PredictionResponseSystemGeneratedEntitiesOutput { return v.SystemGeneratedEntities }).(PredictionResponseSystemGeneratedEntitiesOutput)
}

func (o PredictionOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o PredictionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PredictionOutput{})
}
