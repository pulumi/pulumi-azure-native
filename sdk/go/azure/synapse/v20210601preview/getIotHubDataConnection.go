


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotHubDataConnection(ctx *pulumi.Context, args *LookupIotHubDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupIotHubDataConnectionResult, error) {
	var rv LookupIotHubDataConnectionResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getIotHubDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotHubDataConnectionArgs struct {
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	KustoPoolName      string `pulumi:"kustoPoolName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	WorkspaceName      string `pulumi:"workspaceName"`
}


type LookupIotHubDataConnectionResult struct {
	ConsumerGroup          string             `pulumi:"consumerGroup"`
	DataFormat             *string            `pulumi:"dataFormat"`
	EventSystemProperties  []string           `pulumi:"eventSystemProperties"`
	Id                     string             `pulumi:"id"`
	IotHubResourceId       string             `pulumi:"iotHubResourceId"`
	Kind                   string             `pulumi:"kind"`
	Location               *string            `pulumi:"location"`
	MappingRuleName        *string            `pulumi:"mappingRuleName"`
	Name                   string             `pulumi:"name"`
	ProvisioningState      string             `pulumi:"provisioningState"`
	SharedAccessPolicyName string             `pulumi:"sharedAccessPolicyName"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
	TableName              *string            `pulumi:"tableName"`
	Type                   string             `pulumi:"type"`
}
