


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CspmMonitorAwsOffering struct {
	NativeCloudConnection *CspmMonitorAwsOfferingNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                       `pulumi:"offeringType"`
}

type CspmMonitorAwsOfferingNativeCloudConnection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type CspmMonitorAwsOfferingResponse struct {
	Description           string                                               `pulumi:"description"`
	NativeCloudConnection *CspmMonitorAwsOfferingResponseNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                               `pulumi:"offeringType"`
}

type CspmMonitorAwsOfferingResponseNativeCloudConnection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOffering struct {
	CloudWatchToKinesis   *DefenderForContainersAwsOfferingCloudWatchToKinesis   `pulumi:"cloudWatchToKinesis"`
	KinesisToS3           *DefenderForContainersAwsOfferingKinesisToS3           `pulumi:"kinesisToS3"`
	KubernetesScubaReader *DefenderForContainersAwsOfferingKubernetesScubaReader `pulumi:"kubernetesScubaReader"`
	KubernetesService     *DefenderForContainersAwsOfferingKubernetesService     `pulumi:"kubernetesService"`
	OfferingType          string                                                 `pulumi:"offeringType"`
}

type DefenderForContainersAwsOfferingCloudWatchToKinesis struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingKinesisToS3 struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingKubernetesScubaReader struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingKubernetesService struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponse struct {
	CloudWatchToKinesis   *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis   `pulumi:"cloudWatchToKinesis"`
	Description           string                                                         `pulumi:"description"`
	KinesisToS3           *DefenderForContainersAwsOfferingResponseKinesisToS3           `pulumi:"kinesisToS3"`
	KubernetesScubaReader *DefenderForContainersAwsOfferingResponseKubernetesScubaReader `pulumi:"kubernetesScubaReader"`
	KubernetesService     *DefenderForContainersAwsOfferingResponseKubernetesService     `pulumi:"kubernetesService"`
	OfferingType          string                                                         `pulumi:"offeringType"`
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesis struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseKinesisToS3 struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReader struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseKubernetesService struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForServersAwsOffering struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersAwsOfferingDefenderForServers  `pulumi:"defenderForServers"`
	OfferingType        string                                            `pulumi:"offeringType"`
}

type DefenderForServersAwsOfferingArcAutoProvisioning struct {
	Enabled                        *bool                                                        `pulumi:"enabled"`
	ServicePrincipalSecretMetadata *DefenderForServersAwsOfferingServicePrincipalSecretMetadata `pulumi:"servicePrincipalSecretMetadata"`
}

type DefenderForServersAwsOfferingDefenderForServers struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForServersAwsOfferingResponse struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingResponseArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersAwsOfferingResponseDefenderForServers  `pulumi:"defenderForServers"`
	Description         string                                                    `pulumi:"description"`
	OfferingType        string                                                    `pulumi:"offeringType"`
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioning struct {
	Enabled                        *bool                                                                `pulumi:"enabled"`
	ServicePrincipalSecretMetadata *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata `pulumi:"servicePrincipalSecretMetadata"`
}

type DefenderForServersAwsOfferingResponseDefenderForServers struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata struct {
	ExpiryDate           *string `pulumi:"expiryDate"`
	ParameterNameInStore *string `pulumi:"parameterNameInStore"`
	ParameterStoreRegion *string `pulumi:"parameterStoreRegion"`
}

type DefenderForServersAwsOfferingServicePrincipalSecretMetadata struct {
	ExpiryDate           *string `pulumi:"expiryDate"`
	ParameterNameInStore *string `pulumi:"parameterNameInStore"`
	ParameterStoreRegion *string `pulumi:"parameterStoreRegion"`
}

type InformationProtectionAwsOffering struct {
	InformationProtection *InformationProtectionAwsOfferingInformationProtection `pulumi:"informationProtection"`
	OfferingType          string                                                 `pulumi:"offeringType"`
}

type InformationProtectionAwsOfferingInformationProtection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type InformationProtectionAwsOfferingResponse struct {
	Description           string                                                         `pulumi:"description"`
	InformationProtection *InformationProtectionAwsOfferingResponseInformationProtection `pulumi:"informationProtection"`
	OfferingType          string                                                         `pulumi:"offeringType"`
}

type InformationProtectionAwsOfferingResponseInformationProtection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type SecurityConnectorPropertiesOrganizationalData struct {
	ExcludedAccountIds         []string `pulumi:"excludedAccountIds"`
	OrganizationMembershipType *string  `pulumi:"organizationMembershipType"`
	ParentHierarchyId          *string  `pulumi:"parentHierarchyId"`
	StacksetName               *string  `pulumi:"stacksetName"`
}





type SecurityConnectorPropertiesOrganizationalDataInput interface {
	pulumi.Input

	ToSecurityConnectorPropertiesOrganizationalDataOutput() SecurityConnectorPropertiesOrganizationalDataOutput
	ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(context.Context) SecurityConnectorPropertiesOrganizationalDataOutput
}

type SecurityConnectorPropertiesOrganizationalDataArgs struct {
	ExcludedAccountIds         pulumi.StringArrayInput `pulumi:"excludedAccountIds"`
	OrganizationMembershipType pulumi.StringPtrInput   `pulumi:"organizationMembershipType"`
	ParentHierarchyId          pulumi.StringPtrInput   `pulumi:"parentHierarchyId"`
	StacksetName               pulumi.StringPtrInput   `pulumi:"stacksetName"`
}

func (SecurityConnectorPropertiesOrganizationalDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataOutput() SecurityConnectorPropertiesOrganizationalDataOutput {
	return i.ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(context.Background())
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesOrganizationalDataOutput)
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return i.ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Background())
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesOrganizationalDataOutput).ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx)
}









type SecurityConnectorPropertiesOrganizationalDataPtrInput interface {
	pulumi.Input

	ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput
	ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput
}

type securityConnectorPropertiesOrganizationalDataPtrType SecurityConnectorPropertiesOrganizationalDataArgs

func SecurityConnectorPropertiesOrganizationalDataPtr(v *SecurityConnectorPropertiesOrganizationalDataArgs) SecurityConnectorPropertiesOrganizationalDataPtrInput {
	return (*securityConnectorPropertiesOrganizationalDataPtrType)(v)
}

func (*securityConnectorPropertiesOrganizationalDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (i *securityConnectorPropertiesOrganizationalDataPtrType) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return i.ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Background())
}

func (i *securityConnectorPropertiesOrganizationalDataPtrType) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesOrganizationalDataPtrOutput)
}

type SecurityConnectorPropertiesOrganizationalDataOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesOrganizationalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataOutput() SecurityConnectorPropertiesOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o.ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Background())
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityConnectorPropertiesOrganizationalData) *SecurityConnectorPropertiesOrganizationalData {
		return &v
	}).(SecurityConnectorPropertiesOrganizationalDataPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ExcludedAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) []string { return v.ExcludedAccountIds }).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) *string { return v.OrganizationMembershipType }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) *string { return v.ParentHierarchyId }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) *string { return v.StacksetName }).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesOrganizationalDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesOrganizationalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) Elem() SecurityConnectorPropertiesOrganizationalDataOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) SecurityConnectorPropertiesOrganizationalData {
		if v != nil {
			return *v
		}
		var ret SecurityConnectorPropertiesOrganizationalData
		return ret
	}).(SecurityConnectorPropertiesOrganizationalDataOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ExcludedAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedAccountIds
	}).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationMembershipType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.ParentHierarchyId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.StacksetName
	}).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesResponseOrganizationalData struct {
	ExcludedAccountIds         []string `pulumi:"excludedAccountIds"`
	OrganizationMembershipType *string  `pulumi:"organizationMembershipType"`
	ParentHierarchyId          *string  `pulumi:"parentHierarchyId"`
	StacksetName               *string  `pulumi:"stacksetName"`
}

type SecurityConnectorPropertiesResponseOrganizationalDataOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesResponseOrganizationalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesResponseOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataOutput() SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ExcludedAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) []string { return v.ExcludedAccountIds }).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) *string {
		return v.OrganizationMembershipType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) *string { return v.ParentHierarchyId }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) *string { return v.StacksetName }).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesResponseOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) Elem() SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) SecurityConnectorPropertiesResponseOrganizationalData {
		if v != nil {
			return *v
		}
		var ret SecurityConnectorPropertiesResponseOrganizationalData
		return ret
	}).(SecurityConnectorPropertiesResponseOrganizationalDataOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ExcludedAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedAccountIds
	}).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationMembershipType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.ParentHierarchyId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.StacksetName
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

func init() {
	pulumi.RegisterOutputType(SecurityConnectorPropertiesOrganizationalDataOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesOrganizationalDataPtrOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesResponseOrganizationalDataOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
