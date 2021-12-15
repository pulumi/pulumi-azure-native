


package v20201201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:servicefabric/v20201201preview:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupServiceResult struct {
	CorrelationScheme            []ServiceCorrelationDescriptionResponse     `pulumi:"correlationScheme"`
	DefaultMoveCost              *string                                     `pulumi:"defaultMoveCost"`
	Etag                         string                                      `pulumi:"etag"`
	Id                           string                                      `pulumi:"id"`
	Location                     *string                                     `pulumi:"location"`
	Name                         string                                      `pulumi:"name"`
	PartitionDescription         interface{}                                 `pulumi:"partitionDescription"`
	PlacementConstraints         *string                                     `pulumi:"placementConstraints"`
	ProvisioningState            string                                      `pulumi:"provisioningState"`
	ServiceDnsName               *string                                     `pulumi:"serviceDnsName"`
	ServiceKind                  string                                      `pulumi:"serviceKind"`
	ServiceLoadMetrics           []ServiceLoadMetricDescriptionResponse      `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode *string                                     `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     []ServicePlacementPolicyDescriptionResponse `pulumi:"servicePlacementPolicies"`
	ServiceTypeName              *string                                     `pulumi:"serviceTypeName"`
	SystemData                   SystemDataResponse                          `pulumi:"systemData"`
	Tags                         map[string]string                           `pulumi:"tags"`
	Type                         string                                      `pulumi:"type"`
}
