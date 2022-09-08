


package machinelearningcompute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OperationalizationCluster struct {
	pulumi.CustomResourceState

	AppInsights                AppInsightsPropertiesResponsePtrOutput       `pulumi:"appInsights"`
	ClusterType                pulumi.StringOutput                          `pulumi:"clusterType"`
	ContainerRegistry          ContainerRegistryPropertiesResponsePtrOutput `pulumi:"containerRegistry"`
	ContainerService           AcsClusterPropertiesResponsePtrOutput        `pulumi:"containerService"`
	CreatedOn                  pulumi.StringOutput                          `pulumi:"createdOn"`
	Description                pulumi.StringPtrOutput                       `pulumi:"description"`
	GlobalServiceConfiguration GlobalServiceConfigurationResponsePtrOutput  `pulumi:"globalServiceConfiguration"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	ModifiedOn                 pulumi.StringOutput                          `pulumi:"modifiedOn"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	ProvisioningErrors         ErrorResponseWrapperResponseArrayOutput      `pulumi:"provisioningErrors"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	StorageAccount             StorageAccountPropertiesResponsePtrOutput    `pulumi:"storageAccount"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
}


func NewOperationalizationCluster(ctx *pulumi.Context,
	name string, args *OperationalizationClusterArgs, opts ...pulumi.ResourceOption) (*OperationalizationCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterType == nil {
		return nil, errors.New("invalid value for required argument 'ClusterType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ContainerService != nil {
		args.ContainerService = args.ContainerService.ToAcsClusterPropertiesPtrOutput().ApplyT(func(v *AcsClusterProperties) *AcsClusterProperties { return v.Defaults() }).(AcsClusterPropertiesPtrOutput)
	}
	if args.GlobalServiceConfiguration != nil {
		args.GlobalServiceConfiguration = args.GlobalServiceConfiguration.ToGlobalServiceConfigurationPtrOutput().ApplyT(func(v *GlobalServiceConfiguration) *GlobalServiceConfiguration { return v.Defaults() }).(GlobalServiceConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningcompute/v20170601preview:OperationalizationCluster"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningcompute/v20170801preview:OperationalizationCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource OperationalizationCluster
	err := ctx.RegisterResource("azure-native:machinelearningcompute:OperationalizationCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOperationalizationCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OperationalizationClusterState, opts ...pulumi.ResourceOption) (*OperationalizationCluster, error) {
	var resource OperationalizationCluster
	err := ctx.ReadResource("azure-native:machinelearningcompute:OperationalizationCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type operationalizationClusterState struct {
}

type OperationalizationClusterState struct {
}

func (OperationalizationClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*operationalizationClusterState)(nil)).Elem()
}

type operationalizationClusterArgs struct {
	AppInsights                *AppInsightsProperties       `pulumi:"appInsights"`
	ClusterName                *string                      `pulumi:"clusterName"`
	ClusterType                string                       `pulumi:"clusterType"`
	ContainerRegistry          *ContainerRegistryProperties `pulumi:"containerRegistry"`
	ContainerService           *AcsClusterProperties        `pulumi:"containerService"`
	Description                *string                      `pulumi:"description"`
	GlobalServiceConfiguration *GlobalServiceConfiguration  `pulumi:"globalServiceConfiguration"`
	Location                   *string                      `pulumi:"location"`
	ResourceGroupName          string                       `pulumi:"resourceGroupName"`
	StorageAccount             *StorageAccountProperties    `pulumi:"storageAccount"`
	Tags                       map[string]string            `pulumi:"tags"`
}


type OperationalizationClusterArgs struct {
	AppInsights                AppInsightsPropertiesPtrInput
	ClusterName                pulumi.StringPtrInput
	ClusterType                pulumi.StringInput
	ContainerRegistry          ContainerRegistryPropertiesPtrInput
	ContainerService           AcsClusterPropertiesPtrInput
	Description                pulumi.StringPtrInput
	GlobalServiceConfiguration GlobalServiceConfigurationPtrInput
	Location                   pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	StorageAccount             StorageAccountPropertiesPtrInput
	Tags                       pulumi.StringMapInput
}

func (OperationalizationClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*operationalizationClusterArgs)(nil)).Elem()
}

type OperationalizationClusterInput interface {
	pulumi.Input

	ToOperationalizationClusterOutput() OperationalizationClusterOutput
	ToOperationalizationClusterOutputWithContext(ctx context.Context) OperationalizationClusterOutput
}

func (*OperationalizationCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationalizationCluster)(nil)).Elem()
}

func (i *OperationalizationCluster) ToOperationalizationClusterOutput() OperationalizationClusterOutput {
	return i.ToOperationalizationClusterOutputWithContext(context.Background())
}

func (i *OperationalizationCluster) ToOperationalizationClusterOutputWithContext(ctx context.Context) OperationalizationClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationalizationClusterOutput)
}

type OperationalizationClusterOutput struct{ *pulumi.OutputState }

func (OperationalizationClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationalizationCluster)(nil)).Elem()
}

func (o OperationalizationClusterOutput) ToOperationalizationClusterOutput() OperationalizationClusterOutput {
	return o
}

func (o OperationalizationClusterOutput) ToOperationalizationClusterOutputWithContext(ctx context.Context) OperationalizationClusterOutput {
	return o
}

func (o OperationalizationClusterOutput) AppInsights() AppInsightsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) AppInsightsPropertiesResponsePtrOutput { return v.AppInsights }).(AppInsightsPropertiesResponsePtrOutput)
}

func (o OperationalizationClusterOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

func (o OperationalizationClusterOutput) ContainerRegistry() ContainerRegistryPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) ContainerRegistryPropertiesResponsePtrOutput {
		return v.ContainerRegistry
	}).(ContainerRegistryPropertiesResponsePtrOutput)
}

func (o OperationalizationClusterOutput) ContainerService() AcsClusterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) AcsClusterPropertiesResponsePtrOutput { return v.ContainerService }).(AcsClusterPropertiesResponsePtrOutput)
}

func (o OperationalizationClusterOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o OperationalizationClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o OperationalizationClusterOutput) GlobalServiceConfiguration() GlobalServiceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) GlobalServiceConfigurationResponsePtrOutput {
		return v.GlobalServiceConfiguration
	}).(GlobalServiceConfigurationResponsePtrOutput)
}

func (o OperationalizationClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o OperationalizationClusterOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o OperationalizationClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OperationalizationClusterOutput) ProvisioningErrors() ErrorResponseWrapperResponseArrayOutput {
	return o.ApplyT(func(v *OperationalizationCluster) ErrorResponseWrapperResponseArrayOutput {
		return v.ProvisioningErrors
	}).(ErrorResponseWrapperResponseArrayOutput)
}

func (o OperationalizationClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OperationalizationClusterOutput) StorageAccount() StorageAccountPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) StorageAccountPropertiesResponsePtrOutput { return v.StorageAccount }).(StorageAccountPropertiesResponsePtrOutput)
}

func (o OperationalizationClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OperationalizationClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OperationalizationClusterOutput{})
}
