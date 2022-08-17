


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IscsiTarget struct {
	pulumi.CustomResourceState

	AclMode           pulumi.StringOutput          `pulumi:"aclMode"`
	Endpoints         pulumi.StringArrayOutput     `pulumi:"endpoints"`
	Luns              IscsiLunResponseArrayOutput  `pulumi:"luns"`
	Name              pulumi.StringOutput          `pulumi:"name"`
	Port              pulumi.IntPtrOutput          `pulumi:"port"`
	ProvisioningState pulumi.StringOutput          `pulumi:"provisioningState"`
	StaticAcls        AclResponseArrayOutput       `pulumi:"staticAcls"`
	Status            pulumi.StringOutput          `pulumi:"status"`
	SystemData        SystemMetadataResponseOutput `pulumi:"systemData"`
	TargetIqn         pulumi.StringOutput          `pulumi:"targetIqn"`
	Type              pulumi.StringOutput          `pulumi:"type"`
}


func NewIscsiTarget(ctx *pulumi.Context,
	name string, args *IscsiTargetArgs, opts ...pulumi.ResourceOption) (*IscsiTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclMode == nil {
		return nil, errors.New("invalid value for required argument 'AclMode'")
	}
	if args.DiskPoolName == nil {
		return nil, errors.New("invalid value for required argument 'DiskPoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagepool:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagepool/v20200315preview:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagepool/v20210801:IscsiTarget"),
		},
	})
	opts = append(opts, aliases)
	var resource IscsiTarget
	err := ctx.RegisterResource("azure-native:storagepool/v20210401preview:IscsiTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIscsiTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IscsiTargetState, opts ...pulumi.ResourceOption) (*IscsiTarget, error) {
	var resource IscsiTarget
	err := ctx.ReadResource("azure-native:storagepool/v20210401preview:IscsiTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iscsiTargetState struct {
}

type IscsiTargetState struct {
}

func (IscsiTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiTargetState)(nil)).Elem()
}

type iscsiTargetArgs struct {
	AclMode           string     `pulumi:"aclMode"`
	DiskPoolName      string     `pulumi:"diskPoolName"`
	IscsiTargetName   *string    `pulumi:"iscsiTargetName"`
	Luns              []IscsiLun `pulumi:"luns"`
	ResourceGroupName string     `pulumi:"resourceGroupName"`
	StaticAcls        []Acl      `pulumi:"staticAcls"`
	TargetIqn         *string    `pulumi:"targetIqn"`
}


type IscsiTargetArgs struct {
	AclMode           pulumi.StringInput
	DiskPoolName      pulumi.StringInput
	IscsiTargetName   pulumi.StringPtrInput
	Luns              IscsiLunArrayInput
	ResourceGroupName pulumi.StringInput
	StaticAcls        AclArrayInput
	TargetIqn         pulumi.StringPtrInput
}

func (IscsiTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiTargetArgs)(nil)).Elem()
}

type IscsiTargetInput interface {
	pulumi.Input

	ToIscsiTargetOutput() IscsiTargetOutput
	ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput
}

func (*IscsiTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**IscsiTarget)(nil)).Elem()
}

func (i *IscsiTarget) ToIscsiTargetOutput() IscsiTargetOutput {
	return i.ToIscsiTargetOutputWithContext(context.Background())
}

func (i *IscsiTarget) ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiTargetOutput)
}

type IscsiTargetOutput struct{ *pulumi.OutputState }

func (IscsiTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IscsiTarget)(nil)).Elem()
}

func (o IscsiTargetOutput) ToIscsiTargetOutput() IscsiTargetOutput {
	return o
}

func (o IscsiTargetOutput) ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput {
	return o
}

func (o IscsiTargetOutput) AclMode() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.AclMode }).(pulumi.StringOutput)
}

func (o IscsiTargetOutput) Endpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringArrayOutput { return v.Endpoints }).(pulumi.StringArrayOutput)
}

func (o IscsiTargetOutput) Luns() IscsiLunResponseArrayOutput {
	return o.ApplyT(func(v *IscsiTarget) IscsiLunResponseArrayOutput { return v.Luns }).(IscsiLunResponseArrayOutput)
}

func (o IscsiTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IscsiTargetOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o IscsiTargetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o IscsiTargetOutput) StaticAcls() AclResponseArrayOutput {
	return o.ApplyT(func(v *IscsiTarget) AclResponseArrayOutput { return v.StaticAcls }).(AclResponseArrayOutput)
}

func (o IscsiTargetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o IscsiTargetOutput) SystemData() SystemMetadataResponseOutput {
	return o.ApplyT(func(v *IscsiTarget) SystemMetadataResponseOutput { return v.SystemData }).(SystemMetadataResponseOutput)
}

func (o IscsiTargetOutput) TargetIqn() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.TargetIqn }).(pulumi.StringOutput)
}

func (o IscsiTargetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IscsiTargetOutput{})
}
