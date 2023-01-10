


package v20170401

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
	DataCenter         pulumi.StringPtrOutput `pulumi:"dataCenter"`
	Enabled            pulumi.BoolPtrOutput   `pulumi:"enabled"`
	Location           pulumi.StringPtrOutput `pulumi:"location"`
	MetricId           pulumi.StringOutput    `pulumi:"metricId"`
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
	UpdatedAt          pulumi.StringPtrOutput `pulumi:"updatedAt"`
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
			Type: pulumi.String("azure-native:notificationhubs/v20160301:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:notificationhubs/v20170401:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:notificationhubs/v20170401:Namespace", name, id, state, &resource, opts...)
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
	DataCenter         *string           `pulumi:"dataCenter"`
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
	UpdatedAt          *string           `pulumi:"updatedAt"`
}


type NamespaceArgs struct {
	CreatedAt          pulumi.StringPtrInput
	Critical           pulumi.BoolPtrInput
	DataCenter         pulumi.StringPtrInput
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
	UpdatedAt          pulumi.StringPtrInput
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

func (o NamespaceOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) Critical() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.Critical }).(pulumi.BoolPtrOutput)
}

func (o NamespaceOutput) DataCenter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.DataCenter }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o NamespaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.MetricId }).(pulumi.StringOutput)
}

func (o NamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NamespaceOutput) NamespaceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.NamespaceType }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) ScaleUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.ScaleUnit }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) ServiceBusEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.ServiceBusEndpoint }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o NamespaceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NamespaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NamespaceOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}
