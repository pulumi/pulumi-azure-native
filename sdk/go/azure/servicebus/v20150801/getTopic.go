


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	var rv LookupTopicResult
	err := ctx.Invoke("azure-native:servicebus/v20150801:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TopicName         string `pulumi:"topicName"`
}


type LookupTopicResult struct {
	AccessedAt                          string                      `pulumi:"accessedAt"`
	AutoDeleteOnIdle                    *string                     `pulumi:"autoDeleteOnIdle"`
	CountDetails                        MessageCountDetailsResponse `pulumi:"countDetails"`
	CreatedAt                           string                      `pulumi:"createdAt"`
	DefaultMessageTimeToLive            *string                     `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow *string                     `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations             *bool                       `pulumi:"enableBatchedOperations"`
	EnableExpress                       *bool                       `pulumi:"enableExpress"`
	EnablePartitioning                  *bool                       `pulumi:"enablePartitioning"`
	EntityAvailabilityStatus            *string                     `pulumi:"entityAvailabilityStatus"`
	FilteringMessagesBeforePublishing   *bool                       `pulumi:"filteringMessagesBeforePublishing"`
	Id                                  string                      `pulumi:"id"`
	IsAnonymousAccessible               *bool                       `pulumi:"isAnonymousAccessible"`
	IsExpress                           *bool                       `pulumi:"isExpress"`
	Location                            *string                     `pulumi:"location"`
	MaxSizeInMegabytes                  *float64                    `pulumi:"maxSizeInMegabytes"`
	Name                                string                      `pulumi:"name"`
	RequiresDuplicateDetection          *bool                       `pulumi:"requiresDuplicateDetection"`
	SizeInBytes                         float64                     `pulumi:"sizeInBytes"`
	Status                              *string                     `pulumi:"status"`
	SubscriptionCount                   int                         `pulumi:"subscriptionCount"`
	SupportOrdering                     *bool                       `pulumi:"supportOrdering"`
	Type                                string                      `pulumi:"type"`
	UpdatedAt                           string                      `pulumi:"updatedAt"`
}
