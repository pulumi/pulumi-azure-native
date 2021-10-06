


package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAwsCloudTrailDataConnector(ctx *pulumi.Context, args *LookupAwsCloudTrailDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAwsCloudTrailDataConnectorResult, error) {
	var rv LookupAwsCloudTrailDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20200101:getAwsCloudTrailDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAwsCloudTrailDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAwsCloudTrailDataConnectorResult struct {
	AwsRoleArn *string                                      `pulumi:"awsRoleArn"`
	DataTypes  *AwsCloudTrailDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                                      `pulumi:"etag"`
	Id         string                                       `pulumi:"id"`
	Kind       string                                       `pulumi:"kind"`
	Name       string                                       `pulumi:"name"`
	Type       string                                       `pulumi:"type"`
}
