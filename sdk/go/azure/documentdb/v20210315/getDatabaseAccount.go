


package v20210315

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAccount(ctx *pulumi.Context, args *LookupDatabaseAccountArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountResult, error) {
	var rv LookupDatabaseAccountResult
	err := ctx.Invoke("azure-native:documentdb/v20210315:getDatabaseAccount", args, &rv, opts...)
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
	BackupPolicy                       interface{}                         `pulumi:"backupPolicy"`
	Capabilities                       []CapabilityResponse                `pulumi:"capabilities"`
	ConnectorOffer                     *string                             `pulumi:"connectorOffer"`
	ConsistencyPolicy                  *ConsistencyPolicyResponse          `pulumi:"consistencyPolicy"`
	Cors                               []CorsPolicyResponse                `pulumi:"cors"`
	DatabaseAccountOfferType           string                              `pulumi:"databaseAccountOfferType"`
	DefaultIdentity                    *string                             `pulumi:"defaultIdentity"`
	DisableKeyBasedMetadataWriteAccess *bool                               `pulumi:"disableKeyBasedMetadataWriteAccess"`
	DocumentEndpoint                   string                              `pulumi:"documentEndpoint"`
	EnableAnalyticalStorage            *bool                               `pulumi:"enableAnalyticalStorage"`
	EnableAutomaticFailover            *bool                               `pulumi:"enableAutomaticFailover"`
	EnableCassandraConnector           *bool                               `pulumi:"enableCassandraConnector"`
	EnableFreeTier                     *bool                               `pulumi:"enableFreeTier"`
	EnableMultipleWriteLocations       *bool                               `pulumi:"enableMultipleWriteLocations"`
	FailoverPolicies                   []FailoverPolicyResponse            `pulumi:"failoverPolicies"`
	Id                                 string                              `pulumi:"id"`
	Identity                           *ManagedServiceIdentityResponse     `pulumi:"identity"`
	IpRules                            []IpAddressOrRangeResponse          `pulumi:"ipRules"`
	IsVirtualNetworkFilterEnabled      *bool                               `pulumi:"isVirtualNetworkFilterEnabled"`
	KeyVaultKeyUri                     *string                             `pulumi:"keyVaultKeyUri"`
	Kind                               *string                             `pulumi:"kind"`
	Location                           *string                             `pulumi:"location"`
	Locations                          []LocationResponse                  `pulumi:"locations"`
	Name                               string                              `pulumi:"name"`
	NetworkAclBypass                   *string                             `pulumi:"networkAclBypass"`
	NetworkAclBypassResourceIds        []string                            `pulumi:"networkAclBypassResourceIds"`
	PrivateEndpointConnections         []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState                  string                              `pulumi:"provisioningState"`
	PublicNetworkAccess                *string                             `pulumi:"publicNetworkAccess"`
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

func LookupDatabaseAccountOutput(ctx *pulumi.Context, args LookupDatabaseAccountOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountResult, error) {
			args := v.(LookupDatabaseAccountArgs)
			r, err := LookupDatabaseAccount(ctx, &args, opts...)
			var s LookupDatabaseAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountResultOutput)
}

type LookupDatabaseAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountArgs)(nil)).Elem()
}


type LookupDatabaseAccountResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountResult)(nil)).Elem()
}

func (o LookupDatabaseAccountResultOutput) ToLookupDatabaseAccountResultOutput() LookupDatabaseAccountResultOutput {
	return o
}

func (o LookupDatabaseAccountResultOutput) ToLookupDatabaseAccountResultOutputWithContext(ctx context.Context) LookupDatabaseAccountResultOutput {
	return o
}

func (o LookupDatabaseAccountResultOutput) ApiProperties() ApiPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *ApiPropertiesResponse { return v.ApiProperties }).(ApiPropertiesResponsePtrOutput)
}

func (o LookupDatabaseAccountResultOutput) BackupPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) interface{} { return v.BackupPolicy }).(pulumi.AnyOutput)
}

func (o LookupDatabaseAccountResultOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []CapabilityResponse { return v.Capabilities }).(CapabilityResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) ConnectorOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.ConnectorOffer }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) ConsistencyPolicy() ConsistencyPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *ConsistencyPolicyResponse { return v.ConsistencyPolicy }).(ConsistencyPolicyResponsePtrOutput)
}

func (o LookupDatabaseAccountResultOutput) Cors() CorsPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []CorsPolicyResponse { return v.Cors }).(CorsPolicyResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) DatabaseAccountOfferType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.DatabaseAccountOfferType }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) DefaultIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.DefaultIdentity }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) DisableKeyBasedMetadataWriteAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.DisableKeyBasedMetadataWriteAccess }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) DocumentEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.DocumentEndpoint }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) EnableAnalyticalStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableAnalyticalStorage }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) EnableAutomaticFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableAutomaticFailover }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) EnableCassandraConnector() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableCassandraConnector }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) EnableFreeTier() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableFreeTier }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) EnableMultipleWriteLocations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableMultipleWriteLocations }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) FailoverPolicies() FailoverPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []FailoverPolicyResponse { return v.FailoverPolicies }).(FailoverPolicyResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupDatabaseAccountResultOutput) IpRules() IpAddressOrRangeResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []IpAddressOrRangeResponse { return v.IpRules }).(IpAddressOrRangeResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) IsVirtualNetworkFilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.IsVirtualNetworkFilterEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.KeyVaultKeyUri }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) Locations() LocationResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []LocationResponse { return v.Locations }).(LocationResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) NetworkAclBypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.NetworkAclBypass }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) NetworkAclBypassResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []string { return v.NetworkAclBypassResourceIds }).(pulumi.StringArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) ReadLocations() LocationResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []LocationResponse { return v.ReadLocations }).(LocationResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) WriteLocations() LocationResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []LocationResponse { return v.WriteLocations }).(LocationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountResultOutput{})
}
