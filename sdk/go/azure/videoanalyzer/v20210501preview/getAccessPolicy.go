


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessPolicy(ctx *pulumi.Context, args *LookupAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccessPolicyResult, error) {
	var rv LookupAccessPolicyResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20210501preview:getAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPolicyArgs struct {
	AccessPolicyName  string `pulumi:"accessPolicyName"`
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccessPolicyResult struct {
	Authentication *JwtAuthenticationResponse `pulumi:"authentication"`
	Id             string                     `pulumi:"id"`
	Name           string                     `pulumi:"name"`
	Role           *string                    `pulumi:"role"`
	SystemData     SystemDataResponse         `pulumi:"systemData"`
	Type           string                     `pulumi:"type"`
}

func LookupAccessPolicyOutput(ctx *pulumi.Context, args LookupAccessPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAccessPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessPolicyResult, error) {
			args := v.(LookupAccessPolicyArgs)
			r, err := LookupAccessPolicy(ctx, &args, opts...)
			var s LookupAccessPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessPolicyResultOutput)
}

type LookupAccessPolicyOutputArgs struct {
	AccessPolicyName  pulumi.StringInput `pulumi:"accessPolicyName"`
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccessPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPolicyArgs)(nil)).Elem()
}


type LookupAccessPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPolicyResult)(nil)).Elem()
}

func (o LookupAccessPolicyResultOutput) ToLookupAccessPolicyResultOutput() LookupAccessPolicyResultOutput {
	return o
}

func (o LookupAccessPolicyResultOutput) ToLookupAccessPolicyResultOutputWithContext(ctx context.Context) LookupAccessPolicyResultOutput {
	return o
}

func (o LookupAccessPolicyResultOutput) Authentication() JwtAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) *JwtAuthenticationResponse { return v.Authentication }).(JwtAuthenticationResponsePtrOutput)
}

func (o LookupAccessPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccessPolicyResultOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o LookupAccessPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAccessPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPolicyResultOutput{})
}
