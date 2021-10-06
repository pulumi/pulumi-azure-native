


package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountAgreement(ctx *pulumi.Context, args *LookupIntegrationAccountAgreementArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountAgreementResult, error) {
	var rv LookupIntegrationAccountAgreementResult
	err := ctx.Invoke("azure-native:logic/v20190501:getIntegrationAccountAgreement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountAgreementArgs struct {
	AgreementName          string `pulumi:"agreementName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountAgreementResult struct {
	AgreementType string                   `pulumi:"agreementType"`
	ChangedTime   string                   `pulumi:"changedTime"`
	Content       AgreementContentResponse `pulumi:"content"`
	CreatedTime   string                   `pulumi:"createdTime"`
	GuestIdentity BusinessIdentityResponse `pulumi:"guestIdentity"`
	GuestPartner  string                   `pulumi:"guestPartner"`
	HostIdentity  BusinessIdentityResponse `pulumi:"hostIdentity"`
	HostPartner   string                   `pulumi:"hostPartner"`
	Id            string                   `pulumi:"id"`
	Location      *string                  `pulumi:"location"`
	Metadata      interface{}              `pulumi:"metadata"`
	Name          string                   `pulumi:"name"`
	Tags          map[string]string        `pulumi:"tags"`
	Type          string                   `pulumi:"type"`
}
