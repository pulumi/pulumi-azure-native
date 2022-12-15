


package v20211201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGetPrivateDnsZoneSuffixExecute(ctx *pulumi.Context, args *GetGetPrivateDnsZoneSuffixExecuteArgs, opts ...pulumi.InvokeOption) (*GetGetPrivateDnsZoneSuffixExecuteResult, error) {
	var rv GetGetPrivateDnsZoneSuffixExecuteResult
	err := ctx.Invoke("azure-native:dbformysql/v20211201preview:getGetPrivateDnsZoneSuffixExecute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGetPrivateDnsZoneSuffixExecuteArgs struct {
}


type GetGetPrivateDnsZoneSuffixExecuteResult struct {
	PrivateDnsZoneSuffix *string `pulumi:"privateDnsZoneSuffix"`
}
