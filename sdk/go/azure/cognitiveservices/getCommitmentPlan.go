


package cognitiveservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCommitmentPlan(ctx *pulumi.Context, args *LookupCommitmentPlanArgs, opts ...pulumi.InvokeOption) (*LookupCommitmentPlanResult, error) {
	var rv LookupCommitmentPlanResult
	err := ctx.Invoke("azure-native:cognitiveservices:getCommitmentPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCommitmentPlanArgs struct {
	AccountName        string `pulumi:"accountName"`
	CommitmentPlanName string `pulumi:"commitmentPlanName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupCommitmentPlanResult struct {
	Etag       string                           `pulumi:"etag"`
	Id         string                           `pulumi:"id"`
	Name       string                           `pulumi:"name"`
	Properties CommitmentPlanPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse               `pulumi:"systemData"`
	Type       string                           `pulumi:"type"`
}
