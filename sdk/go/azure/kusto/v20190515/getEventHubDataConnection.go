


package v20190515

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventHubDataConnection(ctx *pulumi.Context, args *LookupEventHubDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventHubDataConnectionResult, error) {
	var rv LookupEventHubDataConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20190515:getEventHubDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubDataConnectionArgs struct {
	ClusterName        string `pulumi:"clusterName"`
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupEventHubDataConnectionResult struct {
	ConsumerGroup         string   `pulumi:"consumerGroup"`
	DataFormat            *string  `pulumi:"dataFormat"`
	EventHubResourceId    string   `pulumi:"eventHubResourceId"`
	EventSystemProperties []string `pulumi:"eventSystemProperties"`
	Id                    string   `pulumi:"id"`
	Kind                  string   `pulumi:"kind"`
	Location              *string  `pulumi:"location"`
	MappingRuleName       *string  `pulumi:"mappingRuleName"`
	Name                  string   `pulumi:"name"`
	TableName             *string  `pulumi:"tableName"`
	Type                  string   `pulumi:"type"`
}
