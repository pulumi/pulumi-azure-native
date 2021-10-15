


package v20200202preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEncryptionProtector(ctx *pulumi.Context, args *LookupEncryptionProtectorArgs, opts ...pulumi.InvokeOption) (*LookupEncryptionProtectorResult, error) {
	var rv LookupEncryptionProtectorResult
	err := ctx.Invoke("azure-native:sql/v20200202preview:getEncryptionProtector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEncryptionProtectorArgs struct {
	EncryptionProtectorName string `pulumi:"encryptionProtectorName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	ServerName              string `pulumi:"serverName"`
}


type LookupEncryptionProtectorResult struct {
	Id            string  `pulumi:"id"`
	Kind          string  `pulumi:"kind"`
	Location      string  `pulumi:"location"`
	Name          string  `pulumi:"name"`
	ServerKeyName *string `pulumi:"serverKeyName"`
	ServerKeyType string  `pulumi:"serverKeyType"`
	Subregion     string  `pulumi:"subregion"`
	Thumbprint    string  `pulumi:"thumbprint"`
	Type          string  `pulumi:"type"`
	Uri           string  `pulumi:"uri"`
}
