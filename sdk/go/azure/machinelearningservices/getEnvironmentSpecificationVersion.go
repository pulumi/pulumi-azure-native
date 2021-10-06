


package machinelearningservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironmentSpecificationVersion(ctx *pulumi.Context, args *LookupEnvironmentSpecificationVersionArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentSpecificationVersionResult, error) {
	var rv LookupEnvironmentSpecificationVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices:getEnvironmentSpecificationVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentSpecificationVersionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupEnvironmentSpecificationVersionResult struct {
	Id         string                                  `pulumi:"id"`
	Name       string                                  `pulumi:"name"`
	Properties EnvironmentSpecificationVersionResponse `pulumi:"properties"`
	SystemData SystemDataResponse                      `pulumi:"systemData"`
	Type       string                                  `pulumi:"type"`
}
