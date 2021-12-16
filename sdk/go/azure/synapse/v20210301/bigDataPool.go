


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BigDataPool struct {
	pulumi.CustomResourceState

	AutoPause                   AutoPausePropertiesResponsePtrOutput       `pulumi:"autoPause"`
	AutoScale                   AutoScalePropertiesResponsePtrOutput       `pulumi:"autoScale"`
	CacheSize                   pulumi.IntPtrOutput                        `pulumi:"cacheSize"`
	CreationDate                pulumi.StringPtrOutput                     `pulumi:"creationDate"`
	CustomLibraries             LibraryInfoResponseArrayOutput             `pulumi:"customLibraries"`
	DefaultSparkLogFolder       pulumi.StringPtrOutput                     `pulumi:"defaultSparkLogFolder"`
	DynamicExecutorAllocation   DynamicExecutorAllocationResponsePtrOutput `pulumi:"dynamicExecutorAllocation"`
	IsComputeIsolationEnabled   pulumi.BoolPtrOutput                       `pulumi:"isComputeIsolationEnabled"`
	LastSucceededTimestamp      pulumi.StringOutput                        `pulumi:"lastSucceededTimestamp"`
	LibraryRequirements         LibraryRequirementsResponsePtrOutput       `pulumi:"libraryRequirements"`
	Location                    pulumi.StringOutput                        `pulumi:"location"`
	Name                        pulumi.StringOutput                        `pulumi:"name"`
	NodeCount                   pulumi.IntPtrOutput                        `pulumi:"nodeCount"`
	NodeSize                    pulumi.StringPtrOutput                     `pulumi:"nodeSize"`
	NodeSizeFamily              pulumi.StringPtrOutput                     `pulumi:"nodeSizeFamily"`
	ProvisioningState           pulumi.StringPtrOutput                     `pulumi:"provisioningState"`
	SessionLevelPackagesEnabled pulumi.BoolPtrOutput                       `pulumi:"sessionLevelPackagesEnabled"`
	SparkConfigProperties       LibraryRequirementsResponsePtrOutput       `pulumi:"sparkConfigProperties"`
	SparkEventsFolder           pulumi.StringPtrOutput                     `pulumi:"sparkEventsFolder"`
	SparkVersion                pulumi.StringPtrOutput                     `pulumi:"sparkVersion"`
	Tags                        pulumi.StringMapOutput                     `pulumi:"tags"`
	Type                        pulumi.StringOutput                        `pulumi:"type"`
}


func NewBigDataPool(ctx *pulumi.Context,
	name string, args *BigDataPoolArgs, opts ...pulumi.ResourceOption) (*BigDataPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:BigDataPool"),
		},
	})
	opts = append(opts, aliases)
	var resource BigDataPool
	err := ctx.RegisterResource("azure-native:synapse/v20210301:BigDataPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBigDataPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigDataPoolState, opts ...pulumi.ResourceOption) (*BigDataPool, error) {
	var resource BigDataPool
	err := ctx.ReadResource("azure-native:synapse/v20210301:BigDataPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type bigDataPoolState struct {
}

type BigDataPoolState struct {
}

func (BigDataPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigDataPoolState)(nil)).Elem()
}

type bigDataPoolArgs struct {
	AutoPause                   *AutoPauseProperties       `pulumi:"autoPause"`
	AutoScale                   *AutoScaleProperties       `pulumi:"autoScale"`
	BigDataPoolName             *string                    `pulumi:"bigDataPoolName"`
	CacheSize                   *int                       `pulumi:"cacheSize"`
	CreationDate                *string                    `pulumi:"creationDate"`
	CustomLibraries             []LibraryInfo              `pulumi:"customLibraries"`
	DefaultSparkLogFolder       *string                    `pulumi:"defaultSparkLogFolder"`
	DynamicExecutorAllocation   *DynamicExecutorAllocation `pulumi:"dynamicExecutorAllocation"`
	Force                       *bool                      `pulumi:"force"`
	IsComputeIsolationEnabled   *bool                      `pulumi:"isComputeIsolationEnabled"`
	LibraryRequirements         *LibraryRequirements       `pulumi:"libraryRequirements"`
	Location                    *string                    `pulumi:"location"`
	NodeCount                   *int                       `pulumi:"nodeCount"`
	NodeSize                    *string                    `pulumi:"nodeSize"`
	NodeSizeFamily              *string                    `pulumi:"nodeSizeFamily"`
	ProvisioningState           *string                    `pulumi:"provisioningState"`
	ResourceGroupName           string                     `pulumi:"resourceGroupName"`
	SessionLevelPackagesEnabled *bool                      `pulumi:"sessionLevelPackagesEnabled"`
	SparkConfigProperties       *LibraryRequirements       `pulumi:"sparkConfigProperties"`
	SparkEventsFolder           *string                    `pulumi:"sparkEventsFolder"`
	SparkVersion                *string                    `pulumi:"sparkVersion"`
	Tags                        map[string]string          `pulumi:"tags"`
	WorkspaceName               string                     `pulumi:"workspaceName"`
}


type BigDataPoolArgs struct {
	AutoPause                   AutoPausePropertiesPtrInput
	AutoScale                   AutoScalePropertiesPtrInput
	BigDataPoolName             pulumi.StringPtrInput
	CacheSize                   pulumi.IntPtrInput
	CreationDate                pulumi.StringPtrInput
	CustomLibraries             LibraryInfoArrayInput
	DefaultSparkLogFolder       pulumi.StringPtrInput
	DynamicExecutorAllocation   DynamicExecutorAllocationPtrInput
	Force                       pulumi.BoolPtrInput
	IsComputeIsolationEnabled   pulumi.BoolPtrInput
	LibraryRequirements         LibraryRequirementsPtrInput
	Location                    pulumi.StringPtrInput
	NodeCount                   pulumi.IntPtrInput
	NodeSize                    pulumi.StringPtrInput
	NodeSizeFamily              pulumi.StringPtrInput
	ProvisioningState           pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	SessionLevelPackagesEnabled pulumi.BoolPtrInput
	SparkConfigProperties       LibraryRequirementsPtrInput
	SparkEventsFolder           pulumi.StringPtrInput
	SparkVersion                pulumi.StringPtrInput
	Tags                        pulumi.StringMapInput
	WorkspaceName               pulumi.StringInput
}

func (BigDataPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigDataPoolArgs)(nil)).Elem()
}

type BigDataPoolInput interface {
	pulumi.Input

	ToBigDataPoolOutput() BigDataPoolOutput
	ToBigDataPoolOutputWithContext(ctx context.Context) BigDataPoolOutput
}

func (*BigDataPool) ElementType() reflect.Type {
	return reflect.TypeOf((**BigDataPool)(nil)).Elem()
}

func (i *BigDataPool) ToBigDataPoolOutput() BigDataPoolOutput {
	return i.ToBigDataPoolOutputWithContext(context.Background())
}

func (i *BigDataPool) ToBigDataPoolOutputWithContext(ctx context.Context) BigDataPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigDataPoolOutput)
}

type BigDataPoolOutput struct{ *pulumi.OutputState }

func (BigDataPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigDataPool)(nil)).Elem()
}

func (o BigDataPoolOutput) ToBigDataPoolOutput() BigDataPoolOutput {
	return o
}

func (o BigDataPoolOutput) ToBigDataPoolOutputWithContext(ctx context.Context) BigDataPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BigDataPoolOutput{})
}
