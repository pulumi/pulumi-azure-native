


package v20191212

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupMachine(ctx *pulumi.Context, args *LookupMachineArgs, opts ...pulumi.InvokeOption) (*LookupMachineResult, error) {
	var rv LookupMachineResult
	err := ctx.Invoke("azure-native:hybridcompute/v20191212:getMachine", args, &rv, opts...)
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
	Identity          *MachineResponseIdentity               `pulumi:"identity"`
	LastStatusChange  string                                 `pulumi:"lastStatusChange"`
	Location          string                                 `pulumi:"location"`
	LocationData      *LocationDataResponse                  `pulumi:"locationData"`
	MachineFqdn       string                                 `pulumi:"machineFqdn"`
	Name              string                                 `pulumi:"name"`
	OsName            string                                 `pulumi:"osName"`
	OsProfile         *MachinePropertiesResponseOsProfile    `pulumi:"osProfile"`
	OsVersion         string                                 `pulumi:"osVersion"`
	ProvisioningState string                                 `pulumi:"provisioningState"`
	Status            string                                 `pulumi:"status"`
	Tags              map[string]string                      `pulumi:"tags"`
	Type              string                                 `pulumi:"type"`
	VmId              *string                                `pulumi:"vmId"`
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

func (o LookupMachineResultOutput) Identity() MachineResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *MachineResponseIdentity { return v.Identity }).(MachineResponseIdentityPtrOutput)
}

func (o LookupMachineResultOutput) LastStatusChange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.LastStatusChange }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) LocationData() LocationDataResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *LocationDataResponse { return v.LocationData }).(LocationDataResponsePtrOutput)
}

func (o LookupMachineResultOutput) MachineFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.MachineFqdn }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.OsName }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) OsProfile() MachinePropertiesResponseOsProfilePtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *MachinePropertiesResponseOsProfile { return v.OsProfile }).(MachinePropertiesResponseOsProfilePtrOutput)
}

func (o LookupMachineResultOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.OsVersion }).(pulumi.StringOutput)
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

func (o LookupMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineResultOutput{})
}
