


package v20191001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDeploymentScript(ctx *pulumi.Context, args *LookupDeploymentScriptArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentScriptResult, error) {
	var rv LookupDeploymentScriptResult
	err := ctx.Invoke("azure-native:resources/v20191001preview:getDeploymentScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentScriptArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScriptName        string `pulumi:"scriptName"`
}


type LookupDeploymentScriptResult struct {
	Id         string                         `pulumi:"id"`
	Identity   ManagedServiceIdentityResponse `pulumi:"identity"`
	Kind       string                         `pulumi:"kind"`
	Location   string                         `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	SystemData SystemDataResponse             `pulumi:"systemData"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}
