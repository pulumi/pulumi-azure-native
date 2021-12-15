


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountProperties struct {
	AllowedFqdnList               []string           `pulumi:"allowedFqdnList"`
	ApiProperties                 *ApiProperties     `pulumi:"apiProperties"`
	CustomSubDomainName           *string            `pulumi:"customSubDomainName"`
	DisableLocalAuth              *bool              `pulumi:"disableLocalAuth"`
	Encryption                    *Encryption        `pulumi:"encryption"`
	MigrationToken                *string            `pulumi:"migrationToken"`
	NetworkAcls                   *NetworkRuleSet    `pulumi:"networkAcls"`
	PublicNetworkAccess           *string            `pulumi:"publicNetworkAccess"`
	Restore                       *bool              `pulumi:"restore"`
	RestrictOutboundNetworkAccess *bool              `pulumi:"restrictOutboundNetworkAccess"`
	UserOwnedStorage              []UserOwnedStorage `pulumi:"userOwnedStorage"`
}





type AccountPropertiesInput interface {
	pulumi.Input

	ToAccountPropertiesOutput() AccountPropertiesOutput
	ToAccountPropertiesOutputWithContext(context.Context) AccountPropertiesOutput
}

type AccountPropertiesArgs struct {
	AllowedFqdnList               pulumi.StringArrayInput    `pulumi:"allowedFqdnList"`
	ApiProperties                 ApiPropertiesPtrInput      `pulumi:"apiProperties"`
	CustomSubDomainName           pulumi.StringPtrInput      `pulumi:"customSubDomainName"`
	DisableLocalAuth              pulumi.BoolPtrInput        `pulumi:"disableLocalAuth"`
	Encryption                    EncryptionPtrInput         `pulumi:"encryption"`
	MigrationToken                pulumi.StringPtrInput      `pulumi:"migrationToken"`
	NetworkAcls                   NetworkRuleSetPtrInput     `pulumi:"networkAcls"`
	PublicNetworkAccess           pulumi.StringPtrInput      `pulumi:"publicNetworkAccess"`
	Restore                       pulumi.BoolPtrInput        `pulumi:"restore"`
	RestrictOutboundNetworkAccess pulumi.BoolPtrInput        `pulumi:"restrictOutboundNetworkAccess"`
	UserOwnedStorage              UserOwnedStorageArrayInput `pulumi:"userOwnedStorage"`
}

func (AccountPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountProperties)(nil)).Elem()
}

func (i AccountPropertiesArgs) ToAccountPropertiesOutput() AccountPropertiesOutput {
	return i.ToAccountPropertiesOutputWithContext(context.Background())
}

func (i AccountPropertiesArgs) ToAccountPropertiesOutputWithContext(ctx context.Context) AccountPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPropertiesOutput)
}

func (i AccountPropertiesArgs) ToAccountPropertiesPtrOutput() AccountPropertiesPtrOutput {
	return i.ToAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i AccountPropertiesArgs) ToAccountPropertiesPtrOutputWithContext(ctx context.Context) AccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPropertiesOutput).ToAccountPropertiesPtrOutputWithContext(ctx)
}









type AccountPropertiesPtrInput interface {
	pulumi.Input

	ToAccountPropertiesPtrOutput() AccountPropertiesPtrOutput
	ToAccountPropertiesPtrOutputWithContext(context.Context) AccountPropertiesPtrOutput
}

type accountPropertiesPtrType AccountPropertiesArgs

func AccountPropertiesPtr(v *AccountPropertiesArgs) AccountPropertiesPtrInput {
	return (*accountPropertiesPtrType)(v)
}

func (*accountPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountProperties)(nil)).Elem()
}

func (i *accountPropertiesPtrType) ToAccountPropertiesPtrOutput() AccountPropertiesPtrOutput {
	return i.ToAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i *accountPropertiesPtrType) ToAccountPropertiesPtrOutputWithContext(ctx context.Context) AccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPropertiesPtrOutput)
}

type AccountPropertiesOutput struct{ *pulumi.OutputState }

func (AccountPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountProperties)(nil)).Elem()
}

func (o AccountPropertiesOutput) ToAccountPropertiesOutput() AccountPropertiesOutput {
	return o
}

func (o AccountPropertiesOutput) ToAccountPropertiesOutputWithContext(ctx context.Context) AccountPropertiesOutput {
	return o
}

func (o AccountPropertiesOutput) ToAccountPropertiesPtrOutput() AccountPropertiesPtrOutput {
	return o.ToAccountPropertiesPtrOutputWithContext(context.Background())
}

func (o AccountPropertiesOutput) ToAccountPropertiesPtrOutputWithContext(ctx context.Context) AccountPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountProperties) *AccountProperties {
		return &v
	}).(AccountPropertiesPtrOutput)
}

func (o AccountPropertiesOutput) AllowedFqdnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AccountProperties) []string { return v.AllowedFqdnList }).(pulumi.StringArrayOutput)
}

func (o AccountPropertiesOutput) ApiProperties() ApiPropertiesPtrOutput {
	return o.ApplyT(func(v AccountProperties) *ApiProperties { return v.ApiProperties }).(ApiPropertiesPtrOutput)
}

func (o AccountPropertiesOutput) CustomSubDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountProperties) *string { return v.CustomSubDomainName }).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccountProperties) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesOutput) Encryption() EncryptionPtrOutput {
	return o.ApplyT(func(v AccountProperties) *Encryption { return v.Encryption }).(EncryptionPtrOutput)
}

func (o AccountPropertiesOutput) MigrationToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountProperties) *string { return v.MigrationToken }).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesOutput) NetworkAcls() NetworkRuleSetPtrOutput {
	return o.ApplyT(func(v AccountProperties) *NetworkRuleSet { return v.NetworkAcls }).(NetworkRuleSetPtrOutput)
}

func (o AccountPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesOutput) Restore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccountProperties) *bool { return v.Restore }).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesOutput) RestrictOutboundNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccountProperties) *bool { return v.RestrictOutboundNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesOutput) UserOwnedStorage() UserOwnedStorageArrayOutput {
	return o.ApplyT(func(v AccountProperties) []UserOwnedStorage { return v.UserOwnedStorage }).(UserOwnedStorageArrayOutput)
}

type AccountPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AccountPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountProperties)(nil)).Elem()
}

func (o AccountPropertiesPtrOutput) ToAccountPropertiesPtrOutput() AccountPropertiesPtrOutput {
	return o
}

func (o AccountPropertiesPtrOutput) ToAccountPropertiesPtrOutputWithContext(ctx context.Context) AccountPropertiesPtrOutput {
	return o
}

func (o AccountPropertiesPtrOutput) Elem() AccountPropertiesOutput {
	return o.ApplyT(func(v *AccountProperties) AccountProperties {
		if v != nil {
			return *v
		}
		var ret AccountProperties
		return ret
	}).(AccountPropertiesOutput)
}

func (o AccountPropertiesPtrOutput) AllowedFqdnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccountProperties) []string {
		if v == nil {
			return nil
		}
		return v.AllowedFqdnList
	}).(pulumi.StringArrayOutput)
}

func (o AccountPropertiesPtrOutput) ApiProperties() ApiPropertiesPtrOutput {
	return o.ApplyT(func(v *AccountProperties) *ApiProperties {
		if v == nil {
			return nil
		}
		return v.ApiProperties
	}).(ApiPropertiesPtrOutput)
}

func (o AccountPropertiesPtrOutput) CustomSubDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountProperties) *string {
		if v == nil {
			return nil
		}
		return v.CustomSubDomainName
	}).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesPtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccountProperties) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesPtrOutput) Encryption() EncryptionPtrOutput {
	return o.ApplyT(func(v *AccountProperties) *Encryption {
		if v == nil {
			return nil
		}
		return v.Encryption
	}).(EncryptionPtrOutput)
}

func (o AccountPropertiesPtrOutput) MigrationToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountProperties) *string {
		if v == nil {
			return nil
		}
		return v.MigrationToken
	}).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesPtrOutput) NetworkAcls() NetworkRuleSetPtrOutput {
	return o.ApplyT(func(v *AccountProperties) *NetworkRuleSet {
		if v == nil {
			return nil
		}
		return v.NetworkAcls
	}).(NetworkRuleSetPtrOutput)
}

func (o AccountPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesPtrOutput) Restore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccountProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Restore
	}).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesPtrOutput) RestrictOutboundNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccountProperties) *bool {
		if v == nil {
			return nil
		}
		return v.RestrictOutboundNetworkAccess
	}).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesPtrOutput) UserOwnedStorage() UserOwnedStorageArrayOutput {
	return o.ApplyT(func(v *AccountProperties) []UserOwnedStorage {
		if v == nil {
			return nil
		}
		return v.UserOwnedStorage
	}).(UserOwnedStorageArrayOutput)
}

type AccountPropertiesResponse struct {
	AllowedFqdnList               []string                            `pulumi:"allowedFqdnList"`
	ApiProperties                 *ApiPropertiesResponse              `pulumi:"apiProperties"`
	CallRateLimit                 CallRateLimitResponse               `pulumi:"callRateLimit"`
	Capabilities                  []SkuCapabilityResponse             `pulumi:"capabilities"`
	CustomSubDomainName           *string                             `pulumi:"customSubDomainName"`
	DateCreated                   string                              `pulumi:"dateCreated"`
	DisableLocalAuth              *bool                               `pulumi:"disableLocalAuth"`
	Encryption                    *EncryptionResponse                 `pulumi:"encryption"`
	Endpoint                      string                              `pulumi:"endpoint"`
	Endpoints                     map[string]string                   `pulumi:"endpoints"`
	InternalId                    string                              `pulumi:"internalId"`
	IsMigrated                    bool                                `pulumi:"isMigrated"`
	MigrationToken                *string                             `pulumi:"migrationToken"`
	NetworkAcls                   *NetworkRuleSetResponse             `pulumi:"networkAcls"`
	PrivateEndpointConnections    []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState             string                              `pulumi:"provisioningState"`
	PublicNetworkAccess           *string                             `pulumi:"publicNetworkAccess"`
	QuotaLimit                    QuotaLimitResponse                  `pulumi:"quotaLimit"`
	Restore                       *bool                               `pulumi:"restore"`
	RestrictOutboundNetworkAccess *bool                               `pulumi:"restrictOutboundNetworkAccess"`
	SkuChangeInfo                 SkuChangeInfoResponse               `pulumi:"skuChangeInfo"`
	UserOwnedStorage              []UserOwnedStorageResponse          `pulumi:"userOwnedStorage"`
}





type AccountPropertiesResponseInput interface {
	pulumi.Input

	ToAccountPropertiesResponseOutput() AccountPropertiesResponseOutput
	ToAccountPropertiesResponseOutputWithContext(context.Context) AccountPropertiesResponseOutput
}

type AccountPropertiesResponseArgs struct {
	AllowedFqdnList               pulumi.StringArrayInput                     `pulumi:"allowedFqdnList"`
	ApiProperties                 ApiPropertiesResponsePtrInput               `pulumi:"apiProperties"`
	CallRateLimit                 CallRateLimitResponseInput                  `pulumi:"callRateLimit"`
	Capabilities                  SkuCapabilityResponseArrayInput             `pulumi:"capabilities"`
	CustomSubDomainName           pulumi.StringPtrInput                       `pulumi:"customSubDomainName"`
	DateCreated                   pulumi.StringInput                          `pulumi:"dateCreated"`
	DisableLocalAuth              pulumi.BoolPtrInput                         `pulumi:"disableLocalAuth"`
	Encryption                    EncryptionResponsePtrInput                  `pulumi:"encryption"`
	Endpoint                      pulumi.StringInput                          `pulumi:"endpoint"`
	Endpoints                     pulumi.StringMapInput                       `pulumi:"endpoints"`
	InternalId                    pulumi.StringInput                          `pulumi:"internalId"`
	IsMigrated                    pulumi.BoolInput                            `pulumi:"isMigrated"`
	MigrationToken                pulumi.StringPtrInput                       `pulumi:"migrationToken"`
	NetworkAcls                   NetworkRuleSetResponsePtrInput              `pulumi:"networkAcls"`
	PrivateEndpointConnections    PrivateEndpointConnectionResponseArrayInput `pulumi:"privateEndpointConnections"`
	ProvisioningState             pulumi.StringInput                          `pulumi:"provisioningState"`
	PublicNetworkAccess           pulumi.StringPtrInput                       `pulumi:"publicNetworkAccess"`
	QuotaLimit                    QuotaLimitResponseInput                     `pulumi:"quotaLimit"`
	Restore                       pulumi.BoolPtrInput                         `pulumi:"restore"`
	RestrictOutboundNetworkAccess pulumi.BoolPtrInput                         `pulumi:"restrictOutboundNetworkAccess"`
	SkuChangeInfo                 SkuChangeInfoResponseInput                  `pulumi:"skuChangeInfo"`
	UserOwnedStorage              UserOwnedStorageResponseArrayInput          `pulumi:"userOwnedStorage"`
}

func (AccountPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountPropertiesResponse)(nil)).Elem()
}

func (i AccountPropertiesResponseArgs) ToAccountPropertiesResponseOutput() AccountPropertiesResponseOutput {
	return i.ToAccountPropertiesResponseOutputWithContext(context.Background())
}

func (i AccountPropertiesResponseArgs) ToAccountPropertiesResponseOutputWithContext(ctx context.Context) AccountPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPropertiesResponseOutput)
}

func (i AccountPropertiesResponseArgs) ToAccountPropertiesResponsePtrOutput() AccountPropertiesResponsePtrOutput {
	return i.ToAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AccountPropertiesResponseArgs) ToAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) AccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPropertiesResponseOutput).ToAccountPropertiesResponsePtrOutputWithContext(ctx)
}









type AccountPropertiesResponsePtrInput interface {
	pulumi.Input

	ToAccountPropertiesResponsePtrOutput() AccountPropertiesResponsePtrOutput
	ToAccountPropertiesResponsePtrOutputWithContext(context.Context) AccountPropertiesResponsePtrOutput
}

type accountPropertiesResponsePtrType AccountPropertiesResponseArgs

func AccountPropertiesResponsePtr(v *AccountPropertiesResponseArgs) AccountPropertiesResponsePtrInput {
	return (*accountPropertiesResponsePtrType)(v)
}

func (*accountPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPropertiesResponse)(nil)).Elem()
}

func (i *accountPropertiesResponsePtrType) ToAccountPropertiesResponsePtrOutput() AccountPropertiesResponsePtrOutput {
	return i.ToAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *accountPropertiesResponsePtrType) ToAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) AccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPropertiesResponsePtrOutput)
}

type AccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountPropertiesResponse)(nil)).Elem()
}

func (o AccountPropertiesResponseOutput) ToAccountPropertiesResponseOutput() AccountPropertiesResponseOutput {
	return o
}

func (o AccountPropertiesResponseOutput) ToAccountPropertiesResponseOutputWithContext(ctx context.Context) AccountPropertiesResponseOutput {
	return o
}

func (o AccountPropertiesResponseOutput) ToAccountPropertiesResponsePtrOutput() AccountPropertiesResponsePtrOutput {
	return o.ToAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AccountPropertiesResponseOutput) ToAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) AccountPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountPropertiesResponse) *AccountPropertiesResponse {
		return &v
	}).(AccountPropertiesResponsePtrOutput)
}

func (o AccountPropertiesResponseOutput) AllowedFqdnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) []string { return v.AllowedFqdnList }).(pulumi.StringArrayOutput)
}

func (o AccountPropertiesResponseOutput) ApiProperties() ApiPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) *ApiPropertiesResponse { return v.ApiProperties }).(ApiPropertiesResponsePtrOutput)
}

func (o AccountPropertiesResponseOutput) CallRateLimit() CallRateLimitResponseOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) CallRateLimitResponse { return v.CallRateLimit }).(CallRateLimitResponseOutput)
}

func (o AccountPropertiesResponseOutput) Capabilities() SkuCapabilityResponseArrayOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) []SkuCapabilityResponse { return v.Capabilities }).(SkuCapabilityResponseArrayOutput)
}

func (o AccountPropertiesResponseOutput) CustomSubDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) *string { return v.CustomSubDomainName }).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesResponseOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) string { return v.DateCreated }).(pulumi.StringOutput)
}

func (o AccountPropertiesResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesResponseOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

func (o AccountPropertiesResponseOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o AccountPropertiesResponseOutput) Endpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) map[string]string { return v.Endpoints }).(pulumi.StringMapOutput)
}

func (o AccountPropertiesResponseOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) string { return v.InternalId }).(pulumi.StringOutput)
}

func (o AccountPropertiesResponseOutput) IsMigrated() pulumi.BoolOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) bool { return v.IsMigrated }).(pulumi.BoolOutput)
}

func (o AccountPropertiesResponseOutput) MigrationToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) *string { return v.MigrationToken }).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesResponseOutput) NetworkAcls() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) *NetworkRuleSetResponse { return v.NetworkAcls }).(NetworkRuleSetResponsePtrOutput)
}

func (o AccountPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o AccountPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AccountPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesResponseOutput) QuotaLimit() QuotaLimitResponseOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) QuotaLimitResponse { return v.QuotaLimit }).(QuotaLimitResponseOutput)
}

func (o AccountPropertiesResponseOutput) Restore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) *bool { return v.Restore }).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesResponseOutput) RestrictOutboundNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) *bool { return v.RestrictOutboundNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesResponseOutput) SkuChangeInfo() SkuChangeInfoResponseOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) SkuChangeInfoResponse { return v.SkuChangeInfo }).(SkuChangeInfoResponseOutput)
}

func (o AccountPropertiesResponseOutput) UserOwnedStorage() UserOwnedStorageResponseArrayOutput {
	return o.ApplyT(func(v AccountPropertiesResponse) []UserOwnedStorageResponse { return v.UserOwnedStorage }).(UserOwnedStorageResponseArrayOutput)
}

type AccountPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AccountPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPropertiesResponse)(nil)).Elem()
}

func (o AccountPropertiesResponsePtrOutput) ToAccountPropertiesResponsePtrOutput() AccountPropertiesResponsePtrOutput {
	return o
}

func (o AccountPropertiesResponsePtrOutput) ToAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) AccountPropertiesResponsePtrOutput {
	return o
}

func (o AccountPropertiesResponsePtrOutput) Elem() AccountPropertiesResponseOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) AccountPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AccountPropertiesResponse
		return ret
	}).(AccountPropertiesResponseOutput)
}

func (o AccountPropertiesResponsePtrOutput) AllowedFqdnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedFqdnList
	}).(pulumi.StringArrayOutput)
}

func (o AccountPropertiesResponsePtrOutput) ApiProperties() ApiPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *ApiPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ApiProperties
	}).(ApiPropertiesResponsePtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) CallRateLimit() CallRateLimitResponsePtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *CallRateLimitResponse {
		if v == nil {
			return nil
		}
		return &v.CallRateLimit
	}).(CallRateLimitResponsePtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) Capabilities() SkuCapabilityResponseArrayOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) []SkuCapabilityResponse {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(SkuCapabilityResponseArrayOutput)
}

func (o AccountPropertiesResponsePtrOutput) CustomSubDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomSubDomainName
	}).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) DateCreated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DateCreated
	}).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *EncryptionResponse {
		if v == nil {
			return nil
		}
		return v.Encryption
	}).(EncryptionResponsePtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) Endpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(pulumi.StringMapOutput)
}

func (o AccountPropertiesResponsePtrOutput) InternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InternalId
	}).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) IsMigrated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsMigrated
	}).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) MigrationToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.MigrationToken
	}).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) NetworkAcls() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *NetworkRuleSetResponse {
		if v == nil {
			return nil
		}
		return v.NetworkAcls
	}).(NetworkRuleSetResponsePtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) []PrivateEndpointConnectionResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o AccountPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) QuotaLimit() QuotaLimitResponsePtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *QuotaLimitResponse {
		if v == nil {
			return nil
		}
		return &v.QuotaLimit
	}).(QuotaLimitResponsePtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) Restore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Restore
	}).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) RestrictOutboundNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RestrictOutboundNetworkAccess
	}).(pulumi.BoolPtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) SkuChangeInfo() SkuChangeInfoResponsePtrOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) *SkuChangeInfoResponse {
		if v == nil {
			return nil
		}
		return &v.SkuChangeInfo
	}).(SkuChangeInfoResponsePtrOutput)
}

func (o AccountPropertiesResponsePtrOutput) UserOwnedStorage() UserOwnedStorageResponseArrayOutput {
	return o.ApplyT(func(v *AccountPropertiesResponse) []UserOwnedStorageResponse {
		if v == nil {
			return nil
		}
		return v.UserOwnedStorage
	}).(UserOwnedStorageResponseArrayOutput)
}

type ApiProperties struct {
	AadClientId                    *string `pulumi:"aadClientId"`
	AadTenantId                    *string `pulumi:"aadTenantId"`
	EventHubConnectionString       *string `pulumi:"eventHubConnectionString"`
	QnaAzureSearchEndpointId       *string `pulumi:"qnaAzureSearchEndpointId"`
	QnaAzureSearchEndpointKey      *string `pulumi:"qnaAzureSearchEndpointKey"`
	QnaRuntimeEndpoint             *string `pulumi:"qnaRuntimeEndpoint"`
	StatisticsEnabled              *bool   `pulumi:"statisticsEnabled"`
	StorageAccountConnectionString *string `pulumi:"storageAccountConnectionString"`
	SuperUser                      *string `pulumi:"superUser"`
	WebsiteName                    *string `pulumi:"websiteName"`
}





type ApiPropertiesInput interface {
	pulumi.Input

	ToApiPropertiesOutput() ApiPropertiesOutput
	ToApiPropertiesOutputWithContext(context.Context) ApiPropertiesOutput
}

type ApiPropertiesArgs struct {
	AadClientId                    pulumi.StringPtrInput `pulumi:"aadClientId"`
	AadTenantId                    pulumi.StringPtrInput `pulumi:"aadTenantId"`
	EventHubConnectionString       pulumi.StringPtrInput `pulumi:"eventHubConnectionString"`
	QnaAzureSearchEndpointId       pulumi.StringPtrInput `pulumi:"qnaAzureSearchEndpointId"`
	QnaAzureSearchEndpointKey      pulumi.StringPtrInput `pulumi:"qnaAzureSearchEndpointKey"`
	QnaRuntimeEndpoint             pulumi.StringPtrInput `pulumi:"qnaRuntimeEndpoint"`
	StatisticsEnabled              pulumi.BoolPtrInput   `pulumi:"statisticsEnabled"`
	StorageAccountConnectionString pulumi.StringPtrInput `pulumi:"storageAccountConnectionString"`
	SuperUser                      pulumi.StringPtrInput `pulumi:"superUser"`
	WebsiteName                    pulumi.StringPtrInput `pulumi:"websiteName"`
}

func (ApiPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiProperties)(nil)).Elem()
}

func (i ApiPropertiesArgs) ToApiPropertiesOutput() ApiPropertiesOutput {
	return i.ToApiPropertiesOutputWithContext(context.Background())
}

func (i ApiPropertiesArgs) ToApiPropertiesOutputWithContext(ctx context.Context) ApiPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesOutput)
}

func (i ApiPropertiesArgs) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return i.ToApiPropertiesPtrOutputWithContext(context.Background())
}

func (i ApiPropertiesArgs) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesOutput).ToApiPropertiesPtrOutputWithContext(ctx)
}









type ApiPropertiesPtrInput interface {
	pulumi.Input

	ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput
	ToApiPropertiesPtrOutputWithContext(context.Context) ApiPropertiesPtrOutput
}

type apiPropertiesPtrType ApiPropertiesArgs

func ApiPropertiesPtr(v *ApiPropertiesArgs) ApiPropertiesPtrInput {
	return (*apiPropertiesPtrType)(v)
}

func (*apiPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiProperties)(nil)).Elem()
}

func (i *apiPropertiesPtrType) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return i.ToApiPropertiesPtrOutputWithContext(context.Background())
}

func (i *apiPropertiesPtrType) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesPtrOutput)
}

type ApiPropertiesOutput struct{ *pulumi.OutputState }

func (ApiPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiProperties)(nil)).Elem()
}

func (o ApiPropertiesOutput) ToApiPropertiesOutput() ApiPropertiesOutput {
	return o
}

func (o ApiPropertiesOutput) ToApiPropertiesOutputWithContext(ctx context.Context) ApiPropertiesOutput {
	return o
}

func (o ApiPropertiesOutput) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return o.ToApiPropertiesPtrOutputWithContext(context.Background())
}

func (o ApiPropertiesOutput) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiProperties) *ApiProperties {
		return &v
	}).(ApiPropertiesPtrOutput)
}

func (o ApiPropertiesOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.AadClientId }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.AadTenantId }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.EventHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesOutput) QnaAzureSearchEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.QnaAzureSearchEndpointId }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesOutput) QnaAzureSearchEndpointKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.QnaAzureSearchEndpointKey }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesOutput) QnaRuntimeEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.QnaRuntimeEndpoint }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesOutput) StatisticsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApiProperties) *bool { return v.StatisticsEnabled }).(pulumi.BoolPtrOutput)
}

func (o ApiPropertiesOutput) StorageAccountConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.StorageAccountConnectionString }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesOutput) SuperUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.SuperUser }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesOutput) WebsiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.WebsiteName }).(pulumi.StringPtrOutput)
}

type ApiPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApiPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiProperties)(nil)).Elem()
}

func (o ApiPropertiesPtrOutput) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return o
}

func (o ApiPropertiesPtrOutput) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return o
}

func (o ApiPropertiesPtrOutput) Elem() ApiPropertiesOutput {
	return o.ApplyT(func(v *ApiProperties) ApiProperties {
		if v != nil {
			return *v
		}
		var ret ApiProperties
		return ret
	}).(ApiPropertiesOutput)
}

func (o ApiPropertiesPtrOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.AadClientId
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesPtrOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.AadTenantId
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesPtrOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.EventHubConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesPtrOutput) QnaAzureSearchEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.QnaAzureSearchEndpointId
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesPtrOutput) QnaAzureSearchEndpointKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.QnaAzureSearchEndpointKey
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesPtrOutput) QnaRuntimeEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.QnaRuntimeEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesPtrOutput) StatisticsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *bool {
		if v == nil {
			return nil
		}
		return v.StatisticsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApiPropertiesPtrOutput) StorageAccountConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesPtrOutput) SuperUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.SuperUser
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesPtrOutput) WebsiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.WebsiteName
	}).(pulumi.StringPtrOutput)
}

type ApiPropertiesResponse struct {
	AadClientId                    *string `pulumi:"aadClientId"`
	AadTenantId                    *string `pulumi:"aadTenantId"`
	EventHubConnectionString       *string `pulumi:"eventHubConnectionString"`
	QnaAzureSearchEndpointId       *string `pulumi:"qnaAzureSearchEndpointId"`
	QnaAzureSearchEndpointKey      *string `pulumi:"qnaAzureSearchEndpointKey"`
	QnaRuntimeEndpoint             *string `pulumi:"qnaRuntimeEndpoint"`
	StatisticsEnabled              *bool   `pulumi:"statisticsEnabled"`
	StorageAccountConnectionString *string `pulumi:"storageAccountConnectionString"`
	SuperUser                      *string `pulumi:"superUser"`
	WebsiteName                    *string `pulumi:"websiteName"`
}





type ApiPropertiesResponseInput interface {
	pulumi.Input

	ToApiPropertiesResponseOutput() ApiPropertiesResponseOutput
	ToApiPropertiesResponseOutputWithContext(context.Context) ApiPropertiesResponseOutput
}

type ApiPropertiesResponseArgs struct {
	AadClientId                    pulumi.StringPtrInput `pulumi:"aadClientId"`
	AadTenantId                    pulumi.StringPtrInput `pulumi:"aadTenantId"`
	EventHubConnectionString       pulumi.StringPtrInput `pulumi:"eventHubConnectionString"`
	QnaAzureSearchEndpointId       pulumi.StringPtrInput `pulumi:"qnaAzureSearchEndpointId"`
	QnaAzureSearchEndpointKey      pulumi.StringPtrInput `pulumi:"qnaAzureSearchEndpointKey"`
	QnaRuntimeEndpoint             pulumi.StringPtrInput `pulumi:"qnaRuntimeEndpoint"`
	StatisticsEnabled              pulumi.BoolPtrInput   `pulumi:"statisticsEnabled"`
	StorageAccountConnectionString pulumi.StringPtrInput `pulumi:"storageAccountConnectionString"`
	SuperUser                      pulumi.StringPtrInput `pulumi:"superUser"`
	WebsiteName                    pulumi.StringPtrInput `pulumi:"websiteName"`
}

func (ApiPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPropertiesResponse)(nil)).Elem()
}

func (i ApiPropertiesResponseArgs) ToApiPropertiesResponseOutput() ApiPropertiesResponseOutput {
	return i.ToApiPropertiesResponseOutputWithContext(context.Background())
}

func (i ApiPropertiesResponseArgs) ToApiPropertiesResponseOutputWithContext(ctx context.Context) ApiPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesResponseOutput)
}

func (i ApiPropertiesResponseArgs) ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput {
	return i.ToApiPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ApiPropertiesResponseArgs) ToApiPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesResponseOutput).ToApiPropertiesResponsePtrOutputWithContext(ctx)
}









type ApiPropertiesResponsePtrInput interface {
	pulumi.Input

	ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput
	ToApiPropertiesResponsePtrOutputWithContext(context.Context) ApiPropertiesResponsePtrOutput
}

type apiPropertiesResponsePtrType ApiPropertiesResponseArgs

func ApiPropertiesResponsePtr(v *ApiPropertiesResponseArgs) ApiPropertiesResponsePtrInput {
	return (*apiPropertiesResponsePtrType)(v)
}

func (*apiPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPropertiesResponse)(nil)).Elem()
}

func (i *apiPropertiesResponsePtrType) ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput {
	return i.ToApiPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *apiPropertiesResponsePtrType) ToApiPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesResponsePtrOutput)
}

type ApiPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApiPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPropertiesResponse)(nil)).Elem()
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponseOutput() ApiPropertiesResponseOutput {
	return o
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponseOutputWithContext(ctx context.Context) ApiPropertiesResponseOutput {
	return o
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput {
	return o.ToApiPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiPropertiesResponse) *ApiPropertiesResponse {
		return &v
	}).(ApiPropertiesResponsePtrOutput)
}

func (o ApiPropertiesResponseOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.AadClientId }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponseOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.AadTenantId }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponseOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.EventHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponseOutput) QnaAzureSearchEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.QnaAzureSearchEndpointId }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponseOutput) QnaAzureSearchEndpointKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.QnaAzureSearchEndpointKey }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponseOutput) QnaRuntimeEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.QnaRuntimeEndpoint }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponseOutput) StatisticsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *bool { return v.StatisticsEnabled }).(pulumi.BoolPtrOutput)
}

func (o ApiPropertiesResponseOutput) StorageAccountConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.StorageAccountConnectionString }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponseOutput) SuperUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.SuperUser }).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponseOutput) WebsiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.WebsiteName }).(pulumi.StringPtrOutput)
}

type ApiPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPropertiesResponse)(nil)).Elem()
}

func (o ApiPropertiesResponsePtrOutput) ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput {
	return o
}

func (o ApiPropertiesResponsePtrOutput) ToApiPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiPropertiesResponsePtrOutput {
	return o
}

func (o ApiPropertiesResponsePtrOutput) Elem() ApiPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) ApiPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ApiPropertiesResponse
		return ret
	}).(ApiPropertiesResponseOutput)
}

func (o ApiPropertiesResponsePtrOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadClientId
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponsePtrOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadTenantId
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponsePtrOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.EventHubConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponsePtrOutput) QnaAzureSearchEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.QnaAzureSearchEndpointId
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponsePtrOutput) QnaAzureSearchEndpointKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.QnaAzureSearchEndpointKey
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponsePtrOutput) QnaRuntimeEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.QnaRuntimeEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponsePtrOutput) StatisticsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.StatisticsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApiPropertiesResponsePtrOutput) StorageAccountConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponsePtrOutput) SuperUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SuperUser
	}).(pulumi.StringPtrOutput)
}

func (o ApiPropertiesResponsePtrOutput) WebsiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WebsiteName
	}).(pulumi.StringPtrOutput)
}

type CallRateLimitResponse struct {
	Count         *float64                 `pulumi:"count"`
	RenewalPeriod *float64                 `pulumi:"renewalPeriod"`
	Rules         []ThrottlingRuleResponse `pulumi:"rules"`
}





type CallRateLimitResponseInput interface {
	pulumi.Input

	ToCallRateLimitResponseOutput() CallRateLimitResponseOutput
	ToCallRateLimitResponseOutputWithContext(context.Context) CallRateLimitResponseOutput
}

type CallRateLimitResponseArgs struct {
	Count         pulumi.Float64PtrInput           `pulumi:"count"`
	RenewalPeriod pulumi.Float64PtrInput           `pulumi:"renewalPeriod"`
	Rules         ThrottlingRuleResponseArrayInput `pulumi:"rules"`
}

func (CallRateLimitResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CallRateLimitResponse)(nil)).Elem()
}

func (i CallRateLimitResponseArgs) ToCallRateLimitResponseOutput() CallRateLimitResponseOutput {
	return i.ToCallRateLimitResponseOutputWithContext(context.Background())
}

func (i CallRateLimitResponseArgs) ToCallRateLimitResponseOutputWithContext(ctx context.Context) CallRateLimitResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CallRateLimitResponseOutput)
}

func (i CallRateLimitResponseArgs) ToCallRateLimitResponsePtrOutput() CallRateLimitResponsePtrOutput {
	return i.ToCallRateLimitResponsePtrOutputWithContext(context.Background())
}

func (i CallRateLimitResponseArgs) ToCallRateLimitResponsePtrOutputWithContext(ctx context.Context) CallRateLimitResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CallRateLimitResponseOutput).ToCallRateLimitResponsePtrOutputWithContext(ctx)
}









type CallRateLimitResponsePtrInput interface {
	pulumi.Input

	ToCallRateLimitResponsePtrOutput() CallRateLimitResponsePtrOutput
	ToCallRateLimitResponsePtrOutputWithContext(context.Context) CallRateLimitResponsePtrOutput
}

type callRateLimitResponsePtrType CallRateLimitResponseArgs

func CallRateLimitResponsePtr(v *CallRateLimitResponseArgs) CallRateLimitResponsePtrInput {
	return (*callRateLimitResponsePtrType)(v)
}

func (*callRateLimitResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CallRateLimitResponse)(nil)).Elem()
}

func (i *callRateLimitResponsePtrType) ToCallRateLimitResponsePtrOutput() CallRateLimitResponsePtrOutput {
	return i.ToCallRateLimitResponsePtrOutputWithContext(context.Background())
}

func (i *callRateLimitResponsePtrType) ToCallRateLimitResponsePtrOutputWithContext(ctx context.Context) CallRateLimitResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CallRateLimitResponsePtrOutput)
}

type CallRateLimitResponseOutput struct{ *pulumi.OutputState }

func (CallRateLimitResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CallRateLimitResponse)(nil)).Elem()
}

func (o CallRateLimitResponseOutput) ToCallRateLimitResponseOutput() CallRateLimitResponseOutput {
	return o
}

func (o CallRateLimitResponseOutput) ToCallRateLimitResponseOutputWithContext(ctx context.Context) CallRateLimitResponseOutput {
	return o
}

func (o CallRateLimitResponseOutput) ToCallRateLimitResponsePtrOutput() CallRateLimitResponsePtrOutput {
	return o.ToCallRateLimitResponsePtrOutputWithContext(context.Background())
}

func (o CallRateLimitResponseOutput) ToCallRateLimitResponsePtrOutputWithContext(ctx context.Context) CallRateLimitResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CallRateLimitResponse) *CallRateLimitResponse {
		return &v
	}).(CallRateLimitResponsePtrOutput)
}

func (o CallRateLimitResponseOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CallRateLimitResponse) *float64 { return v.Count }).(pulumi.Float64PtrOutput)
}

func (o CallRateLimitResponseOutput) RenewalPeriod() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CallRateLimitResponse) *float64 { return v.RenewalPeriod }).(pulumi.Float64PtrOutput)
}

func (o CallRateLimitResponseOutput) Rules() ThrottlingRuleResponseArrayOutput {
	return o.ApplyT(func(v CallRateLimitResponse) []ThrottlingRuleResponse { return v.Rules }).(ThrottlingRuleResponseArrayOutput)
}

type CallRateLimitResponsePtrOutput struct{ *pulumi.OutputState }

func (CallRateLimitResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CallRateLimitResponse)(nil)).Elem()
}

func (o CallRateLimitResponsePtrOutput) ToCallRateLimitResponsePtrOutput() CallRateLimitResponsePtrOutput {
	return o
}

func (o CallRateLimitResponsePtrOutput) ToCallRateLimitResponsePtrOutputWithContext(ctx context.Context) CallRateLimitResponsePtrOutput {
	return o
}

func (o CallRateLimitResponsePtrOutput) Elem() CallRateLimitResponseOutput {
	return o.ApplyT(func(v *CallRateLimitResponse) CallRateLimitResponse {
		if v != nil {
			return *v
		}
		var ret CallRateLimitResponse
		return ret
	}).(CallRateLimitResponseOutput)
}

func (o CallRateLimitResponsePtrOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CallRateLimitResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.Float64PtrOutput)
}

func (o CallRateLimitResponsePtrOutput) RenewalPeriod() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CallRateLimitResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RenewalPeriod
	}).(pulumi.Float64PtrOutput)
}

func (o CallRateLimitResponsePtrOutput) Rules() ThrottlingRuleResponseArrayOutput {
	return o.ApplyT(func(v *CallRateLimitResponse) []ThrottlingRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ThrottlingRuleResponseArrayOutput)
}

type CommitmentPeriod struct {
	Count *int    `pulumi:"count"`
	Tier  *string `pulumi:"tier"`
}





type CommitmentPeriodInput interface {
	pulumi.Input

	ToCommitmentPeriodOutput() CommitmentPeriodOutput
	ToCommitmentPeriodOutputWithContext(context.Context) CommitmentPeriodOutput
}

type CommitmentPeriodArgs struct {
	Count pulumi.IntPtrInput    `pulumi:"count"`
	Tier  pulumi.StringPtrInput `pulumi:"tier"`
}

func (CommitmentPeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPeriod)(nil)).Elem()
}

func (i CommitmentPeriodArgs) ToCommitmentPeriodOutput() CommitmentPeriodOutput {
	return i.ToCommitmentPeriodOutputWithContext(context.Background())
}

func (i CommitmentPeriodArgs) ToCommitmentPeriodOutputWithContext(ctx context.Context) CommitmentPeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPeriodOutput)
}

func (i CommitmentPeriodArgs) ToCommitmentPeriodPtrOutput() CommitmentPeriodPtrOutput {
	return i.ToCommitmentPeriodPtrOutputWithContext(context.Background())
}

func (i CommitmentPeriodArgs) ToCommitmentPeriodPtrOutputWithContext(ctx context.Context) CommitmentPeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPeriodOutput).ToCommitmentPeriodPtrOutputWithContext(ctx)
}









type CommitmentPeriodPtrInput interface {
	pulumi.Input

	ToCommitmentPeriodPtrOutput() CommitmentPeriodPtrOutput
	ToCommitmentPeriodPtrOutputWithContext(context.Context) CommitmentPeriodPtrOutput
}

type commitmentPeriodPtrType CommitmentPeriodArgs

func CommitmentPeriodPtr(v *CommitmentPeriodArgs) CommitmentPeriodPtrInput {
	return (*commitmentPeriodPtrType)(v)
}

func (*commitmentPeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPeriod)(nil)).Elem()
}

func (i *commitmentPeriodPtrType) ToCommitmentPeriodPtrOutput() CommitmentPeriodPtrOutput {
	return i.ToCommitmentPeriodPtrOutputWithContext(context.Background())
}

func (i *commitmentPeriodPtrType) ToCommitmentPeriodPtrOutputWithContext(ctx context.Context) CommitmentPeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPeriodPtrOutput)
}

type CommitmentPeriodOutput struct{ *pulumi.OutputState }

func (CommitmentPeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPeriod)(nil)).Elem()
}

func (o CommitmentPeriodOutput) ToCommitmentPeriodOutput() CommitmentPeriodOutput {
	return o
}

func (o CommitmentPeriodOutput) ToCommitmentPeriodOutputWithContext(ctx context.Context) CommitmentPeriodOutput {
	return o
}

func (o CommitmentPeriodOutput) ToCommitmentPeriodPtrOutput() CommitmentPeriodPtrOutput {
	return o.ToCommitmentPeriodPtrOutputWithContext(context.Background())
}

func (o CommitmentPeriodOutput) ToCommitmentPeriodPtrOutputWithContext(ctx context.Context) CommitmentPeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommitmentPeriod) *CommitmentPeriod {
		return &v
	}).(CommitmentPeriodPtrOutput)
}

func (o CommitmentPeriodOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CommitmentPeriod) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o CommitmentPeriodOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommitmentPeriod) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type CommitmentPeriodPtrOutput struct{ *pulumi.OutputState }

func (CommitmentPeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPeriod)(nil)).Elem()
}

func (o CommitmentPeriodPtrOutput) ToCommitmentPeriodPtrOutput() CommitmentPeriodPtrOutput {
	return o
}

func (o CommitmentPeriodPtrOutput) ToCommitmentPeriodPtrOutputWithContext(ctx context.Context) CommitmentPeriodPtrOutput {
	return o
}

func (o CommitmentPeriodPtrOutput) Elem() CommitmentPeriodOutput {
	return o.ApplyT(func(v *CommitmentPeriod) CommitmentPeriod {
		if v != nil {
			return *v
		}
		var ret CommitmentPeriod
		return ret
	}).(CommitmentPeriodOutput)
}

func (o CommitmentPeriodPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CommitmentPeriod) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o CommitmentPeriodPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPeriod) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type CommitmentPeriodResponse struct {
	Count     *int                    `pulumi:"count"`
	EndDate   string                  `pulumi:"endDate"`
	Quota     CommitmentQuotaResponse `pulumi:"quota"`
	StartDate string                  `pulumi:"startDate"`
	Tier      *string                 `pulumi:"tier"`
}





type CommitmentPeriodResponseInput interface {
	pulumi.Input

	ToCommitmentPeriodResponseOutput() CommitmentPeriodResponseOutput
	ToCommitmentPeriodResponseOutputWithContext(context.Context) CommitmentPeriodResponseOutput
}

type CommitmentPeriodResponseArgs struct {
	Count     pulumi.IntPtrInput           `pulumi:"count"`
	EndDate   pulumi.StringInput           `pulumi:"endDate"`
	Quota     CommitmentQuotaResponseInput `pulumi:"quota"`
	StartDate pulumi.StringInput           `pulumi:"startDate"`
	Tier      pulumi.StringPtrInput        `pulumi:"tier"`
}

func (CommitmentPeriodResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPeriodResponse)(nil)).Elem()
}

func (i CommitmentPeriodResponseArgs) ToCommitmentPeriodResponseOutput() CommitmentPeriodResponseOutput {
	return i.ToCommitmentPeriodResponseOutputWithContext(context.Background())
}

func (i CommitmentPeriodResponseArgs) ToCommitmentPeriodResponseOutputWithContext(ctx context.Context) CommitmentPeriodResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPeriodResponseOutput)
}

func (i CommitmentPeriodResponseArgs) ToCommitmentPeriodResponsePtrOutput() CommitmentPeriodResponsePtrOutput {
	return i.ToCommitmentPeriodResponsePtrOutputWithContext(context.Background())
}

func (i CommitmentPeriodResponseArgs) ToCommitmentPeriodResponsePtrOutputWithContext(ctx context.Context) CommitmentPeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPeriodResponseOutput).ToCommitmentPeriodResponsePtrOutputWithContext(ctx)
}









type CommitmentPeriodResponsePtrInput interface {
	pulumi.Input

	ToCommitmentPeriodResponsePtrOutput() CommitmentPeriodResponsePtrOutput
	ToCommitmentPeriodResponsePtrOutputWithContext(context.Context) CommitmentPeriodResponsePtrOutput
}

type commitmentPeriodResponsePtrType CommitmentPeriodResponseArgs

func CommitmentPeriodResponsePtr(v *CommitmentPeriodResponseArgs) CommitmentPeriodResponsePtrInput {
	return (*commitmentPeriodResponsePtrType)(v)
}

func (*commitmentPeriodResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPeriodResponse)(nil)).Elem()
}

func (i *commitmentPeriodResponsePtrType) ToCommitmentPeriodResponsePtrOutput() CommitmentPeriodResponsePtrOutput {
	return i.ToCommitmentPeriodResponsePtrOutputWithContext(context.Background())
}

func (i *commitmentPeriodResponsePtrType) ToCommitmentPeriodResponsePtrOutputWithContext(ctx context.Context) CommitmentPeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPeriodResponsePtrOutput)
}

type CommitmentPeriodResponseOutput struct{ *pulumi.OutputState }

func (CommitmentPeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPeriodResponse)(nil)).Elem()
}

func (o CommitmentPeriodResponseOutput) ToCommitmentPeriodResponseOutput() CommitmentPeriodResponseOutput {
	return o
}

func (o CommitmentPeriodResponseOutput) ToCommitmentPeriodResponseOutputWithContext(ctx context.Context) CommitmentPeriodResponseOutput {
	return o
}

func (o CommitmentPeriodResponseOutput) ToCommitmentPeriodResponsePtrOutput() CommitmentPeriodResponsePtrOutput {
	return o.ToCommitmentPeriodResponsePtrOutputWithContext(context.Background())
}

func (o CommitmentPeriodResponseOutput) ToCommitmentPeriodResponsePtrOutputWithContext(ctx context.Context) CommitmentPeriodResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommitmentPeriodResponse) *CommitmentPeriodResponse {
		return &v
	}).(CommitmentPeriodResponsePtrOutput)
}

func (o CommitmentPeriodResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CommitmentPeriodResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o CommitmentPeriodResponseOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v CommitmentPeriodResponse) string { return v.EndDate }).(pulumi.StringOutput)
}

func (o CommitmentPeriodResponseOutput) Quota() CommitmentQuotaResponseOutput {
	return o.ApplyT(func(v CommitmentPeriodResponse) CommitmentQuotaResponse { return v.Quota }).(CommitmentQuotaResponseOutput)
}

func (o CommitmentPeriodResponseOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v CommitmentPeriodResponse) string { return v.StartDate }).(pulumi.StringOutput)
}

func (o CommitmentPeriodResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommitmentPeriodResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type CommitmentPeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (CommitmentPeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPeriodResponse)(nil)).Elem()
}

func (o CommitmentPeriodResponsePtrOutput) ToCommitmentPeriodResponsePtrOutput() CommitmentPeriodResponsePtrOutput {
	return o
}

func (o CommitmentPeriodResponsePtrOutput) ToCommitmentPeriodResponsePtrOutputWithContext(ctx context.Context) CommitmentPeriodResponsePtrOutput {
	return o
}

func (o CommitmentPeriodResponsePtrOutput) Elem() CommitmentPeriodResponseOutput {
	return o.ApplyT(func(v *CommitmentPeriodResponse) CommitmentPeriodResponse {
		if v != nil {
			return *v
		}
		var ret CommitmentPeriodResponse
		return ret
	}).(CommitmentPeriodResponseOutput)
}

func (o CommitmentPeriodResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CommitmentPeriodResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o CommitmentPeriodResponsePtrOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndDate
	}).(pulumi.StringPtrOutput)
}

func (o CommitmentPeriodResponsePtrOutput) Quota() CommitmentQuotaResponsePtrOutput {
	return o.ApplyT(func(v *CommitmentPeriodResponse) *CommitmentQuotaResponse {
		if v == nil {
			return nil
		}
		return &v.Quota
	}).(CommitmentQuotaResponsePtrOutput)
}

func (o CommitmentPeriodResponsePtrOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartDate
	}).(pulumi.StringPtrOutput)
}

func (o CommitmentPeriodResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPeriodResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type CommitmentPlanProperties struct {
	AutoRenew    *bool             `pulumi:"autoRenew"`
	Current      *CommitmentPeriod `pulumi:"current"`
	HostingModel *string           `pulumi:"hostingModel"`
	Next         *CommitmentPeriod `pulumi:"next"`
	PlanType     *string           `pulumi:"planType"`
}





type CommitmentPlanPropertiesInput interface {
	pulumi.Input

	ToCommitmentPlanPropertiesOutput() CommitmentPlanPropertiesOutput
	ToCommitmentPlanPropertiesOutputWithContext(context.Context) CommitmentPlanPropertiesOutput
}

type CommitmentPlanPropertiesArgs struct {
	AutoRenew    pulumi.BoolPtrInput      `pulumi:"autoRenew"`
	Current      CommitmentPeriodPtrInput `pulumi:"current"`
	HostingModel pulumi.StringPtrInput    `pulumi:"hostingModel"`
	Next         CommitmentPeriodPtrInput `pulumi:"next"`
	PlanType     pulumi.StringPtrInput    `pulumi:"planType"`
}

func (CommitmentPlanPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPlanProperties)(nil)).Elem()
}

func (i CommitmentPlanPropertiesArgs) ToCommitmentPlanPropertiesOutput() CommitmentPlanPropertiesOutput {
	return i.ToCommitmentPlanPropertiesOutputWithContext(context.Background())
}

func (i CommitmentPlanPropertiesArgs) ToCommitmentPlanPropertiesOutputWithContext(ctx context.Context) CommitmentPlanPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanPropertiesOutput)
}

func (i CommitmentPlanPropertiesArgs) ToCommitmentPlanPropertiesPtrOutput() CommitmentPlanPropertiesPtrOutput {
	return i.ToCommitmentPlanPropertiesPtrOutputWithContext(context.Background())
}

func (i CommitmentPlanPropertiesArgs) ToCommitmentPlanPropertiesPtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanPropertiesOutput).ToCommitmentPlanPropertiesPtrOutputWithContext(ctx)
}









type CommitmentPlanPropertiesPtrInput interface {
	pulumi.Input

	ToCommitmentPlanPropertiesPtrOutput() CommitmentPlanPropertiesPtrOutput
	ToCommitmentPlanPropertiesPtrOutputWithContext(context.Context) CommitmentPlanPropertiesPtrOutput
}

type commitmentPlanPropertiesPtrType CommitmentPlanPropertiesArgs

func CommitmentPlanPropertiesPtr(v *CommitmentPlanPropertiesArgs) CommitmentPlanPropertiesPtrInput {
	return (*commitmentPlanPropertiesPtrType)(v)
}

func (*commitmentPlanPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanProperties)(nil)).Elem()
}

func (i *commitmentPlanPropertiesPtrType) ToCommitmentPlanPropertiesPtrOutput() CommitmentPlanPropertiesPtrOutput {
	return i.ToCommitmentPlanPropertiesPtrOutputWithContext(context.Background())
}

func (i *commitmentPlanPropertiesPtrType) ToCommitmentPlanPropertiesPtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanPropertiesPtrOutput)
}

type CommitmentPlanPropertiesOutput struct{ *pulumi.OutputState }

func (CommitmentPlanPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPlanProperties)(nil)).Elem()
}

func (o CommitmentPlanPropertiesOutput) ToCommitmentPlanPropertiesOutput() CommitmentPlanPropertiesOutput {
	return o
}

func (o CommitmentPlanPropertiesOutput) ToCommitmentPlanPropertiesOutputWithContext(ctx context.Context) CommitmentPlanPropertiesOutput {
	return o
}

func (o CommitmentPlanPropertiesOutput) ToCommitmentPlanPropertiesPtrOutput() CommitmentPlanPropertiesPtrOutput {
	return o.ToCommitmentPlanPropertiesPtrOutputWithContext(context.Background())
}

func (o CommitmentPlanPropertiesOutput) ToCommitmentPlanPropertiesPtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommitmentPlanProperties) *CommitmentPlanProperties {
		return &v
	}).(CommitmentPlanPropertiesPtrOutput)
}

func (o CommitmentPlanPropertiesOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CommitmentPlanProperties) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

func (o CommitmentPlanPropertiesOutput) Current() CommitmentPeriodPtrOutput {
	return o.ApplyT(func(v CommitmentPlanProperties) *CommitmentPeriod { return v.Current }).(CommitmentPeriodPtrOutput)
}

func (o CommitmentPlanPropertiesOutput) HostingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommitmentPlanProperties) *string { return v.HostingModel }).(pulumi.StringPtrOutput)
}

func (o CommitmentPlanPropertiesOutput) Next() CommitmentPeriodPtrOutput {
	return o.ApplyT(func(v CommitmentPlanProperties) *CommitmentPeriod { return v.Next }).(CommitmentPeriodPtrOutput)
}

func (o CommitmentPlanPropertiesOutput) PlanType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommitmentPlanProperties) *string { return v.PlanType }).(pulumi.StringPtrOutput)
}

type CommitmentPlanPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CommitmentPlanPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanProperties)(nil)).Elem()
}

func (o CommitmentPlanPropertiesPtrOutput) ToCommitmentPlanPropertiesPtrOutput() CommitmentPlanPropertiesPtrOutput {
	return o
}

func (o CommitmentPlanPropertiesPtrOutput) ToCommitmentPlanPropertiesPtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesPtrOutput {
	return o
}

func (o CommitmentPlanPropertiesPtrOutput) Elem() CommitmentPlanPropertiesOutput {
	return o.ApplyT(func(v *CommitmentPlanProperties) CommitmentPlanProperties {
		if v != nil {
			return *v
		}
		var ret CommitmentPlanProperties
		return ret
	}).(CommitmentPlanPropertiesOutput)
}

func (o CommitmentPlanPropertiesPtrOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanProperties) *bool {
		if v == nil {
			return nil
		}
		return v.AutoRenew
	}).(pulumi.BoolPtrOutput)
}

func (o CommitmentPlanPropertiesPtrOutput) Current() CommitmentPeriodPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanProperties) *CommitmentPeriod {
		if v == nil {
			return nil
		}
		return v.Current
	}).(CommitmentPeriodPtrOutput)
}

func (o CommitmentPlanPropertiesPtrOutput) HostingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanProperties) *string {
		if v == nil {
			return nil
		}
		return v.HostingModel
	}).(pulumi.StringPtrOutput)
}

func (o CommitmentPlanPropertiesPtrOutput) Next() CommitmentPeriodPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanProperties) *CommitmentPeriod {
		if v == nil {
			return nil
		}
		return v.Next
	}).(CommitmentPeriodPtrOutput)
}

func (o CommitmentPlanPropertiesPtrOutput) PlanType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanProperties) *string {
		if v == nil {
			return nil
		}
		return v.PlanType
	}).(pulumi.StringPtrOutput)
}

type CommitmentPlanPropertiesResponse struct {
	AutoRenew    *bool                     `pulumi:"autoRenew"`
	Current      *CommitmentPeriodResponse `pulumi:"current"`
	HostingModel *string                   `pulumi:"hostingModel"`
	Last         CommitmentPeriodResponse  `pulumi:"last"`
	Next         *CommitmentPeriodResponse `pulumi:"next"`
	PlanType     *string                   `pulumi:"planType"`
}





type CommitmentPlanPropertiesResponseInput interface {
	pulumi.Input

	ToCommitmentPlanPropertiesResponseOutput() CommitmentPlanPropertiesResponseOutput
	ToCommitmentPlanPropertiesResponseOutputWithContext(context.Context) CommitmentPlanPropertiesResponseOutput
}

type CommitmentPlanPropertiesResponseArgs struct {
	AutoRenew    pulumi.BoolPtrInput              `pulumi:"autoRenew"`
	Current      CommitmentPeriodResponsePtrInput `pulumi:"current"`
	HostingModel pulumi.StringPtrInput            `pulumi:"hostingModel"`
	Last         CommitmentPeriodResponseInput    `pulumi:"last"`
	Next         CommitmentPeriodResponsePtrInput `pulumi:"next"`
	PlanType     pulumi.StringPtrInput            `pulumi:"planType"`
}

func (CommitmentPlanPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPlanPropertiesResponse)(nil)).Elem()
}

func (i CommitmentPlanPropertiesResponseArgs) ToCommitmentPlanPropertiesResponseOutput() CommitmentPlanPropertiesResponseOutput {
	return i.ToCommitmentPlanPropertiesResponseOutputWithContext(context.Background())
}

func (i CommitmentPlanPropertiesResponseArgs) ToCommitmentPlanPropertiesResponseOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanPropertiesResponseOutput)
}

func (i CommitmentPlanPropertiesResponseArgs) ToCommitmentPlanPropertiesResponsePtrOutput() CommitmentPlanPropertiesResponsePtrOutput {
	return i.ToCommitmentPlanPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CommitmentPlanPropertiesResponseArgs) ToCommitmentPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanPropertiesResponseOutput).ToCommitmentPlanPropertiesResponsePtrOutputWithContext(ctx)
}









type CommitmentPlanPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCommitmentPlanPropertiesResponsePtrOutput() CommitmentPlanPropertiesResponsePtrOutput
	ToCommitmentPlanPropertiesResponsePtrOutputWithContext(context.Context) CommitmentPlanPropertiesResponsePtrOutput
}

type commitmentPlanPropertiesResponsePtrType CommitmentPlanPropertiesResponseArgs

func CommitmentPlanPropertiesResponsePtr(v *CommitmentPlanPropertiesResponseArgs) CommitmentPlanPropertiesResponsePtrInput {
	return (*commitmentPlanPropertiesResponsePtrType)(v)
}

func (*commitmentPlanPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanPropertiesResponse)(nil)).Elem()
}

func (i *commitmentPlanPropertiesResponsePtrType) ToCommitmentPlanPropertiesResponsePtrOutput() CommitmentPlanPropertiesResponsePtrOutput {
	return i.ToCommitmentPlanPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *commitmentPlanPropertiesResponsePtrType) ToCommitmentPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanPropertiesResponsePtrOutput)
}

type CommitmentPlanPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CommitmentPlanPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPlanPropertiesResponse)(nil)).Elem()
}

func (o CommitmentPlanPropertiesResponseOutput) ToCommitmentPlanPropertiesResponseOutput() CommitmentPlanPropertiesResponseOutput {
	return o
}

func (o CommitmentPlanPropertiesResponseOutput) ToCommitmentPlanPropertiesResponseOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponseOutput {
	return o
}

func (o CommitmentPlanPropertiesResponseOutput) ToCommitmentPlanPropertiesResponsePtrOutput() CommitmentPlanPropertiesResponsePtrOutput {
	return o.ToCommitmentPlanPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CommitmentPlanPropertiesResponseOutput) ToCommitmentPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommitmentPlanPropertiesResponse) *CommitmentPlanPropertiesResponse {
		return &v
	}).(CommitmentPlanPropertiesResponsePtrOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) Current() CommitmentPeriodResponsePtrOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) *CommitmentPeriodResponse { return v.Current }).(CommitmentPeriodResponsePtrOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) HostingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) *string { return v.HostingModel }).(pulumi.StringPtrOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) Last() CommitmentPeriodResponseOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) CommitmentPeriodResponse { return v.Last }).(CommitmentPeriodResponseOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) Next() CommitmentPeriodResponsePtrOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) *CommitmentPeriodResponse { return v.Next }).(CommitmentPeriodResponsePtrOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) PlanType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) *string { return v.PlanType }).(pulumi.StringPtrOutput)
}

type CommitmentPlanPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CommitmentPlanPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanPropertiesResponse)(nil)).Elem()
}

func (o CommitmentPlanPropertiesResponsePtrOutput) ToCommitmentPlanPropertiesResponsePtrOutput() CommitmentPlanPropertiesResponsePtrOutput {
	return o
}

func (o CommitmentPlanPropertiesResponsePtrOutput) ToCommitmentPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponsePtrOutput {
	return o
}

func (o CommitmentPlanPropertiesResponsePtrOutput) Elem() CommitmentPlanPropertiesResponseOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) CommitmentPlanPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CommitmentPlanPropertiesResponse
		return ret
	}).(CommitmentPlanPropertiesResponseOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutoRenew
	}).(pulumi.BoolPtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) Current() CommitmentPeriodResponsePtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *CommitmentPeriodResponse {
		if v == nil {
			return nil
		}
		return v.Current
	}).(CommitmentPeriodResponsePtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) HostingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.HostingModel
	}).(pulumi.StringPtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) Last() CommitmentPeriodResponsePtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *CommitmentPeriodResponse {
		if v == nil {
			return nil
		}
		return &v.Last
	}).(CommitmentPeriodResponsePtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) Next() CommitmentPeriodResponsePtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *CommitmentPeriodResponse {
		if v == nil {
			return nil
		}
		return v.Next
	}).(CommitmentPeriodResponsePtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) PlanType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PlanType
	}).(pulumi.StringPtrOutput)
}

type CommitmentQuotaResponse struct {
	Quantity *float64 `pulumi:"quantity"`
	Unit     *string  `pulumi:"unit"`
}





type CommitmentQuotaResponseInput interface {
	pulumi.Input

	ToCommitmentQuotaResponseOutput() CommitmentQuotaResponseOutput
	ToCommitmentQuotaResponseOutputWithContext(context.Context) CommitmentQuotaResponseOutput
}

type CommitmentQuotaResponseArgs struct {
	Quantity pulumi.Float64PtrInput `pulumi:"quantity"`
	Unit     pulumi.StringPtrInput  `pulumi:"unit"`
}

func (CommitmentQuotaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentQuotaResponse)(nil)).Elem()
}

func (i CommitmentQuotaResponseArgs) ToCommitmentQuotaResponseOutput() CommitmentQuotaResponseOutput {
	return i.ToCommitmentQuotaResponseOutputWithContext(context.Background())
}

func (i CommitmentQuotaResponseArgs) ToCommitmentQuotaResponseOutputWithContext(ctx context.Context) CommitmentQuotaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentQuotaResponseOutput)
}

func (i CommitmentQuotaResponseArgs) ToCommitmentQuotaResponsePtrOutput() CommitmentQuotaResponsePtrOutput {
	return i.ToCommitmentQuotaResponsePtrOutputWithContext(context.Background())
}

func (i CommitmentQuotaResponseArgs) ToCommitmentQuotaResponsePtrOutputWithContext(ctx context.Context) CommitmentQuotaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentQuotaResponseOutput).ToCommitmentQuotaResponsePtrOutputWithContext(ctx)
}









type CommitmentQuotaResponsePtrInput interface {
	pulumi.Input

	ToCommitmentQuotaResponsePtrOutput() CommitmentQuotaResponsePtrOutput
	ToCommitmentQuotaResponsePtrOutputWithContext(context.Context) CommitmentQuotaResponsePtrOutput
}

type commitmentQuotaResponsePtrType CommitmentQuotaResponseArgs

func CommitmentQuotaResponsePtr(v *CommitmentQuotaResponseArgs) CommitmentQuotaResponsePtrInput {
	return (*commitmentQuotaResponsePtrType)(v)
}

func (*commitmentQuotaResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentQuotaResponse)(nil)).Elem()
}

func (i *commitmentQuotaResponsePtrType) ToCommitmentQuotaResponsePtrOutput() CommitmentQuotaResponsePtrOutput {
	return i.ToCommitmentQuotaResponsePtrOutputWithContext(context.Background())
}

func (i *commitmentQuotaResponsePtrType) ToCommitmentQuotaResponsePtrOutputWithContext(ctx context.Context) CommitmentQuotaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentQuotaResponsePtrOutput)
}

type CommitmentQuotaResponseOutput struct{ *pulumi.OutputState }

func (CommitmentQuotaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentQuotaResponse)(nil)).Elem()
}

func (o CommitmentQuotaResponseOutput) ToCommitmentQuotaResponseOutput() CommitmentQuotaResponseOutput {
	return o
}

func (o CommitmentQuotaResponseOutput) ToCommitmentQuotaResponseOutputWithContext(ctx context.Context) CommitmentQuotaResponseOutput {
	return o
}

func (o CommitmentQuotaResponseOutput) ToCommitmentQuotaResponsePtrOutput() CommitmentQuotaResponsePtrOutput {
	return o.ToCommitmentQuotaResponsePtrOutputWithContext(context.Background())
}

func (o CommitmentQuotaResponseOutput) ToCommitmentQuotaResponsePtrOutputWithContext(ctx context.Context) CommitmentQuotaResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommitmentQuotaResponse) *CommitmentQuotaResponse {
		return &v
	}).(CommitmentQuotaResponsePtrOutput)
}

func (o CommitmentQuotaResponseOutput) Quantity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CommitmentQuotaResponse) *float64 { return v.Quantity }).(pulumi.Float64PtrOutput)
}

func (o CommitmentQuotaResponseOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommitmentQuotaResponse) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

type CommitmentQuotaResponsePtrOutput struct{ *pulumi.OutputState }

func (CommitmentQuotaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentQuotaResponse)(nil)).Elem()
}

func (o CommitmentQuotaResponsePtrOutput) ToCommitmentQuotaResponsePtrOutput() CommitmentQuotaResponsePtrOutput {
	return o
}

func (o CommitmentQuotaResponsePtrOutput) ToCommitmentQuotaResponsePtrOutputWithContext(ctx context.Context) CommitmentQuotaResponsePtrOutput {
	return o
}

func (o CommitmentQuotaResponsePtrOutput) Elem() CommitmentQuotaResponseOutput {
	return o.ApplyT(func(v *CommitmentQuotaResponse) CommitmentQuotaResponse {
		if v != nil {
			return *v
		}
		var ret CommitmentQuotaResponse
		return ret
	}).(CommitmentQuotaResponseOutput)
}

func (o CommitmentQuotaResponsePtrOutput) Quantity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CommitmentQuotaResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Quantity
	}).(pulumi.Float64PtrOutput)
}

func (o CommitmentQuotaResponsePtrOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentQuotaResponse) *string {
		if v == nil {
			return nil
		}
		return v.Unit
	}).(pulumi.StringPtrOutput)
}

type DeploymentModel struct {
	Format  *string `pulumi:"format"`
	Name    *string `pulumi:"name"`
	Version *string `pulumi:"version"`
}





type DeploymentModelInput interface {
	pulumi.Input

	ToDeploymentModelOutput() DeploymentModelOutput
	ToDeploymentModelOutputWithContext(context.Context) DeploymentModelOutput
}

type DeploymentModelArgs struct {
	Format  pulumi.StringPtrInput `pulumi:"format"`
	Name    pulumi.StringPtrInput `pulumi:"name"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (DeploymentModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentModel)(nil)).Elem()
}

func (i DeploymentModelArgs) ToDeploymentModelOutput() DeploymentModelOutput {
	return i.ToDeploymentModelOutputWithContext(context.Background())
}

func (i DeploymentModelArgs) ToDeploymentModelOutputWithContext(ctx context.Context) DeploymentModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentModelOutput)
}

func (i DeploymentModelArgs) ToDeploymentModelPtrOutput() DeploymentModelPtrOutput {
	return i.ToDeploymentModelPtrOutputWithContext(context.Background())
}

func (i DeploymentModelArgs) ToDeploymentModelPtrOutputWithContext(ctx context.Context) DeploymentModelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentModelOutput).ToDeploymentModelPtrOutputWithContext(ctx)
}









type DeploymentModelPtrInput interface {
	pulumi.Input

	ToDeploymentModelPtrOutput() DeploymentModelPtrOutput
	ToDeploymentModelPtrOutputWithContext(context.Context) DeploymentModelPtrOutput
}

type deploymentModelPtrType DeploymentModelArgs

func DeploymentModelPtr(v *DeploymentModelArgs) DeploymentModelPtrInput {
	return (*deploymentModelPtrType)(v)
}

func (*deploymentModelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentModel)(nil)).Elem()
}

func (i *deploymentModelPtrType) ToDeploymentModelPtrOutput() DeploymentModelPtrOutput {
	return i.ToDeploymentModelPtrOutputWithContext(context.Background())
}

func (i *deploymentModelPtrType) ToDeploymentModelPtrOutputWithContext(ctx context.Context) DeploymentModelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentModelPtrOutput)
}

type DeploymentModelOutput struct{ *pulumi.OutputState }

func (DeploymentModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentModel)(nil)).Elem()
}

func (o DeploymentModelOutput) ToDeploymentModelOutput() DeploymentModelOutput {
	return o
}

func (o DeploymentModelOutput) ToDeploymentModelOutputWithContext(ctx context.Context) DeploymentModelOutput {
	return o
}

func (o DeploymentModelOutput) ToDeploymentModelPtrOutput() DeploymentModelPtrOutput {
	return o.ToDeploymentModelPtrOutputWithContext(context.Background())
}

func (o DeploymentModelOutput) ToDeploymentModelPtrOutputWithContext(ctx context.Context) DeploymentModelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentModel) *DeploymentModel {
		return &v
	}).(DeploymentModelPtrOutput)
}

func (o DeploymentModelOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentModel) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o DeploymentModelOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentModel) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DeploymentModelOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentModel) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type DeploymentModelPtrOutput struct{ *pulumi.OutputState }

func (DeploymentModelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentModel)(nil)).Elem()
}

func (o DeploymentModelPtrOutput) ToDeploymentModelPtrOutput() DeploymentModelPtrOutput {
	return o
}

func (o DeploymentModelPtrOutput) ToDeploymentModelPtrOutputWithContext(ctx context.Context) DeploymentModelPtrOutput {
	return o
}

func (o DeploymentModelPtrOutput) Elem() DeploymentModelOutput {
	return o.ApplyT(func(v *DeploymentModel) DeploymentModel {
		if v != nil {
			return *v
		}
		var ret DeploymentModel
		return ret
	}).(DeploymentModelOutput)
}

func (o DeploymentModelPtrOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentModel) *string {
		if v == nil {
			return nil
		}
		return v.Format
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentModelPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentModel) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentModelPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentModel) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type DeploymentModelResponse struct {
	Format  *string `pulumi:"format"`
	Name    *string `pulumi:"name"`
	Version *string `pulumi:"version"`
}





type DeploymentModelResponseInput interface {
	pulumi.Input

	ToDeploymentModelResponseOutput() DeploymentModelResponseOutput
	ToDeploymentModelResponseOutputWithContext(context.Context) DeploymentModelResponseOutput
}

type DeploymentModelResponseArgs struct {
	Format  pulumi.StringPtrInput `pulumi:"format"`
	Name    pulumi.StringPtrInput `pulumi:"name"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (DeploymentModelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentModelResponse)(nil)).Elem()
}

func (i DeploymentModelResponseArgs) ToDeploymentModelResponseOutput() DeploymentModelResponseOutput {
	return i.ToDeploymentModelResponseOutputWithContext(context.Background())
}

func (i DeploymentModelResponseArgs) ToDeploymentModelResponseOutputWithContext(ctx context.Context) DeploymentModelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentModelResponseOutput)
}

func (i DeploymentModelResponseArgs) ToDeploymentModelResponsePtrOutput() DeploymentModelResponsePtrOutput {
	return i.ToDeploymentModelResponsePtrOutputWithContext(context.Background())
}

func (i DeploymentModelResponseArgs) ToDeploymentModelResponsePtrOutputWithContext(ctx context.Context) DeploymentModelResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentModelResponseOutput).ToDeploymentModelResponsePtrOutputWithContext(ctx)
}









type DeploymentModelResponsePtrInput interface {
	pulumi.Input

	ToDeploymentModelResponsePtrOutput() DeploymentModelResponsePtrOutput
	ToDeploymentModelResponsePtrOutputWithContext(context.Context) DeploymentModelResponsePtrOutput
}

type deploymentModelResponsePtrType DeploymentModelResponseArgs

func DeploymentModelResponsePtr(v *DeploymentModelResponseArgs) DeploymentModelResponsePtrInput {
	return (*deploymentModelResponsePtrType)(v)
}

func (*deploymentModelResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentModelResponse)(nil)).Elem()
}

func (i *deploymentModelResponsePtrType) ToDeploymentModelResponsePtrOutput() DeploymentModelResponsePtrOutput {
	return i.ToDeploymentModelResponsePtrOutputWithContext(context.Background())
}

func (i *deploymentModelResponsePtrType) ToDeploymentModelResponsePtrOutputWithContext(ctx context.Context) DeploymentModelResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentModelResponsePtrOutput)
}

type DeploymentModelResponseOutput struct{ *pulumi.OutputState }

func (DeploymentModelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentModelResponse)(nil)).Elem()
}

func (o DeploymentModelResponseOutput) ToDeploymentModelResponseOutput() DeploymentModelResponseOutput {
	return o
}

func (o DeploymentModelResponseOutput) ToDeploymentModelResponseOutputWithContext(ctx context.Context) DeploymentModelResponseOutput {
	return o
}

func (o DeploymentModelResponseOutput) ToDeploymentModelResponsePtrOutput() DeploymentModelResponsePtrOutput {
	return o.ToDeploymentModelResponsePtrOutputWithContext(context.Background())
}

func (o DeploymentModelResponseOutput) ToDeploymentModelResponsePtrOutputWithContext(ctx context.Context) DeploymentModelResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentModelResponse) *DeploymentModelResponse {
		return &v
	}).(DeploymentModelResponsePtrOutput)
}

func (o DeploymentModelResponseOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentModelResponse) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o DeploymentModelResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentModelResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DeploymentModelResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentModelResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type DeploymentModelResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentModelResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentModelResponse)(nil)).Elem()
}

func (o DeploymentModelResponsePtrOutput) ToDeploymentModelResponsePtrOutput() DeploymentModelResponsePtrOutput {
	return o
}

func (o DeploymentModelResponsePtrOutput) ToDeploymentModelResponsePtrOutputWithContext(ctx context.Context) DeploymentModelResponsePtrOutput {
	return o
}

func (o DeploymentModelResponsePtrOutput) Elem() DeploymentModelResponseOutput {
	return o.ApplyT(func(v *DeploymentModelResponse) DeploymentModelResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentModelResponse
		return ret
	}).(DeploymentModelResponseOutput)
}

func (o DeploymentModelResponsePtrOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentModelResponse) *string {
		if v == nil {
			return nil
		}
		return v.Format
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentModelResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentModelResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentModelResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentModelResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type DeploymentProperties struct {
	Model         *DeploymentModel         `pulumi:"model"`
	ScaleSettings *DeploymentScaleSettings `pulumi:"scaleSettings"`
}





type DeploymentPropertiesInput interface {
	pulumi.Input

	ToDeploymentPropertiesOutput() DeploymentPropertiesOutput
	ToDeploymentPropertiesOutputWithContext(context.Context) DeploymentPropertiesOutput
}

type DeploymentPropertiesArgs struct {
	Model         DeploymentModelPtrInput         `pulumi:"model"`
	ScaleSettings DeploymentScaleSettingsPtrInput `pulumi:"scaleSettings"`
}

func (DeploymentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentProperties)(nil)).Elem()
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesOutput() DeploymentPropertiesOutput {
	return i.ToDeploymentPropertiesOutputWithContext(context.Background())
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesOutputWithContext(ctx context.Context) DeploymentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesOutput)
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return i.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesOutput).ToDeploymentPropertiesPtrOutputWithContext(ctx)
}









type DeploymentPropertiesPtrInput interface {
	pulumi.Input

	ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput
	ToDeploymentPropertiesPtrOutputWithContext(context.Context) DeploymentPropertiesPtrOutput
}

type deploymentPropertiesPtrType DeploymentPropertiesArgs

func DeploymentPropertiesPtr(v *DeploymentPropertiesArgs) DeploymentPropertiesPtrInput {
	return (*deploymentPropertiesPtrType)(v)
}

func (*deploymentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentProperties)(nil)).Elem()
}

func (i *deploymentPropertiesPtrType) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return i.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (i *deploymentPropertiesPtrType) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesPtrOutput)
}

type DeploymentPropertiesOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentProperties)(nil)).Elem()
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesOutput() DeploymentPropertiesOutput {
	return o
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesOutputWithContext(ctx context.Context) DeploymentPropertiesOutput {
	return o
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return o.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentProperties) *DeploymentProperties {
		return &v
	}).(DeploymentPropertiesPtrOutput)
}

func (o DeploymentPropertiesOutput) Model() DeploymentModelPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *DeploymentModel { return v.Model }).(DeploymentModelPtrOutput)
}

func (o DeploymentPropertiesOutput) ScaleSettings() DeploymentScaleSettingsPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *DeploymentScaleSettings { return v.ScaleSettings }).(DeploymentScaleSettingsPtrOutput)
}

type DeploymentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentProperties)(nil)).Elem()
}

func (o DeploymentPropertiesPtrOutput) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return o
}

func (o DeploymentPropertiesPtrOutput) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return o
}

func (o DeploymentPropertiesPtrOutput) Elem() DeploymentPropertiesOutput {
	return o.ApplyT(func(v *DeploymentProperties) DeploymentProperties {
		if v != nil {
			return *v
		}
		var ret DeploymentProperties
		return ret
	}).(DeploymentPropertiesOutput)
}

func (o DeploymentPropertiesPtrOutput) Model() DeploymentModelPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *DeploymentModel {
		if v == nil {
			return nil
		}
		return v.Model
	}).(DeploymentModelPtrOutput)
}

func (o DeploymentPropertiesPtrOutput) ScaleSettings() DeploymentScaleSettingsPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *DeploymentScaleSettings {
		if v == nil {
			return nil
		}
		return v.ScaleSettings
	}).(DeploymentScaleSettingsPtrOutput)
}

type DeploymentPropertiesResponse struct {
	Model             *DeploymentModelResponse         `pulumi:"model"`
	ProvisioningState string                           `pulumi:"provisioningState"`
	ScaleSettings     *DeploymentScaleSettingsResponse `pulumi:"scaleSettings"`
}





type DeploymentPropertiesResponseInput interface {
	pulumi.Input

	ToDeploymentPropertiesResponseOutput() DeploymentPropertiesResponseOutput
	ToDeploymentPropertiesResponseOutputWithContext(context.Context) DeploymentPropertiesResponseOutput
}

type DeploymentPropertiesResponseArgs struct {
	Model             DeploymentModelResponsePtrInput         `pulumi:"model"`
	ProvisioningState pulumi.StringInput                      `pulumi:"provisioningState"`
	ScaleSettings     DeploymentScaleSettingsResponsePtrInput `pulumi:"scaleSettings"`
}

func (DeploymentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentPropertiesResponse)(nil)).Elem()
}

func (i DeploymentPropertiesResponseArgs) ToDeploymentPropertiesResponseOutput() DeploymentPropertiesResponseOutput {
	return i.ToDeploymentPropertiesResponseOutputWithContext(context.Background())
}

func (i DeploymentPropertiesResponseArgs) ToDeploymentPropertiesResponseOutputWithContext(ctx context.Context) DeploymentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesResponseOutput)
}

func (i DeploymentPropertiesResponseArgs) ToDeploymentPropertiesResponsePtrOutput() DeploymentPropertiesResponsePtrOutput {
	return i.ToDeploymentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i DeploymentPropertiesResponseArgs) ToDeploymentPropertiesResponsePtrOutputWithContext(ctx context.Context) DeploymentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesResponseOutput).ToDeploymentPropertiesResponsePtrOutputWithContext(ctx)
}









type DeploymentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToDeploymentPropertiesResponsePtrOutput() DeploymentPropertiesResponsePtrOutput
	ToDeploymentPropertiesResponsePtrOutputWithContext(context.Context) DeploymentPropertiesResponsePtrOutput
}

type deploymentPropertiesResponsePtrType DeploymentPropertiesResponseArgs

func DeploymentPropertiesResponsePtr(v *DeploymentPropertiesResponseArgs) DeploymentPropertiesResponsePtrInput {
	return (*deploymentPropertiesResponsePtrType)(v)
}

func (*deploymentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentPropertiesResponse)(nil)).Elem()
}

func (i *deploymentPropertiesResponsePtrType) ToDeploymentPropertiesResponsePtrOutput() DeploymentPropertiesResponsePtrOutput {
	return i.ToDeploymentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *deploymentPropertiesResponsePtrType) ToDeploymentPropertiesResponsePtrOutputWithContext(ctx context.Context) DeploymentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesResponsePtrOutput)
}

type DeploymentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentPropertiesResponse)(nil)).Elem()
}

func (o DeploymentPropertiesResponseOutput) ToDeploymentPropertiesResponseOutput() DeploymentPropertiesResponseOutput {
	return o
}

func (o DeploymentPropertiesResponseOutput) ToDeploymentPropertiesResponseOutputWithContext(ctx context.Context) DeploymentPropertiesResponseOutput {
	return o
}

func (o DeploymentPropertiesResponseOutput) ToDeploymentPropertiesResponsePtrOutput() DeploymentPropertiesResponsePtrOutput {
	return o.ToDeploymentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o DeploymentPropertiesResponseOutput) ToDeploymentPropertiesResponsePtrOutputWithContext(ctx context.Context) DeploymentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentPropertiesResponse) *DeploymentPropertiesResponse {
		return &v
	}).(DeploymentPropertiesResponsePtrOutput)
}

func (o DeploymentPropertiesResponseOutput) Model() DeploymentModelResponsePtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesResponse) *DeploymentModelResponse { return v.Model }).(DeploymentModelResponsePtrOutput)
}

func (o DeploymentPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DeploymentPropertiesResponseOutput) ScaleSettings() DeploymentScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesResponse) *DeploymentScaleSettingsResponse { return v.ScaleSettings }).(DeploymentScaleSettingsResponsePtrOutput)
}

type DeploymentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentPropertiesResponse)(nil)).Elem()
}

func (o DeploymentPropertiesResponsePtrOutput) ToDeploymentPropertiesResponsePtrOutput() DeploymentPropertiesResponsePtrOutput {
	return o
}

func (o DeploymentPropertiesResponsePtrOutput) ToDeploymentPropertiesResponsePtrOutputWithContext(ctx context.Context) DeploymentPropertiesResponsePtrOutput {
	return o
}

func (o DeploymentPropertiesResponsePtrOutput) Elem() DeploymentPropertiesResponseOutput {
	return o.ApplyT(func(v *DeploymentPropertiesResponse) DeploymentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentPropertiesResponse
		return ret
	}).(DeploymentPropertiesResponseOutput)
}

func (o DeploymentPropertiesResponsePtrOutput) Model() DeploymentModelResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesResponse) *DeploymentModelResponse {
		if v == nil {
			return nil
		}
		return v.Model
	}).(DeploymentModelResponsePtrOutput)
}

func (o DeploymentPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesResponsePtrOutput) ScaleSettings() DeploymentScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesResponse) *DeploymentScaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ScaleSettings
	}).(DeploymentScaleSettingsResponsePtrOutput)
}

type DeploymentScaleSettings struct {
	Capacity  *int    `pulumi:"capacity"`
	ScaleType *string `pulumi:"scaleType"`
}





type DeploymentScaleSettingsInput interface {
	pulumi.Input

	ToDeploymentScaleSettingsOutput() DeploymentScaleSettingsOutput
	ToDeploymentScaleSettingsOutputWithContext(context.Context) DeploymentScaleSettingsOutput
}

type DeploymentScaleSettingsArgs struct {
	Capacity  pulumi.IntPtrInput    `pulumi:"capacity"`
	ScaleType pulumi.StringPtrInput `pulumi:"scaleType"`
}

func (DeploymentScaleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentScaleSettings)(nil)).Elem()
}

func (i DeploymentScaleSettingsArgs) ToDeploymentScaleSettingsOutput() DeploymentScaleSettingsOutput {
	return i.ToDeploymentScaleSettingsOutputWithContext(context.Background())
}

func (i DeploymentScaleSettingsArgs) ToDeploymentScaleSettingsOutputWithContext(ctx context.Context) DeploymentScaleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentScaleSettingsOutput)
}

func (i DeploymentScaleSettingsArgs) ToDeploymentScaleSettingsPtrOutput() DeploymentScaleSettingsPtrOutput {
	return i.ToDeploymentScaleSettingsPtrOutputWithContext(context.Background())
}

func (i DeploymentScaleSettingsArgs) ToDeploymentScaleSettingsPtrOutputWithContext(ctx context.Context) DeploymentScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentScaleSettingsOutput).ToDeploymentScaleSettingsPtrOutputWithContext(ctx)
}









type DeploymentScaleSettingsPtrInput interface {
	pulumi.Input

	ToDeploymentScaleSettingsPtrOutput() DeploymentScaleSettingsPtrOutput
	ToDeploymentScaleSettingsPtrOutputWithContext(context.Context) DeploymentScaleSettingsPtrOutput
}

type deploymentScaleSettingsPtrType DeploymentScaleSettingsArgs

func DeploymentScaleSettingsPtr(v *DeploymentScaleSettingsArgs) DeploymentScaleSettingsPtrInput {
	return (*deploymentScaleSettingsPtrType)(v)
}

func (*deploymentScaleSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentScaleSettings)(nil)).Elem()
}

func (i *deploymentScaleSettingsPtrType) ToDeploymentScaleSettingsPtrOutput() DeploymentScaleSettingsPtrOutput {
	return i.ToDeploymentScaleSettingsPtrOutputWithContext(context.Background())
}

func (i *deploymentScaleSettingsPtrType) ToDeploymentScaleSettingsPtrOutputWithContext(ctx context.Context) DeploymentScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentScaleSettingsPtrOutput)
}

type DeploymentScaleSettingsOutput struct{ *pulumi.OutputState }

func (DeploymentScaleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentScaleSettings)(nil)).Elem()
}

func (o DeploymentScaleSettingsOutput) ToDeploymentScaleSettingsOutput() DeploymentScaleSettingsOutput {
	return o
}

func (o DeploymentScaleSettingsOutput) ToDeploymentScaleSettingsOutputWithContext(ctx context.Context) DeploymentScaleSettingsOutput {
	return o
}

func (o DeploymentScaleSettingsOutput) ToDeploymentScaleSettingsPtrOutput() DeploymentScaleSettingsPtrOutput {
	return o.ToDeploymentScaleSettingsPtrOutputWithContext(context.Background())
}

func (o DeploymentScaleSettingsOutput) ToDeploymentScaleSettingsPtrOutputWithContext(ctx context.Context) DeploymentScaleSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentScaleSettings) *DeploymentScaleSettings {
		return &v
	}).(DeploymentScaleSettingsPtrOutput)
}

func (o DeploymentScaleSettingsOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentScaleSettings) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o DeploymentScaleSettingsOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentScaleSettings) *string { return v.ScaleType }).(pulumi.StringPtrOutput)
}

type DeploymentScaleSettingsPtrOutput struct{ *pulumi.OutputState }

func (DeploymentScaleSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentScaleSettings)(nil)).Elem()
}

func (o DeploymentScaleSettingsPtrOutput) ToDeploymentScaleSettingsPtrOutput() DeploymentScaleSettingsPtrOutput {
	return o
}

func (o DeploymentScaleSettingsPtrOutput) ToDeploymentScaleSettingsPtrOutputWithContext(ctx context.Context) DeploymentScaleSettingsPtrOutput {
	return o
}

func (o DeploymentScaleSettingsPtrOutput) Elem() DeploymentScaleSettingsOutput {
	return o.ApplyT(func(v *DeploymentScaleSettings) DeploymentScaleSettings {
		if v != nil {
			return *v
		}
		var ret DeploymentScaleSettings
		return ret
	}).(DeploymentScaleSettingsOutput)
}

func (o DeploymentScaleSettingsPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentScaleSettings) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o DeploymentScaleSettingsPtrOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentScaleSettings) *string {
		if v == nil {
			return nil
		}
		return v.ScaleType
	}).(pulumi.StringPtrOutput)
}

type DeploymentScaleSettingsResponse struct {
	Capacity  *int    `pulumi:"capacity"`
	ScaleType *string `pulumi:"scaleType"`
}





type DeploymentScaleSettingsResponseInput interface {
	pulumi.Input

	ToDeploymentScaleSettingsResponseOutput() DeploymentScaleSettingsResponseOutput
	ToDeploymentScaleSettingsResponseOutputWithContext(context.Context) DeploymentScaleSettingsResponseOutput
}

type DeploymentScaleSettingsResponseArgs struct {
	Capacity  pulumi.IntPtrInput    `pulumi:"capacity"`
	ScaleType pulumi.StringPtrInput `pulumi:"scaleType"`
}

func (DeploymentScaleSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentScaleSettingsResponse)(nil)).Elem()
}

func (i DeploymentScaleSettingsResponseArgs) ToDeploymentScaleSettingsResponseOutput() DeploymentScaleSettingsResponseOutput {
	return i.ToDeploymentScaleSettingsResponseOutputWithContext(context.Background())
}

func (i DeploymentScaleSettingsResponseArgs) ToDeploymentScaleSettingsResponseOutputWithContext(ctx context.Context) DeploymentScaleSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentScaleSettingsResponseOutput)
}

func (i DeploymentScaleSettingsResponseArgs) ToDeploymentScaleSettingsResponsePtrOutput() DeploymentScaleSettingsResponsePtrOutput {
	return i.ToDeploymentScaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (i DeploymentScaleSettingsResponseArgs) ToDeploymentScaleSettingsResponsePtrOutputWithContext(ctx context.Context) DeploymentScaleSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentScaleSettingsResponseOutput).ToDeploymentScaleSettingsResponsePtrOutputWithContext(ctx)
}









type DeploymentScaleSettingsResponsePtrInput interface {
	pulumi.Input

	ToDeploymentScaleSettingsResponsePtrOutput() DeploymentScaleSettingsResponsePtrOutput
	ToDeploymentScaleSettingsResponsePtrOutputWithContext(context.Context) DeploymentScaleSettingsResponsePtrOutput
}

type deploymentScaleSettingsResponsePtrType DeploymentScaleSettingsResponseArgs

func DeploymentScaleSettingsResponsePtr(v *DeploymentScaleSettingsResponseArgs) DeploymentScaleSettingsResponsePtrInput {
	return (*deploymentScaleSettingsResponsePtrType)(v)
}

func (*deploymentScaleSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentScaleSettingsResponse)(nil)).Elem()
}

func (i *deploymentScaleSettingsResponsePtrType) ToDeploymentScaleSettingsResponsePtrOutput() DeploymentScaleSettingsResponsePtrOutput {
	return i.ToDeploymentScaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *deploymentScaleSettingsResponsePtrType) ToDeploymentScaleSettingsResponsePtrOutputWithContext(ctx context.Context) DeploymentScaleSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentScaleSettingsResponsePtrOutput)
}

type DeploymentScaleSettingsResponseOutput struct{ *pulumi.OutputState }

func (DeploymentScaleSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentScaleSettingsResponse)(nil)).Elem()
}

func (o DeploymentScaleSettingsResponseOutput) ToDeploymentScaleSettingsResponseOutput() DeploymentScaleSettingsResponseOutput {
	return o
}

func (o DeploymentScaleSettingsResponseOutput) ToDeploymentScaleSettingsResponseOutputWithContext(ctx context.Context) DeploymentScaleSettingsResponseOutput {
	return o
}

func (o DeploymentScaleSettingsResponseOutput) ToDeploymentScaleSettingsResponsePtrOutput() DeploymentScaleSettingsResponsePtrOutput {
	return o.ToDeploymentScaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (o DeploymentScaleSettingsResponseOutput) ToDeploymentScaleSettingsResponsePtrOutputWithContext(ctx context.Context) DeploymentScaleSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentScaleSettingsResponse) *DeploymentScaleSettingsResponse {
		return &v
	}).(DeploymentScaleSettingsResponsePtrOutput)
}

func (o DeploymentScaleSettingsResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentScaleSettingsResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o DeploymentScaleSettingsResponseOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentScaleSettingsResponse) *string { return v.ScaleType }).(pulumi.StringPtrOutput)
}

type DeploymentScaleSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentScaleSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentScaleSettingsResponse)(nil)).Elem()
}

func (o DeploymentScaleSettingsResponsePtrOutput) ToDeploymentScaleSettingsResponsePtrOutput() DeploymentScaleSettingsResponsePtrOutput {
	return o
}

func (o DeploymentScaleSettingsResponsePtrOutput) ToDeploymentScaleSettingsResponsePtrOutputWithContext(ctx context.Context) DeploymentScaleSettingsResponsePtrOutput {
	return o
}

func (o DeploymentScaleSettingsResponsePtrOutput) Elem() DeploymentScaleSettingsResponseOutput {
	return o.ApplyT(func(v *DeploymentScaleSettingsResponse) DeploymentScaleSettingsResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentScaleSettingsResponse
		return ret
	}).(DeploymentScaleSettingsResponseOutput)
}

func (o DeploymentScaleSettingsResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentScaleSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o DeploymentScaleSettingsResponsePtrOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentScaleSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScaleType
	}).(pulumi.StringPtrOutput)
}

type Encryption struct {
	KeySource          *string             `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeySource          pulumi.StringPtrInput      `pulumi:"keySource"`
	KeyVaultProperties KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v Encryption) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *Encryption) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

type EncryptionResponse struct {
	KeySource          *string                     `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
}





type EncryptionResponseInput interface {
	pulumi.Input

	ToEncryptionResponseOutput() EncryptionResponseOutput
	ToEncryptionResponseOutputWithContext(context.Context) EncryptionResponseOutput
}

type EncryptionResponseArgs struct {
	KeySource          pulumi.StringPtrInput              `pulumi:"keySource"`
	KeyVaultProperties KeyVaultPropertiesResponsePtrInput `pulumi:"keyVaultProperties"`
}

func (EncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return i.ToEncryptionResponseOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput)
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput).ToEncryptionResponsePtrOutputWithContext(ctx)
}









type EncryptionResponsePtrInput interface {
	pulumi.Input

	ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput
	ToEncryptionResponsePtrOutputWithContext(context.Context) EncryptionResponsePtrOutput
}

type encryptionResponsePtrType EncryptionResponseArgs

func EncryptionResponsePtr(v *EncryptionResponseArgs) EncryptionResponsePtrInput {
	return (*encryptionResponsePtrType)(v)
}

func (*encryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponsePtrOutput)
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionResponse) *EncryptionResponse {
		return &v
	}).(EncryptionResponsePtrOutput)
}

func (o EncryptionResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionResponse
		return ret
	}).(EncryptionResponseOutput)
}

func (o EncryptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

type Identity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                   `pulumi:"principalId"`
	TenantId               pulumi.StringInput                   `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]UserAssignedIdentityResponse { return v.UserAssignedIdentities }).(UserAssignedIdentityResponseMapOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type IpRule struct {
	Value string `pulumi:"value"`
}





type IpRuleInput interface {
	pulumi.Input

	ToIpRuleOutput() IpRuleOutput
	ToIpRuleOutputWithContext(context.Context) IpRuleOutput
}

type IpRuleArgs struct {
	Value pulumi.StringInput `pulumi:"value"`
}

func (IpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRule)(nil)).Elem()
}

func (i IpRuleArgs) ToIpRuleOutput() IpRuleOutput {
	return i.ToIpRuleOutputWithContext(context.Background())
}

func (i IpRuleArgs) ToIpRuleOutputWithContext(ctx context.Context) IpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleOutput)
}





type IpRuleArrayInput interface {
	pulumi.Input

	ToIpRuleArrayOutput() IpRuleArrayOutput
	ToIpRuleArrayOutputWithContext(context.Context) IpRuleArrayOutput
}

type IpRuleArray []IpRuleInput

func (IpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRule)(nil)).Elem()
}

func (i IpRuleArray) ToIpRuleArrayOutput() IpRuleArrayOutput {
	return i.ToIpRuleArrayOutputWithContext(context.Background())
}

func (i IpRuleArray) ToIpRuleArrayOutputWithContext(ctx context.Context) IpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleArrayOutput)
}

type IpRuleOutput struct{ *pulumi.OutputState }

func (IpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRule)(nil)).Elem()
}

func (o IpRuleOutput) ToIpRuleOutput() IpRuleOutput {
	return o
}

func (o IpRuleOutput) ToIpRuleOutputWithContext(ctx context.Context) IpRuleOutput {
	return o
}

func (o IpRuleOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v IpRule) string { return v.Value }).(pulumi.StringOutput)
}

type IpRuleArrayOutput struct{ *pulumi.OutputState }

func (IpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRule)(nil)).Elem()
}

func (o IpRuleArrayOutput) ToIpRuleArrayOutput() IpRuleArrayOutput {
	return o
}

func (o IpRuleArrayOutput) ToIpRuleArrayOutputWithContext(ctx context.Context) IpRuleArrayOutput {
	return o
}

func (o IpRuleArrayOutput) Index(i pulumi.IntInput) IpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpRule {
		return vs[0].([]IpRule)[vs[1].(int)]
	}).(IpRuleOutput)
}

type IpRuleResponse struct {
	Value string `pulumi:"value"`
}





type IpRuleResponseInput interface {
	pulumi.Input

	ToIpRuleResponseOutput() IpRuleResponseOutput
	ToIpRuleResponseOutputWithContext(context.Context) IpRuleResponseOutput
}

type IpRuleResponseArgs struct {
	Value pulumi.StringInput `pulumi:"value"`
}

func (IpRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRuleResponse)(nil)).Elem()
}

func (i IpRuleResponseArgs) ToIpRuleResponseOutput() IpRuleResponseOutput {
	return i.ToIpRuleResponseOutputWithContext(context.Background())
}

func (i IpRuleResponseArgs) ToIpRuleResponseOutputWithContext(ctx context.Context) IpRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleResponseOutput)
}





type IpRuleResponseArrayInput interface {
	pulumi.Input

	ToIpRuleResponseArrayOutput() IpRuleResponseArrayOutput
	ToIpRuleResponseArrayOutputWithContext(context.Context) IpRuleResponseArrayOutput
}

type IpRuleResponseArray []IpRuleResponseInput

func (IpRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRuleResponse)(nil)).Elem()
}

func (i IpRuleResponseArray) ToIpRuleResponseArrayOutput() IpRuleResponseArrayOutput {
	return i.ToIpRuleResponseArrayOutputWithContext(context.Background())
}

func (i IpRuleResponseArray) ToIpRuleResponseArrayOutputWithContext(ctx context.Context) IpRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleResponseArrayOutput)
}

type IpRuleResponseOutput struct{ *pulumi.OutputState }

func (IpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRuleResponse)(nil)).Elem()
}

func (o IpRuleResponseOutput) ToIpRuleResponseOutput() IpRuleResponseOutput {
	return o
}

func (o IpRuleResponseOutput) ToIpRuleResponseOutputWithContext(ctx context.Context) IpRuleResponseOutput {
	return o
}

func (o IpRuleResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v IpRuleResponse) string { return v.Value }).(pulumi.StringOutput)
}

type IpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRuleResponse)(nil)).Elem()
}

func (o IpRuleResponseArrayOutput) ToIpRuleResponseArrayOutput() IpRuleResponseArrayOutput {
	return o
}

func (o IpRuleResponseArrayOutput) ToIpRuleResponseArrayOutputWithContext(ctx context.Context) IpRuleResponseArrayOutput {
	return o
}

func (o IpRuleResponseArrayOutput) Index(i pulumi.IntInput) IpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpRuleResponse {
		return vs[0].([]IpRuleResponse)[vs[1].(int)]
	}).(IpRuleResponseOutput)
}

type KeyVaultProperties struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	KeyName          *string `pulumi:"keyName"`
	KeyVaultUri      *string `pulumi:"keyVaultUri"`
	KeyVersion       *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	IdentityClientId pulumi.StringPtrInput `pulumi:"identityClientId"`
	KeyName          pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri      pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion       pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.IdentityClientId
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	KeyName          *string `pulumi:"keyName"`
	KeyVaultUri      *string `pulumi:"keyVaultUri"`
	KeyVersion       *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	IdentityClientId pulumi.StringPtrInput `pulumi:"identityClientId"`
	KeyName          pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri      pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion       pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (KeyVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return i.ToKeyVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput)
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput).ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type KeyVaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput
	ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Context) KeyVaultPropertiesResponsePtrOutput
}

type keyVaultPropertiesResponsePtrType KeyVaultPropertiesResponseArgs

func KeyVaultPropertiesResponsePtr(v *KeyVaultPropertiesResponseArgs) KeyVaultPropertiesResponsePtrInput {
	return (*keyVaultPropertiesResponsePtrType)(v)
}

func (*keyVaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponsePtrOutput)
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultPropertiesResponse) *KeyVaultPropertiesResponse {
		return &v
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IdentityClientId
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type NetworkRuleSet struct {
	DefaultAction       *string              `pulumi:"defaultAction"`
	IpRules             []IpRule             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	DefaultAction       pulumi.StringPtrInput        `pulumi:"defaultAction"`
	IpRules             IpRuleArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSet) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetOutput) IpRules() IpRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IpRule { return v.IpRules }).(IpRuleArrayOutput)
}

func (o NetworkRuleSetOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []VirtualNetworkRule { return v.VirtualNetworkRules }).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IpRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IpRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IpRuleArrayOutput)
}

func (o NetworkRuleSetPtrOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []VirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	DefaultAction       *string                      `pulumi:"defaultAction"`
	IpRules             []IpRuleResponse             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}





type NetworkRuleSetResponseInput interface {
	pulumi.Input

	ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput
	ToNetworkRuleSetResponseOutputWithContext(context.Context) NetworkRuleSetResponseOutput
}

type NetworkRuleSetResponseArgs struct {
	DefaultAction       pulumi.StringPtrInput                `pulumi:"defaultAction"`
	IpRules             IpRuleResponseArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleResponseArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return i.ToNetworkRuleSetResponseOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput)
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput).ToNetworkRuleSetResponsePtrOutputWithContext(ctx)
}









type NetworkRuleSetResponsePtrInput interface {
	pulumi.Input

	ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput
	ToNetworkRuleSetResponsePtrOutputWithContext(context.Context) NetworkRuleSetResponsePtrOutput
}

type networkRuleSetResponsePtrType NetworkRuleSetResponseArgs

func NetworkRuleSetResponsePtr(v *NetworkRuleSetResponseArgs) NetworkRuleSetResponsePtrInput {
	return (*networkRuleSetResponsePtrType)(v)
}

func (*networkRuleSetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponsePtrOutput)
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSetResponse) *NetworkRuleSetResponse {
		return &v
	}).(NetworkRuleSetResponsePtrOutput)
}

func (o NetworkRuleSetResponseOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IpRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IpRuleResponse { return v.IpRules }).(IpRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) IpRules() IpRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []IpRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IpRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponsePtrOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []VirtualNetworkRuleResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleResponseArrayOutput)
}

type PrivateEndpointConnectionProperties struct {
	GroupIds                          []string                          `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	GroupIds                          pulumi.StringArrayInput                `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return i.ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput)
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput).ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput
	ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPtrOutput
}

type privateEndpointConnectionPropertiesPtrType PrivateEndpointConnectionPropertiesArgs

func PrivateEndpointConnectionPropertiesPtr(v *PrivateEndpointConnectionPropertiesArgs) PrivateEndpointConnectionPropertiesPtrInput {
	return (*privateEndpointConnectionPropertiesPtrType)(v)
}

func (*privateEndpointConnectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPtrOutput)
}

type PrivateEndpointConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionProperties {
		return &v
	}).(PrivateEndpointConnectionPropertiesPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateEndpointConnectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) Elem() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) PrivateEndpointConnectionProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionProperties
		return ret
	}).(PrivateEndpointConnectionPropertiesOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return &v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	GroupIds                          []string                                  `pulumi:"groupIds"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput
	ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponseOutput
}

type PrivateEndpointConnectionPropertiesResponseArgs struct {
	GroupIds                          pulumi.StringArrayInput                        `pulumi:"groupIds"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                             `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput).ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput
	ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput
}

type privateEndpointConnectionPropertiesResponsePtrType PrivateEndpointConnectionPropertiesResponseArgs

func PrivateEndpointConnectionPropertiesResponsePtr(v *PrivateEndpointConnectionPropertiesResponseArgs) PrivateEndpointConnectionPropertiesResponsePtrInput {
	return (*privateEndpointConnectionPropertiesResponsePtrType)(v)
}

func (*privateEndpointConnectionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponse {
		return &v
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) Elem() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) PrivateEndpointConnectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponse
		return ret
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStateResponse {
		if v == nil {
			return nil
		}
		return &v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Etag       string                                       `pulumi:"etag"`
	Id         string                                       `pulumi:"id"`
	Location   *string                                      `pulumi:"location"`
	Name       string                                       `pulumi:"name"`
	Properties *PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                           `pulumi:"systemData"`
	Type       string                                       `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Etag       pulumi.StringInput                                  `pulumi:"etag"`
	Id         pulumi.StringInput                                  `pulumi:"id"`
	Location   pulumi.StringPtrInput                               `pulumi:"location"`
	Name       pulumi.StringInput                                  `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponsePtrInput `pulumi:"properties"`
	SystemData SystemDataResponseInput                             `pulumi:"systemData"`
	Type       pulumi.StringInput                                  `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type QuotaLimitResponse struct {
	Count         *float64                 `pulumi:"count"`
	RenewalPeriod *float64                 `pulumi:"renewalPeriod"`
	Rules         []ThrottlingRuleResponse `pulumi:"rules"`
}





type QuotaLimitResponseInput interface {
	pulumi.Input

	ToQuotaLimitResponseOutput() QuotaLimitResponseOutput
	ToQuotaLimitResponseOutputWithContext(context.Context) QuotaLimitResponseOutput
}

type QuotaLimitResponseArgs struct {
	Count         pulumi.Float64PtrInput           `pulumi:"count"`
	RenewalPeriod pulumi.Float64PtrInput           `pulumi:"renewalPeriod"`
	Rules         ThrottlingRuleResponseArrayInput `pulumi:"rules"`
}

func (QuotaLimitResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QuotaLimitResponse)(nil)).Elem()
}

func (i QuotaLimitResponseArgs) ToQuotaLimitResponseOutput() QuotaLimitResponseOutput {
	return i.ToQuotaLimitResponseOutputWithContext(context.Background())
}

func (i QuotaLimitResponseArgs) ToQuotaLimitResponseOutputWithContext(ctx context.Context) QuotaLimitResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaLimitResponseOutput)
}

func (i QuotaLimitResponseArgs) ToQuotaLimitResponsePtrOutput() QuotaLimitResponsePtrOutput {
	return i.ToQuotaLimitResponsePtrOutputWithContext(context.Background())
}

func (i QuotaLimitResponseArgs) ToQuotaLimitResponsePtrOutputWithContext(ctx context.Context) QuotaLimitResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaLimitResponseOutput).ToQuotaLimitResponsePtrOutputWithContext(ctx)
}









type QuotaLimitResponsePtrInput interface {
	pulumi.Input

	ToQuotaLimitResponsePtrOutput() QuotaLimitResponsePtrOutput
	ToQuotaLimitResponsePtrOutputWithContext(context.Context) QuotaLimitResponsePtrOutput
}

type quotaLimitResponsePtrType QuotaLimitResponseArgs

func QuotaLimitResponsePtr(v *QuotaLimitResponseArgs) QuotaLimitResponsePtrInput {
	return (*quotaLimitResponsePtrType)(v)
}

func (*quotaLimitResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaLimitResponse)(nil)).Elem()
}

func (i *quotaLimitResponsePtrType) ToQuotaLimitResponsePtrOutput() QuotaLimitResponsePtrOutput {
	return i.ToQuotaLimitResponsePtrOutputWithContext(context.Background())
}

func (i *quotaLimitResponsePtrType) ToQuotaLimitResponsePtrOutputWithContext(ctx context.Context) QuotaLimitResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaLimitResponsePtrOutput)
}

type QuotaLimitResponseOutput struct{ *pulumi.OutputState }

func (QuotaLimitResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuotaLimitResponse)(nil)).Elem()
}

func (o QuotaLimitResponseOutput) ToQuotaLimitResponseOutput() QuotaLimitResponseOutput {
	return o
}

func (o QuotaLimitResponseOutput) ToQuotaLimitResponseOutputWithContext(ctx context.Context) QuotaLimitResponseOutput {
	return o
}

func (o QuotaLimitResponseOutput) ToQuotaLimitResponsePtrOutput() QuotaLimitResponsePtrOutput {
	return o.ToQuotaLimitResponsePtrOutputWithContext(context.Background())
}

func (o QuotaLimitResponseOutput) ToQuotaLimitResponsePtrOutputWithContext(ctx context.Context) QuotaLimitResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QuotaLimitResponse) *QuotaLimitResponse {
		return &v
	}).(QuotaLimitResponsePtrOutput)
}

func (o QuotaLimitResponseOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v QuotaLimitResponse) *float64 { return v.Count }).(pulumi.Float64PtrOutput)
}

func (o QuotaLimitResponseOutput) RenewalPeriod() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v QuotaLimitResponse) *float64 { return v.RenewalPeriod }).(pulumi.Float64PtrOutput)
}

func (o QuotaLimitResponseOutput) Rules() ThrottlingRuleResponseArrayOutput {
	return o.ApplyT(func(v QuotaLimitResponse) []ThrottlingRuleResponse { return v.Rules }).(ThrottlingRuleResponseArrayOutput)
}

type QuotaLimitResponsePtrOutput struct{ *pulumi.OutputState }

func (QuotaLimitResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaLimitResponse)(nil)).Elem()
}

func (o QuotaLimitResponsePtrOutput) ToQuotaLimitResponsePtrOutput() QuotaLimitResponsePtrOutput {
	return o
}

func (o QuotaLimitResponsePtrOutput) ToQuotaLimitResponsePtrOutputWithContext(ctx context.Context) QuotaLimitResponsePtrOutput {
	return o
}

func (o QuotaLimitResponsePtrOutput) Elem() QuotaLimitResponseOutput {
	return o.ApplyT(func(v *QuotaLimitResponse) QuotaLimitResponse {
		if v != nil {
			return *v
		}
		var ret QuotaLimitResponse
		return ret
	}).(QuotaLimitResponseOutput)
}

func (o QuotaLimitResponsePtrOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *QuotaLimitResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.Float64PtrOutput)
}

func (o QuotaLimitResponsePtrOutput) RenewalPeriod() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *QuotaLimitResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RenewalPeriod
	}).(pulumi.Float64PtrOutput)
}

func (o QuotaLimitResponsePtrOutput) Rules() ThrottlingRuleResponseArrayOutput {
	return o.ApplyT(func(v *QuotaLimitResponse) []ThrottlingRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ThrottlingRuleResponseArrayOutput)
}

type RequestMatchPatternResponse struct {
	Method *string `pulumi:"method"`
	Path   *string `pulumi:"path"`
}





type RequestMatchPatternResponseInput interface {
	pulumi.Input

	ToRequestMatchPatternResponseOutput() RequestMatchPatternResponseOutput
	ToRequestMatchPatternResponseOutputWithContext(context.Context) RequestMatchPatternResponseOutput
}

type RequestMatchPatternResponseArgs struct {
	Method pulumi.StringPtrInput `pulumi:"method"`
	Path   pulumi.StringPtrInput `pulumi:"path"`
}

func (RequestMatchPatternResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestMatchPatternResponse)(nil)).Elem()
}

func (i RequestMatchPatternResponseArgs) ToRequestMatchPatternResponseOutput() RequestMatchPatternResponseOutput {
	return i.ToRequestMatchPatternResponseOutputWithContext(context.Background())
}

func (i RequestMatchPatternResponseArgs) ToRequestMatchPatternResponseOutputWithContext(ctx context.Context) RequestMatchPatternResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestMatchPatternResponseOutput)
}





type RequestMatchPatternResponseArrayInput interface {
	pulumi.Input

	ToRequestMatchPatternResponseArrayOutput() RequestMatchPatternResponseArrayOutput
	ToRequestMatchPatternResponseArrayOutputWithContext(context.Context) RequestMatchPatternResponseArrayOutput
}

type RequestMatchPatternResponseArray []RequestMatchPatternResponseInput

func (RequestMatchPatternResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RequestMatchPatternResponse)(nil)).Elem()
}

func (i RequestMatchPatternResponseArray) ToRequestMatchPatternResponseArrayOutput() RequestMatchPatternResponseArrayOutput {
	return i.ToRequestMatchPatternResponseArrayOutputWithContext(context.Background())
}

func (i RequestMatchPatternResponseArray) ToRequestMatchPatternResponseArrayOutputWithContext(ctx context.Context) RequestMatchPatternResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestMatchPatternResponseArrayOutput)
}

type RequestMatchPatternResponseOutput struct{ *pulumi.OutputState }

func (RequestMatchPatternResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestMatchPatternResponse)(nil)).Elem()
}

func (o RequestMatchPatternResponseOutput) ToRequestMatchPatternResponseOutput() RequestMatchPatternResponseOutput {
	return o
}

func (o RequestMatchPatternResponseOutput) ToRequestMatchPatternResponseOutputWithContext(ctx context.Context) RequestMatchPatternResponseOutput {
	return o
}

func (o RequestMatchPatternResponseOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestMatchPatternResponse) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o RequestMatchPatternResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestMatchPatternResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type RequestMatchPatternResponseArrayOutput struct{ *pulumi.OutputState }

func (RequestMatchPatternResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RequestMatchPatternResponse)(nil)).Elem()
}

func (o RequestMatchPatternResponseArrayOutput) ToRequestMatchPatternResponseArrayOutput() RequestMatchPatternResponseArrayOutput {
	return o
}

func (o RequestMatchPatternResponseArrayOutput) ToRequestMatchPatternResponseArrayOutputWithContext(ctx context.Context) RequestMatchPatternResponseArrayOutput {
	return o
}

func (o RequestMatchPatternResponseArrayOutput) Index(i pulumi.IntInput) RequestMatchPatternResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RequestMatchPatternResponse {
		return vs[0].([]RequestMatchPatternResponse)[vs[1].(int)]
	}).(RequestMatchPatternResponseOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuCapabilityResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type SkuCapabilityResponseInput interface {
	pulumi.Input

	ToSkuCapabilityResponseOutput() SkuCapabilityResponseOutput
	ToSkuCapabilityResponseOutputWithContext(context.Context) SkuCapabilityResponseOutput
}

type SkuCapabilityResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (SkuCapabilityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapabilityResponse)(nil)).Elem()
}

func (i SkuCapabilityResponseArgs) ToSkuCapabilityResponseOutput() SkuCapabilityResponseOutput {
	return i.ToSkuCapabilityResponseOutputWithContext(context.Background())
}

func (i SkuCapabilityResponseArgs) ToSkuCapabilityResponseOutputWithContext(ctx context.Context) SkuCapabilityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapabilityResponseOutput)
}





type SkuCapabilityResponseArrayInput interface {
	pulumi.Input

	ToSkuCapabilityResponseArrayOutput() SkuCapabilityResponseArrayOutput
	ToSkuCapabilityResponseArrayOutputWithContext(context.Context) SkuCapabilityResponseArrayOutput
}

type SkuCapabilityResponseArray []SkuCapabilityResponseInput

func (SkuCapabilityResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SkuCapabilityResponse)(nil)).Elem()
}

func (i SkuCapabilityResponseArray) ToSkuCapabilityResponseArrayOutput() SkuCapabilityResponseArrayOutput {
	return i.ToSkuCapabilityResponseArrayOutputWithContext(context.Background())
}

func (i SkuCapabilityResponseArray) ToSkuCapabilityResponseArrayOutputWithContext(ctx context.Context) SkuCapabilityResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapabilityResponseArrayOutput)
}

type SkuCapabilityResponseOutput struct{ *pulumi.OutputState }

func (SkuCapabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapabilityResponse)(nil)).Elem()
}

func (o SkuCapabilityResponseOutput) ToSkuCapabilityResponseOutput() SkuCapabilityResponseOutput {
	return o
}

func (o SkuCapabilityResponseOutput) ToSkuCapabilityResponseOutputWithContext(ctx context.Context) SkuCapabilityResponseOutput {
	return o
}

func (o SkuCapabilityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuCapabilityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuCapabilityResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuCapabilityResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SkuCapabilityResponseArrayOutput struct{ *pulumi.OutputState }

func (SkuCapabilityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SkuCapabilityResponse)(nil)).Elem()
}

func (o SkuCapabilityResponseArrayOutput) ToSkuCapabilityResponseArrayOutput() SkuCapabilityResponseArrayOutput {
	return o
}

func (o SkuCapabilityResponseArrayOutput) ToSkuCapabilityResponseArrayOutputWithContext(ctx context.Context) SkuCapabilityResponseArrayOutput {
	return o
}

func (o SkuCapabilityResponseArrayOutput) Index(i pulumi.IntInput) SkuCapabilityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SkuCapabilityResponse {
		return vs[0].([]SkuCapabilityResponse)[vs[1].(int)]
	}).(SkuCapabilityResponseOutput)
}

type SkuChangeInfoResponse struct {
	CountOfDowngrades              *float64 `pulumi:"countOfDowngrades"`
	CountOfUpgradesAfterDowngrades *float64 `pulumi:"countOfUpgradesAfterDowngrades"`
	LastChangeDate                 *string  `pulumi:"lastChangeDate"`
}





type SkuChangeInfoResponseInput interface {
	pulumi.Input

	ToSkuChangeInfoResponseOutput() SkuChangeInfoResponseOutput
	ToSkuChangeInfoResponseOutputWithContext(context.Context) SkuChangeInfoResponseOutput
}

type SkuChangeInfoResponseArgs struct {
	CountOfDowngrades              pulumi.Float64PtrInput `pulumi:"countOfDowngrades"`
	CountOfUpgradesAfterDowngrades pulumi.Float64PtrInput `pulumi:"countOfUpgradesAfterDowngrades"`
	LastChangeDate                 pulumi.StringPtrInput  `pulumi:"lastChangeDate"`
}

func (SkuChangeInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuChangeInfoResponse)(nil)).Elem()
}

func (i SkuChangeInfoResponseArgs) ToSkuChangeInfoResponseOutput() SkuChangeInfoResponseOutput {
	return i.ToSkuChangeInfoResponseOutputWithContext(context.Background())
}

func (i SkuChangeInfoResponseArgs) ToSkuChangeInfoResponseOutputWithContext(ctx context.Context) SkuChangeInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuChangeInfoResponseOutput)
}

func (i SkuChangeInfoResponseArgs) ToSkuChangeInfoResponsePtrOutput() SkuChangeInfoResponsePtrOutput {
	return i.ToSkuChangeInfoResponsePtrOutputWithContext(context.Background())
}

func (i SkuChangeInfoResponseArgs) ToSkuChangeInfoResponsePtrOutputWithContext(ctx context.Context) SkuChangeInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuChangeInfoResponseOutput).ToSkuChangeInfoResponsePtrOutputWithContext(ctx)
}









type SkuChangeInfoResponsePtrInput interface {
	pulumi.Input

	ToSkuChangeInfoResponsePtrOutput() SkuChangeInfoResponsePtrOutput
	ToSkuChangeInfoResponsePtrOutputWithContext(context.Context) SkuChangeInfoResponsePtrOutput
}

type skuChangeInfoResponsePtrType SkuChangeInfoResponseArgs

func SkuChangeInfoResponsePtr(v *SkuChangeInfoResponseArgs) SkuChangeInfoResponsePtrInput {
	return (*skuChangeInfoResponsePtrType)(v)
}

func (*skuChangeInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuChangeInfoResponse)(nil)).Elem()
}

func (i *skuChangeInfoResponsePtrType) ToSkuChangeInfoResponsePtrOutput() SkuChangeInfoResponsePtrOutput {
	return i.ToSkuChangeInfoResponsePtrOutputWithContext(context.Background())
}

func (i *skuChangeInfoResponsePtrType) ToSkuChangeInfoResponsePtrOutputWithContext(ctx context.Context) SkuChangeInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuChangeInfoResponsePtrOutput)
}

type SkuChangeInfoResponseOutput struct{ *pulumi.OutputState }

func (SkuChangeInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuChangeInfoResponse)(nil)).Elem()
}

func (o SkuChangeInfoResponseOutput) ToSkuChangeInfoResponseOutput() SkuChangeInfoResponseOutput {
	return o
}

func (o SkuChangeInfoResponseOutput) ToSkuChangeInfoResponseOutputWithContext(ctx context.Context) SkuChangeInfoResponseOutput {
	return o
}

func (o SkuChangeInfoResponseOutput) ToSkuChangeInfoResponsePtrOutput() SkuChangeInfoResponsePtrOutput {
	return o.ToSkuChangeInfoResponsePtrOutputWithContext(context.Background())
}

func (o SkuChangeInfoResponseOutput) ToSkuChangeInfoResponsePtrOutputWithContext(ctx context.Context) SkuChangeInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuChangeInfoResponse) *SkuChangeInfoResponse {
		return &v
	}).(SkuChangeInfoResponsePtrOutput)
}

func (o SkuChangeInfoResponseOutput) CountOfDowngrades() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SkuChangeInfoResponse) *float64 { return v.CountOfDowngrades }).(pulumi.Float64PtrOutput)
}

func (o SkuChangeInfoResponseOutput) CountOfUpgradesAfterDowngrades() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SkuChangeInfoResponse) *float64 { return v.CountOfUpgradesAfterDowngrades }).(pulumi.Float64PtrOutput)
}

func (o SkuChangeInfoResponseOutput) LastChangeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuChangeInfoResponse) *string { return v.LastChangeDate }).(pulumi.StringPtrOutput)
}

type SkuChangeInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuChangeInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuChangeInfoResponse)(nil)).Elem()
}

func (o SkuChangeInfoResponsePtrOutput) ToSkuChangeInfoResponsePtrOutput() SkuChangeInfoResponsePtrOutput {
	return o
}

func (o SkuChangeInfoResponsePtrOutput) ToSkuChangeInfoResponsePtrOutputWithContext(ctx context.Context) SkuChangeInfoResponsePtrOutput {
	return o
}

func (o SkuChangeInfoResponsePtrOutput) Elem() SkuChangeInfoResponseOutput {
	return o.ApplyT(func(v *SkuChangeInfoResponse) SkuChangeInfoResponse {
		if v != nil {
			return *v
		}
		var ret SkuChangeInfoResponse
		return ret
	}).(SkuChangeInfoResponseOutput)
}

func (o SkuChangeInfoResponsePtrOutput) CountOfDowngrades() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SkuChangeInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.CountOfDowngrades
	}).(pulumi.Float64PtrOutput)
}

func (o SkuChangeInfoResponsePtrOutput) CountOfUpgradesAfterDowngrades() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SkuChangeInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.CountOfUpgradesAfterDowngrades
	}).(pulumi.Float64PtrOutput)
}

func (o SkuChangeInfoResponsePtrOutput) LastChangeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuChangeInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastChangeDate
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type ThrottlingRuleResponse struct {
	Count                    *float64                      `pulumi:"count"`
	DynamicThrottlingEnabled *bool                         `pulumi:"dynamicThrottlingEnabled"`
	Key                      *string                       `pulumi:"key"`
	MatchPatterns            []RequestMatchPatternResponse `pulumi:"matchPatterns"`
	MinCount                 *float64                      `pulumi:"minCount"`
	RenewalPeriod            *float64                      `pulumi:"renewalPeriod"`
}





type ThrottlingRuleResponseInput interface {
	pulumi.Input

	ToThrottlingRuleResponseOutput() ThrottlingRuleResponseOutput
	ToThrottlingRuleResponseOutputWithContext(context.Context) ThrottlingRuleResponseOutput
}

type ThrottlingRuleResponseArgs struct {
	Count                    pulumi.Float64PtrInput                `pulumi:"count"`
	DynamicThrottlingEnabled pulumi.BoolPtrInput                   `pulumi:"dynamicThrottlingEnabled"`
	Key                      pulumi.StringPtrInput                 `pulumi:"key"`
	MatchPatterns            RequestMatchPatternResponseArrayInput `pulumi:"matchPatterns"`
	MinCount                 pulumi.Float64PtrInput                `pulumi:"minCount"`
	RenewalPeriod            pulumi.Float64PtrInput                `pulumi:"renewalPeriod"`
}

func (ThrottlingRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingRuleResponse)(nil)).Elem()
}

func (i ThrottlingRuleResponseArgs) ToThrottlingRuleResponseOutput() ThrottlingRuleResponseOutput {
	return i.ToThrottlingRuleResponseOutputWithContext(context.Background())
}

func (i ThrottlingRuleResponseArgs) ToThrottlingRuleResponseOutputWithContext(ctx context.Context) ThrottlingRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingRuleResponseOutput)
}





type ThrottlingRuleResponseArrayInput interface {
	pulumi.Input

	ToThrottlingRuleResponseArrayOutput() ThrottlingRuleResponseArrayOutput
	ToThrottlingRuleResponseArrayOutputWithContext(context.Context) ThrottlingRuleResponseArrayOutput
}

type ThrottlingRuleResponseArray []ThrottlingRuleResponseInput

func (ThrottlingRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThrottlingRuleResponse)(nil)).Elem()
}

func (i ThrottlingRuleResponseArray) ToThrottlingRuleResponseArrayOutput() ThrottlingRuleResponseArrayOutput {
	return i.ToThrottlingRuleResponseArrayOutputWithContext(context.Background())
}

func (i ThrottlingRuleResponseArray) ToThrottlingRuleResponseArrayOutputWithContext(ctx context.Context) ThrottlingRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingRuleResponseArrayOutput)
}

type ThrottlingRuleResponseOutput struct{ *pulumi.OutputState }

func (ThrottlingRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingRuleResponse)(nil)).Elem()
}

func (o ThrottlingRuleResponseOutput) ToThrottlingRuleResponseOutput() ThrottlingRuleResponseOutput {
	return o
}

func (o ThrottlingRuleResponseOutput) ToThrottlingRuleResponseOutputWithContext(ctx context.Context) ThrottlingRuleResponseOutput {
	return o
}

func (o ThrottlingRuleResponseOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ThrottlingRuleResponse) *float64 { return v.Count }).(pulumi.Float64PtrOutput)
}

func (o ThrottlingRuleResponseOutput) DynamicThrottlingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ThrottlingRuleResponse) *bool { return v.DynamicThrottlingEnabled }).(pulumi.BoolPtrOutput)
}

func (o ThrottlingRuleResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThrottlingRuleResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o ThrottlingRuleResponseOutput) MatchPatterns() RequestMatchPatternResponseArrayOutput {
	return o.ApplyT(func(v ThrottlingRuleResponse) []RequestMatchPatternResponse { return v.MatchPatterns }).(RequestMatchPatternResponseArrayOutput)
}

func (o ThrottlingRuleResponseOutput) MinCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ThrottlingRuleResponse) *float64 { return v.MinCount }).(pulumi.Float64PtrOutput)
}

func (o ThrottlingRuleResponseOutput) RenewalPeriod() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ThrottlingRuleResponse) *float64 { return v.RenewalPeriod }).(pulumi.Float64PtrOutput)
}

type ThrottlingRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ThrottlingRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThrottlingRuleResponse)(nil)).Elem()
}

func (o ThrottlingRuleResponseArrayOutput) ToThrottlingRuleResponseArrayOutput() ThrottlingRuleResponseArrayOutput {
	return o
}

func (o ThrottlingRuleResponseArrayOutput) ToThrottlingRuleResponseArrayOutputWithContext(ctx context.Context) ThrottlingRuleResponseArrayOutput {
	return o
}

func (o ThrottlingRuleResponseArrayOutput) Index(i pulumi.IntInput) ThrottlingRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThrottlingRuleResponse {
		return vs[0].([]ThrottlingRuleResponse)[vs[1].(int)]
	}).(ThrottlingRuleResponseOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type UserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput
	ToUserAssignedIdentityResponseOutputWithContext(context.Context) UserAssignedIdentityResponseOutput
}

type UserAssignedIdentityResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (UserAssignedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return i.ToUserAssignedIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseOutput)
}





type UserAssignedIdentityResponseMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput
	ToUserAssignedIdentityResponseMapOutputWithContext(context.Context) UserAssignedIdentityResponseMapOutput
}

type UserAssignedIdentityResponseMap map[string]UserAssignedIdentityResponseInput

func (UserAssignedIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return i.ToUserAssignedIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseMapOutput)
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type UserOwnedStorage struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	ResourceId       *string `pulumi:"resourceId"`
}





type UserOwnedStorageInput interface {
	pulumi.Input

	ToUserOwnedStorageOutput() UserOwnedStorageOutput
	ToUserOwnedStorageOutputWithContext(context.Context) UserOwnedStorageOutput
}

type UserOwnedStorageArgs struct {
	IdentityClientId pulumi.StringPtrInput `pulumi:"identityClientId"`
	ResourceId       pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (UserOwnedStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserOwnedStorage)(nil)).Elem()
}

func (i UserOwnedStorageArgs) ToUserOwnedStorageOutput() UserOwnedStorageOutput {
	return i.ToUserOwnedStorageOutputWithContext(context.Background())
}

func (i UserOwnedStorageArgs) ToUserOwnedStorageOutputWithContext(ctx context.Context) UserOwnedStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOwnedStorageOutput)
}





type UserOwnedStorageArrayInput interface {
	pulumi.Input

	ToUserOwnedStorageArrayOutput() UserOwnedStorageArrayOutput
	ToUserOwnedStorageArrayOutputWithContext(context.Context) UserOwnedStorageArrayOutput
}

type UserOwnedStorageArray []UserOwnedStorageInput

func (UserOwnedStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserOwnedStorage)(nil)).Elem()
}

func (i UserOwnedStorageArray) ToUserOwnedStorageArrayOutput() UserOwnedStorageArrayOutput {
	return i.ToUserOwnedStorageArrayOutputWithContext(context.Background())
}

func (i UserOwnedStorageArray) ToUserOwnedStorageArrayOutputWithContext(ctx context.Context) UserOwnedStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOwnedStorageArrayOutput)
}

type UserOwnedStorageOutput struct{ *pulumi.OutputState }

func (UserOwnedStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserOwnedStorage)(nil)).Elem()
}

func (o UserOwnedStorageOutput) ToUserOwnedStorageOutput() UserOwnedStorageOutput {
	return o
}

func (o UserOwnedStorageOutput) ToUserOwnedStorageOutputWithContext(ctx context.Context) UserOwnedStorageOutput {
	return o
}

func (o UserOwnedStorageOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserOwnedStorage) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o UserOwnedStorageOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserOwnedStorage) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type UserOwnedStorageArrayOutput struct{ *pulumi.OutputState }

func (UserOwnedStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserOwnedStorage)(nil)).Elem()
}

func (o UserOwnedStorageArrayOutput) ToUserOwnedStorageArrayOutput() UserOwnedStorageArrayOutput {
	return o
}

func (o UserOwnedStorageArrayOutput) ToUserOwnedStorageArrayOutputWithContext(ctx context.Context) UserOwnedStorageArrayOutput {
	return o
}

func (o UserOwnedStorageArrayOutput) Index(i pulumi.IntInput) UserOwnedStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserOwnedStorage {
		return vs[0].([]UserOwnedStorage)[vs[1].(int)]
	}).(UserOwnedStorageOutput)
}

type UserOwnedStorageResponse struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	ResourceId       *string `pulumi:"resourceId"`
}





type UserOwnedStorageResponseInput interface {
	pulumi.Input

	ToUserOwnedStorageResponseOutput() UserOwnedStorageResponseOutput
	ToUserOwnedStorageResponseOutputWithContext(context.Context) UserOwnedStorageResponseOutput
}

type UserOwnedStorageResponseArgs struct {
	IdentityClientId pulumi.StringPtrInput `pulumi:"identityClientId"`
	ResourceId       pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (UserOwnedStorageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserOwnedStorageResponse)(nil)).Elem()
}

func (i UserOwnedStorageResponseArgs) ToUserOwnedStorageResponseOutput() UserOwnedStorageResponseOutput {
	return i.ToUserOwnedStorageResponseOutputWithContext(context.Background())
}

func (i UserOwnedStorageResponseArgs) ToUserOwnedStorageResponseOutputWithContext(ctx context.Context) UserOwnedStorageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOwnedStorageResponseOutput)
}





type UserOwnedStorageResponseArrayInput interface {
	pulumi.Input

	ToUserOwnedStorageResponseArrayOutput() UserOwnedStorageResponseArrayOutput
	ToUserOwnedStorageResponseArrayOutputWithContext(context.Context) UserOwnedStorageResponseArrayOutput
}

type UserOwnedStorageResponseArray []UserOwnedStorageResponseInput

func (UserOwnedStorageResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserOwnedStorageResponse)(nil)).Elem()
}

func (i UserOwnedStorageResponseArray) ToUserOwnedStorageResponseArrayOutput() UserOwnedStorageResponseArrayOutput {
	return i.ToUserOwnedStorageResponseArrayOutputWithContext(context.Background())
}

func (i UserOwnedStorageResponseArray) ToUserOwnedStorageResponseArrayOutputWithContext(ctx context.Context) UserOwnedStorageResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOwnedStorageResponseArrayOutput)
}

type UserOwnedStorageResponseOutput struct{ *pulumi.OutputState }

func (UserOwnedStorageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserOwnedStorageResponse)(nil)).Elem()
}

func (o UserOwnedStorageResponseOutput) ToUserOwnedStorageResponseOutput() UserOwnedStorageResponseOutput {
	return o
}

func (o UserOwnedStorageResponseOutput) ToUserOwnedStorageResponseOutputWithContext(ctx context.Context) UserOwnedStorageResponseOutput {
	return o
}

func (o UserOwnedStorageResponseOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserOwnedStorageResponse) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o UserOwnedStorageResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserOwnedStorageResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type UserOwnedStorageResponseArrayOutput struct{ *pulumi.OutputState }

func (UserOwnedStorageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserOwnedStorageResponse)(nil)).Elem()
}

func (o UserOwnedStorageResponseArrayOutput) ToUserOwnedStorageResponseArrayOutput() UserOwnedStorageResponseArrayOutput {
	return o
}

func (o UserOwnedStorageResponseArrayOutput) ToUserOwnedStorageResponseArrayOutputWithContext(ctx context.Context) UserOwnedStorageResponseArrayOutput {
	return o
}

func (o UserOwnedStorageResponseArrayOutput) Index(i pulumi.IntInput) UserOwnedStorageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserOwnedStorageResponse {
		return vs[0].([]UserOwnedStorageResponse)[vs[1].(int)]
	}).(UserOwnedStorageResponseOutput)
}

type VirtualNetworkRule struct {
	Id                               string  `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint *bool   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	State                            *string `pulumi:"state"`
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Id                               pulumi.StringInput    `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	State                            pulumi.StringPtrInput `pulumi:"state"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRule) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkRuleOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *string { return v.State }).(pulumi.StringPtrOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Id                               string  `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint *bool   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	State                            *string `pulumi:"state"`
}





type VirtualNetworkRuleResponseInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput
	ToVirtualNetworkRuleResponseOutputWithContext(context.Context) VirtualNetworkRuleResponseOutput
}

type VirtualNetworkRuleResponseArgs struct {
	Id                               pulumi.StringInput    `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	State                            pulumi.StringPtrInput `pulumi:"state"`
}

func (VirtualNetworkRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return i.ToVirtualNetworkRuleResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseOutput)
}





type VirtualNetworkRuleResponseArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput
	ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Context) VirtualNetworkRuleResponseArrayOutput
}

type VirtualNetworkRuleResponseArray []VirtualNetworkRuleResponseInput

func (VirtualNetworkRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return i.ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseArrayOutput)
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleResponseOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountPropertiesOutput{})
	pulumi.RegisterOutputType(AccountPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AccountPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiPropertiesOutput{})
	pulumi.RegisterOutputType(ApiPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApiPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApiPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CallRateLimitResponseOutput{})
	pulumi.RegisterOutputType(CallRateLimitResponsePtrOutput{})
	pulumi.RegisterOutputType(CommitmentPeriodOutput{})
	pulumi.RegisterOutputType(CommitmentPeriodPtrOutput{})
	pulumi.RegisterOutputType(CommitmentPeriodResponseOutput{})
	pulumi.RegisterOutputType(CommitmentPeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(CommitmentPlanPropertiesOutput{})
	pulumi.RegisterOutputType(CommitmentPlanPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CommitmentPlanPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CommitmentPlanPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CommitmentQuotaResponseOutput{})
	pulumi.RegisterOutputType(CommitmentQuotaResponsePtrOutput{})
	pulumi.RegisterOutputType(DeploymentModelOutput{})
	pulumi.RegisterOutputType(DeploymentModelPtrOutput{})
	pulumi.RegisterOutputType(DeploymentModelResponseOutput{})
	pulumi.RegisterOutputType(DeploymentModelResponsePtrOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(DeploymentScaleSettingsOutput{})
	pulumi.RegisterOutputType(DeploymentScaleSettingsPtrOutput{})
	pulumi.RegisterOutputType(DeploymentScaleSettingsResponseOutput{})
	pulumi.RegisterOutputType(DeploymentScaleSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IpRuleOutput{})
	pulumi.RegisterOutputType(IpRuleArrayOutput{})
	pulumi.RegisterOutputType(IpRuleResponseOutput{})
	pulumi.RegisterOutputType(IpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(QuotaLimitResponseOutput{})
	pulumi.RegisterOutputType(QuotaLimitResponsePtrOutput{})
	pulumi.RegisterOutputType(RequestMatchPatternResponseOutput{})
	pulumi.RegisterOutputType(RequestMatchPatternResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuCapabilityResponseOutput{})
	pulumi.RegisterOutputType(SkuCapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuChangeInfoResponseOutput{})
	pulumi.RegisterOutputType(SkuChangeInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(ThrottlingRuleResponseOutput{})
	pulumi.RegisterOutputType(ThrottlingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(UserOwnedStorageOutput{})
	pulumi.RegisterOutputType(UserOwnedStorageArrayOutput{})
	pulumi.RegisterOutputType(UserOwnedStorageResponseOutput{})
	pulumi.RegisterOutputType(UserOwnedStorageResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
