


package v20201216preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomerEvent(ctx *pulumi.Context, args *LookupCustomerEventArgs, opts ...pulumi.InvokeOption) (*LookupCustomerEventResult, error) {
	var rv LookupCustomerEventResult
	err := ctx.Invoke("azure-native:testbase/v20201216preview:getCustomerEvent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomerEventArgs struct {
	CustomerEventName   string `pulumi:"customerEventName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}


type LookupCustomerEventResult struct {
	EventName  string                              `pulumi:"eventName"`
	Id         string                              `pulumi:"id"`
	Name       string                              `pulumi:"name"`
	Receivers  []NotificationEventReceiverResponse `pulumi:"receivers"`
	SystemData SystemDataResponse                  `pulumi:"systemData"`
	Type       string                              `pulumi:"type"`
}
