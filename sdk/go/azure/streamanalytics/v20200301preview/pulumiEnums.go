


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterSkuName string

const (
	// The default SKU.
	ClusterSkuNameDefault = ClusterSkuName("Default")
)

func (ClusterSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSkuName)(nil)).Elem()
}

func (e ClusterSkuName) ToClusterSkuNameOutput() ClusterSkuNameOutput {
	return pulumi.ToOutput(e).(ClusterSkuNameOutput)
}

func (e ClusterSkuName) ToClusterSkuNameOutputWithContext(ctx context.Context) ClusterSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterSkuNameOutput)
}

func (e ClusterSkuName) ToClusterSkuNamePtrOutput() ClusterSkuNamePtrOutput {
	return e.ToClusterSkuNamePtrOutputWithContext(context.Background())
}

func (e ClusterSkuName) ToClusterSkuNamePtrOutputWithContext(ctx context.Context) ClusterSkuNamePtrOutput {
	return ClusterSkuName(e).ToClusterSkuNameOutputWithContext(ctx).ToClusterSkuNamePtrOutputWithContext(ctx)
}

func (e ClusterSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterSkuNameOutput struct{ *pulumi.OutputState }

func (ClusterSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSkuName)(nil)).Elem()
}

func (o ClusterSkuNameOutput) ToClusterSkuNameOutput() ClusterSkuNameOutput {
	return o
}

func (o ClusterSkuNameOutput) ToClusterSkuNameOutputWithContext(ctx context.Context) ClusterSkuNameOutput {
	return o
}

func (o ClusterSkuNameOutput) ToClusterSkuNamePtrOutput() ClusterSkuNamePtrOutput {
	return o.ToClusterSkuNamePtrOutputWithContext(context.Background())
}

func (o ClusterSkuNameOutput) ToClusterSkuNamePtrOutputWithContext(ctx context.Context) ClusterSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSkuName) *ClusterSkuName {
		return &v
	}).(ClusterSkuNamePtrOutput)
}

func (o ClusterSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterSkuNamePtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSkuName)(nil)).Elem()
}

func (o ClusterSkuNamePtrOutput) ToClusterSkuNamePtrOutput() ClusterSkuNamePtrOutput {
	return o
}

func (o ClusterSkuNamePtrOutput) ToClusterSkuNamePtrOutputWithContext(ctx context.Context) ClusterSkuNamePtrOutput {
	return o
}

func (o ClusterSkuNamePtrOutput) Elem() ClusterSkuNameOutput {
	return o.ApplyT(func(v *ClusterSkuName) ClusterSkuName {
		if v != nil {
			return *v
		}
		var ret ClusterSkuName
		return ret
	}).(ClusterSkuNameOutput)
}

func (o ClusterSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClusterSkuNameInput interface {
	pulumi.Input

	ToClusterSkuNameOutput() ClusterSkuNameOutput
	ToClusterSkuNameOutputWithContext(context.Context) ClusterSkuNameOutput
}

var clusterSkuNamePtrType = reflect.TypeOf((**ClusterSkuName)(nil)).Elem()

type ClusterSkuNamePtrInput interface {
	pulumi.Input

	ToClusterSkuNamePtrOutput() ClusterSkuNamePtrOutput
	ToClusterSkuNamePtrOutputWithContext(context.Context) ClusterSkuNamePtrOutput
}

type clusterSkuNamePtr string

func ClusterSkuNamePtr(v string) ClusterSkuNamePtrInput {
	return (*clusterSkuNamePtr)(&v)
}

func (*clusterSkuNamePtr) ElementType() reflect.Type {
	return clusterSkuNamePtrType
}

func (in *clusterSkuNamePtr) ToClusterSkuNamePtrOutput() ClusterSkuNamePtrOutput {
	return pulumi.ToOutput(in).(ClusterSkuNamePtrOutput)
}

func (in *clusterSkuNamePtr) ToClusterSkuNamePtrOutputWithContext(ctx context.Context) ClusterSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterSkuNamePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterSkuNameOutput{})
	pulumi.RegisterOutputType(ClusterSkuNamePtrOutput{})
}
