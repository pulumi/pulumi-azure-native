


package v20151101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSolution(ctx *pulumi.Context, args *LookupSolutionArgs, opts ...pulumi.InvokeOption) (*LookupSolutionResult, error) {
	var rv LookupSolutionResult
	err := ctx.Invoke("azure-native:operationsmanagement/v20151101preview:getSolution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSolutionArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SolutionName      string `pulumi:"solutionName"`
}


type LookupSolutionResult struct {
	Id         string                     `pulumi:"id"`
	Location   *string                    `pulumi:"location"`
	Name       string                     `pulumi:"name"`
	Plan       *SolutionPlanResponse      `pulumi:"plan"`
	Properties SolutionPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string          `pulumi:"tags"`
	Type       string                     `pulumi:"type"`
}
