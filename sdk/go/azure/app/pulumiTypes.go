


package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AllowedAudiencesValidation struct {
	AllowedAudiences []string `pulumi:"allowedAudiences"`
}





type AllowedAudiencesValidationInput interface {
	pulumi.Input

	ToAllowedAudiencesValidationOutput() AllowedAudiencesValidationOutput
	ToAllowedAudiencesValidationOutputWithContext(context.Context) AllowedAudiencesValidationOutput
}

type AllowedAudiencesValidationArgs struct {
	AllowedAudiences pulumi.StringArrayInput `pulumi:"allowedAudiences"`
}

func (AllowedAudiencesValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedAudiencesValidation)(nil)).Elem()
}

func (i AllowedAudiencesValidationArgs) ToAllowedAudiencesValidationOutput() AllowedAudiencesValidationOutput {
	return i.ToAllowedAudiencesValidationOutputWithContext(context.Background())
}

func (i AllowedAudiencesValidationArgs) ToAllowedAudiencesValidationOutputWithContext(ctx context.Context) AllowedAudiencesValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedAudiencesValidationOutput)
}

func (i AllowedAudiencesValidationArgs) ToAllowedAudiencesValidationPtrOutput() AllowedAudiencesValidationPtrOutput {
	return i.ToAllowedAudiencesValidationPtrOutputWithContext(context.Background())
}

func (i AllowedAudiencesValidationArgs) ToAllowedAudiencesValidationPtrOutputWithContext(ctx context.Context) AllowedAudiencesValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedAudiencesValidationOutput).ToAllowedAudiencesValidationPtrOutputWithContext(ctx)
}









type AllowedAudiencesValidationPtrInput interface {
	pulumi.Input

	ToAllowedAudiencesValidationPtrOutput() AllowedAudiencesValidationPtrOutput
	ToAllowedAudiencesValidationPtrOutputWithContext(context.Context) AllowedAudiencesValidationPtrOutput
}

type allowedAudiencesValidationPtrType AllowedAudiencesValidationArgs

func AllowedAudiencesValidationPtr(v *AllowedAudiencesValidationArgs) AllowedAudiencesValidationPtrInput {
	return (*allowedAudiencesValidationPtrType)(v)
}

func (*allowedAudiencesValidationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedAudiencesValidation)(nil)).Elem()
}

func (i *allowedAudiencesValidationPtrType) ToAllowedAudiencesValidationPtrOutput() AllowedAudiencesValidationPtrOutput {
	return i.ToAllowedAudiencesValidationPtrOutputWithContext(context.Background())
}

func (i *allowedAudiencesValidationPtrType) ToAllowedAudiencesValidationPtrOutputWithContext(ctx context.Context) AllowedAudiencesValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedAudiencesValidationPtrOutput)
}

type AllowedAudiencesValidationOutput struct{ *pulumi.OutputState }

func (AllowedAudiencesValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedAudiencesValidation)(nil)).Elem()
}

func (o AllowedAudiencesValidationOutput) ToAllowedAudiencesValidationOutput() AllowedAudiencesValidationOutput {
	return o
}

func (o AllowedAudiencesValidationOutput) ToAllowedAudiencesValidationOutputWithContext(ctx context.Context) AllowedAudiencesValidationOutput {
	return o
}

func (o AllowedAudiencesValidationOutput) ToAllowedAudiencesValidationPtrOutput() AllowedAudiencesValidationPtrOutput {
	return o.ToAllowedAudiencesValidationPtrOutputWithContext(context.Background())
}

func (o AllowedAudiencesValidationOutput) ToAllowedAudiencesValidationPtrOutputWithContext(ctx context.Context) AllowedAudiencesValidationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AllowedAudiencesValidation) *AllowedAudiencesValidation {
		return &v
	}).(AllowedAudiencesValidationPtrOutput)
}

func (o AllowedAudiencesValidationOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedAudiencesValidation) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

type AllowedAudiencesValidationPtrOutput struct{ *pulumi.OutputState }

func (AllowedAudiencesValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedAudiencesValidation)(nil)).Elem()
}

func (o AllowedAudiencesValidationPtrOutput) ToAllowedAudiencesValidationPtrOutput() AllowedAudiencesValidationPtrOutput {
	return o
}

func (o AllowedAudiencesValidationPtrOutput) ToAllowedAudiencesValidationPtrOutputWithContext(ctx context.Context) AllowedAudiencesValidationPtrOutput {
	return o
}

func (o AllowedAudiencesValidationPtrOutput) Elem() AllowedAudiencesValidationOutput {
	return o.ApplyT(func(v *AllowedAudiencesValidation) AllowedAudiencesValidation {
		if v != nil {
			return *v
		}
		var ret AllowedAudiencesValidation
		return ret
	}).(AllowedAudiencesValidationOutput)
}

func (o AllowedAudiencesValidationPtrOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedAudiencesValidation) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAudiences
	}).(pulumi.StringArrayOutput)
}

type AllowedAudiencesValidationResponse struct {
	AllowedAudiences []string `pulumi:"allowedAudiences"`
}

type AllowedAudiencesValidationResponseOutput struct{ *pulumi.OutputState }

func (AllowedAudiencesValidationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedAudiencesValidationResponse)(nil)).Elem()
}

func (o AllowedAudiencesValidationResponseOutput) ToAllowedAudiencesValidationResponseOutput() AllowedAudiencesValidationResponseOutput {
	return o
}

func (o AllowedAudiencesValidationResponseOutput) ToAllowedAudiencesValidationResponseOutputWithContext(ctx context.Context) AllowedAudiencesValidationResponseOutput {
	return o
}

func (o AllowedAudiencesValidationResponseOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedAudiencesValidationResponse) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

type AllowedAudiencesValidationResponsePtrOutput struct{ *pulumi.OutputState }

func (AllowedAudiencesValidationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedAudiencesValidationResponse)(nil)).Elem()
}

func (o AllowedAudiencesValidationResponsePtrOutput) ToAllowedAudiencesValidationResponsePtrOutput() AllowedAudiencesValidationResponsePtrOutput {
	return o
}

func (o AllowedAudiencesValidationResponsePtrOutput) ToAllowedAudiencesValidationResponsePtrOutputWithContext(ctx context.Context) AllowedAudiencesValidationResponsePtrOutput {
	return o
}

func (o AllowedAudiencesValidationResponsePtrOutput) Elem() AllowedAudiencesValidationResponseOutput {
	return o.ApplyT(func(v *AllowedAudiencesValidationResponse) AllowedAudiencesValidationResponse {
		if v != nil {
			return *v
		}
		var ret AllowedAudiencesValidationResponse
		return ret
	}).(AllowedAudiencesValidationResponseOutput)
}

func (o AllowedAudiencesValidationResponsePtrOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedAudiencesValidationResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAudiences
	}).(pulumi.StringArrayOutput)
}

type AllowedPrincipals struct {
	Groups     []string `pulumi:"groups"`
	Identities []string `pulumi:"identities"`
}





type AllowedPrincipalsInput interface {
	pulumi.Input

	ToAllowedPrincipalsOutput() AllowedPrincipalsOutput
	ToAllowedPrincipalsOutputWithContext(context.Context) AllowedPrincipalsOutput
}

type AllowedPrincipalsArgs struct {
	Groups     pulumi.StringArrayInput `pulumi:"groups"`
	Identities pulumi.StringArrayInput `pulumi:"identities"`
}

func (AllowedPrincipalsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedPrincipals)(nil)).Elem()
}

func (i AllowedPrincipalsArgs) ToAllowedPrincipalsOutput() AllowedPrincipalsOutput {
	return i.ToAllowedPrincipalsOutputWithContext(context.Background())
}

func (i AllowedPrincipalsArgs) ToAllowedPrincipalsOutputWithContext(ctx context.Context) AllowedPrincipalsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedPrincipalsOutput)
}

func (i AllowedPrincipalsArgs) ToAllowedPrincipalsPtrOutput() AllowedPrincipalsPtrOutput {
	return i.ToAllowedPrincipalsPtrOutputWithContext(context.Background())
}

func (i AllowedPrincipalsArgs) ToAllowedPrincipalsPtrOutputWithContext(ctx context.Context) AllowedPrincipalsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedPrincipalsOutput).ToAllowedPrincipalsPtrOutputWithContext(ctx)
}









type AllowedPrincipalsPtrInput interface {
	pulumi.Input

	ToAllowedPrincipalsPtrOutput() AllowedPrincipalsPtrOutput
	ToAllowedPrincipalsPtrOutputWithContext(context.Context) AllowedPrincipalsPtrOutput
}

type allowedPrincipalsPtrType AllowedPrincipalsArgs

func AllowedPrincipalsPtr(v *AllowedPrincipalsArgs) AllowedPrincipalsPtrInput {
	return (*allowedPrincipalsPtrType)(v)
}

func (*allowedPrincipalsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedPrincipals)(nil)).Elem()
}

func (i *allowedPrincipalsPtrType) ToAllowedPrincipalsPtrOutput() AllowedPrincipalsPtrOutput {
	return i.ToAllowedPrincipalsPtrOutputWithContext(context.Background())
}

func (i *allowedPrincipalsPtrType) ToAllowedPrincipalsPtrOutputWithContext(ctx context.Context) AllowedPrincipalsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedPrincipalsPtrOutput)
}

type AllowedPrincipalsOutput struct{ *pulumi.OutputState }

func (AllowedPrincipalsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedPrincipals)(nil)).Elem()
}

func (o AllowedPrincipalsOutput) ToAllowedPrincipalsOutput() AllowedPrincipalsOutput {
	return o
}

func (o AllowedPrincipalsOutput) ToAllowedPrincipalsOutputWithContext(ctx context.Context) AllowedPrincipalsOutput {
	return o
}

func (o AllowedPrincipalsOutput) ToAllowedPrincipalsPtrOutput() AllowedPrincipalsPtrOutput {
	return o.ToAllowedPrincipalsPtrOutputWithContext(context.Background())
}

func (o AllowedPrincipalsOutput) ToAllowedPrincipalsPtrOutputWithContext(ctx context.Context) AllowedPrincipalsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AllowedPrincipals) *AllowedPrincipals {
		return &v
	}).(AllowedPrincipalsPtrOutput)
}

func (o AllowedPrincipalsOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedPrincipals) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o AllowedPrincipalsOutput) Identities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedPrincipals) []string { return v.Identities }).(pulumi.StringArrayOutput)
}

type AllowedPrincipalsPtrOutput struct{ *pulumi.OutputState }

func (AllowedPrincipalsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedPrincipals)(nil)).Elem()
}

func (o AllowedPrincipalsPtrOutput) ToAllowedPrincipalsPtrOutput() AllowedPrincipalsPtrOutput {
	return o
}

func (o AllowedPrincipalsPtrOutput) ToAllowedPrincipalsPtrOutputWithContext(ctx context.Context) AllowedPrincipalsPtrOutput {
	return o
}

func (o AllowedPrincipalsPtrOutput) Elem() AllowedPrincipalsOutput {
	return o.ApplyT(func(v *AllowedPrincipals) AllowedPrincipals {
		if v != nil {
			return *v
		}
		var ret AllowedPrincipals
		return ret
	}).(AllowedPrincipalsOutput)
}

func (o AllowedPrincipalsPtrOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedPrincipals) []string {
		if v == nil {
			return nil
		}
		return v.Groups
	}).(pulumi.StringArrayOutput)
}

func (o AllowedPrincipalsPtrOutput) Identities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedPrincipals) []string {
		if v == nil {
			return nil
		}
		return v.Identities
	}).(pulumi.StringArrayOutput)
}

type AllowedPrincipalsResponse struct {
	Groups     []string `pulumi:"groups"`
	Identities []string `pulumi:"identities"`
}

type AllowedPrincipalsResponseOutput struct{ *pulumi.OutputState }

func (AllowedPrincipalsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedPrincipalsResponse)(nil)).Elem()
}

func (o AllowedPrincipalsResponseOutput) ToAllowedPrincipalsResponseOutput() AllowedPrincipalsResponseOutput {
	return o
}

func (o AllowedPrincipalsResponseOutput) ToAllowedPrincipalsResponseOutputWithContext(ctx context.Context) AllowedPrincipalsResponseOutput {
	return o
}

func (o AllowedPrincipalsResponseOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedPrincipalsResponse) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o AllowedPrincipalsResponseOutput) Identities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedPrincipalsResponse) []string { return v.Identities }).(pulumi.StringArrayOutput)
}

type AllowedPrincipalsResponsePtrOutput struct{ *pulumi.OutputState }

func (AllowedPrincipalsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedPrincipalsResponse)(nil)).Elem()
}

func (o AllowedPrincipalsResponsePtrOutput) ToAllowedPrincipalsResponsePtrOutput() AllowedPrincipalsResponsePtrOutput {
	return o
}

func (o AllowedPrincipalsResponsePtrOutput) ToAllowedPrincipalsResponsePtrOutputWithContext(ctx context.Context) AllowedPrincipalsResponsePtrOutput {
	return o
}

func (o AllowedPrincipalsResponsePtrOutput) Elem() AllowedPrincipalsResponseOutput {
	return o.ApplyT(func(v *AllowedPrincipalsResponse) AllowedPrincipalsResponse {
		if v != nil {
			return *v
		}
		var ret AllowedPrincipalsResponse
		return ret
	}).(AllowedPrincipalsResponseOutput)
}

func (o AllowedPrincipalsResponsePtrOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedPrincipalsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Groups
	}).(pulumi.StringArrayOutput)
}

func (o AllowedPrincipalsResponsePtrOutput) Identities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedPrincipalsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Identities
	}).(pulumi.StringArrayOutput)
}

type AppLogsConfiguration struct {
	Destination               *string                    `pulumi:"destination"`
	LogAnalyticsConfiguration *LogAnalyticsConfiguration `pulumi:"logAnalyticsConfiguration"`
}





type AppLogsConfigurationInput interface {
	pulumi.Input

	ToAppLogsConfigurationOutput() AppLogsConfigurationOutput
	ToAppLogsConfigurationOutputWithContext(context.Context) AppLogsConfigurationOutput
}

type AppLogsConfigurationArgs struct {
	Destination               pulumi.StringPtrInput             `pulumi:"destination"`
	LogAnalyticsConfiguration LogAnalyticsConfigurationPtrInput `pulumi:"logAnalyticsConfiguration"`
}

func (AppLogsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLogsConfiguration)(nil)).Elem()
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationOutput() AppLogsConfigurationOutput {
	return i.ToAppLogsConfigurationOutputWithContext(context.Background())
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationOutputWithContext(ctx context.Context) AppLogsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLogsConfigurationOutput)
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return i.ToAppLogsConfigurationPtrOutputWithContext(context.Background())
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLogsConfigurationOutput).ToAppLogsConfigurationPtrOutputWithContext(ctx)
}









type AppLogsConfigurationPtrInput interface {
	pulumi.Input

	ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput
	ToAppLogsConfigurationPtrOutputWithContext(context.Context) AppLogsConfigurationPtrOutput
}

type appLogsConfigurationPtrType AppLogsConfigurationArgs

func AppLogsConfigurationPtr(v *AppLogsConfigurationArgs) AppLogsConfigurationPtrInput {
	return (*appLogsConfigurationPtrType)(v)
}

func (*appLogsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLogsConfiguration)(nil)).Elem()
}

func (i *appLogsConfigurationPtrType) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return i.ToAppLogsConfigurationPtrOutputWithContext(context.Background())
}

func (i *appLogsConfigurationPtrType) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLogsConfigurationPtrOutput)
}

type AppLogsConfigurationOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLogsConfiguration)(nil)).Elem()
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationOutput() AppLogsConfigurationOutput {
	return o
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationOutputWithContext(ctx context.Context) AppLogsConfigurationOutput {
	return o
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return o.ToAppLogsConfigurationPtrOutputWithContext(context.Background())
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppLogsConfiguration) *AppLogsConfiguration {
		return &v
	}).(AppLogsConfigurationPtrOutput)
}

func (o AppLogsConfigurationOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppLogsConfiguration) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationPtrOutput {
	return o.ApplyT(func(v AppLogsConfiguration) *LogAnalyticsConfiguration { return v.LogAnalyticsConfiguration }).(LogAnalyticsConfigurationPtrOutput)
}

type AppLogsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLogsConfiguration)(nil)).Elem()
}

func (o AppLogsConfigurationPtrOutput) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return o
}

func (o AppLogsConfigurationPtrOutput) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return o
}

func (o AppLogsConfigurationPtrOutput) Elem() AppLogsConfigurationOutput {
	return o.ApplyT(func(v *AppLogsConfiguration) AppLogsConfiguration {
		if v != nil {
			return *v
		}
		var ret AppLogsConfiguration
		return ret
	}).(AppLogsConfigurationOutput)
}

func (o AppLogsConfigurationPtrOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppLogsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationPtrOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationPtrOutput {
	return o.ApplyT(func(v *AppLogsConfiguration) *LogAnalyticsConfiguration {
		if v == nil {
			return nil
		}
		return v.LogAnalyticsConfiguration
	}).(LogAnalyticsConfigurationPtrOutput)
}

type AppLogsConfigurationResponse struct {
	Destination               *string                            `pulumi:"destination"`
	LogAnalyticsConfiguration *LogAnalyticsConfigurationResponse `pulumi:"logAnalyticsConfiguration"`
}

type AppLogsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLogsConfigurationResponse)(nil)).Elem()
}

func (o AppLogsConfigurationResponseOutput) ToAppLogsConfigurationResponseOutput() AppLogsConfigurationResponseOutput {
	return o
}

func (o AppLogsConfigurationResponseOutput) ToAppLogsConfigurationResponseOutputWithContext(ctx context.Context) AppLogsConfigurationResponseOutput {
	return o
}

func (o AppLogsConfigurationResponseOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppLogsConfigurationResponse) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationResponseOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AppLogsConfigurationResponse) *LogAnalyticsConfigurationResponse {
		return v.LogAnalyticsConfiguration
	}).(LogAnalyticsConfigurationResponsePtrOutput)
}

type AppLogsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLogsConfigurationResponse)(nil)).Elem()
}

func (o AppLogsConfigurationResponsePtrOutput) ToAppLogsConfigurationResponsePtrOutput() AppLogsConfigurationResponsePtrOutput {
	return o
}

func (o AppLogsConfigurationResponsePtrOutput) ToAppLogsConfigurationResponsePtrOutputWithContext(ctx context.Context) AppLogsConfigurationResponsePtrOutput {
	return o
}

func (o AppLogsConfigurationResponsePtrOutput) Elem() AppLogsConfigurationResponseOutput {
	return o.ApplyT(func(v *AppLogsConfigurationResponse) AppLogsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret AppLogsConfigurationResponse
		return ret
	}).(AppLogsConfigurationResponseOutput)
}

func (o AppLogsConfigurationResponsePtrOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppLogsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationResponsePtrOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AppLogsConfigurationResponse) *LogAnalyticsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LogAnalyticsConfiguration
	}).(LogAnalyticsConfigurationResponsePtrOutput)
}

type AppRegistration struct {
	AppId                *string `pulumi:"appId"`
	AppSecretSettingName *string `pulumi:"appSecretSettingName"`
}





type AppRegistrationInput interface {
	pulumi.Input

	ToAppRegistrationOutput() AppRegistrationOutput
	ToAppRegistrationOutputWithContext(context.Context) AppRegistrationOutput
}

type AppRegistrationArgs struct {
	AppId                pulumi.StringPtrInput `pulumi:"appId"`
	AppSecretSettingName pulumi.StringPtrInput `pulumi:"appSecretSettingName"`
}

func (AppRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppRegistration)(nil)).Elem()
}

func (i AppRegistrationArgs) ToAppRegistrationOutput() AppRegistrationOutput {
	return i.ToAppRegistrationOutputWithContext(context.Background())
}

func (i AppRegistrationArgs) ToAppRegistrationOutputWithContext(ctx context.Context) AppRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppRegistrationOutput)
}

func (i AppRegistrationArgs) ToAppRegistrationPtrOutput() AppRegistrationPtrOutput {
	return i.ToAppRegistrationPtrOutputWithContext(context.Background())
}

func (i AppRegistrationArgs) ToAppRegistrationPtrOutputWithContext(ctx context.Context) AppRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppRegistrationOutput).ToAppRegistrationPtrOutputWithContext(ctx)
}









type AppRegistrationPtrInput interface {
	pulumi.Input

	ToAppRegistrationPtrOutput() AppRegistrationPtrOutput
	ToAppRegistrationPtrOutputWithContext(context.Context) AppRegistrationPtrOutput
}

type appRegistrationPtrType AppRegistrationArgs

func AppRegistrationPtr(v *AppRegistrationArgs) AppRegistrationPtrInput {
	return (*appRegistrationPtrType)(v)
}

func (*appRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppRegistration)(nil)).Elem()
}

func (i *appRegistrationPtrType) ToAppRegistrationPtrOutput() AppRegistrationPtrOutput {
	return i.ToAppRegistrationPtrOutputWithContext(context.Background())
}

func (i *appRegistrationPtrType) ToAppRegistrationPtrOutputWithContext(ctx context.Context) AppRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppRegistrationPtrOutput)
}

type AppRegistrationOutput struct{ *pulumi.OutputState }

func (AppRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppRegistration)(nil)).Elem()
}

func (o AppRegistrationOutput) ToAppRegistrationOutput() AppRegistrationOutput {
	return o
}

func (o AppRegistrationOutput) ToAppRegistrationOutputWithContext(ctx context.Context) AppRegistrationOutput {
	return o
}

func (o AppRegistrationOutput) ToAppRegistrationPtrOutput() AppRegistrationPtrOutput {
	return o.ToAppRegistrationPtrOutputWithContext(context.Background())
}

func (o AppRegistrationOutput) ToAppRegistrationPtrOutputWithContext(ctx context.Context) AppRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppRegistration) *AppRegistration {
		return &v
	}).(AppRegistrationPtrOutput)
}

func (o AppRegistrationOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppRegistration) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o AppRegistrationOutput) AppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppRegistration) *string { return v.AppSecretSettingName }).(pulumi.StringPtrOutput)
}

type AppRegistrationPtrOutput struct{ *pulumi.OutputState }

func (AppRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppRegistration)(nil)).Elem()
}

func (o AppRegistrationPtrOutput) ToAppRegistrationPtrOutput() AppRegistrationPtrOutput {
	return o
}

func (o AppRegistrationPtrOutput) ToAppRegistrationPtrOutputWithContext(ctx context.Context) AppRegistrationPtrOutput {
	return o
}

func (o AppRegistrationPtrOutput) Elem() AppRegistrationOutput {
	return o.ApplyT(func(v *AppRegistration) AppRegistration {
		if v != nil {
			return *v
		}
		var ret AppRegistration
		return ret
	}).(AppRegistrationOutput)
}

func (o AppRegistrationPtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppRegistration) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o AppRegistrationPtrOutput) AppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppRegistration) *string {
		if v == nil {
			return nil
		}
		return v.AppSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type AppRegistrationResponse struct {
	AppId                *string `pulumi:"appId"`
	AppSecretSettingName *string `pulumi:"appSecretSettingName"`
}

type AppRegistrationResponseOutput struct{ *pulumi.OutputState }

func (AppRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppRegistrationResponse)(nil)).Elem()
}

func (o AppRegistrationResponseOutput) ToAppRegistrationResponseOutput() AppRegistrationResponseOutput {
	return o
}

func (o AppRegistrationResponseOutput) ToAppRegistrationResponseOutputWithContext(ctx context.Context) AppRegistrationResponseOutput {
	return o
}

func (o AppRegistrationResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppRegistrationResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o AppRegistrationResponseOutput) AppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppRegistrationResponse) *string { return v.AppSecretSettingName }).(pulumi.StringPtrOutput)
}

type AppRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (AppRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppRegistrationResponse)(nil)).Elem()
}

func (o AppRegistrationResponsePtrOutput) ToAppRegistrationResponsePtrOutput() AppRegistrationResponsePtrOutput {
	return o
}

func (o AppRegistrationResponsePtrOutput) ToAppRegistrationResponsePtrOutputWithContext(ctx context.Context) AppRegistrationResponsePtrOutput {
	return o
}

func (o AppRegistrationResponsePtrOutput) Elem() AppRegistrationResponseOutput {
	return o.ApplyT(func(v *AppRegistrationResponse) AppRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret AppRegistrationResponse
		return ret
	}).(AppRegistrationResponseOutput)
}

func (o AppRegistrationResponsePtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o AppRegistrationResponsePtrOutput) AppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type Apple struct {
	Enabled      *bool              `pulumi:"enabled"`
	Login        *LoginScopes       `pulumi:"login"`
	Registration *AppleRegistration `pulumi:"registration"`
}





type AppleInput interface {
	pulumi.Input

	ToAppleOutput() AppleOutput
	ToAppleOutputWithContext(context.Context) AppleOutput
}

type AppleArgs struct {
	Enabled      pulumi.BoolPtrInput       `pulumi:"enabled"`
	Login        LoginScopesPtrInput       `pulumi:"login"`
	Registration AppleRegistrationPtrInput `pulumi:"registration"`
}

func (AppleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Apple)(nil)).Elem()
}

func (i AppleArgs) ToAppleOutput() AppleOutput {
	return i.ToAppleOutputWithContext(context.Background())
}

func (i AppleArgs) ToAppleOutputWithContext(ctx context.Context) AppleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleOutput)
}

func (i AppleArgs) ToApplePtrOutput() ApplePtrOutput {
	return i.ToApplePtrOutputWithContext(context.Background())
}

func (i AppleArgs) ToApplePtrOutputWithContext(ctx context.Context) ApplePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleOutput).ToApplePtrOutputWithContext(ctx)
}









type ApplePtrInput interface {
	pulumi.Input

	ToApplePtrOutput() ApplePtrOutput
	ToApplePtrOutputWithContext(context.Context) ApplePtrOutput
}

type applePtrType AppleArgs

func ApplePtr(v *AppleArgs) ApplePtrInput {
	return (*applePtrType)(v)
}

func (*applePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Apple)(nil)).Elem()
}

func (i *applePtrType) ToApplePtrOutput() ApplePtrOutput {
	return i.ToApplePtrOutputWithContext(context.Background())
}

func (i *applePtrType) ToApplePtrOutputWithContext(ctx context.Context) ApplePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplePtrOutput)
}

type AppleOutput struct{ *pulumi.OutputState }

func (AppleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Apple)(nil)).Elem()
}

func (o AppleOutput) ToAppleOutput() AppleOutput {
	return o
}

func (o AppleOutput) ToAppleOutputWithContext(ctx context.Context) AppleOutput {
	return o
}

func (o AppleOutput) ToApplePtrOutput() ApplePtrOutput {
	return o.ToApplePtrOutputWithContext(context.Background())
}

func (o AppleOutput) ToApplePtrOutputWithContext(ctx context.Context) ApplePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Apple) *Apple {
		return &v
	}).(ApplePtrOutput)
}

func (o AppleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Apple) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AppleOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v Apple) *LoginScopes { return v.Login }).(LoginScopesPtrOutput)
}

func (o AppleOutput) Registration() AppleRegistrationPtrOutput {
	return o.ApplyT(func(v Apple) *AppleRegistration { return v.Registration }).(AppleRegistrationPtrOutput)
}

type ApplePtrOutput struct{ *pulumi.OutputState }

func (ApplePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Apple)(nil)).Elem()
}

func (o ApplePtrOutput) ToApplePtrOutput() ApplePtrOutput {
	return o
}

func (o ApplePtrOutput) ToApplePtrOutputWithContext(ctx context.Context) ApplePtrOutput {
	return o
}

func (o ApplePtrOutput) Elem() AppleOutput {
	return o.ApplyT(func(v *Apple) Apple {
		if v != nil {
			return *v
		}
		var ret Apple
		return ret
	}).(AppleOutput)
}

func (o ApplePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Apple) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApplePtrOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v *Apple) *LoginScopes {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesPtrOutput)
}

func (o ApplePtrOutput) Registration() AppleRegistrationPtrOutput {
	return o.ApplyT(func(v *Apple) *AppleRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AppleRegistrationPtrOutput)
}

type AppleRegistration struct {
	ClientId                *string `pulumi:"clientId"`
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
}





type AppleRegistrationInput interface {
	pulumi.Input

	ToAppleRegistrationOutput() AppleRegistrationOutput
	ToAppleRegistrationOutputWithContext(context.Context) AppleRegistrationOutput
}

type AppleRegistrationArgs struct {
	ClientId                pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecretSettingName pulumi.StringPtrInput `pulumi:"clientSecretSettingName"`
}

func (AppleRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppleRegistration)(nil)).Elem()
}

func (i AppleRegistrationArgs) ToAppleRegistrationOutput() AppleRegistrationOutput {
	return i.ToAppleRegistrationOutputWithContext(context.Background())
}

func (i AppleRegistrationArgs) ToAppleRegistrationOutputWithContext(ctx context.Context) AppleRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleRegistrationOutput)
}

func (i AppleRegistrationArgs) ToAppleRegistrationPtrOutput() AppleRegistrationPtrOutput {
	return i.ToAppleRegistrationPtrOutputWithContext(context.Background())
}

func (i AppleRegistrationArgs) ToAppleRegistrationPtrOutputWithContext(ctx context.Context) AppleRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleRegistrationOutput).ToAppleRegistrationPtrOutputWithContext(ctx)
}









type AppleRegistrationPtrInput interface {
	pulumi.Input

	ToAppleRegistrationPtrOutput() AppleRegistrationPtrOutput
	ToAppleRegistrationPtrOutputWithContext(context.Context) AppleRegistrationPtrOutput
}

type appleRegistrationPtrType AppleRegistrationArgs

func AppleRegistrationPtr(v *AppleRegistrationArgs) AppleRegistrationPtrInput {
	return (*appleRegistrationPtrType)(v)
}

func (*appleRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleRegistration)(nil)).Elem()
}

func (i *appleRegistrationPtrType) ToAppleRegistrationPtrOutput() AppleRegistrationPtrOutput {
	return i.ToAppleRegistrationPtrOutputWithContext(context.Background())
}

func (i *appleRegistrationPtrType) ToAppleRegistrationPtrOutputWithContext(ctx context.Context) AppleRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleRegistrationPtrOutput)
}

type AppleRegistrationOutput struct{ *pulumi.OutputState }

func (AppleRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppleRegistration)(nil)).Elem()
}

func (o AppleRegistrationOutput) ToAppleRegistrationOutput() AppleRegistrationOutput {
	return o
}

func (o AppleRegistrationOutput) ToAppleRegistrationOutputWithContext(ctx context.Context) AppleRegistrationOutput {
	return o
}

func (o AppleRegistrationOutput) ToAppleRegistrationPtrOutput() AppleRegistrationPtrOutput {
	return o.ToAppleRegistrationPtrOutputWithContext(context.Background())
}

func (o AppleRegistrationOutput) ToAppleRegistrationPtrOutputWithContext(ctx context.Context) AppleRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppleRegistration) *AppleRegistration {
		return &v
	}).(AppleRegistrationPtrOutput)
}

func (o AppleRegistrationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppleRegistration) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AppleRegistrationOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppleRegistration) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

type AppleRegistrationPtrOutput struct{ *pulumi.OutputState }

func (AppleRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleRegistration)(nil)).Elem()
}

func (o AppleRegistrationPtrOutput) ToAppleRegistrationPtrOutput() AppleRegistrationPtrOutput {
	return o
}

func (o AppleRegistrationPtrOutput) ToAppleRegistrationPtrOutputWithContext(ctx context.Context) AppleRegistrationPtrOutput {
	return o
}

func (o AppleRegistrationPtrOutput) Elem() AppleRegistrationOutput {
	return o.ApplyT(func(v *AppleRegistration) AppleRegistration {
		if v != nil {
			return *v
		}
		var ret AppleRegistration
		return ret
	}).(AppleRegistrationOutput)
}

func (o AppleRegistrationPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppleRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AppleRegistrationPtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppleRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type AppleRegistrationResponse struct {
	ClientId                *string `pulumi:"clientId"`
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
}

type AppleRegistrationResponseOutput struct{ *pulumi.OutputState }

func (AppleRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppleRegistrationResponse)(nil)).Elem()
}

func (o AppleRegistrationResponseOutput) ToAppleRegistrationResponseOutput() AppleRegistrationResponseOutput {
	return o
}

func (o AppleRegistrationResponseOutput) ToAppleRegistrationResponseOutputWithContext(ctx context.Context) AppleRegistrationResponseOutput {
	return o
}

func (o AppleRegistrationResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppleRegistrationResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AppleRegistrationResponseOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppleRegistrationResponse) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

type AppleRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (AppleRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleRegistrationResponse)(nil)).Elem()
}

func (o AppleRegistrationResponsePtrOutput) ToAppleRegistrationResponsePtrOutput() AppleRegistrationResponsePtrOutput {
	return o
}

func (o AppleRegistrationResponsePtrOutput) ToAppleRegistrationResponsePtrOutputWithContext(ctx context.Context) AppleRegistrationResponsePtrOutput {
	return o
}

func (o AppleRegistrationResponsePtrOutput) Elem() AppleRegistrationResponseOutput {
	return o.ApplyT(func(v *AppleRegistrationResponse) AppleRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret AppleRegistrationResponse
		return ret
	}).(AppleRegistrationResponseOutput)
}

func (o AppleRegistrationResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppleRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AppleRegistrationResponsePtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppleRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type AppleResponse struct {
	Enabled      *bool                      `pulumi:"enabled"`
	Login        *LoginScopesResponse       `pulumi:"login"`
	Registration *AppleRegistrationResponse `pulumi:"registration"`
}

type AppleResponseOutput struct{ *pulumi.OutputState }

func (AppleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppleResponse)(nil)).Elem()
}

func (o AppleResponseOutput) ToAppleResponseOutput() AppleResponseOutput {
	return o
}

func (o AppleResponseOutput) ToAppleResponseOutputWithContext(ctx context.Context) AppleResponseOutput {
	return o
}

func (o AppleResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppleResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AppleResponseOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v AppleResponse) *LoginScopesResponse { return v.Login }).(LoginScopesResponsePtrOutput)
}

func (o AppleResponseOutput) Registration() AppleRegistrationResponsePtrOutput {
	return o.ApplyT(func(v AppleResponse) *AppleRegistrationResponse { return v.Registration }).(AppleRegistrationResponsePtrOutput)
}

type AppleResponsePtrOutput struct{ *pulumi.OutputState }

func (AppleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleResponse)(nil)).Elem()
}

func (o AppleResponsePtrOutput) ToAppleResponsePtrOutput() AppleResponsePtrOutput {
	return o
}

func (o AppleResponsePtrOutput) ToAppleResponsePtrOutputWithContext(ctx context.Context) AppleResponsePtrOutput {
	return o
}

func (o AppleResponsePtrOutput) Elem() AppleResponseOutput {
	return o.ApplyT(func(v *AppleResponse) AppleResponse {
		if v != nil {
			return *v
		}
		var ret AppleResponse
		return ret
	}).(AppleResponseOutput)
}

func (o AppleResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppleResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AppleResponsePtrOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v *AppleResponse) *LoginScopesResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesResponsePtrOutput)
}

func (o AppleResponsePtrOutput) Registration() AppleRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *AppleResponse) *AppleRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AppleRegistrationResponsePtrOutput)
}

type AuthPlatform struct {
	Enabled        *bool   `pulumi:"enabled"`
	RuntimeVersion *string `pulumi:"runtimeVersion"`
}





type AuthPlatformInput interface {
	pulumi.Input

	ToAuthPlatformOutput() AuthPlatformOutput
	ToAuthPlatformOutputWithContext(context.Context) AuthPlatformOutput
}

type AuthPlatformArgs struct {
	Enabled        pulumi.BoolPtrInput   `pulumi:"enabled"`
	RuntimeVersion pulumi.StringPtrInput `pulumi:"runtimeVersion"`
}

func (AuthPlatformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthPlatform)(nil)).Elem()
}

func (i AuthPlatformArgs) ToAuthPlatformOutput() AuthPlatformOutput {
	return i.ToAuthPlatformOutputWithContext(context.Background())
}

func (i AuthPlatformArgs) ToAuthPlatformOutputWithContext(ctx context.Context) AuthPlatformOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthPlatformOutput)
}

func (i AuthPlatformArgs) ToAuthPlatformPtrOutput() AuthPlatformPtrOutput {
	return i.ToAuthPlatformPtrOutputWithContext(context.Background())
}

func (i AuthPlatformArgs) ToAuthPlatformPtrOutputWithContext(ctx context.Context) AuthPlatformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthPlatformOutput).ToAuthPlatformPtrOutputWithContext(ctx)
}









type AuthPlatformPtrInput interface {
	pulumi.Input

	ToAuthPlatformPtrOutput() AuthPlatformPtrOutput
	ToAuthPlatformPtrOutputWithContext(context.Context) AuthPlatformPtrOutput
}

type authPlatformPtrType AuthPlatformArgs

func AuthPlatformPtr(v *AuthPlatformArgs) AuthPlatformPtrInput {
	return (*authPlatformPtrType)(v)
}

func (*authPlatformPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthPlatform)(nil)).Elem()
}

func (i *authPlatformPtrType) ToAuthPlatformPtrOutput() AuthPlatformPtrOutput {
	return i.ToAuthPlatformPtrOutputWithContext(context.Background())
}

func (i *authPlatformPtrType) ToAuthPlatformPtrOutputWithContext(ctx context.Context) AuthPlatformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthPlatformPtrOutput)
}

type AuthPlatformOutput struct{ *pulumi.OutputState }

func (AuthPlatformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthPlatform)(nil)).Elem()
}

func (o AuthPlatformOutput) ToAuthPlatformOutput() AuthPlatformOutput {
	return o
}

func (o AuthPlatformOutput) ToAuthPlatformOutputWithContext(ctx context.Context) AuthPlatformOutput {
	return o
}

func (o AuthPlatformOutput) ToAuthPlatformPtrOutput() AuthPlatformPtrOutput {
	return o.ToAuthPlatformPtrOutputWithContext(context.Background())
}

func (o AuthPlatformOutput) ToAuthPlatformPtrOutputWithContext(ctx context.Context) AuthPlatformPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthPlatform) *AuthPlatform {
		return &v
	}).(AuthPlatformPtrOutput)
}

func (o AuthPlatformOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AuthPlatform) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AuthPlatformOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthPlatform) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type AuthPlatformPtrOutput struct{ *pulumi.OutputState }

func (AuthPlatformPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthPlatform)(nil)).Elem()
}

func (o AuthPlatformPtrOutput) ToAuthPlatformPtrOutput() AuthPlatformPtrOutput {
	return o
}

func (o AuthPlatformPtrOutput) ToAuthPlatformPtrOutputWithContext(ctx context.Context) AuthPlatformPtrOutput {
	return o
}

func (o AuthPlatformPtrOutput) Elem() AuthPlatformOutput {
	return o.ApplyT(func(v *AuthPlatform) AuthPlatform {
		if v != nil {
			return *v
		}
		var ret AuthPlatform
		return ret
	}).(AuthPlatformOutput)
}

func (o AuthPlatformPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthPlatform) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AuthPlatformPtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthPlatform) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type AuthPlatformResponse struct {
	Enabled        *bool   `pulumi:"enabled"`
	RuntimeVersion *string `pulumi:"runtimeVersion"`
}

type AuthPlatformResponseOutput struct{ *pulumi.OutputState }

func (AuthPlatformResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthPlatformResponse)(nil)).Elem()
}

func (o AuthPlatformResponseOutput) ToAuthPlatformResponseOutput() AuthPlatformResponseOutput {
	return o
}

func (o AuthPlatformResponseOutput) ToAuthPlatformResponseOutputWithContext(ctx context.Context) AuthPlatformResponseOutput {
	return o
}

func (o AuthPlatformResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AuthPlatformResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AuthPlatformResponseOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthPlatformResponse) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type AuthPlatformResponsePtrOutput struct{ *pulumi.OutputState }

func (AuthPlatformResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthPlatformResponse)(nil)).Elem()
}

func (o AuthPlatformResponsePtrOutput) ToAuthPlatformResponsePtrOutput() AuthPlatformResponsePtrOutput {
	return o
}

func (o AuthPlatformResponsePtrOutput) ToAuthPlatformResponsePtrOutputWithContext(ctx context.Context) AuthPlatformResponsePtrOutput {
	return o
}

func (o AuthPlatformResponsePtrOutput) Elem() AuthPlatformResponseOutput {
	return o.ApplyT(func(v *AuthPlatformResponse) AuthPlatformResponse {
		if v != nil {
			return *v
		}
		var ret AuthPlatformResponse
		return ret
	}).(AuthPlatformResponseOutput)
}

func (o AuthPlatformResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthPlatformResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AuthPlatformResponsePtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthPlatformResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type AzureActiveDirectory struct {
	Enabled           *bool                             `pulumi:"enabled"`
	IsAutoProvisioned *bool                             `pulumi:"isAutoProvisioned"`
	Login             *AzureActiveDirectoryLogin        `pulumi:"login"`
	Registration      *AzureActiveDirectoryRegistration `pulumi:"registration"`
	Validation        *AzureActiveDirectoryValidation   `pulumi:"validation"`
}





type AzureActiveDirectoryInput interface {
	pulumi.Input

	ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput
	ToAzureActiveDirectoryOutputWithContext(context.Context) AzureActiveDirectoryOutput
}

type AzureActiveDirectoryArgs struct {
	Enabled           pulumi.BoolPtrInput                      `pulumi:"enabled"`
	IsAutoProvisioned pulumi.BoolPtrInput                      `pulumi:"isAutoProvisioned"`
	Login             AzureActiveDirectoryLoginPtrInput        `pulumi:"login"`
	Registration      AzureActiveDirectoryRegistrationPtrInput `pulumi:"registration"`
	Validation        AzureActiveDirectoryValidationPtrInput   `pulumi:"validation"`
}

func (AzureActiveDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectory)(nil)).Elem()
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput {
	return i.ToAzureActiveDirectoryOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryOutputWithContext(ctx context.Context) AzureActiveDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryOutput)
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return i.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryOutput).ToAzureActiveDirectoryPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput
	ToAzureActiveDirectoryPtrOutputWithContext(context.Context) AzureActiveDirectoryPtrOutput
}

type azureActiveDirectoryPtrType AzureActiveDirectoryArgs

func AzureActiveDirectoryPtr(v *AzureActiveDirectoryArgs) AzureActiveDirectoryPtrInput {
	return (*azureActiveDirectoryPtrType)(v)
}

func (*azureActiveDirectoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectory)(nil)).Elem()
}

func (i *azureActiveDirectoryPtrType) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return i.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryPtrType) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryPtrOutput)
}

type AzureActiveDirectoryOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectory)(nil)).Elem()
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput {
	return o
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryOutputWithContext(ctx context.Context) AzureActiveDirectoryOutput {
	return o
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return o.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectory) *AzureActiveDirectory {
		return &v
	}).(AzureActiveDirectoryPtrOutput)
}

func (o AzureActiveDirectoryOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryOutput) IsAutoProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *bool { return v.IsAutoProvisioned }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryOutput) Login() AzureActiveDirectoryLoginPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *AzureActiveDirectoryLogin { return v.Login }).(AzureActiveDirectoryLoginPtrOutput)
}

func (o AzureActiveDirectoryOutput) Registration() AzureActiveDirectoryRegistrationPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *AzureActiveDirectoryRegistration { return v.Registration }).(AzureActiveDirectoryRegistrationPtrOutput)
}

func (o AzureActiveDirectoryOutput) Validation() AzureActiveDirectoryValidationPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *AzureActiveDirectoryValidation { return v.Validation }).(AzureActiveDirectoryValidationPtrOutput)
}

type AzureActiveDirectoryPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectory)(nil)).Elem()
}

func (o AzureActiveDirectoryPtrOutput) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return o
}

func (o AzureActiveDirectoryPtrOutput) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return o
}

func (o AzureActiveDirectoryPtrOutput) Elem() AzureActiveDirectoryOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) AzureActiveDirectory {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectory
		return ret
	}).(AzureActiveDirectoryOutput)
}

func (o AzureActiveDirectoryPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) IsAutoProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *bool {
		if v == nil {
			return nil
		}
		return v.IsAutoProvisioned
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) Login() AzureActiveDirectoryLoginPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *AzureActiveDirectoryLogin {
		if v == nil {
			return nil
		}
		return v.Login
	}).(AzureActiveDirectoryLoginPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) Registration() AzureActiveDirectoryRegistrationPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *AzureActiveDirectoryRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AzureActiveDirectoryRegistrationPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) Validation() AzureActiveDirectoryValidationPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *AzureActiveDirectoryValidation {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(AzureActiveDirectoryValidationPtrOutput)
}

type AzureActiveDirectoryLogin struct {
	DisableWWWAuthenticate *bool    `pulumi:"disableWWWAuthenticate"`
	LoginParameters        []string `pulumi:"loginParameters"`
}





type AzureActiveDirectoryLoginInput interface {
	pulumi.Input

	ToAzureActiveDirectoryLoginOutput() AzureActiveDirectoryLoginOutput
	ToAzureActiveDirectoryLoginOutputWithContext(context.Context) AzureActiveDirectoryLoginOutput
}

type AzureActiveDirectoryLoginArgs struct {
	DisableWWWAuthenticate pulumi.BoolPtrInput     `pulumi:"disableWWWAuthenticate"`
	LoginParameters        pulumi.StringArrayInput `pulumi:"loginParameters"`
}

func (AzureActiveDirectoryLoginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryLogin)(nil)).Elem()
}

func (i AzureActiveDirectoryLoginArgs) ToAzureActiveDirectoryLoginOutput() AzureActiveDirectoryLoginOutput {
	return i.ToAzureActiveDirectoryLoginOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryLoginArgs) ToAzureActiveDirectoryLoginOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryLoginOutput)
}

func (i AzureActiveDirectoryLoginArgs) ToAzureActiveDirectoryLoginPtrOutput() AzureActiveDirectoryLoginPtrOutput {
	return i.ToAzureActiveDirectoryLoginPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryLoginArgs) ToAzureActiveDirectoryLoginPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryLoginOutput).ToAzureActiveDirectoryLoginPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryLoginPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryLoginPtrOutput() AzureActiveDirectoryLoginPtrOutput
	ToAzureActiveDirectoryLoginPtrOutputWithContext(context.Context) AzureActiveDirectoryLoginPtrOutput
}

type azureActiveDirectoryLoginPtrType AzureActiveDirectoryLoginArgs

func AzureActiveDirectoryLoginPtr(v *AzureActiveDirectoryLoginArgs) AzureActiveDirectoryLoginPtrInput {
	return (*azureActiveDirectoryLoginPtrType)(v)
}

func (*azureActiveDirectoryLoginPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryLogin)(nil)).Elem()
}

func (i *azureActiveDirectoryLoginPtrType) ToAzureActiveDirectoryLoginPtrOutput() AzureActiveDirectoryLoginPtrOutput {
	return i.ToAzureActiveDirectoryLoginPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryLoginPtrType) ToAzureActiveDirectoryLoginPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryLoginPtrOutput)
}

type AzureActiveDirectoryLoginOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryLoginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryLogin)(nil)).Elem()
}

func (o AzureActiveDirectoryLoginOutput) ToAzureActiveDirectoryLoginOutput() AzureActiveDirectoryLoginOutput {
	return o
}

func (o AzureActiveDirectoryLoginOutput) ToAzureActiveDirectoryLoginOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginOutput {
	return o
}

func (o AzureActiveDirectoryLoginOutput) ToAzureActiveDirectoryLoginPtrOutput() AzureActiveDirectoryLoginPtrOutput {
	return o.ToAzureActiveDirectoryLoginPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryLoginOutput) ToAzureActiveDirectoryLoginPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectoryLogin) *AzureActiveDirectoryLogin {
		return &v
	}).(AzureActiveDirectoryLoginPtrOutput)
}

func (o AzureActiveDirectoryLoginOutput) DisableWWWAuthenticate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryLogin) *bool { return v.DisableWWWAuthenticate }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryLoginOutput) LoginParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureActiveDirectoryLogin) []string { return v.LoginParameters }).(pulumi.StringArrayOutput)
}

type AzureActiveDirectoryLoginPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryLoginPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryLogin)(nil)).Elem()
}

func (o AzureActiveDirectoryLoginPtrOutput) ToAzureActiveDirectoryLoginPtrOutput() AzureActiveDirectoryLoginPtrOutput {
	return o
}

func (o AzureActiveDirectoryLoginPtrOutput) ToAzureActiveDirectoryLoginPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginPtrOutput {
	return o
}

func (o AzureActiveDirectoryLoginPtrOutput) Elem() AzureActiveDirectoryLoginOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLogin) AzureActiveDirectoryLogin {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryLogin
		return ret
	}).(AzureActiveDirectoryLoginOutput)
}

func (o AzureActiveDirectoryLoginPtrOutput) DisableWWWAuthenticate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLogin) *bool {
		if v == nil {
			return nil
		}
		return v.DisableWWWAuthenticate
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryLoginPtrOutput) LoginParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLogin) []string {
		if v == nil {
			return nil
		}
		return v.LoginParameters
	}).(pulumi.StringArrayOutput)
}

type AzureActiveDirectoryLoginResponse struct {
	DisableWWWAuthenticate *bool    `pulumi:"disableWWWAuthenticate"`
	LoginParameters        []string `pulumi:"loginParameters"`
}

type AzureActiveDirectoryLoginResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryLoginResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryLoginResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryLoginResponseOutput) ToAzureActiveDirectoryLoginResponseOutput() AzureActiveDirectoryLoginResponseOutput {
	return o
}

func (o AzureActiveDirectoryLoginResponseOutput) ToAzureActiveDirectoryLoginResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginResponseOutput {
	return o
}

func (o AzureActiveDirectoryLoginResponseOutput) DisableWWWAuthenticate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryLoginResponse) *bool { return v.DisableWWWAuthenticate }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryLoginResponseOutput) LoginParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureActiveDirectoryLoginResponse) []string { return v.LoginParameters }).(pulumi.StringArrayOutput)
}

type AzureActiveDirectoryLoginResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryLoginResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryLoginResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryLoginResponsePtrOutput) ToAzureActiveDirectoryLoginResponsePtrOutput() AzureActiveDirectoryLoginResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryLoginResponsePtrOutput) ToAzureActiveDirectoryLoginResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryLoginResponsePtrOutput) Elem() AzureActiveDirectoryLoginResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLoginResponse) AzureActiveDirectoryLoginResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryLoginResponse
		return ret
	}).(AzureActiveDirectoryLoginResponseOutput)
}

func (o AzureActiveDirectoryLoginResponsePtrOutput) DisableWWWAuthenticate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLoginResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableWWWAuthenticate
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryLoginResponsePtrOutput) LoginParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLoginResponse) []string {
		if v == nil {
			return nil
		}
		return v.LoginParameters
	}).(pulumi.StringArrayOutput)
}

type AzureActiveDirectoryRegistration struct {
	ClientId                                      *string `pulumi:"clientId"`
	ClientSecretCertificateIssuer                 *string `pulumi:"clientSecretCertificateIssuer"`
	ClientSecretCertificateSubjectAlternativeName *string `pulumi:"clientSecretCertificateSubjectAlternativeName"`
	ClientSecretCertificateThumbprint             *string `pulumi:"clientSecretCertificateThumbprint"`
	ClientSecretSettingName                       *string `pulumi:"clientSecretSettingName"`
	OpenIdIssuer                                  *string `pulumi:"openIdIssuer"`
}





type AzureActiveDirectoryRegistrationInput interface {
	pulumi.Input

	ToAzureActiveDirectoryRegistrationOutput() AzureActiveDirectoryRegistrationOutput
	ToAzureActiveDirectoryRegistrationOutputWithContext(context.Context) AzureActiveDirectoryRegistrationOutput
}

type AzureActiveDirectoryRegistrationArgs struct {
	ClientId                                      pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecretCertificateIssuer                 pulumi.StringPtrInput `pulumi:"clientSecretCertificateIssuer"`
	ClientSecretCertificateSubjectAlternativeName pulumi.StringPtrInput `pulumi:"clientSecretCertificateSubjectAlternativeName"`
	ClientSecretCertificateThumbprint             pulumi.StringPtrInput `pulumi:"clientSecretCertificateThumbprint"`
	ClientSecretSettingName                       pulumi.StringPtrInput `pulumi:"clientSecretSettingName"`
	OpenIdIssuer                                  pulumi.StringPtrInput `pulumi:"openIdIssuer"`
}

func (AzureActiveDirectoryRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryRegistration)(nil)).Elem()
}

func (i AzureActiveDirectoryRegistrationArgs) ToAzureActiveDirectoryRegistrationOutput() AzureActiveDirectoryRegistrationOutput {
	return i.ToAzureActiveDirectoryRegistrationOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryRegistrationArgs) ToAzureActiveDirectoryRegistrationOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryRegistrationOutput)
}

func (i AzureActiveDirectoryRegistrationArgs) ToAzureActiveDirectoryRegistrationPtrOutput() AzureActiveDirectoryRegistrationPtrOutput {
	return i.ToAzureActiveDirectoryRegistrationPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryRegistrationArgs) ToAzureActiveDirectoryRegistrationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryRegistrationOutput).ToAzureActiveDirectoryRegistrationPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryRegistrationPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryRegistrationPtrOutput() AzureActiveDirectoryRegistrationPtrOutput
	ToAzureActiveDirectoryRegistrationPtrOutputWithContext(context.Context) AzureActiveDirectoryRegistrationPtrOutput
}

type azureActiveDirectoryRegistrationPtrType AzureActiveDirectoryRegistrationArgs

func AzureActiveDirectoryRegistrationPtr(v *AzureActiveDirectoryRegistrationArgs) AzureActiveDirectoryRegistrationPtrInput {
	return (*azureActiveDirectoryRegistrationPtrType)(v)
}

func (*azureActiveDirectoryRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryRegistration)(nil)).Elem()
}

func (i *azureActiveDirectoryRegistrationPtrType) ToAzureActiveDirectoryRegistrationPtrOutput() AzureActiveDirectoryRegistrationPtrOutput {
	return i.ToAzureActiveDirectoryRegistrationPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryRegistrationPtrType) ToAzureActiveDirectoryRegistrationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryRegistrationPtrOutput)
}

type AzureActiveDirectoryRegistrationOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryRegistration)(nil)).Elem()
}

func (o AzureActiveDirectoryRegistrationOutput) ToAzureActiveDirectoryRegistrationOutput() AzureActiveDirectoryRegistrationOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationOutput) ToAzureActiveDirectoryRegistrationOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationOutput) ToAzureActiveDirectoryRegistrationPtrOutput() AzureActiveDirectoryRegistrationPtrOutput {
	return o.ToAzureActiveDirectoryRegistrationPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryRegistrationOutput) ToAzureActiveDirectoryRegistrationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectoryRegistration) *AzureActiveDirectoryRegistration {
		return &v
	}).(AzureActiveDirectoryRegistrationPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) ClientSecretCertificateIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string { return v.ClientSecretCertificateIssuer }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) ClientSecretCertificateSubjectAlternativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string {
		return v.ClientSecretCertificateSubjectAlternativeName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string { return v.ClientSecretCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string { return v.OpenIdIssuer }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryRegistrationPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryRegistration)(nil)).Elem()
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ToAzureActiveDirectoryRegistrationPtrOutput() AzureActiveDirectoryRegistrationPtrOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ToAzureActiveDirectoryRegistrationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationPtrOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationPtrOutput) Elem() AzureActiveDirectoryRegistrationOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) AzureActiveDirectoryRegistration {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryRegistration
		return ret
	}).(AzureActiveDirectoryRegistrationOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ClientSecretCertificateIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateIssuer
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ClientSecretCertificateSubjectAlternativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateSubjectAlternativeName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.OpenIdIssuer
	}).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryRegistrationResponse struct {
	ClientId                                      *string `pulumi:"clientId"`
	ClientSecretCertificateIssuer                 *string `pulumi:"clientSecretCertificateIssuer"`
	ClientSecretCertificateSubjectAlternativeName *string `pulumi:"clientSecretCertificateSubjectAlternativeName"`
	ClientSecretCertificateThumbprint             *string `pulumi:"clientSecretCertificateThumbprint"`
	ClientSecretSettingName                       *string `pulumi:"clientSecretSettingName"`
	OpenIdIssuer                                  *string `pulumi:"openIdIssuer"`
}

type AzureActiveDirectoryRegistrationResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryRegistrationResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ToAzureActiveDirectoryRegistrationResponseOutput() AzureActiveDirectoryRegistrationResponseOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ToAzureActiveDirectoryRegistrationResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationResponseOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ClientSecretCertificateIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string { return v.ClientSecretCertificateIssuer }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ClientSecretCertificateSubjectAlternativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string {
		return v.ClientSecretCertificateSubjectAlternativeName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string { return v.ClientSecretCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponseOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string { return v.OpenIdIssuer }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryRegistrationResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ToAzureActiveDirectoryRegistrationResponsePtrOutput() AzureActiveDirectoryRegistrationResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ToAzureActiveDirectoryRegistrationResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) Elem() AzureActiveDirectoryRegistrationResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) AzureActiveDirectoryRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryRegistrationResponse
		return ret
	}).(AzureActiveDirectoryRegistrationResponseOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ClientSecretCertificateIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateIssuer
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ClientSecretCertificateSubjectAlternativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateSubjectAlternativeName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.OpenIdIssuer
	}).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryResponse struct {
	Enabled           *bool                                     `pulumi:"enabled"`
	IsAutoProvisioned *bool                                     `pulumi:"isAutoProvisioned"`
	Login             *AzureActiveDirectoryLoginResponse        `pulumi:"login"`
	Registration      *AzureActiveDirectoryRegistrationResponse `pulumi:"registration"`
	Validation        *AzureActiveDirectoryValidationResponse   `pulumi:"validation"`
}

type AzureActiveDirectoryResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponseOutput() AzureActiveDirectoryResponseOutput {
	return o
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryResponseOutput {
	return o
}

func (o AzureActiveDirectoryResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) IsAutoProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *bool { return v.IsAutoProvisioned }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) Login() AzureActiveDirectoryLoginResponsePtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *AzureActiveDirectoryLoginResponse { return v.Login }).(AzureActiveDirectoryLoginResponsePtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) Registration() AzureActiveDirectoryRegistrationResponsePtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *AzureActiveDirectoryRegistrationResponse { return v.Registration }).(AzureActiveDirectoryRegistrationResponsePtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) Validation() AzureActiveDirectoryValidationResponsePtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *AzureActiveDirectoryValidationResponse { return v.Validation }).(AzureActiveDirectoryValidationResponsePtrOutput)
}

type AzureActiveDirectoryResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryResponsePtrOutput) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryResponsePtrOutput) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryResponsePtrOutput) Elem() AzureActiveDirectoryResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) AzureActiveDirectoryResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryResponse
		return ret
	}).(AzureActiveDirectoryResponseOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) IsAutoProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsAutoProvisioned
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) Login() AzureActiveDirectoryLoginResponsePtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *AzureActiveDirectoryLoginResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(AzureActiveDirectoryLoginResponsePtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) Registration() AzureActiveDirectoryRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *AzureActiveDirectoryRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AzureActiveDirectoryRegistrationResponsePtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) Validation() AzureActiveDirectoryValidationResponsePtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *AzureActiveDirectoryValidationResponse {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(AzureActiveDirectoryValidationResponsePtrOutput)
}

type AzureActiveDirectoryValidation struct {
	AllowedAudiences           []string                    `pulumi:"allowedAudiences"`
	DefaultAuthorizationPolicy *DefaultAuthorizationPolicy `pulumi:"defaultAuthorizationPolicy"`
	JwtClaimChecks             *JwtClaimChecks             `pulumi:"jwtClaimChecks"`
}





type AzureActiveDirectoryValidationInput interface {
	pulumi.Input

	ToAzureActiveDirectoryValidationOutput() AzureActiveDirectoryValidationOutput
	ToAzureActiveDirectoryValidationOutputWithContext(context.Context) AzureActiveDirectoryValidationOutput
}

type AzureActiveDirectoryValidationArgs struct {
	AllowedAudiences           pulumi.StringArrayInput            `pulumi:"allowedAudiences"`
	DefaultAuthorizationPolicy DefaultAuthorizationPolicyPtrInput `pulumi:"defaultAuthorizationPolicy"`
	JwtClaimChecks             JwtClaimChecksPtrInput             `pulumi:"jwtClaimChecks"`
}

func (AzureActiveDirectoryValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryValidation)(nil)).Elem()
}

func (i AzureActiveDirectoryValidationArgs) ToAzureActiveDirectoryValidationOutput() AzureActiveDirectoryValidationOutput {
	return i.ToAzureActiveDirectoryValidationOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryValidationArgs) ToAzureActiveDirectoryValidationOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryValidationOutput)
}

func (i AzureActiveDirectoryValidationArgs) ToAzureActiveDirectoryValidationPtrOutput() AzureActiveDirectoryValidationPtrOutput {
	return i.ToAzureActiveDirectoryValidationPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryValidationArgs) ToAzureActiveDirectoryValidationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryValidationOutput).ToAzureActiveDirectoryValidationPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryValidationPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryValidationPtrOutput() AzureActiveDirectoryValidationPtrOutput
	ToAzureActiveDirectoryValidationPtrOutputWithContext(context.Context) AzureActiveDirectoryValidationPtrOutput
}

type azureActiveDirectoryValidationPtrType AzureActiveDirectoryValidationArgs

func AzureActiveDirectoryValidationPtr(v *AzureActiveDirectoryValidationArgs) AzureActiveDirectoryValidationPtrInput {
	return (*azureActiveDirectoryValidationPtrType)(v)
}

func (*azureActiveDirectoryValidationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryValidation)(nil)).Elem()
}

func (i *azureActiveDirectoryValidationPtrType) ToAzureActiveDirectoryValidationPtrOutput() AzureActiveDirectoryValidationPtrOutput {
	return i.ToAzureActiveDirectoryValidationPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryValidationPtrType) ToAzureActiveDirectoryValidationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryValidationPtrOutput)
}

type AzureActiveDirectoryValidationOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryValidation)(nil)).Elem()
}

func (o AzureActiveDirectoryValidationOutput) ToAzureActiveDirectoryValidationOutput() AzureActiveDirectoryValidationOutput {
	return o
}

func (o AzureActiveDirectoryValidationOutput) ToAzureActiveDirectoryValidationOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationOutput {
	return o
}

func (o AzureActiveDirectoryValidationOutput) ToAzureActiveDirectoryValidationPtrOutput() AzureActiveDirectoryValidationPtrOutput {
	return o.ToAzureActiveDirectoryValidationPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryValidationOutput) ToAzureActiveDirectoryValidationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectoryValidation) *AzureActiveDirectoryValidation {
		return &v
	}).(AzureActiveDirectoryValidationPtrOutput)
}

func (o AzureActiveDirectoryValidationOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidation) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o AzureActiveDirectoryValidationOutput) DefaultAuthorizationPolicy() DefaultAuthorizationPolicyPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidation) *DefaultAuthorizationPolicy {
		return v.DefaultAuthorizationPolicy
	}).(DefaultAuthorizationPolicyPtrOutput)
}

func (o AzureActiveDirectoryValidationOutput) JwtClaimChecks() JwtClaimChecksPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidation) *JwtClaimChecks { return v.JwtClaimChecks }).(JwtClaimChecksPtrOutput)
}

type AzureActiveDirectoryValidationPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryValidation)(nil)).Elem()
}

func (o AzureActiveDirectoryValidationPtrOutput) ToAzureActiveDirectoryValidationPtrOutput() AzureActiveDirectoryValidationPtrOutput {
	return o
}

func (o AzureActiveDirectoryValidationPtrOutput) ToAzureActiveDirectoryValidationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationPtrOutput {
	return o
}

func (o AzureActiveDirectoryValidationPtrOutput) Elem() AzureActiveDirectoryValidationOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidation) AzureActiveDirectoryValidation {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryValidation
		return ret
	}).(AzureActiveDirectoryValidationOutput)
}

func (o AzureActiveDirectoryValidationPtrOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidation) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAudiences
	}).(pulumi.StringArrayOutput)
}

func (o AzureActiveDirectoryValidationPtrOutput) DefaultAuthorizationPolicy() DefaultAuthorizationPolicyPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidation) *DefaultAuthorizationPolicy {
		if v == nil {
			return nil
		}
		return v.DefaultAuthorizationPolicy
	}).(DefaultAuthorizationPolicyPtrOutput)
}

func (o AzureActiveDirectoryValidationPtrOutput) JwtClaimChecks() JwtClaimChecksPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidation) *JwtClaimChecks {
		if v == nil {
			return nil
		}
		return v.JwtClaimChecks
	}).(JwtClaimChecksPtrOutput)
}

type AzureActiveDirectoryValidationResponse struct {
	AllowedAudiences           []string                            `pulumi:"allowedAudiences"`
	DefaultAuthorizationPolicy *DefaultAuthorizationPolicyResponse `pulumi:"defaultAuthorizationPolicy"`
	JwtClaimChecks             *JwtClaimChecksResponse             `pulumi:"jwtClaimChecks"`
}

type AzureActiveDirectoryValidationResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryValidationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryValidationResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryValidationResponseOutput) ToAzureActiveDirectoryValidationResponseOutput() AzureActiveDirectoryValidationResponseOutput {
	return o
}

func (o AzureActiveDirectoryValidationResponseOutput) ToAzureActiveDirectoryValidationResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationResponseOutput {
	return o
}

func (o AzureActiveDirectoryValidationResponseOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidationResponse) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o AzureActiveDirectoryValidationResponseOutput) DefaultAuthorizationPolicy() DefaultAuthorizationPolicyResponsePtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidationResponse) *DefaultAuthorizationPolicyResponse {
		return v.DefaultAuthorizationPolicy
	}).(DefaultAuthorizationPolicyResponsePtrOutput)
}

func (o AzureActiveDirectoryValidationResponseOutput) JwtClaimChecks() JwtClaimChecksResponsePtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidationResponse) *JwtClaimChecksResponse { return v.JwtClaimChecks }).(JwtClaimChecksResponsePtrOutput)
}

type AzureActiveDirectoryValidationResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryValidationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryValidationResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) ToAzureActiveDirectoryValidationResponsePtrOutput() AzureActiveDirectoryValidationResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) ToAzureActiveDirectoryValidationResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) Elem() AzureActiveDirectoryValidationResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidationResponse) AzureActiveDirectoryValidationResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryValidationResponse
		return ret
	}).(AzureActiveDirectoryValidationResponseOutput)
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidationResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAudiences
	}).(pulumi.StringArrayOutput)
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) DefaultAuthorizationPolicy() DefaultAuthorizationPolicyResponsePtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidationResponse) *DefaultAuthorizationPolicyResponse {
		if v == nil {
			return nil
		}
		return v.DefaultAuthorizationPolicy
	}).(DefaultAuthorizationPolicyResponsePtrOutput)
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) JwtClaimChecks() JwtClaimChecksResponsePtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidationResponse) *JwtClaimChecksResponse {
		if v == nil {
			return nil
		}
		return v.JwtClaimChecks
	}).(JwtClaimChecksResponsePtrOutput)
}

type AzureCredentials struct {
	ClientId       *string `pulumi:"clientId"`
	ClientSecret   *string `pulumi:"clientSecret"`
	SubscriptionId *string `pulumi:"subscriptionId"`
	TenantId       *string `pulumi:"tenantId"`
}





type AzureCredentialsInput interface {
	pulumi.Input

	ToAzureCredentialsOutput() AzureCredentialsOutput
	ToAzureCredentialsOutputWithContext(context.Context) AzureCredentialsOutput
}

type AzureCredentialsArgs struct {
	ClientId       pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret   pulumi.StringPtrInput `pulumi:"clientSecret"`
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
	TenantId       pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AzureCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureCredentials)(nil)).Elem()
}

func (i AzureCredentialsArgs) ToAzureCredentialsOutput() AzureCredentialsOutput {
	return i.ToAzureCredentialsOutputWithContext(context.Background())
}

func (i AzureCredentialsArgs) ToAzureCredentialsOutputWithContext(ctx context.Context) AzureCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCredentialsOutput)
}

func (i AzureCredentialsArgs) ToAzureCredentialsPtrOutput() AzureCredentialsPtrOutput {
	return i.ToAzureCredentialsPtrOutputWithContext(context.Background())
}

func (i AzureCredentialsArgs) ToAzureCredentialsPtrOutputWithContext(ctx context.Context) AzureCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCredentialsOutput).ToAzureCredentialsPtrOutputWithContext(ctx)
}









type AzureCredentialsPtrInput interface {
	pulumi.Input

	ToAzureCredentialsPtrOutput() AzureCredentialsPtrOutput
	ToAzureCredentialsPtrOutputWithContext(context.Context) AzureCredentialsPtrOutput
}

type azureCredentialsPtrType AzureCredentialsArgs

func AzureCredentialsPtr(v *AzureCredentialsArgs) AzureCredentialsPtrInput {
	return (*azureCredentialsPtrType)(v)
}

func (*azureCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCredentials)(nil)).Elem()
}

func (i *azureCredentialsPtrType) ToAzureCredentialsPtrOutput() AzureCredentialsPtrOutput {
	return i.ToAzureCredentialsPtrOutputWithContext(context.Background())
}

func (i *azureCredentialsPtrType) ToAzureCredentialsPtrOutputWithContext(ctx context.Context) AzureCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCredentialsPtrOutput)
}

type AzureCredentialsOutput struct{ *pulumi.OutputState }

func (AzureCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureCredentials)(nil)).Elem()
}

func (o AzureCredentialsOutput) ToAzureCredentialsOutput() AzureCredentialsOutput {
	return o
}

func (o AzureCredentialsOutput) ToAzureCredentialsOutputWithContext(ctx context.Context) AzureCredentialsOutput {
	return o
}

func (o AzureCredentialsOutput) ToAzureCredentialsPtrOutput() AzureCredentialsPtrOutput {
	return o.ToAzureCredentialsPtrOutputWithContext(context.Background())
}

func (o AzureCredentialsOutput) ToAzureCredentialsPtrOutputWithContext(ctx context.Context) AzureCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureCredentials) *AzureCredentials {
		return &v
	}).(AzureCredentialsPtrOutput)
}

func (o AzureCredentialsOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureCredentials) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AzureCredentialsOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureCredentials) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o AzureCredentialsOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureCredentials) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o AzureCredentialsOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureCredentials) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AzureCredentialsPtrOutput struct{ *pulumi.OutputState }

func (AzureCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCredentials)(nil)).Elem()
}

func (o AzureCredentialsPtrOutput) ToAzureCredentialsPtrOutput() AzureCredentialsPtrOutput {
	return o
}

func (o AzureCredentialsPtrOutput) ToAzureCredentialsPtrOutputWithContext(ctx context.Context) AzureCredentialsPtrOutput {
	return o
}

func (o AzureCredentialsPtrOutput) Elem() AzureCredentialsOutput {
	return o.ApplyT(func(v *AzureCredentials) AzureCredentials {
		if v != nil {
			return *v
		}
		var ret AzureCredentials
		return ret
	}).(AzureCredentialsOutput)
}

func (o AzureCredentialsPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCredentials) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AzureCredentialsPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCredentials) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o AzureCredentialsPtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCredentials) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o AzureCredentialsPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCredentials) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type AzureCredentialsResponse struct {
	SubscriptionId *string `pulumi:"subscriptionId"`
}

type AzureCredentialsResponseOutput struct{ *pulumi.OutputState }

func (AzureCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureCredentialsResponse)(nil)).Elem()
}

func (o AzureCredentialsResponseOutput) ToAzureCredentialsResponseOutput() AzureCredentialsResponseOutput {
	return o
}

func (o AzureCredentialsResponseOutput) ToAzureCredentialsResponseOutputWithContext(ctx context.Context) AzureCredentialsResponseOutput {
	return o
}

func (o AzureCredentialsResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureCredentialsResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type AzureCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCredentialsResponse)(nil)).Elem()
}

func (o AzureCredentialsResponsePtrOutput) ToAzureCredentialsResponsePtrOutput() AzureCredentialsResponsePtrOutput {
	return o
}

func (o AzureCredentialsResponsePtrOutput) ToAzureCredentialsResponsePtrOutputWithContext(ctx context.Context) AzureCredentialsResponsePtrOutput {
	return o
}

func (o AzureCredentialsResponsePtrOutput) Elem() AzureCredentialsResponseOutput {
	return o.ApplyT(func(v *AzureCredentialsResponse) AzureCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret AzureCredentialsResponse
		return ret
	}).(AzureCredentialsResponseOutput)
}

func (o AzureCredentialsResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

type AzureFileProperties struct {
	AccessMode  *string `pulumi:"accessMode"`
	AccountKey  *string `pulumi:"accountKey"`
	AccountName *string `pulumi:"accountName"`
	ShareName   *string `pulumi:"shareName"`
}





type AzureFilePropertiesInput interface {
	pulumi.Input

	ToAzureFilePropertiesOutput() AzureFilePropertiesOutput
	ToAzureFilePropertiesOutputWithContext(context.Context) AzureFilePropertiesOutput
}

type AzureFilePropertiesArgs struct {
	AccessMode  pulumi.StringPtrInput `pulumi:"accessMode"`
	AccountKey  pulumi.StringPtrInput `pulumi:"accountKey"`
	AccountName pulumi.StringPtrInput `pulumi:"accountName"`
	ShareName   pulumi.StringPtrInput `pulumi:"shareName"`
}

func (AzureFilePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileProperties)(nil)).Elem()
}

func (i AzureFilePropertiesArgs) ToAzureFilePropertiesOutput() AzureFilePropertiesOutput {
	return i.ToAzureFilePropertiesOutputWithContext(context.Background())
}

func (i AzureFilePropertiesArgs) ToAzureFilePropertiesOutputWithContext(ctx context.Context) AzureFilePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFilePropertiesOutput)
}

func (i AzureFilePropertiesArgs) ToAzureFilePropertiesPtrOutput() AzureFilePropertiesPtrOutput {
	return i.ToAzureFilePropertiesPtrOutputWithContext(context.Background())
}

func (i AzureFilePropertiesArgs) ToAzureFilePropertiesPtrOutputWithContext(ctx context.Context) AzureFilePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFilePropertiesOutput).ToAzureFilePropertiesPtrOutputWithContext(ctx)
}









type AzureFilePropertiesPtrInput interface {
	pulumi.Input

	ToAzureFilePropertiesPtrOutput() AzureFilePropertiesPtrOutput
	ToAzureFilePropertiesPtrOutputWithContext(context.Context) AzureFilePropertiesPtrOutput
}

type azureFilePropertiesPtrType AzureFilePropertiesArgs

func AzureFilePropertiesPtr(v *AzureFilePropertiesArgs) AzureFilePropertiesPtrInput {
	return (*azureFilePropertiesPtrType)(v)
}

func (*azureFilePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileProperties)(nil)).Elem()
}

func (i *azureFilePropertiesPtrType) ToAzureFilePropertiesPtrOutput() AzureFilePropertiesPtrOutput {
	return i.ToAzureFilePropertiesPtrOutputWithContext(context.Background())
}

func (i *azureFilePropertiesPtrType) ToAzureFilePropertiesPtrOutputWithContext(ctx context.Context) AzureFilePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFilePropertiesPtrOutput)
}

type AzureFilePropertiesOutput struct{ *pulumi.OutputState }

func (AzureFilePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileProperties)(nil)).Elem()
}

func (o AzureFilePropertiesOutput) ToAzureFilePropertiesOutput() AzureFilePropertiesOutput {
	return o
}

func (o AzureFilePropertiesOutput) ToAzureFilePropertiesOutputWithContext(ctx context.Context) AzureFilePropertiesOutput {
	return o
}

func (o AzureFilePropertiesOutput) ToAzureFilePropertiesPtrOutput() AzureFilePropertiesPtrOutput {
	return o.ToAzureFilePropertiesPtrOutputWithContext(context.Background())
}

func (o AzureFilePropertiesOutput) ToAzureFilePropertiesPtrOutputWithContext(ctx context.Context) AzureFilePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFileProperties) *AzureFileProperties {
		return &v
	}).(AzureFilePropertiesPtrOutput)
}

func (o AzureFilePropertiesOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileProperties) *string { return v.AccessMode }).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileProperties) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileProperties) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileProperties) *string { return v.ShareName }).(pulumi.StringPtrOutput)
}

type AzureFilePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AzureFilePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileProperties)(nil)).Elem()
}

func (o AzureFilePropertiesPtrOutput) ToAzureFilePropertiesPtrOutput() AzureFilePropertiesPtrOutput {
	return o
}

func (o AzureFilePropertiesPtrOutput) ToAzureFilePropertiesPtrOutputWithContext(ctx context.Context) AzureFilePropertiesPtrOutput {
	return o
}

func (o AzureFilePropertiesPtrOutput) Elem() AzureFilePropertiesOutput {
	return o.ApplyT(func(v *AzureFileProperties) AzureFileProperties {
		if v != nil {
			return *v
		}
		var ret AzureFileProperties
		return ret
	}).(AzureFilePropertiesOutput)
}

func (o AzureFilePropertiesPtrOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileProperties) *string {
		if v == nil {
			return nil
		}
		return v.AccessMode
	}).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesPtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileProperties) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesPtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileProperties) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesPtrOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileProperties) *string {
		if v == nil {
			return nil
		}
		return v.ShareName
	}).(pulumi.StringPtrOutput)
}

type AzureFilePropertiesResponse struct {
	AccessMode  *string `pulumi:"accessMode"`
	AccountKey  *string `pulumi:"accountKey"`
	AccountName *string `pulumi:"accountName"`
	ShareName   *string `pulumi:"shareName"`
}

type AzureFilePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AzureFilePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFilePropertiesResponse)(nil)).Elem()
}

func (o AzureFilePropertiesResponseOutput) ToAzureFilePropertiesResponseOutput() AzureFilePropertiesResponseOutput {
	return o
}

func (o AzureFilePropertiesResponseOutput) ToAzureFilePropertiesResponseOutputWithContext(ctx context.Context) AzureFilePropertiesResponseOutput {
	return o
}

func (o AzureFilePropertiesResponseOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFilePropertiesResponse) *string { return v.AccessMode }).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFilePropertiesResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFilePropertiesResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesResponseOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFilePropertiesResponse) *string { return v.ShareName }).(pulumi.StringPtrOutput)
}

type AzureFilePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureFilePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFilePropertiesResponse)(nil)).Elem()
}

func (o AzureFilePropertiesResponsePtrOutput) ToAzureFilePropertiesResponsePtrOutput() AzureFilePropertiesResponsePtrOutput {
	return o
}

func (o AzureFilePropertiesResponsePtrOutput) ToAzureFilePropertiesResponsePtrOutputWithContext(ctx context.Context) AzureFilePropertiesResponsePtrOutput {
	return o
}

func (o AzureFilePropertiesResponsePtrOutput) Elem() AzureFilePropertiesResponseOutput {
	return o.ApplyT(func(v *AzureFilePropertiesResponse) AzureFilePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AzureFilePropertiesResponse
		return ret
	}).(AzureFilePropertiesResponseOutput)
}

func (o AzureFilePropertiesResponsePtrOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFilePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccessMode
	}).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesResponsePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFilePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesResponsePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFilePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o AzureFilePropertiesResponsePtrOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFilePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ShareName
	}).(pulumi.StringPtrOutput)
}

type AzureStaticWebApps struct {
	Enabled      *bool                           `pulumi:"enabled"`
	Registration *AzureStaticWebAppsRegistration `pulumi:"registration"`
}





type AzureStaticWebAppsInput interface {
	pulumi.Input

	ToAzureStaticWebAppsOutput() AzureStaticWebAppsOutput
	ToAzureStaticWebAppsOutputWithContext(context.Context) AzureStaticWebAppsOutput
}

type AzureStaticWebAppsArgs struct {
	Enabled      pulumi.BoolPtrInput                    `pulumi:"enabled"`
	Registration AzureStaticWebAppsRegistrationPtrInput `pulumi:"registration"`
}

func (AzureStaticWebAppsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebApps)(nil)).Elem()
}

func (i AzureStaticWebAppsArgs) ToAzureStaticWebAppsOutput() AzureStaticWebAppsOutput {
	return i.ToAzureStaticWebAppsOutputWithContext(context.Background())
}

func (i AzureStaticWebAppsArgs) ToAzureStaticWebAppsOutputWithContext(ctx context.Context) AzureStaticWebAppsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsOutput)
}

func (i AzureStaticWebAppsArgs) ToAzureStaticWebAppsPtrOutput() AzureStaticWebAppsPtrOutput {
	return i.ToAzureStaticWebAppsPtrOutputWithContext(context.Background())
}

func (i AzureStaticWebAppsArgs) ToAzureStaticWebAppsPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsOutput).ToAzureStaticWebAppsPtrOutputWithContext(ctx)
}









type AzureStaticWebAppsPtrInput interface {
	pulumi.Input

	ToAzureStaticWebAppsPtrOutput() AzureStaticWebAppsPtrOutput
	ToAzureStaticWebAppsPtrOutputWithContext(context.Context) AzureStaticWebAppsPtrOutput
}

type azureStaticWebAppsPtrType AzureStaticWebAppsArgs

func AzureStaticWebAppsPtr(v *AzureStaticWebAppsArgs) AzureStaticWebAppsPtrInput {
	return (*azureStaticWebAppsPtrType)(v)
}

func (*azureStaticWebAppsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebApps)(nil)).Elem()
}

func (i *azureStaticWebAppsPtrType) ToAzureStaticWebAppsPtrOutput() AzureStaticWebAppsPtrOutput {
	return i.ToAzureStaticWebAppsPtrOutputWithContext(context.Background())
}

func (i *azureStaticWebAppsPtrType) ToAzureStaticWebAppsPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsPtrOutput)
}

type AzureStaticWebAppsOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebApps)(nil)).Elem()
}

func (o AzureStaticWebAppsOutput) ToAzureStaticWebAppsOutput() AzureStaticWebAppsOutput {
	return o
}

func (o AzureStaticWebAppsOutput) ToAzureStaticWebAppsOutputWithContext(ctx context.Context) AzureStaticWebAppsOutput {
	return o
}

func (o AzureStaticWebAppsOutput) ToAzureStaticWebAppsPtrOutput() AzureStaticWebAppsPtrOutput {
	return o.ToAzureStaticWebAppsPtrOutputWithContext(context.Background())
}

func (o AzureStaticWebAppsOutput) ToAzureStaticWebAppsPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureStaticWebApps) *AzureStaticWebApps {
		return &v
	}).(AzureStaticWebAppsPtrOutput)
}

func (o AzureStaticWebAppsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureStaticWebApps) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureStaticWebAppsOutput) Registration() AzureStaticWebAppsRegistrationPtrOutput {
	return o.ApplyT(func(v AzureStaticWebApps) *AzureStaticWebAppsRegistration { return v.Registration }).(AzureStaticWebAppsRegistrationPtrOutput)
}

type AzureStaticWebAppsPtrOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebApps)(nil)).Elem()
}

func (o AzureStaticWebAppsPtrOutput) ToAzureStaticWebAppsPtrOutput() AzureStaticWebAppsPtrOutput {
	return o
}

func (o AzureStaticWebAppsPtrOutput) ToAzureStaticWebAppsPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsPtrOutput {
	return o
}

func (o AzureStaticWebAppsPtrOutput) Elem() AzureStaticWebAppsOutput {
	return o.ApplyT(func(v *AzureStaticWebApps) AzureStaticWebApps {
		if v != nil {
			return *v
		}
		var ret AzureStaticWebApps
		return ret
	}).(AzureStaticWebAppsOutput)
}

func (o AzureStaticWebAppsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureStaticWebApps) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureStaticWebAppsPtrOutput) Registration() AzureStaticWebAppsRegistrationPtrOutput {
	return o.ApplyT(func(v *AzureStaticWebApps) *AzureStaticWebAppsRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AzureStaticWebAppsRegistrationPtrOutput)
}

type AzureStaticWebAppsRegistration struct {
	ClientId *string `pulumi:"clientId"`
}





type AzureStaticWebAppsRegistrationInput interface {
	pulumi.Input

	ToAzureStaticWebAppsRegistrationOutput() AzureStaticWebAppsRegistrationOutput
	ToAzureStaticWebAppsRegistrationOutputWithContext(context.Context) AzureStaticWebAppsRegistrationOutput
}

type AzureStaticWebAppsRegistrationArgs struct {
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
}

func (AzureStaticWebAppsRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebAppsRegistration)(nil)).Elem()
}

func (i AzureStaticWebAppsRegistrationArgs) ToAzureStaticWebAppsRegistrationOutput() AzureStaticWebAppsRegistrationOutput {
	return i.ToAzureStaticWebAppsRegistrationOutputWithContext(context.Background())
}

func (i AzureStaticWebAppsRegistrationArgs) ToAzureStaticWebAppsRegistrationOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsRegistrationOutput)
}

func (i AzureStaticWebAppsRegistrationArgs) ToAzureStaticWebAppsRegistrationPtrOutput() AzureStaticWebAppsRegistrationPtrOutput {
	return i.ToAzureStaticWebAppsRegistrationPtrOutputWithContext(context.Background())
}

func (i AzureStaticWebAppsRegistrationArgs) ToAzureStaticWebAppsRegistrationPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsRegistrationOutput).ToAzureStaticWebAppsRegistrationPtrOutputWithContext(ctx)
}









type AzureStaticWebAppsRegistrationPtrInput interface {
	pulumi.Input

	ToAzureStaticWebAppsRegistrationPtrOutput() AzureStaticWebAppsRegistrationPtrOutput
	ToAzureStaticWebAppsRegistrationPtrOutputWithContext(context.Context) AzureStaticWebAppsRegistrationPtrOutput
}

type azureStaticWebAppsRegistrationPtrType AzureStaticWebAppsRegistrationArgs

func AzureStaticWebAppsRegistrationPtr(v *AzureStaticWebAppsRegistrationArgs) AzureStaticWebAppsRegistrationPtrInput {
	return (*azureStaticWebAppsRegistrationPtrType)(v)
}

func (*azureStaticWebAppsRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebAppsRegistration)(nil)).Elem()
}

func (i *azureStaticWebAppsRegistrationPtrType) ToAzureStaticWebAppsRegistrationPtrOutput() AzureStaticWebAppsRegistrationPtrOutput {
	return i.ToAzureStaticWebAppsRegistrationPtrOutputWithContext(context.Background())
}

func (i *azureStaticWebAppsRegistrationPtrType) ToAzureStaticWebAppsRegistrationPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsRegistrationPtrOutput)
}

type AzureStaticWebAppsRegistrationOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebAppsRegistration)(nil)).Elem()
}

func (o AzureStaticWebAppsRegistrationOutput) ToAzureStaticWebAppsRegistrationOutput() AzureStaticWebAppsRegistrationOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationOutput) ToAzureStaticWebAppsRegistrationOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationOutput) ToAzureStaticWebAppsRegistrationPtrOutput() AzureStaticWebAppsRegistrationPtrOutput {
	return o.ToAzureStaticWebAppsRegistrationPtrOutputWithContext(context.Background())
}

func (o AzureStaticWebAppsRegistrationOutput) ToAzureStaticWebAppsRegistrationPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureStaticWebAppsRegistration) *AzureStaticWebAppsRegistration {
		return &v
	}).(AzureStaticWebAppsRegistrationPtrOutput)
}

func (o AzureStaticWebAppsRegistrationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStaticWebAppsRegistration) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

type AzureStaticWebAppsRegistrationPtrOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebAppsRegistration)(nil)).Elem()
}

func (o AzureStaticWebAppsRegistrationPtrOutput) ToAzureStaticWebAppsRegistrationPtrOutput() AzureStaticWebAppsRegistrationPtrOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationPtrOutput) ToAzureStaticWebAppsRegistrationPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationPtrOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationPtrOutput) Elem() AzureStaticWebAppsRegistrationOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsRegistration) AzureStaticWebAppsRegistration {
		if v != nil {
			return *v
		}
		var ret AzureStaticWebAppsRegistration
		return ret
	}).(AzureStaticWebAppsRegistrationOutput)
}

func (o AzureStaticWebAppsRegistrationPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

type AzureStaticWebAppsRegistrationResponse struct {
	ClientId *string `pulumi:"clientId"`
}

type AzureStaticWebAppsRegistrationResponseOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebAppsRegistrationResponse)(nil)).Elem()
}

func (o AzureStaticWebAppsRegistrationResponseOutput) ToAzureStaticWebAppsRegistrationResponseOutput() AzureStaticWebAppsRegistrationResponseOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationResponseOutput) ToAzureStaticWebAppsRegistrationResponseOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationResponseOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStaticWebAppsRegistrationResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

type AzureStaticWebAppsRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebAppsRegistrationResponse)(nil)).Elem()
}

func (o AzureStaticWebAppsRegistrationResponsePtrOutput) ToAzureStaticWebAppsRegistrationResponsePtrOutput() AzureStaticWebAppsRegistrationResponsePtrOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationResponsePtrOutput) ToAzureStaticWebAppsRegistrationResponsePtrOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationResponsePtrOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationResponsePtrOutput) Elem() AzureStaticWebAppsRegistrationResponseOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsRegistrationResponse) AzureStaticWebAppsRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret AzureStaticWebAppsRegistrationResponse
		return ret
	}).(AzureStaticWebAppsRegistrationResponseOutput)
}

func (o AzureStaticWebAppsRegistrationResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

type AzureStaticWebAppsResponse struct {
	Enabled      *bool                                   `pulumi:"enabled"`
	Registration *AzureStaticWebAppsRegistrationResponse `pulumi:"registration"`
}

type AzureStaticWebAppsResponseOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebAppsResponse)(nil)).Elem()
}

func (o AzureStaticWebAppsResponseOutput) ToAzureStaticWebAppsResponseOutput() AzureStaticWebAppsResponseOutput {
	return o
}

func (o AzureStaticWebAppsResponseOutput) ToAzureStaticWebAppsResponseOutputWithContext(ctx context.Context) AzureStaticWebAppsResponseOutput {
	return o
}

func (o AzureStaticWebAppsResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureStaticWebAppsResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureStaticWebAppsResponseOutput) Registration() AzureStaticWebAppsRegistrationResponsePtrOutput {
	return o.ApplyT(func(v AzureStaticWebAppsResponse) *AzureStaticWebAppsRegistrationResponse { return v.Registration }).(AzureStaticWebAppsRegistrationResponsePtrOutput)
}

type AzureStaticWebAppsResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebAppsResponse)(nil)).Elem()
}

func (o AzureStaticWebAppsResponsePtrOutput) ToAzureStaticWebAppsResponsePtrOutput() AzureStaticWebAppsResponsePtrOutput {
	return o
}

func (o AzureStaticWebAppsResponsePtrOutput) ToAzureStaticWebAppsResponsePtrOutputWithContext(ctx context.Context) AzureStaticWebAppsResponsePtrOutput {
	return o
}

func (o AzureStaticWebAppsResponsePtrOutput) Elem() AzureStaticWebAppsResponseOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsResponse) AzureStaticWebAppsResponse {
		if v != nil {
			return *v
		}
		var ret AzureStaticWebAppsResponse
		return ret
	}).(AzureStaticWebAppsResponseOutput)
}

func (o AzureStaticWebAppsResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureStaticWebAppsResponsePtrOutput) Registration() AzureStaticWebAppsRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsResponse) *AzureStaticWebAppsRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AzureStaticWebAppsRegistrationResponsePtrOutput)
}

type CertificateProperties struct {
	Password *string `pulumi:"password"`
	Value    *string `pulumi:"value"`
}





type CertificatePropertiesInput interface {
	pulumi.Input

	ToCertificatePropertiesOutput() CertificatePropertiesOutput
	ToCertificatePropertiesOutputWithContext(context.Context) CertificatePropertiesOutput
}

type CertificatePropertiesArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Value    pulumi.StringPtrInput `pulumi:"value"`
}

func (CertificatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateProperties)(nil)).Elem()
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesOutput() CertificatePropertiesOutput {
	return i.ToCertificatePropertiesOutputWithContext(context.Background())
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesOutputWithContext(ctx context.Context) CertificatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesOutput)
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return i.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesOutput).ToCertificatePropertiesPtrOutputWithContext(ctx)
}









type CertificatePropertiesPtrInput interface {
	pulumi.Input

	ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput
	ToCertificatePropertiesPtrOutputWithContext(context.Context) CertificatePropertiesPtrOutput
}

type certificatePropertiesPtrType CertificatePropertiesArgs

func CertificatePropertiesPtr(v *CertificatePropertiesArgs) CertificatePropertiesPtrInput {
	return (*certificatePropertiesPtrType)(v)
}

func (*certificatePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProperties)(nil)).Elem()
}

func (i *certificatePropertiesPtrType) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return i.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (i *certificatePropertiesPtrType) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesPtrOutput)
}

type CertificatePropertiesOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateProperties)(nil)).Elem()
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesOutput() CertificatePropertiesOutput {
	return o
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesOutputWithContext(ctx context.Context) CertificatePropertiesOutput {
	return o
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return o.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateProperties) *CertificateProperties {
		return &v
	}).(CertificatePropertiesPtrOutput)
}

func (o CertificatePropertiesOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateProperties) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateProperties) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type CertificatePropertiesPtrOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProperties)(nil)).Elem()
}

func (o CertificatePropertiesPtrOutput) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return o
}

func (o CertificatePropertiesPtrOutput) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return o
}

func (o CertificatePropertiesPtrOutput) Elem() CertificatePropertiesOutput {
	return o.ApplyT(func(v *CertificateProperties) CertificateProperties {
		if v != nil {
			return *v
		}
		var ret CertificateProperties
		return ret
	}).(CertificatePropertiesOutput)
}

func (o CertificatePropertiesPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateProperties) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateProperties) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type CertificateResponseProperties struct {
	ExpirationDate    string `pulumi:"expirationDate"`
	IssueDate         string `pulumi:"issueDate"`
	Issuer            string `pulumi:"issuer"`
	ProvisioningState string `pulumi:"provisioningState"`
	PublicKeyHash     string `pulumi:"publicKeyHash"`
	SubjectName       string `pulumi:"subjectName"`
	Thumbprint        string `pulumi:"thumbprint"`
	Valid             bool   `pulumi:"valid"`
}

type CertificateResponsePropertiesOutput struct{ *pulumi.OutputState }

func (CertificateResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateResponseProperties)(nil)).Elem()
}

func (o CertificateResponsePropertiesOutput) ToCertificateResponsePropertiesOutput() CertificateResponsePropertiesOutput {
	return o
}

func (o CertificateResponsePropertiesOutput) ToCertificateResponsePropertiesOutputWithContext(ctx context.Context) CertificateResponsePropertiesOutput {
	return o
}

func (o CertificateResponsePropertiesOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateResponseProperties) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o CertificateResponsePropertiesOutput) IssueDate() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateResponseProperties) string { return v.IssueDate }).(pulumi.StringOutput)
}

func (o CertificateResponsePropertiesOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateResponseProperties) string { return v.Issuer }).(pulumi.StringOutput)
}

func (o CertificateResponsePropertiesOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateResponseProperties) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CertificateResponsePropertiesOutput) PublicKeyHash() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateResponseProperties) string { return v.PublicKeyHash }).(pulumi.StringOutput)
}

func (o CertificateResponsePropertiesOutput) SubjectName() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateResponseProperties) string { return v.SubjectName }).(pulumi.StringOutput)
}

func (o CertificateResponsePropertiesOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateResponseProperties) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificateResponsePropertiesOutput) Valid() pulumi.BoolOutput {
	return o.ApplyT(func(v CertificateResponseProperties) bool { return v.Valid }).(pulumi.BoolOutput)
}

type ClientRegistration struct {
	ClientId                *string `pulumi:"clientId"`
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
}





type ClientRegistrationInput interface {
	pulumi.Input

	ToClientRegistrationOutput() ClientRegistrationOutput
	ToClientRegistrationOutputWithContext(context.Context) ClientRegistrationOutput
}

type ClientRegistrationArgs struct {
	ClientId                pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecretSettingName pulumi.StringPtrInput `pulumi:"clientSecretSettingName"`
}

func (ClientRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientRegistration)(nil)).Elem()
}

func (i ClientRegistrationArgs) ToClientRegistrationOutput() ClientRegistrationOutput {
	return i.ToClientRegistrationOutputWithContext(context.Background())
}

func (i ClientRegistrationArgs) ToClientRegistrationOutputWithContext(ctx context.Context) ClientRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRegistrationOutput)
}

func (i ClientRegistrationArgs) ToClientRegistrationPtrOutput() ClientRegistrationPtrOutput {
	return i.ToClientRegistrationPtrOutputWithContext(context.Background())
}

func (i ClientRegistrationArgs) ToClientRegistrationPtrOutputWithContext(ctx context.Context) ClientRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRegistrationOutput).ToClientRegistrationPtrOutputWithContext(ctx)
}









type ClientRegistrationPtrInput interface {
	pulumi.Input

	ToClientRegistrationPtrOutput() ClientRegistrationPtrOutput
	ToClientRegistrationPtrOutputWithContext(context.Context) ClientRegistrationPtrOutput
}

type clientRegistrationPtrType ClientRegistrationArgs

func ClientRegistrationPtr(v *ClientRegistrationArgs) ClientRegistrationPtrInput {
	return (*clientRegistrationPtrType)(v)
}

func (*clientRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientRegistration)(nil)).Elem()
}

func (i *clientRegistrationPtrType) ToClientRegistrationPtrOutput() ClientRegistrationPtrOutput {
	return i.ToClientRegistrationPtrOutputWithContext(context.Background())
}

func (i *clientRegistrationPtrType) ToClientRegistrationPtrOutputWithContext(ctx context.Context) ClientRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRegistrationPtrOutput)
}

type ClientRegistrationOutput struct{ *pulumi.OutputState }

func (ClientRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientRegistration)(nil)).Elem()
}

func (o ClientRegistrationOutput) ToClientRegistrationOutput() ClientRegistrationOutput {
	return o
}

func (o ClientRegistrationOutput) ToClientRegistrationOutputWithContext(ctx context.Context) ClientRegistrationOutput {
	return o
}

func (o ClientRegistrationOutput) ToClientRegistrationPtrOutput() ClientRegistrationPtrOutput {
	return o.ToClientRegistrationPtrOutputWithContext(context.Background())
}

func (o ClientRegistrationOutput) ToClientRegistrationPtrOutputWithContext(ctx context.Context) ClientRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClientRegistration) *ClientRegistration {
		return &v
	}).(ClientRegistrationPtrOutput)
}

func (o ClientRegistrationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientRegistration) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ClientRegistrationOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientRegistration) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

type ClientRegistrationPtrOutput struct{ *pulumi.OutputState }

func (ClientRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientRegistration)(nil)).Elem()
}

func (o ClientRegistrationPtrOutput) ToClientRegistrationPtrOutput() ClientRegistrationPtrOutput {
	return o
}

func (o ClientRegistrationPtrOutput) ToClientRegistrationPtrOutputWithContext(ctx context.Context) ClientRegistrationPtrOutput {
	return o
}

func (o ClientRegistrationPtrOutput) Elem() ClientRegistrationOutput {
	return o.ApplyT(func(v *ClientRegistration) ClientRegistration {
		if v != nil {
			return *v
		}
		var ret ClientRegistration
		return ret
	}).(ClientRegistrationOutput)
}

func (o ClientRegistrationPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ClientRegistrationPtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type ClientRegistrationResponse struct {
	ClientId                *string `pulumi:"clientId"`
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
}

type ClientRegistrationResponseOutput struct{ *pulumi.OutputState }

func (ClientRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientRegistrationResponse)(nil)).Elem()
}

func (o ClientRegistrationResponseOutput) ToClientRegistrationResponseOutput() ClientRegistrationResponseOutput {
	return o
}

func (o ClientRegistrationResponseOutput) ToClientRegistrationResponseOutputWithContext(ctx context.Context) ClientRegistrationResponseOutput {
	return o
}

func (o ClientRegistrationResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientRegistrationResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ClientRegistrationResponseOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientRegistrationResponse) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

type ClientRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (ClientRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientRegistrationResponse)(nil)).Elem()
}

func (o ClientRegistrationResponsePtrOutput) ToClientRegistrationResponsePtrOutput() ClientRegistrationResponsePtrOutput {
	return o
}

func (o ClientRegistrationResponsePtrOutput) ToClientRegistrationResponsePtrOutputWithContext(ctx context.Context) ClientRegistrationResponsePtrOutput {
	return o
}

func (o ClientRegistrationResponsePtrOutput) Elem() ClientRegistrationResponseOutput {
	return o.ApplyT(func(v *ClientRegistrationResponse) ClientRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret ClientRegistrationResponse
		return ret
	}).(ClientRegistrationResponseOutput)
}

func (o ClientRegistrationResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ClientRegistrationResponsePtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type Configuration struct {
	ActiveRevisionsMode *string               `pulumi:"activeRevisionsMode"`
	Dapr                *Dapr                 `pulumi:"dapr"`
	Ingress             *Ingress              `pulumi:"ingress"`
	Registries          []RegistryCredentials `pulumi:"registries"`
	Secrets             []Secret              `pulumi:"secrets"`
}


func (val *Configuration) Defaults() *Configuration {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ActiveRevisionsMode) {
		activeRevisionsMode_ := "Single"
		tmp.ActiveRevisionsMode = &activeRevisionsMode_
	}
	tmp.Dapr = tmp.Dapr.Defaults()

	tmp.Ingress = tmp.Ingress.Defaults()

	return &tmp
}





type ConfigurationInput interface {
	pulumi.Input

	ToConfigurationOutput() ConfigurationOutput
	ToConfigurationOutputWithContext(context.Context) ConfigurationOutput
}

type ConfigurationArgs struct {
	ActiveRevisionsMode pulumi.StringPtrInput         `pulumi:"activeRevisionsMode"`
	Dapr                DaprPtrInput                  `pulumi:"dapr"`
	Ingress             IngressPtrInput               `pulumi:"ingress"`
	Registries          RegistryCredentialsArrayInput `pulumi:"registries"`
	Secrets             SecretArrayInput              `pulumi:"secrets"`
}


func (val *ConfigurationArgs) Defaults() *ConfigurationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ActiveRevisionsMode) {
		tmp.ActiveRevisionsMode = pulumi.StringPtr("Single")
	}

	return &tmp
}
func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Configuration)(nil)).Elem()
}

func (i ConfigurationArgs) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i ConfigurationArgs) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

func (i ConfigurationArgs) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return i.ToConfigurationPtrOutputWithContext(context.Background())
}

func (i ConfigurationArgs) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput).ToConfigurationPtrOutputWithContext(ctx)
}









type ConfigurationPtrInput interface {
	pulumi.Input

	ToConfigurationPtrOutput() ConfigurationPtrOutput
	ToConfigurationPtrOutputWithContext(context.Context) ConfigurationPtrOutput
}

type configurationPtrType ConfigurationArgs

func ConfigurationPtr(v *ConfigurationArgs) ConfigurationPtrInput {
	return (*configurationPtrType)(v)
}

func (*configurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (i *configurationPtrType) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return i.ToConfigurationPtrOutputWithContext(context.Background())
}

func (i *configurationPtrType) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationPtrOutput)
}

type ConfigurationOutput struct{ *pulumi.OutputState }

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Configuration)(nil)).Elem()
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return o.ToConfigurationPtrOutputWithContext(context.Background())
}

func (o ConfigurationOutput) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Configuration) *Configuration {
		return &v
	}).(ConfigurationPtrOutput)
}

func (o ConfigurationOutput) ActiveRevisionsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Configuration) *string { return v.ActiveRevisionsMode }).(pulumi.StringPtrOutput)
}

func (o ConfigurationOutput) Dapr() DaprPtrOutput {
	return o.ApplyT(func(v Configuration) *Dapr { return v.Dapr }).(DaprPtrOutput)
}

func (o ConfigurationOutput) Ingress() IngressPtrOutput {
	return o.ApplyT(func(v Configuration) *Ingress { return v.Ingress }).(IngressPtrOutput)
}

func (o ConfigurationOutput) Registries() RegistryCredentialsArrayOutput {
	return o.ApplyT(func(v Configuration) []RegistryCredentials { return v.Registries }).(RegistryCredentialsArrayOutput)
}

func (o ConfigurationOutput) Secrets() SecretArrayOutput {
	return o.ApplyT(func(v Configuration) []Secret { return v.Secrets }).(SecretArrayOutput)
}

type ConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (o ConfigurationPtrOutput) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return o
}

func (o ConfigurationPtrOutput) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return o
}

func (o ConfigurationPtrOutput) Elem() ConfigurationOutput {
	return o.ApplyT(func(v *Configuration) Configuration {
		if v != nil {
			return *v
		}
		var ret Configuration
		return ret
	}).(ConfigurationOutput)
}

func (o ConfigurationPtrOutput) ActiveRevisionsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Configuration) *string {
		if v == nil {
			return nil
		}
		return v.ActiveRevisionsMode
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationPtrOutput) Dapr() DaprPtrOutput {
	return o.ApplyT(func(v *Configuration) *Dapr {
		if v == nil {
			return nil
		}
		return v.Dapr
	}).(DaprPtrOutput)
}

func (o ConfigurationPtrOutput) Ingress() IngressPtrOutput {
	return o.ApplyT(func(v *Configuration) *Ingress {
		if v == nil {
			return nil
		}
		return v.Ingress
	}).(IngressPtrOutput)
}

func (o ConfigurationPtrOutput) Registries() RegistryCredentialsArrayOutput {
	return o.ApplyT(func(v *Configuration) []RegistryCredentials {
		if v == nil {
			return nil
		}
		return v.Registries
	}).(RegistryCredentialsArrayOutput)
}

func (o ConfigurationPtrOutput) Secrets() SecretArrayOutput {
	return o.ApplyT(func(v *Configuration) []Secret {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(SecretArrayOutput)
}

type ConfigurationResponse struct {
	ActiveRevisionsMode *string                       `pulumi:"activeRevisionsMode"`
	Dapr                *DaprResponse                 `pulumi:"dapr"`
	Ingress             *IngressResponse              `pulumi:"ingress"`
	Registries          []RegistryCredentialsResponse `pulumi:"registries"`
	Secrets             []SecretResponse              `pulumi:"secrets"`
}


func (val *ConfigurationResponse) Defaults() *ConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ActiveRevisionsMode) {
		activeRevisionsMode_ := "Single"
		tmp.ActiveRevisionsMode = &activeRevisionsMode_
	}
	tmp.Dapr = tmp.Dapr.Defaults()

	tmp.Ingress = tmp.Ingress.Defaults()

	return &tmp
}

type ConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationResponse)(nil)).Elem()
}

func (o ConfigurationResponseOutput) ToConfigurationResponseOutput() ConfigurationResponseOutput {
	return o
}

func (o ConfigurationResponseOutput) ToConfigurationResponseOutputWithContext(ctx context.Context) ConfigurationResponseOutput {
	return o
}

func (o ConfigurationResponseOutput) ActiveRevisionsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationResponse) *string { return v.ActiveRevisionsMode }).(pulumi.StringPtrOutput)
}

func (o ConfigurationResponseOutput) Dapr() DaprResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationResponse) *DaprResponse { return v.Dapr }).(DaprResponsePtrOutput)
}

func (o ConfigurationResponseOutput) Ingress() IngressResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationResponse) *IngressResponse { return v.Ingress }).(IngressResponsePtrOutput)
}

func (o ConfigurationResponseOutput) Registries() RegistryCredentialsResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationResponse) []RegistryCredentialsResponse { return v.Registries }).(RegistryCredentialsResponseArrayOutput)
}

func (o ConfigurationResponseOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationResponse) []SecretResponse { return v.Secrets }).(SecretResponseArrayOutput)
}

type ConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationResponse)(nil)).Elem()
}

func (o ConfigurationResponsePtrOutput) ToConfigurationResponsePtrOutput() ConfigurationResponsePtrOutput {
	return o
}

func (o ConfigurationResponsePtrOutput) ToConfigurationResponsePtrOutputWithContext(ctx context.Context) ConfigurationResponsePtrOutput {
	return o
}

func (o ConfigurationResponsePtrOutput) Elem() ConfigurationResponseOutput {
	return o.ApplyT(func(v *ConfigurationResponse) ConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationResponse
		return ret
	}).(ConfigurationResponseOutput)
}

func (o ConfigurationResponsePtrOutput) ActiveRevisionsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActiveRevisionsMode
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationResponsePtrOutput) Dapr() DaprResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationResponse) *DaprResponse {
		if v == nil {
			return nil
		}
		return v.Dapr
	}).(DaprResponsePtrOutput)
}

func (o ConfigurationResponsePtrOutput) Ingress() IngressResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationResponse) *IngressResponse {
		if v == nil {
			return nil
		}
		return v.Ingress
	}).(IngressResponsePtrOutput)
}

func (o ConfigurationResponsePtrOutput) Registries() RegistryCredentialsResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationResponse) []RegistryCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.Registries
	}).(RegistryCredentialsResponseArrayOutput)
}

func (o ConfigurationResponsePtrOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationResponse) []SecretResponse {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(SecretResponseArrayOutput)
}

type Container struct {
	Args         []string            `pulumi:"args"`
	Command      []string            `pulumi:"command"`
	Env          []EnvironmentVar    `pulumi:"env"`
	Image        *string             `pulumi:"image"`
	Name         *string             `pulumi:"name"`
	Probes       []ContainerAppProbe `pulumi:"probes"`
	Resources    *ContainerResources `pulumi:"resources"`
	VolumeMounts []VolumeMount       `pulumi:"volumeMounts"`
}





type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(context.Context) ContainerOutput
}

type ContainerArgs struct {
	Args         pulumi.StringArrayInput     `pulumi:"args"`
	Command      pulumi.StringArrayInput     `pulumi:"command"`
	Env          EnvironmentVarArrayInput    `pulumi:"env"`
	Image        pulumi.StringPtrInput       `pulumi:"image"`
	Name         pulumi.StringPtrInput       `pulumi:"name"`
	Probes       ContainerAppProbeArrayInput `pulumi:"probes"`
	Resources    ContainerResourcesPtrInput  `pulumi:"resources"`
	VolumeMounts VolumeMountArrayInput       `pulumi:"volumeMounts"`
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (i ContainerArgs) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i ContainerArgs) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}





type ContainerArrayInput interface {
	pulumi.Input

	ToContainerArrayOutput() ContainerArrayOutput
	ToContainerArrayOutputWithContext(context.Context) ContainerArrayOutput
}

type ContainerArray []ContainerInput

func (ContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Container)(nil)).Elem()
}

func (i ContainerArray) ToContainerArrayOutput() ContainerArrayOutput {
	return i.ToContainerArrayOutputWithContext(context.Background())
}

func (i ContainerArray) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerArrayOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

func (o ContainerOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Container) []string { return v.Args }).(pulumi.StringArrayOutput)
}

func (o ContainerOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Container) []string { return v.Command }).(pulumi.StringArrayOutput)
}

func (o ContainerOutput) Env() EnvironmentVarArrayOutput {
	return o.ApplyT(func(v Container) []EnvironmentVar { return v.Env }).(EnvironmentVarArrayOutput)
}

func (o ContainerOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Container) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o ContainerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Container) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ContainerOutput) Probes() ContainerAppProbeArrayOutput {
	return o.ApplyT(func(v Container) []ContainerAppProbe { return v.Probes }).(ContainerAppProbeArrayOutput)
}

func (o ContainerOutput) Resources() ContainerResourcesPtrOutput {
	return o.ApplyT(func(v Container) *ContainerResources { return v.Resources }).(ContainerResourcesPtrOutput)
}

func (o ContainerOutput) VolumeMounts() VolumeMountArrayOutput {
	return o.ApplyT(func(v Container) []VolumeMount { return v.VolumeMounts }).(VolumeMountArrayOutput)
}

type ContainerArrayOutput struct{ *pulumi.OutputState }

func (ContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Container)(nil)).Elem()
}

func (o ContainerArrayOutput) ToContainerArrayOutput() ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) Index(i pulumi.IntInput) ContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Container {
		return vs[0].([]Container)[vs[1].(int)]
	}).(ContainerOutput)
}

type ContainerAppProbe struct {
	FailureThreshold              *int                        `pulumi:"failureThreshold"`
	HttpGet                       *ContainerAppProbeHttpGet   `pulumi:"httpGet"`
	InitialDelaySeconds           *int                        `pulumi:"initialDelaySeconds"`
	PeriodSeconds                 *int                        `pulumi:"periodSeconds"`
	SuccessThreshold              *int                        `pulumi:"successThreshold"`
	TcpSocket                     *ContainerAppProbeTcpSocket `pulumi:"tcpSocket"`
	TerminationGracePeriodSeconds *float64                    `pulumi:"terminationGracePeriodSeconds"`
	TimeoutSeconds                *int                        `pulumi:"timeoutSeconds"`
	Type                          *string                     `pulumi:"type"`
}





type ContainerAppProbeInput interface {
	pulumi.Input

	ToContainerAppProbeOutput() ContainerAppProbeOutput
	ToContainerAppProbeOutputWithContext(context.Context) ContainerAppProbeOutput
}

type ContainerAppProbeArgs struct {
	FailureThreshold              pulumi.IntPtrInput                 `pulumi:"failureThreshold"`
	HttpGet                       ContainerAppProbeHttpGetPtrInput   `pulumi:"httpGet"`
	InitialDelaySeconds           pulumi.IntPtrInput                 `pulumi:"initialDelaySeconds"`
	PeriodSeconds                 pulumi.IntPtrInput                 `pulumi:"periodSeconds"`
	SuccessThreshold              pulumi.IntPtrInput                 `pulumi:"successThreshold"`
	TcpSocket                     ContainerAppProbeTcpSocketPtrInput `pulumi:"tcpSocket"`
	TerminationGracePeriodSeconds pulumi.Float64PtrInput             `pulumi:"terminationGracePeriodSeconds"`
	TimeoutSeconds                pulumi.IntPtrInput                 `pulumi:"timeoutSeconds"`
	Type                          pulumi.StringPtrInput              `pulumi:"type"`
}

func (ContainerAppProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbe)(nil)).Elem()
}

func (i ContainerAppProbeArgs) ToContainerAppProbeOutput() ContainerAppProbeOutput {
	return i.ToContainerAppProbeOutputWithContext(context.Background())
}

func (i ContainerAppProbeArgs) ToContainerAppProbeOutputWithContext(ctx context.Context) ContainerAppProbeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppProbeOutput)
}





type ContainerAppProbeArrayInput interface {
	pulumi.Input

	ToContainerAppProbeArrayOutput() ContainerAppProbeArrayOutput
	ToContainerAppProbeArrayOutputWithContext(context.Context) ContainerAppProbeArrayOutput
}

type ContainerAppProbeArray []ContainerAppProbeInput

func (ContainerAppProbeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerAppProbe)(nil)).Elem()
}

func (i ContainerAppProbeArray) ToContainerAppProbeArrayOutput() ContainerAppProbeArrayOutput {
	return i.ToContainerAppProbeArrayOutputWithContext(context.Background())
}

func (i ContainerAppProbeArray) ToContainerAppProbeArrayOutputWithContext(ctx context.Context) ContainerAppProbeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppProbeArrayOutput)
}

type ContainerAppProbeOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbe)(nil)).Elem()
}

func (o ContainerAppProbeOutput) ToContainerAppProbeOutput() ContainerAppProbeOutput {
	return o
}

func (o ContainerAppProbeOutput) ToContainerAppProbeOutputWithContext(ctx context.Context) ContainerAppProbeOutput {
	return o
}

func (o ContainerAppProbeOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerAppProbe) *int { return v.FailureThreshold }).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeOutput) HttpGet() ContainerAppProbeHttpGetPtrOutput {
	return o.ApplyT(func(v ContainerAppProbe) *ContainerAppProbeHttpGet { return v.HttpGet }).(ContainerAppProbeHttpGetPtrOutput)
}

func (o ContainerAppProbeOutput) InitialDelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerAppProbe) *int { return v.InitialDelaySeconds }).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeOutput) PeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerAppProbe) *int { return v.PeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeOutput) SuccessThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerAppProbe) *int { return v.SuccessThreshold }).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeOutput) TcpSocket() ContainerAppProbeTcpSocketPtrOutput {
	return o.ApplyT(func(v ContainerAppProbe) *ContainerAppProbeTcpSocket { return v.TcpSocket }).(ContainerAppProbeTcpSocketPtrOutput)
}

func (o ContainerAppProbeOutput) TerminationGracePeriodSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerAppProbe) *float64 { return v.TerminationGracePeriodSeconds }).(pulumi.Float64PtrOutput)
}

func (o ContainerAppProbeOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerAppProbe) *int { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppProbe) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ContainerAppProbeArrayOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerAppProbe)(nil)).Elem()
}

func (o ContainerAppProbeArrayOutput) ToContainerAppProbeArrayOutput() ContainerAppProbeArrayOutput {
	return o
}

func (o ContainerAppProbeArrayOutput) ToContainerAppProbeArrayOutputWithContext(ctx context.Context) ContainerAppProbeArrayOutput {
	return o
}

func (o ContainerAppProbeArrayOutput) Index(i pulumi.IntInput) ContainerAppProbeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerAppProbe {
		return vs[0].([]ContainerAppProbe)[vs[1].(int)]
	}).(ContainerAppProbeOutput)
}

type ContainerAppProbeHttpGet struct {
	Host        *string                        `pulumi:"host"`
	HttpHeaders []ContainerAppProbeHttpHeaders `pulumi:"httpHeaders"`
	Path        *string                        `pulumi:"path"`
	Port        int                            `pulumi:"port"`
	Scheme      *string                        `pulumi:"scheme"`
}





type ContainerAppProbeHttpGetInput interface {
	pulumi.Input

	ToContainerAppProbeHttpGetOutput() ContainerAppProbeHttpGetOutput
	ToContainerAppProbeHttpGetOutputWithContext(context.Context) ContainerAppProbeHttpGetOutput
}

type ContainerAppProbeHttpGetArgs struct {
	Host        pulumi.StringPtrInput                  `pulumi:"host"`
	HttpHeaders ContainerAppProbeHttpHeadersArrayInput `pulumi:"httpHeaders"`
	Path        pulumi.StringPtrInput                  `pulumi:"path"`
	Port        pulumi.IntInput                        `pulumi:"port"`
	Scheme      pulumi.StringPtrInput                  `pulumi:"scheme"`
}

func (ContainerAppProbeHttpGetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbeHttpGet)(nil)).Elem()
}

func (i ContainerAppProbeHttpGetArgs) ToContainerAppProbeHttpGetOutput() ContainerAppProbeHttpGetOutput {
	return i.ToContainerAppProbeHttpGetOutputWithContext(context.Background())
}

func (i ContainerAppProbeHttpGetArgs) ToContainerAppProbeHttpGetOutputWithContext(ctx context.Context) ContainerAppProbeHttpGetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppProbeHttpGetOutput)
}

func (i ContainerAppProbeHttpGetArgs) ToContainerAppProbeHttpGetPtrOutput() ContainerAppProbeHttpGetPtrOutput {
	return i.ToContainerAppProbeHttpGetPtrOutputWithContext(context.Background())
}

func (i ContainerAppProbeHttpGetArgs) ToContainerAppProbeHttpGetPtrOutputWithContext(ctx context.Context) ContainerAppProbeHttpGetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppProbeHttpGetOutput).ToContainerAppProbeHttpGetPtrOutputWithContext(ctx)
}









type ContainerAppProbeHttpGetPtrInput interface {
	pulumi.Input

	ToContainerAppProbeHttpGetPtrOutput() ContainerAppProbeHttpGetPtrOutput
	ToContainerAppProbeHttpGetPtrOutputWithContext(context.Context) ContainerAppProbeHttpGetPtrOutput
}

type containerAppProbeHttpGetPtrType ContainerAppProbeHttpGetArgs

func ContainerAppProbeHttpGetPtr(v *ContainerAppProbeHttpGetArgs) ContainerAppProbeHttpGetPtrInput {
	return (*containerAppProbeHttpGetPtrType)(v)
}

func (*containerAppProbeHttpGetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppProbeHttpGet)(nil)).Elem()
}

func (i *containerAppProbeHttpGetPtrType) ToContainerAppProbeHttpGetPtrOutput() ContainerAppProbeHttpGetPtrOutput {
	return i.ToContainerAppProbeHttpGetPtrOutputWithContext(context.Background())
}

func (i *containerAppProbeHttpGetPtrType) ToContainerAppProbeHttpGetPtrOutputWithContext(ctx context.Context) ContainerAppProbeHttpGetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppProbeHttpGetPtrOutput)
}

type ContainerAppProbeHttpGetOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeHttpGetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbeHttpGet)(nil)).Elem()
}

func (o ContainerAppProbeHttpGetOutput) ToContainerAppProbeHttpGetOutput() ContainerAppProbeHttpGetOutput {
	return o
}

func (o ContainerAppProbeHttpGetOutput) ToContainerAppProbeHttpGetOutputWithContext(ctx context.Context) ContainerAppProbeHttpGetOutput {
	return o
}

func (o ContainerAppProbeHttpGetOutput) ToContainerAppProbeHttpGetPtrOutput() ContainerAppProbeHttpGetPtrOutput {
	return o.ToContainerAppProbeHttpGetPtrOutputWithContext(context.Background())
}

func (o ContainerAppProbeHttpGetOutput) ToContainerAppProbeHttpGetPtrOutputWithContext(ctx context.Context) ContainerAppProbeHttpGetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerAppProbeHttpGet) *ContainerAppProbeHttpGet {
		return &v
	}).(ContainerAppProbeHttpGetPtrOutput)
}

func (o ContainerAppProbeHttpGetOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeHttpGet) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeHttpGetOutput) HttpHeaders() ContainerAppProbeHttpHeadersArrayOutput {
	return o.ApplyT(func(v ContainerAppProbeHttpGet) []ContainerAppProbeHttpHeaders { return v.HttpHeaders }).(ContainerAppProbeHttpHeadersArrayOutput)
}

func (o ContainerAppProbeHttpGetOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeHttpGet) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeHttpGetOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerAppProbeHttpGet) int { return v.Port }).(pulumi.IntOutput)
}

func (o ContainerAppProbeHttpGetOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeHttpGet) *string { return v.Scheme }).(pulumi.StringPtrOutput)
}

type ContainerAppProbeHttpGetPtrOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeHttpGetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppProbeHttpGet)(nil)).Elem()
}

func (o ContainerAppProbeHttpGetPtrOutput) ToContainerAppProbeHttpGetPtrOutput() ContainerAppProbeHttpGetPtrOutput {
	return o
}

func (o ContainerAppProbeHttpGetPtrOutput) ToContainerAppProbeHttpGetPtrOutputWithContext(ctx context.Context) ContainerAppProbeHttpGetPtrOutput {
	return o
}

func (o ContainerAppProbeHttpGetPtrOutput) Elem() ContainerAppProbeHttpGetOutput {
	return o.ApplyT(func(v *ContainerAppProbeHttpGet) ContainerAppProbeHttpGet {
		if v != nil {
			return *v
		}
		var ret ContainerAppProbeHttpGet
		return ret
	}).(ContainerAppProbeHttpGetOutput)
}

func (o ContainerAppProbeHttpGetPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeHttpGet) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeHttpGetPtrOutput) HttpHeaders() ContainerAppProbeHttpHeadersArrayOutput {
	return o.ApplyT(func(v *ContainerAppProbeHttpGet) []ContainerAppProbeHttpHeaders {
		if v == nil {
			return nil
		}
		return v.HttpHeaders
	}).(ContainerAppProbeHttpHeadersArrayOutput)
}

func (o ContainerAppProbeHttpGetPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeHttpGet) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeHttpGetPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeHttpGet) *int {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeHttpGetPtrOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeHttpGet) *string {
		if v == nil {
			return nil
		}
		return v.Scheme
	}).(pulumi.StringPtrOutput)
}

type ContainerAppProbeHttpHeaders struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type ContainerAppProbeHttpHeadersInput interface {
	pulumi.Input

	ToContainerAppProbeHttpHeadersOutput() ContainerAppProbeHttpHeadersOutput
	ToContainerAppProbeHttpHeadersOutputWithContext(context.Context) ContainerAppProbeHttpHeadersOutput
}

type ContainerAppProbeHttpHeadersArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (ContainerAppProbeHttpHeadersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbeHttpHeaders)(nil)).Elem()
}

func (i ContainerAppProbeHttpHeadersArgs) ToContainerAppProbeHttpHeadersOutput() ContainerAppProbeHttpHeadersOutput {
	return i.ToContainerAppProbeHttpHeadersOutputWithContext(context.Background())
}

func (i ContainerAppProbeHttpHeadersArgs) ToContainerAppProbeHttpHeadersOutputWithContext(ctx context.Context) ContainerAppProbeHttpHeadersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppProbeHttpHeadersOutput)
}





type ContainerAppProbeHttpHeadersArrayInput interface {
	pulumi.Input

	ToContainerAppProbeHttpHeadersArrayOutput() ContainerAppProbeHttpHeadersArrayOutput
	ToContainerAppProbeHttpHeadersArrayOutputWithContext(context.Context) ContainerAppProbeHttpHeadersArrayOutput
}

type ContainerAppProbeHttpHeadersArray []ContainerAppProbeHttpHeadersInput

func (ContainerAppProbeHttpHeadersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerAppProbeHttpHeaders)(nil)).Elem()
}

func (i ContainerAppProbeHttpHeadersArray) ToContainerAppProbeHttpHeadersArrayOutput() ContainerAppProbeHttpHeadersArrayOutput {
	return i.ToContainerAppProbeHttpHeadersArrayOutputWithContext(context.Background())
}

func (i ContainerAppProbeHttpHeadersArray) ToContainerAppProbeHttpHeadersArrayOutputWithContext(ctx context.Context) ContainerAppProbeHttpHeadersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppProbeHttpHeadersArrayOutput)
}

type ContainerAppProbeHttpHeadersOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeHttpHeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbeHttpHeaders)(nil)).Elem()
}

func (o ContainerAppProbeHttpHeadersOutput) ToContainerAppProbeHttpHeadersOutput() ContainerAppProbeHttpHeadersOutput {
	return o
}

func (o ContainerAppProbeHttpHeadersOutput) ToContainerAppProbeHttpHeadersOutputWithContext(ctx context.Context) ContainerAppProbeHttpHeadersOutput {
	return o
}

func (o ContainerAppProbeHttpHeadersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerAppProbeHttpHeaders) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerAppProbeHttpHeadersOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerAppProbeHttpHeaders) string { return v.Value }).(pulumi.StringOutput)
}

type ContainerAppProbeHttpHeadersArrayOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeHttpHeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerAppProbeHttpHeaders)(nil)).Elem()
}

func (o ContainerAppProbeHttpHeadersArrayOutput) ToContainerAppProbeHttpHeadersArrayOutput() ContainerAppProbeHttpHeadersArrayOutput {
	return o
}

func (o ContainerAppProbeHttpHeadersArrayOutput) ToContainerAppProbeHttpHeadersArrayOutputWithContext(ctx context.Context) ContainerAppProbeHttpHeadersArrayOutput {
	return o
}

func (o ContainerAppProbeHttpHeadersArrayOutput) Index(i pulumi.IntInput) ContainerAppProbeHttpHeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerAppProbeHttpHeaders {
		return vs[0].([]ContainerAppProbeHttpHeaders)[vs[1].(int)]
	}).(ContainerAppProbeHttpHeadersOutput)
}

type ContainerAppProbeResponse struct {
	FailureThreshold              *int                                `pulumi:"failureThreshold"`
	HttpGet                       *ContainerAppProbeResponseHttpGet   `pulumi:"httpGet"`
	InitialDelaySeconds           *int                                `pulumi:"initialDelaySeconds"`
	PeriodSeconds                 *int                                `pulumi:"periodSeconds"`
	SuccessThreshold              *int                                `pulumi:"successThreshold"`
	TcpSocket                     *ContainerAppProbeResponseTcpSocket `pulumi:"tcpSocket"`
	TerminationGracePeriodSeconds *float64                            `pulumi:"terminationGracePeriodSeconds"`
	TimeoutSeconds                *int                                `pulumi:"timeoutSeconds"`
	Type                          *string                             `pulumi:"type"`
}

type ContainerAppProbeResponseOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbeResponse)(nil)).Elem()
}

func (o ContainerAppProbeResponseOutput) ToContainerAppProbeResponseOutput() ContainerAppProbeResponseOutput {
	return o
}

func (o ContainerAppProbeResponseOutput) ToContainerAppProbeResponseOutputWithContext(ctx context.Context) ContainerAppProbeResponseOutput {
	return o
}

func (o ContainerAppProbeResponseOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponse) *int { return v.FailureThreshold }).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeResponseOutput) HttpGet() ContainerAppProbeResponseHttpGetPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponse) *ContainerAppProbeResponseHttpGet { return v.HttpGet }).(ContainerAppProbeResponseHttpGetPtrOutput)
}

func (o ContainerAppProbeResponseOutput) InitialDelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponse) *int { return v.InitialDelaySeconds }).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeResponseOutput) PeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponse) *int { return v.PeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeResponseOutput) SuccessThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponse) *int { return v.SuccessThreshold }).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeResponseOutput) TcpSocket() ContainerAppProbeResponseTcpSocketPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponse) *ContainerAppProbeResponseTcpSocket { return v.TcpSocket }).(ContainerAppProbeResponseTcpSocketPtrOutput)
}

func (o ContainerAppProbeResponseOutput) TerminationGracePeriodSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponse) *float64 { return v.TerminationGracePeriodSeconds }).(pulumi.Float64PtrOutput)
}

func (o ContainerAppProbeResponseOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponse) *int { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ContainerAppProbeResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerAppProbeResponse)(nil)).Elem()
}

func (o ContainerAppProbeResponseArrayOutput) ToContainerAppProbeResponseArrayOutput() ContainerAppProbeResponseArrayOutput {
	return o
}

func (o ContainerAppProbeResponseArrayOutput) ToContainerAppProbeResponseArrayOutputWithContext(ctx context.Context) ContainerAppProbeResponseArrayOutput {
	return o
}

func (o ContainerAppProbeResponseArrayOutput) Index(i pulumi.IntInput) ContainerAppProbeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerAppProbeResponse {
		return vs[0].([]ContainerAppProbeResponse)[vs[1].(int)]
	}).(ContainerAppProbeResponseOutput)
}

type ContainerAppProbeResponseHttpGet struct {
	Host        *string                                `pulumi:"host"`
	HttpHeaders []ContainerAppProbeResponseHttpHeaders `pulumi:"httpHeaders"`
	Path        *string                                `pulumi:"path"`
	Port        int                                    `pulumi:"port"`
	Scheme      *string                                `pulumi:"scheme"`
}

type ContainerAppProbeResponseHttpGetOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeResponseHttpGetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbeResponseHttpGet)(nil)).Elem()
}

func (o ContainerAppProbeResponseHttpGetOutput) ToContainerAppProbeResponseHttpGetOutput() ContainerAppProbeResponseHttpGetOutput {
	return o
}

func (o ContainerAppProbeResponseHttpGetOutput) ToContainerAppProbeResponseHttpGetOutputWithContext(ctx context.Context) ContainerAppProbeResponseHttpGetOutput {
	return o
}

func (o ContainerAppProbeResponseHttpGetOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponseHttpGet) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeResponseHttpGetOutput) HttpHeaders() ContainerAppProbeResponseHttpHeadersArrayOutput {
	return o.ApplyT(func(v ContainerAppProbeResponseHttpGet) []ContainerAppProbeResponseHttpHeaders { return v.HttpHeaders }).(ContainerAppProbeResponseHttpHeadersArrayOutput)
}

func (o ContainerAppProbeResponseHttpGetOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponseHttpGet) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeResponseHttpGetOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerAppProbeResponseHttpGet) int { return v.Port }).(pulumi.IntOutput)
}

func (o ContainerAppProbeResponseHttpGetOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponseHttpGet) *string { return v.Scheme }).(pulumi.StringPtrOutput)
}

type ContainerAppProbeResponseHttpGetPtrOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeResponseHttpGetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppProbeResponseHttpGet)(nil)).Elem()
}

func (o ContainerAppProbeResponseHttpGetPtrOutput) ToContainerAppProbeResponseHttpGetPtrOutput() ContainerAppProbeResponseHttpGetPtrOutput {
	return o
}

func (o ContainerAppProbeResponseHttpGetPtrOutput) ToContainerAppProbeResponseHttpGetPtrOutputWithContext(ctx context.Context) ContainerAppProbeResponseHttpGetPtrOutput {
	return o
}

func (o ContainerAppProbeResponseHttpGetPtrOutput) Elem() ContainerAppProbeResponseHttpGetOutput {
	return o.ApplyT(func(v *ContainerAppProbeResponseHttpGet) ContainerAppProbeResponseHttpGet {
		if v != nil {
			return *v
		}
		var ret ContainerAppProbeResponseHttpGet
		return ret
	}).(ContainerAppProbeResponseHttpGetOutput)
}

func (o ContainerAppProbeResponseHttpGetPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeResponseHttpGet) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeResponseHttpGetPtrOutput) HttpHeaders() ContainerAppProbeResponseHttpHeadersArrayOutput {
	return o.ApplyT(func(v *ContainerAppProbeResponseHttpGet) []ContainerAppProbeResponseHttpHeaders {
		if v == nil {
			return nil
		}
		return v.HttpHeaders
	}).(ContainerAppProbeResponseHttpHeadersArrayOutput)
}

func (o ContainerAppProbeResponseHttpGetPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeResponseHttpGet) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeResponseHttpGetPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeResponseHttpGet) *int {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.IntPtrOutput)
}

func (o ContainerAppProbeResponseHttpGetPtrOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeResponseHttpGet) *string {
		if v == nil {
			return nil
		}
		return v.Scheme
	}).(pulumi.StringPtrOutput)
}

type ContainerAppProbeResponseHttpHeaders struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type ContainerAppProbeResponseHttpHeadersOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeResponseHttpHeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbeResponseHttpHeaders)(nil)).Elem()
}

func (o ContainerAppProbeResponseHttpHeadersOutput) ToContainerAppProbeResponseHttpHeadersOutput() ContainerAppProbeResponseHttpHeadersOutput {
	return o
}

func (o ContainerAppProbeResponseHttpHeadersOutput) ToContainerAppProbeResponseHttpHeadersOutputWithContext(ctx context.Context) ContainerAppProbeResponseHttpHeadersOutput {
	return o
}

func (o ContainerAppProbeResponseHttpHeadersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerAppProbeResponseHttpHeaders) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerAppProbeResponseHttpHeadersOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerAppProbeResponseHttpHeaders) string { return v.Value }).(pulumi.StringOutput)
}

type ContainerAppProbeResponseHttpHeadersArrayOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeResponseHttpHeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerAppProbeResponseHttpHeaders)(nil)).Elem()
}

func (o ContainerAppProbeResponseHttpHeadersArrayOutput) ToContainerAppProbeResponseHttpHeadersArrayOutput() ContainerAppProbeResponseHttpHeadersArrayOutput {
	return o
}

func (o ContainerAppProbeResponseHttpHeadersArrayOutput) ToContainerAppProbeResponseHttpHeadersArrayOutputWithContext(ctx context.Context) ContainerAppProbeResponseHttpHeadersArrayOutput {
	return o
}

func (o ContainerAppProbeResponseHttpHeadersArrayOutput) Index(i pulumi.IntInput) ContainerAppProbeResponseHttpHeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerAppProbeResponseHttpHeaders {
		return vs[0].([]ContainerAppProbeResponseHttpHeaders)[vs[1].(int)]
	}).(ContainerAppProbeResponseHttpHeadersOutput)
}

type ContainerAppProbeResponseTcpSocket struct {
	Host *string `pulumi:"host"`
	Port int     `pulumi:"port"`
}

type ContainerAppProbeResponseTcpSocketOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeResponseTcpSocketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbeResponseTcpSocket)(nil)).Elem()
}

func (o ContainerAppProbeResponseTcpSocketOutput) ToContainerAppProbeResponseTcpSocketOutput() ContainerAppProbeResponseTcpSocketOutput {
	return o
}

func (o ContainerAppProbeResponseTcpSocketOutput) ToContainerAppProbeResponseTcpSocketOutputWithContext(ctx context.Context) ContainerAppProbeResponseTcpSocketOutput {
	return o
}

func (o ContainerAppProbeResponseTcpSocketOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeResponseTcpSocket) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeResponseTcpSocketOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerAppProbeResponseTcpSocket) int { return v.Port }).(pulumi.IntOutput)
}

type ContainerAppProbeResponseTcpSocketPtrOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeResponseTcpSocketPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppProbeResponseTcpSocket)(nil)).Elem()
}

func (o ContainerAppProbeResponseTcpSocketPtrOutput) ToContainerAppProbeResponseTcpSocketPtrOutput() ContainerAppProbeResponseTcpSocketPtrOutput {
	return o
}

func (o ContainerAppProbeResponseTcpSocketPtrOutput) ToContainerAppProbeResponseTcpSocketPtrOutputWithContext(ctx context.Context) ContainerAppProbeResponseTcpSocketPtrOutput {
	return o
}

func (o ContainerAppProbeResponseTcpSocketPtrOutput) Elem() ContainerAppProbeResponseTcpSocketOutput {
	return o.ApplyT(func(v *ContainerAppProbeResponseTcpSocket) ContainerAppProbeResponseTcpSocket {
		if v != nil {
			return *v
		}
		var ret ContainerAppProbeResponseTcpSocket
		return ret
	}).(ContainerAppProbeResponseTcpSocketOutput)
}

func (o ContainerAppProbeResponseTcpSocketPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeResponseTcpSocket) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeResponseTcpSocketPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeResponseTcpSocket) *int {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.IntPtrOutput)
}

type ContainerAppProbeTcpSocket struct {
	Host *string `pulumi:"host"`
	Port int     `pulumi:"port"`
}





type ContainerAppProbeTcpSocketInput interface {
	pulumi.Input

	ToContainerAppProbeTcpSocketOutput() ContainerAppProbeTcpSocketOutput
	ToContainerAppProbeTcpSocketOutputWithContext(context.Context) ContainerAppProbeTcpSocketOutput
}

type ContainerAppProbeTcpSocketArgs struct {
	Host pulumi.StringPtrInput `pulumi:"host"`
	Port pulumi.IntInput       `pulumi:"port"`
}

func (ContainerAppProbeTcpSocketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbeTcpSocket)(nil)).Elem()
}

func (i ContainerAppProbeTcpSocketArgs) ToContainerAppProbeTcpSocketOutput() ContainerAppProbeTcpSocketOutput {
	return i.ToContainerAppProbeTcpSocketOutputWithContext(context.Background())
}

func (i ContainerAppProbeTcpSocketArgs) ToContainerAppProbeTcpSocketOutputWithContext(ctx context.Context) ContainerAppProbeTcpSocketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppProbeTcpSocketOutput)
}

func (i ContainerAppProbeTcpSocketArgs) ToContainerAppProbeTcpSocketPtrOutput() ContainerAppProbeTcpSocketPtrOutput {
	return i.ToContainerAppProbeTcpSocketPtrOutputWithContext(context.Background())
}

func (i ContainerAppProbeTcpSocketArgs) ToContainerAppProbeTcpSocketPtrOutputWithContext(ctx context.Context) ContainerAppProbeTcpSocketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppProbeTcpSocketOutput).ToContainerAppProbeTcpSocketPtrOutputWithContext(ctx)
}









type ContainerAppProbeTcpSocketPtrInput interface {
	pulumi.Input

	ToContainerAppProbeTcpSocketPtrOutput() ContainerAppProbeTcpSocketPtrOutput
	ToContainerAppProbeTcpSocketPtrOutputWithContext(context.Context) ContainerAppProbeTcpSocketPtrOutput
}

type containerAppProbeTcpSocketPtrType ContainerAppProbeTcpSocketArgs

func ContainerAppProbeTcpSocketPtr(v *ContainerAppProbeTcpSocketArgs) ContainerAppProbeTcpSocketPtrInput {
	return (*containerAppProbeTcpSocketPtrType)(v)
}

func (*containerAppProbeTcpSocketPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppProbeTcpSocket)(nil)).Elem()
}

func (i *containerAppProbeTcpSocketPtrType) ToContainerAppProbeTcpSocketPtrOutput() ContainerAppProbeTcpSocketPtrOutput {
	return i.ToContainerAppProbeTcpSocketPtrOutputWithContext(context.Background())
}

func (i *containerAppProbeTcpSocketPtrType) ToContainerAppProbeTcpSocketPtrOutputWithContext(ctx context.Context) ContainerAppProbeTcpSocketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppProbeTcpSocketPtrOutput)
}

type ContainerAppProbeTcpSocketOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeTcpSocketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppProbeTcpSocket)(nil)).Elem()
}

func (o ContainerAppProbeTcpSocketOutput) ToContainerAppProbeTcpSocketOutput() ContainerAppProbeTcpSocketOutput {
	return o
}

func (o ContainerAppProbeTcpSocketOutput) ToContainerAppProbeTcpSocketOutputWithContext(ctx context.Context) ContainerAppProbeTcpSocketOutput {
	return o
}

func (o ContainerAppProbeTcpSocketOutput) ToContainerAppProbeTcpSocketPtrOutput() ContainerAppProbeTcpSocketPtrOutput {
	return o.ToContainerAppProbeTcpSocketPtrOutputWithContext(context.Background())
}

func (o ContainerAppProbeTcpSocketOutput) ToContainerAppProbeTcpSocketPtrOutputWithContext(ctx context.Context) ContainerAppProbeTcpSocketPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerAppProbeTcpSocket) *ContainerAppProbeTcpSocket {
		return &v
	}).(ContainerAppProbeTcpSocketPtrOutput)
}

func (o ContainerAppProbeTcpSocketOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppProbeTcpSocket) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeTcpSocketOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerAppProbeTcpSocket) int { return v.Port }).(pulumi.IntOutput)
}

type ContainerAppProbeTcpSocketPtrOutput struct{ *pulumi.OutputState }

func (ContainerAppProbeTcpSocketPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppProbeTcpSocket)(nil)).Elem()
}

func (o ContainerAppProbeTcpSocketPtrOutput) ToContainerAppProbeTcpSocketPtrOutput() ContainerAppProbeTcpSocketPtrOutput {
	return o
}

func (o ContainerAppProbeTcpSocketPtrOutput) ToContainerAppProbeTcpSocketPtrOutputWithContext(ctx context.Context) ContainerAppProbeTcpSocketPtrOutput {
	return o
}

func (o ContainerAppProbeTcpSocketPtrOutput) Elem() ContainerAppProbeTcpSocketOutput {
	return o.ApplyT(func(v *ContainerAppProbeTcpSocket) ContainerAppProbeTcpSocket {
		if v != nil {
			return *v
		}
		var ret ContainerAppProbeTcpSocket
		return ret
	}).(ContainerAppProbeTcpSocketOutput)
}

func (o ContainerAppProbeTcpSocketPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeTcpSocket) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppProbeTcpSocketPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerAppProbeTcpSocket) *int {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.IntPtrOutput)
}

type ContainerAppSecretResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type ContainerAppSecretResponseOutput struct{ *pulumi.OutputState }

func (ContainerAppSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppSecretResponse)(nil)).Elem()
}

func (o ContainerAppSecretResponseOutput) ToContainerAppSecretResponseOutput() ContainerAppSecretResponseOutput {
	return o
}

func (o ContainerAppSecretResponseOutput) ToContainerAppSecretResponseOutputWithContext(ctx context.Context) ContainerAppSecretResponseOutput {
	return o
}

func (o ContainerAppSecretResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerAppSecretResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerAppSecretResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerAppSecretResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ContainerAppSecretResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerAppSecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerAppSecretResponse)(nil)).Elem()
}

func (o ContainerAppSecretResponseArrayOutput) ToContainerAppSecretResponseArrayOutput() ContainerAppSecretResponseArrayOutput {
	return o
}

func (o ContainerAppSecretResponseArrayOutput) ToContainerAppSecretResponseArrayOutputWithContext(ctx context.Context) ContainerAppSecretResponseArrayOutput {
	return o
}

func (o ContainerAppSecretResponseArrayOutput) Index(i pulumi.IntInput) ContainerAppSecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerAppSecretResponse {
		return vs[0].([]ContainerAppSecretResponse)[vs[1].(int)]
	}).(ContainerAppSecretResponseOutput)
}

type ContainerResources struct {
	Cpu    *float64 `pulumi:"cpu"`
	Memory *string  `pulumi:"memory"`
}





type ContainerResourcesInput interface {
	pulumi.Input

	ToContainerResourcesOutput() ContainerResourcesOutput
	ToContainerResourcesOutputWithContext(context.Context) ContainerResourcesOutput
}

type ContainerResourcesArgs struct {
	Cpu    pulumi.Float64PtrInput `pulumi:"cpu"`
	Memory pulumi.StringPtrInput  `pulumi:"memory"`
}

func (ContainerResourcesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResources)(nil)).Elem()
}

func (i ContainerResourcesArgs) ToContainerResourcesOutput() ContainerResourcesOutput {
	return i.ToContainerResourcesOutputWithContext(context.Background())
}

func (i ContainerResourcesArgs) ToContainerResourcesOutputWithContext(ctx context.Context) ContainerResourcesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourcesOutput)
}

func (i ContainerResourcesArgs) ToContainerResourcesPtrOutput() ContainerResourcesPtrOutput {
	return i.ToContainerResourcesPtrOutputWithContext(context.Background())
}

func (i ContainerResourcesArgs) ToContainerResourcesPtrOutputWithContext(ctx context.Context) ContainerResourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourcesOutput).ToContainerResourcesPtrOutputWithContext(ctx)
}









type ContainerResourcesPtrInput interface {
	pulumi.Input

	ToContainerResourcesPtrOutput() ContainerResourcesPtrOutput
	ToContainerResourcesPtrOutputWithContext(context.Context) ContainerResourcesPtrOutput
}

type containerResourcesPtrType ContainerResourcesArgs

func ContainerResourcesPtr(v *ContainerResourcesArgs) ContainerResourcesPtrInput {
	return (*containerResourcesPtrType)(v)
}

func (*containerResourcesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerResources)(nil)).Elem()
}

func (i *containerResourcesPtrType) ToContainerResourcesPtrOutput() ContainerResourcesPtrOutput {
	return i.ToContainerResourcesPtrOutputWithContext(context.Background())
}

func (i *containerResourcesPtrType) ToContainerResourcesPtrOutputWithContext(ctx context.Context) ContainerResourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourcesPtrOutput)
}

type ContainerResourcesOutput struct{ *pulumi.OutputState }

func (ContainerResourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResources)(nil)).Elem()
}

func (o ContainerResourcesOutput) ToContainerResourcesOutput() ContainerResourcesOutput {
	return o
}

func (o ContainerResourcesOutput) ToContainerResourcesOutputWithContext(ctx context.Context) ContainerResourcesOutput {
	return o
}

func (o ContainerResourcesOutput) ToContainerResourcesPtrOutput() ContainerResourcesPtrOutput {
	return o.ToContainerResourcesPtrOutputWithContext(context.Background())
}

func (o ContainerResourcesOutput) ToContainerResourcesPtrOutputWithContext(ctx context.Context) ContainerResourcesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerResources) *ContainerResources {
		return &v
	}).(ContainerResourcesPtrOutput)
}

func (o ContainerResourcesOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerResources) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ContainerResourcesOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerResources) *string { return v.Memory }).(pulumi.StringPtrOutput)
}

type ContainerResourcesPtrOutput struct{ *pulumi.OutputState }

func (ContainerResourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerResources)(nil)).Elem()
}

func (o ContainerResourcesPtrOutput) ToContainerResourcesPtrOutput() ContainerResourcesPtrOutput {
	return o
}

func (o ContainerResourcesPtrOutput) ToContainerResourcesPtrOutputWithContext(ctx context.Context) ContainerResourcesPtrOutput {
	return o
}

func (o ContainerResourcesPtrOutput) Elem() ContainerResourcesOutput {
	return o.ApplyT(func(v *ContainerResources) ContainerResources {
		if v != nil {
			return *v
		}
		var ret ContainerResources
		return ret
	}).(ContainerResourcesOutput)
}

func (o ContainerResourcesPtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContainerResources) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ContainerResourcesPtrOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerResources) *string {
		if v == nil {
			return nil
		}
		return v.Memory
	}).(pulumi.StringPtrOutput)
}

type ContainerResourcesResponse struct {
	Cpu              *float64 `pulumi:"cpu"`
	EphemeralStorage string   `pulumi:"ephemeralStorage"`
	Memory           *string  `pulumi:"memory"`
}

type ContainerResourcesResponseOutput struct{ *pulumi.OutputState }

func (ContainerResourcesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResourcesResponse)(nil)).Elem()
}

func (o ContainerResourcesResponseOutput) ToContainerResourcesResponseOutput() ContainerResourcesResponseOutput {
	return o
}

func (o ContainerResourcesResponseOutput) ToContainerResourcesResponseOutputWithContext(ctx context.Context) ContainerResourcesResponseOutput {
	return o
}

func (o ContainerResourcesResponseOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerResourcesResponse) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ContainerResourcesResponseOutput) EphemeralStorage() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerResourcesResponse) string { return v.EphemeralStorage }).(pulumi.StringOutput)
}

func (o ContainerResourcesResponseOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerResourcesResponse) *string { return v.Memory }).(pulumi.StringPtrOutput)
}

type ContainerResourcesResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerResourcesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerResourcesResponse)(nil)).Elem()
}

func (o ContainerResourcesResponsePtrOutput) ToContainerResourcesResponsePtrOutput() ContainerResourcesResponsePtrOutput {
	return o
}

func (o ContainerResourcesResponsePtrOutput) ToContainerResourcesResponsePtrOutputWithContext(ctx context.Context) ContainerResourcesResponsePtrOutput {
	return o
}

func (o ContainerResourcesResponsePtrOutput) Elem() ContainerResourcesResponseOutput {
	return o.ApplyT(func(v *ContainerResourcesResponse) ContainerResourcesResponse {
		if v != nil {
			return *v
		}
		var ret ContainerResourcesResponse
		return ret
	}).(ContainerResourcesResponseOutput)
}

func (o ContainerResourcesResponsePtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContainerResourcesResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ContainerResourcesResponsePtrOutput) EphemeralStorage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerResourcesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EphemeralStorage
	}).(pulumi.StringPtrOutput)
}

func (o ContainerResourcesResponsePtrOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerResourcesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Memory
	}).(pulumi.StringPtrOutput)
}

type ContainerResponse struct {
	Args         []string                    `pulumi:"args"`
	Command      []string                    `pulumi:"command"`
	Env          []EnvironmentVarResponse    `pulumi:"env"`
	Image        *string                     `pulumi:"image"`
	Name         *string                     `pulumi:"name"`
	Probes       []ContainerAppProbeResponse `pulumi:"probes"`
	Resources    *ContainerResourcesResponse `pulumi:"resources"`
	VolumeMounts []VolumeMountResponse       `pulumi:"volumeMounts"`
}

type ContainerResponseOutput struct{ *pulumi.OutputState }

func (ContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResponse)(nil)).Elem()
}

func (o ContainerResponseOutput) ToContainerResponseOutput() ContainerResponseOutput {
	return o
}

func (o ContainerResponseOutput) ToContainerResponseOutputWithContext(ctx context.Context) ContainerResponseOutput {
	return o
}

func (o ContainerResponseOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []string { return v.Args }).(pulumi.StringArrayOutput)
}

func (o ContainerResponseOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []string { return v.Command }).(pulumi.StringArrayOutput)
}

func (o ContainerResponseOutput) Env() EnvironmentVarResponseArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []EnvironmentVarResponse { return v.Env }).(EnvironmentVarResponseArrayOutput)
}

func (o ContainerResponseOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerResponse) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o ContainerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ContainerResponseOutput) Probes() ContainerAppProbeResponseArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []ContainerAppProbeResponse { return v.Probes }).(ContainerAppProbeResponseArrayOutput)
}

func (o ContainerResponseOutput) Resources() ContainerResourcesResponsePtrOutput {
	return o.ApplyT(func(v ContainerResponse) *ContainerResourcesResponse { return v.Resources }).(ContainerResourcesResponsePtrOutput)
}

func (o ContainerResponseOutput) VolumeMounts() VolumeMountResponseArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []VolumeMountResponse { return v.VolumeMounts }).(VolumeMountResponseArrayOutput)
}

type ContainerResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerResponse)(nil)).Elem()
}

func (o ContainerResponseArrayOutput) ToContainerResponseArrayOutput() ContainerResponseArrayOutput {
	return o
}

func (o ContainerResponseArrayOutput) ToContainerResponseArrayOutputWithContext(ctx context.Context) ContainerResponseArrayOutput {
	return o
}

func (o ContainerResponseArrayOutput) Index(i pulumi.IntInput) ContainerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerResponse {
		return vs[0].([]ContainerResponse)[vs[1].(int)]
	}).(ContainerResponseOutput)
}

type CookieExpiration struct {
	Convention       *CookieExpirationConvention `pulumi:"convention"`
	TimeToExpiration *string                     `pulumi:"timeToExpiration"`
}





type CookieExpirationInput interface {
	pulumi.Input

	ToCookieExpirationOutput() CookieExpirationOutput
	ToCookieExpirationOutputWithContext(context.Context) CookieExpirationOutput
}

type CookieExpirationArgs struct {
	Convention       CookieExpirationConventionPtrInput `pulumi:"convention"`
	TimeToExpiration pulumi.StringPtrInput              `pulumi:"timeToExpiration"`
}

func (CookieExpirationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CookieExpiration)(nil)).Elem()
}

func (i CookieExpirationArgs) ToCookieExpirationOutput() CookieExpirationOutput {
	return i.ToCookieExpirationOutputWithContext(context.Background())
}

func (i CookieExpirationArgs) ToCookieExpirationOutputWithContext(ctx context.Context) CookieExpirationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CookieExpirationOutput)
}

func (i CookieExpirationArgs) ToCookieExpirationPtrOutput() CookieExpirationPtrOutput {
	return i.ToCookieExpirationPtrOutputWithContext(context.Background())
}

func (i CookieExpirationArgs) ToCookieExpirationPtrOutputWithContext(ctx context.Context) CookieExpirationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CookieExpirationOutput).ToCookieExpirationPtrOutputWithContext(ctx)
}









type CookieExpirationPtrInput interface {
	pulumi.Input

	ToCookieExpirationPtrOutput() CookieExpirationPtrOutput
	ToCookieExpirationPtrOutputWithContext(context.Context) CookieExpirationPtrOutput
}

type cookieExpirationPtrType CookieExpirationArgs

func CookieExpirationPtr(v *CookieExpirationArgs) CookieExpirationPtrInput {
	return (*cookieExpirationPtrType)(v)
}

func (*cookieExpirationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CookieExpiration)(nil)).Elem()
}

func (i *cookieExpirationPtrType) ToCookieExpirationPtrOutput() CookieExpirationPtrOutput {
	return i.ToCookieExpirationPtrOutputWithContext(context.Background())
}

func (i *cookieExpirationPtrType) ToCookieExpirationPtrOutputWithContext(ctx context.Context) CookieExpirationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CookieExpirationPtrOutput)
}

type CookieExpirationOutput struct{ *pulumi.OutputState }

func (CookieExpirationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CookieExpiration)(nil)).Elem()
}

func (o CookieExpirationOutput) ToCookieExpirationOutput() CookieExpirationOutput {
	return o
}

func (o CookieExpirationOutput) ToCookieExpirationOutputWithContext(ctx context.Context) CookieExpirationOutput {
	return o
}

func (o CookieExpirationOutput) ToCookieExpirationPtrOutput() CookieExpirationPtrOutput {
	return o.ToCookieExpirationPtrOutputWithContext(context.Background())
}

func (o CookieExpirationOutput) ToCookieExpirationPtrOutputWithContext(ctx context.Context) CookieExpirationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CookieExpiration) *CookieExpiration {
		return &v
	}).(CookieExpirationPtrOutput)
}

func (o CookieExpirationOutput) Convention() CookieExpirationConventionPtrOutput {
	return o.ApplyT(func(v CookieExpiration) *CookieExpirationConvention { return v.Convention }).(CookieExpirationConventionPtrOutput)
}

func (o CookieExpirationOutput) TimeToExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CookieExpiration) *string { return v.TimeToExpiration }).(pulumi.StringPtrOutput)
}

type CookieExpirationPtrOutput struct{ *pulumi.OutputState }

func (CookieExpirationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CookieExpiration)(nil)).Elem()
}

func (o CookieExpirationPtrOutput) ToCookieExpirationPtrOutput() CookieExpirationPtrOutput {
	return o
}

func (o CookieExpirationPtrOutput) ToCookieExpirationPtrOutputWithContext(ctx context.Context) CookieExpirationPtrOutput {
	return o
}

func (o CookieExpirationPtrOutput) Elem() CookieExpirationOutput {
	return o.ApplyT(func(v *CookieExpiration) CookieExpiration {
		if v != nil {
			return *v
		}
		var ret CookieExpiration
		return ret
	}).(CookieExpirationOutput)
}

func (o CookieExpirationPtrOutput) Convention() CookieExpirationConventionPtrOutput {
	return o.ApplyT(func(v *CookieExpiration) *CookieExpirationConvention {
		if v == nil {
			return nil
		}
		return v.Convention
	}).(CookieExpirationConventionPtrOutput)
}

func (o CookieExpirationPtrOutput) TimeToExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CookieExpiration) *string {
		if v == nil {
			return nil
		}
		return v.TimeToExpiration
	}).(pulumi.StringPtrOutput)
}

type CookieExpirationResponse struct {
	Convention       *string `pulumi:"convention"`
	TimeToExpiration *string `pulumi:"timeToExpiration"`
}

type CookieExpirationResponseOutput struct{ *pulumi.OutputState }

func (CookieExpirationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CookieExpirationResponse)(nil)).Elem()
}

func (o CookieExpirationResponseOutput) ToCookieExpirationResponseOutput() CookieExpirationResponseOutput {
	return o
}

func (o CookieExpirationResponseOutput) ToCookieExpirationResponseOutputWithContext(ctx context.Context) CookieExpirationResponseOutput {
	return o
}

func (o CookieExpirationResponseOutput) Convention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CookieExpirationResponse) *string { return v.Convention }).(pulumi.StringPtrOutput)
}

func (o CookieExpirationResponseOutput) TimeToExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CookieExpirationResponse) *string { return v.TimeToExpiration }).(pulumi.StringPtrOutput)
}

type CookieExpirationResponsePtrOutput struct{ *pulumi.OutputState }

func (CookieExpirationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CookieExpirationResponse)(nil)).Elem()
}

func (o CookieExpirationResponsePtrOutput) ToCookieExpirationResponsePtrOutput() CookieExpirationResponsePtrOutput {
	return o
}

func (o CookieExpirationResponsePtrOutput) ToCookieExpirationResponsePtrOutputWithContext(ctx context.Context) CookieExpirationResponsePtrOutput {
	return o
}

func (o CookieExpirationResponsePtrOutput) Elem() CookieExpirationResponseOutput {
	return o.ApplyT(func(v *CookieExpirationResponse) CookieExpirationResponse {
		if v != nil {
			return *v
		}
		var ret CookieExpirationResponse
		return ret
	}).(CookieExpirationResponseOutput)
}

func (o CookieExpirationResponsePtrOutput) Convention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CookieExpirationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Convention
	}).(pulumi.StringPtrOutput)
}

func (o CookieExpirationResponsePtrOutput) TimeToExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CookieExpirationResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeToExpiration
	}).(pulumi.StringPtrOutput)
}

type CustomDomain struct {
	BindingType   *string `pulumi:"bindingType"`
	CertificateId string  `pulumi:"certificateId"`
	Name          string  `pulumi:"name"`
}





type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(context.Context) CustomDomainOutput
}

type CustomDomainArgs struct {
	BindingType   pulumi.StringPtrInput `pulumi:"bindingType"`
	CertificateId pulumi.StringInput    `pulumi:"certificateId"`
	Name          pulumi.StringInput    `pulumi:"name"`
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (i CustomDomainArgs) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}





type CustomDomainArrayInput interface {
	pulumi.Input

	ToCustomDomainArrayOutput() CustomDomainArrayOutput
	ToCustomDomainArrayOutputWithContext(context.Context) CustomDomainArrayOutput
}

type CustomDomainArray []CustomDomainInput

func (CustomDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomDomain)(nil)).Elem()
}

func (i CustomDomainArray) ToCustomDomainArrayOutput() CustomDomainArrayOutput {
	return i.ToCustomDomainArrayOutputWithContext(context.Background())
}

func (i CustomDomainArray) ToCustomDomainArrayOutputWithContext(ctx context.Context) CustomDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainArrayOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) BindingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomain) *string { return v.BindingType }).(pulumi.StringPtrOutput)
}

func (o CustomDomainOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomain) string { return v.CertificateId }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomain) string { return v.Name }).(pulumi.StringOutput)
}

type CustomDomainArrayOutput struct{ *pulumi.OutputState }

func (CustomDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomDomain)(nil)).Elem()
}

func (o CustomDomainArrayOutput) ToCustomDomainArrayOutput() CustomDomainArrayOutput {
	return o
}

func (o CustomDomainArrayOutput) ToCustomDomainArrayOutputWithContext(ctx context.Context) CustomDomainArrayOutput {
	return o
}

func (o CustomDomainArrayOutput) Index(i pulumi.IntInput) CustomDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomDomain {
		return vs[0].([]CustomDomain)[vs[1].(int)]
	}).(CustomDomainOutput)
}

type CustomDomainResponse struct {
	BindingType   *string `pulumi:"bindingType"`
	CertificateId string  `pulumi:"certificateId"`
	Name          string  `pulumi:"name"`
}

type CustomDomainResponseOutput struct{ *pulumi.OutputState }

func (CustomDomainResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutput() CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutputWithContext(ctx context.Context) CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) BindingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *string { return v.BindingType }).(pulumi.StringPtrOutput)
}

func (o CustomDomainResponseOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomainResponse) string { return v.CertificateId }).(pulumi.StringOutput)
}

func (o CustomDomainResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomainResponse) string { return v.Name }).(pulumi.StringOutput)
}

type CustomDomainResponseArrayOutput struct{ *pulumi.OutputState }

func (CustomDomainResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponseArrayOutput) ToCustomDomainResponseArrayOutput() CustomDomainResponseArrayOutput {
	return o
}

func (o CustomDomainResponseArrayOutput) ToCustomDomainResponseArrayOutputWithContext(ctx context.Context) CustomDomainResponseArrayOutput {
	return o
}

func (o CustomDomainResponseArrayOutput) Index(i pulumi.IntInput) CustomDomainResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomDomainResponse {
		return vs[0].([]CustomDomainResponse)[vs[1].(int)]
	}).(CustomDomainResponseOutput)
}

type CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfo struct {
	Code    string                                        `pulumi:"code"`
	Details []CustomHostnameAnalysisResultResponseDetails `pulumi:"details"`
	Message string                                        `pulumi:"message"`
	Target  string                                        `pulumi:"target"`
}

type CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput struct{ *pulumi.OutputState }

func (CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfo)(nil)).Elem()
}

func (o CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput) ToCustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput() CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput {
	return o
}

func (o CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput) ToCustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutputWithContext(ctx context.Context) CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput {
	return o
}

func (o CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfo) string { return v.Code }).(pulumi.StringOutput)
}

func (o CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput) Details() CustomHostnameAnalysisResultResponseDetailsArrayOutput {
	return o.ApplyT(func(v CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfo) []CustomHostnameAnalysisResultResponseDetails {
		return v.Details
	}).(CustomHostnameAnalysisResultResponseDetailsArrayOutput)
}

func (o CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfo) string {
		return v.Message
	}).(pulumi.StringOutput)
}

func (o CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfo) string {
		return v.Target
	}).(pulumi.StringOutput)
}

type CustomHostnameAnalysisResultResponseDetails struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
	Target  string `pulumi:"target"`
}

type CustomHostnameAnalysisResultResponseDetailsOutput struct{ *pulumi.OutputState }

func (CustomHostnameAnalysisResultResponseDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHostnameAnalysisResultResponseDetails)(nil)).Elem()
}

func (o CustomHostnameAnalysisResultResponseDetailsOutput) ToCustomHostnameAnalysisResultResponseDetailsOutput() CustomHostnameAnalysisResultResponseDetailsOutput {
	return o
}

func (o CustomHostnameAnalysisResultResponseDetailsOutput) ToCustomHostnameAnalysisResultResponseDetailsOutputWithContext(ctx context.Context) CustomHostnameAnalysisResultResponseDetailsOutput {
	return o
}

func (o CustomHostnameAnalysisResultResponseDetailsOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHostnameAnalysisResultResponseDetails) string { return v.Code }).(pulumi.StringOutput)
}

func (o CustomHostnameAnalysisResultResponseDetailsOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHostnameAnalysisResultResponseDetails) string { return v.Message }).(pulumi.StringOutput)
}

func (o CustomHostnameAnalysisResultResponseDetailsOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHostnameAnalysisResultResponseDetails) string { return v.Target }).(pulumi.StringOutput)
}

type CustomHostnameAnalysisResultResponseDetailsArrayOutput struct{ *pulumi.OutputState }

func (CustomHostnameAnalysisResultResponseDetailsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomHostnameAnalysisResultResponseDetails)(nil)).Elem()
}

func (o CustomHostnameAnalysisResultResponseDetailsArrayOutput) ToCustomHostnameAnalysisResultResponseDetailsArrayOutput() CustomHostnameAnalysisResultResponseDetailsArrayOutput {
	return o
}

func (o CustomHostnameAnalysisResultResponseDetailsArrayOutput) ToCustomHostnameAnalysisResultResponseDetailsArrayOutputWithContext(ctx context.Context) CustomHostnameAnalysisResultResponseDetailsArrayOutput {
	return o
}

func (o CustomHostnameAnalysisResultResponseDetailsArrayOutput) Index(i pulumi.IntInput) CustomHostnameAnalysisResultResponseDetailsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomHostnameAnalysisResultResponseDetails {
		return vs[0].([]CustomHostnameAnalysisResultResponseDetails)[vs[1].(int)]
	}).(CustomHostnameAnalysisResultResponseDetailsOutput)
}

type CustomOpenIdConnectProvider struct {
	Enabled      *bool                      `pulumi:"enabled"`
	Login        *OpenIdConnectLogin        `pulumi:"login"`
	Registration *OpenIdConnectRegistration `pulumi:"registration"`
}





type CustomOpenIdConnectProviderInput interface {
	pulumi.Input

	ToCustomOpenIdConnectProviderOutput() CustomOpenIdConnectProviderOutput
	ToCustomOpenIdConnectProviderOutputWithContext(context.Context) CustomOpenIdConnectProviderOutput
}

type CustomOpenIdConnectProviderArgs struct {
	Enabled      pulumi.BoolPtrInput               `pulumi:"enabled"`
	Login        OpenIdConnectLoginPtrInput        `pulumi:"login"`
	Registration OpenIdConnectRegistrationPtrInput `pulumi:"registration"`
}

func (CustomOpenIdConnectProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomOpenIdConnectProvider)(nil)).Elem()
}

func (i CustomOpenIdConnectProviderArgs) ToCustomOpenIdConnectProviderOutput() CustomOpenIdConnectProviderOutput {
	return i.ToCustomOpenIdConnectProviderOutputWithContext(context.Background())
}

func (i CustomOpenIdConnectProviderArgs) ToCustomOpenIdConnectProviderOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomOpenIdConnectProviderOutput)
}





type CustomOpenIdConnectProviderMapInput interface {
	pulumi.Input

	ToCustomOpenIdConnectProviderMapOutput() CustomOpenIdConnectProviderMapOutput
	ToCustomOpenIdConnectProviderMapOutputWithContext(context.Context) CustomOpenIdConnectProviderMapOutput
}

type CustomOpenIdConnectProviderMap map[string]CustomOpenIdConnectProviderInput

func (CustomOpenIdConnectProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomOpenIdConnectProvider)(nil)).Elem()
}

func (i CustomOpenIdConnectProviderMap) ToCustomOpenIdConnectProviderMapOutput() CustomOpenIdConnectProviderMapOutput {
	return i.ToCustomOpenIdConnectProviderMapOutputWithContext(context.Background())
}

func (i CustomOpenIdConnectProviderMap) ToCustomOpenIdConnectProviderMapOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomOpenIdConnectProviderMapOutput)
}

type CustomOpenIdConnectProviderOutput struct{ *pulumi.OutputState }

func (CustomOpenIdConnectProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomOpenIdConnectProvider)(nil)).Elem()
}

func (o CustomOpenIdConnectProviderOutput) ToCustomOpenIdConnectProviderOutput() CustomOpenIdConnectProviderOutput {
	return o
}

func (o CustomOpenIdConnectProviderOutput) ToCustomOpenIdConnectProviderOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderOutput {
	return o
}

func (o CustomOpenIdConnectProviderOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProvider) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o CustomOpenIdConnectProviderOutput) Login() OpenIdConnectLoginPtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProvider) *OpenIdConnectLogin { return v.Login }).(OpenIdConnectLoginPtrOutput)
}

func (o CustomOpenIdConnectProviderOutput) Registration() OpenIdConnectRegistrationPtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProvider) *OpenIdConnectRegistration { return v.Registration }).(OpenIdConnectRegistrationPtrOutput)
}

type CustomOpenIdConnectProviderMapOutput struct{ *pulumi.OutputState }

func (CustomOpenIdConnectProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomOpenIdConnectProvider)(nil)).Elem()
}

func (o CustomOpenIdConnectProviderMapOutput) ToCustomOpenIdConnectProviderMapOutput() CustomOpenIdConnectProviderMapOutput {
	return o
}

func (o CustomOpenIdConnectProviderMapOutput) ToCustomOpenIdConnectProviderMapOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderMapOutput {
	return o
}

func (o CustomOpenIdConnectProviderMapOutput) MapIndex(k pulumi.StringInput) CustomOpenIdConnectProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomOpenIdConnectProvider {
		return vs[0].(map[string]CustomOpenIdConnectProvider)[vs[1].(string)]
	}).(CustomOpenIdConnectProviderOutput)
}

type CustomOpenIdConnectProviderResponse struct {
	Enabled      *bool                              `pulumi:"enabled"`
	Login        *OpenIdConnectLoginResponse        `pulumi:"login"`
	Registration *OpenIdConnectRegistrationResponse `pulumi:"registration"`
}

type CustomOpenIdConnectProviderResponseOutput struct{ *pulumi.OutputState }

func (CustomOpenIdConnectProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomOpenIdConnectProviderResponse)(nil)).Elem()
}

func (o CustomOpenIdConnectProviderResponseOutput) ToCustomOpenIdConnectProviderResponseOutput() CustomOpenIdConnectProviderResponseOutput {
	return o
}

func (o CustomOpenIdConnectProviderResponseOutput) ToCustomOpenIdConnectProviderResponseOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderResponseOutput {
	return o
}

func (o CustomOpenIdConnectProviderResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProviderResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o CustomOpenIdConnectProviderResponseOutput) Login() OpenIdConnectLoginResponsePtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProviderResponse) *OpenIdConnectLoginResponse { return v.Login }).(OpenIdConnectLoginResponsePtrOutput)
}

func (o CustomOpenIdConnectProviderResponseOutput) Registration() OpenIdConnectRegistrationResponsePtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProviderResponse) *OpenIdConnectRegistrationResponse { return v.Registration }).(OpenIdConnectRegistrationResponsePtrOutput)
}

type CustomOpenIdConnectProviderResponseMapOutput struct{ *pulumi.OutputState }

func (CustomOpenIdConnectProviderResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomOpenIdConnectProviderResponse)(nil)).Elem()
}

func (o CustomOpenIdConnectProviderResponseMapOutput) ToCustomOpenIdConnectProviderResponseMapOutput() CustomOpenIdConnectProviderResponseMapOutput {
	return o
}

func (o CustomOpenIdConnectProviderResponseMapOutput) ToCustomOpenIdConnectProviderResponseMapOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderResponseMapOutput {
	return o
}

func (o CustomOpenIdConnectProviderResponseMapOutput) MapIndex(k pulumi.StringInput) CustomOpenIdConnectProviderResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomOpenIdConnectProviderResponse {
		return vs[0].(map[string]CustomOpenIdConnectProviderResponse)[vs[1].(string)]
	}).(CustomOpenIdConnectProviderResponseOutput)
}

type CustomScaleRule struct {
	Auth     []ScaleRuleAuth   `pulumi:"auth"`
	Metadata map[string]string `pulumi:"metadata"`
	Type     *string           `pulumi:"type"`
}





type CustomScaleRuleInput interface {
	pulumi.Input

	ToCustomScaleRuleOutput() CustomScaleRuleOutput
	ToCustomScaleRuleOutputWithContext(context.Context) CustomScaleRuleOutput
}

type CustomScaleRuleArgs struct {
	Auth     ScaleRuleAuthArrayInput `pulumi:"auth"`
	Metadata pulumi.StringMapInput   `pulumi:"metadata"`
	Type     pulumi.StringPtrInput   `pulumi:"type"`
}

func (CustomScaleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomScaleRule)(nil)).Elem()
}

func (i CustomScaleRuleArgs) ToCustomScaleRuleOutput() CustomScaleRuleOutput {
	return i.ToCustomScaleRuleOutputWithContext(context.Background())
}

func (i CustomScaleRuleArgs) ToCustomScaleRuleOutputWithContext(ctx context.Context) CustomScaleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomScaleRuleOutput)
}

func (i CustomScaleRuleArgs) ToCustomScaleRulePtrOutput() CustomScaleRulePtrOutput {
	return i.ToCustomScaleRulePtrOutputWithContext(context.Background())
}

func (i CustomScaleRuleArgs) ToCustomScaleRulePtrOutputWithContext(ctx context.Context) CustomScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomScaleRuleOutput).ToCustomScaleRulePtrOutputWithContext(ctx)
}









type CustomScaleRulePtrInput interface {
	pulumi.Input

	ToCustomScaleRulePtrOutput() CustomScaleRulePtrOutput
	ToCustomScaleRulePtrOutputWithContext(context.Context) CustomScaleRulePtrOutput
}

type customScaleRulePtrType CustomScaleRuleArgs

func CustomScaleRulePtr(v *CustomScaleRuleArgs) CustomScaleRulePtrInput {
	return (*customScaleRulePtrType)(v)
}

func (*customScaleRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomScaleRule)(nil)).Elem()
}

func (i *customScaleRulePtrType) ToCustomScaleRulePtrOutput() CustomScaleRulePtrOutput {
	return i.ToCustomScaleRulePtrOutputWithContext(context.Background())
}

func (i *customScaleRulePtrType) ToCustomScaleRulePtrOutputWithContext(ctx context.Context) CustomScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomScaleRulePtrOutput)
}

type CustomScaleRuleOutput struct{ *pulumi.OutputState }

func (CustomScaleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomScaleRule)(nil)).Elem()
}

func (o CustomScaleRuleOutput) ToCustomScaleRuleOutput() CustomScaleRuleOutput {
	return o
}

func (o CustomScaleRuleOutput) ToCustomScaleRuleOutputWithContext(ctx context.Context) CustomScaleRuleOutput {
	return o
}

func (o CustomScaleRuleOutput) ToCustomScaleRulePtrOutput() CustomScaleRulePtrOutput {
	return o.ToCustomScaleRulePtrOutputWithContext(context.Background())
}

func (o CustomScaleRuleOutput) ToCustomScaleRulePtrOutputWithContext(ctx context.Context) CustomScaleRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomScaleRule) *CustomScaleRule {
		return &v
	}).(CustomScaleRulePtrOutput)
}

func (o CustomScaleRuleOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v CustomScaleRule) []ScaleRuleAuth { return v.Auth }).(ScaleRuleAuthArrayOutput)
}

func (o CustomScaleRuleOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v CustomScaleRule) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o CustomScaleRuleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomScaleRule) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CustomScaleRulePtrOutput struct{ *pulumi.OutputState }

func (CustomScaleRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomScaleRule)(nil)).Elem()
}

func (o CustomScaleRulePtrOutput) ToCustomScaleRulePtrOutput() CustomScaleRulePtrOutput {
	return o
}

func (o CustomScaleRulePtrOutput) ToCustomScaleRulePtrOutputWithContext(ctx context.Context) CustomScaleRulePtrOutput {
	return o
}

func (o CustomScaleRulePtrOutput) Elem() CustomScaleRuleOutput {
	return o.ApplyT(func(v *CustomScaleRule) CustomScaleRule {
		if v != nil {
			return *v
		}
		var ret CustomScaleRule
		return ret
	}).(CustomScaleRuleOutput)
}

func (o CustomScaleRulePtrOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v *CustomScaleRule) []ScaleRuleAuth {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthArrayOutput)
}

func (o CustomScaleRulePtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomScaleRule) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

func (o CustomScaleRulePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomScaleRule) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type CustomScaleRuleResponse struct {
	Auth     []ScaleRuleAuthResponse `pulumi:"auth"`
	Metadata map[string]string       `pulumi:"metadata"`
	Type     *string                 `pulumi:"type"`
}

type CustomScaleRuleResponseOutput struct{ *pulumi.OutputState }

func (CustomScaleRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomScaleRuleResponse)(nil)).Elem()
}

func (o CustomScaleRuleResponseOutput) ToCustomScaleRuleResponseOutput() CustomScaleRuleResponseOutput {
	return o
}

func (o CustomScaleRuleResponseOutput) ToCustomScaleRuleResponseOutputWithContext(ctx context.Context) CustomScaleRuleResponseOutput {
	return o
}

func (o CustomScaleRuleResponseOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v CustomScaleRuleResponse) []ScaleRuleAuthResponse { return v.Auth }).(ScaleRuleAuthResponseArrayOutput)
}

func (o CustomScaleRuleResponseOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v CustomScaleRuleResponse) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o CustomScaleRuleResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomScaleRuleResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CustomScaleRuleResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomScaleRuleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomScaleRuleResponse)(nil)).Elem()
}

func (o CustomScaleRuleResponsePtrOutput) ToCustomScaleRuleResponsePtrOutput() CustomScaleRuleResponsePtrOutput {
	return o
}

func (o CustomScaleRuleResponsePtrOutput) ToCustomScaleRuleResponsePtrOutputWithContext(ctx context.Context) CustomScaleRuleResponsePtrOutput {
	return o
}

func (o CustomScaleRuleResponsePtrOutput) Elem() CustomScaleRuleResponseOutput {
	return o.ApplyT(func(v *CustomScaleRuleResponse) CustomScaleRuleResponse {
		if v != nil {
			return *v
		}
		var ret CustomScaleRuleResponse
		return ret
	}).(CustomScaleRuleResponseOutput)
}

func (o CustomScaleRuleResponsePtrOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v *CustomScaleRuleResponse) []ScaleRuleAuthResponse {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthResponseArrayOutput)
}

func (o CustomScaleRuleResponsePtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomScaleRuleResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

func (o CustomScaleRuleResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomScaleRuleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type Dapr struct {
	AppId       *string `pulumi:"appId"`
	AppPort     *int    `pulumi:"appPort"`
	AppProtocol *string `pulumi:"appProtocol"`
	Enabled     *bool   `pulumi:"enabled"`
}


func (val *Dapr) Defaults() *Dapr {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AppProtocol) {
		appProtocol_ := "http"
		tmp.AppProtocol = &appProtocol_
	}
	if isZero(tmp.Enabled) {
		enabled_ := false
		tmp.Enabled = &enabled_
	}
	return &tmp
}





type DaprInput interface {
	pulumi.Input

	ToDaprOutput() DaprOutput
	ToDaprOutputWithContext(context.Context) DaprOutput
}

type DaprArgs struct {
	AppId       pulumi.StringPtrInput `pulumi:"appId"`
	AppPort     pulumi.IntPtrInput    `pulumi:"appPort"`
	AppProtocol pulumi.StringPtrInput `pulumi:"appProtocol"`
	Enabled     pulumi.BoolPtrInput   `pulumi:"enabled"`
}


func (val *DaprArgs) Defaults() *DaprArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AppProtocol) {
		tmp.AppProtocol = pulumi.StringPtr("http")
	}
	if isZero(tmp.Enabled) {
		tmp.Enabled = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (DaprArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Dapr)(nil)).Elem()
}

func (i DaprArgs) ToDaprOutput() DaprOutput {
	return i.ToDaprOutputWithContext(context.Background())
}

func (i DaprArgs) ToDaprOutputWithContext(ctx context.Context) DaprOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprOutput)
}

func (i DaprArgs) ToDaprPtrOutput() DaprPtrOutput {
	return i.ToDaprPtrOutputWithContext(context.Background())
}

func (i DaprArgs) ToDaprPtrOutputWithContext(ctx context.Context) DaprPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprOutput).ToDaprPtrOutputWithContext(ctx)
}









type DaprPtrInput interface {
	pulumi.Input

	ToDaprPtrOutput() DaprPtrOutput
	ToDaprPtrOutputWithContext(context.Context) DaprPtrOutput
}

type daprPtrType DaprArgs

func DaprPtr(v *DaprArgs) DaprPtrInput {
	return (*daprPtrType)(v)
}

func (*daprPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Dapr)(nil)).Elem()
}

func (i *daprPtrType) ToDaprPtrOutput() DaprPtrOutput {
	return i.ToDaprPtrOutputWithContext(context.Background())
}

func (i *daprPtrType) ToDaprPtrOutputWithContext(ctx context.Context) DaprPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprPtrOutput)
}

type DaprOutput struct{ *pulumi.OutputState }

func (DaprOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dapr)(nil)).Elem()
}

func (o DaprOutput) ToDaprOutput() DaprOutput {
	return o
}

func (o DaprOutput) ToDaprOutputWithContext(ctx context.Context) DaprOutput {
	return o
}

func (o DaprOutput) ToDaprPtrOutput() DaprPtrOutput {
	return o.ToDaprPtrOutputWithContext(context.Background())
}

func (o DaprOutput) ToDaprPtrOutputWithContext(ctx context.Context) DaprPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Dapr) *Dapr {
		return &v
	}).(DaprPtrOutput)
}

func (o DaprOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Dapr) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o DaprOutput) AppPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Dapr) *int { return v.AppPort }).(pulumi.IntPtrOutput)
}

func (o DaprOutput) AppProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Dapr) *string { return v.AppProtocol }).(pulumi.StringPtrOutput)
}

func (o DaprOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Dapr) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type DaprPtrOutput struct{ *pulumi.OutputState }

func (DaprPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dapr)(nil)).Elem()
}

func (o DaprPtrOutput) ToDaprPtrOutput() DaprPtrOutput {
	return o
}

func (o DaprPtrOutput) ToDaprPtrOutputWithContext(ctx context.Context) DaprPtrOutput {
	return o
}

func (o DaprPtrOutput) Elem() DaprOutput {
	return o.ApplyT(func(v *Dapr) Dapr {
		if v != nil {
			return *v
		}
		var ret Dapr
		return ret
	}).(DaprOutput)
}

func (o DaprPtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dapr) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o DaprPtrOutput) AppPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Dapr) *int {
		if v == nil {
			return nil
		}
		return v.AppPort
	}).(pulumi.IntPtrOutput)
}

func (o DaprPtrOutput) AppProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dapr) *string {
		if v == nil {
			return nil
		}
		return v.AppProtocol
	}).(pulumi.StringPtrOutput)
}

func (o DaprPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Dapr) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type DaprMetadata struct {
	Name      *string `pulumi:"name"`
	SecretRef *string `pulumi:"secretRef"`
	Value     *string `pulumi:"value"`
}





type DaprMetadataInput interface {
	pulumi.Input

	ToDaprMetadataOutput() DaprMetadataOutput
	ToDaprMetadataOutputWithContext(context.Context) DaprMetadataOutput
}

type DaprMetadataArgs struct {
	Name      pulumi.StringPtrInput `pulumi:"name"`
	SecretRef pulumi.StringPtrInput `pulumi:"secretRef"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (DaprMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprMetadata)(nil)).Elem()
}

func (i DaprMetadataArgs) ToDaprMetadataOutput() DaprMetadataOutput {
	return i.ToDaprMetadataOutputWithContext(context.Background())
}

func (i DaprMetadataArgs) ToDaprMetadataOutputWithContext(ctx context.Context) DaprMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprMetadataOutput)
}





type DaprMetadataArrayInput interface {
	pulumi.Input

	ToDaprMetadataArrayOutput() DaprMetadataArrayOutput
	ToDaprMetadataArrayOutputWithContext(context.Context) DaprMetadataArrayOutput
}

type DaprMetadataArray []DaprMetadataInput

func (DaprMetadataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaprMetadata)(nil)).Elem()
}

func (i DaprMetadataArray) ToDaprMetadataArrayOutput() DaprMetadataArrayOutput {
	return i.ToDaprMetadataArrayOutputWithContext(context.Background())
}

func (i DaprMetadataArray) ToDaprMetadataArrayOutputWithContext(ctx context.Context) DaprMetadataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprMetadataArrayOutput)
}

type DaprMetadataOutput struct{ *pulumi.OutputState }

func (DaprMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprMetadata)(nil)).Elem()
}

func (o DaprMetadataOutput) ToDaprMetadataOutput() DaprMetadataOutput {
	return o
}

func (o DaprMetadataOutput) ToDaprMetadataOutputWithContext(ctx context.Context) DaprMetadataOutput {
	return o
}

func (o DaprMetadataOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadata) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DaprMetadataOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadata) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o DaprMetadataOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadata) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DaprMetadataArrayOutput struct{ *pulumi.OutputState }

func (DaprMetadataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaprMetadata)(nil)).Elem()
}

func (o DaprMetadataArrayOutput) ToDaprMetadataArrayOutput() DaprMetadataArrayOutput {
	return o
}

func (o DaprMetadataArrayOutput) ToDaprMetadataArrayOutputWithContext(ctx context.Context) DaprMetadataArrayOutput {
	return o
}

func (o DaprMetadataArrayOutput) Index(i pulumi.IntInput) DaprMetadataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DaprMetadata {
		return vs[0].([]DaprMetadata)[vs[1].(int)]
	}).(DaprMetadataOutput)
}

type DaprMetadataResponse struct {
	Name      *string `pulumi:"name"`
	SecretRef *string `pulumi:"secretRef"`
	Value     *string `pulumi:"value"`
}

type DaprMetadataResponseOutput struct{ *pulumi.OutputState }

func (DaprMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprMetadataResponse)(nil)).Elem()
}

func (o DaprMetadataResponseOutput) ToDaprMetadataResponseOutput() DaprMetadataResponseOutput {
	return o
}

func (o DaprMetadataResponseOutput) ToDaprMetadataResponseOutputWithContext(ctx context.Context) DaprMetadataResponseOutput {
	return o
}

func (o DaprMetadataResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadataResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DaprMetadataResponseOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadataResponse) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o DaprMetadataResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadataResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DaprMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (DaprMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaprMetadataResponse)(nil)).Elem()
}

func (o DaprMetadataResponseArrayOutput) ToDaprMetadataResponseArrayOutput() DaprMetadataResponseArrayOutput {
	return o
}

func (o DaprMetadataResponseArrayOutput) ToDaprMetadataResponseArrayOutputWithContext(ctx context.Context) DaprMetadataResponseArrayOutput {
	return o
}

func (o DaprMetadataResponseArrayOutput) Index(i pulumi.IntInput) DaprMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DaprMetadataResponse {
		return vs[0].([]DaprMetadataResponse)[vs[1].(int)]
	}).(DaprMetadataResponseOutput)
}

type DaprResponse struct {
	AppId       *string `pulumi:"appId"`
	AppPort     *int    `pulumi:"appPort"`
	AppProtocol *string `pulumi:"appProtocol"`
	Enabled     *bool   `pulumi:"enabled"`
}


func (val *DaprResponse) Defaults() *DaprResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AppProtocol) {
		appProtocol_ := "http"
		tmp.AppProtocol = &appProtocol_
	}
	if isZero(tmp.Enabled) {
		enabled_ := false
		tmp.Enabled = &enabled_
	}
	return &tmp
}

type DaprResponseOutput struct{ *pulumi.OutputState }

func (DaprResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprResponse)(nil)).Elem()
}

func (o DaprResponseOutput) ToDaprResponseOutput() DaprResponseOutput {
	return o
}

func (o DaprResponseOutput) ToDaprResponseOutputWithContext(ctx context.Context) DaprResponseOutput {
	return o
}

func (o DaprResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o DaprResponseOutput) AppPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DaprResponse) *int { return v.AppPort }).(pulumi.IntPtrOutput)
}

func (o DaprResponseOutput) AppProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprResponse) *string { return v.AppProtocol }).(pulumi.StringPtrOutput)
}

func (o DaprResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DaprResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type DaprResponsePtrOutput struct{ *pulumi.OutputState }

func (DaprResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DaprResponse)(nil)).Elem()
}

func (o DaprResponsePtrOutput) ToDaprResponsePtrOutput() DaprResponsePtrOutput {
	return o
}

func (o DaprResponsePtrOutput) ToDaprResponsePtrOutputWithContext(ctx context.Context) DaprResponsePtrOutput {
	return o
}

func (o DaprResponsePtrOutput) Elem() DaprResponseOutput {
	return o.ApplyT(func(v *DaprResponse) DaprResponse {
		if v != nil {
			return *v
		}
		var ret DaprResponse
		return ret
	}).(DaprResponseOutput)
}

func (o DaprResponsePtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaprResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o DaprResponsePtrOutput) AppPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DaprResponse) *int {
		if v == nil {
			return nil
		}
		return v.AppPort
	}).(pulumi.IntPtrOutput)
}

func (o DaprResponsePtrOutput) AppProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaprResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppProtocol
	}).(pulumi.StringPtrOutput)
}

func (o DaprResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DaprResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type DaprSecretResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type DaprSecretResponseOutput struct{ *pulumi.OutputState }

func (DaprSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprSecretResponse)(nil)).Elem()
}

func (o DaprSecretResponseOutput) ToDaprSecretResponseOutput() DaprSecretResponseOutput {
	return o
}

func (o DaprSecretResponseOutput) ToDaprSecretResponseOutputWithContext(ctx context.Context) DaprSecretResponseOutput {
	return o
}

func (o DaprSecretResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DaprSecretResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DaprSecretResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DaprSecretResponse) string { return v.Value }).(pulumi.StringOutput)
}

type DaprSecretResponseArrayOutput struct{ *pulumi.OutputState }

func (DaprSecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaprSecretResponse)(nil)).Elem()
}

func (o DaprSecretResponseArrayOutput) ToDaprSecretResponseArrayOutput() DaprSecretResponseArrayOutput {
	return o
}

func (o DaprSecretResponseArrayOutput) ToDaprSecretResponseArrayOutputWithContext(ctx context.Context) DaprSecretResponseArrayOutput {
	return o
}

func (o DaprSecretResponseArrayOutput) Index(i pulumi.IntInput) DaprSecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DaprSecretResponse {
		return vs[0].([]DaprSecretResponse)[vs[1].(int)]
	}).(DaprSecretResponseOutput)
}

type DefaultAuthorizationPolicy struct {
	AllowedApplications []string           `pulumi:"allowedApplications"`
	AllowedPrincipals   *AllowedPrincipals `pulumi:"allowedPrincipals"`
}





type DefaultAuthorizationPolicyInput interface {
	pulumi.Input

	ToDefaultAuthorizationPolicyOutput() DefaultAuthorizationPolicyOutput
	ToDefaultAuthorizationPolicyOutputWithContext(context.Context) DefaultAuthorizationPolicyOutput
}

type DefaultAuthorizationPolicyArgs struct {
	AllowedApplications pulumi.StringArrayInput   `pulumi:"allowedApplications"`
	AllowedPrincipals   AllowedPrincipalsPtrInput `pulumi:"allowedPrincipals"`
}

func (DefaultAuthorizationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAuthorizationPolicy)(nil)).Elem()
}

func (i DefaultAuthorizationPolicyArgs) ToDefaultAuthorizationPolicyOutput() DefaultAuthorizationPolicyOutput {
	return i.ToDefaultAuthorizationPolicyOutputWithContext(context.Background())
}

func (i DefaultAuthorizationPolicyArgs) ToDefaultAuthorizationPolicyOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultAuthorizationPolicyOutput)
}

func (i DefaultAuthorizationPolicyArgs) ToDefaultAuthorizationPolicyPtrOutput() DefaultAuthorizationPolicyPtrOutput {
	return i.ToDefaultAuthorizationPolicyPtrOutputWithContext(context.Background())
}

func (i DefaultAuthorizationPolicyArgs) ToDefaultAuthorizationPolicyPtrOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultAuthorizationPolicyOutput).ToDefaultAuthorizationPolicyPtrOutputWithContext(ctx)
}









type DefaultAuthorizationPolicyPtrInput interface {
	pulumi.Input

	ToDefaultAuthorizationPolicyPtrOutput() DefaultAuthorizationPolicyPtrOutput
	ToDefaultAuthorizationPolicyPtrOutputWithContext(context.Context) DefaultAuthorizationPolicyPtrOutput
}

type defaultAuthorizationPolicyPtrType DefaultAuthorizationPolicyArgs

func DefaultAuthorizationPolicyPtr(v *DefaultAuthorizationPolicyArgs) DefaultAuthorizationPolicyPtrInput {
	return (*defaultAuthorizationPolicyPtrType)(v)
}

func (*defaultAuthorizationPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAuthorizationPolicy)(nil)).Elem()
}

func (i *defaultAuthorizationPolicyPtrType) ToDefaultAuthorizationPolicyPtrOutput() DefaultAuthorizationPolicyPtrOutput {
	return i.ToDefaultAuthorizationPolicyPtrOutputWithContext(context.Background())
}

func (i *defaultAuthorizationPolicyPtrType) ToDefaultAuthorizationPolicyPtrOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultAuthorizationPolicyPtrOutput)
}

type DefaultAuthorizationPolicyOutput struct{ *pulumi.OutputState }

func (DefaultAuthorizationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAuthorizationPolicy)(nil)).Elem()
}

func (o DefaultAuthorizationPolicyOutput) ToDefaultAuthorizationPolicyOutput() DefaultAuthorizationPolicyOutput {
	return o
}

func (o DefaultAuthorizationPolicyOutput) ToDefaultAuthorizationPolicyOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyOutput {
	return o
}

func (o DefaultAuthorizationPolicyOutput) ToDefaultAuthorizationPolicyPtrOutput() DefaultAuthorizationPolicyPtrOutput {
	return o.ToDefaultAuthorizationPolicyPtrOutputWithContext(context.Background())
}

func (o DefaultAuthorizationPolicyOutput) ToDefaultAuthorizationPolicyPtrOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultAuthorizationPolicy) *DefaultAuthorizationPolicy {
		return &v
	}).(DefaultAuthorizationPolicyPtrOutput)
}

func (o DefaultAuthorizationPolicyOutput) AllowedApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DefaultAuthorizationPolicy) []string { return v.AllowedApplications }).(pulumi.StringArrayOutput)
}

func (o DefaultAuthorizationPolicyOutput) AllowedPrincipals() AllowedPrincipalsPtrOutput {
	return o.ApplyT(func(v DefaultAuthorizationPolicy) *AllowedPrincipals { return v.AllowedPrincipals }).(AllowedPrincipalsPtrOutput)
}

type DefaultAuthorizationPolicyPtrOutput struct{ *pulumi.OutputState }

func (DefaultAuthorizationPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAuthorizationPolicy)(nil)).Elem()
}

func (o DefaultAuthorizationPolicyPtrOutput) ToDefaultAuthorizationPolicyPtrOutput() DefaultAuthorizationPolicyPtrOutput {
	return o
}

func (o DefaultAuthorizationPolicyPtrOutput) ToDefaultAuthorizationPolicyPtrOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyPtrOutput {
	return o
}

func (o DefaultAuthorizationPolicyPtrOutput) Elem() DefaultAuthorizationPolicyOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicy) DefaultAuthorizationPolicy {
		if v != nil {
			return *v
		}
		var ret DefaultAuthorizationPolicy
		return ret
	}).(DefaultAuthorizationPolicyOutput)
}

func (o DefaultAuthorizationPolicyPtrOutput) AllowedApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicy) []string {
		if v == nil {
			return nil
		}
		return v.AllowedApplications
	}).(pulumi.StringArrayOutput)
}

func (o DefaultAuthorizationPolicyPtrOutput) AllowedPrincipals() AllowedPrincipalsPtrOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicy) *AllowedPrincipals {
		if v == nil {
			return nil
		}
		return v.AllowedPrincipals
	}).(AllowedPrincipalsPtrOutput)
}

type DefaultAuthorizationPolicyResponse struct {
	AllowedApplications []string                   `pulumi:"allowedApplications"`
	AllowedPrincipals   *AllowedPrincipalsResponse `pulumi:"allowedPrincipals"`
}

type DefaultAuthorizationPolicyResponseOutput struct{ *pulumi.OutputState }

func (DefaultAuthorizationPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAuthorizationPolicyResponse)(nil)).Elem()
}

func (o DefaultAuthorizationPolicyResponseOutput) ToDefaultAuthorizationPolicyResponseOutput() DefaultAuthorizationPolicyResponseOutput {
	return o
}

func (o DefaultAuthorizationPolicyResponseOutput) ToDefaultAuthorizationPolicyResponseOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyResponseOutput {
	return o
}

func (o DefaultAuthorizationPolicyResponseOutput) AllowedApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DefaultAuthorizationPolicyResponse) []string { return v.AllowedApplications }).(pulumi.StringArrayOutput)
}

func (o DefaultAuthorizationPolicyResponseOutput) AllowedPrincipals() AllowedPrincipalsResponsePtrOutput {
	return o.ApplyT(func(v DefaultAuthorizationPolicyResponse) *AllowedPrincipalsResponse { return v.AllowedPrincipals }).(AllowedPrincipalsResponsePtrOutput)
}

type DefaultAuthorizationPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (DefaultAuthorizationPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAuthorizationPolicyResponse)(nil)).Elem()
}

func (o DefaultAuthorizationPolicyResponsePtrOutput) ToDefaultAuthorizationPolicyResponsePtrOutput() DefaultAuthorizationPolicyResponsePtrOutput {
	return o
}

func (o DefaultAuthorizationPolicyResponsePtrOutput) ToDefaultAuthorizationPolicyResponsePtrOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyResponsePtrOutput {
	return o
}

func (o DefaultAuthorizationPolicyResponsePtrOutput) Elem() DefaultAuthorizationPolicyResponseOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicyResponse) DefaultAuthorizationPolicyResponse {
		if v != nil {
			return *v
		}
		var ret DefaultAuthorizationPolicyResponse
		return ret
	}).(DefaultAuthorizationPolicyResponseOutput)
}

func (o DefaultAuthorizationPolicyResponsePtrOutput) AllowedApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicyResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedApplications
	}).(pulumi.StringArrayOutput)
}

func (o DefaultAuthorizationPolicyResponsePtrOutput) AllowedPrincipals() AllowedPrincipalsResponsePtrOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicyResponse) *AllowedPrincipalsResponse {
		if v == nil {
			return nil
		}
		return v.AllowedPrincipals
	}).(AllowedPrincipalsResponsePtrOutput)
}

type EnvironmentVar struct {
	Name      *string `pulumi:"name"`
	SecretRef *string `pulumi:"secretRef"`
	Value     *string `pulumi:"value"`
}





type EnvironmentVarInput interface {
	pulumi.Input

	ToEnvironmentVarOutput() EnvironmentVarOutput
	ToEnvironmentVarOutputWithContext(context.Context) EnvironmentVarOutput
}

type EnvironmentVarArgs struct {
	Name      pulumi.StringPtrInput `pulumi:"name"`
	SecretRef pulumi.StringPtrInput `pulumi:"secretRef"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (EnvironmentVarArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVar)(nil)).Elem()
}

func (i EnvironmentVarArgs) ToEnvironmentVarOutput() EnvironmentVarOutput {
	return i.ToEnvironmentVarOutputWithContext(context.Background())
}

func (i EnvironmentVarArgs) ToEnvironmentVarOutputWithContext(ctx context.Context) EnvironmentVarOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVarOutput)
}





type EnvironmentVarArrayInput interface {
	pulumi.Input

	ToEnvironmentVarArrayOutput() EnvironmentVarArrayOutput
	ToEnvironmentVarArrayOutputWithContext(context.Context) EnvironmentVarArrayOutput
}

type EnvironmentVarArray []EnvironmentVarInput

func (EnvironmentVarArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVar)(nil)).Elem()
}

func (i EnvironmentVarArray) ToEnvironmentVarArrayOutput() EnvironmentVarArrayOutput {
	return i.ToEnvironmentVarArrayOutputWithContext(context.Background())
}

func (i EnvironmentVarArray) ToEnvironmentVarArrayOutputWithContext(ctx context.Context) EnvironmentVarArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVarArrayOutput)
}

type EnvironmentVarOutput struct{ *pulumi.OutputState }

func (EnvironmentVarOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVar)(nil)).Elem()
}

func (o EnvironmentVarOutput) ToEnvironmentVarOutput() EnvironmentVarOutput {
	return o
}

func (o EnvironmentVarOutput) ToEnvironmentVarOutputWithContext(ctx context.Context) EnvironmentVarOutput {
	return o
}

func (o EnvironmentVarOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVar) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVarOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVar) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVarOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVar) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentVarArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVarArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVar)(nil)).Elem()
}

func (o EnvironmentVarArrayOutput) ToEnvironmentVarArrayOutput() EnvironmentVarArrayOutput {
	return o
}

func (o EnvironmentVarArrayOutput) ToEnvironmentVarArrayOutputWithContext(ctx context.Context) EnvironmentVarArrayOutput {
	return o
}

func (o EnvironmentVarArrayOutput) Index(i pulumi.IntInput) EnvironmentVarOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentVar {
		return vs[0].([]EnvironmentVar)[vs[1].(int)]
	}).(EnvironmentVarOutput)
}

type EnvironmentVarResponse struct {
	Name      *string `pulumi:"name"`
	SecretRef *string `pulumi:"secretRef"`
	Value     *string `pulumi:"value"`
}

type EnvironmentVarResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentVarResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVarResponse)(nil)).Elem()
}

func (o EnvironmentVarResponseOutput) ToEnvironmentVarResponseOutput() EnvironmentVarResponseOutput {
	return o
}

func (o EnvironmentVarResponseOutput) ToEnvironmentVarResponseOutputWithContext(ctx context.Context) EnvironmentVarResponseOutput {
	return o
}

func (o EnvironmentVarResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVarResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVarResponseOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVarResponse) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVarResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVarResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentVarResponseArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVarResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVarResponse)(nil)).Elem()
}

func (o EnvironmentVarResponseArrayOutput) ToEnvironmentVarResponseArrayOutput() EnvironmentVarResponseArrayOutput {
	return o
}

func (o EnvironmentVarResponseArrayOutput) ToEnvironmentVarResponseArrayOutputWithContext(ctx context.Context) EnvironmentVarResponseArrayOutput {
	return o
}

func (o EnvironmentVarResponseArrayOutput) Index(i pulumi.IntInput) EnvironmentVarResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentVarResponse {
		return vs[0].([]EnvironmentVarResponse)[vs[1].(int)]
	}).(EnvironmentVarResponseOutput)
}

type Facebook struct {
	Enabled         *bool            `pulumi:"enabled"`
	GraphApiVersion *string          `pulumi:"graphApiVersion"`
	Login           *LoginScopes     `pulumi:"login"`
	Registration    *AppRegistration `pulumi:"registration"`
}





type FacebookInput interface {
	pulumi.Input

	ToFacebookOutput() FacebookOutput
	ToFacebookOutputWithContext(context.Context) FacebookOutput
}

type FacebookArgs struct {
	Enabled         pulumi.BoolPtrInput     `pulumi:"enabled"`
	GraphApiVersion pulumi.StringPtrInput   `pulumi:"graphApiVersion"`
	Login           LoginScopesPtrInput     `pulumi:"login"`
	Registration    AppRegistrationPtrInput `pulumi:"registration"`
}

func (FacebookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Facebook)(nil)).Elem()
}

func (i FacebookArgs) ToFacebookOutput() FacebookOutput {
	return i.ToFacebookOutputWithContext(context.Background())
}

func (i FacebookArgs) ToFacebookOutputWithContext(ctx context.Context) FacebookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookOutput)
}

func (i FacebookArgs) ToFacebookPtrOutput() FacebookPtrOutput {
	return i.ToFacebookPtrOutputWithContext(context.Background())
}

func (i FacebookArgs) ToFacebookPtrOutputWithContext(ctx context.Context) FacebookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookOutput).ToFacebookPtrOutputWithContext(ctx)
}









type FacebookPtrInput interface {
	pulumi.Input

	ToFacebookPtrOutput() FacebookPtrOutput
	ToFacebookPtrOutputWithContext(context.Context) FacebookPtrOutput
}

type facebookPtrType FacebookArgs

func FacebookPtr(v *FacebookArgs) FacebookPtrInput {
	return (*facebookPtrType)(v)
}

func (*facebookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Facebook)(nil)).Elem()
}

func (i *facebookPtrType) ToFacebookPtrOutput() FacebookPtrOutput {
	return i.ToFacebookPtrOutputWithContext(context.Background())
}

func (i *facebookPtrType) ToFacebookPtrOutputWithContext(ctx context.Context) FacebookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookPtrOutput)
}

type FacebookOutput struct{ *pulumi.OutputState }

func (FacebookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Facebook)(nil)).Elem()
}

func (o FacebookOutput) ToFacebookOutput() FacebookOutput {
	return o
}

func (o FacebookOutput) ToFacebookOutputWithContext(ctx context.Context) FacebookOutput {
	return o
}

func (o FacebookOutput) ToFacebookPtrOutput() FacebookPtrOutput {
	return o.ToFacebookPtrOutputWithContext(context.Background())
}

func (o FacebookOutput) ToFacebookPtrOutputWithContext(ctx context.Context) FacebookPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Facebook) *Facebook {
		return &v
	}).(FacebookPtrOutput)
}

func (o FacebookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Facebook) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o FacebookOutput) GraphApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Facebook) *string { return v.GraphApiVersion }).(pulumi.StringPtrOutput)
}

func (o FacebookOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v Facebook) *LoginScopes { return v.Login }).(LoginScopesPtrOutput)
}

func (o FacebookOutput) Registration() AppRegistrationPtrOutput {
	return o.ApplyT(func(v Facebook) *AppRegistration { return v.Registration }).(AppRegistrationPtrOutput)
}

type FacebookPtrOutput struct{ *pulumi.OutputState }

func (FacebookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Facebook)(nil)).Elem()
}

func (o FacebookPtrOutput) ToFacebookPtrOutput() FacebookPtrOutput {
	return o
}

func (o FacebookPtrOutput) ToFacebookPtrOutputWithContext(ctx context.Context) FacebookPtrOutput {
	return o
}

func (o FacebookPtrOutput) Elem() FacebookOutput {
	return o.ApplyT(func(v *Facebook) Facebook {
		if v != nil {
			return *v
		}
		var ret Facebook
		return ret
	}).(FacebookOutput)
}

func (o FacebookPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Facebook) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o FacebookPtrOutput) GraphApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Facebook) *string {
		if v == nil {
			return nil
		}
		return v.GraphApiVersion
	}).(pulumi.StringPtrOutput)
}

func (o FacebookPtrOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v *Facebook) *LoginScopes {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesPtrOutput)
}

func (o FacebookPtrOutput) Registration() AppRegistrationPtrOutput {
	return o.ApplyT(func(v *Facebook) *AppRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AppRegistrationPtrOutput)
}

type FacebookResponse struct {
	Enabled         *bool                    `pulumi:"enabled"`
	GraphApiVersion *string                  `pulumi:"graphApiVersion"`
	Login           *LoginScopesResponse     `pulumi:"login"`
	Registration    *AppRegistrationResponse `pulumi:"registration"`
}

type FacebookResponseOutput struct{ *pulumi.OutputState }

func (FacebookResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookResponse)(nil)).Elem()
}

func (o FacebookResponseOutput) ToFacebookResponseOutput() FacebookResponseOutput {
	return o
}

func (o FacebookResponseOutput) ToFacebookResponseOutputWithContext(ctx context.Context) FacebookResponseOutput {
	return o
}

func (o FacebookResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FacebookResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o FacebookResponseOutput) GraphApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FacebookResponse) *string { return v.GraphApiVersion }).(pulumi.StringPtrOutput)
}

func (o FacebookResponseOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v FacebookResponse) *LoginScopesResponse { return v.Login }).(LoginScopesResponsePtrOutput)
}

func (o FacebookResponseOutput) Registration() AppRegistrationResponsePtrOutput {
	return o.ApplyT(func(v FacebookResponse) *AppRegistrationResponse { return v.Registration }).(AppRegistrationResponsePtrOutput)
}

type FacebookResponsePtrOutput struct{ *pulumi.OutputState }

func (FacebookResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FacebookResponse)(nil)).Elem()
}

func (o FacebookResponsePtrOutput) ToFacebookResponsePtrOutput() FacebookResponsePtrOutput {
	return o
}

func (o FacebookResponsePtrOutput) ToFacebookResponsePtrOutputWithContext(ctx context.Context) FacebookResponsePtrOutput {
	return o
}

func (o FacebookResponsePtrOutput) Elem() FacebookResponseOutput {
	return o.ApplyT(func(v *FacebookResponse) FacebookResponse {
		if v != nil {
			return *v
		}
		var ret FacebookResponse
		return ret
	}).(FacebookResponseOutput)
}

func (o FacebookResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FacebookResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o FacebookResponsePtrOutput) GraphApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FacebookResponse) *string {
		if v == nil {
			return nil
		}
		return v.GraphApiVersion
	}).(pulumi.StringPtrOutput)
}

func (o FacebookResponsePtrOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v *FacebookResponse) *LoginScopesResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesResponsePtrOutput)
}

func (o FacebookResponsePtrOutput) Registration() AppRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *FacebookResponse) *AppRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AppRegistrationResponsePtrOutput)
}

type ForwardProxy struct {
	Convention            *ForwardProxyConvention `pulumi:"convention"`
	CustomHostHeaderName  *string                 `pulumi:"customHostHeaderName"`
	CustomProtoHeaderName *string                 `pulumi:"customProtoHeaderName"`
}





type ForwardProxyInput interface {
	pulumi.Input

	ToForwardProxyOutput() ForwardProxyOutput
	ToForwardProxyOutputWithContext(context.Context) ForwardProxyOutput
}

type ForwardProxyArgs struct {
	Convention            ForwardProxyConventionPtrInput `pulumi:"convention"`
	CustomHostHeaderName  pulumi.StringPtrInput          `pulumi:"customHostHeaderName"`
	CustomProtoHeaderName pulumi.StringPtrInput          `pulumi:"customProtoHeaderName"`
}

func (ForwardProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardProxy)(nil)).Elem()
}

func (i ForwardProxyArgs) ToForwardProxyOutput() ForwardProxyOutput {
	return i.ToForwardProxyOutputWithContext(context.Background())
}

func (i ForwardProxyArgs) ToForwardProxyOutputWithContext(ctx context.Context) ForwardProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardProxyOutput)
}

func (i ForwardProxyArgs) ToForwardProxyPtrOutput() ForwardProxyPtrOutput {
	return i.ToForwardProxyPtrOutputWithContext(context.Background())
}

func (i ForwardProxyArgs) ToForwardProxyPtrOutputWithContext(ctx context.Context) ForwardProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardProxyOutput).ToForwardProxyPtrOutputWithContext(ctx)
}









type ForwardProxyPtrInput interface {
	pulumi.Input

	ToForwardProxyPtrOutput() ForwardProxyPtrOutput
	ToForwardProxyPtrOutputWithContext(context.Context) ForwardProxyPtrOutput
}

type forwardProxyPtrType ForwardProxyArgs

func ForwardProxyPtr(v *ForwardProxyArgs) ForwardProxyPtrInput {
	return (*forwardProxyPtrType)(v)
}

func (*forwardProxyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardProxy)(nil)).Elem()
}

func (i *forwardProxyPtrType) ToForwardProxyPtrOutput() ForwardProxyPtrOutput {
	return i.ToForwardProxyPtrOutputWithContext(context.Background())
}

func (i *forwardProxyPtrType) ToForwardProxyPtrOutputWithContext(ctx context.Context) ForwardProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardProxyPtrOutput)
}

type ForwardProxyOutput struct{ *pulumi.OutputState }

func (ForwardProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardProxy)(nil)).Elem()
}

func (o ForwardProxyOutput) ToForwardProxyOutput() ForwardProxyOutput {
	return o
}

func (o ForwardProxyOutput) ToForwardProxyOutputWithContext(ctx context.Context) ForwardProxyOutput {
	return o
}

func (o ForwardProxyOutput) ToForwardProxyPtrOutput() ForwardProxyPtrOutput {
	return o.ToForwardProxyPtrOutputWithContext(context.Background())
}

func (o ForwardProxyOutput) ToForwardProxyPtrOutputWithContext(ctx context.Context) ForwardProxyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ForwardProxy) *ForwardProxy {
		return &v
	}).(ForwardProxyPtrOutput)
}

func (o ForwardProxyOutput) Convention() ForwardProxyConventionPtrOutput {
	return o.ApplyT(func(v ForwardProxy) *ForwardProxyConvention { return v.Convention }).(ForwardProxyConventionPtrOutput)
}

func (o ForwardProxyOutput) CustomHostHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForwardProxy) *string { return v.CustomHostHeaderName }).(pulumi.StringPtrOutput)
}

func (o ForwardProxyOutput) CustomProtoHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForwardProxy) *string { return v.CustomProtoHeaderName }).(pulumi.StringPtrOutput)
}

type ForwardProxyPtrOutput struct{ *pulumi.OutputState }

func (ForwardProxyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardProxy)(nil)).Elem()
}

func (o ForwardProxyPtrOutput) ToForwardProxyPtrOutput() ForwardProxyPtrOutput {
	return o
}

func (o ForwardProxyPtrOutput) ToForwardProxyPtrOutputWithContext(ctx context.Context) ForwardProxyPtrOutput {
	return o
}

func (o ForwardProxyPtrOutput) Elem() ForwardProxyOutput {
	return o.ApplyT(func(v *ForwardProxy) ForwardProxy {
		if v != nil {
			return *v
		}
		var ret ForwardProxy
		return ret
	}).(ForwardProxyOutput)
}

func (o ForwardProxyPtrOutput) Convention() ForwardProxyConventionPtrOutput {
	return o.ApplyT(func(v *ForwardProxy) *ForwardProxyConvention {
		if v == nil {
			return nil
		}
		return v.Convention
	}).(ForwardProxyConventionPtrOutput)
}

func (o ForwardProxyPtrOutput) CustomHostHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardProxy) *string {
		if v == nil {
			return nil
		}
		return v.CustomHostHeaderName
	}).(pulumi.StringPtrOutput)
}

func (o ForwardProxyPtrOutput) CustomProtoHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardProxy) *string {
		if v == nil {
			return nil
		}
		return v.CustomProtoHeaderName
	}).(pulumi.StringPtrOutput)
}

type ForwardProxyResponse struct {
	Convention            *string `pulumi:"convention"`
	CustomHostHeaderName  *string `pulumi:"customHostHeaderName"`
	CustomProtoHeaderName *string `pulumi:"customProtoHeaderName"`
}

type ForwardProxyResponseOutput struct{ *pulumi.OutputState }

func (ForwardProxyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardProxyResponse)(nil)).Elem()
}

func (o ForwardProxyResponseOutput) ToForwardProxyResponseOutput() ForwardProxyResponseOutput {
	return o
}

func (o ForwardProxyResponseOutput) ToForwardProxyResponseOutputWithContext(ctx context.Context) ForwardProxyResponseOutput {
	return o
}

func (o ForwardProxyResponseOutput) Convention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForwardProxyResponse) *string { return v.Convention }).(pulumi.StringPtrOutput)
}

func (o ForwardProxyResponseOutput) CustomHostHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForwardProxyResponse) *string { return v.CustomHostHeaderName }).(pulumi.StringPtrOutput)
}

func (o ForwardProxyResponseOutput) CustomProtoHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForwardProxyResponse) *string { return v.CustomProtoHeaderName }).(pulumi.StringPtrOutput)
}

type ForwardProxyResponsePtrOutput struct{ *pulumi.OutputState }

func (ForwardProxyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardProxyResponse)(nil)).Elem()
}

func (o ForwardProxyResponsePtrOutput) ToForwardProxyResponsePtrOutput() ForwardProxyResponsePtrOutput {
	return o
}

func (o ForwardProxyResponsePtrOutput) ToForwardProxyResponsePtrOutputWithContext(ctx context.Context) ForwardProxyResponsePtrOutput {
	return o
}

func (o ForwardProxyResponsePtrOutput) Elem() ForwardProxyResponseOutput {
	return o.ApplyT(func(v *ForwardProxyResponse) ForwardProxyResponse {
		if v != nil {
			return *v
		}
		var ret ForwardProxyResponse
		return ret
	}).(ForwardProxyResponseOutput)
}

func (o ForwardProxyResponsePtrOutput) Convention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardProxyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Convention
	}).(pulumi.StringPtrOutput)
}

func (o ForwardProxyResponsePtrOutput) CustomHostHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardProxyResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomHostHeaderName
	}).(pulumi.StringPtrOutput)
}

func (o ForwardProxyResponsePtrOutput) CustomProtoHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardProxyResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomProtoHeaderName
	}).(pulumi.StringPtrOutput)
}

type GitHub struct {
	Enabled      *bool               `pulumi:"enabled"`
	Login        *LoginScopes        `pulumi:"login"`
	Registration *ClientRegistration `pulumi:"registration"`
}





type GitHubInput interface {
	pulumi.Input

	ToGitHubOutput() GitHubOutput
	ToGitHubOutputWithContext(context.Context) GitHubOutput
}

type GitHubArgs struct {
	Enabled      pulumi.BoolPtrInput        `pulumi:"enabled"`
	Login        LoginScopesPtrInput        `pulumi:"login"`
	Registration ClientRegistrationPtrInput `pulumi:"registration"`
}

func (GitHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHub)(nil)).Elem()
}

func (i GitHubArgs) ToGitHubOutput() GitHubOutput {
	return i.ToGitHubOutputWithContext(context.Background())
}

func (i GitHubArgs) ToGitHubOutputWithContext(ctx context.Context) GitHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubOutput)
}

func (i GitHubArgs) ToGitHubPtrOutput() GitHubPtrOutput {
	return i.ToGitHubPtrOutputWithContext(context.Background())
}

func (i GitHubArgs) ToGitHubPtrOutputWithContext(ctx context.Context) GitHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubOutput).ToGitHubPtrOutputWithContext(ctx)
}









type GitHubPtrInput interface {
	pulumi.Input

	ToGitHubPtrOutput() GitHubPtrOutput
	ToGitHubPtrOutputWithContext(context.Context) GitHubPtrOutput
}

type gitHubPtrType GitHubArgs

func GitHubPtr(v *GitHubArgs) GitHubPtrInput {
	return (*gitHubPtrType)(v)
}

func (*gitHubPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHub)(nil)).Elem()
}

func (i *gitHubPtrType) ToGitHubPtrOutput() GitHubPtrOutput {
	return i.ToGitHubPtrOutputWithContext(context.Background())
}

func (i *gitHubPtrType) ToGitHubPtrOutputWithContext(ctx context.Context) GitHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubPtrOutput)
}

type GitHubOutput struct{ *pulumi.OutputState }

func (GitHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHub)(nil)).Elem()
}

func (o GitHubOutput) ToGitHubOutput() GitHubOutput {
	return o
}

func (o GitHubOutput) ToGitHubOutputWithContext(ctx context.Context) GitHubOutput {
	return o
}

func (o GitHubOutput) ToGitHubPtrOutput() GitHubPtrOutput {
	return o.ToGitHubPtrOutputWithContext(context.Background())
}

func (o GitHubOutput) ToGitHubPtrOutputWithContext(ctx context.Context) GitHubPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHub) *GitHub {
		return &v
	}).(GitHubPtrOutput)
}

func (o GitHubOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHub) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o GitHubOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v GitHub) *LoginScopes { return v.Login }).(LoginScopesPtrOutput)
}

func (o GitHubOutput) Registration() ClientRegistrationPtrOutput {
	return o.ApplyT(func(v GitHub) *ClientRegistration { return v.Registration }).(ClientRegistrationPtrOutput)
}

type GitHubPtrOutput struct{ *pulumi.OutputState }

func (GitHubPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHub)(nil)).Elem()
}

func (o GitHubPtrOutput) ToGitHubPtrOutput() GitHubPtrOutput {
	return o
}

func (o GitHubPtrOutput) ToGitHubPtrOutputWithContext(ctx context.Context) GitHubPtrOutput {
	return o
}

func (o GitHubPtrOutput) Elem() GitHubOutput {
	return o.ApplyT(func(v *GitHub) GitHub {
		if v != nil {
			return *v
		}
		var ret GitHub
		return ret
	}).(GitHubOutput)
}

func (o GitHubPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHub) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o GitHubPtrOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v *GitHub) *LoginScopes {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesPtrOutput)
}

func (o GitHubPtrOutput) Registration() ClientRegistrationPtrOutput {
	return o.ApplyT(func(v *GitHub) *ClientRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(ClientRegistrationPtrOutput)
}

type GitHubResponse struct {
	Enabled      *bool                       `pulumi:"enabled"`
	Login        *LoginScopesResponse        `pulumi:"login"`
	Registration *ClientRegistrationResponse `pulumi:"registration"`
}

type GitHubResponseOutput struct{ *pulumi.OutputState }

func (GitHubResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubResponse)(nil)).Elem()
}

func (o GitHubResponseOutput) ToGitHubResponseOutput() GitHubResponseOutput {
	return o
}

func (o GitHubResponseOutput) ToGitHubResponseOutputWithContext(ctx context.Context) GitHubResponseOutput {
	return o
}

func (o GitHubResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHubResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o GitHubResponseOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v GitHubResponse) *LoginScopesResponse { return v.Login }).(LoginScopesResponsePtrOutput)
}

func (o GitHubResponseOutput) Registration() ClientRegistrationResponsePtrOutput {
	return o.ApplyT(func(v GitHubResponse) *ClientRegistrationResponse { return v.Registration }).(ClientRegistrationResponsePtrOutput)
}

type GitHubResponsePtrOutput struct{ *pulumi.OutputState }

func (GitHubResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubResponse)(nil)).Elem()
}

func (o GitHubResponsePtrOutput) ToGitHubResponsePtrOutput() GitHubResponsePtrOutput {
	return o
}

func (o GitHubResponsePtrOutput) ToGitHubResponsePtrOutputWithContext(ctx context.Context) GitHubResponsePtrOutput {
	return o
}

func (o GitHubResponsePtrOutput) Elem() GitHubResponseOutput {
	return o.ApplyT(func(v *GitHubResponse) GitHubResponse {
		if v != nil {
			return *v
		}
		var ret GitHubResponse
		return ret
	}).(GitHubResponseOutput)
}

func (o GitHubResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHubResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o GitHubResponsePtrOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v *GitHubResponse) *LoginScopesResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesResponsePtrOutput)
}

func (o GitHubResponsePtrOutput) Registration() ClientRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *GitHubResponse) *ClientRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(ClientRegistrationResponsePtrOutput)
}

type GithubActionConfiguration struct {
	AzureCredentials *AzureCredentials `pulumi:"azureCredentials"`
	ContextPath      *string           `pulumi:"contextPath"`
	Image            *string           `pulumi:"image"`
	Os               *string           `pulumi:"os"`
	PublishType      *string           `pulumi:"publishType"`
	RegistryInfo     *RegistryInfo     `pulumi:"registryInfo"`
	RuntimeStack     *string           `pulumi:"runtimeStack"`
	RuntimeVersion   *string           `pulumi:"runtimeVersion"`
}





type GithubActionConfigurationInput interface {
	pulumi.Input

	ToGithubActionConfigurationOutput() GithubActionConfigurationOutput
	ToGithubActionConfigurationOutputWithContext(context.Context) GithubActionConfigurationOutput
}

type GithubActionConfigurationArgs struct {
	AzureCredentials AzureCredentialsPtrInput `pulumi:"azureCredentials"`
	ContextPath      pulumi.StringPtrInput    `pulumi:"contextPath"`
	Image            pulumi.StringPtrInput    `pulumi:"image"`
	Os               pulumi.StringPtrInput    `pulumi:"os"`
	PublishType      pulumi.StringPtrInput    `pulumi:"publishType"`
	RegistryInfo     RegistryInfoPtrInput     `pulumi:"registryInfo"`
	RuntimeStack     pulumi.StringPtrInput    `pulumi:"runtimeStack"`
	RuntimeVersion   pulumi.StringPtrInput    `pulumi:"runtimeVersion"`
}

func (GithubActionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GithubActionConfiguration)(nil)).Elem()
}

func (i GithubActionConfigurationArgs) ToGithubActionConfigurationOutput() GithubActionConfigurationOutput {
	return i.ToGithubActionConfigurationOutputWithContext(context.Background())
}

func (i GithubActionConfigurationArgs) ToGithubActionConfigurationOutputWithContext(ctx context.Context) GithubActionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GithubActionConfigurationOutput)
}

func (i GithubActionConfigurationArgs) ToGithubActionConfigurationPtrOutput() GithubActionConfigurationPtrOutput {
	return i.ToGithubActionConfigurationPtrOutputWithContext(context.Background())
}

func (i GithubActionConfigurationArgs) ToGithubActionConfigurationPtrOutputWithContext(ctx context.Context) GithubActionConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GithubActionConfigurationOutput).ToGithubActionConfigurationPtrOutputWithContext(ctx)
}









type GithubActionConfigurationPtrInput interface {
	pulumi.Input

	ToGithubActionConfigurationPtrOutput() GithubActionConfigurationPtrOutput
	ToGithubActionConfigurationPtrOutputWithContext(context.Context) GithubActionConfigurationPtrOutput
}

type githubActionConfigurationPtrType GithubActionConfigurationArgs

func GithubActionConfigurationPtr(v *GithubActionConfigurationArgs) GithubActionConfigurationPtrInput {
	return (*githubActionConfigurationPtrType)(v)
}

func (*githubActionConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GithubActionConfiguration)(nil)).Elem()
}

func (i *githubActionConfigurationPtrType) ToGithubActionConfigurationPtrOutput() GithubActionConfigurationPtrOutput {
	return i.ToGithubActionConfigurationPtrOutputWithContext(context.Background())
}

func (i *githubActionConfigurationPtrType) ToGithubActionConfigurationPtrOutputWithContext(ctx context.Context) GithubActionConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GithubActionConfigurationPtrOutput)
}

type GithubActionConfigurationOutput struct{ *pulumi.OutputState }

func (GithubActionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GithubActionConfiguration)(nil)).Elem()
}

func (o GithubActionConfigurationOutput) ToGithubActionConfigurationOutput() GithubActionConfigurationOutput {
	return o
}

func (o GithubActionConfigurationOutput) ToGithubActionConfigurationOutputWithContext(ctx context.Context) GithubActionConfigurationOutput {
	return o
}

func (o GithubActionConfigurationOutput) ToGithubActionConfigurationPtrOutput() GithubActionConfigurationPtrOutput {
	return o.ToGithubActionConfigurationPtrOutputWithContext(context.Background())
}

func (o GithubActionConfigurationOutput) ToGithubActionConfigurationPtrOutputWithContext(ctx context.Context) GithubActionConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GithubActionConfiguration) *GithubActionConfiguration {
		return &v
	}).(GithubActionConfigurationPtrOutput)
}

func (o GithubActionConfigurationOutput) AzureCredentials() AzureCredentialsPtrOutput {
	return o.ApplyT(func(v GithubActionConfiguration) *AzureCredentials { return v.AzureCredentials }).(AzureCredentialsPtrOutput)
}

func (o GithubActionConfigurationOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfiguration) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfiguration) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfiguration) *string { return v.Os }).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationOutput) PublishType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfiguration) *string { return v.PublishType }).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationOutput) RegistryInfo() RegistryInfoPtrOutput {
	return o.ApplyT(func(v GithubActionConfiguration) *RegistryInfo { return v.RegistryInfo }).(RegistryInfoPtrOutput)
}

func (o GithubActionConfigurationOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfiguration) *string { return v.RuntimeStack }).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfiguration) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type GithubActionConfigurationPtrOutput struct{ *pulumi.OutputState }

func (GithubActionConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GithubActionConfiguration)(nil)).Elem()
}

func (o GithubActionConfigurationPtrOutput) ToGithubActionConfigurationPtrOutput() GithubActionConfigurationPtrOutput {
	return o
}

func (o GithubActionConfigurationPtrOutput) ToGithubActionConfigurationPtrOutputWithContext(ctx context.Context) GithubActionConfigurationPtrOutput {
	return o
}

func (o GithubActionConfigurationPtrOutput) Elem() GithubActionConfigurationOutput {
	return o.ApplyT(func(v *GithubActionConfiguration) GithubActionConfiguration {
		if v != nil {
			return *v
		}
		var ret GithubActionConfiguration
		return ret
	}).(GithubActionConfigurationOutput)
}

func (o GithubActionConfigurationPtrOutput) AzureCredentials() AzureCredentialsPtrOutput {
	return o.ApplyT(func(v *GithubActionConfiguration) *AzureCredentials {
		if v == nil {
			return nil
		}
		return v.AzureCredentials
	}).(AzureCredentialsPtrOutput)
}

func (o GithubActionConfigurationPtrOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ContextPath
	}).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationPtrOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Image
	}).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationPtrOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Os
	}).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationPtrOutput) PublishType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PublishType
	}).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationPtrOutput) RegistryInfo() RegistryInfoPtrOutput {
	return o.ApplyT(func(v *GithubActionConfiguration) *RegistryInfo {
		if v == nil {
			return nil
		}
		return v.RegistryInfo
	}).(RegistryInfoPtrOutput)
}

func (o GithubActionConfigurationPtrOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeStack
	}).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationPtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type GithubActionConfigurationResponse struct {
	AzureCredentials *AzureCredentialsResponse `pulumi:"azureCredentials"`
	ContextPath      *string                   `pulumi:"contextPath"`
	Image            *string                   `pulumi:"image"`
	Os               *string                   `pulumi:"os"`
	PublishType      *string                   `pulumi:"publishType"`
	RegistryInfo     *RegistryInfoResponse     `pulumi:"registryInfo"`
	RuntimeStack     *string                   `pulumi:"runtimeStack"`
	RuntimeVersion   *string                   `pulumi:"runtimeVersion"`
}

type GithubActionConfigurationResponseOutput struct{ *pulumi.OutputState }

func (GithubActionConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GithubActionConfigurationResponse)(nil)).Elem()
}

func (o GithubActionConfigurationResponseOutput) ToGithubActionConfigurationResponseOutput() GithubActionConfigurationResponseOutput {
	return o
}

func (o GithubActionConfigurationResponseOutput) ToGithubActionConfigurationResponseOutputWithContext(ctx context.Context) GithubActionConfigurationResponseOutput {
	return o
}

func (o GithubActionConfigurationResponseOutput) AzureCredentials() AzureCredentialsResponsePtrOutput {
	return o.ApplyT(func(v GithubActionConfigurationResponse) *AzureCredentialsResponse { return v.AzureCredentials }).(AzureCredentialsResponsePtrOutput)
}

func (o GithubActionConfigurationResponseOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfigurationResponse) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationResponseOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfigurationResponse) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationResponseOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfigurationResponse) *string { return v.Os }).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationResponseOutput) PublishType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfigurationResponse) *string { return v.PublishType }).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationResponseOutput) RegistryInfo() RegistryInfoResponsePtrOutput {
	return o.ApplyT(func(v GithubActionConfigurationResponse) *RegistryInfoResponse { return v.RegistryInfo }).(RegistryInfoResponsePtrOutput)
}

func (o GithubActionConfigurationResponseOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfigurationResponse) *string { return v.RuntimeStack }).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationResponseOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GithubActionConfigurationResponse) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type GithubActionConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (GithubActionConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GithubActionConfigurationResponse)(nil)).Elem()
}

func (o GithubActionConfigurationResponsePtrOutput) ToGithubActionConfigurationResponsePtrOutput() GithubActionConfigurationResponsePtrOutput {
	return o
}

func (o GithubActionConfigurationResponsePtrOutput) ToGithubActionConfigurationResponsePtrOutputWithContext(ctx context.Context) GithubActionConfigurationResponsePtrOutput {
	return o
}

func (o GithubActionConfigurationResponsePtrOutput) Elem() GithubActionConfigurationResponseOutput {
	return o.ApplyT(func(v *GithubActionConfigurationResponse) GithubActionConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret GithubActionConfigurationResponse
		return ret
	}).(GithubActionConfigurationResponseOutput)
}

func (o GithubActionConfigurationResponsePtrOutput) AzureCredentials() AzureCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *GithubActionConfigurationResponse) *AzureCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.AzureCredentials
	}).(AzureCredentialsResponsePtrOutput)
}

func (o GithubActionConfigurationResponsePtrOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContextPath
	}).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationResponsePtrOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Image
	}).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationResponsePtrOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Os
	}).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationResponsePtrOutput) PublishType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublishType
	}).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationResponsePtrOutput) RegistryInfo() RegistryInfoResponsePtrOutput {
	return o.ApplyT(func(v *GithubActionConfigurationResponse) *RegistryInfoResponse {
		if v == nil {
			return nil
		}
		return v.RegistryInfo
	}).(RegistryInfoResponsePtrOutput)
}

func (o GithubActionConfigurationResponsePtrOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeStack
	}).(pulumi.StringPtrOutput)
}

func (o GithubActionConfigurationResponsePtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubActionConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type GlobalValidation struct {
	ExcludedPaths               []string                       `pulumi:"excludedPaths"`
	RedirectToProvider          *string                        `pulumi:"redirectToProvider"`
	UnauthenticatedClientAction *UnauthenticatedClientActionV2 `pulumi:"unauthenticatedClientAction"`
}





type GlobalValidationInput interface {
	pulumi.Input

	ToGlobalValidationOutput() GlobalValidationOutput
	ToGlobalValidationOutputWithContext(context.Context) GlobalValidationOutput
}

type GlobalValidationArgs struct {
	ExcludedPaths               pulumi.StringArrayInput               `pulumi:"excludedPaths"`
	RedirectToProvider          pulumi.StringPtrInput                 `pulumi:"redirectToProvider"`
	UnauthenticatedClientAction UnauthenticatedClientActionV2PtrInput `pulumi:"unauthenticatedClientAction"`
}

func (GlobalValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalValidation)(nil)).Elem()
}

func (i GlobalValidationArgs) ToGlobalValidationOutput() GlobalValidationOutput {
	return i.ToGlobalValidationOutputWithContext(context.Background())
}

func (i GlobalValidationArgs) ToGlobalValidationOutputWithContext(ctx context.Context) GlobalValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalValidationOutput)
}

func (i GlobalValidationArgs) ToGlobalValidationPtrOutput() GlobalValidationPtrOutput {
	return i.ToGlobalValidationPtrOutputWithContext(context.Background())
}

func (i GlobalValidationArgs) ToGlobalValidationPtrOutputWithContext(ctx context.Context) GlobalValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalValidationOutput).ToGlobalValidationPtrOutputWithContext(ctx)
}









type GlobalValidationPtrInput interface {
	pulumi.Input

	ToGlobalValidationPtrOutput() GlobalValidationPtrOutput
	ToGlobalValidationPtrOutputWithContext(context.Context) GlobalValidationPtrOutput
}

type globalValidationPtrType GlobalValidationArgs

func GlobalValidationPtr(v *GlobalValidationArgs) GlobalValidationPtrInput {
	return (*globalValidationPtrType)(v)
}

func (*globalValidationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalValidation)(nil)).Elem()
}

func (i *globalValidationPtrType) ToGlobalValidationPtrOutput() GlobalValidationPtrOutput {
	return i.ToGlobalValidationPtrOutputWithContext(context.Background())
}

func (i *globalValidationPtrType) ToGlobalValidationPtrOutputWithContext(ctx context.Context) GlobalValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalValidationPtrOutput)
}

type GlobalValidationOutput struct{ *pulumi.OutputState }

func (GlobalValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalValidation)(nil)).Elem()
}

func (o GlobalValidationOutput) ToGlobalValidationOutput() GlobalValidationOutput {
	return o
}

func (o GlobalValidationOutput) ToGlobalValidationOutputWithContext(ctx context.Context) GlobalValidationOutput {
	return o
}

func (o GlobalValidationOutput) ToGlobalValidationPtrOutput() GlobalValidationPtrOutput {
	return o.ToGlobalValidationPtrOutputWithContext(context.Background())
}

func (o GlobalValidationOutput) ToGlobalValidationPtrOutputWithContext(ctx context.Context) GlobalValidationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GlobalValidation) *GlobalValidation {
		return &v
	}).(GlobalValidationPtrOutput)
}

func (o GlobalValidationOutput) ExcludedPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GlobalValidation) []string { return v.ExcludedPaths }).(pulumi.StringArrayOutput)
}

func (o GlobalValidationOutput) RedirectToProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GlobalValidation) *string { return v.RedirectToProvider }).(pulumi.StringPtrOutput)
}

func (o GlobalValidationOutput) UnauthenticatedClientAction() UnauthenticatedClientActionV2PtrOutput {
	return o.ApplyT(func(v GlobalValidation) *UnauthenticatedClientActionV2 { return v.UnauthenticatedClientAction }).(UnauthenticatedClientActionV2PtrOutput)
}

type GlobalValidationPtrOutput struct{ *pulumi.OutputState }

func (GlobalValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalValidation)(nil)).Elem()
}

func (o GlobalValidationPtrOutput) ToGlobalValidationPtrOutput() GlobalValidationPtrOutput {
	return o
}

func (o GlobalValidationPtrOutput) ToGlobalValidationPtrOutputWithContext(ctx context.Context) GlobalValidationPtrOutput {
	return o
}

func (o GlobalValidationPtrOutput) Elem() GlobalValidationOutput {
	return o.ApplyT(func(v *GlobalValidation) GlobalValidation {
		if v != nil {
			return *v
		}
		var ret GlobalValidation
		return ret
	}).(GlobalValidationOutput)
}

func (o GlobalValidationPtrOutput) ExcludedPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GlobalValidation) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedPaths
	}).(pulumi.StringArrayOutput)
}

func (o GlobalValidationPtrOutput) RedirectToProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalValidation) *string {
		if v == nil {
			return nil
		}
		return v.RedirectToProvider
	}).(pulumi.StringPtrOutput)
}

func (o GlobalValidationPtrOutput) UnauthenticatedClientAction() UnauthenticatedClientActionV2PtrOutput {
	return o.ApplyT(func(v *GlobalValidation) *UnauthenticatedClientActionV2 {
		if v == nil {
			return nil
		}
		return v.UnauthenticatedClientAction
	}).(UnauthenticatedClientActionV2PtrOutput)
}

type GlobalValidationResponse struct {
	ExcludedPaths               []string `pulumi:"excludedPaths"`
	RedirectToProvider          *string  `pulumi:"redirectToProvider"`
	UnauthenticatedClientAction *string  `pulumi:"unauthenticatedClientAction"`
}

type GlobalValidationResponseOutput struct{ *pulumi.OutputState }

func (GlobalValidationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalValidationResponse)(nil)).Elem()
}

func (o GlobalValidationResponseOutput) ToGlobalValidationResponseOutput() GlobalValidationResponseOutput {
	return o
}

func (o GlobalValidationResponseOutput) ToGlobalValidationResponseOutputWithContext(ctx context.Context) GlobalValidationResponseOutput {
	return o
}

func (o GlobalValidationResponseOutput) ExcludedPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GlobalValidationResponse) []string { return v.ExcludedPaths }).(pulumi.StringArrayOutput)
}

func (o GlobalValidationResponseOutput) RedirectToProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GlobalValidationResponse) *string { return v.RedirectToProvider }).(pulumi.StringPtrOutput)
}

func (o GlobalValidationResponseOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GlobalValidationResponse) *string { return v.UnauthenticatedClientAction }).(pulumi.StringPtrOutput)
}

type GlobalValidationResponsePtrOutput struct{ *pulumi.OutputState }

func (GlobalValidationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalValidationResponse)(nil)).Elem()
}

func (o GlobalValidationResponsePtrOutput) ToGlobalValidationResponsePtrOutput() GlobalValidationResponsePtrOutput {
	return o
}

func (o GlobalValidationResponsePtrOutput) ToGlobalValidationResponsePtrOutputWithContext(ctx context.Context) GlobalValidationResponsePtrOutput {
	return o
}

func (o GlobalValidationResponsePtrOutput) Elem() GlobalValidationResponseOutput {
	return o.ApplyT(func(v *GlobalValidationResponse) GlobalValidationResponse {
		if v != nil {
			return *v
		}
		var ret GlobalValidationResponse
		return ret
	}).(GlobalValidationResponseOutput)
}

func (o GlobalValidationResponsePtrOutput) ExcludedPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GlobalValidationResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedPaths
	}).(pulumi.StringArrayOutput)
}

func (o GlobalValidationResponsePtrOutput) RedirectToProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalValidationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RedirectToProvider
	}).(pulumi.StringPtrOutput)
}

func (o GlobalValidationResponsePtrOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalValidationResponse) *string {
		if v == nil {
			return nil
		}
		return v.UnauthenticatedClientAction
	}).(pulumi.StringPtrOutput)
}

type Google struct {
	Enabled      *bool                       `pulumi:"enabled"`
	Login        *LoginScopes                `pulumi:"login"`
	Registration *ClientRegistration         `pulumi:"registration"`
	Validation   *AllowedAudiencesValidation `pulumi:"validation"`
}





type GoogleInput interface {
	pulumi.Input

	ToGoogleOutput() GoogleOutput
	ToGoogleOutputWithContext(context.Context) GoogleOutput
}

type GoogleArgs struct {
	Enabled      pulumi.BoolPtrInput                `pulumi:"enabled"`
	Login        LoginScopesPtrInput                `pulumi:"login"`
	Registration ClientRegistrationPtrInput         `pulumi:"registration"`
	Validation   AllowedAudiencesValidationPtrInput `pulumi:"validation"`
}

func (GoogleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Google)(nil)).Elem()
}

func (i GoogleArgs) ToGoogleOutput() GoogleOutput {
	return i.ToGoogleOutputWithContext(context.Background())
}

func (i GoogleArgs) ToGoogleOutputWithContext(ctx context.Context) GoogleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleOutput)
}

func (i GoogleArgs) ToGooglePtrOutput() GooglePtrOutput {
	return i.ToGooglePtrOutputWithContext(context.Background())
}

func (i GoogleArgs) ToGooglePtrOutputWithContext(ctx context.Context) GooglePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleOutput).ToGooglePtrOutputWithContext(ctx)
}









type GooglePtrInput interface {
	pulumi.Input

	ToGooglePtrOutput() GooglePtrOutput
	ToGooglePtrOutputWithContext(context.Context) GooglePtrOutput
}

type googlePtrType GoogleArgs

func GooglePtr(v *GoogleArgs) GooglePtrInput {
	return (*googlePtrType)(v)
}

func (*googlePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Google)(nil)).Elem()
}

func (i *googlePtrType) ToGooglePtrOutput() GooglePtrOutput {
	return i.ToGooglePtrOutputWithContext(context.Background())
}

func (i *googlePtrType) ToGooglePtrOutputWithContext(ctx context.Context) GooglePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GooglePtrOutput)
}

type GoogleOutput struct{ *pulumi.OutputState }

func (GoogleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Google)(nil)).Elem()
}

func (o GoogleOutput) ToGoogleOutput() GoogleOutput {
	return o
}

func (o GoogleOutput) ToGoogleOutputWithContext(ctx context.Context) GoogleOutput {
	return o
}

func (o GoogleOutput) ToGooglePtrOutput() GooglePtrOutput {
	return o.ToGooglePtrOutputWithContext(context.Background())
}

func (o GoogleOutput) ToGooglePtrOutputWithContext(ctx context.Context) GooglePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Google) *Google {
		return &v
	}).(GooglePtrOutput)
}

func (o GoogleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Google) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o GoogleOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v Google) *LoginScopes { return v.Login }).(LoginScopesPtrOutput)
}

func (o GoogleOutput) Registration() ClientRegistrationPtrOutput {
	return o.ApplyT(func(v Google) *ClientRegistration { return v.Registration }).(ClientRegistrationPtrOutput)
}

func (o GoogleOutput) Validation() AllowedAudiencesValidationPtrOutput {
	return o.ApplyT(func(v Google) *AllowedAudiencesValidation { return v.Validation }).(AllowedAudiencesValidationPtrOutput)
}

type GooglePtrOutput struct{ *pulumi.OutputState }

func (GooglePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Google)(nil)).Elem()
}

func (o GooglePtrOutput) ToGooglePtrOutput() GooglePtrOutput {
	return o
}

func (o GooglePtrOutput) ToGooglePtrOutputWithContext(ctx context.Context) GooglePtrOutput {
	return o
}

func (o GooglePtrOutput) Elem() GoogleOutput {
	return o.ApplyT(func(v *Google) Google {
		if v != nil {
			return *v
		}
		var ret Google
		return ret
	}).(GoogleOutput)
}

func (o GooglePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Google) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o GooglePtrOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v *Google) *LoginScopes {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesPtrOutput)
}

func (o GooglePtrOutput) Registration() ClientRegistrationPtrOutput {
	return o.ApplyT(func(v *Google) *ClientRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(ClientRegistrationPtrOutput)
}

func (o GooglePtrOutput) Validation() AllowedAudiencesValidationPtrOutput {
	return o.ApplyT(func(v *Google) *AllowedAudiencesValidation {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(AllowedAudiencesValidationPtrOutput)
}

type GoogleResponse struct {
	Enabled      *bool                               `pulumi:"enabled"`
	Login        *LoginScopesResponse                `pulumi:"login"`
	Registration *ClientRegistrationResponse         `pulumi:"registration"`
	Validation   *AllowedAudiencesValidationResponse `pulumi:"validation"`
}

type GoogleResponseOutput struct{ *pulumi.OutputState }

func (GoogleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleResponse)(nil)).Elem()
}

func (o GoogleResponseOutput) ToGoogleResponseOutput() GoogleResponseOutput {
	return o
}

func (o GoogleResponseOutput) ToGoogleResponseOutputWithContext(ctx context.Context) GoogleResponseOutput {
	return o
}

func (o GoogleResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GoogleResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o GoogleResponseOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v GoogleResponse) *LoginScopesResponse { return v.Login }).(LoginScopesResponsePtrOutput)
}

func (o GoogleResponseOutput) Registration() ClientRegistrationResponsePtrOutput {
	return o.ApplyT(func(v GoogleResponse) *ClientRegistrationResponse { return v.Registration }).(ClientRegistrationResponsePtrOutput)
}

func (o GoogleResponseOutput) Validation() AllowedAudiencesValidationResponsePtrOutput {
	return o.ApplyT(func(v GoogleResponse) *AllowedAudiencesValidationResponse { return v.Validation }).(AllowedAudiencesValidationResponsePtrOutput)
}

type GoogleResponsePtrOutput struct{ *pulumi.OutputState }

func (GoogleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleResponse)(nil)).Elem()
}

func (o GoogleResponsePtrOutput) ToGoogleResponsePtrOutput() GoogleResponsePtrOutput {
	return o
}

func (o GoogleResponsePtrOutput) ToGoogleResponsePtrOutputWithContext(ctx context.Context) GoogleResponsePtrOutput {
	return o
}

func (o GoogleResponsePtrOutput) Elem() GoogleResponseOutput {
	return o.ApplyT(func(v *GoogleResponse) GoogleResponse {
		if v != nil {
			return *v
		}
		var ret GoogleResponse
		return ret
	}).(GoogleResponseOutput)
}

func (o GoogleResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GoogleResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o GoogleResponsePtrOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v *GoogleResponse) *LoginScopesResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesResponsePtrOutput)
}

func (o GoogleResponsePtrOutput) Registration() ClientRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *GoogleResponse) *ClientRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(ClientRegistrationResponsePtrOutput)
}

func (o GoogleResponsePtrOutput) Validation() AllowedAudiencesValidationResponsePtrOutput {
	return o.ApplyT(func(v *GoogleResponse) *AllowedAudiencesValidationResponse {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(AllowedAudiencesValidationResponsePtrOutput)
}

type HttpScaleRule struct {
	Auth     []ScaleRuleAuth   `pulumi:"auth"`
	Metadata map[string]string `pulumi:"metadata"`
}





type HttpScaleRuleInput interface {
	pulumi.Input

	ToHttpScaleRuleOutput() HttpScaleRuleOutput
	ToHttpScaleRuleOutputWithContext(context.Context) HttpScaleRuleOutput
}

type HttpScaleRuleArgs struct {
	Auth     ScaleRuleAuthArrayInput `pulumi:"auth"`
	Metadata pulumi.StringMapInput   `pulumi:"metadata"`
}

func (HttpScaleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpScaleRule)(nil)).Elem()
}

func (i HttpScaleRuleArgs) ToHttpScaleRuleOutput() HttpScaleRuleOutput {
	return i.ToHttpScaleRuleOutputWithContext(context.Background())
}

func (i HttpScaleRuleArgs) ToHttpScaleRuleOutputWithContext(ctx context.Context) HttpScaleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpScaleRuleOutput)
}

func (i HttpScaleRuleArgs) ToHttpScaleRulePtrOutput() HttpScaleRulePtrOutput {
	return i.ToHttpScaleRulePtrOutputWithContext(context.Background())
}

func (i HttpScaleRuleArgs) ToHttpScaleRulePtrOutputWithContext(ctx context.Context) HttpScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpScaleRuleOutput).ToHttpScaleRulePtrOutputWithContext(ctx)
}









type HttpScaleRulePtrInput interface {
	pulumi.Input

	ToHttpScaleRulePtrOutput() HttpScaleRulePtrOutput
	ToHttpScaleRulePtrOutputWithContext(context.Context) HttpScaleRulePtrOutput
}

type httpScaleRulePtrType HttpScaleRuleArgs

func HttpScaleRulePtr(v *HttpScaleRuleArgs) HttpScaleRulePtrInput {
	return (*httpScaleRulePtrType)(v)
}

func (*httpScaleRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpScaleRule)(nil)).Elem()
}

func (i *httpScaleRulePtrType) ToHttpScaleRulePtrOutput() HttpScaleRulePtrOutput {
	return i.ToHttpScaleRulePtrOutputWithContext(context.Background())
}

func (i *httpScaleRulePtrType) ToHttpScaleRulePtrOutputWithContext(ctx context.Context) HttpScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpScaleRulePtrOutput)
}

type HttpScaleRuleOutput struct{ *pulumi.OutputState }

func (HttpScaleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpScaleRule)(nil)).Elem()
}

func (o HttpScaleRuleOutput) ToHttpScaleRuleOutput() HttpScaleRuleOutput {
	return o
}

func (o HttpScaleRuleOutput) ToHttpScaleRuleOutputWithContext(ctx context.Context) HttpScaleRuleOutput {
	return o
}

func (o HttpScaleRuleOutput) ToHttpScaleRulePtrOutput() HttpScaleRulePtrOutput {
	return o.ToHttpScaleRulePtrOutputWithContext(context.Background())
}

func (o HttpScaleRuleOutput) ToHttpScaleRulePtrOutputWithContext(ctx context.Context) HttpScaleRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpScaleRule) *HttpScaleRule {
		return &v
	}).(HttpScaleRulePtrOutput)
}

func (o HttpScaleRuleOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v HttpScaleRule) []ScaleRuleAuth { return v.Auth }).(ScaleRuleAuthArrayOutput)
}

func (o HttpScaleRuleOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v HttpScaleRule) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

type HttpScaleRulePtrOutput struct{ *pulumi.OutputState }

func (HttpScaleRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpScaleRule)(nil)).Elem()
}

func (o HttpScaleRulePtrOutput) ToHttpScaleRulePtrOutput() HttpScaleRulePtrOutput {
	return o
}

func (o HttpScaleRulePtrOutput) ToHttpScaleRulePtrOutputWithContext(ctx context.Context) HttpScaleRulePtrOutput {
	return o
}

func (o HttpScaleRulePtrOutput) Elem() HttpScaleRuleOutput {
	return o.ApplyT(func(v *HttpScaleRule) HttpScaleRule {
		if v != nil {
			return *v
		}
		var ret HttpScaleRule
		return ret
	}).(HttpScaleRuleOutput)
}

func (o HttpScaleRulePtrOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v *HttpScaleRule) []ScaleRuleAuth {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthArrayOutput)
}

func (o HttpScaleRulePtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HttpScaleRule) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

type HttpScaleRuleResponse struct {
	Auth     []ScaleRuleAuthResponse `pulumi:"auth"`
	Metadata map[string]string       `pulumi:"metadata"`
}

type HttpScaleRuleResponseOutput struct{ *pulumi.OutputState }

func (HttpScaleRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpScaleRuleResponse)(nil)).Elem()
}

func (o HttpScaleRuleResponseOutput) ToHttpScaleRuleResponseOutput() HttpScaleRuleResponseOutput {
	return o
}

func (o HttpScaleRuleResponseOutput) ToHttpScaleRuleResponseOutputWithContext(ctx context.Context) HttpScaleRuleResponseOutput {
	return o
}

func (o HttpScaleRuleResponseOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v HttpScaleRuleResponse) []ScaleRuleAuthResponse { return v.Auth }).(ScaleRuleAuthResponseArrayOutput)
}

func (o HttpScaleRuleResponseOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v HttpScaleRuleResponse) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

type HttpScaleRuleResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpScaleRuleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpScaleRuleResponse)(nil)).Elem()
}

func (o HttpScaleRuleResponsePtrOutput) ToHttpScaleRuleResponsePtrOutput() HttpScaleRuleResponsePtrOutput {
	return o
}

func (o HttpScaleRuleResponsePtrOutput) ToHttpScaleRuleResponsePtrOutputWithContext(ctx context.Context) HttpScaleRuleResponsePtrOutput {
	return o
}

func (o HttpScaleRuleResponsePtrOutput) Elem() HttpScaleRuleResponseOutput {
	return o.ApplyT(func(v *HttpScaleRuleResponse) HttpScaleRuleResponse {
		if v != nil {
			return *v
		}
		var ret HttpScaleRuleResponse
		return ret
	}).(HttpScaleRuleResponseOutput)
}

func (o HttpScaleRuleResponsePtrOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v *HttpScaleRuleResponse) []ScaleRuleAuthResponse {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthResponseArrayOutput)
}

func (o HttpScaleRuleResponsePtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HttpScaleRuleResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

type HttpSettings struct {
	ForwardProxy *ForwardProxy       `pulumi:"forwardProxy"`
	RequireHttps *bool               `pulumi:"requireHttps"`
	Routes       *HttpSettingsRoutes `pulumi:"routes"`
}





type HttpSettingsInput interface {
	pulumi.Input

	ToHttpSettingsOutput() HttpSettingsOutput
	ToHttpSettingsOutputWithContext(context.Context) HttpSettingsOutput
}

type HttpSettingsArgs struct {
	ForwardProxy ForwardProxyPtrInput       `pulumi:"forwardProxy"`
	RequireHttps pulumi.BoolPtrInput        `pulumi:"requireHttps"`
	Routes       HttpSettingsRoutesPtrInput `pulumi:"routes"`
}

func (HttpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettings)(nil)).Elem()
}

func (i HttpSettingsArgs) ToHttpSettingsOutput() HttpSettingsOutput {
	return i.ToHttpSettingsOutputWithContext(context.Background())
}

func (i HttpSettingsArgs) ToHttpSettingsOutputWithContext(ctx context.Context) HttpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsOutput)
}

func (i HttpSettingsArgs) ToHttpSettingsPtrOutput() HttpSettingsPtrOutput {
	return i.ToHttpSettingsPtrOutputWithContext(context.Background())
}

func (i HttpSettingsArgs) ToHttpSettingsPtrOutputWithContext(ctx context.Context) HttpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsOutput).ToHttpSettingsPtrOutputWithContext(ctx)
}









type HttpSettingsPtrInput interface {
	pulumi.Input

	ToHttpSettingsPtrOutput() HttpSettingsPtrOutput
	ToHttpSettingsPtrOutputWithContext(context.Context) HttpSettingsPtrOutput
}

type httpSettingsPtrType HttpSettingsArgs

func HttpSettingsPtr(v *HttpSettingsArgs) HttpSettingsPtrInput {
	return (*httpSettingsPtrType)(v)
}

func (*httpSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettings)(nil)).Elem()
}

func (i *httpSettingsPtrType) ToHttpSettingsPtrOutput() HttpSettingsPtrOutput {
	return i.ToHttpSettingsPtrOutputWithContext(context.Background())
}

func (i *httpSettingsPtrType) ToHttpSettingsPtrOutputWithContext(ctx context.Context) HttpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsPtrOutput)
}

type HttpSettingsOutput struct{ *pulumi.OutputState }

func (HttpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettings)(nil)).Elem()
}

func (o HttpSettingsOutput) ToHttpSettingsOutput() HttpSettingsOutput {
	return o
}

func (o HttpSettingsOutput) ToHttpSettingsOutputWithContext(ctx context.Context) HttpSettingsOutput {
	return o
}

func (o HttpSettingsOutput) ToHttpSettingsPtrOutput() HttpSettingsPtrOutput {
	return o.ToHttpSettingsPtrOutputWithContext(context.Background())
}

func (o HttpSettingsOutput) ToHttpSettingsPtrOutputWithContext(ctx context.Context) HttpSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpSettings) *HttpSettings {
		return &v
	}).(HttpSettingsPtrOutput)
}

func (o HttpSettingsOutput) ForwardProxy() ForwardProxyPtrOutput {
	return o.ApplyT(func(v HttpSettings) *ForwardProxy { return v.ForwardProxy }).(ForwardProxyPtrOutput)
}

func (o HttpSettingsOutput) RequireHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HttpSettings) *bool { return v.RequireHttps }).(pulumi.BoolPtrOutput)
}

func (o HttpSettingsOutput) Routes() HttpSettingsRoutesPtrOutput {
	return o.ApplyT(func(v HttpSettings) *HttpSettingsRoutes { return v.Routes }).(HttpSettingsRoutesPtrOutput)
}

type HttpSettingsPtrOutput struct{ *pulumi.OutputState }

func (HttpSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettings)(nil)).Elem()
}

func (o HttpSettingsPtrOutput) ToHttpSettingsPtrOutput() HttpSettingsPtrOutput {
	return o
}

func (o HttpSettingsPtrOutput) ToHttpSettingsPtrOutputWithContext(ctx context.Context) HttpSettingsPtrOutput {
	return o
}

func (o HttpSettingsPtrOutput) Elem() HttpSettingsOutput {
	return o.ApplyT(func(v *HttpSettings) HttpSettings {
		if v != nil {
			return *v
		}
		var ret HttpSettings
		return ret
	}).(HttpSettingsOutput)
}

func (o HttpSettingsPtrOutput) ForwardProxy() ForwardProxyPtrOutput {
	return o.ApplyT(func(v *HttpSettings) *ForwardProxy {
		if v == nil {
			return nil
		}
		return v.ForwardProxy
	}).(ForwardProxyPtrOutput)
}

func (o HttpSettingsPtrOutput) RequireHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RequireHttps
	}).(pulumi.BoolPtrOutput)
}

func (o HttpSettingsPtrOutput) Routes() HttpSettingsRoutesPtrOutput {
	return o.ApplyT(func(v *HttpSettings) *HttpSettingsRoutes {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(HttpSettingsRoutesPtrOutput)
}

type HttpSettingsResponse struct {
	ForwardProxy *ForwardProxyResponse       `pulumi:"forwardProxy"`
	RequireHttps *bool                       `pulumi:"requireHttps"`
	Routes       *HttpSettingsRoutesResponse `pulumi:"routes"`
}

type HttpSettingsResponseOutput struct{ *pulumi.OutputState }

func (HttpSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettingsResponse)(nil)).Elem()
}

func (o HttpSettingsResponseOutput) ToHttpSettingsResponseOutput() HttpSettingsResponseOutput {
	return o
}

func (o HttpSettingsResponseOutput) ToHttpSettingsResponseOutputWithContext(ctx context.Context) HttpSettingsResponseOutput {
	return o
}

func (o HttpSettingsResponseOutput) ForwardProxy() ForwardProxyResponsePtrOutput {
	return o.ApplyT(func(v HttpSettingsResponse) *ForwardProxyResponse { return v.ForwardProxy }).(ForwardProxyResponsePtrOutput)
}

func (o HttpSettingsResponseOutput) RequireHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HttpSettingsResponse) *bool { return v.RequireHttps }).(pulumi.BoolPtrOutput)
}

func (o HttpSettingsResponseOutput) Routes() HttpSettingsRoutesResponsePtrOutput {
	return o.ApplyT(func(v HttpSettingsResponse) *HttpSettingsRoutesResponse { return v.Routes }).(HttpSettingsRoutesResponsePtrOutput)
}

type HttpSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettingsResponse)(nil)).Elem()
}

func (o HttpSettingsResponsePtrOutput) ToHttpSettingsResponsePtrOutput() HttpSettingsResponsePtrOutput {
	return o
}

func (o HttpSettingsResponsePtrOutput) ToHttpSettingsResponsePtrOutputWithContext(ctx context.Context) HttpSettingsResponsePtrOutput {
	return o
}

func (o HttpSettingsResponsePtrOutput) Elem() HttpSettingsResponseOutput {
	return o.ApplyT(func(v *HttpSettingsResponse) HttpSettingsResponse {
		if v != nil {
			return *v
		}
		var ret HttpSettingsResponse
		return ret
	}).(HttpSettingsResponseOutput)
}

func (o HttpSettingsResponsePtrOutput) ForwardProxy() ForwardProxyResponsePtrOutput {
	return o.ApplyT(func(v *HttpSettingsResponse) *ForwardProxyResponse {
		if v == nil {
			return nil
		}
		return v.ForwardProxy
	}).(ForwardProxyResponsePtrOutput)
}

func (o HttpSettingsResponsePtrOutput) RequireHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequireHttps
	}).(pulumi.BoolPtrOutput)
}

func (o HttpSettingsResponsePtrOutput) Routes() HttpSettingsRoutesResponsePtrOutput {
	return o.ApplyT(func(v *HttpSettingsResponse) *HttpSettingsRoutesResponse {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(HttpSettingsRoutesResponsePtrOutput)
}

type HttpSettingsRoutes struct {
	ApiPrefix *string `pulumi:"apiPrefix"`
}





type HttpSettingsRoutesInput interface {
	pulumi.Input

	ToHttpSettingsRoutesOutput() HttpSettingsRoutesOutput
	ToHttpSettingsRoutesOutputWithContext(context.Context) HttpSettingsRoutesOutput
}

type HttpSettingsRoutesArgs struct {
	ApiPrefix pulumi.StringPtrInput `pulumi:"apiPrefix"`
}

func (HttpSettingsRoutesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettingsRoutes)(nil)).Elem()
}

func (i HttpSettingsRoutesArgs) ToHttpSettingsRoutesOutput() HttpSettingsRoutesOutput {
	return i.ToHttpSettingsRoutesOutputWithContext(context.Background())
}

func (i HttpSettingsRoutesArgs) ToHttpSettingsRoutesOutputWithContext(ctx context.Context) HttpSettingsRoutesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsRoutesOutput)
}

func (i HttpSettingsRoutesArgs) ToHttpSettingsRoutesPtrOutput() HttpSettingsRoutesPtrOutput {
	return i.ToHttpSettingsRoutesPtrOutputWithContext(context.Background())
}

func (i HttpSettingsRoutesArgs) ToHttpSettingsRoutesPtrOutputWithContext(ctx context.Context) HttpSettingsRoutesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsRoutesOutput).ToHttpSettingsRoutesPtrOutputWithContext(ctx)
}









type HttpSettingsRoutesPtrInput interface {
	pulumi.Input

	ToHttpSettingsRoutesPtrOutput() HttpSettingsRoutesPtrOutput
	ToHttpSettingsRoutesPtrOutputWithContext(context.Context) HttpSettingsRoutesPtrOutput
}

type httpSettingsRoutesPtrType HttpSettingsRoutesArgs

func HttpSettingsRoutesPtr(v *HttpSettingsRoutesArgs) HttpSettingsRoutesPtrInput {
	return (*httpSettingsRoutesPtrType)(v)
}

func (*httpSettingsRoutesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettingsRoutes)(nil)).Elem()
}

func (i *httpSettingsRoutesPtrType) ToHttpSettingsRoutesPtrOutput() HttpSettingsRoutesPtrOutput {
	return i.ToHttpSettingsRoutesPtrOutputWithContext(context.Background())
}

func (i *httpSettingsRoutesPtrType) ToHttpSettingsRoutesPtrOutputWithContext(ctx context.Context) HttpSettingsRoutesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsRoutesPtrOutput)
}

type HttpSettingsRoutesOutput struct{ *pulumi.OutputState }

func (HttpSettingsRoutesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettingsRoutes)(nil)).Elem()
}

func (o HttpSettingsRoutesOutput) ToHttpSettingsRoutesOutput() HttpSettingsRoutesOutput {
	return o
}

func (o HttpSettingsRoutesOutput) ToHttpSettingsRoutesOutputWithContext(ctx context.Context) HttpSettingsRoutesOutput {
	return o
}

func (o HttpSettingsRoutesOutput) ToHttpSettingsRoutesPtrOutput() HttpSettingsRoutesPtrOutput {
	return o.ToHttpSettingsRoutesPtrOutputWithContext(context.Background())
}

func (o HttpSettingsRoutesOutput) ToHttpSettingsRoutesPtrOutputWithContext(ctx context.Context) HttpSettingsRoutesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpSettingsRoutes) *HttpSettingsRoutes {
		return &v
	}).(HttpSettingsRoutesPtrOutput)
}

func (o HttpSettingsRoutesOutput) ApiPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpSettingsRoutes) *string { return v.ApiPrefix }).(pulumi.StringPtrOutput)
}

type HttpSettingsRoutesPtrOutput struct{ *pulumi.OutputState }

func (HttpSettingsRoutesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettingsRoutes)(nil)).Elem()
}

func (o HttpSettingsRoutesPtrOutput) ToHttpSettingsRoutesPtrOutput() HttpSettingsRoutesPtrOutput {
	return o
}

func (o HttpSettingsRoutesPtrOutput) ToHttpSettingsRoutesPtrOutputWithContext(ctx context.Context) HttpSettingsRoutesPtrOutput {
	return o
}

func (o HttpSettingsRoutesPtrOutput) Elem() HttpSettingsRoutesOutput {
	return o.ApplyT(func(v *HttpSettingsRoutes) HttpSettingsRoutes {
		if v != nil {
			return *v
		}
		var ret HttpSettingsRoutes
		return ret
	}).(HttpSettingsRoutesOutput)
}

func (o HttpSettingsRoutesPtrOutput) ApiPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpSettingsRoutes) *string {
		if v == nil {
			return nil
		}
		return v.ApiPrefix
	}).(pulumi.StringPtrOutput)
}

type HttpSettingsRoutesResponse struct {
	ApiPrefix *string `pulumi:"apiPrefix"`
}

type HttpSettingsRoutesResponseOutput struct{ *pulumi.OutputState }

func (HttpSettingsRoutesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettingsRoutesResponse)(nil)).Elem()
}

func (o HttpSettingsRoutesResponseOutput) ToHttpSettingsRoutesResponseOutput() HttpSettingsRoutesResponseOutput {
	return o
}

func (o HttpSettingsRoutesResponseOutput) ToHttpSettingsRoutesResponseOutputWithContext(ctx context.Context) HttpSettingsRoutesResponseOutput {
	return o
}

func (o HttpSettingsRoutesResponseOutput) ApiPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpSettingsRoutesResponse) *string { return v.ApiPrefix }).(pulumi.StringPtrOutput)
}

type HttpSettingsRoutesResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpSettingsRoutesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettingsRoutesResponse)(nil)).Elem()
}

func (o HttpSettingsRoutesResponsePtrOutput) ToHttpSettingsRoutesResponsePtrOutput() HttpSettingsRoutesResponsePtrOutput {
	return o
}

func (o HttpSettingsRoutesResponsePtrOutput) ToHttpSettingsRoutesResponsePtrOutputWithContext(ctx context.Context) HttpSettingsRoutesResponsePtrOutput {
	return o
}

func (o HttpSettingsRoutesResponsePtrOutput) Elem() HttpSettingsRoutesResponseOutput {
	return o.ApplyT(func(v *HttpSettingsRoutesResponse) HttpSettingsRoutesResponse {
		if v != nil {
			return *v
		}
		var ret HttpSettingsRoutesResponse
		return ret
	}).(HttpSettingsRoutesResponseOutput)
}

func (o HttpSettingsRoutesResponsePtrOutput) ApiPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpSettingsRoutesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiPrefix
	}).(pulumi.StringPtrOutput)
}

type IdentityProviders struct {
	Apple                        *Apple                                 `pulumi:"apple"`
	AzureActiveDirectory         *AzureActiveDirectory                  `pulumi:"azureActiveDirectory"`
	AzureStaticWebApps           *AzureStaticWebApps                    `pulumi:"azureStaticWebApps"`
	CustomOpenIdConnectProviders map[string]CustomOpenIdConnectProvider `pulumi:"customOpenIdConnectProviders"`
	Facebook                     *Facebook                              `pulumi:"facebook"`
	GitHub                       *GitHub                                `pulumi:"gitHub"`
	Google                       *Google                                `pulumi:"google"`
	Twitter                      *Twitter                               `pulumi:"twitter"`
}





type IdentityProvidersInput interface {
	pulumi.Input

	ToIdentityProvidersOutput() IdentityProvidersOutput
	ToIdentityProvidersOutputWithContext(context.Context) IdentityProvidersOutput
}

type IdentityProvidersArgs struct {
	Apple                        ApplePtrInput                       `pulumi:"apple"`
	AzureActiveDirectory         AzureActiveDirectoryPtrInput        `pulumi:"azureActiveDirectory"`
	AzureStaticWebApps           AzureStaticWebAppsPtrInput          `pulumi:"azureStaticWebApps"`
	CustomOpenIdConnectProviders CustomOpenIdConnectProviderMapInput `pulumi:"customOpenIdConnectProviders"`
	Facebook                     FacebookPtrInput                    `pulumi:"facebook"`
	GitHub                       GitHubPtrInput                      `pulumi:"gitHub"`
	Google                       GooglePtrInput                      `pulumi:"google"`
	Twitter                      TwitterPtrInput                     `pulumi:"twitter"`
}

func (IdentityProvidersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviders)(nil)).Elem()
}

func (i IdentityProvidersArgs) ToIdentityProvidersOutput() IdentityProvidersOutput {
	return i.ToIdentityProvidersOutputWithContext(context.Background())
}

func (i IdentityProvidersArgs) ToIdentityProvidersOutputWithContext(ctx context.Context) IdentityProvidersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProvidersOutput)
}

func (i IdentityProvidersArgs) ToIdentityProvidersPtrOutput() IdentityProvidersPtrOutput {
	return i.ToIdentityProvidersPtrOutputWithContext(context.Background())
}

func (i IdentityProvidersArgs) ToIdentityProvidersPtrOutputWithContext(ctx context.Context) IdentityProvidersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProvidersOutput).ToIdentityProvidersPtrOutputWithContext(ctx)
}









type IdentityProvidersPtrInput interface {
	pulumi.Input

	ToIdentityProvidersPtrOutput() IdentityProvidersPtrOutput
	ToIdentityProvidersPtrOutputWithContext(context.Context) IdentityProvidersPtrOutput
}

type identityProvidersPtrType IdentityProvidersArgs

func IdentityProvidersPtr(v *IdentityProvidersArgs) IdentityProvidersPtrInput {
	return (*identityProvidersPtrType)(v)
}

func (*identityProvidersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviders)(nil)).Elem()
}

func (i *identityProvidersPtrType) ToIdentityProvidersPtrOutput() IdentityProvidersPtrOutput {
	return i.ToIdentityProvidersPtrOutputWithContext(context.Background())
}

func (i *identityProvidersPtrType) ToIdentityProvidersPtrOutputWithContext(ctx context.Context) IdentityProvidersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProvidersPtrOutput)
}

type IdentityProvidersOutput struct{ *pulumi.OutputState }

func (IdentityProvidersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviders)(nil)).Elem()
}

func (o IdentityProvidersOutput) ToIdentityProvidersOutput() IdentityProvidersOutput {
	return o
}

func (o IdentityProvidersOutput) ToIdentityProvidersOutputWithContext(ctx context.Context) IdentityProvidersOutput {
	return o
}

func (o IdentityProvidersOutput) ToIdentityProvidersPtrOutput() IdentityProvidersPtrOutput {
	return o.ToIdentityProvidersPtrOutputWithContext(context.Background())
}

func (o IdentityProvidersOutput) ToIdentityProvidersPtrOutputWithContext(ctx context.Context) IdentityProvidersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProviders) *IdentityProviders {
		return &v
	}).(IdentityProvidersPtrOutput)
}

func (o IdentityProvidersOutput) Apple() ApplePtrOutput {
	return o.ApplyT(func(v IdentityProviders) *Apple { return v.Apple }).(ApplePtrOutput)
}

func (o IdentityProvidersOutput) AzureActiveDirectory() AzureActiveDirectoryPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *AzureActiveDirectory { return v.AzureActiveDirectory }).(AzureActiveDirectoryPtrOutput)
}

func (o IdentityProvidersOutput) AzureStaticWebApps() AzureStaticWebAppsPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *AzureStaticWebApps { return v.AzureStaticWebApps }).(AzureStaticWebAppsPtrOutput)
}

func (o IdentityProvidersOutput) CustomOpenIdConnectProviders() CustomOpenIdConnectProviderMapOutput {
	return o.ApplyT(func(v IdentityProviders) map[string]CustomOpenIdConnectProvider {
		return v.CustomOpenIdConnectProviders
	}).(CustomOpenIdConnectProviderMapOutput)
}

func (o IdentityProvidersOutput) Facebook() FacebookPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *Facebook { return v.Facebook }).(FacebookPtrOutput)
}

func (o IdentityProvidersOutput) GitHub() GitHubPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *GitHub { return v.GitHub }).(GitHubPtrOutput)
}

func (o IdentityProvidersOutput) Google() GooglePtrOutput {
	return o.ApplyT(func(v IdentityProviders) *Google { return v.Google }).(GooglePtrOutput)
}

func (o IdentityProvidersOutput) Twitter() TwitterPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *Twitter { return v.Twitter }).(TwitterPtrOutput)
}

type IdentityProvidersPtrOutput struct{ *pulumi.OutputState }

func (IdentityProvidersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviders)(nil)).Elem()
}

func (o IdentityProvidersPtrOutput) ToIdentityProvidersPtrOutput() IdentityProvidersPtrOutput {
	return o
}

func (o IdentityProvidersPtrOutput) ToIdentityProvidersPtrOutputWithContext(ctx context.Context) IdentityProvidersPtrOutput {
	return o
}

func (o IdentityProvidersPtrOutput) Elem() IdentityProvidersOutput {
	return o.ApplyT(func(v *IdentityProviders) IdentityProviders {
		if v != nil {
			return *v
		}
		var ret IdentityProviders
		return ret
	}).(IdentityProvidersOutput)
}

func (o IdentityProvidersPtrOutput) Apple() ApplePtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *Apple {
		if v == nil {
			return nil
		}
		return v.Apple
	}).(ApplePtrOutput)
}

func (o IdentityProvidersPtrOutput) AzureActiveDirectory() AzureActiveDirectoryPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *AzureActiveDirectory {
		if v == nil {
			return nil
		}
		return v.AzureActiveDirectory
	}).(AzureActiveDirectoryPtrOutput)
}

func (o IdentityProvidersPtrOutput) AzureStaticWebApps() AzureStaticWebAppsPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *AzureStaticWebApps {
		if v == nil {
			return nil
		}
		return v.AzureStaticWebApps
	}).(AzureStaticWebAppsPtrOutput)
}

func (o IdentityProvidersPtrOutput) CustomOpenIdConnectProviders() CustomOpenIdConnectProviderMapOutput {
	return o.ApplyT(func(v *IdentityProviders) map[string]CustomOpenIdConnectProvider {
		if v == nil {
			return nil
		}
		return v.CustomOpenIdConnectProviders
	}).(CustomOpenIdConnectProviderMapOutput)
}

func (o IdentityProvidersPtrOutput) Facebook() FacebookPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *Facebook {
		if v == nil {
			return nil
		}
		return v.Facebook
	}).(FacebookPtrOutput)
}

func (o IdentityProvidersPtrOutput) GitHub() GitHubPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *GitHub {
		if v == nil {
			return nil
		}
		return v.GitHub
	}).(GitHubPtrOutput)
}

func (o IdentityProvidersPtrOutput) Google() GooglePtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *Google {
		if v == nil {
			return nil
		}
		return v.Google
	}).(GooglePtrOutput)
}

func (o IdentityProvidersPtrOutput) Twitter() TwitterPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *Twitter {
		if v == nil {
			return nil
		}
		return v.Twitter
	}).(TwitterPtrOutput)
}

type IdentityProvidersResponse struct {
	Apple                        *AppleResponse                                 `pulumi:"apple"`
	AzureActiveDirectory         *AzureActiveDirectoryResponse                  `pulumi:"azureActiveDirectory"`
	AzureStaticWebApps           *AzureStaticWebAppsResponse                    `pulumi:"azureStaticWebApps"`
	CustomOpenIdConnectProviders map[string]CustomOpenIdConnectProviderResponse `pulumi:"customOpenIdConnectProviders"`
	Facebook                     *FacebookResponse                              `pulumi:"facebook"`
	GitHub                       *GitHubResponse                                `pulumi:"gitHub"`
	Google                       *GoogleResponse                                `pulumi:"google"`
	Twitter                      *TwitterResponse                               `pulumi:"twitter"`
}

type IdentityProvidersResponseOutput struct{ *pulumi.OutputState }

func (IdentityProvidersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProvidersResponse)(nil)).Elem()
}

func (o IdentityProvidersResponseOutput) ToIdentityProvidersResponseOutput() IdentityProvidersResponseOutput {
	return o
}

func (o IdentityProvidersResponseOutput) ToIdentityProvidersResponseOutputWithContext(ctx context.Context) IdentityProvidersResponseOutput {
	return o
}

func (o IdentityProvidersResponseOutput) Apple() AppleResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *AppleResponse { return v.Apple }).(AppleResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *AzureActiveDirectoryResponse { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) AzureStaticWebApps() AzureStaticWebAppsResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *AzureStaticWebAppsResponse { return v.AzureStaticWebApps }).(AzureStaticWebAppsResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) CustomOpenIdConnectProviders() CustomOpenIdConnectProviderResponseMapOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) map[string]CustomOpenIdConnectProviderResponse {
		return v.CustomOpenIdConnectProviders
	}).(CustomOpenIdConnectProviderResponseMapOutput)
}

func (o IdentityProvidersResponseOutput) Facebook() FacebookResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *FacebookResponse { return v.Facebook }).(FacebookResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) GitHub() GitHubResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *GitHubResponse { return v.GitHub }).(GitHubResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) Google() GoogleResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *GoogleResponse { return v.Google }).(GoogleResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) Twitter() TwitterResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *TwitterResponse { return v.Twitter }).(TwitterResponsePtrOutput)
}

type IdentityProvidersResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityProvidersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvidersResponse)(nil)).Elem()
}

func (o IdentityProvidersResponsePtrOutput) ToIdentityProvidersResponsePtrOutput() IdentityProvidersResponsePtrOutput {
	return o
}

func (o IdentityProvidersResponsePtrOutput) ToIdentityProvidersResponsePtrOutputWithContext(ctx context.Context) IdentityProvidersResponsePtrOutput {
	return o
}

func (o IdentityProvidersResponsePtrOutput) Elem() IdentityProvidersResponseOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) IdentityProvidersResponse {
		if v != nil {
			return *v
		}
		var ret IdentityProvidersResponse
		return ret
	}).(IdentityProvidersResponseOutput)
}

func (o IdentityProvidersResponsePtrOutput) Apple() AppleResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *AppleResponse {
		if v == nil {
			return nil
		}
		return v.Apple
	}).(AppleResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *AzureActiveDirectoryResponse {
		if v == nil {
			return nil
		}
		return v.AzureActiveDirectory
	}).(AzureActiveDirectoryResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) AzureStaticWebApps() AzureStaticWebAppsResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *AzureStaticWebAppsResponse {
		if v == nil {
			return nil
		}
		return v.AzureStaticWebApps
	}).(AzureStaticWebAppsResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) CustomOpenIdConnectProviders() CustomOpenIdConnectProviderResponseMapOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) map[string]CustomOpenIdConnectProviderResponse {
		if v == nil {
			return nil
		}
		return v.CustomOpenIdConnectProviders
	}).(CustomOpenIdConnectProviderResponseMapOutput)
}

func (o IdentityProvidersResponsePtrOutput) Facebook() FacebookResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *FacebookResponse {
		if v == nil {
			return nil
		}
		return v.Facebook
	}).(FacebookResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) GitHub() GitHubResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *GitHubResponse {
		if v == nil {
			return nil
		}
		return v.GitHub
	}).(GitHubResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) Google() GoogleResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *GoogleResponse {
		if v == nil {
			return nil
		}
		return v.Google
	}).(GoogleResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) Twitter() TwitterResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *TwitterResponse {
		if v == nil {
			return nil
		}
		return v.Twitter
	}).(TwitterResponsePtrOutput)
}

type Ingress struct {
	AllowInsecure *bool           `pulumi:"allowInsecure"`
	CustomDomains []CustomDomain  `pulumi:"customDomains"`
	External      *bool           `pulumi:"external"`
	TargetPort    *int            `pulumi:"targetPort"`
	Traffic       []TrafficWeight `pulumi:"traffic"`
	Transport     *string         `pulumi:"transport"`
}


func (val *Ingress) Defaults() *Ingress {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllowInsecure) {
		allowInsecure_ := false
		tmp.AllowInsecure = &allowInsecure_
	}
	if isZero(tmp.External) {
		external_ := false
		tmp.External = &external_
	}
	if isZero(tmp.Transport) {
		transport_ := "auto"
		tmp.Transport = &transport_
	}
	return &tmp
}





type IngressInput interface {
	pulumi.Input

	ToIngressOutput() IngressOutput
	ToIngressOutputWithContext(context.Context) IngressOutput
}

type IngressArgs struct {
	AllowInsecure pulumi.BoolPtrInput     `pulumi:"allowInsecure"`
	CustomDomains CustomDomainArrayInput  `pulumi:"customDomains"`
	External      pulumi.BoolPtrInput     `pulumi:"external"`
	TargetPort    pulumi.IntPtrInput      `pulumi:"targetPort"`
	Traffic       TrafficWeightArrayInput `pulumi:"traffic"`
	Transport     pulumi.StringPtrInput   `pulumi:"transport"`
}


func (val *IngressArgs) Defaults() *IngressArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllowInsecure) {
		tmp.AllowInsecure = pulumi.BoolPtr(false)
	}
	if isZero(tmp.External) {
		tmp.External = pulumi.BoolPtr(false)
	}
	if isZero(tmp.Transport) {
		tmp.Transport = pulumi.StringPtr("auto")
	}
	return &tmp
}
func (IngressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ingress)(nil)).Elem()
}

func (i IngressArgs) ToIngressOutput() IngressOutput {
	return i.ToIngressOutputWithContext(context.Background())
}

func (i IngressArgs) ToIngressOutputWithContext(ctx context.Context) IngressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressOutput)
}

func (i IngressArgs) ToIngressPtrOutput() IngressPtrOutput {
	return i.ToIngressPtrOutputWithContext(context.Background())
}

func (i IngressArgs) ToIngressPtrOutputWithContext(ctx context.Context) IngressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressOutput).ToIngressPtrOutputWithContext(ctx)
}









type IngressPtrInput interface {
	pulumi.Input

	ToIngressPtrOutput() IngressPtrOutput
	ToIngressPtrOutputWithContext(context.Context) IngressPtrOutput
}

type ingressPtrType IngressArgs

func IngressPtr(v *IngressArgs) IngressPtrInput {
	return (*ingressPtrType)(v)
}

func (*ingressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ingress)(nil)).Elem()
}

func (i *ingressPtrType) ToIngressPtrOutput() IngressPtrOutput {
	return i.ToIngressPtrOutputWithContext(context.Background())
}

func (i *ingressPtrType) ToIngressPtrOutputWithContext(ctx context.Context) IngressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressPtrOutput)
}

type IngressOutput struct{ *pulumi.OutputState }

func (IngressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ingress)(nil)).Elem()
}

func (o IngressOutput) ToIngressOutput() IngressOutput {
	return o
}

func (o IngressOutput) ToIngressOutputWithContext(ctx context.Context) IngressOutput {
	return o
}

func (o IngressOutput) ToIngressPtrOutput() IngressPtrOutput {
	return o.ToIngressPtrOutputWithContext(context.Background())
}

func (o IngressOutput) ToIngressPtrOutputWithContext(ctx context.Context) IngressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ingress) *Ingress {
		return &v
	}).(IngressPtrOutput)
}

func (o IngressOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ingress) *bool { return v.AllowInsecure }).(pulumi.BoolPtrOutput)
}

func (o IngressOutput) CustomDomains() CustomDomainArrayOutput {
	return o.ApplyT(func(v Ingress) []CustomDomain { return v.CustomDomains }).(CustomDomainArrayOutput)
}

func (o IngressOutput) External() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ingress) *bool { return v.External }).(pulumi.BoolPtrOutput)
}

func (o IngressOutput) TargetPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Ingress) *int { return v.TargetPort }).(pulumi.IntPtrOutput)
}

func (o IngressOutput) Traffic() TrafficWeightArrayOutput {
	return o.ApplyT(func(v Ingress) []TrafficWeight { return v.Traffic }).(TrafficWeightArrayOutput)
}

func (o IngressOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ingress) *string { return v.Transport }).(pulumi.StringPtrOutput)
}

type IngressPtrOutput struct{ *pulumi.OutputState }

func (IngressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ingress)(nil)).Elem()
}

func (o IngressPtrOutput) ToIngressPtrOutput() IngressPtrOutput {
	return o
}

func (o IngressPtrOutput) ToIngressPtrOutputWithContext(ctx context.Context) IngressPtrOutput {
	return o
}

func (o IngressPtrOutput) Elem() IngressOutput {
	return o.ApplyT(func(v *Ingress) Ingress {
		if v != nil {
			return *v
		}
		var ret Ingress
		return ret
	}).(IngressOutput)
}

func (o IngressPtrOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ingress) *bool {
		if v == nil {
			return nil
		}
		return v.AllowInsecure
	}).(pulumi.BoolPtrOutput)
}

func (o IngressPtrOutput) CustomDomains() CustomDomainArrayOutput {
	return o.ApplyT(func(v *Ingress) []CustomDomain {
		if v == nil {
			return nil
		}
		return v.CustomDomains
	}).(CustomDomainArrayOutput)
}

func (o IngressPtrOutput) External() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ingress) *bool {
		if v == nil {
			return nil
		}
		return v.External
	}).(pulumi.BoolPtrOutput)
}

func (o IngressPtrOutput) TargetPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ingress) *int {
		if v == nil {
			return nil
		}
		return v.TargetPort
	}).(pulumi.IntPtrOutput)
}

func (o IngressPtrOutput) Traffic() TrafficWeightArrayOutput {
	return o.ApplyT(func(v *Ingress) []TrafficWeight {
		if v == nil {
			return nil
		}
		return v.Traffic
	}).(TrafficWeightArrayOutput)
}

func (o IngressPtrOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ingress) *string {
		if v == nil {
			return nil
		}
		return v.Transport
	}).(pulumi.StringPtrOutput)
}

type IngressResponse struct {
	AllowInsecure *bool                   `pulumi:"allowInsecure"`
	CustomDomains []CustomDomainResponse  `pulumi:"customDomains"`
	External      *bool                   `pulumi:"external"`
	Fqdn          string                  `pulumi:"fqdn"`
	TargetPort    *int                    `pulumi:"targetPort"`
	Traffic       []TrafficWeightResponse `pulumi:"traffic"`
	Transport     *string                 `pulumi:"transport"`
}


func (val *IngressResponse) Defaults() *IngressResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllowInsecure) {
		allowInsecure_ := false
		tmp.AllowInsecure = &allowInsecure_
	}
	if isZero(tmp.External) {
		external_ := false
		tmp.External = &external_
	}
	if isZero(tmp.Transport) {
		transport_ := "auto"
		tmp.Transport = &transport_
	}
	return &tmp
}

type IngressResponseOutput struct{ *pulumi.OutputState }

func (IngressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressResponse)(nil)).Elem()
}

func (o IngressResponseOutput) ToIngressResponseOutput() IngressResponseOutput {
	return o
}

func (o IngressResponseOutput) ToIngressResponseOutputWithContext(ctx context.Context) IngressResponseOutput {
	return o
}

func (o IngressResponseOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngressResponse) *bool { return v.AllowInsecure }).(pulumi.BoolPtrOutput)
}

func (o IngressResponseOutput) CustomDomains() CustomDomainResponseArrayOutput {
	return o.ApplyT(func(v IngressResponse) []CustomDomainResponse { return v.CustomDomains }).(CustomDomainResponseArrayOutput)
}

func (o IngressResponseOutput) External() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngressResponse) *bool { return v.External }).(pulumi.BoolPtrOutput)
}

func (o IngressResponseOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v IngressResponse) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o IngressResponseOutput) TargetPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IngressResponse) *int { return v.TargetPort }).(pulumi.IntPtrOutput)
}

func (o IngressResponseOutput) Traffic() TrafficWeightResponseArrayOutput {
	return o.ApplyT(func(v IngressResponse) []TrafficWeightResponse { return v.Traffic }).(TrafficWeightResponseArrayOutput)
}

func (o IngressResponseOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressResponse) *string { return v.Transport }).(pulumi.StringPtrOutput)
}

type IngressResponsePtrOutput struct{ *pulumi.OutputState }

func (IngressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressResponse)(nil)).Elem()
}

func (o IngressResponsePtrOutput) ToIngressResponsePtrOutput() IngressResponsePtrOutput {
	return o
}

func (o IngressResponsePtrOutput) ToIngressResponsePtrOutputWithContext(ctx context.Context) IngressResponsePtrOutput {
	return o
}

func (o IngressResponsePtrOutput) Elem() IngressResponseOutput {
	return o.ApplyT(func(v *IngressResponse) IngressResponse {
		if v != nil {
			return *v
		}
		var ret IngressResponse
		return ret
	}).(IngressResponseOutput)
}

func (o IngressResponsePtrOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IngressResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowInsecure
	}).(pulumi.BoolPtrOutput)
}

func (o IngressResponsePtrOutput) CustomDomains() CustomDomainResponseArrayOutput {
	return o.ApplyT(func(v *IngressResponse) []CustomDomainResponse {
		if v == nil {
			return nil
		}
		return v.CustomDomains
	}).(CustomDomainResponseArrayOutput)
}

func (o IngressResponsePtrOutput) External() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IngressResponse) *bool {
		if v == nil {
			return nil
		}
		return v.External
	}).(pulumi.BoolPtrOutput)
}

func (o IngressResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o IngressResponsePtrOutput) TargetPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IngressResponse) *int {
		if v == nil {
			return nil
		}
		return v.TargetPort
	}).(pulumi.IntPtrOutput)
}

func (o IngressResponsePtrOutput) Traffic() TrafficWeightResponseArrayOutput {
	return o.ApplyT(func(v *IngressResponse) []TrafficWeightResponse {
		if v == nil {
			return nil
		}
		return v.Traffic
	}).(TrafficWeightResponseArrayOutput)
}

func (o IngressResponsePtrOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressResponse) *string {
		if v == nil {
			return nil
		}
		return v.Transport
	}).(pulumi.StringPtrOutput)
}

type JwtClaimChecks struct {
	AllowedClientApplications []string `pulumi:"allowedClientApplications"`
	AllowedGroups             []string `pulumi:"allowedGroups"`
}





type JwtClaimChecksInput interface {
	pulumi.Input

	ToJwtClaimChecksOutput() JwtClaimChecksOutput
	ToJwtClaimChecksOutputWithContext(context.Context) JwtClaimChecksOutput
}

type JwtClaimChecksArgs struct {
	AllowedClientApplications pulumi.StringArrayInput `pulumi:"allowedClientApplications"`
	AllowedGroups             pulumi.StringArrayInput `pulumi:"allowedGroups"`
}

func (JwtClaimChecksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtClaimChecks)(nil)).Elem()
}

func (i JwtClaimChecksArgs) ToJwtClaimChecksOutput() JwtClaimChecksOutput {
	return i.ToJwtClaimChecksOutputWithContext(context.Background())
}

func (i JwtClaimChecksArgs) ToJwtClaimChecksOutputWithContext(ctx context.Context) JwtClaimChecksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtClaimChecksOutput)
}

func (i JwtClaimChecksArgs) ToJwtClaimChecksPtrOutput() JwtClaimChecksPtrOutput {
	return i.ToJwtClaimChecksPtrOutputWithContext(context.Background())
}

func (i JwtClaimChecksArgs) ToJwtClaimChecksPtrOutputWithContext(ctx context.Context) JwtClaimChecksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtClaimChecksOutput).ToJwtClaimChecksPtrOutputWithContext(ctx)
}









type JwtClaimChecksPtrInput interface {
	pulumi.Input

	ToJwtClaimChecksPtrOutput() JwtClaimChecksPtrOutput
	ToJwtClaimChecksPtrOutputWithContext(context.Context) JwtClaimChecksPtrOutput
}

type jwtClaimChecksPtrType JwtClaimChecksArgs

func JwtClaimChecksPtr(v *JwtClaimChecksArgs) JwtClaimChecksPtrInput {
	return (*jwtClaimChecksPtrType)(v)
}

func (*jwtClaimChecksPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtClaimChecks)(nil)).Elem()
}

func (i *jwtClaimChecksPtrType) ToJwtClaimChecksPtrOutput() JwtClaimChecksPtrOutput {
	return i.ToJwtClaimChecksPtrOutputWithContext(context.Background())
}

func (i *jwtClaimChecksPtrType) ToJwtClaimChecksPtrOutputWithContext(ctx context.Context) JwtClaimChecksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtClaimChecksPtrOutput)
}

type JwtClaimChecksOutput struct{ *pulumi.OutputState }

func (JwtClaimChecksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtClaimChecks)(nil)).Elem()
}

func (o JwtClaimChecksOutput) ToJwtClaimChecksOutput() JwtClaimChecksOutput {
	return o
}

func (o JwtClaimChecksOutput) ToJwtClaimChecksOutputWithContext(ctx context.Context) JwtClaimChecksOutput {
	return o
}

func (o JwtClaimChecksOutput) ToJwtClaimChecksPtrOutput() JwtClaimChecksPtrOutput {
	return o.ToJwtClaimChecksPtrOutputWithContext(context.Background())
}

func (o JwtClaimChecksOutput) ToJwtClaimChecksPtrOutputWithContext(ctx context.Context) JwtClaimChecksPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JwtClaimChecks) *JwtClaimChecks {
		return &v
	}).(JwtClaimChecksPtrOutput)
}

func (o JwtClaimChecksOutput) AllowedClientApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtClaimChecks) []string { return v.AllowedClientApplications }).(pulumi.StringArrayOutput)
}

func (o JwtClaimChecksOutput) AllowedGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtClaimChecks) []string { return v.AllowedGroups }).(pulumi.StringArrayOutput)
}

type JwtClaimChecksPtrOutput struct{ *pulumi.OutputState }

func (JwtClaimChecksPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtClaimChecks)(nil)).Elem()
}

func (o JwtClaimChecksPtrOutput) ToJwtClaimChecksPtrOutput() JwtClaimChecksPtrOutput {
	return o
}

func (o JwtClaimChecksPtrOutput) ToJwtClaimChecksPtrOutputWithContext(ctx context.Context) JwtClaimChecksPtrOutput {
	return o
}

func (o JwtClaimChecksPtrOutput) Elem() JwtClaimChecksOutput {
	return o.ApplyT(func(v *JwtClaimChecks) JwtClaimChecks {
		if v != nil {
			return *v
		}
		var ret JwtClaimChecks
		return ret
	}).(JwtClaimChecksOutput)
}

func (o JwtClaimChecksPtrOutput) AllowedClientApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtClaimChecks) []string {
		if v == nil {
			return nil
		}
		return v.AllowedClientApplications
	}).(pulumi.StringArrayOutput)
}

func (o JwtClaimChecksPtrOutput) AllowedGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtClaimChecks) []string {
		if v == nil {
			return nil
		}
		return v.AllowedGroups
	}).(pulumi.StringArrayOutput)
}

type JwtClaimChecksResponse struct {
	AllowedClientApplications []string `pulumi:"allowedClientApplications"`
	AllowedGroups             []string `pulumi:"allowedGroups"`
}

type JwtClaimChecksResponseOutput struct{ *pulumi.OutputState }

func (JwtClaimChecksResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtClaimChecksResponse)(nil)).Elem()
}

func (o JwtClaimChecksResponseOutput) ToJwtClaimChecksResponseOutput() JwtClaimChecksResponseOutput {
	return o
}

func (o JwtClaimChecksResponseOutput) ToJwtClaimChecksResponseOutputWithContext(ctx context.Context) JwtClaimChecksResponseOutput {
	return o
}

func (o JwtClaimChecksResponseOutput) AllowedClientApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtClaimChecksResponse) []string { return v.AllowedClientApplications }).(pulumi.StringArrayOutput)
}

func (o JwtClaimChecksResponseOutput) AllowedGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtClaimChecksResponse) []string { return v.AllowedGroups }).(pulumi.StringArrayOutput)
}

type JwtClaimChecksResponsePtrOutput struct{ *pulumi.OutputState }

func (JwtClaimChecksResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtClaimChecksResponse)(nil)).Elem()
}

func (o JwtClaimChecksResponsePtrOutput) ToJwtClaimChecksResponsePtrOutput() JwtClaimChecksResponsePtrOutput {
	return o
}

func (o JwtClaimChecksResponsePtrOutput) ToJwtClaimChecksResponsePtrOutputWithContext(ctx context.Context) JwtClaimChecksResponsePtrOutput {
	return o
}

func (o JwtClaimChecksResponsePtrOutput) Elem() JwtClaimChecksResponseOutput {
	return o.ApplyT(func(v *JwtClaimChecksResponse) JwtClaimChecksResponse {
		if v != nil {
			return *v
		}
		var ret JwtClaimChecksResponse
		return ret
	}).(JwtClaimChecksResponseOutput)
}

func (o JwtClaimChecksResponsePtrOutput) AllowedClientApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtClaimChecksResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedClientApplications
	}).(pulumi.StringArrayOutput)
}

func (o JwtClaimChecksResponsePtrOutput) AllowedGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtClaimChecksResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedGroups
	}).(pulumi.StringArrayOutput)
}

type LogAnalyticsConfiguration struct {
	CustomerId *string `pulumi:"customerId"`
	SharedKey  *string `pulumi:"sharedKey"`
}





type LogAnalyticsConfigurationInput interface {
	pulumi.Input

	ToLogAnalyticsConfigurationOutput() LogAnalyticsConfigurationOutput
	ToLogAnalyticsConfigurationOutputWithContext(context.Context) LogAnalyticsConfigurationOutput
}

type LogAnalyticsConfigurationArgs struct {
	CustomerId pulumi.StringPtrInput `pulumi:"customerId"`
	SharedKey  pulumi.StringPtrInput `pulumi:"sharedKey"`
}

func (LogAnalyticsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsConfiguration)(nil)).Elem()
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationOutput() LogAnalyticsConfigurationOutput {
	return i.ToLogAnalyticsConfigurationOutputWithContext(context.Background())
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationOutputWithContext(ctx context.Context) LogAnalyticsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsConfigurationOutput)
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return i.ToLogAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsConfigurationOutput).ToLogAnalyticsConfigurationPtrOutputWithContext(ctx)
}









type LogAnalyticsConfigurationPtrInput interface {
	pulumi.Input

	ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput
	ToLogAnalyticsConfigurationPtrOutputWithContext(context.Context) LogAnalyticsConfigurationPtrOutput
}

type logAnalyticsConfigurationPtrType LogAnalyticsConfigurationArgs

func LogAnalyticsConfigurationPtr(v *LogAnalyticsConfigurationArgs) LogAnalyticsConfigurationPtrInput {
	return (*logAnalyticsConfigurationPtrType)(v)
}

func (*logAnalyticsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsConfiguration)(nil)).Elem()
}

func (i *logAnalyticsConfigurationPtrType) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return i.ToLogAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (i *logAnalyticsConfigurationPtrType) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsConfigurationPtrOutput)
}

type LogAnalyticsConfigurationOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsConfiguration)(nil)).Elem()
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationOutput() LogAnalyticsConfigurationOutput {
	return o
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationOutputWithContext(ctx context.Context) LogAnalyticsConfigurationOutput {
	return o
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return o.ToLogAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalyticsConfiguration) *LogAnalyticsConfiguration {
		return &v
	}).(LogAnalyticsConfigurationPtrOutput)
}

func (o LogAnalyticsConfigurationOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsConfiguration) *string { return v.CustomerId }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsConfigurationOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsConfiguration) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

type LogAnalyticsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsConfiguration)(nil)).Elem()
}

func (o LogAnalyticsConfigurationPtrOutput) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return o
}

func (o LogAnalyticsConfigurationPtrOutput) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return o
}

func (o LogAnalyticsConfigurationPtrOutput) Elem() LogAnalyticsConfigurationOutput {
	return o.ApplyT(func(v *LogAnalyticsConfiguration) LogAnalyticsConfiguration {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsConfiguration
		return ret
	}).(LogAnalyticsConfigurationOutput)
}

func (o LogAnalyticsConfigurationPtrOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CustomerId
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsConfigurationPtrOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SharedKey
	}).(pulumi.StringPtrOutput)
}

type LogAnalyticsConfigurationResponse struct {
	CustomerId *string `pulumi:"customerId"`
}

type LogAnalyticsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsConfigurationResponse)(nil)).Elem()
}

func (o LogAnalyticsConfigurationResponseOutput) ToLogAnalyticsConfigurationResponseOutput() LogAnalyticsConfigurationResponseOutput {
	return o
}

func (o LogAnalyticsConfigurationResponseOutput) ToLogAnalyticsConfigurationResponseOutputWithContext(ctx context.Context) LogAnalyticsConfigurationResponseOutput {
	return o
}

func (o LogAnalyticsConfigurationResponseOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsConfigurationResponse) *string { return v.CustomerId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsConfigurationResponse)(nil)).Elem()
}

func (o LogAnalyticsConfigurationResponsePtrOutput) ToLogAnalyticsConfigurationResponsePtrOutput() LogAnalyticsConfigurationResponsePtrOutput {
	return o
}

func (o LogAnalyticsConfigurationResponsePtrOutput) ToLogAnalyticsConfigurationResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationResponsePtrOutput {
	return o
}

func (o LogAnalyticsConfigurationResponsePtrOutput) Elem() LogAnalyticsConfigurationResponseOutput {
	return o.ApplyT(func(v *LogAnalyticsConfigurationResponse) LogAnalyticsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsConfigurationResponse
		return ret
	}).(LogAnalyticsConfigurationResponseOutput)
}

func (o LogAnalyticsConfigurationResponsePtrOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomerId
	}).(pulumi.StringPtrOutput)
}

type Login struct {
	AllowedExternalRedirectUrls   []string          `pulumi:"allowedExternalRedirectUrls"`
	CookieExpiration              *CookieExpiration `pulumi:"cookieExpiration"`
	Nonce                         *Nonce            `pulumi:"nonce"`
	PreserveUrlFragmentsForLogins *bool             `pulumi:"preserveUrlFragmentsForLogins"`
	Routes                        *LoginRoutes      `pulumi:"routes"`
}





type LoginInput interface {
	pulumi.Input

	ToLoginOutput() LoginOutput
	ToLoginOutputWithContext(context.Context) LoginOutput
}

type LoginArgs struct {
	AllowedExternalRedirectUrls   pulumi.StringArrayInput  `pulumi:"allowedExternalRedirectUrls"`
	CookieExpiration              CookieExpirationPtrInput `pulumi:"cookieExpiration"`
	Nonce                         NoncePtrInput            `pulumi:"nonce"`
	PreserveUrlFragmentsForLogins pulumi.BoolPtrInput      `pulumi:"preserveUrlFragmentsForLogins"`
	Routes                        LoginRoutesPtrInput      `pulumi:"routes"`
}

func (LoginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Login)(nil)).Elem()
}

func (i LoginArgs) ToLoginOutput() LoginOutput {
	return i.ToLoginOutputWithContext(context.Background())
}

func (i LoginArgs) ToLoginOutputWithContext(ctx context.Context) LoginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginOutput)
}

func (i LoginArgs) ToLoginPtrOutput() LoginPtrOutput {
	return i.ToLoginPtrOutputWithContext(context.Background())
}

func (i LoginArgs) ToLoginPtrOutputWithContext(ctx context.Context) LoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginOutput).ToLoginPtrOutputWithContext(ctx)
}









type LoginPtrInput interface {
	pulumi.Input

	ToLoginPtrOutput() LoginPtrOutput
	ToLoginPtrOutputWithContext(context.Context) LoginPtrOutput
}

type loginPtrType LoginArgs

func LoginPtr(v *LoginArgs) LoginPtrInput {
	return (*loginPtrType)(v)
}

func (*loginPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Login)(nil)).Elem()
}

func (i *loginPtrType) ToLoginPtrOutput() LoginPtrOutput {
	return i.ToLoginPtrOutputWithContext(context.Background())
}

func (i *loginPtrType) ToLoginPtrOutputWithContext(ctx context.Context) LoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginPtrOutput)
}

type LoginOutput struct{ *pulumi.OutputState }

func (LoginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Login)(nil)).Elem()
}

func (o LoginOutput) ToLoginOutput() LoginOutput {
	return o
}

func (o LoginOutput) ToLoginOutputWithContext(ctx context.Context) LoginOutput {
	return o
}

func (o LoginOutput) ToLoginPtrOutput() LoginPtrOutput {
	return o.ToLoginPtrOutputWithContext(context.Background())
}

func (o LoginOutput) ToLoginPtrOutputWithContext(ctx context.Context) LoginPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Login) *Login {
		return &v
	}).(LoginPtrOutput)
}

func (o LoginOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Login) []string { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

func (o LoginOutput) CookieExpiration() CookieExpirationPtrOutput {
	return o.ApplyT(func(v Login) *CookieExpiration { return v.CookieExpiration }).(CookieExpirationPtrOutput)
}

func (o LoginOutput) Nonce() NoncePtrOutput {
	return o.ApplyT(func(v Login) *Nonce { return v.Nonce }).(NoncePtrOutput)
}

func (o LoginOutput) PreserveUrlFragmentsForLogins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Login) *bool { return v.PreserveUrlFragmentsForLogins }).(pulumi.BoolPtrOutput)
}

func (o LoginOutput) Routes() LoginRoutesPtrOutput {
	return o.ApplyT(func(v Login) *LoginRoutes { return v.Routes }).(LoginRoutesPtrOutput)
}

type LoginPtrOutput struct{ *pulumi.OutputState }

func (LoginPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Login)(nil)).Elem()
}

func (o LoginPtrOutput) ToLoginPtrOutput() LoginPtrOutput {
	return o
}

func (o LoginPtrOutput) ToLoginPtrOutputWithContext(ctx context.Context) LoginPtrOutput {
	return o
}

func (o LoginPtrOutput) Elem() LoginOutput {
	return o.ApplyT(func(v *Login) Login {
		if v != nil {
			return *v
		}
		var ret Login
		return ret
	}).(LoginOutput)
}

func (o LoginPtrOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Login) []string {
		if v == nil {
			return nil
		}
		return v.AllowedExternalRedirectUrls
	}).(pulumi.StringArrayOutput)
}

func (o LoginPtrOutput) CookieExpiration() CookieExpirationPtrOutput {
	return o.ApplyT(func(v *Login) *CookieExpiration {
		if v == nil {
			return nil
		}
		return v.CookieExpiration
	}).(CookieExpirationPtrOutput)
}

func (o LoginPtrOutput) Nonce() NoncePtrOutput {
	return o.ApplyT(func(v *Login) *Nonce {
		if v == nil {
			return nil
		}
		return v.Nonce
	}).(NoncePtrOutput)
}

func (o LoginPtrOutput) PreserveUrlFragmentsForLogins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Login) *bool {
		if v == nil {
			return nil
		}
		return v.PreserveUrlFragmentsForLogins
	}).(pulumi.BoolPtrOutput)
}

func (o LoginPtrOutput) Routes() LoginRoutesPtrOutput {
	return o.ApplyT(func(v *Login) *LoginRoutes {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(LoginRoutesPtrOutput)
}

type LoginResponse struct {
	AllowedExternalRedirectUrls   []string                  `pulumi:"allowedExternalRedirectUrls"`
	CookieExpiration              *CookieExpirationResponse `pulumi:"cookieExpiration"`
	Nonce                         *NonceResponse            `pulumi:"nonce"`
	PreserveUrlFragmentsForLogins *bool                     `pulumi:"preserveUrlFragmentsForLogins"`
	Routes                        *LoginRoutesResponse      `pulumi:"routes"`
}

type LoginResponseOutput struct{ *pulumi.OutputState }

func (LoginResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginResponse)(nil)).Elem()
}

func (o LoginResponseOutput) ToLoginResponseOutput() LoginResponseOutput {
	return o
}

func (o LoginResponseOutput) ToLoginResponseOutputWithContext(ctx context.Context) LoginResponseOutput {
	return o
}

func (o LoginResponseOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoginResponse) []string { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

func (o LoginResponseOutput) CookieExpiration() CookieExpirationResponsePtrOutput {
	return o.ApplyT(func(v LoginResponse) *CookieExpirationResponse { return v.CookieExpiration }).(CookieExpirationResponsePtrOutput)
}

func (o LoginResponseOutput) Nonce() NonceResponsePtrOutput {
	return o.ApplyT(func(v LoginResponse) *NonceResponse { return v.Nonce }).(NonceResponsePtrOutput)
}

func (o LoginResponseOutput) PreserveUrlFragmentsForLogins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LoginResponse) *bool { return v.PreserveUrlFragmentsForLogins }).(pulumi.BoolPtrOutput)
}

func (o LoginResponseOutput) Routes() LoginRoutesResponsePtrOutput {
	return o.ApplyT(func(v LoginResponse) *LoginRoutesResponse { return v.Routes }).(LoginRoutesResponsePtrOutput)
}

type LoginResponsePtrOutput struct{ *pulumi.OutputState }

func (LoginResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginResponse)(nil)).Elem()
}

func (o LoginResponsePtrOutput) ToLoginResponsePtrOutput() LoginResponsePtrOutput {
	return o
}

func (o LoginResponsePtrOutput) ToLoginResponsePtrOutputWithContext(ctx context.Context) LoginResponsePtrOutput {
	return o
}

func (o LoginResponsePtrOutput) Elem() LoginResponseOutput {
	return o.ApplyT(func(v *LoginResponse) LoginResponse {
		if v != nil {
			return *v
		}
		var ret LoginResponse
		return ret
	}).(LoginResponseOutput)
}

func (o LoginResponsePtrOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoginResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedExternalRedirectUrls
	}).(pulumi.StringArrayOutput)
}

func (o LoginResponsePtrOutput) CookieExpiration() CookieExpirationResponsePtrOutput {
	return o.ApplyT(func(v *LoginResponse) *CookieExpirationResponse {
		if v == nil {
			return nil
		}
		return v.CookieExpiration
	}).(CookieExpirationResponsePtrOutput)
}

func (o LoginResponsePtrOutput) Nonce() NonceResponsePtrOutput {
	return o.ApplyT(func(v *LoginResponse) *NonceResponse {
		if v == nil {
			return nil
		}
		return v.Nonce
	}).(NonceResponsePtrOutput)
}

func (o LoginResponsePtrOutput) PreserveUrlFragmentsForLogins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoginResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PreserveUrlFragmentsForLogins
	}).(pulumi.BoolPtrOutput)
}

func (o LoginResponsePtrOutput) Routes() LoginRoutesResponsePtrOutput {
	return o.ApplyT(func(v *LoginResponse) *LoginRoutesResponse {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(LoginRoutesResponsePtrOutput)
}

type LoginRoutes struct {
	LogoutEndpoint *string `pulumi:"logoutEndpoint"`
}





type LoginRoutesInput interface {
	pulumi.Input

	ToLoginRoutesOutput() LoginRoutesOutput
	ToLoginRoutesOutputWithContext(context.Context) LoginRoutesOutput
}

type LoginRoutesArgs struct {
	LogoutEndpoint pulumi.StringPtrInput `pulumi:"logoutEndpoint"`
}

func (LoginRoutesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginRoutes)(nil)).Elem()
}

func (i LoginRoutesArgs) ToLoginRoutesOutput() LoginRoutesOutput {
	return i.ToLoginRoutesOutputWithContext(context.Background())
}

func (i LoginRoutesArgs) ToLoginRoutesOutputWithContext(ctx context.Context) LoginRoutesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginRoutesOutput)
}

func (i LoginRoutesArgs) ToLoginRoutesPtrOutput() LoginRoutesPtrOutput {
	return i.ToLoginRoutesPtrOutputWithContext(context.Background())
}

func (i LoginRoutesArgs) ToLoginRoutesPtrOutputWithContext(ctx context.Context) LoginRoutesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginRoutesOutput).ToLoginRoutesPtrOutputWithContext(ctx)
}









type LoginRoutesPtrInput interface {
	pulumi.Input

	ToLoginRoutesPtrOutput() LoginRoutesPtrOutput
	ToLoginRoutesPtrOutputWithContext(context.Context) LoginRoutesPtrOutput
}

type loginRoutesPtrType LoginRoutesArgs

func LoginRoutesPtr(v *LoginRoutesArgs) LoginRoutesPtrInput {
	return (*loginRoutesPtrType)(v)
}

func (*loginRoutesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginRoutes)(nil)).Elem()
}

func (i *loginRoutesPtrType) ToLoginRoutesPtrOutput() LoginRoutesPtrOutput {
	return i.ToLoginRoutesPtrOutputWithContext(context.Background())
}

func (i *loginRoutesPtrType) ToLoginRoutesPtrOutputWithContext(ctx context.Context) LoginRoutesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginRoutesPtrOutput)
}

type LoginRoutesOutput struct{ *pulumi.OutputState }

func (LoginRoutesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginRoutes)(nil)).Elem()
}

func (o LoginRoutesOutput) ToLoginRoutesOutput() LoginRoutesOutput {
	return o
}

func (o LoginRoutesOutput) ToLoginRoutesOutputWithContext(ctx context.Context) LoginRoutesOutput {
	return o
}

func (o LoginRoutesOutput) ToLoginRoutesPtrOutput() LoginRoutesPtrOutput {
	return o.ToLoginRoutesPtrOutputWithContext(context.Background())
}

func (o LoginRoutesOutput) ToLoginRoutesPtrOutputWithContext(ctx context.Context) LoginRoutesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoginRoutes) *LoginRoutes {
		return &v
	}).(LoginRoutesPtrOutput)
}

func (o LoginRoutesOutput) LogoutEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoginRoutes) *string { return v.LogoutEndpoint }).(pulumi.StringPtrOutput)
}

type LoginRoutesPtrOutput struct{ *pulumi.OutputState }

func (LoginRoutesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginRoutes)(nil)).Elem()
}

func (o LoginRoutesPtrOutput) ToLoginRoutesPtrOutput() LoginRoutesPtrOutput {
	return o
}

func (o LoginRoutesPtrOutput) ToLoginRoutesPtrOutputWithContext(ctx context.Context) LoginRoutesPtrOutput {
	return o
}

func (o LoginRoutesPtrOutput) Elem() LoginRoutesOutput {
	return o.ApplyT(func(v *LoginRoutes) LoginRoutes {
		if v != nil {
			return *v
		}
		var ret LoginRoutes
		return ret
	}).(LoginRoutesOutput)
}

func (o LoginRoutesPtrOutput) LogoutEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoginRoutes) *string {
		if v == nil {
			return nil
		}
		return v.LogoutEndpoint
	}).(pulumi.StringPtrOutput)
}

type LoginRoutesResponse struct {
	LogoutEndpoint *string `pulumi:"logoutEndpoint"`
}

type LoginRoutesResponseOutput struct{ *pulumi.OutputState }

func (LoginRoutesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginRoutesResponse)(nil)).Elem()
}

func (o LoginRoutesResponseOutput) ToLoginRoutesResponseOutput() LoginRoutesResponseOutput {
	return o
}

func (o LoginRoutesResponseOutput) ToLoginRoutesResponseOutputWithContext(ctx context.Context) LoginRoutesResponseOutput {
	return o
}

func (o LoginRoutesResponseOutput) LogoutEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoginRoutesResponse) *string { return v.LogoutEndpoint }).(pulumi.StringPtrOutput)
}

type LoginRoutesResponsePtrOutput struct{ *pulumi.OutputState }

func (LoginRoutesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginRoutesResponse)(nil)).Elem()
}

func (o LoginRoutesResponsePtrOutput) ToLoginRoutesResponsePtrOutput() LoginRoutesResponsePtrOutput {
	return o
}

func (o LoginRoutesResponsePtrOutput) ToLoginRoutesResponsePtrOutputWithContext(ctx context.Context) LoginRoutesResponsePtrOutput {
	return o
}

func (o LoginRoutesResponsePtrOutput) Elem() LoginRoutesResponseOutput {
	return o.ApplyT(func(v *LoginRoutesResponse) LoginRoutesResponse {
		if v != nil {
			return *v
		}
		var ret LoginRoutesResponse
		return ret
	}).(LoginRoutesResponseOutput)
}

func (o LoginRoutesResponsePtrOutput) LogoutEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoginRoutesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LogoutEndpoint
	}).(pulumi.StringPtrOutput)
}

type LoginScopes struct {
	Scopes []string `pulumi:"scopes"`
}





type LoginScopesInput interface {
	pulumi.Input

	ToLoginScopesOutput() LoginScopesOutput
	ToLoginScopesOutputWithContext(context.Context) LoginScopesOutput
}

type LoginScopesArgs struct {
	Scopes pulumi.StringArrayInput `pulumi:"scopes"`
}

func (LoginScopesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginScopes)(nil)).Elem()
}

func (i LoginScopesArgs) ToLoginScopesOutput() LoginScopesOutput {
	return i.ToLoginScopesOutputWithContext(context.Background())
}

func (i LoginScopesArgs) ToLoginScopesOutputWithContext(ctx context.Context) LoginScopesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginScopesOutput)
}

func (i LoginScopesArgs) ToLoginScopesPtrOutput() LoginScopesPtrOutput {
	return i.ToLoginScopesPtrOutputWithContext(context.Background())
}

func (i LoginScopesArgs) ToLoginScopesPtrOutputWithContext(ctx context.Context) LoginScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginScopesOutput).ToLoginScopesPtrOutputWithContext(ctx)
}









type LoginScopesPtrInput interface {
	pulumi.Input

	ToLoginScopesPtrOutput() LoginScopesPtrOutput
	ToLoginScopesPtrOutputWithContext(context.Context) LoginScopesPtrOutput
}

type loginScopesPtrType LoginScopesArgs

func LoginScopesPtr(v *LoginScopesArgs) LoginScopesPtrInput {
	return (*loginScopesPtrType)(v)
}

func (*loginScopesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginScopes)(nil)).Elem()
}

func (i *loginScopesPtrType) ToLoginScopesPtrOutput() LoginScopesPtrOutput {
	return i.ToLoginScopesPtrOutputWithContext(context.Background())
}

func (i *loginScopesPtrType) ToLoginScopesPtrOutputWithContext(ctx context.Context) LoginScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginScopesPtrOutput)
}

type LoginScopesOutput struct{ *pulumi.OutputState }

func (LoginScopesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginScopes)(nil)).Elem()
}

func (o LoginScopesOutput) ToLoginScopesOutput() LoginScopesOutput {
	return o
}

func (o LoginScopesOutput) ToLoginScopesOutputWithContext(ctx context.Context) LoginScopesOutput {
	return o
}

func (o LoginScopesOutput) ToLoginScopesPtrOutput() LoginScopesPtrOutput {
	return o.ToLoginScopesPtrOutputWithContext(context.Background())
}

func (o LoginScopesOutput) ToLoginScopesPtrOutputWithContext(ctx context.Context) LoginScopesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoginScopes) *LoginScopes {
		return &v
	}).(LoginScopesPtrOutput)
}

func (o LoginScopesOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoginScopes) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type LoginScopesPtrOutput struct{ *pulumi.OutputState }

func (LoginScopesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginScopes)(nil)).Elem()
}

func (o LoginScopesPtrOutput) ToLoginScopesPtrOutput() LoginScopesPtrOutput {
	return o
}

func (o LoginScopesPtrOutput) ToLoginScopesPtrOutputWithContext(ctx context.Context) LoginScopesPtrOutput {
	return o
}

func (o LoginScopesPtrOutput) Elem() LoginScopesOutput {
	return o.ApplyT(func(v *LoginScopes) LoginScopes {
		if v != nil {
			return *v
		}
		var ret LoginScopes
		return ret
	}).(LoginScopesOutput)
}

func (o LoginScopesPtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoginScopes) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type LoginScopesResponse struct {
	Scopes []string `pulumi:"scopes"`
}

type LoginScopesResponseOutput struct{ *pulumi.OutputState }

func (LoginScopesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginScopesResponse)(nil)).Elem()
}

func (o LoginScopesResponseOutput) ToLoginScopesResponseOutput() LoginScopesResponseOutput {
	return o
}

func (o LoginScopesResponseOutput) ToLoginScopesResponseOutputWithContext(ctx context.Context) LoginScopesResponseOutput {
	return o
}

func (o LoginScopesResponseOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoginScopesResponse) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type LoginScopesResponsePtrOutput struct{ *pulumi.OutputState }

func (LoginScopesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginScopesResponse)(nil)).Elem()
}

func (o LoginScopesResponsePtrOutput) ToLoginScopesResponsePtrOutput() LoginScopesResponsePtrOutput {
	return o
}

func (o LoginScopesResponsePtrOutput) ToLoginScopesResponsePtrOutputWithContext(ctx context.Context) LoginScopesResponsePtrOutput {
	return o
}

func (o LoginScopesResponsePtrOutput) Elem() LoginScopesResponseOutput {
	return o.ApplyT(func(v *LoginScopesResponse) LoginScopesResponse {
		if v != nil {
			return *v
		}
		var ret LoginScopesResponse
		return ret
	}).(LoginScopesResponseOutput)
}

func (o LoginScopesResponsePtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoginScopesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type ManagedEnvironmentStorageProperties struct {
	AzureFile *AzureFileProperties `pulumi:"azureFile"`
}





type ManagedEnvironmentStoragePropertiesInput interface {
	pulumi.Input

	ToManagedEnvironmentStoragePropertiesOutput() ManagedEnvironmentStoragePropertiesOutput
	ToManagedEnvironmentStoragePropertiesOutputWithContext(context.Context) ManagedEnvironmentStoragePropertiesOutput
}

type ManagedEnvironmentStoragePropertiesArgs struct {
	AzureFile AzureFilePropertiesPtrInput `pulumi:"azureFile"`
}

func (ManagedEnvironmentStoragePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedEnvironmentStorageProperties)(nil)).Elem()
}

func (i ManagedEnvironmentStoragePropertiesArgs) ToManagedEnvironmentStoragePropertiesOutput() ManagedEnvironmentStoragePropertiesOutput {
	return i.ToManagedEnvironmentStoragePropertiesOutputWithContext(context.Background())
}

func (i ManagedEnvironmentStoragePropertiesArgs) ToManagedEnvironmentStoragePropertiesOutputWithContext(ctx context.Context) ManagedEnvironmentStoragePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedEnvironmentStoragePropertiesOutput)
}

func (i ManagedEnvironmentStoragePropertiesArgs) ToManagedEnvironmentStoragePropertiesPtrOutput() ManagedEnvironmentStoragePropertiesPtrOutput {
	return i.ToManagedEnvironmentStoragePropertiesPtrOutputWithContext(context.Background())
}

func (i ManagedEnvironmentStoragePropertiesArgs) ToManagedEnvironmentStoragePropertiesPtrOutputWithContext(ctx context.Context) ManagedEnvironmentStoragePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedEnvironmentStoragePropertiesOutput).ToManagedEnvironmentStoragePropertiesPtrOutputWithContext(ctx)
}









type ManagedEnvironmentStoragePropertiesPtrInput interface {
	pulumi.Input

	ToManagedEnvironmentStoragePropertiesPtrOutput() ManagedEnvironmentStoragePropertiesPtrOutput
	ToManagedEnvironmentStoragePropertiesPtrOutputWithContext(context.Context) ManagedEnvironmentStoragePropertiesPtrOutput
}

type managedEnvironmentStoragePropertiesPtrType ManagedEnvironmentStoragePropertiesArgs

func ManagedEnvironmentStoragePropertiesPtr(v *ManagedEnvironmentStoragePropertiesArgs) ManagedEnvironmentStoragePropertiesPtrInput {
	return (*managedEnvironmentStoragePropertiesPtrType)(v)
}

func (*managedEnvironmentStoragePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedEnvironmentStorageProperties)(nil)).Elem()
}

func (i *managedEnvironmentStoragePropertiesPtrType) ToManagedEnvironmentStoragePropertiesPtrOutput() ManagedEnvironmentStoragePropertiesPtrOutput {
	return i.ToManagedEnvironmentStoragePropertiesPtrOutputWithContext(context.Background())
}

func (i *managedEnvironmentStoragePropertiesPtrType) ToManagedEnvironmentStoragePropertiesPtrOutputWithContext(ctx context.Context) ManagedEnvironmentStoragePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedEnvironmentStoragePropertiesPtrOutput)
}

type ManagedEnvironmentStoragePropertiesOutput struct{ *pulumi.OutputState }

func (ManagedEnvironmentStoragePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedEnvironmentStorageProperties)(nil)).Elem()
}

func (o ManagedEnvironmentStoragePropertiesOutput) ToManagedEnvironmentStoragePropertiesOutput() ManagedEnvironmentStoragePropertiesOutput {
	return o
}

func (o ManagedEnvironmentStoragePropertiesOutput) ToManagedEnvironmentStoragePropertiesOutputWithContext(ctx context.Context) ManagedEnvironmentStoragePropertiesOutput {
	return o
}

func (o ManagedEnvironmentStoragePropertiesOutput) ToManagedEnvironmentStoragePropertiesPtrOutput() ManagedEnvironmentStoragePropertiesPtrOutput {
	return o.ToManagedEnvironmentStoragePropertiesPtrOutputWithContext(context.Background())
}

func (o ManagedEnvironmentStoragePropertiesOutput) ToManagedEnvironmentStoragePropertiesPtrOutputWithContext(ctx context.Context) ManagedEnvironmentStoragePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedEnvironmentStorageProperties) *ManagedEnvironmentStorageProperties {
		return &v
	}).(ManagedEnvironmentStoragePropertiesPtrOutput)
}

func (o ManagedEnvironmentStoragePropertiesOutput) AzureFile() AzureFilePropertiesPtrOutput {
	return o.ApplyT(func(v ManagedEnvironmentStorageProperties) *AzureFileProperties { return v.AzureFile }).(AzureFilePropertiesPtrOutput)
}

type ManagedEnvironmentStoragePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ManagedEnvironmentStoragePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedEnvironmentStorageProperties)(nil)).Elem()
}

func (o ManagedEnvironmentStoragePropertiesPtrOutput) ToManagedEnvironmentStoragePropertiesPtrOutput() ManagedEnvironmentStoragePropertiesPtrOutput {
	return o
}

func (o ManagedEnvironmentStoragePropertiesPtrOutput) ToManagedEnvironmentStoragePropertiesPtrOutputWithContext(ctx context.Context) ManagedEnvironmentStoragePropertiesPtrOutput {
	return o
}

func (o ManagedEnvironmentStoragePropertiesPtrOutput) Elem() ManagedEnvironmentStoragePropertiesOutput {
	return o.ApplyT(func(v *ManagedEnvironmentStorageProperties) ManagedEnvironmentStorageProperties {
		if v != nil {
			return *v
		}
		var ret ManagedEnvironmentStorageProperties
		return ret
	}).(ManagedEnvironmentStoragePropertiesOutput)
}

func (o ManagedEnvironmentStoragePropertiesPtrOutput) AzureFile() AzureFilePropertiesPtrOutput {
	return o.ApplyT(func(v *ManagedEnvironmentStorageProperties) *AzureFileProperties {
		if v == nil {
			return nil
		}
		return v.AzureFile
	}).(AzureFilePropertiesPtrOutput)
}

type ManagedEnvironmentStorageResponseProperties struct {
	AzureFile *AzureFilePropertiesResponse `pulumi:"azureFile"`
}

type ManagedEnvironmentStorageResponsePropertiesOutput struct{ *pulumi.OutputState }

func (ManagedEnvironmentStorageResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedEnvironmentStorageResponseProperties)(nil)).Elem()
}

func (o ManagedEnvironmentStorageResponsePropertiesOutput) ToManagedEnvironmentStorageResponsePropertiesOutput() ManagedEnvironmentStorageResponsePropertiesOutput {
	return o
}

func (o ManagedEnvironmentStorageResponsePropertiesOutput) ToManagedEnvironmentStorageResponsePropertiesOutputWithContext(ctx context.Context) ManagedEnvironmentStorageResponsePropertiesOutput {
	return o
}

func (o ManagedEnvironmentStorageResponsePropertiesOutput) AzureFile() AzureFilePropertiesResponsePtrOutput {
	return o.ApplyT(func(v ManagedEnvironmentStorageResponseProperties) *AzureFilePropertiesResponse { return v.AzureFile }).(AzureFilePropertiesResponsePtrOutput)
}

type ManagedServiceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type Nonce struct {
	NonceExpirationInterval *string `pulumi:"nonceExpirationInterval"`
	ValidateNonce           *bool   `pulumi:"validateNonce"`
}





type NonceInput interface {
	pulumi.Input

	ToNonceOutput() NonceOutput
	ToNonceOutputWithContext(context.Context) NonceOutput
}

type NonceArgs struct {
	NonceExpirationInterval pulumi.StringPtrInput `pulumi:"nonceExpirationInterval"`
	ValidateNonce           pulumi.BoolPtrInput   `pulumi:"validateNonce"`
}

func (NonceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Nonce)(nil)).Elem()
}

func (i NonceArgs) ToNonceOutput() NonceOutput {
	return i.ToNonceOutputWithContext(context.Background())
}

func (i NonceArgs) ToNonceOutputWithContext(ctx context.Context) NonceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonceOutput)
}

func (i NonceArgs) ToNoncePtrOutput() NoncePtrOutput {
	return i.ToNoncePtrOutputWithContext(context.Background())
}

func (i NonceArgs) ToNoncePtrOutputWithContext(ctx context.Context) NoncePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonceOutput).ToNoncePtrOutputWithContext(ctx)
}









type NoncePtrInput interface {
	pulumi.Input

	ToNoncePtrOutput() NoncePtrOutput
	ToNoncePtrOutputWithContext(context.Context) NoncePtrOutput
}

type noncePtrType NonceArgs

func NoncePtr(v *NonceArgs) NoncePtrInput {
	return (*noncePtrType)(v)
}

func (*noncePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Nonce)(nil)).Elem()
}

func (i *noncePtrType) ToNoncePtrOutput() NoncePtrOutput {
	return i.ToNoncePtrOutputWithContext(context.Background())
}

func (i *noncePtrType) ToNoncePtrOutputWithContext(ctx context.Context) NoncePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoncePtrOutput)
}

type NonceOutput struct{ *pulumi.OutputState }

func (NonceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Nonce)(nil)).Elem()
}

func (o NonceOutput) ToNonceOutput() NonceOutput {
	return o
}

func (o NonceOutput) ToNonceOutputWithContext(ctx context.Context) NonceOutput {
	return o
}

func (o NonceOutput) ToNoncePtrOutput() NoncePtrOutput {
	return o.ToNoncePtrOutputWithContext(context.Background())
}

func (o NonceOutput) ToNoncePtrOutputWithContext(ctx context.Context) NoncePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Nonce) *Nonce {
		return &v
	}).(NoncePtrOutput)
}

func (o NonceOutput) NonceExpirationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Nonce) *string { return v.NonceExpirationInterval }).(pulumi.StringPtrOutput)
}

func (o NonceOutput) ValidateNonce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Nonce) *bool { return v.ValidateNonce }).(pulumi.BoolPtrOutput)
}

type NoncePtrOutput struct{ *pulumi.OutputState }

func (NoncePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nonce)(nil)).Elem()
}

func (o NoncePtrOutput) ToNoncePtrOutput() NoncePtrOutput {
	return o
}

func (o NoncePtrOutput) ToNoncePtrOutputWithContext(ctx context.Context) NoncePtrOutput {
	return o
}

func (o NoncePtrOutput) Elem() NonceOutput {
	return o.ApplyT(func(v *Nonce) Nonce {
		if v != nil {
			return *v
		}
		var ret Nonce
		return ret
	}).(NonceOutput)
}

func (o NoncePtrOutput) NonceExpirationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nonce) *string {
		if v == nil {
			return nil
		}
		return v.NonceExpirationInterval
	}).(pulumi.StringPtrOutput)
}

func (o NoncePtrOutput) ValidateNonce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Nonce) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateNonce
	}).(pulumi.BoolPtrOutput)
}

type NonceResponse struct {
	NonceExpirationInterval *string `pulumi:"nonceExpirationInterval"`
	ValidateNonce           *bool   `pulumi:"validateNonce"`
}

type NonceResponseOutput struct{ *pulumi.OutputState }

func (NonceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NonceResponse)(nil)).Elem()
}

func (o NonceResponseOutput) ToNonceResponseOutput() NonceResponseOutput {
	return o
}

func (o NonceResponseOutput) ToNonceResponseOutputWithContext(ctx context.Context) NonceResponseOutput {
	return o
}

func (o NonceResponseOutput) NonceExpirationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NonceResponse) *string { return v.NonceExpirationInterval }).(pulumi.StringPtrOutput)
}

func (o NonceResponseOutput) ValidateNonce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NonceResponse) *bool { return v.ValidateNonce }).(pulumi.BoolPtrOutput)
}

type NonceResponsePtrOutput struct{ *pulumi.OutputState }

func (NonceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NonceResponse)(nil)).Elem()
}

func (o NonceResponsePtrOutput) ToNonceResponsePtrOutput() NonceResponsePtrOutput {
	return o
}

func (o NonceResponsePtrOutput) ToNonceResponsePtrOutputWithContext(ctx context.Context) NonceResponsePtrOutput {
	return o
}

func (o NonceResponsePtrOutput) Elem() NonceResponseOutput {
	return o.ApplyT(func(v *NonceResponse) NonceResponse {
		if v != nil {
			return *v
		}
		var ret NonceResponse
		return ret
	}).(NonceResponseOutput)
}

func (o NonceResponsePtrOutput) NonceExpirationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NonceResponse) *string {
		if v == nil {
			return nil
		}
		return v.NonceExpirationInterval
	}).(pulumi.StringPtrOutput)
}

func (o NonceResponsePtrOutput) ValidateNonce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NonceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateNonce
	}).(pulumi.BoolPtrOutput)
}

type OpenIdConnectClientCredential struct {
	ClientSecretSettingName *string                 `pulumi:"clientSecretSettingName"`
	Method                  *ClientCredentialMethod `pulumi:"method"`
}





type OpenIdConnectClientCredentialInput interface {
	pulumi.Input

	ToOpenIdConnectClientCredentialOutput() OpenIdConnectClientCredentialOutput
	ToOpenIdConnectClientCredentialOutputWithContext(context.Context) OpenIdConnectClientCredentialOutput
}

type OpenIdConnectClientCredentialArgs struct {
	ClientSecretSettingName pulumi.StringPtrInput          `pulumi:"clientSecretSettingName"`
	Method                  ClientCredentialMethodPtrInput `pulumi:"method"`
}

func (OpenIdConnectClientCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectClientCredential)(nil)).Elem()
}

func (i OpenIdConnectClientCredentialArgs) ToOpenIdConnectClientCredentialOutput() OpenIdConnectClientCredentialOutput {
	return i.ToOpenIdConnectClientCredentialOutputWithContext(context.Background())
}

func (i OpenIdConnectClientCredentialArgs) ToOpenIdConnectClientCredentialOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectClientCredentialOutput)
}

func (i OpenIdConnectClientCredentialArgs) ToOpenIdConnectClientCredentialPtrOutput() OpenIdConnectClientCredentialPtrOutput {
	return i.ToOpenIdConnectClientCredentialPtrOutputWithContext(context.Background())
}

func (i OpenIdConnectClientCredentialArgs) ToOpenIdConnectClientCredentialPtrOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectClientCredentialOutput).ToOpenIdConnectClientCredentialPtrOutputWithContext(ctx)
}









type OpenIdConnectClientCredentialPtrInput interface {
	pulumi.Input

	ToOpenIdConnectClientCredentialPtrOutput() OpenIdConnectClientCredentialPtrOutput
	ToOpenIdConnectClientCredentialPtrOutputWithContext(context.Context) OpenIdConnectClientCredentialPtrOutput
}

type openIdConnectClientCredentialPtrType OpenIdConnectClientCredentialArgs

func OpenIdConnectClientCredentialPtr(v *OpenIdConnectClientCredentialArgs) OpenIdConnectClientCredentialPtrInput {
	return (*openIdConnectClientCredentialPtrType)(v)
}

func (*openIdConnectClientCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectClientCredential)(nil)).Elem()
}

func (i *openIdConnectClientCredentialPtrType) ToOpenIdConnectClientCredentialPtrOutput() OpenIdConnectClientCredentialPtrOutput {
	return i.ToOpenIdConnectClientCredentialPtrOutputWithContext(context.Background())
}

func (i *openIdConnectClientCredentialPtrType) ToOpenIdConnectClientCredentialPtrOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectClientCredentialPtrOutput)
}

type OpenIdConnectClientCredentialOutput struct{ *pulumi.OutputState }

func (OpenIdConnectClientCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectClientCredential)(nil)).Elem()
}

func (o OpenIdConnectClientCredentialOutput) ToOpenIdConnectClientCredentialOutput() OpenIdConnectClientCredentialOutput {
	return o
}

func (o OpenIdConnectClientCredentialOutput) ToOpenIdConnectClientCredentialOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialOutput {
	return o
}

func (o OpenIdConnectClientCredentialOutput) ToOpenIdConnectClientCredentialPtrOutput() OpenIdConnectClientCredentialPtrOutput {
	return o.ToOpenIdConnectClientCredentialPtrOutputWithContext(context.Background())
}

func (o OpenIdConnectClientCredentialOutput) ToOpenIdConnectClientCredentialPtrOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenIdConnectClientCredential) *OpenIdConnectClientCredential {
		return &v
	}).(OpenIdConnectClientCredentialPtrOutput)
}

func (o OpenIdConnectClientCredentialOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectClientCredential) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectClientCredentialOutput) Method() ClientCredentialMethodPtrOutput {
	return o.ApplyT(func(v OpenIdConnectClientCredential) *ClientCredentialMethod { return v.Method }).(ClientCredentialMethodPtrOutput)
}

type OpenIdConnectClientCredentialPtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectClientCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectClientCredential)(nil)).Elem()
}

func (o OpenIdConnectClientCredentialPtrOutput) ToOpenIdConnectClientCredentialPtrOutput() OpenIdConnectClientCredentialPtrOutput {
	return o
}

func (o OpenIdConnectClientCredentialPtrOutput) ToOpenIdConnectClientCredentialPtrOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialPtrOutput {
	return o
}

func (o OpenIdConnectClientCredentialPtrOutput) Elem() OpenIdConnectClientCredentialOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredential) OpenIdConnectClientCredential {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectClientCredential
		return ret
	}).(OpenIdConnectClientCredentialOutput)
}

func (o OpenIdConnectClientCredentialPtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredential) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectClientCredentialPtrOutput) Method() ClientCredentialMethodPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredential) *ClientCredentialMethod {
		if v == nil {
			return nil
		}
		return v.Method
	}).(ClientCredentialMethodPtrOutput)
}

type OpenIdConnectClientCredentialResponse struct {
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
	Method                  *string `pulumi:"method"`
}

type OpenIdConnectClientCredentialResponseOutput struct{ *pulumi.OutputState }

func (OpenIdConnectClientCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectClientCredentialResponse)(nil)).Elem()
}

func (o OpenIdConnectClientCredentialResponseOutput) ToOpenIdConnectClientCredentialResponseOutput() OpenIdConnectClientCredentialResponseOutput {
	return o
}

func (o OpenIdConnectClientCredentialResponseOutput) ToOpenIdConnectClientCredentialResponseOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialResponseOutput {
	return o
}

func (o OpenIdConnectClientCredentialResponseOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectClientCredentialResponse) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectClientCredentialResponseOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectClientCredentialResponse) *string { return v.Method }).(pulumi.StringPtrOutput)
}

type OpenIdConnectClientCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectClientCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectClientCredentialResponse)(nil)).Elem()
}

func (o OpenIdConnectClientCredentialResponsePtrOutput) ToOpenIdConnectClientCredentialResponsePtrOutput() OpenIdConnectClientCredentialResponsePtrOutput {
	return o
}

func (o OpenIdConnectClientCredentialResponsePtrOutput) ToOpenIdConnectClientCredentialResponsePtrOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialResponsePtrOutput {
	return o
}

func (o OpenIdConnectClientCredentialResponsePtrOutput) Elem() OpenIdConnectClientCredentialResponseOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredentialResponse) OpenIdConnectClientCredentialResponse {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectClientCredentialResponse
		return ret
	}).(OpenIdConnectClientCredentialResponseOutput)
}

func (o OpenIdConnectClientCredentialResponsePtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectClientCredentialResponsePtrOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Method
	}).(pulumi.StringPtrOutput)
}

type OpenIdConnectConfig struct {
	AuthorizationEndpoint        *string `pulumi:"authorizationEndpoint"`
	CertificationUri             *string `pulumi:"certificationUri"`
	Issuer                       *string `pulumi:"issuer"`
	TokenEndpoint                *string `pulumi:"tokenEndpoint"`
	WellKnownOpenIdConfiguration *string `pulumi:"wellKnownOpenIdConfiguration"`
}





type OpenIdConnectConfigInput interface {
	pulumi.Input

	ToOpenIdConnectConfigOutput() OpenIdConnectConfigOutput
	ToOpenIdConnectConfigOutputWithContext(context.Context) OpenIdConnectConfigOutput
}

type OpenIdConnectConfigArgs struct {
	AuthorizationEndpoint        pulumi.StringPtrInput `pulumi:"authorizationEndpoint"`
	CertificationUri             pulumi.StringPtrInput `pulumi:"certificationUri"`
	Issuer                       pulumi.StringPtrInput `pulumi:"issuer"`
	TokenEndpoint                pulumi.StringPtrInput `pulumi:"tokenEndpoint"`
	WellKnownOpenIdConfiguration pulumi.StringPtrInput `pulumi:"wellKnownOpenIdConfiguration"`
}

func (OpenIdConnectConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectConfig)(nil)).Elem()
}

func (i OpenIdConnectConfigArgs) ToOpenIdConnectConfigOutput() OpenIdConnectConfigOutput {
	return i.ToOpenIdConnectConfigOutputWithContext(context.Background())
}

func (i OpenIdConnectConfigArgs) ToOpenIdConnectConfigOutputWithContext(ctx context.Context) OpenIdConnectConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectConfigOutput)
}

func (i OpenIdConnectConfigArgs) ToOpenIdConnectConfigPtrOutput() OpenIdConnectConfigPtrOutput {
	return i.ToOpenIdConnectConfigPtrOutputWithContext(context.Background())
}

func (i OpenIdConnectConfigArgs) ToOpenIdConnectConfigPtrOutputWithContext(ctx context.Context) OpenIdConnectConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectConfigOutput).ToOpenIdConnectConfigPtrOutputWithContext(ctx)
}









type OpenIdConnectConfigPtrInput interface {
	pulumi.Input

	ToOpenIdConnectConfigPtrOutput() OpenIdConnectConfigPtrOutput
	ToOpenIdConnectConfigPtrOutputWithContext(context.Context) OpenIdConnectConfigPtrOutput
}

type openIdConnectConfigPtrType OpenIdConnectConfigArgs

func OpenIdConnectConfigPtr(v *OpenIdConnectConfigArgs) OpenIdConnectConfigPtrInput {
	return (*openIdConnectConfigPtrType)(v)
}

func (*openIdConnectConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectConfig)(nil)).Elem()
}

func (i *openIdConnectConfigPtrType) ToOpenIdConnectConfigPtrOutput() OpenIdConnectConfigPtrOutput {
	return i.ToOpenIdConnectConfigPtrOutputWithContext(context.Background())
}

func (i *openIdConnectConfigPtrType) ToOpenIdConnectConfigPtrOutputWithContext(ctx context.Context) OpenIdConnectConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectConfigPtrOutput)
}

type OpenIdConnectConfigOutput struct{ *pulumi.OutputState }

func (OpenIdConnectConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectConfig)(nil)).Elem()
}

func (o OpenIdConnectConfigOutput) ToOpenIdConnectConfigOutput() OpenIdConnectConfigOutput {
	return o
}

func (o OpenIdConnectConfigOutput) ToOpenIdConnectConfigOutputWithContext(ctx context.Context) OpenIdConnectConfigOutput {
	return o
}

func (o OpenIdConnectConfigOutput) ToOpenIdConnectConfigPtrOutput() OpenIdConnectConfigPtrOutput {
	return o.ToOpenIdConnectConfigPtrOutputWithContext(context.Background())
}

func (o OpenIdConnectConfigOutput) ToOpenIdConnectConfigPtrOutputWithContext(ctx context.Context) OpenIdConnectConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenIdConnectConfig) *OpenIdConnectConfig {
		return &v
	}).(OpenIdConnectConfigPtrOutput)
}

func (o OpenIdConnectConfigOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfig) *string { return v.AuthorizationEndpoint }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigOutput) CertificationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfig) *string { return v.CertificationUri }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfig) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfig) *string { return v.TokenEndpoint }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigOutput) WellKnownOpenIdConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfig) *string { return v.WellKnownOpenIdConfiguration }).(pulumi.StringPtrOutput)
}

type OpenIdConnectConfigPtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectConfig)(nil)).Elem()
}

func (o OpenIdConnectConfigPtrOutput) ToOpenIdConnectConfigPtrOutput() OpenIdConnectConfigPtrOutput {
	return o
}

func (o OpenIdConnectConfigPtrOutput) ToOpenIdConnectConfigPtrOutputWithContext(ctx context.Context) OpenIdConnectConfigPtrOutput {
	return o
}

func (o OpenIdConnectConfigPtrOutput) Elem() OpenIdConnectConfigOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) OpenIdConnectConfig {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectConfig
		return ret
	}).(OpenIdConnectConfigOutput)
}

func (o OpenIdConnectConfigPtrOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigPtrOutput) CertificationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) *string {
		if v == nil {
			return nil
		}
		return v.CertificationUri
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigPtrOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) *string {
		if v == nil {
			return nil
		}
		return v.Issuer
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigPtrOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) *string {
		if v == nil {
			return nil
		}
		return v.TokenEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigPtrOutput) WellKnownOpenIdConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) *string {
		if v == nil {
			return nil
		}
		return v.WellKnownOpenIdConfiguration
	}).(pulumi.StringPtrOutput)
}

type OpenIdConnectConfigResponse struct {
	AuthorizationEndpoint        *string `pulumi:"authorizationEndpoint"`
	CertificationUri             *string `pulumi:"certificationUri"`
	Issuer                       *string `pulumi:"issuer"`
	TokenEndpoint                *string `pulumi:"tokenEndpoint"`
	WellKnownOpenIdConfiguration *string `pulumi:"wellKnownOpenIdConfiguration"`
}

type OpenIdConnectConfigResponseOutput struct{ *pulumi.OutputState }

func (OpenIdConnectConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectConfigResponse)(nil)).Elem()
}

func (o OpenIdConnectConfigResponseOutput) ToOpenIdConnectConfigResponseOutput() OpenIdConnectConfigResponseOutput {
	return o
}

func (o OpenIdConnectConfigResponseOutput) ToOpenIdConnectConfigResponseOutputWithContext(ctx context.Context) OpenIdConnectConfigResponseOutput {
	return o
}

func (o OpenIdConnectConfigResponseOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfigResponse) *string { return v.AuthorizationEndpoint }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponseOutput) CertificationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfigResponse) *string { return v.CertificationUri }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponseOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfigResponse) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponseOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfigResponse) *string { return v.TokenEndpoint }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponseOutput) WellKnownOpenIdConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfigResponse) *string { return v.WellKnownOpenIdConfiguration }).(pulumi.StringPtrOutput)
}

type OpenIdConnectConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectConfigResponse)(nil)).Elem()
}

func (o OpenIdConnectConfigResponsePtrOutput) ToOpenIdConnectConfigResponsePtrOutput() OpenIdConnectConfigResponsePtrOutput {
	return o
}

func (o OpenIdConnectConfigResponsePtrOutput) ToOpenIdConnectConfigResponsePtrOutputWithContext(ctx context.Context) OpenIdConnectConfigResponsePtrOutput {
	return o
}

func (o OpenIdConnectConfigResponsePtrOutput) Elem() OpenIdConnectConfigResponseOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) OpenIdConnectConfigResponse {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectConfigResponse
		return ret
	}).(OpenIdConnectConfigResponseOutput)
}

func (o OpenIdConnectConfigResponsePtrOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponsePtrOutput) CertificationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.CertificationUri
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponsePtrOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Issuer
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponsePtrOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.TokenEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponsePtrOutput) WellKnownOpenIdConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.WellKnownOpenIdConfiguration
	}).(pulumi.StringPtrOutput)
}

type OpenIdConnectLogin struct {
	NameClaimType *string  `pulumi:"nameClaimType"`
	Scopes        []string `pulumi:"scopes"`
}





type OpenIdConnectLoginInput interface {
	pulumi.Input

	ToOpenIdConnectLoginOutput() OpenIdConnectLoginOutput
	ToOpenIdConnectLoginOutputWithContext(context.Context) OpenIdConnectLoginOutput
}

type OpenIdConnectLoginArgs struct {
	NameClaimType pulumi.StringPtrInput   `pulumi:"nameClaimType"`
	Scopes        pulumi.StringArrayInput `pulumi:"scopes"`
}

func (OpenIdConnectLoginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectLogin)(nil)).Elem()
}

func (i OpenIdConnectLoginArgs) ToOpenIdConnectLoginOutput() OpenIdConnectLoginOutput {
	return i.ToOpenIdConnectLoginOutputWithContext(context.Background())
}

func (i OpenIdConnectLoginArgs) ToOpenIdConnectLoginOutputWithContext(ctx context.Context) OpenIdConnectLoginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectLoginOutput)
}

func (i OpenIdConnectLoginArgs) ToOpenIdConnectLoginPtrOutput() OpenIdConnectLoginPtrOutput {
	return i.ToOpenIdConnectLoginPtrOutputWithContext(context.Background())
}

func (i OpenIdConnectLoginArgs) ToOpenIdConnectLoginPtrOutputWithContext(ctx context.Context) OpenIdConnectLoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectLoginOutput).ToOpenIdConnectLoginPtrOutputWithContext(ctx)
}









type OpenIdConnectLoginPtrInput interface {
	pulumi.Input

	ToOpenIdConnectLoginPtrOutput() OpenIdConnectLoginPtrOutput
	ToOpenIdConnectLoginPtrOutputWithContext(context.Context) OpenIdConnectLoginPtrOutput
}

type openIdConnectLoginPtrType OpenIdConnectLoginArgs

func OpenIdConnectLoginPtr(v *OpenIdConnectLoginArgs) OpenIdConnectLoginPtrInput {
	return (*openIdConnectLoginPtrType)(v)
}

func (*openIdConnectLoginPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectLogin)(nil)).Elem()
}

func (i *openIdConnectLoginPtrType) ToOpenIdConnectLoginPtrOutput() OpenIdConnectLoginPtrOutput {
	return i.ToOpenIdConnectLoginPtrOutputWithContext(context.Background())
}

func (i *openIdConnectLoginPtrType) ToOpenIdConnectLoginPtrOutputWithContext(ctx context.Context) OpenIdConnectLoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectLoginPtrOutput)
}

type OpenIdConnectLoginOutput struct{ *pulumi.OutputState }

func (OpenIdConnectLoginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectLogin)(nil)).Elem()
}

func (o OpenIdConnectLoginOutput) ToOpenIdConnectLoginOutput() OpenIdConnectLoginOutput {
	return o
}

func (o OpenIdConnectLoginOutput) ToOpenIdConnectLoginOutputWithContext(ctx context.Context) OpenIdConnectLoginOutput {
	return o
}

func (o OpenIdConnectLoginOutput) ToOpenIdConnectLoginPtrOutput() OpenIdConnectLoginPtrOutput {
	return o.ToOpenIdConnectLoginPtrOutputWithContext(context.Background())
}

func (o OpenIdConnectLoginOutput) ToOpenIdConnectLoginPtrOutputWithContext(ctx context.Context) OpenIdConnectLoginPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenIdConnectLogin) *OpenIdConnectLogin {
		return &v
	}).(OpenIdConnectLoginPtrOutput)
}

func (o OpenIdConnectLoginOutput) NameClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectLogin) *string { return v.NameClaimType }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectLoginOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OpenIdConnectLogin) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type OpenIdConnectLoginPtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectLoginPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectLogin)(nil)).Elem()
}

func (o OpenIdConnectLoginPtrOutput) ToOpenIdConnectLoginPtrOutput() OpenIdConnectLoginPtrOutput {
	return o
}

func (o OpenIdConnectLoginPtrOutput) ToOpenIdConnectLoginPtrOutputWithContext(ctx context.Context) OpenIdConnectLoginPtrOutput {
	return o
}

func (o OpenIdConnectLoginPtrOutput) Elem() OpenIdConnectLoginOutput {
	return o.ApplyT(func(v *OpenIdConnectLogin) OpenIdConnectLogin {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectLogin
		return ret
	}).(OpenIdConnectLoginOutput)
}

func (o OpenIdConnectLoginPtrOutput) NameClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectLogin) *string {
		if v == nil {
			return nil
		}
		return v.NameClaimType
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectLoginPtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OpenIdConnectLogin) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type OpenIdConnectLoginResponse struct {
	NameClaimType *string  `pulumi:"nameClaimType"`
	Scopes        []string `pulumi:"scopes"`
}

type OpenIdConnectLoginResponseOutput struct{ *pulumi.OutputState }

func (OpenIdConnectLoginResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectLoginResponse)(nil)).Elem()
}

func (o OpenIdConnectLoginResponseOutput) ToOpenIdConnectLoginResponseOutput() OpenIdConnectLoginResponseOutput {
	return o
}

func (o OpenIdConnectLoginResponseOutput) ToOpenIdConnectLoginResponseOutputWithContext(ctx context.Context) OpenIdConnectLoginResponseOutput {
	return o
}

func (o OpenIdConnectLoginResponseOutput) NameClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectLoginResponse) *string { return v.NameClaimType }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectLoginResponseOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OpenIdConnectLoginResponse) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type OpenIdConnectLoginResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectLoginResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectLoginResponse)(nil)).Elem()
}

func (o OpenIdConnectLoginResponsePtrOutput) ToOpenIdConnectLoginResponsePtrOutput() OpenIdConnectLoginResponsePtrOutput {
	return o
}

func (o OpenIdConnectLoginResponsePtrOutput) ToOpenIdConnectLoginResponsePtrOutputWithContext(ctx context.Context) OpenIdConnectLoginResponsePtrOutput {
	return o
}

func (o OpenIdConnectLoginResponsePtrOutput) Elem() OpenIdConnectLoginResponseOutput {
	return o.ApplyT(func(v *OpenIdConnectLoginResponse) OpenIdConnectLoginResponse {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectLoginResponse
		return ret
	}).(OpenIdConnectLoginResponseOutput)
}

func (o OpenIdConnectLoginResponsePtrOutput) NameClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectLoginResponse) *string {
		if v == nil {
			return nil
		}
		return v.NameClaimType
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectLoginResponsePtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OpenIdConnectLoginResponse) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type OpenIdConnectRegistration struct {
	ClientCredential           *OpenIdConnectClientCredential `pulumi:"clientCredential"`
	ClientId                   *string                        `pulumi:"clientId"`
	OpenIdConnectConfiguration *OpenIdConnectConfig           `pulumi:"openIdConnectConfiguration"`
}





type OpenIdConnectRegistrationInput interface {
	pulumi.Input

	ToOpenIdConnectRegistrationOutput() OpenIdConnectRegistrationOutput
	ToOpenIdConnectRegistrationOutputWithContext(context.Context) OpenIdConnectRegistrationOutput
}

type OpenIdConnectRegistrationArgs struct {
	ClientCredential           OpenIdConnectClientCredentialPtrInput `pulumi:"clientCredential"`
	ClientId                   pulumi.StringPtrInput                 `pulumi:"clientId"`
	OpenIdConnectConfiguration OpenIdConnectConfigPtrInput           `pulumi:"openIdConnectConfiguration"`
}

func (OpenIdConnectRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectRegistration)(nil)).Elem()
}

func (i OpenIdConnectRegistrationArgs) ToOpenIdConnectRegistrationOutput() OpenIdConnectRegistrationOutput {
	return i.ToOpenIdConnectRegistrationOutputWithContext(context.Background())
}

func (i OpenIdConnectRegistrationArgs) ToOpenIdConnectRegistrationOutputWithContext(ctx context.Context) OpenIdConnectRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectRegistrationOutput)
}

func (i OpenIdConnectRegistrationArgs) ToOpenIdConnectRegistrationPtrOutput() OpenIdConnectRegistrationPtrOutput {
	return i.ToOpenIdConnectRegistrationPtrOutputWithContext(context.Background())
}

func (i OpenIdConnectRegistrationArgs) ToOpenIdConnectRegistrationPtrOutputWithContext(ctx context.Context) OpenIdConnectRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectRegistrationOutput).ToOpenIdConnectRegistrationPtrOutputWithContext(ctx)
}









type OpenIdConnectRegistrationPtrInput interface {
	pulumi.Input

	ToOpenIdConnectRegistrationPtrOutput() OpenIdConnectRegistrationPtrOutput
	ToOpenIdConnectRegistrationPtrOutputWithContext(context.Context) OpenIdConnectRegistrationPtrOutput
}

type openIdConnectRegistrationPtrType OpenIdConnectRegistrationArgs

func OpenIdConnectRegistrationPtr(v *OpenIdConnectRegistrationArgs) OpenIdConnectRegistrationPtrInput {
	return (*openIdConnectRegistrationPtrType)(v)
}

func (*openIdConnectRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectRegistration)(nil)).Elem()
}

func (i *openIdConnectRegistrationPtrType) ToOpenIdConnectRegistrationPtrOutput() OpenIdConnectRegistrationPtrOutput {
	return i.ToOpenIdConnectRegistrationPtrOutputWithContext(context.Background())
}

func (i *openIdConnectRegistrationPtrType) ToOpenIdConnectRegistrationPtrOutputWithContext(ctx context.Context) OpenIdConnectRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectRegistrationPtrOutput)
}

type OpenIdConnectRegistrationOutput struct{ *pulumi.OutputState }

func (OpenIdConnectRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectRegistration)(nil)).Elem()
}

func (o OpenIdConnectRegistrationOutput) ToOpenIdConnectRegistrationOutput() OpenIdConnectRegistrationOutput {
	return o
}

func (o OpenIdConnectRegistrationOutput) ToOpenIdConnectRegistrationOutputWithContext(ctx context.Context) OpenIdConnectRegistrationOutput {
	return o
}

func (o OpenIdConnectRegistrationOutput) ToOpenIdConnectRegistrationPtrOutput() OpenIdConnectRegistrationPtrOutput {
	return o.ToOpenIdConnectRegistrationPtrOutputWithContext(context.Background())
}

func (o OpenIdConnectRegistrationOutput) ToOpenIdConnectRegistrationPtrOutputWithContext(ctx context.Context) OpenIdConnectRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenIdConnectRegistration) *OpenIdConnectRegistration {
		return &v
	}).(OpenIdConnectRegistrationPtrOutput)
}

func (o OpenIdConnectRegistrationOutput) ClientCredential() OpenIdConnectClientCredentialPtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistration) *OpenIdConnectClientCredential { return v.ClientCredential }).(OpenIdConnectClientCredentialPtrOutput)
}

func (o OpenIdConnectRegistrationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistration) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectRegistrationOutput) OpenIdConnectConfiguration() OpenIdConnectConfigPtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistration) *OpenIdConnectConfig { return v.OpenIdConnectConfiguration }).(OpenIdConnectConfigPtrOutput)
}

type OpenIdConnectRegistrationPtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectRegistration)(nil)).Elem()
}

func (o OpenIdConnectRegistrationPtrOutput) ToOpenIdConnectRegistrationPtrOutput() OpenIdConnectRegistrationPtrOutput {
	return o
}

func (o OpenIdConnectRegistrationPtrOutput) ToOpenIdConnectRegistrationPtrOutputWithContext(ctx context.Context) OpenIdConnectRegistrationPtrOutput {
	return o
}

func (o OpenIdConnectRegistrationPtrOutput) Elem() OpenIdConnectRegistrationOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistration) OpenIdConnectRegistration {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectRegistration
		return ret
	}).(OpenIdConnectRegistrationOutput)
}

func (o OpenIdConnectRegistrationPtrOutput) ClientCredential() OpenIdConnectClientCredentialPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistration) *OpenIdConnectClientCredential {
		if v == nil {
			return nil
		}
		return v.ClientCredential
	}).(OpenIdConnectClientCredentialPtrOutput)
}

func (o OpenIdConnectRegistrationPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectRegistrationPtrOutput) OpenIdConnectConfiguration() OpenIdConnectConfigPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistration) *OpenIdConnectConfig {
		if v == nil {
			return nil
		}
		return v.OpenIdConnectConfiguration
	}).(OpenIdConnectConfigPtrOutput)
}

type OpenIdConnectRegistrationResponse struct {
	ClientCredential           *OpenIdConnectClientCredentialResponse `pulumi:"clientCredential"`
	ClientId                   *string                                `pulumi:"clientId"`
	OpenIdConnectConfiguration *OpenIdConnectConfigResponse           `pulumi:"openIdConnectConfiguration"`
}

type OpenIdConnectRegistrationResponseOutput struct{ *pulumi.OutputState }

func (OpenIdConnectRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectRegistrationResponse)(nil)).Elem()
}

func (o OpenIdConnectRegistrationResponseOutput) ToOpenIdConnectRegistrationResponseOutput() OpenIdConnectRegistrationResponseOutput {
	return o
}

func (o OpenIdConnectRegistrationResponseOutput) ToOpenIdConnectRegistrationResponseOutputWithContext(ctx context.Context) OpenIdConnectRegistrationResponseOutput {
	return o
}

func (o OpenIdConnectRegistrationResponseOutput) ClientCredential() OpenIdConnectClientCredentialResponsePtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistrationResponse) *OpenIdConnectClientCredentialResponse {
		return v.ClientCredential
	}).(OpenIdConnectClientCredentialResponsePtrOutput)
}

func (o OpenIdConnectRegistrationResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistrationResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectRegistrationResponseOutput) OpenIdConnectConfiguration() OpenIdConnectConfigResponsePtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistrationResponse) *OpenIdConnectConfigResponse {
		return v.OpenIdConnectConfiguration
	}).(OpenIdConnectConfigResponsePtrOutput)
}

type OpenIdConnectRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectRegistrationResponse)(nil)).Elem()
}

func (o OpenIdConnectRegistrationResponsePtrOutput) ToOpenIdConnectRegistrationResponsePtrOutput() OpenIdConnectRegistrationResponsePtrOutput {
	return o
}

func (o OpenIdConnectRegistrationResponsePtrOutput) ToOpenIdConnectRegistrationResponsePtrOutputWithContext(ctx context.Context) OpenIdConnectRegistrationResponsePtrOutput {
	return o
}

func (o OpenIdConnectRegistrationResponsePtrOutput) Elem() OpenIdConnectRegistrationResponseOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistrationResponse) OpenIdConnectRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectRegistrationResponse
		return ret
	}).(OpenIdConnectRegistrationResponseOutput)
}

func (o OpenIdConnectRegistrationResponsePtrOutput) ClientCredential() OpenIdConnectClientCredentialResponsePtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistrationResponse) *OpenIdConnectClientCredentialResponse {
		if v == nil {
			return nil
		}
		return v.ClientCredential
	}).(OpenIdConnectClientCredentialResponsePtrOutput)
}

func (o OpenIdConnectRegistrationResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectRegistrationResponsePtrOutput) OpenIdConnectConfiguration() OpenIdConnectConfigResponsePtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistrationResponse) *OpenIdConnectConfigResponse {
		if v == nil {
			return nil
		}
		return v.OpenIdConnectConfiguration
	}).(OpenIdConnectConfigResponsePtrOutput)
}

type QueueScaleRule struct {
	Auth        []ScaleRuleAuth `pulumi:"auth"`
	QueueLength *int            `pulumi:"queueLength"`
	QueueName   *string         `pulumi:"queueName"`
}





type QueueScaleRuleInput interface {
	pulumi.Input

	ToQueueScaleRuleOutput() QueueScaleRuleOutput
	ToQueueScaleRuleOutputWithContext(context.Context) QueueScaleRuleOutput
}

type QueueScaleRuleArgs struct {
	Auth        ScaleRuleAuthArrayInput `pulumi:"auth"`
	QueueLength pulumi.IntPtrInput      `pulumi:"queueLength"`
	QueueName   pulumi.StringPtrInput   `pulumi:"queueName"`
}

func (QueueScaleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueScaleRule)(nil)).Elem()
}

func (i QueueScaleRuleArgs) ToQueueScaleRuleOutput() QueueScaleRuleOutput {
	return i.ToQueueScaleRuleOutputWithContext(context.Background())
}

func (i QueueScaleRuleArgs) ToQueueScaleRuleOutputWithContext(ctx context.Context) QueueScaleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueScaleRuleOutput)
}

func (i QueueScaleRuleArgs) ToQueueScaleRulePtrOutput() QueueScaleRulePtrOutput {
	return i.ToQueueScaleRulePtrOutputWithContext(context.Background())
}

func (i QueueScaleRuleArgs) ToQueueScaleRulePtrOutputWithContext(ctx context.Context) QueueScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueScaleRuleOutput).ToQueueScaleRulePtrOutputWithContext(ctx)
}









type QueueScaleRulePtrInput interface {
	pulumi.Input

	ToQueueScaleRulePtrOutput() QueueScaleRulePtrOutput
	ToQueueScaleRulePtrOutputWithContext(context.Context) QueueScaleRulePtrOutput
}

type queueScaleRulePtrType QueueScaleRuleArgs

func QueueScaleRulePtr(v *QueueScaleRuleArgs) QueueScaleRulePtrInput {
	return (*queueScaleRulePtrType)(v)
}

func (*queueScaleRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueScaleRule)(nil)).Elem()
}

func (i *queueScaleRulePtrType) ToQueueScaleRulePtrOutput() QueueScaleRulePtrOutput {
	return i.ToQueueScaleRulePtrOutputWithContext(context.Background())
}

func (i *queueScaleRulePtrType) ToQueueScaleRulePtrOutputWithContext(ctx context.Context) QueueScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueScaleRulePtrOutput)
}

type QueueScaleRuleOutput struct{ *pulumi.OutputState }

func (QueueScaleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueScaleRule)(nil)).Elem()
}

func (o QueueScaleRuleOutput) ToQueueScaleRuleOutput() QueueScaleRuleOutput {
	return o
}

func (o QueueScaleRuleOutput) ToQueueScaleRuleOutputWithContext(ctx context.Context) QueueScaleRuleOutput {
	return o
}

func (o QueueScaleRuleOutput) ToQueueScaleRulePtrOutput() QueueScaleRulePtrOutput {
	return o.ToQueueScaleRulePtrOutputWithContext(context.Background())
}

func (o QueueScaleRuleOutput) ToQueueScaleRulePtrOutputWithContext(ctx context.Context) QueueScaleRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueueScaleRule) *QueueScaleRule {
		return &v
	}).(QueueScaleRulePtrOutput)
}

func (o QueueScaleRuleOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v QueueScaleRule) []ScaleRuleAuth { return v.Auth }).(ScaleRuleAuthArrayOutput)
}

func (o QueueScaleRuleOutput) QueueLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueScaleRule) *int { return v.QueueLength }).(pulumi.IntPtrOutput)
}

func (o QueueScaleRuleOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueScaleRule) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

type QueueScaleRulePtrOutput struct{ *pulumi.OutputState }

func (QueueScaleRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueScaleRule)(nil)).Elem()
}

func (o QueueScaleRulePtrOutput) ToQueueScaleRulePtrOutput() QueueScaleRulePtrOutput {
	return o
}

func (o QueueScaleRulePtrOutput) ToQueueScaleRulePtrOutputWithContext(ctx context.Context) QueueScaleRulePtrOutput {
	return o
}

func (o QueueScaleRulePtrOutput) Elem() QueueScaleRuleOutput {
	return o.ApplyT(func(v *QueueScaleRule) QueueScaleRule {
		if v != nil {
			return *v
		}
		var ret QueueScaleRule
		return ret
	}).(QueueScaleRuleOutput)
}

func (o QueueScaleRulePtrOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v *QueueScaleRule) []ScaleRuleAuth {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthArrayOutput)
}

func (o QueueScaleRulePtrOutput) QueueLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QueueScaleRule) *int {
		if v == nil {
			return nil
		}
		return v.QueueLength
	}).(pulumi.IntPtrOutput)
}

func (o QueueScaleRulePtrOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueScaleRule) *string {
		if v == nil {
			return nil
		}
		return v.QueueName
	}).(pulumi.StringPtrOutput)
}

type QueueScaleRuleResponse struct {
	Auth        []ScaleRuleAuthResponse `pulumi:"auth"`
	QueueLength *int                    `pulumi:"queueLength"`
	QueueName   *string                 `pulumi:"queueName"`
}

type QueueScaleRuleResponseOutput struct{ *pulumi.OutputState }

func (QueueScaleRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueScaleRuleResponse)(nil)).Elem()
}

func (o QueueScaleRuleResponseOutput) ToQueueScaleRuleResponseOutput() QueueScaleRuleResponseOutput {
	return o
}

func (o QueueScaleRuleResponseOutput) ToQueueScaleRuleResponseOutputWithContext(ctx context.Context) QueueScaleRuleResponseOutput {
	return o
}

func (o QueueScaleRuleResponseOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v QueueScaleRuleResponse) []ScaleRuleAuthResponse { return v.Auth }).(ScaleRuleAuthResponseArrayOutput)
}

func (o QueueScaleRuleResponseOutput) QueueLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueScaleRuleResponse) *int { return v.QueueLength }).(pulumi.IntPtrOutput)
}

func (o QueueScaleRuleResponseOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueScaleRuleResponse) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

type QueueScaleRuleResponsePtrOutput struct{ *pulumi.OutputState }

func (QueueScaleRuleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueScaleRuleResponse)(nil)).Elem()
}

func (o QueueScaleRuleResponsePtrOutput) ToQueueScaleRuleResponsePtrOutput() QueueScaleRuleResponsePtrOutput {
	return o
}

func (o QueueScaleRuleResponsePtrOutput) ToQueueScaleRuleResponsePtrOutputWithContext(ctx context.Context) QueueScaleRuleResponsePtrOutput {
	return o
}

func (o QueueScaleRuleResponsePtrOutput) Elem() QueueScaleRuleResponseOutput {
	return o.ApplyT(func(v *QueueScaleRuleResponse) QueueScaleRuleResponse {
		if v != nil {
			return *v
		}
		var ret QueueScaleRuleResponse
		return ret
	}).(QueueScaleRuleResponseOutput)
}

func (o QueueScaleRuleResponsePtrOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v *QueueScaleRuleResponse) []ScaleRuleAuthResponse {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthResponseArrayOutput)
}

func (o QueueScaleRuleResponsePtrOutput) QueueLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QueueScaleRuleResponse) *int {
		if v == nil {
			return nil
		}
		return v.QueueLength
	}).(pulumi.IntPtrOutput)
}

func (o QueueScaleRuleResponsePtrOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueScaleRuleResponse) *string {
		if v == nil {
			return nil
		}
		return v.QueueName
	}).(pulumi.StringPtrOutput)
}

type RegistryCredentials struct {
	Identity          *string `pulumi:"identity"`
	PasswordSecretRef *string `pulumi:"passwordSecretRef"`
	Server            *string `pulumi:"server"`
	Username          *string `pulumi:"username"`
}





type RegistryCredentialsInput interface {
	pulumi.Input

	ToRegistryCredentialsOutput() RegistryCredentialsOutput
	ToRegistryCredentialsOutputWithContext(context.Context) RegistryCredentialsOutput
}

type RegistryCredentialsArgs struct {
	Identity          pulumi.StringPtrInput `pulumi:"identity"`
	PasswordSecretRef pulumi.StringPtrInput `pulumi:"passwordSecretRef"`
	Server            pulumi.StringPtrInput `pulumi:"server"`
	Username          pulumi.StringPtrInput `pulumi:"username"`
}

func (RegistryCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryCredentials)(nil)).Elem()
}

func (i RegistryCredentialsArgs) ToRegistryCredentialsOutput() RegistryCredentialsOutput {
	return i.ToRegistryCredentialsOutputWithContext(context.Background())
}

func (i RegistryCredentialsArgs) ToRegistryCredentialsOutputWithContext(ctx context.Context) RegistryCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryCredentialsOutput)
}





type RegistryCredentialsArrayInput interface {
	pulumi.Input

	ToRegistryCredentialsArrayOutput() RegistryCredentialsArrayOutput
	ToRegistryCredentialsArrayOutputWithContext(context.Context) RegistryCredentialsArrayOutput
}

type RegistryCredentialsArray []RegistryCredentialsInput

func (RegistryCredentialsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryCredentials)(nil)).Elem()
}

func (i RegistryCredentialsArray) ToRegistryCredentialsArrayOutput() RegistryCredentialsArrayOutput {
	return i.ToRegistryCredentialsArrayOutputWithContext(context.Background())
}

func (i RegistryCredentialsArray) ToRegistryCredentialsArrayOutputWithContext(ctx context.Context) RegistryCredentialsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryCredentialsArrayOutput)
}

type RegistryCredentialsOutput struct{ *pulumi.OutputState }

func (RegistryCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryCredentials)(nil)).Elem()
}

func (o RegistryCredentialsOutput) ToRegistryCredentialsOutput() RegistryCredentialsOutput {
	return o
}

func (o RegistryCredentialsOutput) ToRegistryCredentialsOutputWithContext(ctx context.Context) RegistryCredentialsOutput {
	return o
}

func (o RegistryCredentialsOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentials) *string { return v.Identity }).(pulumi.StringPtrOutput)
}

func (o RegistryCredentialsOutput) PasswordSecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentials) *string { return v.PasswordSecretRef }).(pulumi.StringPtrOutput)
}

func (o RegistryCredentialsOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentials) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o RegistryCredentialsOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentials) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type RegistryCredentialsArrayOutput struct{ *pulumi.OutputState }

func (RegistryCredentialsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryCredentials)(nil)).Elem()
}

func (o RegistryCredentialsArrayOutput) ToRegistryCredentialsArrayOutput() RegistryCredentialsArrayOutput {
	return o
}

func (o RegistryCredentialsArrayOutput) ToRegistryCredentialsArrayOutputWithContext(ctx context.Context) RegistryCredentialsArrayOutput {
	return o
}

func (o RegistryCredentialsArrayOutput) Index(i pulumi.IntInput) RegistryCredentialsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryCredentials {
		return vs[0].([]RegistryCredentials)[vs[1].(int)]
	}).(RegistryCredentialsOutput)
}

type RegistryCredentialsResponse struct {
	Identity          *string `pulumi:"identity"`
	PasswordSecretRef *string `pulumi:"passwordSecretRef"`
	Server            *string `pulumi:"server"`
	Username          *string `pulumi:"username"`
}

type RegistryCredentialsResponseOutput struct{ *pulumi.OutputState }

func (RegistryCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryCredentialsResponse)(nil)).Elem()
}

func (o RegistryCredentialsResponseOutput) ToRegistryCredentialsResponseOutput() RegistryCredentialsResponseOutput {
	return o
}

func (o RegistryCredentialsResponseOutput) ToRegistryCredentialsResponseOutputWithContext(ctx context.Context) RegistryCredentialsResponseOutput {
	return o
}

func (o RegistryCredentialsResponseOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentialsResponse) *string { return v.Identity }).(pulumi.StringPtrOutput)
}

func (o RegistryCredentialsResponseOutput) PasswordSecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentialsResponse) *string { return v.PasswordSecretRef }).(pulumi.StringPtrOutput)
}

func (o RegistryCredentialsResponseOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentialsResponse) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o RegistryCredentialsResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentialsResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type RegistryCredentialsResponseArrayOutput struct{ *pulumi.OutputState }

func (RegistryCredentialsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryCredentialsResponse)(nil)).Elem()
}

func (o RegistryCredentialsResponseArrayOutput) ToRegistryCredentialsResponseArrayOutput() RegistryCredentialsResponseArrayOutput {
	return o
}

func (o RegistryCredentialsResponseArrayOutput) ToRegistryCredentialsResponseArrayOutputWithContext(ctx context.Context) RegistryCredentialsResponseArrayOutput {
	return o
}

func (o RegistryCredentialsResponseArrayOutput) Index(i pulumi.IntInput) RegistryCredentialsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryCredentialsResponse {
		return vs[0].([]RegistryCredentialsResponse)[vs[1].(int)]
	}).(RegistryCredentialsResponseOutput)
}

type RegistryInfo struct {
	RegistryPassword *string `pulumi:"registryPassword"`
	RegistryUrl      *string `pulumi:"registryUrl"`
	RegistryUserName *string `pulumi:"registryUserName"`
}





type RegistryInfoInput interface {
	pulumi.Input

	ToRegistryInfoOutput() RegistryInfoOutput
	ToRegistryInfoOutputWithContext(context.Context) RegistryInfoOutput
}

type RegistryInfoArgs struct {
	RegistryPassword pulumi.StringPtrInput `pulumi:"registryPassword"`
	RegistryUrl      pulumi.StringPtrInput `pulumi:"registryUrl"`
	RegistryUserName pulumi.StringPtrInput `pulumi:"registryUserName"`
}

func (RegistryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryInfo)(nil)).Elem()
}

func (i RegistryInfoArgs) ToRegistryInfoOutput() RegistryInfoOutput {
	return i.ToRegistryInfoOutputWithContext(context.Background())
}

func (i RegistryInfoArgs) ToRegistryInfoOutputWithContext(ctx context.Context) RegistryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryInfoOutput)
}

func (i RegistryInfoArgs) ToRegistryInfoPtrOutput() RegistryInfoPtrOutput {
	return i.ToRegistryInfoPtrOutputWithContext(context.Background())
}

func (i RegistryInfoArgs) ToRegistryInfoPtrOutputWithContext(ctx context.Context) RegistryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryInfoOutput).ToRegistryInfoPtrOutputWithContext(ctx)
}









type RegistryInfoPtrInput interface {
	pulumi.Input

	ToRegistryInfoPtrOutput() RegistryInfoPtrOutput
	ToRegistryInfoPtrOutputWithContext(context.Context) RegistryInfoPtrOutput
}

type registryInfoPtrType RegistryInfoArgs

func RegistryInfoPtr(v *RegistryInfoArgs) RegistryInfoPtrInput {
	return (*registryInfoPtrType)(v)
}

func (*registryInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryInfo)(nil)).Elem()
}

func (i *registryInfoPtrType) ToRegistryInfoPtrOutput() RegistryInfoPtrOutput {
	return i.ToRegistryInfoPtrOutputWithContext(context.Background())
}

func (i *registryInfoPtrType) ToRegistryInfoPtrOutputWithContext(ctx context.Context) RegistryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryInfoPtrOutput)
}

type RegistryInfoOutput struct{ *pulumi.OutputState }

func (RegistryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryInfo)(nil)).Elem()
}

func (o RegistryInfoOutput) ToRegistryInfoOutput() RegistryInfoOutput {
	return o
}

func (o RegistryInfoOutput) ToRegistryInfoOutputWithContext(ctx context.Context) RegistryInfoOutput {
	return o
}

func (o RegistryInfoOutput) ToRegistryInfoPtrOutput() RegistryInfoPtrOutput {
	return o.ToRegistryInfoPtrOutputWithContext(context.Background())
}

func (o RegistryInfoOutput) ToRegistryInfoPtrOutputWithContext(ctx context.Context) RegistryInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistryInfo) *RegistryInfo {
		return &v
	}).(RegistryInfoPtrOutput)
}

func (o RegistryInfoOutput) RegistryPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryInfo) *string { return v.RegistryPassword }).(pulumi.StringPtrOutput)
}

func (o RegistryInfoOutput) RegistryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryInfo) *string { return v.RegistryUrl }).(pulumi.StringPtrOutput)
}

func (o RegistryInfoOutput) RegistryUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryInfo) *string { return v.RegistryUserName }).(pulumi.StringPtrOutput)
}

type RegistryInfoPtrOutput struct{ *pulumi.OutputState }

func (RegistryInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryInfo)(nil)).Elem()
}

func (o RegistryInfoPtrOutput) ToRegistryInfoPtrOutput() RegistryInfoPtrOutput {
	return o
}

func (o RegistryInfoPtrOutput) ToRegistryInfoPtrOutputWithContext(ctx context.Context) RegistryInfoPtrOutput {
	return o
}

func (o RegistryInfoPtrOutput) Elem() RegistryInfoOutput {
	return o.ApplyT(func(v *RegistryInfo) RegistryInfo {
		if v != nil {
			return *v
		}
		var ret RegistryInfo
		return ret
	}).(RegistryInfoOutput)
}

func (o RegistryInfoPtrOutput) RegistryPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryInfo) *string {
		if v == nil {
			return nil
		}
		return v.RegistryPassword
	}).(pulumi.StringPtrOutput)
}

func (o RegistryInfoPtrOutput) RegistryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryInfo) *string {
		if v == nil {
			return nil
		}
		return v.RegistryUrl
	}).(pulumi.StringPtrOutput)
}

func (o RegistryInfoPtrOutput) RegistryUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryInfo) *string {
		if v == nil {
			return nil
		}
		return v.RegistryUserName
	}).(pulumi.StringPtrOutput)
}

type RegistryInfoResponse struct {
	RegistryUrl      *string `pulumi:"registryUrl"`
	RegistryUserName *string `pulumi:"registryUserName"`
}

type RegistryInfoResponseOutput struct{ *pulumi.OutputState }

func (RegistryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryInfoResponse)(nil)).Elem()
}

func (o RegistryInfoResponseOutput) ToRegistryInfoResponseOutput() RegistryInfoResponseOutput {
	return o
}

func (o RegistryInfoResponseOutput) ToRegistryInfoResponseOutputWithContext(ctx context.Context) RegistryInfoResponseOutput {
	return o
}

func (o RegistryInfoResponseOutput) RegistryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryInfoResponse) *string { return v.RegistryUrl }).(pulumi.StringPtrOutput)
}

func (o RegistryInfoResponseOutput) RegistryUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryInfoResponse) *string { return v.RegistryUserName }).(pulumi.StringPtrOutput)
}

type RegistryInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (RegistryInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryInfoResponse)(nil)).Elem()
}

func (o RegistryInfoResponsePtrOutput) ToRegistryInfoResponsePtrOutput() RegistryInfoResponsePtrOutput {
	return o
}

func (o RegistryInfoResponsePtrOutput) ToRegistryInfoResponsePtrOutputWithContext(ctx context.Context) RegistryInfoResponsePtrOutput {
	return o
}

func (o RegistryInfoResponsePtrOutput) Elem() RegistryInfoResponseOutput {
	return o.ApplyT(func(v *RegistryInfoResponse) RegistryInfoResponse {
		if v != nil {
			return *v
		}
		var ret RegistryInfoResponse
		return ret
	}).(RegistryInfoResponseOutput)
}

func (o RegistryInfoResponsePtrOutput) RegistryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistryUrl
	}).(pulumi.StringPtrOutput)
}

func (o RegistryInfoResponsePtrOutput) RegistryUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistryUserName
	}).(pulumi.StringPtrOutput)
}

type Scale struct {
	MaxReplicas *int        `pulumi:"maxReplicas"`
	MinReplicas *int        `pulumi:"minReplicas"`
	Rules       []ScaleRule `pulumi:"rules"`
}





type ScaleInput interface {
	pulumi.Input

	ToScaleOutput() ScaleOutput
	ToScaleOutputWithContext(context.Context) ScaleOutput
}

type ScaleArgs struct {
	MaxReplicas pulumi.IntPtrInput  `pulumi:"maxReplicas"`
	MinReplicas pulumi.IntPtrInput  `pulumi:"minReplicas"`
	Rules       ScaleRuleArrayInput `pulumi:"rules"`
}

func (ScaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Scale)(nil)).Elem()
}

func (i ScaleArgs) ToScaleOutput() ScaleOutput {
	return i.ToScaleOutputWithContext(context.Background())
}

func (i ScaleArgs) ToScaleOutputWithContext(ctx context.Context) ScaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleOutput)
}

func (i ScaleArgs) ToScalePtrOutput() ScalePtrOutput {
	return i.ToScalePtrOutputWithContext(context.Background())
}

func (i ScaleArgs) ToScalePtrOutputWithContext(ctx context.Context) ScalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleOutput).ToScalePtrOutputWithContext(ctx)
}









type ScalePtrInput interface {
	pulumi.Input

	ToScalePtrOutput() ScalePtrOutput
	ToScalePtrOutputWithContext(context.Context) ScalePtrOutput
}

type scalePtrType ScaleArgs

func ScalePtr(v *ScaleArgs) ScalePtrInput {
	return (*scalePtrType)(v)
}

func (*scalePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Scale)(nil)).Elem()
}

func (i *scalePtrType) ToScalePtrOutput() ScalePtrOutput {
	return i.ToScalePtrOutputWithContext(context.Background())
}

func (i *scalePtrType) ToScalePtrOutputWithContext(ctx context.Context) ScalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalePtrOutput)
}

type ScaleOutput struct{ *pulumi.OutputState }

func (ScaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Scale)(nil)).Elem()
}

func (o ScaleOutput) ToScaleOutput() ScaleOutput {
	return o
}

func (o ScaleOutput) ToScaleOutputWithContext(ctx context.Context) ScaleOutput {
	return o
}

func (o ScaleOutput) ToScalePtrOutput() ScalePtrOutput {
	return o.ToScalePtrOutputWithContext(context.Background())
}

func (o ScaleOutput) ToScalePtrOutputWithContext(ctx context.Context) ScalePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Scale) *Scale {
		return &v
	}).(ScalePtrOutput)
}

func (o ScaleOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Scale) *int { return v.MaxReplicas }).(pulumi.IntPtrOutput)
}

func (o ScaleOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Scale) *int { return v.MinReplicas }).(pulumi.IntPtrOutput)
}

func (o ScaleOutput) Rules() ScaleRuleArrayOutput {
	return o.ApplyT(func(v Scale) []ScaleRule { return v.Rules }).(ScaleRuleArrayOutput)
}

type ScalePtrOutput struct{ *pulumi.OutputState }

func (ScalePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scale)(nil)).Elem()
}

func (o ScalePtrOutput) ToScalePtrOutput() ScalePtrOutput {
	return o
}

func (o ScalePtrOutput) ToScalePtrOutputWithContext(ctx context.Context) ScalePtrOutput {
	return o
}

func (o ScalePtrOutput) Elem() ScaleOutput {
	return o.ApplyT(func(v *Scale) Scale {
		if v != nil {
			return *v
		}
		var ret Scale
		return ret
	}).(ScaleOutput)
}

func (o ScalePtrOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Scale) *int {
		if v == nil {
			return nil
		}
		return v.MaxReplicas
	}).(pulumi.IntPtrOutput)
}

func (o ScalePtrOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Scale) *int {
		if v == nil {
			return nil
		}
		return v.MinReplicas
	}).(pulumi.IntPtrOutput)
}

func (o ScalePtrOutput) Rules() ScaleRuleArrayOutput {
	return o.ApplyT(func(v *Scale) []ScaleRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ScaleRuleArrayOutput)
}

type ScaleResponse struct {
	MaxReplicas *int                `pulumi:"maxReplicas"`
	MinReplicas *int                `pulumi:"minReplicas"`
	Rules       []ScaleRuleResponse `pulumi:"rules"`
}

type ScaleResponseOutput struct{ *pulumi.OutputState }

func (ScaleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleResponse)(nil)).Elem()
}

func (o ScaleResponseOutput) ToScaleResponseOutput() ScaleResponseOutput {
	return o
}

func (o ScaleResponseOutput) ToScaleResponseOutputWithContext(ctx context.Context) ScaleResponseOutput {
	return o
}

func (o ScaleResponseOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScaleResponse) *int { return v.MaxReplicas }).(pulumi.IntPtrOutput)
}

func (o ScaleResponseOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScaleResponse) *int { return v.MinReplicas }).(pulumi.IntPtrOutput)
}

func (o ScaleResponseOutput) Rules() ScaleRuleResponseArrayOutput {
	return o.ApplyT(func(v ScaleResponse) []ScaleRuleResponse { return v.Rules }).(ScaleRuleResponseArrayOutput)
}

type ScaleResponsePtrOutput struct{ *pulumi.OutputState }

func (ScaleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleResponse)(nil)).Elem()
}

func (o ScaleResponsePtrOutput) ToScaleResponsePtrOutput() ScaleResponsePtrOutput {
	return o
}

func (o ScaleResponsePtrOutput) ToScaleResponsePtrOutputWithContext(ctx context.Context) ScaleResponsePtrOutput {
	return o
}

func (o ScaleResponsePtrOutput) Elem() ScaleResponseOutput {
	return o.ApplyT(func(v *ScaleResponse) ScaleResponse {
		if v != nil {
			return *v
		}
		var ret ScaleResponse
		return ret
	}).(ScaleResponseOutput)
}

func (o ScaleResponsePtrOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxReplicas
	}).(pulumi.IntPtrOutput)
}

func (o ScaleResponsePtrOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinReplicas
	}).(pulumi.IntPtrOutput)
}

func (o ScaleResponsePtrOutput) Rules() ScaleRuleResponseArrayOutput {
	return o.ApplyT(func(v *ScaleResponse) []ScaleRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ScaleRuleResponseArrayOutput)
}

type ScaleRule struct {
	AzureQueue *QueueScaleRule  `pulumi:"azureQueue"`
	Custom     *CustomScaleRule `pulumi:"custom"`
	Http       *HttpScaleRule   `pulumi:"http"`
	Name       *string          `pulumi:"name"`
}





type ScaleRuleInput interface {
	pulumi.Input

	ToScaleRuleOutput() ScaleRuleOutput
	ToScaleRuleOutputWithContext(context.Context) ScaleRuleOutput
}

type ScaleRuleArgs struct {
	AzureQueue QueueScaleRulePtrInput  `pulumi:"azureQueue"`
	Custom     CustomScaleRulePtrInput `pulumi:"custom"`
	Http       HttpScaleRulePtrInput   `pulumi:"http"`
	Name       pulumi.StringPtrInput   `pulumi:"name"`
}

func (ScaleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRule)(nil)).Elem()
}

func (i ScaleRuleArgs) ToScaleRuleOutput() ScaleRuleOutput {
	return i.ToScaleRuleOutputWithContext(context.Background())
}

func (i ScaleRuleArgs) ToScaleRuleOutputWithContext(ctx context.Context) ScaleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleOutput)
}





type ScaleRuleArrayInput interface {
	pulumi.Input

	ToScaleRuleArrayOutput() ScaleRuleArrayOutput
	ToScaleRuleArrayOutputWithContext(context.Context) ScaleRuleArrayOutput
}

type ScaleRuleArray []ScaleRuleInput

func (ScaleRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRule)(nil)).Elem()
}

func (i ScaleRuleArray) ToScaleRuleArrayOutput() ScaleRuleArrayOutput {
	return i.ToScaleRuleArrayOutputWithContext(context.Background())
}

func (i ScaleRuleArray) ToScaleRuleArrayOutputWithContext(ctx context.Context) ScaleRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleArrayOutput)
}

type ScaleRuleOutput struct{ *pulumi.OutputState }

func (ScaleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRule)(nil)).Elem()
}

func (o ScaleRuleOutput) ToScaleRuleOutput() ScaleRuleOutput {
	return o
}

func (o ScaleRuleOutput) ToScaleRuleOutputWithContext(ctx context.Context) ScaleRuleOutput {
	return o
}

func (o ScaleRuleOutput) AzureQueue() QueueScaleRulePtrOutput {
	return o.ApplyT(func(v ScaleRule) *QueueScaleRule { return v.AzureQueue }).(QueueScaleRulePtrOutput)
}

func (o ScaleRuleOutput) Custom() CustomScaleRulePtrOutput {
	return o.ApplyT(func(v ScaleRule) *CustomScaleRule { return v.Custom }).(CustomScaleRulePtrOutput)
}

func (o ScaleRuleOutput) Http() HttpScaleRulePtrOutput {
	return o.ApplyT(func(v ScaleRule) *HttpScaleRule { return v.Http }).(HttpScaleRulePtrOutput)
}

func (o ScaleRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ScaleRuleArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRule)(nil)).Elem()
}

func (o ScaleRuleArrayOutput) ToScaleRuleArrayOutput() ScaleRuleArrayOutput {
	return o
}

func (o ScaleRuleArrayOutput) ToScaleRuleArrayOutputWithContext(ctx context.Context) ScaleRuleArrayOutput {
	return o
}

func (o ScaleRuleArrayOutput) Index(i pulumi.IntInput) ScaleRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRule {
		return vs[0].([]ScaleRule)[vs[1].(int)]
	}).(ScaleRuleOutput)
}

type ScaleRuleAuth struct {
	SecretRef        *string `pulumi:"secretRef"`
	TriggerParameter *string `pulumi:"triggerParameter"`
}





type ScaleRuleAuthInput interface {
	pulumi.Input

	ToScaleRuleAuthOutput() ScaleRuleAuthOutput
	ToScaleRuleAuthOutputWithContext(context.Context) ScaleRuleAuthOutput
}

type ScaleRuleAuthArgs struct {
	SecretRef        pulumi.StringPtrInput `pulumi:"secretRef"`
	TriggerParameter pulumi.StringPtrInput `pulumi:"triggerParameter"`
}

func (ScaleRuleAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleAuth)(nil)).Elem()
}

func (i ScaleRuleAuthArgs) ToScaleRuleAuthOutput() ScaleRuleAuthOutput {
	return i.ToScaleRuleAuthOutputWithContext(context.Background())
}

func (i ScaleRuleAuthArgs) ToScaleRuleAuthOutputWithContext(ctx context.Context) ScaleRuleAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleAuthOutput)
}





type ScaleRuleAuthArrayInput interface {
	pulumi.Input

	ToScaleRuleAuthArrayOutput() ScaleRuleAuthArrayOutput
	ToScaleRuleAuthArrayOutputWithContext(context.Context) ScaleRuleAuthArrayOutput
}

type ScaleRuleAuthArray []ScaleRuleAuthInput

func (ScaleRuleAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleAuth)(nil)).Elem()
}

func (i ScaleRuleAuthArray) ToScaleRuleAuthArrayOutput() ScaleRuleAuthArrayOutput {
	return i.ToScaleRuleAuthArrayOutputWithContext(context.Background())
}

func (i ScaleRuleAuthArray) ToScaleRuleAuthArrayOutputWithContext(ctx context.Context) ScaleRuleAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleAuthArrayOutput)
}

type ScaleRuleAuthOutput struct{ *pulumi.OutputState }

func (ScaleRuleAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleAuth)(nil)).Elem()
}

func (o ScaleRuleAuthOutput) ToScaleRuleAuthOutput() ScaleRuleAuthOutput {
	return o
}

func (o ScaleRuleAuthOutput) ToScaleRuleAuthOutputWithContext(ctx context.Context) ScaleRuleAuthOutput {
	return o
}

func (o ScaleRuleAuthOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRuleAuth) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o ScaleRuleAuthOutput) TriggerParameter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRuleAuth) *string { return v.TriggerParameter }).(pulumi.StringPtrOutput)
}

type ScaleRuleAuthArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleAuth)(nil)).Elem()
}

func (o ScaleRuleAuthArrayOutput) ToScaleRuleAuthArrayOutput() ScaleRuleAuthArrayOutput {
	return o
}

func (o ScaleRuleAuthArrayOutput) ToScaleRuleAuthArrayOutputWithContext(ctx context.Context) ScaleRuleAuthArrayOutput {
	return o
}

func (o ScaleRuleAuthArrayOutput) Index(i pulumi.IntInput) ScaleRuleAuthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRuleAuth {
		return vs[0].([]ScaleRuleAuth)[vs[1].(int)]
	}).(ScaleRuleAuthOutput)
}

type ScaleRuleAuthResponse struct {
	SecretRef        *string `pulumi:"secretRef"`
	TriggerParameter *string `pulumi:"triggerParameter"`
}

type ScaleRuleAuthResponseOutput struct{ *pulumi.OutputState }

func (ScaleRuleAuthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleAuthResponse)(nil)).Elem()
}

func (o ScaleRuleAuthResponseOutput) ToScaleRuleAuthResponseOutput() ScaleRuleAuthResponseOutput {
	return o
}

func (o ScaleRuleAuthResponseOutput) ToScaleRuleAuthResponseOutputWithContext(ctx context.Context) ScaleRuleAuthResponseOutput {
	return o
}

func (o ScaleRuleAuthResponseOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRuleAuthResponse) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o ScaleRuleAuthResponseOutput) TriggerParameter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRuleAuthResponse) *string { return v.TriggerParameter }).(pulumi.StringPtrOutput)
}

type ScaleRuleAuthResponseArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleAuthResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleAuthResponse)(nil)).Elem()
}

func (o ScaleRuleAuthResponseArrayOutput) ToScaleRuleAuthResponseArrayOutput() ScaleRuleAuthResponseArrayOutput {
	return o
}

func (o ScaleRuleAuthResponseArrayOutput) ToScaleRuleAuthResponseArrayOutputWithContext(ctx context.Context) ScaleRuleAuthResponseArrayOutput {
	return o
}

func (o ScaleRuleAuthResponseArrayOutput) Index(i pulumi.IntInput) ScaleRuleAuthResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRuleAuthResponse {
		return vs[0].([]ScaleRuleAuthResponse)[vs[1].(int)]
	}).(ScaleRuleAuthResponseOutput)
}

type ScaleRuleResponse struct {
	AzureQueue *QueueScaleRuleResponse  `pulumi:"azureQueue"`
	Custom     *CustomScaleRuleResponse `pulumi:"custom"`
	Http       *HttpScaleRuleResponse   `pulumi:"http"`
	Name       *string                  `pulumi:"name"`
}

type ScaleRuleResponseOutput struct{ *pulumi.OutputState }

func (ScaleRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleResponse)(nil)).Elem()
}

func (o ScaleRuleResponseOutput) ToScaleRuleResponseOutput() ScaleRuleResponseOutput {
	return o
}

func (o ScaleRuleResponseOutput) ToScaleRuleResponseOutputWithContext(ctx context.Context) ScaleRuleResponseOutput {
	return o
}

func (o ScaleRuleResponseOutput) AzureQueue() QueueScaleRuleResponsePtrOutput {
	return o.ApplyT(func(v ScaleRuleResponse) *QueueScaleRuleResponse { return v.AzureQueue }).(QueueScaleRuleResponsePtrOutput)
}

func (o ScaleRuleResponseOutput) Custom() CustomScaleRuleResponsePtrOutput {
	return o.ApplyT(func(v ScaleRuleResponse) *CustomScaleRuleResponse { return v.Custom }).(CustomScaleRuleResponsePtrOutput)
}

func (o ScaleRuleResponseOutput) Http() HttpScaleRuleResponsePtrOutput {
	return o.ApplyT(func(v ScaleRuleResponse) *HttpScaleRuleResponse { return v.Http }).(HttpScaleRuleResponsePtrOutput)
}

func (o ScaleRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ScaleRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleResponse)(nil)).Elem()
}

func (o ScaleRuleResponseArrayOutput) ToScaleRuleResponseArrayOutput() ScaleRuleResponseArrayOutput {
	return o
}

func (o ScaleRuleResponseArrayOutput) ToScaleRuleResponseArrayOutputWithContext(ctx context.Context) ScaleRuleResponseArrayOutput {
	return o
}

func (o ScaleRuleResponseArrayOutput) Index(i pulumi.IntInput) ScaleRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRuleResponse {
		return vs[0].([]ScaleRuleResponse)[vs[1].(int)]
	}).(ScaleRuleResponseOutput)
}

type Secret struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type SecretInput interface {
	pulumi.Input

	ToSecretOutput() SecretOutput
	ToSecretOutputWithContext(context.Context) SecretOutput
}

type SecretArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Secret)(nil)).Elem()
}

func (i SecretArgs) ToSecretOutput() SecretOutput {
	return i.ToSecretOutputWithContext(context.Background())
}

func (i SecretArgs) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretOutput)
}





type SecretArrayInput interface {
	pulumi.Input

	ToSecretArrayOutput() SecretArrayOutput
	ToSecretArrayOutputWithContext(context.Context) SecretArrayOutput
}

type SecretArray []SecretInput

func (SecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Secret)(nil)).Elem()
}

func (i SecretArray) ToSecretArrayOutput() SecretArrayOutput {
	return i.ToSecretArrayOutputWithContext(context.Background())
}

func (i SecretArray) ToSecretArrayOutputWithContext(ctx context.Context) SecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretArrayOutput)
}

type SecretOutput struct{ *pulumi.OutputState }

func (SecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Secret)(nil)).Elem()
}

func (o SecretOutput) ToSecretOutput() SecretOutput {
	return o
}

func (o SecretOutput) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return o
}

func (o SecretOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Secret) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecretOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Secret) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SecretArrayOutput struct{ *pulumi.OutputState }

func (SecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Secret)(nil)).Elem()
}

func (o SecretArrayOutput) ToSecretArrayOutput() SecretArrayOutput {
	return o
}

func (o SecretArrayOutput) ToSecretArrayOutputWithContext(ctx context.Context) SecretArrayOutput {
	return o
}

func (o SecretArrayOutput) Index(i pulumi.IntInput) SecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Secret {
		return vs[0].([]Secret)[vs[1].(int)]
	}).(SecretOutput)
}

type SecretResponse struct {
	Name *string `pulumi:"name"`
}

type SecretResponseOutput struct{ *pulumi.OutputState }

func (SecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretResponse)(nil)).Elem()
}

func (o SecretResponseOutput) ToSecretResponseOutput() SecretResponseOutput {
	return o
}

func (o SecretResponseOutput) ToSecretResponseOutputWithContext(ctx context.Context) SecretResponseOutput {
	return o
}

func (o SecretResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SecretResponseArrayOutput struct{ *pulumi.OutputState }

func (SecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretResponse)(nil)).Elem()
}

func (o SecretResponseArrayOutput) ToSecretResponseArrayOutput() SecretResponseArrayOutput {
	return o
}

func (o SecretResponseArrayOutput) ToSecretResponseArrayOutputWithContext(ctx context.Context) SecretResponseArrayOutput {
	return o
}

func (o SecretResponseArrayOutput) Index(i pulumi.IntInput) SecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretResponse {
		return vs[0].([]SecretResponse)[vs[1].(int)]
	}).(SecretResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

type Template struct {
	Containers     []Container `pulumi:"containers"`
	RevisionSuffix *string     `pulumi:"revisionSuffix"`
	Scale          *Scale      `pulumi:"scale"`
	Volumes        []Volume    `pulumi:"volumes"`
}





type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(context.Context) TemplateOutput
}

type TemplateArgs struct {
	Containers     ContainerArrayInput   `pulumi:"containers"`
	RevisionSuffix pulumi.StringPtrInput `pulumi:"revisionSuffix"`
	Scale          ScalePtrInput         `pulumi:"scale"`
	Volumes        VolumeArrayInput      `pulumi:"volumes"`
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Template)(nil)).Elem()
}

func (i TemplateArgs) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i TemplateArgs) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

func (i TemplateArgs) ToTemplatePtrOutput() TemplatePtrOutput {
	return i.ToTemplatePtrOutputWithContext(context.Background())
}

func (i TemplateArgs) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput).ToTemplatePtrOutputWithContext(ctx)
}









type TemplatePtrInput interface {
	pulumi.Input

	ToTemplatePtrOutput() TemplatePtrOutput
	ToTemplatePtrOutputWithContext(context.Context) TemplatePtrOutput
}

type templatePtrType TemplateArgs

func TemplatePtr(v *TemplateArgs) TemplatePtrInput {
	return (*templatePtrType)(v)
}

func (*templatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (i *templatePtrType) ToTemplatePtrOutput() TemplatePtrOutput {
	return i.ToTemplatePtrOutputWithContext(context.Background())
}

func (i *templatePtrType) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplatePtrOutput)
}

type TemplateOutput struct{ *pulumi.OutputState }

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Template)(nil)).Elem()
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplatePtrOutput() TemplatePtrOutput {
	return o.ToTemplatePtrOutputWithContext(context.Background())
}

func (o TemplateOutput) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Template) *Template {
		return &v
	}).(TemplatePtrOutput)
}

func (o TemplateOutput) Containers() ContainerArrayOutput {
	return o.ApplyT(func(v Template) []Container { return v.Containers }).(ContainerArrayOutput)
}

func (o TemplateOutput) RevisionSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Template) *string { return v.RevisionSuffix }).(pulumi.StringPtrOutput)
}

func (o TemplateOutput) Scale() ScalePtrOutput {
	return o.ApplyT(func(v Template) *Scale { return v.Scale }).(ScalePtrOutput)
}

func (o TemplateOutput) Volumes() VolumeArrayOutput {
	return o.ApplyT(func(v Template) []Volume { return v.Volumes }).(VolumeArrayOutput)
}

type TemplatePtrOutput struct{ *pulumi.OutputState }

func (TemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (o TemplatePtrOutput) ToTemplatePtrOutput() TemplatePtrOutput {
	return o
}

func (o TemplatePtrOutput) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return o
}

func (o TemplatePtrOutput) Elem() TemplateOutput {
	return o.ApplyT(func(v *Template) Template {
		if v != nil {
			return *v
		}
		var ret Template
		return ret
	}).(TemplateOutput)
}

func (o TemplatePtrOutput) Containers() ContainerArrayOutput {
	return o.ApplyT(func(v *Template) []Container {
		if v == nil {
			return nil
		}
		return v.Containers
	}).(ContainerArrayOutput)
}

func (o TemplatePtrOutput) RevisionSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) *string {
		if v == nil {
			return nil
		}
		return v.RevisionSuffix
	}).(pulumi.StringPtrOutput)
}

func (o TemplatePtrOutput) Scale() ScalePtrOutput {
	return o.ApplyT(func(v *Template) *Scale {
		if v == nil {
			return nil
		}
		return v.Scale
	}).(ScalePtrOutput)
}

func (o TemplatePtrOutput) Volumes() VolumeArrayOutput {
	return o.ApplyT(func(v *Template) []Volume {
		if v == nil {
			return nil
		}
		return v.Volumes
	}).(VolumeArrayOutput)
}

type TemplateResponse struct {
	Containers     []ContainerResponse `pulumi:"containers"`
	RevisionSuffix *string             `pulumi:"revisionSuffix"`
	Scale          *ScaleResponse      `pulumi:"scale"`
	Volumes        []VolumeResponse    `pulumi:"volumes"`
}

type TemplateResponseOutput struct{ *pulumi.OutputState }

func (TemplateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateResponse)(nil)).Elem()
}

func (o TemplateResponseOutput) ToTemplateResponseOutput() TemplateResponseOutput {
	return o
}

func (o TemplateResponseOutput) ToTemplateResponseOutputWithContext(ctx context.Context) TemplateResponseOutput {
	return o
}

func (o TemplateResponseOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v TemplateResponse) []ContainerResponse { return v.Containers }).(ContainerResponseArrayOutput)
}

func (o TemplateResponseOutput) RevisionSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateResponse) *string { return v.RevisionSuffix }).(pulumi.StringPtrOutput)
}

func (o TemplateResponseOutput) Scale() ScaleResponsePtrOutput {
	return o.ApplyT(func(v TemplateResponse) *ScaleResponse { return v.Scale }).(ScaleResponsePtrOutput)
}

func (o TemplateResponseOutput) Volumes() VolumeResponseArrayOutput {
	return o.ApplyT(func(v TemplateResponse) []VolumeResponse { return v.Volumes }).(VolumeResponseArrayOutput)
}

type TemplateResponsePtrOutput struct{ *pulumi.OutputState }

func (TemplateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateResponse)(nil)).Elem()
}

func (o TemplateResponsePtrOutput) ToTemplateResponsePtrOutput() TemplateResponsePtrOutput {
	return o
}

func (o TemplateResponsePtrOutput) ToTemplateResponsePtrOutputWithContext(ctx context.Context) TemplateResponsePtrOutput {
	return o
}

func (o TemplateResponsePtrOutput) Elem() TemplateResponseOutput {
	return o.ApplyT(func(v *TemplateResponse) TemplateResponse {
		if v != nil {
			return *v
		}
		var ret TemplateResponse
		return ret
	}).(TemplateResponseOutput)
}

func (o TemplateResponsePtrOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v *TemplateResponse) []ContainerResponse {
		if v == nil {
			return nil
		}
		return v.Containers
	}).(ContainerResponseArrayOutput)
}

func (o TemplateResponsePtrOutput) RevisionSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateResponse) *string {
		if v == nil {
			return nil
		}
		return v.RevisionSuffix
	}).(pulumi.StringPtrOutput)
}

func (o TemplateResponsePtrOutput) Scale() ScaleResponsePtrOutput {
	return o.ApplyT(func(v *TemplateResponse) *ScaleResponse {
		if v == nil {
			return nil
		}
		return v.Scale
	}).(ScaleResponsePtrOutput)
}

func (o TemplateResponsePtrOutput) Volumes() VolumeResponseArrayOutput {
	return o.ApplyT(func(v *TemplateResponse) []VolumeResponse {
		if v == nil {
			return nil
		}
		return v.Volumes
	}).(VolumeResponseArrayOutput)
}

type TrafficWeight struct {
	Label          *string `pulumi:"label"`
	LatestRevision *bool   `pulumi:"latestRevision"`
	RevisionName   *string `pulumi:"revisionName"`
	Weight         *int    `pulumi:"weight"`
}


func (val *TrafficWeight) Defaults() *TrafficWeight {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LatestRevision) {
		latestRevision_ := false
		tmp.LatestRevision = &latestRevision_
	}
	return &tmp
}





type TrafficWeightInput interface {
	pulumi.Input

	ToTrafficWeightOutput() TrafficWeightOutput
	ToTrafficWeightOutputWithContext(context.Context) TrafficWeightOutput
}

type TrafficWeightArgs struct {
	Label          pulumi.StringPtrInput `pulumi:"label"`
	LatestRevision pulumi.BoolPtrInput   `pulumi:"latestRevision"`
	RevisionName   pulumi.StringPtrInput `pulumi:"revisionName"`
	Weight         pulumi.IntPtrInput    `pulumi:"weight"`
}


func (val *TrafficWeightArgs) Defaults() *TrafficWeightArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LatestRevision) {
		tmp.LatestRevision = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (TrafficWeightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficWeight)(nil)).Elem()
}

func (i TrafficWeightArgs) ToTrafficWeightOutput() TrafficWeightOutput {
	return i.ToTrafficWeightOutputWithContext(context.Background())
}

func (i TrafficWeightArgs) ToTrafficWeightOutputWithContext(ctx context.Context) TrafficWeightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficWeightOutput)
}





type TrafficWeightArrayInput interface {
	pulumi.Input

	ToTrafficWeightArrayOutput() TrafficWeightArrayOutput
	ToTrafficWeightArrayOutputWithContext(context.Context) TrafficWeightArrayOutput
}

type TrafficWeightArray []TrafficWeightInput

func (TrafficWeightArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficWeight)(nil)).Elem()
}

func (i TrafficWeightArray) ToTrafficWeightArrayOutput() TrafficWeightArrayOutput {
	return i.ToTrafficWeightArrayOutputWithContext(context.Background())
}

func (i TrafficWeightArray) ToTrafficWeightArrayOutputWithContext(ctx context.Context) TrafficWeightArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficWeightArrayOutput)
}

type TrafficWeightOutput struct{ *pulumi.OutputState }

func (TrafficWeightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficWeight)(nil)).Elem()
}

func (o TrafficWeightOutput) ToTrafficWeightOutput() TrafficWeightOutput {
	return o
}

func (o TrafficWeightOutput) ToTrafficWeightOutputWithContext(ctx context.Context) TrafficWeightOutput {
	return o
}

func (o TrafficWeightOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficWeight) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o TrafficWeightOutput) LatestRevision() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TrafficWeight) *bool { return v.LatestRevision }).(pulumi.BoolPtrOutput)
}

func (o TrafficWeightOutput) RevisionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficWeight) *string { return v.RevisionName }).(pulumi.StringPtrOutput)
}

func (o TrafficWeightOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TrafficWeight) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type TrafficWeightArrayOutput struct{ *pulumi.OutputState }

func (TrafficWeightArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficWeight)(nil)).Elem()
}

func (o TrafficWeightArrayOutput) ToTrafficWeightArrayOutput() TrafficWeightArrayOutput {
	return o
}

func (o TrafficWeightArrayOutput) ToTrafficWeightArrayOutputWithContext(ctx context.Context) TrafficWeightArrayOutput {
	return o
}

func (o TrafficWeightArrayOutput) Index(i pulumi.IntInput) TrafficWeightOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrafficWeight {
		return vs[0].([]TrafficWeight)[vs[1].(int)]
	}).(TrafficWeightOutput)
}

type TrafficWeightResponse struct {
	Label          *string `pulumi:"label"`
	LatestRevision *bool   `pulumi:"latestRevision"`
	RevisionName   *string `pulumi:"revisionName"`
	Weight         *int    `pulumi:"weight"`
}


func (val *TrafficWeightResponse) Defaults() *TrafficWeightResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LatestRevision) {
		latestRevision_ := false
		tmp.LatestRevision = &latestRevision_
	}
	return &tmp
}

type TrafficWeightResponseOutput struct{ *pulumi.OutputState }

func (TrafficWeightResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficWeightResponse)(nil)).Elem()
}

func (o TrafficWeightResponseOutput) ToTrafficWeightResponseOutput() TrafficWeightResponseOutput {
	return o
}

func (o TrafficWeightResponseOutput) ToTrafficWeightResponseOutputWithContext(ctx context.Context) TrafficWeightResponseOutput {
	return o
}

func (o TrafficWeightResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficWeightResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o TrafficWeightResponseOutput) LatestRevision() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TrafficWeightResponse) *bool { return v.LatestRevision }).(pulumi.BoolPtrOutput)
}

func (o TrafficWeightResponseOutput) RevisionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficWeightResponse) *string { return v.RevisionName }).(pulumi.StringPtrOutput)
}

func (o TrafficWeightResponseOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TrafficWeightResponse) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type TrafficWeightResponseArrayOutput struct{ *pulumi.OutputState }

func (TrafficWeightResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficWeightResponse)(nil)).Elem()
}

func (o TrafficWeightResponseArrayOutput) ToTrafficWeightResponseArrayOutput() TrafficWeightResponseArrayOutput {
	return o
}

func (o TrafficWeightResponseArrayOutput) ToTrafficWeightResponseArrayOutputWithContext(ctx context.Context) TrafficWeightResponseArrayOutput {
	return o
}

func (o TrafficWeightResponseArrayOutput) Index(i pulumi.IntInput) TrafficWeightResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrafficWeightResponse {
		return vs[0].([]TrafficWeightResponse)[vs[1].(int)]
	}).(TrafficWeightResponseOutput)
}

type Twitter struct {
	Enabled      *bool                `pulumi:"enabled"`
	Registration *TwitterRegistration `pulumi:"registration"`
}





type TwitterInput interface {
	pulumi.Input

	ToTwitterOutput() TwitterOutput
	ToTwitterOutputWithContext(context.Context) TwitterOutput
}

type TwitterArgs struct {
	Enabled      pulumi.BoolPtrInput         `pulumi:"enabled"`
	Registration TwitterRegistrationPtrInput `pulumi:"registration"`
}

func (TwitterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Twitter)(nil)).Elem()
}

func (i TwitterArgs) ToTwitterOutput() TwitterOutput {
	return i.ToTwitterOutputWithContext(context.Background())
}

func (i TwitterArgs) ToTwitterOutputWithContext(ctx context.Context) TwitterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterOutput)
}

func (i TwitterArgs) ToTwitterPtrOutput() TwitterPtrOutput {
	return i.ToTwitterPtrOutputWithContext(context.Background())
}

func (i TwitterArgs) ToTwitterPtrOutputWithContext(ctx context.Context) TwitterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterOutput).ToTwitterPtrOutputWithContext(ctx)
}









type TwitterPtrInput interface {
	pulumi.Input

	ToTwitterPtrOutput() TwitterPtrOutput
	ToTwitterPtrOutputWithContext(context.Context) TwitterPtrOutput
}

type twitterPtrType TwitterArgs

func TwitterPtr(v *TwitterArgs) TwitterPtrInput {
	return (*twitterPtrType)(v)
}

func (*twitterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Twitter)(nil)).Elem()
}

func (i *twitterPtrType) ToTwitterPtrOutput() TwitterPtrOutput {
	return i.ToTwitterPtrOutputWithContext(context.Background())
}

func (i *twitterPtrType) ToTwitterPtrOutputWithContext(ctx context.Context) TwitterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterPtrOutput)
}

type TwitterOutput struct{ *pulumi.OutputState }

func (TwitterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Twitter)(nil)).Elem()
}

func (o TwitterOutput) ToTwitterOutput() TwitterOutput {
	return o
}

func (o TwitterOutput) ToTwitterOutputWithContext(ctx context.Context) TwitterOutput {
	return o
}

func (o TwitterOutput) ToTwitterPtrOutput() TwitterPtrOutput {
	return o.ToTwitterPtrOutputWithContext(context.Background())
}

func (o TwitterOutput) ToTwitterPtrOutputWithContext(ctx context.Context) TwitterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Twitter) *Twitter {
		return &v
	}).(TwitterPtrOutput)
}

func (o TwitterOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Twitter) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TwitterOutput) Registration() TwitterRegistrationPtrOutput {
	return o.ApplyT(func(v Twitter) *TwitterRegistration { return v.Registration }).(TwitterRegistrationPtrOutput)
}

type TwitterPtrOutput struct{ *pulumi.OutputState }

func (TwitterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Twitter)(nil)).Elem()
}

func (o TwitterPtrOutput) ToTwitterPtrOutput() TwitterPtrOutput {
	return o
}

func (o TwitterPtrOutput) ToTwitterPtrOutputWithContext(ctx context.Context) TwitterPtrOutput {
	return o
}

func (o TwitterPtrOutput) Elem() TwitterOutput {
	return o.ApplyT(func(v *Twitter) Twitter {
		if v != nil {
			return *v
		}
		var ret Twitter
		return ret
	}).(TwitterOutput)
}

func (o TwitterPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Twitter) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TwitterPtrOutput) Registration() TwitterRegistrationPtrOutput {
	return o.ApplyT(func(v *Twitter) *TwitterRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(TwitterRegistrationPtrOutput)
}

type TwitterRegistration struct {
	ConsumerKey               *string `pulumi:"consumerKey"`
	ConsumerSecretSettingName *string `pulumi:"consumerSecretSettingName"`
}





type TwitterRegistrationInput interface {
	pulumi.Input

	ToTwitterRegistrationOutput() TwitterRegistrationOutput
	ToTwitterRegistrationOutputWithContext(context.Context) TwitterRegistrationOutput
}

type TwitterRegistrationArgs struct {
	ConsumerKey               pulumi.StringPtrInput `pulumi:"consumerKey"`
	ConsumerSecretSettingName pulumi.StringPtrInput `pulumi:"consumerSecretSettingName"`
}

func (TwitterRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TwitterRegistration)(nil)).Elem()
}

func (i TwitterRegistrationArgs) ToTwitterRegistrationOutput() TwitterRegistrationOutput {
	return i.ToTwitterRegistrationOutputWithContext(context.Background())
}

func (i TwitterRegistrationArgs) ToTwitterRegistrationOutputWithContext(ctx context.Context) TwitterRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterRegistrationOutput)
}

func (i TwitterRegistrationArgs) ToTwitterRegistrationPtrOutput() TwitterRegistrationPtrOutput {
	return i.ToTwitterRegistrationPtrOutputWithContext(context.Background())
}

func (i TwitterRegistrationArgs) ToTwitterRegistrationPtrOutputWithContext(ctx context.Context) TwitterRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterRegistrationOutput).ToTwitterRegistrationPtrOutputWithContext(ctx)
}









type TwitterRegistrationPtrInput interface {
	pulumi.Input

	ToTwitterRegistrationPtrOutput() TwitterRegistrationPtrOutput
	ToTwitterRegistrationPtrOutputWithContext(context.Context) TwitterRegistrationPtrOutput
}

type twitterRegistrationPtrType TwitterRegistrationArgs

func TwitterRegistrationPtr(v *TwitterRegistrationArgs) TwitterRegistrationPtrInput {
	return (*twitterRegistrationPtrType)(v)
}

func (*twitterRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TwitterRegistration)(nil)).Elem()
}

func (i *twitterRegistrationPtrType) ToTwitterRegistrationPtrOutput() TwitterRegistrationPtrOutput {
	return i.ToTwitterRegistrationPtrOutputWithContext(context.Background())
}

func (i *twitterRegistrationPtrType) ToTwitterRegistrationPtrOutputWithContext(ctx context.Context) TwitterRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterRegistrationPtrOutput)
}

type TwitterRegistrationOutput struct{ *pulumi.OutputState }

func (TwitterRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TwitterRegistration)(nil)).Elem()
}

func (o TwitterRegistrationOutput) ToTwitterRegistrationOutput() TwitterRegistrationOutput {
	return o
}

func (o TwitterRegistrationOutput) ToTwitterRegistrationOutputWithContext(ctx context.Context) TwitterRegistrationOutput {
	return o
}

func (o TwitterRegistrationOutput) ToTwitterRegistrationPtrOutput() TwitterRegistrationPtrOutput {
	return o.ToTwitterRegistrationPtrOutputWithContext(context.Background())
}

func (o TwitterRegistrationOutput) ToTwitterRegistrationPtrOutputWithContext(ctx context.Context) TwitterRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TwitterRegistration) *TwitterRegistration {
		return &v
	}).(TwitterRegistrationPtrOutput)
}

func (o TwitterRegistrationOutput) ConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TwitterRegistration) *string { return v.ConsumerKey }).(pulumi.StringPtrOutput)
}

func (o TwitterRegistrationOutput) ConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TwitterRegistration) *string { return v.ConsumerSecretSettingName }).(pulumi.StringPtrOutput)
}

type TwitterRegistrationPtrOutput struct{ *pulumi.OutputState }

func (TwitterRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwitterRegistration)(nil)).Elem()
}

func (o TwitterRegistrationPtrOutput) ToTwitterRegistrationPtrOutput() TwitterRegistrationPtrOutput {
	return o
}

func (o TwitterRegistrationPtrOutput) ToTwitterRegistrationPtrOutputWithContext(ctx context.Context) TwitterRegistrationPtrOutput {
	return o
}

func (o TwitterRegistrationPtrOutput) Elem() TwitterRegistrationOutput {
	return o.ApplyT(func(v *TwitterRegistration) TwitterRegistration {
		if v != nil {
			return *v
		}
		var ret TwitterRegistration
		return ret
	}).(TwitterRegistrationOutput)
}

func (o TwitterRegistrationPtrOutput) ConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TwitterRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ConsumerKey
	}).(pulumi.StringPtrOutput)
}

func (o TwitterRegistrationPtrOutput) ConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TwitterRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ConsumerSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type TwitterRegistrationResponse struct {
	ConsumerKey               *string `pulumi:"consumerKey"`
	ConsumerSecretSettingName *string `pulumi:"consumerSecretSettingName"`
}

type TwitterRegistrationResponseOutput struct{ *pulumi.OutputState }

func (TwitterRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TwitterRegistrationResponse)(nil)).Elem()
}

func (o TwitterRegistrationResponseOutput) ToTwitterRegistrationResponseOutput() TwitterRegistrationResponseOutput {
	return o
}

func (o TwitterRegistrationResponseOutput) ToTwitterRegistrationResponseOutputWithContext(ctx context.Context) TwitterRegistrationResponseOutput {
	return o
}

func (o TwitterRegistrationResponseOutput) ConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TwitterRegistrationResponse) *string { return v.ConsumerKey }).(pulumi.StringPtrOutput)
}

func (o TwitterRegistrationResponseOutput) ConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TwitterRegistrationResponse) *string { return v.ConsumerSecretSettingName }).(pulumi.StringPtrOutput)
}

type TwitterRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (TwitterRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwitterRegistrationResponse)(nil)).Elem()
}

func (o TwitterRegistrationResponsePtrOutput) ToTwitterRegistrationResponsePtrOutput() TwitterRegistrationResponsePtrOutput {
	return o
}

func (o TwitterRegistrationResponsePtrOutput) ToTwitterRegistrationResponsePtrOutputWithContext(ctx context.Context) TwitterRegistrationResponsePtrOutput {
	return o
}

func (o TwitterRegistrationResponsePtrOutput) Elem() TwitterRegistrationResponseOutput {
	return o.ApplyT(func(v *TwitterRegistrationResponse) TwitterRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret TwitterRegistrationResponse
		return ret
	}).(TwitterRegistrationResponseOutput)
}

func (o TwitterRegistrationResponsePtrOutput) ConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TwitterRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConsumerKey
	}).(pulumi.StringPtrOutput)
}

func (o TwitterRegistrationResponsePtrOutput) ConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TwitterRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConsumerSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type TwitterResponse struct {
	Enabled      *bool                        `pulumi:"enabled"`
	Registration *TwitterRegistrationResponse `pulumi:"registration"`
}

type TwitterResponseOutput struct{ *pulumi.OutputState }

func (TwitterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TwitterResponse)(nil)).Elem()
}

func (o TwitterResponseOutput) ToTwitterResponseOutput() TwitterResponseOutput {
	return o
}

func (o TwitterResponseOutput) ToTwitterResponseOutputWithContext(ctx context.Context) TwitterResponseOutput {
	return o
}

func (o TwitterResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TwitterResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TwitterResponseOutput) Registration() TwitterRegistrationResponsePtrOutput {
	return o.ApplyT(func(v TwitterResponse) *TwitterRegistrationResponse { return v.Registration }).(TwitterRegistrationResponsePtrOutput)
}

type TwitterResponsePtrOutput struct{ *pulumi.OutputState }

func (TwitterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwitterResponse)(nil)).Elem()
}

func (o TwitterResponsePtrOutput) ToTwitterResponsePtrOutput() TwitterResponsePtrOutput {
	return o
}

func (o TwitterResponsePtrOutput) ToTwitterResponsePtrOutputWithContext(ctx context.Context) TwitterResponsePtrOutput {
	return o
}

func (o TwitterResponsePtrOutput) Elem() TwitterResponseOutput {
	return o.ApplyT(func(v *TwitterResponse) TwitterResponse {
		if v != nil {
			return *v
		}
		var ret TwitterResponse
		return ret
	}).(TwitterResponseOutput)
}

func (o TwitterResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TwitterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TwitterResponsePtrOutput) Registration() TwitterRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *TwitterResponse) *TwitterRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(TwitterRegistrationResponsePtrOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type VnetConfiguration struct {
	DockerBridgeCidr       *string `pulumi:"dockerBridgeCidr"`
	InfrastructureSubnetId *string `pulumi:"infrastructureSubnetId"`
	Internal               *bool   `pulumi:"internal"`
	PlatformReservedCidr   *string `pulumi:"platformReservedCidr"`
	PlatformReservedDnsIP  *string `pulumi:"platformReservedDnsIP"`
	RuntimeSubnetId        *string `pulumi:"runtimeSubnetId"`
}





type VnetConfigurationInput interface {
	pulumi.Input

	ToVnetConfigurationOutput() VnetConfigurationOutput
	ToVnetConfigurationOutputWithContext(context.Context) VnetConfigurationOutput
}

type VnetConfigurationArgs struct {
	DockerBridgeCidr       pulumi.StringPtrInput `pulumi:"dockerBridgeCidr"`
	InfrastructureSubnetId pulumi.StringPtrInput `pulumi:"infrastructureSubnetId"`
	Internal               pulumi.BoolPtrInput   `pulumi:"internal"`
	PlatformReservedCidr   pulumi.StringPtrInput `pulumi:"platformReservedCidr"`
	PlatformReservedDnsIP  pulumi.StringPtrInput `pulumi:"platformReservedDnsIP"`
	RuntimeSubnetId        pulumi.StringPtrInput `pulumi:"runtimeSubnetId"`
}

func (VnetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetConfiguration)(nil)).Elem()
}

func (i VnetConfigurationArgs) ToVnetConfigurationOutput() VnetConfigurationOutput {
	return i.ToVnetConfigurationOutputWithContext(context.Background())
}

func (i VnetConfigurationArgs) ToVnetConfigurationOutputWithContext(ctx context.Context) VnetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetConfigurationOutput)
}

func (i VnetConfigurationArgs) ToVnetConfigurationPtrOutput() VnetConfigurationPtrOutput {
	return i.ToVnetConfigurationPtrOutputWithContext(context.Background())
}

func (i VnetConfigurationArgs) ToVnetConfigurationPtrOutputWithContext(ctx context.Context) VnetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetConfigurationOutput).ToVnetConfigurationPtrOutputWithContext(ctx)
}









type VnetConfigurationPtrInput interface {
	pulumi.Input

	ToVnetConfigurationPtrOutput() VnetConfigurationPtrOutput
	ToVnetConfigurationPtrOutputWithContext(context.Context) VnetConfigurationPtrOutput
}

type vnetConfigurationPtrType VnetConfigurationArgs

func VnetConfigurationPtr(v *VnetConfigurationArgs) VnetConfigurationPtrInput {
	return (*vnetConfigurationPtrType)(v)
}

func (*vnetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VnetConfiguration)(nil)).Elem()
}

func (i *vnetConfigurationPtrType) ToVnetConfigurationPtrOutput() VnetConfigurationPtrOutput {
	return i.ToVnetConfigurationPtrOutputWithContext(context.Background())
}

func (i *vnetConfigurationPtrType) ToVnetConfigurationPtrOutputWithContext(ctx context.Context) VnetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetConfigurationPtrOutput)
}

type VnetConfigurationOutput struct{ *pulumi.OutputState }

func (VnetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetConfiguration)(nil)).Elem()
}

func (o VnetConfigurationOutput) ToVnetConfigurationOutput() VnetConfigurationOutput {
	return o
}

func (o VnetConfigurationOutput) ToVnetConfigurationOutputWithContext(ctx context.Context) VnetConfigurationOutput {
	return o
}

func (o VnetConfigurationOutput) ToVnetConfigurationPtrOutput() VnetConfigurationPtrOutput {
	return o.ToVnetConfigurationPtrOutputWithContext(context.Background())
}

func (o VnetConfigurationOutput) ToVnetConfigurationPtrOutputWithContext(ctx context.Context) VnetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VnetConfiguration) *VnetConfiguration {
		return &v
	}).(VnetConfigurationPtrOutput)
}

func (o VnetConfigurationOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetConfiguration) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationOutput) InfrastructureSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetConfiguration) *string { return v.InfrastructureSubnetId }).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationOutput) Internal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VnetConfiguration) *bool { return v.Internal }).(pulumi.BoolPtrOutput)
}

func (o VnetConfigurationOutput) PlatformReservedCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetConfiguration) *string { return v.PlatformReservedCidr }).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationOutput) PlatformReservedDnsIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetConfiguration) *string { return v.PlatformReservedDnsIP }).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationOutput) RuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetConfiguration) *string { return v.RuntimeSubnetId }).(pulumi.StringPtrOutput)
}

type VnetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VnetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VnetConfiguration)(nil)).Elem()
}

func (o VnetConfigurationPtrOutput) ToVnetConfigurationPtrOutput() VnetConfigurationPtrOutput {
	return o
}

func (o VnetConfigurationPtrOutput) ToVnetConfigurationPtrOutputWithContext(ctx context.Context) VnetConfigurationPtrOutput {
	return o
}

func (o VnetConfigurationPtrOutput) Elem() VnetConfigurationOutput {
	return o.ApplyT(func(v *VnetConfiguration) VnetConfiguration {
		if v != nil {
			return *v
		}
		var ret VnetConfiguration
		return ret
	}).(VnetConfigurationOutput)
}

func (o VnetConfigurationPtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VnetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationPtrOutput) InfrastructureSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VnetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.InfrastructureSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationPtrOutput) Internal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VnetConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.Internal
	}).(pulumi.BoolPtrOutput)
}

func (o VnetConfigurationPtrOutput) PlatformReservedCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VnetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PlatformReservedCidr
	}).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationPtrOutput) PlatformReservedDnsIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VnetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PlatformReservedDnsIP
	}).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationPtrOutput) RuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VnetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeSubnetId
	}).(pulumi.StringPtrOutput)
}

type VnetConfigurationResponse struct {
	DockerBridgeCidr       *string `pulumi:"dockerBridgeCidr"`
	InfrastructureSubnetId *string `pulumi:"infrastructureSubnetId"`
	Internal               *bool   `pulumi:"internal"`
	PlatformReservedCidr   *string `pulumi:"platformReservedCidr"`
	PlatformReservedDnsIP  *string `pulumi:"platformReservedDnsIP"`
	RuntimeSubnetId        *string `pulumi:"runtimeSubnetId"`
}

type VnetConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VnetConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetConfigurationResponse)(nil)).Elem()
}

func (o VnetConfigurationResponseOutput) ToVnetConfigurationResponseOutput() VnetConfigurationResponseOutput {
	return o
}

func (o VnetConfigurationResponseOutput) ToVnetConfigurationResponseOutputWithContext(ctx context.Context) VnetConfigurationResponseOutput {
	return o
}

func (o VnetConfigurationResponseOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetConfigurationResponse) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationResponseOutput) InfrastructureSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetConfigurationResponse) *string { return v.InfrastructureSubnetId }).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationResponseOutput) Internal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VnetConfigurationResponse) *bool { return v.Internal }).(pulumi.BoolPtrOutput)
}

func (o VnetConfigurationResponseOutput) PlatformReservedCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetConfigurationResponse) *string { return v.PlatformReservedCidr }).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationResponseOutput) PlatformReservedDnsIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetConfigurationResponse) *string { return v.PlatformReservedDnsIP }).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationResponseOutput) RuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetConfigurationResponse) *string { return v.RuntimeSubnetId }).(pulumi.StringPtrOutput)
}

type VnetConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VnetConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VnetConfigurationResponse)(nil)).Elem()
}

func (o VnetConfigurationResponsePtrOutput) ToVnetConfigurationResponsePtrOutput() VnetConfigurationResponsePtrOutput {
	return o
}

func (o VnetConfigurationResponsePtrOutput) ToVnetConfigurationResponsePtrOutputWithContext(ctx context.Context) VnetConfigurationResponsePtrOutput {
	return o
}

func (o VnetConfigurationResponsePtrOutput) Elem() VnetConfigurationResponseOutput {
	return o.ApplyT(func(v *VnetConfigurationResponse) VnetConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VnetConfigurationResponse
		return ret
	}).(VnetConfigurationResponseOutput)
}

func (o VnetConfigurationResponsePtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VnetConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationResponsePtrOutput) InfrastructureSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VnetConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.InfrastructureSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationResponsePtrOutput) Internal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VnetConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Internal
	}).(pulumi.BoolPtrOutput)
}

func (o VnetConfigurationResponsePtrOutput) PlatformReservedCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VnetConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PlatformReservedCidr
	}).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationResponsePtrOutput) PlatformReservedDnsIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VnetConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PlatformReservedDnsIP
	}).(pulumi.StringPtrOutput)
}

func (o VnetConfigurationResponsePtrOutput) RuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VnetConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeSubnetId
	}).(pulumi.StringPtrOutput)
}

type Volume struct {
	Name        *string `pulumi:"name"`
	StorageName *string `pulumi:"storageName"`
	StorageType *string `pulumi:"storageType"`
}





type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(context.Context) VolumeOutput
}

type VolumeArgs struct {
	Name        pulumi.StringPtrInput `pulumi:"name"`
	StorageName pulumi.StringPtrInput `pulumi:"storageName"`
	StorageType pulumi.StringPtrInput `pulumi:"storageType"`
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Volume)(nil)).Elem()
}

func (i VolumeArgs) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i VolumeArgs) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}





type VolumeArrayInput interface {
	pulumi.Input

	ToVolumeArrayOutput() VolumeArrayOutput
	ToVolumeArrayOutputWithContext(context.Context) VolumeArrayOutput
}

type VolumeArray []VolumeInput

func (VolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Volume)(nil)).Elem()
}

func (i VolumeArray) ToVolumeArrayOutput() VolumeArrayOutput {
	return i.ToVolumeArrayOutputWithContext(context.Background())
}

func (i VolumeArray) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeArrayOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

func (o VolumeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Volume) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) StorageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Volume) *string { return v.StorageName }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Volume) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

type VolumeArrayOutput struct{ *pulumi.OutputState }

func (VolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Volume)(nil)).Elem()
}

func (o VolumeArrayOutput) ToVolumeArrayOutput() VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) Index(i pulumi.IntInput) VolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Volume {
		return vs[0].([]Volume)[vs[1].(int)]
	}).(VolumeOutput)
}

type VolumeMount struct {
	MountPath  *string `pulumi:"mountPath"`
	VolumeName *string `pulumi:"volumeName"`
}





type VolumeMountInput interface {
	pulumi.Input

	ToVolumeMountOutput() VolumeMountOutput
	ToVolumeMountOutputWithContext(context.Context) VolumeMountOutput
}

type VolumeMountArgs struct {
	MountPath  pulumi.StringPtrInput `pulumi:"mountPath"`
	VolumeName pulumi.StringPtrInput `pulumi:"volumeName"`
}

func (VolumeMountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeMount)(nil)).Elem()
}

func (i VolumeMountArgs) ToVolumeMountOutput() VolumeMountOutput {
	return i.ToVolumeMountOutputWithContext(context.Background())
}

func (i VolumeMountArgs) ToVolumeMountOutputWithContext(ctx context.Context) VolumeMountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMountOutput)
}





type VolumeMountArrayInput interface {
	pulumi.Input

	ToVolumeMountArrayOutput() VolumeMountArrayOutput
	ToVolumeMountArrayOutputWithContext(context.Context) VolumeMountArrayOutput
}

type VolumeMountArray []VolumeMountInput

func (VolumeMountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeMount)(nil)).Elem()
}

func (i VolumeMountArray) ToVolumeMountArrayOutput() VolumeMountArrayOutput {
	return i.ToVolumeMountArrayOutputWithContext(context.Background())
}

func (i VolumeMountArray) ToVolumeMountArrayOutputWithContext(ctx context.Context) VolumeMountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMountArrayOutput)
}

type VolumeMountOutput struct{ *pulumi.OutputState }

func (VolumeMountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeMount)(nil)).Elem()
}

func (o VolumeMountOutput) ToVolumeMountOutput() VolumeMountOutput {
	return o
}

func (o VolumeMountOutput) ToVolumeMountOutputWithContext(ctx context.Context) VolumeMountOutput {
	return o
}

func (o VolumeMountOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeMount) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o VolumeMountOutput) VolumeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeMount) *string { return v.VolumeName }).(pulumi.StringPtrOutput)
}

type VolumeMountArrayOutput struct{ *pulumi.OutputState }

func (VolumeMountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeMount)(nil)).Elem()
}

func (o VolumeMountArrayOutput) ToVolumeMountArrayOutput() VolumeMountArrayOutput {
	return o
}

func (o VolumeMountArrayOutput) ToVolumeMountArrayOutputWithContext(ctx context.Context) VolumeMountArrayOutput {
	return o
}

func (o VolumeMountArrayOutput) Index(i pulumi.IntInput) VolumeMountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeMount {
		return vs[0].([]VolumeMount)[vs[1].(int)]
	}).(VolumeMountOutput)
}

type VolumeMountResponse struct {
	MountPath  *string `pulumi:"mountPath"`
	VolumeName *string `pulumi:"volumeName"`
}

type VolumeMountResponseOutput struct{ *pulumi.OutputState }

func (VolumeMountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeMountResponse)(nil)).Elem()
}

func (o VolumeMountResponseOutput) ToVolumeMountResponseOutput() VolumeMountResponseOutput {
	return o
}

func (o VolumeMountResponseOutput) ToVolumeMountResponseOutputWithContext(ctx context.Context) VolumeMountResponseOutput {
	return o
}

func (o VolumeMountResponseOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeMountResponse) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o VolumeMountResponseOutput) VolumeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeMountResponse) *string { return v.VolumeName }).(pulumi.StringPtrOutput)
}

type VolumeMountResponseArrayOutput struct{ *pulumi.OutputState }

func (VolumeMountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeMountResponse)(nil)).Elem()
}

func (o VolumeMountResponseArrayOutput) ToVolumeMountResponseArrayOutput() VolumeMountResponseArrayOutput {
	return o
}

func (o VolumeMountResponseArrayOutput) ToVolumeMountResponseArrayOutputWithContext(ctx context.Context) VolumeMountResponseArrayOutput {
	return o
}

func (o VolumeMountResponseArrayOutput) Index(i pulumi.IntInput) VolumeMountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeMountResponse {
		return vs[0].([]VolumeMountResponse)[vs[1].(int)]
	}).(VolumeMountResponseOutput)
}

type VolumeResponse struct {
	Name        *string `pulumi:"name"`
	StorageName *string `pulumi:"storageName"`
	StorageType *string `pulumi:"storageType"`
}

type VolumeResponseOutput struct{ *pulumi.OutputState }

func (VolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeResponse)(nil)).Elem()
}

func (o VolumeResponseOutput) ToVolumeResponseOutput() VolumeResponseOutput {
	return o
}

func (o VolumeResponseOutput) ToVolumeResponseOutputWithContext(ctx context.Context) VolumeResponseOutput {
	return o
}

func (o VolumeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VolumeResponseOutput) StorageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeResponse) *string { return v.StorageName }).(pulumi.StringPtrOutput)
}

func (o VolumeResponseOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeResponse) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

type VolumeResponseArrayOutput struct{ *pulumi.OutputState }

func (VolumeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeResponse)(nil)).Elem()
}

func (o VolumeResponseArrayOutput) ToVolumeResponseArrayOutput() VolumeResponseArrayOutput {
	return o
}

func (o VolumeResponseArrayOutput) ToVolumeResponseArrayOutputWithContext(ctx context.Context) VolumeResponseArrayOutput {
	return o
}

func (o VolumeResponseArrayOutput) Index(i pulumi.IntInput) VolumeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeResponse {
		return vs[0].([]VolumeResponse)[vs[1].(int)]
	}).(VolumeResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AllowedAudiencesValidationOutput{})
	pulumi.RegisterOutputType(AllowedAudiencesValidationPtrOutput{})
	pulumi.RegisterOutputType(AllowedAudiencesValidationResponseOutput{})
	pulumi.RegisterOutputType(AllowedAudiencesValidationResponsePtrOutput{})
	pulumi.RegisterOutputType(AllowedPrincipalsOutput{})
	pulumi.RegisterOutputType(AllowedPrincipalsPtrOutput{})
	pulumi.RegisterOutputType(AllowedPrincipalsResponseOutput{})
	pulumi.RegisterOutputType(AllowedPrincipalsResponsePtrOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(AppRegistrationOutput{})
	pulumi.RegisterOutputType(AppRegistrationPtrOutput{})
	pulumi.RegisterOutputType(AppRegistrationResponseOutput{})
	pulumi.RegisterOutputType(AppRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(AppleOutput{})
	pulumi.RegisterOutputType(ApplePtrOutput{})
	pulumi.RegisterOutputType(AppleRegistrationOutput{})
	pulumi.RegisterOutputType(AppleRegistrationPtrOutput{})
	pulumi.RegisterOutputType(AppleRegistrationResponseOutput{})
	pulumi.RegisterOutputType(AppleRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(AppleResponseOutput{})
	pulumi.RegisterOutputType(AppleResponsePtrOutput{})
	pulumi.RegisterOutputType(AuthPlatformOutput{})
	pulumi.RegisterOutputType(AuthPlatformPtrOutput{})
	pulumi.RegisterOutputType(AuthPlatformResponseOutput{})
	pulumi.RegisterOutputType(AuthPlatformResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryLoginOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryLoginPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryLoginResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryLoginResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryRegistrationOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryRegistrationPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryRegistrationResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryValidationOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryValidationPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryValidationResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryValidationResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureCredentialsOutput{})
	pulumi.RegisterOutputType(AzureCredentialsPtrOutput{})
	pulumi.RegisterOutputType(AzureCredentialsResponseOutput{})
	pulumi.RegisterOutputType(AzureCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureFilePropertiesOutput{})
	pulumi.RegisterOutputType(AzureFilePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AzureFilePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AzureFilePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsPtrOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsRegistrationOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsRegistrationPtrOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsRegistrationResponseOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsResponseOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsResponsePtrOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CertificateResponsePropertiesOutput{})
	pulumi.RegisterOutputType(ClientRegistrationOutput{})
	pulumi.RegisterOutputType(ClientRegistrationPtrOutput{})
	pulumi.RegisterOutputType(ClientRegistrationResponseOutput{})
	pulumi.RegisterOutputType(ClientRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationOutput{})
	pulumi.RegisterOutputType(ConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerArrayOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeArrayOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeHttpGetOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeHttpGetPtrOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeHttpHeadersOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeHttpHeadersArrayOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeResponseOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeResponseHttpGetOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeResponseHttpGetPtrOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeResponseHttpHeadersOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeResponseHttpHeadersArrayOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeResponseTcpSocketOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeResponseTcpSocketPtrOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeTcpSocketOutput{})
	pulumi.RegisterOutputType(ContainerAppProbeTcpSocketPtrOutput{})
	pulumi.RegisterOutputType(ContainerAppSecretResponseOutput{})
	pulumi.RegisterOutputType(ContainerAppSecretResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerResourcesOutput{})
	pulumi.RegisterOutputType(ContainerResourcesPtrOutput{})
	pulumi.RegisterOutputType(ContainerResourcesResponseOutput{})
	pulumi.RegisterOutputType(ContainerResourcesResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerResponseOutput{})
	pulumi.RegisterOutputType(ContainerResponseArrayOutput{})
	pulumi.RegisterOutputType(CookieExpirationOutput{})
	pulumi.RegisterOutputType(CookieExpirationPtrOutput{})
	pulumi.RegisterOutputType(CookieExpirationResponseOutput{})
	pulumi.RegisterOutputType(CookieExpirationResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomDomainOutput{})
	pulumi.RegisterOutputType(CustomDomainArrayOutput{})
	pulumi.RegisterOutputType(CustomDomainResponseOutput{})
	pulumi.RegisterOutputType(CustomDomainResponseArrayOutput{})
	pulumi.RegisterOutputType(CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput{})
	pulumi.RegisterOutputType(CustomHostnameAnalysisResultResponseDetailsOutput{})
	pulumi.RegisterOutputType(CustomHostnameAnalysisResultResponseDetailsArrayOutput{})
	pulumi.RegisterOutputType(CustomOpenIdConnectProviderOutput{})
	pulumi.RegisterOutputType(CustomOpenIdConnectProviderMapOutput{})
	pulumi.RegisterOutputType(CustomOpenIdConnectProviderResponseOutput{})
	pulumi.RegisterOutputType(CustomOpenIdConnectProviderResponseMapOutput{})
	pulumi.RegisterOutputType(CustomScaleRuleOutput{})
	pulumi.RegisterOutputType(CustomScaleRulePtrOutput{})
	pulumi.RegisterOutputType(CustomScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(CustomScaleRuleResponsePtrOutput{})
	pulumi.RegisterOutputType(DaprOutput{})
	pulumi.RegisterOutputType(DaprPtrOutput{})
	pulumi.RegisterOutputType(DaprMetadataOutput{})
	pulumi.RegisterOutputType(DaprMetadataArrayOutput{})
	pulumi.RegisterOutputType(DaprMetadataResponseOutput{})
	pulumi.RegisterOutputType(DaprMetadataResponseArrayOutput{})
	pulumi.RegisterOutputType(DaprResponseOutput{})
	pulumi.RegisterOutputType(DaprResponsePtrOutput{})
	pulumi.RegisterOutputType(DaprSecretResponseOutput{})
	pulumi.RegisterOutputType(DaprSecretResponseArrayOutput{})
	pulumi.RegisterOutputType(DefaultAuthorizationPolicyOutput{})
	pulumi.RegisterOutputType(DefaultAuthorizationPolicyPtrOutput{})
	pulumi.RegisterOutputType(DefaultAuthorizationPolicyResponseOutput{})
	pulumi.RegisterOutputType(DefaultAuthorizationPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentVarOutput{})
	pulumi.RegisterOutputType(EnvironmentVarArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVarResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentVarResponseArrayOutput{})
	pulumi.RegisterOutputType(FacebookOutput{})
	pulumi.RegisterOutputType(FacebookPtrOutput{})
	pulumi.RegisterOutputType(FacebookResponseOutput{})
	pulumi.RegisterOutputType(FacebookResponsePtrOutput{})
	pulumi.RegisterOutputType(ForwardProxyOutput{})
	pulumi.RegisterOutputType(ForwardProxyPtrOutput{})
	pulumi.RegisterOutputType(ForwardProxyResponseOutput{})
	pulumi.RegisterOutputType(ForwardProxyResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubOutput{})
	pulumi.RegisterOutputType(GitHubPtrOutput{})
	pulumi.RegisterOutputType(GitHubResponseOutput{})
	pulumi.RegisterOutputType(GitHubResponsePtrOutput{})
	pulumi.RegisterOutputType(GithubActionConfigurationOutput{})
	pulumi.RegisterOutputType(GithubActionConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GithubActionConfigurationResponseOutput{})
	pulumi.RegisterOutputType(GithubActionConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(GlobalValidationOutput{})
	pulumi.RegisterOutputType(GlobalValidationPtrOutput{})
	pulumi.RegisterOutputType(GlobalValidationResponseOutput{})
	pulumi.RegisterOutputType(GlobalValidationResponsePtrOutput{})
	pulumi.RegisterOutputType(GoogleOutput{})
	pulumi.RegisterOutputType(GooglePtrOutput{})
	pulumi.RegisterOutputType(GoogleResponseOutput{})
	pulumi.RegisterOutputType(GoogleResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpScaleRuleOutput{})
	pulumi.RegisterOutputType(HttpScaleRulePtrOutput{})
	pulumi.RegisterOutputType(HttpScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(HttpScaleRuleResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpSettingsOutput{})
	pulumi.RegisterOutputType(HttpSettingsPtrOutput{})
	pulumi.RegisterOutputType(HttpSettingsResponseOutput{})
	pulumi.RegisterOutputType(HttpSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpSettingsRoutesOutput{})
	pulumi.RegisterOutputType(HttpSettingsRoutesPtrOutput{})
	pulumi.RegisterOutputType(HttpSettingsRoutesResponseOutput{})
	pulumi.RegisterOutputType(HttpSettingsRoutesResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityProvidersOutput{})
	pulumi.RegisterOutputType(IdentityProvidersPtrOutput{})
	pulumi.RegisterOutputType(IdentityProvidersResponseOutput{})
	pulumi.RegisterOutputType(IdentityProvidersResponsePtrOutput{})
	pulumi.RegisterOutputType(IngressOutput{})
	pulumi.RegisterOutputType(IngressPtrOutput{})
	pulumi.RegisterOutputType(IngressResponseOutput{})
	pulumi.RegisterOutputType(IngressResponsePtrOutput{})
	pulumi.RegisterOutputType(JwtClaimChecksOutput{})
	pulumi.RegisterOutputType(JwtClaimChecksPtrOutput{})
	pulumi.RegisterOutputType(JwtClaimChecksResponseOutput{})
	pulumi.RegisterOutputType(JwtClaimChecksResponsePtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(LoginOutput{})
	pulumi.RegisterOutputType(LoginPtrOutput{})
	pulumi.RegisterOutputType(LoginResponseOutput{})
	pulumi.RegisterOutputType(LoginResponsePtrOutput{})
	pulumi.RegisterOutputType(LoginRoutesOutput{})
	pulumi.RegisterOutputType(LoginRoutesPtrOutput{})
	pulumi.RegisterOutputType(LoginRoutesResponseOutput{})
	pulumi.RegisterOutputType(LoginRoutesResponsePtrOutput{})
	pulumi.RegisterOutputType(LoginScopesOutput{})
	pulumi.RegisterOutputType(LoginScopesPtrOutput{})
	pulumi.RegisterOutputType(LoginScopesResponseOutput{})
	pulumi.RegisterOutputType(LoginScopesResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedEnvironmentStoragePropertiesOutput{})
	pulumi.RegisterOutputType(ManagedEnvironmentStoragePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagedEnvironmentStorageResponsePropertiesOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(NonceOutput{})
	pulumi.RegisterOutputType(NoncePtrOutput{})
	pulumi.RegisterOutputType(NonceResponseOutput{})
	pulumi.RegisterOutputType(NonceResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectClientCredentialOutput{})
	pulumi.RegisterOutputType(OpenIdConnectClientCredentialPtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectClientCredentialResponseOutput{})
	pulumi.RegisterOutputType(OpenIdConnectClientCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectConfigOutput{})
	pulumi.RegisterOutputType(OpenIdConnectConfigPtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectConfigResponseOutput{})
	pulumi.RegisterOutputType(OpenIdConnectConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectLoginOutput{})
	pulumi.RegisterOutputType(OpenIdConnectLoginPtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectLoginResponseOutput{})
	pulumi.RegisterOutputType(OpenIdConnectLoginResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectRegistrationOutput{})
	pulumi.RegisterOutputType(OpenIdConnectRegistrationPtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectRegistrationResponseOutput{})
	pulumi.RegisterOutputType(OpenIdConnectRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(QueueScaleRuleOutput{})
	pulumi.RegisterOutputType(QueueScaleRulePtrOutput{})
	pulumi.RegisterOutputType(QueueScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(QueueScaleRuleResponsePtrOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsArrayOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsResponseOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsResponseArrayOutput{})
	pulumi.RegisterOutputType(RegistryInfoOutput{})
	pulumi.RegisterOutputType(RegistryInfoPtrOutput{})
	pulumi.RegisterOutputType(RegistryInfoResponseOutput{})
	pulumi.RegisterOutputType(RegistryInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ScaleOutput{})
	pulumi.RegisterOutputType(ScalePtrOutput{})
	pulumi.RegisterOutputType(ScaleResponseOutput{})
	pulumi.RegisterOutputType(ScaleResponsePtrOutput{})
	pulumi.RegisterOutputType(ScaleRuleOutput{})
	pulumi.RegisterOutputType(ScaleRuleArrayOutput{})
	pulumi.RegisterOutputType(ScaleRuleAuthOutput{})
	pulumi.RegisterOutputType(ScaleRuleAuthArrayOutput{})
	pulumi.RegisterOutputType(ScaleRuleAuthResponseOutput{})
	pulumi.RegisterOutputType(ScaleRuleAuthResponseArrayOutput{})
	pulumi.RegisterOutputType(ScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(ScaleRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SecretOutput{})
	pulumi.RegisterOutputType(SecretArrayOutput{})
	pulumi.RegisterOutputType(SecretResponseOutput{})
	pulumi.RegisterOutputType(SecretResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TemplateOutput{})
	pulumi.RegisterOutputType(TemplatePtrOutput{})
	pulumi.RegisterOutputType(TemplateResponseOutput{})
	pulumi.RegisterOutputType(TemplateResponsePtrOutput{})
	pulumi.RegisterOutputType(TrafficWeightOutput{})
	pulumi.RegisterOutputType(TrafficWeightArrayOutput{})
	pulumi.RegisterOutputType(TrafficWeightResponseOutput{})
	pulumi.RegisterOutputType(TrafficWeightResponseArrayOutput{})
	pulumi.RegisterOutputType(TwitterOutput{})
	pulumi.RegisterOutputType(TwitterPtrOutput{})
	pulumi.RegisterOutputType(TwitterRegistrationOutput{})
	pulumi.RegisterOutputType(TwitterRegistrationPtrOutput{})
	pulumi.RegisterOutputType(TwitterRegistrationResponseOutput{})
	pulumi.RegisterOutputType(TwitterRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(TwitterResponseOutput{})
	pulumi.RegisterOutputType(TwitterResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VnetConfigurationOutput{})
	pulumi.RegisterOutputType(VnetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VnetConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VnetConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VolumeOutput{})
	pulumi.RegisterOutputType(VolumeArrayOutput{})
	pulumi.RegisterOutputType(VolumeMountOutput{})
	pulumi.RegisterOutputType(VolumeMountArrayOutput{})
	pulumi.RegisterOutputType(VolumeMountResponseOutput{})
	pulumi.RegisterOutputType(VolumeMountResponseArrayOutput{})
	pulumi.RegisterOutputType(VolumeResponseOutput{})
	pulumi.RegisterOutputType(VolumeResponseArrayOutput{})
}
