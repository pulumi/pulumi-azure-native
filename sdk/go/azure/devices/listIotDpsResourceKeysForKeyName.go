


package devices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIotDpsResourceKeysForKeyName(ctx *pulumi.Context, args *ListIotDpsResourceKeysForKeyNameArgs, opts ...pulumi.InvokeOption) (*ListIotDpsResourceKeysForKeyNameResult, error) {
	var rv ListIotDpsResourceKeysForKeyNameResult
	err := ctx.Invoke("azure-native:devices:listIotDpsResourceKeysForKeyName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotDpsResourceKeysForKeyNameArgs struct {
	KeyName                 string `pulumi:"keyName"`
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type ListIotDpsResourceKeysForKeyNameResult struct {
	KeyName      string  `pulumi:"keyName"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	Rights       string  `pulumi:"rights"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}
