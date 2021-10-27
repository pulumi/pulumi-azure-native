


package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Namespace struct {
	pulumi.CustomResourceState

	CreatedAt          pulumi.StringOutput         `pulumi:"createdAt"`
	Encryption         EncryptionResponsePtrOutput `pulumi:"encryption"`
	Identity           IdentityResponsePtrOutput   `pulumi:"identity"`
	Location           pulumi.StringOutput         `pulumi:"location"`
	MetricId           pulumi.StringOutput         `pulumi:"metricId"`
	Name               pulumi.StringOutput         `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput         `pulumi:"provisioningState"`
	ServiceBusEndpoint pulumi.StringOutput         `pulumi:"serviceBusEndpoint"`
	Sku                SBSkuResponsePtrOutput      `pulumi:"sku"`
	Status             pulumi.StringOutput         `pulumi:"status"`
	Tags               pulumi.StringMapOutput      `pulumi:"tags"`
	Type               pulumi.StringOutput         `pulumi:"type"`
	UpdatedAt          pulumi.StringOutput         `pulumi:"updatedAt"`
	ZoneRedundant      pulumi.BoolPtrOutput        `pulumi:"zoneRedundant"`
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
			Type: pulumi.String("azure-nextgen:servicebus/v20180101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20150801:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20210101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20210601preview:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:servicebus/v20180101preview:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:servicebus/v20180101preview:Namespace", name, id, state, &resource, opts...)
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
	Encryption        *Encryption       `pulumi:"encryption"`
	Identity          *Identity         `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	NamespaceName     *string           `pulumi:"namespaceName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *SBSku            `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	ZoneRedundant     *bool             `pulumi:"zoneRedundant"`
}


type NamespaceArgs struct {
	Encryption        EncryptionPtrInput
	Identity          IdentityPtrInput
	Location          pulumi.StringPtrInput
	NamespaceName     pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SBSkuPtrInput
	Tags              pulumi.StringMapInput
	ZoneRedundant     pulumi.BoolPtrInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceInput)(nil)).Elem(), &Namespace{})
	pulumi.RegisterOutputType(NamespaceOutput{})
}
