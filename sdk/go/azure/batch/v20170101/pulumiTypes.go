


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationPackageResponse struct {
	Format             string `pulumi:"format"`
	Id                 string `pulumi:"id"`
	LastActivationTime string `pulumi:"lastActivationTime"`
	State              string `pulumi:"state"`
	StorageUrl         string `pulumi:"storageUrl"`
	StorageUrlExpiry   string `pulumi:"storageUrlExpiry"`
	Version            string `pulumi:"version"`
}





type ApplicationPackageResponseInput interface {
	pulumi.Input

	ToApplicationPackageResponseOutput() ApplicationPackageResponseOutput
	ToApplicationPackageResponseOutputWithContext(context.Context) ApplicationPackageResponseOutput
}

type ApplicationPackageResponseArgs struct {
	Format             pulumi.StringInput `pulumi:"format"`
	Id                 pulumi.StringInput `pulumi:"id"`
	LastActivationTime pulumi.StringInput `pulumi:"lastActivationTime"`
	State              pulumi.StringInput `pulumi:"state"`
	StorageUrl         pulumi.StringInput `pulumi:"storageUrl"`
	StorageUrlExpiry   pulumi.StringInput `pulumi:"storageUrlExpiry"`
	Version            pulumi.StringInput `pulumi:"version"`
}

func (ApplicationPackageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageResponse)(nil)).Elem()
}

func (i ApplicationPackageResponseArgs) ToApplicationPackageResponseOutput() ApplicationPackageResponseOutput {
	return i.ToApplicationPackageResponseOutputWithContext(context.Background())
}

func (i ApplicationPackageResponseArgs) ToApplicationPackageResponseOutputWithContext(ctx context.Context) ApplicationPackageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageResponseOutput)
}





type ApplicationPackageResponseArrayInput interface {
	pulumi.Input

	ToApplicationPackageResponseArrayOutput() ApplicationPackageResponseArrayOutput
	ToApplicationPackageResponseArrayOutputWithContext(context.Context) ApplicationPackageResponseArrayOutput
}

type ApplicationPackageResponseArray []ApplicationPackageResponseInput

func (ApplicationPackageResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPackageResponse)(nil)).Elem()
}

func (i ApplicationPackageResponseArray) ToApplicationPackageResponseArrayOutput() ApplicationPackageResponseArrayOutput {
	return i.ToApplicationPackageResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationPackageResponseArray) ToApplicationPackageResponseArrayOutputWithContext(ctx context.Context) ApplicationPackageResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageResponseArrayOutput)
}

type ApplicationPackageResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPackageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageResponse)(nil)).Elem()
}

func (o ApplicationPackageResponseOutput) ToApplicationPackageResponseOutput() ApplicationPackageResponseOutput {
	return o
}

func (o ApplicationPackageResponseOutput) ToApplicationPackageResponseOutputWithContext(ctx context.Context) ApplicationPackageResponseOutput {
	return o
}

func (o ApplicationPackageResponseOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.Format }).(pulumi.StringOutput)
}

func (o ApplicationPackageResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ApplicationPackageResponseOutput) LastActivationTime() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.LastActivationTime }).(pulumi.StringOutput)
}

func (o ApplicationPackageResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o ApplicationPackageResponseOutput) StorageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.StorageUrl }).(pulumi.StringOutput)
}

func (o ApplicationPackageResponseOutput) StorageUrlExpiry() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.StorageUrlExpiry }).(pulumi.StringOutput)
}

func (o ApplicationPackageResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.Version }).(pulumi.StringOutput)
}

type ApplicationPackageResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPackageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPackageResponse)(nil)).Elem()
}

func (o ApplicationPackageResponseArrayOutput) ToApplicationPackageResponseArrayOutput() ApplicationPackageResponseArrayOutput {
	return o
}

func (o ApplicationPackageResponseArrayOutput) ToApplicationPackageResponseArrayOutputWithContext(ctx context.Context) ApplicationPackageResponseArrayOutput {
	return o
}

func (o ApplicationPackageResponseArrayOutput) Index(i pulumi.IntInput) ApplicationPackageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationPackageResponse {
		return vs[0].([]ApplicationPackageResponse)[vs[1].(int)]
	}).(ApplicationPackageResponseOutput)
}

type AutoStorageBaseProperties struct {
	StorageAccountId string `pulumi:"storageAccountId"`
}





type AutoStorageBasePropertiesInput interface {
	pulumi.Input

	ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput
	ToAutoStorageBasePropertiesOutputWithContext(context.Context) AutoStorageBasePropertiesOutput
}

type AutoStorageBasePropertiesArgs struct {
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (AutoStorageBasePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStorageBaseProperties)(nil)).Elem()
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput {
	return i.ToAutoStorageBasePropertiesOutputWithContext(context.Background())
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesOutputWithContext(ctx context.Context) AutoStorageBasePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesOutput)
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return i.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesOutput).ToAutoStorageBasePropertiesPtrOutputWithContext(ctx)
}









type AutoStorageBasePropertiesPtrInput interface {
	pulumi.Input

	ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput
	ToAutoStorageBasePropertiesPtrOutputWithContext(context.Context) AutoStorageBasePropertiesPtrOutput
}

type autoStorageBasePropertiesPtrType AutoStorageBasePropertiesArgs

func AutoStorageBasePropertiesPtr(v *AutoStorageBasePropertiesArgs) AutoStorageBasePropertiesPtrInput {
	return (*autoStorageBasePropertiesPtrType)(v)
}

func (*autoStorageBasePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStorageBaseProperties)(nil)).Elem()
}

func (i *autoStorageBasePropertiesPtrType) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return i.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (i *autoStorageBasePropertiesPtrType) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesPtrOutput)
}

type AutoStorageBasePropertiesOutput struct{ *pulumi.OutputState }

func (AutoStorageBasePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStorageBaseProperties)(nil)).Elem()
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput {
	return o
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesOutputWithContext(ctx context.Context) AutoStorageBasePropertiesOutput {
	return o
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return o.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoStorageBaseProperties) *AutoStorageBaseProperties {
		return &v
	}).(AutoStorageBasePropertiesPtrOutput)
}

func (o AutoStorageBasePropertiesOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStorageBaseProperties) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type AutoStorageBasePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AutoStorageBasePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStorageBaseProperties)(nil)).Elem()
}

func (o AutoStorageBasePropertiesPtrOutput) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return o
}

func (o AutoStorageBasePropertiesPtrOutput) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return o
}

func (o AutoStorageBasePropertiesPtrOutput) Elem() AutoStorageBasePropertiesOutput {
	return o.ApplyT(func(v *AutoStorageBaseProperties) AutoStorageBaseProperties {
		if v != nil {
			return *v
		}
		var ret AutoStorageBaseProperties
		return ret
	}).(AutoStorageBasePropertiesOutput)
}

func (o AutoStorageBasePropertiesPtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStorageBaseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

type AutoStoragePropertiesResponse struct {
	LastKeySync      string `pulumi:"lastKeySync"`
	StorageAccountId string `pulumi:"storageAccountId"`
}





type AutoStoragePropertiesResponseInput interface {
	pulumi.Input

	ToAutoStoragePropertiesResponseOutput() AutoStoragePropertiesResponseOutput
	ToAutoStoragePropertiesResponseOutputWithContext(context.Context) AutoStoragePropertiesResponseOutput
}

type AutoStoragePropertiesResponseArgs struct {
	LastKeySync      pulumi.StringInput `pulumi:"lastKeySync"`
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (AutoStoragePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStoragePropertiesResponse)(nil)).Elem()
}

func (i AutoStoragePropertiesResponseArgs) ToAutoStoragePropertiesResponseOutput() AutoStoragePropertiesResponseOutput {
	return i.ToAutoStoragePropertiesResponseOutputWithContext(context.Background())
}

func (i AutoStoragePropertiesResponseArgs) ToAutoStoragePropertiesResponseOutputWithContext(ctx context.Context) AutoStoragePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStoragePropertiesResponseOutput)
}

func (i AutoStoragePropertiesResponseArgs) ToAutoStoragePropertiesResponsePtrOutput() AutoStoragePropertiesResponsePtrOutput {
	return i.ToAutoStoragePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AutoStoragePropertiesResponseArgs) ToAutoStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoStoragePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStoragePropertiesResponseOutput).ToAutoStoragePropertiesResponsePtrOutputWithContext(ctx)
}









type AutoStoragePropertiesResponsePtrInput interface {
	pulumi.Input

	ToAutoStoragePropertiesResponsePtrOutput() AutoStoragePropertiesResponsePtrOutput
	ToAutoStoragePropertiesResponsePtrOutputWithContext(context.Context) AutoStoragePropertiesResponsePtrOutput
}

type autoStoragePropertiesResponsePtrType AutoStoragePropertiesResponseArgs

func AutoStoragePropertiesResponsePtr(v *AutoStoragePropertiesResponseArgs) AutoStoragePropertiesResponsePtrInput {
	return (*autoStoragePropertiesResponsePtrType)(v)
}

func (*autoStoragePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStoragePropertiesResponse)(nil)).Elem()
}

func (i *autoStoragePropertiesResponsePtrType) ToAutoStoragePropertiesResponsePtrOutput() AutoStoragePropertiesResponsePtrOutput {
	return i.ToAutoStoragePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *autoStoragePropertiesResponsePtrType) ToAutoStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoStoragePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStoragePropertiesResponsePtrOutput)
}

type AutoStoragePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutoStoragePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStoragePropertiesResponse)(nil)).Elem()
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponseOutput() AutoStoragePropertiesResponseOutput {
	return o
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponseOutputWithContext(ctx context.Context) AutoStoragePropertiesResponseOutput {
	return o
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponsePtrOutput() AutoStoragePropertiesResponsePtrOutput {
	return o.ToAutoStoragePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoStoragePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoStoragePropertiesResponse) *AutoStoragePropertiesResponse {
		return &v
	}).(AutoStoragePropertiesResponsePtrOutput)
}

func (o AutoStoragePropertiesResponseOutput) LastKeySync() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStoragePropertiesResponse) string { return v.LastKeySync }).(pulumi.StringOutput)
}

func (o AutoStoragePropertiesResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStoragePropertiesResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type AutoStoragePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoStoragePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStoragePropertiesResponse)(nil)).Elem()
}

func (o AutoStoragePropertiesResponsePtrOutput) ToAutoStoragePropertiesResponsePtrOutput() AutoStoragePropertiesResponsePtrOutput {
	return o
}

func (o AutoStoragePropertiesResponsePtrOutput) ToAutoStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoStoragePropertiesResponsePtrOutput {
	return o
}

func (o AutoStoragePropertiesResponsePtrOutput) Elem() AutoStoragePropertiesResponseOutput {
	return o.ApplyT(func(v *AutoStoragePropertiesResponse) AutoStoragePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AutoStoragePropertiesResponse
		return ret
	}).(AutoStoragePropertiesResponseOutput)
}

func (o AutoStoragePropertiesResponsePtrOutput) LastKeySync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastKeySync
	}).(pulumi.StringPtrOutput)
}

func (o AutoStoragePropertiesResponsePtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

type KeyVaultReference struct {
	Id  string `pulumi:"id"`
	Url string `pulumi:"url"`
}





type KeyVaultReferenceInput interface {
	pulumi.Input

	ToKeyVaultReferenceOutput() KeyVaultReferenceOutput
	ToKeyVaultReferenceOutputWithContext(context.Context) KeyVaultReferenceOutput
}

type KeyVaultReferenceArgs struct {
	Id  pulumi.StringInput `pulumi:"id"`
	Url pulumi.StringInput `pulumi:"url"`
}

func (KeyVaultReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReference)(nil)).Elem()
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferenceOutput() KeyVaultReferenceOutput {
	return i.ToKeyVaultReferenceOutputWithContext(context.Background())
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferenceOutputWithContext(ctx context.Context) KeyVaultReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceOutput)
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return i.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceOutput).ToKeyVaultReferencePtrOutputWithContext(ctx)
}









type KeyVaultReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput
	ToKeyVaultReferencePtrOutputWithContext(context.Context) KeyVaultReferencePtrOutput
}

type keyVaultReferencePtrType KeyVaultReferenceArgs

func KeyVaultReferencePtr(v *KeyVaultReferenceArgs) KeyVaultReferencePtrInput {
	return (*keyVaultReferencePtrType)(v)
}

func (*keyVaultReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReference)(nil)).Elem()
}

func (i *keyVaultReferencePtrType) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return i.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultReferencePtrType) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferencePtrOutput)
}

type KeyVaultReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReference)(nil)).Elem()
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferenceOutput() KeyVaultReferenceOutput {
	return o
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferenceOutputWithContext(ctx context.Context) KeyVaultReferenceOutput {
	return o
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return o.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultReference) *KeyVaultReference {
		return &v
	}).(KeyVaultReferencePtrOutput)
}

func (o KeyVaultReferenceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReference) string { return v.Id }).(pulumi.StringOutput)
}

func (o KeyVaultReferenceOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReference) string { return v.Url }).(pulumi.StringOutput)
}

type KeyVaultReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReference)(nil)).Elem()
}

func (o KeyVaultReferencePtrOutput) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return o
}

func (o KeyVaultReferencePtrOutput) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return o
}

func (o KeyVaultReferencePtrOutput) Elem() KeyVaultReferenceOutput {
	return o.ApplyT(func(v *KeyVaultReference) KeyVaultReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultReference
		return ret
	}).(KeyVaultReferenceOutput)
}

func (o KeyVaultReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultReference) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultReferencePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultReference) *string {
		if v == nil {
			return nil
		}
		return &v.Url
	}).(pulumi.StringPtrOutput)
}

type KeyVaultReferenceResponse struct {
	Id  string `pulumi:"id"`
	Url string `pulumi:"url"`
}





type KeyVaultReferenceResponseInput interface {
	pulumi.Input

	ToKeyVaultReferenceResponseOutput() KeyVaultReferenceResponseOutput
	ToKeyVaultReferenceResponseOutputWithContext(context.Context) KeyVaultReferenceResponseOutput
}

type KeyVaultReferenceResponseArgs struct {
	Id  pulumi.StringInput `pulumi:"id"`
	Url pulumi.StringInput `pulumi:"url"`
}

func (KeyVaultReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReferenceResponse)(nil)).Elem()
}

func (i KeyVaultReferenceResponseArgs) ToKeyVaultReferenceResponseOutput() KeyVaultReferenceResponseOutput {
	return i.ToKeyVaultReferenceResponseOutputWithContext(context.Background())
}

func (i KeyVaultReferenceResponseArgs) ToKeyVaultReferenceResponseOutputWithContext(ctx context.Context) KeyVaultReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceResponseOutput)
}

func (i KeyVaultReferenceResponseArgs) ToKeyVaultReferenceResponsePtrOutput() KeyVaultReferenceResponsePtrOutput {
	return i.ToKeyVaultReferenceResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultReferenceResponseArgs) ToKeyVaultReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceResponseOutput).ToKeyVaultReferenceResponsePtrOutputWithContext(ctx)
}









type KeyVaultReferenceResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultReferenceResponsePtrOutput() KeyVaultReferenceResponsePtrOutput
	ToKeyVaultReferenceResponsePtrOutputWithContext(context.Context) KeyVaultReferenceResponsePtrOutput
}

type keyVaultReferenceResponsePtrType KeyVaultReferenceResponseArgs

func KeyVaultReferenceResponsePtr(v *KeyVaultReferenceResponseArgs) KeyVaultReferenceResponsePtrInput {
	return (*keyVaultReferenceResponsePtrType)(v)
}

func (*keyVaultReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReferenceResponse)(nil)).Elem()
}

func (i *keyVaultReferenceResponsePtrType) ToKeyVaultReferenceResponsePtrOutput() KeyVaultReferenceResponsePtrOutput {
	return i.ToKeyVaultReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultReferenceResponsePtrType) ToKeyVaultReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceResponsePtrOutput)
}

type KeyVaultReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReferenceResponse)(nil)).Elem()
}

func (o KeyVaultReferenceResponseOutput) ToKeyVaultReferenceResponseOutput() KeyVaultReferenceResponseOutput {
	return o
}

func (o KeyVaultReferenceResponseOutput) ToKeyVaultReferenceResponseOutputWithContext(ctx context.Context) KeyVaultReferenceResponseOutput {
	return o
}

func (o KeyVaultReferenceResponseOutput) ToKeyVaultReferenceResponsePtrOutput() KeyVaultReferenceResponsePtrOutput {
	return o.ToKeyVaultReferenceResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultReferenceResponseOutput) ToKeyVaultReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultReferenceResponse) *KeyVaultReferenceResponse {
		return &v
	}).(KeyVaultReferenceResponsePtrOutput)
}

func (o KeyVaultReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o KeyVaultReferenceResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReferenceResponse) string { return v.Url }).(pulumi.StringOutput)
}

type KeyVaultReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReferenceResponse)(nil)).Elem()
}

func (o KeyVaultReferenceResponsePtrOutput) ToKeyVaultReferenceResponsePtrOutput() KeyVaultReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultReferenceResponsePtrOutput) ToKeyVaultReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultReferenceResponsePtrOutput) Elem() KeyVaultReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultReferenceResponse) KeyVaultReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultReferenceResponse
		return ret
	}).(KeyVaultReferenceResponseOutput)
}

func (o KeyVaultReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultReferenceResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Url
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPackageResponseInput)(nil)).Elem(), ApplicationPackageResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPackageResponseArrayInput)(nil)).Elem(), ApplicationPackageResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoStorageBasePropertiesInput)(nil)).Elem(), AutoStorageBasePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoStorageBasePropertiesPtrInput)(nil)).Elem(), AutoStorageBasePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoStoragePropertiesResponseInput)(nil)).Elem(), AutoStoragePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoStoragePropertiesResponsePtrInput)(nil)).Elem(), AutoStoragePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultReferenceInput)(nil)).Elem(), KeyVaultReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultReferencePtrInput)(nil)).Elem(), KeyVaultReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultReferenceResponseInput)(nil)).Elem(), KeyVaultReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultReferenceResponsePtrInput)(nil)).Elem(), KeyVaultReferenceResponseArgs{})
	pulumi.RegisterOutputType(ApplicationPackageResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPackageResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoStorageBasePropertiesOutput{})
	pulumi.RegisterOutputType(AutoStorageBasePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AutoStoragePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutoStoragePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceResponsePtrOutput{})
}
