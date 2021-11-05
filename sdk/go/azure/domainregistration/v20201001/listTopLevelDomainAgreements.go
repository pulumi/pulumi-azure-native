


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTopLevelDomainAgreements(ctx *pulumi.Context, args *ListTopLevelDomainAgreementsArgs, opts ...pulumi.InvokeOption) (*ListTopLevelDomainAgreementsResult, error) {
	var rv ListTopLevelDomainAgreementsResult
	err := ctx.Invoke("azure-native:domainregistration/v20201001:listTopLevelDomainAgreements", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTopLevelDomainAgreementsArgs struct {
	ForTransfer    *bool  `pulumi:"forTransfer"`
	IncludePrivacy *bool  `pulumi:"includePrivacy"`
	Name           string `pulumi:"name"`
}


type ListTopLevelDomainAgreementsResult struct {
	NextLink string                      `pulumi:"nextLink"`
	Value    []TldLegalAgreementResponse `pulumi:"value"`
}
