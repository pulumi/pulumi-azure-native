


package v20170131

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainerService(ctx *pulumi.Context, args *LookupContainerServiceArgs, opts ...pulumi.InvokeOption) (*LookupContainerServiceResult, error) {
	var rv LookupContainerServiceResult
	err := ctx.Invoke("azure-native:containerservice/v20170131:getContainerService", args, &rv, opts...)
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
	AgentPoolProfiles       []ContainerServiceAgentPoolProfileResponse       `pulumi:"agentPoolProfiles"`
	CustomProfile           *ContainerServiceCustomProfileResponse           `pulumi:"customProfile"`
	DiagnosticsProfile      *ContainerServiceDiagnosticsProfileResponse      `pulumi:"diagnosticsProfile"`
	Id                      string                                           `pulumi:"id"`
	LinuxProfile            ContainerServiceLinuxProfileResponse             `pulumi:"linuxProfile"`
	Location                string                                           `pulumi:"location"`
	MasterProfile           ContainerServiceMasterProfileResponse            `pulumi:"masterProfile"`
	Name                    string                                           `pulumi:"name"`
	OrchestratorProfile     *ContainerServiceOrchestratorProfileResponse     `pulumi:"orchestratorProfile"`
	ProvisioningState       string                                           `pulumi:"provisioningState"`
	ServicePrincipalProfile *ContainerServiceServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	Tags                    map[string]string                                `pulumi:"tags"`
	Type                    string                                           `pulumi:"type"`
	WindowsProfile          *ContainerServiceWindowsProfileResponse          `pulumi:"windowsProfile"`
}


func (val *LookupContainerServiceResult) Defaults() *LookupContainerServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.MasterProfile = *tmp.MasterProfile.Defaults()

	return &tmp
}
