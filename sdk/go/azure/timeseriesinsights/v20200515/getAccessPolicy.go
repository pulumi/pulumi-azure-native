


package v20200515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessPolicy(ctx *pulumi.Context, args *LookupAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccessPolicyResult, error) {
	var rv LookupAccessPolicyResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20200515:getAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPolicyArgs struct {
	AccessPolicyName  string `pulumi:"accessPolicyName"`
	EnvironmentName   string `pulumi:"environmentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccessPolicyResult struct {
	Description       *string  `pulumi:"description"`
	Id                string   `pulumi:"id"`
	Name              string   `pulumi:"name"`
	PrincipalObjectId *string  `pulumi:"principalObjectId"`
	Roles             []string `pulumi:"roles"`
	Type              string   `pulumi:"type"`
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
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
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

func (o LookupAccessPolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAccessPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccessPolicyResultOutput) PrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) *string { return v.PrincipalObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupAccessPolicyResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o LookupAccessPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPolicyResultOutput{})
}
