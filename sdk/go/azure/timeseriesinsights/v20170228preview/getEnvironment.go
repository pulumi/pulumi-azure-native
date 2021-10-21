


package v20170228preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20170228preview:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEnvironmentResult struct {
	CreationTime                 string            `pulumi:"creationTime"`
	DataAccessFqdn               string            `pulumi:"dataAccessFqdn"`
	DataAccessId                 string            `pulumi:"dataAccessId"`
	DataRetentionTime            string            `pulumi:"dataRetentionTime"`
	Id                           string            `pulumi:"id"`
	Location                     string            `pulumi:"location"`
	Name                         string            `pulumi:"name"`
	ProvisioningState            string            `pulumi:"provisioningState"`
	Sku                          *SkuResponse      `pulumi:"sku"`
	StorageLimitExceededBehavior *string           `pulumi:"storageLimitExceededBehavior"`
	Tags                         map[string]string `pulumi:"tags"`
	Type                         string            `pulumi:"type"`
}
