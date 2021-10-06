


package storagecache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobNfsTarget struct {
	Target     *string `pulumi:"target"`
	UsageModel *string `pulumi:"usageModel"`
}





type BlobNfsTargetInput interface {
	pulumi.Input

	ToBlobNfsTargetOutput() BlobNfsTargetOutput
	ToBlobNfsTargetOutputWithContext(context.Context) BlobNfsTargetOutput
}

type BlobNfsTargetArgs struct {
	Target     pulumi.StringPtrInput `pulumi:"target"`
	UsageModel pulumi.StringPtrInput `pulumi:"usageModel"`
}

func (BlobNfsTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobNfsTarget)(nil)).Elem()
}

func (i BlobNfsTargetArgs) ToBlobNfsTargetOutput() BlobNfsTargetOutput {
	return i.ToBlobNfsTargetOutputWithContext(context.Background())
}

func (i BlobNfsTargetArgs) ToBlobNfsTargetOutputWithContext(ctx context.Context) BlobNfsTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobNfsTargetOutput)
}

func (i BlobNfsTargetArgs) ToBlobNfsTargetPtrOutput() BlobNfsTargetPtrOutput {
	return i.ToBlobNfsTargetPtrOutputWithContext(context.Background())
}

func (i BlobNfsTargetArgs) ToBlobNfsTargetPtrOutputWithContext(ctx context.Context) BlobNfsTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobNfsTargetOutput).ToBlobNfsTargetPtrOutputWithContext(ctx)
}









type BlobNfsTargetPtrInput interface {
	pulumi.Input

	ToBlobNfsTargetPtrOutput() BlobNfsTargetPtrOutput
	ToBlobNfsTargetPtrOutputWithContext(context.Context) BlobNfsTargetPtrOutput
}

type blobNfsTargetPtrType BlobNfsTargetArgs

func BlobNfsTargetPtr(v *BlobNfsTargetArgs) BlobNfsTargetPtrInput {
	return (*blobNfsTargetPtrType)(v)
}

func (*blobNfsTargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobNfsTarget)(nil)).Elem()
}

func (i *blobNfsTargetPtrType) ToBlobNfsTargetPtrOutput() BlobNfsTargetPtrOutput {
	return i.ToBlobNfsTargetPtrOutputWithContext(context.Background())
}

func (i *blobNfsTargetPtrType) ToBlobNfsTargetPtrOutputWithContext(ctx context.Context) BlobNfsTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobNfsTargetPtrOutput)
}

type BlobNfsTargetOutput struct{ *pulumi.OutputState }

func (BlobNfsTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobNfsTarget)(nil)).Elem()
}

func (o BlobNfsTargetOutput) ToBlobNfsTargetOutput() BlobNfsTargetOutput {
	return o
}

func (o BlobNfsTargetOutput) ToBlobNfsTargetOutputWithContext(ctx context.Context) BlobNfsTargetOutput {
	return o
}

func (o BlobNfsTargetOutput) ToBlobNfsTargetPtrOutput() BlobNfsTargetPtrOutput {
	return o.ToBlobNfsTargetPtrOutputWithContext(context.Background())
}

func (o BlobNfsTargetOutput) ToBlobNfsTargetPtrOutputWithContext(ctx context.Context) BlobNfsTargetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobNfsTarget) *BlobNfsTarget {
		return &v
	}).(BlobNfsTargetPtrOutput)
}

func (o BlobNfsTargetOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobNfsTarget) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o BlobNfsTargetOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobNfsTarget) *string { return v.UsageModel }).(pulumi.StringPtrOutput)
}

type BlobNfsTargetPtrOutput struct{ *pulumi.OutputState }

func (BlobNfsTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobNfsTarget)(nil)).Elem()
}

func (o BlobNfsTargetPtrOutput) ToBlobNfsTargetPtrOutput() BlobNfsTargetPtrOutput {
	return o
}

func (o BlobNfsTargetPtrOutput) ToBlobNfsTargetPtrOutputWithContext(ctx context.Context) BlobNfsTargetPtrOutput {
	return o
}

func (o BlobNfsTargetPtrOutput) Elem() BlobNfsTargetOutput {
	return o.ApplyT(func(v *BlobNfsTarget) BlobNfsTarget {
		if v != nil {
			return *v
		}
		var ret BlobNfsTarget
		return ret
	}).(BlobNfsTargetOutput)
}

func (o BlobNfsTargetPtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobNfsTarget) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

func (o BlobNfsTargetPtrOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobNfsTarget) *string {
		if v == nil {
			return nil
		}
		return v.UsageModel
	}).(pulumi.StringPtrOutput)
}

type BlobNfsTargetResponse struct {
	Target     *string `pulumi:"target"`
	UsageModel *string `pulumi:"usageModel"`
}





type BlobNfsTargetResponseInput interface {
	pulumi.Input

	ToBlobNfsTargetResponseOutput() BlobNfsTargetResponseOutput
	ToBlobNfsTargetResponseOutputWithContext(context.Context) BlobNfsTargetResponseOutput
}

type BlobNfsTargetResponseArgs struct {
	Target     pulumi.StringPtrInput `pulumi:"target"`
	UsageModel pulumi.StringPtrInput `pulumi:"usageModel"`
}

func (BlobNfsTargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobNfsTargetResponse)(nil)).Elem()
}

func (i BlobNfsTargetResponseArgs) ToBlobNfsTargetResponseOutput() BlobNfsTargetResponseOutput {
	return i.ToBlobNfsTargetResponseOutputWithContext(context.Background())
}

func (i BlobNfsTargetResponseArgs) ToBlobNfsTargetResponseOutputWithContext(ctx context.Context) BlobNfsTargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobNfsTargetResponseOutput)
}

func (i BlobNfsTargetResponseArgs) ToBlobNfsTargetResponsePtrOutput() BlobNfsTargetResponsePtrOutput {
	return i.ToBlobNfsTargetResponsePtrOutputWithContext(context.Background())
}

func (i BlobNfsTargetResponseArgs) ToBlobNfsTargetResponsePtrOutputWithContext(ctx context.Context) BlobNfsTargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobNfsTargetResponseOutput).ToBlobNfsTargetResponsePtrOutputWithContext(ctx)
}









type BlobNfsTargetResponsePtrInput interface {
	pulumi.Input

	ToBlobNfsTargetResponsePtrOutput() BlobNfsTargetResponsePtrOutput
	ToBlobNfsTargetResponsePtrOutputWithContext(context.Context) BlobNfsTargetResponsePtrOutput
}

type blobNfsTargetResponsePtrType BlobNfsTargetResponseArgs

func BlobNfsTargetResponsePtr(v *BlobNfsTargetResponseArgs) BlobNfsTargetResponsePtrInput {
	return (*blobNfsTargetResponsePtrType)(v)
}

func (*blobNfsTargetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobNfsTargetResponse)(nil)).Elem()
}

func (i *blobNfsTargetResponsePtrType) ToBlobNfsTargetResponsePtrOutput() BlobNfsTargetResponsePtrOutput {
	return i.ToBlobNfsTargetResponsePtrOutputWithContext(context.Background())
}

func (i *blobNfsTargetResponsePtrType) ToBlobNfsTargetResponsePtrOutputWithContext(ctx context.Context) BlobNfsTargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobNfsTargetResponsePtrOutput)
}

type BlobNfsTargetResponseOutput struct{ *pulumi.OutputState }

func (BlobNfsTargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobNfsTargetResponse)(nil)).Elem()
}

func (o BlobNfsTargetResponseOutput) ToBlobNfsTargetResponseOutput() BlobNfsTargetResponseOutput {
	return o
}

func (o BlobNfsTargetResponseOutput) ToBlobNfsTargetResponseOutputWithContext(ctx context.Context) BlobNfsTargetResponseOutput {
	return o
}

func (o BlobNfsTargetResponseOutput) ToBlobNfsTargetResponsePtrOutput() BlobNfsTargetResponsePtrOutput {
	return o.ToBlobNfsTargetResponsePtrOutputWithContext(context.Background())
}

func (o BlobNfsTargetResponseOutput) ToBlobNfsTargetResponsePtrOutputWithContext(ctx context.Context) BlobNfsTargetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobNfsTargetResponse) *BlobNfsTargetResponse {
		return &v
	}).(BlobNfsTargetResponsePtrOutput)
}

func (o BlobNfsTargetResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobNfsTargetResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o BlobNfsTargetResponseOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobNfsTargetResponse) *string { return v.UsageModel }).(pulumi.StringPtrOutput)
}

type BlobNfsTargetResponsePtrOutput struct{ *pulumi.OutputState }

func (BlobNfsTargetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobNfsTargetResponse)(nil)).Elem()
}

func (o BlobNfsTargetResponsePtrOutput) ToBlobNfsTargetResponsePtrOutput() BlobNfsTargetResponsePtrOutput {
	return o
}

func (o BlobNfsTargetResponsePtrOutput) ToBlobNfsTargetResponsePtrOutputWithContext(ctx context.Context) BlobNfsTargetResponsePtrOutput {
	return o
}

func (o BlobNfsTargetResponsePtrOutput) Elem() BlobNfsTargetResponseOutput {
	return o.ApplyT(func(v *BlobNfsTargetResponse) BlobNfsTargetResponse {
		if v != nil {
			return *v
		}
		var ret BlobNfsTargetResponse
		return ret
	}).(BlobNfsTargetResponseOutput)
}

func (o BlobNfsTargetResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobNfsTargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

func (o BlobNfsTargetResponsePtrOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobNfsTargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.UsageModel
	}).(pulumi.StringPtrOutput)
}

type CacheActiveDirectorySettings struct {
	CacheNetBiosName      string                                   `pulumi:"cacheNetBiosName"`
	Credentials           *CacheActiveDirectorySettingsCredentials `pulumi:"credentials"`
	DomainName            string                                   `pulumi:"domainName"`
	DomainNetBiosName     string                                   `pulumi:"domainNetBiosName"`
	PrimaryDnsIpAddress   string                                   `pulumi:"primaryDnsIpAddress"`
	SecondaryDnsIpAddress *string                                  `pulumi:"secondaryDnsIpAddress"`
}





type CacheActiveDirectorySettingsInput interface {
	pulumi.Input

	ToCacheActiveDirectorySettingsOutput() CacheActiveDirectorySettingsOutput
	ToCacheActiveDirectorySettingsOutputWithContext(context.Context) CacheActiveDirectorySettingsOutput
}

type CacheActiveDirectorySettingsArgs struct {
	CacheNetBiosName      pulumi.StringInput                              `pulumi:"cacheNetBiosName"`
	Credentials           CacheActiveDirectorySettingsCredentialsPtrInput `pulumi:"credentials"`
	DomainName            pulumi.StringInput                              `pulumi:"domainName"`
	DomainNetBiosName     pulumi.StringInput                              `pulumi:"domainNetBiosName"`
	PrimaryDnsIpAddress   pulumi.StringInput                              `pulumi:"primaryDnsIpAddress"`
	SecondaryDnsIpAddress pulumi.StringPtrInput                           `pulumi:"secondaryDnsIpAddress"`
}

func (CacheActiveDirectorySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheActiveDirectorySettings)(nil)).Elem()
}

func (i CacheActiveDirectorySettingsArgs) ToCacheActiveDirectorySettingsOutput() CacheActiveDirectorySettingsOutput {
	return i.ToCacheActiveDirectorySettingsOutputWithContext(context.Background())
}

func (i CacheActiveDirectorySettingsArgs) ToCacheActiveDirectorySettingsOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsOutput)
}

func (i CacheActiveDirectorySettingsArgs) ToCacheActiveDirectorySettingsPtrOutput() CacheActiveDirectorySettingsPtrOutput {
	return i.ToCacheActiveDirectorySettingsPtrOutputWithContext(context.Background())
}

func (i CacheActiveDirectorySettingsArgs) ToCacheActiveDirectorySettingsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsOutput).ToCacheActiveDirectorySettingsPtrOutputWithContext(ctx)
}









type CacheActiveDirectorySettingsPtrInput interface {
	pulumi.Input

	ToCacheActiveDirectorySettingsPtrOutput() CacheActiveDirectorySettingsPtrOutput
	ToCacheActiveDirectorySettingsPtrOutputWithContext(context.Context) CacheActiveDirectorySettingsPtrOutput
}

type cacheActiveDirectorySettingsPtrType CacheActiveDirectorySettingsArgs

func CacheActiveDirectorySettingsPtr(v *CacheActiveDirectorySettingsArgs) CacheActiveDirectorySettingsPtrInput {
	return (*cacheActiveDirectorySettingsPtrType)(v)
}

func (*cacheActiveDirectorySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheActiveDirectorySettings)(nil)).Elem()
}

func (i *cacheActiveDirectorySettingsPtrType) ToCacheActiveDirectorySettingsPtrOutput() CacheActiveDirectorySettingsPtrOutput {
	return i.ToCacheActiveDirectorySettingsPtrOutputWithContext(context.Background())
}

func (i *cacheActiveDirectorySettingsPtrType) ToCacheActiveDirectorySettingsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsPtrOutput)
}

type CacheActiveDirectorySettingsOutput struct{ *pulumi.OutputState }

func (CacheActiveDirectorySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheActiveDirectorySettings)(nil)).Elem()
}

func (o CacheActiveDirectorySettingsOutput) ToCacheActiveDirectorySettingsOutput() CacheActiveDirectorySettingsOutput {
	return o
}

func (o CacheActiveDirectorySettingsOutput) ToCacheActiveDirectorySettingsOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsOutput {
	return o
}

func (o CacheActiveDirectorySettingsOutput) ToCacheActiveDirectorySettingsPtrOutput() CacheActiveDirectorySettingsPtrOutput {
	return o.ToCacheActiveDirectorySettingsPtrOutputWithContext(context.Background())
}

func (o CacheActiveDirectorySettingsOutput) ToCacheActiveDirectorySettingsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheActiveDirectorySettings) *CacheActiveDirectorySettings {
		return &v
	}).(CacheActiveDirectorySettingsPtrOutput)
}

func (o CacheActiveDirectorySettingsOutput) CacheNetBiosName() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettings) string { return v.CacheNetBiosName }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsOutput) Credentials() CacheActiveDirectorySettingsCredentialsPtrOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettings) *CacheActiveDirectorySettingsCredentials { return v.Credentials }).(CacheActiveDirectorySettingsCredentialsPtrOutput)
}

func (o CacheActiveDirectorySettingsOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettings) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsOutput) DomainNetBiosName() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettings) string { return v.DomainNetBiosName }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsOutput) PrimaryDnsIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettings) string { return v.PrimaryDnsIpAddress }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsOutput) SecondaryDnsIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettings) *string { return v.SecondaryDnsIpAddress }).(pulumi.StringPtrOutput)
}

type CacheActiveDirectorySettingsPtrOutput struct{ *pulumi.OutputState }

func (CacheActiveDirectorySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheActiveDirectorySettings)(nil)).Elem()
}

func (o CacheActiveDirectorySettingsPtrOutput) ToCacheActiveDirectorySettingsPtrOutput() CacheActiveDirectorySettingsPtrOutput {
	return o
}

func (o CacheActiveDirectorySettingsPtrOutput) ToCacheActiveDirectorySettingsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsPtrOutput {
	return o
}

func (o CacheActiveDirectorySettingsPtrOutput) Elem() CacheActiveDirectorySettingsOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettings) CacheActiveDirectorySettings {
		if v != nil {
			return *v
		}
		var ret CacheActiveDirectorySettings
		return ret
	}).(CacheActiveDirectorySettingsOutput)
}

func (o CacheActiveDirectorySettingsPtrOutput) CacheNetBiosName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettings) *string {
		if v == nil {
			return nil
		}
		return &v.CacheNetBiosName
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsPtrOutput) Credentials() CacheActiveDirectorySettingsCredentialsPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettings) *CacheActiveDirectorySettingsCredentials {
		if v == nil {
			return nil
		}
		return v.Credentials
	}).(CacheActiveDirectorySettingsCredentialsPtrOutput)
}

func (o CacheActiveDirectorySettingsPtrOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettings) *string {
		if v == nil {
			return nil
		}
		return &v.DomainName
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsPtrOutput) DomainNetBiosName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettings) *string {
		if v == nil {
			return nil
		}
		return &v.DomainNetBiosName
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsPtrOutput) PrimaryDnsIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettings) *string {
		if v == nil {
			return nil
		}
		return &v.PrimaryDnsIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsPtrOutput) SecondaryDnsIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettings) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryDnsIpAddress
	}).(pulumi.StringPtrOutput)
}

type CacheActiveDirectorySettingsCredentials struct {
	Password string `pulumi:"password"`
	Username string `pulumi:"username"`
}





type CacheActiveDirectorySettingsCredentialsInput interface {
	pulumi.Input

	ToCacheActiveDirectorySettingsCredentialsOutput() CacheActiveDirectorySettingsCredentialsOutput
	ToCacheActiveDirectorySettingsCredentialsOutputWithContext(context.Context) CacheActiveDirectorySettingsCredentialsOutput
}

type CacheActiveDirectorySettingsCredentialsArgs struct {
	Password pulumi.StringInput `pulumi:"password"`
	Username pulumi.StringInput `pulumi:"username"`
}

func (CacheActiveDirectorySettingsCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheActiveDirectorySettingsCredentials)(nil)).Elem()
}

func (i CacheActiveDirectorySettingsCredentialsArgs) ToCacheActiveDirectorySettingsCredentialsOutput() CacheActiveDirectorySettingsCredentialsOutput {
	return i.ToCacheActiveDirectorySettingsCredentialsOutputWithContext(context.Background())
}

func (i CacheActiveDirectorySettingsCredentialsArgs) ToCacheActiveDirectorySettingsCredentialsOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsCredentialsOutput)
}

func (i CacheActiveDirectorySettingsCredentialsArgs) ToCacheActiveDirectorySettingsCredentialsPtrOutput() CacheActiveDirectorySettingsCredentialsPtrOutput {
	return i.ToCacheActiveDirectorySettingsCredentialsPtrOutputWithContext(context.Background())
}

func (i CacheActiveDirectorySettingsCredentialsArgs) ToCacheActiveDirectorySettingsCredentialsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsCredentialsOutput).ToCacheActiveDirectorySettingsCredentialsPtrOutputWithContext(ctx)
}









type CacheActiveDirectorySettingsCredentialsPtrInput interface {
	pulumi.Input

	ToCacheActiveDirectorySettingsCredentialsPtrOutput() CacheActiveDirectorySettingsCredentialsPtrOutput
	ToCacheActiveDirectorySettingsCredentialsPtrOutputWithContext(context.Context) CacheActiveDirectorySettingsCredentialsPtrOutput
}

type cacheActiveDirectorySettingsCredentialsPtrType CacheActiveDirectorySettingsCredentialsArgs

func CacheActiveDirectorySettingsCredentialsPtr(v *CacheActiveDirectorySettingsCredentialsArgs) CacheActiveDirectorySettingsCredentialsPtrInput {
	return (*cacheActiveDirectorySettingsCredentialsPtrType)(v)
}

func (*cacheActiveDirectorySettingsCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheActiveDirectorySettingsCredentials)(nil)).Elem()
}

func (i *cacheActiveDirectorySettingsCredentialsPtrType) ToCacheActiveDirectorySettingsCredentialsPtrOutput() CacheActiveDirectorySettingsCredentialsPtrOutput {
	return i.ToCacheActiveDirectorySettingsCredentialsPtrOutputWithContext(context.Background())
}

func (i *cacheActiveDirectorySettingsCredentialsPtrType) ToCacheActiveDirectorySettingsCredentialsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsCredentialsPtrOutput)
}

type CacheActiveDirectorySettingsCredentialsOutput struct{ *pulumi.OutputState }

func (CacheActiveDirectorySettingsCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheActiveDirectorySettingsCredentials)(nil)).Elem()
}

func (o CacheActiveDirectorySettingsCredentialsOutput) ToCacheActiveDirectorySettingsCredentialsOutput() CacheActiveDirectorySettingsCredentialsOutput {
	return o
}

func (o CacheActiveDirectorySettingsCredentialsOutput) ToCacheActiveDirectorySettingsCredentialsOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsCredentialsOutput {
	return o
}

func (o CacheActiveDirectorySettingsCredentialsOutput) ToCacheActiveDirectorySettingsCredentialsPtrOutput() CacheActiveDirectorySettingsCredentialsPtrOutput {
	return o.ToCacheActiveDirectorySettingsCredentialsPtrOutputWithContext(context.Background())
}

func (o CacheActiveDirectorySettingsCredentialsOutput) ToCacheActiveDirectorySettingsCredentialsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheActiveDirectorySettingsCredentials) *CacheActiveDirectorySettingsCredentials {
		return &v
	}).(CacheActiveDirectorySettingsCredentialsPtrOutput)
}

func (o CacheActiveDirectorySettingsCredentialsOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsCredentials) string { return v.Password }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsCredentialsOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsCredentials) string { return v.Username }).(pulumi.StringOutput)
}

type CacheActiveDirectorySettingsCredentialsPtrOutput struct{ *pulumi.OutputState }

func (CacheActiveDirectorySettingsCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheActiveDirectorySettingsCredentials)(nil)).Elem()
}

func (o CacheActiveDirectorySettingsCredentialsPtrOutput) ToCacheActiveDirectorySettingsCredentialsPtrOutput() CacheActiveDirectorySettingsCredentialsPtrOutput {
	return o
}

func (o CacheActiveDirectorySettingsCredentialsPtrOutput) ToCacheActiveDirectorySettingsCredentialsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsCredentialsPtrOutput {
	return o
}

func (o CacheActiveDirectorySettingsCredentialsPtrOutput) Elem() CacheActiveDirectorySettingsCredentialsOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsCredentials) CacheActiveDirectorySettingsCredentials {
		if v != nil {
			return *v
		}
		var ret CacheActiveDirectorySettingsCredentials
		return ret
	}).(CacheActiveDirectorySettingsCredentialsOutput)
}

func (o CacheActiveDirectorySettingsCredentialsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsCredentialsPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type CacheActiveDirectorySettingsResponse struct {
	CacheNetBiosName      string                                           `pulumi:"cacheNetBiosName"`
	Credentials           *CacheActiveDirectorySettingsResponseCredentials `pulumi:"credentials"`
	DomainJoined          string                                           `pulumi:"domainJoined"`
	DomainName            string                                           `pulumi:"domainName"`
	DomainNetBiosName     string                                           `pulumi:"domainNetBiosName"`
	PrimaryDnsIpAddress   string                                           `pulumi:"primaryDnsIpAddress"`
	SecondaryDnsIpAddress *string                                          `pulumi:"secondaryDnsIpAddress"`
}





type CacheActiveDirectorySettingsResponseInput interface {
	pulumi.Input

	ToCacheActiveDirectorySettingsResponseOutput() CacheActiveDirectorySettingsResponseOutput
	ToCacheActiveDirectorySettingsResponseOutputWithContext(context.Context) CacheActiveDirectorySettingsResponseOutput
}

type CacheActiveDirectorySettingsResponseArgs struct {
	CacheNetBiosName      pulumi.StringInput                                      `pulumi:"cacheNetBiosName"`
	Credentials           CacheActiveDirectorySettingsResponseCredentialsPtrInput `pulumi:"credentials"`
	DomainJoined          pulumi.StringInput                                      `pulumi:"domainJoined"`
	DomainName            pulumi.StringInput                                      `pulumi:"domainName"`
	DomainNetBiosName     pulumi.StringInput                                      `pulumi:"domainNetBiosName"`
	PrimaryDnsIpAddress   pulumi.StringInput                                      `pulumi:"primaryDnsIpAddress"`
	SecondaryDnsIpAddress pulumi.StringPtrInput                                   `pulumi:"secondaryDnsIpAddress"`
}

func (CacheActiveDirectorySettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheActiveDirectorySettingsResponse)(nil)).Elem()
}

func (i CacheActiveDirectorySettingsResponseArgs) ToCacheActiveDirectorySettingsResponseOutput() CacheActiveDirectorySettingsResponseOutput {
	return i.ToCacheActiveDirectorySettingsResponseOutputWithContext(context.Background())
}

func (i CacheActiveDirectorySettingsResponseArgs) ToCacheActiveDirectorySettingsResponseOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsResponseOutput)
}

func (i CacheActiveDirectorySettingsResponseArgs) ToCacheActiveDirectorySettingsResponsePtrOutput() CacheActiveDirectorySettingsResponsePtrOutput {
	return i.ToCacheActiveDirectorySettingsResponsePtrOutputWithContext(context.Background())
}

func (i CacheActiveDirectorySettingsResponseArgs) ToCacheActiveDirectorySettingsResponsePtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsResponseOutput).ToCacheActiveDirectorySettingsResponsePtrOutputWithContext(ctx)
}









type CacheActiveDirectorySettingsResponsePtrInput interface {
	pulumi.Input

	ToCacheActiveDirectorySettingsResponsePtrOutput() CacheActiveDirectorySettingsResponsePtrOutput
	ToCacheActiveDirectorySettingsResponsePtrOutputWithContext(context.Context) CacheActiveDirectorySettingsResponsePtrOutput
}

type cacheActiveDirectorySettingsResponsePtrType CacheActiveDirectorySettingsResponseArgs

func CacheActiveDirectorySettingsResponsePtr(v *CacheActiveDirectorySettingsResponseArgs) CacheActiveDirectorySettingsResponsePtrInput {
	return (*cacheActiveDirectorySettingsResponsePtrType)(v)
}

func (*cacheActiveDirectorySettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheActiveDirectorySettingsResponse)(nil)).Elem()
}

func (i *cacheActiveDirectorySettingsResponsePtrType) ToCacheActiveDirectorySettingsResponsePtrOutput() CacheActiveDirectorySettingsResponsePtrOutput {
	return i.ToCacheActiveDirectorySettingsResponsePtrOutputWithContext(context.Background())
}

func (i *cacheActiveDirectorySettingsResponsePtrType) ToCacheActiveDirectorySettingsResponsePtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsResponsePtrOutput)
}

type CacheActiveDirectorySettingsResponseOutput struct{ *pulumi.OutputState }

func (CacheActiveDirectorySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheActiveDirectorySettingsResponse)(nil)).Elem()
}

func (o CacheActiveDirectorySettingsResponseOutput) ToCacheActiveDirectorySettingsResponseOutput() CacheActiveDirectorySettingsResponseOutput {
	return o
}

func (o CacheActiveDirectorySettingsResponseOutput) ToCacheActiveDirectorySettingsResponseOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponseOutput {
	return o
}

func (o CacheActiveDirectorySettingsResponseOutput) ToCacheActiveDirectorySettingsResponsePtrOutput() CacheActiveDirectorySettingsResponsePtrOutput {
	return o.ToCacheActiveDirectorySettingsResponsePtrOutputWithContext(context.Background())
}

func (o CacheActiveDirectorySettingsResponseOutput) ToCacheActiveDirectorySettingsResponsePtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheActiveDirectorySettingsResponse) *CacheActiveDirectorySettingsResponse {
		return &v
	}).(CacheActiveDirectorySettingsResponsePtrOutput)
}

func (o CacheActiveDirectorySettingsResponseOutput) CacheNetBiosName() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsResponse) string { return v.CacheNetBiosName }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsResponseOutput) Credentials() CacheActiveDirectorySettingsResponseCredentialsPtrOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsResponse) *CacheActiveDirectorySettingsResponseCredentials {
		return v.Credentials
	}).(CacheActiveDirectorySettingsResponseCredentialsPtrOutput)
}

func (o CacheActiveDirectorySettingsResponseOutput) DomainJoined() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsResponse) string { return v.DomainJoined }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsResponseOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsResponse) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsResponseOutput) DomainNetBiosName() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsResponse) string { return v.DomainNetBiosName }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsResponseOutput) PrimaryDnsIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsResponse) string { return v.PrimaryDnsIpAddress }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsResponseOutput) SecondaryDnsIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsResponse) *string { return v.SecondaryDnsIpAddress }).(pulumi.StringPtrOutput)
}

type CacheActiveDirectorySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheActiveDirectorySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheActiveDirectorySettingsResponse)(nil)).Elem()
}

func (o CacheActiveDirectorySettingsResponsePtrOutput) ToCacheActiveDirectorySettingsResponsePtrOutput() CacheActiveDirectorySettingsResponsePtrOutput {
	return o
}

func (o CacheActiveDirectorySettingsResponsePtrOutput) ToCacheActiveDirectorySettingsResponsePtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponsePtrOutput {
	return o
}

func (o CacheActiveDirectorySettingsResponsePtrOutput) Elem() CacheActiveDirectorySettingsResponseOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponse) CacheActiveDirectorySettingsResponse {
		if v != nil {
			return *v
		}
		var ret CacheActiveDirectorySettingsResponse
		return ret
	}).(CacheActiveDirectorySettingsResponseOutput)
}

func (o CacheActiveDirectorySettingsResponsePtrOutput) CacheNetBiosName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CacheNetBiosName
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsResponsePtrOutput) Credentials() CacheActiveDirectorySettingsResponseCredentialsPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponse) *CacheActiveDirectorySettingsResponseCredentials {
		if v == nil {
			return nil
		}
		return v.Credentials
	}).(CacheActiveDirectorySettingsResponseCredentialsPtrOutput)
}

func (o CacheActiveDirectorySettingsResponsePtrOutput) DomainJoined() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DomainJoined
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsResponsePtrOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DomainName
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsResponsePtrOutput) DomainNetBiosName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DomainNetBiosName
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsResponsePtrOutput) PrimaryDnsIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrimaryDnsIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsResponsePtrOutput) SecondaryDnsIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryDnsIpAddress
	}).(pulumi.StringPtrOutput)
}

type CacheActiveDirectorySettingsResponseCredentials struct {
	Password string `pulumi:"password"`
	Username string `pulumi:"username"`
}





type CacheActiveDirectorySettingsResponseCredentialsInput interface {
	pulumi.Input

	ToCacheActiveDirectorySettingsResponseCredentialsOutput() CacheActiveDirectorySettingsResponseCredentialsOutput
	ToCacheActiveDirectorySettingsResponseCredentialsOutputWithContext(context.Context) CacheActiveDirectorySettingsResponseCredentialsOutput
}

type CacheActiveDirectorySettingsResponseCredentialsArgs struct {
	Password pulumi.StringInput `pulumi:"password"`
	Username pulumi.StringInput `pulumi:"username"`
}

func (CacheActiveDirectorySettingsResponseCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheActiveDirectorySettingsResponseCredentials)(nil)).Elem()
}

func (i CacheActiveDirectorySettingsResponseCredentialsArgs) ToCacheActiveDirectorySettingsResponseCredentialsOutput() CacheActiveDirectorySettingsResponseCredentialsOutput {
	return i.ToCacheActiveDirectorySettingsResponseCredentialsOutputWithContext(context.Background())
}

func (i CacheActiveDirectorySettingsResponseCredentialsArgs) ToCacheActiveDirectorySettingsResponseCredentialsOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponseCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsResponseCredentialsOutput)
}

func (i CacheActiveDirectorySettingsResponseCredentialsArgs) ToCacheActiveDirectorySettingsResponseCredentialsPtrOutput() CacheActiveDirectorySettingsResponseCredentialsPtrOutput {
	return i.ToCacheActiveDirectorySettingsResponseCredentialsPtrOutputWithContext(context.Background())
}

func (i CacheActiveDirectorySettingsResponseCredentialsArgs) ToCacheActiveDirectorySettingsResponseCredentialsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponseCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsResponseCredentialsOutput).ToCacheActiveDirectorySettingsResponseCredentialsPtrOutputWithContext(ctx)
}









type CacheActiveDirectorySettingsResponseCredentialsPtrInput interface {
	pulumi.Input

	ToCacheActiveDirectorySettingsResponseCredentialsPtrOutput() CacheActiveDirectorySettingsResponseCredentialsPtrOutput
	ToCacheActiveDirectorySettingsResponseCredentialsPtrOutputWithContext(context.Context) CacheActiveDirectorySettingsResponseCredentialsPtrOutput
}

type cacheActiveDirectorySettingsResponseCredentialsPtrType CacheActiveDirectorySettingsResponseCredentialsArgs

func CacheActiveDirectorySettingsResponseCredentialsPtr(v *CacheActiveDirectorySettingsResponseCredentialsArgs) CacheActiveDirectorySettingsResponseCredentialsPtrInput {
	return (*cacheActiveDirectorySettingsResponseCredentialsPtrType)(v)
}

func (*cacheActiveDirectorySettingsResponseCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheActiveDirectorySettingsResponseCredentials)(nil)).Elem()
}

func (i *cacheActiveDirectorySettingsResponseCredentialsPtrType) ToCacheActiveDirectorySettingsResponseCredentialsPtrOutput() CacheActiveDirectorySettingsResponseCredentialsPtrOutput {
	return i.ToCacheActiveDirectorySettingsResponseCredentialsPtrOutputWithContext(context.Background())
}

func (i *cacheActiveDirectorySettingsResponseCredentialsPtrType) ToCacheActiveDirectorySettingsResponseCredentialsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponseCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheActiveDirectorySettingsResponseCredentialsPtrOutput)
}

type CacheActiveDirectorySettingsResponseCredentialsOutput struct{ *pulumi.OutputState }

func (CacheActiveDirectorySettingsResponseCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheActiveDirectorySettingsResponseCredentials)(nil)).Elem()
}

func (o CacheActiveDirectorySettingsResponseCredentialsOutput) ToCacheActiveDirectorySettingsResponseCredentialsOutput() CacheActiveDirectorySettingsResponseCredentialsOutput {
	return o
}

func (o CacheActiveDirectorySettingsResponseCredentialsOutput) ToCacheActiveDirectorySettingsResponseCredentialsOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponseCredentialsOutput {
	return o
}

func (o CacheActiveDirectorySettingsResponseCredentialsOutput) ToCacheActiveDirectorySettingsResponseCredentialsPtrOutput() CacheActiveDirectorySettingsResponseCredentialsPtrOutput {
	return o.ToCacheActiveDirectorySettingsResponseCredentialsPtrOutputWithContext(context.Background())
}

func (o CacheActiveDirectorySettingsResponseCredentialsOutput) ToCacheActiveDirectorySettingsResponseCredentialsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponseCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheActiveDirectorySettingsResponseCredentials) *CacheActiveDirectorySettingsResponseCredentials {
		return &v
	}).(CacheActiveDirectorySettingsResponseCredentialsPtrOutput)
}

func (o CacheActiveDirectorySettingsResponseCredentialsOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsResponseCredentials) string { return v.Password }).(pulumi.StringOutput)
}

func (o CacheActiveDirectorySettingsResponseCredentialsOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v CacheActiveDirectorySettingsResponseCredentials) string { return v.Username }).(pulumi.StringOutput)
}

type CacheActiveDirectorySettingsResponseCredentialsPtrOutput struct{ *pulumi.OutputState }

func (CacheActiveDirectorySettingsResponseCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheActiveDirectorySettingsResponseCredentials)(nil)).Elem()
}

func (o CacheActiveDirectorySettingsResponseCredentialsPtrOutput) ToCacheActiveDirectorySettingsResponseCredentialsPtrOutput() CacheActiveDirectorySettingsResponseCredentialsPtrOutput {
	return o
}

func (o CacheActiveDirectorySettingsResponseCredentialsPtrOutput) ToCacheActiveDirectorySettingsResponseCredentialsPtrOutputWithContext(ctx context.Context) CacheActiveDirectorySettingsResponseCredentialsPtrOutput {
	return o
}

func (o CacheActiveDirectorySettingsResponseCredentialsPtrOutput) Elem() CacheActiveDirectorySettingsResponseCredentialsOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponseCredentials) CacheActiveDirectorySettingsResponseCredentials {
		if v != nil {
			return *v
		}
		var ret CacheActiveDirectorySettingsResponseCredentials
		return ret
	}).(CacheActiveDirectorySettingsResponseCredentialsOutput)
}

func (o CacheActiveDirectorySettingsResponseCredentialsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponseCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o CacheActiveDirectorySettingsResponseCredentialsPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheActiveDirectorySettingsResponseCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type CacheDirectorySettings struct {
	ActiveDirectory  *CacheActiveDirectorySettings  `pulumi:"activeDirectory"`
	UsernameDownload *CacheUsernameDownloadSettings `pulumi:"usernameDownload"`
}





type CacheDirectorySettingsInput interface {
	pulumi.Input

	ToCacheDirectorySettingsOutput() CacheDirectorySettingsOutput
	ToCacheDirectorySettingsOutputWithContext(context.Context) CacheDirectorySettingsOutput
}

type CacheDirectorySettingsArgs struct {
	ActiveDirectory  CacheActiveDirectorySettingsPtrInput  `pulumi:"activeDirectory"`
	UsernameDownload CacheUsernameDownloadSettingsPtrInput `pulumi:"usernameDownload"`
}

func (CacheDirectorySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheDirectorySettings)(nil)).Elem()
}

func (i CacheDirectorySettingsArgs) ToCacheDirectorySettingsOutput() CacheDirectorySettingsOutput {
	return i.ToCacheDirectorySettingsOutputWithContext(context.Background())
}

func (i CacheDirectorySettingsArgs) ToCacheDirectorySettingsOutputWithContext(ctx context.Context) CacheDirectorySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheDirectorySettingsOutput)
}

func (i CacheDirectorySettingsArgs) ToCacheDirectorySettingsPtrOutput() CacheDirectorySettingsPtrOutput {
	return i.ToCacheDirectorySettingsPtrOutputWithContext(context.Background())
}

func (i CacheDirectorySettingsArgs) ToCacheDirectorySettingsPtrOutputWithContext(ctx context.Context) CacheDirectorySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheDirectorySettingsOutput).ToCacheDirectorySettingsPtrOutputWithContext(ctx)
}









type CacheDirectorySettingsPtrInput interface {
	pulumi.Input

	ToCacheDirectorySettingsPtrOutput() CacheDirectorySettingsPtrOutput
	ToCacheDirectorySettingsPtrOutputWithContext(context.Context) CacheDirectorySettingsPtrOutput
}

type cacheDirectorySettingsPtrType CacheDirectorySettingsArgs

func CacheDirectorySettingsPtr(v *CacheDirectorySettingsArgs) CacheDirectorySettingsPtrInput {
	return (*cacheDirectorySettingsPtrType)(v)
}

func (*cacheDirectorySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheDirectorySettings)(nil)).Elem()
}

func (i *cacheDirectorySettingsPtrType) ToCacheDirectorySettingsPtrOutput() CacheDirectorySettingsPtrOutput {
	return i.ToCacheDirectorySettingsPtrOutputWithContext(context.Background())
}

func (i *cacheDirectorySettingsPtrType) ToCacheDirectorySettingsPtrOutputWithContext(ctx context.Context) CacheDirectorySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheDirectorySettingsPtrOutput)
}

type CacheDirectorySettingsOutput struct{ *pulumi.OutputState }

func (CacheDirectorySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheDirectorySettings)(nil)).Elem()
}

func (o CacheDirectorySettingsOutput) ToCacheDirectorySettingsOutput() CacheDirectorySettingsOutput {
	return o
}

func (o CacheDirectorySettingsOutput) ToCacheDirectorySettingsOutputWithContext(ctx context.Context) CacheDirectorySettingsOutput {
	return o
}

func (o CacheDirectorySettingsOutput) ToCacheDirectorySettingsPtrOutput() CacheDirectorySettingsPtrOutput {
	return o.ToCacheDirectorySettingsPtrOutputWithContext(context.Background())
}

func (o CacheDirectorySettingsOutput) ToCacheDirectorySettingsPtrOutputWithContext(ctx context.Context) CacheDirectorySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheDirectorySettings) *CacheDirectorySettings {
		return &v
	}).(CacheDirectorySettingsPtrOutput)
}

func (o CacheDirectorySettingsOutput) ActiveDirectory() CacheActiveDirectorySettingsPtrOutput {
	return o.ApplyT(func(v CacheDirectorySettings) *CacheActiveDirectorySettings { return v.ActiveDirectory }).(CacheActiveDirectorySettingsPtrOutput)
}

func (o CacheDirectorySettingsOutput) UsernameDownload() CacheUsernameDownloadSettingsPtrOutput {
	return o.ApplyT(func(v CacheDirectorySettings) *CacheUsernameDownloadSettings { return v.UsernameDownload }).(CacheUsernameDownloadSettingsPtrOutput)
}

type CacheDirectorySettingsPtrOutput struct{ *pulumi.OutputState }

func (CacheDirectorySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheDirectorySettings)(nil)).Elem()
}

func (o CacheDirectorySettingsPtrOutput) ToCacheDirectorySettingsPtrOutput() CacheDirectorySettingsPtrOutput {
	return o
}

func (o CacheDirectorySettingsPtrOutput) ToCacheDirectorySettingsPtrOutputWithContext(ctx context.Context) CacheDirectorySettingsPtrOutput {
	return o
}

func (o CacheDirectorySettingsPtrOutput) Elem() CacheDirectorySettingsOutput {
	return o.ApplyT(func(v *CacheDirectorySettings) CacheDirectorySettings {
		if v != nil {
			return *v
		}
		var ret CacheDirectorySettings
		return ret
	}).(CacheDirectorySettingsOutput)
}

func (o CacheDirectorySettingsPtrOutput) ActiveDirectory() CacheActiveDirectorySettingsPtrOutput {
	return o.ApplyT(func(v *CacheDirectorySettings) *CacheActiveDirectorySettings {
		if v == nil {
			return nil
		}
		return v.ActiveDirectory
	}).(CacheActiveDirectorySettingsPtrOutput)
}

func (o CacheDirectorySettingsPtrOutput) UsernameDownload() CacheUsernameDownloadSettingsPtrOutput {
	return o.ApplyT(func(v *CacheDirectorySettings) *CacheUsernameDownloadSettings {
		if v == nil {
			return nil
		}
		return v.UsernameDownload
	}).(CacheUsernameDownloadSettingsPtrOutput)
}

type CacheDirectorySettingsResponse struct {
	ActiveDirectory  *CacheActiveDirectorySettingsResponse  `pulumi:"activeDirectory"`
	UsernameDownload *CacheUsernameDownloadSettingsResponse `pulumi:"usernameDownload"`
}





type CacheDirectorySettingsResponseInput interface {
	pulumi.Input

	ToCacheDirectorySettingsResponseOutput() CacheDirectorySettingsResponseOutput
	ToCacheDirectorySettingsResponseOutputWithContext(context.Context) CacheDirectorySettingsResponseOutput
}

type CacheDirectorySettingsResponseArgs struct {
	ActiveDirectory  CacheActiveDirectorySettingsResponsePtrInput  `pulumi:"activeDirectory"`
	UsernameDownload CacheUsernameDownloadSettingsResponsePtrInput `pulumi:"usernameDownload"`
}

func (CacheDirectorySettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheDirectorySettingsResponse)(nil)).Elem()
}

func (i CacheDirectorySettingsResponseArgs) ToCacheDirectorySettingsResponseOutput() CacheDirectorySettingsResponseOutput {
	return i.ToCacheDirectorySettingsResponseOutputWithContext(context.Background())
}

func (i CacheDirectorySettingsResponseArgs) ToCacheDirectorySettingsResponseOutputWithContext(ctx context.Context) CacheDirectorySettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheDirectorySettingsResponseOutput)
}

func (i CacheDirectorySettingsResponseArgs) ToCacheDirectorySettingsResponsePtrOutput() CacheDirectorySettingsResponsePtrOutput {
	return i.ToCacheDirectorySettingsResponsePtrOutputWithContext(context.Background())
}

func (i CacheDirectorySettingsResponseArgs) ToCacheDirectorySettingsResponsePtrOutputWithContext(ctx context.Context) CacheDirectorySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheDirectorySettingsResponseOutput).ToCacheDirectorySettingsResponsePtrOutputWithContext(ctx)
}









type CacheDirectorySettingsResponsePtrInput interface {
	pulumi.Input

	ToCacheDirectorySettingsResponsePtrOutput() CacheDirectorySettingsResponsePtrOutput
	ToCacheDirectorySettingsResponsePtrOutputWithContext(context.Context) CacheDirectorySettingsResponsePtrOutput
}

type cacheDirectorySettingsResponsePtrType CacheDirectorySettingsResponseArgs

func CacheDirectorySettingsResponsePtr(v *CacheDirectorySettingsResponseArgs) CacheDirectorySettingsResponsePtrInput {
	return (*cacheDirectorySettingsResponsePtrType)(v)
}

func (*cacheDirectorySettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheDirectorySettingsResponse)(nil)).Elem()
}

func (i *cacheDirectorySettingsResponsePtrType) ToCacheDirectorySettingsResponsePtrOutput() CacheDirectorySettingsResponsePtrOutput {
	return i.ToCacheDirectorySettingsResponsePtrOutputWithContext(context.Background())
}

func (i *cacheDirectorySettingsResponsePtrType) ToCacheDirectorySettingsResponsePtrOutputWithContext(ctx context.Context) CacheDirectorySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheDirectorySettingsResponsePtrOutput)
}

type CacheDirectorySettingsResponseOutput struct{ *pulumi.OutputState }

func (CacheDirectorySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheDirectorySettingsResponse)(nil)).Elem()
}

func (o CacheDirectorySettingsResponseOutput) ToCacheDirectorySettingsResponseOutput() CacheDirectorySettingsResponseOutput {
	return o
}

func (o CacheDirectorySettingsResponseOutput) ToCacheDirectorySettingsResponseOutputWithContext(ctx context.Context) CacheDirectorySettingsResponseOutput {
	return o
}

func (o CacheDirectorySettingsResponseOutput) ToCacheDirectorySettingsResponsePtrOutput() CacheDirectorySettingsResponsePtrOutput {
	return o.ToCacheDirectorySettingsResponsePtrOutputWithContext(context.Background())
}

func (o CacheDirectorySettingsResponseOutput) ToCacheDirectorySettingsResponsePtrOutputWithContext(ctx context.Context) CacheDirectorySettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheDirectorySettingsResponse) *CacheDirectorySettingsResponse {
		return &v
	}).(CacheDirectorySettingsResponsePtrOutput)
}

func (o CacheDirectorySettingsResponseOutput) ActiveDirectory() CacheActiveDirectorySettingsResponsePtrOutput {
	return o.ApplyT(func(v CacheDirectorySettingsResponse) *CacheActiveDirectorySettingsResponse { return v.ActiveDirectory }).(CacheActiveDirectorySettingsResponsePtrOutput)
}

func (o CacheDirectorySettingsResponseOutput) UsernameDownload() CacheUsernameDownloadSettingsResponsePtrOutput {
	return o.ApplyT(func(v CacheDirectorySettingsResponse) *CacheUsernameDownloadSettingsResponse {
		return v.UsernameDownload
	}).(CacheUsernameDownloadSettingsResponsePtrOutput)
}

type CacheDirectorySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheDirectorySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheDirectorySettingsResponse)(nil)).Elem()
}

func (o CacheDirectorySettingsResponsePtrOutput) ToCacheDirectorySettingsResponsePtrOutput() CacheDirectorySettingsResponsePtrOutput {
	return o
}

func (o CacheDirectorySettingsResponsePtrOutput) ToCacheDirectorySettingsResponsePtrOutputWithContext(ctx context.Context) CacheDirectorySettingsResponsePtrOutput {
	return o
}

func (o CacheDirectorySettingsResponsePtrOutput) Elem() CacheDirectorySettingsResponseOutput {
	return o.ApplyT(func(v *CacheDirectorySettingsResponse) CacheDirectorySettingsResponse {
		if v != nil {
			return *v
		}
		var ret CacheDirectorySettingsResponse
		return ret
	}).(CacheDirectorySettingsResponseOutput)
}

func (o CacheDirectorySettingsResponsePtrOutput) ActiveDirectory() CacheActiveDirectorySettingsResponsePtrOutput {
	return o.ApplyT(func(v *CacheDirectorySettingsResponse) *CacheActiveDirectorySettingsResponse {
		if v == nil {
			return nil
		}
		return v.ActiveDirectory
	}).(CacheActiveDirectorySettingsResponsePtrOutput)
}

func (o CacheDirectorySettingsResponsePtrOutput) UsernameDownload() CacheUsernameDownloadSettingsResponsePtrOutput {
	return o.ApplyT(func(v *CacheDirectorySettingsResponse) *CacheUsernameDownloadSettingsResponse {
		if v == nil {
			return nil
		}
		return v.UsernameDownload
	}).(CacheUsernameDownloadSettingsResponsePtrOutput)
}

type CacheEncryptionSettings struct {
	KeyEncryptionKey *KeyVaultKeyReference `pulumi:"keyEncryptionKey"`
}





type CacheEncryptionSettingsInput interface {
	pulumi.Input

	ToCacheEncryptionSettingsOutput() CacheEncryptionSettingsOutput
	ToCacheEncryptionSettingsOutputWithContext(context.Context) CacheEncryptionSettingsOutput
}

type CacheEncryptionSettingsArgs struct {
	KeyEncryptionKey KeyVaultKeyReferencePtrInput `pulumi:"keyEncryptionKey"`
}

func (CacheEncryptionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheEncryptionSettings)(nil)).Elem()
}

func (i CacheEncryptionSettingsArgs) ToCacheEncryptionSettingsOutput() CacheEncryptionSettingsOutput {
	return i.ToCacheEncryptionSettingsOutputWithContext(context.Background())
}

func (i CacheEncryptionSettingsArgs) ToCacheEncryptionSettingsOutputWithContext(ctx context.Context) CacheEncryptionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheEncryptionSettingsOutput)
}

func (i CacheEncryptionSettingsArgs) ToCacheEncryptionSettingsPtrOutput() CacheEncryptionSettingsPtrOutput {
	return i.ToCacheEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (i CacheEncryptionSettingsArgs) ToCacheEncryptionSettingsPtrOutputWithContext(ctx context.Context) CacheEncryptionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheEncryptionSettingsOutput).ToCacheEncryptionSettingsPtrOutputWithContext(ctx)
}









type CacheEncryptionSettingsPtrInput interface {
	pulumi.Input

	ToCacheEncryptionSettingsPtrOutput() CacheEncryptionSettingsPtrOutput
	ToCacheEncryptionSettingsPtrOutputWithContext(context.Context) CacheEncryptionSettingsPtrOutput
}

type cacheEncryptionSettingsPtrType CacheEncryptionSettingsArgs

func CacheEncryptionSettingsPtr(v *CacheEncryptionSettingsArgs) CacheEncryptionSettingsPtrInput {
	return (*cacheEncryptionSettingsPtrType)(v)
}

func (*cacheEncryptionSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheEncryptionSettings)(nil)).Elem()
}

func (i *cacheEncryptionSettingsPtrType) ToCacheEncryptionSettingsPtrOutput() CacheEncryptionSettingsPtrOutput {
	return i.ToCacheEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (i *cacheEncryptionSettingsPtrType) ToCacheEncryptionSettingsPtrOutputWithContext(ctx context.Context) CacheEncryptionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheEncryptionSettingsPtrOutput)
}

type CacheEncryptionSettingsOutput struct{ *pulumi.OutputState }

func (CacheEncryptionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheEncryptionSettings)(nil)).Elem()
}

func (o CacheEncryptionSettingsOutput) ToCacheEncryptionSettingsOutput() CacheEncryptionSettingsOutput {
	return o
}

func (o CacheEncryptionSettingsOutput) ToCacheEncryptionSettingsOutputWithContext(ctx context.Context) CacheEncryptionSettingsOutput {
	return o
}

func (o CacheEncryptionSettingsOutput) ToCacheEncryptionSettingsPtrOutput() CacheEncryptionSettingsPtrOutput {
	return o.ToCacheEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (o CacheEncryptionSettingsOutput) ToCacheEncryptionSettingsPtrOutputWithContext(ctx context.Context) CacheEncryptionSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheEncryptionSettings) *CacheEncryptionSettings {
		return &v
	}).(CacheEncryptionSettingsPtrOutput)
}

func (o CacheEncryptionSettingsOutput) KeyEncryptionKey() KeyVaultKeyReferencePtrOutput {
	return o.ApplyT(func(v CacheEncryptionSettings) *KeyVaultKeyReference { return v.KeyEncryptionKey }).(KeyVaultKeyReferencePtrOutput)
}

type CacheEncryptionSettingsPtrOutput struct{ *pulumi.OutputState }

func (CacheEncryptionSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheEncryptionSettings)(nil)).Elem()
}

func (o CacheEncryptionSettingsPtrOutput) ToCacheEncryptionSettingsPtrOutput() CacheEncryptionSettingsPtrOutput {
	return o
}

func (o CacheEncryptionSettingsPtrOutput) ToCacheEncryptionSettingsPtrOutputWithContext(ctx context.Context) CacheEncryptionSettingsPtrOutput {
	return o
}

func (o CacheEncryptionSettingsPtrOutput) Elem() CacheEncryptionSettingsOutput {
	return o.ApplyT(func(v *CacheEncryptionSettings) CacheEncryptionSettings {
		if v != nil {
			return *v
		}
		var ret CacheEncryptionSettings
		return ret
	}).(CacheEncryptionSettingsOutput)
}

func (o CacheEncryptionSettingsPtrOutput) KeyEncryptionKey() KeyVaultKeyReferencePtrOutput {
	return o.ApplyT(func(v *CacheEncryptionSettings) *KeyVaultKeyReference {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKey
	}).(KeyVaultKeyReferencePtrOutput)
}

type CacheEncryptionSettingsResponse struct {
	KeyEncryptionKey *KeyVaultKeyReferenceResponse `pulumi:"keyEncryptionKey"`
}





type CacheEncryptionSettingsResponseInput interface {
	pulumi.Input

	ToCacheEncryptionSettingsResponseOutput() CacheEncryptionSettingsResponseOutput
	ToCacheEncryptionSettingsResponseOutputWithContext(context.Context) CacheEncryptionSettingsResponseOutput
}

type CacheEncryptionSettingsResponseArgs struct {
	KeyEncryptionKey KeyVaultKeyReferenceResponsePtrInput `pulumi:"keyEncryptionKey"`
}

func (CacheEncryptionSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheEncryptionSettingsResponse)(nil)).Elem()
}

func (i CacheEncryptionSettingsResponseArgs) ToCacheEncryptionSettingsResponseOutput() CacheEncryptionSettingsResponseOutput {
	return i.ToCacheEncryptionSettingsResponseOutputWithContext(context.Background())
}

func (i CacheEncryptionSettingsResponseArgs) ToCacheEncryptionSettingsResponseOutputWithContext(ctx context.Context) CacheEncryptionSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheEncryptionSettingsResponseOutput)
}

func (i CacheEncryptionSettingsResponseArgs) ToCacheEncryptionSettingsResponsePtrOutput() CacheEncryptionSettingsResponsePtrOutput {
	return i.ToCacheEncryptionSettingsResponsePtrOutputWithContext(context.Background())
}

func (i CacheEncryptionSettingsResponseArgs) ToCacheEncryptionSettingsResponsePtrOutputWithContext(ctx context.Context) CacheEncryptionSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheEncryptionSettingsResponseOutput).ToCacheEncryptionSettingsResponsePtrOutputWithContext(ctx)
}









type CacheEncryptionSettingsResponsePtrInput interface {
	pulumi.Input

	ToCacheEncryptionSettingsResponsePtrOutput() CacheEncryptionSettingsResponsePtrOutput
	ToCacheEncryptionSettingsResponsePtrOutputWithContext(context.Context) CacheEncryptionSettingsResponsePtrOutput
}

type cacheEncryptionSettingsResponsePtrType CacheEncryptionSettingsResponseArgs

func CacheEncryptionSettingsResponsePtr(v *CacheEncryptionSettingsResponseArgs) CacheEncryptionSettingsResponsePtrInput {
	return (*cacheEncryptionSettingsResponsePtrType)(v)
}

func (*cacheEncryptionSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheEncryptionSettingsResponse)(nil)).Elem()
}

func (i *cacheEncryptionSettingsResponsePtrType) ToCacheEncryptionSettingsResponsePtrOutput() CacheEncryptionSettingsResponsePtrOutput {
	return i.ToCacheEncryptionSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *cacheEncryptionSettingsResponsePtrType) ToCacheEncryptionSettingsResponsePtrOutputWithContext(ctx context.Context) CacheEncryptionSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheEncryptionSettingsResponsePtrOutput)
}

type CacheEncryptionSettingsResponseOutput struct{ *pulumi.OutputState }

func (CacheEncryptionSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheEncryptionSettingsResponse)(nil)).Elem()
}

func (o CacheEncryptionSettingsResponseOutput) ToCacheEncryptionSettingsResponseOutput() CacheEncryptionSettingsResponseOutput {
	return o
}

func (o CacheEncryptionSettingsResponseOutput) ToCacheEncryptionSettingsResponseOutputWithContext(ctx context.Context) CacheEncryptionSettingsResponseOutput {
	return o
}

func (o CacheEncryptionSettingsResponseOutput) ToCacheEncryptionSettingsResponsePtrOutput() CacheEncryptionSettingsResponsePtrOutput {
	return o.ToCacheEncryptionSettingsResponsePtrOutputWithContext(context.Background())
}

func (o CacheEncryptionSettingsResponseOutput) ToCacheEncryptionSettingsResponsePtrOutputWithContext(ctx context.Context) CacheEncryptionSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheEncryptionSettingsResponse) *CacheEncryptionSettingsResponse {
		return &v
	}).(CacheEncryptionSettingsResponsePtrOutput)
}

func (o CacheEncryptionSettingsResponseOutput) KeyEncryptionKey() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v CacheEncryptionSettingsResponse) *KeyVaultKeyReferenceResponse { return v.KeyEncryptionKey }).(KeyVaultKeyReferenceResponsePtrOutput)
}

type CacheEncryptionSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheEncryptionSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheEncryptionSettingsResponse)(nil)).Elem()
}

func (o CacheEncryptionSettingsResponsePtrOutput) ToCacheEncryptionSettingsResponsePtrOutput() CacheEncryptionSettingsResponsePtrOutput {
	return o
}

func (o CacheEncryptionSettingsResponsePtrOutput) ToCacheEncryptionSettingsResponsePtrOutputWithContext(ctx context.Context) CacheEncryptionSettingsResponsePtrOutput {
	return o
}

func (o CacheEncryptionSettingsResponsePtrOutput) Elem() CacheEncryptionSettingsResponseOutput {
	return o.ApplyT(func(v *CacheEncryptionSettingsResponse) CacheEncryptionSettingsResponse {
		if v != nil {
			return *v
		}
		var ret CacheEncryptionSettingsResponse
		return ret
	}).(CacheEncryptionSettingsResponseOutput)
}

func (o CacheEncryptionSettingsResponsePtrOutput) KeyEncryptionKey() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v *CacheEncryptionSettingsResponse) *KeyVaultKeyReferenceResponse {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKey
	}).(KeyVaultKeyReferenceResponsePtrOutput)
}

type CacheHealthResponse struct {
	Conditions        []ConditionResponse `pulumi:"conditions"`
	State             *string             `pulumi:"state"`
	StatusDescription *string             `pulumi:"statusDescription"`
}





type CacheHealthResponseInput interface {
	pulumi.Input

	ToCacheHealthResponseOutput() CacheHealthResponseOutput
	ToCacheHealthResponseOutputWithContext(context.Context) CacheHealthResponseOutput
}

type CacheHealthResponseArgs struct {
	Conditions        ConditionResponseArrayInput `pulumi:"conditions"`
	State             pulumi.StringPtrInput       `pulumi:"state"`
	StatusDescription pulumi.StringPtrInput       `pulumi:"statusDescription"`
}

func (CacheHealthResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheHealthResponse)(nil)).Elem()
}

func (i CacheHealthResponseArgs) ToCacheHealthResponseOutput() CacheHealthResponseOutput {
	return i.ToCacheHealthResponseOutputWithContext(context.Background())
}

func (i CacheHealthResponseArgs) ToCacheHealthResponseOutputWithContext(ctx context.Context) CacheHealthResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheHealthResponseOutput)
}

func (i CacheHealthResponseArgs) ToCacheHealthResponsePtrOutput() CacheHealthResponsePtrOutput {
	return i.ToCacheHealthResponsePtrOutputWithContext(context.Background())
}

func (i CacheHealthResponseArgs) ToCacheHealthResponsePtrOutputWithContext(ctx context.Context) CacheHealthResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheHealthResponseOutput).ToCacheHealthResponsePtrOutputWithContext(ctx)
}









type CacheHealthResponsePtrInput interface {
	pulumi.Input

	ToCacheHealthResponsePtrOutput() CacheHealthResponsePtrOutput
	ToCacheHealthResponsePtrOutputWithContext(context.Context) CacheHealthResponsePtrOutput
}

type cacheHealthResponsePtrType CacheHealthResponseArgs

func CacheHealthResponsePtr(v *CacheHealthResponseArgs) CacheHealthResponsePtrInput {
	return (*cacheHealthResponsePtrType)(v)
}

func (*cacheHealthResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheHealthResponse)(nil)).Elem()
}

func (i *cacheHealthResponsePtrType) ToCacheHealthResponsePtrOutput() CacheHealthResponsePtrOutput {
	return i.ToCacheHealthResponsePtrOutputWithContext(context.Background())
}

func (i *cacheHealthResponsePtrType) ToCacheHealthResponsePtrOutputWithContext(ctx context.Context) CacheHealthResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheHealthResponsePtrOutput)
}

type CacheHealthResponseOutput struct{ *pulumi.OutputState }

func (CacheHealthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheHealthResponse)(nil)).Elem()
}

func (o CacheHealthResponseOutput) ToCacheHealthResponseOutput() CacheHealthResponseOutput {
	return o
}

func (o CacheHealthResponseOutput) ToCacheHealthResponseOutputWithContext(ctx context.Context) CacheHealthResponseOutput {
	return o
}

func (o CacheHealthResponseOutput) ToCacheHealthResponsePtrOutput() CacheHealthResponsePtrOutput {
	return o.ToCacheHealthResponsePtrOutputWithContext(context.Background())
}

func (o CacheHealthResponseOutput) ToCacheHealthResponsePtrOutputWithContext(ctx context.Context) CacheHealthResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheHealthResponse) *CacheHealthResponse {
		return &v
	}).(CacheHealthResponsePtrOutput)
}

func (o CacheHealthResponseOutput) Conditions() ConditionResponseArrayOutput {
	return o.ApplyT(func(v CacheHealthResponse) []ConditionResponse { return v.Conditions }).(ConditionResponseArrayOutput)
}

func (o CacheHealthResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheHealthResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o CacheHealthResponseOutput) StatusDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheHealthResponse) *string { return v.StatusDescription }).(pulumi.StringPtrOutput)
}

type CacheHealthResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheHealthResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheHealthResponse)(nil)).Elem()
}

func (o CacheHealthResponsePtrOutput) ToCacheHealthResponsePtrOutput() CacheHealthResponsePtrOutput {
	return o
}

func (o CacheHealthResponsePtrOutput) ToCacheHealthResponsePtrOutputWithContext(ctx context.Context) CacheHealthResponsePtrOutput {
	return o
}

func (o CacheHealthResponsePtrOutput) Elem() CacheHealthResponseOutput {
	return o.ApplyT(func(v *CacheHealthResponse) CacheHealthResponse {
		if v != nil {
			return *v
		}
		var ret CacheHealthResponse
		return ret
	}).(CacheHealthResponseOutput)
}

func (o CacheHealthResponsePtrOutput) Conditions() ConditionResponseArrayOutput {
	return o.ApplyT(func(v *CacheHealthResponse) []ConditionResponse {
		if v == nil {
			return nil
		}
		return v.Conditions
	}).(ConditionResponseArrayOutput)
}

func (o CacheHealthResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

func (o CacheHealthResponsePtrOutput) StatusDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.StatusDescription
	}).(pulumi.StringPtrOutput)
}

type CacheIdentity struct {
	Type *CacheIdentityType `pulumi:"type"`
}





type CacheIdentityInput interface {
	pulumi.Input

	ToCacheIdentityOutput() CacheIdentityOutput
	ToCacheIdentityOutputWithContext(context.Context) CacheIdentityOutput
}

type CacheIdentityArgs struct {
	Type CacheIdentityTypePtrInput `pulumi:"type"`
}

func (CacheIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheIdentity)(nil)).Elem()
}

func (i CacheIdentityArgs) ToCacheIdentityOutput() CacheIdentityOutput {
	return i.ToCacheIdentityOutputWithContext(context.Background())
}

func (i CacheIdentityArgs) ToCacheIdentityOutputWithContext(ctx context.Context) CacheIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheIdentityOutput)
}

func (i CacheIdentityArgs) ToCacheIdentityPtrOutput() CacheIdentityPtrOutput {
	return i.ToCacheIdentityPtrOutputWithContext(context.Background())
}

func (i CacheIdentityArgs) ToCacheIdentityPtrOutputWithContext(ctx context.Context) CacheIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheIdentityOutput).ToCacheIdentityPtrOutputWithContext(ctx)
}









type CacheIdentityPtrInput interface {
	pulumi.Input

	ToCacheIdentityPtrOutput() CacheIdentityPtrOutput
	ToCacheIdentityPtrOutputWithContext(context.Context) CacheIdentityPtrOutput
}

type cacheIdentityPtrType CacheIdentityArgs

func CacheIdentityPtr(v *CacheIdentityArgs) CacheIdentityPtrInput {
	return (*cacheIdentityPtrType)(v)
}

func (*cacheIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheIdentity)(nil)).Elem()
}

func (i *cacheIdentityPtrType) ToCacheIdentityPtrOutput() CacheIdentityPtrOutput {
	return i.ToCacheIdentityPtrOutputWithContext(context.Background())
}

func (i *cacheIdentityPtrType) ToCacheIdentityPtrOutputWithContext(ctx context.Context) CacheIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheIdentityPtrOutput)
}

type CacheIdentityOutput struct{ *pulumi.OutputState }

func (CacheIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheIdentity)(nil)).Elem()
}

func (o CacheIdentityOutput) ToCacheIdentityOutput() CacheIdentityOutput {
	return o
}

func (o CacheIdentityOutput) ToCacheIdentityOutputWithContext(ctx context.Context) CacheIdentityOutput {
	return o
}

func (o CacheIdentityOutput) ToCacheIdentityPtrOutput() CacheIdentityPtrOutput {
	return o.ToCacheIdentityPtrOutputWithContext(context.Background())
}

func (o CacheIdentityOutput) ToCacheIdentityPtrOutputWithContext(ctx context.Context) CacheIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheIdentity) *CacheIdentity {
		return &v
	}).(CacheIdentityPtrOutput)
}

func (o CacheIdentityOutput) Type() CacheIdentityTypePtrOutput {
	return o.ApplyT(func(v CacheIdentity) *CacheIdentityType { return v.Type }).(CacheIdentityTypePtrOutput)
}

type CacheIdentityPtrOutput struct{ *pulumi.OutputState }

func (CacheIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheIdentity)(nil)).Elem()
}

func (o CacheIdentityPtrOutput) ToCacheIdentityPtrOutput() CacheIdentityPtrOutput {
	return o
}

func (o CacheIdentityPtrOutput) ToCacheIdentityPtrOutputWithContext(ctx context.Context) CacheIdentityPtrOutput {
	return o
}

func (o CacheIdentityPtrOutput) Elem() CacheIdentityOutput {
	return o.ApplyT(func(v *CacheIdentity) CacheIdentity {
		if v != nil {
			return *v
		}
		var ret CacheIdentity
		return ret
	}).(CacheIdentityOutput)
}

func (o CacheIdentityPtrOutput) Type() CacheIdentityTypePtrOutput {
	return o.ApplyT(func(v *CacheIdentity) *CacheIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(CacheIdentityTypePtrOutput)
}

type CacheIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type CacheIdentityResponseInput interface {
	pulumi.Input

	ToCacheIdentityResponseOutput() CacheIdentityResponseOutput
	ToCacheIdentityResponseOutputWithContext(context.Context) CacheIdentityResponseOutput
}

type CacheIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (CacheIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheIdentityResponse)(nil)).Elem()
}

func (i CacheIdentityResponseArgs) ToCacheIdentityResponseOutput() CacheIdentityResponseOutput {
	return i.ToCacheIdentityResponseOutputWithContext(context.Background())
}

func (i CacheIdentityResponseArgs) ToCacheIdentityResponseOutputWithContext(ctx context.Context) CacheIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheIdentityResponseOutput)
}

func (i CacheIdentityResponseArgs) ToCacheIdentityResponsePtrOutput() CacheIdentityResponsePtrOutput {
	return i.ToCacheIdentityResponsePtrOutputWithContext(context.Background())
}

func (i CacheIdentityResponseArgs) ToCacheIdentityResponsePtrOutputWithContext(ctx context.Context) CacheIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheIdentityResponseOutput).ToCacheIdentityResponsePtrOutputWithContext(ctx)
}









type CacheIdentityResponsePtrInput interface {
	pulumi.Input

	ToCacheIdentityResponsePtrOutput() CacheIdentityResponsePtrOutput
	ToCacheIdentityResponsePtrOutputWithContext(context.Context) CacheIdentityResponsePtrOutput
}

type cacheIdentityResponsePtrType CacheIdentityResponseArgs

func CacheIdentityResponsePtr(v *CacheIdentityResponseArgs) CacheIdentityResponsePtrInput {
	return (*cacheIdentityResponsePtrType)(v)
}

func (*cacheIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheIdentityResponse)(nil)).Elem()
}

func (i *cacheIdentityResponsePtrType) ToCacheIdentityResponsePtrOutput() CacheIdentityResponsePtrOutput {
	return i.ToCacheIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *cacheIdentityResponsePtrType) ToCacheIdentityResponsePtrOutputWithContext(ctx context.Context) CacheIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheIdentityResponsePtrOutput)
}

type CacheIdentityResponseOutput struct{ *pulumi.OutputState }

func (CacheIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheIdentityResponse)(nil)).Elem()
}

func (o CacheIdentityResponseOutput) ToCacheIdentityResponseOutput() CacheIdentityResponseOutput {
	return o
}

func (o CacheIdentityResponseOutput) ToCacheIdentityResponseOutputWithContext(ctx context.Context) CacheIdentityResponseOutput {
	return o
}

func (o CacheIdentityResponseOutput) ToCacheIdentityResponsePtrOutput() CacheIdentityResponsePtrOutput {
	return o.ToCacheIdentityResponsePtrOutputWithContext(context.Background())
}

func (o CacheIdentityResponseOutput) ToCacheIdentityResponsePtrOutputWithContext(ctx context.Context) CacheIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheIdentityResponse) *CacheIdentityResponse {
		return &v
	}).(CacheIdentityResponsePtrOutput)
}

func (o CacheIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v CacheIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o CacheIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v CacheIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o CacheIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CacheIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheIdentityResponse)(nil)).Elem()
}

func (o CacheIdentityResponsePtrOutput) ToCacheIdentityResponsePtrOutput() CacheIdentityResponsePtrOutput {
	return o
}

func (o CacheIdentityResponsePtrOutput) ToCacheIdentityResponsePtrOutputWithContext(ctx context.Context) CacheIdentityResponsePtrOutput {
	return o
}

func (o CacheIdentityResponsePtrOutput) Elem() CacheIdentityResponseOutput {
	return o.ApplyT(func(v *CacheIdentityResponse) CacheIdentityResponse {
		if v != nil {
			return *v
		}
		var ret CacheIdentityResponse
		return ret
	}).(CacheIdentityResponseOutput)
}

func (o CacheIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o CacheIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o CacheIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type CacheNetworkSettings struct {
	DnsSearchDomain *string  `pulumi:"dnsSearchDomain"`
	DnsServers      []string `pulumi:"dnsServers"`
	Mtu             *int     `pulumi:"mtu"`
	NtpServer       *string  `pulumi:"ntpServer"`
}





type CacheNetworkSettingsInput interface {
	pulumi.Input

	ToCacheNetworkSettingsOutput() CacheNetworkSettingsOutput
	ToCacheNetworkSettingsOutputWithContext(context.Context) CacheNetworkSettingsOutput
}

type CacheNetworkSettingsArgs struct {
	DnsSearchDomain pulumi.StringPtrInput   `pulumi:"dnsSearchDomain"`
	DnsServers      pulumi.StringArrayInput `pulumi:"dnsServers"`
	Mtu             pulumi.IntPtrInput      `pulumi:"mtu"`
	NtpServer       pulumi.StringPtrInput   `pulumi:"ntpServer"`
}

func (CacheNetworkSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheNetworkSettings)(nil)).Elem()
}

func (i CacheNetworkSettingsArgs) ToCacheNetworkSettingsOutput() CacheNetworkSettingsOutput {
	return i.ToCacheNetworkSettingsOutputWithContext(context.Background())
}

func (i CacheNetworkSettingsArgs) ToCacheNetworkSettingsOutputWithContext(ctx context.Context) CacheNetworkSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheNetworkSettingsOutput)
}

func (i CacheNetworkSettingsArgs) ToCacheNetworkSettingsPtrOutput() CacheNetworkSettingsPtrOutput {
	return i.ToCacheNetworkSettingsPtrOutputWithContext(context.Background())
}

func (i CacheNetworkSettingsArgs) ToCacheNetworkSettingsPtrOutputWithContext(ctx context.Context) CacheNetworkSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheNetworkSettingsOutput).ToCacheNetworkSettingsPtrOutputWithContext(ctx)
}









type CacheNetworkSettingsPtrInput interface {
	pulumi.Input

	ToCacheNetworkSettingsPtrOutput() CacheNetworkSettingsPtrOutput
	ToCacheNetworkSettingsPtrOutputWithContext(context.Context) CacheNetworkSettingsPtrOutput
}

type cacheNetworkSettingsPtrType CacheNetworkSettingsArgs

func CacheNetworkSettingsPtr(v *CacheNetworkSettingsArgs) CacheNetworkSettingsPtrInput {
	return (*cacheNetworkSettingsPtrType)(v)
}

func (*cacheNetworkSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheNetworkSettings)(nil)).Elem()
}

func (i *cacheNetworkSettingsPtrType) ToCacheNetworkSettingsPtrOutput() CacheNetworkSettingsPtrOutput {
	return i.ToCacheNetworkSettingsPtrOutputWithContext(context.Background())
}

func (i *cacheNetworkSettingsPtrType) ToCacheNetworkSettingsPtrOutputWithContext(ctx context.Context) CacheNetworkSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheNetworkSettingsPtrOutput)
}

type CacheNetworkSettingsOutput struct{ *pulumi.OutputState }

func (CacheNetworkSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheNetworkSettings)(nil)).Elem()
}

func (o CacheNetworkSettingsOutput) ToCacheNetworkSettingsOutput() CacheNetworkSettingsOutput {
	return o
}

func (o CacheNetworkSettingsOutput) ToCacheNetworkSettingsOutputWithContext(ctx context.Context) CacheNetworkSettingsOutput {
	return o
}

func (o CacheNetworkSettingsOutput) ToCacheNetworkSettingsPtrOutput() CacheNetworkSettingsPtrOutput {
	return o.ToCacheNetworkSettingsPtrOutputWithContext(context.Background())
}

func (o CacheNetworkSettingsOutput) ToCacheNetworkSettingsPtrOutputWithContext(ctx context.Context) CacheNetworkSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheNetworkSettings) *CacheNetworkSettings {
		return &v
	}).(CacheNetworkSettingsPtrOutput)
}

func (o CacheNetworkSettingsOutput) DnsSearchDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheNetworkSettings) *string { return v.DnsSearchDomain }).(pulumi.StringPtrOutput)
}

func (o CacheNetworkSettingsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CacheNetworkSettings) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o CacheNetworkSettingsOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheNetworkSettings) *int { return v.Mtu }).(pulumi.IntPtrOutput)
}

func (o CacheNetworkSettingsOutput) NtpServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheNetworkSettings) *string { return v.NtpServer }).(pulumi.StringPtrOutput)
}

type CacheNetworkSettingsPtrOutput struct{ *pulumi.OutputState }

func (CacheNetworkSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheNetworkSettings)(nil)).Elem()
}

func (o CacheNetworkSettingsPtrOutput) ToCacheNetworkSettingsPtrOutput() CacheNetworkSettingsPtrOutput {
	return o
}

func (o CacheNetworkSettingsPtrOutput) ToCacheNetworkSettingsPtrOutputWithContext(ctx context.Context) CacheNetworkSettingsPtrOutput {
	return o
}

func (o CacheNetworkSettingsPtrOutput) Elem() CacheNetworkSettingsOutput {
	return o.ApplyT(func(v *CacheNetworkSettings) CacheNetworkSettings {
		if v != nil {
			return *v
		}
		var ret CacheNetworkSettings
		return ret
	}).(CacheNetworkSettingsOutput)
}

func (o CacheNetworkSettingsPtrOutput) DnsSearchDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheNetworkSettings) *string {
		if v == nil {
			return nil
		}
		return v.DnsSearchDomain
	}).(pulumi.StringPtrOutput)
}

func (o CacheNetworkSettingsPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CacheNetworkSettings) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o CacheNetworkSettingsPtrOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CacheNetworkSettings) *int {
		if v == nil {
			return nil
		}
		return v.Mtu
	}).(pulumi.IntPtrOutput)
}

func (o CacheNetworkSettingsPtrOutput) NtpServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheNetworkSettings) *string {
		if v == nil {
			return nil
		}
		return v.NtpServer
	}).(pulumi.StringPtrOutput)
}

type CacheNetworkSettingsResponse struct {
	DnsSearchDomain  *string  `pulumi:"dnsSearchDomain"`
	DnsServers       []string `pulumi:"dnsServers"`
	Mtu              *int     `pulumi:"mtu"`
	NtpServer        *string  `pulumi:"ntpServer"`
	UtilityAddresses []string `pulumi:"utilityAddresses"`
}





type CacheNetworkSettingsResponseInput interface {
	pulumi.Input

	ToCacheNetworkSettingsResponseOutput() CacheNetworkSettingsResponseOutput
	ToCacheNetworkSettingsResponseOutputWithContext(context.Context) CacheNetworkSettingsResponseOutput
}

type CacheNetworkSettingsResponseArgs struct {
	DnsSearchDomain  pulumi.StringPtrInput   `pulumi:"dnsSearchDomain"`
	DnsServers       pulumi.StringArrayInput `pulumi:"dnsServers"`
	Mtu              pulumi.IntPtrInput      `pulumi:"mtu"`
	NtpServer        pulumi.StringPtrInput   `pulumi:"ntpServer"`
	UtilityAddresses pulumi.StringArrayInput `pulumi:"utilityAddresses"`
}

func (CacheNetworkSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheNetworkSettingsResponse)(nil)).Elem()
}

func (i CacheNetworkSettingsResponseArgs) ToCacheNetworkSettingsResponseOutput() CacheNetworkSettingsResponseOutput {
	return i.ToCacheNetworkSettingsResponseOutputWithContext(context.Background())
}

func (i CacheNetworkSettingsResponseArgs) ToCacheNetworkSettingsResponseOutputWithContext(ctx context.Context) CacheNetworkSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheNetworkSettingsResponseOutput)
}

func (i CacheNetworkSettingsResponseArgs) ToCacheNetworkSettingsResponsePtrOutput() CacheNetworkSettingsResponsePtrOutput {
	return i.ToCacheNetworkSettingsResponsePtrOutputWithContext(context.Background())
}

func (i CacheNetworkSettingsResponseArgs) ToCacheNetworkSettingsResponsePtrOutputWithContext(ctx context.Context) CacheNetworkSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheNetworkSettingsResponseOutput).ToCacheNetworkSettingsResponsePtrOutputWithContext(ctx)
}









type CacheNetworkSettingsResponsePtrInput interface {
	pulumi.Input

	ToCacheNetworkSettingsResponsePtrOutput() CacheNetworkSettingsResponsePtrOutput
	ToCacheNetworkSettingsResponsePtrOutputWithContext(context.Context) CacheNetworkSettingsResponsePtrOutput
}

type cacheNetworkSettingsResponsePtrType CacheNetworkSettingsResponseArgs

func CacheNetworkSettingsResponsePtr(v *CacheNetworkSettingsResponseArgs) CacheNetworkSettingsResponsePtrInput {
	return (*cacheNetworkSettingsResponsePtrType)(v)
}

func (*cacheNetworkSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheNetworkSettingsResponse)(nil)).Elem()
}

func (i *cacheNetworkSettingsResponsePtrType) ToCacheNetworkSettingsResponsePtrOutput() CacheNetworkSettingsResponsePtrOutput {
	return i.ToCacheNetworkSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *cacheNetworkSettingsResponsePtrType) ToCacheNetworkSettingsResponsePtrOutputWithContext(ctx context.Context) CacheNetworkSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheNetworkSettingsResponsePtrOutput)
}

type CacheNetworkSettingsResponseOutput struct{ *pulumi.OutputState }

func (CacheNetworkSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheNetworkSettingsResponse)(nil)).Elem()
}

func (o CacheNetworkSettingsResponseOutput) ToCacheNetworkSettingsResponseOutput() CacheNetworkSettingsResponseOutput {
	return o
}

func (o CacheNetworkSettingsResponseOutput) ToCacheNetworkSettingsResponseOutputWithContext(ctx context.Context) CacheNetworkSettingsResponseOutput {
	return o
}

func (o CacheNetworkSettingsResponseOutput) ToCacheNetworkSettingsResponsePtrOutput() CacheNetworkSettingsResponsePtrOutput {
	return o.ToCacheNetworkSettingsResponsePtrOutputWithContext(context.Background())
}

func (o CacheNetworkSettingsResponseOutput) ToCacheNetworkSettingsResponsePtrOutputWithContext(ctx context.Context) CacheNetworkSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheNetworkSettingsResponse) *CacheNetworkSettingsResponse {
		return &v
	}).(CacheNetworkSettingsResponsePtrOutput)
}

func (o CacheNetworkSettingsResponseOutput) DnsSearchDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheNetworkSettingsResponse) *string { return v.DnsSearchDomain }).(pulumi.StringPtrOutput)
}

func (o CacheNetworkSettingsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CacheNetworkSettingsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o CacheNetworkSettingsResponseOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheNetworkSettingsResponse) *int { return v.Mtu }).(pulumi.IntPtrOutput)
}

func (o CacheNetworkSettingsResponseOutput) NtpServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheNetworkSettingsResponse) *string { return v.NtpServer }).(pulumi.StringPtrOutput)
}

func (o CacheNetworkSettingsResponseOutput) UtilityAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CacheNetworkSettingsResponse) []string { return v.UtilityAddresses }).(pulumi.StringArrayOutput)
}

type CacheNetworkSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheNetworkSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheNetworkSettingsResponse)(nil)).Elem()
}

func (o CacheNetworkSettingsResponsePtrOutput) ToCacheNetworkSettingsResponsePtrOutput() CacheNetworkSettingsResponsePtrOutput {
	return o
}

func (o CacheNetworkSettingsResponsePtrOutput) ToCacheNetworkSettingsResponsePtrOutputWithContext(ctx context.Context) CacheNetworkSettingsResponsePtrOutput {
	return o
}

func (o CacheNetworkSettingsResponsePtrOutput) Elem() CacheNetworkSettingsResponseOutput {
	return o.ApplyT(func(v *CacheNetworkSettingsResponse) CacheNetworkSettingsResponse {
		if v != nil {
			return *v
		}
		var ret CacheNetworkSettingsResponse
		return ret
	}).(CacheNetworkSettingsResponseOutput)
}

func (o CacheNetworkSettingsResponsePtrOutput) DnsSearchDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheNetworkSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsSearchDomain
	}).(pulumi.StringPtrOutput)
}

func (o CacheNetworkSettingsResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CacheNetworkSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o CacheNetworkSettingsResponsePtrOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CacheNetworkSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Mtu
	}).(pulumi.IntPtrOutput)
}

func (o CacheNetworkSettingsResponsePtrOutput) NtpServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheNetworkSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.NtpServer
	}).(pulumi.StringPtrOutput)
}

func (o CacheNetworkSettingsResponsePtrOutput) UtilityAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CacheNetworkSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.UtilityAddresses
	}).(pulumi.StringArrayOutput)
}

type CacheResponseSku struct {
	Name *string `pulumi:"name"`
}





type CacheResponseSkuInput interface {
	pulumi.Input

	ToCacheResponseSkuOutput() CacheResponseSkuOutput
	ToCacheResponseSkuOutputWithContext(context.Context) CacheResponseSkuOutput
}

type CacheResponseSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CacheResponseSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheResponseSku)(nil)).Elem()
}

func (i CacheResponseSkuArgs) ToCacheResponseSkuOutput() CacheResponseSkuOutput {
	return i.ToCacheResponseSkuOutputWithContext(context.Background())
}

func (i CacheResponseSkuArgs) ToCacheResponseSkuOutputWithContext(ctx context.Context) CacheResponseSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheResponseSkuOutput)
}

func (i CacheResponseSkuArgs) ToCacheResponseSkuPtrOutput() CacheResponseSkuPtrOutput {
	return i.ToCacheResponseSkuPtrOutputWithContext(context.Background())
}

func (i CacheResponseSkuArgs) ToCacheResponseSkuPtrOutputWithContext(ctx context.Context) CacheResponseSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheResponseSkuOutput).ToCacheResponseSkuPtrOutputWithContext(ctx)
}









type CacheResponseSkuPtrInput interface {
	pulumi.Input

	ToCacheResponseSkuPtrOutput() CacheResponseSkuPtrOutput
	ToCacheResponseSkuPtrOutputWithContext(context.Context) CacheResponseSkuPtrOutput
}

type cacheResponseSkuPtrType CacheResponseSkuArgs

func CacheResponseSkuPtr(v *CacheResponseSkuArgs) CacheResponseSkuPtrInput {
	return (*cacheResponseSkuPtrType)(v)
}

func (*cacheResponseSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheResponseSku)(nil)).Elem()
}

func (i *cacheResponseSkuPtrType) ToCacheResponseSkuPtrOutput() CacheResponseSkuPtrOutput {
	return i.ToCacheResponseSkuPtrOutputWithContext(context.Background())
}

func (i *cacheResponseSkuPtrType) ToCacheResponseSkuPtrOutputWithContext(ctx context.Context) CacheResponseSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheResponseSkuPtrOutput)
}

type CacheResponseSkuOutput struct{ *pulumi.OutputState }

func (CacheResponseSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheResponseSku)(nil)).Elem()
}

func (o CacheResponseSkuOutput) ToCacheResponseSkuOutput() CacheResponseSkuOutput {
	return o
}

func (o CacheResponseSkuOutput) ToCacheResponseSkuOutputWithContext(ctx context.Context) CacheResponseSkuOutput {
	return o
}

func (o CacheResponseSkuOutput) ToCacheResponseSkuPtrOutput() CacheResponseSkuPtrOutput {
	return o.ToCacheResponseSkuPtrOutputWithContext(context.Background())
}

func (o CacheResponseSkuOutput) ToCacheResponseSkuPtrOutputWithContext(ctx context.Context) CacheResponseSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheResponseSku) *CacheResponseSku {
		return &v
	}).(CacheResponseSkuPtrOutput)
}

func (o CacheResponseSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheResponseSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CacheResponseSkuPtrOutput struct{ *pulumi.OutputState }

func (CacheResponseSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheResponseSku)(nil)).Elem()
}

func (o CacheResponseSkuPtrOutput) ToCacheResponseSkuPtrOutput() CacheResponseSkuPtrOutput {
	return o
}

func (o CacheResponseSkuPtrOutput) ToCacheResponseSkuPtrOutputWithContext(ctx context.Context) CacheResponseSkuPtrOutput {
	return o
}

func (o CacheResponseSkuPtrOutput) Elem() CacheResponseSkuOutput {
	return o.ApplyT(func(v *CacheResponseSku) CacheResponseSku {
		if v != nil {
			return *v
		}
		var ret CacheResponseSku
		return ret
	}).(CacheResponseSkuOutput)
}

func (o CacheResponseSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheResponseSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type CacheSecuritySettings struct {
	AccessPolicies []NfsAccessPolicy `pulumi:"accessPolicies"`
}





type CacheSecuritySettingsInput interface {
	pulumi.Input

	ToCacheSecuritySettingsOutput() CacheSecuritySettingsOutput
	ToCacheSecuritySettingsOutputWithContext(context.Context) CacheSecuritySettingsOutput
}

type CacheSecuritySettingsArgs struct {
	AccessPolicies NfsAccessPolicyArrayInput `pulumi:"accessPolicies"`
}

func (CacheSecuritySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheSecuritySettings)(nil)).Elem()
}

func (i CacheSecuritySettingsArgs) ToCacheSecuritySettingsOutput() CacheSecuritySettingsOutput {
	return i.ToCacheSecuritySettingsOutputWithContext(context.Background())
}

func (i CacheSecuritySettingsArgs) ToCacheSecuritySettingsOutputWithContext(ctx context.Context) CacheSecuritySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSecuritySettingsOutput)
}

func (i CacheSecuritySettingsArgs) ToCacheSecuritySettingsPtrOutput() CacheSecuritySettingsPtrOutput {
	return i.ToCacheSecuritySettingsPtrOutputWithContext(context.Background())
}

func (i CacheSecuritySettingsArgs) ToCacheSecuritySettingsPtrOutputWithContext(ctx context.Context) CacheSecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSecuritySettingsOutput).ToCacheSecuritySettingsPtrOutputWithContext(ctx)
}









type CacheSecuritySettingsPtrInput interface {
	pulumi.Input

	ToCacheSecuritySettingsPtrOutput() CacheSecuritySettingsPtrOutput
	ToCacheSecuritySettingsPtrOutputWithContext(context.Context) CacheSecuritySettingsPtrOutput
}

type cacheSecuritySettingsPtrType CacheSecuritySettingsArgs

func CacheSecuritySettingsPtr(v *CacheSecuritySettingsArgs) CacheSecuritySettingsPtrInput {
	return (*cacheSecuritySettingsPtrType)(v)
}

func (*cacheSecuritySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheSecuritySettings)(nil)).Elem()
}

func (i *cacheSecuritySettingsPtrType) ToCacheSecuritySettingsPtrOutput() CacheSecuritySettingsPtrOutput {
	return i.ToCacheSecuritySettingsPtrOutputWithContext(context.Background())
}

func (i *cacheSecuritySettingsPtrType) ToCacheSecuritySettingsPtrOutputWithContext(ctx context.Context) CacheSecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSecuritySettingsPtrOutput)
}

type CacheSecuritySettingsOutput struct{ *pulumi.OutputState }

func (CacheSecuritySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheSecuritySettings)(nil)).Elem()
}

func (o CacheSecuritySettingsOutput) ToCacheSecuritySettingsOutput() CacheSecuritySettingsOutput {
	return o
}

func (o CacheSecuritySettingsOutput) ToCacheSecuritySettingsOutputWithContext(ctx context.Context) CacheSecuritySettingsOutput {
	return o
}

func (o CacheSecuritySettingsOutput) ToCacheSecuritySettingsPtrOutput() CacheSecuritySettingsPtrOutput {
	return o.ToCacheSecuritySettingsPtrOutputWithContext(context.Background())
}

func (o CacheSecuritySettingsOutput) ToCacheSecuritySettingsPtrOutputWithContext(ctx context.Context) CacheSecuritySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheSecuritySettings) *CacheSecuritySettings {
		return &v
	}).(CacheSecuritySettingsPtrOutput)
}

func (o CacheSecuritySettingsOutput) AccessPolicies() NfsAccessPolicyArrayOutput {
	return o.ApplyT(func(v CacheSecuritySettings) []NfsAccessPolicy { return v.AccessPolicies }).(NfsAccessPolicyArrayOutput)
}

type CacheSecuritySettingsPtrOutput struct{ *pulumi.OutputState }

func (CacheSecuritySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheSecuritySettings)(nil)).Elem()
}

func (o CacheSecuritySettingsPtrOutput) ToCacheSecuritySettingsPtrOutput() CacheSecuritySettingsPtrOutput {
	return o
}

func (o CacheSecuritySettingsPtrOutput) ToCacheSecuritySettingsPtrOutputWithContext(ctx context.Context) CacheSecuritySettingsPtrOutput {
	return o
}

func (o CacheSecuritySettingsPtrOutput) Elem() CacheSecuritySettingsOutput {
	return o.ApplyT(func(v *CacheSecuritySettings) CacheSecuritySettings {
		if v != nil {
			return *v
		}
		var ret CacheSecuritySettings
		return ret
	}).(CacheSecuritySettingsOutput)
}

func (o CacheSecuritySettingsPtrOutput) AccessPolicies() NfsAccessPolicyArrayOutput {
	return o.ApplyT(func(v *CacheSecuritySettings) []NfsAccessPolicy {
		if v == nil {
			return nil
		}
		return v.AccessPolicies
	}).(NfsAccessPolicyArrayOutput)
}

type CacheSecuritySettingsResponse struct {
	AccessPolicies []NfsAccessPolicyResponse `pulumi:"accessPolicies"`
}





type CacheSecuritySettingsResponseInput interface {
	pulumi.Input

	ToCacheSecuritySettingsResponseOutput() CacheSecuritySettingsResponseOutput
	ToCacheSecuritySettingsResponseOutputWithContext(context.Context) CacheSecuritySettingsResponseOutput
}

type CacheSecuritySettingsResponseArgs struct {
	AccessPolicies NfsAccessPolicyResponseArrayInput `pulumi:"accessPolicies"`
}

func (CacheSecuritySettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheSecuritySettingsResponse)(nil)).Elem()
}

func (i CacheSecuritySettingsResponseArgs) ToCacheSecuritySettingsResponseOutput() CacheSecuritySettingsResponseOutput {
	return i.ToCacheSecuritySettingsResponseOutputWithContext(context.Background())
}

func (i CacheSecuritySettingsResponseArgs) ToCacheSecuritySettingsResponseOutputWithContext(ctx context.Context) CacheSecuritySettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSecuritySettingsResponseOutput)
}

func (i CacheSecuritySettingsResponseArgs) ToCacheSecuritySettingsResponsePtrOutput() CacheSecuritySettingsResponsePtrOutput {
	return i.ToCacheSecuritySettingsResponsePtrOutputWithContext(context.Background())
}

func (i CacheSecuritySettingsResponseArgs) ToCacheSecuritySettingsResponsePtrOutputWithContext(ctx context.Context) CacheSecuritySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSecuritySettingsResponseOutput).ToCacheSecuritySettingsResponsePtrOutputWithContext(ctx)
}









type CacheSecuritySettingsResponsePtrInput interface {
	pulumi.Input

	ToCacheSecuritySettingsResponsePtrOutput() CacheSecuritySettingsResponsePtrOutput
	ToCacheSecuritySettingsResponsePtrOutputWithContext(context.Context) CacheSecuritySettingsResponsePtrOutput
}

type cacheSecuritySettingsResponsePtrType CacheSecuritySettingsResponseArgs

func CacheSecuritySettingsResponsePtr(v *CacheSecuritySettingsResponseArgs) CacheSecuritySettingsResponsePtrInput {
	return (*cacheSecuritySettingsResponsePtrType)(v)
}

func (*cacheSecuritySettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheSecuritySettingsResponse)(nil)).Elem()
}

func (i *cacheSecuritySettingsResponsePtrType) ToCacheSecuritySettingsResponsePtrOutput() CacheSecuritySettingsResponsePtrOutput {
	return i.ToCacheSecuritySettingsResponsePtrOutputWithContext(context.Background())
}

func (i *cacheSecuritySettingsResponsePtrType) ToCacheSecuritySettingsResponsePtrOutputWithContext(ctx context.Context) CacheSecuritySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSecuritySettingsResponsePtrOutput)
}

type CacheSecuritySettingsResponseOutput struct{ *pulumi.OutputState }

func (CacheSecuritySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheSecuritySettingsResponse)(nil)).Elem()
}

func (o CacheSecuritySettingsResponseOutput) ToCacheSecuritySettingsResponseOutput() CacheSecuritySettingsResponseOutput {
	return o
}

func (o CacheSecuritySettingsResponseOutput) ToCacheSecuritySettingsResponseOutputWithContext(ctx context.Context) CacheSecuritySettingsResponseOutput {
	return o
}

func (o CacheSecuritySettingsResponseOutput) ToCacheSecuritySettingsResponsePtrOutput() CacheSecuritySettingsResponsePtrOutput {
	return o.ToCacheSecuritySettingsResponsePtrOutputWithContext(context.Background())
}

func (o CacheSecuritySettingsResponseOutput) ToCacheSecuritySettingsResponsePtrOutputWithContext(ctx context.Context) CacheSecuritySettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheSecuritySettingsResponse) *CacheSecuritySettingsResponse {
		return &v
	}).(CacheSecuritySettingsResponsePtrOutput)
}

func (o CacheSecuritySettingsResponseOutput) AccessPolicies() NfsAccessPolicyResponseArrayOutput {
	return o.ApplyT(func(v CacheSecuritySettingsResponse) []NfsAccessPolicyResponse { return v.AccessPolicies }).(NfsAccessPolicyResponseArrayOutput)
}

type CacheSecuritySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheSecuritySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheSecuritySettingsResponse)(nil)).Elem()
}

func (o CacheSecuritySettingsResponsePtrOutput) ToCacheSecuritySettingsResponsePtrOutput() CacheSecuritySettingsResponsePtrOutput {
	return o
}

func (o CacheSecuritySettingsResponsePtrOutput) ToCacheSecuritySettingsResponsePtrOutputWithContext(ctx context.Context) CacheSecuritySettingsResponsePtrOutput {
	return o
}

func (o CacheSecuritySettingsResponsePtrOutput) Elem() CacheSecuritySettingsResponseOutput {
	return o.ApplyT(func(v *CacheSecuritySettingsResponse) CacheSecuritySettingsResponse {
		if v != nil {
			return *v
		}
		var ret CacheSecuritySettingsResponse
		return ret
	}).(CacheSecuritySettingsResponseOutput)
}

func (o CacheSecuritySettingsResponsePtrOutput) AccessPolicies() NfsAccessPolicyResponseArrayOutput {
	return o.ApplyT(func(v *CacheSecuritySettingsResponse) []NfsAccessPolicyResponse {
		if v == nil {
			return nil
		}
		return v.AccessPolicies
	}).(NfsAccessPolicyResponseArrayOutput)
}

type CacheSku struct {
	Name *string `pulumi:"name"`
}





type CacheSkuInput interface {
	pulumi.Input

	ToCacheSkuOutput() CacheSkuOutput
	ToCacheSkuOutputWithContext(context.Context) CacheSkuOutput
}

type CacheSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CacheSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheSku)(nil)).Elem()
}

func (i CacheSkuArgs) ToCacheSkuOutput() CacheSkuOutput {
	return i.ToCacheSkuOutputWithContext(context.Background())
}

func (i CacheSkuArgs) ToCacheSkuOutputWithContext(ctx context.Context) CacheSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSkuOutput)
}

func (i CacheSkuArgs) ToCacheSkuPtrOutput() CacheSkuPtrOutput {
	return i.ToCacheSkuPtrOutputWithContext(context.Background())
}

func (i CacheSkuArgs) ToCacheSkuPtrOutputWithContext(ctx context.Context) CacheSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSkuOutput).ToCacheSkuPtrOutputWithContext(ctx)
}









type CacheSkuPtrInput interface {
	pulumi.Input

	ToCacheSkuPtrOutput() CacheSkuPtrOutput
	ToCacheSkuPtrOutputWithContext(context.Context) CacheSkuPtrOutput
}

type cacheSkuPtrType CacheSkuArgs

func CacheSkuPtr(v *CacheSkuArgs) CacheSkuPtrInput {
	return (*cacheSkuPtrType)(v)
}

func (*cacheSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheSku)(nil)).Elem()
}

func (i *cacheSkuPtrType) ToCacheSkuPtrOutput() CacheSkuPtrOutput {
	return i.ToCacheSkuPtrOutputWithContext(context.Background())
}

func (i *cacheSkuPtrType) ToCacheSkuPtrOutputWithContext(ctx context.Context) CacheSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSkuPtrOutput)
}

type CacheSkuOutput struct{ *pulumi.OutputState }

func (CacheSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheSku)(nil)).Elem()
}

func (o CacheSkuOutput) ToCacheSkuOutput() CacheSkuOutput {
	return o
}

func (o CacheSkuOutput) ToCacheSkuOutputWithContext(ctx context.Context) CacheSkuOutput {
	return o
}

func (o CacheSkuOutput) ToCacheSkuPtrOutput() CacheSkuPtrOutput {
	return o.ToCacheSkuPtrOutputWithContext(context.Background())
}

func (o CacheSkuOutput) ToCacheSkuPtrOutputWithContext(ctx context.Context) CacheSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheSku) *CacheSku {
		return &v
	}).(CacheSkuPtrOutput)
}

func (o CacheSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CacheSkuPtrOutput struct{ *pulumi.OutputState }

func (CacheSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheSku)(nil)).Elem()
}

func (o CacheSkuPtrOutput) ToCacheSkuPtrOutput() CacheSkuPtrOutput {
	return o
}

func (o CacheSkuPtrOutput) ToCacheSkuPtrOutputWithContext(ctx context.Context) CacheSkuPtrOutput {
	return o
}

func (o CacheSkuPtrOutput) Elem() CacheSkuOutput {
	return o.ApplyT(func(v *CacheSku) CacheSku {
		if v != nil {
			return *v
		}
		var ret CacheSku
		return ret
	}).(CacheSkuOutput)
}

func (o CacheSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type CacheUpgradeStatusResponse struct {
	CurrentFirmwareVersion string `pulumi:"currentFirmwareVersion"`
	FirmwareUpdateDeadline string `pulumi:"firmwareUpdateDeadline"`
	FirmwareUpdateStatus   string `pulumi:"firmwareUpdateStatus"`
	LastFirmwareUpdate     string `pulumi:"lastFirmwareUpdate"`
	PendingFirmwareVersion string `pulumi:"pendingFirmwareVersion"`
}





type CacheUpgradeStatusResponseInput interface {
	pulumi.Input

	ToCacheUpgradeStatusResponseOutput() CacheUpgradeStatusResponseOutput
	ToCacheUpgradeStatusResponseOutputWithContext(context.Context) CacheUpgradeStatusResponseOutput
}

type CacheUpgradeStatusResponseArgs struct {
	CurrentFirmwareVersion pulumi.StringInput `pulumi:"currentFirmwareVersion"`
	FirmwareUpdateDeadline pulumi.StringInput `pulumi:"firmwareUpdateDeadline"`
	FirmwareUpdateStatus   pulumi.StringInput `pulumi:"firmwareUpdateStatus"`
	LastFirmwareUpdate     pulumi.StringInput `pulumi:"lastFirmwareUpdate"`
	PendingFirmwareVersion pulumi.StringInput `pulumi:"pendingFirmwareVersion"`
}

func (CacheUpgradeStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUpgradeStatusResponse)(nil)).Elem()
}

func (i CacheUpgradeStatusResponseArgs) ToCacheUpgradeStatusResponseOutput() CacheUpgradeStatusResponseOutput {
	return i.ToCacheUpgradeStatusResponseOutputWithContext(context.Background())
}

func (i CacheUpgradeStatusResponseArgs) ToCacheUpgradeStatusResponseOutputWithContext(ctx context.Context) CacheUpgradeStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUpgradeStatusResponseOutput)
}

func (i CacheUpgradeStatusResponseArgs) ToCacheUpgradeStatusResponsePtrOutput() CacheUpgradeStatusResponsePtrOutput {
	return i.ToCacheUpgradeStatusResponsePtrOutputWithContext(context.Background())
}

func (i CacheUpgradeStatusResponseArgs) ToCacheUpgradeStatusResponsePtrOutputWithContext(ctx context.Context) CacheUpgradeStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUpgradeStatusResponseOutput).ToCacheUpgradeStatusResponsePtrOutputWithContext(ctx)
}









type CacheUpgradeStatusResponsePtrInput interface {
	pulumi.Input

	ToCacheUpgradeStatusResponsePtrOutput() CacheUpgradeStatusResponsePtrOutput
	ToCacheUpgradeStatusResponsePtrOutputWithContext(context.Context) CacheUpgradeStatusResponsePtrOutput
}

type cacheUpgradeStatusResponsePtrType CacheUpgradeStatusResponseArgs

func CacheUpgradeStatusResponsePtr(v *CacheUpgradeStatusResponseArgs) CacheUpgradeStatusResponsePtrInput {
	return (*cacheUpgradeStatusResponsePtrType)(v)
}

func (*cacheUpgradeStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUpgradeStatusResponse)(nil)).Elem()
}

func (i *cacheUpgradeStatusResponsePtrType) ToCacheUpgradeStatusResponsePtrOutput() CacheUpgradeStatusResponsePtrOutput {
	return i.ToCacheUpgradeStatusResponsePtrOutputWithContext(context.Background())
}

func (i *cacheUpgradeStatusResponsePtrType) ToCacheUpgradeStatusResponsePtrOutputWithContext(ctx context.Context) CacheUpgradeStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUpgradeStatusResponsePtrOutput)
}

type CacheUpgradeStatusResponseOutput struct{ *pulumi.OutputState }

func (CacheUpgradeStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUpgradeStatusResponse)(nil)).Elem()
}

func (o CacheUpgradeStatusResponseOutput) ToCacheUpgradeStatusResponseOutput() CacheUpgradeStatusResponseOutput {
	return o
}

func (o CacheUpgradeStatusResponseOutput) ToCacheUpgradeStatusResponseOutputWithContext(ctx context.Context) CacheUpgradeStatusResponseOutput {
	return o
}

func (o CacheUpgradeStatusResponseOutput) ToCacheUpgradeStatusResponsePtrOutput() CacheUpgradeStatusResponsePtrOutput {
	return o.ToCacheUpgradeStatusResponsePtrOutputWithContext(context.Background())
}

func (o CacheUpgradeStatusResponseOutput) ToCacheUpgradeStatusResponsePtrOutputWithContext(ctx context.Context) CacheUpgradeStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheUpgradeStatusResponse) *CacheUpgradeStatusResponse {
		return &v
	}).(CacheUpgradeStatusResponsePtrOutput)
}

func (o CacheUpgradeStatusResponseOutput) CurrentFirmwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUpgradeStatusResponse) string { return v.CurrentFirmwareVersion }).(pulumi.StringOutput)
}

func (o CacheUpgradeStatusResponseOutput) FirmwareUpdateDeadline() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUpgradeStatusResponse) string { return v.FirmwareUpdateDeadline }).(pulumi.StringOutput)
}

func (o CacheUpgradeStatusResponseOutput) FirmwareUpdateStatus() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUpgradeStatusResponse) string { return v.FirmwareUpdateStatus }).(pulumi.StringOutput)
}

func (o CacheUpgradeStatusResponseOutput) LastFirmwareUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUpgradeStatusResponse) string { return v.LastFirmwareUpdate }).(pulumi.StringOutput)
}

func (o CacheUpgradeStatusResponseOutput) PendingFirmwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUpgradeStatusResponse) string { return v.PendingFirmwareVersion }).(pulumi.StringOutput)
}

type CacheUpgradeStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheUpgradeStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUpgradeStatusResponse)(nil)).Elem()
}

func (o CacheUpgradeStatusResponsePtrOutput) ToCacheUpgradeStatusResponsePtrOutput() CacheUpgradeStatusResponsePtrOutput {
	return o
}

func (o CacheUpgradeStatusResponsePtrOutput) ToCacheUpgradeStatusResponsePtrOutputWithContext(ctx context.Context) CacheUpgradeStatusResponsePtrOutput {
	return o
}

func (o CacheUpgradeStatusResponsePtrOutput) Elem() CacheUpgradeStatusResponseOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) CacheUpgradeStatusResponse {
		if v != nil {
			return *v
		}
		var ret CacheUpgradeStatusResponse
		return ret
	}).(CacheUpgradeStatusResponseOutput)
}

func (o CacheUpgradeStatusResponsePtrOutput) CurrentFirmwareVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CurrentFirmwareVersion
	}).(pulumi.StringPtrOutput)
}

func (o CacheUpgradeStatusResponsePtrOutput) FirmwareUpdateDeadline() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FirmwareUpdateDeadline
	}).(pulumi.StringPtrOutput)
}

func (o CacheUpgradeStatusResponsePtrOutput) FirmwareUpdateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FirmwareUpdateStatus
	}).(pulumi.StringPtrOutput)
}

func (o CacheUpgradeStatusResponsePtrOutput) LastFirmwareUpdate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastFirmwareUpdate
	}).(pulumi.StringPtrOutput)
}

func (o CacheUpgradeStatusResponsePtrOutput) PendingFirmwareVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PendingFirmwareVersion
	}).(pulumi.StringPtrOutput)
}

type CacheUsernameDownloadSettings struct {
	AutoDownloadCertificate *bool                                     `pulumi:"autoDownloadCertificate"`
	CaCertificateURI        *string                                   `pulumi:"caCertificateURI"`
	Credentials             *CacheUsernameDownloadSettingsCredentials `pulumi:"credentials"`
	EncryptLdapConnection   *bool                                     `pulumi:"encryptLdapConnection"`
	ExtendedGroups          *bool                                     `pulumi:"extendedGroups"`
	GroupFileURI            *string                                   `pulumi:"groupFileURI"`
	LdapBaseDN              *string                                   `pulumi:"ldapBaseDN"`
	LdapServer              *string                                   `pulumi:"ldapServer"`
	RequireValidCertificate *bool                                     `pulumi:"requireValidCertificate"`
	UserFileURI             *string                                   `pulumi:"userFileURI"`
	UsernameSource          *string                                   `pulumi:"usernameSource"`
}





type CacheUsernameDownloadSettingsInput interface {
	pulumi.Input

	ToCacheUsernameDownloadSettingsOutput() CacheUsernameDownloadSettingsOutput
	ToCacheUsernameDownloadSettingsOutputWithContext(context.Context) CacheUsernameDownloadSettingsOutput
}

type CacheUsernameDownloadSettingsArgs struct {
	AutoDownloadCertificate pulumi.BoolPtrInput                              `pulumi:"autoDownloadCertificate"`
	CaCertificateURI        pulumi.StringPtrInput                            `pulumi:"caCertificateURI"`
	Credentials             CacheUsernameDownloadSettingsCredentialsPtrInput `pulumi:"credentials"`
	EncryptLdapConnection   pulumi.BoolPtrInput                              `pulumi:"encryptLdapConnection"`
	ExtendedGroups          pulumi.BoolPtrInput                              `pulumi:"extendedGroups"`
	GroupFileURI            pulumi.StringPtrInput                            `pulumi:"groupFileURI"`
	LdapBaseDN              pulumi.StringPtrInput                            `pulumi:"ldapBaseDN"`
	LdapServer              pulumi.StringPtrInput                            `pulumi:"ldapServer"`
	RequireValidCertificate pulumi.BoolPtrInput                              `pulumi:"requireValidCertificate"`
	UserFileURI             pulumi.StringPtrInput                            `pulumi:"userFileURI"`
	UsernameSource          pulumi.StringPtrInput                            `pulumi:"usernameSource"`
}

func (CacheUsernameDownloadSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUsernameDownloadSettings)(nil)).Elem()
}

func (i CacheUsernameDownloadSettingsArgs) ToCacheUsernameDownloadSettingsOutput() CacheUsernameDownloadSettingsOutput {
	return i.ToCacheUsernameDownloadSettingsOutputWithContext(context.Background())
}

func (i CacheUsernameDownloadSettingsArgs) ToCacheUsernameDownloadSettingsOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsOutput)
}

func (i CacheUsernameDownloadSettingsArgs) ToCacheUsernameDownloadSettingsPtrOutput() CacheUsernameDownloadSettingsPtrOutput {
	return i.ToCacheUsernameDownloadSettingsPtrOutputWithContext(context.Background())
}

func (i CacheUsernameDownloadSettingsArgs) ToCacheUsernameDownloadSettingsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsOutput).ToCacheUsernameDownloadSettingsPtrOutputWithContext(ctx)
}









type CacheUsernameDownloadSettingsPtrInput interface {
	pulumi.Input

	ToCacheUsernameDownloadSettingsPtrOutput() CacheUsernameDownloadSettingsPtrOutput
	ToCacheUsernameDownloadSettingsPtrOutputWithContext(context.Context) CacheUsernameDownloadSettingsPtrOutput
}

type cacheUsernameDownloadSettingsPtrType CacheUsernameDownloadSettingsArgs

func CacheUsernameDownloadSettingsPtr(v *CacheUsernameDownloadSettingsArgs) CacheUsernameDownloadSettingsPtrInput {
	return (*cacheUsernameDownloadSettingsPtrType)(v)
}

func (*cacheUsernameDownloadSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUsernameDownloadSettings)(nil)).Elem()
}

func (i *cacheUsernameDownloadSettingsPtrType) ToCacheUsernameDownloadSettingsPtrOutput() CacheUsernameDownloadSettingsPtrOutput {
	return i.ToCacheUsernameDownloadSettingsPtrOutputWithContext(context.Background())
}

func (i *cacheUsernameDownloadSettingsPtrType) ToCacheUsernameDownloadSettingsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsPtrOutput)
}

type CacheUsernameDownloadSettingsOutput struct{ *pulumi.OutputState }

func (CacheUsernameDownloadSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUsernameDownloadSettings)(nil)).Elem()
}

func (o CacheUsernameDownloadSettingsOutput) ToCacheUsernameDownloadSettingsOutput() CacheUsernameDownloadSettingsOutput {
	return o
}

func (o CacheUsernameDownloadSettingsOutput) ToCacheUsernameDownloadSettingsOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsOutput {
	return o
}

func (o CacheUsernameDownloadSettingsOutput) ToCacheUsernameDownloadSettingsPtrOutput() CacheUsernameDownloadSettingsPtrOutput {
	return o.ToCacheUsernameDownloadSettingsPtrOutputWithContext(context.Background())
}

func (o CacheUsernameDownloadSettingsOutput) ToCacheUsernameDownloadSettingsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheUsernameDownloadSettings) *CacheUsernameDownloadSettings {
		return &v
	}).(CacheUsernameDownloadSettingsPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) AutoDownloadCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *bool { return v.AutoDownloadCertificate }).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) CaCertificateURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *string { return v.CaCertificateURI }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) Credentials() CacheUsernameDownloadSettingsCredentialsPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *CacheUsernameDownloadSettingsCredentials { return v.Credentials }).(CacheUsernameDownloadSettingsCredentialsPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) EncryptLdapConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *bool { return v.EncryptLdapConnection }).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) ExtendedGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *bool { return v.ExtendedGroups }).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) GroupFileURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *string { return v.GroupFileURI }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) LdapBaseDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *string { return v.LdapBaseDN }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) LdapServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *string { return v.LdapServer }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) RequireValidCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *bool { return v.RequireValidCertificate }).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) UserFileURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *string { return v.UserFileURI }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsOutput) UsernameSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettings) *string { return v.UsernameSource }).(pulumi.StringPtrOutput)
}

type CacheUsernameDownloadSettingsPtrOutput struct{ *pulumi.OutputState }

func (CacheUsernameDownloadSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUsernameDownloadSettings)(nil)).Elem()
}

func (o CacheUsernameDownloadSettingsPtrOutput) ToCacheUsernameDownloadSettingsPtrOutput() CacheUsernameDownloadSettingsPtrOutput {
	return o
}

func (o CacheUsernameDownloadSettingsPtrOutput) ToCacheUsernameDownloadSettingsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsPtrOutput {
	return o
}

func (o CacheUsernameDownloadSettingsPtrOutput) Elem() CacheUsernameDownloadSettingsOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) CacheUsernameDownloadSettings {
		if v != nil {
			return *v
		}
		var ret CacheUsernameDownloadSettings
		return ret
	}).(CacheUsernameDownloadSettingsOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) AutoDownloadCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *bool {
		if v == nil {
			return nil
		}
		return v.AutoDownloadCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) CaCertificateURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *string {
		if v == nil {
			return nil
		}
		return v.CaCertificateURI
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) Credentials() CacheUsernameDownloadSettingsCredentialsPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *CacheUsernameDownloadSettingsCredentials {
		if v == nil {
			return nil
		}
		return v.Credentials
	}).(CacheUsernameDownloadSettingsCredentialsPtrOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) EncryptLdapConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptLdapConnection
	}).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) ExtendedGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ExtendedGroups
	}).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) GroupFileURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupFileURI
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) LdapBaseDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *string {
		if v == nil {
			return nil
		}
		return v.LdapBaseDN
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) LdapServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *string {
		if v == nil {
			return nil
		}
		return v.LdapServer
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) RequireValidCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RequireValidCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) UserFileURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *string {
		if v == nil {
			return nil
		}
		return v.UserFileURI
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsPtrOutput) UsernameSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettings) *string {
		if v == nil {
			return nil
		}
		return v.UsernameSource
	}).(pulumi.StringPtrOutput)
}

type CacheUsernameDownloadSettingsCredentials struct {
	BindDn       *string `pulumi:"bindDn"`
	BindPassword *string `pulumi:"bindPassword"`
}





type CacheUsernameDownloadSettingsCredentialsInput interface {
	pulumi.Input

	ToCacheUsernameDownloadSettingsCredentialsOutput() CacheUsernameDownloadSettingsCredentialsOutput
	ToCacheUsernameDownloadSettingsCredentialsOutputWithContext(context.Context) CacheUsernameDownloadSettingsCredentialsOutput
}

type CacheUsernameDownloadSettingsCredentialsArgs struct {
	BindDn       pulumi.StringPtrInput `pulumi:"bindDn"`
	BindPassword pulumi.StringPtrInput `pulumi:"bindPassword"`
}

func (CacheUsernameDownloadSettingsCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUsernameDownloadSettingsCredentials)(nil)).Elem()
}

func (i CacheUsernameDownloadSettingsCredentialsArgs) ToCacheUsernameDownloadSettingsCredentialsOutput() CacheUsernameDownloadSettingsCredentialsOutput {
	return i.ToCacheUsernameDownloadSettingsCredentialsOutputWithContext(context.Background())
}

func (i CacheUsernameDownloadSettingsCredentialsArgs) ToCacheUsernameDownloadSettingsCredentialsOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsCredentialsOutput)
}

func (i CacheUsernameDownloadSettingsCredentialsArgs) ToCacheUsernameDownloadSettingsCredentialsPtrOutput() CacheUsernameDownloadSettingsCredentialsPtrOutput {
	return i.ToCacheUsernameDownloadSettingsCredentialsPtrOutputWithContext(context.Background())
}

func (i CacheUsernameDownloadSettingsCredentialsArgs) ToCacheUsernameDownloadSettingsCredentialsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsCredentialsOutput).ToCacheUsernameDownloadSettingsCredentialsPtrOutputWithContext(ctx)
}









type CacheUsernameDownloadSettingsCredentialsPtrInput interface {
	pulumi.Input

	ToCacheUsernameDownloadSettingsCredentialsPtrOutput() CacheUsernameDownloadSettingsCredentialsPtrOutput
	ToCacheUsernameDownloadSettingsCredentialsPtrOutputWithContext(context.Context) CacheUsernameDownloadSettingsCredentialsPtrOutput
}

type cacheUsernameDownloadSettingsCredentialsPtrType CacheUsernameDownloadSettingsCredentialsArgs

func CacheUsernameDownloadSettingsCredentialsPtr(v *CacheUsernameDownloadSettingsCredentialsArgs) CacheUsernameDownloadSettingsCredentialsPtrInput {
	return (*cacheUsernameDownloadSettingsCredentialsPtrType)(v)
}

func (*cacheUsernameDownloadSettingsCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUsernameDownloadSettingsCredentials)(nil)).Elem()
}

func (i *cacheUsernameDownloadSettingsCredentialsPtrType) ToCacheUsernameDownloadSettingsCredentialsPtrOutput() CacheUsernameDownloadSettingsCredentialsPtrOutput {
	return i.ToCacheUsernameDownloadSettingsCredentialsPtrOutputWithContext(context.Background())
}

func (i *cacheUsernameDownloadSettingsCredentialsPtrType) ToCacheUsernameDownloadSettingsCredentialsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsCredentialsPtrOutput)
}

type CacheUsernameDownloadSettingsCredentialsOutput struct{ *pulumi.OutputState }

func (CacheUsernameDownloadSettingsCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUsernameDownloadSettingsCredentials)(nil)).Elem()
}

func (o CacheUsernameDownloadSettingsCredentialsOutput) ToCacheUsernameDownloadSettingsCredentialsOutput() CacheUsernameDownloadSettingsCredentialsOutput {
	return o
}

func (o CacheUsernameDownloadSettingsCredentialsOutput) ToCacheUsernameDownloadSettingsCredentialsOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsCredentialsOutput {
	return o
}

func (o CacheUsernameDownloadSettingsCredentialsOutput) ToCacheUsernameDownloadSettingsCredentialsPtrOutput() CacheUsernameDownloadSettingsCredentialsPtrOutput {
	return o.ToCacheUsernameDownloadSettingsCredentialsPtrOutputWithContext(context.Background())
}

func (o CacheUsernameDownloadSettingsCredentialsOutput) ToCacheUsernameDownloadSettingsCredentialsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheUsernameDownloadSettingsCredentials) *CacheUsernameDownloadSettingsCredentials {
		return &v
	}).(CacheUsernameDownloadSettingsCredentialsPtrOutput)
}

func (o CacheUsernameDownloadSettingsCredentialsOutput) BindDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsCredentials) *string { return v.BindDn }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsCredentialsOutput) BindPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsCredentials) *string { return v.BindPassword }).(pulumi.StringPtrOutput)
}

type CacheUsernameDownloadSettingsCredentialsPtrOutput struct{ *pulumi.OutputState }

func (CacheUsernameDownloadSettingsCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUsernameDownloadSettingsCredentials)(nil)).Elem()
}

func (o CacheUsernameDownloadSettingsCredentialsPtrOutput) ToCacheUsernameDownloadSettingsCredentialsPtrOutput() CacheUsernameDownloadSettingsCredentialsPtrOutput {
	return o
}

func (o CacheUsernameDownloadSettingsCredentialsPtrOutput) ToCacheUsernameDownloadSettingsCredentialsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsCredentialsPtrOutput {
	return o
}

func (o CacheUsernameDownloadSettingsCredentialsPtrOutput) Elem() CacheUsernameDownloadSettingsCredentialsOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsCredentials) CacheUsernameDownloadSettingsCredentials {
		if v != nil {
			return *v
		}
		var ret CacheUsernameDownloadSettingsCredentials
		return ret
	}).(CacheUsernameDownloadSettingsCredentialsOutput)
}

func (o CacheUsernameDownloadSettingsCredentialsPtrOutput) BindDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsCredentials) *string {
		if v == nil {
			return nil
		}
		return v.BindDn
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsCredentialsPtrOutput) BindPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsCredentials) *string {
		if v == nil {
			return nil
		}
		return v.BindPassword
	}).(pulumi.StringPtrOutput)
}

type CacheUsernameDownloadSettingsResponse struct {
	AutoDownloadCertificate *bool                                             `pulumi:"autoDownloadCertificate"`
	CaCertificateURI        *string                                           `pulumi:"caCertificateURI"`
	Credentials             *CacheUsernameDownloadSettingsResponseCredentials `pulumi:"credentials"`
	EncryptLdapConnection   *bool                                             `pulumi:"encryptLdapConnection"`
	ExtendedGroups          *bool                                             `pulumi:"extendedGroups"`
	GroupFileURI            *string                                           `pulumi:"groupFileURI"`
	LdapBaseDN              *string                                           `pulumi:"ldapBaseDN"`
	LdapServer              *string                                           `pulumi:"ldapServer"`
	RequireValidCertificate *bool                                             `pulumi:"requireValidCertificate"`
	UserFileURI             *string                                           `pulumi:"userFileURI"`
	UsernameDownloaded      string                                            `pulumi:"usernameDownloaded"`
	UsernameSource          *string                                           `pulumi:"usernameSource"`
}





type CacheUsernameDownloadSettingsResponseInput interface {
	pulumi.Input

	ToCacheUsernameDownloadSettingsResponseOutput() CacheUsernameDownloadSettingsResponseOutput
	ToCacheUsernameDownloadSettingsResponseOutputWithContext(context.Context) CacheUsernameDownloadSettingsResponseOutput
}

type CacheUsernameDownloadSettingsResponseArgs struct {
	AutoDownloadCertificate pulumi.BoolPtrInput                                      `pulumi:"autoDownloadCertificate"`
	CaCertificateURI        pulumi.StringPtrInput                                    `pulumi:"caCertificateURI"`
	Credentials             CacheUsernameDownloadSettingsResponseCredentialsPtrInput `pulumi:"credentials"`
	EncryptLdapConnection   pulumi.BoolPtrInput                                      `pulumi:"encryptLdapConnection"`
	ExtendedGroups          pulumi.BoolPtrInput                                      `pulumi:"extendedGroups"`
	GroupFileURI            pulumi.StringPtrInput                                    `pulumi:"groupFileURI"`
	LdapBaseDN              pulumi.StringPtrInput                                    `pulumi:"ldapBaseDN"`
	LdapServer              pulumi.StringPtrInput                                    `pulumi:"ldapServer"`
	RequireValidCertificate pulumi.BoolPtrInput                                      `pulumi:"requireValidCertificate"`
	UserFileURI             pulumi.StringPtrInput                                    `pulumi:"userFileURI"`
	UsernameDownloaded      pulumi.StringInput                                       `pulumi:"usernameDownloaded"`
	UsernameSource          pulumi.StringPtrInput                                    `pulumi:"usernameSource"`
}

func (CacheUsernameDownloadSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUsernameDownloadSettingsResponse)(nil)).Elem()
}

func (i CacheUsernameDownloadSettingsResponseArgs) ToCacheUsernameDownloadSettingsResponseOutput() CacheUsernameDownloadSettingsResponseOutput {
	return i.ToCacheUsernameDownloadSettingsResponseOutputWithContext(context.Background())
}

func (i CacheUsernameDownloadSettingsResponseArgs) ToCacheUsernameDownloadSettingsResponseOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsResponseOutput)
}

func (i CacheUsernameDownloadSettingsResponseArgs) ToCacheUsernameDownloadSettingsResponsePtrOutput() CacheUsernameDownloadSettingsResponsePtrOutput {
	return i.ToCacheUsernameDownloadSettingsResponsePtrOutputWithContext(context.Background())
}

func (i CacheUsernameDownloadSettingsResponseArgs) ToCacheUsernameDownloadSettingsResponsePtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsResponseOutput).ToCacheUsernameDownloadSettingsResponsePtrOutputWithContext(ctx)
}









type CacheUsernameDownloadSettingsResponsePtrInput interface {
	pulumi.Input

	ToCacheUsernameDownloadSettingsResponsePtrOutput() CacheUsernameDownloadSettingsResponsePtrOutput
	ToCacheUsernameDownloadSettingsResponsePtrOutputWithContext(context.Context) CacheUsernameDownloadSettingsResponsePtrOutput
}

type cacheUsernameDownloadSettingsResponsePtrType CacheUsernameDownloadSettingsResponseArgs

func CacheUsernameDownloadSettingsResponsePtr(v *CacheUsernameDownloadSettingsResponseArgs) CacheUsernameDownloadSettingsResponsePtrInput {
	return (*cacheUsernameDownloadSettingsResponsePtrType)(v)
}

func (*cacheUsernameDownloadSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUsernameDownloadSettingsResponse)(nil)).Elem()
}

func (i *cacheUsernameDownloadSettingsResponsePtrType) ToCacheUsernameDownloadSettingsResponsePtrOutput() CacheUsernameDownloadSettingsResponsePtrOutput {
	return i.ToCacheUsernameDownloadSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *cacheUsernameDownloadSettingsResponsePtrType) ToCacheUsernameDownloadSettingsResponsePtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsResponsePtrOutput)
}

type CacheUsernameDownloadSettingsResponseOutput struct{ *pulumi.OutputState }

func (CacheUsernameDownloadSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUsernameDownloadSettingsResponse)(nil)).Elem()
}

func (o CacheUsernameDownloadSettingsResponseOutput) ToCacheUsernameDownloadSettingsResponseOutput() CacheUsernameDownloadSettingsResponseOutput {
	return o
}

func (o CacheUsernameDownloadSettingsResponseOutput) ToCacheUsernameDownloadSettingsResponseOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponseOutput {
	return o
}

func (o CacheUsernameDownloadSettingsResponseOutput) ToCacheUsernameDownloadSettingsResponsePtrOutput() CacheUsernameDownloadSettingsResponsePtrOutput {
	return o.ToCacheUsernameDownloadSettingsResponsePtrOutputWithContext(context.Background())
}

func (o CacheUsernameDownloadSettingsResponseOutput) ToCacheUsernameDownloadSettingsResponsePtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheUsernameDownloadSettingsResponse) *CacheUsernameDownloadSettingsResponse {
		return &v
	}).(CacheUsernameDownloadSettingsResponsePtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) AutoDownloadCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *bool { return v.AutoDownloadCertificate }).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) CaCertificateURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *string { return v.CaCertificateURI }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) Credentials() CacheUsernameDownloadSettingsResponseCredentialsPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *CacheUsernameDownloadSettingsResponseCredentials {
		return v.Credentials
	}).(CacheUsernameDownloadSettingsResponseCredentialsPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) EncryptLdapConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *bool { return v.EncryptLdapConnection }).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) ExtendedGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *bool { return v.ExtendedGroups }).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) GroupFileURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *string { return v.GroupFileURI }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) LdapBaseDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *string { return v.LdapBaseDN }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) LdapServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *string { return v.LdapServer }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) RequireValidCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *bool { return v.RequireValidCertificate }).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) UserFileURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *string { return v.UserFileURI }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) UsernameDownloaded() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) string { return v.UsernameDownloaded }).(pulumi.StringOutput)
}

func (o CacheUsernameDownloadSettingsResponseOutput) UsernameSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponse) *string { return v.UsernameSource }).(pulumi.StringPtrOutput)
}

type CacheUsernameDownloadSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheUsernameDownloadSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUsernameDownloadSettingsResponse)(nil)).Elem()
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) ToCacheUsernameDownloadSettingsResponsePtrOutput() CacheUsernameDownloadSettingsResponsePtrOutput {
	return o
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) ToCacheUsernameDownloadSettingsResponsePtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponsePtrOutput {
	return o
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) Elem() CacheUsernameDownloadSettingsResponseOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) CacheUsernameDownloadSettingsResponse {
		if v != nil {
			return *v
		}
		var ret CacheUsernameDownloadSettingsResponse
		return ret
	}).(CacheUsernameDownloadSettingsResponseOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) AutoDownloadCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutoDownloadCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) CaCertificateURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.CaCertificateURI
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) Credentials() CacheUsernameDownloadSettingsResponseCredentialsPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *CacheUsernameDownloadSettingsResponseCredentials {
		if v == nil {
			return nil
		}
		return v.Credentials
	}).(CacheUsernameDownloadSettingsResponseCredentialsPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) EncryptLdapConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptLdapConnection
	}).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) ExtendedGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ExtendedGroups
	}).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) GroupFileURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupFileURI
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) LdapBaseDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.LdapBaseDN
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) LdapServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.LdapServer
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) RequireValidCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequireValidCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) UserFileURI() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserFileURI
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) UsernameDownloaded() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UsernameDownloaded
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponsePtrOutput) UsernameSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.UsernameSource
	}).(pulumi.StringPtrOutput)
}

type CacheUsernameDownloadSettingsResponseCredentials struct {
	BindDn       *string `pulumi:"bindDn"`
	BindPassword *string `pulumi:"bindPassword"`
}





type CacheUsernameDownloadSettingsResponseCredentialsInput interface {
	pulumi.Input

	ToCacheUsernameDownloadSettingsResponseCredentialsOutput() CacheUsernameDownloadSettingsResponseCredentialsOutput
	ToCacheUsernameDownloadSettingsResponseCredentialsOutputWithContext(context.Context) CacheUsernameDownloadSettingsResponseCredentialsOutput
}

type CacheUsernameDownloadSettingsResponseCredentialsArgs struct {
	BindDn       pulumi.StringPtrInput `pulumi:"bindDn"`
	BindPassword pulumi.StringPtrInput `pulumi:"bindPassword"`
}

func (CacheUsernameDownloadSettingsResponseCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUsernameDownloadSettingsResponseCredentials)(nil)).Elem()
}

func (i CacheUsernameDownloadSettingsResponseCredentialsArgs) ToCacheUsernameDownloadSettingsResponseCredentialsOutput() CacheUsernameDownloadSettingsResponseCredentialsOutput {
	return i.ToCacheUsernameDownloadSettingsResponseCredentialsOutputWithContext(context.Background())
}

func (i CacheUsernameDownloadSettingsResponseCredentialsArgs) ToCacheUsernameDownloadSettingsResponseCredentialsOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponseCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsResponseCredentialsOutput)
}

func (i CacheUsernameDownloadSettingsResponseCredentialsArgs) ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutput() CacheUsernameDownloadSettingsResponseCredentialsPtrOutput {
	return i.ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutputWithContext(context.Background())
}

func (i CacheUsernameDownloadSettingsResponseCredentialsArgs) ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponseCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsResponseCredentialsOutput).ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutputWithContext(ctx)
}









type CacheUsernameDownloadSettingsResponseCredentialsPtrInput interface {
	pulumi.Input

	ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutput() CacheUsernameDownloadSettingsResponseCredentialsPtrOutput
	ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutputWithContext(context.Context) CacheUsernameDownloadSettingsResponseCredentialsPtrOutput
}

type cacheUsernameDownloadSettingsResponseCredentialsPtrType CacheUsernameDownloadSettingsResponseCredentialsArgs

func CacheUsernameDownloadSettingsResponseCredentialsPtr(v *CacheUsernameDownloadSettingsResponseCredentialsArgs) CacheUsernameDownloadSettingsResponseCredentialsPtrInput {
	return (*cacheUsernameDownloadSettingsResponseCredentialsPtrType)(v)
}

func (*cacheUsernameDownloadSettingsResponseCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUsernameDownloadSettingsResponseCredentials)(nil)).Elem()
}

func (i *cacheUsernameDownloadSettingsResponseCredentialsPtrType) ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutput() CacheUsernameDownloadSettingsResponseCredentialsPtrOutput {
	return i.ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutputWithContext(context.Background())
}

func (i *cacheUsernameDownloadSettingsResponseCredentialsPtrType) ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponseCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUsernameDownloadSettingsResponseCredentialsPtrOutput)
}

type CacheUsernameDownloadSettingsResponseCredentialsOutput struct{ *pulumi.OutputState }

func (CacheUsernameDownloadSettingsResponseCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUsernameDownloadSettingsResponseCredentials)(nil)).Elem()
}

func (o CacheUsernameDownloadSettingsResponseCredentialsOutput) ToCacheUsernameDownloadSettingsResponseCredentialsOutput() CacheUsernameDownloadSettingsResponseCredentialsOutput {
	return o
}

func (o CacheUsernameDownloadSettingsResponseCredentialsOutput) ToCacheUsernameDownloadSettingsResponseCredentialsOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponseCredentialsOutput {
	return o
}

func (o CacheUsernameDownloadSettingsResponseCredentialsOutput) ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutput() CacheUsernameDownloadSettingsResponseCredentialsPtrOutput {
	return o.ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutputWithContext(context.Background())
}

func (o CacheUsernameDownloadSettingsResponseCredentialsOutput) ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponseCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheUsernameDownloadSettingsResponseCredentials) *CacheUsernameDownloadSettingsResponseCredentials {
		return &v
	}).(CacheUsernameDownloadSettingsResponseCredentialsPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseCredentialsOutput) BindDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponseCredentials) *string { return v.BindDn }).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseCredentialsOutput) BindPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheUsernameDownloadSettingsResponseCredentials) *string { return v.BindPassword }).(pulumi.StringPtrOutput)
}

type CacheUsernameDownloadSettingsResponseCredentialsPtrOutput struct{ *pulumi.OutputState }

func (CacheUsernameDownloadSettingsResponseCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUsernameDownloadSettingsResponseCredentials)(nil)).Elem()
}

func (o CacheUsernameDownloadSettingsResponseCredentialsPtrOutput) ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutput() CacheUsernameDownloadSettingsResponseCredentialsPtrOutput {
	return o
}

func (o CacheUsernameDownloadSettingsResponseCredentialsPtrOutput) ToCacheUsernameDownloadSettingsResponseCredentialsPtrOutputWithContext(ctx context.Context) CacheUsernameDownloadSettingsResponseCredentialsPtrOutput {
	return o
}

func (o CacheUsernameDownloadSettingsResponseCredentialsPtrOutput) Elem() CacheUsernameDownloadSettingsResponseCredentialsOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponseCredentials) CacheUsernameDownloadSettingsResponseCredentials {
		if v != nil {
			return *v
		}
		var ret CacheUsernameDownloadSettingsResponseCredentials
		return ret
	}).(CacheUsernameDownloadSettingsResponseCredentialsOutput)
}

func (o CacheUsernameDownloadSettingsResponseCredentialsPtrOutput) BindDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponseCredentials) *string {
		if v == nil {
			return nil
		}
		return v.BindDn
	}).(pulumi.StringPtrOutput)
}

func (o CacheUsernameDownloadSettingsResponseCredentialsPtrOutput) BindPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUsernameDownloadSettingsResponseCredentials) *string {
		if v == nil {
			return nil
		}
		return v.BindPassword
	}).(pulumi.StringPtrOutput)
}

type ClfsTarget struct {
	Target *string `pulumi:"target"`
}





type ClfsTargetInput interface {
	pulumi.Input

	ToClfsTargetOutput() ClfsTargetOutput
	ToClfsTargetOutputWithContext(context.Context) ClfsTargetOutput
}

type ClfsTargetArgs struct {
	Target pulumi.StringPtrInput `pulumi:"target"`
}

func (ClfsTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClfsTarget)(nil)).Elem()
}

func (i ClfsTargetArgs) ToClfsTargetOutput() ClfsTargetOutput {
	return i.ToClfsTargetOutputWithContext(context.Background())
}

func (i ClfsTargetArgs) ToClfsTargetOutputWithContext(ctx context.Context) ClfsTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetOutput)
}

func (i ClfsTargetArgs) ToClfsTargetPtrOutput() ClfsTargetPtrOutput {
	return i.ToClfsTargetPtrOutputWithContext(context.Background())
}

func (i ClfsTargetArgs) ToClfsTargetPtrOutputWithContext(ctx context.Context) ClfsTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetOutput).ToClfsTargetPtrOutputWithContext(ctx)
}









type ClfsTargetPtrInput interface {
	pulumi.Input

	ToClfsTargetPtrOutput() ClfsTargetPtrOutput
	ToClfsTargetPtrOutputWithContext(context.Context) ClfsTargetPtrOutput
}

type clfsTargetPtrType ClfsTargetArgs

func ClfsTargetPtr(v *ClfsTargetArgs) ClfsTargetPtrInput {
	return (*clfsTargetPtrType)(v)
}

func (*clfsTargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClfsTarget)(nil)).Elem()
}

func (i *clfsTargetPtrType) ToClfsTargetPtrOutput() ClfsTargetPtrOutput {
	return i.ToClfsTargetPtrOutputWithContext(context.Background())
}

func (i *clfsTargetPtrType) ToClfsTargetPtrOutputWithContext(ctx context.Context) ClfsTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetPtrOutput)
}

type ClfsTargetOutput struct{ *pulumi.OutputState }

func (ClfsTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClfsTarget)(nil)).Elem()
}

func (o ClfsTargetOutput) ToClfsTargetOutput() ClfsTargetOutput {
	return o
}

func (o ClfsTargetOutput) ToClfsTargetOutputWithContext(ctx context.Context) ClfsTargetOutput {
	return o
}

func (o ClfsTargetOutput) ToClfsTargetPtrOutput() ClfsTargetPtrOutput {
	return o.ToClfsTargetPtrOutputWithContext(context.Background())
}

func (o ClfsTargetOutput) ToClfsTargetPtrOutputWithContext(ctx context.Context) ClfsTargetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClfsTarget) *ClfsTarget {
		return &v
	}).(ClfsTargetPtrOutput)
}

func (o ClfsTargetOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClfsTarget) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ClfsTargetPtrOutput struct{ *pulumi.OutputState }

func (ClfsTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClfsTarget)(nil)).Elem()
}

func (o ClfsTargetPtrOutput) ToClfsTargetPtrOutput() ClfsTargetPtrOutput {
	return o
}

func (o ClfsTargetPtrOutput) ToClfsTargetPtrOutputWithContext(ctx context.Context) ClfsTargetPtrOutput {
	return o
}

func (o ClfsTargetPtrOutput) Elem() ClfsTargetOutput {
	return o.ApplyT(func(v *ClfsTarget) ClfsTarget {
		if v != nil {
			return *v
		}
		var ret ClfsTarget
		return ret
	}).(ClfsTargetOutput)
}

func (o ClfsTargetPtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClfsTarget) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type ClfsTargetResponse struct {
	Target *string `pulumi:"target"`
}





type ClfsTargetResponseInput interface {
	pulumi.Input

	ToClfsTargetResponseOutput() ClfsTargetResponseOutput
	ToClfsTargetResponseOutputWithContext(context.Context) ClfsTargetResponseOutput
}

type ClfsTargetResponseArgs struct {
	Target pulumi.StringPtrInput `pulumi:"target"`
}

func (ClfsTargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClfsTargetResponse)(nil)).Elem()
}

func (i ClfsTargetResponseArgs) ToClfsTargetResponseOutput() ClfsTargetResponseOutput {
	return i.ToClfsTargetResponseOutputWithContext(context.Background())
}

func (i ClfsTargetResponseArgs) ToClfsTargetResponseOutputWithContext(ctx context.Context) ClfsTargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetResponseOutput)
}

func (i ClfsTargetResponseArgs) ToClfsTargetResponsePtrOutput() ClfsTargetResponsePtrOutput {
	return i.ToClfsTargetResponsePtrOutputWithContext(context.Background())
}

func (i ClfsTargetResponseArgs) ToClfsTargetResponsePtrOutputWithContext(ctx context.Context) ClfsTargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetResponseOutput).ToClfsTargetResponsePtrOutputWithContext(ctx)
}









type ClfsTargetResponsePtrInput interface {
	pulumi.Input

	ToClfsTargetResponsePtrOutput() ClfsTargetResponsePtrOutput
	ToClfsTargetResponsePtrOutputWithContext(context.Context) ClfsTargetResponsePtrOutput
}

type clfsTargetResponsePtrType ClfsTargetResponseArgs

func ClfsTargetResponsePtr(v *ClfsTargetResponseArgs) ClfsTargetResponsePtrInput {
	return (*clfsTargetResponsePtrType)(v)
}

func (*clfsTargetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClfsTargetResponse)(nil)).Elem()
}

func (i *clfsTargetResponsePtrType) ToClfsTargetResponsePtrOutput() ClfsTargetResponsePtrOutput {
	return i.ToClfsTargetResponsePtrOutputWithContext(context.Background())
}

func (i *clfsTargetResponsePtrType) ToClfsTargetResponsePtrOutputWithContext(ctx context.Context) ClfsTargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetResponsePtrOutput)
}

type ClfsTargetResponseOutput struct{ *pulumi.OutputState }

func (ClfsTargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClfsTargetResponse)(nil)).Elem()
}

func (o ClfsTargetResponseOutput) ToClfsTargetResponseOutput() ClfsTargetResponseOutput {
	return o
}

func (o ClfsTargetResponseOutput) ToClfsTargetResponseOutputWithContext(ctx context.Context) ClfsTargetResponseOutput {
	return o
}

func (o ClfsTargetResponseOutput) ToClfsTargetResponsePtrOutput() ClfsTargetResponsePtrOutput {
	return o.ToClfsTargetResponsePtrOutputWithContext(context.Background())
}

func (o ClfsTargetResponseOutput) ToClfsTargetResponsePtrOutputWithContext(ctx context.Context) ClfsTargetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClfsTargetResponse) *ClfsTargetResponse {
		return &v
	}).(ClfsTargetResponsePtrOutput)
}

func (o ClfsTargetResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClfsTargetResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ClfsTargetResponsePtrOutput struct{ *pulumi.OutputState }

func (ClfsTargetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClfsTargetResponse)(nil)).Elem()
}

func (o ClfsTargetResponsePtrOutput) ToClfsTargetResponsePtrOutput() ClfsTargetResponsePtrOutput {
	return o
}

func (o ClfsTargetResponsePtrOutput) ToClfsTargetResponsePtrOutputWithContext(ctx context.Context) ClfsTargetResponsePtrOutput {
	return o
}

func (o ClfsTargetResponsePtrOutput) Elem() ClfsTargetResponseOutput {
	return o.ApplyT(func(v *ClfsTargetResponse) ClfsTargetResponse {
		if v != nil {
			return *v
		}
		var ret ClfsTargetResponse
		return ret
	}).(ClfsTargetResponseOutput)
}

func (o ClfsTargetResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClfsTargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type ConditionResponse struct {
	Message   string `pulumi:"message"`
	Timestamp string `pulumi:"timestamp"`
}





type ConditionResponseInput interface {
	pulumi.Input

	ToConditionResponseOutput() ConditionResponseOutput
	ToConditionResponseOutputWithContext(context.Context) ConditionResponseOutput
}

type ConditionResponseArgs struct {
	Message   pulumi.StringInput `pulumi:"message"`
	Timestamp pulumi.StringInput `pulumi:"timestamp"`
}

func (ConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionResponse)(nil)).Elem()
}

func (i ConditionResponseArgs) ToConditionResponseOutput() ConditionResponseOutput {
	return i.ToConditionResponseOutputWithContext(context.Background())
}

func (i ConditionResponseArgs) ToConditionResponseOutputWithContext(ctx context.Context) ConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponseOutput)
}





type ConditionResponseArrayInput interface {
	pulumi.Input

	ToConditionResponseArrayOutput() ConditionResponseArrayOutput
	ToConditionResponseArrayOutputWithContext(context.Context) ConditionResponseArrayOutput
}

type ConditionResponseArray []ConditionResponseInput

func (ConditionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConditionResponse)(nil)).Elem()
}

func (i ConditionResponseArray) ToConditionResponseArrayOutput() ConditionResponseArrayOutput {
	return i.ToConditionResponseArrayOutputWithContext(context.Background())
}

func (i ConditionResponseArray) ToConditionResponseArrayOutputWithContext(ctx context.Context) ConditionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponseArrayOutput)
}

type ConditionResponseOutput struct{ *pulumi.OutputState }

func (ConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionResponse)(nil)).Elem()
}

func (o ConditionResponseOutput) ToConditionResponseOutput() ConditionResponseOutput {
	return o
}

func (o ConditionResponseOutput) ToConditionResponseOutputWithContext(ctx context.Context) ConditionResponseOutput {
	return o
}

func (o ConditionResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ConditionResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ConditionResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ConditionResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

type ConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (ConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConditionResponse)(nil)).Elem()
}

func (o ConditionResponseArrayOutput) ToConditionResponseArrayOutput() ConditionResponseArrayOutput {
	return o
}

func (o ConditionResponseArrayOutput) ToConditionResponseArrayOutputWithContext(ctx context.Context) ConditionResponseArrayOutput {
	return o
}

func (o ConditionResponseArrayOutput) Index(i pulumi.IntInput) ConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConditionResponse {
		return vs[0].([]ConditionResponse)[vs[1].(int)]
	}).(ConditionResponseOutput)
}

type KeyVaultKeyReference struct {
	KeyUrl      string                          `pulumi:"keyUrl"`
	SourceVault KeyVaultKeyReferenceSourceVault `pulumi:"sourceVault"`
}





type KeyVaultKeyReferenceInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput
	ToKeyVaultKeyReferenceOutputWithContext(context.Context) KeyVaultKeyReferenceOutput
}

type KeyVaultKeyReferenceArgs struct {
	KeyUrl      pulumi.StringInput                   `pulumi:"keyUrl"`
	SourceVault KeyVaultKeyReferenceSourceVaultInput `pulumi:"sourceVault"`
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

func (o KeyVaultKeyReferenceOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceOutput) SourceVault() KeyVaultKeyReferenceSourceVaultOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) KeyVaultKeyReferenceSourceVault { return v.SourceVault }).(KeyVaultKeyReferenceSourceVaultOutput)
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

func (o KeyVaultKeyReferencePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferencePtrOutput) SourceVault() KeyVaultKeyReferenceSourceVaultPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *KeyVaultKeyReferenceSourceVault {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(KeyVaultKeyReferenceSourceVaultPtrOutput)
}

type KeyVaultKeyReferenceResponse struct {
	KeyUrl      string                                  `pulumi:"keyUrl"`
	SourceVault KeyVaultKeyReferenceResponseSourceVault `pulumi:"sourceVault"`
}





type KeyVaultKeyReferenceResponseInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceResponseOutput() KeyVaultKeyReferenceResponseOutput
	ToKeyVaultKeyReferenceResponseOutputWithContext(context.Context) KeyVaultKeyReferenceResponseOutput
}

type KeyVaultKeyReferenceResponseArgs struct {
	KeyUrl      pulumi.StringInput                           `pulumi:"keyUrl"`
	SourceVault KeyVaultKeyReferenceResponseSourceVaultInput `pulumi:"sourceVault"`
}

func (KeyVaultKeyReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (i KeyVaultKeyReferenceResponseArgs) ToKeyVaultKeyReferenceResponseOutput() KeyVaultKeyReferenceResponseOutput {
	return i.ToKeyVaultKeyReferenceResponseOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceResponseArgs) ToKeyVaultKeyReferenceResponseOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceResponseOutput)
}

func (i KeyVaultKeyReferenceResponseArgs) ToKeyVaultKeyReferenceResponsePtrOutput() KeyVaultKeyReferenceResponsePtrOutput {
	return i.ToKeyVaultKeyReferenceResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceResponseArgs) ToKeyVaultKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceResponseOutput).ToKeyVaultKeyReferenceResponsePtrOutputWithContext(ctx)
}









type KeyVaultKeyReferenceResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceResponsePtrOutput() KeyVaultKeyReferenceResponsePtrOutput
	ToKeyVaultKeyReferenceResponsePtrOutputWithContext(context.Context) KeyVaultKeyReferenceResponsePtrOutput
}

type keyVaultKeyReferenceResponsePtrType KeyVaultKeyReferenceResponseArgs

func KeyVaultKeyReferenceResponsePtr(v *KeyVaultKeyReferenceResponseArgs) KeyVaultKeyReferenceResponsePtrInput {
	return (*keyVaultKeyReferenceResponsePtrType)(v)
}

func (*keyVaultKeyReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (i *keyVaultKeyReferenceResponsePtrType) ToKeyVaultKeyReferenceResponsePtrOutput() KeyVaultKeyReferenceResponsePtrOutput {
	return i.ToKeyVaultKeyReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyReferenceResponsePtrType) ToKeyVaultKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceResponsePtrOutput)
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

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponsePtrOutput() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ToKeyVaultKeyReferenceResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKeyReferenceResponse) *KeyVaultKeyReferenceResponse {
		return &v
	}).(KeyVaultKeyReferenceResponsePtrOutput)
}

func (o KeyVaultKeyReferenceResponseOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceResponseOutput) SourceVault() KeyVaultKeyReferenceResponseSourceVaultOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) KeyVaultKeyReferenceResponseSourceVault { return v.SourceVault }).(KeyVaultKeyReferenceResponseSourceVaultOutput)
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

func (o KeyVaultKeyReferenceResponsePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) SourceVault() KeyVaultKeyReferenceResponseSourceVaultPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *KeyVaultKeyReferenceResponseSourceVault {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(KeyVaultKeyReferenceResponseSourceVaultPtrOutput)
}

type KeyVaultKeyReferenceResponseSourceVault struct {
	Id *string `pulumi:"id"`
}





type KeyVaultKeyReferenceResponseSourceVaultInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceResponseSourceVaultOutput() KeyVaultKeyReferenceResponseSourceVaultOutput
	ToKeyVaultKeyReferenceResponseSourceVaultOutputWithContext(context.Context) KeyVaultKeyReferenceResponseSourceVaultOutput
}

type KeyVaultKeyReferenceResponseSourceVaultArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (KeyVaultKeyReferenceResponseSourceVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceResponseSourceVault)(nil)).Elem()
}

func (i KeyVaultKeyReferenceResponseSourceVaultArgs) ToKeyVaultKeyReferenceResponseSourceVaultOutput() KeyVaultKeyReferenceResponseSourceVaultOutput {
	return i.ToKeyVaultKeyReferenceResponseSourceVaultOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceResponseSourceVaultArgs) ToKeyVaultKeyReferenceResponseSourceVaultOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseSourceVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceResponseSourceVaultOutput)
}

func (i KeyVaultKeyReferenceResponseSourceVaultArgs) ToKeyVaultKeyReferenceResponseSourceVaultPtrOutput() KeyVaultKeyReferenceResponseSourceVaultPtrOutput {
	return i.ToKeyVaultKeyReferenceResponseSourceVaultPtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceResponseSourceVaultArgs) ToKeyVaultKeyReferenceResponseSourceVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseSourceVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceResponseSourceVaultOutput).ToKeyVaultKeyReferenceResponseSourceVaultPtrOutputWithContext(ctx)
}









type KeyVaultKeyReferenceResponseSourceVaultPtrInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceResponseSourceVaultPtrOutput() KeyVaultKeyReferenceResponseSourceVaultPtrOutput
	ToKeyVaultKeyReferenceResponseSourceVaultPtrOutputWithContext(context.Context) KeyVaultKeyReferenceResponseSourceVaultPtrOutput
}

type keyVaultKeyReferenceResponseSourceVaultPtrType KeyVaultKeyReferenceResponseSourceVaultArgs

func KeyVaultKeyReferenceResponseSourceVaultPtr(v *KeyVaultKeyReferenceResponseSourceVaultArgs) KeyVaultKeyReferenceResponseSourceVaultPtrInput {
	return (*keyVaultKeyReferenceResponseSourceVaultPtrType)(v)
}

func (*keyVaultKeyReferenceResponseSourceVaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceResponseSourceVault)(nil)).Elem()
}

func (i *keyVaultKeyReferenceResponseSourceVaultPtrType) ToKeyVaultKeyReferenceResponseSourceVaultPtrOutput() KeyVaultKeyReferenceResponseSourceVaultPtrOutput {
	return i.ToKeyVaultKeyReferenceResponseSourceVaultPtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyReferenceResponseSourceVaultPtrType) ToKeyVaultKeyReferenceResponseSourceVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseSourceVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceResponseSourceVaultPtrOutput)
}

type KeyVaultKeyReferenceResponseSourceVaultOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseSourceVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceResponseSourceVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseSourceVaultOutput) ToKeyVaultKeyReferenceResponseSourceVaultOutput() KeyVaultKeyReferenceResponseSourceVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseSourceVaultOutput) ToKeyVaultKeyReferenceResponseSourceVaultOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseSourceVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseSourceVaultOutput) ToKeyVaultKeyReferenceResponseSourceVaultPtrOutput() KeyVaultKeyReferenceResponseSourceVaultPtrOutput {
	return o.ToKeyVaultKeyReferenceResponseSourceVaultPtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyReferenceResponseSourceVaultOutput) ToKeyVaultKeyReferenceResponseSourceVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseSourceVaultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKeyReferenceResponseSourceVault) *KeyVaultKeyReferenceResponseSourceVault {
		return &v
	}).(KeyVaultKeyReferenceResponseSourceVaultPtrOutput)
}

func (o KeyVaultKeyReferenceResponseSourceVaultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponseSourceVault) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceResponseSourceVaultPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseSourceVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceResponseSourceVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseSourceVaultPtrOutput) ToKeyVaultKeyReferenceResponseSourceVaultPtrOutput() KeyVaultKeyReferenceResponseSourceVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseSourceVaultPtrOutput) ToKeyVaultKeyReferenceResponseSourceVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseSourceVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseSourceVaultPtrOutput) Elem() KeyVaultKeyReferenceResponseSourceVaultOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponseSourceVault) KeyVaultKeyReferenceResponseSourceVault {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceResponseSourceVault
		return ret
	}).(KeyVaultKeyReferenceResponseSourceVaultOutput)
}

func (o KeyVaultKeyReferenceResponseSourceVaultPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponseSourceVault) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceSourceVault struct {
	Id *string `pulumi:"id"`
}





type KeyVaultKeyReferenceSourceVaultInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceSourceVaultOutput() KeyVaultKeyReferenceSourceVaultOutput
	ToKeyVaultKeyReferenceSourceVaultOutputWithContext(context.Context) KeyVaultKeyReferenceSourceVaultOutput
}

type KeyVaultKeyReferenceSourceVaultArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (KeyVaultKeyReferenceSourceVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceSourceVault)(nil)).Elem()
}

func (i KeyVaultKeyReferenceSourceVaultArgs) ToKeyVaultKeyReferenceSourceVaultOutput() KeyVaultKeyReferenceSourceVaultOutput {
	return i.ToKeyVaultKeyReferenceSourceVaultOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceSourceVaultArgs) ToKeyVaultKeyReferenceSourceVaultOutputWithContext(ctx context.Context) KeyVaultKeyReferenceSourceVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceSourceVaultOutput)
}

func (i KeyVaultKeyReferenceSourceVaultArgs) ToKeyVaultKeyReferenceSourceVaultPtrOutput() KeyVaultKeyReferenceSourceVaultPtrOutput {
	return i.ToKeyVaultKeyReferenceSourceVaultPtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceSourceVaultArgs) ToKeyVaultKeyReferenceSourceVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceSourceVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceSourceVaultOutput).ToKeyVaultKeyReferenceSourceVaultPtrOutputWithContext(ctx)
}









type KeyVaultKeyReferenceSourceVaultPtrInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceSourceVaultPtrOutput() KeyVaultKeyReferenceSourceVaultPtrOutput
	ToKeyVaultKeyReferenceSourceVaultPtrOutputWithContext(context.Context) KeyVaultKeyReferenceSourceVaultPtrOutput
}

type keyVaultKeyReferenceSourceVaultPtrType KeyVaultKeyReferenceSourceVaultArgs

func KeyVaultKeyReferenceSourceVaultPtr(v *KeyVaultKeyReferenceSourceVaultArgs) KeyVaultKeyReferenceSourceVaultPtrInput {
	return (*keyVaultKeyReferenceSourceVaultPtrType)(v)
}

func (*keyVaultKeyReferenceSourceVaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceSourceVault)(nil)).Elem()
}

func (i *keyVaultKeyReferenceSourceVaultPtrType) ToKeyVaultKeyReferenceSourceVaultPtrOutput() KeyVaultKeyReferenceSourceVaultPtrOutput {
	return i.ToKeyVaultKeyReferenceSourceVaultPtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyReferenceSourceVaultPtrType) ToKeyVaultKeyReferenceSourceVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceSourceVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceSourceVaultPtrOutput)
}

type KeyVaultKeyReferenceSourceVaultOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceSourceVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceSourceVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceSourceVaultOutput) ToKeyVaultKeyReferenceSourceVaultOutput() KeyVaultKeyReferenceSourceVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceSourceVaultOutput) ToKeyVaultKeyReferenceSourceVaultOutputWithContext(ctx context.Context) KeyVaultKeyReferenceSourceVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceSourceVaultOutput) ToKeyVaultKeyReferenceSourceVaultPtrOutput() KeyVaultKeyReferenceSourceVaultPtrOutput {
	return o.ToKeyVaultKeyReferenceSourceVaultPtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyReferenceSourceVaultOutput) ToKeyVaultKeyReferenceSourceVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceSourceVaultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKeyReferenceSourceVault) *KeyVaultKeyReferenceSourceVault {
		return &v
	}).(KeyVaultKeyReferenceSourceVaultPtrOutput)
}

func (o KeyVaultKeyReferenceSourceVaultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceSourceVault) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceSourceVaultPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceSourceVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceSourceVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceSourceVaultPtrOutput) ToKeyVaultKeyReferenceSourceVaultPtrOutput() KeyVaultKeyReferenceSourceVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceSourceVaultPtrOutput) ToKeyVaultKeyReferenceSourceVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceSourceVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceSourceVaultPtrOutput) Elem() KeyVaultKeyReferenceSourceVaultOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceSourceVault) KeyVaultKeyReferenceSourceVault {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceSourceVault
		return ret
	}).(KeyVaultKeyReferenceSourceVaultOutput)
}

func (o KeyVaultKeyReferenceSourceVaultPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceSourceVault) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type NamespaceJunction struct {
	NamespacePath   *string `pulumi:"namespacePath"`
	NfsAccessPolicy *string `pulumi:"nfsAccessPolicy"`
	NfsExport       *string `pulumi:"nfsExport"`
	TargetPath      *string `pulumi:"targetPath"`
}





type NamespaceJunctionInput interface {
	pulumi.Input

	ToNamespaceJunctionOutput() NamespaceJunctionOutput
	ToNamespaceJunctionOutputWithContext(context.Context) NamespaceJunctionOutput
}

type NamespaceJunctionArgs struct {
	NamespacePath   pulumi.StringPtrInput `pulumi:"namespacePath"`
	NfsAccessPolicy pulumi.StringPtrInput `pulumi:"nfsAccessPolicy"`
	NfsExport       pulumi.StringPtrInput `pulumi:"nfsExport"`
	TargetPath      pulumi.StringPtrInput `pulumi:"targetPath"`
}

func (NamespaceJunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceJunction)(nil)).Elem()
}

func (i NamespaceJunctionArgs) ToNamespaceJunctionOutput() NamespaceJunctionOutput {
	return i.ToNamespaceJunctionOutputWithContext(context.Background())
}

func (i NamespaceJunctionArgs) ToNamespaceJunctionOutputWithContext(ctx context.Context) NamespaceJunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceJunctionOutput)
}





type NamespaceJunctionArrayInput interface {
	pulumi.Input

	ToNamespaceJunctionArrayOutput() NamespaceJunctionArrayOutput
	ToNamespaceJunctionArrayOutputWithContext(context.Context) NamespaceJunctionArrayOutput
}

type NamespaceJunctionArray []NamespaceJunctionInput

func (NamespaceJunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceJunction)(nil)).Elem()
}

func (i NamespaceJunctionArray) ToNamespaceJunctionArrayOutput() NamespaceJunctionArrayOutput {
	return i.ToNamespaceJunctionArrayOutputWithContext(context.Background())
}

func (i NamespaceJunctionArray) ToNamespaceJunctionArrayOutputWithContext(ctx context.Context) NamespaceJunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceJunctionArrayOutput)
}

type NamespaceJunctionOutput struct{ *pulumi.OutputState }

func (NamespaceJunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceJunction)(nil)).Elem()
}

func (o NamespaceJunctionOutput) ToNamespaceJunctionOutput() NamespaceJunctionOutput {
	return o
}

func (o NamespaceJunctionOutput) ToNamespaceJunctionOutputWithContext(ctx context.Context) NamespaceJunctionOutput {
	return o
}

func (o NamespaceJunctionOutput) NamespacePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunction) *string { return v.NamespacePath }).(pulumi.StringPtrOutput)
}

func (o NamespaceJunctionOutput) NfsAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunction) *string { return v.NfsAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o NamespaceJunctionOutput) NfsExport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunction) *string { return v.NfsExport }).(pulumi.StringPtrOutput)
}

func (o NamespaceJunctionOutput) TargetPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunction) *string { return v.TargetPath }).(pulumi.StringPtrOutput)
}

type NamespaceJunctionArrayOutput struct{ *pulumi.OutputState }

func (NamespaceJunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceJunction)(nil)).Elem()
}

func (o NamespaceJunctionArrayOutput) ToNamespaceJunctionArrayOutput() NamespaceJunctionArrayOutput {
	return o
}

func (o NamespaceJunctionArrayOutput) ToNamespaceJunctionArrayOutputWithContext(ctx context.Context) NamespaceJunctionArrayOutput {
	return o
}

func (o NamespaceJunctionArrayOutput) Index(i pulumi.IntInput) NamespaceJunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamespaceJunction {
		return vs[0].([]NamespaceJunction)[vs[1].(int)]
	}).(NamespaceJunctionOutput)
}

type NamespaceJunctionResponse struct {
	NamespacePath   *string `pulumi:"namespacePath"`
	NfsAccessPolicy *string `pulumi:"nfsAccessPolicy"`
	NfsExport       *string `pulumi:"nfsExport"`
	TargetPath      *string `pulumi:"targetPath"`
}





type NamespaceJunctionResponseInput interface {
	pulumi.Input

	ToNamespaceJunctionResponseOutput() NamespaceJunctionResponseOutput
	ToNamespaceJunctionResponseOutputWithContext(context.Context) NamespaceJunctionResponseOutput
}

type NamespaceJunctionResponseArgs struct {
	NamespacePath   pulumi.StringPtrInput `pulumi:"namespacePath"`
	NfsAccessPolicy pulumi.StringPtrInput `pulumi:"nfsAccessPolicy"`
	NfsExport       pulumi.StringPtrInput `pulumi:"nfsExport"`
	TargetPath      pulumi.StringPtrInput `pulumi:"targetPath"`
}

func (NamespaceJunctionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceJunctionResponse)(nil)).Elem()
}

func (i NamespaceJunctionResponseArgs) ToNamespaceJunctionResponseOutput() NamespaceJunctionResponseOutput {
	return i.ToNamespaceJunctionResponseOutputWithContext(context.Background())
}

func (i NamespaceJunctionResponseArgs) ToNamespaceJunctionResponseOutputWithContext(ctx context.Context) NamespaceJunctionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceJunctionResponseOutput)
}





type NamespaceJunctionResponseArrayInput interface {
	pulumi.Input

	ToNamespaceJunctionResponseArrayOutput() NamespaceJunctionResponseArrayOutput
	ToNamespaceJunctionResponseArrayOutputWithContext(context.Context) NamespaceJunctionResponseArrayOutput
}

type NamespaceJunctionResponseArray []NamespaceJunctionResponseInput

func (NamespaceJunctionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceJunctionResponse)(nil)).Elem()
}

func (i NamespaceJunctionResponseArray) ToNamespaceJunctionResponseArrayOutput() NamespaceJunctionResponseArrayOutput {
	return i.ToNamespaceJunctionResponseArrayOutputWithContext(context.Background())
}

func (i NamespaceJunctionResponseArray) ToNamespaceJunctionResponseArrayOutputWithContext(ctx context.Context) NamespaceJunctionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceJunctionResponseArrayOutput)
}

type NamespaceJunctionResponseOutput struct{ *pulumi.OutputState }

func (NamespaceJunctionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceJunctionResponse)(nil)).Elem()
}

func (o NamespaceJunctionResponseOutput) ToNamespaceJunctionResponseOutput() NamespaceJunctionResponseOutput {
	return o
}

func (o NamespaceJunctionResponseOutput) ToNamespaceJunctionResponseOutputWithContext(ctx context.Context) NamespaceJunctionResponseOutput {
	return o
}

func (o NamespaceJunctionResponseOutput) NamespacePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunctionResponse) *string { return v.NamespacePath }).(pulumi.StringPtrOutput)
}

func (o NamespaceJunctionResponseOutput) NfsAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunctionResponse) *string { return v.NfsAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o NamespaceJunctionResponseOutput) NfsExport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunctionResponse) *string { return v.NfsExport }).(pulumi.StringPtrOutput)
}

func (o NamespaceJunctionResponseOutput) TargetPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunctionResponse) *string { return v.TargetPath }).(pulumi.StringPtrOutput)
}

type NamespaceJunctionResponseArrayOutput struct{ *pulumi.OutputState }

func (NamespaceJunctionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceJunctionResponse)(nil)).Elem()
}

func (o NamespaceJunctionResponseArrayOutput) ToNamespaceJunctionResponseArrayOutput() NamespaceJunctionResponseArrayOutput {
	return o
}

func (o NamespaceJunctionResponseArrayOutput) ToNamespaceJunctionResponseArrayOutputWithContext(ctx context.Context) NamespaceJunctionResponseArrayOutput {
	return o
}

func (o NamespaceJunctionResponseArrayOutput) Index(i pulumi.IntInput) NamespaceJunctionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamespaceJunctionResponse {
		return vs[0].([]NamespaceJunctionResponse)[vs[1].(int)]
	}).(NamespaceJunctionResponseOutput)
}

type Nfs3Target struct {
	Target     *string `pulumi:"target"`
	UsageModel *string `pulumi:"usageModel"`
}





type Nfs3TargetInput interface {
	pulumi.Input

	ToNfs3TargetOutput() Nfs3TargetOutput
	ToNfs3TargetOutputWithContext(context.Context) Nfs3TargetOutput
}

type Nfs3TargetArgs struct {
	Target     pulumi.StringPtrInput `pulumi:"target"`
	UsageModel pulumi.StringPtrInput `pulumi:"usageModel"`
}

func (Nfs3TargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Nfs3Target)(nil)).Elem()
}

func (i Nfs3TargetArgs) ToNfs3TargetOutput() Nfs3TargetOutput {
	return i.ToNfs3TargetOutputWithContext(context.Background())
}

func (i Nfs3TargetArgs) ToNfs3TargetOutputWithContext(ctx context.Context) Nfs3TargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetOutput)
}

func (i Nfs3TargetArgs) ToNfs3TargetPtrOutput() Nfs3TargetPtrOutput {
	return i.ToNfs3TargetPtrOutputWithContext(context.Background())
}

func (i Nfs3TargetArgs) ToNfs3TargetPtrOutputWithContext(ctx context.Context) Nfs3TargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetOutput).ToNfs3TargetPtrOutputWithContext(ctx)
}









type Nfs3TargetPtrInput interface {
	pulumi.Input

	ToNfs3TargetPtrOutput() Nfs3TargetPtrOutput
	ToNfs3TargetPtrOutputWithContext(context.Context) Nfs3TargetPtrOutput
}

type nfs3TargetPtrType Nfs3TargetArgs

func Nfs3TargetPtr(v *Nfs3TargetArgs) Nfs3TargetPtrInput {
	return (*nfs3TargetPtrType)(v)
}

func (*nfs3TargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Nfs3Target)(nil)).Elem()
}

func (i *nfs3TargetPtrType) ToNfs3TargetPtrOutput() Nfs3TargetPtrOutput {
	return i.ToNfs3TargetPtrOutputWithContext(context.Background())
}

func (i *nfs3TargetPtrType) ToNfs3TargetPtrOutputWithContext(ctx context.Context) Nfs3TargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetPtrOutput)
}

type Nfs3TargetOutput struct{ *pulumi.OutputState }

func (Nfs3TargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Nfs3Target)(nil)).Elem()
}

func (o Nfs3TargetOutput) ToNfs3TargetOutput() Nfs3TargetOutput {
	return o
}

func (o Nfs3TargetOutput) ToNfs3TargetOutputWithContext(ctx context.Context) Nfs3TargetOutput {
	return o
}

func (o Nfs3TargetOutput) ToNfs3TargetPtrOutput() Nfs3TargetPtrOutput {
	return o.ToNfs3TargetPtrOutputWithContext(context.Background())
}

func (o Nfs3TargetOutput) ToNfs3TargetPtrOutputWithContext(ctx context.Context) Nfs3TargetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Nfs3Target) *Nfs3Target {
		return &v
	}).(Nfs3TargetPtrOutput)
}

func (o Nfs3TargetOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Nfs3Target) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o Nfs3TargetOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Nfs3Target) *string { return v.UsageModel }).(pulumi.StringPtrOutput)
}

type Nfs3TargetPtrOutput struct{ *pulumi.OutputState }

func (Nfs3TargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nfs3Target)(nil)).Elem()
}

func (o Nfs3TargetPtrOutput) ToNfs3TargetPtrOutput() Nfs3TargetPtrOutput {
	return o
}

func (o Nfs3TargetPtrOutput) ToNfs3TargetPtrOutputWithContext(ctx context.Context) Nfs3TargetPtrOutput {
	return o
}

func (o Nfs3TargetPtrOutput) Elem() Nfs3TargetOutput {
	return o.ApplyT(func(v *Nfs3Target) Nfs3Target {
		if v != nil {
			return *v
		}
		var ret Nfs3Target
		return ret
	}).(Nfs3TargetOutput)
}

func (o Nfs3TargetPtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nfs3Target) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

func (o Nfs3TargetPtrOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nfs3Target) *string {
		if v == nil {
			return nil
		}
		return v.UsageModel
	}).(pulumi.StringPtrOutput)
}

type Nfs3TargetResponse struct {
	Target     *string `pulumi:"target"`
	UsageModel *string `pulumi:"usageModel"`
}





type Nfs3TargetResponseInput interface {
	pulumi.Input

	ToNfs3TargetResponseOutput() Nfs3TargetResponseOutput
	ToNfs3TargetResponseOutputWithContext(context.Context) Nfs3TargetResponseOutput
}

type Nfs3TargetResponseArgs struct {
	Target     pulumi.StringPtrInput `pulumi:"target"`
	UsageModel pulumi.StringPtrInput `pulumi:"usageModel"`
}

func (Nfs3TargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Nfs3TargetResponse)(nil)).Elem()
}

func (i Nfs3TargetResponseArgs) ToNfs3TargetResponseOutput() Nfs3TargetResponseOutput {
	return i.ToNfs3TargetResponseOutputWithContext(context.Background())
}

func (i Nfs3TargetResponseArgs) ToNfs3TargetResponseOutputWithContext(ctx context.Context) Nfs3TargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetResponseOutput)
}

func (i Nfs3TargetResponseArgs) ToNfs3TargetResponsePtrOutput() Nfs3TargetResponsePtrOutput {
	return i.ToNfs3TargetResponsePtrOutputWithContext(context.Background())
}

func (i Nfs3TargetResponseArgs) ToNfs3TargetResponsePtrOutputWithContext(ctx context.Context) Nfs3TargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetResponseOutput).ToNfs3TargetResponsePtrOutputWithContext(ctx)
}









type Nfs3TargetResponsePtrInput interface {
	pulumi.Input

	ToNfs3TargetResponsePtrOutput() Nfs3TargetResponsePtrOutput
	ToNfs3TargetResponsePtrOutputWithContext(context.Context) Nfs3TargetResponsePtrOutput
}

type nfs3TargetResponsePtrType Nfs3TargetResponseArgs

func Nfs3TargetResponsePtr(v *Nfs3TargetResponseArgs) Nfs3TargetResponsePtrInput {
	return (*nfs3TargetResponsePtrType)(v)
}

func (*nfs3TargetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Nfs3TargetResponse)(nil)).Elem()
}

func (i *nfs3TargetResponsePtrType) ToNfs3TargetResponsePtrOutput() Nfs3TargetResponsePtrOutput {
	return i.ToNfs3TargetResponsePtrOutputWithContext(context.Background())
}

func (i *nfs3TargetResponsePtrType) ToNfs3TargetResponsePtrOutputWithContext(ctx context.Context) Nfs3TargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetResponsePtrOutput)
}

type Nfs3TargetResponseOutput struct{ *pulumi.OutputState }

func (Nfs3TargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Nfs3TargetResponse)(nil)).Elem()
}

func (o Nfs3TargetResponseOutput) ToNfs3TargetResponseOutput() Nfs3TargetResponseOutput {
	return o
}

func (o Nfs3TargetResponseOutput) ToNfs3TargetResponseOutputWithContext(ctx context.Context) Nfs3TargetResponseOutput {
	return o
}

func (o Nfs3TargetResponseOutput) ToNfs3TargetResponsePtrOutput() Nfs3TargetResponsePtrOutput {
	return o.ToNfs3TargetResponsePtrOutputWithContext(context.Background())
}

func (o Nfs3TargetResponseOutput) ToNfs3TargetResponsePtrOutputWithContext(ctx context.Context) Nfs3TargetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Nfs3TargetResponse) *Nfs3TargetResponse {
		return &v
	}).(Nfs3TargetResponsePtrOutput)
}

func (o Nfs3TargetResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Nfs3TargetResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o Nfs3TargetResponseOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Nfs3TargetResponse) *string { return v.UsageModel }).(pulumi.StringPtrOutput)
}

type Nfs3TargetResponsePtrOutput struct{ *pulumi.OutputState }

func (Nfs3TargetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nfs3TargetResponse)(nil)).Elem()
}

func (o Nfs3TargetResponsePtrOutput) ToNfs3TargetResponsePtrOutput() Nfs3TargetResponsePtrOutput {
	return o
}

func (o Nfs3TargetResponsePtrOutput) ToNfs3TargetResponsePtrOutputWithContext(ctx context.Context) Nfs3TargetResponsePtrOutput {
	return o
}

func (o Nfs3TargetResponsePtrOutput) Elem() Nfs3TargetResponseOutput {
	return o.ApplyT(func(v *Nfs3TargetResponse) Nfs3TargetResponse {
		if v != nil {
			return *v
		}
		var ret Nfs3TargetResponse
		return ret
	}).(Nfs3TargetResponseOutput)
}

func (o Nfs3TargetResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nfs3TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

func (o Nfs3TargetResponsePtrOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nfs3TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.UsageModel
	}).(pulumi.StringPtrOutput)
}

type NfsAccessPolicy struct {
	AccessRules []NfsAccessRule `pulumi:"accessRules"`
	Name        string          `pulumi:"name"`
}





type NfsAccessPolicyInput interface {
	pulumi.Input

	ToNfsAccessPolicyOutput() NfsAccessPolicyOutput
	ToNfsAccessPolicyOutputWithContext(context.Context) NfsAccessPolicyOutput
}

type NfsAccessPolicyArgs struct {
	AccessRules NfsAccessRuleArrayInput `pulumi:"accessRules"`
	Name        pulumi.StringInput      `pulumi:"name"`
}

func (NfsAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessPolicy)(nil)).Elem()
}

func (i NfsAccessPolicyArgs) ToNfsAccessPolicyOutput() NfsAccessPolicyOutput {
	return i.ToNfsAccessPolicyOutputWithContext(context.Background())
}

func (i NfsAccessPolicyArgs) ToNfsAccessPolicyOutputWithContext(ctx context.Context) NfsAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsAccessPolicyOutput)
}





type NfsAccessPolicyArrayInput interface {
	pulumi.Input

	ToNfsAccessPolicyArrayOutput() NfsAccessPolicyArrayOutput
	ToNfsAccessPolicyArrayOutputWithContext(context.Context) NfsAccessPolicyArrayOutput
}

type NfsAccessPolicyArray []NfsAccessPolicyInput

func (NfsAccessPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsAccessPolicy)(nil)).Elem()
}

func (i NfsAccessPolicyArray) ToNfsAccessPolicyArrayOutput() NfsAccessPolicyArrayOutput {
	return i.ToNfsAccessPolicyArrayOutputWithContext(context.Background())
}

func (i NfsAccessPolicyArray) ToNfsAccessPolicyArrayOutputWithContext(ctx context.Context) NfsAccessPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsAccessPolicyArrayOutput)
}

type NfsAccessPolicyOutput struct{ *pulumi.OutputState }

func (NfsAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessPolicy)(nil)).Elem()
}

func (o NfsAccessPolicyOutput) ToNfsAccessPolicyOutput() NfsAccessPolicyOutput {
	return o
}

func (o NfsAccessPolicyOutput) ToNfsAccessPolicyOutputWithContext(ctx context.Context) NfsAccessPolicyOutput {
	return o
}

func (o NfsAccessPolicyOutput) AccessRules() NfsAccessRuleArrayOutput {
	return o.ApplyT(func(v NfsAccessPolicy) []NfsAccessRule { return v.AccessRules }).(NfsAccessRuleArrayOutput)
}

func (o NfsAccessPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NfsAccessPolicy) string { return v.Name }).(pulumi.StringOutput)
}

type NfsAccessPolicyArrayOutput struct{ *pulumi.OutputState }

func (NfsAccessPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsAccessPolicy)(nil)).Elem()
}

func (o NfsAccessPolicyArrayOutput) ToNfsAccessPolicyArrayOutput() NfsAccessPolicyArrayOutput {
	return o
}

func (o NfsAccessPolicyArrayOutput) ToNfsAccessPolicyArrayOutputWithContext(ctx context.Context) NfsAccessPolicyArrayOutput {
	return o
}

func (o NfsAccessPolicyArrayOutput) Index(i pulumi.IntInput) NfsAccessPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NfsAccessPolicy {
		return vs[0].([]NfsAccessPolicy)[vs[1].(int)]
	}).(NfsAccessPolicyOutput)
}

type NfsAccessPolicyResponse struct {
	AccessRules []NfsAccessRuleResponse `pulumi:"accessRules"`
	Name        string                  `pulumi:"name"`
}





type NfsAccessPolicyResponseInput interface {
	pulumi.Input

	ToNfsAccessPolicyResponseOutput() NfsAccessPolicyResponseOutput
	ToNfsAccessPolicyResponseOutputWithContext(context.Context) NfsAccessPolicyResponseOutput
}

type NfsAccessPolicyResponseArgs struct {
	AccessRules NfsAccessRuleResponseArrayInput `pulumi:"accessRules"`
	Name        pulumi.StringInput              `pulumi:"name"`
}

func (NfsAccessPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessPolicyResponse)(nil)).Elem()
}

func (i NfsAccessPolicyResponseArgs) ToNfsAccessPolicyResponseOutput() NfsAccessPolicyResponseOutput {
	return i.ToNfsAccessPolicyResponseOutputWithContext(context.Background())
}

func (i NfsAccessPolicyResponseArgs) ToNfsAccessPolicyResponseOutputWithContext(ctx context.Context) NfsAccessPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsAccessPolicyResponseOutput)
}





type NfsAccessPolicyResponseArrayInput interface {
	pulumi.Input

	ToNfsAccessPolicyResponseArrayOutput() NfsAccessPolicyResponseArrayOutput
	ToNfsAccessPolicyResponseArrayOutputWithContext(context.Context) NfsAccessPolicyResponseArrayOutput
}

type NfsAccessPolicyResponseArray []NfsAccessPolicyResponseInput

func (NfsAccessPolicyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsAccessPolicyResponse)(nil)).Elem()
}

func (i NfsAccessPolicyResponseArray) ToNfsAccessPolicyResponseArrayOutput() NfsAccessPolicyResponseArrayOutput {
	return i.ToNfsAccessPolicyResponseArrayOutputWithContext(context.Background())
}

func (i NfsAccessPolicyResponseArray) ToNfsAccessPolicyResponseArrayOutputWithContext(ctx context.Context) NfsAccessPolicyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsAccessPolicyResponseArrayOutput)
}

type NfsAccessPolicyResponseOutput struct{ *pulumi.OutputState }

func (NfsAccessPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessPolicyResponse)(nil)).Elem()
}

func (o NfsAccessPolicyResponseOutput) ToNfsAccessPolicyResponseOutput() NfsAccessPolicyResponseOutput {
	return o
}

func (o NfsAccessPolicyResponseOutput) ToNfsAccessPolicyResponseOutputWithContext(ctx context.Context) NfsAccessPolicyResponseOutput {
	return o
}

func (o NfsAccessPolicyResponseOutput) AccessRules() NfsAccessRuleResponseArrayOutput {
	return o.ApplyT(func(v NfsAccessPolicyResponse) []NfsAccessRuleResponse { return v.AccessRules }).(NfsAccessRuleResponseArrayOutput)
}

func (o NfsAccessPolicyResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NfsAccessPolicyResponse) string { return v.Name }).(pulumi.StringOutput)
}

type NfsAccessPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (NfsAccessPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsAccessPolicyResponse)(nil)).Elem()
}

func (o NfsAccessPolicyResponseArrayOutput) ToNfsAccessPolicyResponseArrayOutput() NfsAccessPolicyResponseArrayOutput {
	return o
}

func (o NfsAccessPolicyResponseArrayOutput) ToNfsAccessPolicyResponseArrayOutputWithContext(ctx context.Context) NfsAccessPolicyResponseArrayOutput {
	return o
}

func (o NfsAccessPolicyResponseArrayOutput) Index(i pulumi.IntInput) NfsAccessPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NfsAccessPolicyResponse {
		return vs[0].([]NfsAccessPolicyResponse)[vs[1].(int)]
	}).(NfsAccessPolicyResponseOutput)
}

type NfsAccessRule struct {
	Access         string  `pulumi:"access"`
	AnonymousGID   *string `pulumi:"anonymousGID"`
	AnonymousUID   *string `pulumi:"anonymousUID"`
	Filter         *string `pulumi:"filter"`
	RootSquash     *bool   `pulumi:"rootSquash"`
	Scope          string  `pulumi:"scope"`
	SubmountAccess *bool   `pulumi:"submountAccess"`
	Suid           *bool   `pulumi:"suid"`
}





type NfsAccessRuleInput interface {
	pulumi.Input

	ToNfsAccessRuleOutput() NfsAccessRuleOutput
	ToNfsAccessRuleOutputWithContext(context.Context) NfsAccessRuleOutput
}

type NfsAccessRuleArgs struct {
	Access         pulumi.StringInput    `pulumi:"access"`
	AnonymousGID   pulumi.StringPtrInput `pulumi:"anonymousGID"`
	AnonymousUID   pulumi.StringPtrInput `pulumi:"anonymousUID"`
	Filter         pulumi.StringPtrInput `pulumi:"filter"`
	RootSquash     pulumi.BoolPtrInput   `pulumi:"rootSquash"`
	Scope          pulumi.StringInput    `pulumi:"scope"`
	SubmountAccess pulumi.BoolPtrInput   `pulumi:"submountAccess"`
	Suid           pulumi.BoolPtrInput   `pulumi:"suid"`
}

func (NfsAccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessRule)(nil)).Elem()
}

func (i NfsAccessRuleArgs) ToNfsAccessRuleOutput() NfsAccessRuleOutput {
	return i.ToNfsAccessRuleOutputWithContext(context.Background())
}

func (i NfsAccessRuleArgs) ToNfsAccessRuleOutputWithContext(ctx context.Context) NfsAccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsAccessRuleOutput)
}





type NfsAccessRuleArrayInput interface {
	pulumi.Input

	ToNfsAccessRuleArrayOutput() NfsAccessRuleArrayOutput
	ToNfsAccessRuleArrayOutputWithContext(context.Context) NfsAccessRuleArrayOutput
}

type NfsAccessRuleArray []NfsAccessRuleInput

func (NfsAccessRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsAccessRule)(nil)).Elem()
}

func (i NfsAccessRuleArray) ToNfsAccessRuleArrayOutput() NfsAccessRuleArrayOutput {
	return i.ToNfsAccessRuleArrayOutputWithContext(context.Background())
}

func (i NfsAccessRuleArray) ToNfsAccessRuleArrayOutputWithContext(ctx context.Context) NfsAccessRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsAccessRuleArrayOutput)
}

type NfsAccessRuleOutput struct{ *pulumi.OutputState }

func (NfsAccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessRule)(nil)).Elem()
}

func (o NfsAccessRuleOutput) ToNfsAccessRuleOutput() NfsAccessRuleOutput {
	return o
}

func (o NfsAccessRuleOutput) ToNfsAccessRuleOutputWithContext(ctx context.Context) NfsAccessRuleOutput {
	return o
}

func (o NfsAccessRuleOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v NfsAccessRule) string { return v.Access }).(pulumi.StringOutput)
}

func (o NfsAccessRuleOutput) AnonymousGID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NfsAccessRule) *string { return v.AnonymousGID }).(pulumi.StringPtrOutput)
}

func (o NfsAccessRuleOutput) AnonymousUID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NfsAccessRule) *string { return v.AnonymousUID }).(pulumi.StringPtrOutput)
}

func (o NfsAccessRuleOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NfsAccessRule) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o NfsAccessRuleOutput) RootSquash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NfsAccessRule) *bool { return v.RootSquash }).(pulumi.BoolPtrOutput)
}

func (o NfsAccessRuleOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v NfsAccessRule) string { return v.Scope }).(pulumi.StringOutput)
}

func (o NfsAccessRuleOutput) SubmountAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NfsAccessRule) *bool { return v.SubmountAccess }).(pulumi.BoolPtrOutput)
}

func (o NfsAccessRuleOutput) Suid() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NfsAccessRule) *bool { return v.Suid }).(pulumi.BoolPtrOutput)
}

type NfsAccessRuleArrayOutput struct{ *pulumi.OutputState }

func (NfsAccessRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsAccessRule)(nil)).Elem()
}

func (o NfsAccessRuleArrayOutput) ToNfsAccessRuleArrayOutput() NfsAccessRuleArrayOutput {
	return o
}

func (o NfsAccessRuleArrayOutput) ToNfsAccessRuleArrayOutputWithContext(ctx context.Context) NfsAccessRuleArrayOutput {
	return o
}

func (o NfsAccessRuleArrayOutput) Index(i pulumi.IntInput) NfsAccessRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NfsAccessRule {
		return vs[0].([]NfsAccessRule)[vs[1].(int)]
	}).(NfsAccessRuleOutput)
}

type NfsAccessRuleResponse struct {
	Access         string  `pulumi:"access"`
	AnonymousGID   *string `pulumi:"anonymousGID"`
	AnonymousUID   *string `pulumi:"anonymousUID"`
	Filter         *string `pulumi:"filter"`
	RootSquash     *bool   `pulumi:"rootSquash"`
	Scope          string  `pulumi:"scope"`
	SubmountAccess *bool   `pulumi:"submountAccess"`
	Suid           *bool   `pulumi:"suid"`
}





type NfsAccessRuleResponseInput interface {
	pulumi.Input

	ToNfsAccessRuleResponseOutput() NfsAccessRuleResponseOutput
	ToNfsAccessRuleResponseOutputWithContext(context.Context) NfsAccessRuleResponseOutput
}

type NfsAccessRuleResponseArgs struct {
	Access         pulumi.StringInput    `pulumi:"access"`
	AnonymousGID   pulumi.StringPtrInput `pulumi:"anonymousGID"`
	AnonymousUID   pulumi.StringPtrInput `pulumi:"anonymousUID"`
	Filter         pulumi.StringPtrInput `pulumi:"filter"`
	RootSquash     pulumi.BoolPtrInput   `pulumi:"rootSquash"`
	Scope          pulumi.StringInput    `pulumi:"scope"`
	SubmountAccess pulumi.BoolPtrInput   `pulumi:"submountAccess"`
	Suid           pulumi.BoolPtrInput   `pulumi:"suid"`
}

func (NfsAccessRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessRuleResponse)(nil)).Elem()
}

func (i NfsAccessRuleResponseArgs) ToNfsAccessRuleResponseOutput() NfsAccessRuleResponseOutput {
	return i.ToNfsAccessRuleResponseOutputWithContext(context.Background())
}

func (i NfsAccessRuleResponseArgs) ToNfsAccessRuleResponseOutputWithContext(ctx context.Context) NfsAccessRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsAccessRuleResponseOutput)
}





type NfsAccessRuleResponseArrayInput interface {
	pulumi.Input

	ToNfsAccessRuleResponseArrayOutput() NfsAccessRuleResponseArrayOutput
	ToNfsAccessRuleResponseArrayOutputWithContext(context.Context) NfsAccessRuleResponseArrayOutput
}

type NfsAccessRuleResponseArray []NfsAccessRuleResponseInput

func (NfsAccessRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsAccessRuleResponse)(nil)).Elem()
}

func (i NfsAccessRuleResponseArray) ToNfsAccessRuleResponseArrayOutput() NfsAccessRuleResponseArrayOutput {
	return i.ToNfsAccessRuleResponseArrayOutputWithContext(context.Background())
}

func (i NfsAccessRuleResponseArray) ToNfsAccessRuleResponseArrayOutputWithContext(ctx context.Context) NfsAccessRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsAccessRuleResponseArrayOutput)
}

type NfsAccessRuleResponseOutput struct{ *pulumi.OutputState }

func (NfsAccessRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessRuleResponse)(nil)).Elem()
}

func (o NfsAccessRuleResponseOutput) ToNfsAccessRuleResponseOutput() NfsAccessRuleResponseOutput {
	return o
}

func (o NfsAccessRuleResponseOutput) ToNfsAccessRuleResponseOutputWithContext(ctx context.Context) NfsAccessRuleResponseOutput {
	return o
}

func (o NfsAccessRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v NfsAccessRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o NfsAccessRuleResponseOutput) AnonymousGID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NfsAccessRuleResponse) *string { return v.AnonymousGID }).(pulumi.StringPtrOutput)
}

func (o NfsAccessRuleResponseOutput) AnonymousUID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NfsAccessRuleResponse) *string { return v.AnonymousUID }).(pulumi.StringPtrOutput)
}

func (o NfsAccessRuleResponseOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NfsAccessRuleResponse) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o NfsAccessRuleResponseOutput) RootSquash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NfsAccessRuleResponse) *bool { return v.RootSquash }).(pulumi.BoolPtrOutput)
}

func (o NfsAccessRuleResponseOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v NfsAccessRuleResponse) string { return v.Scope }).(pulumi.StringOutput)
}

func (o NfsAccessRuleResponseOutput) SubmountAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NfsAccessRuleResponse) *bool { return v.SubmountAccess }).(pulumi.BoolPtrOutput)
}

func (o NfsAccessRuleResponseOutput) Suid() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NfsAccessRuleResponse) *bool { return v.Suid }).(pulumi.BoolPtrOutput)
}

type NfsAccessRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (NfsAccessRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsAccessRuleResponse)(nil)).Elem()
}

func (o NfsAccessRuleResponseArrayOutput) ToNfsAccessRuleResponseArrayOutput() NfsAccessRuleResponseArrayOutput {
	return o
}

func (o NfsAccessRuleResponseArrayOutput) ToNfsAccessRuleResponseArrayOutputWithContext(ctx context.Context) NfsAccessRuleResponseArrayOutput {
	return o
}

func (o NfsAccessRuleResponseArrayOutput) Index(i pulumi.IntInput) NfsAccessRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NfsAccessRuleResponse {
		return vs[0].([]NfsAccessRuleResponse)[vs[1].(int)]
	}).(NfsAccessRuleResponseOutput)
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

type UnknownTarget struct {
	Attributes map[string]string `pulumi:"attributes"`
}





type UnknownTargetInput interface {
	pulumi.Input

	ToUnknownTargetOutput() UnknownTargetOutput
	ToUnknownTargetOutputWithContext(context.Context) UnknownTargetOutput
}

type UnknownTargetArgs struct {
	Attributes pulumi.StringMapInput `pulumi:"attributes"`
}

func (UnknownTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UnknownTarget)(nil)).Elem()
}

func (i UnknownTargetArgs) ToUnknownTargetOutput() UnknownTargetOutput {
	return i.ToUnknownTargetOutputWithContext(context.Background())
}

func (i UnknownTargetArgs) ToUnknownTargetOutputWithContext(ctx context.Context) UnknownTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetOutput)
}

func (i UnknownTargetArgs) ToUnknownTargetPtrOutput() UnknownTargetPtrOutput {
	return i.ToUnknownTargetPtrOutputWithContext(context.Background())
}

func (i UnknownTargetArgs) ToUnknownTargetPtrOutputWithContext(ctx context.Context) UnknownTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetOutput).ToUnknownTargetPtrOutputWithContext(ctx)
}









type UnknownTargetPtrInput interface {
	pulumi.Input

	ToUnknownTargetPtrOutput() UnknownTargetPtrOutput
	ToUnknownTargetPtrOutputWithContext(context.Context) UnknownTargetPtrOutput
}

type unknownTargetPtrType UnknownTargetArgs

func UnknownTargetPtr(v *UnknownTargetArgs) UnknownTargetPtrInput {
	return (*unknownTargetPtrType)(v)
}

func (*unknownTargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UnknownTarget)(nil)).Elem()
}

func (i *unknownTargetPtrType) ToUnknownTargetPtrOutput() UnknownTargetPtrOutput {
	return i.ToUnknownTargetPtrOutputWithContext(context.Background())
}

func (i *unknownTargetPtrType) ToUnknownTargetPtrOutputWithContext(ctx context.Context) UnknownTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetPtrOutput)
}

type UnknownTargetOutput struct{ *pulumi.OutputState }

func (UnknownTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnknownTarget)(nil)).Elem()
}

func (o UnknownTargetOutput) ToUnknownTargetOutput() UnknownTargetOutput {
	return o
}

func (o UnknownTargetOutput) ToUnknownTargetOutputWithContext(ctx context.Context) UnknownTargetOutput {
	return o
}

func (o UnknownTargetOutput) ToUnknownTargetPtrOutput() UnknownTargetPtrOutput {
	return o.ToUnknownTargetPtrOutputWithContext(context.Background())
}

func (o UnknownTargetOutput) ToUnknownTargetPtrOutputWithContext(ctx context.Context) UnknownTargetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UnknownTarget) *UnknownTarget {
		return &v
	}).(UnknownTargetPtrOutput)
}

func (o UnknownTargetOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v UnknownTarget) map[string]string { return v.Attributes }).(pulumi.StringMapOutput)
}

type UnknownTargetPtrOutput struct{ *pulumi.OutputState }

func (UnknownTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnknownTarget)(nil)).Elem()
}

func (o UnknownTargetPtrOutput) ToUnknownTargetPtrOutput() UnknownTargetPtrOutput {
	return o
}

func (o UnknownTargetPtrOutput) ToUnknownTargetPtrOutputWithContext(ctx context.Context) UnknownTargetPtrOutput {
	return o
}

func (o UnknownTargetPtrOutput) Elem() UnknownTargetOutput {
	return o.ApplyT(func(v *UnknownTarget) UnknownTarget {
		if v != nil {
			return *v
		}
		var ret UnknownTarget
		return ret
	}).(UnknownTargetOutput)
}

func (o UnknownTargetPtrOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *UnknownTarget) map[string]string {
		if v == nil {
			return nil
		}
		return v.Attributes
	}).(pulumi.StringMapOutput)
}

type UnknownTargetResponse struct {
	Attributes map[string]string `pulumi:"attributes"`
}





type UnknownTargetResponseInput interface {
	pulumi.Input

	ToUnknownTargetResponseOutput() UnknownTargetResponseOutput
	ToUnknownTargetResponseOutputWithContext(context.Context) UnknownTargetResponseOutput
}

type UnknownTargetResponseArgs struct {
	Attributes pulumi.StringMapInput `pulumi:"attributes"`
}

func (UnknownTargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UnknownTargetResponse)(nil)).Elem()
}

func (i UnknownTargetResponseArgs) ToUnknownTargetResponseOutput() UnknownTargetResponseOutput {
	return i.ToUnknownTargetResponseOutputWithContext(context.Background())
}

func (i UnknownTargetResponseArgs) ToUnknownTargetResponseOutputWithContext(ctx context.Context) UnknownTargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetResponseOutput)
}

func (i UnknownTargetResponseArgs) ToUnknownTargetResponsePtrOutput() UnknownTargetResponsePtrOutput {
	return i.ToUnknownTargetResponsePtrOutputWithContext(context.Background())
}

func (i UnknownTargetResponseArgs) ToUnknownTargetResponsePtrOutputWithContext(ctx context.Context) UnknownTargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetResponseOutput).ToUnknownTargetResponsePtrOutputWithContext(ctx)
}









type UnknownTargetResponsePtrInput interface {
	pulumi.Input

	ToUnknownTargetResponsePtrOutput() UnknownTargetResponsePtrOutput
	ToUnknownTargetResponsePtrOutputWithContext(context.Context) UnknownTargetResponsePtrOutput
}

type unknownTargetResponsePtrType UnknownTargetResponseArgs

func UnknownTargetResponsePtr(v *UnknownTargetResponseArgs) UnknownTargetResponsePtrInput {
	return (*unknownTargetResponsePtrType)(v)
}

func (*unknownTargetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UnknownTargetResponse)(nil)).Elem()
}

func (i *unknownTargetResponsePtrType) ToUnknownTargetResponsePtrOutput() UnknownTargetResponsePtrOutput {
	return i.ToUnknownTargetResponsePtrOutputWithContext(context.Background())
}

func (i *unknownTargetResponsePtrType) ToUnknownTargetResponsePtrOutputWithContext(ctx context.Context) UnknownTargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetResponsePtrOutput)
}

type UnknownTargetResponseOutput struct{ *pulumi.OutputState }

func (UnknownTargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnknownTargetResponse)(nil)).Elem()
}

func (o UnknownTargetResponseOutput) ToUnknownTargetResponseOutput() UnknownTargetResponseOutput {
	return o
}

func (o UnknownTargetResponseOutput) ToUnknownTargetResponseOutputWithContext(ctx context.Context) UnknownTargetResponseOutput {
	return o
}

func (o UnknownTargetResponseOutput) ToUnknownTargetResponsePtrOutput() UnknownTargetResponsePtrOutput {
	return o.ToUnknownTargetResponsePtrOutputWithContext(context.Background())
}

func (o UnknownTargetResponseOutput) ToUnknownTargetResponsePtrOutputWithContext(ctx context.Context) UnknownTargetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UnknownTargetResponse) *UnknownTargetResponse {
		return &v
	}).(UnknownTargetResponsePtrOutput)
}

func (o UnknownTargetResponseOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v UnknownTargetResponse) map[string]string { return v.Attributes }).(pulumi.StringMapOutput)
}

type UnknownTargetResponsePtrOutput struct{ *pulumi.OutputState }

func (UnknownTargetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnknownTargetResponse)(nil)).Elem()
}

func (o UnknownTargetResponsePtrOutput) ToUnknownTargetResponsePtrOutput() UnknownTargetResponsePtrOutput {
	return o
}

func (o UnknownTargetResponsePtrOutput) ToUnknownTargetResponsePtrOutputWithContext(ctx context.Context) UnknownTargetResponsePtrOutput {
	return o
}

func (o UnknownTargetResponsePtrOutput) Elem() UnknownTargetResponseOutput {
	return o.ApplyT(func(v *UnknownTargetResponse) UnknownTargetResponse {
		if v != nil {
			return *v
		}
		var ret UnknownTargetResponse
		return ret
	}).(UnknownTargetResponseOutput)
}

func (o UnknownTargetResponsePtrOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *UnknownTargetResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Attributes
	}).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobNfsTargetOutput{})
	pulumi.RegisterOutputType(BlobNfsTargetPtrOutput{})
	pulumi.RegisterOutputType(BlobNfsTargetResponseOutput{})
	pulumi.RegisterOutputType(BlobNfsTargetResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheActiveDirectorySettingsOutput{})
	pulumi.RegisterOutputType(CacheActiveDirectorySettingsPtrOutput{})
	pulumi.RegisterOutputType(CacheActiveDirectorySettingsCredentialsOutput{})
	pulumi.RegisterOutputType(CacheActiveDirectorySettingsCredentialsPtrOutput{})
	pulumi.RegisterOutputType(CacheActiveDirectorySettingsResponseOutput{})
	pulumi.RegisterOutputType(CacheActiveDirectorySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheActiveDirectorySettingsResponseCredentialsOutput{})
	pulumi.RegisterOutputType(CacheActiveDirectorySettingsResponseCredentialsPtrOutput{})
	pulumi.RegisterOutputType(CacheDirectorySettingsOutput{})
	pulumi.RegisterOutputType(CacheDirectorySettingsPtrOutput{})
	pulumi.RegisterOutputType(CacheDirectorySettingsResponseOutput{})
	pulumi.RegisterOutputType(CacheDirectorySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheEncryptionSettingsOutput{})
	pulumi.RegisterOutputType(CacheEncryptionSettingsPtrOutput{})
	pulumi.RegisterOutputType(CacheEncryptionSettingsResponseOutput{})
	pulumi.RegisterOutputType(CacheEncryptionSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheHealthResponseOutput{})
	pulumi.RegisterOutputType(CacheHealthResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheIdentityOutput{})
	pulumi.RegisterOutputType(CacheIdentityPtrOutput{})
	pulumi.RegisterOutputType(CacheIdentityResponseOutput{})
	pulumi.RegisterOutputType(CacheIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheNetworkSettingsOutput{})
	pulumi.RegisterOutputType(CacheNetworkSettingsPtrOutput{})
	pulumi.RegisterOutputType(CacheNetworkSettingsResponseOutput{})
	pulumi.RegisterOutputType(CacheNetworkSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheResponseSkuOutput{})
	pulumi.RegisterOutputType(CacheResponseSkuPtrOutput{})
	pulumi.RegisterOutputType(CacheSecuritySettingsOutput{})
	pulumi.RegisterOutputType(CacheSecuritySettingsPtrOutput{})
	pulumi.RegisterOutputType(CacheSecuritySettingsResponseOutput{})
	pulumi.RegisterOutputType(CacheSecuritySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheSkuOutput{})
	pulumi.RegisterOutputType(CacheSkuPtrOutput{})
	pulumi.RegisterOutputType(CacheUpgradeStatusResponseOutput{})
	pulumi.RegisterOutputType(CacheUpgradeStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheUsernameDownloadSettingsOutput{})
	pulumi.RegisterOutputType(CacheUsernameDownloadSettingsPtrOutput{})
	pulumi.RegisterOutputType(CacheUsernameDownloadSettingsCredentialsOutput{})
	pulumi.RegisterOutputType(CacheUsernameDownloadSettingsCredentialsPtrOutput{})
	pulumi.RegisterOutputType(CacheUsernameDownloadSettingsResponseOutput{})
	pulumi.RegisterOutputType(CacheUsernameDownloadSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheUsernameDownloadSettingsResponseCredentialsOutput{})
	pulumi.RegisterOutputType(CacheUsernameDownloadSettingsResponseCredentialsPtrOutput{})
	pulumi.RegisterOutputType(ClfsTargetOutput{})
	pulumi.RegisterOutputType(ClfsTargetPtrOutput{})
	pulumi.RegisterOutputType(ClfsTargetResponseOutput{})
	pulumi.RegisterOutputType(ClfsTargetResponsePtrOutput{})
	pulumi.RegisterOutputType(ConditionResponseOutput{})
	pulumi.RegisterOutputType(ConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseSourceVaultOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseSourceVaultPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceSourceVaultOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceSourceVaultPtrOutput{})
	pulumi.RegisterOutputType(NamespaceJunctionOutput{})
	pulumi.RegisterOutputType(NamespaceJunctionArrayOutput{})
	pulumi.RegisterOutputType(NamespaceJunctionResponseOutput{})
	pulumi.RegisterOutputType(NamespaceJunctionResponseArrayOutput{})
	pulumi.RegisterOutputType(Nfs3TargetOutput{})
	pulumi.RegisterOutputType(Nfs3TargetPtrOutput{})
	pulumi.RegisterOutputType(Nfs3TargetResponseOutput{})
	pulumi.RegisterOutputType(Nfs3TargetResponsePtrOutput{})
	pulumi.RegisterOutputType(NfsAccessPolicyOutput{})
	pulumi.RegisterOutputType(NfsAccessPolicyArrayOutput{})
	pulumi.RegisterOutputType(NfsAccessPolicyResponseOutput{})
	pulumi.RegisterOutputType(NfsAccessPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(NfsAccessRuleOutput{})
	pulumi.RegisterOutputType(NfsAccessRuleArrayOutput{})
	pulumi.RegisterOutputType(NfsAccessRuleResponseOutput{})
	pulumi.RegisterOutputType(NfsAccessRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UnknownTargetOutput{})
	pulumi.RegisterOutputType(UnknownTargetPtrOutput{})
	pulumi.RegisterOutputType(UnknownTargetResponseOutput{})
	pulumi.RegisterOutputType(UnknownTargetResponsePtrOutput{})
}
