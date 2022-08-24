


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceFailoverGroup struct {
	pulumi.CustomResourceState

	ManagedInstancePairs ManagedInstancePairInfoResponseArrayOutput             `pulumi:"managedInstancePairs"`
	Name                 pulumi.StringOutput                                    `pulumi:"name"`
	PartnerRegions       PartnerRegionInfoResponseArrayOutput                   `pulumi:"partnerRegions"`
	ReadOnlyEndpoint     InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput `pulumi:"readOnlyEndpoint"`
	ReadWriteEndpoint    InstanceFailoverGroupReadWriteEndpointResponseOutput   `pulumi:"readWriteEndpoint"`
	ReplicationRole      pulumi.StringOutput                                    `pulumi:"replicationRole"`
	ReplicationState     pulumi.StringOutput                                    `pulumi:"replicationState"`
	Type                 pulumi.StringOutput                                    `pulumi:"type"`
}


func NewInstanceFailoverGroup(ctx *pulumi.Context,
	name string, args *InstanceFailoverGroupArgs, opts ...pulumi.ResourceOption) (*InstanceFailoverGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocationName == nil {
		return nil, errors.New("invalid value for required argument 'LocationName'")
	}
	if args.ManagedInstancePairs == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstancePairs'")
	}
	if args.PartnerRegions == nil {
		return nil, errors.New("invalid value for required argument 'PartnerRegions'")
	}
	if args.ReadWriteEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ReadWriteEndpoint'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:InstanceFailoverGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource InstanceFailoverGroup
	err := ctx.RegisterResource("azure-native:sql/v20220201preview:InstanceFailoverGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetInstanceFailoverGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceFailoverGroupState, opts ...pulumi.ResourceOption) (*InstanceFailoverGroup, error) {
	var resource InstanceFailoverGroup
	err := ctx.ReadResource("azure-native:sql/v20220201preview:InstanceFailoverGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type instanceFailoverGroupState struct {
}

type InstanceFailoverGroupState struct {
}

func (InstanceFailoverGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFailoverGroupState)(nil)).Elem()
}

type instanceFailoverGroupArgs struct {
	FailoverGroupName    *string                                `pulumi:"failoverGroupName"`
	LocationName         string                                 `pulumi:"locationName"`
	ManagedInstancePairs []ManagedInstancePairInfo              `pulumi:"managedInstancePairs"`
	PartnerRegions       []PartnerRegionInfo                    `pulumi:"partnerRegions"`
	ReadOnlyEndpoint     *InstanceFailoverGroupReadOnlyEndpoint `pulumi:"readOnlyEndpoint"`
	ReadWriteEndpoint    InstanceFailoverGroupReadWriteEndpoint `pulumi:"readWriteEndpoint"`
	ResourceGroupName    string                                 `pulumi:"resourceGroupName"`
}


type InstanceFailoverGroupArgs struct {
	FailoverGroupName    pulumi.StringPtrInput
	LocationName         pulumi.StringInput
	ManagedInstancePairs ManagedInstancePairInfoArrayInput
	PartnerRegions       PartnerRegionInfoArrayInput
	ReadOnlyEndpoint     InstanceFailoverGroupReadOnlyEndpointPtrInput
	ReadWriteEndpoint    InstanceFailoverGroupReadWriteEndpointInput
	ResourceGroupName    pulumi.StringInput
}

func (InstanceFailoverGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFailoverGroupArgs)(nil)).Elem()
}

type InstanceFailoverGroupInput interface {
	pulumi.Input

	ToInstanceFailoverGroupOutput() InstanceFailoverGroupOutput
	ToInstanceFailoverGroupOutputWithContext(ctx context.Context) InstanceFailoverGroupOutput
}

func (*InstanceFailoverGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroup)(nil)).Elem()
}

func (i *InstanceFailoverGroup) ToInstanceFailoverGroupOutput() InstanceFailoverGroupOutput {
	return i.ToInstanceFailoverGroupOutputWithContext(context.Background())
}

func (i *InstanceFailoverGroup) ToInstanceFailoverGroupOutputWithContext(ctx context.Context) InstanceFailoverGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupOutput)
}

type InstanceFailoverGroupOutput struct{ *pulumi.OutputState }

func (InstanceFailoverGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroup)(nil)).Elem()
}

func (o InstanceFailoverGroupOutput) ToInstanceFailoverGroupOutput() InstanceFailoverGroupOutput {
	return o
}

func (o InstanceFailoverGroupOutput) ToInstanceFailoverGroupOutputWithContext(ctx context.Context) InstanceFailoverGroupOutput {
	return o
}

func (o InstanceFailoverGroupOutput) ManagedInstancePairs() ManagedInstancePairInfoResponseArrayOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) ManagedInstancePairInfoResponseArrayOutput {
		return v.ManagedInstancePairs
	}).(ManagedInstancePairInfoResponseArrayOutput)
}

func (o InstanceFailoverGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o InstanceFailoverGroupOutput) PartnerRegions() PartnerRegionInfoResponseArrayOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) PartnerRegionInfoResponseArrayOutput { return v.PartnerRegions }).(PartnerRegionInfoResponseArrayOutput)
}

func (o InstanceFailoverGroupOutput) ReadOnlyEndpoint() InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
		return v.ReadOnlyEndpoint
	}).(InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput)
}

func (o InstanceFailoverGroupOutput) ReadWriteEndpoint() InstanceFailoverGroupReadWriteEndpointResponseOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) InstanceFailoverGroupReadWriteEndpointResponseOutput {
		return v.ReadWriteEndpoint
	}).(InstanceFailoverGroupReadWriteEndpointResponseOutput)
}

func (o InstanceFailoverGroupOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) pulumi.StringOutput { return v.ReplicationRole }).(pulumi.StringOutput)
}

func (o InstanceFailoverGroupOutput) ReplicationState() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) pulumi.StringOutput { return v.ReplicationState }).(pulumi.StringOutput)
}

func (o InstanceFailoverGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceFailoverGroupOutput{})
}
