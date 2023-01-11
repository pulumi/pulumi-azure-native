


package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	CreatedAt pulumi.StringOutput         `pulumi:"createdAt"`
	Location  pulumi.StringPtrOutput      `pulumi:"location"`
	MetricId  pulumi.StringOutput         `pulumi:"metricId"`
	Name      pulumi.StringOutput         `pulumi:"name"`
	Sku       ClusterSkuResponsePtrOutput `pulumi:"sku"`
	Status    pulumi.StringOutput         `pulumi:"status"`
	Tags      pulumi.StringMapOutput      `pulumi:"tags"`
	Type      pulumi.StringOutput         `pulumi:"type"`
	UpdatedAt pulumi.StringOutput         `pulumi:"updatedAt"`
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
			Type: pulumi.String("azure-native:eventhub/v20180101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20221001preview:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:eventhub:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:eventhub:Cluster", name, id, state, &resource, opts...)
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
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *ClusterSku       `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type ClusterArgs struct {
	ClusterName       pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               ClusterSkuPtrInput
	Tags              pulumi.StringMapInput
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

func (o ClusterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ClusterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.MetricId }).(pulumi.StringOutput)
}

func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterOutput) Sku() ClusterSkuResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterSkuResponsePtrOutput { return v.Sku }).(ClusterSkuResponsePtrOutput)
}

func (o ClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ClusterOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
