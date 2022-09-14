


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttachedDataNetwork struct {
	pulumi.CustomResourceState

	CreatedAt                            pulumi.StringPtrOutput             `pulumi:"createdAt"`
	CreatedBy                            pulumi.StringPtrOutput             `pulumi:"createdBy"`
	CreatedByType                        pulumi.StringPtrOutput             `pulumi:"createdByType"`
	DnsAddresses                         pulumi.StringArrayOutput           `pulumi:"dnsAddresses"`
	LastModifiedAt                       pulumi.StringPtrOutput             `pulumi:"lastModifiedAt"`
	LastModifiedBy                       pulumi.StringPtrOutput             `pulumi:"lastModifiedBy"`
	LastModifiedByType                   pulumi.StringPtrOutput             `pulumi:"lastModifiedByType"`
	Location                             pulumi.StringOutput                `pulumi:"location"`
	Name                                 pulumi.StringOutput                `pulumi:"name"`
	NaptConfiguration                    NaptConfigurationResponsePtrOutput `pulumi:"naptConfiguration"`
	ProvisioningState                    pulumi.StringOutput                `pulumi:"provisioningState"`
	SystemData                           SystemDataResponseOutput           `pulumi:"systemData"`
	Tags                                 pulumi.StringMapOutput             `pulumi:"tags"`
	Type                                 pulumi.StringOutput                `pulumi:"type"`
	UserEquipmentAddressPoolPrefix       pulumi.StringArrayOutput           `pulumi:"userEquipmentAddressPoolPrefix"`
	UserEquipmentStaticAddressPoolPrefix pulumi.StringArrayOutput           `pulumi:"userEquipmentStaticAddressPoolPrefix"`
	UserPlaneDataInterface               InterfacePropertiesResponseOutput  `pulumi:"userPlaneDataInterface"`
}


func NewAttachedDataNetwork(ctx *pulumi.Context,
	name string, args *AttachedDataNetworkArgs, opts ...pulumi.ResourceOption) (*AttachedDataNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PacketCoreControlPlaneName == nil {
		return nil, errors.New("invalid value for required argument 'PacketCoreControlPlaneName'")
	}
	if args.PacketCoreDataPlaneName == nil {
		return nil, errors.New("invalid value for required argument 'PacketCoreDataPlaneName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserPlaneDataInterface == nil {
		return nil, errors.New("invalid value for required argument 'UserPlaneDataInterface'")
	}
	if args.NaptConfiguration != nil {
		args.NaptConfiguration = args.NaptConfiguration.ToNaptConfigurationPtrOutput().ApplyT(func(v *NaptConfiguration) *NaptConfiguration { return v.Defaults() }).(NaptConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:AttachedDataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:AttachedDataNetwork"),
		},
	})
	opts = append(opts, aliases)
	var resource AttachedDataNetwork
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220401preview:AttachedDataNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAttachedDataNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedDataNetworkState, opts ...pulumi.ResourceOption) (*AttachedDataNetwork, error) {
	var resource AttachedDataNetwork
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220401preview:AttachedDataNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type attachedDataNetworkState struct {
}

type AttachedDataNetworkState struct {
}

func (AttachedDataNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDataNetworkState)(nil)).Elem()
}

type attachedDataNetworkArgs struct {
	AttachedDataNetworkName              *string             `pulumi:"attachedDataNetworkName"`
	CreatedAt                            *string             `pulumi:"createdAt"`
	CreatedBy                            *string             `pulumi:"createdBy"`
	CreatedByType                        *string             `pulumi:"createdByType"`
	DnsAddresses                         []string            `pulumi:"dnsAddresses"`
	LastModifiedAt                       *string             `pulumi:"lastModifiedAt"`
	LastModifiedBy                       *string             `pulumi:"lastModifiedBy"`
	LastModifiedByType                   *string             `pulumi:"lastModifiedByType"`
	Location                             *string             `pulumi:"location"`
	NaptConfiguration                    *NaptConfiguration  `pulumi:"naptConfiguration"`
	PacketCoreControlPlaneName           string              `pulumi:"packetCoreControlPlaneName"`
	PacketCoreDataPlaneName              string              `pulumi:"packetCoreDataPlaneName"`
	ResourceGroupName                    string              `pulumi:"resourceGroupName"`
	Tags                                 map[string]string   `pulumi:"tags"`
	UserEquipmentAddressPoolPrefix       []string            `pulumi:"userEquipmentAddressPoolPrefix"`
	UserEquipmentStaticAddressPoolPrefix []string            `pulumi:"userEquipmentStaticAddressPoolPrefix"`
	UserPlaneDataInterface               InterfaceProperties `pulumi:"userPlaneDataInterface"`
}


type AttachedDataNetworkArgs struct {
	AttachedDataNetworkName              pulumi.StringPtrInput
	CreatedAt                            pulumi.StringPtrInput
	CreatedBy                            pulumi.StringPtrInput
	CreatedByType                        pulumi.StringPtrInput
	DnsAddresses                         pulumi.StringArrayInput
	LastModifiedAt                       pulumi.StringPtrInput
	LastModifiedBy                       pulumi.StringPtrInput
	LastModifiedByType                   pulumi.StringPtrInput
	Location                             pulumi.StringPtrInput
	NaptConfiguration                    NaptConfigurationPtrInput
	PacketCoreControlPlaneName           pulumi.StringInput
	PacketCoreDataPlaneName              pulumi.StringInput
	ResourceGroupName                    pulumi.StringInput
	Tags                                 pulumi.StringMapInput
	UserEquipmentAddressPoolPrefix       pulumi.StringArrayInput
	UserEquipmentStaticAddressPoolPrefix pulumi.StringArrayInput
	UserPlaneDataInterface               InterfacePropertiesInput
}

func (AttachedDataNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDataNetworkArgs)(nil)).Elem()
}

type AttachedDataNetworkInput interface {
	pulumi.Input

	ToAttachedDataNetworkOutput() AttachedDataNetworkOutput
	ToAttachedDataNetworkOutputWithContext(ctx context.Context) AttachedDataNetworkOutput
}

func (*AttachedDataNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDataNetwork)(nil)).Elem()
}

func (i *AttachedDataNetwork) ToAttachedDataNetworkOutput() AttachedDataNetworkOutput {
	return i.ToAttachedDataNetworkOutputWithContext(context.Background())
}

func (i *AttachedDataNetwork) ToAttachedDataNetworkOutputWithContext(ctx context.Context) AttachedDataNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDataNetworkOutput)
}

type AttachedDataNetworkOutput struct{ *pulumi.OutputState }

func (AttachedDataNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDataNetwork)(nil)).Elem()
}

func (o AttachedDataNetworkOutput) ToAttachedDataNetworkOutput() AttachedDataNetworkOutput {
	return o
}

func (o AttachedDataNetworkOutput) ToAttachedDataNetworkOutputWithContext(ctx context.Context) AttachedDataNetworkOutput {
	return o
}

func (o AttachedDataNetworkOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o AttachedDataNetworkOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o AttachedDataNetworkOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o AttachedDataNetworkOutput) DnsAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringArrayOutput { return v.DnsAddresses }).(pulumi.StringArrayOutput)
}

func (o AttachedDataNetworkOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o AttachedDataNetworkOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o AttachedDataNetworkOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o AttachedDataNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AttachedDataNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AttachedDataNetworkOutput) NaptConfiguration() NaptConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) NaptConfigurationResponsePtrOutput { return v.NaptConfiguration }).(NaptConfigurationResponsePtrOutput)
}

func (o AttachedDataNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AttachedDataNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AttachedDataNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AttachedDataNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AttachedDataNetworkOutput) UserEquipmentAddressPoolPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringArrayOutput { return v.UserEquipmentAddressPoolPrefix }).(pulumi.StringArrayOutput)
}

func (o AttachedDataNetworkOutput) UserEquipmentStaticAddressPoolPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringArrayOutput { return v.UserEquipmentStaticAddressPoolPrefix }).(pulumi.StringArrayOutput)
}

func (o AttachedDataNetworkOutput) UserPlaneDataInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) InterfacePropertiesResponseOutput { return v.UserPlaneDataInterface }).(InterfacePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AttachedDataNetworkOutput{})
}
