


package v20221110

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachine(ctx *pulumi.Context, args *LookupMachineArgs, opts ...pulumi.InvokeOption) (*LookupMachineResult, error) {
	var rv LookupMachineResult
	err := ctx.Invoke("azure-native:hybridcompute/v20221110:getMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineArgs struct {
	Expand            *string `pulumi:"expand"`
	MachineName       string  `pulumi:"machineName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupMachineResult struct {
	AdFqdn                     string                                 `pulumi:"adFqdn"`
	AgentConfiguration         AgentConfigurationResponse             `pulumi:"agentConfiguration"`
	AgentVersion               string                                 `pulumi:"agentVersion"`
	ClientPublicKey            *string                                `pulumi:"clientPublicKey"`
	CloudMetadata              *CloudMetadataResponse                 `pulumi:"cloudMetadata"`
	DetectedProperties         map[string]string                      `pulumi:"detectedProperties"`
	DisplayName                string                                 `pulumi:"displayName"`
	DnsFqdn                    string                                 `pulumi:"dnsFqdn"`
	DomainName                 string                                 `pulumi:"domainName"`
	ErrorDetails               []ErrorDetailResponse                  `pulumi:"errorDetails"`
	Extensions                 []MachineExtensionInstanceViewResponse `pulumi:"extensions"`
	Id                         string                                 `pulumi:"id"`
	Identity                   *IdentityResponse                      `pulumi:"identity"`
	LastStatusChange           string                                 `pulumi:"lastStatusChange"`
	Location                   string                                 `pulumi:"location"`
	LocationData               *LocationDataResponse                  `pulumi:"locationData"`
	MachineFqdn                string                                 `pulumi:"machineFqdn"`
	MssqlDiscovered            *string                                `pulumi:"mssqlDiscovered"`
	Name                       string                                 `pulumi:"name"`
	OsName                     string                                 `pulumi:"osName"`
	OsProfile                  *OSProfileResponse                     `pulumi:"osProfile"`
	OsSku                      string                                 `pulumi:"osSku"`
	OsType                     *string                                `pulumi:"osType"`
	OsVersion                  string                                 `pulumi:"osVersion"`
	ParentClusterResourceId    *string                                `pulumi:"parentClusterResourceId"`
	PrivateLinkScopeResourceId *string                                `pulumi:"privateLinkScopeResourceId"`
	ProvisioningState          string                                 `pulumi:"provisioningState"`
	Resources                  []MachineExtensionResponse             `pulumi:"resources"`
	ServiceStatuses            *ServiceStatusesResponse               `pulumi:"serviceStatuses"`
	Status                     string                                 `pulumi:"status"`
	SystemData                 SystemDataResponse                     `pulumi:"systemData"`
	Tags                       map[string]string                      `pulumi:"tags"`
	Type                       string                                 `pulumi:"type"`
	VmId                       *string                                `pulumi:"vmId"`
	VmUuid                     string                                 `pulumi:"vmUuid"`
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
	MachineName       pulumi.StringInput    `pulumi:"machineName"`
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

func (o LookupMachineResultOutput) AdFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.AdFqdn }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) AgentConfiguration() AgentConfigurationResponseOutput {
	return o.ApplyT(func(v LookupMachineResult) AgentConfigurationResponse { return v.AgentConfiguration }).(AgentConfigurationResponseOutput)
}

func (o LookupMachineResultOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) ClientPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *string { return v.ClientPublicKey }).(pulumi.StringPtrOutput)
}

func (o LookupMachineResultOutput) CloudMetadata() CloudMetadataResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *CloudMetadataResponse { return v.CloudMetadata }).(CloudMetadataResponsePtrOutput)
}

func (o LookupMachineResultOutput) DetectedProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMachineResult) map[string]string { return v.DetectedProperties }).(pulumi.StringMapOutput)
}

func (o LookupMachineResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) DnsFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.DnsFqdn }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.DomainName }).(pulumi.StringOutput)
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

func (o LookupMachineResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
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

func (o LookupMachineResultOutput) MssqlDiscovered() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *string { return v.MssqlDiscovered }).(pulumi.StringPtrOutput)
}

func (o LookupMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.OsName }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *OSProfileResponse { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

func (o LookupMachineResultOutput) OsSku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.OsSku }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupMachineResultOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.OsVersion }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) ParentClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *string { return v.ParentClusterResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupMachineResultOutput) PrivateLinkScopeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *string { return v.PrivateLinkScopeResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Resources() MachineExtensionResponseArrayOutput {
	return o.ApplyT(func(v LookupMachineResult) []MachineExtensionResponse { return v.Resources }).(MachineExtensionResponseArrayOutput)
}

func (o LookupMachineResultOutput) ServiceStatuses() ServiceStatusesResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *ServiceStatusesResponse { return v.ServiceStatuses }).(ServiceStatusesResponsePtrOutput)
}

func (o LookupMachineResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMachineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
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

func (o LookupMachineResultOutput) VmUuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.VmUuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineResultOutput{})
}
