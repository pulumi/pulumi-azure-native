


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssessment(ctx *pulumi.Context, args *LookupAssessmentArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentResult, error) {
	var rv LookupAssessmentResult
	err := ctx.Invoke("azure-native:security/v20210601:getAssessment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentArgs struct {
	AssessmentName string  `pulumi:"assessmentName"`
	Expand         *string `pulumi:"expand"`
	ResourceId     string  `pulumi:"resourceId"`
}


type LookupAssessmentResult struct {
	AdditionalData  map[string]string                             `pulumi:"additionalData"`
	DisplayName     string                                        `pulumi:"displayName"`
	Id              string                                        `pulumi:"id"`
	Links           AssessmentLinksResponse                       `pulumi:"links"`
	Metadata        *SecurityAssessmentMetadataPropertiesResponse `pulumi:"metadata"`
	Name            string                                        `pulumi:"name"`
	PartnersData    *SecurityAssessmentPartnerDataResponse        `pulumi:"partnersData"`
	ResourceDetails interface{}                                   `pulumi:"resourceDetails"`
	Status          AssessmentStatusResponseResponse              `pulumi:"status"`
	Type            string                                        `pulumi:"type"`
}
