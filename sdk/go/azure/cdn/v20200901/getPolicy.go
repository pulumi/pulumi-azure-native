


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	var rv LookupPolicyResult
	err := ctx.Invoke("azure-native:cdn/v20200901:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyArgs struct {
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPolicyResult struct {
	CustomRules       *CustomRuleListResponse     `pulumi:"customRules"`
	EndpointLinks     []CdnEndpointResponse       `pulumi:"endpointLinks"`
	Etag              *string                     `pulumi:"etag"`
	Id                string                      `pulumi:"id"`
	Location          string                      `pulumi:"location"`
	ManagedRules      *ManagedRuleSetListResponse `pulumi:"managedRules"`
	Name              string                      `pulumi:"name"`
	PolicySettings    *PolicySettingsResponse     `pulumi:"policySettings"`
	ProvisioningState string                      `pulumi:"provisioningState"`
	RateLimitRules    *RateLimitRuleListResponse  `pulumi:"rateLimitRules"`
	ResourceState     string                      `pulumi:"resourceState"`
	Sku               SkuResponse                 `pulumi:"sku"`
	SystemData        SystemDataResponse          `pulumi:"systemData"`
	Tags              map[string]string           `pulumi:"tags"`
	Type              string                      `pulumi:"type"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

type LookupPolicyOutputArgs struct {
	PolicyName        pulumi.StringInput `pulumi:"policyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}


type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) CustomRules() CustomRuleListResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *CustomRuleListResponse { return v.CustomRules }).(CustomRuleListResponsePtrOutput)
}

func (o LookupPolicyResultOutput) EndpointLinks() CdnEndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []CdnEndpointResponse { return v.EndpointLinks }).(CdnEndpointResponseArrayOutput)
}

func (o LookupPolicyResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) ManagedRules() ManagedRuleSetListResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *ManagedRuleSetListResponse { return v.ManagedRules }).(ManagedRuleSetListResponsePtrOutput)
}

func (o LookupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) PolicySettings() PolicySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *PolicySettingsResponse { return v.PolicySettings }).(PolicySettingsResponsePtrOutput)
}

func (o LookupPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) RateLimitRules() RateLimitRuleListResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *RateLimitRuleListResponse { return v.RateLimitRules }).(RateLimitRuleListResponsePtrOutput)
}

func (o LookupPolicyResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupPolicyResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
