


package v20140901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQueue(ctx *pulumi.Context, args *LookupQueueArgs, opts ...pulumi.InvokeOption) (*LookupQueueResult, error) {
	var rv LookupQueueResult
	err := ctx.Invoke("azure-native:servicebus/v20140901:getQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	QueueName         string `pulumi:"queueName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupQueueResult struct {
	AccessedAt                          string                      `pulumi:"accessedAt"`
	AutoDeleteOnIdle                    *string                     `pulumi:"autoDeleteOnIdle"`
	CountDetails                        MessageCountDetailsResponse `pulumi:"countDetails"`
	CreatedAt                           string                      `pulumi:"createdAt"`
	DeadLetteringOnMessageExpiration    *bool                       `pulumi:"deadLetteringOnMessageExpiration"`
	DefaultMessageTimeToLive            *string                     `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow *string                     `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations             *bool                       `pulumi:"enableBatchedOperations"`
	EnableExpress                       *bool                       `pulumi:"enableExpress"`
	EnablePartitioning                  *bool                       `pulumi:"enablePartitioning"`
	EntityAvailabilityStatus            *string                     `pulumi:"entityAvailabilityStatus"`
	Id                                  string                      `pulumi:"id"`
	IsAnonymousAccessible               *bool                       `pulumi:"isAnonymousAccessible"`
	Location                            *string                     `pulumi:"location"`
	LockDuration                        *string                     `pulumi:"lockDuration"`
	MaxDeliveryCount                    *int                        `pulumi:"maxDeliveryCount"`
	MaxSizeInMegabytes                  *float64                    `pulumi:"maxSizeInMegabytes"`
	MessageCount                        float64                     `pulumi:"messageCount"`
	Name                                string                      `pulumi:"name"`
	RequiresDuplicateDetection          *bool                       `pulumi:"requiresDuplicateDetection"`
	RequiresSession                     *bool                       `pulumi:"requiresSession"`
	SizeInBytes                         float64                     `pulumi:"sizeInBytes"`
	Status                              *string                     `pulumi:"status"`
	SupportOrdering                     *bool                       `pulumi:"supportOrdering"`
	Type                                string                      `pulumi:"type"`
	UpdatedAt                           string                      `pulumi:"updatedAt"`
}
