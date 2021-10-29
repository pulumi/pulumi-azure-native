


package v20171012

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomDomain(ctx *pulumi.Context, args *LookupCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupCustomDomainResult, error) {
	var rv LookupCustomDomainResult
	err := ctx.Invoke("azure-native:cdn/v20171012:getCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomDomainArgs struct {
	CustomDomainName  string `pulumi:"customDomainName"`
	EndpointName      string `pulumi:"endpointName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCustomDomainResult struct {
	CustomHttpsProvisioningState    string  `pulumi:"customHttpsProvisioningState"`
	CustomHttpsProvisioningSubstate string  `pulumi:"customHttpsProvisioningSubstate"`
	HostName                        string  `pulumi:"hostName"`
	Id                              string  `pulumi:"id"`
	Name                            string  `pulumi:"name"`
	ProvisioningState               string  `pulumi:"provisioningState"`
	ResourceState                   string  `pulumi:"resourceState"`
	Type                            string  `pulumi:"type"`
	ValidationData                  *string `pulumi:"validationData"`
}
