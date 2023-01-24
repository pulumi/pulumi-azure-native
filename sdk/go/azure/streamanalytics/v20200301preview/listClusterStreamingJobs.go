


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListClusterStreamingJobs(ctx *pulumi.Context, args *ListClusterStreamingJobsArgs, opts ...pulumi.InvokeOption) (*ListClusterStreamingJobsResult, error) {
	var rv ListClusterStreamingJobsResult
	err := ctx.Invoke("azure-native:streamanalytics/v20200301preview:listClusterStreamingJobs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterStreamingJobsArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListClusterStreamingJobsResult struct {
	NextLink string               `pulumi:"nextLink"`
	Value    []ClusterJobResponse `pulumi:"value"`
}

func ListClusterStreamingJobsOutput(ctx *pulumi.Context, args ListClusterStreamingJobsOutputArgs, opts ...pulumi.InvokeOption) ListClusterStreamingJobsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListClusterStreamingJobsResult, error) {
			args := v.(ListClusterStreamingJobsArgs)
			r, err := ListClusterStreamingJobs(ctx, &args, opts...)
			var s ListClusterStreamingJobsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListClusterStreamingJobsResultOutput)
}

type ListClusterStreamingJobsOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListClusterStreamingJobsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterStreamingJobsArgs)(nil)).Elem()
}


type ListClusterStreamingJobsResultOutput struct{ *pulumi.OutputState }

func (ListClusterStreamingJobsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterStreamingJobsResult)(nil)).Elem()
}

func (o ListClusterStreamingJobsResultOutput) ToListClusterStreamingJobsResultOutput() ListClusterStreamingJobsResultOutput {
	return o
}

func (o ListClusterStreamingJobsResultOutput) ToListClusterStreamingJobsResultOutputWithContext(ctx context.Context) ListClusterStreamingJobsResultOutput {
	return o
}

func (o ListClusterStreamingJobsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListClusterStreamingJobsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListClusterStreamingJobsResultOutput) Value() ClusterJobResponseArrayOutput {
	return o.ApplyT(func(v ListClusterStreamingJobsResult) []ClusterJobResponse { return v.Value }).(ClusterJobResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListClusterStreamingJobsResultOutput{})
}
