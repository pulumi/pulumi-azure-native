


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKubernetesRole(ctx *pulumi.Context, args *LookupKubernetesRoleArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesRoleResult, error) {
	var rv LookupKubernetesRoleResult
	err := ctx.Invoke("azure-native:databoxedge/v20200901:getKubernetesRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKubernetesRoleArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupKubernetesRoleResult struct {
	HostPlatform            string                          `pulumi:"hostPlatform"`
	HostPlatformType        string                          `pulumi:"hostPlatformType"`
	Id                      string                          `pulumi:"id"`
	Kind                    string                          `pulumi:"kind"`
	KubernetesClusterInfo   KubernetesClusterInfoResponse   `pulumi:"kubernetesClusterInfo"`
	KubernetesRoleResources KubernetesRoleResourcesResponse `pulumi:"kubernetesRoleResources"`
	Name                    string                          `pulumi:"name"`
	ProvisioningState       string                          `pulumi:"provisioningState"`
	RoleStatus              string                          `pulumi:"roleStatus"`
	SystemData              SystemDataResponse              `pulumi:"systemData"`
	Type                    string                          `pulumi:"type"`
}
