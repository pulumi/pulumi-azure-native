


package v20200802

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ErrorDetailResponse struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
	Target  *string               `pulumi:"target"`
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorDetailResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
}

type LocationData struct {
	City            *string `pulumi:"city"`
	CountryOrRegion *string `pulumi:"countryOrRegion"`
	District        *string `pulumi:"district"`
	Name            string  `pulumi:"name"`
}





type LocationDataInput interface {
	pulumi.Input

	ToLocationDataOutput() LocationDataOutput
	ToLocationDataOutputWithContext(context.Context) LocationDataOutput
}

type LocationDataArgs struct {
	City            pulumi.StringPtrInput `pulumi:"city"`
	CountryOrRegion pulumi.StringPtrInput `pulumi:"countryOrRegion"`
	District        pulumi.StringPtrInput `pulumi:"district"`
	Name            pulumi.StringInput    `pulumi:"name"`
}

func (LocationDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationData)(nil)).Elem()
}

func (i LocationDataArgs) ToLocationDataOutput() LocationDataOutput {
	return i.ToLocationDataOutputWithContext(context.Background())
}

func (i LocationDataArgs) ToLocationDataOutputWithContext(ctx context.Context) LocationDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataOutput)
}

func (i LocationDataArgs) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return i.ToLocationDataPtrOutputWithContext(context.Background())
}

func (i LocationDataArgs) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataOutput).ToLocationDataPtrOutputWithContext(ctx)
}









type LocationDataPtrInput interface {
	pulumi.Input

	ToLocationDataPtrOutput() LocationDataPtrOutput
	ToLocationDataPtrOutputWithContext(context.Context) LocationDataPtrOutput
}

type locationDataPtrType LocationDataArgs

func LocationDataPtr(v *LocationDataArgs) LocationDataPtrInput {
	return (*locationDataPtrType)(v)
}

func (*locationDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationData)(nil)).Elem()
}

func (i *locationDataPtrType) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return i.ToLocationDataPtrOutputWithContext(context.Background())
}

func (i *locationDataPtrType) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataPtrOutput)
}

type LocationDataOutput struct{ *pulumi.OutputState }

func (LocationDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationData)(nil)).Elem()
}

func (o LocationDataOutput) ToLocationDataOutput() LocationDataOutput {
	return o
}

func (o LocationDataOutput) ToLocationDataOutputWithContext(ctx context.Context) LocationDataOutput {
	return o
}

func (o LocationDataOutput) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return o.ToLocationDataPtrOutputWithContext(context.Background())
}

func (o LocationDataOutput) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocationData) *LocationData {
		return &v
	}).(LocationDataPtrOutput)
}

func (o LocationDataOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o LocationDataOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

func (o LocationDataOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.District }).(pulumi.StringPtrOutput)
}

func (o LocationDataOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LocationData) string { return v.Name }).(pulumi.StringOutput)
}

type LocationDataPtrOutput struct{ *pulumi.OutputState }

func (LocationDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationData)(nil)).Elem()
}

func (o LocationDataPtrOutput) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return o
}

func (o LocationDataPtrOutput) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return o
}

func (o LocationDataPtrOutput) Elem() LocationDataOutput {
	return o.ApplyT(func(v *LocationData) LocationData {
		if v != nil {
			return *v
		}
		var ret LocationData
		return ret
	}).(LocationDataOutput)
}

func (o LocationDataPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataPtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataPtrOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.District
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type LocationDataResponse struct {
	City            *string `pulumi:"city"`
	CountryOrRegion *string `pulumi:"countryOrRegion"`
	District        *string `pulumi:"district"`
	Name            string  `pulumi:"name"`
}

type LocationDataResponseOutput struct{ *pulumi.OutputState }

func (LocationDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationDataResponse)(nil)).Elem()
}

func (o LocationDataResponseOutput) ToLocationDataResponseOutput() LocationDataResponseOutput {
	return o
}

func (o LocationDataResponseOutput) ToLocationDataResponseOutputWithContext(ctx context.Context) LocationDataResponseOutput {
	return o
}

func (o LocationDataResponseOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationDataResponse) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o LocationDataResponseOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationDataResponse) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

func (o LocationDataResponseOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationDataResponse) *string { return v.District }).(pulumi.StringPtrOutput)
}

func (o LocationDataResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LocationDataResponse) string { return v.Name }).(pulumi.StringOutput)
}

type LocationDataResponsePtrOutput struct{ *pulumi.OutputState }

func (LocationDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationDataResponse)(nil)).Elem()
}

func (o LocationDataResponsePtrOutput) ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput {
	return o
}

func (o LocationDataResponsePtrOutput) ToLocationDataResponsePtrOutputWithContext(ctx context.Context) LocationDataResponsePtrOutput {
	return o
}

func (o LocationDataResponsePtrOutput) Elem() LocationDataResponseOutput {
	return o.ApplyT(func(v *LocationDataResponse) LocationDataResponse {
		if v != nil {
			return *v
		}
		var ret LocationDataResponse
		return ret
	}).(LocationDataResponseOutput)
}

func (o LocationDataResponsePtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataResponsePtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataResponsePtrOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.District
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewResponse struct {
	Name               string                                      `pulumi:"name"`
	Status             *MachineExtensionInstanceViewResponseStatus `pulumi:"status"`
	Type               string                                      `pulumi:"type"`
	TypeHandlerVersion string                                      `pulumi:"typeHandlerVersion"`
}

type MachineExtensionInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseOutput) ToMachineExtensionInstanceViewResponseOutput() MachineExtensionInstanceViewResponseOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseOutput) ToMachineExtensionInstanceViewResponseOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) *MachineExtensionInstanceViewResponseStatus {
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type MachineExtensionInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseArrayOutput) ToMachineExtensionInstanceViewResponseArrayOutput() MachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseArrayOutput) ToMachineExtensionInstanceViewResponseArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) MachineExtensionInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineExtensionInstanceViewResponse {
		return vs[0].([]MachineExtensionInstanceViewResponse)[vs[1].(int)]
	}).(MachineExtensionInstanceViewResponseOutput)
}

type MachineExtensionInstanceViewResponseStatus struct {
	Code          string `pulumi:"code"`
	DisplayStatus string `pulumi:"displayStatus"`
	Level         string `pulumi:"level"`
	Message       string `pulumi:"message"`
	Time          string `pulumi:"time"`
}

type MachineExtensionInstanceViewResponseStatusOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Code }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) DisplayStatus() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.DisplayStatus }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Level }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Message }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Time }).(pulumi.StringOutput)
}

type MachineExtensionInstanceViewResponseStatusPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Elem() MachineExtensionInstanceViewResponseStatusOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) MachineExtensionInstanceViewResponseStatus {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceViewResponseStatus
		return ret
	}).(MachineExtensionInstanceViewResponseStatusOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Level
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Time
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionPropertiesResponseInstanceView struct {
	Name               string                                      `pulumi:"name"`
	Status             *MachineExtensionInstanceViewResponseStatus `pulumi:"status"`
	Type               string                                      `pulumi:"type"`
	TypeHandlerVersion string                                      `pulumi:"typeHandlerVersion"`
}

type MachineExtensionPropertiesResponseInstanceViewOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesResponseInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionPropertiesResponseInstanceView)(nil)).Elem()
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) ToMachineExtensionPropertiesResponseInstanceViewOutput() MachineExtensionPropertiesResponseInstanceViewOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) ToMachineExtensionPropertiesResponseInstanceViewOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.Name }).(pulumi.StringOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) *MachineExtensionInstanceViewResponseStatus {
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.Type }).(pulumi.StringOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type MachineExtensionPropertiesResponseInstanceViewPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesResponseInstanceViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionPropertiesResponseInstanceView)(nil)).Elem()
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) ToMachineExtensionPropertiesResponseInstanceViewPtrOutput() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Elem() MachineExtensionPropertiesResponseInstanceViewOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) MachineExtensionPropertiesResponseInstanceView {
		if v != nil {
			return *v
		}
		var ret MachineExtensionPropertiesResponseInstanceView
		return ret
	}).(MachineExtensionPropertiesResponseInstanceViewOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *MachineExtensionInstanceViewResponseStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type MachineIdentity struct {
	Type *string `pulumi:"type"`
}





type MachineIdentityInput interface {
	pulumi.Input

	ToMachineIdentityOutput() MachineIdentityOutput
	ToMachineIdentityOutputWithContext(context.Context) MachineIdentityOutput
}

type MachineIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (MachineIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineIdentity)(nil)).Elem()
}

func (i MachineIdentityArgs) ToMachineIdentityOutput() MachineIdentityOutput {
	return i.ToMachineIdentityOutputWithContext(context.Background())
}

func (i MachineIdentityArgs) ToMachineIdentityOutputWithContext(ctx context.Context) MachineIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineIdentityOutput)
}

func (i MachineIdentityArgs) ToMachineIdentityPtrOutput() MachineIdentityPtrOutput {
	return i.ToMachineIdentityPtrOutputWithContext(context.Background())
}

func (i MachineIdentityArgs) ToMachineIdentityPtrOutputWithContext(ctx context.Context) MachineIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineIdentityOutput).ToMachineIdentityPtrOutputWithContext(ctx)
}









type MachineIdentityPtrInput interface {
	pulumi.Input

	ToMachineIdentityPtrOutput() MachineIdentityPtrOutput
	ToMachineIdentityPtrOutputWithContext(context.Context) MachineIdentityPtrOutput
}

type machineIdentityPtrType MachineIdentityArgs

func MachineIdentityPtr(v *MachineIdentityArgs) MachineIdentityPtrInput {
	return (*machineIdentityPtrType)(v)
}

func (*machineIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineIdentity)(nil)).Elem()
}

func (i *machineIdentityPtrType) ToMachineIdentityPtrOutput() MachineIdentityPtrOutput {
	return i.ToMachineIdentityPtrOutputWithContext(context.Background())
}

func (i *machineIdentityPtrType) ToMachineIdentityPtrOutputWithContext(ctx context.Context) MachineIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineIdentityPtrOutput)
}

type MachineIdentityOutput struct{ *pulumi.OutputState }

func (MachineIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineIdentity)(nil)).Elem()
}

func (o MachineIdentityOutput) ToMachineIdentityOutput() MachineIdentityOutput {
	return o
}

func (o MachineIdentityOutput) ToMachineIdentityOutputWithContext(ctx context.Context) MachineIdentityOutput {
	return o
}

func (o MachineIdentityOutput) ToMachineIdentityPtrOutput() MachineIdentityPtrOutput {
	return o.ToMachineIdentityPtrOutputWithContext(context.Background())
}

func (o MachineIdentityOutput) ToMachineIdentityPtrOutputWithContext(ctx context.Context) MachineIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineIdentity) *MachineIdentity {
		return &v
	}).(MachineIdentityPtrOutput)
}

func (o MachineIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type MachineIdentityPtrOutput struct{ *pulumi.OutputState }

func (MachineIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineIdentity)(nil)).Elem()
}

func (o MachineIdentityPtrOutput) ToMachineIdentityPtrOutput() MachineIdentityPtrOutput {
	return o
}

func (o MachineIdentityPtrOutput) ToMachineIdentityPtrOutputWithContext(ctx context.Context) MachineIdentityPtrOutput {
	return o
}

func (o MachineIdentityPtrOutput) Elem() MachineIdentityOutput {
	return o.ApplyT(func(v *MachineIdentity) MachineIdentity {
		if v != nil {
			return *v
		}
		var ret MachineIdentity
		return ret
	}).(MachineIdentityOutput)
}

func (o MachineIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type MachinePropertiesResponseOsProfile struct {
	ComputerName string `pulumi:"computerName"`
}

type MachinePropertiesResponseOsProfileOutput struct{ *pulumi.OutputState }

func (MachinePropertiesResponseOsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachinePropertiesResponseOsProfile)(nil)).Elem()
}

func (o MachinePropertiesResponseOsProfileOutput) ToMachinePropertiesResponseOsProfileOutput() MachinePropertiesResponseOsProfileOutput {
	return o
}

func (o MachinePropertiesResponseOsProfileOutput) ToMachinePropertiesResponseOsProfileOutputWithContext(ctx context.Context) MachinePropertiesResponseOsProfileOutput {
	return o
}

func (o MachinePropertiesResponseOsProfileOutput) ComputerName() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponseOsProfile) string { return v.ComputerName }).(pulumi.StringOutput)
}

type MachinePropertiesResponseOsProfilePtrOutput struct{ *pulumi.OutputState }

func (MachinePropertiesResponseOsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachinePropertiesResponseOsProfile)(nil)).Elem()
}

func (o MachinePropertiesResponseOsProfilePtrOutput) ToMachinePropertiesResponseOsProfilePtrOutput() MachinePropertiesResponseOsProfilePtrOutput {
	return o
}

func (o MachinePropertiesResponseOsProfilePtrOutput) ToMachinePropertiesResponseOsProfilePtrOutputWithContext(ctx context.Context) MachinePropertiesResponseOsProfilePtrOutput {
	return o
}

func (o MachinePropertiesResponseOsProfilePtrOutput) Elem() MachinePropertiesResponseOsProfileOutput {
	return o.ApplyT(func(v *MachinePropertiesResponseOsProfile) MachinePropertiesResponseOsProfile {
		if v != nil {
			return *v
		}
		var ret MachinePropertiesResponseOsProfile
		return ret
	}).(MachinePropertiesResponseOsProfileOutput)
}

func (o MachinePropertiesResponseOsProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponseOsProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ComputerName
	}).(pulumi.StringPtrOutput)
}

type MachineResponseIdentity struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type MachineResponseIdentityOutput struct{ *pulumi.OutputState }

func (MachineResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineResponseIdentity)(nil)).Elem()
}

func (o MachineResponseIdentityOutput) ToMachineResponseIdentityOutput() MachineResponseIdentityOutput {
	return o
}

func (o MachineResponseIdentityOutput) ToMachineResponseIdentityOutputWithContext(ctx context.Context) MachineResponseIdentityOutput {
	return o
}

func (o MachineResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v MachineResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o MachineResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v MachineResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o MachineResponseIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineResponseIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type MachineResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (MachineResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineResponseIdentity)(nil)).Elem()
}

func (o MachineResponseIdentityPtrOutput) ToMachineResponseIdentityPtrOutput() MachineResponseIdentityPtrOutput {
	return o
}

func (o MachineResponseIdentityPtrOutput) ToMachineResponseIdentityPtrOutputWithContext(ctx context.Context) MachineResponseIdentityPtrOutput {
	return o
}

func (o MachineResponseIdentityPtrOutput) Elem() MachineResponseIdentityOutput {
	return o.ApplyT(func(v *MachineResponseIdentity) MachineResponseIdentity {
		if v != nil {
			return *v
		}
		var ret MachineResponseIdentity
		return ret
	}).(MachineResponseIdentityOutput)
}

func (o MachineResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o MachineResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o MachineResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(LocationDataOutput{})
	pulumi.RegisterOutputType(LocationDataPtrOutput{})
	pulumi.RegisterOutputType(LocationDataResponseOutput{})
	pulumi.RegisterOutputType(LocationDataResponsePtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionPropertiesResponseInstanceViewOutput{})
	pulumi.RegisterOutputType(MachineExtensionPropertiesResponseInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(MachineIdentityOutput{})
	pulumi.RegisterOutputType(MachineIdentityPtrOutput{})
	pulumi.RegisterOutputType(MachinePropertiesResponseOsProfileOutput{})
	pulumi.RegisterOutputType(MachinePropertiesResponseOsProfilePtrOutput{})
	pulumi.RegisterOutputType(MachineResponseIdentityOutput{})
	pulumi.RegisterOutputType(MachineResponseIdentityPtrOutput{})
}
