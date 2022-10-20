


package v20210601preview

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
	CreationDate                pulumi.StringOutput                        `pulumi:"creationDate"`
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
	SparkConfigProperties       SparkConfigPropertiesResponsePtrOutput     `pulumi:"sparkConfigProperties"`
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
			Type: pulumi.String("azure-native:synapse/v20210301:BigDataPool"),
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
	})
	opts = append(opts, aliases)
	var resource BigDataPool
	err := ctx.RegisterResource("azure-native:synapse/v20210601preview:BigDataPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBigDataPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigDataPoolState, opts ...pulumi.ResourceOption) (*BigDataPool, error) {
	var resource BigDataPool
	err := ctx.ReadResource("azure-native:synapse/v20210601preview:BigDataPool", name, id, state, &resource, opts...)
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
	SparkConfigProperties       *SparkConfigProperties     `pulumi:"sparkConfigProperties"`
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
	SparkConfigProperties       SparkConfigPropertiesPtrInput
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

func (o BigDataPoolOutput) AutoPause() AutoPausePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BigDataPool) AutoPausePropertiesResponsePtrOutput { return v.AutoPause }).(AutoPausePropertiesResponsePtrOutput)
}

func (o BigDataPoolOutput) AutoScale() AutoScalePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BigDataPool) AutoScalePropertiesResponsePtrOutput { return v.AutoScale }).(AutoScalePropertiesResponsePtrOutput)
}

func (o BigDataPoolOutput) CacheSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.IntPtrOutput { return v.CacheSize }).(pulumi.IntPtrOutput)
}

func (o BigDataPoolOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o BigDataPoolOutput) CustomLibraries() LibraryInfoResponseArrayOutput {
	return o.ApplyT(func(v *BigDataPool) LibraryInfoResponseArrayOutput { return v.CustomLibraries }).(LibraryInfoResponseArrayOutput)
}

func (o BigDataPoolOutput) DefaultSparkLogFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.DefaultSparkLogFolder }).(pulumi.StringPtrOutput)
}

func (o BigDataPoolOutput) DynamicExecutorAllocation() DynamicExecutorAllocationResponsePtrOutput {
	return o.ApplyT(func(v *BigDataPool) DynamicExecutorAllocationResponsePtrOutput { return v.DynamicExecutorAllocation }).(DynamicExecutorAllocationResponsePtrOutput)
}

func (o BigDataPoolOutput) IsComputeIsolationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.BoolPtrOutput { return v.IsComputeIsolationEnabled }).(pulumi.BoolPtrOutput)
}

func (o BigDataPoolOutput) LastSucceededTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringOutput { return v.LastSucceededTimestamp }).(pulumi.StringOutput)
}

func (o BigDataPoolOutput) LibraryRequirements() LibraryRequirementsResponsePtrOutput {
	return o.ApplyT(func(v *BigDataPool) LibraryRequirementsResponsePtrOutput { return v.LibraryRequirements }).(LibraryRequirementsResponsePtrOutput)
}

func (o BigDataPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BigDataPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BigDataPoolOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.IntPtrOutput { return v.NodeCount }).(pulumi.IntPtrOutput)
}

func (o BigDataPoolOutput) NodeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.NodeSize }).(pulumi.StringPtrOutput)
}

func (o BigDataPoolOutput) NodeSizeFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.NodeSizeFamily }).(pulumi.StringPtrOutput)
}

func (o BigDataPoolOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o BigDataPoolOutput) SessionLevelPackagesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.BoolPtrOutput { return v.SessionLevelPackagesEnabled }).(pulumi.BoolPtrOutput)
}

func (o BigDataPoolOutput) SparkConfigProperties() SparkConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BigDataPool) SparkConfigPropertiesResponsePtrOutput { return v.SparkConfigProperties }).(SparkConfigPropertiesResponsePtrOutput)
}

func (o BigDataPoolOutput) SparkEventsFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.SparkEventsFolder }).(pulumi.StringPtrOutput)
}

func (o BigDataPoolOutput) SparkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.SparkVersion }).(pulumi.StringPtrOutput)
}

func (o BigDataPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BigDataPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BigDataPoolOutput{})
}
