


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReplicationvCenter struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput          `pulumi:"location"`
	Name       pulumi.StringOutput             `pulumi:"name"`
	Properties VCenterPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput             `pulumi:"type"`
}


func NewReplicationvCenter(ctx *pulumi.Context,
	name string, args *ReplicationvCenterArgs, opts ...pulumi.ResourceOption) (*ReplicationvCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationvCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationvCenter
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210801:ReplicationvCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplicationvCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationvCenterState, opts ...pulumi.ResourceOption) (*ReplicationvCenter, error) {
	var resource ReplicationvCenter
	err := ctx.ReadResource("azure-native:recoveryservices/v20210801:ReplicationvCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicationvCenterState struct {
}

type ReplicationvCenterState struct {
}

func (ReplicationvCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationvCenterState)(nil)).Elem()
}

type replicationvCenterArgs struct {
	FabricName        string                       `pulumi:"fabricName"`
	Properties        *AddVCenterRequestProperties `pulumi:"properties"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	ResourceName      string                       `pulumi:"resourceName"`
	VcenterName       *string                      `pulumi:"vcenterName"`
}


type ReplicationvCenterArgs struct {
	FabricName        pulumi.StringInput
	Properties        AddVCenterRequestPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
	VcenterName       pulumi.StringPtrInput
}

func (ReplicationvCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationvCenterArgs)(nil)).Elem()
}

type ReplicationvCenterInput interface {
	pulumi.Input

	ToReplicationvCenterOutput() ReplicationvCenterOutput
	ToReplicationvCenterOutputWithContext(ctx context.Context) ReplicationvCenterOutput
}

func (*ReplicationvCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationvCenter)(nil)).Elem()
}

func (i *ReplicationvCenter) ToReplicationvCenterOutput() ReplicationvCenterOutput {
	return i.ToReplicationvCenterOutputWithContext(context.Background())
}

func (i *ReplicationvCenter) ToReplicationvCenterOutputWithContext(ctx context.Context) ReplicationvCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationvCenterOutput)
}

type ReplicationvCenterOutput struct{ *pulumi.OutputState }

func (ReplicationvCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationvCenter)(nil)).Elem()
}

func (o ReplicationvCenterOutput) ToReplicationvCenterOutput() ReplicationvCenterOutput {
	return o
}

func (o ReplicationvCenterOutput) ToReplicationvCenterOutputWithContext(ctx context.Context) ReplicationvCenterOutput {
	return o
}

func (o ReplicationvCenterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationvCenter) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ReplicationvCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationvCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReplicationvCenterOutput) Properties() VCenterPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationvCenter) VCenterPropertiesResponseOutput { return v.Properties }).(VCenterPropertiesResponseOutput)
}

func (o ReplicationvCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationvCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationvCenterOutput{})
}
