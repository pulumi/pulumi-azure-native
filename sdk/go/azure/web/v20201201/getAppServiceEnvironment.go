


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceEnvironment(ctx *pulumi.Context, args *LookupAppServiceEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceEnvironmentResult, error) {
	var rv LookupAppServiceEnvironmentResult
	err := ctx.Invoke("azure-native:web/v20201201:getAppServiceEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceEnvironmentArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAppServiceEnvironmentResult struct {
	ClusterSettings           []NameValuePairResponse       `pulumi:"clusterSettings"`
	DedicatedHostCount        int                           `pulumi:"dedicatedHostCount"`
	DnsSuffix                 *string                       `pulumi:"dnsSuffix"`
	FrontEndScaleFactor       *int                          `pulumi:"frontEndScaleFactor"`
	HasLinuxWorkers           bool                          `pulumi:"hasLinuxWorkers"`
	Id                        string                        `pulumi:"id"`
	InternalLoadBalancingMode *string                       `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount         *int                          `pulumi:"ipsslAddressCount"`
	Kind                      *string                       `pulumi:"kind"`
	Location                  string                        `pulumi:"location"`
	MaximumNumberOfMachines   int                           `pulumi:"maximumNumberOfMachines"`
	MultiRoleCount            int                           `pulumi:"multiRoleCount"`
	MultiSize                 *string                       `pulumi:"multiSize"`
	Name                      string                        `pulumi:"name"`
	ProvisioningState         string                        `pulumi:"provisioningState"`
	Status                    string                        `pulumi:"status"`
	Suspended                 bool                          `pulumi:"suspended"`
	Tags                      map[string]string             `pulumi:"tags"`
	Type                      string                        `pulumi:"type"`
	UserWhitelistedIpRanges   []string                      `pulumi:"userWhitelistedIpRanges"`
	VirtualNetwork            VirtualNetworkProfileResponse `pulumi:"virtualNetwork"`
}

func LookupAppServiceEnvironmentOutput(ctx *pulumi.Context, args LookupAppServiceEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupAppServiceEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppServiceEnvironmentResult, error) {
			args := v.(LookupAppServiceEnvironmentArgs)
			r, err := LookupAppServiceEnvironment(ctx, &args, opts...)
			var s LookupAppServiceEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppServiceEnvironmentResultOutput)
}

type LookupAppServiceEnvironmentOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAppServiceEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceEnvironmentArgs)(nil)).Elem()
}


type LookupAppServiceEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupAppServiceEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceEnvironmentResult)(nil)).Elem()
}

func (o LookupAppServiceEnvironmentResultOutput) ToLookupAppServiceEnvironmentResultOutput() LookupAppServiceEnvironmentResultOutput {
	return o
}

func (o LookupAppServiceEnvironmentResultOutput) ToLookupAppServiceEnvironmentResultOutputWithContext(ctx context.Context) LookupAppServiceEnvironmentResultOutput {
	return o
}

func (o LookupAppServiceEnvironmentResultOutput) ClusterSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) []NameValuePairResponse { return v.ClusterSettings }).(NameValuePairResponseArrayOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) DedicatedHostCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) int { return v.DedicatedHostCount }).(pulumi.IntOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) FrontEndScaleFactor() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *int { return v.FrontEndScaleFactor }).(pulumi.IntPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) HasLinuxWorkers() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) bool { return v.HasLinuxWorkers }).(pulumi.BoolOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) InternalLoadBalancingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.InternalLoadBalancingMode }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) IpsslAddressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *int { return v.IpsslAddressCount }).(pulumi.IntPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) MaximumNumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) int { return v.MaximumNumberOfMachines }).(pulumi.IntOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) MultiRoleCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) int { return v.MultiRoleCount }).(pulumi.IntOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) MultiSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.MultiSize }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Suspended() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) bool { return v.Suspended }).(pulumi.BoolOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) UserWhitelistedIpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) []string { return v.UserWhitelistedIpRanges }).(pulumi.StringArrayOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) VirtualNetwork() VirtualNetworkProfileResponseOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) VirtualNetworkProfileResponse { return v.VirtualNetwork }).(VirtualNetworkProfileResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppServiceEnvironmentResultOutput{})
}
