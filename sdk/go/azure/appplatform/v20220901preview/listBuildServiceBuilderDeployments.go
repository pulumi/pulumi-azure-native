


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBuildServiceBuilderDeployments(ctx *pulumi.Context, args *ListBuildServiceBuilderDeploymentsArgs, opts ...pulumi.InvokeOption) (*ListBuildServiceBuilderDeploymentsResult, error) {
	var rv ListBuildServiceBuilderDeploymentsResult
	err := ctx.Invoke("azure-native:appplatform/v20220901preview:listBuildServiceBuilderDeployments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBuildServiceBuilderDeploymentsArgs struct {
	BuildServiceName  string `pulumi:"buildServiceName"`
	BuilderName       string `pulumi:"builderName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListBuildServiceBuilderDeploymentsResult struct {
	Deployments []string `pulumi:"deployments"`
}

func ListBuildServiceBuilderDeploymentsOutput(ctx *pulumi.Context, args ListBuildServiceBuilderDeploymentsOutputArgs, opts ...pulumi.InvokeOption) ListBuildServiceBuilderDeploymentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBuildServiceBuilderDeploymentsResult, error) {
			args := v.(ListBuildServiceBuilderDeploymentsArgs)
			r, err := ListBuildServiceBuilderDeployments(ctx, &args, opts...)
			var s ListBuildServiceBuilderDeploymentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBuildServiceBuilderDeploymentsResultOutput)
}

type ListBuildServiceBuilderDeploymentsOutputArgs struct {
	BuildServiceName  pulumi.StringInput `pulumi:"buildServiceName"`
	BuilderName       pulumi.StringInput `pulumi:"builderName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (ListBuildServiceBuilderDeploymentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBuildServiceBuilderDeploymentsArgs)(nil)).Elem()
}


type ListBuildServiceBuilderDeploymentsResultOutput struct{ *pulumi.OutputState }

func (ListBuildServiceBuilderDeploymentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBuildServiceBuilderDeploymentsResult)(nil)).Elem()
}

func (o ListBuildServiceBuilderDeploymentsResultOutput) ToListBuildServiceBuilderDeploymentsResultOutput() ListBuildServiceBuilderDeploymentsResultOutput {
	return o
}

func (o ListBuildServiceBuilderDeploymentsResultOutput) ToListBuildServiceBuilderDeploymentsResultOutputWithContext(ctx context.Context) ListBuildServiceBuilderDeploymentsResultOutput {
	return o
}

func (o ListBuildServiceBuilderDeploymentsResultOutput) Deployments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListBuildServiceBuilderDeploymentsResult) []string { return v.Deployments }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBuildServiceBuilderDeploymentsResultOutput{})
}
