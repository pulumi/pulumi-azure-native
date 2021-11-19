


package v20171111preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupArtifact(ctx *pulumi.Context, args *LookupArtifactArgs, opts ...pulumi.InvokeOption) (*LookupArtifactResult, error) {
	var rv LookupArtifactResult
	err := ctx.Invoke("azure-native:blueprint/v20171111preview:getArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactArgs struct {
	ArtifactName        string `pulumi:"artifactName"`
	BlueprintName       string `pulumi:"blueprintName"`
	ManagementGroupName string `pulumi:"managementGroupName"`
}


type LookupArtifactResult struct {
	Id   string `pulumi:"id"`
	Kind string `pulumi:"kind"`
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}
