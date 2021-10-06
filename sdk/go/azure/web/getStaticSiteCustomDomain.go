


package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSiteCustomDomain(ctx *pulumi.Context, args *LookupStaticSiteCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteCustomDomainResult, error) {
	var rv LookupStaticSiteCustomDomainResult
	err := ctx.Invoke("azure-native:web:getStaticSiteCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteCustomDomainArgs struct {
	DomainName        string `pulumi:"domainName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteCustomDomainResult struct {
	CreatedOn       string  `pulumi:"createdOn"`
	DomainName      string  `pulumi:"domainName"`
	ErrorMessage    string  `pulumi:"errorMessage"`
	Id              string  `pulumi:"id"`
	Kind            *string `pulumi:"kind"`
	Name            string  `pulumi:"name"`
	Status          string  `pulumi:"status"`
	Type            string  `pulumi:"type"`
	ValidationToken string  `pulumi:"validationToken"`
}
