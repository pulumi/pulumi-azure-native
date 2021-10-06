


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSiteHostNameBindingSlot(ctx *pulumi.Context, args *LookupSiteHostNameBindingSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteHostNameBindingSlotResult, error) {
	var rv LookupSiteHostNameBindingSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteHostNameBindingSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteHostNameBindingSlotArgs struct {
	HostName          string `pulumi:"hostName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupSiteHostNameBindingSlotResult struct {
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
