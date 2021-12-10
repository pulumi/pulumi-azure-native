


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAwsS3DataConnector(ctx *pulumi.Context, args *LookupAwsS3DataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAwsS3DataConnectorResult, error) {
	var rv LookupAwsS3DataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getAwsS3DataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAwsS3DataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAwsS3DataConnectorResult struct {
	DataTypes        AwsS3DataConnectorDataTypesResponse `pulumi:"dataTypes"`
	DestinationTable string                              `pulumi:"destinationTable"`
	Etag             *string                             `pulumi:"etag"`
	Id               string                              `pulumi:"id"`
	Kind             string                              `pulumi:"kind"`
	Name             string                              `pulumi:"name"`
	RoleArn          string                              `pulumi:"roleArn"`
	SqsUrls          []string                            `pulumi:"sqsUrls"`
	SystemData       SystemDataResponse                  `pulumi:"systemData"`
	Type             string                              `pulumi:"type"`
}
