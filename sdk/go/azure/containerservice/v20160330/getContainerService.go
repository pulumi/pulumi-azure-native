


package v20160330

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupContainerService(ctx *pulumi.Context, args *LookupContainerServiceArgs, opts ...pulumi.InvokeOption) (*LookupContainerServiceResult, error) {
	var rv LookupContainerServiceResult
	err := ctx.Invoke("azure-native:containerservice/v20160330:getContainerService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupContainerServiceArgs struct {
	ContainerServiceName string `pulumi:"containerServiceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupContainerServiceResult struct {
	AgentPoolProfiles   []ContainerServiceAgentPoolProfileResponse   `pulumi:"agentPoolProfiles"`
	DiagnosticsProfile  *ContainerServiceDiagnosticsProfileResponse  `pulumi:"diagnosticsProfile"`
	Id                  string                                       `pulumi:"id"`
	LinuxProfile        ContainerServiceLinuxProfileResponse         `pulumi:"linuxProfile"`
	Location            string                                       `pulumi:"location"`
	MasterProfile       ContainerServiceMasterProfileResponse        `pulumi:"masterProfile"`
	Name                string                                       `pulumi:"name"`
	OrchestratorProfile *ContainerServiceOrchestratorProfileResponse `pulumi:"orchestratorProfile"`
	ProvisioningState   string                                       `pulumi:"provisioningState"`
	Tags                map[string]string                            `pulumi:"tags"`
	Type                string                                       `pulumi:"type"`
	WindowsProfile      *ContainerServiceWindowsProfileResponse      `pulumi:"windowsProfile"`
}


func (val *LookupContainerServiceResult) Defaults() *LookupContainerServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.MasterProfile = *tmp.MasterProfile.Defaults()

	return &tmp
}

func LookupContainerServiceOutput(ctx *pulumi.Context, args LookupContainerServiceOutputArgs, opts ...pulumi.InvokeOption) LookupContainerServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerServiceResult, error) {
			args := v.(LookupContainerServiceArgs)
			r, err := LookupContainerService(ctx, &args, opts...)
			var s LookupContainerServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerServiceResultOutput)
}

type LookupContainerServiceOutputArgs struct {
	ContainerServiceName pulumi.StringInput `pulumi:"containerServiceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContainerServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerServiceArgs)(nil)).Elem()
}


type LookupContainerServiceResultOutput struct{ *pulumi.OutputState }

func (LookupContainerServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerServiceResult)(nil)).Elem()
}

func (o LookupContainerServiceResultOutput) ToLookupContainerServiceResultOutput() LookupContainerServiceResultOutput {
	return o
}

func (o LookupContainerServiceResultOutput) ToLookupContainerServiceResultOutputWithContext(ctx context.Context) LookupContainerServiceResultOutput {
	return o
}

func (o LookupContainerServiceResultOutput) AgentPoolProfiles() ContainerServiceAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) []ContainerServiceAgentPoolProfileResponse {
		return v.AgentPoolProfiles
	}).(ContainerServiceAgentPoolProfileResponseArrayOutput)
}

func (o LookupContainerServiceResultOutput) DiagnosticsProfile() ContainerServiceDiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) *ContainerServiceDiagnosticsProfileResponse {
		return v.DiagnosticsProfile
	}).(ContainerServiceDiagnosticsProfileResponsePtrOutput)
}

func (o LookupContainerServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContainerServiceResultOutput) LinuxProfile() ContainerServiceLinuxProfileResponseOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) ContainerServiceLinuxProfileResponse { return v.LinuxProfile }).(ContainerServiceLinuxProfileResponseOutput)
}

func (o LookupContainerServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupContainerServiceResultOutput) MasterProfile() ContainerServiceMasterProfileResponseOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) ContainerServiceMasterProfileResponse { return v.MasterProfile }).(ContainerServiceMasterProfileResponseOutput)
}

func (o LookupContainerServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContainerServiceResultOutput) OrchestratorProfile() ContainerServiceOrchestratorProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) *ContainerServiceOrchestratorProfileResponse {
		return v.OrchestratorProfile
	}).(ContainerServiceOrchestratorProfileResponsePtrOutput)
}

func (o LookupContainerServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupContainerServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupContainerServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupContainerServiceResultOutput) WindowsProfile() ContainerServiceWindowsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerServiceResult) *ContainerServiceWindowsProfileResponse { return v.WindowsProfile }).(ContainerServiceWindowsProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerServiceResultOutput{})
}
