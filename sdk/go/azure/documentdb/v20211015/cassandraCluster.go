


package v20211015

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CassandraCluster struct {
	pulumi.CustomResourceState

	Identity   ManagedCassandraManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput                                  `pulumi:"location"`
	Name       pulumi.StringOutput                                     `pulumi:"name"`
	Properties ClusterResourceResponsePropertiesOutput                 `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type       pulumi.StringOutput                                     `pulumi:"type"`
}


func NewCassandraCluster(ctx *pulumi.Context,
	name string, args *CassandraClusterArgs, opts ...pulumi.ResourceOption) (*CassandraCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:CassandraCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource CassandraCluster
	err := ctx.RegisterResource("azure-native:documentdb/v20211015:CassandraCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCassandraCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraClusterState, opts ...pulumi.ResourceOption) (*CassandraCluster, error) {
	var resource CassandraCluster
	err := ctx.ReadResource("azure-native:documentdb/v20211015:CassandraCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cassandraClusterState struct {
}

type CassandraClusterState struct {
}

func (CassandraClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraClusterState)(nil)).Elem()
}

type cassandraClusterArgs struct {
	ClusterName       *string                                 `pulumi:"clusterName"`
	Identity          *ManagedCassandraManagedServiceIdentity `pulumi:"identity"`
	Location          *string                                 `pulumi:"location"`
	Properties        *ClusterResourceProperties              `pulumi:"properties"`
	ResourceGroupName string                                  `pulumi:"resourceGroupName"`
	Tags              map[string]string                       `pulumi:"tags"`
}


type CassandraClusterArgs struct {
	ClusterName       pulumi.StringPtrInput
	Identity          ManagedCassandraManagedServiceIdentityPtrInput
	Location          pulumi.StringPtrInput
	Properties        ClusterResourcePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (CassandraClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraClusterArgs)(nil)).Elem()
}

type CassandraClusterInput interface {
	pulumi.Input

	ToCassandraClusterOutput() CassandraClusterOutput
	ToCassandraClusterOutputWithContext(ctx context.Context) CassandraClusterOutput
}

func (*CassandraCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraCluster)(nil)).Elem()
}

func (i *CassandraCluster) ToCassandraClusterOutput() CassandraClusterOutput {
	return i.ToCassandraClusterOutputWithContext(context.Background())
}

func (i *CassandraCluster) ToCassandraClusterOutputWithContext(ctx context.Context) CassandraClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraClusterOutput)
}

type CassandraClusterOutput struct{ *pulumi.OutputState }

func (CassandraClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraCluster)(nil)).Elem()
}

func (o CassandraClusterOutput) ToCassandraClusterOutput() CassandraClusterOutput {
	return o
}

func (o CassandraClusterOutput) ToCassandraClusterOutputWithContext(ctx context.Context) CassandraClusterOutput {
	return o
}

func (o CassandraClusterOutput) Identity() ManagedCassandraManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *CassandraCluster) ManagedCassandraManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedCassandraManagedServiceIdentityResponsePtrOutput)
}

func (o CassandraClusterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraCluster) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o CassandraClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CassandraClusterOutput) Properties() ClusterResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *CassandraCluster) ClusterResourceResponsePropertiesOutput { return v.Properties }).(ClusterResourceResponsePropertiesOutput)
}

func (o CassandraClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CassandraCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CassandraClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CassandraClusterOutput{})
}
