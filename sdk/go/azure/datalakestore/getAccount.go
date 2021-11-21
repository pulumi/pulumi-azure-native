


package datalakestore

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:datalakestore:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	AccountId                   string                       `pulumi:"accountId"`
	CreationTime                string                       `pulumi:"creationTime"`
	CurrentTier                 string                       `pulumi:"currentTier"`
	DefaultGroup                string                       `pulumi:"defaultGroup"`
	EncryptionConfig            EncryptionConfigResponse     `pulumi:"encryptionConfig"`
	EncryptionProvisioningState string                       `pulumi:"encryptionProvisioningState"`
	EncryptionState             string                       `pulumi:"encryptionState"`
	Endpoint                    string                       `pulumi:"endpoint"`
	FirewallAllowAzureIps       string                       `pulumi:"firewallAllowAzureIps"`
	FirewallRules               []FirewallRuleResponse       `pulumi:"firewallRules"`
	FirewallState               string                       `pulumi:"firewallState"`
	Id                          string                       `pulumi:"id"`
	Identity                    EncryptionIdentityResponse   `pulumi:"identity"`
	LastModifiedTime            string                       `pulumi:"lastModifiedTime"`
	Location                    string                       `pulumi:"location"`
	Name                        string                       `pulumi:"name"`
	NewTier                     string                       `pulumi:"newTier"`
	ProvisioningState           string                       `pulumi:"provisioningState"`
	State                       string                       `pulumi:"state"`
	Tags                        map[string]string            `pulumi:"tags"`
	TrustedIdProviderState      string                       `pulumi:"trustedIdProviderState"`
	TrustedIdProviders          []TrustedIdProviderResponse  `pulumi:"trustedIdProviders"`
	Type                        string                       `pulumi:"type"`
	VirtualNetworkRules         []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}
