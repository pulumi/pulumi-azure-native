


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageSkuName string

const (
	StorageSkuName_Standard_LRS    = StorageSkuName("Standard_LRS")
	StorageSkuName_Standard_GRS    = StorageSkuName("Standard_GRS")
	StorageSkuName_Standard_Ragrs  = StorageSkuName("Standard_Ragrs")
	StorageSkuName_Standard_ZRS    = StorageSkuName("Standard_ZRS")
	StorageSkuName_Premium_LRS     = StorageSkuName("Premium_LRS")
	StorageSkuName_Premium_ZRS     = StorageSkuName("Premium_ZRS")
	StorageSkuName_Standard_Gzrs   = StorageSkuName("Standard_Gzrs")
	StorageSkuName_Standard_Ragzrs = StorageSkuName("Standard_Ragzrs")
)

func (StorageSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSkuName)(nil)).Elem()
}

func (e StorageSkuName) ToStorageSkuNameOutput() StorageSkuNameOutput {
	return pulumi.ToOutput(e).(StorageSkuNameOutput)
}

func (e StorageSkuName) ToStorageSkuNameOutputWithContext(ctx context.Context) StorageSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageSkuNameOutput)
}

func (e StorageSkuName) ToStorageSkuNamePtrOutput() StorageSkuNamePtrOutput {
	return e.ToStorageSkuNamePtrOutputWithContext(context.Background())
}

func (e StorageSkuName) ToStorageSkuNamePtrOutputWithContext(ctx context.Context) StorageSkuNamePtrOutput {
	return StorageSkuName(e).ToStorageSkuNameOutputWithContext(ctx).ToStorageSkuNamePtrOutputWithContext(ctx)
}

func (e StorageSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageSkuNameOutput struct{ *pulumi.OutputState }

func (StorageSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSkuName)(nil)).Elem()
}

func (o StorageSkuNameOutput) ToStorageSkuNameOutput() StorageSkuNameOutput {
	return o
}

func (o StorageSkuNameOutput) ToStorageSkuNameOutputWithContext(ctx context.Context) StorageSkuNameOutput {
	return o
}

func (o StorageSkuNameOutput) ToStorageSkuNamePtrOutput() StorageSkuNamePtrOutput {
	return o.ToStorageSkuNamePtrOutputWithContext(context.Background())
}

func (o StorageSkuNameOutput) ToStorageSkuNamePtrOutputWithContext(ctx context.Context) StorageSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageSkuName) *StorageSkuName {
		return &v
	}).(StorageSkuNamePtrOutput)
}

func (o StorageSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageSkuNamePtrOutput struct{ *pulumi.OutputState }

func (StorageSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSkuName)(nil)).Elem()
}

func (o StorageSkuNamePtrOutput) ToStorageSkuNamePtrOutput() StorageSkuNamePtrOutput {
	return o
}

func (o StorageSkuNamePtrOutput) ToStorageSkuNamePtrOutputWithContext(ctx context.Context) StorageSkuNamePtrOutput {
	return o
}

func (o StorageSkuNamePtrOutput) Elem() StorageSkuNameOutput {
	return o.ApplyT(func(v *StorageSkuName) StorageSkuName {
		if v != nil {
			return *v
		}
		var ret StorageSkuName
		return ret
	}).(StorageSkuNameOutput)
}

func (o StorageSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageSkuNameInput interface {
	pulumi.Input

	ToStorageSkuNameOutput() StorageSkuNameOutput
	ToStorageSkuNameOutputWithContext(context.Context) StorageSkuNameOutput
}

var storageSkuNamePtrType = reflect.TypeOf((**StorageSkuName)(nil)).Elem()

type StorageSkuNamePtrInput interface {
	pulumi.Input

	ToStorageSkuNamePtrOutput() StorageSkuNamePtrOutput
	ToStorageSkuNamePtrOutputWithContext(context.Context) StorageSkuNamePtrOutput
}

type storageSkuNamePtr string

func StorageSkuNamePtr(v string) StorageSkuNamePtrInput {
	return (*storageSkuNamePtr)(&v)
}

func (*storageSkuNamePtr) ElementType() reflect.Type {
	return storageSkuNamePtrType
}

func (in *storageSkuNamePtr) ToStorageSkuNamePtrOutput() StorageSkuNamePtrOutput {
	return pulumi.ToOutput(in).(StorageSkuNamePtrOutput)
}

func (in *storageSkuNamePtr) ToStorageSkuNamePtrOutputWithContext(ctx context.Context) StorageSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageSkuNamePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageSkuNameOutput{})
	pulumi.RegisterOutputType(StorageSkuNamePtrOutput{})
}
