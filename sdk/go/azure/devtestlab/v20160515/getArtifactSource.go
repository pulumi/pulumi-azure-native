


package v20160515

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupArtifactSource(ctx *pulumi.Context, args *LookupArtifactSourceArgs, opts ...pulumi.InvokeOption) (*LookupArtifactSourceResult, error) {
	var rv LookupArtifactSourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:getArtifactSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactSourceArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupArtifactSourceResult struct {
	ArmTemplateFolderPath *string           `pulumi:"armTemplateFolderPath"`
	BranchRef             *string           `pulumi:"branchRef"`
	CreatedDate           string            `pulumi:"createdDate"`
	DisplayName           *string           `pulumi:"displayName"`
	FolderPath            *string           `pulumi:"folderPath"`
	Id                    string            `pulumi:"id"`
	Location              *string           `pulumi:"location"`
	Name                  string            `pulumi:"name"`
	ProvisioningState     *string           `pulumi:"provisioningState"`
	SecurityToken         *string           `pulumi:"securityToken"`
	SourceType            *string           `pulumi:"sourceType"`
	Status                *string           `pulumi:"status"`
	Tags                  map[string]string `pulumi:"tags"`
	Type                  string            `pulumi:"type"`
	UniqueIdentifier      *string           `pulumi:"uniqueIdentifier"`
	Uri                   *string           `pulumi:"uri"`
}
