


package v20181101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleAssignmentArtifact(ctx *pulumi.Context, args *LookupRoleAssignmentArtifactArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentArtifactResult, error) {
	var rv LookupRoleAssignmentArtifactResult
	err := ctx.Invoke("azure-native:blueprint/v20181101preview:getRoleAssignmentArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAssignmentArtifactArgs struct {
	ArtifactName  string `pulumi:"artifactName"`
	BlueprintName string `pulumi:"blueprintName"`
	ResourceScope string `pulumi:"resourceScope"`
}


type LookupRoleAssignmentArtifactResult struct {
	DependsOn        []string    `pulumi:"dependsOn"`
	Description      *string     `pulumi:"description"`
	DisplayName      *string     `pulumi:"displayName"`
	Id               string      `pulumi:"id"`
	Kind             string      `pulumi:"kind"`
	Name             string      `pulumi:"name"`
	PrincipalIds     interface{} `pulumi:"principalIds"`
	ResourceGroup    *string     `pulumi:"resourceGroup"`
	RoleDefinitionId string      `pulumi:"roleDefinitionId"`
	Type             string      `pulumi:"type"`
}
