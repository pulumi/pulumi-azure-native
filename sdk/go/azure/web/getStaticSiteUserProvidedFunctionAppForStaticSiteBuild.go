


package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuild(ctx *pulumi.Context, args *LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult, error) {
	var rv LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult
	err := ctx.Invoke("azure-native:web:getStaticSiteUserProvidedFunctionAppForStaticSiteBuild", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	FunctionAppName   string `pulumi:"functionAppName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult struct {
	CreatedOn             string  `pulumi:"createdOn"`
	FunctionAppRegion     *string `pulumi:"functionAppRegion"`
	FunctionAppResourceId *string `pulumi:"functionAppResourceId"`
	Id                    string  `pulumi:"id"`
	Kind                  *string `pulumi:"kind"`
	Name                  string  `pulumi:"name"`
	Type                  string  `pulumi:"type"`
}
