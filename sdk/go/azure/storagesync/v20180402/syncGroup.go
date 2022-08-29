


package v20180402

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SyncGroup struct {
	pulumi.CustomResourceState

	Name            pulumi.StringOutput    `pulumi:"name"`
	SyncGroupStatus pulumi.StringOutput    `pulumi:"syncGroupStatus"`
	Type            pulumi.StringOutput    `pulumi:"type"`
	UniqueId        pulumi.StringPtrOutput `pulumi:"uniqueId"`
}


func NewSyncGroup(ctx *pulumi.Context,
	name string, args *SyncGroupArgs, opts ...pulumi.ResourceOption) (*SyncGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageSyncServiceName == nil {
		return nil, errors.New("invalid value for required argument 'StorageSyncServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagesync:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20170605preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180701:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20181001:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190201:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190301:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190601:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20191001:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200301:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200901:SyncGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource SyncGroup
	err := ctx.RegisterResource("azure-native:storagesync/v20180402:SyncGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSyncGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncGroupState, opts ...pulumi.ResourceOption) (*SyncGroup, error) {
	var resource SyncGroup
	err := ctx.ReadResource("azure-native:storagesync/v20180402:SyncGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type syncGroupState struct {
}

type SyncGroupState struct {
}

func (SyncGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGroupState)(nil)).Elem()
}

type syncGroupArgs struct {
	Location               *string           `pulumi:"location"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	StorageSyncServiceName string            `pulumi:"storageSyncServiceName"`
	SyncGroupName          *string           `pulumi:"syncGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
}


type SyncGroupArgs struct {
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	StorageSyncServiceName pulumi.StringInput
	SyncGroupName          pulumi.StringPtrInput
	Tags                   pulumi.StringMapInput
}

func (SyncGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGroupArgs)(nil)).Elem()
}

type SyncGroupInput interface {
	pulumi.Input

	ToSyncGroupOutput() SyncGroupOutput
	ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput
}

func (*SyncGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroup)(nil)).Elem()
}

func (i *SyncGroup) ToSyncGroupOutput() SyncGroupOutput {
	return i.ToSyncGroupOutputWithContext(context.Background())
}

func (i *SyncGroup) ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupOutput)
}

type SyncGroupOutput struct{ *pulumi.OutputState }

func (SyncGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroup)(nil)).Elem()
}

func (o SyncGroupOutput) ToSyncGroupOutput() SyncGroupOutput {
	return o
}

func (o SyncGroupOutput) ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput {
	return o
}

func (o SyncGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SyncGroupOutput) SyncGroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringOutput { return v.SyncGroupStatus }).(pulumi.StringOutput)
}

func (o SyncGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SyncGroupOutput) UniqueId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringPtrOutput { return v.UniqueId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SyncGroupOutput{})
}
