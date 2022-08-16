


package v20171001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DedicatedCapacityAdministrators struct {
	Members []string `pulumi:"members"`
}





type DedicatedCapacityAdministratorsInput interface {
	pulumi.Input

	ToDedicatedCapacityAdministratorsOutput() DedicatedCapacityAdministratorsOutput
	ToDedicatedCapacityAdministratorsOutputWithContext(context.Context) DedicatedCapacityAdministratorsOutput
}

type DedicatedCapacityAdministratorsArgs struct {
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (DedicatedCapacityAdministratorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministrators)(nil)).Elem()
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsOutput() DedicatedCapacityAdministratorsOutput {
	return i.ToDedicatedCapacityAdministratorsOutputWithContext(context.Background())
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsOutput)
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return i.ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Background())
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsOutput).ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx)
}









type DedicatedCapacityAdministratorsPtrInput interface {
	pulumi.Input

	ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput
	ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Context) DedicatedCapacityAdministratorsPtrOutput
}

type dedicatedCapacityAdministratorsPtrType DedicatedCapacityAdministratorsArgs

func DedicatedCapacityAdministratorsPtr(v *DedicatedCapacityAdministratorsArgs) DedicatedCapacityAdministratorsPtrInput {
	return (*dedicatedCapacityAdministratorsPtrType)(v)
}

func (*dedicatedCapacityAdministratorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministrators)(nil)).Elem()
}

func (i *dedicatedCapacityAdministratorsPtrType) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return i.ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Background())
}

func (i *dedicatedCapacityAdministratorsPtrType) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsPtrOutput)
}

type DedicatedCapacityAdministratorsOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministrators)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsOutput() DedicatedCapacityAdministratorsOutput {
	return o
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsOutput {
	return o
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return o.ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Background())
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DedicatedCapacityAdministrators) *DedicatedCapacityAdministrators {
		return &v
	}).(DedicatedCapacityAdministratorsPtrOutput)
}

func (o DedicatedCapacityAdministratorsOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DedicatedCapacityAdministrators) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type DedicatedCapacityAdministratorsPtrOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministrators)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsPtrOutput) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsPtrOutput) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsPtrOutput) Elem() DedicatedCapacityAdministratorsOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministrators) DedicatedCapacityAdministrators {
		if v != nil {
			return *v
		}
		var ret DedicatedCapacityAdministrators
		return ret
	}).(DedicatedCapacityAdministratorsOutput)
}

func (o DedicatedCapacityAdministratorsPtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministrators) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

type DedicatedCapacityAdministratorsResponse struct {
	Members []string `pulumi:"members"`
}

type DedicatedCapacityAdministratorsResponseOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministratorsResponse)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsResponseOutput) ToDedicatedCapacityAdministratorsResponseOutput() DedicatedCapacityAdministratorsResponseOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponseOutput) ToDedicatedCapacityAdministratorsResponseOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponseOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponseOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DedicatedCapacityAdministratorsResponse) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type DedicatedCapacityAdministratorsResponsePtrOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministratorsResponse)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponsePtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) Elem() DedicatedCapacityAdministratorsResponseOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministratorsResponse) DedicatedCapacityAdministratorsResponse {
		if v != nil {
			return *v
		}
		var ret DedicatedCapacityAdministratorsResponse
		return ret
	}).(DedicatedCapacityAdministratorsResponseOutput)
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministratorsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

type ResourceSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
}
