


package v20180901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMigrateProject(ctx *pulumi.Context, args *LookupMigrateProjectArgs, opts ...pulumi.InvokeOption) (*LookupMigrateProjectResult, error) {
	var rv LookupMigrateProjectResult
	err := ctx.Invoke("azure-native:migrate/v20180901preview:getMigrateProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrateProjectArgs struct {
	MigrateProjectName string `pulumi:"migrateProjectName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupMigrateProjectResult struct {
	ETag       *string                          `pulumi:"eTag"`
	Id         string                           `pulumi:"id"`
	Location   *string                          `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties MigrateProjectPropertiesResponse `pulumi:"properties"`
	Tags       *MigrateProjectResponseTags      `pulumi:"tags"`
	Type       string                           `pulumi:"type"`
}
