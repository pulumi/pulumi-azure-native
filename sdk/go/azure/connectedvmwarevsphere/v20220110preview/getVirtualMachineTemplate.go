


package v20220110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineTemplate(ctx *pulumi.Context, args *LookupVirtualMachineTemplateArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineTemplateResult, error) {
	var rv LookupVirtualMachineTemplateResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20220110preview:getVirtualMachineTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineTemplateArgs struct {
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	VirtualMachineTemplateName string `pulumi:"virtualMachineTemplateName"`
}


type LookupVirtualMachineTemplateResult struct {
	CustomResourceName string                     `pulumi:"customResourceName"`
	Disks              []VirtualDiskResponse      `pulumi:"disks"`
	ExtendedLocation   *ExtendedLocationResponse  `pulumi:"extendedLocation"`
	FirmwareType       string                     `pulumi:"firmwareType"`
	FolderPath         string                     `pulumi:"folderPath"`
	Id                 string                     `pulumi:"id"`
	InventoryItemId    *string                    `pulumi:"inventoryItemId"`
	Kind               *string                    `pulumi:"kind"`
	Location           string                     `pulumi:"location"`
	MemorySizeMB       int                        `pulumi:"memorySizeMB"`
	MoName             string                     `pulumi:"moName"`
	MoRefId            *string                    `pulumi:"moRefId"`
	Name               string                     `pulumi:"name"`
	NetworkInterfaces  []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	NumCPUs            int                        `pulumi:"numCPUs"`
	NumCoresPerSocket  int                        `pulumi:"numCoresPerSocket"`
	OsName             string                     `pulumi:"osName"`
	OsType             string                     `pulumi:"osType"`
	ProvisioningState  string                     `pulumi:"provisioningState"`
	Statuses           []ResourceStatusResponse   `pulumi:"statuses"`
	SystemData         SystemDataResponse         `pulumi:"systemData"`
	Tags               map[string]string          `pulumi:"tags"`
	ToolsVersion       string                     `pulumi:"toolsVersion"`
	ToolsVersionStatus string                     `pulumi:"toolsVersionStatus"`
	Type               string                     `pulumi:"type"`
	Uuid               string                     `pulumi:"uuid"`
	VCenterId          *string                    `pulumi:"vCenterId"`
}

func LookupVirtualMachineTemplateOutput(ctx *pulumi.Context, args LookupVirtualMachineTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineTemplateResult, error) {
			args := v.(LookupVirtualMachineTemplateArgs)
			r, err := LookupVirtualMachineTemplate(ctx, &args, opts...)
			var s LookupVirtualMachineTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineTemplateResultOutput)
}

type LookupVirtualMachineTemplateOutputArgs struct {
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualMachineTemplateName pulumi.StringInput `pulumi:"virtualMachineTemplateName"`
}

func (LookupVirtualMachineTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineTemplateArgs)(nil)).Elem()
}


type LookupVirtualMachineTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineTemplateResult)(nil)).Elem()
}

func (o LookupVirtualMachineTemplateResultOutput) ToLookupVirtualMachineTemplateResultOutput() LookupVirtualMachineTemplateResultOutput {
	return o
}

func (o LookupVirtualMachineTemplateResultOutput) ToLookupVirtualMachineTemplateResultOutputWithContext(ctx context.Context) LookupVirtualMachineTemplateResultOutput {
	return o
}

func (o LookupVirtualMachineTemplateResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) []VirtualDiskResponse { return v.Disks }).(VirtualDiskResponseArrayOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) FirmwareType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.FirmwareType }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.FolderPath }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) MemorySizeMB() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.MemorySizeMB }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.MoName }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) NumCPUs() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.NumCPUs }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) NumCoresPerSocket() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.NumCoresPerSocket }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.OsName }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.OsType }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) ToolsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.ToolsVersion }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) ToolsVersionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.ToolsVersionStatus }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineTemplateResultOutput{})
}
