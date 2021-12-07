


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlPoolTransparentDataEncryption(ctx *pulumi.Context, args *LookupSqlPoolTransparentDataEncryptionArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolTransparentDataEncryptionResult, error) {
	var rv LookupSqlPoolTransparentDataEncryptionResult
	err := ctx.Invoke("azure-native:synapse/v20210401preview:getSqlPoolTransparentDataEncryption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolTransparentDataEncryptionArgs struct {
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	SqlPoolName                   string `pulumi:"sqlPoolName"`
	TransparentDataEncryptionName string `pulumi:"transparentDataEncryptionName"`
	WorkspaceName                 string `pulumi:"workspaceName"`
}


type LookupSqlPoolTransparentDataEncryptionResult struct {
	Id       string  `pulumi:"id"`
	Location string  `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Status   *string `pulumi:"status"`
	Type     string  `pulumi:"type"`
}
