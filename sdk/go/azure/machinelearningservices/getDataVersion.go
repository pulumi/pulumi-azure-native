


package machinelearningservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataVersion(ctx *pulumi.Context, args *LookupDataVersionArgs, opts ...pulumi.InvokeOption) (*LookupDataVersionResult, error) {
	var rv LookupDataVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices:getDataVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataVersionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDataVersionResult struct {
	Id         string              `pulumi:"id"`
	Name       string              `pulumi:"name"`
	Properties DataVersionResponse `pulumi:"properties"`
	SystemData SystemDataResponse  `pulumi:"systemData"`
	Type       string              `pulumi:"type"`
}
