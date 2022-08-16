


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppAzureStorageAccountsSlot(ctx *pulumi.Context, args *ListWebAppAzureStorageAccountsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppAzureStorageAccountsSlotResult, error) {
	var rv ListWebAppAzureStorageAccountsSlotResult
	err := ctx.Invoke("azure-native:web/v20201001:listWebAppAzureStorageAccountsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppAzureStorageAccountsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppAzureStorageAccountsSlotResult struct {
	Id         string                                   `pulumi:"id"`
	Kind       *string                                  `pulumi:"kind"`
	Name       string                                   `pulumi:"name"`
	Properties map[string]AzureStorageInfoValueResponse `pulumi:"properties"`
	SystemData SystemDataResponse                       `pulumi:"systemData"`
	Type       string                                   `pulumi:"type"`
}

func ListWebAppAzureStorageAccountsSlotOutput(ctx *pulumi.Context, args ListWebAppAzureStorageAccountsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppAzureStorageAccountsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppAzureStorageAccountsSlotResult, error) {
			args := v.(ListWebAppAzureStorageAccountsSlotArgs)
			r, err := ListWebAppAzureStorageAccountsSlot(ctx, &args, opts...)
			var s ListWebAppAzureStorageAccountsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppAzureStorageAccountsSlotResultOutput)
}

type ListWebAppAzureStorageAccountsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppAzureStorageAccountsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAzureStorageAccountsSlotArgs)(nil)).Elem()
}


type ListWebAppAzureStorageAccountsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppAzureStorageAccountsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAzureStorageAccountsSlotResult)(nil)).Elem()
}

func (o ListWebAppAzureStorageAccountsSlotResultOutput) ToListWebAppAzureStorageAccountsSlotResultOutput() ListWebAppAzureStorageAccountsSlotResultOutput {
	return o
}

func (o ListWebAppAzureStorageAccountsSlotResultOutput) ToListWebAppAzureStorageAccountsSlotResultOutputWithContext(ctx context.Context) ListWebAppAzureStorageAccountsSlotResultOutput {
	return o
}

func (o ListWebAppAzureStorageAccountsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppAzureStorageAccountsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAzureStorageAccountsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppAzureStorageAccountsSlotResultOutput) Properties() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) map[string]AzureStorageInfoValueResponse {
		return v.Properties
	}).(AzureStorageInfoValueResponseMapOutput)
}

func (o ListWebAppAzureStorageAccountsSlotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ListWebAppAzureStorageAccountsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppAzureStorageAccountsSlotResultOutput{})
}
