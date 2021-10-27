


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	ClusterId         pulumi.IntOutput         `pulumi:"clusterId"`
	ClusterSize       pulumi.IntPtrOutput      `pulumi:"clusterSize"`
	Hosts             pulumi.StringArrayOutput `pulumi:"hosts"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Sku               SkuResponseOutput        `pulumi:"sku"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:avs/v20211201:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs:Cluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200320:Cluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20200320:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20200717preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:Cluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210601:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:avs/v20211201:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:avs/v20211201:Cluster", name, id, state, &resource, opts...)
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
	ClusterName       *string  `pulumi:"clusterName"`
	ClusterSize       *int     `pulumi:"clusterSize"`
	Hosts             []string `pulumi:"hosts"`
	PrivateCloudName  string   `pulumi:"privateCloudName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Sku               Sku      `pulumi:"sku"`
}


type ClusterArgs struct {
	ClusterName       pulumi.StringPtrInput
	ClusterSize       pulumi.IntPtrInput
	Hosts             pulumi.StringArrayInput
	PrivateCloudName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuInput
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
	return reflect.TypeOf((*Cluster)(nil))
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterOutputType(ClusterOutput{})
}
