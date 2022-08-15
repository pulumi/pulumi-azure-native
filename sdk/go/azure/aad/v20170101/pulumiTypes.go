


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DomainSecuritySettings struct {
	NtlmV1            *string `pulumi:"ntlmV1"`
	SyncNtlmPasswords *string `pulumi:"syncNtlmPasswords"`
	TlsV1             *string `pulumi:"tlsV1"`
}





type DomainSecuritySettingsInput interface {
	pulumi.Input

	ToDomainSecuritySettingsOutput() DomainSecuritySettingsOutput
	ToDomainSecuritySettingsOutputWithContext(context.Context) DomainSecuritySettingsOutput
}

type DomainSecuritySettingsArgs struct {
	NtlmV1            pulumi.StringPtrInput `pulumi:"ntlmV1"`
	SyncNtlmPasswords pulumi.StringPtrInput `pulumi:"syncNtlmPasswords"`
	TlsV1             pulumi.StringPtrInput `pulumi:"tlsV1"`
}

func (DomainSecuritySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainSecuritySettings)(nil)).Elem()
}

func (i DomainSecuritySettingsArgs) ToDomainSecuritySettingsOutput() DomainSecuritySettingsOutput {
	return i.ToDomainSecuritySettingsOutputWithContext(context.Background())
}

func (i DomainSecuritySettingsArgs) ToDomainSecuritySettingsOutputWithContext(ctx context.Context) DomainSecuritySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSecuritySettingsOutput)
}

func (i DomainSecuritySettingsArgs) ToDomainSecuritySettingsPtrOutput() DomainSecuritySettingsPtrOutput {
	return i.ToDomainSecuritySettingsPtrOutputWithContext(context.Background())
}

func (i DomainSecuritySettingsArgs) ToDomainSecuritySettingsPtrOutputWithContext(ctx context.Context) DomainSecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSecuritySettingsOutput).ToDomainSecuritySettingsPtrOutputWithContext(ctx)
}









type DomainSecuritySettingsPtrInput interface {
	pulumi.Input

	ToDomainSecuritySettingsPtrOutput() DomainSecuritySettingsPtrOutput
	ToDomainSecuritySettingsPtrOutputWithContext(context.Context) DomainSecuritySettingsPtrOutput
}

type domainSecuritySettingsPtrType DomainSecuritySettingsArgs

func DomainSecuritySettingsPtr(v *DomainSecuritySettingsArgs) DomainSecuritySettingsPtrInput {
	return (*domainSecuritySettingsPtrType)(v)
}

func (*domainSecuritySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainSecuritySettings)(nil)).Elem()
}

func (i *domainSecuritySettingsPtrType) ToDomainSecuritySettingsPtrOutput() DomainSecuritySettingsPtrOutput {
	return i.ToDomainSecuritySettingsPtrOutputWithContext(context.Background())
}

func (i *domainSecuritySettingsPtrType) ToDomainSecuritySettingsPtrOutputWithContext(ctx context.Context) DomainSecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSecuritySettingsPtrOutput)
}

type DomainSecuritySettingsOutput struct{ *pulumi.OutputState }

func (DomainSecuritySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainSecuritySettings)(nil)).Elem()
}

func (o DomainSecuritySettingsOutput) ToDomainSecuritySettingsOutput() DomainSecuritySettingsOutput {
	return o
}

func (o DomainSecuritySettingsOutput) ToDomainSecuritySettingsOutputWithContext(ctx context.Context) DomainSecuritySettingsOutput {
	return o
}

func (o DomainSecuritySettingsOutput) ToDomainSecuritySettingsPtrOutput() DomainSecuritySettingsPtrOutput {
	return o.ToDomainSecuritySettingsPtrOutputWithContext(context.Background())
}

func (o DomainSecuritySettingsOutput) ToDomainSecuritySettingsPtrOutputWithContext(ctx context.Context) DomainSecuritySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainSecuritySettings) *DomainSecuritySettings {
		return &v
	}).(DomainSecuritySettingsPtrOutput)
}

func (o DomainSecuritySettingsOutput) NtlmV1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettings) *string { return v.NtlmV1 }).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsOutput) SyncNtlmPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettings) *string { return v.SyncNtlmPasswords }).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsOutput) TlsV1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettings) *string { return v.TlsV1 }).(pulumi.StringPtrOutput)
}

type DomainSecuritySettingsPtrOutput struct{ *pulumi.OutputState }

func (DomainSecuritySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainSecuritySettings)(nil)).Elem()
}

func (o DomainSecuritySettingsPtrOutput) ToDomainSecuritySettingsPtrOutput() DomainSecuritySettingsPtrOutput {
	return o
}

func (o DomainSecuritySettingsPtrOutput) ToDomainSecuritySettingsPtrOutputWithContext(ctx context.Context) DomainSecuritySettingsPtrOutput {
	return o
}

func (o DomainSecuritySettingsPtrOutput) Elem() DomainSecuritySettingsOutput {
	return o.ApplyT(func(v *DomainSecuritySettings) DomainSecuritySettings {
		if v != nil {
			return *v
		}
		var ret DomainSecuritySettings
		return ret
	}).(DomainSecuritySettingsOutput)
}

func (o DomainSecuritySettingsPtrOutput) NtlmV1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainSecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.NtlmV1
	}).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsPtrOutput) SyncNtlmPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainSecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.SyncNtlmPasswords
	}).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsPtrOutput) TlsV1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainSecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.TlsV1
	}).(pulumi.StringPtrOutput)
}

type DomainSecuritySettingsResponse struct {
	NtlmV1            *string `pulumi:"ntlmV1"`
	SyncNtlmPasswords *string `pulumi:"syncNtlmPasswords"`
	TlsV1             *string `pulumi:"tlsV1"`
}

type DomainSecuritySettingsResponseOutput struct{ *pulumi.OutputState }

func (DomainSecuritySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainSecuritySettingsResponse)(nil)).Elem()
}

func (o DomainSecuritySettingsResponseOutput) ToDomainSecuritySettingsResponseOutput() DomainSecuritySettingsResponseOutput {
	return o
}

func (o DomainSecuritySettingsResponseOutput) ToDomainSecuritySettingsResponseOutputWithContext(ctx context.Context) DomainSecuritySettingsResponseOutput {
	return o
}

func (o DomainSecuritySettingsResponseOutput) NtlmV1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettingsResponse) *string { return v.NtlmV1 }).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsResponseOutput) SyncNtlmPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettingsResponse) *string { return v.SyncNtlmPasswords }).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsResponseOutput) TlsV1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettingsResponse) *string { return v.TlsV1 }).(pulumi.StringPtrOutput)
}

type DomainSecuritySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (DomainSecuritySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainSecuritySettingsResponse)(nil)).Elem()
}

func (o DomainSecuritySettingsResponsePtrOutput) ToDomainSecuritySettingsResponsePtrOutput() DomainSecuritySettingsResponsePtrOutput {
	return o
}

func (o DomainSecuritySettingsResponsePtrOutput) ToDomainSecuritySettingsResponsePtrOutputWithContext(ctx context.Context) DomainSecuritySettingsResponsePtrOutput {
	return o
}

func (o DomainSecuritySettingsResponsePtrOutput) Elem() DomainSecuritySettingsResponseOutput {
	return o.ApplyT(func(v *DomainSecuritySettingsResponse) DomainSecuritySettingsResponse {
		if v != nil {
			return *v
		}
		var ret DomainSecuritySettingsResponse
		return ret
	}).(DomainSecuritySettingsResponseOutput)
}

func (o DomainSecuritySettingsResponsePtrOutput) NtlmV1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainSecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.NtlmV1
	}).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsResponsePtrOutput) SyncNtlmPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainSecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SyncNtlmPasswords
	}).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsResponsePtrOutput) TlsV1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainSecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TlsV1
	}).(pulumi.StringPtrOutput)
}

type HealthAlertResponse struct {
	Id            string `pulumi:"id"`
	Issue         string `pulumi:"issue"`
	LastDetected  string `pulumi:"lastDetected"`
	Name          string `pulumi:"name"`
	Raised        string `pulumi:"raised"`
	ResolutionUri string `pulumi:"resolutionUri"`
	Severity      string `pulumi:"severity"`
}

type HealthAlertResponseOutput struct{ *pulumi.OutputState }

func (HealthAlertResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertResponse)(nil)).Elem()
}

func (o HealthAlertResponseOutput) ToHealthAlertResponseOutput() HealthAlertResponseOutput {
	return o
}

func (o HealthAlertResponseOutput) ToHealthAlertResponseOutputWithContext(ctx context.Context) HealthAlertResponseOutput {
	return o
}

func (o HealthAlertResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v HealthAlertResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o HealthAlertResponseOutput) Issue() pulumi.StringOutput {
	return o.ApplyT(func(v HealthAlertResponse) string { return v.Issue }).(pulumi.StringOutput)
}

func (o HealthAlertResponseOutput) LastDetected() pulumi.StringOutput {
	return o.ApplyT(func(v HealthAlertResponse) string { return v.LastDetected }).(pulumi.StringOutput)
}

func (o HealthAlertResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HealthAlertResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o HealthAlertResponseOutput) Raised() pulumi.StringOutput {
	return o.ApplyT(func(v HealthAlertResponse) string { return v.Raised }).(pulumi.StringOutput)
}

func (o HealthAlertResponseOutput) ResolutionUri() pulumi.StringOutput {
	return o.ApplyT(func(v HealthAlertResponse) string { return v.ResolutionUri }).(pulumi.StringOutput)
}

func (o HealthAlertResponseOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v HealthAlertResponse) string { return v.Severity }).(pulumi.StringOutput)
}

type HealthAlertResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthAlertResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertResponse)(nil)).Elem()
}

func (o HealthAlertResponseArrayOutput) ToHealthAlertResponseArrayOutput() HealthAlertResponseArrayOutput {
	return o
}

func (o HealthAlertResponseArrayOutput) ToHealthAlertResponseArrayOutputWithContext(ctx context.Context) HealthAlertResponseArrayOutput {
	return o
}

func (o HealthAlertResponseArrayOutput) Index(i pulumi.IntInput) HealthAlertResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthAlertResponse {
		return vs[0].([]HealthAlertResponse)[vs[1].(int)]
	}).(HealthAlertResponseOutput)
}

type HealthMonitorResponse struct {
	Details string `pulumi:"details"`
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
}

type HealthMonitorResponseOutput struct{ *pulumi.OutputState }

func (HealthMonitorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthMonitorResponse)(nil)).Elem()
}

func (o HealthMonitorResponseOutput) ToHealthMonitorResponseOutput() HealthMonitorResponseOutput {
	return o
}

func (o HealthMonitorResponseOutput) ToHealthMonitorResponseOutputWithContext(ctx context.Context) HealthMonitorResponseOutput {
	return o
}

func (o HealthMonitorResponseOutput) Details() pulumi.StringOutput {
	return o.ApplyT(func(v HealthMonitorResponse) string { return v.Details }).(pulumi.StringOutput)
}

func (o HealthMonitorResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v HealthMonitorResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o HealthMonitorResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HealthMonitorResponse) string { return v.Name }).(pulumi.StringOutput)
}

type HealthMonitorResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthMonitorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthMonitorResponse)(nil)).Elem()
}

func (o HealthMonitorResponseArrayOutput) ToHealthMonitorResponseArrayOutput() HealthMonitorResponseArrayOutput {
	return o
}

func (o HealthMonitorResponseArrayOutput) ToHealthMonitorResponseArrayOutputWithContext(ctx context.Context) HealthMonitorResponseArrayOutput {
	return o
}

func (o HealthMonitorResponseArrayOutput) Index(i pulumi.IntInput) HealthMonitorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthMonitorResponse {
		return vs[0].([]HealthMonitorResponse)[vs[1].(int)]
	}).(HealthMonitorResponseOutput)
}

type LdapsSettings struct {
	ExternalAccess         *string `pulumi:"externalAccess"`
	Ldaps                  *string `pulumi:"ldaps"`
	PfxCertificate         *string `pulumi:"pfxCertificate"`
	PfxCertificatePassword *string `pulumi:"pfxCertificatePassword"`
}





type LdapsSettingsInput interface {
	pulumi.Input

	ToLdapsSettingsOutput() LdapsSettingsOutput
	ToLdapsSettingsOutputWithContext(context.Context) LdapsSettingsOutput
}

type LdapsSettingsArgs struct {
	ExternalAccess         pulumi.StringPtrInput `pulumi:"externalAccess"`
	Ldaps                  pulumi.StringPtrInput `pulumi:"ldaps"`
	PfxCertificate         pulumi.StringPtrInput `pulumi:"pfxCertificate"`
	PfxCertificatePassword pulumi.StringPtrInput `pulumi:"pfxCertificatePassword"`
}

func (LdapsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LdapsSettings)(nil)).Elem()
}

func (i LdapsSettingsArgs) ToLdapsSettingsOutput() LdapsSettingsOutput {
	return i.ToLdapsSettingsOutputWithContext(context.Background())
}

func (i LdapsSettingsArgs) ToLdapsSettingsOutputWithContext(ctx context.Context) LdapsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapsSettingsOutput)
}

func (i LdapsSettingsArgs) ToLdapsSettingsPtrOutput() LdapsSettingsPtrOutput {
	return i.ToLdapsSettingsPtrOutputWithContext(context.Background())
}

func (i LdapsSettingsArgs) ToLdapsSettingsPtrOutputWithContext(ctx context.Context) LdapsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapsSettingsOutput).ToLdapsSettingsPtrOutputWithContext(ctx)
}









type LdapsSettingsPtrInput interface {
	pulumi.Input

	ToLdapsSettingsPtrOutput() LdapsSettingsPtrOutput
	ToLdapsSettingsPtrOutputWithContext(context.Context) LdapsSettingsPtrOutput
}

type ldapsSettingsPtrType LdapsSettingsArgs

func LdapsSettingsPtr(v *LdapsSettingsArgs) LdapsSettingsPtrInput {
	return (*ldapsSettingsPtrType)(v)
}

func (*ldapsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapsSettings)(nil)).Elem()
}

func (i *ldapsSettingsPtrType) ToLdapsSettingsPtrOutput() LdapsSettingsPtrOutput {
	return i.ToLdapsSettingsPtrOutputWithContext(context.Background())
}

func (i *ldapsSettingsPtrType) ToLdapsSettingsPtrOutputWithContext(ctx context.Context) LdapsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapsSettingsPtrOutput)
}

type LdapsSettingsOutput struct{ *pulumi.OutputState }

func (LdapsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LdapsSettings)(nil)).Elem()
}

func (o LdapsSettingsOutput) ToLdapsSettingsOutput() LdapsSettingsOutput {
	return o
}

func (o LdapsSettingsOutput) ToLdapsSettingsOutputWithContext(ctx context.Context) LdapsSettingsOutput {
	return o
}

func (o LdapsSettingsOutput) ToLdapsSettingsPtrOutput() LdapsSettingsPtrOutput {
	return o.ToLdapsSettingsPtrOutputWithContext(context.Background())
}

func (o LdapsSettingsOutput) ToLdapsSettingsPtrOutputWithContext(ctx context.Context) LdapsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LdapsSettings) *LdapsSettings {
		return &v
	}).(LdapsSettingsPtrOutput)
}

func (o LdapsSettingsOutput) ExternalAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LdapsSettings) *string { return v.ExternalAccess }).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsOutput) Ldaps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LdapsSettings) *string { return v.Ldaps }).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsOutput) PfxCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LdapsSettings) *string { return v.PfxCertificate }).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsOutput) PfxCertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LdapsSettings) *string { return v.PfxCertificatePassword }).(pulumi.StringPtrOutput)
}

type LdapsSettingsPtrOutput struct{ *pulumi.OutputState }

func (LdapsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapsSettings)(nil)).Elem()
}

func (o LdapsSettingsPtrOutput) ToLdapsSettingsPtrOutput() LdapsSettingsPtrOutput {
	return o
}

func (o LdapsSettingsPtrOutput) ToLdapsSettingsPtrOutputWithContext(ctx context.Context) LdapsSettingsPtrOutput {
	return o
}

func (o LdapsSettingsPtrOutput) Elem() LdapsSettingsOutput {
	return o.ApplyT(func(v *LdapsSettings) LdapsSettings {
		if v != nil {
			return *v
		}
		var ret LdapsSettings
		return ret
	}).(LdapsSettingsOutput)
}

func (o LdapsSettingsPtrOutput) ExternalAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettings) *string {
		if v == nil {
			return nil
		}
		return v.ExternalAccess
	}).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsPtrOutput) Ldaps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettings) *string {
		if v == nil {
			return nil
		}
		return v.Ldaps
	}).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsPtrOutput) PfxCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettings) *string {
		if v == nil {
			return nil
		}
		return v.PfxCertificate
	}).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsPtrOutput) PfxCertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettings) *string {
		if v == nil {
			return nil
		}
		return v.PfxCertificatePassword
	}).(pulumi.StringPtrOutput)
}

type LdapsSettingsResponse struct {
	CertificateNotAfter     string  `pulumi:"certificateNotAfter"`
	CertificateThumbprint   string  `pulumi:"certificateThumbprint"`
	ExternalAccess          *string `pulumi:"externalAccess"`
	ExternalAccessIpAddress string  `pulumi:"externalAccessIpAddress"`
	Ldaps                   *string `pulumi:"ldaps"`
	PfxCertificate          *string `pulumi:"pfxCertificate"`
	PfxCertificatePassword  *string `pulumi:"pfxCertificatePassword"`
	PublicCertificate       string  `pulumi:"publicCertificate"`
}

type LdapsSettingsResponseOutput struct{ *pulumi.OutputState }

func (LdapsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LdapsSettingsResponse)(nil)).Elem()
}

func (o LdapsSettingsResponseOutput) ToLdapsSettingsResponseOutput() LdapsSettingsResponseOutput {
	return o
}

func (o LdapsSettingsResponseOutput) ToLdapsSettingsResponseOutputWithContext(ctx context.Context) LdapsSettingsResponseOutput {
	return o
}

func (o LdapsSettingsResponseOutput) CertificateNotAfter() pulumi.StringOutput {
	return o.ApplyT(func(v LdapsSettingsResponse) string { return v.CertificateNotAfter }).(pulumi.StringOutput)
}

func (o LdapsSettingsResponseOutput) CertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LdapsSettingsResponse) string { return v.CertificateThumbprint }).(pulumi.StringOutput)
}

func (o LdapsSettingsResponseOutput) ExternalAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LdapsSettingsResponse) *string { return v.ExternalAccess }).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponseOutput) ExternalAccessIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LdapsSettingsResponse) string { return v.ExternalAccessIpAddress }).(pulumi.StringOutput)
}

func (o LdapsSettingsResponseOutput) Ldaps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LdapsSettingsResponse) *string { return v.Ldaps }).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponseOutput) PfxCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LdapsSettingsResponse) *string { return v.PfxCertificate }).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponseOutput) PfxCertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LdapsSettingsResponse) *string { return v.PfxCertificatePassword }).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponseOutput) PublicCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LdapsSettingsResponse) string { return v.PublicCertificate }).(pulumi.StringOutput)
}

type LdapsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (LdapsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapsSettingsResponse)(nil)).Elem()
}

func (o LdapsSettingsResponsePtrOutput) ToLdapsSettingsResponsePtrOutput() LdapsSettingsResponsePtrOutput {
	return o
}

func (o LdapsSettingsResponsePtrOutput) ToLdapsSettingsResponsePtrOutputWithContext(ctx context.Context) LdapsSettingsResponsePtrOutput {
	return o
}

func (o LdapsSettingsResponsePtrOutput) Elem() LdapsSettingsResponseOutput {
	return o.ApplyT(func(v *LdapsSettingsResponse) LdapsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret LdapsSettingsResponse
		return ret
	}).(LdapsSettingsResponseOutput)
}

func (o LdapsSettingsResponsePtrOutput) CertificateNotAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CertificateNotAfter
	}).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponsePtrOutput) CertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponsePtrOutput) ExternalAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExternalAccess
	}).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponsePtrOutput) ExternalAccessIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExternalAccessIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponsePtrOutput) Ldaps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Ldaps
	}).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponsePtrOutput) PfxCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PfxCertificate
	}).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponsePtrOutput) PfxCertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PfxCertificatePassword
	}).(pulumi.StringPtrOutput)
}

func (o LdapsSettingsResponsePtrOutput) PublicCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PublicCertificate
	}).(pulumi.StringPtrOutput)
}

type NotificationSettings struct {
	AdditionalRecipients []string `pulumi:"additionalRecipients"`
	NotifyDcAdmins       *string  `pulumi:"notifyDcAdmins"`
	NotifyGlobalAdmins   *string  `pulumi:"notifyGlobalAdmins"`
}





type NotificationSettingsInput interface {
	pulumi.Input

	ToNotificationSettingsOutput() NotificationSettingsOutput
	ToNotificationSettingsOutputWithContext(context.Context) NotificationSettingsOutput
}

type NotificationSettingsArgs struct {
	AdditionalRecipients pulumi.StringArrayInput `pulumi:"additionalRecipients"`
	NotifyDcAdmins       pulumi.StringPtrInput   `pulumi:"notifyDcAdmins"`
	NotifyGlobalAdmins   pulumi.StringPtrInput   `pulumi:"notifyGlobalAdmins"`
}

func (NotificationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettings)(nil)).Elem()
}

func (i NotificationSettingsArgs) ToNotificationSettingsOutput() NotificationSettingsOutput {
	return i.ToNotificationSettingsOutputWithContext(context.Background())
}

func (i NotificationSettingsArgs) ToNotificationSettingsOutputWithContext(ctx context.Context) NotificationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsOutput)
}

func (i NotificationSettingsArgs) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return i.ToNotificationSettingsPtrOutputWithContext(context.Background())
}

func (i NotificationSettingsArgs) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsOutput).ToNotificationSettingsPtrOutputWithContext(ctx)
}









type NotificationSettingsPtrInput interface {
	pulumi.Input

	ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput
	ToNotificationSettingsPtrOutputWithContext(context.Context) NotificationSettingsPtrOutput
}

type notificationSettingsPtrType NotificationSettingsArgs

func NotificationSettingsPtr(v *NotificationSettingsArgs) NotificationSettingsPtrInput {
	return (*notificationSettingsPtrType)(v)
}

func (*notificationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettings)(nil)).Elem()
}

func (i *notificationSettingsPtrType) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return i.ToNotificationSettingsPtrOutputWithContext(context.Background())
}

func (i *notificationSettingsPtrType) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsPtrOutput)
}

type NotificationSettingsOutput struct{ *pulumi.OutputState }

func (NotificationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettings)(nil)).Elem()
}

func (o NotificationSettingsOutput) ToNotificationSettingsOutput() NotificationSettingsOutput {
	return o
}

func (o NotificationSettingsOutput) ToNotificationSettingsOutputWithContext(ctx context.Context) NotificationSettingsOutput {
	return o
}

func (o NotificationSettingsOutput) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return o.ToNotificationSettingsPtrOutputWithContext(context.Background())
}

func (o NotificationSettingsOutput) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationSettings) *NotificationSettings {
		return &v
	}).(NotificationSettingsPtrOutput)
}

func (o NotificationSettingsOutput) AdditionalRecipients() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationSettings) []string { return v.AdditionalRecipients }).(pulumi.StringArrayOutput)
}

func (o NotificationSettingsOutput) NotifyDcAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettings) *string { return v.NotifyDcAdmins }).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsOutput) NotifyGlobalAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettings) *string { return v.NotifyGlobalAdmins }).(pulumi.StringPtrOutput)
}

type NotificationSettingsPtrOutput struct{ *pulumi.OutputState }

func (NotificationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettings)(nil)).Elem()
}

func (o NotificationSettingsPtrOutput) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return o
}

func (o NotificationSettingsPtrOutput) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return o
}

func (o NotificationSettingsPtrOutput) Elem() NotificationSettingsOutput {
	return o.ApplyT(func(v *NotificationSettings) NotificationSettings {
		if v != nil {
			return *v
		}
		var ret NotificationSettings
		return ret
	}).(NotificationSettingsOutput)
}

func (o NotificationSettingsPtrOutput) AdditionalRecipients() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NotificationSettings) []string {
		if v == nil {
			return nil
		}
		return v.AdditionalRecipients
	}).(pulumi.StringArrayOutput)
}

func (o NotificationSettingsPtrOutput) NotifyDcAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettings) *string {
		if v == nil {
			return nil
		}
		return v.NotifyDcAdmins
	}).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsPtrOutput) NotifyGlobalAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettings) *string {
		if v == nil {
			return nil
		}
		return v.NotifyGlobalAdmins
	}).(pulumi.StringPtrOutput)
}

type NotificationSettingsResponse struct {
	AdditionalRecipients []string `pulumi:"additionalRecipients"`
	NotifyDcAdmins       *string  `pulumi:"notifyDcAdmins"`
	NotifyGlobalAdmins   *string  `pulumi:"notifyGlobalAdmins"`
}

type NotificationSettingsResponseOutput struct{ *pulumi.OutputState }

func (NotificationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettingsResponse)(nil)).Elem()
}

func (o NotificationSettingsResponseOutput) ToNotificationSettingsResponseOutput() NotificationSettingsResponseOutput {
	return o
}

func (o NotificationSettingsResponseOutput) ToNotificationSettingsResponseOutputWithContext(ctx context.Context) NotificationSettingsResponseOutput {
	return o
}

func (o NotificationSettingsResponseOutput) AdditionalRecipients() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationSettingsResponse) []string { return v.AdditionalRecipients }).(pulumi.StringArrayOutput)
}

func (o NotificationSettingsResponseOutput) NotifyDcAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettingsResponse) *string { return v.NotifyDcAdmins }).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsResponseOutput) NotifyGlobalAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettingsResponse) *string { return v.NotifyGlobalAdmins }).(pulumi.StringPtrOutput)
}

type NotificationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (NotificationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettingsResponse)(nil)).Elem()
}

func (o NotificationSettingsResponsePtrOutput) ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput {
	return o
}

func (o NotificationSettingsResponsePtrOutput) ToNotificationSettingsResponsePtrOutputWithContext(ctx context.Context) NotificationSettingsResponsePtrOutput {
	return o
}

func (o NotificationSettingsResponsePtrOutput) Elem() NotificationSettingsResponseOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) NotificationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret NotificationSettingsResponse
		return ret
	}).(NotificationSettingsResponseOutput)
}

func (o NotificationSettingsResponsePtrOutput) AdditionalRecipients() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.AdditionalRecipients
	}).(pulumi.StringArrayOutput)
}

func (o NotificationSettingsResponsePtrOutput) NotifyDcAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.NotifyDcAdmins
	}).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsResponsePtrOutput) NotifyGlobalAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.NotifyGlobalAdmins
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainSecuritySettingsOutput{})
	pulumi.RegisterOutputType(DomainSecuritySettingsPtrOutput{})
	pulumi.RegisterOutputType(DomainSecuritySettingsResponseOutput{})
	pulumi.RegisterOutputType(DomainSecuritySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(HealthAlertResponseOutput{})
	pulumi.RegisterOutputType(HealthAlertResponseArrayOutput{})
	pulumi.RegisterOutputType(HealthMonitorResponseOutput{})
	pulumi.RegisterOutputType(HealthMonitorResponseArrayOutput{})
	pulumi.RegisterOutputType(LdapsSettingsOutput{})
	pulumi.RegisterOutputType(LdapsSettingsPtrOutput{})
	pulumi.RegisterOutputType(LdapsSettingsResponseOutput{})
	pulumi.RegisterOutputType(LdapsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(NotificationSettingsOutput{})
	pulumi.RegisterOutputType(NotificationSettingsPtrOutput{})
	pulumi.RegisterOutputType(NotificationSettingsResponseOutput{})
	pulumi.RegisterOutputType(NotificationSettingsResponsePtrOutput{})
}
