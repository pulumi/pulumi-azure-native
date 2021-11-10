


package v20210901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Snapshot struct {
	pulumi.CustomResourceState

	CreationData      CreationDataResponsePtrOutput `pulumi:"creationData"`
	EnableFIPS        pulumi.BoolOutput             `pulumi:"enableFIPS"`
	KubernetesVersion pulumi.StringOutput           `pulumi:"kubernetesVersion"`
	Location          pulumi.StringOutput           `pulumi:"location"`
	Name              pulumi.StringOutput           `pulumi:"name"`
	NodeImageVersion  pulumi.StringOutput           `pulumi:"nodeImageVersion"`
	OsSku             pulumi.StringOutput           `pulumi:"osSku"`
	OsType            pulumi.StringOutput           `pulumi:"osType"`
	SnapshotType      pulumi.StringPtrOutput        `pulumi:"snapshotType"`
	SystemData        SystemDataResponseOutput      `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput        `pulumi:"tags"`
	Type              pulumi.StringOutput           `pulumi:"type"`
	VmSize            pulumi.StringOutput           `pulumi:"vmSize"`
}


func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource Snapshot
	err := ctx.RegisterResource("azure-native:containerservice/v20210901:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure-native:containerservice/v20210901:Snapshot", name, id, state, &resource, opts...)
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
	CreationData      *CreationData     `pulumi:"creationData"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	SnapshotType      *string           `pulumi:"snapshotType"`
	Tags              map[string]string `pulumi:"tags"`
}


type SnapshotArgs struct {
	CreationData      CreationDataPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	SnapshotType      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
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
	pulumi.RegisterOutputType(SnapshotOutput{})
}
