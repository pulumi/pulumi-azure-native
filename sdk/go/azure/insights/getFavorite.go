


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFavorite(ctx *pulumi.Context, args *LookupFavoriteArgs, opts ...pulumi.InvokeOption) (*LookupFavoriteResult, error) {
	var rv LookupFavoriteResult
	err := ctx.Invoke("azure-native:insights:getFavorite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFavoriteArgs struct {
	FavoriteId        string `pulumi:"favoriteId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupFavoriteResult struct {
	Category                *string  `pulumi:"category"`
	Config                  *string  `pulumi:"config"`
	FavoriteId              string   `pulumi:"favoriteId"`
	FavoriteType            *string  `pulumi:"favoriteType"`
	IsGeneratedFromTemplate *bool    `pulumi:"isGeneratedFromTemplate"`
	Name                    *string  `pulumi:"name"`
	SourceType              *string  `pulumi:"sourceType"`
	Tags                    []string `pulumi:"tags"`
	TimeModified            string   `pulumi:"timeModified"`
	UserId                  string   `pulumi:"userId"`
	Version                 *string  `pulumi:"version"`
}
