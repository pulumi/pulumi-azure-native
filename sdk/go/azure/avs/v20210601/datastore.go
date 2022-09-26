


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Datastore struct {
	pulumi.CustomResourceState

	DiskPoolVolume    DiskPoolVolumeResponsePtrOutput `pulumi:"diskPoolVolume"`
	Name              pulumi.StringOutput             `pulumi:"name"`
	NetAppVolume      NetAppVolumeResponsePtrOutput   `pulumi:"netAppVolume"`
	ProvisioningState pulumi.StringOutput             `pulumi:"provisioningState"`
	Type              pulumi.StringOutput             `pulumi:"type"`
}


func NewDatastore(ctx *pulumi.Context,
	name string, args *DatastoreArgs, opts ...pulumi.ResourceOption) (*Datastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DiskPoolVolume != nil {
		args.DiskPoolVolume = args.DiskPoolVolume.ToDiskPoolVolumePtrOutput().ApplyT(func(v *DiskPoolVolume) *DiskPoolVolume { return v.Defaults() }).(DiskPoolVolumePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:Datastore"),
		},
	})
	opts = append(opts, aliases)
	var resource Datastore
	err := ctx.RegisterResource("azure-native:avs/v20210601:Datastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatastoreState, opts ...pulumi.ResourceOption) (*Datastore, error) {
	var resource Datastore
	err := ctx.ReadResource("azure-native:avs/v20210601:Datastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type datastoreState struct {
}

type DatastoreState struct {
}

func (DatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreState)(nil)).Elem()
}

type datastoreArgs struct {
	ClusterName       string          `pulumi:"clusterName"`
	DatastoreName     *string         `pulumi:"datastoreName"`
	DiskPoolVolume    *DiskPoolVolume `pulumi:"diskPoolVolume"`
	NetAppVolume      *NetAppVolume   `pulumi:"netAppVolume"`
	PrivateCloudName  string          `pulumi:"privateCloudName"`
	ResourceGroupName string          `pulumi:"resourceGroupName"`
}


type DatastoreArgs struct {
	ClusterName       pulumi.StringInput
	DatastoreName     pulumi.StringPtrInput
	DiskPoolVolume    DiskPoolVolumePtrInput
	NetAppVolume      NetAppVolumePtrInput
	PrivateCloudName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (DatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreArgs)(nil)).Elem()
}

type DatastoreInput interface {
	pulumi.Input

	ToDatastoreOutput() DatastoreOutput
	ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput
}

func (*Datastore) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (i *Datastore) ToDatastoreOutput() DatastoreOutput {
	return i.ToDatastoreOutputWithContext(context.Background())
}

func (i *Datastore) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreOutput)
}

type DatastoreOutput struct{ *pulumi.OutputState }

func (DatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (o DatastoreOutput) ToDatastoreOutput() DatastoreOutput {
	return o
}

func (o DatastoreOutput) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return o
}

func (o DatastoreOutput) DiskPoolVolume() DiskPoolVolumeResponsePtrOutput {
	return o.ApplyT(func(v *Datastore) DiskPoolVolumeResponsePtrOutput { return v.DiskPoolVolume }).(DiskPoolVolumeResponsePtrOutput)
}

func (o DatastoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatastoreOutput) NetAppVolume() NetAppVolumeResponsePtrOutput {
	return o.ApplyT(func(v *Datastore) NetAppVolumeResponsePtrOutput { return v.NetAppVolume }).(NetAppVolumeResponsePtrOutput)
}

func (o DatastoreOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DatastoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatastoreOutput{})
}
