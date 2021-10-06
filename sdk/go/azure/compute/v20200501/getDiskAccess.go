


package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiskAccess(ctx *pulumi.Context, args *LookupDiskAccessArgs, opts ...pulumi.InvokeOption) (*LookupDiskAccessResult, error) {
	var rv LookupDiskAccessResult
	err := ctx.Invoke("azure-native:compute/v20200501:getDiskAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskAccessArgs struct {
	DiskAccessName    string `pulumi:"diskAccessName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDiskAccessResult struct {
	Id                         string                              `pulumi:"id"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	Tags                       map[string]string                   `pulumi:"tags"`
	TimeCreated                string                              `pulumi:"timeCreated"`
	Type                       string                              `pulumi:"type"`
}
