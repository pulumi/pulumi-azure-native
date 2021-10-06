


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDynamics365DataConnector(ctx *pulumi.Context, args *LookupDynamics365DataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupDynamics365DataConnectorResult, error) {
	var rv LookupDynamics365DataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getDynamics365DataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDynamics365DataConnectorArgs struct {
	DataConnectorId                     string `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupDynamics365DataConnectorResult struct {
	DataTypes  Dynamics365DataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                                   `pulumi:"etag"`
	Id         string                                    `pulumi:"id"`
	Kind       string                                    `pulumi:"kind"`
	Name       string                                    `pulumi:"name"`
	SystemData SystemDataResponse                        `pulumi:"systemData"`
	TenantId   string                                    `pulumi:"tenantId"`
	Type       string                                    `pulumi:"type"`
}
