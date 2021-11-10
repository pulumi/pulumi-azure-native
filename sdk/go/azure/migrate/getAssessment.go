


package migrate

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssessment(ctx *pulumi.Context, args *LookupAssessmentArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentResult, error) {
	var rv LookupAssessmentResult
	err := ctx.Invoke("azure-native:migrate:getAssessment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentArgs struct {
	AssessmentName    string `pulumi:"assessmentName"`
	GroupName         string `pulumi:"groupName"`
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAssessmentResult struct {
	ETag       *string                      `pulumi:"eTag"`
	Id         string                       `pulumi:"id"`
	Name       string                       `pulumi:"name"`
	Properties AssessmentPropertiesResponse `pulumi:"properties"`
	Type       string                       `pulumi:"type"`
}
