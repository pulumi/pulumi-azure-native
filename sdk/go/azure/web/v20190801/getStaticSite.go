


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSite(ctx *pulumi.Context, args *LookupStaticSiteArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteResult, error) {
	var rv LookupStaticSiteResult
	err := ctx.Invoke("azure-native:web/v20190801:getStaticSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteResult struct {
	Branch          *string                            `pulumi:"branch"`
	BuildProperties *StaticSiteBuildPropertiesResponse `pulumi:"buildProperties"`
	CustomDomains   []string                           `pulumi:"customDomains"`
	DefaultHostname string                             `pulumi:"defaultHostname"`
	Id              string                             `pulumi:"id"`
	Kind            *string                            `pulumi:"kind"`
	Location        string                             `pulumi:"location"`
	Name            string                             `pulumi:"name"`
	RepositoryToken *string                            `pulumi:"repositoryToken"`
	RepositoryUrl   *string                            `pulumi:"repositoryUrl"`
	Sku             *SkuDescriptionResponse            `pulumi:"sku"`
	Tags            map[string]string                  `pulumi:"tags"`
	Type            string                             `pulumi:"type"`
}
