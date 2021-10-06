


package azurestack

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedSubscription(ctx *pulumi.Context, args *LookupLinkedSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupLinkedSubscriptionResult, error) {
	var rv LookupLinkedSubscriptionResult
	err := ctx.Invoke("azure-native:azurestack:getLinkedSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedSubscriptionArgs struct {
	LinkedSubscriptionName string `pulumi:"linkedSubscriptionName"`
	ResourceGroup          string `pulumi:"resourceGroup"`
}


type LookupLinkedSubscriptionResult struct {
	DeviceConnectionStatus string             `pulumi:"deviceConnectionStatus"`
	DeviceId               string             `pulumi:"deviceId"`
	DeviceLinkState        string             `pulumi:"deviceLinkState"`
	DeviceObjectId         string             `pulumi:"deviceObjectId"`
	Etag                   *string            `pulumi:"etag"`
	Id                     string             `pulumi:"id"`
	Kind                   string             `pulumi:"kind"`
	LastConnectedTime      string             `pulumi:"lastConnectedTime"`
	LinkedSubscriptionId   *string            `pulumi:"linkedSubscriptionId"`
	Location               string             `pulumi:"location"`
	Name                   string             `pulumi:"name"`
	RegistrationResourceId *string            `pulumi:"registrationResourceId"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
	Tags                   map[string]string  `pulumi:"tags"`
	Type                   string             `pulumi:"type"`
}
