


package v20201216preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFavoriteProcess(ctx *pulumi.Context, args *LookupFavoriteProcessArgs, opts ...pulumi.InvokeOption) (*LookupFavoriteProcessResult, error) {
	var rv LookupFavoriteProcessResult
	err := ctx.Invoke("azure-native:testbase/v20201216preview:getFavoriteProcess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFavoriteProcessArgs struct {
	FavoriteProcessResourceName string `pulumi:"favoriteProcessResourceName"`
	PackageName                 string `pulumi:"packageName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	TestBaseAccountName         string `pulumi:"testBaseAccountName"`
}


type LookupFavoriteProcessResult struct {
	ActualProcessName string             `pulumi:"actualProcessName"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}
