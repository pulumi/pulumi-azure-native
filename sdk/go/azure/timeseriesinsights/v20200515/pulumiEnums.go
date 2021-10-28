


package v20200515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessPolicyRole string

const (
	AccessPolicyRoleReader      = AccessPolicyRole("Reader")
	AccessPolicyRoleContributor = AccessPolicyRole("Contributor")
)

func (AccessPolicyRole) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyRole)(nil)).Elem()
}

func (e AccessPolicyRole) ToAccessPolicyRoleOutput() AccessPolicyRoleOutput {
	return pulumi.ToOutput(e).(AccessPolicyRoleOutput)
}

func (e AccessPolicyRole) ToAccessPolicyRoleOutputWithContext(ctx context.Context) AccessPolicyRoleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessPolicyRoleOutput)
}

func (e AccessPolicyRole) ToAccessPolicyRolePtrOutput() AccessPolicyRolePtrOutput {
	return e.ToAccessPolicyRolePtrOutputWithContext(context.Background())
}

func (e AccessPolicyRole) ToAccessPolicyRolePtrOutputWithContext(ctx context.Context) AccessPolicyRolePtrOutput {
	return AccessPolicyRole(e).ToAccessPolicyRoleOutputWithContext(ctx).ToAccessPolicyRolePtrOutputWithContext(ctx)
}

func (e AccessPolicyRole) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessPolicyRole) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessPolicyRole) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessPolicyRole) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessPolicyRoleOutput struct{ *pulumi.OutputState }

func (AccessPolicyRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyRole)(nil)).Elem()
}

func (o AccessPolicyRoleOutput) ToAccessPolicyRoleOutput() AccessPolicyRoleOutput {
	return o
}

func (o AccessPolicyRoleOutput) ToAccessPolicyRoleOutputWithContext(ctx context.Context) AccessPolicyRoleOutput {
	return o
}

func (o AccessPolicyRoleOutput) ToAccessPolicyRolePtrOutput() AccessPolicyRolePtrOutput {
	return o.ToAccessPolicyRolePtrOutputWithContext(context.Background())
}

func (o AccessPolicyRoleOutput) ToAccessPolicyRolePtrOutputWithContext(ctx context.Context) AccessPolicyRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessPolicyRole) *AccessPolicyRole {
		return &v
	}).(AccessPolicyRolePtrOutput)
}

func (o AccessPolicyRoleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessPolicyRoleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessPolicyRole) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessPolicyRoleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPolicyRoleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessPolicyRole) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessPolicyRolePtrOutput struct{ *pulumi.OutputState }

func (AccessPolicyRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicyRole)(nil)).Elem()
}

func (o AccessPolicyRolePtrOutput) ToAccessPolicyRolePtrOutput() AccessPolicyRolePtrOutput {
	return o
}

func (o AccessPolicyRolePtrOutput) ToAccessPolicyRolePtrOutputWithContext(ctx context.Context) AccessPolicyRolePtrOutput {
	return o
}

func (o AccessPolicyRolePtrOutput) Elem() AccessPolicyRoleOutput {
	return o.ApplyT(func(v *AccessPolicyRole) AccessPolicyRole {
		if v != nil {
			return *v
		}
		var ret AccessPolicyRole
		return ret
	}).(AccessPolicyRoleOutput)
}

func (o AccessPolicyRolePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPolicyRolePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessPolicyRole) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessPolicyRoleInput interface {
	pulumi.Input

	ToAccessPolicyRoleOutput() AccessPolicyRoleOutput
	ToAccessPolicyRoleOutputWithContext(context.Context) AccessPolicyRoleOutput
}

var accessPolicyRolePtrType = reflect.TypeOf((**AccessPolicyRole)(nil)).Elem()

type AccessPolicyRolePtrInput interface {
	pulumi.Input

	ToAccessPolicyRolePtrOutput() AccessPolicyRolePtrOutput
	ToAccessPolicyRolePtrOutputWithContext(context.Context) AccessPolicyRolePtrOutput
}

type accessPolicyRolePtr string

func AccessPolicyRolePtr(v string) AccessPolicyRolePtrInput {
	return (*accessPolicyRolePtr)(&v)
}

func (*accessPolicyRolePtr) ElementType() reflect.Type {
	return accessPolicyRolePtrType
}

func (in *accessPolicyRolePtr) ToAccessPolicyRolePtrOutput() AccessPolicyRolePtrOutput {
	return pulumi.ToOutput(in).(AccessPolicyRolePtrOutput)
}

func (in *accessPolicyRolePtr) ToAccessPolicyRolePtrOutputWithContext(ctx context.Context) AccessPolicyRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessPolicyRolePtrOutput)
}

type DataStringComparisonBehavior string

const (
	DataStringComparisonBehaviorOrdinal           = DataStringComparisonBehavior("Ordinal")
	DataStringComparisonBehaviorOrdinalIgnoreCase = DataStringComparisonBehavior("OrdinalIgnoreCase")
)

func (DataStringComparisonBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStringComparisonBehavior)(nil)).Elem()
}

func (e DataStringComparisonBehavior) ToDataStringComparisonBehaviorOutput() DataStringComparisonBehaviorOutput {
	return pulumi.ToOutput(e).(DataStringComparisonBehaviorOutput)
}

func (e DataStringComparisonBehavior) ToDataStringComparisonBehaviorOutputWithContext(ctx context.Context) DataStringComparisonBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataStringComparisonBehaviorOutput)
}

func (e DataStringComparisonBehavior) ToDataStringComparisonBehaviorPtrOutput() DataStringComparisonBehaviorPtrOutput {
	return e.ToDataStringComparisonBehaviorPtrOutputWithContext(context.Background())
}

func (e DataStringComparisonBehavior) ToDataStringComparisonBehaviorPtrOutputWithContext(ctx context.Context) DataStringComparisonBehaviorPtrOutput {
	return DataStringComparisonBehavior(e).ToDataStringComparisonBehaviorOutputWithContext(ctx).ToDataStringComparisonBehaviorPtrOutputWithContext(ctx)
}

func (e DataStringComparisonBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataStringComparisonBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataStringComparisonBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataStringComparisonBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataStringComparisonBehaviorOutput struct{ *pulumi.OutputState }

func (DataStringComparisonBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStringComparisonBehavior)(nil)).Elem()
}

func (o DataStringComparisonBehaviorOutput) ToDataStringComparisonBehaviorOutput() DataStringComparisonBehaviorOutput {
	return o
}

func (o DataStringComparisonBehaviorOutput) ToDataStringComparisonBehaviorOutputWithContext(ctx context.Context) DataStringComparisonBehaviorOutput {
	return o
}

func (o DataStringComparisonBehaviorOutput) ToDataStringComparisonBehaviorPtrOutput() DataStringComparisonBehaviorPtrOutput {
	return o.ToDataStringComparisonBehaviorPtrOutputWithContext(context.Background())
}

func (o DataStringComparisonBehaviorOutput) ToDataStringComparisonBehaviorPtrOutputWithContext(ctx context.Context) DataStringComparisonBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataStringComparisonBehavior) *DataStringComparisonBehavior {
		return &v
	}).(DataStringComparisonBehaviorPtrOutput)
}

func (o DataStringComparisonBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataStringComparisonBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataStringComparisonBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataStringComparisonBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataStringComparisonBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataStringComparisonBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataStringComparisonBehaviorPtrOutput struct{ *pulumi.OutputState }

func (DataStringComparisonBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataStringComparisonBehavior)(nil)).Elem()
}

func (o DataStringComparisonBehaviorPtrOutput) ToDataStringComparisonBehaviorPtrOutput() DataStringComparisonBehaviorPtrOutput {
	return o
}

func (o DataStringComparisonBehaviorPtrOutput) ToDataStringComparisonBehaviorPtrOutputWithContext(ctx context.Context) DataStringComparisonBehaviorPtrOutput {
	return o
}

func (o DataStringComparisonBehaviorPtrOutput) Elem() DataStringComparisonBehaviorOutput {
	return o.ApplyT(func(v *DataStringComparisonBehavior) DataStringComparisonBehavior {
		if v != nil {
			return *v
		}
		var ret DataStringComparisonBehavior
		return ret
	}).(DataStringComparisonBehaviorOutput)
}

func (o DataStringComparisonBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataStringComparisonBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataStringComparisonBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataStringComparisonBehaviorInput interface {
	pulumi.Input

	ToDataStringComparisonBehaviorOutput() DataStringComparisonBehaviorOutput
	ToDataStringComparisonBehaviorOutputWithContext(context.Context) DataStringComparisonBehaviorOutput
}

var dataStringComparisonBehaviorPtrType = reflect.TypeOf((**DataStringComparisonBehavior)(nil)).Elem()

type DataStringComparisonBehaviorPtrInput interface {
	pulumi.Input

	ToDataStringComparisonBehaviorPtrOutput() DataStringComparisonBehaviorPtrOutput
	ToDataStringComparisonBehaviorPtrOutputWithContext(context.Context) DataStringComparisonBehaviorPtrOutput
}

type dataStringComparisonBehaviorPtr string

func DataStringComparisonBehaviorPtr(v string) DataStringComparisonBehaviorPtrInput {
	return (*dataStringComparisonBehaviorPtr)(&v)
}

func (*dataStringComparisonBehaviorPtr) ElementType() reflect.Type {
	return dataStringComparisonBehaviorPtrType
}

func (in *dataStringComparisonBehaviorPtr) ToDataStringComparisonBehaviorPtrOutput() DataStringComparisonBehaviorPtrOutput {
	return pulumi.ToOutput(in).(DataStringComparisonBehaviorPtrOutput)
}

func (in *dataStringComparisonBehaviorPtr) ToDataStringComparisonBehaviorPtrOutputWithContext(ctx context.Context) DataStringComparisonBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataStringComparisonBehaviorPtrOutput)
}

type EnvironmentKind string

const (
	EnvironmentKindGen1 = EnvironmentKind("Gen1")
	EnvironmentKindGen2 = EnvironmentKind("Gen2")
)

func (EnvironmentKind) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentKind)(nil)).Elem()
}

func (e EnvironmentKind) ToEnvironmentKindOutput() EnvironmentKindOutput {
	return pulumi.ToOutput(e).(EnvironmentKindOutput)
}

func (e EnvironmentKind) ToEnvironmentKindOutputWithContext(ctx context.Context) EnvironmentKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnvironmentKindOutput)
}

func (e EnvironmentKind) ToEnvironmentKindPtrOutput() EnvironmentKindPtrOutput {
	return e.ToEnvironmentKindPtrOutputWithContext(context.Background())
}

func (e EnvironmentKind) ToEnvironmentKindPtrOutputWithContext(ctx context.Context) EnvironmentKindPtrOutput {
	return EnvironmentKind(e).ToEnvironmentKindOutputWithContext(ctx).ToEnvironmentKindPtrOutputWithContext(ctx)
}

func (e EnvironmentKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnvironmentKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnvironmentKindOutput struct{ *pulumi.OutputState }

func (EnvironmentKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentKind)(nil)).Elem()
}

func (o EnvironmentKindOutput) ToEnvironmentKindOutput() EnvironmentKindOutput {
	return o
}

func (o EnvironmentKindOutput) ToEnvironmentKindOutputWithContext(ctx context.Context) EnvironmentKindOutput {
	return o
}

func (o EnvironmentKindOutput) ToEnvironmentKindPtrOutput() EnvironmentKindPtrOutput {
	return o.ToEnvironmentKindPtrOutputWithContext(context.Background())
}

func (o EnvironmentKindOutput) ToEnvironmentKindPtrOutputWithContext(ctx context.Context) EnvironmentKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentKind) *EnvironmentKind {
		return &v
	}).(EnvironmentKindPtrOutput)
}

func (o EnvironmentKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnvironmentKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnvironmentKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnvironmentKindPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentKind)(nil)).Elem()
}

func (o EnvironmentKindPtrOutput) ToEnvironmentKindPtrOutput() EnvironmentKindPtrOutput {
	return o
}

func (o EnvironmentKindPtrOutput) ToEnvironmentKindPtrOutputWithContext(ctx context.Context) EnvironmentKindPtrOutput {
	return o
}

func (o EnvironmentKindPtrOutput) Elem() EnvironmentKindOutput {
	return o.ApplyT(func(v *EnvironmentKind) EnvironmentKind {
		if v != nil {
			return *v
		}
		var ret EnvironmentKind
		return ret
	}).(EnvironmentKindOutput)
}

func (o EnvironmentKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnvironmentKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EnvironmentKindInput interface {
	pulumi.Input

	ToEnvironmentKindOutput() EnvironmentKindOutput
	ToEnvironmentKindOutputWithContext(context.Context) EnvironmentKindOutput
}

var environmentKindPtrType = reflect.TypeOf((**EnvironmentKind)(nil)).Elem()

type EnvironmentKindPtrInput interface {
	pulumi.Input

	ToEnvironmentKindPtrOutput() EnvironmentKindPtrOutput
	ToEnvironmentKindPtrOutputWithContext(context.Context) EnvironmentKindPtrOutput
}

type environmentKindPtr string

func EnvironmentKindPtr(v string) EnvironmentKindPtrInput {
	return (*environmentKindPtr)(&v)
}

func (*environmentKindPtr) ElementType() reflect.Type {
	return environmentKindPtrType
}

func (in *environmentKindPtr) ToEnvironmentKindPtrOutput() EnvironmentKindPtrOutput {
	return pulumi.ToOutput(in).(EnvironmentKindPtrOutput)
}

func (in *environmentKindPtr) ToEnvironmentKindPtrOutputWithContext(ctx context.Context) EnvironmentKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnvironmentKindPtrOutput)
}

type EventSourceKind string

const (
	EventSourceKind_Microsoft_EventHub = EventSourceKind("Microsoft.EventHub")
	EventSourceKind_Microsoft_IoTHub   = EventSourceKind("Microsoft.IoTHub")
)

func (EventSourceKind) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSourceKind)(nil)).Elem()
}

func (e EventSourceKind) ToEventSourceKindOutput() EventSourceKindOutput {
	return pulumi.ToOutput(e).(EventSourceKindOutput)
}

func (e EventSourceKind) ToEventSourceKindOutputWithContext(ctx context.Context) EventSourceKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventSourceKindOutput)
}

func (e EventSourceKind) ToEventSourceKindPtrOutput() EventSourceKindPtrOutput {
	return e.ToEventSourceKindPtrOutputWithContext(context.Background())
}

func (e EventSourceKind) ToEventSourceKindPtrOutputWithContext(ctx context.Context) EventSourceKindPtrOutput {
	return EventSourceKind(e).ToEventSourceKindOutputWithContext(ctx).ToEventSourceKindPtrOutputWithContext(ctx)
}

func (e EventSourceKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSourceKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSourceKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventSourceKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventSourceKindOutput struct{ *pulumi.OutputState }

func (EventSourceKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSourceKind)(nil)).Elem()
}

func (o EventSourceKindOutput) ToEventSourceKindOutput() EventSourceKindOutput {
	return o
}

func (o EventSourceKindOutput) ToEventSourceKindOutputWithContext(ctx context.Context) EventSourceKindOutput {
	return o
}

func (o EventSourceKindOutput) ToEventSourceKindPtrOutput() EventSourceKindPtrOutput {
	return o.ToEventSourceKindPtrOutputWithContext(context.Background())
}

func (o EventSourceKindOutput) ToEventSourceKindPtrOutputWithContext(ctx context.Context) EventSourceKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSourceKind) *EventSourceKind {
		return &v
	}).(EventSourceKindPtrOutput)
}

func (o EventSourceKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventSourceKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSourceKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventSourceKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSourceKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSourceKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventSourceKindPtrOutput struct{ *pulumi.OutputState }

func (EventSourceKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSourceKind)(nil)).Elem()
}

func (o EventSourceKindPtrOutput) ToEventSourceKindPtrOutput() EventSourceKindPtrOutput {
	return o
}

func (o EventSourceKindPtrOutput) ToEventSourceKindPtrOutputWithContext(ctx context.Context) EventSourceKindPtrOutput {
	return o
}

func (o EventSourceKindPtrOutput) Elem() EventSourceKindOutput {
	return o.ApplyT(func(v *EventSourceKind) EventSourceKind {
		if v != nil {
			return *v
		}
		var ret EventSourceKind
		return ret
	}).(EventSourceKindOutput)
}

func (o EventSourceKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSourceKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventSourceKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EventSourceKindInput interface {
	pulumi.Input

	ToEventSourceKindOutput() EventSourceKindOutput
	ToEventSourceKindOutputWithContext(context.Context) EventSourceKindOutput
}

var eventSourceKindPtrType = reflect.TypeOf((**EventSourceKind)(nil)).Elem()

type EventSourceKindPtrInput interface {
	pulumi.Input

	ToEventSourceKindPtrOutput() EventSourceKindPtrOutput
	ToEventSourceKindPtrOutputWithContext(context.Context) EventSourceKindPtrOutput
}

type eventSourceKindPtr string

func EventSourceKindPtr(v string) EventSourceKindPtrInput {
	return (*eventSourceKindPtr)(&v)
}

func (*eventSourceKindPtr) ElementType() reflect.Type {
	return eventSourceKindPtrType
}

func (in *eventSourceKindPtr) ToEventSourceKindPtrOutput() EventSourceKindPtrOutput {
	return pulumi.ToOutput(in).(EventSourceKindPtrOutput)
}

func (in *eventSourceKindPtr) ToEventSourceKindPtrOutputWithContext(ctx context.Context) EventSourceKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventSourceKindPtrOutput)
}

type IngressStartAtType string

const (
	IngressStartAtTypeEarliestAvailable       = IngressStartAtType("EarliestAvailable")
	IngressStartAtTypeEventSourceCreationTime = IngressStartAtType("EventSourceCreationTime")
	IngressStartAtTypeCustomEnqueuedTime      = IngressStartAtType("CustomEnqueuedTime")
)

func (IngressStartAtType) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressStartAtType)(nil)).Elem()
}

func (e IngressStartAtType) ToIngressStartAtTypeOutput() IngressStartAtTypeOutput {
	return pulumi.ToOutput(e).(IngressStartAtTypeOutput)
}

func (e IngressStartAtType) ToIngressStartAtTypeOutputWithContext(ctx context.Context) IngressStartAtTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IngressStartAtTypeOutput)
}

func (e IngressStartAtType) ToIngressStartAtTypePtrOutput() IngressStartAtTypePtrOutput {
	return e.ToIngressStartAtTypePtrOutputWithContext(context.Background())
}

func (e IngressStartAtType) ToIngressStartAtTypePtrOutputWithContext(ctx context.Context) IngressStartAtTypePtrOutput {
	return IngressStartAtType(e).ToIngressStartAtTypeOutputWithContext(ctx).ToIngressStartAtTypePtrOutputWithContext(ctx)
}

func (e IngressStartAtType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IngressStartAtType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IngressStartAtType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IngressStartAtType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IngressStartAtTypeOutput struct{ *pulumi.OutputState }

func (IngressStartAtTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressStartAtType)(nil)).Elem()
}

func (o IngressStartAtTypeOutput) ToIngressStartAtTypeOutput() IngressStartAtTypeOutput {
	return o
}

func (o IngressStartAtTypeOutput) ToIngressStartAtTypeOutputWithContext(ctx context.Context) IngressStartAtTypeOutput {
	return o
}

func (o IngressStartAtTypeOutput) ToIngressStartAtTypePtrOutput() IngressStartAtTypePtrOutput {
	return o.ToIngressStartAtTypePtrOutputWithContext(context.Background())
}

func (o IngressStartAtTypeOutput) ToIngressStartAtTypePtrOutputWithContext(ctx context.Context) IngressStartAtTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IngressStartAtType) *IngressStartAtType {
		return &v
	}).(IngressStartAtTypePtrOutput)
}

func (o IngressStartAtTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IngressStartAtTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IngressStartAtType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IngressStartAtTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IngressStartAtTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IngressStartAtType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IngressStartAtTypePtrOutput struct{ *pulumi.OutputState }

func (IngressStartAtTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressStartAtType)(nil)).Elem()
}

func (o IngressStartAtTypePtrOutput) ToIngressStartAtTypePtrOutput() IngressStartAtTypePtrOutput {
	return o
}

func (o IngressStartAtTypePtrOutput) ToIngressStartAtTypePtrOutputWithContext(ctx context.Context) IngressStartAtTypePtrOutput {
	return o
}

func (o IngressStartAtTypePtrOutput) Elem() IngressStartAtTypeOutput {
	return o.ApplyT(func(v *IngressStartAtType) IngressStartAtType {
		if v != nil {
			return *v
		}
		var ret IngressStartAtType
		return ret
	}).(IngressStartAtTypeOutput)
}

func (o IngressStartAtTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IngressStartAtTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IngressStartAtType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IngressStartAtTypeInput interface {
	pulumi.Input

	ToIngressStartAtTypeOutput() IngressStartAtTypeOutput
	ToIngressStartAtTypeOutputWithContext(context.Context) IngressStartAtTypeOutput
}

var ingressStartAtTypePtrType = reflect.TypeOf((**IngressStartAtType)(nil)).Elem()

type IngressStartAtTypePtrInput interface {
	pulumi.Input

	ToIngressStartAtTypePtrOutput() IngressStartAtTypePtrOutput
	ToIngressStartAtTypePtrOutputWithContext(context.Context) IngressStartAtTypePtrOutput
}

type ingressStartAtTypePtr string

func IngressStartAtTypePtr(v string) IngressStartAtTypePtrInput {
	return (*ingressStartAtTypePtr)(&v)
}

func (*ingressStartAtTypePtr) ElementType() reflect.Type {
	return ingressStartAtTypePtrType
}

func (in *ingressStartAtTypePtr) ToIngressStartAtTypePtrOutput() IngressStartAtTypePtrOutput {
	return pulumi.ToOutput(in).(IngressStartAtTypePtrOutput)
}

func (in *ingressStartAtTypePtr) ToIngressStartAtTypePtrOutputWithContext(ctx context.Context) IngressStartAtTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IngressStartAtTypePtrOutput)
}

type LocalTimestampFormat string

const (
	LocalTimestampFormatEmbedded = LocalTimestampFormat("Embedded")
)

func (LocalTimestampFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalTimestampFormat)(nil)).Elem()
}

func (e LocalTimestampFormat) ToLocalTimestampFormatOutput() LocalTimestampFormatOutput {
	return pulumi.ToOutput(e).(LocalTimestampFormatOutput)
}

func (e LocalTimestampFormat) ToLocalTimestampFormatOutputWithContext(ctx context.Context) LocalTimestampFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LocalTimestampFormatOutput)
}

func (e LocalTimestampFormat) ToLocalTimestampFormatPtrOutput() LocalTimestampFormatPtrOutput {
	return e.ToLocalTimestampFormatPtrOutputWithContext(context.Background())
}

func (e LocalTimestampFormat) ToLocalTimestampFormatPtrOutputWithContext(ctx context.Context) LocalTimestampFormatPtrOutput {
	return LocalTimestampFormat(e).ToLocalTimestampFormatOutputWithContext(ctx).ToLocalTimestampFormatPtrOutputWithContext(ctx)
}

func (e LocalTimestampFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LocalTimestampFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LocalTimestampFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LocalTimestampFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LocalTimestampFormatOutput struct{ *pulumi.OutputState }

func (LocalTimestampFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalTimestampFormat)(nil)).Elem()
}

func (o LocalTimestampFormatOutput) ToLocalTimestampFormatOutput() LocalTimestampFormatOutput {
	return o
}

func (o LocalTimestampFormatOutput) ToLocalTimestampFormatOutputWithContext(ctx context.Context) LocalTimestampFormatOutput {
	return o
}

func (o LocalTimestampFormatOutput) ToLocalTimestampFormatPtrOutput() LocalTimestampFormatPtrOutput {
	return o.ToLocalTimestampFormatPtrOutputWithContext(context.Background())
}

func (o LocalTimestampFormatOutput) ToLocalTimestampFormatPtrOutputWithContext(ctx context.Context) LocalTimestampFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalTimestampFormat) *LocalTimestampFormat {
		return &v
	}).(LocalTimestampFormatPtrOutput)
}

func (o LocalTimestampFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LocalTimestampFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LocalTimestampFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LocalTimestampFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LocalTimestampFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LocalTimestampFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LocalTimestampFormatPtrOutput struct{ *pulumi.OutputState }

func (LocalTimestampFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTimestampFormat)(nil)).Elem()
}

func (o LocalTimestampFormatPtrOutput) ToLocalTimestampFormatPtrOutput() LocalTimestampFormatPtrOutput {
	return o
}

func (o LocalTimestampFormatPtrOutput) ToLocalTimestampFormatPtrOutputWithContext(ctx context.Context) LocalTimestampFormatPtrOutput {
	return o
}

func (o LocalTimestampFormatPtrOutput) Elem() LocalTimestampFormatOutput {
	return o.ApplyT(func(v *LocalTimestampFormat) LocalTimestampFormat {
		if v != nil {
			return *v
		}
		var ret LocalTimestampFormat
		return ret
	}).(LocalTimestampFormatOutput)
}

func (o LocalTimestampFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LocalTimestampFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LocalTimestampFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LocalTimestampFormatInput interface {
	pulumi.Input

	ToLocalTimestampFormatOutput() LocalTimestampFormatOutput
	ToLocalTimestampFormatOutputWithContext(context.Context) LocalTimestampFormatOutput
}

var localTimestampFormatPtrType = reflect.TypeOf((**LocalTimestampFormat)(nil)).Elem()

type LocalTimestampFormatPtrInput interface {
	pulumi.Input

	ToLocalTimestampFormatPtrOutput() LocalTimestampFormatPtrOutput
	ToLocalTimestampFormatPtrOutputWithContext(context.Context) LocalTimestampFormatPtrOutput
}

type localTimestampFormatPtr string

func LocalTimestampFormatPtr(v string) LocalTimestampFormatPtrInput {
	return (*localTimestampFormatPtr)(&v)
}

func (*localTimestampFormatPtr) ElementType() reflect.Type {
	return localTimestampFormatPtrType
}

func (in *localTimestampFormatPtr) ToLocalTimestampFormatPtrOutput() LocalTimestampFormatPtrOutput {
	return pulumi.ToOutput(in).(LocalTimestampFormatPtrOutput)
}

func (in *localTimestampFormatPtr) ToLocalTimestampFormatPtrOutputWithContext(ctx context.Context) LocalTimestampFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LocalTimestampFormatPtrOutput)
}

type PropertyType string

const (
	PropertyTypeString = PropertyType("String")
)

func (PropertyType) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyType)(nil)).Elem()
}

func (e PropertyType) ToPropertyTypeOutput() PropertyTypeOutput {
	return pulumi.ToOutput(e).(PropertyTypeOutput)
}

func (e PropertyType) ToPropertyTypeOutputWithContext(ctx context.Context) PropertyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PropertyTypeOutput)
}

func (e PropertyType) ToPropertyTypePtrOutput() PropertyTypePtrOutput {
	return e.ToPropertyTypePtrOutputWithContext(context.Background())
}

func (e PropertyType) ToPropertyTypePtrOutputWithContext(ctx context.Context) PropertyTypePtrOutput {
	return PropertyType(e).ToPropertyTypeOutputWithContext(ctx).ToPropertyTypePtrOutputWithContext(ctx)
}

func (e PropertyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PropertyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PropertyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PropertyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PropertyTypeOutput struct{ *pulumi.OutputState }

func (PropertyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyType)(nil)).Elem()
}

func (o PropertyTypeOutput) ToPropertyTypeOutput() PropertyTypeOutput {
	return o
}

func (o PropertyTypeOutput) ToPropertyTypeOutputWithContext(ctx context.Context) PropertyTypeOutput {
	return o
}

func (o PropertyTypeOutput) ToPropertyTypePtrOutput() PropertyTypePtrOutput {
	return o.ToPropertyTypePtrOutputWithContext(context.Background())
}

func (o PropertyTypeOutput) ToPropertyTypePtrOutputWithContext(ctx context.Context) PropertyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PropertyType) *PropertyType {
		return &v
	}).(PropertyTypePtrOutput)
}

func (o PropertyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PropertyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PropertyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PropertyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PropertyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PropertyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PropertyTypePtrOutput struct{ *pulumi.OutputState }

func (PropertyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyType)(nil)).Elem()
}

func (o PropertyTypePtrOutput) ToPropertyTypePtrOutput() PropertyTypePtrOutput {
	return o
}

func (o PropertyTypePtrOutput) ToPropertyTypePtrOutputWithContext(ctx context.Context) PropertyTypePtrOutput {
	return o
}

func (o PropertyTypePtrOutput) Elem() PropertyTypeOutput {
	return o.ApplyT(func(v *PropertyType) PropertyType {
		if v != nil {
			return *v
		}
		var ret PropertyType
		return ret
	}).(PropertyTypeOutput)
}

func (o PropertyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PropertyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PropertyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PropertyTypeInput interface {
	pulumi.Input

	ToPropertyTypeOutput() PropertyTypeOutput
	ToPropertyTypeOutputWithContext(context.Context) PropertyTypeOutput
}

var propertyTypePtrType = reflect.TypeOf((**PropertyType)(nil)).Elem()

type PropertyTypePtrInput interface {
	pulumi.Input

	ToPropertyTypePtrOutput() PropertyTypePtrOutput
	ToPropertyTypePtrOutputWithContext(context.Context) PropertyTypePtrOutput
}

type propertyTypePtr string

func PropertyTypePtr(v string) PropertyTypePtrInput {
	return (*propertyTypePtr)(&v)
}

func (*propertyTypePtr) ElementType() reflect.Type {
	return propertyTypePtrType
}

func (in *propertyTypePtr) ToPropertyTypePtrOutput() PropertyTypePtrOutput {
	return pulumi.ToOutput(in).(PropertyTypePtrOutput)
}

func (in *propertyTypePtr) ToPropertyTypePtrOutputWithContext(ctx context.Context) PropertyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PropertyTypePtrOutput)
}

type ReferenceDataKeyPropertyType string

const (
	ReferenceDataKeyPropertyTypeString   = ReferenceDataKeyPropertyType("String")
	ReferenceDataKeyPropertyTypeDouble   = ReferenceDataKeyPropertyType("Double")
	ReferenceDataKeyPropertyTypeBool     = ReferenceDataKeyPropertyType("Bool")
	ReferenceDataKeyPropertyTypeDateTime = ReferenceDataKeyPropertyType("DateTime")
)

func (ReferenceDataKeyPropertyType) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataKeyPropertyType)(nil)).Elem()
}

func (e ReferenceDataKeyPropertyType) ToReferenceDataKeyPropertyTypeOutput() ReferenceDataKeyPropertyTypeOutput {
	return pulumi.ToOutput(e).(ReferenceDataKeyPropertyTypeOutput)
}

func (e ReferenceDataKeyPropertyType) ToReferenceDataKeyPropertyTypeOutputWithContext(ctx context.Context) ReferenceDataKeyPropertyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReferenceDataKeyPropertyTypeOutput)
}

func (e ReferenceDataKeyPropertyType) ToReferenceDataKeyPropertyTypePtrOutput() ReferenceDataKeyPropertyTypePtrOutput {
	return e.ToReferenceDataKeyPropertyTypePtrOutputWithContext(context.Background())
}

func (e ReferenceDataKeyPropertyType) ToReferenceDataKeyPropertyTypePtrOutputWithContext(ctx context.Context) ReferenceDataKeyPropertyTypePtrOutput {
	return ReferenceDataKeyPropertyType(e).ToReferenceDataKeyPropertyTypeOutputWithContext(ctx).ToReferenceDataKeyPropertyTypePtrOutputWithContext(ctx)
}

func (e ReferenceDataKeyPropertyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReferenceDataKeyPropertyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReferenceDataKeyPropertyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReferenceDataKeyPropertyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReferenceDataKeyPropertyTypeOutput struct{ *pulumi.OutputState }

func (ReferenceDataKeyPropertyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataKeyPropertyType)(nil)).Elem()
}

func (o ReferenceDataKeyPropertyTypeOutput) ToReferenceDataKeyPropertyTypeOutput() ReferenceDataKeyPropertyTypeOutput {
	return o
}

func (o ReferenceDataKeyPropertyTypeOutput) ToReferenceDataKeyPropertyTypeOutputWithContext(ctx context.Context) ReferenceDataKeyPropertyTypeOutput {
	return o
}

func (o ReferenceDataKeyPropertyTypeOutput) ToReferenceDataKeyPropertyTypePtrOutput() ReferenceDataKeyPropertyTypePtrOutput {
	return o.ToReferenceDataKeyPropertyTypePtrOutputWithContext(context.Background())
}

func (o ReferenceDataKeyPropertyTypeOutput) ToReferenceDataKeyPropertyTypePtrOutputWithContext(ctx context.Context) ReferenceDataKeyPropertyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReferenceDataKeyPropertyType) *ReferenceDataKeyPropertyType {
		return &v
	}).(ReferenceDataKeyPropertyTypePtrOutput)
}

func (o ReferenceDataKeyPropertyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReferenceDataKeyPropertyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReferenceDataKeyPropertyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReferenceDataKeyPropertyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReferenceDataKeyPropertyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReferenceDataKeyPropertyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReferenceDataKeyPropertyTypePtrOutput struct{ *pulumi.OutputState }

func (ReferenceDataKeyPropertyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceDataKeyPropertyType)(nil)).Elem()
}

func (o ReferenceDataKeyPropertyTypePtrOutput) ToReferenceDataKeyPropertyTypePtrOutput() ReferenceDataKeyPropertyTypePtrOutput {
	return o
}

func (o ReferenceDataKeyPropertyTypePtrOutput) ToReferenceDataKeyPropertyTypePtrOutputWithContext(ctx context.Context) ReferenceDataKeyPropertyTypePtrOutput {
	return o
}

func (o ReferenceDataKeyPropertyTypePtrOutput) Elem() ReferenceDataKeyPropertyTypeOutput {
	return o.ApplyT(func(v *ReferenceDataKeyPropertyType) ReferenceDataKeyPropertyType {
		if v != nil {
			return *v
		}
		var ret ReferenceDataKeyPropertyType
		return ret
	}).(ReferenceDataKeyPropertyTypeOutput)
}

func (o ReferenceDataKeyPropertyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReferenceDataKeyPropertyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReferenceDataKeyPropertyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ReferenceDataKeyPropertyTypeInput interface {
	pulumi.Input

	ToReferenceDataKeyPropertyTypeOutput() ReferenceDataKeyPropertyTypeOutput
	ToReferenceDataKeyPropertyTypeOutputWithContext(context.Context) ReferenceDataKeyPropertyTypeOutput
}

var referenceDataKeyPropertyTypePtrType = reflect.TypeOf((**ReferenceDataKeyPropertyType)(nil)).Elem()

type ReferenceDataKeyPropertyTypePtrInput interface {
	pulumi.Input

	ToReferenceDataKeyPropertyTypePtrOutput() ReferenceDataKeyPropertyTypePtrOutput
	ToReferenceDataKeyPropertyTypePtrOutputWithContext(context.Context) ReferenceDataKeyPropertyTypePtrOutput
}

type referenceDataKeyPropertyTypePtr string

func ReferenceDataKeyPropertyTypePtr(v string) ReferenceDataKeyPropertyTypePtrInput {
	return (*referenceDataKeyPropertyTypePtr)(&v)
}

func (*referenceDataKeyPropertyTypePtr) ElementType() reflect.Type {
	return referenceDataKeyPropertyTypePtrType
}

func (in *referenceDataKeyPropertyTypePtr) ToReferenceDataKeyPropertyTypePtrOutput() ReferenceDataKeyPropertyTypePtrOutput {
	return pulumi.ToOutput(in).(ReferenceDataKeyPropertyTypePtrOutput)
}

func (in *referenceDataKeyPropertyTypePtr) ToReferenceDataKeyPropertyTypePtrOutputWithContext(ctx context.Context) ReferenceDataKeyPropertyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReferenceDataKeyPropertyTypePtrOutput)
}

type SkuName string

const (
	SkuNameS1 = SkuName("S1")
	SkuNameS2 = SkuName("S2")
	SkuNameP1 = SkuName("P1")
	SkuNameL1 = SkuName("L1")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

type StorageLimitExceededBehavior string

const (
	StorageLimitExceededBehaviorPurgeOldData = StorageLimitExceededBehavior("PurgeOldData")
	StorageLimitExceededBehaviorPauseIngress = StorageLimitExceededBehavior("PauseIngress")
)

func (StorageLimitExceededBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageLimitExceededBehavior)(nil)).Elem()
}

func (e StorageLimitExceededBehavior) ToStorageLimitExceededBehaviorOutput() StorageLimitExceededBehaviorOutput {
	return pulumi.ToOutput(e).(StorageLimitExceededBehaviorOutput)
}

func (e StorageLimitExceededBehavior) ToStorageLimitExceededBehaviorOutputWithContext(ctx context.Context) StorageLimitExceededBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageLimitExceededBehaviorOutput)
}

func (e StorageLimitExceededBehavior) ToStorageLimitExceededBehaviorPtrOutput() StorageLimitExceededBehaviorPtrOutput {
	return e.ToStorageLimitExceededBehaviorPtrOutputWithContext(context.Background())
}

func (e StorageLimitExceededBehavior) ToStorageLimitExceededBehaviorPtrOutputWithContext(ctx context.Context) StorageLimitExceededBehaviorPtrOutput {
	return StorageLimitExceededBehavior(e).ToStorageLimitExceededBehaviorOutputWithContext(ctx).ToStorageLimitExceededBehaviorPtrOutputWithContext(ctx)
}

func (e StorageLimitExceededBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageLimitExceededBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageLimitExceededBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageLimitExceededBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageLimitExceededBehaviorOutput struct{ *pulumi.OutputState }

func (StorageLimitExceededBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageLimitExceededBehavior)(nil)).Elem()
}

func (o StorageLimitExceededBehaviorOutput) ToStorageLimitExceededBehaviorOutput() StorageLimitExceededBehaviorOutput {
	return o
}

func (o StorageLimitExceededBehaviorOutput) ToStorageLimitExceededBehaviorOutputWithContext(ctx context.Context) StorageLimitExceededBehaviorOutput {
	return o
}

func (o StorageLimitExceededBehaviorOutput) ToStorageLimitExceededBehaviorPtrOutput() StorageLimitExceededBehaviorPtrOutput {
	return o.ToStorageLimitExceededBehaviorPtrOutputWithContext(context.Background())
}

func (o StorageLimitExceededBehaviorOutput) ToStorageLimitExceededBehaviorPtrOutputWithContext(ctx context.Context) StorageLimitExceededBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageLimitExceededBehavior) *StorageLimitExceededBehavior {
		return &v
	}).(StorageLimitExceededBehaviorPtrOutput)
}

func (o StorageLimitExceededBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageLimitExceededBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageLimitExceededBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageLimitExceededBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageLimitExceededBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageLimitExceededBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageLimitExceededBehaviorPtrOutput struct{ *pulumi.OutputState }

func (StorageLimitExceededBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageLimitExceededBehavior)(nil)).Elem()
}

func (o StorageLimitExceededBehaviorPtrOutput) ToStorageLimitExceededBehaviorPtrOutput() StorageLimitExceededBehaviorPtrOutput {
	return o
}

func (o StorageLimitExceededBehaviorPtrOutput) ToStorageLimitExceededBehaviorPtrOutputWithContext(ctx context.Context) StorageLimitExceededBehaviorPtrOutput {
	return o
}

func (o StorageLimitExceededBehaviorPtrOutput) Elem() StorageLimitExceededBehaviorOutput {
	return o.ApplyT(func(v *StorageLimitExceededBehavior) StorageLimitExceededBehavior {
		if v != nil {
			return *v
		}
		var ret StorageLimitExceededBehavior
		return ret
	}).(StorageLimitExceededBehaviorOutput)
}

func (o StorageLimitExceededBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageLimitExceededBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageLimitExceededBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageLimitExceededBehaviorInput interface {
	pulumi.Input

	ToStorageLimitExceededBehaviorOutput() StorageLimitExceededBehaviorOutput
	ToStorageLimitExceededBehaviorOutputWithContext(context.Context) StorageLimitExceededBehaviorOutput
}

var storageLimitExceededBehaviorPtrType = reflect.TypeOf((**StorageLimitExceededBehavior)(nil)).Elem()

type StorageLimitExceededBehaviorPtrInput interface {
	pulumi.Input

	ToStorageLimitExceededBehaviorPtrOutput() StorageLimitExceededBehaviorPtrOutput
	ToStorageLimitExceededBehaviorPtrOutputWithContext(context.Context) StorageLimitExceededBehaviorPtrOutput
}

type storageLimitExceededBehaviorPtr string

func StorageLimitExceededBehaviorPtr(v string) StorageLimitExceededBehaviorPtrInput {
	return (*storageLimitExceededBehaviorPtr)(&v)
}

func (*storageLimitExceededBehaviorPtr) ElementType() reflect.Type {
	return storageLimitExceededBehaviorPtrType
}

func (in *storageLimitExceededBehaviorPtr) ToStorageLimitExceededBehaviorPtrOutput() StorageLimitExceededBehaviorPtrOutput {
	return pulumi.ToOutput(in).(StorageLimitExceededBehaviorPtrOutput)
}

func (in *storageLimitExceededBehaviorPtr) ToStorageLimitExceededBehaviorPtrOutputWithContext(ctx context.Context) StorageLimitExceededBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageLimitExceededBehaviorPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyRoleInput)(nil)).Elem(), AccessPolicyRole("Reader"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyRolePtrInput)(nil)).Elem(), AccessPolicyRole("Reader"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataStringComparisonBehaviorInput)(nil)).Elem(), DataStringComparisonBehavior("Ordinal"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataStringComparisonBehaviorPtrInput)(nil)).Elem(), DataStringComparisonBehavior("Ordinal"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentKindInput)(nil)).Elem(), EnvironmentKind("Gen1"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentKindPtrInput)(nil)).Elem(), EnvironmentKind("Gen1"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventSourceKindInput)(nil)).Elem(), EventSourceKind("Microsoft.EventHub"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventSourceKindPtrInput)(nil)).Elem(), EventSourceKind("Microsoft.EventHub"))
	pulumi.RegisterInputType(reflect.TypeOf((*IngressStartAtTypeInput)(nil)).Elem(), IngressStartAtType("EarliestAvailable"))
	pulumi.RegisterInputType(reflect.TypeOf((*IngressStartAtTypePtrInput)(nil)).Elem(), IngressStartAtType("EarliestAvailable"))
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTimestampFormatInput)(nil)).Elem(), LocalTimestampFormat("Embedded"))
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTimestampFormatPtrInput)(nil)).Elem(), LocalTimestampFormat("Embedded"))
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyTypeInput)(nil)).Elem(), PropertyType("String"))
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyTypePtrInput)(nil)).Elem(), PropertyType("String"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceDataKeyPropertyTypeInput)(nil)).Elem(), ReferenceDataKeyPropertyType("String"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceDataKeyPropertyTypePtrInput)(nil)).Elem(), ReferenceDataKeyPropertyType("String"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNameInput)(nil)).Elem(), SkuName("S1"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNamePtrInput)(nil)).Elem(), SkuName("S1"))
	pulumi.RegisterInputType(reflect.TypeOf((*StorageLimitExceededBehaviorInput)(nil)).Elem(), StorageLimitExceededBehavior("PurgeOldData"))
	pulumi.RegisterInputType(reflect.TypeOf((*StorageLimitExceededBehaviorPtrInput)(nil)).Elem(), StorageLimitExceededBehavior("PurgeOldData"))
	pulumi.RegisterOutputType(AccessPolicyRoleOutput{})
	pulumi.RegisterOutputType(AccessPolicyRolePtrOutput{})
	pulumi.RegisterOutputType(DataStringComparisonBehaviorOutput{})
	pulumi.RegisterOutputType(DataStringComparisonBehaviorPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentKindOutput{})
	pulumi.RegisterOutputType(EnvironmentKindPtrOutput{})
	pulumi.RegisterOutputType(EventSourceKindOutput{})
	pulumi.RegisterOutputType(EventSourceKindPtrOutput{})
	pulumi.RegisterOutputType(IngressStartAtTypeOutput{})
	pulumi.RegisterOutputType(IngressStartAtTypePtrOutput{})
	pulumi.RegisterOutputType(LocalTimestampFormatOutput{})
	pulumi.RegisterOutputType(LocalTimestampFormatPtrOutput{})
	pulumi.RegisterOutputType(PropertyTypeOutput{})
	pulumi.RegisterOutputType(PropertyTypePtrOutput{})
	pulumi.RegisterOutputType(ReferenceDataKeyPropertyTypeOutput{})
	pulumi.RegisterOutputType(ReferenceDataKeyPropertyTypePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(StorageLimitExceededBehaviorOutput{})
	pulumi.RegisterOutputType(StorageLimitExceededBehaviorPtrOutput{})
}
