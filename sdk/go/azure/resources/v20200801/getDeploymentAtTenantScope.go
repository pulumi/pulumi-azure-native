


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeploymentAtTenantScope(ctx *pulumi.Context, args *LookupDeploymentAtTenantScopeArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentAtTenantScopeResult, error) {
	var rv LookupDeploymentAtTenantScopeResult
	err := ctx.Invoke("azure-native:resources/v20200801:getDeploymentAtTenantScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentAtTenantScopeArgs struct {
	DeploymentName string `pulumi:"deploymentName"`
}


type LookupDeploymentAtTenantScopeResult struct {
	Id         string                               `pulumi:"id"`
	Location   *string                              `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponse `pulumi:"properties"`
	Tags       map[string]string                    `pulumi:"tags"`
	Type       string                               `pulumi:"type"`
}
