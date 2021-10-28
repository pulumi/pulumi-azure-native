


package v20150501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationType string

const (
	ApplicationTypeWeb   = ApplicationType("web")
	ApplicationTypeOther = ApplicationType("other")
)

func (ApplicationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationType)(nil)).Elem()
}

func (e ApplicationType) ToApplicationTypeOutput() ApplicationTypeOutput {
	return pulumi.ToOutput(e).(ApplicationTypeOutput)
}

func (e ApplicationType) ToApplicationTypeOutputWithContext(ctx context.Context) ApplicationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationTypeOutput)
}

func (e ApplicationType) ToApplicationTypePtrOutput() ApplicationTypePtrOutput {
	return e.ToApplicationTypePtrOutputWithContext(context.Background())
}

func (e ApplicationType) ToApplicationTypePtrOutputWithContext(ctx context.Context) ApplicationTypePtrOutput {
	return ApplicationType(e).ToApplicationTypeOutputWithContext(ctx).ToApplicationTypePtrOutputWithContext(ctx)
}

func (e ApplicationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationTypeOutput struct{ *pulumi.OutputState }

func (ApplicationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationType)(nil)).Elem()
}

func (o ApplicationTypeOutput) ToApplicationTypeOutput() ApplicationTypeOutput {
	return o
}

func (o ApplicationTypeOutput) ToApplicationTypeOutputWithContext(ctx context.Context) ApplicationTypeOutput {
	return o
}

func (o ApplicationTypeOutput) ToApplicationTypePtrOutput() ApplicationTypePtrOutput {
	return o.ToApplicationTypePtrOutputWithContext(context.Background())
}

func (o ApplicationTypeOutput) ToApplicationTypePtrOutputWithContext(ctx context.Context) ApplicationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationType) *ApplicationType {
		return &v
	}).(ApplicationTypePtrOutput)
}

func (o ApplicationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationTypePtrOutput struct{ *pulumi.OutputState }

func (ApplicationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationType)(nil)).Elem()
}

func (o ApplicationTypePtrOutput) ToApplicationTypePtrOutput() ApplicationTypePtrOutput {
	return o
}

func (o ApplicationTypePtrOutput) ToApplicationTypePtrOutputWithContext(ctx context.Context) ApplicationTypePtrOutput {
	return o
}

func (o ApplicationTypePtrOutput) Elem() ApplicationTypeOutput {
	return o.ApplyT(func(v *ApplicationType) ApplicationType {
		if v != nil {
			return *v
		}
		var ret ApplicationType
		return ret
	}).(ApplicationTypeOutput)
}

func (o ApplicationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationTypeInput interface {
	pulumi.Input

	ToApplicationTypeOutput() ApplicationTypeOutput
	ToApplicationTypeOutputWithContext(context.Context) ApplicationTypeOutput
}

var applicationTypePtrType = reflect.TypeOf((**ApplicationType)(nil)).Elem()

type ApplicationTypePtrInput interface {
	pulumi.Input

	ToApplicationTypePtrOutput() ApplicationTypePtrOutput
	ToApplicationTypePtrOutputWithContext(context.Context) ApplicationTypePtrOutput
}

type applicationTypePtr string

func ApplicationTypePtr(v string) ApplicationTypePtrInput {
	return (*applicationTypePtr)(&v)
}

func (*applicationTypePtr) ElementType() reflect.Type {
	return applicationTypePtrType
}

func (in *applicationTypePtr) ToApplicationTypePtrOutput() ApplicationTypePtrOutput {
	return pulumi.ToOutput(in).(ApplicationTypePtrOutput)
}

func (in *applicationTypePtr) ToApplicationTypePtrOutputWithContext(ctx context.Context) ApplicationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationTypePtrOutput)
}

type FavoriteType string

const (
	FavoriteTypeShared = FavoriteType("shared")
	FavoriteTypeUser   = FavoriteType("user")
)

func (FavoriteType) ElementType() reflect.Type {
	return reflect.TypeOf((*FavoriteType)(nil)).Elem()
}

func (e FavoriteType) ToFavoriteTypeOutput() FavoriteTypeOutput {
	return pulumi.ToOutput(e).(FavoriteTypeOutput)
}

func (e FavoriteType) ToFavoriteTypeOutputWithContext(ctx context.Context) FavoriteTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FavoriteTypeOutput)
}

func (e FavoriteType) ToFavoriteTypePtrOutput() FavoriteTypePtrOutput {
	return e.ToFavoriteTypePtrOutputWithContext(context.Background())
}

func (e FavoriteType) ToFavoriteTypePtrOutputWithContext(ctx context.Context) FavoriteTypePtrOutput {
	return FavoriteType(e).ToFavoriteTypeOutputWithContext(ctx).ToFavoriteTypePtrOutputWithContext(ctx)
}

func (e FavoriteType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FavoriteType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FavoriteType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FavoriteType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FavoriteTypeOutput struct{ *pulumi.OutputState }

func (FavoriteTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FavoriteType)(nil)).Elem()
}

func (o FavoriteTypeOutput) ToFavoriteTypeOutput() FavoriteTypeOutput {
	return o
}

func (o FavoriteTypeOutput) ToFavoriteTypeOutputWithContext(ctx context.Context) FavoriteTypeOutput {
	return o
}

func (o FavoriteTypeOutput) ToFavoriteTypePtrOutput() FavoriteTypePtrOutput {
	return o.ToFavoriteTypePtrOutputWithContext(context.Background())
}

func (o FavoriteTypeOutput) ToFavoriteTypePtrOutputWithContext(ctx context.Context) FavoriteTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FavoriteType) *FavoriteType {
		return &v
	}).(FavoriteTypePtrOutput)
}

func (o FavoriteTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FavoriteTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FavoriteType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FavoriteTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FavoriteTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FavoriteType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FavoriteTypePtrOutput struct{ *pulumi.OutputState }

func (FavoriteTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FavoriteType)(nil)).Elem()
}

func (o FavoriteTypePtrOutput) ToFavoriteTypePtrOutput() FavoriteTypePtrOutput {
	return o
}

func (o FavoriteTypePtrOutput) ToFavoriteTypePtrOutputWithContext(ctx context.Context) FavoriteTypePtrOutput {
	return o
}

func (o FavoriteTypePtrOutput) Elem() FavoriteTypeOutput {
	return o.ApplyT(func(v *FavoriteType) FavoriteType {
		if v != nil {
			return *v
		}
		var ret FavoriteType
		return ret
	}).(FavoriteTypeOutput)
}

func (o FavoriteTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FavoriteTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FavoriteType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FavoriteTypeInput interface {
	pulumi.Input

	ToFavoriteTypeOutput() FavoriteTypeOutput
	ToFavoriteTypeOutputWithContext(context.Context) FavoriteTypeOutput
}

var favoriteTypePtrType = reflect.TypeOf((**FavoriteType)(nil)).Elem()

type FavoriteTypePtrInput interface {
	pulumi.Input

	ToFavoriteTypePtrOutput() FavoriteTypePtrOutput
	ToFavoriteTypePtrOutputWithContext(context.Context) FavoriteTypePtrOutput
}

type favoriteTypePtr string

func FavoriteTypePtr(v string) FavoriteTypePtrInput {
	return (*favoriteTypePtr)(&v)
}

func (*favoriteTypePtr) ElementType() reflect.Type {
	return favoriteTypePtrType
}

func (in *favoriteTypePtr) ToFavoriteTypePtrOutput() FavoriteTypePtrOutput {
	return pulumi.ToOutput(in).(FavoriteTypePtrOutput)
}

func (in *favoriteTypePtr) ToFavoriteTypePtrOutputWithContext(ctx context.Context) FavoriteTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FavoriteTypePtrOutput)
}

type FlowType string

const (
	FlowTypeBluefield = FlowType("Bluefield")
)

func (FlowType) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowType)(nil)).Elem()
}

func (e FlowType) ToFlowTypeOutput() FlowTypeOutput {
	return pulumi.ToOutput(e).(FlowTypeOutput)
}

func (e FlowType) ToFlowTypeOutputWithContext(ctx context.Context) FlowTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FlowTypeOutput)
}

func (e FlowType) ToFlowTypePtrOutput() FlowTypePtrOutput {
	return e.ToFlowTypePtrOutputWithContext(context.Background())
}

func (e FlowType) ToFlowTypePtrOutputWithContext(ctx context.Context) FlowTypePtrOutput {
	return FlowType(e).ToFlowTypeOutputWithContext(ctx).ToFlowTypePtrOutputWithContext(ctx)
}

func (e FlowType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FlowType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FlowType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FlowType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FlowTypeOutput struct{ *pulumi.OutputState }

func (FlowTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowType)(nil)).Elem()
}

func (o FlowTypeOutput) ToFlowTypeOutput() FlowTypeOutput {
	return o
}

func (o FlowTypeOutput) ToFlowTypeOutputWithContext(ctx context.Context) FlowTypeOutput {
	return o
}

func (o FlowTypeOutput) ToFlowTypePtrOutput() FlowTypePtrOutput {
	return o.ToFlowTypePtrOutputWithContext(context.Background())
}

func (o FlowTypeOutput) ToFlowTypePtrOutputWithContext(ctx context.Context) FlowTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FlowType) *FlowType {
		return &v
	}).(FlowTypePtrOutput)
}

func (o FlowTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FlowTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FlowType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FlowTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FlowTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FlowType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FlowTypePtrOutput struct{ *pulumi.OutputState }

func (FlowTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowType)(nil)).Elem()
}

func (o FlowTypePtrOutput) ToFlowTypePtrOutput() FlowTypePtrOutput {
	return o
}

func (o FlowTypePtrOutput) ToFlowTypePtrOutputWithContext(ctx context.Context) FlowTypePtrOutput {
	return o
}

func (o FlowTypePtrOutput) Elem() FlowTypeOutput {
	return o.ApplyT(func(v *FlowType) FlowType {
		if v != nil {
			return *v
		}
		var ret FlowType
		return ret
	}).(FlowTypeOutput)
}

func (o FlowTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FlowTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FlowType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FlowTypeInput interface {
	pulumi.Input

	ToFlowTypeOutput() FlowTypeOutput
	ToFlowTypeOutputWithContext(context.Context) FlowTypeOutput
}

var flowTypePtrType = reflect.TypeOf((**FlowType)(nil)).Elem()

type FlowTypePtrInput interface {
	pulumi.Input

	ToFlowTypePtrOutput() FlowTypePtrOutput
	ToFlowTypePtrOutputWithContext(context.Context) FlowTypePtrOutput
}

type flowTypePtr string

func FlowTypePtr(v string) FlowTypePtrInput {
	return (*flowTypePtr)(&v)
}

func (*flowTypePtr) ElementType() reflect.Type {
	return flowTypePtrType
}

func (in *flowTypePtr) ToFlowTypePtrOutput() FlowTypePtrOutput {
	return pulumi.ToOutput(in).(FlowTypePtrOutput)
}

func (in *flowTypePtr) ToFlowTypePtrOutputWithContext(ctx context.Context) FlowTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FlowTypePtrOutput)
}

type IngestionMode string

const (
	IngestionModeApplicationInsights                       = IngestionMode("ApplicationInsights")
	IngestionModeApplicationInsightsWithDiagnosticSettings = IngestionMode("ApplicationInsightsWithDiagnosticSettings")
	IngestionModeLogAnalytics                              = IngestionMode("LogAnalytics")
)

func (IngestionMode) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionMode)(nil)).Elem()
}

func (e IngestionMode) ToIngestionModeOutput() IngestionModeOutput {
	return pulumi.ToOutput(e).(IngestionModeOutput)
}

func (e IngestionMode) ToIngestionModeOutputWithContext(ctx context.Context) IngestionModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IngestionModeOutput)
}

func (e IngestionMode) ToIngestionModePtrOutput() IngestionModePtrOutput {
	return e.ToIngestionModePtrOutputWithContext(context.Background())
}

func (e IngestionMode) ToIngestionModePtrOutputWithContext(ctx context.Context) IngestionModePtrOutput {
	return IngestionMode(e).ToIngestionModeOutputWithContext(ctx).ToIngestionModePtrOutputWithContext(ctx)
}

func (e IngestionMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IngestionMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IngestionMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IngestionMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IngestionModeOutput struct{ *pulumi.OutputState }

func (IngestionModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionMode)(nil)).Elem()
}

func (o IngestionModeOutput) ToIngestionModeOutput() IngestionModeOutput {
	return o
}

func (o IngestionModeOutput) ToIngestionModeOutputWithContext(ctx context.Context) IngestionModeOutput {
	return o
}

func (o IngestionModeOutput) ToIngestionModePtrOutput() IngestionModePtrOutput {
	return o.ToIngestionModePtrOutputWithContext(context.Background())
}

func (o IngestionModeOutput) ToIngestionModePtrOutputWithContext(ctx context.Context) IngestionModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IngestionMode) *IngestionMode {
		return &v
	}).(IngestionModePtrOutput)
}

func (o IngestionModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IngestionModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IngestionMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IngestionModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IngestionModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IngestionMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IngestionModePtrOutput struct{ *pulumi.OutputState }

func (IngestionModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngestionMode)(nil)).Elem()
}

func (o IngestionModePtrOutput) ToIngestionModePtrOutput() IngestionModePtrOutput {
	return o
}

func (o IngestionModePtrOutput) ToIngestionModePtrOutputWithContext(ctx context.Context) IngestionModePtrOutput {
	return o
}

func (o IngestionModePtrOutput) Elem() IngestionModeOutput {
	return o.ApplyT(func(v *IngestionMode) IngestionMode {
		if v != nil {
			return *v
		}
		var ret IngestionMode
		return ret
	}).(IngestionModeOutput)
}

func (o IngestionModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IngestionModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IngestionMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IngestionModeInput interface {
	pulumi.Input

	ToIngestionModeOutput() IngestionModeOutput
	ToIngestionModeOutputWithContext(context.Context) IngestionModeOutput
}

var ingestionModePtrType = reflect.TypeOf((**IngestionMode)(nil)).Elem()

type IngestionModePtrInput interface {
	pulumi.Input

	ToIngestionModePtrOutput() IngestionModePtrOutput
	ToIngestionModePtrOutputWithContext(context.Context) IngestionModePtrOutput
}

type ingestionModePtr string

func IngestionModePtr(v string) IngestionModePtrInput {
	return (*ingestionModePtr)(&v)
}

func (*ingestionModePtr) ElementType() reflect.Type {
	return ingestionModePtrType
}

func (in *ingestionModePtr) ToIngestionModePtrOutput() IngestionModePtrOutput {
	return pulumi.ToOutput(in).(IngestionModePtrOutput)
}

func (in *ingestionModePtr) ToIngestionModePtrOutputWithContext(ctx context.Context) IngestionModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IngestionModePtrOutput)
}

type ItemScope string

const (
	ItemScopeShared = ItemScope("shared")
	ItemScopeUser   = ItemScope("user")
)

func (ItemScope) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemScope)(nil)).Elem()
}

func (e ItemScope) ToItemScopeOutput() ItemScopeOutput {
	return pulumi.ToOutput(e).(ItemScopeOutput)
}

func (e ItemScope) ToItemScopeOutputWithContext(ctx context.Context) ItemScopeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ItemScopeOutput)
}

func (e ItemScope) ToItemScopePtrOutput() ItemScopePtrOutput {
	return e.ToItemScopePtrOutputWithContext(context.Background())
}

func (e ItemScope) ToItemScopePtrOutputWithContext(ctx context.Context) ItemScopePtrOutput {
	return ItemScope(e).ToItemScopeOutputWithContext(ctx).ToItemScopePtrOutputWithContext(ctx)
}

func (e ItemScope) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ItemScope) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ItemScope) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ItemScope) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ItemScopeOutput struct{ *pulumi.OutputState }

func (ItemScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemScope)(nil)).Elem()
}

func (o ItemScopeOutput) ToItemScopeOutput() ItemScopeOutput {
	return o
}

func (o ItemScopeOutput) ToItemScopeOutputWithContext(ctx context.Context) ItemScopeOutput {
	return o
}

func (o ItemScopeOutput) ToItemScopePtrOutput() ItemScopePtrOutput {
	return o.ToItemScopePtrOutputWithContext(context.Background())
}

func (o ItemScopeOutput) ToItemScopePtrOutputWithContext(ctx context.Context) ItemScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ItemScope) *ItemScope {
		return &v
	}).(ItemScopePtrOutput)
}

func (o ItemScopeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ItemScopeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ItemScope) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ItemScopeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ItemScopeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ItemScope) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ItemScopePtrOutput struct{ *pulumi.OutputState }

func (ItemScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ItemScope)(nil)).Elem()
}

func (o ItemScopePtrOutput) ToItemScopePtrOutput() ItemScopePtrOutput {
	return o
}

func (o ItemScopePtrOutput) ToItemScopePtrOutputWithContext(ctx context.Context) ItemScopePtrOutput {
	return o
}

func (o ItemScopePtrOutput) Elem() ItemScopeOutput {
	return o.ApplyT(func(v *ItemScope) ItemScope {
		if v != nil {
			return *v
		}
		var ret ItemScope
		return ret
	}).(ItemScopeOutput)
}

func (o ItemScopePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ItemScopePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ItemScope) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ItemScopeInput interface {
	pulumi.Input

	ToItemScopeOutput() ItemScopeOutput
	ToItemScopeOutputWithContext(context.Context) ItemScopeOutput
}

var itemScopePtrType = reflect.TypeOf((**ItemScope)(nil)).Elem()

type ItemScopePtrInput interface {
	pulumi.Input

	ToItemScopePtrOutput() ItemScopePtrOutput
	ToItemScopePtrOutputWithContext(context.Context) ItemScopePtrOutput
}

type itemScopePtr string

func ItemScopePtr(v string) ItemScopePtrInput {
	return (*itemScopePtr)(&v)
}

func (*itemScopePtr) ElementType() reflect.Type {
	return itemScopePtrType
}

func (in *itemScopePtr) ToItemScopePtrOutput() ItemScopePtrOutput {
	return pulumi.ToOutput(in).(ItemScopePtrOutput)
}

func (in *itemScopePtr) ToItemScopePtrOutputWithContext(ctx context.Context) ItemScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ItemScopePtrOutput)
}

type ItemType string

const (
	ItemTypeQuery    = ItemType("query")
	ItemTypeFunction = ItemType("function")
	ItemTypeFolder   = ItemType("folder")
	ItemTypeRecent   = ItemType("recent")
)

func (ItemType) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemType)(nil)).Elem()
}

func (e ItemType) ToItemTypeOutput() ItemTypeOutput {
	return pulumi.ToOutput(e).(ItemTypeOutput)
}

func (e ItemType) ToItemTypeOutputWithContext(ctx context.Context) ItemTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ItemTypeOutput)
}

func (e ItemType) ToItemTypePtrOutput() ItemTypePtrOutput {
	return e.ToItemTypePtrOutputWithContext(context.Background())
}

func (e ItemType) ToItemTypePtrOutputWithContext(ctx context.Context) ItemTypePtrOutput {
	return ItemType(e).ToItemTypeOutputWithContext(ctx).ToItemTypePtrOutputWithContext(ctx)
}

func (e ItemType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ItemType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ItemType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ItemType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ItemTypeOutput struct{ *pulumi.OutputState }

func (ItemTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemType)(nil)).Elem()
}

func (o ItemTypeOutput) ToItemTypeOutput() ItemTypeOutput {
	return o
}

func (o ItemTypeOutput) ToItemTypeOutputWithContext(ctx context.Context) ItemTypeOutput {
	return o
}

func (o ItemTypeOutput) ToItemTypePtrOutput() ItemTypePtrOutput {
	return o.ToItemTypePtrOutputWithContext(context.Background())
}

func (o ItemTypeOutput) ToItemTypePtrOutputWithContext(ctx context.Context) ItemTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ItemType) *ItemType {
		return &v
	}).(ItemTypePtrOutput)
}

func (o ItemTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ItemTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ItemType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ItemTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ItemTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ItemType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ItemTypePtrOutput struct{ *pulumi.OutputState }

func (ItemTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ItemType)(nil)).Elem()
}

func (o ItemTypePtrOutput) ToItemTypePtrOutput() ItemTypePtrOutput {
	return o
}

func (o ItemTypePtrOutput) ToItemTypePtrOutputWithContext(ctx context.Context) ItemTypePtrOutput {
	return o
}

func (o ItemTypePtrOutput) Elem() ItemTypeOutput {
	return o.ApplyT(func(v *ItemType) ItemType {
		if v != nil {
			return *v
		}
		var ret ItemType
		return ret
	}).(ItemTypeOutput)
}

func (o ItemTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ItemTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ItemType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ItemTypeInput interface {
	pulumi.Input

	ToItemTypeOutput() ItemTypeOutput
	ToItemTypeOutputWithContext(context.Context) ItemTypeOutput
}

var itemTypePtrType = reflect.TypeOf((**ItemType)(nil)).Elem()

type ItemTypePtrInput interface {
	pulumi.Input

	ToItemTypePtrOutput() ItemTypePtrOutput
	ToItemTypePtrOutputWithContext(context.Context) ItemTypePtrOutput
}

type itemTypePtr string

func ItemTypePtr(v string) ItemTypePtrInput {
	return (*itemTypePtr)(&v)
}

func (*itemTypePtr) ElementType() reflect.Type {
	return itemTypePtrType
}

func (in *itemTypePtr) ToItemTypePtrOutput() ItemTypePtrOutput {
	return pulumi.ToOutput(in).(ItemTypePtrOutput)
}

func (in *itemTypePtr) ToItemTypePtrOutputWithContext(ctx context.Context) ItemTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ItemTypePtrOutput)
}

type RequestSource string

const (
	RequestSourceRest = RequestSource("rest")
)

func (RequestSource) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestSource)(nil)).Elem()
}

func (e RequestSource) ToRequestSourceOutput() RequestSourceOutput {
	return pulumi.ToOutput(e).(RequestSourceOutput)
}

func (e RequestSource) ToRequestSourceOutputWithContext(ctx context.Context) RequestSourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RequestSourceOutput)
}

func (e RequestSource) ToRequestSourcePtrOutput() RequestSourcePtrOutput {
	return e.ToRequestSourcePtrOutputWithContext(context.Background())
}

func (e RequestSource) ToRequestSourcePtrOutputWithContext(ctx context.Context) RequestSourcePtrOutput {
	return RequestSource(e).ToRequestSourceOutputWithContext(ctx).ToRequestSourcePtrOutputWithContext(ctx)
}

func (e RequestSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RequestSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RequestSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RequestSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RequestSourceOutput struct{ *pulumi.OutputState }

func (RequestSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestSource)(nil)).Elem()
}

func (o RequestSourceOutput) ToRequestSourceOutput() RequestSourceOutput {
	return o
}

func (o RequestSourceOutput) ToRequestSourceOutputWithContext(ctx context.Context) RequestSourceOutput {
	return o
}

func (o RequestSourceOutput) ToRequestSourcePtrOutput() RequestSourcePtrOutput {
	return o.ToRequestSourcePtrOutputWithContext(context.Background())
}

func (o RequestSourceOutput) ToRequestSourcePtrOutputWithContext(ctx context.Context) RequestSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RequestSource) *RequestSource {
		return &v
	}).(RequestSourcePtrOutput)
}

func (o RequestSourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RequestSourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RequestSource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RequestSourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RequestSourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RequestSource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RequestSourcePtrOutput struct{ *pulumi.OutputState }

func (RequestSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestSource)(nil)).Elem()
}

func (o RequestSourcePtrOutput) ToRequestSourcePtrOutput() RequestSourcePtrOutput {
	return o
}

func (o RequestSourcePtrOutput) ToRequestSourcePtrOutputWithContext(ctx context.Context) RequestSourcePtrOutput {
	return o
}

func (o RequestSourcePtrOutput) Elem() RequestSourceOutput {
	return o.ApplyT(func(v *RequestSource) RequestSource {
		if v != nil {
			return *v
		}
		var ret RequestSource
		return ret
	}).(RequestSourceOutput)
}

func (o RequestSourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RequestSourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RequestSource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RequestSourceInput interface {
	pulumi.Input

	ToRequestSourceOutput() RequestSourceOutput
	ToRequestSourceOutputWithContext(context.Context) RequestSourceOutput
}

var requestSourcePtrType = reflect.TypeOf((**RequestSource)(nil)).Elem()

type RequestSourcePtrInput interface {
	pulumi.Input

	ToRequestSourcePtrOutput() RequestSourcePtrOutput
	ToRequestSourcePtrOutputWithContext(context.Context) RequestSourcePtrOutput
}

type requestSourcePtr string

func RequestSourcePtr(v string) RequestSourcePtrInput {
	return (*requestSourcePtr)(&v)
}

func (*requestSourcePtr) ElementType() reflect.Type {
	return requestSourcePtrType
}

func (in *requestSourcePtr) ToRequestSourcePtrOutput() RequestSourcePtrOutput {
	return pulumi.ToOutput(in).(RequestSourcePtrOutput)
}

func (in *requestSourcePtr) ToRequestSourcePtrOutputWithContext(ctx context.Context) RequestSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RequestSourcePtrOutput)
}

type SharedTypeKind string

const (
	SharedTypeKindShared = SharedTypeKind("shared")
	SharedTypeKindUser   = SharedTypeKind("user")
)

func (SharedTypeKind) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedTypeKind)(nil)).Elem()
}

func (e SharedTypeKind) ToSharedTypeKindOutput() SharedTypeKindOutput {
	return pulumi.ToOutput(e).(SharedTypeKindOutput)
}

func (e SharedTypeKind) ToSharedTypeKindOutputWithContext(ctx context.Context) SharedTypeKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SharedTypeKindOutput)
}

func (e SharedTypeKind) ToSharedTypeKindPtrOutput() SharedTypeKindPtrOutput {
	return e.ToSharedTypeKindPtrOutputWithContext(context.Background())
}

func (e SharedTypeKind) ToSharedTypeKindPtrOutputWithContext(ctx context.Context) SharedTypeKindPtrOutput {
	return SharedTypeKind(e).ToSharedTypeKindOutputWithContext(ctx).ToSharedTypeKindPtrOutputWithContext(ctx)
}

func (e SharedTypeKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedTypeKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedTypeKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SharedTypeKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SharedTypeKindOutput struct{ *pulumi.OutputState }

func (SharedTypeKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedTypeKind)(nil)).Elem()
}

func (o SharedTypeKindOutput) ToSharedTypeKindOutput() SharedTypeKindOutput {
	return o
}

func (o SharedTypeKindOutput) ToSharedTypeKindOutputWithContext(ctx context.Context) SharedTypeKindOutput {
	return o
}

func (o SharedTypeKindOutput) ToSharedTypeKindPtrOutput() SharedTypeKindPtrOutput {
	return o.ToSharedTypeKindPtrOutputWithContext(context.Background())
}

func (o SharedTypeKindOutput) ToSharedTypeKindPtrOutputWithContext(ctx context.Context) SharedTypeKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharedTypeKind) *SharedTypeKind {
		return &v
	}).(SharedTypeKindPtrOutput)
}

func (o SharedTypeKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SharedTypeKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SharedTypeKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SharedTypeKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SharedTypeKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SharedTypeKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SharedTypeKindPtrOutput struct{ *pulumi.OutputState }

func (SharedTypeKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedTypeKind)(nil)).Elem()
}

func (o SharedTypeKindPtrOutput) ToSharedTypeKindPtrOutput() SharedTypeKindPtrOutput {
	return o
}

func (o SharedTypeKindPtrOutput) ToSharedTypeKindPtrOutputWithContext(ctx context.Context) SharedTypeKindPtrOutput {
	return o
}

func (o SharedTypeKindPtrOutput) Elem() SharedTypeKindOutput {
	return o.ApplyT(func(v *SharedTypeKind) SharedTypeKind {
		if v != nil {
			return *v
		}
		var ret SharedTypeKind
		return ret
	}).(SharedTypeKindOutput)
}

func (o SharedTypeKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SharedTypeKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SharedTypeKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SharedTypeKindInput interface {
	pulumi.Input

	ToSharedTypeKindOutput() SharedTypeKindOutput
	ToSharedTypeKindOutputWithContext(context.Context) SharedTypeKindOutput
}

var sharedTypeKindPtrType = reflect.TypeOf((**SharedTypeKind)(nil)).Elem()

type SharedTypeKindPtrInput interface {
	pulumi.Input

	ToSharedTypeKindPtrOutput() SharedTypeKindPtrOutput
	ToSharedTypeKindPtrOutputWithContext(context.Context) SharedTypeKindPtrOutput
}

type sharedTypeKindPtr string

func SharedTypeKindPtr(v string) SharedTypeKindPtrInput {
	return (*sharedTypeKindPtr)(&v)
}

func (*sharedTypeKindPtr) ElementType() reflect.Type {
	return sharedTypeKindPtrType
}

func (in *sharedTypeKindPtr) ToSharedTypeKindPtrOutput() SharedTypeKindPtrOutput {
	return pulumi.ToOutput(in).(SharedTypeKindPtrOutput)
}

func (in *sharedTypeKindPtr) ToSharedTypeKindPtrOutputWithContext(ctx context.Context) SharedTypeKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SharedTypeKindPtrOutput)
}

type WebTestKind string

const (
	WebTestKindPing      = WebTestKind("ping")
	WebTestKindMultistep = WebTestKind("multistep")
)

func (WebTestKind) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestKind)(nil)).Elem()
}

func (e WebTestKind) ToWebTestKindOutput() WebTestKindOutput {
	return pulumi.ToOutput(e).(WebTestKindOutput)
}

func (e WebTestKind) ToWebTestKindOutputWithContext(ctx context.Context) WebTestKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebTestKindOutput)
}

func (e WebTestKind) ToWebTestKindPtrOutput() WebTestKindPtrOutput {
	return e.ToWebTestKindPtrOutputWithContext(context.Background())
}

func (e WebTestKind) ToWebTestKindPtrOutputWithContext(ctx context.Context) WebTestKindPtrOutput {
	return WebTestKind(e).ToWebTestKindOutputWithContext(ctx).ToWebTestKindPtrOutputWithContext(ctx)
}

func (e WebTestKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebTestKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebTestKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebTestKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebTestKindOutput struct{ *pulumi.OutputState }

func (WebTestKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestKind)(nil)).Elem()
}

func (o WebTestKindOutput) ToWebTestKindOutput() WebTestKindOutput {
	return o
}

func (o WebTestKindOutput) ToWebTestKindOutputWithContext(ctx context.Context) WebTestKindOutput {
	return o
}

func (o WebTestKindOutput) ToWebTestKindPtrOutput() WebTestKindPtrOutput {
	return o.ToWebTestKindPtrOutputWithContext(context.Background())
}

func (o WebTestKindOutput) ToWebTestKindPtrOutputWithContext(ctx context.Context) WebTestKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebTestKind) *WebTestKind {
		return &v
	}).(WebTestKindPtrOutput)
}

func (o WebTestKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebTestKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebTestKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebTestKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebTestKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebTestKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebTestKindPtrOutput struct{ *pulumi.OutputState }

func (WebTestKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestKind)(nil)).Elem()
}

func (o WebTestKindPtrOutput) ToWebTestKindPtrOutput() WebTestKindPtrOutput {
	return o
}

func (o WebTestKindPtrOutput) ToWebTestKindPtrOutputWithContext(ctx context.Context) WebTestKindPtrOutput {
	return o
}

func (o WebTestKindPtrOutput) Elem() WebTestKindOutput {
	return o.ApplyT(func(v *WebTestKind) WebTestKind {
		if v != nil {
			return *v
		}
		var ret WebTestKind
		return ret
	}).(WebTestKindOutput)
}

func (o WebTestKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebTestKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebTestKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebTestKindInput interface {
	pulumi.Input

	ToWebTestKindOutput() WebTestKindOutput
	ToWebTestKindOutputWithContext(context.Context) WebTestKindOutput
}

var webTestKindPtrType = reflect.TypeOf((**WebTestKind)(nil)).Elem()

type WebTestKindPtrInput interface {
	pulumi.Input

	ToWebTestKindPtrOutput() WebTestKindPtrOutput
	ToWebTestKindPtrOutputWithContext(context.Context) WebTestKindPtrOutput
}

type webTestKindPtr string

func WebTestKindPtr(v string) WebTestKindPtrInput {
	return (*webTestKindPtr)(&v)
}

func (*webTestKindPtr) ElementType() reflect.Type {
	return webTestKindPtrType
}

func (in *webTestKindPtr) ToWebTestKindPtrOutput() WebTestKindPtrOutput {
	return pulumi.ToOutput(in).(WebTestKindPtrOutput)
}

func (in *webTestKindPtr) ToWebTestKindPtrOutputWithContext(ctx context.Context) WebTestKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebTestKindPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationTypeInput)(nil)).Elem(), ApplicationType("web"))
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationTypePtrInput)(nil)).Elem(), ApplicationType("web"))
	pulumi.RegisterInputType(reflect.TypeOf((*FavoriteTypeInput)(nil)).Elem(), FavoriteType("shared"))
	pulumi.RegisterInputType(reflect.TypeOf((*FavoriteTypePtrInput)(nil)).Elem(), FavoriteType("shared"))
	pulumi.RegisterInputType(reflect.TypeOf((*FlowTypeInput)(nil)).Elem(), FlowType("Bluefield"))
	pulumi.RegisterInputType(reflect.TypeOf((*FlowTypePtrInput)(nil)).Elem(), FlowType("Bluefield"))
	pulumi.RegisterInputType(reflect.TypeOf((*IngestionModeInput)(nil)).Elem(), IngestionMode("ApplicationInsights"))
	pulumi.RegisterInputType(reflect.TypeOf((*IngestionModePtrInput)(nil)).Elem(), IngestionMode("ApplicationInsights"))
	pulumi.RegisterInputType(reflect.TypeOf((*ItemScopeInput)(nil)).Elem(), ItemScope("shared"))
	pulumi.RegisterInputType(reflect.TypeOf((*ItemScopePtrInput)(nil)).Elem(), ItemScope("shared"))
	pulumi.RegisterInputType(reflect.TypeOf((*ItemTypeInput)(nil)).Elem(), ItemType("query"))
	pulumi.RegisterInputType(reflect.TypeOf((*ItemTypePtrInput)(nil)).Elem(), ItemType("query"))
	pulumi.RegisterInputType(reflect.TypeOf((*RequestSourceInput)(nil)).Elem(), RequestSource("rest"))
	pulumi.RegisterInputType(reflect.TypeOf((*RequestSourcePtrInput)(nil)).Elem(), RequestSource("rest"))
	pulumi.RegisterInputType(reflect.TypeOf((*SharedTypeKindInput)(nil)).Elem(), SharedTypeKind("shared"))
	pulumi.RegisterInputType(reflect.TypeOf((*SharedTypeKindPtrInput)(nil)).Elem(), SharedTypeKind("shared"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebTestKindInput)(nil)).Elem(), WebTestKind("ping"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebTestKindPtrInput)(nil)).Elem(), WebTestKind("ping"))
	pulumi.RegisterOutputType(ApplicationTypeOutput{})
	pulumi.RegisterOutputType(ApplicationTypePtrOutput{})
	pulumi.RegisterOutputType(FavoriteTypeOutput{})
	pulumi.RegisterOutputType(FavoriteTypePtrOutput{})
	pulumi.RegisterOutputType(FlowTypeOutput{})
	pulumi.RegisterOutputType(FlowTypePtrOutput{})
	pulumi.RegisterOutputType(IngestionModeOutput{})
	pulumi.RegisterOutputType(IngestionModePtrOutput{})
	pulumi.RegisterOutputType(ItemScopeOutput{})
	pulumi.RegisterOutputType(ItemScopePtrOutput{})
	pulumi.RegisterOutputType(ItemTypeOutput{})
	pulumi.RegisterOutputType(ItemTypePtrOutput{})
	pulumi.RegisterOutputType(RequestSourceOutput{})
	pulumi.RegisterOutputType(RequestSourcePtrOutput{})
	pulumi.RegisterOutputType(SharedTypeKindOutput{})
	pulumi.RegisterOutputType(SharedTypeKindPtrOutput{})
	pulumi.RegisterOutputType(WebTestKindOutput{})
	pulumi.RegisterOutputType(WebTestKindPtrOutput{})
}
