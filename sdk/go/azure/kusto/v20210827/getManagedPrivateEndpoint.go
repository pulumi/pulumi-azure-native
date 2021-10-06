


package v20210827

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedPrivateEndpoint(ctx *pulumi.Context, args *LookupManagedPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupManagedPrivateEndpointResult, error) {
	var rv LookupManagedPrivateEndpointResult
	err := ctx.Invoke("azure-native:kusto/v20210827:getManagedPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedPrivateEndpointArgs struct {
	ClusterName                string `pulumi:"clusterName"`
	ManagedPrivateEndpointName string `pulumi:"managedPrivateEndpointName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}


type LookupManagedPrivateEndpointResult struct {
	GroupId                   string             `pulumi:"groupId"`
	Id                        string             `pulumi:"id"`
	Name                      string             `pulumi:"name"`
	PrivateLinkResourceId     string             `pulumi:"privateLinkResourceId"`
	PrivateLinkResourceRegion *string            `pulumi:"privateLinkResourceRegion"`
	ProvisioningState         string             `pulumi:"provisioningState"`
	RequestMessage            *string            `pulumi:"requestMessage"`
	SystemData                SystemDataResponse `pulumi:"systemData"`
	Type                      string             `pulumi:"type"`
}
