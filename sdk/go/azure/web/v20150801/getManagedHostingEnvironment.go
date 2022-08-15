


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupManagedHostingEnvironment(ctx *pulumi.Context, args *LookupManagedHostingEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupManagedHostingEnvironmentResult, error) {
	var rv LookupManagedHostingEnvironmentResult
	err := ctx.Invoke("azure-native:web/v20150801:getManagedHostingEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedHostingEnvironmentArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagedHostingEnvironmentResult struct {
	ApiManagementAccount *string                        `pulumi:"apiManagementAccount"`
	DnsSuffix            *string                        `pulumi:"dnsSuffix"`
	EnvironmentIsHealthy *bool                          `pulumi:"environmentIsHealthy"`
	EnvironmentStatus    *string                        `pulumi:"environmentStatus"`
	Id                   *string                        `pulumi:"id"`
	IpsslAddressCount    *int                           `pulumi:"ipsslAddressCount"`
	Kind                 *string                        `pulumi:"kind"`
	Location             string                         `pulumi:"location"`
	Name                 *string                        `pulumi:"name"`
	ResourceGroup        *string                        `pulumi:"resourceGroup"`
	Status               string                         `pulumi:"status"`
	SubscriptionId       *string                        `pulumi:"subscriptionId"`
	Suspended            *bool                          `pulumi:"suspended"`
	Tags                 map[string]string              `pulumi:"tags"`
	Type                 *string                        `pulumi:"type"`
	VirtualNetwork       *VirtualNetworkProfileResponse `pulumi:"virtualNetwork"`
}

func LookupManagedHostingEnvironmentOutput(ctx *pulumi.Context, args LookupManagedHostingEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupManagedHostingEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedHostingEnvironmentResult, error) {
			args := v.(LookupManagedHostingEnvironmentArgs)
			r, err := LookupManagedHostingEnvironment(ctx, &args, opts...)
			var s LookupManagedHostingEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedHostingEnvironmentResultOutput)
}

type LookupManagedHostingEnvironmentOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedHostingEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedHostingEnvironmentArgs)(nil)).Elem()
}


type LookupManagedHostingEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupManagedHostingEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedHostingEnvironmentResult)(nil)).Elem()
}

func (o LookupManagedHostingEnvironmentResultOutput) ToLookupManagedHostingEnvironmentResultOutput() LookupManagedHostingEnvironmentResultOutput {
	return o
}

func (o LookupManagedHostingEnvironmentResultOutput) ToLookupManagedHostingEnvironmentResultOutputWithContext(ctx context.Context) LookupManagedHostingEnvironmentResultOutput {
	return o
}

func (o LookupManagedHostingEnvironmentResultOutput) ApiManagementAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *string { return v.ApiManagementAccount }).(pulumi.StringPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *string { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) EnvironmentIsHealthy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *bool { return v.EnvironmentIsHealthy }).(pulumi.BoolPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) EnvironmentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *string { return v.EnvironmentStatus }).(pulumi.StringPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) IpsslAddressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *int { return v.IpsslAddressCount }).(pulumi.IntPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) Suspended() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *bool { return v.Suspended }).(pulumi.BoolPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupManagedHostingEnvironmentResultOutput) VirtualNetwork() VirtualNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedHostingEnvironmentResult) *VirtualNetworkProfileResponse { return v.VirtualNetwork }).(VirtualNetworkProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedHostingEnvironmentResultOutput{})
}
