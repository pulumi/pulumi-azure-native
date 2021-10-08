


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventHandlerSettings struct {
	Items map[string][]EventHandlerTemplate `pulumi:"items"`
}





type EventHandlerSettingsInput interface {
	pulumi.Input

	ToEventHandlerSettingsOutput() EventHandlerSettingsOutput
	ToEventHandlerSettingsOutputWithContext(context.Context) EventHandlerSettingsOutput
}

type EventHandlerSettingsArgs struct {
	Items EventHandlerTemplateArrayMapInput `pulumi:"items"`
}

func (EventHandlerSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerSettings)(nil)).Elem()
}

func (i EventHandlerSettingsArgs) ToEventHandlerSettingsOutput() EventHandlerSettingsOutput {
	return i.ToEventHandlerSettingsOutputWithContext(context.Background())
}

func (i EventHandlerSettingsArgs) ToEventHandlerSettingsOutputWithContext(ctx context.Context) EventHandlerSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerSettingsOutput)
}

func (i EventHandlerSettingsArgs) ToEventHandlerSettingsPtrOutput() EventHandlerSettingsPtrOutput {
	return i.ToEventHandlerSettingsPtrOutputWithContext(context.Background())
}

func (i EventHandlerSettingsArgs) ToEventHandlerSettingsPtrOutputWithContext(ctx context.Context) EventHandlerSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerSettingsOutput).ToEventHandlerSettingsPtrOutputWithContext(ctx)
}









type EventHandlerSettingsPtrInput interface {
	pulumi.Input

	ToEventHandlerSettingsPtrOutput() EventHandlerSettingsPtrOutput
	ToEventHandlerSettingsPtrOutputWithContext(context.Context) EventHandlerSettingsPtrOutput
}

type eventHandlerSettingsPtrType EventHandlerSettingsArgs

func EventHandlerSettingsPtr(v *EventHandlerSettingsArgs) EventHandlerSettingsPtrInput {
	return (*eventHandlerSettingsPtrType)(v)
}

func (*eventHandlerSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHandlerSettings)(nil)).Elem()
}

func (i *eventHandlerSettingsPtrType) ToEventHandlerSettingsPtrOutput() EventHandlerSettingsPtrOutput {
	return i.ToEventHandlerSettingsPtrOutputWithContext(context.Background())
}

func (i *eventHandlerSettingsPtrType) ToEventHandlerSettingsPtrOutputWithContext(ctx context.Context) EventHandlerSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerSettingsPtrOutput)
}

type EventHandlerSettingsOutput struct{ *pulumi.OutputState }

func (EventHandlerSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerSettings)(nil)).Elem()
}

func (o EventHandlerSettingsOutput) ToEventHandlerSettingsOutput() EventHandlerSettingsOutput {
	return o
}

func (o EventHandlerSettingsOutput) ToEventHandlerSettingsOutputWithContext(ctx context.Context) EventHandlerSettingsOutput {
	return o
}

func (o EventHandlerSettingsOutput) ToEventHandlerSettingsPtrOutput() EventHandlerSettingsPtrOutput {
	return o.ToEventHandlerSettingsPtrOutputWithContext(context.Background())
}

func (o EventHandlerSettingsOutput) ToEventHandlerSettingsPtrOutputWithContext(ctx context.Context) EventHandlerSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventHandlerSettings) *EventHandlerSettings {
		return &v
	}).(EventHandlerSettingsPtrOutput)
}

func (o EventHandlerSettingsOutput) Items() EventHandlerTemplateArrayMapOutput {
	return o.ApplyT(func(v EventHandlerSettings) map[string][]EventHandlerTemplate { return v.Items }).(EventHandlerTemplateArrayMapOutput)
}

type EventHandlerSettingsPtrOutput struct{ *pulumi.OutputState }

func (EventHandlerSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHandlerSettings)(nil)).Elem()
}

func (o EventHandlerSettingsPtrOutput) ToEventHandlerSettingsPtrOutput() EventHandlerSettingsPtrOutput {
	return o
}

func (o EventHandlerSettingsPtrOutput) ToEventHandlerSettingsPtrOutputWithContext(ctx context.Context) EventHandlerSettingsPtrOutput {
	return o
}

func (o EventHandlerSettingsPtrOutput) Elem() EventHandlerSettingsOutput {
	return o.ApplyT(func(v *EventHandlerSettings) EventHandlerSettings {
		if v != nil {
			return *v
		}
		var ret EventHandlerSettings
		return ret
	}).(EventHandlerSettingsOutput)
}

func (o EventHandlerSettingsPtrOutput) Items() EventHandlerTemplateArrayMapOutput {
	return o.ApplyT(func(v *EventHandlerSettings) map[string][]EventHandlerTemplate {
		if v == nil {
			return nil
		}
		return v.Items
	}).(EventHandlerTemplateArrayMapOutput)
}

type EventHandlerSettingsResponse struct {
	Items map[string][]EventHandlerTemplateResponse `pulumi:"items"`
}





type EventHandlerSettingsResponseInput interface {
	pulumi.Input

	ToEventHandlerSettingsResponseOutput() EventHandlerSettingsResponseOutput
	ToEventHandlerSettingsResponseOutputWithContext(context.Context) EventHandlerSettingsResponseOutput
}

type EventHandlerSettingsResponseArgs struct {
	Items EventHandlerTemplateResponseArrayMapInput `pulumi:"items"`
}

func (EventHandlerSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerSettingsResponse)(nil)).Elem()
}

func (i EventHandlerSettingsResponseArgs) ToEventHandlerSettingsResponseOutput() EventHandlerSettingsResponseOutput {
	return i.ToEventHandlerSettingsResponseOutputWithContext(context.Background())
}

func (i EventHandlerSettingsResponseArgs) ToEventHandlerSettingsResponseOutputWithContext(ctx context.Context) EventHandlerSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerSettingsResponseOutput)
}

func (i EventHandlerSettingsResponseArgs) ToEventHandlerSettingsResponsePtrOutput() EventHandlerSettingsResponsePtrOutput {
	return i.ToEventHandlerSettingsResponsePtrOutputWithContext(context.Background())
}

func (i EventHandlerSettingsResponseArgs) ToEventHandlerSettingsResponsePtrOutputWithContext(ctx context.Context) EventHandlerSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerSettingsResponseOutput).ToEventHandlerSettingsResponsePtrOutputWithContext(ctx)
}









type EventHandlerSettingsResponsePtrInput interface {
	pulumi.Input

	ToEventHandlerSettingsResponsePtrOutput() EventHandlerSettingsResponsePtrOutput
	ToEventHandlerSettingsResponsePtrOutputWithContext(context.Context) EventHandlerSettingsResponsePtrOutput
}

type eventHandlerSettingsResponsePtrType EventHandlerSettingsResponseArgs

func EventHandlerSettingsResponsePtr(v *EventHandlerSettingsResponseArgs) EventHandlerSettingsResponsePtrInput {
	return (*eventHandlerSettingsResponsePtrType)(v)
}

func (*eventHandlerSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHandlerSettingsResponse)(nil)).Elem()
}

func (i *eventHandlerSettingsResponsePtrType) ToEventHandlerSettingsResponsePtrOutput() EventHandlerSettingsResponsePtrOutput {
	return i.ToEventHandlerSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *eventHandlerSettingsResponsePtrType) ToEventHandlerSettingsResponsePtrOutputWithContext(ctx context.Context) EventHandlerSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerSettingsResponsePtrOutput)
}

type EventHandlerSettingsResponseOutput struct{ *pulumi.OutputState }

func (EventHandlerSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerSettingsResponse)(nil)).Elem()
}

func (o EventHandlerSettingsResponseOutput) ToEventHandlerSettingsResponseOutput() EventHandlerSettingsResponseOutput {
	return o
}

func (o EventHandlerSettingsResponseOutput) ToEventHandlerSettingsResponseOutputWithContext(ctx context.Context) EventHandlerSettingsResponseOutput {
	return o
}

func (o EventHandlerSettingsResponseOutput) ToEventHandlerSettingsResponsePtrOutput() EventHandlerSettingsResponsePtrOutput {
	return o.ToEventHandlerSettingsResponsePtrOutputWithContext(context.Background())
}

func (o EventHandlerSettingsResponseOutput) ToEventHandlerSettingsResponsePtrOutputWithContext(ctx context.Context) EventHandlerSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventHandlerSettingsResponse) *EventHandlerSettingsResponse {
		return &v
	}).(EventHandlerSettingsResponsePtrOutput)
}

func (o EventHandlerSettingsResponseOutput) Items() EventHandlerTemplateResponseArrayMapOutput {
	return o.ApplyT(func(v EventHandlerSettingsResponse) map[string][]EventHandlerTemplateResponse { return v.Items }).(EventHandlerTemplateResponseArrayMapOutput)
}

type EventHandlerSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EventHandlerSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHandlerSettingsResponse)(nil)).Elem()
}

func (o EventHandlerSettingsResponsePtrOutput) ToEventHandlerSettingsResponsePtrOutput() EventHandlerSettingsResponsePtrOutput {
	return o
}

func (o EventHandlerSettingsResponsePtrOutput) ToEventHandlerSettingsResponsePtrOutputWithContext(ctx context.Context) EventHandlerSettingsResponsePtrOutput {
	return o
}

func (o EventHandlerSettingsResponsePtrOutput) Elem() EventHandlerSettingsResponseOutput {
	return o.ApplyT(func(v *EventHandlerSettingsResponse) EventHandlerSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EventHandlerSettingsResponse
		return ret
	}).(EventHandlerSettingsResponseOutput)
}

func (o EventHandlerSettingsResponsePtrOutput) Items() EventHandlerTemplateResponseArrayMapOutput {
	return o.ApplyT(func(v *EventHandlerSettingsResponse) map[string][]EventHandlerTemplateResponse {
		if v == nil {
			return nil
		}
		return v.Items
	}).(EventHandlerTemplateResponseArrayMapOutput)
}

type EventHandlerTemplate struct {
	Auth               *UpstreamAuthSettings `pulumi:"auth"`
	SystemEventPattern *string               `pulumi:"systemEventPattern"`
	UrlTemplate        string                `pulumi:"urlTemplate"`
	UserEventPattern   *string               `pulumi:"userEventPattern"`
}





type EventHandlerTemplateInput interface {
	pulumi.Input

	ToEventHandlerTemplateOutput() EventHandlerTemplateOutput
	ToEventHandlerTemplateOutputWithContext(context.Context) EventHandlerTemplateOutput
}

type EventHandlerTemplateArgs struct {
	Auth               UpstreamAuthSettingsPtrInput `pulumi:"auth"`
	SystemEventPattern pulumi.StringPtrInput        `pulumi:"systemEventPattern"`
	UrlTemplate        pulumi.StringInput           `pulumi:"urlTemplate"`
	UserEventPattern   pulumi.StringPtrInput        `pulumi:"userEventPattern"`
}

func (EventHandlerTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerTemplate)(nil)).Elem()
}

func (i EventHandlerTemplateArgs) ToEventHandlerTemplateOutput() EventHandlerTemplateOutput {
	return i.ToEventHandlerTemplateOutputWithContext(context.Background())
}

func (i EventHandlerTemplateArgs) ToEventHandlerTemplateOutputWithContext(ctx context.Context) EventHandlerTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerTemplateOutput)
}





type EventHandlerTemplateArrayInput interface {
	pulumi.Input

	ToEventHandlerTemplateArrayOutput() EventHandlerTemplateArrayOutput
	ToEventHandlerTemplateArrayOutputWithContext(context.Context) EventHandlerTemplateArrayOutput
}

type EventHandlerTemplateArray []EventHandlerTemplateInput

func (EventHandlerTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHandlerTemplate)(nil)).Elem()
}

func (i EventHandlerTemplateArray) ToEventHandlerTemplateArrayOutput() EventHandlerTemplateArrayOutput {
	return i.ToEventHandlerTemplateArrayOutputWithContext(context.Background())
}

func (i EventHandlerTemplateArray) ToEventHandlerTemplateArrayOutputWithContext(ctx context.Context) EventHandlerTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerTemplateArrayOutput)
}

type EventHandlerTemplateOutput struct{ *pulumi.OutputState }

func (EventHandlerTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerTemplate)(nil)).Elem()
}

func (o EventHandlerTemplateOutput) ToEventHandlerTemplateOutput() EventHandlerTemplateOutput {
	return o
}

func (o EventHandlerTemplateOutput) ToEventHandlerTemplateOutputWithContext(ctx context.Context) EventHandlerTemplateOutput {
	return o
}

func (o EventHandlerTemplateOutput) Auth() UpstreamAuthSettingsPtrOutput {
	return o.ApplyT(func(v EventHandlerTemplate) *UpstreamAuthSettings { return v.Auth }).(UpstreamAuthSettingsPtrOutput)
}

func (o EventHandlerTemplateOutput) SystemEventPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHandlerTemplate) *string { return v.SystemEventPattern }).(pulumi.StringPtrOutput)
}

func (o EventHandlerTemplateOutput) UrlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v EventHandlerTemplate) string { return v.UrlTemplate }).(pulumi.StringOutput)
}

func (o EventHandlerTemplateOutput) UserEventPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHandlerTemplate) *string { return v.UserEventPattern }).(pulumi.StringPtrOutput)
}

type EventHandlerTemplateArrayOutput struct{ *pulumi.OutputState }

func (EventHandlerTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHandlerTemplate)(nil)).Elem()
}

func (o EventHandlerTemplateArrayOutput) ToEventHandlerTemplateArrayOutput() EventHandlerTemplateArrayOutput {
	return o
}

func (o EventHandlerTemplateArrayOutput) ToEventHandlerTemplateArrayOutputWithContext(ctx context.Context) EventHandlerTemplateArrayOutput {
	return o
}

func (o EventHandlerTemplateArrayOutput) Index(i pulumi.IntInput) EventHandlerTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventHandlerTemplate {
		return vs[0].([]EventHandlerTemplate)[vs[1].(int)]
	}).(EventHandlerTemplateOutput)
}

type EventHandlerTemplateResponse struct {
	Auth               *UpstreamAuthSettingsResponse `pulumi:"auth"`
	SystemEventPattern *string                       `pulumi:"systemEventPattern"`
	UrlTemplate        string                        `pulumi:"urlTemplate"`
	UserEventPattern   *string                       `pulumi:"userEventPattern"`
}





type EventHandlerTemplateResponseInput interface {
	pulumi.Input

	ToEventHandlerTemplateResponseOutput() EventHandlerTemplateResponseOutput
	ToEventHandlerTemplateResponseOutputWithContext(context.Context) EventHandlerTemplateResponseOutput
}

type EventHandlerTemplateResponseArgs struct {
	Auth               UpstreamAuthSettingsResponsePtrInput `pulumi:"auth"`
	SystemEventPattern pulumi.StringPtrInput                `pulumi:"systemEventPattern"`
	UrlTemplate        pulumi.StringInput                   `pulumi:"urlTemplate"`
	UserEventPattern   pulumi.StringPtrInput                `pulumi:"userEventPattern"`
}

func (EventHandlerTemplateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerTemplateResponse)(nil)).Elem()
}

func (i EventHandlerTemplateResponseArgs) ToEventHandlerTemplateResponseOutput() EventHandlerTemplateResponseOutput {
	return i.ToEventHandlerTemplateResponseOutputWithContext(context.Background())
}

func (i EventHandlerTemplateResponseArgs) ToEventHandlerTemplateResponseOutputWithContext(ctx context.Context) EventHandlerTemplateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerTemplateResponseOutput)
}





type EventHandlerTemplateResponseArrayInput interface {
	pulumi.Input

	ToEventHandlerTemplateResponseArrayOutput() EventHandlerTemplateResponseArrayOutput
	ToEventHandlerTemplateResponseArrayOutputWithContext(context.Context) EventHandlerTemplateResponseArrayOutput
}

type EventHandlerTemplateResponseArray []EventHandlerTemplateResponseInput

func (EventHandlerTemplateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHandlerTemplateResponse)(nil)).Elem()
}

func (i EventHandlerTemplateResponseArray) ToEventHandlerTemplateResponseArrayOutput() EventHandlerTemplateResponseArrayOutput {
	return i.ToEventHandlerTemplateResponseArrayOutputWithContext(context.Background())
}

func (i EventHandlerTemplateResponseArray) ToEventHandlerTemplateResponseArrayOutputWithContext(ctx context.Context) EventHandlerTemplateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerTemplateResponseArrayOutput)
}

type EventHandlerTemplateResponseOutput struct{ *pulumi.OutputState }

func (EventHandlerTemplateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerTemplateResponse)(nil)).Elem()
}

func (o EventHandlerTemplateResponseOutput) ToEventHandlerTemplateResponseOutput() EventHandlerTemplateResponseOutput {
	return o
}

func (o EventHandlerTemplateResponseOutput) ToEventHandlerTemplateResponseOutputWithContext(ctx context.Context) EventHandlerTemplateResponseOutput {
	return o
}

func (o EventHandlerTemplateResponseOutput) Auth() UpstreamAuthSettingsResponsePtrOutput {
	return o.ApplyT(func(v EventHandlerTemplateResponse) *UpstreamAuthSettingsResponse { return v.Auth }).(UpstreamAuthSettingsResponsePtrOutput)
}

func (o EventHandlerTemplateResponseOutput) SystemEventPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHandlerTemplateResponse) *string { return v.SystemEventPattern }).(pulumi.StringPtrOutput)
}

func (o EventHandlerTemplateResponseOutput) UrlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v EventHandlerTemplateResponse) string { return v.UrlTemplate }).(pulumi.StringOutput)
}

func (o EventHandlerTemplateResponseOutput) UserEventPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHandlerTemplateResponse) *string { return v.UserEventPattern }).(pulumi.StringPtrOutput)
}

type EventHandlerTemplateResponseArrayOutput struct{ *pulumi.OutputState }

func (EventHandlerTemplateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHandlerTemplateResponse)(nil)).Elem()
}

func (o EventHandlerTemplateResponseArrayOutput) ToEventHandlerTemplateResponseArrayOutput() EventHandlerTemplateResponseArrayOutput {
	return o
}

func (o EventHandlerTemplateResponseArrayOutput) ToEventHandlerTemplateResponseArrayOutputWithContext(ctx context.Context) EventHandlerTemplateResponseArrayOutput {
	return o
}

func (o EventHandlerTemplateResponseArrayOutput) Index(i pulumi.IntInput) EventHandlerTemplateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventHandlerTemplateResponse {
		return vs[0].([]EventHandlerTemplateResponse)[vs[1].(int)]
	}).(EventHandlerTemplateResponseOutput)
}

type LiveTraceCategory struct {
	Enabled *string `pulumi:"enabled"`
	Name    *string `pulumi:"name"`
}





type LiveTraceCategoryInput interface {
	pulumi.Input

	ToLiveTraceCategoryOutput() LiveTraceCategoryOutput
	ToLiveTraceCategoryOutputWithContext(context.Context) LiveTraceCategoryOutput
}

type LiveTraceCategoryArgs struct {
	Enabled pulumi.StringPtrInput `pulumi:"enabled"`
	Name    pulumi.StringPtrInput `pulumi:"name"`
}

func (LiveTraceCategoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceCategory)(nil)).Elem()
}

func (i LiveTraceCategoryArgs) ToLiveTraceCategoryOutput() LiveTraceCategoryOutput {
	return i.ToLiveTraceCategoryOutputWithContext(context.Background())
}

func (i LiveTraceCategoryArgs) ToLiveTraceCategoryOutputWithContext(ctx context.Context) LiveTraceCategoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceCategoryOutput)
}





type LiveTraceCategoryArrayInput interface {
	pulumi.Input

	ToLiveTraceCategoryArrayOutput() LiveTraceCategoryArrayOutput
	ToLiveTraceCategoryArrayOutputWithContext(context.Context) LiveTraceCategoryArrayOutput
}

type LiveTraceCategoryArray []LiveTraceCategoryInput

func (LiveTraceCategoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveTraceCategory)(nil)).Elem()
}

func (i LiveTraceCategoryArray) ToLiveTraceCategoryArrayOutput() LiveTraceCategoryArrayOutput {
	return i.ToLiveTraceCategoryArrayOutputWithContext(context.Background())
}

func (i LiveTraceCategoryArray) ToLiveTraceCategoryArrayOutputWithContext(ctx context.Context) LiveTraceCategoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceCategoryArrayOutput)
}

type LiveTraceCategoryOutput struct{ *pulumi.OutputState }

func (LiveTraceCategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceCategory)(nil)).Elem()
}

func (o LiveTraceCategoryOutput) ToLiveTraceCategoryOutput() LiveTraceCategoryOutput {
	return o
}

func (o LiveTraceCategoryOutput) ToLiveTraceCategoryOutputWithContext(ctx context.Context) LiveTraceCategoryOutput {
	return o
}

func (o LiveTraceCategoryOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceCategory) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o LiveTraceCategoryOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceCategory) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type LiveTraceCategoryArrayOutput struct{ *pulumi.OutputState }

func (LiveTraceCategoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveTraceCategory)(nil)).Elem()
}

func (o LiveTraceCategoryArrayOutput) ToLiveTraceCategoryArrayOutput() LiveTraceCategoryArrayOutput {
	return o
}

func (o LiveTraceCategoryArrayOutput) ToLiveTraceCategoryArrayOutputWithContext(ctx context.Context) LiveTraceCategoryArrayOutput {
	return o
}

func (o LiveTraceCategoryArrayOutput) Index(i pulumi.IntInput) LiveTraceCategoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveTraceCategory {
		return vs[0].([]LiveTraceCategory)[vs[1].(int)]
	}).(LiveTraceCategoryOutput)
}

type LiveTraceCategoryResponse struct {
	Enabled *string `pulumi:"enabled"`
	Name    *string `pulumi:"name"`
}





type LiveTraceCategoryResponseInput interface {
	pulumi.Input

	ToLiveTraceCategoryResponseOutput() LiveTraceCategoryResponseOutput
	ToLiveTraceCategoryResponseOutputWithContext(context.Context) LiveTraceCategoryResponseOutput
}

type LiveTraceCategoryResponseArgs struct {
	Enabled pulumi.StringPtrInput `pulumi:"enabled"`
	Name    pulumi.StringPtrInput `pulumi:"name"`
}

func (LiveTraceCategoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceCategoryResponse)(nil)).Elem()
}

func (i LiveTraceCategoryResponseArgs) ToLiveTraceCategoryResponseOutput() LiveTraceCategoryResponseOutput {
	return i.ToLiveTraceCategoryResponseOutputWithContext(context.Background())
}

func (i LiveTraceCategoryResponseArgs) ToLiveTraceCategoryResponseOutputWithContext(ctx context.Context) LiveTraceCategoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceCategoryResponseOutput)
}





type LiveTraceCategoryResponseArrayInput interface {
	pulumi.Input

	ToLiveTraceCategoryResponseArrayOutput() LiveTraceCategoryResponseArrayOutput
	ToLiveTraceCategoryResponseArrayOutputWithContext(context.Context) LiveTraceCategoryResponseArrayOutput
}

type LiveTraceCategoryResponseArray []LiveTraceCategoryResponseInput

func (LiveTraceCategoryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveTraceCategoryResponse)(nil)).Elem()
}

func (i LiveTraceCategoryResponseArray) ToLiveTraceCategoryResponseArrayOutput() LiveTraceCategoryResponseArrayOutput {
	return i.ToLiveTraceCategoryResponseArrayOutputWithContext(context.Background())
}

func (i LiveTraceCategoryResponseArray) ToLiveTraceCategoryResponseArrayOutputWithContext(ctx context.Context) LiveTraceCategoryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceCategoryResponseArrayOutput)
}

type LiveTraceCategoryResponseOutput struct{ *pulumi.OutputState }

func (LiveTraceCategoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceCategoryResponse)(nil)).Elem()
}

func (o LiveTraceCategoryResponseOutput) ToLiveTraceCategoryResponseOutput() LiveTraceCategoryResponseOutput {
	return o
}

func (o LiveTraceCategoryResponseOutput) ToLiveTraceCategoryResponseOutputWithContext(ctx context.Context) LiveTraceCategoryResponseOutput {
	return o
}

func (o LiveTraceCategoryResponseOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceCategoryResponse) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o LiveTraceCategoryResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceCategoryResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type LiveTraceCategoryResponseArrayOutput struct{ *pulumi.OutputState }

func (LiveTraceCategoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveTraceCategoryResponse)(nil)).Elem()
}

func (o LiveTraceCategoryResponseArrayOutput) ToLiveTraceCategoryResponseArrayOutput() LiveTraceCategoryResponseArrayOutput {
	return o
}

func (o LiveTraceCategoryResponseArrayOutput) ToLiveTraceCategoryResponseArrayOutputWithContext(ctx context.Context) LiveTraceCategoryResponseArrayOutput {
	return o
}

func (o LiveTraceCategoryResponseArrayOutput) Index(i pulumi.IntInput) LiveTraceCategoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveTraceCategoryResponse {
		return vs[0].([]LiveTraceCategoryResponse)[vs[1].(int)]
	}).(LiveTraceCategoryResponseOutput)
}

type LiveTraceConfiguration struct {
	Categories []LiveTraceCategory `pulumi:"categories"`
	Enabled    *string             `pulumi:"enabled"`
}





type LiveTraceConfigurationInput interface {
	pulumi.Input

	ToLiveTraceConfigurationOutput() LiveTraceConfigurationOutput
	ToLiveTraceConfigurationOutputWithContext(context.Context) LiveTraceConfigurationOutput
}

type LiveTraceConfigurationArgs struct {
	Categories LiveTraceCategoryArrayInput `pulumi:"categories"`
	Enabled    pulumi.StringPtrInput       `pulumi:"enabled"`
}

func (LiveTraceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceConfiguration)(nil)).Elem()
}

func (i LiveTraceConfigurationArgs) ToLiveTraceConfigurationOutput() LiveTraceConfigurationOutput {
	return i.ToLiveTraceConfigurationOutputWithContext(context.Background())
}

func (i LiveTraceConfigurationArgs) ToLiveTraceConfigurationOutputWithContext(ctx context.Context) LiveTraceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceConfigurationOutput)
}

func (i LiveTraceConfigurationArgs) ToLiveTraceConfigurationPtrOutput() LiveTraceConfigurationPtrOutput {
	return i.ToLiveTraceConfigurationPtrOutputWithContext(context.Background())
}

func (i LiveTraceConfigurationArgs) ToLiveTraceConfigurationPtrOutputWithContext(ctx context.Context) LiveTraceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceConfigurationOutput).ToLiveTraceConfigurationPtrOutputWithContext(ctx)
}









type LiveTraceConfigurationPtrInput interface {
	pulumi.Input

	ToLiveTraceConfigurationPtrOutput() LiveTraceConfigurationPtrOutput
	ToLiveTraceConfigurationPtrOutputWithContext(context.Context) LiveTraceConfigurationPtrOutput
}

type liveTraceConfigurationPtrType LiveTraceConfigurationArgs

func LiveTraceConfigurationPtr(v *LiveTraceConfigurationArgs) LiveTraceConfigurationPtrInput {
	return (*liveTraceConfigurationPtrType)(v)
}

func (*liveTraceConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveTraceConfiguration)(nil)).Elem()
}

func (i *liveTraceConfigurationPtrType) ToLiveTraceConfigurationPtrOutput() LiveTraceConfigurationPtrOutput {
	return i.ToLiveTraceConfigurationPtrOutputWithContext(context.Background())
}

func (i *liveTraceConfigurationPtrType) ToLiveTraceConfigurationPtrOutputWithContext(ctx context.Context) LiveTraceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceConfigurationPtrOutput)
}

type LiveTraceConfigurationOutput struct{ *pulumi.OutputState }

func (LiveTraceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceConfiguration)(nil)).Elem()
}

func (o LiveTraceConfigurationOutput) ToLiveTraceConfigurationOutput() LiveTraceConfigurationOutput {
	return o
}

func (o LiveTraceConfigurationOutput) ToLiveTraceConfigurationOutputWithContext(ctx context.Context) LiveTraceConfigurationOutput {
	return o
}

func (o LiveTraceConfigurationOutput) ToLiveTraceConfigurationPtrOutput() LiveTraceConfigurationPtrOutput {
	return o.ToLiveTraceConfigurationPtrOutputWithContext(context.Background())
}

func (o LiveTraceConfigurationOutput) ToLiveTraceConfigurationPtrOutputWithContext(ctx context.Context) LiveTraceConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveTraceConfiguration) *LiveTraceConfiguration {
		return &v
	}).(LiveTraceConfigurationPtrOutput)
}

func (o LiveTraceConfigurationOutput) Categories() LiveTraceCategoryArrayOutput {
	return o.ApplyT(func(v LiveTraceConfiguration) []LiveTraceCategory { return v.Categories }).(LiveTraceCategoryArrayOutput)
}

func (o LiveTraceConfigurationOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceConfiguration) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

type LiveTraceConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LiveTraceConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveTraceConfiguration)(nil)).Elem()
}

func (o LiveTraceConfigurationPtrOutput) ToLiveTraceConfigurationPtrOutput() LiveTraceConfigurationPtrOutput {
	return o
}

func (o LiveTraceConfigurationPtrOutput) ToLiveTraceConfigurationPtrOutputWithContext(ctx context.Context) LiveTraceConfigurationPtrOutput {
	return o
}

func (o LiveTraceConfigurationPtrOutput) Elem() LiveTraceConfigurationOutput {
	return o.ApplyT(func(v *LiveTraceConfiguration) LiveTraceConfiguration {
		if v != nil {
			return *v
		}
		var ret LiveTraceConfiguration
		return ret
	}).(LiveTraceConfigurationOutput)
}

func (o LiveTraceConfigurationPtrOutput) Categories() LiveTraceCategoryArrayOutput {
	return o.ApplyT(func(v *LiveTraceConfiguration) []LiveTraceCategory {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(LiveTraceCategoryArrayOutput)
}

func (o LiveTraceConfigurationPtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveTraceConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

type LiveTraceConfigurationResponse struct {
	Categories []LiveTraceCategoryResponse `pulumi:"categories"`
	Enabled    *string                     `pulumi:"enabled"`
}





type LiveTraceConfigurationResponseInput interface {
	pulumi.Input

	ToLiveTraceConfigurationResponseOutput() LiveTraceConfigurationResponseOutput
	ToLiveTraceConfigurationResponseOutputWithContext(context.Context) LiveTraceConfigurationResponseOutput
}

type LiveTraceConfigurationResponseArgs struct {
	Categories LiveTraceCategoryResponseArrayInput `pulumi:"categories"`
	Enabled    pulumi.StringPtrInput               `pulumi:"enabled"`
}

func (LiveTraceConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceConfigurationResponse)(nil)).Elem()
}

func (i LiveTraceConfigurationResponseArgs) ToLiveTraceConfigurationResponseOutput() LiveTraceConfigurationResponseOutput {
	return i.ToLiveTraceConfigurationResponseOutputWithContext(context.Background())
}

func (i LiveTraceConfigurationResponseArgs) ToLiveTraceConfigurationResponseOutputWithContext(ctx context.Context) LiveTraceConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceConfigurationResponseOutput)
}

func (i LiveTraceConfigurationResponseArgs) ToLiveTraceConfigurationResponsePtrOutput() LiveTraceConfigurationResponsePtrOutput {
	return i.ToLiveTraceConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i LiveTraceConfigurationResponseArgs) ToLiveTraceConfigurationResponsePtrOutputWithContext(ctx context.Context) LiveTraceConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceConfigurationResponseOutput).ToLiveTraceConfigurationResponsePtrOutputWithContext(ctx)
}









type LiveTraceConfigurationResponsePtrInput interface {
	pulumi.Input

	ToLiveTraceConfigurationResponsePtrOutput() LiveTraceConfigurationResponsePtrOutput
	ToLiveTraceConfigurationResponsePtrOutputWithContext(context.Context) LiveTraceConfigurationResponsePtrOutput
}

type liveTraceConfigurationResponsePtrType LiveTraceConfigurationResponseArgs

func LiveTraceConfigurationResponsePtr(v *LiveTraceConfigurationResponseArgs) LiveTraceConfigurationResponsePtrInput {
	return (*liveTraceConfigurationResponsePtrType)(v)
}

func (*liveTraceConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveTraceConfigurationResponse)(nil)).Elem()
}

func (i *liveTraceConfigurationResponsePtrType) ToLiveTraceConfigurationResponsePtrOutput() LiveTraceConfigurationResponsePtrOutput {
	return i.ToLiveTraceConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *liveTraceConfigurationResponsePtrType) ToLiveTraceConfigurationResponsePtrOutputWithContext(ctx context.Context) LiveTraceConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceConfigurationResponsePtrOutput)
}

type LiveTraceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LiveTraceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceConfigurationResponse)(nil)).Elem()
}

func (o LiveTraceConfigurationResponseOutput) ToLiveTraceConfigurationResponseOutput() LiveTraceConfigurationResponseOutput {
	return o
}

func (o LiveTraceConfigurationResponseOutput) ToLiveTraceConfigurationResponseOutputWithContext(ctx context.Context) LiveTraceConfigurationResponseOutput {
	return o
}

func (o LiveTraceConfigurationResponseOutput) ToLiveTraceConfigurationResponsePtrOutput() LiveTraceConfigurationResponsePtrOutput {
	return o.ToLiveTraceConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o LiveTraceConfigurationResponseOutput) ToLiveTraceConfigurationResponsePtrOutputWithContext(ctx context.Context) LiveTraceConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveTraceConfigurationResponse) *LiveTraceConfigurationResponse {
		return &v
	}).(LiveTraceConfigurationResponsePtrOutput)
}

func (o LiveTraceConfigurationResponseOutput) Categories() LiveTraceCategoryResponseArrayOutput {
	return o.ApplyT(func(v LiveTraceConfigurationResponse) []LiveTraceCategoryResponse { return v.Categories }).(LiveTraceCategoryResponseArrayOutput)
}

func (o LiveTraceConfigurationResponseOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceConfigurationResponse) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

type LiveTraceConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveTraceConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveTraceConfigurationResponse)(nil)).Elem()
}

func (o LiveTraceConfigurationResponsePtrOutput) ToLiveTraceConfigurationResponsePtrOutput() LiveTraceConfigurationResponsePtrOutput {
	return o
}

func (o LiveTraceConfigurationResponsePtrOutput) ToLiveTraceConfigurationResponsePtrOutputWithContext(ctx context.Context) LiveTraceConfigurationResponsePtrOutput {
	return o
}

func (o LiveTraceConfigurationResponsePtrOutput) Elem() LiveTraceConfigurationResponseOutput {
	return o.ApplyT(func(v *LiveTraceConfigurationResponse) LiveTraceConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LiveTraceConfigurationResponse
		return ret
	}).(LiveTraceConfigurationResponseOutput)
}

func (o LiveTraceConfigurationResponsePtrOutput) Categories() LiveTraceCategoryResponseArrayOutput {
	return o.ApplyT(func(v *LiveTraceConfigurationResponse) []LiveTraceCategoryResponse {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(LiveTraceCategoryResponseArrayOutput)
}

func (o LiveTraceConfigurationResponsePtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveTraceConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentity struct {
	Type                   *string                `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedIdentityInput interface {
	pulumi.Input

	ToManagedIdentityOutput() ManagedIdentityOutput
	ToManagedIdentityOutputWithContext(context.Context) ManagedIdentityOutput
}

type ManagedIdentityArgs struct {
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput       `pulumi:"userAssignedIdentities"`
}

func (ManagedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentity)(nil)).Elem()
}

func (i ManagedIdentityArgs) ToManagedIdentityOutput() ManagedIdentityOutput {
	return i.ToManagedIdentityOutputWithContext(context.Background())
}

func (i ManagedIdentityArgs) ToManagedIdentityOutputWithContext(ctx context.Context) ManagedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityOutput)
}

func (i ManagedIdentityArgs) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return i.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedIdentityArgs) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityOutput).ToManagedIdentityPtrOutputWithContext(ctx)
}









type ManagedIdentityPtrInput interface {
	pulumi.Input

	ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput
	ToManagedIdentityPtrOutputWithContext(context.Context) ManagedIdentityPtrOutput
}

type managedIdentityPtrType ManagedIdentityArgs

func ManagedIdentityPtr(v *ManagedIdentityArgs) ManagedIdentityPtrInput {
	return (*managedIdentityPtrType)(v)
}

func (*managedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentity)(nil)).Elem()
}

func (i *managedIdentityPtrType) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return i.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (i *managedIdentityPtrType) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPtrOutput)
}

type ManagedIdentityOutput struct{ *pulumi.OutputState }

func (ManagedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentity)(nil)).Elem()
}

func (o ManagedIdentityOutput) ToManagedIdentityOutput() ManagedIdentityOutput {
	return o
}

func (o ManagedIdentityOutput) ToManagedIdentityOutputWithContext(ctx context.Context) ManagedIdentityOutput {
	return o
}

func (o ManagedIdentityOutput) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return o.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityOutput) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentity) *ManagedIdentity {
		return &v
	}).(ManagedIdentityPtrOutput)
}

func (o ManagedIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentity)(nil)).Elem()
}

func (o ManagedIdentityPtrOutput) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return o
}

func (o ManagedIdentityPtrOutput) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return o
}

func (o ManagedIdentityPtrOutput) Elem() ManagedIdentityOutput {
	return o.ApplyT(func(v *ManagedIdentity) ManagedIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedIdentity
		return ret
	}).(ManagedIdentityOutput)
}

func (o ManagedIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedIdentityResponse struct {
	PrincipalId            string                                          `pulumi:"principalId"`
	TenantId               string                                          `pulumi:"tenantId"`
	Type                   *string                                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityPropertyResponse `pulumi:"userAssignedIdentities"`
}





type ManagedIdentityResponseInput interface {
	pulumi.Input

	ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput
	ToManagedIdentityResponseOutputWithContext(context.Context) ManagedIdentityResponseOutput
}

type ManagedIdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                           `pulumi:"principalId"`
	TenantId               pulumi.StringInput                           `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                        `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityPropertyResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (ManagedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityResponse)(nil)).Elem()
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput {
	return i.ToManagedIdentityResponseOutputWithContext(context.Background())
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponseOutputWithContext(ctx context.Context) ManagedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityResponseOutput)
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return i.ToManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityResponseOutput).ToManagedIdentityResponsePtrOutputWithContext(ctx)
}









type ManagedIdentityResponsePtrInput interface {
	pulumi.Input

	ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput
	ToManagedIdentityResponsePtrOutputWithContext(context.Context) ManagedIdentityResponsePtrOutput
}

type managedIdentityResponsePtrType ManagedIdentityResponseArgs

func ManagedIdentityResponsePtr(v *ManagedIdentityResponseArgs) ManagedIdentityResponsePtrInput {
	return (*managedIdentityResponsePtrType)(v)
}

func (*managedIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityResponse)(nil)).Elem()
}

func (i *managedIdentityResponsePtrType) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return i.ToManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *managedIdentityResponsePtrType) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityResponsePtrOutput)
}

type ManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityResponse)(nil)).Elem()
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput {
	return o
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponseOutputWithContext(ctx context.Context) ManagedIdentityResponseOutput {
	return o
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return o.ToManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityResponse) *ManagedIdentityResponse {
		return &v
	}).(ManagedIdentityResponsePtrOutput)
}

func (o ManagedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityPropertyResponseMapOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) map[string]UserAssignedIdentityPropertyResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityPropertyResponseMapOutput)
}

type ManagedIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityResponse)(nil)).Elem()
}

func (o ManagedIdentityResponsePtrOutput) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return o
}

func (o ManagedIdentityResponsePtrOutput) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return o
}

func (o ManagedIdentityResponsePtrOutput) Elem() ManagedIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) ManagedIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityResponse
		return ret
	}).(ManagedIdentityResponseOutput)
}

func (o ManagedIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityPropertyResponseMapOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) map[string]UserAssignedIdentityPropertyResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityPropertyResponseMapOutput)
}

type ManagedIdentitySettings struct {
	Resource *string `pulumi:"resource"`
}





type ManagedIdentitySettingsInput interface {
	pulumi.Input

	ToManagedIdentitySettingsOutput() ManagedIdentitySettingsOutput
	ToManagedIdentitySettingsOutputWithContext(context.Context) ManagedIdentitySettingsOutput
}

type ManagedIdentitySettingsArgs struct {
	Resource pulumi.StringPtrInput `pulumi:"resource"`
}

func (ManagedIdentitySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentitySettings)(nil)).Elem()
}

func (i ManagedIdentitySettingsArgs) ToManagedIdentitySettingsOutput() ManagedIdentitySettingsOutput {
	return i.ToManagedIdentitySettingsOutputWithContext(context.Background())
}

func (i ManagedIdentitySettingsArgs) ToManagedIdentitySettingsOutputWithContext(ctx context.Context) ManagedIdentitySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentitySettingsOutput)
}

func (i ManagedIdentitySettingsArgs) ToManagedIdentitySettingsPtrOutput() ManagedIdentitySettingsPtrOutput {
	return i.ToManagedIdentitySettingsPtrOutputWithContext(context.Background())
}

func (i ManagedIdentitySettingsArgs) ToManagedIdentitySettingsPtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentitySettingsOutput).ToManagedIdentitySettingsPtrOutputWithContext(ctx)
}









type ManagedIdentitySettingsPtrInput interface {
	pulumi.Input

	ToManagedIdentitySettingsPtrOutput() ManagedIdentitySettingsPtrOutput
	ToManagedIdentitySettingsPtrOutputWithContext(context.Context) ManagedIdentitySettingsPtrOutput
}

type managedIdentitySettingsPtrType ManagedIdentitySettingsArgs

func ManagedIdentitySettingsPtr(v *ManagedIdentitySettingsArgs) ManagedIdentitySettingsPtrInput {
	return (*managedIdentitySettingsPtrType)(v)
}

func (*managedIdentitySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentitySettings)(nil)).Elem()
}

func (i *managedIdentitySettingsPtrType) ToManagedIdentitySettingsPtrOutput() ManagedIdentitySettingsPtrOutput {
	return i.ToManagedIdentitySettingsPtrOutputWithContext(context.Background())
}

func (i *managedIdentitySettingsPtrType) ToManagedIdentitySettingsPtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentitySettingsPtrOutput)
}

type ManagedIdentitySettingsOutput struct{ *pulumi.OutputState }

func (ManagedIdentitySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentitySettings)(nil)).Elem()
}

func (o ManagedIdentitySettingsOutput) ToManagedIdentitySettingsOutput() ManagedIdentitySettingsOutput {
	return o
}

func (o ManagedIdentitySettingsOutput) ToManagedIdentitySettingsOutputWithContext(ctx context.Context) ManagedIdentitySettingsOutput {
	return o
}

func (o ManagedIdentitySettingsOutput) ToManagedIdentitySettingsPtrOutput() ManagedIdentitySettingsPtrOutput {
	return o.ToManagedIdentitySettingsPtrOutputWithContext(context.Background())
}

func (o ManagedIdentitySettingsOutput) ToManagedIdentitySettingsPtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentitySettings) *ManagedIdentitySettings {
		return &v
	}).(ManagedIdentitySettingsPtrOutput)
}

func (o ManagedIdentitySettingsOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentitySettings) *string { return v.Resource }).(pulumi.StringPtrOutput)
}

type ManagedIdentitySettingsPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentitySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentitySettings)(nil)).Elem()
}

func (o ManagedIdentitySettingsPtrOutput) ToManagedIdentitySettingsPtrOutput() ManagedIdentitySettingsPtrOutput {
	return o
}

func (o ManagedIdentitySettingsPtrOutput) ToManagedIdentitySettingsPtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsPtrOutput {
	return o
}

func (o ManagedIdentitySettingsPtrOutput) Elem() ManagedIdentitySettingsOutput {
	return o.ApplyT(func(v *ManagedIdentitySettings) ManagedIdentitySettings {
		if v != nil {
			return *v
		}
		var ret ManagedIdentitySettings
		return ret
	}).(ManagedIdentitySettingsOutput)
}

func (o ManagedIdentitySettingsPtrOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentitySettings) *string {
		if v == nil {
			return nil
		}
		return v.Resource
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentitySettingsResponse struct {
	Resource *string `pulumi:"resource"`
}





type ManagedIdentitySettingsResponseInput interface {
	pulumi.Input

	ToManagedIdentitySettingsResponseOutput() ManagedIdentitySettingsResponseOutput
	ToManagedIdentitySettingsResponseOutputWithContext(context.Context) ManagedIdentitySettingsResponseOutput
}

type ManagedIdentitySettingsResponseArgs struct {
	Resource pulumi.StringPtrInput `pulumi:"resource"`
}

func (ManagedIdentitySettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentitySettingsResponse)(nil)).Elem()
}

func (i ManagedIdentitySettingsResponseArgs) ToManagedIdentitySettingsResponseOutput() ManagedIdentitySettingsResponseOutput {
	return i.ToManagedIdentitySettingsResponseOutputWithContext(context.Background())
}

func (i ManagedIdentitySettingsResponseArgs) ToManagedIdentitySettingsResponseOutputWithContext(ctx context.Context) ManagedIdentitySettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentitySettingsResponseOutput)
}

func (i ManagedIdentitySettingsResponseArgs) ToManagedIdentitySettingsResponsePtrOutput() ManagedIdentitySettingsResponsePtrOutput {
	return i.ToManagedIdentitySettingsResponsePtrOutputWithContext(context.Background())
}

func (i ManagedIdentitySettingsResponseArgs) ToManagedIdentitySettingsResponsePtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentitySettingsResponseOutput).ToManagedIdentitySettingsResponsePtrOutputWithContext(ctx)
}









type ManagedIdentitySettingsResponsePtrInput interface {
	pulumi.Input

	ToManagedIdentitySettingsResponsePtrOutput() ManagedIdentitySettingsResponsePtrOutput
	ToManagedIdentitySettingsResponsePtrOutputWithContext(context.Context) ManagedIdentitySettingsResponsePtrOutput
}

type managedIdentitySettingsResponsePtrType ManagedIdentitySettingsResponseArgs

func ManagedIdentitySettingsResponsePtr(v *ManagedIdentitySettingsResponseArgs) ManagedIdentitySettingsResponsePtrInput {
	return (*managedIdentitySettingsResponsePtrType)(v)
}

func (*managedIdentitySettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentitySettingsResponse)(nil)).Elem()
}

func (i *managedIdentitySettingsResponsePtrType) ToManagedIdentitySettingsResponsePtrOutput() ManagedIdentitySettingsResponsePtrOutput {
	return i.ToManagedIdentitySettingsResponsePtrOutputWithContext(context.Background())
}

func (i *managedIdentitySettingsResponsePtrType) ToManagedIdentitySettingsResponsePtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentitySettingsResponsePtrOutput)
}

type ManagedIdentitySettingsResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentitySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentitySettingsResponse)(nil)).Elem()
}

func (o ManagedIdentitySettingsResponseOutput) ToManagedIdentitySettingsResponseOutput() ManagedIdentitySettingsResponseOutput {
	return o
}

func (o ManagedIdentitySettingsResponseOutput) ToManagedIdentitySettingsResponseOutputWithContext(ctx context.Context) ManagedIdentitySettingsResponseOutput {
	return o
}

func (o ManagedIdentitySettingsResponseOutput) ToManagedIdentitySettingsResponsePtrOutput() ManagedIdentitySettingsResponsePtrOutput {
	return o.ToManagedIdentitySettingsResponsePtrOutputWithContext(context.Background())
}

func (o ManagedIdentitySettingsResponseOutput) ToManagedIdentitySettingsResponsePtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentitySettingsResponse) *ManagedIdentitySettingsResponse {
		return &v
	}).(ManagedIdentitySettingsResponsePtrOutput)
}

func (o ManagedIdentitySettingsResponseOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentitySettingsResponse) *string { return v.Resource }).(pulumi.StringPtrOutput)
}

type ManagedIdentitySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentitySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentitySettingsResponse)(nil)).Elem()
}

func (o ManagedIdentitySettingsResponsePtrOutput) ToManagedIdentitySettingsResponsePtrOutput() ManagedIdentitySettingsResponsePtrOutput {
	return o
}

func (o ManagedIdentitySettingsResponsePtrOutput) ToManagedIdentitySettingsResponsePtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsResponsePtrOutput {
	return o
}

func (o ManagedIdentitySettingsResponsePtrOutput) Elem() ManagedIdentitySettingsResponseOutput {
	return o.ApplyT(func(v *ManagedIdentitySettingsResponse) ManagedIdentitySettingsResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentitySettingsResponse
		return ret
	}).(ManagedIdentitySettingsResponseOutput)
}

func (o ManagedIdentitySettingsResponsePtrOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentitySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Resource
	}).(pulumi.StringPtrOutput)
}

type NetworkACL struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
}





type NetworkACLInput interface {
	pulumi.Input

	ToNetworkACLOutput() NetworkACLOutput
	ToNetworkACLOutputWithContext(context.Context) NetworkACLOutput
}

type NetworkACLArgs struct {
	Allow pulumi.StringArrayInput `pulumi:"allow"`
	Deny  pulumi.StringArrayInput `pulumi:"deny"`
}

func (NetworkACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkACL)(nil)).Elem()
}

func (i NetworkACLArgs) ToNetworkACLOutput() NetworkACLOutput {
	return i.ToNetworkACLOutputWithContext(context.Background())
}

func (i NetworkACLArgs) ToNetworkACLOutputWithContext(ctx context.Context) NetworkACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLOutput)
}

func (i NetworkACLArgs) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return i.ToNetworkACLPtrOutputWithContext(context.Background())
}

func (i NetworkACLArgs) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLOutput).ToNetworkACLPtrOutputWithContext(ctx)
}









type NetworkACLPtrInput interface {
	pulumi.Input

	ToNetworkACLPtrOutput() NetworkACLPtrOutput
	ToNetworkACLPtrOutputWithContext(context.Context) NetworkACLPtrOutput
}

type networkACLPtrType NetworkACLArgs

func NetworkACLPtr(v *NetworkACLArgs) NetworkACLPtrInput {
	return (*networkACLPtrType)(v)
}

func (*networkACLPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkACL)(nil)).Elem()
}

func (i *networkACLPtrType) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return i.ToNetworkACLPtrOutputWithContext(context.Background())
}

func (i *networkACLPtrType) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLPtrOutput)
}

type NetworkACLOutput struct{ *pulumi.OutputState }

func (NetworkACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkACL)(nil)).Elem()
}

func (o NetworkACLOutput) ToNetworkACLOutput() NetworkACLOutput {
	return o
}

func (o NetworkACLOutput) ToNetworkACLOutputWithContext(ctx context.Context) NetworkACLOutput {
	return o
}

func (o NetworkACLOutput) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return o.ToNetworkACLPtrOutputWithContext(context.Background())
}

func (o NetworkACLOutput) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkACL) *NetworkACL {
		return &v
	}).(NetworkACLPtrOutput)
}

func (o NetworkACLOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACL) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o NetworkACLOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACL) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

type NetworkACLPtrOutput struct{ *pulumi.OutputState }

func (NetworkACLPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkACL)(nil)).Elem()
}

func (o NetworkACLPtrOutput) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return o
}

func (o NetworkACLPtrOutput) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return o
}

func (o NetworkACLPtrOutput) Elem() NetworkACLOutput {
	return o.ApplyT(func(v *NetworkACL) NetworkACL {
		if v != nil {
			return *v
		}
		var ret NetworkACL
		return ret
	}).(NetworkACLOutput)
}

func (o NetworkACLPtrOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACL) []string {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(pulumi.StringArrayOutput)
}

func (o NetworkACLPtrOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACL) []string {
		if v == nil {
			return nil
		}
		return v.Deny
	}).(pulumi.StringArrayOutput)
}

type NetworkACLResponse struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
}





type NetworkACLResponseInput interface {
	pulumi.Input

	ToNetworkACLResponseOutput() NetworkACLResponseOutput
	ToNetworkACLResponseOutputWithContext(context.Context) NetworkACLResponseOutput
}

type NetworkACLResponseArgs struct {
	Allow pulumi.StringArrayInput `pulumi:"allow"`
	Deny  pulumi.StringArrayInput `pulumi:"deny"`
}

func (NetworkACLResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkACLResponse)(nil)).Elem()
}

func (i NetworkACLResponseArgs) ToNetworkACLResponseOutput() NetworkACLResponseOutput {
	return i.ToNetworkACLResponseOutputWithContext(context.Background())
}

func (i NetworkACLResponseArgs) ToNetworkACLResponseOutputWithContext(ctx context.Context) NetworkACLResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLResponseOutput)
}

func (i NetworkACLResponseArgs) ToNetworkACLResponsePtrOutput() NetworkACLResponsePtrOutput {
	return i.ToNetworkACLResponsePtrOutputWithContext(context.Background())
}

func (i NetworkACLResponseArgs) ToNetworkACLResponsePtrOutputWithContext(ctx context.Context) NetworkACLResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLResponseOutput).ToNetworkACLResponsePtrOutputWithContext(ctx)
}









type NetworkACLResponsePtrInput interface {
	pulumi.Input

	ToNetworkACLResponsePtrOutput() NetworkACLResponsePtrOutput
	ToNetworkACLResponsePtrOutputWithContext(context.Context) NetworkACLResponsePtrOutput
}

type networkACLResponsePtrType NetworkACLResponseArgs

func NetworkACLResponsePtr(v *NetworkACLResponseArgs) NetworkACLResponsePtrInput {
	return (*networkACLResponsePtrType)(v)
}

func (*networkACLResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkACLResponse)(nil)).Elem()
}

func (i *networkACLResponsePtrType) ToNetworkACLResponsePtrOutput() NetworkACLResponsePtrOutput {
	return i.ToNetworkACLResponsePtrOutputWithContext(context.Background())
}

func (i *networkACLResponsePtrType) ToNetworkACLResponsePtrOutputWithContext(ctx context.Context) NetworkACLResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLResponsePtrOutput)
}

type NetworkACLResponseOutput struct{ *pulumi.OutputState }

func (NetworkACLResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkACLResponse)(nil)).Elem()
}

func (o NetworkACLResponseOutput) ToNetworkACLResponseOutput() NetworkACLResponseOutput {
	return o
}

func (o NetworkACLResponseOutput) ToNetworkACLResponseOutputWithContext(ctx context.Context) NetworkACLResponseOutput {
	return o
}

func (o NetworkACLResponseOutput) ToNetworkACLResponsePtrOutput() NetworkACLResponsePtrOutput {
	return o.ToNetworkACLResponsePtrOutputWithContext(context.Background())
}

func (o NetworkACLResponseOutput) ToNetworkACLResponsePtrOutputWithContext(ctx context.Context) NetworkACLResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkACLResponse) *NetworkACLResponse {
		return &v
	}).(NetworkACLResponsePtrOutput)
}

func (o NetworkACLResponseOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACLResponse) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o NetworkACLResponseOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACLResponse) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

type NetworkACLResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkACLResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkACLResponse)(nil)).Elem()
}

func (o NetworkACLResponsePtrOutput) ToNetworkACLResponsePtrOutput() NetworkACLResponsePtrOutput {
	return o
}

func (o NetworkACLResponsePtrOutput) ToNetworkACLResponsePtrOutputWithContext(ctx context.Context) NetworkACLResponsePtrOutput {
	return o
}

func (o NetworkACLResponsePtrOutput) Elem() NetworkACLResponseOutput {
	return o.ApplyT(func(v *NetworkACLResponse) NetworkACLResponse {
		if v != nil {
			return *v
		}
		var ret NetworkACLResponse
		return ret
	}).(NetworkACLResponseOutput)
}

func (o NetworkACLResponsePtrOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACLResponse) []string {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(pulumi.StringArrayOutput)
}

func (o NetworkACLResponsePtrOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACLResponse) []string {
		if v == nil {
			return nil
		}
		return v.Deny
	}).(pulumi.StringArrayOutput)
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointACL struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
	Name  string   `pulumi:"name"`
}





type PrivateEndpointACLInput interface {
	pulumi.Input

	ToPrivateEndpointACLOutput() PrivateEndpointACLOutput
	ToPrivateEndpointACLOutputWithContext(context.Context) PrivateEndpointACLOutput
}

type PrivateEndpointACLArgs struct {
	Allow pulumi.StringArrayInput `pulumi:"allow"`
	Deny  pulumi.StringArrayInput `pulumi:"deny"`
	Name  pulumi.StringInput      `pulumi:"name"`
}

func (PrivateEndpointACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointACL)(nil)).Elem()
}

func (i PrivateEndpointACLArgs) ToPrivateEndpointACLOutput() PrivateEndpointACLOutput {
	return i.ToPrivateEndpointACLOutputWithContext(context.Background())
}

func (i PrivateEndpointACLArgs) ToPrivateEndpointACLOutputWithContext(ctx context.Context) PrivateEndpointACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointACLOutput)
}





type PrivateEndpointACLArrayInput interface {
	pulumi.Input

	ToPrivateEndpointACLArrayOutput() PrivateEndpointACLArrayOutput
	ToPrivateEndpointACLArrayOutputWithContext(context.Context) PrivateEndpointACLArrayOutput
}

type PrivateEndpointACLArray []PrivateEndpointACLInput

func (PrivateEndpointACLArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointACL)(nil)).Elem()
}

func (i PrivateEndpointACLArray) ToPrivateEndpointACLArrayOutput() PrivateEndpointACLArrayOutput {
	return i.ToPrivateEndpointACLArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointACLArray) ToPrivateEndpointACLArrayOutputWithContext(ctx context.Context) PrivateEndpointACLArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointACLArrayOutput)
}

type PrivateEndpointACLOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointACL)(nil)).Elem()
}

func (o PrivateEndpointACLOutput) ToPrivateEndpointACLOutput() PrivateEndpointACLOutput {
	return o
}

func (o PrivateEndpointACLOutput) ToPrivateEndpointACLOutputWithContext(ctx context.Context) PrivateEndpointACLOutput {
	return o
}

func (o PrivateEndpointACLOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACL) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACL) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointACL) string { return v.Name }).(pulumi.StringOutput)
}

type PrivateEndpointACLArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointACL)(nil)).Elem()
}

func (o PrivateEndpointACLArrayOutput) ToPrivateEndpointACLArrayOutput() PrivateEndpointACLArrayOutput {
	return o
}

func (o PrivateEndpointACLArrayOutput) ToPrivateEndpointACLArrayOutputWithContext(ctx context.Context) PrivateEndpointACLArrayOutput {
	return o
}

func (o PrivateEndpointACLArrayOutput) Index(i pulumi.IntInput) PrivateEndpointACLOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointACL {
		return vs[0].([]PrivateEndpointACL)[vs[1].(int)]
	}).(PrivateEndpointACLOutput)
}

type PrivateEndpointACLResponse struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
	Name  string   `pulumi:"name"`
}





type PrivateEndpointACLResponseInput interface {
	pulumi.Input

	ToPrivateEndpointACLResponseOutput() PrivateEndpointACLResponseOutput
	ToPrivateEndpointACLResponseOutputWithContext(context.Context) PrivateEndpointACLResponseOutput
}

type PrivateEndpointACLResponseArgs struct {
	Allow pulumi.StringArrayInput `pulumi:"allow"`
	Deny  pulumi.StringArrayInput `pulumi:"deny"`
	Name  pulumi.StringInput      `pulumi:"name"`
}

func (PrivateEndpointACLResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointACLResponse)(nil)).Elem()
}

func (i PrivateEndpointACLResponseArgs) ToPrivateEndpointACLResponseOutput() PrivateEndpointACLResponseOutput {
	return i.ToPrivateEndpointACLResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointACLResponseArgs) ToPrivateEndpointACLResponseOutputWithContext(ctx context.Context) PrivateEndpointACLResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointACLResponseOutput)
}





type PrivateEndpointACLResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointACLResponseArrayOutput() PrivateEndpointACLResponseArrayOutput
	ToPrivateEndpointACLResponseArrayOutputWithContext(context.Context) PrivateEndpointACLResponseArrayOutput
}

type PrivateEndpointACLResponseArray []PrivateEndpointACLResponseInput

func (PrivateEndpointACLResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointACLResponse)(nil)).Elem()
}

func (i PrivateEndpointACLResponseArray) ToPrivateEndpointACLResponseArrayOutput() PrivateEndpointACLResponseArrayOutput {
	return i.ToPrivateEndpointACLResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointACLResponseArray) ToPrivateEndpointACLResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointACLResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointACLResponseArrayOutput)
}

type PrivateEndpointACLResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointACLResponse)(nil)).Elem()
}

func (o PrivateEndpointACLResponseOutput) ToPrivateEndpointACLResponseOutput() PrivateEndpointACLResponseOutput {
	return o
}

func (o PrivateEndpointACLResponseOutput) ToPrivateEndpointACLResponseOutputWithContext(ctx context.Context) PrivateEndpointACLResponseOutput {
	return o
}

func (o PrivateEndpointACLResponseOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACLResponse) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLResponseOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACLResponse) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointACLResponse) string { return v.Name }).(pulumi.StringOutput)
}

type PrivateEndpointACLResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointACLResponse)(nil)).Elem()
}

func (o PrivateEndpointACLResponseArrayOutput) ToPrivateEndpointACLResponseArrayOutput() PrivateEndpointACLResponseArrayOutput {
	return o
}

func (o PrivateEndpointACLResponseArrayOutput) ToPrivateEndpointACLResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointACLResponseArrayOutput {
	return o
}

func (o PrivateEndpointACLResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointACLResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointACLResponse {
		return vs[0].([]PrivateEndpointACLResponse)[vs[1].(int)]
	}).(PrivateEndpointACLResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                     `pulumi:"id"`
	Name                              string                                     `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                         `pulumi:"systemData"`
	Type                              string                                     `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                                `pulumi:"id"`
	Name                              pulumi.StringInput                                `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseInput                           `pulumi:"systemData"`
	Type                              pulumi.StringInput                                `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ResourceSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   string  `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     string  `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuResponseInput interface {
	pulumi.Input

	ToResourceSkuResponseOutput() ResourceSkuResponseOutput
	ToResourceSkuResponseOutputWithContext(context.Context) ResourceSkuResponseOutput
}

type ResourceSkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringInput    `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringInput    `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return i.ToResourceSkuResponseOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput)
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput).ToResourceSkuResponsePtrOutputWithContext(ctx)
}









type ResourceSkuResponsePtrInput interface {
	pulumi.Input

	ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput
	ToResourceSkuResponsePtrOutputWithContext(context.Context) ResourceSkuResponsePtrOutput
}

type resourceSkuResponsePtrType ResourceSkuResponseArgs

func ResourceSkuResponsePtr(v *ResourceSkuResponseArgs) ResourceSkuResponsePtrInput {
	return (*resourceSkuResponsePtrType)(v)
}

func (*resourceSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponsePtrOutput)
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSkuResponse) *ResourceSkuResponse {
		return &v
	}).(ResourceSkuResponsePtrOutput)
}

func (o ResourceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Size }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Size
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourceResponse struct {
	GroupId               string             `pulumi:"groupId"`
	Id                    string             `pulumi:"id"`
	Name                  string             `pulumi:"name"`
	PrivateLinkResourceId string             `pulumi:"privateLinkResourceId"`
	ProvisioningState     string             `pulumi:"provisioningState"`
	RequestMessage        *string            `pulumi:"requestMessage"`
	Status                string             `pulumi:"status"`
	SystemData            SystemDataResponse `pulumi:"systemData"`
	Type                  string             `pulumi:"type"`
}





type SharedPrivateLinkResourceResponseInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput
	ToSharedPrivateLinkResourceResponseOutputWithContext(context.Context) SharedPrivateLinkResourceResponseOutput
}

type SharedPrivateLinkResourceResponseArgs struct {
	GroupId               pulumi.StringInput      `pulumi:"groupId"`
	Id                    pulumi.StringInput      `pulumi:"id"`
	Name                  pulumi.StringInput      `pulumi:"name"`
	PrivateLinkResourceId pulumi.StringInput      `pulumi:"privateLinkResourceId"`
	ProvisioningState     pulumi.StringInput      `pulumi:"provisioningState"`
	RequestMessage        pulumi.StringPtrInput   `pulumi:"requestMessage"`
	Status                pulumi.StringInput      `pulumi:"status"`
	SystemData            SystemDataResponseInput `pulumi:"systemData"`
	Type                  pulumi.StringInput      `pulumi:"type"`
}

func (SharedPrivateLinkResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (i SharedPrivateLinkResourceResponseArgs) ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput {
	return i.ToSharedPrivateLinkResourceResponseOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceResponseArgs) ToSharedPrivateLinkResourceResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceResponseOutput)
}





type SharedPrivateLinkResourceResponseArrayInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput
	ToSharedPrivateLinkResourceResponseArrayOutputWithContext(context.Context) SharedPrivateLinkResourceResponseArrayOutput
}

type SharedPrivateLinkResourceResponseArray []SharedPrivateLinkResourceResponseInput

func (SharedPrivateLinkResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (i SharedPrivateLinkResourceResponseArray) ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput {
	return i.ToSharedPrivateLinkResourceResponseArrayOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceResponseArray) ToSharedPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceResponseArrayOutput)
}

type SharedPrivateLinkResourceResponseOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourceResponseOutput) ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseOutput) ToSharedPrivateLinkResourceResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) PrivateLinkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.PrivateLinkResourceId }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SharedPrivateLinkResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourceResponseArrayOutput) ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseArrayOutput) ToSharedPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseArrayOutput) Index(i pulumi.IntInput) SharedPrivateLinkResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedPrivateLinkResourceResponse {
		return vs[0].([]SharedPrivateLinkResourceResponse)[vs[1].(int)]
	}).(SharedPrivateLinkResourceResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type UpstreamAuthSettings struct {
	ManagedIdentity *ManagedIdentitySettings `pulumi:"managedIdentity"`
	Type            *string                  `pulumi:"type"`
}





type UpstreamAuthSettingsInput interface {
	pulumi.Input

	ToUpstreamAuthSettingsOutput() UpstreamAuthSettingsOutput
	ToUpstreamAuthSettingsOutputWithContext(context.Context) UpstreamAuthSettingsOutput
}

type UpstreamAuthSettingsArgs struct {
	ManagedIdentity ManagedIdentitySettingsPtrInput `pulumi:"managedIdentity"`
	Type            pulumi.StringPtrInput           `pulumi:"type"`
}

func (UpstreamAuthSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamAuthSettings)(nil)).Elem()
}

func (i UpstreamAuthSettingsArgs) ToUpstreamAuthSettingsOutput() UpstreamAuthSettingsOutput {
	return i.ToUpstreamAuthSettingsOutputWithContext(context.Background())
}

func (i UpstreamAuthSettingsArgs) ToUpstreamAuthSettingsOutputWithContext(ctx context.Context) UpstreamAuthSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamAuthSettingsOutput)
}

func (i UpstreamAuthSettingsArgs) ToUpstreamAuthSettingsPtrOutput() UpstreamAuthSettingsPtrOutput {
	return i.ToUpstreamAuthSettingsPtrOutputWithContext(context.Background())
}

func (i UpstreamAuthSettingsArgs) ToUpstreamAuthSettingsPtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamAuthSettingsOutput).ToUpstreamAuthSettingsPtrOutputWithContext(ctx)
}









type UpstreamAuthSettingsPtrInput interface {
	pulumi.Input

	ToUpstreamAuthSettingsPtrOutput() UpstreamAuthSettingsPtrOutput
	ToUpstreamAuthSettingsPtrOutputWithContext(context.Context) UpstreamAuthSettingsPtrOutput
}

type upstreamAuthSettingsPtrType UpstreamAuthSettingsArgs

func UpstreamAuthSettingsPtr(v *UpstreamAuthSettingsArgs) UpstreamAuthSettingsPtrInput {
	return (*upstreamAuthSettingsPtrType)(v)
}

func (*upstreamAuthSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UpstreamAuthSettings)(nil)).Elem()
}

func (i *upstreamAuthSettingsPtrType) ToUpstreamAuthSettingsPtrOutput() UpstreamAuthSettingsPtrOutput {
	return i.ToUpstreamAuthSettingsPtrOutputWithContext(context.Background())
}

func (i *upstreamAuthSettingsPtrType) ToUpstreamAuthSettingsPtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamAuthSettingsPtrOutput)
}

type UpstreamAuthSettingsOutput struct{ *pulumi.OutputState }

func (UpstreamAuthSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamAuthSettings)(nil)).Elem()
}

func (o UpstreamAuthSettingsOutput) ToUpstreamAuthSettingsOutput() UpstreamAuthSettingsOutput {
	return o
}

func (o UpstreamAuthSettingsOutput) ToUpstreamAuthSettingsOutputWithContext(ctx context.Context) UpstreamAuthSettingsOutput {
	return o
}

func (o UpstreamAuthSettingsOutput) ToUpstreamAuthSettingsPtrOutput() UpstreamAuthSettingsPtrOutput {
	return o.ToUpstreamAuthSettingsPtrOutputWithContext(context.Background())
}

func (o UpstreamAuthSettingsOutput) ToUpstreamAuthSettingsPtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpstreamAuthSettings) *UpstreamAuthSettings {
		return &v
	}).(UpstreamAuthSettingsPtrOutput)
}

func (o UpstreamAuthSettingsOutput) ManagedIdentity() ManagedIdentitySettingsPtrOutput {
	return o.ApplyT(func(v UpstreamAuthSettings) *ManagedIdentitySettings { return v.ManagedIdentity }).(ManagedIdentitySettingsPtrOutput)
}

func (o UpstreamAuthSettingsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpstreamAuthSettings) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type UpstreamAuthSettingsPtrOutput struct{ *pulumi.OutputState }

func (UpstreamAuthSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpstreamAuthSettings)(nil)).Elem()
}

func (o UpstreamAuthSettingsPtrOutput) ToUpstreamAuthSettingsPtrOutput() UpstreamAuthSettingsPtrOutput {
	return o
}

func (o UpstreamAuthSettingsPtrOutput) ToUpstreamAuthSettingsPtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsPtrOutput {
	return o
}

func (o UpstreamAuthSettingsPtrOutput) Elem() UpstreamAuthSettingsOutput {
	return o.ApplyT(func(v *UpstreamAuthSettings) UpstreamAuthSettings {
		if v != nil {
			return *v
		}
		var ret UpstreamAuthSettings
		return ret
	}).(UpstreamAuthSettingsOutput)
}

func (o UpstreamAuthSettingsPtrOutput) ManagedIdentity() ManagedIdentitySettingsPtrOutput {
	return o.ApplyT(func(v *UpstreamAuthSettings) *ManagedIdentitySettings {
		if v == nil {
			return nil
		}
		return v.ManagedIdentity
	}).(ManagedIdentitySettingsPtrOutput)
}

func (o UpstreamAuthSettingsPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpstreamAuthSettings) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type UpstreamAuthSettingsResponse struct {
	ManagedIdentity *ManagedIdentitySettingsResponse `pulumi:"managedIdentity"`
	Type            *string                          `pulumi:"type"`
}





type UpstreamAuthSettingsResponseInput interface {
	pulumi.Input

	ToUpstreamAuthSettingsResponseOutput() UpstreamAuthSettingsResponseOutput
	ToUpstreamAuthSettingsResponseOutputWithContext(context.Context) UpstreamAuthSettingsResponseOutput
}

type UpstreamAuthSettingsResponseArgs struct {
	ManagedIdentity ManagedIdentitySettingsResponsePtrInput `pulumi:"managedIdentity"`
	Type            pulumi.StringPtrInput                   `pulumi:"type"`
}

func (UpstreamAuthSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamAuthSettingsResponse)(nil)).Elem()
}

func (i UpstreamAuthSettingsResponseArgs) ToUpstreamAuthSettingsResponseOutput() UpstreamAuthSettingsResponseOutput {
	return i.ToUpstreamAuthSettingsResponseOutputWithContext(context.Background())
}

func (i UpstreamAuthSettingsResponseArgs) ToUpstreamAuthSettingsResponseOutputWithContext(ctx context.Context) UpstreamAuthSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamAuthSettingsResponseOutput)
}

func (i UpstreamAuthSettingsResponseArgs) ToUpstreamAuthSettingsResponsePtrOutput() UpstreamAuthSettingsResponsePtrOutput {
	return i.ToUpstreamAuthSettingsResponsePtrOutputWithContext(context.Background())
}

func (i UpstreamAuthSettingsResponseArgs) ToUpstreamAuthSettingsResponsePtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamAuthSettingsResponseOutput).ToUpstreamAuthSettingsResponsePtrOutputWithContext(ctx)
}









type UpstreamAuthSettingsResponsePtrInput interface {
	pulumi.Input

	ToUpstreamAuthSettingsResponsePtrOutput() UpstreamAuthSettingsResponsePtrOutput
	ToUpstreamAuthSettingsResponsePtrOutputWithContext(context.Context) UpstreamAuthSettingsResponsePtrOutput
}

type upstreamAuthSettingsResponsePtrType UpstreamAuthSettingsResponseArgs

func UpstreamAuthSettingsResponsePtr(v *UpstreamAuthSettingsResponseArgs) UpstreamAuthSettingsResponsePtrInput {
	return (*upstreamAuthSettingsResponsePtrType)(v)
}

func (*upstreamAuthSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UpstreamAuthSettingsResponse)(nil)).Elem()
}

func (i *upstreamAuthSettingsResponsePtrType) ToUpstreamAuthSettingsResponsePtrOutput() UpstreamAuthSettingsResponsePtrOutput {
	return i.ToUpstreamAuthSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *upstreamAuthSettingsResponsePtrType) ToUpstreamAuthSettingsResponsePtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamAuthSettingsResponsePtrOutput)
}

type UpstreamAuthSettingsResponseOutput struct{ *pulumi.OutputState }

func (UpstreamAuthSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamAuthSettingsResponse)(nil)).Elem()
}

func (o UpstreamAuthSettingsResponseOutput) ToUpstreamAuthSettingsResponseOutput() UpstreamAuthSettingsResponseOutput {
	return o
}

func (o UpstreamAuthSettingsResponseOutput) ToUpstreamAuthSettingsResponseOutputWithContext(ctx context.Context) UpstreamAuthSettingsResponseOutput {
	return o
}

func (o UpstreamAuthSettingsResponseOutput) ToUpstreamAuthSettingsResponsePtrOutput() UpstreamAuthSettingsResponsePtrOutput {
	return o.ToUpstreamAuthSettingsResponsePtrOutputWithContext(context.Background())
}

func (o UpstreamAuthSettingsResponseOutput) ToUpstreamAuthSettingsResponsePtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpstreamAuthSettingsResponse) *UpstreamAuthSettingsResponse {
		return &v
	}).(UpstreamAuthSettingsResponsePtrOutput)
}

func (o UpstreamAuthSettingsResponseOutput) ManagedIdentity() ManagedIdentitySettingsResponsePtrOutput {
	return o.ApplyT(func(v UpstreamAuthSettingsResponse) *ManagedIdentitySettingsResponse { return v.ManagedIdentity }).(ManagedIdentitySettingsResponsePtrOutput)
}

func (o UpstreamAuthSettingsResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpstreamAuthSettingsResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type UpstreamAuthSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (UpstreamAuthSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpstreamAuthSettingsResponse)(nil)).Elem()
}

func (o UpstreamAuthSettingsResponsePtrOutput) ToUpstreamAuthSettingsResponsePtrOutput() UpstreamAuthSettingsResponsePtrOutput {
	return o
}

func (o UpstreamAuthSettingsResponsePtrOutput) ToUpstreamAuthSettingsResponsePtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsResponsePtrOutput {
	return o
}

func (o UpstreamAuthSettingsResponsePtrOutput) Elem() UpstreamAuthSettingsResponseOutput {
	return o.ApplyT(func(v *UpstreamAuthSettingsResponse) UpstreamAuthSettingsResponse {
		if v != nil {
			return *v
		}
		var ret UpstreamAuthSettingsResponse
		return ret
	}).(UpstreamAuthSettingsResponseOutput)
}

func (o UpstreamAuthSettingsResponsePtrOutput) ManagedIdentity() ManagedIdentitySettingsResponsePtrOutput {
	return o.ApplyT(func(v *UpstreamAuthSettingsResponse) *ManagedIdentitySettingsResponse {
		if v == nil {
			return nil
		}
		return v.ManagedIdentity
	}).(ManagedIdentitySettingsResponsePtrOutput)
}

func (o UpstreamAuthSettingsResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpstreamAuthSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityPropertyResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type UserAssignedIdentityPropertyResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityPropertyResponseOutput() UserAssignedIdentityPropertyResponseOutput
	ToUserAssignedIdentityPropertyResponseOutputWithContext(context.Context) UserAssignedIdentityPropertyResponseOutput
}

type UserAssignedIdentityPropertyResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (UserAssignedIdentityPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityPropertyResponse)(nil)).Elem()
}

func (i UserAssignedIdentityPropertyResponseArgs) ToUserAssignedIdentityPropertyResponseOutput() UserAssignedIdentityPropertyResponseOutput {
	return i.ToUserAssignedIdentityPropertyResponseOutputWithContext(context.Background())
}

func (i UserAssignedIdentityPropertyResponseArgs) ToUserAssignedIdentityPropertyResponseOutputWithContext(ctx context.Context) UserAssignedIdentityPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityPropertyResponseOutput)
}





type UserAssignedIdentityPropertyResponseMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityPropertyResponseMapOutput() UserAssignedIdentityPropertyResponseMapOutput
	ToUserAssignedIdentityPropertyResponseMapOutputWithContext(context.Context) UserAssignedIdentityPropertyResponseMapOutput
}

type UserAssignedIdentityPropertyResponseMap map[string]UserAssignedIdentityPropertyResponseInput

func (UserAssignedIdentityPropertyResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityPropertyResponse)(nil)).Elem()
}

func (i UserAssignedIdentityPropertyResponseMap) ToUserAssignedIdentityPropertyResponseMapOutput() UserAssignedIdentityPropertyResponseMapOutput {
	return i.ToUserAssignedIdentityPropertyResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityPropertyResponseMap) ToUserAssignedIdentityPropertyResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityPropertyResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityPropertyResponseMapOutput)
}

type UserAssignedIdentityPropertyResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityPropertyResponse)(nil)).Elem()
}

func (o UserAssignedIdentityPropertyResponseOutput) ToUserAssignedIdentityPropertyResponseOutput() UserAssignedIdentityPropertyResponseOutput {
	return o
}

func (o UserAssignedIdentityPropertyResponseOutput) ToUserAssignedIdentityPropertyResponseOutputWithContext(ctx context.Context) UserAssignedIdentityPropertyResponseOutput {
	return o
}

func (o UserAssignedIdentityPropertyResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityPropertyResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityPropertyResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityPropertyResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityPropertyResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertyResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityPropertyResponse)(nil)).Elem()
}

func (o UserAssignedIdentityPropertyResponseMapOutput) ToUserAssignedIdentityPropertyResponseMapOutput() UserAssignedIdentityPropertyResponseMapOutput {
	return o
}

func (o UserAssignedIdentityPropertyResponseMapOutput) ToUserAssignedIdentityPropertyResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityPropertyResponseMapOutput {
	return o
}

func (o UserAssignedIdentityPropertyResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityPropertyResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityPropertyResponse {
		return vs[0].(map[string]UserAssignedIdentityPropertyResponse)[vs[1].(string)]
	}).(UserAssignedIdentityPropertyResponseOutput)
}

type WebPubSubNetworkACLs struct {
	DefaultAction    *string              `pulumi:"defaultAction"`
	PrivateEndpoints []PrivateEndpointACL `pulumi:"privateEndpoints"`
	PublicNetwork    *NetworkACL          `pulumi:"publicNetwork"`
}





type WebPubSubNetworkACLsInput interface {
	pulumi.Input

	ToWebPubSubNetworkACLsOutput() WebPubSubNetworkACLsOutput
	ToWebPubSubNetworkACLsOutputWithContext(context.Context) WebPubSubNetworkACLsOutput
}

type WebPubSubNetworkACLsArgs struct {
	DefaultAction    pulumi.StringPtrInput        `pulumi:"defaultAction"`
	PrivateEndpoints PrivateEndpointACLArrayInput `pulumi:"privateEndpoints"`
	PublicNetwork    NetworkACLPtrInput           `pulumi:"publicNetwork"`
}

func (WebPubSubNetworkACLsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubNetworkACLs)(nil)).Elem()
}

func (i WebPubSubNetworkACLsArgs) ToWebPubSubNetworkACLsOutput() WebPubSubNetworkACLsOutput {
	return i.ToWebPubSubNetworkACLsOutputWithContext(context.Background())
}

func (i WebPubSubNetworkACLsArgs) ToWebPubSubNetworkACLsOutputWithContext(ctx context.Context) WebPubSubNetworkACLsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubNetworkACLsOutput)
}

func (i WebPubSubNetworkACLsArgs) ToWebPubSubNetworkACLsPtrOutput() WebPubSubNetworkACLsPtrOutput {
	return i.ToWebPubSubNetworkACLsPtrOutputWithContext(context.Background())
}

func (i WebPubSubNetworkACLsArgs) ToWebPubSubNetworkACLsPtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubNetworkACLsOutput).ToWebPubSubNetworkACLsPtrOutputWithContext(ctx)
}









type WebPubSubNetworkACLsPtrInput interface {
	pulumi.Input

	ToWebPubSubNetworkACLsPtrOutput() WebPubSubNetworkACLsPtrOutput
	ToWebPubSubNetworkACLsPtrOutputWithContext(context.Context) WebPubSubNetworkACLsPtrOutput
}

type webPubSubNetworkACLsPtrType WebPubSubNetworkACLsArgs

func WebPubSubNetworkACLsPtr(v *WebPubSubNetworkACLsArgs) WebPubSubNetworkACLsPtrInput {
	return (*webPubSubNetworkACLsPtrType)(v)
}

func (*webPubSubNetworkACLsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubNetworkACLs)(nil)).Elem()
}

func (i *webPubSubNetworkACLsPtrType) ToWebPubSubNetworkACLsPtrOutput() WebPubSubNetworkACLsPtrOutput {
	return i.ToWebPubSubNetworkACLsPtrOutputWithContext(context.Background())
}

func (i *webPubSubNetworkACLsPtrType) ToWebPubSubNetworkACLsPtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubNetworkACLsPtrOutput)
}

type WebPubSubNetworkACLsOutput struct{ *pulumi.OutputState }

func (WebPubSubNetworkACLsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubNetworkACLs)(nil)).Elem()
}

func (o WebPubSubNetworkACLsOutput) ToWebPubSubNetworkACLsOutput() WebPubSubNetworkACLsOutput {
	return o
}

func (o WebPubSubNetworkACLsOutput) ToWebPubSubNetworkACLsOutputWithContext(ctx context.Context) WebPubSubNetworkACLsOutput {
	return o
}

func (o WebPubSubNetworkACLsOutput) ToWebPubSubNetworkACLsPtrOutput() WebPubSubNetworkACLsPtrOutput {
	return o.ToWebPubSubNetworkACLsPtrOutputWithContext(context.Background())
}

func (o WebPubSubNetworkACLsOutput) ToWebPubSubNetworkACLsPtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebPubSubNetworkACLs) *WebPubSubNetworkACLs {
		return &v
	}).(WebPubSubNetworkACLsPtrOutput)
}

func (o WebPubSubNetworkACLsOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLs) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o WebPubSubNetworkACLsOutput) PrivateEndpoints() PrivateEndpointACLArrayOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLs) []PrivateEndpointACL { return v.PrivateEndpoints }).(PrivateEndpointACLArrayOutput)
}

func (o WebPubSubNetworkACLsOutput) PublicNetwork() NetworkACLPtrOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLs) *NetworkACL { return v.PublicNetwork }).(NetworkACLPtrOutput)
}

type WebPubSubNetworkACLsPtrOutput struct{ *pulumi.OutputState }

func (WebPubSubNetworkACLsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubNetworkACLs)(nil)).Elem()
}

func (o WebPubSubNetworkACLsPtrOutput) ToWebPubSubNetworkACLsPtrOutput() WebPubSubNetworkACLsPtrOutput {
	return o
}

func (o WebPubSubNetworkACLsPtrOutput) ToWebPubSubNetworkACLsPtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsPtrOutput {
	return o
}

func (o WebPubSubNetworkACLsPtrOutput) Elem() WebPubSubNetworkACLsOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLs) WebPubSubNetworkACLs {
		if v != nil {
			return *v
		}
		var ret WebPubSubNetworkACLs
		return ret
	}).(WebPubSubNetworkACLsOutput)
}

func (o WebPubSubNetworkACLsPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLs) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o WebPubSubNetworkACLsPtrOutput) PrivateEndpoints() PrivateEndpointACLArrayOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLs) []PrivateEndpointACL {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoints
	}).(PrivateEndpointACLArrayOutput)
}

func (o WebPubSubNetworkACLsPtrOutput) PublicNetwork() NetworkACLPtrOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLs) *NetworkACL {
		if v == nil {
			return nil
		}
		return v.PublicNetwork
	}).(NetworkACLPtrOutput)
}

type WebPubSubNetworkACLsResponse struct {
	DefaultAction    *string                      `pulumi:"defaultAction"`
	PrivateEndpoints []PrivateEndpointACLResponse `pulumi:"privateEndpoints"`
	PublicNetwork    *NetworkACLResponse          `pulumi:"publicNetwork"`
}





type WebPubSubNetworkACLsResponseInput interface {
	pulumi.Input

	ToWebPubSubNetworkACLsResponseOutput() WebPubSubNetworkACLsResponseOutput
	ToWebPubSubNetworkACLsResponseOutputWithContext(context.Context) WebPubSubNetworkACLsResponseOutput
}

type WebPubSubNetworkACLsResponseArgs struct {
	DefaultAction    pulumi.StringPtrInput                `pulumi:"defaultAction"`
	PrivateEndpoints PrivateEndpointACLResponseArrayInput `pulumi:"privateEndpoints"`
	PublicNetwork    NetworkACLResponsePtrInput           `pulumi:"publicNetwork"`
}

func (WebPubSubNetworkACLsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubNetworkACLsResponse)(nil)).Elem()
}

func (i WebPubSubNetworkACLsResponseArgs) ToWebPubSubNetworkACLsResponseOutput() WebPubSubNetworkACLsResponseOutput {
	return i.ToWebPubSubNetworkACLsResponseOutputWithContext(context.Background())
}

func (i WebPubSubNetworkACLsResponseArgs) ToWebPubSubNetworkACLsResponseOutputWithContext(ctx context.Context) WebPubSubNetworkACLsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubNetworkACLsResponseOutput)
}

func (i WebPubSubNetworkACLsResponseArgs) ToWebPubSubNetworkACLsResponsePtrOutput() WebPubSubNetworkACLsResponsePtrOutput {
	return i.ToWebPubSubNetworkACLsResponsePtrOutputWithContext(context.Background())
}

func (i WebPubSubNetworkACLsResponseArgs) ToWebPubSubNetworkACLsResponsePtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubNetworkACLsResponseOutput).ToWebPubSubNetworkACLsResponsePtrOutputWithContext(ctx)
}









type WebPubSubNetworkACLsResponsePtrInput interface {
	pulumi.Input

	ToWebPubSubNetworkACLsResponsePtrOutput() WebPubSubNetworkACLsResponsePtrOutput
	ToWebPubSubNetworkACLsResponsePtrOutputWithContext(context.Context) WebPubSubNetworkACLsResponsePtrOutput
}

type webPubSubNetworkACLsResponsePtrType WebPubSubNetworkACLsResponseArgs

func WebPubSubNetworkACLsResponsePtr(v *WebPubSubNetworkACLsResponseArgs) WebPubSubNetworkACLsResponsePtrInput {
	return (*webPubSubNetworkACLsResponsePtrType)(v)
}

func (*webPubSubNetworkACLsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubNetworkACLsResponse)(nil)).Elem()
}

func (i *webPubSubNetworkACLsResponsePtrType) ToWebPubSubNetworkACLsResponsePtrOutput() WebPubSubNetworkACLsResponsePtrOutput {
	return i.ToWebPubSubNetworkACLsResponsePtrOutputWithContext(context.Background())
}

func (i *webPubSubNetworkACLsResponsePtrType) ToWebPubSubNetworkACLsResponsePtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubNetworkACLsResponsePtrOutput)
}

type WebPubSubNetworkACLsResponseOutput struct{ *pulumi.OutputState }

func (WebPubSubNetworkACLsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubNetworkACLsResponse)(nil)).Elem()
}

func (o WebPubSubNetworkACLsResponseOutput) ToWebPubSubNetworkACLsResponseOutput() WebPubSubNetworkACLsResponseOutput {
	return o
}

func (o WebPubSubNetworkACLsResponseOutput) ToWebPubSubNetworkACLsResponseOutputWithContext(ctx context.Context) WebPubSubNetworkACLsResponseOutput {
	return o
}

func (o WebPubSubNetworkACLsResponseOutput) ToWebPubSubNetworkACLsResponsePtrOutput() WebPubSubNetworkACLsResponsePtrOutput {
	return o.ToWebPubSubNetworkACLsResponsePtrOutputWithContext(context.Background())
}

func (o WebPubSubNetworkACLsResponseOutput) ToWebPubSubNetworkACLsResponsePtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebPubSubNetworkACLsResponse) *WebPubSubNetworkACLsResponse {
		return &v
	}).(WebPubSubNetworkACLsResponsePtrOutput)
}

func (o WebPubSubNetworkACLsResponseOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLsResponse) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o WebPubSubNetworkACLsResponseOutput) PrivateEndpoints() PrivateEndpointACLResponseArrayOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLsResponse) []PrivateEndpointACLResponse { return v.PrivateEndpoints }).(PrivateEndpointACLResponseArrayOutput)
}

func (o WebPubSubNetworkACLsResponseOutput) PublicNetwork() NetworkACLResponsePtrOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLsResponse) *NetworkACLResponse { return v.PublicNetwork }).(NetworkACLResponsePtrOutput)
}

type WebPubSubNetworkACLsResponsePtrOutput struct{ *pulumi.OutputState }

func (WebPubSubNetworkACLsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubNetworkACLsResponse)(nil)).Elem()
}

func (o WebPubSubNetworkACLsResponsePtrOutput) ToWebPubSubNetworkACLsResponsePtrOutput() WebPubSubNetworkACLsResponsePtrOutput {
	return o
}

func (o WebPubSubNetworkACLsResponsePtrOutput) ToWebPubSubNetworkACLsResponsePtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsResponsePtrOutput {
	return o
}

func (o WebPubSubNetworkACLsResponsePtrOutput) Elem() WebPubSubNetworkACLsResponseOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLsResponse) WebPubSubNetworkACLsResponse {
		if v != nil {
			return *v
		}
		var ret WebPubSubNetworkACLsResponse
		return ret
	}).(WebPubSubNetworkACLsResponseOutput)
}

func (o WebPubSubNetworkACLsResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o WebPubSubNetworkACLsResponsePtrOutput) PrivateEndpoints() PrivateEndpointACLResponseArrayOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLsResponse) []PrivateEndpointACLResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoints
	}).(PrivateEndpointACLResponseArrayOutput)
}

func (o WebPubSubNetworkACLsResponsePtrOutput) PublicNetwork() NetworkACLResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLsResponse) *NetworkACLResponse {
		if v == nil {
			return nil
		}
		return v.PublicNetwork
	}).(NetworkACLResponsePtrOutput)
}

type WebPubSubTlsSettings struct {
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
}





type WebPubSubTlsSettingsInput interface {
	pulumi.Input

	ToWebPubSubTlsSettingsOutput() WebPubSubTlsSettingsOutput
	ToWebPubSubTlsSettingsOutputWithContext(context.Context) WebPubSubTlsSettingsOutput
}

type WebPubSubTlsSettingsArgs struct {
	ClientCertEnabled pulumi.BoolPtrInput `pulumi:"clientCertEnabled"`
}

func (WebPubSubTlsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubTlsSettings)(nil)).Elem()
}

func (i WebPubSubTlsSettingsArgs) ToWebPubSubTlsSettingsOutput() WebPubSubTlsSettingsOutput {
	return i.ToWebPubSubTlsSettingsOutputWithContext(context.Background())
}

func (i WebPubSubTlsSettingsArgs) ToWebPubSubTlsSettingsOutputWithContext(ctx context.Context) WebPubSubTlsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubTlsSettingsOutput)
}

func (i WebPubSubTlsSettingsArgs) ToWebPubSubTlsSettingsPtrOutput() WebPubSubTlsSettingsPtrOutput {
	return i.ToWebPubSubTlsSettingsPtrOutputWithContext(context.Background())
}

func (i WebPubSubTlsSettingsArgs) ToWebPubSubTlsSettingsPtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubTlsSettingsOutput).ToWebPubSubTlsSettingsPtrOutputWithContext(ctx)
}









type WebPubSubTlsSettingsPtrInput interface {
	pulumi.Input

	ToWebPubSubTlsSettingsPtrOutput() WebPubSubTlsSettingsPtrOutput
	ToWebPubSubTlsSettingsPtrOutputWithContext(context.Context) WebPubSubTlsSettingsPtrOutput
}

type webPubSubTlsSettingsPtrType WebPubSubTlsSettingsArgs

func WebPubSubTlsSettingsPtr(v *WebPubSubTlsSettingsArgs) WebPubSubTlsSettingsPtrInput {
	return (*webPubSubTlsSettingsPtrType)(v)
}

func (*webPubSubTlsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubTlsSettings)(nil)).Elem()
}

func (i *webPubSubTlsSettingsPtrType) ToWebPubSubTlsSettingsPtrOutput() WebPubSubTlsSettingsPtrOutput {
	return i.ToWebPubSubTlsSettingsPtrOutputWithContext(context.Background())
}

func (i *webPubSubTlsSettingsPtrType) ToWebPubSubTlsSettingsPtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubTlsSettingsPtrOutput)
}

type WebPubSubTlsSettingsOutput struct{ *pulumi.OutputState }

func (WebPubSubTlsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubTlsSettings)(nil)).Elem()
}

func (o WebPubSubTlsSettingsOutput) ToWebPubSubTlsSettingsOutput() WebPubSubTlsSettingsOutput {
	return o
}

func (o WebPubSubTlsSettingsOutput) ToWebPubSubTlsSettingsOutputWithContext(ctx context.Context) WebPubSubTlsSettingsOutput {
	return o
}

func (o WebPubSubTlsSettingsOutput) ToWebPubSubTlsSettingsPtrOutput() WebPubSubTlsSettingsPtrOutput {
	return o.ToWebPubSubTlsSettingsPtrOutputWithContext(context.Background())
}

func (o WebPubSubTlsSettingsOutput) ToWebPubSubTlsSettingsPtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebPubSubTlsSettings) *WebPubSubTlsSettings {
		return &v
	}).(WebPubSubTlsSettingsPtrOutput)
}

func (o WebPubSubTlsSettingsOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebPubSubTlsSettings) *bool { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

type WebPubSubTlsSettingsPtrOutput struct{ *pulumi.OutputState }

func (WebPubSubTlsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubTlsSettings)(nil)).Elem()
}

func (o WebPubSubTlsSettingsPtrOutput) ToWebPubSubTlsSettingsPtrOutput() WebPubSubTlsSettingsPtrOutput {
	return o
}

func (o WebPubSubTlsSettingsPtrOutput) ToWebPubSubTlsSettingsPtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsPtrOutput {
	return o
}

func (o WebPubSubTlsSettingsPtrOutput) Elem() WebPubSubTlsSettingsOutput {
	return o.ApplyT(func(v *WebPubSubTlsSettings) WebPubSubTlsSettings {
		if v != nil {
			return *v
		}
		var ret WebPubSubTlsSettings
		return ret
	}).(WebPubSubTlsSettingsOutput)
}

func (o WebPubSubTlsSettingsPtrOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebPubSubTlsSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ClientCertEnabled
	}).(pulumi.BoolPtrOutput)
}

type WebPubSubTlsSettingsResponse struct {
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
}





type WebPubSubTlsSettingsResponseInput interface {
	pulumi.Input

	ToWebPubSubTlsSettingsResponseOutput() WebPubSubTlsSettingsResponseOutput
	ToWebPubSubTlsSettingsResponseOutputWithContext(context.Context) WebPubSubTlsSettingsResponseOutput
}

type WebPubSubTlsSettingsResponseArgs struct {
	ClientCertEnabled pulumi.BoolPtrInput `pulumi:"clientCertEnabled"`
}

func (WebPubSubTlsSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubTlsSettingsResponse)(nil)).Elem()
}

func (i WebPubSubTlsSettingsResponseArgs) ToWebPubSubTlsSettingsResponseOutput() WebPubSubTlsSettingsResponseOutput {
	return i.ToWebPubSubTlsSettingsResponseOutputWithContext(context.Background())
}

func (i WebPubSubTlsSettingsResponseArgs) ToWebPubSubTlsSettingsResponseOutputWithContext(ctx context.Context) WebPubSubTlsSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubTlsSettingsResponseOutput)
}

func (i WebPubSubTlsSettingsResponseArgs) ToWebPubSubTlsSettingsResponsePtrOutput() WebPubSubTlsSettingsResponsePtrOutput {
	return i.ToWebPubSubTlsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i WebPubSubTlsSettingsResponseArgs) ToWebPubSubTlsSettingsResponsePtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubTlsSettingsResponseOutput).ToWebPubSubTlsSettingsResponsePtrOutputWithContext(ctx)
}









type WebPubSubTlsSettingsResponsePtrInput interface {
	pulumi.Input

	ToWebPubSubTlsSettingsResponsePtrOutput() WebPubSubTlsSettingsResponsePtrOutput
	ToWebPubSubTlsSettingsResponsePtrOutputWithContext(context.Context) WebPubSubTlsSettingsResponsePtrOutput
}

type webPubSubTlsSettingsResponsePtrType WebPubSubTlsSettingsResponseArgs

func WebPubSubTlsSettingsResponsePtr(v *WebPubSubTlsSettingsResponseArgs) WebPubSubTlsSettingsResponsePtrInput {
	return (*webPubSubTlsSettingsResponsePtrType)(v)
}

func (*webPubSubTlsSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubTlsSettingsResponse)(nil)).Elem()
}

func (i *webPubSubTlsSettingsResponsePtrType) ToWebPubSubTlsSettingsResponsePtrOutput() WebPubSubTlsSettingsResponsePtrOutput {
	return i.ToWebPubSubTlsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *webPubSubTlsSettingsResponsePtrType) ToWebPubSubTlsSettingsResponsePtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubTlsSettingsResponsePtrOutput)
}

type WebPubSubTlsSettingsResponseOutput struct{ *pulumi.OutputState }

func (WebPubSubTlsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubTlsSettingsResponse)(nil)).Elem()
}

func (o WebPubSubTlsSettingsResponseOutput) ToWebPubSubTlsSettingsResponseOutput() WebPubSubTlsSettingsResponseOutput {
	return o
}

func (o WebPubSubTlsSettingsResponseOutput) ToWebPubSubTlsSettingsResponseOutputWithContext(ctx context.Context) WebPubSubTlsSettingsResponseOutput {
	return o
}

func (o WebPubSubTlsSettingsResponseOutput) ToWebPubSubTlsSettingsResponsePtrOutput() WebPubSubTlsSettingsResponsePtrOutput {
	return o.ToWebPubSubTlsSettingsResponsePtrOutputWithContext(context.Background())
}

func (o WebPubSubTlsSettingsResponseOutput) ToWebPubSubTlsSettingsResponsePtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebPubSubTlsSettingsResponse) *WebPubSubTlsSettingsResponse {
		return &v
	}).(WebPubSubTlsSettingsResponsePtrOutput)
}

func (o WebPubSubTlsSettingsResponseOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebPubSubTlsSettingsResponse) *bool { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

type WebPubSubTlsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (WebPubSubTlsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubTlsSettingsResponse)(nil)).Elem()
}

func (o WebPubSubTlsSettingsResponsePtrOutput) ToWebPubSubTlsSettingsResponsePtrOutput() WebPubSubTlsSettingsResponsePtrOutput {
	return o
}

func (o WebPubSubTlsSettingsResponsePtrOutput) ToWebPubSubTlsSettingsResponsePtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsResponsePtrOutput {
	return o
}

func (o WebPubSubTlsSettingsResponsePtrOutput) Elem() WebPubSubTlsSettingsResponseOutput {
	return o.ApplyT(func(v *WebPubSubTlsSettingsResponse) WebPubSubTlsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret WebPubSubTlsSettingsResponse
		return ret
	}).(WebPubSubTlsSettingsResponseOutput)
}

func (o WebPubSubTlsSettingsResponsePtrOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebPubSubTlsSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ClientCertEnabled
	}).(pulumi.BoolPtrOutput)
}

type EventHandlerTemplateArrayMap map[string]EventHandlerTemplateArrayInput

func (EventHandlerTemplateArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerTemplateArray)(nil)).Elem()
}

func (i EventHandlerTemplateArrayMap) ToEventHandlerTemplateArrayMapOutput() EventHandlerTemplateArrayMapOutput {
	return i.ToEventHandlerTemplateArrayMapOutputWithContext(context.Background())
}

func (i EventHandlerTemplateArrayMap) ToEventHandlerTemplateArrayMapOutputWithContext(ctx context.Context) EventHandlerTemplateArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerTemplateArrayMapOutput)
}

type EventHandlerTemplateArrayMapOutput struct{ *pulumi.OutputState }

func (EventHandlerTemplateArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventHandlerTemplateArray)(nil)).Elem()
}

func (o EventHandlerTemplateArrayMapOutput) ToEventHandlerTemplateArrayMapOutput() EventHandlerTemplateArrayMapOutput {
	return o
}

func (o EventHandlerTemplateArrayMapOutput) ToEventHandlerTemplateArrayMapOutputWithContext(ctx context.Context) EventHandlerTemplateArrayMapOutput {
	return o
}

func (o EventHandlerTemplateArrayMapOutput) MapIndex(k pulumi.StringInput) EventHandlerTemplateArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventHandlerTemplateArray {
		return vs[0].(map[string]EventHandlerTemplateArray)[vs[1].(string)]
	}).(EventHandlerTemplateArrayOutput)
}





type EventHandlerTemplateArrayMapInput interface {
	pulumi.Input

	ToEventHandlerTemplateArrayMapOutput() EventHandlerTemplateArrayMapOutput
	ToEventHandlerTemplateArrayMapOutputWithContext(context.Context) EventHandlerTemplateArrayMapOutput
}

type EventHandlerTemplateResponseArrayMap map[string]EventHandlerTemplateResponseArrayInput

func (EventHandlerTemplateResponseArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerTemplateResponseArray)(nil)).Elem()
}

func (i EventHandlerTemplateResponseArrayMap) ToEventHandlerTemplateResponseArrayMapOutput() EventHandlerTemplateResponseArrayMapOutput {
	return i.ToEventHandlerTemplateResponseArrayMapOutputWithContext(context.Background())
}

func (i EventHandlerTemplateResponseArrayMap) ToEventHandlerTemplateResponseArrayMapOutputWithContext(ctx context.Context) EventHandlerTemplateResponseArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerTemplateResponseArrayMapOutput)
}

type EventHandlerTemplateResponseArrayMapOutput struct{ *pulumi.OutputState }

func (EventHandlerTemplateResponseArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventHandlerTemplateResponseArray)(nil)).Elem()
}

func (o EventHandlerTemplateResponseArrayMapOutput) ToEventHandlerTemplateResponseArrayMapOutput() EventHandlerTemplateResponseArrayMapOutput {
	return o
}

func (o EventHandlerTemplateResponseArrayMapOutput) ToEventHandlerTemplateResponseArrayMapOutputWithContext(ctx context.Context) EventHandlerTemplateResponseArrayMapOutput {
	return o
}

func (o EventHandlerTemplateResponseArrayMapOutput) MapIndex(k pulumi.StringInput) EventHandlerTemplateResponseArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventHandlerTemplateResponseArray {
		return vs[0].(map[string]EventHandlerTemplateResponseArray)[vs[1].(string)]
	}).(EventHandlerTemplateResponseArrayOutput)
}





type EventHandlerTemplateResponseArrayMapInput interface {
	pulumi.Input

	ToEventHandlerTemplateResponseArrayMapOutput() EventHandlerTemplateResponseArrayMapOutput
	ToEventHandlerTemplateResponseArrayMapOutputWithContext(context.Context) EventHandlerTemplateResponseArrayMapOutput
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventHandlerSettingsInput)(nil)).Elem(), EventHandlerSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHandlerSettingsPtrInput)(nil)).Elem(), EventHandlerSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHandlerSettingsResponseInput)(nil)).Elem(), EventHandlerSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHandlerSettingsResponsePtrInput)(nil)).Elem(), EventHandlerSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHandlerTemplateInput)(nil)).Elem(), EventHandlerTemplateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHandlerTemplateArrayInput)(nil)).Elem(), EventHandlerTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHandlerTemplateResponseInput)(nil)).Elem(), EventHandlerTemplateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHandlerTemplateResponseArrayInput)(nil)).Elem(), EventHandlerTemplateResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveTraceCategoryInput)(nil)).Elem(), LiveTraceCategoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveTraceCategoryArrayInput)(nil)).Elem(), LiveTraceCategoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveTraceCategoryResponseInput)(nil)).Elem(), LiveTraceCategoryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveTraceCategoryResponseArrayInput)(nil)).Elem(), LiveTraceCategoryResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveTraceConfigurationInput)(nil)).Elem(), LiveTraceConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveTraceConfigurationPtrInput)(nil)).Elem(), LiveTraceConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveTraceConfigurationResponseInput)(nil)).Elem(), LiveTraceConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveTraceConfigurationResponsePtrInput)(nil)).Elem(), LiveTraceConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentityInput)(nil)).Elem(), ManagedIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentityPtrInput)(nil)).Elem(), ManagedIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentityResponseInput)(nil)).Elem(), ManagedIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentityResponsePtrInput)(nil)).Elem(), ManagedIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentitySettingsInput)(nil)).Elem(), ManagedIdentitySettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentitySettingsPtrInput)(nil)).Elem(), ManagedIdentitySettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentitySettingsResponseInput)(nil)).Elem(), ManagedIdentitySettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentitySettingsResponsePtrInput)(nil)).Elem(), ManagedIdentitySettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkACLInput)(nil)).Elem(), NetworkACLArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkACLPtrInput)(nil)).Elem(), NetworkACLArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkACLResponseInput)(nil)).Elem(), NetworkACLResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkACLResponsePtrInput)(nil)).Elem(), NetworkACLResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointInput)(nil)).Elem(), PrivateEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointPtrInput)(nil)).Elem(), PrivateEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointACLInput)(nil)).Elem(), PrivateEndpointACLArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointACLArrayInput)(nil)).Elem(), PrivateEndpointACLArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointACLResponseInput)(nil)).Elem(), PrivateEndpointACLResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointACLResponseArrayInput)(nil)).Elem(), PrivateEndpointACLResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseInput)(nil)).Elem(), PrivateEndpointConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseArrayInput)(nil)).Elem(), PrivateEndpointConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponseInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponsePtrInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponsePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuInput)(nil)).Elem(), ResourceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuPtrInput)(nil)).Elem(), ResourceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuResponseInput)(nil)).Elem(), ResourceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuResponsePtrInput)(nil)).Elem(), ResourceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourceResponseInput)(nil)).Elem(), SharedPrivateLinkResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourceResponseArrayInput)(nil)).Elem(), SharedPrivateLinkResourceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpstreamAuthSettingsInput)(nil)).Elem(), UpstreamAuthSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpstreamAuthSettingsPtrInput)(nil)).Elem(), UpstreamAuthSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpstreamAuthSettingsResponseInput)(nil)).Elem(), UpstreamAuthSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpstreamAuthSettingsResponsePtrInput)(nil)).Elem(), UpstreamAuthSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedIdentityPropertyResponseInput)(nil)).Elem(), UserAssignedIdentityPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedIdentityPropertyResponseMapInput)(nil)).Elem(), UserAssignedIdentityPropertyResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebPubSubNetworkACLsInput)(nil)).Elem(), WebPubSubNetworkACLsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebPubSubNetworkACLsPtrInput)(nil)).Elem(), WebPubSubNetworkACLsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebPubSubNetworkACLsResponseInput)(nil)).Elem(), WebPubSubNetworkACLsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebPubSubNetworkACLsResponsePtrInput)(nil)).Elem(), WebPubSubNetworkACLsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebPubSubTlsSettingsInput)(nil)).Elem(), WebPubSubTlsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebPubSubTlsSettingsPtrInput)(nil)).Elem(), WebPubSubTlsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebPubSubTlsSettingsResponseInput)(nil)).Elem(), WebPubSubTlsSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebPubSubTlsSettingsResponsePtrInput)(nil)).Elem(), WebPubSubTlsSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHandlerTemplateArrayMapInput)(nil)).Elem(), EventHandlerTemplateArrayMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHandlerTemplateResponseArrayMapInput)(nil)).Elem(), EventHandlerTemplateResponseArrayMap{})
	pulumi.RegisterOutputType(EventHandlerSettingsOutput{})
	pulumi.RegisterOutputType(EventHandlerSettingsPtrOutput{})
	pulumi.RegisterOutputType(EventHandlerSettingsResponseOutput{})
	pulumi.RegisterOutputType(EventHandlerSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EventHandlerTemplateOutput{})
	pulumi.RegisterOutputType(EventHandlerTemplateArrayOutput{})
	pulumi.RegisterOutputType(EventHandlerTemplateResponseOutput{})
	pulumi.RegisterOutputType(EventHandlerTemplateResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveTraceCategoryOutput{})
	pulumi.RegisterOutputType(LiveTraceCategoryArrayOutput{})
	pulumi.RegisterOutputType(LiveTraceCategoryResponseOutput{})
	pulumi.RegisterOutputType(LiveTraceCategoryResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveTraceConfigurationOutput{})
	pulumi.RegisterOutputType(LiveTraceConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LiveTraceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LiveTraceConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentitySettingsOutput{})
	pulumi.RegisterOutputType(ManagedIdentitySettingsPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentitySettingsResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentitySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkACLOutput{})
	pulumi.RegisterOutputType(NetworkACLPtrOutput{})
	pulumi.RegisterOutputType(NetworkACLResponseOutput{})
	pulumi.RegisterOutputType(NetworkACLResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UpstreamAuthSettingsOutput{})
	pulumi.RegisterOutputType(UpstreamAuthSettingsPtrOutput{})
	pulumi.RegisterOutputType(UpstreamAuthSettingsResponseOutput{})
	pulumi.RegisterOutputType(UpstreamAuthSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertyResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertyResponseMapOutput{})
	pulumi.RegisterOutputType(WebPubSubNetworkACLsOutput{})
	pulumi.RegisterOutputType(WebPubSubNetworkACLsPtrOutput{})
	pulumi.RegisterOutputType(WebPubSubNetworkACLsResponseOutput{})
	pulumi.RegisterOutputType(WebPubSubNetworkACLsResponsePtrOutput{})
	pulumi.RegisterOutputType(WebPubSubTlsSettingsOutput{})
	pulumi.RegisterOutputType(WebPubSubTlsSettingsPtrOutput{})
	pulumi.RegisterOutputType(WebPubSubTlsSettingsResponseOutput{})
	pulumi.RegisterOutputType(WebPubSubTlsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EventHandlerTemplateArrayMapOutput{})
	pulumi.RegisterOutputType(EventHandlerTemplateResponseArrayMapOutput{})
}
