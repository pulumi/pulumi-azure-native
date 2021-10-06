


package v20210325preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineExtension(ctx *pulumi.Context, args *LookupMachineExtensionArgs, opts ...pulumi.InvokeOption) (*LookupMachineExtensionResult, error) {
	var rv LookupMachineExtensionResult
	err := ctx.Invoke("azure-native:hybridcompute/v20210325preview:getMachineExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineExtensionArgs struct {
	ExtensionName     string `pulumi:"extensionName"`
	MachineName       string `pulumi:"machineName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMachineExtensionResult struct {
	Id         string                             `pulumi:"id"`
	Location   string                             `pulumi:"location"`
	Name       string                             `pulumi:"name"`
	Properties MachineExtensionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                 `pulumi:"systemData"`
	Tags       map[string]string                  `pulumi:"tags"`
	Type       string                             `pulumi:"type"`
}
