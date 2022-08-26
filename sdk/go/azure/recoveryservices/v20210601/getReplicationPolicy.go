


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationPolicy(ctx *pulumi.Context, args *LookupReplicationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupReplicationPolicyResult, error) {
	var rv LookupReplicationPolicyResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210601:getReplicationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationPolicyArgs struct {
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupReplicationPolicyResult struct {
	Id         string                   `pulumi:"id"`
	Location   *string                  `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Properties PolicyPropertiesResponse `pulumi:"properties"`
	Type       string                   `pulumi:"type"`
}

func LookupReplicationPolicyOutput(ctx *pulumi.Context, args LookupReplicationPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationPolicyResult, error) {
			args := v.(LookupReplicationPolicyArgs)
			r, err := LookupReplicationPolicy(ctx, &args, opts...)
			var s LookupReplicationPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationPolicyResultOutput)
}

type LookupReplicationPolicyOutputArgs struct {
	PolicyName        pulumi.StringInput `pulumi:"policyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationPolicyArgs)(nil)).Elem()
}


type LookupReplicationPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationPolicyResult)(nil)).Elem()
}

func (o LookupReplicationPolicyResultOutput) ToLookupReplicationPolicyResultOutput() LookupReplicationPolicyResultOutput {
	return o
}

func (o LookupReplicationPolicyResultOutput) ToLookupReplicationPolicyResultOutputWithContext(ctx context.Context) LookupReplicationPolicyResultOutput {
	return o
}

func (o LookupReplicationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationPolicyResultOutput) Properties() PolicyPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationPolicyResult) PolicyPropertiesResponse { return v.Properties }).(PolicyPropertiesResponseOutput)
}

func (o LookupReplicationPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationPolicyResultOutput{})
}
