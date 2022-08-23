


package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSAPApplicationServerInstance(ctx *pulumi.Context, args *LookupSAPApplicationServerInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSAPApplicationServerInstanceResult, error) {
	var rv LookupSAPApplicationServerInstanceResult
	err := ctx.Invoke("azure-native:workloads:getSAPApplicationServerInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSAPApplicationServerInstanceArgs struct {
	ApplicationInstanceName string `pulumi:"applicationInstanceName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SapVirtualInstanceName  string `pulumi:"sapVirtualInstanceName"`
}


type LookupSAPApplicationServerInstanceResult struct {
	Errors            SAPVirtualInstanceErrorResponse `pulumi:"errors"`
	GatewayPort       float64                         `pulumi:"gatewayPort"`
	Health            string                          `pulumi:"health"`
	Hostname          string                          `pulumi:"hostname"`
	IcmHttpPort       float64                         `pulumi:"icmHttpPort"`
	IcmHttpsPort      float64                         `pulumi:"icmHttpsPort"`
	Id                string                          `pulumi:"id"`
	InstanceNo        string                          `pulumi:"instanceNo"`
	IpAddress         string                          `pulumi:"ipAddress"`
	KernelPatch       string                          `pulumi:"kernelPatch"`
	KernelVersion     string                          `pulumi:"kernelVersion"`
	Location          string                          `pulumi:"location"`
	Name              string                          `pulumi:"name"`
	ProvisioningState string                          `pulumi:"provisioningState"`
	Status            string                          `pulumi:"status"`
	Subnet            string                          `pulumi:"subnet"`
	SystemData        SystemDataResponse              `pulumi:"systemData"`
	Tags              map[string]string               `pulumi:"tags"`
	Type              string                          `pulumi:"type"`
	VirtualMachineId  string                          `pulumi:"virtualMachineId"`
}

func LookupSAPApplicationServerInstanceOutput(ctx *pulumi.Context, args LookupSAPApplicationServerInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupSAPApplicationServerInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSAPApplicationServerInstanceResult, error) {
			args := v.(LookupSAPApplicationServerInstanceArgs)
			r, err := LookupSAPApplicationServerInstance(ctx, &args, opts...)
			var s LookupSAPApplicationServerInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSAPApplicationServerInstanceResultOutput)
}

type LookupSAPApplicationServerInstanceOutputArgs struct {
	ApplicationInstanceName pulumi.StringInput `pulumi:"applicationInstanceName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	SapVirtualInstanceName  pulumi.StringInput `pulumi:"sapVirtualInstanceName"`
}

func (LookupSAPApplicationServerInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSAPApplicationServerInstanceArgs)(nil)).Elem()
}


type LookupSAPApplicationServerInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupSAPApplicationServerInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSAPApplicationServerInstanceResult)(nil)).Elem()
}

func (o LookupSAPApplicationServerInstanceResultOutput) ToLookupSAPApplicationServerInstanceResultOutput() LookupSAPApplicationServerInstanceResultOutput {
	return o
}

func (o LookupSAPApplicationServerInstanceResultOutput) ToLookupSAPApplicationServerInstanceResultOutputWithContext(ctx context.Context) LookupSAPApplicationServerInstanceResultOutput {
	return o
}

func (o LookupSAPApplicationServerInstanceResultOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) SAPVirtualInstanceErrorResponse { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) GatewayPort() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) float64 { return v.GatewayPort }).(pulumi.Float64Output)
}

func (o LookupSAPApplicationServerInstanceResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Health }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) IcmHttpPort() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) float64 { return v.IcmHttpPort }).(pulumi.Float64Output)
}

func (o LookupSAPApplicationServerInstanceResultOutput) IcmHttpsPort() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) float64 { return v.IcmHttpsPort }).(pulumi.Float64Output)
}

func (o LookupSAPApplicationServerInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) InstanceNo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.InstanceNo }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) KernelPatch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.KernelPatch }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) KernelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.KernelVersion }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Subnet }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSAPApplicationServerInstanceResultOutput) VirtualMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.VirtualMachineId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSAPApplicationServerInstanceResultOutput{})
}
