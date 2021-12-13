


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AwAssumeRoleAuthenticationDetailsProperties struct {
	AuthenticationType string `pulumi:"authenticationType"`
	AwsAssumeRoleArn   string `pulumi:"awsAssumeRoleArn"`
	AwsExternalId      string `pulumi:"awsExternalId"`
}

type AwAssumeRoleAuthenticationDetailsPropertiesResponse struct {
	AccountId                       string   `pulumi:"accountId"`
	AuthenticationProvisioningState string   `pulumi:"authenticationProvisioningState"`
	AuthenticationType              string   `pulumi:"authenticationType"`
	AwsAssumeRoleArn                string   `pulumi:"awsAssumeRoleArn"`
	AwsExternalId                   string   `pulumi:"awsExternalId"`
	GrantedPermissions              []string `pulumi:"grantedPermissions"`
}

type AwsCredsAuthenticationDetailsProperties struct {
	AuthenticationType string `pulumi:"authenticationType"`
	AwsAccessKeyId     string `pulumi:"awsAccessKeyId"`
	AwsSecretAccessKey string `pulumi:"awsSecretAccessKey"`
}

type AwsCredsAuthenticationDetailsPropertiesResponse struct {
	AccountId                       string   `pulumi:"accountId"`
	AuthenticationProvisioningState string   `pulumi:"authenticationProvisioningState"`
	AuthenticationType              string   `pulumi:"authenticationType"`
	AwsAccessKeyId                  string   `pulumi:"awsAccessKeyId"`
	AwsSecretAccessKey              string   `pulumi:"awsSecretAccessKey"`
	GrantedPermissions              []string `pulumi:"grantedPermissions"`
}

type GcpCredentialsDetailsProperties struct {
	AuthProviderX509CertUrl string `pulumi:"authProviderX509CertUrl"`
	AuthUri                 string `pulumi:"authUri"`
	AuthenticationType      string `pulumi:"authenticationType"`
	ClientEmail             string `pulumi:"clientEmail"`
	ClientId                string `pulumi:"clientId"`
	ClientX509CertUrl       string `pulumi:"clientX509CertUrl"`
	OrganizationId          string `pulumi:"organizationId"`
	PrivateKey              string `pulumi:"privateKey"`
	PrivateKeyId            string `pulumi:"privateKeyId"`
	ProjectId               string `pulumi:"projectId"`
	TokenUri                string `pulumi:"tokenUri"`
	Type                    string `pulumi:"type"`
}

type GcpCredentialsDetailsPropertiesResponse struct {
	AuthProviderX509CertUrl         string   `pulumi:"authProviderX509CertUrl"`
	AuthUri                         string   `pulumi:"authUri"`
	AuthenticationProvisioningState string   `pulumi:"authenticationProvisioningState"`
	AuthenticationType              string   `pulumi:"authenticationType"`
	ClientEmail                     string   `pulumi:"clientEmail"`
	ClientId                        string   `pulumi:"clientId"`
	ClientX509CertUrl               string   `pulumi:"clientX509CertUrl"`
	GrantedPermissions              []string `pulumi:"grantedPermissions"`
	OrganizationId                  string   `pulumi:"organizationId"`
	PrivateKey                      string   `pulumi:"privateKey"`
	PrivateKeyId                    string   `pulumi:"privateKeyId"`
	ProjectId                       string   `pulumi:"projectId"`
	TokenUri                        string   `pulumi:"tokenUri"`
	Type                            string   `pulumi:"type"`
}

type HybridComputeSettingsProperties struct {
	AutoProvision     string                      `pulumi:"autoProvision"`
	ProxyServer       *ProxyServerProperties      `pulumi:"proxyServer"`
	Region            *string                     `pulumi:"region"`
	ResourceGroupName *string                     `pulumi:"resourceGroupName"`
	ServicePrincipal  *ServicePrincipalProperties `pulumi:"servicePrincipal"`
}





type HybridComputeSettingsPropertiesInput interface {
	pulumi.Input

	ToHybridComputeSettingsPropertiesOutput() HybridComputeSettingsPropertiesOutput
	ToHybridComputeSettingsPropertiesOutputWithContext(context.Context) HybridComputeSettingsPropertiesOutput
}

type HybridComputeSettingsPropertiesArgs struct {
	AutoProvision     pulumi.StringInput                 `pulumi:"autoProvision"`
	ProxyServer       ProxyServerPropertiesPtrInput      `pulumi:"proxyServer"`
	Region            pulumi.StringPtrInput              `pulumi:"region"`
	ResourceGroupName pulumi.StringPtrInput              `pulumi:"resourceGroupName"`
	ServicePrincipal  ServicePrincipalPropertiesPtrInput `pulumi:"servicePrincipal"`
}

func (HybridComputeSettingsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputeSettingsProperties)(nil)).Elem()
}

func (i HybridComputeSettingsPropertiesArgs) ToHybridComputeSettingsPropertiesOutput() HybridComputeSettingsPropertiesOutput {
	return i.ToHybridComputeSettingsPropertiesOutputWithContext(context.Background())
}

func (i HybridComputeSettingsPropertiesArgs) ToHybridComputeSettingsPropertiesOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputeSettingsPropertiesOutput)
}

func (i HybridComputeSettingsPropertiesArgs) ToHybridComputeSettingsPropertiesPtrOutput() HybridComputeSettingsPropertiesPtrOutput {
	return i.ToHybridComputeSettingsPropertiesPtrOutputWithContext(context.Background())
}

func (i HybridComputeSettingsPropertiesArgs) ToHybridComputeSettingsPropertiesPtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputeSettingsPropertiesOutput).ToHybridComputeSettingsPropertiesPtrOutputWithContext(ctx)
}









type HybridComputeSettingsPropertiesPtrInput interface {
	pulumi.Input

	ToHybridComputeSettingsPropertiesPtrOutput() HybridComputeSettingsPropertiesPtrOutput
	ToHybridComputeSettingsPropertiesPtrOutputWithContext(context.Context) HybridComputeSettingsPropertiesPtrOutput
}

type hybridComputeSettingsPropertiesPtrType HybridComputeSettingsPropertiesArgs

func HybridComputeSettingsPropertiesPtr(v *HybridComputeSettingsPropertiesArgs) HybridComputeSettingsPropertiesPtrInput {
	return (*hybridComputeSettingsPropertiesPtrType)(v)
}

func (*hybridComputeSettingsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputeSettingsProperties)(nil)).Elem()
}

func (i *hybridComputeSettingsPropertiesPtrType) ToHybridComputeSettingsPropertiesPtrOutput() HybridComputeSettingsPropertiesPtrOutput {
	return i.ToHybridComputeSettingsPropertiesPtrOutputWithContext(context.Background())
}

func (i *hybridComputeSettingsPropertiesPtrType) ToHybridComputeSettingsPropertiesPtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputeSettingsPropertiesPtrOutput)
}

type HybridComputeSettingsPropertiesOutput struct{ *pulumi.OutputState }

func (HybridComputeSettingsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputeSettingsProperties)(nil)).Elem()
}

func (o HybridComputeSettingsPropertiesOutput) ToHybridComputeSettingsPropertiesOutput() HybridComputeSettingsPropertiesOutput {
	return o
}

func (o HybridComputeSettingsPropertiesOutput) ToHybridComputeSettingsPropertiesOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesOutput {
	return o
}

func (o HybridComputeSettingsPropertiesOutput) ToHybridComputeSettingsPropertiesPtrOutput() HybridComputeSettingsPropertiesPtrOutput {
	return o.ToHybridComputeSettingsPropertiesPtrOutputWithContext(context.Background())
}

func (o HybridComputeSettingsPropertiesOutput) ToHybridComputeSettingsPropertiesPtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HybridComputeSettingsProperties) *HybridComputeSettingsProperties {
		return &v
	}).(HybridComputeSettingsPropertiesPtrOutput)
}

func (o HybridComputeSettingsPropertiesOutput) AutoProvision() pulumi.StringOutput {
	return o.ApplyT(func(v HybridComputeSettingsProperties) string { return v.AutoProvision }).(pulumi.StringOutput)
}

func (o HybridComputeSettingsPropertiesOutput) ProxyServer() ProxyServerPropertiesPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsProperties) *ProxyServerProperties { return v.ProxyServer }).(ProxyServerPropertiesPtrOutput)
}

func (o HybridComputeSettingsPropertiesOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsProperties) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsProperties) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesOutput) ServicePrincipal() ServicePrincipalPropertiesPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsProperties) *ServicePrincipalProperties { return v.ServicePrincipal }).(ServicePrincipalPropertiesPtrOutput)
}

type HybridComputeSettingsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (HybridComputeSettingsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputeSettingsProperties)(nil)).Elem()
}

func (o HybridComputeSettingsPropertiesPtrOutput) ToHybridComputeSettingsPropertiesPtrOutput() HybridComputeSettingsPropertiesPtrOutput {
	return o
}

func (o HybridComputeSettingsPropertiesPtrOutput) ToHybridComputeSettingsPropertiesPtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesPtrOutput {
	return o
}

func (o HybridComputeSettingsPropertiesPtrOutput) Elem() HybridComputeSettingsPropertiesOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) HybridComputeSettingsProperties {
		if v != nil {
			return *v
		}
		var ret HybridComputeSettingsProperties
		return ret
	}).(HybridComputeSettingsPropertiesOutput)
}

func (o HybridComputeSettingsPropertiesPtrOutput) AutoProvision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AutoProvision
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesPtrOutput) ProxyServer() ProxyServerPropertiesPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) *ProxyServerProperties {
		if v == nil {
			return nil
		}
		return v.ProxyServer
	}).(ProxyServerPropertiesPtrOutput)
}

func (o HybridComputeSettingsPropertiesPtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) *string {
		if v == nil {
			return nil
		}
		return v.Region
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesPtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesPtrOutput) ServicePrincipal() ServicePrincipalPropertiesPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) *ServicePrincipalProperties {
		if v == nil {
			return nil
		}
		return v.ServicePrincipal
	}).(ServicePrincipalPropertiesPtrOutput)
}

type HybridComputeSettingsPropertiesResponse struct {
	AutoProvision                  string                              `pulumi:"autoProvision"`
	HybridComputeProvisioningState string                              `pulumi:"hybridComputeProvisioningState"`
	ProxyServer                    *ProxyServerPropertiesResponse      `pulumi:"proxyServer"`
	Region                         *string                             `pulumi:"region"`
	ResourceGroupName              *string                             `pulumi:"resourceGroupName"`
	ServicePrincipal               *ServicePrincipalPropertiesResponse `pulumi:"servicePrincipal"`
}

type HybridComputeSettingsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (HybridComputeSettingsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputeSettingsPropertiesResponse)(nil)).Elem()
}

func (o HybridComputeSettingsPropertiesResponseOutput) ToHybridComputeSettingsPropertiesResponseOutput() HybridComputeSettingsPropertiesResponseOutput {
	return o
}

func (o HybridComputeSettingsPropertiesResponseOutput) ToHybridComputeSettingsPropertiesResponseOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesResponseOutput {
	return o
}

func (o HybridComputeSettingsPropertiesResponseOutput) AutoProvision() pulumi.StringOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) string { return v.AutoProvision }).(pulumi.StringOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) HybridComputeProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) string { return v.HybridComputeProvisioningState }).(pulumi.StringOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) ProxyServer() ProxyServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) *ProxyServerPropertiesResponse { return v.ProxyServer }).(ProxyServerPropertiesResponsePtrOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) ServicePrincipal() ServicePrincipalPropertiesResponsePtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) *ServicePrincipalPropertiesResponse {
		return v.ServicePrincipal
	}).(ServicePrincipalPropertiesResponsePtrOutput)
}

type HybridComputeSettingsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (HybridComputeSettingsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputeSettingsPropertiesResponse)(nil)).Elem()
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) ToHybridComputeSettingsPropertiesResponsePtrOutput() HybridComputeSettingsPropertiesResponsePtrOutput {
	return o
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) ToHybridComputeSettingsPropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesResponsePtrOutput {
	return o
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) Elem() HybridComputeSettingsPropertiesResponseOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) HybridComputeSettingsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret HybridComputeSettingsPropertiesResponse
		return ret
	}).(HybridComputeSettingsPropertiesResponseOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) AutoProvision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AutoProvision
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) HybridComputeProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HybridComputeProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) ProxyServer() ProxyServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *ProxyServerPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ProxyServer
	}).(ProxyServerPropertiesResponsePtrOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Region
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) ServicePrincipal() ServicePrincipalPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *ServicePrincipalPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ServicePrincipal
	}).(ServicePrincipalPropertiesResponsePtrOutput)
}

type ProxyServerProperties struct {
	Ip   *string `pulumi:"ip"`
	Port *string `pulumi:"port"`
}





type ProxyServerPropertiesInput interface {
	pulumi.Input

	ToProxyServerPropertiesOutput() ProxyServerPropertiesOutput
	ToProxyServerPropertiesOutputWithContext(context.Context) ProxyServerPropertiesOutput
}

type ProxyServerPropertiesArgs struct {
	Ip   pulumi.StringPtrInput `pulumi:"ip"`
	Port pulumi.StringPtrInput `pulumi:"port"`
}

func (ProxyServerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyServerProperties)(nil)).Elem()
}

func (i ProxyServerPropertiesArgs) ToProxyServerPropertiesOutput() ProxyServerPropertiesOutput {
	return i.ToProxyServerPropertiesOutputWithContext(context.Background())
}

func (i ProxyServerPropertiesArgs) ToProxyServerPropertiesOutputWithContext(ctx context.Context) ProxyServerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyServerPropertiesOutput)
}

func (i ProxyServerPropertiesArgs) ToProxyServerPropertiesPtrOutput() ProxyServerPropertiesPtrOutput {
	return i.ToProxyServerPropertiesPtrOutputWithContext(context.Background())
}

func (i ProxyServerPropertiesArgs) ToProxyServerPropertiesPtrOutputWithContext(ctx context.Context) ProxyServerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyServerPropertiesOutput).ToProxyServerPropertiesPtrOutputWithContext(ctx)
}









type ProxyServerPropertiesPtrInput interface {
	pulumi.Input

	ToProxyServerPropertiesPtrOutput() ProxyServerPropertiesPtrOutput
	ToProxyServerPropertiesPtrOutputWithContext(context.Context) ProxyServerPropertiesPtrOutput
}

type proxyServerPropertiesPtrType ProxyServerPropertiesArgs

func ProxyServerPropertiesPtr(v *ProxyServerPropertiesArgs) ProxyServerPropertiesPtrInput {
	return (*proxyServerPropertiesPtrType)(v)
}

func (*proxyServerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyServerProperties)(nil)).Elem()
}

func (i *proxyServerPropertiesPtrType) ToProxyServerPropertiesPtrOutput() ProxyServerPropertiesPtrOutput {
	return i.ToProxyServerPropertiesPtrOutputWithContext(context.Background())
}

func (i *proxyServerPropertiesPtrType) ToProxyServerPropertiesPtrOutputWithContext(ctx context.Context) ProxyServerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyServerPropertiesPtrOutput)
}

type ProxyServerPropertiesOutput struct{ *pulumi.OutputState }

func (ProxyServerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyServerProperties)(nil)).Elem()
}

func (o ProxyServerPropertiesOutput) ToProxyServerPropertiesOutput() ProxyServerPropertiesOutput {
	return o
}

func (o ProxyServerPropertiesOutput) ToProxyServerPropertiesOutputWithContext(ctx context.Context) ProxyServerPropertiesOutput {
	return o
}

func (o ProxyServerPropertiesOutput) ToProxyServerPropertiesPtrOutput() ProxyServerPropertiesPtrOutput {
	return o.ToProxyServerPropertiesPtrOutputWithContext(context.Background())
}

func (o ProxyServerPropertiesOutput) ToProxyServerPropertiesPtrOutputWithContext(ctx context.Context) ProxyServerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProxyServerProperties) *ProxyServerProperties {
		return &v
	}).(ProxyServerPropertiesPtrOutput)
}

func (o ProxyServerPropertiesOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyServerProperties) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o ProxyServerPropertiesOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyServerProperties) *string { return v.Port }).(pulumi.StringPtrOutput)
}

type ProxyServerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ProxyServerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyServerProperties)(nil)).Elem()
}

func (o ProxyServerPropertiesPtrOutput) ToProxyServerPropertiesPtrOutput() ProxyServerPropertiesPtrOutput {
	return o
}

func (o ProxyServerPropertiesPtrOutput) ToProxyServerPropertiesPtrOutputWithContext(ctx context.Context) ProxyServerPropertiesPtrOutput {
	return o
}

func (o ProxyServerPropertiesPtrOutput) Elem() ProxyServerPropertiesOutput {
	return o.ApplyT(func(v *ProxyServerProperties) ProxyServerProperties {
		if v != nil {
			return *v
		}
		var ret ProxyServerProperties
		return ret
	}).(ProxyServerPropertiesOutput)
}

func (o ProxyServerPropertiesPtrOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyServerProperties) *string {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(pulumi.StringPtrOutput)
}

func (o ProxyServerPropertiesPtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyServerProperties) *string {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.StringPtrOutput)
}

type ProxyServerPropertiesResponse struct {
	Ip   *string `pulumi:"ip"`
	Port *string `pulumi:"port"`
}

type ProxyServerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ProxyServerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyServerPropertiesResponse)(nil)).Elem()
}

func (o ProxyServerPropertiesResponseOutput) ToProxyServerPropertiesResponseOutput() ProxyServerPropertiesResponseOutput {
	return o
}

func (o ProxyServerPropertiesResponseOutput) ToProxyServerPropertiesResponseOutputWithContext(ctx context.Context) ProxyServerPropertiesResponseOutput {
	return o
}

func (o ProxyServerPropertiesResponseOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyServerPropertiesResponse) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o ProxyServerPropertiesResponseOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyServerPropertiesResponse) *string { return v.Port }).(pulumi.StringPtrOutput)
}

type ProxyServerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ProxyServerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyServerPropertiesResponse)(nil)).Elem()
}

func (o ProxyServerPropertiesResponsePtrOutput) ToProxyServerPropertiesResponsePtrOutput() ProxyServerPropertiesResponsePtrOutput {
	return o
}

func (o ProxyServerPropertiesResponsePtrOutput) ToProxyServerPropertiesResponsePtrOutputWithContext(ctx context.Context) ProxyServerPropertiesResponsePtrOutput {
	return o
}

func (o ProxyServerPropertiesResponsePtrOutput) Elem() ProxyServerPropertiesResponseOutput {
	return o.ApplyT(func(v *ProxyServerPropertiesResponse) ProxyServerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ProxyServerPropertiesResponse
		return ret
	}).(ProxyServerPropertiesResponseOutput)
}

func (o ProxyServerPropertiesResponsePtrOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(pulumi.StringPtrOutput)
}

func (o ProxyServerPropertiesResponsePtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesAlertNotifications struct {
	MinimalSeverity *string `pulumi:"minimalSeverity"`
	State           *string `pulumi:"state"`
}





type SecurityContactPropertiesAlertNotificationsInput interface {
	pulumi.Input

	ToSecurityContactPropertiesAlertNotificationsOutput() SecurityContactPropertiesAlertNotificationsOutput
	ToSecurityContactPropertiesAlertNotificationsOutputWithContext(context.Context) SecurityContactPropertiesAlertNotificationsOutput
}

type SecurityContactPropertiesAlertNotificationsArgs struct {
	MinimalSeverity pulumi.StringPtrInput `pulumi:"minimalSeverity"`
	State           pulumi.StringPtrInput `pulumi:"state"`
}

func (SecurityContactPropertiesAlertNotificationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesAlertNotifications)(nil)).Elem()
}

func (i SecurityContactPropertiesAlertNotificationsArgs) ToSecurityContactPropertiesAlertNotificationsOutput() SecurityContactPropertiesAlertNotificationsOutput {
	return i.ToSecurityContactPropertiesAlertNotificationsOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesAlertNotificationsArgs) ToSecurityContactPropertiesAlertNotificationsOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesAlertNotificationsOutput)
}

func (i SecurityContactPropertiesAlertNotificationsArgs) ToSecurityContactPropertiesAlertNotificationsPtrOutput() SecurityContactPropertiesAlertNotificationsPtrOutput {
	return i.ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesAlertNotificationsArgs) ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesAlertNotificationsOutput).ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(ctx)
}









type SecurityContactPropertiesAlertNotificationsPtrInput interface {
	pulumi.Input

	ToSecurityContactPropertiesAlertNotificationsPtrOutput() SecurityContactPropertiesAlertNotificationsPtrOutput
	ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(context.Context) SecurityContactPropertiesAlertNotificationsPtrOutput
}

type securityContactPropertiesAlertNotificationsPtrType SecurityContactPropertiesAlertNotificationsArgs

func SecurityContactPropertiesAlertNotificationsPtr(v *SecurityContactPropertiesAlertNotificationsArgs) SecurityContactPropertiesAlertNotificationsPtrInput {
	return (*securityContactPropertiesAlertNotificationsPtrType)(v)
}

func (*securityContactPropertiesAlertNotificationsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesAlertNotifications)(nil)).Elem()
}

func (i *securityContactPropertiesAlertNotificationsPtrType) ToSecurityContactPropertiesAlertNotificationsPtrOutput() SecurityContactPropertiesAlertNotificationsPtrOutput {
	return i.ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(context.Background())
}

func (i *securityContactPropertiesAlertNotificationsPtrType) ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesAlertNotificationsPtrOutput)
}

type SecurityContactPropertiesAlertNotificationsOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesAlertNotificationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesAlertNotifications)(nil)).Elem()
}

func (o SecurityContactPropertiesAlertNotificationsOutput) ToSecurityContactPropertiesAlertNotificationsOutput() SecurityContactPropertiesAlertNotificationsOutput {
	return o
}

func (o SecurityContactPropertiesAlertNotificationsOutput) ToSecurityContactPropertiesAlertNotificationsOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsOutput {
	return o
}

func (o SecurityContactPropertiesAlertNotificationsOutput) ToSecurityContactPropertiesAlertNotificationsPtrOutput() SecurityContactPropertiesAlertNotificationsPtrOutput {
	return o.ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(context.Background())
}

func (o SecurityContactPropertiesAlertNotificationsOutput) ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityContactPropertiesAlertNotifications) *SecurityContactPropertiesAlertNotifications {
		return &v
	}).(SecurityContactPropertiesAlertNotificationsPtrOutput)
}

func (o SecurityContactPropertiesAlertNotificationsOutput) MinimalSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesAlertNotifications) *string { return v.MinimalSeverity }).(pulumi.StringPtrOutput)
}

func (o SecurityContactPropertiesAlertNotificationsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesAlertNotifications) *string { return v.State }).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesAlertNotificationsPtrOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesAlertNotificationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesAlertNotifications)(nil)).Elem()
}

func (o SecurityContactPropertiesAlertNotificationsPtrOutput) ToSecurityContactPropertiesAlertNotificationsPtrOutput() SecurityContactPropertiesAlertNotificationsPtrOutput {
	return o
}

func (o SecurityContactPropertiesAlertNotificationsPtrOutput) ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsPtrOutput {
	return o
}

func (o SecurityContactPropertiesAlertNotificationsPtrOutput) Elem() SecurityContactPropertiesAlertNotificationsOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesAlertNotifications) SecurityContactPropertiesAlertNotifications {
		if v != nil {
			return *v
		}
		var ret SecurityContactPropertiesAlertNotifications
		return ret
	}).(SecurityContactPropertiesAlertNotificationsOutput)
}

func (o SecurityContactPropertiesAlertNotificationsPtrOutput) MinimalSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesAlertNotifications) *string {
		if v == nil {
			return nil
		}
		return v.MinimalSeverity
	}).(pulumi.StringPtrOutput)
}

func (o SecurityContactPropertiesAlertNotificationsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesAlertNotifications) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesNotificationsByRole struct {
	Roles []string `pulumi:"roles"`
	State *string  `pulumi:"state"`
}





type SecurityContactPropertiesNotificationsByRoleInput interface {
	pulumi.Input

	ToSecurityContactPropertiesNotificationsByRoleOutput() SecurityContactPropertiesNotificationsByRoleOutput
	ToSecurityContactPropertiesNotificationsByRoleOutputWithContext(context.Context) SecurityContactPropertiesNotificationsByRoleOutput
}

type SecurityContactPropertiesNotificationsByRoleArgs struct {
	Roles pulumi.StringArrayInput `pulumi:"roles"`
	State pulumi.StringPtrInput   `pulumi:"state"`
}

func (SecurityContactPropertiesNotificationsByRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesNotificationsByRole)(nil)).Elem()
}

func (i SecurityContactPropertiesNotificationsByRoleArgs) ToSecurityContactPropertiesNotificationsByRoleOutput() SecurityContactPropertiesNotificationsByRoleOutput {
	return i.ToSecurityContactPropertiesNotificationsByRoleOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesNotificationsByRoleArgs) ToSecurityContactPropertiesNotificationsByRoleOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesNotificationsByRoleOutput)
}

func (i SecurityContactPropertiesNotificationsByRoleArgs) ToSecurityContactPropertiesNotificationsByRolePtrOutput() SecurityContactPropertiesNotificationsByRolePtrOutput {
	return i.ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesNotificationsByRoleArgs) ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesNotificationsByRoleOutput).ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(ctx)
}









type SecurityContactPropertiesNotificationsByRolePtrInput interface {
	pulumi.Input

	ToSecurityContactPropertiesNotificationsByRolePtrOutput() SecurityContactPropertiesNotificationsByRolePtrOutput
	ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(context.Context) SecurityContactPropertiesNotificationsByRolePtrOutput
}

type securityContactPropertiesNotificationsByRolePtrType SecurityContactPropertiesNotificationsByRoleArgs

func SecurityContactPropertiesNotificationsByRolePtr(v *SecurityContactPropertiesNotificationsByRoleArgs) SecurityContactPropertiesNotificationsByRolePtrInput {
	return (*securityContactPropertiesNotificationsByRolePtrType)(v)
}

func (*securityContactPropertiesNotificationsByRolePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesNotificationsByRole)(nil)).Elem()
}

func (i *securityContactPropertiesNotificationsByRolePtrType) ToSecurityContactPropertiesNotificationsByRolePtrOutput() SecurityContactPropertiesNotificationsByRolePtrOutput {
	return i.ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(context.Background())
}

func (i *securityContactPropertiesNotificationsByRolePtrType) ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesNotificationsByRolePtrOutput)
}

type SecurityContactPropertiesNotificationsByRoleOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesNotificationsByRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesNotificationsByRole)(nil)).Elem()
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) ToSecurityContactPropertiesNotificationsByRoleOutput() SecurityContactPropertiesNotificationsByRoleOutput {
	return o
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) ToSecurityContactPropertiesNotificationsByRoleOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRoleOutput {
	return o
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) ToSecurityContactPropertiesNotificationsByRolePtrOutput() SecurityContactPropertiesNotificationsByRolePtrOutput {
	return o.ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(context.Background())
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityContactPropertiesNotificationsByRole) *SecurityContactPropertiesNotificationsByRole {
		return &v
	}).(SecurityContactPropertiesNotificationsByRolePtrOutput)
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityContactPropertiesNotificationsByRole) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesNotificationsByRole) *string { return v.State }).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesNotificationsByRolePtrOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesNotificationsByRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesNotificationsByRole)(nil)).Elem()
}

func (o SecurityContactPropertiesNotificationsByRolePtrOutput) ToSecurityContactPropertiesNotificationsByRolePtrOutput() SecurityContactPropertiesNotificationsByRolePtrOutput {
	return o
}

func (o SecurityContactPropertiesNotificationsByRolePtrOutput) ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRolePtrOutput {
	return o
}

func (o SecurityContactPropertiesNotificationsByRolePtrOutput) Elem() SecurityContactPropertiesNotificationsByRoleOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesNotificationsByRole) SecurityContactPropertiesNotificationsByRole {
		if v != nil {
			return *v
		}
		var ret SecurityContactPropertiesNotificationsByRole
		return ret
	}).(SecurityContactPropertiesNotificationsByRoleOutput)
}

func (o SecurityContactPropertiesNotificationsByRolePtrOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesNotificationsByRole) []string {
		if v == nil {
			return nil
		}
		return v.Roles
	}).(pulumi.StringArrayOutput)
}

func (o SecurityContactPropertiesNotificationsByRolePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesNotificationsByRole) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesResponseAlertNotifications struct {
	MinimalSeverity *string `pulumi:"minimalSeverity"`
	State           *string `pulumi:"state"`
}

type SecurityContactPropertiesResponseAlertNotificationsOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesResponseAlertNotificationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesResponseAlertNotifications)(nil)).Elem()
}

func (o SecurityContactPropertiesResponseAlertNotificationsOutput) ToSecurityContactPropertiesResponseAlertNotificationsOutput() SecurityContactPropertiesResponseAlertNotificationsOutput {
	return o
}

func (o SecurityContactPropertiesResponseAlertNotificationsOutput) ToSecurityContactPropertiesResponseAlertNotificationsOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseAlertNotificationsOutput {
	return o
}

func (o SecurityContactPropertiesResponseAlertNotificationsOutput) MinimalSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesResponseAlertNotifications) *string { return v.MinimalSeverity }).(pulumi.StringPtrOutput)
}

func (o SecurityContactPropertiesResponseAlertNotificationsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesResponseAlertNotifications) *string { return v.State }).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesResponseAlertNotificationsPtrOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesResponseAlertNotificationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesResponseAlertNotifications)(nil)).Elem()
}

func (o SecurityContactPropertiesResponseAlertNotificationsPtrOutput) ToSecurityContactPropertiesResponseAlertNotificationsPtrOutput() SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return o
}

func (o SecurityContactPropertiesResponseAlertNotificationsPtrOutput) ToSecurityContactPropertiesResponseAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return o
}

func (o SecurityContactPropertiesResponseAlertNotificationsPtrOutput) Elem() SecurityContactPropertiesResponseAlertNotificationsOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseAlertNotifications) SecurityContactPropertiesResponseAlertNotifications {
		if v != nil {
			return *v
		}
		var ret SecurityContactPropertiesResponseAlertNotifications
		return ret
	}).(SecurityContactPropertiesResponseAlertNotificationsOutput)
}

func (o SecurityContactPropertiesResponseAlertNotificationsPtrOutput) MinimalSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseAlertNotifications) *string {
		if v == nil {
			return nil
		}
		return v.MinimalSeverity
	}).(pulumi.StringPtrOutput)
}

func (o SecurityContactPropertiesResponseAlertNotificationsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseAlertNotifications) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesResponseNotificationsByRole struct {
	Roles []string `pulumi:"roles"`
	State *string  `pulumi:"state"`
}

type SecurityContactPropertiesResponseNotificationsByRoleOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesResponseNotificationsByRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesResponseNotificationsByRole)(nil)).Elem()
}

func (o SecurityContactPropertiesResponseNotificationsByRoleOutput) ToSecurityContactPropertiesResponseNotificationsByRoleOutput() SecurityContactPropertiesResponseNotificationsByRoleOutput {
	return o
}

func (o SecurityContactPropertiesResponseNotificationsByRoleOutput) ToSecurityContactPropertiesResponseNotificationsByRoleOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseNotificationsByRoleOutput {
	return o
}

func (o SecurityContactPropertiesResponseNotificationsByRoleOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityContactPropertiesResponseNotificationsByRole) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o SecurityContactPropertiesResponseNotificationsByRoleOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesResponseNotificationsByRole) *string { return v.State }).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesResponseNotificationsByRolePtrOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesResponseNotificationsByRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesResponseNotificationsByRole)(nil)).Elem()
}

func (o SecurityContactPropertiesResponseNotificationsByRolePtrOutput) ToSecurityContactPropertiesResponseNotificationsByRolePtrOutput() SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return o
}

func (o SecurityContactPropertiesResponseNotificationsByRolePtrOutput) ToSecurityContactPropertiesResponseNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return o
}

func (o SecurityContactPropertiesResponseNotificationsByRolePtrOutput) Elem() SecurityContactPropertiesResponseNotificationsByRoleOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseNotificationsByRole) SecurityContactPropertiesResponseNotificationsByRole {
		if v != nil {
			return *v
		}
		var ret SecurityContactPropertiesResponseNotificationsByRole
		return ret
	}).(SecurityContactPropertiesResponseNotificationsByRoleOutput)
}

func (o SecurityContactPropertiesResponseNotificationsByRolePtrOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseNotificationsByRole) []string {
		if v == nil {
			return nil
		}
		return v.Roles
	}).(pulumi.StringArrayOutput)
}

func (o SecurityContactPropertiesResponseNotificationsByRolePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseNotificationsByRole) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalProperties struct {
	ApplicationId *string `pulumi:"applicationId"`
	Secret        *string `pulumi:"secret"`
}





type ServicePrincipalPropertiesInput interface {
	pulumi.Input

	ToServicePrincipalPropertiesOutput() ServicePrincipalPropertiesOutput
	ToServicePrincipalPropertiesOutputWithContext(context.Context) ServicePrincipalPropertiesOutput
}

type ServicePrincipalPropertiesArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	Secret        pulumi.StringPtrInput `pulumi:"secret"`
}

func (ServicePrincipalPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalProperties)(nil)).Elem()
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesOutput() ServicePrincipalPropertiesOutput {
	return i.ToServicePrincipalPropertiesOutputWithContext(context.Background())
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesOutputWithContext(ctx context.Context) ServicePrincipalPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesOutput)
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return i.ToServicePrincipalPropertiesPtrOutputWithContext(context.Background())
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesOutput).ToServicePrincipalPropertiesPtrOutputWithContext(ctx)
}









type ServicePrincipalPropertiesPtrInput interface {
	pulumi.Input

	ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput
	ToServicePrincipalPropertiesPtrOutputWithContext(context.Context) ServicePrincipalPropertiesPtrOutput
}

type servicePrincipalPropertiesPtrType ServicePrincipalPropertiesArgs

func ServicePrincipalPropertiesPtr(v *ServicePrincipalPropertiesArgs) ServicePrincipalPropertiesPtrInput {
	return (*servicePrincipalPropertiesPtrType)(v)
}

func (*servicePrincipalPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalProperties)(nil)).Elem()
}

func (i *servicePrincipalPropertiesPtrType) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return i.ToServicePrincipalPropertiesPtrOutputWithContext(context.Background())
}

func (i *servicePrincipalPropertiesPtrType) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesPtrOutput)
}

type ServicePrincipalPropertiesOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalProperties)(nil)).Elem()
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesOutput() ServicePrincipalPropertiesOutput {
	return o
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesOutputWithContext(ctx context.Context) ServicePrincipalPropertiesOutput {
	return o
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return o.ToServicePrincipalPropertiesPtrOutputWithContext(context.Background())
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePrincipalProperties) *ServicePrincipalProperties {
		return &v
	}).(ServicePrincipalPropertiesPtrOutput)
}

func (o ServicePrincipalPropertiesOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalProperties) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalPropertiesOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalProperties) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ServicePrincipalPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalProperties)(nil)).Elem()
}

func (o ServicePrincipalPropertiesPtrOutput) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return o
}

func (o ServicePrincipalPropertiesPtrOutput) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return o
}

func (o ServicePrincipalPropertiesPtrOutput) Elem() ServicePrincipalPropertiesOutput {
	return o.ApplyT(func(v *ServicePrincipalProperties) ServicePrincipalProperties {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalProperties
		return ret
	}).(ServicePrincipalPropertiesOutput)
}

func (o ServicePrincipalPropertiesPtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalPropertiesPtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalProperties) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalPropertiesResponse struct {
	ApplicationId *string `pulumi:"applicationId"`
	Secret        *string `pulumi:"secret"`
}

type ServicePrincipalPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalPropertiesResponse)(nil)).Elem()
}

func (o ServicePrincipalPropertiesResponseOutput) ToServicePrincipalPropertiesResponseOutput() ServicePrincipalPropertiesResponseOutput {
	return o
}

func (o ServicePrincipalPropertiesResponseOutput) ToServicePrincipalPropertiesResponseOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponseOutput {
	return o
}

func (o ServicePrincipalPropertiesResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalPropertiesResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalPropertiesResponseOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalPropertiesResponse) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ServicePrincipalPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalPropertiesResponse)(nil)).Elem()
}

func (o ServicePrincipalPropertiesResponsePtrOutput) ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput {
	return o
}

func (o ServicePrincipalPropertiesResponsePtrOutput) ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponsePtrOutput {
	return o
}

func (o ServicePrincipalPropertiesResponsePtrOutput) Elem() ServicePrincipalPropertiesResponseOutput {
	return o.ApplyT(func(v *ServicePrincipalPropertiesResponse) ServicePrincipalPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalPropertiesResponse
		return ret
	}).(ServicePrincipalPropertiesResponseOutput)
}

func (o ServicePrincipalPropertiesResponsePtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalPropertiesResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridComputeSettingsPropertiesOutput{})
	pulumi.RegisterOutputType(HybridComputeSettingsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(HybridComputeSettingsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HybridComputeSettingsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ProxyServerPropertiesOutput{})
	pulumi.RegisterOutputType(ProxyServerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ProxyServerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ProxyServerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesAlertNotificationsOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesAlertNotificationsPtrOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesNotificationsByRoleOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesNotificationsByRolePtrOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesResponseAlertNotificationsOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesResponseAlertNotificationsPtrOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesResponseNotificationsByRoleOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesResponseNotificationsByRolePtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesResponsePtrOutput{})
}
