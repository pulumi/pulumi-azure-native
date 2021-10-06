


package v20200202preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerAzureADOnlyAuthentication(ctx *pulumi.Context, args *LookupServerAzureADOnlyAuthenticationArgs, opts ...pulumi.InvokeOption) (*LookupServerAzureADOnlyAuthenticationResult, error) {
	var rv LookupServerAzureADOnlyAuthenticationResult
	err := ctx.Invoke("azure-native:sql/v20200202preview:getServerAzureADOnlyAuthentication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerAzureADOnlyAuthenticationArgs struct {
	AuthenticationName string `pulumi:"authenticationName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	ServerName         string `pulumi:"serverName"`
}


type LookupServerAzureADOnlyAuthenticationResult struct {
	AzureADOnlyAuthentication bool   `pulumi:"azureADOnlyAuthentication"`
	Id                        string `pulumi:"id"`
	Name                      string `pulumi:"name"`
	Type                      string `pulumi:"type"`
}
