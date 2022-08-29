


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupWorkspaceAadAdmin(ctx *pulumi.Context, args *LookupWorkspaceAadAdminArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceAadAdminResult, error) {
	var rv LookupWorkspaceAadAdminResult
	err := ctx.Invoke("azure-native:synapse/v20190601preview:getWorkspaceAadAdmin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceAadAdminArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceAadAdminResult struct {
	AdministratorType *string `pulumi:"administratorType"`
	Id                string  `pulumi:"id"`
	Login             *string `pulumi:"login"`
	Name              string  `pulumi:"name"`
	Sid               *string `pulumi:"sid"`
	TenantId          *string `pulumi:"tenantId"`
	Type              string  `pulumi:"type"`
}

func LookupWorkspaceAadAdminOutput(ctx *pulumi.Context, args LookupWorkspaceAadAdminOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceAadAdminResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceAadAdminResult, error) {
			args := v.(LookupWorkspaceAadAdminArgs)
			r, err := LookupWorkspaceAadAdmin(ctx, &args, opts...)
			var s LookupWorkspaceAadAdminResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceAadAdminResultOutput)
}

type LookupWorkspaceAadAdminOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceAadAdminOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceAadAdminArgs)(nil)).Elem()
}


type LookupWorkspaceAadAdminResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceAadAdminResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceAadAdminResult)(nil)).Elem()
}

func (o LookupWorkspaceAadAdminResultOutput) ToLookupWorkspaceAadAdminResultOutput() LookupWorkspaceAadAdminResultOutput {
	return o
}

func (o LookupWorkspaceAadAdminResultOutput) ToLookupWorkspaceAadAdminResultOutputWithContext(ctx context.Context) LookupWorkspaceAadAdminResultOutput {
	return o
}

func (o LookupWorkspaceAadAdminResultOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) *string { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceAadAdminResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceAadAdminResultOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) *string { return v.Login }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceAadAdminResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceAadAdminResultOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceAadAdminResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceAadAdminResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceAadAdminResultOutput{})
}
