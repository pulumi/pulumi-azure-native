


package v20200113preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAutomationAccount(ctx *pulumi.Context, args *LookupAutomationAccountArgs, opts ...pulumi.InvokeOption) (*LookupAutomationAccountResult, error) {
	var rv LookupAutomationAccountResult
	err := ctx.Invoke("azure-native:automation/v20200113preview:getAutomationAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutomationAccountArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupAutomationAccountResult struct {
	CreationTime               string                              `pulumi:"creationTime"`
	Description                *string                             `pulumi:"description"`
	Encryption                 *EncryptionPropertiesResponse       `pulumi:"encryption"`
	Etag                       *string                             `pulumi:"etag"`
	Id                         string                              `pulumi:"id"`
	Identity                   *IdentityResponse                   `pulumi:"identity"`
	LastModifiedBy             *string                             `pulumi:"lastModifiedBy"`
	LastModifiedTime           string                              `pulumi:"lastModifiedTime"`
	Location                   *string                             `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        *bool                               `pulumi:"publicNetworkAccess"`
	Sku                        *SkuResponse                        `pulumi:"sku"`
	State                      string                              `pulumi:"state"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
}
