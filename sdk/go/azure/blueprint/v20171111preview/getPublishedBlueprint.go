


package v20171111preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPublishedBlueprint(ctx *pulumi.Context, args *LookupPublishedBlueprintArgs, opts ...pulumi.InvokeOption) (*LookupPublishedBlueprintResult, error) {
	var rv LookupPublishedBlueprintResult
	err := ctx.Invoke("azure-native:blueprint/v20171111preview:getPublishedBlueprint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublishedBlueprintArgs struct {
	BlueprintName       string `pulumi:"blueprintName"`
	ManagementGroupName string `pulumi:"managementGroupName"`
	VersionId           string `pulumi:"versionId"`
}


type LookupPublishedBlueprintResult struct {
	BlueprintName  *string                                    `pulumi:"blueprintName"`
	ChangeNotes    *string                                    `pulumi:"changeNotes"`
	Description    *string                                    `pulumi:"description"`
	DisplayName    *string                                    `pulumi:"displayName"`
	Id             string                                     `pulumi:"id"`
	Name           string                                     `pulumi:"name"`
	Parameters     map[string]ParameterDefinitionResponse     `pulumi:"parameters"`
	ResourceGroups map[string]ResourceGroupDefinitionResponse `pulumi:"resourceGroups"`
	Status         BlueprintStatusResponse                    `pulumi:"status"`
	TargetScope    *string                                    `pulumi:"targetScope"`
	Type           string                                     `pulumi:"type"`
}
