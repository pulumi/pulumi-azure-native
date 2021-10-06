


package v20191212

import (
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
