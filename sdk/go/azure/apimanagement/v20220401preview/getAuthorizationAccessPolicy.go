


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAuthorizationAccessPolicy(ctx *pulumi.Context, args *LookupAuthorizationAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationAccessPolicyResult, error) {
	var rv LookupAuthorizationAccessPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getAuthorizationAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationAccessPolicyArgs struct {
	AuthorizationAccessPolicyId string `pulumi:"authorizationAccessPolicyId"`
	AuthorizationId             string `pulumi:"authorizationId"`
	AuthorizationProviderId     string `pulumi:"authorizationProviderId"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	ServiceName                 string `pulumi:"serviceName"`
}


type LookupAuthorizationAccessPolicyResult struct {
	Id       string  `pulumi:"id"`
	Name     string  `pulumi:"name"`
	ObjectId *string `pulumi:"objectId"`
	TenantId *string `pulumi:"tenantId"`
	Type     string  `pulumi:"type"`
}

func LookupAuthorizationAccessPolicyOutput(ctx *pulumi.Context, args LookupAuthorizationAccessPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationAccessPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationAccessPolicyResult, error) {
			args := v.(LookupAuthorizationAccessPolicyArgs)
			r, err := LookupAuthorizationAccessPolicy(ctx, &args, opts...)
			var s LookupAuthorizationAccessPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizationAccessPolicyResultOutput)
}

type LookupAuthorizationAccessPolicyOutputArgs struct {
	AuthorizationAccessPolicyId pulumi.StringInput `pulumi:"authorizationAccessPolicyId"`
	AuthorizationId             pulumi.StringInput `pulumi:"authorizationId"`
	AuthorizationProviderId     pulumi.StringInput `pulumi:"authorizationProviderId"`
	ResourceGroupName           pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName                 pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupAuthorizationAccessPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationAccessPolicyArgs)(nil)).Elem()
}


type LookupAuthorizationAccessPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationAccessPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationAccessPolicyResult)(nil)).Elem()
}

func (o LookupAuthorizationAccessPolicyResultOutput) ToLookupAuthorizationAccessPolicyResultOutput() LookupAuthorizationAccessPolicyResultOutput {
	return o
}

func (o LookupAuthorizationAccessPolicyResultOutput) ToLookupAuthorizationAccessPolicyResultOutputWithContext(ctx context.Context) LookupAuthorizationAccessPolicyResultOutput {
	return o
}

func (o LookupAuthorizationAccessPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationAccessPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAuthorizationAccessPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationAccessPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAuthorizationAccessPolicyResultOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationAccessPolicyResult) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationAccessPolicyResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationAccessPolicyResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationAccessPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationAccessPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationAccessPolicyResultOutput{})
}
