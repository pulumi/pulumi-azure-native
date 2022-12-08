


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	CustomResourceName pulumi.StringOutput               `pulumi:"customResourceName"`
	DatastoreIds       pulumi.StringArrayOutput          `pulumi:"datastoreIds"`
	ExtendedLocation   ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	InventoryItemId    pulumi.StringPtrOutput            `pulumi:"inventoryItemId"`
	Kind               pulumi.StringPtrOutput            `pulumi:"kind"`
	Location           pulumi.StringOutput               `pulumi:"location"`
	MoName             pulumi.StringOutput               `pulumi:"moName"`
	MoRefId            pulumi.StringPtrOutput            `pulumi:"moRefId"`
	Name               pulumi.StringOutput               `pulumi:"name"`
	NetworkIds         pulumi.StringArrayOutput          `pulumi:"networkIds"`
	ProvisioningState  pulumi.StringOutput               `pulumi:"provisioningState"`
	Statuses           ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	SystemData         SystemDataResponseOutput          `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput            `pulumi:"tags"`
	Type               pulumi.StringOutput               `pulumi:"type"`
	Uuid               pulumi.StringOutput               `pulumi:"uuid"`
	VCenterId          pulumi.StringPtrOutput            `pulumi:"vCenterId"`
}


func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20201001preview:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20201001preview:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	ClusterName       *string           `pulumi:"clusterName"`
	ExtendedLocation  *ExtendedLocation `pulumi:"extendedLocation"`
	InventoryItemId   *string           `pulumi:"inventoryItemId"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	MoRefId           *string           `pulumi:"moRefId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VCenterId         *string           `pulumi:"vCenterId"`
}


type ClusterArgs struct {
	ClusterName       pulumi.StringPtrInput
	ExtendedLocation  ExtendedLocationPtrInput
	InventoryItemId   pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	MoRefId           pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VCenterId         pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o ClusterOutput) DatastoreIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.DatastoreIds }).(pulumi.StringArrayOutput)
}

func (o ClusterOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o ClusterOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ClusterOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.MoName }).(pulumi.StringOutput)
}

func (o ClusterOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterOutput) NetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.NetworkIds }).(pulumi.StringArrayOutput)
}

func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ClusterOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o ClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Cluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ClusterOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o ClusterOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
