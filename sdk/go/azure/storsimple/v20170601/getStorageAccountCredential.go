


package v20170601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageAccountCredential(ctx *pulumi.Context, args *LookupStorageAccountCredentialArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountCredentialResult, error) {
	var rv LookupStorageAccountCredentialResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:getStorageAccountCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountCredentialArgs struct {
	ManagerName                  string `pulumi:"managerName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
	StorageAccountCredentialName string `pulumi:"storageAccountCredentialName"`
}


type LookupStorageAccountCredentialResult struct {
	AccessKey    *AsymmetricEncryptedSecretResponse `pulumi:"accessKey"`
	EndPoint     string                             `pulumi:"endPoint"`
	Id           string                             `pulumi:"id"`
	Kind         *string                            `pulumi:"kind"`
	Name         string                             `pulumi:"name"`
	SslStatus    string                             `pulumi:"sslStatus"`
	Type         string                             `pulumi:"type"`
	VolumesCount int                                `pulumi:"volumesCount"`
}
