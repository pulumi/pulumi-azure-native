


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageRedundancy string

const (
	StorageRedundancyLocal   = StorageRedundancy("Local")
	StorageRedundancyGeo     = StorageRedundancy("Geo")
	StorageRedundancyZone    = StorageRedundancy("Zone")
	StorageRedundancyGeoZone = StorageRedundancy("GeoZone")
)

func (StorageRedundancy) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageRedundancy)(nil)).Elem()
}

func (e StorageRedundancy) ToStorageRedundancyOutput() StorageRedundancyOutput {
	return pulumi.ToOutput(e).(StorageRedundancyOutput)
}

func (e StorageRedundancy) ToStorageRedundancyOutputWithContext(ctx context.Context) StorageRedundancyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageRedundancyOutput)
}

func (e StorageRedundancy) ToStorageRedundancyPtrOutput() StorageRedundancyPtrOutput {
	return e.ToStorageRedundancyPtrOutputWithContext(context.Background())
}

func (e StorageRedundancy) ToStorageRedundancyPtrOutputWithContext(ctx context.Context) StorageRedundancyPtrOutput {
	return StorageRedundancy(e).ToStorageRedundancyOutputWithContext(ctx).ToStorageRedundancyPtrOutputWithContext(ctx)
}

func (e StorageRedundancy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageRedundancy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageRedundancy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageRedundancy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageRedundancyOutput struct{ *pulumi.OutputState }

func (StorageRedundancyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageRedundancy)(nil)).Elem()
}

func (o StorageRedundancyOutput) ToStorageRedundancyOutput() StorageRedundancyOutput {
	return o
}

func (o StorageRedundancyOutput) ToStorageRedundancyOutputWithContext(ctx context.Context) StorageRedundancyOutput {
	return o
}

func (o StorageRedundancyOutput) ToStorageRedundancyPtrOutput() StorageRedundancyPtrOutput {
	return o.ToStorageRedundancyPtrOutputWithContext(context.Background())
}

func (o StorageRedundancyOutput) ToStorageRedundancyPtrOutputWithContext(ctx context.Context) StorageRedundancyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageRedundancy) *StorageRedundancy {
		return &v
	}).(StorageRedundancyPtrOutput)
}

func (o StorageRedundancyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageRedundancyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageRedundancy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageRedundancyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageRedundancyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageRedundancy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageRedundancyPtrOutput struct{ *pulumi.OutputState }

func (StorageRedundancyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageRedundancy)(nil)).Elem()
}

func (o StorageRedundancyPtrOutput) ToStorageRedundancyPtrOutput() StorageRedundancyPtrOutput {
	return o
}

func (o StorageRedundancyPtrOutput) ToStorageRedundancyPtrOutputWithContext(ctx context.Context) StorageRedundancyPtrOutput {
	return o
}

func (o StorageRedundancyPtrOutput) Elem() StorageRedundancyOutput {
	return o.ApplyT(func(v *StorageRedundancy) StorageRedundancy {
		if v != nil {
			return *v
		}
		var ret StorageRedundancy
		return ret
	}).(StorageRedundancyOutput)
}

func (o StorageRedundancyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageRedundancyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageRedundancy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageRedundancyInput interface {
	pulumi.Input

	ToStorageRedundancyOutput() StorageRedundancyOutput
	ToStorageRedundancyOutputWithContext(context.Context) StorageRedundancyOutput
}

var storageRedundancyPtrType = reflect.TypeOf((**StorageRedundancy)(nil)).Elem()

type StorageRedundancyPtrInput interface {
	pulumi.Input

	ToStorageRedundancyPtrOutput() StorageRedundancyPtrOutput
	ToStorageRedundancyPtrOutputWithContext(context.Context) StorageRedundancyPtrOutput
}

type storageRedundancyPtr string

func StorageRedundancyPtr(v string) StorageRedundancyPtrInput {
	return (*storageRedundancyPtr)(&v)
}

func (*storageRedundancyPtr) ElementType() reflect.Type {
	return storageRedundancyPtrType
}

func (in *storageRedundancyPtr) ToStorageRedundancyPtrOutput() StorageRedundancyPtrOutput {
	return pulumi.ToOutput(in).(StorageRedundancyPtrOutput)
}

func (in *storageRedundancyPtr) ToStorageRedundancyPtrOutputWithContext(ctx context.Context) StorageRedundancyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageRedundancyPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageRedundancyInput)(nil)).Elem(), StorageRedundancy("Local"))
	pulumi.RegisterInputType(reflect.TypeOf((*StorageRedundancyPtrInput)(nil)).Elem(), StorageRedundancy("Local"))
	pulumi.RegisterOutputType(StorageRedundancyOutput{})
	pulumi.RegisterOutputType(StorageRedundancyPtrOutput{})
}
