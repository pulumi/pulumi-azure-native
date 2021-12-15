


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountEncryption struct {
	KeySource *string `pulumi:"keySource"`
}





type AccountEncryptionInput interface {
	pulumi.Input

	ToAccountEncryptionOutput() AccountEncryptionOutput
	ToAccountEncryptionOutputWithContext(context.Context) AccountEncryptionOutput
}

type AccountEncryptionArgs struct {
	KeySource pulumi.StringPtrInput `pulumi:"keySource"`
}

func (AccountEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryption)(nil)).Elem()
}

func (i AccountEncryptionArgs) ToAccountEncryptionOutput() AccountEncryptionOutput {
	return i.ToAccountEncryptionOutputWithContext(context.Background())
}

func (i AccountEncryptionArgs) ToAccountEncryptionOutputWithContext(ctx context.Context) AccountEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionOutput)
}

func (i AccountEncryptionArgs) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return i.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (i AccountEncryptionArgs) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionOutput).ToAccountEncryptionPtrOutputWithContext(ctx)
}









type AccountEncryptionPtrInput interface {
	pulumi.Input

	ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput
	ToAccountEncryptionPtrOutputWithContext(context.Context) AccountEncryptionPtrOutput
}

type accountEncryptionPtrType AccountEncryptionArgs

func AccountEncryptionPtr(v *AccountEncryptionArgs) AccountEncryptionPtrInput {
	return (*accountEncryptionPtrType)(v)
}

func (*accountEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryption)(nil)).Elem()
}

func (i *accountEncryptionPtrType) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return i.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (i *accountEncryptionPtrType) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionPtrOutput)
}

type AccountEncryptionOutput struct{ *pulumi.OutputState }

func (AccountEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryption)(nil)).Elem()
}

func (o AccountEncryptionOutput) ToAccountEncryptionOutput() AccountEncryptionOutput {
	return o
}

func (o AccountEncryptionOutput) ToAccountEncryptionOutputWithContext(ctx context.Context) AccountEncryptionOutput {
	return o
}

func (o AccountEncryptionOutput) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return o.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (o AccountEncryptionOutput) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountEncryption) *AccountEncryption {
		return &v
	}).(AccountEncryptionPtrOutput)
}

func (o AccountEncryptionOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountEncryption) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

type AccountEncryptionPtrOutput struct{ *pulumi.OutputState }

func (AccountEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryption)(nil)).Elem()
}

func (o AccountEncryptionPtrOutput) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return o
}

func (o AccountEncryptionPtrOutput) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return o
}

func (o AccountEncryptionPtrOutput) Elem() AccountEncryptionOutput {
	return o.ApplyT(func(v *AccountEncryption) AccountEncryption {
		if v != nil {
			return *v
		}
		var ret AccountEncryption
		return ret
	}).(AccountEncryptionOutput)
}

func (o AccountEncryptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryption) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

type AccountEncryptionResponse struct {
	KeySource *string `pulumi:"keySource"`
}





type AccountEncryptionResponseInput interface {
	pulumi.Input

	ToAccountEncryptionResponseOutput() AccountEncryptionResponseOutput
	ToAccountEncryptionResponseOutputWithContext(context.Context) AccountEncryptionResponseOutput
}

type AccountEncryptionResponseArgs struct {
	KeySource pulumi.StringPtrInput `pulumi:"keySource"`
}

func (AccountEncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryptionResponse)(nil)).Elem()
}

func (i AccountEncryptionResponseArgs) ToAccountEncryptionResponseOutput() AccountEncryptionResponseOutput {
	return i.ToAccountEncryptionResponseOutputWithContext(context.Background())
}

func (i AccountEncryptionResponseArgs) ToAccountEncryptionResponseOutputWithContext(ctx context.Context) AccountEncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionResponseOutput)
}

func (i AccountEncryptionResponseArgs) ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput {
	return i.ToAccountEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i AccountEncryptionResponseArgs) ToAccountEncryptionResponsePtrOutputWithContext(ctx context.Context) AccountEncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionResponseOutput).ToAccountEncryptionResponsePtrOutputWithContext(ctx)
}









type AccountEncryptionResponsePtrInput interface {
	pulumi.Input

	ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput
	ToAccountEncryptionResponsePtrOutputWithContext(context.Context) AccountEncryptionResponsePtrOutput
}

type accountEncryptionResponsePtrType AccountEncryptionResponseArgs

func AccountEncryptionResponsePtr(v *AccountEncryptionResponseArgs) AccountEncryptionResponsePtrInput {
	return (*accountEncryptionResponsePtrType)(v)
}

func (*accountEncryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryptionResponse)(nil)).Elem()
}

func (i *accountEncryptionResponsePtrType) ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput {
	return i.ToAccountEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *accountEncryptionResponsePtrType) ToAccountEncryptionResponsePtrOutputWithContext(ctx context.Context) AccountEncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionResponsePtrOutput)
}

type AccountEncryptionResponseOutput struct{ *pulumi.OutputState }

func (AccountEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryptionResponse)(nil)).Elem()
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponseOutput() AccountEncryptionResponseOutput {
	return o
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponseOutputWithContext(ctx context.Context) AccountEncryptionResponseOutput {
	return o
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput {
	return o.ToAccountEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponsePtrOutputWithContext(ctx context.Context) AccountEncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountEncryptionResponse) *AccountEncryptionResponse {
		return &v
	}).(AccountEncryptionResponsePtrOutput)
}

func (o AccountEncryptionResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

type AccountEncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (AccountEncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryptionResponse)(nil)).Elem()
}

func (o AccountEncryptionResponsePtrOutput) ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput {
	return o
}

func (o AccountEncryptionResponsePtrOutput) ToAccountEncryptionResponsePtrOutputWithContext(ctx context.Context) AccountEncryptionResponsePtrOutput {
	return o
}

func (o AccountEncryptionResponsePtrOutput) Elem() AccountEncryptionResponseOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) AccountEncryptionResponse {
		if v != nil {
			return *v
		}
		var ret AccountEncryptionResponse
		return ret
	}).(AccountEncryptionResponseOutput)
}

func (o AccountEncryptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

type ActiveDirectory struct {
	ActiveDirectoryId          *string  `pulumi:"activeDirectoryId"`
	AdName                     *string  `pulumi:"adName"`
	Administrators             []string `pulumi:"administrators"`
	AesEncryption              *bool    `pulumi:"aesEncryption"`
	AllowLocalNfsUsersWithLdap *bool    `pulumi:"allowLocalNfsUsersWithLdap"`
	BackupOperators            []string `pulumi:"backupOperators"`
	Dns                        *string  `pulumi:"dns"`
	Domain                     *string  `pulumi:"domain"`
	EncryptDCConnections       *bool    `pulumi:"encryptDCConnections"`
	KdcIP                      *string  `pulumi:"kdcIP"`
	LdapOverTLS                *bool    `pulumi:"ldapOverTLS"`
	LdapSigning                *bool    `pulumi:"ldapSigning"`
	OrganizationalUnit         *string  `pulumi:"organizationalUnit"`
	Password                   *string  `pulumi:"password"`
	SecurityOperators          []string `pulumi:"securityOperators"`
	ServerRootCACertificate    *string  `pulumi:"serverRootCACertificate"`
	Site                       *string  `pulumi:"site"`
	SmbServerName              *string  `pulumi:"smbServerName"`
	Username                   *string  `pulumi:"username"`
}


func (val *ActiveDirectory) Defaults() *ActiveDirectory {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.OrganizationalUnit) {
		organizationalUnit_ := "CN=Computers"
		tmp.OrganizationalUnit = &organizationalUnit_
	}
	return &tmp
}





type ActiveDirectoryInput interface {
	pulumi.Input

	ToActiveDirectoryOutput() ActiveDirectoryOutput
	ToActiveDirectoryOutputWithContext(context.Context) ActiveDirectoryOutput
}

type ActiveDirectoryArgs struct {
	ActiveDirectoryId          pulumi.StringPtrInput   `pulumi:"activeDirectoryId"`
	AdName                     pulumi.StringPtrInput   `pulumi:"adName"`
	Administrators             pulumi.StringArrayInput `pulumi:"administrators"`
	AesEncryption              pulumi.BoolPtrInput     `pulumi:"aesEncryption"`
	AllowLocalNfsUsersWithLdap pulumi.BoolPtrInput     `pulumi:"allowLocalNfsUsersWithLdap"`
	BackupOperators            pulumi.StringArrayInput `pulumi:"backupOperators"`
	Dns                        pulumi.StringPtrInput   `pulumi:"dns"`
	Domain                     pulumi.StringPtrInput   `pulumi:"domain"`
	EncryptDCConnections       pulumi.BoolPtrInput     `pulumi:"encryptDCConnections"`
	KdcIP                      pulumi.StringPtrInput   `pulumi:"kdcIP"`
	LdapOverTLS                pulumi.BoolPtrInput     `pulumi:"ldapOverTLS"`
	LdapSigning                pulumi.BoolPtrInput     `pulumi:"ldapSigning"`
	OrganizationalUnit         pulumi.StringPtrInput   `pulumi:"organizationalUnit"`
	Password                   pulumi.StringPtrInput   `pulumi:"password"`
	SecurityOperators          pulumi.StringArrayInput `pulumi:"securityOperators"`
	ServerRootCACertificate    pulumi.StringPtrInput   `pulumi:"serverRootCACertificate"`
	Site                       pulumi.StringPtrInput   `pulumi:"site"`
	SmbServerName              pulumi.StringPtrInput   `pulumi:"smbServerName"`
	Username                   pulumi.StringPtrInput   `pulumi:"username"`
}

func (ActiveDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectory)(nil)).Elem()
}

func (i ActiveDirectoryArgs) ToActiveDirectoryOutput() ActiveDirectoryOutput {
	return i.ToActiveDirectoryOutputWithContext(context.Background())
}

func (i ActiveDirectoryArgs) ToActiveDirectoryOutputWithContext(ctx context.Context) ActiveDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryOutput)
}





type ActiveDirectoryArrayInput interface {
	pulumi.Input

	ToActiveDirectoryArrayOutput() ActiveDirectoryArrayOutput
	ToActiveDirectoryArrayOutputWithContext(context.Context) ActiveDirectoryArrayOutput
}

type ActiveDirectoryArray []ActiveDirectoryInput

func (ActiveDirectoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectory)(nil)).Elem()
}

func (i ActiveDirectoryArray) ToActiveDirectoryArrayOutput() ActiveDirectoryArrayOutput {
	return i.ToActiveDirectoryArrayOutputWithContext(context.Background())
}

func (i ActiveDirectoryArray) ToActiveDirectoryArrayOutputWithContext(ctx context.Context) ActiveDirectoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryArrayOutput)
}

type ActiveDirectoryOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectory)(nil)).Elem()
}

func (o ActiveDirectoryOutput) ToActiveDirectoryOutput() ActiveDirectoryOutput {
	return o
}

func (o ActiveDirectoryOutput) ToActiveDirectoryOutputWithContext(ctx context.Context) ActiveDirectoryOutput {
	return o
}

func (o ActiveDirectoryOutput) ActiveDirectoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.ActiveDirectoryId }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) AdName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.AdName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) Administrators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDirectory) []string { return v.Administrators }).(pulumi.StringArrayOutput)
}

func (o ActiveDirectoryOutput) AesEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *bool { return v.AesEncryption }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryOutput) AllowLocalNfsUsersWithLdap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *bool { return v.AllowLocalNfsUsersWithLdap }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryOutput) BackupOperators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDirectory) []string { return v.BackupOperators }).(pulumi.StringArrayOutput)
}

func (o ActiveDirectoryOutput) Dns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Dns }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) EncryptDCConnections() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *bool { return v.EncryptDCConnections }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryOutput) KdcIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.KdcIP }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) LdapOverTLS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *bool { return v.LdapOverTLS }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryOutput) LdapSigning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *bool { return v.LdapSigning }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.OrganizationalUnit }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) SecurityOperators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDirectory) []string { return v.SecurityOperators }).(pulumi.StringArrayOutput)
}

func (o ActiveDirectoryOutput) ServerRootCACertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.ServerRootCACertificate }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) Site() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Site }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) SmbServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.SmbServerName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ActiveDirectoryArrayOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectory)(nil)).Elem()
}

func (o ActiveDirectoryArrayOutput) ToActiveDirectoryArrayOutput() ActiveDirectoryArrayOutput {
	return o
}

func (o ActiveDirectoryArrayOutput) ToActiveDirectoryArrayOutputWithContext(ctx context.Context) ActiveDirectoryArrayOutput {
	return o
}

func (o ActiveDirectoryArrayOutput) Index(i pulumi.IntInput) ActiveDirectoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActiveDirectory {
		return vs[0].([]ActiveDirectory)[vs[1].(int)]
	}).(ActiveDirectoryOutput)
}

type ActiveDirectoryResponse struct {
	ActiveDirectoryId          *string  `pulumi:"activeDirectoryId"`
	AdName                     *string  `pulumi:"adName"`
	Administrators             []string `pulumi:"administrators"`
	AesEncryption              *bool    `pulumi:"aesEncryption"`
	AllowLocalNfsUsersWithLdap *bool    `pulumi:"allowLocalNfsUsersWithLdap"`
	BackupOperators            []string `pulumi:"backupOperators"`
	Dns                        *string  `pulumi:"dns"`
	Domain                     *string  `pulumi:"domain"`
	EncryptDCConnections       *bool    `pulumi:"encryptDCConnections"`
	KdcIP                      *string  `pulumi:"kdcIP"`
	LdapOverTLS                *bool    `pulumi:"ldapOverTLS"`
	LdapSigning                *bool    `pulumi:"ldapSigning"`
	OrganizationalUnit         *string  `pulumi:"organizationalUnit"`
	Password                   *string  `pulumi:"password"`
	SecurityOperators          []string `pulumi:"securityOperators"`
	ServerRootCACertificate    *string  `pulumi:"serverRootCACertificate"`
	Site                       *string  `pulumi:"site"`
	SmbServerName              *string  `pulumi:"smbServerName"`
	Status                     string   `pulumi:"status"`
	StatusDetails              string   `pulumi:"statusDetails"`
	Username                   *string  `pulumi:"username"`
}


func (val *ActiveDirectoryResponse) Defaults() *ActiveDirectoryResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.OrganizationalUnit) {
		organizationalUnit_ := "CN=Computers"
		tmp.OrganizationalUnit = &organizationalUnit_
	}
	return &tmp
}





type ActiveDirectoryResponseInput interface {
	pulumi.Input

	ToActiveDirectoryResponseOutput() ActiveDirectoryResponseOutput
	ToActiveDirectoryResponseOutputWithContext(context.Context) ActiveDirectoryResponseOutput
}

type ActiveDirectoryResponseArgs struct {
	ActiveDirectoryId          pulumi.StringPtrInput   `pulumi:"activeDirectoryId"`
	AdName                     pulumi.StringPtrInput   `pulumi:"adName"`
	Administrators             pulumi.StringArrayInput `pulumi:"administrators"`
	AesEncryption              pulumi.BoolPtrInput     `pulumi:"aesEncryption"`
	AllowLocalNfsUsersWithLdap pulumi.BoolPtrInput     `pulumi:"allowLocalNfsUsersWithLdap"`
	BackupOperators            pulumi.StringArrayInput `pulumi:"backupOperators"`
	Dns                        pulumi.StringPtrInput   `pulumi:"dns"`
	Domain                     pulumi.StringPtrInput   `pulumi:"domain"`
	EncryptDCConnections       pulumi.BoolPtrInput     `pulumi:"encryptDCConnections"`
	KdcIP                      pulumi.StringPtrInput   `pulumi:"kdcIP"`
	LdapOverTLS                pulumi.BoolPtrInput     `pulumi:"ldapOverTLS"`
	LdapSigning                pulumi.BoolPtrInput     `pulumi:"ldapSigning"`
	OrganizationalUnit         pulumi.StringPtrInput   `pulumi:"organizationalUnit"`
	Password                   pulumi.StringPtrInput   `pulumi:"password"`
	SecurityOperators          pulumi.StringArrayInput `pulumi:"securityOperators"`
	ServerRootCACertificate    pulumi.StringPtrInput   `pulumi:"serverRootCACertificate"`
	Site                       pulumi.StringPtrInput   `pulumi:"site"`
	SmbServerName              pulumi.StringPtrInput   `pulumi:"smbServerName"`
	Status                     pulumi.StringInput      `pulumi:"status"`
	StatusDetails              pulumi.StringInput      `pulumi:"statusDetails"`
	Username                   pulumi.StringPtrInput   `pulumi:"username"`
}

func (ActiveDirectoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryResponse)(nil)).Elem()
}

func (i ActiveDirectoryResponseArgs) ToActiveDirectoryResponseOutput() ActiveDirectoryResponseOutput {
	return i.ToActiveDirectoryResponseOutputWithContext(context.Background())
}

func (i ActiveDirectoryResponseArgs) ToActiveDirectoryResponseOutputWithContext(ctx context.Context) ActiveDirectoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryResponseOutput)
}





type ActiveDirectoryResponseArrayInput interface {
	pulumi.Input

	ToActiveDirectoryResponseArrayOutput() ActiveDirectoryResponseArrayOutput
	ToActiveDirectoryResponseArrayOutputWithContext(context.Context) ActiveDirectoryResponseArrayOutput
}

type ActiveDirectoryResponseArray []ActiveDirectoryResponseInput

func (ActiveDirectoryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectoryResponse)(nil)).Elem()
}

func (i ActiveDirectoryResponseArray) ToActiveDirectoryResponseArrayOutput() ActiveDirectoryResponseArrayOutput {
	return i.ToActiveDirectoryResponseArrayOutputWithContext(context.Background())
}

func (i ActiveDirectoryResponseArray) ToActiveDirectoryResponseArrayOutputWithContext(ctx context.Context) ActiveDirectoryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryResponseArrayOutput)
}

type ActiveDirectoryResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryResponse)(nil)).Elem()
}

func (o ActiveDirectoryResponseOutput) ToActiveDirectoryResponseOutput() ActiveDirectoryResponseOutput {
	return o
}

func (o ActiveDirectoryResponseOutput) ToActiveDirectoryResponseOutputWithContext(ctx context.Context) ActiveDirectoryResponseOutput {
	return o
}

func (o ActiveDirectoryResponseOutput) ActiveDirectoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.ActiveDirectoryId }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) AdName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.AdName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) Administrators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) []string { return v.Administrators }).(pulumi.StringArrayOutput)
}

func (o ActiveDirectoryResponseOutput) AesEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *bool { return v.AesEncryption }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryResponseOutput) AllowLocalNfsUsersWithLdap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *bool { return v.AllowLocalNfsUsersWithLdap }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryResponseOutput) BackupOperators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) []string { return v.BackupOperators }).(pulumi.StringArrayOutput)
}

func (o ActiveDirectoryResponseOutput) Dns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Dns }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) EncryptDCConnections() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *bool { return v.EncryptDCConnections }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryResponseOutput) KdcIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.KdcIP }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) LdapOverTLS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *bool { return v.LdapOverTLS }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryResponseOutput) LdapSigning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *bool { return v.LdapSigning }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryResponseOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.OrganizationalUnit }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) SecurityOperators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) []string { return v.SecurityOperators }).(pulumi.StringArrayOutput)
}

func (o ActiveDirectoryResponseOutput) ServerRootCACertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.ServerRootCACertificate }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) Site() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Site }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) SmbServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.SmbServerName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o ActiveDirectoryResponseOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) string { return v.StatusDetails }).(pulumi.StringOutput)
}

func (o ActiveDirectoryResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ActiveDirectoryResponseArrayOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectoryResponse)(nil)).Elem()
}

func (o ActiveDirectoryResponseArrayOutput) ToActiveDirectoryResponseArrayOutput() ActiveDirectoryResponseArrayOutput {
	return o
}

func (o ActiveDirectoryResponseArrayOutput) ToActiveDirectoryResponseArrayOutputWithContext(ctx context.Context) ActiveDirectoryResponseArrayOutput {
	return o
}

func (o ActiveDirectoryResponseArrayOutput) Index(i pulumi.IntInput) ActiveDirectoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActiveDirectoryResponse {
		return vs[0].([]ActiveDirectoryResponse)[vs[1].(int)]
	}).(ActiveDirectoryResponseOutput)
}

type DailySchedule struct {
	Hour            *int     `pulumi:"hour"`
	Minute          *int     `pulumi:"minute"`
	SnapshotsToKeep *int     `pulumi:"snapshotsToKeep"`
	UsedBytes       *float64 `pulumi:"usedBytes"`
}





type DailyScheduleInput interface {
	pulumi.Input

	ToDailyScheduleOutput() DailyScheduleOutput
	ToDailyScheduleOutputWithContext(context.Context) DailyScheduleOutput
}

type DailyScheduleArgs struct {
	Hour            pulumi.IntPtrInput     `pulumi:"hour"`
	Minute          pulumi.IntPtrInput     `pulumi:"minute"`
	SnapshotsToKeep pulumi.IntPtrInput     `pulumi:"snapshotsToKeep"`
	UsedBytes       pulumi.Float64PtrInput `pulumi:"usedBytes"`
}

func (DailyScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailySchedule)(nil)).Elem()
}

func (i DailyScheduleArgs) ToDailyScheduleOutput() DailyScheduleOutput {
	return i.ToDailyScheduleOutputWithContext(context.Background())
}

func (i DailyScheduleArgs) ToDailyScheduleOutputWithContext(ctx context.Context) DailyScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyScheduleOutput)
}

func (i DailyScheduleArgs) ToDailySchedulePtrOutput() DailySchedulePtrOutput {
	return i.ToDailySchedulePtrOutputWithContext(context.Background())
}

func (i DailyScheduleArgs) ToDailySchedulePtrOutputWithContext(ctx context.Context) DailySchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyScheduleOutput).ToDailySchedulePtrOutputWithContext(ctx)
}









type DailySchedulePtrInput interface {
	pulumi.Input

	ToDailySchedulePtrOutput() DailySchedulePtrOutput
	ToDailySchedulePtrOutputWithContext(context.Context) DailySchedulePtrOutput
}

type dailySchedulePtrType DailyScheduleArgs

func DailySchedulePtr(v *DailyScheduleArgs) DailySchedulePtrInput {
	return (*dailySchedulePtrType)(v)
}

func (*dailySchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DailySchedule)(nil)).Elem()
}

func (i *dailySchedulePtrType) ToDailySchedulePtrOutput() DailySchedulePtrOutput {
	return i.ToDailySchedulePtrOutputWithContext(context.Background())
}

func (i *dailySchedulePtrType) ToDailySchedulePtrOutputWithContext(ctx context.Context) DailySchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailySchedulePtrOutput)
}

type DailyScheduleOutput struct{ *pulumi.OutputState }

func (DailyScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailySchedule)(nil)).Elem()
}

func (o DailyScheduleOutput) ToDailyScheduleOutput() DailyScheduleOutput {
	return o
}

func (o DailyScheduleOutput) ToDailyScheduleOutputWithContext(ctx context.Context) DailyScheduleOutput {
	return o
}

func (o DailyScheduleOutput) ToDailySchedulePtrOutput() DailySchedulePtrOutput {
	return o.ToDailySchedulePtrOutputWithContext(context.Background())
}

func (o DailyScheduleOutput) ToDailySchedulePtrOutputWithContext(ctx context.Context) DailySchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DailySchedule) *DailySchedule {
		return &v
	}).(DailySchedulePtrOutput)
}

func (o DailyScheduleOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DailySchedule) *int { return v.Hour }).(pulumi.IntPtrOutput)
}

func (o DailyScheduleOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DailySchedule) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

func (o DailyScheduleOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DailySchedule) *int { return v.SnapshotsToKeep }).(pulumi.IntPtrOutput)
}

func (o DailyScheduleOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DailySchedule) *float64 { return v.UsedBytes }).(pulumi.Float64PtrOutput)
}

type DailySchedulePtrOutput struct{ *pulumi.OutputState }

func (DailySchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DailySchedule)(nil)).Elem()
}

func (o DailySchedulePtrOutput) ToDailySchedulePtrOutput() DailySchedulePtrOutput {
	return o
}

func (o DailySchedulePtrOutput) ToDailySchedulePtrOutputWithContext(ctx context.Context) DailySchedulePtrOutput {
	return o
}

func (o DailySchedulePtrOutput) Elem() DailyScheduleOutput {
	return o.ApplyT(func(v *DailySchedule) DailySchedule {
		if v != nil {
			return *v
		}
		var ret DailySchedule
		return ret
	}).(DailyScheduleOutput)
}

func (o DailySchedulePtrOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DailySchedule) *int {
		if v == nil {
			return nil
		}
		return v.Hour
	}).(pulumi.IntPtrOutput)
}

func (o DailySchedulePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DailySchedule) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

func (o DailySchedulePtrOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DailySchedule) *int {
		if v == nil {
			return nil
		}
		return v.SnapshotsToKeep
	}).(pulumi.IntPtrOutput)
}

func (o DailySchedulePtrOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DailySchedule) *float64 {
		if v == nil {
			return nil
		}
		return v.UsedBytes
	}).(pulumi.Float64PtrOutput)
}

type DailyScheduleResponse struct {
	Hour            *int     `pulumi:"hour"`
	Minute          *int     `pulumi:"minute"`
	SnapshotsToKeep *int     `pulumi:"snapshotsToKeep"`
	UsedBytes       *float64 `pulumi:"usedBytes"`
}





type DailyScheduleResponseInput interface {
	pulumi.Input

	ToDailyScheduleResponseOutput() DailyScheduleResponseOutput
	ToDailyScheduleResponseOutputWithContext(context.Context) DailyScheduleResponseOutput
}

type DailyScheduleResponseArgs struct {
	Hour            pulumi.IntPtrInput     `pulumi:"hour"`
	Minute          pulumi.IntPtrInput     `pulumi:"minute"`
	SnapshotsToKeep pulumi.IntPtrInput     `pulumi:"snapshotsToKeep"`
	UsedBytes       pulumi.Float64PtrInput `pulumi:"usedBytes"`
}

func (DailyScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyScheduleResponse)(nil)).Elem()
}

func (i DailyScheduleResponseArgs) ToDailyScheduleResponseOutput() DailyScheduleResponseOutput {
	return i.ToDailyScheduleResponseOutputWithContext(context.Background())
}

func (i DailyScheduleResponseArgs) ToDailyScheduleResponseOutputWithContext(ctx context.Context) DailyScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyScheduleResponseOutput)
}

func (i DailyScheduleResponseArgs) ToDailyScheduleResponsePtrOutput() DailyScheduleResponsePtrOutput {
	return i.ToDailyScheduleResponsePtrOutputWithContext(context.Background())
}

func (i DailyScheduleResponseArgs) ToDailyScheduleResponsePtrOutputWithContext(ctx context.Context) DailyScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyScheduleResponseOutput).ToDailyScheduleResponsePtrOutputWithContext(ctx)
}









type DailyScheduleResponsePtrInput interface {
	pulumi.Input

	ToDailyScheduleResponsePtrOutput() DailyScheduleResponsePtrOutput
	ToDailyScheduleResponsePtrOutputWithContext(context.Context) DailyScheduleResponsePtrOutput
}

type dailyScheduleResponsePtrType DailyScheduleResponseArgs

func DailyScheduleResponsePtr(v *DailyScheduleResponseArgs) DailyScheduleResponsePtrInput {
	return (*dailyScheduleResponsePtrType)(v)
}

func (*dailyScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyScheduleResponse)(nil)).Elem()
}

func (i *dailyScheduleResponsePtrType) ToDailyScheduleResponsePtrOutput() DailyScheduleResponsePtrOutput {
	return i.ToDailyScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *dailyScheduleResponsePtrType) ToDailyScheduleResponsePtrOutputWithContext(ctx context.Context) DailyScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyScheduleResponsePtrOutput)
}

type DailyScheduleResponseOutput struct{ *pulumi.OutputState }

func (DailyScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyScheduleResponse)(nil)).Elem()
}

func (o DailyScheduleResponseOutput) ToDailyScheduleResponseOutput() DailyScheduleResponseOutput {
	return o
}

func (o DailyScheduleResponseOutput) ToDailyScheduleResponseOutputWithContext(ctx context.Context) DailyScheduleResponseOutput {
	return o
}

func (o DailyScheduleResponseOutput) ToDailyScheduleResponsePtrOutput() DailyScheduleResponsePtrOutput {
	return o.ToDailyScheduleResponsePtrOutputWithContext(context.Background())
}

func (o DailyScheduleResponseOutput) ToDailyScheduleResponsePtrOutputWithContext(ctx context.Context) DailyScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DailyScheduleResponse) *DailyScheduleResponse {
		return &v
	}).(DailyScheduleResponsePtrOutput)
}

func (o DailyScheduleResponseOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DailyScheduleResponse) *int { return v.Hour }).(pulumi.IntPtrOutput)
}

func (o DailyScheduleResponseOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DailyScheduleResponse) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

func (o DailyScheduleResponseOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DailyScheduleResponse) *int { return v.SnapshotsToKeep }).(pulumi.IntPtrOutput)
}

func (o DailyScheduleResponseOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DailyScheduleResponse) *float64 { return v.UsedBytes }).(pulumi.Float64PtrOutput)
}

type DailyScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (DailyScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyScheduleResponse)(nil)).Elem()
}

func (o DailyScheduleResponsePtrOutput) ToDailyScheduleResponsePtrOutput() DailyScheduleResponsePtrOutput {
	return o
}

func (o DailyScheduleResponsePtrOutput) ToDailyScheduleResponsePtrOutputWithContext(ctx context.Context) DailyScheduleResponsePtrOutput {
	return o
}

func (o DailyScheduleResponsePtrOutput) Elem() DailyScheduleResponseOutput {
	return o.ApplyT(func(v *DailyScheduleResponse) DailyScheduleResponse {
		if v != nil {
			return *v
		}
		var ret DailyScheduleResponse
		return ret
	}).(DailyScheduleResponseOutput)
}

func (o DailyScheduleResponsePtrOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DailyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.Hour
	}).(pulumi.IntPtrOutput)
}

func (o DailyScheduleResponsePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DailyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

func (o DailyScheduleResponsePtrOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DailyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.SnapshotsToKeep
	}).(pulumi.IntPtrOutput)
}

func (o DailyScheduleResponsePtrOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DailyScheduleResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.UsedBytes
	}).(pulumi.Float64PtrOutput)
}

type ExportPolicyRule struct {
	AllowedClients      *string `pulumi:"allowedClients"`
	ChownMode           *string `pulumi:"chownMode"`
	Cifs                *bool   `pulumi:"cifs"`
	HasRootAccess       *bool   `pulumi:"hasRootAccess"`
	Kerberos5ReadOnly   *bool   `pulumi:"kerberos5ReadOnly"`
	Kerberos5ReadWrite  *bool   `pulumi:"kerberos5ReadWrite"`
	Kerberos5iReadOnly  *bool   `pulumi:"kerberos5iReadOnly"`
	Kerberos5iReadWrite *bool   `pulumi:"kerberos5iReadWrite"`
	Kerberos5pReadOnly  *bool   `pulumi:"kerberos5pReadOnly"`
	Kerberos5pReadWrite *bool   `pulumi:"kerberos5pReadWrite"`
	Nfsv3               *bool   `pulumi:"nfsv3"`
	Nfsv41              *bool   `pulumi:"nfsv41"`
	RuleIndex           *int    `pulumi:"ruleIndex"`
	UnixReadOnly        *bool   `pulumi:"unixReadOnly"`
	UnixReadWrite       *bool   `pulumi:"unixReadWrite"`
}


func (val *ExportPolicyRule) Defaults() *ExportPolicyRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ChownMode) {
		chownMode_ := "Restricted"
		tmp.ChownMode = &chownMode_
	}
	if isZero(tmp.HasRootAccess) {
		hasRootAccess_ := true
		tmp.HasRootAccess = &hasRootAccess_
	}
	if isZero(tmp.Kerberos5ReadOnly) {
		kerberos5ReadOnly_ := false
		tmp.Kerberos5ReadOnly = &kerberos5ReadOnly_
	}
	if isZero(tmp.Kerberos5ReadWrite) {
		kerberos5ReadWrite_ := false
		tmp.Kerberos5ReadWrite = &kerberos5ReadWrite_
	}
	if isZero(tmp.Kerberos5iReadOnly) {
		kerberos5iReadOnly_ := false
		tmp.Kerberos5iReadOnly = &kerberos5iReadOnly_
	}
	if isZero(tmp.Kerberos5iReadWrite) {
		kerberos5iReadWrite_ := false
		tmp.Kerberos5iReadWrite = &kerberos5iReadWrite_
	}
	if isZero(tmp.Kerberos5pReadOnly) {
		kerberos5pReadOnly_ := false
		tmp.Kerberos5pReadOnly = &kerberos5pReadOnly_
	}
	if isZero(tmp.Kerberos5pReadWrite) {
		kerberos5pReadWrite_ := false
		tmp.Kerberos5pReadWrite = &kerberos5pReadWrite_
	}
	return &tmp
}





type ExportPolicyRuleInput interface {
	pulumi.Input

	ToExportPolicyRuleOutput() ExportPolicyRuleOutput
	ToExportPolicyRuleOutputWithContext(context.Context) ExportPolicyRuleOutput
}

type ExportPolicyRuleArgs struct {
	AllowedClients      pulumi.StringPtrInput `pulumi:"allowedClients"`
	ChownMode           pulumi.StringPtrInput `pulumi:"chownMode"`
	Cifs                pulumi.BoolPtrInput   `pulumi:"cifs"`
	HasRootAccess       pulumi.BoolPtrInput   `pulumi:"hasRootAccess"`
	Kerberos5ReadOnly   pulumi.BoolPtrInput   `pulumi:"kerberos5ReadOnly"`
	Kerberos5ReadWrite  pulumi.BoolPtrInput   `pulumi:"kerberos5ReadWrite"`
	Kerberos5iReadOnly  pulumi.BoolPtrInput   `pulumi:"kerberos5iReadOnly"`
	Kerberos5iReadWrite pulumi.BoolPtrInput   `pulumi:"kerberos5iReadWrite"`
	Kerberos5pReadOnly  pulumi.BoolPtrInput   `pulumi:"kerberos5pReadOnly"`
	Kerberos5pReadWrite pulumi.BoolPtrInput   `pulumi:"kerberos5pReadWrite"`
	Nfsv3               pulumi.BoolPtrInput   `pulumi:"nfsv3"`
	Nfsv41              pulumi.BoolPtrInput   `pulumi:"nfsv41"`
	RuleIndex           pulumi.IntPtrInput    `pulumi:"ruleIndex"`
	UnixReadOnly        pulumi.BoolPtrInput   `pulumi:"unixReadOnly"`
	UnixReadWrite       pulumi.BoolPtrInput   `pulumi:"unixReadWrite"`
}

func (ExportPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPolicyRule)(nil)).Elem()
}

func (i ExportPolicyRuleArgs) ToExportPolicyRuleOutput() ExportPolicyRuleOutput {
	return i.ToExportPolicyRuleOutputWithContext(context.Background())
}

func (i ExportPolicyRuleArgs) ToExportPolicyRuleOutputWithContext(ctx context.Context) ExportPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPolicyRuleOutput)
}





type ExportPolicyRuleArrayInput interface {
	pulumi.Input

	ToExportPolicyRuleArrayOutput() ExportPolicyRuleArrayOutput
	ToExportPolicyRuleArrayOutputWithContext(context.Context) ExportPolicyRuleArrayOutput
}

type ExportPolicyRuleArray []ExportPolicyRuleInput

func (ExportPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExportPolicyRule)(nil)).Elem()
}

func (i ExportPolicyRuleArray) ToExportPolicyRuleArrayOutput() ExportPolicyRuleArrayOutput {
	return i.ToExportPolicyRuleArrayOutputWithContext(context.Background())
}

func (i ExportPolicyRuleArray) ToExportPolicyRuleArrayOutputWithContext(ctx context.Context) ExportPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPolicyRuleArrayOutput)
}

type ExportPolicyRuleOutput struct{ *pulumi.OutputState }

func (ExportPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPolicyRule)(nil)).Elem()
}

func (o ExportPolicyRuleOutput) ToExportPolicyRuleOutput() ExportPolicyRuleOutput {
	return o
}

func (o ExportPolicyRuleOutput) ToExportPolicyRuleOutputWithContext(ctx context.Context) ExportPolicyRuleOutput {
	return o
}

func (o ExportPolicyRuleOutput) AllowedClients() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *string { return v.AllowedClients }).(pulumi.StringPtrOutput)
}

func (o ExportPolicyRuleOutput) ChownMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *string { return v.ChownMode }).(pulumi.StringPtrOutput)
}

func (o ExportPolicyRuleOutput) Cifs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Cifs }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) HasRootAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.HasRootAccess }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) Kerberos5ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Kerberos5ReadOnly }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) Kerberos5ReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Kerberos5ReadWrite }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) Kerberos5iReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Kerberos5iReadOnly }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) Kerberos5iReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Kerberos5iReadWrite }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) Kerberos5pReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Kerberos5pReadOnly }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) Kerberos5pReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Kerberos5pReadWrite }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) Nfsv3() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Nfsv3 }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) Nfsv41() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Nfsv41 }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) RuleIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *int { return v.RuleIndex }).(pulumi.IntPtrOutput)
}

func (o ExportPolicyRuleOutput) UnixReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.UnixReadOnly }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) UnixReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.UnixReadWrite }).(pulumi.BoolPtrOutput)
}

type ExportPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (ExportPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExportPolicyRule)(nil)).Elem()
}

func (o ExportPolicyRuleArrayOutput) ToExportPolicyRuleArrayOutput() ExportPolicyRuleArrayOutput {
	return o
}

func (o ExportPolicyRuleArrayOutput) ToExportPolicyRuleArrayOutputWithContext(ctx context.Context) ExportPolicyRuleArrayOutput {
	return o
}

func (o ExportPolicyRuleArrayOutput) Index(i pulumi.IntInput) ExportPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExportPolicyRule {
		return vs[0].([]ExportPolicyRule)[vs[1].(int)]
	}).(ExportPolicyRuleOutput)
}

type ExportPolicyRuleResponse struct {
	AllowedClients      *string `pulumi:"allowedClients"`
	ChownMode           *string `pulumi:"chownMode"`
	Cifs                *bool   `pulumi:"cifs"`
	HasRootAccess       *bool   `pulumi:"hasRootAccess"`
	Kerberos5ReadOnly   *bool   `pulumi:"kerberos5ReadOnly"`
	Kerberos5ReadWrite  *bool   `pulumi:"kerberos5ReadWrite"`
	Kerberos5iReadOnly  *bool   `pulumi:"kerberos5iReadOnly"`
	Kerberos5iReadWrite *bool   `pulumi:"kerberos5iReadWrite"`
	Kerberos5pReadOnly  *bool   `pulumi:"kerberos5pReadOnly"`
	Kerberos5pReadWrite *bool   `pulumi:"kerberos5pReadWrite"`
	Nfsv3               *bool   `pulumi:"nfsv3"`
	Nfsv41              *bool   `pulumi:"nfsv41"`
	RuleIndex           *int    `pulumi:"ruleIndex"`
	UnixReadOnly        *bool   `pulumi:"unixReadOnly"`
	UnixReadWrite       *bool   `pulumi:"unixReadWrite"`
}


func (val *ExportPolicyRuleResponse) Defaults() *ExportPolicyRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ChownMode) {
		chownMode_ := "Restricted"
		tmp.ChownMode = &chownMode_
	}
	if isZero(tmp.HasRootAccess) {
		hasRootAccess_ := true
		tmp.HasRootAccess = &hasRootAccess_
	}
	if isZero(tmp.Kerberos5ReadOnly) {
		kerberos5ReadOnly_ := false
		tmp.Kerberos5ReadOnly = &kerberos5ReadOnly_
	}
	if isZero(tmp.Kerberos5ReadWrite) {
		kerberos5ReadWrite_ := false
		tmp.Kerberos5ReadWrite = &kerberos5ReadWrite_
	}
	if isZero(tmp.Kerberos5iReadOnly) {
		kerberos5iReadOnly_ := false
		tmp.Kerberos5iReadOnly = &kerberos5iReadOnly_
	}
	if isZero(tmp.Kerberos5iReadWrite) {
		kerberos5iReadWrite_ := false
		tmp.Kerberos5iReadWrite = &kerberos5iReadWrite_
	}
	if isZero(tmp.Kerberos5pReadOnly) {
		kerberos5pReadOnly_ := false
		tmp.Kerberos5pReadOnly = &kerberos5pReadOnly_
	}
	if isZero(tmp.Kerberos5pReadWrite) {
		kerberos5pReadWrite_ := false
		tmp.Kerberos5pReadWrite = &kerberos5pReadWrite_
	}
	return &tmp
}





type ExportPolicyRuleResponseInput interface {
	pulumi.Input

	ToExportPolicyRuleResponseOutput() ExportPolicyRuleResponseOutput
	ToExportPolicyRuleResponseOutputWithContext(context.Context) ExportPolicyRuleResponseOutput
}

type ExportPolicyRuleResponseArgs struct {
	AllowedClients      pulumi.StringPtrInput `pulumi:"allowedClients"`
	ChownMode           pulumi.StringPtrInput `pulumi:"chownMode"`
	Cifs                pulumi.BoolPtrInput   `pulumi:"cifs"`
	HasRootAccess       pulumi.BoolPtrInput   `pulumi:"hasRootAccess"`
	Kerberos5ReadOnly   pulumi.BoolPtrInput   `pulumi:"kerberos5ReadOnly"`
	Kerberos5ReadWrite  pulumi.BoolPtrInput   `pulumi:"kerberos5ReadWrite"`
	Kerberos5iReadOnly  pulumi.BoolPtrInput   `pulumi:"kerberos5iReadOnly"`
	Kerberos5iReadWrite pulumi.BoolPtrInput   `pulumi:"kerberos5iReadWrite"`
	Kerberos5pReadOnly  pulumi.BoolPtrInput   `pulumi:"kerberos5pReadOnly"`
	Kerberos5pReadWrite pulumi.BoolPtrInput   `pulumi:"kerberos5pReadWrite"`
	Nfsv3               pulumi.BoolPtrInput   `pulumi:"nfsv3"`
	Nfsv41              pulumi.BoolPtrInput   `pulumi:"nfsv41"`
	RuleIndex           pulumi.IntPtrInput    `pulumi:"ruleIndex"`
	UnixReadOnly        pulumi.BoolPtrInput   `pulumi:"unixReadOnly"`
	UnixReadWrite       pulumi.BoolPtrInput   `pulumi:"unixReadWrite"`
}

func (ExportPolicyRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPolicyRuleResponse)(nil)).Elem()
}

func (i ExportPolicyRuleResponseArgs) ToExportPolicyRuleResponseOutput() ExportPolicyRuleResponseOutput {
	return i.ToExportPolicyRuleResponseOutputWithContext(context.Background())
}

func (i ExportPolicyRuleResponseArgs) ToExportPolicyRuleResponseOutputWithContext(ctx context.Context) ExportPolicyRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPolicyRuleResponseOutput)
}





type ExportPolicyRuleResponseArrayInput interface {
	pulumi.Input

	ToExportPolicyRuleResponseArrayOutput() ExportPolicyRuleResponseArrayOutput
	ToExportPolicyRuleResponseArrayOutputWithContext(context.Context) ExportPolicyRuleResponseArrayOutput
}

type ExportPolicyRuleResponseArray []ExportPolicyRuleResponseInput

func (ExportPolicyRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExportPolicyRuleResponse)(nil)).Elem()
}

func (i ExportPolicyRuleResponseArray) ToExportPolicyRuleResponseArrayOutput() ExportPolicyRuleResponseArrayOutput {
	return i.ToExportPolicyRuleResponseArrayOutputWithContext(context.Background())
}

func (i ExportPolicyRuleResponseArray) ToExportPolicyRuleResponseArrayOutputWithContext(ctx context.Context) ExportPolicyRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPolicyRuleResponseArrayOutput)
}

type ExportPolicyRuleResponseOutput struct{ *pulumi.OutputState }

func (ExportPolicyRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPolicyRuleResponse)(nil)).Elem()
}

func (o ExportPolicyRuleResponseOutput) ToExportPolicyRuleResponseOutput() ExportPolicyRuleResponseOutput {
	return o
}

func (o ExportPolicyRuleResponseOutput) ToExportPolicyRuleResponseOutputWithContext(ctx context.Context) ExportPolicyRuleResponseOutput {
	return o
}

func (o ExportPolicyRuleResponseOutput) AllowedClients() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *string { return v.AllowedClients }).(pulumi.StringPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) ChownMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *string { return v.ChownMode }).(pulumi.StringPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Cifs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Cifs }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) HasRootAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.HasRootAccess }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Kerberos5ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Kerberos5ReadOnly }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Kerberos5ReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Kerberos5ReadWrite }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Kerberos5iReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Kerberos5iReadOnly }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Kerberos5iReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Kerberos5iReadWrite }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Kerberos5pReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Kerberos5pReadOnly }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Kerberos5pReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Kerberos5pReadWrite }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Nfsv3() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Nfsv3 }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Nfsv41() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Nfsv41 }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) RuleIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *int { return v.RuleIndex }).(pulumi.IntPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) UnixReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.UnixReadOnly }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) UnixReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.UnixReadWrite }).(pulumi.BoolPtrOutput)
}

type ExportPolicyRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ExportPolicyRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExportPolicyRuleResponse)(nil)).Elem()
}

func (o ExportPolicyRuleResponseArrayOutput) ToExportPolicyRuleResponseArrayOutput() ExportPolicyRuleResponseArrayOutput {
	return o
}

func (o ExportPolicyRuleResponseArrayOutput) ToExportPolicyRuleResponseArrayOutputWithContext(ctx context.Context) ExportPolicyRuleResponseArrayOutput {
	return o
}

func (o ExportPolicyRuleResponseArrayOutput) Index(i pulumi.IntInput) ExportPolicyRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExportPolicyRuleResponse {
		return vs[0].([]ExportPolicyRuleResponse)[vs[1].(int)]
	}).(ExportPolicyRuleResponseOutput)
}

type HourlySchedule struct {
	Minute          *int     `pulumi:"minute"`
	SnapshotsToKeep *int     `pulumi:"snapshotsToKeep"`
	UsedBytes       *float64 `pulumi:"usedBytes"`
}





type HourlyScheduleInput interface {
	pulumi.Input

	ToHourlyScheduleOutput() HourlyScheduleOutput
	ToHourlyScheduleOutputWithContext(context.Context) HourlyScheduleOutput
}

type HourlyScheduleArgs struct {
	Minute          pulumi.IntPtrInput     `pulumi:"minute"`
	SnapshotsToKeep pulumi.IntPtrInput     `pulumi:"snapshotsToKeep"`
	UsedBytes       pulumi.Float64PtrInput `pulumi:"usedBytes"`
}

func (HourlyScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HourlySchedule)(nil)).Elem()
}

func (i HourlyScheduleArgs) ToHourlyScheduleOutput() HourlyScheduleOutput {
	return i.ToHourlyScheduleOutputWithContext(context.Background())
}

func (i HourlyScheduleArgs) ToHourlyScheduleOutputWithContext(ctx context.Context) HourlyScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlyScheduleOutput)
}

func (i HourlyScheduleArgs) ToHourlySchedulePtrOutput() HourlySchedulePtrOutput {
	return i.ToHourlySchedulePtrOutputWithContext(context.Background())
}

func (i HourlyScheduleArgs) ToHourlySchedulePtrOutputWithContext(ctx context.Context) HourlySchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlyScheduleOutput).ToHourlySchedulePtrOutputWithContext(ctx)
}









type HourlySchedulePtrInput interface {
	pulumi.Input

	ToHourlySchedulePtrOutput() HourlySchedulePtrOutput
	ToHourlySchedulePtrOutputWithContext(context.Context) HourlySchedulePtrOutput
}

type hourlySchedulePtrType HourlyScheduleArgs

func HourlySchedulePtr(v *HourlyScheduleArgs) HourlySchedulePtrInput {
	return (*hourlySchedulePtrType)(v)
}

func (*hourlySchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HourlySchedule)(nil)).Elem()
}

func (i *hourlySchedulePtrType) ToHourlySchedulePtrOutput() HourlySchedulePtrOutput {
	return i.ToHourlySchedulePtrOutputWithContext(context.Background())
}

func (i *hourlySchedulePtrType) ToHourlySchedulePtrOutputWithContext(ctx context.Context) HourlySchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlySchedulePtrOutput)
}

type HourlyScheduleOutput struct{ *pulumi.OutputState }

func (HourlyScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HourlySchedule)(nil)).Elem()
}

func (o HourlyScheduleOutput) ToHourlyScheduleOutput() HourlyScheduleOutput {
	return o
}

func (o HourlyScheduleOutput) ToHourlyScheduleOutputWithContext(ctx context.Context) HourlyScheduleOutput {
	return o
}

func (o HourlyScheduleOutput) ToHourlySchedulePtrOutput() HourlySchedulePtrOutput {
	return o.ToHourlySchedulePtrOutputWithContext(context.Background())
}

func (o HourlyScheduleOutput) ToHourlySchedulePtrOutputWithContext(ctx context.Context) HourlySchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HourlySchedule) *HourlySchedule {
		return &v
	}).(HourlySchedulePtrOutput)
}

func (o HourlyScheduleOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourlySchedule) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourlySchedule) *int { return v.SnapshotsToKeep }).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HourlySchedule) *float64 { return v.UsedBytes }).(pulumi.Float64PtrOutput)
}

type HourlySchedulePtrOutput struct{ *pulumi.OutputState }

func (HourlySchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HourlySchedule)(nil)).Elem()
}

func (o HourlySchedulePtrOutput) ToHourlySchedulePtrOutput() HourlySchedulePtrOutput {
	return o
}

func (o HourlySchedulePtrOutput) ToHourlySchedulePtrOutputWithContext(ctx context.Context) HourlySchedulePtrOutput {
	return o
}

func (o HourlySchedulePtrOutput) Elem() HourlyScheduleOutput {
	return o.ApplyT(func(v *HourlySchedule) HourlySchedule {
		if v != nil {
			return *v
		}
		var ret HourlySchedule
		return ret
	}).(HourlyScheduleOutput)
}

func (o HourlySchedulePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourlySchedule) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

func (o HourlySchedulePtrOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourlySchedule) *int {
		if v == nil {
			return nil
		}
		return v.SnapshotsToKeep
	}).(pulumi.IntPtrOutput)
}

func (o HourlySchedulePtrOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *HourlySchedule) *float64 {
		if v == nil {
			return nil
		}
		return v.UsedBytes
	}).(pulumi.Float64PtrOutput)
}

type HourlyScheduleResponse struct {
	Minute          *int     `pulumi:"minute"`
	SnapshotsToKeep *int     `pulumi:"snapshotsToKeep"`
	UsedBytes       *float64 `pulumi:"usedBytes"`
}





type HourlyScheduleResponseInput interface {
	pulumi.Input

	ToHourlyScheduleResponseOutput() HourlyScheduleResponseOutput
	ToHourlyScheduleResponseOutputWithContext(context.Context) HourlyScheduleResponseOutput
}

type HourlyScheduleResponseArgs struct {
	Minute          pulumi.IntPtrInput     `pulumi:"minute"`
	SnapshotsToKeep pulumi.IntPtrInput     `pulumi:"snapshotsToKeep"`
	UsedBytes       pulumi.Float64PtrInput `pulumi:"usedBytes"`
}

func (HourlyScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HourlyScheduleResponse)(nil)).Elem()
}

func (i HourlyScheduleResponseArgs) ToHourlyScheduleResponseOutput() HourlyScheduleResponseOutput {
	return i.ToHourlyScheduleResponseOutputWithContext(context.Background())
}

func (i HourlyScheduleResponseArgs) ToHourlyScheduleResponseOutputWithContext(ctx context.Context) HourlyScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlyScheduleResponseOutput)
}

func (i HourlyScheduleResponseArgs) ToHourlyScheduleResponsePtrOutput() HourlyScheduleResponsePtrOutput {
	return i.ToHourlyScheduleResponsePtrOutputWithContext(context.Background())
}

func (i HourlyScheduleResponseArgs) ToHourlyScheduleResponsePtrOutputWithContext(ctx context.Context) HourlyScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlyScheduleResponseOutput).ToHourlyScheduleResponsePtrOutputWithContext(ctx)
}









type HourlyScheduleResponsePtrInput interface {
	pulumi.Input

	ToHourlyScheduleResponsePtrOutput() HourlyScheduleResponsePtrOutput
	ToHourlyScheduleResponsePtrOutputWithContext(context.Context) HourlyScheduleResponsePtrOutput
}

type hourlyScheduleResponsePtrType HourlyScheduleResponseArgs

func HourlyScheduleResponsePtr(v *HourlyScheduleResponseArgs) HourlyScheduleResponsePtrInput {
	return (*hourlyScheduleResponsePtrType)(v)
}

func (*hourlyScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HourlyScheduleResponse)(nil)).Elem()
}

func (i *hourlyScheduleResponsePtrType) ToHourlyScheduleResponsePtrOutput() HourlyScheduleResponsePtrOutput {
	return i.ToHourlyScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *hourlyScheduleResponsePtrType) ToHourlyScheduleResponsePtrOutputWithContext(ctx context.Context) HourlyScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlyScheduleResponsePtrOutput)
}

type HourlyScheduleResponseOutput struct{ *pulumi.OutputState }

func (HourlyScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HourlyScheduleResponse)(nil)).Elem()
}

func (o HourlyScheduleResponseOutput) ToHourlyScheduleResponseOutput() HourlyScheduleResponseOutput {
	return o
}

func (o HourlyScheduleResponseOutput) ToHourlyScheduleResponseOutputWithContext(ctx context.Context) HourlyScheduleResponseOutput {
	return o
}

func (o HourlyScheduleResponseOutput) ToHourlyScheduleResponsePtrOutput() HourlyScheduleResponsePtrOutput {
	return o.ToHourlyScheduleResponsePtrOutputWithContext(context.Background())
}

func (o HourlyScheduleResponseOutput) ToHourlyScheduleResponsePtrOutputWithContext(ctx context.Context) HourlyScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HourlyScheduleResponse) *HourlyScheduleResponse {
		return &v
	}).(HourlyScheduleResponsePtrOutput)
}

func (o HourlyScheduleResponseOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourlyScheduleResponse) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleResponseOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourlyScheduleResponse) *int { return v.SnapshotsToKeep }).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleResponseOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HourlyScheduleResponse) *float64 { return v.UsedBytes }).(pulumi.Float64PtrOutput)
}

type HourlyScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (HourlyScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HourlyScheduleResponse)(nil)).Elem()
}

func (o HourlyScheduleResponsePtrOutput) ToHourlyScheduleResponsePtrOutput() HourlyScheduleResponsePtrOutput {
	return o
}

func (o HourlyScheduleResponsePtrOutput) ToHourlyScheduleResponsePtrOutputWithContext(ctx context.Context) HourlyScheduleResponsePtrOutput {
	return o
}

func (o HourlyScheduleResponsePtrOutput) Elem() HourlyScheduleResponseOutput {
	return o.ApplyT(func(v *HourlyScheduleResponse) HourlyScheduleResponse {
		if v != nil {
			return *v
		}
		var ret HourlyScheduleResponse
		return ret
	}).(HourlyScheduleResponseOutput)
}

func (o HourlyScheduleResponsePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourlyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleResponsePtrOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourlyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.SnapshotsToKeep
	}).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleResponsePtrOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *HourlyScheduleResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.UsedBytes
	}).(pulumi.Float64PtrOutput)
}

type MonthlySchedule struct {
	DaysOfMonth     *string  `pulumi:"daysOfMonth"`
	Hour            *int     `pulumi:"hour"`
	Minute          *int     `pulumi:"minute"`
	SnapshotsToKeep *int     `pulumi:"snapshotsToKeep"`
	UsedBytes       *float64 `pulumi:"usedBytes"`
}





type MonthlyScheduleInput interface {
	pulumi.Input

	ToMonthlyScheduleOutput() MonthlyScheduleOutput
	ToMonthlyScheduleOutputWithContext(context.Context) MonthlyScheduleOutput
}

type MonthlyScheduleArgs struct {
	DaysOfMonth     pulumi.StringPtrInput  `pulumi:"daysOfMonth"`
	Hour            pulumi.IntPtrInput     `pulumi:"hour"`
	Minute          pulumi.IntPtrInput     `pulumi:"minute"`
	SnapshotsToKeep pulumi.IntPtrInput     `pulumi:"snapshotsToKeep"`
	UsedBytes       pulumi.Float64PtrInput `pulumi:"usedBytes"`
}

func (MonthlyScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlySchedule)(nil)).Elem()
}

func (i MonthlyScheduleArgs) ToMonthlyScheduleOutput() MonthlyScheduleOutput {
	return i.ToMonthlyScheduleOutputWithContext(context.Background())
}

func (i MonthlyScheduleArgs) ToMonthlyScheduleOutputWithContext(ctx context.Context) MonthlyScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyScheduleOutput)
}

func (i MonthlyScheduleArgs) ToMonthlySchedulePtrOutput() MonthlySchedulePtrOutput {
	return i.ToMonthlySchedulePtrOutputWithContext(context.Background())
}

func (i MonthlyScheduleArgs) ToMonthlySchedulePtrOutputWithContext(ctx context.Context) MonthlySchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyScheduleOutput).ToMonthlySchedulePtrOutputWithContext(ctx)
}









type MonthlySchedulePtrInput interface {
	pulumi.Input

	ToMonthlySchedulePtrOutput() MonthlySchedulePtrOutput
	ToMonthlySchedulePtrOutputWithContext(context.Context) MonthlySchedulePtrOutput
}

type monthlySchedulePtrType MonthlyScheduleArgs

func MonthlySchedulePtr(v *MonthlyScheduleArgs) MonthlySchedulePtrInput {
	return (*monthlySchedulePtrType)(v)
}

func (*monthlySchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlySchedule)(nil)).Elem()
}

func (i *monthlySchedulePtrType) ToMonthlySchedulePtrOutput() MonthlySchedulePtrOutput {
	return i.ToMonthlySchedulePtrOutputWithContext(context.Background())
}

func (i *monthlySchedulePtrType) ToMonthlySchedulePtrOutputWithContext(ctx context.Context) MonthlySchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlySchedulePtrOutput)
}

type MonthlyScheduleOutput struct{ *pulumi.OutputState }

func (MonthlyScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlySchedule)(nil)).Elem()
}

func (o MonthlyScheduleOutput) ToMonthlyScheduleOutput() MonthlyScheduleOutput {
	return o
}

func (o MonthlyScheduleOutput) ToMonthlyScheduleOutputWithContext(ctx context.Context) MonthlyScheduleOutput {
	return o
}

func (o MonthlyScheduleOutput) ToMonthlySchedulePtrOutput() MonthlySchedulePtrOutput {
	return o.ToMonthlySchedulePtrOutputWithContext(context.Background())
}

func (o MonthlyScheduleOutput) ToMonthlySchedulePtrOutputWithContext(ctx context.Context) MonthlySchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonthlySchedule) *MonthlySchedule {
		return &v
	}).(MonthlySchedulePtrOutput)
}

func (o MonthlyScheduleOutput) DaysOfMonth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonthlySchedule) *string { return v.DaysOfMonth }).(pulumi.StringPtrOutput)
}

func (o MonthlyScheduleOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonthlySchedule) *int { return v.Hour }).(pulumi.IntPtrOutput)
}

func (o MonthlyScheduleOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonthlySchedule) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

func (o MonthlyScheduleOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonthlySchedule) *int { return v.SnapshotsToKeep }).(pulumi.IntPtrOutput)
}

func (o MonthlyScheduleOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonthlySchedule) *float64 { return v.UsedBytes }).(pulumi.Float64PtrOutput)
}

type MonthlySchedulePtrOutput struct{ *pulumi.OutputState }

func (MonthlySchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlySchedule)(nil)).Elem()
}

func (o MonthlySchedulePtrOutput) ToMonthlySchedulePtrOutput() MonthlySchedulePtrOutput {
	return o
}

func (o MonthlySchedulePtrOutput) ToMonthlySchedulePtrOutputWithContext(ctx context.Context) MonthlySchedulePtrOutput {
	return o
}

func (o MonthlySchedulePtrOutput) Elem() MonthlyScheduleOutput {
	return o.ApplyT(func(v *MonthlySchedule) MonthlySchedule {
		if v != nil {
			return *v
		}
		var ret MonthlySchedule
		return ret
	}).(MonthlyScheduleOutput)
}

func (o MonthlySchedulePtrOutput) DaysOfMonth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonthlySchedule) *string {
		if v == nil {
			return nil
		}
		return v.DaysOfMonth
	}).(pulumi.StringPtrOutput)
}

func (o MonthlySchedulePtrOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MonthlySchedule) *int {
		if v == nil {
			return nil
		}
		return v.Hour
	}).(pulumi.IntPtrOutput)
}

func (o MonthlySchedulePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MonthlySchedule) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

func (o MonthlySchedulePtrOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MonthlySchedule) *int {
		if v == nil {
			return nil
		}
		return v.SnapshotsToKeep
	}).(pulumi.IntPtrOutput)
}

func (o MonthlySchedulePtrOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonthlySchedule) *float64 {
		if v == nil {
			return nil
		}
		return v.UsedBytes
	}).(pulumi.Float64PtrOutput)
}

type MonthlyScheduleResponse struct {
	DaysOfMonth     *string  `pulumi:"daysOfMonth"`
	Hour            *int     `pulumi:"hour"`
	Minute          *int     `pulumi:"minute"`
	SnapshotsToKeep *int     `pulumi:"snapshotsToKeep"`
	UsedBytes       *float64 `pulumi:"usedBytes"`
}





type MonthlyScheduleResponseInput interface {
	pulumi.Input

	ToMonthlyScheduleResponseOutput() MonthlyScheduleResponseOutput
	ToMonthlyScheduleResponseOutputWithContext(context.Context) MonthlyScheduleResponseOutput
}

type MonthlyScheduleResponseArgs struct {
	DaysOfMonth     pulumi.StringPtrInput  `pulumi:"daysOfMonth"`
	Hour            pulumi.IntPtrInput     `pulumi:"hour"`
	Minute          pulumi.IntPtrInput     `pulumi:"minute"`
	SnapshotsToKeep pulumi.IntPtrInput     `pulumi:"snapshotsToKeep"`
	UsedBytes       pulumi.Float64PtrInput `pulumi:"usedBytes"`
}

func (MonthlyScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyScheduleResponse)(nil)).Elem()
}

func (i MonthlyScheduleResponseArgs) ToMonthlyScheduleResponseOutput() MonthlyScheduleResponseOutput {
	return i.ToMonthlyScheduleResponseOutputWithContext(context.Background())
}

func (i MonthlyScheduleResponseArgs) ToMonthlyScheduleResponseOutputWithContext(ctx context.Context) MonthlyScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyScheduleResponseOutput)
}

func (i MonthlyScheduleResponseArgs) ToMonthlyScheduleResponsePtrOutput() MonthlyScheduleResponsePtrOutput {
	return i.ToMonthlyScheduleResponsePtrOutputWithContext(context.Background())
}

func (i MonthlyScheduleResponseArgs) ToMonthlyScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyScheduleResponseOutput).ToMonthlyScheduleResponsePtrOutputWithContext(ctx)
}









type MonthlyScheduleResponsePtrInput interface {
	pulumi.Input

	ToMonthlyScheduleResponsePtrOutput() MonthlyScheduleResponsePtrOutput
	ToMonthlyScheduleResponsePtrOutputWithContext(context.Context) MonthlyScheduleResponsePtrOutput
}

type monthlyScheduleResponsePtrType MonthlyScheduleResponseArgs

func MonthlyScheduleResponsePtr(v *MonthlyScheduleResponseArgs) MonthlyScheduleResponsePtrInput {
	return (*monthlyScheduleResponsePtrType)(v)
}

func (*monthlyScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlyScheduleResponse)(nil)).Elem()
}

func (i *monthlyScheduleResponsePtrType) ToMonthlyScheduleResponsePtrOutput() MonthlyScheduleResponsePtrOutput {
	return i.ToMonthlyScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *monthlyScheduleResponsePtrType) ToMonthlyScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyScheduleResponsePtrOutput)
}

type MonthlyScheduleResponseOutput struct{ *pulumi.OutputState }

func (MonthlyScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyScheduleResponse)(nil)).Elem()
}

func (o MonthlyScheduleResponseOutput) ToMonthlyScheduleResponseOutput() MonthlyScheduleResponseOutput {
	return o
}

func (o MonthlyScheduleResponseOutput) ToMonthlyScheduleResponseOutputWithContext(ctx context.Context) MonthlyScheduleResponseOutput {
	return o
}

func (o MonthlyScheduleResponseOutput) ToMonthlyScheduleResponsePtrOutput() MonthlyScheduleResponsePtrOutput {
	return o.ToMonthlyScheduleResponsePtrOutputWithContext(context.Background())
}

func (o MonthlyScheduleResponseOutput) ToMonthlyScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonthlyScheduleResponse) *MonthlyScheduleResponse {
		return &v
	}).(MonthlyScheduleResponsePtrOutput)
}

func (o MonthlyScheduleResponseOutput) DaysOfMonth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonthlyScheduleResponse) *string { return v.DaysOfMonth }).(pulumi.StringPtrOutput)
}

func (o MonthlyScheduleResponseOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonthlyScheduleResponse) *int { return v.Hour }).(pulumi.IntPtrOutput)
}

func (o MonthlyScheduleResponseOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonthlyScheduleResponse) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

func (o MonthlyScheduleResponseOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonthlyScheduleResponse) *int { return v.SnapshotsToKeep }).(pulumi.IntPtrOutput)
}

func (o MonthlyScheduleResponseOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonthlyScheduleResponse) *float64 { return v.UsedBytes }).(pulumi.Float64PtrOutput)
}

type MonthlyScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (MonthlyScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlyScheduleResponse)(nil)).Elem()
}

func (o MonthlyScheduleResponsePtrOutput) ToMonthlyScheduleResponsePtrOutput() MonthlyScheduleResponsePtrOutput {
	return o
}

func (o MonthlyScheduleResponsePtrOutput) ToMonthlyScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyScheduleResponsePtrOutput {
	return o
}

func (o MonthlyScheduleResponsePtrOutput) Elem() MonthlyScheduleResponseOutput {
	return o.ApplyT(func(v *MonthlyScheduleResponse) MonthlyScheduleResponse {
		if v != nil {
			return *v
		}
		var ret MonthlyScheduleResponse
		return ret
	}).(MonthlyScheduleResponseOutput)
}

func (o MonthlyScheduleResponsePtrOutput) DaysOfMonth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonthlyScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.DaysOfMonth
	}).(pulumi.StringPtrOutput)
}

func (o MonthlyScheduleResponsePtrOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MonthlyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.Hour
	}).(pulumi.IntPtrOutput)
}

func (o MonthlyScheduleResponsePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MonthlyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

func (o MonthlyScheduleResponsePtrOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MonthlyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.SnapshotsToKeep
	}).(pulumi.IntPtrOutput)
}

func (o MonthlyScheduleResponsePtrOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonthlyScheduleResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.UsedBytes
	}).(pulumi.Float64PtrOutput)
}

type MountTargetPropertiesResponse struct {
	FileSystemId  string  `pulumi:"fileSystemId"`
	IpAddress     string  `pulumi:"ipAddress"`
	MountTargetId string  `pulumi:"mountTargetId"`
	SmbServerFqdn *string `pulumi:"smbServerFqdn"`
}





type MountTargetPropertiesResponseInput interface {
	pulumi.Input

	ToMountTargetPropertiesResponseOutput() MountTargetPropertiesResponseOutput
	ToMountTargetPropertiesResponseOutputWithContext(context.Context) MountTargetPropertiesResponseOutput
}

type MountTargetPropertiesResponseArgs struct {
	FileSystemId  pulumi.StringInput    `pulumi:"fileSystemId"`
	IpAddress     pulumi.StringInput    `pulumi:"ipAddress"`
	MountTargetId pulumi.StringInput    `pulumi:"mountTargetId"`
	SmbServerFqdn pulumi.StringPtrInput `pulumi:"smbServerFqdn"`
}

func (MountTargetPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MountTargetPropertiesResponse)(nil)).Elem()
}

func (i MountTargetPropertiesResponseArgs) ToMountTargetPropertiesResponseOutput() MountTargetPropertiesResponseOutput {
	return i.ToMountTargetPropertiesResponseOutputWithContext(context.Background())
}

func (i MountTargetPropertiesResponseArgs) ToMountTargetPropertiesResponseOutputWithContext(ctx context.Context) MountTargetPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountTargetPropertiesResponseOutput)
}





type MountTargetPropertiesResponseArrayInput interface {
	pulumi.Input

	ToMountTargetPropertiesResponseArrayOutput() MountTargetPropertiesResponseArrayOutput
	ToMountTargetPropertiesResponseArrayOutputWithContext(context.Context) MountTargetPropertiesResponseArrayOutput
}

type MountTargetPropertiesResponseArray []MountTargetPropertiesResponseInput

func (MountTargetPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MountTargetPropertiesResponse)(nil)).Elem()
}

func (i MountTargetPropertiesResponseArray) ToMountTargetPropertiesResponseArrayOutput() MountTargetPropertiesResponseArrayOutput {
	return i.ToMountTargetPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i MountTargetPropertiesResponseArray) ToMountTargetPropertiesResponseArrayOutputWithContext(ctx context.Context) MountTargetPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountTargetPropertiesResponseArrayOutput)
}

type MountTargetPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MountTargetPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MountTargetPropertiesResponse)(nil)).Elem()
}

func (o MountTargetPropertiesResponseOutput) ToMountTargetPropertiesResponseOutput() MountTargetPropertiesResponseOutput {
	return o
}

func (o MountTargetPropertiesResponseOutput) ToMountTargetPropertiesResponseOutputWithContext(ctx context.Context) MountTargetPropertiesResponseOutput {
	return o
}

func (o MountTargetPropertiesResponseOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) string { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o MountTargetPropertiesResponseOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o MountTargetPropertiesResponseOutput) MountTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) string { return v.MountTargetId }).(pulumi.StringOutput)
}

func (o MountTargetPropertiesResponseOutput) SmbServerFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) *string { return v.SmbServerFqdn }).(pulumi.StringPtrOutput)
}

type MountTargetPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (MountTargetPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MountTargetPropertiesResponse)(nil)).Elem()
}

func (o MountTargetPropertiesResponseArrayOutput) ToMountTargetPropertiesResponseArrayOutput() MountTargetPropertiesResponseArrayOutput {
	return o
}

func (o MountTargetPropertiesResponseArrayOutput) ToMountTargetPropertiesResponseArrayOutputWithContext(ctx context.Context) MountTargetPropertiesResponseArrayOutput {
	return o
}

func (o MountTargetPropertiesResponseArrayOutput) Index(i pulumi.IntInput) MountTargetPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MountTargetPropertiesResponse {
		return vs[0].([]MountTargetPropertiesResponse)[vs[1].(int)]
	}).(MountTargetPropertiesResponseOutput)
}

type PlacementKeyValuePairs struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}





type PlacementKeyValuePairsInput interface {
	pulumi.Input

	ToPlacementKeyValuePairsOutput() PlacementKeyValuePairsOutput
	ToPlacementKeyValuePairsOutputWithContext(context.Context) PlacementKeyValuePairsOutput
}

type PlacementKeyValuePairsArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (PlacementKeyValuePairsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementKeyValuePairs)(nil)).Elem()
}

func (i PlacementKeyValuePairsArgs) ToPlacementKeyValuePairsOutput() PlacementKeyValuePairsOutput {
	return i.ToPlacementKeyValuePairsOutputWithContext(context.Background())
}

func (i PlacementKeyValuePairsArgs) ToPlacementKeyValuePairsOutputWithContext(ctx context.Context) PlacementKeyValuePairsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementKeyValuePairsOutput)
}





type PlacementKeyValuePairsArrayInput interface {
	pulumi.Input

	ToPlacementKeyValuePairsArrayOutput() PlacementKeyValuePairsArrayOutput
	ToPlacementKeyValuePairsArrayOutputWithContext(context.Context) PlacementKeyValuePairsArrayOutput
}

type PlacementKeyValuePairsArray []PlacementKeyValuePairsInput

func (PlacementKeyValuePairsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlacementKeyValuePairs)(nil)).Elem()
}

func (i PlacementKeyValuePairsArray) ToPlacementKeyValuePairsArrayOutput() PlacementKeyValuePairsArrayOutput {
	return i.ToPlacementKeyValuePairsArrayOutputWithContext(context.Background())
}

func (i PlacementKeyValuePairsArray) ToPlacementKeyValuePairsArrayOutputWithContext(ctx context.Context) PlacementKeyValuePairsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementKeyValuePairsArrayOutput)
}

type PlacementKeyValuePairsOutput struct{ *pulumi.OutputState }

func (PlacementKeyValuePairsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementKeyValuePairs)(nil)).Elem()
}

func (o PlacementKeyValuePairsOutput) ToPlacementKeyValuePairsOutput() PlacementKeyValuePairsOutput {
	return o
}

func (o PlacementKeyValuePairsOutput) ToPlacementKeyValuePairsOutputWithContext(ctx context.Context) PlacementKeyValuePairsOutput {
	return o
}

func (o PlacementKeyValuePairsOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v PlacementKeyValuePairs) string { return v.Key }).(pulumi.StringOutput)
}

func (o PlacementKeyValuePairsOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v PlacementKeyValuePairs) string { return v.Value }).(pulumi.StringOutput)
}

type PlacementKeyValuePairsArrayOutput struct{ *pulumi.OutputState }

func (PlacementKeyValuePairsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlacementKeyValuePairs)(nil)).Elem()
}

func (o PlacementKeyValuePairsArrayOutput) ToPlacementKeyValuePairsArrayOutput() PlacementKeyValuePairsArrayOutput {
	return o
}

func (o PlacementKeyValuePairsArrayOutput) ToPlacementKeyValuePairsArrayOutputWithContext(ctx context.Context) PlacementKeyValuePairsArrayOutput {
	return o
}

func (o PlacementKeyValuePairsArrayOutput) Index(i pulumi.IntInput) PlacementKeyValuePairsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlacementKeyValuePairs {
		return vs[0].([]PlacementKeyValuePairs)[vs[1].(int)]
	}).(PlacementKeyValuePairsOutput)
}

type PlacementKeyValuePairsResponse struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}





type PlacementKeyValuePairsResponseInput interface {
	pulumi.Input

	ToPlacementKeyValuePairsResponseOutput() PlacementKeyValuePairsResponseOutput
	ToPlacementKeyValuePairsResponseOutputWithContext(context.Context) PlacementKeyValuePairsResponseOutput
}

type PlacementKeyValuePairsResponseArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (PlacementKeyValuePairsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementKeyValuePairsResponse)(nil)).Elem()
}

func (i PlacementKeyValuePairsResponseArgs) ToPlacementKeyValuePairsResponseOutput() PlacementKeyValuePairsResponseOutput {
	return i.ToPlacementKeyValuePairsResponseOutputWithContext(context.Background())
}

func (i PlacementKeyValuePairsResponseArgs) ToPlacementKeyValuePairsResponseOutputWithContext(ctx context.Context) PlacementKeyValuePairsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementKeyValuePairsResponseOutput)
}





type PlacementKeyValuePairsResponseArrayInput interface {
	pulumi.Input

	ToPlacementKeyValuePairsResponseArrayOutput() PlacementKeyValuePairsResponseArrayOutput
	ToPlacementKeyValuePairsResponseArrayOutputWithContext(context.Context) PlacementKeyValuePairsResponseArrayOutput
}

type PlacementKeyValuePairsResponseArray []PlacementKeyValuePairsResponseInput

func (PlacementKeyValuePairsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlacementKeyValuePairsResponse)(nil)).Elem()
}

func (i PlacementKeyValuePairsResponseArray) ToPlacementKeyValuePairsResponseArrayOutput() PlacementKeyValuePairsResponseArrayOutput {
	return i.ToPlacementKeyValuePairsResponseArrayOutputWithContext(context.Background())
}

func (i PlacementKeyValuePairsResponseArray) ToPlacementKeyValuePairsResponseArrayOutputWithContext(ctx context.Context) PlacementKeyValuePairsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementKeyValuePairsResponseArrayOutput)
}

type PlacementKeyValuePairsResponseOutput struct{ *pulumi.OutputState }

func (PlacementKeyValuePairsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementKeyValuePairsResponse)(nil)).Elem()
}

func (o PlacementKeyValuePairsResponseOutput) ToPlacementKeyValuePairsResponseOutput() PlacementKeyValuePairsResponseOutput {
	return o
}

func (o PlacementKeyValuePairsResponseOutput) ToPlacementKeyValuePairsResponseOutputWithContext(ctx context.Context) PlacementKeyValuePairsResponseOutput {
	return o
}

func (o PlacementKeyValuePairsResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v PlacementKeyValuePairsResponse) string { return v.Key }).(pulumi.StringOutput)
}

func (o PlacementKeyValuePairsResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v PlacementKeyValuePairsResponse) string { return v.Value }).(pulumi.StringOutput)
}

type PlacementKeyValuePairsResponseArrayOutput struct{ *pulumi.OutputState }

func (PlacementKeyValuePairsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlacementKeyValuePairsResponse)(nil)).Elem()
}

func (o PlacementKeyValuePairsResponseArrayOutput) ToPlacementKeyValuePairsResponseArrayOutput() PlacementKeyValuePairsResponseArrayOutput {
	return o
}

func (o PlacementKeyValuePairsResponseArrayOutput) ToPlacementKeyValuePairsResponseArrayOutputWithContext(ctx context.Context) PlacementKeyValuePairsResponseArrayOutput {
	return o
}

func (o PlacementKeyValuePairsResponseArrayOutput) Index(i pulumi.IntInput) PlacementKeyValuePairsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlacementKeyValuePairsResponse {
		return vs[0].([]PlacementKeyValuePairsResponse)[vs[1].(int)]
	}).(PlacementKeyValuePairsResponseOutput)
}

type ReplicationObject struct {
	EndpointType           *string `pulumi:"endpointType"`
	RemoteVolumeRegion     *string `pulumi:"remoteVolumeRegion"`
	RemoteVolumeResourceId string  `pulumi:"remoteVolumeResourceId"`
	ReplicationId          *string `pulumi:"replicationId"`
	ReplicationSchedule    *string `pulumi:"replicationSchedule"`
}





type ReplicationObjectInput interface {
	pulumi.Input

	ToReplicationObjectOutput() ReplicationObjectOutput
	ToReplicationObjectOutputWithContext(context.Context) ReplicationObjectOutput
}

type ReplicationObjectArgs struct {
	EndpointType           pulumi.StringPtrInput `pulumi:"endpointType"`
	RemoteVolumeRegion     pulumi.StringPtrInput `pulumi:"remoteVolumeRegion"`
	RemoteVolumeResourceId pulumi.StringInput    `pulumi:"remoteVolumeResourceId"`
	ReplicationId          pulumi.StringPtrInput `pulumi:"replicationId"`
	ReplicationSchedule    pulumi.StringPtrInput `pulumi:"replicationSchedule"`
}

func (ReplicationObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationObject)(nil)).Elem()
}

func (i ReplicationObjectArgs) ToReplicationObjectOutput() ReplicationObjectOutput {
	return i.ToReplicationObjectOutputWithContext(context.Background())
}

func (i ReplicationObjectArgs) ToReplicationObjectOutputWithContext(ctx context.Context) ReplicationObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectOutput)
}

func (i ReplicationObjectArgs) ToReplicationObjectPtrOutput() ReplicationObjectPtrOutput {
	return i.ToReplicationObjectPtrOutputWithContext(context.Background())
}

func (i ReplicationObjectArgs) ToReplicationObjectPtrOutputWithContext(ctx context.Context) ReplicationObjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectOutput).ToReplicationObjectPtrOutputWithContext(ctx)
}









type ReplicationObjectPtrInput interface {
	pulumi.Input

	ToReplicationObjectPtrOutput() ReplicationObjectPtrOutput
	ToReplicationObjectPtrOutputWithContext(context.Context) ReplicationObjectPtrOutput
}

type replicationObjectPtrType ReplicationObjectArgs

func ReplicationObjectPtr(v *ReplicationObjectArgs) ReplicationObjectPtrInput {
	return (*replicationObjectPtrType)(v)
}

func (*replicationObjectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationObject)(nil)).Elem()
}

func (i *replicationObjectPtrType) ToReplicationObjectPtrOutput() ReplicationObjectPtrOutput {
	return i.ToReplicationObjectPtrOutputWithContext(context.Background())
}

func (i *replicationObjectPtrType) ToReplicationObjectPtrOutputWithContext(ctx context.Context) ReplicationObjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectPtrOutput)
}

type ReplicationObjectOutput struct{ *pulumi.OutputState }

func (ReplicationObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationObject)(nil)).Elem()
}

func (o ReplicationObjectOutput) ToReplicationObjectOutput() ReplicationObjectOutput {
	return o
}

func (o ReplicationObjectOutput) ToReplicationObjectOutputWithContext(ctx context.Context) ReplicationObjectOutput {
	return o
}

func (o ReplicationObjectOutput) ToReplicationObjectPtrOutput() ReplicationObjectPtrOutput {
	return o.ToReplicationObjectPtrOutputWithContext(context.Background())
}

func (o ReplicationObjectOutput) ToReplicationObjectPtrOutputWithContext(ctx context.Context) ReplicationObjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationObject) *ReplicationObject {
		return &v
	}).(ReplicationObjectPtrOutput)
}

func (o ReplicationObjectOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObject) *string { return v.EndpointType }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectOutput) RemoteVolumeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObject) *string { return v.RemoteVolumeRegion }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectOutput) RemoteVolumeResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicationObject) string { return v.RemoteVolumeResourceId }).(pulumi.StringOutput)
}

func (o ReplicationObjectOutput) ReplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObject) *string { return v.ReplicationId }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectOutput) ReplicationSchedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObject) *string { return v.ReplicationSchedule }).(pulumi.StringPtrOutput)
}

type ReplicationObjectPtrOutput struct{ *pulumi.OutputState }

func (ReplicationObjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationObject)(nil)).Elem()
}

func (o ReplicationObjectPtrOutput) ToReplicationObjectPtrOutput() ReplicationObjectPtrOutput {
	return o
}

func (o ReplicationObjectPtrOutput) ToReplicationObjectPtrOutputWithContext(ctx context.Context) ReplicationObjectPtrOutput {
	return o
}

func (o ReplicationObjectPtrOutput) Elem() ReplicationObjectOutput {
	return o.ApplyT(func(v *ReplicationObject) ReplicationObject {
		if v != nil {
			return *v
		}
		var ret ReplicationObject
		return ret
	}).(ReplicationObjectOutput)
}

func (o ReplicationObjectPtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObject) *string {
		if v == nil {
			return nil
		}
		return v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectPtrOutput) RemoteVolumeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObject) *string {
		if v == nil {
			return nil
		}
		return v.RemoteVolumeRegion
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectPtrOutput) RemoteVolumeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObject) *string {
		if v == nil {
			return nil
		}
		return &v.RemoteVolumeResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectPtrOutput) ReplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObject) *string {
		if v == nil {
			return nil
		}
		return v.ReplicationId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectPtrOutput) ReplicationSchedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObject) *string {
		if v == nil {
			return nil
		}
		return v.ReplicationSchedule
	}).(pulumi.StringPtrOutput)
}

type ReplicationObjectResponse struct {
	EndpointType           *string `pulumi:"endpointType"`
	RemoteVolumeRegion     *string `pulumi:"remoteVolumeRegion"`
	RemoteVolumeResourceId string  `pulumi:"remoteVolumeResourceId"`
	ReplicationId          *string `pulumi:"replicationId"`
	ReplicationSchedule    *string `pulumi:"replicationSchedule"`
}





type ReplicationObjectResponseInput interface {
	pulumi.Input

	ToReplicationObjectResponseOutput() ReplicationObjectResponseOutput
	ToReplicationObjectResponseOutputWithContext(context.Context) ReplicationObjectResponseOutput
}

type ReplicationObjectResponseArgs struct {
	EndpointType           pulumi.StringPtrInput `pulumi:"endpointType"`
	RemoteVolumeRegion     pulumi.StringPtrInput `pulumi:"remoteVolumeRegion"`
	RemoteVolumeResourceId pulumi.StringInput    `pulumi:"remoteVolumeResourceId"`
	ReplicationId          pulumi.StringPtrInput `pulumi:"replicationId"`
	ReplicationSchedule    pulumi.StringPtrInput `pulumi:"replicationSchedule"`
}

func (ReplicationObjectResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationObjectResponse)(nil)).Elem()
}

func (i ReplicationObjectResponseArgs) ToReplicationObjectResponseOutput() ReplicationObjectResponseOutput {
	return i.ToReplicationObjectResponseOutputWithContext(context.Background())
}

func (i ReplicationObjectResponseArgs) ToReplicationObjectResponseOutputWithContext(ctx context.Context) ReplicationObjectResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectResponseOutput)
}

func (i ReplicationObjectResponseArgs) ToReplicationObjectResponsePtrOutput() ReplicationObjectResponsePtrOutput {
	return i.ToReplicationObjectResponsePtrOutputWithContext(context.Background())
}

func (i ReplicationObjectResponseArgs) ToReplicationObjectResponsePtrOutputWithContext(ctx context.Context) ReplicationObjectResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectResponseOutput).ToReplicationObjectResponsePtrOutputWithContext(ctx)
}









type ReplicationObjectResponsePtrInput interface {
	pulumi.Input

	ToReplicationObjectResponsePtrOutput() ReplicationObjectResponsePtrOutput
	ToReplicationObjectResponsePtrOutputWithContext(context.Context) ReplicationObjectResponsePtrOutput
}

type replicationObjectResponsePtrType ReplicationObjectResponseArgs

func ReplicationObjectResponsePtr(v *ReplicationObjectResponseArgs) ReplicationObjectResponsePtrInput {
	return (*replicationObjectResponsePtrType)(v)
}

func (*replicationObjectResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationObjectResponse)(nil)).Elem()
}

func (i *replicationObjectResponsePtrType) ToReplicationObjectResponsePtrOutput() ReplicationObjectResponsePtrOutput {
	return i.ToReplicationObjectResponsePtrOutputWithContext(context.Background())
}

func (i *replicationObjectResponsePtrType) ToReplicationObjectResponsePtrOutputWithContext(ctx context.Context) ReplicationObjectResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectResponsePtrOutput)
}

type ReplicationObjectResponseOutput struct{ *pulumi.OutputState }

func (ReplicationObjectResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationObjectResponse)(nil)).Elem()
}

func (o ReplicationObjectResponseOutput) ToReplicationObjectResponseOutput() ReplicationObjectResponseOutput {
	return o
}

func (o ReplicationObjectResponseOutput) ToReplicationObjectResponseOutputWithContext(ctx context.Context) ReplicationObjectResponseOutput {
	return o
}

func (o ReplicationObjectResponseOutput) ToReplicationObjectResponsePtrOutput() ReplicationObjectResponsePtrOutput {
	return o.ToReplicationObjectResponsePtrOutputWithContext(context.Background())
}

func (o ReplicationObjectResponseOutput) ToReplicationObjectResponsePtrOutputWithContext(ctx context.Context) ReplicationObjectResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationObjectResponse) *ReplicationObjectResponse {
		return &v
	}).(ReplicationObjectResponsePtrOutput)
}

func (o ReplicationObjectResponseOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObjectResponse) *string { return v.EndpointType }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponseOutput) RemoteVolumeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObjectResponse) *string { return v.RemoteVolumeRegion }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponseOutput) RemoteVolumeResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicationObjectResponse) string { return v.RemoteVolumeResourceId }).(pulumi.StringOutput)
}

func (o ReplicationObjectResponseOutput) ReplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObjectResponse) *string { return v.ReplicationId }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponseOutput) ReplicationSchedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObjectResponse) *string { return v.ReplicationSchedule }).(pulumi.StringPtrOutput)
}

type ReplicationObjectResponsePtrOutput struct{ *pulumi.OutputState }

func (ReplicationObjectResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationObjectResponse)(nil)).Elem()
}

func (o ReplicationObjectResponsePtrOutput) ToReplicationObjectResponsePtrOutput() ReplicationObjectResponsePtrOutput {
	return o
}

func (o ReplicationObjectResponsePtrOutput) ToReplicationObjectResponsePtrOutputWithContext(ctx context.Context) ReplicationObjectResponsePtrOutput {
	return o
}

func (o ReplicationObjectResponsePtrOutput) Elem() ReplicationObjectResponseOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) ReplicationObjectResponse {
		if v != nil {
			return *v
		}
		var ret ReplicationObjectResponse
		return ret
	}).(ReplicationObjectResponseOutput)
}

func (o ReplicationObjectResponsePtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponsePtrOutput) RemoteVolumeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.RemoteVolumeRegion
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponsePtrOutput) RemoteVolumeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RemoteVolumeResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponsePtrOutput) ReplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplicationId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponsePtrOutput) ReplicationSchedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplicationSchedule
	}).(pulumi.StringPtrOutput)
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

type VolumeBackupProperties struct {
	BackupEnabled  *bool   `pulumi:"backupEnabled"`
	BackupPolicyId *string `pulumi:"backupPolicyId"`
	PolicyEnforced *bool   `pulumi:"policyEnforced"`
	VaultId        *string `pulumi:"vaultId"`
}





type VolumeBackupPropertiesInput interface {
	pulumi.Input

	ToVolumeBackupPropertiesOutput() VolumeBackupPropertiesOutput
	ToVolumeBackupPropertiesOutputWithContext(context.Context) VolumeBackupPropertiesOutput
}

type VolumeBackupPropertiesArgs struct {
	BackupEnabled  pulumi.BoolPtrInput   `pulumi:"backupEnabled"`
	BackupPolicyId pulumi.StringPtrInput `pulumi:"backupPolicyId"`
	PolicyEnforced pulumi.BoolPtrInput   `pulumi:"policyEnforced"`
	VaultId        pulumi.StringPtrInput `pulumi:"vaultId"`
}

func (VolumeBackupPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeBackupProperties)(nil)).Elem()
}

func (i VolumeBackupPropertiesArgs) ToVolumeBackupPropertiesOutput() VolumeBackupPropertiesOutput {
	return i.ToVolumeBackupPropertiesOutputWithContext(context.Background())
}

func (i VolumeBackupPropertiesArgs) ToVolumeBackupPropertiesOutputWithContext(ctx context.Context) VolumeBackupPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupPropertiesOutput)
}

func (i VolumeBackupPropertiesArgs) ToVolumeBackupPropertiesPtrOutput() VolumeBackupPropertiesPtrOutput {
	return i.ToVolumeBackupPropertiesPtrOutputWithContext(context.Background())
}

func (i VolumeBackupPropertiesArgs) ToVolumeBackupPropertiesPtrOutputWithContext(ctx context.Context) VolumeBackupPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupPropertiesOutput).ToVolumeBackupPropertiesPtrOutputWithContext(ctx)
}









type VolumeBackupPropertiesPtrInput interface {
	pulumi.Input

	ToVolumeBackupPropertiesPtrOutput() VolumeBackupPropertiesPtrOutput
	ToVolumeBackupPropertiesPtrOutputWithContext(context.Context) VolumeBackupPropertiesPtrOutput
}

type volumeBackupPropertiesPtrType VolumeBackupPropertiesArgs

func VolumeBackupPropertiesPtr(v *VolumeBackupPropertiesArgs) VolumeBackupPropertiesPtrInput {
	return (*volumeBackupPropertiesPtrType)(v)
}

func (*volumeBackupPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeBackupProperties)(nil)).Elem()
}

func (i *volumeBackupPropertiesPtrType) ToVolumeBackupPropertiesPtrOutput() VolumeBackupPropertiesPtrOutput {
	return i.ToVolumeBackupPropertiesPtrOutputWithContext(context.Background())
}

func (i *volumeBackupPropertiesPtrType) ToVolumeBackupPropertiesPtrOutputWithContext(ctx context.Context) VolumeBackupPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupPropertiesPtrOutput)
}

type VolumeBackupPropertiesOutput struct{ *pulumi.OutputState }

func (VolumeBackupPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeBackupProperties)(nil)).Elem()
}

func (o VolumeBackupPropertiesOutput) ToVolumeBackupPropertiesOutput() VolumeBackupPropertiesOutput {
	return o
}

func (o VolumeBackupPropertiesOutput) ToVolumeBackupPropertiesOutputWithContext(ctx context.Context) VolumeBackupPropertiesOutput {
	return o
}

func (o VolumeBackupPropertiesOutput) ToVolumeBackupPropertiesPtrOutput() VolumeBackupPropertiesPtrOutput {
	return o.ToVolumeBackupPropertiesPtrOutputWithContext(context.Background())
}

func (o VolumeBackupPropertiesOutput) ToVolumeBackupPropertiesPtrOutputWithContext(ctx context.Context) VolumeBackupPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeBackupProperties) *VolumeBackupProperties {
		return &v
	}).(VolumeBackupPropertiesPtrOutput)
}

func (o VolumeBackupPropertiesOutput) BackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeBackupProperties) *bool { return v.BackupEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeBackupPropertiesOutput) BackupPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeBackupProperties) *string { return v.BackupPolicyId }).(pulumi.StringPtrOutput)
}

func (o VolumeBackupPropertiesOutput) PolicyEnforced() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeBackupProperties) *bool { return v.PolicyEnforced }).(pulumi.BoolPtrOutput)
}

func (o VolumeBackupPropertiesOutput) VaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeBackupProperties) *string { return v.VaultId }).(pulumi.StringPtrOutput)
}

type VolumeBackupPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VolumeBackupPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeBackupProperties)(nil)).Elem()
}

func (o VolumeBackupPropertiesPtrOutput) ToVolumeBackupPropertiesPtrOutput() VolumeBackupPropertiesPtrOutput {
	return o
}

func (o VolumeBackupPropertiesPtrOutput) ToVolumeBackupPropertiesPtrOutputWithContext(ctx context.Context) VolumeBackupPropertiesPtrOutput {
	return o
}

func (o VolumeBackupPropertiesPtrOutput) Elem() VolumeBackupPropertiesOutput {
	return o.ApplyT(func(v *VolumeBackupProperties) VolumeBackupProperties {
		if v != nil {
			return *v
		}
		var ret VolumeBackupProperties
		return ret
	}).(VolumeBackupPropertiesOutput)
}

func (o VolumeBackupPropertiesPtrOutput) BackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VolumeBackupProperties) *bool {
		if v == nil {
			return nil
		}
		return v.BackupEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o VolumeBackupPropertiesPtrOutput) BackupPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeBackupProperties) *string {
		if v == nil {
			return nil
		}
		return v.BackupPolicyId
	}).(pulumi.StringPtrOutput)
}

func (o VolumeBackupPropertiesPtrOutput) PolicyEnforced() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VolumeBackupProperties) *bool {
		if v == nil {
			return nil
		}
		return v.PolicyEnforced
	}).(pulumi.BoolPtrOutput)
}

func (o VolumeBackupPropertiesPtrOutput) VaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeBackupProperties) *string {
		if v == nil {
			return nil
		}
		return v.VaultId
	}).(pulumi.StringPtrOutput)
}

type VolumeBackupPropertiesResponse struct {
	BackupEnabled  *bool   `pulumi:"backupEnabled"`
	BackupPolicyId *string `pulumi:"backupPolicyId"`
	PolicyEnforced *bool   `pulumi:"policyEnforced"`
	VaultId        *string `pulumi:"vaultId"`
}





type VolumeBackupPropertiesResponseInput interface {
	pulumi.Input

	ToVolumeBackupPropertiesResponseOutput() VolumeBackupPropertiesResponseOutput
	ToVolumeBackupPropertiesResponseOutputWithContext(context.Context) VolumeBackupPropertiesResponseOutput
}

type VolumeBackupPropertiesResponseArgs struct {
	BackupEnabled  pulumi.BoolPtrInput   `pulumi:"backupEnabled"`
	BackupPolicyId pulumi.StringPtrInput `pulumi:"backupPolicyId"`
	PolicyEnforced pulumi.BoolPtrInput   `pulumi:"policyEnforced"`
	VaultId        pulumi.StringPtrInput `pulumi:"vaultId"`
}

func (VolumeBackupPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeBackupPropertiesResponse)(nil)).Elem()
}

func (i VolumeBackupPropertiesResponseArgs) ToVolumeBackupPropertiesResponseOutput() VolumeBackupPropertiesResponseOutput {
	return i.ToVolumeBackupPropertiesResponseOutputWithContext(context.Background())
}

func (i VolumeBackupPropertiesResponseArgs) ToVolumeBackupPropertiesResponseOutputWithContext(ctx context.Context) VolumeBackupPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupPropertiesResponseOutput)
}

func (i VolumeBackupPropertiesResponseArgs) ToVolumeBackupPropertiesResponsePtrOutput() VolumeBackupPropertiesResponsePtrOutput {
	return i.ToVolumeBackupPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i VolumeBackupPropertiesResponseArgs) ToVolumeBackupPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeBackupPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupPropertiesResponseOutput).ToVolumeBackupPropertiesResponsePtrOutputWithContext(ctx)
}









type VolumeBackupPropertiesResponsePtrInput interface {
	pulumi.Input

	ToVolumeBackupPropertiesResponsePtrOutput() VolumeBackupPropertiesResponsePtrOutput
	ToVolumeBackupPropertiesResponsePtrOutputWithContext(context.Context) VolumeBackupPropertiesResponsePtrOutput
}

type volumeBackupPropertiesResponsePtrType VolumeBackupPropertiesResponseArgs

func VolumeBackupPropertiesResponsePtr(v *VolumeBackupPropertiesResponseArgs) VolumeBackupPropertiesResponsePtrInput {
	return (*volumeBackupPropertiesResponsePtrType)(v)
}

func (*volumeBackupPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeBackupPropertiesResponse)(nil)).Elem()
}

func (i *volumeBackupPropertiesResponsePtrType) ToVolumeBackupPropertiesResponsePtrOutput() VolumeBackupPropertiesResponsePtrOutput {
	return i.ToVolumeBackupPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *volumeBackupPropertiesResponsePtrType) ToVolumeBackupPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeBackupPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupPropertiesResponsePtrOutput)
}

type VolumeBackupPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VolumeBackupPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeBackupPropertiesResponse)(nil)).Elem()
}

func (o VolumeBackupPropertiesResponseOutput) ToVolumeBackupPropertiesResponseOutput() VolumeBackupPropertiesResponseOutput {
	return o
}

func (o VolumeBackupPropertiesResponseOutput) ToVolumeBackupPropertiesResponseOutputWithContext(ctx context.Context) VolumeBackupPropertiesResponseOutput {
	return o
}

func (o VolumeBackupPropertiesResponseOutput) ToVolumeBackupPropertiesResponsePtrOutput() VolumeBackupPropertiesResponsePtrOutput {
	return o.ToVolumeBackupPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o VolumeBackupPropertiesResponseOutput) ToVolumeBackupPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeBackupPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeBackupPropertiesResponse) *VolumeBackupPropertiesResponse {
		return &v
	}).(VolumeBackupPropertiesResponsePtrOutput)
}

func (o VolumeBackupPropertiesResponseOutput) BackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeBackupPropertiesResponse) *bool { return v.BackupEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeBackupPropertiesResponseOutput) BackupPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeBackupPropertiesResponse) *string { return v.BackupPolicyId }).(pulumi.StringPtrOutput)
}

func (o VolumeBackupPropertiesResponseOutput) PolicyEnforced() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeBackupPropertiesResponse) *bool { return v.PolicyEnforced }).(pulumi.BoolPtrOutput)
}

func (o VolumeBackupPropertiesResponseOutput) VaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeBackupPropertiesResponse) *string { return v.VaultId }).(pulumi.StringPtrOutput)
}

type VolumeBackupPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VolumeBackupPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeBackupPropertiesResponse)(nil)).Elem()
}

func (o VolumeBackupPropertiesResponsePtrOutput) ToVolumeBackupPropertiesResponsePtrOutput() VolumeBackupPropertiesResponsePtrOutput {
	return o
}

func (o VolumeBackupPropertiesResponsePtrOutput) ToVolumeBackupPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeBackupPropertiesResponsePtrOutput {
	return o
}

func (o VolumeBackupPropertiesResponsePtrOutput) Elem() VolumeBackupPropertiesResponseOutput {
	return o.ApplyT(func(v *VolumeBackupPropertiesResponse) VolumeBackupPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VolumeBackupPropertiesResponse
		return ret
	}).(VolumeBackupPropertiesResponseOutput)
}

func (o VolumeBackupPropertiesResponsePtrOutput) BackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VolumeBackupPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.BackupEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o VolumeBackupPropertiesResponsePtrOutput) BackupPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeBackupPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.BackupPolicyId
	}).(pulumi.StringPtrOutput)
}

func (o VolumeBackupPropertiesResponsePtrOutput) PolicyEnforced() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VolumeBackupPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PolicyEnforced
	}).(pulumi.BoolPtrOutput)
}

func (o VolumeBackupPropertiesResponsePtrOutput) VaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeBackupPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VaultId
	}).(pulumi.StringPtrOutput)
}

type VolumeBackupsResponse struct {
	BackupsCount  *int    `pulumi:"backupsCount"`
	PolicyEnabled *bool   `pulumi:"policyEnabled"`
	VolumeName    *string `pulumi:"volumeName"`
}





type VolumeBackupsResponseInput interface {
	pulumi.Input

	ToVolumeBackupsResponseOutput() VolumeBackupsResponseOutput
	ToVolumeBackupsResponseOutputWithContext(context.Context) VolumeBackupsResponseOutput
}

type VolumeBackupsResponseArgs struct {
	BackupsCount  pulumi.IntPtrInput    `pulumi:"backupsCount"`
	PolicyEnabled pulumi.BoolPtrInput   `pulumi:"policyEnabled"`
	VolumeName    pulumi.StringPtrInput `pulumi:"volumeName"`
}

func (VolumeBackupsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeBackupsResponse)(nil)).Elem()
}

func (i VolumeBackupsResponseArgs) ToVolumeBackupsResponseOutput() VolumeBackupsResponseOutput {
	return i.ToVolumeBackupsResponseOutputWithContext(context.Background())
}

func (i VolumeBackupsResponseArgs) ToVolumeBackupsResponseOutputWithContext(ctx context.Context) VolumeBackupsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupsResponseOutput)
}





type VolumeBackupsResponseArrayInput interface {
	pulumi.Input

	ToVolumeBackupsResponseArrayOutput() VolumeBackupsResponseArrayOutput
	ToVolumeBackupsResponseArrayOutputWithContext(context.Context) VolumeBackupsResponseArrayOutput
}

type VolumeBackupsResponseArray []VolumeBackupsResponseInput

func (VolumeBackupsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeBackupsResponse)(nil)).Elem()
}

func (i VolumeBackupsResponseArray) ToVolumeBackupsResponseArrayOutput() VolumeBackupsResponseArrayOutput {
	return i.ToVolumeBackupsResponseArrayOutputWithContext(context.Background())
}

func (i VolumeBackupsResponseArray) ToVolumeBackupsResponseArrayOutputWithContext(ctx context.Context) VolumeBackupsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupsResponseArrayOutput)
}

type VolumeBackupsResponseOutput struct{ *pulumi.OutputState }

func (VolumeBackupsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeBackupsResponse)(nil)).Elem()
}

func (o VolumeBackupsResponseOutput) ToVolumeBackupsResponseOutput() VolumeBackupsResponseOutput {
	return o
}

func (o VolumeBackupsResponseOutput) ToVolumeBackupsResponseOutputWithContext(ctx context.Context) VolumeBackupsResponseOutput {
	return o
}

func (o VolumeBackupsResponseOutput) BackupsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VolumeBackupsResponse) *int { return v.BackupsCount }).(pulumi.IntPtrOutput)
}

func (o VolumeBackupsResponseOutput) PolicyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeBackupsResponse) *bool { return v.PolicyEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeBackupsResponseOutput) VolumeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeBackupsResponse) *string { return v.VolumeName }).(pulumi.StringPtrOutput)
}

type VolumeBackupsResponseArrayOutput struct{ *pulumi.OutputState }

func (VolumeBackupsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeBackupsResponse)(nil)).Elem()
}

func (o VolumeBackupsResponseArrayOutput) ToVolumeBackupsResponseArrayOutput() VolumeBackupsResponseArrayOutput {
	return o
}

func (o VolumeBackupsResponseArrayOutput) ToVolumeBackupsResponseArrayOutputWithContext(ctx context.Context) VolumeBackupsResponseArrayOutput {
	return o
}

func (o VolumeBackupsResponseArrayOutput) Index(i pulumi.IntInput) VolumeBackupsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeBackupsResponse {
		return vs[0].([]VolumeBackupsResponse)[vs[1].(int)]
	}).(VolumeBackupsResponseOutput)
}

type VolumeGroupMetaData struct {
	ApplicationIdentifier *string                  `pulumi:"applicationIdentifier"`
	ApplicationType       *string                  `pulumi:"applicationType"`
	DeploymentSpecId      *string                  `pulumi:"deploymentSpecId"`
	GlobalPlacementRules  []PlacementKeyValuePairs `pulumi:"globalPlacementRules"`
	GroupDescription      *string                  `pulumi:"groupDescription"`
}





type VolumeGroupMetaDataInput interface {
	pulumi.Input

	ToVolumeGroupMetaDataOutput() VolumeGroupMetaDataOutput
	ToVolumeGroupMetaDataOutputWithContext(context.Context) VolumeGroupMetaDataOutput
}

type VolumeGroupMetaDataArgs struct {
	ApplicationIdentifier pulumi.StringPtrInput            `pulumi:"applicationIdentifier"`
	ApplicationType       pulumi.StringPtrInput            `pulumi:"applicationType"`
	DeploymentSpecId      pulumi.StringPtrInput            `pulumi:"deploymentSpecId"`
	GlobalPlacementRules  PlacementKeyValuePairsArrayInput `pulumi:"globalPlacementRules"`
	GroupDescription      pulumi.StringPtrInput            `pulumi:"groupDescription"`
}

func (VolumeGroupMetaDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeGroupMetaData)(nil)).Elem()
}

func (i VolumeGroupMetaDataArgs) ToVolumeGroupMetaDataOutput() VolumeGroupMetaDataOutput {
	return i.ToVolumeGroupMetaDataOutputWithContext(context.Background())
}

func (i VolumeGroupMetaDataArgs) ToVolumeGroupMetaDataOutputWithContext(ctx context.Context) VolumeGroupMetaDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupMetaDataOutput)
}

func (i VolumeGroupMetaDataArgs) ToVolumeGroupMetaDataPtrOutput() VolumeGroupMetaDataPtrOutput {
	return i.ToVolumeGroupMetaDataPtrOutputWithContext(context.Background())
}

func (i VolumeGroupMetaDataArgs) ToVolumeGroupMetaDataPtrOutputWithContext(ctx context.Context) VolumeGroupMetaDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupMetaDataOutput).ToVolumeGroupMetaDataPtrOutputWithContext(ctx)
}









type VolumeGroupMetaDataPtrInput interface {
	pulumi.Input

	ToVolumeGroupMetaDataPtrOutput() VolumeGroupMetaDataPtrOutput
	ToVolumeGroupMetaDataPtrOutputWithContext(context.Context) VolumeGroupMetaDataPtrOutput
}

type volumeGroupMetaDataPtrType VolumeGroupMetaDataArgs

func VolumeGroupMetaDataPtr(v *VolumeGroupMetaDataArgs) VolumeGroupMetaDataPtrInput {
	return (*volumeGroupMetaDataPtrType)(v)
}

func (*volumeGroupMetaDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeGroupMetaData)(nil)).Elem()
}

func (i *volumeGroupMetaDataPtrType) ToVolumeGroupMetaDataPtrOutput() VolumeGroupMetaDataPtrOutput {
	return i.ToVolumeGroupMetaDataPtrOutputWithContext(context.Background())
}

func (i *volumeGroupMetaDataPtrType) ToVolumeGroupMetaDataPtrOutputWithContext(ctx context.Context) VolumeGroupMetaDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupMetaDataPtrOutput)
}

type VolumeGroupMetaDataOutput struct{ *pulumi.OutputState }

func (VolumeGroupMetaDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeGroupMetaData)(nil)).Elem()
}

func (o VolumeGroupMetaDataOutput) ToVolumeGroupMetaDataOutput() VolumeGroupMetaDataOutput {
	return o
}

func (o VolumeGroupMetaDataOutput) ToVolumeGroupMetaDataOutputWithContext(ctx context.Context) VolumeGroupMetaDataOutput {
	return o
}

func (o VolumeGroupMetaDataOutput) ToVolumeGroupMetaDataPtrOutput() VolumeGroupMetaDataPtrOutput {
	return o.ToVolumeGroupMetaDataPtrOutputWithContext(context.Background())
}

func (o VolumeGroupMetaDataOutput) ToVolumeGroupMetaDataPtrOutputWithContext(ctx context.Context) VolumeGroupMetaDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeGroupMetaData) *VolumeGroupMetaData {
		return &v
	}).(VolumeGroupMetaDataPtrOutput)
}

func (o VolumeGroupMetaDataOutput) ApplicationIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupMetaData) *string { return v.ApplicationIdentifier }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataOutput) ApplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupMetaData) *string { return v.ApplicationType }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataOutput) DeploymentSpecId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupMetaData) *string { return v.DeploymentSpecId }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataOutput) GlobalPlacementRules() PlacementKeyValuePairsArrayOutput {
	return o.ApplyT(func(v VolumeGroupMetaData) []PlacementKeyValuePairs { return v.GlobalPlacementRules }).(PlacementKeyValuePairsArrayOutput)
}

func (o VolumeGroupMetaDataOutput) GroupDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupMetaData) *string { return v.GroupDescription }).(pulumi.StringPtrOutput)
}

type VolumeGroupMetaDataPtrOutput struct{ *pulumi.OutputState }

func (VolumeGroupMetaDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeGroupMetaData)(nil)).Elem()
}

func (o VolumeGroupMetaDataPtrOutput) ToVolumeGroupMetaDataPtrOutput() VolumeGroupMetaDataPtrOutput {
	return o
}

func (o VolumeGroupMetaDataPtrOutput) ToVolumeGroupMetaDataPtrOutputWithContext(ctx context.Context) VolumeGroupMetaDataPtrOutput {
	return o
}

func (o VolumeGroupMetaDataPtrOutput) Elem() VolumeGroupMetaDataOutput {
	return o.ApplyT(func(v *VolumeGroupMetaData) VolumeGroupMetaData {
		if v != nil {
			return *v
		}
		var ret VolumeGroupMetaData
		return ret
	}).(VolumeGroupMetaDataOutput)
}

func (o VolumeGroupMetaDataPtrOutput) ApplicationIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroupMetaData) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataPtrOutput) ApplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroupMetaData) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationType
	}).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataPtrOutput) DeploymentSpecId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroupMetaData) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentSpecId
	}).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataPtrOutput) GlobalPlacementRules() PlacementKeyValuePairsArrayOutput {
	return o.ApplyT(func(v *VolumeGroupMetaData) []PlacementKeyValuePairs {
		if v == nil {
			return nil
		}
		return v.GlobalPlacementRules
	}).(PlacementKeyValuePairsArrayOutput)
}

func (o VolumeGroupMetaDataPtrOutput) GroupDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroupMetaData) *string {
		if v == nil {
			return nil
		}
		return v.GroupDescription
	}).(pulumi.StringPtrOutput)
}

type VolumeGroupMetaDataResponse struct {
	ApplicationIdentifier *string                          `pulumi:"applicationIdentifier"`
	ApplicationType       *string                          `pulumi:"applicationType"`
	DeploymentSpecId      *string                          `pulumi:"deploymentSpecId"`
	GlobalPlacementRules  []PlacementKeyValuePairsResponse `pulumi:"globalPlacementRules"`
	GroupDescription      *string                          `pulumi:"groupDescription"`
	VolumesCount          float64                          `pulumi:"volumesCount"`
}





type VolumeGroupMetaDataResponseInput interface {
	pulumi.Input

	ToVolumeGroupMetaDataResponseOutput() VolumeGroupMetaDataResponseOutput
	ToVolumeGroupMetaDataResponseOutputWithContext(context.Context) VolumeGroupMetaDataResponseOutput
}

type VolumeGroupMetaDataResponseArgs struct {
	ApplicationIdentifier pulumi.StringPtrInput                    `pulumi:"applicationIdentifier"`
	ApplicationType       pulumi.StringPtrInput                    `pulumi:"applicationType"`
	DeploymentSpecId      pulumi.StringPtrInput                    `pulumi:"deploymentSpecId"`
	GlobalPlacementRules  PlacementKeyValuePairsResponseArrayInput `pulumi:"globalPlacementRules"`
	GroupDescription      pulumi.StringPtrInput                    `pulumi:"groupDescription"`
	VolumesCount          pulumi.Float64Input                      `pulumi:"volumesCount"`
}

func (VolumeGroupMetaDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeGroupMetaDataResponse)(nil)).Elem()
}

func (i VolumeGroupMetaDataResponseArgs) ToVolumeGroupMetaDataResponseOutput() VolumeGroupMetaDataResponseOutput {
	return i.ToVolumeGroupMetaDataResponseOutputWithContext(context.Background())
}

func (i VolumeGroupMetaDataResponseArgs) ToVolumeGroupMetaDataResponseOutputWithContext(ctx context.Context) VolumeGroupMetaDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupMetaDataResponseOutput)
}

func (i VolumeGroupMetaDataResponseArgs) ToVolumeGroupMetaDataResponsePtrOutput() VolumeGroupMetaDataResponsePtrOutput {
	return i.ToVolumeGroupMetaDataResponsePtrOutputWithContext(context.Background())
}

func (i VolumeGroupMetaDataResponseArgs) ToVolumeGroupMetaDataResponsePtrOutputWithContext(ctx context.Context) VolumeGroupMetaDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupMetaDataResponseOutput).ToVolumeGroupMetaDataResponsePtrOutputWithContext(ctx)
}









type VolumeGroupMetaDataResponsePtrInput interface {
	pulumi.Input

	ToVolumeGroupMetaDataResponsePtrOutput() VolumeGroupMetaDataResponsePtrOutput
	ToVolumeGroupMetaDataResponsePtrOutputWithContext(context.Context) VolumeGroupMetaDataResponsePtrOutput
}

type volumeGroupMetaDataResponsePtrType VolumeGroupMetaDataResponseArgs

func VolumeGroupMetaDataResponsePtr(v *VolumeGroupMetaDataResponseArgs) VolumeGroupMetaDataResponsePtrInput {
	return (*volumeGroupMetaDataResponsePtrType)(v)
}

func (*volumeGroupMetaDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeGroupMetaDataResponse)(nil)).Elem()
}

func (i *volumeGroupMetaDataResponsePtrType) ToVolumeGroupMetaDataResponsePtrOutput() VolumeGroupMetaDataResponsePtrOutput {
	return i.ToVolumeGroupMetaDataResponsePtrOutputWithContext(context.Background())
}

func (i *volumeGroupMetaDataResponsePtrType) ToVolumeGroupMetaDataResponsePtrOutputWithContext(ctx context.Context) VolumeGroupMetaDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupMetaDataResponsePtrOutput)
}

type VolumeGroupMetaDataResponseOutput struct{ *pulumi.OutputState }

func (VolumeGroupMetaDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeGroupMetaDataResponse)(nil)).Elem()
}

func (o VolumeGroupMetaDataResponseOutput) ToVolumeGroupMetaDataResponseOutput() VolumeGroupMetaDataResponseOutput {
	return o
}

func (o VolumeGroupMetaDataResponseOutput) ToVolumeGroupMetaDataResponseOutputWithContext(ctx context.Context) VolumeGroupMetaDataResponseOutput {
	return o
}

func (o VolumeGroupMetaDataResponseOutput) ToVolumeGroupMetaDataResponsePtrOutput() VolumeGroupMetaDataResponsePtrOutput {
	return o.ToVolumeGroupMetaDataResponsePtrOutputWithContext(context.Background())
}

func (o VolumeGroupMetaDataResponseOutput) ToVolumeGroupMetaDataResponsePtrOutputWithContext(ctx context.Context) VolumeGroupMetaDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeGroupMetaDataResponse) *VolumeGroupMetaDataResponse {
		return &v
	}).(VolumeGroupMetaDataResponsePtrOutput)
}

func (o VolumeGroupMetaDataResponseOutput) ApplicationIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupMetaDataResponse) *string { return v.ApplicationIdentifier }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataResponseOutput) ApplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupMetaDataResponse) *string { return v.ApplicationType }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataResponseOutput) DeploymentSpecId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupMetaDataResponse) *string { return v.DeploymentSpecId }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataResponseOutput) GlobalPlacementRules() PlacementKeyValuePairsResponseArrayOutput {
	return o.ApplyT(func(v VolumeGroupMetaDataResponse) []PlacementKeyValuePairsResponse { return v.GlobalPlacementRules }).(PlacementKeyValuePairsResponseArrayOutput)
}

func (o VolumeGroupMetaDataResponseOutput) GroupDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupMetaDataResponse) *string { return v.GroupDescription }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataResponseOutput) VolumesCount() pulumi.Float64Output {
	return o.ApplyT(func(v VolumeGroupMetaDataResponse) float64 { return v.VolumesCount }).(pulumi.Float64Output)
}

type VolumeGroupMetaDataResponsePtrOutput struct{ *pulumi.OutputState }

func (VolumeGroupMetaDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeGroupMetaDataResponse)(nil)).Elem()
}

func (o VolumeGroupMetaDataResponsePtrOutput) ToVolumeGroupMetaDataResponsePtrOutput() VolumeGroupMetaDataResponsePtrOutput {
	return o
}

func (o VolumeGroupMetaDataResponsePtrOutput) ToVolumeGroupMetaDataResponsePtrOutputWithContext(ctx context.Context) VolumeGroupMetaDataResponsePtrOutput {
	return o
}

func (o VolumeGroupMetaDataResponsePtrOutput) Elem() VolumeGroupMetaDataResponseOutput {
	return o.ApplyT(func(v *VolumeGroupMetaDataResponse) VolumeGroupMetaDataResponse {
		if v != nil {
			return *v
		}
		var ret VolumeGroupMetaDataResponse
		return ret
	}).(VolumeGroupMetaDataResponseOutput)
}

func (o VolumeGroupMetaDataResponsePtrOutput) ApplicationIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroupMetaDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataResponsePtrOutput) ApplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroupMetaDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationType
	}).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataResponsePtrOutput) DeploymentSpecId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroupMetaDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentSpecId
	}).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataResponsePtrOutput) GlobalPlacementRules() PlacementKeyValuePairsResponseArrayOutput {
	return o.ApplyT(func(v *VolumeGroupMetaDataResponse) []PlacementKeyValuePairsResponse {
		if v == nil {
			return nil
		}
		return v.GlobalPlacementRules
	}).(PlacementKeyValuePairsResponseArrayOutput)
}

func (o VolumeGroupMetaDataResponsePtrOutput) GroupDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroupMetaDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupDescription
	}).(pulumi.StringPtrOutput)
}

func (o VolumeGroupMetaDataResponsePtrOutput) VolumesCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VolumeGroupMetaDataResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.VolumesCount
	}).(pulumi.Float64PtrOutput)
}

type VolumeGroupVolumeProperties struct {
	AvsDataStore             *string                         `pulumi:"avsDataStore"`
	BackupId                 *string                         `pulumi:"backupId"`
	CapacityPoolResourceId   *string                         `pulumi:"capacityPoolResourceId"`
	CoolAccess               *bool                           `pulumi:"coolAccess"`
	CoolnessPeriod           *int                            `pulumi:"coolnessPeriod"`
	CreationToken            string                          `pulumi:"creationToken"`
	DataProtection           *VolumePropertiesDataProtection `pulumi:"dataProtection"`
	DefaultGroupQuotaInKiBs  *float64                        `pulumi:"defaultGroupQuotaInKiBs"`
	DefaultUserQuotaInKiBs   *float64                        `pulumi:"defaultUserQuotaInKiBs"`
	EncryptionKeySource      *string                         `pulumi:"encryptionKeySource"`
	ExportPolicy             *VolumePropertiesExportPolicy   `pulumi:"exportPolicy"`
	IsDefaultQuotaEnabled    *bool                           `pulumi:"isDefaultQuotaEnabled"`
	IsRestoring              *bool                           `pulumi:"isRestoring"`
	KerberosEnabled          *bool                           `pulumi:"kerberosEnabled"`
	LdapEnabled              *bool                           `pulumi:"ldapEnabled"`
	Name                     *string                         `pulumi:"name"`
	NetworkFeatures          *string                         `pulumi:"networkFeatures"`
	PlacementRules           []PlacementKeyValuePairs        `pulumi:"placementRules"`
	ProtocolTypes            []string                        `pulumi:"protocolTypes"`
	ProximityPlacementGroup  *string                         `pulumi:"proximityPlacementGroup"`
	SecurityStyle            *string                         `pulumi:"securityStyle"`
	ServiceLevel             *string                         `pulumi:"serviceLevel"`
	SmbContinuouslyAvailable *bool                           `pulumi:"smbContinuouslyAvailable"`
	SmbEncryption            *bool                           `pulumi:"smbEncryption"`
	SnapshotDirectoryVisible *bool                           `pulumi:"snapshotDirectoryVisible"`
	SnapshotId               *string                         `pulumi:"snapshotId"`
	SubnetId                 string                          `pulumi:"subnetId"`
	Tags                     map[string]string               `pulumi:"tags"`
	ThroughputMibps          *float64                        `pulumi:"throughputMibps"`
	UnixPermissions          *string                         `pulumi:"unixPermissions"`
	UsageThreshold           float64                         `pulumi:"usageThreshold"`
	VolumeSpecName           *string                         `pulumi:"volumeSpecName"`
	VolumeType               *string                         `pulumi:"volumeType"`
}


func (val *VolumeGroupVolumeProperties) Defaults() *VolumeGroupVolumeProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AvsDataStore) {
		avsDataStore_ := "Disabled"
		tmp.AvsDataStore = &avsDataStore_
	}
	if isZero(tmp.CoolAccess) {
		coolAccess_ := false
		tmp.CoolAccess = &coolAccess_
	}
	if isZero(tmp.DefaultGroupQuotaInKiBs) {
		defaultGroupQuotaInKiBs_ := 0.0
		tmp.DefaultGroupQuotaInKiBs = &defaultGroupQuotaInKiBs_
	}
	if isZero(tmp.DefaultUserQuotaInKiBs) {
		defaultUserQuotaInKiBs_ := 0.0
		tmp.DefaultUserQuotaInKiBs = &defaultUserQuotaInKiBs_
	}
	if isZero(tmp.IsDefaultQuotaEnabled) {
		isDefaultQuotaEnabled_ := false
		tmp.IsDefaultQuotaEnabled = &isDefaultQuotaEnabled_
	}
	if isZero(tmp.KerberosEnabled) {
		kerberosEnabled_ := false
		tmp.KerberosEnabled = &kerberosEnabled_
	}
	if isZero(tmp.LdapEnabled) {
		ldapEnabled_ := false
		tmp.LdapEnabled = &ldapEnabled_
	}
	if isZero(tmp.NetworkFeatures) {
		networkFeatures_ := "Basic"
		tmp.NetworkFeatures = &networkFeatures_
	}
	if isZero(tmp.SecurityStyle) {
		securityStyle_ := "unix"
		tmp.SecurityStyle = &securityStyle_
	}
	if isZero(tmp.SmbContinuouslyAvailable) {
		smbContinuouslyAvailable_ := false
		tmp.SmbContinuouslyAvailable = &smbContinuouslyAvailable_
	}
	if isZero(tmp.SmbEncryption) {
		smbEncryption_ := false
		tmp.SmbEncryption = &smbEncryption_
	}
	if isZero(tmp.SnapshotDirectoryVisible) {
		snapshotDirectoryVisible_ := true
		tmp.SnapshotDirectoryVisible = &snapshotDirectoryVisible_
	}
	if isZero(tmp.ThroughputMibps) {
		throughputMibps_ := 0.0
		tmp.ThroughputMibps = &throughputMibps_
	}
	if isZero(tmp.UnixPermissions) {
		unixPermissions_ := "0770"
		tmp.UnixPermissions = &unixPermissions_
	}
	if isZero(tmp.UsageThreshold) {
		tmp.UsageThreshold = 107374182400.0
	}
	return &tmp
}





type VolumeGroupVolumePropertiesInput interface {
	pulumi.Input

	ToVolumeGroupVolumePropertiesOutput() VolumeGroupVolumePropertiesOutput
	ToVolumeGroupVolumePropertiesOutputWithContext(context.Context) VolumeGroupVolumePropertiesOutput
}

type VolumeGroupVolumePropertiesArgs struct {
	AvsDataStore             pulumi.StringPtrInput                  `pulumi:"avsDataStore"`
	BackupId                 pulumi.StringPtrInput                  `pulumi:"backupId"`
	CapacityPoolResourceId   pulumi.StringPtrInput                  `pulumi:"capacityPoolResourceId"`
	CoolAccess               pulumi.BoolPtrInput                    `pulumi:"coolAccess"`
	CoolnessPeriod           pulumi.IntPtrInput                     `pulumi:"coolnessPeriod"`
	CreationToken            pulumi.StringInput                     `pulumi:"creationToken"`
	DataProtection           VolumePropertiesDataProtectionPtrInput `pulumi:"dataProtection"`
	DefaultGroupQuotaInKiBs  pulumi.Float64PtrInput                 `pulumi:"defaultGroupQuotaInKiBs"`
	DefaultUserQuotaInKiBs   pulumi.Float64PtrInput                 `pulumi:"defaultUserQuotaInKiBs"`
	EncryptionKeySource      pulumi.StringPtrInput                  `pulumi:"encryptionKeySource"`
	ExportPolicy             VolumePropertiesExportPolicyPtrInput   `pulumi:"exportPolicy"`
	IsDefaultQuotaEnabled    pulumi.BoolPtrInput                    `pulumi:"isDefaultQuotaEnabled"`
	IsRestoring              pulumi.BoolPtrInput                    `pulumi:"isRestoring"`
	KerberosEnabled          pulumi.BoolPtrInput                    `pulumi:"kerberosEnabled"`
	LdapEnabled              pulumi.BoolPtrInput                    `pulumi:"ldapEnabled"`
	Name                     pulumi.StringPtrInput                  `pulumi:"name"`
	NetworkFeatures          pulumi.StringPtrInput                  `pulumi:"networkFeatures"`
	PlacementRules           PlacementKeyValuePairsArrayInput       `pulumi:"placementRules"`
	ProtocolTypes            pulumi.StringArrayInput                `pulumi:"protocolTypes"`
	ProximityPlacementGroup  pulumi.StringPtrInput                  `pulumi:"proximityPlacementGroup"`
	SecurityStyle            pulumi.StringPtrInput                  `pulumi:"securityStyle"`
	ServiceLevel             pulumi.StringPtrInput                  `pulumi:"serviceLevel"`
	SmbContinuouslyAvailable pulumi.BoolPtrInput                    `pulumi:"smbContinuouslyAvailable"`
	SmbEncryption            pulumi.BoolPtrInput                    `pulumi:"smbEncryption"`
	SnapshotDirectoryVisible pulumi.BoolPtrInput                    `pulumi:"snapshotDirectoryVisible"`
	SnapshotId               pulumi.StringPtrInput                  `pulumi:"snapshotId"`
	SubnetId                 pulumi.StringInput                     `pulumi:"subnetId"`
	Tags                     pulumi.StringMapInput                  `pulumi:"tags"`
	ThroughputMibps          pulumi.Float64PtrInput                 `pulumi:"throughputMibps"`
	UnixPermissions          pulumi.StringPtrInput                  `pulumi:"unixPermissions"`
	UsageThreshold           pulumi.Float64Input                    `pulumi:"usageThreshold"`
	VolumeSpecName           pulumi.StringPtrInput                  `pulumi:"volumeSpecName"`
	VolumeType               pulumi.StringPtrInput                  `pulumi:"volumeType"`
}

func (VolumeGroupVolumePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeGroupVolumeProperties)(nil)).Elem()
}

func (i VolumeGroupVolumePropertiesArgs) ToVolumeGroupVolumePropertiesOutput() VolumeGroupVolumePropertiesOutput {
	return i.ToVolumeGroupVolumePropertiesOutputWithContext(context.Background())
}

func (i VolumeGroupVolumePropertiesArgs) ToVolumeGroupVolumePropertiesOutputWithContext(ctx context.Context) VolumeGroupVolumePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupVolumePropertiesOutput)
}





type VolumeGroupVolumePropertiesArrayInput interface {
	pulumi.Input

	ToVolumeGroupVolumePropertiesArrayOutput() VolumeGroupVolumePropertiesArrayOutput
	ToVolumeGroupVolumePropertiesArrayOutputWithContext(context.Context) VolumeGroupVolumePropertiesArrayOutput
}

type VolumeGroupVolumePropertiesArray []VolumeGroupVolumePropertiesInput

func (VolumeGroupVolumePropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeGroupVolumeProperties)(nil)).Elem()
}

func (i VolumeGroupVolumePropertiesArray) ToVolumeGroupVolumePropertiesArrayOutput() VolumeGroupVolumePropertiesArrayOutput {
	return i.ToVolumeGroupVolumePropertiesArrayOutputWithContext(context.Background())
}

func (i VolumeGroupVolumePropertiesArray) ToVolumeGroupVolumePropertiesArrayOutputWithContext(ctx context.Context) VolumeGroupVolumePropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupVolumePropertiesArrayOutput)
}

type VolumeGroupVolumePropertiesOutput struct{ *pulumi.OutputState }

func (VolumeGroupVolumePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeGroupVolumeProperties)(nil)).Elem()
}

func (o VolumeGroupVolumePropertiesOutput) ToVolumeGroupVolumePropertiesOutput() VolumeGroupVolumePropertiesOutput {
	return o
}

func (o VolumeGroupVolumePropertiesOutput) ToVolumeGroupVolumePropertiesOutputWithContext(ctx context.Context) VolumeGroupVolumePropertiesOutput {
	return o
}

func (o VolumeGroupVolumePropertiesOutput) AvsDataStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.AvsDataStore }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.BackupId }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) CapacityPoolResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.CapacityPoolResourceId }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) CoolAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *bool { return v.CoolAccess }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) CoolnessPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *int { return v.CoolnessPeriod }).(pulumi.IntPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) string { return v.CreationToken }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesOutput) DataProtection() VolumePropertiesDataProtectionPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *VolumePropertiesDataProtection { return v.DataProtection }).(VolumePropertiesDataProtectionPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) DefaultGroupQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *float64 { return v.DefaultGroupQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) DefaultUserQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *float64 { return v.DefaultUserQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) EncryptionKeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.EncryptionKeySource }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) ExportPolicy() VolumePropertiesExportPolicyPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *VolumePropertiesExportPolicy { return v.ExportPolicy }).(VolumePropertiesExportPolicyPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) IsDefaultQuotaEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *bool { return v.IsDefaultQuotaEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) IsRestoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *bool { return v.IsRestoring }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) KerberosEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *bool { return v.KerberosEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) LdapEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *bool { return v.LdapEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) NetworkFeatures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.NetworkFeatures }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) PlacementRules() PlacementKeyValuePairsArrayOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) []PlacementKeyValuePairs { return v.PlacementRules }).(PlacementKeyValuePairsArrayOutput)
}

func (o VolumeGroupVolumePropertiesOutput) ProtocolTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) []string { return v.ProtocolTypes }).(pulumi.StringArrayOutput)
}

func (o VolumeGroupVolumePropertiesOutput) ProximityPlacementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.ProximityPlacementGroup }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) SecurityStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.SecurityStyle }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) SmbContinuouslyAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *bool { return v.SmbContinuouslyAvailable }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) SmbEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *bool { return v.SmbEncryption }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) SnapshotDirectoryVisible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *bool { return v.SnapshotDirectoryVisible }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VolumeGroupVolumePropertiesOutput) ThroughputMibps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *float64 { return v.ThroughputMibps }).(pulumi.Float64PtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) UnixPermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.UnixPermissions }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) UsageThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) float64 { return v.UsageThreshold }).(pulumi.Float64Output)
}

func (o VolumeGroupVolumePropertiesOutput) VolumeSpecName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.VolumeSpecName }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumeProperties) *string { return v.VolumeType }).(pulumi.StringPtrOutput)
}

type VolumeGroupVolumePropertiesArrayOutput struct{ *pulumi.OutputState }

func (VolumeGroupVolumePropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeGroupVolumeProperties)(nil)).Elem()
}

func (o VolumeGroupVolumePropertiesArrayOutput) ToVolumeGroupVolumePropertiesArrayOutput() VolumeGroupVolumePropertiesArrayOutput {
	return o
}

func (o VolumeGroupVolumePropertiesArrayOutput) ToVolumeGroupVolumePropertiesArrayOutputWithContext(ctx context.Context) VolumeGroupVolumePropertiesArrayOutput {
	return o
}

func (o VolumeGroupVolumePropertiesArrayOutput) Index(i pulumi.IntInput) VolumeGroupVolumePropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeGroupVolumeProperties {
		return vs[0].([]VolumeGroupVolumeProperties)[vs[1].(int)]
	}).(VolumeGroupVolumePropertiesOutput)
}

type VolumeGroupVolumePropertiesResponse struct {
	AvsDataStore              *string                                 `pulumi:"avsDataStore"`
	BackupId                  *string                                 `pulumi:"backupId"`
	BaremetalTenantId         string                                  `pulumi:"baremetalTenantId"`
	CapacityPoolResourceId    *string                                 `pulumi:"capacityPoolResourceId"`
	CloneProgress             int                                     `pulumi:"cloneProgress"`
	CoolAccess                *bool                                   `pulumi:"coolAccess"`
	CoolnessPeriod            *int                                    `pulumi:"coolnessPeriod"`
	CreationToken             string                                  `pulumi:"creationToken"`
	DataProtection            *VolumePropertiesResponseDataProtection `pulumi:"dataProtection"`
	DefaultGroupQuotaInKiBs   *float64                                `pulumi:"defaultGroupQuotaInKiBs"`
	DefaultUserQuotaInKiBs    *float64                                `pulumi:"defaultUserQuotaInKiBs"`
	EncryptionKeySource       *string                                 `pulumi:"encryptionKeySource"`
	ExportPolicy              *VolumePropertiesResponseExportPolicy   `pulumi:"exportPolicy"`
	FileSystemId              string                                  `pulumi:"fileSystemId"`
	Id                        string                                  `pulumi:"id"`
	IsDefaultQuotaEnabled     *bool                                   `pulumi:"isDefaultQuotaEnabled"`
	IsRestoring               *bool                                   `pulumi:"isRestoring"`
	KerberosEnabled           *bool                                   `pulumi:"kerberosEnabled"`
	LdapEnabled               *bool                                   `pulumi:"ldapEnabled"`
	MountTargets              []MountTargetPropertiesResponse         `pulumi:"mountTargets"`
	Name                      *string                                 `pulumi:"name"`
	NetworkFeatures           *string                                 `pulumi:"networkFeatures"`
	NetworkSiblingSetId       string                                  `pulumi:"networkSiblingSetId"`
	PlacementRules            []PlacementKeyValuePairsResponse        `pulumi:"placementRules"`
	ProtocolTypes             []string                                `pulumi:"protocolTypes"`
	ProvisioningState         string                                  `pulumi:"provisioningState"`
	ProximityPlacementGroup   *string                                 `pulumi:"proximityPlacementGroup"`
	SecurityStyle             *string                                 `pulumi:"securityStyle"`
	ServiceLevel              *string                                 `pulumi:"serviceLevel"`
	SmbContinuouslyAvailable  *bool                                   `pulumi:"smbContinuouslyAvailable"`
	SmbEncryption             *bool                                   `pulumi:"smbEncryption"`
	SnapshotDirectoryVisible  *bool                                   `pulumi:"snapshotDirectoryVisible"`
	SnapshotId                *string                                 `pulumi:"snapshotId"`
	StorageToNetworkProximity string                                  `pulumi:"storageToNetworkProximity"`
	SubnetId                  string                                  `pulumi:"subnetId"`
	T2Network                 string                                  `pulumi:"t2Network"`
	Tags                      map[string]string                       `pulumi:"tags"`
	ThroughputMibps           *float64                                `pulumi:"throughputMibps"`
	Type                      string                                  `pulumi:"type"`
	UnixPermissions           *string                                 `pulumi:"unixPermissions"`
	UsageThreshold            float64                                 `pulumi:"usageThreshold"`
	VolumeGroupName           string                                  `pulumi:"volumeGroupName"`
	VolumeSpecName            *string                                 `pulumi:"volumeSpecName"`
	VolumeType                *string                                 `pulumi:"volumeType"`
}


func (val *VolumeGroupVolumePropertiesResponse) Defaults() *VolumeGroupVolumePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AvsDataStore) {
		avsDataStore_ := "Disabled"
		tmp.AvsDataStore = &avsDataStore_
	}
	if isZero(tmp.CoolAccess) {
		coolAccess_ := false
		tmp.CoolAccess = &coolAccess_
	}
	if isZero(tmp.DefaultGroupQuotaInKiBs) {
		defaultGroupQuotaInKiBs_ := 0.0
		tmp.DefaultGroupQuotaInKiBs = &defaultGroupQuotaInKiBs_
	}
	if isZero(tmp.DefaultUserQuotaInKiBs) {
		defaultUserQuotaInKiBs_ := 0.0
		tmp.DefaultUserQuotaInKiBs = &defaultUserQuotaInKiBs_
	}
	if isZero(tmp.IsDefaultQuotaEnabled) {
		isDefaultQuotaEnabled_ := false
		tmp.IsDefaultQuotaEnabled = &isDefaultQuotaEnabled_
	}
	if isZero(tmp.KerberosEnabled) {
		kerberosEnabled_ := false
		tmp.KerberosEnabled = &kerberosEnabled_
	}
	if isZero(tmp.LdapEnabled) {
		ldapEnabled_ := false
		tmp.LdapEnabled = &ldapEnabled_
	}
	if isZero(tmp.NetworkFeatures) {
		networkFeatures_ := "Basic"
		tmp.NetworkFeatures = &networkFeatures_
	}
	if isZero(tmp.SecurityStyle) {
		securityStyle_ := "unix"
		tmp.SecurityStyle = &securityStyle_
	}
	if isZero(tmp.SmbContinuouslyAvailable) {
		smbContinuouslyAvailable_ := false
		tmp.SmbContinuouslyAvailable = &smbContinuouslyAvailable_
	}
	if isZero(tmp.SmbEncryption) {
		smbEncryption_ := false
		tmp.SmbEncryption = &smbEncryption_
	}
	if isZero(tmp.SnapshotDirectoryVisible) {
		snapshotDirectoryVisible_ := true
		tmp.SnapshotDirectoryVisible = &snapshotDirectoryVisible_
	}
	if isZero(tmp.ThroughputMibps) {
		throughputMibps_ := 0.0
		tmp.ThroughputMibps = &throughputMibps_
	}
	if isZero(tmp.UnixPermissions) {
		unixPermissions_ := "0770"
		tmp.UnixPermissions = &unixPermissions_
	}
	if isZero(tmp.UsageThreshold) {
		tmp.UsageThreshold = 107374182400.0
	}
	return &tmp
}





type VolumeGroupVolumePropertiesResponseInput interface {
	pulumi.Input

	ToVolumeGroupVolumePropertiesResponseOutput() VolumeGroupVolumePropertiesResponseOutput
	ToVolumeGroupVolumePropertiesResponseOutputWithContext(context.Context) VolumeGroupVolumePropertiesResponseOutput
}

type VolumeGroupVolumePropertiesResponseArgs struct {
	AvsDataStore              pulumi.StringPtrInput                          `pulumi:"avsDataStore"`
	BackupId                  pulumi.StringPtrInput                          `pulumi:"backupId"`
	BaremetalTenantId         pulumi.StringInput                             `pulumi:"baremetalTenantId"`
	CapacityPoolResourceId    pulumi.StringPtrInput                          `pulumi:"capacityPoolResourceId"`
	CloneProgress             pulumi.IntInput                                `pulumi:"cloneProgress"`
	CoolAccess                pulumi.BoolPtrInput                            `pulumi:"coolAccess"`
	CoolnessPeriod            pulumi.IntPtrInput                             `pulumi:"coolnessPeriod"`
	CreationToken             pulumi.StringInput                             `pulumi:"creationToken"`
	DataProtection            VolumePropertiesResponseDataProtectionPtrInput `pulumi:"dataProtection"`
	DefaultGroupQuotaInKiBs   pulumi.Float64PtrInput                         `pulumi:"defaultGroupQuotaInKiBs"`
	DefaultUserQuotaInKiBs    pulumi.Float64PtrInput                         `pulumi:"defaultUserQuotaInKiBs"`
	EncryptionKeySource       pulumi.StringPtrInput                          `pulumi:"encryptionKeySource"`
	ExportPolicy              VolumePropertiesResponseExportPolicyPtrInput   `pulumi:"exportPolicy"`
	FileSystemId              pulumi.StringInput                             `pulumi:"fileSystemId"`
	Id                        pulumi.StringInput                             `pulumi:"id"`
	IsDefaultQuotaEnabled     pulumi.BoolPtrInput                            `pulumi:"isDefaultQuotaEnabled"`
	IsRestoring               pulumi.BoolPtrInput                            `pulumi:"isRestoring"`
	KerberosEnabled           pulumi.BoolPtrInput                            `pulumi:"kerberosEnabled"`
	LdapEnabled               pulumi.BoolPtrInput                            `pulumi:"ldapEnabled"`
	MountTargets              MountTargetPropertiesResponseArrayInput        `pulumi:"mountTargets"`
	Name                      pulumi.StringPtrInput                          `pulumi:"name"`
	NetworkFeatures           pulumi.StringPtrInput                          `pulumi:"networkFeatures"`
	NetworkSiblingSetId       pulumi.StringInput                             `pulumi:"networkSiblingSetId"`
	PlacementRules            PlacementKeyValuePairsResponseArrayInput       `pulumi:"placementRules"`
	ProtocolTypes             pulumi.StringArrayInput                        `pulumi:"protocolTypes"`
	ProvisioningState         pulumi.StringInput                             `pulumi:"provisioningState"`
	ProximityPlacementGroup   pulumi.StringPtrInput                          `pulumi:"proximityPlacementGroup"`
	SecurityStyle             pulumi.StringPtrInput                          `pulumi:"securityStyle"`
	ServiceLevel              pulumi.StringPtrInput                          `pulumi:"serviceLevel"`
	SmbContinuouslyAvailable  pulumi.BoolPtrInput                            `pulumi:"smbContinuouslyAvailable"`
	SmbEncryption             pulumi.BoolPtrInput                            `pulumi:"smbEncryption"`
	SnapshotDirectoryVisible  pulumi.BoolPtrInput                            `pulumi:"snapshotDirectoryVisible"`
	SnapshotId                pulumi.StringPtrInput                          `pulumi:"snapshotId"`
	StorageToNetworkProximity pulumi.StringInput                             `pulumi:"storageToNetworkProximity"`
	SubnetId                  pulumi.StringInput                             `pulumi:"subnetId"`
	T2Network                 pulumi.StringInput                             `pulumi:"t2Network"`
	Tags                      pulumi.StringMapInput                          `pulumi:"tags"`
	ThroughputMibps           pulumi.Float64PtrInput                         `pulumi:"throughputMibps"`
	Type                      pulumi.StringInput                             `pulumi:"type"`
	UnixPermissions           pulumi.StringPtrInput                          `pulumi:"unixPermissions"`
	UsageThreshold            pulumi.Float64Input                            `pulumi:"usageThreshold"`
	VolumeGroupName           pulumi.StringInput                             `pulumi:"volumeGroupName"`
	VolumeSpecName            pulumi.StringPtrInput                          `pulumi:"volumeSpecName"`
	VolumeType                pulumi.StringPtrInput                          `pulumi:"volumeType"`
}

func (VolumeGroupVolumePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeGroupVolumePropertiesResponse)(nil)).Elem()
}

func (i VolumeGroupVolumePropertiesResponseArgs) ToVolumeGroupVolumePropertiesResponseOutput() VolumeGroupVolumePropertiesResponseOutput {
	return i.ToVolumeGroupVolumePropertiesResponseOutputWithContext(context.Background())
}

func (i VolumeGroupVolumePropertiesResponseArgs) ToVolumeGroupVolumePropertiesResponseOutputWithContext(ctx context.Context) VolumeGroupVolumePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupVolumePropertiesResponseOutput)
}





type VolumeGroupVolumePropertiesResponseArrayInput interface {
	pulumi.Input

	ToVolumeGroupVolumePropertiesResponseArrayOutput() VolumeGroupVolumePropertiesResponseArrayOutput
	ToVolumeGroupVolumePropertiesResponseArrayOutputWithContext(context.Context) VolumeGroupVolumePropertiesResponseArrayOutput
}

type VolumeGroupVolumePropertiesResponseArray []VolumeGroupVolumePropertiesResponseInput

func (VolumeGroupVolumePropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeGroupVolumePropertiesResponse)(nil)).Elem()
}

func (i VolumeGroupVolumePropertiesResponseArray) ToVolumeGroupVolumePropertiesResponseArrayOutput() VolumeGroupVolumePropertiesResponseArrayOutput {
	return i.ToVolumeGroupVolumePropertiesResponseArrayOutputWithContext(context.Background())
}

func (i VolumeGroupVolumePropertiesResponseArray) ToVolumeGroupVolumePropertiesResponseArrayOutputWithContext(ctx context.Context) VolumeGroupVolumePropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupVolumePropertiesResponseArrayOutput)
}

type VolumeGroupVolumePropertiesResponseOutput struct{ *pulumi.OutputState }

func (VolumeGroupVolumePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeGroupVolumePropertiesResponse)(nil)).Elem()
}

func (o VolumeGroupVolumePropertiesResponseOutput) ToVolumeGroupVolumePropertiesResponseOutput() VolumeGroupVolumePropertiesResponseOutput {
	return o
}

func (o VolumeGroupVolumePropertiesResponseOutput) ToVolumeGroupVolumePropertiesResponseOutputWithContext(ctx context.Context) VolumeGroupVolumePropertiesResponseOutput {
	return o
}

func (o VolumeGroupVolumePropertiesResponseOutput) AvsDataStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.AvsDataStore }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.BackupId }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) BaremetalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.BaremetalTenantId }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) CapacityPoolResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.CapacityPoolResourceId }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) CloneProgress() pulumi.IntOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) int { return v.CloneProgress }).(pulumi.IntOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) CoolAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *bool { return v.CoolAccess }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) CoolnessPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *int { return v.CoolnessPeriod }).(pulumi.IntPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.CreationToken }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) DataProtection() VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *VolumePropertiesResponseDataProtection {
		return v.DataProtection
	}).(VolumePropertiesResponseDataProtectionPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) DefaultGroupQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *float64 { return v.DefaultGroupQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) DefaultUserQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *float64 { return v.DefaultUserQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) EncryptionKeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.EncryptionKeySource }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) ExportPolicy() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *VolumePropertiesResponseExportPolicy {
		return v.ExportPolicy
	}).(VolumePropertiesResponseExportPolicyPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) IsDefaultQuotaEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *bool { return v.IsDefaultQuotaEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) IsRestoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *bool { return v.IsRestoring }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) KerberosEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *bool { return v.KerberosEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) LdapEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *bool { return v.LdapEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) MountTargets() MountTargetPropertiesResponseArrayOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) []MountTargetPropertiesResponse { return v.MountTargets }).(MountTargetPropertiesResponseArrayOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) NetworkFeatures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.NetworkFeatures }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) NetworkSiblingSetId() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.NetworkSiblingSetId }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) PlacementRules() PlacementKeyValuePairsResponseArrayOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) []PlacementKeyValuePairsResponse { return v.PlacementRules }).(PlacementKeyValuePairsResponseArrayOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) ProtocolTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) []string { return v.ProtocolTypes }).(pulumi.StringArrayOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) ProximityPlacementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.ProximityPlacementGroup }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) SecurityStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.SecurityStyle }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) SmbContinuouslyAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *bool { return v.SmbContinuouslyAvailable }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) SmbEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *bool { return v.SmbEncryption }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) SnapshotDirectoryVisible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *bool { return v.SnapshotDirectoryVisible }).(pulumi.BoolPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) StorageToNetworkProximity() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.StorageToNetworkProximity }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) T2Network() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.T2Network }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) ThroughputMibps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *float64 { return v.ThroughputMibps }).(pulumi.Float64PtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) UnixPermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.UnixPermissions }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) UsageThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) float64 { return v.UsageThreshold }).(pulumi.Float64Output)
}

func (o VolumeGroupVolumePropertiesResponseOutput) VolumeGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) string { return v.VolumeGroupName }).(pulumi.StringOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) VolumeSpecName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.VolumeSpecName }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupVolumePropertiesResponseOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeGroupVolumePropertiesResponse) *string { return v.VolumeType }).(pulumi.StringPtrOutput)
}

type VolumeGroupVolumePropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (VolumeGroupVolumePropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeGroupVolumePropertiesResponse)(nil)).Elem()
}

func (o VolumeGroupVolumePropertiesResponseArrayOutput) ToVolumeGroupVolumePropertiesResponseArrayOutput() VolumeGroupVolumePropertiesResponseArrayOutput {
	return o
}

func (o VolumeGroupVolumePropertiesResponseArrayOutput) ToVolumeGroupVolumePropertiesResponseArrayOutputWithContext(ctx context.Context) VolumeGroupVolumePropertiesResponseArrayOutput {
	return o
}

func (o VolumeGroupVolumePropertiesResponseArrayOutput) Index(i pulumi.IntInput) VolumeGroupVolumePropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeGroupVolumePropertiesResponse {
		return vs[0].([]VolumeGroupVolumePropertiesResponse)[vs[1].(int)]
	}).(VolumeGroupVolumePropertiesResponseOutput)
}

type VolumePropertiesDataProtection struct {
	Backup      *VolumeBackupProperties   `pulumi:"backup"`
	Replication *ReplicationObject        `pulumi:"replication"`
	Snapshot    *VolumeSnapshotProperties `pulumi:"snapshot"`
}





type VolumePropertiesDataProtectionInput interface {
	pulumi.Input

	ToVolumePropertiesDataProtectionOutput() VolumePropertiesDataProtectionOutput
	ToVolumePropertiesDataProtectionOutputWithContext(context.Context) VolumePropertiesDataProtectionOutput
}

type VolumePropertiesDataProtectionArgs struct {
	Backup      VolumeBackupPropertiesPtrInput   `pulumi:"backup"`
	Replication ReplicationObjectPtrInput        `pulumi:"replication"`
	Snapshot    VolumeSnapshotPropertiesPtrInput `pulumi:"snapshot"`
}

func (VolumePropertiesDataProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesDataProtection)(nil)).Elem()
}

func (i VolumePropertiesDataProtectionArgs) ToVolumePropertiesDataProtectionOutput() VolumePropertiesDataProtectionOutput {
	return i.ToVolumePropertiesDataProtectionOutputWithContext(context.Background())
}

func (i VolumePropertiesDataProtectionArgs) ToVolumePropertiesDataProtectionOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesDataProtectionOutput)
}

func (i VolumePropertiesDataProtectionArgs) ToVolumePropertiesDataProtectionPtrOutput() VolumePropertiesDataProtectionPtrOutput {
	return i.ToVolumePropertiesDataProtectionPtrOutputWithContext(context.Background())
}

func (i VolumePropertiesDataProtectionArgs) ToVolumePropertiesDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesDataProtectionOutput).ToVolumePropertiesDataProtectionPtrOutputWithContext(ctx)
}









type VolumePropertiesDataProtectionPtrInput interface {
	pulumi.Input

	ToVolumePropertiesDataProtectionPtrOutput() VolumePropertiesDataProtectionPtrOutput
	ToVolumePropertiesDataProtectionPtrOutputWithContext(context.Context) VolumePropertiesDataProtectionPtrOutput
}

type volumePropertiesDataProtectionPtrType VolumePropertiesDataProtectionArgs

func VolumePropertiesDataProtectionPtr(v *VolumePropertiesDataProtectionArgs) VolumePropertiesDataProtectionPtrInput {
	return (*volumePropertiesDataProtectionPtrType)(v)
}

func (*volumePropertiesDataProtectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesDataProtection)(nil)).Elem()
}

func (i *volumePropertiesDataProtectionPtrType) ToVolumePropertiesDataProtectionPtrOutput() VolumePropertiesDataProtectionPtrOutput {
	return i.ToVolumePropertiesDataProtectionPtrOutputWithContext(context.Background())
}

func (i *volumePropertiesDataProtectionPtrType) ToVolumePropertiesDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesDataProtectionPtrOutput)
}

type VolumePropertiesDataProtectionOutput struct{ *pulumi.OutputState }

func (VolumePropertiesDataProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesDataProtection)(nil)).Elem()
}

func (o VolumePropertiesDataProtectionOutput) ToVolumePropertiesDataProtectionOutput() VolumePropertiesDataProtectionOutput {
	return o
}

func (o VolumePropertiesDataProtectionOutput) ToVolumePropertiesDataProtectionOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionOutput {
	return o
}

func (o VolumePropertiesDataProtectionOutput) ToVolumePropertiesDataProtectionPtrOutput() VolumePropertiesDataProtectionPtrOutput {
	return o.ToVolumePropertiesDataProtectionPtrOutputWithContext(context.Background())
}

func (o VolumePropertiesDataProtectionOutput) ToVolumePropertiesDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumePropertiesDataProtection) *VolumePropertiesDataProtection {
		return &v
	}).(VolumePropertiesDataProtectionPtrOutput)
}

func (o VolumePropertiesDataProtectionOutput) Backup() VolumeBackupPropertiesPtrOutput {
	return o.ApplyT(func(v VolumePropertiesDataProtection) *VolumeBackupProperties { return v.Backup }).(VolumeBackupPropertiesPtrOutput)
}

func (o VolumePropertiesDataProtectionOutput) Replication() ReplicationObjectPtrOutput {
	return o.ApplyT(func(v VolumePropertiesDataProtection) *ReplicationObject { return v.Replication }).(ReplicationObjectPtrOutput)
}

func (o VolumePropertiesDataProtectionOutput) Snapshot() VolumeSnapshotPropertiesPtrOutput {
	return o.ApplyT(func(v VolumePropertiesDataProtection) *VolumeSnapshotProperties { return v.Snapshot }).(VolumeSnapshotPropertiesPtrOutput)
}

type VolumePropertiesDataProtectionPtrOutput struct{ *pulumi.OutputState }

func (VolumePropertiesDataProtectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesDataProtection)(nil)).Elem()
}

func (o VolumePropertiesDataProtectionPtrOutput) ToVolumePropertiesDataProtectionPtrOutput() VolumePropertiesDataProtectionPtrOutput {
	return o
}

func (o VolumePropertiesDataProtectionPtrOutput) ToVolumePropertiesDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionPtrOutput {
	return o
}

func (o VolumePropertiesDataProtectionPtrOutput) Elem() VolumePropertiesDataProtectionOutput {
	return o.ApplyT(func(v *VolumePropertiesDataProtection) VolumePropertiesDataProtection {
		if v != nil {
			return *v
		}
		var ret VolumePropertiesDataProtection
		return ret
	}).(VolumePropertiesDataProtectionOutput)
}

func (o VolumePropertiesDataProtectionPtrOutput) Backup() VolumeBackupPropertiesPtrOutput {
	return o.ApplyT(func(v *VolumePropertiesDataProtection) *VolumeBackupProperties {
		if v == nil {
			return nil
		}
		return v.Backup
	}).(VolumeBackupPropertiesPtrOutput)
}

func (o VolumePropertiesDataProtectionPtrOutput) Replication() ReplicationObjectPtrOutput {
	return o.ApplyT(func(v *VolumePropertiesDataProtection) *ReplicationObject {
		if v == nil {
			return nil
		}
		return v.Replication
	}).(ReplicationObjectPtrOutput)
}

func (o VolumePropertiesDataProtectionPtrOutput) Snapshot() VolumeSnapshotPropertiesPtrOutput {
	return o.ApplyT(func(v *VolumePropertiesDataProtection) *VolumeSnapshotProperties {
		if v == nil {
			return nil
		}
		return v.Snapshot
	}).(VolumeSnapshotPropertiesPtrOutput)
}

type VolumePropertiesExportPolicy struct {
	Rules []ExportPolicyRule `pulumi:"rules"`
}





type VolumePropertiesExportPolicyInput interface {
	pulumi.Input

	ToVolumePropertiesExportPolicyOutput() VolumePropertiesExportPolicyOutput
	ToVolumePropertiesExportPolicyOutputWithContext(context.Context) VolumePropertiesExportPolicyOutput
}

type VolumePropertiesExportPolicyArgs struct {
	Rules ExportPolicyRuleArrayInput `pulumi:"rules"`
}

func (VolumePropertiesExportPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesExportPolicy)(nil)).Elem()
}

func (i VolumePropertiesExportPolicyArgs) ToVolumePropertiesExportPolicyOutput() VolumePropertiesExportPolicyOutput {
	return i.ToVolumePropertiesExportPolicyOutputWithContext(context.Background())
}

func (i VolumePropertiesExportPolicyArgs) ToVolumePropertiesExportPolicyOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesExportPolicyOutput)
}

func (i VolumePropertiesExportPolicyArgs) ToVolumePropertiesExportPolicyPtrOutput() VolumePropertiesExportPolicyPtrOutput {
	return i.ToVolumePropertiesExportPolicyPtrOutputWithContext(context.Background())
}

func (i VolumePropertiesExportPolicyArgs) ToVolumePropertiesExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesExportPolicyOutput).ToVolumePropertiesExportPolicyPtrOutputWithContext(ctx)
}









type VolumePropertiesExportPolicyPtrInput interface {
	pulumi.Input

	ToVolumePropertiesExportPolicyPtrOutput() VolumePropertiesExportPolicyPtrOutput
	ToVolumePropertiesExportPolicyPtrOutputWithContext(context.Context) VolumePropertiesExportPolicyPtrOutput
}

type volumePropertiesExportPolicyPtrType VolumePropertiesExportPolicyArgs

func VolumePropertiesExportPolicyPtr(v *VolumePropertiesExportPolicyArgs) VolumePropertiesExportPolicyPtrInput {
	return (*volumePropertiesExportPolicyPtrType)(v)
}

func (*volumePropertiesExportPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesExportPolicy)(nil)).Elem()
}

func (i *volumePropertiesExportPolicyPtrType) ToVolumePropertiesExportPolicyPtrOutput() VolumePropertiesExportPolicyPtrOutput {
	return i.ToVolumePropertiesExportPolicyPtrOutputWithContext(context.Background())
}

func (i *volumePropertiesExportPolicyPtrType) ToVolumePropertiesExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesExportPolicyPtrOutput)
}

type VolumePropertiesExportPolicyOutput struct{ *pulumi.OutputState }

func (VolumePropertiesExportPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesExportPolicy)(nil)).Elem()
}

func (o VolumePropertiesExportPolicyOutput) ToVolumePropertiesExportPolicyOutput() VolumePropertiesExportPolicyOutput {
	return o
}

func (o VolumePropertiesExportPolicyOutput) ToVolumePropertiesExportPolicyOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyOutput {
	return o
}

func (o VolumePropertiesExportPolicyOutput) ToVolumePropertiesExportPolicyPtrOutput() VolumePropertiesExportPolicyPtrOutput {
	return o.ToVolumePropertiesExportPolicyPtrOutputWithContext(context.Background())
}

func (o VolumePropertiesExportPolicyOutput) ToVolumePropertiesExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumePropertiesExportPolicy) *VolumePropertiesExportPolicy {
		return &v
	}).(VolumePropertiesExportPolicyPtrOutput)
}

func (o VolumePropertiesExportPolicyOutput) Rules() ExportPolicyRuleArrayOutput {
	return o.ApplyT(func(v VolumePropertiesExportPolicy) []ExportPolicyRule { return v.Rules }).(ExportPolicyRuleArrayOutput)
}

type VolumePropertiesExportPolicyPtrOutput struct{ *pulumi.OutputState }

func (VolumePropertiesExportPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesExportPolicy)(nil)).Elem()
}

func (o VolumePropertiesExportPolicyPtrOutput) ToVolumePropertiesExportPolicyPtrOutput() VolumePropertiesExportPolicyPtrOutput {
	return o
}

func (o VolumePropertiesExportPolicyPtrOutput) ToVolumePropertiesExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyPtrOutput {
	return o
}

func (o VolumePropertiesExportPolicyPtrOutput) Elem() VolumePropertiesExportPolicyOutput {
	return o.ApplyT(func(v *VolumePropertiesExportPolicy) VolumePropertiesExportPolicy {
		if v != nil {
			return *v
		}
		var ret VolumePropertiesExportPolicy
		return ret
	}).(VolumePropertiesExportPolicyOutput)
}

func (o VolumePropertiesExportPolicyPtrOutput) Rules() ExportPolicyRuleArrayOutput {
	return o.ApplyT(func(v *VolumePropertiesExportPolicy) []ExportPolicyRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ExportPolicyRuleArrayOutput)
}

type VolumePropertiesResponseDataProtection struct {
	Backup      *VolumeBackupPropertiesResponse   `pulumi:"backup"`
	Replication *ReplicationObjectResponse        `pulumi:"replication"`
	Snapshot    *VolumeSnapshotPropertiesResponse `pulumi:"snapshot"`
}





type VolumePropertiesResponseDataProtectionInput interface {
	pulumi.Input

	ToVolumePropertiesResponseDataProtectionOutput() VolumePropertiesResponseDataProtectionOutput
	ToVolumePropertiesResponseDataProtectionOutputWithContext(context.Context) VolumePropertiesResponseDataProtectionOutput
}

type VolumePropertiesResponseDataProtectionArgs struct {
	Backup      VolumeBackupPropertiesResponsePtrInput   `pulumi:"backup"`
	Replication ReplicationObjectResponsePtrInput        `pulumi:"replication"`
	Snapshot    VolumeSnapshotPropertiesResponsePtrInput `pulumi:"snapshot"`
}

func (VolumePropertiesResponseDataProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesResponseDataProtection)(nil)).Elem()
}

func (i VolumePropertiesResponseDataProtectionArgs) ToVolumePropertiesResponseDataProtectionOutput() VolumePropertiesResponseDataProtectionOutput {
	return i.ToVolumePropertiesResponseDataProtectionOutputWithContext(context.Background())
}

func (i VolumePropertiesResponseDataProtectionArgs) ToVolumePropertiesResponseDataProtectionOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseDataProtectionOutput)
}

func (i VolumePropertiesResponseDataProtectionArgs) ToVolumePropertiesResponseDataProtectionPtrOutput() VolumePropertiesResponseDataProtectionPtrOutput {
	return i.ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(context.Background())
}

func (i VolumePropertiesResponseDataProtectionArgs) ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseDataProtectionOutput).ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(ctx)
}









type VolumePropertiesResponseDataProtectionPtrInput interface {
	pulumi.Input

	ToVolumePropertiesResponseDataProtectionPtrOutput() VolumePropertiesResponseDataProtectionPtrOutput
	ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(context.Context) VolumePropertiesResponseDataProtectionPtrOutput
}

type volumePropertiesResponseDataProtectionPtrType VolumePropertiesResponseDataProtectionArgs

func VolumePropertiesResponseDataProtectionPtr(v *VolumePropertiesResponseDataProtectionArgs) VolumePropertiesResponseDataProtectionPtrInput {
	return (*volumePropertiesResponseDataProtectionPtrType)(v)
}

func (*volumePropertiesResponseDataProtectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesResponseDataProtection)(nil)).Elem()
}

func (i *volumePropertiesResponseDataProtectionPtrType) ToVolumePropertiesResponseDataProtectionPtrOutput() VolumePropertiesResponseDataProtectionPtrOutput {
	return i.ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(context.Background())
}

func (i *volumePropertiesResponseDataProtectionPtrType) ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseDataProtectionPtrOutput)
}

type VolumePropertiesResponseDataProtectionOutput struct{ *pulumi.OutputState }

func (VolumePropertiesResponseDataProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesResponseDataProtection)(nil)).Elem()
}

func (o VolumePropertiesResponseDataProtectionOutput) ToVolumePropertiesResponseDataProtectionOutput() VolumePropertiesResponseDataProtectionOutput {
	return o
}

func (o VolumePropertiesResponseDataProtectionOutput) ToVolumePropertiesResponseDataProtectionOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionOutput {
	return o
}

func (o VolumePropertiesResponseDataProtectionOutput) ToVolumePropertiesResponseDataProtectionPtrOutput() VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(context.Background())
}

func (o VolumePropertiesResponseDataProtectionOutput) ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumePropertiesResponseDataProtection) *VolumePropertiesResponseDataProtection {
		return &v
	}).(VolumePropertiesResponseDataProtectionPtrOutput)
}

func (o VolumePropertiesResponseDataProtectionOutput) Backup() VolumeBackupPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VolumePropertiesResponseDataProtection) *VolumeBackupPropertiesResponse { return v.Backup }).(VolumeBackupPropertiesResponsePtrOutput)
}

func (o VolumePropertiesResponseDataProtectionOutput) Replication() ReplicationObjectResponsePtrOutput {
	return o.ApplyT(func(v VolumePropertiesResponseDataProtection) *ReplicationObjectResponse { return v.Replication }).(ReplicationObjectResponsePtrOutput)
}

func (o VolumePropertiesResponseDataProtectionOutput) Snapshot() VolumeSnapshotPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VolumePropertiesResponseDataProtection) *VolumeSnapshotPropertiesResponse { return v.Snapshot }).(VolumeSnapshotPropertiesResponsePtrOutput)
}

type VolumePropertiesResponseDataProtectionPtrOutput struct{ *pulumi.OutputState }

func (VolumePropertiesResponseDataProtectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesResponseDataProtection)(nil)).Elem()
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) ToVolumePropertiesResponseDataProtectionPtrOutput() VolumePropertiesResponseDataProtectionPtrOutput {
	return o
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionPtrOutput {
	return o
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) Elem() VolumePropertiesResponseDataProtectionOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseDataProtection) VolumePropertiesResponseDataProtection {
		if v != nil {
			return *v
		}
		var ret VolumePropertiesResponseDataProtection
		return ret
	}).(VolumePropertiesResponseDataProtectionOutput)
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) Backup() VolumeBackupPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseDataProtection) *VolumeBackupPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Backup
	}).(VolumeBackupPropertiesResponsePtrOutput)
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) Replication() ReplicationObjectResponsePtrOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseDataProtection) *ReplicationObjectResponse {
		if v == nil {
			return nil
		}
		return v.Replication
	}).(ReplicationObjectResponsePtrOutput)
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) Snapshot() VolumeSnapshotPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseDataProtection) *VolumeSnapshotPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Snapshot
	}).(VolumeSnapshotPropertiesResponsePtrOutput)
}

type VolumePropertiesResponseExportPolicy struct {
	Rules []ExportPolicyRuleResponse `pulumi:"rules"`
}





type VolumePropertiesResponseExportPolicyInput interface {
	pulumi.Input

	ToVolumePropertiesResponseExportPolicyOutput() VolumePropertiesResponseExportPolicyOutput
	ToVolumePropertiesResponseExportPolicyOutputWithContext(context.Context) VolumePropertiesResponseExportPolicyOutput
}

type VolumePropertiesResponseExportPolicyArgs struct {
	Rules ExportPolicyRuleResponseArrayInput `pulumi:"rules"`
}

func (VolumePropertiesResponseExportPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesResponseExportPolicy)(nil)).Elem()
}

func (i VolumePropertiesResponseExportPolicyArgs) ToVolumePropertiesResponseExportPolicyOutput() VolumePropertiesResponseExportPolicyOutput {
	return i.ToVolumePropertiesResponseExportPolicyOutputWithContext(context.Background())
}

func (i VolumePropertiesResponseExportPolicyArgs) ToVolumePropertiesResponseExportPolicyOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseExportPolicyOutput)
}

func (i VolumePropertiesResponseExportPolicyArgs) ToVolumePropertiesResponseExportPolicyPtrOutput() VolumePropertiesResponseExportPolicyPtrOutput {
	return i.ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(context.Background())
}

func (i VolumePropertiesResponseExportPolicyArgs) ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseExportPolicyOutput).ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(ctx)
}









type VolumePropertiesResponseExportPolicyPtrInput interface {
	pulumi.Input

	ToVolumePropertiesResponseExportPolicyPtrOutput() VolumePropertiesResponseExportPolicyPtrOutput
	ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(context.Context) VolumePropertiesResponseExportPolicyPtrOutput
}

type volumePropertiesResponseExportPolicyPtrType VolumePropertiesResponseExportPolicyArgs

func VolumePropertiesResponseExportPolicyPtr(v *VolumePropertiesResponseExportPolicyArgs) VolumePropertiesResponseExportPolicyPtrInput {
	return (*volumePropertiesResponseExportPolicyPtrType)(v)
}

func (*volumePropertiesResponseExportPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesResponseExportPolicy)(nil)).Elem()
}

func (i *volumePropertiesResponseExportPolicyPtrType) ToVolumePropertiesResponseExportPolicyPtrOutput() VolumePropertiesResponseExportPolicyPtrOutput {
	return i.ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(context.Background())
}

func (i *volumePropertiesResponseExportPolicyPtrType) ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseExportPolicyPtrOutput)
}

type VolumePropertiesResponseExportPolicyOutput struct{ *pulumi.OutputState }

func (VolumePropertiesResponseExportPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesResponseExportPolicy)(nil)).Elem()
}

func (o VolumePropertiesResponseExportPolicyOutput) ToVolumePropertiesResponseExportPolicyOutput() VolumePropertiesResponseExportPolicyOutput {
	return o
}

func (o VolumePropertiesResponseExportPolicyOutput) ToVolumePropertiesResponseExportPolicyOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyOutput {
	return o
}

func (o VolumePropertiesResponseExportPolicyOutput) ToVolumePropertiesResponseExportPolicyPtrOutput() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(context.Background())
}

func (o VolumePropertiesResponseExportPolicyOutput) ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumePropertiesResponseExportPolicy) *VolumePropertiesResponseExportPolicy {
		return &v
	}).(VolumePropertiesResponseExportPolicyPtrOutput)
}

func (o VolumePropertiesResponseExportPolicyOutput) Rules() ExportPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v VolumePropertiesResponseExportPolicy) []ExportPolicyRuleResponse { return v.Rules }).(ExportPolicyRuleResponseArrayOutput)
}

type VolumePropertiesResponseExportPolicyPtrOutput struct{ *pulumi.OutputState }

func (VolumePropertiesResponseExportPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesResponseExportPolicy)(nil)).Elem()
}

func (o VolumePropertiesResponseExportPolicyPtrOutput) ToVolumePropertiesResponseExportPolicyPtrOutput() VolumePropertiesResponseExportPolicyPtrOutput {
	return o
}

func (o VolumePropertiesResponseExportPolicyPtrOutput) ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyPtrOutput {
	return o
}

func (o VolumePropertiesResponseExportPolicyPtrOutput) Elem() VolumePropertiesResponseExportPolicyOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseExportPolicy) VolumePropertiesResponseExportPolicy {
		if v != nil {
			return *v
		}
		var ret VolumePropertiesResponseExportPolicy
		return ret
	}).(VolumePropertiesResponseExportPolicyOutput)
}

func (o VolumePropertiesResponseExportPolicyPtrOutput) Rules() ExportPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseExportPolicy) []ExportPolicyRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ExportPolicyRuleResponseArrayOutput)
}

type VolumeSnapshotProperties struct {
	SnapshotPolicyId *string `pulumi:"snapshotPolicyId"`
}





type VolumeSnapshotPropertiesInput interface {
	pulumi.Input

	ToVolumeSnapshotPropertiesOutput() VolumeSnapshotPropertiesOutput
	ToVolumeSnapshotPropertiesOutputWithContext(context.Context) VolumeSnapshotPropertiesOutput
}

type VolumeSnapshotPropertiesArgs struct {
	SnapshotPolicyId pulumi.StringPtrInput `pulumi:"snapshotPolicyId"`
}

func (VolumeSnapshotPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSnapshotProperties)(nil)).Elem()
}

func (i VolumeSnapshotPropertiesArgs) ToVolumeSnapshotPropertiesOutput() VolumeSnapshotPropertiesOutput {
	return i.ToVolumeSnapshotPropertiesOutputWithContext(context.Background())
}

func (i VolumeSnapshotPropertiesArgs) ToVolumeSnapshotPropertiesOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesOutput)
}

func (i VolumeSnapshotPropertiesArgs) ToVolumeSnapshotPropertiesPtrOutput() VolumeSnapshotPropertiesPtrOutput {
	return i.ToVolumeSnapshotPropertiesPtrOutputWithContext(context.Background())
}

func (i VolumeSnapshotPropertiesArgs) ToVolumeSnapshotPropertiesPtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesOutput).ToVolumeSnapshotPropertiesPtrOutputWithContext(ctx)
}









type VolumeSnapshotPropertiesPtrInput interface {
	pulumi.Input

	ToVolumeSnapshotPropertiesPtrOutput() VolumeSnapshotPropertiesPtrOutput
	ToVolumeSnapshotPropertiesPtrOutputWithContext(context.Context) VolumeSnapshotPropertiesPtrOutput
}

type volumeSnapshotPropertiesPtrType VolumeSnapshotPropertiesArgs

func VolumeSnapshotPropertiesPtr(v *VolumeSnapshotPropertiesArgs) VolumeSnapshotPropertiesPtrInput {
	return (*volumeSnapshotPropertiesPtrType)(v)
}

func (*volumeSnapshotPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeSnapshotProperties)(nil)).Elem()
}

func (i *volumeSnapshotPropertiesPtrType) ToVolumeSnapshotPropertiesPtrOutput() VolumeSnapshotPropertiesPtrOutput {
	return i.ToVolumeSnapshotPropertiesPtrOutputWithContext(context.Background())
}

func (i *volumeSnapshotPropertiesPtrType) ToVolumeSnapshotPropertiesPtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesPtrOutput)
}

type VolumeSnapshotPropertiesOutput struct{ *pulumi.OutputState }

func (VolumeSnapshotPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSnapshotProperties)(nil)).Elem()
}

func (o VolumeSnapshotPropertiesOutput) ToVolumeSnapshotPropertiesOutput() VolumeSnapshotPropertiesOutput {
	return o
}

func (o VolumeSnapshotPropertiesOutput) ToVolumeSnapshotPropertiesOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesOutput {
	return o
}

func (o VolumeSnapshotPropertiesOutput) ToVolumeSnapshotPropertiesPtrOutput() VolumeSnapshotPropertiesPtrOutput {
	return o.ToVolumeSnapshotPropertiesPtrOutputWithContext(context.Background())
}

func (o VolumeSnapshotPropertiesOutput) ToVolumeSnapshotPropertiesPtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeSnapshotProperties) *VolumeSnapshotProperties {
		return &v
	}).(VolumeSnapshotPropertiesPtrOutput)
}

func (o VolumeSnapshotPropertiesOutput) SnapshotPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeSnapshotProperties) *string { return v.SnapshotPolicyId }).(pulumi.StringPtrOutput)
}

type VolumeSnapshotPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VolumeSnapshotPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeSnapshotProperties)(nil)).Elem()
}

func (o VolumeSnapshotPropertiesPtrOutput) ToVolumeSnapshotPropertiesPtrOutput() VolumeSnapshotPropertiesPtrOutput {
	return o
}

func (o VolumeSnapshotPropertiesPtrOutput) ToVolumeSnapshotPropertiesPtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesPtrOutput {
	return o
}

func (o VolumeSnapshotPropertiesPtrOutput) Elem() VolumeSnapshotPropertiesOutput {
	return o.ApplyT(func(v *VolumeSnapshotProperties) VolumeSnapshotProperties {
		if v != nil {
			return *v
		}
		var ret VolumeSnapshotProperties
		return ret
	}).(VolumeSnapshotPropertiesOutput)
}

func (o VolumeSnapshotPropertiesPtrOutput) SnapshotPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeSnapshotProperties) *string {
		if v == nil {
			return nil
		}
		return v.SnapshotPolicyId
	}).(pulumi.StringPtrOutput)
}

type VolumeSnapshotPropertiesResponse struct {
	SnapshotPolicyId *string `pulumi:"snapshotPolicyId"`
}





type VolumeSnapshotPropertiesResponseInput interface {
	pulumi.Input

	ToVolumeSnapshotPropertiesResponseOutput() VolumeSnapshotPropertiesResponseOutput
	ToVolumeSnapshotPropertiesResponseOutputWithContext(context.Context) VolumeSnapshotPropertiesResponseOutput
}

type VolumeSnapshotPropertiesResponseArgs struct {
	SnapshotPolicyId pulumi.StringPtrInput `pulumi:"snapshotPolicyId"`
}

func (VolumeSnapshotPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSnapshotPropertiesResponse)(nil)).Elem()
}

func (i VolumeSnapshotPropertiesResponseArgs) ToVolumeSnapshotPropertiesResponseOutput() VolumeSnapshotPropertiesResponseOutput {
	return i.ToVolumeSnapshotPropertiesResponseOutputWithContext(context.Background())
}

func (i VolumeSnapshotPropertiesResponseArgs) ToVolumeSnapshotPropertiesResponseOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesResponseOutput)
}

func (i VolumeSnapshotPropertiesResponseArgs) ToVolumeSnapshotPropertiesResponsePtrOutput() VolumeSnapshotPropertiesResponsePtrOutput {
	return i.ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i VolumeSnapshotPropertiesResponseArgs) ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesResponseOutput).ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(ctx)
}









type VolumeSnapshotPropertiesResponsePtrInput interface {
	pulumi.Input

	ToVolumeSnapshotPropertiesResponsePtrOutput() VolumeSnapshotPropertiesResponsePtrOutput
	ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(context.Context) VolumeSnapshotPropertiesResponsePtrOutput
}

type volumeSnapshotPropertiesResponsePtrType VolumeSnapshotPropertiesResponseArgs

func VolumeSnapshotPropertiesResponsePtr(v *VolumeSnapshotPropertiesResponseArgs) VolumeSnapshotPropertiesResponsePtrInput {
	return (*volumeSnapshotPropertiesResponsePtrType)(v)
}

func (*volumeSnapshotPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeSnapshotPropertiesResponse)(nil)).Elem()
}

func (i *volumeSnapshotPropertiesResponsePtrType) ToVolumeSnapshotPropertiesResponsePtrOutput() VolumeSnapshotPropertiesResponsePtrOutput {
	return i.ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *volumeSnapshotPropertiesResponsePtrType) ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesResponsePtrOutput)
}

type VolumeSnapshotPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VolumeSnapshotPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSnapshotPropertiesResponse)(nil)).Elem()
}

func (o VolumeSnapshotPropertiesResponseOutput) ToVolumeSnapshotPropertiesResponseOutput() VolumeSnapshotPropertiesResponseOutput {
	return o
}

func (o VolumeSnapshotPropertiesResponseOutput) ToVolumeSnapshotPropertiesResponseOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponseOutput {
	return o
}

func (o VolumeSnapshotPropertiesResponseOutput) ToVolumeSnapshotPropertiesResponsePtrOutput() VolumeSnapshotPropertiesResponsePtrOutput {
	return o.ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o VolumeSnapshotPropertiesResponseOutput) ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeSnapshotPropertiesResponse) *VolumeSnapshotPropertiesResponse {
		return &v
	}).(VolumeSnapshotPropertiesResponsePtrOutput)
}

func (o VolumeSnapshotPropertiesResponseOutput) SnapshotPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeSnapshotPropertiesResponse) *string { return v.SnapshotPolicyId }).(pulumi.StringPtrOutput)
}

type VolumeSnapshotPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VolumeSnapshotPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeSnapshotPropertiesResponse)(nil)).Elem()
}

func (o VolumeSnapshotPropertiesResponsePtrOutput) ToVolumeSnapshotPropertiesResponsePtrOutput() VolumeSnapshotPropertiesResponsePtrOutput {
	return o
}

func (o VolumeSnapshotPropertiesResponsePtrOutput) ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponsePtrOutput {
	return o
}

func (o VolumeSnapshotPropertiesResponsePtrOutput) Elem() VolumeSnapshotPropertiesResponseOutput {
	return o.ApplyT(func(v *VolumeSnapshotPropertiesResponse) VolumeSnapshotPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VolumeSnapshotPropertiesResponse
		return ret
	}).(VolumeSnapshotPropertiesResponseOutput)
}

func (o VolumeSnapshotPropertiesResponsePtrOutput) SnapshotPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeSnapshotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SnapshotPolicyId
	}).(pulumi.StringPtrOutput)
}

type WeeklySchedule struct {
	Day             *string  `pulumi:"day"`
	Hour            *int     `pulumi:"hour"`
	Minute          *int     `pulumi:"minute"`
	SnapshotsToKeep *int     `pulumi:"snapshotsToKeep"`
	UsedBytes       *float64 `pulumi:"usedBytes"`
}





type WeeklyScheduleInput interface {
	pulumi.Input

	ToWeeklyScheduleOutput() WeeklyScheduleOutput
	ToWeeklyScheduleOutputWithContext(context.Context) WeeklyScheduleOutput
}

type WeeklyScheduleArgs struct {
	Day             pulumi.StringPtrInput  `pulumi:"day"`
	Hour            pulumi.IntPtrInput     `pulumi:"hour"`
	Minute          pulumi.IntPtrInput     `pulumi:"minute"`
	SnapshotsToKeep pulumi.IntPtrInput     `pulumi:"snapshotsToKeep"`
	UsedBytes       pulumi.Float64PtrInput `pulumi:"usedBytes"`
}

func (WeeklyScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklySchedule)(nil)).Elem()
}

func (i WeeklyScheduleArgs) ToWeeklyScheduleOutput() WeeklyScheduleOutput {
	return i.ToWeeklyScheduleOutputWithContext(context.Background())
}

func (i WeeklyScheduleArgs) ToWeeklyScheduleOutputWithContext(ctx context.Context) WeeklyScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyScheduleOutput)
}

func (i WeeklyScheduleArgs) ToWeeklySchedulePtrOutput() WeeklySchedulePtrOutput {
	return i.ToWeeklySchedulePtrOutputWithContext(context.Background())
}

func (i WeeklyScheduleArgs) ToWeeklySchedulePtrOutputWithContext(ctx context.Context) WeeklySchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyScheduleOutput).ToWeeklySchedulePtrOutputWithContext(ctx)
}









type WeeklySchedulePtrInput interface {
	pulumi.Input

	ToWeeklySchedulePtrOutput() WeeklySchedulePtrOutput
	ToWeeklySchedulePtrOutputWithContext(context.Context) WeeklySchedulePtrOutput
}

type weeklySchedulePtrType WeeklyScheduleArgs

func WeeklySchedulePtr(v *WeeklyScheduleArgs) WeeklySchedulePtrInput {
	return (*weeklySchedulePtrType)(v)
}

func (*weeklySchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklySchedule)(nil)).Elem()
}

func (i *weeklySchedulePtrType) ToWeeklySchedulePtrOutput() WeeklySchedulePtrOutput {
	return i.ToWeeklySchedulePtrOutputWithContext(context.Background())
}

func (i *weeklySchedulePtrType) ToWeeklySchedulePtrOutputWithContext(ctx context.Context) WeeklySchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklySchedulePtrOutput)
}

type WeeklyScheduleOutput struct{ *pulumi.OutputState }

func (WeeklyScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklySchedule)(nil)).Elem()
}

func (o WeeklyScheduleOutput) ToWeeklyScheduleOutput() WeeklyScheduleOutput {
	return o
}

func (o WeeklyScheduleOutput) ToWeeklyScheduleOutputWithContext(ctx context.Context) WeeklyScheduleOutput {
	return o
}

func (o WeeklyScheduleOutput) ToWeeklySchedulePtrOutput() WeeklySchedulePtrOutput {
	return o.ToWeeklySchedulePtrOutputWithContext(context.Background())
}

func (o WeeklyScheduleOutput) ToWeeklySchedulePtrOutputWithContext(ctx context.Context) WeeklySchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeeklySchedule) *WeeklySchedule {
		return &v
	}).(WeeklySchedulePtrOutput)
}

func (o WeeklyScheduleOutput) Day() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WeeklySchedule) *string { return v.Day }).(pulumi.StringPtrOutput)
}

func (o WeeklyScheduleOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WeeklySchedule) *int { return v.Hour }).(pulumi.IntPtrOutput)
}

func (o WeeklyScheduleOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WeeklySchedule) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

func (o WeeklyScheduleOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WeeklySchedule) *int { return v.SnapshotsToKeep }).(pulumi.IntPtrOutput)
}

func (o WeeklyScheduleOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WeeklySchedule) *float64 { return v.UsedBytes }).(pulumi.Float64PtrOutput)
}

type WeeklySchedulePtrOutput struct{ *pulumi.OutputState }

func (WeeklySchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklySchedule)(nil)).Elem()
}

func (o WeeklySchedulePtrOutput) ToWeeklySchedulePtrOutput() WeeklySchedulePtrOutput {
	return o
}

func (o WeeklySchedulePtrOutput) ToWeeklySchedulePtrOutputWithContext(ctx context.Context) WeeklySchedulePtrOutput {
	return o
}

func (o WeeklySchedulePtrOutput) Elem() WeeklyScheduleOutput {
	return o.ApplyT(func(v *WeeklySchedule) WeeklySchedule {
		if v != nil {
			return *v
		}
		var ret WeeklySchedule
		return ret
	}).(WeeklyScheduleOutput)
}

func (o WeeklySchedulePtrOutput) Day() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WeeklySchedule) *string {
		if v == nil {
			return nil
		}
		return v.Day
	}).(pulumi.StringPtrOutput)
}

func (o WeeklySchedulePtrOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WeeklySchedule) *int {
		if v == nil {
			return nil
		}
		return v.Hour
	}).(pulumi.IntPtrOutput)
}

func (o WeeklySchedulePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WeeklySchedule) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

func (o WeeklySchedulePtrOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WeeklySchedule) *int {
		if v == nil {
			return nil
		}
		return v.SnapshotsToKeep
	}).(pulumi.IntPtrOutput)
}

func (o WeeklySchedulePtrOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WeeklySchedule) *float64 {
		if v == nil {
			return nil
		}
		return v.UsedBytes
	}).(pulumi.Float64PtrOutput)
}

type WeeklyScheduleResponse struct {
	Day             *string  `pulumi:"day"`
	Hour            *int     `pulumi:"hour"`
	Minute          *int     `pulumi:"minute"`
	SnapshotsToKeep *int     `pulumi:"snapshotsToKeep"`
	UsedBytes       *float64 `pulumi:"usedBytes"`
}





type WeeklyScheduleResponseInput interface {
	pulumi.Input

	ToWeeklyScheduleResponseOutput() WeeklyScheduleResponseOutput
	ToWeeklyScheduleResponseOutputWithContext(context.Context) WeeklyScheduleResponseOutput
}

type WeeklyScheduleResponseArgs struct {
	Day             pulumi.StringPtrInput  `pulumi:"day"`
	Hour            pulumi.IntPtrInput     `pulumi:"hour"`
	Minute          pulumi.IntPtrInput     `pulumi:"minute"`
	SnapshotsToKeep pulumi.IntPtrInput     `pulumi:"snapshotsToKeep"`
	UsedBytes       pulumi.Float64PtrInput `pulumi:"usedBytes"`
}

func (WeeklyScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyScheduleResponse)(nil)).Elem()
}

func (i WeeklyScheduleResponseArgs) ToWeeklyScheduleResponseOutput() WeeklyScheduleResponseOutput {
	return i.ToWeeklyScheduleResponseOutputWithContext(context.Background())
}

func (i WeeklyScheduleResponseArgs) ToWeeklyScheduleResponseOutputWithContext(ctx context.Context) WeeklyScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyScheduleResponseOutput)
}

func (i WeeklyScheduleResponseArgs) ToWeeklyScheduleResponsePtrOutput() WeeklyScheduleResponsePtrOutput {
	return i.ToWeeklyScheduleResponsePtrOutputWithContext(context.Background())
}

func (i WeeklyScheduleResponseArgs) ToWeeklyScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyScheduleResponseOutput).ToWeeklyScheduleResponsePtrOutputWithContext(ctx)
}









type WeeklyScheduleResponsePtrInput interface {
	pulumi.Input

	ToWeeklyScheduleResponsePtrOutput() WeeklyScheduleResponsePtrOutput
	ToWeeklyScheduleResponsePtrOutputWithContext(context.Context) WeeklyScheduleResponsePtrOutput
}

type weeklyScheduleResponsePtrType WeeklyScheduleResponseArgs

func WeeklyScheduleResponsePtr(v *WeeklyScheduleResponseArgs) WeeklyScheduleResponsePtrInput {
	return (*weeklyScheduleResponsePtrType)(v)
}

func (*weeklyScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyScheduleResponse)(nil)).Elem()
}

func (i *weeklyScheduleResponsePtrType) ToWeeklyScheduleResponsePtrOutput() WeeklyScheduleResponsePtrOutput {
	return i.ToWeeklyScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *weeklyScheduleResponsePtrType) ToWeeklyScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyScheduleResponsePtrOutput)
}

type WeeklyScheduleResponseOutput struct{ *pulumi.OutputState }

func (WeeklyScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyScheduleResponse)(nil)).Elem()
}

func (o WeeklyScheduleResponseOutput) ToWeeklyScheduleResponseOutput() WeeklyScheduleResponseOutput {
	return o
}

func (o WeeklyScheduleResponseOutput) ToWeeklyScheduleResponseOutputWithContext(ctx context.Context) WeeklyScheduleResponseOutput {
	return o
}

func (o WeeklyScheduleResponseOutput) ToWeeklyScheduleResponsePtrOutput() WeeklyScheduleResponsePtrOutput {
	return o.ToWeeklyScheduleResponsePtrOutputWithContext(context.Background())
}

func (o WeeklyScheduleResponseOutput) ToWeeklyScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeeklyScheduleResponse) *WeeklyScheduleResponse {
		return &v
	}).(WeeklyScheduleResponsePtrOutput)
}

func (o WeeklyScheduleResponseOutput) Day() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WeeklyScheduleResponse) *string { return v.Day }).(pulumi.StringPtrOutput)
}

func (o WeeklyScheduleResponseOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WeeklyScheduleResponse) *int { return v.Hour }).(pulumi.IntPtrOutput)
}

func (o WeeklyScheduleResponseOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WeeklyScheduleResponse) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

func (o WeeklyScheduleResponseOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WeeklyScheduleResponse) *int { return v.SnapshotsToKeep }).(pulumi.IntPtrOutput)
}

func (o WeeklyScheduleResponseOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WeeklyScheduleResponse) *float64 { return v.UsedBytes }).(pulumi.Float64PtrOutput)
}

type WeeklyScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (WeeklyScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyScheduleResponse)(nil)).Elem()
}

func (o WeeklyScheduleResponsePtrOutput) ToWeeklyScheduleResponsePtrOutput() WeeklyScheduleResponsePtrOutput {
	return o
}

func (o WeeklyScheduleResponsePtrOutput) ToWeeklyScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyScheduleResponsePtrOutput {
	return o
}

func (o WeeklyScheduleResponsePtrOutput) Elem() WeeklyScheduleResponseOutput {
	return o.ApplyT(func(v *WeeklyScheduleResponse) WeeklyScheduleResponse {
		if v != nil {
			return *v
		}
		var ret WeeklyScheduleResponse
		return ret
	}).(WeeklyScheduleResponseOutput)
}

func (o WeeklyScheduleResponsePtrOutput) Day() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WeeklyScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Day
	}).(pulumi.StringPtrOutput)
}

func (o WeeklyScheduleResponsePtrOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WeeklyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.Hour
	}).(pulumi.IntPtrOutput)
}

func (o WeeklyScheduleResponsePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WeeklyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

func (o WeeklyScheduleResponsePtrOutput) SnapshotsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WeeklyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.SnapshotsToKeep
	}).(pulumi.IntPtrOutput)
}

func (o WeeklyScheduleResponsePtrOutput) UsedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WeeklyScheduleResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.UsedBytes
	}).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountEncryptionOutput{})
	pulumi.RegisterOutputType(AccountEncryptionPtrOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponseOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryArrayOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryResponseArrayOutput{})
	pulumi.RegisterOutputType(DailyScheduleOutput{})
	pulumi.RegisterOutputType(DailySchedulePtrOutput{})
	pulumi.RegisterOutputType(DailyScheduleResponseOutput{})
	pulumi.RegisterOutputType(DailyScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleResponseOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(HourlyScheduleOutput{})
	pulumi.RegisterOutputType(HourlySchedulePtrOutput{})
	pulumi.RegisterOutputType(HourlyScheduleResponseOutput{})
	pulumi.RegisterOutputType(HourlyScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(MonthlyScheduleOutput{})
	pulumi.RegisterOutputType(MonthlySchedulePtrOutput{})
	pulumi.RegisterOutputType(MonthlyScheduleResponseOutput{})
	pulumi.RegisterOutputType(MonthlyScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(MountTargetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MountTargetPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(PlacementKeyValuePairsOutput{})
	pulumi.RegisterOutputType(PlacementKeyValuePairsArrayOutput{})
	pulumi.RegisterOutputType(PlacementKeyValuePairsResponseOutput{})
	pulumi.RegisterOutputType(PlacementKeyValuePairsResponseArrayOutput{})
	pulumi.RegisterOutputType(ReplicationObjectOutput{})
	pulumi.RegisterOutputType(ReplicationObjectPtrOutput{})
	pulumi.RegisterOutputType(ReplicationObjectResponseOutput{})
	pulumi.RegisterOutputType(ReplicationObjectResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(VolumeBackupPropertiesOutput{})
	pulumi.RegisterOutputType(VolumeBackupPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VolumeBackupPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VolumeBackupPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VolumeBackupsResponseOutput{})
	pulumi.RegisterOutputType(VolumeBackupsResponseArrayOutput{})
	pulumi.RegisterOutputType(VolumeGroupMetaDataOutput{})
	pulumi.RegisterOutputType(VolumeGroupMetaDataPtrOutput{})
	pulumi.RegisterOutputType(VolumeGroupMetaDataResponseOutput{})
	pulumi.RegisterOutputType(VolumeGroupMetaDataResponsePtrOutput{})
	pulumi.RegisterOutputType(VolumeGroupVolumePropertiesOutput{})
	pulumi.RegisterOutputType(VolumeGroupVolumePropertiesArrayOutput{})
	pulumi.RegisterOutputType(VolumeGroupVolumePropertiesResponseOutput{})
	pulumi.RegisterOutputType(VolumeGroupVolumePropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(VolumePropertiesDataProtectionOutput{})
	pulumi.RegisterOutputType(VolumePropertiesDataProtectionPtrOutput{})
	pulumi.RegisterOutputType(VolumePropertiesExportPolicyOutput{})
	pulumi.RegisterOutputType(VolumePropertiesExportPolicyPtrOutput{})
	pulumi.RegisterOutputType(VolumePropertiesResponseDataProtectionOutput{})
	pulumi.RegisterOutputType(VolumePropertiesResponseDataProtectionPtrOutput{})
	pulumi.RegisterOutputType(VolumePropertiesResponseExportPolicyOutput{})
	pulumi.RegisterOutputType(VolumePropertiesResponseExportPolicyPtrOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotPropertiesOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(WeeklyScheduleOutput{})
	pulumi.RegisterOutputType(WeeklySchedulePtrOutput{})
	pulumi.RegisterOutputType(WeeklyScheduleResponseOutput{})
	pulumi.RegisterOutputType(WeeklyScheduleResponsePtrOutput{})
}
