


package v20190802preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupMachine(ctx *pulumi.Context, args *LookupMachineArgs, opts ...pulumi.InvokeOption) (*LookupMachineResult, error) {
	var rv LookupMachineResult
	err := ctx.Invoke("azure-native:hybridcompute/v20190802preview:getMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineArgs struct {
	Expand            *string `pulumi:"expand"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupMachineResult struct {
	AgentVersion      string                                 `pulumi:"agentVersion"`
	ClientPublicKey   *string                                `pulumi:"clientPublicKey"`
	DisplayName       string                                 `pulumi:"displayName"`
	ErrorDetails      []ErrorDetailResponse                  `pulumi:"errorDetails"`
	Extensions        []MachineExtensionInstanceViewResponse `pulumi:"extensions"`
	Id                string                                 `pulumi:"id"`
	LastStatusChange  string                                 `pulumi:"lastStatusChange"`
	Location          string                                 `pulumi:"location"`
	MachineFqdn       string                                 `pulumi:"machineFqdn"`
	Name              string                                 `pulumi:"name"`
	OsName            *string                                `pulumi:"osName"`
	OsProfile         OSProfileResponse                      `pulumi:"osProfile"`
	OsVersion         *string                                `pulumi:"osVersion"`
	PhysicalLocation  *string                                `pulumi:"physicalLocation"`
	PrincipalId       string                                 `pulumi:"principalId"`
	ProvisioningState string                                 `pulumi:"provisioningState"`
	Status            string                                 `pulumi:"status"`
	Tags              map[string]string                      `pulumi:"tags"`
	TenantId          string                                 `pulumi:"tenantId"`
	Type              string                                 `pulumi:"type"`
	VmId              string                                 `pulumi:"vmId"`
}

func LookupMachineOutput(ctx *pulumi.Context, args LookupMachineOutputArgs, opts ...pulumi.InvokeOption) LookupMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineResult, error) {
			args := v.(LookupMachineArgs)
			r, err := LookupMachine(ctx, &args, opts...)
			var s LookupMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineResultOutput)
}

type LookupMachineOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineArgs)(nil)).Elem()
}


type LookupMachineResultOutput struct{ *pulumi.OutputState }

func (LookupMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineResult)(nil)).Elem()
}

func (o LookupMachineResultOutput) ToLookupMachineResultOutput() LookupMachineResultOutput {
	return o
}

func (o LookupMachineResultOutput) ToLookupMachineResultOutputWithContext(ctx context.Context) LookupMachineResultOutput {
	return o
}

func (o LookupMachineResultOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) ClientPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *string { return v.ClientPublicKey }).(pulumi.StringPtrOutput)
}

func (o LookupMachineResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v LookupMachineResult) []ErrorDetailResponse { return v.ErrorDetails }).(ErrorDetailResponseArrayOutput)
}

func (o LookupMachineResultOutput) Extensions() MachineExtensionInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v LookupMachineResult) []MachineExtensionInstanceViewResponse { return v.Extensions }).(MachineExtensionInstanceViewResponseArrayOutput)
}

func (o LookupMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) LastStatusChange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.LastStatusChange }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) MachineFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.MachineFqdn }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) OsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *string { return v.OsName }).(pulumi.StringPtrOutput)
}

func (o LookupMachineResultOutput) OsProfile() OSProfileResponseOutput {
	return o.ApplyT(func(v LookupMachineResult) OSProfileResponse { return v.OsProfile }).(OSProfileResponseOutput)
}

func (o LookupMachineResultOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *string { return v.OsVersion }).(pulumi.StringPtrOutput)
}

func (o LookupMachineResultOutput) PhysicalLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *string { return v.PhysicalLocation }).(pulumi.StringPtrOutput)
}

func (o LookupMachineResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMachineResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.VmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineResultOutput{})
}
