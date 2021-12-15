


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCodelessUiDataConnector(ctx *pulumi.Context, args *LookupCodelessUiDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupCodelessUiDataConnectorResult, error) {
	var rv LookupCodelessUiDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getCodelessUiDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCodelessUiDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupCodelessUiDataConnectorResult struct {
	ConnectorUiConfig *CodelessUiConnectorConfigPropertiesResponse `pulumi:"connectorUiConfig"`
	Etag              *string                                      `pulumi:"etag"`
	Id                string                                       `pulumi:"id"`
	Kind              string                                       `pulumi:"kind"`
	Name              string                                       `pulumi:"name"`
	SystemData        SystemDataResponse                           `pulumi:"systemData"`
	Type              string                                       `pulumi:"type"`
}
