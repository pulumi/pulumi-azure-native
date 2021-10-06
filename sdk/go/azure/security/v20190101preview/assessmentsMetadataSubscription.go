


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssessmentsMetadataSubscription struct {
	pulumi.CustomResourceState

	AssessmentType         pulumi.StringOutput      `pulumi:"assessmentType"`
	Categories             pulumi.StringArrayOutput `pulumi:"categories"`
	Description            pulumi.StringPtrOutput   `pulumi:"description"`
	DisplayName            pulumi.StringOutput      `pulumi:"displayName"`
	ImplementationEffort   pulumi.StringPtrOutput   `pulumi:"implementationEffort"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	PolicyDefinitionId     pulumi.StringOutput      `pulumi:"policyDefinitionId"`
	Preview                pulumi.BoolPtrOutput     `pulumi:"preview"`
	RemediationDescription pulumi.StringPtrOutput   `pulumi:"remediationDescription"`
	Severity               pulumi.StringOutput      `pulumi:"severity"`
	Threats                pulumi.StringArrayOutput `pulumi:"threats"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
	UserImpact             pulumi.StringPtrOutput   `pulumi:"userImpact"`
}


func NewAssessmentsMetadataSubscription(ctx *pulumi.Context,
	name string, args *AssessmentsMetadataSubscriptionArgs, opts ...pulumi.ResourceOption) (*AssessmentsMetadataSubscription, error) {
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
			Type: pulumi.String("azure-nextgen:security/v20190101preview:AssessmentsMetadataSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security:AssessmentsMetadataSubscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:security:AssessmentsMetadataSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101:AssessmentsMetadataSubscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20200101:AssessmentsMetadataSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security/v20210601:AssessmentsMetadataSubscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20210601:AssessmentsMetadataSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource AssessmentsMetadataSubscription
	err := ctx.RegisterResource("azure-native:security/v20190101preview:AssessmentsMetadataSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAssessmentsMetadataSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentsMetadataSubscriptionState, opts ...pulumi.ResourceOption) (*AssessmentsMetadataSubscription, error) {
	var resource AssessmentsMetadataSubscription
	err := ctx.ReadResource("azure-native:security/v20190101preview:AssessmentsMetadataSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type assessmentsMetadataSubscriptionState struct {
}

type AssessmentsMetadataSubscriptionState struct {
}

func (AssessmentsMetadataSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentsMetadataSubscriptionState)(nil)).Elem()
}

type assessmentsMetadataSubscriptionArgs struct {
	AssessmentMetadataName *string  `pulumi:"assessmentMetadataName"`
	AssessmentType         string   `pulumi:"assessmentType"`
	Categories             []string `pulumi:"categories"`
	Description            *string  `pulumi:"description"`
	DisplayName            string   `pulumi:"displayName"`
	ImplementationEffort   *string  `pulumi:"implementationEffort"`
	Preview                *bool    `pulumi:"preview"`
	RemediationDescription *string  `pulumi:"remediationDescription"`
	Severity               string   `pulumi:"severity"`
	Threats                []string `pulumi:"threats"`
	UserImpact             *string  `pulumi:"userImpact"`
}


type AssessmentsMetadataSubscriptionArgs struct {
	AssessmentMetadataName pulumi.StringPtrInput
	AssessmentType         pulumi.StringInput
	Categories             pulumi.StringArrayInput
	Description            pulumi.StringPtrInput
	DisplayName            pulumi.StringInput
	ImplementationEffort   pulumi.StringPtrInput
	Preview                pulumi.BoolPtrInput
	RemediationDescription pulumi.StringPtrInput
	Severity               pulumi.StringInput
	Threats                pulumi.StringArrayInput
	UserImpact             pulumi.StringPtrInput
}

func (AssessmentsMetadataSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentsMetadataSubscriptionArgs)(nil)).Elem()
}

type AssessmentsMetadataSubscriptionInput interface {
	pulumi.Input

	ToAssessmentsMetadataSubscriptionOutput() AssessmentsMetadataSubscriptionOutput
	ToAssessmentsMetadataSubscriptionOutputWithContext(ctx context.Context) AssessmentsMetadataSubscriptionOutput
}

func (*AssessmentsMetadataSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentsMetadataSubscription)(nil))
}

func (i *AssessmentsMetadataSubscription) ToAssessmentsMetadataSubscriptionOutput() AssessmentsMetadataSubscriptionOutput {
	return i.ToAssessmentsMetadataSubscriptionOutputWithContext(context.Background())
}

func (i *AssessmentsMetadataSubscription) ToAssessmentsMetadataSubscriptionOutputWithContext(ctx context.Context) AssessmentsMetadataSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentsMetadataSubscriptionOutput)
}

type AssessmentsMetadataSubscriptionOutput struct{ *pulumi.OutputState }

func (AssessmentsMetadataSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentsMetadataSubscription)(nil))
}

func (o AssessmentsMetadataSubscriptionOutput) ToAssessmentsMetadataSubscriptionOutput() AssessmentsMetadataSubscriptionOutput {
	return o
}

func (o AssessmentsMetadataSubscriptionOutput) ToAssessmentsMetadataSubscriptionOutputWithContext(ctx context.Context) AssessmentsMetadataSubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AssessmentsMetadataSubscriptionOutput{})
}
