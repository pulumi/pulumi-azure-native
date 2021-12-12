


package v20200801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQueueServiceProperties(ctx *pulumi.Context, args *LookupQueueServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupQueueServicePropertiesResult, error) {
	var rv LookupQueueServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20200801preview:getQueueServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueServicePropertiesArgs struct {
	AccountName       string `pulumi:"accountName"`
	QueueServiceName  string `pulumi:"queueServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupQueueServicePropertiesResult struct {
	Cors *CorsRulesResponse `pulumi:"cors"`
	Id   string             `pulumi:"id"`
	Name string             `pulumi:"name"`
	Type string             `pulumi:"type"`
}
