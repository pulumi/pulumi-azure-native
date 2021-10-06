


package connectedvmwarevsphere

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestAgent(ctx *pulumi.Context, args *LookupGuestAgentArgs, opts ...pulumi.InvokeOption) (*LookupGuestAgentResult, error) {
	var rv LookupGuestAgentResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere:getGuestAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGuestAgentArgs struct {
	Name               string `pulumi:"name"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VirtualMachineName string `pulumi:"virtualMachineName"`
}


type LookupGuestAgentResult struct {
	Credentials        *GuestCredentialResponse        `pulumi:"credentials"`
	CustomResourceName string                          `pulumi:"customResourceName"`
	HttpProxyConfig    *HttpProxyConfigurationResponse `pulumi:"httpProxyConfig"`
	Id                 string                          `pulumi:"id"`
	Name               string                          `pulumi:"name"`
	ProvisioningAction *string                         `pulumi:"provisioningAction"`
	ProvisioningState  string                          `pulumi:"provisioningState"`
	Status             string                          `pulumi:"status"`
	Statuses           []ResourceStatusResponse        `pulumi:"statuses"`
	SystemData         SystemDataResponse              `pulumi:"systemData"`
	Type               string                          `pulumi:"type"`
	Uuid               string                          `pulumi:"uuid"`
}
