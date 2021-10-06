


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLiveToken(ctx *pulumi.Context, args *GetLiveTokenArgs, opts ...pulumi.InvokeOption) (*GetLiveTokenResult, error) {
	var rv GetLiveTokenResult
	err := ctx.Invoke("azure-native:insights:getLiveToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLiveTokenArgs struct {
	ResourceUri string `pulumi:"resourceUri"`
}


type GetLiveTokenResult struct {
	LiveToken string `pulumi:"liveToken"`
}
