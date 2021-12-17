


package v20170601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OperationalizationCluster struct {
	pulumi.CustomResourceState

	AppInsights                AppInsightsCredentialsResponsePtrOutput      `pulumi:"appInsights"`
	ClusterType                pulumi.StringOutput                          `pulumi:"clusterType"`
	ContainerRegistry          ContainerRegistryPropertiesResponsePtrOutput `pulumi:"containerRegistry"`
	ContainerService           AcsClusterPropertiesResponseOutput           `pulumi:"containerService"`
	CreatedOn                  pulumi.StringOutput                          `pulumi:"createdOn"`
	Description                pulumi.StringPtrOutput                       `pulumi:"description"`
	GlobalServiceConfiguration GlobalServiceConfigurationResponsePtrOutput  `pulumi:"globalServiceConfiguration"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	ModifiedOn                 pulumi.StringOutput                          `pulumi:"modifiedOn"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
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
	if args.ContainerService == nil {
		return nil, errors.New("invalid value for required argument 'ContainerService'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	containerServiceApplier := func(v AcsClusterProperties) *AcsClusterProperties { return v.Defaults() }
	args.ContainerService = args.ContainerService.ToAcsClusterPropertiesOutput().ApplyT(containerServiceApplier).(AcsClusterPropertiesPtrOutput).Elem()
	globalServiceConfigurationApplier := func(v GlobalServiceConfiguration) *GlobalServiceConfiguration { return v.Defaults() }
	if args.GlobalServiceConfiguration != nil {
		args.GlobalServiceConfiguration = args.GlobalServiceConfiguration.ToGlobalServiceConfigurationPtrOutput().Elem().ApplyT(globalServiceConfigurationApplier).(GlobalServiceConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningcompute:OperationalizationCluster"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningcompute/v20170801preview:OperationalizationCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource OperationalizationCluster
	err := ctx.RegisterResource("azure-native:machinelearningcompute/v20170601preview:OperationalizationCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOperationalizationCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OperationalizationClusterState, opts ...pulumi.ResourceOption) (*OperationalizationCluster, error) {
	var resource OperationalizationCluster
	err := ctx.ReadResource("azure-native:machinelearningcompute/v20170601preview:OperationalizationCluster", name, id, state, &resource, opts...)
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
	AppInsights                *AppInsightsCredentials      `pulumi:"appInsights"`
	ClusterName                *string                      `pulumi:"clusterName"`
	ClusterType                string                       `pulumi:"clusterType"`
	ContainerRegistry          *ContainerRegistryProperties `pulumi:"containerRegistry"`
	ContainerService           AcsClusterProperties         `pulumi:"containerService"`
	Description                *string                      `pulumi:"description"`
	GlobalServiceConfiguration *GlobalServiceConfiguration  `pulumi:"globalServiceConfiguration"`
	Location                   *string                      `pulumi:"location"`
	ResourceGroupName          string                       `pulumi:"resourceGroupName"`
	StorageAccount             *StorageAccountProperties    `pulumi:"storageAccount"`
	Tags                       map[string]string            `pulumi:"tags"`
}


type OperationalizationClusterArgs struct {
	AppInsights                AppInsightsCredentialsPtrInput
	ClusterName                pulumi.StringPtrInput
	ClusterType                pulumi.StringInput
	ContainerRegistry          ContainerRegistryPropertiesPtrInput
	ContainerService           AcsClusterPropertiesInput
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

func init() {
	pulumi.RegisterOutputType(OperationalizationClusterOutput{})
}
