


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEntityQuery(ctx *pulumi.Context, args *LookupEntityQueryArgs, opts ...pulumi.InvokeOption) (*LookupEntityQueryResult, error) {
	var rv LookupEntityQueryResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getEntityQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEntityQueryArgs struct {
	EntityQueryId     string `pulumi:"entityQueryId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupEntityQueryResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}
