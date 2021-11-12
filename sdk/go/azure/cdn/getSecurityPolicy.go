


package cdn

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityPolicy(ctx *pulumi.Context, args *LookupSecurityPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSecurityPolicyResult, error) {
	var rv LookupSecurityPolicyResult
	err := ctx.Invoke("azure-native:cdn:getSecurityPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityPolicyArgs struct {
	ProfileName        string `pulumi:"profileName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	SecurityPolicyName string `pulumi:"securityPolicyName"`
}


type LookupSecurityPolicyResult struct {
	DeploymentStatus  string                                                  `pulumi:"deploymentStatus"`
	Id                string                                                  `pulumi:"id"`
	Name              string                                                  `pulumi:"name"`
	Parameters        *SecurityPolicyWebApplicationFirewallParametersResponse `pulumi:"parameters"`
	ProvisioningState string                                                  `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                                      `pulumi:"systemData"`
	Type              string                                                  `pulumi:"type"`
}
