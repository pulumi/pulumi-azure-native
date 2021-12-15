


package v20160301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Namespace struct {
	pulumi.CustomResourceState

	CreatedAt          pulumi.StringPtrOutput `pulumi:"createdAt"`
	Critical           pulumi.BoolPtrOutput   `pulumi:"critical"`
	Enabled            pulumi.BoolPtrOutput   `pulumi:"enabled"`
	Location           pulumi.StringOutput    `pulumi:"location"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	NamespaceType      pulumi.StringPtrOutput `pulumi:"namespaceType"`
	ProvisioningState  pulumi.StringPtrOutput `pulumi:"provisioningState"`
	Region             pulumi.StringPtrOutput `pulumi:"region"`
	ScaleUnit          pulumi.StringPtrOutput `pulumi:"scaleUnit"`
	ServiceBusEndpoint pulumi.StringPtrOutput `pulumi:"serviceBusEndpoint"`
	Sku                SkuResponsePtrOutput   `pulumi:"sku"`
	Status             pulumi.StringPtrOutput `pulumi:"status"`
	SubscriptionId     pulumi.StringPtrOutput `pulumi:"subscriptionId"`
	Tags               pulumi.StringMapOutput `pulumi:"tags"`
	Type               pulumi.StringOutput    `pulumi:"type"`
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
			Type: pulumi.String("azure-native:notificationhubs:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20170401:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:notificationhubs/v20160301:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:notificationhubs/v20160301:Namespace", name, id, state, &resource, opts...)
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
	CreatedAt          *string           `pulumi:"createdAt"`
	Critical           *bool             `pulumi:"critical"`
	Enabled            *bool             `pulumi:"enabled"`
	Location           *string           `pulumi:"location"`
	Name               *string           `pulumi:"name"`
	NamespaceName      *string           `pulumi:"namespaceName"`
	NamespaceType      *NamespaceType    `pulumi:"namespaceType"`
	ProvisioningState  *string           `pulumi:"provisioningState"`
	Region             *string           `pulumi:"region"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	ScaleUnit          *string           `pulumi:"scaleUnit"`
	ServiceBusEndpoint *string           `pulumi:"serviceBusEndpoint"`
	Sku                *Sku              `pulumi:"sku"`
	Status             *string           `pulumi:"status"`
	SubscriptionId     *string           `pulumi:"subscriptionId"`
	Tags               map[string]string `pulumi:"tags"`
}


type NamespaceArgs struct {
	CreatedAt          pulumi.StringPtrInput
	Critical           pulumi.BoolPtrInput
	Enabled            pulumi.BoolPtrInput
	Location           pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	NamespaceName      pulumi.StringPtrInput
	NamespaceType      NamespaceTypePtrInput
	ProvisioningState  pulumi.StringPtrInput
	Region             pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	ScaleUnit          pulumi.StringPtrInput
	ServiceBusEndpoint pulumi.StringPtrInput
	Sku                SkuPtrInput
	Status             pulumi.StringPtrInput
	SubscriptionId     pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
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
