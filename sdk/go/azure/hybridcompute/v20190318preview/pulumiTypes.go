


package v20190318preview

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
	pulumi.RegisterInputType(reflect.TypeOf((*OSProfileResponseInput)(nil)).Elem(), OSProfileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OSProfileResponsePtrInput)(nil)).Elem(), OSProfileResponseArgs{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(OSProfileResponseOutput{})
	pulumi.RegisterOutputType(OSProfileResponsePtrOutput{})
}
