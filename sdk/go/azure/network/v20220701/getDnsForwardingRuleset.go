


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDnsForwardingRuleset(ctx *pulumi.Context, args *LookupDnsForwardingRulesetArgs, opts ...pulumi.InvokeOption) (*LookupDnsForwardingRulesetResult, error) {
	var rv LookupDnsForwardingRulesetResult
	err := ctx.Invoke("azure-native:network/v20220701:getDnsForwardingRuleset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDnsForwardingRulesetArgs struct {
	DnsForwardingRulesetName string `pulumi:"dnsForwardingRulesetName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type LookupDnsForwardingRulesetResult struct {
	DnsResolverOutboundEndpoints []SubResourceResponse `pulumi:"dnsResolverOutboundEndpoints"`
	Etag                         string                `pulumi:"etag"`
	Id                           string                `pulumi:"id"`
	Location                     string                `pulumi:"location"`
	Name                         string                `pulumi:"name"`
	ProvisioningState            string                `pulumi:"provisioningState"`
	ResourceGuid                 string                `pulumi:"resourceGuid"`
	SystemData                   SystemDataResponse    `pulumi:"systemData"`
	Tags                         map[string]string     `pulumi:"tags"`
	Type                         string                `pulumi:"type"`
}

func LookupDnsForwardingRulesetOutput(ctx *pulumi.Context, args LookupDnsForwardingRulesetOutputArgs, opts ...pulumi.InvokeOption) LookupDnsForwardingRulesetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDnsForwardingRulesetResult, error) {
			args := v.(LookupDnsForwardingRulesetArgs)
			r, err := LookupDnsForwardingRuleset(ctx, &args, opts...)
			var s LookupDnsForwardingRulesetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDnsForwardingRulesetResultOutput)
}

type LookupDnsForwardingRulesetOutputArgs struct {
	DnsForwardingRulesetName pulumi.StringInput `pulumi:"dnsForwardingRulesetName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDnsForwardingRulesetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsForwardingRulesetArgs)(nil)).Elem()
}


type LookupDnsForwardingRulesetResultOutput struct{ *pulumi.OutputState }

func (LookupDnsForwardingRulesetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsForwardingRulesetResult)(nil)).Elem()
}

func (o LookupDnsForwardingRulesetResultOutput) ToLookupDnsForwardingRulesetResultOutput() LookupDnsForwardingRulesetResultOutput {
	return o
}

func (o LookupDnsForwardingRulesetResultOutput) ToLookupDnsForwardingRulesetResultOutputWithContext(ctx context.Context) LookupDnsForwardingRulesetResultOutput {
	return o
}

func (o LookupDnsForwardingRulesetResultOutput) DnsResolverOutboundEndpoints() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupDnsForwardingRulesetResult) []SubResourceResponse { return v.DnsResolverOutboundEndpoints }).(SubResourceResponseArrayOutput)
}

func (o LookupDnsForwardingRulesetResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsForwardingRulesetResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDnsForwardingRulesetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsForwardingRulesetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDnsForwardingRulesetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsForwardingRulesetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDnsForwardingRulesetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsForwardingRulesetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDnsForwardingRulesetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsForwardingRulesetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDnsForwardingRulesetResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsForwardingRulesetResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupDnsForwardingRulesetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDnsForwardingRulesetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDnsForwardingRulesetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDnsForwardingRulesetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDnsForwardingRulesetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsForwardingRulesetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDnsForwardingRulesetResultOutput{})
}
