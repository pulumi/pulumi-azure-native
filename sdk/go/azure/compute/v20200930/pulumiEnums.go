


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

func (DiskCreateOption) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskCreateOption)(nil)).Elem()
}

func (e DiskCreateOption) ToDiskCreateOptionOutput() DiskCreateOptionOutput {
	return pulumi.ToOutput(e).(DiskCreateOptionOutput)
}

func (e DiskCreateOption) ToDiskCreateOptionOutputWithContext(ctx context.Context) DiskCreateOptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskCreateOptionOutput)
}

func (e DiskCreateOption) ToDiskCreateOptionPtrOutput() DiskCreateOptionPtrOutput {
	return e.ToDiskCreateOptionPtrOutputWithContext(context.Background())
}

func (e DiskCreateOption) ToDiskCreateOptionPtrOutputWithContext(ctx context.Context) DiskCreateOptionPtrOutput {
	return DiskCreateOption(e).ToDiskCreateOptionOutputWithContext(ctx).ToDiskCreateOptionPtrOutputWithContext(ctx)
}

func (e DiskCreateOption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskCreateOption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskCreateOption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskCreateOption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskCreateOptionOutput struct{ *pulumi.OutputState }

func (DiskCreateOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskCreateOption)(nil)).Elem()
}

func (o DiskCreateOptionOutput) ToDiskCreateOptionOutput() DiskCreateOptionOutput {
	return o
}

func (o DiskCreateOptionOutput) ToDiskCreateOptionOutputWithContext(ctx context.Context) DiskCreateOptionOutput {
	return o
}

func (o DiskCreateOptionOutput) ToDiskCreateOptionPtrOutput() DiskCreateOptionPtrOutput {
	return o.ToDiskCreateOptionPtrOutputWithContext(context.Background())
}

func (o DiskCreateOptionOutput) ToDiskCreateOptionPtrOutputWithContext(ctx context.Context) DiskCreateOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskCreateOption) *DiskCreateOption {
		return &v
	}).(DiskCreateOptionPtrOutput)
}

func (o DiskCreateOptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskCreateOptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskCreateOption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskCreateOptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskCreateOptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskCreateOption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskCreateOptionPtrOutput struct{ *pulumi.OutputState }

func (DiskCreateOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskCreateOption)(nil)).Elem()
}

func (o DiskCreateOptionPtrOutput) ToDiskCreateOptionPtrOutput() DiskCreateOptionPtrOutput {
	return o
}

func (o DiskCreateOptionPtrOutput) ToDiskCreateOptionPtrOutputWithContext(ctx context.Context) DiskCreateOptionPtrOutput {
	return o
}

func (o DiskCreateOptionPtrOutput) Elem() DiskCreateOptionOutput {
	return o.ApplyT(func(v *DiskCreateOption) DiskCreateOption {
		if v != nil {
			return *v
		}
		var ret DiskCreateOption
		return ret
	}).(DiskCreateOptionOutput)
}

func (o DiskCreateOptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskCreateOptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskCreateOption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiskCreateOptionInput interface {
	pulumi.Input

	ToDiskCreateOptionOutput() DiskCreateOptionOutput
	ToDiskCreateOptionOutputWithContext(context.Context) DiskCreateOptionOutput
}

var diskCreateOptionPtrType = reflect.TypeOf((**DiskCreateOption)(nil)).Elem()

type DiskCreateOptionPtrInput interface {
	pulumi.Input

	ToDiskCreateOptionPtrOutput() DiskCreateOptionPtrOutput
	ToDiskCreateOptionPtrOutputWithContext(context.Context) DiskCreateOptionPtrOutput
}

type diskCreateOptionPtr string

func DiskCreateOptionPtr(v string) DiskCreateOptionPtrInput {
	return (*diskCreateOptionPtr)(&v)
}

func (*diskCreateOptionPtr) ElementType() reflect.Type {
	return diskCreateOptionPtrType
}

func (in *diskCreateOptionPtr) ToDiskCreateOptionPtrOutput() DiskCreateOptionPtrOutput {
	return pulumi.ToOutput(in).(DiskCreateOptionPtrOutput)
}

func (in *diskCreateOptionPtr) ToDiskCreateOptionPtrOutputWithContext(ctx context.Context) DiskCreateOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskCreateOptionPtrOutput)
}

type DiskEncryptionSetIdentityType string

const (
	DiskEncryptionSetIdentityTypeSystemAssigned = DiskEncryptionSetIdentityType("SystemAssigned")
	DiskEncryptionSetIdentityTypeNone           = DiskEncryptionSetIdentityType("None")
)

func (DiskEncryptionSetIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetIdentityType)(nil)).Elem()
}

func (e DiskEncryptionSetIdentityType) ToDiskEncryptionSetIdentityTypeOutput() DiskEncryptionSetIdentityTypeOutput {
	return pulumi.ToOutput(e).(DiskEncryptionSetIdentityTypeOutput)
}

func (e DiskEncryptionSetIdentityType) ToDiskEncryptionSetIdentityTypeOutputWithContext(ctx context.Context) DiskEncryptionSetIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskEncryptionSetIdentityTypeOutput)
}

func (e DiskEncryptionSetIdentityType) ToDiskEncryptionSetIdentityTypePtrOutput() DiskEncryptionSetIdentityTypePtrOutput {
	return e.ToDiskEncryptionSetIdentityTypePtrOutputWithContext(context.Background())
}

func (e DiskEncryptionSetIdentityType) ToDiskEncryptionSetIdentityTypePtrOutputWithContext(ctx context.Context) DiskEncryptionSetIdentityTypePtrOutput {
	return DiskEncryptionSetIdentityType(e).ToDiskEncryptionSetIdentityTypeOutputWithContext(ctx).ToDiskEncryptionSetIdentityTypePtrOutputWithContext(ctx)
}

func (e DiskEncryptionSetIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskEncryptionSetIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskEncryptionSetIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskEncryptionSetIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskEncryptionSetIdentityTypeOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetIdentityType)(nil)).Elem()
}

func (o DiskEncryptionSetIdentityTypeOutput) ToDiskEncryptionSetIdentityTypeOutput() DiskEncryptionSetIdentityTypeOutput {
	return o
}

func (o DiskEncryptionSetIdentityTypeOutput) ToDiskEncryptionSetIdentityTypeOutputWithContext(ctx context.Context) DiskEncryptionSetIdentityTypeOutput {
	return o
}

func (o DiskEncryptionSetIdentityTypeOutput) ToDiskEncryptionSetIdentityTypePtrOutput() DiskEncryptionSetIdentityTypePtrOutput {
	return o.ToDiskEncryptionSetIdentityTypePtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSetIdentityTypeOutput) ToDiskEncryptionSetIdentityTypePtrOutputWithContext(ctx context.Context) DiskEncryptionSetIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskEncryptionSetIdentityType) *DiskEncryptionSetIdentityType {
		return &v
	}).(DiskEncryptionSetIdentityTypePtrOutput)
}

func (o DiskEncryptionSetIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskEncryptionSetIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskEncryptionSetIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskEncryptionSetIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSetIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskEncryptionSetIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskEncryptionSetIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSetIdentityType)(nil)).Elem()
}

func (o DiskEncryptionSetIdentityTypePtrOutput) ToDiskEncryptionSetIdentityTypePtrOutput() DiskEncryptionSetIdentityTypePtrOutput {
	return o
}

func (o DiskEncryptionSetIdentityTypePtrOutput) ToDiskEncryptionSetIdentityTypePtrOutputWithContext(ctx context.Context) DiskEncryptionSetIdentityTypePtrOutput {
	return o
}

func (o DiskEncryptionSetIdentityTypePtrOutput) Elem() DiskEncryptionSetIdentityTypeOutput {
	return o.ApplyT(func(v *DiskEncryptionSetIdentityType) DiskEncryptionSetIdentityType {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSetIdentityType
		return ret
	}).(DiskEncryptionSetIdentityTypeOutput)
}

func (o DiskEncryptionSetIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSetIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskEncryptionSetIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiskEncryptionSetIdentityTypeInput interface {
	pulumi.Input

	ToDiskEncryptionSetIdentityTypeOutput() DiskEncryptionSetIdentityTypeOutput
	ToDiskEncryptionSetIdentityTypeOutputWithContext(context.Context) DiskEncryptionSetIdentityTypeOutput
}

var diskEncryptionSetIdentityTypePtrType = reflect.TypeOf((**DiskEncryptionSetIdentityType)(nil)).Elem()

type DiskEncryptionSetIdentityTypePtrInput interface {
	pulumi.Input

	ToDiskEncryptionSetIdentityTypePtrOutput() DiskEncryptionSetIdentityTypePtrOutput
	ToDiskEncryptionSetIdentityTypePtrOutputWithContext(context.Context) DiskEncryptionSetIdentityTypePtrOutput
}

type diskEncryptionSetIdentityTypePtr string

func DiskEncryptionSetIdentityTypePtr(v string) DiskEncryptionSetIdentityTypePtrInput {
	return (*diskEncryptionSetIdentityTypePtr)(&v)
}

func (*diskEncryptionSetIdentityTypePtr) ElementType() reflect.Type {
	return diskEncryptionSetIdentityTypePtrType
}

func (in *diskEncryptionSetIdentityTypePtr) ToDiskEncryptionSetIdentityTypePtrOutput() DiskEncryptionSetIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(DiskEncryptionSetIdentityTypePtrOutput)
}

func (in *diskEncryptionSetIdentityTypePtr) ToDiskEncryptionSetIdentityTypePtrOutputWithContext(ctx context.Context) DiskEncryptionSetIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskEncryptionSetIdentityTypePtrOutput)
}

type DiskEncryptionSetType string

const (
	// Resource using diskEncryptionSet would be encrypted at rest with Customer managed key that can be changed and revoked by a customer.
	DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey = DiskEncryptionSetType("EncryptionAtRestWithCustomerKey")
	// Resource using diskEncryptionSet would be encrypted at rest with two layers of encryption. One of the keys is Customer managed and the other key is Platform managed.
	DiskEncryptionSetTypeEncryptionAtRestWithPlatformAndCustomerKeys = DiskEncryptionSetType("EncryptionAtRestWithPlatformAndCustomerKeys")
)

func (DiskEncryptionSetType) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetType)(nil)).Elem()
}

func (e DiskEncryptionSetType) ToDiskEncryptionSetTypeOutput() DiskEncryptionSetTypeOutput {
	return pulumi.ToOutput(e).(DiskEncryptionSetTypeOutput)
}

func (e DiskEncryptionSetType) ToDiskEncryptionSetTypeOutputWithContext(ctx context.Context) DiskEncryptionSetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskEncryptionSetTypeOutput)
}

func (e DiskEncryptionSetType) ToDiskEncryptionSetTypePtrOutput() DiskEncryptionSetTypePtrOutput {
	return e.ToDiskEncryptionSetTypePtrOutputWithContext(context.Background())
}

func (e DiskEncryptionSetType) ToDiskEncryptionSetTypePtrOutputWithContext(ctx context.Context) DiskEncryptionSetTypePtrOutput {
	return DiskEncryptionSetType(e).ToDiskEncryptionSetTypeOutputWithContext(ctx).ToDiskEncryptionSetTypePtrOutputWithContext(ctx)
}

func (e DiskEncryptionSetType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskEncryptionSetType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskEncryptionSetType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskEncryptionSetType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskEncryptionSetTypeOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetType)(nil)).Elem()
}

func (o DiskEncryptionSetTypeOutput) ToDiskEncryptionSetTypeOutput() DiskEncryptionSetTypeOutput {
	return o
}

func (o DiskEncryptionSetTypeOutput) ToDiskEncryptionSetTypeOutputWithContext(ctx context.Context) DiskEncryptionSetTypeOutput {
	return o
}

func (o DiskEncryptionSetTypeOutput) ToDiskEncryptionSetTypePtrOutput() DiskEncryptionSetTypePtrOutput {
	return o.ToDiskEncryptionSetTypePtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSetTypeOutput) ToDiskEncryptionSetTypePtrOutputWithContext(ctx context.Context) DiskEncryptionSetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskEncryptionSetType) *DiskEncryptionSetType {
		return &v
	}).(DiskEncryptionSetTypePtrOutput)
}

func (o DiskEncryptionSetTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskEncryptionSetTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskEncryptionSetType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskEncryptionSetTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSetTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskEncryptionSetType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskEncryptionSetTypePtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSetType)(nil)).Elem()
}

func (o DiskEncryptionSetTypePtrOutput) ToDiskEncryptionSetTypePtrOutput() DiskEncryptionSetTypePtrOutput {
	return o
}

func (o DiskEncryptionSetTypePtrOutput) ToDiskEncryptionSetTypePtrOutputWithContext(ctx context.Context) DiskEncryptionSetTypePtrOutput {
	return o
}

func (o DiskEncryptionSetTypePtrOutput) Elem() DiskEncryptionSetTypeOutput {
	return o.ApplyT(func(v *DiskEncryptionSetType) DiskEncryptionSetType {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSetType
		return ret
	}).(DiskEncryptionSetTypeOutput)
}

func (o DiskEncryptionSetTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSetTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskEncryptionSetType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiskEncryptionSetTypeInput interface {
	pulumi.Input

	ToDiskEncryptionSetTypeOutput() DiskEncryptionSetTypeOutput
	ToDiskEncryptionSetTypeOutputWithContext(context.Context) DiskEncryptionSetTypeOutput
}

var diskEncryptionSetTypePtrType = reflect.TypeOf((**DiskEncryptionSetType)(nil)).Elem()

type DiskEncryptionSetTypePtrInput interface {
	pulumi.Input

	ToDiskEncryptionSetTypePtrOutput() DiskEncryptionSetTypePtrOutput
	ToDiskEncryptionSetTypePtrOutputWithContext(context.Context) DiskEncryptionSetTypePtrOutput
}

type diskEncryptionSetTypePtr string

func DiskEncryptionSetTypePtr(v string) DiskEncryptionSetTypePtrInput {
	return (*diskEncryptionSetTypePtr)(&v)
}

func (*diskEncryptionSetTypePtr) ElementType() reflect.Type {
	return diskEncryptionSetTypePtrType
}

func (in *diskEncryptionSetTypePtr) ToDiskEncryptionSetTypePtrOutput() DiskEncryptionSetTypePtrOutput {
	return pulumi.ToOutput(in).(DiskEncryptionSetTypePtrOutput)
}

func (in *diskEncryptionSetTypePtr) ToDiskEncryptionSetTypePtrOutputWithContext(ctx context.Context) DiskEncryptionSetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskEncryptionSetTypePtrOutput)
}

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

func (DiskStorageAccountTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskStorageAccountTypes)(nil)).Elem()
}

func (e DiskStorageAccountTypes) ToDiskStorageAccountTypesOutput() DiskStorageAccountTypesOutput {
	return pulumi.ToOutput(e).(DiskStorageAccountTypesOutput)
}

func (e DiskStorageAccountTypes) ToDiskStorageAccountTypesOutputWithContext(ctx context.Context) DiskStorageAccountTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskStorageAccountTypesOutput)
}

func (e DiskStorageAccountTypes) ToDiskStorageAccountTypesPtrOutput() DiskStorageAccountTypesPtrOutput {
	return e.ToDiskStorageAccountTypesPtrOutputWithContext(context.Background())
}

func (e DiskStorageAccountTypes) ToDiskStorageAccountTypesPtrOutputWithContext(ctx context.Context) DiskStorageAccountTypesPtrOutput {
	return DiskStorageAccountTypes(e).ToDiskStorageAccountTypesOutputWithContext(ctx).ToDiskStorageAccountTypesPtrOutputWithContext(ctx)
}

func (e DiskStorageAccountTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskStorageAccountTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskStorageAccountTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskStorageAccountTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskStorageAccountTypesOutput struct{ *pulumi.OutputState }

func (DiskStorageAccountTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskStorageAccountTypes)(nil)).Elem()
}

func (o DiskStorageAccountTypesOutput) ToDiskStorageAccountTypesOutput() DiskStorageAccountTypesOutput {
	return o
}

func (o DiskStorageAccountTypesOutput) ToDiskStorageAccountTypesOutputWithContext(ctx context.Context) DiskStorageAccountTypesOutput {
	return o
}

func (o DiskStorageAccountTypesOutput) ToDiskStorageAccountTypesPtrOutput() DiskStorageAccountTypesPtrOutput {
	return o.ToDiskStorageAccountTypesPtrOutputWithContext(context.Background())
}

func (o DiskStorageAccountTypesOutput) ToDiskStorageAccountTypesPtrOutputWithContext(ctx context.Context) DiskStorageAccountTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskStorageAccountTypes) *DiskStorageAccountTypes {
		return &v
	}).(DiskStorageAccountTypesPtrOutput)
}

func (o DiskStorageAccountTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskStorageAccountTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskStorageAccountTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskStorageAccountTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskStorageAccountTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskStorageAccountTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskStorageAccountTypesPtrOutput struct{ *pulumi.OutputState }

func (DiskStorageAccountTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskStorageAccountTypes)(nil)).Elem()
}

func (o DiskStorageAccountTypesPtrOutput) ToDiskStorageAccountTypesPtrOutput() DiskStorageAccountTypesPtrOutput {
	return o
}

func (o DiskStorageAccountTypesPtrOutput) ToDiskStorageAccountTypesPtrOutputWithContext(ctx context.Context) DiskStorageAccountTypesPtrOutput {
	return o
}

func (o DiskStorageAccountTypesPtrOutput) Elem() DiskStorageAccountTypesOutput {
	return o.ApplyT(func(v *DiskStorageAccountTypes) DiskStorageAccountTypes {
		if v != nil {
			return *v
		}
		var ret DiskStorageAccountTypes
		return ret
	}).(DiskStorageAccountTypesOutput)
}

func (o DiskStorageAccountTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskStorageAccountTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskStorageAccountTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiskStorageAccountTypesInput interface {
	pulumi.Input

	ToDiskStorageAccountTypesOutput() DiskStorageAccountTypesOutput
	ToDiskStorageAccountTypesOutputWithContext(context.Context) DiskStorageAccountTypesOutput
}

var diskStorageAccountTypesPtrType = reflect.TypeOf((**DiskStorageAccountTypes)(nil)).Elem()

type DiskStorageAccountTypesPtrInput interface {
	pulumi.Input

	ToDiskStorageAccountTypesPtrOutput() DiskStorageAccountTypesPtrOutput
	ToDiskStorageAccountTypesPtrOutputWithContext(context.Context) DiskStorageAccountTypesPtrOutput
}

type diskStorageAccountTypesPtr string

func DiskStorageAccountTypesPtr(v string) DiskStorageAccountTypesPtrInput {
	return (*diskStorageAccountTypesPtr)(&v)
}

func (*diskStorageAccountTypesPtr) ElementType() reflect.Type {
	return diskStorageAccountTypesPtrType
}

func (in *diskStorageAccountTypesPtr) ToDiskStorageAccountTypesPtrOutput() DiskStorageAccountTypesPtrOutput {
	return pulumi.ToOutput(in).(DiskStorageAccountTypesPtrOutput)
}

func (in *diskStorageAccountTypesPtr) ToDiskStorageAccountTypesPtrOutputWithContext(ctx context.Context) DiskStorageAccountTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskStorageAccountTypesPtrOutput)
}

type EncryptionType string

const (
	// Disk is encrypted at rest with Platform managed key. It is the default encryption type. This is not a valid encryption type for disk encryption sets.
	EncryptionTypeEncryptionAtRestWithPlatformKey = EncryptionType("EncryptionAtRestWithPlatformKey")
	// Disk is encrypted at rest with Customer managed key that can be changed and revoked by a customer.
	EncryptionTypeEncryptionAtRestWithCustomerKey = EncryptionType("EncryptionAtRestWithCustomerKey")
	// Disk is encrypted at rest with 2 layers of encryption. One of the keys is Customer managed and the other key is Platform managed.
	EncryptionTypeEncryptionAtRestWithPlatformAndCustomerKeys = EncryptionType("EncryptionAtRestWithPlatformAndCustomerKeys")
)

func (EncryptionType) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionType)(nil)).Elem()
}

func (e EncryptionType) ToEncryptionTypeOutput() EncryptionTypeOutput {
	return pulumi.ToOutput(e).(EncryptionTypeOutput)
}

func (e EncryptionType) ToEncryptionTypeOutputWithContext(ctx context.Context) EncryptionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncryptionTypeOutput)
}

func (e EncryptionType) ToEncryptionTypePtrOutput() EncryptionTypePtrOutput {
	return e.ToEncryptionTypePtrOutputWithContext(context.Background())
}

func (e EncryptionType) ToEncryptionTypePtrOutputWithContext(ctx context.Context) EncryptionTypePtrOutput {
	return EncryptionType(e).ToEncryptionTypeOutputWithContext(ctx).ToEncryptionTypePtrOutputWithContext(ctx)
}

func (e EncryptionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncryptionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncryptionTypeOutput struct{ *pulumi.OutputState }

func (EncryptionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionType)(nil)).Elem()
}

func (o EncryptionTypeOutput) ToEncryptionTypeOutput() EncryptionTypeOutput {
	return o
}

func (o EncryptionTypeOutput) ToEncryptionTypeOutputWithContext(ctx context.Context) EncryptionTypeOutput {
	return o
}

func (o EncryptionTypeOutput) ToEncryptionTypePtrOutput() EncryptionTypePtrOutput {
	return o.ToEncryptionTypePtrOutputWithContext(context.Background())
}

func (o EncryptionTypeOutput) ToEncryptionTypePtrOutputWithContext(ctx context.Context) EncryptionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionType) *EncryptionType {
		return &v
	}).(EncryptionTypePtrOutput)
}

func (o EncryptionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncryptionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncryptionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncryptionTypePtrOutput struct{ *pulumi.OutputState }

func (EncryptionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionType)(nil)).Elem()
}

func (o EncryptionTypePtrOutput) ToEncryptionTypePtrOutput() EncryptionTypePtrOutput {
	return o
}

func (o EncryptionTypePtrOutput) ToEncryptionTypePtrOutputWithContext(ctx context.Context) EncryptionTypePtrOutput {
	return o
}

func (o EncryptionTypePtrOutput) Elem() EncryptionTypeOutput {
	return o.ApplyT(func(v *EncryptionType) EncryptionType {
		if v != nil {
			return *v
		}
		var ret EncryptionType
		return ret
	}).(EncryptionTypeOutput)
}

func (o EncryptionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncryptionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncryptionTypeInput interface {
	pulumi.Input

	ToEncryptionTypeOutput() EncryptionTypeOutput
	ToEncryptionTypeOutputWithContext(context.Context) EncryptionTypeOutput
}

var encryptionTypePtrType = reflect.TypeOf((**EncryptionType)(nil)).Elem()

type EncryptionTypePtrInput interface {
	pulumi.Input

	ToEncryptionTypePtrOutput() EncryptionTypePtrOutput
	ToEncryptionTypePtrOutputWithContext(context.Context) EncryptionTypePtrOutput
}

type encryptionTypePtr string

func EncryptionTypePtr(v string) EncryptionTypePtrInput {
	return (*encryptionTypePtr)(&v)
}

func (*encryptionTypePtr) ElementType() reflect.Type {
	return encryptionTypePtrType
}

func (in *encryptionTypePtr) ToEncryptionTypePtrOutput() EncryptionTypePtrOutput {
	return pulumi.ToOutput(in).(EncryptionTypePtrOutput)
}

func (in *encryptionTypePtr) ToEncryptionTypePtrOutputWithContext(ctx context.Context) EncryptionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncryptionTypePtrOutput)
}

type ExtendedLocationTypes string

const (
	ExtendedLocationTypesEdgeZone = ExtendedLocationTypes("EdgeZone")
)

func (ExtendedLocationTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationTypes)(nil)).Elem()
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput {
	return pulumi.ToOutput(e).(ExtendedLocationTypesOutput)
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesOutputWithContext(ctx context.Context) ExtendedLocationTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExtendedLocationTypesOutput)
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return e.ToExtendedLocationTypesPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return ExtendedLocationTypes(e).ToExtendedLocationTypesOutputWithContext(ctx).ToExtendedLocationTypesPtrOutputWithContext(ctx)
}

func (e ExtendedLocationTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExtendedLocationTypesOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationTypes)(nil)).Elem()
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput {
	return o
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesOutputWithContext(ctx context.Context) ExtendedLocationTypesOutput {
	return o
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return o.ToExtendedLocationTypesPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationTypes) *ExtendedLocationTypes {
		return &v
	}).(ExtendedLocationTypesPtrOutput)
}

func (o ExtendedLocationTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExtendedLocationTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationTypesPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationTypes)(nil)).Elem()
}

func (o ExtendedLocationTypesPtrOutput) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return o
}

func (o ExtendedLocationTypesPtrOutput) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return o
}

func (o ExtendedLocationTypesPtrOutput) Elem() ExtendedLocationTypesOutput {
	return o.ApplyT(func(v *ExtendedLocationTypes) ExtendedLocationTypes {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationTypes
		return ret
	}).(ExtendedLocationTypesOutput)
}

func (o ExtendedLocationTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExtendedLocationTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExtendedLocationTypesInput interface {
	pulumi.Input

	ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput
	ToExtendedLocationTypesOutputWithContext(context.Context) ExtendedLocationTypesOutput
}

var extendedLocationTypesPtrType = reflect.TypeOf((**ExtendedLocationTypes)(nil)).Elem()

type ExtendedLocationTypesPtrInput interface {
	pulumi.Input

	ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput
	ToExtendedLocationTypesPtrOutputWithContext(context.Context) ExtendedLocationTypesPtrOutput
}

type extendedLocationTypesPtr string

func ExtendedLocationTypesPtr(v string) ExtendedLocationTypesPtrInput {
	return (*extendedLocationTypesPtr)(&v)
}

func (*extendedLocationTypesPtr) ElementType() reflect.Type {
	return extendedLocationTypesPtrType
}

func (in *extendedLocationTypesPtr) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return pulumi.ToOutput(in).(ExtendedLocationTypesPtrOutput)
}

func (in *extendedLocationTypesPtr) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExtendedLocationTypesPtrOutput)
}

type GallerySharingPermissionTypes string

const (
	GallerySharingPermissionTypesPrivate = GallerySharingPermissionTypes("Private")
	GallerySharingPermissionTypesGroups  = GallerySharingPermissionTypes("Groups")
)

func (GallerySharingPermissionTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*GallerySharingPermissionTypes)(nil)).Elem()
}

func (e GallerySharingPermissionTypes) ToGallerySharingPermissionTypesOutput() GallerySharingPermissionTypesOutput {
	return pulumi.ToOutput(e).(GallerySharingPermissionTypesOutput)
}

func (e GallerySharingPermissionTypes) ToGallerySharingPermissionTypesOutputWithContext(ctx context.Context) GallerySharingPermissionTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GallerySharingPermissionTypesOutput)
}

func (e GallerySharingPermissionTypes) ToGallerySharingPermissionTypesPtrOutput() GallerySharingPermissionTypesPtrOutput {
	return e.ToGallerySharingPermissionTypesPtrOutputWithContext(context.Background())
}

func (e GallerySharingPermissionTypes) ToGallerySharingPermissionTypesPtrOutputWithContext(ctx context.Context) GallerySharingPermissionTypesPtrOutput {
	return GallerySharingPermissionTypes(e).ToGallerySharingPermissionTypesOutputWithContext(ctx).ToGallerySharingPermissionTypesPtrOutputWithContext(ctx)
}

func (e GallerySharingPermissionTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GallerySharingPermissionTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GallerySharingPermissionTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GallerySharingPermissionTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GallerySharingPermissionTypesOutput struct{ *pulumi.OutputState }

func (GallerySharingPermissionTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GallerySharingPermissionTypes)(nil)).Elem()
}

func (o GallerySharingPermissionTypesOutput) ToGallerySharingPermissionTypesOutput() GallerySharingPermissionTypesOutput {
	return o
}

func (o GallerySharingPermissionTypesOutput) ToGallerySharingPermissionTypesOutputWithContext(ctx context.Context) GallerySharingPermissionTypesOutput {
	return o
}

func (o GallerySharingPermissionTypesOutput) ToGallerySharingPermissionTypesPtrOutput() GallerySharingPermissionTypesPtrOutput {
	return o.ToGallerySharingPermissionTypesPtrOutputWithContext(context.Background())
}

func (o GallerySharingPermissionTypesOutput) ToGallerySharingPermissionTypesPtrOutputWithContext(ctx context.Context) GallerySharingPermissionTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GallerySharingPermissionTypes) *GallerySharingPermissionTypes {
		return &v
	}).(GallerySharingPermissionTypesPtrOutput)
}

func (o GallerySharingPermissionTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GallerySharingPermissionTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GallerySharingPermissionTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GallerySharingPermissionTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GallerySharingPermissionTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GallerySharingPermissionTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GallerySharingPermissionTypesPtrOutput struct{ *pulumi.OutputState }

func (GallerySharingPermissionTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GallerySharingPermissionTypes)(nil)).Elem()
}

func (o GallerySharingPermissionTypesPtrOutput) ToGallerySharingPermissionTypesPtrOutput() GallerySharingPermissionTypesPtrOutput {
	return o
}

func (o GallerySharingPermissionTypesPtrOutput) ToGallerySharingPermissionTypesPtrOutputWithContext(ctx context.Context) GallerySharingPermissionTypesPtrOutput {
	return o
}

func (o GallerySharingPermissionTypesPtrOutput) Elem() GallerySharingPermissionTypesOutput {
	return o.ApplyT(func(v *GallerySharingPermissionTypes) GallerySharingPermissionTypes {
		if v != nil {
			return *v
		}
		var ret GallerySharingPermissionTypes
		return ret
	}).(GallerySharingPermissionTypesOutput)
}

func (o GallerySharingPermissionTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GallerySharingPermissionTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GallerySharingPermissionTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GallerySharingPermissionTypesInput interface {
	pulumi.Input

	ToGallerySharingPermissionTypesOutput() GallerySharingPermissionTypesOutput
	ToGallerySharingPermissionTypesOutputWithContext(context.Context) GallerySharingPermissionTypesOutput
}

var gallerySharingPermissionTypesPtrType = reflect.TypeOf((**GallerySharingPermissionTypes)(nil)).Elem()

type GallerySharingPermissionTypesPtrInput interface {
	pulumi.Input

	ToGallerySharingPermissionTypesPtrOutput() GallerySharingPermissionTypesPtrOutput
	ToGallerySharingPermissionTypesPtrOutputWithContext(context.Context) GallerySharingPermissionTypesPtrOutput
}

type gallerySharingPermissionTypesPtr string

func GallerySharingPermissionTypesPtr(v string) GallerySharingPermissionTypesPtrInput {
	return (*gallerySharingPermissionTypesPtr)(&v)
}

func (*gallerySharingPermissionTypesPtr) ElementType() reflect.Type {
	return gallerySharingPermissionTypesPtrType
}

func (in *gallerySharingPermissionTypesPtr) ToGallerySharingPermissionTypesPtrOutput() GallerySharingPermissionTypesPtrOutput {
	return pulumi.ToOutput(in).(GallerySharingPermissionTypesPtrOutput)
}

func (in *gallerySharingPermissionTypesPtr) ToGallerySharingPermissionTypesPtrOutputWithContext(ctx context.Context) GallerySharingPermissionTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GallerySharingPermissionTypesPtrOutput)
}

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

func (HyperVGeneration) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVGeneration)(nil)).Elem()
}

func (e HyperVGeneration) ToHyperVGenerationOutput() HyperVGenerationOutput {
	return pulumi.ToOutput(e).(HyperVGenerationOutput)
}

func (e HyperVGeneration) ToHyperVGenerationOutputWithContext(ctx context.Context) HyperVGenerationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HyperVGenerationOutput)
}

func (e HyperVGeneration) ToHyperVGenerationPtrOutput() HyperVGenerationPtrOutput {
	return e.ToHyperVGenerationPtrOutputWithContext(context.Background())
}

func (e HyperVGeneration) ToHyperVGenerationPtrOutputWithContext(ctx context.Context) HyperVGenerationPtrOutput {
	return HyperVGeneration(e).ToHyperVGenerationOutputWithContext(ctx).ToHyperVGenerationPtrOutputWithContext(ctx)
}

func (e HyperVGeneration) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HyperVGeneration) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HyperVGeneration) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HyperVGeneration) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HyperVGenerationOutput struct{ *pulumi.OutputState }

func (HyperVGenerationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVGeneration)(nil)).Elem()
}

func (o HyperVGenerationOutput) ToHyperVGenerationOutput() HyperVGenerationOutput {
	return o
}

func (o HyperVGenerationOutput) ToHyperVGenerationOutputWithContext(ctx context.Context) HyperVGenerationOutput {
	return o
}

func (o HyperVGenerationOutput) ToHyperVGenerationPtrOutput() HyperVGenerationPtrOutput {
	return o.ToHyperVGenerationPtrOutputWithContext(context.Background())
}

func (o HyperVGenerationOutput) ToHyperVGenerationPtrOutputWithContext(ctx context.Context) HyperVGenerationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HyperVGeneration) *HyperVGeneration {
		return &v
	}).(HyperVGenerationPtrOutput)
}

func (o HyperVGenerationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HyperVGenerationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HyperVGeneration) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HyperVGenerationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HyperVGenerationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HyperVGeneration) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HyperVGenerationPtrOutput struct{ *pulumi.OutputState }

func (HyperVGenerationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HyperVGeneration)(nil)).Elem()
}

func (o HyperVGenerationPtrOutput) ToHyperVGenerationPtrOutput() HyperVGenerationPtrOutput {
	return o
}

func (o HyperVGenerationPtrOutput) ToHyperVGenerationPtrOutputWithContext(ctx context.Context) HyperVGenerationPtrOutput {
	return o
}

func (o HyperVGenerationPtrOutput) Elem() HyperVGenerationOutput {
	return o.ApplyT(func(v *HyperVGeneration) HyperVGeneration {
		if v != nil {
			return *v
		}
		var ret HyperVGeneration
		return ret
	}).(HyperVGenerationOutput)
}

func (o HyperVGenerationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HyperVGenerationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HyperVGeneration) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HyperVGenerationInput interface {
	pulumi.Input

	ToHyperVGenerationOutput() HyperVGenerationOutput
	ToHyperVGenerationOutputWithContext(context.Context) HyperVGenerationOutput
}

var hyperVGenerationPtrType = reflect.TypeOf((**HyperVGeneration)(nil)).Elem()

type HyperVGenerationPtrInput interface {
	pulumi.Input

	ToHyperVGenerationPtrOutput() HyperVGenerationPtrOutput
	ToHyperVGenerationPtrOutputWithContext(context.Context) HyperVGenerationPtrOutput
}

type hyperVGenerationPtr string

func HyperVGenerationPtr(v string) HyperVGenerationPtrInput {
	return (*hyperVGenerationPtr)(&v)
}

func (*hyperVGenerationPtr) ElementType() reflect.Type {
	return hyperVGenerationPtrType
}

func (in *hyperVGenerationPtr) ToHyperVGenerationPtrOutput() HyperVGenerationPtrOutput {
	return pulumi.ToOutput(in).(HyperVGenerationPtrOutput)
}

func (in *hyperVGenerationPtr) ToHyperVGenerationPtrOutputWithContext(ctx context.Context) HyperVGenerationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HyperVGenerationPtrOutput)
}

type NetworkAccessPolicy string

const (
	// The disk can be exported or uploaded to from any network.
	NetworkAccessPolicyAllowAll = NetworkAccessPolicy("AllowAll")
	// The disk can be exported or uploaded to using a DiskAccess resource's private endpoints.
	NetworkAccessPolicyAllowPrivate = NetworkAccessPolicy("AllowPrivate")
	// The disk cannot be exported.
	NetworkAccessPolicyDenyAll = NetworkAccessPolicy("DenyAll")
)

func (NetworkAccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessPolicy)(nil)).Elem()
}

func (e NetworkAccessPolicy) ToNetworkAccessPolicyOutput() NetworkAccessPolicyOutput {
	return pulumi.ToOutput(e).(NetworkAccessPolicyOutput)
}

func (e NetworkAccessPolicy) ToNetworkAccessPolicyOutputWithContext(ctx context.Context) NetworkAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkAccessPolicyOutput)
}

func (e NetworkAccessPolicy) ToNetworkAccessPolicyPtrOutput() NetworkAccessPolicyPtrOutput {
	return e.ToNetworkAccessPolicyPtrOutputWithContext(context.Background())
}

func (e NetworkAccessPolicy) ToNetworkAccessPolicyPtrOutputWithContext(ctx context.Context) NetworkAccessPolicyPtrOutput {
	return NetworkAccessPolicy(e).ToNetworkAccessPolicyOutputWithContext(ctx).ToNetworkAccessPolicyPtrOutputWithContext(ctx)
}

func (e NetworkAccessPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkAccessPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkAccessPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkAccessPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkAccessPolicyOutput struct{ *pulumi.OutputState }

func (NetworkAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessPolicy)(nil)).Elem()
}

func (o NetworkAccessPolicyOutput) ToNetworkAccessPolicyOutput() NetworkAccessPolicyOutput {
	return o
}

func (o NetworkAccessPolicyOutput) ToNetworkAccessPolicyOutputWithContext(ctx context.Context) NetworkAccessPolicyOutput {
	return o
}

func (o NetworkAccessPolicyOutput) ToNetworkAccessPolicyPtrOutput() NetworkAccessPolicyPtrOutput {
	return o.ToNetworkAccessPolicyPtrOutputWithContext(context.Background())
}

func (o NetworkAccessPolicyOutput) ToNetworkAccessPolicyPtrOutputWithContext(ctx context.Context) NetworkAccessPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkAccessPolicy) *NetworkAccessPolicy {
		return &v
	}).(NetworkAccessPolicyPtrOutput)
}

func (o NetworkAccessPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkAccessPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkAccessPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkAccessPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkAccessPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkAccessPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkAccessPolicyPtrOutput struct{ *pulumi.OutputState }

func (NetworkAccessPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAccessPolicy)(nil)).Elem()
}

func (o NetworkAccessPolicyPtrOutput) ToNetworkAccessPolicyPtrOutput() NetworkAccessPolicyPtrOutput {
	return o
}

func (o NetworkAccessPolicyPtrOutput) ToNetworkAccessPolicyPtrOutputWithContext(ctx context.Context) NetworkAccessPolicyPtrOutput {
	return o
}

func (o NetworkAccessPolicyPtrOutput) Elem() NetworkAccessPolicyOutput {
	return o.ApplyT(func(v *NetworkAccessPolicy) NetworkAccessPolicy {
		if v != nil {
			return *v
		}
		var ret NetworkAccessPolicy
		return ret
	}).(NetworkAccessPolicyOutput)
}

func (o NetworkAccessPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkAccessPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkAccessPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkAccessPolicyInput interface {
	pulumi.Input

	ToNetworkAccessPolicyOutput() NetworkAccessPolicyOutput
	ToNetworkAccessPolicyOutputWithContext(context.Context) NetworkAccessPolicyOutput
}

var networkAccessPolicyPtrType = reflect.TypeOf((**NetworkAccessPolicy)(nil)).Elem()

type NetworkAccessPolicyPtrInput interface {
	pulumi.Input

	ToNetworkAccessPolicyPtrOutput() NetworkAccessPolicyPtrOutput
	ToNetworkAccessPolicyPtrOutputWithContext(context.Context) NetworkAccessPolicyPtrOutput
}

type networkAccessPolicyPtr string

func NetworkAccessPolicyPtr(v string) NetworkAccessPolicyPtrInput {
	return (*networkAccessPolicyPtr)(&v)
}

func (*networkAccessPolicyPtr) ElementType() reflect.Type {
	return networkAccessPolicyPtrType
}

func (in *networkAccessPolicyPtr) ToNetworkAccessPolicyPtrOutput() NetworkAccessPolicyPtrOutput {
	return pulumi.ToOutput(in).(NetworkAccessPolicyPtrOutput)
}

func (in *networkAccessPolicyPtr) ToNetworkAccessPolicyPtrOutputWithContext(ctx context.Context) NetworkAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkAccessPolicyPtrOutput)
}

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

func (PrivateEndpointServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return e.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return PrivateEndpointServiceConnectionStatus(e).ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx).ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateEndpointServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointServiceConnectionStatus) *PrivateEndpointServiceConnectionStatus {
		return &v
	}).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) Elem() PrivateEndpointServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateEndpointServiceConnectionStatus) PrivateEndpointServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointServiceConnectionStatus
		return ret
	}).(PrivateEndpointServiceConnectionStatusOutput)
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateEndpointServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateEndpointServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput
	ToPrivateEndpointServiceConnectionStatusOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusOutput
}

var privateEndpointServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()

type PrivateEndpointServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput
	ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusPtrOutput
}

type privateEndpointServiceConnectionStatusPtr string

func PrivateEndpointServiceConnectionStatusPtr(v string) PrivateEndpointServiceConnectionStatusPtrInput {
	return (*privateEndpointServiceConnectionStatusPtr)(&v)
}

func (*privateEndpointServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateEndpointServiceConnectionStatusPtrType
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

type SnapshotStorageAccountTypes string

const (
	// Standard HDD locally redundant storage
	SnapshotStorageAccountTypes_Standard_LRS = SnapshotStorageAccountTypes("Standard_LRS")
	// Premium SSD locally redundant storage
	SnapshotStorageAccountTypes_Premium_LRS = SnapshotStorageAccountTypes("Premium_LRS")
	// Standard zone redundant storage
	SnapshotStorageAccountTypes_Standard_ZRS = SnapshotStorageAccountTypes("Standard_ZRS")
)

func (SnapshotStorageAccountTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotStorageAccountTypes)(nil)).Elem()
}

func (e SnapshotStorageAccountTypes) ToSnapshotStorageAccountTypesOutput() SnapshotStorageAccountTypesOutput {
	return pulumi.ToOutput(e).(SnapshotStorageAccountTypesOutput)
}

func (e SnapshotStorageAccountTypes) ToSnapshotStorageAccountTypesOutputWithContext(ctx context.Context) SnapshotStorageAccountTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SnapshotStorageAccountTypesOutput)
}

func (e SnapshotStorageAccountTypes) ToSnapshotStorageAccountTypesPtrOutput() SnapshotStorageAccountTypesPtrOutput {
	return e.ToSnapshotStorageAccountTypesPtrOutputWithContext(context.Background())
}

func (e SnapshotStorageAccountTypes) ToSnapshotStorageAccountTypesPtrOutputWithContext(ctx context.Context) SnapshotStorageAccountTypesPtrOutput {
	return SnapshotStorageAccountTypes(e).ToSnapshotStorageAccountTypesOutputWithContext(ctx).ToSnapshotStorageAccountTypesPtrOutputWithContext(ctx)
}

func (e SnapshotStorageAccountTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SnapshotStorageAccountTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SnapshotStorageAccountTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SnapshotStorageAccountTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SnapshotStorageAccountTypesOutput struct{ *pulumi.OutputState }

func (SnapshotStorageAccountTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotStorageAccountTypes)(nil)).Elem()
}

func (o SnapshotStorageAccountTypesOutput) ToSnapshotStorageAccountTypesOutput() SnapshotStorageAccountTypesOutput {
	return o
}

func (o SnapshotStorageAccountTypesOutput) ToSnapshotStorageAccountTypesOutputWithContext(ctx context.Context) SnapshotStorageAccountTypesOutput {
	return o
}

func (o SnapshotStorageAccountTypesOutput) ToSnapshotStorageAccountTypesPtrOutput() SnapshotStorageAccountTypesPtrOutput {
	return o.ToSnapshotStorageAccountTypesPtrOutputWithContext(context.Background())
}

func (o SnapshotStorageAccountTypesOutput) ToSnapshotStorageAccountTypesPtrOutputWithContext(ctx context.Context) SnapshotStorageAccountTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SnapshotStorageAccountTypes) *SnapshotStorageAccountTypes {
		return &v
	}).(SnapshotStorageAccountTypesPtrOutput)
}

func (o SnapshotStorageAccountTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SnapshotStorageAccountTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SnapshotStorageAccountTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SnapshotStorageAccountTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SnapshotStorageAccountTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SnapshotStorageAccountTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SnapshotStorageAccountTypesPtrOutput struct{ *pulumi.OutputState }

func (SnapshotStorageAccountTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotStorageAccountTypes)(nil)).Elem()
}

func (o SnapshotStorageAccountTypesPtrOutput) ToSnapshotStorageAccountTypesPtrOutput() SnapshotStorageAccountTypesPtrOutput {
	return o
}

func (o SnapshotStorageAccountTypesPtrOutput) ToSnapshotStorageAccountTypesPtrOutputWithContext(ctx context.Context) SnapshotStorageAccountTypesPtrOutput {
	return o
}

func (o SnapshotStorageAccountTypesPtrOutput) Elem() SnapshotStorageAccountTypesOutput {
	return o.ApplyT(func(v *SnapshotStorageAccountTypes) SnapshotStorageAccountTypes {
		if v != nil {
			return *v
		}
		var ret SnapshotStorageAccountTypes
		return ret
	}).(SnapshotStorageAccountTypesOutput)
}

func (o SnapshotStorageAccountTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SnapshotStorageAccountTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SnapshotStorageAccountTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SnapshotStorageAccountTypesInput interface {
	pulumi.Input

	ToSnapshotStorageAccountTypesOutput() SnapshotStorageAccountTypesOutput
	ToSnapshotStorageAccountTypesOutputWithContext(context.Context) SnapshotStorageAccountTypesOutput
}

var snapshotStorageAccountTypesPtrType = reflect.TypeOf((**SnapshotStorageAccountTypes)(nil)).Elem()

type SnapshotStorageAccountTypesPtrInput interface {
	pulumi.Input

	ToSnapshotStorageAccountTypesPtrOutput() SnapshotStorageAccountTypesPtrOutput
	ToSnapshotStorageAccountTypesPtrOutputWithContext(context.Context) SnapshotStorageAccountTypesPtrOutput
}

type snapshotStorageAccountTypesPtr string

func SnapshotStorageAccountTypesPtr(v string) SnapshotStorageAccountTypesPtrInput {
	return (*snapshotStorageAccountTypesPtr)(&v)
}

func (*snapshotStorageAccountTypesPtr) ElementType() reflect.Type {
	return snapshotStorageAccountTypesPtrType
}

func (in *snapshotStorageAccountTypesPtr) ToSnapshotStorageAccountTypesPtrOutput() SnapshotStorageAccountTypesPtrOutput {
	return pulumi.ToOutput(in).(SnapshotStorageAccountTypesPtrOutput)
}

func (in *snapshotStorageAccountTypesPtr) ToSnapshotStorageAccountTypesPtrOutputWithContext(ctx context.Context) SnapshotStorageAccountTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SnapshotStorageAccountTypesPtrOutput)
}

type StorageAccountType string

const (
	StorageAccountType_Standard_LRS = StorageAccountType("Standard_LRS")
	StorageAccountType_Standard_ZRS = StorageAccountType("Standard_ZRS")
	StorageAccountType_Premium_LRS  = StorageAccountType("Premium_LRS")
)

func (StorageAccountType) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountType)(nil)).Elem()
}

func (e StorageAccountType) ToStorageAccountTypeOutput() StorageAccountTypeOutput {
	return pulumi.ToOutput(e).(StorageAccountTypeOutput)
}

func (e StorageAccountType) ToStorageAccountTypeOutputWithContext(ctx context.Context) StorageAccountTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageAccountTypeOutput)
}

func (e StorageAccountType) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return e.ToStorageAccountTypePtrOutputWithContext(context.Background())
}

func (e StorageAccountType) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return StorageAccountType(e).ToStorageAccountTypeOutputWithContext(ctx).ToStorageAccountTypePtrOutputWithContext(ctx)
}

func (e StorageAccountType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageAccountType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageAccountTypeOutput struct{ *pulumi.OutputState }

func (StorageAccountTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountType)(nil)).Elem()
}

func (o StorageAccountTypeOutput) ToStorageAccountTypeOutput() StorageAccountTypeOutput {
	return o
}

func (o StorageAccountTypeOutput) ToStorageAccountTypeOutputWithContext(ctx context.Context) StorageAccountTypeOutput {
	return o
}

func (o StorageAccountTypeOutput) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return o.ToStorageAccountTypePtrOutputWithContext(context.Background())
}

func (o StorageAccountTypeOutput) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountType) *StorageAccountType {
		return &v
	}).(StorageAccountTypePtrOutput)
}

func (o StorageAccountTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageAccountTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAccountType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageAccountTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAccountTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAccountType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageAccountTypePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountType)(nil)).Elem()
}

func (o StorageAccountTypePtrOutput) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return o
}

func (o StorageAccountTypePtrOutput) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return o
}

func (o StorageAccountTypePtrOutput) Elem() StorageAccountTypeOutput {
	return o.ApplyT(func(v *StorageAccountType) StorageAccountType {
		if v != nil {
			return *v
		}
		var ret StorageAccountType
		return ret
	}).(StorageAccountTypeOutput)
}

func (o StorageAccountTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAccountTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageAccountType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageAccountTypeInput interface {
	pulumi.Input

	ToStorageAccountTypeOutput() StorageAccountTypeOutput
	ToStorageAccountTypeOutputWithContext(context.Context) StorageAccountTypeOutput
}

var storageAccountTypePtrType = reflect.TypeOf((**StorageAccountType)(nil)).Elem()

type StorageAccountTypePtrInput interface {
	pulumi.Input

	ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput
	ToStorageAccountTypePtrOutputWithContext(context.Context) StorageAccountTypePtrOutput
}

type storageAccountTypePtr string

func StorageAccountTypePtr(v string) StorageAccountTypePtrInput {
	return (*storageAccountTypePtr)(&v)
}

func (*storageAccountTypePtr) ElementType() reflect.Type {
	return storageAccountTypePtrType
}

func (in *storageAccountTypePtr) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return pulumi.ToOutput(in).(StorageAccountTypePtrOutput)
}

func (in *storageAccountTypePtr) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageAccountTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskCreateOptionOutput{})
	pulumi.RegisterOutputType(DiskCreateOptionPtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetIdentityTypeOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetTypeOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetTypePtrOutput{})
	pulumi.RegisterOutputType(DiskStorageAccountTypesOutput{})
	pulumi.RegisterOutputType(DiskStorageAccountTypesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionTypeOutput{})
	pulumi.RegisterOutputType(EncryptionTypePtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypesOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypesPtrOutput{})
	pulumi.RegisterOutputType(GallerySharingPermissionTypesOutput{})
	pulumi.RegisterOutputType(GallerySharingPermissionTypesPtrOutput{})
	pulumi.RegisterOutputType(HostCachingOutput{})
	pulumi.RegisterOutputType(HostCachingPtrOutput{})
	pulumi.RegisterOutputType(HyperVGenerationOutput{})
	pulumi.RegisterOutputType(HyperVGenerationPtrOutput{})
	pulumi.RegisterOutputType(NetworkAccessPolicyOutput{})
	pulumi.RegisterOutputType(NetworkAccessPolicyPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemStateTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemStateTypesPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(SnapshotStorageAccountTypesOutput{})
	pulumi.RegisterOutputType(SnapshotStorageAccountTypesPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountTypeOutput{})
	pulumi.RegisterOutputType(StorageAccountTypePtrOutput{})
}
