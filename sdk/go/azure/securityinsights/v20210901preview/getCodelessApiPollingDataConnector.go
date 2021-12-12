


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCodelessApiPollingDataConnector(ctx *pulumi.Context, args *LookupCodelessApiPollingDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupCodelessApiPollingDataConnectorResult, error) {
	var rv LookupCodelessApiPollingDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getCodelessApiPollingDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCodelessApiPollingDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupCodelessApiPollingDataConnectorResult struct {
	ConnectorUiConfig *CodelessUiConnectorConfigPropertiesResponse      `pulumi:"connectorUiConfig"`
	Etag              *string                                           `pulumi:"etag"`
	Id                string                                            `pulumi:"id"`
	Kind              string                                            `pulumi:"kind"`
	Name              string                                            `pulumi:"name"`
	PollingConfig     *CodelessConnectorPollingConfigPropertiesResponse `pulumi:"pollingConfig"`
	SystemData        SystemDataResponse                                `pulumi:"systemData"`
	Type              string                                            `pulumi:"type"`
}
