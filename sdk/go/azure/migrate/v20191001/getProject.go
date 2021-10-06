


package v20191001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:migrate/v20191001:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupProjectResult struct {
	ETag       *string                   `pulumi:"eTag"`
	Id         string                    `pulumi:"id"`
	Location   *string                   `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties ProjectPropertiesResponse `pulumi:"properties"`
	Tags       interface{}               `pulumi:"tags"`
	Type       string                    `pulumi:"type"`
}
