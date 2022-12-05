


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementPolicy(ctx *pulumi.Context, args *LookupManagementPolicyArgs, opts ...pulumi.InvokeOption) (*LookupManagementPolicyResult, error) {
	var rv LookupManagementPolicyResult
	err := ctx.Invoke("azure-native:storage/v20210401:getManagementPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementPolicyArgs struct {
	AccountName          string `pulumi:"accountName"`
	ManagementPolicyName string `pulumi:"managementPolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupManagementPolicyResult struct {
	Id               string                         `pulumi:"id"`
	LastModifiedTime string                         `pulumi:"lastModifiedTime"`
	Name             string                         `pulumi:"name"`
	Policy           ManagementPolicySchemaResponse `pulumi:"policy"`
	Type             string                         `pulumi:"type"`
}

func LookupManagementPolicyOutput(ctx *pulumi.Context, args LookupManagementPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupManagementPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementPolicyResult, error) {
			args := v.(LookupManagementPolicyArgs)
			r, err := LookupManagementPolicy(ctx, &args, opts...)
			var s LookupManagementPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementPolicyResultOutput)
}

type LookupManagementPolicyOutputArgs struct {
	AccountName          pulumi.StringInput `pulumi:"accountName"`
	ManagementPolicyName pulumi.StringInput `pulumi:"managementPolicyName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagementPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementPolicyArgs)(nil)).Elem()
}


type LookupManagementPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupManagementPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementPolicyResult)(nil)).Elem()
}

func (o LookupManagementPolicyResultOutput) ToLookupManagementPolicyResultOutput() LookupManagementPolicyResultOutput {
	return o
}

func (o LookupManagementPolicyResultOutput) ToLookupManagementPolicyResultOutputWithContext(ctx context.Context) LookupManagementPolicyResultOutput {
	return o
}

func (o LookupManagementPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementPolicyResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementPolicyResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupManagementPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementPolicyResultOutput) Policy() ManagementPolicySchemaResponseOutput {
	return o.ApplyT(func(v LookupManagementPolicyResult) ManagementPolicySchemaResponse { return v.Policy }).(ManagementPolicySchemaResponseOutput)
}

func (o LookupManagementPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementPolicyResultOutput{})
}
