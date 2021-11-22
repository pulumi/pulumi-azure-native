


package servicebus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Namespace struct {
	pulumi.CustomResourceState

	CreatedAt          pulumi.StringOutput    `pulumi:"createdAt"`
	Location           pulumi.StringOutput    `pulumi:"location"`
	MetricId           pulumi.StringOutput    `pulumi:"metricId"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput    `pulumi:"provisioningState"`
	ServiceBusEndpoint pulumi.StringOutput    `pulumi:"serviceBusEndpoint"`
	Sku                SBSkuResponsePtrOutput `pulumi:"sku"`
	Tags               pulumi.StringMapOutput `pulumi:"tags"`
	Type               pulumi.StringOutput    `pulumi:"type"`
	UpdatedAt          pulumi.StringOutput    `pulumi:"updatedAt"`
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
			Type: pulumi.String("azure-native:servicebus/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20211101:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:servicebus:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:servicebus:Namespace", name, id, state, &resource, opts...)
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
	Location          *string           `pulumi:"location"`
	NamespaceName     *string           `pulumi:"namespaceName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *SBSku            `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type NamespaceArgs struct {
	Location          pulumi.StringPtrInput
	NamespaceName     pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SBSkuPtrInput
	Tags              pulumi.StringMapInput
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
