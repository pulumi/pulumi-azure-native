


package hdinsight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput             `pulumi:"etag"`
	Identity   ClusterIdentityResponsePtrOutput   `pulumi:"identity"`
	Location   pulumi.StringPtrOutput             `pulumi:"location"`
	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties ClusterGetPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput             `pulumi:"tags"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToClusterCreatePropertiesPtrOutput().ApplyT(func(v *ClusterCreateProperties) *ClusterCreateProperties { return v.Defaults() }).(ClusterCreatePropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hdinsight/v20150301preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20180601preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20210601:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:hdinsight:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:hdinsight:Cluster", name, id, state, &resource, opts...)
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
	ClusterName       *string                  `pulumi:"clusterName"`
	Identity          *ClusterIdentity         `pulumi:"identity"`
	Location          *string                  `pulumi:"location"`
	Properties        *ClusterCreateProperties `pulumi:"properties"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	Tags              map[string]string        `pulumi:"tags"`
}


type ClusterArgs struct {
	ClusterName       pulumi.StringPtrInput
	Identity          ClusterIdentityPtrInput
	Location          pulumi.StringPtrInput
	Properties        ClusterCreatePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
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

func (o ClusterOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) Identity() ClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterIdentityResponsePtrOutput { return v.Identity }).(ClusterIdentityResponsePtrOutput)
}

func (o ClusterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterOutput) Properties() ClusterGetPropertiesResponseOutput {
	return o.ApplyT(func(v *Cluster) ClusterGetPropertiesResponseOutput { return v.Properties }).(ClusterGetPropertiesResponseOutput)
}

func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
