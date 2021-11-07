


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemTopic(ctx *pulumi.Context, args *LookupSystemTopicArgs, opts ...pulumi.InvokeOption) (*LookupSystemTopicResult, error) {
	var rv LookupSystemTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20210601preview:getSystemTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSystemTopicArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SystemTopicName   string `pulumi:"systemTopicName"`
}


type LookupSystemTopicResult struct {
	Id                string                `pulumi:"id"`
	Identity          *IdentityInfoResponse `pulumi:"identity"`
	Location          string                `pulumi:"location"`
	MetricResourceId  string                `pulumi:"metricResourceId"`
	Name              string                `pulumi:"name"`
	ProvisioningState string                `pulumi:"provisioningState"`
	Source            *string               `pulumi:"source"`
	SystemData        SystemDataResponse    `pulumi:"systemData"`
	Tags              map[string]string     `pulumi:"tags"`
	TopicType         *string               `pulumi:"topicType"`
	Type              string                `pulumi:"type"`
}
