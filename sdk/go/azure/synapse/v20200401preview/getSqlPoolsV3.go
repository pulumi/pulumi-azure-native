


package v20200401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlPoolsV3(ctx *pulumi.Context, args *LookupSqlPoolsV3Args, opts ...pulumi.InvokeOption) (*LookupSqlPoolsV3Result, error) {
	var rv LookupSqlPoolsV3Result
	err := ctx.Invoke("azure-native:synapse/v20200401preview:getSqlPoolsV3", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolsV3Args struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SqlPoolName       string `pulumi:"sqlPoolName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSqlPoolsV3Result struct {
	AutoPauseTimer                *int               `pulumi:"autoPauseTimer"`
	AutoResume                    *bool              `pulumi:"autoResume"`
	CurrentServiceObjectiveName   string             `pulumi:"currentServiceObjectiveName"`
	Id                            string             `pulumi:"id"`
	Kind                          string             `pulumi:"kind"`
	Location                      string             `pulumi:"location"`
	MaxServiceObjectiveName       *string            `pulumi:"maxServiceObjectiveName"`
	Name                          string             `pulumi:"name"`
	RequestedServiceObjectiveName string             `pulumi:"requestedServiceObjectiveName"`
	Sku                           *SkuV3Response     `pulumi:"sku"`
	SqlPoolGuid                   string             `pulumi:"sqlPoolGuid"`
	Status                        string             `pulumi:"status"`
	SystemData                    SystemDataResponse `pulumi:"systemData"`
	Tags                          map[string]string  `pulumi:"tags"`
	Type                          string             `pulumi:"type"`
}
