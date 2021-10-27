


package v20190802preview

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





type ErrorDetailResponseInput interface {
	pulumi.Input

	ToErrorDetailResponseOutput() ErrorDetailResponseOutput
	ToErrorDetailResponseOutputWithContext(context.Context) ErrorDetailResponseOutput
}

type ErrorDetailResponseArgs struct {
	Code    pulumi.StringInput            `pulumi:"code"`
	Details ErrorDetailResponseArrayInput `pulumi:"details"`
	Message pulumi.StringInput            `pulumi:"message"`
	Target  pulumi.StringPtrInput         `pulumi:"target"`
}

func (ErrorDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return i.ToErrorDetailResponseOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput)
}





type ErrorDetailResponseArrayInput interface {
	pulumi.Input

	ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput
	ToErrorDetailResponseArrayOutputWithContext(context.Context) ErrorDetailResponseArrayOutput
}

type ErrorDetailResponseArray []ErrorDetailResponseInput

func (ErrorDetailResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return i.ToErrorDetailResponseArrayOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseArrayOutput)
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

type MachineExtensionInstanceView struct {
	Name               *string                             `pulumi:"name"`
	Status             *MachineExtensionInstanceViewStatus `pulumi:"status"`
	Type               *string                             `pulumi:"type"`
	TypeHandlerVersion *string                             `pulumi:"typeHandlerVersion"`
}





type MachineExtensionInstanceViewInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewOutput() MachineExtensionInstanceViewOutput
	ToMachineExtensionInstanceViewOutputWithContext(context.Context) MachineExtensionInstanceViewOutput
}

type MachineExtensionInstanceViewArgs struct {
	Name               pulumi.StringPtrInput                      `pulumi:"name"`
	Status             MachineExtensionInstanceViewStatusPtrInput `pulumi:"status"`
	Type               pulumi.StringPtrInput                      `pulumi:"type"`
	TypeHandlerVersion pulumi.StringPtrInput                      `pulumi:"typeHandlerVersion"`
}

func (MachineExtensionInstanceViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceView)(nil)).Elem()
}

func (i MachineExtensionInstanceViewArgs) ToMachineExtensionInstanceViewOutput() MachineExtensionInstanceViewOutput {
	return i.ToMachineExtensionInstanceViewOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewArgs) ToMachineExtensionInstanceViewOutputWithContext(ctx context.Context) MachineExtensionInstanceViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewOutput)
}

func (i MachineExtensionInstanceViewArgs) ToMachineExtensionInstanceViewPtrOutput() MachineExtensionInstanceViewPtrOutput {
	return i.ToMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewArgs) ToMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewOutput).ToMachineExtensionInstanceViewPtrOutputWithContext(ctx)
}









type MachineExtensionInstanceViewPtrInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewPtrOutput() MachineExtensionInstanceViewPtrOutput
	ToMachineExtensionInstanceViewPtrOutputWithContext(context.Context) MachineExtensionInstanceViewPtrOutput
}

type machineExtensionInstanceViewPtrType MachineExtensionInstanceViewArgs

func MachineExtensionInstanceViewPtr(v *MachineExtensionInstanceViewArgs) MachineExtensionInstanceViewPtrInput {
	return (*machineExtensionInstanceViewPtrType)(v)
}

func (*machineExtensionInstanceViewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceView)(nil)).Elem()
}

func (i *machineExtensionInstanceViewPtrType) ToMachineExtensionInstanceViewPtrOutput() MachineExtensionInstanceViewPtrOutput {
	return i.ToMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (i *machineExtensionInstanceViewPtrType) ToMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewPtrOutput)
}





type MachineExtensionInstanceViewArrayInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewArrayOutput() MachineExtensionInstanceViewArrayOutput
	ToMachineExtensionInstanceViewArrayOutputWithContext(context.Context) MachineExtensionInstanceViewArrayOutput
}

type MachineExtensionInstanceViewArray []MachineExtensionInstanceViewInput

func (MachineExtensionInstanceViewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceView)(nil)).Elem()
}

func (i MachineExtensionInstanceViewArray) ToMachineExtensionInstanceViewArrayOutput() MachineExtensionInstanceViewArrayOutput {
	return i.ToMachineExtensionInstanceViewArrayOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewArray) ToMachineExtensionInstanceViewArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewArrayOutput)
}

type MachineExtensionInstanceViewOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceView)(nil)).Elem()
}

func (o MachineExtensionInstanceViewOutput) ToMachineExtensionInstanceViewOutput() MachineExtensionInstanceViewOutput {
	return o
}

func (o MachineExtensionInstanceViewOutput) ToMachineExtensionInstanceViewOutputWithContext(ctx context.Context) MachineExtensionInstanceViewOutput {
	return o
}

func (o MachineExtensionInstanceViewOutput) ToMachineExtensionInstanceViewPtrOutput() MachineExtensionInstanceViewPtrOutput {
	return o.ToMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (o MachineExtensionInstanceViewOutput) ToMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionInstanceView) *MachineExtensionInstanceView {
		return &v
	}).(MachineExtensionInstanceViewPtrOutput)
}

func (o MachineExtensionInstanceViewOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewOutput) Status() MachineExtensionInstanceViewStatusPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *MachineExtensionInstanceViewStatus { return v.Status }).(MachineExtensionInstanceViewStatusPtrOutput)
}

func (o MachineExtensionInstanceViewOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceView)(nil)).Elem()
}

func (o MachineExtensionInstanceViewPtrOutput) ToMachineExtensionInstanceViewPtrOutput() MachineExtensionInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewPtrOutput) ToMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewPtrOutput) Elem() MachineExtensionInstanceViewOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceView) MachineExtensionInstanceView {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceView
		return ret
	}).(MachineExtensionInstanceViewOutput)
}

func (o MachineExtensionInstanceViewPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewPtrOutput) Status() MachineExtensionInstanceViewStatusPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceView) *MachineExtensionInstanceViewStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(MachineExtensionInstanceViewStatusPtrOutput)
}

func (o MachineExtensionInstanceViewPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewArrayOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceView)(nil)).Elem()
}

func (o MachineExtensionInstanceViewArrayOutput) ToMachineExtensionInstanceViewArrayOutput() MachineExtensionInstanceViewArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewArrayOutput) ToMachineExtensionInstanceViewArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewArrayOutput) Index(i pulumi.IntInput) MachineExtensionInstanceViewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineExtensionInstanceView {
		return vs[0].([]MachineExtensionInstanceView)[vs[1].(int)]
	}).(MachineExtensionInstanceViewOutput)
}

type MachineExtensionInstanceViewResponse struct {
	Name               *string                                     `pulumi:"name"`
	Status             *MachineExtensionInstanceViewResponseStatus `pulumi:"status"`
	Type               *string                                     `pulumi:"type"`
	TypeHandlerVersion *string                                     `pulumi:"typeHandlerVersion"`
}





type MachineExtensionInstanceViewResponseInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseOutput() MachineExtensionInstanceViewResponseOutput
	ToMachineExtensionInstanceViewResponseOutputWithContext(context.Context) MachineExtensionInstanceViewResponseOutput
}

type MachineExtensionInstanceViewResponseArgs struct {
	Name               pulumi.StringPtrInput                              `pulumi:"name"`
	Status             MachineExtensionInstanceViewResponseStatusPtrInput `pulumi:"status"`
	Type               pulumi.StringPtrInput                              `pulumi:"type"`
	TypeHandlerVersion pulumi.StringPtrInput                              `pulumi:"typeHandlerVersion"`
}

func (MachineExtensionInstanceViewResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (i MachineExtensionInstanceViewResponseArgs) ToMachineExtensionInstanceViewResponseOutput() MachineExtensionInstanceViewResponseOutput {
	return i.ToMachineExtensionInstanceViewResponseOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseArgs) ToMachineExtensionInstanceViewResponseOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseOutput)
}

func (i MachineExtensionInstanceViewResponseArgs) ToMachineExtensionInstanceViewResponsePtrOutput() MachineExtensionInstanceViewResponsePtrOutput {
	return i.ToMachineExtensionInstanceViewResponsePtrOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseArgs) ToMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseOutput).ToMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx)
}









type MachineExtensionInstanceViewResponsePtrInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponsePtrOutput() MachineExtensionInstanceViewResponsePtrOutput
	ToMachineExtensionInstanceViewResponsePtrOutputWithContext(context.Context) MachineExtensionInstanceViewResponsePtrOutput
}

type machineExtensionInstanceViewResponsePtrType MachineExtensionInstanceViewResponseArgs

func MachineExtensionInstanceViewResponsePtr(v *MachineExtensionInstanceViewResponseArgs) MachineExtensionInstanceViewResponsePtrInput {
	return (*machineExtensionInstanceViewResponsePtrType)(v)
}

func (*machineExtensionInstanceViewResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (i *machineExtensionInstanceViewResponsePtrType) ToMachineExtensionInstanceViewResponsePtrOutput() MachineExtensionInstanceViewResponsePtrOutput {
	return i.ToMachineExtensionInstanceViewResponsePtrOutputWithContext(context.Background())
}

func (i *machineExtensionInstanceViewResponsePtrType) ToMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponsePtrOutput)
}





type MachineExtensionInstanceViewResponseArrayInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseArrayOutput() MachineExtensionInstanceViewResponseArrayOutput
	ToMachineExtensionInstanceViewResponseArrayOutputWithContext(context.Context) MachineExtensionInstanceViewResponseArrayOutput
}

type MachineExtensionInstanceViewResponseArray []MachineExtensionInstanceViewResponseInput

func (MachineExtensionInstanceViewResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (i MachineExtensionInstanceViewResponseArray) ToMachineExtensionInstanceViewResponseArrayOutput() MachineExtensionInstanceViewResponseArrayOutput {
	return i.ToMachineExtensionInstanceViewResponseArrayOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseArray) ToMachineExtensionInstanceViewResponseArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseArrayOutput)
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

func (o MachineExtensionInstanceViewResponseOutput) ToMachineExtensionInstanceViewResponsePtrOutput() MachineExtensionInstanceViewResponsePtrOutput {
	return o.ToMachineExtensionInstanceViewResponsePtrOutputWithContext(context.Background())
}

func (o MachineExtensionInstanceViewResponseOutput) ToMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionInstanceViewResponse) *MachineExtensionInstanceViewResponse {
		return &v
	}).(MachineExtensionInstanceViewResponsePtrOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) *MachineExtensionInstanceViewResponseStatus {
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponsePtrOutput) ToMachineExtensionInstanceViewResponsePtrOutput() MachineExtensionInstanceViewResponsePtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponsePtrOutput) ToMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponsePtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponsePtrOutput) Elem() MachineExtensionInstanceViewResponseOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponse) MachineExtensionInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceViewResponse
		return ret
	}).(MachineExtensionInstanceViewResponseOutput)
}

func (o MachineExtensionInstanceViewResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponsePtrOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponse) *MachineExtensionInstanceViewResponseStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionInstanceViewResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponsePtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
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
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}





type MachineExtensionInstanceViewResponseStatusInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput
	ToMachineExtensionInstanceViewResponseStatusOutputWithContext(context.Context) MachineExtensionInstanceViewResponseStatusOutput
}

type MachineExtensionInstanceViewResponseStatusArgs struct {
	Code          pulumi.StringPtrInput `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput `pulumi:"displayStatus"`
	Level         pulumi.StringPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput `pulumi:"message"`
	Time          pulumi.StringPtrInput `pulumi:"time"`
}

func (MachineExtensionInstanceViewResponseStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput {
	return i.ToMachineExtensionInstanceViewResponseStatusOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseStatusOutput)
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return i.ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseStatusOutput).ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx)
}









type MachineExtensionInstanceViewResponseStatusPtrInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput
	ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput
}

type machineExtensionInstanceViewResponseStatusPtrType MachineExtensionInstanceViewResponseStatusArgs

func MachineExtensionInstanceViewResponseStatusPtr(v *MachineExtensionInstanceViewResponseStatusArgs) MachineExtensionInstanceViewResponseStatusPtrInput {
	return (*machineExtensionInstanceViewResponseStatusPtrType)(v)
}

func (*machineExtensionInstanceViewResponseStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (i *machineExtensionInstanceViewResponseStatusPtrType) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return i.ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Background())
}

func (i *machineExtensionInstanceViewResponseStatusPtrType) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseStatusPtrOutput)
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

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Background())
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionInstanceViewResponseStatus) *MachineExtensionInstanceViewResponseStatus {
		return &v
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) *string { return v.Time }).(pulumi.StringPtrOutput)
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
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewStatus struct {
	Code          *string           `pulumi:"code"`
	DisplayStatus *string           `pulumi:"displayStatus"`
	Level         *StatusLevelTypes `pulumi:"level"`
	Message       *string           `pulumi:"message"`
	Time          *string           `pulumi:"time"`
}





type MachineExtensionInstanceViewStatusInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewStatusOutput() MachineExtensionInstanceViewStatusOutput
	ToMachineExtensionInstanceViewStatusOutputWithContext(context.Context) MachineExtensionInstanceViewStatusOutput
}

type MachineExtensionInstanceViewStatusArgs struct {
	Code          pulumi.StringPtrInput    `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput    `pulumi:"displayStatus"`
	Level         StatusLevelTypesPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput    `pulumi:"message"`
	Time          pulumi.StringPtrInput    `pulumi:"time"`
}

func (MachineExtensionInstanceViewStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewStatus)(nil)).Elem()
}

func (i MachineExtensionInstanceViewStatusArgs) ToMachineExtensionInstanceViewStatusOutput() MachineExtensionInstanceViewStatusOutput {
	return i.ToMachineExtensionInstanceViewStatusOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewStatusArgs) ToMachineExtensionInstanceViewStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewStatusOutput)
}

func (i MachineExtensionInstanceViewStatusArgs) ToMachineExtensionInstanceViewStatusPtrOutput() MachineExtensionInstanceViewStatusPtrOutput {
	return i.ToMachineExtensionInstanceViewStatusPtrOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewStatusArgs) ToMachineExtensionInstanceViewStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewStatusOutput).ToMachineExtensionInstanceViewStatusPtrOutputWithContext(ctx)
}









type MachineExtensionInstanceViewStatusPtrInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewStatusPtrOutput() MachineExtensionInstanceViewStatusPtrOutput
	ToMachineExtensionInstanceViewStatusPtrOutputWithContext(context.Context) MachineExtensionInstanceViewStatusPtrOutput
}

type machineExtensionInstanceViewStatusPtrType MachineExtensionInstanceViewStatusArgs

func MachineExtensionInstanceViewStatusPtr(v *MachineExtensionInstanceViewStatusArgs) MachineExtensionInstanceViewStatusPtrInput {
	return (*machineExtensionInstanceViewStatusPtrType)(v)
}

func (*machineExtensionInstanceViewStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewStatus)(nil)).Elem()
}

func (i *machineExtensionInstanceViewStatusPtrType) ToMachineExtensionInstanceViewStatusPtrOutput() MachineExtensionInstanceViewStatusPtrOutput {
	return i.ToMachineExtensionInstanceViewStatusPtrOutputWithContext(context.Background())
}

func (i *machineExtensionInstanceViewStatusPtrType) ToMachineExtensionInstanceViewStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewStatusPtrOutput)
}

type MachineExtensionInstanceViewStatusOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewStatusOutput) ToMachineExtensionInstanceViewStatusOutput() MachineExtensionInstanceViewStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewStatusOutput) ToMachineExtensionInstanceViewStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewStatusOutput) ToMachineExtensionInstanceViewStatusPtrOutput() MachineExtensionInstanceViewStatusPtrOutput {
	return o.ToMachineExtensionInstanceViewStatusPtrOutputWithContext(context.Background())
}

func (o MachineExtensionInstanceViewStatusOutput) ToMachineExtensionInstanceViewStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionInstanceViewStatus) *MachineExtensionInstanceViewStatus {
		return &v
	}).(MachineExtensionInstanceViewStatusPtrOutput)
}

func (o MachineExtensionInstanceViewStatusOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewStatus) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewStatus) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusOutput) Level() StatusLevelTypesPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewStatus) *StatusLevelTypes { return v.Level }).(StatusLevelTypesPtrOutput)
}

func (o MachineExtensionInstanceViewStatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewStatus) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewStatus) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewStatusPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewStatusPtrOutput) ToMachineExtensionInstanceViewStatusPtrOutput() MachineExtensionInstanceViewStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewStatusPtrOutput) ToMachineExtensionInstanceViewStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewStatusPtrOutput) Elem() MachineExtensionInstanceViewStatusOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) MachineExtensionInstanceViewStatus {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceViewStatus
		return ret
	}).(MachineExtensionInstanceViewStatusOutput)
}

func (o MachineExtensionInstanceViewStatusPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusPtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusPtrOutput) Level() StatusLevelTypesPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) *StatusLevelTypes {
		if v == nil {
			return nil
		}
		return v.Level
	}).(StatusLevelTypesPtrOutput)
}

func (o MachineExtensionInstanceViewStatusPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type OSProfileResponse struct {
	ComputerName string `pulumi:"computerName"`
}





type OSProfileResponseInput interface {
	pulumi.Input

	ToOSProfileResponseOutput() OSProfileResponseOutput
	ToOSProfileResponseOutputWithContext(context.Context) OSProfileResponseOutput
}

type OSProfileResponseArgs struct {
	ComputerName pulumi.StringInput `pulumi:"computerName"`
}

func (OSProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponse)(nil)).Elem()
}

func (i OSProfileResponseArgs) ToOSProfileResponseOutput() OSProfileResponseOutput {
	return i.ToOSProfileResponseOutputWithContext(context.Background())
}

func (i OSProfileResponseArgs) ToOSProfileResponseOutputWithContext(ctx context.Context) OSProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileResponseOutput)
}

func (i OSProfileResponseArgs) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return i.ToOSProfileResponsePtrOutputWithContext(context.Background())
}

func (i OSProfileResponseArgs) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileResponseOutput).ToOSProfileResponsePtrOutputWithContext(ctx)
}









type OSProfileResponsePtrInput interface {
	pulumi.Input

	ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput
	ToOSProfileResponsePtrOutputWithContext(context.Context) OSProfileResponsePtrOutput
}

type osprofileResponsePtrType OSProfileResponseArgs

func OSProfileResponsePtr(v *OSProfileResponseArgs) OSProfileResponsePtrInput {
	return (*osprofileResponsePtrType)(v)
}

func (*osprofileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponse)(nil)).Elem()
}

func (i *osprofileResponsePtrType) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return i.ToOSProfileResponsePtrOutputWithContext(context.Background())
}

func (i *osprofileResponsePtrType) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileResponsePtrOutput)
}

type OSProfileResponseOutput struct{ *pulumi.OutputState }

func (OSProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutput() OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutputWithContext(ctx context.Context) OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return o.ToOSProfileResponsePtrOutputWithContext(context.Background())
}

func (o OSProfileResponseOutput) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSProfileResponse) *OSProfileResponse {
		return &v
	}).(OSProfileResponsePtrOutput)
}

func (o OSProfileResponseOutput) ComputerName() pulumi.StringOutput {
	return o.ApplyT(func(v OSProfileResponse) string { return v.ComputerName }).(pulumi.StringOutput)
}

type OSProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OSProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) Elem() OSProfileResponseOutput {
	return o.ApplyT(func(v *OSProfileResponse) OSProfileResponse {
		if v != nil {
			return *v
		}
		var ret OSProfileResponse
		return ret
	}).(OSProfileResponseOutput)
}

func (o OSProfileResponsePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorDetailResponseInput)(nil)).Elem(), ErrorDetailResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorDetailResponseArrayInput)(nil)).Elem(), ErrorDetailResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineExtensionInstanceViewInput)(nil)).Elem(), MachineExtensionInstanceViewArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineExtensionInstanceViewPtrInput)(nil)).Elem(), MachineExtensionInstanceViewArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineExtensionInstanceViewArrayInput)(nil)).Elem(), MachineExtensionInstanceViewArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineExtensionInstanceViewResponseInput)(nil)).Elem(), MachineExtensionInstanceViewResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineExtensionInstanceViewResponsePtrInput)(nil)).Elem(), MachineExtensionInstanceViewResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineExtensionInstanceViewResponseArrayInput)(nil)).Elem(), MachineExtensionInstanceViewResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineExtensionInstanceViewResponseStatusInput)(nil)).Elem(), MachineExtensionInstanceViewResponseStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineExtensionInstanceViewResponseStatusPtrInput)(nil)).Elem(), MachineExtensionInstanceViewResponseStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineExtensionInstanceViewStatusInput)(nil)).Elem(), MachineExtensionInstanceViewStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineExtensionInstanceViewStatusPtrInput)(nil)).Elem(), MachineExtensionInstanceViewStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OSProfileResponseInput)(nil)).Elem(), OSProfileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OSProfileResponsePtrInput)(nil)).Elem(), OSProfileResponseArgs{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewArrayOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewStatusOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewStatusPtrOutput{})
	pulumi.RegisterOutputType(OSProfileResponseOutput{})
	pulumi.RegisterOutputType(OSProfileResponsePtrOutput{})
}
