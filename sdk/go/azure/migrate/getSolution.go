


package migrate

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSolution(ctx *pulumi.Context, args *LookupSolutionArgs, opts ...pulumi.InvokeOption) (*LookupSolutionResult, error) {
	var rv LookupSolutionResult
	err := ctx.Invoke("azure-native:migrate:getSolution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSolutionArgs struct {
	MigrateProjectName string `pulumi:"migrateProjectName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	SolutionName       string `pulumi:"solutionName"`
}


type LookupSolutionResult struct {
	Etag       *string                    `pulumi:"etag"`
	Id         string                     `pulumi:"id"`
	Name       string                     `pulumi:"name"`
	Properties SolutionPropertiesResponse `pulumi:"properties"`
	Type       string                     `pulumi:"type"`
}
