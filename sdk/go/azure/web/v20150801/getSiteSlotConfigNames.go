


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSiteSlotConfigNames(ctx *pulumi.Context, args *LookupSiteSlotConfigNamesArgs, opts ...pulumi.InvokeOption) (*LookupSiteSlotConfigNamesResult, error) {
	var rv LookupSiteSlotConfigNamesResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteSlotConfigNames", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteSlotConfigNamesArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteSlotConfigNamesResult struct {
	AppSettingNames       []string          `pulumi:"appSettingNames"`
	ConnectionStringNames []string          `pulumi:"connectionStringNames"`
	Id                    *string           `pulumi:"id"`
	Kind                  *string           `pulumi:"kind"`
	Location              string            `pulumi:"location"`
	Name                  *string           `pulumi:"name"`
	Tags                  map[string]string `pulumi:"tags"`
	Type                  *string           `pulumi:"type"`
}
