


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOnlineDeployment(ctx *pulumi.Context, args *LookupOnlineDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupOnlineDeploymentResult, error) {
	var rv LookupOnlineDeploymentResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getOnlineDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOnlineDeploymentArgs struct {
	DeploymentName    string `pulumi:"deploymentName"`
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type LookupOnlineDeploymentResult struct {
	Id         string                    `pulumi:"id"`
	Identity   *ResourceIdentityResponse `pulumi:"identity"`
	Kind       *string                   `pulumi:"kind"`
	Location   string                    `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties interface{}               `pulumi:"properties"`
	SystemData SystemDataResponse        `pulumi:"systemData"`
	Tags       map[string]string         `pulumi:"tags"`
	Type       string                    `pulumi:"type"`
}
