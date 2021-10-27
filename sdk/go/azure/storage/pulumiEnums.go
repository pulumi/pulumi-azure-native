


package storage

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

type BlobAccessTier string

const (
	// Optimized for storing data that is accessed frequently.
	BlobAccessTierHot = BlobAccessTier("Hot")
	// Optimized for storing data that is infrequently accessed and stored for at least 30 days.
	BlobAccessTierCool = BlobAccessTier("Cool")
	// Optimized for storing data that is rarely accessed and stored for at least 180 days with flexible latency requirements, on the order of hours.
	BlobAccessTierArchive = BlobAccessTier("Archive")
)

func (BlobAccessTier) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobAccessTier)(nil)).Elem()
}

func (e BlobAccessTier) ToBlobAccessTierOutput() BlobAccessTierOutput {
	return pulumi.ToOutput(e).(BlobAccessTierOutput)
}

func (e BlobAccessTier) ToBlobAccessTierOutputWithContext(ctx context.Context) BlobAccessTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BlobAccessTierOutput)
}

func (e BlobAccessTier) ToBlobAccessTierPtrOutput() BlobAccessTierPtrOutput {
	return e.ToBlobAccessTierPtrOutputWithContext(context.Background())
}

func (e BlobAccessTier) ToBlobAccessTierPtrOutputWithContext(ctx context.Context) BlobAccessTierPtrOutput {
	return BlobAccessTier(e).ToBlobAccessTierOutputWithContext(ctx).ToBlobAccessTierPtrOutputWithContext(ctx)
}

func (e BlobAccessTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobAccessTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobAccessTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BlobAccessTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BlobAccessTierOutput struct{ *pulumi.OutputState }

func (BlobAccessTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobAccessTier)(nil)).Elem()
}

func (o BlobAccessTierOutput) ToBlobAccessTierOutput() BlobAccessTierOutput {
	return o
}

func (o BlobAccessTierOutput) ToBlobAccessTierOutputWithContext(ctx context.Context) BlobAccessTierOutput {
	return o
}

func (o BlobAccessTierOutput) ToBlobAccessTierPtrOutput() BlobAccessTierPtrOutput {
	return o.ToBlobAccessTierPtrOutputWithContext(context.Background())
}

func (o BlobAccessTierOutput) ToBlobAccessTierPtrOutputWithContext(ctx context.Context) BlobAccessTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobAccessTier) *BlobAccessTier {
		return &v
	}).(BlobAccessTierPtrOutput)
}

func (o BlobAccessTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BlobAccessTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobAccessTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BlobAccessTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobAccessTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobAccessTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BlobAccessTierPtrOutput struct{ *pulumi.OutputState }

func (BlobAccessTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobAccessTier)(nil)).Elem()
}

func (o BlobAccessTierPtrOutput) ToBlobAccessTierPtrOutput() BlobAccessTierPtrOutput {
	return o
}

func (o BlobAccessTierPtrOutput) ToBlobAccessTierPtrOutputWithContext(ctx context.Context) BlobAccessTierPtrOutput {
	return o
}

func (o BlobAccessTierPtrOutput) Elem() BlobAccessTierOutput {
	return o.ApplyT(func(v *BlobAccessTier) BlobAccessTier {
		if v != nil {
			return *v
		}
		var ret BlobAccessTier
		return ret
	}).(BlobAccessTierOutput)
}

func (o BlobAccessTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobAccessTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BlobAccessTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BlobAccessTierInput interface {
	pulumi.Input

	ToBlobAccessTierOutput() BlobAccessTierOutput
	ToBlobAccessTierOutputWithContext(context.Context) BlobAccessTierOutput
}

var blobAccessTierPtrType = reflect.TypeOf((**BlobAccessTier)(nil)).Elem()

type BlobAccessTierPtrInput interface {
	pulumi.Input

	ToBlobAccessTierPtrOutput() BlobAccessTierPtrOutput
	ToBlobAccessTierPtrOutputWithContext(context.Context) BlobAccessTierPtrOutput
}

type blobAccessTierPtr string

func BlobAccessTierPtr(v string) BlobAccessTierPtrInput {
	return (*blobAccessTierPtr)(&v)
}

func (*blobAccessTierPtr) ElementType() reflect.Type {
	return blobAccessTierPtrType
}

func (in *blobAccessTierPtr) ToBlobAccessTierPtrOutput() BlobAccessTierPtrOutput {
	return pulumi.ToOutput(in).(BlobAccessTierPtrOutput)
}

func (in *blobAccessTierPtr) ToBlobAccessTierPtrOutputWithContext(ctx context.Context) BlobAccessTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BlobAccessTierPtrOutput)
}

type BlobType string

const (
	// Block blobs store text and binary data. Block blobs are made up of blocks of data that can be managed individually.
	BlobTypeBlock = BlobType("Block")
	// Append blobs are made up of blocks like block blobs, but are optimized for append operations.
	BlobTypeAppend = BlobType("Append")
)

func (BlobType) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobType)(nil)).Elem()
}

func (e BlobType) ToBlobTypeOutput() BlobTypeOutput {
	return pulumi.ToOutput(e).(BlobTypeOutput)
}

func (e BlobType) ToBlobTypeOutputWithContext(ctx context.Context) BlobTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BlobTypeOutput)
}

func (e BlobType) ToBlobTypePtrOutput() BlobTypePtrOutput {
	return e.ToBlobTypePtrOutputWithContext(context.Background())
}

func (e BlobType) ToBlobTypePtrOutputWithContext(ctx context.Context) BlobTypePtrOutput {
	return BlobType(e).ToBlobTypeOutputWithContext(ctx).ToBlobTypePtrOutputWithContext(ctx)
}

func (e BlobType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BlobType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BlobTypeOutput struct{ *pulumi.OutputState }

func (BlobTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobType)(nil)).Elem()
}

func (o BlobTypeOutput) ToBlobTypeOutput() BlobTypeOutput {
	return o
}

func (o BlobTypeOutput) ToBlobTypeOutputWithContext(ctx context.Context) BlobTypeOutput {
	return o
}

func (o BlobTypeOutput) ToBlobTypePtrOutput() BlobTypePtrOutput {
	return o.ToBlobTypePtrOutputWithContext(context.Background())
}

func (o BlobTypeOutput) ToBlobTypePtrOutputWithContext(ctx context.Context) BlobTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobType) *BlobType {
		return &v
	}).(BlobTypePtrOutput)
}

func (o BlobTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BlobTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BlobTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BlobTypePtrOutput struct{ *pulumi.OutputState }

func (BlobTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobType)(nil)).Elem()
}

func (o BlobTypePtrOutput) ToBlobTypePtrOutput() BlobTypePtrOutput {
	return o
}

func (o BlobTypePtrOutput) ToBlobTypePtrOutputWithContext(ctx context.Context) BlobTypePtrOutput {
	return o
}

func (o BlobTypePtrOutput) Elem() BlobTypeOutput {
	return o.ApplyT(func(v *BlobType) BlobType {
		if v != nil {
			return *v
		}
		var ret BlobType
		return ret
	}).(BlobTypeOutput)
}

func (o BlobTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BlobType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BlobTypeInput interface {
	pulumi.Input

	ToBlobTypeOutput() BlobTypeOutput
	ToBlobTypeOutputWithContext(context.Context) BlobTypeOutput
}

var blobTypePtrType = reflect.TypeOf((**BlobType)(nil)).Elem()

type BlobTypePtrInput interface {
	pulumi.Input

	ToBlobTypePtrOutput() BlobTypePtrOutput
	ToBlobTypePtrOutputWithContext(context.Context) BlobTypePtrOutput
}

type blobTypePtr string

func BlobTypePtr(v string) BlobTypePtrInput {
	return (*blobTypePtr)(&v)
}

func (*blobTypePtr) ElementType() reflect.Type {
	return blobTypePtrType
}

func (in *blobTypePtr) ToBlobTypePtrOutput() BlobTypePtrOutput {
	return pulumi.ToOutput(in).(BlobTypePtrOutput)
}

func (in *blobTypePtr) ToBlobTypePtrOutputWithContext(ctx context.Context) BlobTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BlobTypePtrOutput)
}

type Bypass string

const (
	BypassNone          = Bypass("None")
	BypassLogging       = Bypass("Logging")
	BypassMetrics       = Bypass("Metrics")
	BypassAzureServices = Bypass("AzureServices")
)

func (Bypass) ElementType() reflect.Type {
	return reflect.TypeOf((*Bypass)(nil)).Elem()
}

func (e Bypass) ToBypassOutput() BypassOutput {
	return pulumi.ToOutput(e).(BypassOutput)
}

func (e Bypass) ToBypassOutputWithContext(ctx context.Context) BypassOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BypassOutput)
}

func (e Bypass) ToBypassPtrOutput() BypassPtrOutput {
	return e.ToBypassPtrOutputWithContext(context.Background())
}

func (e Bypass) ToBypassPtrOutputWithContext(ctx context.Context) BypassPtrOutput {
	return Bypass(e).ToBypassOutputWithContext(ctx).ToBypassPtrOutputWithContext(ctx)
}

func (e Bypass) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Bypass) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Bypass) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Bypass) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BypassOutput struct{ *pulumi.OutputState }

func (BypassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Bypass)(nil)).Elem()
}

func (o BypassOutput) ToBypassOutput() BypassOutput {
	return o
}

func (o BypassOutput) ToBypassOutputWithContext(ctx context.Context) BypassOutput {
	return o
}

func (o BypassOutput) ToBypassPtrOutput() BypassPtrOutput {
	return o.ToBypassPtrOutputWithContext(context.Background())
}

func (o BypassOutput) ToBypassPtrOutputWithContext(ctx context.Context) BypassPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Bypass) *Bypass {
		return &v
	}).(BypassPtrOutput)
}

func (o BypassOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BypassOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Bypass) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BypassOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BypassOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Bypass) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BypassPtrOutput struct{ *pulumi.OutputState }

func (BypassPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bypass)(nil)).Elem()
}

func (o BypassPtrOutput) ToBypassPtrOutput() BypassPtrOutput {
	return o
}

func (o BypassPtrOutput) ToBypassPtrOutputWithContext(ctx context.Context) BypassPtrOutput {
	return o
}

func (o BypassPtrOutput) Elem() BypassOutput {
	return o.ApplyT(func(v *Bypass) Bypass {
		if v != nil {
			return *v
		}
		var ret Bypass
		return ret
	}).(BypassOutput)
}

func (o BypassPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BypassPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Bypass) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BypassInput interface {
	pulumi.Input

	ToBypassOutput() BypassOutput
	ToBypassOutputWithContext(context.Context) BypassOutput
}

var bypassPtrType = reflect.TypeOf((**Bypass)(nil)).Elem()

type BypassPtrInput interface {
	pulumi.Input

	ToBypassPtrOutput() BypassPtrOutput
	ToBypassPtrOutputWithContext(context.Context) BypassPtrOutput
}

type bypassPtr string

func BypassPtr(v string) BypassPtrInput {
	return (*bypassPtr)(&v)
}

func (*bypassPtr) ElementType() reflect.Type {
	return bypassPtrType
}

func (in *bypassPtr) ToBypassPtrOutput() BypassPtrOutput {
	return pulumi.ToOutput(in).(BypassPtrOutput)
}

func (in *bypassPtr) ToBypassPtrOutputWithContext(ctx context.Context) BypassPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BypassPtrOutput)
}

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

type DirectoryServiceOptions string

const (
	DirectoryServiceOptionsNone  = DirectoryServiceOptions("None")
	DirectoryServiceOptionsAADDS = DirectoryServiceOptions("AADDS")
	DirectoryServiceOptionsAD    = DirectoryServiceOptions("AD")
)

func (DirectoryServiceOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectoryServiceOptions)(nil)).Elem()
}

func (e DirectoryServiceOptions) ToDirectoryServiceOptionsOutput() DirectoryServiceOptionsOutput {
	return pulumi.ToOutput(e).(DirectoryServiceOptionsOutput)
}

func (e DirectoryServiceOptions) ToDirectoryServiceOptionsOutputWithContext(ctx context.Context) DirectoryServiceOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DirectoryServiceOptionsOutput)
}

func (e DirectoryServiceOptions) ToDirectoryServiceOptionsPtrOutput() DirectoryServiceOptionsPtrOutput {
	return e.ToDirectoryServiceOptionsPtrOutputWithContext(context.Background())
}

func (e DirectoryServiceOptions) ToDirectoryServiceOptionsPtrOutputWithContext(ctx context.Context) DirectoryServiceOptionsPtrOutput {
	return DirectoryServiceOptions(e).ToDirectoryServiceOptionsOutputWithContext(ctx).ToDirectoryServiceOptionsPtrOutputWithContext(ctx)
}

func (e DirectoryServiceOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DirectoryServiceOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DirectoryServiceOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DirectoryServiceOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DirectoryServiceOptionsOutput struct{ *pulumi.OutputState }

func (DirectoryServiceOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectoryServiceOptions)(nil)).Elem()
}

func (o DirectoryServiceOptionsOutput) ToDirectoryServiceOptionsOutput() DirectoryServiceOptionsOutput {
	return o
}

func (o DirectoryServiceOptionsOutput) ToDirectoryServiceOptionsOutputWithContext(ctx context.Context) DirectoryServiceOptionsOutput {
	return o
}

func (o DirectoryServiceOptionsOutput) ToDirectoryServiceOptionsPtrOutput() DirectoryServiceOptionsPtrOutput {
	return o.ToDirectoryServiceOptionsPtrOutputWithContext(context.Background())
}

func (o DirectoryServiceOptionsOutput) ToDirectoryServiceOptionsPtrOutputWithContext(ctx context.Context) DirectoryServiceOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DirectoryServiceOptions) *DirectoryServiceOptions {
		return &v
	}).(DirectoryServiceOptionsPtrOutput)
}

func (o DirectoryServiceOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DirectoryServiceOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DirectoryServiceOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DirectoryServiceOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DirectoryServiceOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DirectoryServiceOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DirectoryServiceOptionsPtrOutput struct{ *pulumi.OutputState }

func (DirectoryServiceOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryServiceOptions)(nil)).Elem()
}

func (o DirectoryServiceOptionsPtrOutput) ToDirectoryServiceOptionsPtrOutput() DirectoryServiceOptionsPtrOutput {
	return o
}

func (o DirectoryServiceOptionsPtrOutput) ToDirectoryServiceOptionsPtrOutputWithContext(ctx context.Context) DirectoryServiceOptionsPtrOutput {
	return o
}

func (o DirectoryServiceOptionsPtrOutput) Elem() DirectoryServiceOptionsOutput {
	return o.ApplyT(func(v *DirectoryServiceOptions) DirectoryServiceOptions {
		if v != nil {
			return *v
		}
		var ret DirectoryServiceOptions
		return ret
	}).(DirectoryServiceOptionsOutput)
}

func (o DirectoryServiceOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DirectoryServiceOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DirectoryServiceOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DirectoryServiceOptionsInput interface {
	pulumi.Input

	ToDirectoryServiceOptionsOutput() DirectoryServiceOptionsOutput
	ToDirectoryServiceOptionsOutputWithContext(context.Context) DirectoryServiceOptionsOutput
}

var directoryServiceOptionsPtrType = reflect.TypeOf((**DirectoryServiceOptions)(nil)).Elem()

type DirectoryServiceOptionsPtrInput interface {
	pulumi.Input

	ToDirectoryServiceOptionsPtrOutput() DirectoryServiceOptionsPtrOutput
	ToDirectoryServiceOptionsPtrOutputWithContext(context.Context) DirectoryServiceOptionsPtrOutput
}

type directoryServiceOptionsPtr string

func DirectoryServiceOptionsPtr(v string) DirectoryServiceOptionsPtrInput {
	return (*directoryServiceOptionsPtr)(&v)
}

func (*directoryServiceOptionsPtr) ElementType() reflect.Type {
	return directoryServiceOptionsPtrType
}

func (in *directoryServiceOptionsPtr) ToDirectoryServiceOptionsPtrOutput() DirectoryServiceOptionsPtrOutput {
	return pulumi.ToOutput(in).(DirectoryServiceOptionsPtrOutput)
}

func (in *directoryServiceOptionsPtr) ToDirectoryServiceOptionsPtrOutputWithContext(ctx context.Context) DirectoryServiceOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DirectoryServiceOptionsPtrOutput)
}

type EnabledProtocols string

const (
	EnabledProtocolsSMB = EnabledProtocols("SMB")
	EnabledProtocolsNFS = EnabledProtocols("NFS")
)

func (EnabledProtocols) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledProtocols)(nil)).Elem()
}

func (e EnabledProtocols) ToEnabledProtocolsOutput() EnabledProtocolsOutput {
	return pulumi.ToOutput(e).(EnabledProtocolsOutput)
}

func (e EnabledProtocols) ToEnabledProtocolsOutputWithContext(ctx context.Context) EnabledProtocolsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnabledProtocolsOutput)
}

func (e EnabledProtocols) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return e.ToEnabledProtocolsPtrOutputWithContext(context.Background())
}

func (e EnabledProtocols) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return EnabledProtocols(e).ToEnabledProtocolsOutputWithContext(ctx).ToEnabledProtocolsPtrOutputWithContext(ctx)
}

func (e EnabledProtocols) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnabledProtocols) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnabledProtocols) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnabledProtocols) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnabledProtocolsOutput struct{ *pulumi.OutputState }

func (EnabledProtocolsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledProtocols)(nil)).Elem()
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsOutput() EnabledProtocolsOutput {
	return o
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsOutputWithContext(ctx context.Context) EnabledProtocolsOutput {
	return o
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return o.ToEnabledProtocolsPtrOutputWithContext(context.Background())
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnabledProtocols) *EnabledProtocols {
		return &v
	}).(EnabledProtocolsPtrOutput)
}

func (o EnabledProtocolsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnabledProtocolsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnabledProtocols) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnabledProtocolsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnabledProtocolsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnabledProtocols) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnabledProtocolsPtrOutput struct{ *pulumi.OutputState }

func (EnabledProtocolsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledProtocols)(nil)).Elem()
}

func (o EnabledProtocolsPtrOutput) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return o
}

func (o EnabledProtocolsPtrOutput) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return o
}

func (o EnabledProtocolsPtrOutput) Elem() EnabledProtocolsOutput {
	return o.ApplyT(func(v *EnabledProtocols) EnabledProtocols {
		if v != nil {
			return *v
		}
		var ret EnabledProtocols
		return ret
	}).(EnabledProtocolsOutput)
}

func (o EnabledProtocolsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnabledProtocolsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnabledProtocols) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EnabledProtocolsInput interface {
	pulumi.Input

	ToEnabledProtocolsOutput() EnabledProtocolsOutput
	ToEnabledProtocolsOutputWithContext(context.Context) EnabledProtocolsOutput
}

var enabledProtocolsPtrType = reflect.TypeOf((**EnabledProtocols)(nil)).Elem()

type EnabledProtocolsPtrInput interface {
	pulumi.Input

	ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput
	ToEnabledProtocolsPtrOutputWithContext(context.Context) EnabledProtocolsPtrOutput
}

type enabledProtocolsPtr string

func EnabledProtocolsPtr(v string) EnabledProtocolsPtrInput {
	return (*enabledProtocolsPtr)(&v)
}

func (*enabledProtocolsPtr) ElementType() reflect.Type {
	return enabledProtocolsPtrType
}

func (in *enabledProtocolsPtr) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return pulumi.ToOutput(in).(EnabledProtocolsPtrOutput)
}

func (in *enabledProtocolsPtr) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnabledProtocolsPtrOutput)
}

type EncryptionScopeSource string

const (
	EncryptionScopeSource_Microsoft_Storage  = EncryptionScopeSource("Microsoft.Storage")
	EncryptionScopeSource_Microsoft_KeyVault = EncryptionScopeSource("Microsoft.KeyVault")
)

func (EncryptionScopeSource) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionScopeSource)(nil)).Elem()
}

func (e EncryptionScopeSource) ToEncryptionScopeSourceOutput() EncryptionScopeSourceOutput {
	return pulumi.ToOutput(e).(EncryptionScopeSourceOutput)
}

func (e EncryptionScopeSource) ToEncryptionScopeSourceOutputWithContext(ctx context.Context) EncryptionScopeSourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncryptionScopeSourceOutput)
}

func (e EncryptionScopeSource) ToEncryptionScopeSourcePtrOutput() EncryptionScopeSourcePtrOutput {
	return e.ToEncryptionScopeSourcePtrOutputWithContext(context.Background())
}

func (e EncryptionScopeSource) ToEncryptionScopeSourcePtrOutputWithContext(ctx context.Context) EncryptionScopeSourcePtrOutput {
	return EncryptionScopeSource(e).ToEncryptionScopeSourceOutputWithContext(ctx).ToEncryptionScopeSourcePtrOutputWithContext(ctx)
}

func (e EncryptionScopeSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionScopeSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionScopeSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncryptionScopeSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncryptionScopeSourceOutput struct{ *pulumi.OutputState }

func (EncryptionScopeSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionScopeSource)(nil)).Elem()
}

func (o EncryptionScopeSourceOutput) ToEncryptionScopeSourceOutput() EncryptionScopeSourceOutput {
	return o
}

func (o EncryptionScopeSourceOutput) ToEncryptionScopeSourceOutputWithContext(ctx context.Context) EncryptionScopeSourceOutput {
	return o
}

func (o EncryptionScopeSourceOutput) ToEncryptionScopeSourcePtrOutput() EncryptionScopeSourcePtrOutput {
	return o.ToEncryptionScopeSourcePtrOutputWithContext(context.Background())
}

func (o EncryptionScopeSourceOutput) ToEncryptionScopeSourcePtrOutputWithContext(ctx context.Context) EncryptionScopeSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionScopeSource) *EncryptionScopeSource {
		return &v
	}).(EncryptionScopeSourcePtrOutput)
}

func (o EncryptionScopeSourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncryptionScopeSourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionScopeSource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncryptionScopeSourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionScopeSourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionScopeSource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncryptionScopeSourcePtrOutput struct{ *pulumi.OutputState }

func (EncryptionScopeSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionScopeSource)(nil)).Elem()
}

func (o EncryptionScopeSourcePtrOutput) ToEncryptionScopeSourcePtrOutput() EncryptionScopeSourcePtrOutput {
	return o
}

func (o EncryptionScopeSourcePtrOutput) ToEncryptionScopeSourcePtrOutputWithContext(ctx context.Context) EncryptionScopeSourcePtrOutput {
	return o
}

func (o EncryptionScopeSourcePtrOutput) Elem() EncryptionScopeSourceOutput {
	return o.ApplyT(func(v *EncryptionScopeSource) EncryptionScopeSource {
		if v != nil {
			return *v
		}
		var ret EncryptionScopeSource
		return ret
	}).(EncryptionScopeSourceOutput)
}

func (o EncryptionScopeSourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionScopeSourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncryptionScopeSource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncryptionScopeSourceInput interface {
	pulumi.Input

	ToEncryptionScopeSourceOutput() EncryptionScopeSourceOutput
	ToEncryptionScopeSourceOutputWithContext(context.Context) EncryptionScopeSourceOutput
}

var encryptionScopeSourcePtrType = reflect.TypeOf((**EncryptionScopeSource)(nil)).Elem()

type EncryptionScopeSourcePtrInput interface {
	pulumi.Input

	ToEncryptionScopeSourcePtrOutput() EncryptionScopeSourcePtrOutput
	ToEncryptionScopeSourcePtrOutputWithContext(context.Context) EncryptionScopeSourcePtrOutput
}

type encryptionScopeSourcePtr string

func EncryptionScopeSourcePtr(v string) EncryptionScopeSourcePtrInput {
	return (*encryptionScopeSourcePtr)(&v)
}

func (*encryptionScopeSourcePtr) ElementType() reflect.Type {
	return encryptionScopeSourcePtrType
}

func (in *encryptionScopeSourcePtr) ToEncryptionScopeSourcePtrOutput() EncryptionScopeSourcePtrOutput {
	return pulumi.ToOutput(in).(EncryptionScopeSourcePtrOutput)
}

func (in *encryptionScopeSourcePtr) ToEncryptionScopeSourcePtrOutputWithContext(ctx context.Context) EncryptionScopeSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncryptionScopeSourcePtrOutput)
}

type EncryptionScopeStateEnum string

const (
	EncryptionScopeStateEnumEnabled  = EncryptionScopeStateEnum("Enabled")
	EncryptionScopeStateEnumDisabled = EncryptionScopeStateEnum("Disabled")
)

func (EncryptionScopeStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionScopeStateEnum)(nil)).Elem()
}

func (e EncryptionScopeStateEnum) ToEncryptionScopeStateEnumOutput() EncryptionScopeStateEnumOutput {
	return pulumi.ToOutput(e).(EncryptionScopeStateEnumOutput)
}

func (e EncryptionScopeStateEnum) ToEncryptionScopeStateEnumOutputWithContext(ctx context.Context) EncryptionScopeStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncryptionScopeStateEnumOutput)
}

func (e EncryptionScopeStateEnum) ToEncryptionScopeStateEnumPtrOutput() EncryptionScopeStateEnumPtrOutput {
	return e.ToEncryptionScopeStateEnumPtrOutputWithContext(context.Background())
}

func (e EncryptionScopeStateEnum) ToEncryptionScopeStateEnumPtrOutputWithContext(ctx context.Context) EncryptionScopeStateEnumPtrOutput {
	return EncryptionScopeStateEnum(e).ToEncryptionScopeStateEnumOutputWithContext(ctx).ToEncryptionScopeStateEnumPtrOutputWithContext(ctx)
}

func (e EncryptionScopeStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionScopeStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionScopeStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncryptionScopeStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncryptionScopeStateEnumOutput struct{ *pulumi.OutputState }

func (EncryptionScopeStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionScopeStateEnum)(nil)).Elem()
}

func (o EncryptionScopeStateEnumOutput) ToEncryptionScopeStateEnumOutput() EncryptionScopeStateEnumOutput {
	return o
}

func (o EncryptionScopeStateEnumOutput) ToEncryptionScopeStateEnumOutputWithContext(ctx context.Context) EncryptionScopeStateEnumOutput {
	return o
}

func (o EncryptionScopeStateEnumOutput) ToEncryptionScopeStateEnumPtrOutput() EncryptionScopeStateEnumPtrOutput {
	return o.ToEncryptionScopeStateEnumPtrOutputWithContext(context.Background())
}

func (o EncryptionScopeStateEnumOutput) ToEncryptionScopeStateEnumPtrOutputWithContext(ctx context.Context) EncryptionScopeStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionScopeStateEnum) *EncryptionScopeStateEnum {
		return &v
	}).(EncryptionScopeStateEnumPtrOutput)
}

func (o EncryptionScopeStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncryptionScopeStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionScopeStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncryptionScopeStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionScopeStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionScopeStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncryptionScopeStateEnumPtrOutput struct{ *pulumi.OutputState }

func (EncryptionScopeStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionScopeStateEnum)(nil)).Elem()
}

func (o EncryptionScopeStateEnumPtrOutput) ToEncryptionScopeStateEnumPtrOutput() EncryptionScopeStateEnumPtrOutput {
	return o
}

func (o EncryptionScopeStateEnumPtrOutput) ToEncryptionScopeStateEnumPtrOutputWithContext(ctx context.Context) EncryptionScopeStateEnumPtrOutput {
	return o
}

func (o EncryptionScopeStateEnumPtrOutput) Elem() EncryptionScopeStateEnumOutput {
	return o.ApplyT(func(v *EncryptionScopeStateEnum) EncryptionScopeStateEnum {
		if v != nil {
			return *v
		}
		var ret EncryptionScopeStateEnum
		return ret
	}).(EncryptionScopeStateEnumOutput)
}

func (o EncryptionScopeStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionScopeStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncryptionScopeStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncryptionScopeStateEnumInput interface {
	pulumi.Input

	ToEncryptionScopeStateEnumOutput() EncryptionScopeStateEnumOutput
	ToEncryptionScopeStateEnumOutputWithContext(context.Context) EncryptionScopeStateEnumOutput
}

var encryptionScopeStateEnumPtrType = reflect.TypeOf((**EncryptionScopeStateEnum)(nil)).Elem()

type EncryptionScopeStateEnumPtrInput interface {
	pulumi.Input

	ToEncryptionScopeStateEnumPtrOutput() EncryptionScopeStateEnumPtrOutput
	ToEncryptionScopeStateEnumPtrOutputWithContext(context.Context) EncryptionScopeStateEnumPtrOutput
}

type encryptionScopeStateEnumPtr string

func EncryptionScopeStateEnumPtr(v string) EncryptionScopeStateEnumPtrInput {
	return (*encryptionScopeStateEnumPtr)(&v)
}

func (*encryptionScopeStateEnumPtr) ElementType() reflect.Type {
	return encryptionScopeStateEnumPtrType
}

func (in *encryptionScopeStateEnumPtr) ToEncryptionScopeStateEnumPtrOutput() EncryptionScopeStateEnumPtrOutput {
	return pulumi.ToOutput(in).(EncryptionScopeStateEnumPtrOutput)
}

func (in *encryptionScopeStateEnumPtr) ToEncryptionScopeStateEnumPtrOutputWithContext(ctx context.Context) EncryptionScopeStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncryptionScopeStateEnumPtrOutput)
}

type ExpirationAction string

const (
	ExpirationActionLog = ExpirationAction("Log")
)

func (ExpirationAction) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpirationAction)(nil)).Elem()
}

func (e ExpirationAction) ToExpirationActionOutput() ExpirationActionOutput {
	return pulumi.ToOutput(e).(ExpirationActionOutput)
}

func (e ExpirationAction) ToExpirationActionOutputWithContext(ctx context.Context) ExpirationActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpirationActionOutput)
}

func (e ExpirationAction) ToExpirationActionPtrOutput() ExpirationActionPtrOutput {
	return e.ToExpirationActionPtrOutputWithContext(context.Background())
}

func (e ExpirationAction) ToExpirationActionPtrOutputWithContext(ctx context.Context) ExpirationActionPtrOutput {
	return ExpirationAction(e).ToExpirationActionOutputWithContext(ctx).ToExpirationActionPtrOutputWithContext(ctx)
}

func (e ExpirationAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpirationAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpirationAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpirationAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpirationActionOutput struct{ *pulumi.OutputState }

func (ExpirationActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpirationAction)(nil)).Elem()
}

func (o ExpirationActionOutput) ToExpirationActionOutput() ExpirationActionOutput {
	return o
}

func (o ExpirationActionOutput) ToExpirationActionOutputWithContext(ctx context.Context) ExpirationActionOutput {
	return o
}

func (o ExpirationActionOutput) ToExpirationActionPtrOutput() ExpirationActionPtrOutput {
	return o.ToExpirationActionPtrOutputWithContext(context.Background())
}

func (o ExpirationActionOutput) ToExpirationActionPtrOutputWithContext(ctx context.Context) ExpirationActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpirationAction) *ExpirationAction {
		return &v
	}).(ExpirationActionPtrOutput)
}

func (o ExpirationActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpirationActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpirationAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpirationActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpirationActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpirationAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpirationActionPtrOutput struct{ *pulumi.OutputState }

func (ExpirationActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpirationAction)(nil)).Elem()
}

func (o ExpirationActionPtrOutput) ToExpirationActionPtrOutput() ExpirationActionPtrOutput {
	return o
}

func (o ExpirationActionPtrOutput) ToExpirationActionPtrOutputWithContext(ctx context.Context) ExpirationActionPtrOutput {
	return o
}

func (o ExpirationActionPtrOutput) Elem() ExpirationActionOutput {
	return o.ApplyT(func(v *ExpirationAction) ExpirationAction {
		if v != nil {
			return *v
		}
		var ret ExpirationAction
		return ret
	}).(ExpirationActionOutput)
}

func (o ExpirationActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpirationActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpirationAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpirationActionInput interface {
	pulumi.Input

	ToExpirationActionOutput() ExpirationActionOutput
	ToExpirationActionOutputWithContext(context.Context) ExpirationActionOutput
}

var expirationActionPtrType = reflect.TypeOf((**ExpirationAction)(nil)).Elem()

type ExpirationActionPtrInput interface {
	pulumi.Input

	ToExpirationActionPtrOutput() ExpirationActionPtrOutput
	ToExpirationActionPtrOutputWithContext(context.Context) ExpirationActionPtrOutput
}

type expirationActionPtr string

func ExpirationActionPtr(v string) ExpirationActionPtrInput {
	return (*expirationActionPtr)(&v)
}

func (*expirationActionPtr) ElementType() reflect.Type {
	return expirationActionPtrType
}

func (in *expirationActionPtr) ToExpirationActionPtrOutput() ExpirationActionPtrOutput {
	return pulumi.ToOutput(in).(ExpirationActionPtrOutput)
}

func (in *expirationActionPtr) ToExpirationActionPtrOutputWithContext(ctx context.Context) ExpirationActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpirationActionPtrOutput)
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

type HttpProtocol string

const (
	HttpProtocol_Https_http = HttpProtocol("https,http")
	HttpProtocolHttps       = HttpProtocol("https")
)

func (HttpProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProtocol)(nil)).Elem()
}

func (e HttpProtocol) ToHttpProtocolOutput() HttpProtocolOutput {
	return pulumi.ToOutput(e).(HttpProtocolOutput)
}

func (e HttpProtocol) ToHttpProtocolOutputWithContext(ctx context.Context) HttpProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HttpProtocolOutput)
}

func (e HttpProtocol) ToHttpProtocolPtrOutput() HttpProtocolPtrOutput {
	return e.ToHttpProtocolPtrOutputWithContext(context.Background())
}

func (e HttpProtocol) ToHttpProtocolPtrOutputWithContext(ctx context.Context) HttpProtocolPtrOutput {
	return HttpProtocol(e).ToHttpProtocolOutputWithContext(ctx).ToHttpProtocolPtrOutputWithContext(ctx)
}

func (e HttpProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HttpProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HttpProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HttpProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HttpProtocolOutput struct{ *pulumi.OutputState }

func (HttpProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProtocol)(nil)).Elem()
}

func (o HttpProtocolOutput) ToHttpProtocolOutput() HttpProtocolOutput {
	return o
}

func (o HttpProtocolOutput) ToHttpProtocolOutputWithContext(ctx context.Context) HttpProtocolOutput {
	return o
}

func (o HttpProtocolOutput) ToHttpProtocolPtrOutput() HttpProtocolPtrOutput {
	return o.ToHttpProtocolPtrOutputWithContext(context.Background())
}

func (o HttpProtocolOutput) ToHttpProtocolPtrOutputWithContext(ctx context.Context) HttpProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpProtocol) *HttpProtocol {
		return &v
	}).(HttpProtocolPtrOutput)
}

func (o HttpProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HttpProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HttpProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HttpProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HttpProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HttpProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HttpProtocolPtrOutput struct{ *pulumi.OutputState }

func (HttpProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProtocol)(nil)).Elem()
}

func (o HttpProtocolPtrOutput) ToHttpProtocolPtrOutput() HttpProtocolPtrOutput {
	return o
}

func (o HttpProtocolPtrOutput) ToHttpProtocolPtrOutputWithContext(ctx context.Context) HttpProtocolPtrOutput {
	return o
}

func (o HttpProtocolPtrOutput) Elem() HttpProtocolOutput {
	return o.ApplyT(func(v *HttpProtocol) HttpProtocol {
		if v != nil {
			return *v
		}
		var ret HttpProtocol
		return ret
	}).(HttpProtocolOutput)
}

func (o HttpProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HttpProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HttpProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HttpProtocolInput interface {
	pulumi.Input

	ToHttpProtocolOutput() HttpProtocolOutput
	ToHttpProtocolOutputWithContext(context.Context) HttpProtocolOutput
}

var httpProtocolPtrType = reflect.TypeOf((**HttpProtocol)(nil)).Elem()

type HttpProtocolPtrInput interface {
	pulumi.Input

	ToHttpProtocolPtrOutput() HttpProtocolPtrOutput
	ToHttpProtocolPtrOutputWithContext(context.Context) HttpProtocolPtrOutput
}

type httpProtocolPtr string

func HttpProtocolPtr(v string) HttpProtocolPtrInput {
	return (*httpProtocolPtr)(&v)
}

func (*httpProtocolPtr) ElementType() reflect.Type {
	return httpProtocolPtrType
}

func (in *httpProtocolPtr) ToHttpProtocolPtrOutput() HttpProtocolPtrOutput {
	return pulumi.ToOutput(in).(HttpProtocolPtrOutput)
}

func (in *httpProtocolPtr) ToHttpProtocolPtrOutputWithContext(ctx context.Context) HttpProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HttpProtocolPtrOutput)
}

type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned,UserAssigned")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (e IdentityType) ToIdentityTypeOutput() IdentityTypeOutput {
	return pulumi.ToOutput(e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return e.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (e IdentityType) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return IdentityType(e).ToIdentityTypeOutputWithContext(ctx).ToIdentityTypePtrOutputWithContext(ctx)
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityTypeOutput struct{ *pulumi.OutputState }

func (IdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (o IdentityTypeOutput) ToIdentityTypeOutput() IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityType) *IdentityType {
		return &v
	}).(IdentityTypePtrOutput)
}

func (o IdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityType)(nil)).Elem()
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) Elem() IdentityTypeOutput {
	return o.ApplyT(func(v *IdentityType) IdentityType {
		if v != nil {
			return *v
		}
		var ret IdentityType
		return ret
	}).(IdentityTypeOutput)
}

func (o IdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IdentityTypeInput interface {
	pulumi.Input

	ToIdentityTypeOutput() IdentityTypeOutput
	ToIdentityTypeOutputWithContext(context.Context) IdentityTypeOutput
}

var identityTypePtrType = reflect.TypeOf((**IdentityType)(nil)).Elem()

type IdentityTypePtrInput interface {
	pulumi.Input

	ToIdentityTypePtrOutput() IdentityTypePtrOutput
	ToIdentityTypePtrOutputWithContext(context.Context) IdentityTypePtrOutput
}

type identityTypePtr string

func IdentityTypePtr(v string) IdentityTypePtrInput {
	return (*identityTypePtr)(&v)
}

func (*identityTypePtr) ElementType() reflect.Type {
	return identityTypePtrType
}

func (in *identityTypePtr) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityTypePtrOutput)
}

type InventoryRuleType string

const (
	InventoryRuleTypeInventory = InventoryRuleType("Inventory")
)

func (InventoryRuleType) ElementType() reflect.Type {
	return reflect.TypeOf((*InventoryRuleType)(nil)).Elem()
}

func (e InventoryRuleType) ToInventoryRuleTypeOutput() InventoryRuleTypeOutput {
	return pulumi.ToOutput(e).(InventoryRuleTypeOutput)
}

func (e InventoryRuleType) ToInventoryRuleTypeOutputWithContext(ctx context.Context) InventoryRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InventoryRuleTypeOutput)
}

func (e InventoryRuleType) ToInventoryRuleTypePtrOutput() InventoryRuleTypePtrOutput {
	return e.ToInventoryRuleTypePtrOutputWithContext(context.Background())
}

func (e InventoryRuleType) ToInventoryRuleTypePtrOutputWithContext(ctx context.Context) InventoryRuleTypePtrOutput {
	return InventoryRuleType(e).ToInventoryRuleTypeOutputWithContext(ctx).ToInventoryRuleTypePtrOutputWithContext(ctx)
}

func (e InventoryRuleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InventoryRuleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InventoryRuleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InventoryRuleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InventoryRuleTypeOutput struct{ *pulumi.OutputState }

func (InventoryRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InventoryRuleType)(nil)).Elem()
}

func (o InventoryRuleTypeOutput) ToInventoryRuleTypeOutput() InventoryRuleTypeOutput {
	return o
}

func (o InventoryRuleTypeOutput) ToInventoryRuleTypeOutputWithContext(ctx context.Context) InventoryRuleTypeOutput {
	return o
}

func (o InventoryRuleTypeOutput) ToInventoryRuleTypePtrOutput() InventoryRuleTypePtrOutput {
	return o.ToInventoryRuleTypePtrOutputWithContext(context.Background())
}

func (o InventoryRuleTypeOutput) ToInventoryRuleTypePtrOutputWithContext(ctx context.Context) InventoryRuleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InventoryRuleType) *InventoryRuleType {
		return &v
	}).(InventoryRuleTypePtrOutput)
}

func (o InventoryRuleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InventoryRuleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InventoryRuleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InventoryRuleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InventoryRuleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InventoryRuleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InventoryRuleTypePtrOutput struct{ *pulumi.OutputState }

func (InventoryRuleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InventoryRuleType)(nil)).Elem()
}

func (o InventoryRuleTypePtrOutput) ToInventoryRuleTypePtrOutput() InventoryRuleTypePtrOutput {
	return o
}

func (o InventoryRuleTypePtrOutput) ToInventoryRuleTypePtrOutputWithContext(ctx context.Context) InventoryRuleTypePtrOutput {
	return o
}

func (o InventoryRuleTypePtrOutput) Elem() InventoryRuleTypeOutput {
	return o.ApplyT(func(v *InventoryRuleType) InventoryRuleType {
		if v != nil {
			return *v
		}
		var ret InventoryRuleType
		return ret
	}).(InventoryRuleTypeOutput)
}

func (o InventoryRuleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InventoryRuleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InventoryRuleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InventoryRuleTypeInput interface {
	pulumi.Input

	ToInventoryRuleTypeOutput() InventoryRuleTypeOutput
	ToInventoryRuleTypeOutputWithContext(context.Context) InventoryRuleTypeOutput
}

var inventoryRuleTypePtrType = reflect.TypeOf((**InventoryRuleType)(nil)).Elem()

type InventoryRuleTypePtrInput interface {
	pulumi.Input

	ToInventoryRuleTypePtrOutput() InventoryRuleTypePtrOutput
	ToInventoryRuleTypePtrOutputWithContext(context.Context) InventoryRuleTypePtrOutput
}

type inventoryRuleTypePtr string

func InventoryRuleTypePtr(v string) InventoryRuleTypePtrInput {
	return (*inventoryRuleTypePtr)(&v)
}

func (*inventoryRuleTypePtr) ElementType() reflect.Type {
	return inventoryRuleTypePtrType
}

func (in *inventoryRuleTypePtr) ToInventoryRuleTypePtrOutput() InventoryRuleTypePtrOutput {
	return pulumi.ToOutput(in).(InventoryRuleTypePtrOutput)
}

func (in *inventoryRuleTypePtr) ToInventoryRuleTypePtrOutputWithContext(ctx context.Context) InventoryRuleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InventoryRuleTypePtrOutput)
}

type KeySource string

const (
	KeySource_Microsoft_Storage  = KeySource("Microsoft.Storage")
	KeySource_Microsoft_Keyvault = KeySource("Microsoft.Keyvault")
)

func (KeySource) ElementType() reflect.Type {
	return reflect.TypeOf((*KeySource)(nil)).Elem()
}

func (e KeySource) ToKeySourceOutput() KeySourceOutput {
	return pulumi.ToOutput(e).(KeySourceOutput)
}

func (e KeySource) ToKeySourceOutputWithContext(ctx context.Context) KeySourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeySourceOutput)
}

func (e KeySource) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return e.ToKeySourcePtrOutputWithContext(context.Background())
}

func (e KeySource) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return KeySource(e).ToKeySourceOutputWithContext(ctx).ToKeySourcePtrOutputWithContext(ctx)
}

func (e KeySource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeySource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeySource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeySource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeySourceOutput struct{ *pulumi.OutputState }

func (KeySourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeySource)(nil)).Elem()
}

func (o KeySourceOutput) ToKeySourceOutput() KeySourceOutput {
	return o
}

func (o KeySourceOutput) ToKeySourceOutputWithContext(ctx context.Context) KeySourceOutput {
	return o
}

func (o KeySourceOutput) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return o.ToKeySourcePtrOutputWithContext(context.Background())
}

func (o KeySourceOutput) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeySource) *KeySource {
		return &v
	}).(KeySourcePtrOutput)
}

func (o KeySourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeySourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeySource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeySourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeySourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeySource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeySourcePtrOutput struct{ *pulumi.OutputState }

func (KeySourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeySource)(nil)).Elem()
}

func (o KeySourcePtrOutput) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return o
}

func (o KeySourcePtrOutput) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return o
}

func (o KeySourcePtrOutput) Elem() KeySourceOutput {
	return o.ApplyT(func(v *KeySource) KeySource {
		if v != nil {
			return *v
		}
		var ret KeySource
		return ret
	}).(KeySourceOutput)
}

func (o KeySourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeySourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeySource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KeySourceInput interface {
	pulumi.Input

	ToKeySourceOutput() KeySourceOutput
	ToKeySourceOutputWithContext(context.Context) KeySourceOutput
}

var keySourcePtrType = reflect.TypeOf((**KeySource)(nil)).Elem()

type KeySourcePtrInput interface {
	pulumi.Input

	ToKeySourcePtrOutput() KeySourcePtrOutput
	ToKeySourcePtrOutputWithContext(context.Context) KeySourcePtrOutput
}

type keySourcePtr string

func KeySourcePtr(v string) KeySourcePtrInput {
	return (*keySourcePtr)(&v)
}

func (*keySourcePtr) ElementType() reflect.Type {
	return keySourcePtrType
}

func (in *keySourcePtr) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return pulumi.ToOutput(in).(KeySourcePtrOutput)
}

func (in *keySourcePtr) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeySourcePtrOutput)
}

type KeyType string

const (
	KeyTypeService = KeyType("Service")
	KeyTypeAccount = KeyType("Account")
)

func (KeyType) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyType)(nil)).Elem()
}

func (e KeyType) ToKeyTypeOutput() KeyTypeOutput {
	return pulumi.ToOutput(e).(KeyTypeOutput)
}

func (e KeyType) ToKeyTypeOutputWithContext(ctx context.Context) KeyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeyTypeOutput)
}

func (e KeyType) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return e.ToKeyTypePtrOutputWithContext(context.Background())
}

func (e KeyType) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return KeyType(e).ToKeyTypeOutputWithContext(ctx).ToKeyTypePtrOutputWithContext(ctx)
}

func (e KeyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeyTypeOutput struct{ *pulumi.OutputState }

func (KeyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyType)(nil)).Elem()
}

func (o KeyTypeOutput) ToKeyTypeOutput() KeyTypeOutput {
	return o
}

func (o KeyTypeOutput) ToKeyTypeOutputWithContext(ctx context.Context) KeyTypeOutput {
	return o
}

func (o KeyTypeOutput) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return o.ToKeyTypePtrOutputWithContext(context.Background())
}

func (o KeyTypeOutput) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyType) *KeyType {
		return &v
	}).(KeyTypePtrOutput)
}

func (o KeyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeyTypePtrOutput struct{ *pulumi.OutputState }

func (KeyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyType)(nil)).Elem()
}

func (o KeyTypePtrOutput) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return o
}

func (o KeyTypePtrOutput) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return o
}

func (o KeyTypePtrOutput) Elem() KeyTypeOutput {
	return o.ApplyT(func(v *KeyType) KeyType {
		if v != nil {
			return *v
		}
		var ret KeyType
		return ret
	}).(KeyTypeOutput)
}

func (o KeyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KeyTypeInput interface {
	pulumi.Input

	ToKeyTypeOutput() KeyTypeOutput
	ToKeyTypeOutputWithContext(context.Context) KeyTypeOutput
}

var keyTypePtrType = reflect.TypeOf((**KeyType)(nil)).Elem()

type KeyTypePtrInput interface {
	pulumi.Input

	ToKeyTypePtrOutput() KeyTypePtrOutput
	ToKeyTypePtrOutputWithContext(context.Context) KeyTypePtrOutput
}

type keyTypePtr string

func KeyTypePtr(v string) KeyTypePtrInput {
	return (*keyTypePtr)(&v)
}

func (*keyTypePtr) ElementType() reflect.Type {
	return keyTypePtrType
}

func (in *keyTypePtr) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return pulumi.ToOutput(in).(KeyTypePtrOutput)
}

func (in *keyTypePtr) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeyTypePtrOutput)
}

type Kind string

const (
	KindStorage          = Kind("Storage")
	KindStorageV2        = Kind("StorageV2")
	KindBlobStorage      = Kind("BlobStorage")
	KindFileStorage      = Kind("FileStorage")
	KindBlockBlobStorage = Kind("BlockBlobStorage")
)

func (Kind) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (e Kind) ToKindOutput() KindOutput {
	return pulumi.ToOutput(e).(KindOutput)
}

func (e Kind) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KindOutput)
}

func (e Kind) ToKindPtrOutput() KindPtrOutput {
	return e.ToKindPtrOutputWithContext(context.Background())
}

func (e Kind) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return Kind(e).ToKindOutputWithContext(ctx).ToKindPtrOutputWithContext(ctx)
}

func (e Kind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Kind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KindOutput struct{ *pulumi.OutputState }

func (KindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (o KindOutput) ToKindOutput() KindOutput {
	return o
}

func (o KindOutput) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return o
}

func (o KindOutput) ToKindPtrOutput() KindPtrOutput {
	return o.ToKindPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Kind) *Kind {
		return &v
	}).(KindPtrOutput)
}

func (o KindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KindPtrOutput struct{ *pulumi.OutputState }

func (KindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kind)(nil)).Elem()
}

func (o KindPtrOutput) ToKindPtrOutput() KindPtrOutput {
	return o
}

func (o KindPtrOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o
}

func (o KindPtrOutput) Elem() KindOutput {
	return o.ApplyT(func(v *Kind) Kind {
		if v != nil {
			return *v
		}
		var ret Kind
		return ret
	}).(KindOutput)
}

func (o KindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Kind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KindInput interface {
	pulumi.Input

	ToKindOutput() KindOutput
	ToKindOutputWithContext(context.Context) KindOutput
}

var kindPtrType = reflect.TypeOf((**Kind)(nil)).Elem()

type KindPtrInput interface {
	pulumi.Input

	ToKindPtrOutput() KindPtrOutput
	ToKindPtrOutputWithContext(context.Context) KindPtrOutput
}

type kindPtr string

func KindPtr(v string) KindPtrInput {
	return (*kindPtr)(&v)
}

func (*kindPtr) ElementType() reflect.Type {
	return kindPtrType
}

func (in *kindPtr) ToKindPtrOutput() KindPtrOutput {
	return pulumi.ToOutput(in).(KindPtrOutput)
}

func (in *kindPtr) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KindPtrOutput)
}

type LargeFileSharesState string

const (
	LargeFileSharesStateDisabled = LargeFileSharesState("Disabled")
	LargeFileSharesStateEnabled  = LargeFileSharesState("Enabled")
)

func (LargeFileSharesState) ElementType() reflect.Type {
	return reflect.TypeOf((*LargeFileSharesState)(nil)).Elem()
}

func (e LargeFileSharesState) ToLargeFileSharesStateOutput() LargeFileSharesStateOutput {
	return pulumi.ToOutput(e).(LargeFileSharesStateOutput)
}

func (e LargeFileSharesState) ToLargeFileSharesStateOutputWithContext(ctx context.Context) LargeFileSharesStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LargeFileSharesStateOutput)
}

func (e LargeFileSharesState) ToLargeFileSharesStatePtrOutput() LargeFileSharesStatePtrOutput {
	return e.ToLargeFileSharesStatePtrOutputWithContext(context.Background())
}

func (e LargeFileSharesState) ToLargeFileSharesStatePtrOutputWithContext(ctx context.Context) LargeFileSharesStatePtrOutput {
	return LargeFileSharesState(e).ToLargeFileSharesStateOutputWithContext(ctx).ToLargeFileSharesStatePtrOutputWithContext(ctx)
}

func (e LargeFileSharesState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LargeFileSharesState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LargeFileSharesState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LargeFileSharesState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LargeFileSharesStateOutput struct{ *pulumi.OutputState }

func (LargeFileSharesStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LargeFileSharesState)(nil)).Elem()
}

func (o LargeFileSharesStateOutput) ToLargeFileSharesStateOutput() LargeFileSharesStateOutput {
	return o
}

func (o LargeFileSharesStateOutput) ToLargeFileSharesStateOutputWithContext(ctx context.Context) LargeFileSharesStateOutput {
	return o
}

func (o LargeFileSharesStateOutput) ToLargeFileSharesStatePtrOutput() LargeFileSharesStatePtrOutput {
	return o.ToLargeFileSharesStatePtrOutputWithContext(context.Background())
}

func (o LargeFileSharesStateOutput) ToLargeFileSharesStatePtrOutputWithContext(ctx context.Context) LargeFileSharesStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LargeFileSharesState) *LargeFileSharesState {
		return &v
	}).(LargeFileSharesStatePtrOutput)
}

func (o LargeFileSharesStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LargeFileSharesStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LargeFileSharesState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LargeFileSharesStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LargeFileSharesStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LargeFileSharesState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LargeFileSharesStatePtrOutput struct{ *pulumi.OutputState }

func (LargeFileSharesStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LargeFileSharesState)(nil)).Elem()
}

func (o LargeFileSharesStatePtrOutput) ToLargeFileSharesStatePtrOutput() LargeFileSharesStatePtrOutput {
	return o
}

func (o LargeFileSharesStatePtrOutput) ToLargeFileSharesStatePtrOutputWithContext(ctx context.Context) LargeFileSharesStatePtrOutput {
	return o
}

func (o LargeFileSharesStatePtrOutput) Elem() LargeFileSharesStateOutput {
	return o.ApplyT(func(v *LargeFileSharesState) LargeFileSharesState {
		if v != nil {
			return *v
		}
		var ret LargeFileSharesState
		return ret
	}).(LargeFileSharesStateOutput)
}

func (o LargeFileSharesStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LargeFileSharesStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LargeFileSharesState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LargeFileSharesStateInput interface {
	pulumi.Input

	ToLargeFileSharesStateOutput() LargeFileSharesStateOutput
	ToLargeFileSharesStateOutputWithContext(context.Context) LargeFileSharesStateOutput
}

var largeFileSharesStatePtrType = reflect.TypeOf((**LargeFileSharesState)(nil)).Elem()

type LargeFileSharesStatePtrInput interface {
	pulumi.Input

	ToLargeFileSharesStatePtrOutput() LargeFileSharesStatePtrOutput
	ToLargeFileSharesStatePtrOutputWithContext(context.Context) LargeFileSharesStatePtrOutput
}

type largeFileSharesStatePtr string

func LargeFileSharesStatePtr(v string) LargeFileSharesStatePtrInput {
	return (*largeFileSharesStatePtr)(&v)
}

func (*largeFileSharesStatePtr) ElementType() reflect.Type {
	return largeFileSharesStatePtrType
}

func (in *largeFileSharesStatePtr) ToLargeFileSharesStatePtrOutput() LargeFileSharesStatePtrOutput {
	return pulumi.ToOutput(in).(LargeFileSharesStatePtrOutput)
}

func (in *largeFileSharesStatePtr) ToLargeFileSharesStatePtrOutputWithContext(ctx context.Context) LargeFileSharesStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LargeFileSharesStatePtrOutput)
}

type MinimumTlsVersion string

const (
	MinimumTlsVersion_TLS1_0 = MinimumTlsVersion("TLS1_0")
	MinimumTlsVersion_TLS1_1 = MinimumTlsVersion("TLS1_1")
	MinimumTlsVersion_TLS1_2 = MinimumTlsVersion("TLS1_2")
)

func (MinimumTlsVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*MinimumTlsVersion)(nil)).Elem()
}

func (e MinimumTlsVersion) ToMinimumTlsVersionOutput() MinimumTlsVersionOutput {
	return pulumi.ToOutput(e).(MinimumTlsVersionOutput)
}

func (e MinimumTlsVersion) ToMinimumTlsVersionOutputWithContext(ctx context.Context) MinimumTlsVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MinimumTlsVersionOutput)
}

func (e MinimumTlsVersion) ToMinimumTlsVersionPtrOutput() MinimumTlsVersionPtrOutput {
	return e.ToMinimumTlsVersionPtrOutputWithContext(context.Background())
}

func (e MinimumTlsVersion) ToMinimumTlsVersionPtrOutputWithContext(ctx context.Context) MinimumTlsVersionPtrOutput {
	return MinimumTlsVersion(e).ToMinimumTlsVersionOutputWithContext(ctx).ToMinimumTlsVersionPtrOutputWithContext(ctx)
}

func (e MinimumTlsVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MinimumTlsVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MinimumTlsVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MinimumTlsVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MinimumTlsVersionOutput struct{ *pulumi.OutputState }

func (MinimumTlsVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MinimumTlsVersion)(nil)).Elem()
}

func (o MinimumTlsVersionOutput) ToMinimumTlsVersionOutput() MinimumTlsVersionOutput {
	return o
}

func (o MinimumTlsVersionOutput) ToMinimumTlsVersionOutputWithContext(ctx context.Context) MinimumTlsVersionOutput {
	return o
}

func (o MinimumTlsVersionOutput) ToMinimumTlsVersionPtrOutput() MinimumTlsVersionPtrOutput {
	return o.ToMinimumTlsVersionPtrOutputWithContext(context.Background())
}

func (o MinimumTlsVersionOutput) ToMinimumTlsVersionPtrOutputWithContext(ctx context.Context) MinimumTlsVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MinimumTlsVersion) *MinimumTlsVersion {
		return &v
	}).(MinimumTlsVersionPtrOutput)
}

func (o MinimumTlsVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MinimumTlsVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MinimumTlsVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MinimumTlsVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MinimumTlsVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MinimumTlsVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MinimumTlsVersionPtrOutput struct{ *pulumi.OutputState }

func (MinimumTlsVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MinimumTlsVersion)(nil)).Elem()
}

func (o MinimumTlsVersionPtrOutput) ToMinimumTlsVersionPtrOutput() MinimumTlsVersionPtrOutput {
	return o
}

func (o MinimumTlsVersionPtrOutput) ToMinimumTlsVersionPtrOutputWithContext(ctx context.Context) MinimumTlsVersionPtrOutput {
	return o
}

func (o MinimumTlsVersionPtrOutput) Elem() MinimumTlsVersionOutput {
	return o.ApplyT(func(v *MinimumTlsVersion) MinimumTlsVersion {
		if v != nil {
			return *v
		}
		var ret MinimumTlsVersion
		return ret
	}).(MinimumTlsVersionOutput)
}

func (o MinimumTlsVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MinimumTlsVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MinimumTlsVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MinimumTlsVersionInput interface {
	pulumi.Input

	ToMinimumTlsVersionOutput() MinimumTlsVersionOutput
	ToMinimumTlsVersionOutputWithContext(context.Context) MinimumTlsVersionOutput
}

var minimumTlsVersionPtrType = reflect.TypeOf((**MinimumTlsVersion)(nil)).Elem()

type MinimumTlsVersionPtrInput interface {
	pulumi.Input

	ToMinimumTlsVersionPtrOutput() MinimumTlsVersionPtrOutput
	ToMinimumTlsVersionPtrOutputWithContext(context.Context) MinimumTlsVersionPtrOutput
}

type minimumTlsVersionPtr string

func MinimumTlsVersionPtr(v string) MinimumTlsVersionPtrInput {
	return (*minimumTlsVersionPtr)(&v)
}

func (*minimumTlsVersionPtr) ElementType() reflect.Type {
	return minimumTlsVersionPtrType
}

func (in *minimumTlsVersionPtr) ToMinimumTlsVersionPtrOutput() MinimumTlsVersionPtrOutput {
	return pulumi.ToOutput(in).(MinimumTlsVersionPtrOutput)
}

func (in *minimumTlsVersionPtr) ToMinimumTlsVersionPtrOutputWithContext(ctx context.Context) MinimumTlsVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MinimumTlsVersionPtrOutput)
}

type Name string

const (
	NameAccessTimeTracking = Name("AccessTimeTracking")
)

func (Name) ElementType() reflect.Type {
	return reflect.TypeOf((*Name)(nil)).Elem()
}

func (e Name) ToNameOutput() NameOutput {
	return pulumi.ToOutput(e).(NameOutput)
}

func (e Name) ToNameOutputWithContext(ctx context.Context) NameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NameOutput)
}

func (e Name) ToNamePtrOutput() NamePtrOutput {
	return e.ToNamePtrOutputWithContext(context.Background())
}

func (e Name) ToNamePtrOutputWithContext(ctx context.Context) NamePtrOutput {
	return Name(e).ToNameOutputWithContext(ctx).ToNamePtrOutputWithContext(ctx)
}

func (e Name) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Name) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Name) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Name) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NameOutput struct{ *pulumi.OutputState }

func (NameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Name)(nil)).Elem()
}

func (o NameOutput) ToNameOutput() NameOutput {
	return o
}

func (o NameOutput) ToNameOutputWithContext(ctx context.Context) NameOutput {
	return o
}

func (o NameOutput) ToNamePtrOutput() NamePtrOutput {
	return o.ToNamePtrOutputWithContext(context.Background())
}

func (o NameOutput) ToNamePtrOutputWithContext(ctx context.Context) NamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Name) *Name {
		return &v
	}).(NamePtrOutput)
}

func (o NameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Name) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Name) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NamePtrOutput struct{ *pulumi.OutputState }

func (NamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Name)(nil)).Elem()
}

func (o NamePtrOutput) ToNamePtrOutput() NamePtrOutput {
	return o
}

func (o NamePtrOutput) ToNamePtrOutputWithContext(ctx context.Context) NamePtrOutput {
	return o
}

func (o NamePtrOutput) Elem() NameOutput {
	return o.ApplyT(func(v *Name) Name {
		if v != nil {
			return *v
		}
		var ret Name
		return ret
	}).(NameOutput)
}

func (o NamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Name) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NameInput interface {
	pulumi.Input

	ToNameOutput() NameOutput
	ToNameOutputWithContext(context.Context) NameOutput
}

var namePtrType = reflect.TypeOf((**Name)(nil)).Elem()

type NamePtrInput interface {
	pulumi.Input

	ToNamePtrOutput() NamePtrOutput
	ToNamePtrOutputWithContext(context.Context) NamePtrOutput
}

type namePtr string

func NamePtr(v string) NamePtrInput {
	return (*namePtr)(&v)
}

func (*namePtr) ElementType() reflect.Type {
	return namePtrType
}

func (in *namePtr) ToNamePtrOutput() NamePtrOutput {
	return pulumi.ToOutput(in).(NamePtrOutput)
}

func (in *namePtr) ToNamePtrOutputWithContext(ctx context.Context) NamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NamePtrOutput)
}

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

func (Permissions) ElementType() reflect.Type {
	return reflect.TypeOf((*Permissions)(nil)).Elem()
}

func (e Permissions) ToPermissionsOutput() PermissionsOutput {
	return pulumi.ToOutput(e).(PermissionsOutput)
}

func (e Permissions) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PermissionsOutput)
}

func (e Permissions) ToPermissionsPtrOutput() PermissionsPtrOutput {
	return e.ToPermissionsPtrOutputWithContext(context.Background())
}

func (e Permissions) ToPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPtrOutput {
	return Permissions(e).ToPermissionsOutputWithContext(ctx).ToPermissionsPtrOutputWithContext(ctx)
}

func (e Permissions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Permissions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Permissions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Permissions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PermissionsOutput struct{ *pulumi.OutputState }

func (PermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Permissions)(nil)).Elem()
}

func (o PermissionsOutput) ToPermissionsOutput() PermissionsOutput {
	return o
}

func (o PermissionsOutput) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return o
}

func (o PermissionsOutput) ToPermissionsPtrOutput() PermissionsPtrOutput {
	return o.ToPermissionsPtrOutputWithContext(context.Background())
}

func (o PermissionsOutput) ToPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Permissions) *Permissions {
		return &v
	}).(PermissionsPtrOutput)
}

func (o PermissionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PermissionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Permissions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PermissionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PermissionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Permissions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PermissionsPtrOutput struct{ *pulumi.OutputState }

func (PermissionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Permissions)(nil)).Elem()
}

func (o PermissionsPtrOutput) ToPermissionsPtrOutput() PermissionsPtrOutput {
	return o
}

func (o PermissionsPtrOutput) ToPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPtrOutput {
	return o
}

func (o PermissionsPtrOutput) Elem() PermissionsOutput {
	return o.ApplyT(func(v *Permissions) Permissions {
		if v != nil {
			return *v
		}
		var ret Permissions
		return ret
	}).(PermissionsOutput)
}

func (o PermissionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PermissionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Permissions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PermissionsInput interface {
	pulumi.Input

	ToPermissionsOutput() PermissionsOutput
	ToPermissionsOutputWithContext(context.Context) PermissionsOutput
}

var permissionsPtrType = reflect.TypeOf((**Permissions)(nil)).Elem()

type PermissionsPtrInput interface {
	pulumi.Input

	ToPermissionsPtrOutput() PermissionsPtrOutput
	ToPermissionsPtrOutputWithContext(context.Context) PermissionsPtrOutput
}

type permissionsPtr string

func PermissionsPtr(v string) PermissionsPtrInput {
	return (*permissionsPtr)(&v)
}

func (*permissionsPtr) ElementType() reflect.Type {
	return permissionsPtrType
}

func (in *permissionsPtr) ToPermissionsPtrOutput() PermissionsPtrOutput {
	return pulumi.ToOutput(in).(PermissionsPtrOutput)
}

func (in *permissionsPtr) ToPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PermissionsPtrOutput)
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

func (RootSquashType) ElementType() reflect.Type {
	return reflect.TypeOf((*RootSquashType)(nil)).Elem()
}

func (e RootSquashType) ToRootSquashTypeOutput() RootSquashTypeOutput {
	return pulumi.ToOutput(e).(RootSquashTypeOutput)
}

func (e RootSquashType) ToRootSquashTypeOutputWithContext(ctx context.Context) RootSquashTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RootSquashTypeOutput)
}

func (e RootSquashType) ToRootSquashTypePtrOutput() RootSquashTypePtrOutput {
	return e.ToRootSquashTypePtrOutputWithContext(context.Background())
}

func (e RootSquashType) ToRootSquashTypePtrOutputWithContext(ctx context.Context) RootSquashTypePtrOutput {
	return RootSquashType(e).ToRootSquashTypeOutputWithContext(ctx).ToRootSquashTypePtrOutputWithContext(ctx)
}

func (e RootSquashType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RootSquashType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RootSquashType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RootSquashType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RootSquashTypeOutput struct{ *pulumi.OutputState }

func (RootSquashTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RootSquashType)(nil)).Elem()
}

func (o RootSquashTypeOutput) ToRootSquashTypeOutput() RootSquashTypeOutput {
	return o
}

func (o RootSquashTypeOutput) ToRootSquashTypeOutputWithContext(ctx context.Context) RootSquashTypeOutput {
	return o
}

func (o RootSquashTypeOutput) ToRootSquashTypePtrOutput() RootSquashTypePtrOutput {
	return o.ToRootSquashTypePtrOutputWithContext(context.Background())
}

func (o RootSquashTypeOutput) ToRootSquashTypePtrOutputWithContext(ctx context.Context) RootSquashTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RootSquashType) *RootSquashType {
		return &v
	}).(RootSquashTypePtrOutput)
}

func (o RootSquashTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RootSquashTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RootSquashType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RootSquashTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RootSquashTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RootSquashType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RootSquashTypePtrOutput struct{ *pulumi.OutputState }

func (RootSquashTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RootSquashType)(nil)).Elem()
}

func (o RootSquashTypePtrOutput) ToRootSquashTypePtrOutput() RootSquashTypePtrOutput {
	return o
}

func (o RootSquashTypePtrOutput) ToRootSquashTypePtrOutputWithContext(ctx context.Context) RootSquashTypePtrOutput {
	return o
}

func (o RootSquashTypePtrOutput) Elem() RootSquashTypeOutput {
	return o.ApplyT(func(v *RootSquashType) RootSquashType {
		if v != nil {
			return *v
		}
		var ret RootSquashType
		return ret
	}).(RootSquashTypeOutput)
}

func (o RootSquashTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RootSquashTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RootSquashType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RootSquashTypeInput interface {
	pulumi.Input

	ToRootSquashTypeOutput() RootSquashTypeOutput
	ToRootSquashTypeOutputWithContext(context.Context) RootSquashTypeOutput
}

var rootSquashTypePtrType = reflect.TypeOf((**RootSquashType)(nil)).Elem()

type RootSquashTypePtrInput interface {
	pulumi.Input

	ToRootSquashTypePtrOutput() RootSquashTypePtrOutput
	ToRootSquashTypePtrOutputWithContext(context.Context) RootSquashTypePtrOutput
}

type rootSquashTypePtr string

func RootSquashTypePtr(v string) RootSquashTypePtrInput {
	return (*rootSquashTypePtr)(&v)
}

func (*rootSquashTypePtr) ElementType() reflect.Type {
	return rootSquashTypePtrType
}

func (in *rootSquashTypePtr) ToRootSquashTypePtrOutput() RootSquashTypePtrOutput {
	return pulumi.ToOutput(in).(RootSquashTypePtrOutput)
}

func (in *rootSquashTypePtr) ToRootSquashTypePtrOutputWithContext(ctx context.Context) RootSquashTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RootSquashTypePtrOutput)
}

type RoutingChoice string

const (
	RoutingChoiceMicrosoftRouting = RoutingChoice("MicrosoftRouting")
	RoutingChoiceInternetRouting  = RoutingChoice("InternetRouting")
)

func (RoutingChoice) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingChoice)(nil)).Elem()
}

func (e RoutingChoice) ToRoutingChoiceOutput() RoutingChoiceOutput {
	return pulumi.ToOutput(e).(RoutingChoiceOutput)
}

func (e RoutingChoice) ToRoutingChoiceOutputWithContext(ctx context.Context) RoutingChoiceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RoutingChoiceOutput)
}

func (e RoutingChoice) ToRoutingChoicePtrOutput() RoutingChoicePtrOutput {
	return e.ToRoutingChoicePtrOutputWithContext(context.Background())
}

func (e RoutingChoice) ToRoutingChoicePtrOutputWithContext(ctx context.Context) RoutingChoicePtrOutput {
	return RoutingChoice(e).ToRoutingChoiceOutputWithContext(ctx).ToRoutingChoicePtrOutputWithContext(ctx)
}

func (e RoutingChoice) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoutingChoice) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoutingChoice) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RoutingChoice) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RoutingChoiceOutput struct{ *pulumi.OutputState }

func (RoutingChoiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingChoice)(nil)).Elem()
}

func (o RoutingChoiceOutput) ToRoutingChoiceOutput() RoutingChoiceOutput {
	return o
}

func (o RoutingChoiceOutput) ToRoutingChoiceOutputWithContext(ctx context.Context) RoutingChoiceOutput {
	return o
}

func (o RoutingChoiceOutput) ToRoutingChoicePtrOutput() RoutingChoicePtrOutput {
	return o.ToRoutingChoicePtrOutputWithContext(context.Background())
}

func (o RoutingChoiceOutput) ToRoutingChoicePtrOutputWithContext(ctx context.Context) RoutingChoicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingChoice) *RoutingChoice {
		return &v
	}).(RoutingChoicePtrOutput)
}

func (o RoutingChoiceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RoutingChoiceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoutingChoice) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RoutingChoiceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoutingChoiceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoutingChoice) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RoutingChoicePtrOutput struct{ *pulumi.OutputState }

func (RoutingChoicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingChoice)(nil)).Elem()
}

func (o RoutingChoicePtrOutput) ToRoutingChoicePtrOutput() RoutingChoicePtrOutput {
	return o
}

func (o RoutingChoicePtrOutput) ToRoutingChoicePtrOutputWithContext(ctx context.Context) RoutingChoicePtrOutput {
	return o
}

func (o RoutingChoicePtrOutput) Elem() RoutingChoiceOutput {
	return o.ApplyT(func(v *RoutingChoice) RoutingChoice {
		if v != nil {
			return *v
		}
		var ret RoutingChoice
		return ret
	}).(RoutingChoiceOutput)
}

func (o RoutingChoicePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoutingChoicePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RoutingChoice) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RoutingChoiceInput interface {
	pulumi.Input

	ToRoutingChoiceOutput() RoutingChoiceOutput
	ToRoutingChoiceOutputWithContext(context.Context) RoutingChoiceOutput
}

var routingChoicePtrType = reflect.TypeOf((**RoutingChoice)(nil)).Elem()

type RoutingChoicePtrInput interface {
	pulumi.Input

	ToRoutingChoicePtrOutput() RoutingChoicePtrOutput
	ToRoutingChoicePtrOutputWithContext(context.Context) RoutingChoicePtrOutput
}

type routingChoicePtr string

func RoutingChoicePtr(v string) RoutingChoicePtrInput {
	return (*routingChoicePtr)(&v)
}

func (*routingChoicePtr) ElementType() reflect.Type {
	return routingChoicePtrType
}

func (in *routingChoicePtr) ToRoutingChoicePtrOutput() RoutingChoicePtrOutput {
	return pulumi.ToOutput(in).(RoutingChoicePtrOutput)
}

func (in *routingChoicePtr) ToRoutingChoicePtrOutputWithContext(ctx context.Context) RoutingChoicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RoutingChoicePtrOutput)
}

type RuleType string

const (
	RuleTypeLifecycle = RuleType("Lifecycle")
)

func (RuleType) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleType)(nil)).Elem()
}

func (e RuleType) ToRuleTypeOutput() RuleTypeOutput {
	return pulumi.ToOutput(e).(RuleTypeOutput)
}

func (e RuleType) ToRuleTypeOutputWithContext(ctx context.Context) RuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RuleTypeOutput)
}

func (e RuleType) ToRuleTypePtrOutput() RuleTypePtrOutput {
	return e.ToRuleTypePtrOutputWithContext(context.Background())
}

func (e RuleType) ToRuleTypePtrOutputWithContext(ctx context.Context) RuleTypePtrOutput {
	return RuleType(e).ToRuleTypeOutputWithContext(ctx).ToRuleTypePtrOutputWithContext(ctx)
}

func (e RuleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RuleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RuleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RuleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RuleTypeOutput struct{ *pulumi.OutputState }

func (RuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleType)(nil)).Elem()
}

func (o RuleTypeOutput) ToRuleTypeOutput() RuleTypeOutput {
	return o
}

func (o RuleTypeOutput) ToRuleTypeOutputWithContext(ctx context.Context) RuleTypeOutput {
	return o
}

func (o RuleTypeOutput) ToRuleTypePtrOutput() RuleTypePtrOutput {
	return o.ToRuleTypePtrOutputWithContext(context.Background())
}

func (o RuleTypeOutput) ToRuleTypePtrOutputWithContext(ctx context.Context) RuleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleType) *RuleType {
		return &v
	}).(RuleTypePtrOutput)
}

func (o RuleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RuleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RuleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RuleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RuleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RuleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RuleTypePtrOutput struct{ *pulumi.OutputState }

func (RuleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleType)(nil)).Elem()
}

func (o RuleTypePtrOutput) ToRuleTypePtrOutput() RuleTypePtrOutput {
	return o
}

func (o RuleTypePtrOutput) ToRuleTypePtrOutputWithContext(ctx context.Context) RuleTypePtrOutput {
	return o
}

func (o RuleTypePtrOutput) Elem() RuleTypeOutput {
	return o.ApplyT(func(v *RuleType) RuleType {
		if v != nil {
			return *v
		}
		var ret RuleType
		return ret
	}).(RuleTypeOutput)
}

func (o RuleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RuleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RuleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RuleTypeInput interface {
	pulumi.Input

	ToRuleTypeOutput() RuleTypeOutput
	ToRuleTypeOutputWithContext(context.Context) RuleTypeOutput
}

var ruleTypePtrType = reflect.TypeOf((**RuleType)(nil)).Elem()

type RuleTypePtrInput interface {
	pulumi.Input

	ToRuleTypePtrOutput() RuleTypePtrOutput
	ToRuleTypePtrOutputWithContext(context.Context) RuleTypePtrOutput
}

type ruleTypePtr string

func RuleTypePtr(v string) RuleTypePtrInput {
	return (*ruleTypePtr)(&v)
}

func (*ruleTypePtr) ElementType() reflect.Type {
	return ruleTypePtrType
}

func (in *ruleTypePtr) ToRuleTypePtrOutput() RuleTypePtrOutput {
	return pulumi.ToOutput(in).(RuleTypePtrOutput)
}

func (in *ruleTypePtr) ToRuleTypePtrOutputWithContext(ctx context.Context) RuleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RuleTypePtrOutput)
}

type Services string

const (
	ServicesB = Services("b")
	ServicesQ = Services("q")
	ServicesT = Services("t")
	ServicesF = Services("f")
)

func (Services) ElementType() reflect.Type {
	return reflect.TypeOf((*Services)(nil)).Elem()
}

func (e Services) ToServicesOutput() ServicesOutput {
	return pulumi.ToOutput(e).(ServicesOutput)
}

func (e Services) ToServicesOutputWithContext(ctx context.Context) ServicesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServicesOutput)
}

func (e Services) ToServicesPtrOutput() ServicesPtrOutput {
	return e.ToServicesPtrOutputWithContext(context.Background())
}

func (e Services) ToServicesPtrOutputWithContext(ctx context.Context) ServicesPtrOutput {
	return Services(e).ToServicesOutputWithContext(ctx).ToServicesPtrOutputWithContext(ctx)
}

func (e Services) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Services) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Services) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Services) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServicesOutput struct{ *pulumi.OutputState }

func (ServicesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Services)(nil)).Elem()
}

func (o ServicesOutput) ToServicesOutput() ServicesOutput {
	return o
}

func (o ServicesOutput) ToServicesOutputWithContext(ctx context.Context) ServicesOutput {
	return o
}

func (o ServicesOutput) ToServicesPtrOutput() ServicesPtrOutput {
	return o.ToServicesPtrOutputWithContext(context.Background())
}

func (o ServicesOutput) ToServicesPtrOutputWithContext(ctx context.Context) ServicesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Services) *Services {
		return &v
	}).(ServicesPtrOutput)
}

func (o ServicesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServicesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Services) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServicesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServicesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Services) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServicesPtrOutput struct{ *pulumi.OutputState }

func (ServicesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Services)(nil)).Elem()
}

func (o ServicesPtrOutput) ToServicesPtrOutput() ServicesPtrOutput {
	return o
}

func (o ServicesPtrOutput) ToServicesPtrOutputWithContext(ctx context.Context) ServicesPtrOutput {
	return o
}

func (o ServicesPtrOutput) Elem() ServicesOutput {
	return o.ApplyT(func(v *Services) Services {
		if v != nil {
			return *v
		}
		var ret Services
		return ret
	}).(ServicesOutput)
}

func (o ServicesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServicesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Services) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServicesInput interface {
	pulumi.Input

	ToServicesOutput() ServicesOutput
	ToServicesOutputWithContext(context.Context) ServicesOutput
}

var servicesPtrType = reflect.TypeOf((**Services)(nil)).Elem()

type ServicesPtrInput interface {
	pulumi.Input

	ToServicesPtrOutput() ServicesPtrOutput
	ToServicesPtrOutputWithContext(context.Context) ServicesPtrOutput
}

type servicesPtr string

func ServicesPtr(v string) ServicesPtrInput {
	return (*servicesPtr)(&v)
}

func (*servicesPtr) ElementType() reflect.Type {
	return servicesPtrType
}

func (in *servicesPtr) ToServicesPtrOutput() ServicesPtrOutput {
	return pulumi.ToOutput(in).(ServicesPtrOutput)
}

func (in *servicesPtr) ToServicesPtrOutputWithContext(ctx context.Context) ServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServicesPtrOutput)
}

type ShareAccessTier string

const (
	ShareAccessTierTransactionOptimized = ShareAccessTier("TransactionOptimized")
	ShareAccessTierHot                  = ShareAccessTier("Hot")
	ShareAccessTierCool                 = ShareAccessTier("Cool")
	ShareAccessTierPremium              = ShareAccessTier("Premium")
)

func (ShareAccessTier) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccessTier)(nil)).Elem()
}

func (e ShareAccessTier) ToShareAccessTierOutput() ShareAccessTierOutput {
	return pulumi.ToOutput(e).(ShareAccessTierOutput)
}

func (e ShareAccessTier) ToShareAccessTierOutputWithContext(ctx context.Context) ShareAccessTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ShareAccessTierOutput)
}

func (e ShareAccessTier) ToShareAccessTierPtrOutput() ShareAccessTierPtrOutput {
	return e.ToShareAccessTierPtrOutputWithContext(context.Background())
}

func (e ShareAccessTier) ToShareAccessTierPtrOutputWithContext(ctx context.Context) ShareAccessTierPtrOutput {
	return ShareAccessTier(e).ToShareAccessTierOutputWithContext(ctx).ToShareAccessTierPtrOutputWithContext(ctx)
}

func (e ShareAccessTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareAccessTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareAccessTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ShareAccessTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ShareAccessTierOutput struct{ *pulumi.OutputState }

func (ShareAccessTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccessTier)(nil)).Elem()
}

func (o ShareAccessTierOutput) ToShareAccessTierOutput() ShareAccessTierOutput {
	return o
}

func (o ShareAccessTierOutput) ToShareAccessTierOutputWithContext(ctx context.Context) ShareAccessTierOutput {
	return o
}

func (o ShareAccessTierOutput) ToShareAccessTierPtrOutput() ShareAccessTierPtrOutput {
	return o.ToShareAccessTierPtrOutputWithContext(context.Background())
}

func (o ShareAccessTierOutput) ToShareAccessTierPtrOutputWithContext(ctx context.Context) ShareAccessTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShareAccessTier) *ShareAccessTier {
		return &v
	}).(ShareAccessTierPtrOutput)
}

func (o ShareAccessTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ShareAccessTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShareAccessTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ShareAccessTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShareAccessTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShareAccessTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ShareAccessTierPtrOutput struct{ *pulumi.OutputState }

func (ShareAccessTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareAccessTier)(nil)).Elem()
}

func (o ShareAccessTierPtrOutput) ToShareAccessTierPtrOutput() ShareAccessTierPtrOutput {
	return o
}

func (o ShareAccessTierPtrOutput) ToShareAccessTierPtrOutputWithContext(ctx context.Context) ShareAccessTierPtrOutput {
	return o
}

func (o ShareAccessTierPtrOutput) Elem() ShareAccessTierOutput {
	return o.ApplyT(func(v *ShareAccessTier) ShareAccessTier {
		if v != nil {
			return *v
		}
		var ret ShareAccessTier
		return ret
	}).(ShareAccessTierOutput)
}

func (o ShareAccessTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShareAccessTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ShareAccessTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ShareAccessTierInput interface {
	pulumi.Input

	ToShareAccessTierOutput() ShareAccessTierOutput
	ToShareAccessTierOutputWithContext(context.Context) ShareAccessTierOutput
}

var shareAccessTierPtrType = reflect.TypeOf((**ShareAccessTier)(nil)).Elem()

type ShareAccessTierPtrInput interface {
	pulumi.Input

	ToShareAccessTierPtrOutput() ShareAccessTierPtrOutput
	ToShareAccessTierPtrOutputWithContext(context.Context) ShareAccessTierPtrOutput
}

type shareAccessTierPtr string

func ShareAccessTierPtr(v string) ShareAccessTierPtrInput {
	return (*shareAccessTierPtr)(&v)
}

func (*shareAccessTierPtr) ElementType() reflect.Type {
	return shareAccessTierPtrType
}

func (in *shareAccessTierPtr) ToShareAccessTierPtrOutput() ShareAccessTierPtrOutput {
	return pulumi.ToOutput(in).(ShareAccessTierPtrOutput)
}

func (in *shareAccessTierPtr) ToShareAccessTierPtrOutputWithContext(ctx context.Context) ShareAccessTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ShareAccessTierPtrOutput)
}

type SignedResource string

const (
	SignedResourceB = SignedResource("b")
	SignedResourceC = SignedResource("c")
	SignedResourceF = SignedResource("f")
	SignedResourceS = SignedResource("s")
)

func (SignedResource) ElementType() reflect.Type {
	return reflect.TypeOf((*SignedResource)(nil)).Elem()
}

func (e SignedResource) ToSignedResourceOutput() SignedResourceOutput {
	return pulumi.ToOutput(e).(SignedResourceOutput)
}

func (e SignedResource) ToSignedResourceOutputWithContext(ctx context.Context) SignedResourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SignedResourceOutput)
}

func (e SignedResource) ToSignedResourcePtrOutput() SignedResourcePtrOutput {
	return e.ToSignedResourcePtrOutputWithContext(context.Background())
}

func (e SignedResource) ToSignedResourcePtrOutputWithContext(ctx context.Context) SignedResourcePtrOutput {
	return SignedResource(e).ToSignedResourceOutputWithContext(ctx).ToSignedResourcePtrOutputWithContext(ctx)
}

func (e SignedResource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignedResource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignedResource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SignedResource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SignedResourceOutput struct{ *pulumi.OutputState }

func (SignedResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignedResource)(nil)).Elem()
}

func (o SignedResourceOutput) ToSignedResourceOutput() SignedResourceOutput {
	return o
}

func (o SignedResourceOutput) ToSignedResourceOutputWithContext(ctx context.Context) SignedResourceOutput {
	return o
}

func (o SignedResourceOutput) ToSignedResourcePtrOutput() SignedResourcePtrOutput {
	return o.ToSignedResourcePtrOutputWithContext(context.Background())
}

func (o SignedResourceOutput) ToSignedResourcePtrOutputWithContext(ctx context.Context) SignedResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignedResource) *SignedResource {
		return &v
	}).(SignedResourcePtrOutput)
}

func (o SignedResourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SignedResourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignedResource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SignedResourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignedResourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignedResource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SignedResourcePtrOutput struct{ *pulumi.OutputState }

func (SignedResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignedResource)(nil)).Elem()
}

func (o SignedResourcePtrOutput) ToSignedResourcePtrOutput() SignedResourcePtrOutput {
	return o
}

func (o SignedResourcePtrOutput) ToSignedResourcePtrOutputWithContext(ctx context.Context) SignedResourcePtrOutput {
	return o
}

func (o SignedResourcePtrOutput) Elem() SignedResourceOutput {
	return o.ApplyT(func(v *SignedResource) SignedResource {
		if v != nil {
			return *v
		}
		var ret SignedResource
		return ret
	}).(SignedResourceOutput)
}

func (o SignedResourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignedResourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SignedResource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SignedResourceInput interface {
	pulumi.Input

	ToSignedResourceOutput() SignedResourceOutput
	ToSignedResourceOutputWithContext(context.Context) SignedResourceOutput
}

var signedResourcePtrType = reflect.TypeOf((**SignedResource)(nil)).Elem()

type SignedResourcePtrInput interface {
	pulumi.Input

	ToSignedResourcePtrOutput() SignedResourcePtrOutput
	ToSignedResourcePtrOutputWithContext(context.Context) SignedResourcePtrOutput
}

type signedResourcePtr string

func SignedResourcePtr(v string) SignedResourcePtrInput {
	return (*signedResourcePtr)(&v)
}

func (*signedResourcePtr) ElementType() reflect.Type {
	return signedResourcePtrType
}

func (in *signedResourcePtr) ToSignedResourcePtrOutput() SignedResourcePtrOutput {
	return pulumi.ToOutput(in).(SignedResourcePtrOutput)
}

func (in *signedResourcePtr) ToSignedResourcePtrOutputWithContext(ctx context.Context) SignedResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SignedResourcePtrOutput)
}

type SignedResourceTypes string

const (
	SignedResourceTypesS = SignedResourceTypes("s")
	SignedResourceTypesC = SignedResourceTypes("c")
	SignedResourceTypesO = SignedResourceTypes("o")
)

func (SignedResourceTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*SignedResourceTypes)(nil)).Elem()
}

func (e SignedResourceTypes) ToSignedResourceTypesOutput() SignedResourceTypesOutput {
	return pulumi.ToOutput(e).(SignedResourceTypesOutput)
}

func (e SignedResourceTypes) ToSignedResourceTypesOutputWithContext(ctx context.Context) SignedResourceTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SignedResourceTypesOutput)
}

func (e SignedResourceTypes) ToSignedResourceTypesPtrOutput() SignedResourceTypesPtrOutput {
	return e.ToSignedResourceTypesPtrOutputWithContext(context.Background())
}

func (e SignedResourceTypes) ToSignedResourceTypesPtrOutputWithContext(ctx context.Context) SignedResourceTypesPtrOutput {
	return SignedResourceTypes(e).ToSignedResourceTypesOutputWithContext(ctx).ToSignedResourceTypesPtrOutputWithContext(ctx)
}

func (e SignedResourceTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignedResourceTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignedResourceTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SignedResourceTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SignedResourceTypesOutput struct{ *pulumi.OutputState }

func (SignedResourceTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignedResourceTypes)(nil)).Elem()
}

func (o SignedResourceTypesOutput) ToSignedResourceTypesOutput() SignedResourceTypesOutput {
	return o
}

func (o SignedResourceTypesOutput) ToSignedResourceTypesOutputWithContext(ctx context.Context) SignedResourceTypesOutput {
	return o
}

func (o SignedResourceTypesOutput) ToSignedResourceTypesPtrOutput() SignedResourceTypesPtrOutput {
	return o.ToSignedResourceTypesPtrOutputWithContext(context.Background())
}

func (o SignedResourceTypesOutput) ToSignedResourceTypesPtrOutputWithContext(ctx context.Context) SignedResourceTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignedResourceTypes) *SignedResourceTypes {
		return &v
	}).(SignedResourceTypesPtrOutput)
}

func (o SignedResourceTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SignedResourceTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignedResourceTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SignedResourceTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignedResourceTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignedResourceTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SignedResourceTypesPtrOutput struct{ *pulumi.OutputState }

func (SignedResourceTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignedResourceTypes)(nil)).Elem()
}

func (o SignedResourceTypesPtrOutput) ToSignedResourceTypesPtrOutput() SignedResourceTypesPtrOutput {
	return o
}

func (o SignedResourceTypesPtrOutput) ToSignedResourceTypesPtrOutputWithContext(ctx context.Context) SignedResourceTypesPtrOutput {
	return o
}

func (o SignedResourceTypesPtrOutput) Elem() SignedResourceTypesOutput {
	return o.ApplyT(func(v *SignedResourceTypes) SignedResourceTypes {
		if v != nil {
			return *v
		}
		var ret SignedResourceTypes
		return ret
	}).(SignedResourceTypesOutput)
}

func (o SignedResourceTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignedResourceTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SignedResourceTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SignedResourceTypesInput interface {
	pulumi.Input

	ToSignedResourceTypesOutput() SignedResourceTypesOutput
	ToSignedResourceTypesOutputWithContext(context.Context) SignedResourceTypesOutput
}

var signedResourceTypesPtrType = reflect.TypeOf((**SignedResourceTypes)(nil)).Elem()

type SignedResourceTypesPtrInput interface {
	pulumi.Input

	ToSignedResourceTypesPtrOutput() SignedResourceTypesPtrOutput
	ToSignedResourceTypesPtrOutputWithContext(context.Context) SignedResourceTypesPtrOutput
}

type signedResourceTypesPtr string

func SignedResourceTypesPtr(v string) SignedResourceTypesPtrInput {
	return (*signedResourceTypesPtr)(&v)
}

func (*signedResourceTypesPtr) ElementType() reflect.Type {
	return signedResourceTypesPtrType
}

func (in *signedResourceTypesPtr) ToSignedResourceTypesPtrOutput() SignedResourceTypesPtrOutput {
	return pulumi.ToOutput(in).(SignedResourceTypesPtrOutput)
}

func (in *signedResourceTypesPtr) ToSignedResourceTypesPtrOutputWithContext(ctx context.Context) SignedResourceTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SignedResourceTypesPtrOutput)
}

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

type State string

const (
	StateProvisioning         = State("provisioning")
	StateDeprovisioning       = State("deprovisioning")
	StateSucceeded            = State("succeeded")
	StateFailed               = State("failed")
	StateNetworkSourceDeleted = State("networkSourceDeleted")
)

func (State) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (e State) ToStateOutput() StateOutput {
	return pulumi.ToOutput(e).(StateOutput)
}

func (e State) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StateOutput)
}

func (e State) ToStatePtrOutput() StatePtrOutput {
	return e.ToStatePtrOutputWithContext(context.Background())
}

func (e State) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return State(e).ToStateOutputWithContext(ctx).ToStatePtrOutputWithContext(ctx)
}

func (e State) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e State) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StateOutput struct{ *pulumi.OutputState }

func (StateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (o StateOutput) ToStateOutput() StateOutput {
	return o
}

func (o StateOutput) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return o
}

func (o StateOutput) ToStatePtrOutput() StatePtrOutput {
	return o.ToStatePtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v State) *State {
		return &v
	}).(StatePtrOutput)
}

func (o StateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatePtrOutput struct{ *pulumi.OutputState }

func (StatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**State)(nil)).Elem()
}

func (o StatePtrOutput) ToStatePtrOutput() StatePtrOutput {
	return o
}

func (o StatePtrOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o
}

func (o StatePtrOutput) Elem() StateOutput {
	return o.ApplyT(func(v *State) State {
		if v != nil {
			return *v
		}
		var ret State
		return ret
	}).(StateOutput)
}

func (o StatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *State) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StateInput interface {
	pulumi.Input

	ToStateOutput() StateOutput
	ToStateOutputWithContext(context.Context) StateOutput
}

var statePtrType = reflect.TypeOf((**State)(nil)).Elem()

type StatePtrInput interface {
	pulumi.Input

	ToStatePtrOutput() StatePtrOutput
	ToStatePtrOutputWithContext(context.Context) StatePtrOutput
}

type statePtr string

func StatePtr(v string) StatePtrInput {
	return (*statePtr)(&v)
}

func (*statePtr) ElementType() reflect.Type {
	return statePtrType
}

func (in *statePtr) ToStatePtrOutput() StatePtrOutput {
	return pulumi.ToOutput(in).(StatePtrOutput)
}

func (in *statePtr) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessTierInput)(nil)).Elem(), AccessTier("Hot"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessTierPtrInput)(nil)).Elem(), AccessTier("Hot"))
	pulumi.RegisterInputType(reflect.TypeOf((*ActionInput)(nil)).Elem(), Action("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*ActionPtrInput)(nil)).Elem(), Action("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*BlobAccessTierInput)(nil)).Elem(), BlobAccessTier("Hot"))
	pulumi.RegisterInputType(reflect.TypeOf((*BlobAccessTierPtrInput)(nil)).Elem(), BlobAccessTier("Hot"))
	pulumi.RegisterInputType(reflect.TypeOf((*BlobTypeInput)(nil)).Elem(), BlobType("Block"))
	pulumi.RegisterInputType(reflect.TypeOf((*BlobTypePtrInput)(nil)).Elem(), BlobType("Block"))
	pulumi.RegisterInputType(reflect.TypeOf((*BypassInput)(nil)).Elem(), Bypass("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*BypassPtrInput)(nil)).Elem(), Bypass("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultActionInput)(nil)).Elem(), DefaultAction("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultActionPtrInput)(nil)).Elem(), DefaultAction("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryServiceOptionsInput)(nil)).Elem(), DirectoryServiceOptions("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryServiceOptionsPtrInput)(nil)).Elem(), DirectoryServiceOptions("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledProtocolsInput)(nil)).Elem(), EnabledProtocols("SMB"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledProtocolsPtrInput)(nil)).Elem(), EnabledProtocols("SMB"))
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionScopeSourceInput)(nil)).Elem(), EncryptionScopeSource("Microsoft.Storage"))
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionScopeSourcePtrInput)(nil)).Elem(), EncryptionScopeSource("Microsoft.Storage"))
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionScopeStateEnumInput)(nil)).Elem(), EncryptionScopeStateEnum("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionScopeStateEnumPtrInput)(nil)).Elem(), EncryptionScopeStateEnum("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExpirationActionInput)(nil)).Elem(), ExpirationAction("Log"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExpirationActionPtrInput)(nil)).Elem(), ExpirationAction("Log"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationTypesInput)(nil)).Elem(), ExtendedLocationTypes("EdgeZone"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationTypesPtrInput)(nil)).Elem(), ExtendedLocationTypes("EdgeZone"))
	pulumi.RegisterInputType(reflect.TypeOf((*HttpProtocolInput)(nil)).Elem(), HttpProtocol("https,http"))
	pulumi.RegisterInputType(reflect.TypeOf((*HttpProtocolPtrInput)(nil)).Elem(), HttpProtocol("https,http"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypeInput)(nil)).Elem(), IdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypePtrInput)(nil)).Elem(), IdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*InventoryRuleTypeInput)(nil)).Elem(), InventoryRuleType("Inventory"))
	pulumi.RegisterInputType(reflect.TypeOf((*InventoryRuleTypePtrInput)(nil)).Elem(), InventoryRuleType("Inventory"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeySourceInput)(nil)).Elem(), KeySource("Microsoft.Storage"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeySourcePtrInput)(nil)).Elem(), KeySource("Microsoft.Storage"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeyTypeInput)(nil)).Elem(), KeyType("Service"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeyTypePtrInput)(nil)).Elem(), KeyType("Service"))
	pulumi.RegisterInputType(reflect.TypeOf((*KindInput)(nil)).Elem(), Kind("Storage"))
	pulumi.RegisterInputType(reflect.TypeOf((*KindPtrInput)(nil)).Elem(), Kind("Storage"))
	pulumi.RegisterInputType(reflect.TypeOf((*LargeFileSharesStateInput)(nil)).Elem(), LargeFileSharesState("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*LargeFileSharesStatePtrInput)(nil)).Elem(), LargeFileSharesState("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*MinimumTlsVersionInput)(nil)).Elem(), MinimumTlsVersion("TLS1_0"))
	pulumi.RegisterInputType(reflect.TypeOf((*MinimumTlsVersionPtrInput)(nil)).Elem(), MinimumTlsVersion("TLS1_0"))
	pulumi.RegisterInputType(reflect.TypeOf((*NameInput)(nil)).Elem(), Name("AccessTimeTracking"))
	pulumi.RegisterInputType(reflect.TypeOf((*NamePtrInput)(nil)).Elem(), Name("AccessTimeTracking"))
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsInput)(nil)).Elem(), Permissions("r"))
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsPtrInput)(nil)).Elem(), Permissions("r"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointServiceConnectionStatusInput)(nil)).Elem(), PrivateEndpointServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointServiceConnectionStatusPtrInput)(nil)).Elem(), PrivateEndpointServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicAccessInput)(nil)).Elem(), PublicAccess("Container"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicAccessPtrInput)(nil)).Elem(), PublicAccess("Container"))
	pulumi.RegisterInputType(reflect.TypeOf((*RootSquashTypeInput)(nil)).Elem(), RootSquashType("NoRootSquash"))
	pulumi.RegisterInputType(reflect.TypeOf((*RootSquashTypePtrInput)(nil)).Elem(), RootSquashType("NoRootSquash"))
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingChoiceInput)(nil)).Elem(), RoutingChoice("MicrosoftRouting"))
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingChoicePtrInput)(nil)).Elem(), RoutingChoice("MicrosoftRouting"))
	pulumi.RegisterInputType(reflect.TypeOf((*RuleTypeInput)(nil)).Elem(), RuleType("Lifecycle"))
	pulumi.RegisterInputType(reflect.TypeOf((*RuleTypePtrInput)(nil)).Elem(), RuleType("Lifecycle"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServicesInput)(nil)).Elem(), Services("b"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServicesPtrInput)(nil)).Elem(), Services("b"))
	pulumi.RegisterInputType(reflect.TypeOf((*ShareAccessTierInput)(nil)).Elem(), ShareAccessTier("TransactionOptimized"))
	pulumi.RegisterInputType(reflect.TypeOf((*ShareAccessTierPtrInput)(nil)).Elem(), ShareAccessTier("TransactionOptimized"))
	pulumi.RegisterInputType(reflect.TypeOf((*SignedResourceInput)(nil)).Elem(), SignedResource("b"))
	pulumi.RegisterInputType(reflect.TypeOf((*SignedResourcePtrInput)(nil)).Elem(), SignedResource("b"))
	pulumi.RegisterInputType(reflect.TypeOf((*SignedResourceTypesInput)(nil)).Elem(), SignedResourceTypes("s"))
	pulumi.RegisterInputType(reflect.TypeOf((*SignedResourceTypesPtrInput)(nil)).Elem(), SignedResourceTypes("s"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNameInput)(nil)).Elem(), SkuName("Standard_LRS"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNamePtrInput)(nil)).Elem(), SkuName("Standard_LRS"))
	pulumi.RegisterInputType(reflect.TypeOf((*StateInput)(nil)).Elem(), State("provisioning"))
	pulumi.RegisterInputType(reflect.TypeOf((*StatePtrInput)(nil)).Elem(), State("provisioning"))
	pulumi.RegisterOutputType(AccessTierOutput{})
	pulumi.RegisterOutputType(AccessTierPtrOutput{})
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(BlobAccessTierOutput{})
	pulumi.RegisterOutputType(BlobAccessTierPtrOutput{})
	pulumi.RegisterOutputType(BlobTypeOutput{})
	pulumi.RegisterOutputType(BlobTypePtrOutput{})
	pulumi.RegisterOutputType(BypassOutput{})
	pulumi.RegisterOutputType(BypassPtrOutput{})
	pulumi.RegisterOutputType(DefaultActionOutput{})
	pulumi.RegisterOutputType(DefaultActionPtrOutput{})
	pulumi.RegisterOutputType(DirectoryServiceOptionsOutput{})
	pulumi.RegisterOutputType(DirectoryServiceOptionsPtrOutput{})
	pulumi.RegisterOutputType(EnabledProtocolsOutput{})
	pulumi.RegisterOutputType(EnabledProtocolsPtrOutput{})
	pulumi.RegisterOutputType(EncryptionScopeSourceOutput{})
	pulumi.RegisterOutputType(EncryptionScopeSourcePtrOutput{})
	pulumi.RegisterOutputType(EncryptionScopeStateEnumOutput{})
	pulumi.RegisterOutputType(EncryptionScopeStateEnumPtrOutput{})
	pulumi.RegisterOutputType(ExpirationActionOutput{})
	pulumi.RegisterOutputType(ExpirationActionPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypesOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypesPtrOutput{})
	pulumi.RegisterOutputType(HttpProtocolOutput{})
	pulumi.RegisterOutputType(HttpProtocolPtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(InventoryRuleTypeOutput{})
	pulumi.RegisterOutputType(InventoryRuleTypePtrOutput{})
	pulumi.RegisterOutputType(KeySourceOutput{})
	pulumi.RegisterOutputType(KeySourcePtrOutput{})
	pulumi.RegisterOutputType(KeyTypeOutput{})
	pulumi.RegisterOutputType(KeyTypePtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
	pulumi.RegisterOutputType(LargeFileSharesStateOutput{})
	pulumi.RegisterOutputType(LargeFileSharesStatePtrOutput{})
	pulumi.RegisterOutputType(MinimumTlsVersionOutput{})
	pulumi.RegisterOutputType(MinimumTlsVersionPtrOutput{})
	pulumi.RegisterOutputType(NameOutput{})
	pulumi.RegisterOutputType(NamePtrOutput{})
	pulumi.RegisterOutputType(PermissionsOutput{})
	pulumi.RegisterOutputType(PermissionsPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicAccessOutput{})
	pulumi.RegisterOutputType(PublicAccessPtrOutput{})
	pulumi.RegisterOutputType(RootSquashTypeOutput{})
	pulumi.RegisterOutputType(RootSquashTypePtrOutput{})
	pulumi.RegisterOutputType(RoutingChoiceOutput{})
	pulumi.RegisterOutputType(RoutingChoicePtrOutput{})
	pulumi.RegisterOutputType(RuleTypeOutput{})
	pulumi.RegisterOutputType(RuleTypePtrOutput{})
	pulumi.RegisterOutputType(ServicesOutput{})
	pulumi.RegisterOutputType(ServicesPtrOutput{})
	pulumi.RegisterOutputType(ShareAccessTierOutput{})
	pulumi.RegisterOutputType(ShareAccessTierPtrOutput{})
	pulumi.RegisterOutputType(SignedResourceOutput{})
	pulumi.RegisterOutputType(SignedResourcePtrOutput{})
	pulumi.RegisterOutputType(SignedResourceTypesOutput{})
	pulumi.RegisterOutputType(SignedResourceTypesPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(StateOutput{})
	pulumi.RegisterOutputType(StatePtrOutput{})
}
