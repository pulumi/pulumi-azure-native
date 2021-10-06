


package v20200815preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineExtension(ctx *pulumi.Context, args *LookupMachineExtensionArgs, opts ...pulumi.InvokeOption) (*LookupMachineExtensionResult, error) {
	var rv LookupMachineExtensionResult
	err := ctx.Invoke("azure-native:hybridcompute/v20200815preview:getMachineExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineExtensionArgs struct {
	ExtensionName     string `pulumi:"extensionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMachineExtensionResult struct {
	AutoUpgradeMinorVersion *bool                                           `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          *string                                         `pulumi:"forceUpdateTag"`
	Id                      string                                          `pulumi:"id"`
	InstanceView            *MachineExtensionPropertiesResponseInstanceView `pulumi:"instanceView"`
	Location                string                                          `pulumi:"location"`
	Name                    string                                          `pulumi:"name"`
	ProtectedSettings       interface{}                                     `pulumi:"protectedSettings"`
	ProvisioningState       string                                          `pulumi:"provisioningState"`
	Publisher               *string                                         `pulumi:"publisher"`
	Settings                interface{}                                     `pulumi:"settings"`
	Tags                    map[string]string                               `pulumi:"tags"`
	Type                    string                                          `pulumi:"type"`
	TypeHandlerVersion      *string                                         `pulumi:"typeHandlerVersion"`
}
