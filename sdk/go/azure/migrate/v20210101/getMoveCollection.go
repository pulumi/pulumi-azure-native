


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMoveCollection(ctx *pulumi.Context, args *LookupMoveCollectionArgs, opts ...pulumi.InvokeOption) (*LookupMoveCollectionResult, error) {
	var rv LookupMoveCollectionResult
	err := ctx.Invoke("azure-native:migrate/v20210101:getMoveCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMoveCollectionArgs struct {
	MoveCollectionName string `pulumi:"moveCollectionName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupMoveCollectionResult struct {
	Etag       string                           `pulumi:"etag"`
	Id         string                           `pulumi:"id"`
	Identity   *IdentityResponse                `pulumi:"identity"`
	Location   *string                          `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties MoveCollectionPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                `pulumi:"tags"`
	Type       string                           `pulumi:"type"`
}
