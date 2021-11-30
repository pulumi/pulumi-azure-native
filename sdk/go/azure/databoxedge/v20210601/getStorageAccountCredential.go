


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageAccountCredential(ctx *pulumi.Context, args *LookupStorageAccountCredentialArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountCredentialResult, error) {
	var rv LookupStorageAccountCredentialResult
	err := ctx.Invoke("azure-native:databoxedge/v20210601:getStorageAccountCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountCredentialArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStorageAccountCredentialResult struct {
	AccountKey       *AsymmetricEncryptedSecretResponse `pulumi:"accountKey"`
	AccountType      string                             `pulumi:"accountType"`
	Alias            string                             `pulumi:"alias"`
	BlobDomainName   *string                            `pulumi:"blobDomainName"`
	ConnectionString *string                            `pulumi:"connectionString"`
	Id               string                             `pulumi:"id"`
	Name             string                             `pulumi:"name"`
	SslStatus        string                             `pulumi:"sslStatus"`
	StorageAccountId *string                            `pulumi:"storageAccountId"`
	SystemData       SystemDataResponse                 `pulumi:"systemData"`
	Type             string                             `pulumi:"type"`
	UserName         *string                            `pulumi:"userName"`
}
