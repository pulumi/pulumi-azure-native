


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


func (val *ResourceSkuArgs) Defaults() *ResourceSkuArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Capacity) {
		tmp.Capacity = pulumi.IntPtr(1)
	}
	return &tmp
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
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsResponsePtrOutput{})
}
