


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessTier string

const (
	AccessTierHot  = AccessTier("Hot")
	AccessTierCool = AccessTier("Cool")
)

func (AccessTier) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessTier)(nil)).Elem()
}

func (e AccessTier) ToAccessTierOutput() AccessTierOutput {
	return pulumi.ToOutput(e).(AccessTierOutput)
}

func (e AccessTier) ToAccessTierOutputWithContext(ctx context.Context) AccessTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessTierOutput)
}

func (e AccessTier) ToAccessTierPtrOutput() AccessTierPtrOutput {
	return e.ToAccessTierPtrOutputWithContext(context.Background())
}

func (e AccessTier) ToAccessTierPtrOutputWithContext(ctx context.Context) AccessTierPtrOutput {
	return AccessTier(e).ToAccessTierOutputWithContext(ctx).ToAccessTierPtrOutputWithContext(ctx)
}

func (e AccessTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessTierOutput struct{ *pulumi.OutputState }

func (AccessTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessTier)(nil)).Elem()
}

func (o AccessTierOutput) ToAccessTierOutput() AccessTierOutput {
	return o
}

func (o AccessTierOutput) ToAccessTierOutputWithContext(ctx context.Context) AccessTierOutput {
	return o
}

func (o AccessTierOutput) ToAccessTierPtrOutput() AccessTierPtrOutput {
	return o.ToAccessTierPtrOutputWithContext(context.Background())
}

func (o AccessTierOutput) ToAccessTierPtrOutputWithContext(ctx context.Context) AccessTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessTier) *AccessTier {
		return &v
	}).(AccessTierPtrOutput)
}

func (o AccessTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessTierPtrOutput struct{ *pulumi.OutputState }

func (AccessTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessTier)(nil)).Elem()
}

func (o AccessTierPtrOutput) ToAccessTierPtrOutput() AccessTierPtrOutput {
	return o
}

func (o AccessTierPtrOutput) ToAccessTierPtrOutputWithContext(ctx context.Context) AccessTierPtrOutput {
	return o
}

func (o AccessTierPtrOutput) Elem() AccessTierOutput {
	return o.ApplyT(func(v *AccessTier) AccessTier {
		if v != nil {
			return *v
		}
		var ret AccessTier
		return ret
	}).(AccessTierOutput)
}

func (o AccessTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessTierInput interface {
	pulumi.Input

	ToAccessTierOutput() AccessTierOutput
	ToAccessTierOutputWithContext(context.Context) AccessTierOutput
}

var accessTierPtrType = reflect.TypeOf((**AccessTier)(nil)).Elem()

type AccessTierPtrInput interface {
	pulumi.Input

	ToAccessTierPtrOutput() AccessTierPtrOutput
	ToAccessTierPtrOutputWithContext(context.Context) AccessTierPtrOutput
}

type accessTierPtr string

func AccessTierPtr(v string) AccessTierPtrInput {
	return (*accessTierPtr)(&v)
}

func (*accessTierPtr) ElementType() reflect.Type {
	return accessTierPtrType
}

func (in *accessTierPtr) ToAccessTierPtrOutput() AccessTierPtrOutput {
	return pulumi.ToOutput(in).(AccessTierPtrOutput)
}

func (in *accessTierPtr) ToAccessTierPtrOutputWithContext(ctx context.Context) AccessTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessTierPtrOutput)
}

type Action string

const (
	ActionAllow = Action("Allow")
)

func (Action) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (e Action) ToActionOutput() ActionOutput {
	return pulumi.ToOutput(e).(ActionOutput)
}

func (e Action) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionOutput)
}

func (e Action) ToActionPtrOutput() ActionPtrOutput {
	return e.ToActionPtrOutputWithContext(context.Background())
}

func (e Action) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return Action(e).ToActionOutputWithContext(ctx).ToActionPtrOutputWithContext(ctx)
}

func (e Action) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Action) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Action) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Action) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

func (o ActionOutput) ToActionPtrOutput() ActionPtrOutput {
	return o.ToActionPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Action) *Action {
		return &v
	}).(ActionPtrOutput)
}

func (o ActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Action) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Action) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionPtrOutput struct{ *pulumi.OutputState }

func (ActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (o ActionPtrOutput) ToActionPtrOutput() ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) Elem() ActionOutput {
	return o.ApplyT(func(v *Action) Action {
		if v != nil {
			return *v
		}
		var ret Action
		return ret
	}).(ActionOutput)
}

func (o ActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Action) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(context.Context) ActionOutput
}

var actionPtrType = reflect.TypeOf((**Action)(nil)).Elem()

type ActionPtrInput interface {
	pulumi.Input

	ToActionPtrOutput() ActionPtrOutput
	ToActionPtrOutputWithContext(context.Context) ActionPtrOutput
}

type actionPtr string

func ActionPtr(v string) ActionPtrInput {
	return (*actionPtr)(&v)
}

func (*actionPtr) ElementType() reflect.Type {
	return actionPtrType
}

func (in *actionPtr) ToActionPtrOutput() ActionPtrOutput {
	return pulumi.ToOutput(in).(ActionPtrOutput)
}

func (in *actionPtr) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionPtrOutput)
}

type Bypass string

const (
	BypassNone          = Bypass("None")
	BypassLogging       = Bypass("Logging")
	BypassMetrics       = Bypass("Metrics")
	BypassAzureServices = Bypass("AzureServices")
)

type DefaultAction string

const (
	DefaultActionAllow = DefaultAction("Allow")
	DefaultActionDeny  = DefaultAction("Deny")
)

func (DefaultAction) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAction)(nil)).Elem()
}

func (e DefaultAction) ToDefaultActionOutput() DefaultActionOutput {
	return pulumi.ToOutput(e).(DefaultActionOutput)
}

func (e DefaultAction) ToDefaultActionOutputWithContext(ctx context.Context) DefaultActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DefaultActionOutput)
}

func (e DefaultAction) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return e.ToDefaultActionPtrOutputWithContext(context.Background())
}

func (e DefaultAction) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return DefaultAction(e).ToDefaultActionOutputWithContext(ctx).ToDefaultActionPtrOutputWithContext(ctx)
}

func (e DefaultAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefaultAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DefaultActionOutput struct{ *pulumi.OutputState }

func (DefaultActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAction)(nil)).Elem()
}

func (o DefaultActionOutput) ToDefaultActionOutput() DefaultActionOutput {
	return o
}

func (o DefaultActionOutput) ToDefaultActionOutputWithContext(ctx context.Context) DefaultActionOutput {
	return o
}

func (o DefaultActionOutput) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return o.ToDefaultActionPtrOutputWithContext(context.Background())
}

func (o DefaultActionOutput) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultAction) *DefaultAction {
		return &v
	}).(DefaultActionPtrOutput)
}

func (o DefaultActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DefaultActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DefaultActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DefaultActionPtrOutput struct{ *pulumi.OutputState }

func (DefaultActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAction)(nil)).Elem()
}

func (o DefaultActionPtrOutput) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return o
}

func (o DefaultActionPtrOutput) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return o
}

func (o DefaultActionPtrOutput) Elem() DefaultActionOutput {
	return o.ApplyT(func(v *DefaultAction) DefaultAction {
		if v != nil {
			return *v
		}
		var ret DefaultAction
		return ret
	}).(DefaultActionOutput)
}

func (o DefaultActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DefaultAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DefaultActionInput interface {
	pulumi.Input

	ToDefaultActionOutput() DefaultActionOutput
	ToDefaultActionOutputWithContext(context.Context) DefaultActionOutput
}

var defaultActionPtrType = reflect.TypeOf((**DefaultAction)(nil)).Elem()

type DefaultActionPtrInput interface {
	pulumi.Input

	ToDefaultActionPtrOutput() DefaultActionPtrOutput
	ToDefaultActionPtrOutputWithContext(context.Context) DefaultActionPtrOutput
}

type defaultActionPtr string

func DefaultActionPtr(v string) DefaultActionPtrInput {
	return (*defaultActionPtr)(&v)
}

func (*defaultActionPtr) ElementType() reflect.Type {
	return defaultActionPtrType
}

func (in *defaultActionPtr) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return pulumi.ToOutput(in).(DefaultActionPtrOutput)
}

func (in *defaultActionPtr) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DefaultActionPtrOutput)
}

type DefaultSharePermission string

const (
	DefaultSharePermissionNone                                       = DefaultSharePermission("None")
	DefaultSharePermissionStorageFileDataSmbShareReader              = DefaultSharePermission("StorageFileDataSmbShareReader")
	DefaultSharePermissionStorageFileDataSmbShareContributor         = DefaultSharePermission("StorageFileDataSmbShareContributor")
	DefaultSharePermissionStorageFileDataSmbShareElevatedContributor = DefaultSharePermission("StorageFileDataSmbShareElevatedContributor")
	DefaultSharePermissionStorageFileDataSmbShareOwner               = DefaultSharePermission("StorageFileDataSmbShareOwner")
)

type DirectoryServiceOptions string

const (
	DirectoryServiceOptionsNone  = DirectoryServiceOptions("None")
	DirectoryServiceOptionsAADDS = DirectoryServiceOptions("AADDS")
	DirectoryServiceOptionsAD    = DirectoryServiceOptions("AD")
)

type EnabledProtocols string

const (
	EnabledProtocolsSMB = EnabledProtocols("SMB")
	EnabledProtocolsNFS = EnabledProtocols("NFS")
)

type EncryptionScopeSource string

const (
	EncryptionScopeSource_Microsoft_Storage  = EncryptionScopeSource("Microsoft.Storage")
	EncryptionScopeSource_Microsoft_KeyVault = EncryptionScopeSource("Microsoft.KeyVault")
)

type EncryptionScopeStateEnum string

const (
	EncryptionScopeStateEnumEnabled  = EncryptionScopeStateEnum("Enabled")
	EncryptionScopeStateEnumDisabled = EncryptionScopeStateEnum("Disabled")
)

type ExpirationAction string

const (
	ExpirationActionLog = ExpirationAction("Log")
)

type ExtendedLocationTypes string

const (
	ExtendedLocationTypesEdgeZone = ExtendedLocationTypes("EdgeZone")
)

type Format string

const (
	FormatCsv     = Format("Csv")
	FormatParquet = Format("Parquet")
)

type HttpProtocol string

const (
	HttpProtocol_Https_http = HttpProtocol("https,http")
	HttpProtocolHttps       = HttpProtocol("https")
)

type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned,UserAssigned")
)

type InventoryRuleType string

const (
	InventoryRuleTypeInventory = InventoryRuleType("Inventory")
)

type KeySource string

const (
	KeySource_Microsoft_Storage  = KeySource("Microsoft.Storage")
	KeySource_Microsoft_Keyvault = KeySource("Microsoft.Keyvault")
)

type KeyType string

const (
	KeyTypeService = KeyType("Service")
	KeyTypeAccount = KeyType("Account")
)

type Kind string

const (
	KindStorage          = Kind("Storage")
	KindStorageV2        = Kind("StorageV2")
	KindBlobStorage      = Kind("BlobStorage")
	KindFileStorage      = Kind("FileStorage")
	KindBlockBlobStorage = Kind("BlockBlobStorage")
)

type LargeFileSharesState string

const (
	LargeFileSharesStateDisabled = LargeFileSharesState("Disabled")
	LargeFileSharesStateEnabled  = LargeFileSharesState("Enabled")
)

type MinimumTlsVersion string

const (
	MinimumTlsVersion_TLS1_0 = MinimumTlsVersion("TLS1_0")
	MinimumTlsVersion_TLS1_1 = MinimumTlsVersion("TLS1_1")
	MinimumTlsVersion_TLS1_2 = MinimumTlsVersion("TLS1_2")
)

type Name string

const (
	NameAccessTimeTracking = Name("AccessTimeTracking")
)

type ObjectType string

const (
	ObjectTypeBlob      = ObjectType("Blob")
	ObjectTypeContainer = ObjectType("Container")
)

type Permissions string

const (
	PermissionsR = Permissions("r")
	PermissionsD = Permissions("d")
	PermissionsW = Permissions("w")
	PermissionsL = Permissions("l")
	PermissionsA = Permissions("a")
	PermissionsC = Permissions("c")
	PermissionsU = Permissions("u")
	PermissionsP = Permissions("p")
)

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

type PublicAccess string

const (
	PublicAccessContainer = PublicAccess("Container")
	PublicAccessBlob      = PublicAccess("Blob")
	PublicAccessNone      = PublicAccess("None")
)

func (PublicAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicAccess)(nil)).Elem()
}

func (e PublicAccess) ToPublicAccessOutput() PublicAccessOutput {
	return pulumi.ToOutput(e).(PublicAccessOutput)
}

func (e PublicAccess) ToPublicAccessOutputWithContext(ctx context.Context) PublicAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicAccessOutput)
}

func (e PublicAccess) ToPublicAccessPtrOutput() PublicAccessPtrOutput {
	return e.ToPublicAccessPtrOutputWithContext(context.Background())
}

func (e PublicAccess) ToPublicAccessPtrOutputWithContext(ctx context.Context) PublicAccessPtrOutput {
	return PublicAccess(e).ToPublicAccessOutputWithContext(ctx).ToPublicAccessPtrOutputWithContext(ctx)
}

func (e PublicAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicAccessOutput struct{ *pulumi.OutputState }

func (PublicAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicAccess)(nil)).Elem()
}

func (o PublicAccessOutput) ToPublicAccessOutput() PublicAccessOutput {
	return o
}

func (o PublicAccessOutput) ToPublicAccessOutputWithContext(ctx context.Context) PublicAccessOutput {
	return o
}

func (o PublicAccessOutput) ToPublicAccessPtrOutput() PublicAccessPtrOutput {
	return o.ToPublicAccessPtrOutputWithContext(context.Background())
}

func (o PublicAccessOutput) ToPublicAccessPtrOutputWithContext(ctx context.Context) PublicAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicAccess) *PublicAccess {
		return &v
	}).(PublicAccessPtrOutput)
}

func (o PublicAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicAccessPtrOutput struct{ *pulumi.OutputState }

func (PublicAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicAccess)(nil)).Elem()
}

func (o PublicAccessPtrOutput) ToPublicAccessPtrOutput() PublicAccessPtrOutput {
	return o
}

func (o PublicAccessPtrOutput) ToPublicAccessPtrOutputWithContext(ctx context.Context) PublicAccessPtrOutput {
	return o
}

func (o PublicAccessPtrOutput) Elem() PublicAccessOutput {
	return o.ApplyT(func(v *PublicAccess) PublicAccess {
		if v != nil {
			return *v
		}
		var ret PublicAccess
		return ret
	}).(PublicAccessOutput)
}

func (o PublicAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicAccessInput interface {
	pulumi.Input

	ToPublicAccessOutput() PublicAccessOutput
	ToPublicAccessOutputWithContext(context.Context) PublicAccessOutput
}

var publicAccessPtrType = reflect.TypeOf((**PublicAccess)(nil)).Elem()

type PublicAccessPtrInput interface {
	pulumi.Input

	ToPublicAccessPtrOutput() PublicAccessPtrOutput
	ToPublicAccessPtrOutputWithContext(context.Context) PublicAccessPtrOutput
}

type publicAccessPtr string

func PublicAccessPtr(v string) PublicAccessPtrInput {
	return (*publicAccessPtr)(&v)
}

func (*publicAccessPtr) ElementType() reflect.Type {
	return publicAccessPtrType
}

func (in *publicAccessPtr) ToPublicAccessPtrOutput() PublicAccessPtrOutput {
	return pulumi.ToOutput(in).(PublicAccessPtrOutput)
}

func (in *publicAccessPtr) ToPublicAccessPtrOutputWithContext(ctx context.Context) PublicAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicAccessPtrOutput)
}

type RootSquashType string

const (
	RootSquashTypeNoRootSquash = RootSquashType("NoRootSquash")
	RootSquashTypeRootSquash   = RootSquashType("RootSquash")
	RootSquashTypeAllSquash    = RootSquashType("AllSquash")
)

type RoutingChoice string

const (
	RoutingChoiceMicrosoftRouting = RoutingChoice("MicrosoftRouting")
	RoutingChoiceInternetRouting  = RoutingChoice("InternetRouting")
)

type RuleType string

const (
	RuleTypeLifecycle = RuleType("Lifecycle")
)

type Schedule string

const (
	ScheduleDaily  = Schedule("Daily")
	ScheduleWeekly = Schedule("Weekly")
)

type Services string

const (
	ServicesB = Services("b")
	ServicesQ = Services("q")
	ServicesT = Services("t")
	ServicesF = Services("f")
)

type ShareAccessTier string

const (
	ShareAccessTierTransactionOptimized = ShareAccessTier("TransactionOptimized")
	ShareAccessTierHot                  = ShareAccessTier("Hot")
	ShareAccessTierCool                 = ShareAccessTier("Cool")
	ShareAccessTierPremium              = ShareAccessTier("Premium")
)

type SignedResource string

const (
	SignedResourceB = SignedResource("b")
	SignedResourceC = SignedResource("c")
	SignedResourceF = SignedResource("f")
	SignedResourceS = SignedResource("s")
)

type SignedResourceTypes string

const (
	SignedResourceTypesS = SignedResourceTypes("s")
	SignedResourceTypesC = SignedResourceTypes("c")
	SignedResourceTypesO = SignedResourceTypes("o")
)

type SkuName string

const (
	SkuName_Standard_LRS    = SkuName("Standard_LRS")
	SkuName_Standard_GRS    = SkuName("Standard_GRS")
	SkuName_Standard_RAGRS  = SkuName("Standard_RAGRS")
	SkuName_Standard_ZRS    = SkuName("Standard_ZRS")
	SkuName_Premium_LRS     = SkuName("Premium_LRS")
	SkuName_Premium_ZRS     = SkuName("Premium_ZRS")
	SkuName_Standard_GZRS   = SkuName("Standard_GZRS")
	SkuName_Standard_RAGZRS = SkuName("Standard_RAGZRS")
)

type State string

const (
	StateProvisioning         = State("Provisioning")
	StateDeprovisioning       = State("Deprovisioning")
	StateSucceeded            = State("Succeeded")
	StateFailed               = State("Failed")
	StateNetworkSourceDeleted = State("NetworkSourceDeleted")
)

func init() {
	pulumi.RegisterOutputType(AccessTierOutput{})
	pulumi.RegisterOutputType(AccessTierPtrOutput{})
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(DefaultActionOutput{})
	pulumi.RegisterOutputType(DefaultActionPtrOutput{})
	pulumi.RegisterOutputType(PublicAccessOutput{})
	pulumi.RegisterOutputType(PublicAccessPtrOutput{})
}
