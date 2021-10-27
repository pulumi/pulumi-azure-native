


package v20190401

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
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
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

type ReasonCode string

const (
	ReasonCodeQuotaId                     = ReasonCode("QuotaId")
	ReasonCodeNotAvailableForSubscription = ReasonCode("NotAvailableForSubscription")
)

func (ReasonCode) ElementType() reflect.Type {
	return reflect.TypeOf((*ReasonCode)(nil)).Elem()
}

func (e ReasonCode) ToReasonCodeOutput() ReasonCodeOutput {
	return pulumi.ToOutput(e).(ReasonCodeOutput)
}

func (e ReasonCode) ToReasonCodeOutputWithContext(ctx context.Context) ReasonCodeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReasonCodeOutput)
}

func (e ReasonCode) ToReasonCodePtrOutput() ReasonCodePtrOutput {
	return e.ToReasonCodePtrOutputWithContext(context.Background())
}

func (e ReasonCode) ToReasonCodePtrOutputWithContext(ctx context.Context) ReasonCodePtrOutput {
	return ReasonCode(e).ToReasonCodeOutputWithContext(ctx).ToReasonCodePtrOutputWithContext(ctx)
}

func (e ReasonCode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReasonCode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReasonCode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReasonCode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReasonCodeOutput struct{ *pulumi.OutputState }

func (ReasonCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReasonCode)(nil)).Elem()
}

func (o ReasonCodeOutput) ToReasonCodeOutput() ReasonCodeOutput {
	return o
}

func (o ReasonCodeOutput) ToReasonCodeOutputWithContext(ctx context.Context) ReasonCodeOutput {
	return o
}

func (o ReasonCodeOutput) ToReasonCodePtrOutput() ReasonCodePtrOutput {
	return o.ToReasonCodePtrOutputWithContext(context.Background())
}

func (o ReasonCodeOutput) ToReasonCodePtrOutputWithContext(ctx context.Context) ReasonCodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReasonCode) *ReasonCode {
		return &v
	}).(ReasonCodePtrOutput)
}

func (o ReasonCodeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReasonCodeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReasonCode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReasonCodeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReasonCodeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReasonCode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReasonCodePtrOutput struct{ *pulumi.OutputState }

func (ReasonCodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReasonCode)(nil)).Elem()
}

func (o ReasonCodePtrOutput) ToReasonCodePtrOutput() ReasonCodePtrOutput {
	return o
}

func (o ReasonCodePtrOutput) ToReasonCodePtrOutputWithContext(ctx context.Context) ReasonCodePtrOutput {
	return o
}

func (o ReasonCodePtrOutput) Elem() ReasonCodeOutput {
	return o.ApplyT(func(v *ReasonCode) ReasonCode {
		if v != nil {
			return *v
		}
		var ret ReasonCode
		return ret
	}).(ReasonCodeOutput)
}

func (o ReasonCodePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReasonCodePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReasonCode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ReasonCodeInput interface {
	pulumi.Input

	ToReasonCodeOutput() ReasonCodeOutput
	ToReasonCodeOutputWithContext(context.Context) ReasonCodeOutput
}

var reasonCodePtrType = reflect.TypeOf((**ReasonCode)(nil)).Elem()

type ReasonCodePtrInput interface {
	pulumi.Input

	ToReasonCodePtrOutput() ReasonCodePtrOutput
	ToReasonCodePtrOutputWithContext(context.Context) ReasonCodePtrOutput
}

type reasonCodePtr string

func ReasonCodePtr(v string) ReasonCodePtrInput {
	return (*reasonCodePtr)(&v)
}

func (*reasonCodePtr) ElementType() reflect.Type {
	return reasonCodePtrType
}

func (in *reasonCodePtr) ToReasonCodePtrOutput() ReasonCodePtrOutput {
	return pulumi.ToOutput(in).(ReasonCodePtrOutput)
}

func (in *reasonCodePtr) ToReasonCodePtrOutputWithContext(ctx context.Context) ReasonCodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReasonCodePtrOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*BypassInput)(nil)).Elem(), Bypass("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*BypassPtrInput)(nil)).Elem(), Bypass("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultActionInput)(nil)).Elem(), DefaultAction("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultActionPtrInput)(nil)).Elem(), DefaultAction("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryServiceOptionsInput)(nil)).Elem(), DirectoryServiceOptions("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryServiceOptionsPtrInput)(nil)).Elem(), DirectoryServiceOptions("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*HttpProtocolInput)(nil)).Elem(), HttpProtocol("https,http"))
	pulumi.RegisterInputType(reflect.TypeOf((*HttpProtocolPtrInput)(nil)).Elem(), HttpProtocol("https,http"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypeInput)(nil)).Elem(), IdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypePtrInput)(nil)).Elem(), IdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeySourceInput)(nil)).Elem(), KeySource("Microsoft.Storage"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeySourcePtrInput)(nil)).Elem(), KeySource("Microsoft.Storage"))
	pulumi.RegisterInputType(reflect.TypeOf((*KindInput)(nil)).Elem(), Kind("Storage"))
	pulumi.RegisterInputType(reflect.TypeOf((*KindPtrInput)(nil)).Elem(), Kind("Storage"))
	pulumi.RegisterInputType(reflect.TypeOf((*LargeFileSharesStateInput)(nil)).Elem(), LargeFileSharesState("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*LargeFileSharesStatePtrInput)(nil)).Elem(), LargeFileSharesState("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*MinimumTlsVersionInput)(nil)).Elem(), MinimumTlsVersion("TLS1_0"))
	pulumi.RegisterInputType(reflect.TypeOf((*MinimumTlsVersionPtrInput)(nil)).Elem(), MinimumTlsVersion("TLS1_0"))
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsInput)(nil)).Elem(), Permissions("r"))
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsPtrInput)(nil)).Elem(), Permissions("r"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicAccessInput)(nil)).Elem(), PublicAccess("Container"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicAccessPtrInput)(nil)).Elem(), PublicAccess("Container"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReasonCodeInput)(nil)).Elem(), ReasonCode("QuotaId"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReasonCodePtrInput)(nil)).Elem(), ReasonCode("QuotaId"))
	pulumi.RegisterInputType(reflect.TypeOf((*RuleTypeInput)(nil)).Elem(), RuleType("Lifecycle"))
	pulumi.RegisterInputType(reflect.TypeOf((*RuleTypePtrInput)(nil)).Elem(), RuleType("Lifecycle"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServicesInput)(nil)).Elem(), Services("b"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServicesPtrInput)(nil)).Elem(), Services("b"))
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
	pulumi.RegisterOutputType(BypassOutput{})
	pulumi.RegisterOutputType(BypassPtrOutput{})
	pulumi.RegisterOutputType(DefaultActionOutput{})
	pulumi.RegisterOutputType(DefaultActionPtrOutput{})
	pulumi.RegisterOutputType(DirectoryServiceOptionsOutput{})
	pulumi.RegisterOutputType(DirectoryServiceOptionsPtrOutput{})
	pulumi.RegisterOutputType(HttpProtocolOutput{})
	pulumi.RegisterOutputType(HttpProtocolPtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(KeySourceOutput{})
	pulumi.RegisterOutputType(KeySourcePtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
	pulumi.RegisterOutputType(LargeFileSharesStateOutput{})
	pulumi.RegisterOutputType(LargeFileSharesStatePtrOutput{})
	pulumi.RegisterOutputType(MinimumTlsVersionOutput{})
	pulumi.RegisterOutputType(MinimumTlsVersionPtrOutput{})
	pulumi.RegisterOutputType(PermissionsOutput{})
	pulumi.RegisterOutputType(PermissionsPtrOutput{})
	pulumi.RegisterOutputType(PublicAccessOutput{})
	pulumi.RegisterOutputType(PublicAccessPtrOutput{})
	pulumi.RegisterOutputType(ReasonCodeOutput{})
	pulumi.RegisterOutputType(ReasonCodePtrOutput{})
	pulumi.RegisterOutputType(RuleTypeOutput{})
	pulumi.RegisterOutputType(RuleTypePtrOutput{})
	pulumi.RegisterOutputType(ServicesOutput{})
	pulumi.RegisterOutputType(ServicesPtrOutput{})
	pulumi.RegisterOutputType(SignedResourceOutput{})
	pulumi.RegisterOutputType(SignedResourcePtrOutput{})
	pulumi.RegisterOutputType(SignedResourceTypesOutput{})
	pulumi.RegisterOutputType(SignedResourceTypesPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(StateOutput{})
	pulumi.RegisterOutputType(StatePtrOutput{})
}
