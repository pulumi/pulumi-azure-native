


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceLevel string

const (
	// Standard service level
	ServiceLevelStandard = ServiceLevel("Standard")
	// Premium service level
	ServiceLevelPremium = ServiceLevel("Premium")
	// Ultra service level
	ServiceLevelUltra = ServiceLevel("Ultra")
)

func (ServiceLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLevel)(nil)).Elem()
}

func (e ServiceLevel) ToServiceLevelOutput() ServiceLevelOutput {
	return pulumi.ToOutput(e).(ServiceLevelOutput)
}

func (e ServiceLevel) ToServiceLevelOutputWithContext(ctx context.Context) ServiceLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceLevelOutput)
}

func (e ServiceLevel) ToServiceLevelPtrOutput() ServiceLevelPtrOutput {
	return e.ToServiceLevelPtrOutputWithContext(context.Background())
}

func (e ServiceLevel) ToServiceLevelPtrOutputWithContext(ctx context.Context) ServiceLevelPtrOutput {
	return ServiceLevel(e).ToServiceLevelOutputWithContext(ctx).ToServiceLevelPtrOutputWithContext(ctx)
}

func (e ServiceLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceLevelOutput struct{ *pulumi.OutputState }

func (ServiceLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLevel)(nil)).Elem()
}

func (o ServiceLevelOutput) ToServiceLevelOutput() ServiceLevelOutput {
	return o
}

func (o ServiceLevelOutput) ToServiceLevelOutputWithContext(ctx context.Context) ServiceLevelOutput {
	return o
}

func (o ServiceLevelOutput) ToServiceLevelPtrOutput() ServiceLevelPtrOutput {
	return o.ToServiceLevelPtrOutputWithContext(context.Background())
}

func (o ServiceLevelOutput) ToServiceLevelPtrOutputWithContext(ctx context.Context) ServiceLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceLevel) *ServiceLevel {
		return &v
	}).(ServiceLevelPtrOutput)
}

func (o ServiceLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceLevelPtrOutput struct{ *pulumi.OutputState }

func (ServiceLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceLevel)(nil)).Elem()
}

func (o ServiceLevelPtrOutput) ToServiceLevelPtrOutput() ServiceLevelPtrOutput {
	return o
}

func (o ServiceLevelPtrOutput) ToServiceLevelPtrOutputWithContext(ctx context.Context) ServiceLevelPtrOutput {
	return o
}

func (o ServiceLevelPtrOutput) Elem() ServiceLevelOutput {
	return o.ApplyT(func(v *ServiceLevel) ServiceLevel {
		if v != nil {
			return *v
		}
		var ret ServiceLevel
		return ret
	}).(ServiceLevelOutput)
}

func (o ServiceLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServiceLevelInput interface {
	pulumi.Input

	ToServiceLevelOutput() ServiceLevelOutput
	ToServiceLevelOutputWithContext(context.Context) ServiceLevelOutput
}

var serviceLevelPtrType = reflect.TypeOf((**ServiceLevel)(nil)).Elem()

type ServiceLevelPtrInput interface {
	pulumi.Input

	ToServiceLevelPtrOutput() ServiceLevelPtrOutput
	ToServiceLevelPtrOutputWithContext(context.Context) ServiceLevelPtrOutput
}

type serviceLevelPtr string

func ServiceLevelPtr(v string) ServiceLevelPtrInput {
	return (*serviceLevelPtr)(&v)
}

func (*serviceLevelPtr) ElementType() reflect.Type {
	return serviceLevelPtrType
}

func (in *serviceLevelPtr) ToServiceLevelPtrOutput() ServiceLevelPtrOutput {
	return pulumi.ToOutput(in).(ServiceLevelPtrOutput)
}

func (in *serviceLevelPtr) ToServiceLevelPtrOutputWithContext(ctx context.Context) ServiceLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceLevelPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLevelInput)(nil)).Elem(), ServiceLevel("Standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLevelPtrInput)(nil)).Elem(), ServiceLevel("Standard"))
	pulumi.RegisterOutputType(ServiceLevelOutput{})
	pulumi.RegisterOutputType(ServiceLevelPtrOutput{})
}
