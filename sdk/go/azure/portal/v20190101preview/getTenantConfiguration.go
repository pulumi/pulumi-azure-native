


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTenantConfiguration(ctx *pulumi.Context, args *LookupTenantConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupTenantConfigurationResult, error) {
	var rv LookupTenantConfigurationResult
	err := ctx.Invoke("azure-native:portal/v20190101preview:getTenantConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTenantConfigurationArgs struct {
	ConfigurationName string `pulumi:"configurationName"`
}


type LookupTenantConfigurationResult struct {
	EnforcePrivateMarkdownStorage *bool  `pulumi:"enforcePrivateMarkdownStorage"`
	Id                            string `pulumi:"id"`
	Name                          string `pulumi:"name"`
	Type                          string `pulumi:"type"`
}
