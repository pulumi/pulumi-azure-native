


package v20160601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomApi(ctx *pulumi.Context, args *LookupCustomApiArgs, opts ...pulumi.InvokeOption) (*LookupCustomApiResult, error) {
	var rv LookupCustomApiResult
	err := ctx.Invoke("azure-native:web/v20160601:getCustomApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomApiArgs struct {
	ApiName           string  `pulumi:"apiName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SubscriptionId    *string `pulumi:"subscriptionId"`
}


type LookupCustomApiResult struct {
	Etag       *string                               `pulumi:"etag"`
	Id         string                                `pulumi:"id"`
	Location   *string                               `pulumi:"location"`
	Name       string                                `pulumi:"name"`
	Properties CustomApiPropertiesDefinitionResponse `pulumi:"properties"`
	Tags       map[string]string                     `pulumi:"tags"`
	Type       string                                `pulumi:"type"`
}
