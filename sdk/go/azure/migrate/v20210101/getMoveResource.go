


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMoveResource(ctx *pulumi.Context, args *LookupMoveResourceArgs, opts ...pulumi.InvokeOption) (*LookupMoveResourceResult, error) {
	var rv LookupMoveResourceResult
	err := ctx.Invoke("azure-native:migrate/v20210101:getMoveResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMoveResourceArgs struct {
	MoveCollectionName string `pulumi:"moveCollectionName"`
	MoveResourceName   string `pulumi:"moveResourceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupMoveResourceResult struct {
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties MoveResourcePropertiesResponse `pulumi:"properties"`
	Type       string                         `pulumi:"type"`
}
