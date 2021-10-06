


package v20210115preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIngestionSettingTokens(ctx *pulumi.Context, args *ListIngestionSettingTokensArgs, opts ...pulumi.InvokeOption) (*ListIngestionSettingTokensResult, error) {
	var rv ListIngestionSettingTokensResult
	err := ctx.Invoke("azure-native:security/v20210115preview:listIngestionSettingTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIngestionSettingTokensArgs struct {
	IngestionSettingName string `pulumi:"ingestionSettingName"`
}


type ListIngestionSettingTokensResult struct {
	Token string `pulumi:"token"`
}
