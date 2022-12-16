


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkDnsService(ctx *pulumi.Context, args *LookupWorkloadNetworkDnsServiceArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkDnsServiceResult, error) {
	var rv LookupWorkloadNetworkDnsServiceResult
	err := ctx.Invoke("azure-native:avs/v20210101preview:getWorkloadNetworkDnsService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkDnsServiceArgs struct {
	DnsServiceId      string `pulumi:"dnsServiceId"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWorkloadNetworkDnsServiceResult struct {
	DefaultDnsZone    *string  `pulumi:"defaultDnsZone"`
	DisplayName       *string  `pulumi:"displayName"`
	DnsServiceIp      *string  `pulumi:"dnsServiceIp"`
	FqdnZones         []string `pulumi:"fqdnZones"`
	Id                string   `pulumi:"id"`
	LogLevel          *string  `pulumi:"logLevel"`
	Name              string   `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	Status            string   `pulumi:"status"`
	Type              string   `pulumi:"type"`
}

func LookupWorkloadNetworkDnsServiceOutput(ctx *pulumi.Context, args LookupWorkloadNetworkDnsServiceOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadNetworkDnsServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadNetworkDnsServiceResult, error) {
			args := v.(LookupWorkloadNetworkDnsServiceArgs)
			r, err := LookupWorkloadNetworkDnsService(ctx, &args, opts...)
			var s LookupWorkloadNetworkDnsServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadNetworkDnsServiceResultOutput)
}

type LookupWorkloadNetworkDnsServiceOutputArgs struct {
	DnsServiceId      pulumi.StringInput `pulumi:"dnsServiceId"`
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWorkloadNetworkDnsServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkDnsServiceArgs)(nil)).Elem()
}


type LookupWorkloadNetworkDnsServiceResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadNetworkDnsServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkDnsServiceResult)(nil)).Elem()
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) ToLookupWorkloadNetworkDnsServiceResultOutput() LookupWorkloadNetworkDnsServiceResultOutput {
	return o
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) ToLookupWorkloadNetworkDnsServiceResultOutputWithContext(ctx context.Context) LookupWorkloadNetworkDnsServiceResultOutput {
	return o
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) DefaultDnsZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) *string { return v.DefaultDnsZone }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) DnsServiceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) *string { return v.DnsServiceIp }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) FqdnZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) []string { return v.FqdnZones }).(pulumi.StringArrayOutput)
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) *string { return v.LogLevel }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDnsServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadNetworkDnsServiceResultOutput{})
}
