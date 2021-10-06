


package v20201020preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Getgetbot(ctx *pulumi.Context, args *GetgetbotArgs, opts ...pulumi.InvokeOption) (*GetgetbotResult, error) {
	var rv GetgetbotResult
	err := ctx.Invoke("azure-native:healthbot/v20201020preview:getgetbot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetgetbotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetgetbotResult struct {
	Id         string                      `pulumi:"id"`
	Location   string                      `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties HealthBotPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                `pulumi:"sku"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}
