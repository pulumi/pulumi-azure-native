


package subscription

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workload string

const (
	WorkloadProduction = Workload("Production")
	WorkloadDevTest    = Workload("DevTest")
)

func (Workload) ElementType() reflect.Type {
	return reflect.TypeOf((*Workload)(nil)).Elem()
}

func (e Workload) ToWorkloadOutput() WorkloadOutput {
	return pulumi.ToOutput(e).(WorkloadOutput)
}

func (e Workload) ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkloadOutput)
}

func (e Workload) ToWorkloadPtrOutput() WorkloadPtrOutput {
	return e.ToWorkloadPtrOutputWithContext(context.Background())
}

func (e Workload) ToWorkloadPtrOutputWithContext(ctx context.Context) WorkloadPtrOutput {
	return Workload(e).ToWorkloadOutputWithContext(ctx).ToWorkloadPtrOutputWithContext(ctx)
}

func (e Workload) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Workload) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Workload) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Workload) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkloadOutput struct{ *pulumi.OutputState }

func (WorkloadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Workload)(nil)).Elem()
}

func (o WorkloadOutput) ToWorkloadOutput() WorkloadOutput {
	return o
}

func (o WorkloadOutput) ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput {
	return o
}

func (o WorkloadOutput) ToWorkloadPtrOutput() WorkloadPtrOutput {
	return o.ToWorkloadPtrOutputWithContext(context.Background())
}

func (o WorkloadOutput) ToWorkloadPtrOutputWithContext(ctx context.Context) WorkloadPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Workload) *Workload {
		return &v
	}).(WorkloadPtrOutput)
}

func (o WorkloadOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkloadOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Workload) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkloadOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Workload) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkloadPtrOutput struct{ *pulumi.OutputState }

func (WorkloadPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workload)(nil)).Elem()
}

func (o WorkloadPtrOutput) ToWorkloadPtrOutput() WorkloadPtrOutput {
	return o
}

func (o WorkloadPtrOutput) ToWorkloadPtrOutputWithContext(ctx context.Context) WorkloadPtrOutput {
	return o
}

func (o WorkloadPtrOutput) Elem() WorkloadOutput {
	return o.ApplyT(func(v *Workload) Workload {
		if v != nil {
			return *v
		}
		var ret Workload
		return ret
	}).(WorkloadOutput)
}

func (o WorkloadPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Workload) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WorkloadInput interface {
	pulumi.Input

	ToWorkloadOutput() WorkloadOutput
	ToWorkloadOutputWithContext(context.Context) WorkloadOutput
}

var workloadPtrType = reflect.TypeOf((**Workload)(nil)).Elem()

type WorkloadPtrInput interface {
	pulumi.Input

	ToWorkloadPtrOutput() WorkloadPtrOutput
	ToWorkloadPtrOutputWithContext(context.Context) WorkloadPtrOutput
}

type workloadPtr string

func WorkloadPtr(v string) WorkloadPtrInput {
	return (*workloadPtr)(&v)
}

func (*workloadPtr) ElementType() reflect.Type {
	return workloadPtrType
}

func (in *workloadPtr) ToWorkloadPtrOutput() WorkloadPtrOutput {
	return pulumi.ToOutput(in).(WorkloadPtrOutput)
}

func (in *workloadPtr) ToWorkloadPtrOutputWithContext(ctx context.Context) WorkloadPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkloadPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadInput)(nil)).Elem(), Workload("Production"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadPtrInput)(nil)).Elem(), Workload("Production"))
	pulumi.RegisterOutputType(WorkloadOutput{})
	pulumi.RegisterOutputType(WorkloadPtrOutput{})
}
