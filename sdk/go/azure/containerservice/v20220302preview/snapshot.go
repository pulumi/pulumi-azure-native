


package v20220302preview

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
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211001:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource Snapshot
	err := ctx.RegisterResource("azure-native:containerservice/v20220302preview:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure-native:containerservice/v20220302preview:Snapshot", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

func (o SnapshotOutput) CreationData() CreationDataResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) CreationDataResponsePtrOutput { return v.CreationData }).(CreationDataResponsePtrOutput)
}

func (o SnapshotOutput) EnableFIPS() pulumi.BoolOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolOutput { return v.EnableFIPS }).(pulumi.BoolOutput)
}

func (o SnapshotOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.KubernetesVersion }).(pulumi.StringOutput)
}

func (o SnapshotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SnapshotOutput) NodeImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.NodeImageVersion }).(pulumi.StringOutput)
}

func (o SnapshotOutput) OsSku() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.OsSku }).(pulumi.StringOutput)
}

func (o SnapshotOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

func (o SnapshotOutput) SnapshotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.SnapshotType }).(pulumi.StringPtrOutput)
}

func (o SnapshotOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Snapshot) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SnapshotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SnapshotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SnapshotOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.VmSize }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotOutput{})
}
