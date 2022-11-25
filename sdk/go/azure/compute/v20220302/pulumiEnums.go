


package v20220302

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Architecture string

const (
	ArchitectureX64   = Architecture("x64")
	ArchitectureArm64 = Architecture("Arm64")
)

type CopyCompletionErrorReason string

const (
	// Indicates that the source snapshot was deleted while the background copy of the resource created via CopyStart operation was in progress.
	CopyCompletionErrorReasonCopySourceNotFound = CopyCompletionErrorReason("CopySourceNotFound")
)

type DataAccessAuthMode string

const (
	// When export/upload URL is used, the system checks if the user has an identity in Azure Active Directory and has necessary permissions to export/upload the data. Please refer to aka.ms/DisksAzureADAuth.
	DataAccessAuthModeAzureActiveDirectory = DataAccessAuthMode("AzureActiveDirectory")
	// No additional authentication would be performed when accessing export/upload URL.
	DataAccessAuthModeNone = DataAccessAuthMode("None")
)

type DiskCreateOption string

const (
	// Create an empty data disk of a size given by diskSizeGB.
	DiskCreateOptionEmpty = DiskCreateOption("Empty")
	// Disk will be attached to a VM.
	DiskCreateOptionAttach = DiskCreateOption("Attach")
	// Create a new disk from a platform image specified by the given imageReference or galleryImageReference.
	DiskCreateOptionFromImage = DiskCreateOption("FromImage")
	// Create a disk by importing from a blob specified by a sourceUri in a storage account specified by storageAccountId.
	DiskCreateOptionImport = DiskCreateOption("Import")
	// Create a new disk or snapshot by copying from a disk or snapshot specified by the given sourceResourceId.
	DiskCreateOptionCopy = DiskCreateOption("Copy")
	// Create a new disk by copying from a backup recovery point.
	DiskCreateOptionRestore = DiskCreateOption("Restore")
	// Create a new disk by obtaining a write token and using it to directly upload the contents of the disk.
	DiskCreateOptionUpload = DiskCreateOption("Upload")
	// Create a new disk by using a deep copy process, where the resource creation is considered complete only after all data has been copied from the source.
	DiskCreateOptionCopyStart = DiskCreateOption("CopyStart")
	// Similar to Import create option. Create a new Trusted Launch VM or Confidential VM supported disk by importing additional blob for VM guest state specified by securityDataUri in storage account specified by storageAccountId
	DiskCreateOptionImportSecure = DiskCreateOption("ImportSecure")
	// Similar to Upload create option. Create a new Trusted Launch VM or Confidential VM supported disk and upload using write token in both disk and VM guest state
	DiskCreateOptionUploadPreparedSecure = DiskCreateOption("UploadPreparedSecure")
)

type DiskEncryptionSetIdentityType string

const (
	DiskEncryptionSetIdentityTypeSystemAssigned               = DiskEncryptionSetIdentityType("SystemAssigned")
	DiskEncryptionSetIdentityTypeUserAssigned                 = DiskEncryptionSetIdentityType("UserAssigned")
	DiskEncryptionSetIdentityType_SystemAssigned_UserAssigned = DiskEncryptionSetIdentityType("SystemAssigned, UserAssigned")
	DiskEncryptionSetIdentityTypeNone                         = DiskEncryptionSetIdentityType("None")
)

type DiskEncryptionSetType string

const (
	// Resource using diskEncryptionSet would be encrypted at rest with Customer managed key that can be changed and revoked by a customer.
	DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey = DiskEncryptionSetType("EncryptionAtRestWithCustomerKey")
	// Resource using diskEncryptionSet would be encrypted at rest with two layers of encryption. One of the keys is Customer managed and the other key is Platform managed.
	DiskEncryptionSetTypeEncryptionAtRestWithPlatformAndCustomerKeys = DiskEncryptionSetType("EncryptionAtRestWithPlatformAndCustomerKeys")
	// Confidential VM supported disk and VM guest state would be encrypted with customer managed key.
	DiskEncryptionSetTypeConfidentialVmEncryptedWithCustomerKey = DiskEncryptionSetType("ConfidentialVmEncryptedWithCustomerKey")
)

type DiskSecurityTypes string

const (
	// Trusted Launch provides security features such as secure boot and virtual Trusted Platform Module (vTPM)
	DiskSecurityTypesTrustedLaunch = DiskSecurityTypes("TrustedLaunch")
	// Indicates Confidential VM disk with only VM guest state encrypted
	DiskSecurityTypes_ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey = DiskSecurityTypes("ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey")
	// Indicates Confidential VM disk with both OS disk and VM guest state encrypted with a platform managed key
	DiskSecurityTypes_ConfidentialVM_DiskEncryptedWithPlatformKey = DiskSecurityTypes("ConfidentialVM_DiskEncryptedWithPlatformKey")
	// Indicates Confidential VM disk with both OS disk and VM guest state encrypted with a customer managed key
	DiskSecurityTypes_ConfidentialVM_DiskEncryptedWithCustomerKey = DiskSecurityTypes("ConfidentialVM_DiskEncryptedWithCustomerKey")
)

type DiskStorageAccountTypes string

const (
	// Standard HDD locally redundant storage. Best for backup, non-critical, and infrequent access.
	DiskStorageAccountTypes_Standard_LRS = DiskStorageAccountTypes("Standard_LRS")
	// Premium SSD locally redundant storage. Best for production and performance sensitive workloads.
	DiskStorageAccountTypes_Premium_LRS = DiskStorageAccountTypes("Premium_LRS")
	// Standard SSD locally redundant storage. Best for web servers, lightly used enterprise applications and dev/test.
	DiskStorageAccountTypes_StandardSSD_LRS = DiskStorageAccountTypes("StandardSSD_LRS")
	// Ultra SSD locally redundant storage. Best for IO-intensive workloads such as SAP HANA, top tier databases (for example, SQL, Oracle), and other transaction-heavy workloads.
	DiskStorageAccountTypes_UltraSSD_LRS = DiskStorageAccountTypes("UltraSSD_LRS")
	// Premium SSD zone redundant storage. Best for the production workloads that need storage resiliency against zone failures.
	DiskStorageAccountTypes_Premium_ZRS = DiskStorageAccountTypes("Premium_ZRS")
	// Standard SSD zone redundant storage. Best for web servers, lightly used enterprise applications and dev/test that need storage resiliency against zone failures.
	DiskStorageAccountTypes_StandardSSD_ZRS = DiskStorageAccountTypes("StandardSSD_ZRS")
	// Premium SSD v2 locally redundant storage. Best for production and performance-sensitive workloads that consistently require low latency and high IOPS and throughput.
	DiskStorageAccountTypes_PremiumV2_LRS = DiskStorageAccountTypes("PremiumV2_LRS")
)

type EncryptionType string

const (
	// Disk is encrypted at rest with Platform managed key. It is the default encryption type. This is not a valid encryption type for disk encryption sets.
	EncryptionTypeEncryptionAtRestWithPlatformKey = EncryptionType("EncryptionAtRestWithPlatformKey")
	// Disk is encrypted at rest with Customer managed key that can be changed and revoked by a customer.
	EncryptionTypeEncryptionAtRestWithCustomerKey = EncryptionType("EncryptionAtRestWithCustomerKey")
	// Disk is encrypted at rest with 2 layers of encryption. One of the keys is Customer managed and the other key is Platform managed.
	EncryptionTypeEncryptionAtRestWithPlatformAndCustomerKeys = EncryptionType("EncryptionAtRestWithPlatformAndCustomerKeys")
)

type ExtendedLocationTypes string

const (
	ExtendedLocationTypesEdgeZone = ExtendedLocationTypes("EdgeZone")
)

type HyperVGeneration string

const (
	HyperVGenerationV1 = HyperVGeneration("V1")
	HyperVGenerationV2 = HyperVGeneration("V2")
)

type NetworkAccessPolicy string

const (
	// The disk can be exported or uploaded to from any network.
	NetworkAccessPolicyAllowAll = NetworkAccessPolicy("AllowAll")
	// The disk can be exported or uploaded to using a DiskAccess resource's private endpoints.
	NetworkAccessPolicyAllowPrivate = NetworkAccessPolicy("AllowPrivate")
	// The disk cannot be exported.
	NetworkAccessPolicyDenyAll = NetworkAccessPolicy("DenyAll")
)

type OperatingSystemTypes string

const (
	OperatingSystemTypesWindows = OperatingSystemTypes("Windows")
	OperatingSystemTypesLinux   = OperatingSystemTypes("Linux")
)

func (OperatingSystemTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemTypes)(nil)).Elem()
}

func (e OperatingSystemTypes) ToOperatingSystemTypesOutput() OperatingSystemTypesOutput {
	return pulumi.ToOutput(e).(OperatingSystemTypesOutput)
}

func (e OperatingSystemTypes) ToOperatingSystemTypesOutputWithContext(ctx context.Context) OperatingSystemTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatingSystemTypesOutput)
}

func (e OperatingSystemTypes) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return e.ToOperatingSystemTypesPtrOutputWithContext(context.Background())
}

func (e OperatingSystemTypes) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return OperatingSystemTypes(e).ToOperatingSystemTypesOutputWithContext(ctx).ToOperatingSystemTypesPtrOutputWithContext(ctx)
}

func (e OperatingSystemTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatingSystemTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatingSystemTypesOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemTypes)(nil)).Elem()
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesOutput() OperatingSystemTypesOutput {
	return o
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesOutputWithContext(ctx context.Context) OperatingSystemTypesOutput {
	return o
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return o.ToOperatingSystemTypesPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatingSystemTypes) *OperatingSystemTypes {
		return &v
	}).(OperatingSystemTypesPtrOutput)
}

func (o OperatingSystemTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatingSystemTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatingSystemTypesPtrOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatingSystemTypes)(nil)).Elem()
}

func (o OperatingSystemTypesPtrOutput) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return o
}

func (o OperatingSystemTypesPtrOutput) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return o
}

func (o OperatingSystemTypesPtrOutput) Elem() OperatingSystemTypesOutput {
	return o.ApplyT(func(v *OperatingSystemTypes) OperatingSystemTypes {
		if v != nil {
			return *v
		}
		var ret OperatingSystemTypes
		return ret
	}).(OperatingSystemTypesOutput)
}

func (o OperatingSystemTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatingSystemTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatingSystemTypesInput interface {
	pulumi.Input

	ToOperatingSystemTypesOutput() OperatingSystemTypesOutput
	ToOperatingSystemTypesOutputWithContext(context.Context) OperatingSystemTypesOutput
}

var operatingSystemTypesPtrType = reflect.TypeOf((**OperatingSystemTypes)(nil)).Elem()

type OperatingSystemTypesPtrInput interface {
	pulumi.Input

	ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput
	ToOperatingSystemTypesPtrOutputWithContext(context.Context) OperatingSystemTypesPtrOutput
}

type operatingSystemTypesPtr string

func OperatingSystemTypesPtr(v string) OperatingSystemTypesPtrInput {
	return (*operatingSystemTypesPtr)(&v)
}

func (*operatingSystemTypesPtr) ElementType() reflect.Type {
	return operatingSystemTypesPtrType
}

func (in *operatingSystemTypesPtr) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return pulumi.ToOutput(in).(OperatingSystemTypesPtrOutput)
}

func (in *operatingSystemTypesPtr) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatingSystemTypesPtrOutput)
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

type PublicNetworkAccess string

const (
	// You can generate a SAS URI to access the underlying data of the disk publicly on the internet when NetworkAccessPolicy is set to AllowAll. You can access the data via the SAS URI only from your trusted Azure VNET when NetworkAccessPolicy is set to AllowPrivate.
	PublicNetworkAccessEnabled = PublicNetworkAccess("Enabled")
	// You cannot access the underlying data of the disk publicly on the internet even when NetworkAccessPolicy is set to AllowAll. You can access the data via the SAS URI only from your trusted Azure VNET when NetworkAccessPolicy is set to AllowPrivate.
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

type SnapshotStorageAccountTypes string

const (
	// Standard HDD locally redundant storage
	SnapshotStorageAccountTypes_Standard_LRS = SnapshotStorageAccountTypes("Standard_LRS")
	// Premium SSD locally redundant storage
	SnapshotStorageAccountTypes_Premium_LRS = SnapshotStorageAccountTypes("Premium_LRS")
	// Standard zone redundant storage
	SnapshotStorageAccountTypes_Standard_ZRS = SnapshotStorageAccountTypes("Standard_ZRS")
)

func init() {
	pulumi.RegisterOutputType(OperatingSystemTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesPtrOutput{})
}
