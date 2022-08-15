


package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLabPlan(ctx *pulumi.Context, args *LookupLabPlanArgs, opts ...pulumi.InvokeOption) (*LookupLabPlanResult, error) {
	var rv LookupLabPlanResult
	err := ctx.Invoke("azure-native:labservices/v20220801:getLabPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupLabPlanArgs struct {
	LabPlanName       string `pulumi:"labPlanName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLabPlanResult struct {
	AllowedRegions             []string                       `pulumi:"allowedRegions"`
	DefaultAutoShutdownProfile *AutoShutdownProfileResponse   `pulumi:"defaultAutoShutdownProfile"`
	DefaultConnectionProfile   *ConnectionProfileResponse     `pulumi:"defaultConnectionProfile"`
	DefaultNetworkProfile      *LabPlanNetworkProfileResponse `pulumi:"defaultNetworkProfile"`
	Id                         string                         `pulumi:"id"`
	Identity                   *IdentityResponse              `pulumi:"identity"`
	LinkedLmsInstance          *string                        `pulumi:"linkedLmsInstance"`
	Location                   string                         `pulumi:"location"`
	Name                       string                         `pulumi:"name"`
	ProvisioningState          string                         `pulumi:"provisioningState"`
	SharedGalleryId            *string                        `pulumi:"sharedGalleryId"`
	SupportInfo                *SupportInfoResponse           `pulumi:"supportInfo"`
	SystemData                 SystemDataResponse             `pulumi:"systemData"`
	Tags                       map[string]string              `pulumi:"tags"`
	Type                       string                         `pulumi:"type"`
}


func (val *LookupLabPlanResult) Defaults() *LookupLabPlanResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DefaultAutoShutdownProfile = tmp.DefaultAutoShutdownProfile.Defaults()

	tmp.DefaultConnectionProfile = tmp.DefaultConnectionProfile.Defaults()

	return &tmp
}

func LookupLabPlanOutput(ctx *pulumi.Context, args LookupLabPlanOutputArgs, opts ...pulumi.InvokeOption) LookupLabPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabPlanResult, error) {
			args := v.(LookupLabPlanArgs)
			r, err := LookupLabPlan(ctx, &args, opts...)
			var s LookupLabPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabPlanResultOutput)
}

type LookupLabPlanOutputArgs struct {
	LabPlanName       pulumi.StringInput `pulumi:"labPlanName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLabPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabPlanArgs)(nil)).Elem()
}


type LookupLabPlanResultOutput struct{ *pulumi.OutputState }

func (LookupLabPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabPlanResult)(nil)).Elem()
}

func (o LookupLabPlanResultOutput) ToLookupLabPlanResultOutput() LookupLabPlanResultOutput {
	return o
}

func (o LookupLabPlanResultOutput) ToLookupLabPlanResultOutputWithContext(ctx context.Context) LookupLabPlanResultOutput {
	return o
}

func (o LookupLabPlanResultOutput) AllowedRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLabPlanResult) []string { return v.AllowedRegions }).(pulumi.StringArrayOutput)
}

func (o LookupLabPlanResultOutput) DefaultAutoShutdownProfile() AutoShutdownProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupLabPlanResult) *AutoShutdownProfileResponse { return v.DefaultAutoShutdownProfile }).(AutoShutdownProfileResponsePtrOutput)
}

func (o LookupLabPlanResultOutput) DefaultConnectionProfile() ConnectionProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupLabPlanResult) *ConnectionProfileResponse { return v.DefaultConnectionProfile }).(ConnectionProfileResponsePtrOutput)
}

func (o LookupLabPlanResultOutput) DefaultNetworkProfile() LabPlanNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupLabPlanResult) *LabPlanNetworkProfileResponse { return v.DefaultNetworkProfile }).(LabPlanNetworkProfileResponsePtrOutput)
}

func (o LookupLabPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLabPlanResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupLabPlanResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupLabPlanResultOutput) LinkedLmsInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabPlanResult) *string { return v.LinkedLmsInstance }).(pulumi.StringPtrOutput)
}

func (o LookupLabPlanResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabPlanResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupLabPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLabPlanResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabPlanResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLabPlanResultOutput) SharedGalleryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabPlanResult) *string { return v.SharedGalleryId }).(pulumi.StringPtrOutput)
}

func (o LookupLabPlanResultOutput) SupportInfo() SupportInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupLabPlanResult) *SupportInfoResponse { return v.SupportInfo }).(SupportInfoResponsePtrOutput)
}

func (o LookupLabPlanResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLabPlanResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLabPlanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLabPlanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLabPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabPlanResultOutput{})
}
