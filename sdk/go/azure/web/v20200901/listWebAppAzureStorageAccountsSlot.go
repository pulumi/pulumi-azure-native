


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppAzureStorageAccountsSlot(ctx *pulumi.Context, args *ListWebAppAzureStorageAccountsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppAzureStorageAccountsSlotResult, error) {
	var rv ListWebAppAzureStorageAccountsSlotResult
	err := ctx.Invoke("azure-native:web/v20200901:listWebAppAzureStorageAccountsSlot", args, &rv, opts...)
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
