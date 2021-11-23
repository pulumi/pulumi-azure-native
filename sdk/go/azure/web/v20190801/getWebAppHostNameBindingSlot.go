


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppHostNameBindingSlot(ctx *pulumi.Context, args *LookupWebAppHostNameBindingSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppHostNameBindingSlotResult, error) {
	var rv LookupWebAppHostNameBindingSlotResult
	err := ctx.Invoke("azure-native:web/v20190801:getWebAppHostNameBindingSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppHostNameBindingSlotArgs struct {
	HostName          string `pulumi:"hostName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppHostNameBindingSlotResult struct {
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
