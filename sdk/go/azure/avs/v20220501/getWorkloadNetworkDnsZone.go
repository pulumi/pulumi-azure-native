


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkDnsZone(ctx *pulumi.Context, args *LookupWorkloadNetworkDnsZoneArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkDnsZoneResult, error) {
	var rv LookupWorkloadNetworkDnsZoneResult
	err := ctx.Invoke("azure-native:avs/v20220501:getWorkloadNetworkDnsZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkDnsZoneArgs struct {
	DnsZoneId         string `pulumi:"dnsZoneId"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWorkloadNetworkDnsZoneResult struct {
	DisplayName       *string  `pulumi:"displayName"`
	DnsServerIps      []string `pulumi:"dnsServerIps"`
	DnsServices       *float64 `pulumi:"dnsServices"`
	Domain            []string `pulumi:"domain"`
	Id                string   `pulumi:"id"`
	Name              string   `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	SourceIp          *string  `pulumi:"sourceIp"`
	Type              string   `pulumi:"type"`
}

func LookupWorkloadNetworkDnsZoneOutput(ctx *pulumi.Context, args LookupWorkloadNetworkDnsZoneOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadNetworkDnsZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadNetworkDnsZoneResult, error) {
			args := v.(LookupWorkloadNetworkDnsZoneArgs)
			r, err := LookupWorkloadNetworkDnsZone(ctx, &args, opts...)
			var s LookupWorkloadNetworkDnsZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadNetworkDnsZoneResultOutput)
}

type LookupWorkloadNetworkDnsZoneOutputArgs struct {
	DnsZoneId         pulumi.StringInput `pulumi:"dnsZoneId"`
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWorkloadNetworkDnsZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkDnsZoneArgs)(nil)).Elem()
}


type LookupWorkloadNetworkDnsZoneResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadNetworkDnsZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkDnsZoneResult)(nil)).Elem()
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) ToLookupWorkloadNetworkDnsZoneResultOutput() LookupWorkloadNetworkDnsZoneResultOutput {
	return o
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) ToLookupWorkloadNetworkDnsZoneResultOutputWithContext(ctx context.Context) LookupWorkloadNetworkDnsZoneResultOutput {
	return o
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) DnsServerIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) []string { return v.DnsServerIps }).(pulumi.StringArrayOutput)
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) DnsServices() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) *float64 { return v.DnsServices }).(pulumi.Float64PtrOutput)
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) Domain() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) []string { return v.Domain }).(pulumi.StringArrayOutput)
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) *string { return v.SourceIp }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadNetworkDnsZoneResultOutput{})
}
