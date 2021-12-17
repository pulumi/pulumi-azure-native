


package v20170601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOperationalizationClusterKeys(ctx *pulumi.Context, args *ListOperationalizationClusterKeysArgs, opts ...pulumi.InvokeOption) (*ListOperationalizationClusterKeysResult, error) {
	var rv ListOperationalizationClusterKeysResult
	err := ctx.Invoke("azure-native:machinelearningcompute/v20170601preview:listOperationalizationClusterKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListOperationalizationClusterKeysArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListOperationalizationClusterKeysResult struct {
	AppInsights              *AppInsightsCredentialsResponse       `pulumi:"appInsights"`
	ContainerRegistry        *ContainerRegistryCredentialsResponse `pulumi:"containerRegistry"`
	ContainerService         *ContainerServiceCredentialsResponse  `pulumi:"containerService"`
	ServiceAuthConfiguration *ServiceAuthConfigurationResponse     `pulumi:"serviceAuthConfiguration"`
	SslConfiguration         *SslConfigurationResponse             `pulumi:"sslConfiguration"`
	StorageAccount           *StorageAccountCredentialsResponse    `pulumi:"storageAccount"`
}


func (val *ListOperationalizationClusterKeysResult) Defaults() *ListOperationalizationClusterKeysResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SslConfiguration = tmp.SslConfiguration.Defaults()

	return &tmp
}
