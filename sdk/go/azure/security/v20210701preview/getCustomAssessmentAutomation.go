


package v20210701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomAssessmentAutomation(ctx *pulumi.Context, args *LookupCustomAssessmentAutomationArgs, opts ...pulumi.InvokeOption) (*LookupCustomAssessmentAutomationResult, error) {
	var rv LookupCustomAssessmentAutomationResult
	err := ctx.Invoke("azure-native:security/v20210701preview:getCustomAssessmentAutomation", args, &rv, opts...)
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
	CompressedQuery        *string `pulumi:"compressedQuery"`
	Description            *string `pulumi:"description"`
	Id                     string  `pulumi:"id"`
	Name                   string  `pulumi:"name"`
	RemediationDescription *string `pulumi:"remediationDescription"`
	Severity               *string `pulumi:"severity"`
	SupportedCloud         *string `pulumi:"supportedCloud"`
	Type                   string  `pulumi:"type"`
}
