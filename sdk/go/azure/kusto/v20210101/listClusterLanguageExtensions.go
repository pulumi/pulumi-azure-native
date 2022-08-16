


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListClusterLanguageExtensions(ctx *pulumi.Context, args *ListClusterLanguageExtensionsArgs, opts ...pulumi.InvokeOption) (*ListClusterLanguageExtensionsResult, error) {
	var rv ListClusterLanguageExtensionsResult
	err := ctx.Invoke("azure-native:kusto/v20210101:listClusterLanguageExtensions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterLanguageExtensionsArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListClusterLanguageExtensionsResult struct {
	Value []LanguageExtensionResponse `pulumi:"value"`
}

func ListClusterLanguageExtensionsOutput(ctx *pulumi.Context, args ListClusterLanguageExtensionsOutputArgs, opts ...pulumi.InvokeOption) ListClusterLanguageExtensionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListClusterLanguageExtensionsResult, error) {
			args := v.(ListClusterLanguageExtensionsArgs)
			r, err := ListClusterLanguageExtensions(ctx, &args, opts...)
			var s ListClusterLanguageExtensionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListClusterLanguageExtensionsResultOutput)
}

type ListClusterLanguageExtensionsOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListClusterLanguageExtensionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterLanguageExtensionsArgs)(nil)).Elem()
}


type ListClusterLanguageExtensionsResultOutput struct{ *pulumi.OutputState }

func (ListClusterLanguageExtensionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterLanguageExtensionsResult)(nil)).Elem()
}

func (o ListClusterLanguageExtensionsResultOutput) ToListClusterLanguageExtensionsResultOutput() ListClusterLanguageExtensionsResultOutput {
	return o
}

func (o ListClusterLanguageExtensionsResultOutput) ToListClusterLanguageExtensionsResultOutputWithContext(ctx context.Context) ListClusterLanguageExtensionsResultOutput {
	return o
}

func (o ListClusterLanguageExtensionsResultOutput) Value() LanguageExtensionResponseArrayOutput {
	return o.ApplyT(func(v ListClusterLanguageExtensionsResult) []LanguageExtensionResponse { return v.Value }).(LanguageExtensionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListClusterLanguageExtensionsResultOutput{})
}
