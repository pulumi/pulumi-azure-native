


package v20170601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomerSubscription(ctx *pulumi.Context, args *LookupCustomerSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupCustomerSubscriptionResult, error) {
	var rv LookupCustomerSubscriptionResult
	err := ctx.Invoke("azure-native:azurestack/v20170601:getCustomerSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomerSubscriptionArgs struct {
	CustomerSubscriptionName string `pulumi:"customerSubscriptionName"`
	RegistrationName         string `pulumi:"registrationName"`
	ResourceGroup            string `pulumi:"resourceGroup"`
}


type LookupCustomerSubscriptionResult struct {
	Etag     *string `pulumi:"etag"`
	Id       string  `pulumi:"id"`
	Name     string  `pulumi:"name"`
	TenantId *string `pulumi:"tenantId"`
	Type     string  `pulumi:"type"`
}
