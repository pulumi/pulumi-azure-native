


package machinelearningcompute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOperationalizationCluster(ctx *pulumi.Context, args *LookupOperationalizationClusterArgs, opts ...pulumi.InvokeOption) (*LookupOperationalizationClusterResult, error) {
	var rv LookupOperationalizationClusterResult
	err := ctx.Invoke("azure-native:machinelearningcompute:getOperationalizationCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOperationalizationClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOperationalizationClusterResult struct {
	AppInsights                *AppInsightsPropertiesResponse       `pulumi:"appInsights"`
	ClusterType                string                               `pulumi:"clusterType"`
	ContainerRegistry          *ContainerRegistryPropertiesResponse `pulumi:"containerRegistry"`
	ContainerService           *AcsClusterPropertiesResponse        `pulumi:"containerService"`
	CreatedOn                  string                               `pulumi:"createdOn"`
	Description                *string                              `pulumi:"description"`
	GlobalServiceConfiguration *GlobalServiceConfigurationResponse  `pulumi:"globalServiceConfiguration"`
	Id                         string                               `pulumi:"id"`
	Location                   string                               `pulumi:"location"`
	ModifiedOn                 string                               `pulumi:"modifiedOn"`
	Name                       string                               `pulumi:"name"`
	ProvisioningErrors         []ErrorResponseWrapperResponse       `pulumi:"provisioningErrors"`
	ProvisioningState          string                               `pulumi:"provisioningState"`
	StorageAccount             *StorageAccountPropertiesResponse    `pulumi:"storageAccount"`
	Tags                       map[string]string                    `pulumi:"tags"`
	Type                       string                               `pulumi:"type"`
}
