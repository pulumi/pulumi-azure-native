


package v20170601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupOperationalizationCluster(ctx *pulumi.Context, args *LookupOperationalizationClusterArgs, opts ...pulumi.InvokeOption) (*LookupOperationalizationClusterResult, error) {
	var rv LookupOperationalizationClusterResult
	err := ctx.Invoke("azure-native:machinelearningcompute/v20170601preview:getOperationalizationCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupOperationalizationClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOperationalizationClusterResult struct {
	AppInsights                *AppInsightsCredentialsResponse      `pulumi:"appInsights"`
	ClusterType                string                               `pulumi:"clusterType"`
	ContainerRegistry          *ContainerRegistryPropertiesResponse `pulumi:"containerRegistry"`
	ContainerService           AcsClusterPropertiesResponse         `pulumi:"containerService"`
	CreatedOn                  string                               `pulumi:"createdOn"`
	Description                *string                              `pulumi:"description"`
	GlobalServiceConfiguration *GlobalServiceConfigurationResponse  `pulumi:"globalServiceConfiguration"`
	Id                         string                               `pulumi:"id"`
	Location                   string                               `pulumi:"location"`
	ModifiedOn                 string                               `pulumi:"modifiedOn"`
	Name                       string                               `pulumi:"name"`
	ProvisioningState          string                               `pulumi:"provisioningState"`
	StorageAccount             *StorageAccountPropertiesResponse    `pulumi:"storageAccount"`
	Tags                       map[string]string                    `pulumi:"tags"`
	Type                       string                               `pulumi:"type"`
}


func (val *LookupOperationalizationClusterResult) Defaults() *LookupOperationalizationClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ContainerService = *tmp.ContainerService.Defaults()

	tmp.GlobalServiceConfiguration = tmp.GlobalServiceConfiguration.Defaults()

	return &tmp
}

func LookupOperationalizationClusterOutput(ctx *pulumi.Context, args LookupOperationalizationClusterOutputArgs, opts ...pulumi.InvokeOption) LookupOperationalizationClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOperationalizationClusterResult, error) {
			args := v.(LookupOperationalizationClusterArgs)
			r, err := LookupOperationalizationCluster(ctx, &args, opts...)
			var s LookupOperationalizationClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOperationalizationClusterResultOutput)
}

type LookupOperationalizationClusterOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOperationalizationClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOperationalizationClusterArgs)(nil)).Elem()
}


type LookupOperationalizationClusterResultOutput struct{ *pulumi.OutputState }

func (LookupOperationalizationClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOperationalizationClusterResult)(nil)).Elem()
}

func (o LookupOperationalizationClusterResultOutput) ToLookupOperationalizationClusterResultOutput() LookupOperationalizationClusterResultOutput {
	return o
}

func (o LookupOperationalizationClusterResultOutput) ToLookupOperationalizationClusterResultOutputWithContext(ctx context.Context) LookupOperationalizationClusterResultOutput {
	return o
}

func (o LookupOperationalizationClusterResultOutput) AppInsights() AppInsightsCredentialsResponsePtrOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) *AppInsightsCredentialsResponse { return v.AppInsights }).(AppInsightsCredentialsResponsePtrOutput)
}

func (o LookupOperationalizationClusterResultOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) string { return v.ClusterType }).(pulumi.StringOutput)
}

func (o LookupOperationalizationClusterResultOutput) ContainerRegistry() ContainerRegistryPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) *ContainerRegistryPropertiesResponse {
		return v.ContainerRegistry
	}).(ContainerRegistryPropertiesResponsePtrOutput)
}

func (o LookupOperationalizationClusterResultOutput) ContainerService() AcsClusterPropertiesResponseOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) AcsClusterPropertiesResponse { return v.ContainerService }).(AcsClusterPropertiesResponseOutput)
}

func (o LookupOperationalizationClusterResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupOperationalizationClusterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupOperationalizationClusterResultOutput) GlobalServiceConfiguration() GlobalServiceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) *GlobalServiceConfigurationResponse {
		return v.GlobalServiceConfiguration
	}).(GlobalServiceConfigurationResponsePtrOutput)
}

func (o LookupOperationalizationClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOperationalizationClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupOperationalizationClusterResultOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o LookupOperationalizationClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOperationalizationClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupOperationalizationClusterResultOutput) StorageAccount() StorageAccountPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) *StorageAccountPropertiesResponse {
		return v.StorageAccount
	}).(StorageAccountPropertiesResponsePtrOutput)
}

func (o LookupOperationalizationClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOperationalizationClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOperationalizationClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOperationalizationClusterResultOutput{})
}
