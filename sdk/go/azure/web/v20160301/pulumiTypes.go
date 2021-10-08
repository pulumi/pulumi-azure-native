


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HostingEnvironmentProfileResponse struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}





type HostingEnvironmentProfileResponseInput interface {
	pulumi.Input

	ToHostingEnvironmentProfileResponseOutput() HostingEnvironmentProfileResponseOutput
	ToHostingEnvironmentProfileResponseOutputWithContext(context.Context) HostingEnvironmentProfileResponseOutput
}

type HostingEnvironmentProfileResponseArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringInput    `pulumi:"name"`
	Type pulumi.StringInput    `pulumi:"type"`
}

func (HostingEnvironmentProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (i HostingEnvironmentProfileResponseArgs) ToHostingEnvironmentProfileResponseOutput() HostingEnvironmentProfileResponseOutput {
	return i.ToHostingEnvironmentProfileResponseOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileResponseArgs) ToHostingEnvironmentProfileResponseOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileResponseOutput)
}

func (i HostingEnvironmentProfileResponseArgs) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return i.ToHostingEnvironmentProfileResponsePtrOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileResponseArgs) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileResponseOutput).ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx)
}









type HostingEnvironmentProfileResponsePtrInput interface {
	pulumi.Input

	ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput
	ToHostingEnvironmentProfileResponsePtrOutputWithContext(context.Context) HostingEnvironmentProfileResponsePtrOutput
}

type hostingEnvironmentProfileResponsePtrType HostingEnvironmentProfileResponseArgs

func HostingEnvironmentProfileResponsePtr(v *HostingEnvironmentProfileResponseArgs) HostingEnvironmentProfileResponsePtrInput {
	return (*hostingEnvironmentProfileResponsePtrType)(v)
}

func (*hostingEnvironmentProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (i *hostingEnvironmentProfileResponsePtrType) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return i.ToHostingEnvironmentProfileResponsePtrOutputWithContext(context.Background())
}

func (i *hostingEnvironmentProfileResponsePtrType) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileResponsePtrOutput)
}

type HostingEnvironmentProfileResponseOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponseOutput() HostingEnvironmentProfileResponseOutput {
	return o
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponseOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponseOutput {
	return o
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return o.ToHostingEnvironmentProfileResponsePtrOutputWithContext(context.Background())
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostingEnvironmentProfileResponse) *HostingEnvironmentProfileResponse {
		return &v
	}).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) string { return v.Type }).(pulumi.StringOutput)
}

type HostingEnvironmentProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (o HostingEnvironmentProfileResponsePtrOutput) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return o
}

func (o HostingEnvironmentProfileResponsePtrOutput) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return o
}

func (o HostingEnvironmentProfileResponsePtrOutput) Elem() HostingEnvironmentProfileResponseOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) HostingEnvironmentProfileResponse {
		if v != nil {
			return *v
		}
		var ret HostingEnvironmentProfileResponse
		return ret
	}).(HostingEnvironmentProfileResponseOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentifierResponse struct {
	Id   string  `pulumi:"id"`
	Kind *string `pulumi:"kind"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}





type IdentifierResponseInput interface {
	pulumi.Input

	ToIdentifierResponseOutput() IdentifierResponseOutput
	ToIdentifierResponseOutputWithContext(context.Context) IdentifierResponseOutput
}

type IdentifierResponseArgs struct {
	Id   pulumi.StringInput    `pulumi:"id"`
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	Name pulumi.StringInput    `pulumi:"name"`
	Type pulumi.StringInput    `pulumi:"type"`
}

func (IdentifierResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentifierResponse)(nil)).Elem()
}

func (i IdentifierResponseArgs) ToIdentifierResponseOutput() IdentifierResponseOutput {
	return i.ToIdentifierResponseOutputWithContext(context.Background())
}

func (i IdentifierResponseArgs) ToIdentifierResponseOutputWithContext(ctx context.Context) IdentifierResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentifierResponseOutput)
}





type IdentifierResponseArrayInput interface {
	pulumi.Input

	ToIdentifierResponseArrayOutput() IdentifierResponseArrayOutput
	ToIdentifierResponseArrayOutputWithContext(context.Context) IdentifierResponseArrayOutput
}

type IdentifierResponseArray []IdentifierResponseInput

func (IdentifierResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentifierResponse)(nil)).Elem()
}

func (i IdentifierResponseArray) ToIdentifierResponseArrayOutput() IdentifierResponseArrayOutput {
	return i.ToIdentifierResponseArrayOutputWithContext(context.Background())
}

func (i IdentifierResponseArray) ToIdentifierResponseArrayOutputWithContext(ctx context.Context) IdentifierResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentifierResponseArrayOutput)
}

type IdentifierResponseOutput struct{ *pulumi.OutputState }

func (IdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentifierResponse)(nil)).Elem()
}

func (o IdentifierResponseOutput) ToIdentifierResponseOutput() IdentifierResponseOutput {
	return o
}

func (o IdentifierResponseOutput) ToIdentifierResponseOutputWithContext(ctx context.Context) IdentifierResponseOutput {
	return o
}

func (o IdentifierResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IdentifierResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o IdentifierResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentifierResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o IdentifierResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IdentifierResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o IdentifierResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentifierResponse) string { return v.Type }).(pulumi.StringOutput)
}

type IdentifierResponseArrayOutput struct{ *pulumi.OutputState }

func (IdentifierResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentifierResponse)(nil)).Elem()
}

func (o IdentifierResponseArrayOutput) ToIdentifierResponseArrayOutput() IdentifierResponseArrayOutput {
	return o
}

func (o IdentifierResponseArrayOutput) ToIdentifierResponseArrayOutputWithContext(ctx context.Context) IdentifierResponseArrayOutput {
	return o
}

func (o IdentifierResponseArrayOutput) Index(i pulumi.IntInput) IdentifierResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentifierResponse {
		return vs[0].([]IdentifierResponse)[vs[1].(int)]
	}).(IdentifierResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostingEnvironmentProfileResponseInput)(nil)).Elem(), HostingEnvironmentProfileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostingEnvironmentProfileResponsePtrInput)(nil)).Elem(), HostingEnvironmentProfileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentifierResponseInput)(nil)).Elem(), IdentifierResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentifierResponseArrayInput)(nil)).Elem(), IdentifierResponseArray{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileResponseOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentifierResponseOutput{})
	pulumi.RegisterOutputType(IdentifierResponseArrayOutput{})
}
