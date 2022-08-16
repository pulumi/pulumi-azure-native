


package v20190401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:vmwarecloudsimple/v20190401:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VirtualMachineName string `pulumi:"virtualMachineName"`
}


type LookupVirtualMachineResult struct {
	AmountOfRam       int                             `pulumi:"amountOfRam"`
	Controllers       []VirtualDiskControllerResponse `pulumi:"controllers"`
	Customization     *GuestOSCustomizationResponse   `pulumi:"customization"`
	Disks             []VirtualDiskResponse           `pulumi:"disks"`
	Dnsname           string                          `pulumi:"dnsname"`
	ExposeToGuestVM   *bool                           `pulumi:"exposeToGuestVM"`
	Folder            string                          `pulumi:"folder"`
	GuestOS           string                          `pulumi:"guestOS"`
	GuestOSType       string                          `pulumi:"guestOSType"`
	Id                string                          `pulumi:"id"`
	Location          string                          `pulumi:"location"`
	Name              string                          `pulumi:"name"`
	Nics              []VirtualNicResponse            `pulumi:"nics"`
	NumberOfCores     int                             `pulumi:"numberOfCores"`
	Password          *string                         `pulumi:"password"`
	PrivateCloudId    string                          `pulumi:"privateCloudId"`
	ProvisioningState string                          `pulumi:"provisioningState"`
	PublicIP          string                          `pulumi:"publicIP"`
	ResourcePool      *ResourcePoolResponse           `pulumi:"resourcePool"`
	Status            string                          `pulumi:"status"`
	Tags              map[string]string               `pulumi:"tags"`
	TemplateId        *string                         `pulumi:"templateId"`
	Type              string                          `pulumi:"type"`
	Username          *string                         `pulumi:"username"`
	VSphereNetworks   []string                        `pulumi:"vSphereNetworks"`
	VmId              string                          `pulumi:"vmId"`
	Vmwaretools       string                          `pulumi:"vmwaretools"`
}

func LookupVirtualMachineOutput(ctx *pulumi.Context, args LookupVirtualMachineOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineResult, error) {
			args := v.(LookupVirtualMachineArgs)
			r, err := LookupVirtualMachine(ctx, &args, opts...)
			var s LookupVirtualMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineResultOutput)
}

type LookupVirtualMachineOutputArgs struct {
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualMachineName pulumi.StringInput `pulumi:"virtualMachineName"`
}

func (LookupVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineArgs)(nil)).Elem()
}


type LookupVirtualMachineResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineResult)(nil)).Elem()
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutput() LookupVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutputWithContext(ctx context.Context) LookupVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineResultOutput) AmountOfRam() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) int { return v.AmountOfRam }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineResultOutput) Controllers() VirtualDiskControllerResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []VirtualDiskControllerResponse { return v.Controllers }).(VirtualDiskControllerResponseArrayOutput)
}

func (o LookupVirtualMachineResultOutput) Customization() GuestOSCustomizationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *GuestOSCustomizationResponse { return v.Customization }).(GuestOSCustomizationResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []VirtualDiskResponse { return v.Disks }).(VirtualDiskResponseArrayOutput)
}

func (o LookupVirtualMachineResultOutput) Dnsname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Dnsname }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ExposeToGuestVM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.ExposeToGuestVM }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Folder }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) GuestOS() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.GuestOS }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) GuestOSType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.GuestOSType }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Nics() VirtualNicResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []VirtualNicResponse { return v.Nics }).(VirtualNicResponseArrayOutput)
}

func (o LookupVirtualMachineResultOutput) NumberOfCores() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) int { return v.NumberOfCores }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) PrivateCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.PrivateCloudId }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) PublicIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.PublicIP }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ResourcePool() ResourcePoolResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *ResourcePoolResponse { return v.ResourcePool }).(ResourcePoolResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineResultOutput) TemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.TemplateId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) VSphereNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []string { return v.VSphereNetworks }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualMachineResultOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.VmId }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Vmwaretools() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Vmwaretools }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
