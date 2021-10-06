


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssessmentsMetadataSubscription(ctx *pulumi.Context, args *LookupAssessmentsMetadataSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentsMetadataSubscriptionResult, error) {
	var rv LookupAssessmentsMetadataSubscriptionResult
	err := ctx.Invoke("azure-native:security/v20190101preview:getAssessmentsMetadataSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentsMetadataSubscriptionArgs struct {
	AssessmentMetadataName string `pulumi:"assessmentMetadataName"`
}


type LookupAssessmentsMetadataSubscriptionResult struct {
	AssessmentType         string   `pulumi:"assessmentType"`
	Categories             []string `pulumi:"categories"`
	Description            *string  `pulumi:"description"`
	DisplayName            string   `pulumi:"displayName"`
	Id                     string   `pulumi:"id"`
	ImplementationEffort   *string  `pulumi:"implementationEffort"`
	Name                   string   `pulumi:"name"`
	PolicyDefinitionId     string   `pulumi:"policyDefinitionId"`
	Preview                *bool    `pulumi:"preview"`
	RemediationDescription *string  `pulumi:"remediationDescription"`
	Severity               string   `pulumi:"severity"`
	Threats                []string `pulumi:"threats"`
	Type                   string   `pulumi:"type"`
	UserImpact             *string  `pulumi:"userImpact"`
}
