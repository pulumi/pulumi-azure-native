


package v20170301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistryPasswordResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type RegistryPasswordResponseOutput struct{ *pulumi.OutputState }

func (RegistryPasswordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryPasswordResponse)(nil)).Elem()
}

func (o RegistryPasswordResponseOutput) ToRegistryPasswordResponseOutput() RegistryPasswordResponseOutput {
	return o
}

func (o RegistryPasswordResponseOutput) ToRegistryPasswordResponseOutputWithContext(ctx context.Context) RegistryPasswordResponseOutput {
	return o
}

func (o RegistryPasswordResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryPasswordResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RegistryPasswordResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryPasswordResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type RegistryPasswordResponseArrayOutput struct{ *pulumi.OutputState }

func (RegistryPasswordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryPasswordResponse)(nil)).Elem()
}

func (o RegistryPasswordResponseArrayOutput) ToRegistryPasswordResponseArrayOutput() RegistryPasswordResponseArrayOutput {
	return o
}

func (o RegistryPasswordResponseArrayOutput) ToRegistryPasswordResponseArrayOutputWithContext(ctx context.Context) RegistryPasswordResponseArrayOutput {
	return o
}

func (o RegistryPasswordResponseArrayOutput) Index(i pulumi.IntInput) RegistryPasswordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryPasswordResponse {
		return vs[0].([]RegistryPasswordResponse)[vs[1].(int)]
	}).(RegistryPasswordResponseOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type StorageAccountParameters struct {
	AccessKey string `pulumi:"accessKey"`
	Name      string `pulumi:"name"`
}





type StorageAccountParametersInput interface {
	pulumi.Input

	ToStorageAccountParametersOutput() StorageAccountParametersOutput
	ToStorageAccountParametersOutputWithContext(context.Context) StorageAccountParametersOutput
}

type StorageAccountParametersArgs struct {
	AccessKey pulumi.StringInput `pulumi:"accessKey"`
	Name      pulumi.StringInput `pulumi:"name"`
}

func (StorageAccountParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountParameters)(nil)).Elem()
}

func (i StorageAccountParametersArgs) ToStorageAccountParametersOutput() StorageAccountParametersOutput {
	return i.ToStorageAccountParametersOutputWithContext(context.Background())
}

func (i StorageAccountParametersArgs) ToStorageAccountParametersOutputWithContext(ctx context.Context) StorageAccountParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountParametersOutput)
}

type StorageAccountParametersOutput struct{ *pulumi.OutputState }

func (StorageAccountParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountParameters)(nil)).Elem()
}

func (o StorageAccountParametersOutput) ToStorageAccountParametersOutput() StorageAccountParametersOutput {
	return o
}

func (o StorageAccountParametersOutput) ToStorageAccountParametersOutputWithContext(ctx context.Context) StorageAccountParametersOutput {
	return o
}

func (o StorageAccountParametersOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountParameters) string { return v.AccessKey }).(pulumi.StringOutput)
}

func (o StorageAccountParametersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountParameters) string { return v.Name }).(pulumi.StringOutput)
}

type StorageAccountPropertiesResponse struct {
	Name *string `pulumi:"name"`
}

type StorageAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutputWithContext(ctx context.Context) StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type StorageAccountPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) Elem() StorageAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) StorageAccountPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountPropertiesResponse
		return ret
	}).(StorageAccountPropertiesResponseOutput)
}

func (o StorageAccountPropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryPasswordResponseOutput{})
	pulumi.RegisterOutputType(RegistryPasswordResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountParametersOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponsePtrOutput{})
}
