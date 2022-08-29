


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AbsoluteDeleteOption struct {
	Duration   string `pulumi:"duration"`
	ObjectType string `pulumi:"objectType"`
}

type AbsoluteDeleteOptionResponse struct {
	Duration   string `pulumi:"duration"`
	ObjectType string `pulumi:"objectType"`
}

type AdhocBasedTaggingCriteria struct {
	TagInfo *RetentionTag `pulumi:"tagInfo"`
}

type AdhocBasedTaggingCriteriaResponse struct {
	TagInfo *RetentionTagResponse `pulumi:"tagInfo"`
}

type AdhocBasedTriggerContext struct {
	ObjectType      string                    `pulumi:"objectType"`
	TaggingCriteria AdhocBasedTaggingCriteria `pulumi:"taggingCriteria"`
}

type AdhocBasedTriggerContextResponse struct {
	ObjectType      string                            `pulumi:"objectType"`
	TaggingCriteria AdhocBasedTaggingCriteriaResponse `pulumi:"taggingCriteria"`
}

type AzureBackupParams struct {
	BackupType string `pulumi:"backupType"`
	ObjectType string `pulumi:"objectType"`
}

type AzureBackupParamsResponse struct {
	BackupType string `pulumi:"backupType"`
	ObjectType string `pulumi:"objectType"`
}

type AzureBackupRule struct {
	BackupParameters *AzureBackupParams `pulumi:"backupParameters"`
	DataStore        DataStoreInfoBase  `pulumi:"dataStore"`
	Name             string             `pulumi:"name"`
	ObjectType       string             `pulumi:"objectType"`
	Trigger          interface{}        `pulumi:"trigger"`
}

type AzureBackupRuleResponse struct {
	BackupParameters *AzureBackupParamsResponse `pulumi:"backupParameters"`
	DataStore        DataStoreInfoBaseResponse  `pulumi:"dataStore"`
	Name             string                     `pulumi:"name"`
	ObjectType       string                     `pulumi:"objectType"`
	Trigger          interface{}                `pulumi:"trigger"`
}

type AzureOperationalStoreParameters struct {
	DataStoreType   string  `pulumi:"dataStoreType"`
	ObjectType      string  `pulumi:"objectType"`
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}





type AzureOperationalStoreParametersInput interface {
	pulumi.Input

	ToAzureOperationalStoreParametersOutput() AzureOperationalStoreParametersOutput
	ToAzureOperationalStoreParametersOutputWithContext(context.Context) AzureOperationalStoreParametersOutput
}

type AzureOperationalStoreParametersArgs struct {
	DataStoreType   pulumi.StringInput    `pulumi:"dataStoreType"`
	ObjectType      pulumi.StringInput    `pulumi:"objectType"`
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
}

func (AzureOperationalStoreParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureOperationalStoreParameters)(nil)).Elem()
}

func (i AzureOperationalStoreParametersArgs) ToAzureOperationalStoreParametersOutput() AzureOperationalStoreParametersOutput {
	return i.ToAzureOperationalStoreParametersOutputWithContext(context.Background())
}

func (i AzureOperationalStoreParametersArgs) ToAzureOperationalStoreParametersOutputWithContext(ctx context.Context) AzureOperationalStoreParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureOperationalStoreParametersOutput)
}





type AzureOperationalStoreParametersArrayInput interface {
	pulumi.Input

	ToAzureOperationalStoreParametersArrayOutput() AzureOperationalStoreParametersArrayOutput
	ToAzureOperationalStoreParametersArrayOutputWithContext(context.Context) AzureOperationalStoreParametersArrayOutput
}

type AzureOperationalStoreParametersArray []AzureOperationalStoreParametersInput

func (AzureOperationalStoreParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureOperationalStoreParameters)(nil)).Elem()
}

func (i AzureOperationalStoreParametersArray) ToAzureOperationalStoreParametersArrayOutput() AzureOperationalStoreParametersArrayOutput {
	return i.ToAzureOperationalStoreParametersArrayOutputWithContext(context.Background())
}

func (i AzureOperationalStoreParametersArray) ToAzureOperationalStoreParametersArrayOutputWithContext(ctx context.Context) AzureOperationalStoreParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureOperationalStoreParametersArrayOutput)
}

type AzureOperationalStoreParametersOutput struct{ *pulumi.OutputState }

func (AzureOperationalStoreParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureOperationalStoreParameters)(nil)).Elem()
}

func (o AzureOperationalStoreParametersOutput) ToAzureOperationalStoreParametersOutput() AzureOperationalStoreParametersOutput {
	return o
}

func (o AzureOperationalStoreParametersOutput) ToAzureOperationalStoreParametersOutputWithContext(ctx context.Context) AzureOperationalStoreParametersOutput {
	return o
}

func (o AzureOperationalStoreParametersOutput) DataStoreType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureOperationalStoreParameters) string { return v.DataStoreType }).(pulumi.StringOutput)
}

func (o AzureOperationalStoreParametersOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureOperationalStoreParameters) string { return v.ObjectType }).(pulumi.StringOutput)
}

func (o AzureOperationalStoreParametersOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureOperationalStoreParameters) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

type AzureOperationalStoreParametersArrayOutput struct{ *pulumi.OutputState }

func (AzureOperationalStoreParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureOperationalStoreParameters)(nil)).Elem()
}

func (o AzureOperationalStoreParametersArrayOutput) ToAzureOperationalStoreParametersArrayOutput() AzureOperationalStoreParametersArrayOutput {
	return o
}

func (o AzureOperationalStoreParametersArrayOutput) ToAzureOperationalStoreParametersArrayOutputWithContext(ctx context.Context) AzureOperationalStoreParametersArrayOutput {
	return o
}

func (o AzureOperationalStoreParametersArrayOutput) Index(i pulumi.IntInput) AzureOperationalStoreParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureOperationalStoreParameters {
		return vs[0].([]AzureOperationalStoreParameters)[vs[1].(int)]
	}).(AzureOperationalStoreParametersOutput)
}

type AzureOperationalStoreParametersResponse struct {
	DataStoreType   string  `pulumi:"dataStoreType"`
	ObjectType      string  `pulumi:"objectType"`
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}

type AzureOperationalStoreParametersResponseOutput struct{ *pulumi.OutputState }

func (AzureOperationalStoreParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureOperationalStoreParametersResponse)(nil)).Elem()
}

func (o AzureOperationalStoreParametersResponseOutput) ToAzureOperationalStoreParametersResponseOutput() AzureOperationalStoreParametersResponseOutput {
	return o
}

func (o AzureOperationalStoreParametersResponseOutput) ToAzureOperationalStoreParametersResponseOutputWithContext(ctx context.Context) AzureOperationalStoreParametersResponseOutput {
	return o
}

func (o AzureOperationalStoreParametersResponseOutput) DataStoreType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureOperationalStoreParametersResponse) string { return v.DataStoreType }).(pulumi.StringOutput)
}

func (o AzureOperationalStoreParametersResponseOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureOperationalStoreParametersResponse) string { return v.ObjectType }).(pulumi.StringOutput)
}

func (o AzureOperationalStoreParametersResponseOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureOperationalStoreParametersResponse) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

type AzureOperationalStoreParametersResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureOperationalStoreParametersResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureOperationalStoreParametersResponse)(nil)).Elem()
}

func (o AzureOperationalStoreParametersResponseArrayOutput) ToAzureOperationalStoreParametersResponseArrayOutput() AzureOperationalStoreParametersResponseArrayOutput {
	return o
}

func (o AzureOperationalStoreParametersResponseArrayOutput) ToAzureOperationalStoreParametersResponseArrayOutputWithContext(ctx context.Context) AzureOperationalStoreParametersResponseArrayOutput {
	return o
}

func (o AzureOperationalStoreParametersResponseArrayOutput) Index(i pulumi.IntInput) AzureOperationalStoreParametersResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureOperationalStoreParametersResponse {
		return vs[0].([]AzureOperationalStoreParametersResponse)[vs[1].(int)]
	}).(AzureOperationalStoreParametersResponseOutput)
}

type AzureRetentionRule struct {
	IsDefault  *bool             `pulumi:"isDefault"`
	Lifecycles []SourceLifeCycle `pulumi:"lifecycles"`
	Name       string            `pulumi:"name"`
	ObjectType string            `pulumi:"objectType"`
}

type AzureRetentionRuleResponse struct {
	IsDefault  *bool                     `pulumi:"isDefault"`
	Lifecycles []SourceLifeCycleResponse `pulumi:"lifecycles"`
	Name       string                    `pulumi:"name"`
	ObjectType string                    `pulumi:"objectType"`
}

type BackupInstanceType struct {
	DataSourceInfo    Datasource     `pulumi:"dataSourceInfo"`
	DataSourceSetInfo *DatasourceSet `pulumi:"dataSourceSetInfo"`
	FriendlyName      *string        `pulumi:"friendlyName"`
	ObjectType        string         `pulumi:"objectType"`
	PolicyInfo        PolicyInfo     `pulumi:"policyInfo"`
}





type BackupInstanceTypeInput interface {
	pulumi.Input

	ToBackupInstanceTypeOutput() BackupInstanceTypeOutput
	ToBackupInstanceTypeOutputWithContext(context.Context) BackupInstanceTypeOutput
}

type BackupInstanceTypeArgs struct {
	DataSourceInfo    DatasourceInput       `pulumi:"dataSourceInfo"`
	DataSourceSetInfo DatasourceSetPtrInput `pulumi:"dataSourceSetInfo"`
	FriendlyName      pulumi.StringPtrInput `pulumi:"friendlyName"`
	ObjectType        pulumi.StringInput    `pulumi:"objectType"`
	PolicyInfo        PolicyInfoInput       `pulumi:"policyInfo"`
}

func (BackupInstanceTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupInstanceType)(nil)).Elem()
}

func (i BackupInstanceTypeArgs) ToBackupInstanceTypeOutput() BackupInstanceTypeOutput {
	return i.ToBackupInstanceTypeOutputWithContext(context.Background())
}

func (i BackupInstanceTypeArgs) ToBackupInstanceTypeOutputWithContext(ctx context.Context) BackupInstanceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceTypeOutput)
}

func (i BackupInstanceTypeArgs) ToBackupInstanceTypePtrOutput() BackupInstanceTypePtrOutput {
	return i.ToBackupInstanceTypePtrOutputWithContext(context.Background())
}

func (i BackupInstanceTypeArgs) ToBackupInstanceTypePtrOutputWithContext(ctx context.Context) BackupInstanceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceTypeOutput).ToBackupInstanceTypePtrOutputWithContext(ctx)
}









type BackupInstanceTypePtrInput interface {
	pulumi.Input

	ToBackupInstanceTypePtrOutput() BackupInstanceTypePtrOutput
	ToBackupInstanceTypePtrOutputWithContext(context.Context) BackupInstanceTypePtrOutput
}

type backupInstanceTypePtrType BackupInstanceTypeArgs

func BackupInstanceTypePtr(v *BackupInstanceTypeArgs) BackupInstanceTypePtrInput {
	return (*backupInstanceTypePtrType)(v)
}

func (*backupInstanceTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupInstanceType)(nil)).Elem()
}

func (i *backupInstanceTypePtrType) ToBackupInstanceTypePtrOutput() BackupInstanceTypePtrOutput {
	return i.ToBackupInstanceTypePtrOutputWithContext(context.Background())
}

func (i *backupInstanceTypePtrType) ToBackupInstanceTypePtrOutputWithContext(ctx context.Context) BackupInstanceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceTypePtrOutput)
}

type BackupInstanceTypeOutput struct{ *pulumi.OutputState }

func (BackupInstanceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupInstanceType)(nil)).Elem()
}

func (o BackupInstanceTypeOutput) ToBackupInstanceTypeOutput() BackupInstanceTypeOutput {
	return o
}

func (o BackupInstanceTypeOutput) ToBackupInstanceTypeOutputWithContext(ctx context.Context) BackupInstanceTypeOutput {
	return o
}

func (o BackupInstanceTypeOutput) ToBackupInstanceTypePtrOutput() BackupInstanceTypePtrOutput {
	return o.ToBackupInstanceTypePtrOutputWithContext(context.Background())
}

func (o BackupInstanceTypeOutput) ToBackupInstanceTypePtrOutputWithContext(ctx context.Context) BackupInstanceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupInstanceType) *BackupInstanceType {
		return &v
	}).(BackupInstanceTypePtrOutput)
}

func (o BackupInstanceTypeOutput) DataSourceInfo() DatasourceOutput {
	return o.ApplyT(func(v BackupInstanceType) Datasource { return v.DataSourceInfo }).(DatasourceOutput)
}

func (o BackupInstanceTypeOutput) DataSourceSetInfo() DatasourceSetPtrOutput {
	return o.ApplyT(func(v BackupInstanceType) *DatasourceSet { return v.DataSourceSetInfo }).(DatasourceSetPtrOutput)
}

func (o BackupInstanceTypeOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupInstanceType) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o BackupInstanceTypeOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v BackupInstanceType) string { return v.ObjectType }).(pulumi.StringOutput)
}

func (o BackupInstanceTypeOutput) PolicyInfo() PolicyInfoOutput {
	return o.ApplyT(func(v BackupInstanceType) PolicyInfo { return v.PolicyInfo }).(PolicyInfoOutput)
}

type BackupInstanceTypePtrOutput struct{ *pulumi.OutputState }

func (BackupInstanceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupInstanceType)(nil)).Elem()
}

func (o BackupInstanceTypePtrOutput) ToBackupInstanceTypePtrOutput() BackupInstanceTypePtrOutput {
	return o
}

func (o BackupInstanceTypePtrOutput) ToBackupInstanceTypePtrOutputWithContext(ctx context.Context) BackupInstanceTypePtrOutput {
	return o
}

func (o BackupInstanceTypePtrOutput) Elem() BackupInstanceTypeOutput {
	return o.ApplyT(func(v *BackupInstanceType) BackupInstanceType {
		if v != nil {
			return *v
		}
		var ret BackupInstanceType
		return ret
	}).(BackupInstanceTypeOutput)
}

func (o BackupInstanceTypePtrOutput) DataSourceInfo() DatasourcePtrOutput {
	return o.ApplyT(func(v *BackupInstanceType) *Datasource {
		if v == nil {
			return nil
		}
		return &v.DataSourceInfo
	}).(DatasourcePtrOutput)
}

func (o BackupInstanceTypePtrOutput) DataSourceSetInfo() DatasourceSetPtrOutput {
	return o.ApplyT(func(v *BackupInstanceType) *DatasourceSet {
		if v == nil {
			return nil
		}
		return v.DataSourceSetInfo
	}).(DatasourceSetPtrOutput)
}

func (o BackupInstanceTypePtrOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupInstanceType) *string {
		if v == nil {
			return nil
		}
		return v.FriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o BackupInstanceTypePtrOutput) ObjectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupInstanceType) *string {
		if v == nil {
			return nil
		}
		return &v.ObjectType
	}).(pulumi.StringPtrOutput)
}

func (o BackupInstanceTypePtrOutput) PolicyInfo() PolicyInfoPtrOutput {
	return o.ApplyT(func(v *BackupInstanceType) *PolicyInfo {
		if v == nil {
			return nil
		}
		return &v.PolicyInfo
	}).(PolicyInfoPtrOutput)
}

type BackupInstanceResponse struct {
	CurrentProtectionState string                          `pulumi:"currentProtectionState"`
	DataSourceInfo         DatasourceResponse              `pulumi:"dataSourceInfo"`
	DataSourceSetInfo      *DatasourceSetResponse          `pulumi:"dataSourceSetInfo"`
	FriendlyName           *string                         `pulumi:"friendlyName"`
	ObjectType             string                          `pulumi:"objectType"`
	PolicyInfo             PolicyInfoResponse              `pulumi:"policyInfo"`
	ProtectionErrorDetails UserFacingErrorResponse         `pulumi:"protectionErrorDetails"`
	ProtectionStatus       ProtectionStatusDetailsResponse `pulumi:"protectionStatus"`
	ProvisioningState      string                          `pulumi:"provisioningState"`
}

type BackupInstanceResponseOutput struct{ *pulumi.OutputState }

func (BackupInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupInstanceResponse)(nil)).Elem()
}

func (o BackupInstanceResponseOutput) ToBackupInstanceResponseOutput() BackupInstanceResponseOutput {
	return o
}

func (o BackupInstanceResponseOutput) ToBackupInstanceResponseOutputWithContext(ctx context.Context) BackupInstanceResponseOutput {
	return o
}

func (o BackupInstanceResponseOutput) CurrentProtectionState() pulumi.StringOutput {
	return o.ApplyT(func(v BackupInstanceResponse) string { return v.CurrentProtectionState }).(pulumi.StringOutput)
}

func (o BackupInstanceResponseOutput) DataSourceInfo() DatasourceResponseOutput {
	return o.ApplyT(func(v BackupInstanceResponse) DatasourceResponse { return v.DataSourceInfo }).(DatasourceResponseOutput)
}

func (o BackupInstanceResponseOutput) DataSourceSetInfo() DatasourceSetResponsePtrOutput {
	return o.ApplyT(func(v BackupInstanceResponse) *DatasourceSetResponse { return v.DataSourceSetInfo }).(DatasourceSetResponsePtrOutput)
}

func (o BackupInstanceResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupInstanceResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o BackupInstanceResponseOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v BackupInstanceResponse) string { return v.ObjectType }).(pulumi.StringOutput)
}

func (o BackupInstanceResponseOutput) PolicyInfo() PolicyInfoResponseOutput {
	return o.ApplyT(func(v BackupInstanceResponse) PolicyInfoResponse { return v.PolicyInfo }).(PolicyInfoResponseOutput)
}

func (o BackupInstanceResponseOutput) ProtectionErrorDetails() UserFacingErrorResponseOutput {
	return o.ApplyT(func(v BackupInstanceResponse) UserFacingErrorResponse { return v.ProtectionErrorDetails }).(UserFacingErrorResponseOutput)
}

func (o BackupInstanceResponseOutput) ProtectionStatus() ProtectionStatusDetailsResponseOutput {
	return o.ApplyT(func(v BackupInstanceResponse) ProtectionStatusDetailsResponse { return v.ProtectionStatus }).(ProtectionStatusDetailsResponseOutput)
}

func (o BackupInstanceResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v BackupInstanceResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type BackupPolicyType struct {
	DatasourceTypes []string      `pulumi:"datasourceTypes"`
	ObjectType      string        `pulumi:"objectType"`
	PolicyRules     []interface{} `pulumi:"policyRules"`
}





type BackupPolicyTypeInput interface {
	pulumi.Input

	ToBackupPolicyTypeOutput() BackupPolicyTypeOutput
	ToBackupPolicyTypeOutputWithContext(context.Context) BackupPolicyTypeOutput
}

type BackupPolicyTypeArgs struct {
	DatasourceTypes pulumi.StringArrayInput `pulumi:"datasourceTypes"`
	ObjectType      pulumi.StringInput      `pulumi:"objectType"`
	PolicyRules     pulumi.ArrayInput       `pulumi:"policyRules"`
}

func (BackupPolicyTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyType)(nil)).Elem()
}

func (i BackupPolicyTypeArgs) ToBackupPolicyTypeOutput() BackupPolicyTypeOutput {
	return i.ToBackupPolicyTypeOutputWithContext(context.Background())
}

func (i BackupPolicyTypeArgs) ToBackupPolicyTypeOutputWithContext(ctx context.Context) BackupPolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyTypeOutput)
}

func (i BackupPolicyTypeArgs) ToBackupPolicyTypePtrOutput() BackupPolicyTypePtrOutput {
	return i.ToBackupPolicyTypePtrOutputWithContext(context.Background())
}

func (i BackupPolicyTypeArgs) ToBackupPolicyTypePtrOutputWithContext(ctx context.Context) BackupPolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyTypeOutput).ToBackupPolicyTypePtrOutputWithContext(ctx)
}









type BackupPolicyTypePtrInput interface {
	pulumi.Input

	ToBackupPolicyTypePtrOutput() BackupPolicyTypePtrOutput
	ToBackupPolicyTypePtrOutputWithContext(context.Context) BackupPolicyTypePtrOutput
}

type backupPolicyTypePtrType BackupPolicyTypeArgs

func BackupPolicyTypePtr(v *BackupPolicyTypeArgs) BackupPolicyTypePtrInput {
	return (*backupPolicyTypePtrType)(v)
}

func (*backupPolicyTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicyType)(nil)).Elem()
}

func (i *backupPolicyTypePtrType) ToBackupPolicyTypePtrOutput() BackupPolicyTypePtrOutput {
	return i.ToBackupPolicyTypePtrOutputWithContext(context.Background())
}

func (i *backupPolicyTypePtrType) ToBackupPolicyTypePtrOutputWithContext(ctx context.Context) BackupPolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyTypePtrOutput)
}

type BackupPolicyTypeOutput struct{ *pulumi.OutputState }

func (BackupPolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyType)(nil)).Elem()
}

func (o BackupPolicyTypeOutput) ToBackupPolicyTypeOutput() BackupPolicyTypeOutput {
	return o
}

func (o BackupPolicyTypeOutput) ToBackupPolicyTypeOutputWithContext(ctx context.Context) BackupPolicyTypeOutput {
	return o
}

func (o BackupPolicyTypeOutput) ToBackupPolicyTypePtrOutput() BackupPolicyTypePtrOutput {
	return o.ToBackupPolicyTypePtrOutputWithContext(context.Background())
}

func (o BackupPolicyTypeOutput) ToBackupPolicyTypePtrOutputWithContext(ctx context.Context) BackupPolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupPolicyType) *BackupPolicyType {
		return &v
	}).(BackupPolicyTypePtrOutput)
}

func (o BackupPolicyTypeOutput) DatasourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackupPolicyType) []string { return v.DatasourceTypes }).(pulumi.StringArrayOutput)
}

func (o BackupPolicyTypeOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v BackupPolicyType) string { return v.ObjectType }).(pulumi.StringOutput)
}

func (o BackupPolicyTypeOutput) PolicyRules() pulumi.ArrayOutput {
	return o.ApplyT(func(v BackupPolicyType) []interface{} { return v.PolicyRules }).(pulumi.ArrayOutput)
}

type BackupPolicyTypePtrOutput struct{ *pulumi.OutputState }

func (BackupPolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicyType)(nil)).Elem()
}

func (o BackupPolicyTypePtrOutput) ToBackupPolicyTypePtrOutput() BackupPolicyTypePtrOutput {
	return o
}

func (o BackupPolicyTypePtrOutput) ToBackupPolicyTypePtrOutputWithContext(ctx context.Context) BackupPolicyTypePtrOutput {
	return o
}

func (o BackupPolicyTypePtrOutput) Elem() BackupPolicyTypeOutput {
	return o.ApplyT(func(v *BackupPolicyType) BackupPolicyType {
		if v != nil {
			return *v
		}
		var ret BackupPolicyType
		return ret
	}).(BackupPolicyTypeOutput)
}

func (o BackupPolicyTypePtrOutput) DatasourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackupPolicyType) []string {
		if v == nil {
			return nil
		}
		return v.DatasourceTypes
	}).(pulumi.StringArrayOutput)
}

func (o BackupPolicyTypePtrOutput) ObjectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicyType) *string {
		if v == nil {
			return nil
		}
		return &v.ObjectType
	}).(pulumi.StringPtrOutput)
}

func (o BackupPolicyTypePtrOutput) PolicyRules() pulumi.ArrayOutput {
	return o.ApplyT(func(v *BackupPolicyType) []interface{} {
		if v == nil {
			return nil
		}
		return v.PolicyRules
	}).(pulumi.ArrayOutput)
}

type BackupPolicyResponse struct {
	DatasourceTypes []string      `pulumi:"datasourceTypes"`
	ObjectType      string        `pulumi:"objectType"`
	PolicyRules     []interface{} `pulumi:"policyRules"`
}

type BackupPolicyResponseOutput struct{ *pulumi.OutputState }

func (BackupPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyResponse)(nil)).Elem()
}

func (o BackupPolicyResponseOutput) ToBackupPolicyResponseOutput() BackupPolicyResponseOutput {
	return o
}

func (o BackupPolicyResponseOutput) ToBackupPolicyResponseOutputWithContext(ctx context.Context) BackupPolicyResponseOutput {
	return o
}

func (o BackupPolicyResponseOutput) DatasourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackupPolicyResponse) []string { return v.DatasourceTypes }).(pulumi.StringArrayOutput)
}

func (o BackupPolicyResponseOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v BackupPolicyResponse) string { return v.ObjectType }).(pulumi.StringOutput)
}

func (o BackupPolicyResponseOutput) PolicyRules() pulumi.ArrayOutput {
	return o.ApplyT(func(v BackupPolicyResponse) []interface{} { return v.PolicyRules }).(pulumi.ArrayOutput)
}

type BackupSchedule struct {
	RepeatingTimeIntervals []string `pulumi:"repeatingTimeIntervals"`
}

type BackupScheduleResponse struct {
	RepeatingTimeIntervals []string `pulumi:"repeatingTimeIntervals"`
}

type BackupVaultType struct {
	StorageSettings []StorageSetting `pulumi:"storageSettings"`
}





type BackupVaultTypeInput interface {
	pulumi.Input

	ToBackupVaultTypeOutput() BackupVaultTypeOutput
	ToBackupVaultTypeOutputWithContext(context.Context) BackupVaultTypeOutput
}

type BackupVaultTypeArgs struct {
	StorageSettings StorageSettingArrayInput `pulumi:"storageSettings"`
}

func (BackupVaultTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupVaultType)(nil)).Elem()
}

func (i BackupVaultTypeArgs) ToBackupVaultTypeOutput() BackupVaultTypeOutput {
	return i.ToBackupVaultTypeOutputWithContext(context.Background())
}

func (i BackupVaultTypeArgs) ToBackupVaultTypeOutputWithContext(ctx context.Context) BackupVaultTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultTypeOutput)
}

type BackupVaultTypeOutput struct{ *pulumi.OutputState }

func (BackupVaultTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupVaultType)(nil)).Elem()
}

func (o BackupVaultTypeOutput) ToBackupVaultTypeOutput() BackupVaultTypeOutput {
	return o
}

func (o BackupVaultTypeOutput) ToBackupVaultTypeOutputWithContext(ctx context.Context) BackupVaultTypeOutput {
	return o
}

func (o BackupVaultTypeOutput) StorageSettings() StorageSettingArrayOutput {
	return o.ApplyT(func(v BackupVaultType) []StorageSetting { return v.StorageSettings }).(StorageSettingArrayOutput)
}

type BackupVaultResponse struct {
	ProvisioningState string                   `pulumi:"provisioningState"`
	StorageSettings   []StorageSettingResponse `pulumi:"storageSettings"`
}

type BackupVaultResponseOutput struct{ *pulumi.OutputState }

func (BackupVaultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupVaultResponse)(nil)).Elem()
}

func (o BackupVaultResponseOutput) ToBackupVaultResponseOutput() BackupVaultResponseOutput {
	return o
}

func (o BackupVaultResponseOutput) ToBackupVaultResponseOutputWithContext(ctx context.Context) BackupVaultResponseOutput {
	return o
}

func (o BackupVaultResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v BackupVaultResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o BackupVaultResponseOutput) StorageSettings() StorageSettingResponseArrayOutput {
	return o.ApplyT(func(v BackupVaultResponse) []StorageSettingResponse { return v.StorageSettings }).(StorageSettingResponseArrayOutput)
}

type CopyOnExpiryOption struct {
	ObjectType string `pulumi:"objectType"`
}

type CopyOnExpiryOptionResponse struct {
	ObjectType string `pulumi:"objectType"`
}

type CustomCopyOption struct {
	Duration   *string `pulumi:"duration"`
	ObjectType string  `pulumi:"objectType"`
}

type CustomCopyOptionResponse struct {
	Duration   *string `pulumi:"duration"`
	ObjectType string  `pulumi:"objectType"`
}

type DataStoreInfoBase struct {
	DataStoreType string `pulumi:"dataStoreType"`
	ObjectType    string `pulumi:"objectType"`
}

type DataStoreInfoBaseResponse struct {
	DataStoreType string `pulumi:"dataStoreType"`
	ObjectType    string `pulumi:"objectType"`
}

type Datasource struct {
	DatasourceType   *string `pulumi:"datasourceType"`
	ObjectType       *string `pulumi:"objectType"`
	ResourceID       string  `pulumi:"resourceID"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceName     *string `pulumi:"resourceName"`
	ResourceType     *string `pulumi:"resourceType"`
	ResourceUri      *string `pulumi:"resourceUri"`
}





type DatasourceInput interface {
	pulumi.Input

	ToDatasourceOutput() DatasourceOutput
	ToDatasourceOutputWithContext(context.Context) DatasourceOutput
}

type DatasourceArgs struct {
	DatasourceType   pulumi.StringPtrInput `pulumi:"datasourceType"`
	ObjectType       pulumi.StringPtrInput `pulumi:"objectType"`
	ResourceID       pulumi.StringInput    `pulumi:"resourceID"`
	ResourceLocation pulumi.StringPtrInput `pulumi:"resourceLocation"`
	ResourceName     pulumi.StringPtrInput `pulumi:"resourceName"`
	ResourceType     pulumi.StringPtrInput `pulumi:"resourceType"`
	ResourceUri      pulumi.StringPtrInput `pulumi:"resourceUri"`
}

func (DatasourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Datasource)(nil)).Elem()
}

func (i DatasourceArgs) ToDatasourceOutput() DatasourceOutput {
	return i.ToDatasourceOutputWithContext(context.Background())
}

func (i DatasourceArgs) ToDatasourceOutputWithContext(ctx context.Context) DatasourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceOutput)
}

func (i DatasourceArgs) ToDatasourcePtrOutput() DatasourcePtrOutput {
	return i.ToDatasourcePtrOutputWithContext(context.Background())
}

func (i DatasourceArgs) ToDatasourcePtrOutputWithContext(ctx context.Context) DatasourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceOutput).ToDatasourcePtrOutputWithContext(ctx)
}









type DatasourcePtrInput interface {
	pulumi.Input

	ToDatasourcePtrOutput() DatasourcePtrOutput
	ToDatasourcePtrOutputWithContext(context.Context) DatasourcePtrOutput
}

type datasourcePtrType DatasourceArgs

func DatasourcePtr(v *DatasourceArgs) DatasourcePtrInput {
	return (*datasourcePtrType)(v)
}

func (*datasourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Datasource)(nil)).Elem()
}

func (i *datasourcePtrType) ToDatasourcePtrOutput() DatasourcePtrOutput {
	return i.ToDatasourcePtrOutputWithContext(context.Background())
}

func (i *datasourcePtrType) ToDatasourcePtrOutputWithContext(ctx context.Context) DatasourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourcePtrOutput)
}

type DatasourceOutput struct{ *pulumi.OutputState }

func (DatasourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Datasource)(nil)).Elem()
}

func (o DatasourceOutput) ToDatasourceOutput() DatasourceOutput {
	return o
}

func (o DatasourceOutput) ToDatasourceOutputWithContext(ctx context.Context) DatasourceOutput {
	return o
}

func (o DatasourceOutput) ToDatasourcePtrOutput() DatasourcePtrOutput {
	return o.ToDatasourcePtrOutputWithContext(context.Background())
}

func (o DatasourceOutput) ToDatasourcePtrOutputWithContext(ctx context.Context) DatasourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Datasource) *Datasource {
		return &v
	}).(DatasourcePtrOutput)
}

func (o DatasourceOutput) DatasourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Datasource) *string { return v.DatasourceType }).(pulumi.StringPtrOutput)
}

func (o DatasourceOutput) ObjectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Datasource) *string { return v.ObjectType }).(pulumi.StringPtrOutput)
}

func (o DatasourceOutput) ResourceID() pulumi.StringOutput {
	return o.ApplyT(func(v Datasource) string { return v.ResourceID }).(pulumi.StringOutput)
}

func (o DatasourceOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Datasource) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o DatasourceOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Datasource) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o DatasourceOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Datasource) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o DatasourceOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Datasource) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

type DatasourcePtrOutput struct{ *pulumi.OutputState }

func (DatasourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datasource)(nil)).Elem()
}

func (o DatasourcePtrOutput) ToDatasourcePtrOutput() DatasourcePtrOutput {
	return o
}

func (o DatasourcePtrOutput) ToDatasourcePtrOutputWithContext(ctx context.Context) DatasourcePtrOutput {
	return o
}

func (o DatasourcePtrOutput) Elem() DatasourceOutput {
	return o.ApplyT(func(v *Datasource) Datasource {
		if v != nil {
			return *v
		}
		var ret Datasource
		return ret
	}).(DatasourceOutput)
}

func (o DatasourcePtrOutput) DatasourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) *string {
		if v == nil {
			return nil
		}
		return v.DatasourceType
	}).(pulumi.StringPtrOutput)
}

func (o DatasourcePtrOutput) ObjectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) *string {
		if v == nil {
			return nil
		}
		return v.ObjectType
	}).(pulumi.StringPtrOutput)
}

func (o DatasourcePtrOutput) ResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceID
	}).(pulumi.StringPtrOutput)
}

func (o DatasourcePtrOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) *string {
		if v == nil {
			return nil
		}
		return v.ResourceLocation
	}).(pulumi.StringPtrOutput)
}

func (o DatasourcePtrOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) *string {
		if v == nil {
			return nil
		}
		return v.ResourceName
	}).(pulumi.StringPtrOutput)
}

func (o DatasourcePtrOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) *string {
		if v == nil {
			return nil
		}
		return v.ResourceType
	}).(pulumi.StringPtrOutput)
}

func (o DatasourcePtrOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) *string {
		if v == nil {
			return nil
		}
		return v.ResourceUri
	}).(pulumi.StringPtrOutput)
}

type DatasourceResponse struct {
	DatasourceType   *string `pulumi:"datasourceType"`
	ObjectType       *string `pulumi:"objectType"`
	ResourceID       string  `pulumi:"resourceID"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceName     *string `pulumi:"resourceName"`
	ResourceType     *string `pulumi:"resourceType"`
	ResourceUri      *string `pulumi:"resourceUri"`
}

type DatasourceResponseOutput struct{ *pulumi.OutputState }

func (DatasourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasourceResponse)(nil)).Elem()
}

func (o DatasourceResponseOutput) ToDatasourceResponseOutput() DatasourceResponseOutput {
	return o
}

func (o DatasourceResponseOutput) ToDatasourceResponseOutputWithContext(ctx context.Context) DatasourceResponseOutput {
	return o
}

func (o DatasourceResponseOutput) DatasourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceResponse) *string { return v.DatasourceType }).(pulumi.StringPtrOutput)
}

func (o DatasourceResponseOutput) ObjectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceResponse) *string { return v.ObjectType }).(pulumi.StringPtrOutput)
}

func (o DatasourceResponseOutput) ResourceID() pulumi.StringOutput {
	return o.ApplyT(func(v DatasourceResponse) string { return v.ResourceID }).(pulumi.StringOutput)
}

func (o DatasourceResponseOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceResponse) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o DatasourceResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o DatasourceResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o DatasourceResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

type DatasourceSet struct {
	DatasourceType   *string `pulumi:"datasourceType"`
	ObjectType       *string `pulumi:"objectType"`
	ResourceID       string  `pulumi:"resourceID"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceName     *string `pulumi:"resourceName"`
	ResourceType     *string `pulumi:"resourceType"`
	ResourceUri      *string `pulumi:"resourceUri"`
}





type DatasourceSetInput interface {
	pulumi.Input

	ToDatasourceSetOutput() DatasourceSetOutput
	ToDatasourceSetOutputWithContext(context.Context) DatasourceSetOutput
}

type DatasourceSetArgs struct {
	DatasourceType   pulumi.StringPtrInput `pulumi:"datasourceType"`
	ObjectType       pulumi.StringPtrInput `pulumi:"objectType"`
	ResourceID       pulumi.StringInput    `pulumi:"resourceID"`
	ResourceLocation pulumi.StringPtrInput `pulumi:"resourceLocation"`
	ResourceName     pulumi.StringPtrInput `pulumi:"resourceName"`
	ResourceType     pulumi.StringPtrInput `pulumi:"resourceType"`
	ResourceUri      pulumi.StringPtrInput `pulumi:"resourceUri"`
}

func (DatasourceSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasourceSet)(nil)).Elem()
}

func (i DatasourceSetArgs) ToDatasourceSetOutput() DatasourceSetOutput {
	return i.ToDatasourceSetOutputWithContext(context.Background())
}

func (i DatasourceSetArgs) ToDatasourceSetOutputWithContext(ctx context.Context) DatasourceSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceSetOutput)
}

func (i DatasourceSetArgs) ToDatasourceSetPtrOutput() DatasourceSetPtrOutput {
	return i.ToDatasourceSetPtrOutputWithContext(context.Background())
}

func (i DatasourceSetArgs) ToDatasourceSetPtrOutputWithContext(ctx context.Context) DatasourceSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceSetOutput).ToDatasourceSetPtrOutputWithContext(ctx)
}









type DatasourceSetPtrInput interface {
	pulumi.Input

	ToDatasourceSetPtrOutput() DatasourceSetPtrOutput
	ToDatasourceSetPtrOutputWithContext(context.Context) DatasourceSetPtrOutput
}

type datasourceSetPtrType DatasourceSetArgs

func DatasourceSetPtr(v *DatasourceSetArgs) DatasourceSetPtrInput {
	return (*datasourceSetPtrType)(v)
}

func (*datasourceSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasourceSet)(nil)).Elem()
}

func (i *datasourceSetPtrType) ToDatasourceSetPtrOutput() DatasourceSetPtrOutput {
	return i.ToDatasourceSetPtrOutputWithContext(context.Background())
}

func (i *datasourceSetPtrType) ToDatasourceSetPtrOutputWithContext(ctx context.Context) DatasourceSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceSetPtrOutput)
}

type DatasourceSetOutput struct{ *pulumi.OutputState }

func (DatasourceSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasourceSet)(nil)).Elem()
}

func (o DatasourceSetOutput) ToDatasourceSetOutput() DatasourceSetOutput {
	return o
}

func (o DatasourceSetOutput) ToDatasourceSetOutputWithContext(ctx context.Context) DatasourceSetOutput {
	return o
}

func (o DatasourceSetOutput) ToDatasourceSetPtrOutput() DatasourceSetPtrOutput {
	return o.ToDatasourceSetPtrOutputWithContext(context.Background())
}

func (o DatasourceSetOutput) ToDatasourceSetPtrOutputWithContext(ctx context.Context) DatasourceSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasourceSet) *DatasourceSet {
		return &v
	}).(DatasourceSetPtrOutput)
}

func (o DatasourceSetOutput) DatasourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSet) *string { return v.DatasourceType }).(pulumi.StringPtrOutput)
}

func (o DatasourceSetOutput) ObjectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSet) *string { return v.ObjectType }).(pulumi.StringPtrOutput)
}

func (o DatasourceSetOutput) ResourceID() pulumi.StringOutput {
	return o.ApplyT(func(v DatasourceSet) string { return v.ResourceID }).(pulumi.StringOutput)
}

func (o DatasourceSetOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSet) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o DatasourceSetOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSet) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o DatasourceSetOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSet) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o DatasourceSetOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSet) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

type DatasourceSetPtrOutput struct{ *pulumi.OutputState }

func (DatasourceSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasourceSet)(nil)).Elem()
}

func (o DatasourceSetPtrOutput) ToDatasourceSetPtrOutput() DatasourceSetPtrOutput {
	return o
}

func (o DatasourceSetPtrOutput) ToDatasourceSetPtrOutputWithContext(ctx context.Context) DatasourceSetPtrOutput {
	return o
}

func (o DatasourceSetPtrOutput) Elem() DatasourceSetOutput {
	return o.ApplyT(func(v *DatasourceSet) DatasourceSet {
		if v != nil {
			return *v
		}
		var ret DatasourceSet
		return ret
	}).(DatasourceSetOutput)
}

func (o DatasourceSetPtrOutput) DatasourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSet) *string {
		if v == nil {
			return nil
		}
		return v.DatasourceType
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetPtrOutput) ObjectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSet) *string {
		if v == nil {
			return nil
		}
		return v.ObjectType
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetPtrOutput) ResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSet) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceID
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetPtrOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSet) *string {
		if v == nil {
			return nil
		}
		return v.ResourceLocation
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetPtrOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSet) *string {
		if v == nil {
			return nil
		}
		return v.ResourceName
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetPtrOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSet) *string {
		if v == nil {
			return nil
		}
		return v.ResourceType
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetPtrOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSet) *string {
		if v == nil {
			return nil
		}
		return v.ResourceUri
	}).(pulumi.StringPtrOutput)
}

type DatasourceSetResponse struct {
	DatasourceType   *string `pulumi:"datasourceType"`
	ObjectType       *string `pulumi:"objectType"`
	ResourceID       string  `pulumi:"resourceID"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceName     *string `pulumi:"resourceName"`
	ResourceType     *string `pulumi:"resourceType"`
	ResourceUri      *string `pulumi:"resourceUri"`
}

type DatasourceSetResponseOutput struct{ *pulumi.OutputState }

func (DatasourceSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasourceSetResponse)(nil)).Elem()
}

func (o DatasourceSetResponseOutput) ToDatasourceSetResponseOutput() DatasourceSetResponseOutput {
	return o
}

func (o DatasourceSetResponseOutput) ToDatasourceSetResponseOutputWithContext(ctx context.Context) DatasourceSetResponseOutput {
	return o
}

func (o DatasourceSetResponseOutput) DatasourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSetResponse) *string { return v.DatasourceType }).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponseOutput) ObjectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSetResponse) *string { return v.ObjectType }).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponseOutput) ResourceID() pulumi.StringOutput {
	return o.ApplyT(func(v DatasourceSetResponse) string { return v.ResourceID }).(pulumi.StringOutput)
}

func (o DatasourceSetResponseOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSetResponse) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSetResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSetResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasourceSetResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

type DatasourceSetResponsePtrOutput struct{ *pulumi.OutputState }

func (DatasourceSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasourceSetResponse)(nil)).Elem()
}

func (o DatasourceSetResponsePtrOutput) ToDatasourceSetResponsePtrOutput() DatasourceSetResponsePtrOutput {
	return o
}

func (o DatasourceSetResponsePtrOutput) ToDatasourceSetResponsePtrOutputWithContext(ctx context.Context) DatasourceSetResponsePtrOutput {
	return o
}

func (o DatasourceSetResponsePtrOutput) Elem() DatasourceSetResponseOutput {
	return o.ApplyT(func(v *DatasourceSetResponse) DatasourceSetResponse {
		if v != nil {
			return *v
		}
		var ret DatasourceSetResponse
		return ret
	}).(DatasourceSetResponseOutput)
}

func (o DatasourceSetResponsePtrOutput) DatasourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.DatasourceType
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponsePtrOutput) ObjectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectType
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponsePtrOutput) ResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceID
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponsePtrOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceLocation
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponsePtrOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceName
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponsePtrOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceType
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceSetResponsePtrOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceUri
	}).(pulumi.StringPtrOutput)
}

type Day struct {
	Date   *int  `pulumi:"date"`
	IsLast *bool `pulumi:"isLast"`
}

type DayResponse struct {
	Date   *int  `pulumi:"date"`
	IsLast *bool `pulumi:"isLast"`
}

type DppIdentityDetails struct {
	Type *string `pulumi:"type"`
}





type DppIdentityDetailsInput interface {
	pulumi.Input

	ToDppIdentityDetailsOutput() DppIdentityDetailsOutput
	ToDppIdentityDetailsOutputWithContext(context.Context) DppIdentityDetailsOutput
}

type DppIdentityDetailsArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (DppIdentityDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DppIdentityDetails)(nil)).Elem()
}

func (i DppIdentityDetailsArgs) ToDppIdentityDetailsOutput() DppIdentityDetailsOutput {
	return i.ToDppIdentityDetailsOutputWithContext(context.Background())
}

func (i DppIdentityDetailsArgs) ToDppIdentityDetailsOutputWithContext(ctx context.Context) DppIdentityDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DppIdentityDetailsOutput)
}

func (i DppIdentityDetailsArgs) ToDppIdentityDetailsPtrOutput() DppIdentityDetailsPtrOutput {
	return i.ToDppIdentityDetailsPtrOutputWithContext(context.Background())
}

func (i DppIdentityDetailsArgs) ToDppIdentityDetailsPtrOutputWithContext(ctx context.Context) DppIdentityDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DppIdentityDetailsOutput).ToDppIdentityDetailsPtrOutputWithContext(ctx)
}









type DppIdentityDetailsPtrInput interface {
	pulumi.Input

	ToDppIdentityDetailsPtrOutput() DppIdentityDetailsPtrOutput
	ToDppIdentityDetailsPtrOutputWithContext(context.Context) DppIdentityDetailsPtrOutput
}

type dppIdentityDetailsPtrType DppIdentityDetailsArgs

func DppIdentityDetailsPtr(v *DppIdentityDetailsArgs) DppIdentityDetailsPtrInput {
	return (*dppIdentityDetailsPtrType)(v)
}

func (*dppIdentityDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DppIdentityDetails)(nil)).Elem()
}

func (i *dppIdentityDetailsPtrType) ToDppIdentityDetailsPtrOutput() DppIdentityDetailsPtrOutput {
	return i.ToDppIdentityDetailsPtrOutputWithContext(context.Background())
}

func (i *dppIdentityDetailsPtrType) ToDppIdentityDetailsPtrOutputWithContext(ctx context.Context) DppIdentityDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DppIdentityDetailsPtrOutput)
}

type DppIdentityDetailsOutput struct{ *pulumi.OutputState }

func (DppIdentityDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DppIdentityDetails)(nil)).Elem()
}

func (o DppIdentityDetailsOutput) ToDppIdentityDetailsOutput() DppIdentityDetailsOutput {
	return o
}

func (o DppIdentityDetailsOutput) ToDppIdentityDetailsOutputWithContext(ctx context.Context) DppIdentityDetailsOutput {
	return o
}

func (o DppIdentityDetailsOutput) ToDppIdentityDetailsPtrOutput() DppIdentityDetailsPtrOutput {
	return o.ToDppIdentityDetailsPtrOutputWithContext(context.Background())
}

func (o DppIdentityDetailsOutput) ToDppIdentityDetailsPtrOutputWithContext(ctx context.Context) DppIdentityDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DppIdentityDetails) *DppIdentityDetails {
		return &v
	}).(DppIdentityDetailsPtrOutput)
}

func (o DppIdentityDetailsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DppIdentityDetails) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type DppIdentityDetailsPtrOutput struct{ *pulumi.OutputState }

func (DppIdentityDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DppIdentityDetails)(nil)).Elem()
}

func (o DppIdentityDetailsPtrOutput) ToDppIdentityDetailsPtrOutput() DppIdentityDetailsPtrOutput {
	return o
}

func (o DppIdentityDetailsPtrOutput) ToDppIdentityDetailsPtrOutputWithContext(ctx context.Context) DppIdentityDetailsPtrOutput {
	return o
}

func (o DppIdentityDetailsPtrOutput) Elem() DppIdentityDetailsOutput {
	return o.ApplyT(func(v *DppIdentityDetails) DppIdentityDetails {
		if v != nil {
			return *v
		}
		var ret DppIdentityDetails
		return ret
	}).(DppIdentityDetailsOutput)
}

func (o DppIdentityDetailsPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DppIdentityDetails) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type DppIdentityDetailsResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type DppIdentityDetailsResponseOutput struct{ *pulumi.OutputState }

func (DppIdentityDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DppIdentityDetailsResponse)(nil)).Elem()
}

func (o DppIdentityDetailsResponseOutput) ToDppIdentityDetailsResponseOutput() DppIdentityDetailsResponseOutput {
	return o
}

func (o DppIdentityDetailsResponseOutput) ToDppIdentityDetailsResponseOutputWithContext(ctx context.Context) DppIdentityDetailsResponseOutput {
	return o
}

func (o DppIdentityDetailsResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v DppIdentityDetailsResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o DppIdentityDetailsResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v DppIdentityDetailsResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o DppIdentityDetailsResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DppIdentityDetailsResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type DppIdentityDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (DppIdentityDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DppIdentityDetailsResponse)(nil)).Elem()
}

func (o DppIdentityDetailsResponsePtrOutput) ToDppIdentityDetailsResponsePtrOutput() DppIdentityDetailsResponsePtrOutput {
	return o
}

func (o DppIdentityDetailsResponsePtrOutput) ToDppIdentityDetailsResponsePtrOutputWithContext(ctx context.Context) DppIdentityDetailsResponsePtrOutput {
	return o
}

func (o DppIdentityDetailsResponsePtrOutput) Elem() DppIdentityDetailsResponseOutput {
	return o.ApplyT(func(v *DppIdentityDetailsResponse) DppIdentityDetailsResponse {
		if v != nil {
			return *v
		}
		var ret DppIdentityDetailsResponse
		return ret
	}).(DppIdentityDetailsResponseOutput)
}

func (o DppIdentityDetailsResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DppIdentityDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o DppIdentityDetailsResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DppIdentityDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o DppIdentityDetailsResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DppIdentityDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ImmediateCopyOption struct {
	ObjectType string `pulumi:"objectType"`
}

type ImmediateCopyOptionResponse struct {
	ObjectType string `pulumi:"objectType"`
}

type InnerErrorResponse struct {
	AdditionalInfo     map[string]string   `pulumi:"additionalInfo"`
	Code               *string             `pulumi:"code"`
	EmbeddedInnerError *InnerErrorResponse `pulumi:"embeddedInnerError"`
}

type InnerErrorResponseOutput struct{ *pulumi.OutputState }

func (InnerErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InnerErrorResponse)(nil)).Elem()
}

func (o InnerErrorResponseOutput) ToInnerErrorResponseOutput() InnerErrorResponseOutput {
	return o
}

func (o InnerErrorResponseOutput) ToInnerErrorResponseOutputWithContext(ctx context.Context) InnerErrorResponseOutput {
	return o
}

func (o InnerErrorResponseOutput) AdditionalInfo() pulumi.StringMapOutput {
	return o.ApplyT(func(v InnerErrorResponse) map[string]string { return v.AdditionalInfo }).(pulumi.StringMapOutput)
}

func (o InnerErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o InnerErrorResponseOutput) EmbeddedInnerError() InnerErrorResponsePtrOutput {
	return o.ApplyT(func(v InnerErrorResponse) *InnerErrorResponse { return v.EmbeddedInnerError }).(InnerErrorResponsePtrOutput)
}

type InnerErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (InnerErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InnerErrorResponse)(nil)).Elem()
}

func (o InnerErrorResponsePtrOutput) ToInnerErrorResponsePtrOutput() InnerErrorResponsePtrOutput {
	return o
}

func (o InnerErrorResponsePtrOutput) ToInnerErrorResponsePtrOutputWithContext(ctx context.Context) InnerErrorResponsePtrOutput {
	return o
}

func (o InnerErrorResponsePtrOutput) Elem() InnerErrorResponseOutput {
	return o.ApplyT(func(v *InnerErrorResponse) InnerErrorResponse {
		if v != nil {
			return *v
		}
		var ret InnerErrorResponse
		return ret
	}).(InnerErrorResponseOutput)
}

func (o InnerErrorResponsePtrOutput) AdditionalInfo() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InnerErrorResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.AdditionalInfo
	}).(pulumi.StringMapOutput)
}

func (o InnerErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InnerErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o InnerErrorResponsePtrOutput) EmbeddedInnerError() InnerErrorResponsePtrOutput {
	return o.ApplyT(func(v *InnerErrorResponse) *InnerErrorResponse {
		if v == nil {
			return nil
		}
		return v.EmbeddedInnerError
	}).(InnerErrorResponsePtrOutput)
}

type PolicyInfo struct {
	PolicyId         string            `pulumi:"policyId"`
	PolicyParameters *PolicyParameters `pulumi:"policyParameters"`
}





type PolicyInfoInput interface {
	pulumi.Input

	ToPolicyInfoOutput() PolicyInfoOutput
	ToPolicyInfoOutputWithContext(context.Context) PolicyInfoOutput
}

type PolicyInfoArgs struct {
	PolicyId         pulumi.StringInput       `pulumi:"policyId"`
	PolicyParameters PolicyParametersPtrInput `pulumi:"policyParameters"`
}

func (PolicyInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyInfo)(nil)).Elem()
}

func (i PolicyInfoArgs) ToPolicyInfoOutput() PolicyInfoOutput {
	return i.ToPolicyInfoOutputWithContext(context.Background())
}

func (i PolicyInfoArgs) ToPolicyInfoOutputWithContext(ctx context.Context) PolicyInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyInfoOutput)
}

func (i PolicyInfoArgs) ToPolicyInfoPtrOutput() PolicyInfoPtrOutput {
	return i.ToPolicyInfoPtrOutputWithContext(context.Background())
}

func (i PolicyInfoArgs) ToPolicyInfoPtrOutputWithContext(ctx context.Context) PolicyInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyInfoOutput).ToPolicyInfoPtrOutputWithContext(ctx)
}









type PolicyInfoPtrInput interface {
	pulumi.Input

	ToPolicyInfoPtrOutput() PolicyInfoPtrOutput
	ToPolicyInfoPtrOutputWithContext(context.Context) PolicyInfoPtrOutput
}

type policyInfoPtrType PolicyInfoArgs

func PolicyInfoPtr(v *PolicyInfoArgs) PolicyInfoPtrInput {
	return (*policyInfoPtrType)(v)
}

func (*policyInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyInfo)(nil)).Elem()
}

func (i *policyInfoPtrType) ToPolicyInfoPtrOutput() PolicyInfoPtrOutput {
	return i.ToPolicyInfoPtrOutputWithContext(context.Background())
}

func (i *policyInfoPtrType) ToPolicyInfoPtrOutputWithContext(ctx context.Context) PolicyInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyInfoPtrOutput)
}

type PolicyInfoOutput struct{ *pulumi.OutputState }

func (PolicyInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyInfo)(nil)).Elem()
}

func (o PolicyInfoOutput) ToPolicyInfoOutput() PolicyInfoOutput {
	return o
}

func (o PolicyInfoOutput) ToPolicyInfoOutputWithContext(ctx context.Context) PolicyInfoOutput {
	return o
}

func (o PolicyInfoOutput) ToPolicyInfoPtrOutput() PolicyInfoPtrOutput {
	return o.ToPolicyInfoPtrOutputWithContext(context.Background())
}

func (o PolicyInfoOutput) ToPolicyInfoPtrOutputWithContext(ctx context.Context) PolicyInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyInfo) *PolicyInfo {
		return &v
	}).(PolicyInfoPtrOutput)
}

func (o PolicyInfoOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyInfo) string { return v.PolicyId }).(pulumi.StringOutput)
}

func (o PolicyInfoOutput) PolicyParameters() PolicyParametersPtrOutput {
	return o.ApplyT(func(v PolicyInfo) *PolicyParameters { return v.PolicyParameters }).(PolicyParametersPtrOutput)
}

type PolicyInfoPtrOutput struct{ *pulumi.OutputState }

func (PolicyInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyInfo)(nil)).Elem()
}

func (o PolicyInfoPtrOutput) ToPolicyInfoPtrOutput() PolicyInfoPtrOutput {
	return o
}

func (o PolicyInfoPtrOutput) ToPolicyInfoPtrOutputWithContext(ctx context.Context) PolicyInfoPtrOutput {
	return o
}

func (o PolicyInfoPtrOutput) Elem() PolicyInfoOutput {
	return o.ApplyT(func(v *PolicyInfo) PolicyInfo {
		if v != nil {
			return *v
		}
		var ret PolicyInfo
		return ret
	}).(PolicyInfoOutput)
}

func (o PolicyInfoPtrOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyInfo) *string {
		if v == nil {
			return nil
		}
		return &v.PolicyId
	}).(pulumi.StringPtrOutput)
}

func (o PolicyInfoPtrOutput) PolicyParameters() PolicyParametersPtrOutput {
	return o.ApplyT(func(v *PolicyInfo) *PolicyParameters {
		if v == nil {
			return nil
		}
		return v.PolicyParameters
	}).(PolicyParametersPtrOutput)
}

type PolicyInfoResponse struct {
	PolicyId         string                    `pulumi:"policyId"`
	PolicyParameters *PolicyParametersResponse `pulumi:"policyParameters"`
	PolicyVersion    string                    `pulumi:"policyVersion"`
}

type PolicyInfoResponseOutput struct{ *pulumi.OutputState }

func (PolicyInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyInfoResponse)(nil)).Elem()
}

func (o PolicyInfoResponseOutput) ToPolicyInfoResponseOutput() PolicyInfoResponseOutput {
	return o
}

func (o PolicyInfoResponseOutput) ToPolicyInfoResponseOutputWithContext(ctx context.Context) PolicyInfoResponseOutput {
	return o
}

func (o PolicyInfoResponseOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyInfoResponse) string { return v.PolicyId }).(pulumi.StringOutput)
}

func (o PolicyInfoResponseOutput) PolicyParameters() PolicyParametersResponsePtrOutput {
	return o.ApplyT(func(v PolicyInfoResponse) *PolicyParametersResponse { return v.PolicyParameters }).(PolicyParametersResponsePtrOutput)
}

func (o PolicyInfoResponseOutput) PolicyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyInfoResponse) string { return v.PolicyVersion }).(pulumi.StringOutput)
}

type PolicyParameters struct {
	DataStoreParametersList []AzureOperationalStoreParameters `pulumi:"dataStoreParametersList"`
}





type PolicyParametersInput interface {
	pulumi.Input

	ToPolicyParametersOutput() PolicyParametersOutput
	ToPolicyParametersOutputWithContext(context.Context) PolicyParametersOutput
}

type PolicyParametersArgs struct {
	DataStoreParametersList AzureOperationalStoreParametersArrayInput `pulumi:"dataStoreParametersList"`
}

func (PolicyParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyParameters)(nil)).Elem()
}

func (i PolicyParametersArgs) ToPolicyParametersOutput() PolicyParametersOutput {
	return i.ToPolicyParametersOutputWithContext(context.Background())
}

func (i PolicyParametersArgs) ToPolicyParametersOutputWithContext(ctx context.Context) PolicyParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyParametersOutput)
}

func (i PolicyParametersArgs) ToPolicyParametersPtrOutput() PolicyParametersPtrOutput {
	return i.ToPolicyParametersPtrOutputWithContext(context.Background())
}

func (i PolicyParametersArgs) ToPolicyParametersPtrOutputWithContext(ctx context.Context) PolicyParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyParametersOutput).ToPolicyParametersPtrOutputWithContext(ctx)
}









type PolicyParametersPtrInput interface {
	pulumi.Input

	ToPolicyParametersPtrOutput() PolicyParametersPtrOutput
	ToPolicyParametersPtrOutputWithContext(context.Context) PolicyParametersPtrOutput
}

type policyParametersPtrType PolicyParametersArgs

func PolicyParametersPtr(v *PolicyParametersArgs) PolicyParametersPtrInput {
	return (*policyParametersPtrType)(v)
}

func (*policyParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyParameters)(nil)).Elem()
}

func (i *policyParametersPtrType) ToPolicyParametersPtrOutput() PolicyParametersPtrOutput {
	return i.ToPolicyParametersPtrOutputWithContext(context.Background())
}

func (i *policyParametersPtrType) ToPolicyParametersPtrOutputWithContext(ctx context.Context) PolicyParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyParametersPtrOutput)
}

type PolicyParametersOutput struct{ *pulumi.OutputState }

func (PolicyParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyParameters)(nil)).Elem()
}

func (o PolicyParametersOutput) ToPolicyParametersOutput() PolicyParametersOutput {
	return o
}

func (o PolicyParametersOutput) ToPolicyParametersOutputWithContext(ctx context.Context) PolicyParametersOutput {
	return o
}

func (o PolicyParametersOutput) ToPolicyParametersPtrOutput() PolicyParametersPtrOutput {
	return o.ToPolicyParametersPtrOutputWithContext(context.Background())
}

func (o PolicyParametersOutput) ToPolicyParametersPtrOutputWithContext(ctx context.Context) PolicyParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyParameters) *PolicyParameters {
		return &v
	}).(PolicyParametersPtrOutput)
}

func (o PolicyParametersOutput) DataStoreParametersList() AzureOperationalStoreParametersArrayOutput {
	return o.ApplyT(func(v PolicyParameters) []AzureOperationalStoreParameters { return v.DataStoreParametersList }).(AzureOperationalStoreParametersArrayOutput)
}

type PolicyParametersPtrOutput struct{ *pulumi.OutputState }

func (PolicyParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyParameters)(nil)).Elem()
}

func (o PolicyParametersPtrOutput) ToPolicyParametersPtrOutput() PolicyParametersPtrOutput {
	return o
}

func (o PolicyParametersPtrOutput) ToPolicyParametersPtrOutputWithContext(ctx context.Context) PolicyParametersPtrOutput {
	return o
}

func (o PolicyParametersPtrOutput) Elem() PolicyParametersOutput {
	return o.ApplyT(func(v *PolicyParameters) PolicyParameters {
		if v != nil {
			return *v
		}
		var ret PolicyParameters
		return ret
	}).(PolicyParametersOutput)
}

func (o PolicyParametersPtrOutput) DataStoreParametersList() AzureOperationalStoreParametersArrayOutput {
	return o.ApplyT(func(v *PolicyParameters) []AzureOperationalStoreParameters {
		if v == nil {
			return nil
		}
		return v.DataStoreParametersList
	}).(AzureOperationalStoreParametersArrayOutput)
}

type PolicyParametersResponse struct {
	DataStoreParametersList []AzureOperationalStoreParametersResponse `pulumi:"dataStoreParametersList"`
}

type PolicyParametersResponseOutput struct{ *pulumi.OutputState }

func (PolicyParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyParametersResponse)(nil)).Elem()
}

func (o PolicyParametersResponseOutput) ToPolicyParametersResponseOutput() PolicyParametersResponseOutput {
	return o
}

func (o PolicyParametersResponseOutput) ToPolicyParametersResponseOutputWithContext(ctx context.Context) PolicyParametersResponseOutput {
	return o
}

func (o PolicyParametersResponseOutput) DataStoreParametersList() AzureOperationalStoreParametersResponseArrayOutput {
	return o.ApplyT(func(v PolicyParametersResponse) []AzureOperationalStoreParametersResponse {
		return v.DataStoreParametersList
	}).(AzureOperationalStoreParametersResponseArrayOutput)
}

type PolicyParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (PolicyParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyParametersResponse)(nil)).Elem()
}

func (o PolicyParametersResponsePtrOutput) ToPolicyParametersResponsePtrOutput() PolicyParametersResponsePtrOutput {
	return o
}

func (o PolicyParametersResponsePtrOutput) ToPolicyParametersResponsePtrOutputWithContext(ctx context.Context) PolicyParametersResponsePtrOutput {
	return o
}

func (o PolicyParametersResponsePtrOutput) Elem() PolicyParametersResponseOutput {
	return o.ApplyT(func(v *PolicyParametersResponse) PolicyParametersResponse {
		if v != nil {
			return *v
		}
		var ret PolicyParametersResponse
		return ret
	}).(PolicyParametersResponseOutput)
}

func (o PolicyParametersResponsePtrOutput) DataStoreParametersList() AzureOperationalStoreParametersResponseArrayOutput {
	return o.ApplyT(func(v *PolicyParametersResponse) []AzureOperationalStoreParametersResponse {
		if v == nil {
			return nil
		}
		return v.DataStoreParametersList
	}).(AzureOperationalStoreParametersResponseArrayOutput)
}

type ProtectionStatusDetailsResponse struct {
	ErrorDetails *UserFacingErrorResponse `pulumi:"errorDetails"`
	Status       *string                  `pulumi:"status"`
}

type ProtectionStatusDetailsResponseOutput struct{ *pulumi.OutputState }

func (ProtectionStatusDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionStatusDetailsResponse)(nil)).Elem()
}

func (o ProtectionStatusDetailsResponseOutput) ToProtectionStatusDetailsResponseOutput() ProtectionStatusDetailsResponseOutput {
	return o
}

func (o ProtectionStatusDetailsResponseOutput) ToProtectionStatusDetailsResponseOutputWithContext(ctx context.Context) ProtectionStatusDetailsResponseOutput {
	return o
}

func (o ProtectionStatusDetailsResponseOutput) ErrorDetails() UserFacingErrorResponsePtrOutput {
	return o.ApplyT(func(v ProtectionStatusDetailsResponse) *UserFacingErrorResponse { return v.ErrorDetails }).(UserFacingErrorResponsePtrOutput)
}

func (o ProtectionStatusDetailsResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionStatusDetailsResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RetentionTag struct {
	TagName string `pulumi:"tagName"`
}

type RetentionTagResponse struct {
	ETag    string `pulumi:"eTag"`
	Id      string `pulumi:"id"`
	TagName string `pulumi:"tagName"`
}

type ScheduleBasedBackupCriteria struct {
	AbsoluteCriteria []string `pulumi:"absoluteCriteria"`
	DaysOfMonth      []Day    `pulumi:"daysOfMonth"`
	DaysOfTheWeek    []string `pulumi:"daysOfTheWeek"`
	MonthsOfYear     []string `pulumi:"monthsOfYear"`
	ObjectType       string   `pulumi:"objectType"`
	ScheduleTimes    []string `pulumi:"scheduleTimes"`
	WeeksOfTheMonth  []string `pulumi:"weeksOfTheMonth"`
}

type ScheduleBasedBackupCriteriaResponse struct {
	AbsoluteCriteria []string      `pulumi:"absoluteCriteria"`
	DaysOfMonth      []DayResponse `pulumi:"daysOfMonth"`
	DaysOfTheWeek    []string      `pulumi:"daysOfTheWeek"`
	MonthsOfYear     []string      `pulumi:"monthsOfYear"`
	ObjectType       string        `pulumi:"objectType"`
	ScheduleTimes    []string      `pulumi:"scheduleTimes"`
	WeeksOfTheMonth  []string      `pulumi:"weeksOfTheMonth"`
}

type ScheduleBasedTriggerContext struct {
	ObjectType      string            `pulumi:"objectType"`
	Schedule        BackupSchedule    `pulumi:"schedule"`
	TaggingCriteria []TaggingCriteria `pulumi:"taggingCriteria"`
}

type ScheduleBasedTriggerContextResponse struct {
	ObjectType      string                    `pulumi:"objectType"`
	Schedule        BackupScheduleResponse    `pulumi:"schedule"`
	TaggingCriteria []TaggingCriteriaResponse `pulumi:"taggingCriteria"`
}

type SourceLifeCycle struct {
	DeleteAfter                 AbsoluteDeleteOption `pulumi:"deleteAfter"`
	SourceDataStore             DataStoreInfoBase    `pulumi:"sourceDataStore"`
	TargetDataStoreCopySettings []TargetCopySetting  `pulumi:"targetDataStoreCopySettings"`
}

type SourceLifeCycleResponse struct {
	DeleteAfter                 AbsoluteDeleteOptionResponse `pulumi:"deleteAfter"`
	SourceDataStore             DataStoreInfoBaseResponse    `pulumi:"sourceDataStore"`
	TargetDataStoreCopySettings []TargetCopySettingResponse  `pulumi:"targetDataStoreCopySettings"`
}

type StorageSetting struct {
	DatastoreType *string `pulumi:"datastoreType"`
	Type          *string `pulumi:"type"`
}





type StorageSettingInput interface {
	pulumi.Input

	ToStorageSettingOutput() StorageSettingOutput
	ToStorageSettingOutputWithContext(context.Context) StorageSettingOutput
}

type StorageSettingArgs struct {
	DatastoreType pulumi.StringPtrInput `pulumi:"datastoreType"`
	Type          pulumi.StringPtrInput `pulumi:"type"`
}

func (StorageSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSetting)(nil)).Elem()
}

func (i StorageSettingArgs) ToStorageSettingOutput() StorageSettingOutput {
	return i.ToStorageSettingOutputWithContext(context.Background())
}

func (i StorageSettingArgs) ToStorageSettingOutputWithContext(ctx context.Context) StorageSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSettingOutput)
}





type StorageSettingArrayInput interface {
	pulumi.Input

	ToStorageSettingArrayOutput() StorageSettingArrayOutput
	ToStorageSettingArrayOutputWithContext(context.Context) StorageSettingArrayOutput
}

type StorageSettingArray []StorageSettingInput

func (StorageSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageSetting)(nil)).Elem()
}

func (i StorageSettingArray) ToStorageSettingArrayOutput() StorageSettingArrayOutput {
	return i.ToStorageSettingArrayOutputWithContext(context.Background())
}

func (i StorageSettingArray) ToStorageSettingArrayOutputWithContext(ctx context.Context) StorageSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSettingArrayOutput)
}

type StorageSettingOutput struct{ *pulumi.OutputState }

func (StorageSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSetting)(nil)).Elem()
}

func (o StorageSettingOutput) ToStorageSettingOutput() StorageSettingOutput {
	return o
}

func (o StorageSettingOutput) ToStorageSettingOutputWithContext(ctx context.Context) StorageSettingOutput {
	return o
}

func (o StorageSettingOutput) DatastoreType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSetting) *string { return v.DatastoreType }).(pulumi.StringPtrOutput)
}

func (o StorageSettingOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSetting) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type StorageSettingArrayOutput struct{ *pulumi.OutputState }

func (StorageSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageSetting)(nil)).Elem()
}

func (o StorageSettingArrayOutput) ToStorageSettingArrayOutput() StorageSettingArrayOutput {
	return o
}

func (o StorageSettingArrayOutput) ToStorageSettingArrayOutputWithContext(ctx context.Context) StorageSettingArrayOutput {
	return o
}

func (o StorageSettingArrayOutput) Index(i pulumi.IntInput) StorageSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageSetting {
		return vs[0].([]StorageSetting)[vs[1].(int)]
	}).(StorageSettingOutput)
}

type StorageSettingResponse struct {
	DatastoreType *string `pulumi:"datastoreType"`
	Type          *string `pulumi:"type"`
}

type StorageSettingResponseOutput struct{ *pulumi.OutputState }

func (StorageSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSettingResponse)(nil)).Elem()
}

func (o StorageSettingResponseOutput) ToStorageSettingResponseOutput() StorageSettingResponseOutput {
	return o
}

func (o StorageSettingResponseOutput) ToStorageSettingResponseOutputWithContext(ctx context.Context) StorageSettingResponseOutput {
	return o
}

func (o StorageSettingResponseOutput) DatastoreType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSettingResponse) *string { return v.DatastoreType }).(pulumi.StringPtrOutput)
}

func (o StorageSettingResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSettingResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type StorageSettingResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageSettingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageSettingResponse)(nil)).Elem()
}

func (o StorageSettingResponseArrayOutput) ToStorageSettingResponseArrayOutput() StorageSettingResponseArrayOutput {
	return o
}

func (o StorageSettingResponseArrayOutput) ToStorageSettingResponseArrayOutputWithContext(ctx context.Context) StorageSettingResponseArrayOutput {
	return o
}

func (o StorageSettingResponseArrayOutput) Index(i pulumi.IntInput) StorageSettingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageSettingResponse {
		return vs[0].([]StorageSettingResponse)[vs[1].(int)]
	}).(StorageSettingResponseOutput)
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

type TaggingCriteria struct {
	Criteria        []ScheduleBasedBackupCriteria `pulumi:"criteria"`
	IsDefault       bool                          `pulumi:"isDefault"`
	TagInfo         RetentionTag                  `pulumi:"tagInfo"`
	TaggingPriority float64                       `pulumi:"taggingPriority"`
}

type TaggingCriteriaResponse struct {
	Criteria        []ScheduleBasedBackupCriteriaResponse `pulumi:"criteria"`
	IsDefault       bool                                  `pulumi:"isDefault"`
	TagInfo         RetentionTagResponse                  `pulumi:"tagInfo"`
	TaggingPriority float64                               `pulumi:"taggingPriority"`
}

type TargetCopySetting struct {
	CopyAfter interface{}       `pulumi:"copyAfter"`
	DataStore DataStoreInfoBase `pulumi:"dataStore"`
}

type TargetCopySettingResponse struct {
	CopyAfter interface{}               `pulumi:"copyAfter"`
	DataStore DataStoreInfoBaseResponse `pulumi:"dataStore"`
}

type UserFacingErrorResponse struct {
	Code              *string                   `pulumi:"code"`
	Details           []UserFacingErrorResponse `pulumi:"details"`
	InnerError        *InnerErrorResponse       `pulumi:"innerError"`
	IsRetryable       *bool                     `pulumi:"isRetryable"`
	IsUserError       *bool                     `pulumi:"isUserError"`
	Message           *string                   `pulumi:"message"`
	Properties        map[string]string         `pulumi:"properties"`
	RecommendedAction []string                  `pulumi:"recommendedAction"`
	Target            *string                   `pulumi:"target"`
}

type UserFacingErrorResponseOutput struct{ *pulumi.OutputState }

func (UserFacingErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserFacingErrorResponse)(nil)).Elem()
}

func (o UserFacingErrorResponseOutput) ToUserFacingErrorResponseOutput() UserFacingErrorResponseOutput {
	return o
}

func (o UserFacingErrorResponseOutput) ToUserFacingErrorResponseOutputWithContext(ctx context.Context) UserFacingErrorResponseOutput {
	return o
}

func (o UserFacingErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserFacingErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o UserFacingErrorResponseOutput) Details() UserFacingErrorResponseArrayOutput {
	return o.ApplyT(func(v UserFacingErrorResponse) []UserFacingErrorResponse { return v.Details }).(UserFacingErrorResponseArrayOutput)
}

func (o UserFacingErrorResponseOutput) InnerError() InnerErrorResponsePtrOutput {
	return o.ApplyT(func(v UserFacingErrorResponse) *InnerErrorResponse { return v.InnerError }).(InnerErrorResponsePtrOutput)
}

func (o UserFacingErrorResponseOutput) IsRetryable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UserFacingErrorResponse) *bool { return v.IsRetryable }).(pulumi.BoolPtrOutput)
}

func (o UserFacingErrorResponseOutput) IsUserError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UserFacingErrorResponse) *bool { return v.IsUserError }).(pulumi.BoolPtrOutput)
}

func (o UserFacingErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserFacingErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o UserFacingErrorResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v UserFacingErrorResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o UserFacingErrorResponseOutput) RecommendedAction() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UserFacingErrorResponse) []string { return v.RecommendedAction }).(pulumi.StringArrayOutput)
}

func (o UserFacingErrorResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserFacingErrorResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type UserFacingErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (UserFacingErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFacingErrorResponse)(nil)).Elem()
}

func (o UserFacingErrorResponsePtrOutput) ToUserFacingErrorResponsePtrOutput() UserFacingErrorResponsePtrOutput {
	return o
}

func (o UserFacingErrorResponsePtrOutput) ToUserFacingErrorResponsePtrOutputWithContext(ctx context.Context) UserFacingErrorResponsePtrOutput {
	return o
}

func (o UserFacingErrorResponsePtrOutput) Elem() UserFacingErrorResponseOutput {
	return o.ApplyT(func(v *UserFacingErrorResponse) UserFacingErrorResponse {
		if v != nil {
			return *v
		}
		var ret UserFacingErrorResponse
		return ret
	}).(UserFacingErrorResponseOutput)
}

func (o UserFacingErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFacingErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o UserFacingErrorResponsePtrOutput) Details() UserFacingErrorResponseArrayOutput {
	return o.ApplyT(func(v *UserFacingErrorResponse) []UserFacingErrorResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(UserFacingErrorResponseArrayOutput)
}

func (o UserFacingErrorResponsePtrOutput) InnerError() InnerErrorResponsePtrOutput {
	return o.ApplyT(func(v *UserFacingErrorResponse) *InnerErrorResponse {
		if v == nil {
			return nil
		}
		return v.InnerError
	}).(InnerErrorResponsePtrOutput)
}

func (o UserFacingErrorResponsePtrOutput) IsRetryable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserFacingErrorResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsRetryable
	}).(pulumi.BoolPtrOutput)
}

func (o UserFacingErrorResponsePtrOutput) IsUserError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserFacingErrorResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsUserError
	}).(pulumi.BoolPtrOutput)
}

func (o UserFacingErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFacingErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o UserFacingErrorResponsePtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *UserFacingErrorResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o UserFacingErrorResponsePtrOutput) RecommendedAction() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserFacingErrorResponse) []string {
		if v == nil {
			return nil
		}
		return v.RecommendedAction
	}).(pulumi.StringArrayOutput)
}

func (o UserFacingErrorResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFacingErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type UserFacingErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (UserFacingErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserFacingErrorResponse)(nil)).Elem()
}

func (o UserFacingErrorResponseArrayOutput) ToUserFacingErrorResponseArrayOutput() UserFacingErrorResponseArrayOutput {
	return o
}

func (o UserFacingErrorResponseArrayOutput) ToUserFacingErrorResponseArrayOutputWithContext(ctx context.Context) UserFacingErrorResponseArrayOutput {
	return o
}

func (o UserFacingErrorResponseArrayOutput) Index(i pulumi.IntInput) UserFacingErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserFacingErrorResponse {
		return vs[0].([]UserFacingErrorResponse)[vs[1].(int)]
	}).(UserFacingErrorResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureOperationalStoreParametersOutput{})
	pulumi.RegisterOutputType(AzureOperationalStoreParametersArrayOutput{})
	pulumi.RegisterOutputType(AzureOperationalStoreParametersResponseOutput{})
	pulumi.RegisterOutputType(AzureOperationalStoreParametersResponseArrayOutput{})
	pulumi.RegisterOutputType(BackupInstanceTypeOutput{})
	pulumi.RegisterOutputType(BackupInstanceTypePtrOutput{})
	pulumi.RegisterOutputType(BackupInstanceResponseOutput{})
	pulumi.RegisterOutputType(BackupPolicyTypeOutput{})
	pulumi.RegisterOutputType(BackupPolicyTypePtrOutput{})
	pulumi.RegisterOutputType(BackupPolicyResponseOutput{})
	pulumi.RegisterOutputType(BackupVaultTypeOutput{})
	pulumi.RegisterOutputType(BackupVaultResponseOutput{})
	pulumi.RegisterOutputType(DatasourceOutput{})
	pulumi.RegisterOutputType(DatasourcePtrOutput{})
	pulumi.RegisterOutputType(DatasourceResponseOutput{})
	pulumi.RegisterOutputType(DatasourceSetOutput{})
	pulumi.RegisterOutputType(DatasourceSetPtrOutput{})
	pulumi.RegisterOutputType(DatasourceSetResponseOutput{})
	pulumi.RegisterOutputType(DatasourceSetResponsePtrOutput{})
	pulumi.RegisterOutputType(DppIdentityDetailsOutput{})
	pulumi.RegisterOutputType(DppIdentityDetailsPtrOutput{})
	pulumi.RegisterOutputType(DppIdentityDetailsResponseOutput{})
	pulumi.RegisterOutputType(DppIdentityDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(InnerErrorResponseOutput{})
	pulumi.RegisterOutputType(InnerErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(PolicyInfoOutput{})
	pulumi.RegisterOutputType(PolicyInfoPtrOutput{})
	pulumi.RegisterOutputType(PolicyInfoResponseOutput{})
	pulumi.RegisterOutputType(PolicyParametersOutput{})
	pulumi.RegisterOutputType(PolicyParametersPtrOutput{})
	pulumi.RegisterOutputType(PolicyParametersResponseOutput{})
	pulumi.RegisterOutputType(PolicyParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(ProtectionStatusDetailsResponseOutput{})
	pulumi.RegisterOutputType(StorageSettingOutput{})
	pulumi.RegisterOutputType(StorageSettingArrayOutput{})
	pulumi.RegisterOutputType(StorageSettingResponseOutput{})
	pulumi.RegisterOutputType(StorageSettingResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserFacingErrorResponseOutput{})
	pulumi.RegisterOutputType(UserFacingErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(UserFacingErrorResponseArrayOutput{})
}
