


package v20200930

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
)

type DiskEncryptionSetIdentityType string

const (
	DiskEncryptionSetIdentityTypeSystemAssigned = DiskEncryptionSetIdentityType("SystemAssigned")
	DiskEncryptionSetIdentityTypeNone           = DiskEncryptionSetIdentityType("None")
)

type DiskEncryptionSetType string

const (
	// Resource using diskEncryptionSet would be encrypted at rest with Customer managed key that can be changed and revoked by a customer.
	DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey = DiskEncryptionSetType("EncryptionAtRestWithCustomerKey")
	// Resource using diskEncryptionSet would be encrypted at rest with two layers of encryption. One of the keys is Customer managed and the other key is Platform managed.
	DiskEncryptionSetTypeEncryptionAtRestWithPlatformAndCustomerKeys = DiskEncryptionSetType("EncryptionAtRestWithPlatformAndCustomerKeys")
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

type GallerySharingPermissionTypes string

const (
	GallerySharingPermissionTypesPrivate = GallerySharingPermissionTypes("Private")
	GallerySharingPermissionTypesGroups  = GallerySharingPermissionTypes("Groups")
)

type HostCaching string

const (
	HostCachingNone      = HostCaching("None")
	HostCachingReadOnly  = HostCaching("ReadOnly")
	HostCachingReadWrite = HostCaching("ReadWrite")
)

func (HostCaching) ElementType() reflect.Type {
	return reflect.TypeOf((*HostCaching)(nil)).Elem()
}

func (e HostCaching) ToHostCachingOutput() HostCachingOutput {
	return pulumi.ToOutput(e).(HostCachingOutput)
}

func (e HostCaching) ToHostCachingOutputWithContext(ctx context.Context) HostCachingOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostCachingOutput)
}

func (e HostCaching) ToHostCachingPtrOutput() HostCachingPtrOutput {
	return e.ToHostCachingPtrOutputWithContext(context.Background())
}

func (e HostCaching) ToHostCachingPtrOutputWithContext(ctx context.Context) HostCachingPtrOutput {
	return HostCaching(e).ToHostCachingOutputWithContext(ctx).ToHostCachingPtrOutputWithContext(ctx)
}

func (e HostCaching) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostCaching) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostCaching) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostCaching) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostCachingOutput struct{ *pulumi.OutputState }

func (HostCachingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostCaching)(nil)).Elem()
}

func (o HostCachingOutput) ToHostCachingOutput() HostCachingOutput {
	return o
}

func (o HostCachingOutput) ToHostCachingOutputWithContext(ctx context.Context) HostCachingOutput {
	return o
}

func (o HostCachingOutput) ToHostCachingPtrOutput() HostCachingPtrOutput {
	return o.ToHostCachingPtrOutputWithContext(context.Background())
}

func (o HostCachingOutput) ToHostCachingPtrOutputWithContext(ctx context.Context) HostCachingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostCaching) *HostCaching {
		return &v
	}).(HostCachingPtrOutput)
}

func (o HostCachingOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostCachingOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostCaching) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostCachingOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostCachingOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostCaching) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostCachingPtrOutput struct{ *pulumi.OutputState }

func (HostCachingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostCaching)(nil)).Elem()
}

func (o HostCachingPtrOutput) ToHostCachingPtrOutput() HostCachingPtrOutput {
	return o
}

func (o HostCachingPtrOutput) ToHostCachingPtrOutputWithContext(ctx context.Context) HostCachingPtrOutput {
	return o
}

func (o HostCachingPtrOutput) Elem() HostCachingOutput {
	return o.ApplyT(func(v *HostCaching) HostCaching {
		if v != nil {
			return *v
		}
		var ret HostCaching
		return ret
	}).(HostCachingOutput)
}

func (o HostCachingPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostCachingPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostCaching) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostCachingInput interface {
	pulumi.Input

	ToHostCachingOutput() HostCachingOutput
	ToHostCachingOutputWithContext(context.Context) HostCachingOutput
}

var hostCachingPtrType = reflect.TypeOf((**HostCaching)(nil)).Elem()

type HostCachingPtrInput interface {
	pulumi.Input

	ToHostCachingPtrOutput() HostCachingPtrOutput
	ToHostCachingPtrOutputWithContext(context.Context) HostCachingPtrOutput
}

type hostCachingPtr string

func HostCachingPtr(v string) HostCachingPtrInput {
	return (*hostCachingPtr)(&v)
}

func (*hostCachingPtr) ElementType() reflect.Type {
	return hostCachingPtrType
}

func (in *hostCachingPtr) ToHostCachingPtrOutput() HostCachingPtrOutput {
	return pulumi.ToOutput(in).(HostCachingPtrOutput)
}

func (in *hostCachingPtr) ToHostCachingPtrOutputWithContext(ctx context.Context) HostCachingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostCachingPtrOutput)
}

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

type OperatingSystemStateTypes string

const (
	OperatingSystemStateTypesGeneralized = OperatingSystemStateTypes("Generalized")
	OperatingSystemStateTypesSpecialized = OperatingSystemStateTypes("Specialized")
)

func (OperatingSystemStateTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemStateTypes)(nil)).Elem()
}

func (e OperatingSystemStateTypes) ToOperatingSystemStateTypesOutput() OperatingSystemStateTypesOutput {
	return pulumi.ToOutput(e).(OperatingSystemStateTypesOutput)
}

func (e OperatingSystemStateTypes) ToOperatingSystemStateTypesOutputWithContext(ctx context.Context) OperatingSystemStateTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatingSystemStateTypesOutput)
}

func (e OperatingSystemStateTypes) ToOperatingSystemStateTypesPtrOutput() OperatingSystemStateTypesPtrOutput {
	return e.ToOperatingSystemStateTypesPtrOutputWithContext(context.Background())
}

func (e OperatingSystemStateTypes) ToOperatingSystemStateTypesPtrOutputWithContext(ctx context.Context) OperatingSystemStateTypesPtrOutput {
	return OperatingSystemStateTypes(e).ToOperatingSystemStateTypesOutputWithContext(ctx).ToOperatingSystemStateTypesPtrOutputWithContext(ctx)
}

func (e OperatingSystemStateTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemStateTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemStateTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatingSystemStateTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatingSystemStateTypesOutput struct{ *pulumi.OutputState }

func (OperatingSystemStateTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemStateTypes)(nil)).Elem()
}

func (o OperatingSystemStateTypesOutput) ToOperatingSystemStateTypesOutput() OperatingSystemStateTypesOutput {
	return o
}

func (o OperatingSystemStateTypesOutput) ToOperatingSystemStateTypesOutputWithContext(ctx context.Context) OperatingSystemStateTypesOutput {
	return o
}

func (o OperatingSystemStateTypesOutput) ToOperatingSystemStateTypesPtrOutput() OperatingSystemStateTypesPtrOutput {
	return o.ToOperatingSystemStateTypesPtrOutputWithContext(context.Background())
}

func (o OperatingSystemStateTypesOutput) ToOperatingSystemStateTypesPtrOutputWithContext(ctx context.Context) OperatingSystemStateTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatingSystemStateTypes) *OperatingSystemStateTypes {
		return &v
	}).(OperatingSystemStateTypesPtrOutput)
}

func (o OperatingSystemStateTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatingSystemStateTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemStateTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatingSystemStateTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemStateTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemStateTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatingSystemStateTypesPtrOutput struct{ *pulumi.OutputState }

func (OperatingSystemStateTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatingSystemStateTypes)(nil)).Elem()
}

func (o OperatingSystemStateTypesPtrOutput) ToOperatingSystemStateTypesPtrOutput() OperatingSystemStateTypesPtrOutput {
	return o
}

func (o OperatingSystemStateTypesPtrOutput) ToOperatingSystemStateTypesPtrOutputWithContext(ctx context.Context) OperatingSystemStateTypesPtrOutput {
	return o
}

func (o OperatingSystemStateTypesPtrOutput) Elem() OperatingSystemStateTypesOutput {
	return o.ApplyT(func(v *OperatingSystemStateTypes) OperatingSystemStateTypes {
		if v != nil {
			return *v
		}
		var ret OperatingSystemStateTypes
		return ret
	}).(OperatingSystemStateTypesOutput)
}

func (o OperatingSystemStateTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemStateTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatingSystemStateTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatingSystemStateTypesInput interface {
	pulumi.Input

	ToOperatingSystemStateTypesOutput() OperatingSystemStateTypesOutput
	ToOperatingSystemStateTypesOutputWithContext(context.Context) OperatingSystemStateTypesOutput
}

var operatingSystemStateTypesPtrType = reflect.TypeOf((**OperatingSystemStateTypes)(nil)).Elem()

type OperatingSystemStateTypesPtrInput interface {
	pulumi.Input

	ToOperatingSystemStateTypesPtrOutput() OperatingSystemStateTypesPtrOutput
	ToOperatingSystemStateTypesPtrOutputWithContext(context.Context) OperatingSystemStateTypesPtrOutput
}

type operatingSystemStateTypesPtr string

func OperatingSystemStateTypesPtr(v string) OperatingSystemStateTypesPtrInput {
	return (*operatingSystemStateTypesPtr)(&v)
}

func (*operatingSystemStateTypesPtr) ElementType() reflect.Type {
	return operatingSystemStateTypesPtrType
}

func (in *operatingSystemStateTypesPtr) ToOperatingSystemStateTypesPtrOutput() OperatingSystemStateTypesPtrOutput {
	return pulumi.ToOutput(in).(OperatingSystemStateTypesPtrOutput)
}

func (in *operatingSystemStateTypesPtr) ToOperatingSystemStateTypesPtrOutputWithContext(ctx context.Context) OperatingSystemStateTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatingSystemStateTypesPtrOutput)
}

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

type SnapshotStorageAccountTypes string

const (
	// Standard HDD locally redundant storage
	SnapshotStorageAccountTypes_Standard_LRS = SnapshotStorageAccountTypes("Standard_LRS")
	// Premium SSD locally redundant storage
	SnapshotStorageAccountTypes_Premium_LRS = SnapshotStorageAccountTypes("Premium_LRS")
	// Standard zone redundant storage
	SnapshotStorageAccountTypes_Standard_ZRS = SnapshotStorageAccountTypes("Standard_ZRS")
)

type StorageAccountType string

const (
	StorageAccountType_Standard_LRS = StorageAccountType("Standard_LRS")
	StorageAccountType_Standard_ZRS = StorageAccountType("Standard_ZRS")
	StorageAccountType_Premium_LRS  = StorageAccountType("Premium_LRS")
)

func init() {
	pulumi.RegisterOutputType(HostCachingOutput{})
	pulumi.RegisterOutputType(HostCachingPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemStateTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemStateTypesPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesPtrOutput{})
}
