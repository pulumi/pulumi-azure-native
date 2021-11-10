


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Namespace struct {
	pulumi.CustomResourceState

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
			Type: pulumi.String("azure-native:eventhub/v20211101:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:eventhub/v20210601preview:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:eventhub/v20210601preview:Namespace", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*Namespace)(nil))
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil))
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}
