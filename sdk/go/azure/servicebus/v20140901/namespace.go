


package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Namespace struct {
	pulumi.CustomResourceState

	CreateACSNamespace pulumi.BoolPtrOutput   `pulumi:"createACSNamespace"`
	CreatedAt          pulumi.StringOutput    `pulumi:"createdAt"`
	Enabled            pulumi.BoolPtrOutput   `pulumi:"enabled"`
	Location           pulumi.StringOutput    `pulumi:"location"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput    `pulumi:"provisioningState"`
	ServiceBusEndpoint pulumi.StringOutput    `pulumi:"serviceBusEndpoint"`
	Sku                SkuResponsePtrOutput   `pulumi:"sku"`
	Status             pulumi.StringPtrOutput `pulumi:"status"`
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
			Type: pulumi.String("azure-native:servicebus:Namespace"),
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
		{
			Type: pulumi.String("azure-native:servicebus/v20220101preview:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:servicebus/v20140901:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:servicebus/v20140901:Namespace", name, id, state, &resource, opts...)
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
	CreateACSNamespace *bool               `pulumi:"createACSNamespace"`
	Enabled            *bool               `pulumi:"enabled"`
	Location           *string             `pulumi:"location"`
	NamespaceName      *string             `pulumi:"namespaceName"`
	ResourceGroupName  string              `pulumi:"resourceGroupName"`
	Sku                *Sku                `pulumi:"sku"`
	Status             *NamespaceStateEnum `pulumi:"status"`
	Tags               map[string]string   `pulumi:"tags"`
}


type NamespaceArgs struct {
	CreateACSNamespace pulumi.BoolPtrInput
	Enabled            pulumi.BoolPtrInput
	Location           pulumi.StringPtrInput
	NamespaceName      pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Sku                SkuPtrInput
	Status             NamespaceStateEnumPtrInput
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

func (o NamespaceOutput) CreateACSNamespace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.CreateACSNamespace }).(pulumi.BoolPtrOutput)
}

func (o NamespaceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o NamespaceOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o NamespaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
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

func (o NamespaceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
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

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}
