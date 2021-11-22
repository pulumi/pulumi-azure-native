


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIotHubResourceKeysForKeyName(ctx *pulumi.Context, args *ListIotHubResourceKeysForKeyNameArgs, opts ...pulumi.InvokeOption) (*ListIotHubResourceKeysForKeyNameResult, error) {
	var rv ListIotHubResourceKeysForKeyNameResult
	err := ctx.Invoke("azure-native:devices/v20210201preview:listIotHubResourceKeysForKeyName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotHubResourceKeysForKeyNameArgs struct {
	KeyName           string `pulumi:"keyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListIotHubResourceKeysForKeyNameResult struct {
	KeyName      string  `pulumi:"keyName"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	Rights       string  `pulumi:"rights"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}
