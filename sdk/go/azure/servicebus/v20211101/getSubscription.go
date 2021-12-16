


package v20211101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubscription(ctx *pulumi.Context, args *LookupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionResult, error) {
	var rv LookupSubscriptionResult
	err := ctx.Invoke("azure-native:servicebus/v20211101:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SubscriptionName  string `pulumi:"subscriptionName"`
	TopicName         string `pulumi:"topicName"`
}


type LookupSubscriptionResult struct {
	AccessedAt                                string                            `pulumi:"accessedAt"`
	AutoDeleteOnIdle                          *string                           `pulumi:"autoDeleteOnIdle"`
	ClientAffineProperties                    *SBClientAffinePropertiesResponse `pulumi:"clientAffineProperties"`
	CountDetails                              MessageCountDetailsResponse       `pulumi:"countDetails"`
	CreatedAt                                 string                            `pulumi:"createdAt"`
	DeadLetteringOnFilterEvaluationExceptions *bool                             `pulumi:"deadLetteringOnFilterEvaluationExceptions"`
	DeadLetteringOnMessageExpiration          *bool                             `pulumi:"deadLetteringOnMessageExpiration"`
	DefaultMessageTimeToLive                  *string                           `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow       *string                           `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations                   *bool                             `pulumi:"enableBatchedOperations"`
	ForwardDeadLetteredMessagesTo             *string                           `pulumi:"forwardDeadLetteredMessagesTo"`
	ForwardTo                                 *string                           `pulumi:"forwardTo"`
	Id                                        string                            `pulumi:"id"`
	IsClientAffine                            *bool                             `pulumi:"isClientAffine"`
	Location                                  string                            `pulumi:"location"`
	LockDuration                              *string                           `pulumi:"lockDuration"`
	MaxDeliveryCount                          *int                              `pulumi:"maxDeliveryCount"`
	MessageCount                              float64                           `pulumi:"messageCount"`
	Name                                      string                            `pulumi:"name"`
	RequiresSession                           *bool                             `pulumi:"requiresSession"`
	Status                                    *string                           `pulumi:"status"`
	SystemData                                SystemDataResponse                `pulumi:"systemData"`
	Type                                      string                            `pulumi:"type"`
	UpdatedAt                                 string                            `pulumi:"updatedAt"`
}
