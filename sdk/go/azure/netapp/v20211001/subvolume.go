


package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Subvolume struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput      `pulumi:"name"`
	ParentPath        pulumi.StringPtrOutput   `pulumi:"parentPath"`
	Path              pulumi.StringPtrOutput   `pulumi:"path"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewSubvolume(ctx *pulumi.Context,
	name string, args *SubvolumeArgs, opts ...pulumi.ResourceOption) (*Subvolume, error) {
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
			Type: pulumi.String("azure-native:netapp:Subvolume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:Subvolume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:Subvolume"),
		},
	})
	opts = append(opts, aliases)
	var resource Subvolume
	err := ctx.RegisterResource("azure-native:netapp/v20211001:Subvolume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubvolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubvolumeState, opts ...pulumi.ResourceOption) (*Subvolume, error) {
	var resource Subvolume
	err := ctx.ReadResource("azure-native:netapp/v20211001:Subvolume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subvolumeState struct {
}

type SubvolumeState struct {
}

func (SubvolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*subvolumeState)(nil)).Elem()
}

type subvolumeArgs struct {
	AccountName       string   `pulumi:"accountName"`
	ParentPath        *string  `pulumi:"parentPath"`
	Path              *string  `pulumi:"path"`
	PoolName          string   `pulumi:"poolName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Size              *float64 `pulumi:"size"`
	SubvolumeName     *string  `pulumi:"subvolumeName"`
	VolumeName        string   `pulumi:"volumeName"`
}


type SubvolumeArgs struct {
	AccountName       pulumi.StringInput
	ParentPath        pulumi.StringPtrInput
	Path              pulumi.StringPtrInput
	PoolName          pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Size              pulumi.Float64PtrInput
	SubvolumeName     pulumi.StringPtrInput
	VolumeName        pulumi.StringInput
}

func (SubvolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subvolumeArgs)(nil)).Elem()
}

type SubvolumeInput interface {
	pulumi.Input

	ToSubvolumeOutput() SubvolumeOutput
	ToSubvolumeOutputWithContext(ctx context.Context) SubvolumeOutput
}

func (*Subvolume) ElementType() reflect.Type {
	return reflect.TypeOf((**Subvolume)(nil)).Elem()
}

func (i *Subvolume) ToSubvolumeOutput() SubvolumeOutput {
	return i.ToSubvolumeOutputWithContext(context.Background())
}

func (i *Subvolume) ToSubvolumeOutputWithContext(ctx context.Context) SubvolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubvolumeOutput)
}

type SubvolumeOutput struct{ *pulumi.OutputState }

func (SubvolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subvolume)(nil)).Elem()
}

func (o SubvolumeOutput) ToSubvolumeOutput() SubvolumeOutput {
	return o
}

func (o SubvolumeOutput) ToSubvolumeOutputWithContext(ctx context.Context) SubvolumeOutput {
	return o
}

func (o SubvolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Subvolume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SubvolumeOutput) ParentPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subvolume) pulumi.StringPtrOutput { return v.ParentPath }).(pulumi.StringPtrOutput)
}

func (o SubvolumeOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subvolume) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

func (o SubvolumeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Subvolume) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SubvolumeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Subvolume) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SubvolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Subvolume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubvolumeOutput{})
}
