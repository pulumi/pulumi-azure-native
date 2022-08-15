


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationSharingPolicy string

const (
	ApplicationSharingPolicyPersonal = ApplicationSharingPolicy("Personal")
	ApplicationSharingPolicyShared   = ApplicationSharingPolicy("Shared")
)

type ComputeEnvironmentType string

const (
	ComputeEnvironmentTypeACI = ComputeEnvironmentType("ACI")
	ComputeEnvironmentTypeAKS = ComputeEnvironmentType("AKS")
)

type ComputeInstanceAuthorizationType string

const (
	ComputeInstanceAuthorizationTypePersonal = ComputeInstanceAuthorizationType("personal")
)

type ComputeType string

const (
	ComputeTypeAKS               = ComputeType("AKS")
	ComputeTypeAmlCompute        = ComputeType("AmlCompute")
	ComputeTypeComputeInstance   = ComputeType("ComputeInstance")
	ComputeTypeDataFactory       = ComputeType("DataFactory")
	ComputeTypeVirtualMachine    = ComputeType("VirtualMachine")
	ComputeTypeHDInsight         = ComputeType("HDInsight")
	ComputeTypeDatabricks        = ComputeType("Databricks")
	ComputeTypeDataLakeAnalytics = ComputeType("DataLakeAnalytics")
)

type EncryptionStatus string

const (
	EncryptionStatusEnabled  = EncryptionStatus("Enabled")
	EncryptionStatusDisabled = EncryptionStatus("Disabled")
)

type ImageAnnotationType string

const (
	ImageAnnotationTypeClassification       = ImageAnnotationType("Classification")
	ImageAnnotationTypeBoundingBox          = ImageAnnotationType("BoundingBox")
	ImageAnnotationTypeInstanceSegmentation = ImageAnnotationType("InstanceSegmentation")
)

type LinkedServiceLinkType string

const (
	LinkedServiceLinkTypeSynapse = LinkedServiceLinkType("Synapse")
)

func (LinkedServiceLinkType) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceLinkType)(nil)).Elem()
}

func (e LinkedServiceLinkType) ToLinkedServiceLinkTypeOutput() LinkedServiceLinkTypeOutput {
	return pulumi.ToOutput(e).(LinkedServiceLinkTypeOutput)
}

func (e LinkedServiceLinkType) ToLinkedServiceLinkTypeOutputWithContext(ctx context.Context) LinkedServiceLinkTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LinkedServiceLinkTypeOutput)
}

func (e LinkedServiceLinkType) ToLinkedServiceLinkTypePtrOutput() LinkedServiceLinkTypePtrOutput {
	return e.ToLinkedServiceLinkTypePtrOutputWithContext(context.Background())
}

func (e LinkedServiceLinkType) ToLinkedServiceLinkTypePtrOutputWithContext(ctx context.Context) LinkedServiceLinkTypePtrOutput {
	return LinkedServiceLinkType(e).ToLinkedServiceLinkTypeOutputWithContext(ctx).ToLinkedServiceLinkTypePtrOutputWithContext(ctx)
}

func (e LinkedServiceLinkType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LinkedServiceLinkType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LinkedServiceLinkType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LinkedServiceLinkType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LinkedServiceLinkTypeOutput struct{ *pulumi.OutputState }

func (LinkedServiceLinkTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceLinkType)(nil)).Elem()
}

func (o LinkedServiceLinkTypeOutput) ToLinkedServiceLinkTypeOutput() LinkedServiceLinkTypeOutput {
	return o
}

func (o LinkedServiceLinkTypeOutput) ToLinkedServiceLinkTypeOutputWithContext(ctx context.Context) LinkedServiceLinkTypeOutput {
	return o
}

func (o LinkedServiceLinkTypeOutput) ToLinkedServiceLinkTypePtrOutput() LinkedServiceLinkTypePtrOutput {
	return o.ToLinkedServiceLinkTypePtrOutputWithContext(context.Background())
}

func (o LinkedServiceLinkTypeOutput) ToLinkedServiceLinkTypePtrOutputWithContext(ctx context.Context) LinkedServiceLinkTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinkedServiceLinkType) *LinkedServiceLinkType {
		return &v
	}).(LinkedServiceLinkTypePtrOutput)
}

func (o LinkedServiceLinkTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LinkedServiceLinkTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LinkedServiceLinkType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LinkedServiceLinkTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LinkedServiceLinkTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LinkedServiceLinkType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LinkedServiceLinkTypePtrOutput struct{ *pulumi.OutputState }

func (LinkedServiceLinkTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceLinkType)(nil)).Elem()
}

func (o LinkedServiceLinkTypePtrOutput) ToLinkedServiceLinkTypePtrOutput() LinkedServiceLinkTypePtrOutput {
	return o
}

func (o LinkedServiceLinkTypePtrOutput) ToLinkedServiceLinkTypePtrOutputWithContext(ctx context.Context) LinkedServiceLinkTypePtrOutput {
	return o
}

func (o LinkedServiceLinkTypePtrOutput) Elem() LinkedServiceLinkTypeOutput {
	return o.ApplyT(func(v *LinkedServiceLinkType) LinkedServiceLinkType {
		if v != nil {
			return *v
		}
		var ret LinkedServiceLinkType
		return ret
	}).(LinkedServiceLinkTypeOutput)
}

func (o LinkedServiceLinkTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LinkedServiceLinkTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LinkedServiceLinkType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LinkedServiceLinkTypeInput interface {
	pulumi.Input

	ToLinkedServiceLinkTypeOutput() LinkedServiceLinkTypeOutput
	ToLinkedServiceLinkTypeOutputWithContext(context.Context) LinkedServiceLinkTypeOutput
}

var linkedServiceLinkTypePtrType = reflect.TypeOf((**LinkedServiceLinkType)(nil)).Elem()

type LinkedServiceLinkTypePtrInput interface {
	pulumi.Input

	ToLinkedServiceLinkTypePtrOutput() LinkedServiceLinkTypePtrOutput
	ToLinkedServiceLinkTypePtrOutputWithContext(context.Context) LinkedServiceLinkTypePtrOutput
}

type linkedServiceLinkTypePtr string

func LinkedServiceLinkTypePtr(v string) LinkedServiceLinkTypePtrInput {
	return (*linkedServiceLinkTypePtr)(&v)
}

func (*linkedServiceLinkTypePtr) ElementType() reflect.Type {
	return linkedServiceLinkTypePtrType
}

func (in *linkedServiceLinkTypePtr) ToLinkedServiceLinkTypePtrOutput() LinkedServiceLinkTypePtrOutput {
	return pulumi.ToOutput(in).(LinkedServiceLinkTypePtrOutput)
}

func (in *linkedServiceLinkTypePtr) ToLinkedServiceLinkTypePtrOutputWithContext(ctx context.Context) LinkedServiceLinkTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LinkedServiceLinkTypePtrOutput)
}

type MediaType string

const (
	MediaTypeImage = MediaType("Image")
	MediaTypeText  = MediaType("Text")
)

type OsType string

const (
	OsTypeLinux   = OsType("Linux")
	OsTypeWindows = OsType("Windows")
)

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending      = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved     = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected     = PrivateEndpointServiceConnectionStatus("Rejected")
	PrivateEndpointServiceConnectionStatusDisconnected = PrivateEndpointServiceConnectionStatus("Disconnected")
	PrivateEndpointServiceConnectionStatusTimeout      = PrivateEndpointServiceConnectionStatus("Timeout")
)

type RemoteLoginPortPublicAccess string

const (
	RemoteLoginPortPublicAccessEnabled      = RemoteLoginPortPublicAccess("Enabled")
	RemoteLoginPortPublicAccessDisabled     = RemoteLoginPortPublicAccess("Disabled")
	RemoteLoginPortPublicAccessNotSpecified = RemoteLoginPortPublicAccess("NotSpecified")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned,UserAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
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

type SshPublicAccess string

const (
	SshPublicAccessEnabled  = SshPublicAccess("Enabled")
	SshPublicAccessDisabled = SshPublicAccess("Disabled")
)

type VariantType string

const (
	VariantTypeControl   = VariantType("Control")
	VariantTypeTreatment = VariantType("Treatment")
)

type VmPriority string

const (
	VmPriorityDedicated   = VmPriority("Dedicated")
	VmPriorityLowPriority = VmPriority("LowPriority")
)

func init() {
	pulumi.RegisterOutputType(LinkedServiceLinkTypeOutput{})
	pulumi.RegisterOutputType(LinkedServiceLinkTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
