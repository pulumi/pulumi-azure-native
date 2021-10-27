


package delegatednetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OrchestratorKind string

const (
	OrchestratorKindKubernetes = OrchestratorKind("Kubernetes")
)

func (OrchestratorKind) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratorKind)(nil)).Elem()
}

func (e OrchestratorKind) ToOrchestratorKindOutput() OrchestratorKindOutput {
	return pulumi.ToOutput(e).(OrchestratorKindOutput)
}

func (e OrchestratorKind) ToOrchestratorKindOutputWithContext(ctx context.Context) OrchestratorKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OrchestratorKindOutput)
}

func (e OrchestratorKind) ToOrchestratorKindPtrOutput() OrchestratorKindPtrOutput {
	return e.ToOrchestratorKindPtrOutputWithContext(context.Background())
}

func (e OrchestratorKind) ToOrchestratorKindPtrOutputWithContext(ctx context.Context) OrchestratorKindPtrOutput {
	return OrchestratorKind(e).ToOrchestratorKindOutputWithContext(ctx).ToOrchestratorKindPtrOutputWithContext(ctx)
}

func (e OrchestratorKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrchestratorKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrchestratorKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OrchestratorKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OrchestratorKindOutput struct{ *pulumi.OutputState }

func (OrchestratorKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratorKind)(nil)).Elem()
}

func (o OrchestratorKindOutput) ToOrchestratorKindOutput() OrchestratorKindOutput {
	return o
}

func (o OrchestratorKindOutput) ToOrchestratorKindOutputWithContext(ctx context.Context) OrchestratorKindOutput {
	return o
}

func (o OrchestratorKindOutput) ToOrchestratorKindPtrOutput() OrchestratorKindPtrOutput {
	return o.ToOrchestratorKindPtrOutputWithContext(context.Background())
}

func (o OrchestratorKindOutput) ToOrchestratorKindPtrOutputWithContext(ctx context.Context) OrchestratorKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrchestratorKind) *OrchestratorKind {
		return &v
	}).(OrchestratorKindPtrOutput)
}

func (o OrchestratorKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OrchestratorKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OrchestratorKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OrchestratorKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OrchestratorKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OrchestratorKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OrchestratorKindPtrOutput struct{ *pulumi.OutputState }

func (OrchestratorKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratorKind)(nil)).Elem()
}

func (o OrchestratorKindPtrOutput) ToOrchestratorKindPtrOutput() OrchestratorKindPtrOutput {
	return o
}

func (o OrchestratorKindPtrOutput) ToOrchestratorKindPtrOutputWithContext(ctx context.Context) OrchestratorKindPtrOutput {
	return o
}

func (o OrchestratorKindPtrOutput) Elem() OrchestratorKindOutput {
	return o.ApplyT(func(v *OrchestratorKind) OrchestratorKind {
		if v != nil {
			return *v
		}
		var ret OrchestratorKind
		return ret
	}).(OrchestratorKindOutput)
}

func (o OrchestratorKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OrchestratorKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OrchestratorKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OrchestratorKindInput interface {
	pulumi.Input

	ToOrchestratorKindOutput() OrchestratorKindOutput
	ToOrchestratorKindOutputWithContext(context.Context) OrchestratorKindOutput
}

var orchestratorKindPtrType = reflect.TypeOf((**OrchestratorKind)(nil)).Elem()

type OrchestratorKindPtrInput interface {
	pulumi.Input

	ToOrchestratorKindPtrOutput() OrchestratorKindPtrOutput
	ToOrchestratorKindPtrOutputWithContext(context.Context) OrchestratorKindPtrOutput
}

type orchestratorKindPtr string

func OrchestratorKindPtr(v string) OrchestratorKindPtrInput {
	return (*orchestratorKindPtr)(&v)
}

func (*orchestratorKindPtr) ElementType() reflect.Type {
	return orchestratorKindPtrType
}

func (in *orchestratorKindPtr) ToOrchestratorKindPtrOutput() OrchestratorKindPtrOutput {
	return pulumi.ToOutput(in).(OrchestratorKindPtrOutput)
}

func (in *orchestratorKindPtr) ToOrchestratorKindPtrOutputWithContext(ctx context.Context) OrchestratorKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OrchestratorKindPtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrchestratorKindInput)(nil)).Elem(), OrchestratorKind("Kubernetes"))
	pulumi.RegisterInputType(reflect.TypeOf((*OrchestratorKindPtrInput)(nil)).Elem(), OrchestratorKind("Kubernetes"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterOutputType(OrchestratorKindOutput{})
	pulumi.RegisterOutputType(OrchestratorKindPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
