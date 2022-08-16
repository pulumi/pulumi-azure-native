


package v20211130

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDedicatedHsm(ctx *pulumi.Context, args *LookupDedicatedHsmArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHsmResult, error) {
	var rv LookupDedicatedHsmResult
	err := ctx.Invoke("azure-native:hardwaresecuritymodules/v20211130:getDedicatedHsm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHsmArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDedicatedHsmResult struct {
	Id                       string                  `pulumi:"id"`
	Location                 string                  `pulumi:"location"`
	ManagementNetworkProfile *NetworkProfileResponse `pulumi:"managementNetworkProfile"`
	Name                     string                  `pulumi:"name"`
	NetworkProfile           *NetworkProfileResponse `pulumi:"networkProfile"`
	ProvisioningState        string                  `pulumi:"provisioningState"`
	Sku                      SkuResponse             `pulumi:"sku"`
	StampId                  *string                 `pulumi:"stampId"`
	StatusMessage            string                  `pulumi:"statusMessage"`
	SystemData               SystemDataResponse      `pulumi:"systemData"`
	Tags                     map[string]string       `pulumi:"tags"`
	Type                     string                  `pulumi:"type"`
	Zones                    []string                `pulumi:"zones"`
}

func LookupDedicatedHsmOutput(ctx *pulumi.Context, args LookupDedicatedHsmOutputArgs, opts ...pulumi.InvokeOption) LookupDedicatedHsmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDedicatedHsmResult, error) {
			args := v.(LookupDedicatedHsmArgs)
			r, err := LookupDedicatedHsm(ctx, &args, opts...)
			var s LookupDedicatedHsmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDedicatedHsmResultOutput)
}

type LookupDedicatedHsmOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDedicatedHsmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHsmArgs)(nil)).Elem()
}


type LookupDedicatedHsmResultOutput struct{ *pulumi.OutputState }

func (LookupDedicatedHsmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHsmResult)(nil)).Elem()
}

func (o LookupDedicatedHsmResultOutput) ToLookupDedicatedHsmResultOutput() LookupDedicatedHsmResultOutput {
	return o
}

func (o LookupDedicatedHsmResultOutput) ToLookupDedicatedHsmResultOutputWithContext(ctx context.Context) LookupDedicatedHsmResultOutput {
	return o
}

func (o LookupDedicatedHsmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDedicatedHsmResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDedicatedHsmResultOutput) ManagementNetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) *NetworkProfileResponse { return v.ManagementNetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o LookupDedicatedHsmResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDedicatedHsmResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o LookupDedicatedHsmResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDedicatedHsmResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupDedicatedHsmResultOutput) StampId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) *string { return v.StampId }).(pulumi.StringPtrOutput)
}

func (o LookupDedicatedHsmResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

func (o LookupDedicatedHsmResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDedicatedHsmResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDedicatedHsmResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDedicatedHsmResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDedicatedHsmResultOutput{})
}
