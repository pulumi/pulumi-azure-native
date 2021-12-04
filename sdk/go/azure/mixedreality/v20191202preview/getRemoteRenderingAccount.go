


package v20191202preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRemoteRenderingAccount(ctx *pulumi.Context, args *LookupRemoteRenderingAccountArgs, opts ...pulumi.InvokeOption) (*LookupRemoteRenderingAccountResult, error) {
	var rv LookupRemoteRenderingAccountResult
	err := ctx.Invoke("azure-native:mixedreality/v20191202preview:getRemoteRenderingAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemoteRenderingAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRemoteRenderingAccountResult struct {
	AccountDomain string            `pulumi:"accountDomain"`
	AccountId     string            `pulumi:"accountId"`
	Id            string            `pulumi:"id"`
	Identity      *IdentityResponse `pulumi:"identity"`
	Location      string            `pulumi:"location"`
	Name          string            `pulumi:"name"`
	Tags          map[string]string `pulumi:"tags"`
	Type          string            `pulumi:"type"`
}
