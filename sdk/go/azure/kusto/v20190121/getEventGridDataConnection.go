


package v20190121

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventGridDataConnection(ctx *pulumi.Context, args *LookupEventGridDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventGridDataConnectionResult, error) {
	var rv LookupEventGridDataConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20190121:getEventGridDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventGridDataConnectionArgs struct {
	ClusterName        string `pulumi:"clusterName"`
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupEventGridDataConnectionResult struct {
	ConsumerGroup            string  `pulumi:"consumerGroup"`
	DataFormat               string  `pulumi:"dataFormat"`
	EventHubResourceId       string  `pulumi:"eventHubResourceId"`
	Id                       string  `pulumi:"id"`
	Kind                     string  `pulumi:"kind"`
	Location                 *string `pulumi:"location"`
	MappingRuleName          *string `pulumi:"mappingRuleName"`
	Name                     string  `pulumi:"name"`
	StorageAccountResourceId string  `pulumi:"storageAccountResourceId"`
	TableName                string  `pulumi:"tableName"`
	Type                     string  `pulumi:"type"`
}
