


package v20211115preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlResourceSqlRoleAssignment(ctx *pulumi.Context, args *LookupSqlResourceSqlRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlRoleAssignmentResult, error) {
	var rv LookupSqlResourceSqlRoleAssignmentResult
	err := ctx.Invoke("azure-native:documentdb/v20211115preview:getSqlResourceSqlRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlRoleAssignmentArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoleAssignmentId  string `pulumi:"roleAssignmentId"`
}


type LookupSqlResourceSqlRoleAssignmentResult struct {
	Id               string  `pulumi:"id"`
	Name             string  `pulumi:"name"`
	PrincipalId      *string `pulumi:"principalId"`
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	Scope            *string `pulumi:"scope"`
	Type             string  `pulumi:"type"`
}

func LookupSqlResourceSqlRoleAssignmentOutput(ctx *pulumi.Context, args LookupSqlResourceSqlRoleAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlRoleAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlRoleAssignmentResult, error) {
			args := v.(LookupSqlResourceSqlRoleAssignmentArgs)
			r, err := LookupSqlResourceSqlRoleAssignment(ctx, &args, opts...)
			var s LookupSqlResourceSqlRoleAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlRoleAssignmentResultOutput)
}

type LookupSqlResourceSqlRoleAssignmentOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RoleAssignmentId  pulumi.StringInput `pulumi:"roleAssignmentId"`
}

func (LookupSqlResourceSqlRoleAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlRoleAssignmentArgs)(nil)).Elem()
}


type LookupSqlResourceSqlRoleAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlRoleAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlRoleAssignmentResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlRoleAssignmentResultOutput) ToLookupSqlResourceSqlRoleAssignmentResultOutput() LookupSqlResourceSqlRoleAssignmentResultOutput {
	return o
}

func (o LookupSqlResourceSqlRoleAssignmentResultOutput) ToLookupSqlResourceSqlRoleAssignmentResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlRoleAssignmentResultOutput {
	return o
}

func (o LookupSqlResourceSqlRoleAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlRoleAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlRoleAssignmentResultOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleAssignmentResult) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o LookupSqlResourceSqlRoleAssignmentResultOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleAssignmentResult) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupSqlResourceSqlRoleAssignmentResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleAssignmentResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupSqlResourceSqlRoleAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlRoleAssignmentResultOutput{})
}
