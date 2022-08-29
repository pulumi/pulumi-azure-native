


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DnsForwardingRuleset struct {
	pulumi.CustomResourceState

	DnsResolverOutboundEndpoints SubResourceResponseArrayOutput `pulumi:"dnsResolverOutboundEndpoints"`
	Etag                         pulumi.StringOutput            `pulumi:"etag"`
	Location                     pulumi.StringOutput            `pulumi:"location"`
	Name                         pulumi.StringOutput            `pulumi:"name"`
	ProvisioningState            pulumi.StringOutput            `pulumi:"provisioningState"`
	ResourceGuid                 pulumi.StringOutput            `pulumi:"resourceGuid"`
	SystemData                   SystemDataResponseOutput       `pulumi:"systemData"`
	Tags                         pulumi.StringMapOutput         `pulumi:"tags"`
	Type                         pulumi.StringOutput            `pulumi:"type"`
}


func NewDnsForwardingRuleset(ctx *pulumi.Context,
	name string, args *DnsForwardingRulesetArgs, opts ...pulumi.ResourceOption) (*DnsForwardingRuleset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20200401preview:DnsForwardingRuleset"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:DnsForwardingRuleset"),
		},
	})
	opts = append(opts, aliases)
	var resource DnsForwardingRuleset
	err := ctx.RegisterResource("azure-native:network:DnsForwardingRuleset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDnsForwardingRuleset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsForwardingRulesetState, opts ...pulumi.ResourceOption) (*DnsForwardingRuleset, error) {
	var resource DnsForwardingRuleset
	err := ctx.ReadResource("azure-native:network:DnsForwardingRuleset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dnsForwardingRulesetState struct {
}

type DnsForwardingRulesetState struct {
}

func (DnsForwardingRulesetState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsForwardingRulesetState)(nil)).Elem()
}

type dnsForwardingRulesetArgs struct {
	DnsForwardingRulesetName     *string           `pulumi:"dnsForwardingRulesetName"`
	DnsResolverOutboundEndpoints []SubResource     `pulumi:"dnsResolverOutboundEndpoints"`
	Location                     *string           `pulumi:"location"`
	ResourceGroupName            string            `pulumi:"resourceGroupName"`
	Tags                         map[string]string `pulumi:"tags"`
}


type DnsForwardingRulesetArgs struct {
	DnsForwardingRulesetName     pulumi.StringPtrInput
	DnsResolverOutboundEndpoints SubResourceArrayInput
	Location                     pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
}

func (DnsForwardingRulesetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsForwardingRulesetArgs)(nil)).Elem()
}

type DnsForwardingRulesetInput interface {
	pulumi.Input

	ToDnsForwardingRulesetOutput() DnsForwardingRulesetOutput
	ToDnsForwardingRulesetOutputWithContext(ctx context.Context) DnsForwardingRulesetOutput
}

func (*DnsForwardingRuleset) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsForwardingRuleset)(nil)).Elem()
}

func (i *DnsForwardingRuleset) ToDnsForwardingRulesetOutput() DnsForwardingRulesetOutput {
	return i.ToDnsForwardingRulesetOutputWithContext(context.Background())
}

func (i *DnsForwardingRuleset) ToDnsForwardingRulesetOutputWithContext(ctx context.Context) DnsForwardingRulesetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsForwardingRulesetOutput)
}

type DnsForwardingRulesetOutput struct{ *pulumi.OutputState }

func (DnsForwardingRulesetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsForwardingRuleset)(nil)).Elem()
}

func (o DnsForwardingRulesetOutput) ToDnsForwardingRulesetOutput() DnsForwardingRulesetOutput {
	return o
}

func (o DnsForwardingRulesetOutput) ToDnsForwardingRulesetOutputWithContext(ctx context.Context) DnsForwardingRulesetOutput {
	return o
}

func (o DnsForwardingRulesetOutput) DnsResolverOutboundEndpoints() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) SubResourceResponseArrayOutput { return v.DnsResolverOutboundEndpoints }).(SubResourceResponseArrayOutput)
}

func (o DnsForwardingRulesetOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DnsForwardingRulesetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DnsForwardingRulesetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DnsForwardingRulesetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DnsForwardingRulesetOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o DnsForwardingRulesetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DnsForwardingRulesetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DnsForwardingRulesetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DnsForwardingRulesetOutput{})
}
