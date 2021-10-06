


package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountPartner(ctx *pulumi.Context, args *LookupIntegrationAccountPartnerArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountPartnerResult, error) {
	var rv LookupIntegrationAccountPartnerResult
	err := ctx.Invoke("azure-native:logic:getIntegrationAccountPartner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountPartnerArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	PartnerName            string `pulumi:"partnerName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountPartnerResult struct {
	ChangedTime string                 `pulumi:"changedTime"`
	Content     PartnerContentResponse `pulumi:"content"`
	CreatedTime string                 `pulumi:"createdTime"`
	Id          string                 `pulumi:"id"`
	Location    *string                `pulumi:"location"`
	Metadata    interface{}            `pulumi:"metadata"`
	Name        string                 `pulumi:"name"`
	PartnerType string                 `pulumi:"partnerType"`
	Tags        map[string]string      `pulumi:"tags"`
	Type        string                 `pulumi:"type"`
}
