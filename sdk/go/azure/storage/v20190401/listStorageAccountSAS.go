


package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStorageAccountSAS(ctx *pulumi.Context, args *ListStorageAccountSASArgs, opts ...pulumi.InvokeOption) (*ListStorageAccountSASResult, error) {
	var rv ListStorageAccountSASResult
	err := ctx.Invoke("azure-native:storage/v20190401:listStorageAccountSAS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStorageAccountSASArgs struct {
	AccountName            string        `pulumi:"accountName"`
	IPAddressOrRange       *string       `pulumi:"iPAddressOrRange"`
	KeyToSign              *string       `pulumi:"keyToSign"`
	Permissions            string        `pulumi:"permissions"`
	Protocols              *HttpProtocol `pulumi:"protocols"`
	ResourceGroupName      string        `pulumi:"resourceGroupName"`
	ResourceTypes          string        `pulumi:"resourceTypes"`
	Services               string        `pulumi:"services"`
	SharedAccessExpiryTime string        `pulumi:"sharedAccessExpiryTime"`
	SharedAccessStartTime  *string       `pulumi:"sharedAccessStartTime"`
}


type ListStorageAccountSASResult struct {
	AccountSasToken string `pulumi:"accountSasToken"`
}
