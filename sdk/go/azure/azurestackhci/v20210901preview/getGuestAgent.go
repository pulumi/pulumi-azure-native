


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestAgent(ctx *pulumi.Context, args *LookupGuestAgentArgs, opts ...pulumi.InvokeOption) (*LookupGuestAgentResult, error) {
	var rv LookupGuestAgentResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210901preview:getGuestAgent", args, &rv, opts...)
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
	HttpProxyConfig    *HttpProxyConfigurationResponse `pulumi:"httpProxyConfig"`
	Id                 string                          `pulumi:"id"`
	Name               string                          `pulumi:"name"`
	ProvisioningAction *string                         `pulumi:"provisioningAction"`
	ProvisioningState  string                          `pulumi:"provisioningState"`
	Status             string                          `pulumi:"status"`
	SystemData         SystemDataResponse              `pulumi:"systemData"`
	Type               string                          `pulumi:"type"`
}

func LookupGuestAgentOutput(ctx *pulumi.Context, args LookupGuestAgentOutputArgs, opts ...pulumi.InvokeOption) LookupGuestAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGuestAgentResult, error) {
			args := v.(LookupGuestAgentArgs)
			r, err := LookupGuestAgent(ctx, &args, opts...)
			var s LookupGuestAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGuestAgentResultOutput)
}

type LookupGuestAgentOutputArgs struct {
	Name               pulumi.StringInput `pulumi:"name"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualMachineName pulumi.StringInput `pulumi:"virtualMachineName"`
}

func (LookupGuestAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestAgentArgs)(nil)).Elem()
}


type LookupGuestAgentResultOutput struct{ *pulumi.OutputState }

func (LookupGuestAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestAgentResult)(nil)).Elem()
}

func (o LookupGuestAgentResultOutput) ToLookupGuestAgentResultOutput() LookupGuestAgentResultOutput {
	return o
}

func (o LookupGuestAgentResultOutput) ToLookupGuestAgentResultOutputWithContext(ctx context.Context) LookupGuestAgentResultOutput {
	return o
}

func (o LookupGuestAgentResultOutput) Credentials() GuestCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) *GuestCredentialResponse { return v.Credentials }).(GuestCredentialResponsePtrOutput)
}

func (o LookupGuestAgentResultOutput) HttpProxyConfig() HttpProxyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) *HttpProxyConfigurationResponse { return v.HttpProxyConfig }).(HttpProxyConfigurationResponsePtrOutput)
}

func (o LookupGuestAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGuestAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGuestAgentResultOutput) ProvisioningAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) *string { return v.ProvisioningAction }).(pulumi.StringPtrOutput)
}

func (o LookupGuestAgentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupGuestAgentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupGuestAgentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGuestAgentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGuestAgentResultOutput{})
}
