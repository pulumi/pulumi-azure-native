


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBotConnectionWithSecrets(ctx *pulumi.Context, args *ListBotConnectionWithSecretsArgs, opts ...pulumi.InvokeOption) (*ListBotConnectionWithSecretsResult, error) {
	var rv ListBotConnectionWithSecretsResult
	err := ctx.Invoke("azure-native:botservice/v20210501preview:listBotConnectionWithSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBotConnectionWithSecretsArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListBotConnectionWithSecretsResult struct {
	Etag       *string                             `pulumi:"etag"`
	Id         string                              `pulumi:"id"`
	Kind       *string                             `pulumi:"kind"`
	Location   *string                             `pulumi:"location"`
	Name       string                              `pulumi:"name"`
	Properties ConnectionSettingPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                        `pulumi:"sku"`
	Tags       map[string]string                   `pulumi:"tags"`
	Type       string                              `pulumi:"type"`
}
