


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListKustoPoolLanguageExtensions(ctx *pulumi.Context, args *ListKustoPoolLanguageExtensionsArgs, opts ...pulumi.InvokeOption) (*ListKustoPoolLanguageExtensionsResult, error) {
	var rv ListKustoPoolLanguageExtensionsResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:listKustoPoolLanguageExtensions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListKustoPoolLanguageExtensionsArgs struct {
	KustoPoolName     string `pulumi:"kustoPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListKustoPoolLanguageExtensionsResult struct {
	Value []LanguageExtensionResponse `pulumi:"value"`
}

func ListKustoPoolLanguageExtensionsOutput(ctx *pulumi.Context, args ListKustoPoolLanguageExtensionsOutputArgs, opts ...pulumi.InvokeOption) ListKustoPoolLanguageExtensionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListKustoPoolLanguageExtensionsResult, error) {
			args := v.(ListKustoPoolLanguageExtensionsArgs)
			r, err := ListKustoPoolLanguageExtensions(ctx, &args, opts...)
			var s ListKustoPoolLanguageExtensionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListKustoPoolLanguageExtensionsResultOutput)
}

type ListKustoPoolLanguageExtensionsOutputArgs struct {
	KustoPoolName     pulumi.StringInput `pulumi:"kustoPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListKustoPoolLanguageExtensionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKustoPoolLanguageExtensionsArgs)(nil)).Elem()
}


type ListKustoPoolLanguageExtensionsResultOutput struct{ *pulumi.OutputState }

func (ListKustoPoolLanguageExtensionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKustoPoolLanguageExtensionsResult)(nil)).Elem()
}

func (o ListKustoPoolLanguageExtensionsResultOutput) ToListKustoPoolLanguageExtensionsResultOutput() ListKustoPoolLanguageExtensionsResultOutput {
	return o
}

func (o ListKustoPoolLanguageExtensionsResultOutput) ToListKustoPoolLanguageExtensionsResultOutputWithContext(ctx context.Context) ListKustoPoolLanguageExtensionsResultOutput {
	return o
}

func (o ListKustoPoolLanguageExtensionsResultOutput) Value() LanguageExtensionResponseArrayOutput {
	return o.ApplyT(func(v ListKustoPoolLanguageExtensionsResult) []LanguageExtensionResponse { return v.Value }).(LanguageExtensionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListKustoPoolLanguageExtensionsResultOutput{})
}
