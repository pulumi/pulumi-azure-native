


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessControl struct {
	DefaultAction *string  `pulumi:"defaultAction"`
	IpAllowList   []string `pulumi:"ipAllowList"`
}





type AccessControlInput interface {
	pulumi.Input

	ToAccessControlOutput() AccessControlOutput
	ToAccessControlOutputWithContext(context.Context) AccessControlOutput
}

type AccessControlArgs struct {
	DefaultAction pulumi.StringPtrInput   `pulumi:"defaultAction"`
	IpAllowList   pulumi.StringArrayInput `pulumi:"ipAllowList"`
}

func (AccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessControl)(nil)).Elem()
}

func (i AccessControlArgs) ToAccessControlOutput() AccessControlOutput {
	return i.ToAccessControlOutputWithContext(context.Background())
}

func (i AccessControlArgs) ToAccessControlOutputWithContext(ctx context.Context) AccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessControlOutput)
}

func (i AccessControlArgs) ToAccessControlPtrOutput() AccessControlPtrOutput {
	return i.ToAccessControlPtrOutputWithContext(context.Background())
}

func (i AccessControlArgs) ToAccessControlPtrOutputWithContext(ctx context.Context) AccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessControlOutput).ToAccessControlPtrOutputWithContext(ctx)
}









type AccessControlPtrInput interface {
	pulumi.Input

	ToAccessControlPtrOutput() AccessControlPtrOutput
	ToAccessControlPtrOutputWithContext(context.Context) AccessControlPtrOutput
}

type accessControlPtrType AccessControlArgs

func AccessControlPtr(v *AccessControlArgs) AccessControlPtrInput {
	return (*accessControlPtrType)(v)
}

func (*accessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessControl)(nil)).Elem()
}

func (i *accessControlPtrType) ToAccessControlPtrOutput() AccessControlPtrOutput {
	return i.ToAccessControlPtrOutputWithContext(context.Background())
}

func (i *accessControlPtrType) ToAccessControlPtrOutputWithContext(ctx context.Context) AccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessControlPtrOutput)
}

type AccessControlOutput struct{ *pulumi.OutputState }

func (AccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessControl)(nil)).Elem()
}

func (o AccessControlOutput) ToAccessControlOutput() AccessControlOutput {
	return o
}

func (o AccessControlOutput) ToAccessControlOutputWithContext(ctx context.Context) AccessControlOutput {
	return o
}

func (o AccessControlOutput) ToAccessControlPtrOutput() AccessControlPtrOutput {
	return o.ToAccessControlPtrOutputWithContext(context.Background())
}

func (o AccessControlOutput) ToAccessControlPtrOutputWithContext(ctx context.Context) AccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessControl) *AccessControl {
		return &v
	}).(AccessControlPtrOutput)
}

func (o AccessControlOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessControl) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o AccessControlOutput) IpAllowList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AccessControl) []string { return v.IpAllowList }).(pulumi.StringArrayOutput)
}

type AccessControlPtrOutput struct{ *pulumi.OutputState }

func (AccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessControl)(nil)).Elem()
}

func (o AccessControlPtrOutput) ToAccessControlPtrOutput() AccessControlPtrOutput {
	return o
}

func (o AccessControlPtrOutput) ToAccessControlPtrOutputWithContext(ctx context.Context) AccessControlPtrOutput {
	return o
}

func (o AccessControlPtrOutput) Elem() AccessControlOutput {
	return o.ApplyT(func(v *AccessControl) AccessControl {
		if v != nil {
			return *v
		}
		var ret AccessControl
		return ret
	}).(AccessControlOutput)
}

func (o AccessControlPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessControl) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o AccessControlPtrOutput) IpAllowList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessControl) []string {
		if v == nil {
			return nil
		}
		return v.IpAllowList
	}).(pulumi.StringArrayOutput)
}

type AccessControlResponse struct {
	DefaultAction *string  `pulumi:"defaultAction"`
	IpAllowList   []string `pulumi:"ipAllowList"`
}

type AccessControlResponseOutput struct{ *pulumi.OutputState }

func (AccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessControlResponse)(nil)).Elem()
}

func (o AccessControlResponseOutput) ToAccessControlResponseOutput() AccessControlResponseOutput {
	return o
}

func (o AccessControlResponseOutput) ToAccessControlResponseOutputWithContext(ctx context.Context) AccessControlResponseOutput {
	return o
}

func (o AccessControlResponseOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessControlResponse) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o AccessControlResponseOutput) IpAllowList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AccessControlResponse) []string { return v.IpAllowList }).(pulumi.StringArrayOutput)
}

type AccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (AccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessControlResponse)(nil)).Elem()
}

func (o AccessControlResponsePtrOutput) ToAccessControlResponsePtrOutput() AccessControlResponsePtrOutput {
	return o
}

func (o AccessControlResponsePtrOutput) ToAccessControlResponsePtrOutputWithContext(ctx context.Context) AccessControlResponsePtrOutput {
	return o
}

func (o AccessControlResponsePtrOutput) Elem() AccessControlResponseOutput {
	return o.ApplyT(func(v *AccessControlResponse) AccessControlResponse {
		if v != nil {
			return *v
		}
		var ret AccessControlResponse
		return ret
	}).(AccessControlResponseOutput)
}

func (o AccessControlResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessControlResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o AccessControlResponsePtrOutput) IpAllowList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessControlResponse) []string {
		if v == nil {
			return nil
		}
		return v.IpAllowList
	}).(pulumi.StringArrayOutput)
}

type AccountEncryption struct {
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	Type               string              `pulumi:"type"`
}





type AccountEncryptionInput interface {
	pulumi.Input

	ToAccountEncryptionOutput() AccountEncryptionOutput
	ToAccountEncryptionOutputWithContext(context.Context) AccountEncryptionOutput
}

type AccountEncryptionArgs struct {
	KeyVaultProperties KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
	Type               pulumi.StringInput         `pulumi:"type"`
}

func (AccountEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryption)(nil)).Elem()
}

func (i AccountEncryptionArgs) ToAccountEncryptionOutput() AccountEncryptionOutput {
	return i.ToAccountEncryptionOutputWithContext(context.Background())
}

func (i AccountEncryptionArgs) ToAccountEncryptionOutputWithContext(ctx context.Context) AccountEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionOutput)
}

func (i AccountEncryptionArgs) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return i.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (i AccountEncryptionArgs) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionOutput).ToAccountEncryptionPtrOutputWithContext(ctx)
}









type AccountEncryptionPtrInput interface {
	pulumi.Input

	ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput
	ToAccountEncryptionPtrOutputWithContext(context.Context) AccountEncryptionPtrOutput
}

type accountEncryptionPtrType AccountEncryptionArgs

func AccountEncryptionPtr(v *AccountEncryptionArgs) AccountEncryptionPtrInput {
	return (*accountEncryptionPtrType)(v)
}

func (*accountEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryption)(nil)).Elem()
}

func (i *accountEncryptionPtrType) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return i.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (i *accountEncryptionPtrType) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionPtrOutput)
}

type AccountEncryptionOutput struct{ *pulumi.OutputState }

func (AccountEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryption)(nil)).Elem()
}

func (o AccountEncryptionOutput) ToAccountEncryptionOutput() AccountEncryptionOutput {
	return o
}

func (o AccountEncryptionOutput) ToAccountEncryptionOutputWithContext(ctx context.Context) AccountEncryptionOutput {
	return o
}

func (o AccountEncryptionOutput) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return o.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (o AccountEncryptionOutput) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountEncryption) *AccountEncryption {
		return &v
	}).(AccountEncryptionPtrOutput)
}

func (o AccountEncryptionOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v AccountEncryption) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

func (o AccountEncryptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccountEncryption) string { return v.Type }).(pulumi.StringOutput)
}

type AccountEncryptionPtrOutput struct{ *pulumi.OutputState }

func (AccountEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryption)(nil)).Elem()
}

func (o AccountEncryptionPtrOutput) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return o
}

func (o AccountEncryptionPtrOutput) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return o
}

func (o AccountEncryptionPtrOutput) Elem() AccountEncryptionOutput {
	return o.ApplyT(func(v *AccountEncryption) AccountEncryption {
		if v != nil {
			return *v
		}
		var ret AccountEncryption
		return ret
	}).(AccountEncryptionOutput)
}

func (o AccountEncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *AccountEncryption) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

func (o AccountEncryptionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryption) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type AccountEncryptionResponse struct {
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Type               string                      `pulumi:"type"`
}

type AccountEncryptionResponseOutput struct{ *pulumi.OutputState }

func (AccountEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryptionResponse)(nil)).Elem()
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponseOutput() AccountEncryptionResponseOutput {
	return o
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponseOutputWithContext(ctx context.Context) AccountEncryptionResponseOutput {
	return o
}

func (o AccountEncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o AccountEncryptionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AccountEncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (AccountEncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryptionResponse)(nil)).Elem()
}

func (o AccountEncryptionResponsePtrOutput) ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput {
	return o
}

func (o AccountEncryptionResponsePtrOutput) ToAccountEncryptionResponsePtrOutputWithContext(ctx context.Context) AccountEncryptionResponsePtrOutput {
	return o
}

func (o AccountEncryptionResponsePtrOutput) Elem() AccountEncryptionResponseOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) AccountEncryptionResponse {
		if v != nil {
			return *v
		}
		var ret AccountEncryptionResponse
		return ret
	}).(AccountEncryptionResponseOutput)
}

func (o AccountEncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o AccountEncryptionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type EdgeUsageDataCollectionPolicyResponse struct {
	DataCollectionFrequency           *string                        `pulumi:"dataCollectionFrequency"`
	DataReportingFrequency            *string                        `pulumi:"dataReportingFrequency"`
	EventHubDetails                   *EdgeUsageDataEventHubResponse `pulumi:"eventHubDetails"`
	MaxAllowedUnreportedUsageDuration *string                        `pulumi:"maxAllowedUnreportedUsageDuration"`
}

type EdgeUsageDataCollectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (EdgeUsageDataCollectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeUsageDataCollectionPolicyResponse)(nil)).Elem()
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) ToEdgeUsageDataCollectionPolicyResponseOutput() EdgeUsageDataCollectionPolicyResponseOutput {
	return o
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) ToEdgeUsageDataCollectionPolicyResponseOutputWithContext(ctx context.Context) EdgeUsageDataCollectionPolicyResponseOutput {
	return o
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) DataCollectionFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataCollectionPolicyResponse) *string { return v.DataCollectionFrequency }).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) DataReportingFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataCollectionPolicyResponse) *string { return v.DataReportingFrequency }).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) EventHubDetails() EdgeUsageDataEventHubResponsePtrOutput {
	return o.ApplyT(func(v EdgeUsageDataCollectionPolicyResponse) *EdgeUsageDataEventHubResponse { return v.EventHubDetails }).(EdgeUsageDataEventHubResponsePtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) MaxAllowedUnreportedUsageDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataCollectionPolicyResponse) *string { return v.MaxAllowedUnreportedUsageDuration }).(pulumi.StringPtrOutput)
}

type EdgeUsageDataCollectionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (EdgeUsageDataCollectionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeUsageDataCollectionPolicyResponse)(nil)).Elem()
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) ToEdgeUsageDataCollectionPolicyResponsePtrOutput() EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return o
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) ToEdgeUsageDataCollectionPolicyResponsePtrOutputWithContext(ctx context.Context) EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return o
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) Elem() EdgeUsageDataCollectionPolicyResponseOutput {
	return o.ApplyT(func(v *EdgeUsageDataCollectionPolicyResponse) EdgeUsageDataCollectionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret EdgeUsageDataCollectionPolicyResponse
		return ret
	}).(EdgeUsageDataCollectionPolicyResponseOutput)
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) DataCollectionFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataCollectionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.DataCollectionFrequency
	}).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) DataReportingFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataCollectionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.DataReportingFrequency
	}).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) EventHubDetails() EdgeUsageDataEventHubResponsePtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataCollectionPolicyResponse) *EdgeUsageDataEventHubResponse {
		if v == nil {
			return nil
		}
		return v.EventHubDetails
	}).(EdgeUsageDataEventHubResponsePtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) MaxAllowedUnreportedUsageDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataCollectionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaxAllowedUnreportedUsageDuration
	}).(pulumi.StringPtrOutput)
}

type EdgeUsageDataEventHubResponse struct {
	Name      *string `pulumi:"name"`
	Namespace *string `pulumi:"namespace"`
	Token     *string `pulumi:"token"`
}

type EdgeUsageDataEventHubResponseOutput struct{ *pulumi.OutputState }

func (EdgeUsageDataEventHubResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeUsageDataEventHubResponse)(nil)).Elem()
}

func (o EdgeUsageDataEventHubResponseOutput) ToEdgeUsageDataEventHubResponseOutput() EdgeUsageDataEventHubResponseOutput {
	return o
}

func (o EdgeUsageDataEventHubResponseOutput) ToEdgeUsageDataEventHubResponseOutputWithContext(ctx context.Context) EdgeUsageDataEventHubResponseOutput {
	return o
}

func (o EdgeUsageDataEventHubResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataEventHubResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataEventHubResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataEventHubResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataEventHubResponseOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataEventHubResponse) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type EdgeUsageDataEventHubResponsePtrOutput struct{ *pulumi.OutputState }

func (EdgeUsageDataEventHubResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeUsageDataEventHubResponse)(nil)).Elem()
}

func (o EdgeUsageDataEventHubResponsePtrOutput) ToEdgeUsageDataEventHubResponsePtrOutput() EdgeUsageDataEventHubResponsePtrOutput {
	return o
}

func (o EdgeUsageDataEventHubResponsePtrOutput) ToEdgeUsageDataEventHubResponsePtrOutputWithContext(ctx context.Context) EdgeUsageDataEventHubResponsePtrOutput {
	return o
}

func (o EdgeUsageDataEventHubResponsePtrOutput) Elem() EdgeUsageDataEventHubResponseOutput {
	return o.ApplyT(func(v *EdgeUsageDataEventHubResponse) EdgeUsageDataEventHubResponse {
		if v != nil {
			return *v
		}
		var ret EdgeUsageDataEventHubResponse
		return ret
	}).(EdgeUsageDataEventHubResponseOutput)
}

func (o EdgeUsageDataEventHubResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataEventHubResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataEventHubResponsePtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataEventHubResponse) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataEventHubResponsePtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataEventHubResponse) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

type KeyDelivery struct {
	AccessControl *AccessControl `pulumi:"accessControl"`
}





type KeyDeliveryInput interface {
	pulumi.Input

	ToKeyDeliveryOutput() KeyDeliveryOutput
	ToKeyDeliveryOutputWithContext(context.Context) KeyDeliveryOutput
}

type KeyDeliveryArgs struct {
	AccessControl AccessControlPtrInput `pulumi:"accessControl"`
}

func (KeyDeliveryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyDelivery)(nil)).Elem()
}

func (i KeyDeliveryArgs) ToKeyDeliveryOutput() KeyDeliveryOutput {
	return i.ToKeyDeliveryOutputWithContext(context.Background())
}

func (i KeyDeliveryArgs) ToKeyDeliveryOutputWithContext(ctx context.Context) KeyDeliveryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyDeliveryOutput)
}

func (i KeyDeliveryArgs) ToKeyDeliveryPtrOutput() KeyDeliveryPtrOutput {
	return i.ToKeyDeliveryPtrOutputWithContext(context.Background())
}

func (i KeyDeliveryArgs) ToKeyDeliveryPtrOutputWithContext(ctx context.Context) KeyDeliveryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyDeliveryOutput).ToKeyDeliveryPtrOutputWithContext(ctx)
}









type KeyDeliveryPtrInput interface {
	pulumi.Input

	ToKeyDeliveryPtrOutput() KeyDeliveryPtrOutput
	ToKeyDeliveryPtrOutputWithContext(context.Context) KeyDeliveryPtrOutput
}

type keyDeliveryPtrType KeyDeliveryArgs

func KeyDeliveryPtr(v *KeyDeliveryArgs) KeyDeliveryPtrInput {
	return (*keyDeliveryPtrType)(v)
}

func (*keyDeliveryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyDelivery)(nil)).Elem()
}

func (i *keyDeliveryPtrType) ToKeyDeliveryPtrOutput() KeyDeliveryPtrOutput {
	return i.ToKeyDeliveryPtrOutputWithContext(context.Background())
}

func (i *keyDeliveryPtrType) ToKeyDeliveryPtrOutputWithContext(ctx context.Context) KeyDeliveryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyDeliveryPtrOutput)
}

type KeyDeliveryOutput struct{ *pulumi.OutputState }

func (KeyDeliveryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyDelivery)(nil)).Elem()
}

func (o KeyDeliveryOutput) ToKeyDeliveryOutput() KeyDeliveryOutput {
	return o
}

func (o KeyDeliveryOutput) ToKeyDeliveryOutputWithContext(ctx context.Context) KeyDeliveryOutput {
	return o
}

func (o KeyDeliveryOutput) ToKeyDeliveryPtrOutput() KeyDeliveryPtrOutput {
	return o.ToKeyDeliveryPtrOutputWithContext(context.Background())
}

func (o KeyDeliveryOutput) ToKeyDeliveryPtrOutputWithContext(ctx context.Context) KeyDeliveryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyDelivery) *KeyDelivery {
		return &v
	}).(KeyDeliveryPtrOutput)
}

func (o KeyDeliveryOutput) AccessControl() AccessControlPtrOutput {
	return o.ApplyT(func(v KeyDelivery) *AccessControl { return v.AccessControl }).(AccessControlPtrOutput)
}

type KeyDeliveryPtrOutput struct{ *pulumi.OutputState }

func (KeyDeliveryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyDelivery)(nil)).Elem()
}

func (o KeyDeliveryPtrOutput) ToKeyDeliveryPtrOutput() KeyDeliveryPtrOutput {
	return o
}

func (o KeyDeliveryPtrOutput) ToKeyDeliveryPtrOutputWithContext(ctx context.Context) KeyDeliveryPtrOutput {
	return o
}

func (o KeyDeliveryPtrOutput) Elem() KeyDeliveryOutput {
	return o.ApplyT(func(v *KeyDelivery) KeyDelivery {
		if v != nil {
			return *v
		}
		var ret KeyDelivery
		return ret
	}).(KeyDeliveryOutput)
}

func (o KeyDeliveryPtrOutput) AccessControl() AccessControlPtrOutput {
	return o.ApplyT(func(v *KeyDelivery) *AccessControl {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(AccessControlPtrOutput)
}

type KeyDeliveryResponse struct {
	AccessControl *AccessControlResponse `pulumi:"accessControl"`
}

type KeyDeliveryResponseOutput struct{ *pulumi.OutputState }

func (KeyDeliveryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyDeliveryResponse)(nil)).Elem()
}

func (o KeyDeliveryResponseOutput) ToKeyDeliveryResponseOutput() KeyDeliveryResponseOutput {
	return o
}

func (o KeyDeliveryResponseOutput) ToKeyDeliveryResponseOutputWithContext(ctx context.Context) KeyDeliveryResponseOutput {
	return o
}

func (o KeyDeliveryResponseOutput) AccessControl() AccessControlResponsePtrOutput {
	return o.ApplyT(func(v KeyDeliveryResponse) *AccessControlResponse { return v.AccessControl }).(AccessControlResponsePtrOutput)
}

type KeyDeliveryResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyDeliveryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyDeliveryResponse)(nil)).Elem()
}

func (o KeyDeliveryResponsePtrOutput) ToKeyDeliveryResponsePtrOutput() KeyDeliveryResponsePtrOutput {
	return o
}

func (o KeyDeliveryResponsePtrOutput) ToKeyDeliveryResponsePtrOutputWithContext(ctx context.Context) KeyDeliveryResponsePtrOutput {
	return o
}

func (o KeyDeliveryResponsePtrOutput) Elem() KeyDeliveryResponseOutput {
	return o.ApplyT(func(v *KeyDeliveryResponse) KeyDeliveryResponse {
		if v != nil {
			return *v
		}
		var ret KeyDeliveryResponse
		return ret
	}).(KeyDeliveryResponseOutput)
}

func (o KeyDeliveryResponsePtrOutput) AccessControl() AccessControlResponsePtrOutput {
	return o.ApplyT(func(v *KeyDeliveryResponse) *AccessControlResponse {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(AccessControlResponsePtrOutput)
}

type KeyVaultProperties struct {
	KeyIdentifier *string `pulumi:"keyIdentifier"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyIdentifier pulumi.StringPtrInput `pulumi:"keyIdentifier"`
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

func (o KeyVaultPropertiesOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
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

func (o KeyVaultPropertiesPtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	CurrentKeyIdentifier string  `pulumi:"currentKeyIdentifier"`
	KeyIdentifier        *string `pulumi:"keyIdentifier"`
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

func (o KeyVaultPropertiesResponseOutput) CurrentKeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.CurrentKeyIdentifier }).(pulumi.StringOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
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

func (o KeyVaultPropertiesResponsePtrOutput) CurrentKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CurrentKeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

type MediaServiceIdentity struct {
	Type string `pulumi:"type"`
}





type MediaServiceIdentityInput interface {
	pulumi.Input

	ToMediaServiceIdentityOutput() MediaServiceIdentityOutput
	ToMediaServiceIdentityOutputWithContext(context.Context) MediaServiceIdentityOutput
}

type MediaServiceIdentityArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (MediaServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServiceIdentity)(nil)).Elem()
}

func (i MediaServiceIdentityArgs) ToMediaServiceIdentityOutput() MediaServiceIdentityOutput {
	return i.ToMediaServiceIdentityOutputWithContext(context.Background())
}

func (i MediaServiceIdentityArgs) ToMediaServiceIdentityOutputWithContext(ctx context.Context) MediaServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceIdentityOutput)
}

func (i MediaServiceIdentityArgs) ToMediaServiceIdentityPtrOutput() MediaServiceIdentityPtrOutput {
	return i.ToMediaServiceIdentityPtrOutputWithContext(context.Background())
}

func (i MediaServiceIdentityArgs) ToMediaServiceIdentityPtrOutputWithContext(ctx context.Context) MediaServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceIdentityOutput).ToMediaServiceIdentityPtrOutputWithContext(ctx)
}









type MediaServiceIdentityPtrInput interface {
	pulumi.Input

	ToMediaServiceIdentityPtrOutput() MediaServiceIdentityPtrOutput
	ToMediaServiceIdentityPtrOutputWithContext(context.Context) MediaServiceIdentityPtrOutput
}

type mediaServiceIdentityPtrType MediaServiceIdentityArgs

func MediaServiceIdentityPtr(v *MediaServiceIdentityArgs) MediaServiceIdentityPtrInput {
	return (*mediaServiceIdentityPtrType)(v)
}

func (*mediaServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServiceIdentity)(nil)).Elem()
}

func (i *mediaServiceIdentityPtrType) ToMediaServiceIdentityPtrOutput() MediaServiceIdentityPtrOutput {
	return i.ToMediaServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *mediaServiceIdentityPtrType) ToMediaServiceIdentityPtrOutputWithContext(ctx context.Context) MediaServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceIdentityPtrOutput)
}

type MediaServiceIdentityOutput struct{ *pulumi.OutputState }

func (MediaServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServiceIdentity)(nil)).Elem()
}

func (o MediaServiceIdentityOutput) ToMediaServiceIdentityOutput() MediaServiceIdentityOutput {
	return o
}

func (o MediaServiceIdentityOutput) ToMediaServiceIdentityOutputWithContext(ctx context.Context) MediaServiceIdentityOutput {
	return o
}

func (o MediaServiceIdentityOutput) ToMediaServiceIdentityPtrOutput() MediaServiceIdentityPtrOutput {
	return o.ToMediaServiceIdentityPtrOutputWithContext(context.Background())
}

func (o MediaServiceIdentityOutput) ToMediaServiceIdentityPtrOutputWithContext(ctx context.Context) MediaServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaServiceIdentity) *MediaServiceIdentity {
		return &v
	}).(MediaServiceIdentityPtrOutput)
}

func (o MediaServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MediaServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type MediaServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (MediaServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServiceIdentity)(nil)).Elem()
}

func (o MediaServiceIdentityPtrOutput) ToMediaServiceIdentityPtrOutput() MediaServiceIdentityPtrOutput {
	return o
}

func (o MediaServiceIdentityPtrOutput) ToMediaServiceIdentityPtrOutputWithContext(ctx context.Context) MediaServiceIdentityPtrOutput {
	return o
}

func (o MediaServiceIdentityPtrOutput) Elem() MediaServiceIdentityOutput {
	return o.ApplyT(func(v *MediaServiceIdentity) MediaServiceIdentity {
		if v != nil {
			return *v
		}
		var ret MediaServiceIdentity
		return ret
	}).(MediaServiceIdentityOutput)
}

func (o MediaServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type MediaServiceIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

type MediaServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (MediaServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServiceIdentityResponse)(nil)).Elem()
}

func (o MediaServiceIdentityResponseOutput) ToMediaServiceIdentityResponseOutput() MediaServiceIdentityResponseOutput {
	return o
}

func (o MediaServiceIdentityResponseOutput) ToMediaServiceIdentityResponseOutputWithContext(ctx context.Context) MediaServiceIdentityResponseOutput {
	return o
}

func (o MediaServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v MediaServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o MediaServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v MediaServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o MediaServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MediaServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type MediaServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (MediaServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServiceIdentityResponse)(nil)).Elem()
}

func (o MediaServiceIdentityResponsePtrOutput) ToMediaServiceIdentityResponsePtrOutput() MediaServiceIdentityResponsePtrOutput {
	return o
}

func (o MediaServiceIdentityResponsePtrOutput) ToMediaServiceIdentityResponsePtrOutputWithContext(ctx context.Context) MediaServiceIdentityResponsePtrOutput {
	return o
}

func (o MediaServiceIdentityResponsePtrOutput) Elem() MediaServiceIdentityResponseOutput {
	return o.ApplyT(func(v *MediaServiceIdentityResponse) MediaServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret MediaServiceIdentityResponse
		return ret
	}).(MediaServiceIdentityResponseOutput)
}

func (o MediaServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o MediaServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o MediaServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
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

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type StorageAccount struct {
	Id   *string `pulumi:"id"`
	Type string  `pulumi:"type"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Type pulumi.StringInput    `pulumi:"type"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}





type StorageAccountArrayInput interface {
	pulumi.Input

	ToStorageAccountArrayOutput() StorageAccountArrayOutput
	ToStorageAccountArrayOutputWithContext(context.Context) StorageAccountArrayOutput
}

type StorageAccountArray []StorageAccountInput

func (StorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (i StorageAccountArray) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return i.ToStorageAccountArrayOutputWithContext(context.Background())
}

func (i StorageAccountArray) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountArrayOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Type }).(pulumi.StringOutput)
}

type StorageAccountArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) Index(i pulumi.IntInput) StorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccount {
		return vs[0].([]StorageAccount)[vs[1].(int)]
	}).(StorageAccountOutput)
}

type StorageAccountResponse struct {
	Id   *string `pulumi:"id"`
	Type string  `pulumi:"type"`
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Type }).(pulumi.StringOutput)
}

type StorageAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountResponse {
		return vs[0].([]StorageAccountResponse)[vs[1].(int)]
	}).(StorageAccountResponseOutput)
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
	pulumi.RegisterOutputType(AccessControlOutput{})
	pulumi.RegisterOutputType(AccessControlPtrOutput{})
	pulumi.RegisterOutputType(AccessControlResponseOutput{})
	pulumi.RegisterOutputType(AccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(AccountEncryptionOutput{})
	pulumi.RegisterOutputType(AccountEncryptionPtrOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponseOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EdgeUsageDataCollectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(EdgeUsageDataCollectionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(EdgeUsageDataEventHubResponseOutput{})
	pulumi.RegisterOutputType(EdgeUsageDataEventHubResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyDeliveryOutput{})
	pulumi.RegisterOutputType(KeyDeliveryPtrOutput{})
	pulumi.RegisterOutputType(KeyDeliveryResponseOutput{})
	pulumi.RegisterOutputType(KeyDeliveryResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MediaServiceIdentityOutput{})
	pulumi.RegisterOutputType(MediaServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(MediaServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(MediaServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
