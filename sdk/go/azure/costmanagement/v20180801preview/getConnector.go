


package v20180801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	var rv LookupConnectorResult
	err := ctx.Invoke("azure-native:costmanagement/v20180801preview:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	ConnectorName     string `pulumi:"connectorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectorResult struct {
	Collection        ConnectorCollectionInfoResponse `pulumi:"collection"`
	CreatedOn         string                          `pulumi:"createdOn"`
	CredentialsKey    *string                         `pulumi:"credentialsKey"`
	DisplayName       *string                         `pulumi:"displayName"`
	Id                string                          `pulumi:"id"`
	Kind              *string                         `pulumi:"kind"`
	Location          *string                         `pulumi:"location"`
	ModifiedOn        string                          `pulumi:"modifiedOn"`
	Name              string                          `pulumi:"name"`
	ProviderAccountId string                          `pulumi:"providerAccountId"`
	ReportId          *string                         `pulumi:"reportId"`
	Status            *string                         `pulumi:"status"`
	Tags              map[string]string               `pulumi:"tags"`
	Type              string                          `pulumi:"type"`
}
