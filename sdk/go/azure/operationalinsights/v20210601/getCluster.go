


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:operationalinsights/v20210601:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupClusterResult struct {
	AssociatedWorkspaces          []AssociatedWorkspaceResponse          `pulumi:"associatedWorkspaces"`
	BillingType                   *string                                `pulumi:"billingType"`
	CapacityReservationProperties *CapacityReservationPropertiesResponse `pulumi:"capacityReservationProperties"`
	ClusterId                     string                                 `pulumi:"clusterId"`
	CreatedDate                   string                                 `pulumi:"createdDate"`
	Id                            string                                 `pulumi:"id"`
	Identity                      *IdentityResponse                      `pulumi:"identity"`
	IsAvailabilityZonesEnabled    *bool                                  `pulumi:"isAvailabilityZonesEnabled"`
	KeyVaultProperties            *KeyVaultPropertiesResponse            `pulumi:"keyVaultProperties"`
	LastModifiedDate              string                                 `pulumi:"lastModifiedDate"`
	Location                      string                                 `pulumi:"location"`
	Name                          string                                 `pulumi:"name"`
	ProvisioningState             string                                 `pulumi:"provisioningState"`
	Sku                           *ClusterSkuResponse                    `pulumi:"sku"`
	Tags                          map[string]string                      `pulumi:"tags"`
	Type                          string                                 `pulumi:"type"`
}
