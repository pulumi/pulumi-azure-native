


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeploymentAtScope(ctx *pulumi.Context, args *LookupDeploymentAtScopeArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentAtScopeResult, error) {
	var rv LookupDeploymentAtScopeResult
	err := ctx.Invoke("azure-native:resources/v20190801:getDeploymentAtScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentAtScopeArgs struct {
	DeploymentName string `pulumi:"deploymentName"`
	Scope          string `pulumi:"scope"`
}


type LookupDeploymentAtScopeResult struct {
	Id         string                               `pulumi:"id"`
	Location   *string                              `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponse `pulumi:"properties"`
	Type       string                               `pulumi:"type"`
}
