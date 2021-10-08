


package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountAssemblyContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountAssemblyContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountAssemblyContentCallbackUrlResult, error) {
	var rv ListIntegrationAccountAssemblyContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20190501:listIntegrationAccountAssemblyContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountAssemblyContentCallbackUrlArgs struct {
	AssemblyArtifactName   string `pulumi:"assemblyArtifactName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type ListIntegrationAccountAssemblyContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}
