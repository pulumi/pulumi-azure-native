


package v20210115

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSiteUserProvidedFunctionAppForStaticSite(ctx *pulumi.Context, args *LookupStaticSiteUserProvidedFunctionAppForStaticSiteArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult, error) {
	var rv LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult
	err := ctx.Invoke("azure-native:web/v20210115:getStaticSiteUserProvidedFunctionAppForStaticSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteUserProvidedFunctionAppForStaticSiteArgs struct {
	FunctionAppName   string `pulumi:"functionAppName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult struct {
	CreatedOn             string  `pulumi:"createdOn"`
	FunctionAppRegion     *string `pulumi:"functionAppRegion"`
	FunctionAppResourceId *string `pulumi:"functionAppResourceId"`
	Id                    string  `pulumi:"id"`
	Kind                  *string `pulumi:"kind"`
	Name                  string  `pulumi:"name"`
	Type                  string  `pulumi:"type"`
}
