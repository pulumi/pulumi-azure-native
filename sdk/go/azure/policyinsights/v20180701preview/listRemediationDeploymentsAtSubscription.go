


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRemediationDeploymentsAtSubscription(ctx *pulumi.Context, args *ListRemediationDeploymentsAtSubscriptionArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtSubscriptionResult, error) {
	var rv ListRemediationDeploymentsAtSubscriptionResult
	err := ctx.Invoke("azure-native:policyinsights/v20180701preview:listRemediationDeploymentsAtSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtSubscriptionArgs struct {
	RemediationName string `pulumi:"remediationName"`
	Top             *int   `pulumi:"top"`
}


type ListRemediationDeploymentsAtSubscriptionResult struct {
	NextLink string                          `pulumi:"nextLink"`
	Value    []RemediationDeploymentResponse `pulumi:"value"`
}

func ListRemediationDeploymentsAtSubscriptionOutput(ctx *pulumi.Context, args ListRemediationDeploymentsAtSubscriptionOutputArgs, opts ...pulumi.InvokeOption) ListRemediationDeploymentsAtSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRemediationDeploymentsAtSubscriptionResult, error) {
			args := v.(ListRemediationDeploymentsAtSubscriptionArgs)
			r, err := ListRemediationDeploymentsAtSubscription(ctx, &args, opts...)
			var s ListRemediationDeploymentsAtSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRemediationDeploymentsAtSubscriptionResultOutput)
}

type ListRemediationDeploymentsAtSubscriptionOutputArgs struct {
	RemediationName pulumi.StringInput `pulumi:"remediationName"`
	Top             pulumi.IntPtrInput `pulumi:"top"`
}

func (ListRemediationDeploymentsAtSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtSubscriptionArgs)(nil)).Elem()
}


type ListRemediationDeploymentsAtSubscriptionResultOutput struct{ *pulumi.OutputState }

func (ListRemediationDeploymentsAtSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtSubscriptionResult)(nil)).Elem()
}

func (o ListRemediationDeploymentsAtSubscriptionResultOutput) ToListRemediationDeploymentsAtSubscriptionResultOutput() ListRemediationDeploymentsAtSubscriptionResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtSubscriptionResultOutput) ToListRemediationDeploymentsAtSubscriptionResultOutputWithContext(ctx context.Context) ListRemediationDeploymentsAtSubscriptionResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtSubscriptionResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtSubscriptionResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListRemediationDeploymentsAtSubscriptionResultOutput) Value() RemediationDeploymentResponseArrayOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtSubscriptionResult) []RemediationDeploymentResponse { return v.Value }).(RemediationDeploymentResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRemediationDeploymentsAtSubscriptionResultOutput{})
}
