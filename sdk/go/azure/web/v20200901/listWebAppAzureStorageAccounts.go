


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppAzureStorageAccounts(ctx *pulumi.Context, args *ListWebAppAzureStorageAccountsArgs, opts ...pulumi.InvokeOption) (*ListWebAppAzureStorageAccountsResult, error) {
	var rv ListWebAppAzureStorageAccountsResult
	err := ctx.Invoke("azure-native:web/v20200901:listWebAppAzureStorageAccounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppAzureStorageAccountsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppAzureStorageAccountsResult struct {
	Id         string                                   `pulumi:"id"`
	Kind       *string                                  `pulumi:"kind"`
	Name       string                                   `pulumi:"name"`
	Properties map[string]AzureStorageInfoValueResponse `pulumi:"properties"`
	SystemData SystemDataResponse                       `pulumi:"systemData"`
	Type       string                                   `pulumi:"type"`
}

func ListWebAppAzureStorageAccountsOutput(ctx *pulumi.Context, args ListWebAppAzureStorageAccountsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppAzureStorageAccountsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppAzureStorageAccountsResult, error) {
			args := v.(ListWebAppAzureStorageAccountsArgs)
			r, err := ListWebAppAzureStorageAccounts(ctx, &args, opts...)
			var s ListWebAppAzureStorageAccountsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppAzureStorageAccountsResultOutput)
}

type ListWebAppAzureStorageAccountsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppAzureStorageAccountsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAzureStorageAccountsArgs)(nil)).Elem()
}


type ListWebAppAzureStorageAccountsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppAzureStorageAccountsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAzureStorageAccountsResult)(nil)).Elem()
}

func (o ListWebAppAzureStorageAccountsResultOutput) ToListWebAppAzureStorageAccountsResultOutput() ListWebAppAzureStorageAccountsResultOutput {
	return o
}

func (o ListWebAppAzureStorageAccountsResultOutput) ToListWebAppAzureStorageAccountsResultOutputWithContext(ctx context.Context) ListWebAppAzureStorageAccountsResultOutput {
	return o
}

func (o ListWebAppAzureStorageAccountsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppAzureStorageAccountsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAzureStorageAccountsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppAzureStorageAccountsResultOutput) Properties() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsResult) map[string]AzureStorageInfoValueResponse {
		return v.Properties
	}).(AzureStorageInfoValueResponseMapOutput)
}

func (o ListWebAppAzureStorageAccountsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ListWebAppAzureStorageAccountsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppAzureStorageAccountsResultOutput{})
}
