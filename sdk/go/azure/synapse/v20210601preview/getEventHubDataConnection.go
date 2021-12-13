


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventHubDataConnection(ctx *pulumi.Context, args *LookupEventHubDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventHubDataConnectionResult, error) {
	var rv LookupEventHubDataConnectionResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getEventHubDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubDataConnectionArgs struct {
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	KustoPoolName      string `pulumi:"kustoPoolName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	WorkspaceName      string `pulumi:"workspaceName"`
}


type LookupEventHubDataConnectionResult struct {
	Compression               *string            `pulumi:"compression"`
	ConsumerGroup             string             `pulumi:"consumerGroup"`
	DataFormat                *string            `pulumi:"dataFormat"`
	EventHubResourceId        string             `pulumi:"eventHubResourceId"`
	EventSystemProperties     []string           `pulumi:"eventSystemProperties"`
	Id                        string             `pulumi:"id"`
	Kind                      string             `pulumi:"kind"`
	Location                  *string            `pulumi:"location"`
	ManagedIdentityResourceId *string            `pulumi:"managedIdentityResourceId"`
	MappingRuleName           *string            `pulumi:"mappingRuleName"`
	Name                      string             `pulumi:"name"`
	ProvisioningState         string             `pulumi:"provisioningState"`
	SystemData                SystemDataResponse `pulumi:"systemData"`
	TableName                 *string            `pulumi:"tableName"`
	Type                      string             `pulumi:"type"`
}
