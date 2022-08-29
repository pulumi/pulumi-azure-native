


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRemediationDeploymentsAtResourceGroup(ctx *pulumi.Context, args *ListRemediationDeploymentsAtResourceGroupArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtResourceGroupResult, error) {
	var rv ListRemediationDeploymentsAtResourceGroupResult
	err := ctx.Invoke("azure-native:policyinsights/v20211001:listRemediationDeploymentsAtResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtResourceGroupArgs struct {
	RemediationName   string `pulumi:"remediationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Top               *int   `pulumi:"top"`
}


type ListRemediationDeploymentsAtResourceGroupResult struct {
	NextLink string                          `pulumi:"nextLink"`
	Value    []RemediationDeploymentResponse `pulumi:"value"`
}

func ListRemediationDeploymentsAtResourceGroupOutput(ctx *pulumi.Context, args ListRemediationDeploymentsAtResourceGroupOutputArgs, opts ...pulumi.InvokeOption) ListRemediationDeploymentsAtResourceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRemediationDeploymentsAtResourceGroupResult, error) {
			args := v.(ListRemediationDeploymentsAtResourceGroupArgs)
			r, err := ListRemediationDeploymentsAtResourceGroup(ctx, &args, opts...)
			var s ListRemediationDeploymentsAtResourceGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRemediationDeploymentsAtResourceGroupResultOutput)
}

type ListRemediationDeploymentsAtResourceGroupOutputArgs struct {
	RemediationName   pulumi.StringInput `pulumi:"remediationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Top               pulumi.IntPtrInput `pulumi:"top"`
}

func (ListRemediationDeploymentsAtResourceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtResourceGroupArgs)(nil)).Elem()
}


type ListRemediationDeploymentsAtResourceGroupResultOutput struct{ *pulumi.OutputState }

func (ListRemediationDeploymentsAtResourceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtResourceGroupResult)(nil)).Elem()
}

func (o ListRemediationDeploymentsAtResourceGroupResultOutput) ToListRemediationDeploymentsAtResourceGroupResultOutput() ListRemediationDeploymentsAtResourceGroupResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtResourceGroupResultOutput) ToListRemediationDeploymentsAtResourceGroupResultOutputWithContext(ctx context.Context) ListRemediationDeploymentsAtResourceGroupResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtResourceGroupResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtResourceGroupResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListRemediationDeploymentsAtResourceGroupResultOutput) Value() RemediationDeploymentResponseArrayOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtResourceGroupResult) []RemediationDeploymentResponse {
		return v.Value
	}).(RemediationDeploymentResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRemediationDeploymentsAtResourceGroupResultOutput{})
}
