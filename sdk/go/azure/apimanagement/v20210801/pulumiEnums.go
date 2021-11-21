


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlwaysLog string

const (
	// Always log all erroneous request regardless of sampling settings.
	AlwaysLogAllErrors = AlwaysLog("allErrors")
)

func (AlwaysLog) ElementType() reflect.Type {
	return reflect.TypeOf((*AlwaysLog)(nil)).Elem()
}

func (e AlwaysLog) ToAlwaysLogOutput() AlwaysLogOutput {
	return pulumi.ToOutput(e).(AlwaysLogOutput)
}

func (e AlwaysLog) ToAlwaysLogOutputWithContext(ctx context.Context) AlwaysLogOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AlwaysLogOutput)
}

func (e AlwaysLog) ToAlwaysLogPtrOutput() AlwaysLogPtrOutput {
	return e.ToAlwaysLogPtrOutputWithContext(context.Background())
}

func (e AlwaysLog) ToAlwaysLogPtrOutputWithContext(ctx context.Context) AlwaysLogPtrOutput {
	return AlwaysLog(e).ToAlwaysLogOutputWithContext(ctx).ToAlwaysLogPtrOutputWithContext(ctx)
}

func (e AlwaysLog) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlwaysLog) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlwaysLog) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AlwaysLog) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AlwaysLogOutput struct{ *pulumi.OutputState }

func (AlwaysLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlwaysLog)(nil)).Elem()
}

func (o AlwaysLogOutput) ToAlwaysLogOutput() AlwaysLogOutput {
	return o
}

func (o AlwaysLogOutput) ToAlwaysLogOutputWithContext(ctx context.Context) AlwaysLogOutput {
	return o
}

func (o AlwaysLogOutput) ToAlwaysLogPtrOutput() AlwaysLogPtrOutput {
	return o.ToAlwaysLogPtrOutputWithContext(context.Background())
}

func (o AlwaysLogOutput) ToAlwaysLogPtrOutputWithContext(ctx context.Context) AlwaysLogPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlwaysLog) *AlwaysLog {
		return &v
	}).(AlwaysLogPtrOutput)
}

func (o AlwaysLogOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AlwaysLogOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AlwaysLog) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AlwaysLogOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlwaysLogOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AlwaysLog) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AlwaysLogPtrOutput struct{ *pulumi.OutputState }

func (AlwaysLogPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlwaysLog)(nil)).Elem()
}

func (o AlwaysLogPtrOutput) ToAlwaysLogPtrOutput() AlwaysLogPtrOutput {
	return o
}

func (o AlwaysLogPtrOutput) ToAlwaysLogPtrOutputWithContext(ctx context.Context) AlwaysLogPtrOutput {
	return o
}

func (o AlwaysLogPtrOutput) Elem() AlwaysLogOutput {
	return o.ApplyT(func(v *AlwaysLog) AlwaysLog {
		if v != nil {
			return *v
		}
		var ret AlwaysLog
		return ret
	}).(AlwaysLogOutput)
}

func (o AlwaysLogPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlwaysLogPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AlwaysLog) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AlwaysLogInput interface {
	pulumi.Input

	ToAlwaysLogOutput() AlwaysLogOutput
	ToAlwaysLogOutputWithContext(context.Context) AlwaysLogOutput
}

var alwaysLogPtrType = reflect.TypeOf((**AlwaysLog)(nil)).Elem()

type AlwaysLogPtrInput interface {
	pulumi.Input

	ToAlwaysLogPtrOutput() AlwaysLogPtrOutput
	ToAlwaysLogPtrOutputWithContext(context.Context) AlwaysLogPtrOutput
}

type alwaysLogPtr string

func AlwaysLogPtr(v string) AlwaysLogPtrInput {
	return (*alwaysLogPtr)(&v)
}

func (*alwaysLogPtr) ElementType() reflect.Type {
	return alwaysLogPtrType
}

func (in *alwaysLogPtr) ToAlwaysLogPtrOutput() AlwaysLogPtrOutput {
	return pulumi.ToOutput(in).(AlwaysLogPtrOutput)
}

func (in *alwaysLogPtr) ToAlwaysLogPtrOutputWithContext(ctx context.Context) AlwaysLogPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AlwaysLogPtrOutput)
}

type ApiType string

const (
	ApiTypeHttp      = ApiType("http")
	ApiTypeSoap      = ApiType("soap")
	ApiTypeWebsocket = ApiType("websocket")
	ApiTypeGraphql   = ApiType("graphql")
)

func (ApiType) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiType)(nil)).Elem()
}

func (e ApiType) ToApiTypeOutput() ApiTypeOutput {
	return pulumi.ToOutput(e).(ApiTypeOutput)
}

func (e ApiType) ToApiTypeOutputWithContext(ctx context.Context) ApiTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApiTypeOutput)
}

func (e ApiType) ToApiTypePtrOutput() ApiTypePtrOutput {
	return e.ToApiTypePtrOutputWithContext(context.Background())
}

func (e ApiType) ToApiTypePtrOutputWithContext(ctx context.Context) ApiTypePtrOutput {
	return ApiType(e).ToApiTypeOutputWithContext(ctx).ToApiTypePtrOutputWithContext(ctx)
}

func (e ApiType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApiType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApiType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApiType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApiTypeOutput struct{ *pulumi.OutputState }

func (ApiTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiType)(nil)).Elem()
}

func (o ApiTypeOutput) ToApiTypeOutput() ApiTypeOutput {
	return o
}

func (o ApiTypeOutput) ToApiTypeOutputWithContext(ctx context.Context) ApiTypeOutput {
	return o
}

func (o ApiTypeOutput) ToApiTypePtrOutput() ApiTypePtrOutput {
	return o.ToApiTypePtrOutputWithContext(context.Background())
}

func (o ApiTypeOutput) ToApiTypePtrOutputWithContext(ctx context.Context) ApiTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiType) *ApiType {
		return &v
	}).(ApiTypePtrOutput)
}

func (o ApiTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApiTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApiType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApiTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApiTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApiType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApiTypePtrOutput struct{ *pulumi.OutputState }

func (ApiTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiType)(nil)).Elem()
}

func (o ApiTypePtrOutput) ToApiTypePtrOutput() ApiTypePtrOutput {
	return o
}

func (o ApiTypePtrOutput) ToApiTypePtrOutputWithContext(ctx context.Context) ApiTypePtrOutput {
	return o
}

func (o ApiTypePtrOutput) Elem() ApiTypeOutput {
	return o.ApplyT(func(v *ApiType) ApiType {
		if v != nil {
			return *v
		}
		var ret ApiType
		return ret
	}).(ApiTypeOutput)
}

func (o ApiTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApiTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApiType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApiTypeInput interface {
	pulumi.Input

	ToApiTypeOutput() ApiTypeOutput
	ToApiTypeOutputWithContext(context.Context) ApiTypeOutput
}

var apiTypePtrType = reflect.TypeOf((**ApiType)(nil)).Elem()

type ApiTypePtrInput interface {
	pulumi.Input

	ToApiTypePtrOutput() ApiTypePtrOutput
	ToApiTypePtrOutputWithContext(context.Context) ApiTypePtrOutput
}

type apiTypePtr string

func ApiTypePtr(v string) ApiTypePtrInput {
	return (*apiTypePtr)(&v)
}

func (*apiTypePtr) ElementType() reflect.Type {
	return apiTypePtrType
}

func (in *apiTypePtr) ToApiTypePtrOutput() ApiTypePtrOutput {
	return pulumi.ToOutput(in).(ApiTypePtrOutput)
}

func (in *apiTypePtr) ToApiTypePtrOutputWithContext(ctx context.Context) ApiTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApiTypePtrOutput)
}

type ApimIdentityType string

const (
	ApimIdentityTypeSystemAssigned               = ApimIdentityType("SystemAssigned")
	ApimIdentityTypeUserAssigned                 = ApimIdentityType("UserAssigned")
	ApimIdentityType_SystemAssigned_UserAssigned = ApimIdentityType("SystemAssigned, UserAssigned")
	ApimIdentityTypeNone                         = ApimIdentityType("None")
)

func (ApimIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ApimIdentityType)(nil)).Elem()
}

func (e ApimIdentityType) ToApimIdentityTypeOutput() ApimIdentityTypeOutput {
	return pulumi.ToOutput(e).(ApimIdentityTypeOutput)
}

func (e ApimIdentityType) ToApimIdentityTypeOutputWithContext(ctx context.Context) ApimIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApimIdentityTypeOutput)
}

func (e ApimIdentityType) ToApimIdentityTypePtrOutput() ApimIdentityTypePtrOutput {
	return e.ToApimIdentityTypePtrOutputWithContext(context.Background())
}

func (e ApimIdentityType) ToApimIdentityTypePtrOutputWithContext(ctx context.Context) ApimIdentityTypePtrOutput {
	return ApimIdentityType(e).ToApimIdentityTypeOutputWithContext(ctx).ToApimIdentityTypePtrOutputWithContext(ctx)
}

func (e ApimIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApimIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApimIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApimIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApimIdentityTypeOutput struct{ *pulumi.OutputState }

func (ApimIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApimIdentityType)(nil)).Elem()
}

func (o ApimIdentityTypeOutput) ToApimIdentityTypeOutput() ApimIdentityTypeOutput {
	return o
}

func (o ApimIdentityTypeOutput) ToApimIdentityTypeOutputWithContext(ctx context.Context) ApimIdentityTypeOutput {
	return o
}

func (o ApimIdentityTypeOutput) ToApimIdentityTypePtrOutput() ApimIdentityTypePtrOutput {
	return o.ToApimIdentityTypePtrOutputWithContext(context.Background())
}

func (o ApimIdentityTypeOutput) ToApimIdentityTypePtrOutputWithContext(ctx context.Context) ApimIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApimIdentityType) *ApimIdentityType {
		return &v
	}).(ApimIdentityTypePtrOutput)
}

func (o ApimIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApimIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApimIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApimIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApimIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApimIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApimIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ApimIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApimIdentityType)(nil)).Elem()
}

func (o ApimIdentityTypePtrOutput) ToApimIdentityTypePtrOutput() ApimIdentityTypePtrOutput {
	return o
}

func (o ApimIdentityTypePtrOutput) ToApimIdentityTypePtrOutputWithContext(ctx context.Context) ApimIdentityTypePtrOutput {
	return o
}

func (o ApimIdentityTypePtrOutput) Elem() ApimIdentityTypeOutput {
	return o.ApplyT(func(v *ApimIdentityType) ApimIdentityType {
		if v != nil {
			return *v
		}
		var ret ApimIdentityType
		return ret
	}).(ApimIdentityTypeOutput)
}

func (o ApimIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApimIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApimIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApimIdentityTypeInput interface {
	pulumi.Input

	ToApimIdentityTypeOutput() ApimIdentityTypeOutput
	ToApimIdentityTypeOutputWithContext(context.Context) ApimIdentityTypeOutput
}

var apimIdentityTypePtrType = reflect.TypeOf((**ApimIdentityType)(nil)).Elem()

type ApimIdentityTypePtrInput interface {
	pulumi.Input

	ToApimIdentityTypePtrOutput() ApimIdentityTypePtrOutput
	ToApimIdentityTypePtrOutputWithContext(context.Context) ApimIdentityTypePtrOutput
}

type apimIdentityTypePtr string

func ApimIdentityTypePtr(v string) ApimIdentityTypePtrInput {
	return (*apimIdentityTypePtr)(&v)
}

func (*apimIdentityTypePtr) ElementType() reflect.Type {
	return apimIdentityTypePtrType
}

func (in *apimIdentityTypePtr) ToApimIdentityTypePtrOutput() ApimIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ApimIdentityTypePtrOutput)
}

func (in *apimIdentityTypePtr) ToApimIdentityTypePtrOutputWithContext(ctx context.Context) ApimIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApimIdentityTypePtrOutput)
}

type AppType string

const (
	// User create request was sent by legacy developer portal.
	AppTypePortal = AppType("portal")
	// User create request was sent by new developer portal.
	AppTypeDeveloperPortal = AppType("developerPortal")
)

func (AppType) ElementType() reflect.Type {
	return reflect.TypeOf((*AppType)(nil)).Elem()
}

func (e AppType) ToAppTypeOutput() AppTypeOutput {
	return pulumi.ToOutput(e).(AppTypeOutput)
}

func (e AppType) ToAppTypeOutputWithContext(ctx context.Context) AppTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AppTypeOutput)
}

func (e AppType) ToAppTypePtrOutput() AppTypePtrOutput {
	return e.ToAppTypePtrOutputWithContext(context.Background())
}

func (e AppType) ToAppTypePtrOutputWithContext(ctx context.Context) AppTypePtrOutput {
	return AppType(e).ToAppTypeOutputWithContext(ctx).ToAppTypePtrOutputWithContext(ctx)
}

func (e AppType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AppType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AppTypeOutput struct{ *pulumi.OutputState }

func (AppTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppType)(nil)).Elem()
}

func (o AppTypeOutput) ToAppTypeOutput() AppTypeOutput {
	return o
}

func (o AppTypeOutput) ToAppTypeOutputWithContext(ctx context.Context) AppTypeOutput {
	return o
}

func (o AppTypeOutput) ToAppTypePtrOutput() AppTypePtrOutput {
	return o.ToAppTypePtrOutputWithContext(context.Background())
}

func (o AppTypeOutput) ToAppTypePtrOutputWithContext(ctx context.Context) AppTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppType) *AppType {
		return &v
	}).(AppTypePtrOutput)
}

func (o AppTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AppTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AppTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AppTypePtrOutput struct{ *pulumi.OutputState }

func (AppTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppType)(nil)).Elem()
}

func (o AppTypePtrOutput) ToAppTypePtrOutput() AppTypePtrOutput {
	return o
}

func (o AppTypePtrOutput) ToAppTypePtrOutputWithContext(ctx context.Context) AppTypePtrOutput {
	return o
}

func (o AppTypePtrOutput) Elem() AppTypeOutput {
	return o.ApplyT(func(v *AppType) AppType {
		if v != nil {
			return *v
		}
		var ret AppType
		return ret
	}).(AppTypeOutput)
}

func (o AppTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AppType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AppTypeInput interface {
	pulumi.Input

	ToAppTypeOutput() AppTypeOutput
	ToAppTypeOutputWithContext(context.Context) AppTypeOutput
}

var appTypePtrType = reflect.TypeOf((**AppType)(nil)).Elem()

type AppTypePtrInput interface {
	pulumi.Input

	ToAppTypePtrOutput() AppTypePtrOutput
	ToAppTypePtrOutputWithContext(context.Context) AppTypePtrOutput
}

type appTypePtr string

func AppTypePtr(v string) AppTypePtrInput {
	return (*appTypePtr)(&v)
}

func (*appTypePtr) ElementType() reflect.Type {
	return appTypePtrType
}

func (in *appTypePtr) ToAppTypePtrOutput() AppTypePtrOutput {
	return pulumi.ToOutput(in).(AppTypePtrOutput)
}

func (in *appTypePtr) ToAppTypePtrOutputWithContext(ctx context.Context) AppTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AppTypePtrOutput)
}

type AuthorizationMethod string

const (
	AuthorizationMethodHEAD    = AuthorizationMethod("HEAD")
	AuthorizationMethodOPTIONS = AuthorizationMethod("OPTIONS")
	AuthorizationMethodTRACE   = AuthorizationMethod("TRACE")
	AuthorizationMethodGET     = AuthorizationMethod("GET")
	AuthorizationMethodPOST    = AuthorizationMethod("POST")
	AuthorizationMethodPUT     = AuthorizationMethod("PUT")
	AuthorizationMethodPATCH   = AuthorizationMethod("PATCH")
	AuthorizationMethodDELETE  = AuthorizationMethod("DELETE")
)

func (AuthorizationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationMethod)(nil)).Elem()
}

func (e AuthorizationMethod) ToAuthorizationMethodOutput() AuthorizationMethodOutput {
	return pulumi.ToOutput(e).(AuthorizationMethodOutput)
}

func (e AuthorizationMethod) ToAuthorizationMethodOutputWithContext(ctx context.Context) AuthorizationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthorizationMethodOutput)
}

func (e AuthorizationMethod) ToAuthorizationMethodPtrOutput() AuthorizationMethodPtrOutput {
	return e.ToAuthorizationMethodPtrOutputWithContext(context.Background())
}

func (e AuthorizationMethod) ToAuthorizationMethodPtrOutputWithContext(ctx context.Context) AuthorizationMethodPtrOutput {
	return AuthorizationMethod(e).ToAuthorizationMethodOutputWithContext(ctx).ToAuthorizationMethodPtrOutputWithContext(ctx)
}

func (e AuthorizationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthorizationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthorizationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthorizationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthorizationMethodOutput struct{ *pulumi.OutputState }

func (AuthorizationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationMethod)(nil)).Elem()
}

func (o AuthorizationMethodOutput) ToAuthorizationMethodOutput() AuthorizationMethodOutput {
	return o
}

func (o AuthorizationMethodOutput) ToAuthorizationMethodOutputWithContext(ctx context.Context) AuthorizationMethodOutput {
	return o
}

func (o AuthorizationMethodOutput) ToAuthorizationMethodPtrOutput() AuthorizationMethodPtrOutput {
	return o.ToAuthorizationMethodPtrOutputWithContext(context.Background())
}

func (o AuthorizationMethodOutput) ToAuthorizationMethodPtrOutputWithContext(ctx context.Context) AuthorizationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthorizationMethod) *AuthorizationMethod {
		return &v
	}).(AuthorizationMethodPtrOutput)
}

func (o AuthorizationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthorizationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthorizationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthorizationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthorizationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthorizationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthorizationMethodPtrOutput struct{ *pulumi.OutputState }

func (AuthorizationMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationMethod)(nil)).Elem()
}

func (o AuthorizationMethodPtrOutput) ToAuthorizationMethodPtrOutput() AuthorizationMethodPtrOutput {
	return o
}

func (o AuthorizationMethodPtrOutput) ToAuthorizationMethodPtrOutputWithContext(ctx context.Context) AuthorizationMethodPtrOutput {
	return o
}

func (o AuthorizationMethodPtrOutput) Elem() AuthorizationMethodOutput {
	return o.ApplyT(func(v *AuthorizationMethod) AuthorizationMethod {
		if v != nil {
			return *v
		}
		var ret AuthorizationMethod
		return ret
	}).(AuthorizationMethodOutput)
}

func (o AuthorizationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthorizationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthorizationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AuthorizationMethodInput interface {
	pulumi.Input

	ToAuthorizationMethodOutput() AuthorizationMethodOutput
	ToAuthorizationMethodOutputWithContext(context.Context) AuthorizationMethodOutput
}

var authorizationMethodPtrType = reflect.TypeOf((**AuthorizationMethod)(nil)).Elem()

type AuthorizationMethodPtrInput interface {
	pulumi.Input

	ToAuthorizationMethodPtrOutput() AuthorizationMethodPtrOutput
	ToAuthorizationMethodPtrOutputWithContext(context.Context) AuthorizationMethodPtrOutput
}

type authorizationMethodPtr string

func AuthorizationMethodPtr(v string) AuthorizationMethodPtrInput {
	return (*authorizationMethodPtr)(&v)
}

func (*authorizationMethodPtr) ElementType() reflect.Type {
	return authorizationMethodPtrType
}

func (in *authorizationMethodPtr) ToAuthorizationMethodPtrOutput() AuthorizationMethodPtrOutput {
	return pulumi.ToOutput(in).(AuthorizationMethodPtrOutput)
}

func (in *authorizationMethodPtr) ToAuthorizationMethodPtrOutputWithContext(ctx context.Context) AuthorizationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthorizationMethodPtrOutput)
}





type AuthorizationMethodArrayInput interface {
	pulumi.Input

	ToAuthorizationMethodArrayOutput() AuthorizationMethodArrayOutput
	ToAuthorizationMethodArrayOutputWithContext(context.Context) AuthorizationMethodArrayOutput
}

type AuthorizationMethodArray []AuthorizationMethod

func (AuthorizationMethodArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthorizationMethod)(nil)).Elem()
}

func (i AuthorizationMethodArray) ToAuthorizationMethodArrayOutput() AuthorizationMethodArrayOutput {
	return i.ToAuthorizationMethodArrayOutputWithContext(context.Background())
}

func (i AuthorizationMethodArray) ToAuthorizationMethodArrayOutputWithContext(ctx context.Context) AuthorizationMethodArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationMethodArrayOutput)
}

type AuthorizationMethodArrayOutput struct{ *pulumi.OutputState }

func (AuthorizationMethodArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthorizationMethod)(nil)).Elem()
}

func (o AuthorizationMethodArrayOutput) ToAuthorizationMethodArrayOutput() AuthorizationMethodArrayOutput {
	return o
}

func (o AuthorizationMethodArrayOutput) ToAuthorizationMethodArrayOutputWithContext(ctx context.Context) AuthorizationMethodArrayOutput {
	return o
}

func (o AuthorizationMethodArrayOutput) Index(i pulumi.IntInput) AuthorizationMethodOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthorizationMethod {
		return vs[0].([]AuthorizationMethod)[vs[1].(int)]
	}).(AuthorizationMethodOutput)
}

type BackendProtocol string

const (
	// The Backend is a RESTful service.
	BackendProtocolHttp = BackendProtocol("http")
	// The Backend is a SOAP service.
	BackendProtocolSoap = BackendProtocol("soap")
)

func (BackendProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendProtocol)(nil)).Elem()
}

func (e BackendProtocol) ToBackendProtocolOutput() BackendProtocolOutput {
	return pulumi.ToOutput(e).(BackendProtocolOutput)
}

func (e BackendProtocol) ToBackendProtocolOutputWithContext(ctx context.Context) BackendProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BackendProtocolOutput)
}

func (e BackendProtocol) ToBackendProtocolPtrOutput() BackendProtocolPtrOutput {
	return e.ToBackendProtocolPtrOutputWithContext(context.Background())
}

func (e BackendProtocol) ToBackendProtocolPtrOutputWithContext(ctx context.Context) BackendProtocolPtrOutput {
	return BackendProtocol(e).ToBackendProtocolOutputWithContext(ctx).ToBackendProtocolPtrOutputWithContext(ctx)
}

func (e BackendProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackendProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackendProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BackendProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BackendProtocolOutput struct{ *pulumi.OutputState }

func (BackendProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendProtocol)(nil)).Elem()
}

func (o BackendProtocolOutput) ToBackendProtocolOutput() BackendProtocolOutput {
	return o
}

func (o BackendProtocolOutput) ToBackendProtocolOutputWithContext(ctx context.Context) BackendProtocolOutput {
	return o
}

func (o BackendProtocolOutput) ToBackendProtocolPtrOutput() BackendProtocolPtrOutput {
	return o.ToBackendProtocolPtrOutputWithContext(context.Background())
}

func (o BackendProtocolOutput) ToBackendProtocolPtrOutputWithContext(ctx context.Context) BackendProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendProtocol) *BackendProtocol {
		return &v
	}).(BackendProtocolPtrOutput)
}

func (o BackendProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BackendProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackendProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BackendProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackendProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackendProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BackendProtocolPtrOutput struct{ *pulumi.OutputState }

func (BackendProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendProtocol)(nil)).Elem()
}

func (o BackendProtocolPtrOutput) ToBackendProtocolPtrOutput() BackendProtocolPtrOutput {
	return o
}

func (o BackendProtocolPtrOutput) ToBackendProtocolPtrOutputWithContext(ctx context.Context) BackendProtocolPtrOutput {
	return o
}

func (o BackendProtocolPtrOutput) Elem() BackendProtocolOutput {
	return o.ApplyT(func(v *BackendProtocol) BackendProtocol {
		if v != nil {
			return *v
		}
		var ret BackendProtocol
		return ret
	}).(BackendProtocolOutput)
}

func (o BackendProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackendProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BackendProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BackendProtocolInput interface {
	pulumi.Input

	ToBackendProtocolOutput() BackendProtocolOutput
	ToBackendProtocolOutputWithContext(context.Context) BackendProtocolOutput
}

var backendProtocolPtrType = reflect.TypeOf((**BackendProtocol)(nil)).Elem()

type BackendProtocolPtrInput interface {
	pulumi.Input

	ToBackendProtocolPtrOutput() BackendProtocolPtrOutput
	ToBackendProtocolPtrOutputWithContext(context.Context) BackendProtocolPtrOutput
}

type backendProtocolPtr string

func BackendProtocolPtr(v string) BackendProtocolPtrInput {
	return (*backendProtocolPtr)(&v)
}

func (*backendProtocolPtr) ElementType() reflect.Type {
	return backendProtocolPtrType
}

func (in *backendProtocolPtr) ToBackendProtocolPtrOutput() BackendProtocolPtrOutput {
	return pulumi.ToOutput(in).(BackendProtocolPtrOutput)
}

func (in *backendProtocolPtr) ToBackendProtocolPtrOutputWithContext(ctx context.Context) BackendProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BackendProtocolPtrOutput)
}

type BearerTokenSendingMethod string

const (
	BearerTokenSendingMethodAuthorizationHeader = BearerTokenSendingMethod("authorizationHeader")
	BearerTokenSendingMethodQuery               = BearerTokenSendingMethod("query")
)

func (BearerTokenSendingMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*BearerTokenSendingMethod)(nil)).Elem()
}

func (e BearerTokenSendingMethod) ToBearerTokenSendingMethodOutput() BearerTokenSendingMethodOutput {
	return pulumi.ToOutput(e).(BearerTokenSendingMethodOutput)
}

func (e BearerTokenSendingMethod) ToBearerTokenSendingMethodOutputWithContext(ctx context.Context) BearerTokenSendingMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BearerTokenSendingMethodOutput)
}

func (e BearerTokenSendingMethod) ToBearerTokenSendingMethodPtrOutput() BearerTokenSendingMethodPtrOutput {
	return e.ToBearerTokenSendingMethodPtrOutputWithContext(context.Background())
}

func (e BearerTokenSendingMethod) ToBearerTokenSendingMethodPtrOutputWithContext(ctx context.Context) BearerTokenSendingMethodPtrOutput {
	return BearerTokenSendingMethod(e).ToBearerTokenSendingMethodOutputWithContext(ctx).ToBearerTokenSendingMethodPtrOutputWithContext(ctx)
}

func (e BearerTokenSendingMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BearerTokenSendingMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BearerTokenSendingMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BearerTokenSendingMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BearerTokenSendingMethodOutput struct{ *pulumi.OutputState }

func (BearerTokenSendingMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BearerTokenSendingMethod)(nil)).Elem()
}

func (o BearerTokenSendingMethodOutput) ToBearerTokenSendingMethodOutput() BearerTokenSendingMethodOutput {
	return o
}

func (o BearerTokenSendingMethodOutput) ToBearerTokenSendingMethodOutputWithContext(ctx context.Context) BearerTokenSendingMethodOutput {
	return o
}

func (o BearerTokenSendingMethodOutput) ToBearerTokenSendingMethodPtrOutput() BearerTokenSendingMethodPtrOutput {
	return o.ToBearerTokenSendingMethodPtrOutputWithContext(context.Background())
}

func (o BearerTokenSendingMethodOutput) ToBearerTokenSendingMethodPtrOutputWithContext(ctx context.Context) BearerTokenSendingMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BearerTokenSendingMethod) *BearerTokenSendingMethod {
		return &v
	}).(BearerTokenSendingMethodPtrOutput)
}

func (o BearerTokenSendingMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BearerTokenSendingMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BearerTokenSendingMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BearerTokenSendingMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BearerTokenSendingMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BearerTokenSendingMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BearerTokenSendingMethodPtrOutput struct{ *pulumi.OutputState }

func (BearerTokenSendingMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BearerTokenSendingMethod)(nil)).Elem()
}

func (o BearerTokenSendingMethodPtrOutput) ToBearerTokenSendingMethodPtrOutput() BearerTokenSendingMethodPtrOutput {
	return o
}

func (o BearerTokenSendingMethodPtrOutput) ToBearerTokenSendingMethodPtrOutputWithContext(ctx context.Context) BearerTokenSendingMethodPtrOutput {
	return o
}

func (o BearerTokenSendingMethodPtrOutput) Elem() BearerTokenSendingMethodOutput {
	return o.ApplyT(func(v *BearerTokenSendingMethod) BearerTokenSendingMethod {
		if v != nil {
			return *v
		}
		var ret BearerTokenSendingMethod
		return ret
	}).(BearerTokenSendingMethodOutput)
}

func (o BearerTokenSendingMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BearerTokenSendingMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BearerTokenSendingMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BearerTokenSendingMethodInput interface {
	pulumi.Input

	ToBearerTokenSendingMethodOutput() BearerTokenSendingMethodOutput
	ToBearerTokenSendingMethodOutputWithContext(context.Context) BearerTokenSendingMethodOutput
}

var bearerTokenSendingMethodPtrType = reflect.TypeOf((**BearerTokenSendingMethod)(nil)).Elem()

type BearerTokenSendingMethodPtrInput interface {
	pulumi.Input

	ToBearerTokenSendingMethodPtrOutput() BearerTokenSendingMethodPtrOutput
	ToBearerTokenSendingMethodPtrOutputWithContext(context.Context) BearerTokenSendingMethodPtrOutput
}

type bearerTokenSendingMethodPtr string

func BearerTokenSendingMethodPtr(v string) BearerTokenSendingMethodPtrInput {
	return (*bearerTokenSendingMethodPtr)(&v)
}

func (*bearerTokenSendingMethodPtr) ElementType() reflect.Type {
	return bearerTokenSendingMethodPtrType
}

func (in *bearerTokenSendingMethodPtr) ToBearerTokenSendingMethodPtrOutput() BearerTokenSendingMethodPtrOutput {
	return pulumi.ToOutput(in).(BearerTokenSendingMethodPtrOutput)
}

func (in *bearerTokenSendingMethodPtr) ToBearerTokenSendingMethodPtrOutputWithContext(ctx context.Context) BearerTokenSendingMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BearerTokenSendingMethodPtrOutput)
}

type BearerTokenSendingMethods string

const (
	// Access token will be transmitted in the Authorization header using Bearer schema
	BearerTokenSendingMethodsAuthorizationHeader = BearerTokenSendingMethods("authorizationHeader")
	// Access token will be transmitted as query parameters.
	BearerTokenSendingMethodsQuery = BearerTokenSendingMethods("query")
)

func (BearerTokenSendingMethods) ElementType() reflect.Type {
	return reflect.TypeOf((*BearerTokenSendingMethods)(nil)).Elem()
}

func (e BearerTokenSendingMethods) ToBearerTokenSendingMethodsOutput() BearerTokenSendingMethodsOutput {
	return pulumi.ToOutput(e).(BearerTokenSendingMethodsOutput)
}

func (e BearerTokenSendingMethods) ToBearerTokenSendingMethodsOutputWithContext(ctx context.Context) BearerTokenSendingMethodsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BearerTokenSendingMethodsOutput)
}

func (e BearerTokenSendingMethods) ToBearerTokenSendingMethodsPtrOutput() BearerTokenSendingMethodsPtrOutput {
	return e.ToBearerTokenSendingMethodsPtrOutputWithContext(context.Background())
}

func (e BearerTokenSendingMethods) ToBearerTokenSendingMethodsPtrOutputWithContext(ctx context.Context) BearerTokenSendingMethodsPtrOutput {
	return BearerTokenSendingMethods(e).ToBearerTokenSendingMethodsOutputWithContext(ctx).ToBearerTokenSendingMethodsPtrOutputWithContext(ctx)
}

func (e BearerTokenSendingMethods) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BearerTokenSendingMethods) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BearerTokenSendingMethods) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BearerTokenSendingMethods) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BearerTokenSendingMethodsOutput struct{ *pulumi.OutputState }

func (BearerTokenSendingMethodsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BearerTokenSendingMethods)(nil)).Elem()
}

func (o BearerTokenSendingMethodsOutput) ToBearerTokenSendingMethodsOutput() BearerTokenSendingMethodsOutput {
	return o
}

func (o BearerTokenSendingMethodsOutput) ToBearerTokenSendingMethodsOutputWithContext(ctx context.Context) BearerTokenSendingMethodsOutput {
	return o
}

func (o BearerTokenSendingMethodsOutput) ToBearerTokenSendingMethodsPtrOutput() BearerTokenSendingMethodsPtrOutput {
	return o.ToBearerTokenSendingMethodsPtrOutputWithContext(context.Background())
}

func (o BearerTokenSendingMethodsOutput) ToBearerTokenSendingMethodsPtrOutputWithContext(ctx context.Context) BearerTokenSendingMethodsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BearerTokenSendingMethods) *BearerTokenSendingMethods {
		return &v
	}).(BearerTokenSendingMethodsPtrOutput)
}

func (o BearerTokenSendingMethodsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BearerTokenSendingMethodsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BearerTokenSendingMethods) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BearerTokenSendingMethodsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BearerTokenSendingMethodsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BearerTokenSendingMethods) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BearerTokenSendingMethodsPtrOutput struct{ *pulumi.OutputState }

func (BearerTokenSendingMethodsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BearerTokenSendingMethods)(nil)).Elem()
}

func (o BearerTokenSendingMethodsPtrOutput) ToBearerTokenSendingMethodsPtrOutput() BearerTokenSendingMethodsPtrOutput {
	return o
}

func (o BearerTokenSendingMethodsPtrOutput) ToBearerTokenSendingMethodsPtrOutputWithContext(ctx context.Context) BearerTokenSendingMethodsPtrOutput {
	return o
}

func (o BearerTokenSendingMethodsPtrOutput) Elem() BearerTokenSendingMethodsOutput {
	return o.ApplyT(func(v *BearerTokenSendingMethods) BearerTokenSendingMethods {
		if v != nil {
			return *v
		}
		var ret BearerTokenSendingMethods
		return ret
	}).(BearerTokenSendingMethodsOutput)
}

func (o BearerTokenSendingMethodsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BearerTokenSendingMethodsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BearerTokenSendingMethods) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BearerTokenSendingMethodsInput interface {
	pulumi.Input

	ToBearerTokenSendingMethodsOutput() BearerTokenSendingMethodsOutput
	ToBearerTokenSendingMethodsOutputWithContext(context.Context) BearerTokenSendingMethodsOutput
}

var bearerTokenSendingMethodsPtrType = reflect.TypeOf((**BearerTokenSendingMethods)(nil)).Elem()

type BearerTokenSendingMethodsPtrInput interface {
	pulumi.Input

	ToBearerTokenSendingMethodsPtrOutput() BearerTokenSendingMethodsPtrOutput
	ToBearerTokenSendingMethodsPtrOutputWithContext(context.Context) BearerTokenSendingMethodsPtrOutput
}

type bearerTokenSendingMethodsPtr string

func BearerTokenSendingMethodsPtr(v string) BearerTokenSendingMethodsPtrInput {
	return (*bearerTokenSendingMethodsPtr)(&v)
}

func (*bearerTokenSendingMethodsPtr) ElementType() reflect.Type {
	return bearerTokenSendingMethodsPtrType
}

func (in *bearerTokenSendingMethodsPtr) ToBearerTokenSendingMethodsPtrOutput() BearerTokenSendingMethodsPtrOutput {
	return pulumi.ToOutput(in).(BearerTokenSendingMethodsPtrOutput)
}

func (in *bearerTokenSendingMethodsPtr) ToBearerTokenSendingMethodsPtrOutputWithContext(ctx context.Context) BearerTokenSendingMethodsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BearerTokenSendingMethodsPtrOutput)
}

type CertificateSource string

const (
	CertificateSourceManaged  = CertificateSource("Managed")
	CertificateSourceKeyVault = CertificateSource("KeyVault")
	CertificateSourceCustom   = CertificateSource("Custom")
	CertificateSourceBuiltIn  = CertificateSource("BuiltIn")
)

func (CertificateSource) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateSource)(nil)).Elem()
}

func (e CertificateSource) ToCertificateSourceOutput() CertificateSourceOutput {
	return pulumi.ToOutput(e).(CertificateSourceOutput)
}

func (e CertificateSource) ToCertificateSourceOutputWithContext(ctx context.Context) CertificateSourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CertificateSourceOutput)
}

func (e CertificateSource) ToCertificateSourcePtrOutput() CertificateSourcePtrOutput {
	return e.ToCertificateSourcePtrOutputWithContext(context.Background())
}

func (e CertificateSource) ToCertificateSourcePtrOutputWithContext(ctx context.Context) CertificateSourcePtrOutput {
	return CertificateSource(e).ToCertificateSourceOutputWithContext(ctx).ToCertificateSourcePtrOutputWithContext(ctx)
}

func (e CertificateSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CertificateSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CertificateSourceOutput struct{ *pulumi.OutputState }

func (CertificateSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateSource)(nil)).Elem()
}

func (o CertificateSourceOutput) ToCertificateSourceOutput() CertificateSourceOutput {
	return o
}

func (o CertificateSourceOutput) ToCertificateSourceOutputWithContext(ctx context.Context) CertificateSourceOutput {
	return o
}

func (o CertificateSourceOutput) ToCertificateSourcePtrOutput() CertificateSourcePtrOutput {
	return o.ToCertificateSourcePtrOutputWithContext(context.Background())
}

func (o CertificateSourceOutput) ToCertificateSourcePtrOutputWithContext(ctx context.Context) CertificateSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateSource) *CertificateSource {
		return &v
	}).(CertificateSourcePtrOutput)
}

func (o CertificateSourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CertificateSourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateSource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CertificateSourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateSourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateSource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CertificateSourcePtrOutput struct{ *pulumi.OutputState }

func (CertificateSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateSource)(nil)).Elem()
}

func (o CertificateSourcePtrOutput) ToCertificateSourcePtrOutput() CertificateSourcePtrOutput {
	return o
}

func (o CertificateSourcePtrOutput) ToCertificateSourcePtrOutputWithContext(ctx context.Context) CertificateSourcePtrOutput {
	return o
}

func (o CertificateSourcePtrOutput) Elem() CertificateSourceOutput {
	return o.ApplyT(func(v *CertificateSource) CertificateSource {
		if v != nil {
			return *v
		}
		var ret CertificateSource
		return ret
	}).(CertificateSourceOutput)
}

func (o CertificateSourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateSourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CertificateSource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CertificateSourceInput interface {
	pulumi.Input

	ToCertificateSourceOutput() CertificateSourceOutput
	ToCertificateSourceOutputWithContext(context.Context) CertificateSourceOutput
}

var certificateSourcePtrType = reflect.TypeOf((**CertificateSource)(nil)).Elem()

type CertificateSourcePtrInput interface {
	pulumi.Input

	ToCertificateSourcePtrOutput() CertificateSourcePtrOutput
	ToCertificateSourcePtrOutputWithContext(context.Context) CertificateSourcePtrOutput
}

type certificateSourcePtr string

func CertificateSourcePtr(v string) CertificateSourcePtrInput {
	return (*certificateSourcePtr)(&v)
}

func (*certificateSourcePtr) ElementType() reflect.Type {
	return certificateSourcePtrType
}

func (in *certificateSourcePtr) ToCertificateSourcePtrOutput() CertificateSourcePtrOutput {
	return pulumi.ToOutput(in).(CertificateSourcePtrOutput)
}

func (in *certificateSourcePtr) ToCertificateSourcePtrOutputWithContext(ctx context.Context) CertificateSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CertificateSourcePtrOutput)
}

type CertificateStatus string

const (
	CertificateStatusCompleted  = CertificateStatus("Completed")
	CertificateStatusFailed     = CertificateStatus("Failed")
	CertificateStatusInProgress = CertificateStatus("InProgress")
)

func (CertificateStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateStatus)(nil)).Elem()
}

func (e CertificateStatus) ToCertificateStatusOutput() CertificateStatusOutput {
	return pulumi.ToOutput(e).(CertificateStatusOutput)
}

func (e CertificateStatus) ToCertificateStatusOutputWithContext(ctx context.Context) CertificateStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CertificateStatusOutput)
}

func (e CertificateStatus) ToCertificateStatusPtrOutput() CertificateStatusPtrOutput {
	return e.ToCertificateStatusPtrOutputWithContext(context.Background())
}

func (e CertificateStatus) ToCertificateStatusPtrOutputWithContext(ctx context.Context) CertificateStatusPtrOutput {
	return CertificateStatus(e).ToCertificateStatusOutputWithContext(ctx).ToCertificateStatusPtrOutputWithContext(ctx)
}

func (e CertificateStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CertificateStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CertificateStatusOutput struct{ *pulumi.OutputState }

func (CertificateStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateStatus)(nil)).Elem()
}

func (o CertificateStatusOutput) ToCertificateStatusOutput() CertificateStatusOutput {
	return o
}

func (o CertificateStatusOutput) ToCertificateStatusOutputWithContext(ctx context.Context) CertificateStatusOutput {
	return o
}

func (o CertificateStatusOutput) ToCertificateStatusPtrOutput() CertificateStatusPtrOutput {
	return o.ToCertificateStatusPtrOutputWithContext(context.Background())
}

func (o CertificateStatusOutput) ToCertificateStatusPtrOutputWithContext(ctx context.Context) CertificateStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateStatus) *CertificateStatus {
		return &v
	}).(CertificateStatusPtrOutput)
}

func (o CertificateStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CertificateStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CertificateStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CertificateStatusPtrOutput struct{ *pulumi.OutputState }

func (CertificateStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateStatus)(nil)).Elem()
}

func (o CertificateStatusPtrOutput) ToCertificateStatusPtrOutput() CertificateStatusPtrOutput {
	return o
}

func (o CertificateStatusPtrOutput) ToCertificateStatusPtrOutputWithContext(ctx context.Context) CertificateStatusPtrOutput {
	return o
}

func (o CertificateStatusPtrOutput) Elem() CertificateStatusOutput {
	return o.ApplyT(func(v *CertificateStatus) CertificateStatus {
		if v != nil {
			return *v
		}
		var ret CertificateStatus
		return ret
	}).(CertificateStatusOutput)
}

func (o CertificateStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CertificateStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CertificateStatusInput interface {
	pulumi.Input

	ToCertificateStatusOutput() CertificateStatusOutput
	ToCertificateStatusOutputWithContext(context.Context) CertificateStatusOutput
}

var certificateStatusPtrType = reflect.TypeOf((**CertificateStatus)(nil)).Elem()

type CertificateStatusPtrInput interface {
	pulumi.Input

	ToCertificateStatusPtrOutput() CertificateStatusPtrOutput
	ToCertificateStatusPtrOutputWithContext(context.Context) CertificateStatusPtrOutput
}

type certificateStatusPtr string

func CertificateStatusPtr(v string) CertificateStatusPtrInput {
	return (*certificateStatusPtr)(&v)
}

func (*certificateStatusPtr) ElementType() reflect.Type {
	return certificateStatusPtrType
}

func (in *certificateStatusPtr) ToCertificateStatusPtrOutput() CertificateStatusPtrOutput {
	return pulumi.ToOutput(in).(CertificateStatusPtrOutput)
}

func (in *certificateStatusPtr) ToCertificateStatusPtrOutputWithContext(ctx context.Context) CertificateStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CertificateStatusPtrOutput)
}

type ClientAuthenticationMethod string

const (
	// Basic Client Authentication method.
	ClientAuthenticationMethodBasic = ClientAuthenticationMethod("Basic")
	// Body based Authentication method.
	ClientAuthenticationMethodBody = ClientAuthenticationMethod("Body")
)

func (ClientAuthenticationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthenticationMethod)(nil)).Elem()
}

func (e ClientAuthenticationMethod) ToClientAuthenticationMethodOutput() ClientAuthenticationMethodOutput {
	return pulumi.ToOutput(e).(ClientAuthenticationMethodOutput)
}

func (e ClientAuthenticationMethod) ToClientAuthenticationMethodOutputWithContext(ctx context.Context) ClientAuthenticationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClientAuthenticationMethodOutput)
}

func (e ClientAuthenticationMethod) ToClientAuthenticationMethodPtrOutput() ClientAuthenticationMethodPtrOutput {
	return e.ToClientAuthenticationMethodPtrOutputWithContext(context.Background())
}

func (e ClientAuthenticationMethod) ToClientAuthenticationMethodPtrOutputWithContext(ctx context.Context) ClientAuthenticationMethodPtrOutput {
	return ClientAuthenticationMethod(e).ToClientAuthenticationMethodOutputWithContext(ctx).ToClientAuthenticationMethodPtrOutputWithContext(ctx)
}

func (e ClientAuthenticationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientAuthenticationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientAuthenticationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClientAuthenticationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClientAuthenticationMethodOutput struct{ *pulumi.OutputState }

func (ClientAuthenticationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthenticationMethod)(nil)).Elem()
}

func (o ClientAuthenticationMethodOutput) ToClientAuthenticationMethodOutput() ClientAuthenticationMethodOutput {
	return o
}

func (o ClientAuthenticationMethodOutput) ToClientAuthenticationMethodOutputWithContext(ctx context.Context) ClientAuthenticationMethodOutput {
	return o
}

func (o ClientAuthenticationMethodOutput) ToClientAuthenticationMethodPtrOutput() ClientAuthenticationMethodPtrOutput {
	return o.ToClientAuthenticationMethodPtrOutputWithContext(context.Background())
}

func (o ClientAuthenticationMethodOutput) ToClientAuthenticationMethodPtrOutputWithContext(ctx context.Context) ClientAuthenticationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClientAuthenticationMethod) *ClientAuthenticationMethod {
		return &v
	}).(ClientAuthenticationMethodPtrOutput)
}

func (o ClientAuthenticationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClientAuthenticationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientAuthenticationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClientAuthenticationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientAuthenticationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientAuthenticationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClientAuthenticationMethodPtrOutput struct{ *pulumi.OutputState }

func (ClientAuthenticationMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthenticationMethod)(nil)).Elem()
}

func (o ClientAuthenticationMethodPtrOutput) ToClientAuthenticationMethodPtrOutput() ClientAuthenticationMethodPtrOutput {
	return o
}

func (o ClientAuthenticationMethodPtrOutput) ToClientAuthenticationMethodPtrOutputWithContext(ctx context.Context) ClientAuthenticationMethodPtrOutput {
	return o
}

func (o ClientAuthenticationMethodPtrOutput) Elem() ClientAuthenticationMethodOutput {
	return o.ApplyT(func(v *ClientAuthenticationMethod) ClientAuthenticationMethod {
		if v != nil {
			return *v
		}
		var ret ClientAuthenticationMethod
		return ret
	}).(ClientAuthenticationMethodOutput)
}

func (o ClientAuthenticationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientAuthenticationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClientAuthenticationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClientAuthenticationMethodInput interface {
	pulumi.Input

	ToClientAuthenticationMethodOutput() ClientAuthenticationMethodOutput
	ToClientAuthenticationMethodOutputWithContext(context.Context) ClientAuthenticationMethodOutput
}

var clientAuthenticationMethodPtrType = reflect.TypeOf((**ClientAuthenticationMethod)(nil)).Elem()

type ClientAuthenticationMethodPtrInput interface {
	pulumi.Input

	ToClientAuthenticationMethodPtrOutput() ClientAuthenticationMethodPtrOutput
	ToClientAuthenticationMethodPtrOutputWithContext(context.Context) ClientAuthenticationMethodPtrOutput
}

type clientAuthenticationMethodPtr string

func ClientAuthenticationMethodPtr(v string) ClientAuthenticationMethodPtrInput {
	return (*clientAuthenticationMethodPtr)(&v)
}

func (*clientAuthenticationMethodPtr) ElementType() reflect.Type {
	return clientAuthenticationMethodPtrType
}

func (in *clientAuthenticationMethodPtr) ToClientAuthenticationMethodPtrOutput() ClientAuthenticationMethodPtrOutput {
	return pulumi.ToOutput(in).(ClientAuthenticationMethodPtrOutput)
}

func (in *clientAuthenticationMethodPtr) ToClientAuthenticationMethodPtrOutputWithContext(ctx context.Context) ClientAuthenticationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClientAuthenticationMethodPtrOutput)
}

type Confirmation string

const (
	// Send an e-mail to the user confirming they have successfully signed up.
	ConfirmationSignup = Confirmation("signup")
	// Send an e-mail inviting the user to sign-up and complete registration.
	ConfirmationInvite = Confirmation("invite")
)

func (Confirmation) ElementType() reflect.Type {
	return reflect.TypeOf((*Confirmation)(nil)).Elem()
}

func (e Confirmation) ToConfirmationOutput() ConfirmationOutput {
	return pulumi.ToOutput(e).(ConfirmationOutput)
}

func (e Confirmation) ToConfirmationOutputWithContext(ctx context.Context) ConfirmationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConfirmationOutput)
}

func (e Confirmation) ToConfirmationPtrOutput() ConfirmationPtrOutput {
	return e.ToConfirmationPtrOutputWithContext(context.Background())
}

func (e Confirmation) ToConfirmationPtrOutputWithContext(ctx context.Context) ConfirmationPtrOutput {
	return Confirmation(e).ToConfirmationOutputWithContext(ctx).ToConfirmationPtrOutputWithContext(ctx)
}

func (e Confirmation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Confirmation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Confirmation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Confirmation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConfirmationOutput struct{ *pulumi.OutputState }

func (ConfirmationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Confirmation)(nil)).Elem()
}

func (o ConfirmationOutput) ToConfirmationOutput() ConfirmationOutput {
	return o
}

func (o ConfirmationOutput) ToConfirmationOutputWithContext(ctx context.Context) ConfirmationOutput {
	return o
}

func (o ConfirmationOutput) ToConfirmationPtrOutput() ConfirmationPtrOutput {
	return o.ToConfirmationPtrOutputWithContext(context.Background())
}

func (o ConfirmationOutput) ToConfirmationPtrOutputWithContext(ctx context.Context) ConfirmationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Confirmation) *Confirmation {
		return &v
	}).(ConfirmationPtrOutput)
}

func (o ConfirmationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConfirmationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Confirmation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConfirmationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfirmationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Confirmation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConfirmationPtrOutput struct{ *pulumi.OutputState }

func (ConfirmationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Confirmation)(nil)).Elem()
}

func (o ConfirmationPtrOutput) ToConfirmationPtrOutput() ConfirmationPtrOutput {
	return o
}

func (o ConfirmationPtrOutput) ToConfirmationPtrOutputWithContext(ctx context.Context) ConfirmationPtrOutput {
	return o
}

func (o ConfirmationPtrOutput) Elem() ConfirmationOutput {
	return o.ApplyT(func(v *Confirmation) Confirmation {
		if v != nil {
			return *v
		}
		var ret Confirmation
		return ret
	}).(ConfirmationOutput)
}

func (o ConfirmationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfirmationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Confirmation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConfirmationInput interface {
	pulumi.Input

	ToConfirmationOutput() ConfirmationOutput
	ToConfirmationOutputWithContext(context.Context) ConfirmationOutput
}

var confirmationPtrType = reflect.TypeOf((**Confirmation)(nil)).Elem()

type ConfirmationPtrInput interface {
	pulumi.Input

	ToConfirmationPtrOutput() ConfirmationPtrOutput
	ToConfirmationPtrOutputWithContext(context.Context) ConfirmationPtrOutput
}

type confirmationPtr string

func ConfirmationPtr(v string) ConfirmationPtrInput {
	return (*confirmationPtr)(&v)
}

func (*confirmationPtr) ElementType() reflect.Type {
	return confirmationPtrType
}

func (in *confirmationPtr) ToConfirmationPtrOutput() ConfirmationPtrOutput {
	return pulumi.ToOutput(in).(ConfirmationPtrOutput)
}

func (in *confirmationPtr) ToConfirmationPtrOutputWithContext(ctx context.Context) ConfirmationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConfirmationPtrOutput)
}

type ContentFormat string

const (
	// The contents are inline and Content type is a WADL document.
	ContentFormat_Wadl_xml = ContentFormat("wadl-xml")
	// The WADL document is hosted on a publicly accessible internet address.
	ContentFormat_Wadl_link_json = ContentFormat("wadl-link-json")
	// The contents are inline and Content Type is a OpenAPI 2.0 JSON Document.
	ContentFormat_Swagger_json = ContentFormat("swagger-json")
	// The OpenAPI 2.0 JSON document is hosted on a publicly accessible internet address.
	ContentFormat_Swagger_link_json = ContentFormat("swagger-link-json")
	// The contents are inline and the document is a WSDL/Soap document.
	ContentFormatWsdl = ContentFormat("wsdl")
	// The WSDL document is hosted on a publicly accessible internet address.
	ContentFormat_Wsdl_link = ContentFormat("wsdl-link")
	// The contents are inline and Content Type is a OpenAPI 3.0 YAML Document.
	ContentFormatOpenapi = ContentFormat("openapi")
	// The contents are inline and Content Type is a OpenAPI 3.0 JSON Document.
	ContentFormat_Openapi_json = ContentFormat("openapi+json")
	// The OpenAPI 3.0 YAML document is hosted on a publicly accessible internet address.
	ContentFormat_Openapi_link = ContentFormat("openapi-link")
	// The OpenAPI 3.0 JSON document is hosted on a publicly accessible internet address.
	ContentFormat_Openapi_json_link = ContentFormat("openapi+json-link")
	// The GraphQL API endpoint hosted on a publicly accessible internet address.
	ContentFormat_Graphql_link = ContentFormat("graphql-link")
)

func (ContentFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentFormat)(nil)).Elem()
}

func (e ContentFormat) ToContentFormatOutput() ContentFormatOutput {
	return pulumi.ToOutput(e).(ContentFormatOutput)
}

func (e ContentFormat) ToContentFormatOutputWithContext(ctx context.Context) ContentFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContentFormatOutput)
}

func (e ContentFormat) ToContentFormatPtrOutput() ContentFormatPtrOutput {
	return e.ToContentFormatPtrOutputWithContext(context.Background())
}

func (e ContentFormat) ToContentFormatPtrOutputWithContext(ctx context.Context) ContentFormatPtrOutput {
	return ContentFormat(e).ToContentFormatOutputWithContext(ctx).ToContentFormatPtrOutputWithContext(ctx)
}

func (e ContentFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContentFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContentFormatOutput struct{ *pulumi.OutputState }

func (ContentFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentFormat)(nil)).Elem()
}

func (o ContentFormatOutput) ToContentFormatOutput() ContentFormatOutput {
	return o
}

func (o ContentFormatOutput) ToContentFormatOutputWithContext(ctx context.Context) ContentFormatOutput {
	return o
}

func (o ContentFormatOutput) ToContentFormatPtrOutput() ContentFormatPtrOutput {
	return o.ToContentFormatPtrOutputWithContext(context.Background())
}

func (o ContentFormatOutput) ToContentFormatPtrOutputWithContext(ctx context.Context) ContentFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentFormat) *ContentFormat {
		return &v
	}).(ContentFormatPtrOutput)
}

func (o ContentFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContentFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContentFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContentFormatPtrOutput struct{ *pulumi.OutputState }

func (ContentFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentFormat)(nil)).Elem()
}

func (o ContentFormatPtrOutput) ToContentFormatPtrOutput() ContentFormatPtrOutput {
	return o
}

func (o ContentFormatPtrOutput) ToContentFormatPtrOutputWithContext(ctx context.Context) ContentFormatPtrOutput {
	return o
}

func (o ContentFormatPtrOutput) Elem() ContentFormatOutput {
	return o.ApplyT(func(v *ContentFormat) ContentFormat {
		if v != nil {
			return *v
		}
		var ret ContentFormat
		return ret
	}).(ContentFormatOutput)
}

func (o ContentFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContentFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContentFormatInput interface {
	pulumi.Input

	ToContentFormatOutput() ContentFormatOutput
	ToContentFormatOutputWithContext(context.Context) ContentFormatOutput
}

var contentFormatPtrType = reflect.TypeOf((**ContentFormat)(nil)).Elem()

type ContentFormatPtrInput interface {
	pulumi.Input

	ToContentFormatPtrOutput() ContentFormatPtrOutput
	ToContentFormatPtrOutputWithContext(context.Context) ContentFormatPtrOutput
}

type contentFormatPtr string

func ContentFormatPtr(v string) ContentFormatPtrInput {
	return (*contentFormatPtr)(&v)
}

func (*contentFormatPtr) ElementType() reflect.Type {
	return contentFormatPtrType
}

func (in *contentFormatPtr) ToContentFormatPtrOutput() ContentFormatPtrOutput {
	return pulumi.ToOutput(in).(ContentFormatPtrOutput)
}

func (in *contentFormatPtr) ToContentFormatPtrOutputWithContext(ctx context.Context) ContentFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContentFormatPtrOutput)
}

type DataMaskingMode string

const (
	// Mask the value of an entity.
	DataMaskingModeMask = DataMaskingMode("Mask")
	// Hide the presence of an entity.
	DataMaskingModeHide = DataMaskingMode("Hide")
)

func (DataMaskingMode) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMaskingMode)(nil)).Elem()
}

func (e DataMaskingMode) ToDataMaskingModeOutput() DataMaskingModeOutput {
	return pulumi.ToOutput(e).(DataMaskingModeOutput)
}

func (e DataMaskingMode) ToDataMaskingModeOutputWithContext(ctx context.Context) DataMaskingModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataMaskingModeOutput)
}

func (e DataMaskingMode) ToDataMaskingModePtrOutput() DataMaskingModePtrOutput {
	return e.ToDataMaskingModePtrOutputWithContext(context.Background())
}

func (e DataMaskingMode) ToDataMaskingModePtrOutputWithContext(ctx context.Context) DataMaskingModePtrOutput {
	return DataMaskingMode(e).ToDataMaskingModeOutputWithContext(ctx).ToDataMaskingModePtrOutputWithContext(ctx)
}

func (e DataMaskingMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataMaskingMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataMaskingMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataMaskingMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataMaskingModeOutput struct{ *pulumi.OutputState }

func (DataMaskingModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMaskingMode)(nil)).Elem()
}

func (o DataMaskingModeOutput) ToDataMaskingModeOutput() DataMaskingModeOutput {
	return o
}

func (o DataMaskingModeOutput) ToDataMaskingModeOutputWithContext(ctx context.Context) DataMaskingModeOutput {
	return o
}

func (o DataMaskingModeOutput) ToDataMaskingModePtrOutput() DataMaskingModePtrOutput {
	return o.ToDataMaskingModePtrOutputWithContext(context.Background())
}

func (o DataMaskingModeOutput) ToDataMaskingModePtrOutputWithContext(ctx context.Context) DataMaskingModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataMaskingMode) *DataMaskingMode {
		return &v
	}).(DataMaskingModePtrOutput)
}

func (o DataMaskingModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataMaskingModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataMaskingMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataMaskingModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataMaskingModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataMaskingMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataMaskingModePtrOutput struct{ *pulumi.OutputState }

func (DataMaskingModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataMaskingMode)(nil)).Elem()
}

func (o DataMaskingModePtrOutput) ToDataMaskingModePtrOutput() DataMaskingModePtrOutput {
	return o
}

func (o DataMaskingModePtrOutput) ToDataMaskingModePtrOutputWithContext(ctx context.Context) DataMaskingModePtrOutput {
	return o
}

func (o DataMaskingModePtrOutput) Elem() DataMaskingModeOutput {
	return o.ApplyT(func(v *DataMaskingMode) DataMaskingMode {
		if v != nil {
			return *v
		}
		var ret DataMaskingMode
		return ret
	}).(DataMaskingModeOutput)
}

func (o DataMaskingModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataMaskingModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataMaskingMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataMaskingModeInput interface {
	pulumi.Input

	ToDataMaskingModeOutput() DataMaskingModeOutput
	ToDataMaskingModeOutputWithContext(context.Context) DataMaskingModeOutput
}

var dataMaskingModePtrType = reflect.TypeOf((**DataMaskingMode)(nil)).Elem()

type DataMaskingModePtrInput interface {
	pulumi.Input

	ToDataMaskingModePtrOutput() DataMaskingModePtrOutput
	ToDataMaskingModePtrOutputWithContext(context.Context) DataMaskingModePtrOutput
}

type dataMaskingModePtr string

func DataMaskingModePtr(v string) DataMaskingModePtrInput {
	return (*dataMaskingModePtr)(&v)
}

func (*dataMaskingModePtr) ElementType() reflect.Type {
	return dataMaskingModePtrType
}

func (in *dataMaskingModePtr) ToDataMaskingModePtrOutput() DataMaskingModePtrOutput {
	return pulumi.ToOutput(in).(DataMaskingModePtrOutput)
}

func (in *dataMaskingModePtr) ToDataMaskingModePtrOutputWithContext(ctx context.Context) DataMaskingModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataMaskingModePtrOutput)
}

type GrantType string

const (
	// Authorization Code Grant flow as described https://tools.ietf.org/html/rfc6749#section-4.1.
	GrantTypeAuthorizationCode = GrantType("authorizationCode")
	// Implicit Code Grant flow as described https://tools.ietf.org/html/rfc6749#section-4.2.
	GrantTypeImplicit = GrantType("implicit")
	// Resource Owner Password Grant flow as described https://tools.ietf.org/html/rfc6749#section-4.3.
	GrantTypeResourceOwnerPassword = GrantType("resourceOwnerPassword")
	// Client Credentials Grant flow as described https://tools.ietf.org/html/rfc6749#section-4.4.
	GrantTypeClientCredentials = GrantType("clientCredentials")
)

func (GrantType) ElementType() reflect.Type {
	return reflect.TypeOf((*GrantType)(nil)).Elem()
}

func (e GrantType) ToGrantTypeOutput() GrantTypeOutput {
	return pulumi.ToOutput(e).(GrantTypeOutput)
}

func (e GrantType) ToGrantTypeOutputWithContext(ctx context.Context) GrantTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GrantTypeOutput)
}

func (e GrantType) ToGrantTypePtrOutput() GrantTypePtrOutput {
	return e.ToGrantTypePtrOutputWithContext(context.Background())
}

func (e GrantType) ToGrantTypePtrOutputWithContext(ctx context.Context) GrantTypePtrOutput {
	return GrantType(e).ToGrantTypeOutputWithContext(ctx).ToGrantTypePtrOutputWithContext(ctx)
}

func (e GrantType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GrantType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GrantType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GrantType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GrantTypeOutput struct{ *pulumi.OutputState }

func (GrantTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GrantType)(nil)).Elem()
}

func (o GrantTypeOutput) ToGrantTypeOutput() GrantTypeOutput {
	return o
}

func (o GrantTypeOutput) ToGrantTypeOutputWithContext(ctx context.Context) GrantTypeOutput {
	return o
}

func (o GrantTypeOutput) ToGrantTypePtrOutput() GrantTypePtrOutput {
	return o.ToGrantTypePtrOutputWithContext(context.Background())
}

func (o GrantTypeOutput) ToGrantTypePtrOutputWithContext(ctx context.Context) GrantTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GrantType) *GrantType {
		return &v
	}).(GrantTypePtrOutput)
}

func (o GrantTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GrantTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GrantType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GrantTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GrantTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GrantType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GrantTypePtrOutput struct{ *pulumi.OutputState }

func (GrantTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GrantType)(nil)).Elem()
}

func (o GrantTypePtrOutput) ToGrantTypePtrOutput() GrantTypePtrOutput {
	return o
}

func (o GrantTypePtrOutput) ToGrantTypePtrOutputWithContext(ctx context.Context) GrantTypePtrOutput {
	return o
}

func (o GrantTypePtrOutput) Elem() GrantTypeOutput {
	return o.ApplyT(func(v *GrantType) GrantType {
		if v != nil {
			return *v
		}
		var ret GrantType
		return ret
	}).(GrantTypeOutput)
}

func (o GrantTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GrantTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GrantType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GrantTypeInput interface {
	pulumi.Input

	ToGrantTypeOutput() GrantTypeOutput
	ToGrantTypeOutputWithContext(context.Context) GrantTypeOutput
}

var grantTypePtrType = reflect.TypeOf((**GrantType)(nil)).Elem()

type GrantTypePtrInput interface {
	pulumi.Input

	ToGrantTypePtrOutput() GrantTypePtrOutput
	ToGrantTypePtrOutputWithContext(context.Context) GrantTypePtrOutput
}

type grantTypePtr string

func GrantTypePtr(v string) GrantTypePtrInput {
	return (*grantTypePtr)(&v)
}

func (*grantTypePtr) ElementType() reflect.Type {
	return grantTypePtrType
}

func (in *grantTypePtr) ToGrantTypePtrOutput() GrantTypePtrOutput {
	return pulumi.ToOutput(in).(GrantTypePtrOutput)
}

func (in *grantTypePtr) ToGrantTypePtrOutputWithContext(ctx context.Context) GrantTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GrantTypePtrOutput)
}

type GroupType string

const (
	GroupTypeCustom   = GroupType("custom")
	GroupTypeSystem   = GroupType("system")
	GroupTypeExternal = GroupType("external")
)

func (GroupType) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupType)(nil)).Elem()
}

func (e GroupType) ToGroupTypeOutput() GroupTypeOutput {
	return pulumi.ToOutput(e).(GroupTypeOutput)
}

func (e GroupType) ToGroupTypeOutputWithContext(ctx context.Context) GroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GroupTypeOutput)
}

func (e GroupType) ToGroupTypePtrOutput() GroupTypePtrOutput {
	return e.ToGroupTypePtrOutputWithContext(context.Background())
}

func (e GroupType) ToGroupTypePtrOutputWithContext(ctx context.Context) GroupTypePtrOutput {
	return GroupType(e).ToGroupTypeOutputWithContext(ctx).ToGroupTypePtrOutputWithContext(ctx)
}

func (e GroupType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GroupType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GroupType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GroupType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GroupTypeOutput struct{ *pulumi.OutputState }

func (GroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupType)(nil)).Elem()
}

func (o GroupTypeOutput) ToGroupTypeOutput() GroupTypeOutput {
	return o
}

func (o GroupTypeOutput) ToGroupTypeOutputWithContext(ctx context.Context) GroupTypeOutput {
	return o
}

func (o GroupTypeOutput) ToGroupTypePtrOutput() GroupTypePtrOutput {
	return o.ToGroupTypePtrOutputWithContext(context.Background())
}

func (o GroupTypeOutput) ToGroupTypePtrOutputWithContext(ctx context.Context) GroupTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupType) *GroupType {
		return &v
	}).(GroupTypePtrOutput)
}

func (o GroupTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GroupTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GroupType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GroupTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GroupTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GroupType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GroupTypePtrOutput struct{ *pulumi.OutputState }

func (GroupTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupType)(nil)).Elem()
}

func (o GroupTypePtrOutput) ToGroupTypePtrOutput() GroupTypePtrOutput {
	return o
}

func (o GroupTypePtrOutput) ToGroupTypePtrOutputWithContext(ctx context.Context) GroupTypePtrOutput {
	return o
}

func (o GroupTypePtrOutput) Elem() GroupTypeOutput {
	return o.ApplyT(func(v *GroupType) GroupType {
		if v != nil {
			return *v
		}
		var ret GroupType
		return ret
	}).(GroupTypeOutput)
}

func (o GroupTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GroupTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GroupType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GroupTypeInput interface {
	pulumi.Input

	ToGroupTypeOutput() GroupTypeOutput
	ToGroupTypeOutputWithContext(context.Context) GroupTypeOutput
}

var groupTypePtrType = reflect.TypeOf((**GroupType)(nil)).Elem()

type GroupTypePtrInput interface {
	pulumi.Input

	ToGroupTypePtrOutput() GroupTypePtrOutput
	ToGroupTypePtrOutputWithContext(context.Context) GroupTypePtrOutput
}

type groupTypePtr string

func GroupTypePtr(v string) GroupTypePtrInput {
	return (*groupTypePtr)(&v)
}

func (*groupTypePtr) ElementType() reflect.Type {
	return groupTypePtrType
}

func (in *groupTypePtr) ToGroupTypePtrOutput() GroupTypePtrOutput {
	return pulumi.ToOutput(in).(GroupTypePtrOutput)
}

func (in *groupTypePtr) ToGroupTypePtrOutputWithContext(ctx context.Context) GroupTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GroupTypePtrOutput)
}

type HostnameType string

const (
	HostnameTypeProxy           = HostnameType("Proxy")
	HostnameTypePortal          = HostnameType("Portal")
	HostnameTypeManagement      = HostnameType("Management")
	HostnameTypeScm             = HostnameType("Scm")
	HostnameTypeDeveloperPortal = HostnameType("DeveloperPortal")
)

func (HostnameType) ElementType() reflect.Type {
	return reflect.TypeOf((*HostnameType)(nil)).Elem()
}

func (e HostnameType) ToHostnameTypeOutput() HostnameTypeOutput {
	return pulumi.ToOutput(e).(HostnameTypeOutput)
}

func (e HostnameType) ToHostnameTypeOutputWithContext(ctx context.Context) HostnameTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostnameTypeOutput)
}

func (e HostnameType) ToHostnameTypePtrOutput() HostnameTypePtrOutput {
	return e.ToHostnameTypePtrOutputWithContext(context.Background())
}

func (e HostnameType) ToHostnameTypePtrOutputWithContext(ctx context.Context) HostnameTypePtrOutput {
	return HostnameType(e).ToHostnameTypeOutputWithContext(ctx).ToHostnameTypePtrOutputWithContext(ctx)
}

func (e HostnameType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostnameType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostnameType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostnameType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostnameTypeOutput struct{ *pulumi.OutputState }

func (HostnameTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostnameType)(nil)).Elem()
}

func (o HostnameTypeOutput) ToHostnameTypeOutput() HostnameTypeOutput {
	return o
}

func (o HostnameTypeOutput) ToHostnameTypeOutputWithContext(ctx context.Context) HostnameTypeOutput {
	return o
}

func (o HostnameTypeOutput) ToHostnameTypePtrOutput() HostnameTypePtrOutput {
	return o.ToHostnameTypePtrOutputWithContext(context.Background())
}

func (o HostnameTypeOutput) ToHostnameTypePtrOutputWithContext(ctx context.Context) HostnameTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostnameType) *HostnameType {
		return &v
	}).(HostnameTypePtrOutput)
}

func (o HostnameTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostnameTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostnameType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostnameTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostnameTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostnameType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostnameTypePtrOutput struct{ *pulumi.OutputState }

func (HostnameTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostnameType)(nil)).Elem()
}

func (o HostnameTypePtrOutput) ToHostnameTypePtrOutput() HostnameTypePtrOutput {
	return o
}

func (o HostnameTypePtrOutput) ToHostnameTypePtrOutputWithContext(ctx context.Context) HostnameTypePtrOutput {
	return o
}

func (o HostnameTypePtrOutput) Elem() HostnameTypeOutput {
	return o.ApplyT(func(v *HostnameType) HostnameType {
		if v != nil {
			return *v
		}
		var ret HostnameType
		return ret
	}).(HostnameTypeOutput)
}

func (o HostnameTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostnameTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostnameType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostnameTypeInput interface {
	pulumi.Input

	ToHostnameTypeOutput() HostnameTypeOutput
	ToHostnameTypeOutputWithContext(context.Context) HostnameTypeOutput
}

var hostnameTypePtrType = reflect.TypeOf((**HostnameType)(nil)).Elem()

type HostnameTypePtrInput interface {
	pulumi.Input

	ToHostnameTypePtrOutput() HostnameTypePtrOutput
	ToHostnameTypePtrOutputWithContext(context.Context) HostnameTypePtrOutput
}

type hostnameTypePtr string

func HostnameTypePtr(v string) HostnameTypePtrInput {
	return (*hostnameTypePtr)(&v)
}

func (*hostnameTypePtr) ElementType() reflect.Type {
	return hostnameTypePtrType
}

func (in *hostnameTypePtr) ToHostnameTypePtrOutput() HostnameTypePtrOutput {
	return pulumi.ToOutput(in).(HostnameTypePtrOutput)
}

func (in *hostnameTypePtr) ToHostnameTypePtrOutputWithContext(ctx context.Context) HostnameTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostnameTypePtrOutput)
}

type HttpCorrelationProtocol string

const (
	// Do not read and inject correlation headers.
	HttpCorrelationProtocolNone = HttpCorrelationProtocol("None")
	// Inject Request-Id and Request-Context headers with request correlation data. See https://github.com/dotnet/corefx/blob/master/src/System.Diagnostics.DiagnosticSource/src/HttpCorrelationProtocol.md.
	HttpCorrelationProtocolLegacy = HttpCorrelationProtocol("Legacy")
	// Inject Trace Context headers. See https://w3c.github.io/trace-context.
	HttpCorrelationProtocolW3C = HttpCorrelationProtocol("W3C")
)

func (HttpCorrelationProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpCorrelationProtocol)(nil)).Elem()
}

func (e HttpCorrelationProtocol) ToHttpCorrelationProtocolOutput() HttpCorrelationProtocolOutput {
	return pulumi.ToOutput(e).(HttpCorrelationProtocolOutput)
}

func (e HttpCorrelationProtocol) ToHttpCorrelationProtocolOutputWithContext(ctx context.Context) HttpCorrelationProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HttpCorrelationProtocolOutput)
}

func (e HttpCorrelationProtocol) ToHttpCorrelationProtocolPtrOutput() HttpCorrelationProtocolPtrOutput {
	return e.ToHttpCorrelationProtocolPtrOutputWithContext(context.Background())
}

func (e HttpCorrelationProtocol) ToHttpCorrelationProtocolPtrOutputWithContext(ctx context.Context) HttpCorrelationProtocolPtrOutput {
	return HttpCorrelationProtocol(e).ToHttpCorrelationProtocolOutputWithContext(ctx).ToHttpCorrelationProtocolPtrOutputWithContext(ctx)
}

func (e HttpCorrelationProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HttpCorrelationProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HttpCorrelationProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HttpCorrelationProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HttpCorrelationProtocolOutput struct{ *pulumi.OutputState }

func (HttpCorrelationProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpCorrelationProtocol)(nil)).Elem()
}

func (o HttpCorrelationProtocolOutput) ToHttpCorrelationProtocolOutput() HttpCorrelationProtocolOutput {
	return o
}

func (o HttpCorrelationProtocolOutput) ToHttpCorrelationProtocolOutputWithContext(ctx context.Context) HttpCorrelationProtocolOutput {
	return o
}

func (o HttpCorrelationProtocolOutput) ToHttpCorrelationProtocolPtrOutput() HttpCorrelationProtocolPtrOutput {
	return o.ToHttpCorrelationProtocolPtrOutputWithContext(context.Background())
}

func (o HttpCorrelationProtocolOutput) ToHttpCorrelationProtocolPtrOutputWithContext(ctx context.Context) HttpCorrelationProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpCorrelationProtocol) *HttpCorrelationProtocol {
		return &v
	}).(HttpCorrelationProtocolPtrOutput)
}

func (o HttpCorrelationProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HttpCorrelationProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HttpCorrelationProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HttpCorrelationProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HttpCorrelationProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HttpCorrelationProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HttpCorrelationProtocolPtrOutput struct{ *pulumi.OutputState }

func (HttpCorrelationProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpCorrelationProtocol)(nil)).Elem()
}

func (o HttpCorrelationProtocolPtrOutput) ToHttpCorrelationProtocolPtrOutput() HttpCorrelationProtocolPtrOutput {
	return o
}

func (o HttpCorrelationProtocolPtrOutput) ToHttpCorrelationProtocolPtrOutputWithContext(ctx context.Context) HttpCorrelationProtocolPtrOutput {
	return o
}

func (o HttpCorrelationProtocolPtrOutput) Elem() HttpCorrelationProtocolOutput {
	return o.ApplyT(func(v *HttpCorrelationProtocol) HttpCorrelationProtocol {
		if v != nil {
			return *v
		}
		var ret HttpCorrelationProtocol
		return ret
	}).(HttpCorrelationProtocolOutput)
}

func (o HttpCorrelationProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HttpCorrelationProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HttpCorrelationProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HttpCorrelationProtocolInput interface {
	pulumi.Input

	ToHttpCorrelationProtocolOutput() HttpCorrelationProtocolOutput
	ToHttpCorrelationProtocolOutputWithContext(context.Context) HttpCorrelationProtocolOutput
}

var httpCorrelationProtocolPtrType = reflect.TypeOf((**HttpCorrelationProtocol)(nil)).Elem()

type HttpCorrelationProtocolPtrInput interface {
	pulumi.Input

	ToHttpCorrelationProtocolPtrOutput() HttpCorrelationProtocolPtrOutput
	ToHttpCorrelationProtocolPtrOutputWithContext(context.Context) HttpCorrelationProtocolPtrOutput
}

type httpCorrelationProtocolPtr string

func HttpCorrelationProtocolPtr(v string) HttpCorrelationProtocolPtrInput {
	return (*httpCorrelationProtocolPtr)(&v)
}

func (*httpCorrelationProtocolPtr) ElementType() reflect.Type {
	return httpCorrelationProtocolPtrType
}

func (in *httpCorrelationProtocolPtr) ToHttpCorrelationProtocolPtrOutput() HttpCorrelationProtocolPtrOutput {
	return pulumi.ToOutput(in).(HttpCorrelationProtocolPtrOutput)
}

func (in *httpCorrelationProtocolPtr) ToHttpCorrelationProtocolPtrOutputWithContext(ctx context.Context) HttpCorrelationProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HttpCorrelationProtocolPtrOutput)
}

type IdentityProviderType string

const (
	// Facebook as Identity provider.
	IdentityProviderTypeFacebook = IdentityProviderType("facebook")
	// Google as Identity provider.
	IdentityProviderTypeGoogle = IdentityProviderType("google")
	// Microsoft Live as Identity provider.
	IdentityProviderTypeMicrosoft = IdentityProviderType("microsoft")
	// Twitter as Identity provider.
	IdentityProviderTypeTwitter = IdentityProviderType("twitter")
	// Azure Active Directory as Identity provider.
	IdentityProviderTypeAad = IdentityProviderType("aad")
	// Azure Active Directory B2C as Identity provider.
	IdentityProviderTypeAadB2C = IdentityProviderType("aadB2C")
)

func (IdentityProviderType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderType)(nil)).Elem()
}

func (e IdentityProviderType) ToIdentityProviderTypeOutput() IdentityProviderTypeOutput {
	return pulumi.ToOutput(e).(IdentityProviderTypeOutput)
}

func (e IdentityProviderType) ToIdentityProviderTypeOutputWithContext(ctx context.Context) IdentityProviderTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityProviderTypeOutput)
}

func (e IdentityProviderType) ToIdentityProviderTypePtrOutput() IdentityProviderTypePtrOutput {
	return e.ToIdentityProviderTypePtrOutputWithContext(context.Background())
}

func (e IdentityProviderType) ToIdentityProviderTypePtrOutputWithContext(ctx context.Context) IdentityProviderTypePtrOutput {
	return IdentityProviderType(e).ToIdentityProviderTypeOutputWithContext(ctx).ToIdentityProviderTypePtrOutputWithContext(ctx)
}

func (e IdentityProviderType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityProviderType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityProviderType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityProviderType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityProviderTypeOutput struct{ *pulumi.OutputState }

func (IdentityProviderTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderType)(nil)).Elem()
}

func (o IdentityProviderTypeOutput) ToIdentityProviderTypeOutput() IdentityProviderTypeOutput {
	return o
}

func (o IdentityProviderTypeOutput) ToIdentityProviderTypeOutputWithContext(ctx context.Context) IdentityProviderTypeOutput {
	return o
}

func (o IdentityProviderTypeOutput) ToIdentityProviderTypePtrOutput() IdentityProviderTypePtrOutput {
	return o.ToIdentityProviderTypePtrOutputWithContext(context.Background())
}

func (o IdentityProviderTypeOutput) ToIdentityProviderTypePtrOutputWithContext(ctx context.Context) IdentityProviderTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProviderType) *IdentityProviderType {
		return &v
	}).(IdentityProviderTypePtrOutput)
}

func (o IdentityProviderTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityProviderTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityProviderType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityProviderTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityProviderTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityProviderType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityProviderTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityProviderTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderType)(nil)).Elem()
}

func (o IdentityProviderTypePtrOutput) ToIdentityProviderTypePtrOutput() IdentityProviderTypePtrOutput {
	return o
}

func (o IdentityProviderTypePtrOutput) ToIdentityProviderTypePtrOutputWithContext(ctx context.Context) IdentityProviderTypePtrOutput {
	return o
}

func (o IdentityProviderTypePtrOutput) Elem() IdentityProviderTypeOutput {
	return o.ApplyT(func(v *IdentityProviderType) IdentityProviderType {
		if v != nil {
			return *v
		}
		var ret IdentityProviderType
		return ret
	}).(IdentityProviderTypeOutput)
}

func (o IdentityProviderTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityProviderTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityProviderType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IdentityProviderTypeInput interface {
	pulumi.Input

	ToIdentityProviderTypeOutput() IdentityProviderTypeOutput
	ToIdentityProviderTypeOutputWithContext(context.Context) IdentityProviderTypeOutput
}

var identityProviderTypePtrType = reflect.TypeOf((**IdentityProviderType)(nil)).Elem()

type IdentityProviderTypePtrInput interface {
	pulumi.Input

	ToIdentityProviderTypePtrOutput() IdentityProviderTypePtrOutput
	ToIdentityProviderTypePtrOutputWithContext(context.Context) IdentityProviderTypePtrOutput
}

type identityProviderTypePtr string

func IdentityProviderTypePtr(v string) IdentityProviderTypePtrInput {
	return (*identityProviderTypePtr)(&v)
}

func (*identityProviderTypePtr) ElementType() reflect.Type {
	return identityProviderTypePtrType
}

func (in *identityProviderTypePtr) ToIdentityProviderTypePtrOutput() IdentityProviderTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityProviderTypePtrOutput)
}

func (in *identityProviderTypePtr) ToIdentityProviderTypePtrOutputWithContext(ctx context.Context) IdentityProviderTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityProviderTypePtrOutput)
}

type KeyType string

const (
	KeyTypePrimary   = KeyType("primary")
	KeyTypeSecondary = KeyType("secondary")
)

func (KeyType) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyType)(nil)).Elem()
}

func (e KeyType) ToKeyTypeOutput() KeyTypeOutput {
	return pulumi.ToOutput(e).(KeyTypeOutput)
}

func (e KeyType) ToKeyTypeOutputWithContext(ctx context.Context) KeyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeyTypeOutput)
}

func (e KeyType) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return e.ToKeyTypePtrOutputWithContext(context.Background())
}

func (e KeyType) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return KeyType(e).ToKeyTypeOutputWithContext(ctx).ToKeyTypePtrOutputWithContext(ctx)
}

func (e KeyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeyTypeOutput struct{ *pulumi.OutputState }

func (KeyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyType)(nil)).Elem()
}

func (o KeyTypeOutput) ToKeyTypeOutput() KeyTypeOutput {
	return o
}

func (o KeyTypeOutput) ToKeyTypeOutputWithContext(ctx context.Context) KeyTypeOutput {
	return o
}

func (o KeyTypeOutput) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return o.ToKeyTypePtrOutputWithContext(context.Background())
}

func (o KeyTypeOutput) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyType) *KeyType {
		return &v
	}).(KeyTypePtrOutput)
}

func (o KeyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeyTypePtrOutput struct{ *pulumi.OutputState }

func (KeyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyType)(nil)).Elem()
}

func (o KeyTypePtrOutput) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return o
}

func (o KeyTypePtrOutput) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return o
}

func (o KeyTypePtrOutput) Elem() KeyTypeOutput {
	return o.ApplyT(func(v *KeyType) KeyType {
		if v != nil {
			return *v
		}
		var ret KeyType
		return ret
	}).(KeyTypeOutput)
}

func (o KeyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KeyTypeInput interface {
	pulumi.Input

	ToKeyTypeOutput() KeyTypeOutput
	ToKeyTypeOutputWithContext(context.Context) KeyTypeOutput
}

var keyTypePtrType = reflect.TypeOf((**KeyType)(nil)).Elem()

type KeyTypePtrInput interface {
	pulumi.Input

	ToKeyTypePtrOutput() KeyTypePtrOutput
	ToKeyTypePtrOutputWithContext(context.Context) KeyTypePtrOutput
}

type keyTypePtr string

func KeyTypePtr(v string) KeyTypePtrInput {
	return (*keyTypePtr)(&v)
}

func (*keyTypePtr) ElementType() reflect.Type {
	return keyTypePtrType
}

func (in *keyTypePtr) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return pulumi.ToOutput(in).(KeyTypePtrOutput)
}

func (in *keyTypePtr) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeyTypePtrOutput)
}

type LoggerType string

const (
	// Azure Event Hub as log destination.
	LoggerTypeAzureEventHub = LoggerType("azureEventHub")
	// Azure Application Insights as log destination.
	LoggerTypeApplicationInsights = LoggerType("applicationInsights")
	// Azure Monitor
	LoggerTypeAzureMonitor = LoggerType("azureMonitor")
)

func (LoggerType) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggerType)(nil)).Elem()
}

func (e LoggerType) ToLoggerTypeOutput() LoggerTypeOutput {
	return pulumi.ToOutput(e).(LoggerTypeOutput)
}

func (e LoggerType) ToLoggerTypeOutputWithContext(ctx context.Context) LoggerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LoggerTypeOutput)
}

func (e LoggerType) ToLoggerTypePtrOutput() LoggerTypePtrOutput {
	return e.ToLoggerTypePtrOutputWithContext(context.Background())
}

func (e LoggerType) ToLoggerTypePtrOutputWithContext(ctx context.Context) LoggerTypePtrOutput {
	return LoggerType(e).ToLoggerTypeOutputWithContext(ctx).ToLoggerTypePtrOutputWithContext(ctx)
}

func (e LoggerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoggerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoggerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoggerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LoggerTypeOutput struct{ *pulumi.OutputState }

func (LoggerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggerType)(nil)).Elem()
}

func (o LoggerTypeOutput) ToLoggerTypeOutput() LoggerTypeOutput {
	return o
}

func (o LoggerTypeOutput) ToLoggerTypeOutputWithContext(ctx context.Context) LoggerTypeOutput {
	return o
}

func (o LoggerTypeOutput) ToLoggerTypePtrOutput() LoggerTypePtrOutput {
	return o.ToLoggerTypePtrOutputWithContext(context.Background())
}

func (o LoggerTypeOutput) ToLoggerTypePtrOutputWithContext(ctx context.Context) LoggerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoggerType) *LoggerType {
		return &v
	}).(LoggerTypePtrOutput)
}

func (o LoggerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LoggerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoggerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LoggerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoggerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoggerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LoggerTypePtrOutput struct{ *pulumi.OutputState }

func (LoggerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggerType)(nil)).Elem()
}

func (o LoggerTypePtrOutput) ToLoggerTypePtrOutput() LoggerTypePtrOutput {
	return o
}

func (o LoggerTypePtrOutput) ToLoggerTypePtrOutputWithContext(ctx context.Context) LoggerTypePtrOutput {
	return o
}

func (o LoggerTypePtrOutput) Elem() LoggerTypeOutput {
	return o.ApplyT(func(v *LoggerType) LoggerType {
		if v != nil {
			return *v
		}
		var ret LoggerType
		return ret
	}).(LoggerTypeOutput)
}

func (o LoggerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoggerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LoggerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LoggerTypeInput interface {
	pulumi.Input

	ToLoggerTypeOutput() LoggerTypeOutput
	ToLoggerTypeOutputWithContext(context.Context) LoggerTypeOutput
}

var loggerTypePtrType = reflect.TypeOf((**LoggerType)(nil)).Elem()

type LoggerTypePtrInput interface {
	pulumi.Input

	ToLoggerTypePtrOutput() LoggerTypePtrOutput
	ToLoggerTypePtrOutputWithContext(context.Context) LoggerTypePtrOutput
}

type loggerTypePtr string

func LoggerTypePtr(v string) LoggerTypePtrInput {
	return (*loggerTypePtr)(&v)
}

func (*loggerTypePtr) ElementType() reflect.Type {
	return loggerTypePtrType
}

func (in *loggerTypePtr) ToLoggerTypePtrOutput() LoggerTypePtrOutput {
	return pulumi.ToOutput(in).(LoggerTypePtrOutput)
}

func (in *loggerTypePtr) ToLoggerTypePtrOutputWithContext(ctx context.Context) LoggerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LoggerTypePtrOutput)
}

type OperationNameFormat string

const (
	// API_NAME;rev=API_REVISION - OPERATION_NAME
	OperationNameFormatName = OperationNameFormat("Name")
	// HTTP_VERB URL
	OperationNameFormatUrl = OperationNameFormat("Url")
)

func (OperationNameFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationNameFormat)(nil)).Elem()
}

func (e OperationNameFormat) ToOperationNameFormatOutput() OperationNameFormatOutput {
	return pulumi.ToOutput(e).(OperationNameFormatOutput)
}

func (e OperationNameFormat) ToOperationNameFormatOutputWithContext(ctx context.Context) OperationNameFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperationNameFormatOutput)
}

func (e OperationNameFormat) ToOperationNameFormatPtrOutput() OperationNameFormatPtrOutput {
	return e.ToOperationNameFormatPtrOutputWithContext(context.Background())
}

func (e OperationNameFormat) ToOperationNameFormatPtrOutputWithContext(ctx context.Context) OperationNameFormatPtrOutput {
	return OperationNameFormat(e).ToOperationNameFormatOutputWithContext(ctx).ToOperationNameFormatPtrOutputWithContext(ctx)
}

func (e OperationNameFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperationNameFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperationNameFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperationNameFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperationNameFormatOutput struct{ *pulumi.OutputState }

func (OperationNameFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationNameFormat)(nil)).Elem()
}

func (o OperationNameFormatOutput) ToOperationNameFormatOutput() OperationNameFormatOutput {
	return o
}

func (o OperationNameFormatOutput) ToOperationNameFormatOutputWithContext(ctx context.Context) OperationNameFormatOutput {
	return o
}

func (o OperationNameFormatOutput) ToOperationNameFormatPtrOutput() OperationNameFormatPtrOutput {
	return o.ToOperationNameFormatPtrOutputWithContext(context.Background())
}

func (o OperationNameFormatOutput) ToOperationNameFormatPtrOutputWithContext(ctx context.Context) OperationNameFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperationNameFormat) *OperationNameFormat {
		return &v
	}).(OperationNameFormatPtrOutput)
}

func (o OperationNameFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperationNameFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperationNameFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperationNameFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperationNameFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperationNameFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperationNameFormatPtrOutput struct{ *pulumi.OutputState }

func (OperationNameFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationNameFormat)(nil)).Elem()
}

func (o OperationNameFormatPtrOutput) ToOperationNameFormatPtrOutput() OperationNameFormatPtrOutput {
	return o
}

func (o OperationNameFormatPtrOutput) ToOperationNameFormatPtrOutputWithContext(ctx context.Context) OperationNameFormatPtrOutput {
	return o
}

func (o OperationNameFormatPtrOutput) Elem() OperationNameFormatOutput {
	return o.ApplyT(func(v *OperationNameFormat) OperationNameFormat {
		if v != nil {
			return *v
		}
		var ret OperationNameFormat
		return ret
	}).(OperationNameFormatOutput)
}

func (o OperationNameFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperationNameFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperationNameFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperationNameFormatInput interface {
	pulumi.Input

	ToOperationNameFormatOutput() OperationNameFormatOutput
	ToOperationNameFormatOutputWithContext(context.Context) OperationNameFormatOutput
}

var operationNameFormatPtrType = reflect.TypeOf((**OperationNameFormat)(nil)).Elem()

type OperationNameFormatPtrInput interface {
	pulumi.Input

	ToOperationNameFormatPtrOutput() OperationNameFormatPtrOutput
	ToOperationNameFormatPtrOutputWithContext(context.Context) OperationNameFormatPtrOutput
}

type operationNameFormatPtr string

func OperationNameFormatPtr(v string) OperationNameFormatPtrInput {
	return (*operationNameFormatPtr)(&v)
}

func (*operationNameFormatPtr) ElementType() reflect.Type {
	return operationNameFormatPtrType
}

func (in *operationNameFormatPtr) ToOperationNameFormatPtrOutput() OperationNameFormatPtrOutput {
	return pulumi.ToOutput(in).(OperationNameFormatPtrOutput)
}

func (in *operationNameFormatPtr) ToOperationNameFormatPtrOutputWithContext(ctx context.Context) OperationNameFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperationNameFormatPtrOutput)
}

type PolicyContentFormat string

const (
	// The contents are inline and Content type is an XML document.
	PolicyContentFormatXml = PolicyContentFormat("xml")
	// The policy XML document is hosted on a http endpoint accessible from the API Management service.
	PolicyContentFormat_Xml_link = PolicyContentFormat("xml-link")
	// The contents are inline and Content type is a non XML encoded policy document.
	PolicyContentFormatRawxml = PolicyContentFormat("rawxml")
	// The policy document is not Xml encoded and is hosted on a http endpoint accessible from the API Management service.
	PolicyContentFormat_Rawxml_link = PolicyContentFormat("rawxml-link")
)

func (PolicyContentFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyContentFormat)(nil)).Elem()
}

func (e PolicyContentFormat) ToPolicyContentFormatOutput() PolicyContentFormatOutput {
	return pulumi.ToOutput(e).(PolicyContentFormatOutput)
}

func (e PolicyContentFormat) ToPolicyContentFormatOutputWithContext(ctx context.Context) PolicyContentFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolicyContentFormatOutput)
}

func (e PolicyContentFormat) ToPolicyContentFormatPtrOutput() PolicyContentFormatPtrOutput {
	return e.ToPolicyContentFormatPtrOutputWithContext(context.Background())
}

func (e PolicyContentFormat) ToPolicyContentFormatPtrOutputWithContext(ctx context.Context) PolicyContentFormatPtrOutput {
	return PolicyContentFormat(e).ToPolicyContentFormatOutputWithContext(ctx).ToPolicyContentFormatPtrOutputWithContext(ctx)
}

func (e PolicyContentFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyContentFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyContentFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolicyContentFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolicyContentFormatOutput struct{ *pulumi.OutputState }

func (PolicyContentFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyContentFormat)(nil)).Elem()
}

func (o PolicyContentFormatOutput) ToPolicyContentFormatOutput() PolicyContentFormatOutput {
	return o
}

func (o PolicyContentFormatOutput) ToPolicyContentFormatOutputWithContext(ctx context.Context) PolicyContentFormatOutput {
	return o
}

func (o PolicyContentFormatOutput) ToPolicyContentFormatPtrOutput() PolicyContentFormatPtrOutput {
	return o.ToPolicyContentFormatPtrOutputWithContext(context.Background())
}

func (o PolicyContentFormatOutput) ToPolicyContentFormatPtrOutputWithContext(ctx context.Context) PolicyContentFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyContentFormat) *PolicyContentFormat {
		return &v
	}).(PolicyContentFormatPtrOutput)
}

func (o PolicyContentFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyContentFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyContentFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyContentFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyContentFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyContentFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyContentFormatPtrOutput struct{ *pulumi.OutputState }

func (PolicyContentFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyContentFormat)(nil)).Elem()
}

func (o PolicyContentFormatPtrOutput) ToPolicyContentFormatPtrOutput() PolicyContentFormatPtrOutput {
	return o
}

func (o PolicyContentFormatPtrOutput) ToPolicyContentFormatPtrOutputWithContext(ctx context.Context) PolicyContentFormatPtrOutput {
	return o
}

func (o PolicyContentFormatPtrOutput) Elem() PolicyContentFormatOutput {
	return o.ApplyT(func(v *PolicyContentFormat) PolicyContentFormat {
		if v != nil {
			return *v
		}
		var ret PolicyContentFormat
		return ret
	}).(PolicyContentFormatOutput)
}

func (o PolicyContentFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyContentFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyContentFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PolicyContentFormatInput interface {
	pulumi.Input

	ToPolicyContentFormatOutput() PolicyContentFormatOutput
	ToPolicyContentFormatOutputWithContext(context.Context) PolicyContentFormatOutput
}

var policyContentFormatPtrType = reflect.TypeOf((**PolicyContentFormat)(nil)).Elem()

type PolicyContentFormatPtrInput interface {
	pulumi.Input

	ToPolicyContentFormatPtrOutput() PolicyContentFormatPtrOutput
	ToPolicyContentFormatPtrOutputWithContext(context.Context) PolicyContentFormatPtrOutput
}

type policyContentFormatPtr string

func PolicyContentFormatPtr(v string) PolicyContentFormatPtrInput {
	return (*policyContentFormatPtr)(&v)
}

func (*policyContentFormatPtr) ElementType() reflect.Type {
	return policyContentFormatPtrType
}

func (in *policyContentFormatPtr) ToPolicyContentFormatPtrOutput() PolicyContentFormatPtrOutput {
	return pulumi.ToOutput(in).(PolicyContentFormatPtrOutput)
}

func (in *policyContentFormatPtr) ToPolicyContentFormatPtrOutputWithContext(ctx context.Context) PolicyContentFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolicyContentFormatPtrOutput)
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

func (PrivateEndpointServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return e.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return PrivateEndpointServiceConnectionStatus(e).ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx).ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateEndpointServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointServiceConnectionStatus) *PrivateEndpointServiceConnectionStatus {
		return &v
	}).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) Elem() PrivateEndpointServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateEndpointServiceConnectionStatus) PrivateEndpointServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointServiceConnectionStatus
		return ret
	}).(PrivateEndpointServiceConnectionStatusOutput)
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateEndpointServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateEndpointServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput
	ToPrivateEndpointServiceConnectionStatusOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusOutput
}

var privateEndpointServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()

type PrivateEndpointServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput
	ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusPtrOutput
}

type privateEndpointServiceConnectionStatusPtr string

func PrivateEndpointServiceConnectionStatusPtr(v string) PrivateEndpointServiceConnectionStatusPtrInput {
	return (*privateEndpointServiceConnectionStatusPtr)(&v)
}

func (*privateEndpointServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateEndpointServiceConnectionStatusPtrType
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

type ProductStateEnum string

const (
	ProductStateEnumNotPublished = ProductStateEnum("notPublished")
	ProductStateEnumPublished    = ProductStateEnum("published")
)

func (ProductStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductStateEnum)(nil)).Elem()
}

func (e ProductStateEnum) ToProductStateEnumOutput() ProductStateEnumOutput {
	return pulumi.ToOutput(e).(ProductStateEnumOutput)
}

func (e ProductStateEnum) ToProductStateEnumOutputWithContext(ctx context.Context) ProductStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProductStateEnumOutput)
}

func (e ProductStateEnum) ToProductStateEnumPtrOutput() ProductStateEnumPtrOutput {
	return e.ToProductStateEnumPtrOutputWithContext(context.Background())
}

func (e ProductStateEnum) ToProductStateEnumPtrOutputWithContext(ctx context.Context) ProductStateEnumPtrOutput {
	return ProductStateEnum(e).ToProductStateEnumOutputWithContext(ctx).ToProductStateEnumPtrOutputWithContext(ctx)
}

func (e ProductStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProductStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProductStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProductStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProductStateEnumOutput struct{ *pulumi.OutputState }

func (ProductStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductStateEnum)(nil)).Elem()
}

func (o ProductStateEnumOutput) ToProductStateEnumOutput() ProductStateEnumOutput {
	return o
}

func (o ProductStateEnumOutput) ToProductStateEnumOutputWithContext(ctx context.Context) ProductStateEnumOutput {
	return o
}

func (o ProductStateEnumOutput) ToProductStateEnumPtrOutput() ProductStateEnumPtrOutput {
	return o.ToProductStateEnumPtrOutputWithContext(context.Background())
}

func (o ProductStateEnumOutput) ToProductStateEnumPtrOutputWithContext(ctx context.Context) ProductStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProductStateEnum) *ProductStateEnum {
		return &v
	}).(ProductStateEnumPtrOutput)
}

func (o ProductStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProductStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProductStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProductStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProductStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProductStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProductStateEnumPtrOutput struct{ *pulumi.OutputState }

func (ProductStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductStateEnum)(nil)).Elem()
}

func (o ProductStateEnumPtrOutput) ToProductStateEnumPtrOutput() ProductStateEnumPtrOutput {
	return o
}

func (o ProductStateEnumPtrOutput) ToProductStateEnumPtrOutputWithContext(ctx context.Context) ProductStateEnumPtrOutput {
	return o
}

func (o ProductStateEnumPtrOutput) Elem() ProductStateEnumOutput {
	return o.ApplyT(func(v *ProductStateEnum) ProductStateEnum {
		if v != nil {
			return *v
		}
		var ret ProductStateEnum
		return ret
	}).(ProductStateEnumOutput)
}

func (o ProductStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProductStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProductStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProductStateEnumInput interface {
	pulumi.Input

	ToProductStateEnumOutput() ProductStateEnumOutput
	ToProductStateEnumOutputWithContext(context.Context) ProductStateEnumOutput
}

var productStateEnumPtrType = reflect.TypeOf((**ProductStateEnum)(nil)).Elem()

type ProductStateEnumPtrInput interface {
	pulumi.Input

	ToProductStateEnumPtrOutput() ProductStateEnumPtrOutput
	ToProductStateEnumPtrOutputWithContext(context.Context) ProductStateEnumPtrOutput
}

type productStateEnumPtr string

func ProductStateEnumPtr(v string) ProductStateEnumPtrInput {
	return (*productStateEnumPtr)(&v)
}

func (*productStateEnumPtr) ElementType() reflect.Type {
	return productStateEnumPtrType
}

func (in *productStateEnumPtr) ToProductStateEnumPtrOutput() ProductStateEnumPtrOutput {
	return pulumi.ToOutput(in).(ProductStateEnumPtrOutput)
}

func (in *productStateEnumPtr) ToProductStateEnumPtrOutputWithContext(ctx context.Context) ProductStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProductStateEnumPtrOutput)
}

type Protocol string

const (
	ProtocolHttp  = Protocol("http")
	ProtocolHttps = Protocol("https")
	ProtocolWs    = Protocol("ws")
	ProtocolWss   = Protocol("wss")
)

func (Protocol) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (e Protocol) ToProtocolOutput() ProtocolOutput {
	return pulumi.ToOutput(e).(ProtocolOutput)
}

func (e Protocol) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtocolOutput)
}

func (e Protocol) ToProtocolPtrOutput() ProtocolPtrOutput {
	return e.ToProtocolPtrOutputWithContext(context.Background())
}

func (e Protocol) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return Protocol(e).ToProtocolOutputWithContext(ctx).ToProtocolPtrOutputWithContext(ctx)
}

func (e Protocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Protocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtocolOutput struct{ *pulumi.OutputState }

func (ProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (o ProtocolOutput) ToProtocolOutput() ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o.ToProtocolPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Protocol) *Protocol {
		return &v
	}).(ProtocolPtrOutput)
}

func (o ProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtocolPtrOutput struct{ *pulumi.OutputState }

func (ProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Protocol)(nil)).Elem()
}

func (o ProtocolPtrOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) Elem() ProtocolOutput {
	return o.ApplyT(func(v *Protocol) Protocol {
		if v != nil {
			return *v
		}
		var ret Protocol
		return ret
	}).(ProtocolOutput)
}

func (o ProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Protocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtocolInput interface {
	pulumi.Input

	ToProtocolOutput() ProtocolOutput
	ToProtocolOutputWithContext(context.Context) ProtocolOutput
}

var protocolPtrType = reflect.TypeOf((**Protocol)(nil)).Elem()

type ProtocolPtrInput interface {
	pulumi.Input

	ToProtocolPtrOutput() ProtocolPtrOutput
	ToProtocolPtrOutputWithContext(context.Context) ProtocolPtrOutput
}

type protocolPtr string

func ProtocolPtr(v string) ProtocolPtrInput {
	return (*protocolPtr)(&v)
}

func (*protocolPtr) ElementType() reflect.Type {
	return protocolPtrType
}

func (in *protocolPtr) ToProtocolPtrOutput() ProtocolPtrOutput {
	return pulumi.ToOutput(in).(ProtocolPtrOutput)
}

func (in *protocolPtr) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtocolPtrOutput)
}

type ProvisioningState string

const (
	ProvisioningStateCreated = ProvisioningState("created")
)

func (ProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (e ProvisioningState) ToProvisioningStateOutput() ProvisioningStateOutput {
	return pulumi.ToOutput(e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return e.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return ProvisioningState(e).ToProvisioningStateOutputWithContext(ctx).ToProvisioningStatePtrOutputWithContext(ctx)
}

func (e ProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningStateOutput struct{ *pulumi.OutputState }

func (ProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStateOutput) ToProvisioningStateOutput() ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningState) *ProvisioningState {
		return &v
	}).(ProvisioningStatePtrOutput)
}

func (o ProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) Elem() ProvisioningStateOutput {
	return o.ApplyT(func(v *ProvisioningState) ProvisioningState {
		if v != nil {
			return *v
		}
		var ret ProvisioningState
		return ret
	}).(ProvisioningStateOutput)
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProvisioningStateInput interface {
	pulumi.Input

	ToProvisioningStateOutput() ProvisioningStateOutput
	ToProvisioningStateOutputWithContext(context.Context) ProvisioningStateOutput
}

var provisioningStatePtrType = reflect.TypeOf((**ProvisioningState)(nil)).Elem()

type ProvisioningStatePtrInput interface {
	pulumi.Input

	ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput
	ToProvisioningStatePtrOutputWithContext(context.Context) ProvisioningStatePtrOutput
}

type provisioningStatePtr string

func ProvisioningStatePtr(v string) ProvisioningStatePtrInput {
	return (*provisioningStatePtr)(&v)
}

func (*provisioningStatePtr) ElementType() reflect.Type {
	return provisioningStatePtrType
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ProvisioningStatePtrOutput)
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningStatePtrOutput)
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

func (PublicNetworkAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return e.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return PublicNetworkAccess(e).ToPublicNetworkAccessOutputWithContext(ctx).ToPublicNetworkAccessPtrOutputWithContext(ctx)
}

func (e PublicNetworkAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccess) *PublicNetworkAccess {
		return &v
	}).(PublicNetworkAccessPtrOutput)
}

func (o PublicNetworkAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessPtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) Elem() PublicNetworkAccessOutput {
	return o.ApplyT(func(v *PublicNetworkAccess) PublicNetworkAccess {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccess
		return ret
	}).(PublicNetworkAccessOutput)
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicNetworkAccessInput interface {
	pulumi.Input

	ToPublicNetworkAccessOutput() PublicNetworkAccessOutput
	ToPublicNetworkAccessOutputWithContext(context.Context) PublicNetworkAccessOutput
}

var publicNetworkAccessPtrType = reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()

type PublicNetworkAccessPtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput
	ToPublicNetworkAccessPtrOutputWithContext(context.Context) PublicNetworkAccessPtrOutput
}

type publicNetworkAccessPtr string

func PublicNetworkAccessPtr(v string) PublicNetworkAccessPtrInput {
	return (*publicNetworkAccessPtr)(&v)
}

func (*publicNetworkAccessPtr) ElementType() reflect.Type {
	return publicNetworkAccessPtrType
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessPtrOutput)
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessPtrOutput)
}

type SamplingType string

const (
	// Fixed-rate sampling.
	SamplingTypeFixed = SamplingType("fixed")
)

func (SamplingType) ElementType() reflect.Type {
	return reflect.TypeOf((*SamplingType)(nil)).Elem()
}

func (e SamplingType) ToSamplingTypeOutput() SamplingTypeOutput {
	return pulumi.ToOutput(e).(SamplingTypeOutput)
}

func (e SamplingType) ToSamplingTypeOutputWithContext(ctx context.Context) SamplingTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SamplingTypeOutput)
}

func (e SamplingType) ToSamplingTypePtrOutput() SamplingTypePtrOutput {
	return e.ToSamplingTypePtrOutputWithContext(context.Background())
}

func (e SamplingType) ToSamplingTypePtrOutputWithContext(ctx context.Context) SamplingTypePtrOutput {
	return SamplingType(e).ToSamplingTypeOutputWithContext(ctx).ToSamplingTypePtrOutputWithContext(ctx)
}

func (e SamplingType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SamplingType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SamplingType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SamplingType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SamplingTypeOutput struct{ *pulumi.OutputState }

func (SamplingTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SamplingType)(nil)).Elem()
}

func (o SamplingTypeOutput) ToSamplingTypeOutput() SamplingTypeOutput {
	return o
}

func (o SamplingTypeOutput) ToSamplingTypeOutputWithContext(ctx context.Context) SamplingTypeOutput {
	return o
}

func (o SamplingTypeOutput) ToSamplingTypePtrOutput() SamplingTypePtrOutput {
	return o.ToSamplingTypePtrOutputWithContext(context.Background())
}

func (o SamplingTypeOutput) ToSamplingTypePtrOutputWithContext(ctx context.Context) SamplingTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SamplingType) *SamplingType {
		return &v
	}).(SamplingTypePtrOutput)
}

func (o SamplingTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SamplingTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SamplingType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SamplingTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SamplingTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SamplingType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SamplingTypePtrOutput struct{ *pulumi.OutputState }

func (SamplingTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamplingType)(nil)).Elem()
}

func (o SamplingTypePtrOutput) ToSamplingTypePtrOutput() SamplingTypePtrOutput {
	return o
}

func (o SamplingTypePtrOutput) ToSamplingTypePtrOutputWithContext(ctx context.Context) SamplingTypePtrOutput {
	return o
}

func (o SamplingTypePtrOutput) Elem() SamplingTypeOutput {
	return o.ApplyT(func(v *SamplingType) SamplingType {
		if v != nil {
			return *v
		}
		var ret SamplingType
		return ret
	}).(SamplingTypeOutput)
}

func (o SamplingTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SamplingTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SamplingType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SamplingTypeInput interface {
	pulumi.Input

	ToSamplingTypeOutput() SamplingTypeOutput
	ToSamplingTypeOutputWithContext(context.Context) SamplingTypeOutput
}

var samplingTypePtrType = reflect.TypeOf((**SamplingType)(nil)).Elem()

type SamplingTypePtrInput interface {
	pulumi.Input

	ToSamplingTypePtrOutput() SamplingTypePtrOutput
	ToSamplingTypePtrOutputWithContext(context.Context) SamplingTypePtrOutput
}

type samplingTypePtr string

func SamplingTypePtr(v string) SamplingTypePtrInput {
	return (*samplingTypePtr)(&v)
}

func (*samplingTypePtr) ElementType() reflect.Type {
	return samplingTypePtrType
}

func (in *samplingTypePtr) ToSamplingTypePtrOutput() SamplingTypePtrOutput {
	return pulumi.ToOutput(in).(SamplingTypePtrOutput)
}

func (in *samplingTypePtr) ToSamplingTypePtrOutputWithContext(ctx context.Context) SamplingTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SamplingTypePtrOutput)
}

type SkuType string

const (
	// Developer SKU of Api Management.
	SkuTypeDeveloper = SkuType("Developer")
	// Standard SKU of Api Management.
	SkuTypeStandard = SkuType("Standard")
	// Premium SKU of Api Management.
	SkuTypePremium = SkuType("Premium")
	// Basic SKU of Api Management.
	SkuTypeBasic = SkuType("Basic")
	// Consumption SKU of Api Management.
	SkuTypeConsumption = SkuType("Consumption")
	// Isolated SKU of Api Management.
	SkuTypeIsolated = SkuType("Isolated")
)

func (SkuType) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuType)(nil)).Elem()
}

func (e SkuType) ToSkuTypeOutput() SkuTypeOutput {
	return pulumi.ToOutput(e).(SkuTypeOutput)
}

func (e SkuType) ToSkuTypeOutputWithContext(ctx context.Context) SkuTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuTypeOutput)
}

func (e SkuType) ToSkuTypePtrOutput() SkuTypePtrOutput {
	return e.ToSkuTypePtrOutputWithContext(context.Background())
}

func (e SkuType) ToSkuTypePtrOutputWithContext(ctx context.Context) SkuTypePtrOutput {
	return SkuType(e).ToSkuTypeOutputWithContext(ctx).ToSkuTypePtrOutputWithContext(ctx)
}

func (e SkuType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuTypeOutput struct{ *pulumi.OutputState }

func (SkuTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuType)(nil)).Elem()
}

func (o SkuTypeOutput) ToSkuTypeOutput() SkuTypeOutput {
	return o
}

func (o SkuTypeOutput) ToSkuTypeOutputWithContext(ctx context.Context) SkuTypeOutput {
	return o
}

func (o SkuTypeOutput) ToSkuTypePtrOutput() SkuTypePtrOutput {
	return o.ToSkuTypePtrOutputWithContext(context.Background())
}

func (o SkuTypeOutput) ToSkuTypePtrOutputWithContext(ctx context.Context) SkuTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuType) *SkuType {
		return &v
	}).(SkuTypePtrOutput)
}

func (o SkuTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuTypePtrOutput struct{ *pulumi.OutputState }

func (SkuTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuType)(nil)).Elem()
}

func (o SkuTypePtrOutput) ToSkuTypePtrOutput() SkuTypePtrOutput {
	return o
}

func (o SkuTypePtrOutput) ToSkuTypePtrOutputWithContext(ctx context.Context) SkuTypePtrOutput {
	return o
}

func (o SkuTypePtrOutput) Elem() SkuTypeOutput {
	return o.ApplyT(func(v *SkuType) SkuType {
		if v != nil {
			return *v
		}
		var ret SkuType
		return ret
	}).(SkuTypeOutput)
}

func (o SkuTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuTypeInput interface {
	pulumi.Input

	ToSkuTypeOutput() SkuTypeOutput
	ToSkuTypeOutputWithContext(context.Context) SkuTypeOutput
}

var skuTypePtrType = reflect.TypeOf((**SkuType)(nil)).Elem()

type SkuTypePtrInput interface {
	pulumi.Input

	ToSkuTypePtrOutput() SkuTypePtrOutput
	ToSkuTypePtrOutputWithContext(context.Context) SkuTypePtrOutput
}

type skuTypePtr string

func SkuTypePtr(v string) SkuTypePtrInput {
	return (*skuTypePtr)(&v)
}

func (*skuTypePtr) ElementType() reflect.Type {
	return skuTypePtrType
}

func (in *skuTypePtr) ToSkuTypePtrOutput() SkuTypePtrOutput {
	return pulumi.ToOutput(in).(SkuTypePtrOutput)
}

func (in *skuTypePtr) ToSkuTypePtrOutputWithContext(ctx context.Context) SkuTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuTypePtrOutput)
}

type SoapApiType string

const (
	// Imports a SOAP API having a RESTful front end.
	SoapApiTypeSoapToRest = SoapApiType("http")
	// Imports the SOAP API having a SOAP front end.
	SoapApiTypeSoapPassThrough = SoapApiType("soap")
	// Imports the API having a Websocket front end.
	SoapApiTypeWebSocket = SoapApiType("websocket")
	// Imports the API having a GraphQL front end.
	SoapApiTypeGraphQL = SoapApiType("graphql")
)

func (SoapApiType) ElementType() reflect.Type {
	return reflect.TypeOf((*SoapApiType)(nil)).Elem()
}

func (e SoapApiType) ToSoapApiTypeOutput() SoapApiTypeOutput {
	return pulumi.ToOutput(e).(SoapApiTypeOutput)
}

func (e SoapApiType) ToSoapApiTypeOutputWithContext(ctx context.Context) SoapApiTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SoapApiTypeOutput)
}

func (e SoapApiType) ToSoapApiTypePtrOutput() SoapApiTypePtrOutput {
	return e.ToSoapApiTypePtrOutputWithContext(context.Background())
}

func (e SoapApiType) ToSoapApiTypePtrOutputWithContext(ctx context.Context) SoapApiTypePtrOutput {
	return SoapApiType(e).ToSoapApiTypeOutputWithContext(ctx).ToSoapApiTypePtrOutputWithContext(ctx)
}

func (e SoapApiType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SoapApiType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SoapApiType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SoapApiType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SoapApiTypeOutput struct{ *pulumi.OutputState }

func (SoapApiTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoapApiType)(nil)).Elem()
}

func (o SoapApiTypeOutput) ToSoapApiTypeOutput() SoapApiTypeOutput {
	return o
}

func (o SoapApiTypeOutput) ToSoapApiTypeOutputWithContext(ctx context.Context) SoapApiTypeOutput {
	return o
}

func (o SoapApiTypeOutput) ToSoapApiTypePtrOutput() SoapApiTypePtrOutput {
	return o.ToSoapApiTypePtrOutputWithContext(context.Background())
}

func (o SoapApiTypeOutput) ToSoapApiTypePtrOutputWithContext(ctx context.Context) SoapApiTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SoapApiType) *SoapApiType {
		return &v
	}).(SoapApiTypePtrOutput)
}

func (o SoapApiTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SoapApiTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SoapApiType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SoapApiTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SoapApiTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SoapApiType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SoapApiTypePtrOutput struct{ *pulumi.OutputState }

func (SoapApiTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoapApiType)(nil)).Elem()
}

func (o SoapApiTypePtrOutput) ToSoapApiTypePtrOutput() SoapApiTypePtrOutput {
	return o
}

func (o SoapApiTypePtrOutput) ToSoapApiTypePtrOutputWithContext(ctx context.Context) SoapApiTypePtrOutput {
	return o
}

func (o SoapApiTypePtrOutput) Elem() SoapApiTypeOutput {
	return o.ApplyT(func(v *SoapApiType) SoapApiType {
		if v != nil {
			return *v
		}
		var ret SoapApiType
		return ret
	}).(SoapApiTypeOutput)
}

func (o SoapApiTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SoapApiTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SoapApiType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SoapApiTypeInput interface {
	pulumi.Input

	ToSoapApiTypeOutput() SoapApiTypeOutput
	ToSoapApiTypeOutputWithContext(context.Context) SoapApiTypeOutput
}

var soapApiTypePtrType = reflect.TypeOf((**SoapApiType)(nil)).Elem()

type SoapApiTypePtrInput interface {
	pulumi.Input

	ToSoapApiTypePtrOutput() SoapApiTypePtrOutput
	ToSoapApiTypePtrOutputWithContext(context.Context) SoapApiTypePtrOutput
}

type soapApiTypePtr string

func SoapApiTypePtr(v string) SoapApiTypePtrInput {
	return (*soapApiTypePtr)(&v)
}

func (*soapApiTypePtr) ElementType() reflect.Type {
	return soapApiTypePtrType
}

func (in *soapApiTypePtr) ToSoapApiTypePtrOutput() SoapApiTypePtrOutput {
	return pulumi.ToOutput(in).(SoapApiTypePtrOutput)
}

func (in *soapApiTypePtr) ToSoapApiTypePtrOutputWithContext(ctx context.Context) SoapApiTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SoapApiTypePtrOutput)
}

type State string

const (
	// The issue is proposed.
	StateProposed = State("proposed")
	// The issue is opened.
	StateOpen = State("open")
	// The issue was removed.
	StateRemoved = State("removed")
	// The issue is now resolved.
	StateResolved = State("resolved")
	// The issue was closed.
	StateClosed = State("closed")
)

func (State) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (e State) ToStateOutput() StateOutput {
	return pulumi.ToOutput(e).(StateOutput)
}

func (e State) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StateOutput)
}

func (e State) ToStatePtrOutput() StatePtrOutput {
	return e.ToStatePtrOutputWithContext(context.Background())
}

func (e State) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return State(e).ToStateOutputWithContext(ctx).ToStatePtrOutputWithContext(ctx)
}

func (e State) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e State) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StateOutput struct{ *pulumi.OutputState }

func (StateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (o StateOutput) ToStateOutput() StateOutput {
	return o
}

func (o StateOutput) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return o
}

func (o StateOutput) ToStatePtrOutput() StatePtrOutput {
	return o.ToStatePtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v State) *State {
		return &v
	}).(StatePtrOutput)
}

func (o StateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatePtrOutput struct{ *pulumi.OutputState }

func (StatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**State)(nil)).Elem()
}

func (o StatePtrOutput) ToStatePtrOutput() StatePtrOutput {
	return o
}

func (o StatePtrOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o
}

func (o StatePtrOutput) Elem() StateOutput {
	return o.ApplyT(func(v *State) State {
		if v != nil {
			return *v
		}
		var ret State
		return ret
	}).(StateOutput)
}

func (o StatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *State) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StateInput interface {
	pulumi.Input

	ToStateOutput() StateOutput
	ToStateOutputWithContext(context.Context) StateOutput
}

var statePtrType = reflect.TypeOf((**State)(nil)).Elem()

type StatePtrInput interface {
	pulumi.Input

	ToStatePtrOutput() StatePtrOutput
	ToStatePtrOutputWithContext(context.Context) StatePtrOutput
}

type statePtr string

func StatePtr(v string) StatePtrInput {
	return (*statePtr)(&v)
}

func (*statePtr) ElementType() reflect.Type {
	return statePtrType
}

func (in *statePtr) ToStatePtrOutput() StatePtrOutput {
	return pulumi.ToOutput(in).(StatePtrOutput)
}

func (in *statePtr) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatePtrOutput)
}

type SubscriptionStateEnum string

const (
	SubscriptionStateEnumSuspended = SubscriptionStateEnum("suspended")
	SubscriptionStateEnumActive    = SubscriptionStateEnum("active")
	SubscriptionStateEnumExpired   = SubscriptionStateEnum("expired")
	SubscriptionStateEnumSubmitted = SubscriptionStateEnum("submitted")
	SubscriptionStateEnumRejected  = SubscriptionStateEnum("rejected")
	SubscriptionStateEnumCancelled = SubscriptionStateEnum("cancelled")
)

func (SubscriptionStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionStateEnum)(nil)).Elem()
}

func (e SubscriptionStateEnum) ToSubscriptionStateEnumOutput() SubscriptionStateEnumOutput {
	return pulumi.ToOutput(e).(SubscriptionStateEnumOutput)
}

func (e SubscriptionStateEnum) ToSubscriptionStateEnumOutputWithContext(ctx context.Context) SubscriptionStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SubscriptionStateEnumOutput)
}

func (e SubscriptionStateEnum) ToSubscriptionStateEnumPtrOutput() SubscriptionStateEnumPtrOutput {
	return e.ToSubscriptionStateEnumPtrOutputWithContext(context.Background())
}

func (e SubscriptionStateEnum) ToSubscriptionStateEnumPtrOutputWithContext(ctx context.Context) SubscriptionStateEnumPtrOutput {
	return SubscriptionStateEnum(e).ToSubscriptionStateEnumOutputWithContext(ctx).ToSubscriptionStateEnumPtrOutputWithContext(ctx)
}

func (e SubscriptionStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubscriptionStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubscriptionStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SubscriptionStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SubscriptionStateEnumOutput struct{ *pulumi.OutputState }

func (SubscriptionStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionStateEnum)(nil)).Elem()
}

func (o SubscriptionStateEnumOutput) ToSubscriptionStateEnumOutput() SubscriptionStateEnumOutput {
	return o
}

func (o SubscriptionStateEnumOutput) ToSubscriptionStateEnumOutputWithContext(ctx context.Context) SubscriptionStateEnumOutput {
	return o
}

func (o SubscriptionStateEnumOutput) ToSubscriptionStateEnumPtrOutput() SubscriptionStateEnumPtrOutput {
	return o.ToSubscriptionStateEnumPtrOutputWithContext(context.Background())
}

func (o SubscriptionStateEnumOutput) ToSubscriptionStateEnumPtrOutputWithContext(ctx context.Context) SubscriptionStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionStateEnum) *SubscriptionStateEnum {
		return &v
	}).(SubscriptionStateEnumPtrOutput)
}

func (o SubscriptionStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SubscriptionStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriptionStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SubscriptionStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriptionStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriptionStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SubscriptionStateEnumPtrOutput struct{ *pulumi.OutputState }

func (SubscriptionStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionStateEnum)(nil)).Elem()
}

func (o SubscriptionStateEnumPtrOutput) ToSubscriptionStateEnumPtrOutput() SubscriptionStateEnumPtrOutput {
	return o
}

func (o SubscriptionStateEnumPtrOutput) ToSubscriptionStateEnumPtrOutputWithContext(ctx context.Context) SubscriptionStateEnumPtrOutput {
	return o
}

func (o SubscriptionStateEnumPtrOutput) Elem() SubscriptionStateEnumOutput {
	return o.ApplyT(func(v *SubscriptionStateEnum) SubscriptionStateEnum {
		if v != nil {
			return *v
		}
		var ret SubscriptionStateEnum
		return ret
	}).(SubscriptionStateEnumOutput)
}

func (o SubscriptionStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriptionStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SubscriptionStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SubscriptionStateEnumInput interface {
	pulumi.Input

	ToSubscriptionStateEnumOutput() SubscriptionStateEnumOutput
	ToSubscriptionStateEnumOutputWithContext(context.Context) SubscriptionStateEnumOutput
}

var subscriptionStateEnumPtrType = reflect.TypeOf((**SubscriptionStateEnum)(nil)).Elem()

type SubscriptionStateEnumPtrInput interface {
	pulumi.Input

	ToSubscriptionStateEnumPtrOutput() SubscriptionStateEnumPtrOutput
	ToSubscriptionStateEnumPtrOutputWithContext(context.Context) SubscriptionStateEnumPtrOutput
}

type subscriptionStateEnumPtr string

func SubscriptionStateEnumPtr(v string) SubscriptionStateEnumPtrInput {
	return (*subscriptionStateEnumPtr)(&v)
}

func (*subscriptionStateEnumPtr) ElementType() reflect.Type {
	return subscriptionStateEnumPtrType
}

func (in *subscriptionStateEnumPtr) ToSubscriptionStateEnumPtrOutput() SubscriptionStateEnumPtrOutput {
	return pulumi.ToOutput(in).(SubscriptionStateEnumPtrOutput)
}

func (in *subscriptionStateEnumPtr) ToSubscriptionStateEnumPtrOutputWithContext(ctx context.Context) SubscriptionStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SubscriptionStateEnumPtrOutput)
}

type UserStateEnum string

const (
	// User state is active.
	UserStateEnumActive = UserStateEnum("active")
	// User is blocked. Blocked users cannot authenticate at developer portal or call API.
	UserStateEnumBlocked = UserStateEnum("blocked")
	// User account is pending. Requires identity confirmation before it can be made active.
	UserStateEnumPending = UserStateEnum("pending")
	// User account is closed. All identities and related entities are removed.
	UserStateEnumDeleted = UserStateEnum("deleted")
)

func (UserStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*UserStateEnum)(nil)).Elem()
}

func (e UserStateEnum) ToUserStateEnumOutput() UserStateEnumOutput {
	return pulumi.ToOutput(e).(UserStateEnumOutput)
}

func (e UserStateEnum) ToUserStateEnumOutputWithContext(ctx context.Context) UserStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UserStateEnumOutput)
}

func (e UserStateEnum) ToUserStateEnumPtrOutput() UserStateEnumPtrOutput {
	return e.ToUserStateEnumPtrOutputWithContext(context.Background())
}

func (e UserStateEnum) ToUserStateEnumPtrOutputWithContext(ctx context.Context) UserStateEnumPtrOutput {
	return UserStateEnum(e).ToUserStateEnumOutputWithContext(ctx).ToUserStateEnumPtrOutputWithContext(ctx)
}

func (e UserStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UserStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UserStateEnumOutput struct{ *pulumi.OutputState }

func (UserStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserStateEnum)(nil)).Elem()
}

func (o UserStateEnumOutput) ToUserStateEnumOutput() UserStateEnumOutput {
	return o
}

func (o UserStateEnumOutput) ToUserStateEnumOutputWithContext(ctx context.Context) UserStateEnumOutput {
	return o
}

func (o UserStateEnumOutput) ToUserStateEnumPtrOutput() UserStateEnumPtrOutput {
	return o.ToUserStateEnumPtrOutputWithContext(context.Background())
}

func (o UserStateEnumOutput) ToUserStateEnumPtrOutputWithContext(ctx context.Context) UserStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserStateEnum) *UserStateEnum {
		return &v
	}).(UserStateEnumPtrOutput)
}

func (o UserStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UserStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UserStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UserStateEnumPtrOutput struct{ *pulumi.OutputState }

func (UserStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserStateEnum)(nil)).Elem()
}

func (o UserStateEnumPtrOutput) ToUserStateEnumPtrOutput() UserStateEnumPtrOutput {
	return o
}

func (o UserStateEnumPtrOutput) ToUserStateEnumPtrOutputWithContext(ctx context.Context) UserStateEnumPtrOutput {
	return o
}

func (o UserStateEnumPtrOutput) Elem() UserStateEnumOutput {
	return o.ApplyT(func(v *UserStateEnum) UserStateEnum {
		if v != nil {
			return *v
		}
		var ret UserStateEnum
		return ret
	}).(UserStateEnumOutput)
}

func (o UserStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UserStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UserStateEnumInput interface {
	pulumi.Input

	ToUserStateEnumOutput() UserStateEnumOutput
	ToUserStateEnumOutputWithContext(context.Context) UserStateEnumOutput
}

var userStateEnumPtrType = reflect.TypeOf((**UserStateEnum)(nil)).Elem()

type UserStateEnumPtrInput interface {
	pulumi.Input

	ToUserStateEnumPtrOutput() UserStateEnumPtrOutput
	ToUserStateEnumPtrOutputWithContext(context.Context) UserStateEnumPtrOutput
}

type userStateEnumPtr string

func UserStateEnumPtr(v string) UserStateEnumPtrInput {
	return (*userStateEnumPtr)(&v)
}

func (*userStateEnumPtr) ElementType() reflect.Type {
	return userStateEnumPtrType
}

func (in *userStateEnumPtr) ToUserStateEnumPtrOutput() UserStateEnumPtrOutput {
	return pulumi.ToOutput(in).(UserStateEnumPtrOutput)
}

func (in *userStateEnumPtr) ToUserStateEnumPtrOutputWithContext(ctx context.Context) UserStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UserStateEnumPtrOutput)
}

type Verbosity string

const (
	// All the traces emitted by trace policies will be sent to the logger attached to this diagnostic instance.
	VerbosityVerbose = Verbosity("verbose")
	// Traces with 'severity' set to 'information' and 'error' will be sent to the logger attached to this diagnostic instance.
	VerbosityInformation = Verbosity("information")
	// Only traces with 'severity' set to 'error' will be sent to the logger attached to this diagnostic instance.
	VerbosityError = Verbosity("error")
)

func (Verbosity) ElementType() reflect.Type {
	return reflect.TypeOf((*Verbosity)(nil)).Elem()
}

func (e Verbosity) ToVerbosityOutput() VerbosityOutput {
	return pulumi.ToOutput(e).(VerbosityOutput)
}

func (e Verbosity) ToVerbosityOutputWithContext(ctx context.Context) VerbosityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VerbosityOutput)
}

func (e Verbosity) ToVerbosityPtrOutput() VerbosityPtrOutput {
	return e.ToVerbosityPtrOutputWithContext(context.Background())
}

func (e Verbosity) ToVerbosityPtrOutputWithContext(ctx context.Context) VerbosityPtrOutput {
	return Verbosity(e).ToVerbosityOutputWithContext(ctx).ToVerbosityPtrOutputWithContext(ctx)
}

func (e Verbosity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Verbosity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Verbosity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Verbosity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VerbosityOutput struct{ *pulumi.OutputState }

func (VerbosityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Verbosity)(nil)).Elem()
}

func (o VerbosityOutput) ToVerbosityOutput() VerbosityOutput {
	return o
}

func (o VerbosityOutput) ToVerbosityOutputWithContext(ctx context.Context) VerbosityOutput {
	return o
}

func (o VerbosityOutput) ToVerbosityPtrOutput() VerbosityPtrOutput {
	return o.ToVerbosityPtrOutputWithContext(context.Background())
}

func (o VerbosityOutput) ToVerbosityPtrOutputWithContext(ctx context.Context) VerbosityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Verbosity) *Verbosity {
		return &v
	}).(VerbosityPtrOutput)
}

func (o VerbosityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VerbosityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Verbosity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VerbosityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VerbosityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Verbosity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VerbosityPtrOutput struct{ *pulumi.OutputState }

func (VerbosityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Verbosity)(nil)).Elem()
}

func (o VerbosityPtrOutput) ToVerbosityPtrOutput() VerbosityPtrOutput {
	return o
}

func (o VerbosityPtrOutput) ToVerbosityPtrOutputWithContext(ctx context.Context) VerbosityPtrOutput {
	return o
}

func (o VerbosityPtrOutput) Elem() VerbosityOutput {
	return o.ApplyT(func(v *Verbosity) Verbosity {
		if v != nil {
			return *v
		}
		var ret Verbosity
		return ret
	}).(VerbosityOutput)
}

func (o VerbosityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VerbosityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Verbosity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VerbosityInput interface {
	pulumi.Input

	ToVerbosityOutput() VerbosityOutput
	ToVerbosityOutputWithContext(context.Context) VerbosityOutput
}

var verbosityPtrType = reflect.TypeOf((**Verbosity)(nil)).Elem()

type VerbosityPtrInput interface {
	pulumi.Input

	ToVerbosityPtrOutput() VerbosityPtrOutput
	ToVerbosityPtrOutputWithContext(context.Context) VerbosityPtrOutput
}

type verbosityPtr string

func VerbosityPtr(v string) VerbosityPtrInput {
	return (*verbosityPtr)(&v)
}

func (*verbosityPtr) ElementType() reflect.Type {
	return verbosityPtrType
}

func (in *verbosityPtr) ToVerbosityPtrOutput() VerbosityPtrOutput {
	return pulumi.ToOutput(in).(VerbosityPtrOutput)
}

func (in *verbosityPtr) ToVerbosityPtrOutputWithContext(ctx context.Context) VerbosityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VerbosityPtrOutput)
}

type VersioningScheme string

const (
	// The API Version is passed in a path segment.
	VersioningSchemeSegment = VersioningScheme("Segment")
	// The API Version is passed in a query parameter.
	VersioningSchemeQuery = VersioningScheme("Query")
	// The API Version is passed in a HTTP header.
	VersioningSchemeHeader = VersioningScheme("Header")
)

func (VersioningScheme) ElementType() reflect.Type {
	return reflect.TypeOf((*VersioningScheme)(nil)).Elem()
}

func (e VersioningScheme) ToVersioningSchemeOutput() VersioningSchemeOutput {
	return pulumi.ToOutput(e).(VersioningSchemeOutput)
}

func (e VersioningScheme) ToVersioningSchemeOutputWithContext(ctx context.Context) VersioningSchemeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VersioningSchemeOutput)
}

func (e VersioningScheme) ToVersioningSchemePtrOutput() VersioningSchemePtrOutput {
	return e.ToVersioningSchemePtrOutputWithContext(context.Background())
}

func (e VersioningScheme) ToVersioningSchemePtrOutputWithContext(ctx context.Context) VersioningSchemePtrOutput {
	return VersioningScheme(e).ToVersioningSchemeOutputWithContext(ctx).ToVersioningSchemePtrOutputWithContext(ctx)
}

func (e VersioningScheme) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VersioningScheme) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VersioningScheme) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VersioningScheme) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VersioningSchemeOutput struct{ *pulumi.OutputState }

func (VersioningSchemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VersioningScheme)(nil)).Elem()
}

func (o VersioningSchemeOutput) ToVersioningSchemeOutput() VersioningSchemeOutput {
	return o
}

func (o VersioningSchemeOutput) ToVersioningSchemeOutputWithContext(ctx context.Context) VersioningSchemeOutput {
	return o
}

func (o VersioningSchemeOutput) ToVersioningSchemePtrOutput() VersioningSchemePtrOutput {
	return o.ToVersioningSchemePtrOutputWithContext(context.Background())
}

func (o VersioningSchemeOutput) ToVersioningSchemePtrOutputWithContext(ctx context.Context) VersioningSchemePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VersioningScheme) *VersioningScheme {
		return &v
	}).(VersioningSchemePtrOutput)
}

func (o VersioningSchemeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VersioningSchemeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VersioningScheme) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VersioningSchemeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VersioningSchemeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VersioningScheme) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VersioningSchemePtrOutput struct{ *pulumi.OutputState }

func (VersioningSchemePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VersioningScheme)(nil)).Elem()
}

func (o VersioningSchemePtrOutput) ToVersioningSchemePtrOutput() VersioningSchemePtrOutput {
	return o
}

func (o VersioningSchemePtrOutput) ToVersioningSchemePtrOutputWithContext(ctx context.Context) VersioningSchemePtrOutput {
	return o
}

func (o VersioningSchemePtrOutput) Elem() VersioningSchemeOutput {
	return o.ApplyT(func(v *VersioningScheme) VersioningScheme {
		if v != nil {
			return *v
		}
		var ret VersioningScheme
		return ret
	}).(VersioningSchemeOutput)
}

func (o VersioningSchemePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VersioningSchemePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VersioningScheme) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VersioningSchemeInput interface {
	pulumi.Input

	ToVersioningSchemeOutput() VersioningSchemeOutput
	ToVersioningSchemeOutputWithContext(context.Context) VersioningSchemeOutput
}

var versioningSchemePtrType = reflect.TypeOf((**VersioningScheme)(nil)).Elem()

type VersioningSchemePtrInput interface {
	pulumi.Input

	ToVersioningSchemePtrOutput() VersioningSchemePtrOutput
	ToVersioningSchemePtrOutputWithContext(context.Context) VersioningSchemePtrOutput
}

type versioningSchemePtr string

func VersioningSchemePtr(v string) VersioningSchemePtrInput {
	return (*versioningSchemePtr)(&v)
}

func (*versioningSchemePtr) ElementType() reflect.Type {
	return versioningSchemePtrType
}

func (in *versioningSchemePtr) ToVersioningSchemePtrOutput() VersioningSchemePtrOutput {
	return pulumi.ToOutput(in).(VersioningSchemePtrOutput)
}

func (in *versioningSchemePtr) ToVersioningSchemePtrOutputWithContext(ctx context.Context) VersioningSchemePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VersioningSchemePtrOutput)
}

type VirtualNetworkType string

const (
	// The service is not part of any Virtual Network.
	VirtualNetworkTypeNone = VirtualNetworkType("None")
	// The service is part of Virtual Network and it is accessible from Internet.
	VirtualNetworkTypeExternal = VirtualNetworkType("External")
	// The service is part of Virtual Network and it is only accessible from within the virtual network.
	VirtualNetworkTypeInternal = VirtualNetworkType("Internal")
)

func (VirtualNetworkType) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkType)(nil)).Elem()
}

func (e VirtualNetworkType) ToVirtualNetworkTypeOutput() VirtualNetworkTypeOutput {
	return pulumi.ToOutput(e).(VirtualNetworkTypeOutput)
}

func (e VirtualNetworkType) ToVirtualNetworkTypeOutputWithContext(ctx context.Context) VirtualNetworkTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkTypeOutput)
}

func (e VirtualNetworkType) ToVirtualNetworkTypePtrOutput() VirtualNetworkTypePtrOutput {
	return e.ToVirtualNetworkTypePtrOutputWithContext(context.Background())
}

func (e VirtualNetworkType) ToVirtualNetworkTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTypePtrOutput {
	return VirtualNetworkType(e).ToVirtualNetworkTypeOutputWithContext(ctx).ToVirtualNetworkTypePtrOutputWithContext(ctx)
}

func (e VirtualNetworkType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkTypeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkType)(nil)).Elem()
}

func (o VirtualNetworkTypeOutput) ToVirtualNetworkTypeOutput() VirtualNetworkTypeOutput {
	return o
}

func (o VirtualNetworkTypeOutput) ToVirtualNetworkTypeOutputWithContext(ctx context.Context) VirtualNetworkTypeOutput {
	return o
}

func (o VirtualNetworkTypeOutput) ToVirtualNetworkTypePtrOutput() VirtualNetworkTypePtrOutput {
	return o.ToVirtualNetworkTypePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkTypeOutput) ToVirtualNetworkTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkType) *VirtualNetworkType {
		return &v
	}).(VirtualNetworkTypePtrOutput)
}

func (o VirtualNetworkTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkTypePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkType)(nil)).Elem()
}

func (o VirtualNetworkTypePtrOutput) ToVirtualNetworkTypePtrOutput() VirtualNetworkTypePtrOutput {
	return o
}

func (o VirtualNetworkTypePtrOutput) ToVirtualNetworkTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTypePtrOutput {
	return o
}

func (o VirtualNetworkTypePtrOutput) Elem() VirtualNetworkTypeOutput {
	return o.ApplyT(func(v *VirtualNetworkType) VirtualNetworkType {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkType
		return ret
	}).(VirtualNetworkTypeOutput)
}

func (o VirtualNetworkTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkTypeInput interface {
	pulumi.Input

	ToVirtualNetworkTypeOutput() VirtualNetworkTypeOutput
	ToVirtualNetworkTypeOutputWithContext(context.Context) VirtualNetworkTypeOutput
}

var virtualNetworkTypePtrType = reflect.TypeOf((**VirtualNetworkType)(nil)).Elem()

type VirtualNetworkTypePtrInput interface {
	pulumi.Input

	ToVirtualNetworkTypePtrOutput() VirtualNetworkTypePtrOutput
	ToVirtualNetworkTypePtrOutputWithContext(context.Context) VirtualNetworkTypePtrOutput
}

type virtualNetworkTypePtr string

func VirtualNetworkTypePtr(v string) VirtualNetworkTypePtrInput {
	return (*virtualNetworkTypePtr)(&v)
}

func (*virtualNetworkTypePtr) ElementType() reflect.Type {
	return virtualNetworkTypePtrType
}

func (in *virtualNetworkTypePtr) ToVirtualNetworkTypePtrOutput() VirtualNetworkTypePtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkTypePtrOutput)
}

func (in *virtualNetworkTypePtr) ToVirtualNetworkTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AlwaysLogOutput{})
	pulumi.RegisterOutputType(AlwaysLogPtrOutput{})
	pulumi.RegisterOutputType(ApiTypeOutput{})
	pulumi.RegisterOutputType(ApiTypePtrOutput{})
	pulumi.RegisterOutputType(ApimIdentityTypeOutput{})
	pulumi.RegisterOutputType(ApimIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(AppTypeOutput{})
	pulumi.RegisterOutputType(AppTypePtrOutput{})
	pulumi.RegisterOutputType(AuthorizationMethodOutput{})
	pulumi.RegisterOutputType(AuthorizationMethodPtrOutput{})
	pulumi.RegisterOutputType(AuthorizationMethodArrayOutput{})
	pulumi.RegisterOutputType(BackendProtocolOutput{})
	pulumi.RegisterOutputType(BackendProtocolPtrOutput{})
	pulumi.RegisterOutputType(BearerTokenSendingMethodOutput{})
	pulumi.RegisterOutputType(BearerTokenSendingMethodPtrOutput{})
	pulumi.RegisterOutputType(BearerTokenSendingMethodsOutput{})
	pulumi.RegisterOutputType(BearerTokenSendingMethodsPtrOutput{})
	pulumi.RegisterOutputType(CertificateSourceOutput{})
	pulumi.RegisterOutputType(CertificateSourcePtrOutput{})
	pulumi.RegisterOutputType(CertificateStatusOutput{})
	pulumi.RegisterOutputType(CertificateStatusPtrOutput{})
	pulumi.RegisterOutputType(ClientAuthenticationMethodOutput{})
	pulumi.RegisterOutputType(ClientAuthenticationMethodPtrOutput{})
	pulumi.RegisterOutputType(ConfirmationOutput{})
	pulumi.RegisterOutputType(ConfirmationPtrOutput{})
	pulumi.RegisterOutputType(ContentFormatOutput{})
	pulumi.RegisterOutputType(ContentFormatPtrOutput{})
	pulumi.RegisterOutputType(DataMaskingModeOutput{})
	pulumi.RegisterOutputType(DataMaskingModePtrOutput{})
	pulumi.RegisterOutputType(GrantTypeOutput{})
	pulumi.RegisterOutputType(GrantTypePtrOutput{})
	pulumi.RegisterOutputType(GroupTypeOutput{})
	pulumi.RegisterOutputType(GroupTypePtrOutput{})
	pulumi.RegisterOutputType(HostnameTypeOutput{})
	pulumi.RegisterOutputType(HostnameTypePtrOutput{})
	pulumi.RegisterOutputType(HttpCorrelationProtocolOutput{})
	pulumi.RegisterOutputType(HttpCorrelationProtocolPtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderTypeOutput{})
	pulumi.RegisterOutputType(IdentityProviderTypePtrOutput{})
	pulumi.RegisterOutputType(KeyTypeOutput{})
	pulumi.RegisterOutputType(KeyTypePtrOutput{})
	pulumi.RegisterOutputType(LoggerTypeOutput{})
	pulumi.RegisterOutputType(LoggerTypePtrOutput{})
	pulumi.RegisterOutputType(OperationNameFormatOutput{})
	pulumi.RegisterOutputType(OperationNameFormatPtrOutput{})
	pulumi.RegisterOutputType(PolicyContentFormatOutput{})
	pulumi.RegisterOutputType(PolicyContentFormatPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(ProductStateEnumOutput{})
	pulumi.RegisterOutputType(ProductStateEnumPtrOutput{})
	pulumi.RegisterOutputType(ProtocolOutput{})
	pulumi.RegisterOutputType(ProtocolPtrOutput{})
	pulumi.RegisterOutputType(ProvisioningStateOutput{})
	pulumi.RegisterOutputType(ProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(SamplingTypeOutput{})
	pulumi.RegisterOutputType(SamplingTypePtrOutput{})
	pulumi.RegisterOutputType(SkuTypeOutput{})
	pulumi.RegisterOutputType(SkuTypePtrOutput{})
	pulumi.RegisterOutputType(SoapApiTypeOutput{})
	pulumi.RegisterOutputType(SoapApiTypePtrOutput{})
	pulumi.RegisterOutputType(StateOutput{})
	pulumi.RegisterOutputType(StatePtrOutput{})
	pulumi.RegisterOutputType(SubscriptionStateEnumOutput{})
	pulumi.RegisterOutputType(SubscriptionStateEnumPtrOutput{})
	pulumi.RegisterOutputType(UserStateEnumOutput{})
	pulumi.RegisterOutputType(UserStateEnumPtrOutput{})
	pulumi.RegisterOutputType(VerbosityOutput{})
	pulumi.RegisterOutputType(VerbosityPtrOutput{})
	pulumi.RegisterOutputType(VersioningSchemeOutput{})
	pulumi.RegisterOutputType(VersioningSchemePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTypePtrOutput{})
}
