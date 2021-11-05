


package security

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomAssessmentAutomation struct {
	pulumi.CustomResourceState

	AssessmentKey          pulumi.StringPtrOutput   `pulumi:"assessmentKey"`
	CompressedQuery        pulumi.StringPtrOutput   `pulumi:"compressedQuery"`
	Description            pulumi.StringPtrOutput   `pulumi:"description"`
	DisplayName            pulumi.StringPtrOutput   `pulumi:"displayName"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	RemediationDescription pulumi.StringPtrOutput   `pulumi:"remediationDescription"`
	Severity               pulumi.StringPtrOutput   `pulumi:"severity"`
	SupportedCloud         pulumi.StringPtrOutput   `pulumi:"supportedCloud"`
	SystemData             SystemDataResponseOutput `pulumi:"systemData"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
}


func NewCustomAssessmentAutomation(ctx *pulumi.Context,
	name string, args *CustomAssessmentAutomationArgs, opts ...pulumi.ResourceOption) (*CustomAssessmentAutomation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security:CustomAssessmentAutomation"),
		},
		{
			Type: pulumi.String("azure-native:security/v20210701preview:CustomAssessmentAutomation"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20210701preview:CustomAssessmentAutomation"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomAssessmentAutomation
	err := ctx.RegisterResource("azure-native:security:CustomAssessmentAutomation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomAssessmentAutomation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomAssessmentAutomationState, opts ...pulumi.ResourceOption) (*CustomAssessmentAutomation, error) {
	var resource CustomAssessmentAutomation
	err := ctx.ReadResource("azure-native:security:CustomAssessmentAutomation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customAssessmentAutomationState struct {
}

type CustomAssessmentAutomationState struct {
}

func (CustomAssessmentAutomationState) ElementType() reflect.Type {
	return reflect.TypeOf((*customAssessmentAutomationState)(nil)).Elem()
}

type customAssessmentAutomationArgs struct {
	CompressedQuery                *string `pulumi:"compressedQuery"`
	CustomAssessmentAutomationName *string `pulumi:"customAssessmentAutomationName"`
	Description                    *string `pulumi:"description"`
	DisplayName                    *string `pulumi:"displayName"`
	RemediationDescription         *string `pulumi:"remediationDescription"`
	ResourceGroupName              string  `pulumi:"resourceGroupName"`
	Severity                       *string `pulumi:"severity"`
	SupportedCloud                 *string `pulumi:"supportedCloud"`
}


type CustomAssessmentAutomationArgs struct {
	CompressedQuery                pulumi.StringPtrInput
	CustomAssessmentAutomationName pulumi.StringPtrInput
	Description                    pulumi.StringPtrInput
	DisplayName                    pulumi.StringPtrInput
	RemediationDescription         pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	Severity                       pulumi.StringPtrInput
	SupportedCloud                 pulumi.StringPtrInput
}

func (CustomAssessmentAutomationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customAssessmentAutomationArgs)(nil)).Elem()
}

type CustomAssessmentAutomationInput interface {
	pulumi.Input

	ToCustomAssessmentAutomationOutput() CustomAssessmentAutomationOutput
	ToCustomAssessmentAutomationOutputWithContext(ctx context.Context) CustomAssessmentAutomationOutput
}

func (*CustomAssessmentAutomation) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomAssessmentAutomation)(nil))
}

func (i *CustomAssessmentAutomation) ToCustomAssessmentAutomationOutput() CustomAssessmentAutomationOutput {
	return i.ToCustomAssessmentAutomationOutputWithContext(context.Background())
}

func (i *CustomAssessmentAutomation) ToCustomAssessmentAutomationOutputWithContext(ctx context.Context) CustomAssessmentAutomationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomAssessmentAutomationOutput)
}

type CustomAssessmentAutomationOutput struct{ *pulumi.OutputState }

func (CustomAssessmentAutomationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomAssessmentAutomation)(nil))
}

func (o CustomAssessmentAutomationOutput) ToCustomAssessmentAutomationOutput() CustomAssessmentAutomationOutput {
	return o
}

func (o CustomAssessmentAutomationOutput) ToCustomAssessmentAutomationOutputWithContext(ctx context.Context) CustomAssessmentAutomationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomAssessmentAutomationOutput{})
}
