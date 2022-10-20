


package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SignalR struct {
	pulumi.CustomResourceState

	ExternalIP        pulumi.StringOutput          `pulumi:"externalIP"`
	HostName          pulumi.StringOutput          `pulumi:"hostName"`
	HostNamePrefix    pulumi.StringPtrOutput       `pulumi:"hostNamePrefix"`
	Location          pulumi.StringPtrOutput       `pulumi:"location"`
	Name              pulumi.StringOutput          `pulumi:"name"`
	ProvisioningState pulumi.StringOutput          `pulumi:"provisioningState"`
	PublicPort        pulumi.IntOutput             `pulumi:"publicPort"`
	ServerPort        pulumi.IntOutput             `pulumi:"serverPort"`
	Sku               ResourceSkuResponsePtrOutput `pulumi:"sku"`
	Tags              pulumi.StringMapOutput       `pulumi:"tags"`
	Type              pulumi.StringOutput          `pulumi:"type"`
	Version           pulumi.StringPtrOutput       `pulumi:"version"`
}


func NewSignalR(ctx *pulumi.Context,
	name string, args *SignalRArgs, opts ...pulumi.ResourceOption) (*SignalR, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:signalrservice:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20181001:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20200501:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20200701preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210401preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210601preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210901preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20211001:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220201:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220801preview:SignalR"),
		},
	})
	opts = append(opts, aliases)
	var resource SignalR
	err := ctx.RegisterResource("azure-native:signalrservice/v20180301preview:SignalR", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSignalR(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRState, opts ...pulumi.ResourceOption) (*SignalR, error) {
	var resource SignalR
	err := ctx.ReadResource("azure-native:signalrservice/v20180301preview:SignalR", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type signalRState struct {
}

type SignalRState struct {
}

func (SignalRState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRState)(nil)).Elem()
}

type signalRArgs struct {
	Location          *string                          `pulumi:"location"`
	Properties        *SignalRCreateOrUpdateProperties `pulumi:"properties"`
	ResourceGroupName string                           `pulumi:"resourceGroupName"`
	ResourceName      *string                          `pulumi:"resourceName"`
	Sku               *ResourceSku                     `pulumi:"sku"`
	Tags              map[string]string                `pulumi:"tags"`
}


type SignalRArgs struct {
	Location          pulumi.StringPtrInput
	Properties        SignalRCreateOrUpdatePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Sku               ResourceSkuPtrInput
	Tags              pulumi.StringMapInput
}

func (SignalRArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRArgs)(nil)).Elem()
}

type SignalRInput interface {
	pulumi.Input

	ToSignalROutput() SignalROutput
	ToSignalROutputWithContext(ctx context.Context) SignalROutput
}

func (*SignalR) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalR)(nil)).Elem()
}

func (i *SignalR) ToSignalROutput() SignalROutput {
	return i.ToSignalROutputWithContext(context.Background())
}

func (i *SignalR) ToSignalROutputWithContext(ctx context.Context) SignalROutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalROutput)
}

type SignalROutput struct{ *pulumi.OutputState }

func (SignalROutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalR)(nil)).Elem()
}

func (o SignalROutput) ToSignalROutput() SignalROutput {
	return o
}

func (o SignalROutput) ToSignalROutputWithContext(ctx context.Context) SignalROutput {
	return o
}

func (o SignalROutput) ExternalIP() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.ExternalIP }).(pulumi.StringOutput)
}

func (o SignalROutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o SignalROutput) HostNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.HostNamePrefix }).(pulumi.StringPtrOutput)
}

func (o SignalROutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o SignalROutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SignalROutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SignalROutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SignalR) pulumi.IntOutput { return v.PublicPort }).(pulumi.IntOutput)
}

func (o SignalROutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SignalR) pulumi.IntOutput { return v.ServerPort }).(pulumi.IntOutput)
}

func (o SignalROutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

func (o SignalROutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SignalROutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SignalROutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalROutput{})
}
