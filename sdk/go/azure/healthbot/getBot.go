


package healthbot

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBot(ctx *pulumi.Context, args *LookupBotArgs, opts ...pulumi.InvokeOption) (*LookupBotResult, error) {
	var rv LookupBotResult
	err := ctx.Invoke("azure-native:healthbot:getBot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBotArgs struct {
	BotName           string `pulumi:"botName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBotResult struct {
	Id         string                      `pulumi:"id"`
	Identity   *IdentityResponse           `pulumi:"identity"`
	Location   string                      `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties HealthBotPropertiesResponse `pulumi:"properties"`
	Sku        SkuResponse                 `pulumi:"sku"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}
