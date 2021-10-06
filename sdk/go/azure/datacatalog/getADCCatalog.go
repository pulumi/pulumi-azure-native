


package datacatalog

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADCCatalog(ctx *pulumi.Context, args *LookupADCCatalogArgs, opts ...pulumi.InvokeOption) (*LookupADCCatalogResult, error) {
	var rv LookupADCCatalogResult
	err := ctx.Invoke("azure-native:datacatalog:getADCCatalog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADCCatalogArgs struct {
	CatalogName       string `pulumi:"catalogName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupADCCatalogResult struct {
	Admins                        []PrincipalsResponse `pulumi:"admins"`
	EnableAutomaticUnitAdjustment *bool                `pulumi:"enableAutomaticUnitAdjustment"`
	Etag                          *string              `pulumi:"etag"`
	Id                            string               `pulumi:"id"`
	Location                      *string              `pulumi:"location"`
	Name                          string               `pulumi:"name"`
	Sku                           *string              `pulumi:"sku"`
	SuccessfullyProvisioned       *bool                `pulumi:"successfullyProvisioned"`
	Tags                          map[string]string    `pulumi:"tags"`
	Type                          string               `pulumi:"type"`
	Units                         *int                 `pulumi:"units"`
	Users                         []PrincipalsResponse `pulumi:"users"`
}
