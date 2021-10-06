


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMediaService(ctx *pulumi.Context, args *LookupMediaServiceArgs, opts ...pulumi.InvokeOption) (*LookupMediaServiceResult, error) {
	var rv LookupMediaServiceResult
	err := ctx.Invoke("azure-native:media/v20210601:getMediaService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMediaServiceArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMediaServiceResult struct {
	Encryption            *AccountEncryptionResponse    `pulumi:"encryption"`
	Id                    string                        `pulumi:"id"`
	Identity              *MediaServiceIdentityResponse `pulumi:"identity"`
	KeyDelivery           *KeyDeliveryResponse          `pulumi:"keyDelivery"`
	Location              string                        `pulumi:"location"`
	MediaServiceId        string                        `pulumi:"mediaServiceId"`
	Name                  string                        `pulumi:"name"`
	PublicNetworkAccess   *string                       `pulumi:"publicNetworkAccess"`
	StorageAccounts       []StorageAccountResponse      `pulumi:"storageAccounts"`
	StorageAuthentication *string                       `pulumi:"storageAuthentication"`
	SystemData            SystemDataResponse            `pulumi:"systemData"`
	Tags                  map[string]string             `pulumi:"tags"`
	Type                  string                        `pulumi:"type"`
}
