


package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountAssembly(ctx *pulumi.Context, args *LookupIntegrationAccountAssemblyArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountAssemblyResult, error) {
	var rv LookupIntegrationAccountAssemblyResult
	err := ctx.Invoke("azure-native:logic:getIntegrationAccountAssembly", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountAssemblyArgs struct {
	AssemblyArtifactName   string `pulumi:"assemblyArtifactName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountAssemblyResult struct {
	Id         string                     `pulumi:"id"`
	Location   *string                    `pulumi:"location"`
	Name       string                     `pulumi:"name"`
	Properties AssemblyPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string          `pulumi:"tags"`
	Type       string                     `pulumi:"type"`
}
