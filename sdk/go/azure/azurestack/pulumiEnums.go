


package azurestack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Location string

const (
	LocationGlobal = Location("global")
)

func (Location) ElementType() reflect.Type {
	return reflect.TypeOf((*Location)(nil)).Elem()
}

func (e Location) ToLocationOutput() LocationOutput {
	return pulumi.ToOutput(e).(LocationOutput)
}

func (e Location) ToLocationOutputWithContext(ctx context.Context) LocationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LocationOutput)
}

func (e Location) ToLocationPtrOutput() LocationPtrOutput {
	return e.ToLocationPtrOutputWithContext(context.Background())
}

func (e Location) ToLocationPtrOutputWithContext(ctx context.Context) LocationPtrOutput {
	return Location(e).ToLocationOutputWithContext(ctx).ToLocationPtrOutputWithContext(ctx)
}

func (e Location) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Location) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Location) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Location) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LocationOutput struct{ *pulumi.OutputState }

func (LocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Location)(nil)).Elem()
}

func (o LocationOutput) ToLocationOutput() LocationOutput {
	return o
}

func (o LocationOutput) ToLocationOutputWithContext(ctx context.Context) LocationOutput {
	return o
}

func (o LocationOutput) ToLocationPtrOutput() LocationPtrOutput {
	return o.ToLocationPtrOutputWithContext(context.Background())
}

func (o LocationOutput) ToLocationPtrOutputWithContext(ctx context.Context) LocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Location) *Location {
		return &v
	}).(LocationPtrOutput)
}

func (o LocationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LocationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Location) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LocationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LocationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Location) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LocationPtrOutput struct{ *pulumi.OutputState }

func (LocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Location)(nil)).Elem()
}

func (o LocationPtrOutput) ToLocationPtrOutput() LocationPtrOutput {
	return o
}

func (o LocationPtrOutput) ToLocationPtrOutputWithContext(ctx context.Context) LocationPtrOutput {
	return o
}

func (o LocationPtrOutput) Elem() LocationOutput {
	return o.ApplyT(func(v *Location) Location {
		if v != nil {
			return *v
		}
		var ret Location
		return ret
	}).(LocationOutput)
}

func (o LocationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LocationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Location) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LocationInput interface {
	pulumi.Input

	ToLocationOutput() LocationOutput
	ToLocationOutputWithContext(context.Context) LocationOutput
}

var locationPtrType = reflect.TypeOf((**Location)(nil)).Elem()

type LocationPtrInput interface {
	pulumi.Input

	ToLocationPtrOutput() LocationPtrOutput
	ToLocationPtrOutputWithContext(context.Context) LocationPtrOutput
}

type locationPtr string

func LocationPtr(v string) LocationPtrInput {
	return (*locationPtr)(&v)
}

func (*locationPtr) ElementType() reflect.Type {
	return locationPtrType
}

func (in *locationPtr) ToLocationPtrOutput() LocationPtrOutput {
	return pulumi.ToOutput(in).(LocationPtrOutput)
}

func (in *locationPtr) ToLocationPtrOutputWithContext(ctx context.Context) LocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LocationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LocationOutput{})
	pulumi.RegisterOutputType(LocationPtrOutput{})
}
