


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationvCenter(ctx *pulumi.Context, args *LookupReplicationvCenterArgs, opts ...pulumi.InvokeOption) (*LookupReplicationvCenterResult, error) {
	var rv LookupReplicationvCenterResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210801:getReplicationvCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationvCenterArgs struct {
	FabricName        string `pulumi:"fabricName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
	VcenterName       string `pulumi:"vcenterName"`
}


type LookupReplicationvCenterResult struct {
	Id         string                    `pulumi:"id"`
	Location   *string                   `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties VCenterPropertiesResponse `pulumi:"properties"`
	Type       string                    `pulumi:"type"`
}
