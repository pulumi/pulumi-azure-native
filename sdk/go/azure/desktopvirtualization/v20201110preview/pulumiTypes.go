


package v20201110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MsixPackageApplications struct {
	AppId          *string `pulumi:"appId"`
	AppUserModelID *string `pulumi:"appUserModelID"`
	Description    *string `pulumi:"description"`
	FriendlyName   *string `pulumi:"friendlyName"`
	IconImageName  *string `pulumi:"iconImageName"`
	RawIcon        *string `pulumi:"rawIcon"`
	RawPng         *string `pulumi:"rawPng"`
}





type MsixPackageApplicationsInput interface {
	pulumi.Input

	ToMsixPackageApplicationsOutput() MsixPackageApplicationsOutput
	ToMsixPackageApplicationsOutputWithContext(context.Context) MsixPackageApplicationsOutput
}

type MsixPackageApplicationsArgs struct {
	AppId          pulumi.StringPtrInput `pulumi:"appId"`
	AppUserModelID pulumi.StringPtrInput `pulumi:"appUserModelID"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	FriendlyName   pulumi.StringPtrInput `pulumi:"friendlyName"`
	IconImageName  pulumi.StringPtrInput `pulumi:"iconImageName"`
	RawIcon        pulumi.StringPtrInput `pulumi:"rawIcon"`
	RawPng         pulumi.StringPtrInput `pulumi:"rawPng"`
}

func (MsixPackageApplicationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageApplications)(nil)).Elem()
}

func (i MsixPackageApplicationsArgs) ToMsixPackageApplicationsOutput() MsixPackageApplicationsOutput {
	return i.ToMsixPackageApplicationsOutputWithContext(context.Background())
}

func (i MsixPackageApplicationsArgs) ToMsixPackageApplicationsOutputWithContext(ctx context.Context) MsixPackageApplicationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageApplicationsOutput)
}





type MsixPackageApplicationsArrayInput interface {
	pulumi.Input

	ToMsixPackageApplicationsArrayOutput() MsixPackageApplicationsArrayOutput
	ToMsixPackageApplicationsArrayOutputWithContext(context.Context) MsixPackageApplicationsArrayOutput
}

type MsixPackageApplicationsArray []MsixPackageApplicationsInput

func (MsixPackageApplicationsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageApplications)(nil)).Elem()
}

func (i MsixPackageApplicationsArray) ToMsixPackageApplicationsArrayOutput() MsixPackageApplicationsArrayOutput {
	return i.ToMsixPackageApplicationsArrayOutputWithContext(context.Background())
}

func (i MsixPackageApplicationsArray) ToMsixPackageApplicationsArrayOutputWithContext(ctx context.Context) MsixPackageApplicationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageApplicationsArrayOutput)
}

type MsixPackageApplicationsOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageApplications)(nil)).Elem()
}

func (o MsixPackageApplicationsOutput) ToMsixPackageApplicationsOutput() MsixPackageApplicationsOutput {
	return o
}

func (o MsixPackageApplicationsOutput) ToMsixPackageApplicationsOutputWithContext(ctx context.Context) MsixPackageApplicationsOutput {
	return o
}

func (o MsixPackageApplicationsOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) AppUserModelID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.AppUserModelID }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) IconImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.IconImageName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) RawIcon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.RawIcon }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) RawPng() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.RawPng }).(pulumi.StringPtrOutput)
}

type MsixPackageApplicationsArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageApplications)(nil)).Elem()
}

func (o MsixPackageApplicationsArrayOutput) ToMsixPackageApplicationsArrayOutput() MsixPackageApplicationsArrayOutput {
	return o
}

func (o MsixPackageApplicationsArrayOutput) ToMsixPackageApplicationsArrayOutputWithContext(ctx context.Context) MsixPackageApplicationsArrayOutput {
	return o
}

func (o MsixPackageApplicationsArrayOutput) Index(i pulumi.IntInput) MsixPackageApplicationsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageApplications {
		return vs[0].([]MsixPackageApplications)[vs[1].(int)]
	}).(MsixPackageApplicationsOutput)
}

type MsixPackageApplicationsResponse struct {
	AppId          *string `pulumi:"appId"`
	AppUserModelID *string `pulumi:"appUserModelID"`
	Description    *string `pulumi:"description"`
	FriendlyName   *string `pulumi:"friendlyName"`
	IconImageName  *string `pulumi:"iconImageName"`
	RawIcon        *string `pulumi:"rawIcon"`
	RawPng         *string `pulumi:"rawPng"`
}





type MsixPackageApplicationsResponseInput interface {
	pulumi.Input

	ToMsixPackageApplicationsResponseOutput() MsixPackageApplicationsResponseOutput
	ToMsixPackageApplicationsResponseOutputWithContext(context.Context) MsixPackageApplicationsResponseOutput
}

type MsixPackageApplicationsResponseArgs struct {
	AppId          pulumi.StringPtrInput `pulumi:"appId"`
	AppUserModelID pulumi.StringPtrInput `pulumi:"appUserModelID"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	FriendlyName   pulumi.StringPtrInput `pulumi:"friendlyName"`
	IconImageName  pulumi.StringPtrInput `pulumi:"iconImageName"`
	RawIcon        pulumi.StringPtrInput `pulumi:"rawIcon"`
	RawPng         pulumi.StringPtrInput `pulumi:"rawPng"`
}

func (MsixPackageApplicationsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageApplicationsResponse)(nil)).Elem()
}

func (i MsixPackageApplicationsResponseArgs) ToMsixPackageApplicationsResponseOutput() MsixPackageApplicationsResponseOutput {
	return i.ToMsixPackageApplicationsResponseOutputWithContext(context.Background())
}

func (i MsixPackageApplicationsResponseArgs) ToMsixPackageApplicationsResponseOutputWithContext(ctx context.Context) MsixPackageApplicationsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageApplicationsResponseOutput)
}





type MsixPackageApplicationsResponseArrayInput interface {
	pulumi.Input

	ToMsixPackageApplicationsResponseArrayOutput() MsixPackageApplicationsResponseArrayOutput
	ToMsixPackageApplicationsResponseArrayOutputWithContext(context.Context) MsixPackageApplicationsResponseArrayOutput
}

type MsixPackageApplicationsResponseArray []MsixPackageApplicationsResponseInput

func (MsixPackageApplicationsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageApplicationsResponse)(nil)).Elem()
}

func (i MsixPackageApplicationsResponseArray) ToMsixPackageApplicationsResponseArrayOutput() MsixPackageApplicationsResponseArrayOutput {
	return i.ToMsixPackageApplicationsResponseArrayOutputWithContext(context.Background())
}

func (i MsixPackageApplicationsResponseArray) ToMsixPackageApplicationsResponseArrayOutputWithContext(ctx context.Context) MsixPackageApplicationsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageApplicationsResponseArrayOutput)
}

type MsixPackageApplicationsResponseOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageApplicationsResponse)(nil)).Elem()
}

func (o MsixPackageApplicationsResponseOutput) ToMsixPackageApplicationsResponseOutput() MsixPackageApplicationsResponseOutput {
	return o
}

func (o MsixPackageApplicationsResponseOutput) ToMsixPackageApplicationsResponseOutputWithContext(ctx context.Context) MsixPackageApplicationsResponseOutput {
	return o
}

func (o MsixPackageApplicationsResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) AppUserModelID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.AppUserModelID }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) IconImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.IconImageName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) RawIcon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.RawIcon }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) RawPng() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.RawPng }).(pulumi.StringPtrOutput)
}

type MsixPackageApplicationsResponseArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageApplicationsResponse)(nil)).Elem()
}

func (o MsixPackageApplicationsResponseArrayOutput) ToMsixPackageApplicationsResponseArrayOutput() MsixPackageApplicationsResponseArrayOutput {
	return o
}

func (o MsixPackageApplicationsResponseArrayOutput) ToMsixPackageApplicationsResponseArrayOutputWithContext(ctx context.Context) MsixPackageApplicationsResponseArrayOutput {
	return o
}

func (o MsixPackageApplicationsResponseArrayOutput) Index(i pulumi.IntInput) MsixPackageApplicationsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageApplicationsResponse {
		return vs[0].([]MsixPackageApplicationsResponse)[vs[1].(int)]
	}).(MsixPackageApplicationsResponseOutput)
}

type MsixPackageDependencies struct {
	DependencyName *string `pulumi:"dependencyName"`
	MinVersion     *string `pulumi:"minVersion"`
	Publisher      *string `pulumi:"publisher"`
}





type MsixPackageDependenciesInput interface {
	pulumi.Input

	ToMsixPackageDependenciesOutput() MsixPackageDependenciesOutput
	ToMsixPackageDependenciesOutputWithContext(context.Context) MsixPackageDependenciesOutput
}

type MsixPackageDependenciesArgs struct {
	DependencyName pulumi.StringPtrInput `pulumi:"dependencyName"`
	MinVersion     pulumi.StringPtrInput `pulumi:"minVersion"`
	Publisher      pulumi.StringPtrInput `pulumi:"publisher"`
}

func (MsixPackageDependenciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageDependencies)(nil)).Elem()
}

func (i MsixPackageDependenciesArgs) ToMsixPackageDependenciesOutput() MsixPackageDependenciesOutput {
	return i.ToMsixPackageDependenciesOutputWithContext(context.Background())
}

func (i MsixPackageDependenciesArgs) ToMsixPackageDependenciesOutputWithContext(ctx context.Context) MsixPackageDependenciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageDependenciesOutput)
}





type MsixPackageDependenciesArrayInput interface {
	pulumi.Input

	ToMsixPackageDependenciesArrayOutput() MsixPackageDependenciesArrayOutput
	ToMsixPackageDependenciesArrayOutputWithContext(context.Context) MsixPackageDependenciesArrayOutput
}

type MsixPackageDependenciesArray []MsixPackageDependenciesInput

func (MsixPackageDependenciesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageDependencies)(nil)).Elem()
}

func (i MsixPackageDependenciesArray) ToMsixPackageDependenciesArrayOutput() MsixPackageDependenciesArrayOutput {
	return i.ToMsixPackageDependenciesArrayOutputWithContext(context.Background())
}

func (i MsixPackageDependenciesArray) ToMsixPackageDependenciesArrayOutputWithContext(ctx context.Context) MsixPackageDependenciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageDependenciesArrayOutput)
}

type MsixPackageDependenciesOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageDependencies)(nil)).Elem()
}

func (o MsixPackageDependenciesOutput) ToMsixPackageDependenciesOutput() MsixPackageDependenciesOutput {
	return o
}

func (o MsixPackageDependenciesOutput) ToMsixPackageDependenciesOutputWithContext(ctx context.Context) MsixPackageDependenciesOutput {
	return o
}

func (o MsixPackageDependenciesOutput) DependencyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependencies) *string { return v.DependencyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesOutput) MinVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependencies) *string { return v.MinVersion }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependencies) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type MsixPackageDependenciesArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageDependencies)(nil)).Elem()
}

func (o MsixPackageDependenciesArrayOutput) ToMsixPackageDependenciesArrayOutput() MsixPackageDependenciesArrayOutput {
	return o
}

func (o MsixPackageDependenciesArrayOutput) ToMsixPackageDependenciesArrayOutputWithContext(ctx context.Context) MsixPackageDependenciesArrayOutput {
	return o
}

func (o MsixPackageDependenciesArrayOutput) Index(i pulumi.IntInput) MsixPackageDependenciesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageDependencies {
		return vs[0].([]MsixPackageDependencies)[vs[1].(int)]
	}).(MsixPackageDependenciesOutput)
}

type MsixPackageDependenciesResponse struct {
	DependencyName *string `pulumi:"dependencyName"`
	MinVersion     *string `pulumi:"minVersion"`
	Publisher      *string `pulumi:"publisher"`
}





type MsixPackageDependenciesResponseInput interface {
	pulumi.Input

	ToMsixPackageDependenciesResponseOutput() MsixPackageDependenciesResponseOutput
	ToMsixPackageDependenciesResponseOutputWithContext(context.Context) MsixPackageDependenciesResponseOutput
}

type MsixPackageDependenciesResponseArgs struct {
	DependencyName pulumi.StringPtrInput `pulumi:"dependencyName"`
	MinVersion     pulumi.StringPtrInput `pulumi:"minVersion"`
	Publisher      pulumi.StringPtrInput `pulumi:"publisher"`
}

func (MsixPackageDependenciesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageDependenciesResponse)(nil)).Elem()
}

func (i MsixPackageDependenciesResponseArgs) ToMsixPackageDependenciesResponseOutput() MsixPackageDependenciesResponseOutput {
	return i.ToMsixPackageDependenciesResponseOutputWithContext(context.Background())
}

func (i MsixPackageDependenciesResponseArgs) ToMsixPackageDependenciesResponseOutputWithContext(ctx context.Context) MsixPackageDependenciesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageDependenciesResponseOutput)
}





type MsixPackageDependenciesResponseArrayInput interface {
	pulumi.Input

	ToMsixPackageDependenciesResponseArrayOutput() MsixPackageDependenciesResponseArrayOutput
	ToMsixPackageDependenciesResponseArrayOutputWithContext(context.Context) MsixPackageDependenciesResponseArrayOutput
}

type MsixPackageDependenciesResponseArray []MsixPackageDependenciesResponseInput

func (MsixPackageDependenciesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageDependenciesResponse)(nil)).Elem()
}

func (i MsixPackageDependenciesResponseArray) ToMsixPackageDependenciesResponseArrayOutput() MsixPackageDependenciesResponseArrayOutput {
	return i.ToMsixPackageDependenciesResponseArrayOutputWithContext(context.Background())
}

func (i MsixPackageDependenciesResponseArray) ToMsixPackageDependenciesResponseArrayOutputWithContext(ctx context.Context) MsixPackageDependenciesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageDependenciesResponseArrayOutput)
}

type MsixPackageDependenciesResponseOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageDependenciesResponse)(nil)).Elem()
}

func (o MsixPackageDependenciesResponseOutput) ToMsixPackageDependenciesResponseOutput() MsixPackageDependenciesResponseOutput {
	return o
}

func (o MsixPackageDependenciesResponseOutput) ToMsixPackageDependenciesResponseOutputWithContext(ctx context.Context) MsixPackageDependenciesResponseOutput {
	return o
}

func (o MsixPackageDependenciesResponseOutput) DependencyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependenciesResponse) *string { return v.DependencyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesResponseOutput) MinVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependenciesResponse) *string { return v.MinVersion }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependenciesResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type MsixPackageDependenciesResponseArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageDependenciesResponse)(nil)).Elem()
}

func (o MsixPackageDependenciesResponseArrayOutput) ToMsixPackageDependenciesResponseArrayOutput() MsixPackageDependenciesResponseArrayOutput {
	return o
}

func (o MsixPackageDependenciesResponseArrayOutput) ToMsixPackageDependenciesResponseArrayOutputWithContext(ctx context.Context) MsixPackageDependenciesResponseArrayOutput {
	return o
}

func (o MsixPackageDependenciesResponseArrayOutput) Index(i pulumi.IntInput) MsixPackageDependenciesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageDependenciesResponse {
		return vs[0].([]MsixPackageDependenciesResponse)[vs[1].(int)]
	}).(MsixPackageDependenciesResponseOutput)
}

type RegistrationInfo struct {
	ExpirationTime             *string `pulumi:"expirationTime"`
	RegistrationTokenOperation *string `pulumi:"registrationTokenOperation"`
	Token                      *string `pulumi:"token"`
}





type RegistrationInfoInput interface {
	pulumi.Input

	ToRegistrationInfoOutput() RegistrationInfoOutput
	ToRegistrationInfoOutputWithContext(context.Context) RegistrationInfoOutput
}

type RegistrationInfoArgs struct {
	ExpirationTime             pulumi.StringPtrInput `pulumi:"expirationTime"`
	RegistrationTokenOperation pulumi.StringPtrInput `pulumi:"registrationTokenOperation"`
	Token                      pulumi.StringPtrInput `pulumi:"token"`
}

func (RegistrationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfo)(nil)).Elem()
}

func (i RegistrationInfoArgs) ToRegistrationInfoOutput() RegistrationInfoOutput {
	return i.ToRegistrationInfoOutputWithContext(context.Background())
}

func (i RegistrationInfoArgs) ToRegistrationInfoOutputWithContext(ctx context.Context) RegistrationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoOutput)
}

func (i RegistrationInfoArgs) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return i.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (i RegistrationInfoArgs) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoOutput).ToRegistrationInfoPtrOutputWithContext(ctx)
}









type RegistrationInfoPtrInput interface {
	pulumi.Input

	ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput
	ToRegistrationInfoPtrOutputWithContext(context.Context) RegistrationInfoPtrOutput
}

type registrationInfoPtrType RegistrationInfoArgs

func RegistrationInfoPtr(v *RegistrationInfoArgs) RegistrationInfoPtrInput {
	return (*registrationInfoPtrType)(v)
}

func (*registrationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfo)(nil)).Elem()
}

func (i *registrationInfoPtrType) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return i.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (i *registrationInfoPtrType) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoPtrOutput)
}

type RegistrationInfoOutput struct{ *pulumi.OutputState }

func (RegistrationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfo)(nil)).Elem()
}

func (o RegistrationInfoOutput) ToRegistrationInfoOutput() RegistrationInfoOutput {
	return o
}

func (o RegistrationInfoOutput) ToRegistrationInfoOutputWithContext(ctx context.Context) RegistrationInfoOutput {
	return o
}

func (o RegistrationInfoOutput) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return o.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (o RegistrationInfoOutput) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationInfo) *RegistrationInfo {
		return &v
	}).(RegistrationInfoPtrOutput)
}

func (o RegistrationInfoOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.ExpirationTime }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.RegistrationTokenOperation }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type RegistrationInfoPtrOutput struct{ *pulumi.OutputState }

func (RegistrationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfo)(nil)).Elem()
}

func (o RegistrationInfoPtrOutput) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return o
}

func (o RegistrationInfoPtrOutput) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return o
}

func (o RegistrationInfoPtrOutput) Elem() RegistrationInfoOutput {
	return o.ApplyT(func(v *RegistrationInfo) RegistrationInfo {
		if v != nil {
			return *v
		}
		var ret RegistrationInfo
		return ret
	}).(RegistrationInfoOutput)
}

func (o RegistrationInfoPtrOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoPtrOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationTokenOperation
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoPtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

type RegistrationInfoResponse struct {
	ExpirationTime             *string `pulumi:"expirationTime"`
	RegistrationTokenOperation *string `pulumi:"registrationTokenOperation"`
	Token                      *string `pulumi:"token"`
}





type RegistrationInfoResponseInput interface {
	pulumi.Input

	ToRegistrationInfoResponseOutput() RegistrationInfoResponseOutput
	ToRegistrationInfoResponseOutputWithContext(context.Context) RegistrationInfoResponseOutput
}

type RegistrationInfoResponseArgs struct {
	ExpirationTime             pulumi.StringPtrInput `pulumi:"expirationTime"`
	RegistrationTokenOperation pulumi.StringPtrInput `pulumi:"registrationTokenOperation"`
	Token                      pulumi.StringPtrInput `pulumi:"token"`
}

func (RegistrationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfoResponse)(nil)).Elem()
}

func (i RegistrationInfoResponseArgs) ToRegistrationInfoResponseOutput() RegistrationInfoResponseOutput {
	return i.ToRegistrationInfoResponseOutputWithContext(context.Background())
}

func (i RegistrationInfoResponseArgs) ToRegistrationInfoResponseOutputWithContext(ctx context.Context) RegistrationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoResponseOutput)
}

func (i RegistrationInfoResponseArgs) ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput {
	return i.ToRegistrationInfoResponsePtrOutputWithContext(context.Background())
}

func (i RegistrationInfoResponseArgs) ToRegistrationInfoResponsePtrOutputWithContext(ctx context.Context) RegistrationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoResponseOutput).ToRegistrationInfoResponsePtrOutputWithContext(ctx)
}









type RegistrationInfoResponsePtrInput interface {
	pulumi.Input

	ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput
	ToRegistrationInfoResponsePtrOutputWithContext(context.Context) RegistrationInfoResponsePtrOutput
}

type registrationInfoResponsePtrType RegistrationInfoResponseArgs

func RegistrationInfoResponsePtr(v *RegistrationInfoResponseArgs) RegistrationInfoResponsePtrInput {
	return (*registrationInfoResponsePtrType)(v)
}

func (*registrationInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfoResponse)(nil)).Elem()
}

func (i *registrationInfoResponsePtrType) ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput {
	return i.ToRegistrationInfoResponsePtrOutputWithContext(context.Background())
}

func (i *registrationInfoResponsePtrType) ToRegistrationInfoResponsePtrOutputWithContext(ctx context.Context) RegistrationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoResponsePtrOutput)
}

type RegistrationInfoResponseOutput struct{ *pulumi.OutputState }

func (RegistrationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfoResponse)(nil)).Elem()
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponseOutput() RegistrationInfoResponseOutput {
	return o
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponseOutputWithContext(ctx context.Context) RegistrationInfoResponseOutput {
	return o
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput {
	return o.ToRegistrationInfoResponsePtrOutputWithContext(context.Background())
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponsePtrOutputWithContext(ctx context.Context) RegistrationInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationInfoResponse) *RegistrationInfoResponse {
		return &v
	}).(RegistrationInfoResponsePtrOutput)
}

func (o RegistrationInfoResponseOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.ExpirationTime }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponseOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.RegistrationTokenOperation }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponseOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type RegistrationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (RegistrationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfoResponse)(nil)).Elem()
}

func (o RegistrationInfoResponsePtrOutput) ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput {
	return o
}

func (o RegistrationInfoResponsePtrOutput) ToRegistrationInfoResponsePtrOutputWithContext(ctx context.Context) RegistrationInfoResponsePtrOutput {
	return o
}

func (o RegistrationInfoResponsePtrOutput) Elem() RegistrationInfoResponseOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) RegistrationInfoResponse {
		if v != nil {
			return *v
		}
		var ret RegistrationInfoResponse
		return ret
	}).(RegistrationInfoResponseOutput)
}

func (o RegistrationInfoResponsePtrOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponsePtrOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationTokenOperation
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponsePtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

type ScalingHostPoolReference struct {
	HostPoolArmPath    *string `pulumi:"hostPoolArmPath"`
	ScalingPlanEnabled *bool   `pulumi:"scalingPlanEnabled"`
}





type ScalingHostPoolReferenceInput interface {
	pulumi.Input

	ToScalingHostPoolReferenceOutput() ScalingHostPoolReferenceOutput
	ToScalingHostPoolReferenceOutputWithContext(context.Context) ScalingHostPoolReferenceOutput
}

type ScalingHostPoolReferenceArgs struct {
	HostPoolArmPath    pulumi.StringPtrInput `pulumi:"hostPoolArmPath"`
	ScalingPlanEnabled pulumi.BoolPtrInput   `pulumi:"scalingPlanEnabled"`
}

func (ScalingHostPoolReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingHostPoolReference)(nil)).Elem()
}

func (i ScalingHostPoolReferenceArgs) ToScalingHostPoolReferenceOutput() ScalingHostPoolReferenceOutput {
	return i.ToScalingHostPoolReferenceOutputWithContext(context.Background())
}

func (i ScalingHostPoolReferenceArgs) ToScalingHostPoolReferenceOutputWithContext(ctx context.Context) ScalingHostPoolReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingHostPoolReferenceOutput)
}





type ScalingHostPoolReferenceArrayInput interface {
	pulumi.Input

	ToScalingHostPoolReferenceArrayOutput() ScalingHostPoolReferenceArrayOutput
	ToScalingHostPoolReferenceArrayOutputWithContext(context.Context) ScalingHostPoolReferenceArrayOutput
}

type ScalingHostPoolReferenceArray []ScalingHostPoolReferenceInput

func (ScalingHostPoolReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingHostPoolReference)(nil)).Elem()
}

func (i ScalingHostPoolReferenceArray) ToScalingHostPoolReferenceArrayOutput() ScalingHostPoolReferenceArrayOutput {
	return i.ToScalingHostPoolReferenceArrayOutputWithContext(context.Background())
}

func (i ScalingHostPoolReferenceArray) ToScalingHostPoolReferenceArrayOutputWithContext(ctx context.Context) ScalingHostPoolReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingHostPoolReferenceArrayOutput)
}

type ScalingHostPoolReferenceOutput struct{ *pulumi.OutputState }

func (ScalingHostPoolReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingHostPoolReference)(nil)).Elem()
}

func (o ScalingHostPoolReferenceOutput) ToScalingHostPoolReferenceOutput() ScalingHostPoolReferenceOutput {
	return o
}

func (o ScalingHostPoolReferenceOutput) ToScalingHostPoolReferenceOutputWithContext(ctx context.Context) ScalingHostPoolReferenceOutput {
	return o
}

func (o ScalingHostPoolReferenceOutput) HostPoolArmPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingHostPoolReference) *string { return v.HostPoolArmPath }).(pulumi.StringPtrOutput)
}

func (o ScalingHostPoolReferenceOutput) ScalingPlanEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalingHostPoolReference) *bool { return v.ScalingPlanEnabled }).(pulumi.BoolPtrOutput)
}

type ScalingHostPoolReferenceArrayOutput struct{ *pulumi.OutputState }

func (ScalingHostPoolReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingHostPoolReference)(nil)).Elem()
}

func (o ScalingHostPoolReferenceArrayOutput) ToScalingHostPoolReferenceArrayOutput() ScalingHostPoolReferenceArrayOutput {
	return o
}

func (o ScalingHostPoolReferenceArrayOutput) ToScalingHostPoolReferenceArrayOutputWithContext(ctx context.Context) ScalingHostPoolReferenceArrayOutput {
	return o
}

func (o ScalingHostPoolReferenceArrayOutput) Index(i pulumi.IntInput) ScalingHostPoolReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingHostPoolReference {
		return vs[0].([]ScalingHostPoolReference)[vs[1].(int)]
	}).(ScalingHostPoolReferenceOutput)
}

type ScalingHostPoolReferenceResponse struct {
	HostPoolArmPath    *string `pulumi:"hostPoolArmPath"`
	ScalingPlanEnabled *bool   `pulumi:"scalingPlanEnabled"`
}





type ScalingHostPoolReferenceResponseInput interface {
	pulumi.Input

	ToScalingHostPoolReferenceResponseOutput() ScalingHostPoolReferenceResponseOutput
	ToScalingHostPoolReferenceResponseOutputWithContext(context.Context) ScalingHostPoolReferenceResponseOutput
}

type ScalingHostPoolReferenceResponseArgs struct {
	HostPoolArmPath    pulumi.StringPtrInput `pulumi:"hostPoolArmPath"`
	ScalingPlanEnabled pulumi.BoolPtrInput   `pulumi:"scalingPlanEnabled"`
}

func (ScalingHostPoolReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingHostPoolReferenceResponse)(nil)).Elem()
}

func (i ScalingHostPoolReferenceResponseArgs) ToScalingHostPoolReferenceResponseOutput() ScalingHostPoolReferenceResponseOutput {
	return i.ToScalingHostPoolReferenceResponseOutputWithContext(context.Background())
}

func (i ScalingHostPoolReferenceResponseArgs) ToScalingHostPoolReferenceResponseOutputWithContext(ctx context.Context) ScalingHostPoolReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingHostPoolReferenceResponseOutput)
}





type ScalingHostPoolReferenceResponseArrayInput interface {
	pulumi.Input

	ToScalingHostPoolReferenceResponseArrayOutput() ScalingHostPoolReferenceResponseArrayOutput
	ToScalingHostPoolReferenceResponseArrayOutputWithContext(context.Context) ScalingHostPoolReferenceResponseArrayOutput
}

type ScalingHostPoolReferenceResponseArray []ScalingHostPoolReferenceResponseInput

func (ScalingHostPoolReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingHostPoolReferenceResponse)(nil)).Elem()
}

func (i ScalingHostPoolReferenceResponseArray) ToScalingHostPoolReferenceResponseArrayOutput() ScalingHostPoolReferenceResponseArrayOutput {
	return i.ToScalingHostPoolReferenceResponseArrayOutputWithContext(context.Background())
}

func (i ScalingHostPoolReferenceResponseArray) ToScalingHostPoolReferenceResponseArrayOutputWithContext(ctx context.Context) ScalingHostPoolReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingHostPoolReferenceResponseArrayOutput)
}

type ScalingHostPoolReferenceResponseOutput struct{ *pulumi.OutputState }

func (ScalingHostPoolReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingHostPoolReferenceResponse)(nil)).Elem()
}

func (o ScalingHostPoolReferenceResponseOutput) ToScalingHostPoolReferenceResponseOutput() ScalingHostPoolReferenceResponseOutput {
	return o
}

func (o ScalingHostPoolReferenceResponseOutput) ToScalingHostPoolReferenceResponseOutputWithContext(ctx context.Context) ScalingHostPoolReferenceResponseOutput {
	return o
}

func (o ScalingHostPoolReferenceResponseOutput) HostPoolArmPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingHostPoolReferenceResponse) *string { return v.HostPoolArmPath }).(pulumi.StringPtrOutput)
}

func (o ScalingHostPoolReferenceResponseOutput) ScalingPlanEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalingHostPoolReferenceResponse) *bool { return v.ScalingPlanEnabled }).(pulumi.BoolPtrOutput)
}

type ScalingHostPoolReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ScalingHostPoolReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingHostPoolReferenceResponse)(nil)).Elem()
}

func (o ScalingHostPoolReferenceResponseArrayOutput) ToScalingHostPoolReferenceResponseArrayOutput() ScalingHostPoolReferenceResponseArrayOutput {
	return o
}

func (o ScalingHostPoolReferenceResponseArrayOutput) ToScalingHostPoolReferenceResponseArrayOutputWithContext(ctx context.Context) ScalingHostPoolReferenceResponseArrayOutput {
	return o
}

func (o ScalingHostPoolReferenceResponseArrayOutput) Index(i pulumi.IntInput) ScalingHostPoolReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingHostPoolReferenceResponse {
		return vs[0].([]ScalingHostPoolReferenceResponse)[vs[1].(int)]
	}).(ScalingHostPoolReferenceResponseOutput)
}

type ScalingSchedule struct {
	DaysOfWeek                     []string `pulumi:"daysOfWeek"`
	Name                           *string  `pulumi:"name"`
	OffPeakLoadBalancingAlgorithm  *string  `pulumi:"offPeakLoadBalancingAlgorithm"`
	OffPeakStartTime               *string  `pulumi:"offPeakStartTime"`
	PeakLoadBalancingAlgorithm     *string  `pulumi:"peakLoadBalancingAlgorithm"`
	PeakStartTime                  *string  `pulumi:"peakStartTime"`
	RampDownCapacityThresholdPct   *int     `pulumi:"rampDownCapacityThresholdPct"`
	RampDownForceLogoffUsers       *bool    `pulumi:"rampDownForceLogoffUsers"`
	RampDownLoadBalancingAlgorithm *string  `pulumi:"rampDownLoadBalancingAlgorithm"`
	RampDownMinimumHostsPct        *int     `pulumi:"rampDownMinimumHostsPct"`
	RampDownNotificationMessage    *string  `pulumi:"rampDownNotificationMessage"`
	RampDownStartTime              *string  `pulumi:"rampDownStartTime"`
	RampDownStopHostsWhen          *string  `pulumi:"rampDownStopHostsWhen"`
	RampDownWaitTimeMinutes        *int     `pulumi:"rampDownWaitTimeMinutes"`
	RampUpCapacityThresholdPct     *int     `pulumi:"rampUpCapacityThresholdPct"`
	RampUpLoadBalancingAlgorithm   *string  `pulumi:"rampUpLoadBalancingAlgorithm"`
	RampUpMinimumHostsPct          *int     `pulumi:"rampUpMinimumHostsPct"`
	RampUpStartTime                *string  `pulumi:"rampUpStartTime"`
}





type ScalingScheduleInput interface {
	pulumi.Input

	ToScalingScheduleOutput() ScalingScheduleOutput
	ToScalingScheduleOutputWithContext(context.Context) ScalingScheduleOutput
}

type ScalingScheduleArgs struct {
	DaysOfWeek                     pulumi.StringArrayInput `pulumi:"daysOfWeek"`
	Name                           pulumi.StringPtrInput   `pulumi:"name"`
	OffPeakLoadBalancingAlgorithm  pulumi.StringPtrInput   `pulumi:"offPeakLoadBalancingAlgorithm"`
	OffPeakStartTime               pulumi.StringPtrInput   `pulumi:"offPeakStartTime"`
	PeakLoadBalancingAlgorithm     pulumi.StringPtrInput   `pulumi:"peakLoadBalancingAlgorithm"`
	PeakStartTime                  pulumi.StringPtrInput   `pulumi:"peakStartTime"`
	RampDownCapacityThresholdPct   pulumi.IntPtrInput      `pulumi:"rampDownCapacityThresholdPct"`
	RampDownForceLogoffUsers       pulumi.BoolPtrInput     `pulumi:"rampDownForceLogoffUsers"`
	RampDownLoadBalancingAlgorithm pulumi.StringPtrInput   `pulumi:"rampDownLoadBalancingAlgorithm"`
	RampDownMinimumHostsPct        pulumi.IntPtrInput      `pulumi:"rampDownMinimumHostsPct"`
	RampDownNotificationMessage    pulumi.StringPtrInput   `pulumi:"rampDownNotificationMessage"`
	RampDownStartTime              pulumi.StringPtrInput   `pulumi:"rampDownStartTime"`
	RampDownStopHostsWhen          pulumi.StringPtrInput   `pulumi:"rampDownStopHostsWhen"`
	RampDownWaitTimeMinutes        pulumi.IntPtrInput      `pulumi:"rampDownWaitTimeMinutes"`
	RampUpCapacityThresholdPct     pulumi.IntPtrInput      `pulumi:"rampUpCapacityThresholdPct"`
	RampUpLoadBalancingAlgorithm   pulumi.StringPtrInput   `pulumi:"rampUpLoadBalancingAlgorithm"`
	RampUpMinimumHostsPct          pulumi.IntPtrInput      `pulumi:"rampUpMinimumHostsPct"`
	RampUpStartTime                pulumi.StringPtrInput   `pulumi:"rampUpStartTime"`
}

func (ScalingScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingSchedule)(nil)).Elem()
}

func (i ScalingScheduleArgs) ToScalingScheduleOutput() ScalingScheduleOutput {
	return i.ToScalingScheduleOutputWithContext(context.Background())
}

func (i ScalingScheduleArgs) ToScalingScheduleOutputWithContext(ctx context.Context) ScalingScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingScheduleOutput)
}





type ScalingScheduleArrayInput interface {
	pulumi.Input

	ToScalingScheduleArrayOutput() ScalingScheduleArrayOutput
	ToScalingScheduleArrayOutputWithContext(context.Context) ScalingScheduleArrayOutput
}

type ScalingScheduleArray []ScalingScheduleInput

func (ScalingScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingSchedule)(nil)).Elem()
}

func (i ScalingScheduleArray) ToScalingScheduleArrayOutput() ScalingScheduleArrayOutput {
	return i.ToScalingScheduleArrayOutputWithContext(context.Background())
}

func (i ScalingScheduleArray) ToScalingScheduleArrayOutputWithContext(ctx context.Context) ScalingScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingScheduleArrayOutput)
}

type ScalingScheduleOutput struct{ *pulumi.OutputState }

func (ScalingScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingSchedule)(nil)).Elem()
}

func (o ScalingScheduleOutput) ToScalingScheduleOutput() ScalingScheduleOutput {
	return o
}

func (o ScalingScheduleOutput) ToScalingScheduleOutputWithContext(ctx context.Context) ScalingScheduleOutput {
	return o
}

func (o ScalingScheduleOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScalingSchedule) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o ScalingScheduleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) OffPeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.OffPeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) OffPeakStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.OffPeakStartTime }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) PeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.PeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) PeakStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.PeakStartTime }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) RampDownCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *int { return v.RampDownCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleOutput) RampDownForceLogoffUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *bool { return v.RampDownForceLogoffUsers }).(pulumi.BoolPtrOutput)
}

func (o ScalingScheduleOutput) RampDownLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.RampDownLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) RampDownMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *int { return v.RampDownMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleOutput) RampDownNotificationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.RampDownNotificationMessage }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) RampDownStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.RampDownStartTime }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) RampDownStopHostsWhen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.RampDownStopHostsWhen }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) RampDownWaitTimeMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *int { return v.RampDownWaitTimeMinutes }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleOutput) RampUpCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *int { return v.RampUpCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleOutput) RampUpLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.RampUpLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) RampUpMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *int { return v.RampUpMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleOutput) RampUpStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.RampUpStartTime }).(pulumi.StringPtrOutput)
}

type ScalingScheduleArrayOutput struct{ *pulumi.OutputState }

func (ScalingScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingSchedule)(nil)).Elem()
}

func (o ScalingScheduleArrayOutput) ToScalingScheduleArrayOutput() ScalingScheduleArrayOutput {
	return o
}

func (o ScalingScheduleArrayOutput) ToScalingScheduleArrayOutputWithContext(ctx context.Context) ScalingScheduleArrayOutput {
	return o
}

func (o ScalingScheduleArrayOutput) Index(i pulumi.IntInput) ScalingScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingSchedule {
		return vs[0].([]ScalingSchedule)[vs[1].(int)]
	}).(ScalingScheduleOutput)
}

type ScalingScheduleResponse struct {
	DaysOfWeek                     []string `pulumi:"daysOfWeek"`
	Name                           *string  `pulumi:"name"`
	OffPeakLoadBalancingAlgorithm  *string  `pulumi:"offPeakLoadBalancingAlgorithm"`
	OffPeakStartTime               *string  `pulumi:"offPeakStartTime"`
	PeakLoadBalancingAlgorithm     *string  `pulumi:"peakLoadBalancingAlgorithm"`
	PeakStartTime                  *string  `pulumi:"peakStartTime"`
	RampDownCapacityThresholdPct   *int     `pulumi:"rampDownCapacityThresholdPct"`
	RampDownForceLogoffUsers       *bool    `pulumi:"rampDownForceLogoffUsers"`
	RampDownLoadBalancingAlgorithm *string  `pulumi:"rampDownLoadBalancingAlgorithm"`
	RampDownMinimumHostsPct        *int     `pulumi:"rampDownMinimumHostsPct"`
	RampDownNotificationMessage    *string  `pulumi:"rampDownNotificationMessage"`
	RampDownStartTime              *string  `pulumi:"rampDownStartTime"`
	RampDownStopHostsWhen          *string  `pulumi:"rampDownStopHostsWhen"`
	RampDownWaitTimeMinutes        *int     `pulumi:"rampDownWaitTimeMinutes"`
	RampUpCapacityThresholdPct     *int     `pulumi:"rampUpCapacityThresholdPct"`
	RampUpLoadBalancingAlgorithm   *string  `pulumi:"rampUpLoadBalancingAlgorithm"`
	RampUpMinimumHostsPct          *int     `pulumi:"rampUpMinimumHostsPct"`
	RampUpStartTime                *string  `pulumi:"rampUpStartTime"`
}





type ScalingScheduleResponseInput interface {
	pulumi.Input

	ToScalingScheduleResponseOutput() ScalingScheduleResponseOutput
	ToScalingScheduleResponseOutputWithContext(context.Context) ScalingScheduleResponseOutput
}

type ScalingScheduleResponseArgs struct {
	DaysOfWeek                     pulumi.StringArrayInput `pulumi:"daysOfWeek"`
	Name                           pulumi.StringPtrInput   `pulumi:"name"`
	OffPeakLoadBalancingAlgorithm  pulumi.StringPtrInput   `pulumi:"offPeakLoadBalancingAlgorithm"`
	OffPeakStartTime               pulumi.StringPtrInput   `pulumi:"offPeakStartTime"`
	PeakLoadBalancingAlgorithm     pulumi.StringPtrInput   `pulumi:"peakLoadBalancingAlgorithm"`
	PeakStartTime                  pulumi.StringPtrInput   `pulumi:"peakStartTime"`
	RampDownCapacityThresholdPct   pulumi.IntPtrInput      `pulumi:"rampDownCapacityThresholdPct"`
	RampDownForceLogoffUsers       pulumi.BoolPtrInput     `pulumi:"rampDownForceLogoffUsers"`
	RampDownLoadBalancingAlgorithm pulumi.StringPtrInput   `pulumi:"rampDownLoadBalancingAlgorithm"`
	RampDownMinimumHostsPct        pulumi.IntPtrInput      `pulumi:"rampDownMinimumHostsPct"`
	RampDownNotificationMessage    pulumi.StringPtrInput   `pulumi:"rampDownNotificationMessage"`
	RampDownStartTime              pulumi.StringPtrInput   `pulumi:"rampDownStartTime"`
	RampDownStopHostsWhen          pulumi.StringPtrInput   `pulumi:"rampDownStopHostsWhen"`
	RampDownWaitTimeMinutes        pulumi.IntPtrInput      `pulumi:"rampDownWaitTimeMinutes"`
	RampUpCapacityThresholdPct     pulumi.IntPtrInput      `pulumi:"rampUpCapacityThresholdPct"`
	RampUpLoadBalancingAlgorithm   pulumi.StringPtrInput   `pulumi:"rampUpLoadBalancingAlgorithm"`
	RampUpMinimumHostsPct          pulumi.IntPtrInput      `pulumi:"rampUpMinimumHostsPct"`
	RampUpStartTime                pulumi.StringPtrInput   `pulumi:"rampUpStartTime"`
}

func (ScalingScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingScheduleResponse)(nil)).Elem()
}

func (i ScalingScheduleResponseArgs) ToScalingScheduleResponseOutput() ScalingScheduleResponseOutput {
	return i.ToScalingScheduleResponseOutputWithContext(context.Background())
}

func (i ScalingScheduleResponseArgs) ToScalingScheduleResponseOutputWithContext(ctx context.Context) ScalingScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingScheduleResponseOutput)
}





type ScalingScheduleResponseArrayInput interface {
	pulumi.Input

	ToScalingScheduleResponseArrayOutput() ScalingScheduleResponseArrayOutput
	ToScalingScheduleResponseArrayOutputWithContext(context.Context) ScalingScheduleResponseArrayOutput
}

type ScalingScheduleResponseArray []ScalingScheduleResponseInput

func (ScalingScheduleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingScheduleResponse)(nil)).Elem()
}

func (i ScalingScheduleResponseArray) ToScalingScheduleResponseArrayOutput() ScalingScheduleResponseArrayOutput {
	return i.ToScalingScheduleResponseArrayOutputWithContext(context.Background())
}

func (i ScalingScheduleResponseArray) ToScalingScheduleResponseArrayOutputWithContext(ctx context.Context) ScalingScheduleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingScheduleResponseArrayOutput)
}

type ScalingScheduleResponseOutput struct{ *pulumi.OutputState }

func (ScalingScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingScheduleResponse)(nil)).Elem()
}

func (o ScalingScheduleResponseOutput) ToScalingScheduleResponseOutput() ScalingScheduleResponseOutput {
	return o
}

func (o ScalingScheduleResponseOutput) ToScalingScheduleResponseOutputWithContext(ctx context.Context) ScalingScheduleResponseOutput {
	return o
}

func (o ScalingScheduleResponseOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o ScalingScheduleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) OffPeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.OffPeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) OffPeakStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.OffPeakStartTime }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) PeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.PeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) PeakStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.PeakStartTime }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *int { return v.RampDownCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownForceLogoffUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *bool { return v.RampDownForceLogoffUsers }).(pulumi.BoolPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.RampDownLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *int { return v.RampDownMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownNotificationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.RampDownNotificationMessage }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.RampDownStartTime }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownStopHostsWhen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.RampDownStopHostsWhen }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownWaitTimeMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *int { return v.RampDownWaitTimeMinutes }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampUpCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *int { return v.RampUpCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampUpLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.RampUpLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampUpMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *int { return v.RampUpMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampUpStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.RampUpStartTime }).(pulumi.StringPtrOutput)
}

type ScalingScheduleResponseArrayOutput struct{ *pulumi.OutputState }

func (ScalingScheduleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingScheduleResponse)(nil)).Elem()
}

func (o ScalingScheduleResponseArrayOutput) ToScalingScheduleResponseArrayOutput() ScalingScheduleResponseArrayOutput {
	return o
}

func (o ScalingScheduleResponseArrayOutput) ToScalingScheduleResponseArrayOutputWithContext(ctx context.Context) ScalingScheduleResponseArrayOutput {
	return o
}

func (o ScalingScheduleResponseArrayOutput) Index(i pulumi.IntInput) ScalingScheduleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingScheduleResponse {
		return vs[0].([]ScalingScheduleResponse)[vs[1].(int)]
	}).(ScalingScheduleResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MsixPackageApplicationsInput)(nil)).Elem(), MsixPackageApplicationsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsixPackageApplicationsArrayInput)(nil)).Elem(), MsixPackageApplicationsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsixPackageApplicationsResponseInput)(nil)).Elem(), MsixPackageApplicationsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsixPackageApplicationsResponseArrayInput)(nil)).Elem(), MsixPackageApplicationsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsixPackageDependenciesInput)(nil)).Elem(), MsixPackageDependenciesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsixPackageDependenciesArrayInput)(nil)).Elem(), MsixPackageDependenciesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsixPackageDependenciesResponseInput)(nil)).Elem(), MsixPackageDependenciesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsixPackageDependenciesResponseArrayInput)(nil)).Elem(), MsixPackageDependenciesResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistrationInfoInput)(nil)).Elem(), RegistrationInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistrationInfoPtrInput)(nil)).Elem(), RegistrationInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistrationInfoResponseInput)(nil)).Elem(), RegistrationInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistrationInfoResponsePtrInput)(nil)).Elem(), RegistrationInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingHostPoolReferenceInput)(nil)).Elem(), ScalingHostPoolReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingHostPoolReferenceArrayInput)(nil)).Elem(), ScalingHostPoolReferenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingHostPoolReferenceResponseInput)(nil)).Elem(), ScalingHostPoolReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingHostPoolReferenceResponseArrayInput)(nil)).Elem(), ScalingHostPoolReferenceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingScheduleInput)(nil)).Elem(), ScalingScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingScheduleArrayInput)(nil)).Elem(), ScalingScheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingScheduleResponseInput)(nil)).Elem(), ScalingScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingScheduleResponseArrayInput)(nil)).Elem(), ScalingScheduleResponseArray{})
	pulumi.RegisterOutputType(MsixPackageApplicationsOutput{})
	pulumi.RegisterOutputType(MsixPackageApplicationsArrayOutput{})
	pulumi.RegisterOutputType(MsixPackageApplicationsResponseOutput{})
	pulumi.RegisterOutputType(MsixPackageApplicationsResponseArrayOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesArrayOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesResponseOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesResponseArrayOutput{})
	pulumi.RegisterOutputType(RegistrationInfoOutput{})
	pulumi.RegisterOutputType(RegistrationInfoPtrOutput{})
	pulumi.RegisterOutputType(RegistrationInfoResponseOutput{})
	pulumi.RegisterOutputType(RegistrationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ScalingHostPoolReferenceOutput{})
	pulumi.RegisterOutputType(ScalingHostPoolReferenceArrayOutput{})
	pulumi.RegisterOutputType(ScalingHostPoolReferenceResponseOutput{})
	pulumi.RegisterOutputType(ScalingHostPoolReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(ScalingScheduleOutput{})
	pulumi.RegisterOutputType(ScalingScheduleArrayOutput{})
	pulumi.RegisterOutputType(ScalingScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScalingScheduleResponseArrayOutput{})
}
