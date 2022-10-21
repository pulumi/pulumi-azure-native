


package avs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkPublicIP(ctx *pulumi.Context, args *LookupWorkloadNetworkPublicIPArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkPublicIPResult, error) {
	var rv LookupWorkloadNetworkPublicIPResult
	err := ctx.Invoke("azure-native:avs:getWorkloadNetworkPublicIP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkPublicIPArgs struct {
	PrivateCloudName  string `pulumi:"privateCloudName"`
	PublicIPId        string `pulumi:"publicIPId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWorkloadNetworkPublicIPResult struct {
	DisplayName       *string  `pulumi:"displayName"`
	Id                string   `pulumi:"id"`
	Name              string   `pulumi:"name"`
	NumberOfPublicIPs *float64 `pulumi:"numberOfPublicIPs"`
	ProvisioningState string   `pulumi:"provisioningState"`
	PublicIPBlock     string   `pulumi:"publicIPBlock"`
	Type              string   `pulumi:"type"`
}

func LookupWorkloadNetworkPublicIPOutput(ctx *pulumi.Context, args LookupWorkloadNetworkPublicIPOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadNetworkPublicIPResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadNetworkPublicIPResult, error) {
			args := v.(LookupWorkloadNetworkPublicIPArgs)
			r, err := LookupWorkloadNetworkPublicIP(ctx, &args, opts...)
			var s LookupWorkloadNetworkPublicIPResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadNetworkPublicIPResultOutput)
}

type LookupWorkloadNetworkPublicIPOutputArgs struct {
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	PublicIPId        pulumi.StringInput `pulumi:"publicIPId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWorkloadNetworkPublicIPOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkPublicIPArgs)(nil)).Elem()
}


type LookupWorkloadNetworkPublicIPResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadNetworkPublicIPResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkPublicIPResult)(nil)).Elem()
}

func (o LookupWorkloadNetworkPublicIPResultOutput) ToLookupWorkloadNetworkPublicIPResultOutput() LookupWorkloadNetworkPublicIPResultOutput {
	return o
}

func (o LookupWorkloadNetworkPublicIPResultOutput) ToLookupWorkloadNetworkPublicIPResultOutputWithContext(ctx context.Context) LookupWorkloadNetworkPublicIPResultOutput {
	return o
}

func (o LookupWorkloadNetworkPublicIPResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPublicIPResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkPublicIPResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPublicIPResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkPublicIPResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPublicIPResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkPublicIPResultOutput) NumberOfPublicIPs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPublicIPResult) *float64 { return v.NumberOfPublicIPs }).(pulumi.Float64PtrOutput)
}

func (o LookupWorkloadNetworkPublicIPResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPublicIPResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkPublicIPResultOutput) PublicIPBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPublicIPResult) string { return v.PublicIPBlock }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkPublicIPResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPublicIPResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadNetworkPublicIPResultOutput{})
}
