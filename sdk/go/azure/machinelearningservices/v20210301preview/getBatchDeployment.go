


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBatchDeployment(ctx *pulumi.Context, args *LookupBatchDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupBatchDeploymentResult, error) {
	var rv LookupBatchDeploymentResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getBatchDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBatchDeploymentArgs struct {
	DeploymentName    string `pulumi:"deploymentName"`
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type LookupBatchDeploymentResult struct {
	Id         string                    `pulumi:"id"`
	Identity   *ResourceIdentityResponse `pulumi:"identity"`
	Kind       *string                   `pulumi:"kind"`
	Location   string                    `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties BatchDeploymentResponse   `pulumi:"properties"`
	SystemData SystemDataResponse        `pulumi:"systemData"`
	Tags       map[string]string         `pulumi:"tags"`
	Type       string                    `pulumi:"type"`
}
