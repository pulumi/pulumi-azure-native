


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkDhcp(ctx *pulumi.Context, args *LookupWorkloadNetworkDhcpArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkDhcpResult, error) {
	var rv LookupWorkloadNetworkDhcpResult
	err := ctx.Invoke("azure-native:avs/v20210601:getWorkloadNetworkDhcp", args, &rv, opts...)
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
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
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

func (o LookupWorkloadNetworkDhcpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDhcpResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkDhcpResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupWorkloadNetworkDhcpResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDhcpResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadNetworkDhcpResultOutput{})
}
