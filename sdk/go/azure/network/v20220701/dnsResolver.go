


package v20220701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DnsResolver struct {
	pulumi.CustomResourceState

	DnsResolverState  pulumi.StringOutput       `pulumi:"dnsResolverState"`
	Etag              pulumi.StringOutput       `pulumi:"etag"`
	Location          pulumi.StringOutput       `pulumi:"location"`
	Name              pulumi.StringOutput       `pulumi:"name"`
	ProvisioningState pulumi.StringOutput       `pulumi:"provisioningState"`
	ResourceGuid      pulumi.StringOutput       `pulumi:"resourceGuid"`
	SystemData        SystemDataResponseOutput  `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput    `pulumi:"tags"`
	Type              pulumi.StringOutput       `pulumi:"type"`
	VirtualNetwork    SubResourceResponseOutput `pulumi:"virtualNetwork"`
}


func NewDnsResolver(ctx *pulumi.Context,
	name string, args *DnsResolverArgs, opts ...pulumi.ResourceOption) (*DnsResolver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetwork == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetwork'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:DnsResolver"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401preview:DnsResolver"),
		},
	})
	opts = append(opts, aliases)
	var resource DnsResolver
	err := ctx.RegisterResource("azure-native:network/v20220701:DnsResolver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDnsResolver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsResolverState, opts ...pulumi.ResourceOption) (*DnsResolver, error) {
	var resource DnsResolver
	err := ctx.ReadResource("azure-native:network/v20220701:DnsResolver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dnsResolverState struct {
}

type DnsResolverState struct {
}

func (DnsResolverState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsResolverState)(nil)).Elem()
}

type dnsResolverArgs struct {
	DnsResolverName   *string           `pulumi:"dnsResolverName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VirtualNetwork    SubResource       `pulumi:"virtualNetwork"`
}


type DnsResolverArgs struct {
	DnsResolverName   pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VirtualNetwork    SubResourceInput
}

func (DnsResolverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsResolverArgs)(nil)).Elem()
}

type DnsResolverInput interface {
	pulumi.Input

	ToDnsResolverOutput() DnsResolverOutput
	ToDnsResolverOutputWithContext(ctx context.Context) DnsResolverOutput
}

func (*DnsResolver) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsResolver)(nil)).Elem()
}

func (i *DnsResolver) ToDnsResolverOutput() DnsResolverOutput {
	return i.ToDnsResolverOutputWithContext(context.Background())
}

func (i *DnsResolver) ToDnsResolverOutputWithContext(ctx context.Context) DnsResolverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsResolverOutput)
}

type DnsResolverOutput struct{ *pulumi.OutputState }

func (DnsResolverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsResolver)(nil)).Elem()
}

func (o DnsResolverOutput) ToDnsResolverOutput() DnsResolverOutput {
	return o
}

func (o DnsResolverOutput) ToDnsResolverOutputWithContext(ctx context.Context) DnsResolverOutput {
	return o
}

func (o DnsResolverOutput) DnsResolverState() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolver) pulumi.StringOutput { return v.DnsResolverState }).(pulumi.StringOutput)
}

func (o DnsResolverOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolver) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DnsResolverOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolver) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DnsResolverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolver) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DnsResolverOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolver) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DnsResolverOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolver) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o DnsResolverOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DnsResolver) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DnsResolverOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DnsResolver) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DnsResolverOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolver) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DnsResolverOutput) VirtualNetwork() SubResourceResponseOutput {
	return o.ApplyT(func(v *DnsResolver) SubResourceResponseOutput { return v.VirtualNetwork }).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(DnsResolverOutput{})
}
