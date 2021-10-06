


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:azurestackhci/v20201001:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupClusterResult struct {
	AadClientId           string                             `pulumi:"aadClientId"`
	AadTenantId           string                             `pulumi:"aadTenantId"`
	BillingModel          string                             `pulumi:"billingModel"`
	CloudId               string                             `pulumi:"cloudId"`
	CreatedAt             *string                            `pulumi:"createdAt"`
	CreatedBy             *string                            `pulumi:"createdBy"`
	CreatedByType         *string                            `pulumi:"createdByType"`
	Id                    string                             `pulumi:"id"`
	LastBillingTimestamp  string                             `pulumi:"lastBillingTimestamp"`
	LastModifiedAt        *string                            `pulumi:"lastModifiedAt"`
	LastModifiedBy        *string                            `pulumi:"lastModifiedBy"`
	LastModifiedByType    *string                            `pulumi:"lastModifiedByType"`
	LastSyncTimestamp     string                             `pulumi:"lastSyncTimestamp"`
	Location              string                             `pulumi:"location"`
	Name                  string                             `pulumi:"name"`
	ProvisioningState     string                             `pulumi:"provisioningState"`
	RegistrationTimestamp string                             `pulumi:"registrationTimestamp"`
	ReportedProperties    *ClusterReportedPropertiesResponse `pulumi:"reportedProperties"`
	Status                string                             `pulumi:"status"`
	Tags                  map[string]string                  `pulumi:"tags"`
	TrialDaysRemaining    float64                            `pulumi:"trialDaysRemaining"`
	Type                  string                             `pulumi:"type"`
}
