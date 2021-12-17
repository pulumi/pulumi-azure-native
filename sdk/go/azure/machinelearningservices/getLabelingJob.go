


package machinelearningservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLabelingJob(ctx *pulumi.Context, args *LookupLabelingJobArgs, opts ...pulumi.InvokeOption) (*LookupLabelingJobResult, error) {
	var rv LookupLabelingJobResult
	err := ctx.Invoke("azure-native:machinelearningservices:getLabelingJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabelingJobArgs struct {
	IncludeJobInstructions *bool  `pulumi:"includeJobInstructions"`
	IncludeLabelCategories *bool  `pulumi:"includeLabelCategories"`
	LabelingJobId          string `pulumi:"labelingJobId"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type LookupLabelingJobResult struct {
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties LabelingJobPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Type       string                        `pulumi:"type"`
}
