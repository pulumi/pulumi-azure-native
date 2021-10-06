


package migrate

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSolutionConfig(ctx *pulumi.Context, args *GetSolutionConfigArgs, opts ...pulumi.InvokeOption) (*GetSolutionConfigResult, error) {
	var rv GetSolutionConfigResult
	err := ctx.Invoke("azure-native:migrate:getSolutionConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSolutionConfigArgs struct {
	MigrateProjectName string `pulumi:"migrateProjectName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	SolutionName       string `pulumi:"solutionName"`
}


type GetSolutionConfigResult struct {
	PublisherSasUri *string `pulumi:"publisherSasUri"`
}
