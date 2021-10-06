


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeploymentAtManagementGroupScope(ctx *pulumi.Context, args *LookupDeploymentAtManagementGroupScopeArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentAtManagementGroupScopeResult, error) {
	var rv LookupDeploymentAtManagementGroupScopeResult
	err := ctx.Invoke("azure-native:resources/v20200601:getDeploymentAtManagementGroupScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentAtManagementGroupScopeArgs struct {
	DeploymentName string `pulumi:"deploymentName"`
	GroupId        string `pulumi:"groupId"`
}


type LookupDeploymentAtManagementGroupScopeResult struct {
	Id         string                               `pulumi:"id"`
	Location   *string                              `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponse `pulumi:"properties"`
	Tags       map[string]string                    `pulumi:"tags"`
	Type       string                               `pulumi:"type"`
}
