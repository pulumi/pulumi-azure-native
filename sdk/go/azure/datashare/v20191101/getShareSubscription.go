


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupShareSubscription(ctx *pulumi.Context, args *LookupShareSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupShareSubscriptionResult, error) {
	var rv LookupShareSubscriptionResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getShareSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupShareSubscriptionArgs struct {
	AccountName           string `pulumi:"accountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupShareSubscriptionResult struct {
	CreatedAt               string `pulumi:"createdAt"`
	Id                      string `pulumi:"id"`
	InvitationId            string `pulumi:"invitationId"`
	Name                    string `pulumi:"name"`
	ProviderEmail           string `pulumi:"providerEmail"`
	ProviderName            string `pulumi:"providerName"`
	ProviderTenantName      string `pulumi:"providerTenantName"`
	ProvisioningState       string `pulumi:"provisioningState"`
	ShareDescription        string `pulumi:"shareDescription"`
	ShareKind               string `pulumi:"shareKind"`
	ShareName               string `pulumi:"shareName"`
	ShareSubscriptionStatus string `pulumi:"shareSubscriptionStatus"`
	ShareTerms              string `pulumi:"shareTerms"`
	SourceShareLocation     string `pulumi:"sourceShareLocation"`
	Type                    string `pulumi:"type"`
	UserEmail               string `pulumi:"userEmail"`
	UserName                string `pulumi:"userName"`
}
