


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRemediationDeploymentsAtManagementGroup(ctx *pulumi.Context, args *ListRemediationDeploymentsAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtManagementGroupResult, error) {
	var rv ListRemediationDeploymentsAtManagementGroupResult
	err := ctx.Invoke("azure-native:policyinsights/v20180701preview:listRemediationDeploymentsAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtManagementGroupArgs struct {
	ManagementGroupId         string `pulumi:"managementGroupId"`
	ManagementGroupsNamespace string `pulumi:"managementGroupsNamespace"`
	RemediationName           string `pulumi:"remediationName"`
	Top                       *int   `pulumi:"top"`
}


type ListRemediationDeploymentsAtManagementGroupResult struct {
	NextLink string                          `pulumi:"nextLink"`
	Value    []RemediationDeploymentResponse `pulumi:"value"`
}

func ListRemediationDeploymentsAtManagementGroupOutput(ctx *pulumi.Context, args ListRemediationDeploymentsAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) ListRemediationDeploymentsAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRemediationDeploymentsAtManagementGroupResult, error) {
			args := v.(ListRemediationDeploymentsAtManagementGroupArgs)
			r, err := ListRemediationDeploymentsAtManagementGroup(ctx, &args, opts...)
			var s ListRemediationDeploymentsAtManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRemediationDeploymentsAtManagementGroupResultOutput)
}

type ListRemediationDeploymentsAtManagementGroupOutputArgs struct {
	ManagementGroupId         pulumi.StringInput `pulumi:"managementGroupId"`
	ManagementGroupsNamespace pulumi.StringInput `pulumi:"managementGroupsNamespace"`
	RemediationName           pulumi.StringInput `pulumi:"remediationName"`
	Top                       pulumi.IntPtrInput `pulumi:"top"`
}

func (ListRemediationDeploymentsAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtManagementGroupArgs)(nil)).Elem()
}


type ListRemediationDeploymentsAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (ListRemediationDeploymentsAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtManagementGroupResult)(nil)).Elem()
}

func (o ListRemediationDeploymentsAtManagementGroupResultOutput) ToListRemediationDeploymentsAtManagementGroupResultOutput() ListRemediationDeploymentsAtManagementGroupResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtManagementGroupResultOutput) ToListRemediationDeploymentsAtManagementGroupResultOutputWithContext(ctx context.Context) ListRemediationDeploymentsAtManagementGroupResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtManagementGroupResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtManagementGroupResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListRemediationDeploymentsAtManagementGroupResultOutput) Value() RemediationDeploymentResponseArrayOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtManagementGroupResult) []RemediationDeploymentResponse {
		return v.Value
	}).(RemediationDeploymentResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRemediationDeploymentsAtManagementGroupResultOutput{})
}
