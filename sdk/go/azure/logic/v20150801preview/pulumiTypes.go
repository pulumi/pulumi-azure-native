


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AS2AcknowledgementConnectionSettings struct {
	IgnoreCertificateNameMismatch *bool `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       *bool `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue *bool `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             *bool `pulumi:"unfoldHttpHeaders"`
}





type AS2AcknowledgementConnectionSettingsInput interface {
	pulumi.Input

	ToAS2AcknowledgementConnectionSettingsOutput() AS2AcknowledgementConnectionSettingsOutput
	ToAS2AcknowledgementConnectionSettingsOutputWithContext(context.Context) AS2AcknowledgementConnectionSettingsOutput
}

type AS2AcknowledgementConnectionSettingsArgs struct {
	IgnoreCertificateNameMismatch pulumi.BoolPtrInput `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       pulumi.BoolPtrInput `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue pulumi.BoolPtrInput `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             pulumi.BoolPtrInput `pulumi:"unfoldHttpHeaders"`
}

func (AS2AcknowledgementConnectionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AcknowledgementConnectionSettings)(nil)).Elem()
}

func (i AS2AcknowledgementConnectionSettingsArgs) ToAS2AcknowledgementConnectionSettingsOutput() AS2AcknowledgementConnectionSettingsOutput {
	return i.ToAS2AcknowledgementConnectionSettingsOutputWithContext(context.Background())
}

func (i AS2AcknowledgementConnectionSettingsArgs) ToAS2AcknowledgementConnectionSettingsOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AcknowledgementConnectionSettingsOutput)
}

func (i AS2AcknowledgementConnectionSettingsArgs) ToAS2AcknowledgementConnectionSettingsPtrOutput() AS2AcknowledgementConnectionSettingsPtrOutput {
	return i.ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(context.Background())
}

func (i AS2AcknowledgementConnectionSettingsArgs) ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AcknowledgementConnectionSettingsOutput).ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(ctx)
}









type AS2AcknowledgementConnectionSettingsPtrInput interface {
	pulumi.Input

	ToAS2AcknowledgementConnectionSettingsPtrOutput() AS2AcknowledgementConnectionSettingsPtrOutput
	ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(context.Context) AS2AcknowledgementConnectionSettingsPtrOutput
}

type as2acknowledgementConnectionSettingsPtrType AS2AcknowledgementConnectionSettingsArgs

func AS2AcknowledgementConnectionSettingsPtr(v *AS2AcknowledgementConnectionSettingsArgs) AS2AcknowledgementConnectionSettingsPtrInput {
	return (*as2acknowledgementConnectionSettingsPtrType)(v)
}

func (*as2acknowledgementConnectionSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AcknowledgementConnectionSettings)(nil)).Elem()
}

func (i *as2acknowledgementConnectionSettingsPtrType) ToAS2AcknowledgementConnectionSettingsPtrOutput() AS2AcknowledgementConnectionSettingsPtrOutput {
	return i.ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(context.Background())
}

func (i *as2acknowledgementConnectionSettingsPtrType) ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AcknowledgementConnectionSettingsPtrOutput)
}

type AS2AcknowledgementConnectionSettingsOutput struct{ *pulumi.OutputState }

func (AS2AcknowledgementConnectionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AcknowledgementConnectionSettings)(nil)).Elem()
}

func (o AS2AcknowledgementConnectionSettingsOutput) ToAS2AcknowledgementConnectionSettingsOutput() AS2AcknowledgementConnectionSettingsOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsOutput) ToAS2AcknowledgementConnectionSettingsOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsOutput) ToAS2AcknowledgementConnectionSettingsPtrOutput() AS2AcknowledgementConnectionSettingsPtrOutput {
	return o.ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(context.Background())
}

func (o AS2AcknowledgementConnectionSettingsOutput) ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2AcknowledgementConnectionSettings) *AS2AcknowledgementConnectionSettings {
		return &v
	}).(AS2AcknowledgementConnectionSettingsPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettings) *bool { return v.IgnoreCertificateNameMismatch }).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettings) *bool { return v.KeepHttpConnectionAlive }).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettings) *bool { return v.SupportHttpStatusCodeContinue }).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettings) *bool { return v.UnfoldHttpHeaders }).(pulumi.BoolPtrOutput)
}

type AS2AcknowledgementConnectionSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2AcknowledgementConnectionSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AcknowledgementConnectionSettings)(nil)).Elem()
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) ToAS2AcknowledgementConnectionSettingsPtrOutput() AS2AcknowledgementConnectionSettingsPtrOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsPtrOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) Elem() AS2AcknowledgementConnectionSettingsOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettings) AS2AcknowledgementConnectionSettings {
		if v != nil {
			return *v
		}
		var ret AS2AcknowledgementConnectionSettings
		return ret
	}).(AS2AcknowledgementConnectionSettingsOutput)
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.IgnoreCertificateNameMismatch
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.KeepHttpConnectionAlive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SupportHttpStatusCodeContinue
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.UnfoldHttpHeaders
	}).(pulumi.BoolPtrOutput)
}

type AS2AcknowledgementConnectionSettingsResponse struct {
	IgnoreCertificateNameMismatch *bool `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       *bool `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue *bool `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             *bool `pulumi:"unfoldHttpHeaders"`
}

type AS2AcknowledgementConnectionSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2AcknowledgementConnectionSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AcknowledgementConnectionSettingsResponse)(nil)).Elem()
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) ToAS2AcknowledgementConnectionSettingsResponseOutput() AS2AcknowledgementConnectionSettingsResponseOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) ToAS2AcknowledgementConnectionSettingsResponseOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsResponseOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettingsResponse) *bool { return v.IgnoreCertificateNameMismatch }).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettingsResponse) *bool { return v.KeepHttpConnectionAlive }).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettingsResponse) *bool { return v.SupportHttpStatusCodeContinue }).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettingsResponse) *bool { return v.UnfoldHttpHeaders }).(pulumi.BoolPtrOutput)
}

type AS2AcknowledgementConnectionSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2AcknowledgementConnectionSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AcknowledgementConnectionSettingsResponse)(nil)).Elem()
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) ToAS2AcknowledgementConnectionSettingsResponsePtrOutput() AS2AcknowledgementConnectionSettingsResponsePtrOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) ToAS2AcknowledgementConnectionSettingsResponsePtrOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsResponsePtrOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) Elem() AS2AcknowledgementConnectionSettingsResponseOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettingsResponse) AS2AcknowledgementConnectionSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2AcknowledgementConnectionSettingsResponse
		return ret
	}).(AS2AcknowledgementConnectionSettingsResponseOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IgnoreCertificateNameMismatch
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.KeepHttpConnectionAlive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SupportHttpStatusCodeContinue
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UnfoldHttpHeaders
	}).(pulumi.BoolPtrOutput)
}

type AS2AgreementContent struct {
	ReceiveAgreement *AS2OneWayAgreement `pulumi:"receiveAgreement"`
	SendAgreement    *AS2OneWayAgreement `pulumi:"sendAgreement"`
}





type AS2AgreementContentInput interface {
	pulumi.Input

	ToAS2AgreementContentOutput() AS2AgreementContentOutput
	ToAS2AgreementContentOutputWithContext(context.Context) AS2AgreementContentOutput
}

type AS2AgreementContentArgs struct {
	ReceiveAgreement AS2OneWayAgreementPtrInput `pulumi:"receiveAgreement"`
	SendAgreement    AS2OneWayAgreementPtrInput `pulumi:"sendAgreement"`
}

func (AS2AgreementContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AgreementContent)(nil)).Elem()
}

func (i AS2AgreementContentArgs) ToAS2AgreementContentOutput() AS2AgreementContentOutput {
	return i.ToAS2AgreementContentOutputWithContext(context.Background())
}

func (i AS2AgreementContentArgs) ToAS2AgreementContentOutputWithContext(ctx context.Context) AS2AgreementContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AgreementContentOutput)
}

func (i AS2AgreementContentArgs) ToAS2AgreementContentPtrOutput() AS2AgreementContentPtrOutput {
	return i.ToAS2AgreementContentPtrOutputWithContext(context.Background())
}

func (i AS2AgreementContentArgs) ToAS2AgreementContentPtrOutputWithContext(ctx context.Context) AS2AgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AgreementContentOutput).ToAS2AgreementContentPtrOutputWithContext(ctx)
}









type AS2AgreementContentPtrInput interface {
	pulumi.Input

	ToAS2AgreementContentPtrOutput() AS2AgreementContentPtrOutput
	ToAS2AgreementContentPtrOutputWithContext(context.Context) AS2AgreementContentPtrOutput
}

type as2agreementContentPtrType AS2AgreementContentArgs

func AS2AgreementContentPtr(v *AS2AgreementContentArgs) AS2AgreementContentPtrInput {
	return (*as2agreementContentPtrType)(v)
}

func (*as2agreementContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AgreementContent)(nil)).Elem()
}

func (i *as2agreementContentPtrType) ToAS2AgreementContentPtrOutput() AS2AgreementContentPtrOutput {
	return i.ToAS2AgreementContentPtrOutputWithContext(context.Background())
}

func (i *as2agreementContentPtrType) ToAS2AgreementContentPtrOutputWithContext(ctx context.Context) AS2AgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AgreementContentPtrOutput)
}

type AS2AgreementContentOutput struct{ *pulumi.OutputState }

func (AS2AgreementContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AgreementContent)(nil)).Elem()
}

func (o AS2AgreementContentOutput) ToAS2AgreementContentOutput() AS2AgreementContentOutput {
	return o
}

func (o AS2AgreementContentOutput) ToAS2AgreementContentOutputWithContext(ctx context.Context) AS2AgreementContentOutput {
	return o
}

func (o AS2AgreementContentOutput) ToAS2AgreementContentPtrOutput() AS2AgreementContentPtrOutput {
	return o.ToAS2AgreementContentPtrOutputWithContext(context.Background())
}

func (o AS2AgreementContentOutput) ToAS2AgreementContentPtrOutputWithContext(ctx context.Context) AS2AgreementContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2AgreementContent) *AS2AgreementContent {
		return &v
	}).(AS2AgreementContentPtrOutput)
}

func (o AS2AgreementContentOutput) ReceiveAgreement() AS2OneWayAgreementPtrOutput {
	return o.ApplyT(func(v AS2AgreementContent) *AS2OneWayAgreement { return v.ReceiveAgreement }).(AS2OneWayAgreementPtrOutput)
}

func (o AS2AgreementContentOutput) SendAgreement() AS2OneWayAgreementPtrOutput {
	return o.ApplyT(func(v AS2AgreementContent) *AS2OneWayAgreement { return v.SendAgreement }).(AS2OneWayAgreementPtrOutput)
}

type AS2AgreementContentPtrOutput struct{ *pulumi.OutputState }

func (AS2AgreementContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AgreementContent)(nil)).Elem()
}

func (o AS2AgreementContentPtrOutput) ToAS2AgreementContentPtrOutput() AS2AgreementContentPtrOutput {
	return o
}

func (o AS2AgreementContentPtrOutput) ToAS2AgreementContentPtrOutputWithContext(ctx context.Context) AS2AgreementContentPtrOutput {
	return o
}

func (o AS2AgreementContentPtrOutput) Elem() AS2AgreementContentOutput {
	return o.ApplyT(func(v *AS2AgreementContent) AS2AgreementContent {
		if v != nil {
			return *v
		}
		var ret AS2AgreementContent
		return ret
	}).(AS2AgreementContentOutput)
}

func (o AS2AgreementContentPtrOutput) ReceiveAgreement() AS2OneWayAgreementPtrOutput {
	return o.ApplyT(func(v *AS2AgreementContent) *AS2OneWayAgreement {
		if v == nil {
			return nil
		}
		return v.ReceiveAgreement
	}).(AS2OneWayAgreementPtrOutput)
}

func (o AS2AgreementContentPtrOutput) SendAgreement() AS2OneWayAgreementPtrOutput {
	return o.ApplyT(func(v *AS2AgreementContent) *AS2OneWayAgreement {
		if v == nil {
			return nil
		}
		return v.SendAgreement
	}).(AS2OneWayAgreementPtrOutput)
}

type AS2AgreementContentResponse struct {
	ReceiveAgreement *AS2OneWayAgreementResponse `pulumi:"receiveAgreement"`
	SendAgreement    *AS2OneWayAgreementResponse `pulumi:"sendAgreement"`
}

type AS2AgreementContentResponseOutput struct{ *pulumi.OutputState }

func (AS2AgreementContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AgreementContentResponse)(nil)).Elem()
}

func (o AS2AgreementContentResponseOutput) ToAS2AgreementContentResponseOutput() AS2AgreementContentResponseOutput {
	return o
}

func (o AS2AgreementContentResponseOutput) ToAS2AgreementContentResponseOutputWithContext(ctx context.Context) AS2AgreementContentResponseOutput {
	return o
}

func (o AS2AgreementContentResponseOutput) ReceiveAgreement() AS2OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v AS2AgreementContentResponse) *AS2OneWayAgreementResponse { return v.ReceiveAgreement }).(AS2OneWayAgreementResponsePtrOutput)
}

func (o AS2AgreementContentResponseOutput) SendAgreement() AS2OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v AS2AgreementContentResponse) *AS2OneWayAgreementResponse { return v.SendAgreement }).(AS2OneWayAgreementResponsePtrOutput)
}

type AS2AgreementContentResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2AgreementContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AgreementContentResponse)(nil)).Elem()
}

func (o AS2AgreementContentResponsePtrOutput) ToAS2AgreementContentResponsePtrOutput() AS2AgreementContentResponsePtrOutput {
	return o
}

func (o AS2AgreementContentResponsePtrOutput) ToAS2AgreementContentResponsePtrOutputWithContext(ctx context.Context) AS2AgreementContentResponsePtrOutput {
	return o
}

func (o AS2AgreementContentResponsePtrOutput) Elem() AS2AgreementContentResponseOutput {
	return o.ApplyT(func(v *AS2AgreementContentResponse) AS2AgreementContentResponse {
		if v != nil {
			return *v
		}
		var ret AS2AgreementContentResponse
		return ret
	}).(AS2AgreementContentResponseOutput)
}

func (o AS2AgreementContentResponsePtrOutput) ReceiveAgreement() AS2OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *AS2AgreementContentResponse) *AS2OneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return v.ReceiveAgreement
	}).(AS2OneWayAgreementResponsePtrOutput)
}

func (o AS2AgreementContentResponsePtrOutput) SendAgreement() AS2OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *AS2AgreementContentResponse) *AS2OneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return v.SendAgreement
	}).(AS2OneWayAgreementResponsePtrOutput)
}

type AS2EnvelopeSettings struct {
	AutogenerateFileName                    *bool   `pulumi:"autogenerateFileName"`
	FileNameTemplate                        *string `pulumi:"fileNameTemplate"`
	MessageContentType                      *string `pulumi:"messageContentType"`
	SuspendMessageOnFileNameGenerationError *bool   `pulumi:"suspendMessageOnFileNameGenerationError"`
	TransmitFileNameInMimeHeader            *bool   `pulumi:"transmitFileNameInMimeHeader"`
}





type AS2EnvelopeSettingsInput interface {
	pulumi.Input

	ToAS2EnvelopeSettingsOutput() AS2EnvelopeSettingsOutput
	ToAS2EnvelopeSettingsOutputWithContext(context.Context) AS2EnvelopeSettingsOutput
}

type AS2EnvelopeSettingsArgs struct {
	AutogenerateFileName                    pulumi.BoolPtrInput   `pulumi:"autogenerateFileName"`
	FileNameTemplate                        pulumi.StringPtrInput `pulumi:"fileNameTemplate"`
	MessageContentType                      pulumi.StringPtrInput `pulumi:"messageContentType"`
	SuspendMessageOnFileNameGenerationError pulumi.BoolPtrInput   `pulumi:"suspendMessageOnFileNameGenerationError"`
	TransmitFileNameInMimeHeader            pulumi.BoolPtrInput   `pulumi:"transmitFileNameInMimeHeader"`
}

func (AS2EnvelopeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2EnvelopeSettings)(nil)).Elem()
}

func (i AS2EnvelopeSettingsArgs) ToAS2EnvelopeSettingsOutput() AS2EnvelopeSettingsOutput {
	return i.ToAS2EnvelopeSettingsOutputWithContext(context.Background())
}

func (i AS2EnvelopeSettingsArgs) ToAS2EnvelopeSettingsOutputWithContext(ctx context.Context) AS2EnvelopeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2EnvelopeSettingsOutput)
}

func (i AS2EnvelopeSettingsArgs) ToAS2EnvelopeSettingsPtrOutput() AS2EnvelopeSettingsPtrOutput {
	return i.ToAS2EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i AS2EnvelopeSettingsArgs) ToAS2EnvelopeSettingsPtrOutputWithContext(ctx context.Context) AS2EnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2EnvelopeSettingsOutput).ToAS2EnvelopeSettingsPtrOutputWithContext(ctx)
}









type AS2EnvelopeSettingsPtrInput interface {
	pulumi.Input

	ToAS2EnvelopeSettingsPtrOutput() AS2EnvelopeSettingsPtrOutput
	ToAS2EnvelopeSettingsPtrOutputWithContext(context.Context) AS2EnvelopeSettingsPtrOutput
}

type as2envelopeSettingsPtrType AS2EnvelopeSettingsArgs

func AS2EnvelopeSettingsPtr(v *AS2EnvelopeSettingsArgs) AS2EnvelopeSettingsPtrInput {
	return (*as2envelopeSettingsPtrType)(v)
}

func (*as2envelopeSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2EnvelopeSettings)(nil)).Elem()
}

func (i *as2envelopeSettingsPtrType) ToAS2EnvelopeSettingsPtrOutput() AS2EnvelopeSettingsPtrOutput {
	return i.ToAS2EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i *as2envelopeSettingsPtrType) ToAS2EnvelopeSettingsPtrOutputWithContext(ctx context.Context) AS2EnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2EnvelopeSettingsPtrOutput)
}

type AS2EnvelopeSettingsOutput struct{ *pulumi.OutputState }

func (AS2EnvelopeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2EnvelopeSettings)(nil)).Elem()
}

func (o AS2EnvelopeSettingsOutput) ToAS2EnvelopeSettingsOutput() AS2EnvelopeSettingsOutput {
	return o
}

func (o AS2EnvelopeSettingsOutput) ToAS2EnvelopeSettingsOutputWithContext(ctx context.Context) AS2EnvelopeSettingsOutput {
	return o
}

func (o AS2EnvelopeSettingsOutput) ToAS2EnvelopeSettingsPtrOutput() AS2EnvelopeSettingsPtrOutput {
	return o.ToAS2EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (o AS2EnvelopeSettingsOutput) ToAS2EnvelopeSettingsPtrOutputWithContext(ctx context.Context) AS2EnvelopeSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2EnvelopeSettings) *AS2EnvelopeSettings {
		return &v
	}).(AS2EnvelopeSettingsPtrOutput)
}

func (o AS2EnvelopeSettingsOutput) AutogenerateFileName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2EnvelopeSettings) *bool { return v.AutogenerateFileName }).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsOutput) FileNameTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2EnvelopeSettings) *string { return v.FileNameTemplate }).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsOutput) MessageContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2EnvelopeSettings) *string { return v.MessageContentType }).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsOutput) SuspendMessageOnFileNameGenerationError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2EnvelopeSettings) *bool { return v.SuspendMessageOnFileNameGenerationError }).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsOutput) TransmitFileNameInMimeHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2EnvelopeSettings) *bool { return v.TransmitFileNameInMimeHeader }).(pulumi.BoolPtrOutput)
}

type AS2EnvelopeSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2EnvelopeSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2EnvelopeSettings)(nil)).Elem()
}

func (o AS2EnvelopeSettingsPtrOutput) ToAS2EnvelopeSettingsPtrOutput() AS2EnvelopeSettingsPtrOutput {
	return o
}

func (o AS2EnvelopeSettingsPtrOutput) ToAS2EnvelopeSettingsPtrOutputWithContext(ctx context.Context) AS2EnvelopeSettingsPtrOutput {
	return o
}

func (o AS2EnvelopeSettingsPtrOutput) Elem() AS2EnvelopeSettingsOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) AS2EnvelopeSettings {
		if v != nil {
			return *v
		}
		var ret AS2EnvelopeSettings
		return ret
	}).(AS2EnvelopeSettingsOutput)
}

func (o AS2EnvelopeSettingsPtrOutput) AutogenerateFileName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.AutogenerateFileName
	}).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsPtrOutput) FileNameTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.FileNameTemplate
	}).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsPtrOutput) MessageContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.MessageContentType
	}).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsPtrOutput) SuspendMessageOnFileNameGenerationError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SuspendMessageOnFileNameGenerationError
	}).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsPtrOutput) TransmitFileNameInMimeHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.TransmitFileNameInMimeHeader
	}).(pulumi.BoolPtrOutput)
}

type AS2EnvelopeSettingsResponse struct {
	AutogenerateFileName                    *bool   `pulumi:"autogenerateFileName"`
	FileNameTemplate                        *string `pulumi:"fileNameTemplate"`
	MessageContentType                      *string `pulumi:"messageContentType"`
	SuspendMessageOnFileNameGenerationError *bool   `pulumi:"suspendMessageOnFileNameGenerationError"`
	TransmitFileNameInMimeHeader            *bool   `pulumi:"transmitFileNameInMimeHeader"`
}

type AS2EnvelopeSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2EnvelopeSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2EnvelopeSettingsResponse)(nil)).Elem()
}

func (o AS2EnvelopeSettingsResponseOutput) ToAS2EnvelopeSettingsResponseOutput() AS2EnvelopeSettingsResponseOutput {
	return o
}

func (o AS2EnvelopeSettingsResponseOutput) ToAS2EnvelopeSettingsResponseOutputWithContext(ctx context.Context) AS2EnvelopeSettingsResponseOutput {
	return o
}

func (o AS2EnvelopeSettingsResponseOutput) AutogenerateFileName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2EnvelopeSettingsResponse) *bool { return v.AutogenerateFileName }).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsResponseOutput) FileNameTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2EnvelopeSettingsResponse) *string { return v.FileNameTemplate }).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsResponseOutput) MessageContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2EnvelopeSettingsResponse) *string { return v.MessageContentType }).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsResponseOutput) SuspendMessageOnFileNameGenerationError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2EnvelopeSettingsResponse) *bool { return v.SuspendMessageOnFileNameGenerationError }).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsResponseOutput) TransmitFileNameInMimeHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2EnvelopeSettingsResponse) *bool { return v.TransmitFileNameInMimeHeader }).(pulumi.BoolPtrOutput)
}

type AS2EnvelopeSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2EnvelopeSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2EnvelopeSettingsResponse)(nil)).Elem()
}

func (o AS2EnvelopeSettingsResponsePtrOutput) ToAS2EnvelopeSettingsResponsePtrOutput() AS2EnvelopeSettingsResponsePtrOutput {
	return o
}

func (o AS2EnvelopeSettingsResponsePtrOutput) ToAS2EnvelopeSettingsResponsePtrOutputWithContext(ctx context.Context) AS2EnvelopeSettingsResponsePtrOutput {
	return o
}

func (o AS2EnvelopeSettingsResponsePtrOutput) Elem() AS2EnvelopeSettingsResponseOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) AS2EnvelopeSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2EnvelopeSettingsResponse
		return ret
	}).(AS2EnvelopeSettingsResponseOutput)
}

func (o AS2EnvelopeSettingsResponsePtrOutput) AutogenerateFileName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutogenerateFileName
	}).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsResponsePtrOutput) FileNameTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FileNameTemplate
	}).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsResponsePtrOutput) MessageContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageContentType
	}).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsResponsePtrOutput) SuspendMessageOnFileNameGenerationError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SuspendMessageOnFileNameGenerationError
	}).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsResponsePtrOutput) TransmitFileNameInMimeHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.TransmitFileNameInMimeHeader
	}).(pulumi.BoolPtrOutput)
}

type AS2ErrorSettings struct {
	ResendIfMdnNotReceived  *bool `pulumi:"resendIfMdnNotReceived"`
	SuspendDuplicateMessage *bool `pulumi:"suspendDuplicateMessage"`
}





type AS2ErrorSettingsInput interface {
	pulumi.Input

	ToAS2ErrorSettingsOutput() AS2ErrorSettingsOutput
	ToAS2ErrorSettingsOutputWithContext(context.Context) AS2ErrorSettingsOutput
}

type AS2ErrorSettingsArgs struct {
	ResendIfMdnNotReceived  pulumi.BoolPtrInput `pulumi:"resendIfMdnNotReceived"`
	SuspendDuplicateMessage pulumi.BoolPtrInput `pulumi:"suspendDuplicateMessage"`
}

func (AS2ErrorSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ErrorSettings)(nil)).Elem()
}

func (i AS2ErrorSettingsArgs) ToAS2ErrorSettingsOutput() AS2ErrorSettingsOutput {
	return i.ToAS2ErrorSettingsOutputWithContext(context.Background())
}

func (i AS2ErrorSettingsArgs) ToAS2ErrorSettingsOutputWithContext(ctx context.Context) AS2ErrorSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ErrorSettingsOutput)
}

func (i AS2ErrorSettingsArgs) ToAS2ErrorSettingsPtrOutput() AS2ErrorSettingsPtrOutput {
	return i.ToAS2ErrorSettingsPtrOutputWithContext(context.Background())
}

func (i AS2ErrorSettingsArgs) ToAS2ErrorSettingsPtrOutputWithContext(ctx context.Context) AS2ErrorSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ErrorSettingsOutput).ToAS2ErrorSettingsPtrOutputWithContext(ctx)
}









type AS2ErrorSettingsPtrInput interface {
	pulumi.Input

	ToAS2ErrorSettingsPtrOutput() AS2ErrorSettingsPtrOutput
	ToAS2ErrorSettingsPtrOutputWithContext(context.Context) AS2ErrorSettingsPtrOutput
}

type as2errorSettingsPtrType AS2ErrorSettingsArgs

func AS2ErrorSettingsPtr(v *AS2ErrorSettingsArgs) AS2ErrorSettingsPtrInput {
	return (*as2errorSettingsPtrType)(v)
}

func (*as2errorSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ErrorSettings)(nil)).Elem()
}

func (i *as2errorSettingsPtrType) ToAS2ErrorSettingsPtrOutput() AS2ErrorSettingsPtrOutput {
	return i.ToAS2ErrorSettingsPtrOutputWithContext(context.Background())
}

func (i *as2errorSettingsPtrType) ToAS2ErrorSettingsPtrOutputWithContext(ctx context.Context) AS2ErrorSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ErrorSettingsPtrOutput)
}

type AS2ErrorSettingsOutput struct{ *pulumi.OutputState }

func (AS2ErrorSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ErrorSettings)(nil)).Elem()
}

func (o AS2ErrorSettingsOutput) ToAS2ErrorSettingsOutput() AS2ErrorSettingsOutput {
	return o
}

func (o AS2ErrorSettingsOutput) ToAS2ErrorSettingsOutputWithContext(ctx context.Context) AS2ErrorSettingsOutput {
	return o
}

func (o AS2ErrorSettingsOutput) ToAS2ErrorSettingsPtrOutput() AS2ErrorSettingsPtrOutput {
	return o.ToAS2ErrorSettingsPtrOutputWithContext(context.Background())
}

func (o AS2ErrorSettingsOutput) ToAS2ErrorSettingsPtrOutputWithContext(ctx context.Context) AS2ErrorSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2ErrorSettings) *AS2ErrorSettings {
		return &v
	}).(AS2ErrorSettingsPtrOutput)
}

func (o AS2ErrorSettingsOutput) ResendIfMdnNotReceived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ErrorSettings) *bool { return v.ResendIfMdnNotReceived }).(pulumi.BoolPtrOutput)
}

func (o AS2ErrorSettingsOutput) SuspendDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ErrorSettings) *bool { return v.SuspendDuplicateMessage }).(pulumi.BoolPtrOutput)
}

type AS2ErrorSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2ErrorSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ErrorSettings)(nil)).Elem()
}

func (o AS2ErrorSettingsPtrOutput) ToAS2ErrorSettingsPtrOutput() AS2ErrorSettingsPtrOutput {
	return o
}

func (o AS2ErrorSettingsPtrOutput) ToAS2ErrorSettingsPtrOutputWithContext(ctx context.Context) AS2ErrorSettingsPtrOutput {
	return o
}

func (o AS2ErrorSettingsPtrOutput) Elem() AS2ErrorSettingsOutput {
	return o.ApplyT(func(v *AS2ErrorSettings) AS2ErrorSettings {
		if v != nil {
			return *v
		}
		var ret AS2ErrorSettings
		return ret
	}).(AS2ErrorSettingsOutput)
}

func (o AS2ErrorSettingsPtrOutput) ResendIfMdnNotReceived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ErrorSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ResendIfMdnNotReceived
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ErrorSettingsPtrOutput) SuspendDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ErrorSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SuspendDuplicateMessage
	}).(pulumi.BoolPtrOutput)
}

type AS2ErrorSettingsResponse struct {
	ResendIfMdnNotReceived  *bool `pulumi:"resendIfMdnNotReceived"`
	SuspendDuplicateMessage *bool `pulumi:"suspendDuplicateMessage"`
}

type AS2ErrorSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2ErrorSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ErrorSettingsResponse)(nil)).Elem()
}

func (o AS2ErrorSettingsResponseOutput) ToAS2ErrorSettingsResponseOutput() AS2ErrorSettingsResponseOutput {
	return o
}

func (o AS2ErrorSettingsResponseOutput) ToAS2ErrorSettingsResponseOutputWithContext(ctx context.Context) AS2ErrorSettingsResponseOutput {
	return o
}

func (o AS2ErrorSettingsResponseOutput) ResendIfMdnNotReceived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ErrorSettingsResponse) *bool { return v.ResendIfMdnNotReceived }).(pulumi.BoolPtrOutput)
}

func (o AS2ErrorSettingsResponseOutput) SuspendDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ErrorSettingsResponse) *bool { return v.SuspendDuplicateMessage }).(pulumi.BoolPtrOutput)
}

type AS2ErrorSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2ErrorSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ErrorSettingsResponse)(nil)).Elem()
}

func (o AS2ErrorSettingsResponsePtrOutput) ToAS2ErrorSettingsResponsePtrOutput() AS2ErrorSettingsResponsePtrOutput {
	return o
}

func (o AS2ErrorSettingsResponsePtrOutput) ToAS2ErrorSettingsResponsePtrOutputWithContext(ctx context.Context) AS2ErrorSettingsResponsePtrOutput {
	return o
}

func (o AS2ErrorSettingsResponsePtrOutput) Elem() AS2ErrorSettingsResponseOutput {
	return o.ApplyT(func(v *AS2ErrorSettingsResponse) AS2ErrorSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2ErrorSettingsResponse
		return ret
	}).(AS2ErrorSettingsResponseOutput)
}

func (o AS2ErrorSettingsResponsePtrOutput) ResendIfMdnNotReceived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ErrorSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ResendIfMdnNotReceived
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ErrorSettingsResponsePtrOutput) SuspendDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ErrorSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SuspendDuplicateMessage
	}).(pulumi.BoolPtrOutput)
}

type AS2MdnSettings struct {
	DispositionNotificationTo  *string           `pulumi:"dispositionNotificationTo"`
	MdnText                    *string           `pulumi:"mdnText"`
	MicHashingAlgorithm        *HashingAlgorithm `pulumi:"micHashingAlgorithm"`
	NeedMdn                    *bool             `pulumi:"needMdn"`
	ReceiptDeliveryUrl         *string           `pulumi:"receiptDeliveryUrl"`
	SendInboundMdnToMessageBox *bool             `pulumi:"sendInboundMdnToMessageBox"`
	SendMdnAsynchronously      *bool             `pulumi:"sendMdnAsynchronously"`
	SignMdn                    *bool             `pulumi:"signMdn"`
	SignOutboundMdnIfOptional  *bool             `pulumi:"signOutboundMdnIfOptional"`
}





type AS2MdnSettingsInput interface {
	pulumi.Input

	ToAS2MdnSettingsOutput() AS2MdnSettingsOutput
	ToAS2MdnSettingsOutputWithContext(context.Context) AS2MdnSettingsOutput
}

type AS2MdnSettingsArgs struct {
	DispositionNotificationTo  pulumi.StringPtrInput    `pulumi:"dispositionNotificationTo"`
	MdnText                    pulumi.StringPtrInput    `pulumi:"mdnText"`
	MicHashingAlgorithm        HashingAlgorithmPtrInput `pulumi:"micHashingAlgorithm"`
	NeedMdn                    pulumi.BoolPtrInput      `pulumi:"needMdn"`
	ReceiptDeliveryUrl         pulumi.StringPtrInput    `pulumi:"receiptDeliveryUrl"`
	SendInboundMdnToMessageBox pulumi.BoolPtrInput      `pulumi:"sendInboundMdnToMessageBox"`
	SendMdnAsynchronously      pulumi.BoolPtrInput      `pulumi:"sendMdnAsynchronously"`
	SignMdn                    pulumi.BoolPtrInput      `pulumi:"signMdn"`
	SignOutboundMdnIfOptional  pulumi.BoolPtrInput      `pulumi:"signOutboundMdnIfOptional"`
}

func (AS2MdnSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MdnSettings)(nil)).Elem()
}

func (i AS2MdnSettingsArgs) ToAS2MdnSettingsOutput() AS2MdnSettingsOutput {
	return i.ToAS2MdnSettingsOutputWithContext(context.Background())
}

func (i AS2MdnSettingsArgs) ToAS2MdnSettingsOutputWithContext(ctx context.Context) AS2MdnSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MdnSettingsOutput)
}

func (i AS2MdnSettingsArgs) ToAS2MdnSettingsPtrOutput() AS2MdnSettingsPtrOutput {
	return i.ToAS2MdnSettingsPtrOutputWithContext(context.Background())
}

func (i AS2MdnSettingsArgs) ToAS2MdnSettingsPtrOutputWithContext(ctx context.Context) AS2MdnSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MdnSettingsOutput).ToAS2MdnSettingsPtrOutputWithContext(ctx)
}









type AS2MdnSettingsPtrInput interface {
	pulumi.Input

	ToAS2MdnSettingsPtrOutput() AS2MdnSettingsPtrOutput
	ToAS2MdnSettingsPtrOutputWithContext(context.Context) AS2MdnSettingsPtrOutput
}

type as2mdnSettingsPtrType AS2MdnSettingsArgs

func AS2MdnSettingsPtr(v *AS2MdnSettingsArgs) AS2MdnSettingsPtrInput {
	return (*as2mdnSettingsPtrType)(v)
}

func (*as2mdnSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MdnSettings)(nil)).Elem()
}

func (i *as2mdnSettingsPtrType) ToAS2MdnSettingsPtrOutput() AS2MdnSettingsPtrOutput {
	return i.ToAS2MdnSettingsPtrOutputWithContext(context.Background())
}

func (i *as2mdnSettingsPtrType) ToAS2MdnSettingsPtrOutputWithContext(ctx context.Context) AS2MdnSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MdnSettingsPtrOutput)
}

type AS2MdnSettingsOutput struct{ *pulumi.OutputState }

func (AS2MdnSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MdnSettings)(nil)).Elem()
}

func (o AS2MdnSettingsOutput) ToAS2MdnSettingsOutput() AS2MdnSettingsOutput {
	return o
}

func (o AS2MdnSettingsOutput) ToAS2MdnSettingsOutputWithContext(ctx context.Context) AS2MdnSettingsOutput {
	return o
}

func (o AS2MdnSettingsOutput) ToAS2MdnSettingsPtrOutput() AS2MdnSettingsPtrOutput {
	return o.ToAS2MdnSettingsPtrOutputWithContext(context.Background())
}

func (o AS2MdnSettingsOutput) ToAS2MdnSettingsPtrOutputWithContext(ctx context.Context) AS2MdnSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2MdnSettings) *AS2MdnSettings {
		return &v
	}).(AS2MdnSettingsPtrOutput)
}

func (o AS2MdnSettingsOutput) DispositionNotificationTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *string { return v.DispositionNotificationTo }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsOutput) MdnText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *string { return v.MdnText }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsOutput) MicHashingAlgorithm() HashingAlgorithmPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *HashingAlgorithm { return v.MicHashingAlgorithm }).(HashingAlgorithmPtrOutput)
}

func (o AS2MdnSettingsOutput) NeedMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *bool { return v.NeedMdn }).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsOutput) ReceiptDeliveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *string { return v.ReceiptDeliveryUrl }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsOutput) SendInboundMdnToMessageBox() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *bool { return v.SendInboundMdnToMessageBox }).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsOutput) SendMdnAsynchronously() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *bool { return v.SendMdnAsynchronously }).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsOutput) SignMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *bool { return v.SignMdn }).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsOutput) SignOutboundMdnIfOptional() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *bool { return v.SignOutboundMdnIfOptional }).(pulumi.BoolPtrOutput)
}

type AS2MdnSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2MdnSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MdnSettings)(nil)).Elem()
}

func (o AS2MdnSettingsPtrOutput) ToAS2MdnSettingsPtrOutput() AS2MdnSettingsPtrOutput {
	return o
}

func (o AS2MdnSettingsPtrOutput) ToAS2MdnSettingsPtrOutputWithContext(ctx context.Context) AS2MdnSettingsPtrOutput {
	return o
}

func (o AS2MdnSettingsPtrOutput) Elem() AS2MdnSettingsOutput {
	return o.ApplyT(func(v *AS2MdnSettings) AS2MdnSettings {
		if v != nil {
			return *v
		}
		var ret AS2MdnSettings
		return ret
	}).(AS2MdnSettingsOutput)
}

func (o AS2MdnSettingsPtrOutput) DispositionNotificationTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *string {
		if v == nil {
			return nil
		}
		return v.DispositionNotificationTo
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) MdnText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *string {
		if v == nil {
			return nil
		}
		return v.MdnText
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) MicHashingAlgorithm() HashingAlgorithmPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *HashingAlgorithm {
		if v == nil {
			return nil
		}
		return v.MicHashingAlgorithm
	}).(HashingAlgorithmPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) NeedMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *bool {
		if v == nil {
			return nil
		}
		return v.NeedMdn
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) ReceiptDeliveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReceiptDeliveryUrl
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) SendInboundMdnToMessageBox() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SendInboundMdnToMessageBox
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) SendMdnAsynchronously() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SendMdnAsynchronously
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) SignMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SignMdn
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) SignOutboundMdnIfOptional() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SignOutboundMdnIfOptional
	}).(pulumi.BoolPtrOutput)
}

type AS2MdnSettingsResponse struct {
	DispositionNotificationTo  *string `pulumi:"dispositionNotificationTo"`
	MdnText                    *string `pulumi:"mdnText"`
	MicHashingAlgorithm        *string `pulumi:"micHashingAlgorithm"`
	NeedMdn                    *bool   `pulumi:"needMdn"`
	ReceiptDeliveryUrl         *string `pulumi:"receiptDeliveryUrl"`
	SendInboundMdnToMessageBox *bool   `pulumi:"sendInboundMdnToMessageBox"`
	SendMdnAsynchronously      *bool   `pulumi:"sendMdnAsynchronously"`
	SignMdn                    *bool   `pulumi:"signMdn"`
	SignOutboundMdnIfOptional  *bool   `pulumi:"signOutboundMdnIfOptional"`
}

type AS2MdnSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2MdnSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MdnSettingsResponse)(nil)).Elem()
}

func (o AS2MdnSettingsResponseOutput) ToAS2MdnSettingsResponseOutput() AS2MdnSettingsResponseOutput {
	return o
}

func (o AS2MdnSettingsResponseOutput) ToAS2MdnSettingsResponseOutputWithContext(ctx context.Context) AS2MdnSettingsResponseOutput {
	return o
}

func (o AS2MdnSettingsResponseOutput) DispositionNotificationTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *string { return v.DispositionNotificationTo }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) MdnText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *string { return v.MdnText }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) MicHashingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *string { return v.MicHashingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) NeedMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *bool { return v.NeedMdn }).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) ReceiptDeliveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *string { return v.ReceiptDeliveryUrl }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) SendInboundMdnToMessageBox() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *bool { return v.SendInboundMdnToMessageBox }).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) SendMdnAsynchronously() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *bool { return v.SendMdnAsynchronously }).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) SignMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *bool { return v.SignMdn }).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) SignOutboundMdnIfOptional() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *bool { return v.SignOutboundMdnIfOptional }).(pulumi.BoolPtrOutput)
}

type AS2MdnSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2MdnSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MdnSettingsResponse)(nil)).Elem()
}

func (o AS2MdnSettingsResponsePtrOutput) ToAS2MdnSettingsResponsePtrOutput() AS2MdnSettingsResponsePtrOutput {
	return o
}

func (o AS2MdnSettingsResponsePtrOutput) ToAS2MdnSettingsResponsePtrOutputWithContext(ctx context.Context) AS2MdnSettingsResponsePtrOutput {
	return o
}

func (o AS2MdnSettingsResponsePtrOutput) Elem() AS2MdnSettingsResponseOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) AS2MdnSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2MdnSettingsResponse
		return ret
	}).(AS2MdnSettingsResponseOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) DispositionNotificationTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DispositionNotificationTo
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) MdnText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.MdnText
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) MicHashingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.MicHashingAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) NeedMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.NeedMdn
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) ReceiptDeliveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReceiptDeliveryUrl
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) SendInboundMdnToMessageBox() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendInboundMdnToMessageBox
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) SendMdnAsynchronously() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendMdnAsynchronously
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) SignMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SignMdn
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) SignOutboundMdnIfOptional() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SignOutboundMdnIfOptional
	}).(pulumi.BoolPtrOutput)
}

type AS2MessageConnectionSettings struct {
	IgnoreCertificateNameMismatch *bool `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       *bool `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue *bool `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             *bool `pulumi:"unfoldHttpHeaders"`
}





type AS2MessageConnectionSettingsInput interface {
	pulumi.Input

	ToAS2MessageConnectionSettingsOutput() AS2MessageConnectionSettingsOutput
	ToAS2MessageConnectionSettingsOutputWithContext(context.Context) AS2MessageConnectionSettingsOutput
}

type AS2MessageConnectionSettingsArgs struct {
	IgnoreCertificateNameMismatch pulumi.BoolPtrInput `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       pulumi.BoolPtrInput `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue pulumi.BoolPtrInput `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             pulumi.BoolPtrInput `pulumi:"unfoldHttpHeaders"`
}

func (AS2MessageConnectionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MessageConnectionSettings)(nil)).Elem()
}

func (i AS2MessageConnectionSettingsArgs) ToAS2MessageConnectionSettingsOutput() AS2MessageConnectionSettingsOutput {
	return i.ToAS2MessageConnectionSettingsOutputWithContext(context.Background())
}

func (i AS2MessageConnectionSettingsArgs) ToAS2MessageConnectionSettingsOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MessageConnectionSettingsOutput)
}

func (i AS2MessageConnectionSettingsArgs) ToAS2MessageConnectionSettingsPtrOutput() AS2MessageConnectionSettingsPtrOutput {
	return i.ToAS2MessageConnectionSettingsPtrOutputWithContext(context.Background())
}

func (i AS2MessageConnectionSettingsArgs) ToAS2MessageConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MessageConnectionSettingsOutput).ToAS2MessageConnectionSettingsPtrOutputWithContext(ctx)
}









type AS2MessageConnectionSettingsPtrInput interface {
	pulumi.Input

	ToAS2MessageConnectionSettingsPtrOutput() AS2MessageConnectionSettingsPtrOutput
	ToAS2MessageConnectionSettingsPtrOutputWithContext(context.Context) AS2MessageConnectionSettingsPtrOutput
}

type as2messageConnectionSettingsPtrType AS2MessageConnectionSettingsArgs

func AS2MessageConnectionSettingsPtr(v *AS2MessageConnectionSettingsArgs) AS2MessageConnectionSettingsPtrInput {
	return (*as2messageConnectionSettingsPtrType)(v)
}

func (*as2messageConnectionSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MessageConnectionSettings)(nil)).Elem()
}

func (i *as2messageConnectionSettingsPtrType) ToAS2MessageConnectionSettingsPtrOutput() AS2MessageConnectionSettingsPtrOutput {
	return i.ToAS2MessageConnectionSettingsPtrOutputWithContext(context.Background())
}

func (i *as2messageConnectionSettingsPtrType) ToAS2MessageConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MessageConnectionSettingsPtrOutput)
}

type AS2MessageConnectionSettingsOutput struct{ *pulumi.OutputState }

func (AS2MessageConnectionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MessageConnectionSettings)(nil)).Elem()
}

func (o AS2MessageConnectionSettingsOutput) ToAS2MessageConnectionSettingsOutput() AS2MessageConnectionSettingsOutput {
	return o
}

func (o AS2MessageConnectionSettingsOutput) ToAS2MessageConnectionSettingsOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsOutput {
	return o
}

func (o AS2MessageConnectionSettingsOutput) ToAS2MessageConnectionSettingsPtrOutput() AS2MessageConnectionSettingsPtrOutput {
	return o.ToAS2MessageConnectionSettingsPtrOutputWithContext(context.Background())
}

func (o AS2MessageConnectionSettingsOutput) ToAS2MessageConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2MessageConnectionSettings) *AS2MessageConnectionSettings {
		return &v
	}).(AS2MessageConnectionSettingsPtrOutput)
}

func (o AS2MessageConnectionSettingsOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettings) *bool { return v.IgnoreCertificateNameMismatch }).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettings) *bool { return v.KeepHttpConnectionAlive }).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettings) *bool { return v.SupportHttpStatusCodeContinue }).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettings) *bool { return v.UnfoldHttpHeaders }).(pulumi.BoolPtrOutput)
}

type AS2MessageConnectionSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2MessageConnectionSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MessageConnectionSettings)(nil)).Elem()
}

func (o AS2MessageConnectionSettingsPtrOutput) ToAS2MessageConnectionSettingsPtrOutput() AS2MessageConnectionSettingsPtrOutput {
	return o
}

func (o AS2MessageConnectionSettingsPtrOutput) ToAS2MessageConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsPtrOutput {
	return o
}

func (o AS2MessageConnectionSettingsPtrOutput) Elem() AS2MessageConnectionSettingsOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettings) AS2MessageConnectionSettings {
		if v != nil {
			return *v
		}
		var ret AS2MessageConnectionSettings
		return ret
	}).(AS2MessageConnectionSettingsOutput)
}

func (o AS2MessageConnectionSettingsPtrOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.IgnoreCertificateNameMismatch
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsPtrOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.KeepHttpConnectionAlive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsPtrOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SupportHttpStatusCodeContinue
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsPtrOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.UnfoldHttpHeaders
	}).(pulumi.BoolPtrOutput)
}

type AS2MessageConnectionSettingsResponse struct {
	IgnoreCertificateNameMismatch *bool `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       *bool `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue *bool `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             *bool `pulumi:"unfoldHttpHeaders"`
}

type AS2MessageConnectionSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2MessageConnectionSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MessageConnectionSettingsResponse)(nil)).Elem()
}

func (o AS2MessageConnectionSettingsResponseOutput) ToAS2MessageConnectionSettingsResponseOutput() AS2MessageConnectionSettingsResponseOutput {
	return o
}

func (o AS2MessageConnectionSettingsResponseOutput) ToAS2MessageConnectionSettingsResponseOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsResponseOutput {
	return o
}

func (o AS2MessageConnectionSettingsResponseOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettingsResponse) *bool { return v.IgnoreCertificateNameMismatch }).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsResponseOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettingsResponse) *bool { return v.KeepHttpConnectionAlive }).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsResponseOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettingsResponse) *bool { return v.SupportHttpStatusCodeContinue }).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsResponseOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettingsResponse) *bool { return v.UnfoldHttpHeaders }).(pulumi.BoolPtrOutput)
}

type AS2MessageConnectionSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2MessageConnectionSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MessageConnectionSettingsResponse)(nil)).Elem()
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) ToAS2MessageConnectionSettingsResponsePtrOutput() AS2MessageConnectionSettingsResponsePtrOutput {
	return o
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) ToAS2MessageConnectionSettingsResponsePtrOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsResponsePtrOutput {
	return o
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) Elem() AS2MessageConnectionSettingsResponseOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettingsResponse) AS2MessageConnectionSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2MessageConnectionSettingsResponse
		return ret
	}).(AS2MessageConnectionSettingsResponseOutput)
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IgnoreCertificateNameMismatch
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.KeepHttpConnectionAlive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SupportHttpStatusCodeContinue
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UnfoldHttpHeaders
	}).(pulumi.BoolPtrOutput)
}

type AS2OneWayAgreement struct {
	ProtocolSettings         *AS2ProtocolSettings `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity *BusinessIdentity    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   *BusinessIdentity    `pulumi:"senderBusinessIdentity"`
}





type AS2OneWayAgreementInput interface {
	pulumi.Input

	ToAS2OneWayAgreementOutput() AS2OneWayAgreementOutput
	ToAS2OneWayAgreementOutputWithContext(context.Context) AS2OneWayAgreementOutput
}

type AS2OneWayAgreementArgs struct {
	ProtocolSettings         AS2ProtocolSettingsPtrInput `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentityPtrInput    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentityPtrInput    `pulumi:"senderBusinessIdentity"`
}

func (AS2OneWayAgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2OneWayAgreement)(nil)).Elem()
}

func (i AS2OneWayAgreementArgs) ToAS2OneWayAgreementOutput() AS2OneWayAgreementOutput {
	return i.ToAS2OneWayAgreementOutputWithContext(context.Background())
}

func (i AS2OneWayAgreementArgs) ToAS2OneWayAgreementOutputWithContext(ctx context.Context) AS2OneWayAgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2OneWayAgreementOutput)
}

func (i AS2OneWayAgreementArgs) ToAS2OneWayAgreementPtrOutput() AS2OneWayAgreementPtrOutput {
	return i.ToAS2OneWayAgreementPtrOutputWithContext(context.Background())
}

func (i AS2OneWayAgreementArgs) ToAS2OneWayAgreementPtrOutputWithContext(ctx context.Context) AS2OneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2OneWayAgreementOutput).ToAS2OneWayAgreementPtrOutputWithContext(ctx)
}









type AS2OneWayAgreementPtrInput interface {
	pulumi.Input

	ToAS2OneWayAgreementPtrOutput() AS2OneWayAgreementPtrOutput
	ToAS2OneWayAgreementPtrOutputWithContext(context.Context) AS2OneWayAgreementPtrOutput
}

type as2oneWayAgreementPtrType AS2OneWayAgreementArgs

func AS2OneWayAgreementPtr(v *AS2OneWayAgreementArgs) AS2OneWayAgreementPtrInput {
	return (*as2oneWayAgreementPtrType)(v)
}

func (*as2oneWayAgreementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2OneWayAgreement)(nil)).Elem()
}

func (i *as2oneWayAgreementPtrType) ToAS2OneWayAgreementPtrOutput() AS2OneWayAgreementPtrOutput {
	return i.ToAS2OneWayAgreementPtrOutputWithContext(context.Background())
}

func (i *as2oneWayAgreementPtrType) ToAS2OneWayAgreementPtrOutputWithContext(ctx context.Context) AS2OneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2OneWayAgreementPtrOutput)
}

type AS2OneWayAgreementOutput struct{ *pulumi.OutputState }

func (AS2OneWayAgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2OneWayAgreement)(nil)).Elem()
}

func (o AS2OneWayAgreementOutput) ToAS2OneWayAgreementOutput() AS2OneWayAgreementOutput {
	return o
}

func (o AS2OneWayAgreementOutput) ToAS2OneWayAgreementOutputWithContext(ctx context.Context) AS2OneWayAgreementOutput {
	return o
}

func (o AS2OneWayAgreementOutput) ToAS2OneWayAgreementPtrOutput() AS2OneWayAgreementPtrOutput {
	return o.ToAS2OneWayAgreementPtrOutputWithContext(context.Background())
}

func (o AS2OneWayAgreementOutput) ToAS2OneWayAgreementPtrOutputWithContext(ctx context.Context) AS2OneWayAgreementPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2OneWayAgreement) *AS2OneWayAgreement {
		return &v
	}).(AS2OneWayAgreementPtrOutput)
}

func (o AS2OneWayAgreementOutput) ProtocolSettings() AS2ProtocolSettingsPtrOutput {
	return o.ApplyT(func(v AS2OneWayAgreement) *AS2ProtocolSettings { return v.ProtocolSettings }).(AS2ProtocolSettingsPtrOutput)
}

func (o AS2OneWayAgreementOutput) ReceiverBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v AS2OneWayAgreement) *BusinessIdentity { return v.ReceiverBusinessIdentity }).(BusinessIdentityPtrOutput)
}

func (o AS2OneWayAgreementOutput) SenderBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v AS2OneWayAgreement) *BusinessIdentity { return v.SenderBusinessIdentity }).(BusinessIdentityPtrOutput)
}

type AS2OneWayAgreementPtrOutput struct{ *pulumi.OutputState }

func (AS2OneWayAgreementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2OneWayAgreement)(nil)).Elem()
}

func (o AS2OneWayAgreementPtrOutput) ToAS2OneWayAgreementPtrOutput() AS2OneWayAgreementPtrOutput {
	return o
}

func (o AS2OneWayAgreementPtrOutput) ToAS2OneWayAgreementPtrOutputWithContext(ctx context.Context) AS2OneWayAgreementPtrOutput {
	return o
}

func (o AS2OneWayAgreementPtrOutput) Elem() AS2OneWayAgreementOutput {
	return o.ApplyT(func(v *AS2OneWayAgreement) AS2OneWayAgreement {
		if v != nil {
			return *v
		}
		var ret AS2OneWayAgreement
		return ret
	}).(AS2OneWayAgreementOutput)
}

func (o AS2OneWayAgreementPtrOutput) ProtocolSettings() AS2ProtocolSettingsPtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreement) *AS2ProtocolSettings {
		if v == nil {
			return nil
		}
		return v.ProtocolSettings
	}).(AS2ProtocolSettingsPtrOutput)
}

func (o AS2OneWayAgreementPtrOutput) ReceiverBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return v.ReceiverBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

func (o AS2OneWayAgreementPtrOutput) SenderBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return v.SenderBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

type AS2OneWayAgreementResponse struct {
	ProtocolSettings         *AS2ProtocolSettingsResponse `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity *BusinessIdentityResponse    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   *BusinessIdentityResponse    `pulumi:"senderBusinessIdentity"`
}

type AS2OneWayAgreementResponseOutput struct{ *pulumi.OutputState }

func (AS2OneWayAgreementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2OneWayAgreementResponse)(nil)).Elem()
}

func (o AS2OneWayAgreementResponseOutput) ToAS2OneWayAgreementResponseOutput() AS2OneWayAgreementResponseOutput {
	return o
}

func (o AS2OneWayAgreementResponseOutput) ToAS2OneWayAgreementResponseOutputWithContext(ctx context.Context) AS2OneWayAgreementResponseOutput {
	return o
}

func (o AS2OneWayAgreementResponseOutput) ProtocolSettings() AS2ProtocolSettingsResponsePtrOutput {
	return o.ApplyT(func(v AS2OneWayAgreementResponse) *AS2ProtocolSettingsResponse { return v.ProtocolSettings }).(AS2ProtocolSettingsResponsePtrOutput)
}

func (o AS2OneWayAgreementResponseOutput) ReceiverBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v AS2OneWayAgreementResponse) *BusinessIdentityResponse { return v.ReceiverBusinessIdentity }).(BusinessIdentityResponsePtrOutput)
}

func (o AS2OneWayAgreementResponseOutput) SenderBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v AS2OneWayAgreementResponse) *BusinessIdentityResponse { return v.SenderBusinessIdentity }).(BusinessIdentityResponsePtrOutput)
}

type AS2OneWayAgreementResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2OneWayAgreementResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2OneWayAgreementResponse)(nil)).Elem()
}

func (o AS2OneWayAgreementResponsePtrOutput) ToAS2OneWayAgreementResponsePtrOutput() AS2OneWayAgreementResponsePtrOutput {
	return o
}

func (o AS2OneWayAgreementResponsePtrOutput) ToAS2OneWayAgreementResponsePtrOutputWithContext(ctx context.Context) AS2OneWayAgreementResponsePtrOutput {
	return o
}

func (o AS2OneWayAgreementResponsePtrOutput) Elem() AS2OneWayAgreementResponseOutput {
	return o.ApplyT(func(v *AS2OneWayAgreementResponse) AS2OneWayAgreementResponse {
		if v != nil {
			return *v
		}
		var ret AS2OneWayAgreementResponse
		return ret
	}).(AS2OneWayAgreementResponseOutput)
}

func (o AS2OneWayAgreementResponsePtrOutput) ProtocolSettings() AS2ProtocolSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreementResponse) *AS2ProtocolSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ProtocolSettings
	}).(AS2ProtocolSettingsResponsePtrOutput)
}

func (o AS2OneWayAgreementResponsePtrOutput) ReceiverBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return v.ReceiverBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

func (o AS2OneWayAgreementResponsePtrOutput) SenderBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return v.SenderBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

type AS2ProtocolSettings struct {
	AcknowledgementConnectionSettings *AS2AcknowledgementConnectionSettings `pulumi:"acknowledgementConnectionSettings"`
	EnvelopeSettings                  *AS2EnvelopeSettings                  `pulumi:"envelopeSettings"`
	ErrorSettings                     *AS2ErrorSettings                     `pulumi:"errorSettings"`
	MdnSettings                       *AS2MdnSettings                       `pulumi:"mdnSettings"`
	MessageConnectionSettings         *AS2MessageConnectionSettings         `pulumi:"messageConnectionSettings"`
	SecuritySettings                  *AS2SecuritySettings                  `pulumi:"securitySettings"`
	ValidationSettings                *AS2ValidationSettings                `pulumi:"validationSettings"`
}





type AS2ProtocolSettingsInput interface {
	pulumi.Input

	ToAS2ProtocolSettingsOutput() AS2ProtocolSettingsOutput
	ToAS2ProtocolSettingsOutputWithContext(context.Context) AS2ProtocolSettingsOutput
}

type AS2ProtocolSettingsArgs struct {
	AcknowledgementConnectionSettings AS2AcknowledgementConnectionSettingsPtrInput `pulumi:"acknowledgementConnectionSettings"`
	EnvelopeSettings                  AS2EnvelopeSettingsPtrInput                  `pulumi:"envelopeSettings"`
	ErrorSettings                     AS2ErrorSettingsPtrInput                     `pulumi:"errorSettings"`
	MdnSettings                       AS2MdnSettingsPtrInput                       `pulumi:"mdnSettings"`
	MessageConnectionSettings         AS2MessageConnectionSettingsPtrInput         `pulumi:"messageConnectionSettings"`
	SecuritySettings                  AS2SecuritySettingsPtrInput                  `pulumi:"securitySettings"`
	ValidationSettings                AS2ValidationSettingsPtrInput                `pulumi:"validationSettings"`
}

func (AS2ProtocolSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ProtocolSettings)(nil)).Elem()
}

func (i AS2ProtocolSettingsArgs) ToAS2ProtocolSettingsOutput() AS2ProtocolSettingsOutput {
	return i.ToAS2ProtocolSettingsOutputWithContext(context.Background())
}

func (i AS2ProtocolSettingsArgs) ToAS2ProtocolSettingsOutputWithContext(ctx context.Context) AS2ProtocolSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ProtocolSettingsOutput)
}

func (i AS2ProtocolSettingsArgs) ToAS2ProtocolSettingsPtrOutput() AS2ProtocolSettingsPtrOutput {
	return i.ToAS2ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i AS2ProtocolSettingsArgs) ToAS2ProtocolSettingsPtrOutputWithContext(ctx context.Context) AS2ProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ProtocolSettingsOutput).ToAS2ProtocolSettingsPtrOutputWithContext(ctx)
}









type AS2ProtocolSettingsPtrInput interface {
	pulumi.Input

	ToAS2ProtocolSettingsPtrOutput() AS2ProtocolSettingsPtrOutput
	ToAS2ProtocolSettingsPtrOutputWithContext(context.Context) AS2ProtocolSettingsPtrOutput
}

type as2protocolSettingsPtrType AS2ProtocolSettingsArgs

func AS2ProtocolSettingsPtr(v *AS2ProtocolSettingsArgs) AS2ProtocolSettingsPtrInput {
	return (*as2protocolSettingsPtrType)(v)
}

func (*as2protocolSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ProtocolSettings)(nil)).Elem()
}

func (i *as2protocolSettingsPtrType) ToAS2ProtocolSettingsPtrOutput() AS2ProtocolSettingsPtrOutput {
	return i.ToAS2ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i *as2protocolSettingsPtrType) ToAS2ProtocolSettingsPtrOutputWithContext(ctx context.Context) AS2ProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ProtocolSettingsPtrOutput)
}

type AS2ProtocolSettingsOutput struct{ *pulumi.OutputState }

func (AS2ProtocolSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ProtocolSettings)(nil)).Elem()
}

func (o AS2ProtocolSettingsOutput) ToAS2ProtocolSettingsOutput() AS2ProtocolSettingsOutput {
	return o
}

func (o AS2ProtocolSettingsOutput) ToAS2ProtocolSettingsOutputWithContext(ctx context.Context) AS2ProtocolSettingsOutput {
	return o
}

func (o AS2ProtocolSettingsOutput) ToAS2ProtocolSettingsPtrOutput() AS2ProtocolSettingsPtrOutput {
	return o.ToAS2ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (o AS2ProtocolSettingsOutput) ToAS2ProtocolSettingsPtrOutputWithContext(ctx context.Context) AS2ProtocolSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2ProtocolSettings) *AS2ProtocolSettings {
		return &v
	}).(AS2ProtocolSettingsPtrOutput)
}

func (o AS2ProtocolSettingsOutput) AcknowledgementConnectionSettings() AS2AcknowledgementConnectionSettingsPtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) *AS2AcknowledgementConnectionSettings {
		return v.AcknowledgementConnectionSettings
	}).(AS2AcknowledgementConnectionSettingsPtrOutput)
}

func (o AS2ProtocolSettingsOutput) EnvelopeSettings() AS2EnvelopeSettingsPtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) *AS2EnvelopeSettings { return v.EnvelopeSettings }).(AS2EnvelopeSettingsPtrOutput)
}

func (o AS2ProtocolSettingsOutput) ErrorSettings() AS2ErrorSettingsPtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) *AS2ErrorSettings { return v.ErrorSettings }).(AS2ErrorSettingsPtrOutput)
}

func (o AS2ProtocolSettingsOutput) MdnSettings() AS2MdnSettingsPtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) *AS2MdnSettings { return v.MdnSettings }).(AS2MdnSettingsPtrOutput)
}

func (o AS2ProtocolSettingsOutput) MessageConnectionSettings() AS2MessageConnectionSettingsPtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) *AS2MessageConnectionSettings { return v.MessageConnectionSettings }).(AS2MessageConnectionSettingsPtrOutput)
}

func (o AS2ProtocolSettingsOutput) SecuritySettings() AS2SecuritySettingsPtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) *AS2SecuritySettings { return v.SecuritySettings }).(AS2SecuritySettingsPtrOutput)
}

func (o AS2ProtocolSettingsOutput) ValidationSettings() AS2ValidationSettingsPtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) *AS2ValidationSettings { return v.ValidationSettings }).(AS2ValidationSettingsPtrOutput)
}

type AS2ProtocolSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2ProtocolSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ProtocolSettings)(nil)).Elem()
}

func (o AS2ProtocolSettingsPtrOutput) ToAS2ProtocolSettingsPtrOutput() AS2ProtocolSettingsPtrOutput {
	return o
}

func (o AS2ProtocolSettingsPtrOutput) ToAS2ProtocolSettingsPtrOutputWithContext(ctx context.Context) AS2ProtocolSettingsPtrOutput {
	return o
}

func (o AS2ProtocolSettingsPtrOutput) Elem() AS2ProtocolSettingsOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) AS2ProtocolSettings {
		if v != nil {
			return *v
		}
		var ret AS2ProtocolSettings
		return ret
	}).(AS2ProtocolSettingsOutput)
}

func (o AS2ProtocolSettingsPtrOutput) AcknowledgementConnectionSettings() AS2AcknowledgementConnectionSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2AcknowledgementConnectionSettings {
		if v == nil {
			return nil
		}
		return v.AcknowledgementConnectionSettings
	}).(AS2AcknowledgementConnectionSettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) EnvelopeSettings() AS2EnvelopeSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2EnvelopeSettings {
		if v == nil {
			return nil
		}
		return v.EnvelopeSettings
	}).(AS2EnvelopeSettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) ErrorSettings() AS2ErrorSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2ErrorSettings {
		if v == nil {
			return nil
		}
		return v.ErrorSettings
	}).(AS2ErrorSettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) MdnSettings() AS2MdnSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2MdnSettings {
		if v == nil {
			return nil
		}
		return v.MdnSettings
	}).(AS2MdnSettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) MessageConnectionSettings() AS2MessageConnectionSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2MessageConnectionSettings {
		if v == nil {
			return nil
		}
		return v.MessageConnectionSettings
	}).(AS2MessageConnectionSettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) SecuritySettings() AS2SecuritySettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2SecuritySettings {
		if v == nil {
			return nil
		}
		return v.SecuritySettings
	}).(AS2SecuritySettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) ValidationSettings() AS2ValidationSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2ValidationSettings {
		if v == nil {
			return nil
		}
		return v.ValidationSettings
	}).(AS2ValidationSettingsPtrOutput)
}

type AS2ProtocolSettingsResponse struct {
	AcknowledgementConnectionSettings *AS2AcknowledgementConnectionSettingsResponse `pulumi:"acknowledgementConnectionSettings"`
	EnvelopeSettings                  *AS2EnvelopeSettingsResponse                  `pulumi:"envelopeSettings"`
	ErrorSettings                     *AS2ErrorSettingsResponse                     `pulumi:"errorSettings"`
	MdnSettings                       *AS2MdnSettingsResponse                       `pulumi:"mdnSettings"`
	MessageConnectionSettings         *AS2MessageConnectionSettingsResponse         `pulumi:"messageConnectionSettings"`
	SecuritySettings                  *AS2SecuritySettingsResponse                  `pulumi:"securitySettings"`
	ValidationSettings                *AS2ValidationSettingsResponse                `pulumi:"validationSettings"`
}

type AS2ProtocolSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2ProtocolSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ProtocolSettingsResponse)(nil)).Elem()
}

func (o AS2ProtocolSettingsResponseOutput) ToAS2ProtocolSettingsResponseOutput() AS2ProtocolSettingsResponseOutput {
	return o
}

func (o AS2ProtocolSettingsResponseOutput) ToAS2ProtocolSettingsResponseOutputWithContext(ctx context.Context) AS2ProtocolSettingsResponseOutput {
	return o
}

func (o AS2ProtocolSettingsResponseOutput) AcknowledgementConnectionSettings() AS2AcknowledgementConnectionSettingsResponsePtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) *AS2AcknowledgementConnectionSettingsResponse {
		return v.AcknowledgementConnectionSettings
	}).(AS2AcknowledgementConnectionSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponseOutput) EnvelopeSettings() AS2EnvelopeSettingsResponsePtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) *AS2EnvelopeSettingsResponse { return v.EnvelopeSettings }).(AS2EnvelopeSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponseOutput) ErrorSettings() AS2ErrorSettingsResponsePtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) *AS2ErrorSettingsResponse { return v.ErrorSettings }).(AS2ErrorSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponseOutput) MdnSettings() AS2MdnSettingsResponsePtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) *AS2MdnSettingsResponse { return v.MdnSettings }).(AS2MdnSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponseOutput) MessageConnectionSettings() AS2MessageConnectionSettingsResponsePtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) *AS2MessageConnectionSettingsResponse {
		return v.MessageConnectionSettings
	}).(AS2MessageConnectionSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponseOutput) SecuritySettings() AS2SecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) *AS2SecuritySettingsResponse { return v.SecuritySettings }).(AS2SecuritySettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponseOutput) ValidationSettings() AS2ValidationSettingsResponsePtrOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) *AS2ValidationSettingsResponse { return v.ValidationSettings }).(AS2ValidationSettingsResponsePtrOutput)
}

type AS2ProtocolSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2ProtocolSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ProtocolSettingsResponse)(nil)).Elem()
}

func (o AS2ProtocolSettingsResponsePtrOutput) ToAS2ProtocolSettingsResponsePtrOutput() AS2ProtocolSettingsResponsePtrOutput {
	return o
}

func (o AS2ProtocolSettingsResponsePtrOutput) ToAS2ProtocolSettingsResponsePtrOutputWithContext(ctx context.Context) AS2ProtocolSettingsResponsePtrOutput {
	return o
}

func (o AS2ProtocolSettingsResponsePtrOutput) Elem() AS2ProtocolSettingsResponseOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) AS2ProtocolSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2ProtocolSettingsResponse
		return ret
	}).(AS2ProtocolSettingsResponseOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) AcknowledgementConnectionSettings() AS2AcknowledgementConnectionSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2AcknowledgementConnectionSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AcknowledgementConnectionSettings
	}).(AS2AcknowledgementConnectionSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) EnvelopeSettings() AS2EnvelopeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2EnvelopeSettingsResponse {
		if v == nil {
			return nil
		}
		return v.EnvelopeSettings
	}).(AS2EnvelopeSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) ErrorSettings() AS2ErrorSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2ErrorSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ErrorSettings
	}).(AS2ErrorSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) MdnSettings() AS2MdnSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2MdnSettingsResponse {
		if v == nil {
			return nil
		}
		return v.MdnSettings
	}).(AS2MdnSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) MessageConnectionSettings() AS2MessageConnectionSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2MessageConnectionSettingsResponse {
		if v == nil {
			return nil
		}
		return v.MessageConnectionSettings
	}).(AS2MessageConnectionSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) SecuritySettings() AS2SecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2SecuritySettingsResponse {
		if v == nil {
			return nil
		}
		return v.SecuritySettings
	}).(AS2SecuritySettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) ValidationSettings() AS2ValidationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2ValidationSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ValidationSettings
	}).(AS2ValidationSettingsResponsePtrOutput)
}

type AS2SecuritySettings struct {
	EnableNrrForInboundDecodedMessages  *bool   `pulumi:"enableNrrForInboundDecodedMessages"`
	EnableNrrForInboundEncodedMessages  *bool   `pulumi:"enableNrrForInboundEncodedMessages"`
	EnableNrrForInboundMdn              *bool   `pulumi:"enableNrrForInboundMdn"`
	EnableNrrForOutboundDecodedMessages *bool   `pulumi:"enableNrrForOutboundDecodedMessages"`
	EnableNrrForOutboundEncodedMessages *bool   `pulumi:"enableNrrForOutboundEncodedMessages"`
	EnableNrrForOutboundMdn             *bool   `pulumi:"enableNrrForOutboundMdn"`
	EncryptionCertificateName           *string `pulumi:"encryptionCertificateName"`
	OverrideGroupSigningCertificate     *bool   `pulumi:"overrideGroupSigningCertificate"`
	SigningCertificateName              *string `pulumi:"signingCertificateName"`
}





type AS2SecuritySettingsInput interface {
	pulumi.Input

	ToAS2SecuritySettingsOutput() AS2SecuritySettingsOutput
	ToAS2SecuritySettingsOutputWithContext(context.Context) AS2SecuritySettingsOutput
}

type AS2SecuritySettingsArgs struct {
	EnableNrrForInboundDecodedMessages  pulumi.BoolPtrInput   `pulumi:"enableNrrForInboundDecodedMessages"`
	EnableNrrForInboundEncodedMessages  pulumi.BoolPtrInput   `pulumi:"enableNrrForInboundEncodedMessages"`
	EnableNrrForInboundMdn              pulumi.BoolPtrInput   `pulumi:"enableNrrForInboundMdn"`
	EnableNrrForOutboundDecodedMessages pulumi.BoolPtrInput   `pulumi:"enableNrrForOutboundDecodedMessages"`
	EnableNrrForOutboundEncodedMessages pulumi.BoolPtrInput   `pulumi:"enableNrrForOutboundEncodedMessages"`
	EnableNrrForOutboundMdn             pulumi.BoolPtrInput   `pulumi:"enableNrrForOutboundMdn"`
	EncryptionCertificateName           pulumi.StringPtrInput `pulumi:"encryptionCertificateName"`
	OverrideGroupSigningCertificate     pulumi.BoolPtrInput   `pulumi:"overrideGroupSigningCertificate"`
	SigningCertificateName              pulumi.StringPtrInput `pulumi:"signingCertificateName"`
}

func (AS2SecuritySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2SecuritySettings)(nil)).Elem()
}

func (i AS2SecuritySettingsArgs) ToAS2SecuritySettingsOutput() AS2SecuritySettingsOutput {
	return i.ToAS2SecuritySettingsOutputWithContext(context.Background())
}

func (i AS2SecuritySettingsArgs) ToAS2SecuritySettingsOutputWithContext(ctx context.Context) AS2SecuritySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2SecuritySettingsOutput)
}

func (i AS2SecuritySettingsArgs) ToAS2SecuritySettingsPtrOutput() AS2SecuritySettingsPtrOutput {
	return i.ToAS2SecuritySettingsPtrOutputWithContext(context.Background())
}

func (i AS2SecuritySettingsArgs) ToAS2SecuritySettingsPtrOutputWithContext(ctx context.Context) AS2SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2SecuritySettingsOutput).ToAS2SecuritySettingsPtrOutputWithContext(ctx)
}









type AS2SecuritySettingsPtrInput interface {
	pulumi.Input

	ToAS2SecuritySettingsPtrOutput() AS2SecuritySettingsPtrOutput
	ToAS2SecuritySettingsPtrOutputWithContext(context.Context) AS2SecuritySettingsPtrOutput
}

type as2securitySettingsPtrType AS2SecuritySettingsArgs

func AS2SecuritySettingsPtr(v *AS2SecuritySettingsArgs) AS2SecuritySettingsPtrInput {
	return (*as2securitySettingsPtrType)(v)
}

func (*as2securitySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2SecuritySettings)(nil)).Elem()
}

func (i *as2securitySettingsPtrType) ToAS2SecuritySettingsPtrOutput() AS2SecuritySettingsPtrOutput {
	return i.ToAS2SecuritySettingsPtrOutputWithContext(context.Background())
}

func (i *as2securitySettingsPtrType) ToAS2SecuritySettingsPtrOutputWithContext(ctx context.Context) AS2SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2SecuritySettingsPtrOutput)
}

type AS2SecuritySettingsOutput struct{ *pulumi.OutputState }

func (AS2SecuritySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2SecuritySettings)(nil)).Elem()
}

func (o AS2SecuritySettingsOutput) ToAS2SecuritySettingsOutput() AS2SecuritySettingsOutput {
	return o
}

func (o AS2SecuritySettingsOutput) ToAS2SecuritySettingsOutputWithContext(ctx context.Context) AS2SecuritySettingsOutput {
	return o
}

func (o AS2SecuritySettingsOutput) ToAS2SecuritySettingsPtrOutput() AS2SecuritySettingsPtrOutput {
	return o.ToAS2SecuritySettingsPtrOutputWithContext(context.Background())
}

func (o AS2SecuritySettingsOutput) ToAS2SecuritySettingsPtrOutputWithContext(ctx context.Context) AS2SecuritySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2SecuritySettings) *AS2SecuritySettings {
		return &v
	}).(AS2SecuritySettingsPtrOutput)
}

func (o AS2SecuritySettingsOutput) EnableNrrForInboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *bool { return v.EnableNrrForInboundDecodedMessages }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsOutput) EnableNrrForInboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *bool { return v.EnableNrrForInboundEncodedMessages }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsOutput) EnableNrrForInboundMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *bool { return v.EnableNrrForInboundMdn }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsOutput) EnableNrrForOutboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *bool { return v.EnableNrrForOutboundDecodedMessages }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsOutput) EnableNrrForOutboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *bool { return v.EnableNrrForOutboundEncodedMessages }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsOutput) EnableNrrForOutboundMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *bool { return v.EnableNrrForOutboundMdn }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsOutput) EncryptionCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *string { return v.EncryptionCertificateName }).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsOutput) OverrideGroupSigningCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *bool { return v.OverrideGroupSigningCertificate }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsOutput) SigningCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *string { return v.SigningCertificateName }).(pulumi.StringPtrOutput)
}

type AS2SecuritySettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2SecuritySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2SecuritySettings)(nil)).Elem()
}

func (o AS2SecuritySettingsPtrOutput) ToAS2SecuritySettingsPtrOutput() AS2SecuritySettingsPtrOutput {
	return o
}

func (o AS2SecuritySettingsPtrOutput) ToAS2SecuritySettingsPtrOutputWithContext(ctx context.Context) AS2SecuritySettingsPtrOutput {
	return o
}

func (o AS2SecuritySettingsPtrOutput) Elem() AS2SecuritySettingsOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) AS2SecuritySettings {
		if v != nil {
			return *v
		}
		var ret AS2SecuritySettings
		return ret
	}).(AS2SecuritySettingsOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNrrForInboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForInboundDecodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNrrForInboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForInboundEncodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNrrForInboundMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForInboundMdn
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNrrForOutboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForOutboundDecodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNrrForOutboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForOutboundEncodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNrrForOutboundMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForOutboundMdn
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EncryptionCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionCertificateName
	}).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) OverrideGroupSigningCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return v.OverrideGroupSigningCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) SigningCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.SigningCertificateName
	}).(pulumi.StringPtrOutput)
}

type AS2SecuritySettingsResponse struct {
	EnableNrrForInboundDecodedMessages  *bool   `pulumi:"enableNrrForInboundDecodedMessages"`
	EnableNrrForInboundEncodedMessages  *bool   `pulumi:"enableNrrForInboundEncodedMessages"`
	EnableNrrForInboundMdn              *bool   `pulumi:"enableNrrForInboundMdn"`
	EnableNrrForOutboundDecodedMessages *bool   `pulumi:"enableNrrForOutboundDecodedMessages"`
	EnableNrrForOutboundEncodedMessages *bool   `pulumi:"enableNrrForOutboundEncodedMessages"`
	EnableNrrForOutboundMdn             *bool   `pulumi:"enableNrrForOutboundMdn"`
	EncryptionCertificateName           *string `pulumi:"encryptionCertificateName"`
	OverrideGroupSigningCertificate     *bool   `pulumi:"overrideGroupSigningCertificate"`
	SigningCertificateName              *string `pulumi:"signingCertificateName"`
}

type AS2SecuritySettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2SecuritySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2SecuritySettingsResponse)(nil)).Elem()
}

func (o AS2SecuritySettingsResponseOutput) ToAS2SecuritySettingsResponseOutput() AS2SecuritySettingsResponseOutput {
	return o
}

func (o AS2SecuritySettingsResponseOutput) ToAS2SecuritySettingsResponseOutputWithContext(ctx context.Context) AS2SecuritySettingsResponseOutput {
	return o
}

func (o AS2SecuritySettingsResponseOutput) EnableNrrForInboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *bool { return v.EnableNrrForInboundDecodedMessages }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponseOutput) EnableNrrForInboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *bool { return v.EnableNrrForInboundEncodedMessages }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponseOutput) EnableNrrForInboundMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *bool { return v.EnableNrrForInboundMdn }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponseOutput) EnableNrrForOutboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *bool { return v.EnableNrrForOutboundDecodedMessages }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponseOutput) EnableNrrForOutboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *bool { return v.EnableNrrForOutboundEncodedMessages }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponseOutput) EnableNrrForOutboundMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *bool { return v.EnableNrrForOutboundMdn }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponseOutput) EncryptionCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *string { return v.EncryptionCertificateName }).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsResponseOutput) OverrideGroupSigningCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *bool { return v.OverrideGroupSigningCertificate }).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponseOutput) SigningCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *string { return v.SigningCertificateName }).(pulumi.StringPtrOutput)
}

type AS2SecuritySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2SecuritySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2SecuritySettingsResponse)(nil)).Elem()
}

func (o AS2SecuritySettingsResponsePtrOutput) ToAS2SecuritySettingsResponsePtrOutput() AS2SecuritySettingsResponsePtrOutput {
	return o
}

func (o AS2SecuritySettingsResponsePtrOutput) ToAS2SecuritySettingsResponsePtrOutputWithContext(ctx context.Context) AS2SecuritySettingsResponsePtrOutput {
	return o
}

func (o AS2SecuritySettingsResponsePtrOutput) Elem() AS2SecuritySettingsResponseOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) AS2SecuritySettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2SecuritySettingsResponse
		return ret
	}).(AS2SecuritySettingsResponseOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNrrForInboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForInboundDecodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNrrForInboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForInboundEncodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNrrForInboundMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForInboundMdn
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNrrForOutboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForOutboundDecodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNrrForOutboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForOutboundEncodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNrrForOutboundMdn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNrrForOutboundMdn
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EncryptionCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionCertificateName
	}).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) OverrideGroupSigningCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.OverrideGroupSigningCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) SigningCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SigningCertificateName
	}).(pulumi.StringPtrOutput)
}

type AS2ValidationSettings struct {
	CheckCertificateRevocationListOnReceive *bool                `pulumi:"checkCertificateRevocationListOnReceive"`
	CheckCertificateRevocationListOnSend    *bool                `pulumi:"checkCertificateRevocationListOnSend"`
	CheckDuplicateMessage                   *bool                `pulumi:"checkDuplicateMessage"`
	CompressMessage                         *bool                `pulumi:"compressMessage"`
	EncryptMessage                          *bool                `pulumi:"encryptMessage"`
	EncryptionAlgorithm                     *EncryptionAlgorithm `pulumi:"encryptionAlgorithm"`
	InterchangeDuplicatesValidityDays       *int                 `pulumi:"interchangeDuplicatesValidityDays"`
	OverrideMessageProperties               *bool                `pulumi:"overrideMessageProperties"`
	SignMessage                             *bool                `pulumi:"signMessage"`
}





type AS2ValidationSettingsInput interface {
	pulumi.Input

	ToAS2ValidationSettingsOutput() AS2ValidationSettingsOutput
	ToAS2ValidationSettingsOutputWithContext(context.Context) AS2ValidationSettingsOutput
}

type AS2ValidationSettingsArgs struct {
	CheckCertificateRevocationListOnReceive pulumi.BoolPtrInput         `pulumi:"checkCertificateRevocationListOnReceive"`
	CheckCertificateRevocationListOnSend    pulumi.BoolPtrInput         `pulumi:"checkCertificateRevocationListOnSend"`
	CheckDuplicateMessage                   pulumi.BoolPtrInput         `pulumi:"checkDuplicateMessage"`
	CompressMessage                         pulumi.BoolPtrInput         `pulumi:"compressMessage"`
	EncryptMessage                          pulumi.BoolPtrInput         `pulumi:"encryptMessage"`
	EncryptionAlgorithm                     EncryptionAlgorithmPtrInput `pulumi:"encryptionAlgorithm"`
	InterchangeDuplicatesValidityDays       pulumi.IntPtrInput          `pulumi:"interchangeDuplicatesValidityDays"`
	OverrideMessageProperties               pulumi.BoolPtrInput         `pulumi:"overrideMessageProperties"`
	SignMessage                             pulumi.BoolPtrInput         `pulumi:"signMessage"`
}

func (AS2ValidationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ValidationSettings)(nil)).Elem()
}

func (i AS2ValidationSettingsArgs) ToAS2ValidationSettingsOutput() AS2ValidationSettingsOutput {
	return i.ToAS2ValidationSettingsOutputWithContext(context.Background())
}

func (i AS2ValidationSettingsArgs) ToAS2ValidationSettingsOutputWithContext(ctx context.Context) AS2ValidationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ValidationSettingsOutput)
}

func (i AS2ValidationSettingsArgs) ToAS2ValidationSettingsPtrOutput() AS2ValidationSettingsPtrOutput {
	return i.ToAS2ValidationSettingsPtrOutputWithContext(context.Background())
}

func (i AS2ValidationSettingsArgs) ToAS2ValidationSettingsPtrOutputWithContext(ctx context.Context) AS2ValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ValidationSettingsOutput).ToAS2ValidationSettingsPtrOutputWithContext(ctx)
}









type AS2ValidationSettingsPtrInput interface {
	pulumi.Input

	ToAS2ValidationSettingsPtrOutput() AS2ValidationSettingsPtrOutput
	ToAS2ValidationSettingsPtrOutputWithContext(context.Context) AS2ValidationSettingsPtrOutput
}

type as2validationSettingsPtrType AS2ValidationSettingsArgs

func AS2ValidationSettingsPtr(v *AS2ValidationSettingsArgs) AS2ValidationSettingsPtrInput {
	return (*as2validationSettingsPtrType)(v)
}

func (*as2validationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ValidationSettings)(nil)).Elem()
}

func (i *as2validationSettingsPtrType) ToAS2ValidationSettingsPtrOutput() AS2ValidationSettingsPtrOutput {
	return i.ToAS2ValidationSettingsPtrOutputWithContext(context.Background())
}

func (i *as2validationSettingsPtrType) ToAS2ValidationSettingsPtrOutputWithContext(ctx context.Context) AS2ValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ValidationSettingsPtrOutput)
}

type AS2ValidationSettingsOutput struct{ *pulumi.OutputState }

func (AS2ValidationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ValidationSettings)(nil)).Elem()
}

func (o AS2ValidationSettingsOutput) ToAS2ValidationSettingsOutput() AS2ValidationSettingsOutput {
	return o
}

func (o AS2ValidationSettingsOutput) ToAS2ValidationSettingsOutputWithContext(ctx context.Context) AS2ValidationSettingsOutput {
	return o
}

func (o AS2ValidationSettingsOutput) ToAS2ValidationSettingsPtrOutput() AS2ValidationSettingsPtrOutput {
	return o.ToAS2ValidationSettingsPtrOutputWithContext(context.Background())
}

func (o AS2ValidationSettingsOutput) ToAS2ValidationSettingsPtrOutputWithContext(ctx context.Context) AS2ValidationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2ValidationSettings) *AS2ValidationSettings {
		return &v
	}).(AS2ValidationSettingsPtrOutput)
}

func (o AS2ValidationSettingsOutput) CheckCertificateRevocationListOnReceive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettings) *bool { return v.CheckCertificateRevocationListOnReceive }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsOutput) CheckCertificateRevocationListOnSend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettings) *bool { return v.CheckCertificateRevocationListOnSend }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsOutput) CheckDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettings) *bool { return v.CheckDuplicateMessage }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsOutput) CompressMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettings) *bool { return v.CompressMessage }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsOutput) EncryptMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettings) *bool { return v.EncryptMessage }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsOutput) EncryptionAlgorithm() EncryptionAlgorithmPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettings) *EncryptionAlgorithm { return v.EncryptionAlgorithm }).(EncryptionAlgorithmPtrOutput)
}

func (o AS2ValidationSettingsOutput) InterchangeDuplicatesValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettings) *int { return v.InterchangeDuplicatesValidityDays }).(pulumi.IntPtrOutput)
}

func (o AS2ValidationSettingsOutput) OverrideMessageProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettings) *bool { return v.OverrideMessageProperties }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsOutput) SignMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettings) *bool { return v.SignMessage }).(pulumi.BoolPtrOutput)
}

type AS2ValidationSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2ValidationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ValidationSettings)(nil)).Elem()
}

func (o AS2ValidationSettingsPtrOutput) ToAS2ValidationSettingsPtrOutput() AS2ValidationSettingsPtrOutput {
	return o
}

func (o AS2ValidationSettingsPtrOutput) ToAS2ValidationSettingsPtrOutputWithContext(ctx context.Context) AS2ValidationSettingsPtrOutput {
	return o
}

func (o AS2ValidationSettingsPtrOutput) Elem() AS2ValidationSettingsOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) AS2ValidationSettings {
		if v != nil {
			return *v
		}
		var ret AS2ValidationSettings
		return ret
	}).(AS2ValidationSettingsOutput)
}

func (o AS2ValidationSettingsPtrOutput) CheckCertificateRevocationListOnReceive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CheckCertificateRevocationListOnReceive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) CheckCertificateRevocationListOnSend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CheckCertificateRevocationListOnSend
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) CheckDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) CompressMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CompressMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) EncryptMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) EncryptionAlgorithm() EncryptionAlgorithmPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *EncryptionAlgorithm {
		if v == nil {
			return nil
		}
		return v.EncryptionAlgorithm
	}).(EncryptionAlgorithmPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) InterchangeDuplicatesValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *int {
		if v == nil {
			return nil
		}
		return v.InterchangeDuplicatesValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) OverrideMessageProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.OverrideMessageProperties
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) SignMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SignMessage
	}).(pulumi.BoolPtrOutput)
}

type AS2ValidationSettingsResponse struct {
	CheckCertificateRevocationListOnReceive *bool   `pulumi:"checkCertificateRevocationListOnReceive"`
	CheckCertificateRevocationListOnSend    *bool   `pulumi:"checkCertificateRevocationListOnSend"`
	CheckDuplicateMessage                   *bool   `pulumi:"checkDuplicateMessage"`
	CompressMessage                         *bool   `pulumi:"compressMessage"`
	EncryptMessage                          *bool   `pulumi:"encryptMessage"`
	EncryptionAlgorithm                     *string `pulumi:"encryptionAlgorithm"`
	InterchangeDuplicatesValidityDays       *int    `pulumi:"interchangeDuplicatesValidityDays"`
	OverrideMessageProperties               *bool   `pulumi:"overrideMessageProperties"`
	SignMessage                             *bool   `pulumi:"signMessage"`
}

type AS2ValidationSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2ValidationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ValidationSettingsResponse)(nil)).Elem()
}

func (o AS2ValidationSettingsResponseOutput) ToAS2ValidationSettingsResponseOutput() AS2ValidationSettingsResponseOutput {
	return o
}

func (o AS2ValidationSettingsResponseOutput) ToAS2ValidationSettingsResponseOutputWithContext(ctx context.Context) AS2ValidationSettingsResponseOutput {
	return o
}

func (o AS2ValidationSettingsResponseOutput) CheckCertificateRevocationListOnReceive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) *bool { return v.CheckCertificateRevocationListOnReceive }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponseOutput) CheckCertificateRevocationListOnSend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) *bool { return v.CheckCertificateRevocationListOnSend }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponseOutput) CheckDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) *bool { return v.CheckDuplicateMessage }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponseOutput) CompressMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) *bool { return v.CompressMessage }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponseOutput) EncryptMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) *bool { return v.EncryptMessage }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponseOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) *string { return v.EncryptionAlgorithm }).(pulumi.StringPtrOutput)
}

func (o AS2ValidationSettingsResponseOutput) InterchangeDuplicatesValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) *int { return v.InterchangeDuplicatesValidityDays }).(pulumi.IntPtrOutput)
}

func (o AS2ValidationSettingsResponseOutput) OverrideMessageProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) *bool { return v.OverrideMessageProperties }).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponseOutput) SignMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) *bool { return v.SignMessage }).(pulumi.BoolPtrOutput)
}

type AS2ValidationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2ValidationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ValidationSettingsResponse)(nil)).Elem()
}

func (o AS2ValidationSettingsResponsePtrOutput) ToAS2ValidationSettingsResponsePtrOutput() AS2ValidationSettingsResponsePtrOutput {
	return o
}

func (o AS2ValidationSettingsResponsePtrOutput) ToAS2ValidationSettingsResponsePtrOutputWithContext(ctx context.Context) AS2ValidationSettingsResponsePtrOutput {
	return o
}

func (o AS2ValidationSettingsResponsePtrOutput) Elem() AS2ValidationSettingsResponseOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) AS2ValidationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2ValidationSettingsResponse
		return ret
	}).(AS2ValidationSettingsResponseOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) CheckCertificateRevocationListOnReceive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CheckCertificateRevocationListOnReceive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) CheckCertificateRevocationListOnSend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CheckCertificateRevocationListOnSend
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) CheckDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) CompressMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CompressMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) EncryptMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) InterchangeDuplicatesValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.InterchangeDuplicatesValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) OverrideMessageProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.OverrideMessageProperties
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) SignMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SignMessage
	}).(pulumi.BoolPtrOutput)
}

type AgreementContent struct {
	AS2     *AS2AgreementContent     `pulumi:"aS2"`
	Edifact *EdifactAgreementContent `pulumi:"edifact"`
	X12     *X12AgreementContent     `pulumi:"x12"`
}





type AgreementContentInput interface {
	pulumi.Input

	ToAgreementContentOutput() AgreementContentOutput
	ToAgreementContentOutputWithContext(context.Context) AgreementContentOutput
}

type AgreementContentArgs struct {
	AS2     AS2AgreementContentPtrInput     `pulumi:"aS2"`
	Edifact EdifactAgreementContentPtrInput `pulumi:"edifact"`
	X12     X12AgreementContentPtrInput     `pulumi:"x12"`
}

func (AgreementContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgreementContent)(nil)).Elem()
}

func (i AgreementContentArgs) ToAgreementContentOutput() AgreementContentOutput {
	return i.ToAgreementContentOutputWithContext(context.Background())
}

func (i AgreementContentArgs) ToAgreementContentOutputWithContext(ctx context.Context) AgreementContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgreementContentOutput)
}

func (i AgreementContentArgs) ToAgreementContentPtrOutput() AgreementContentPtrOutput {
	return i.ToAgreementContentPtrOutputWithContext(context.Background())
}

func (i AgreementContentArgs) ToAgreementContentPtrOutputWithContext(ctx context.Context) AgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgreementContentOutput).ToAgreementContentPtrOutputWithContext(ctx)
}









type AgreementContentPtrInput interface {
	pulumi.Input

	ToAgreementContentPtrOutput() AgreementContentPtrOutput
	ToAgreementContentPtrOutputWithContext(context.Context) AgreementContentPtrOutput
}

type agreementContentPtrType AgreementContentArgs

func AgreementContentPtr(v *AgreementContentArgs) AgreementContentPtrInput {
	return (*agreementContentPtrType)(v)
}

func (*agreementContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgreementContent)(nil)).Elem()
}

func (i *agreementContentPtrType) ToAgreementContentPtrOutput() AgreementContentPtrOutput {
	return i.ToAgreementContentPtrOutputWithContext(context.Background())
}

func (i *agreementContentPtrType) ToAgreementContentPtrOutputWithContext(ctx context.Context) AgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgreementContentPtrOutput)
}

type AgreementContentOutput struct{ *pulumi.OutputState }

func (AgreementContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgreementContent)(nil)).Elem()
}

func (o AgreementContentOutput) ToAgreementContentOutput() AgreementContentOutput {
	return o
}

func (o AgreementContentOutput) ToAgreementContentOutputWithContext(ctx context.Context) AgreementContentOutput {
	return o
}

func (o AgreementContentOutput) ToAgreementContentPtrOutput() AgreementContentPtrOutput {
	return o.ToAgreementContentPtrOutputWithContext(context.Background())
}

func (o AgreementContentOutput) ToAgreementContentPtrOutputWithContext(ctx context.Context) AgreementContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgreementContent) *AgreementContent {
		return &v
	}).(AgreementContentPtrOutput)
}

func (o AgreementContentOutput) AS2() AS2AgreementContentPtrOutput {
	return o.ApplyT(func(v AgreementContent) *AS2AgreementContent { return v.AS2 }).(AS2AgreementContentPtrOutput)
}

func (o AgreementContentOutput) Edifact() EdifactAgreementContentPtrOutput {
	return o.ApplyT(func(v AgreementContent) *EdifactAgreementContent { return v.Edifact }).(EdifactAgreementContentPtrOutput)
}

func (o AgreementContentOutput) X12() X12AgreementContentPtrOutput {
	return o.ApplyT(func(v AgreementContent) *X12AgreementContent { return v.X12 }).(X12AgreementContentPtrOutput)
}

type AgreementContentPtrOutput struct{ *pulumi.OutputState }

func (AgreementContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgreementContent)(nil)).Elem()
}

func (o AgreementContentPtrOutput) ToAgreementContentPtrOutput() AgreementContentPtrOutput {
	return o
}

func (o AgreementContentPtrOutput) ToAgreementContentPtrOutputWithContext(ctx context.Context) AgreementContentPtrOutput {
	return o
}

func (o AgreementContentPtrOutput) Elem() AgreementContentOutput {
	return o.ApplyT(func(v *AgreementContent) AgreementContent {
		if v != nil {
			return *v
		}
		var ret AgreementContent
		return ret
	}).(AgreementContentOutput)
}

func (o AgreementContentPtrOutput) AS2() AS2AgreementContentPtrOutput {
	return o.ApplyT(func(v *AgreementContent) *AS2AgreementContent {
		if v == nil {
			return nil
		}
		return v.AS2
	}).(AS2AgreementContentPtrOutput)
}

func (o AgreementContentPtrOutput) Edifact() EdifactAgreementContentPtrOutput {
	return o.ApplyT(func(v *AgreementContent) *EdifactAgreementContent {
		if v == nil {
			return nil
		}
		return v.Edifact
	}).(EdifactAgreementContentPtrOutput)
}

func (o AgreementContentPtrOutput) X12() X12AgreementContentPtrOutput {
	return o.ApplyT(func(v *AgreementContent) *X12AgreementContent {
		if v == nil {
			return nil
		}
		return v.X12
	}).(X12AgreementContentPtrOutput)
}

type AgreementContentResponse struct {
	AS2     *AS2AgreementContentResponse     `pulumi:"aS2"`
	Edifact *EdifactAgreementContentResponse `pulumi:"edifact"`
	X12     *X12AgreementContentResponse     `pulumi:"x12"`
}

type AgreementContentResponseOutput struct{ *pulumi.OutputState }

func (AgreementContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgreementContentResponse)(nil)).Elem()
}

func (o AgreementContentResponseOutput) ToAgreementContentResponseOutput() AgreementContentResponseOutput {
	return o
}

func (o AgreementContentResponseOutput) ToAgreementContentResponseOutputWithContext(ctx context.Context) AgreementContentResponseOutput {
	return o
}

func (o AgreementContentResponseOutput) AS2() AS2AgreementContentResponsePtrOutput {
	return o.ApplyT(func(v AgreementContentResponse) *AS2AgreementContentResponse { return v.AS2 }).(AS2AgreementContentResponsePtrOutput)
}

func (o AgreementContentResponseOutput) Edifact() EdifactAgreementContentResponsePtrOutput {
	return o.ApplyT(func(v AgreementContentResponse) *EdifactAgreementContentResponse { return v.Edifact }).(EdifactAgreementContentResponsePtrOutput)
}

func (o AgreementContentResponseOutput) X12() X12AgreementContentResponsePtrOutput {
	return o.ApplyT(func(v AgreementContentResponse) *X12AgreementContentResponse { return v.X12 }).(X12AgreementContentResponsePtrOutput)
}

type AgreementContentResponsePtrOutput struct{ *pulumi.OutputState }

func (AgreementContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgreementContentResponse)(nil)).Elem()
}

func (o AgreementContentResponsePtrOutput) ToAgreementContentResponsePtrOutput() AgreementContentResponsePtrOutput {
	return o
}

func (o AgreementContentResponsePtrOutput) ToAgreementContentResponsePtrOutputWithContext(ctx context.Context) AgreementContentResponsePtrOutput {
	return o
}

func (o AgreementContentResponsePtrOutput) Elem() AgreementContentResponseOutput {
	return o.ApplyT(func(v *AgreementContentResponse) AgreementContentResponse {
		if v != nil {
			return *v
		}
		var ret AgreementContentResponse
		return ret
	}).(AgreementContentResponseOutput)
}

func (o AgreementContentResponsePtrOutput) AS2() AS2AgreementContentResponsePtrOutput {
	return o.ApplyT(func(v *AgreementContentResponse) *AS2AgreementContentResponse {
		if v == nil {
			return nil
		}
		return v.AS2
	}).(AS2AgreementContentResponsePtrOutput)
}

func (o AgreementContentResponsePtrOutput) Edifact() EdifactAgreementContentResponsePtrOutput {
	return o.ApplyT(func(v *AgreementContentResponse) *EdifactAgreementContentResponse {
		if v == nil {
			return nil
		}
		return v.Edifact
	}).(EdifactAgreementContentResponsePtrOutput)
}

func (o AgreementContentResponsePtrOutput) X12() X12AgreementContentResponsePtrOutput {
	return o.ApplyT(func(v *AgreementContentResponse) *X12AgreementContentResponse {
		if v == nil {
			return nil
		}
		return v.X12
	}).(X12AgreementContentResponsePtrOutput)
}

type B2BPartnerContent struct {
	BusinessIdentities []BusinessIdentity `pulumi:"businessIdentities"`
}





type B2BPartnerContentInput interface {
	pulumi.Input

	ToB2BPartnerContentOutput() B2BPartnerContentOutput
	ToB2BPartnerContentOutputWithContext(context.Context) B2BPartnerContentOutput
}

type B2BPartnerContentArgs struct {
	BusinessIdentities BusinessIdentityArrayInput `pulumi:"businessIdentities"`
}

func (B2BPartnerContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*B2BPartnerContent)(nil)).Elem()
}

func (i B2BPartnerContentArgs) ToB2BPartnerContentOutput() B2BPartnerContentOutput {
	return i.ToB2BPartnerContentOutputWithContext(context.Background())
}

func (i B2BPartnerContentArgs) ToB2BPartnerContentOutputWithContext(ctx context.Context) B2BPartnerContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2BPartnerContentOutput)
}

func (i B2BPartnerContentArgs) ToB2BPartnerContentPtrOutput() B2BPartnerContentPtrOutput {
	return i.ToB2BPartnerContentPtrOutputWithContext(context.Background())
}

func (i B2BPartnerContentArgs) ToB2BPartnerContentPtrOutputWithContext(ctx context.Context) B2BPartnerContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2BPartnerContentOutput).ToB2BPartnerContentPtrOutputWithContext(ctx)
}









type B2BPartnerContentPtrInput interface {
	pulumi.Input

	ToB2BPartnerContentPtrOutput() B2BPartnerContentPtrOutput
	ToB2BPartnerContentPtrOutputWithContext(context.Context) B2BPartnerContentPtrOutput
}

type b2bpartnerContentPtrType B2BPartnerContentArgs

func B2BPartnerContentPtr(v *B2BPartnerContentArgs) B2BPartnerContentPtrInput {
	return (*b2bpartnerContentPtrType)(v)
}

func (*b2bpartnerContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**B2BPartnerContent)(nil)).Elem()
}

func (i *b2bpartnerContentPtrType) ToB2BPartnerContentPtrOutput() B2BPartnerContentPtrOutput {
	return i.ToB2BPartnerContentPtrOutputWithContext(context.Background())
}

func (i *b2bpartnerContentPtrType) ToB2BPartnerContentPtrOutputWithContext(ctx context.Context) B2BPartnerContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2BPartnerContentPtrOutput)
}

type B2BPartnerContentOutput struct{ *pulumi.OutputState }

func (B2BPartnerContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*B2BPartnerContent)(nil)).Elem()
}

func (o B2BPartnerContentOutput) ToB2BPartnerContentOutput() B2BPartnerContentOutput {
	return o
}

func (o B2BPartnerContentOutput) ToB2BPartnerContentOutputWithContext(ctx context.Context) B2BPartnerContentOutput {
	return o
}

func (o B2BPartnerContentOutput) ToB2BPartnerContentPtrOutput() B2BPartnerContentPtrOutput {
	return o.ToB2BPartnerContentPtrOutputWithContext(context.Background())
}

func (o B2BPartnerContentOutput) ToB2BPartnerContentPtrOutputWithContext(ctx context.Context) B2BPartnerContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v B2BPartnerContent) *B2BPartnerContent {
		return &v
	}).(B2BPartnerContentPtrOutput)
}

func (o B2BPartnerContentOutput) BusinessIdentities() BusinessIdentityArrayOutput {
	return o.ApplyT(func(v B2BPartnerContent) []BusinessIdentity { return v.BusinessIdentities }).(BusinessIdentityArrayOutput)
}

type B2BPartnerContentPtrOutput struct{ *pulumi.OutputState }

func (B2BPartnerContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2BPartnerContent)(nil)).Elem()
}

func (o B2BPartnerContentPtrOutput) ToB2BPartnerContentPtrOutput() B2BPartnerContentPtrOutput {
	return o
}

func (o B2BPartnerContentPtrOutput) ToB2BPartnerContentPtrOutputWithContext(ctx context.Context) B2BPartnerContentPtrOutput {
	return o
}

func (o B2BPartnerContentPtrOutput) Elem() B2BPartnerContentOutput {
	return o.ApplyT(func(v *B2BPartnerContent) B2BPartnerContent {
		if v != nil {
			return *v
		}
		var ret B2BPartnerContent
		return ret
	}).(B2BPartnerContentOutput)
}

func (o B2BPartnerContentPtrOutput) BusinessIdentities() BusinessIdentityArrayOutput {
	return o.ApplyT(func(v *B2BPartnerContent) []BusinessIdentity {
		if v == nil {
			return nil
		}
		return v.BusinessIdentities
	}).(BusinessIdentityArrayOutput)
}

type B2BPartnerContentResponse struct {
	BusinessIdentities []BusinessIdentityResponse `pulumi:"businessIdentities"`
}

type B2BPartnerContentResponseOutput struct{ *pulumi.OutputState }

func (B2BPartnerContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*B2BPartnerContentResponse)(nil)).Elem()
}

func (o B2BPartnerContentResponseOutput) ToB2BPartnerContentResponseOutput() B2BPartnerContentResponseOutput {
	return o
}

func (o B2BPartnerContentResponseOutput) ToB2BPartnerContentResponseOutputWithContext(ctx context.Context) B2BPartnerContentResponseOutput {
	return o
}

func (o B2BPartnerContentResponseOutput) BusinessIdentities() BusinessIdentityResponseArrayOutput {
	return o.ApplyT(func(v B2BPartnerContentResponse) []BusinessIdentityResponse { return v.BusinessIdentities }).(BusinessIdentityResponseArrayOutput)
}

type B2BPartnerContentResponsePtrOutput struct{ *pulumi.OutputState }

func (B2BPartnerContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2BPartnerContentResponse)(nil)).Elem()
}

func (o B2BPartnerContentResponsePtrOutput) ToB2BPartnerContentResponsePtrOutput() B2BPartnerContentResponsePtrOutput {
	return o
}

func (o B2BPartnerContentResponsePtrOutput) ToB2BPartnerContentResponsePtrOutputWithContext(ctx context.Context) B2BPartnerContentResponsePtrOutput {
	return o
}

func (o B2BPartnerContentResponsePtrOutput) Elem() B2BPartnerContentResponseOutput {
	return o.ApplyT(func(v *B2BPartnerContentResponse) B2BPartnerContentResponse {
		if v != nil {
			return *v
		}
		var ret B2BPartnerContentResponse
		return ret
	}).(B2BPartnerContentResponseOutput)
}

func (o B2BPartnerContentResponsePtrOutput) BusinessIdentities() BusinessIdentityResponseArrayOutput {
	return o.ApplyT(func(v *B2BPartnerContentResponse) []BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return v.BusinessIdentities
	}).(BusinessIdentityResponseArrayOutput)
}

type BusinessIdentity struct {
	Qualifier *string `pulumi:"qualifier"`
	Value     *string `pulumi:"value"`
}





type BusinessIdentityInput interface {
	pulumi.Input

	ToBusinessIdentityOutput() BusinessIdentityOutput
	ToBusinessIdentityOutputWithContext(context.Context) BusinessIdentityOutput
}

type BusinessIdentityArgs struct {
	Qualifier pulumi.StringPtrInput `pulumi:"qualifier"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (BusinessIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BusinessIdentity)(nil)).Elem()
}

func (i BusinessIdentityArgs) ToBusinessIdentityOutput() BusinessIdentityOutput {
	return i.ToBusinessIdentityOutputWithContext(context.Background())
}

func (i BusinessIdentityArgs) ToBusinessIdentityOutputWithContext(ctx context.Context) BusinessIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIdentityOutput)
}

func (i BusinessIdentityArgs) ToBusinessIdentityPtrOutput() BusinessIdentityPtrOutput {
	return i.ToBusinessIdentityPtrOutputWithContext(context.Background())
}

func (i BusinessIdentityArgs) ToBusinessIdentityPtrOutputWithContext(ctx context.Context) BusinessIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIdentityOutput).ToBusinessIdentityPtrOutputWithContext(ctx)
}









type BusinessIdentityPtrInput interface {
	pulumi.Input

	ToBusinessIdentityPtrOutput() BusinessIdentityPtrOutput
	ToBusinessIdentityPtrOutputWithContext(context.Context) BusinessIdentityPtrOutput
}

type businessIdentityPtrType BusinessIdentityArgs

func BusinessIdentityPtr(v *BusinessIdentityArgs) BusinessIdentityPtrInput {
	return (*businessIdentityPtrType)(v)
}

func (*businessIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessIdentity)(nil)).Elem()
}

func (i *businessIdentityPtrType) ToBusinessIdentityPtrOutput() BusinessIdentityPtrOutput {
	return i.ToBusinessIdentityPtrOutputWithContext(context.Background())
}

func (i *businessIdentityPtrType) ToBusinessIdentityPtrOutputWithContext(ctx context.Context) BusinessIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIdentityPtrOutput)
}





type BusinessIdentityArrayInput interface {
	pulumi.Input

	ToBusinessIdentityArrayOutput() BusinessIdentityArrayOutput
	ToBusinessIdentityArrayOutputWithContext(context.Context) BusinessIdentityArrayOutput
}

type BusinessIdentityArray []BusinessIdentityInput

func (BusinessIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BusinessIdentity)(nil)).Elem()
}

func (i BusinessIdentityArray) ToBusinessIdentityArrayOutput() BusinessIdentityArrayOutput {
	return i.ToBusinessIdentityArrayOutputWithContext(context.Background())
}

func (i BusinessIdentityArray) ToBusinessIdentityArrayOutputWithContext(ctx context.Context) BusinessIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIdentityArrayOutput)
}

type BusinessIdentityOutput struct{ *pulumi.OutputState }

func (BusinessIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BusinessIdentity)(nil)).Elem()
}

func (o BusinessIdentityOutput) ToBusinessIdentityOutput() BusinessIdentityOutput {
	return o
}

func (o BusinessIdentityOutput) ToBusinessIdentityOutputWithContext(ctx context.Context) BusinessIdentityOutput {
	return o
}

func (o BusinessIdentityOutput) ToBusinessIdentityPtrOutput() BusinessIdentityPtrOutput {
	return o.ToBusinessIdentityPtrOutputWithContext(context.Background())
}

func (o BusinessIdentityOutput) ToBusinessIdentityPtrOutputWithContext(ctx context.Context) BusinessIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BusinessIdentity) *BusinessIdentity {
		return &v
	}).(BusinessIdentityPtrOutput)
}

func (o BusinessIdentityOutput) Qualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BusinessIdentity) *string { return v.Qualifier }).(pulumi.StringPtrOutput)
}

func (o BusinessIdentityOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BusinessIdentity) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type BusinessIdentityPtrOutput struct{ *pulumi.OutputState }

func (BusinessIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessIdentity)(nil)).Elem()
}

func (o BusinessIdentityPtrOutput) ToBusinessIdentityPtrOutput() BusinessIdentityPtrOutput {
	return o
}

func (o BusinessIdentityPtrOutput) ToBusinessIdentityPtrOutputWithContext(ctx context.Context) BusinessIdentityPtrOutput {
	return o
}

func (o BusinessIdentityPtrOutput) Elem() BusinessIdentityOutput {
	return o.ApplyT(func(v *BusinessIdentity) BusinessIdentity {
		if v != nil {
			return *v
		}
		var ret BusinessIdentity
		return ret
	}).(BusinessIdentityOutput)
}

func (o BusinessIdentityPtrOutput) Qualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BusinessIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Qualifier
	}).(pulumi.StringPtrOutput)
}

func (o BusinessIdentityPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BusinessIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type BusinessIdentityArrayOutput struct{ *pulumi.OutputState }

func (BusinessIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BusinessIdentity)(nil)).Elem()
}

func (o BusinessIdentityArrayOutput) ToBusinessIdentityArrayOutput() BusinessIdentityArrayOutput {
	return o
}

func (o BusinessIdentityArrayOutput) ToBusinessIdentityArrayOutputWithContext(ctx context.Context) BusinessIdentityArrayOutput {
	return o
}

func (o BusinessIdentityArrayOutput) Index(i pulumi.IntInput) BusinessIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BusinessIdentity {
		return vs[0].([]BusinessIdentity)[vs[1].(int)]
	}).(BusinessIdentityOutput)
}

type BusinessIdentityResponse struct {
	Qualifier *string `pulumi:"qualifier"`
	Value     *string `pulumi:"value"`
}

type BusinessIdentityResponseOutput struct{ *pulumi.OutputState }

func (BusinessIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BusinessIdentityResponse)(nil)).Elem()
}

func (o BusinessIdentityResponseOutput) ToBusinessIdentityResponseOutput() BusinessIdentityResponseOutput {
	return o
}

func (o BusinessIdentityResponseOutput) ToBusinessIdentityResponseOutputWithContext(ctx context.Context) BusinessIdentityResponseOutput {
	return o
}

func (o BusinessIdentityResponseOutput) Qualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BusinessIdentityResponse) *string { return v.Qualifier }).(pulumi.StringPtrOutput)
}

func (o BusinessIdentityResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BusinessIdentityResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type BusinessIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (BusinessIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessIdentityResponse)(nil)).Elem()
}

func (o BusinessIdentityResponsePtrOutput) ToBusinessIdentityResponsePtrOutput() BusinessIdentityResponsePtrOutput {
	return o
}

func (o BusinessIdentityResponsePtrOutput) ToBusinessIdentityResponsePtrOutputWithContext(ctx context.Context) BusinessIdentityResponsePtrOutput {
	return o
}

func (o BusinessIdentityResponsePtrOutput) Elem() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v *BusinessIdentityResponse) BusinessIdentityResponse {
		if v != nil {
			return *v
		}
		var ret BusinessIdentityResponse
		return ret
	}).(BusinessIdentityResponseOutput)
}

func (o BusinessIdentityResponsePtrOutput) Qualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BusinessIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Qualifier
	}).(pulumi.StringPtrOutput)
}

func (o BusinessIdentityResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BusinessIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type BusinessIdentityResponseArrayOutput struct{ *pulumi.OutputState }

func (BusinessIdentityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BusinessIdentityResponse)(nil)).Elem()
}

func (o BusinessIdentityResponseArrayOutput) ToBusinessIdentityResponseArrayOutput() BusinessIdentityResponseArrayOutput {
	return o
}

func (o BusinessIdentityResponseArrayOutput) ToBusinessIdentityResponseArrayOutputWithContext(ctx context.Context) BusinessIdentityResponseArrayOutput {
	return o
}

func (o BusinessIdentityResponseArrayOutput) Index(i pulumi.IntInput) BusinessIdentityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BusinessIdentityResponse {
		return vs[0].([]BusinessIdentityResponse)[vs[1].(int)]
	}).(BusinessIdentityResponseOutput)
}

type EdifactAcknowledgementSettings struct {
	AcknowledgementControlNumberLowerBound *int    `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     *string `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     *string `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound *int    `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        *bool   `pulumi:"batchFunctionalAcknowledgements"`
	BatchTechnicalAcknowledgements         *bool   `pulumi:"batchTechnicalAcknowledgements"`
	NeedFunctionalAcknowledgement          *bool   `pulumi:"needFunctionalAcknowledgement"`
	NeedLoopForValidMessages               *bool   `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           *bool   `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   *bool   `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         *bool   `pulumi:"sendSynchronousAcknowledgement"`
}





type EdifactAcknowledgementSettingsInput interface {
	pulumi.Input

	ToEdifactAcknowledgementSettingsOutput() EdifactAcknowledgementSettingsOutput
	ToEdifactAcknowledgementSettingsOutputWithContext(context.Context) EdifactAcknowledgementSettingsOutput
}

type EdifactAcknowledgementSettingsArgs struct {
	AcknowledgementControlNumberLowerBound pulumi.IntPtrInput    `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     pulumi.StringPtrInput `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     pulumi.StringPtrInput `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound pulumi.IntPtrInput    `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        pulumi.BoolPtrInput   `pulumi:"batchFunctionalAcknowledgements"`
	BatchTechnicalAcknowledgements         pulumi.BoolPtrInput   `pulumi:"batchTechnicalAcknowledgements"`
	NeedFunctionalAcknowledgement          pulumi.BoolPtrInput   `pulumi:"needFunctionalAcknowledgement"`
	NeedLoopForValidMessages               pulumi.BoolPtrInput   `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           pulumi.BoolPtrInput   `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   pulumi.BoolPtrInput   `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         pulumi.BoolPtrInput   `pulumi:"sendSynchronousAcknowledgement"`
}

func (EdifactAcknowledgementSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAcknowledgementSettings)(nil)).Elem()
}

func (i EdifactAcknowledgementSettingsArgs) ToEdifactAcknowledgementSettingsOutput() EdifactAcknowledgementSettingsOutput {
	return i.ToEdifactAcknowledgementSettingsOutputWithContext(context.Background())
}

func (i EdifactAcknowledgementSettingsArgs) ToEdifactAcknowledgementSettingsOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAcknowledgementSettingsOutput)
}

func (i EdifactAcknowledgementSettingsArgs) ToEdifactAcknowledgementSettingsPtrOutput() EdifactAcknowledgementSettingsPtrOutput {
	return i.ToEdifactAcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactAcknowledgementSettingsArgs) ToEdifactAcknowledgementSettingsPtrOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAcknowledgementSettingsOutput).ToEdifactAcknowledgementSettingsPtrOutputWithContext(ctx)
}









type EdifactAcknowledgementSettingsPtrInput interface {
	pulumi.Input

	ToEdifactAcknowledgementSettingsPtrOutput() EdifactAcknowledgementSettingsPtrOutput
	ToEdifactAcknowledgementSettingsPtrOutputWithContext(context.Context) EdifactAcknowledgementSettingsPtrOutput
}

type edifactAcknowledgementSettingsPtrType EdifactAcknowledgementSettingsArgs

func EdifactAcknowledgementSettingsPtr(v *EdifactAcknowledgementSettingsArgs) EdifactAcknowledgementSettingsPtrInput {
	return (*edifactAcknowledgementSettingsPtrType)(v)
}

func (*edifactAcknowledgementSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAcknowledgementSettings)(nil)).Elem()
}

func (i *edifactAcknowledgementSettingsPtrType) ToEdifactAcknowledgementSettingsPtrOutput() EdifactAcknowledgementSettingsPtrOutput {
	return i.ToEdifactAcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactAcknowledgementSettingsPtrType) ToEdifactAcknowledgementSettingsPtrOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAcknowledgementSettingsPtrOutput)
}

type EdifactAcknowledgementSettingsOutput struct{ *pulumi.OutputState }

func (EdifactAcknowledgementSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAcknowledgementSettings)(nil)).Elem()
}

func (o EdifactAcknowledgementSettingsOutput) ToEdifactAcknowledgementSettingsOutput() EdifactAcknowledgementSettingsOutput {
	return o
}

func (o EdifactAcknowledgementSettingsOutput) ToEdifactAcknowledgementSettingsOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsOutput {
	return o
}

func (o EdifactAcknowledgementSettingsOutput) ToEdifactAcknowledgementSettingsPtrOutput() EdifactAcknowledgementSettingsPtrOutput {
	return o.ToEdifactAcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactAcknowledgementSettingsOutput) ToEdifactAcknowledgementSettingsPtrOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactAcknowledgementSettings) *EdifactAcknowledgementSettings {
		return &v
	}).(EdifactAcknowledgementSettingsPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *int { return v.AcknowledgementControlNumberLowerBound }).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *string { return v.AcknowledgementControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *string { return v.AcknowledgementControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *int { return v.AcknowledgementControlNumberUpperBound }).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *bool { return v.BatchFunctionalAcknowledgements }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *bool { return v.BatchTechnicalAcknowledgements }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *bool { return v.NeedFunctionalAcknowledgement }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *bool { return v.NeedLoopForValidMessages }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *bool { return v.NeedTechnicalAcknowledgement }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *bool { return v.RolloverAcknowledgementControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *bool { return v.SendSynchronousAcknowledgement }).(pulumi.BoolPtrOutput)
}

type EdifactAcknowledgementSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactAcknowledgementSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAcknowledgementSettings)(nil)).Elem()
}

func (o EdifactAcknowledgementSettingsPtrOutput) ToEdifactAcknowledgementSettingsPtrOutput() EdifactAcknowledgementSettingsPtrOutput {
	return o
}

func (o EdifactAcknowledgementSettingsPtrOutput) ToEdifactAcknowledgementSettingsPtrOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsPtrOutput {
	return o
}

func (o EdifactAcknowledgementSettingsPtrOutput) Elem() EdifactAcknowledgementSettingsOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) EdifactAcknowledgementSettings {
		if v != nil {
			return *v
		}
		var ret EdifactAcknowledgementSettings
		return ret
	}).(EdifactAcknowledgementSettingsOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *int {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *int {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BatchFunctionalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BatchTechnicalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.NeedFunctionalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.NeedLoopForValidMessages
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.NeedTechnicalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverAcknowledgementControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SendSynchronousAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

type EdifactAcknowledgementSettingsResponse struct {
	AcknowledgementControlNumberLowerBound *int    `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     *string `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     *string `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound *int    `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        *bool   `pulumi:"batchFunctionalAcknowledgements"`
	BatchTechnicalAcknowledgements         *bool   `pulumi:"batchTechnicalAcknowledgements"`
	NeedFunctionalAcknowledgement          *bool   `pulumi:"needFunctionalAcknowledgement"`
	NeedLoopForValidMessages               *bool   `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           *bool   `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   *bool   `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         *bool   `pulumi:"sendSynchronousAcknowledgement"`
}

type EdifactAcknowledgementSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactAcknowledgementSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAcknowledgementSettingsResponse)(nil)).Elem()
}

func (o EdifactAcknowledgementSettingsResponseOutput) ToEdifactAcknowledgementSettingsResponseOutput() EdifactAcknowledgementSettingsResponseOutput {
	return o
}

func (o EdifactAcknowledgementSettingsResponseOutput) ToEdifactAcknowledgementSettingsResponseOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsResponseOutput {
	return o
}

func (o EdifactAcknowledgementSettingsResponseOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *int { return v.AcknowledgementControlNumberLowerBound }).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *string { return v.AcknowledgementControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *string { return v.AcknowledgementControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *int { return v.AcknowledgementControlNumberUpperBound }).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *bool { return v.BatchFunctionalAcknowledgements }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *bool { return v.BatchTechnicalAcknowledgements }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *bool { return v.NeedFunctionalAcknowledgement }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *bool { return v.NeedLoopForValidMessages }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *bool { return v.NeedTechnicalAcknowledgement }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *bool { return v.RolloverAcknowledgementControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *bool { return v.SendSynchronousAcknowledgement }).(pulumi.BoolPtrOutput)
}

type EdifactAcknowledgementSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactAcknowledgementSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAcknowledgementSettingsResponse)(nil)).Elem()
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) ToEdifactAcknowledgementSettingsResponsePtrOutput() EdifactAcknowledgementSettingsResponsePtrOutput {
	return o
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) ToEdifactAcknowledgementSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsResponsePtrOutput {
	return o
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) Elem() EdifactAcknowledgementSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) EdifactAcknowledgementSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactAcknowledgementSettingsResponse
		return ret
	}).(EdifactAcknowledgementSettingsResponseOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.BatchFunctionalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.BatchTechnicalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.NeedFunctionalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.NeedLoopForValidMessages
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.NeedTechnicalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverAcknowledgementControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendSynchronousAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

type EdifactAgreementContent struct {
	ReceiveAgreement *EdifactOneWayAgreement `pulumi:"receiveAgreement"`
	SendAgreement    *EdifactOneWayAgreement `pulumi:"sendAgreement"`
}





type EdifactAgreementContentInput interface {
	pulumi.Input

	ToEdifactAgreementContentOutput() EdifactAgreementContentOutput
	ToEdifactAgreementContentOutputWithContext(context.Context) EdifactAgreementContentOutput
}

type EdifactAgreementContentArgs struct {
	ReceiveAgreement EdifactOneWayAgreementPtrInput `pulumi:"receiveAgreement"`
	SendAgreement    EdifactOneWayAgreementPtrInput `pulumi:"sendAgreement"`
}

func (EdifactAgreementContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAgreementContent)(nil)).Elem()
}

func (i EdifactAgreementContentArgs) ToEdifactAgreementContentOutput() EdifactAgreementContentOutput {
	return i.ToEdifactAgreementContentOutputWithContext(context.Background())
}

func (i EdifactAgreementContentArgs) ToEdifactAgreementContentOutputWithContext(ctx context.Context) EdifactAgreementContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAgreementContentOutput)
}

func (i EdifactAgreementContentArgs) ToEdifactAgreementContentPtrOutput() EdifactAgreementContentPtrOutput {
	return i.ToEdifactAgreementContentPtrOutputWithContext(context.Background())
}

func (i EdifactAgreementContentArgs) ToEdifactAgreementContentPtrOutputWithContext(ctx context.Context) EdifactAgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAgreementContentOutput).ToEdifactAgreementContentPtrOutputWithContext(ctx)
}









type EdifactAgreementContentPtrInput interface {
	pulumi.Input

	ToEdifactAgreementContentPtrOutput() EdifactAgreementContentPtrOutput
	ToEdifactAgreementContentPtrOutputWithContext(context.Context) EdifactAgreementContentPtrOutput
}

type edifactAgreementContentPtrType EdifactAgreementContentArgs

func EdifactAgreementContentPtr(v *EdifactAgreementContentArgs) EdifactAgreementContentPtrInput {
	return (*edifactAgreementContentPtrType)(v)
}

func (*edifactAgreementContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAgreementContent)(nil)).Elem()
}

func (i *edifactAgreementContentPtrType) ToEdifactAgreementContentPtrOutput() EdifactAgreementContentPtrOutput {
	return i.ToEdifactAgreementContentPtrOutputWithContext(context.Background())
}

func (i *edifactAgreementContentPtrType) ToEdifactAgreementContentPtrOutputWithContext(ctx context.Context) EdifactAgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAgreementContentPtrOutput)
}

type EdifactAgreementContentOutput struct{ *pulumi.OutputState }

func (EdifactAgreementContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAgreementContent)(nil)).Elem()
}

func (o EdifactAgreementContentOutput) ToEdifactAgreementContentOutput() EdifactAgreementContentOutput {
	return o
}

func (o EdifactAgreementContentOutput) ToEdifactAgreementContentOutputWithContext(ctx context.Context) EdifactAgreementContentOutput {
	return o
}

func (o EdifactAgreementContentOutput) ToEdifactAgreementContentPtrOutput() EdifactAgreementContentPtrOutput {
	return o.ToEdifactAgreementContentPtrOutputWithContext(context.Background())
}

func (o EdifactAgreementContentOutput) ToEdifactAgreementContentPtrOutputWithContext(ctx context.Context) EdifactAgreementContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactAgreementContent) *EdifactAgreementContent {
		return &v
	}).(EdifactAgreementContentPtrOutput)
}

func (o EdifactAgreementContentOutput) ReceiveAgreement() EdifactOneWayAgreementPtrOutput {
	return o.ApplyT(func(v EdifactAgreementContent) *EdifactOneWayAgreement { return v.ReceiveAgreement }).(EdifactOneWayAgreementPtrOutput)
}

func (o EdifactAgreementContentOutput) SendAgreement() EdifactOneWayAgreementPtrOutput {
	return o.ApplyT(func(v EdifactAgreementContent) *EdifactOneWayAgreement { return v.SendAgreement }).(EdifactOneWayAgreementPtrOutput)
}

type EdifactAgreementContentPtrOutput struct{ *pulumi.OutputState }

func (EdifactAgreementContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAgreementContent)(nil)).Elem()
}

func (o EdifactAgreementContentPtrOutput) ToEdifactAgreementContentPtrOutput() EdifactAgreementContentPtrOutput {
	return o
}

func (o EdifactAgreementContentPtrOutput) ToEdifactAgreementContentPtrOutputWithContext(ctx context.Context) EdifactAgreementContentPtrOutput {
	return o
}

func (o EdifactAgreementContentPtrOutput) Elem() EdifactAgreementContentOutput {
	return o.ApplyT(func(v *EdifactAgreementContent) EdifactAgreementContent {
		if v != nil {
			return *v
		}
		var ret EdifactAgreementContent
		return ret
	}).(EdifactAgreementContentOutput)
}

func (o EdifactAgreementContentPtrOutput) ReceiveAgreement() EdifactOneWayAgreementPtrOutput {
	return o.ApplyT(func(v *EdifactAgreementContent) *EdifactOneWayAgreement {
		if v == nil {
			return nil
		}
		return v.ReceiveAgreement
	}).(EdifactOneWayAgreementPtrOutput)
}

func (o EdifactAgreementContentPtrOutput) SendAgreement() EdifactOneWayAgreementPtrOutput {
	return o.ApplyT(func(v *EdifactAgreementContent) *EdifactOneWayAgreement {
		if v == nil {
			return nil
		}
		return v.SendAgreement
	}).(EdifactOneWayAgreementPtrOutput)
}

type EdifactAgreementContentResponse struct {
	ReceiveAgreement *EdifactOneWayAgreementResponse `pulumi:"receiveAgreement"`
	SendAgreement    *EdifactOneWayAgreementResponse `pulumi:"sendAgreement"`
}

type EdifactAgreementContentResponseOutput struct{ *pulumi.OutputState }

func (EdifactAgreementContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAgreementContentResponse)(nil)).Elem()
}

func (o EdifactAgreementContentResponseOutput) ToEdifactAgreementContentResponseOutput() EdifactAgreementContentResponseOutput {
	return o
}

func (o EdifactAgreementContentResponseOutput) ToEdifactAgreementContentResponseOutputWithContext(ctx context.Context) EdifactAgreementContentResponseOutput {
	return o
}

func (o EdifactAgreementContentResponseOutput) ReceiveAgreement() EdifactOneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v EdifactAgreementContentResponse) *EdifactOneWayAgreementResponse { return v.ReceiveAgreement }).(EdifactOneWayAgreementResponsePtrOutput)
}

func (o EdifactAgreementContentResponseOutput) SendAgreement() EdifactOneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v EdifactAgreementContentResponse) *EdifactOneWayAgreementResponse { return v.SendAgreement }).(EdifactOneWayAgreementResponsePtrOutput)
}

type EdifactAgreementContentResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactAgreementContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAgreementContentResponse)(nil)).Elem()
}

func (o EdifactAgreementContentResponsePtrOutput) ToEdifactAgreementContentResponsePtrOutput() EdifactAgreementContentResponsePtrOutput {
	return o
}

func (o EdifactAgreementContentResponsePtrOutput) ToEdifactAgreementContentResponsePtrOutputWithContext(ctx context.Context) EdifactAgreementContentResponsePtrOutput {
	return o
}

func (o EdifactAgreementContentResponsePtrOutput) Elem() EdifactAgreementContentResponseOutput {
	return o.ApplyT(func(v *EdifactAgreementContentResponse) EdifactAgreementContentResponse {
		if v != nil {
			return *v
		}
		var ret EdifactAgreementContentResponse
		return ret
	}).(EdifactAgreementContentResponseOutput)
}

func (o EdifactAgreementContentResponsePtrOutput) ReceiveAgreement() EdifactOneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *EdifactAgreementContentResponse) *EdifactOneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return v.ReceiveAgreement
	}).(EdifactOneWayAgreementResponsePtrOutput)
}

func (o EdifactAgreementContentResponsePtrOutput) SendAgreement() EdifactOneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *EdifactAgreementContentResponse) *EdifactOneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return v.SendAgreement
	}).(EdifactOneWayAgreementResponsePtrOutput)
}

type EdifactDelimiterOverride struct {
	ComponentSeparator             *int                     `pulumi:"componentSeparator"`
	DataElementSeparator           *int                     `pulumi:"dataElementSeparator"`
	DecimalPointIndicator          *EdifactDecimalIndicator `pulumi:"decimalPointIndicator"`
	MessageAssociationAssignedCode *string                  `pulumi:"messageAssociationAssignedCode"`
	MessageId                      *string                  `pulumi:"messageId"`
	MessageRelease                 *string                  `pulumi:"messageRelease"`
	MessageVersion                 *string                  `pulumi:"messageVersion"`
	ReleaseIndicator               *int                     `pulumi:"releaseIndicator"`
	RepetitionSeparator            *int                     `pulumi:"repetitionSeparator"`
	SegmentTerminator              *int                     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix        *SegmentTerminatorSuffix `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace                *string                  `pulumi:"targetNamespace"`
}





type EdifactDelimiterOverrideInput interface {
	pulumi.Input

	ToEdifactDelimiterOverrideOutput() EdifactDelimiterOverrideOutput
	ToEdifactDelimiterOverrideOutputWithContext(context.Context) EdifactDelimiterOverrideOutput
}

type EdifactDelimiterOverrideArgs struct {
	ComponentSeparator             pulumi.IntPtrInput              `pulumi:"componentSeparator"`
	DataElementSeparator           pulumi.IntPtrInput              `pulumi:"dataElementSeparator"`
	DecimalPointIndicator          EdifactDecimalIndicatorPtrInput `pulumi:"decimalPointIndicator"`
	MessageAssociationAssignedCode pulumi.StringPtrInput           `pulumi:"messageAssociationAssignedCode"`
	MessageId                      pulumi.StringPtrInput           `pulumi:"messageId"`
	MessageRelease                 pulumi.StringPtrInput           `pulumi:"messageRelease"`
	MessageVersion                 pulumi.StringPtrInput           `pulumi:"messageVersion"`
	ReleaseIndicator               pulumi.IntPtrInput              `pulumi:"releaseIndicator"`
	RepetitionSeparator            pulumi.IntPtrInput              `pulumi:"repetitionSeparator"`
	SegmentTerminator              pulumi.IntPtrInput              `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix        SegmentTerminatorSuffixPtrInput `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace                pulumi.StringPtrInput           `pulumi:"targetNamespace"`
}

func (EdifactDelimiterOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactDelimiterOverride)(nil)).Elem()
}

func (i EdifactDelimiterOverrideArgs) ToEdifactDelimiterOverrideOutput() EdifactDelimiterOverrideOutput {
	return i.ToEdifactDelimiterOverrideOutputWithContext(context.Background())
}

func (i EdifactDelimiterOverrideArgs) ToEdifactDelimiterOverrideOutputWithContext(ctx context.Context) EdifactDelimiterOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactDelimiterOverrideOutput)
}





type EdifactDelimiterOverrideArrayInput interface {
	pulumi.Input

	ToEdifactDelimiterOverrideArrayOutput() EdifactDelimiterOverrideArrayOutput
	ToEdifactDelimiterOverrideArrayOutputWithContext(context.Context) EdifactDelimiterOverrideArrayOutput
}

type EdifactDelimiterOverrideArray []EdifactDelimiterOverrideInput

func (EdifactDelimiterOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactDelimiterOverride)(nil)).Elem()
}

func (i EdifactDelimiterOverrideArray) ToEdifactDelimiterOverrideArrayOutput() EdifactDelimiterOverrideArrayOutput {
	return i.ToEdifactDelimiterOverrideArrayOutputWithContext(context.Background())
}

func (i EdifactDelimiterOverrideArray) ToEdifactDelimiterOverrideArrayOutputWithContext(ctx context.Context) EdifactDelimiterOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactDelimiterOverrideArrayOutput)
}

type EdifactDelimiterOverrideOutput struct{ *pulumi.OutputState }

func (EdifactDelimiterOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactDelimiterOverride)(nil)).Elem()
}

func (o EdifactDelimiterOverrideOutput) ToEdifactDelimiterOverrideOutput() EdifactDelimiterOverrideOutput {
	return o
}

func (o EdifactDelimiterOverrideOutput) ToEdifactDelimiterOverrideOutputWithContext(ctx context.Context) EdifactDelimiterOverrideOutput {
	return o
}

func (o EdifactDelimiterOverrideOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *int { return v.ComponentSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *int { return v.DataElementSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) DecimalPointIndicator() EdifactDecimalIndicatorPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *EdifactDecimalIndicator { return v.DecimalPointIndicator }).(EdifactDecimalIndicatorPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) MessageAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *string { return v.MessageAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) MessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *string { return v.MessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) MessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *string { return v.MessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) ReleaseIndicator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *int { return v.ReleaseIndicator }).(pulumi.IntPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) RepetitionSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *int { return v.RepetitionSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *int { return v.SegmentTerminator }).(pulumi.IntPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *SegmentTerminatorSuffix { return v.SegmentTerminatorSuffix }).(SegmentTerminatorSuffixPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type EdifactDelimiterOverrideArrayOutput struct{ *pulumi.OutputState }

func (EdifactDelimiterOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactDelimiterOverride)(nil)).Elem()
}

func (o EdifactDelimiterOverrideArrayOutput) ToEdifactDelimiterOverrideArrayOutput() EdifactDelimiterOverrideArrayOutput {
	return o
}

func (o EdifactDelimiterOverrideArrayOutput) ToEdifactDelimiterOverrideArrayOutputWithContext(ctx context.Context) EdifactDelimiterOverrideArrayOutput {
	return o
}

func (o EdifactDelimiterOverrideArrayOutput) Index(i pulumi.IntInput) EdifactDelimiterOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactDelimiterOverride {
		return vs[0].([]EdifactDelimiterOverride)[vs[1].(int)]
	}).(EdifactDelimiterOverrideOutput)
}

type EdifactDelimiterOverrideResponse struct {
	ComponentSeparator             *int    `pulumi:"componentSeparator"`
	DataElementSeparator           *int    `pulumi:"dataElementSeparator"`
	DecimalPointIndicator          *string `pulumi:"decimalPointIndicator"`
	MessageAssociationAssignedCode *string `pulumi:"messageAssociationAssignedCode"`
	MessageId                      *string `pulumi:"messageId"`
	MessageRelease                 *string `pulumi:"messageRelease"`
	MessageVersion                 *string `pulumi:"messageVersion"`
	ReleaseIndicator               *int    `pulumi:"releaseIndicator"`
	RepetitionSeparator            *int    `pulumi:"repetitionSeparator"`
	SegmentTerminator              *int    `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix        *string `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace                *string `pulumi:"targetNamespace"`
}

type EdifactDelimiterOverrideResponseOutput struct{ *pulumi.OutputState }

func (EdifactDelimiterOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactDelimiterOverrideResponse)(nil)).Elem()
}

func (o EdifactDelimiterOverrideResponseOutput) ToEdifactDelimiterOverrideResponseOutput() EdifactDelimiterOverrideResponseOutput {
	return o
}

func (o EdifactDelimiterOverrideResponseOutput) ToEdifactDelimiterOverrideResponseOutputWithContext(ctx context.Context) EdifactDelimiterOverrideResponseOutput {
	return o
}

func (o EdifactDelimiterOverrideResponseOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *int { return v.ComponentSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *int { return v.DataElementSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) DecimalPointIndicator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.DecimalPointIndicator }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) MessageAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.MessageAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) MessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.MessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) MessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.MessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) ReleaseIndicator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *int { return v.ReleaseIndicator }).(pulumi.IntPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) RepetitionSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *int { return v.RepetitionSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *int { return v.SegmentTerminator }).(pulumi.IntPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) SegmentTerminatorSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.SegmentTerminatorSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type EdifactDelimiterOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (EdifactDelimiterOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactDelimiterOverrideResponse)(nil)).Elem()
}

func (o EdifactDelimiterOverrideResponseArrayOutput) ToEdifactDelimiterOverrideResponseArrayOutput() EdifactDelimiterOverrideResponseArrayOutput {
	return o
}

func (o EdifactDelimiterOverrideResponseArrayOutput) ToEdifactDelimiterOverrideResponseArrayOutputWithContext(ctx context.Context) EdifactDelimiterOverrideResponseArrayOutput {
	return o
}

func (o EdifactDelimiterOverrideResponseArrayOutput) Index(i pulumi.IntInput) EdifactDelimiterOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactDelimiterOverrideResponse {
		return vs[0].([]EdifactDelimiterOverrideResponse)[vs[1].(int)]
	}).(EdifactDelimiterOverrideResponseOutput)
}

type EdifactEnvelopeOverride struct {
	ApplicationPassword            *string `pulumi:"applicationPassword"`
	AssociationAssignedCode        *string `pulumi:"associationAssignedCode"`
	ControllingAgencyCode          *string `pulumi:"controllingAgencyCode"`
	FunctionalGroupId              *string `pulumi:"functionalGroupId"`
	GroupHeaderMessageRelease      *string `pulumi:"groupHeaderMessageRelease"`
	GroupHeaderMessageVersion      *string `pulumi:"groupHeaderMessageVersion"`
	MessageAssociationAssignedCode *string `pulumi:"messageAssociationAssignedCode"`
	MessageId                      *string `pulumi:"messageId"`
	MessageRelease                 *string `pulumi:"messageRelease"`
	MessageVersion                 *string `pulumi:"messageVersion"`
	ReceiverApplicationId          *string `pulumi:"receiverApplicationId"`
	ReceiverApplicationQualifier   *string `pulumi:"receiverApplicationQualifier"`
	SenderApplicationId            *string `pulumi:"senderApplicationId"`
	SenderApplicationQualifier     *string `pulumi:"senderApplicationQualifier"`
	TargetNamespace                *string `pulumi:"targetNamespace"`
}





type EdifactEnvelopeOverrideInput interface {
	pulumi.Input

	ToEdifactEnvelopeOverrideOutput() EdifactEnvelopeOverrideOutput
	ToEdifactEnvelopeOverrideOutputWithContext(context.Context) EdifactEnvelopeOverrideOutput
}

type EdifactEnvelopeOverrideArgs struct {
	ApplicationPassword            pulumi.StringPtrInput `pulumi:"applicationPassword"`
	AssociationAssignedCode        pulumi.StringPtrInput `pulumi:"associationAssignedCode"`
	ControllingAgencyCode          pulumi.StringPtrInput `pulumi:"controllingAgencyCode"`
	FunctionalGroupId              pulumi.StringPtrInput `pulumi:"functionalGroupId"`
	GroupHeaderMessageRelease      pulumi.StringPtrInput `pulumi:"groupHeaderMessageRelease"`
	GroupHeaderMessageVersion      pulumi.StringPtrInput `pulumi:"groupHeaderMessageVersion"`
	MessageAssociationAssignedCode pulumi.StringPtrInput `pulumi:"messageAssociationAssignedCode"`
	MessageId                      pulumi.StringPtrInput `pulumi:"messageId"`
	MessageRelease                 pulumi.StringPtrInput `pulumi:"messageRelease"`
	MessageVersion                 pulumi.StringPtrInput `pulumi:"messageVersion"`
	ReceiverApplicationId          pulumi.StringPtrInput `pulumi:"receiverApplicationId"`
	ReceiverApplicationQualifier   pulumi.StringPtrInput `pulumi:"receiverApplicationQualifier"`
	SenderApplicationId            pulumi.StringPtrInput `pulumi:"senderApplicationId"`
	SenderApplicationQualifier     pulumi.StringPtrInput `pulumi:"senderApplicationQualifier"`
	TargetNamespace                pulumi.StringPtrInput `pulumi:"targetNamespace"`
}

func (EdifactEnvelopeOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeOverride)(nil)).Elem()
}

func (i EdifactEnvelopeOverrideArgs) ToEdifactEnvelopeOverrideOutput() EdifactEnvelopeOverrideOutput {
	return i.ToEdifactEnvelopeOverrideOutputWithContext(context.Background())
}

func (i EdifactEnvelopeOverrideArgs) ToEdifactEnvelopeOverrideOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactEnvelopeOverrideOutput)
}





type EdifactEnvelopeOverrideArrayInput interface {
	pulumi.Input

	ToEdifactEnvelopeOverrideArrayOutput() EdifactEnvelopeOverrideArrayOutput
	ToEdifactEnvelopeOverrideArrayOutputWithContext(context.Context) EdifactEnvelopeOverrideArrayOutput
}

type EdifactEnvelopeOverrideArray []EdifactEnvelopeOverrideInput

func (EdifactEnvelopeOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactEnvelopeOverride)(nil)).Elem()
}

func (i EdifactEnvelopeOverrideArray) ToEdifactEnvelopeOverrideArrayOutput() EdifactEnvelopeOverrideArrayOutput {
	return i.ToEdifactEnvelopeOverrideArrayOutputWithContext(context.Background())
}

func (i EdifactEnvelopeOverrideArray) ToEdifactEnvelopeOverrideArrayOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactEnvelopeOverrideArrayOutput)
}

type EdifactEnvelopeOverrideOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeOverride)(nil)).Elem()
}

func (o EdifactEnvelopeOverrideOutput) ToEdifactEnvelopeOverrideOutput() EdifactEnvelopeOverrideOutput {
	return o
}

func (o EdifactEnvelopeOverrideOutput) ToEdifactEnvelopeOverrideOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideOutput {
	return o
}

func (o EdifactEnvelopeOverrideOutput) ApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.ApplicationPassword }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) AssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.AssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) ControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.ControllingAgencyCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) GroupHeaderMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.GroupHeaderMessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) GroupHeaderMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.GroupHeaderMessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) MessageAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.MessageAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) MessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.MessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) MessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.MessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.ReceiverApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) ReceiverApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.ReceiverApplicationQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) SenderApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.SenderApplicationQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type EdifactEnvelopeOverrideArrayOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactEnvelopeOverride)(nil)).Elem()
}

func (o EdifactEnvelopeOverrideArrayOutput) ToEdifactEnvelopeOverrideArrayOutput() EdifactEnvelopeOverrideArrayOutput {
	return o
}

func (o EdifactEnvelopeOverrideArrayOutput) ToEdifactEnvelopeOverrideArrayOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideArrayOutput {
	return o
}

func (o EdifactEnvelopeOverrideArrayOutput) Index(i pulumi.IntInput) EdifactEnvelopeOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactEnvelopeOverride {
		return vs[0].([]EdifactEnvelopeOverride)[vs[1].(int)]
	}).(EdifactEnvelopeOverrideOutput)
}

type EdifactEnvelopeOverrideResponse struct {
	ApplicationPassword            *string `pulumi:"applicationPassword"`
	AssociationAssignedCode        *string `pulumi:"associationAssignedCode"`
	ControllingAgencyCode          *string `pulumi:"controllingAgencyCode"`
	FunctionalGroupId              *string `pulumi:"functionalGroupId"`
	GroupHeaderMessageRelease      *string `pulumi:"groupHeaderMessageRelease"`
	GroupHeaderMessageVersion      *string `pulumi:"groupHeaderMessageVersion"`
	MessageAssociationAssignedCode *string `pulumi:"messageAssociationAssignedCode"`
	MessageId                      *string `pulumi:"messageId"`
	MessageRelease                 *string `pulumi:"messageRelease"`
	MessageVersion                 *string `pulumi:"messageVersion"`
	ReceiverApplicationId          *string `pulumi:"receiverApplicationId"`
	ReceiverApplicationQualifier   *string `pulumi:"receiverApplicationQualifier"`
	SenderApplicationId            *string `pulumi:"senderApplicationId"`
	SenderApplicationQualifier     *string `pulumi:"senderApplicationQualifier"`
	TargetNamespace                *string `pulumi:"targetNamespace"`
}

type EdifactEnvelopeOverrideResponseOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeOverrideResponse)(nil)).Elem()
}

func (o EdifactEnvelopeOverrideResponseOutput) ToEdifactEnvelopeOverrideResponseOutput() EdifactEnvelopeOverrideResponseOutput {
	return o
}

func (o EdifactEnvelopeOverrideResponseOutput) ToEdifactEnvelopeOverrideResponseOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideResponseOutput {
	return o
}

func (o EdifactEnvelopeOverrideResponseOutput) ApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.ApplicationPassword }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) AssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.AssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) ControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.ControllingAgencyCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) GroupHeaderMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.GroupHeaderMessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) GroupHeaderMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.GroupHeaderMessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) MessageAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.MessageAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) MessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.MessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) MessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.MessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.ReceiverApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) ReceiverApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.ReceiverApplicationQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) SenderApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.SenderApplicationQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type EdifactEnvelopeOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactEnvelopeOverrideResponse)(nil)).Elem()
}

func (o EdifactEnvelopeOverrideResponseArrayOutput) ToEdifactEnvelopeOverrideResponseArrayOutput() EdifactEnvelopeOverrideResponseArrayOutput {
	return o
}

func (o EdifactEnvelopeOverrideResponseArrayOutput) ToEdifactEnvelopeOverrideResponseArrayOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideResponseArrayOutput {
	return o
}

func (o EdifactEnvelopeOverrideResponseArrayOutput) Index(i pulumi.IntInput) EdifactEnvelopeOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactEnvelopeOverrideResponse {
		return vs[0].([]EdifactEnvelopeOverrideResponse)[vs[1].(int)]
	}).(EdifactEnvelopeOverrideResponseOutput)
}

type EdifactEnvelopeSettings struct {
	ApplicationReferenceId                       *string  `pulumi:"applicationReferenceId"`
	ApplyDelimiterStringAdvice                   *bool    `pulumi:"applyDelimiterStringAdvice"`
	CommunicationAgreementId                     *string  `pulumi:"communicationAgreementId"`
	CreateGroupingSegments                       *bool    `pulumi:"createGroupingSegments"`
	EnableDefaultGroupHeaders                    *bool    `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            *string  `pulumi:"functionalGroupId"`
	GroupApplicationPassword                     *string  `pulumi:"groupApplicationPassword"`
	GroupApplicationReceiverId                   *string  `pulumi:"groupApplicationReceiverId"`
	GroupApplicationReceiverQualifier            *string  `pulumi:"groupApplicationReceiverQualifier"`
	GroupApplicationSenderId                     *string  `pulumi:"groupApplicationSenderId"`
	GroupApplicationSenderQualifier              *string  `pulumi:"groupApplicationSenderQualifier"`
	GroupAssociationAssignedCode                 *string  `pulumi:"groupAssociationAssignedCode"`
	GroupControlNumberLowerBound                 *float64 `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberPrefix                     *string  `pulumi:"groupControlNumberPrefix"`
	GroupControlNumberSuffix                     *string  `pulumi:"groupControlNumberSuffix"`
	GroupControlNumberUpperBound                 *float64 `pulumi:"groupControlNumberUpperBound"`
	GroupControllingAgencyCode                   *string  `pulumi:"groupControllingAgencyCode"`
	GroupMessageRelease                          *string  `pulumi:"groupMessageRelease"`
	GroupMessageVersion                          *string  `pulumi:"groupMessageVersion"`
	InterchangeControlNumberLowerBound           *float64 `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberPrefix               *string  `pulumi:"interchangeControlNumberPrefix"`
	InterchangeControlNumberSuffix               *string  `pulumi:"interchangeControlNumberSuffix"`
	InterchangeControlNumberUpperBound           *float64 `pulumi:"interchangeControlNumberUpperBound"`
	IsTestInterchange                            *bool    `pulumi:"isTestInterchange"`
	OverwriteExistingTransactionSetControlNumber *bool    `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ProcessingPriorityCode                       *string  `pulumi:"processingPriorityCode"`
	ReceiverInternalIdentification               *string  `pulumi:"receiverInternalIdentification"`
	ReceiverInternalSubIdentification            *string  `pulumi:"receiverInternalSubIdentification"`
	ReceiverReverseRoutingAddress                *string  `pulumi:"receiverReverseRoutingAddress"`
	RecipientReferencePasswordQualifier          *string  `pulumi:"recipientReferencePasswordQualifier"`
	RecipientReferencePasswordValue              *string  `pulumi:"recipientReferencePasswordValue"`
	RolloverGroupControlNumber                   *bool    `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             *bool    `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          *bool    `pulumi:"rolloverTransactionSetControlNumber"`
	SenderInternalIdentification                 *string  `pulumi:"senderInternalIdentification"`
	SenderInternalSubIdentification              *string  `pulumi:"senderInternalSubIdentification"`
	SenderReverseRoutingAddress                  *string  `pulumi:"senderReverseRoutingAddress"`
	TransactionSetControlNumberLowerBound        *float64 `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            *string  `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            *string  `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        *float64 `pulumi:"transactionSetControlNumberUpperBound"`
}





type EdifactEnvelopeSettingsInput interface {
	pulumi.Input

	ToEdifactEnvelopeSettingsOutput() EdifactEnvelopeSettingsOutput
	ToEdifactEnvelopeSettingsOutputWithContext(context.Context) EdifactEnvelopeSettingsOutput
}

type EdifactEnvelopeSettingsArgs struct {
	ApplicationReferenceId                       pulumi.StringPtrInput  `pulumi:"applicationReferenceId"`
	ApplyDelimiterStringAdvice                   pulumi.BoolPtrInput    `pulumi:"applyDelimiterStringAdvice"`
	CommunicationAgreementId                     pulumi.StringPtrInput  `pulumi:"communicationAgreementId"`
	CreateGroupingSegments                       pulumi.BoolPtrInput    `pulumi:"createGroupingSegments"`
	EnableDefaultGroupHeaders                    pulumi.BoolPtrInput    `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            pulumi.StringPtrInput  `pulumi:"functionalGroupId"`
	GroupApplicationPassword                     pulumi.StringPtrInput  `pulumi:"groupApplicationPassword"`
	GroupApplicationReceiverId                   pulumi.StringPtrInput  `pulumi:"groupApplicationReceiverId"`
	GroupApplicationReceiverQualifier            pulumi.StringPtrInput  `pulumi:"groupApplicationReceiverQualifier"`
	GroupApplicationSenderId                     pulumi.StringPtrInput  `pulumi:"groupApplicationSenderId"`
	GroupApplicationSenderQualifier              pulumi.StringPtrInput  `pulumi:"groupApplicationSenderQualifier"`
	GroupAssociationAssignedCode                 pulumi.StringPtrInput  `pulumi:"groupAssociationAssignedCode"`
	GroupControlNumberLowerBound                 pulumi.Float64PtrInput `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberPrefix                     pulumi.StringPtrInput  `pulumi:"groupControlNumberPrefix"`
	GroupControlNumberSuffix                     pulumi.StringPtrInput  `pulumi:"groupControlNumberSuffix"`
	GroupControlNumberUpperBound                 pulumi.Float64PtrInput `pulumi:"groupControlNumberUpperBound"`
	GroupControllingAgencyCode                   pulumi.StringPtrInput  `pulumi:"groupControllingAgencyCode"`
	GroupMessageRelease                          pulumi.StringPtrInput  `pulumi:"groupMessageRelease"`
	GroupMessageVersion                          pulumi.StringPtrInput  `pulumi:"groupMessageVersion"`
	InterchangeControlNumberLowerBound           pulumi.Float64PtrInput `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberPrefix               pulumi.StringPtrInput  `pulumi:"interchangeControlNumberPrefix"`
	InterchangeControlNumberSuffix               pulumi.StringPtrInput  `pulumi:"interchangeControlNumberSuffix"`
	InterchangeControlNumberUpperBound           pulumi.Float64PtrInput `pulumi:"interchangeControlNumberUpperBound"`
	IsTestInterchange                            pulumi.BoolPtrInput    `pulumi:"isTestInterchange"`
	OverwriteExistingTransactionSetControlNumber pulumi.BoolPtrInput    `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ProcessingPriorityCode                       pulumi.StringPtrInput  `pulumi:"processingPriorityCode"`
	ReceiverInternalIdentification               pulumi.StringPtrInput  `pulumi:"receiverInternalIdentification"`
	ReceiverInternalSubIdentification            pulumi.StringPtrInput  `pulumi:"receiverInternalSubIdentification"`
	ReceiverReverseRoutingAddress                pulumi.StringPtrInput  `pulumi:"receiverReverseRoutingAddress"`
	RecipientReferencePasswordQualifier          pulumi.StringPtrInput  `pulumi:"recipientReferencePasswordQualifier"`
	RecipientReferencePasswordValue              pulumi.StringPtrInput  `pulumi:"recipientReferencePasswordValue"`
	RolloverGroupControlNumber                   pulumi.BoolPtrInput    `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             pulumi.BoolPtrInput    `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          pulumi.BoolPtrInput    `pulumi:"rolloverTransactionSetControlNumber"`
	SenderInternalIdentification                 pulumi.StringPtrInput  `pulumi:"senderInternalIdentification"`
	SenderInternalSubIdentification              pulumi.StringPtrInput  `pulumi:"senderInternalSubIdentification"`
	SenderReverseRoutingAddress                  pulumi.StringPtrInput  `pulumi:"senderReverseRoutingAddress"`
	TransactionSetControlNumberLowerBound        pulumi.Float64PtrInput `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            pulumi.StringPtrInput  `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            pulumi.StringPtrInput  `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        pulumi.Float64PtrInput `pulumi:"transactionSetControlNumberUpperBound"`
}

func (EdifactEnvelopeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeSettings)(nil)).Elem()
}

func (i EdifactEnvelopeSettingsArgs) ToEdifactEnvelopeSettingsOutput() EdifactEnvelopeSettingsOutput {
	return i.ToEdifactEnvelopeSettingsOutputWithContext(context.Background())
}

func (i EdifactEnvelopeSettingsArgs) ToEdifactEnvelopeSettingsOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactEnvelopeSettingsOutput)
}

func (i EdifactEnvelopeSettingsArgs) ToEdifactEnvelopeSettingsPtrOutput() EdifactEnvelopeSettingsPtrOutput {
	return i.ToEdifactEnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactEnvelopeSettingsArgs) ToEdifactEnvelopeSettingsPtrOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactEnvelopeSettingsOutput).ToEdifactEnvelopeSettingsPtrOutputWithContext(ctx)
}









type EdifactEnvelopeSettingsPtrInput interface {
	pulumi.Input

	ToEdifactEnvelopeSettingsPtrOutput() EdifactEnvelopeSettingsPtrOutput
	ToEdifactEnvelopeSettingsPtrOutputWithContext(context.Context) EdifactEnvelopeSettingsPtrOutput
}

type edifactEnvelopeSettingsPtrType EdifactEnvelopeSettingsArgs

func EdifactEnvelopeSettingsPtr(v *EdifactEnvelopeSettingsArgs) EdifactEnvelopeSettingsPtrInput {
	return (*edifactEnvelopeSettingsPtrType)(v)
}

func (*edifactEnvelopeSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactEnvelopeSettings)(nil)).Elem()
}

func (i *edifactEnvelopeSettingsPtrType) ToEdifactEnvelopeSettingsPtrOutput() EdifactEnvelopeSettingsPtrOutput {
	return i.ToEdifactEnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactEnvelopeSettingsPtrType) ToEdifactEnvelopeSettingsPtrOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactEnvelopeSettingsPtrOutput)
}

type EdifactEnvelopeSettingsOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeSettings)(nil)).Elem()
}

func (o EdifactEnvelopeSettingsOutput) ToEdifactEnvelopeSettingsOutput() EdifactEnvelopeSettingsOutput {
	return o
}

func (o EdifactEnvelopeSettingsOutput) ToEdifactEnvelopeSettingsOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsOutput {
	return o
}

func (o EdifactEnvelopeSettingsOutput) ToEdifactEnvelopeSettingsPtrOutput() EdifactEnvelopeSettingsPtrOutput {
	return o.ToEdifactEnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactEnvelopeSettingsOutput) ToEdifactEnvelopeSettingsPtrOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactEnvelopeSettings) *EdifactEnvelopeSettings {
		return &v
	}).(EdifactEnvelopeSettingsPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ApplicationReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.ApplicationReferenceId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ApplyDelimiterStringAdvice() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *bool { return v.ApplyDelimiterStringAdvice }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) CommunicationAgreementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.CommunicationAgreementId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) CreateGroupingSegments() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *bool { return v.CreateGroupingSegments }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *bool { return v.EnableDefaultGroupHeaders }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupApplicationPassword }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupApplicationReceiverId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupApplicationReceiverId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupApplicationReceiverQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupApplicationReceiverQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupApplicationSenderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupApplicationSenderId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupApplicationSenderQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupApplicationSenderQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *float64 { return v.GroupControlNumberLowerBound }).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *float64 { return v.GroupControlNumberUpperBound }).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupControllingAgencyCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupMessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupMessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) InterchangeControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *float64 { return v.InterchangeControlNumberLowerBound }).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) InterchangeControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.InterchangeControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) InterchangeControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.InterchangeControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) InterchangeControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *float64 { return v.InterchangeControlNumberUpperBound }).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) IsTestInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *bool { return v.IsTestInterchange }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *bool { return v.OverwriteExistingTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ProcessingPriorityCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.ProcessingPriorityCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ReceiverInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.ReceiverInternalIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ReceiverInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.ReceiverInternalSubIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ReceiverReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.ReceiverReverseRoutingAddress }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) RecipientReferencePasswordQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.RecipientReferencePasswordQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) RecipientReferencePasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.RecipientReferencePasswordValue }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *bool { return v.RolloverGroupControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *bool { return v.RolloverInterchangeControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *bool { return v.RolloverTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) SenderInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.SenderInternalIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) SenderInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.SenderInternalSubIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) SenderReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.SenderReverseRoutingAddress }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) TransactionSetControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *float64 { return v.TransactionSetControlNumberLowerBound }).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.TransactionSetControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.TransactionSetControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) TransactionSetControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *float64 { return v.TransactionSetControlNumberUpperBound }).(pulumi.Float64PtrOutput)
}

type EdifactEnvelopeSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactEnvelopeSettings)(nil)).Elem()
}

func (o EdifactEnvelopeSettingsPtrOutput) ToEdifactEnvelopeSettingsPtrOutput() EdifactEnvelopeSettingsPtrOutput {
	return o
}

func (o EdifactEnvelopeSettingsPtrOutput) ToEdifactEnvelopeSettingsPtrOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsPtrOutput {
	return o
}

func (o EdifactEnvelopeSettingsPtrOutput) Elem() EdifactEnvelopeSettingsOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) EdifactEnvelopeSettings {
		if v != nil {
			return *v
		}
		var ret EdifactEnvelopeSettings
		return ret
	}).(EdifactEnvelopeSettingsOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ApplicationReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationReferenceId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ApplyDelimiterStringAdvice() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ApplyDelimiterStringAdvice
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) CommunicationAgreementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.CommunicationAgreementId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) CreateGroupingSegments() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CreateGroupingSegments
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDefaultGroupHeaders
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalGroupId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationPassword
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupApplicationReceiverId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationReceiverId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupApplicationReceiverQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationReceiverQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupApplicationSenderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationSenderId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupApplicationSenderQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationSenderQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupAssociationAssignedCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupControllingAgencyCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupMessageRelease
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupMessageVersion
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) InterchangeControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) InterchangeControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) InterchangeControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) InterchangeControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) IsTestInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.IsTestInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.OverwriteExistingTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ProcessingPriorityCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ProcessingPriorityCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ReceiverInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverInternalIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ReceiverInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverInternalSubIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ReceiverReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverReverseRoutingAddress
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) RecipientReferencePasswordQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.RecipientReferencePasswordQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) RecipientReferencePasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.RecipientReferencePasswordValue
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) SenderInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.SenderInternalIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) SenderInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.SenderInternalSubIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) SenderReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.SenderReverseRoutingAddress
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) TransactionSetControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) TransactionSetControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

type EdifactEnvelopeSettingsResponse struct {
	ApplicationReferenceId                       *string  `pulumi:"applicationReferenceId"`
	ApplyDelimiterStringAdvice                   *bool    `pulumi:"applyDelimiterStringAdvice"`
	CommunicationAgreementId                     *string  `pulumi:"communicationAgreementId"`
	CreateGroupingSegments                       *bool    `pulumi:"createGroupingSegments"`
	EnableDefaultGroupHeaders                    *bool    `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            *string  `pulumi:"functionalGroupId"`
	GroupApplicationPassword                     *string  `pulumi:"groupApplicationPassword"`
	GroupApplicationReceiverId                   *string  `pulumi:"groupApplicationReceiverId"`
	GroupApplicationReceiverQualifier            *string  `pulumi:"groupApplicationReceiverQualifier"`
	GroupApplicationSenderId                     *string  `pulumi:"groupApplicationSenderId"`
	GroupApplicationSenderQualifier              *string  `pulumi:"groupApplicationSenderQualifier"`
	GroupAssociationAssignedCode                 *string  `pulumi:"groupAssociationAssignedCode"`
	GroupControlNumberLowerBound                 *float64 `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberPrefix                     *string  `pulumi:"groupControlNumberPrefix"`
	GroupControlNumberSuffix                     *string  `pulumi:"groupControlNumberSuffix"`
	GroupControlNumberUpperBound                 *float64 `pulumi:"groupControlNumberUpperBound"`
	GroupControllingAgencyCode                   *string  `pulumi:"groupControllingAgencyCode"`
	GroupMessageRelease                          *string  `pulumi:"groupMessageRelease"`
	GroupMessageVersion                          *string  `pulumi:"groupMessageVersion"`
	InterchangeControlNumberLowerBound           *float64 `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberPrefix               *string  `pulumi:"interchangeControlNumberPrefix"`
	InterchangeControlNumberSuffix               *string  `pulumi:"interchangeControlNumberSuffix"`
	InterchangeControlNumberUpperBound           *float64 `pulumi:"interchangeControlNumberUpperBound"`
	IsTestInterchange                            *bool    `pulumi:"isTestInterchange"`
	OverwriteExistingTransactionSetControlNumber *bool    `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ProcessingPriorityCode                       *string  `pulumi:"processingPriorityCode"`
	ReceiverInternalIdentification               *string  `pulumi:"receiverInternalIdentification"`
	ReceiverInternalSubIdentification            *string  `pulumi:"receiverInternalSubIdentification"`
	ReceiverReverseRoutingAddress                *string  `pulumi:"receiverReverseRoutingAddress"`
	RecipientReferencePasswordQualifier          *string  `pulumi:"recipientReferencePasswordQualifier"`
	RecipientReferencePasswordValue              *string  `pulumi:"recipientReferencePasswordValue"`
	RolloverGroupControlNumber                   *bool    `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             *bool    `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          *bool    `pulumi:"rolloverTransactionSetControlNumber"`
	SenderInternalIdentification                 *string  `pulumi:"senderInternalIdentification"`
	SenderInternalSubIdentification              *string  `pulumi:"senderInternalSubIdentification"`
	SenderReverseRoutingAddress                  *string  `pulumi:"senderReverseRoutingAddress"`
	TransactionSetControlNumberLowerBound        *float64 `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            *string  `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            *string  `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        *float64 `pulumi:"transactionSetControlNumberUpperBound"`
}

type EdifactEnvelopeSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeSettingsResponse)(nil)).Elem()
}

func (o EdifactEnvelopeSettingsResponseOutput) ToEdifactEnvelopeSettingsResponseOutput() EdifactEnvelopeSettingsResponseOutput {
	return o
}

func (o EdifactEnvelopeSettingsResponseOutput) ToEdifactEnvelopeSettingsResponseOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsResponseOutput {
	return o
}

func (o EdifactEnvelopeSettingsResponseOutput) ApplicationReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.ApplicationReferenceId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) ApplyDelimiterStringAdvice() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *bool { return v.ApplyDelimiterStringAdvice }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) CommunicationAgreementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.CommunicationAgreementId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) CreateGroupingSegments() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *bool { return v.CreateGroupingSegments }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *bool { return v.EnableDefaultGroupHeaders }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupApplicationPassword }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupApplicationReceiverId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupApplicationReceiverId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupApplicationReceiverQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupApplicationReceiverQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupApplicationSenderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupApplicationSenderId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupApplicationSenderQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupApplicationSenderQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *float64 { return v.GroupControlNumberLowerBound }).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *float64 { return v.GroupControlNumberUpperBound }).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupControllingAgencyCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupMessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupMessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) InterchangeControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *float64 { return v.InterchangeControlNumberLowerBound }).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) InterchangeControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.InterchangeControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) InterchangeControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.InterchangeControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) InterchangeControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *float64 { return v.InterchangeControlNumberUpperBound }).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) IsTestInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *bool { return v.IsTestInterchange }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *bool { return v.OverwriteExistingTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) ProcessingPriorityCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.ProcessingPriorityCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) ReceiverInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.ReceiverInternalIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) ReceiverInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.ReceiverInternalSubIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) ReceiverReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.ReceiverReverseRoutingAddress }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) RecipientReferencePasswordQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.RecipientReferencePasswordQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) RecipientReferencePasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.RecipientReferencePasswordValue }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *bool { return v.RolloverGroupControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *bool { return v.RolloverInterchangeControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *bool { return v.RolloverTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) SenderInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.SenderInternalIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) SenderInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.SenderInternalSubIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) SenderReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.SenderReverseRoutingAddress }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) TransactionSetControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *float64 { return v.TransactionSetControlNumberLowerBound }).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.TransactionSetControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.TransactionSetControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) TransactionSetControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *float64 { return v.TransactionSetControlNumberUpperBound }).(pulumi.Float64PtrOutput)
}

type EdifactEnvelopeSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactEnvelopeSettingsResponse)(nil)).Elem()
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ToEdifactEnvelopeSettingsResponsePtrOutput() EdifactEnvelopeSettingsResponsePtrOutput {
	return o
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ToEdifactEnvelopeSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsResponsePtrOutput {
	return o
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) Elem() EdifactEnvelopeSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) EdifactEnvelopeSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactEnvelopeSettingsResponse
		return ret
	}).(EdifactEnvelopeSettingsResponseOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ApplicationReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationReferenceId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ApplyDelimiterStringAdvice() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ApplyDelimiterStringAdvice
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) CommunicationAgreementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.CommunicationAgreementId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) CreateGroupingSegments() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CreateGroupingSegments
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDefaultGroupHeaders
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalGroupId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationPassword
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupApplicationReceiverId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationReceiverId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupApplicationReceiverQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationReceiverQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupApplicationSenderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationSenderId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupApplicationSenderQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationSenderQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupAssociationAssignedCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupControllingAgencyCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupMessageRelease
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupMessageVersion
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) InterchangeControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) InterchangeControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) InterchangeControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) InterchangeControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) IsTestInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsTestInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.OverwriteExistingTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ProcessingPriorityCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProcessingPriorityCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ReceiverInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverInternalIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ReceiverInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverInternalSubIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ReceiverReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverReverseRoutingAddress
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) RecipientReferencePasswordQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecipientReferencePasswordQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) RecipientReferencePasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecipientReferencePasswordValue
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) SenderInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SenderInternalIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) SenderInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SenderInternalSubIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) SenderReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SenderReverseRoutingAddress
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

type EdifactFramingSettings struct {
	CharacterEncoding               *string                  `pulumi:"characterEncoding"`
	CharacterSet                    *EdifactCharacterSet     `pulumi:"characterSet"`
	ComponentSeparator              *int                     `pulumi:"componentSeparator"`
	DataElementSeparator            *int                     `pulumi:"dataElementSeparator"`
	DecimalPointIndicator           *EdifactDecimalIndicator `pulumi:"decimalPointIndicator"`
	ProtocolVersion                 *int                     `pulumi:"protocolVersion"`
	ReleaseIndicator                *int                     `pulumi:"releaseIndicator"`
	RepetitionSeparator             *int                     `pulumi:"repetitionSeparator"`
	SegmentTerminator               *int                     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix         *SegmentTerminatorSuffix `pulumi:"segmentTerminatorSuffix"`
	ServiceCodeListDirectoryVersion *string                  `pulumi:"serviceCodeListDirectoryVersion"`
}





type EdifactFramingSettingsInput interface {
	pulumi.Input

	ToEdifactFramingSettingsOutput() EdifactFramingSettingsOutput
	ToEdifactFramingSettingsOutputWithContext(context.Context) EdifactFramingSettingsOutput
}

type EdifactFramingSettingsArgs struct {
	CharacterEncoding               pulumi.StringPtrInput           `pulumi:"characterEncoding"`
	CharacterSet                    EdifactCharacterSetPtrInput     `pulumi:"characterSet"`
	ComponentSeparator              pulumi.IntPtrInput              `pulumi:"componentSeparator"`
	DataElementSeparator            pulumi.IntPtrInput              `pulumi:"dataElementSeparator"`
	DecimalPointIndicator           EdifactDecimalIndicatorPtrInput `pulumi:"decimalPointIndicator"`
	ProtocolVersion                 pulumi.IntPtrInput              `pulumi:"protocolVersion"`
	ReleaseIndicator                pulumi.IntPtrInput              `pulumi:"releaseIndicator"`
	RepetitionSeparator             pulumi.IntPtrInput              `pulumi:"repetitionSeparator"`
	SegmentTerminator               pulumi.IntPtrInput              `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix         SegmentTerminatorSuffixPtrInput `pulumi:"segmentTerminatorSuffix"`
	ServiceCodeListDirectoryVersion pulumi.StringPtrInput           `pulumi:"serviceCodeListDirectoryVersion"`
}

func (EdifactFramingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactFramingSettings)(nil)).Elem()
}

func (i EdifactFramingSettingsArgs) ToEdifactFramingSettingsOutput() EdifactFramingSettingsOutput {
	return i.ToEdifactFramingSettingsOutputWithContext(context.Background())
}

func (i EdifactFramingSettingsArgs) ToEdifactFramingSettingsOutputWithContext(ctx context.Context) EdifactFramingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactFramingSettingsOutput)
}

func (i EdifactFramingSettingsArgs) ToEdifactFramingSettingsPtrOutput() EdifactFramingSettingsPtrOutput {
	return i.ToEdifactFramingSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactFramingSettingsArgs) ToEdifactFramingSettingsPtrOutputWithContext(ctx context.Context) EdifactFramingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactFramingSettingsOutput).ToEdifactFramingSettingsPtrOutputWithContext(ctx)
}









type EdifactFramingSettingsPtrInput interface {
	pulumi.Input

	ToEdifactFramingSettingsPtrOutput() EdifactFramingSettingsPtrOutput
	ToEdifactFramingSettingsPtrOutputWithContext(context.Context) EdifactFramingSettingsPtrOutput
}

type edifactFramingSettingsPtrType EdifactFramingSettingsArgs

func EdifactFramingSettingsPtr(v *EdifactFramingSettingsArgs) EdifactFramingSettingsPtrInput {
	return (*edifactFramingSettingsPtrType)(v)
}

func (*edifactFramingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactFramingSettings)(nil)).Elem()
}

func (i *edifactFramingSettingsPtrType) ToEdifactFramingSettingsPtrOutput() EdifactFramingSettingsPtrOutput {
	return i.ToEdifactFramingSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactFramingSettingsPtrType) ToEdifactFramingSettingsPtrOutputWithContext(ctx context.Context) EdifactFramingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactFramingSettingsPtrOutput)
}

type EdifactFramingSettingsOutput struct{ *pulumi.OutputState }

func (EdifactFramingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactFramingSettings)(nil)).Elem()
}

func (o EdifactFramingSettingsOutput) ToEdifactFramingSettingsOutput() EdifactFramingSettingsOutput {
	return o
}

func (o EdifactFramingSettingsOutput) ToEdifactFramingSettingsOutputWithContext(ctx context.Context) EdifactFramingSettingsOutput {
	return o
}

func (o EdifactFramingSettingsOutput) ToEdifactFramingSettingsPtrOutput() EdifactFramingSettingsPtrOutput {
	return o.ToEdifactFramingSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactFramingSettingsOutput) ToEdifactFramingSettingsPtrOutputWithContext(ctx context.Context) EdifactFramingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactFramingSettings) *EdifactFramingSettings {
		return &v
	}).(EdifactFramingSettingsPtrOutput)
}

func (o EdifactFramingSettingsOutput) CharacterEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *string { return v.CharacterEncoding }).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsOutput) CharacterSet() EdifactCharacterSetPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *EdifactCharacterSet { return v.CharacterSet }).(EdifactCharacterSetPtrOutput)
}

func (o EdifactFramingSettingsOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *int { return v.ComponentSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *int { return v.DataElementSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsOutput) DecimalPointIndicator() EdifactDecimalIndicatorPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *EdifactDecimalIndicator { return v.DecimalPointIndicator }).(EdifactDecimalIndicatorPtrOutput)
}

func (o EdifactFramingSettingsOutput) ProtocolVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *int { return v.ProtocolVersion }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsOutput) ReleaseIndicator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *int { return v.ReleaseIndicator }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsOutput) RepetitionSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *int { return v.RepetitionSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *int { return v.SegmentTerminator }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *SegmentTerminatorSuffix { return v.SegmentTerminatorSuffix }).(SegmentTerminatorSuffixPtrOutput)
}

func (o EdifactFramingSettingsOutput) ServiceCodeListDirectoryVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *string { return v.ServiceCodeListDirectoryVersion }).(pulumi.StringPtrOutput)
}

type EdifactFramingSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactFramingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactFramingSettings)(nil)).Elem()
}

func (o EdifactFramingSettingsPtrOutput) ToEdifactFramingSettingsPtrOutput() EdifactFramingSettingsPtrOutput {
	return o
}

func (o EdifactFramingSettingsPtrOutput) ToEdifactFramingSettingsPtrOutputWithContext(ctx context.Context) EdifactFramingSettingsPtrOutput {
	return o
}

func (o EdifactFramingSettingsPtrOutput) Elem() EdifactFramingSettingsOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) EdifactFramingSettings {
		if v != nil {
			return *v
		}
		var ret EdifactFramingSettings
		return ret
	}).(EdifactFramingSettingsOutput)
}

func (o EdifactFramingSettingsPtrOutput) CharacterEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *string {
		if v == nil {
			return nil
		}
		return v.CharacterEncoding
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) CharacterSet() EdifactCharacterSetPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *EdifactCharacterSet {
		if v == nil {
			return nil
		}
		return v.CharacterSet
	}).(EdifactCharacterSetPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return v.ComponentSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return v.DataElementSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) DecimalPointIndicator() EdifactDecimalIndicatorPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *EdifactDecimalIndicator {
		if v == nil {
			return nil
		}
		return v.DecimalPointIndicator
	}).(EdifactDecimalIndicatorPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) ProtocolVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return v.ProtocolVersion
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) ReleaseIndicator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return v.ReleaseIndicator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) RepetitionSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return v.RepetitionSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return v.SegmentTerminator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *SegmentTerminatorSuffix {
		if v == nil {
			return nil
		}
		return v.SegmentTerminatorSuffix
	}).(SegmentTerminatorSuffixPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) ServiceCodeListDirectoryVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCodeListDirectoryVersion
	}).(pulumi.StringPtrOutput)
}

type EdifactFramingSettingsResponse struct {
	CharacterEncoding               *string `pulumi:"characterEncoding"`
	CharacterSet                    *string `pulumi:"characterSet"`
	ComponentSeparator              *int    `pulumi:"componentSeparator"`
	DataElementSeparator            *int    `pulumi:"dataElementSeparator"`
	DecimalPointIndicator           *string `pulumi:"decimalPointIndicator"`
	ProtocolVersion                 *int    `pulumi:"protocolVersion"`
	ReleaseIndicator                *int    `pulumi:"releaseIndicator"`
	RepetitionSeparator             *int    `pulumi:"repetitionSeparator"`
	SegmentTerminator               *int    `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix         *string `pulumi:"segmentTerminatorSuffix"`
	ServiceCodeListDirectoryVersion *string `pulumi:"serviceCodeListDirectoryVersion"`
}

type EdifactFramingSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactFramingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactFramingSettingsResponse)(nil)).Elem()
}

func (o EdifactFramingSettingsResponseOutput) ToEdifactFramingSettingsResponseOutput() EdifactFramingSettingsResponseOutput {
	return o
}

func (o EdifactFramingSettingsResponseOutput) ToEdifactFramingSettingsResponseOutputWithContext(ctx context.Context) EdifactFramingSettingsResponseOutput {
	return o
}

func (o EdifactFramingSettingsResponseOutput) CharacterEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *string { return v.CharacterEncoding }).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) CharacterSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *string { return v.CharacterSet }).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *int { return v.ComponentSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *int { return v.DataElementSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) DecimalPointIndicator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *string { return v.DecimalPointIndicator }).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) ProtocolVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *int { return v.ProtocolVersion }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) ReleaseIndicator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *int { return v.ReleaseIndicator }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) RepetitionSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *int { return v.RepetitionSeparator }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *int { return v.SegmentTerminator }).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) SegmentTerminatorSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *string { return v.SegmentTerminatorSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) ServiceCodeListDirectoryVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *string { return v.ServiceCodeListDirectoryVersion }).(pulumi.StringPtrOutput)
}

type EdifactFramingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactFramingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactFramingSettingsResponse)(nil)).Elem()
}

func (o EdifactFramingSettingsResponsePtrOutput) ToEdifactFramingSettingsResponsePtrOutput() EdifactFramingSettingsResponsePtrOutput {
	return o
}

func (o EdifactFramingSettingsResponsePtrOutput) ToEdifactFramingSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactFramingSettingsResponsePtrOutput {
	return o
}

func (o EdifactFramingSettingsResponsePtrOutput) Elem() EdifactFramingSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) EdifactFramingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactFramingSettingsResponse
		return ret
	}).(EdifactFramingSettingsResponseOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) CharacterEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.CharacterEncoding
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) CharacterSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.CharacterSet
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.ComponentSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.DataElementSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) DecimalPointIndicator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DecimalPointIndicator
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) ProtocolVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.ProtocolVersion
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) ReleaseIndicator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.ReleaseIndicator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) RepetitionSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.RepetitionSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.SegmentTerminator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) SegmentTerminatorSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SegmentTerminatorSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) ServiceCodeListDirectoryVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCodeListDirectoryVersion
	}).(pulumi.StringPtrOutput)
}

type EdifactMessageFilter struct {
	MessageFilterType *MessageFilterType `pulumi:"messageFilterType"`
}





type EdifactMessageFilterInput interface {
	pulumi.Input

	ToEdifactMessageFilterOutput() EdifactMessageFilterOutput
	ToEdifactMessageFilterOutputWithContext(context.Context) EdifactMessageFilterOutput
}

type EdifactMessageFilterArgs struct {
	MessageFilterType MessageFilterTypePtrInput `pulumi:"messageFilterType"`
}

func (EdifactMessageFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageFilter)(nil)).Elem()
}

func (i EdifactMessageFilterArgs) ToEdifactMessageFilterOutput() EdifactMessageFilterOutput {
	return i.ToEdifactMessageFilterOutputWithContext(context.Background())
}

func (i EdifactMessageFilterArgs) ToEdifactMessageFilterOutputWithContext(ctx context.Context) EdifactMessageFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactMessageFilterOutput)
}

func (i EdifactMessageFilterArgs) ToEdifactMessageFilterPtrOutput() EdifactMessageFilterPtrOutput {
	return i.ToEdifactMessageFilterPtrOutputWithContext(context.Background())
}

func (i EdifactMessageFilterArgs) ToEdifactMessageFilterPtrOutputWithContext(ctx context.Context) EdifactMessageFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactMessageFilterOutput).ToEdifactMessageFilterPtrOutputWithContext(ctx)
}









type EdifactMessageFilterPtrInput interface {
	pulumi.Input

	ToEdifactMessageFilterPtrOutput() EdifactMessageFilterPtrOutput
	ToEdifactMessageFilterPtrOutputWithContext(context.Context) EdifactMessageFilterPtrOutput
}

type edifactMessageFilterPtrType EdifactMessageFilterArgs

func EdifactMessageFilterPtr(v *EdifactMessageFilterArgs) EdifactMessageFilterPtrInput {
	return (*edifactMessageFilterPtrType)(v)
}

func (*edifactMessageFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactMessageFilter)(nil)).Elem()
}

func (i *edifactMessageFilterPtrType) ToEdifactMessageFilterPtrOutput() EdifactMessageFilterPtrOutput {
	return i.ToEdifactMessageFilterPtrOutputWithContext(context.Background())
}

func (i *edifactMessageFilterPtrType) ToEdifactMessageFilterPtrOutputWithContext(ctx context.Context) EdifactMessageFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactMessageFilterPtrOutput)
}

type EdifactMessageFilterOutput struct{ *pulumi.OutputState }

func (EdifactMessageFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageFilter)(nil)).Elem()
}

func (o EdifactMessageFilterOutput) ToEdifactMessageFilterOutput() EdifactMessageFilterOutput {
	return o
}

func (o EdifactMessageFilterOutput) ToEdifactMessageFilterOutputWithContext(ctx context.Context) EdifactMessageFilterOutput {
	return o
}

func (o EdifactMessageFilterOutput) ToEdifactMessageFilterPtrOutput() EdifactMessageFilterPtrOutput {
	return o.ToEdifactMessageFilterPtrOutputWithContext(context.Background())
}

func (o EdifactMessageFilterOutput) ToEdifactMessageFilterPtrOutputWithContext(ctx context.Context) EdifactMessageFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactMessageFilter) *EdifactMessageFilter {
		return &v
	}).(EdifactMessageFilterPtrOutput)
}

func (o EdifactMessageFilterOutput) MessageFilterType() MessageFilterTypePtrOutput {
	return o.ApplyT(func(v EdifactMessageFilter) *MessageFilterType { return v.MessageFilterType }).(MessageFilterTypePtrOutput)
}

type EdifactMessageFilterPtrOutput struct{ *pulumi.OutputState }

func (EdifactMessageFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactMessageFilter)(nil)).Elem()
}

func (o EdifactMessageFilterPtrOutput) ToEdifactMessageFilterPtrOutput() EdifactMessageFilterPtrOutput {
	return o
}

func (o EdifactMessageFilterPtrOutput) ToEdifactMessageFilterPtrOutputWithContext(ctx context.Context) EdifactMessageFilterPtrOutput {
	return o
}

func (o EdifactMessageFilterPtrOutput) Elem() EdifactMessageFilterOutput {
	return o.ApplyT(func(v *EdifactMessageFilter) EdifactMessageFilter {
		if v != nil {
			return *v
		}
		var ret EdifactMessageFilter
		return ret
	}).(EdifactMessageFilterOutput)
}

func (o EdifactMessageFilterPtrOutput) MessageFilterType() MessageFilterTypePtrOutput {
	return o.ApplyT(func(v *EdifactMessageFilter) *MessageFilterType {
		if v == nil {
			return nil
		}
		return v.MessageFilterType
	}).(MessageFilterTypePtrOutput)
}

type EdifactMessageFilterResponse struct {
	MessageFilterType *string `pulumi:"messageFilterType"`
}

type EdifactMessageFilterResponseOutput struct{ *pulumi.OutputState }

func (EdifactMessageFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageFilterResponse)(nil)).Elem()
}

func (o EdifactMessageFilterResponseOutput) ToEdifactMessageFilterResponseOutput() EdifactMessageFilterResponseOutput {
	return o
}

func (o EdifactMessageFilterResponseOutput) ToEdifactMessageFilterResponseOutputWithContext(ctx context.Context) EdifactMessageFilterResponseOutput {
	return o
}

func (o EdifactMessageFilterResponseOutput) MessageFilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactMessageFilterResponse) *string { return v.MessageFilterType }).(pulumi.StringPtrOutput)
}

type EdifactMessageFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactMessageFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactMessageFilterResponse)(nil)).Elem()
}

func (o EdifactMessageFilterResponsePtrOutput) ToEdifactMessageFilterResponsePtrOutput() EdifactMessageFilterResponsePtrOutput {
	return o
}

func (o EdifactMessageFilterResponsePtrOutput) ToEdifactMessageFilterResponsePtrOutputWithContext(ctx context.Context) EdifactMessageFilterResponsePtrOutput {
	return o
}

func (o EdifactMessageFilterResponsePtrOutput) Elem() EdifactMessageFilterResponseOutput {
	return o.ApplyT(func(v *EdifactMessageFilterResponse) EdifactMessageFilterResponse {
		if v != nil {
			return *v
		}
		var ret EdifactMessageFilterResponse
		return ret
	}).(EdifactMessageFilterResponseOutput)
}

func (o EdifactMessageFilterResponsePtrOutput) MessageFilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactMessageFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageFilterType
	}).(pulumi.StringPtrOutput)
}

type EdifactMessageIdentifier struct {
	MessageId *string `pulumi:"messageId"`
}





type EdifactMessageIdentifierInput interface {
	pulumi.Input

	ToEdifactMessageIdentifierOutput() EdifactMessageIdentifierOutput
	ToEdifactMessageIdentifierOutputWithContext(context.Context) EdifactMessageIdentifierOutput
}

type EdifactMessageIdentifierArgs struct {
	MessageId pulumi.StringPtrInput `pulumi:"messageId"`
}

func (EdifactMessageIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageIdentifier)(nil)).Elem()
}

func (i EdifactMessageIdentifierArgs) ToEdifactMessageIdentifierOutput() EdifactMessageIdentifierOutput {
	return i.ToEdifactMessageIdentifierOutputWithContext(context.Background())
}

func (i EdifactMessageIdentifierArgs) ToEdifactMessageIdentifierOutputWithContext(ctx context.Context) EdifactMessageIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactMessageIdentifierOutput)
}





type EdifactMessageIdentifierArrayInput interface {
	pulumi.Input

	ToEdifactMessageIdentifierArrayOutput() EdifactMessageIdentifierArrayOutput
	ToEdifactMessageIdentifierArrayOutputWithContext(context.Context) EdifactMessageIdentifierArrayOutput
}

type EdifactMessageIdentifierArray []EdifactMessageIdentifierInput

func (EdifactMessageIdentifierArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactMessageIdentifier)(nil)).Elem()
}

func (i EdifactMessageIdentifierArray) ToEdifactMessageIdentifierArrayOutput() EdifactMessageIdentifierArrayOutput {
	return i.ToEdifactMessageIdentifierArrayOutputWithContext(context.Background())
}

func (i EdifactMessageIdentifierArray) ToEdifactMessageIdentifierArrayOutputWithContext(ctx context.Context) EdifactMessageIdentifierArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactMessageIdentifierArrayOutput)
}

type EdifactMessageIdentifierOutput struct{ *pulumi.OutputState }

func (EdifactMessageIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageIdentifier)(nil)).Elem()
}

func (o EdifactMessageIdentifierOutput) ToEdifactMessageIdentifierOutput() EdifactMessageIdentifierOutput {
	return o
}

func (o EdifactMessageIdentifierOutput) ToEdifactMessageIdentifierOutputWithContext(ctx context.Context) EdifactMessageIdentifierOutput {
	return o
}

func (o EdifactMessageIdentifierOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactMessageIdentifier) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

type EdifactMessageIdentifierArrayOutput struct{ *pulumi.OutputState }

func (EdifactMessageIdentifierArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactMessageIdentifier)(nil)).Elem()
}

func (o EdifactMessageIdentifierArrayOutput) ToEdifactMessageIdentifierArrayOutput() EdifactMessageIdentifierArrayOutput {
	return o
}

func (o EdifactMessageIdentifierArrayOutput) ToEdifactMessageIdentifierArrayOutputWithContext(ctx context.Context) EdifactMessageIdentifierArrayOutput {
	return o
}

func (o EdifactMessageIdentifierArrayOutput) Index(i pulumi.IntInput) EdifactMessageIdentifierOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactMessageIdentifier {
		return vs[0].([]EdifactMessageIdentifier)[vs[1].(int)]
	}).(EdifactMessageIdentifierOutput)
}

type EdifactMessageIdentifierResponse struct {
	MessageId *string `pulumi:"messageId"`
}

type EdifactMessageIdentifierResponseOutput struct{ *pulumi.OutputState }

func (EdifactMessageIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageIdentifierResponse)(nil)).Elem()
}

func (o EdifactMessageIdentifierResponseOutput) ToEdifactMessageIdentifierResponseOutput() EdifactMessageIdentifierResponseOutput {
	return o
}

func (o EdifactMessageIdentifierResponseOutput) ToEdifactMessageIdentifierResponseOutputWithContext(ctx context.Context) EdifactMessageIdentifierResponseOutput {
	return o
}

func (o EdifactMessageIdentifierResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactMessageIdentifierResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

type EdifactMessageIdentifierResponseArrayOutput struct{ *pulumi.OutputState }

func (EdifactMessageIdentifierResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactMessageIdentifierResponse)(nil)).Elem()
}

func (o EdifactMessageIdentifierResponseArrayOutput) ToEdifactMessageIdentifierResponseArrayOutput() EdifactMessageIdentifierResponseArrayOutput {
	return o
}

func (o EdifactMessageIdentifierResponseArrayOutput) ToEdifactMessageIdentifierResponseArrayOutputWithContext(ctx context.Context) EdifactMessageIdentifierResponseArrayOutput {
	return o
}

func (o EdifactMessageIdentifierResponseArrayOutput) Index(i pulumi.IntInput) EdifactMessageIdentifierResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactMessageIdentifierResponse {
		return vs[0].([]EdifactMessageIdentifierResponse)[vs[1].(int)]
	}).(EdifactMessageIdentifierResponseOutput)
}

type EdifactOneWayAgreement struct {
	ProtocolSettings         *EdifactProtocolSettings `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity *BusinessIdentity        `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   *BusinessIdentity        `pulumi:"senderBusinessIdentity"`
}





type EdifactOneWayAgreementInput interface {
	pulumi.Input

	ToEdifactOneWayAgreementOutput() EdifactOneWayAgreementOutput
	ToEdifactOneWayAgreementOutputWithContext(context.Context) EdifactOneWayAgreementOutput
}

type EdifactOneWayAgreementArgs struct {
	ProtocolSettings         EdifactProtocolSettingsPtrInput `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentityPtrInput        `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentityPtrInput        `pulumi:"senderBusinessIdentity"`
}

func (EdifactOneWayAgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactOneWayAgreement)(nil)).Elem()
}

func (i EdifactOneWayAgreementArgs) ToEdifactOneWayAgreementOutput() EdifactOneWayAgreementOutput {
	return i.ToEdifactOneWayAgreementOutputWithContext(context.Background())
}

func (i EdifactOneWayAgreementArgs) ToEdifactOneWayAgreementOutputWithContext(ctx context.Context) EdifactOneWayAgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactOneWayAgreementOutput)
}

func (i EdifactOneWayAgreementArgs) ToEdifactOneWayAgreementPtrOutput() EdifactOneWayAgreementPtrOutput {
	return i.ToEdifactOneWayAgreementPtrOutputWithContext(context.Background())
}

func (i EdifactOneWayAgreementArgs) ToEdifactOneWayAgreementPtrOutputWithContext(ctx context.Context) EdifactOneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactOneWayAgreementOutput).ToEdifactOneWayAgreementPtrOutputWithContext(ctx)
}









type EdifactOneWayAgreementPtrInput interface {
	pulumi.Input

	ToEdifactOneWayAgreementPtrOutput() EdifactOneWayAgreementPtrOutput
	ToEdifactOneWayAgreementPtrOutputWithContext(context.Context) EdifactOneWayAgreementPtrOutput
}

type edifactOneWayAgreementPtrType EdifactOneWayAgreementArgs

func EdifactOneWayAgreementPtr(v *EdifactOneWayAgreementArgs) EdifactOneWayAgreementPtrInput {
	return (*edifactOneWayAgreementPtrType)(v)
}

func (*edifactOneWayAgreementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactOneWayAgreement)(nil)).Elem()
}

func (i *edifactOneWayAgreementPtrType) ToEdifactOneWayAgreementPtrOutput() EdifactOneWayAgreementPtrOutput {
	return i.ToEdifactOneWayAgreementPtrOutputWithContext(context.Background())
}

func (i *edifactOneWayAgreementPtrType) ToEdifactOneWayAgreementPtrOutputWithContext(ctx context.Context) EdifactOneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactOneWayAgreementPtrOutput)
}

type EdifactOneWayAgreementOutput struct{ *pulumi.OutputState }

func (EdifactOneWayAgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactOneWayAgreement)(nil)).Elem()
}

func (o EdifactOneWayAgreementOutput) ToEdifactOneWayAgreementOutput() EdifactOneWayAgreementOutput {
	return o
}

func (o EdifactOneWayAgreementOutput) ToEdifactOneWayAgreementOutputWithContext(ctx context.Context) EdifactOneWayAgreementOutput {
	return o
}

func (o EdifactOneWayAgreementOutput) ToEdifactOneWayAgreementPtrOutput() EdifactOneWayAgreementPtrOutput {
	return o.ToEdifactOneWayAgreementPtrOutputWithContext(context.Background())
}

func (o EdifactOneWayAgreementOutput) ToEdifactOneWayAgreementPtrOutputWithContext(ctx context.Context) EdifactOneWayAgreementPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactOneWayAgreement) *EdifactOneWayAgreement {
		return &v
	}).(EdifactOneWayAgreementPtrOutput)
}

func (o EdifactOneWayAgreementOutput) ProtocolSettings() EdifactProtocolSettingsPtrOutput {
	return o.ApplyT(func(v EdifactOneWayAgreement) *EdifactProtocolSettings { return v.ProtocolSettings }).(EdifactProtocolSettingsPtrOutput)
}

func (o EdifactOneWayAgreementOutput) ReceiverBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v EdifactOneWayAgreement) *BusinessIdentity { return v.ReceiverBusinessIdentity }).(BusinessIdentityPtrOutput)
}

func (o EdifactOneWayAgreementOutput) SenderBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v EdifactOneWayAgreement) *BusinessIdentity { return v.SenderBusinessIdentity }).(BusinessIdentityPtrOutput)
}

type EdifactOneWayAgreementPtrOutput struct{ *pulumi.OutputState }

func (EdifactOneWayAgreementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactOneWayAgreement)(nil)).Elem()
}

func (o EdifactOneWayAgreementPtrOutput) ToEdifactOneWayAgreementPtrOutput() EdifactOneWayAgreementPtrOutput {
	return o
}

func (o EdifactOneWayAgreementPtrOutput) ToEdifactOneWayAgreementPtrOutputWithContext(ctx context.Context) EdifactOneWayAgreementPtrOutput {
	return o
}

func (o EdifactOneWayAgreementPtrOutput) Elem() EdifactOneWayAgreementOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreement) EdifactOneWayAgreement {
		if v != nil {
			return *v
		}
		var ret EdifactOneWayAgreement
		return ret
	}).(EdifactOneWayAgreementOutput)
}

func (o EdifactOneWayAgreementPtrOutput) ProtocolSettings() EdifactProtocolSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreement) *EdifactProtocolSettings {
		if v == nil {
			return nil
		}
		return v.ProtocolSettings
	}).(EdifactProtocolSettingsPtrOutput)
}

func (o EdifactOneWayAgreementPtrOutput) ReceiverBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return v.ReceiverBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

func (o EdifactOneWayAgreementPtrOutput) SenderBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return v.SenderBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

type EdifactOneWayAgreementResponse struct {
	ProtocolSettings         *EdifactProtocolSettingsResponse `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity *BusinessIdentityResponse        `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   *BusinessIdentityResponse        `pulumi:"senderBusinessIdentity"`
}

type EdifactOneWayAgreementResponseOutput struct{ *pulumi.OutputState }

func (EdifactOneWayAgreementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactOneWayAgreementResponse)(nil)).Elem()
}

func (o EdifactOneWayAgreementResponseOutput) ToEdifactOneWayAgreementResponseOutput() EdifactOneWayAgreementResponseOutput {
	return o
}

func (o EdifactOneWayAgreementResponseOutput) ToEdifactOneWayAgreementResponseOutputWithContext(ctx context.Context) EdifactOneWayAgreementResponseOutput {
	return o
}

func (o EdifactOneWayAgreementResponseOutput) ProtocolSettings() EdifactProtocolSettingsResponsePtrOutput {
	return o.ApplyT(func(v EdifactOneWayAgreementResponse) *EdifactProtocolSettingsResponse { return v.ProtocolSettings }).(EdifactProtocolSettingsResponsePtrOutput)
}

func (o EdifactOneWayAgreementResponseOutput) ReceiverBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v EdifactOneWayAgreementResponse) *BusinessIdentityResponse { return v.ReceiverBusinessIdentity }).(BusinessIdentityResponsePtrOutput)
}

func (o EdifactOneWayAgreementResponseOutput) SenderBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v EdifactOneWayAgreementResponse) *BusinessIdentityResponse { return v.SenderBusinessIdentity }).(BusinessIdentityResponsePtrOutput)
}

type EdifactOneWayAgreementResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactOneWayAgreementResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactOneWayAgreementResponse)(nil)).Elem()
}

func (o EdifactOneWayAgreementResponsePtrOutput) ToEdifactOneWayAgreementResponsePtrOutput() EdifactOneWayAgreementResponsePtrOutput {
	return o
}

func (o EdifactOneWayAgreementResponsePtrOutput) ToEdifactOneWayAgreementResponsePtrOutputWithContext(ctx context.Context) EdifactOneWayAgreementResponsePtrOutput {
	return o
}

func (o EdifactOneWayAgreementResponsePtrOutput) Elem() EdifactOneWayAgreementResponseOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreementResponse) EdifactOneWayAgreementResponse {
		if v != nil {
			return *v
		}
		var ret EdifactOneWayAgreementResponse
		return ret
	}).(EdifactOneWayAgreementResponseOutput)
}

func (o EdifactOneWayAgreementResponsePtrOutput) ProtocolSettings() EdifactProtocolSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreementResponse) *EdifactProtocolSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ProtocolSettings
	}).(EdifactProtocolSettingsResponsePtrOutput)
}

func (o EdifactOneWayAgreementResponsePtrOutput) ReceiverBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return v.ReceiverBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

func (o EdifactOneWayAgreementResponsePtrOutput) SenderBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return v.SenderBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

type EdifactProcessingSettings struct {
	CreateEmptyXmlTagsForTrailingSeparators *bool `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        *bool `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     *bool `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               *bool `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                *bool `pulumi:"useDotAsDecimalSeparator"`
}





type EdifactProcessingSettingsInput interface {
	pulumi.Input

	ToEdifactProcessingSettingsOutput() EdifactProcessingSettingsOutput
	ToEdifactProcessingSettingsOutputWithContext(context.Context) EdifactProcessingSettingsOutput
}

type EdifactProcessingSettingsArgs struct {
	CreateEmptyXmlTagsForTrailingSeparators pulumi.BoolPtrInput `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        pulumi.BoolPtrInput `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     pulumi.BoolPtrInput `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               pulumi.BoolPtrInput `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                pulumi.BoolPtrInput `pulumi:"useDotAsDecimalSeparator"`
}

func (EdifactProcessingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProcessingSettings)(nil)).Elem()
}

func (i EdifactProcessingSettingsArgs) ToEdifactProcessingSettingsOutput() EdifactProcessingSettingsOutput {
	return i.ToEdifactProcessingSettingsOutputWithContext(context.Background())
}

func (i EdifactProcessingSettingsArgs) ToEdifactProcessingSettingsOutputWithContext(ctx context.Context) EdifactProcessingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProcessingSettingsOutput)
}

func (i EdifactProcessingSettingsArgs) ToEdifactProcessingSettingsPtrOutput() EdifactProcessingSettingsPtrOutput {
	return i.ToEdifactProcessingSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactProcessingSettingsArgs) ToEdifactProcessingSettingsPtrOutputWithContext(ctx context.Context) EdifactProcessingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProcessingSettingsOutput).ToEdifactProcessingSettingsPtrOutputWithContext(ctx)
}









type EdifactProcessingSettingsPtrInput interface {
	pulumi.Input

	ToEdifactProcessingSettingsPtrOutput() EdifactProcessingSettingsPtrOutput
	ToEdifactProcessingSettingsPtrOutputWithContext(context.Context) EdifactProcessingSettingsPtrOutput
}

type edifactProcessingSettingsPtrType EdifactProcessingSettingsArgs

func EdifactProcessingSettingsPtr(v *EdifactProcessingSettingsArgs) EdifactProcessingSettingsPtrInput {
	return (*edifactProcessingSettingsPtrType)(v)
}

func (*edifactProcessingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProcessingSettings)(nil)).Elem()
}

func (i *edifactProcessingSettingsPtrType) ToEdifactProcessingSettingsPtrOutput() EdifactProcessingSettingsPtrOutput {
	return i.ToEdifactProcessingSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactProcessingSettingsPtrType) ToEdifactProcessingSettingsPtrOutputWithContext(ctx context.Context) EdifactProcessingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProcessingSettingsPtrOutput)
}

type EdifactProcessingSettingsOutput struct{ *pulumi.OutputState }

func (EdifactProcessingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProcessingSettings)(nil)).Elem()
}

func (o EdifactProcessingSettingsOutput) ToEdifactProcessingSettingsOutput() EdifactProcessingSettingsOutput {
	return o
}

func (o EdifactProcessingSettingsOutput) ToEdifactProcessingSettingsOutputWithContext(ctx context.Context) EdifactProcessingSettingsOutput {
	return o
}

func (o EdifactProcessingSettingsOutput) ToEdifactProcessingSettingsPtrOutput() EdifactProcessingSettingsPtrOutput {
	return o.ToEdifactProcessingSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactProcessingSettingsOutput) ToEdifactProcessingSettingsPtrOutputWithContext(ctx context.Context) EdifactProcessingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactProcessingSettings) *EdifactProcessingSettings {
		return &v
	}).(EdifactProcessingSettingsPtrOutput)
}

func (o EdifactProcessingSettingsOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactProcessingSettings) *bool { return v.CreateEmptyXmlTagsForTrailingSeparators }).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactProcessingSettings) *bool { return v.MaskSecurityInfo }).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactProcessingSettings) *bool { return v.PreserveInterchange }).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactProcessingSettings) *bool { return v.SuspendInterchangeOnError }).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactProcessingSettings) *bool { return v.UseDotAsDecimalSeparator }).(pulumi.BoolPtrOutput)
}

type EdifactProcessingSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactProcessingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProcessingSettings)(nil)).Elem()
}

func (o EdifactProcessingSettingsPtrOutput) ToEdifactProcessingSettingsPtrOutput() EdifactProcessingSettingsPtrOutput {
	return o
}

func (o EdifactProcessingSettingsPtrOutput) ToEdifactProcessingSettingsPtrOutputWithContext(ctx context.Context) EdifactProcessingSettingsPtrOutput {
	return o
}

func (o EdifactProcessingSettingsPtrOutput) Elem() EdifactProcessingSettingsOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) EdifactProcessingSettings {
		if v != nil {
			return *v
		}
		var ret EdifactProcessingSettings
		return ret
	}).(EdifactProcessingSettingsOutput)
}

func (o EdifactProcessingSettingsPtrOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CreateEmptyXmlTagsForTrailingSeparators
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsPtrOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.MaskSecurityInfo
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsPtrOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.PreserveInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsPtrOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SuspendInterchangeOnError
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsPtrOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.UseDotAsDecimalSeparator
	}).(pulumi.BoolPtrOutput)
}

type EdifactProcessingSettingsResponse struct {
	CreateEmptyXmlTagsForTrailingSeparators *bool `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        *bool `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     *bool `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               *bool `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                *bool `pulumi:"useDotAsDecimalSeparator"`
}

type EdifactProcessingSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactProcessingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProcessingSettingsResponse)(nil)).Elem()
}

func (o EdifactProcessingSettingsResponseOutput) ToEdifactProcessingSettingsResponseOutput() EdifactProcessingSettingsResponseOutput {
	return o
}

func (o EdifactProcessingSettingsResponseOutput) ToEdifactProcessingSettingsResponseOutputWithContext(ctx context.Context) EdifactProcessingSettingsResponseOutput {
	return o
}

func (o EdifactProcessingSettingsResponseOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactProcessingSettingsResponse) *bool { return v.CreateEmptyXmlTagsForTrailingSeparators }).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponseOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactProcessingSettingsResponse) *bool { return v.MaskSecurityInfo }).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponseOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactProcessingSettingsResponse) *bool { return v.PreserveInterchange }).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponseOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactProcessingSettingsResponse) *bool { return v.SuspendInterchangeOnError }).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponseOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactProcessingSettingsResponse) *bool { return v.UseDotAsDecimalSeparator }).(pulumi.BoolPtrOutput)
}

type EdifactProcessingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactProcessingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProcessingSettingsResponse)(nil)).Elem()
}

func (o EdifactProcessingSettingsResponsePtrOutput) ToEdifactProcessingSettingsResponsePtrOutput() EdifactProcessingSettingsResponsePtrOutput {
	return o
}

func (o EdifactProcessingSettingsResponsePtrOutput) ToEdifactProcessingSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactProcessingSettingsResponsePtrOutput {
	return o
}

func (o EdifactProcessingSettingsResponsePtrOutput) Elem() EdifactProcessingSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) EdifactProcessingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactProcessingSettingsResponse
		return ret
	}).(EdifactProcessingSettingsResponseOutput)
}

func (o EdifactProcessingSettingsResponsePtrOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CreateEmptyXmlTagsForTrailingSeparators
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponsePtrOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.MaskSecurityInfo
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponsePtrOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PreserveInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponsePtrOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SuspendInterchangeOnError
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponsePtrOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseDotAsDecimalSeparator
	}).(pulumi.BoolPtrOutput)
}

type EdifactProtocolSettings struct {
	AcknowledgementSettings   *EdifactAcknowledgementSettings `pulumi:"acknowledgementSettings"`
	EdifactDelimiterOverrides []EdifactDelimiterOverride      `pulumi:"edifactDelimiterOverrides"`
	EnvelopeOverrides         []EdifactEnvelopeOverride       `pulumi:"envelopeOverrides"`
	EnvelopeSettings          *EdifactEnvelopeSettings        `pulumi:"envelopeSettings"`
	FramingSettings           *EdifactFramingSettings         `pulumi:"framingSettings"`
	MessageFilter             *EdifactMessageFilter           `pulumi:"messageFilter"`
	MessageFilterList         []EdifactMessageIdentifier      `pulumi:"messageFilterList"`
	ProcessingSettings        *EdifactProcessingSettings      `pulumi:"processingSettings"`
	SchemaReferences          []EdifactSchemaReference        `pulumi:"schemaReferences"`
	ValidationOverrides       []EdifactValidationOverride     `pulumi:"validationOverrides"`
	ValidationSettings        *EdifactValidationSettings      `pulumi:"validationSettings"`
}





type EdifactProtocolSettingsInput interface {
	pulumi.Input

	ToEdifactProtocolSettingsOutput() EdifactProtocolSettingsOutput
	ToEdifactProtocolSettingsOutputWithContext(context.Context) EdifactProtocolSettingsOutput
}

type EdifactProtocolSettingsArgs struct {
	AcknowledgementSettings   EdifactAcknowledgementSettingsPtrInput `pulumi:"acknowledgementSettings"`
	EdifactDelimiterOverrides EdifactDelimiterOverrideArrayInput     `pulumi:"edifactDelimiterOverrides"`
	EnvelopeOverrides         EdifactEnvelopeOverrideArrayInput      `pulumi:"envelopeOverrides"`
	EnvelopeSettings          EdifactEnvelopeSettingsPtrInput        `pulumi:"envelopeSettings"`
	FramingSettings           EdifactFramingSettingsPtrInput         `pulumi:"framingSettings"`
	MessageFilter             EdifactMessageFilterPtrInput           `pulumi:"messageFilter"`
	MessageFilterList         EdifactMessageIdentifierArrayInput     `pulumi:"messageFilterList"`
	ProcessingSettings        EdifactProcessingSettingsPtrInput      `pulumi:"processingSettings"`
	SchemaReferences          EdifactSchemaReferenceArrayInput       `pulumi:"schemaReferences"`
	ValidationOverrides       EdifactValidationOverrideArrayInput    `pulumi:"validationOverrides"`
	ValidationSettings        EdifactValidationSettingsPtrInput      `pulumi:"validationSettings"`
}

func (EdifactProtocolSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProtocolSettings)(nil)).Elem()
}

func (i EdifactProtocolSettingsArgs) ToEdifactProtocolSettingsOutput() EdifactProtocolSettingsOutput {
	return i.ToEdifactProtocolSettingsOutputWithContext(context.Background())
}

func (i EdifactProtocolSettingsArgs) ToEdifactProtocolSettingsOutputWithContext(ctx context.Context) EdifactProtocolSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProtocolSettingsOutput)
}

func (i EdifactProtocolSettingsArgs) ToEdifactProtocolSettingsPtrOutput() EdifactProtocolSettingsPtrOutput {
	return i.ToEdifactProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactProtocolSettingsArgs) ToEdifactProtocolSettingsPtrOutputWithContext(ctx context.Context) EdifactProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProtocolSettingsOutput).ToEdifactProtocolSettingsPtrOutputWithContext(ctx)
}









type EdifactProtocolSettingsPtrInput interface {
	pulumi.Input

	ToEdifactProtocolSettingsPtrOutput() EdifactProtocolSettingsPtrOutput
	ToEdifactProtocolSettingsPtrOutputWithContext(context.Context) EdifactProtocolSettingsPtrOutput
}

type edifactProtocolSettingsPtrType EdifactProtocolSettingsArgs

func EdifactProtocolSettingsPtr(v *EdifactProtocolSettingsArgs) EdifactProtocolSettingsPtrInput {
	return (*edifactProtocolSettingsPtrType)(v)
}

func (*edifactProtocolSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProtocolSettings)(nil)).Elem()
}

func (i *edifactProtocolSettingsPtrType) ToEdifactProtocolSettingsPtrOutput() EdifactProtocolSettingsPtrOutput {
	return i.ToEdifactProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactProtocolSettingsPtrType) ToEdifactProtocolSettingsPtrOutputWithContext(ctx context.Context) EdifactProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProtocolSettingsPtrOutput)
}

type EdifactProtocolSettingsOutput struct{ *pulumi.OutputState }

func (EdifactProtocolSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProtocolSettings)(nil)).Elem()
}

func (o EdifactProtocolSettingsOutput) ToEdifactProtocolSettingsOutput() EdifactProtocolSettingsOutput {
	return o
}

func (o EdifactProtocolSettingsOutput) ToEdifactProtocolSettingsOutputWithContext(ctx context.Context) EdifactProtocolSettingsOutput {
	return o
}

func (o EdifactProtocolSettingsOutput) ToEdifactProtocolSettingsPtrOutput() EdifactProtocolSettingsPtrOutput {
	return o.ToEdifactProtocolSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactProtocolSettingsOutput) ToEdifactProtocolSettingsPtrOutputWithContext(ctx context.Context) EdifactProtocolSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactProtocolSettings) *EdifactProtocolSettings {
		return &v
	}).(EdifactProtocolSettingsPtrOutput)
}

func (o EdifactProtocolSettingsOutput) AcknowledgementSettings() EdifactAcknowledgementSettingsPtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) *EdifactAcknowledgementSettings { return v.AcknowledgementSettings }).(EdifactAcknowledgementSettingsPtrOutput)
}

func (o EdifactProtocolSettingsOutput) EdifactDelimiterOverrides() EdifactDelimiterOverrideArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) []EdifactDelimiterOverride { return v.EdifactDelimiterOverrides }).(EdifactDelimiterOverrideArrayOutput)
}

func (o EdifactProtocolSettingsOutput) EnvelopeOverrides() EdifactEnvelopeOverrideArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) []EdifactEnvelopeOverride { return v.EnvelopeOverrides }).(EdifactEnvelopeOverrideArrayOutput)
}

func (o EdifactProtocolSettingsOutput) EnvelopeSettings() EdifactEnvelopeSettingsPtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) *EdifactEnvelopeSettings { return v.EnvelopeSettings }).(EdifactEnvelopeSettingsPtrOutput)
}

func (o EdifactProtocolSettingsOutput) FramingSettings() EdifactFramingSettingsPtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) *EdifactFramingSettings { return v.FramingSettings }).(EdifactFramingSettingsPtrOutput)
}

func (o EdifactProtocolSettingsOutput) MessageFilter() EdifactMessageFilterPtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) *EdifactMessageFilter { return v.MessageFilter }).(EdifactMessageFilterPtrOutput)
}

func (o EdifactProtocolSettingsOutput) MessageFilterList() EdifactMessageIdentifierArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) []EdifactMessageIdentifier { return v.MessageFilterList }).(EdifactMessageIdentifierArrayOutput)
}

func (o EdifactProtocolSettingsOutput) ProcessingSettings() EdifactProcessingSettingsPtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) *EdifactProcessingSettings { return v.ProcessingSettings }).(EdifactProcessingSettingsPtrOutput)
}

func (o EdifactProtocolSettingsOutput) SchemaReferences() EdifactSchemaReferenceArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) []EdifactSchemaReference { return v.SchemaReferences }).(EdifactSchemaReferenceArrayOutput)
}

func (o EdifactProtocolSettingsOutput) ValidationOverrides() EdifactValidationOverrideArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) []EdifactValidationOverride { return v.ValidationOverrides }).(EdifactValidationOverrideArrayOutput)
}

func (o EdifactProtocolSettingsOutput) ValidationSettings() EdifactValidationSettingsPtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) *EdifactValidationSettings { return v.ValidationSettings }).(EdifactValidationSettingsPtrOutput)
}

type EdifactProtocolSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactProtocolSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProtocolSettings)(nil)).Elem()
}

func (o EdifactProtocolSettingsPtrOutput) ToEdifactProtocolSettingsPtrOutput() EdifactProtocolSettingsPtrOutput {
	return o
}

func (o EdifactProtocolSettingsPtrOutput) ToEdifactProtocolSettingsPtrOutputWithContext(ctx context.Context) EdifactProtocolSettingsPtrOutput {
	return o
}

func (o EdifactProtocolSettingsPtrOutput) Elem() EdifactProtocolSettingsOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) EdifactProtocolSettings {
		if v != nil {
			return *v
		}
		var ret EdifactProtocolSettings
		return ret
	}).(EdifactProtocolSettingsOutput)
}

func (o EdifactProtocolSettingsPtrOutput) AcknowledgementSettings() EdifactAcknowledgementSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactAcknowledgementSettings {
		if v == nil {
			return nil
		}
		return v.AcknowledgementSettings
	}).(EdifactAcknowledgementSettingsPtrOutput)
}

func (o EdifactProtocolSettingsPtrOutput) EdifactDelimiterOverrides() EdifactDelimiterOverrideArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) []EdifactDelimiterOverride {
		if v == nil {
			return nil
		}
		return v.EdifactDelimiterOverrides
	}).(EdifactDelimiterOverrideArrayOutput)
}

func (o EdifactProtocolSettingsPtrOutput) EnvelopeOverrides() EdifactEnvelopeOverrideArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) []EdifactEnvelopeOverride {
		if v == nil {
			return nil
		}
		return v.EnvelopeOverrides
	}).(EdifactEnvelopeOverrideArrayOutput)
}

func (o EdifactProtocolSettingsPtrOutput) EnvelopeSettings() EdifactEnvelopeSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactEnvelopeSettings {
		if v == nil {
			return nil
		}
		return v.EnvelopeSettings
	}).(EdifactEnvelopeSettingsPtrOutput)
}

func (o EdifactProtocolSettingsPtrOutput) FramingSettings() EdifactFramingSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactFramingSettings {
		if v == nil {
			return nil
		}
		return v.FramingSettings
	}).(EdifactFramingSettingsPtrOutput)
}

func (o EdifactProtocolSettingsPtrOutput) MessageFilter() EdifactMessageFilterPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactMessageFilter {
		if v == nil {
			return nil
		}
		return v.MessageFilter
	}).(EdifactMessageFilterPtrOutput)
}

func (o EdifactProtocolSettingsPtrOutput) MessageFilterList() EdifactMessageIdentifierArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) []EdifactMessageIdentifier {
		if v == nil {
			return nil
		}
		return v.MessageFilterList
	}).(EdifactMessageIdentifierArrayOutput)
}

func (o EdifactProtocolSettingsPtrOutput) ProcessingSettings() EdifactProcessingSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactProcessingSettings {
		if v == nil {
			return nil
		}
		return v.ProcessingSettings
	}).(EdifactProcessingSettingsPtrOutput)
}

func (o EdifactProtocolSettingsPtrOutput) SchemaReferences() EdifactSchemaReferenceArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) []EdifactSchemaReference {
		if v == nil {
			return nil
		}
		return v.SchemaReferences
	}).(EdifactSchemaReferenceArrayOutput)
}

func (o EdifactProtocolSettingsPtrOutput) ValidationOverrides() EdifactValidationOverrideArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) []EdifactValidationOverride {
		if v == nil {
			return nil
		}
		return v.ValidationOverrides
	}).(EdifactValidationOverrideArrayOutput)
}

func (o EdifactProtocolSettingsPtrOutput) ValidationSettings() EdifactValidationSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactValidationSettings {
		if v == nil {
			return nil
		}
		return v.ValidationSettings
	}).(EdifactValidationSettingsPtrOutput)
}

type EdifactProtocolSettingsResponse struct {
	AcknowledgementSettings   *EdifactAcknowledgementSettingsResponse `pulumi:"acknowledgementSettings"`
	EdifactDelimiterOverrides []EdifactDelimiterOverrideResponse      `pulumi:"edifactDelimiterOverrides"`
	EnvelopeOverrides         []EdifactEnvelopeOverrideResponse       `pulumi:"envelopeOverrides"`
	EnvelopeSettings          *EdifactEnvelopeSettingsResponse        `pulumi:"envelopeSettings"`
	FramingSettings           *EdifactFramingSettingsResponse         `pulumi:"framingSettings"`
	MessageFilter             *EdifactMessageFilterResponse           `pulumi:"messageFilter"`
	MessageFilterList         []EdifactMessageIdentifierResponse      `pulumi:"messageFilterList"`
	ProcessingSettings        *EdifactProcessingSettingsResponse      `pulumi:"processingSettings"`
	SchemaReferences          []EdifactSchemaReferenceResponse        `pulumi:"schemaReferences"`
	ValidationOverrides       []EdifactValidationOverrideResponse     `pulumi:"validationOverrides"`
	ValidationSettings        *EdifactValidationSettingsResponse      `pulumi:"validationSettings"`
}

type EdifactProtocolSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactProtocolSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProtocolSettingsResponse)(nil)).Elem()
}

func (o EdifactProtocolSettingsResponseOutput) ToEdifactProtocolSettingsResponseOutput() EdifactProtocolSettingsResponseOutput {
	return o
}

func (o EdifactProtocolSettingsResponseOutput) ToEdifactProtocolSettingsResponseOutputWithContext(ctx context.Context) EdifactProtocolSettingsResponseOutput {
	return o
}

func (o EdifactProtocolSettingsResponseOutput) AcknowledgementSettings() EdifactAcknowledgementSettingsResponsePtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) *EdifactAcknowledgementSettingsResponse {
		return v.AcknowledgementSettings
	}).(EdifactAcknowledgementSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponseOutput) EdifactDelimiterOverrides() EdifactDelimiterOverrideResponseArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) []EdifactDelimiterOverrideResponse {
		return v.EdifactDelimiterOverrides
	}).(EdifactDelimiterOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponseOutput) EnvelopeOverrides() EdifactEnvelopeOverrideResponseArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) []EdifactEnvelopeOverrideResponse { return v.EnvelopeOverrides }).(EdifactEnvelopeOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponseOutput) EnvelopeSettings() EdifactEnvelopeSettingsResponsePtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) *EdifactEnvelopeSettingsResponse { return v.EnvelopeSettings }).(EdifactEnvelopeSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponseOutput) FramingSettings() EdifactFramingSettingsResponsePtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) *EdifactFramingSettingsResponse { return v.FramingSettings }).(EdifactFramingSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponseOutput) MessageFilter() EdifactMessageFilterResponsePtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) *EdifactMessageFilterResponse { return v.MessageFilter }).(EdifactMessageFilterResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponseOutput) MessageFilterList() EdifactMessageIdentifierResponseArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) []EdifactMessageIdentifierResponse { return v.MessageFilterList }).(EdifactMessageIdentifierResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponseOutput) ProcessingSettings() EdifactProcessingSettingsResponsePtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) *EdifactProcessingSettingsResponse {
		return v.ProcessingSettings
	}).(EdifactProcessingSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponseOutput) SchemaReferences() EdifactSchemaReferenceResponseArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) []EdifactSchemaReferenceResponse { return v.SchemaReferences }).(EdifactSchemaReferenceResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponseOutput) ValidationOverrides() EdifactValidationOverrideResponseArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) []EdifactValidationOverrideResponse {
		return v.ValidationOverrides
	}).(EdifactValidationOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponseOutput) ValidationSettings() EdifactValidationSettingsResponsePtrOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) *EdifactValidationSettingsResponse {
		return v.ValidationSettings
	}).(EdifactValidationSettingsResponsePtrOutput)
}

type EdifactProtocolSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactProtocolSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProtocolSettingsResponse)(nil)).Elem()
}

func (o EdifactProtocolSettingsResponsePtrOutput) ToEdifactProtocolSettingsResponsePtrOutput() EdifactProtocolSettingsResponsePtrOutput {
	return o
}

func (o EdifactProtocolSettingsResponsePtrOutput) ToEdifactProtocolSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactProtocolSettingsResponsePtrOutput {
	return o
}

func (o EdifactProtocolSettingsResponsePtrOutput) Elem() EdifactProtocolSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) EdifactProtocolSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactProtocolSettingsResponse
		return ret
	}).(EdifactProtocolSettingsResponseOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) AcknowledgementSettings() EdifactAcknowledgementSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactAcknowledgementSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AcknowledgementSettings
	}).(EdifactAcknowledgementSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) EdifactDelimiterOverrides() EdifactDelimiterOverrideResponseArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) []EdifactDelimiterOverrideResponse {
		if v == nil {
			return nil
		}
		return v.EdifactDelimiterOverrides
	}).(EdifactDelimiterOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) EnvelopeOverrides() EdifactEnvelopeOverrideResponseArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) []EdifactEnvelopeOverrideResponse {
		if v == nil {
			return nil
		}
		return v.EnvelopeOverrides
	}).(EdifactEnvelopeOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) EnvelopeSettings() EdifactEnvelopeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactEnvelopeSettingsResponse {
		if v == nil {
			return nil
		}
		return v.EnvelopeSettings
	}).(EdifactEnvelopeSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) FramingSettings() EdifactFramingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactFramingSettingsResponse {
		if v == nil {
			return nil
		}
		return v.FramingSettings
	}).(EdifactFramingSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) MessageFilter() EdifactMessageFilterResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactMessageFilterResponse {
		if v == nil {
			return nil
		}
		return v.MessageFilter
	}).(EdifactMessageFilterResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) MessageFilterList() EdifactMessageIdentifierResponseArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) []EdifactMessageIdentifierResponse {
		if v == nil {
			return nil
		}
		return v.MessageFilterList
	}).(EdifactMessageIdentifierResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) ProcessingSettings() EdifactProcessingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactProcessingSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ProcessingSettings
	}).(EdifactProcessingSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) SchemaReferences() EdifactSchemaReferenceResponseArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) []EdifactSchemaReferenceResponse {
		if v == nil {
			return nil
		}
		return v.SchemaReferences
	}).(EdifactSchemaReferenceResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) ValidationOverrides() EdifactValidationOverrideResponseArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) []EdifactValidationOverrideResponse {
		if v == nil {
			return nil
		}
		return v.ValidationOverrides
	}).(EdifactValidationOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) ValidationSettings() EdifactValidationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactValidationSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ValidationSettings
	}).(EdifactValidationSettingsResponsePtrOutput)
}

type EdifactSchemaReference struct {
	AssociationAssignedCode    *string `pulumi:"associationAssignedCode"`
	MessageId                  *string `pulumi:"messageId"`
	MessageRelease             *string `pulumi:"messageRelease"`
	MessageVersion             *string `pulumi:"messageVersion"`
	SchemaName                 *string `pulumi:"schemaName"`
	SenderApplicationId        *string `pulumi:"senderApplicationId"`
	SenderApplicationQualifier *string `pulumi:"senderApplicationQualifier"`
}





type EdifactSchemaReferenceInput interface {
	pulumi.Input

	ToEdifactSchemaReferenceOutput() EdifactSchemaReferenceOutput
	ToEdifactSchemaReferenceOutputWithContext(context.Context) EdifactSchemaReferenceOutput
}

type EdifactSchemaReferenceArgs struct {
	AssociationAssignedCode    pulumi.StringPtrInput `pulumi:"associationAssignedCode"`
	MessageId                  pulumi.StringPtrInput `pulumi:"messageId"`
	MessageRelease             pulumi.StringPtrInput `pulumi:"messageRelease"`
	MessageVersion             pulumi.StringPtrInput `pulumi:"messageVersion"`
	SchemaName                 pulumi.StringPtrInput `pulumi:"schemaName"`
	SenderApplicationId        pulumi.StringPtrInput `pulumi:"senderApplicationId"`
	SenderApplicationQualifier pulumi.StringPtrInput `pulumi:"senderApplicationQualifier"`
}

func (EdifactSchemaReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactSchemaReference)(nil)).Elem()
}

func (i EdifactSchemaReferenceArgs) ToEdifactSchemaReferenceOutput() EdifactSchemaReferenceOutput {
	return i.ToEdifactSchemaReferenceOutputWithContext(context.Background())
}

func (i EdifactSchemaReferenceArgs) ToEdifactSchemaReferenceOutputWithContext(ctx context.Context) EdifactSchemaReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactSchemaReferenceOutput)
}





type EdifactSchemaReferenceArrayInput interface {
	pulumi.Input

	ToEdifactSchemaReferenceArrayOutput() EdifactSchemaReferenceArrayOutput
	ToEdifactSchemaReferenceArrayOutputWithContext(context.Context) EdifactSchemaReferenceArrayOutput
}

type EdifactSchemaReferenceArray []EdifactSchemaReferenceInput

func (EdifactSchemaReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactSchemaReference)(nil)).Elem()
}

func (i EdifactSchemaReferenceArray) ToEdifactSchemaReferenceArrayOutput() EdifactSchemaReferenceArrayOutput {
	return i.ToEdifactSchemaReferenceArrayOutputWithContext(context.Background())
}

func (i EdifactSchemaReferenceArray) ToEdifactSchemaReferenceArrayOutputWithContext(ctx context.Context) EdifactSchemaReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactSchemaReferenceArrayOutput)
}

type EdifactSchemaReferenceOutput struct{ *pulumi.OutputState }

func (EdifactSchemaReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactSchemaReference)(nil)).Elem()
}

func (o EdifactSchemaReferenceOutput) ToEdifactSchemaReferenceOutput() EdifactSchemaReferenceOutput {
	return o
}

func (o EdifactSchemaReferenceOutput) ToEdifactSchemaReferenceOutputWithContext(ctx context.Context) EdifactSchemaReferenceOutput {
	return o
}

func (o EdifactSchemaReferenceOutput) AssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReference) *string { return v.AssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReference) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceOutput) MessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReference) *string { return v.MessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceOutput) MessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReference) *string { return v.MessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReference) *string { return v.SchemaName }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReference) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceOutput) SenderApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReference) *string { return v.SenderApplicationQualifier }).(pulumi.StringPtrOutput)
}

type EdifactSchemaReferenceArrayOutput struct{ *pulumi.OutputState }

func (EdifactSchemaReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactSchemaReference)(nil)).Elem()
}

func (o EdifactSchemaReferenceArrayOutput) ToEdifactSchemaReferenceArrayOutput() EdifactSchemaReferenceArrayOutput {
	return o
}

func (o EdifactSchemaReferenceArrayOutput) ToEdifactSchemaReferenceArrayOutputWithContext(ctx context.Context) EdifactSchemaReferenceArrayOutput {
	return o
}

func (o EdifactSchemaReferenceArrayOutput) Index(i pulumi.IntInput) EdifactSchemaReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactSchemaReference {
		return vs[0].([]EdifactSchemaReference)[vs[1].(int)]
	}).(EdifactSchemaReferenceOutput)
}

type EdifactSchemaReferenceResponse struct {
	AssociationAssignedCode    *string `pulumi:"associationAssignedCode"`
	MessageId                  *string `pulumi:"messageId"`
	MessageRelease             *string `pulumi:"messageRelease"`
	MessageVersion             *string `pulumi:"messageVersion"`
	SchemaName                 *string `pulumi:"schemaName"`
	SenderApplicationId        *string `pulumi:"senderApplicationId"`
	SenderApplicationQualifier *string `pulumi:"senderApplicationQualifier"`
}

type EdifactSchemaReferenceResponseOutput struct{ *pulumi.OutputState }

func (EdifactSchemaReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactSchemaReferenceResponse)(nil)).Elem()
}

func (o EdifactSchemaReferenceResponseOutput) ToEdifactSchemaReferenceResponseOutput() EdifactSchemaReferenceResponseOutput {
	return o
}

func (o EdifactSchemaReferenceResponseOutput) ToEdifactSchemaReferenceResponseOutputWithContext(ctx context.Context) EdifactSchemaReferenceResponseOutput {
	return o
}

func (o EdifactSchemaReferenceResponseOutput) AssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) *string { return v.AssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceResponseOutput) MessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) *string { return v.MessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceResponseOutput) MessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) *string { return v.MessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceResponseOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) *string { return v.SchemaName }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceResponseOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceResponseOutput) SenderApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) *string { return v.SenderApplicationQualifier }).(pulumi.StringPtrOutput)
}

type EdifactSchemaReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (EdifactSchemaReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactSchemaReferenceResponse)(nil)).Elem()
}

func (o EdifactSchemaReferenceResponseArrayOutput) ToEdifactSchemaReferenceResponseArrayOutput() EdifactSchemaReferenceResponseArrayOutput {
	return o
}

func (o EdifactSchemaReferenceResponseArrayOutput) ToEdifactSchemaReferenceResponseArrayOutputWithContext(ctx context.Context) EdifactSchemaReferenceResponseArrayOutput {
	return o
}

func (o EdifactSchemaReferenceResponseArrayOutput) Index(i pulumi.IntInput) EdifactSchemaReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactSchemaReferenceResponse {
		return vs[0].([]EdifactSchemaReferenceResponse)[vs[1].(int)]
	}).(EdifactSchemaReferenceResponseOutput)
}

type EdifactValidationOverride struct {
	AllowLeadingAndTrailingSpacesAndZeroes *bool                    `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	EnforceCharacterSet                    *bool                    `pulumi:"enforceCharacterSet"`
	MessageId                              *string                  `pulumi:"messageId"`
	TrailingSeparatorPolicy                *TrailingSeparatorPolicy `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  *bool                    `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateEDITypes                       *bool                    `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       *bool                    `pulumi:"validateXSDTypes"`
}





type EdifactValidationOverrideInput interface {
	pulumi.Input

	ToEdifactValidationOverrideOutput() EdifactValidationOverrideOutput
	ToEdifactValidationOverrideOutputWithContext(context.Context) EdifactValidationOverrideOutput
}

type EdifactValidationOverrideArgs struct {
	AllowLeadingAndTrailingSpacesAndZeroes pulumi.BoolPtrInput             `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	EnforceCharacterSet                    pulumi.BoolPtrInput             `pulumi:"enforceCharacterSet"`
	MessageId                              pulumi.StringPtrInput           `pulumi:"messageId"`
	TrailingSeparatorPolicy                TrailingSeparatorPolicyPtrInput `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  pulumi.BoolPtrInput             `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateEDITypes                       pulumi.BoolPtrInput             `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       pulumi.BoolPtrInput             `pulumi:"validateXSDTypes"`
}

func (EdifactValidationOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationOverride)(nil)).Elem()
}

func (i EdifactValidationOverrideArgs) ToEdifactValidationOverrideOutput() EdifactValidationOverrideOutput {
	return i.ToEdifactValidationOverrideOutputWithContext(context.Background())
}

func (i EdifactValidationOverrideArgs) ToEdifactValidationOverrideOutputWithContext(ctx context.Context) EdifactValidationOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactValidationOverrideOutput)
}





type EdifactValidationOverrideArrayInput interface {
	pulumi.Input

	ToEdifactValidationOverrideArrayOutput() EdifactValidationOverrideArrayOutput
	ToEdifactValidationOverrideArrayOutputWithContext(context.Context) EdifactValidationOverrideArrayOutput
}

type EdifactValidationOverrideArray []EdifactValidationOverrideInput

func (EdifactValidationOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactValidationOverride)(nil)).Elem()
}

func (i EdifactValidationOverrideArray) ToEdifactValidationOverrideArrayOutput() EdifactValidationOverrideArrayOutput {
	return i.ToEdifactValidationOverrideArrayOutputWithContext(context.Background())
}

func (i EdifactValidationOverrideArray) ToEdifactValidationOverrideArrayOutputWithContext(ctx context.Context) EdifactValidationOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactValidationOverrideArrayOutput)
}

type EdifactValidationOverrideOutput struct{ *pulumi.OutputState }

func (EdifactValidationOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationOverride)(nil)).Elem()
}

func (o EdifactValidationOverrideOutput) ToEdifactValidationOverrideOutput() EdifactValidationOverrideOutput {
	return o
}

func (o EdifactValidationOverrideOutput) ToEdifactValidationOverrideOutputWithContext(ctx context.Context) EdifactValidationOverrideOutput {
	return o
}

func (o EdifactValidationOverrideOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverride) *bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationOverrideOutput) EnforceCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverride) *bool { return v.EnforceCharacterSet }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationOverrideOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverride) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactValidationOverrideOutput) TrailingSeparatorPolicy() TrailingSeparatorPolicyPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverride) *TrailingSeparatorPolicy { return v.TrailingSeparatorPolicy }).(TrailingSeparatorPolicyPtrOutput)
}

func (o EdifactValidationOverrideOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverride) *bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationOverrideOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverride) *bool { return v.ValidateEDITypes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationOverrideOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverride) *bool { return v.ValidateXSDTypes }).(pulumi.BoolPtrOutput)
}

type EdifactValidationOverrideArrayOutput struct{ *pulumi.OutputState }

func (EdifactValidationOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactValidationOverride)(nil)).Elem()
}

func (o EdifactValidationOverrideArrayOutput) ToEdifactValidationOverrideArrayOutput() EdifactValidationOverrideArrayOutput {
	return o
}

func (o EdifactValidationOverrideArrayOutput) ToEdifactValidationOverrideArrayOutputWithContext(ctx context.Context) EdifactValidationOverrideArrayOutput {
	return o
}

func (o EdifactValidationOverrideArrayOutput) Index(i pulumi.IntInput) EdifactValidationOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactValidationOverride {
		return vs[0].([]EdifactValidationOverride)[vs[1].(int)]
	}).(EdifactValidationOverrideOutput)
}

type EdifactValidationOverrideResponse struct {
	AllowLeadingAndTrailingSpacesAndZeroes *bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	EnforceCharacterSet                    *bool   `pulumi:"enforceCharacterSet"`
	MessageId                              *string `pulumi:"messageId"`
	TrailingSeparatorPolicy                *string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  *bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateEDITypes                       *bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       *bool   `pulumi:"validateXSDTypes"`
}

type EdifactValidationOverrideResponseOutput struct{ *pulumi.OutputState }

func (EdifactValidationOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationOverrideResponse)(nil)).Elem()
}

func (o EdifactValidationOverrideResponseOutput) ToEdifactValidationOverrideResponseOutput() EdifactValidationOverrideResponseOutput {
	return o
}

func (o EdifactValidationOverrideResponseOutput) ToEdifactValidationOverrideResponseOutputWithContext(ctx context.Context) EdifactValidationOverrideResponseOutput {
	return o
}

func (o EdifactValidationOverrideResponseOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) *bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationOverrideResponseOutput) EnforceCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) *bool { return v.EnforceCharacterSet }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationOverrideResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactValidationOverrideResponseOutput) TrailingSeparatorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) *string { return v.TrailingSeparatorPolicy }).(pulumi.StringPtrOutput)
}

func (o EdifactValidationOverrideResponseOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) *bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationOverrideResponseOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) *bool { return v.ValidateEDITypes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationOverrideResponseOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) *bool { return v.ValidateXSDTypes }).(pulumi.BoolPtrOutput)
}

type EdifactValidationOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (EdifactValidationOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactValidationOverrideResponse)(nil)).Elem()
}

func (o EdifactValidationOverrideResponseArrayOutput) ToEdifactValidationOverrideResponseArrayOutput() EdifactValidationOverrideResponseArrayOutput {
	return o
}

func (o EdifactValidationOverrideResponseArrayOutput) ToEdifactValidationOverrideResponseArrayOutputWithContext(ctx context.Context) EdifactValidationOverrideResponseArrayOutput {
	return o
}

func (o EdifactValidationOverrideResponseArrayOutput) Index(i pulumi.IntInput) EdifactValidationOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactValidationOverrideResponse {
		return vs[0].([]EdifactValidationOverrideResponse)[vs[1].(int)]
	}).(EdifactValidationOverrideResponseOutput)
}

type EdifactValidationSettings struct {
	AllowLeadingAndTrailingSpacesAndZeroes    *bool                    `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          *bool                    `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    *bool                    `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber *bool                    `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      *int                     `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   *TrailingSeparatorPolicy `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     *bool                    `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      *bool                    `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          *bool                    `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          *bool                    `pulumi:"validateXSDTypes"`
}





type EdifactValidationSettingsInput interface {
	pulumi.Input

	ToEdifactValidationSettingsOutput() EdifactValidationSettingsOutput
	ToEdifactValidationSettingsOutputWithContext(context.Context) EdifactValidationSettingsOutput
}

type EdifactValidationSettingsArgs struct {
	AllowLeadingAndTrailingSpacesAndZeroes    pulumi.BoolPtrInput             `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          pulumi.BoolPtrInput             `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    pulumi.BoolPtrInput             `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber pulumi.BoolPtrInput             `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      pulumi.IntPtrInput              `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   TrailingSeparatorPolicyPtrInput `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     pulumi.BoolPtrInput             `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      pulumi.BoolPtrInput             `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          pulumi.BoolPtrInput             `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          pulumi.BoolPtrInput             `pulumi:"validateXSDTypes"`
}

func (EdifactValidationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationSettings)(nil)).Elem()
}

func (i EdifactValidationSettingsArgs) ToEdifactValidationSettingsOutput() EdifactValidationSettingsOutput {
	return i.ToEdifactValidationSettingsOutputWithContext(context.Background())
}

func (i EdifactValidationSettingsArgs) ToEdifactValidationSettingsOutputWithContext(ctx context.Context) EdifactValidationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactValidationSettingsOutput)
}

func (i EdifactValidationSettingsArgs) ToEdifactValidationSettingsPtrOutput() EdifactValidationSettingsPtrOutput {
	return i.ToEdifactValidationSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactValidationSettingsArgs) ToEdifactValidationSettingsPtrOutputWithContext(ctx context.Context) EdifactValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactValidationSettingsOutput).ToEdifactValidationSettingsPtrOutputWithContext(ctx)
}









type EdifactValidationSettingsPtrInput interface {
	pulumi.Input

	ToEdifactValidationSettingsPtrOutput() EdifactValidationSettingsPtrOutput
	ToEdifactValidationSettingsPtrOutputWithContext(context.Context) EdifactValidationSettingsPtrOutput
}

type edifactValidationSettingsPtrType EdifactValidationSettingsArgs

func EdifactValidationSettingsPtr(v *EdifactValidationSettingsArgs) EdifactValidationSettingsPtrInput {
	return (*edifactValidationSettingsPtrType)(v)
}

func (*edifactValidationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactValidationSettings)(nil)).Elem()
}

func (i *edifactValidationSettingsPtrType) ToEdifactValidationSettingsPtrOutput() EdifactValidationSettingsPtrOutput {
	return i.ToEdifactValidationSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactValidationSettingsPtrType) ToEdifactValidationSettingsPtrOutputWithContext(ctx context.Context) EdifactValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactValidationSettingsPtrOutput)
}

type EdifactValidationSettingsOutput struct{ *pulumi.OutputState }

func (EdifactValidationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationSettings)(nil)).Elem()
}

func (o EdifactValidationSettingsOutput) ToEdifactValidationSettingsOutput() EdifactValidationSettingsOutput {
	return o
}

func (o EdifactValidationSettingsOutput) ToEdifactValidationSettingsOutputWithContext(ctx context.Context) EdifactValidationSettingsOutput {
	return o
}

func (o EdifactValidationSettingsOutput) ToEdifactValidationSettingsPtrOutput() EdifactValidationSettingsPtrOutput {
	return o.ToEdifactValidationSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactValidationSettingsOutput) ToEdifactValidationSettingsPtrOutputWithContext(ctx context.Context) EdifactValidationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactValidationSettings) *EdifactValidationSettings {
		return &v
	}).(EdifactValidationSettingsPtrOutput)
}

func (o EdifactValidationSettingsOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettings) *bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettings) *bool { return v.CheckDuplicateGroupControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettings) *bool { return v.CheckDuplicateInterchangeControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettings) *bool { return v.CheckDuplicateTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettings) *int { return v.InterchangeControlNumberValidityDays }).(pulumi.IntPtrOutput)
}

func (o EdifactValidationSettingsOutput) TrailingSeparatorPolicy() TrailingSeparatorPolicyPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettings) *TrailingSeparatorPolicy { return v.TrailingSeparatorPolicy }).(TrailingSeparatorPolicyPtrOutput)
}

func (o EdifactValidationSettingsOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettings) *bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettings) *bool { return v.ValidateCharacterSet }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettings) *bool { return v.ValidateEDITypes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettings) *bool { return v.ValidateXSDTypes }).(pulumi.BoolPtrOutput)
}

type EdifactValidationSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactValidationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactValidationSettings)(nil)).Elem()
}

func (o EdifactValidationSettingsPtrOutput) ToEdifactValidationSettingsPtrOutput() EdifactValidationSettingsPtrOutput {
	return o
}

func (o EdifactValidationSettingsPtrOutput) ToEdifactValidationSettingsPtrOutputWithContext(ctx context.Context) EdifactValidationSettingsPtrOutput {
	return o
}

func (o EdifactValidationSettingsPtrOutput) Elem() EdifactValidationSettingsOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) EdifactValidationSettings {
		if v != nil {
			return *v
		}
		var ret EdifactValidationSettings
		return ret
	}).(EdifactValidationSettingsOutput)
}

func (o EdifactValidationSettingsPtrOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.AllowLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *int {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) TrailingSeparatorPolicy() TrailingSeparatorPolicyPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *TrailingSeparatorPolicy {
		if v == nil {
			return nil
		}
		return v.TrailingSeparatorPolicy
	}).(TrailingSeparatorPolicyPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.TrimLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateCharacterSet
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateEDITypes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateXSDTypes
	}).(pulumi.BoolPtrOutput)
}

type EdifactValidationSettingsResponse struct {
	AllowLeadingAndTrailingSpacesAndZeroes    *bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          *bool   `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    *bool   `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber *bool   `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      *int    `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   *string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     *bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      *bool   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          *bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          *bool   `pulumi:"validateXSDTypes"`
}

type EdifactValidationSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactValidationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationSettingsResponse)(nil)).Elem()
}

func (o EdifactValidationSettingsResponseOutput) ToEdifactValidationSettingsResponseOutput() EdifactValidationSettingsResponseOutput {
	return o
}

func (o EdifactValidationSettingsResponseOutput) ToEdifactValidationSettingsResponseOutputWithContext(ctx context.Context) EdifactValidationSettingsResponseOutput {
	return o
}

func (o EdifactValidationSettingsResponseOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) *bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponseOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) *bool { return v.CheckDuplicateGroupControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponseOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) *bool { return v.CheckDuplicateInterchangeControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponseOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) *bool { return v.CheckDuplicateTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponseOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) *int { return v.InterchangeControlNumberValidityDays }).(pulumi.IntPtrOutput)
}

func (o EdifactValidationSettingsResponseOutput) TrailingSeparatorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) *string { return v.TrailingSeparatorPolicy }).(pulumi.StringPtrOutput)
}

func (o EdifactValidationSettingsResponseOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) *bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponseOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) *bool { return v.ValidateCharacterSet }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponseOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) *bool { return v.ValidateEDITypes }).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponseOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) *bool { return v.ValidateXSDTypes }).(pulumi.BoolPtrOutput)
}

type EdifactValidationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactValidationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactValidationSettingsResponse)(nil)).Elem()
}

func (o EdifactValidationSettingsResponsePtrOutput) ToEdifactValidationSettingsResponsePtrOutput() EdifactValidationSettingsResponsePtrOutput {
	return o
}

func (o EdifactValidationSettingsResponsePtrOutput) ToEdifactValidationSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactValidationSettingsResponsePtrOutput {
	return o
}

func (o EdifactValidationSettingsResponsePtrOutput) Elem() EdifactValidationSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) EdifactValidationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactValidationSettingsResponse
		return ret
	}).(EdifactValidationSettingsResponseOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) TrailingSeparatorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TrailingSeparatorPolicy
	}).(pulumi.StringPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.TrimLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateCharacterSet
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateEDITypes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateXSDTypes
	}).(pulumi.BoolPtrOutput)
}

type IntegrationAccountContentHashResponse struct {
	Algorithm *string `pulumi:"algorithm"`
	Value     *string `pulumi:"value"`
}

type IntegrationAccountContentHashResponseOutput struct{ *pulumi.OutputState }

func (IntegrationAccountContentHashResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountContentHashResponse)(nil)).Elem()
}

func (o IntegrationAccountContentHashResponseOutput) ToIntegrationAccountContentHashResponseOutput() IntegrationAccountContentHashResponseOutput {
	return o
}

func (o IntegrationAccountContentHashResponseOutput) ToIntegrationAccountContentHashResponseOutputWithContext(ctx context.Context) IntegrationAccountContentHashResponseOutput {
	return o
}

func (o IntegrationAccountContentHashResponseOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationAccountContentHashResponse) *string { return v.Algorithm }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountContentHashResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationAccountContentHashResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type IntegrationAccountContentHashResponsePtrOutput struct{ *pulumi.OutputState }

func (IntegrationAccountContentHashResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountContentHashResponse)(nil)).Elem()
}

func (o IntegrationAccountContentHashResponsePtrOutput) ToIntegrationAccountContentHashResponsePtrOutput() IntegrationAccountContentHashResponsePtrOutput {
	return o
}

func (o IntegrationAccountContentHashResponsePtrOutput) ToIntegrationAccountContentHashResponsePtrOutputWithContext(ctx context.Context) IntegrationAccountContentHashResponsePtrOutput {
	return o
}

func (o IntegrationAccountContentHashResponsePtrOutput) Elem() IntegrationAccountContentHashResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountContentHashResponse) IntegrationAccountContentHashResponse {
		if v != nil {
			return *v
		}
		var ret IntegrationAccountContentHashResponse
		return ret
	}).(IntegrationAccountContentHashResponseOutput)
}

func (o IntegrationAccountContentHashResponsePtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountContentHashResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type IntegrationAccountContentLinkResponse struct {
	ContentHash    *IntegrationAccountContentHashResponse `pulumi:"contentHash"`
	ContentSize    *float64                               `pulumi:"contentSize"`
	ContentVersion *string                                `pulumi:"contentVersion"`
	Metadata       interface{}                            `pulumi:"metadata"`
	Uri            *string                                `pulumi:"uri"`
}

type IntegrationAccountContentLinkResponseOutput struct{ *pulumi.OutputState }

func (IntegrationAccountContentLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountContentLinkResponse)(nil)).Elem()
}

func (o IntegrationAccountContentLinkResponseOutput) ToIntegrationAccountContentLinkResponseOutput() IntegrationAccountContentLinkResponseOutput {
	return o
}

func (o IntegrationAccountContentLinkResponseOutput) ToIntegrationAccountContentLinkResponseOutputWithContext(ctx context.Context) IntegrationAccountContentLinkResponseOutput {
	return o
}

func (o IntegrationAccountContentLinkResponseOutput) ContentHash() IntegrationAccountContentHashResponsePtrOutput {
	return o.ApplyT(func(v IntegrationAccountContentLinkResponse) *IntegrationAccountContentHashResponse {
		return v.ContentHash
	}).(IntegrationAccountContentHashResponsePtrOutput)
}

func (o IntegrationAccountContentLinkResponseOutput) ContentSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IntegrationAccountContentLinkResponse) *float64 { return v.ContentSize }).(pulumi.Float64PtrOutput)
}

func (o IntegrationAccountContentLinkResponseOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationAccountContentLinkResponse) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountContentLinkResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v IntegrationAccountContentLinkResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o IntegrationAccountContentLinkResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationAccountContentLinkResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type IntegrationAccountSku struct {
	Name *SkuName `pulumi:"name"`
}





type IntegrationAccountSkuInput interface {
	pulumi.Input

	ToIntegrationAccountSkuOutput() IntegrationAccountSkuOutput
	ToIntegrationAccountSkuOutputWithContext(context.Context) IntegrationAccountSkuOutput
}

type IntegrationAccountSkuArgs struct {
	Name SkuNamePtrInput `pulumi:"name"`
}

func (IntegrationAccountSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountSku)(nil)).Elem()
}

func (i IntegrationAccountSkuArgs) ToIntegrationAccountSkuOutput() IntegrationAccountSkuOutput {
	return i.ToIntegrationAccountSkuOutputWithContext(context.Background())
}

func (i IntegrationAccountSkuArgs) ToIntegrationAccountSkuOutputWithContext(ctx context.Context) IntegrationAccountSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountSkuOutput)
}

func (i IntegrationAccountSkuArgs) ToIntegrationAccountSkuPtrOutput() IntegrationAccountSkuPtrOutput {
	return i.ToIntegrationAccountSkuPtrOutputWithContext(context.Background())
}

func (i IntegrationAccountSkuArgs) ToIntegrationAccountSkuPtrOutputWithContext(ctx context.Context) IntegrationAccountSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountSkuOutput).ToIntegrationAccountSkuPtrOutputWithContext(ctx)
}









type IntegrationAccountSkuPtrInput interface {
	pulumi.Input

	ToIntegrationAccountSkuPtrOutput() IntegrationAccountSkuPtrOutput
	ToIntegrationAccountSkuPtrOutputWithContext(context.Context) IntegrationAccountSkuPtrOutput
}

type integrationAccountSkuPtrType IntegrationAccountSkuArgs

func IntegrationAccountSkuPtr(v *IntegrationAccountSkuArgs) IntegrationAccountSkuPtrInput {
	return (*integrationAccountSkuPtrType)(v)
}

func (*integrationAccountSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSku)(nil)).Elem()
}

func (i *integrationAccountSkuPtrType) ToIntegrationAccountSkuPtrOutput() IntegrationAccountSkuPtrOutput {
	return i.ToIntegrationAccountSkuPtrOutputWithContext(context.Background())
}

func (i *integrationAccountSkuPtrType) ToIntegrationAccountSkuPtrOutputWithContext(ctx context.Context) IntegrationAccountSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountSkuPtrOutput)
}

type IntegrationAccountSkuOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountSku)(nil)).Elem()
}

func (o IntegrationAccountSkuOutput) ToIntegrationAccountSkuOutput() IntegrationAccountSkuOutput {
	return o
}

func (o IntegrationAccountSkuOutput) ToIntegrationAccountSkuOutputWithContext(ctx context.Context) IntegrationAccountSkuOutput {
	return o
}

func (o IntegrationAccountSkuOutput) ToIntegrationAccountSkuPtrOutput() IntegrationAccountSkuPtrOutput {
	return o.ToIntegrationAccountSkuPtrOutputWithContext(context.Background())
}

func (o IntegrationAccountSkuOutput) ToIntegrationAccountSkuPtrOutputWithContext(ctx context.Context) IntegrationAccountSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationAccountSku) *IntegrationAccountSku {
		return &v
	}).(IntegrationAccountSkuPtrOutput)
}

func (o IntegrationAccountSkuOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v IntegrationAccountSku) *SkuName { return v.Name }).(SkuNamePtrOutput)
}

type IntegrationAccountSkuPtrOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSku)(nil)).Elem()
}

func (o IntegrationAccountSkuPtrOutput) ToIntegrationAccountSkuPtrOutput() IntegrationAccountSkuPtrOutput {
	return o
}

func (o IntegrationAccountSkuPtrOutput) ToIntegrationAccountSkuPtrOutputWithContext(ctx context.Context) IntegrationAccountSkuPtrOutput {
	return o
}

func (o IntegrationAccountSkuPtrOutput) Elem() IntegrationAccountSkuOutput {
	return o.ApplyT(func(v *IntegrationAccountSku) IntegrationAccountSku {
		if v != nil {
			return *v
		}
		var ret IntegrationAccountSku
		return ret
	}).(IntegrationAccountSkuOutput)
}

func (o IntegrationAccountSkuPtrOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSku) *SkuName {
		if v == nil {
			return nil
		}
		return v.Name
	}).(SkuNamePtrOutput)
}

type IntegrationAccountSkuResponse struct {
	Name *string `pulumi:"name"`
}

type IntegrationAccountSkuResponseOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountSkuResponse)(nil)).Elem()
}

func (o IntegrationAccountSkuResponseOutput) ToIntegrationAccountSkuResponseOutput() IntegrationAccountSkuResponseOutput {
	return o
}

func (o IntegrationAccountSkuResponseOutput) ToIntegrationAccountSkuResponseOutputWithContext(ctx context.Context) IntegrationAccountSkuResponseOutput {
	return o
}

func (o IntegrationAccountSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationAccountSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type IntegrationAccountSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSkuResponse)(nil)).Elem()
}

func (o IntegrationAccountSkuResponsePtrOutput) ToIntegrationAccountSkuResponsePtrOutput() IntegrationAccountSkuResponsePtrOutput {
	return o
}

func (o IntegrationAccountSkuResponsePtrOutput) ToIntegrationAccountSkuResponsePtrOutputWithContext(ctx context.Context) IntegrationAccountSkuResponsePtrOutput {
	return o
}

func (o IntegrationAccountSkuResponsePtrOutput) Elem() IntegrationAccountSkuResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountSkuResponse) IntegrationAccountSkuResponse {
		if v != nil {
			return *v
		}
		var ret IntegrationAccountSkuResponse
		return ret
	}).(IntegrationAccountSkuResponseOutput)
}

func (o IntegrationAccountSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReference struct {
	KeyName    *string                       `pulumi:"keyName"`
	KeyVault   *KeyVaultKeyReferenceKeyVault `pulumi:"keyVault"`
	KeyVersion *string                       `pulumi:"keyVersion"`
}





type KeyVaultKeyReferenceInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput
	ToKeyVaultKeyReferenceOutputWithContext(context.Context) KeyVaultKeyReferenceOutput
}

type KeyVaultKeyReferenceArgs struct {
	KeyName    pulumi.StringPtrInput                `pulumi:"keyName"`
	KeyVault   KeyVaultKeyReferenceKeyVaultPtrInput `pulumi:"keyVault"`
	KeyVersion pulumi.StringPtrInput                `pulumi:"keyVersion"`
}

func (KeyVaultKeyReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReference)(nil)).Elem()
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput {
	return i.ToKeyVaultKeyReferenceOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferenceOutputWithContext(ctx context.Context) KeyVaultKeyReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceOutput)
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return i.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceOutput).ToKeyVaultKeyReferencePtrOutputWithContext(ctx)
}









type KeyVaultKeyReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput
	ToKeyVaultKeyReferencePtrOutputWithContext(context.Context) KeyVaultKeyReferencePtrOutput
}

type keyVaultKeyReferencePtrType KeyVaultKeyReferenceArgs

func KeyVaultKeyReferencePtr(v *KeyVaultKeyReferenceArgs) KeyVaultKeyReferencePtrInput {
	return (*keyVaultKeyReferencePtrType)(v)
}

func (*keyVaultKeyReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReference)(nil)).Elem()
}

func (i *keyVaultKeyReferencePtrType) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return i.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyReferencePtrType) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferencePtrOutput)
}

type KeyVaultKeyReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReference)(nil)).Elem()
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput {
	return o
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferenceOutputWithContext(ctx context.Context) KeyVaultKeyReferenceOutput {
	return o
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return o.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKeyReference) *KeyVaultKeyReference {
		return &v
	}).(KeyVaultKeyReferencePtrOutput)
}

func (o KeyVaultKeyReferenceOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceOutput) KeyVault() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) *KeyVaultKeyReferenceKeyVault { return v.KeyVault }).(KeyVaultKeyReferenceKeyVaultPtrOutput)
}

func (o KeyVaultKeyReferenceOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReference)(nil)).Elem()
}

func (o KeyVaultKeyReferencePtrOutput) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return o
}

func (o KeyVaultKeyReferencePtrOutput) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return o
}

func (o KeyVaultKeyReferencePtrOutput) Elem() KeyVaultKeyReferenceOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) KeyVaultKeyReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReference
		return ret
	}).(KeyVaultKeyReferenceOutput)
}

func (o KeyVaultKeyReferencePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferencePtrOutput) KeyVault() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *KeyVaultKeyReferenceKeyVault {
		if v == nil {
			return nil
		}
		return v.KeyVault
	}).(KeyVaultKeyReferenceKeyVaultPtrOutput)
}

func (o KeyVaultKeyReferencePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceKeyVault struct {
	Id *string `pulumi:"id"`
}





type KeyVaultKeyReferenceKeyVaultInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceKeyVaultOutput() KeyVaultKeyReferenceKeyVaultOutput
	ToKeyVaultKeyReferenceKeyVaultOutputWithContext(context.Context) KeyVaultKeyReferenceKeyVaultOutput
}

type KeyVaultKeyReferenceKeyVaultArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (KeyVaultKeyReferenceKeyVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceKeyVault)(nil)).Elem()
}

func (i KeyVaultKeyReferenceKeyVaultArgs) ToKeyVaultKeyReferenceKeyVaultOutput() KeyVaultKeyReferenceKeyVaultOutput {
	return i.ToKeyVaultKeyReferenceKeyVaultOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceKeyVaultArgs) ToKeyVaultKeyReferenceKeyVaultOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceKeyVaultOutput)
}

func (i KeyVaultKeyReferenceKeyVaultArgs) ToKeyVaultKeyReferenceKeyVaultPtrOutput() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return i.ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceKeyVaultArgs) ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceKeyVaultOutput).ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(ctx)
}









type KeyVaultKeyReferenceKeyVaultPtrInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceKeyVaultPtrOutput() KeyVaultKeyReferenceKeyVaultPtrOutput
	ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(context.Context) KeyVaultKeyReferenceKeyVaultPtrOutput
}

type keyVaultKeyReferenceKeyVaultPtrType KeyVaultKeyReferenceKeyVaultArgs

func KeyVaultKeyReferenceKeyVaultPtr(v *KeyVaultKeyReferenceKeyVaultArgs) KeyVaultKeyReferenceKeyVaultPtrInput {
	return (*keyVaultKeyReferenceKeyVaultPtrType)(v)
}

func (*keyVaultKeyReferenceKeyVaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceKeyVault)(nil)).Elem()
}

func (i *keyVaultKeyReferenceKeyVaultPtrType) ToKeyVaultKeyReferenceKeyVaultPtrOutput() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return i.ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyReferenceKeyVaultPtrType) ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceKeyVaultPtrOutput)
}

type KeyVaultKeyReferenceKeyVaultOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceKeyVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceKeyVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceKeyVaultOutput) ToKeyVaultKeyReferenceKeyVaultOutput() KeyVaultKeyReferenceKeyVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceKeyVaultOutput) ToKeyVaultKeyReferenceKeyVaultOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceKeyVaultOutput) ToKeyVaultKeyReferenceKeyVaultPtrOutput() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o.ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyReferenceKeyVaultOutput) ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKeyReferenceKeyVault) *KeyVaultKeyReferenceKeyVault {
		return &v
	}).(KeyVaultKeyReferenceKeyVaultPtrOutput)
}

func (o KeyVaultKeyReferenceKeyVaultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceKeyVault) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceKeyVaultPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceKeyVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceKeyVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceKeyVaultPtrOutput) ToKeyVaultKeyReferenceKeyVaultPtrOutput() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceKeyVaultPtrOutput) ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceKeyVaultPtrOutput) Elem() KeyVaultKeyReferenceKeyVaultOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceKeyVault) KeyVaultKeyReferenceKeyVault {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceKeyVault
		return ret
	}).(KeyVaultKeyReferenceKeyVaultOutput)
}

func (o KeyVaultKeyReferenceKeyVaultPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceKeyVault) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceResponse struct {
	KeyName    *string                               `pulumi:"keyName"`
	KeyVault   *KeyVaultKeyReferenceResponseKeyVault `pulumi:"keyVault"`
	KeyVersion *string                               `pulumi:"keyVersion"`
}

type KeyVaultKeyReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponseOutput() KeyVaultKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponseOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponseOutput) KeyVault() KeyVaultKeyReferenceResponseKeyVaultPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) *KeyVaultKeyReferenceResponseKeyVault { return v.KeyVault }).(KeyVaultKeyReferenceResponseKeyVaultPtrOutput)
}

func (o KeyVaultKeyReferenceResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponsePtrOutput) ToKeyVaultKeyReferenceResponsePtrOutput() KeyVaultKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponsePtrOutput) ToKeyVaultKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponsePtrOutput) Elem() KeyVaultKeyReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) KeyVaultKeyReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceResponse
		return ret
	}).(KeyVaultKeyReferenceResponseOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) KeyVault() KeyVaultKeyReferenceResponseKeyVaultPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *KeyVaultKeyReferenceResponseKeyVault {
		if v == nil {
			return nil
		}
		return v.KeyVault
	}).(KeyVaultKeyReferenceResponseKeyVaultPtrOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceResponseKeyVault struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type KeyVaultKeyReferenceResponseKeyVaultOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseKeyVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceResponseKeyVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseKeyVaultOutput) ToKeyVaultKeyReferenceResponseKeyVaultOutput() KeyVaultKeyReferenceResponseKeyVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseKeyVaultOutput) ToKeyVaultKeyReferenceResponseKeyVaultOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseKeyVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseKeyVaultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponseKeyVault) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponseKeyVaultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponseKeyVault) string { return v.Name }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceResponseKeyVaultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponseKeyVault) string { return v.Type }).(pulumi.StringOutput)
}

type KeyVaultKeyReferenceResponseKeyVaultPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseKeyVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceResponseKeyVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) ToKeyVaultKeyReferenceResponseKeyVaultPtrOutput() KeyVaultKeyReferenceResponseKeyVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) ToKeyVaultKeyReferenceResponseKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseKeyVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) Elem() KeyVaultKeyReferenceResponseKeyVaultOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponseKeyVault) KeyVaultKeyReferenceResponseKeyVault {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceResponseKeyVault
		return ret
	}).(KeyVaultKeyReferenceResponseKeyVaultOutput)
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponseKeyVault) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponseKeyVault) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponseKeyVault) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type PartnerContent struct {
	B2b *B2BPartnerContent `pulumi:"b2b"`
}





type PartnerContentInput interface {
	pulumi.Input

	ToPartnerContentOutput() PartnerContentOutput
	ToPartnerContentOutputWithContext(context.Context) PartnerContentOutput
}

type PartnerContentArgs struct {
	B2b B2BPartnerContentPtrInput `pulumi:"b2b"`
}

func (PartnerContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerContent)(nil)).Elem()
}

func (i PartnerContentArgs) ToPartnerContentOutput() PartnerContentOutput {
	return i.ToPartnerContentOutputWithContext(context.Background())
}

func (i PartnerContentArgs) ToPartnerContentOutputWithContext(ctx context.Context) PartnerContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerContentOutput)
}

func (i PartnerContentArgs) ToPartnerContentPtrOutput() PartnerContentPtrOutput {
	return i.ToPartnerContentPtrOutputWithContext(context.Background())
}

func (i PartnerContentArgs) ToPartnerContentPtrOutputWithContext(ctx context.Context) PartnerContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerContentOutput).ToPartnerContentPtrOutputWithContext(ctx)
}









type PartnerContentPtrInput interface {
	pulumi.Input

	ToPartnerContentPtrOutput() PartnerContentPtrOutput
	ToPartnerContentPtrOutputWithContext(context.Context) PartnerContentPtrOutput
}

type partnerContentPtrType PartnerContentArgs

func PartnerContentPtr(v *PartnerContentArgs) PartnerContentPtrInput {
	return (*partnerContentPtrType)(v)
}

func (*partnerContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerContent)(nil)).Elem()
}

func (i *partnerContentPtrType) ToPartnerContentPtrOutput() PartnerContentPtrOutput {
	return i.ToPartnerContentPtrOutputWithContext(context.Background())
}

func (i *partnerContentPtrType) ToPartnerContentPtrOutputWithContext(ctx context.Context) PartnerContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerContentPtrOutput)
}

type PartnerContentOutput struct{ *pulumi.OutputState }

func (PartnerContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerContent)(nil)).Elem()
}

func (o PartnerContentOutput) ToPartnerContentOutput() PartnerContentOutput {
	return o
}

func (o PartnerContentOutput) ToPartnerContentOutputWithContext(ctx context.Context) PartnerContentOutput {
	return o
}

func (o PartnerContentOutput) ToPartnerContentPtrOutput() PartnerContentPtrOutput {
	return o.ToPartnerContentPtrOutputWithContext(context.Background())
}

func (o PartnerContentOutput) ToPartnerContentPtrOutputWithContext(ctx context.Context) PartnerContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartnerContent) *PartnerContent {
		return &v
	}).(PartnerContentPtrOutput)
}

func (o PartnerContentOutput) B2b() B2BPartnerContentPtrOutput {
	return o.ApplyT(func(v PartnerContent) *B2BPartnerContent { return v.B2b }).(B2BPartnerContentPtrOutput)
}

type PartnerContentPtrOutput struct{ *pulumi.OutputState }

func (PartnerContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerContent)(nil)).Elem()
}

func (o PartnerContentPtrOutput) ToPartnerContentPtrOutput() PartnerContentPtrOutput {
	return o
}

func (o PartnerContentPtrOutput) ToPartnerContentPtrOutputWithContext(ctx context.Context) PartnerContentPtrOutput {
	return o
}

func (o PartnerContentPtrOutput) Elem() PartnerContentOutput {
	return o.ApplyT(func(v *PartnerContent) PartnerContent {
		if v != nil {
			return *v
		}
		var ret PartnerContent
		return ret
	}).(PartnerContentOutput)
}

func (o PartnerContentPtrOutput) B2b() B2BPartnerContentPtrOutput {
	return o.ApplyT(func(v *PartnerContent) *B2BPartnerContent {
		if v == nil {
			return nil
		}
		return v.B2b
	}).(B2BPartnerContentPtrOutput)
}

type PartnerContentResponse struct {
	B2b *B2BPartnerContentResponse `pulumi:"b2b"`
}

type PartnerContentResponseOutput struct{ *pulumi.OutputState }

func (PartnerContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerContentResponse)(nil)).Elem()
}

func (o PartnerContentResponseOutput) ToPartnerContentResponseOutput() PartnerContentResponseOutput {
	return o
}

func (o PartnerContentResponseOutput) ToPartnerContentResponseOutputWithContext(ctx context.Context) PartnerContentResponseOutput {
	return o
}

func (o PartnerContentResponseOutput) B2b() B2BPartnerContentResponsePtrOutput {
	return o.ApplyT(func(v PartnerContentResponse) *B2BPartnerContentResponse { return v.B2b }).(B2BPartnerContentResponsePtrOutput)
}

type PartnerContentResponsePtrOutput struct{ *pulumi.OutputState }

func (PartnerContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerContentResponse)(nil)).Elem()
}

func (o PartnerContentResponsePtrOutput) ToPartnerContentResponsePtrOutput() PartnerContentResponsePtrOutput {
	return o
}

func (o PartnerContentResponsePtrOutput) ToPartnerContentResponsePtrOutputWithContext(ctx context.Context) PartnerContentResponsePtrOutput {
	return o
}

func (o PartnerContentResponsePtrOutput) Elem() PartnerContentResponseOutput {
	return o.ApplyT(func(v *PartnerContentResponse) PartnerContentResponse {
		if v != nil {
			return *v
		}
		var ret PartnerContentResponse
		return ret
	}).(PartnerContentResponseOutput)
}

func (o PartnerContentResponsePtrOutput) B2b() B2BPartnerContentResponsePtrOutput {
	return o.ApplyT(func(v *PartnerContentResponse) *B2BPartnerContentResponse {
		if v == nil {
			return nil
		}
		return v.B2b
	}).(B2BPartnerContentResponsePtrOutput)
}

type X12AcknowledgementSettings struct {
	AcknowledgementControlNumberLowerBound *int    `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     *string `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     *string `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound *int    `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        *bool   `pulumi:"batchFunctionalAcknowledgements"`
	BatchImplementationAcknowledgements    *bool   `pulumi:"batchImplementationAcknowledgements"`
	BatchTechnicalAcknowledgements         *bool   `pulumi:"batchTechnicalAcknowledgements"`
	FunctionalAcknowledgementVersion       *string `pulumi:"functionalAcknowledgementVersion"`
	ImplementationAcknowledgementVersion   *string `pulumi:"implementationAcknowledgementVersion"`
	NeedFunctionalAcknowledgement          *bool   `pulumi:"needFunctionalAcknowledgement"`
	NeedImplementationAcknowledgement      *bool   `pulumi:"needImplementationAcknowledgement"`
	NeedLoopForValidMessages               *bool   `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           *bool   `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   *bool   `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         *bool   `pulumi:"sendSynchronousAcknowledgement"`
}





type X12AcknowledgementSettingsInput interface {
	pulumi.Input

	ToX12AcknowledgementSettingsOutput() X12AcknowledgementSettingsOutput
	ToX12AcknowledgementSettingsOutputWithContext(context.Context) X12AcknowledgementSettingsOutput
}

type X12AcknowledgementSettingsArgs struct {
	AcknowledgementControlNumberLowerBound pulumi.IntPtrInput    `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     pulumi.StringPtrInput `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     pulumi.StringPtrInput `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound pulumi.IntPtrInput    `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        pulumi.BoolPtrInput   `pulumi:"batchFunctionalAcknowledgements"`
	BatchImplementationAcknowledgements    pulumi.BoolPtrInput   `pulumi:"batchImplementationAcknowledgements"`
	BatchTechnicalAcknowledgements         pulumi.BoolPtrInput   `pulumi:"batchTechnicalAcknowledgements"`
	FunctionalAcknowledgementVersion       pulumi.StringPtrInput `pulumi:"functionalAcknowledgementVersion"`
	ImplementationAcknowledgementVersion   pulumi.StringPtrInput `pulumi:"implementationAcknowledgementVersion"`
	NeedFunctionalAcknowledgement          pulumi.BoolPtrInput   `pulumi:"needFunctionalAcknowledgement"`
	NeedImplementationAcknowledgement      pulumi.BoolPtrInput   `pulumi:"needImplementationAcknowledgement"`
	NeedLoopForValidMessages               pulumi.BoolPtrInput   `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           pulumi.BoolPtrInput   `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   pulumi.BoolPtrInput   `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         pulumi.BoolPtrInput   `pulumi:"sendSynchronousAcknowledgement"`
}

func (X12AcknowledgementSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AcknowledgementSettings)(nil)).Elem()
}

func (i X12AcknowledgementSettingsArgs) ToX12AcknowledgementSettingsOutput() X12AcknowledgementSettingsOutput {
	return i.ToX12AcknowledgementSettingsOutputWithContext(context.Background())
}

func (i X12AcknowledgementSettingsArgs) ToX12AcknowledgementSettingsOutputWithContext(ctx context.Context) X12AcknowledgementSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AcknowledgementSettingsOutput)
}

func (i X12AcknowledgementSettingsArgs) ToX12AcknowledgementSettingsPtrOutput() X12AcknowledgementSettingsPtrOutput {
	return i.ToX12AcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (i X12AcknowledgementSettingsArgs) ToX12AcknowledgementSettingsPtrOutputWithContext(ctx context.Context) X12AcknowledgementSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AcknowledgementSettingsOutput).ToX12AcknowledgementSettingsPtrOutputWithContext(ctx)
}









type X12AcknowledgementSettingsPtrInput interface {
	pulumi.Input

	ToX12AcknowledgementSettingsPtrOutput() X12AcknowledgementSettingsPtrOutput
	ToX12AcknowledgementSettingsPtrOutputWithContext(context.Context) X12AcknowledgementSettingsPtrOutput
}

type x12acknowledgementSettingsPtrType X12AcknowledgementSettingsArgs

func X12AcknowledgementSettingsPtr(v *X12AcknowledgementSettingsArgs) X12AcknowledgementSettingsPtrInput {
	return (*x12acknowledgementSettingsPtrType)(v)
}

func (*x12acknowledgementSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AcknowledgementSettings)(nil)).Elem()
}

func (i *x12acknowledgementSettingsPtrType) ToX12AcknowledgementSettingsPtrOutput() X12AcknowledgementSettingsPtrOutput {
	return i.ToX12AcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (i *x12acknowledgementSettingsPtrType) ToX12AcknowledgementSettingsPtrOutputWithContext(ctx context.Context) X12AcknowledgementSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AcknowledgementSettingsPtrOutput)
}

type X12AcknowledgementSettingsOutput struct{ *pulumi.OutputState }

func (X12AcknowledgementSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AcknowledgementSettings)(nil)).Elem()
}

func (o X12AcknowledgementSettingsOutput) ToX12AcknowledgementSettingsOutput() X12AcknowledgementSettingsOutput {
	return o
}

func (o X12AcknowledgementSettingsOutput) ToX12AcknowledgementSettingsOutputWithContext(ctx context.Context) X12AcknowledgementSettingsOutput {
	return o
}

func (o X12AcknowledgementSettingsOutput) ToX12AcknowledgementSettingsPtrOutput() X12AcknowledgementSettingsPtrOutput {
	return o.ToX12AcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (o X12AcknowledgementSettingsOutput) ToX12AcknowledgementSettingsPtrOutputWithContext(ctx context.Context) X12AcknowledgementSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12AcknowledgementSettings) *X12AcknowledgementSettings {
		return &v
	}).(X12AcknowledgementSettingsPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *int { return v.AcknowledgementControlNumberLowerBound }).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *string { return v.AcknowledgementControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *string { return v.AcknowledgementControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *int { return v.AcknowledgementControlNumberUpperBound }).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *bool { return v.BatchFunctionalAcknowledgements }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) BatchImplementationAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *bool { return v.BatchImplementationAcknowledgements }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *bool { return v.BatchTechnicalAcknowledgements }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) FunctionalAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *string { return v.FunctionalAcknowledgementVersion }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) ImplementationAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *string { return v.ImplementationAcknowledgementVersion }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *bool { return v.NeedFunctionalAcknowledgement }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) NeedImplementationAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *bool { return v.NeedImplementationAcknowledgement }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *bool { return v.NeedLoopForValidMessages }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *bool { return v.NeedTechnicalAcknowledgement }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *bool { return v.RolloverAcknowledgementControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *bool { return v.SendSynchronousAcknowledgement }).(pulumi.BoolPtrOutput)
}

type X12AcknowledgementSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12AcknowledgementSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AcknowledgementSettings)(nil)).Elem()
}

func (o X12AcknowledgementSettingsPtrOutput) ToX12AcknowledgementSettingsPtrOutput() X12AcknowledgementSettingsPtrOutput {
	return o
}

func (o X12AcknowledgementSettingsPtrOutput) ToX12AcknowledgementSettingsPtrOutputWithContext(ctx context.Context) X12AcknowledgementSettingsPtrOutput {
	return o
}

func (o X12AcknowledgementSettingsPtrOutput) Elem() X12AcknowledgementSettingsOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) X12AcknowledgementSettings {
		if v != nil {
			return *v
		}
		var ret X12AcknowledgementSettings
		return ret
	}).(X12AcknowledgementSettingsOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *int {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *int {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BatchFunctionalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) BatchImplementationAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BatchImplementationAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BatchTechnicalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) FunctionalAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalAcknowledgementVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) ImplementationAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.ImplementationAcknowledgementVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.NeedFunctionalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) NeedImplementationAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.NeedImplementationAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.NeedLoopForValidMessages
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.NeedTechnicalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverAcknowledgementControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SendSynchronousAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

type X12AcknowledgementSettingsResponse struct {
	AcknowledgementControlNumberLowerBound *int    `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     *string `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     *string `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound *int    `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        *bool   `pulumi:"batchFunctionalAcknowledgements"`
	BatchImplementationAcknowledgements    *bool   `pulumi:"batchImplementationAcknowledgements"`
	BatchTechnicalAcknowledgements         *bool   `pulumi:"batchTechnicalAcknowledgements"`
	FunctionalAcknowledgementVersion       *string `pulumi:"functionalAcknowledgementVersion"`
	ImplementationAcknowledgementVersion   *string `pulumi:"implementationAcknowledgementVersion"`
	NeedFunctionalAcknowledgement          *bool   `pulumi:"needFunctionalAcknowledgement"`
	NeedImplementationAcknowledgement      *bool   `pulumi:"needImplementationAcknowledgement"`
	NeedLoopForValidMessages               *bool   `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           *bool   `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   *bool   `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         *bool   `pulumi:"sendSynchronousAcknowledgement"`
}

type X12AcknowledgementSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12AcknowledgementSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AcknowledgementSettingsResponse)(nil)).Elem()
}

func (o X12AcknowledgementSettingsResponseOutput) ToX12AcknowledgementSettingsResponseOutput() X12AcknowledgementSettingsResponseOutput {
	return o
}

func (o X12AcknowledgementSettingsResponseOutput) ToX12AcknowledgementSettingsResponseOutputWithContext(ctx context.Context) X12AcknowledgementSettingsResponseOutput {
	return o
}

func (o X12AcknowledgementSettingsResponseOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *int { return v.AcknowledgementControlNumberLowerBound }).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *string { return v.AcknowledgementControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *string { return v.AcknowledgementControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *int { return v.AcknowledgementControlNumberUpperBound }).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *bool { return v.BatchFunctionalAcknowledgements }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) BatchImplementationAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *bool { return v.BatchImplementationAcknowledgements }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *bool { return v.BatchTechnicalAcknowledgements }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) FunctionalAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *string { return v.FunctionalAcknowledgementVersion }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) ImplementationAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *string { return v.ImplementationAcknowledgementVersion }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *bool { return v.NeedFunctionalAcknowledgement }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) NeedImplementationAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *bool { return v.NeedImplementationAcknowledgement }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *bool { return v.NeedLoopForValidMessages }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *bool { return v.NeedTechnicalAcknowledgement }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *bool { return v.RolloverAcknowledgementControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *bool { return v.SendSynchronousAcknowledgement }).(pulumi.BoolPtrOutput)
}

type X12AcknowledgementSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12AcknowledgementSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AcknowledgementSettingsResponse)(nil)).Elem()
}

func (o X12AcknowledgementSettingsResponsePtrOutput) ToX12AcknowledgementSettingsResponsePtrOutput() X12AcknowledgementSettingsResponsePtrOutput {
	return o
}

func (o X12AcknowledgementSettingsResponsePtrOutput) ToX12AcknowledgementSettingsResponsePtrOutputWithContext(ctx context.Context) X12AcknowledgementSettingsResponsePtrOutput {
	return o
}

func (o X12AcknowledgementSettingsResponsePtrOutput) Elem() X12AcknowledgementSettingsResponseOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) X12AcknowledgementSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12AcknowledgementSettingsResponse
		return ret
	}).(X12AcknowledgementSettingsResponseOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.BatchFunctionalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) BatchImplementationAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.BatchImplementationAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.BatchTechnicalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) FunctionalAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalAcknowledgementVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) ImplementationAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImplementationAcknowledgementVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.NeedFunctionalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) NeedImplementationAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.NeedImplementationAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.NeedLoopForValidMessages
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.NeedTechnicalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverAcknowledgementControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendSynchronousAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

type X12AgreementContent struct {
	ReceiveAgreement *X12OneWayAgreement `pulumi:"receiveAgreement"`
	SendAgreement    *X12OneWayAgreement `pulumi:"sendAgreement"`
}





type X12AgreementContentInput interface {
	pulumi.Input

	ToX12AgreementContentOutput() X12AgreementContentOutput
	ToX12AgreementContentOutputWithContext(context.Context) X12AgreementContentOutput
}

type X12AgreementContentArgs struct {
	ReceiveAgreement X12OneWayAgreementPtrInput `pulumi:"receiveAgreement"`
	SendAgreement    X12OneWayAgreementPtrInput `pulumi:"sendAgreement"`
}

func (X12AgreementContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AgreementContent)(nil)).Elem()
}

func (i X12AgreementContentArgs) ToX12AgreementContentOutput() X12AgreementContentOutput {
	return i.ToX12AgreementContentOutputWithContext(context.Background())
}

func (i X12AgreementContentArgs) ToX12AgreementContentOutputWithContext(ctx context.Context) X12AgreementContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AgreementContentOutput)
}

func (i X12AgreementContentArgs) ToX12AgreementContentPtrOutput() X12AgreementContentPtrOutput {
	return i.ToX12AgreementContentPtrOutputWithContext(context.Background())
}

func (i X12AgreementContentArgs) ToX12AgreementContentPtrOutputWithContext(ctx context.Context) X12AgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AgreementContentOutput).ToX12AgreementContentPtrOutputWithContext(ctx)
}









type X12AgreementContentPtrInput interface {
	pulumi.Input

	ToX12AgreementContentPtrOutput() X12AgreementContentPtrOutput
	ToX12AgreementContentPtrOutputWithContext(context.Context) X12AgreementContentPtrOutput
}

type x12agreementContentPtrType X12AgreementContentArgs

func X12AgreementContentPtr(v *X12AgreementContentArgs) X12AgreementContentPtrInput {
	return (*x12agreementContentPtrType)(v)
}

func (*x12agreementContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AgreementContent)(nil)).Elem()
}

func (i *x12agreementContentPtrType) ToX12AgreementContentPtrOutput() X12AgreementContentPtrOutput {
	return i.ToX12AgreementContentPtrOutputWithContext(context.Background())
}

func (i *x12agreementContentPtrType) ToX12AgreementContentPtrOutputWithContext(ctx context.Context) X12AgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AgreementContentPtrOutput)
}

type X12AgreementContentOutput struct{ *pulumi.OutputState }

func (X12AgreementContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AgreementContent)(nil)).Elem()
}

func (o X12AgreementContentOutput) ToX12AgreementContentOutput() X12AgreementContentOutput {
	return o
}

func (o X12AgreementContentOutput) ToX12AgreementContentOutputWithContext(ctx context.Context) X12AgreementContentOutput {
	return o
}

func (o X12AgreementContentOutput) ToX12AgreementContentPtrOutput() X12AgreementContentPtrOutput {
	return o.ToX12AgreementContentPtrOutputWithContext(context.Background())
}

func (o X12AgreementContentOutput) ToX12AgreementContentPtrOutputWithContext(ctx context.Context) X12AgreementContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12AgreementContent) *X12AgreementContent {
		return &v
	}).(X12AgreementContentPtrOutput)
}

func (o X12AgreementContentOutput) ReceiveAgreement() X12OneWayAgreementPtrOutput {
	return o.ApplyT(func(v X12AgreementContent) *X12OneWayAgreement { return v.ReceiveAgreement }).(X12OneWayAgreementPtrOutput)
}

func (o X12AgreementContentOutput) SendAgreement() X12OneWayAgreementPtrOutput {
	return o.ApplyT(func(v X12AgreementContent) *X12OneWayAgreement { return v.SendAgreement }).(X12OneWayAgreementPtrOutput)
}

type X12AgreementContentPtrOutput struct{ *pulumi.OutputState }

func (X12AgreementContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AgreementContent)(nil)).Elem()
}

func (o X12AgreementContentPtrOutput) ToX12AgreementContentPtrOutput() X12AgreementContentPtrOutput {
	return o
}

func (o X12AgreementContentPtrOutput) ToX12AgreementContentPtrOutputWithContext(ctx context.Context) X12AgreementContentPtrOutput {
	return o
}

func (o X12AgreementContentPtrOutput) Elem() X12AgreementContentOutput {
	return o.ApplyT(func(v *X12AgreementContent) X12AgreementContent {
		if v != nil {
			return *v
		}
		var ret X12AgreementContent
		return ret
	}).(X12AgreementContentOutput)
}

func (o X12AgreementContentPtrOutput) ReceiveAgreement() X12OneWayAgreementPtrOutput {
	return o.ApplyT(func(v *X12AgreementContent) *X12OneWayAgreement {
		if v == nil {
			return nil
		}
		return v.ReceiveAgreement
	}).(X12OneWayAgreementPtrOutput)
}

func (o X12AgreementContentPtrOutput) SendAgreement() X12OneWayAgreementPtrOutput {
	return o.ApplyT(func(v *X12AgreementContent) *X12OneWayAgreement {
		if v == nil {
			return nil
		}
		return v.SendAgreement
	}).(X12OneWayAgreementPtrOutput)
}

type X12AgreementContentResponse struct {
	ReceiveAgreement *X12OneWayAgreementResponse `pulumi:"receiveAgreement"`
	SendAgreement    *X12OneWayAgreementResponse `pulumi:"sendAgreement"`
}

type X12AgreementContentResponseOutput struct{ *pulumi.OutputState }

func (X12AgreementContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AgreementContentResponse)(nil)).Elem()
}

func (o X12AgreementContentResponseOutput) ToX12AgreementContentResponseOutput() X12AgreementContentResponseOutput {
	return o
}

func (o X12AgreementContentResponseOutput) ToX12AgreementContentResponseOutputWithContext(ctx context.Context) X12AgreementContentResponseOutput {
	return o
}

func (o X12AgreementContentResponseOutput) ReceiveAgreement() X12OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v X12AgreementContentResponse) *X12OneWayAgreementResponse { return v.ReceiveAgreement }).(X12OneWayAgreementResponsePtrOutput)
}

func (o X12AgreementContentResponseOutput) SendAgreement() X12OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v X12AgreementContentResponse) *X12OneWayAgreementResponse { return v.SendAgreement }).(X12OneWayAgreementResponsePtrOutput)
}

type X12AgreementContentResponsePtrOutput struct{ *pulumi.OutputState }

func (X12AgreementContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AgreementContentResponse)(nil)).Elem()
}

func (o X12AgreementContentResponsePtrOutput) ToX12AgreementContentResponsePtrOutput() X12AgreementContentResponsePtrOutput {
	return o
}

func (o X12AgreementContentResponsePtrOutput) ToX12AgreementContentResponsePtrOutputWithContext(ctx context.Context) X12AgreementContentResponsePtrOutput {
	return o
}

func (o X12AgreementContentResponsePtrOutput) Elem() X12AgreementContentResponseOutput {
	return o.ApplyT(func(v *X12AgreementContentResponse) X12AgreementContentResponse {
		if v != nil {
			return *v
		}
		var ret X12AgreementContentResponse
		return ret
	}).(X12AgreementContentResponseOutput)
}

func (o X12AgreementContentResponsePtrOutput) ReceiveAgreement() X12OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *X12AgreementContentResponse) *X12OneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return v.ReceiveAgreement
	}).(X12OneWayAgreementResponsePtrOutput)
}

func (o X12AgreementContentResponsePtrOutput) SendAgreement() X12OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *X12AgreementContentResponse) *X12OneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return v.SendAgreement
	}).(X12OneWayAgreementResponsePtrOutput)
}

type X12DelimiterOverrides struct {
	ComponentSeparator         *int                     `pulumi:"componentSeparator"`
	DataElementSeparator       *int                     `pulumi:"dataElementSeparator"`
	MessageId                  *string                  `pulumi:"messageId"`
	ProtocolVersion            *string                  `pulumi:"protocolVersion"`
	ReplaceCharacter           *int                     `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload *bool                    `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          *int                     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    *SegmentTerminatorSuffix `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace            *string                  `pulumi:"targetNamespace"`
}





type X12DelimiterOverridesInput interface {
	pulumi.Input

	ToX12DelimiterOverridesOutput() X12DelimiterOverridesOutput
	ToX12DelimiterOverridesOutputWithContext(context.Context) X12DelimiterOverridesOutput
}

type X12DelimiterOverridesArgs struct {
	ComponentSeparator         pulumi.IntPtrInput              `pulumi:"componentSeparator"`
	DataElementSeparator       pulumi.IntPtrInput              `pulumi:"dataElementSeparator"`
	MessageId                  pulumi.StringPtrInput           `pulumi:"messageId"`
	ProtocolVersion            pulumi.StringPtrInput           `pulumi:"protocolVersion"`
	ReplaceCharacter           pulumi.IntPtrInput              `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload pulumi.BoolPtrInput             `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          pulumi.IntPtrInput              `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    SegmentTerminatorSuffixPtrInput `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace            pulumi.StringPtrInput           `pulumi:"targetNamespace"`
}

func (X12DelimiterOverridesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12DelimiterOverrides)(nil)).Elem()
}

func (i X12DelimiterOverridesArgs) ToX12DelimiterOverridesOutput() X12DelimiterOverridesOutput {
	return i.ToX12DelimiterOverridesOutputWithContext(context.Background())
}

func (i X12DelimiterOverridesArgs) ToX12DelimiterOverridesOutputWithContext(ctx context.Context) X12DelimiterOverridesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12DelimiterOverridesOutput)
}





type X12DelimiterOverridesArrayInput interface {
	pulumi.Input

	ToX12DelimiterOverridesArrayOutput() X12DelimiterOverridesArrayOutput
	ToX12DelimiterOverridesArrayOutputWithContext(context.Context) X12DelimiterOverridesArrayOutput
}

type X12DelimiterOverridesArray []X12DelimiterOverridesInput

func (X12DelimiterOverridesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12DelimiterOverrides)(nil)).Elem()
}

func (i X12DelimiterOverridesArray) ToX12DelimiterOverridesArrayOutput() X12DelimiterOverridesArrayOutput {
	return i.ToX12DelimiterOverridesArrayOutputWithContext(context.Background())
}

func (i X12DelimiterOverridesArray) ToX12DelimiterOverridesArrayOutputWithContext(ctx context.Context) X12DelimiterOverridesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12DelimiterOverridesArrayOutput)
}

type X12DelimiterOverridesOutput struct{ *pulumi.OutputState }

func (X12DelimiterOverridesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12DelimiterOverrides)(nil)).Elem()
}

func (o X12DelimiterOverridesOutput) ToX12DelimiterOverridesOutput() X12DelimiterOverridesOutput {
	return o
}

func (o X12DelimiterOverridesOutput) ToX12DelimiterOverridesOutputWithContext(ctx context.Context) X12DelimiterOverridesOutput {
	return o
}

func (o X12DelimiterOverridesOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *int { return v.ComponentSeparator }).(pulumi.IntPtrOutput)
}

func (o X12DelimiterOverridesOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *int { return v.DataElementSeparator }).(pulumi.IntPtrOutput)
}

func (o X12DelimiterOverridesOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o X12DelimiterOverridesOutput) ProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *string { return v.ProtocolVersion }).(pulumi.StringPtrOutput)
}

func (o X12DelimiterOverridesOutput) ReplaceCharacter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *int { return v.ReplaceCharacter }).(pulumi.IntPtrOutput)
}

func (o X12DelimiterOverridesOutput) ReplaceSeparatorsInPayload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *bool { return v.ReplaceSeparatorsInPayload }).(pulumi.BoolPtrOutput)
}

func (o X12DelimiterOverridesOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *int { return v.SegmentTerminator }).(pulumi.IntPtrOutput)
}

func (o X12DelimiterOverridesOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *SegmentTerminatorSuffix { return v.SegmentTerminatorSuffix }).(SegmentTerminatorSuffixPtrOutput)
}

func (o X12DelimiterOverridesOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type X12DelimiterOverridesArrayOutput struct{ *pulumi.OutputState }

func (X12DelimiterOverridesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12DelimiterOverrides)(nil)).Elem()
}

func (o X12DelimiterOverridesArrayOutput) ToX12DelimiterOverridesArrayOutput() X12DelimiterOverridesArrayOutput {
	return o
}

func (o X12DelimiterOverridesArrayOutput) ToX12DelimiterOverridesArrayOutputWithContext(ctx context.Context) X12DelimiterOverridesArrayOutput {
	return o
}

func (o X12DelimiterOverridesArrayOutput) Index(i pulumi.IntInput) X12DelimiterOverridesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12DelimiterOverrides {
		return vs[0].([]X12DelimiterOverrides)[vs[1].(int)]
	}).(X12DelimiterOverridesOutput)
}

type X12DelimiterOverridesResponse struct {
	ComponentSeparator         *int    `pulumi:"componentSeparator"`
	DataElementSeparator       *int    `pulumi:"dataElementSeparator"`
	MessageId                  *string `pulumi:"messageId"`
	ProtocolVersion            *string `pulumi:"protocolVersion"`
	ReplaceCharacter           *int    `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload *bool   `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          *int    `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    *string `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace            *string `pulumi:"targetNamespace"`
}

type X12DelimiterOverridesResponseOutput struct{ *pulumi.OutputState }

func (X12DelimiterOverridesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12DelimiterOverridesResponse)(nil)).Elem()
}

func (o X12DelimiterOverridesResponseOutput) ToX12DelimiterOverridesResponseOutput() X12DelimiterOverridesResponseOutput {
	return o
}

func (o X12DelimiterOverridesResponseOutput) ToX12DelimiterOverridesResponseOutputWithContext(ctx context.Context) X12DelimiterOverridesResponseOutput {
	return o
}

func (o X12DelimiterOverridesResponseOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *int { return v.ComponentSeparator }).(pulumi.IntPtrOutput)
}

func (o X12DelimiterOverridesResponseOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *int { return v.DataElementSeparator }).(pulumi.IntPtrOutput)
}

func (o X12DelimiterOverridesResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o X12DelimiterOverridesResponseOutput) ProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *string { return v.ProtocolVersion }).(pulumi.StringPtrOutput)
}

func (o X12DelimiterOverridesResponseOutput) ReplaceCharacter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *int { return v.ReplaceCharacter }).(pulumi.IntPtrOutput)
}

func (o X12DelimiterOverridesResponseOutput) ReplaceSeparatorsInPayload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *bool { return v.ReplaceSeparatorsInPayload }).(pulumi.BoolPtrOutput)
}

func (o X12DelimiterOverridesResponseOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *int { return v.SegmentTerminator }).(pulumi.IntPtrOutput)
}

func (o X12DelimiterOverridesResponseOutput) SegmentTerminatorSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *string { return v.SegmentTerminatorSuffix }).(pulumi.StringPtrOutput)
}

func (o X12DelimiterOverridesResponseOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type X12DelimiterOverridesResponseArrayOutput struct{ *pulumi.OutputState }

func (X12DelimiterOverridesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12DelimiterOverridesResponse)(nil)).Elem()
}

func (o X12DelimiterOverridesResponseArrayOutput) ToX12DelimiterOverridesResponseArrayOutput() X12DelimiterOverridesResponseArrayOutput {
	return o
}

func (o X12DelimiterOverridesResponseArrayOutput) ToX12DelimiterOverridesResponseArrayOutputWithContext(ctx context.Context) X12DelimiterOverridesResponseArrayOutput {
	return o
}

func (o X12DelimiterOverridesResponseArrayOutput) Index(i pulumi.IntInput) X12DelimiterOverridesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12DelimiterOverridesResponse {
		return vs[0].([]X12DelimiterOverridesResponse)[vs[1].(int)]
	}).(X12DelimiterOverridesResponseOutput)
}

type X12EnvelopeOverride struct {
	DateFormat               *X12DateFormat `pulumi:"dateFormat"`
	FunctionalIdentifierCode *string        `pulumi:"functionalIdentifierCode"`
	HeaderVersion            *string        `pulumi:"headerVersion"`
	MessageId                *string        `pulumi:"messageId"`
	ProtocolVersion          *string        `pulumi:"protocolVersion"`
	ReceiverApplicationId    *string        `pulumi:"receiverApplicationId"`
	ResponsibleAgencyCode    *int           `pulumi:"responsibleAgencyCode"`
	SenderApplicationId      *string        `pulumi:"senderApplicationId"`
	TargetNamespace          *string        `pulumi:"targetNamespace"`
	TimeFormat               *X12TimeFormat `pulumi:"timeFormat"`
}





type X12EnvelopeOverrideInput interface {
	pulumi.Input

	ToX12EnvelopeOverrideOutput() X12EnvelopeOverrideOutput
	ToX12EnvelopeOverrideOutputWithContext(context.Context) X12EnvelopeOverrideOutput
}

type X12EnvelopeOverrideArgs struct {
	DateFormat               X12DateFormatPtrInput `pulumi:"dateFormat"`
	FunctionalIdentifierCode pulumi.StringPtrInput `pulumi:"functionalIdentifierCode"`
	HeaderVersion            pulumi.StringPtrInput `pulumi:"headerVersion"`
	MessageId                pulumi.StringPtrInput `pulumi:"messageId"`
	ProtocolVersion          pulumi.StringPtrInput `pulumi:"protocolVersion"`
	ReceiverApplicationId    pulumi.StringPtrInput `pulumi:"receiverApplicationId"`
	ResponsibleAgencyCode    pulumi.IntPtrInput    `pulumi:"responsibleAgencyCode"`
	SenderApplicationId      pulumi.StringPtrInput `pulumi:"senderApplicationId"`
	TargetNamespace          pulumi.StringPtrInput `pulumi:"targetNamespace"`
	TimeFormat               X12TimeFormatPtrInput `pulumi:"timeFormat"`
}

func (X12EnvelopeOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeOverride)(nil)).Elem()
}

func (i X12EnvelopeOverrideArgs) ToX12EnvelopeOverrideOutput() X12EnvelopeOverrideOutput {
	return i.ToX12EnvelopeOverrideOutputWithContext(context.Background())
}

func (i X12EnvelopeOverrideArgs) ToX12EnvelopeOverrideOutputWithContext(ctx context.Context) X12EnvelopeOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12EnvelopeOverrideOutput)
}





type X12EnvelopeOverrideArrayInput interface {
	pulumi.Input

	ToX12EnvelopeOverrideArrayOutput() X12EnvelopeOverrideArrayOutput
	ToX12EnvelopeOverrideArrayOutputWithContext(context.Context) X12EnvelopeOverrideArrayOutput
}

type X12EnvelopeOverrideArray []X12EnvelopeOverrideInput

func (X12EnvelopeOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12EnvelopeOverride)(nil)).Elem()
}

func (i X12EnvelopeOverrideArray) ToX12EnvelopeOverrideArrayOutput() X12EnvelopeOverrideArrayOutput {
	return i.ToX12EnvelopeOverrideArrayOutputWithContext(context.Background())
}

func (i X12EnvelopeOverrideArray) ToX12EnvelopeOverrideArrayOutputWithContext(ctx context.Context) X12EnvelopeOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12EnvelopeOverrideArrayOutput)
}

type X12EnvelopeOverrideOutput struct{ *pulumi.OutputState }

func (X12EnvelopeOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeOverride)(nil)).Elem()
}

func (o X12EnvelopeOverrideOutput) ToX12EnvelopeOverrideOutput() X12EnvelopeOverrideOutput {
	return o
}

func (o X12EnvelopeOverrideOutput) ToX12EnvelopeOverrideOutputWithContext(ctx context.Context) X12EnvelopeOverrideOutput {
	return o
}

func (o X12EnvelopeOverrideOutput) DateFormat() X12DateFormatPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *X12DateFormat { return v.DateFormat }).(X12DateFormatPtrOutput)
}

func (o X12EnvelopeOverrideOutput) FunctionalIdentifierCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *string { return v.FunctionalIdentifierCode }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideOutput) HeaderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *string { return v.HeaderVersion }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideOutput) ProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *string { return v.ProtocolVersion }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *string { return v.ReceiverApplicationId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideOutput) ResponsibleAgencyCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *int { return v.ResponsibleAgencyCode }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeOverrideOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideOutput) TimeFormat() X12TimeFormatPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *X12TimeFormat { return v.TimeFormat }).(X12TimeFormatPtrOutput)
}

type X12EnvelopeOverrideArrayOutput struct{ *pulumi.OutputState }

func (X12EnvelopeOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12EnvelopeOverride)(nil)).Elem()
}

func (o X12EnvelopeOverrideArrayOutput) ToX12EnvelopeOverrideArrayOutput() X12EnvelopeOverrideArrayOutput {
	return o
}

func (o X12EnvelopeOverrideArrayOutput) ToX12EnvelopeOverrideArrayOutputWithContext(ctx context.Context) X12EnvelopeOverrideArrayOutput {
	return o
}

func (o X12EnvelopeOverrideArrayOutput) Index(i pulumi.IntInput) X12EnvelopeOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12EnvelopeOverride {
		return vs[0].([]X12EnvelopeOverride)[vs[1].(int)]
	}).(X12EnvelopeOverrideOutput)
}

type X12EnvelopeOverrideResponse struct {
	DateFormat               *string `pulumi:"dateFormat"`
	FunctionalIdentifierCode *string `pulumi:"functionalIdentifierCode"`
	HeaderVersion            *string `pulumi:"headerVersion"`
	MessageId                *string `pulumi:"messageId"`
	ProtocolVersion          *string `pulumi:"protocolVersion"`
	ReceiverApplicationId    *string `pulumi:"receiverApplicationId"`
	ResponsibleAgencyCode    *int    `pulumi:"responsibleAgencyCode"`
	SenderApplicationId      *string `pulumi:"senderApplicationId"`
	TargetNamespace          *string `pulumi:"targetNamespace"`
	TimeFormat               *string `pulumi:"timeFormat"`
}

type X12EnvelopeOverrideResponseOutput struct{ *pulumi.OutputState }

func (X12EnvelopeOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeOverrideResponse)(nil)).Elem()
}

func (o X12EnvelopeOverrideResponseOutput) ToX12EnvelopeOverrideResponseOutput() X12EnvelopeOverrideResponseOutput {
	return o
}

func (o X12EnvelopeOverrideResponseOutput) ToX12EnvelopeOverrideResponseOutputWithContext(ctx context.Context) X12EnvelopeOverrideResponseOutput {
	return o
}

func (o X12EnvelopeOverrideResponseOutput) DateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *string { return v.DateFormat }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideResponseOutput) FunctionalIdentifierCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *string { return v.FunctionalIdentifierCode }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideResponseOutput) HeaderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *string { return v.HeaderVersion }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideResponseOutput) ProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *string { return v.ProtocolVersion }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideResponseOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *string { return v.ReceiverApplicationId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideResponseOutput) ResponsibleAgencyCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *int { return v.ResponsibleAgencyCode }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeOverrideResponseOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideResponseOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideResponseOutput) TimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *string { return v.TimeFormat }).(pulumi.StringPtrOutput)
}

type X12EnvelopeOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (X12EnvelopeOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12EnvelopeOverrideResponse)(nil)).Elem()
}

func (o X12EnvelopeOverrideResponseArrayOutput) ToX12EnvelopeOverrideResponseArrayOutput() X12EnvelopeOverrideResponseArrayOutput {
	return o
}

func (o X12EnvelopeOverrideResponseArrayOutput) ToX12EnvelopeOverrideResponseArrayOutputWithContext(ctx context.Context) X12EnvelopeOverrideResponseArrayOutput {
	return o
}

func (o X12EnvelopeOverrideResponseArrayOutput) Index(i pulumi.IntInput) X12EnvelopeOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12EnvelopeOverrideResponse {
		return vs[0].([]X12EnvelopeOverrideResponse)[vs[1].(int)]
	}).(X12EnvelopeOverrideResponseOutput)
}

type X12EnvelopeSettings struct {
	ControlStandardsId                           *int            `pulumi:"controlStandardsId"`
	ControlVersionNumber                         *string         `pulumi:"controlVersionNumber"`
	EnableDefaultGroupHeaders                    *bool           `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            *string         `pulumi:"functionalGroupId"`
	GroupControlNumberLowerBound                 *int            `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberUpperBound                 *int            `pulumi:"groupControlNumberUpperBound"`
	GroupHeaderAgencyCode                        *string         `pulumi:"groupHeaderAgencyCode"`
	GroupHeaderDateFormat                        *X12DateFormat  `pulumi:"groupHeaderDateFormat"`
	GroupHeaderTimeFormat                        *X12TimeFormat  `pulumi:"groupHeaderTimeFormat"`
	GroupHeaderVersion                           *string         `pulumi:"groupHeaderVersion"`
	InterchangeControlNumberLowerBound           *int            `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberUpperBound           *int            `pulumi:"interchangeControlNumberUpperBound"`
	OverwriteExistingTransactionSetControlNumber *bool           `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ReceiverApplicationId                        *string         `pulumi:"receiverApplicationId"`
	RolloverGroupControlNumber                   *bool           `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             *bool           `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          *bool           `pulumi:"rolloverTransactionSetControlNumber"`
	SenderApplicationId                          *string         `pulumi:"senderApplicationId"`
	TransactionSetControlNumberLowerBound        *int            `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            *string         `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            *string         `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        *int            `pulumi:"transactionSetControlNumberUpperBound"`
	UsageIndicator                               *UsageIndicator `pulumi:"usageIndicator"`
	UseControlStandardsIdAsRepetitionCharacter   *bool           `pulumi:"useControlStandardsIdAsRepetitionCharacter"`
}





type X12EnvelopeSettingsInput interface {
	pulumi.Input

	ToX12EnvelopeSettingsOutput() X12EnvelopeSettingsOutput
	ToX12EnvelopeSettingsOutputWithContext(context.Context) X12EnvelopeSettingsOutput
}

type X12EnvelopeSettingsArgs struct {
	ControlStandardsId                           pulumi.IntPtrInput     `pulumi:"controlStandardsId"`
	ControlVersionNumber                         pulumi.StringPtrInput  `pulumi:"controlVersionNumber"`
	EnableDefaultGroupHeaders                    pulumi.BoolPtrInput    `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            pulumi.StringPtrInput  `pulumi:"functionalGroupId"`
	GroupControlNumberLowerBound                 pulumi.IntPtrInput     `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberUpperBound                 pulumi.IntPtrInput     `pulumi:"groupControlNumberUpperBound"`
	GroupHeaderAgencyCode                        pulumi.StringPtrInput  `pulumi:"groupHeaderAgencyCode"`
	GroupHeaderDateFormat                        X12DateFormatPtrInput  `pulumi:"groupHeaderDateFormat"`
	GroupHeaderTimeFormat                        X12TimeFormatPtrInput  `pulumi:"groupHeaderTimeFormat"`
	GroupHeaderVersion                           pulumi.StringPtrInput  `pulumi:"groupHeaderVersion"`
	InterchangeControlNumberLowerBound           pulumi.IntPtrInput     `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberUpperBound           pulumi.IntPtrInput     `pulumi:"interchangeControlNumberUpperBound"`
	OverwriteExistingTransactionSetControlNumber pulumi.BoolPtrInput    `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ReceiverApplicationId                        pulumi.StringPtrInput  `pulumi:"receiverApplicationId"`
	RolloverGroupControlNumber                   pulumi.BoolPtrInput    `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             pulumi.BoolPtrInput    `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          pulumi.BoolPtrInput    `pulumi:"rolloverTransactionSetControlNumber"`
	SenderApplicationId                          pulumi.StringPtrInput  `pulumi:"senderApplicationId"`
	TransactionSetControlNumberLowerBound        pulumi.IntPtrInput     `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            pulumi.StringPtrInput  `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            pulumi.StringPtrInput  `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        pulumi.IntPtrInput     `pulumi:"transactionSetControlNumberUpperBound"`
	UsageIndicator                               UsageIndicatorPtrInput `pulumi:"usageIndicator"`
	UseControlStandardsIdAsRepetitionCharacter   pulumi.BoolPtrInput    `pulumi:"useControlStandardsIdAsRepetitionCharacter"`
}

func (X12EnvelopeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeSettings)(nil)).Elem()
}

func (i X12EnvelopeSettingsArgs) ToX12EnvelopeSettingsOutput() X12EnvelopeSettingsOutput {
	return i.ToX12EnvelopeSettingsOutputWithContext(context.Background())
}

func (i X12EnvelopeSettingsArgs) ToX12EnvelopeSettingsOutputWithContext(ctx context.Context) X12EnvelopeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12EnvelopeSettingsOutput)
}

func (i X12EnvelopeSettingsArgs) ToX12EnvelopeSettingsPtrOutput() X12EnvelopeSettingsPtrOutput {
	return i.ToX12EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i X12EnvelopeSettingsArgs) ToX12EnvelopeSettingsPtrOutputWithContext(ctx context.Context) X12EnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12EnvelopeSettingsOutput).ToX12EnvelopeSettingsPtrOutputWithContext(ctx)
}









type X12EnvelopeSettingsPtrInput interface {
	pulumi.Input

	ToX12EnvelopeSettingsPtrOutput() X12EnvelopeSettingsPtrOutput
	ToX12EnvelopeSettingsPtrOutputWithContext(context.Context) X12EnvelopeSettingsPtrOutput
}

type x12envelopeSettingsPtrType X12EnvelopeSettingsArgs

func X12EnvelopeSettingsPtr(v *X12EnvelopeSettingsArgs) X12EnvelopeSettingsPtrInput {
	return (*x12envelopeSettingsPtrType)(v)
}

func (*x12envelopeSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12EnvelopeSettings)(nil)).Elem()
}

func (i *x12envelopeSettingsPtrType) ToX12EnvelopeSettingsPtrOutput() X12EnvelopeSettingsPtrOutput {
	return i.ToX12EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i *x12envelopeSettingsPtrType) ToX12EnvelopeSettingsPtrOutputWithContext(ctx context.Context) X12EnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12EnvelopeSettingsPtrOutput)
}

type X12EnvelopeSettingsOutput struct{ *pulumi.OutputState }

func (X12EnvelopeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeSettings)(nil)).Elem()
}

func (o X12EnvelopeSettingsOutput) ToX12EnvelopeSettingsOutput() X12EnvelopeSettingsOutput {
	return o
}

func (o X12EnvelopeSettingsOutput) ToX12EnvelopeSettingsOutputWithContext(ctx context.Context) X12EnvelopeSettingsOutput {
	return o
}

func (o X12EnvelopeSettingsOutput) ToX12EnvelopeSettingsPtrOutput() X12EnvelopeSettingsPtrOutput {
	return o.ToX12EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (o X12EnvelopeSettingsOutput) ToX12EnvelopeSettingsPtrOutputWithContext(ctx context.Context) X12EnvelopeSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12EnvelopeSettings) *X12EnvelopeSettings {
		return &v
	}).(X12EnvelopeSettingsPtrOutput)
}

func (o X12EnvelopeSettingsOutput) ControlStandardsId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *int { return v.ControlStandardsId }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsOutput) ControlVersionNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.ControlVersionNumber }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *bool { return v.EnableDefaultGroupHeaders }).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) GroupControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *int { return v.GroupControlNumberLowerBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsOutput) GroupControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *int { return v.GroupControlNumberUpperBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsOutput) GroupHeaderAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.GroupHeaderAgencyCode }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) GroupHeaderDateFormat() X12DateFormatPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *X12DateFormat { return v.GroupHeaderDateFormat }).(X12DateFormatPtrOutput)
}

func (o X12EnvelopeSettingsOutput) GroupHeaderTimeFormat() X12TimeFormatPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *X12TimeFormat { return v.GroupHeaderTimeFormat }).(X12TimeFormatPtrOutput)
}

func (o X12EnvelopeSettingsOutput) GroupHeaderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.GroupHeaderVersion }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) InterchangeControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *int { return v.InterchangeControlNumberLowerBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsOutput) InterchangeControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *int { return v.InterchangeControlNumberUpperBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *bool { return v.OverwriteExistingTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.ReceiverApplicationId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *bool { return v.RolloverGroupControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *bool { return v.RolloverInterchangeControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *bool { return v.RolloverTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) TransactionSetControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *int { return v.TransactionSetControlNumberLowerBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.TransactionSetControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.TransactionSetControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) TransactionSetControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *int { return v.TransactionSetControlNumberUpperBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsOutput) UsageIndicator() UsageIndicatorPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *UsageIndicator { return v.UsageIndicator }).(UsageIndicatorPtrOutput)
}

func (o X12EnvelopeSettingsOutput) UseControlStandardsIdAsRepetitionCharacter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *bool { return v.UseControlStandardsIdAsRepetitionCharacter }).(pulumi.BoolPtrOutput)
}

type X12EnvelopeSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12EnvelopeSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12EnvelopeSettings)(nil)).Elem()
}

func (o X12EnvelopeSettingsPtrOutput) ToX12EnvelopeSettingsPtrOutput() X12EnvelopeSettingsPtrOutput {
	return o
}

func (o X12EnvelopeSettingsPtrOutput) ToX12EnvelopeSettingsPtrOutputWithContext(ctx context.Context) X12EnvelopeSettingsPtrOutput {
	return o
}

func (o X12EnvelopeSettingsPtrOutput) Elem() X12EnvelopeSettingsOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) X12EnvelopeSettings {
		if v != nil {
			return *v
		}
		var ret X12EnvelopeSettings
		return ret
	}).(X12EnvelopeSettingsOutput)
}

func (o X12EnvelopeSettingsPtrOutput) ControlStandardsId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return v.ControlStandardsId
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) ControlVersionNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ControlVersionNumber
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDefaultGroupHeaders
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalGroupId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupHeaderAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupHeaderAgencyCode
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupHeaderDateFormat() X12DateFormatPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *X12DateFormat {
		if v == nil {
			return nil
		}
		return v.GroupHeaderDateFormat
	}).(X12DateFormatPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupHeaderTimeFormat() X12TimeFormatPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *X12TimeFormat {
		if v == nil {
			return nil
		}
		return v.GroupHeaderTimeFormat
	}).(X12TimeFormatPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupHeaderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupHeaderVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) InterchangeControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) InterchangeControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.OverwriteExistingTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.SenderApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) TransactionSetControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) TransactionSetControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) UsageIndicator() UsageIndicatorPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *UsageIndicator {
		if v == nil {
			return nil
		}
		return v.UsageIndicator
	}).(UsageIndicatorPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) UseControlStandardsIdAsRepetitionCharacter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return v.UseControlStandardsIdAsRepetitionCharacter
	}).(pulumi.BoolPtrOutput)
}

type X12EnvelopeSettingsResponse struct {
	ControlStandardsId                           *int    `pulumi:"controlStandardsId"`
	ControlVersionNumber                         *string `pulumi:"controlVersionNumber"`
	EnableDefaultGroupHeaders                    *bool   `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            *string `pulumi:"functionalGroupId"`
	GroupControlNumberLowerBound                 *int    `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberUpperBound                 *int    `pulumi:"groupControlNumberUpperBound"`
	GroupHeaderAgencyCode                        *string `pulumi:"groupHeaderAgencyCode"`
	GroupHeaderDateFormat                        *string `pulumi:"groupHeaderDateFormat"`
	GroupHeaderTimeFormat                        *string `pulumi:"groupHeaderTimeFormat"`
	GroupHeaderVersion                           *string `pulumi:"groupHeaderVersion"`
	InterchangeControlNumberLowerBound           *int    `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberUpperBound           *int    `pulumi:"interchangeControlNumberUpperBound"`
	OverwriteExistingTransactionSetControlNumber *bool   `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ReceiverApplicationId                        *string `pulumi:"receiverApplicationId"`
	RolloverGroupControlNumber                   *bool   `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             *bool   `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          *bool   `pulumi:"rolloverTransactionSetControlNumber"`
	SenderApplicationId                          *string `pulumi:"senderApplicationId"`
	TransactionSetControlNumberLowerBound        *int    `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            *string `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            *string `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        *int    `pulumi:"transactionSetControlNumberUpperBound"`
	UsageIndicator                               *string `pulumi:"usageIndicator"`
	UseControlStandardsIdAsRepetitionCharacter   *bool   `pulumi:"useControlStandardsIdAsRepetitionCharacter"`
}

type X12EnvelopeSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12EnvelopeSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeSettingsResponse)(nil)).Elem()
}

func (o X12EnvelopeSettingsResponseOutput) ToX12EnvelopeSettingsResponseOutput() X12EnvelopeSettingsResponseOutput {
	return o
}

func (o X12EnvelopeSettingsResponseOutput) ToX12EnvelopeSettingsResponseOutputWithContext(ctx context.Context) X12EnvelopeSettingsResponseOutput {
	return o
}

func (o X12EnvelopeSettingsResponseOutput) ControlStandardsId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *int { return v.ControlStandardsId }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) ControlVersionNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.ControlVersionNumber }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *bool { return v.EnableDefaultGroupHeaders }).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *int { return v.GroupControlNumberLowerBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *int { return v.GroupControlNumberUpperBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupHeaderAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.GroupHeaderAgencyCode }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupHeaderDateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.GroupHeaderDateFormat }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupHeaderTimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.GroupHeaderTimeFormat }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupHeaderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.GroupHeaderVersion }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) InterchangeControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *int { return v.InterchangeControlNumberLowerBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) InterchangeControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *int { return v.InterchangeControlNumberUpperBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *bool { return v.OverwriteExistingTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.ReceiverApplicationId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *bool { return v.RolloverGroupControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *bool { return v.RolloverInterchangeControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *bool { return v.RolloverTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) TransactionSetControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *int { return v.TransactionSetControlNumberLowerBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.TransactionSetControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.TransactionSetControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) TransactionSetControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *int { return v.TransactionSetControlNumberUpperBound }).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) UsageIndicator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.UsageIndicator }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) UseControlStandardsIdAsRepetitionCharacter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *bool { return v.UseControlStandardsIdAsRepetitionCharacter }).(pulumi.BoolPtrOutput)
}

type X12EnvelopeSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12EnvelopeSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12EnvelopeSettingsResponse)(nil)).Elem()
}

func (o X12EnvelopeSettingsResponsePtrOutput) ToX12EnvelopeSettingsResponsePtrOutput() X12EnvelopeSettingsResponsePtrOutput {
	return o
}

func (o X12EnvelopeSettingsResponsePtrOutput) ToX12EnvelopeSettingsResponsePtrOutputWithContext(ctx context.Context) X12EnvelopeSettingsResponsePtrOutput {
	return o
}

func (o X12EnvelopeSettingsResponsePtrOutput) Elem() X12EnvelopeSettingsResponseOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) X12EnvelopeSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12EnvelopeSettingsResponse
		return ret
	}).(X12EnvelopeSettingsResponseOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) ControlStandardsId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.ControlStandardsId
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) ControlVersionNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ControlVersionNumber
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDefaultGroupHeaders
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalGroupId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupHeaderAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupHeaderAgencyCode
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupHeaderDateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupHeaderDateFormat
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupHeaderTimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupHeaderTimeFormat
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupHeaderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupHeaderVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) InterchangeControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) InterchangeControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.OverwriteExistingTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RolloverTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SenderApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) UsageIndicator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.UsageIndicator
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) UseControlStandardsIdAsRepetitionCharacter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseControlStandardsIdAsRepetitionCharacter
	}).(pulumi.BoolPtrOutput)
}

type X12FramingSettings struct {
	CharacterSet               *X12CharacterSet         `pulumi:"characterSet"`
	ComponentSeparator         *int                     `pulumi:"componentSeparator"`
	DataElementSeparator       *int                     `pulumi:"dataElementSeparator"`
	ReplaceCharacter           *int                     `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload *bool                    `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          *int                     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    *SegmentTerminatorSuffix `pulumi:"segmentTerminatorSuffix"`
}





type X12FramingSettingsInput interface {
	pulumi.Input

	ToX12FramingSettingsOutput() X12FramingSettingsOutput
	ToX12FramingSettingsOutputWithContext(context.Context) X12FramingSettingsOutput
}

type X12FramingSettingsArgs struct {
	CharacterSet               X12CharacterSetPtrInput         `pulumi:"characterSet"`
	ComponentSeparator         pulumi.IntPtrInput              `pulumi:"componentSeparator"`
	DataElementSeparator       pulumi.IntPtrInput              `pulumi:"dataElementSeparator"`
	ReplaceCharacter           pulumi.IntPtrInput              `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload pulumi.BoolPtrInput             `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          pulumi.IntPtrInput              `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    SegmentTerminatorSuffixPtrInput `pulumi:"segmentTerminatorSuffix"`
}

func (X12FramingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12FramingSettings)(nil)).Elem()
}

func (i X12FramingSettingsArgs) ToX12FramingSettingsOutput() X12FramingSettingsOutput {
	return i.ToX12FramingSettingsOutputWithContext(context.Background())
}

func (i X12FramingSettingsArgs) ToX12FramingSettingsOutputWithContext(ctx context.Context) X12FramingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12FramingSettingsOutput)
}

func (i X12FramingSettingsArgs) ToX12FramingSettingsPtrOutput() X12FramingSettingsPtrOutput {
	return i.ToX12FramingSettingsPtrOutputWithContext(context.Background())
}

func (i X12FramingSettingsArgs) ToX12FramingSettingsPtrOutputWithContext(ctx context.Context) X12FramingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12FramingSettingsOutput).ToX12FramingSettingsPtrOutputWithContext(ctx)
}









type X12FramingSettingsPtrInput interface {
	pulumi.Input

	ToX12FramingSettingsPtrOutput() X12FramingSettingsPtrOutput
	ToX12FramingSettingsPtrOutputWithContext(context.Context) X12FramingSettingsPtrOutput
}

type x12framingSettingsPtrType X12FramingSettingsArgs

func X12FramingSettingsPtr(v *X12FramingSettingsArgs) X12FramingSettingsPtrInput {
	return (*x12framingSettingsPtrType)(v)
}

func (*x12framingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12FramingSettings)(nil)).Elem()
}

func (i *x12framingSettingsPtrType) ToX12FramingSettingsPtrOutput() X12FramingSettingsPtrOutput {
	return i.ToX12FramingSettingsPtrOutputWithContext(context.Background())
}

func (i *x12framingSettingsPtrType) ToX12FramingSettingsPtrOutputWithContext(ctx context.Context) X12FramingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12FramingSettingsPtrOutput)
}

type X12FramingSettingsOutput struct{ *pulumi.OutputState }

func (X12FramingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12FramingSettings)(nil)).Elem()
}

func (o X12FramingSettingsOutput) ToX12FramingSettingsOutput() X12FramingSettingsOutput {
	return o
}

func (o X12FramingSettingsOutput) ToX12FramingSettingsOutputWithContext(ctx context.Context) X12FramingSettingsOutput {
	return o
}

func (o X12FramingSettingsOutput) ToX12FramingSettingsPtrOutput() X12FramingSettingsPtrOutput {
	return o.ToX12FramingSettingsPtrOutputWithContext(context.Background())
}

func (o X12FramingSettingsOutput) ToX12FramingSettingsPtrOutputWithContext(ctx context.Context) X12FramingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12FramingSettings) *X12FramingSettings {
		return &v
	}).(X12FramingSettingsPtrOutput)
}

func (o X12FramingSettingsOutput) CharacterSet() X12CharacterSetPtrOutput {
	return o.ApplyT(func(v X12FramingSettings) *X12CharacterSet { return v.CharacterSet }).(X12CharacterSetPtrOutput)
}

func (o X12FramingSettingsOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12FramingSettings) *int { return v.ComponentSeparator }).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12FramingSettings) *int { return v.DataElementSeparator }).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsOutput) ReplaceCharacter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12FramingSettings) *int { return v.ReplaceCharacter }).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsOutput) ReplaceSeparatorsInPayload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12FramingSettings) *bool { return v.ReplaceSeparatorsInPayload }).(pulumi.BoolPtrOutput)
}

func (o X12FramingSettingsOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12FramingSettings) *int { return v.SegmentTerminator }).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixPtrOutput {
	return o.ApplyT(func(v X12FramingSettings) *SegmentTerminatorSuffix { return v.SegmentTerminatorSuffix }).(SegmentTerminatorSuffixPtrOutput)
}

type X12FramingSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12FramingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12FramingSettings)(nil)).Elem()
}

func (o X12FramingSettingsPtrOutput) ToX12FramingSettingsPtrOutput() X12FramingSettingsPtrOutput {
	return o
}

func (o X12FramingSettingsPtrOutput) ToX12FramingSettingsPtrOutputWithContext(ctx context.Context) X12FramingSettingsPtrOutput {
	return o
}

func (o X12FramingSettingsPtrOutput) Elem() X12FramingSettingsOutput {
	return o.ApplyT(func(v *X12FramingSettings) X12FramingSettings {
		if v != nil {
			return *v
		}
		var ret X12FramingSettings
		return ret
	}).(X12FramingSettingsOutput)
}

func (o X12FramingSettingsPtrOutput) CharacterSet() X12CharacterSetPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *X12CharacterSet {
		if v == nil {
			return nil
		}
		return v.CharacterSet
	}).(X12CharacterSetPtrOutput)
}

func (o X12FramingSettingsPtrOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *int {
		if v == nil {
			return nil
		}
		return v.ComponentSeparator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsPtrOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *int {
		if v == nil {
			return nil
		}
		return v.DataElementSeparator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsPtrOutput) ReplaceCharacter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *int {
		if v == nil {
			return nil
		}
		return v.ReplaceCharacter
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsPtrOutput) ReplaceSeparatorsInPayload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ReplaceSeparatorsInPayload
	}).(pulumi.BoolPtrOutput)
}

func (o X12FramingSettingsPtrOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *int {
		if v == nil {
			return nil
		}
		return v.SegmentTerminator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsPtrOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *SegmentTerminatorSuffix {
		if v == nil {
			return nil
		}
		return v.SegmentTerminatorSuffix
	}).(SegmentTerminatorSuffixPtrOutput)
}

type X12FramingSettingsResponse struct {
	CharacterSet               *string `pulumi:"characterSet"`
	ComponentSeparator         *int    `pulumi:"componentSeparator"`
	DataElementSeparator       *int    `pulumi:"dataElementSeparator"`
	ReplaceCharacter           *int    `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload *bool   `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          *int    `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    *string `pulumi:"segmentTerminatorSuffix"`
}

type X12FramingSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12FramingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12FramingSettingsResponse)(nil)).Elem()
}

func (o X12FramingSettingsResponseOutput) ToX12FramingSettingsResponseOutput() X12FramingSettingsResponseOutput {
	return o
}

func (o X12FramingSettingsResponseOutput) ToX12FramingSettingsResponseOutputWithContext(ctx context.Context) X12FramingSettingsResponseOutput {
	return o
}

func (o X12FramingSettingsResponseOutput) CharacterSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) *string { return v.CharacterSet }).(pulumi.StringPtrOutput)
}

func (o X12FramingSettingsResponseOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) *int { return v.ComponentSeparator }).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponseOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) *int { return v.DataElementSeparator }).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponseOutput) ReplaceCharacter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) *int { return v.ReplaceCharacter }).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponseOutput) ReplaceSeparatorsInPayload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) *bool { return v.ReplaceSeparatorsInPayload }).(pulumi.BoolPtrOutput)
}

func (o X12FramingSettingsResponseOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) *int { return v.SegmentTerminator }).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponseOutput) SegmentTerminatorSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) *string { return v.SegmentTerminatorSuffix }).(pulumi.StringPtrOutput)
}

type X12FramingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12FramingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12FramingSettingsResponse)(nil)).Elem()
}

func (o X12FramingSettingsResponsePtrOutput) ToX12FramingSettingsResponsePtrOutput() X12FramingSettingsResponsePtrOutput {
	return o
}

func (o X12FramingSettingsResponsePtrOutput) ToX12FramingSettingsResponsePtrOutputWithContext(ctx context.Context) X12FramingSettingsResponsePtrOutput {
	return o
}

func (o X12FramingSettingsResponsePtrOutput) Elem() X12FramingSettingsResponseOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) X12FramingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12FramingSettingsResponse
		return ret
	}).(X12FramingSettingsResponseOutput)
}

func (o X12FramingSettingsResponsePtrOutput) CharacterSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.CharacterSet
	}).(pulumi.StringPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.ComponentSeparator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.DataElementSeparator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) ReplaceCharacter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.ReplaceCharacter
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) ReplaceSeparatorsInPayload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ReplaceSeparatorsInPayload
	}).(pulumi.BoolPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.SegmentTerminator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) SegmentTerminatorSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SegmentTerminatorSuffix
	}).(pulumi.StringPtrOutput)
}

type X12MessageFilter struct {
	MessageFilterType *MessageFilterType `pulumi:"messageFilterType"`
}





type X12MessageFilterInput interface {
	pulumi.Input

	ToX12MessageFilterOutput() X12MessageFilterOutput
	ToX12MessageFilterOutputWithContext(context.Context) X12MessageFilterOutput
}

type X12MessageFilterArgs struct {
	MessageFilterType MessageFilterTypePtrInput `pulumi:"messageFilterType"`
}

func (X12MessageFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageFilter)(nil)).Elem()
}

func (i X12MessageFilterArgs) ToX12MessageFilterOutput() X12MessageFilterOutput {
	return i.ToX12MessageFilterOutputWithContext(context.Background())
}

func (i X12MessageFilterArgs) ToX12MessageFilterOutputWithContext(ctx context.Context) X12MessageFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12MessageFilterOutput)
}

func (i X12MessageFilterArgs) ToX12MessageFilterPtrOutput() X12MessageFilterPtrOutput {
	return i.ToX12MessageFilterPtrOutputWithContext(context.Background())
}

func (i X12MessageFilterArgs) ToX12MessageFilterPtrOutputWithContext(ctx context.Context) X12MessageFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12MessageFilterOutput).ToX12MessageFilterPtrOutputWithContext(ctx)
}









type X12MessageFilterPtrInput interface {
	pulumi.Input

	ToX12MessageFilterPtrOutput() X12MessageFilterPtrOutput
	ToX12MessageFilterPtrOutputWithContext(context.Context) X12MessageFilterPtrOutput
}

type x12messageFilterPtrType X12MessageFilterArgs

func X12MessageFilterPtr(v *X12MessageFilterArgs) X12MessageFilterPtrInput {
	return (*x12messageFilterPtrType)(v)
}

func (*x12messageFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12MessageFilter)(nil)).Elem()
}

func (i *x12messageFilterPtrType) ToX12MessageFilterPtrOutput() X12MessageFilterPtrOutput {
	return i.ToX12MessageFilterPtrOutputWithContext(context.Background())
}

func (i *x12messageFilterPtrType) ToX12MessageFilterPtrOutputWithContext(ctx context.Context) X12MessageFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12MessageFilterPtrOutput)
}

type X12MessageFilterOutput struct{ *pulumi.OutputState }

func (X12MessageFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageFilter)(nil)).Elem()
}

func (o X12MessageFilterOutput) ToX12MessageFilterOutput() X12MessageFilterOutput {
	return o
}

func (o X12MessageFilterOutput) ToX12MessageFilterOutputWithContext(ctx context.Context) X12MessageFilterOutput {
	return o
}

func (o X12MessageFilterOutput) ToX12MessageFilterPtrOutput() X12MessageFilterPtrOutput {
	return o.ToX12MessageFilterPtrOutputWithContext(context.Background())
}

func (o X12MessageFilterOutput) ToX12MessageFilterPtrOutputWithContext(ctx context.Context) X12MessageFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12MessageFilter) *X12MessageFilter {
		return &v
	}).(X12MessageFilterPtrOutput)
}

func (o X12MessageFilterOutput) MessageFilterType() MessageFilterTypePtrOutput {
	return o.ApplyT(func(v X12MessageFilter) *MessageFilterType { return v.MessageFilterType }).(MessageFilterTypePtrOutput)
}

type X12MessageFilterPtrOutput struct{ *pulumi.OutputState }

func (X12MessageFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12MessageFilter)(nil)).Elem()
}

func (o X12MessageFilterPtrOutput) ToX12MessageFilterPtrOutput() X12MessageFilterPtrOutput {
	return o
}

func (o X12MessageFilterPtrOutput) ToX12MessageFilterPtrOutputWithContext(ctx context.Context) X12MessageFilterPtrOutput {
	return o
}

func (o X12MessageFilterPtrOutput) Elem() X12MessageFilterOutput {
	return o.ApplyT(func(v *X12MessageFilter) X12MessageFilter {
		if v != nil {
			return *v
		}
		var ret X12MessageFilter
		return ret
	}).(X12MessageFilterOutput)
}

func (o X12MessageFilterPtrOutput) MessageFilterType() MessageFilterTypePtrOutput {
	return o.ApplyT(func(v *X12MessageFilter) *MessageFilterType {
		if v == nil {
			return nil
		}
		return v.MessageFilterType
	}).(MessageFilterTypePtrOutput)
}

type X12MessageFilterResponse struct {
	MessageFilterType *string `pulumi:"messageFilterType"`
}

type X12MessageFilterResponseOutput struct{ *pulumi.OutputState }

func (X12MessageFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageFilterResponse)(nil)).Elem()
}

func (o X12MessageFilterResponseOutput) ToX12MessageFilterResponseOutput() X12MessageFilterResponseOutput {
	return o
}

func (o X12MessageFilterResponseOutput) ToX12MessageFilterResponseOutputWithContext(ctx context.Context) X12MessageFilterResponseOutput {
	return o
}

func (o X12MessageFilterResponseOutput) MessageFilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12MessageFilterResponse) *string { return v.MessageFilterType }).(pulumi.StringPtrOutput)
}

type X12MessageFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (X12MessageFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12MessageFilterResponse)(nil)).Elem()
}

func (o X12MessageFilterResponsePtrOutput) ToX12MessageFilterResponsePtrOutput() X12MessageFilterResponsePtrOutput {
	return o
}

func (o X12MessageFilterResponsePtrOutput) ToX12MessageFilterResponsePtrOutputWithContext(ctx context.Context) X12MessageFilterResponsePtrOutput {
	return o
}

func (o X12MessageFilterResponsePtrOutput) Elem() X12MessageFilterResponseOutput {
	return o.ApplyT(func(v *X12MessageFilterResponse) X12MessageFilterResponse {
		if v != nil {
			return *v
		}
		var ret X12MessageFilterResponse
		return ret
	}).(X12MessageFilterResponseOutput)
}

func (o X12MessageFilterResponsePtrOutput) MessageFilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12MessageFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageFilterType
	}).(pulumi.StringPtrOutput)
}

type X12MessageIdentifier struct {
	MessageId *string `pulumi:"messageId"`
}





type X12MessageIdentifierInput interface {
	pulumi.Input

	ToX12MessageIdentifierOutput() X12MessageIdentifierOutput
	ToX12MessageIdentifierOutputWithContext(context.Context) X12MessageIdentifierOutput
}

type X12MessageIdentifierArgs struct {
	MessageId pulumi.StringPtrInput `pulumi:"messageId"`
}

func (X12MessageIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageIdentifier)(nil)).Elem()
}

func (i X12MessageIdentifierArgs) ToX12MessageIdentifierOutput() X12MessageIdentifierOutput {
	return i.ToX12MessageIdentifierOutputWithContext(context.Background())
}

func (i X12MessageIdentifierArgs) ToX12MessageIdentifierOutputWithContext(ctx context.Context) X12MessageIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12MessageIdentifierOutput)
}





type X12MessageIdentifierArrayInput interface {
	pulumi.Input

	ToX12MessageIdentifierArrayOutput() X12MessageIdentifierArrayOutput
	ToX12MessageIdentifierArrayOutputWithContext(context.Context) X12MessageIdentifierArrayOutput
}

type X12MessageIdentifierArray []X12MessageIdentifierInput

func (X12MessageIdentifierArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12MessageIdentifier)(nil)).Elem()
}

func (i X12MessageIdentifierArray) ToX12MessageIdentifierArrayOutput() X12MessageIdentifierArrayOutput {
	return i.ToX12MessageIdentifierArrayOutputWithContext(context.Background())
}

func (i X12MessageIdentifierArray) ToX12MessageIdentifierArrayOutputWithContext(ctx context.Context) X12MessageIdentifierArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12MessageIdentifierArrayOutput)
}

type X12MessageIdentifierOutput struct{ *pulumi.OutputState }

func (X12MessageIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageIdentifier)(nil)).Elem()
}

func (o X12MessageIdentifierOutput) ToX12MessageIdentifierOutput() X12MessageIdentifierOutput {
	return o
}

func (o X12MessageIdentifierOutput) ToX12MessageIdentifierOutputWithContext(ctx context.Context) X12MessageIdentifierOutput {
	return o
}

func (o X12MessageIdentifierOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12MessageIdentifier) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

type X12MessageIdentifierArrayOutput struct{ *pulumi.OutputState }

func (X12MessageIdentifierArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12MessageIdentifier)(nil)).Elem()
}

func (o X12MessageIdentifierArrayOutput) ToX12MessageIdentifierArrayOutput() X12MessageIdentifierArrayOutput {
	return o
}

func (o X12MessageIdentifierArrayOutput) ToX12MessageIdentifierArrayOutputWithContext(ctx context.Context) X12MessageIdentifierArrayOutput {
	return o
}

func (o X12MessageIdentifierArrayOutput) Index(i pulumi.IntInput) X12MessageIdentifierOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12MessageIdentifier {
		return vs[0].([]X12MessageIdentifier)[vs[1].(int)]
	}).(X12MessageIdentifierOutput)
}

type X12MessageIdentifierResponse struct {
	MessageId *string `pulumi:"messageId"`
}

type X12MessageIdentifierResponseOutput struct{ *pulumi.OutputState }

func (X12MessageIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageIdentifierResponse)(nil)).Elem()
}

func (o X12MessageIdentifierResponseOutput) ToX12MessageIdentifierResponseOutput() X12MessageIdentifierResponseOutput {
	return o
}

func (o X12MessageIdentifierResponseOutput) ToX12MessageIdentifierResponseOutputWithContext(ctx context.Context) X12MessageIdentifierResponseOutput {
	return o
}

func (o X12MessageIdentifierResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12MessageIdentifierResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

type X12MessageIdentifierResponseArrayOutput struct{ *pulumi.OutputState }

func (X12MessageIdentifierResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12MessageIdentifierResponse)(nil)).Elem()
}

func (o X12MessageIdentifierResponseArrayOutput) ToX12MessageIdentifierResponseArrayOutput() X12MessageIdentifierResponseArrayOutput {
	return o
}

func (o X12MessageIdentifierResponseArrayOutput) ToX12MessageIdentifierResponseArrayOutputWithContext(ctx context.Context) X12MessageIdentifierResponseArrayOutput {
	return o
}

func (o X12MessageIdentifierResponseArrayOutput) Index(i pulumi.IntInput) X12MessageIdentifierResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12MessageIdentifierResponse {
		return vs[0].([]X12MessageIdentifierResponse)[vs[1].(int)]
	}).(X12MessageIdentifierResponseOutput)
}

type X12OneWayAgreement struct {
	ProtocolSettings         *X12ProtocolSettings `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity *BusinessIdentity    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   *BusinessIdentity    `pulumi:"senderBusinessIdentity"`
}





type X12OneWayAgreementInput interface {
	pulumi.Input

	ToX12OneWayAgreementOutput() X12OneWayAgreementOutput
	ToX12OneWayAgreementOutputWithContext(context.Context) X12OneWayAgreementOutput
}

type X12OneWayAgreementArgs struct {
	ProtocolSettings         X12ProtocolSettingsPtrInput `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentityPtrInput    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentityPtrInput    `pulumi:"senderBusinessIdentity"`
}

func (X12OneWayAgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12OneWayAgreement)(nil)).Elem()
}

func (i X12OneWayAgreementArgs) ToX12OneWayAgreementOutput() X12OneWayAgreementOutput {
	return i.ToX12OneWayAgreementOutputWithContext(context.Background())
}

func (i X12OneWayAgreementArgs) ToX12OneWayAgreementOutputWithContext(ctx context.Context) X12OneWayAgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12OneWayAgreementOutput)
}

func (i X12OneWayAgreementArgs) ToX12OneWayAgreementPtrOutput() X12OneWayAgreementPtrOutput {
	return i.ToX12OneWayAgreementPtrOutputWithContext(context.Background())
}

func (i X12OneWayAgreementArgs) ToX12OneWayAgreementPtrOutputWithContext(ctx context.Context) X12OneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12OneWayAgreementOutput).ToX12OneWayAgreementPtrOutputWithContext(ctx)
}









type X12OneWayAgreementPtrInput interface {
	pulumi.Input

	ToX12OneWayAgreementPtrOutput() X12OneWayAgreementPtrOutput
	ToX12OneWayAgreementPtrOutputWithContext(context.Context) X12OneWayAgreementPtrOutput
}

type x12oneWayAgreementPtrType X12OneWayAgreementArgs

func X12OneWayAgreementPtr(v *X12OneWayAgreementArgs) X12OneWayAgreementPtrInput {
	return (*x12oneWayAgreementPtrType)(v)
}

func (*x12oneWayAgreementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12OneWayAgreement)(nil)).Elem()
}

func (i *x12oneWayAgreementPtrType) ToX12OneWayAgreementPtrOutput() X12OneWayAgreementPtrOutput {
	return i.ToX12OneWayAgreementPtrOutputWithContext(context.Background())
}

func (i *x12oneWayAgreementPtrType) ToX12OneWayAgreementPtrOutputWithContext(ctx context.Context) X12OneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12OneWayAgreementPtrOutput)
}

type X12OneWayAgreementOutput struct{ *pulumi.OutputState }

func (X12OneWayAgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12OneWayAgreement)(nil)).Elem()
}

func (o X12OneWayAgreementOutput) ToX12OneWayAgreementOutput() X12OneWayAgreementOutput {
	return o
}

func (o X12OneWayAgreementOutput) ToX12OneWayAgreementOutputWithContext(ctx context.Context) X12OneWayAgreementOutput {
	return o
}

func (o X12OneWayAgreementOutput) ToX12OneWayAgreementPtrOutput() X12OneWayAgreementPtrOutput {
	return o.ToX12OneWayAgreementPtrOutputWithContext(context.Background())
}

func (o X12OneWayAgreementOutput) ToX12OneWayAgreementPtrOutputWithContext(ctx context.Context) X12OneWayAgreementPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12OneWayAgreement) *X12OneWayAgreement {
		return &v
	}).(X12OneWayAgreementPtrOutput)
}

func (o X12OneWayAgreementOutput) ProtocolSettings() X12ProtocolSettingsPtrOutput {
	return o.ApplyT(func(v X12OneWayAgreement) *X12ProtocolSettings { return v.ProtocolSettings }).(X12ProtocolSettingsPtrOutput)
}

func (o X12OneWayAgreementOutput) ReceiverBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v X12OneWayAgreement) *BusinessIdentity { return v.ReceiverBusinessIdentity }).(BusinessIdentityPtrOutput)
}

func (o X12OneWayAgreementOutput) SenderBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v X12OneWayAgreement) *BusinessIdentity { return v.SenderBusinessIdentity }).(BusinessIdentityPtrOutput)
}

type X12OneWayAgreementPtrOutput struct{ *pulumi.OutputState }

func (X12OneWayAgreementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12OneWayAgreement)(nil)).Elem()
}

func (o X12OneWayAgreementPtrOutput) ToX12OneWayAgreementPtrOutput() X12OneWayAgreementPtrOutput {
	return o
}

func (o X12OneWayAgreementPtrOutput) ToX12OneWayAgreementPtrOutputWithContext(ctx context.Context) X12OneWayAgreementPtrOutput {
	return o
}

func (o X12OneWayAgreementPtrOutput) Elem() X12OneWayAgreementOutput {
	return o.ApplyT(func(v *X12OneWayAgreement) X12OneWayAgreement {
		if v != nil {
			return *v
		}
		var ret X12OneWayAgreement
		return ret
	}).(X12OneWayAgreementOutput)
}

func (o X12OneWayAgreementPtrOutput) ProtocolSettings() X12ProtocolSettingsPtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreement) *X12ProtocolSettings {
		if v == nil {
			return nil
		}
		return v.ProtocolSettings
	}).(X12ProtocolSettingsPtrOutput)
}

func (o X12OneWayAgreementPtrOutput) ReceiverBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return v.ReceiverBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

func (o X12OneWayAgreementPtrOutput) SenderBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return v.SenderBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

type X12OneWayAgreementResponse struct {
	ProtocolSettings         *X12ProtocolSettingsResponse `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity *BusinessIdentityResponse    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   *BusinessIdentityResponse    `pulumi:"senderBusinessIdentity"`
}

type X12OneWayAgreementResponseOutput struct{ *pulumi.OutputState }

func (X12OneWayAgreementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12OneWayAgreementResponse)(nil)).Elem()
}

func (o X12OneWayAgreementResponseOutput) ToX12OneWayAgreementResponseOutput() X12OneWayAgreementResponseOutput {
	return o
}

func (o X12OneWayAgreementResponseOutput) ToX12OneWayAgreementResponseOutputWithContext(ctx context.Context) X12OneWayAgreementResponseOutput {
	return o
}

func (o X12OneWayAgreementResponseOutput) ProtocolSettings() X12ProtocolSettingsResponsePtrOutput {
	return o.ApplyT(func(v X12OneWayAgreementResponse) *X12ProtocolSettingsResponse { return v.ProtocolSettings }).(X12ProtocolSettingsResponsePtrOutput)
}

func (o X12OneWayAgreementResponseOutput) ReceiverBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v X12OneWayAgreementResponse) *BusinessIdentityResponse { return v.ReceiverBusinessIdentity }).(BusinessIdentityResponsePtrOutput)
}

func (o X12OneWayAgreementResponseOutput) SenderBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v X12OneWayAgreementResponse) *BusinessIdentityResponse { return v.SenderBusinessIdentity }).(BusinessIdentityResponsePtrOutput)
}

type X12OneWayAgreementResponsePtrOutput struct{ *pulumi.OutputState }

func (X12OneWayAgreementResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12OneWayAgreementResponse)(nil)).Elem()
}

func (o X12OneWayAgreementResponsePtrOutput) ToX12OneWayAgreementResponsePtrOutput() X12OneWayAgreementResponsePtrOutput {
	return o
}

func (o X12OneWayAgreementResponsePtrOutput) ToX12OneWayAgreementResponsePtrOutputWithContext(ctx context.Context) X12OneWayAgreementResponsePtrOutput {
	return o
}

func (o X12OneWayAgreementResponsePtrOutput) Elem() X12OneWayAgreementResponseOutput {
	return o.ApplyT(func(v *X12OneWayAgreementResponse) X12OneWayAgreementResponse {
		if v != nil {
			return *v
		}
		var ret X12OneWayAgreementResponse
		return ret
	}).(X12OneWayAgreementResponseOutput)
}

func (o X12OneWayAgreementResponsePtrOutput) ProtocolSettings() X12ProtocolSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreementResponse) *X12ProtocolSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ProtocolSettings
	}).(X12ProtocolSettingsResponsePtrOutput)
}

func (o X12OneWayAgreementResponsePtrOutput) ReceiverBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return v.ReceiverBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

func (o X12OneWayAgreementResponsePtrOutput) SenderBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return v.SenderBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

type X12ProcessingSettings struct {
	ConvertImpliedDecimal                   *bool `pulumi:"convertImpliedDecimal"`
	CreateEmptyXmlTagsForTrailingSeparators *bool `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        *bool `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     *bool `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               *bool `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                *bool `pulumi:"useDotAsDecimalSeparator"`
}





type X12ProcessingSettingsInput interface {
	pulumi.Input

	ToX12ProcessingSettingsOutput() X12ProcessingSettingsOutput
	ToX12ProcessingSettingsOutputWithContext(context.Context) X12ProcessingSettingsOutput
}

type X12ProcessingSettingsArgs struct {
	ConvertImpliedDecimal                   pulumi.BoolPtrInput `pulumi:"convertImpliedDecimal"`
	CreateEmptyXmlTagsForTrailingSeparators pulumi.BoolPtrInput `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        pulumi.BoolPtrInput `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     pulumi.BoolPtrInput `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               pulumi.BoolPtrInput `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                pulumi.BoolPtrInput `pulumi:"useDotAsDecimalSeparator"`
}

func (X12ProcessingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProcessingSettings)(nil)).Elem()
}

func (i X12ProcessingSettingsArgs) ToX12ProcessingSettingsOutput() X12ProcessingSettingsOutput {
	return i.ToX12ProcessingSettingsOutputWithContext(context.Background())
}

func (i X12ProcessingSettingsArgs) ToX12ProcessingSettingsOutputWithContext(ctx context.Context) X12ProcessingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProcessingSettingsOutput)
}

func (i X12ProcessingSettingsArgs) ToX12ProcessingSettingsPtrOutput() X12ProcessingSettingsPtrOutput {
	return i.ToX12ProcessingSettingsPtrOutputWithContext(context.Background())
}

func (i X12ProcessingSettingsArgs) ToX12ProcessingSettingsPtrOutputWithContext(ctx context.Context) X12ProcessingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProcessingSettingsOutput).ToX12ProcessingSettingsPtrOutputWithContext(ctx)
}









type X12ProcessingSettingsPtrInput interface {
	pulumi.Input

	ToX12ProcessingSettingsPtrOutput() X12ProcessingSettingsPtrOutput
	ToX12ProcessingSettingsPtrOutputWithContext(context.Context) X12ProcessingSettingsPtrOutput
}

type x12processingSettingsPtrType X12ProcessingSettingsArgs

func X12ProcessingSettingsPtr(v *X12ProcessingSettingsArgs) X12ProcessingSettingsPtrInput {
	return (*x12processingSettingsPtrType)(v)
}

func (*x12processingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProcessingSettings)(nil)).Elem()
}

func (i *x12processingSettingsPtrType) ToX12ProcessingSettingsPtrOutput() X12ProcessingSettingsPtrOutput {
	return i.ToX12ProcessingSettingsPtrOutputWithContext(context.Background())
}

func (i *x12processingSettingsPtrType) ToX12ProcessingSettingsPtrOutputWithContext(ctx context.Context) X12ProcessingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProcessingSettingsPtrOutput)
}

type X12ProcessingSettingsOutput struct{ *pulumi.OutputState }

func (X12ProcessingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProcessingSettings)(nil)).Elem()
}

func (o X12ProcessingSettingsOutput) ToX12ProcessingSettingsOutput() X12ProcessingSettingsOutput {
	return o
}

func (o X12ProcessingSettingsOutput) ToX12ProcessingSettingsOutputWithContext(ctx context.Context) X12ProcessingSettingsOutput {
	return o
}

func (o X12ProcessingSettingsOutput) ToX12ProcessingSettingsPtrOutput() X12ProcessingSettingsPtrOutput {
	return o.ToX12ProcessingSettingsPtrOutputWithContext(context.Background())
}

func (o X12ProcessingSettingsOutput) ToX12ProcessingSettingsPtrOutputWithContext(ctx context.Context) X12ProcessingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12ProcessingSettings) *X12ProcessingSettings {
		return &v
	}).(X12ProcessingSettingsPtrOutput)
}

func (o X12ProcessingSettingsOutput) ConvertImpliedDecimal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettings) *bool { return v.ConvertImpliedDecimal }).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettings) *bool { return v.CreateEmptyXmlTagsForTrailingSeparators }).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettings) *bool { return v.MaskSecurityInfo }).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettings) *bool { return v.PreserveInterchange }).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettings) *bool { return v.SuspendInterchangeOnError }).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettings) *bool { return v.UseDotAsDecimalSeparator }).(pulumi.BoolPtrOutput)
}

type X12ProcessingSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12ProcessingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProcessingSettings)(nil)).Elem()
}

func (o X12ProcessingSettingsPtrOutput) ToX12ProcessingSettingsPtrOutput() X12ProcessingSettingsPtrOutput {
	return o
}

func (o X12ProcessingSettingsPtrOutput) ToX12ProcessingSettingsPtrOutputWithContext(ctx context.Context) X12ProcessingSettingsPtrOutput {
	return o
}

func (o X12ProcessingSettingsPtrOutput) Elem() X12ProcessingSettingsOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) X12ProcessingSettings {
		if v != nil {
			return *v
		}
		var ret X12ProcessingSettings
		return ret
	}).(X12ProcessingSettingsOutput)
}

func (o X12ProcessingSettingsPtrOutput) ConvertImpliedDecimal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ConvertImpliedDecimal
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsPtrOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CreateEmptyXmlTagsForTrailingSeparators
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsPtrOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.MaskSecurityInfo
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsPtrOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.PreserveInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsPtrOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SuspendInterchangeOnError
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsPtrOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.UseDotAsDecimalSeparator
	}).(pulumi.BoolPtrOutput)
}

type X12ProcessingSettingsResponse struct {
	ConvertImpliedDecimal                   *bool `pulumi:"convertImpliedDecimal"`
	CreateEmptyXmlTagsForTrailingSeparators *bool `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        *bool `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     *bool `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               *bool `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                *bool `pulumi:"useDotAsDecimalSeparator"`
}

type X12ProcessingSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12ProcessingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProcessingSettingsResponse)(nil)).Elem()
}

func (o X12ProcessingSettingsResponseOutput) ToX12ProcessingSettingsResponseOutput() X12ProcessingSettingsResponseOutput {
	return o
}

func (o X12ProcessingSettingsResponseOutput) ToX12ProcessingSettingsResponseOutputWithContext(ctx context.Context) X12ProcessingSettingsResponseOutput {
	return o
}

func (o X12ProcessingSettingsResponseOutput) ConvertImpliedDecimal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) *bool { return v.ConvertImpliedDecimal }).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponseOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) *bool { return v.CreateEmptyXmlTagsForTrailingSeparators }).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponseOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) *bool { return v.MaskSecurityInfo }).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponseOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) *bool { return v.PreserveInterchange }).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponseOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) *bool { return v.SuspendInterchangeOnError }).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponseOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) *bool { return v.UseDotAsDecimalSeparator }).(pulumi.BoolPtrOutput)
}

type X12ProcessingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12ProcessingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProcessingSettingsResponse)(nil)).Elem()
}

func (o X12ProcessingSettingsResponsePtrOutput) ToX12ProcessingSettingsResponsePtrOutput() X12ProcessingSettingsResponsePtrOutput {
	return o
}

func (o X12ProcessingSettingsResponsePtrOutput) ToX12ProcessingSettingsResponsePtrOutputWithContext(ctx context.Context) X12ProcessingSettingsResponsePtrOutput {
	return o
}

func (o X12ProcessingSettingsResponsePtrOutput) Elem() X12ProcessingSettingsResponseOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) X12ProcessingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12ProcessingSettingsResponse
		return ret
	}).(X12ProcessingSettingsResponseOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) ConvertImpliedDecimal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ConvertImpliedDecimal
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CreateEmptyXmlTagsForTrailingSeparators
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.MaskSecurityInfo
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PreserveInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SuspendInterchangeOnError
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseDotAsDecimalSeparator
	}).(pulumi.BoolPtrOutput)
}

type X12ProtocolSettings struct {
	AcknowledgementSettings *X12AcknowledgementSettings `pulumi:"acknowledgementSettings"`
	EnvelopeOverrides       []X12EnvelopeOverride       `pulumi:"envelopeOverrides"`
	EnvelopeSettings        *X12EnvelopeSettings        `pulumi:"envelopeSettings"`
	FramingSettings         *X12FramingSettings         `pulumi:"framingSettings"`
	MessageFilter           *X12MessageFilter           `pulumi:"messageFilter"`
	MessageFilterList       []X12MessageIdentifier      `pulumi:"messageFilterList"`
	ProcessingSettings      *X12ProcessingSettings      `pulumi:"processingSettings"`
	SchemaReferences        []X12SchemaReference        `pulumi:"schemaReferences"`
	SecuritySettings        *X12SecuritySettings        `pulumi:"securitySettings"`
	ValidationOverrides     []X12ValidationOverride     `pulumi:"validationOverrides"`
	ValidationSettings      *X12ValidationSettings      `pulumi:"validationSettings"`
	X12DelimiterOverrides   []X12DelimiterOverrides     `pulumi:"x12DelimiterOverrides"`
}





type X12ProtocolSettingsInput interface {
	pulumi.Input

	ToX12ProtocolSettingsOutput() X12ProtocolSettingsOutput
	ToX12ProtocolSettingsOutputWithContext(context.Context) X12ProtocolSettingsOutput
}

type X12ProtocolSettingsArgs struct {
	AcknowledgementSettings X12AcknowledgementSettingsPtrInput `pulumi:"acknowledgementSettings"`
	EnvelopeOverrides       X12EnvelopeOverrideArrayInput      `pulumi:"envelopeOverrides"`
	EnvelopeSettings        X12EnvelopeSettingsPtrInput        `pulumi:"envelopeSettings"`
	FramingSettings         X12FramingSettingsPtrInput         `pulumi:"framingSettings"`
	MessageFilter           X12MessageFilterPtrInput           `pulumi:"messageFilter"`
	MessageFilterList       X12MessageIdentifierArrayInput     `pulumi:"messageFilterList"`
	ProcessingSettings      X12ProcessingSettingsPtrInput      `pulumi:"processingSettings"`
	SchemaReferences        X12SchemaReferenceArrayInput       `pulumi:"schemaReferences"`
	SecuritySettings        X12SecuritySettingsPtrInput        `pulumi:"securitySettings"`
	ValidationOverrides     X12ValidationOverrideArrayInput    `pulumi:"validationOverrides"`
	ValidationSettings      X12ValidationSettingsPtrInput      `pulumi:"validationSettings"`
	X12DelimiterOverrides   X12DelimiterOverridesArrayInput    `pulumi:"x12DelimiterOverrides"`
}

func (X12ProtocolSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProtocolSettings)(nil)).Elem()
}

func (i X12ProtocolSettingsArgs) ToX12ProtocolSettingsOutput() X12ProtocolSettingsOutput {
	return i.ToX12ProtocolSettingsOutputWithContext(context.Background())
}

func (i X12ProtocolSettingsArgs) ToX12ProtocolSettingsOutputWithContext(ctx context.Context) X12ProtocolSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProtocolSettingsOutput)
}

func (i X12ProtocolSettingsArgs) ToX12ProtocolSettingsPtrOutput() X12ProtocolSettingsPtrOutput {
	return i.ToX12ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i X12ProtocolSettingsArgs) ToX12ProtocolSettingsPtrOutputWithContext(ctx context.Context) X12ProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProtocolSettingsOutput).ToX12ProtocolSettingsPtrOutputWithContext(ctx)
}









type X12ProtocolSettingsPtrInput interface {
	pulumi.Input

	ToX12ProtocolSettingsPtrOutput() X12ProtocolSettingsPtrOutput
	ToX12ProtocolSettingsPtrOutputWithContext(context.Context) X12ProtocolSettingsPtrOutput
}

type x12protocolSettingsPtrType X12ProtocolSettingsArgs

func X12ProtocolSettingsPtr(v *X12ProtocolSettingsArgs) X12ProtocolSettingsPtrInput {
	return (*x12protocolSettingsPtrType)(v)
}

func (*x12protocolSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProtocolSettings)(nil)).Elem()
}

func (i *x12protocolSettingsPtrType) ToX12ProtocolSettingsPtrOutput() X12ProtocolSettingsPtrOutput {
	return i.ToX12ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i *x12protocolSettingsPtrType) ToX12ProtocolSettingsPtrOutputWithContext(ctx context.Context) X12ProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProtocolSettingsPtrOutput)
}

type X12ProtocolSettingsOutput struct{ *pulumi.OutputState }

func (X12ProtocolSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProtocolSettings)(nil)).Elem()
}

func (o X12ProtocolSettingsOutput) ToX12ProtocolSettingsOutput() X12ProtocolSettingsOutput {
	return o
}

func (o X12ProtocolSettingsOutput) ToX12ProtocolSettingsOutputWithContext(ctx context.Context) X12ProtocolSettingsOutput {
	return o
}

func (o X12ProtocolSettingsOutput) ToX12ProtocolSettingsPtrOutput() X12ProtocolSettingsPtrOutput {
	return o.ToX12ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (o X12ProtocolSettingsOutput) ToX12ProtocolSettingsPtrOutputWithContext(ctx context.Context) X12ProtocolSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12ProtocolSettings) *X12ProtocolSettings {
		return &v
	}).(X12ProtocolSettingsPtrOutput)
}

func (o X12ProtocolSettingsOutput) AcknowledgementSettings() X12AcknowledgementSettingsPtrOutput {
	return o.ApplyT(func(v X12ProtocolSettings) *X12AcknowledgementSettings { return v.AcknowledgementSettings }).(X12AcknowledgementSettingsPtrOutput)
}

func (o X12ProtocolSettingsOutput) EnvelopeOverrides() X12EnvelopeOverrideArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettings) []X12EnvelopeOverride { return v.EnvelopeOverrides }).(X12EnvelopeOverrideArrayOutput)
}

func (o X12ProtocolSettingsOutput) EnvelopeSettings() X12EnvelopeSettingsPtrOutput {
	return o.ApplyT(func(v X12ProtocolSettings) *X12EnvelopeSettings { return v.EnvelopeSettings }).(X12EnvelopeSettingsPtrOutput)
}

func (o X12ProtocolSettingsOutput) FramingSettings() X12FramingSettingsPtrOutput {
	return o.ApplyT(func(v X12ProtocolSettings) *X12FramingSettings { return v.FramingSettings }).(X12FramingSettingsPtrOutput)
}

func (o X12ProtocolSettingsOutput) MessageFilter() X12MessageFilterPtrOutput {
	return o.ApplyT(func(v X12ProtocolSettings) *X12MessageFilter { return v.MessageFilter }).(X12MessageFilterPtrOutput)
}

func (o X12ProtocolSettingsOutput) MessageFilterList() X12MessageIdentifierArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettings) []X12MessageIdentifier { return v.MessageFilterList }).(X12MessageIdentifierArrayOutput)
}

func (o X12ProtocolSettingsOutput) ProcessingSettings() X12ProcessingSettingsPtrOutput {
	return o.ApplyT(func(v X12ProtocolSettings) *X12ProcessingSettings { return v.ProcessingSettings }).(X12ProcessingSettingsPtrOutput)
}

func (o X12ProtocolSettingsOutput) SchemaReferences() X12SchemaReferenceArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettings) []X12SchemaReference { return v.SchemaReferences }).(X12SchemaReferenceArrayOutput)
}

func (o X12ProtocolSettingsOutput) SecuritySettings() X12SecuritySettingsPtrOutput {
	return o.ApplyT(func(v X12ProtocolSettings) *X12SecuritySettings { return v.SecuritySettings }).(X12SecuritySettingsPtrOutput)
}

func (o X12ProtocolSettingsOutput) ValidationOverrides() X12ValidationOverrideArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettings) []X12ValidationOverride { return v.ValidationOverrides }).(X12ValidationOverrideArrayOutput)
}

func (o X12ProtocolSettingsOutput) ValidationSettings() X12ValidationSettingsPtrOutput {
	return o.ApplyT(func(v X12ProtocolSettings) *X12ValidationSettings { return v.ValidationSettings }).(X12ValidationSettingsPtrOutput)
}

func (o X12ProtocolSettingsOutput) X12DelimiterOverrides() X12DelimiterOverridesArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettings) []X12DelimiterOverrides { return v.X12DelimiterOverrides }).(X12DelimiterOverridesArrayOutput)
}

type X12ProtocolSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12ProtocolSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProtocolSettings)(nil)).Elem()
}

func (o X12ProtocolSettingsPtrOutput) ToX12ProtocolSettingsPtrOutput() X12ProtocolSettingsPtrOutput {
	return o
}

func (o X12ProtocolSettingsPtrOutput) ToX12ProtocolSettingsPtrOutputWithContext(ctx context.Context) X12ProtocolSettingsPtrOutput {
	return o
}

func (o X12ProtocolSettingsPtrOutput) Elem() X12ProtocolSettingsOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) X12ProtocolSettings {
		if v != nil {
			return *v
		}
		var ret X12ProtocolSettings
		return ret
	}).(X12ProtocolSettingsOutput)
}

func (o X12ProtocolSettingsPtrOutput) AcknowledgementSettings() X12AcknowledgementSettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12AcknowledgementSettings {
		if v == nil {
			return nil
		}
		return v.AcknowledgementSettings
	}).(X12AcknowledgementSettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) EnvelopeOverrides() X12EnvelopeOverrideArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) []X12EnvelopeOverride {
		if v == nil {
			return nil
		}
		return v.EnvelopeOverrides
	}).(X12EnvelopeOverrideArrayOutput)
}

func (o X12ProtocolSettingsPtrOutput) EnvelopeSettings() X12EnvelopeSettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12EnvelopeSettings {
		if v == nil {
			return nil
		}
		return v.EnvelopeSettings
	}).(X12EnvelopeSettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) FramingSettings() X12FramingSettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12FramingSettings {
		if v == nil {
			return nil
		}
		return v.FramingSettings
	}).(X12FramingSettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) MessageFilter() X12MessageFilterPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12MessageFilter {
		if v == nil {
			return nil
		}
		return v.MessageFilter
	}).(X12MessageFilterPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) MessageFilterList() X12MessageIdentifierArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) []X12MessageIdentifier {
		if v == nil {
			return nil
		}
		return v.MessageFilterList
	}).(X12MessageIdentifierArrayOutput)
}

func (o X12ProtocolSettingsPtrOutput) ProcessingSettings() X12ProcessingSettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12ProcessingSettings {
		if v == nil {
			return nil
		}
		return v.ProcessingSettings
	}).(X12ProcessingSettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) SchemaReferences() X12SchemaReferenceArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) []X12SchemaReference {
		if v == nil {
			return nil
		}
		return v.SchemaReferences
	}).(X12SchemaReferenceArrayOutput)
}

func (o X12ProtocolSettingsPtrOutput) SecuritySettings() X12SecuritySettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12SecuritySettings {
		if v == nil {
			return nil
		}
		return v.SecuritySettings
	}).(X12SecuritySettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) ValidationOverrides() X12ValidationOverrideArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) []X12ValidationOverride {
		if v == nil {
			return nil
		}
		return v.ValidationOverrides
	}).(X12ValidationOverrideArrayOutput)
}

func (o X12ProtocolSettingsPtrOutput) ValidationSettings() X12ValidationSettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12ValidationSettings {
		if v == nil {
			return nil
		}
		return v.ValidationSettings
	}).(X12ValidationSettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) X12DelimiterOverrides() X12DelimiterOverridesArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) []X12DelimiterOverrides {
		if v == nil {
			return nil
		}
		return v.X12DelimiterOverrides
	}).(X12DelimiterOverridesArrayOutput)
}

type X12ProtocolSettingsResponse struct {
	AcknowledgementSettings *X12AcknowledgementSettingsResponse `pulumi:"acknowledgementSettings"`
	EnvelopeOverrides       []X12EnvelopeOverrideResponse       `pulumi:"envelopeOverrides"`
	EnvelopeSettings        *X12EnvelopeSettingsResponse        `pulumi:"envelopeSettings"`
	FramingSettings         *X12FramingSettingsResponse         `pulumi:"framingSettings"`
	MessageFilter           *X12MessageFilterResponse           `pulumi:"messageFilter"`
	MessageFilterList       []X12MessageIdentifierResponse      `pulumi:"messageFilterList"`
	ProcessingSettings      *X12ProcessingSettingsResponse      `pulumi:"processingSettings"`
	SchemaReferences        []X12SchemaReferenceResponse        `pulumi:"schemaReferences"`
	SecuritySettings        *X12SecuritySettingsResponse        `pulumi:"securitySettings"`
	ValidationOverrides     []X12ValidationOverrideResponse     `pulumi:"validationOverrides"`
	ValidationSettings      *X12ValidationSettingsResponse      `pulumi:"validationSettings"`
	X12DelimiterOverrides   []X12DelimiterOverridesResponse     `pulumi:"x12DelimiterOverrides"`
}

type X12ProtocolSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12ProtocolSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProtocolSettingsResponse)(nil)).Elem()
}

func (o X12ProtocolSettingsResponseOutput) ToX12ProtocolSettingsResponseOutput() X12ProtocolSettingsResponseOutput {
	return o
}

func (o X12ProtocolSettingsResponseOutput) ToX12ProtocolSettingsResponseOutputWithContext(ctx context.Context) X12ProtocolSettingsResponseOutput {
	return o
}

func (o X12ProtocolSettingsResponseOutput) AcknowledgementSettings() X12AcknowledgementSettingsResponsePtrOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) *X12AcknowledgementSettingsResponse {
		return v.AcknowledgementSettings
	}).(X12AcknowledgementSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponseOutput) EnvelopeOverrides() X12EnvelopeOverrideResponseArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) []X12EnvelopeOverrideResponse { return v.EnvelopeOverrides }).(X12EnvelopeOverrideResponseArrayOutput)
}

func (o X12ProtocolSettingsResponseOutput) EnvelopeSettings() X12EnvelopeSettingsResponsePtrOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) *X12EnvelopeSettingsResponse { return v.EnvelopeSettings }).(X12EnvelopeSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponseOutput) FramingSettings() X12FramingSettingsResponsePtrOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) *X12FramingSettingsResponse { return v.FramingSettings }).(X12FramingSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponseOutput) MessageFilter() X12MessageFilterResponsePtrOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) *X12MessageFilterResponse { return v.MessageFilter }).(X12MessageFilterResponsePtrOutput)
}

func (o X12ProtocolSettingsResponseOutput) MessageFilterList() X12MessageIdentifierResponseArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) []X12MessageIdentifierResponse { return v.MessageFilterList }).(X12MessageIdentifierResponseArrayOutput)
}

func (o X12ProtocolSettingsResponseOutput) ProcessingSettings() X12ProcessingSettingsResponsePtrOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) *X12ProcessingSettingsResponse { return v.ProcessingSettings }).(X12ProcessingSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponseOutput) SchemaReferences() X12SchemaReferenceResponseArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) []X12SchemaReferenceResponse { return v.SchemaReferences }).(X12SchemaReferenceResponseArrayOutput)
}

func (o X12ProtocolSettingsResponseOutput) SecuritySettings() X12SecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) *X12SecuritySettingsResponse { return v.SecuritySettings }).(X12SecuritySettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponseOutput) ValidationOverrides() X12ValidationOverrideResponseArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) []X12ValidationOverrideResponse { return v.ValidationOverrides }).(X12ValidationOverrideResponseArrayOutput)
}

func (o X12ProtocolSettingsResponseOutput) ValidationSettings() X12ValidationSettingsResponsePtrOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) *X12ValidationSettingsResponse { return v.ValidationSettings }).(X12ValidationSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponseOutput) X12DelimiterOverrides() X12DelimiterOverridesResponseArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) []X12DelimiterOverridesResponse { return v.X12DelimiterOverrides }).(X12DelimiterOverridesResponseArrayOutput)
}

type X12ProtocolSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12ProtocolSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProtocolSettingsResponse)(nil)).Elem()
}

func (o X12ProtocolSettingsResponsePtrOutput) ToX12ProtocolSettingsResponsePtrOutput() X12ProtocolSettingsResponsePtrOutput {
	return o
}

func (o X12ProtocolSettingsResponsePtrOutput) ToX12ProtocolSettingsResponsePtrOutputWithContext(ctx context.Context) X12ProtocolSettingsResponsePtrOutput {
	return o
}

func (o X12ProtocolSettingsResponsePtrOutput) Elem() X12ProtocolSettingsResponseOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) X12ProtocolSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12ProtocolSettingsResponse
		return ret
	}).(X12ProtocolSettingsResponseOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) AcknowledgementSettings() X12AcknowledgementSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12AcknowledgementSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AcknowledgementSettings
	}).(X12AcknowledgementSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) EnvelopeOverrides() X12EnvelopeOverrideResponseArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) []X12EnvelopeOverrideResponse {
		if v == nil {
			return nil
		}
		return v.EnvelopeOverrides
	}).(X12EnvelopeOverrideResponseArrayOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) EnvelopeSettings() X12EnvelopeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12EnvelopeSettingsResponse {
		if v == nil {
			return nil
		}
		return v.EnvelopeSettings
	}).(X12EnvelopeSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) FramingSettings() X12FramingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12FramingSettingsResponse {
		if v == nil {
			return nil
		}
		return v.FramingSettings
	}).(X12FramingSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) MessageFilter() X12MessageFilterResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12MessageFilterResponse {
		if v == nil {
			return nil
		}
		return v.MessageFilter
	}).(X12MessageFilterResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) MessageFilterList() X12MessageIdentifierResponseArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) []X12MessageIdentifierResponse {
		if v == nil {
			return nil
		}
		return v.MessageFilterList
	}).(X12MessageIdentifierResponseArrayOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) ProcessingSettings() X12ProcessingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12ProcessingSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ProcessingSettings
	}).(X12ProcessingSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) SchemaReferences() X12SchemaReferenceResponseArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) []X12SchemaReferenceResponse {
		if v == nil {
			return nil
		}
		return v.SchemaReferences
	}).(X12SchemaReferenceResponseArrayOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) SecuritySettings() X12SecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12SecuritySettingsResponse {
		if v == nil {
			return nil
		}
		return v.SecuritySettings
	}).(X12SecuritySettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) ValidationOverrides() X12ValidationOverrideResponseArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) []X12ValidationOverrideResponse {
		if v == nil {
			return nil
		}
		return v.ValidationOverrides
	}).(X12ValidationOverrideResponseArrayOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) ValidationSettings() X12ValidationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12ValidationSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ValidationSettings
	}).(X12ValidationSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) X12DelimiterOverrides() X12DelimiterOverridesResponseArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) []X12DelimiterOverridesResponse {
		if v == nil {
			return nil
		}
		return v.X12DelimiterOverrides
	}).(X12DelimiterOverridesResponseArrayOutput)
}

type X12SchemaReference struct {
	MessageId           *string `pulumi:"messageId"`
	SchemaName          *string `pulumi:"schemaName"`
	SchemaVersion       *string `pulumi:"schemaVersion"`
	SenderApplicationId *string `pulumi:"senderApplicationId"`
}





type X12SchemaReferenceInput interface {
	pulumi.Input

	ToX12SchemaReferenceOutput() X12SchemaReferenceOutput
	ToX12SchemaReferenceOutputWithContext(context.Context) X12SchemaReferenceOutput
}

type X12SchemaReferenceArgs struct {
	MessageId           pulumi.StringPtrInput `pulumi:"messageId"`
	SchemaName          pulumi.StringPtrInput `pulumi:"schemaName"`
	SchemaVersion       pulumi.StringPtrInput `pulumi:"schemaVersion"`
	SenderApplicationId pulumi.StringPtrInput `pulumi:"senderApplicationId"`
}

func (X12SchemaReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SchemaReference)(nil)).Elem()
}

func (i X12SchemaReferenceArgs) ToX12SchemaReferenceOutput() X12SchemaReferenceOutput {
	return i.ToX12SchemaReferenceOutputWithContext(context.Background())
}

func (i X12SchemaReferenceArgs) ToX12SchemaReferenceOutputWithContext(ctx context.Context) X12SchemaReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12SchemaReferenceOutput)
}





type X12SchemaReferenceArrayInput interface {
	pulumi.Input

	ToX12SchemaReferenceArrayOutput() X12SchemaReferenceArrayOutput
	ToX12SchemaReferenceArrayOutputWithContext(context.Context) X12SchemaReferenceArrayOutput
}

type X12SchemaReferenceArray []X12SchemaReferenceInput

func (X12SchemaReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12SchemaReference)(nil)).Elem()
}

func (i X12SchemaReferenceArray) ToX12SchemaReferenceArrayOutput() X12SchemaReferenceArrayOutput {
	return i.ToX12SchemaReferenceArrayOutputWithContext(context.Background())
}

func (i X12SchemaReferenceArray) ToX12SchemaReferenceArrayOutputWithContext(ctx context.Context) X12SchemaReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12SchemaReferenceArrayOutput)
}

type X12SchemaReferenceOutput struct{ *pulumi.OutputState }

func (X12SchemaReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SchemaReference)(nil)).Elem()
}

func (o X12SchemaReferenceOutput) ToX12SchemaReferenceOutput() X12SchemaReferenceOutput {
	return o
}

func (o X12SchemaReferenceOutput) ToX12SchemaReferenceOutputWithContext(ctx context.Context) X12SchemaReferenceOutput {
	return o
}

func (o X12SchemaReferenceOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SchemaReference) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o X12SchemaReferenceOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SchemaReference) *string { return v.SchemaName }).(pulumi.StringPtrOutput)
}

func (o X12SchemaReferenceOutput) SchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SchemaReference) *string { return v.SchemaVersion }).(pulumi.StringPtrOutput)
}

func (o X12SchemaReferenceOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SchemaReference) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

type X12SchemaReferenceArrayOutput struct{ *pulumi.OutputState }

func (X12SchemaReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12SchemaReference)(nil)).Elem()
}

func (o X12SchemaReferenceArrayOutput) ToX12SchemaReferenceArrayOutput() X12SchemaReferenceArrayOutput {
	return o
}

func (o X12SchemaReferenceArrayOutput) ToX12SchemaReferenceArrayOutputWithContext(ctx context.Context) X12SchemaReferenceArrayOutput {
	return o
}

func (o X12SchemaReferenceArrayOutput) Index(i pulumi.IntInput) X12SchemaReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12SchemaReference {
		return vs[0].([]X12SchemaReference)[vs[1].(int)]
	}).(X12SchemaReferenceOutput)
}

type X12SchemaReferenceResponse struct {
	MessageId           *string `pulumi:"messageId"`
	SchemaName          *string `pulumi:"schemaName"`
	SchemaVersion       *string `pulumi:"schemaVersion"`
	SenderApplicationId *string `pulumi:"senderApplicationId"`
}

type X12SchemaReferenceResponseOutput struct{ *pulumi.OutputState }

func (X12SchemaReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SchemaReferenceResponse)(nil)).Elem()
}

func (o X12SchemaReferenceResponseOutput) ToX12SchemaReferenceResponseOutput() X12SchemaReferenceResponseOutput {
	return o
}

func (o X12SchemaReferenceResponseOutput) ToX12SchemaReferenceResponseOutputWithContext(ctx context.Context) X12SchemaReferenceResponseOutput {
	return o
}

func (o X12SchemaReferenceResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SchemaReferenceResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o X12SchemaReferenceResponseOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SchemaReferenceResponse) *string { return v.SchemaName }).(pulumi.StringPtrOutput)
}

func (o X12SchemaReferenceResponseOutput) SchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SchemaReferenceResponse) *string { return v.SchemaVersion }).(pulumi.StringPtrOutput)
}

func (o X12SchemaReferenceResponseOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SchemaReferenceResponse) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

type X12SchemaReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (X12SchemaReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12SchemaReferenceResponse)(nil)).Elem()
}

func (o X12SchemaReferenceResponseArrayOutput) ToX12SchemaReferenceResponseArrayOutput() X12SchemaReferenceResponseArrayOutput {
	return o
}

func (o X12SchemaReferenceResponseArrayOutput) ToX12SchemaReferenceResponseArrayOutputWithContext(ctx context.Context) X12SchemaReferenceResponseArrayOutput {
	return o
}

func (o X12SchemaReferenceResponseArrayOutput) Index(i pulumi.IntInput) X12SchemaReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12SchemaReferenceResponse {
		return vs[0].([]X12SchemaReferenceResponse)[vs[1].(int)]
	}).(X12SchemaReferenceResponseOutput)
}

type X12SecuritySettings struct {
	AuthorizationQualifier *string `pulumi:"authorizationQualifier"`
	AuthorizationValue     *string `pulumi:"authorizationValue"`
	PasswordValue          *string `pulumi:"passwordValue"`
	SecurityQualifier      *string `pulumi:"securityQualifier"`
}





type X12SecuritySettingsInput interface {
	pulumi.Input

	ToX12SecuritySettingsOutput() X12SecuritySettingsOutput
	ToX12SecuritySettingsOutputWithContext(context.Context) X12SecuritySettingsOutput
}

type X12SecuritySettingsArgs struct {
	AuthorizationQualifier pulumi.StringPtrInput `pulumi:"authorizationQualifier"`
	AuthorizationValue     pulumi.StringPtrInput `pulumi:"authorizationValue"`
	PasswordValue          pulumi.StringPtrInput `pulumi:"passwordValue"`
	SecurityQualifier      pulumi.StringPtrInput `pulumi:"securityQualifier"`
}

func (X12SecuritySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SecuritySettings)(nil)).Elem()
}

func (i X12SecuritySettingsArgs) ToX12SecuritySettingsOutput() X12SecuritySettingsOutput {
	return i.ToX12SecuritySettingsOutputWithContext(context.Background())
}

func (i X12SecuritySettingsArgs) ToX12SecuritySettingsOutputWithContext(ctx context.Context) X12SecuritySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12SecuritySettingsOutput)
}

func (i X12SecuritySettingsArgs) ToX12SecuritySettingsPtrOutput() X12SecuritySettingsPtrOutput {
	return i.ToX12SecuritySettingsPtrOutputWithContext(context.Background())
}

func (i X12SecuritySettingsArgs) ToX12SecuritySettingsPtrOutputWithContext(ctx context.Context) X12SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12SecuritySettingsOutput).ToX12SecuritySettingsPtrOutputWithContext(ctx)
}









type X12SecuritySettingsPtrInput interface {
	pulumi.Input

	ToX12SecuritySettingsPtrOutput() X12SecuritySettingsPtrOutput
	ToX12SecuritySettingsPtrOutputWithContext(context.Context) X12SecuritySettingsPtrOutput
}

type x12securitySettingsPtrType X12SecuritySettingsArgs

func X12SecuritySettingsPtr(v *X12SecuritySettingsArgs) X12SecuritySettingsPtrInput {
	return (*x12securitySettingsPtrType)(v)
}

func (*x12securitySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12SecuritySettings)(nil)).Elem()
}

func (i *x12securitySettingsPtrType) ToX12SecuritySettingsPtrOutput() X12SecuritySettingsPtrOutput {
	return i.ToX12SecuritySettingsPtrOutputWithContext(context.Background())
}

func (i *x12securitySettingsPtrType) ToX12SecuritySettingsPtrOutputWithContext(ctx context.Context) X12SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12SecuritySettingsPtrOutput)
}

type X12SecuritySettingsOutput struct{ *pulumi.OutputState }

func (X12SecuritySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SecuritySettings)(nil)).Elem()
}

func (o X12SecuritySettingsOutput) ToX12SecuritySettingsOutput() X12SecuritySettingsOutput {
	return o
}

func (o X12SecuritySettingsOutput) ToX12SecuritySettingsOutputWithContext(ctx context.Context) X12SecuritySettingsOutput {
	return o
}

func (o X12SecuritySettingsOutput) ToX12SecuritySettingsPtrOutput() X12SecuritySettingsPtrOutput {
	return o.ToX12SecuritySettingsPtrOutputWithContext(context.Background())
}

func (o X12SecuritySettingsOutput) ToX12SecuritySettingsPtrOutputWithContext(ctx context.Context) X12SecuritySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12SecuritySettings) *X12SecuritySettings {
		return &v
	}).(X12SecuritySettingsPtrOutput)
}

func (o X12SecuritySettingsOutput) AuthorizationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettings) *string { return v.AuthorizationQualifier }).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsOutput) AuthorizationValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettings) *string { return v.AuthorizationValue }).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsOutput) PasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettings) *string { return v.PasswordValue }).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsOutput) SecurityQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettings) *string { return v.SecurityQualifier }).(pulumi.StringPtrOutput)
}

type X12SecuritySettingsPtrOutput struct{ *pulumi.OutputState }

func (X12SecuritySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12SecuritySettings)(nil)).Elem()
}

func (o X12SecuritySettingsPtrOutput) ToX12SecuritySettingsPtrOutput() X12SecuritySettingsPtrOutput {
	return o
}

func (o X12SecuritySettingsPtrOutput) ToX12SecuritySettingsPtrOutputWithContext(ctx context.Context) X12SecuritySettingsPtrOutput {
	return o
}

func (o X12SecuritySettingsPtrOutput) Elem() X12SecuritySettingsOutput {
	return o.ApplyT(func(v *X12SecuritySettings) X12SecuritySettings {
		if v != nil {
			return *v
		}
		var ret X12SecuritySettings
		return ret
	}).(X12SecuritySettingsOutput)
}

func (o X12SecuritySettingsPtrOutput) AuthorizationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationQualifier
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsPtrOutput) AuthorizationValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationValue
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsPtrOutput) PasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.PasswordValue
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsPtrOutput) SecurityQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.SecurityQualifier
	}).(pulumi.StringPtrOutput)
}

type X12SecuritySettingsResponse struct {
	AuthorizationQualifier *string `pulumi:"authorizationQualifier"`
	AuthorizationValue     *string `pulumi:"authorizationValue"`
	PasswordValue          *string `pulumi:"passwordValue"`
	SecurityQualifier      *string `pulumi:"securityQualifier"`
}

type X12SecuritySettingsResponseOutput struct{ *pulumi.OutputState }

func (X12SecuritySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SecuritySettingsResponse)(nil)).Elem()
}

func (o X12SecuritySettingsResponseOutput) ToX12SecuritySettingsResponseOutput() X12SecuritySettingsResponseOutput {
	return o
}

func (o X12SecuritySettingsResponseOutput) ToX12SecuritySettingsResponseOutputWithContext(ctx context.Context) X12SecuritySettingsResponseOutput {
	return o
}

func (o X12SecuritySettingsResponseOutput) AuthorizationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettingsResponse) *string { return v.AuthorizationQualifier }).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponseOutput) AuthorizationValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettingsResponse) *string { return v.AuthorizationValue }).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponseOutput) PasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettingsResponse) *string { return v.PasswordValue }).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponseOutput) SecurityQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettingsResponse) *string { return v.SecurityQualifier }).(pulumi.StringPtrOutput)
}

type X12SecuritySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12SecuritySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12SecuritySettingsResponse)(nil)).Elem()
}

func (o X12SecuritySettingsResponsePtrOutput) ToX12SecuritySettingsResponsePtrOutput() X12SecuritySettingsResponsePtrOutput {
	return o
}

func (o X12SecuritySettingsResponsePtrOutput) ToX12SecuritySettingsResponsePtrOutputWithContext(ctx context.Context) X12SecuritySettingsResponsePtrOutput {
	return o
}

func (o X12SecuritySettingsResponsePtrOutput) Elem() X12SecuritySettingsResponseOutput {
	return o.ApplyT(func(v *X12SecuritySettingsResponse) X12SecuritySettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12SecuritySettingsResponse
		return ret
	}).(X12SecuritySettingsResponseOutput)
}

func (o X12SecuritySettingsResponsePtrOutput) AuthorizationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationQualifier
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponsePtrOutput) AuthorizationValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationValue
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponsePtrOutput) PasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PasswordValue
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponsePtrOutput) SecurityQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecurityQualifier
	}).(pulumi.StringPtrOutput)
}

type X12ValidationOverride struct {
	AllowLeadingAndTrailingSpacesAndZeroes *bool                    `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	MessageId                              *string                  `pulumi:"messageId"`
	TrailingSeparatorPolicy                *TrailingSeparatorPolicy `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  *bool                    `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                   *bool                    `pulumi:"validateCharacterSet"`
	ValidateEDITypes                       *bool                    `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       *bool                    `pulumi:"validateXSDTypes"`
}





type X12ValidationOverrideInput interface {
	pulumi.Input

	ToX12ValidationOverrideOutput() X12ValidationOverrideOutput
	ToX12ValidationOverrideOutputWithContext(context.Context) X12ValidationOverrideOutput
}

type X12ValidationOverrideArgs struct {
	AllowLeadingAndTrailingSpacesAndZeroes pulumi.BoolPtrInput             `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	MessageId                              pulumi.StringPtrInput           `pulumi:"messageId"`
	TrailingSeparatorPolicy                TrailingSeparatorPolicyPtrInput `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  pulumi.BoolPtrInput             `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                   pulumi.BoolPtrInput             `pulumi:"validateCharacterSet"`
	ValidateEDITypes                       pulumi.BoolPtrInput             `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       pulumi.BoolPtrInput             `pulumi:"validateXSDTypes"`
}

func (X12ValidationOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationOverride)(nil)).Elem()
}

func (i X12ValidationOverrideArgs) ToX12ValidationOverrideOutput() X12ValidationOverrideOutput {
	return i.ToX12ValidationOverrideOutputWithContext(context.Background())
}

func (i X12ValidationOverrideArgs) ToX12ValidationOverrideOutputWithContext(ctx context.Context) X12ValidationOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ValidationOverrideOutput)
}





type X12ValidationOverrideArrayInput interface {
	pulumi.Input

	ToX12ValidationOverrideArrayOutput() X12ValidationOverrideArrayOutput
	ToX12ValidationOverrideArrayOutputWithContext(context.Context) X12ValidationOverrideArrayOutput
}

type X12ValidationOverrideArray []X12ValidationOverrideInput

func (X12ValidationOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12ValidationOverride)(nil)).Elem()
}

func (i X12ValidationOverrideArray) ToX12ValidationOverrideArrayOutput() X12ValidationOverrideArrayOutput {
	return i.ToX12ValidationOverrideArrayOutputWithContext(context.Background())
}

func (i X12ValidationOverrideArray) ToX12ValidationOverrideArrayOutputWithContext(ctx context.Context) X12ValidationOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ValidationOverrideArrayOutput)
}

type X12ValidationOverrideOutput struct{ *pulumi.OutputState }

func (X12ValidationOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationOverride)(nil)).Elem()
}

func (o X12ValidationOverrideOutput) ToX12ValidationOverrideOutput() X12ValidationOverrideOutput {
	return o
}

func (o X12ValidationOverrideOutput) ToX12ValidationOverrideOutputWithContext(ctx context.Context) X12ValidationOverrideOutput {
	return o
}

func (o X12ValidationOverrideOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationOverride) *bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationOverrideOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12ValidationOverride) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o X12ValidationOverrideOutput) TrailingSeparatorPolicy() TrailingSeparatorPolicyPtrOutput {
	return o.ApplyT(func(v X12ValidationOverride) *TrailingSeparatorPolicy { return v.TrailingSeparatorPolicy }).(TrailingSeparatorPolicyPtrOutput)
}

func (o X12ValidationOverrideOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationOverride) *bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationOverrideOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationOverride) *bool { return v.ValidateCharacterSet }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationOverrideOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationOverride) *bool { return v.ValidateEDITypes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationOverrideOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationOverride) *bool { return v.ValidateXSDTypes }).(pulumi.BoolPtrOutput)
}

type X12ValidationOverrideArrayOutput struct{ *pulumi.OutputState }

func (X12ValidationOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12ValidationOverride)(nil)).Elem()
}

func (o X12ValidationOverrideArrayOutput) ToX12ValidationOverrideArrayOutput() X12ValidationOverrideArrayOutput {
	return o
}

func (o X12ValidationOverrideArrayOutput) ToX12ValidationOverrideArrayOutputWithContext(ctx context.Context) X12ValidationOverrideArrayOutput {
	return o
}

func (o X12ValidationOverrideArrayOutput) Index(i pulumi.IntInput) X12ValidationOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12ValidationOverride {
		return vs[0].([]X12ValidationOverride)[vs[1].(int)]
	}).(X12ValidationOverrideOutput)
}

type X12ValidationOverrideResponse struct {
	AllowLeadingAndTrailingSpacesAndZeroes *bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	MessageId                              *string `pulumi:"messageId"`
	TrailingSeparatorPolicy                *string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  *bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                   *bool   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                       *bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       *bool   `pulumi:"validateXSDTypes"`
}

type X12ValidationOverrideResponseOutput struct{ *pulumi.OutputState }

func (X12ValidationOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationOverrideResponse)(nil)).Elem()
}

func (o X12ValidationOverrideResponseOutput) ToX12ValidationOverrideResponseOutput() X12ValidationOverrideResponseOutput {
	return o
}

func (o X12ValidationOverrideResponseOutput) ToX12ValidationOverrideResponseOutputWithContext(ctx context.Context) X12ValidationOverrideResponseOutput {
	return o
}

func (o X12ValidationOverrideResponseOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) *bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationOverrideResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o X12ValidationOverrideResponseOutput) TrailingSeparatorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) *string { return v.TrailingSeparatorPolicy }).(pulumi.StringPtrOutput)
}

func (o X12ValidationOverrideResponseOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) *bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationOverrideResponseOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) *bool { return v.ValidateCharacterSet }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationOverrideResponseOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) *bool { return v.ValidateEDITypes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationOverrideResponseOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) *bool { return v.ValidateXSDTypes }).(pulumi.BoolPtrOutput)
}

type X12ValidationOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (X12ValidationOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12ValidationOverrideResponse)(nil)).Elem()
}

func (o X12ValidationOverrideResponseArrayOutput) ToX12ValidationOverrideResponseArrayOutput() X12ValidationOverrideResponseArrayOutput {
	return o
}

func (o X12ValidationOverrideResponseArrayOutput) ToX12ValidationOverrideResponseArrayOutputWithContext(ctx context.Context) X12ValidationOverrideResponseArrayOutput {
	return o
}

func (o X12ValidationOverrideResponseArrayOutput) Index(i pulumi.IntInput) X12ValidationOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12ValidationOverrideResponse {
		return vs[0].([]X12ValidationOverrideResponse)[vs[1].(int)]
	}).(X12ValidationOverrideResponseOutput)
}

type X12ValidationSettings struct {
	AllowLeadingAndTrailingSpacesAndZeroes    *bool                    `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          *bool                    `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    *bool                    `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber *bool                    `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      *int                     `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   *TrailingSeparatorPolicy `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     *bool                    `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      *bool                    `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          *bool                    `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          *bool                    `pulumi:"validateXSDTypes"`
}





type X12ValidationSettingsInput interface {
	pulumi.Input

	ToX12ValidationSettingsOutput() X12ValidationSettingsOutput
	ToX12ValidationSettingsOutputWithContext(context.Context) X12ValidationSettingsOutput
}

type X12ValidationSettingsArgs struct {
	AllowLeadingAndTrailingSpacesAndZeroes    pulumi.BoolPtrInput             `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          pulumi.BoolPtrInput             `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    pulumi.BoolPtrInput             `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber pulumi.BoolPtrInput             `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      pulumi.IntPtrInput              `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   TrailingSeparatorPolicyPtrInput `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     pulumi.BoolPtrInput             `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      pulumi.BoolPtrInput             `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          pulumi.BoolPtrInput             `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          pulumi.BoolPtrInput             `pulumi:"validateXSDTypes"`
}

func (X12ValidationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationSettings)(nil)).Elem()
}

func (i X12ValidationSettingsArgs) ToX12ValidationSettingsOutput() X12ValidationSettingsOutput {
	return i.ToX12ValidationSettingsOutputWithContext(context.Background())
}

func (i X12ValidationSettingsArgs) ToX12ValidationSettingsOutputWithContext(ctx context.Context) X12ValidationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ValidationSettingsOutput)
}

func (i X12ValidationSettingsArgs) ToX12ValidationSettingsPtrOutput() X12ValidationSettingsPtrOutput {
	return i.ToX12ValidationSettingsPtrOutputWithContext(context.Background())
}

func (i X12ValidationSettingsArgs) ToX12ValidationSettingsPtrOutputWithContext(ctx context.Context) X12ValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ValidationSettingsOutput).ToX12ValidationSettingsPtrOutputWithContext(ctx)
}









type X12ValidationSettingsPtrInput interface {
	pulumi.Input

	ToX12ValidationSettingsPtrOutput() X12ValidationSettingsPtrOutput
	ToX12ValidationSettingsPtrOutputWithContext(context.Context) X12ValidationSettingsPtrOutput
}

type x12validationSettingsPtrType X12ValidationSettingsArgs

func X12ValidationSettingsPtr(v *X12ValidationSettingsArgs) X12ValidationSettingsPtrInput {
	return (*x12validationSettingsPtrType)(v)
}

func (*x12validationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ValidationSettings)(nil)).Elem()
}

func (i *x12validationSettingsPtrType) ToX12ValidationSettingsPtrOutput() X12ValidationSettingsPtrOutput {
	return i.ToX12ValidationSettingsPtrOutputWithContext(context.Background())
}

func (i *x12validationSettingsPtrType) ToX12ValidationSettingsPtrOutputWithContext(ctx context.Context) X12ValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ValidationSettingsPtrOutput)
}

type X12ValidationSettingsOutput struct{ *pulumi.OutputState }

func (X12ValidationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationSettings)(nil)).Elem()
}

func (o X12ValidationSettingsOutput) ToX12ValidationSettingsOutput() X12ValidationSettingsOutput {
	return o
}

func (o X12ValidationSettingsOutput) ToX12ValidationSettingsOutputWithContext(ctx context.Context) X12ValidationSettingsOutput {
	return o
}

func (o X12ValidationSettingsOutput) ToX12ValidationSettingsPtrOutput() X12ValidationSettingsPtrOutput {
	return o.ToX12ValidationSettingsPtrOutputWithContext(context.Background())
}

func (o X12ValidationSettingsOutput) ToX12ValidationSettingsPtrOutputWithContext(ctx context.Context) X12ValidationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12ValidationSettings) *X12ValidationSettings {
		return &v
	}).(X12ValidationSettingsPtrOutput)
}

func (o X12ValidationSettingsOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettings) *bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettings) *bool { return v.CheckDuplicateGroupControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettings) *bool { return v.CheckDuplicateInterchangeControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettings) *bool { return v.CheckDuplicateTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12ValidationSettings) *int { return v.InterchangeControlNumberValidityDays }).(pulumi.IntPtrOutput)
}

func (o X12ValidationSettingsOutput) TrailingSeparatorPolicy() TrailingSeparatorPolicyPtrOutput {
	return o.ApplyT(func(v X12ValidationSettings) *TrailingSeparatorPolicy { return v.TrailingSeparatorPolicy }).(TrailingSeparatorPolicyPtrOutput)
}

func (o X12ValidationSettingsOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettings) *bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettings) *bool { return v.ValidateCharacterSet }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettings) *bool { return v.ValidateEDITypes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettings) *bool { return v.ValidateXSDTypes }).(pulumi.BoolPtrOutput)
}

type X12ValidationSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12ValidationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ValidationSettings)(nil)).Elem()
}

func (o X12ValidationSettingsPtrOutput) ToX12ValidationSettingsPtrOutput() X12ValidationSettingsPtrOutput {
	return o
}

func (o X12ValidationSettingsPtrOutput) ToX12ValidationSettingsPtrOutputWithContext(ctx context.Context) X12ValidationSettingsPtrOutput {
	return o
}

func (o X12ValidationSettingsPtrOutput) Elem() X12ValidationSettingsOutput {
	return o.ApplyT(func(v *X12ValidationSettings) X12ValidationSettings {
		if v != nil {
			return *v
		}
		var ret X12ValidationSettings
		return ret
	}).(X12ValidationSettingsOutput)
}

func (o X12ValidationSettingsPtrOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.AllowLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *int {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) TrailingSeparatorPolicy() TrailingSeparatorPolicyPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *TrailingSeparatorPolicy {
		if v == nil {
			return nil
		}
		return v.TrailingSeparatorPolicy
	}).(TrailingSeparatorPolicyPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.TrimLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateCharacterSet
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateEDITypes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateXSDTypes
	}).(pulumi.BoolPtrOutput)
}

type X12ValidationSettingsResponse struct {
	AllowLeadingAndTrailingSpacesAndZeroes    *bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          *bool   `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    *bool   `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber *bool   `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      *int    `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   *string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     *bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      *bool   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          *bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          *bool   `pulumi:"validateXSDTypes"`
}

type X12ValidationSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12ValidationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationSettingsResponse)(nil)).Elem()
}

func (o X12ValidationSettingsResponseOutput) ToX12ValidationSettingsResponseOutput() X12ValidationSettingsResponseOutput {
	return o
}

func (o X12ValidationSettingsResponseOutput) ToX12ValidationSettingsResponseOutputWithContext(ctx context.Context) X12ValidationSettingsResponseOutput {
	return o
}

func (o X12ValidationSettingsResponseOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) *bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponseOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) *bool { return v.CheckDuplicateGroupControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponseOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) *bool { return v.CheckDuplicateInterchangeControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponseOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) *bool { return v.CheckDuplicateTransactionSetControlNumber }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponseOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) *int { return v.InterchangeControlNumberValidityDays }).(pulumi.IntPtrOutput)
}

func (o X12ValidationSettingsResponseOutput) TrailingSeparatorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) *string { return v.TrailingSeparatorPolicy }).(pulumi.StringPtrOutput)
}

func (o X12ValidationSettingsResponseOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) *bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponseOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) *bool { return v.ValidateCharacterSet }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponseOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) *bool { return v.ValidateEDITypes }).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponseOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) *bool { return v.ValidateXSDTypes }).(pulumi.BoolPtrOutput)
}

type X12ValidationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12ValidationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ValidationSettingsResponse)(nil)).Elem()
}

func (o X12ValidationSettingsResponsePtrOutput) ToX12ValidationSettingsResponsePtrOutput() X12ValidationSettingsResponsePtrOutput {
	return o
}

func (o X12ValidationSettingsResponsePtrOutput) ToX12ValidationSettingsResponsePtrOutputWithContext(ctx context.Context) X12ValidationSettingsResponsePtrOutput {
	return o
}

func (o X12ValidationSettingsResponsePtrOutput) Elem() X12ValidationSettingsResponseOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) X12ValidationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12ValidationSettingsResponse
		return ret
	}).(X12ValidationSettingsResponseOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CheckDuplicateTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) TrailingSeparatorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TrailingSeparatorPolicy
	}).(pulumi.StringPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.TrimLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateCharacterSet
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateEDITypes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateXSDTypes
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AS2AcknowledgementConnectionSettingsOutput{})
	pulumi.RegisterOutputType(AS2AcknowledgementConnectionSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2AcknowledgementConnectionSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2AcknowledgementConnectionSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2AgreementContentOutput{})
	pulumi.RegisterOutputType(AS2AgreementContentPtrOutput{})
	pulumi.RegisterOutputType(AS2AgreementContentResponseOutput{})
	pulumi.RegisterOutputType(AS2AgreementContentResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2EnvelopeSettingsOutput{})
	pulumi.RegisterOutputType(AS2EnvelopeSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2EnvelopeSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2EnvelopeSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2ErrorSettingsOutput{})
	pulumi.RegisterOutputType(AS2ErrorSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2ErrorSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2ErrorSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2MdnSettingsOutput{})
	pulumi.RegisterOutputType(AS2MdnSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2MdnSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2MdnSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2MessageConnectionSettingsOutput{})
	pulumi.RegisterOutputType(AS2MessageConnectionSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2MessageConnectionSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2MessageConnectionSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2OneWayAgreementOutput{})
	pulumi.RegisterOutputType(AS2OneWayAgreementPtrOutput{})
	pulumi.RegisterOutputType(AS2OneWayAgreementResponseOutput{})
	pulumi.RegisterOutputType(AS2OneWayAgreementResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2ProtocolSettingsOutput{})
	pulumi.RegisterOutputType(AS2ProtocolSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2ProtocolSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2ProtocolSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2SecuritySettingsOutput{})
	pulumi.RegisterOutputType(AS2SecuritySettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2SecuritySettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2SecuritySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2ValidationSettingsOutput{})
	pulumi.RegisterOutputType(AS2ValidationSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2ValidationSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2ValidationSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AgreementContentOutput{})
	pulumi.RegisterOutputType(AgreementContentPtrOutput{})
	pulumi.RegisterOutputType(AgreementContentResponseOutput{})
	pulumi.RegisterOutputType(AgreementContentResponsePtrOutput{})
	pulumi.RegisterOutputType(B2BPartnerContentOutput{})
	pulumi.RegisterOutputType(B2BPartnerContentPtrOutput{})
	pulumi.RegisterOutputType(B2BPartnerContentResponseOutput{})
	pulumi.RegisterOutputType(B2BPartnerContentResponsePtrOutput{})
	pulumi.RegisterOutputType(BusinessIdentityOutput{})
	pulumi.RegisterOutputType(BusinessIdentityPtrOutput{})
	pulumi.RegisterOutputType(BusinessIdentityArrayOutput{})
	pulumi.RegisterOutputType(BusinessIdentityResponseOutput{})
	pulumi.RegisterOutputType(BusinessIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(BusinessIdentityResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactAcknowledgementSettingsOutput{})
	pulumi.RegisterOutputType(EdifactAcknowledgementSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactAcknowledgementSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactAcknowledgementSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactAgreementContentOutput{})
	pulumi.RegisterOutputType(EdifactAgreementContentPtrOutput{})
	pulumi.RegisterOutputType(EdifactAgreementContentResponseOutput{})
	pulumi.RegisterOutputType(EdifactAgreementContentResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactDelimiterOverrideOutput{})
	pulumi.RegisterOutputType(EdifactDelimiterOverrideArrayOutput{})
	pulumi.RegisterOutputType(EdifactDelimiterOverrideResponseOutput{})
	pulumi.RegisterOutputType(EdifactDelimiterOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeOverrideOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeOverrideArrayOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeOverrideResponseOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeSettingsOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactFramingSettingsOutput{})
	pulumi.RegisterOutputType(EdifactFramingSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactFramingSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactFramingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactMessageFilterOutput{})
	pulumi.RegisterOutputType(EdifactMessageFilterPtrOutput{})
	pulumi.RegisterOutputType(EdifactMessageFilterResponseOutput{})
	pulumi.RegisterOutputType(EdifactMessageFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactMessageIdentifierOutput{})
	pulumi.RegisterOutputType(EdifactMessageIdentifierArrayOutput{})
	pulumi.RegisterOutputType(EdifactMessageIdentifierResponseOutput{})
	pulumi.RegisterOutputType(EdifactMessageIdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactOneWayAgreementOutput{})
	pulumi.RegisterOutputType(EdifactOneWayAgreementPtrOutput{})
	pulumi.RegisterOutputType(EdifactOneWayAgreementResponseOutput{})
	pulumi.RegisterOutputType(EdifactOneWayAgreementResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactProcessingSettingsOutput{})
	pulumi.RegisterOutputType(EdifactProcessingSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactProcessingSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactProcessingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactProtocolSettingsOutput{})
	pulumi.RegisterOutputType(EdifactProtocolSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactProtocolSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactProtocolSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactSchemaReferenceOutput{})
	pulumi.RegisterOutputType(EdifactSchemaReferenceArrayOutput{})
	pulumi.RegisterOutputType(EdifactSchemaReferenceResponseOutput{})
	pulumi.RegisterOutputType(EdifactSchemaReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactValidationOverrideOutput{})
	pulumi.RegisterOutputType(EdifactValidationOverrideArrayOutput{})
	pulumi.RegisterOutputType(EdifactValidationOverrideResponseOutput{})
	pulumi.RegisterOutputType(EdifactValidationOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactValidationSettingsOutput{})
	pulumi.RegisterOutputType(EdifactValidationSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactValidationSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactValidationSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(IntegrationAccountContentHashResponseOutput{})
	pulumi.RegisterOutputType(IntegrationAccountContentHashResponsePtrOutput{})
	pulumi.RegisterOutputType(IntegrationAccountContentLinkResponseOutput{})
	pulumi.RegisterOutputType(IntegrationAccountSkuOutput{})
	pulumi.RegisterOutputType(IntegrationAccountSkuPtrOutput{})
	pulumi.RegisterOutputType(IntegrationAccountSkuResponseOutput{})
	pulumi.RegisterOutputType(IntegrationAccountSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceKeyVaultOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceKeyVaultPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseKeyVaultOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseKeyVaultPtrOutput{})
	pulumi.RegisterOutputType(PartnerContentOutput{})
	pulumi.RegisterOutputType(PartnerContentPtrOutput{})
	pulumi.RegisterOutputType(PartnerContentResponseOutput{})
	pulumi.RegisterOutputType(PartnerContentResponsePtrOutput{})
	pulumi.RegisterOutputType(X12AcknowledgementSettingsOutput{})
	pulumi.RegisterOutputType(X12AcknowledgementSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12AcknowledgementSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12AcknowledgementSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12AgreementContentOutput{})
	pulumi.RegisterOutputType(X12AgreementContentPtrOutput{})
	pulumi.RegisterOutputType(X12AgreementContentResponseOutput{})
	pulumi.RegisterOutputType(X12AgreementContentResponsePtrOutput{})
	pulumi.RegisterOutputType(X12DelimiterOverridesOutput{})
	pulumi.RegisterOutputType(X12DelimiterOverridesArrayOutput{})
	pulumi.RegisterOutputType(X12DelimiterOverridesResponseOutput{})
	pulumi.RegisterOutputType(X12DelimiterOverridesResponseArrayOutput{})
	pulumi.RegisterOutputType(X12EnvelopeOverrideOutput{})
	pulumi.RegisterOutputType(X12EnvelopeOverrideArrayOutput{})
	pulumi.RegisterOutputType(X12EnvelopeOverrideResponseOutput{})
	pulumi.RegisterOutputType(X12EnvelopeOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(X12EnvelopeSettingsOutput{})
	pulumi.RegisterOutputType(X12EnvelopeSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12EnvelopeSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12EnvelopeSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12FramingSettingsOutput{})
	pulumi.RegisterOutputType(X12FramingSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12FramingSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12FramingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12MessageFilterOutput{})
	pulumi.RegisterOutputType(X12MessageFilterPtrOutput{})
	pulumi.RegisterOutputType(X12MessageFilterResponseOutput{})
	pulumi.RegisterOutputType(X12MessageFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(X12MessageIdentifierOutput{})
	pulumi.RegisterOutputType(X12MessageIdentifierArrayOutput{})
	pulumi.RegisterOutputType(X12MessageIdentifierResponseOutput{})
	pulumi.RegisterOutputType(X12MessageIdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(X12OneWayAgreementOutput{})
	pulumi.RegisterOutputType(X12OneWayAgreementPtrOutput{})
	pulumi.RegisterOutputType(X12OneWayAgreementResponseOutput{})
	pulumi.RegisterOutputType(X12OneWayAgreementResponsePtrOutput{})
	pulumi.RegisterOutputType(X12ProcessingSettingsOutput{})
	pulumi.RegisterOutputType(X12ProcessingSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12ProcessingSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12ProcessingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12ProtocolSettingsOutput{})
	pulumi.RegisterOutputType(X12ProtocolSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12ProtocolSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12ProtocolSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12SchemaReferenceOutput{})
	pulumi.RegisterOutputType(X12SchemaReferenceArrayOutput{})
	pulumi.RegisterOutputType(X12SchemaReferenceResponseOutput{})
	pulumi.RegisterOutputType(X12SchemaReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(X12SecuritySettingsOutput{})
	pulumi.RegisterOutputType(X12SecuritySettingsPtrOutput{})
	pulumi.RegisterOutputType(X12SecuritySettingsResponseOutput{})
	pulumi.RegisterOutputType(X12SecuritySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12ValidationOverrideOutput{})
	pulumi.RegisterOutputType(X12ValidationOverrideArrayOutput{})
	pulumi.RegisterOutputType(X12ValidationOverrideResponseOutput{})
	pulumi.RegisterOutputType(X12ValidationOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(X12ValidationSettingsOutput{})
	pulumi.RegisterOutputType(X12ValidationSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12ValidationSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12ValidationSettingsResponsePtrOutput{})
}
