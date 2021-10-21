


package security

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomAssessmentAutomation(ctx *pulumi.Context, args *LookupCustomAssessmentAutomationArgs, opts ...pulumi.InvokeOption) (*LookupCustomAssessmentAutomationResult, error) {
	var rv LookupCustomAssessmentAutomationResult
	err := ctx.Invoke("azure-native:security:getCustomAssessmentAutomation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomAssessmentAutomationArgs struct {
	CustomAssessmentAutomationName string `pulumi:"customAssessmentAutomationName"`
	ResourceGroupName              string `pulumi:"resourceGroupName"`
}


type LookupCustomAssessmentAutomationResult struct {
	AssessmentKey          *string            `pulumi:"assessmentKey"`
	CompressedQuery        *string            `pulumi:"compressedQuery"`
	Description            *string            `pulumi:"description"`
	Id                     string             `pulumi:"id"`
	ImplementationEffort   *string            `pulumi:"implementationEffort"`
	Name                   string             `pulumi:"name"`
	RemediationDescription *string            `pulumi:"remediationDescription"`
	Severity               *string            `pulumi:"severity"`
	SupportedCloud         *string            `pulumi:"supportedCloud"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
	Type                   string             `pulumi:"type"`
	UserImpact             *string            `pulumi:"userImpact"`
}
