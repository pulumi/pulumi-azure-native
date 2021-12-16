


package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerAccountResponse struct {
	AccountName *string `pulumi:"accountName"`
	Password    *string `pulumi:"password"`
	Spn         *string `pulumi:"spn"`
}





type ContainerAccountResponseInput interface {
	pulumi.Input

	ToContainerAccountResponseOutput() ContainerAccountResponseOutput
	ToContainerAccountResponseOutputWithContext(context.Context) ContainerAccountResponseOutput
}

type ContainerAccountResponseArgs struct {
	AccountName pulumi.StringPtrInput `pulumi:"accountName"`
	Password    pulumi.StringPtrInput `pulumi:"password"`
	Spn         pulumi.StringPtrInput `pulumi:"spn"`
}

func (ContainerAccountResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAccountResponse)(nil)).Elem()
}

func (i ContainerAccountResponseArgs) ToContainerAccountResponseOutput() ContainerAccountResponseOutput {
	return i.ToContainerAccountResponseOutputWithContext(context.Background())
}

func (i ContainerAccountResponseArgs) ToContainerAccountResponseOutputWithContext(ctx context.Context) ContainerAccountResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAccountResponseOutput)
}





type ContainerAccountResponseArrayInput interface {
	pulumi.Input

	ToContainerAccountResponseArrayOutput() ContainerAccountResponseArrayOutput
	ToContainerAccountResponseArrayOutputWithContext(context.Context) ContainerAccountResponseArrayOutput
}

type ContainerAccountResponseArray []ContainerAccountResponseInput

func (ContainerAccountResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerAccountResponse)(nil)).Elem()
}

func (i ContainerAccountResponseArray) ToContainerAccountResponseArrayOutput() ContainerAccountResponseArrayOutput {
	return i.ToContainerAccountResponseArrayOutputWithContext(context.Background())
}

func (i ContainerAccountResponseArray) ToContainerAccountResponseArrayOutputWithContext(ctx context.Context) ContainerAccountResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAccountResponseArrayOutput)
}

type ContainerAccountResponseOutput struct{ *pulumi.OutputState }

func (ContainerAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAccountResponse)(nil)).Elem()
}

func (o ContainerAccountResponseOutput) ToContainerAccountResponseOutput() ContainerAccountResponseOutput {
	return o
}

func (o ContainerAccountResponseOutput) ToContainerAccountResponseOutputWithContext(ctx context.Context) ContainerAccountResponseOutput {
	return o
}

func (o ContainerAccountResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAccountResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o ContainerAccountResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAccountResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ContainerAccountResponseOutput) Spn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAccountResponse) *string { return v.Spn }).(pulumi.StringPtrOutput)
}

type ContainerAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerAccountResponse)(nil)).Elem()
}

func (o ContainerAccountResponseArrayOutput) ToContainerAccountResponseArrayOutput() ContainerAccountResponseArrayOutput {
	return o
}

func (o ContainerAccountResponseArrayOutput) ToContainerAccountResponseArrayOutputWithContext(ctx context.Context) ContainerAccountResponseArrayOutput {
	return o
}

func (o ContainerAccountResponseArrayOutput) Index(i pulumi.IntInput) ContainerAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerAccountResponse {
		return vs[0].([]ContainerAccountResponse)[vs[1].(int)]
	}).(ContainerAccountResponseOutput)
}

type DomainSecuritySettings struct {
	NtlmV1                *string `pulumi:"ntlmV1"`
	SyncKerberosPasswords *string `pulumi:"syncKerberosPasswords"`
	SyncNtlmPasswords     *string `pulumi:"syncNtlmPasswords"`
	SyncOnPremPasswords   *string `pulumi:"syncOnPremPasswords"`
	TlsV1                 *string `pulumi:"tlsV1"`
}


func (val *DomainSecuritySettings) Defaults() *DomainSecuritySettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.NtlmV1) {
		ntlmV1_ := "Enabled"
		tmp.NtlmV1 = &ntlmV1_
	}
	if isZero(tmp.SyncKerberosPasswords) {
		syncKerberosPasswords_ := "Enabled"
		tmp.SyncKerberosPasswords = &syncKerberosPasswords_
	}
	if isZero(tmp.SyncNtlmPasswords) {
		syncNtlmPasswords_ := "Enabled"
		tmp.SyncNtlmPasswords = &syncNtlmPasswords_
	}
	if isZero(tmp.SyncOnPremPasswords) {
		syncOnPremPasswords_ := "Enabled"
		tmp.SyncOnPremPasswords = &syncOnPremPasswords_
	}
	if isZero(tmp.TlsV1) {
		tlsV1_ := "Enabled"
		tmp.TlsV1 = &tlsV1_
	}
	return &tmp
}





type DomainSecuritySettingsInput interface {
	pulumi.Input

	ToDomainSecuritySettingsOutput() DomainSecuritySettingsOutput
	ToDomainSecuritySettingsOutputWithContext(context.Context) DomainSecuritySettingsOutput
}

type DomainSecuritySettingsArgs struct {
	NtlmV1                pulumi.StringPtrInput `pulumi:"ntlmV1"`
	SyncKerberosPasswords pulumi.StringPtrInput `pulumi:"syncKerberosPasswords"`
	SyncNtlmPasswords     pulumi.StringPtrInput `pulumi:"syncNtlmPasswords"`
	SyncOnPremPasswords   pulumi.StringPtrInput `pulumi:"syncOnPremPasswords"`
	TlsV1                 pulumi.StringPtrInput `pulumi:"tlsV1"`
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

func (o DomainSecuritySettingsOutput) SyncKerberosPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettings) *string { return v.SyncKerberosPasswords }).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsOutput) SyncNtlmPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettings) *string { return v.SyncNtlmPasswords }).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsOutput) SyncOnPremPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettings) *string { return v.SyncOnPremPasswords }).(pulumi.StringPtrOutput)
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

func (o DomainSecuritySettingsPtrOutput) SyncKerberosPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainSecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.SyncKerberosPasswords
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

func (o DomainSecuritySettingsPtrOutput) SyncOnPremPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainSecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.SyncOnPremPasswords
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
	NtlmV1                *string `pulumi:"ntlmV1"`
	SyncKerberosPasswords *string `pulumi:"syncKerberosPasswords"`
	SyncNtlmPasswords     *string `pulumi:"syncNtlmPasswords"`
	SyncOnPremPasswords   *string `pulumi:"syncOnPremPasswords"`
	TlsV1                 *string `pulumi:"tlsV1"`
}


func (val *DomainSecuritySettingsResponse) Defaults() *DomainSecuritySettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.NtlmV1) {
		ntlmV1_ := "Enabled"
		tmp.NtlmV1 = &ntlmV1_
	}
	if isZero(tmp.SyncKerberosPasswords) {
		syncKerberosPasswords_ := "Enabled"
		tmp.SyncKerberosPasswords = &syncKerberosPasswords_
	}
	if isZero(tmp.SyncNtlmPasswords) {
		syncNtlmPasswords_ := "Enabled"
		tmp.SyncNtlmPasswords = &syncNtlmPasswords_
	}
	if isZero(tmp.SyncOnPremPasswords) {
		syncOnPremPasswords_ := "Enabled"
		tmp.SyncOnPremPasswords = &syncOnPremPasswords_
	}
	if isZero(tmp.TlsV1) {
		tlsV1_ := "Enabled"
		tmp.TlsV1 = &tlsV1_
	}
	return &tmp
}





type DomainSecuritySettingsResponseInput interface {
	pulumi.Input

	ToDomainSecuritySettingsResponseOutput() DomainSecuritySettingsResponseOutput
	ToDomainSecuritySettingsResponseOutputWithContext(context.Context) DomainSecuritySettingsResponseOutput
}

type DomainSecuritySettingsResponseArgs struct {
	NtlmV1                pulumi.StringPtrInput `pulumi:"ntlmV1"`
	SyncKerberosPasswords pulumi.StringPtrInput `pulumi:"syncKerberosPasswords"`
	SyncNtlmPasswords     pulumi.StringPtrInput `pulumi:"syncNtlmPasswords"`
	SyncOnPremPasswords   pulumi.StringPtrInput `pulumi:"syncOnPremPasswords"`
	TlsV1                 pulumi.StringPtrInput `pulumi:"tlsV1"`
}

func (DomainSecuritySettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainSecuritySettingsResponse)(nil)).Elem()
}

func (i DomainSecuritySettingsResponseArgs) ToDomainSecuritySettingsResponseOutput() DomainSecuritySettingsResponseOutput {
	return i.ToDomainSecuritySettingsResponseOutputWithContext(context.Background())
}

func (i DomainSecuritySettingsResponseArgs) ToDomainSecuritySettingsResponseOutputWithContext(ctx context.Context) DomainSecuritySettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSecuritySettingsResponseOutput)
}

func (i DomainSecuritySettingsResponseArgs) ToDomainSecuritySettingsResponsePtrOutput() DomainSecuritySettingsResponsePtrOutput {
	return i.ToDomainSecuritySettingsResponsePtrOutputWithContext(context.Background())
}

func (i DomainSecuritySettingsResponseArgs) ToDomainSecuritySettingsResponsePtrOutputWithContext(ctx context.Context) DomainSecuritySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSecuritySettingsResponseOutput).ToDomainSecuritySettingsResponsePtrOutputWithContext(ctx)
}









type DomainSecuritySettingsResponsePtrInput interface {
	pulumi.Input

	ToDomainSecuritySettingsResponsePtrOutput() DomainSecuritySettingsResponsePtrOutput
	ToDomainSecuritySettingsResponsePtrOutputWithContext(context.Context) DomainSecuritySettingsResponsePtrOutput
}

type domainSecuritySettingsResponsePtrType DomainSecuritySettingsResponseArgs

func DomainSecuritySettingsResponsePtr(v *DomainSecuritySettingsResponseArgs) DomainSecuritySettingsResponsePtrInput {
	return (*domainSecuritySettingsResponsePtrType)(v)
}

func (*domainSecuritySettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainSecuritySettingsResponse)(nil)).Elem()
}

func (i *domainSecuritySettingsResponsePtrType) ToDomainSecuritySettingsResponsePtrOutput() DomainSecuritySettingsResponsePtrOutput {
	return i.ToDomainSecuritySettingsResponsePtrOutputWithContext(context.Background())
}

func (i *domainSecuritySettingsResponsePtrType) ToDomainSecuritySettingsResponsePtrOutputWithContext(ctx context.Context) DomainSecuritySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSecuritySettingsResponsePtrOutput)
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

func (o DomainSecuritySettingsResponseOutput) ToDomainSecuritySettingsResponsePtrOutput() DomainSecuritySettingsResponsePtrOutput {
	return o.ToDomainSecuritySettingsResponsePtrOutputWithContext(context.Background())
}

func (o DomainSecuritySettingsResponseOutput) ToDomainSecuritySettingsResponsePtrOutputWithContext(ctx context.Context) DomainSecuritySettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainSecuritySettingsResponse) *DomainSecuritySettingsResponse {
		return &v
	}).(DomainSecuritySettingsResponsePtrOutput)
}

func (o DomainSecuritySettingsResponseOutput) NtlmV1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettingsResponse) *string { return v.NtlmV1 }).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsResponseOutput) SyncKerberosPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettingsResponse) *string { return v.SyncKerberosPasswords }).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsResponseOutput) SyncNtlmPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettingsResponse) *string { return v.SyncNtlmPasswords }).(pulumi.StringPtrOutput)
}

func (o DomainSecuritySettingsResponseOutput) SyncOnPremPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSecuritySettingsResponse) *string { return v.SyncOnPremPasswords }).(pulumi.StringPtrOutput)
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

func (o DomainSecuritySettingsResponsePtrOutput) SyncKerberosPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainSecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SyncKerberosPasswords
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

func (o DomainSecuritySettingsResponsePtrOutput) SyncOnPremPasswords() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainSecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SyncOnPremPasswords
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

type ForestTrust struct {
	FriendlyName      *string `pulumi:"friendlyName"`
	RemoteDnsIps      *string `pulumi:"remoteDnsIps"`
	TrustDirection    *string `pulumi:"trustDirection"`
	TrustPassword     *string `pulumi:"trustPassword"`
	TrustedDomainFqdn *string `pulumi:"trustedDomainFqdn"`
}





type ForestTrustInput interface {
	pulumi.Input

	ToForestTrustOutput() ForestTrustOutput
	ToForestTrustOutputWithContext(context.Context) ForestTrustOutput
}

type ForestTrustArgs struct {
	FriendlyName      pulumi.StringPtrInput `pulumi:"friendlyName"`
	RemoteDnsIps      pulumi.StringPtrInput `pulumi:"remoteDnsIps"`
	TrustDirection    pulumi.StringPtrInput `pulumi:"trustDirection"`
	TrustPassword     pulumi.StringPtrInput `pulumi:"trustPassword"`
	TrustedDomainFqdn pulumi.StringPtrInput `pulumi:"trustedDomainFqdn"`
}

func (ForestTrustArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ForestTrust)(nil)).Elem()
}

func (i ForestTrustArgs) ToForestTrustOutput() ForestTrustOutput {
	return i.ToForestTrustOutputWithContext(context.Background())
}

func (i ForestTrustArgs) ToForestTrustOutputWithContext(ctx context.Context) ForestTrustOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForestTrustOutput)
}





type ForestTrustArrayInput interface {
	pulumi.Input

	ToForestTrustArrayOutput() ForestTrustArrayOutput
	ToForestTrustArrayOutputWithContext(context.Context) ForestTrustArrayOutput
}

type ForestTrustArray []ForestTrustInput

func (ForestTrustArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ForestTrust)(nil)).Elem()
}

func (i ForestTrustArray) ToForestTrustArrayOutput() ForestTrustArrayOutput {
	return i.ToForestTrustArrayOutputWithContext(context.Background())
}

func (i ForestTrustArray) ToForestTrustArrayOutputWithContext(ctx context.Context) ForestTrustArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForestTrustArrayOutput)
}

type ForestTrustOutput struct{ *pulumi.OutputState }

func (ForestTrustOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForestTrust)(nil)).Elem()
}

func (o ForestTrustOutput) ToForestTrustOutput() ForestTrustOutput {
	return o
}

func (o ForestTrustOutput) ToForestTrustOutputWithContext(ctx context.Context) ForestTrustOutput {
	return o
}

func (o ForestTrustOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForestTrust) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o ForestTrustOutput) RemoteDnsIps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForestTrust) *string { return v.RemoteDnsIps }).(pulumi.StringPtrOutput)
}

func (o ForestTrustOutput) TrustDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForestTrust) *string { return v.TrustDirection }).(pulumi.StringPtrOutput)
}

func (o ForestTrustOutput) TrustPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForestTrust) *string { return v.TrustPassword }).(pulumi.StringPtrOutput)
}

func (o ForestTrustOutput) TrustedDomainFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForestTrust) *string { return v.TrustedDomainFqdn }).(pulumi.StringPtrOutput)
}

type ForestTrustArrayOutput struct{ *pulumi.OutputState }

func (ForestTrustArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ForestTrust)(nil)).Elem()
}

func (o ForestTrustArrayOutput) ToForestTrustArrayOutput() ForestTrustArrayOutput {
	return o
}

func (o ForestTrustArrayOutput) ToForestTrustArrayOutputWithContext(ctx context.Context) ForestTrustArrayOutput {
	return o
}

func (o ForestTrustArrayOutput) Index(i pulumi.IntInput) ForestTrustOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ForestTrust {
		return vs[0].([]ForestTrust)[vs[1].(int)]
	}).(ForestTrustOutput)
}

type ForestTrustResponse struct {
	FriendlyName      *string `pulumi:"friendlyName"`
	RemoteDnsIps      *string `pulumi:"remoteDnsIps"`
	TrustDirection    *string `pulumi:"trustDirection"`
	TrustPassword     *string `pulumi:"trustPassword"`
	TrustedDomainFqdn *string `pulumi:"trustedDomainFqdn"`
}





type ForestTrustResponseInput interface {
	pulumi.Input

	ToForestTrustResponseOutput() ForestTrustResponseOutput
	ToForestTrustResponseOutputWithContext(context.Context) ForestTrustResponseOutput
}

type ForestTrustResponseArgs struct {
	FriendlyName      pulumi.StringPtrInput `pulumi:"friendlyName"`
	RemoteDnsIps      pulumi.StringPtrInput `pulumi:"remoteDnsIps"`
	TrustDirection    pulumi.StringPtrInput `pulumi:"trustDirection"`
	TrustPassword     pulumi.StringPtrInput `pulumi:"trustPassword"`
	TrustedDomainFqdn pulumi.StringPtrInput `pulumi:"trustedDomainFqdn"`
}

func (ForestTrustResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ForestTrustResponse)(nil)).Elem()
}

func (i ForestTrustResponseArgs) ToForestTrustResponseOutput() ForestTrustResponseOutput {
	return i.ToForestTrustResponseOutputWithContext(context.Background())
}

func (i ForestTrustResponseArgs) ToForestTrustResponseOutputWithContext(ctx context.Context) ForestTrustResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForestTrustResponseOutput)
}





type ForestTrustResponseArrayInput interface {
	pulumi.Input

	ToForestTrustResponseArrayOutput() ForestTrustResponseArrayOutput
	ToForestTrustResponseArrayOutputWithContext(context.Context) ForestTrustResponseArrayOutput
}

type ForestTrustResponseArray []ForestTrustResponseInput

func (ForestTrustResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ForestTrustResponse)(nil)).Elem()
}

func (i ForestTrustResponseArray) ToForestTrustResponseArrayOutput() ForestTrustResponseArrayOutput {
	return i.ToForestTrustResponseArrayOutputWithContext(context.Background())
}

func (i ForestTrustResponseArray) ToForestTrustResponseArrayOutputWithContext(ctx context.Context) ForestTrustResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForestTrustResponseArrayOutput)
}

type ForestTrustResponseOutput struct{ *pulumi.OutputState }

func (ForestTrustResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForestTrustResponse)(nil)).Elem()
}

func (o ForestTrustResponseOutput) ToForestTrustResponseOutput() ForestTrustResponseOutput {
	return o
}

func (o ForestTrustResponseOutput) ToForestTrustResponseOutputWithContext(ctx context.Context) ForestTrustResponseOutput {
	return o
}

func (o ForestTrustResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForestTrustResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o ForestTrustResponseOutput) RemoteDnsIps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForestTrustResponse) *string { return v.RemoteDnsIps }).(pulumi.StringPtrOutput)
}

func (o ForestTrustResponseOutput) TrustDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForestTrustResponse) *string { return v.TrustDirection }).(pulumi.StringPtrOutput)
}

func (o ForestTrustResponseOutput) TrustPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForestTrustResponse) *string { return v.TrustPassword }).(pulumi.StringPtrOutput)
}

func (o ForestTrustResponseOutput) TrustedDomainFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForestTrustResponse) *string { return v.TrustedDomainFqdn }).(pulumi.StringPtrOutput)
}

type ForestTrustResponseArrayOutput struct{ *pulumi.OutputState }

func (ForestTrustResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ForestTrustResponse)(nil)).Elem()
}

func (o ForestTrustResponseArrayOutput) ToForestTrustResponseArrayOutput() ForestTrustResponseArrayOutput {
	return o
}

func (o ForestTrustResponseArrayOutput) ToForestTrustResponseArrayOutputWithContext(ctx context.Context) ForestTrustResponseArrayOutput {
	return o
}

func (o ForestTrustResponseArrayOutput) Index(i pulumi.IntInput) ForestTrustResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ForestTrustResponse {
		return vs[0].([]ForestTrustResponse)[vs[1].(int)]
	}).(ForestTrustResponseOutput)
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





type HealthAlertResponseInput interface {
	pulumi.Input

	ToHealthAlertResponseOutput() HealthAlertResponseOutput
	ToHealthAlertResponseOutputWithContext(context.Context) HealthAlertResponseOutput
}

type HealthAlertResponseArgs struct {
	Id            pulumi.StringInput `pulumi:"id"`
	Issue         pulumi.StringInput `pulumi:"issue"`
	LastDetected  pulumi.StringInput `pulumi:"lastDetected"`
	Name          pulumi.StringInput `pulumi:"name"`
	Raised        pulumi.StringInput `pulumi:"raised"`
	ResolutionUri pulumi.StringInput `pulumi:"resolutionUri"`
	Severity      pulumi.StringInput `pulumi:"severity"`
}

func (HealthAlertResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertResponse)(nil)).Elem()
}

func (i HealthAlertResponseArgs) ToHealthAlertResponseOutput() HealthAlertResponseOutput {
	return i.ToHealthAlertResponseOutputWithContext(context.Background())
}

func (i HealthAlertResponseArgs) ToHealthAlertResponseOutputWithContext(ctx context.Context) HealthAlertResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertResponseOutput)
}





type HealthAlertResponseArrayInput interface {
	pulumi.Input

	ToHealthAlertResponseArrayOutput() HealthAlertResponseArrayOutput
	ToHealthAlertResponseArrayOutputWithContext(context.Context) HealthAlertResponseArrayOutput
}

type HealthAlertResponseArray []HealthAlertResponseInput

func (HealthAlertResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertResponse)(nil)).Elem()
}

func (i HealthAlertResponseArray) ToHealthAlertResponseArrayOutput() HealthAlertResponseArrayOutput {
	return i.ToHealthAlertResponseArrayOutputWithContext(context.Background())
}

func (i HealthAlertResponseArray) ToHealthAlertResponseArrayOutputWithContext(ctx context.Context) HealthAlertResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertResponseArrayOutput)
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





type HealthMonitorResponseInput interface {
	pulumi.Input

	ToHealthMonitorResponseOutput() HealthMonitorResponseOutput
	ToHealthMonitorResponseOutputWithContext(context.Context) HealthMonitorResponseOutput
}

type HealthMonitorResponseArgs struct {
	Details pulumi.StringInput `pulumi:"details"`
	Id      pulumi.StringInput `pulumi:"id"`
	Name    pulumi.StringInput `pulumi:"name"`
}

func (HealthMonitorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthMonitorResponse)(nil)).Elem()
}

func (i HealthMonitorResponseArgs) ToHealthMonitorResponseOutput() HealthMonitorResponseOutput {
	return i.ToHealthMonitorResponseOutputWithContext(context.Background())
}

func (i HealthMonitorResponseArgs) ToHealthMonitorResponseOutputWithContext(ctx context.Context) HealthMonitorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthMonitorResponseOutput)
}





type HealthMonitorResponseArrayInput interface {
	pulumi.Input

	ToHealthMonitorResponseArrayOutput() HealthMonitorResponseArrayOutput
	ToHealthMonitorResponseArrayOutputWithContext(context.Context) HealthMonitorResponseArrayOutput
}

type HealthMonitorResponseArray []HealthMonitorResponseInput

func (HealthMonitorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthMonitorResponse)(nil)).Elem()
}

func (i HealthMonitorResponseArray) ToHealthMonitorResponseArrayOutput() HealthMonitorResponseArrayOutput {
	return i.ToHealthMonitorResponseArrayOutputWithContext(context.Background())
}

func (i HealthMonitorResponseArray) ToHealthMonitorResponseArrayOutputWithContext(ctx context.Context) HealthMonitorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthMonitorResponseArrayOutput)
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


func (val *LdapsSettings) Defaults() *LdapsSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExternalAccess) {
		externalAccess_ := "Disabled"
		tmp.ExternalAccess = &externalAccess_
	}
	if isZero(tmp.Ldaps) {
		ldaps_ := "Disabled"
		tmp.Ldaps = &ldaps_
	}
	return &tmp
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


func (val *LdapsSettingsResponse) Defaults() *LdapsSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExternalAccess) {
		externalAccess_ := "Disabled"
		tmp.ExternalAccess = &externalAccess_
	}
	if isZero(tmp.Ldaps) {
		ldaps_ := "Disabled"
		tmp.Ldaps = &ldaps_
	}
	return &tmp
}





type LdapsSettingsResponseInput interface {
	pulumi.Input

	ToLdapsSettingsResponseOutput() LdapsSettingsResponseOutput
	ToLdapsSettingsResponseOutputWithContext(context.Context) LdapsSettingsResponseOutput
}

type LdapsSettingsResponseArgs struct {
	CertificateNotAfter     pulumi.StringInput    `pulumi:"certificateNotAfter"`
	CertificateThumbprint   pulumi.StringInput    `pulumi:"certificateThumbprint"`
	ExternalAccess          pulumi.StringPtrInput `pulumi:"externalAccess"`
	ExternalAccessIpAddress pulumi.StringInput    `pulumi:"externalAccessIpAddress"`
	Ldaps                   pulumi.StringPtrInput `pulumi:"ldaps"`
	PfxCertificate          pulumi.StringPtrInput `pulumi:"pfxCertificate"`
	PfxCertificatePassword  pulumi.StringPtrInput `pulumi:"pfxCertificatePassword"`
	PublicCertificate       pulumi.StringInput    `pulumi:"publicCertificate"`
}

func (LdapsSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LdapsSettingsResponse)(nil)).Elem()
}

func (i LdapsSettingsResponseArgs) ToLdapsSettingsResponseOutput() LdapsSettingsResponseOutput {
	return i.ToLdapsSettingsResponseOutputWithContext(context.Background())
}

func (i LdapsSettingsResponseArgs) ToLdapsSettingsResponseOutputWithContext(ctx context.Context) LdapsSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapsSettingsResponseOutput)
}

func (i LdapsSettingsResponseArgs) ToLdapsSettingsResponsePtrOutput() LdapsSettingsResponsePtrOutput {
	return i.ToLdapsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i LdapsSettingsResponseArgs) ToLdapsSettingsResponsePtrOutputWithContext(ctx context.Context) LdapsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapsSettingsResponseOutput).ToLdapsSettingsResponsePtrOutputWithContext(ctx)
}









type LdapsSettingsResponsePtrInput interface {
	pulumi.Input

	ToLdapsSettingsResponsePtrOutput() LdapsSettingsResponsePtrOutput
	ToLdapsSettingsResponsePtrOutputWithContext(context.Context) LdapsSettingsResponsePtrOutput
}

type ldapsSettingsResponsePtrType LdapsSettingsResponseArgs

func LdapsSettingsResponsePtr(v *LdapsSettingsResponseArgs) LdapsSettingsResponsePtrInput {
	return (*ldapsSettingsResponsePtrType)(v)
}

func (*ldapsSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapsSettingsResponse)(nil)).Elem()
}

func (i *ldapsSettingsResponsePtrType) ToLdapsSettingsResponsePtrOutput() LdapsSettingsResponsePtrOutput {
	return i.ToLdapsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *ldapsSettingsResponsePtrType) ToLdapsSettingsResponsePtrOutputWithContext(ctx context.Context) LdapsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapsSettingsResponsePtrOutput)
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

func (o LdapsSettingsResponseOutput) ToLdapsSettingsResponsePtrOutput() LdapsSettingsResponsePtrOutput {
	return o.ToLdapsSettingsResponsePtrOutputWithContext(context.Background())
}

func (o LdapsSettingsResponseOutput) ToLdapsSettingsResponsePtrOutputWithContext(ctx context.Context) LdapsSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LdapsSettingsResponse) *LdapsSettingsResponse {
		return &v
	}).(LdapsSettingsResponsePtrOutput)
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

type MigrationProgressResponse struct {
	CompletionPercentage *float64 `pulumi:"completionPercentage"`
	ProgressMessage      *string  `pulumi:"progressMessage"`
}





type MigrationProgressResponseInput interface {
	pulumi.Input

	ToMigrationProgressResponseOutput() MigrationProgressResponseOutput
	ToMigrationProgressResponseOutputWithContext(context.Context) MigrationProgressResponseOutput
}

type MigrationProgressResponseArgs struct {
	CompletionPercentage pulumi.Float64PtrInput `pulumi:"completionPercentage"`
	ProgressMessage      pulumi.StringPtrInput  `pulumi:"progressMessage"`
}

func (MigrationProgressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationProgressResponse)(nil)).Elem()
}

func (i MigrationProgressResponseArgs) ToMigrationProgressResponseOutput() MigrationProgressResponseOutput {
	return i.ToMigrationProgressResponseOutputWithContext(context.Background())
}

func (i MigrationProgressResponseArgs) ToMigrationProgressResponseOutputWithContext(ctx context.Context) MigrationProgressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationProgressResponseOutput)
}

func (i MigrationProgressResponseArgs) ToMigrationProgressResponsePtrOutput() MigrationProgressResponsePtrOutput {
	return i.ToMigrationProgressResponsePtrOutputWithContext(context.Background())
}

func (i MigrationProgressResponseArgs) ToMigrationProgressResponsePtrOutputWithContext(ctx context.Context) MigrationProgressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationProgressResponseOutput).ToMigrationProgressResponsePtrOutputWithContext(ctx)
}









type MigrationProgressResponsePtrInput interface {
	pulumi.Input

	ToMigrationProgressResponsePtrOutput() MigrationProgressResponsePtrOutput
	ToMigrationProgressResponsePtrOutputWithContext(context.Context) MigrationProgressResponsePtrOutput
}

type migrationProgressResponsePtrType MigrationProgressResponseArgs

func MigrationProgressResponsePtr(v *MigrationProgressResponseArgs) MigrationProgressResponsePtrInput {
	return (*migrationProgressResponsePtrType)(v)
}

func (*migrationProgressResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationProgressResponse)(nil)).Elem()
}

func (i *migrationProgressResponsePtrType) ToMigrationProgressResponsePtrOutput() MigrationProgressResponsePtrOutput {
	return i.ToMigrationProgressResponsePtrOutputWithContext(context.Background())
}

func (i *migrationProgressResponsePtrType) ToMigrationProgressResponsePtrOutputWithContext(ctx context.Context) MigrationProgressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationProgressResponsePtrOutput)
}

type MigrationProgressResponseOutput struct{ *pulumi.OutputState }

func (MigrationProgressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationProgressResponse)(nil)).Elem()
}

func (o MigrationProgressResponseOutput) ToMigrationProgressResponseOutput() MigrationProgressResponseOutput {
	return o
}

func (o MigrationProgressResponseOutput) ToMigrationProgressResponseOutputWithContext(ctx context.Context) MigrationProgressResponseOutput {
	return o
}

func (o MigrationProgressResponseOutput) ToMigrationProgressResponsePtrOutput() MigrationProgressResponsePtrOutput {
	return o.ToMigrationProgressResponsePtrOutputWithContext(context.Background())
}

func (o MigrationProgressResponseOutput) ToMigrationProgressResponsePtrOutputWithContext(ctx context.Context) MigrationProgressResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrationProgressResponse) *MigrationProgressResponse {
		return &v
	}).(MigrationProgressResponsePtrOutput)
}

func (o MigrationProgressResponseOutput) CompletionPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MigrationProgressResponse) *float64 { return v.CompletionPercentage }).(pulumi.Float64PtrOutput)
}

func (o MigrationProgressResponseOutput) ProgressMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrationProgressResponse) *string { return v.ProgressMessage }).(pulumi.StringPtrOutput)
}

type MigrationProgressResponsePtrOutput struct{ *pulumi.OutputState }

func (MigrationProgressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationProgressResponse)(nil)).Elem()
}

func (o MigrationProgressResponsePtrOutput) ToMigrationProgressResponsePtrOutput() MigrationProgressResponsePtrOutput {
	return o
}

func (o MigrationProgressResponsePtrOutput) ToMigrationProgressResponsePtrOutputWithContext(ctx context.Context) MigrationProgressResponsePtrOutput {
	return o
}

func (o MigrationProgressResponsePtrOutput) Elem() MigrationProgressResponseOutput {
	return o.ApplyT(func(v *MigrationProgressResponse) MigrationProgressResponse {
		if v != nil {
			return *v
		}
		var ret MigrationProgressResponse
		return ret
	}).(MigrationProgressResponseOutput)
}

func (o MigrationProgressResponsePtrOutput) CompletionPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MigrationProgressResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.CompletionPercentage
	}).(pulumi.Float64PtrOutput)
}

func (o MigrationProgressResponsePtrOutput) ProgressMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationProgressResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProgressMessage
	}).(pulumi.StringPtrOutput)
}

type MigrationPropertiesResponse struct {
	MigrationProgress MigrationProgressResponse `pulumi:"migrationProgress"`
	OldSubnetId       string                    `pulumi:"oldSubnetId"`
	OldVnetSiteId     string                    `pulumi:"oldVnetSiteId"`
}





type MigrationPropertiesResponseInput interface {
	pulumi.Input

	ToMigrationPropertiesResponseOutput() MigrationPropertiesResponseOutput
	ToMigrationPropertiesResponseOutputWithContext(context.Context) MigrationPropertiesResponseOutput
}

type MigrationPropertiesResponseArgs struct {
	MigrationProgress MigrationProgressResponseInput `pulumi:"migrationProgress"`
	OldSubnetId       pulumi.StringInput             `pulumi:"oldSubnetId"`
	OldVnetSiteId     pulumi.StringInput             `pulumi:"oldVnetSiteId"`
}

func (MigrationPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationPropertiesResponse)(nil)).Elem()
}

func (i MigrationPropertiesResponseArgs) ToMigrationPropertiesResponseOutput() MigrationPropertiesResponseOutput {
	return i.ToMigrationPropertiesResponseOutputWithContext(context.Background())
}

func (i MigrationPropertiesResponseArgs) ToMigrationPropertiesResponseOutputWithContext(ctx context.Context) MigrationPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationPropertiesResponseOutput)
}

func (i MigrationPropertiesResponseArgs) ToMigrationPropertiesResponsePtrOutput() MigrationPropertiesResponsePtrOutput {
	return i.ToMigrationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MigrationPropertiesResponseArgs) ToMigrationPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationPropertiesResponseOutput).ToMigrationPropertiesResponsePtrOutputWithContext(ctx)
}









type MigrationPropertiesResponsePtrInput interface {
	pulumi.Input

	ToMigrationPropertiesResponsePtrOutput() MigrationPropertiesResponsePtrOutput
	ToMigrationPropertiesResponsePtrOutputWithContext(context.Context) MigrationPropertiesResponsePtrOutput
}

type migrationPropertiesResponsePtrType MigrationPropertiesResponseArgs

func MigrationPropertiesResponsePtr(v *MigrationPropertiesResponseArgs) MigrationPropertiesResponsePtrInput {
	return (*migrationPropertiesResponsePtrType)(v)
}

func (*migrationPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationPropertiesResponse)(nil)).Elem()
}

func (i *migrationPropertiesResponsePtrType) ToMigrationPropertiesResponsePtrOutput() MigrationPropertiesResponsePtrOutput {
	return i.ToMigrationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *migrationPropertiesResponsePtrType) ToMigrationPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationPropertiesResponsePtrOutput)
}

type MigrationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MigrationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationPropertiesResponse)(nil)).Elem()
}

func (o MigrationPropertiesResponseOutput) ToMigrationPropertiesResponseOutput() MigrationPropertiesResponseOutput {
	return o
}

func (o MigrationPropertiesResponseOutput) ToMigrationPropertiesResponseOutputWithContext(ctx context.Context) MigrationPropertiesResponseOutput {
	return o
}

func (o MigrationPropertiesResponseOutput) ToMigrationPropertiesResponsePtrOutput() MigrationPropertiesResponsePtrOutput {
	return o.ToMigrationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MigrationPropertiesResponseOutput) ToMigrationPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrationPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrationPropertiesResponse) *MigrationPropertiesResponse {
		return &v
	}).(MigrationPropertiesResponsePtrOutput)
}

func (o MigrationPropertiesResponseOutput) MigrationProgress() MigrationProgressResponseOutput {
	return o.ApplyT(func(v MigrationPropertiesResponse) MigrationProgressResponse { return v.MigrationProgress }).(MigrationProgressResponseOutput)
}

func (o MigrationPropertiesResponseOutput) OldSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationPropertiesResponse) string { return v.OldSubnetId }).(pulumi.StringOutput)
}

func (o MigrationPropertiesResponseOutput) OldVnetSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationPropertiesResponse) string { return v.OldVnetSiteId }).(pulumi.StringOutput)
}

type MigrationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MigrationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationPropertiesResponse)(nil)).Elem()
}

func (o MigrationPropertiesResponsePtrOutput) ToMigrationPropertiesResponsePtrOutput() MigrationPropertiesResponsePtrOutput {
	return o
}

func (o MigrationPropertiesResponsePtrOutput) ToMigrationPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrationPropertiesResponsePtrOutput {
	return o
}

func (o MigrationPropertiesResponsePtrOutput) Elem() MigrationPropertiesResponseOutput {
	return o.ApplyT(func(v *MigrationPropertiesResponse) MigrationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MigrationPropertiesResponse
		return ret
	}).(MigrationPropertiesResponseOutput)
}

func (o MigrationPropertiesResponsePtrOutput) MigrationProgress() MigrationProgressResponsePtrOutput {
	return o.ApplyT(func(v *MigrationPropertiesResponse) *MigrationProgressResponse {
		if v == nil {
			return nil
		}
		return &v.MigrationProgress
	}).(MigrationProgressResponsePtrOutput)
}

func (o MigrationPropertiesResponsePtrOutput) OldSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OldSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o MigrationPropertiesResponsePtrOutput) OldVnetSiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OldVnetSiteId
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





type NotificationSettingsResponseInput interface {
	pulumi.Input

	ToNotificationSettingsResponseOutput() NotificationSettingsResponseOutput
	ToNotificationSettingsResponseOutputWithContext(context.Context) NotificationSettingsResponseOutput
}

type NotificationSettingsResponseArgs struct {
	AdditionalRecipients pulumi.StringArrayInput `pulumi:"additionalRecipients"`
	NotifyDcAdmins       pulumi.StringPtrInput   `pulumi:"notifyDcAdmins"`
	NotifyGlobalAdmins   pulumi.StringPtrInput   `pulumi:"notifyGlobalAdmins"`
}

func (NotificationSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettingsResponse)(nil)).Elem()
}

func (i NotificationSettingsResponseArgs) ToNotificationSettingsResponseOutput() NotificationSettingsResponseOutput {
	return i.ToNotificationSettingsResponseOutputWithContext(context.Background())
}

func (i NotificationSettingsResponseArgs) ToNotificationSettingsResponseOutputWithContext(ctx context.Context) NotificationSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsResponseOutput)
}

func (i NotificationSettingsResponseArgs) ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput {
	return i.ToNotificationSettingsResponsePtrOutputWithContext(context.Background())
}

func (i NotificationSettingsResponseArgs) ToNotificationSettingsResponsePtrOutputWithContext(ctx context.Context) NotificationSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsResponseOutput).ToNotificationSettingsResponsePtrOutputWithContext(ctx)
}









type NotificationSettingsResponsePtrInput interface {
	pulumi.Input

	ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput
	ToNotificationSettingsResponsePtrOutputWithContext(context.Context) NotificationSettingsResponsePtrOutput
}

type notificationSettingsResponsePtrType NotificationSettingsResponseArgs

func NotificationSettingsResponsePtr(v *NotificationSettingsResponseArgs) NotificationSettingsResponsePtrInput {
	return (*notificationSettingsResponsePtrType)(v)
}

func (*notificationSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettingsResponse)(nil)).Elem()
}

func (i *notificationSettingsResponsePtrType) ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput {
	return i.ToNotificationSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *notificationSettingsResponsePtrType) ToNotificationSettingsResponsePtrOutputWithContext(ctx context.Context) NotificationSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsResponsePtrOutput)
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

func (o NotificationSettingsResponseOutput) ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput {
	return o.ToNotificationSettingsResponsePtrOutputWithContext(context.Background())
}

func (o NotificationSettingsResponseOutput) ToNotificationSettingsResponsePtrOutputWithContext(ctx context.Context) NotificationSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationSettingsResponse) *NotificationSettingsResponse {
		return &v
	}).(NotificationSettingsResponsePtrOutput)
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

type ResourceForestSettings struct {
	ResourceForest *string       `pulumi:"resourceForest"`
	Settings       []ForestTrust `pulumi:"settings"`
}





type ResourceForestSettingsInput interface {
	pulumi.Input

	ToResourceForestSettingsOutput() ResourceForestSettingsOutput
	ToResourceForestSettingsOutputWithContext(context.Context) ResourceForestSettingsOutput
}

type ResourceForestSettingsArgs struct {
	ResourceForest pulumi.StringPtrInput `pulumi:"resourceForest"`
	Settings       ForestTrustArrayInput `pulumi:"settings"`
}

func (ResourceForestSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceForestSettings)(nil)).Elem()
}

func (i ResourceForestSettingsArgs) ToResourceForestSettingsOutput() ResourceForestSettingsOutput {
	return i.ToResourceForestSettingsOutputWithContext(context.Background())
}

func (i ResourceForestSettingsArgs) ToResourceForestSettingsOutputWithContext(ctx context.Context) ResourceForestSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceForestSettingsOutput)
}

func (i ResourceForestSettingsArgs) ToResourceForestSettingsPtrOutput() ResourceForestSettingsPtrOutput {
	return i.ToResourceForestSettingsPtrOutputWithContext(context.Background())
}

func (i ResourceForestSettingsArgs) ToResourceForestSettingsPtrOutputWithContext(ctx context.Context) ResourceForestSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceForestSettingsOutput).ToResourceForestSettingsPtrOutputWithContext(ctx)
}









type ResourceForestSettingsPtrInput interface {
	pulumi.Input

	ToResourceForestSettingsPtrOutput() ResourceForestSettingsPtrOutput
	ToResourceForestSettingsPtrOutputWithContext(context.Context) ResourceForestSettingsPtrOutput
}

type resourceForestSettingsPtrType ResourceForestSettingsArgs

func ResourceForestSettingsPtr(v *ResourceForestSettingsArgs) ResourceForestSettingsPtrInput {
	return (*resourceForestSettingsPtrType)(v)
}

func (*resourceForestSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceForestSettings)(nil)).Elem()
}

func (i *resourceForestSettingsPtrType) ToResourceForestSettingsPtrOutput() ResourceForestSettingsPtrOutput {
	return i.ToResourceForestSettingsPtrOutputWithContext(context.Background())
}

func (i *resourceForestSettingsPtrType) ToResourceForestSettingsPtrOutputWithContext(ctx context.Context) ResourceForestSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceForestSettingsPtrOutput)
}

type ResourceForestSettingsOutput struct{ *pulumi.OutputState }

func (ResourceForestSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceForestSettings)(nil)).Elem()
}

func (o ResourceForestSettingsOutput) ToResourceForestSettingsOutput() ResourceForestSettingsOutput {
	return o
}

func (o ResourceForestSettingsOutput) ToResourceForestSettingsOutputWithContext(ctx context.Context) ResourceForestSettingsOutput {
	return o
}

func (o ResourceForestSettingsOutput) ToResourceForestSettingsPtrOutput() ResourceForestSettingsPtrOutput {
	return o.ToResourceForestSettingsPtrOutputWithContext(context.Background())
}

func (o ResourceForestSettingsOutput) ToResourceForestSettingsPtrOutputWithContext(ctx context.Context) ResourceForestSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceForestSettings) *ResourceForestSettings {
		return &v
	}).(ResourceForestSettingsPtrOutput)
}

func (o ResourceForestSettingsOutput) ResourceForest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceForestSettings) *string { return v.ResourceForest }).(pulumi.StringPtrOutput)
}

func (o ResourceForestSettingsOutput) Settings() ForestTrustArrayOutput {
	return o.ApplyT(func(v ResourceForestSettings) []ForestTrust { return v.Settings }).(ForestTrustArrayOutput)
}

type ResourceForestSettingsPtrOutput struct{ *pulumi.OutputState }

func (ResourceForestSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceForestSettings)(nil)).Elem()
}

func (o ResourceForestSettingsPtrOutput) ToResourceForestSettingsPtrOutput() ResourceForestSettingsPtrOutput {
	return o
}

func (o ResourceForestSettingsPtrOutput) ToResourceForestSettingsPtrOutputWithContext(ctx context.Context) ResourceForestSettingsPtrOutput {
	return o
}

func (o ResourceForestSettingsPtrOutput) Elem() ResourceForestSettingsOutput {
	return o.ApplyT(func(v *ResourceForestSettings) ResourceForestSettings {
		if v != nil {
			return *v
		}
		var ret ResourceForestSettings
		return ret
	}).(ResourceForestSettingsOutput)
}

func (o ResourceForestSettingsPtrOutput) ResourceForest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceForestSettings) *string {
		if v == nil {
			return nil
		}
		return v.ResourceForest
	}).(pulumi.StringPtrOutput)
}

func (o ResourceForestSettingsPtrOutput) Settings() ForestTrustArrayOutput {
	return o.ApplyT(func(v *ResourceForestSettings) []ForestTrust {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(ForestTrustArrayOutput)
}

type ResourceForestSettingsResponse struct {
	ResourceForest *string               `pulumi:"resourceForest"`
	Settings       []ForestTrustResponse `pulumi:"settings"`
}





type ResourceForestSettingsResponseInput interface {
	pulumi.Input

	ToResourceForestSettingsResponseOutput() ResourceForestSettingsResponseOutput
	ToResourceForestSettingsResponseOutputWithContext(context.Context) ResourceForestSettingsResponseOutput
}

type ResourceForestSettingsResponseArgs struct {
	ResourceForest pulumi.StringPtrInput         `pulumi:"resourceForest"`
	Settings       ForestTrustResponseArrayInput `pulumi:"settings"`
}

func (ResourceForestSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceForestSettingsResponse)(nil)).Elem()
}

func (i ResourceForestSettingsResponseArgs) ToResourceForestSettingsResponseOutput() ResourceForestSettingsResponseOutput {
	return i.ToResourceForestSettingsResponseOutputWithContext(context.Background())
}

func (i ResourceForestSettingsResponseArgs) ToResourceForestSettingsResponseOutputWithContext(ctx context.Context) ResourceForestSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceForestSettingsResponseOutput)
}

func (i ResourceForestSettingsResponseArgs) ToResourceForestSettingsResponsePtrOutput() ResourceForestSettingsResponsePtrOutput {
	return i.ToResourceForestSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ResourceForestSettingsResponseArgs) ToResourceForestSettingsResponsePtrOutputWithContext(ctx context.Context) ResourceForestSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceForestSettingsResponseOutput).ToResourceForestSettingsResponsePtrOutputWithContext(ctx)
}









type ResourceForestSettingsResponsePtrInput interface {
	pulumi.Input

	ToResourceForestSettingsResponsePtrOutput() ResourceForestSettingsResponsePtrOutput
	ToResourceForestSettingsResponsePtrOutputWithContext(context.Context) ResourceForestSettingsResponsePtrOutput
}

type resourceForestSettingsResponsePtrType ResourceForestSettingsResponseArgs

func ResourceForestSettingsResponsePtr(v *ResourceForestSettingsResponseArgs) ResourceForestSettingsResponsePtrInput {
	return (*resourceForestSettingsResponsePtrType)(v)
}

func (*resourceForestSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceForestSettingsResponse)(nil)).Elem()
}

func (i *resourceForestSettingsResponsePtrType) ToResourceForestSettingsResponsePtrOutput() ResourceForestSettingsResponsePtrOutput {
	return i.ToResourceForestSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *resourceForestSettingsResponsePtrType) ToResourceForestSettingsResponsePtrOutputWithContext(ctx context.Context) ResourceForestSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceForestSettingsResponsePtrOutput)
}

type ResourceForestSettingsResponseOutput struct{ *pulumi.OutputState }

func (ResourceForestSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceForestSettingsResponse)(nil)).Elem()
}

func (o ResourceForestSettingsResponseOutput) ToResourceForestSettingsResponseOutput() ResourceForestSettingsResponseOutput {
	return o
}

func (o ResourceForestSettingsResponseOutput) ToResourceForestSettingsResponseOutputWithContext(ctx context.Context) ResourceForestSettingsResponseOutput {
	return o
}

func (o ResourceForestSettingsResponseOutput) ToResourceForestSettingsResponsePtrOutput() ResourceForestSettingsResponsePtrOutput {
	return o.ToResourceForestSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ResourceForestSettingsResponseOutput) ToResourceForestSettingsResponsePtrOutputWithContext(ctx context.Context) ResourceForestSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceForestSettingsResponse) *ResourceForestSettingsResponse {
		return &v
	}).(ResourceForestSettingsResponsePtrOutput)
}

func (o ResourceForestSettingsResponseOutput) ResourceForest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceForestSettingsResponse) *string { return v.ResourceForest }).(pulumi.StringPtrOutput)
}

func (o ResourceForestSettingsResponseOutput) Settings() ForestTrustResponseArrayOutput {
	return o.ApplyT(func(v ResourceForestSettingsResponse) []ForestTrustResponse { return v.Settings }).(ForestTrustResponseArrayOutput)
}

type ResourceForestSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceForestSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceForestSettingsResponse)(nil)).Elem()
}

func (o ResourceForestSettingsResponsePtrOutput) ToResourceForestSettingsResponsePtrOutput() ResourceForestSettingsResponsePtrOutput {
	return o
}

func (o ResourceForestSettingsResponsePtrOutput) ToResourceForestSettingsResponsePtrOutputWithContext(ctx context.Context) ResourceForestSettingsResponsePtrOutput {
	return o
}

func (o ResourceForestSettingsResponsePtrOutput) Elem() ResourceForestSettingsResponseOutput {
	return o.ApplyT(func(v *ResourceForestSettingsResponse) ResourceForestSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ResourceForestSettingsResponse
		return ret
	}).(ResourceForestSettingsResponseOutput)
}

func (o ResourceForestSettingsResponsePtrOutput) ResourceForest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceForestSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceForest
	}).(pulumi.StringPtrOutput)
}

func (o ResourceForestSettingsResponsePtrOutput) Settings() ForestTrustResponseArrayOutput {
	return o.ApplyT(func(v *ResourceForestSettingsResponse) []ForestTrustResponse {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(ForestTrustResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerAccountResponseOutput{})
	pulumi.RegisterOutputType(ContainerAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(DomainSecuritySettingsOutput{})
	pulumi.RegisterOutputType(DomainSecuritySettingsPtrOutput{})
	pulumi.RegisterOutputType(DomainSecuritySettingsResponseOutput{})
	pulumi.RegisterOutputType(DomainSecuritySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ForestTrustOutput{})
	pulumi.RegisterOutputType(ForestTrustArrayOutput{})
	pulumi.RegisterOutputType(ForestTrustResponseOutput{})
	pulumi.RegisterOutputType(ForestTrustResponseArrayOutput{})
	pulumi.RegisterOutputType(HealthAlertResponseOutput{})
	pulumi.RegisterOutputType(HealthAlertResponseArrayOutput{})
	pulumi.RegisterOutputType(HealthMonitorResponseOutput{})
	pulumi.RegisterOutputType(HealthMonitorResponseArrayOutput{})
	pulumi.RegisterOutputType(LdapsSettingsOutput{})
	pulumi.RegisterOutputType(LdapsSettingsPtrOutput{})
	pulumi.RegisterOutputType(LdapsSettingsResponseOutput{})
	pulumi.RegisterOutputType(LdapsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(MigrationProgressResponseOutput{})
	pulumi.RegisterOutputType(MigrationProgressResponsePtrOutput{})
	pulumi.RegisterOutputType(MigrationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MigrationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NotificationSettingsOutput{})
	pulumi.RegisterOutputType(NotificationSettingsPtrOutput{})
	pulumi.RegisterOutputType(NotificationSettingsResponseOutput{})
	pulumi.RegisterOutputType(NotificationSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceForestSettingsOutput{})
	pulumi.RegisterOutputType(ResourceForestSettingsPtrOutput{})
	pulumi.RegisterOutputType(ResourceForestSettingsResponseOutput{})
	pulumi.RegisterOutputType(ResourceForestSettingsResponsePtrOutput{})
}
