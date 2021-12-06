


package migrate

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("azure-native:migrate:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	GroupName         string `pulumi:"groupName"`
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupGroupResult struct {
	ETag       *string                 `pulumi:"eTag"`
	Id         string                  `pulumi:"id"`
	Name       string                  `pulumi:"name"`
	Properties GroupPropertiesResponse `pulumi:"properties"`
	Type       string                  `pulumi:"type"`
}
