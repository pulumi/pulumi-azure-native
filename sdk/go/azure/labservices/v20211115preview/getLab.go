


package v20211115preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:labservices/v20211115preview:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupLabArgs struct {
	LabName           string `pulumi:"labName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLabResult struct {
	AutoShutdownProfile   AutoShutdownProfileResponse   `pulumi:"autoShutdownProfile"`
	ConnectionProfile     ConnectionProfileResponse     `pulumi:"connectionProfile"`
	Description           *string                       `pulumi:"description"`
	Id                    string                        `pulumi:"id"`
	LabPlanId             *string                       `pulumi:"labPlanId"`
	Location              string                        `pulumi:"location"`
	Name                  string                        `pulumi:"name"`
	NetworkProfile        *LabNetworkProfileResponse    `pulumi:"networkProfile"`
	ProvisioningState     string                        `pulumi:"provisioningState"`
	RosterProfile         *RosterProfileResponse        `pulumi:"rosterProfile"`
	SecurityProfile       SecurityProfileResponse       `pulumi:"securityProfile"`
	State                 string                        `pulumi:"state"`
	SystemData            SystemDataResponse            `pulumi:"systemData"`
	Tags                  map[string]string             `pulumi:"tags"`
	Title                 *string                       `pulumi:"title"`
	Type                  string                        `pulumi:"type"`
	VirtualMachineProfile VirtualMachineProfileResponse `pulumi:"virtualMachineProfile"`
}


func (val *LookupLabResult) Defaults() *LookupLabResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.AutoShutdownProfile = *tmp.AutoShutdownProfile.Defaults()

	tmp.ConnectionProfile = *tmp.ConnectionProfile.Defaults()

	tmp.VirtualMachineProfile = *tmp.VirtualMachineProfile.Defaults()

	return &tmp
}

func LookupLabOutput(ctx *pulumi.Context, args LookupLabOutputArgs, opts ...pulumi.InvokeOption) LookupLabResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabResult, error) {
			args := v.(LookupLabArgs)
			r, err := LookupLab(ctx, &args, opts...)
			var s LookupLabResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabResultOutput)
}

type LookupLabOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLabOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabArgs)(nil)).Elem()
}


type LookupLabResultOutput struct{ *pulumi.OutputState }

func (LookupLabResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabResult)(nil)).Elem()
}

func (o LookupLabResultOutput) ToLookupLabResultOutput() LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) ToLookupLabResultOutputWithContext(ctx context.Context) LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) AutoShutdownProfile() AutoShutdownProfileResponseOutput {
	return o.ApplyT(func(v LookupLabResult) AutoShutdownProfileResponse { return v.AutoShutdownProfile }).(AutoShutdownProfileResponseOutput)
}

func (o LookupLabResultOutput) ConnectionProfile() ConnectionProfileResponseOutput {
	return o.ApplyT(func(v LookupLabResult) ConnectionProfileResponse { return v.ConnectionProfile }).(ConnectionProfileResponseOutput)
}

func (o LookupLabResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) LabPlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.LabPlanId }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) NetworkProfile() LabNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupLabResult) *LabNetworkProfileResponse { return v.NetworkProfile }).(LabNetworkProfileResponsePtrOutput)
}

func (o LookupLabResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) RosterProfile() RosterProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupLabResult) *RosterProfileResponse { return v.RosterProfile }).(RosterProfileResponsePtrOutput)
}

func (o LookupLabResultOutput) SecurityProfile() SecurityProfileResponseOutput {
	return o.ApplyT(func(v LookupLabResult) SecurityProfileResponse { return v.SecurityProfile }).(SecurityProfileResponseOutput)
}

func (o LookupLabResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLabResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLabResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLabResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLabResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) VirtualMachineProfile() VirtualMachineProfileResponseOutput {
	return o.ApplyT(func(v LookupLabResult) VirtualMachineProfileResponse { return v.VirtualMachineProfile }).(VirtualMachineProfileResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabResultOutput{})
}
