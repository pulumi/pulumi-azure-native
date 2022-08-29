


package v20151001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAccountSasTokens(ctx *pulumi.Context, args *ListAccountSasTokensArgs, opts ...pulumi.InvokeOption) (*ListAccountSasTokensResult, error) {
	var rv ListAccountSasTokensResult
	err := ctx.Invoke("azure-native:datalakeanalytics/v20151001preview:listAccountSasTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountSasTokensArgs struct {
	AccountName        string `pulumi:"accountName"`
	ContainerName      string `pulumi:"containerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageAccountName string `pulumi:"storageAccountName"`
}


type ListAccountSasTokensResult struct {
	NextLink string                        `pulumi:"nextLink"`
	Value    []SasTokenInformationResponse `pulumi:"value"`
}

func ListAccountSasTokensOutput(ctx *pulumi.Context, args ListAccountSasTokensOutputArgs, opts ...pulumi.InvokeOption) ListAccountSasTokensResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAccountSasTokensResult, error) {
			args := v.(ListAccountSasTokensArgs)
			r, err := ListAccountSasTokens(ctx, &args, opts...)
			var s ListAccountSasTokensResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAccountSasTokensResultOutput)
}

type ListAccountSasTokensOutputArgs struct {
	AccountName        pulumi.StringInput `pulumi:"accountName"`
	ContainerName      pulumi.StringInput `pulumi:"containerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}

func (ListAccountSasTokensOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountSasTokensArgs)(nil)).Elem()
}


type ListAccountSasTokensResultOutput struct{ *pulumi.OutputState }

func (ListAccountSasTokensResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountSasTokensResult)(nil)).Elem()
}

func (o ListAccountSasTokensResultOutput) ToListAccountSasTokensResultOutput() ListAccountSasTokensResultOutput {
	return o
}

func (o ListAccountSasTokensResultOutput) ToListAccountSasTokensResultOutputWithContext(ctx context.Context) ListAccountSasTokensResultOutput {
	return o
}

func (o ListAccountSasTokensResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountSasTokensResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListAccountSasTokensResultOutput) Value() SasTokenInformationResponseArrayOutput {
	return o.ApplyT(func(v ListAccountSasTokensResult) []SasTokenInformationResponse { return v.Value }).(SasTokenInformationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAccountSasTokensResultOutput{})
}
