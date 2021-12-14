


package v20190501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Replication struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput    `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Status            StatusResponseOutput   `pulumi:"status"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewReplication(ctx *pulumi.Context,
	name string, args *ReplicationArgs, opts ...pulumi.ResourceOption) (*Replication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:Replication"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20170601preview:Replication"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20171001:Replication"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20191201preview:Replication"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:Replication"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:Replication"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:Replication"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210901:Replication"),
		},
	})
	opts = append(opts, aliases)
	var resource Replication
	err := ctx.RegisterResource("azure-native:containerregistry/v20190501:Replication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationState, opts ...pulumi.ResourceOption) (*Replication, error) {
	var resource Replication
	err := ctx.ReadResource("azure-native:containerregistry/v20190501:Replication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicationState struct {
}

type ReplicationState struct {
}

func (ReplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationState)(nil)).Elem()
}

type replicationArgs struct {
	Location          *string           `pulumi:"location"`
	RegistryName      string            `pulumi:"registryName"`
	ReplicationName   *string           `pulumi:"replicationName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type ReplicationArgs struct {
	Location          pulumi.StringPtrInput
	RegistryName      pulumi.StringInput
	ReplicationName   pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (ReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationArgs)(nil)).Elem()
}

type ReplicationInput interface {
	pulumi.Input

	ToReplicationOutput() ReplicationOutput
	ToReplicationOutputWithContext(ctx context.Context) ReplicationOutput
}

func (*Replication) ElementType() reflect.Type {
	return reflect.TypeOf((**Replication)(nil)).Elem()
}

func (i *Replication) ToReplicationOutput() ReplicationOutput {
	return i.ToReplicationOutputWithContext(context.Background())
}

func (i *Replication) ToReplicationOutputWithContext(ctx context.Context) ReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationOutput)
}

type ReplicationOutput struct{ *pulumi.OutputState }

func (ReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Replication)(nil)).Elem()
}

func (o ReplicationOutput) ToReplicationOutput() ReplicationOutput {
	return o
}

func (o ReplicationOutput) ToReplicationOutputWithContext(ctx context.Context) ReplicationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationOutput{})
}
