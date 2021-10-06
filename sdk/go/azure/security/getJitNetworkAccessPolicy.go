


package security

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJitNetworkAccessPolicy(ctx *pulumi.Context, args *LookupJitNetworkAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupJitNetworkAccessPolicyResult, error) {
	var rv LookupJitNetworkAccessPolicyResult
	err := ctx.Invoke("azure-native:security:getJitNetworkAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJitNetworkAccessPolicyArgs struct {
	AscLocation                string `pulumi:"ascLocation"`
	JitNetworkAccessPolicyName string `pulumi:"jitNetworkAccessPolicyName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}

type LookupJitNetworkAccessPolicyResult struct {
	Id                string                                         `pulumi:"id"`
	Kind              *string                                        `pulumi:"kind"`
	Location          string                                         `pulumi:"location"`
	Name              string                                         `pulumi:"name"`
	ProvisioningState string                                         `pulumi:"provisioningState"`
	Requests          []JitNetworkAccessRequestResponse              `pulumi:"requests"`
	Type              string                                         `pulumi:"type"`
	VirtualMachines   []JitNetworkAccessPolicyVirtualMachineResponse `pulumi:"virtualMachines"`
}
