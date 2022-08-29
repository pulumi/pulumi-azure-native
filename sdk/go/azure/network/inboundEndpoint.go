


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InboundEndpoint struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput                               `pulumi:"etag"`
	IpConfigurations  InboundEndpointIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	Location          pulumi.StringOutput                               `pulumi:"location"`
	Name              pulumi.StringOutput                               `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                               `pulumi:"provisioningState"`
	ResourceGuid      pulumi.StringOutput                               `pulumi:"resourceGuid"`
	SystemData        SystemDataResponseOutput                          `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                            `pulumi:"tags"`
	Type              pulumi.StringOutput                               `pulumi:"type"`
}


func NewInboundEndpoint(ctx *pulumi.Context,
	name string, args *InboundEndpointArgs, opts ...pulumi.ResourceOption) (*InboundEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsResolverName == nil {
		return nil, errors.New("invalid value for required argument 'DnsResolverName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20200401preview:InboundEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:InboundEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource InboundEndpoint
	err := ctx.RegisterResource("azure-native:network:InboundEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetInboundEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InboundEndpointState, opts ...pulumi.ResourceOption) (*InboundEndpoint, error) {
	var resource InboundEndpoint
	err := ctx.ReadResource("azure-native:network:InboundEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type inboundEndpointState struct {
}

type InboundEndpointState struct {
}

func (InboundEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundEndpointState)(nil)).Elem()
}

type inboundEndpointArgs struct {
	DnsResolverName     string                           `pulumi:"dnsResolverName"`
	InboundEndpointName *string                          `pulumi:"inboundEndpointName"`
	IpConfigurations    []InboundEndpointIPConfiguration `pulumi:"ipConfigurations"`
	Location            *string                          `pulumi:"location"`
	ResourceGroupName   string                           `pulumi:"resourceGroupName"`
	Tags                map[string]string                `pulumi:"tags"`
}


type InboundEndpointArgs struct {
	DnsResolverName     pulumi.StringInput
	InboundEndpointName pulumi.StringPtrInput
	IpConfigurations    InboundEndpointIPConfigurationArrayInput
	Location            pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
}

func (InboundEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundEndpointArgs)(nil)).Elem()
}

type InboundEndpointInput interface {
	pulumi.Input

	ToInboundEndpointOutput() InboundEndpointOutput
	ToInboundEndpointOutputWithContext(ctx context.Context) InboundEndpointOutput
}

func (*InboundEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundEndpoint)(nil)).Elem()
}

func (i *InboundEndpoint) ToInboundEndpointOutput() InboundEndpointOutput {
	return i.ToInboundEndpointOutputWithContext(context.Background())
}

func (i *InboundEndpoint) ToInboundEndpointOutputWithContext(ctx context.Context) InboundEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundEndpointOutput)
}

type InboundEndpointOutput struct{ *pulumi.OutputState }

func (InboundEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundEndpoint)(nil)).Elem()
}

func (o InboundEndpointOutput) ToInboundEndpointOutput() InboundEndpointOutput {
	return o
}

func (o InboundEndpointOutput) ToInboundEndpointOutputWithContext(ctx context.Context) InboundEndpointOutput {
	return o
}

func (o InboundEndpointOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o InboundEndpointOutput) IpConfigurations() InboundEndpointIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *InboundEndpoint) InboundEndpointIPConfigurationResponseArrayOutput { return v.IpConfigurations }).(InboundEndpointIPConfigurationResponseArrayOutput)
}

func (o InboundEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o InboundEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o InboundEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o InboundEndpointOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o InboundEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *InboundEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o InboundEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o InboundEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InboundEndpointOutput{})
}
