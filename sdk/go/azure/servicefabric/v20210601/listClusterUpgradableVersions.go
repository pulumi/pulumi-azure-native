


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListClusterUpgradableVersions(ctx *pulumi.Context, args *ListClusterUpgradableVersionsArgs, opts ...pulumi.InvokeOption) (*ListClusterUpgradableVersionsResult, error) {
	var rv ListClusterUpgradableVersionsResult
	err := ctx.Invoke("azure-native:servicefabric/v20210601:listClusterUpgradableVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterUpgradableVersionsArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TargetVersion     string `pulumi:"targetVersion"`
}


type ListClusterUpgradableVersionsResult struct {
	SupportedPath []string `pulumi:"supportedPath"`
}

func ListClusterUpgradableVersionsOutput(ctx *pulumi.Context, args ListClusterUpgradableVersionsOutputArgs, opts ...pulumi.InvokeOption) ListClusterUpgradableVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListClusterUpgradableVersionsResult, error) {
			args := v.(ListClusterUpgradableVersionsArgs)
			r, err := ListClusterUpgradableVersions(ctx, &args, opts...)
			var s ListClusterUpgradableVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListClusterUpgradableVersionsResultOutput)
}

type ListClusterUpgradableVersionsOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TargetVersion     pulumi.StringInput `pulumi:"targetVersion"`
}

func (ListClusterUpgradableVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterUpgradableVersionsArgs)(nil)).Elem()
}


type ListClusterUpgradableVersionsResultOutput struct{ *pulumi.OutputState }

func (ListClusterUpgradableVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterUpgradableVersionsResult)(nil)).Elem()
}

func (o ListClusterUpgradableVersionsResultOutput) ToListClusterUpgradableVersionsResultOutput() ListClusterUpgradableVersionsResultOutput {
	return o
}

func (o ListClusterUpgradableVersionsResultOutput) ToListClusterUpgradableVersionsResultOutputWithContext(ctx context.Context) ListClusterUpgradableVersionsResultOutput {
	return o
}

func (o ListClusterUpgradableVersionsResultOutput) SupportedPath() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListClusterUpgradableVersionsResult) []string { return v.SupportedPath }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListClusterUpgradableVersionsResultOutput{})
}
