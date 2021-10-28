


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppHostNameBinding(ctx *pulumi.Context, args *LookupWebAppHostNameBindingArgs, opts ...pulumi.InvokeOption) (*LookupWebAppHostNameBindingResult, error) {
	var rv LookupWebAppHostNameBindingResult
	err := ctx.Invoke("azure-native:web/v20200601:getWebAppHostNameBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppHostNameBindingArgs struct {
	HostName          string `pulumi:"hostName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppHostNameBindingResult struct {
	AzureResourceName           *string `pulumi:"azureResourceName"`
	AzureResourceType           *string `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType *string `pulumi:"customHostNameDnsRecordType"`
	DomainId                    *string `pulumi:"domainId"`
	HostNameType                *string `pulumi:"hostNameType"`
	Id                          string  `pulumi:"id"`
	Kind                        *string `pulumi:"kind"`
	Name                        string  `pulumi:"name"`
	SiteName                    *string `pulumi:"siteName"`
	SslState                    *string `pulumi:"sslState"`
	Thumbprint                  *string `pulumi:"thumbprint"`
	Type                        string  `pulumi:"type"`
	VirtualIP                   string  `pulumi:"virtualIP"`
}
