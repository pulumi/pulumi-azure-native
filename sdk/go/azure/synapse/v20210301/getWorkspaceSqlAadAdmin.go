


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspaceSqlAadAdmin(ctx *pulumi.Context, args *LookupWorkspaceSqlAadAdminArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceSqlAadAdminResult, error) {
	var rv LookupWorkspaceSqlAadAdminResult
	err := ctx.Invoke("azure-native:synapse/v20210301:getWorkspaceSqlAadAdmin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceSqlAadAdminArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceSqlAadAdminResult struct {
	AdministratorType *string `pulumi:"administratorType"`
	Id                string  `pulumi:"id"`
	Login             *string `pulumi:"login"`
	Name              string  `pulumi:"name"`
	Sid               *string `pulumi:"sid"`
	TenantId          *string `pulumi:"tenantId"`
	Type              string  `pulumi:"type"`
}

func LookupWorkspaceSqlAadAdminOutput(ctx *pulumi.Context, args LookupWorkspaceSqlAadAdminOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceSqlAadAdminResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceSqlAadAdminResult, error) {
			args := v.(LookupWorkspaceSqlAadAdminArgs)
			r, err := LookupWorkspaceSqlAadAdmin(ctx, &args, opts...)
			var s LookupWorkspaceSqlAadAdminResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceSqlAadAdminResultOutput)
}

type LookupWorkspaceSqlAadAdminOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceSqlAadAdminOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceSqlAadAdminArgs)(nil)).Elem()
}


type LookupWorkspaceSqlAadAdminResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceSqlAadAdminResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceSqlAadAdminResult)(nil)).Elem()
}

func (o LookupWorkspaceSqlAadAdminResultOutput) ToLookupWorkspaceSqlAadAdminResultOutput() LookupWorkspaceSqlAadAdminResultOutput {
	return o
}

func (o LookupWorkspaceSqlAadAdminResultOutput) ToLookupWorkspaceSqlAadAdminResultOutputWithContext(ctx context.Context) LookupWorkspaceSqlAadAdminResultOutput {
	return o
}

func (o LookupWorkspaceSqlAadAdminResultOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) *string { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceSqlAadAdminResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceSqlAadAdminResultOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) *string { return v.Login }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceSqlAadAdminResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceSqlAadAdminResultOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceSqlAadAdminResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceSqlAadAdminResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceSqlAadAdminResultOutput{})
}
