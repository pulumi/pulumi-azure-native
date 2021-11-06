


package v20210422preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachine(ctx *pulumi.Context, args *LookupMachineArgs, opts ...pulumi.InvokeOption) (*LookupMachineResult, error) {
	var rv LookupMachineResult
	err := ctx.Invoke("azure-native:hybridcompute/v20210422preview:getMachine", args, &rv, opts...)
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
	Id         string                    `pulumi:"id"`
	Identity   *IdentityResponse         `pulumi:"identity"`
	Location   string                    `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties MachinePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse        `pulumi:"systemData"`
	Tags       map[string]string         `pulumi:"tags"`
	Type       string                    `pulumi:"type"`
}
