


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





type DedicatedCapacityAdministratorsResponseInput interface {
	pulumi.Input

	ToDedicatedCapacityAdministratorsResponseOutput() DedicatedCapacityAdministratorsResponseOutput
	ToDedicatedCapacityAdministratorsResponseOutputWithContext(context.Context) DedicatedCapacityAdministratorsResponseOutput
}

type DedicatedCapacityAdministratorsResponseArgs struct {
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (DedicatedCapacityAdministratorsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministratorsResponse)(nil)).Elem()
}

func (i DedicatedCapacityAdministratorsResponseArgs) ToDedicatedCapacityAdministratorsResponseOutput() DedicatedCapacityAdministratorsResponseOutput {
	return i.ToDedicatedCapacityAdministratorsResponseOutputWithContext(context.Background())
}

func (i DedicatedCapacityAdministratorsResponseArgs) ToDedicatedCapacityAdministratorsResponseOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsResponseOutput)
}

func (i DedicatedCapacityAdministratorsResponseArgs) ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput {
	return i.ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (i DedicatedCapacityAdministratorsResponseArgs) ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsResponseOutput).ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx)
}









type DedicatedCapacityAdministratorsResponsePtrInput interface {
	pulumi.Input

	ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput
	ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(context.Context) DedicatedCapacityAdministratorsResponsePtrOutput
}

type dedicatedCapacityAdministratorsResponsePtrType DedicatedCapacityAdministratorsResponseArgs

func DedicatedCapacityAdministratorsResponsePtr(v *DedicatedCapacityAdministratorsResponseArgs) DedicatedCapacityAdministratorsResponsePtrInput {
	return (*dedicatedCapacityAdministratorsResponsePtrType)(v)
}

func (*dedicatedCapacityAdministratorsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministratorsResponse)(nil)).Elem()
}

func (i *dedicatedCapacityAdministratorsResponsePtrType) ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput {
	return i.ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (i *dedicatedCapacityAdministratorsResponsePtrType) ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsResponsePtrOutput)
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

func (o DedicatedCapacityAdministratorsResponseOutput) ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput {
	return o.ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (o DedicatedCapacityAdministratorsResponseOutput) ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DedicatedCapacityAdministratorsResponse) *DedicatedCapacityAdministratorsResponse {
		return &v
	}).(DedicatedCapacityAdministratorsResponsePtrOutput)
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
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
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

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
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

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type ResourceSkuResponseInput interface {
	pulumi.Input

	ToResourceSkuResponseOutput() ResourceSkuResponseOutput
	ToResourceSkuResponseOutputWithContext(context.Context) ResourceSkuResponseOutput
}

type ResourceSkuResponseArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return i.ToResourceSkuResponseOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput)
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput).ToResourceSkuResponsePtrOutputWithContext(ctx)
}









type ResourceSkuResponsePtrInput interface {
	pulumi.Input

	ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput
	ToResourceSkuResponsePtrOutputWithContext(context.Context) ResourceSkuResponsePtrOutput
}

type resourceSkuResponsePtrType ResourceSkuResponseArgs

func ResourceSkuResponsePtr(v *ResourceSkuResponseArgs) ResourceSkuResponsePtrInput {
	return (*resourceSkuResponsePtrType)(v)
}

func (*resourceSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponsePtrOutput)
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

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSkuResponse) *ResourceSkuResponse {
		return &v
	}).(ResourceSkuResponsePtrOutput)
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCapacityAdministratorsInput)(nil)).Elem(), DedicatedCapacityAdministratorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCapacityAdministratorsPtrInput)(nil)).Elem(), DedicatedCapacityAdministratorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCapacityAdministratorsResponseInput)(nil)).Elem(), DedicatedCapacityAdministratorsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCapacityAdministratorsResponsePtrInput)(nil)).Elem(), DedicatedCapacityAdministratorsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuInput)(nil)).Elem(), ResourceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuPtrInput)(nil)).Elem(), ResourceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuResponseInput)(nil)).Elem(), ResourceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuResponsePtrInput)(nil)).Elem(), ResourceSkuResponseArgs{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
}
