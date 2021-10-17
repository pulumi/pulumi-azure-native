


package webpubsub

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebPubSubKeys(ctx *pulumi.Context, args *ListWebPubSubKeysArgs, opts ...pulumi.InvokeOption) (*ListWebPubSubKeysResult, error) {
	var rv ListWebPubSubKeysResult
	err := ctx.Invoke("azure-native:webpubsub:listWebPubSubKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebPubSubKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListWebPubSubKeysResult struct {
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}
