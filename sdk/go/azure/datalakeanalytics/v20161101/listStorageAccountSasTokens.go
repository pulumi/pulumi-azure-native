


package v20161101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStorageAccountSasTokens(ctx *pulumi.Context, args *ListStorageAccountSasTokensArgs, opts ...pulumi.InvokeOption) (*ListStorageAccountSasTokensResult, error) {
	var rv ListStorageAccountSasTokensResult
	err := ctx.Invoke("azure-native:datalakeanalytics/v20161101:listStorageAccountSasTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStorageAccountSasTokensArgs struct {
	AccountName        string `pulumi:"accountName"`
	ContainerName      string `pulumi:"containerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageAccountName string `pulumi:"storageAccountName"`
}


type ListStorageAccountSasTokensResult struct {
	NextLink string                        `pulumi:"nextLink"`
	Value    []SasTokenInformationResponse `pulumi:"value"`
}

func ListStorageAccountSasTokensOutput(ctx *pulumi.Context, args ListStorageAccountSasTokensOutputArgs, opts ...pulumi.InvokeOption) ListStorageAccountSasTokensResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStorageAccountSasTokensResult, error) {
			args := v.(ListStorageAccountSasTokensArgs)
			r, err := ListStorageAccountSasTokens(ctx, &args, opts...)
			var s ListStorageAccountSasTokensResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStorageAccountSasTokensResultOutput)
}

type ListStorageAccountSasTokensOutputArgs struct {
	AccountName        pulumi.StringInput `pulumi:"accountName"`
	ContainerName      pulumi.StringInput `pulumi:"containerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}

func (ListStorageAccountSasTokensOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountSasTokensArgs)(nil)).Elem()
}


type ListStorageAccountSasTokensResultOutput struct{ *pulumi.OutputState }

func (ListStorageAccountSasTokensResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountSasTokensResult)(nil)).Elem()
}

func (o ListStorageAccountSasTokensResultOutput) ToListStorageAccountSasTokensResultOutput() ListStorageAccountSasTokensResultOutput {
	return o
}

func (o ListStorageAccountSasTokensResultOutput) ToListStorageAccountSasTokensResultOutputWithContext(ctx context.Context) ListStorageAccountSasTokensResultOutput {
	return o
}

func (o ListStorageAccountSasTokensResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListStorageAccountSasTokensResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListStorageAccountSasTokensResultOutput) Value() SasTokenInformationResponseArrayOutput {
	return o.ApplyT(func(v ListStorageAccountSasTokensResult) []SasTokenInformationResponse { return v.Value }).(SasTokenInformationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStorageAccountSasTokensResultOutput{})
}
