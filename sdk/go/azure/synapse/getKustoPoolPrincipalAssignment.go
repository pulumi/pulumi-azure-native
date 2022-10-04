


package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoPoolPrincipalAssignment(ctx *pulumi.Context, args *LookupKustoPoolPrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolPrincipalAssignmentResult, error) {
	var rv LookupKustoPoolPrincipalAssignmentResult
	err := ctx.Invoke("azure-native:synapse:getKustoPoolPrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolPrincipalAssignmentArgs struct {
	KustoPoolName           string `pulumi:"kustoPoolName"`
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	WorkspaceName           string `pulumi:"workspaceName"`
}


type LookupKustoPoolPrincipalAssignmentResult struct {
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	PrincipalId       string             `pulumi:"principalId"`
	PrincipalName     string             `pulumi:"principalName"`
	PrincipalType     string             `pulumi:"principalType"`
	ProvisioningState string             `pulumi:"provisioningState"`
	Role              string             `pulumi:"role"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	TenantId          *string            `pulumi:"tenantId"`
	TenantName        string             `pulumi:"tenantName"`
	Type              string             `pulumi:"type"`
}

func LookupKustoPoolPrincipalAssignmentOutput(ctx *pulumi.Context, args LookupKustoPoolPrincipalAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupKustoPoolPrincipalAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoPoolPrincipalAssignmentResult, error) {
			args := v.(LookupKustoPoolPrincipalAssignmentArgs)
			r, err := LookupKustoPoolPrincipalAssignment(ctx, &args, opts...)
			var s LookupKustoPoolPrincipalAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoPoolPrincipalAssignmentResultOutput)
}

type LookupKustoPoolPrincipalAssignmentOutputArgs struct {
	KustoPoolName           pulumi.StringInput `pulumi:"kustoPoolName"`
	PrincipalAssignmentName pulumi.StringInput `pulumi:"principalAssignmentName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName           pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupKustoPoolPrincipalAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolPrincipalAssignmentArgs)(nil)).Elem()
}


type LookupKustoPoolPrincipalAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupKustoPoolPrincipalAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolPrincipalAssignmentResult)(nil)).Elem()
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) ToLookupKustoPoolPrincipalAssignmentResultOutput() LookupKustoPoolPrincipalAssignmentResultOutput {
	return o
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) ToLookupKustoPoolPrincipalAssignmentResultOutputWithContext(ctx context.Context) LookupKustoPoolPrincipalAssignmentResultOutput {
	return o
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.TenantName }).(pulumi.StringOutput)
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoPoolPrincipalAssignmentResultOutput{})
}
