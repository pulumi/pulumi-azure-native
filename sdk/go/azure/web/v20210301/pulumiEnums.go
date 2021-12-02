


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActiveRevisionsMode string

const (
	ActiveRevisionsModeMultiple = ActiveRevisionsMode("multiple")
	ActiveRevisionsModeSingle   = ActiveRevisionsMode("single")
)

func (ActiveRevisionsMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveRevisionsMode)(nil)).Elem()
}

func (e ActiveRevisionsMode) ToActiveRevisionsModeOutput() ActiveRevisionsModeOutput {
	return pulumi.ToOutput(e).(ActiveRevisionsModeOutput)
}

func (e ActiveRevisionsMode) ToActiveRevisionsModeOutputWithContext(ctx context.Context) ActiveRevisionsModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActiveRevisionsModeOutput)
}

func (e ActiveRevisionsMode) ToActiveRevisionsModePtrOutput() ActiveRevisionsModePtrOutput {
	return e.ToActiveRevisionsModePtrOutputWithContext(context.Background())
}

func (e ActiveRevisionsMode) ToActiveRevisionsModePtrOutputWithContext(ctx context.Context) ActiveRevisionsModePtrOutput {
	return ActiveRevisionsMode(e).ToActiveRevisionsModeOutputWithContext(ctx).ToActiveRevisionsModePtrOutputWithContext(ctx)
}

func (e ActiveRevisionsMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActiveRevisionsMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActiveRevisionsMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ActiveRevisionsMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActiveRevisionsModeOutput struct{ *pulumi.OutputState }

func (ActiveRevisionsModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveRevisionsMode)(nil)).Elem()
}

func (o ActiveRevisionsModeOutput) ToActiveRevisionsModeOutput() ActiveRevisionsModeOutput {
	return o
}

func (o ActiveRevisionsModeOutput) ToActiveRevisionsModeOutputWithContext(ctx context.Context) ActiveRevisionsModeOutput {
	return o
}

func (o ActiveRevisionsModeOutput) ToActiveRevisionsModePtrOutput() ActiveRevisionsModePtrOutput {
	return o.ToActiveRevisionsModePtrOutputWithContext(context.Background())
}

func (o ActiveRevisionsModeOutput) ToActiveRevisionsModePtrOutputWithContext(ctx context.Context) ActiveRevisionsModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActiveRevisionsMode) *ActiveRevisionsMode {
		return &v
	}).(ActiveRevisionsModePtrOutput)
}

func (o ActiveRevisionsModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActiveRevisionsModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActiveRevisionsMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActiveRevisionsModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActiveRevisionsModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActiveRevisionsMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActiveRevisionsModePtrOutput struct{ *pulumi.OutputState }

func (ActiveRevisionsModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveRevisionsMode)(nil)).Elem()
}

func (o ActiveRevisionsModePtrOutput) ToActiveRevisionsModePtrOutput() ActiveRevisionsModePtrOutput {
	return o
}

func (o ActiveRevisionsModePtrOutput) ToActiveRevisionsModePtrOutputWithContext(ctx context.Context) ActiveRevisionsModePtrOutput {
	return o
}

func (o ActiveRevisionsModePtrOutput) Elem() ActiveRevisionsModeOutput {
	return o.ApplyT(func(v *ActiveRevisionsMode) ActiveRevisionsMode {
		if v != nil {
			return *v
		}
		var ret ActiveRevisionsMode
		return ret
	}).(ActiveRevisionsModeOutput)
}

func (o ActiveRevisionsModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActiveRevisionsModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ActiveRevisionsMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ActiveRevisionsModeInput interface {
	pulumi.Input

	ToActiveRevisionsModeOutput() ActiveRevisionsModeOutput
	ToActiveRevisionsModeOutputWithContext(context.Context) ActiveRevisionsModeOutput
}

var activeRevisionsModePtrType = reflect.TypeOf((**ActiveRevisionsMode)(nil)).Elem()

type ActiveRevisionsModePtrInput interface {
	pulumi.Input

	ToActiveRevisionsModePtrOutput() ActiveRevisionsModePtrOutput
	ToActiveRevisionsModePtrOutputWithContext(context.Context) ActiveRevisionsModePtrOutput
}

type activeRevisionsModePtr string

func ActiveRevisionsModePtr(v string) ActiveRevisionsModePtrInput {
	return (*activeRevisionsModePtr)(&v)
}

func (*activeRevisionsModePtr) ElementType() reflect.Type {
	return activeRevisionsModePtrType
}

func (in *activeRevisionsModePtr) ToActiveRevisionsModePtrOutput() ActiveRevisionsModePtrOutput {
	return pulumi.ToOutput(in).(ActiveRevisionsModePtrOutput)
}

func (in *activeRevisionsModePtr) ToActiveRevisionsModePtrOutputWithContext(ctx context.Context) ActiveRevisionsModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActiveRevisionsModePtrOutput)
}

type FrontEndServiceType string

const (
	FrontEndServiceTypeNodePort     = FrontEndServiceType("NodePort")
	FrontEndServiceTypeLoadBalancer = FrontEndServiceType("LoadBalancer")
)

func (FrontEndServiceType) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontEndServiceType)(nil)).Elem()
}

func (e FrontEndServiceType) ToFrontEndServiceTypeOutput() FrontEndServiceTypeOutput {
	return pulumi.ToOutput(e).(FrontEndServiceTypeOutput)
}

func (e FrontEndServiceType) ToFrontEndServiceTypeOutputWithContext(ctx context.Context) FrontEndServiceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FrontEndServiceTypeOutput)
}

func (e FrontEndServiceType) ToFrontEndServiceTypePtrOutput() FrontEndServiceTypePtrOutput {
	return e.ToFrontEndServiceTypePtrOutputWithContext(context.Background())
}

func (e FrontEndServiceType) ToFrontEndServiceTypePtrOutputWithContext(ctx context.Context) FrontEndServiceTypePtrOutput {
	return FrontEndServiceType(e).ToFrontEndServiceTypeOutputWithContext(ctx).ToFrontEndServiceTypePtrOutputWithContext(ctx)
}

func (e FrontEndServiceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrontEndServiceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrontEndServiceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FrontEndServiceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FrontEndServiceTypeOutput struct{ *pulumi.OutputState }

func (FrontEndServiceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontEndServiceType)(nil)).Elem()
}

func (o FrontEndServiceTypeOutput) ToFrontEndServiceTypeOutput() FrontEndServiceTypeOutput {
	return o
}

func (o FrontEndServiceTypeOutput) ToFrontEndServiceTypeOutputWithContext(ctx context.Context) FrontEndServiceTypeOutput {
	return o
}

func (o FrontEndServiceTypeOutput) ToFrontEndServiceTypePtrOutput() FrontEndServiceTypePtrOutput {
	return o.ToFrontEndServiceTypePtrOutputWithContext(context.Background())
}

func (o FrontEndServiceTypeOutput) ToFrontEndServiceTypePtrOutputWithContext(ctx context.Context) FrontEndServiceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontEndServiceType) *FrontEndServiceType {
		return &v
	}).(FrontEndServiceTypePtrOutput)
}

func (o FrontEndServiceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FrontEndServiceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrontEndServiceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FrontEndServiceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrontEndServiceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrontEndServiceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FrontEndServiceTypePtrOutput struct{ *pulumi.OutputState }

func (FrontEndServiceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontEndServiceType)(nil)).Elem()
}

func (o FrontEndServiceTypePtrOutput) ToFrontEndServiceTypePtrOutput() FrontEndServiceTypePtrOutput {
	return o
}

func (o FrontEndServiceTypePtrOutput) ToFrontEndServiceTypePtrOutputWithContext(ctx context.Context) FrontEndServiceTypePtrOutput {
	return o
}

func (o FrontEndServiceTypePtrOutput) Elem() FrontEndServiceTypeOutput {
	return o.ApplyT(func(v *FrontEndServiceType) FrontEndServiceType {
		if v != nil {
			return *v
		}
		var ret FrontEndServiceType
		return ret
	}).(FrontEndServiceTypeOutput)
}

func (o FrontEndServiceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrontEndServiceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FrontEndServiceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FrontEndServiceTypeInput interface {
	pulumi.Input

	ToFrontEndServiceTypeOutput() FrontEndServiceTypeOutput
	ToFrontEndServiceTypeOutputWithContext(context.Context) FrontEndServiceTypeOutput
}

var frontEndServiceTypePtrType = reflect.TypeOf((**FrontEndServiceType)(nil)).Elem()

type FrontEndServiceTypePtrInput interface {
	pulumi.Input

	ToFrontEndServiceTypePtrOutput() FrontEndServiceTypePtrOutput
	ToFrontEndServiceTypePtrOutputWithContext(context.Context) FrontEndServiceTypePtrOutput
}

type frontEndServiceTypePtr string

func FrontEndServiceTypePtr(v string) FrontEndServiceTypePtrInput {
	return (*frontEndServiceTypePtr)(&v)
}

func (*frontEndServiceTypePtr) ElementType() reflect.Type {
	return frontEndServiceTypePtrType
}

func (in *frontEndServiceTypePtr) ToFrontEndServiceTypePtrOutput() FrontEndServiceTypePtrOutput {
	return pulumi.ToOutput(in).(FrontEndServiceTypePtrOutput)
}

func (in *frontEndServiceTypePtr) ToFrontEndServiceTypePtrOutputWithContext(ctx context.Context) FrontEndServiceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FrontEndServiceTypePtrOutput)
}

type IngressTransportMethod string

const (
	IngressTransportMethodAuto  = IngressTransportMethod("auto")
	IngressTransportMethodHttp  = IngressTransportMethod("http")
	IngressTransportMethodHttp2 = IngressTransportMethod("http2")
)

func (IngressTransportMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressTransportMethod)(nil)).Elem()
}

func (e IngressTransportMethod) ToIngressTransportMethodOutput() IngressTransportMethodOutput {
	return pulumi.ToOutput(e).(IngressTransportMethodOutput)
}

func (e IngressTransportMethod) ToIngressTransportMethodOutputWithContext(ctx context.Context) IngressTransportMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IngressTransportMethodOutput)
}

func (e IngressTransportMethod) ToIngressTransportMethodPtrOutput() IngressTransportMethodPtrOutput {
	return e.ToIngressTransportMethodPtrOutputWithContext(context.Background())
}

func (e IngressTransportMethod) ToIngressTransportMethodPtrOutputWithContext(ctx context.Context) IngressTransportMethodPtrOutput {
	return IngressTransportMethod(e).ToIngressTransportMethodOutputWithContext(ctx).ToIngressTransportMethodPtrOutputWithContext(ctx)
}

func (e IngressTransportMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IngressTransportMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IngressTransportMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IngressTransportMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IngressTransportMethodOutput struct{ *pulumi.OutputState }

func (IngressTransportMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressTransportMethod)(nil)).Elem()
}

func (o IngressTransportMethodOutput) ToIngressTransportMethodOutput() IngressTransportMethodOutput {
	return o
}

func (o IngressTransportMethodOutput) ToIngressTransportMethodOutputWithContext(ctx context.Context) IngressTransportMethodOutput {
	return o
}

func (o IngressTransportMethodOutput) ToIngressTransportMethodPtrOutput() IngressTransportMethodPtrOutput {
	return o.ToIngressTransportMethodPtrOutputWithContext(context.Background())
}

func (o IngressTransportMethodOutput) ToIngressTransportMethodPtrOutputWithContext(ctx context.Context) IngressTransportMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IngressTransportMethod) *IngressTransportMethod {
		return &v
	}).(IngressTransportMethodPtrOutput)
}

func (o IngressTransportMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IngressTransportMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IngressTransportMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IngressTransportMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IngressTransportMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IngressTransportMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IngressTransportMethodPtrOutput struct{ *pulumi.OutputState }

func (IngressTransportMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressTransportMethod)(nil)).Elem()
}

func (o IngressTransportMethodPtrOutput) ToIngressTransportMethodPtrOutput() IngressTransportMethodPtrOutput {
	return o
}

func (o IngressTransportMethodPtrOutput) ToIngressTransportMethodPtrOutputWithContext(ctx context.Context) IngressTransportMethodPtrOutput {
	return o
}

func (o IngressTransportMethodPtrOutput) Elem() IngressTransportMethodOutput {
	return o.ApplyT(func(v *IngressTransportMethod) IngressTransportMethod {
		if v != nil {
			return *v
		}
		var ret IngressTransportMethod
		return ret
	}).(IngressTransportMethodOutput)
}

func (o IngressTransportMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IngressTransportMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IngressTransportMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IngressTransportMethodInput interface {
	pulumi.Input

	ToIngressTransportMethodOutput() IngressTransportMethodOutput
	ToIngressTransportMethodOutputWithContext(context.Context) IngressTransportMethodOutput
}

var ingressTransportMethodPtrType = reflect.TypeOf((**IngressTransportMethod)(nil)).Elem()

type IngressTransportMethodPtrInput interface {
	pulumi.Input

	ToIngressTransportMethodPtrOutput() IngressTransportMethodPtrOutput
	ToIngressTransportMethodPtrOutputWithContext(context.Context) IngressTransportMethodPtrOutput
}

type ingressTransportMethodPtr string

func IngressTransportMethodPtr(v string) IngressTransportMethodPtrInput {
	return (*ingressTransportMethodPtr)(&v)
}

func (*ingressTransportMethodPtr) ElementType() reflect.Type {
	return ingressTransportMethodPtrType
}

func (in *ingressTransportMethodPtr) ToIngressTransportMethodPtrOutput() IngressTransportMethodPtrOutput {
	return pulumi.ToOutput(in).(IngressTransportMethodPtrOutput)
}

func (in *ingressTransportMethodPtr) ToIngressTransportMethodPtrOutputWithContext(ctx context.Context) IngressTransportMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IngressTransportMethodPtrOutput)
}

type StorageType string

const (
	StorageTypeLocalNode         = StorageType("LocalNode")
	StorageTypeNetworkFileSystem = StorageType("NetworkFileSystem")
)

func (StorageType) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageType)(nil)).Elem()
}

func (e StorageType) ToStorageTypeOutput() StorageTypeOutput {
	return pulumi.ToOutput(e).(StorageTypeOutput)
}

func (e StorageType) ToStorageTypeOutputWithContext(ctx context.Context) StorageTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageTypeOutput)
}

func (e StorageType) ToStorageTypePtrOutput() StorageTypePtrOutput {
	return e.ToStorageTypePtrOutputWithContext(context.Background())
}

func (e StorageType) ToStorageTypePtrOutputWithContext(ctx context.Context) StorageTypePtrOutput {
	return StorageType(e).ToStorageTypeOutputWithContext(ctx).ToStorageTypePtrOutputWithContext(ctx)
}

func (e StorageType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageTypeOutput struct{ *pulumi.OutputState }

func (StorageTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageType)(nil)).Elem()
}

func (o StorageTypeOutput) ToStorageTypeOutput() StorageTypeOutput {
	return o
}

func (o StorageTypeOutput) ToStorageTypeOutputWithContext(ctx context.Context) StorageTypeOutput {
	return o
}

func (o StorageTypeOutput) ToStorageTypePtrOutput() StorageTypePtrOutput {
	return o.ToStorageTypePtrOutputWithContext(context.Background())
}

func (o StorageTypeOutput) ToStorageTypePtrOutputWithContext(ctx context.Context) StorageTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageType) *StorageType {
		return &v
	}).(StorageTypePtrOutput)
}

func (o StorageTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageTypePtrOutput struct{ *pulumi.OutputState }

func (StorageTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageType)(nil)).Elem()
}

func (o StorageTypePtrOutput) ToStorageTypePtrOutput() StorageTypePtrOutput {
	return o
}

func (o StorageTypePtrOutput) ToStorageTypePtrOutputWithContext(ctx context.Context) StorageTypePtrOutput {
	return o
}

func (o StorageTypePtrOutput) Elem() StorageTypeOutput {
	return o.ApplyT(func(v *StorageType) StorageType {
		if v != nil {
			return *v
		}
		var ret StorageType
		return ret
	}).(StorageTypeOutput)
}

func (o StorageTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageTypeInput interface {
	pulumi.Input

	ToStorageTypeOutput() StorageTypeOutput
	ToStorageTypeOutputWithContext(context.Context) StorageTypeOutput
}

var storageTypePtrType = reflect.TypeOf((**StorageType)(nil)).Elem()

type StorageTypePtrInput interface {
	pulumi.Input

	ToStorageTypePtrOutput() StorageTypePtrOutput
	ToStorageTypePtrOutputWithContext(context.Context) StorageTypePtrOutput
}

type storageTypePtr string

func StorageTypePtr(v string) StorageTypePtrInput {
	return (*storageTypePtr)(&v)
}

func (*storageTypePtr) ElementType() reflect.Type {
	return storageTypePtrType
}

func (in *storageTypePtr) ToStorageTypePtrOutput() StorageTypePtrOutput {
	return pulumi.ToOutput(in).(StorageTypePtrOutput)
}

func (in *storageTypePtr) ToStorageTypePtrOutputWithContext(ctx context.Context) StorageTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveRevisionsModeOutput{})
	pulumi.RegisterOutputType(ActiveRevisionsModePtrOutput{})
	pulumi.RegisterOutputType(FrontEndServiceTypeOutput{})
	pulumi.RegisterOutputType(FrontEndServiceTypePtrOutput{})
	pulumi.RegisterOutputType(IngressTransportMethodOutput{})
	pulumi.RegisterOutputType(IngressTransportMethodPtrOutput{})
	pulumi.RegisterOutputType(StorageTypeOutput{})
	pulumi.RegisterOutputType(StorageTypePtrOutput{})
}
