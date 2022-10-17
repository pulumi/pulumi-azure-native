


package v20220701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OutboundEndpoint struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput       `pulumi:"etag"`
	Location          pulumi.StringOutput       `pulumi:"location"`
	Name              pulumi.StringOutput       `pulumi:"name"`
	ProvisioningState pulumi.StringOutput       `pulumi:"provisioningState"`
	ResourceGuid      pulumi.StringOutput       `pulumi:"resourceGuid"`
	Subnet            SubResourceResponseOutput `pulumi:"subnet"`
	SystemData        SystemDataResponseOutput  `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput    `pulumi:"tags"`
	Type              pulumi.StringOutput       `pulumi:"type"`
}


func NewOutboundEndpoint(ctx *pulumi.Context,
	name string, args *OutboundEndpointArgs, opts ...pulumi.ResourceOption) (*OutboundEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsResolverName == nil {
		return nil, errors.New("invalid value for required argument 'DnsResolverName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:OutboundEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401preview:OutboundEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource OutboundEndpoint
	err := ctx.RegisterResource("azure-native:network/v20220701:OutboundEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOutboundEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutboundEndpointState, opts ...pulumi.ResourceOption) (*OutboundEndpoint, error) {
	var resource OutboundEndpoint
	err := ctx.ReadResource("azure-native:network/v20220701:OutboundEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type outboundEndpointState struct {
}

type OutboundEndpointState struct {
}

func (OutboundEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*outboundEndpointState)(nil)).Elem()
}

type outboundEndpointArgs struct {
	DnsResolverName      string            `pulumi:"dnsResolverName"`
	Location             *string           `pulumi:"location"`
	OutboundEndpointName *string           `pulumi:"outboundEndpointName"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	Subnet               SubResource       `pulumi:"subnet"`
	Tags                 map[string]string `pulumi:"tags"`
}


type OutboundEndpointArgs struct {
	DnsResolverName      pulumi.StringInput
	Location             pulumi.StringPtrInput
	OutboundEndpointName pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Subnet               SubResourceInput
	Tags                 pulumi.StringMapInput
}

func (OutboundEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outboundEndpointArgs)(nil)).Elem()
}

type OutboundEndpointInput interface {
	pulumi.Input

	ToOutboundEndpointOutput() OutboundEndpointOutput
	ToOutboundEndpointOutputWithContext(ctx context.Context) OutboundEndpointOutput
}

func (*OutboundEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**OutboundEndpoint)(nil)).Elem()
}

func (i *OutboundEndpoint) ToOutboundEndpointOutput() OutboundEndpointOutput {
	return i.ToOutboundEndpointOutputWithContext(context.Background())
}

func (i *OutboundEndpoint) ToOutboundEndpointOutputWithContext(ctx context.Context) OutboundEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundEndpointOutput)
}

type OutboundEndpointOutput struct{ *pulumi.OutputState }

func (OutboundEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutboundEndpoint)(nil)).Elem()
}

func (o OutboundEndpointOutput) ToOutboundEndpointOutput() OutboundEndpointOutput {
	return o
}

func (o OutboundEndpointOutput) ToOutboundEndpointOutputWithContext(ctx context.Context) OutboundEndpointOutput {
	return o
}

func (o OutboundEndpointOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o OutboundEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o OutboundEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OutboundEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OutboundEndpointOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o OutboundEndpointOutput) Subnet() SubResourceResponseOutput {
	return o.ApplyT(func(v *OutboundEndpoint) SubResourceResponseOutput { return v.Subnet }).(SubResourceResponseOutput)
}

func (o OutboundEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OutboundEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o OutboundEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OutboundEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OutboundEndpointOutput{})
}
