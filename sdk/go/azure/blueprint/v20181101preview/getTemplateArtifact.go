


package v20181101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTemplateArtifact(ctx *pulumi.Context, args *LookupTemplateArtifactArgs, opts ...pulumi.InvokeOption) (*LookupTemplateArtifactResult, error) {
	var rv LookupTemplateArtifactResult
	err := ctx.Invoke("azure-native:blueprint/v20181101preview:getTemplateArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateArtifactArgs struct {
	ArtifactName  string `pulumi:"artifactName"`
	BlueprintName string `pulumi:"blueprintName"`
	ResourceScope string `pulumi:"resourceScope"`
}


type LookupTemplateArtifactResult struct {
	DependsOn     []string                          `pulumi:"dependsOn"`
	Description   *string                           `pulumi:"description"`
	DisplayName   *string                           `pulumi:"displayName"`
	Id            string                            `pulumi:"id"`
	Kind          string                            `pulumi:"kind"`
	Name          string                            `pulumi:"name"`
	Parameters    map[string]ParameterValueResponse `pulumi:"parameters"`
	ResourceGroup *string                           `pulumi:"resourceGroup"`
	Template      interface{}                       `pulumi:"template"`
	Type          string                            `pulumi:"type"`
}
