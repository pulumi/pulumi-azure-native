


package v20200101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssessmentMetadataInSubscription struct {
	pulumi.CustomResourceState

	AssessmentType         pulumi.StringOutput                                    `pulumi:"assessmentType"`
	Categories             pulumi.StringArrayOutput                               `pulumi:"categories"`
	Description            pulumi.StringPtrOutput                                 `pulumi:"description"`
	DisplayName            pulumi.StringOutput                                    `pulumi:"displayName"`
	ImplementationEffort   pulumi.StringPtrOutput                                 `pulumi:"implementationEffort"`
	Name                   pulumi.StringOutput                                    `pulumi:"name"`
	PartnerData            SecurityAssessmentMetadataPartnerDataResponsePtrOutput `pulumi:"partnerData"`
	PolicyDefinitionId     pulumi.StringOutput                                    `pulumi:"policyDefinitionId"`
	Preview                pulumi.BoolPtrOutput                                   `pulumi:"preview"`
	RemediationDescription pulumi.StringPtrOutput                                 `pulumi:"remediationDescription"`
	Severity               pulumi.StringOutput                                    `pulumi:"severity"`
	Threats                pulumi.StringArrayOutput                               `pulumi:"threats"`
	Type                   pulumi.StringOutput                                    `pulumi:"type"`
	UserImpact             pulumi.StringPtrOutput                                 `pulumi:"userImpact"`
}


func NewAssessmentMetadataInSubscription(ctx *pulumi.Context,
	name string, args *AssessmentMetadataInSubscriptionArgs, opts ...pulumi.ResourceOption) (*AssessmentMetadataInSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssessmentType == nil {
		return nil, errors.New("invalid value for required argument 'AssessmentType'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:AssessmentMetadataInSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security/v20190101preview:AssessmentMetadataInSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security/v20210601:AssessmentMetadataInSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource AssessmentMetadataInSubscription
	err := ctx.RegisterResource("azure-native:security/v20200101:AssessmentMetadataInSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAssessmentMetadataInSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentMetadataInSubscriptionState, opts ...pulumi.ResourceOption) (*AssessmentMetadataInSubscription, error) {
	var resource AssessmentMetadataInSubscription
	err := ctx.ReadResource("azure-native:security/v20200101:AssessmentMetadataInSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type assessmentMetadataInSubscriptionState struct {
}

type AssessmentMetadataInSubscriptionState struct {
}

func (AssessmentMetadataInSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentMetadataInSubscriptionState)(nil)).Elem()
}

type assessmentMetadataInSubscriptionArgs struct {
	AssessmentMetadataName *string                                `pulumi:"assessmentMetadataName"`
	AssessmentType         string                                 `pulumi:"assessmentType"`
	Categories             []string                               `pulumi:"categories"`
	Description            *string                                `pulumi:"description"`
	DisplayName            string                                 `pulumi:"displayName"`
	ImplementationEffort   *string                                `pulumi:"implementationEffort"`
	PartnerData            *SecurityAssessmentMetadataPartnerData `pulumi:"partnerData"`
	Preview                *bool                                  `pulumi:"preview"`
	RemediationDescription *string                                `pulumi:"remediationDescription"`
	Severity               string                                 `pulumi:"severity"`
	Threats                []string                               `pulumi:"threats"`
	UserImpact             *string                                `pulumi:"userImpact"`
}


type AssessmentMetadataInSubscriptionArgs struct {
	AssessmentMetadataName pulumi.StringPtrInput
	AssessmentType         pulumi.StringInput
	Categories             pulumi.StringArrayInput
	Description            pulumi.StringPtrInput
	DisplayName            pulumi.StringInput
	ImplementationEffort   pulumi.StringPtrInput
	PartnerData            SecurityAssessmentMetadataPartnerDataPtrInput
	Preview                pulumi.BoolPtrInput
	RemediationDescription pulumi.StringPtrInput
	Severity               pulumi.StringInput
	Threats                pulumi.StringArrayInput
	UserImpact             pulumi.StringPtrInput
}

func (AssessmentMetadataInSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentMetadataInSubscriptionArgs)(nil)).Elem()
}

type AssessmentMetadataInSubscriptionInput interface {
	pulumi.Input

	ToAssessmentMetadataInSubscriptionOutput() AssessmentMetadataInSubscriptionOutput
	ToAssessmentMetadataInSubscriptionOutputWithContext(ctx context.Context) AssessmentMetadataInSubscriptionOutput
}

func (*AssessmentMetadataInSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentMetadataInSubscription)(nil)).Elem()
}

func (i *AssessmentMetadataInSubscription) ToAssessmentMetadataInSubscriptionOutput() AssessmentMetadataInSubscriptionOutput {
	return i.ToAssessmentMetadataInSubscriptionOutputWithContext(context.Background())
}

func (i *AssessmentMetadataInSubscription) ToAssessmentMetadataInSubscriptionOutputWithContext(ctx context.Context) AssessmentMetadataInSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentMetadataInSubscriptionOutput)
}

type AssessmentMetadataInSubscriptionOutput struct{ *pulumi.OutputState }

func (AssessmentMetadataInSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentMetadataInSubscription)(nil)).Elem()
}

func (o AssessmentMetadataInSubscriptionOutput) ToAssessmentMetadataInSubscriptionOutput() AssessmentMetadataInSubscriptionOutput {
	return o
}

func (o AssessmentMetadataInSubscriptionOutput) ToAssessmentMetadataInSubscriptionOutputWithContext(ctx context.Context) AssessmentMetadataInSubscriptionOutput {
	return o
}

func (o AssessmentMetadataInSubscriptionOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringArrayOutput { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringPtrOutput { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) PartnerData() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.BoolPtrOutput { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringPtrOutput { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringArrayOutput { return v.Threats }).(pulumi.StringArrayOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringPtrOutput { return v.UserImpact }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentMetadataInSubscriptionOutput{})
}
