


package v20200701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerKey(ctx *pulumi.Context, args *LookupServerKeyArgs, opts ...pulumi.InvokeOption) (*LookupServerKeyResult, error) {
	var rv LookupServerKeyResult
	err := ctx.Invoke("azure-native:dbformysql/v20200701preview:getServerKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerKeyArgs struct {
	KeyName           string `pulumi:"keyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerKeyResult struct {
	CreationDate  string  `pulumi:"creationDate"`
	Id            string  `pulumi:"id"`
	Kind          string  `pulumi:"kind"`
	Name          string  `pulumi:"name"`
	ServerKeyType string  `pulumi:"serverKeyType"`
	Type          string  `pulumi:"type"`
	Uri           *string `pulumi:"uri"`
}
