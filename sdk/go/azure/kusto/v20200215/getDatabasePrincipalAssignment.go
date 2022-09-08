


package v20200215

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabasePrincipalAssignment(ctx *pulumi.Context, args *LookupDatabasePrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupDatabasePrincipalAssignmentResult, error) {
	var rv LookupDatabasePrincipalAssignmentResult
	err := ctx.Invoke("azure-native:kusto/v20200215:getDatabasePrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabasePrincipalAssignmentArgs struct {
	ClusterName             string `pulumi:"clusterName"`
	DatabaseName            string `pulumi:"databaseName"`
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupDatabasePrincipalAssignmentResult struct {
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	PrincipalId       string  `pulumi:"principalId"`
	PrincipalName     string  `pulumi:"principalName"`
	PrincipalType     string  `pulumi:"principalType"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Role              string  `pulumi:"role"`
	TenantId          *string `pulumi:"tenantId"`
	TenantName        string  `pulumi:"tenantName"`
	Type              string  `pulumi:"type"`
}

func LookupDatabasePrincipalAssignmentOutput(ctx *pulumi.Context, args LookupDatabasePrincipalAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupDatabasePrincipalAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabasePrincipalAssignmentResult, error) {
			args := v.(LookupDatabasePrincipalAssignmentArgs)
			r, err := LookupDatabasePrincipalAssignment(ctx, &args, opts...)
			var s LookupDatabasePrincipalAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabasePrincipalAssignmentResultOutput)
}

type LookupDatabasePrincipalAssignmentOutputArgs struct {
	ClusterName             pulumi.StringInput `pulumi:"clusterName"`
	DatabaseName            pulumi.StringInput `pulumi:"databaseName"`
	PrincipalAssignmentName pulumi.StringInput `pulumi:"principalAssignmentName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabasePrincipalAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasePrincipalAssignmentArgs)(nil)).Elem()
}


type LookupDatabasePrincipalAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupDatabasePrincipalAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasePrincipalAssignmentResult)(nil)).Elem()
}

func (o LookupDatabasePrincipalAssignmentResultOutput) ToLookupDatabasePrincipalAssignmentResultOutput() LookupDatabasePrincipalAssignmentResultOutput {
	return o
}

func (o LookupDatabasePrincipalAssignmentResultOutput) ToLookupDatabasePrincipalAssignmentResultOutputWithContext(ctx context.Context) LookupDatabasePrincipalAssignmentResultOutput {
	return o
}

func (o LookupDatabasePrincipalAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabasePrincipalAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabasePrincipalAssignmentResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupDatabasePrincipalAssignmentResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o LookupDatabasePrincipalAssignmentResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o LookupDatabasePrincipalAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDatabasePrincipalAssignmentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o LookupDatabasePrincipalAssignmentResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupDatabasePrincipalAssignmentResultOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.TenantName }).(pulumi.StringOutput)
}

func (o LookupDatabasePrincipalAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabasePrincipalAssignmentResultOutput{})
}
