


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityPolicy(ctx *pulumi.Context, args *LookupSecurityPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSecurityPolicyResult, error) {
	var rv LookupSecurityPolicyResult
	err := ctx.Invoke("azure-native:cdn/v20210601:getSecurityPolicy", args, &rv, opts...)
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
	ProfileName       string                                                  `pulumi:"profileName"`
	ProvisioningState string                                                  `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                                      `pulumi:"systemData"`
	Type              string                                                  `pulumi:"type"`
}

func LookupSecurityPolicyOutput(ctx *pulumi.Context, args LookupSecurityPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityPolicyResult, error) {
			args := v.(LookupSecurityPolicyArgs)
			r, err := LookupSecurityPolicy(ctx, &args, opts...)
			var s LookupSecurityPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityPolicyResultOutput)
}

type LookupSecurityPolicyOutputArgs struct {
	ProfileName        pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	SecurityPolicyName pulumi.StringInput `pulumi:"securityPolicyName"`
}

func (LookupSecurityPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityPolicyArgs)(nil)).Elem()
}


type LookupSecurityPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityPolicyResult)(nil)).Elem()
}

func (o LookupSecurityPolicyResultOutput) ToLookupSecurityPolicyResultOutput() LookupSecurityPolicyResultOutput {
	return o
}

func (o LookupSecurityPolicyResultOutput) ToLookupSecurityPolicyResultOutputWithContext(ctx context.Context) LookupSecurityPolicyResultOutput {
	return o
}

func (o LookupSecurityPolicyResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPolicyResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o LookupSecurityPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecurityPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecurityPolicyResultOutput) Parameters() SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupSecurityPolicyResult) *SecurityPolicyWebApplicationFirewallParametersResponse {
		return v.Parameters
	}).(SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput)
}

func (o LookupSecurityPolicyResultOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPolicyResult) string { return v.ProfileName }).(pulumi.StringOutput)
}

func (o LookupSecurityPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSecurityPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecurityPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSecurityPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityPolicyResultOutput{})
}
