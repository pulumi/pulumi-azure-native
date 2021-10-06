


package v20180915preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	var rv LookupTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20180915preview:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TopicName         string `pulumi:"topicName"`
}


type LookupTopicResult struct {
	Endpoint           string                          `pulumi:"endpoint"`
	Id                 string                          `pulumi:"id"`
	InputSchema        *string                         `pulumi:"inputSchema"`
	InputSchemaMapping *JsonInputSchemaMappingResponse `pulumi:"inputSchemaMapping"`
	Location           string                          `pulumi:"location"`
	Name               string                          `pulumi:"name"`
	ProvisioningState  string                          `pulumi:"provisioningState"`
	Tags               map[string]string               `pulumi:"tags"`
	Type               string                          `pulumi:"type"`
}
