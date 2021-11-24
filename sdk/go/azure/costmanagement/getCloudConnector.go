


package costmanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudConnector(ctx *pulumi.Context, args *LookupCloudConnectorArgs, opts ...pulumi.InvokeOption) (*LookupCloudConnectorResult, error) {
	var rv LookupCloudConnectorResult
	err := ctx.Invoke("azure-native:costmanagement:getCloudConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudConnectorArgs struct {
	ConnectorName string  `pulumi:"connectorName"`
	Expand        *string `pulumi:"expand"`
}


type LookupCloudConnectorResult struct {
	BillingModel                      *string                         `pulumi:"billingModel"`
	CollectionInfo                    ConnectorCollectionInfoResponse `pulumi:"collectionInfo"`
	CreatedOn                         string                          `pulumi:"createdOn"`
	CredentialsKey                    *string                         `pulumi:"credentialsKey"`
	DaysTrialRemaining                int                             `pulumi:"daysTrialRemaining"`
	DefaultManagementGroupId          *string                         `pulumi:"defaultManagementGroupId"`
	DisplayName                       *string                         `pulumi:"displayName"`
	ExternalBillingAccountId          string                          `pulumi:"externalBillingAccountId"`
	Id                                string                          `pulumi:"id"`
	Kind                              *string                         `pulumi:"kind"`
	ModifiedOn                        string                          `pulumi:"modifiedOn"`
	Name                              string                          `pulumi:"name"`
	ProviderBillingAccountDisplayName string                          `pulumi:"providerBillingAccountDisplayName"`
	ProviderBillingAccountId          string                          `pulumi:"providerBillingAccountId"`
	ReportId                          *string                         `pulumi:"reportId"`
	Status                            string                          `pulumi:"status"`
	SubscriptionId                    *string                         `pulumi:"subscriptionId"`
	Type                              string                          `pulumi:"type"`
}
