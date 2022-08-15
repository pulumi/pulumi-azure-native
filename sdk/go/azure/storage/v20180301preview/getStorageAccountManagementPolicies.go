


package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupStorageAccountManagementPolicies(ctx *pulumi.Context, args *LookupStorageAccountManagementPoliciesArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountManagementPoliciesResult, error) {
	var rv LookupStorageAccountManagementPoliciesResult
	err := ctx.Invoke("azure-native:storage/v20180301preview:getStorageAccountManagementPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountManagementPoliciesArgs struct {
	AccountName          string `pulumi:"accountName"`
	ManagementPolicyName string `pulumi:"managementPolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupStorageAccountManagementPoliciesResult struct {
	Id               string      `pulumi:"id"`
	LastModifiedTime string      `pulumi:"lastModifiedTime"`
	Name             string      `pulumi:"name"`
	Policy           interface{} `pulumi:"policy"`
	Type             string      `pulumi:"type"`
}

func LookupStorageAccountManagementPoliciesOutput(ctx *pulumi.Context, args LookupStorageAccountManagementPoliciesOutputArgs, opts ...pulumi.InvokeOption) LookupStorageAccountManagementPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageAccountManagementPoliciesResult, error) {
			args := v.(LookupStorageAccountManagementPoliciesArgs)
			r, err := LookupStorageAccountManagementPolicies(ctx, &args, opts...)
			var s LookupStorageAccountManagementPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageAccountManagementPoliciesResultOutput)
}

type LookupStorageAccountManagementPoliciesOutputArgs struct {
	AccountName          pulumi.StringInput `pulumi:"accountName"`
	ManagementPolicyName pulumi.StringInput `pulumi:"managementPolicyName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStorageAccountManagementPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountManagementPoliciesArgs)(nil)).Elem()
}


type LookupStorageAccountManagementPoliciesResultOutput struct{ *pulumi.OutputState }

func (LookupStorageAccountManagementPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountManagementPoliciesResult)(nil)).Elem()
}

func (o LookupStorageAccountManagementPoliciesResultOutput) ToLookupStorageAccountManagementPoliciesResultOutput() LookupStorageAccountManagementPoliciesResultOutput {
	return o
}

func (o LookupStorageAccountManagementPoliciesResultOutput) ToLookupStorageAccountManagementPoliciesResultOutputWithContext(ctx context.Context) LookupStorageAccountManagementPoliciesResultOutput {
	return o
}

func (o LookupStorageAccountManagementPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountManagementPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageAccountManagementPoliciesResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountManagementPoliciesResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupStorageAccountManagementPoliciesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountManagementPoliciesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageAccountManagementPoliciesResultOutput) Policy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupStorageAccountManagementPoliciesResult) interface{} { return v.Policy }).(pulumi.AnyOutput)
}

func (o LookupStorageAccountManagementPoliciesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountManagementPoliciesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountManagementPoliciesResultOutput{})
}
