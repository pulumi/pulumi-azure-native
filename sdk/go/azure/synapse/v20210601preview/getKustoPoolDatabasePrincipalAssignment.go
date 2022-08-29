


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoPoolDatabasePrincipalAssignment(ctx *pulumi.Context, args *LookupKustoPoolDatabasePrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolDatabasePrincipalAssignmentResult, error) {
	var rv LookupKustoPoolDatabasePrincipalAssignmentResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getKustoPoolDatabasePrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolDatabasePrincipalAssignmentArgs struct {
	DatabaseName            string `pulumi:"databaseName"`
	KustoPoolName           string `pulumi:"kustoPoolName"`
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	WorkspaceName           string `pulumi:"workspaceName"`
}


type LookupKustoPoolDatabasePrincipalAssignmentResult struct {
	AadObjectId       string             `pulumi:"aadObjectId"`
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

func LookupKustoPoolDatabasePrincipalAssignmentOutput(ctx *pulumi.Context, args LookupKustoPoolDatabasePrincipalAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupKustoPoolDatabasePrincipalAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoPoolDatabasePrincipalAssignmentResult, error) {
			args := v.(LookupKustoPoolDatabasePrincipalAssignmentArgs)
			r, err := LookupKustoPoolDatabasePrincipalAssignment(ctx, &args, opts...)
			var s LookupKustoPoolDatabasePrincipalAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoPoolDatabasePrincipalAssignmentResultOutput)
}

type LookupKustoPoolDatabasePrincipalAssignmentOutputArgs struct {
	DatabaseName            pulumi.StringInput `pulumi:"databaseName"`
	KustoPoolName           pulumi.StringInput `pulumi:"kustoPoolName"`
	PrincipalAssignmentName pulumi.StringInput `pulumi:"principalAssignmentName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName           pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupKustoPoolDatabasePrincipalAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolDatabasePrincipalAssignmentArgs)(nil)).Elem()
}


type LookupKustoPoolDatabasePrincipalAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupKustoPoolDatabasePrincipalAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolDatabasePrincipalAssignmentResult)(nil)).Elem()
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) ToLookupKustoPoolDatabasePrincipalAssignmentResultOutput() LookupKustoPoolDatabasePrincipalAssignmentResultOutput {
	return o
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) ToLookupKustoPoolDatabasePrincipalAssignmentResultOutputWithContext(ctx context.Context) LookupKustoPoolDatabasePrincipalAssignmentResultOutput {
	return o
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) AadObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.AadObjectId }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.TenantName }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoPoolDatabasePrincipalAssignmentResultOutput{})
}
