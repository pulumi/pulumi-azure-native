


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebPubSubSharedPrivateLinkResource(ctx *pulumi.Context, args *LookupWebPubSubSharedPrivateLinkResourceArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubSharedPrivateLinkResourceResult, error) {
	var rv LookupWebPubSubSharedPrivateLinkResourceResult
	err := ctx.Invoke("azure-native:webpubsub/v20210901preview:getWebPubSubSharedPrivateLinkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubSharedPrivateLinkResourceArgs struct {
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
	SharedPrivateLinkResourceName string `pulumi:"sharedPrivateLinkResourceName"`
}


type LookupWebPubSubSharedPrivateLinkResourceResult struct {
	GroupId               string             `pulumi:"groupId"`
	Id                    string             `pulumi:"id"`
	Name                  string             `pulumi:"name"`
	PrivateLinkResourceId string             `pulumi:"privateLinkResourceId"`
	ProvisioningState     string             `pulumi:"provisioningState"`
	RequestMessage        *string            `pulumi:"requestMessage"`
	Status                string             `pulumi:"status"`
	SystemData            SystemDataResponse `pulumi:"systemData"`
	Type                  string             `pulumi:"type"`
}
