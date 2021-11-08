


package v20210701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDedicatedHost(ctx *pulumi.Context, args *LookupDedicatedHostArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostResult, error) {
	var rv LookupDedicatedHostResult
	err := ctx.Invoke("azure-native:compute/v20210701:getDedicatedHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHostArgs struct {
	Expand            *string `pulumi:"expand"`
	HostGroupName     string  `pulumi:"hostGroupName"`
	HostName          string  `pulumi:"hostName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupDedicatedHostResult struct {
	AutoReplaceOnFailure *bool                             `pulumi:"autoReplaceOnFailure"`
	HostId               string                            `pulumi:"hostId"`
	Id                   string                            `pulumi:"id"`
	InstanceView         DedicatedHostInstanceViewResponse `pulumi:"instanceView"`
	LicenseType          *string                           `pulumi:"licenseType"`
	Location             string                            `pulumi:"location"`
	Name                 string                            `pulumi:"name"`
	PlatformFaultDomain  *int                              `pulumi:"platformFaultDomain"`
	ProvisioningState    string                            `pulumi:"provisioningState"`
	ProvisioningTime     string                            `pulumi:"provisioningTime"`
	Sku                  SkuResponse                       `pulumi:"sku"`
	Tags                 map[string]string                 `pulumi:"tags"`
	Type                 string                            `pulumi:"type"`
	VirtualMachines      []SubResourceReadOnlyResponse     `pulumi:"virtualMachines"`
}
