


package v20200301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Snapshot struct {
	pulumi.CustomResourceState

	Created           pulumi.StringOutput `pulumi:"created"`
	Location          pulumi.StringOutput `pulumi:"location"`
	Name              pulumi.StringOutput `pulumi:"name"`
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	SnapshotId        pulumi.StringOutput `pulumi:"snapshotId"`
	Type              pulumi.StringOutput `pulumi:"type"`
}


func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VolumeName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20170815:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20170815:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190801:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190801:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191001:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20191001:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20191101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200801:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200901:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20201101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20201201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210401preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210601:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource Snapshot
	err := ctx.RegisterResource("azure-native:netapp/v20200301:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure-native:netapp/v20200301:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type snapshotState struct {
}

type SnapshotState struct {
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	AccountName       string  `pulumi:"accountName"`
	Location          *string `pulumi:"location"`
	PoolName          string  `pulumi:"poolName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SnapshotName      *string `pulumi:"snapshotName"`
	VolumeName        string  `pulumi:"volumeName"`
}


type SnapshotArgs struct {
	AccountName       pulumi.StringInput
	Location          pulumi.StringPtrInput
	PoolName          pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SnapshotName      pulumi.StringPtrInput
	VolumeName        pulumi.StringInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterOutputType(SnapshotOutput{})
}
