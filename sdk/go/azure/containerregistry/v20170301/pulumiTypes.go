


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





type RegistryPasswordResponseInput interface {
	pulumi.Input

	ToRegistryPasswordResponseOutput() RegistryPasswordResponseOutput
	ToRegistryPasswordResponseOutputWithContext(context.Context) RegistryPasswordResponseOutput
}

type RegistryPasswordResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (RegistryPasswordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryPasswordResponse)(nil)).Elem()
}

func (i RegistryPasswordResponseArgs) ToRegistryPasswordResponseOutput() RegistryPasswordResponseOutput {
	return i.ToRegistryPasswordResponseOutputWithContext(context.Background())
}

func (i RegistryPasswordResponseArgs) ToRegistryPasswordResponseOutputWithContext(ctx context.Context) RegistryPasswordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPasswordResponseOutput)
}





type RegistryPasswordResponseArrayInput interface {
	pulumi.Input

	ToRegistryPasswordResponseArrayOutput() RegistryPasswordResponseArrayOutput
	ToRegistryPasswordResponseArrayOutputWithContext(context.Context) RegistryPasswordResponseArrayOutput
}

type RegistryPasswordResponseArray []RegistryPasswordResponseInput

func (RegistryPasswordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryPasswordResponse)(nil)).Elem()
}

func (i RegistryPasswordResponseArray) ToRegistryPasswordResponseArrayOutput() RegistryPasswordResponseArrayOutput {
	return i.ToRegistryPasswordResponseArrayOutputWithContext(context.Background())
}

func (i RegistryPasswordResponseArray) ToRegistryPasswordResponseArrayOutputWithContext(ctx context.Context) RegistryPasswordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPasswordResponseArrayOutput)
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

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
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

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
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

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
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

func (i StorageAccountParametersArgs) ToStorageAccountParametersPtrOutput() StorageAccountParametersPtrOutput {
	return i.ToStorageAccountParametersPtrOutputWithContext(context.Background())
}

func (i StorageAccountParametersArgs) ToStorageAccountParametersPtrOutputWithContext(ctx context.Context) StorageAccountParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountParametersOutput).ToStorageAccountParametersPtrOutputWithContext(ctx)
}









type StorageAccountParametersPtrInput interface {
	pulumi.Input

	ToStorageAccountParametersPtrOutput() StorageAccountParametersPtrOutput
	ToStorageAccountParametersPtrOutputWithContext(context.Context) StorageAccountParametersPtrOutput
}

type storageAccountParametersPtrType StorageAccountParametersArgs

func StorageAccountParametersPtr(v *StorageAccountParametersArgs) StorageAccountParametersPtrInput {
	return (*storageAccountParametersPtrType)(v)
}

func (*storageAccountParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountParameters)(nil)).Elem()
}

func (i *storageAccountParametersPtrType) ToStorageAccountParametersPtrOutput() StorageAccountParametersPtrOutput {
	return i.ToStorageAccountParametersPtrOutputWithContext(context.Background())
}

func (i *storageAccountParametersPtrType) ToStorageAccountParametersPtrOutputWithContext(ctx context.Context) StorageAccountParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountParametersPtrOutput)
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

func (o StorageAccountParametersOutput) ToStorageAccountParametersPtrOutput() StorageAccountParametersPtrOutput {
	return o.ToStorageAccountParametersPtrOutputWithContext(context.Background())
}

func (o StorageAccountParametersOutput) ToStorageAccountParametersPtrOutputWithContext(ctx context.Context) StorageAccountParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountParameters) *StorageAccountParameters {
		return &v
	}).(StorageAccountParametersPtrOutput)
}

func (o StorageAccountParametersOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountParameters) string { return v.AccessKey }).(pulumi.StringOutput)
}

func (o StorageAccountParametersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountParameters) string { return v.Name }).(pulumi.StringOutput)
}

type StorageAccountParametersPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountParameters)(nil)).Elem()
}

func (o StorageAccountParametersPtrOutput) ToStorageAccountParametersPtrOutput() StorageAccountParametersPtrOutput {
	return o
}

func (o StorageAccountParametersPtrOutput) ToStorageAccountParametersPtrOutputWithContext(ctx context.Context) StorageAccountParametersPtrOutput {
	return o
}

func (o StorageAccountParametersPtrOutput) Elem() StorageAccountParametersOutput {
	return o.ApplyT(func(v *StorageAccountParameters) StorageAccountParameters {
		if v != nil {
			return *v
		}
		var ret StorageAccountParameters
		return ret
	}).(StorageAccountParametersOutput)
}

func (o StorageAccountParametersPtrOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountParameters) *string {
		if v == nil {
			return nil
		}
		return &v.AccessKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountParametersPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountParameters) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type StorageAccountPropertiesResponse struct {
	Name *string `pulumi:"name"`
}





type StorageAccountPropertiesResponseInput interface {
	pulumi.Input

	ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput
	ToStorageAccountPropertiesResponseOutputWithContext(context.Context) StorageAccountPropertiesResponseOutput
}

type StorageAccountPropertiesResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (StorageAccountPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountPropertiesResponse)(nil)).Elem()
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput {
	return i.ToStorageAccountPropertiesResponseOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponseOutputWithContext(ctx context.Context) StorageAccountPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesResponseOutput)
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return i.ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesResponseOutput).ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx)
}









type StorageAccountPropertiesResponsePtrInput interface {
	pulumi.Input

	ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput
	ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Context) StorageAccountPropertiesResponsePtrOutput
}

type storageAccountPropertiesResponsePtrType StorageAccountPropertiesResponseArgs

func StorageAccountPropertiesResponsePtr(v *StorageAccountPropertiesResponseArgs) StorageAccountPropertiesResponsePtrInput {
	return (*storageAccountPropertiesResponsePtrType)(v)
}

func (*storageAccountPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountPropertiesResponse)(nil)).Elem()
}

func (i *storageAccountPropertiesResponsePtrType) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return i.ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *storageAccountPropertiesResponsePtrType) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesResponsePtrOutput)
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

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return o.ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountPropertiesResponse) *StorageAccountPropertiesResponse {
		return &v
	}).(StorageAccountPropertiesResponsePtrOutput)
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
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountParametersOutput{})
	pulumi.RegisterOutputType(StorageAccountParametersPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponsePtrOutput{})
}
