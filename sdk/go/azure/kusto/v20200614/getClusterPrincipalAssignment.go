


package v20200614

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupClusterPrincipalAssignment(ctx *pulumi.Context, args *LookupClusterPrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupClusterPrincipalAssignmentResult, error) {
	var rv LookupClusterPrincipalAssignmentResult
	err := ctx.Invoke("azure-native:kusto/v20200614:getClusterPrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterPrincipalAssignmentArgs struct {
	ClusterName             string `pulumi:"clusterName"`
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupClusterPrincipalAssignmentResult struct {
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

func LookupClusterPrincipalAssignmentOutput(ctx *pulumi.Context, args LookupClusterPrincipalAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupClusterPrincipalAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterPrincipalAssignmentResult, error) {
			args := v.(LookupClusterPrincipalAssignmentArgs)
			r, err := LookupClusterPrincipalAssignment(ctx, &args, opts...)
			var s LookupClusterPrincipalAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterPrincipalAssignmentResultOutput)
}

type LookupClusterPrincipalAssignmentOutputArgs struct {
	ClusterName             pulumi.StringInput `pulumi:"clusterName"`
	PrincipalAssignmentName pulumi.StringInput `pulumi:"principalAssignmentName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterPrincipalAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterPrincipalAssignmentArgs)(nil)).Elem()
}


type LookupClusterPrincipalAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupClusterPrincipalAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterPrincipalAssignmentResult)(nil)).Elem()
}

func (o LookupClusterPrincipalAssignmentResultOutput) ToLookupClusterPrincipalAssignmentResultOutput() LookupClusterPrincipalAssignmentResultOutput {
	return o
}

func (o LookupClusterPrincipalAssignmentResultOutput) ToLookupClusterPrincipalAssignmentResultOutputWithContext(ctx context.Context) LookupClusterPrincipalAssignmentResultOutput {
	return o
}

func (o LookupClusterPrincipalAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterPrincipalAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupClusterPrincipalAssignmentResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupClusterPrincipalAssignmentResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o LookupClusterPrincipalAssignmentResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o LookupClusterPrincipalAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupClusterPrincipalAssignmentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o LookupClusterPrincipalAssignmentResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupClusterPrincipalAssignmentResultOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.TenantName }).(pulumi.StringOutput)
}

func (o LookupClusterPrincipalAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterPrincipalAssignmentResultOutput{})
}
