


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkDhcp(ctx *pulumi.Context, args *LookupWorkloadNetworkDhcpArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkDhcpResult, error) {
	var rv LookupWorkloadNetworkDhcpResult
	err := ctx.Invoke("azure-native:avs/v20210101preview:getWorkloadNetworkDhcp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkDhcpArgs struct {
	DhcpId            string `pulumi:"dhcpId"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWorkloadNetworkDhcpResult struct {
	DhcpType          string   `pulumi:"dhcpType"`
	DisplayName       *string  `pulumi:"displayName"`
	Id                string   `pulumi:"id"`
	Name              string   `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	Segments          []string `pulumi:"segments"`
	Type              string   `pulumi:"type"`
}

func LookupWorkloadNetworkDhcpOutput(ctx *pulumi.Context, args LookupWorkloadNetworkDhcpOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadNetworkDhcpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadNetworkDhcpResult, error) {
			args := v.(LookupWorkloadNetworkDhcpArgs)
			r, err := LookupWorkloadNetworkDhcp(ctx, &args, opts...)
			var s LookupWorkloadNetworkDhcpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadNetworkDhcpResultOutput)
}

type LookupWorkloadNetworkDhcpOutputArgs struct {
	DhcpId            pulumi.StringInput `pulumi:"dhcpId"`
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWorkloadNetworkDhcpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkDhcpArgs)(nil)).Elem()
}


type LookupWorkloadNetworkDhcpResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadNetworkDhcpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkDhcpResult)(nil)).Elem()
}

func (o LookupWorkloadNetworkDhcpResultOutput) ToLookupWorkloadNetworkDhcpResultOutput() LookupWorkloadNetworkDhcpResultOutput {
	return o
}

func (o LookupWorkloadNetworkDhcpResultOutput) ToLookupWorkloadNetworkDhcpResultOutputWithContext(ctx context.Context) LookupWorkloadNetworkDhcpResultOutput {
	return o
}

func (o LookupWorkloadNetworkDhcpResultOutput) DhcpType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) string { return v.DhcpType }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDhcpResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkDhcpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDhcpResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDhcpResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDhcpResultOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o LookupWorkloadNetworkDhcpResultOutput) Segments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) []string { return v.Segments }).(pulumi.StringArrayOutput)
}

func (o LookupWorkloadNetworkDhcpResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadNetworkDhcpResultOutput{})
}
