


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FailoverGroup struct {
	pulumi.CustomResourceState

	Databases         pulumi.StringArrayOutput                       `pulumi:"databases"`
	Location          pulumi.StringOutput                            `pulumi:"location"`
	Name              pulumi.StringOutput                            `pulumi:"name"`
	PartnerServers    PartnerInfoResponseArrayOutput                 `pulumi:"partnerServers"`
	ReadOnlyEndpoint  FailoverGroupReadOnlyEndpointResponsePtrOutput `pulumi:"readOnlyEndpoint"`
	ReadWriteEndpoint FailoverGroupReadWriteEndpointResponseOutput   `pulumi:"readWriteEndpoint"`
	ReplicationRole   pulumi.StringOutput                            `pulumi:"replicationRole"`
	ReplicationState  pulumi.StringOutput                            `pulumi:"replicationState"`
	Tags              pulumi.StringMapOutput                         `pulumi:"tags"`
	Type              pulumi.StringOutput                            `pulumi:"type"`
}


func NewFailoverGroup(ctx *pulumi.Context,
	name string, args *FailoverGroupArgs, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PartnerServers == nil {
		return nil, errors.New("invalid value for required argument 'PartnerServers'")
	}
	if args.ReadWriteEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ReadWriteEndpoint'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:FailoverGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource FailoverGroup
	err := ctx.RegisterResource("azure-native:sql/v20220501preview:FailoverGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFailoverGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FailoverGroupState, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	var resource FailoverGroup
	err := ctx.ReadResource("azure-native:sql/v20220501preview:FailoverGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type failoverGroupState struct {
}

type FailoverGroupState struct {
}

func (FailoverGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupState)(nil)).Elem()
}

type failoverGroupArgs struct {
	Databases         []string                       `pulumi:"databases"`
	FailoverGroupName *string                        `pulumi:"failoverGroupName"`
	PartnerServers    []PartnerInfo                  `pulumi:"partnerServers"`
	ReadOnlyEndpoint  *FailoverGroupReadOnlyEndpoint `pulumi:"readOnlyEndpoint"`
	ReadWriteEndpoint FailoverGroupReadWriteEndpoint `pulumi:"readWriteEndpoint"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	ServerName        string                         `pulumi:"serverName"`
	Tags              map[string]string              `pulumi:"tags"`
}


type FailoverGroupArgs struct {
	Databases         pulumi.StringArrayInput
	FailoverGroupName pulumi.StringPtrInput
	PartnerServers    PartnerInfoArrayInput
	ReadOnlyEndpoint  FailoverGroupReadOnlyEndpointPtrInput
	ReadWriteEndpoint FailoverGroupReadWriteEndpointInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (FailoverGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupArgs)(nil)).Elem()
}

type FailoverGroupInput interface {
	pulumi.Input

	ToFailoverGroupOutput() FailoverGroupOutput
	ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput
}

func (*FailoverGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroup)(nil)).Elem()
}

func (i *FailoverGroup) ToFailoverGroupOutput() FailoverGroupOutput {
	return i.ToFailoverGroupOutputWithContext(context.Background())
}

func (i *FailoverGroup) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupOutput)
}

type FailoverGroupOutput struct{ *pulumi.OutputState }

func (FailoverGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroup)(nil)).Elem()
}

func (o FailoverGroupOutput) ToFailoverGroupOutput() FailoverGroupOutput {
	return o
}

func (o FailoverGroupOutput) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return o
}

func (o FailoverGroupOutput) Databases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringArrayOutput { return v.Databases }).(pulumi.StringArrayOutput)
}

func (o FailoverGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o FailoverGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FailoverGroupOutput) PartnerServers() PartnerInfoResponseArrayOutput {
	return o.ApplyT(func(v *FailoverGroup) PartnerInfoResponseArrayOutput { return v.PartnerServers }).(PartnerInfoResponseArrayOutput)
}

func (o FailoverGroupOutput) ReadOnlyEndpoint() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ApplyT(func(v *FailoverGroup) FailoverGroupReadOnlyEndpointResponsePtrOutput { return v.ReadOnlyEndpoint }).(FailoverGroupReadOnlyEndpointResponsePtrOutput)
}

func (o FailoverGroupOutput) ReadWriteEndpoint() FailoverGroupReadWriteEndpointResponseOutput {
	return o.ApplyT(func(v *FailoverGroup) FailoverGroupReadWriteEndpointResponseOutput { return v.ReadWriteEndpoint }).(FailoverGroupReadWriteEndpointResponseOutput)
}

func (o FailoverGroupOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.ReplicationRole }).(pulumi.StringOutput)
}

func (o FailoverGroupOutput) ReplicationState() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.ReplicationState }).(pulumi.StringOutput)
}

func (o FailoverGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FailoverGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FailoverGroupOutput{})
}
