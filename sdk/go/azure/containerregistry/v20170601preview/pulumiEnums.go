


package v20170601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SkuName string

const (
	SkuNameBasic             = SkuName("Basic")
	SkuName_Managed_Basic    = SkuName("Managed_Basic")
	SkuName_Managed_Standard = SkuName("Managed_Standard")
	SkuName_Managed_Premium  = SkuName("Managed_Premium")
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

type WebhookAction string

const (
	WebhookActionPush   = WebhookAction("push")
	WebhookActionDelete = WebhookAction("delete")
)

func (WebhookAction) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookAction)(nil)).Elem()
}

func (e WebhookAction) ToWebhookActionOutput() WebhookActionOutput {
	return pulumi.ToOutput(e).(WebhookActionOutput)
}

func (e WebhookAction) ToWebhookActionOutputWithContext(ctx context.Context) WebhookActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebhookActionOutput)
}

func (e WebhookAction) ToWebhookActionPtrOutput() WebhookActionPtrOutput {
	return e.ToWebhookActionPtrOutputWithContext(context.Background())
}

func (e WebhookAction) ToWebhookActionPtrOutputWithContext(ctx context.Context) WebhookActionPtrOutput {
	return WebhookAction(e).ToWebhookActionOutputWithContext(ctx).ToWebhookActionPtrOutputWithContext(ctx)
}

func (e WebhookAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebhookAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebhookAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebhookAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebhookActionOutput struct{ *pulumi.OutputState }

func (WebhookActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookAction)(nil)).Elem()
}

func (o WebhookActionOutput) ToWebhookActionOutput() WebhookActionOutput {
	return o
}

func (o WebhookActionOutput) ToWebhookActionOutputWithContext(ctx context.Context) WebhookActionOutput {
	return o
}

func (o WebhookActionOutput) ToWebhookActionPtrOutput() WebhookActionPtrOutput {
	return o.ToWebhookActionPtrOutputWithContext(context.Background())
}

func (o WebhookActionOutput) ToWebhookActionPtrOutputWithContext(ctx context.Context) WebhookActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebhookAction) *WebhookAction {
		return &v
	}).(WebhookActionPtrOutput)
}

func (o WebhookActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebhookActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebhookAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebhookActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebhookActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebhookAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebhookActionPtrOutput struct{ *pulumi.OutputState }

func (WebhookActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookAction)(nil)).Elem()
}

func (o WebhookActionPtrOutput) ToWebhookActionPtrOutput() WebhookActionPtrOutput {
	return o
}

func (o WebhookActionPtrOutput) ToWebhookActionPtrOutputWithContext(ctx context.Context) WebhookActionPtrOutput {
	return o
}

func (o WebhookActionPtrOutput) Elem() WebhookActionOutput {
	return o.ApplyT(func(v *WebhookAction) WebhookAction {
		if v != nil {
			return *v
		}
		var ret WebhookAction
		return ret
	}).(WebhookActionOutput)
}

func (o WebhookActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebhookActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebhookAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebhookActionInput interface {
	pulumi.Input

	ToWebhookActionOutput() WebhookActionOutput
	ToWebhookActionOutputWithContext(context.Context) WebhookActionOutput
}

var webhookActionPtrType = reflect.TypeOf((**WebhookAction)(nil)).Elem()

type WebhookActionPtrInput interface {
	pulumi.Input

	ToWebhookActionPtrOutput() WebhookActionPtrOutput
	ToWebhookActionPtrOutputWithContext(context.Context) WebhookActionPtrOutput
}

type webhookActionPtr string

func WebhookActionPtr(v string) WebhookActionPtrInput {
	return (*webhookActionPtr)(&v)
}

func (*webhookActionPtr) ElementType() reflect.Type {
	return webhookActionPtrType
}

func (in *webhookActionPtr) ToWebhookActionPtrOutput() WebhookActionPtrOutput {
	return pulumi.ToOutput(in).(WebhookActionPtrOutput)
}

func (in *webhookActionPtr) ToWebhookActionPtrOutputWithContext(ctx context.Context) WebhookActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebhookActionPtrOutput)
}

type WebhookStatus string

const (
	WebhookStatusEnabled  = WebhookStatus("enabled")
	WebhookStatusDisabled = WebhookStatus("disabled")
)

func (WebhookStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookStatus)(nil)).Elem()
}

func (e WebhookStatus) ToWebhookStatusOutput() WebhookStatusOutput {
	return pulumi.ToOutput(e).(WebhookStatusOutput)
}

func (e WebhookStatus) ToWebhookStatusOutputWithContext(ctx context.Context) WebhookStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebhookStatusOutput)
}

func (e WebhookStatus) ToWebhookStatusPtrOutput() WebhookStatusPtrOutput {
	return e.ToWebhookStatusPtrOutputWithContext(context.Background())
}

func (e WebhookStatus) ToWebhookStatusPtrOutputWithContext(ctx context.Context) WebhookStatusPtrOutput {
	return WebhookStatus(e).ToWebhookStatusOutputWithContext(ctx).ToWebhookStatusPtrOutputWithContext(ctx)
}

func (e WebhookStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebhookStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebhookStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebhookStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebhookStatusOutput struct{ *pulumi.OutputState }

func (WebhookStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookStatus)(nil)).Elem()
}

func (o WebhookStatusOutput) ToWebhookStatusOutput() WebhookStatusOutput {
	return o
}

func (o WebhookStatusOutput) ToWebhookStatusOutputWithContext(ctx context.Context) WebhookStatusOutput {
	return o
}

func (o WebhookStatusOutput) ToWebhookStatusPtrOutput() WebhookStatusPtrOutput {
	return o.ToWebhookStatusPtrOutputWithContext(context.Background())
}

func (o WebhookStatusOutput) ToWebhookStatusPtrOutputWithContext(ctx context.Context) WebhookStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebhookStatus) *WebhookStatus {
		return &v
	}).(WebhookStatusPtrOutput)
}

func (o WebhookStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebhookStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebhookStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebhookStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebhookStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebhookStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebhookStatusPtrOutput struct{ *pulumi.OutputState }

func (WebhookStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookStatus)(nil)).Elem()
}

func (o WebhookStatusPtrOutput) ToWebhookStatusPtrOutput() WebhookStatusPtrOutput {
	return o
}

func (o WebhookStatusPtrOutput) ToWebhookStatusPtrOutputWithContext(ctx context.Context) WebhookStatusPtrOutput {
	return o
}

func (o WebhookStatusPtrOutput) Elem() WebhookStatusOutput {
	return o.ApplyT(func(v *WebhookStatus) WebhookStatus {
		if v != nil {
			return *v
		}
		var ret WebhookStatus
		return ret
	}).(WebhookStatusOutput)
}

func (o WebhookStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebhookStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebhookStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebhookStatusInput interface {
	pulumi.Input

	ToWebhookStatusOutput() WebhookStatusOutput
	ToWebhookStatusOutputWithContext(context.Context) WebhookStatusOutput
}

var webhookStatusPtrType = reflect.TypeOf((**WebhookStatus)(nil)).Elem()

type WebhookStatusPtrInput interface {
	pulumi.Input

	ToWebhookStatusPtrOutput() WebhookStatusPtrOutput
	ToWebhookStatusPtrOutputWithContext(context.Context) WebhookStatusPtrOutput
}

type webhookStatusPtr string

func WebhookStatusPtr(v string) WebhookStatusPtrInput {
	return (*webhookStatusPtr)(&v)
}

func (*webhookStatusPtr) ElementType() reflect.Type {
	return webhookStatusPtrType
}

func (in *webhookStatusPtr) ToWebhookStatusPtrOutput() WebhookStatusPtrOutput {
	return pulumi.ToOutput(in).(WebhookStatusPtrOutput)
}

func (in *webhookStatusPtr) ToWebhookStatusPtrOutputWithContext(ctx context.Context) WebhookStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebhookStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNameInput)(nil)).Elem(), SkuName("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNamePtrInput)(nil)).Elem(), SkuName("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookActionInput)(nil)).Elem(), WebhookAction("push"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookActionPtrInput)(nil)).Elem(), WebhookAction("push"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookStatusInput)(nil)).Elem(), WebhookStatus("enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookStatusPtrInput)(nil)).Elem(), WebhookStatus("enabled"))
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(WebhookActionOutput{})
	pulumi.RegisterOutputType(WebhookActionPtrOutput{})
	pulumi.RegisterOutputType(WebhookStatusOutput{})
	pulumi.RegisterOutputType(WebhookStatusPtrOutput{})
}
