


package v20151001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiEndpointResponse struct {
	Endpoint     *string `pulumi:"endpoint"`
	MajorVersion *string `pulumi:"majorVersion"`
}





type ApiEndpointResponseInput interface {
	pulumi.Input

	ToApiEndpointResponseOutput() ApiEndpointResponseOutput
	ToApiEndpointResponseOutputWithContext(context.Context) ApiEndpointResponseOutput
}

type ApiEndpointResponseArgs struct {
	Endpoint     pulumi.StringPtrInput `pulumi:"endpoint"`
	MajorVersion pulumi.StringPtrInput `pulumi:"majorVersion"`
}

func (ApiEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEndpointResponse)(nil)).Elem()
}

func (i ApiEndpointResponseArgs) ToApiEndpointResponseOutput() ApiEndpointResponseOutput {
	return i.ToApiEndpointResponseOutputWithContext(context.Background())
}

func (i ApiEndpointResponseArgs) ToApiEndpointResponseOutputWithContext(ctx context.Context) ApiEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEndpointResponseOutput)
}





type ApiEndpointResponseArrayInput interface {
	pulumi.Input

	ToApiEndpointResponseArrayOutput() ApiEndpointResponseArrayOutput
	ToApiEndpointResponseArrayOutputWithContext(context.Context) ApiEndpointResponseArrayOutput
}

type ApiEndpointResponseArray []ApiEndpointResponseInput

func (ApiEndpointResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiEndpointResponse)(nil)).Elem()
}

func (i ApiEndpointResponseArray) ToApiEndpointResponseArrayOutput() ApiEndpointResponseArrayOutput {
	return i.ToApiEndpointResponseArrayOutputWithContext(context.Background())
}

func (i ApiEndpointResponseArray) ToApiEndpointResponseArrayOutputWithContext(ctx context.Context) ApiEndpointResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEndpointResponseArrayOutput)
}

type ApiEndpointResponseOutput struct{ *pulumi.OutputState }

func (ApiEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEndpointResponse)(nil)).Elem()
}

func (o ApiEndpointResponseOutput) ToApiEndpointResponseOutput() ApiEndpointResponseOutput {
	return o
}

func (o ApiEndpointResponseOutput) ToApiEndpointResponseOutputWithContext(ctx context.Context) ApiEndpointResponseOutput {
	return o
}

func (o ApiEndpointResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEndpointResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o ApiEndpointResponseOutput) MajorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEndpointResponse) *string { return v.MajorVersion }).(pulumi.StringPtrOutput)
}

type ApiEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiEndpointResponse)(nil)).Elem()
}

func (o ApiEndpointResponseArrayOutput) ToApiEndpointResponseArrayOutput() ApiEndpointResponseArrayOutput {
	return o
}

func (o ApiEndpointResponseArrayOutput) ToApiEndpointResponseArrayOutputWithContext(ctx context.Context) ApiEndpointResponseArrayOutput {
	return o
}

func (o ApiEndpointResponseArrayOutput) Index(i pulumi.IntInput) ApiEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiEndpointResponse {
		return vs[0].([]ApiEndpointResponse)[vs[1].(int)]
	}).(ApiEndpointResponseOutput)
}

type StorageAccount struct {
	Id        string `pulumi:"id"`
	IsPrimary bool   `pulumi:"isPrimary"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Id        pulumi.StringInput `pulumi:"id"`
	IsPrimary pulumi.BoolInput   `pulumi:"isPrimary"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}





type StorageAccountArrayInput interface {
	pulumi.Input

	ToStorageAccountArrayOutput() StorageAccountArrayOutput
	ToStorageAccountArrayOutputWithContext(context.Context) StorageAccountArrayOutput
}

type StorageAccountArray []StorageAccountInput

func (StorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (i StorageAccountArray) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return i.ToStorageAccountArrayOutputWithContext(context.Background())
}

func (i StorageAccountArray) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountArrayOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v StorageAccount) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

type StorageAccountArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) Index(i pulumi.IntInput) StorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccount {
		return vs[0].([]StorageAccount)[vs[1].(int)]
	}).(StorageAccountOutput)
}

type StorageAccountResponse struct {
	Id        string `pulumi:"id"`
	IsPrimary bool   `pulumi:"isPrimary"`
}





type StorageAccountResponseInput interface {
	pulumi.Input

	ToStorageAccountResponseOutput() StorageAccountResponseOutput
	ToStorageAccountResponseOutputWithContext(context.Context) StorageAccountResponseOutput
}

type StorageAccountResponseArgs struct {
	Id        pulumi.StringInput `pulumi:"id"`
	IsPrimary pulumi.BoolInput   `pulumi:"isPrimary"`
}

func (StorageAccountResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (i StorageAccountResponseArgs) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return i.ToStorageAccountResponseOutputWithContext(context.Background())
}

func (i StorageAccountResponseArgs) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponseOutput)
}





type StorageAccountResponseArrayInput interface {
	pulumi.Input

	ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput
	ToStorageAccountResponseArrayOutputWithContext(context.Context) StorageAccountResponseArrayOutput
}

type StorageAccountResponseArray []StorageAccountResponseInput

func (StorageAccountResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (i StorageAccountResponseArray) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return i.ToStorageAccountResponseArrayOutputWithContext(context.Background())
}

func (i StorageAccountResponseArray) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponseArrayOutput)
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountResponseOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v StorageAccountResponse) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

type StorageAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountResponse {
		return vs[0].([]StorageAccountResponse)[vs[1].(int)]
	}).(StorageAccountResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiEndpointResponseInput)(nil)).Elem(), ApiEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiEndpointResponseArrayInput)(nil)).Elem(), ApiEndpointResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountInput)(nil)).Elem(), StorageAccountArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountArrayInput)(nil)).Elem(), StorageAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountResponseInput)(nil)).Elem(), StorageAccountResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountResponseArrayInput)(nil)).Elem(), StorageAccountResponseArray{})
	pulumi.RegisterOutputType(ApiEndpointResponseOutput{})
	pulumi.RegisterOutputType(ApiEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseArrayOutput{})
}
