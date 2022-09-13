


package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListListUpgradableVersionPost(ctx *pulumi.Context, args *ListListUpgradableVersionPostArgs, opts ...pulumi.InvokeOption) (*ListListUpgradableVersionPostResult, error) {
	var rv ListListUpgradableVersionPostResult
	err := ctx.Invoke("azure-native:servicefabric/v20201201preview:listListUpgradableVersionPost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListUpgradableVersionPostArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TargetVersion     string `pulumi:"targetVersion"`
}


type ListListUpgradableVersionPostResult struct {
	SupportedPath []string `pulumi:"supportedPath"`
}

func ListListUpgradableVersionPostOutput(ctx *pulumi.Context, args ListListUpgradableVersionPostOutputArgs, opts ...pulumi.InvokeOption) ListListUpgradableVersionPostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListListUpgradableVersionPostResult, error) {
			args := v.(ListListUpgradableVersionPostArgs)
			r, err := ListListUpgradableVersionPost(ctx, &args, opts...)
			var s ListListUpgradableVersionPostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListListUpgradableVersionPostResultOutput)
}

type ListListUpgradableVersionPostOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TargetVersion     pulumi.StringInput `pulumi:"targetVersion"`
}

func (ListListUpgradableVersionPostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListUpgradableVersionPostArgs)(nil)).Elem()
}


type ListListUpgradableVersionPostResultOutput struct{ *pulumi.OutputState }

func (ListListUpgradableVersionPostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListUpgradableVersionPostResult)(nil)).Elem()
}

func (o ListListUpgradableVersionPostResultOutput) ToListListUpgradableVersionPostResultOutput() ListListUpgradableVersionPostResultOutput {
	return o
}

func (o ListListUpgradableVersionPostResultOutput) ToListListUpgradableVersionPostResultOutputWithContext(ctx context.Context) ListListUpgradableVersionPostResultOutput {
	return o
}

func (o ListListUpgradableVersionPostResultOutput) SupportedPath() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListListUpgradableVersionPostResult) []string { return v.SupportedPath }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListListUpgradableVersionPostResultOutput{})
}
