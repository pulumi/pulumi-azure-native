


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Namespace struct {
	pulumi.CustomResourceState

	AlternateName              pulumi.StringPtrOutput                       `pulumi:"alternateName"`
	ClusterArmId               pulumi.StringPtrOutput                       `pulumi:"clusterArmId"`
	CreatedAt                  pulumi.StringOutput                          `pulumi:"createdAt"`
	DisableLocalAuth           pulumi.BoolPtrOutput                         `pulumi:"disableLocalAuth"`
	Encryption                 EncryptionResponsePtrOutput                  `pulumi:"encryption"`
	Identity                   IdentityResponsePtrOutput                    `pulumi:"identity"`
	IsAutoInflateEnabled       pulumi.BoolPtrOutput                         `pulumi:"isAutoInflateEnabled"`
	KafkaEnabled               pulumi.BoolPtrOutput                         `pulumi:"kafkaEnabled"`
	Location                   pulumi.StringPtrOutput                       `pulumi:"location"`
	MaximumThroughputUnits     pulumi.IntPtrOutput                          `pulumi:"maximumThroughputUnits"`
	MetricId                   pulumi.StringOutput                          `pulumi:"metricId"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	ServiceBusEndpoint         pulumi.StringOutput                          `pulumi:"serviceBusEndpoint"`
	Sku                        SkuResponsePtrOutput                         `pulumi:"sku"`
	Status                     pulumi.StringOutput                          `pulumi:"status"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
	UpdatedAt                  pulumi.StringOutput                          `pulumi:"updatedAt"`
	ZoneRedundant              pulumi.BoolPtrOutput                         `pulumi:"zoneRedundant"`
}


func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Encryption != nil {
		args.Encryption = args.Encryption.ToEncryptionPtrOutput().ApplyT(func(v *Encryption) *Encryption { return v.Defaults() }).(EncryptionPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20150801:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:eventhub/v20211101:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:eventhub/v20211101:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type namespaceState struct {
}

type NamespaceState struct {
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	AlternateName              *string                         `pulumi:"alternateName"`
	ClusterArmId               *string                         `pulumi:"clusterArmId"`
	DisableLocalAuth           *bool                           `pulumi:"disableLocalAuth"`
	Encryption                 *Encryption                     `pulumi:"encryption"`
	Identity                   *Identity                       `pulumi:"identity"`
	IsAutoInflateEnabled       *bool                           `pulumi:"isAutoInflateEnabled"`
	KafkaEnabled               *bool                           `pulumi:"kafkaEnabled"`
	Location                   *string                         `pulumi:"location"`
	MaximumThroughputUnits     *int                            `pulumi:"maximumThroughputUnits"`
	NamespaceName              *string                         `pulumi:"namespaceName"`
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	ResourceGroupName          string                          `pulumi:"resourceGroupName"`
	Sku                        *Sku                            `pulumi:"sku"`
	Tags                       map[string]string               `pulumi:"tags"`
	ZoneRedundant              *bool                           `pulumi:"zoneRedundant"`
}


type NamespaceArgs struct {
	AlternateName              pulumi.StringPtrInput
	ClusterArmId               pulumi.StringPtrInput
	DisableLocalAuth           pulumi.BoolPtrInput
	Encryption                 EncryptionPtrInput
	Identity                   IdentityPtrInput
	IsAutoInflateEnabled       pulumi.BoolPtrInput
	KafkaEnabled               pulumi.BoolPtrInput
	Location                   pulumi.StringPtrInput
	MaximumThroughputUnits     pulumi.IntPtrInput
	NamespaceName              pulumi.StringPtrInput
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	ResourceGroupName          pulumi.StringInput
	Sku                        SkuPtrInput
	Tags                       pulumi.StringMapInput
	ZoneRedundant              pulumi.BoolPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

func (o NamespaceOutput) AlternateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.AlternateName }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) ClusterArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.ClusterArmId }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o NamespaceOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o NamespaceOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) EncryptionResponsePtrOutput { return v.Encryption }).(EncryptionResponsePtrOutput)
}

func (o NamespaceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o NamespaceOutput) IsAutoInflateEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.IsAutoInflateEnabled }).(pulumi.BoolPtrOutput)
}

func (o NamespaceOutput) KafkaEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.KafkaEnabled }).(pulumi.BoolPtrOutput)
}

func (o NamespaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) MaximumThroughputUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.IntPtrOutput { return v.MaximumThroughputUnits }).(pulumi.IntPtrOutput)
}

func (o NamespaceOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.MetricId }).(pulumi.StringOutput)
}

func (o NamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NamespaceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Namespace) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o NamespaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NamespaceOutput) ServiceBusEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.ServiceBusEndpoint }).(pulumi.StringOutput)
}

func (o NamespaceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o NamespaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o NamespaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Namespace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o NamespaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NamespaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NamespaceOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o NamespaceOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}
