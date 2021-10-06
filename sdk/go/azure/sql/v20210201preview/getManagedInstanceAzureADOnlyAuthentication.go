


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedInstanceAzureADOnlyAuthentication(ctx *pulumi.Context, args *LookupManagedInstanceAzureADOnlyAuthenticationArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceAzureADOnlyAuthenticationResult, error) {
	var rv LookupManagedInstanceAzureADOnlyAuthenticationResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getManagedInstanceAzureADOnlyAuthentication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceAzureADOnlyAuthenticationArgs struct {
	AuthenticationName  string `pulumi:"authenticationName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedInstanceAzureADOnlyAuthenticationResult struct {
	AzureADOnlyAuthentication bool   `pulumi:"azureADOnlyAuthentication"`
	Id                        string `pulumi:"id"`
	Name                      string `pulumi:"name"`
	Type                      string `pulumi:"type"`
}
