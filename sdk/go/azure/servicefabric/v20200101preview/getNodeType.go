


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNodeType(ctx *pulumi.Context, args *LookupNodeTypeArgs, opts ...pulumi.InvokeOption) (*LookupNodeTypeResult, error) {
	var rv LookupNodeTypeResult
	err := ctx.Invoke("azure-native:servicefabric/v20200101preview:getNodeType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNodeTypeArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	NodeTypeName      string `pulumi:"nodeTypeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNodeTypeResult struct {
	ApplicationPorts    *EndpointRangeDescriptionResponse `pulumi:"applicationPorts"`
	Capacities          map[string]string                 `pulumi:"capacities"`
	DataDiskSizeGB      int                               `pulumi:"dataDiskSizeGB"`
	EphemeralPorts      *EndpointRangeDescriptionResponse `pulumi:"ephemeralPorts"`
	Id                  string                            `pulumi:"id"`
	IsPrimary           bool                              `pulumi:"isPrimary"`
	Name                string                            `pulumi:"name"`
	PlacementProperties map[string]string                 `pulumi:"placementProperties"`
	ProvisioningState   string                            `pulumi:"provisioningState"`
	Tags                map[string]string                 `pulumi:"tags"`
	Type                string                            `pulumi:"type"`
	VmExtensions        []VMSSExtensionResponse           `pulumi:"vmExtensions"`
	VmImageOffer        *string                           `pulumi:"vmImageOffer"`
	VmImagePublisher    *string                           `pulumi:"vmImagePublisher"`
	VmImageSku          *string                           `pulumi:"vmImageSku"`
	VmImageVersion      *string                           `pulumi:"vmImageVersion"`
	VmInstanceCount     int                               `pulumi:"vmInstanceCount"`
	VmSecrets           []VaultSecretGroupResponse        `pulumi:"vmSecrets"`
	VmSize              *string                           `pulumi:"vmSize"`
}
