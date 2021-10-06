


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTransparentDataEncryption(ctx *pulumi.Context, args *LookupTransparentDataEncryptionArgs, opts ...pulumi.InvokeOption) (*LookupTransparentDataEncryptionResult, error) {
	var rv LookupTransparentDataEncryptionResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getTransparentDataEncryption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransparentDataEncryptionArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	TdeName           string `pulumi:"tdeName"`
}


type LookupTransparentDataEncryptionResult struct {
	Id    string `pulumi:"id"`
	Name  string `pulumi:"name"`
	State string `pulumi:"state"`
	Type  string `pulumi:"type"`
}
