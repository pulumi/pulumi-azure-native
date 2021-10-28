


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudName string

const (
	CloudNameAzure = CloudName("Azure")
	CloudNameAWS   = CloudName("AWS")
	CloudNameGCP   = CloudName("GCP")
)

func (CloudName) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudName)(nil)).Elem()
}

func (e CloudName) ToCloudNameOutput() CloudNameOutput {
	return pulumi.ToOutput(e).(CloudNameOutput)
}

func (e CloudName) ToCloudNameOutputWithContext(ctx context.Context) CloudNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CloudNameOutput)
}

func (e CloudName) ToCloudNamePtrOutput() CloudNamePtrOutput {
	return e.ToCloudNamePtrOutputWithContext(context.Background())
}

func (e CloudName) ToCloudNamePtrOutputWithContext(ctx context.Context) CloudNamePtrOutput {
	return CloudName(e).ToCloudNameOutputWithContext(ctx).ToCloudNamePtrOutputWithContext(ctx)
}

func (e CloudName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CloudName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CloudName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CloudName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CloudNameOutput struct{ *pulumi.OutputState }

func (CloudNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudName)(nil)).Elem()
}

func (o CloudNameOutput) ToCloudNameOutput() CloudNameOutput {
	return o
}

func (o CloudNameOutput) ToCloudNameOutputWithContext(ctx context.Context) CloudNameOutput {
	return o
}

func (o CloudNameOutput) ToCloudNamePtrOutput() CloudNamePtrOutput {
	return o.ToCloudNamePtrOutputWithContext(context.Background())
}

func (o CloudNameOutput) ToCloudNamePtrOutputWithContext(ctx context.Context) CloudNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudName) *CloudName {
		return &v
	}).(CloudNamePtrOutput)
}

func (o CloudNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CloudNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CloudName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CloudNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CloudNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CloudName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CloudNamePtrOutput struct{ *pulumi.OutputState }

func (CloudNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudName)(nil)).Elem()
}

func (o CloudNamePtrOutput) ToCloudNamePtrOutput() CloudNamePtrOutput {
	return o
}

func (o CloudNamePtrOutput) ToCloudNamePtrOutputWithContext(ctx context.Context) CloudNamePtrOutput {
	return o
}

func (o CloudNamePtrOutput) Elem() CloudNameOutput {
	return o.ApplyT(func(v *CloudName) CloudName {
		if v != nil {
			return *v
		}
		var ret CloudName
		return ret
	}).(CloudNameOutput)
}

func (o CloudNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CloudNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CloudName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CloudNameInput interface {
	pulumi.Input

	ToCloudNameOutput() CloudNameOutput
	ToCloudNameOutputWithContext(context.Context) CloudNameOutput
}

var cloudNamePtrType = reflect.TypeOf((**CloudName)(nil)).Elem()

type CloudNamePtrInput interface {
	pulumi.Input

	ToCloudNamePtrOutput() CloudNamePtrOutput
	ToCloudNamePtrOutputWithContext(context.Context) CloudNamePtrOutput
}

type cloudNamePtr string

func CloudNamePtr(v string) CloudNamePtrInput {
	return (*cloudNamePtr)(&v)
}

func (*cloudNamePtr) ElementType() reflect.Type {
	return cloudNamePtrType
}

func (in *cloudNamePtr) ToCloudNamePtrOutput() CloudNamePtrOutput {
	return pulumi.ToOutput(in).(CloudNamePtrOutput)
}

func (in *cloudNamePtr) ToCloudNamePtrOutputWithContext(ctx context.Context) CloudNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CloudNamePtrOutput)
}

type ImplementationEffortEnum string

const (
	ImplementationEffortEnumHigh     = ImplementationEffortEnum("High")
	ImplementationEffortEnumModerate = ImplementationEffortEnum("Moderate")
	ImplementationEffortEnumLow      = ImplementationEffortEnum("Low")
)

func (ImplementationEffortEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ImplementationEffortEnum)(nil)).Elem()
}

func (e ImplementationEffortEnum) ToImplementationEffortEnumOutput() ImplementationEffortEnumOutput {
	return pulumi.ToOutput(e).(ImplementationEffortEnumOutput)
}

func (e ImplementationEffortEnum) ToImplementationEffortEnumOutputWithContext(ctx context.Context) ImplementationEffortEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ImplementationEffortEnumOutput)
}

func (e ImplementationEffortEnum) ToImplementationEffortEnumPtrOutput() ImplementationEffortEnumPtrOutput {
	return e.ToImplementationEffortEnumPtrOutputWithContext(context.Background())
}

func (e ImplementationEffortEnum) ToImplementationEffortEnumPtrOutputWithContext(ctx context.Context) ImplementationEffortEnumPtrOutput {
	return ImplementationEffortEnum(e).ToImplementationEffortEnumOutputWithContext(ctx).ToImplementationEffortEnumPtrOutputWithContext(ctx)
}

func (e ImplementationEffortEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImplementationEffortEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImplementationEffortEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ImplementationEffortEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ImplementationEffortEnumOutput struct{ *pulumi.OutputState }

func (ImplementationEffortEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImplementationEffortEnum)(nil)).Elem()
}

func (o ImplementationEffortEnumOutput) ToImplementationEffortEnumOutput() ImplementationEffortEnumOutput {
	return o
}

func (o ImplementationEffortEnumOutput) ToImplementationEffortEnumOutputWithContext(ctx context.Context) ImplementationEffortEnumOutput {
	return o
}

func (o ImplementationEffortEnumOutput) ToImplementationEffortEnumPtrOutput() ImplementationEffortEnumPtrOutput {
	return o.ToImplementationEffortEnumPtrOutputWithContext(context.Background())
}

func (o ImplementationEffortEnumOutput) ToImplementationEffortEnumPtrOutputWithContext(ctx context.Context) ImplementationEffortEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImplementationEffortEnum) *ImplementationEffortEnum {
		return &v
	}).(ImplementationEffortEnumPtrOutput)
}

func (o ImplementationEffortEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ImplementationEffortEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ImplementationEffortEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ImplementationEffortEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ImplementationEffortEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ImplementationEffortEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ImplementationEffortEnumPtrOutput struct{ *pulumi.OutputState }

func (ImplementationEffortEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImplementationEffortEnum)(nil)).Elem()
}

func (o ImplementationEffortEnumPtrOutput) ToImplementationEffortEnumPtrOutput() ImplementationEffortEnumPtrOutput {
	return o
}

func (o ImplementationEffortEnumPtrOutput) ToImplementationEffortEnumPtrOutputWithContext(ctx context.Context) ImplementationEffortEnumPtrOutput {
	return o
}

func (o ImplementationEffortEnumPtrOutput) Elem() ImplementationEffortEnumOutput {
	return o.ApplyT(func(v *ImplementationEffortEnum) ImplementationEffortEnum {
		if v != nil {
			return *v
		}
		var ret ImplementationEffortEnum
		return ret
	}).(ImplementationEffortEnumOutput)
}

func (o ImplementationEffortEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ImplementationEffortEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ImplementationEffortEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ImplementationEffortEnumInput interface {
	pulumi.Input

	ToImplementationEffortEnumOutput() ImplementationEffortEnumOutput
	ToImplementationEffortEnumOutputWithContext(context.Context) ImplementationEffortEnumOutput
}

var implementationEffortEnumPtrType = reflect.TypeOf((**ImplementationEffortEnum)(nil)).Elem()

type ImplementationEffortEnumPtrInput interface {
	pulumi.Input

	ToImplementationEffortEnumPtrOutput() ImplementationEffortEnumPtrOutput
	ToImplementationEffortEnumPtrOutputWithContext(context.Context) ImplementationEffortEnumPtrOutput
}

type implementationEffortEnumPtr string

func ImplementationEffortEnumPtr(v string) ImplementationEffortEnumPtrInput {
	return (*implementationEffortEnumPtr)(&v)
}

func (*implementationEffortEnumPtr) ElementType() reflect.Type {
	return implementationEffortEnumPtrType
}

func (in *implementationEffortEnumPtr) ToImplementationEffortEnumPtrOutput() ImplementationEffortEnumPtrOutput {
	return pulumi.ToOutput(in).(ImplementationEffortEnumPtrOutput)
}

func (in *implementationEffortEnumPtr) ToImplementationEffortEnumPtrOutputWithContext(ctx context.Context) ImplementationEffortEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ImplementationEffortEnumPtrOutput)
}

type OfferingType string

const (
	OfferingTypeCspmMonitorAws           = OfferingType("CspmMonitorAws")
	OfferingTypeDefenderForContainersAws = OfferingType("DefenderForContainersAws")
	OfferingTypeDefenderForServersAws    = OfferingType("DefenderForServersAws")
)

func (OfferingType) ElementType() reflect.Type {
	return reflect.TypeOf((*OfferingType)(nil)).Elem()
}

func (e OfferingType) ToOfferingTypeOutput() OfferingTypeOutput {
	return pulumi.ToOutput(e).(OfferingTypeOutput)
}

func (e OfferingType) ToOfferingTypeOutputWithContext(ctx context.Context) OfferingTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OfferingTypeOutput)
}

func (e OfferingType) ToOfferingTypePtrOutput() OfferingTypePtrOutput {
	return e.ToOfferingTypePtrOutputWithContext(context.Background())
}

func (e OfferingType) ToOfferingTypePtrOutputWithContext(ctx context.Context) OfferingTypePtrOutput {
	return OfferingType(e).ToOfferingTypeOutputWithContext(ctx).ToOfferingTypePtrOutputWithContext(ctx)
}

func (e OfferingType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OfferingType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OfferingType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OfferingType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OfferingTypeOutput struct{ *pulumi.OutputState }

func (OfferingTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfferingType)(nil)).Elem()
}

func (o OfferingTypeOutput) ToOfferingTypeOutput() OfferingTypeOutput {
	return o
}

func (o OfferingTypeOutput) ToOfferingTypeOutputWithContext(ctx context.Context) OfferingTypeOutput {
	return o
}

func (o OfferingTypeOutput) ToOfferingTypePtrOutput() OfferingTypePtrOutput {
	return o.ToOfferingTypePtrOutputWithContext(context.Background())
}

func (o OfferingTypeOutput) ToOfferingTypePtrOutputWithContext(ctx context.Context) OfferingTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfferingType) *OfferingType {
		return &v
	}).(OfferingTypePtrOutput)
}

func (o OfferingTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OfferingTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OfferingType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OfferingTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OfferingTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OfferingType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OfferingTypePtrOutput struct{ *pulumi.OutputState }

func (OfferingTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfferingType)(nil)).Elem()
}

func (o OfferingTypePtrOutput) ToOfferingTypePtrOutput() OfferingTypePtrOutput {
	return o
}

func (o OfferingTypePtrOutput) ToOfferingTypePtrOutputWithContext(ctx context.Context) OfferingTypePtrOutput {
	return o
}

func (o OfferingTypePtrOutput) Elem() OfferingTypeOutput {
	return o.ApplyT(func(v *OfferingType) OfferingType {
		if v != nil {
			return *v
		}
		var ret OfferingType
		return ret
	}).(OfferingTypeOutput)
}

func (o OfferingTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OfferingTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OfferingType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OfferingTypeInput interface {
	pulumi.Input

	ToOfferingTypeOutput() OfferingTypeOutput
	ToOfferingTypeOutputWithContext(context.Context) OfferingTypeOutput
}

var offeringTypePtrType = reflect.TypeOf((**OfferingType)(nil)).Elem()

type OfferingTypePtrInput interface {
	pulumi.Input

	ToOfferingTypePtrOutput() OfferingTypePtrOutput
	ToOfferingTypePtrOutputWithContext(context.Context) OfferingTypePtrOutput
}

type offeringTypePtr string

func OfferingTypePtr(v string) OfferingTypePtrInput {
	return (*offeringTypePtr)(&v)
}

func (*offeringTypePtr) ElementType() reflect.Type {
	return offeringTypePtrType
}

func (in *offeringTypePtr) ToOfferingTypePtrOutput() OfferingTypePtrOutput {
	return pulumi.ToOutput(in).(OfferingTypePtrOutput)
}

func (in *offeringTypePtr) ToOfferingTypePtrOutputWithContext(ctx context.Context) OfferingTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OfferingTypePtrOutput)
}

type OrganizationMembershipType string

const (
	OrganizationMembershipTypeMember       = OrganizationMembershipType("Member")
	OrganizationMembershipTypeOrganization = OrganizationMembershipType("Organization")
)

func (OrganizationMembershipType) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationMembershipType)(nil)).Elem()
}

func (e OrganizationMembershipType) ToOrganizationMembershipTypeOutput() OrganizationMembershipTypeOutput {
	return pulumi.ToOutput(e).(OrganizationMembershipTypeOutput)
}

func (e OrganizationMembershipType) ToOrganizationMembershipTypeOutputWithContext(ctx context.Context) OrganizationMembershipTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OrganizationMembershipTypeOutput)
}

func (e OrganizationMembershipType) ToOrganizationMembershipTypePtrOutput() OrganizationMembershipTypePtrOutput {
	return e.ToOrganizationMembershipTypePtrOutputWithContext(context.Background())
}

func (e OrganizationMembershipType) ToOrganizationMembershipTypePtrOutputWithContext(ctx context.Context) OrganizationMembershipTypePtrOutput {
	return OrganizationMembershipType(e).ToOrganizationMembershipTypeOutputWithContext(ctx).ToOrganizationMembershipTypePtrOutputWithContext(ctx)
}

func (e OrganizationMembershipType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrganizationMembershipType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrganizationMembershipType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OrganizationMembershipType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OrganizationMembershipTypeOutput struct{ *pulumi.OutputState }

func (OrganizationMembershipTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationMembershipType)(nil)).Elem()
}

func (o OrganizationMembershipTypeOutput) ToOrganizationMembershipTypeOutput() OrganizationMembershipTypeOutput {
	return o
}

func (o OrganizationMembershipTypeOutput) ToOrganizationMembershipTypeOutputWithContext(ctx context.Context) OrganizationMembershipTypeOutput {
	return o
}

func (o OrganizationMembershipTypeOutput) ToOrganizationMembershipTypePtrOutput() OrganizationMembershipTypePtrOutput {
	return o.ToOrganizationMembershipTypePtrOutputWithContext(context.Background())
}

func (o OrganizationMembershipTypeOutput) ToOrganizationMembershipTypePtrOutputWithContext(ctx context.Context) OrganizationMembershipTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrganizationMembershipType) *OrganizationMembershipType {
		return &v
	}).(OrganizationMembershipTypePtrOutput)
}

func (o OrganizationMembershipTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OrganizationMembershipTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OrganizationMembershipType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OrganizationMembershipTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OrganizationMembershipTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OrganizationMembershipType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OrganizationMembershipTypePtrOutput struct{ *pulumi.OutputState }

func (OrganizationMembershipTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationMembershipType)(nil)).Elem()
}

func (o OrganizationMembershipTypePtrOutput) ToOrganizationMembershipTypePtrOutput() OrganizationMembershipTypePtrOutput {
	return o
}

func (o OrganizationMembershipTypePtrOutput) ToOrganizationMembershipTypePtrOutputWithContext(ctx context.Context) OrganizationMembershipTypePtrOutput {
	return o
}

func (o OrganizationMembershipTypePtrOutput) Elem() OrganizationMembershipTypeOutput {
	return o.ApplyT(func(v *OrganizationMembershipType) OrganizationMembershipType {
		if v != nil {
			return *v
		}
		var ret OrganizationMembershipType
		return ret
	}).(OrganizationMembershipTypeOutput)
}

func (o OrganizationMembershipTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OrganizationMembershipTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OrganizationMembershipType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OrganizationMembershipTypeInput interface {
	pulumi.Input

	ToOrganizationMembershipTypeOutput() OrganizationMembershipTypeOutput
	ToOrganizationMembershipTypeOutputWithContext(context.Context) OrganizationMembershipTypeOutput
}

var organizationMembershipTypePtrType = reflect.TypeOf((**OrganizationMembershipType)(nil)).Elem()

type OrganizationMembershipTypePtrInput interface {
	pulumi.Input

	ToOrganizationMembershipTypePtrOutput() OrganizationMembershipTypePtrOutput
	ToOrganizationMembershipTypePtrOutputWithContext(context.Context) OrganizationMembershipTypePtrOutput
}

type organizationMembershipTypePtr string

func OrganizationMembershipTypePtr(v string) OrganizationMembershipTypePtrInput {
	return (*organizationMembershipTypePtr)(&v)
}

func (*organizationMembershipTypePtr) ElementType() reflect.Type {
	return organizationMembershipTypePtrType
}

func (in *organizationMembershipTypePtr) ToOrganizationMembershipTypePtrOutput() OrganizationMembershipTypePtrOutput {
	return pulumi.ToOutput(in).(OrganizationMembershipTypePtrOutput)
}

func (in *organizationMembershipTypePtr) ToOrganizationMembershipTypePtrOutputWithContext(ctx context.Context) OrganizationMembershipTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OrganizationMembershipTypePtrOutput)
}

type SeverityEnum string

const (
	SeverityEnumHigh   = SeverityEnum("High")
	SeverityEnumMedium = SeverityEnum("Medium")
	SeverityEnumLow    = SeverityEnum("Low")
)

func (SeverityEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SeverityEnum)(nil)).Elem()
}

func (e SeverityEnum) ToSeverityEnumOutput() SeverityEnumOutput {
	return pulumi.ToOutput(e).(SeverityEnumOutput)
}

func (e SeverityEnum) ToSeverityEnumOutputWithContext(ctx context.Context) SeverityEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SeverityEnumOutput)
}

func (e SeverityEnum) ToSeverityEnumPtrOutput() SeverityEnumPtrOutput {
	return e.ToSeverityEnumPtrOutputWithContext(context.Background())
}

func (e SeverityEnum) ToSeverityEnumPtrOutputWithContext(ctx context.Context) SeverityEnumPtrOutput {
	return SeverityEnum(e).ToSeverityEnumOutputWithContext(ctx).ToSeverityEnumPtrOutputWithContext(ctx)
}

func (e SeverityEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SeverityEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SeverityEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SeverityEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SeverityEnumOutput struct{ *pulumi.OutputState }

func (SeverityEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SeverityEnum)(nil)).Elem()
}

func (o SeverityEnumOutput) ToSeverityEnumOutput() SeverityEnumOutput {
	return o
}

func (o SeverityEnumOutput) ToSeverityEnumOutputWithContext(ctx context.Context) SeverityEnumOutput {
	return o
}

func (o SeverityEnumOutput) ToSeverityEnumPtrOutput() SeverityEnumPtrOutput {
	return o.ToSeverityEnumPtrOutputWithContext(context.Background())
}

func (o SeverityEnumOutput) ToSeverityEnumPtrOutputWithContext(ctx context.Context) SeverityEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SeverityEnum) *SeverityEnum {
		return &v
	}).(SeverityEnumPtrOutput)
}

func (o SeverityEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SeverityEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SeverityEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SeverityEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SeverityEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SeverityEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SeverityEnumPtrOutput struct{ *pulumi.OutputState }

func (SeverityEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SeverityEnum)(nil)).Elem()
}

func (o SeverityEnumPtrOutput) ToSeverityEnumPtrOutput() SeverityEnumPtrOutput {
	return o
}

func (o SeverityEnumPtrOutput) ToSeverityEnumPtrOutputWithContext(ctx context.Context) SeverityEnumPtrOutput {
	return o
}

func (o SeverityEnumPtrOutput) Elem() SeverityEnumOutput {
	return o.ApplyT(func(v *SeverityEnum) SeverityEnum {
		if v != nil {
			return *v
		}
		var ret SeverityEnum
		return ret
	}).(SeverityEnumOutput)
}

func (o SeverityEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SeverityEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SeverityEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SeverityEnumInput interface {
	pulumi.Input

	ToSeverityEnumOutput() SeverityEnumOutput
	ToSeverityEnumOutputWithContext(context.Context) SeverityEnumOutput
}

var severityEnumPtrType = reflect.TypeOf((**SeverityEnum)(nil)).Elem()

type SeverityEnumPtrInput interface {
	pulumi.Input

	ToSeverityEnumPtrOutput() SeverityEnumPtrOutput
	ToSeverityEnumPtrOutputWithContext(context.Context) SeverityEnumPtrOutput
}

type severityEnumPtr string

func SeverityEnumPtr(v string) SeverityEnumPtrInput {
	return (*severityEnumPtr)(&v)
}

func (*severityEnumPtr) ElementType() reflect.Type {
	return severityEnumPtrType
}

func (in *severityEnumPtr) ToSeverityEnumPtrOutput() SeverityEnumPtrOutput {
	return pulumi.ToOutput(in).(SeverityEnumPtrOutput)
}

func (in *severityEnumPtr) ToSeverityEnumPtrOutputWithContext(ctx context.Context) SeverityEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SeverityEnumPtrOutput)
}

type SupportedCloudEnum string

const (
	SupportedCloudEnumAWS = SupportedCloudEnum("AWS")
)

func (SupportedCloudEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedCloudEnum)(nil)).Elem()
}

func (e SupportedCloudEnum) ToSupportedCloudEnumOutput() SupportedCloudEnumOutput {
	return pulumi.ToOutput(e).(SupportedCloudEnumOutput)
}

func (e SupportedCloudEnum) ToSupportedCloudEnumOutputWithContext(ctx context.Context) SupportedCloudEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SupportedCloudEnumOutput)
}

func (e SupportedCloudEnum) ToSupportedCloudEnumPtrOutput() SupportedCloudEnumPtrOutput {
	return e.ToSupportedCloudEnumPtrOutputWithContext(context.Background())
}

func (e SupportedCloudEnum) ToSupportedCloudEnumPtrOutputWithContext(ctx context.Context) SupportedCloudEnumPtrOutput {
	return SupportedCloudEnum(e).ToSupportedCloudEnumOutputWithContext(ctx).ToSupportedCloudEnumPtrOutputWithContext(ctx)
}

func (e SupportedCloudEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedCloudEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedCloudEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SupportedCloudEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SupportedCloudEnumOutput struct{ *pulumi.OutputState }

func (SupportedCloudEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedCloudEnum)(nil)).Elem()
}

func (o SupportedCloudEnumOutput) ToSupportedCloudEnumOutput() SupportedCloudEnumOutput {
	return o
}

func (o SupportedCloudEnumOutput) ToSupportedCloudEnumOutputWithContext(ctx context.Context) SupportedCloudEnumOutput {
	return o
}

func (o SupportedCloudEnumOutput) ToSupportedCloudEnumPtrOutput() SupportedCloudEnumPtrOutput {
	return o.ToSupportedCloudEnumPtrOutputWithContext(context.Background())
}

func (o SupportedCloudEnumOutput) ToSupportedCloudEnumPtrOutputWithContext(ctx context.Context) SupportedCloudEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SupportedCloudEnum) *SupportedCloudEnum {
		return &v
	}).(SupportedCloudEnumPtrOutput)
}

func (o SupportedCloudEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SupportedCloudEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedCloudEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SupportedCloudEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedCloudEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedCloudEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SupportedCloudEnumPtrOutput struct{ *pulumi.OutputState }

func (SupportedCloudEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportedCloudEnum)(nil)).Elem()
}

func (o SupportedCloudEnumPtrOutput) ToSupportedCloudEnumPtrOutput() SupportedCloudEnumPtrOutput {
	return o
}

func (o SupportedCloudEnumPtrOutput) ToSupportedCloudEnumPtrOutputWithContext(ctx context.Context) SupportedCloudEnumPtrOutput {
	return o
}

func (o SupportedCloudEnumPtrOutput) Elem() SupportedCloudEnumOutput {
	return o.ApplyT(func(v *SupportedCloudEnum) SupportedCloudEnum {
		if v != nil {
			return *v
		}
		var ret SupportedCloudEnum
		return ret
	}).(SupportedCloudEnumOutput)
}

func (o SupportedCloudEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedCloudEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SupportedCloudEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SupportedCloudEnumInput interface {
	pulumi.Input

	ToSupportedCloudEnumOutput() SupportedCloudEnumOutput
	ToSupportedCloudEnumOutputWithContext(context.Context) SupportedCloudEnumOutput
}

var supportedCloudEnumPtrType = reflect.TypeOf((**SupportedCloudEnum)(nil)).Elem()

type SupportedCloudEnumPtrInput interface {
	pulumi.Input

	ToSupportedCloudEnumPtrOutput() SupportedCloudEnumPtrOutput
	ToSupportedCloudEnumPtrOutputWithContext(context.Context) SupportedCloudEnumPtrOutput
}

type supportedCloudEnumPtr string

func SupportedCloudEnumPtr(v string) SupportedCloudEnumPtrInput {
	return (*supportedCloudEnumPtr)(&v)
}

func (*supportedCloudEnumPtr) ElementType() reflect.Type {
	return supportedCloudEnumPtrType
}

func (in *supportedCloudEnumPtr) ToSupportedCloudEnumPtrOutput() SupportedCloudEnumPtrOutput {
	return pulumi.ToOutput(in).(SupportedCloudEnumPtrOutput)
}

func (in *supportedCloudEnumPtr) ToSupportedCloudEnumPtrOutputWithContext(ctx context.Context) SupportedCloudEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SupportedCloudEnumPtrOutput)
}

type UserImpactEnum string

const (
	UserImpactEnumHigh     = UserImpactEnum("High")
	UserImpactEnumModerate = UserImpactEnum("Moderate")
	UserImpactEnumLow      = UserImpactEnum("Low")
)

func (UserImpactEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*UserImpactEnum)(nil)).Elem()
}

func (e UserImpactEnum) ToUserImpactEnumOutput() UserImpactEnumOutput {
	return pulumi.ToOutput(e).(UserImpactEnumOutput)
}

func (e UserImpactEnum) ToUserImpactEnumOutputWithContext(ctx context.Context) UserImpactEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UserImpactEnumOutput)
}

func (e UserImpactEnum) ToUserImpactEnumPtrOutput() UserImpactEnumPtrOutput {
	return e.ToUserImpactEnumPtrOutputWithContext(context.Background())
}

func (e UserImpactEnum) ToUserImpactEnumPtrOutputWithContext(ctx context.Context) UserImpactEnumPtrOutput {
	return UserImpactEnum(e).ToUserImpactEnumOutputWithContext(ctx).ToUserImpactEnumPtrOutputWithContext(ctx)
}

func (e UserImpactEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserImpactEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserImpactEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UserImpactEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UserImpactEnumOutput struct{ *pulumi.OutputState }

func (UserImpactEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserImpactEnum)(nil)).Elem()
}

func (o UserImpactEnumOutput) ToUserImpactEnumOutput() UserImpactEnumOutput {
	return o
}

func (o UserImpactEnumOutput) ToUserImpactEnumOutputWithContext(ctx context.Context) UserImpactEnumOutput {
	return o
}

func (o UserImpactEnumOutput) ToUserImpactEnumPtrOutput() UserImpactEnumPtrOutput {
	return o.ToUserImpactEnumPtrOutputWithContext(context.Background())
}

func (o UserImpactEnumOutput) ToUserImpactEnumPtrOutputWithContext(ctx context.Context) UserImpactEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserImpactEnum) *UserImpactEnum {
		return &v
	}).(UserImpactEnumPtrOutput)
}

func (o UserImpactEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UserImpactEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserImpactEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UserImpactEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserImpactEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserImpactEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UserImpactEnumPtrOutput struct{ *pulumi.OutputState }

func (UserImpactEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserImpactEnum)(nil)).Elem()
}

func (o UserImpactEnumPtrOutput) ToUserImpactEnumPtrOutput() UserImpactEnumPtrOutput {
	return o
}

func (o UserImpactEnumPtrOutput) ToUserImpactEnumPtrOutputWithContext(ctx context.Context) UserImpactEnumPtrOutput {
	return o
}

func (o UserImpactEnumPtrOutput) Elem() UserImpactEnumOutput {
	return o.ApplyT(func(v *UserImpactEnum) UserImpactEnum {
		if v != nil {
			return *v
		}
		var ret UserImpactEnum
		return ret
	}).(UserImpactEnumOutput)
}

func (o UserImpactEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserImpactEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UserImpactEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UserImpactEnumInput interface {
	pulumi.Input

	ToUserImpactEnumOutput() UserImpactEnumOutput
	ToUserImpactEnumOutputWithContext(context.Context) UserImpactEnumOutput
}

var userImpactEnumPtrType = reflect.TypeOf((**UserImpactEnum)(nil)).Elem()

type UserImpactEnumPtrInput interface {
	pulumi.Input

	ToUserImpactEnumPtrOutput() UserImpactEnumPtrOutput
	ToUserImpactEnumPtrOutputWithContext(context.Context) UserImpactEnumPtrOutput
}

type userImpactEnumPtr string

func UserImpactEnumPtr(v string) UserImpactEnumPtrInput {
	return (*userImpactEnumPtr)(&v)
}

func (*userImpactEnumPtr) ElementType() reflect.Type {
	return userImpactEnumPtrType
}

func (in *userImpactEnumPtr) ToUserImpactEnumPtrOutput() UserImpactEnumPtrOutput {
	return pulumi.ToOutput(in).(UserImpactEnumPtrOutput)
}

func (in *userImpactEnumPtr) ToUserImpactEnumPtrOutputWithContext(ctx context.Context) UserImpactEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UserImpactEnumPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudNameInput)(nil)).Elem(), CloudName("Azure"))
	pulumi.RegisterInputType(reflect.TypeOf((*CloudNamePtrInput)(nil)).Elem(), CloudName("Azure"))
	pulumi.RegisterInputType(reflect.TypeOf((*ImplementationEffortEnumInput)(nil)).Elem(), ImplementationEffortEnum("High"))
	pulumi.RegisterInputType(reflect.TypeOf((*ImplementationEffortEnumPtrInput)(nil)).Elem(), ImplementationEffortEnum("High"))
	pulumi.RegisterInputType(reflect.TypeOf((*OfferingTypeInput)(nil)).Elem(), OfferingType("CspmMonitorAws"))
	pulumi.RegisterInputType(reflect.TypeOf((*OfferingTypePtrInput)(nil)).Elem(), OfferingType("CspmMonitorAws"))
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationMembershipTypeInput)(nil)).Elem(), OrganizationMembershipType("Member"))
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationMembershipTypePtrInput)(nil)).Elem(), OrganizationMembershipType("Member"))
	pulumi.RegisterInputType(reflect.TypeOf((*SeverityEnumInput)(nil)).Elem(), SeverityEnum("High"))
	pulumi.RegisterInputType(reflect.TypeOf((*SeverityEnumPtrInput)(nil)).Elem(), SeverityEnum("High"))
	pulumi.RegisterInputType(reflect.TypeOf((*SupportedCloudEnumInput)(nil)).Elem(), SupportedCloudEnum("AWS"))
	pulumi.RegisterInputType(reflect.TypeOf((*SupportedCloudEnumPtrInput)(nil)).Elem(), SupportedCloudEnum("AWS"))
	pulumi.RegisterInputType(reflect.TypeOf((*UserImpactEnumInput)(nil)).Elem(), UserImpactEnum("High"))
	pulumi.RegisterInputType(reflect.TypeOf((*UserImpactEnumPtrInput)(nil)).Elem(), UserImpactEnum("High"))
	pulumi.RegisterOutputType(CloudNameOutput{})
	pulumi.RegisterOutputType(CloudNamePtrOutput{})
	pulumi.RegisterOutputType(ImplementationEffortEnumOutput{})
	pulumi.RegisterOutputType(ImplementationEffortEnumPtrOutput{})
	pulumi.RegisterOutputType(OfferingTypeOutput{})
	pulumi.RegisterOutputType(OfferingTypePtrOutput{})
	pulumi.RegisterOutputType(OrganizationMembershipTypeOutput{})
	pulumi.RegisterOutputType(OrganizationMembershipTypePtrOutput{})
	pulumi.RegisterOutputType(SeverityEnumOutput{})
	pulumi.RegisterOutputType(SeverityEnumPtrOutput{})
	pulumi.RegisterOutputType(SupportedCloudEnumOutput{})
	pulumi.RegisterOutputType(SupportedCloudEnumPtrOutput{})
	pulumi.RegisterOutputType(UserImpactEnumOutput{})
	pulumi.RegisterOutputType(UserImpactEnumPtrOutput{})
}
