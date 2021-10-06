


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspaceConnection(ctx *pulumi.Context, args *LookupWorkspaceConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceConnectionResult, error) {
	var rv LookupWorkspaceConnectionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210101:getWorkspaceConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceConnectionResult struct {
	AuthType    *string `pulumi:"authType"`
	Category    *string `pulumi:"category"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	Target      *string `pulumi:"target"`
	Type        string  `pulumi:"type"`
	Value       *string `pulumi:"value"`
	ValueFormat *string `pulumi:"valueFormat"`
}
