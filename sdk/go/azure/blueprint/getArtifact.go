


package blueprint

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupArtifact(ctx *pulumi.Context, args *LookupArtifactArgs, opts ...pulumi.InvokeOption) (*LookupArtifactResult, error) {
	var rv LookupArtifactResult
	err := ctx.Invoke("azure-native:blueprint:getArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactArgs struct {
	ArtifactName  string `pulumi:"artifactName"`
	BlueprintName string `pulumi:"blueprintName"`
	ResourceScope string `pulumi:"resourceScope"`
}


type LookupArtifactResult struct {
	Id   string `pulumi:"id"`
	Kind string `pulumi:"kind"`
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}
