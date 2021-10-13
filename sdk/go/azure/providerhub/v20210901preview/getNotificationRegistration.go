


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotificationRegistration(ctx *pulumi.Context, args *LookupNotificationRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupNotificationRegistrationResult, error) {
	var rv LookupNotificationRegistrationResult
	err := ctx.Invoke("azure-native:providerhub/v20210901preview:getNotificationRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationRegistrationArgs struct {
	NotificationRegistrationName string `pulumi:"notificationRegistrationName"`
	ProviderNamespace            string `pulumi:"providerNamespace"`
}


type LookupNotificationRegistrationResult struct {
	Id         string                                     `pulumi:"id"`
	Name       string                                     `pulumi:"name"`
	Properties NotificationRegistrationResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse                         `pulumi:"systemData"`
	Type       string                                     `pulumi:"type"`
}
