


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataPoolEncryption struct {
	KeyName              string  `pulumi:"keyName"`
	KeyVaultUri          string  `pulumi:"keyVaultUri"`
	KeyVersion           *string `pulumi:"keyVersion"`
	UserAssignedIdentity string  `pulumi:"userAssignedIdentity"`
}





type DataPoolEncryptionInput interface {
	pulumi.Input

	ToDataPoolEncryptionOutput() DataPoolEncryptionOutput
	ToDataPoolEncryptionOutputWithContext(context.Context) DataPoolEncryptionOutput
}

type DataPoolEncryptionArgs struct {
	KeyName              pulumi.StringInput    `pulumi:"keyName"`
	KeyVaultUri          pulumi.StringInput    `pulumi:"keyVaultUri"`
	KeyVersion           pulumi.StringPtrInput `pulumi:"keyVersion"`
	UserAssignedIdentity pulumi.StringInput    `pulumi:"userAssignedIdentity"`
}

func (DataPoolEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPoolEncryption)(nil)).Elem()
}

func (i DataPoolEncryptionArgs) ToDataPoolEncryptionOutput() DataPoolEncryptionOutput {
	return i.ToDataPoolEncryptionOutputWithContext(context.Background())
}

func (i DataPoolEncryptionArgs) ToDataPoolEncryptionOutputWithContext(ctx context.Context) DataPoolEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPoolEncryptionOutput)
}

func (i DataPoolEncryptionArgs) ToDataPoolEncryptionPtrOutput() DataPoolEncryptionPtrOutput {
	return i.ToDataPoolEncryptionPtrOutputWithContext(context.Background())
}

func (i DataPoolEncryptionArgs) ToDataPoolEncryptionPtrOutputWithContext(ctx context.Context) DataPoolEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPoolEncryptionOutput).ToDataPoolEncryptionPtrOutputWithContext(ctx)
}









type DataPoolEncryptionPtrInput interface {
	pulumi.Input

	ToDataPoolEncryptionPtrOutput() DataPoolEncryptionPtrOutput
	ToDataPoolEncryptionPtrOutputWithContext(context.Context) DataPoolEncryptionPtrOutput
}

type dataPoolEncryptionPtrType DataPoolEncryptionArgs

func DataPoolEncryptionPtr(v *DataPoolEncryptionArgs) DataPoolEncryptionPtrInput {
	return (*dataPoolEncryptionPtrType)(v)
}

func (*dataPoolEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPoolEncryption)(nil)).Elem()
}

func (i *dataPoolEncryptionPtrType) ToDataPoolEncryptionPtrOutput() DataPoolEncryptionPtrOutput {
	return i.ToDataPoolEncryptionPtrOutputWithContext(context.Background())
}

func (i *dataPoolEncryptionPtrType) ToDataPoolEncryptionPtrOutputWithContext(ctx context.Context) DataPoolEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPoolEncryptionPtrOutput)
}

type DataPoolEncryptionOutput struct{ *pulumi.OutputState }

func (DataPoolEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPoolEncryption)(nil)).Elem()
}

func (o DataPoolEncryptionOutput) ToDataPoolEncryptionOutput() DataPoolEncryptionOutput {
	return o
}

func (o DataPoolEncryptionOutput) ToDataPoolEncryptionOutputWithContext(ctx context.Context) DataPoolEncryptionOutput {
	return o
}

func (o DataPoolEncryptionOutput) ToDataPoolEncryptionPtrOutput() DataPoolEncryptionPtrOutput {
	return o.ToDataPoolEncryptionPtrOutputWithContext(context.Background())
}

func (o DataPoolEncryptionOutput) ToDataPoolEncryptionPtrOutputWithContext(ctx context.Context) DataPoolEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataPoolEncryption) *DataPoolEncryption {
		return &v
	}).(DataPoolEncryptionPtrOutput)
}

func (o DataPoolEncryptionOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v DataPoolEncryption) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o DataPoolEncryptionOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v DataPoolEncryption) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o DataPoolEncryptionOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataPoolEncryption) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o DataPoolEncryptionOutput) UserAssignedIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v DataPoolEncryption) string { return v.UserAssignedIdentity }).(pulumi.StringOutput)
}

type DataPoolEncryptionPtrOutput struct{ *pulumi.OutputState }

func (DataPoolEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPoolEncryption)(nil)).Elem()
}

func (o DataPoolEncryptionPtrOutput) ToDataPoolEncryptionPtrOutput() DataPoolEncryptionPtrOutput {
	return o
}

func (o DataPoolEncryptionPtrOutput) ToDataPoolEncryptionPtrOutputWithContext(ctx context.Context) DataPoolEncryptionPtrOutput {
	return o
}

func (o DataPoolEncryptionPtrOutput) Elem() DataPoolEncryptionOutput {
	return o.ApplyT(func(v *DataPoolEncryption) DataPoolEncryption {
		if v != nil {
			return *v
		}
		var ret DataPoolEncryption
		return ret
	}).(DataPoolEncryptionOutput)
}

func (o DataPoolEncryptionPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataPoolEncryption) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o DataPoolEncryptionPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataPoolEncryption) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o DataPoolEncryptionPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataPoolEncryption) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o DataPoolEncryptionPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataPoolEncryption) *string {
		if v == nil {
			return nil
		}
		return &v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type DataPoolEncryptionResponse struct {
	KeyName              string  `pulumi:"keyName"`
	KeyVaultUri          string  `pulumi:"keyVaultUri"`
	KeyVersion           *string `pulumi:"keyVersion"`
	UserAssignedIdentity string  `pulumi:"userAssignedIdentity"`
}

type DataPoolEncryptionResponseOutput struct{ *pulumi.OutputState }

func (DataPoolEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPoolEncryptionResponse)(nil)).Elem()
}

func (o DataPoolEncryptionResponseOutput) ToDataPoolEncryptionResponseOutput() DataPoolEncryptionResponseOutput {
	return o
}

func (o DataPoolEncryptionResponseOutput) ToDataPoolEncryptionResponseOutputWithContext(ctx context.Context) DataPoolEncryptionResponseOutput {
	return o
}

func (o DataPoolEncryptionResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v DataPoolEncryptionResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o DataPoolEncryptionResponseOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v DataPoolEncryptionResponse) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o DataPoolEncryptionResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataPoolEncryptionResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o DataPoolEncryptionResponseOutput) UserAssignedIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v DataPoolEncryptionResponse) string { return v.UserAssignedIdentity }).(pulumi.StringOutput)
}

type DataPoolEncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (DataPoolEncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPoolEncryptionResponse)(nil)).Elem()
}

func (o DataPoolEncryptionResponsePtrOutput) ToDataPoolEncryptionResponsePtrOutput() DataPoolEncryptionResponsePtrOutput {
	return o
}

func (o DataPoolEncryptionResponsePtrOutput) ToDataPoolEncryptionResponsePtrOutputWithContext(ctx context.Context) DataPoolEncryptionResponsePtrOutput {
	return o
}

func (o DataPoolEncryptionResponsePtrOutput) Elem() DataPoolEncryptionResponseOutput {
	return o.ApplyT(func(v *DataPoolEncryptionResponse) DataPoolEncryptionResponse {
		if v != nil {
			return *v
		}
		var ret DataPoolEncryptionResponse
		return ret
	}).(DataPoolEncryptionResponseOutput)
}

func (o DataPoolEncryptionResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataPoolEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o DataPoolEncryptionResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataPoolEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o DataPoolEncryptionResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataPoolEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o DataPoolEncryptionResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataPoolEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type DataPoolLocation struct {
	Encryption *DataPoolEncryption `pulumi:"encryption"`
	Name       string              `pulumi:"name"`
}





type DataPoolLocationInput interface {
	pulumi.Input

	ToDataPoolLocationOutput() DataPoolLocationOutput
	ToDataPoolLocationOutputWithContext(context.Context) DataPoolLocationOutput
}

type DataPoolLocationArgs struct {
	Encryption DataPoolEncryptionPtrInput `pulumi:"encryption"`
	Name       pulumi.StringInput         `pulumi:"name"`
}

func (DataPoolLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPoolLocation)(nil)).Elem()
}

func (i DataPoolLocationArgs) ToDataPoolLocationOutput() DataPoolLocationOutput {
	return i.ToDataPoolLocationOutputWithContext(context.Background())
}

func (i DataPoolLocationArgs) ToDataPoolLocationOutputWithContext(ctx context.Context) DataPoolLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPoolLocationOutput)
}





type DataPoolLocationArrayInput interface {
	pulumi.Input

	ToDataPoolLocationArrayOutput() DataPoolLocationArrayOutput
	ToDataPoolLocationArrayOutputWithContext(context.Context) DataPoolLocationArrayOutput
}

type DataPoolLocationArray []DataPoolLocationInput

func (DataPoolLocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPoolLocation)(nil)).Elem()
}

func (i DataPoolLocationArray) ToDataPoolLocationArrayOutput() DataPoolLocationArrayOutput {
	return i.ToDataPoolLocationArrayOutputWithContext(context.Background())
}

func (i DataPoolLocationArray) ToDataPoolLocationArrayOutputWithContext(ctx context.Context) DataPoolLocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPoolLocationArrayOutput)
}

type DataPoolLocationOutput struct{ *pulumi.OutputState }

func (DataPoolLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPoolLocation)(nil)).Elem()
}

func (o DataPoolLocationOutput) ToDataPoolLocationOutput() DataPoolLocationOutput {
	return o
}

func (o DataPoolLocationOutput) ToDataPoolLocationOutputWithContext(ctx context.Context) DataPoolLocationOutput {
	return o
}

func (o DataPoolLocationOutput) Encryption() DataPoolEncryptionPtrOutput {
	return o.ApplyT(func(v DataPoolLocation) *DataPoolEncryption { return v.Encryption }).(DataPoolEncryptionPtrOutput)
}

func (o DataPoolLocationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataPoolLocation) string { return v.Name }).(pulumi.StringOutput)
}

type DataPoolLocationArrayOutput struct{ *pulumi.OutputState }

func (DataPoolLocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPoolLocation)(nil)).Elem()
}

func (o DataPoolLocationArrayOutput) ToDataPoolLocationArrayOutput() DataPoolLocationArrayOutput {
	return o
}

func (o DataPoolLocationArrayOutput) ToDataPoolLocationArrayOutputWithContext(ctx context.Context) DataPoolLocationArrayOutput {
	return o
}

func (o DataPoolLocationArrayOutput) Index(i pulumi.IntInput) DataPoolLocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataPoolLocation {
		return vs[0].([]DataPoolLocation)[vs[1].(int)]
	}).(DataPoolLocationOutput)
}

type DataPoolLocationResponse struct {
	Encryption *DataPoolEncryptionResponse `pulumi:"encryption"`
	Name       string                      `pulumi:"name"`
}

type DataPoolLocationResponseOutput struct{ *pulumi.OutputState }

func (DataPoolLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPoolLocationResponse)(nil)).Elem()
}

func (o DataPoolLocationResponseOutput) ToDataPoolLocationResponseOutput() DataPoolLocationResponseOutput {
	return o
}

func (o DataPoolLocationResponseOutput) ToDataPoolLocationResponseOutputWithContext(ctx context.Context) DataPoolLocationResponseOutput {
	return o
}

func (o DataPoolLocationResponseOutput) Encryption() DataPoolEncryptionResponsePtrOutput {
	return o.ApplyT(func(v DataPoolLocationResponse) *DataPoolEncryptionResponse { return v.Encryption }).(DataPoolEncryptionResponsePtrOutput)
}

func (o DataPoolLocationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataPoolLocationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type DataPoolLocationResponseArrayOutput struct{ *pulumi.OutputState }

func (DataPoolLocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPoolLocationResponse)(nil)).Elem()
}

func (o DataPoolLocationResponseArrayOutput) ToDataPoolLocationResponseArrayOutput() DataPoolLocationResponseArrayOutput {
	return o
}

func (o DataPoolLocationResponseArrayOutput) ToDataPoolLocationResponseArrayOutputWithContext(ctx context.Context) DataPoolLocationResponseArrayOutput {
	return o
}

func (o DataPoolLocationResponseArrayOutput) Index(i pulumi.IntInput) DataPoolLocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataPoolLocationResponse {
		return vs[0].([]DataPoolLocationResponse)[vs[1].(int)]
	}).(DataPoolLocationResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DataPoolEncryptionOutput{})
	pulumi.RegisterOutputType(DataPoolEncryptionPtrOutput{})
	pulumi.RegisterOutputType(DataPoolEncryptionResponseOutput{})
	pulumi.RegisterOutputType(DataPoolEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(DataPoolLocationOutput{})
	pulumi.RegisterOutputType(DataPoolLocationArrayOutput{})
	pulumi.RegisterOutputType(DataPoolLocationResponseOutput{})
	pulumi.RegisterOutputType(DataPoolLocationResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
