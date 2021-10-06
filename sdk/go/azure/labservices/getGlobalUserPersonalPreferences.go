


package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGlobalUserPersonalPreferences(ctx *pulumi.Context, args *GetGlobalUserPersonalPreferencesArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserPersonalPreferencesResult, error) {
	var rv GetGlobalUserPersonalPreferencesResult
	err := ctx.Invoke("azure-native:labservices:getGlobalUserPersonalPreferences", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserPersonalPreferencesArgs struct {
	AddRemove            *string `pulumi:"addRemove"`
	LabAccountResourceId *string `pulumi:"labAccountResourceId"`
	LabResourceId        *string `pulumi:"labResourceId"`
	UserName             string  `pulumi:"userName"`
}


type GetGlobalUserPersonalPreferencesResult struct {
	FavoriteLabResourceIds []string `pulumi:"favoriteLabResourceIds"`
	Id                     *string  `pulumi:"id"`
}
