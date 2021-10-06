


package v20170907privatepreview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventHubConnection(ctx *pulumi.Context, args *LookupEventHubConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventHubConnectionResult, error) {
	var rv LookupEventHubConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20170907privatepreview:getEventHubConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubConnectionArgs struct {
	ClusterName            string `pulumi:"clusterName"`
	DatabaseName           string `pulumi:"databaseName"`
	EventHubConnectionName string `pulumi:"eventHubConnectionName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupEventHubConnectionResult struct {
	ConsumerGroup      string  `pulumi:"consumerGroup"`
	DataFormat         *string `pulumi:"dataFormat"`
	EventHubResourceId string  `pulumi:"eventHubResourceId"`
	Id                 string  `pulumi:"id"`
	Location           *string `pulumi:"location"`
	MappingRuleName    *string `pulumi:"mappingRuleName"`
	Name               string  `pulumi:"name"`
	TableName          *string `pulumi:"tableName"`
	Type               string  `pulumi:"type"`
}
