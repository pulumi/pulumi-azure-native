


package v20170501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccountProperties struct {
	AccessKey        string `pulumi:"accessKey"`
	StorageAccountId string `pulumi:"storageAccountId"`
}





type StorageAccountPropertiesInput interface {
	pulumi.Input

	ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput
	ToStorageAccountPropertiesOutputWithContext(context.Context) StorageAccountPropertiesOutput
}

type StorageAccountPropertiesArgs struct {
	AccessKey        pulumi.StringInput `pulumi:"accessKey"`
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (StorageAccountPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return i.ToStorageAccountPropertiesOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput)
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return i.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput).ToStorageAccountPropertiesPtrOutputWithContext(ctx)
}









type StorageAccountPropertiesPtrInput interface {
	pulumi.Input

	ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput
	ToStorageAccountPropertiesPtrOutputWithContext(context.Context) StorageAccountPropertiesPtrOutput
}

type storageAccountPropertiesPtrType StorageAccountPropertiesArgs

func StorageAccountPropertiesPtr(v *StorageAccountPropertiesArgs) StorageAccountPropertiesPtrInput {
	return (*storageAccountPropertiesPtrType)(v)
}

func (*storageAccountPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountProperties)(nil)).Elem()
}

func (i *storageAccountPropertiesPtrType) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return i.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i *storageAccountPropertiesPtrType) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesPtrOutput)
}

type StorageAccountPropertiesOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return o.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountProperties) *StorageAccountProperties {
		return &v
	}).(StorageAccountPropertiesPtrOutput)
}

func (o StorageAccountPropertiesOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountProperties) string { return v.AccessKey }).(pulumi.StringOutput)
}

func (o StorageAccountPropertiesOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountProperties) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type StorageAccountPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesPtrOutput) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return o
}

func (o StorageAccountPropertiesPtrOutput) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return o
}

func (o StorageAccountPropertiesPtrOutput) Elem() StorageAccountPropertiesOutput {
	return o.ApplyT(func(v *StorageAccountProperties) StorageAccountProperties {
		if v != nil {
			return *v
		}
		var ret StorageAccountProperties
		return ret
	}).(StorageAccountPropertiesOutput)
}

func (o StorageAccountPropertiesPtrOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AccessKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountPropertiesPtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountProperties) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

type StorageAccountPropertiesResponse struct {
	AccessKey        string `pulumi:"accessKey"`
	StorageAccountId string `pulumi:"storageAccountId"`
}





type StorageAccountPropertiesResponseInput interface {
	pulumi.Input

	ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput
	ToStorageAccountPropertiesResponseOutputWithContext(context.Context) StorageAccountPropertiesResponseOutput
}

type StorageAccountPropertiesResponseArgs struct {
	AccessKey        pulumi.StringInput `pulumi:"accessKey"`
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
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

func (o StorageAccountPropertiesResponseOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountPropertiesResponse) string { return v.AccessKey }).(pulumi.StringOutput)
}

func (o StorageAccountPropertiesResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountPropertiesResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
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

func (o StorageAccountPropertiesResponsePtrOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AccessKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountPropertiesResponsePtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountPropertiesInput)(nil)).Elem(), StorageAccountPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountPropertiesPtrInput)(nil)).Elem(), StorageAccountPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountPropertiesResponseInput)(nil)).Elem(), StorageAccountPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountPropertiesResponsePtrInput)(nil)).Elem(), StorageAccountPropertiesResponseArgs{})
	pulumi.RegisterOutputType(StorageAccountPropertiesOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponsePtrOutput{})
}
