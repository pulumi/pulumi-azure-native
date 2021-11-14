


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSiteHostNameBinding(ctx *pulumi.Context, args *LookupSiteHostNameBindingArgs, opts ...pulumi.InvokeOption) (*LookupSiteHostNameBindingResult, error) {
	var rv LookupSiteHostNameBindingResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteHostNameBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteHostNameBindingArgs struct {
	HostName          string `pulumi:"hostName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteHostNameBindingResult struct {
	AzureResourceName           *string           `pulumi:"azureResourceName"`
	AzureResourceType           *string           `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType *string           `pulumi:"customHostNameDnsRecordType"`
	DomainId                    *string           `pulumi:"domainId"`
	HostNameType                *string           `pulumi:"hostNameType"`
	Id                          *string           `pulumi:"id"`
	Kind                        *string           `pulumi:"kind"`
	Location                    string            `pulumi:"location"`
	Name                        *string           `pulumi:"name"`
	SiteName                    *string           `pulumi:"siteName"`
	Tags                        map[string]string `pulumi:"tags"`
	Type                        *string           `pulumi:"type"`
}
