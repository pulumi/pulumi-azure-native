


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteSecrets(ctx *pulumi.Context, args *ListStaticSiteSecretsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteSecretsResult, error) {
	var rv ListStaticSiteSecretsResult
	err := ctx.Invoke("azure-native:web/v20200601:listStaticSiteSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteSecretsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteSecretsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
