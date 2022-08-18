


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDedicatedHost(ctx *pulumi.Context, args *LookupDedicatedHostArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostResult, error) {
	var rv LookupDedicatedHostResult
	err := ctx.Invoke("azure-native:compute/v20210701:getDedicatedHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHostArgs struct {
	Expand            *string `pulumi:"expand"`
	HostGroupName     string  `pulumi:"hostGroupName"`
	HostName          string  `pulumi:"hostName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupDedicatedHostResult struct {
	AutoReplaceOnFailure *bool                             `pulumi:"autoReplaceOnFailure"`
	HostId               string                            `pulumi:"hostId"`
	Id                   string                            `pulumi:"id"`
	InstanceView         DedicatedHostInstanceViewResponse `pulumi:"instanceView"`
	LicenseType          *string                           `pulumi:"licenseType"`
	Location             string                            `pulumi:"location"`
	Name                 string                            `pulumi:"name"`
	PlatformFaultDomain  *int                              `pulumi:"platformFaultDomain"`
	ProvisioningState    string                            `pulumi:"provisioningState"`
	ProvisioningTime     string                            `pulumi:"provisioningTime"`
	Sku                  SkuResponse                       `pulumi:"sku"`
	Tags                 map[string]string                 `pulumi:"tags"`
	Type                 string                            `pulumi:"type"`
	VirtualMachines      []SubResourceReadOnlyResponse     `pulumi:"virtualMachines"`
}

func LookupDedicatedHostOutput(ctx *pulumi.Context, args LookupDedicatedHostOutputArgs, opts ...pulumi.InvokeOption) LookupDedicatedHostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDedicatedHostResult, error) {
			args := v.(LookupDedicatedHostArgs)
			r, err := LookupDedicatedHost(ctx, &args, opts...)
			var s LookupDedicatedHostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDedicatedHostResultOutput)
}

type LookupDedicatedHostOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	HostGroupName     pulumi.StringInput    `pulumi:"hostGroupName"`
	HostName          pulumi.StringInput    `pulumi:"hostName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupDedicatedHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHostArgs)(nil)).Elem()
}


type LookupDedicatedHostResultOutput struct{ *pulumi.OutputState }

func (LookupDedicatedHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHostResult)(nil)).Elem()
}

func (o LookupDedicatedHostResultOutput) ToLookupDedicatedHostResultOutput() LookupDedicatedHostResultOutput {
	return o
}

func (o LookupDedicatedHostResultOutput) ToLookupDedicatedHostResultOutputWithContext(ctx context.Context) LookupDedicatedHostResultOutput {
	return o
}

func (o LookupDedicatedHostResultOutput) AutoReplaceOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) *bool { return v.AutoReplaceOnFailure }).(pulumi.BoolPtrOutput)
}

func (o LookupDedicatedHostResultOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.HostId }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostResultOutput) InstanceView() DedicatedHostInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) DedicatedHostInstanceViewResponse { return v.InstanceView }).(DedicatedHostInstanceViewResponseOutput)
}

func (o LookupDedicatedHostResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o LookupDedicatedHostResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostResultOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) *int { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

func (o LookupDedicatedHostResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostResultOutput) ProvisioningTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.ProvisioningTime }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupDedicatedHostResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDedicatedHostResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDedicatedHostResultOutput) VirtualMachines() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) []SubResourceReadOnlyResponse { return v.VirtualMachines }).(SubResourceReadOnlyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDedicatedHostResultOutput{})
}
