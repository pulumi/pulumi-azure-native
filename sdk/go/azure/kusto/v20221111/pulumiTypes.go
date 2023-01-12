


package v20221111

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AcceptedAudiences struct {
	Value *string `pulumi:"value"`
}





type AcceptedAudiencesInput interface {
	pulumi.Input

	ToAcceptedAudiencesOutput() AcceptedAudiencesOutput
	ToAcceptedAudiencesOutputWithContext(context.Context) AcceptedAudiencesOutput
}

type AcceptedAudiencesArgs struct {
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (AcceptedAudiencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceptedAudiences)(nil)).Elem()
}

func (i AcceptedAudiencesArgs) ToAcceptedAudiencesOutput() AcceptedAudiencesOutput {
	return i.ToAcceptedAudiencesOutputWithContext(context.Background())
}

func (i AcceptedAudiencesArgs) ToAcceptedAudiencesOutputWithContext(ctx context.Context) AcceptedAudiencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceptedAudiencesOutput)
}





type AcceptedAudiencesArrayInput interface {
	pulumi.Input

	ToAcceptedAudiencesArrayOutput() AcceptedAudiencesArrayOutput
	ToAcceptedAudiencesArrayOutputWithContext(context.Context) AcceptedAudiencesArrayOutput
}

type AcceptedAudiencesArray []AcceptedAudiencesInput

func (AcceptedAudiencesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AcceptedAudiences)(nil)).Elem()
}

func (i AcceptedAudiencesArray) ToAcceptedAudiencesArrayOutput() AcceptedAudiencesArrayOutput {
	return i.ToAcceptedAudiencesArrayOutputWithContext(context.Background())
}

func (i AcceptedAudiencesArray) ToAcceptedAudiencesArrayOutputWithContext(ctx context.Context) AcceptedAudiencesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceptedAudiencesArrayOutput)
}

type AcceptedAudiencesOutput struct{ *pulumi.OutputState }

func (AcceptedAudiencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceptedAudiences)(nil)).Elem()
}

func (o AcceptedAudiencesOutput) ToAcceptedAudiencesOutput() AcceptedAudiencesOutput {
	return o
}

func (o AcceptedAudiencesOutput) ToAcceptedAudiencesOutputWithContext(ctx context.Context) AcceptedAudiencesOutput {
	return o
}

func (o AcceptedAudiencesOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AcceptedAudiences) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type AcceptedAudiencesArrayOutput struct{ *pulumi.OutputState }

func (AcceptedAudiencesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AcceptedAudiences)(nil)).Elem()
}

func (o AcceptedAudiencesArrayOutput) ToAcceptedAudiencesArrayOutput() AcceptedAudiencesArrayOutput {
	return o
}

func (o AcceptedAudiencesArrayOutput) ToAcceptedAudiencesArrayOutputWithContext(ctx context.Context) AcceptedAudiencesArrayOutput {
	return o
}

func (o AcceptedAudiencesArrayOutput) Index(i pulumi.IntInput) AcceptedAudiencesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AcceptedAudiences {
		return vs[0].([]AcceptedAudiences)[vs[1].(int)]
	}).(AcceptedAudiencesOutput)
}

type AcceptedAudiencesResponse struct {
	Value *string `pulumi:"value"`
}

type AcceptedAudiencesResponseOutput struct{ *pulumi.OutputState }

func (AcceptedAudiencesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceptedAudiencesResponse)(nil)).Elem()
}

func (o AcceptedAudiencesResponseOutput) ToAcceptedAudiencesResponseOutput() AcceptedAudiencesResponseOutput {
	return o
}

func (o AcceptedAudiencesResponseOutput) ToAcceptedAudiencesResponseOutputWithContext(ctx context.Context) AcceptedAudiencesResponseOutput {
	return o
}

func (o AcceptedAudiencesResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AcceptedAudiencesResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type AcceptedAudiencesResponseArrayOutput struct{ *pulumi.OutputState }

func (AcceptedAudiencesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AcceptedAudiencesResponse)(nil)).Elem()
}

func (o AcceptedAudiencesResponseArrayOutput) ToAcceptedAudiencesResponseArrayOutput() AcceptedAudiencesResponseArrayOutput {
	return o
}

func (o AcceptedAudiencesResponseArrayOutput) ToAcceptedAudiencesResponseArrayOutputWithContext(ctx context.Context) AcceptedAudiencesResponseArrayOutput {
	return o
}

func (o AcceptedAudiencesResponseArrayOutput) Index(i pulumi.IntInput) AcceptedAudiencesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AcceptedAudiencesResponse {
		return vs[0].([]AcceptedAudiencesResponse)[vs[1].(int)]
	}).(AcceptedAudiencesResponseOutput)
}

type AzureSku struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
	Tier     string `pulumi:"tier"`
}





type AzureSkuInput interface {
	pulumi.Input

	ToAzureSkuOutput() AzureSkuOutput
	ToAzureSkuOutputWithContext(context.Context) AzureSkuOutput
}

type AzureSkuArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
	Tier     pulumi.StringInput `pulumi:"tier"`
}

func (AzureSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (i AzureSkuArgs) ToAzureSkuOutput() AzureSkuOutput {
	return i.ToAzureSkuOutputWithContext(context.Background())
}

func (i AzureSkuArgs) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuOutput)
}

type AzureSkuOutput struct{ *pulumi.OutputState }

func (AzureSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (o AzureSkuOutput) ToAzureSkuOutput() AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o AzureSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureSkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Tier }).(pulumi.StringOutput)
}

type AzureSkuResponse struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
	Tier     string `pulumi:"tier"`
}

type AzureSkuResponseOutput struct{ *pulumi.OutputState }

func (AzureSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuResponse)(nil)).Elem()
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutput() AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutputWithContext(ctx context.Context) AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o AzureSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureSkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type DatabasePrincipalResponse struct {
	AppId      *string `pulumi:"appId"`
	Email      *string `pulumi:"email"`
	Fqn        *string `pulumi:"fqn"`
	Name       string  `pulumi:"name"`
	Role       string  `pulumi:"role"`
	TenantName string  `pulumi:"tenantName"`
	Type       string  `pulumi:"type"`
}

type DatabasePrincipalResponseOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipalResponse)(nil)).Elem()
}

func (o DatabasePrincipalResponseOutput) ToDatabasePrincipalResponseOutput() DatabasePrincipalResponseOutput {
	return o
}

func (o DatabasePrincipalResponseOutput) ToDatabasePrincipalResponseOutputWithContext(ctx context.Context) DatabasePrincipalResponseOutput {
	return o
}

func (o DatabasePrincipalResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o DatabasePrincipalResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o DatabasePrincipalResponseOutput) Fqn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) *string { return v.Fqn }).(pulumi.StringPtrOutput)
}

func (o DatabasePrincipalResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DatabasePrincipalResponseOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) string { return v.Role }).(pulumi.StringOutput)
}

func (o DatabasePrincipalResponseOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) string { return v.TenantName }).(pulumi.StringOutput)
}

func (o DatabasePrincipalResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) string { return v.Type }).(pulumi.StringOutput)
}

type DatabasePrincipalResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabasePrincipalResponse)(nil)).Elem()
}

func (o DatabasePrincipalResponseArrayOutput) ToDatabasePrincipalResponseArrayOutput() DatabasePrincipalResponseArrayOutput {
	return o
}

func (o DatabasePrincipalResponseArrayOutput) ToDatabasePrincipalResponseArrayOutputWithContext(ctx context.Context) DatabasePrincipalResponseArrayOutput {
	return o
}

func (o DatabasePrincipalResponseArrayOutput) Index(i pulumi.IntInput) DatabasePrincipalResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabasePrincipalResponse {
		return vs[0].([]DatabasePrincipalResponse)[vs[1].(int)]
	}).(DatabasePrincipalResponseOutput)
}

type DatabaseStatisticsResponse struct {
	Size *float64 `pulumi:"size"`
}

type DatabaseStatisticsResponseOutput struct{ *pulumi.OutputState }

func (DatabaseStatisticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseStatisticsResponse)(nil)).Elem()
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponseOutput() DatabaseStatisticsResponseOutput {
	return o
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponseOutputWithContext(ctx context.Context) DatabaseStatisticsResponseOutput {
	return o
}

func (o DatabaseStatisticsResponseOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DatabaseStatisticsResponse) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

type FollowerDatabaseDefinitionResponse struct {
	AttachedDatabaseConfigurationName string                              `pulumi:"attachedDatabaseConfigurationName"`
	ClusterResourceId                 string                              `pulumi:"clusterResourceId"`
	DatabaseName                      string                              `pulumi:"databaseName"`
	DatabaseShareOrigin               string                              `pulumi:"databaseShareOrigin"`
	TableLevelSharingProperties       TableLevelSharingPropertiesResponse `pulumi:"tableLevelSharingProperties"`
}

type FollowerDatabaseDefinitionResponseOutput struct{ *pulumi.OutputState }

func (FollowerDatabaseDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FollowerDatabaseDefinitionResponse)(nil)).Elem()
}

func (o FollowerDatabaseDefinitionResponseOutput) ToFollowerDatabaseDefinitionResponseOutput() FollowerDatabaseDefinitionResponseOutput {
	return o
}

func (o FollowerDatabaseDefinitionResponseOutput) ToFollowerDatabaseDefinitionResponseOutputWithContext(ctx context.Context) FollowerDatabaseDefinitionResponseOutput {
	return o
}

func (o FollowerDatabaseDefinitionResponseOutput) AttachedDatabaseConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v FollowerDatabaseDefinitionResponse) string { return v.AttachedDatabaseConfigurationName }).(pulumi.StringOutput)
}

func (o FollowerDatabaseDefinitionResponseOutput) ClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v FollowerDatabaseDefinitionResponse) string { return v.ClusterResourceId }).(pulumi.StringOutput)
}

func (o FollowerDatabaseDefinitionResponseOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v FollowerDatabaseDefinitionResponse) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o FollowerDatabaseDefinitionResponseOutput) DatabaseShareOrigin() pulumi.StringOutput {
	return o.ApplyT(func(v FollowerDatabaseDefinitionResponse) string { return v.DatabaseShareOrigin }).(pulumi.StringOutput)
}

func (o FollowerDatabaseDefinitionResponseOutput) TableLevelSharingProperties() TableLevelSharingPropertiesResponseOutput {
	return o.ApplyT(func(v FollowerDatabaseDefinitionResponse) TableLevelSharingPropertiesResponse {
		return v.TableLevelSharingProperties
	}).(TableLevelSharingPropertiesResponseOutput)
}

type FollowerDatabaseDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (FollowerDatabaseDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FollowerDatabaseDefinitionResponse)(nil)).Elem()
}

func (o FollowerDatabaseDefinitionResponseArrayOutput) ToFollowerDatabaseDefinitionResponseArrayOutput() FollowerDatabaseDefinitionResponseArrayOutput {
	return o
}

func (o FollowerDatabaseDefinitionResponseArrayOutput) ToFollowerDatabaseDefinitionResponseArrayOutputWithContext(ctx context.Context) FollowerDatabaseDefinitionResponseArrayOutput {
	return o
}

func (o FollowerDatabaseDefinitionResponseArrayOutput) Index(i pulumi.IntInput) FollowerDatabaseDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FollowerDatabaseDefinitionResponse {
		return vs[0].([]FollowerDatabaseDefinitionResponse)[vs[1].(int)]
	}).(FollowerDatabaseDefinitionResponseOutput)
}

type Identity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
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

func (o IdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Identity) string { return v.Type }).(pulumi.StringOutput)
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

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
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
	PrincipalId            string                                            `pulumi:"principalId"`
	TenantId               string                                            `pulumi:"tenantId"`
	Type                   string                                            `pulumi:"type"`
	UserAssignedIdentities map[string]IdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
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

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
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
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
}

type IdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type IdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutput() IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o IdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type IdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutput() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) IdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]IdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(IdentityResponseUserAssignedIdentitiesOutput)
}

type KeyVaultProperties struct {
	KeyName      *string `pulumi:"keyName"`
	KeyVaultUri  *string `pulumi:"keyVaultUri"`
	KeyVersion   *string `pulumi:"keyVersion"`
	UserIdentity *string `pulumi:"userIdentity"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyName      pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri  pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion   pulumi.StringPtrInput `pulumi:"keyVersion"`
	UserIdentity pulumi.StringPtrInput `pulumi:"userIdentity"`
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

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) UserIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.UserIdentity }).(pulumi.StringPtrOutput)
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

func (o KeyVaultPropertiesPtrOutput) UserIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.UserIdentity
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	KeyName      *string `pulumi:"keyName"`
	KeyVaultUri  *string `pulumi:"keyVaultUri"`
	KeyVersion   *string `pulumi:"keyVersion"`
	UserIdentity *string `pulumi:"userIdentity"`
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

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) UserIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.UserIdentity }).(pulumi.StringPtrOutput)
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

func (o KeyVaultPropertiesResponsePtrOutput) UserIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserIdentity
	}).(pulumi.StringPtrOutput)
}

type LanguageExtension struct {
	LanguageExtensionImageName *string `pulumi:"languageExtensionImageName"`
	LanguageExtensionName      *string `pulumi:"languageExtensionName"`
}





type LanguageExtensionInput interface {
	pulumi.Input

	ToLanguageExtensionOutput() LanguageExtensionOutput
	ToLanguageExtensionOutputWithContext(context.Context) LanguageExtensionOutput
}

type LanguageExtensionArgs struct {
	LanguageExtensionImageName pulumi.StringPtrInput `pulumi:"languageExtensionImageName"`
	LanguageExtensionName      pulumi.StringPtrInput `pulumi:"languageExtensionName"`
}

func (LanguageExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LanguageExtension)(nil)).Elem()
}

func (i LanguageExtensionArgs) ToLanguageExtensionOutput() LanguageExtensionOutput {
	return i.ToLanguageExtensionOutputWithContext(context.Background())
}

func (i LanguageExtensionArgs) ToLanguageExtensionOutputWithContext(ctx context.Context) LanguageExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LanguageExtensionOutput)
}





type LanguageExtensionArrayInput interface {
	pulumi.Input

	ToLanguageExtensionArrayOutput() LanguageExtensionArrayOutput
	ToLanguageExtensionArrayOutputWithContext(context.Context) LanguageExtensionArrayOutput
}

type LanguageExtensionArray []LanguageExtensionInput

func (LanguageExtensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LanguageExtension)(nil)).Elem()
}

func (i LanguageExtensionArray) ToLanguageExtensionArrayOutput() LanguageExtensionArrayOutput {
	return i.ToLanguageExtensionArrayOutputWithContext(context.Background())
}

func (i LanguageExtensionArray) ToLanguageExtensionArrayOutputWithContext(ctx context.Context) LanguageExtensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LanguageExtensionArrayOutput)
}

type LanguageExtensionOutput struct{ *pulumi.OutputState }

func (LanguageExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LanguageExtension)(nil)).Elem()
}

func (o LanguageExtensionOutput) ToLanguageExtensionOutput() LanguageExtensionOutput {
	return o
}

func (o LanguageExtensionOutput) ToLanguageExtensionOutputWithContext(ctx context.Context) LanguageExtensionOutput {
	return o
}

func (o LanguageExtensionOutput) LanguageExtensionImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LanguageExtension) *string { return v.LanguageExtensionImageName }).(pulumi.StringPtrOutput)
}

func (o LanguageExtensionOutput) LanguageExtensionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LanguageExtension) *string { return v.LanguageExtensionName }).(pulumi.StringPtrOutput)
}

type LanguageExtensionArrayOutput struct{ *pulumi.OutputState }

func (LanguageExtensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LanguageExtension)(nil)).Elem()
}

func (o LanguageExtensionArrayOutput) ToLanguageExtensionArrayOutput() LanguageExtensionArrayOutput {
	return o
}

func (o LanguageExtensionArrayOutput) ToLanguageExtensionArrayOutputWithContext(ctx context.Context) LanguageExtensionArrayOutput {
	return o
}

func (o LanguageExtensionArrayOutput) Index(i pulumi.IntInput) LanguageExtensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LanguageExtension {
		return vs[0].([]LanguageExtension)[vs[1].(int)]
	}).(LanguageExtensionOutput)
}

type LanguageExtensionResponse struct {
	LanguageExtensionImageName *string `pulumi:"languageExtensionImageName"`
	LanguageExtensionName      *string `pulumi:"languageExtensionName"`
}

type LanguageExtensionResponseOutput struct{ *pulumi.OutputState }

func (LanguageExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LanguageExtensionResponse)(nil)).Elem()
}

func (o LanguageExtensionResponseOutput) ToLanguageExtensionResponseOutput() LanguageExtensionResponseOutput {
	return o
}

func (o LanguageExtensionResponseOutput) ToLanguageExtensionResponseOutputWithContext(ctx context.Context) LanguageExtensionResponseOutput {
	return o
}

func (o LanguageExtensionResponseOutput) LanguageExtensionImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LanguageExtensionResponse) *string { return v.LanguageExtensionImageName }).(pulumi.StringPtrOutput)
}

func (o LanguageExtensionResponseOutput) LanguageExtensionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LanguageExtensionResponse) *string { return v.LanguageExtensionName }).(pulumi.StringPtrOutput)
}

type LanguageExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (LanguageExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LanguageExtensionResponse)(nil)).Elem()
}

func (o LanguageExtensionResponseArrayOutput) ToLanguageExtensionResponseArrayOutput() LanguageExtensionResponseArrayOutput {
	return o
}

func (o LanguageExtensionResponseArrayOutput) ToLanguageExtensionResponseArrayOutputWithContext(ctx context.Context) LanguageExtensionResponseArrayOutput {
	return o
}

func (o LanguageExtensionResponseArrayOutput) Index(i pulumi.IntInput) LanguageExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LanguageExtensionResponse {
		return vs[0].([]LanguageExtensionResponse)[vs[1].(int)]
	}).(LanguageExtensionResponseOutput)
}

type LanguageExtensionsList struct {
	Value []LanguageExtension `pulumi:"value"`
}





type LanguageExtensionsListInput interface {
	pulumi.Input

	ToLanguageExtensionsListOutput() LanguageExtensionsListOutput
	ToLanguageExtensionsListOutputWithContext(context.Context) LanguageExtensionsListOutput
}

type LanguageExtensionsListArgs struct {
	Value LanguageExtensionArrayInput `pulumi:"value"`
}

func (LanguageExtensionsListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LanguageExtensionsList)(nil)).Elem()
}

func (i LanguageExtensionsListArgs) ToLanguageExtensionsListOutput() LanguageExtensionsListOutput {
	return i.ToLanguageExtensionsListOutputWithContext(context.Background())
}

func (i LanguageExtensionsListArgs) ToLanguageExtensionsListOutputWithContext(ctx context.Context) LanguageExtensionsListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LanguageExtensionsListOutput)
}

func (i LanguageExtensionsListArgs) ToLanguageExtensionsListPtrOutput() LanguageExtensionsListPtrOutput {
	return i.ToLanguageExtensionsListPtrOutputWithContext(context.Background())
}

func (i LanguageExtensionsListArgs) ToLanguageExtensionsListPtrOutputWithContext(ctx context.Context) LanguageExtensionsListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LanguageExtensionsListOutput).ToLanguageExtensionsListPtrOutputWithContext(ctx)
}









type LanguageExtensionsListPtrInput interface {
	pulumi.Input

	ToLanguageExtensionsListPtrOutput() LanguageExtensionsListPtrOutput
	ToLanguageExtensionsListPtrOutputWithContext(context.Context) LanguageExtensionsListPtrOutput
}

type languageExtensionsListPtrType LanguageExtensionsListArgs

func LanguageExtensionsListPtr(v *LanguageExtensionsListArgs) LanguageExtensionsListPtrInput {
	return (*languageExtensionsListPtrType)(v)
}

func (*languageExtensionsListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LanguageExtensionsList)(nil)).Elem()
}

func (i *languageExtensionsListPtrType) ToLanguageExtensionsListPtrOutput() LanguageExtensionsListPtrOutput {
	return i.ToLanguageExtensionsListPtrOutputWithContext(context.Background())
}

func (i *languageExtensionsListPtrType) ToLanguageExtensionsListPtrOutputWithContext(ctx context.Context) LanguageExtensionsListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LanguageExtensionsListPtrOutput)
}

type LanguageExtensionsListOutput struct{ *pulumi.OutputState }

func (LanguageExtensionsListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LanguageExtensionsList)(nil)).Elem()
}

func (o LanguageExtensionsListOutput) ToLanguageExtensionsListOutput() LanguageExtensionsListOutput {
	return o
}

func (o LanguageExtensionsListOutput) ToLanguageExtensionsListOutputWithContext(ctx context.Context) LanguageExtensionsListOutput {
	return o
}

func (o LanguageExtensionsListOutput) ToLanguageExtensionsListPtrOutput() LanguageExtensionsListPtrOutput {
	return o.ToLanguageExtensionsListPtrOutputWithContext(context.Background())
}

func (o LanguageExtensionsListOutput) ToLanguageExtensionsListPtrOutputWithContext(ctx context.Context) LanguageExtensionsListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LanguageExtensionsList) *LanguageExtensionsList {
		return &v
	}).(LanguageExtensionsListPtrOutput)
}

func (o LanguageExtensionsListOutput) Value() LanguageExtensionArrayOutput {
	return o.ApplyT(func(v LanguageExtensionsList) []LanguageExtension { return v.Value }).(LanguageExtensionArrayOutput)
}

type LanguageExtensionsListPtrOutput struct{ *pulumi.OutputState }

func (LanguageExtensionsListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LanguageExtensionsList)(nil)).Elem()
}

func (o LanguageExtensionsListPtrOutput) ToLanguageExtensionsListPtrOutput() LanguageExtensionsListPtrOutput {
	return o
}

func (o LanguageExtensionsListPtrOutput) ToLanguageExtensionsListPtrOutputWithContext(ctx context.Context) LanguageExtensionsListPtrOutput {
	return o
}

func (o LanguageExtensionsListPtrOutput) Elem() LanguageExtensionsListOutput {
	return o.ApplyT(func(v *LanguageExtensionsList) LanguageExtensionsList {
		if v != nil {
			return *v
		}
		var ret LanguageExtensionsList
		return ret
	}).(LanguageExtensionsListOutput)
}

func (o LanguageExtensionsListPtrOutput) Value() LanguageExtensionArrayOutput {
	return o.ApplyT(func(v *LanguageExtensionsList) []LanguageExtension {
		if v == nil {
			return nil
		}
		return v.Value
	}).(LanguageExtensionArrayOutput)
}

type LanguageExtensionsListResponse struct {
	Value []LanguageExtensionResponse `pulumi:"value"`
}

type LanguageExtensionsListResponseOutput struct{ *pulumi.OutputState }

func (LanguageExtensionsListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LanguageExtensionsListResponse)(nil)).Elem()
}

func (o LanguageExtensionsListResponseOutput) ToLanguageExtensionsListResponseOutput() LanguageExtensionsListResponseOutput {
	return o
}

func (o LanguageExtensionsListResponseOutput) ToLanguageExtensionsListResponseOutputWithContext(ctx context.Context) LanguageExtensionsListResponseOutput {
	return o
}

func (o LanguageExtensionsListResponseOutput) Value() LanguageExtensionResponseArrayOutput {
	return o.ApplyT(func(v LanguageExtensionsListResponse) []LanguageExtensionResponse { return v.Value }).(LanguageExtensionResponseArrayOutput)
}

type LanguageExtensionsListResponsePtrOutput struct{ *pulumi.OutputState }

func (LanguageExtensionsListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LanguageExtensionsListResponse)(nil)).Elem()
}

func (o LanguageExtensionsListResponsePtrOutput) ToLanguageExtensionsListResponsePtrOutput() LanguageExtensionsListResponsePtrOutput {
	return o
}

func (o LanguageExtensionsListResponsePtrOutput) ToLanguageExtensionsListResponsePtrOutputWithContext(ctx context.Context) LanguageExtensionsListResponsePtrOutput {
	return o
}

func (o LanguageExtensionsListResponsePtrOutput) Elem() LanguageExtensionsListResponseOutput {
	return o.ApplyT(func(v *LanguageExtensionsListResponse) LanguageExtensionsListResponse {
		if v != nil {
			return *v
		}
		var ret LanguageExtensionsListResponse
		return ret
	}).(LanguageExtensionsListResponseOutput)
}

func (o LanguageExtensionsListResponsePtrOutput) Value() LanguageExtensionResponseArrayOutput {
	return o.ApplyT(func(v *LanguageExtensionsListResponse) []LanguageExtensionResponse {
		if v == nil {
			return nil
		}
		return v.Value
	}).(LanguageExtensionResponseArrayOutput)
}

type OptimizedAutoscale struct {
	IsEnabled bool `pulumi:"isEnabled"`
	Maximum   int  `pulumi:"maximum"`
	Minimum   int  `pulumi:"minimum"`
	Version   int  `pulumi:"version"`
}





type OptimizedAutoscaleInput interface {
	pulumi.Input

	ToOptimizedAutoscaleOutput() OptimizedAutoscaleOutput
	ToOptimizedAutoscaleOutputWithContext(context.Context) OptimizedAutoscaleOutput
}

type OptimizedAutoscaleArgs struct {
	IsEnabled pulumi.BoolInput `pulumi:"isEnabled"`
	Maximum   pulumi.IntInput  `pulumi:"maximum"`
	Minimum   pulumi.IntInput  `pulumi:"minimum"`
	Version   pulumi.IntInput  `pulumi:"version"`
}

func (OptimizedAutoscaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OptimizedAutoscale)(nil)).Elem()
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscaleOutput() OptimizedAutoscaleOutput {
	return i.ToOptimizedAutoscaleOutputWithContext(context.Background())
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscaleOutputWithContext(ctx context.Context) OptimizedAutoscaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptimizedAutoscaleOutput)
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return i.ToOptimizedAutoscalePtrOutputWithContext(context.Background())
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptimizedAutoscaleOutput).ToOptimizedAutoscalePtrOutputWithContext(ctx)
}









type OptimizedAutoscalePtrInput interface {
	pulumi.Input

	ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput
	ToOptimizedAutoscalePtrOutputWithContext(context.Context) OptimizedAutoscalePtrOutput
}

type optimizedAutoscalePtrType OptimizedAutoscaleArgs

func OptimizedAutoscalePtr(v *OptimizedAutoscaleArgs) OptimizedAutoscalePtrInput {
	return (*optimizedAutoscalePtrType)(v)
}

func (*optimizedAutoscalePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OptimizedAutoscale)(nil)).Elem()
}

func (i *optimizedAutoscalePtrType) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return i.ToOptimizedAutoscalePtrOutputWithContext(context.Background())
}

func (i *optimizedAutoscalePtrType) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptimizedAutoscalePtrOutput)
}

type OptimizedAutoscaleOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OptimizedAutoscale)(nil)).Elem()
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscaleOutput() OptimizedAutoscaleOutput {
	return o
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscaleOutputWithContext(ctx context.Context) OptimizedAutoscaleOutput {
	return o
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return o.ToOptimizedAutoscalePtrOutputWithContext(context.Background())
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OptimizedAutoscale) *OptimizedAutoscale {
		return &v
	}).(OptimizedAutoscalePtrOutput)
}

func (o OptimizedAutoscaleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v OptimizedAutoscale) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o OptimizedAutoscaleOutput) Maximum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscale) int { return v.Maximum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleOutput) Minimum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscale) int { return v.Minimum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscale) int { return v.Version }).(pulumi.IntOutput)
}

type OptimizedAutoscalePtrOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscalePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OptimizedAutoscale)(nil)).Elem()
}

func (o OptimizedAutoscalePtrOutput) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return o
}

func (o OptimizedAutoscalePtrOutput) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return o
}

func (o OptimizedAutoscalePtrOutput) Elem() OptimizedAutoscaleOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) OptimizedAutoscale {
		if v != nil {
			return *v
		}
		var ret OptimizedAutoscale
		return ret
	}).(OptimizedAutoscaleOutput)
}

func (o OptimizedAutoscalePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o OptimizedAutoscalePtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *int {
		if v == nil {
			return nil
		}
		return &v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscalePtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *int {
		if v == nil {
			return nil
		}
		return &v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscalePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *int {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.IntPtrOutput)
}

type OptimizedAutoscaleResponse struct {
	IsEnabled bool `pulumi:"isEnabled"`
	Maximum   int  `pulumi:"maximum"`
	Minimum   int  `pulumi:"minimum"`
	Version   int  `pulumi:"version"`
}

type OptimizedAutoscaleResponseOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscaleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OptimizedAutoscaleResponse)(nil)).Elem()
}

func (o OptimizedAutoscaleResponseOutput) ToOptimizedAutoscaleResponseOutput() OptimizedAutoscaleResponseOutput {
	return o
}

func (o OptimizedAutoscaleResponseOutput) ToOptimizedAutoscaleResponseOutputWithContext(ctx context.Context) OptimizedAutoscaleResponseOutput {
	return o
}

func (o OptimizedAutoscaleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o OptimizedAutoscaleResponseOutput) Maximum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) int { return v.Maximum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleResponseOutput) Minimum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) int { return v.Minimum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleResponseOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) int { return v.Version }).(pulumi.IntOutput)
}

type OptimizedAutoscaleResponsePtrOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscaleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OptimizedAutoscaleResponse)(nil)).Elem()
}

func (o OptimizedAutoscaleResponsePtrOutput) ToOptimizedAutoscaleResponsePtrOutput() OptimizedAutoscaleResponsePtrOutput {
	return o
}

func (o OptimizedAutoscaleResponsePtrOutput) ToOptimizedAutoscaleResponsePtrOutputWithContext(ctx context.Context) OptimizedAutoscaleResponsePtrOutput {
	return o
}

func (o OptimizedAutoscaleResponsePtrOutput) Elem() OptimizedAutoscaleResponseOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) OptimizedAutoscaleResponse {
		if v != nil {
			return *v
		}
		var ret OptimizedAutoscaleResponse
		return ret
	}).(OptimizedAutoscaleResponseOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.IntPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	GroupId                           string                                            `pulumi:"groupId"`
	Id                                string                                            `pulumi:"id"`
	Name                              string                                            `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                            `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                                `pulumi:"systemData"`
	Type                              string                                            `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointPropertyResponse { return v.PrivateEndpoint }).(PrivateEndpointPropertyResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
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

type PrivateEndpointPropertyResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointPropertyResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStateProperty struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput
	ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyOutput
}

type PrivateLinkServiceConnectionStatePropertyArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput)
}

type PrivateLinkServiceConnectionStatePropertyOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string  `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

type TableLevelSharingProperties struct {
	ExternalTablesToExclude    []string `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    []string `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude []string `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude []string `pulumi:"materializedViewsToInclude"`
	TablesToExclude            []string `pulumi:"tablesToExclude"`
	TablesToInclude            []string `pulumi:"tablesToInclude"`
}





type TableLevelSharingPropertiesInput interface {
	pulumi.Input

	ToTableLevelSharingPropertiesOutput() TableLevelSharingPropertiesOutput
	ToTableLevelSharingPropertiesOutputWithContext(context.Context) TableLevelSharingPropertiesOutput
}

type TableLevelSharingPropertiesArgs struct {
	ExternalTablesToExclude    pulumi.StringArrayInput `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    pulumi.StringArrayInput `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude pulumi.StringArrayInput `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude pulumi.StringArrayInput `pulumi:"materializedViewsToInclude"`
	TablesToExclude            pulumi.StringArrayInput `pulumi:"tablesToExclude"`
	TablesToInclude            pulumi.StringArrayInput `pulumi:"tablesToInclude"`
}

func (TableLevelSharingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableLevelSharingProperties)(nil)).Elem()
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesOutput() TableLevelSharingPropertiesOutput {
	return i.ToTableLevelSharingPropertiesOutputWithContext(context.Background())
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesOutputWithContext(ctx context.Context) TableLevelSharingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesOutput)
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return i.ToTableLevelSharingPropertiesPtrOutputWithContext(context.Background())
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesOutput).ToTableLevelSharingPropertiesPtrOutputWithContext(ctx)
}









type TableLevelSharingPropertiesPtrInput interface {
	pulumi.Input

	ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput
	ToTableLevelSharingPropertiesPtrOutputWithContext(context.Context) TableLevelSharingPropertiesPtrOutput
}

type tableLevelSharingPropertiesPtrType TableLevelSharingPropertiesArgs

func TableLevelSharingPropertiesPtr(v *TableLevelSharingPropertiesArgs) TableLevelSharingPropertiesPtrInput {
	return (*tableLevelSharingPropertiesPtrType)(v)
}

func (*tableLevelSharingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableLevelSharingProperties)(nil)).Elem()
}

func (i *tableLevelSharingPropertiesPtrType) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return i.ToTableLevelSharingPropertiesPtrOutputWithContext(context.Background())
}

func (i *tableLevelSharingPropertiesPtrType) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesPtrOutput)
}

type TableLevelSharingPropertiesOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableLevelSharingProperties)(nil)).Elem()
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesOutput() TableLevelSharingPropertiesOutput {
	return o
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesOutputWithContext(ctx context.Context) TableLevelSharingPropertiesOutput {
	return o
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return o.ToTableLevelSharingPropertiesPtrOutputWithContext(context.Background())
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableLevelSharingProperties) *TableLevelSharingProperties {
		return &v
	}).(TableLevelSharingPropertiesPtrOutput)
}

func (o TableLevelSharingPropertiesOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.ExternalTablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.ExternalTablesToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.MaterializedViewsToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.MaterializedViewsToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.TablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.TablesToInclude }).(pulumi.StringArrayOutput)
}

type TableLevelSharingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableLevelSharingProperties)(nil)).Elem()
}

func (o TableLevelSharingPropertiesPtrOutput) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return o
}

func (o TableLevelSharingPropertiesPtrOutput) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return o
}

func (o TableLevelSharingPropertiesPtrOutput) Elem() TableLevelSharingPropertiesOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) TableLevelSharingProperties {
		if v != nil {
			return *v
		}
		var ret TableLevelSharingProperties
		return ret
	}).(TableLevelSharingPropertiesOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.TablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.TablesToInclude
	}).(pulumi.StringArrayOutput)
}

type TableLevelSharingPropertiesResponse struct {
	ExternalTablesToExclude    []string `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    []string `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude []string `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude []string `pulumi:"materializedViewsToInclude"`
	TablesToExclude            []string `pulumi:"tablesToExclude"`
	TablesToInclude            []string `pulumi:"tablesToInclude"`
}

type TableLevelSharingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableLevelSharingPropertiesResponse)(nil)).Elem()
}

func (o TableLevelSharingPropertiesResponseOutput) ToTableLevelSharingPropertiesResponseOutput() TableLevelSharingPropertiesResponseOutput {
	return o
}

func (o TableLevelSharingPropertiesResponseOutput) ToTableLevelSharingPropertiesResponseOutputWithContext(ctx context.Context) TableLevelSharingPropertiesResponseOutput {
	return o
}

func (o TableLevelSharingPropertiesResponseOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.ExternalTablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.ExternalTablesToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.MaterializedViewsToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.MaterializedViewsToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.TablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.TablesToInclude }).(pulumi.StringArrayOutput)
}

type TableLevelSharingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableLevelSharingPropertiesResponse)(nil)).Elem()
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ToTableLevelSharingPropertiesResponsePtrOutput() TableLevelSharingPropertiesResponsePtrOutput {
	return o
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ToTableLevelSharingPropertiesResponsePtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesResponsePtrOutput {
	return o
}

func (o TableLevelSharingPropertiesResponsePtrOutput) Elem() TableLevelSharingPropertiesResponseOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) TableLevelSharingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TableLevelSharingPropertiesResponse
		return ret
	}).(TableLevelSharingPropertiesResponseOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.TablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.TablesToInclude
	}).(pulumi.StringArrayOutput)
}

type TrustedExternalTenant struct {
	Value *string `pulumi:"value"`
}





type TrustedExternalTenantInput interface {
	pulumi.Input

	ToTrustedExternalTenantOutput() TrustedExternalTenantOutput
	ToTrustedExternalTenantOutputWithContext(context.Context) TrustedExternalTenantOutput
}

type TrustedExternalTenantArgs struct {
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (TrustedExternalTenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenant)(nil)).Elem()
}

func (i TrustedExternalTenantArgs) ToTrustedExternalTenantOutput() TrustedExternalTenantOutput {
	return i.ToTrustedExternalTenantOutputWithContext(context.Background())
}

func (i TrustedExternalTenantArgs) ToTrustedExternalTenantOutputWithContext(ctx context.Context) TrustedExternalTenantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedExternalTenantOutput)
}





type TrustedExternalTenantArrayInput interface {
	pulumi.Input

	ToTrustedExternalTenantArrayOutput() TrustedExternalTenantArrayOutput
	ToTrustedExternalTenantArrayOutputWithContext(context.Context) TrustedExternalTenantArrayOutput
}

type TrustedExternalTenantArray []TrustedExternalTenantInput

func (TrustedExternalTenantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenant)(nil)).Elem()
}

func (i TrustedExternalTenantArray) ToTrustedExternalTenantArrayOutput() TrustedExternalTenantArrayOutput {
	return i.ToTrustedExternalTenantArrayOutputWithContext(context.Background())
}

func (i TrustedExternalTenantArray) ToTrustedExternalTenantArrayOutputWithContext(ctx context.Context) TrustedExternalTenantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedExternalTenantArrayOutput)
}

type TrustedExternalTenantOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenant)(nil)).Elem()
}

func (o TrustedExternalTenantOutput) ToTrustedExternalTenantOutput() TrustedExternalTenantOutput {
	return o
}

func (o TrustedExternalTenantOutput) ToTrustedExternalTenantOutputWithContext(ctx context.Context) TrustedExternalTenantOutput {
	return o
}

func (o TrustedExternalTenantOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustedExternalTenant) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type TrustedExternalTenantArrayOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenant)(nil)).Elem()
}

func (o TrustedExternalTenantArrayOutput) ToTrustedExternalTenantArrayOutput() TrustedExternalTenantArrayOutput {
	return o
}

func (o TrustedExternalTenantArrayOutput) ToTrustedExternalTenantArrayOutputWithContext(ctx context.Context) TrustedExternalTenantArrayOutput {
	return o
}

func (o TrustedExternalTenantArrayOutput) Index(i pulumi.IntInput) TrustedExternalTenantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrustedExternalTenant {
		return vs[0].([]TrustedExternalTenant)[vs[1].(int)]
	}).(TrustedExternalTenantOutput)
}

type TrustedExternalTenantResponse struct {
	Value *string `pulumi:"value"`
}

type TrustedExternalTenantResponseOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenantResponse)(nil)).Elem()
}

func (o TrustedExternalTenantResponseOutput) ToTrustedExternalTenantResponseOutput() TrustedExternalTenantResponseOutput {
	return o
}

func (o TrustedExternalTenantResponseOutput) ToTrustedExternalTenantResponseOutputWithContext(ctx context.Context) TrustedExternalTenantResponseOutput {
	return o
}

func (o TrustedExternalTenantResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustedExternalTenantResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type TrustedExternalTenantResponseArrayOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenantResponse)(nil)).Elem()
}

func (o TrustedExternalTenantResponseArrayOutput) ToTrustedExternalTenantResponseArrayOutput() TrustedExternalTenantResponseArrayOutput {
	return o
}

func (o TrustedExternalTenantResponseArrayOutput) ToTrustedExternalTenantResponseArrayOutputWithContext(ctx context.Context) TrustedExternalTenantResponseArrayOutput {
	return o
}

func (o TrustedExternalTenantResponseArrayOutput) Index(i pulumi.IntInput) TrustedExternalTenantResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrustedExternalTenantResponse {
		return vs[0].([]TrustedExternalTenantResponse)[vs[1].(int)]
	}).(TrustedExternalTenantResponseOutput)
}

type VirtualNetworkConfiguration struct {
	DataManagementPublicIpId string `pulumi:"dataManagementPublicIpId"`
	EnginePublicIpId         string `pulumi:"enginePublicIpId"`
	SubnetId                 string `pulumi:"subnetId"`
}





type VirtualNetworkConfigurationInput interface {
	pulumi.Input

	ToVirtualNetworkConfigurationOutput() VirtualNetworkConfigurationOutput
	ToVirtualNetworkConfigurationOutputWithContext(context.Context) VirtualNetworkConfigurationOutput
}

type VirtualNetworkConfigurationArgs struct {
	DataManagementPublicIpId pulumi.StringInput `pulumi:"dataManagementPublicIpId"`
	EnginePublicIpId         pulumi.StringInput `pulumi:"enginePublicIpId"`
	SubnetId                 pulumi.StringInput `pulumi:"subnetId"`
}

func (VirtualNetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfiguration)(nil)).Elem()
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationOutput() VirtualNetworkConfigurationOutput {
	return i.ToVirtualNetworkConfigurationOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationOutputWithContext(ctx context.Context) VirtualNetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationOutput)
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return i.ToVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationOutput).ToVirtualNetworkConfigurationPtrOutputWithContext(ctx)
}









type VirtualNetworkConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput
	ToVirtualNetworkConfigurationPtrOutputWithContext(context.Context) VirtualNetworkConfigurationPtrOutput
}

type virtualNetworkConfigurationPtrType VirtualNetworkConfigurationArgs

func VirtualNetworkConfigurationPtr(v *VirtualNetworkConfigurationArgs) VirtualNetworkConfigurationPtrInput {
	return (*virtualNetworkConfigurationPtrType)(v)
}

func (*virtualNetworkConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfiguration)(nil)).Elem()
}

func (i *virtualNetworkConfigurationPtrType) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return i.ToVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkConfigurationPtrType) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationPtrOutput)
}

type VirtualNetworkConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfiguration)(nil)).Elem()
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationOutput() VirtualNetworkConfigurationOutput {
	return o
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationOutputWithContext(ctx context.Context) VirtualNetworkConfigurationOutput {
	return o
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return o.ToVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkConfiguration) *VirtualNetworkConfiguration {
		return &v
	}).(VirtualNetworkConfigurationPtrOutput)
}

func (o VirtualNetworkConfigurationOutput) DataManagementPublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfiguration) string { return v.DataManagementPublicIpId }).(pulumi.StringOutput)
}

func (o VirtualNetworkConfigurationOutput) EnginePublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfiguration) string { return v.EnginePublicIpId }).(pulumi.StringOutput)
}

func (o VirtualNetworkConfigurationOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfiguration) string { return v.SubnetId }).(pulumi.StringOutput)
}

type VirtualNetworkConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfiguration)(nil)).Elem()
}

func (o VirtualNetworkConfigurationPtrOutput) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return o
}

func (o VirtualNetworkConfigurationPtrOutput) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return o
}

func (o VirtualNetworkConfigurationPtrOutput) Elem() VirtualNetworkConfigurationOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) VirtualNetworkConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkConfiguration
		return ret
	}).(VirtualNetworkConfigurationOutput)
}

func (o VirtualNetworkConfigurationPtrOutput) DataManagementPublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.DataManagementPublicIpId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationPtrOutput) EnginePublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.EnginePublicIpId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationPtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigurationResponse struct {
	DataManagementPublicIpId string `pulumi:"dataManagementPublicIpId"`
	EnginePublicIpId         string `pulumi:"enginePublicIpId"`
	SubnetId                 string `pulumi:"subnetId"`
}

type VirtualNetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkConfigurationResponseOutput) ToVirtualNetworkConfigurationResponseOutput() VirtualNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkConfigurationResponseOutput) ToVirtualNetworkConfigurationResponseOutputWithContext(ctx context.Context) VirtualNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkConfigurationResponseOutput) DataManagementPublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfigurationResponse) string { return v.DataManagementPublicIpId }).(pulumi.StringOutput)
}

func (o VirtualNetworkConfigurationResponseOutput) EnginePublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfigurationResponse) string { return v.EnginePublicIpId }).(pulumi.StringOutput)
}

func (o VirtualNetworkConfigurationResponseOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfigurationResponse) string { return v.SubnetId }).(pulumi.StringOutput)
}

type VirtualNetworkConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkConfigurationResponsePtrOutput) ToVirtualNetworkConfigurationResponsePtrOutput() VirtualNetworkConfigurationResponsePtrOutput {
	return o
}

func (o VirtualNetworkConfigurationResponsePtrOutput) ToVirtualNetworkConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationResponsePtrOutput {
	return o
}

func (o VirtualNetworkConfigurationResponsePtrOutput) Elem() VirtualNetworkConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) VirtualNetworkConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkConfigurationResponse
		return ret
	}).(VirtualNetworkConfigurationResponseOutput)
}

func (o VirtualNetworkConfigurationResponsePtrOutput) DataManagementPublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataManagementPublicIpId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationResponsePtrOutput) EnginePublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EnginePublicIpId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SubnetId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AcceptedAudiencesOutput{})
	pulumi.RegisterOutputType(AcceptedAudiencesArrayOutput{})
	pulumi.RegisterOutputType(AcceptedAudiencesResponseOutput{})
	pulumi.RegisterOutputType(AcceptedAudiencesResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureSkuOutput{})
	pulumi.RegisterOutputType(AzureSkuResponseOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalResponseOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalResponseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseStatisticsResponseOutput{})
	pulumi.RegisterOutputType(FollowerDatabaseDefinitionResponseOutput{})
	pulumi.RegisterOutputType(FollowerDatabaseDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LanguageExtensionOutput{})
	pulumi.RegisterOutputType(LanguageExtensionArrayOutput{})
	pulumi.RegisterOutputType(LanguageExtensionResponseOutput{})
	pulumi.RegisterOutputType(LanguageExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(LanguageExtensionsListOutput{})
	pulumi.RegisterOutputType(LanguageExtensionsListPtrOutput{})
	pulumi.RegisterOutputType(LanguageExtensionsListResponseOutput{})
	pulumi.RegisterOutputType(LanguageExtensionsListResponsePtrOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscaleOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscalePtrOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscaleResponseOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscaleResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantArrayOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantResponseOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationResponsePtrOutput{})
}
