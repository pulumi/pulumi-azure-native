


package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTransparentDataEncryption(ctx *pulumi.Context, args *LookupTransparentDataEncryptionArgs, opts ...pulumi.InvokeOption) (*LookupTransparentDataEncryptionResult, error) {
	var rv LookupTransparentDataEncryptionResult
	err := ctx.Invoke("azure-native:sql:getTransparentDataEncryption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransparentDataEncryptionArgs struct {
	DatabaseName                  string `pulumi:"databaseName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ServerName                    string `pulumi:"serverName"`
	TransparentDataEncryptionName string `pulumi:"transparentDataEncryptionName"`
}


type LookupTransparentDataEncryptionResult struct {
	Id       string  `pulumi:"id"`
	Location string  `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Status   *string `pulumi:"status"`
	Type     string  `pulumi:"type"`
}
