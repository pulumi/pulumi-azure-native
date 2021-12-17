


package v20170714

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayDetails struct {
	GatewayResourceId *string `pulumi:"gatewayResourceId"`
}





type GatewayDetailsInput interface {
	pulumi.Input

	ToGatewayDetailsOutput() GatewayDetailsOutput
	ToGatewayDetailsOutputWithContext(context.Context) GatewayDetailsOutput
}

type GatewayDetailsArgs struct {
	GatewayResourceId pulumi.StringPtrInput `pulumi:"gatewayResourceId"`
}

func (GatewayDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDetails)(nil)).Elem()
}

func (i GatewayDetailsArgs) ToGatewayDetailsOutput() GatewayDetailsOutput {
	return i.ToGatewayDetailsOutputWithContext(context.Background())
}

func (i GatewayDetailsArgs) ToGatewayDetailsOutputWithContext(ctx context.Context) GatewayDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsOutput)
}

func (i GatewayDetailsArgs) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return i.ToGatewayDetailsPtrOutputWithContext(context.Background())
}

func (i GatewayDetailsArgs) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsOutput).ToGatewayDetailsPtrOutputWithContext(ctx)
}









type GatewayDetailsPtrInput interface {
	pulumi.Input

	ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput
	ToGatewayDetailsPtrOutputWithContext(context.Context) GatewayDetailsPtrOutput
}

type gatewayDetailsPtrType GatewayDetailsArgs

func GatewayDetailsPtr(v *GatewayDetailsArgs) GatewayDetailsPtrInput {
	return (*gatewayDetailsPtrType)(v)
}

func (*gatewayDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayDetails)(nil)).Elem()
}

func (i *gatewayDetailsPtrType) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return i.ToGatewayDetailsPtrOutputWithContext(context.Background())
}

func (i *gatewayDetailsPtrType) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsPtrOutput)
}

type GatewayDetailsOutput struct{ *pulumi.OutputState }

func (GatewayDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDetails)(nil)).Elem()
}

func (o GatewayDetailsOutput) ToGatewayDetailsOutput() GatewayDetailsOutput {
	return o
}

func (o GatewayDetailsOutput) ToGatewayDetailsOutputWithContext(ctx context.Context) GatewayDetailsOutput {
	return o
}

func (o GatewayDetailsOutput) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return o.ToGatewayDetailsPtrOutputWithContext(context.Background())
}

func (o GatewayDetailsOutput) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayDetails) *GatewayDetails {
		return &v
	}).(GatewayDetailsPtrOutput)
}

func (o GatewayDetailsOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayDetails) *string { return v.GatewayResourceId }).(pulumi.StringPtrOutput)
}

type GatewayDetailsPtrOutput struct{ *pulumi.OutputState }

func (GatewayDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayDetails)(nil)).Elem()
}

func (o GatewayDetailsPtrOutput) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return o
}

func (o GatewayDetailsPtrOutput) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return o
}

func (o GatewayDetailsPtrOutput) Elem() GatewayDetailsOutput {
	return o.ApplyT(func(v *GatewayDetails) GatewayDetails {
		if v != nil {
			return *v
		}
		var ret GatewayDetails
		return ret
	}).(GatewayDetailsOutput)
}

func (o GatewayDetailsPtrOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetails) *string {
		if v == nil {
			return nil
		}
		return v.GatewayResourceId
	}).(pulumi.StringPtrOutput)
}

type GatewayDetailsResponse struct {
	DmtsClusterUri    string  `pulumi:"dmtsClusterUri"`
	GatewayObjectId   string  `pulumi:"gatewayObjectId"`
	GatewayResourceId *string `pulumi:"gatewayResourceId"`
}

type GatewayDetailsResponseOutput struct{ *pulumi.OutputState }

func (GatewayDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDetailsResponse)(nil)).Elem()
}

func (o GatewayDetailsResponseOutput) ToGatewayDetailsResponseOutput() GatewayDetailsResponseOutput {
	return o
}

func (o GatewayDetailsResponseOutput) ToGatewayDetailsResponseOutputWithContext(ctx context.Context) GatewayDetailsResponseOutput {
	return o
}

func (o GatewayDetailsResponseOutput) DmtsClusterUri() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDetailsResponse) string { return v.DmtsClusterUri }).(pulumi.StringOutput)
}

func (o GatewayDetailsResponseOutput) GatewayObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDetailsResponse) string { return v.GatewayObjectId }).(pulumi.StringOutput)
}

func (o GatewayDetailsResponseOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayDetailsResponse) *string { return v.GatewayResourceId }).(pulumi.StringPtrOutput)
}

type GatewayDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (GatewayDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayDetailsResponse)(nil)).Elem()
}

func (o GatewayDetailsResponsePtrOutput) ToGatewayDetailsResponsePtrOutput() GatewayDetailsResponsePtrOutput {
	return o
}

func (o GatewayDetailsResponsePtrOutput) ToGatewayDetailsResponsePtrOutputWithContext(ctx context.Context) GatewayDetailsResponsePtrOutput {
	return o
}

func (o GatewayDetailsResponsePtrOutput) Elem() GatewayDetailsResponseOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) GatewayDetailsResponse {
		if v != nil {
			return *v
		}
		var ret GatewayDetailsResponse
		return ret
	}).(GatewayDetailsResponseOutput)
}

func (o GatewayDetailsResponsePtrOutput) DmtsClusterUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DmtsClusterUri
	}).(pulumi.StringPtrOutput)
}

func (o GatewayDetailsResponsePtrOutput) GatewayObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GatewayObjectId
	}).(pulumi.StringPtrOutput)
}

func (o GatewayDetailsResponsePtrOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GatewayResourceId
	}).(pulumi.StringPtrOutput)
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

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
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
	pulumi.RegisterOutputType(GatewayDetailsOutput{})
	pulumi.RegisterOutputType(GatewayDetailsPtrOutput{})
	pulumi.RegisterOutputType(GatewayDetailsResponseOutput{})
	pulumi.RegisterOutputType(GatewayDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsResponsePtrOutput{})
}
