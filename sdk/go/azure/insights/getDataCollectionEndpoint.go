


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataCollectionEndpoint(ctx *pulumi.Context, args *LookupDataCollectionEndpointArgs, opts ...pulumi.InvokeOption) (*LookupDataCollectionEndpointResult, error) {
	var rv LookupDataCollectionEndpointResult
	err := ctx.Invoke("azure-native:insights:getDataCollectionEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataCollectionEndpointArgs struct {
	DataCollectionEndpointName string `pulumi:"dataCollectionEndpointName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}


type LookupDataCollectionEndpointResult struct {
	ConfigurationAccess *DataCollectionEndpointResponseConfigurationAccess `pulumi:"configurationAccess"`
	Description         *string                                            `pulumi:"description"`
	Etag                string                                             `pulumi:"etag"`
	Id                  string                                             `pulumi:"id"`
	ImmutableId         *string                                            `pulumi:"immutableId"`
	Kind                *string                                            `pulumi:"kind"`
	Location            string                                             `pulumi:"location"`
	LogsIngestion       *DataCollectionEndpointResponseLogsIngestion       `pulumi:"logsIngestion"`
	Name                string                                             `pulumi:"name"`
	NetworkAcls         *DataCollectionEndpointResponseNetworkAcls         `pulumi:"networkAcls"`
	ProvisioningState   string                                             `pulumi:"provisioningState"`
	SystemData          DataCollectionEndpointResourceResponseSystemData   `pulumi:"systemData"`
	Tags                map[string]string                                  `pulumi:"tags"`
	Type                string                                             `pulumi:"type"`
}
