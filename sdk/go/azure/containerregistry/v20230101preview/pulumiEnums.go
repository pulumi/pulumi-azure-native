


package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Action string

const (
	ActionAllow = Action("Allow")
)

type ActionsRequired string

const (
	ActionsRequiredNone     = ActionsRequired("None")
	ActionsRequiredRecreate = ActionsRequired("Recreate")
)

type AuditLogStatus string

const (
	AuditLogStatusEnabled  = AuditLogStatus("Enabled")
	AuditLogStatusDisabled = AuditLogStatus("Disabled")
)

type AzureADAuthenticationAsArmPolicyStatus string

const (
	AzureADAuthenticationAsArmPolicyStatusEnabled  = AzureADAuthenticationAsArmPolicyStatus("enabled")
	AzureADAuthenticationAsArmPolicyStatusDisabled = AzureADAuthenticationAsArmPolicyStatus("disabled")
)

type ConnectedRegistryMode string

const (
	ConnectedRegistryModeReadWrite = ConnectedRegistryMode("ReadWrite")
	ConnectedRegistryModeReadOnly  = ConnectedRegistryMode("ReadOnly")
	ConnectedRegistryModeRegistry  = ConnectedRegistryMode("Registry")
	ConnectedRegistryModeMirror    = ConnectedRegistryMode("Mirror")
)

type ConnectionStatus string

const (
	ConnectionStatusApproved     = ConnectionStatus("Approved")
	ConnectionStatusPending      = ConnectionStatus("Pending")
	ConnectionStatusRejected     = ConnectionStatus("Rejected")
	ConnectionStatusDisconnected = ConnectionStatus("Disconnected")
)

type CredentialName string

const (
	CredentialNameCredential1 = CredentialName("Credential1")
)

type DefaultAction string

const (
	DefaultActionAllow = DefaultAction("Allow")
	DefaultActionDeny  = DefaultAction("Deny")
)

type EncryptionStatus string

const (
	EncryptionStatusEnabled  = EncryptionStatus("enabled")
	EncryptionStatusDisabled = EncryptionStatus("disabled")
)

type ExportPolicyStatus string

const (
	ExportPolicyStatusEnabled  = ExportPolicyStatus("enabled")
	ExportPolicyStatusDisabled = ExportPolicyStatus("disabled")
)

type LogLevel string

const (
	LogLevelDebug       = LogLevel("Debug")
	LogLevelInformation = LogLevel("Information")
	LogLevelWarning     = LogLevel("Warning")
	LogLevelError       = LogLevel("Error")
	LogLevelNone        = LogLevel("None")
)

type NetworkRuleBypassOptions string

const (
	NetworkRuleBypassOptionsAzureServices = NetworkRuleBypassOptions("AzureServices")
	NetworkRuleBypassOptionsNone          = NetworkRuleBypassOptions("None")
)

type PipelineOptions string

const (
	PipelineOptionsOverwriteTags             = PipelineOptions("OverwriteTags")
	PipelineOptionsOverwriteBlobs            = PipelineOptions("OverwriteBlobs")
	PipelineOptionsDeleteSourceBlobOnSuccess = PipelineOptions("DeleteSourceBlobOnSuccess")
	PipelineOptionsContinueOnErrors          = PipelineOptions("ContinueOnErrors")
)

type PipelineRunSourceType string

const (
	PipelineRunSourceTypeAzureStorageBlob = PipelineRunSourceType("AzureStorageBlob")
)

type PipelineRunTargetType string

const (
	PipelineRunTargetTypeAzureStorageBlob = PipelineRunTargetType("AzureStorageBlob")
)

type PipelineSourceType string

const (
	PipelineSourceTypeAzureStorageBlobContainer = PipelineSourceType("AzureStorageBlobContainer")
)

type PolicyStatus string

const (
	PolicyStatusEnabled  = PolicyStatus("enabled")
	PolicyStatusDisabled = PolicyStatus("disabled")
)

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type SkuName string

const (
	SkuNameClassic  = SkuName("Classic")
	SkuNameBasic    = SkuName("Basic")
	SkuNameStandard = SkuName("Standard")
	SkuNamePremium  = SkuName("Premium")
)

type TokenCertificateName string

const (
	TokenCertificateNameCertificate1 = TokenCertificateName("certificate1")
	TokenCertificateNameCertificate2 = TokenCertificateName("certificate2")
)

type TokenPasswordName string

const (
	TokenPasswordNamePassword1 = TokenPasswordName("password1")
	TokenPasswordNamePassword2 = TokenPasswordName("password2")
)

type TokenStatus string

const (
	TokenStatusEnabled  = TokenStatus("enabled")
	TokenStatusDisabled = TokenStatus("disabled")
)

type TriggerStatus string

const (
	TriggerStatusEnabled  = TriggerStatus("Enabled")
	TriggerStatusDisabled = TriggerStatus("Disabled")
)

type TrustPolicyType string

const (
	TrustPolicyTypeNotary = TrustPolicyType("Notary")
)

type WebhookAction string

const (
	WebhookActionPush          = WebhookAction("push")
	WebhookActionDelete        = WebhookAction("delete")
	WebhookActionQuarantine    = WebhookAction("quarantine")
	WebhookAction_Chart_push   = WebhookAction("chart_push")
	WebhookAction_Chart_delete = WebhookAction("chart_delete")
)

type WebhookStatus string

const (
	WebhookStatusEnabled  = WebhookStatus("enabled")
	WebhookStatusDisabled = WebhookStatus("disabled")
)

type ZoneRedundancy string

const (
	ZoneRedundancyEnabled  = ZoneRedundancy("Enabled")
	ZoneRedundancyDisabled = ZoneRedundancy("Disabled")
)

func init() {
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
