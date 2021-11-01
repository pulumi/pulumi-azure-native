


package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageDomain(ctx *pulumi.Context, args *LookupStorageDomainArgs, opts ...pulumi.InvokeOption) (*LookupStorageDomainResult, error) {
	var rv LookupStorageDomainResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getStorageDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageDomainArgs struct {
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StorageDomainName string `pulumi:"storageDomainName"`
}


type LookupStorageDomainResult struct {
	EncryptionKey               *AsymmetricEncryptedSecretResponse `pulumi:"encryptionKey"`
	EncryptionStatus            string                             `pulumi:"encryptionStatus"`
	Id                          string                             `pulumi:"id"`
	Name                        string                             `pulumi:"name"`
	StorageAccountCredentialIds []string                           `pulumi:"storageAccountCredentialIds"`
	Type                        string                             `pulumi:"type"`
}
