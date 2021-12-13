


package blueprint

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyAssignmentArtifact(ctx *pulumi.Context, args *LookupPolicyAssignmentArtifactArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentArtifactResult, error) {
	var rv LookupPolicyAssignmentArtifactResult
	err := ctx.Invoke("azure-native:blueprint:getPolicyAssignmentArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyAssignmentArtifactArgs struct {
	ArtifactName  string `pulumi:"artifactName"`
	BlueprintName string `pulumi:"blueprintName"`
	ResourceScope string `pulumi:"resourceScope"`
}


type LookupPolicyAssignmentArtifactResult struct {
	DependsOn          []string                          `pulumi:"dependsOn"`
	Description        *string                           `pulumi:"description"`
	DisplayName        *string                           `pulumi:"displayName"`
	Id                 string                            `pulumi:"id"`
	Kind               string                            `pulumi:"kind"`
	Name               string                            `pulumi:"name"`
	Parameters         map[string]ParameterValueResponse `pulumi:"parameters"`
	PolicyDefinitionId string                            `pulumi:"policyDefinitionId"`
	ResourceGroup      *string                           `pulumi:"resourceGroup"`
	Type               string                            `pulumi:"type"`
}
