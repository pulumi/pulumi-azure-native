


package v20210801

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
	ManagedBy         pulumi.StringOutput          `pulumi:"managedBy"`
	ManagedByExtended pulumi.StringArrayOutput     `pulumi:"managedByExtended"`
	Name              pulumi.StringOutput          `pulumi:"name"`
	Port              pulumi.IntPtrOutput          `pulumi:"port"`
	ProvisioningState pulumi.StringOutput          `pulumi:"provisioningState"`
	Sessions          pulumi.StringArrayOutput     `pulumi:"sessions"`
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
			Type: pulumi.String("azure-nextgen:storagepool/v20210801:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagepool:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagepool:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagepool/v20200315preview:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagepool/v20200315preview:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagepool/v20210401preview:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagepool/v20210401preview:IscsiTarget"),
		},
	})
	opts = append(opts, aliases)
	var resource IscsiTarget
	err := ctx.RegisterResource("azure-native:storagepool/v20210801:IscsiTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIscsiTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IscsiTargetState, opts ...pulumi.ResourceOption) (*IscsiTarget, error) {
	var resource IscsiTarget
	err := ctx.ReadResource("azure-native:storagepool/v20210801:IscsiTarget", name, id, state, &resource, opts...)
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
	ManagedBy         *string    `pulumi:"managedBy"`
	ManagedByExtended []string   `pulumi:"managedByExtended"`
	ResourceGroupName string     `pulumi:"resourceGroupName"`
	StaticAcls        []Acl      `pulumi:"staticAcls"`
	TargetIqn         *string    `pulumi:"targetIqn"`
}


type IscsiTargetArgs struct {
	AclMode           pulumi.StringInput
	DiskPoolName      pulumi.StringInput
	IscsiTargetName   pulumi.StringPtrInput
	Luns              IscsiLunArrayInput
	ManagedBy         pulumi.StringPtrInput
	ManagedByExtended pulumi.StringArrayInput
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
	return reflect.TypeOf((*IscsiTarget)(nil))
}

func (i *IscsiTarget) ToIscsiTargetOutput() IscsiTargetOutput {
	return i.ToIscsiTargetOutputWithContext(context.Background())
}

func (i *IscsiTarget) ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiTargetOutput)
}

type IscsiTargetOutput struct{ *pulumi.OutputState }

func (IscsiTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiTarget)(nil))
}

func (o IscsiTargetOutput) ToIscsiTargetOutput() IscsiTargetOutput {
	return o
}

func (o IscsiTargetOutput) ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IscsiTargetInput)(nil)).Elem(), &IscsiTarget{})
	pulumi.RegisterOutputType(IscsiTargetOutput{})
}
