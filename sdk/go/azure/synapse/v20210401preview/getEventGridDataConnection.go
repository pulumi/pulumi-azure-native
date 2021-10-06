


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventGridDataConnection(ctx *pulumi.Context, args *LookupEventGridDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventGridDataConnectionResult, error) {
	var rv LookupEventGridDataConnectionResult
	err := ctx.Invoke("azure-native:synapse/v20210401preview:getEventGridDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventGridDataConnectionArgs struct {
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	KustoPoolName      string `pulumi:"kustoPoolName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	WorkspaceName      string `pulumi:"workspaceName"`
}


type LookupEventGridDataConnectionResult struct {
	BlobStorageEventType     *string            `pulumi:"blobStorageEventType"`
	ConsumerGroup            string             `pulumi:"consumerGroup"`
	DataFormat               *string            `pulumi:"dataFormat"`
	EventHubResourceId       string             `pulumi:"eventHubResourceId"`
	Id                       string             `pulumi:"id"`
	IgnoreFirstRecord        *bool              `pulumi:"ignoreFirstRecord"`
	Kind                     string             `pulumi:"kind"`
	Location                 *string            `pulumi:"location"`
	MappingRuleName          *string            `pulumi:"mappingRuleName"`
	Name                     string             `pulumi:"name"`
	ProvisioningState        string             `pulumi:"provisioningState"`
	StorageAccountResourceId string             `pulumi:"storageAccountResourceId"`
	SystemData               SystemDataResponse `pulumi:"systemData"`
	TableName                *string            `pulumi:"tableName"`
	Type                     string             `pulumi:"type"`
}
