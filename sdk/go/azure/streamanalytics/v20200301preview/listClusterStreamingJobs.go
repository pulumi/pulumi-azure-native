


package v20200301preview

import (
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
