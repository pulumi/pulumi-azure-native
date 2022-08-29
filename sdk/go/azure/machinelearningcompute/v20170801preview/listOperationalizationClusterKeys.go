


package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOperationalizationClusterKeys(ctx *pulumi.Context, args *ListOperationalizationClusterKeysArgs, opts ...pulumi.InvokeOption) (*ListOperationalizationClusterKeysResult, error) {
	var rv ListOperationalizationClusterKeysResult
	err := ctx.Invoke("azure-native:machinelearningcompute/v20170801preview:listOperationalizationClusterKeys", args, &rv, opts...)
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

func ListOperationalizationClusterKeysOutput(ctx *pulumi.Context, args ListOperationalizationClusterKeysOutputArgs, opts ...pulumi.InvokeOption) ListOperationalizationClusterKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOperationalizationClusterKeysResult, error) {
			args := v.(ListOperationalizationClusterKeysArgs)
			r, err := ListOperationalizationClusterKeys(ctx, &args, opts...)
			var s ListOperationalizationClusterKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOperationalizationClusterKeysResultOutput)
}

type ListOperationalizationClusterKeysOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListOperationalizationClusterKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOperationalizationClusterKeysArgs)(nil)).Elem()
}


type ListOperationalizationClusterKeysResultOutput struct{ *pulumi.OutputState }

func (ListOperationalizationClusterKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOperationalizationClusterKeysResult)(nil)).Elem()
}

func (o ListOperationalizationClusterKeysResultOutput) ToListOperationalizationClusterKeysResultOutput() ListOperationalizationClusterKeysResultOutput {
	return o
}

func (o ListOperationalizationClusterKeysResultOutput) ToListOperationalizationClusterKeysResultOutputWithContext(ctx context.Context) ListOperationalizationClusterKeysResultOutput {
	return o
}

func (o ListOperationalizationClusterKeysResultOutput) AppInsights() AppInsightsCredentialsResponsePtrOutput {
	return o.ApplyT(func(v ListOperationalizationClusterKeysResult) *AppInsightsCredentialsResponse { return v.AppInsights }).(AppInsightsCredentialsResponsePtrOutput)
}

func (o ListOperationalizationClusterKeysResultOutput) ContainerRegistry() ContainerRegistryCredentialsResponsePtrOutput {
	return o.ApplyT(func(v ListOperationalizationClusterKeysResult) *ContainerRegistryCredentialsResponse {
		return v.ContainerRegistry
	}).(ContainerRegistryCredentialsResponsePtrOutput)
}

func (o ListOperationalizationClusterKeysResultOutput) ContainerService() ContainerServiceCredentialsResponsePtrOutput {
	return o.ApplyT(func(v ListOperationalizationClusterKeysResult) *ContainerServiceCredentialsResponse {
		return v.ContainerService
	}).(ContainerServiceCredentialsResponsePtrOutput)
}

func (o ListOperationalizationClusterKeysResultOutput) ServiceAuthConfiguration() ServiceAuthConfigurationResponsePtrOutput {
	return o.ApplyT(func(v ListOperationalizationClusterKeysResult) *ServiceAuthConfigurationResponse {
		return v.ServiceAuthConfiguration
	}).(ServiceAuthConfigurationResponsePtrOutput)
}

func (o ListOperationalizationClusterKeysResultOutput) SslConfiguration() SslConfigurationResponsePtrOutput {
	return o.ApplyT(func(v ListOperationalizationClusterKeysResult) *SslConfigurationResponse { return v.SslConfiguration }).(SslConfigurationResponsePtrOutput)
}

func (o ListOperationalizationClusterKeysResultOutput) StorageAccount() StorageAccountCredentialsResponsePtrOutput {
	return o.ApplyT(func(v ListOperationalizationClusterKeysResult) *StorageAccountCredentialsResponse {
		return v.StorageAccount
	}).(StorageAccountCredentialsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOperationalizationClusterKeysResultOutput{})
}
