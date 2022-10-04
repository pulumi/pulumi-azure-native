


package v20220803preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedClusterSnapshot struct {
	pulumi.CustomResourceState

	CreationData                     CreationDataResponsePtrOutput                     `pulumi:"creationData"`
	Location                         pulumi.StringOutput                               `pulumi:"location"`
	ManagedClusterPropertiesReadOnly ManagedClusterPropertiesForSnapshotResponseOutput `pulumi:"managedClusterPropertiesReadOnly"`
	Name                             pulumi.StringOutput                               `pulumi:"name"`
	SnapshotType                     pulumi.StringPtrOutput                            `pulumi:"snapshotType"`
	SystemData                       SystemDataResponseOutput                          `pulumi:"systemData"`
	Tags                             pulumi.StringMapOutput                            `pulumi:"tags"`
	Type                             pulumi.StringOutput                               `pulumi:"type"`
}


func NewManagedClusterSnapshot(ctx *pulumi.Context,
	name string, args *ManagedClusterSnapshotArgs, opts ...pulumi.ResourceOption) (*ManagedClusterSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:ManagedClusterSnapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedClusterSnapshot
	err := ctx.RegisterResource("azure-native:containerservice/v20220803preview:ManagedClusterSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedClusterSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterSnapshotState, opts ...pulumi.ResourceOption) (*ManagedClusterSnapshot, error) {
	var resource ManagedClusterSnapshot
	err := ctx.ReadResource("azure-native:containerservice/v20220803preview:ManagedClusterSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedClusterSnapshotState struct {
}

type ManagedClusterSnapshotState struct {
}

func (ManagedClusterSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterSnapshotState)(nil)).Elem()
}

type managedClusterSnapshotArgs struct {
	CreationData      *CreationData     `pulumi:"creationData"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	SnapshotType      *string           `pulumi:"snapshotType"`
	Tags              map[string]string `pulumi:"tags"`
}


type ManagedClusterSnapshotArgs struct {
	CreationData      CreationDataPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	SnapshotType      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (ManagedClusterSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterSnapshotArgs)(nil)).Elem()
}

type ManagedClusterSnapshotInput interface {
	pulumi.Input

	ToManagedClusterSnapshotOutput() ManagedClusterSnapshotOutput
	ToManagedClusterSnapshotOutputWithContext(ctx context.Context) ManagedClusterSnapshotOutput
}

func (*ManagedClusterSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterSnapshot)(nil)).Elem()
}

func (i *ManagedClusterSnapshot) ToManagedClusterSnapshotOutput() ManagedClusterSnapshotOutput {
	return i.ToManagedClusterSnapshotOutputWithContext(context.Background())
}

func (i *ManagedClusterSnapshot) ToManagedClusterSnapshotOutputWithContext(ctx context.Context) ManagedClusterSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterSnapshotOutput)
}

type ManagedClusterSnapshotOutput struct{ *pulumi.OutputState }

func (ManagedClusterSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterSnapshot)(nil)).Elem()
}

func (o ManagedClusterSnapshotOutput) ToManagedClusterSnapshotOutput() ManagedClusterSnapshotOutput {
	return o
}

func (o ManagedClusterSnapshotOutput) ToManagedClusterSnapshotOutputWithContext(ctx context.Context) ManagedClusterSnapshotOutput {
	return o
}

func (o ManagedClusterSnapshotOutput) CreationData() CreationDataResponsePtrOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) CreationDataResponsePtrOutput { return v.CreationData }).(CreationDataResponsePtrOutput)
}

func (o ManagedClusterSnapshotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ManagedClusterSnapshotOutput) ManagedClusterPropertiesReadOnly() ManagedClusterPropertiesForSnapshotResponseOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) ManagedClusterPropertiesForSnapshotResponseOutput {
		return v.ManagedClusterPropertiesReadOnly
	}).(ManagedClusterPropertiesForSnapshotResponseOutput)
}

func (o ManagedClusterSnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterSnapshotOutput) SnapshotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) pulumi.StringPtrOutput { return v.SnapshotType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterSnapshotOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ManagedClusterSnapshotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedClusterSnapshotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterSnapshotOutput{})
}
