package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := appplatform.NewStorage(ctx, "storage", &appplatform.StorageArgs{
Properties: interface{}{
AccountKey: pulumi.String("account-key-of-storage-account"),
AccountName: pulumi.String("storage-account-name"),
StorageType: pulumi.String("StorageAccount"),
},
ResourceGroupName: pulumi.String("myResourceGroup"),
ServiceName: pulumi.String("myservice"),
StorageName: pulumi.String("mystorage"),
})
if err != nil {
return err
}
return nil
})
}
