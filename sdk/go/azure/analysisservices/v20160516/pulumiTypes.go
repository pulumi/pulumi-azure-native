


package v20160516

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}


func (val *ResourceSku) Defaults() *ResourceSku {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Capacity) {
		capacity_ := 1
		tmp.Capacity = &capacity_
	}
	return &tmp
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

func (o ResourceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
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

func (o ResourceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
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
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}


func (val *ResourceSkuResponse) Defaults() *ResourceSkuResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Capacity) {
		capacity_ := 1
		tmp.Capacity = &capacity_
	}
	return &tmp
}





type ResourceSkuResponseInput interface {
	pulumi.Input

	ToResourceSkuResponseOutput() ResourceSkuResponseOutput
	ToResourceSkuResponseOutputWithContext(context.Context) ResourceSkuResponseOutput
}

type ResourceSkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o ResourceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
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

func (o ResourceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
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

type ServerAdministrators struct {
	Members []string `pulumi:"members"`
}





type ServerAdministratorsInput interface {
	pulumi.Input

	ToServerAdministratorsOutput() ServerAdministratorsOutput
	ToServerAdministratorsOutputWithContext(context.Context) ServerAdministratorsOutput
}

type ServerAdministratorsArgs struct {
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (ServerAdministratorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministrators)(nil)).Elem()
}

func (i ServerAdministratorsArgs) ToServerAdministratorsOutput() ServerAdministratorsOutput {
	return i.ToServerAdministratorsOutputWithContext(context.Background())
}

func (i ServerAdministratorsArgs) ToServerAdministratorsOutputWithContext(ctx context.Context) ServerAdministratorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsOutput)
}

func (i ServerAdministratorsArgs) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return i.ToServerAdministratorsPtrOutputWithContext(context.Background())
}

func (i ServerAdministratorsArgs) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsOutput).ToServerAdministratorsPtrOutputWithContext(ctx)
}









type ServerAdministratorsPtrInput interface {
	pulumi.Input

	ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput
	ToServerAdministratorsPtrOutputWithContext(context.Context) ServerAdministratorsPtrOutput
}

type serverAdministratorsPtrType ServerAdministratorsArgs

func ServerAdministratorsPtr(v *ServerAdministratorsArgs) ServerAdministratorsPtrInput {
	return (*serverAdministratorsPtrType)(v)
}

func (*serverAdministratorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministrators)(nil)).Elem()
}

func (i *serverAdministratorsPtrType) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return i.ToServerAdministratorsPtrOutputWithContext(context.Background())
}

func (i *serverAdministratorsPtrType) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsPtrOutput)
}

type ServerAdministratorsOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministrators)(nil)).Elem()
}

func (o ServerAdministratorsOutput) ToServerAdministratorsOutput() ServerAdministratorsOutput {
	return o
}

func (o ServerAdministratorsOutput) ToServerAdministratorsOutputWithContext(ctx context.Context) ServerAdministratorsOutput {
	return o
}

func (o ServerAdministratorsOutput) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return o.ToServerAdministratorsPtrOutputWithContext(context.Background())
}

func (o ServerAdministratorsOutput) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerAdministrators) *ServerAdministrators {
		return &v
	}).(ServerAdministratorsPtrOutput)
}

func (o ServerAdministratorsOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServerAdministrators) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type ServerAdministratorsPtrOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministrators)(nil)).Elem()
}

func (o ServerAdministratorsPtrOutput) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return o
}

func (o ServerAdministratorsPtrOutput) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return o
}

func (o ServerAdministratorsPtrOutput) Elem() ServerAdministratorsOutput {
	return o.ApplyT(func(v *ServerAdministrators) ServerAdministrators {
		if v != nil {
			return *v
		}
		var ret ServerAdministrators
		return ret
	}).(ServerAdministratorsOutput)
}

func (o ServerAdministratorsPtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerAdministrators) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

type ServerAdministratorsResponse struct {
	Members []string `pulumi:"members"`
}





type ServerAdministratorsResponseInput interface {
	pulumi.Input

	ToServerAdministratorsResponseOutput() ServerAdministratorsResponseOutput
	ToServerAdministratorsResponseOutputWithContext(context.Context) ServerAdministratorsResponseOutput
}

type ServerAdministratorsResponseArgs struct {
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (ServerAdministratorsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministratorsResponse)(nil)).Elem()
}

func (i ServerAdministratorsResponseArgs) ToServerAdministratorsResponseOutput() ServerAdministratorsResponseOutput {
	return i.ToServerAdministratorsResponseOutputWithContext(context.Background())
}

func (i ServerAdministratorsResponseArgs) ToServerAdministratorsResponseOutputWithContext(ctx context.Context) ServerAdministratorsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsResponseOutput)
}

func (i ServerAdministratorsResponseArgs) ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput {
	return i.ToServerAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (i ServerAdministratorsResponseArgs) ToServerAdministratorsResponsePtrOutputWithContext(ctx context.Context) ServerAdministratorsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsResponseOutput).ToServerAdministratorsResponsePtrOutputWithContext(ctx)
}









type ServerAdministratorsResponsePtrInput interface {
	pulumi.Input

	ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput
	ToServerAdministratorsResponsePtrOutputWithContext(context.Context) ServerAdministratorsResponsePtrOutput
}

type serverAdministratorsResponsePtrType ServerAdministratorsResponseArgs

func ServerAdministratorsResponsePtr(v *ServerAdministratorsResponseArgs) ServerAdministratorsResponsePtrInput {
	return (*serverAdministratorsResponsePtrType)(v)
}

func (*serverAdministratorsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministratorsResponse)(nil)).Elem()
}

func (i *serverAdministratorsResponsePtrType) ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput {
	return i.ToServerAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (i *serverAdministratorsResponsePtrType) ToServerAdministratorsResponsePtrOutputWithContext(ctx context.Context) ServerAdministratorsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsResponsePtrOutput)
}

type ServerAdministratorsResponseOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministratorsResponse)(nil)).Elem()
}

func (o ServerAdministratorsResponseOutput) ToServerAdministratorsResponseOutput() ServerAdministratorsResponseOutput {
	return o
}

func (o ServerAdministratorsResponseOutput) ToServerAdministratorsResponseOutputWithContext(ctx context.Context) ServerAdministratorsResponseOutput {
	return o
}

func (o ServerAdministratorsResponseOutput) ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput {
	return o.ToServerAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (o ServerAdministratorsResponseOutput) ToServerAdministratorsResponsePtrOutputWithContext(ctx context.Context) ServerAdministratorsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerAdministratorsResponse) *ServerAdministratorsResponse {
		return &v
	}).(ServerAdministratorsResponsePtrOutput)
}

func (o ServerAdministratorsResponseOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServerAdministratorsResponse) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type ServerAdministratorsResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministratorsResponse)(nil)).Elem()
}

func (o ServerAdministratorsResponsePtrOutput) ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput {
	return o
}

func (o ServerAdministratorsResponsePtrOutput) ToServerAdministratorsResponsePtrOutputWithContext(ctx context.Context) ServerAdministratorsResponsePtrOutput {
	return o
}

func (o ServerAdministratorsResponsePtrOutput) Elem() ServerAdministratorsResponseOutput {
	return o.ApplyT(func(v *ServerAdministratorsResponse) ServerAdministratorsResponse {
		if v != nil {
			return *v
		}
		var ret ServerAdministratorsResponse
		return ret
	}).(ServerAdministratorsResponseOutput)
}

func (o ServerAdministratorsResponsePtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerAdministratorsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsResponsePtrOutput{})
}
