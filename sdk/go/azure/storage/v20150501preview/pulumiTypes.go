


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomDomainResponse struct {
	Name             *string `pulumi:"name"`
	UseSubDomainName *bool   `pulumi:"useSubDomainName"`
}





type CustomDomainResponseInput interface {
	pulumi.Input

	ToCustomDomainResponseOutput() CustomDomainResponseOutput
	ToCustomDomainResponseOutputWithContext(context.Context) CustomDomainResponseOutput
}

type CustomDomainResponseArgs struct {
	Name             pulumi.StringPtrInput `pulumi:"name"`
	UseSubDomainName pulumi.BoolPtrInput   `pulumi:"useSubDomainName"`
}

func (CustomDomainResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainResponse)(nil)).Elem()
}

func (i CustomDomainResponseArgs) ToCustomDomainResponseOutput() CustomDomainResponseOutput {
	return i.ToCustomDomainResponseOutputWithContext(context.Background())
}

func (i CustomDomainResponseArgs) ToCustomDomainResponseOutputWithContext(ctx context.Context) CustomDomainResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainResponseOutput)
}

func (i CustomDomainResponseArgs) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return i.ToCustomDomainResponsePtrOutputWithContext(context.Background())
}

func (i CustomDomainResponseArgs) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainResponseOutput).ToCustomDomainResponsePtrOutputWithContext(ctx)
}









type CustomDomainResponsePtrInput interface {
	pulumi.Input

	ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput
	ToCustomDomainResponsePtrOutputWithContext(context.Context) CustomDomainResponsePtrOutput
}

type customDomainResponsePtrType CustomDomainResponseArgs

func CustomDomainResponsePtr(v *CustomDomainResponseArgs) CustomDomainResponsePtrInput {
	return (*customDomainResponsePtrType)(v)
}

func (*customDomainResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainResponse)(nil)).Elem()
}

func (i *customDomainResponsePtrType) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return i.ToCustomDomainResponsePtrOutputWithContext(context.Background())
}

func (i *customDomainResponsePtrType) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainResponsePtrOutput)
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

func (o CustomDomainResponseOutput) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return o.ToCustomDomainResponsePtrOutputWithContext(context.Background())
}

func (o CustomDomainResponseOutput) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomainResponse) *CustomDomainResponse {
		return &v
	}).(CustomDomainResponsePtrOutput)
}

func (o CustomDomainResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CustomDomainResponseOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type CustomDomainResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomDomainResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponsePtrOutput) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return o
}

func (o CustomDomainResponsePtrOutput) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return o
}

func (o CustomDomainResponsePtrOutput) Elem() CustomDomainResponseOutput {
	return o.ApplyT(func(v *CustomDomainResponse) CustomDomainResponse {
		if v != nil {
			return *v
		}
		var ret CustomDomainResponse
		return ret
	}).(CustomDomainResponseOutput)
}

func (o CustomDomainResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainResponsePtrOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomainResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseSubDomainName
	}).(pulumi.BoolPtrOutput)
}

type EndpointsResponse struct {
	Blob  *string `pulumi:"blob"`
	Queue *string `pulumi:"queue"`
	Table *string `pulumi:"table"`
}





type EndpointsResponseInput interface {
	pulumi.Input

	ToEndpointsResponseOutput() EndpointsResponseOutput
	ToEndpointsResponseOutputWithContext(context.Context) EndpointsResponseOutput
}

type EndpointsResponseArgs struct {
	Blob  pulumi.StringPtrInput `pulumi:"blob"`
	Queue pulumi.StringPtrInput `pulumi:"queue"`
	Table pulumi.StringPtrInput `pulumi:"table"`
}

func (EndpointsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (i EndpointsResponseArgs) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return i.ToEndpointsResponseOutputWithContext(context.Background())
}

func (i EndpointsResponseArgs) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponseOutput)
}

func (i EndpointsResponseArgs) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return i.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i EndpointsResponseArgs) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponseOutput).ToEndpointsResponsePtrOutputWithContext(ctx)
}









type EndpointsResponsePtrInput interface {
	pulumi.Input

	ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput
	ToEndpointsResponsePtrOutputWithContext(context.Context) EndpointsResponsePtrOutput
}

type endpointsResponsePtrType EndpointsResponseArgs

func EndpointsResponsePtr(v *EndpointsResponseArgs) EndpointsResponsePtrInput {
	return (*endpointsResponsePtrType)(v)
}

func (*endpointsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsResponse)(nil)).Elem()
}

func (i *endpointsResponsePtrType) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return i.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i *endpointsResponsePtrType) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponsePtrOutput)
}

type EndpointsResponseOutput struct{ *pulumi.OutputState }

func (EndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (o EndpointsResponseOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointsResponse) *EndpointsResponse {
		return &v
	}).(EndpointsResponsePtrOutput)
}

func (o EndpointsResponseOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointsResponse) *string { return v.Blob }).(pulumi.StringPtrOutput)
}

func (o EndpointsResponseOutput) Queue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointsResponse) *string { return v.Queue }).(pulumi.StringPtrOutput)
}

func (o EndpointsResponseOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointsResponse) *string { return v.Table }).(pulumi.StringPtrOutput)
}

type EndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (EndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) Elem() EndpointsResponseOutput {
	return o.ApplyT(func(v *EndpointsResponse) EndpointsResponse {
		if v != nil {
			return *v
		}
		var ret EndpointsResponse
		return ret
	}).(EndpointsResponseOutput)
}

func (o EndpointsResponsePtrOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Queue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Queue
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Table
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomDomainResponseOutput{})
	pulumi.RegisterOutputType(CustomDomainResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointsResponseOutput{})
	pulumi.RegisterOutputType(EndpointsResponsePtrOutput{})
}
