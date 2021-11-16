


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTiTaxiiDataConnector(ctx *pulumi.Context, args *LookupTiTaxiiDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupTiTaxiiDataConnectorResult, error) {
	var rv LookupTiTaxiiDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getTiTaxiiDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTiTaxiiDataConnectorArgs struct {
	DataConnectorId                     string `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupTiTaxiiDataConnectorResult struct {
	CollectionId        *string                               `pulumi:"collectionId"`
	DataTypes           TiTaxiiDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag                *string                               `pulumi:"etag"`
	FriendlyName        *string                               `pulumi:"friendlyName"`
	Id                  string                                `pulumi:"id"`
	Kind                string                                `pulumi:"kind"`
	Name                string                                `pulumi:"name"`
	Password            *string                               `pulumi:"password"`
	PollingFrequency    string                                `pulumi:"pollingFrequency"`
	SystemData          SystemDataResponse                    `pulumi:"systemData"`
	TaxiiLookbackPeriod *string                               `pulumi:"taxiiLookbackPeriod"`
	TaxiiServer         *string                               `pulumi:"taxiiServer"`
	TenantId            string                                `pulumi:"tenantId"`
	Type                string                                `pulumi:"type"`
	UserName            *string                               `pulumi:"userName"`
	WorkspaceId         *string                               `pulumi:"workspaceId"`
}
