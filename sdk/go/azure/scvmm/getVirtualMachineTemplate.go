


package scvmm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineTemplate(ctx *pulumi.Context, args *LookupVirtualMachineTemplateArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineTemplateResult, error) {
	var rv LookupVirtualMachineTemplateResult
	err := ctx.Invoke("azure-native:scvmm:getVirtualMachineTemplate", args, &rv, opts...)
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
	ComputerName         string                      `pulumi:"computerName"`
	CpuCount             int                         `pulumi:"cpuCount"`
	Disks                []VirtualDiskResponse       `pulumi:"disks"`
	DynamicMemoryEnabled string                      `pulumi:"dynamicMemoryEnabled"`
	DynamicMemoryMaxMB   int                         `pulumi:"dynamicMemoryMaxMB"`
	DynamicMemoryMinMB   int                         `pulumi:"dynamicMemoryMinMB"`
	ExtendedLocation     ExtendedLocationResponse    `pulumi:"extendedLocation"`
	Generation           int                         `pulumi:"generation"`
	Id                   string                      `pulumi:"id"`
	InventoryItemId      *string                     `pulumi:"inventoryItemId"`
	IsCustomizable       string                      `pulumi:"isCustomizable"`
	IsHighlyAvailable    string                      `pulumi:"isHighlyAvailable"`
	LimitCpuForMigration string                      `pulumi:"limitCpuForMigration"`
	Location             string                      `pulumi:"location"`
	MemoryMB             int                         `pulumi:"memoryMB"`
	Name                 string                      `pulumi:"name"`
	NetworkInterfaces    []NetworkInterfacesResponse `pulumi:"networkInterfaces"`
	OsName               string                      `pulumi:"osName"`
	OsType               string                      `pulumi:"osType"`
	ProvisioningState    string                      `pulumi:"provisioningState"`
	SystemData           SystemDataResponse          `pulumi:"systemData"`
	Tags                 map[string]string           `pulumi:"tags"`
	Type                 string                      `pulumi:"type"`
	Uuid                 *string                     `pulumi:"uuid"`
	VmmServerId          *string                     `pulumi:"vmmServerId"`
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

func (o LookupVirtualMachineTemplateResultOutput) ComputerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.ComputerName }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) CpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.CpuCount }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) []VirtualDiskResponse { return v.Disks }).(VirtualDiskResponseArrayOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) DynamicMemoryEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.DynamicMemoryEnabled }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) DynamicMemoryMaxMB() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.DynamicMemoryMaxMB }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) DynamicMemoryMinMB() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.DynamicMemoryMinMB }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Generation() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.Generation }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) IsCustomizable() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.IsCustomizable }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) IsHighlyAvailable() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.IsHighlyAvailable }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) LimitCpuForMigration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.LimitCpuForMigration }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) MemoryMB() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.MemoryMB }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) NetworkInterfaces() NetworkInterfacesResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) []NetworkInterfacesResponse { return v.NetworkInterfaces }).(NetworkInterfacesResponseArrayOutput)
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

func (o LookupVirtualMachineTemplateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineTemplateResultOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineTemplateResultOutput{})
}
