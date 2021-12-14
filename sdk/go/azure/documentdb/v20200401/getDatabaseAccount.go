


package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAccount(ctx *pulumi.Context, args *LookupDatabaseAccountArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountResult, error) {
	var rv LookupDatabaseAccountResult
	err := ctx.Invoke("azure-native:documentdb/v20200401:getDatabaseAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDatabaseAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountResult struct {
	ApiProperties                      *ApiPropertiesResponse              `pulumi:"apiProperties"`
	Capabilities                       []CapabilityResponse                `pulumi:"capabilities"`
	ConnectorOffer                     *string                             `pulumi:"connectorOffer"`
	ConsistencyPolicy                  *ConsistencyPolicyResponse          `pulumi:"consistencyPolicy"`
	Cors                               []CorsPolicyResponse                `pulumi:"cors"`
	DatabaseAccountOfferType           string                              `pulumi:"databaseAccountOfferType"`
	DisableKeyBasedMetadataWriteAccess *bool                               `pulumi:"disableKeyBasedMetadataWriteAccess"`
	DocumentEndpoint                   string                              `pulumi:"documentEndpoint"`
	EnableAnalyticalStorage            *bool                               `pulumi:"enableAnalyticalStorage"`
	EnableAutomaticFailover            *bool                               `pulumi:"enableAutomaticFailover"`
	EnableCassandraConnector           *bool                               `pulumi:"enableCassandraConnector"`
	EnableFreeTier                     *bool                               `pulumi:"enableFreeTier"`
	EnableMultipleWriteLocations       *bool                               `pulumi:"enableMultipleWriteLocations"`
	FailoverPolicies                   []FailoverPolicyResponse            `pulumi:"failoverPolicies"`
	Id                                 string                              `pulumi:"id"`
	IpRules                            []IpAddressOrRangeResponse          `pulumi:"ipRules"`
	IsVirtualNetworkFilterEnabled      *bool                               `pulumi:"isVirtualNetworkFilterEnabled"`
	KeyVaultKeyUri                     *string                             `pulumi:"keyVaultKeyUri"`
	Kind                               *string                             `pulumi:"kind"`
	Location                           *string                             `pulumi:"location"`
	Locations                          []LocationResponse                  `pulumi:"locations"`
	Name                               string                              `pulumi:"name"`
	PrivateEndpointConnections         []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState                  string                              `pulumi:"provisioningState"`
	PublicNetworkAccess                string                              `pulumi:"publicNetworkAccess"`
	ReadLocations                      []LocationResponse                  `pulumi:"readLocations"`
	Tags                               map[string]string                   `pulumi:"tags"`
	Type                               string                              `pulumi:"type"`
	VirtualNetworkRules                []VirtualNetworkRuleResponse        `pulumi:"virtualNetworkRules"`
	WriteLocations                     []LocationResponse                  `pulumi:"writeLocations"`
}


func (val *LookupDatabaseAccountResult) Defaults() *LookupDatabaseAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Kind) {
		kind_ := "GlobalDocumentDB"
		tmp.Kind = &kind_
	}
	return &tmp
}
