


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupJitRequest(ctx *pulumi.Context, args *LookupJitRequestArgs, opts ...pulumi.InvokeOption) (*LookupJitRequestResult, error) {
	var rv LookupJitRequestResult
	err := ctx.Invoke("azure-native:solutions/v20180901preview:getJitRequest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJitRequestArgs struct {
	JitRequestName    string `pulumi:"jitRequestName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupJitRequestResult struct {
	ApplicationResourceId    string                             `pulumi:"applicationResourceId"`
	CreatedBy                ApplicationClientDetailsResponse   `pulumi:"createdBy"`
	Id                       string                             `pulumi:"id"`
	JitAuthorizationPolicies []JitAuthorizationPoliciesResponse `pulumi:"jitAuthorizationPolicies"`
	JitRequestState          string                             `pulumi:"jitRequestState"`
	JitSchedulingPolicy      JitSchedulingPolicyResponse        `pulumi:"jitSchedulingPolicy"`
	Location                 *string                            `pulumi:"location"`
	Name                     string                             `pulumi:"name"`
	ProvisioningState        string                             `pulumi:"provisioningState"`
	PublisherTenantId        string                             `pulumi:"publisherTenantId"`
	Tags                     map[string]string                  `pulumi:"tags"`
	Type                     string                             `pulumi:"type"`
	UpdatedBy                ApplicationClientDetailsResponse   `pulumi:"updatedBy"`
}

func LookupJitRequestOutput(ctx *pulumi.Context, args LookupJitRequestOutputArgs, opts ...pulumi.InvokeOption) LookupJitRequestResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJitRequestResult, error) {
			args := v.(LookupJitRequestArgs)
			r, err := LookupJitRequest(ctx, &args, opts...)
			var s LookupJitRequestResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJitRequestResultOutput)
}

type LookupJitRequestOutputArgs struct {
	JitRequestName    pulumi.StringInput `pulumi:"jitRequestName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJitRequestOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJitRequestArgs)(nil)).Elem()
}


type LookupJitRequestResultOutput struct{ *pulumi.OutputState }

func (LookupJitRequestResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJitRequestResult)(nil)).Elem()
}

func (o LookupJitRequestResultOutput) ToLookupJitRequestResultOutput() LookupJitRequestResultOutput {
	return o
}

func (o LookupJitRequestResultOutput) ToLookupJitRequestResultOutputWithContext(ctx context.Context) LookupJitRequestResultOutput {
	return o
}

func (o LookupJitRequestResultOutput) ApplicationResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitRequestResult) string { return v.ApplicationResourceId }).(pulumi.StringOutput)
}

func (o LookupJitRequestResultOutput) CreatedBy() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v LookupJitRequestResult) ApplicationClientDetailsResponse { return v.CreatedBy }).(ApplicationClientDetailsResponseOutput)
}

func (o LookupJitRequestResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitRequestResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJitRequestResultOutput) JitAuthorizationPolicies() JitAuthorizationPoliciesResponseArrayOutput {
	return o.ApplyT(func(v LookupJitRequestResult) []JitAuthorizationPoliciesResponse { return v.JitAuthorizationPolicies }).(JitAuthorizationPoliciesResponseArrayOutput)
}

func (o LookupJitRequestResultOutput) JitRequestState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitRequestResult) string { return v.JitRequestState }).(pulumi.StringOutput)
}

func (o LookupJitRequestResultOutput) JitSchedulingPolicy() JitSchedulingPolicyResponseOutput {
	return o.ApplyT(func(v LookupJitRequestResult) JitSchedulingPolicyResponse { return v.JitSchedulingPolicy }).(JitSchedulingPolicyResponseOutput)
}

func (o LookupJitRequestResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJitRequestResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupJitRequestResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitRequestResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJitRequestResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitRequestResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupJitRequestResultOutput) PublisherTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitRequestResult) string { return v.PublisherTenantId }).(pulumi.StringOutput)
}

func (o LookupJitRequestResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJitRequestResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupJitRequestResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitRequestResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupJitRequestResultOutput) UpdatedBy() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v LookupJitRequestResult) ApplicationClientDetailsResponse { return v.UpdatedBy }).(ApplicationClientDetailsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJitRequestResultOutput{})
}
