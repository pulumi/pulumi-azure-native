


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationRuntime(ctx *pulumi.Context, args *LookupIntegrationRuntimeArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationRuntimeResult, error) {
	var rv LookupIntegrationRuntimeResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getIntegrationRuntime", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationRuntimeArgs struct {
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type LookupIntegrationRuntimeResult struct {
	Etag       string      `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}
