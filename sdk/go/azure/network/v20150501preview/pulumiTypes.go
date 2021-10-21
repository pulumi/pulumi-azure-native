


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddressSpace struct {
	AddressPrefixes []string `pulumi:"addressPrefixes"`
}





type AddressSpaceInput interface {
	pulumi.Input

	ToAddressSpaceOutput() AddressSpaceOutput
	ToAddressSpaceOutputWithContext(context.Context) AddressSpaceOutput
}

type AddressSpaceArgs struct {
	AddressPrefixes pulumi.StringArrayInput `pulumi:"addressPrefixes"`
}

func (AddressSpaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpace)(nil)).Elem()
}

func (i AddressSpaceArgs) ToAddressSpaceOutput() AddressSpaceOutput {
	return i.ToAddressSpaceOutputWithContext(context.Background())
}

func (i AddressSpaceArgs) ToAddressSpaceOutputWithContext(ctx context.Context) AddressSpaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceOutput)
}

func (i AddressSpaceArgs) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return i.ToAddressSpacePtrOutputWithContext(context.Background())
}

func (i AddressSpaceArgs) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceOutput).ToAddressSpacePtrOutputWithContext(ctx)
}









type AddressSpacePtrInput interface {
	pulumi.Input

	ToAddressSpacePtrOutput() AddressSpacePtrOutput
	ToAddressSpacePtrOutputWithContext(context.Context) AddressSpacePtrOutput
}

type addressSpacePtrType AddressSpaceArgs

func AddressSpacePtr(v *AddressSpaceArgs) AddressSpacePtrInput {
	return (*addressSpacePtrType)(v)
}

func (*addressSpacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpace)(nil)).Elem()
}

func (i *addressSpacePtrType) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return i.ToAddressSpacePtrOutputWithContext(context.Background())
}

func (i *addressSpacePtrType) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpacePtrOutput)
}

type AddressSpaceOutput struct{ *pulumi.OutputState }

func (AddressSpaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpace)(nil)).Elem()
}

func (o AddressSpaceOutput) ToAddressSpaceOutput() AddressSpaceOutput {
	return o
}

func (o AddressSpaceOutput) ToAddressSpaceOutputWithContext(ctx context.Context) AddressSpaceOutput {
	return o
}

func (o AddressSpaceOutput) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return o.ToAddressSpacePtrOutputWithContext(context.Background())
}

func (o AddressSpaceOutput) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressSpace) *AddressSpace {
		return &v
	}).(AddressSpacePtrOutput)
}

func (o AddressSpaceOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddressSpace) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

type AddressSpacePtrOutput struct{ *pulumi.OutputState }

func (AddressSpacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpace)(nil)).Elem()
}

func (o AddressSpacePtrOutput) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return o
}

func (o AddressSpacePtrOutput) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return o
}

func (o AddressSpacePtrOutput) Elem() AddressSpaceOutput {
	return o.ApplyT(func(v *AddressSpace) AddressSpace {
		if v != nil {
			return *v
		}
		var ret AddressSpace
		return ret
	}).(AddressSpaceOutput)
}

func (o AddressSpacePtrOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AddressSpace) []string {
		if v == nil {
			return nil
		}
		return v.AddressPrefixes
	}).(pulumi.StringArrayOutput)
}

type AddressSpaceResponse struct {
	AddressPrefixes []string `pulumi:"addressPrefixes"`
}





type AddressSpaceResponseInput interface {
	pulumi.Input

	ToAddressSpaceResponseOutput() AddressSpaceResponseOutput
	ToAddressSpaceResponseOutputWithContext(context.Context) AddressSpaceResponseOutput
}

type AddressSpaceResponseArgs struct {
	AddressPrefixes pulumi.StringArrayInput `pulumi:"addressPrefixes"`
}

func (AddressSpaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpaceResponse)(nil)).Elem()
}

func (i AddressSpaceResponseArgs) ToAddressSpaceResponseOutput() AddressSpaceResponseOutput {
	return i.ToAddressSpaceResponseOutputWithContext(context.Background())
}

func (i AddressSpaceResponseArgs) ToAddressSpaceResponseOutputWithContext(ctx context.Context) AddressSpaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceResponseOutput)
}

func (i AddressSpaceResponseArgs) ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput {
	return i.ToAddressSpaceResponsePtrOutputWithContext(context.Background())
}

func (i AddressSpaceResponseArgs) ToAddressSpaceResponsePtrOutputWithContext(ctx context.Context) AddressSpaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceResponseOutput).ToAddressSpaceResponsePtrOutputWithContext(ctx)
}









type AddressSpaceResponsePtrInput interface {
	pulumi.Input

	ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput
	ToAddressSpaceResponsePtrOutputWithContext(context.Context) AddressSpaceResponsePtrOutput
}

type addressSpaceResponsePtrType AddressSpaceResponseArgs

func AddressSpaceResponsePtr(v *AddressSpaceResponseArgs) AddressSpaceResponsePtrInput {
	return (*addressSpaceResponsePtrType)(v)
}

func (*addressSpaceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpaceResponse)(nil)).Elem()
}

func (i *addressSpaceResponsePtrType) ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput {
	return i.ToAddressSpaceResponsePtrOutputWithContext(context.Background())
}

func (i *addressSpaceResponsePtrType) ToAddressSpaceResponsePtrOutputWithContext(ctx context.Context) AddressSpaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceResponsePtrOutput)
}

type AddressSpaceResponseOutput struct{ *pulumi.OutputState }

func (AddressSpaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpaceResponse)(nil)).Elem()
}

func (o AddressSpaceResponseOutput) ToAddressSpaceResponseOutput() AddressSpaceResponseOutput {
	return o
}

func (o AddressSpaceResponseOutput) ToAddressSpaceResponseOutputWithContext(ctx context.Context) AddressSpaceResponseOutput {
	return o
}

func (o AddressSpaceResponseOutput) ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput {
	return o.ToAddressSpaceResponsePtrOutputWithContext(context.Background())
}

func (o AddressSpaceResponseOutput) ToAddressSpaceResponsePtrOutputWithContext(ctx context.Context) AddressSpaceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressSpaceResponse) *AddressSpaceResponse {
		return &v
	}).(AddressSpaceResponsePtrOutput)
}

func (o AddressSpaceResponseOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddressSpaceResponse) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

type AddressSpaceResponsePtrOutput struct{ *pulumi.OutputState }

func (AddressSpaceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpaceResponse)(nil)).Elem()
}

func (o AddressSpaceResponsePtrOutput) ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput {
	return o
}

func (o AddressSpaceResponsePtrOutput) ToAddressSpaceResponsePtrOutputWithContext(ctx context.Context) AddressSpaceResponsePtrOutput {
	return o
}

func (o AddressSpaceResponsePtrOutput) Elem() AddressSpaceResponseOutput {
	return o.ApplyT(func(v *AddressSpaceResponse) AddressSpaceResponse {
		if v != nil {
			return *v
		}
		var ret AddressSpaceResponse
		return ret
	}).(AddressSpaceResponseOutput)
}

func (o AddressSpaceResponsePtrOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AddressSpaceResponse) []string {
		if v == nil {
			return nil
		}
		return v.AddressPrefixes
	}).(pulumi.StringArrayOutput)
}

type ApplicationGatewayBackendAddress struct {
	Fqdn      *string `pulumi:"fqdn"`
	IpAddress *string `pulumi:"ipAddress"`
}





type ApplicationGatewayBackendAddressInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressOutput() ApplicationGatewayBackendAddressOutput
	ToApplicationGatewayBackendAddressOutputWithContext(context.Context) ApplicationGatewayBackendAddressOutput
}

type ApplicationGatewayBackendAddressArgs struct {
	Fqdn      pulumi.StringPtrInput `pulumi:"fqdn"`
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
}

func (ApplicationGatewayBackendAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddress)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressArgs) ToApplicationGatewayBackendAddressOutput() ApplicationGatewayBackendAddressOutput {
	return i.ToApplicationGatewayBackendAddressOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressArgs) ToApplicationGatewayBackendAddressOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressOutput)
}





type ApplicationGatewayBackendAddressArrayInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressArrayOutput() ApplicationGatewayBackendAddressArrayOutput
	ToApplicationGatewayBackendAddressArrayOutputWithContext(context.Context) ApplicationGatewayBackendAddressArrayOutput
}

type ApplicationGatewayBackendAddressArray []ApplicationGatewayBackendAddressInput

func (ApplicationGatewayBackendAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddress)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressArray) ToApplicationGatewayBackendAddressArrayOutput() ApplicationGatewayBackendAddressArrayOutput {
	return i.ToApplicationGatewayBackendAddressArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressArray) ToApplicationGatewayBackendAddressArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressArrayOutput)
}

type ApplicationGatewayBackendAddressOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddress)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressOutput) ToApplicationGatewayBackendAddressOutput() ApplicationGatewayBackendAddressOutput {
	return o
}

func (o ApplicationGatewayBackendAddressOutput) ToApplicationGatewayBackendAddressOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressOutput {
	return o
}

func (o ApplicationGatewayBackendAddressOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddress) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddress) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayBackendAddressArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddress)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressArrayOutput) ToApplicationGatewayBackendAddressArrayOutput() ApplicationGatewayBackendAddressArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressArrayOutput) ToApplicationGatewayBackendAddressArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendAddress {
		return vs[0].([]ApplicationGatewayBackendAddress)[vs[1].(int)]
	}).(ApplicationGatewayBackendAddressOutput)
}

type ApplicationGatewayBackendAddressPool struct {
	BackendAddresses        []ApplicationGatewayBackendAddress `pulumi:"backendAddresses"`
	BackendIPConfigurations []SubResource                      `pulumi:"backendIPConfigurations"`
	Etag                    *string                            `pulumi:"etag"`
	Id                      *string                            `pulumi:"id"`
	Name                    *string                            `pulumi:"name"`
	ProvisioningState       *string                            `pulumi:"provisioningState"`
}





type ApplicationGatewayBackendAddressPoolInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressPoolOutput() ApplicationGatewayBackendAddressPoolOutput
	ToApplicationGatewayBackendAddressPoolOutputWithContext(context.Context) ApplicationGatewayBackendAddressPoolOutput
}

type ApplicationGatewayBackendAddressPoolArgs struct {
	BackendAddresses        ApplicationGatewayBackendAddressArrayInput `pulumi:"backendAddresses"`
	BackendIPConfigurations SubResourceArrayInput                      `pulumi:"backendIPConfigurations"`
	Etag                    pulumi.StringPtrInput                      `pulumi:"etag"`
	Id                      pulumi.StringPtrInput                      `pulumi:"id"`
	Name                    pulumi.StringPtrInput                      `pulumi:"name"`
	ProvisioningState       pulumi.StringPtrInput                      `pulumi:"provisioningState"`
}

func (ApplicationGatewayBackendAddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddressPool)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressPoolArgs) ToApplicationGatewayBackendAddressPoolOutput() ApplicationGatewayBackendAddressPoolOutput {
	return i.ToApplicationGatewayBackendAddressPoolOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressPoolArgs) ToApplicationGatewayBackendAddressPoolOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressPoolOutput)
}





type ApplicationGatewayBackendAddressPoolArrayInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressPoolArrayOutput() ApplicationGatewayBackendAddressPoolArrayOutput
	ToApplicationGatewayBackendAddressPoolArrayOutputWithContext(context.Context) ApplicationGatewayBackendAddressPoolArrayOutput
}

type ApplicationGatewayBackendAddressPoolArray []ApplicationGatewayBackendAddressPoolInput

func (ApplicationGatewayBackendAddressPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddressPool)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressPoolArray) ToApplicationGatewayBackendAddressPoolArrayOutput() ApplicationGatewayBackendAddressPoolArrayOutput {
	return i.ToApplicationGatewayBackendAddressPoolArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressPoolArray) ToApplicationGatewayBackendAddressPoolArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressPoolArrayOutput)
}

type ApplicationGatewayBackendAddressPoolOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddressPool)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressPoolOutput) ToApplicationGatewayBackendAddressPoolOutput() ApplicationGatewayBackendAddressPoolOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolOutput) ToApplicationGatewayBackendAddressPoolOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolOutput) BackendAddresses() ApplicationGatewayBackendAddressArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) []ApplicationGatewayBackendAddress {
		return v.BackendAddresses
	}).(ApplicationGatewayBackendAddressArrayOutput)
}

func (o ApplicationGatewayBackendAddressPoolOutput) BackendIPConfigurations() SubResourceArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) []SubResource { return v.BackendIPConfigurations }).(SubResourceArrayOutput)
}

func (o ApplicationGatewayBackendAddressPoolOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayBackendAddressPoolArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddressPool)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressPoolArrayOutput) ToApplicationGatewayBackendAddressPoolArrayOutput() ApplicationGatewayBackendAddressPoolArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolArrayOutput) ToApplicationGatewayBackendAddressPoolArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendAddressPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendAddressPool {
		return vs[0].([]ApplicationGatewayBackendAddressPool)[vs[1].(int)]
	}).(ApplicationGatewayBackendAddressPoolOutput)
}

type ApplicationGatewayBackendAddressPoolResponse struct {
	BackendAddresses        []ApplicationGatewayBackendAddressResponse `pulumi:"backendAddresses"`
	BackendIPConfigurations []SubResourceResponse                      `pulumi:"backendIPConfigurations"`
	Etag                    *string                                    `pulumi:"etag"`
	Id                      *string                                    `pulumi:"id"`
	Name                    *string                                    `pulumi:"name"`
	ProvisioningState       *string                                    `pulumi:"provisioningState"`
}





type ApplicationGatewayBackendAddressPoolResponseInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressPoolResponseOutput() ApplicationGatewayBackendAddressPoolResponseOutput
	ToApplicationGatewayBackendAddressPoolResponseOutputWithContext(context.Context) ApplicationGatewayBackendAddressPoolResponseOutput
}

type ApplicationGatewayBackendAddressPoolResponseArgs struct {
	BackendAddresses        ApplicationGatewayBackendAddressResponseArrayInput `pulumi:"backendAddresses"`
	BackendIPConfigurations SubResourceResponseArrayInput                      `pulumi:"backendIPConfigurations"`
	Etag                    pulumi.StringPtrInput                              `pulumi:"etag"`
	Id                      pulumi.StringPtrInput                              `pulumi:"id"`
	Name                    pulumi.StringPtrInput                              `pulumi:"name"`
	ProvisioningState       pulumi.StringPtrInput                              `pulumi:"provisioningState"`
}

func (ApplicationGatewayBackendAddressPoolResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddressPoolResponse)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressPoolResponseArgs) ToApplicationGatewayBackendAddressPoolResponseOutput() ApplicationGatewayBackendAddressPoolResponseOutput {
	return i.ToApplicationGatewayBackendAddressPoolResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressPoolResponseArgs) ToApplicationGatewayBackendAddressPoolResponseOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressPoolResponseOutput)
}





type ApplicationGatewayBackendAddressPoolResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressPoolResponseArrayOutput() ApplicationGatewayBackendAddressPoolResponseArrayOutput
	ToApplicationGatewayBackendAddressPoolResponseArrayOutputWithContext(context.Context) ApplicationGatewayBackendAddressPoolResponseArrayOutput
}

type ApplicationGatewayBackendAddressPoolResponseArray []ApplicationGatewayBackendAddressPoolResponseInput

func (ApplicationGatewayBackendAddressPoolResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddressPoolResponse)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressPoolResponseArray) ToApplicationGatewayBackendAddressPoolResponseArrayOutput() ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return i.ToApplicationGatewayBackendAddressPoolResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressPoolResponseArray) ToApplicationGatewayBackendAddressPoolResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressPoolResponseArrayOutput)
}

type ApplicationGatewayBackendAddressPoolResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddressPoolResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) ToApplicationGatewayBackendAddressPoolResponseOutput() ApplicationGatewayBackendAddressPoolResponseOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) ToApplicationGatewayBackendAddressPoolResponseOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolResponseOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) BackendAddresses() ApplicationGatewayBackendAddressResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) []ApplicationGatewayBackendAddressResponse {
		return v.BackendAddresses
	}).(ApplicationGatewayBackendAddressResponseArrayOutput)
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) BackendIPConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) []SubResourceResponse {
		return v.BackendIPConfigurations
	}).(SubResourceResponseArrayOutput)
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayBackendAddressPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddressPoolResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressPoolResponseArrayOutput) ToApplicationGatewayBackendAddressPoolResponseArrayOutput() ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolResponseArrayOutput) ToApplicationGatewayBackendAddressPoolResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendAddressPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendAddressPoolResponse {
		return vs[0].([]ApplicationGatewayBackendAddressPoolResponse)[vs[1].(int)]
	}).(ApplicationGatewayBackendAddressPoolResponseOutput)
}

type ApplicationGatewayBackendAddressResponse struct {
	Fqdn      *string `pulumi:"fqdn"`
	IpAddress *string `pulumi:"ipAddress"`
}





type ApplicationGatewayBackendAddressResponseInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressResponseOutput() ApplicationGatewayBackendAddressResponseOutput
	ToApplicationGatewayBackendAddressResponseOutputWithContext(context.Context) ApplicationGatewayBackendAddressResponseOutput
}

type ApplicationGatewayBackendAddressResponseArgs struct {
	Fqdn      pulumi.StringPtrInput `pulumi:"fqdn"`
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
}

func (ApplicationGatewayBackendAddressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddressResponse)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressResponseArgs) ToApplicationGatewayBackendAddressResponseOutput() ApplicationGatewayBackendAddressResponseOutput {
	return i.ToApplicationGatewayBackendAddressResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressResponseArgs) ToApplicationGatewayBackendAddressResponseOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressResponseOutput)
}





type ApplicationGatewayBackendAddressResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressResponseArrayOutput() ApplicationGatewayBackendAddressResponseArrayOutput
	ToApplicationGatewayBackendAddressResponseArrayOutputWithContext(context.Context) ApplicationGatewayBackendAddressResponseArrayOutput
}

type ApplicationGatewayBackendAddressResponseArray []ApplicationGatewayBackendAddressResponseInput

func (ApplicationGatewayBackendAddressResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddressResponse)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressResponseArray) ToApplicationGatewayBackendAddressResponseArrayOutput() ApplicationGatewayBackendAddressResponseArrayOutput {
	return i.ToApplicationGatewayBackendAddressResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressResponseArray) ToApplicationGatewayBackendAddressResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressResponseArrayOutput)
}

type ApplicationGatewayBackendAddressResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddressResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressResponseOutput) ToApplicationGatewayBackendAddressResponseOutput() ApplicationGatewayBackendAddressResponseOutput {
	return o
}

func (o ApplicationGatewayBackendAddressResponseOutput) ToApplicationGatewayBackendAddressResponseOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressResponseOutput {
	return o
}

func (o ApplicationGatewayBackendAddressResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayBackendAddressResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddressResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressResponseArrayOutput) ToApplicationGatewayBackendAddressResponseArrayOutput() ApplicationGatewayBackendAddressResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressResponseArrayOutput) ToApplicationGatewayBackendAddressResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendAddressResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendAddressResponse {
		return vs[0].([]ApplicationGatewayBackendAddressResponse)[vs[1].(int)]
	}).(ApplicationGatewayBackendAddressResponseOutput)
}

type ApplicationGatewayBackendHttpSettings struct {
	CookieBasedAffinity *string `pulumi:"cookieBasedAffinity"`
	Etag                *string `pulumi:"etag"`
	Id                  *string `pulumi:"id"`
	Name                *string `pulumi:"name"`
	Port                *int    `pulumi:"port"`
	Protocol            *string `pulumi:"protocol"`
	ProvisioningState   *string `pulumi:"provisioningState"`
}





type ApplicationGatewayBackendHttpSettingsInput interface {
	pulumi.Input

	ToApplicationGatewayBackendHttpSettingsOutput() ApplicationGatewayBackendHttpSettingsOutput
	ToApplicationGatewayBackendHttpSettingsOutputWithContext(context.Context) ApplicationGatewayBackendHttpSettingsOutput
}

type ApplicationGatewayBackendHttpSettingsArgs struct {
	CookieBasedAffinity pulumi.StringPtrInput `pulumi:"cookieBasedAffinity"`
	Etag                pulumi.StringPtrInput `pulumi:"etag"`
	Id                  pulumi.StringPtrInput `pulumi:"id"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	Port                pulumi.IntPtrInput    `pulumi:"port"`
	Protocol            pulumi.StringPtrInput `pulumi:"protocol"`
	ProvisioningState   pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (ApplicationGatewayBackendHttpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendHttpSettings)(nil)).Elem()
}

func (i ApplicationGatewayBackendHttpSettingsArgs) ToApplicationGatewayBackendHttpSettingsOutput() ApplicationGatewayBackendHttpSettingsOutput {
	return i.ToApplicationGatewayBackendHttpSettingsOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendHttpSettingsArgs) ToApplicationGatewayBackendHttpSettingsOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendHttpSettingsOutput)
}





type ApplicationGatewayBackendHttpSettingsArrayInput interface {
	pulumi.Input

	ToApplicationGatewayBackendHttpSettingsArrayOutput() ApplicationGatewayBackendHttpSettingsArrayOutput
	ToApplicationGatewayBackendHttpSettingsArrayOutputWithContext(context.Context) ApplicationGatewayBackendHttpSettingsArrayOutput
}

type ApplicationGatewayBackendHttpSettingsArray []ApplicationGatewayBackendHttpSettingsInput

func (ApplicationGatewayBackendHttpSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendHttpSettings)(nil)).Elem()
}

func (i ApplicationGatewayBackendHttpSettingsArray) ToApplicationGatewayBackendHttpSettingsArrayOutput() ApplicationGatewayBackendHttpSettingsArrayOutput {
	return i.ToApplicationGatewayBackendHttpSettingsArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendHttpSettingsArray) ToApplicationGatewayBackendHttpSettingsArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendHttpSettingsArrayOutput)
}

type ApplicationGatewayBackendHttpSettingsOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendHttpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendHttpSettings)(nil)).Elem()
}

func (o ApplicationGatewayBackendHttpSettingsOutput) ToApplicationGatewayBackendHttpSettingsOutput() ApplicationGatewayBackendHttpSettingsOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsOutput) ToApplicationGatewayBackendHttpSettingsOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsOutput) CookieBasedAffinity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.CookieBasedAffinity }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayBackendHttpSettingsArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendHttpSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendHttpSettings)(nil)).Elem()
}

func (o ApplicationGatewayBackendHttpSettingsArrayOutput) ToApplicationGatewayBackendHttpSettingsArrayOutput() ApplicationGatewayBackendHttpSettingsArrayOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsArrayOutput) ToApplicationGatewayBackendHttpSettingsArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsArrayOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendHttpSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendHttpSettings {
		return vs[0].([]ApplicationGatewayBackendHttpSettings)[vs[1].(int)]
	}).(ApplicationGatewayBackendHttpSettingsOutput)
}

type ApplicationGatewayBackendHttpSettingsResponse struct {
	CookieBasedAffinity *string `pulumi:"cookieBasedAffinity"`
	Etag                *string `pulumi:"etag"`
	Id                  *string `pulumi:"id"`
	Name                *string `pulumi:"name"`
	Port                *int    `pulumi:"port"`
	Protocol            *string `pulumi:"protocol"`
	ProvisioningState   *string `pulumi:"provisioningState"`
}





type ApplicationGatewayBackendHttpSettingsResponseInput interface {
	pulumi.Input

	ToApplicationGatewayBackendHttpSettingsResponseOutput() ApplicationGatewayBackendHttpSettingsResponseOutput
	ToApplicationGatewayBackendHttpSettingsResponseOutputWithContext(context.Context) ApplicationGatewayBackendHttpSettingsResponseOutput
}

type ApplicationGatewayBackendHttpSettingsResponseArgs struct {
	CookieBasedAffinity pulumi.StringPtrInput `pulumi:"cookieBasedAffinity"`
	Etag                pulumi.StringPtrInput `pulumi:"etag"`
	Id                  pulumi.StringPtrInput `pulumi:"id"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	Port                pulumi.IntPtrInput    `pulumi:"port"`
	Protocol            pulumi.StringPtrInput `pulumi:"protocol"`
	ProvisioningState   pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (ApplicationGatewayBackendHttpSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendHttpSettingsResponse)(nil)).Elem()
}

func (i ApplicationGatewayBackendHttpSettingsResponseArgs) ToApplicationGatewayBackendHttpSettingsResponseOutput() ApplicationGatewayBackendHttpSettingsResponseOutput {
	return i.ToApplicationGatewayBackendHttpSettingsResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendHttpSettingsResponseArgs) ToApplicationGatewayBackendHttpSettingsResponseOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendHttpSettingsResponseOutput)
}





type ApplicationGatewayBackendHttpSettingsResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayBackendHttpSettingsResponseArrayOutput() ApplicationGatewayBackendHttpSettingsResponseArrayOutput
	ToApplicationGatewayBackendHttpSettingsResponseArrayOutputWithContext(context.Context) ApplicationGatewayBackendHttpSettingsResponseArrayOutput
}

type ApplicationGatewayBackendHttpSettingsResponseArray []ApplicationGatewayBackendHttpSettingsResponseInput

func (ApplicationGatewayBackendHttpSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendHttpSettingsResponse)(nil)).Elem()
}

func (i ApplicationGatewayBackendHttpSettingsResponseArray) ToApplicationGatewayBackendHttpSettingsResponseArrayOutput() ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
	return i.ToApplicationGatewayBackendHttpSettingsResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendHttpSettingsResponseArray) ToApplicationGatewayBackendHttpSettingsResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendHttpSettingsResponseArrayOutput)
}

type ApplicationGatewayBackendHttpSettingsResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendHttpSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendHttpSettingsResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) ToApplicationGatewayBackendHttpSettingsResponseOutput() ApplicationGatewayBackendHttpSettingsResponseOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) ToApplicationGatewayBackendHttpSettingsResponseOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsResponseOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) CookieBasedAffinity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.CookieBasedAffinity }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayBackendHttpSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendHttpSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendHttpSettingsResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendHttpSettingsResponseArrayOutput) ToApplicationGatewayBackendHttpSettingsResponseArrayOutput() ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsResponseArrayOutput) ToApplicationGatewayBackendHttpSettingsResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendHttpSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendHttpSettingsResponse {
		return vs[0].([]ApplicationGatewayBackendHttpSettingsResponse)[vs[1].(int)]
	}).(ApplicationGatewayBackendHttpSettingsResponseOutput)
}

type ApplicationGatewayFrontendIPConfiguration struct {
	Etag                      *string      `pulumi:"etag"`
	Id                        *string      `pulumi:"id"`
	Name                      *string      `pulumi:"name"`
	PrivateIPAddress          *string      `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string      `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string      `pulumi:"provisioningState"`
	PublicIPAddress           *SubResource `pulumi:"publicIPAddress"`
	Subnet                    *SubResource `pulumi:"subnet"`
}





type ApplicationGatewayFrontendIPConfigurationInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendIPConfigurationOutput() ApplicationGatewayFrontendIPConfigurationOutput
	ToApplicationGatewayFrontendIPConfigurationOutputWithContext(context.Context) ApplicationGatewayFrontendIPConfigurationOutput
}

type ApplicationGatewayFrontendIPConfigurationArgs struct {
	Etag                      pulumi.StringPtrInput `pulumi:"etag"`
	Id                        pulumi.StringPtrInput `pulumi:"id"`
	Name                      pulumi.StringPtrInput `pulumi:"name"`
	PrivateIPAddress          pulumi.StringPtrInput `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringPtrInput `pulumi:"provisioningState"`
	PublicIPAddress           SubResourcePtrInput   `pulumi:"publicIPAddress"`
	Subnet                    SubResourcePtrInput   `pulumi:"subnet"`
}

func (ApplicationGatewayFrontendIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendIPConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayFrontendIPConfigurationArgs) ToApplicationGatewayFrontendIPConfigurationOutput() ApplicationGatewayFrontendIPConfigurationOutput {
	return i.ToApplicationGatewayFrontendIPConfigurationOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendIPConfigurationArgs) ToApplicationGatewayFrontendIPConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendIPConfigurationOutput)
}





type ApplicationGatewayFrontendIPConfigurationArrayInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendIPConfigurationArrayOutput() ApplicationGatewayFrontendIPConfigurationArrayOutput
	ToApplicationGatewayFrontendIPConfigurationArrayOutputWithContext(context.Context) ApplicationGatewayFrontendIPConfigurationArrayOutput
}

type ApplicationGatewayFrontendIPConfigurationArray []ApplicationGatewayFrontendIPConfigurationInput

func (ApplicationGatewayFrontendIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendIPConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayFrontendIPConfigurationArray) ToApplicationGatewayFrontendIPConfigurationArrayOutput() ApplicationGatewayFrontendIPConfigurationArrayOutput {
	return i.ToApplicationGatewayFrontendIPConfigurationArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendIPConfigurationArray) ToApplicationGatewayFrontendIPConfigurationArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendIPConfigurationArrayOutput)
}

type ApplicationGatewayFrontendIPConfigurationOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendIPConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) ToApplicationGatewayFrontendIPConfigurationOutput() ApplicationGatewayFrontendIPConfigurationOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) ToApplicationGatewayFrontendIPConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) PublicIPAddress() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *SubResource { return v.PublicIPAddress }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type ApplicationGatewayFrontendIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendIPConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayFrontendIPConfigurationArrayOutput) ToApplicationGatewayFrontendIPConfigurationArrayOutput() ApplicationGatewayFrontendIPConfigurationArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationArrayOutput) ToApplicationGatewayFrontendIPConfigurationArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFrontendIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFrontendIPConfiguration {
		return vs[0].([]ApplicationGatewayFrontendIPConfiguration)[vs[1].(int)]
	}).(ApplicationGatewayFrontendIPConfigurationOutput)
}

type ApplicationGatewayFrontendIPConfigurationResponse struct {
	Etag                      *string              `pulumi:"etag"`
	Id                        *string              `pulumi:"id"`
	Name                      *string              `pulumi:"name"`
	PrivateIPAddress          *string              `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string              `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string              `pulumi:"provisioningState"`
	PublicIPAddress           *SubResourceResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubResourceResponse `pulumi:"subnet"`
}





type ApplicationGatewayFrontendIPConfigurationResponseInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendIPConfigurationResponseOutput() ApplicationGatewayFrontendIPConfigurationResponseOutput
	ToApplicationGatewayFrontendIPConfigurationResponseOutputWithContext(context.Context) ApplicationGatewayFrontendIPConfigurationResponseOutput
}

type ApplicationGatewayFrontendIPConfigurationResponseArgs struct {
	Etag                      pulumi.StringPtrInput       `pulumi:"etag"`
	Id                        pulumi.StringPtrInput       `pulumi:"id"`
	Name                      pulumi.StringPtrInput       `pulumi:"name"`
	PrivateIPAddress          pulumi.StringPtrInput       `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput       `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringPtrInput       `pulumi:"provisioningState"`
	PublicIPAddress           SubResourceResponsePtrInput `pulumi:"publicIPAddress"`
	Subnet                    SubResourceResponsePtrInput `pulumi:"subnet"`
}

func (ApplicationGatewayFrontendIPConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendIPConfigurationResponse)(nil)).Elem()
}

func (i ApplicationGatewayFrontendIPConfigurationResponseArgs) ToApplicationGatewayFrontendIPConfigurationResponseOutput() ApplicationGatewayFrontendIPConfigurationResponseOutput {
	return i.ToApplicationGatewayFrontendIPConfigurationResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendIPConfigurationResponseArgs) ToApplicationGatewayFrontendIPConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendIPConfigurationResponseOutput)
}





type ApplicationGatewayFrontendIPConfigurationResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendIPConfigurationResponseArrayOutput() ApplicationGatewayFrontendIPConfigurationResponseArrayOutput
	ToApplicationGatewayFrontendIPConfigurationResponseArrayOutputWithContext(context.Context) ApplicationGatewayFrontendIPConfigurationResponseArrayOutput
}

type ApplicationGatewayFrontendIPConfigurationResponseArray []ApplicationGatewayFrontendIPConfigurationResponseInput

func (ApplicationGatewayFrontendIPConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendIPConfigurationResponse)(nil)).Elem()
}

func (i ApplicationGatewayFrontendIPConfigurationResponseArray) ToApplicationGatewayFrontendIPConfigurationResponseArrayOutput() ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
	return i.ToApplicationGatewayFrontendIPConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendIPConfigurationResponseArray) ToApplicationGatewayFrontendIPConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendIPConfigurationResponseArrayOutput)
}

type ApplicationGatewayFrontendIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendIPConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) ToApplicationGatewayFrontendIPConfigurationResponseOutput() ApplicationGatewayFrontendIPConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) ToApplicationGatewayFrontendIPConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) PublicIPAddress() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *SubResourceResponse {
		return v.PublicIPAddress
	}).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type ApplicationGatewayFrontendIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendIPConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayFrontendIPConfigurationResponseArrayOutput) ToApplicationGatewayFrontendIPConfigurationResponseArrayOutput() ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationResponseArrayOutput) ToApplicationGatewayFrontendIPConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFrontendIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFrontendIPConfigurationResponse {
		return vs[0].([]ApplicationGatewayFrontendIPConfigurationResponse)[vs[1].(int)]
	}).(ApplicationGatewayFrontendIPConfigurationResponseOutput)
}

type ApplicationGatewayFrontendPort struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	Port              *int    `pulumi:"port"`
	ProvisioningState *string `pulumi:"provisioningState"`
}





type ApplicationGatewayFrontendPortInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendPortOutput() ApplicationGatewayFrontendPortOutput
	ToApplicationGatewayFrontendPortOutputWithContext(context.Context) ApplicationGatewayFrontendPortOutput
}

type ApplicationGatewayFrontendPortArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	Port              pulumi.IntPtrInput    `pulumi:"port"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (ApplicationGatewayFrontendPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendPort)(nil)).Elem()
}

func (i ApplicationGatewayFrontendPortArgs) ToApplicationGatewayFrontendPortOutput() ApplicationGatewayFrontendPortOutput {
	return i.ToApplicationGatewayFrontendPortOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendPortArgs) ToApplicationGatewayFrontendPortOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendPortOutput)
}





type ApplicationGatewayFrontendPortArrayInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendPortArrayOutput() ApplicationGatewayFrontendPortArrayOutput
	ToApplicationGatewayFrontendPortArrayOutputWithContext(context.Context) ApplicationGatewayFrontendPortArrayOutput
}

type ApplicationGatewayFrontendPortArray []ApplicationGatewayFrontendPortInput

func (ApplicationGatewayFrontendPortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendPort)(nil)).Elem()
}

func (i ApplicationGatewayFrontendPortArray) ToApplicationGatewayFrontendPortArrayOutput() ApplicationGatewayFrontendPortArrayOutput {
	return i.ToApplicationGatewayFrontendPortArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendPortArray) ToApplicationGatewayFrontendPortArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendPortArrayOutput)
}

type ApplicationGatewayFrontendPortOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendPort)(nil)).Elem()
}

func (o ApplicationGatewayFrontendPortOutput) ToApplicationGatewayFrontendPortOutput() ApplicationGatewayFrontendPortOutput {
	return o
}

func (o ApplicationGatewayFrontendPortOutput) ToApplicationGatewayFrontendPortOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortOutput {
	return o
}

func (o ApplicationGatewayFrontendPortOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayFrontendPortOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayFrontendPortArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendPortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendPort)(nil)).Elem()
}

func (o ApplicationGatewayFrontendPortArrayOutput) ToApplicationGatewayFrontendPortArrayOutput() ApplicationGatewayFrontendPortArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendPortArrayOutput) ToApplicationGatewayFrontendPortArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendPortArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFrontendPortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFrontendPort {
		return vs[0].([]ApplicationGatewayFrontendPort)[vs[1].(int)]
	}).(ApplicationGatewayFrontendPortOutput)
}

type ApplicationGatewayFrontendPortResponse struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	Port              *int    `pulumi:"port"`
	ProvisioningState *string `pulumi:"provisioningState"`
}





type ApplicationGatewayFrontendPortResponseInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendPortResponseOutput() ApplicationGatewayFrontendPortResponseOutput
	ToApplicationGatewayFrontendPortResponseOutputWithContext(context.Context) ApplicationGatewayFrontendPortResponseOutput
}

type ApplicationGatewayFrontendPortResponseArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	Port              pulumi.IntPtrInput    `pulumi:"port"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (ApplicationGatewayFrontendPortResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendPortResponse)(nil)).Elem()
}

func (i ApplicationGatewayFrontendPortResponseArgs) ToApplicationGatewayFrontendPortResponseOutput() ApplicationGatewayFrontendPortResponseOutput {
	return i.ToApplicationGatewayFrontendPortResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendPortResponseArgs) ToApplicationGatewayFrontendPortResponseOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendPortResponseOutput)
}





type ApplicationGatewayFrontendPortResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendPortResponseArrayOutput() ApplicationGatewayFrontendPortResponseArrayOutput
	ToApplicationGatewayFrontendPortResponseArrayOutputWithContext(context.Context) ApplicationGatewayFrontendPortResponseArrayOutput
}

type ApplicationGatewayFrontendPortResponseArray []ApplicationGatewayFrontendPortResponseInput

func (ApplicationGatewayFrontendPortResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendPortResponse)(nil)).Elem()
}

func (i ApplicationGatewayFrontendPortResponseArray) ToApplicationGatewayFrontendPortResponseArrayOutput() ApplicationGatewayFrontendPortResponseArrayOutput {
	return i.ToApplicationGatewayFrontendPortResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendPortResponseArray) ToApplicationGatewayFrontendPortResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendPortResponseArrayOutput)
}

type ApplicationGatewayFrontendPortResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendPortResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendPortResponse)(nil)).Elem()
}

func (o ApplicationGatewayFrontendPortResponseOutput) ToApplicationGatewayFrontendPortResponseOutput() ApplicationGatewayFrontendPortResponseOutput {
	return o
}

func (o ApplicationGatewayFrontendPortResponseOutput) ToApplicationGatewayFrontendPortResponseOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortResponseOutput {
	return o
}

func (o ApplicationGatewayFrontendPortResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayFrontendPortResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayFrontendPortResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendPortResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendPortResponse)(nil)).Elem()
}

func (o ApplicationGatewayFrontendPortResponseArrayOutput) ToApplicationGatewayFrontendPortResponseArrayOutput() ApplicationGatewayFrontendPortResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendPortResponseArrayOutput) ToApplicationGatewayFrontendPortResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendPortResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFrontendPortResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFrontendPortResponse {
		return vs[0].([]ApplicationGatewayFrontendPortResponse)[vs[1].(int)]
	}).(ApplicationGatewayFrontendPortResponseOutput)
}

type ApplicationGatewayHttpListener struct {
	Etag                    *string      `pulumi:"etag"`
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPort            *SubResource `pulumi:"frontendPort"`
	Id                      *string      `pulumi:"id"`
	Name                    *string      `pulumi:"name"`
	Protocol                *string      `pulumi:"protocol"`
	ProvisioningState       *string      `pulumi:"provisioningState"`
	SslCertificate          *SubResource `pulumi:"sslCertificate"`
}





type ApplicationGatewayHttpListenerInput interface {
	pulumi.Input

	ToApplicationGatewayHttpListenerOutput() ApplicationGatewayHttpListenerOutput
	ToApplicationGatewayHttpListenerOutputWithContext(context.Context) ApplicationGatewayHttpListenerOutput
}

type ApplicationGatewayHttpListenerArgs struct {
	Etag                    pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPort            SubResourcePtrInput   `pulumi:"frontendPort"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	Protocol                pulumi.StringPtrInput `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput `pulumi:"provisioningState"`
	SslCertificate          SubResourcePtrInput   `pulumi:"sslCertificate"`
}

func (ApplicationGatewayHttpListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayHttpListener)(nil)).Elem()
}

func (i ApplicationGatewayHttpListenerArgs) ToApplicationGatewayHttpListenerOutput() ApplicationGatewayHttpListenerOutput {
	return i.ToApplicationGatewayHttpListenerOutputWithContext(context.Background())
}

func (i ApplicationGatewayHttpListenerArgs) ToApplicationGatewayHttpListenerOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayHttpListenerOutput)
}





type ApplicationGatewayHttpListenerArrayInput interface {
	pulumi.Input

	ToApplicationGatewayHttpListenerArrayOutput() ApplicationGatewayHttpListenerArrayOutput
	ToApplicationGatewayHttpListenerArrayOutputWithContext(context.Context) ApplicationGatewayHttpListenerArrayOutput
}

type ApplicationGatewayHttpListenerArray []ApplicationGatewayHttpListenerInput

func (ApplicationGatewayHttpListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayHttpListener)(nil)).Elem()
}

func (i ApplicationGatewayHttpListenerArray) ToApplicationGatewayHttpListenerArrayOutput() ApplicationGatewayHttpListenerArrayOutput {
	return i.ToApplicationGatewayHttpListenerArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayHttpListenerArray) ToApplicationGatewayHttpListenerArrayOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayHttpListenerArrayOutput)
}

type ApplicationGatewayHttpListenerOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayHttpListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayHttpListener)(nil)).Elem()
}

func (o ApplicationGatewayHttpListenerOutput) ToApplicationGatewayHttpListenerOutput() ApplicationGatewayHttpListenerOutput {
	return o
}

func (o ApplicationGatewayHttpListenerOutput) ToApplicationGatewayHttpListenerOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerOutput {
	return o
}

func (o ApplicationGatewayHttpListenerOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) FrontendIPConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *SubResource { return v.FrontendIPConfiguration }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) FrontendPort() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *SubResource { return v.FrontendPort }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) SslCertificate() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *SubResource { return v.SslCertificate }).(SubResourcePtrOutput)
}

type ApplicationGatewayHttpListenerArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayHttpListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayHttpListener)(nil)).Elem()
}

func (o ApplicationGatewayHttpListenerArrayOutput) ToApplicationGatewayHttpListenerArrayOutput() ApplicationGatewayHttpListenerArrayOutput {
	return o
}

func (o ApplicationGatewayHttpListenerArrayOutput) ToApplicationGatewayHttpListenerArrayOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerArrayOutput {
	return o
}

func (o ApplicationGatewayHttpListenerArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayHttpListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayHttpListener {
		return vs[0].([]ApplicationGatewayHttpListener)[vs[1].(int)]
	}).(ApplicationGatewayHttpListenerOutput)
}

type ApplicationGatewayHttpListenerResponse struct {
	Etag                    *string              `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	FrontendPort            *SubResourceResponse `pulumi:"frontendPort"`
	Id                      *string              `pulumi:"id"`
	Name                    *string              `pulumi:"name"`
	Protocol                *string              `pulumi:"protocol"`
	ProvisioningState       *string              `pulumi:"provisioningState"`
	SslCertificate          *SubResourceResponse `pulumi:"sslCertificate"`
}





type ApplicationGatewayHttpListenerResponseInput interface {
	pulumi.Input

	ToApplicationGatewayHttpListenerResponseOutput() ApplicationGatewayHttpListenerResponseOutput
	ToApplicationGatewayHttpListenerResponseOutputWithContext(context.Context) ApplicationGatewayHttpListenerResponseOutput
}

type ApplicationGatewayHttpListenerResponseArgs struct {
	Etag                    pulumi.StringPtrInput       `pulumi:"etag"`
	FrontendIPConfiguration SubResourceResponsePtrInput `pulumi:"frontendIPConfiguration"`
	FrontendPort            SubResourceResponsePtrInput `pulumi:"frontendPort"`
	Id                      pulumi.StringPtrInput       `pulumi:"id"`
	Name                    pulumi.StringPtrInput       `pulumi:"name"`
	Protocol                pulumi.StringPtrInput       `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput       `pulumi:"provisioningState"`
	SslCertificate          SubResourceResponsePtrInput `pulumi:"sslCertificate"`
}

func (ApplicationGatewayHttpListenerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayHttpListenerResponse)(nil)).Elem()
}

func (i ApplicationGatewayHttpListenerResponseArgs) ToApplicationGatewayHttpListenerResponseOutput() ApplicationGatewayHttpListenerResponseOutput {
	return i.ToApplicationGatewayHttpListenerResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayHttpListenerResponseArgs) ToApplicationGatewayHttpListenerResponseOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayHttpListenerResponseOutput)
}





type ApplicationGatewayHttpListenerResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayHttpListenerResponseArrayOutput() ApplicationGatewayHttpListenerResponseArrayOutput
	ToApplicationGatewayHttpListenerResponseArrayOutputWithContext(context.Context) ApplicationGatewayHttpListenerResponseArrayOutput
}

type ApplicationGatewayHttpListenerResponseArray []ApplicationGatewayHttpListenerResponseInput

func (ApplicationGatewayHttpListenerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayHttpListenerResponse)(nil)).Elem()
}

func (i ApplicationGatewayHttpListenerResponseArray) ToApplicationGatewayHttpListenerResponseArrayOutput() ApplicationGatewayHttpListenerResponseArrayOutput {
	return i.ToApplicationGatewayHttpListenerResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayHttpListenerResponseArray) ToApplicationGatewayHttpListenerResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayHttpListenerResponseArrayOutput)
}

type ApplicationGatewayHttpListenerResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayHttpListenerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayHttpListenerResponse)(nil)).Elem()
}

func (o ApplicationGatewayHttpListenerResponseOutput) ToApplicationGatewayHttpListenerResponseOutput() ApplicationGatewayHttpListenerResponseOutput {
	return o
}

func (o ApplicationGatewayHttpListenerResponseOutput) ToApplicationGatewayHttpListenerResponseOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerResponseOutput {
	return o
}

func (o ApplicationGatewayHttpListenerResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) FrontendPort() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *SubResourceResponse { return v.FrontendPort }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) SslCertificate() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *SubResourceResponse { return v.SslCertificate }).(SubResourceResponsePtrOutput)
}

type ApplicationGatewayHttpListenerResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayHttpListenerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayHttpListenerResponse)(nil)).Elem()
}

func (o ApplicationGatewayHttpListenerResponseArrayOutput) ToApplicationGatewayHttpListenerResponseArrayOutput() ApplicationGatewayHttpListenerResponseArrayOutput {
	return o
}

func (o ApplicationGatewayHttpListenerResponseArrayOutput) ToApplicationGatewayHttpListenerResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerResponseArrayOutput {
	return o
}

func (o ApplicationGatewayHttpListenerResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayHttpListenerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayHttpListenerResponse {
		return vs[0].([]ApplicationGatewayHttpListenerResponse)[vs[1].(int)]
	}).(ApplicationGatewayHttpListenerResponseOutput)
}

type ApplicationGatewayIPConfiguration struct {
	Etag              *string      `pulumi:"etag"`
	Id                *string      `pulumi:"id"`
	Name              *string      `pulumi:"name"`
	ProvisioningState *string      `pulumi:"provisioningState"`
	Subnet            *SubResource `pulumi:"subnet"`
}





type ApplicationGatewayIPConfigurationInput interface {
	pulumi.Input

	ToApplicationGatewayIPConfigurationOutput() ApplicationGatewayIPConfigurationOutput
	ToApplicationGatewayIPConfigurationOutputWithContext(context.Context) ApplicationGatewayIPConfigurationOutput
}

type ApplicationGatewayIPConfigurationArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	Subnet            SubResourcePtrInput   `pulumi:"subnet"`
}

func (ApplicationGatewayIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayIPConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayIPConfigurationArgs) ToApplicationGatewayIPConfigurationOutput() ApplicationGatewayIPConfigurationOutput {
	return i.ToApplicationGatewayIPConfigurationOutputWithContext(context.Background())
}

func (i ApplicationGatewayIPConfigurationArgs) ToApplicationGatewayIPConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayIPConfigurationOutput)
}





type ApplicationGatewayIPConfigurationArrayInput interface {
	pulumi.Input

	ToApplicationGatewayIPConfigurationArrayOutput() ApplicationGatewayIPConfigurationArrayOutput
	ToApplicationGatewayIPConfigurationArrayOutputWithContext(context.Context) ApplicationGatewayIPConfigurationArrayOutput
}

type ApplicationGatewayIPConfigurationArray []ApplicationGatewayIPConfigurationInput

func (ApplicationGatewayIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayIPConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayIPConfigurationArray) ToApplicationGatewayIPConfigurationArrayOutput() ApplicationGatewayIPConfigurationArrayOutput {
	return i.ToApplicationGatewayIPConfigurationArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayIPConfigurationArray) ToApplicationGatewayIPConfigurationArrayOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayIPConfigurationArrayOutput)
}

type ApplicationGatewayIPConfigurationOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayIPConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayIPConfigurationOutput) ToApplicationGatewayIPConfigurationOutput() ApplicationGatewayIPConfigurationOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationOutput) ToApplicationGatewayIPConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type ApplicationGatewayIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayIPConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayIPConfigurationArrayOutput) ToApplicationGatewayIPConfigurationArrayOutput() ApplicationGatewayIPConfigurationArrayOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationArrayOutput) ToApplicationGatewayIPConfigurationArrayOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationArrayOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayIPConfiguration {
		return vs[0].([]ApplicationGatewayIPConfiguration)[vs[1].(int)]
	}).(ApplicationGatewayIPConfigurationOutput)
}

type ApplicationGatewayIPConfigurationResponse struct {
	Etag              *string              `pulumi:"etag"`
	Id                *string              `pulumi:"id"`
	Name              *string              `pulumi:"name"`
	ProvisioningState *string              `pulumi:"provisioningState"`
	Subnet            *SubResourceResponse `pulumi:"subnet"`
}





type ApplicationGatewayIPConfigurationResponseInput interface {
	pulumi.Input

	ToApplicationGatewayIPConfigurationResponseOutput() ApplicationGatewayIPConfigurationResponseOutput
	ToApplicationGatewayIPConfigurationResponseOutputWithContext(context.Context) ApplicationGatewayIPConfigurationResponseOutput
}

type ApplicationGatewayIPConfigurationResponseArgs struct {
	Etag              pulumi.StringPtrInput       `pulumi:"etag"`
	Id                pulumi.StringPtrInput       `pulumi:"id"`
	Name              pulumi.StringPtrInput       `pulumi:"name"`
	ProvisioningState pulumi.StringPtrInput       `pulumi:"provisioningState"`
	Subnet            SubResourceResponsePtrInput `pulumi:"subnet"`
}

func (ApplicationGatewayIPConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayIPConfigurationResponse)(nil)).Elem()
}

func (i ApplicationGatewayIPConfigurationResponseArgs) ToApplicationGatewayIPConfigurationResponseOutput() ApplicationGatewayIPConfigurationResponseOutput {
	return i.ToApplicationGatewayIPConfigurationResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayIPConfigurationResponseArgs) ToApplicationGatewayIPConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayIPConfigurationResponseOutput)
}





type ApplicationGatewayIPConfigurationResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayIPConfigurationResponseArrayOutput() ApplicationGatewayIPConfigurationResponseArrayOutput
	ToApplicationGatewayIPConfigurationResponseArrayOutputWithContext(context.Context) ApplicationGatewayIPConfigurationResponseArrayOutput
}

type ApplicationGatewayIPConfigurationResponseArray []ApplicationGatewayIPConfigurationResponseInput

func (ApplicationGatewayIPConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayIPConfigurationResponse)(nil)).Elem()
}

func (i ApplicationGatewayIPConfigurationResponseArray) ToApplicationGatewayIPConfigurationResponseArrayOutput() ApplicationGatewayIPConfigurationResponseArrayOutput {
	return i.ToApplicationGatewayIPConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayIPConfigurationResponseArray) ToApplicationGatewayIPConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayIPConfigurationResponseArrayOutput)
}

type ApplicationGatewayIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayIPConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayIPConfigurationResponseOutput) ToApplicationGatewayIPConfigurationResponseOutput() ApplicationGatewayIPConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationResponseOutput) ToApplicationGatewayIPConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type ApplicationGatewayIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayIPConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayIPConfigurationResponseArrayOutput) ToApplicationGatewayIPConfigurationResponseArrayOutput() ApplicationGatewayIPConfigurationResponseArrayOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationResponseArrayOutput) ToApplicationGatewayIPConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationResponseArrayOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayIPConfigurationResponse {
		return vs[0].([]ApplicationGatewayIPConfigurationResponse)[vs[1].(int)]
	}).(ApplicationGatewayIPConfigurationResponseOutput)
}

type ApplicationGatewayRequestRoutingRule struct {
	BackendAddressPool  *SubResource `pulumi:"backendAddressPool"`
	BackendHttpSettings *SubResource `pulumi:"backendHttpSettings"`
	Etag                *string      `pulumi:"etag"`
	HttpListener        *SubResource `pulumi:"httpListener"`
	Id                  *string      `pulumi:"id"`
	Name                *string      `pulumi:"name"`
	ProvisioningState   *string      `pulumi:"provisioningState"`
	RuleType            *string      `pulumi:"ruleType"`
}





type ApplicationGatewayRequestRoutingRuleInput interface {
	pulumi.Input

	ToApplicationGatewayRequestRoutingRuleOutput() ApplicationGatewayRequestRoutingRuleOutput
	ToApplicationGatewayRequestRoutingRuleOutputWithContext(context.Context) ApplicationGatewayRequestRoutingRuleOutput
}

type ApplicationGatewayRequestRoutingRuleArgs struct {
	BackendAddressPool  SubResourcePtrInput   `pulumi:"backendAddressPool"`
	BackendHttpSettings SubResourcePtrInput   `pulumi:"backendHttpSettings"`
	Etag                pulumi.StringPtrInput `pulumi:"etag"`
	HttpListener        SubResourcePtrInput   `pulumi:"httpListener"`
	Id                  pulumi.StringPtrInput `pulumi:"id"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState   pulumi.StringPtrInput `pulumi:"provisioningState"`
	RuleType            pulumi.StringPtrInput `pulumi:"ruleType"`
}

func (ApplicationGatewayRequestRoutingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRequestRoutingRule)(nil)).Elem()
}

func (i ApplicationGatewayRequestRoutingRuleArgs) ToApplicationGatewayRequestRoutingRuleOutput() ApplicationGatewayRequestRoutingRuleOutput {
	return i.ToApplicationGatewayRequestRoutingRuleOutputWithContext(context.Background())
}

func (i ApplicationGatewayRequestRoutingRuleArgs) ToApplicationGatewayRequestRoutingRuleOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayRequestRoutingRuleOutput)
}





type ApplicationGatewayRequestRoutingRuleArrayInput interface {
	pulumi.Input

	ToApplicationGatewayRequestRoutingRuleArrayOutput() ApplicationGatewayRequestRoutingRuleArrayOutput
	ToApplicationGatewayRequestRoutingRuleArrayOutputWithContext(context.Context) ApplicationGatewayRequestRoutingRuleArrayOutput
}

type ApplicationGatewayRequestRoutingRuleArray []ApplicationGatewayRequestRoutingRuleInput

func (ApplicationGatewayRequestRoutingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRequestRoutingRule)(nil)).Elem()
}

func (i ApplicationGatewayRequestRoutingRuleArray) ToApplicationGatewayRequestRoutingRuleArrayOutput() ApplicationGatewayRequestRoutingRuleArrayOutput {
	return i.ToApplicationGatewayRequestRoutingRuleArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayRequestRoutingRuleArray) ToApplicationGatewayRequestRoutingRuleArrayOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayRequestRoutingRuleArrayOutput)
}

type ApplicationGatewayRequestRoutingRuleOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRequestRoutingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRequestRoutingRule)(nil)).Elem()
}

func (o ApplicationGatewayRequestRoutingRuleOutput) ToApplicationGatewayRequestRoutingRuleOutput() ApplicationGatewayRequestRoutingRuleOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleOutput) ToApplicationGatewayRequestRoutingRuleOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleOutput) BackendAddressPool() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *SubResource { return v.BackendAddressPool }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) BackendHttpSettings() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *SubResource { return v.BackendHttpSettings }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) HttpListener() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *SubResource { return v.HttpListener }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) RuleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.RuleType }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayRequestRoutingRuleArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRequestRoutingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRequestRoutingRule)(nil)).Elem()
}

func (o ApplicationGatewayRequestRoutingRuleArrayOutput) ToApplicationGatewayRequestRoutingRuleArrayOutput() ApplicationGatewayRequestRoutingRuleArrayOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleArrayOutput) ToApplicationGatewayRequestRoutingRuleArrayOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleArrayOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayRequestRoutingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayRequestRoutingRule {
		return vs[0].([]ApplicationGatewayRequestRoutingRule)[vs[1].(int)]
	}).(ApplicationGatewayRequestRoutingRuleOutput)
}

type ApplicationGatewayRequestRoutingRuleResponse struct {
	BackendAddressPool  *SubResourceResponse `pulumi:"backendAddressPool"`
	BackendHttpSettings *SubResourceResponse `pulumi:"backendHttpSettings"`
	Etag                *string              `pulumi:"etag"`
	HttpListener        *SubResourceResponse `pulumi:"httpListener"`
	Id                  *string              `pulumi:"id"`
	Name                *string              `pulumi:"name"`
	ProvisioningState   *string              `pulumi:"provisioningState"`
	RuleType            *string              `pulumi:"ruleType"`
}





type ApplicationGatewayRequestRoutingRuleResponseInput interface {
	pulumi.Input

	ToApplicationGatewayRequestRoutingRuleResponseOutput() ApplicationGatewayRequestRoutingRuleResponseOutput
	ToApplicationGatewayRequestRoutingRuleResponseOutputWithContext(context.Context) ApplicationGatewayRequestRoutingRuleResponseOutput
}

type ApplicationGatewayRequestRoutingRuleResponseArgs struct {
	BackendAddressPool  SubResourceResponsePtrInput `pulumi:"backendAddressPool"`
	BackendHttpSettings SubResourceResponsePtrInput `pulumi:"backendHttpSettings"`
	Etag                pulumi.StringPtrInput       `pulumi:"etag"`
	HttpListener        SubResourceResponsePtrInput `pulumi:"httpListener"`
	Id                  pulumi.StringPtrInput       `pulumi:"id"`
	Name                pulumi.StringPtrInput       `pulumi:"name"`
	ProvisioningState   pulumi.StringPtrInput       `pulumi:"provisioningState"`
	RuleType            pulumi.StringPtrInput       `pulumi:"ruleType"`
}

func (ApplicationGatewayRequestRoutingRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRequestRoutingRuleResponse)(nil)).Elem()
}

func (i ApplicationGatewayRequestRoutingRuleResponseArgs) ToApplicationGatewayRequestRoutingRuleResponseOutput() ApplicationGatewayRequestRoutingRuleResponseOutput {
	return i.ToApplicationGatewayRequestRoutingRuleResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayRequestRoutingRuleResponseArgs) ToApplicationGatewayRequestRoutingRuleResponseOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayRequestRoutingRuleResponseOutput)
}





type ApplicationGatewayRequestRoutingRuleResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayRequestRoutingRuleResponseArrayOutput() ApplicationGatewayRequestRoutingRuleResponseArrayOutput
	ToApplicationGatewayRequestRoutingRuleResponseArrayOutputWithContext(context.Context) ApplicationGatewayRequestRoutingRuleResponseArrayOutput
}

type ApplicationGatewayRequestRoutingRuleResponseArray []ApplicationGatewayRequestRoutingRuleResponseInput

func (ApplicationGatewayRequestRoutingRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRequestRoutingRuleResponse)(nil)).Elem()
}

func (i ApplicationGatewayRequestRoutingRuleResponseArray) ToApplicationGatewayRequestRoutingRuleResponseArrayOutput() ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return i.ToApplicationGatewayRequestRoutingRuleResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayRequestRoutingRuleResponseArray) ToApplicationGatewayRequestRoutingRuleResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayRequestRoutingRuleResponseArrayOutput)
}

type ApplicationGatewayRequestRoutingRuleResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRequestRoutingRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRequestRoutingRuleResponse)(nil)).Elem()
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) ToApplicationGatewayRequestRoutingRuleResponseOutput() ApplicationGatewayRequestRoutingRuleResponseOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) ToApplicationGatewayRequestRoutingRuleResponseOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleResponseOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) BackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) BackendHttpSettings() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *SubResourceResponse {
		return v.BackendHttpSettings
	}).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) HttpListener() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *SubResourceResponse { return v.HttpListener }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) RuleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.RuleType }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayRequestRoutingRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRequestRoutingRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRequestRoutingRuleResponse)(nil)).Elem()
}

func (o ApplicationGatewayRequestRoutingRuleResponseArrayOutput) ToApplicationGatewayRequestRoutingRuleResponseArrayOutput() ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleResponseArrayOutput) ToApplicationGatewayRequestRoutingRuleResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayRequestRoutingRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayRequestRoutingRuleResponse {
		return vs[0].([]ApplicationGatewayRequestRoutingRuleResponse)[vs[1].(int)]
	}).(ApplicationGatewayRequestRoutingRuleResponseOutput)
}

type ApplicationGatewaySku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ApplicationGatewaySkuInput interface {
	pulumi.Input

	ToApplicationGatewaySkuOutput() ApplicationGatewaySkuOutput
	ToApplicationGatewaySkuOutputWithContext(context.Context) ApplicationGatewaySkuOutput
}

type ApplicationGatewaySkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ApplicationGatewaySkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySku)(nil)).Elem()
}

func (i ApplicationGatewaySkuArgs) ToApplicationGatewaySkuOutput() ApplicationGatewaySkuOutput {
	return i.ToApplicationGatewaySkuOutputWithContext(context.Background())
}

func (i ApplicationGatewaySkuArgs) ToApplicationGatewaySkuOutputWithContext(ctx context.Context) ApplicationGatewaySkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySkuOutput)
}

func (i ApplicationGatewaySkuArgs) ToApplicationGatewaySkuPtrOutput() ApplicationGatewaySkuPtrOutput {
	return i.ToApplicationGatewaySkuPtrOutputWithContext(context.Background())
}

func (i ApplicationGatewaySkuArgs) ToApplicationGatewaySkuPtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySkuOutput).ToApplicationGatewaySkuPtrOutputWithContext(ctx)
}









type ApplicationGatewaySkuPtrInput interface {
	pulumi.Input

	ToApplicationGatewaySkuPtrOutput() ApplicationGatewaySkuPtrOutput
	ToApplicationGatewaySkuPtrOutputWithContext(context.Context) ApplicationGatewaySkuPtrOutput
}

type applicationGatewaySkuPtrType ApplicationGatewaySkuArgs

func ApplicationGatewaySkuPtr(v *ApplicationGatewaySkuArgs) ApplicationGatewaySkuPtrInput {
	return (*applicationGatewaySkuPtrType)(v)
}

func (*applicationGatewaySkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySku)(nil)).Elem()
}

func (i *applicationGatewaySkuPtrType) ToApplicationGatewaySkuPtrOutput() ApplicationGatewaySkuPtrOutput {
	return i.ToApplicationGatewaySkuPtrOutputWithContext(context.Background())
}

func (i *applicationGatewaySkuPtrType) ToApplicationGatewaySkuPtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySkuPtrOutput)
}

type ApplicationGatewaySkuOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySku)(nil)).Elem()
}

func (o ApplicationGatewaySkuOutput) ToApplicationGatewaySkuOutput() ApplicationGatewaySkuOutput {
	return o
}

func (o ApplicationGatewaySkuOutput) ToApplicationGatewaySkuOutputWithContext(ctx context.Context) ApplicationGatewaySkuOutput {
	return o
}

func (o ApplicationGatewaySkuOutput) ToApplicationGatewaySkuPtrOutput() ApplicationGatewaySkuPtrOutput {
	return o.ToApplicationGatewaySkuPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySkuOutput) ToApplicationGatewaySkuPtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySku) *ApplicationGatewaySku {
		return &v
	}).(ApplicationGatewaySkuPtrOutput)
}

func (o ApplicationGatewaySkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewaySkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySkuPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySku)(nil)).Elem()
}

func (o ApplicationGatewaySkuPtrOutput) ToApplicationGatewaySkuPtrOutput() ApplicationGatewaySkuPtrOutput {
	return o
}

func (o ApplicationGatewaySkuPtrOutput) ToApplicationGatewaySkuPtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuPtrOutput {
	return o
}

func (o ApplicationGatewaySkuPtrOutput) Elem() ApplicationGatewaySkuOutput {
	return o.ApplyT(func(v *ApplicationGatewaySku) ApplicationGatewaySku {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySku
		return ret
	}).(ApplicationGatewaySkuOutput)
}

func (o ApplicationGatewaySkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewaySkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ApplicationGatewaySkuResponseInput interface {
	pulumi.Input

	ToApplicationGatewaySkuResponseOutput() ApplicationGatewaySkuResponseOutput
	ToApplicationGatewaySkuResponseOutputWithContext(context.Context) ApplicationGatewaySkuResponseOutput
}

type ApplicationGatewaySkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ApplicationGatewaySkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySkuResponse)(nil)).Elem()
}

func (i ApplicationGatewaySkuResponseArgs) ToApplicationGatewaySkuResponseOutput() ApplicationGatewaySkuResponseOutput {
	return i.ToApplicationGatewaySkuResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewaySkuResponseArgs) ToApplicationGatewaySkuResponseOutputWithContext(ctx context.Context) ApplicationGatewaySkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySkuResponseOutput)
}

func (i ApplicationGatewaySkuResponseArgs) ToApplicationGatewaySkuResponsePtrOutput() ApplicationGatewaySkuResponsePtrOutput {
	return i.ToApplicationGatewaySkuResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationGatewaySkuResponseArgs) ToApplicationGatewaySkuResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySkuResponseOutput).ToApplicationGatewaySkuResponsePtrOutputWithContext(ctx)
}









type ApplicationGatewaySkuResponsePtrInput interface {
	pulumi.Input

	ToApplicationGatewaySkuResponsePtrOutput() ApplicationGatewaySkuResponsePtrOutput
	ToApplicationGatewaySkuResponsePtrOutputWithContext(context.Context) ApplicationGatewaySkuResponsePtrOutput
}

type applicationGatewaySkuResponsePtrType ApplicationGatewaySkuResponseArgs

func ApplicationGatewaySkuResponsePtr(v *ApplicationGatewaySkuResponseArgs) ApplicationGatewaySkuResponsePtrInput {
	return (*applicationGatewaySkuResponsePtrType)(v)
}

func (*applicationGatewaySkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySkuResponse)(nil)).Elem()
}

func (i *applicationGatewaySkuResponsePtrType) ToApplicationGatewaySkuResponsePtrOutput() ApplicationGatewaySkuResponsePtrOutput {
	return i.ToApplicationGatewaySkuResponsePtrOutputWithContext(context.Background())
}

func (i *applicationGatewaySkuResponsePtrType) ToApplicationGatewaySkuResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySkuResponsePtrOutput)
}

type ApplicationGatewaySkuResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySkuResponse)(nil)).Elem()
}

func (o ApplicationGatewaySkuResponseOutput) ToApplicationGatewaySkuResponseOutput() ApplicationGatewaySkuResponseOutput {
	return o
}

func (o ApplicationGatewaySkuResponseOutput) ToApplicationGatewaySkuResponseOutputWithContext(ctx context.Context) ApplicationGatewaySkuResponseOutput {
	return o
}

func (o ApplicationGatewaySkuResponseOutput) ToApplicationGatewaySkuResponsePtrOutput() ApplicationGatewaySkuResponsePtrOutput {
	return o.ToApplicationGatewaySkuResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySkuResponseOutput) ToApplicationGatewaySkuResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySkuResponse) *ApplicationGatewaySkuResponse {
		return &v
	}).(ApplicationGatewaySkuResponsePtrOutput)
}

func (o ApplicationGatewaySkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewaySkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySkuResponse)(nil)).Elem()
}

func (o ApplicationGatewaySkuResponsePtrOutput) ToApplicationGatewaySkuResponsePtrOutput() ApplicationGatewaySkuResponsePtrOutput {
	return o
}

func (o ApplicationGatewaySkuResponsePtrOutput) ToApplicationGatewaySkuResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuResponsePtrOutput {
	return o
}

func (o ApplicationGatewaySkuResponsePtrOutput) Elem() ApplicationGatewaySkuResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewaySkuResponse) ApplicationGatewaySkuResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySkuResponse
		return ret
	}).(ApplicationGatewaySkuResponseOutput)
}

func (o ApplicationGatewaySkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewaySkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslCertificate struct {
	Data              *string `pulumi:"data"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	Password          *string `pulumi:"password"`
	ProvisioningState *string `pulumi:"provisioningState"`
	PublicCertData    *string `pulumi:"publicCertData"`
}





type ApplicationGatewaySslCertificateInput interface {
	pulumi.Input

	ToApplicationGatewaySslCertificateOutput() ApplicationGatewaySslCertificateOutput
	ToApplicationGatewaySslCertificateOutputWithContext(context.Context) ApplicationGatewaySslCertificateOutput
}

type ApplicationGatewaySslCertificateArgs struct {
	Data              pulumi.StringPtrInput `pulumi:"data"`
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	Password          pulumi.StringPtrInput `pulumi:"password"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	PublicCertData    pulumi.StringPtrInput `pulumi:"publicCertData"`
}

func (ApplicationGatewaySslCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslCertificate)(nil)).Elem()
}

func (i ApplicationGatewaySslCertificateArgs) ToApplicationGatewaySslCertificateOutput() ApplicationGatewaySslCertificateOutput {
	return i.ToApplicationGatewaySslCertificateOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslCertificateArgs) ToApplicationGatewaySslCertificateOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslCertificateOutput)
}





type ApplicationGatewaySslCertificateArrayInput interface {
	pulumi.Input

	ToApplicationGatewaySslCertificateArrayOutput() ApplicationGatewaySslCertificateArrayOutput
	ToApplicationGatewaySslCertificateArrayOutputWithContext(context.Context) ApplicationGatewaySslCertificateArrayOutput
}

type ApplicationGatewaySslCertificateArray []ApplicationGatewaySslCertificateInput

func (ApplicationGatewaySslCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewaySslCertificate)(nil)).Elem()
}

func (i ApplicationGatewaySslCertificateArray) ToApplicationGatewaySslCertificateArrayOutput() ApplicationGatewaySslCertificateArrayOutput {
	return i.ToApplicationGatewaySslCertificateArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslCertificateArray) ToApplicationGatewaySslCertificateArrayOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslCertificateArrayOutput)
}

type ApplicationGatewaySslCertificateOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslCertificate)(nil)).Elem()
}

func (o ApplicationGatewaySslCertificateOutput) ToApplicationGatewaySslCertificateOutput() ApplicationGatewaySslCertificateOutput {
	return o
}

func (o ApplicationGatewaySslCertificateOutput) ToApplicationGatewaySslCertificateOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateOutput {
	return o
}

func (o ApplicationGatewaySslCertificateOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslCertificateArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewaySslCertificate)(nil)).Elem()
}

func (o ApplicationGatewaySslCertificateArrayOutput) ToApplicationGatewaySslCertificateArrayOutput() ApplicationGatewaySslCertificateArrayOutput {
	return o
}

func (o ApplicationGatewaySslCertificateArrayOutput) ToApplicationGatewaySslCertificateArrayOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateArrayOutput {
	return o
}

func (o ApplicationGatewaySslCertificateArrayOutput) Index(i pulumi.IntInput) ApplicationGatewaySslCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewaySslCertificate {
		return vs[0].([]ApplicationGatewaySslCertificate)[vs[1].(int)]
	}).(ApplicationGatewaySslCertificateOutput)
}

type ApplicationGatewaySslCertificateResponse struct {
	Data              *string `pulumi:"data"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	Password          *string `pulumi:"password"`
	ProvisioningState *string `pulumi:"provisioningState"`
	PublicCertData    *string `pulumi:"publicCertData"`
}





type ApplicationGatewaySslCertificateResponseInput interface {
	pulumi.Input

	ToApplicationGatewaySslCertificateResponseOutput() ApplicationGatewaySslCertificateResponseOutput
	ToApplicationGatewaySslCertificateResponseOutputWithContext(context.Context) ApplicationGatewaySslCertificateResponseOutput
}

type ApplicationGatewaySslCertificateResponseArgs struct {
	Data              pulumi.StringPtrInput `pulumi:"data"`
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	Password          pulumi.StringPtrInput `pulumi:"password"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	PublicCertData    pulumi.StringPtrInput `pulumi:"publicCertData"`
}

func (ApplicationGatewaySslCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslCertificateResponse)(nil)).Elem()
}

func (i ApplicationGatewaySslCertificateResponseArgs) ToApplicationGatewaySslCertificateResponseOutput() ApplicationGatewaySslCertificateResponseOutput {
	return i.ToApplicationGatewaySslCertificateResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslCertificateResponseArgs) ToApplicationGatewaySslCertificateResponseOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslCertificateResponseOutput)
}





type ApplicationGatewaySslCertificateResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewaySslCertificateResponseArrayOutput() ApplicationGatewaySslCertificateResponseArrayOutput
	ToApplicationGatewaySslCertificateResponseArrayOutputWithContext(context.Context) ApplicationGatewaySslCertificateResponseArrayOutput
}

type ApplicationGatewaySslCertificateResponseArray []ApplicationGatewaySslCertificateResponseInput

func (ApplicationGatewaySslCertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewaySslCertificateResponse)(nil)).Elem()
}

func (i ApplicationGatewaySslCertificateResponseArray) ToApplicationGatewaySslCertificateResponseArrayOutput() ApplicationGatewaySslCertificateResponseArrayOutput {
	return i.ToApplicationGatewaySslCertificateResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslCertificateResponseArray) ToApplicationGatewaySslCertificateResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslCertificateResponseArrayOutput)
}

type ApplicationGatewaySslCertificateResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslCertificateResponse)(nil)).Elem()
}

func (o ApplicationGatewaySslCertificateResponseOutput) ToApplicationGatewaySslCertificateResponseOutput() ApplicationGatewaySslCertificateResponseOutput {
	return o
}

func (o ApplicationGatewaySslCertificateResponseOutput) ToApplicationGatewaySslCertificateResponseOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateResponseOutput {
	return o
}

func (o ApplicationGatewaySslCertificateResponseOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewaySslCertificateResponse)(nil)).Elem()
}

func (o ApplicationGatewaySslCertificateResponseArrayOutput) ToApplicationGatewaySslCertificateResponseArrayOutput() ApplicationGatewaySslCertificateResponseArrayOutput {
	return o
}

func (o ApplicationGatewaySslCertificateResponseArrayOutput) ToApplicationGatewaySslCertificateResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateResponseArrayOutput {
	return o
}

func (o ApplicationGatewaySslCertificateResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewaySslCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewaySslCertificateResponse {
		return vs[0].([]ApplicationGatewaySslCertificateResponse)[vs[1].(int)]
	}).(ApplicationGatewaySslCertificateResponseOutput)
}

type BackendAddressPool struct {
	BackendIPConfigurations []SubResource `pulumi:"backendIPConfigurations"`
	Etag                    *string       `pulumi:"etag"`
	Id                      *string       `pulumi:"id"`
	LoadBalancingRules      []SubResource `pulumi:"loadBalancingRules"`
	Name                    *string       `pulumi:"name"`
	OutboundNatRule         *SubResource  `pulumi:"outboundNatRule"`
	ProvisioningState       *string       `pulumi:"provisioningState"`
}





type BackendAddressPoolInput interface {
	pulumi.Input

	ToBackendAddressPoolOutput() BackendAddressPoolOutput
	ToBackendAddressPoolOutputWithContext(context.Context) BackendAddressPoolOutput
}

type BackendAddressPoolArgs struct {
	BackendIPConfigurations SubResourceArrayInput `pulumi:"backendIPConfigurations"`
	Etag                    pulumi.StringPtrInput `pulumi:"etag"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	LoadBalancingRules      SubResourceArrayInput `pulumi:"loadBalancingRules"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	OutboundNatRule         SubResourcePtrInput   `pulumi:"outboundNatRule"`
	ProvisioningState       pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (BackendAddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAddressPool)(nil)).Elem()
}

func (i BackendAddressPoolArgs) ToBackendAddressPoolOutput() BackendAddressPoolOutput {
	return i.ToBackendAddressPoolOutputWithContext(context.Background())
}

func (i BackendAddressPoolArgs) ToBackendAddressPoolOutputWithContext(ctx context.Context) BackendAddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolOutput)
}





type BackendAddressPoolArrayInput interface {
	pulumi.Input

	ToBackendAddressPoolArrayOutput() BackendAddressPoolArrayOutput
	ToBackendAddressPoolArrayOutputWithContext(context.Context) BackendAddressPoolArrayOutput
}

type BackendAddressPoolArray []BackendAddressPoolInput

func (BackendAddressPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendAddressPool)(nil)).Elem()
}

func (i BackendAddressPoolArray) ToBackendAddressPoolArrayOutput() BackendAddressPoolArrayOutput {
	return i.ToBackendAddressPoolArrayOutputWithContext(context.Background())
}

func (i BackendAddressPoolArray) ToBackendAddressPoolArrayOutputWithContext(ctx context.Context) BackendAddressPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolArrayOutput)
}

type BackendAddressPoolOutput struct{ *pulumi.OutputState }

func (BackendAddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAddressPool)(nil)).Elem()
}

func (o BackendAddressPoolOutput) ToBackendAddressPoolOutput() BackendAddressPoolOutput {
	return o
}

func (o BackendAddressPoolOutput) ToBackendAddressPoolOutputWithContext(ctx context.Context) BackendAddressPoolOutput {
	return o
}

func (o BackendAddressPoolOutput) BackendIPConfigurations() SubResourceArrayOutput {
	return o.ApplyT(func(v BackendAddressPool) []SubResource { return v.BackendIPConfigurations }).(SubResourceArrayOutput)
}

func (o BackendAddressPoolOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolOutput) LoadBalancingRules() SubResourceArrayOutput {
	return o.ApplyT(func(v BackendAddressPool) []SubResource { return v.LoadBalancingRules }).(SubResourceArrayOutput)
}

func (o BackendAddressPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolOutput) OutboundNatRule() SubResourcePtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *SubResource { return v.OutboundNatRule }).(SubResourcePtrOutput)
}

func (o BackendAddressPoolOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type BackendAddressPoolArrayOutput struct{ *pulumi.OutputState }

func (BackendAddressPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendAddressPool)(nil)).Elem()
}

func (o BackendAddressPoolArrayOutput) ToBackendAddressPoolArrayOutput() BackendAddressPoolArrayOutput {
	return o
}

func (o BackendAddressPoolArrayOutput) ToBackendAddressPoolArrayOutputWithContext(ctx context.Context) BackendAddressPoolArrayOutput {
	return o
}

func (o BackendAddressPoolArrayOutput) Index(i pulumi.IntInput) BackendAddressPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendAddressPool {
		return vs[0].([]BackendAddressPool)[vs[1].(int)]
	}).(BackendAddressPoolOutput)
}

type BackendAddressPoolResponse struct {
	BackendIPConfigurations []SubResourceResponse `pulumi:"backendIPConfigurations"`
	Etag                    *string               `pulumi:"etag"`
	Id                      *string               `pulumi:"id"`
	LoadBalancingRules      []SubResourceResponse `pulumi:"loadBalancingRules"`
	Name                    *string               `pulumi:"name"`
	OutboundNatRule         *SubResourceResponse  `pulumi:"outboundNatRule"`
	ProvisioningState       *string               `pulumi:"provisioningState"`
}





type BackendAddressPoolResponseInput interface {
	pulumi.Input

	ToBackendAddressPoolResponseOutput() BackendAddressPoolResponseOutput
	ToBackendAddressPoolResponseOutputWithContext(context.Context) BackendAddressPoolResponseOutput
}

type BackendAddressPoolResponseArgs struct {
	BackendIPConfigurations SubResourceResponseArrayInput `pulumi:"backendIPConfigurations"`
	Etag                    pulumi.StringPtrInput         `pulumi:"etag"`
	Id                      pulumi.StringPtrInput         `pulumi:"id"`
	LoadBalancingRules      SubResourceResponseArrayInput `pulumi:"loadBalancingRules"`
	Name                    pulumi.StringPtrInput         `pulumi:"name"`
	OutboundNatRule         SubResourceResponsePtrInput   `pulumi:"outboundNatRule"`
	ProvisioningState       pulumi.StringPtrInput         `pulumi:"provisioningState"`
}

func (BackendAddressPoolResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAddressPoolResponse)(nil)).Elem()
}

func (i BackendAddressPoolResponseArgs) ToBackendAddressPoolResponseOutput() BackendAddressPoolResponseOutput {
	return i.ToBackendAddressPoolResponseOutputWithContext(context.Background())
}

func (i BackendAddressPoolResponseArgs) ToBackendAddressPoolResponseOutputWithContext(ctx context.Context) BackendAddressPoolResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolResponseOutput)
}





type BackendAddressPoolResponseArrayInput interface {
	pulumi.Input

	ToBackendAddressPoolResponseArrayOutput() BackendAddressPoolResponseArrayOutput
	ToBackendAddressPoolResponseArrayOutputWithContext(context.Context) BackendAddressPoolResponseArrayOutput
}

type BackendAddressPoolResponseArray []BackendAddressPoolResponseInput

func (BackendAddressPoolResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendAddressPoolResponse)(nil)).Elem()
}

func (i BackendAddressPoolResponseArray) ToBackendAddressPoolResponseArrayOutput() BackendAddressPoolResponseArrayOutput {
	return i.ToBackendAddressPoolResponseArrayOutputWithContext(context.Background())
}

func (i BackendAddressPoolResponseArray) ToBackendAddressPoolResponseArrayOutputWithContext(ctx context.Context) BackendAddressPoolResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolResponseArrayOutput)
}

type BackendAddressPoolResponseOutput struct{ *pulumi.OutputState }

func (BackendAddressPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAddressPoolResponse)(nil)).Elem()
}

func (o BackendAddressPoolResponseOutput) ToBackendAddressPoolResponseOutput() BackendAddressPoolResponseOutput {
	return o
}

func (o BackendAddressPoolResponseOutput) ToBackendAddressPoolResponseOutputWithContext(ctx context.Context) BackendAddressPoolResponseOutput {
	return o
}

func (o BackendAddressPoolResponseOutput) BackendIPConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) []SubResourceResponse { return v.BackendIPConfigurations }).(SubResourceResponseArrayOutput)
}

func (o BackendAddressPoolResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolResponseOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) []SubResourceResponse { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

func (o BackendAddressPoolResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolResponseOutput) OutboundNatRule() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) *SubResourceResponse { return v.OutboundNatRule }).(SubResourceResponsePtrOutput)
}

func (o BackendAddressPoolResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type BackendAddressPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (BackendAddressPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendAddressPoolResponse)(nil)).Elem()
}

func (o BackendAddressPoolResponseArrayOutput) ToBackendAddressPoolResponseArrayOutput() BackendAddressPoolResponseArrayOutput {
	return o
}

func (o BackendAddressPoolResponseArrayOutput) ToBackendAddressPoolResponseArrayOutputWithContext(ctx context.Context) BackendAddressPoolResponseArrayOutput {
	return o
}

func (o BackendAddressPoolResponseArrayOutput) Index(i pulumi.IntInput) BackendAddressPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendAddressPoolResponse {
		return vs[0].([]BackendAddressPoolResponse)[vs[1].(int)]
	}).(BackendAddressPoolResponseOutput)
}

type DhcpOptions struct {
	DnsServers []string `pulumi:"dnsServers"`
}





type DhcpOptionsInput interface {
	pulumi.Input

	ToDhcpOptionsOutput() DhcpOptionsOutput
	ToDhcpOptionsOutputWithContext(context.Context) DhcpOptionsOutput
}

type DhcpOptionsArgs struct {
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
}

func (DhcpOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DhcpOptions)(nil)).Elem()
}

func (i DhcpOptionsArgs) ToDhcpOptionsOutput() DhcpOptionsOutput {
	return i.ToDhcpOptionsOutputWithContext(context.Background())
}

func (i DhcpOptionsArgs) ToDhcpOptionsOutputWithContext(ctx context.Context) DhcpOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsOutput)
}

func (i DhcpOptionsArgs) ToDhcpOptionsPtrOutput() DhcpOptionsPtrOutput {
	return i.ToDhcpOptionsPtrOutputWithContext(context.Background())
}

func (i DhcpOptionsArgs) ToDhcpOptionsPtrOutputWithContext(ctx context.Context) DhcpOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsOutput).ToDhcpOptionsPtrOutputWithContext(ctx)
}









type DhcpOptionsPtrInput interface {
	pulumi.Input

	ToDhcpOptionsPtrOutput() DhcpOptionsPtrOutput
	ToDhcpOptionsPtrOutputWithContext(context.Context) DhcpOptionsPtrOutput
}

type dhcpOptionsPtrType DhcpOptionsArgs

func DhcpOptionsPtr(v *DhcpOptionsArgs) DhcpOptionsPtrInput {
	return (*dhcpOptionsPtrType)(v)
}

func (*dhcpOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptions)(nil)).Elem()
}

func (i *dhcpOptionsPtrType) ToDhcpOptionsPtrOutput() DhcpOptionsPtrOutput {
	return i.ToDhcpOptionsPtrOutputWithContext(context.Background())
}

func (i *dhcpOptionsPtrType) ToDhcpOptionsPtrOutputWithContext(ctx context.Context) DhcpOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsPtrOutput)
}

type DhcpOptionsOutput struct{ *pulumi.OutputState }

func (DhcpOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DhcpOptions)(nil)).Elem()
}

func (o DhcpOptionsOutput) ToDhcpOptionsOutput() DhcpOptionsOutput {
	return o
}

func (o DhcpOptionsOutput) ToDhcpOptionsOutputWithContext(ctx context.Context) DhcpOptionsOutput {
	return o
}

func (o DhcpOptionsOutput) ToDhcpOptionsPtrOutput() DhcpOptionsPtrOutput {
	return o.ToDhcpOptionsPtrOutputWithContext(context.Background())
}

func (o DhcpOptionsOutput) ToDhcpOptionsPtrOutputWithContext(ctx context.Context) DhcpOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DhcpOptions) *DhcpOptions {
		return &v
	}).(DhcpOptionsPtrOutput)
}

func (o DhcpOptionsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DhcpOptions) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type DhcpOptionsPtrOutput struct{ *pulumi.OutputState }

func (DhcpOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptions)(nil)).Elem()
}

func (o DhcpOptionsPtrOutput) ToDhcpOptionsPtrOutput() DhcpOptionsPtrOutput {
	return o
}

func (o DhcpOptionsPtrOutput) ToDhcpOptionsPtrOutputWithContext(ctx context.Context) DhcpOptionsPtrOutput {
	return o
}

func (o DhcpOptionsPtrOutput) Elem() DhcpOptionsOutput {
	return o.ApplyT(func(v *DhcpOptions) DhcpOptions {
		if v != nil {
			return *v
		}
		var ret DhcpOptions
		return ret
	}).(DhcpOptionsOutput)
}

func (o DhcpOptionsPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DhcpOptions) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type DhcpOptionsResponse struct {
	DnsServers []string `pulumi:"dnsServers"`
}





type DhcpOptionsResponseInput interface {
	pulumi.Input

	ToDhcpOptionsResponseOutput() DhcpOptionsResponseOutput
	ToDhcpOptionsResponseOutputWithContext(context.Context) DhcpOptionsResponseOutput
}

type DhcpOptionsResponseArgs struct {
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
}

func (DhcpOptionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DhcpOptionsResponse)(nil)).Elem()
}

func (i DhcpOptionsResponseArgs) ToDhcpOptionsResponseOutput() DhcpOptionsResponseOutput {
	return i.ToDhcpOptionsResponseOutputWithContext(context.Background())
}

func (i DhcpOptionsResponseArgs) ToDhcpOptionsResponseOutputWithContext(ctx context.Context) DhcpOptionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsResponseOutput)
}

func (i DhcpOptionsResponseArgs) ToDhcpOptionsResponsePtrOutput() DhcpOptionsResponsePtrOutput {
	return i.ToDhcpOptionsResponsePtrOutputWithContext(context.Background())
}

func (i DhcpOptionsResponseArgs) ToDhcpOptionsResponsePtrOutputWithContext(ctx context.Context) DhcpOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsResponseOutput).ToDhcpOptionsResponsePtrOutputWithContext(ctx)
}









type DhcpOptionsResponsePtrInput interface {
	pulumi.Input

	ToDhcpOptionsResponsePtrOutput() DhcpOptionsResponsePtrOutput
	ToDhcpOptionsResponsePtrOutputWithContext(context.Context) DhcpOptionsResponsePtrOutput
}

type dhcpOptionsResponsePtrType DhcpOptionsResponseArgs

func DhcpOptionsResponsePtr(v *DhcpOptionsResponseArgs) DhcpOptionsResponsePtrInput {
	return (*dhcpOptionsResponsePtrType)(v)
}

func (*dhcpOptionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptionsResponse)(nil)).Elem()
}

func (i *dhcpOptionsResponsePtrType) ToDhcpOptionsResponsePtrOutput() DhcpOptionsResponsePtrOutput {
	return i.ToDhcpOptionsResponsePtrOutputWithContext(context.Background())
}

func (i *dhcpOptionsResponsePtrType) ToDhcpOptionsResponsePtrOutputWithContext(ctx context.Context) DhcpOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsResponsePtrOutput)
}

type DhcpOptionsResponseOutput struct{ *pulumi.OutputState }

func (DhcpOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DhcpOptionsResponse)(nil)).Elem()
}

func (o DhcpOptionsResponseOutput) ToDhcpOptionsResponseOutput() DhcpOptionsResponseOutput {
	return o
}

func (o DhcpOptionsResponseOutput) ToDhcpOptionsResponseOutputWithContext(ctx context.Context) DhcpOptionsResponseOutput {
	return o
}

func (o DhcpOptionsResponseOutput) ToDhcpOptionsResponsePtrOutput() DhcpOptionsResponsePtrOutput {
	return o.ToDhcpOptionsResponsePtrOutputWithContext(context.Background())
}

func (o DhcpOptionsResponseOutput) ToDhcpOptionsResponsePtrOutputWithContext(ctx context.Context) DhcpOptionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DhcpOptionsResponse) *DhcpOptionsResponse {
		return &v
	}).(DhcpOptionsResponsePtrOutput)
}

func (o DhcpOptionsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DhcpOptionsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type DhcpOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (DhcpOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptionsResponse)(nil)).Elem()
}

func (o DhcpOptionsResponsePtrOutput) ToDhcpOptionsResponsePtrOutput() DhcpOptionsResponsePtrOutput {
	return o
}

func (o DhcpOptionsResponsePtrOutput) ToDhcpOptionsResponsePtrOutputWithContext(ctx context.Context) DhcpOptionsResponsePtrOutput {
	return o
}

func (o DhcpOptionsResponsePtrOutput) Elem() DhcpOptionsResponseOutput {
	return o.ApplyT(func(v *DhcpOptionsResponse) DhcpOptionsResponse {
		if v != nil {
			return *v
		}
		var ret DhcpOptionsResponse
		return ret
	}).(DhcpOptionsResponseOutput)
}

func (o DhcpOptionsResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DhcpOptionsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type ExpressRouteCircuitAuthorizationType struct {
	AuthorizationKey       *string `pulumi:"authorizationKey"`
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	Etag                   *string `pulumi:"etag"`
	Id                     *string `pulumi:"id"`
	Name                   *string `pulumi:"name"`
	ProvisioningState      *string `pulumi:"provisioningState"`
}





type ExpressRouteCircuitAuthorizationTypeInput interface {
	pulumi.Input

	ToExpressRouteCircuitAuthorizationTypeOutput() ExpressRouteCircuitAuthorizationTypeOutput
	ToExpressRouteCircuitAuthorizationTypeOutputWithContext(context.Context) ExpressRouteCircuitAuthorizationTypeOutput
}

type ExpressRouteCircuitAuthorizationTypeArgs struct {
	AuthorizationKey       pulumi.StringPtrInput `pulumi:"authorizationKey"`
	AuthorizationUseStatus pulumi.StringPtrInput `pulumi:"authorizationUseStatus"`
	Etag                   pulumi.StringPtrInput `pulumi:"etag"`
	Id                     pulumi.StringPtrInput `pulumi:"id"`
	Name                   pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState      pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (ExpressRouteCircuitAuthorizationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorizationType)(nil)).Elem()
}

func (i ExpressRouteCircuitAuthorizationTypeArgs) ToExpressRouteCircuitAuthorizationTypeOutput() ExpressRouteCircuitAuthorizationTypeOutput {
	return i.ToExpressRouteCircuitAuthorizationTypeOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitAuthorizationTypeArgs) ToExpressRouteCircuitAuthorizationTypeOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitAuthorizationTypeOutput)
}





type ExpressRouteCircuitAuthorizationTypeArrayInput interface {
	pulumi.Input

	ToExpressRouteCircuitAuthorizationTypeArrayOutput() ExpressRouteCircuitAuthorizationTypeArrayOutput
	ToExpressRouteCircuitAuthorizationTypeArrayOutputWithContext(context.Context) ExpressRouteCircuitAuthorizationTypeArrayOutput
}

type ExpressRouteCircuitAuthorizationTypeArray []ExpressRouteCircuitAuthorizationTypeInput

func (ExpressRouteCircuitAuthorizationTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitAuthorizationType)(nil)).Elem()
}

func (i ExpressRouteCircuitAuthorizationTypeArray) ToExpressRouteCircuitAuthorizationTypeArrayOutput() ExpressRouteCircuitAuthorizationTypeArrayOutput {
	return i.ToExpressRouteCircuitAuthorizationTypeArrayOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitAuthorizationTypeArray) ToExpressRouteCircuitAuthorizationTypeArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitAuthorizationTypeArrayOutput)
}

type ExpressRouteCircuitAuthorizationTypeOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitAuthorizationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorizationType)(nil)).Elem()
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) ToExpressRouteCircuitAuthorizationTypeOutput() ExpressRouteCircuitAuthorizationTypeOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) ToExpressRouteCircuitAuthorizationTypeOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationTypeOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) AuthorizationUseStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.AuthorizationUseStatus }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitAuthorizationTypeArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitAuthorizationTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitAuthorizationType)(nil)).Elem()
}

func (o ExpressRouteCircuitAuthorizationTypeArrayOutput) ToExpressRouteCircuitAuthorizationTypeArrayOutput() ExpressRouteCircuitAuthorizationTypeArrayOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationTypeArrayOutput) ToExpressRouteCircuitAuthorizationTypeArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationTypeArrayOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationTypeArrayOutput) Index(i pulumi.IntInput) ExpressRouteCircuitAuthorizationTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteCircuitAuthorizationType {
		return vs[0].([]ExpressRouteCircuitAuthorizationType)[vs[1].(int)]
	}).(ExpressRouteCircuitAuthorizationTypeOutput)
}

type ExpressRouteCircuitAuthorizationResponse struct {
	AuthorizationKey       *string `pulumi:"authorizationKey"`
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	Etag                   *string `pulumi:"etag"`
	Id                     *string `pulumi:"id"`
	Name                   *string `pulumi:"name"`
	ProvisioningState      *string `pulumi:"provisioningState"`
}





type ExpressRouteCircuitAuthorizationResponseInput interface {
	pulumi.Input

	ToExpressRouteCircuitAuthorizationResponseOutput() ExpressRouteCircuitAuthorizationResponseOutput
	ToExpressRouteCircuitAuthorizationResponseOutputWithContext(context.Context) ExpressRouteCircuitAuthorizationResponseOutput
}

type ExpressRouteCircuitAuthorizationResponseArgs struct {
	AuthorizationKey       pulumi.StringPtrInput `pulumi:"authorizationKey"`
	AuthorizationUseStatus pulumi.StringPtrInput `pulumi:"authorizationUseStatus"`
	Etag                   pulumi.StringPtrInput `pulumi:"etag"`
	Id                     pulumi.StringPtrInput `pulumi:"id"`
	Name                   pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState      pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (ExpressRouteCircuitAuthorizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorizationResponse)(nil)).Elem()
}

func (i ExpressRouteCircuitAuthorizationResponseArgs) ToExpressRouteCircuitAuthorizationResponseOutput() ExpressRouteCircuitAuthorizationResponseOutput {
	return i.ToExpressRouteCircuitAuthorizationResponseOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitAuthorizationResponseArgs) ToExpressRouteCircuitAuthorizationResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitAuthorizationResponseOutput)
}





type ExpressRouteCircuitAuthorizationResponseArrayInput interface {
	pulumi.Input

	ToExpressRouteCircuitAuthorizationResponseArrayOutput() ExpressRouteCircuitAuthorizationResponseArrayOutput
	ToExpressRouteCircuitAuthorizationResponseArrayOutputWithContext(context.Context) ExpressRouteCircuitAuthorizationResponseArrayOutput
}

type ExpressRouteCircuitAuthorizationResponseArray []ExpressRouteCircuitAuthorizationResponseInput

func (ExpressRouteCircuitAuthorizationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitAuthorizationResponse)(nil)).Elem()
}

func (i ExpressRouteCircuitAuthorizationResponseArray) ToExpressRouteCircuitAuthorizationResponseArrayOutput() ExpressRouteCircuitAuthorizationResponseArrayOutput {
	return i.ToExpressRouteCircuitAuthorizationResponseArrayOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitAuthorizationResponseArray) ToExpressRouteCircuitAuthorizationResponseArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitAuthorizationResponseArrayOutput)
}

type ExpressRouteCircuitAuthorizationResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitAuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorizationResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) ToExpressRouteCircuitAuthorizationResponseOutput() ExpressRouteCircuitAuthorizationResponseOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) ToExpressRouteCircuitAuthorizationResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationResponseOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) AuthorizationUseStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.AuthorizationUseStatus }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitAuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitAuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitAuthorizationResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitAuthorizationResponseArrayOutput) ToExpressRouteCircuitAuthorizationResponseArrayOutput() ExpressRouteCircuitAuthorizationResponseArrayOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationResponseArrayOutput) ToExpressRouteCircuitAuthorizationResponseArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationResponseArrayOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationResponseArrayOutput) Index(i pulumi.IntInput) ExpressRouteCircuitAuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteCircuitAuthorizationResponse {
		return vs[0].([]ExpressRouteCircuitAuthorizationResponse)[vs[1].(int)]
	}).(ExpressRouteCircuitAuthorizationResponseOutput)
}

type ExpressRouteCircuitPeeringType struct {
	AzureASN                   *int                              `pulumi:"azureASN"`
	Etag                       *string                           `pulumi:"etag"`
	Id                         *string                           `pulumi:"id"`
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfig `pulumi:"microsoftPeeringConfig"`
	Name                       *string                           `pulumi:"name"`
	PeerASN                    *int                              `pulumi:"peerASN"`
	PeeringType                *string                           `pulumi:"peeringType"`
	PrimaryAzurePort           *string                           `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   *string                           `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          *string                           `pulumi:"provisioningState"`
	SecondaryAzurePort         *string                           `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix *string                           `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  *string                           `pulumi:"sharedKey"`
	State                      *string                           `pulumi:"state"`
	Stats                      *ExpressRouteCircuitStats         `pulumi:"stats"`
	VlanId                     *int                              `pulumi:"vlanId"`
}





type ExpressRouteCircuitPeeringTypeInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringTypeOutput() ExpressRouteCircuitPeeringTypeOutput
	ToExpressRouteCircuitPeeringTypeOutputWithContext(context.Context) ExpressRouteCircuitPeeringTypeOutput
}

type ExpressRouteCircuitPeeringTypeArgs struct {
	AzureASN                   pulumi.IntPtrInput                       `pulumi:"azureASN"`
	Etag                       pulumi.StringPtrInput                    `pulumi:"etag"`
	Id                         pulumi.StringPtrInput                    `pulumi:"id"`
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigPtrInput `pulumi:"microsoftPeeringConfig"`
	Name                       pulumi.StringPtrInput                    `pulumi:"name"`
	PeerASN                    pulumi.IntPtrInput                       `pulumi:"peerASN"`
	PeeringType                pulumi.StringPtrInput                    `pulumi:"peeringType"`
	PrimaryAzurePort           pulumi.StringPtrInput                    `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   pulumi.StringPtrInput                    `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          pulumi.StringPtrInput                    `pulumi:"provisioningState"`
	SecondaryAzurePort         pulumi.StringPtrInput                    `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix pulumi.StringPtrInput                    `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  pulumi.StringPtrInput                    `pulumi:"sharedKey"`
	State                      pulumi.StringPtrInput                    `pulumi:"state"`
	Stats                      ExpressRouteCircuitStatsPtrInput         `pulumi:"stats"`
	VlanId                     pulumi.IntPtrInput                       `pulumi:"vlanId"`
}

func (ExpressRouteCircuitPeeringTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringType)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringTypeArgs) ToExpressRouteCircuitPeeringTypeOutput() ExpressRouteCircuitPeeringTypeOutput {
	return i.ToExpressRouteCircuitPeeringTypeOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringTypeArgs) ToExpressRouteCircuitPeeringTypeOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringTypeOutput)
}





type ExpressRouteCircuitPeeringTypeArrayInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringTypeArrayOutput() ExpressRouteCircuitPeeringTypeArrayOutput
	ToExpressRouteCircuitPeeringTypeArrayOutputWithContext(context.Context) ExpressRouteCircuitPeeringTypeArrayOutput
}

type ExpressRouteCircuitPeeringTypeArray []ExpressRouteCircuitPeeringTypeInput

func (ExpressRouteCircuitPeeringTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitPeeringType)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringTypeArray) ToExpressRouteCircuitPeeringTypeArrayOutput() ExpressRouteCircuitPeeringTypeArrayOutput {
	return i.ToExpressRouteCircuitPeeringTypeArrayOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringTypeArray) ToExpressRouteCircuitPeeringTypeArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringTypeArrayOutput)
}

type ExpressRouteCircuitPeeringTypeOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringType)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringTypeOutput) ToExpressRouteCircuitPeeringTypeOutput() ExpressRouteCircuitPeeringTypeOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeOutput) ToExpressRouteCircuitPeeringTypeOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeOutput) AzureASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *int { return v.AzureASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *ExpressRouteCircuitPeeringConfig {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) PeerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *int { return v.PeerASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.PeeringType }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) PrimaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.PrimaryAzurePort }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) SecondaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.SecondaryAzurePort }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Stats() ExpressRouteCircuitStatsPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *ExpressRouteCircuitStats { return v.Stats }).(ExpressRouteCircuitStatsPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

type ExpressRouteCircuitPeeringTypeArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitPeeringType)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringTypeArrayOutput) ToExpressRouteCircuitPeeringTypeArrayOutput() ExpressRouteCircuitPeeringTypeArrayOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeArrayOutput) ToExpressRouteCircuitPeeringTypeArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeArrayOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeArrayOutput) Index(i pulumi.IntInput) ExpressRouteCircuitPeeringTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteCircuitPeeringType {
		return vs[0].([]ExpressRouteCircuitPeeringType)[vs[1].(int)]
	}).(ExpressRouteCircuitPeeringTypeOutput)
}

type ExpressRouteCircuitPeeringConfig struct {
	AdvertisedPublicPrefixes      []string `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState *string  `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   *int     `pulumi:"customerASN"`
	RoutingRegistryName           *string  `pulumi:"routingRegistryName"`
}





type ExpressRouteCircuitPeeringConfigInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringConfigOutput() ExpressRouteCircuitPeeringConfigOutput
	ToExpressRouteCircuitPeeringConfigOutputWithContext(context.Context) ExpressRouteCircuitPeeringConfigOutput
}

type ExpressRouteCircuitPeeringConfigArgs struct {
	AdvertisedPublicPrefixes      pulumi.StringArrayInput `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState pulumi.StringPtrInput   `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   pulumi.IntPtrInput      `pulumi:"customerASN"`
	RoutingRegistryName           pulumi.StringPtrInput   `pulumi:"routingRegistryName"`
}

func (ExpressRouteCircuitPeeringConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringConfigArgs) ToExpressRouteCircuitPeeringConfigOutput() ExpressRouteCircuitPeeringConfigOutput {
	return i.ToExpressRouteCircuitPeeringConfigOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringConfigArgs) ToExpressRouteCircuitPeeringConfigOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringConfigOutput)
}

func (i ExpressRouteCircuitPeeringConfigArgs) ToExpressRouteCircuitPeeringConfigPtrOutput() ExpressRouteCircuitPeeringConfigPtrOutput {
	return i.ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringConfigArgs) ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringConfigOutput).ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx)
}









type ExpressRouteCircuitPeeringConfigPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringConfigPtrOutput() ExpressRouteCircuitPeeringConfigPtrOutput
	ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Context) ExpressRouteCircuitPeeringConfigPtrOutput
}

type expressRouteCircuitPeeringConfigPtrType ExpressRouteCircuitPeeringConfigArgs

func ExpressRouteCircuitPeeringConfigPtr(v *ExpressRouteCircuitPeeringConfigArgs) ExpressRouteCircuitPeeringConfigPtrInput {
	return (*expressRouteCircuitPeeringConfigPtrType)(v)
}

func (*expressRouteCircuitPeeringConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (i *expressRouteCircuitPeeringConfigPtrType) ToExpressRouteCircuitPeeringConfigPtrOutput() ExpressRouteCircuitPeeringConfigPtrOutput {
	return i.ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitPeeringConfigPtrType) ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

type ExpressRouteCircuitPeeringConfigOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringConfigOutput) ToExpressRouteCircuitPeeringConfigOutput() ExpressRouteCircuitPeeringConfigOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigOutput) ToExpressRouteCircuitPeeringConfigOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigOutput) ToExpressRouteCircuitPeeringConfigPtrOutput() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringConfigOutput) ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitPeeringConfig) *ExpressRouteCircuitPeeringConfig {
		return &v
	}).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigOutput) AdvertisedPublicPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfig) []string { return v.AdvertisedPublicPrefixes }).(pulumi.StringArrayOutput)
}

func (o ExpressRouteCircuitPeeringConfigOutput) AdvertisedPublicPrefixesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfig) *string { return v.AdvertisedPublicPrefixesState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigOutput) CustomerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfig) *int { return v.CustomerASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigOutput) RoutingRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfig) *string { return v.RoutingRegistryName }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringConfigPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) ToExpressRouteCircuitPeeringConfigPtrOutput() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigPtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) Elem() ExpressRouteCircuitPeeringConfigOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) ExpressRouteCircuitPeeringConfig {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitPeeringConfig
		return ret
	}).(ExpressRouteCircuitPeeringConfigOutput)
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) AdvertisedPublicPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) []string {
		if v == nil {
			return nil
		}
		return v.AdvertisedPublicPrefixes
	}).(pulumi.StringArrayOutput)
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) AdvertisedPublicPrefixesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) *string {
		if v == nil {
			return nil
		}
		return v.AdvertisedPublicPrefixesState
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) CustomerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) *int {
		if v == nil {
			return nil
		}
		return v.CustomerASN
	}).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) RoutingRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) *string {
		if v == nil {
			return nil
		}
		return v.RoutingRegistryName
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringConfigResponse struct {
	AdvertisedPublicPrefixes      []string `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState *string  `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   *int     `pulumi:"customerASN"`
	RoutingRegistryName           *string  `pulumi:"routingRegistryName"`
}





type ExpressRouteCircuitPeeringConfigResponseInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringConfigResponseOutput() ExpressRouteCircuitPeeringConfigResponseOutput
	ToExpressRouteCircuitPeeringConfigResponseOutputWithContext(context.Context) ExpressRouteCircuitPeeringConfigResponseOutput
}

type ExpressRouteCircuitPeeringConfigResponseArgs struct {
	AdvertisedPublicPrefixes      pulumi.StringArrayInput `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState pulumi.StringPtrInput   `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   pulumi.IntPtrInput      `pulumi:"customerASN"`
	RoutingRegistryName           pulumi.StringPtrInput   `pulumi:"routingRegistryName"`
}

func (ExpressRouteCircuitPeeringConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringConfigResponseArgs) ToExpressRouteCircuitPeeringConfigResponseOutput() ExpressRouteCircuitPeeringConfigResponseOutput {
	return i.ToExpressRouteCircuitPeeringConfigResponseOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringConfigResponseArgs) ToExpressRouteCircuitPeeringConfigResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringConfigResponseOutput)
}

func (i ExpressRouteCircuitPeeringConfigResponseArgs) ToExpressRouteCircuitPeeringConfigResponsePtrOutput() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return i.ToExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringConfigResponseArgs) ToExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringConfigResponseOutput).ToExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx)
}









type ExpressRouteCircuitPeeringConfigResponsePtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringConfigResponsePtrOutput() ExpressRouteCircuitPeeringConfigResponsePtrOutput
	ToExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(context.Context) ExpressRouteCircuitPeeringConfigResponsePtrOutput
}

type expressRouteCircuitPeeringConfigResponsePtrType ExpressRouteCircuitPeeringConfigResponseArgs

func ExpressRouteCircuitPeeringConfigResponsePtr(v *ExpressRouteCircuitPeeringConfigResponseArgs) ExpressRouteCircuitPeeringConfigResponsePtrInput {
	return (*expressRouteCircuitPeeringConfigResponsePtrType)(v)
}

func (*expressRouteCircuitPeeringConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (i *expressRouteCircuitPeeringConfigResponsePtrType) ToExpressRouteCircuitPeeringConfigResponsePtrOutput() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return i.ToExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitPeeringConfigResponsePtrType) ToExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

type ExpressRouteCircuitPeeringConfigResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) ToExpressRouteCircuitPeeringConfigResponseOutput() ExpressRouteCircuitPeeringConfigResponseOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) ToExpressRouteCircuitPeeringConfigResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigResponseOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) ToExpressRouteCircuitPeeringConfigResponsePtrOutput() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ToExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) ToExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitPeeringConfigResponse) *ExpressRouteCircuitPeeringConfigResponse {
		return &v
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) AdvertisedPublicPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfigResponse) []string { return v.AdvertisedPublicPrefixes }).(pulumi.StringArrayOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) AdvertisedPublicPrefixesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfigResponse) *string { return v.AdvertisedPublicPrefixesState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) CustomerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfigResponse) *int { return v.CustomerASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) RoutingRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfigResponse) *string { return v.RoutingRegistryName }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) ToExpressRouteCircuitPeeringConfigResponsePtrOutput() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) ToExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) Elem() ExpressRouteCircuitPeeringConfigResponseOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) ExpressRouteCircuitPeeringConfigResponse {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitPeeringConfigResponse
		return ret
	}).(ExpressRouteCircuitPeeringConfigResponseOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) AdvertisedPublicPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) []string {
		if v == nil {
			return nil
		}
		return v.AdvertisedPublicPrefixes
	}).(pulumi.StringArrayOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) AdvertisedPublicPrefixesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdvertisedPublicPrefixesState
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) CustomerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.CustomerASN
	}).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) RoutingRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RoutingRegistryName
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringResponse struct {
	AzureASN                   *int                                      `pulumi:"azureASN"`
	Etag                       *string                                   `pulumi:"etag"`
	Id                         *string                                   `pulumi:"id"`
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfigResponse `pulumi:"microsoftPeeringConfig"`
	Name                       *string                                   `pulumi:"name"`
	PeerASN                    *int                                      `pulumi:"peerASN"`
	PeeringType                *string                                   `pulumi:"peeringType"`
	PrimaryAzurePort           *string                                   `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   *string                                   `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          *string                                   `pulumi:"provisioningState"`
	SecondaryAzurePort         *string                                   `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix *string                                   `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  *string                                   `pulumi:"sharedKey"`
	State                      *string                                   `pulumi:"state"`
	Stats                      *ExpressRouteCircuitStatsResponse         `pulumi:"stats"`
	VlanId                     *int                                      `pulumi:"vlanId"`
}





type ExpressRouteCircuitPeeringResponseInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringResponseOutput() ExpressRouteCircuitPeeringResponseOutput
	ToExpressRouteCircuitPeeringResponseOutputWithContext(context.Context) ExpressRouteCircuitPeeringResponseOutput
}

type ExpressRouteCircuitPeeringResponseArgs struct {
	AzureASN                   pulumi.IntPtrInput                               `pulumi:"azureASN"`
	Etag                       pulumi.StringPtrInput                            `pulumi:"etag"`
	Id                         pulumi.StringPtrInput                            `pulumi:"id"`
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigResponsePtrInput `pulumi:"microsoftPeeringConfig"`
	Name                       pulumi.StringPtrInput                            `pulumi:"name"`
	PeerASN                    pulumi.IntPtrInput                               `pulumi:"peerASN"`
	PeeringType                pulumi.StringPtrInput                            `pulumi:"peeringType"`
	PrimaryAzurePort           pulumi.StringPtrInput                            `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   pulumi.StringPtrInput                            `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          pulumi.StringPtrInput                            `pulumi:"provisioningState"`
	SecondaryAzurePort         pulumi.StringPtrInput                            `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix pulumi.StringPtrInput                            `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  pulumi.StringPtrInput                            `pulumi:"sharedKey"`
	State                      pulumi.StringPtrInput                            `pulumi:"state"`
	Stats                      ExpressRouteCircuitStatsResponsePtrInput         `pulumi:"stats"`
	VlanId                     pulumi.IntPtrInput                               `pulumi:"vlanId"`
}

func (ExpressRouteCircuitPeeringResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringResponse)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringResponseArgs) ToExpressRouteCircuitPeeringResponseOutput() ExpressRouteCircuitPeeringResponseOutput {
	return i.ToExpressRouteCircuitPeeringResponseOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringResponseArgs) ToExpressRouteCircuitPeeringResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringResponseOutput)
}





type ExpressRouteCircuitPeeringResponseArrayInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringResponseArrayOutput() ExpressRouteCircuitPeeringResponseArrayOutput
	ToExpressRouteCircuitPeeringResponseArrayOutputWithContext(context.Context) ExpressRouteCircuitPeeringResponseArrayOutput
}

type ExpressRouteCircuitPeeringResponseArray []ExpressRouteCircuitPeeringResponseInput

func (ExpressRouteCircuitPeeringResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitPeeringResponse)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringResponseArray) ToExpressRouteCircuitPeeringResponseArrayOutput() ExpressRouteCircuitPeeringResponseArrayOutput {
	return i.ToExpressRouteCircuitPeeringResponseArrayOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringResponseArray) ToExpressRouteCircuitPeeringResponseArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

type ExpressRouteCircuitPeeringResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringResponseOutput) ToExpressRouteCircuitPeeringResponseOutput() ExpressRouteCircuitPeeringResponseOutput {
	return o
}

func (o ExpressRouteCircuitPeeringResponseOutput) ToExpressRouteCircuitPeeringResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringResponseOutput {
	return o
}

func (o ExpressRouteCircuitPeeringResponseOutput) AzureASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *int { return v.AzureASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *ExpressRouteCircuitPeeringConfigResponse {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) PeerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *int { return v.PeerASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.PeeringType }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) PrimaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.PrimaryAzurePort }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) SecondaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.SecondaryAzurePort }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Stats() ExpressRouteCircuitStatsResponsePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *ExpressRouteCircuitStatsResponse { return v.Stats }).(ExpressRouteCircuitStatsResponsePtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

type ExpressRouteCircuitPeeringResponseArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitPeeringResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringResponseArrayOutput) ToExpressRouteCircuitPeeringResponseArrayOutput() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o
}

func (o ExpressRouteCircuitPeeringResponseArrayOutput) ToExpressRouteCircuitPeeringResponseArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringResponseArrayOutput {
	return o
}

func (o ExpressRouteCircuitPeeringResponseArrayOutput) Index(i pulumi.IntInput) ExpressRouteCircuitPeeringResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteCircuitPeeringResponse {
		return vs[0].([]ExpressRouteCircuitPeeringResponse)[vs[1].(int)]
	}).(ExpressRouteCircuitPeeringResponseOutput)
}

type ExpressRouteCircuitServiceProviderProperties struct {
	BandwidthInMbps     *int    `pulumi:"bandwidthInMbps"`
	PeeringLocation     *string `pulumi:"peeringLocation"`
	ServiceProviderName *string `pulumi:"serviceProviderName"`
}





type ExpressRouteCircuitServiceProviderPropertiesInput interface {
	pulumi.Input

	ToExpressRouteCircuitServiceProviderPropertiesOutput() ExpressRouteCircuitServiceProviderPropertiesOutput
	ToExpressRouteCircuitServiceProviderPropertiesOutputWithContext(context.Context) ExpressRouteCircuitServiceProviderPropertiesOutput
}

type ExpressRouteCircuitServiceProviderPropertiesArgs struct {
	BandwidthInMbps     pulumi.IntPtrInput    `pulumi:"bandwidthInMbps"`
	PeeringLocation     pulumi.StringPtrInput `pulumi:"peeringLocation"`
	ServiceProviderName pulumi.StringPtrInput `pulumi:"serviceProviderName"`
}

func (ExpressRouteCircuitServiceProviderPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitServiceProviderProperties)(nil)).Elem()
}

func (i ExpressRouteCircuitServiceProviderPropertiesArgs) ToExpressRouteCircuitServiceProviderPropertiesOutput() ExpressRouteCircuitServiceProviderPropertiesOutput {
	return i.ToExpressRouteCircuitServiceProviderPropertiesOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitServiceProviderPropertiesArgs) ToExpressRouteCircuitServiceProviderPropertiesOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitServiceProviderPropertiesOutput)
}

func (i ExpressRouteCircuitServiceProviderPropertiesArgs) ToExpressRouteCircuitServiceProviderPropertiesPtrOutput() ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return i.ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitServiceProviderPropertiesArgs) ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitServiceProviderPropertiesOutput).ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(ctx)
}









type ExpressRouteCircuitServiceProviderPropertiesPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitServiceProviderPropertiesPtrOutput() ExpressRouteCircuitServiceProviderPropertiesPtrOutput
	ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(context.Context) ExpressRouteCircuitServiceProviderPropertiesPtrOutput
}

type expressRouteCircuitServiceProviderPropertiesPtrType ExpressRouteCircuitServiceProviderPropertiesArgs

func ExpressRouteCircuitServiceProviderPropertiesPtr(v *ExpressRouteCircuitServiceProviderPropertiesArgs) ExpressRouteCircuitServiceProviderPropertiesPtrInput {
	return (*expressRouteCircuitServiceProviderPropertiesPtrType)(v)
}

func (*expressRouteCircuitServiceProviderPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitServiceProviderProperties)(nil)).Elem()
}

func (i *expressRouteCircuitServiceProviderPropertiesPtrType) ToExpressRouteCircuitServiceProviderPropertiesPtrOutput() ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return i.ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitServiceProviderPropertiesPtrType) ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitServiceProviderPropertiesPtrOutput)
}

type ExpressRouteCircuitServiceProviderPropertiesOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitServiceProviderPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitServiceProviderProperties)(nil)).Elem()
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) ToExpressRouteCircuitServiceProviderPropertiesOutput() ExpressRouteCircuitServiceProviderPropertiesOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) ToExpressRouteCircuitServiceProviderPropertiesOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) ToExpressRouteCircuitServiceProviderPropertiesPtrOutput() ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return o.ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitServiceProviderProperties) *ExpressRouteCircuitServiceProviderProperties {
		return &v
	}).(ExpressRouteCircuitServiceProviderPropertiesPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) BandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderProperties) *int { return v.BandwidthInMbps }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderProperties) *string { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) ServiceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderProperties) *string { return v.ServiceProviderName }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitServiceProviderPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitServiceProviderPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitServiceProviderProperties)(nil)).Elem()
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) ToExpressRouteCircuitServiceProviderPropertiesPtrOutput() ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) Elem() ExpressRouteCircuitServiceProviderPropertiesOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderProperties) ExpressRouteCircuitServiceProviderProperties {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitServiceProviderProperties
		return ret
	}).(ExpressRouteCircuitServiceProviderPropertiesOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) BandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderProperties) *int {
		if v == nil {
			return nil
		}
		return v.BandwidthInMbps
	}).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderProperties) *string {
		if v == nil {
			return nil
		}
		return v.PeeringLocation
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) ServiceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServiceProviderName
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitServiceProviderPropertiesResponse struct {
	BandwidthInMbps     *int    `pulumi:"bandwidthInMbps"`
	PeeringLocation     *string `pulumi:"peeringLocation"`
	ServiceProviderName *string `pulumi:"serviceProviderName"`
}





type ExpressRouteCircuitServiceProviderPropertiesResponseInput interface {
	pulumi.Input

	ToExpressRouteCircuitServiceProviderPropertiesResponseOutput() ExpressRouteCircuitServiceProviderPropertiesResponseOutput
	ToExpressRouteCircuitServiceProviderPropertiesResponseOutputWithContext(context.Context) ExpressRouteCircuitServiceProviderPropertiesResponseOutput
}

type ExpressRouteCircuitServiceProviderPropertiesResponseArgs struct {
	BandwidthInMbps     pulumi.IntPtrInput    `pulumi:"bandwidthInMbps"`
	PeeringLocation     pulumi.StringPtrInput `pulumi:"peeringLocation"`
	ServiceProviderName pulumi.StringPtrInput `pulumi:"serviceProviderName"`
}

func (ExpressRouteCircuitServiceProviderPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitServiceProviderPropertiesResponse)(nil)).Elem()
}

func (i ExpressRouteCircuitServiceProviderPropertiesResponseArgs) ToExpressRouteCircuitServiceProviderPropertiesResponseOutput() ExpressRouteCircuitServiceProviderPropertiesResponseOutput {
	return i.ToExpressRouteCircuitServiceProviderPropertiesResponseOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitServiceProviderPropertiesResponseArgs) ToExpressRouteCircuitServiceProviderPropertiesResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitServiceProviderPropertiesResponseOutput)
}

func (i ExpressRouteCircuitServiceProviderPropertiesResponseArgs) ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput() ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return i.ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitServiceProviderPropertiesResponseArgs) ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitServiceProviderPropertiesResponseOutput).ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutputWithContext(ctx)
}









type ExpressRouteCircuitServiceProviderPropertiesResponsePtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput() ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput
	ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutputWithContext(context.Context) ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput
}

type expressRouteCircuitServiceProviderPropertiesResponsePtrType ExpressRouteCircuitServiceProviderPropertiesResponseArgs

func ExpressRouteCircuitServiceProviderPropertiesResponsePtr(v *ExpressRouteCircuitServiceProviderPropertiesResponseArgs) ExpressRouteCircuitServiceProviderPropertiesResponsePtrInput {
	return (*expressRouteCircuitServiceProviderPropertiesResponsePtrType)(v)
}

func (*expressRouteCircuitServiceProviderPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitServiceProviderPropertiesResponse)(nil)).Elem()
}

func (i *expressRouteCircuitServiceProviderPropertiesResponsePtrType) ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput() ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return i.ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitServiceProviderPropertiesResponsePtrType) ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput)
}

type ExpressRouteCircuitServiceProviderPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitServiceProviderPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitServiceProviderPropertiesResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) ToExpressRouteCircuitServiceProviderPropertiesResponseOutput() ExpressRouteCircuitServiceProviderPropertiesResponseOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) ToExpressRouteCircuitServiceProviderPropertiesResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesResponseOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput() ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return o.ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitServiceProviderPropertiesResponse) *ExpressRouteCircuitServiceProviderPropertiesResponse {
		return &v
	}).(ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) BandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderPropertiesResponse) *int { return v.BandwidthInMbps }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderPropertiesResponse) *string { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) ServiceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderPropertiesResponse) *string { return v.ServiceProviderName }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitServiceProviderPropertiesResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput() ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) Elem() ExpressRouteCircuitServiceProviderPropertiesResponseOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderPropertiesResponse) ExpressRouteCircuitServiceProviderPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitServiceProviderPropertiesResponse
		return ret
	}).(ExpressRouteCircuitServiceProviderPropertiesResponseOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) BandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.BandwidthInMbps
	}).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PeeringLocation
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) ServiceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceProviderName
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitSku struct {
	Family *string `pulumi:"family"`
	Name   *string `pulumi:"name"`
	Tier   *string `pulumi:"tier"`
}





type ExpressRouteCircuitSkuInput interface {
	pulumi.Input

	ToExpressRouteCircuitSkuOutput() ExpressRouteCircuitSkuOutput
	ToExpressRouteCircuitSkuOutputWithContext(context.Context) ExpressRouteCircuitSkuOutput
}

type ExpressRouteCircuitSkuArgs struct {
	Family pulumi.StringPtrInput `pulumi:"family"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Tier   pulumi.StringPtrInput `pulumi:"tier"`
}

func (ExpressRouteCircuitSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSku)(nil)).Elem()
}

func (i ExpressRouteCircuitSkuArgs) ToExpressRouteCircuitSkuOutput() ExpressRouteCircuitSkuOutput {
	return i.ToExpressRouteCircuitSkuOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitSkuArgs) ToExpressRouteCircuitSkuOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitSkuOutput)
}

func (i ExpressRouteCircuitSkuArgs) ToExpressRouteCircuitSkuPtrOutput() ExpressRouteCircuitSkuPtrOutput {
	return i.ToExpressRouteCircuitSkuPtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitSkuArgs) ToExpressRouteCircuitSkuPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitSkuOutput).ToExpressRouteCircuitSkuPtrOutputWithContext(ctx)
}









type ExpressRouteCircuitSkuPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitSkuPtrOutput() ExpressRouteCircuitSkuPtrOutput
	ToExpressRouteCircuitSkuPtrOutputWithContext(context.Context) ExpressRouteCircuitSkuPtrOutput
}

type expressRouteCircuitSkuPtrType ExpressRouteCircuitSkuArgs

func ExpressRouteCircuitSkuPtr(v *ExpressRouteCircuitSkuArgs) ExpressRouteCircuitSkuPtrInput {
	return (*expressRouteCircuitSkuPtrType)(v)
}

func (*expressRouteCircuitSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitSku)(nil)).Elem()
}

func (i *expressRouteCircuitSkuPtrType) ToExpressRouteCircuitSkuPtrOutput() ExpressRouteCircuitSkuPtrOutput {
	return i.ToExpressRouteCircuitSkuPtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitSkuPtrType) ToExpressRouteCircuitSkuPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitSkuPtrOutput)
}

type ExpressRouteCircuitSkuOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSku)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuOutput) ToExpressRouteCircuitSkuOutput() ExpressRouteCircuitSkuOutput {
	return o
}

func (o ExpressRouteCircuitSkuOutput) ToExpressRouteCircuitSkuOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuOutput {
	return o
}

func (o ExpressRouteCircuitSkuOutput) ToExpressRouteCircuitSkuPtrOutput() ExpressRouteCircuitSkuPtrOutput {
	return o.ToExpressRouteCircuitSkuPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuOutput) ToExpressRouteCircuitSkuPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitSku) *ExpressRouteCircuitSku {
		return &v
	}).(ExpressRouteCircuitSkuPtrOutput)
}

func (o ExpressRouteCircuitSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitSkuPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitSku)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuPtrOutput) ToExpressRouteCircuitSkuPtrOutput() ExpressRouteCircuitSkuPtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuPtrOutput) ToExpressRouteCircuitSkuPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuPtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuPtrOutput) Elem() ExpressRouteCircuitSkuOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSku) ExpressRouteCircuitSku {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitSku
		return ret
	}).(ExpressRouteCircuitSkuOutput)
}

func (o ExpressRouteCircuitSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitSkuResponse struct {
	Family *string `pulumi:"family"`
	Name   *string `pulumi:"name"`
	Tier   *string `pulumi:"tier"`
}





type ExpressRouteCircuitSkuResponseInput interface {
	pulumi.Input

	ToExpressRouteCircuitSkuResponseOutput() ExpressRouteCircuitSkuResponseOutput
	ToExpressRouteCircuitSkuResponseOutputWithContext(context.Context) ExpressRouteCircuitSkuResponseOutput
}

type ExpressRouteCircuitSkuResponseArgs struct {
	Family pulumi.StringPtrInput `pulumi:"family"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Tier   pulumi.StringPtrInput `pulumi:"tier"`
}

func (ExpressRouteCircuitSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSkuResponse)(nil)).Elem()
}

func (i ExpressRouteCircuitSkuResponseArgs) ToExpressRouteCircuitSkuResponseOutput() ExpressRouteCircuitSkuResponseOutput {
	return i.ToExpressRouteCircuitSkuResponseOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitSkuResponseArgs) ToExpressRouteCircuitSkuResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitSkuResponseOutput)
}

func (i ExpressRouteCircuitSkuResponseArgs) ToExpressRouteCircuitSkuResponsePtrOutput() ExpressRouteCircuitSkuResponsePtrOutput {
	return i.ToExpressRouteCircuitSkuResponsePtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitSkuResponseArgs) ToExpressRouteCircuitSkuResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitSkuResponseOutput).ToExpressRouteCircuitSkuResponsePtrOutputWithContext(ctx)
}









type ExpressRouteCircuitSkuResponsePtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitSkuResponsePtrOutput() ExpressRouteCircuitSkuResponsePtrOutput
	ToExpressRouteCircuitSkuResponsePtrOutputWithContext(context.Context) ExpressRouteCircuitSkuResponsePtrOutput
}

type expressRouteCircuitSkuResponsePtrType ExpressRouteCircuitSkuResponseArgs

func ExpressRouteCircuitSkuResponsePtr(v *ExpressRouteCircuitSkuResponseArgs) ExpressRouteCircuitSkuResponsePtrInput {
	return (*expressRouteCircuitSkuResponsePtrType)(v)
}

func (*expressRouteCircuitSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitSkuResponse)(nil)).Elem()
}

func (i *expressRouteCircuitSkuResponsePtrType) ToExpressRouteCircuitSkuResponsePtrOutput() ExpressRouteCircuitSkuResponsePtrOutput {
	return i.ToExpressRouteCircuitSkuResponsePtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitSkuResponsePtrType) ToExpressRouteCircuitSkuResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitSkuResponsePtrOutput)
}

type ExpressRouteCircuitSkuResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSkuResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuResponseOutput) ToExpressRouteCircuitSkuResponseOutput() ExpressRouteCircuitSkuResponseOutput {
	return o
}

func (o ExpressRouteCircuitSkuResponseOutput) ToExpressRouteCircuitSkuResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuResponseOutput {
	return o
}

func (o ExpressRouteCircuitSkuResponseOutput) ToExpressRouteCircuitSkuResponsePtrOutput() ExpressRouteCircuitSkuResponsePtrOutput {
	return o.ToExpressRouteCircuitSkuResponsePtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuResponseOutput) ToExpressRouteCircuitSkuResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitSkuResponse) *ExpressRouteCircuitSkuResponse {
		return &v
	}).(ExpressRouteCircuitSkuResponsePtrOutput)
}

func (o ExpressRouteCircuitSkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitSkuResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) ToExpressRouteCircuitSkuResponsePtrOutput() ExpressRouteCircuitSkuResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) ToExpressRouteCircuitSkuResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) Elem() ExpressRouteCircuitSkuResponseOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSkuResponse) ExpressRouteCircuitSkuResponse {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitSkuResponse
		return ret
	}).(ExpressRouteCircuitSkuResponseOutput)
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitStats struct {
	BytesIn  *int `pulumi:"bytesIn"`
	BytesOut *int `pulumi:"bytesOut"`
}





type ExpressRouteCircuitStatsInput interface {
	pulumi.Input

	ToExpressRouteCircuitStatsOutput() ExpressRouteCircuitStatsOutput
	ToExpressRouteCircuitStatsOutputWithContext(context.Context) ExpressRouteCircuitStatsOutput
}

type ExpressRouteCircuitStatsArgs struct {
	BytesIn  pulumi.IntPtrInput `pulumi:"bytesIn"`
	BytesOut pulumi.IntPtrInput `pulumi:"bytesOut"`
}

func (ExpressRouteCircuitStatsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitStats)(nil)).Elem()
}

func (i ExpressRouteCircuitStatsArgs) ToExpressRouteCircuitStatsOutput() ExpressRouteCircuitStatsOutput {
	return i.ToExpressRouteCircuitStatsOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitStatsArgs) ToExpressRouteCircuitStatsOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitStatsOutput)
}

func (i ExpressRouteCircuitStatsArgs) ToExpressRouteCircuitStatsPtrOutput() ExpressRouteCircuitStatsPtrOutput {
	return i.ToExpressRouteCircuitStatsPtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitStatsArgs) ToExpressRouteCircuitStatsPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitStatsOutput).ToExpressRouteCircuitStatsPtrOutputWithContext(ctx)
}









type ExpressRouteCircuitStatsPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitStatsPtrOutput() ExpressRouteCircuitStatsPtrOutput
	ToExpressRouteCircuitStatsPtrOutputWithContext(context.Context) ExpressRouteCircuitStatsPtrOutput
}

type expressRouteCircuitStatsPtrType ExpressRouteCircuitStatsArgs

func ExpressRouteCircuitStatsPtr(v *ExpressRouteCircuitStatsArgs) ExpressRouteCircuitStatsPtrInput {
	return (*expressRouteCircuitStatsPtrType)(v)
}

func (*expressRouteCircuitStatsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitStats)(nil)).Elem()
}

func (i *expressRouteCircuitStatsPtrType) ToExpressRouteCircuitStatsPtrOutput() ExpressRouteCircuitStatsPtrOutput {
	return i.ToExpressRouteCircuitStatsPtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitStatsPtrType) ToExpressRouteCircuitStatsPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitStatsPtrOutput)
}

type ExpressRouteCircuitStatsOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitStatsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitStats)(nil)).Elem()
}

func (o ExpressRouteCircuitStatsOutput) ToExpressRouteCircuitStatsOutput() ExpressRouteCircuitStatsOutput {
	return o
}

func (o ExpressRouteCircuitStatsOutput) ToExpressRouteCircuitStatsOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsOutput {
	return o
}

func (o ExpressRouteCircuitStatsOutput) ToExpressRouteCircuitStatsPtrOutput() ExpressRouteCircuitStatsPtrOutput {
	return o.ToExpressRouteCircuitStatsPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitStatsOutput) ToExpressRouteCircuitStatsPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitStats) *ExpressRouteCircuitStats {
		return &v
	}).(ExpressRouteCircuitStatsPtrOutput)
}

func (o ExpressRouteCircuitStatsOutput) BytesIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStats) *int { return v.BytesIn }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitStatsOutput) BytesOut() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStats) *int { return v.BytesOut }).(pulumi.IntPtrOutput)
}

type ExpressRouteCircuitStatsPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitStatsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitStats)(nil)).Elem()
}

func (o ExpressRouteCircuitStatsPtrOutput) ToExpressRouteCircuitStatsPtrOutput() ExpressRouteCircuitStatsPtrOutput {
	return o
}

func (o ExpressRouteCircuitStatsPtrOutput) ToExpressRouteCircuitStatsPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsPtrOutput {
	return o
}

func (o ExpressRouteCircuitStatsPtrOutput) Elem() ExpressRouteCircuitStatsOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) ExpressRouteCircuitStats {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitStats
		return ret
	}).(ExpressRouteCircuitStatsOutput)
}

func (o ExpressRouteCircuitStatsPtrOutput) BytesIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) *int {
		if v == nil {
			return nil
		}
		return v.BytesIn
	}).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitStatsPtrOutput) BytesOut() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) *int {
		if v == nil {
			return nil
		}
		return v.BytesOut
	}).(pulumi.IntPtrOutput)
}

type ExpressRouteCircuitStatsResponse struct {
	BytesIn  *int `pulumi:"bytesIn"`
	BytesOut *int `pulumi:"bytesOut"`
}





type ExpressRouteCircuitStatsResponseInput interface {
	pulumi.Input

	ToExpressRouteCircuitStatsResponseOutput() ExpressRouteCircuitStatsResponseOutput
	ToExpressRouteCircuitStatsResponseOutputWithContext(context.Context) ExpressRouteCircuitStatsResponseOutput
}

type ExpressRouteCircuitStatsResponseArgs struct {
	BytesIn  pulumi.IntPtrInput `pulumi:"bytesIn"`
	BytesOut pulumi.IntPtrInput `pulumi:"bytesOut"`
}

func (ExpressRouteCircuitStatsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitStatsResponse)(nil)).Elem()
}

func (i ExpressRouteCircuitStatsResponseArgs) ToExpressRouteCircuitStatsResponseOutput() ExpressRouteCircuitStatsResponseOutput {
	return i.ToExpressRouteCircuitStatsResponseOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitStatsResponseArgs) ToExpressRouteCircuitStatsResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitStatsResponseOutput)
}

func (i ExpressRouteCircuitStatsResponseArgs) ToExpressRouteCircuitStatsResponsePtrOutput() ExpressRouteCircuitStatsResponsePtrOutput {
	return i.ToExpressRouteCircuitStatsResponsePtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitStatsResponseArgs) ToExpressRouteCircuitStatsResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitStatsResponseOutput).ToExpressRouteCircuitStatsResponsePtrOutputWithContext(ctx)
}









type ExpressRouteCircuitStatsResponsePtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitStatsResponsePtrOutput() ExpressRouteCircuitStatsResponsePtrOutput
	ToExpressRouteCircuitStatsResponsePtrOutputWithContext(context.Context) ExpressRouteCircuitStatsResponsePtrOutput
}

type expressRouteCircuitStatsResponsePtrType ExpressRouteCircuitStatsResponseArgs

func ExpressRouteCircuitStatsResponsePtr(v *ExpressRouteCircuitStatsResponseArgs) ExpressRouteCircuitStatsResponsePtrInput {
	return (*expressRouteCircuitStatsResponsePtrType)(v)
}

func (*expressRouteCircuitStatsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitStatsResponse)(nil)).Elem()
}

func (i *expressRouteCircuitStatsResponsePtrType) ToExpressRouteCircuitStatsResponsePtrOutput() ExpressRouteCircuitStatsResponsePtrOutput {
	return i.ToExpressRouteCircuitStatsResponsePtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitStatsResponsePtrType) ToExpressRouteCircuitStatsResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitStatsResponsePtrOutput)
}

type ExpressRouteCircuitStatsResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitStatsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitStatsResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitStatsResponseOutput) ToExpressRouteCircuitStatsResponseOutput() ExpressRouteCircuitStatsResponseOutput {
	return o
}

func (o ExpressRouteCircuitStatsResponseOutput) ToExpressRouteCircuitStatsResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsResponseOutput {
	return o
}

func (o ExpressRouteCircuitStatsResponseOutput) ToExpressRouteCircuitStatsResponsePtrOutput() ExpressRouteCircuitStatsResponsePtrOutput {
	return o.ToExpressRouteCircuitStatsResponsePtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitStatsResponseOutput) ToExpressRouteCircuitStatsResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitStatsResponse) *ExpressRouteCircuitStatsResponse {
		return &v
	}).(ExpressRouteCircuitStatsResponsePtrOutput)
}

func (o ExpressRouteCircuitStatsResponseOutput) BytesIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStatsResponse) *int { return v.BytesIn }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitStatsResponseOutput) BytesOut() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStatsResponse) *int { return v.BytesOut }).(pulumi.IntPtrOutput)
}

type ExpressRouteCircuitStatsResponsePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitStatsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitStatsResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) ToExpressRouteCircuitStatsResponsePtrOutput() ExpressRouteCircuitStatsResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) ToExpressRouteCircuitStatsResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) Elem() ExpressRouteCircuitStatsResponseOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) ExpressRouteCircuitStatsResponse {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitStatsResponse
		return ret
	}).(ExpressRouteCircuitStatsResponseOutput)
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) BytesIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) *int {
		if v == nil {
			return nil
		}
		return v.BytesIn
	}).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) BytesOut() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) *int {
		if v == nil {
			return nil
		}
		return v.BytesOut
	}).(pulumi.IntPtrOutput)
}

type FrontendIpConfiguration struct {
	Etag                      *string       `pulumi:"etag"`
	Id                        *string       `pulumi:"id"`
	InboundNatPools           []SubResource `pulumi:"inboundNatPools"`
	InboundNatRules           []SubResource `pulumi:"inboundNatRules"`
	LoadBalancingRules        []SubResource `pulumi:"loadBalancingRules"`
	Name                      *string       `pulumi:"name"`
	OutboundNatRules          []SubResource `pulumi:"outboundNatRules"`
	PrivateIPAddress          *string       `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string       `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string       `pulumi:"provisioningState"`
	PublicIPAddress           *SubResource  `pulumi:"publicIPAddress"`
	Subnet                    *SubResource  `pulumi:"subnet"`
}





type FrontendIpConfigurationInput interface {
	pulumi.Input

	ToFrontendIpConfigurationOutput() FrontendIpConfigurationOutput
	ToFrontendIpConfigurationOutputWithContext(context.Context) FrontendIpConfigurationOutput
}

type FrontendIpConfigurationArgs struct {
	Etag                      pulumi.StringPtrInput `pulumi:"etag"`
	Id                        pulumi.StringPtrInput `pulumi:"id"`
	InboundNatPools           SubResourceArrayInput `pulumi:"inboundNatPools"`
	InboundNatRules           SubResourceArrayInput `pulumi:"inboundNatRules"`
	LoadBalancingRules        SubResourceArrayInput `pulumi:"loadBalancingRules"`
	Name                      pulumi.StringPtrInput `pulumi:"name"`
	OutboundNatRules          SubResourceArrayInput `pulumi:"outboundNatRules"`
	PrivateIPAddress          pulumi.StringPtrInput `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringPtrInput `pulumi:"provisioningState"`
	PublicIPAddress           SubResourcePtrInput   `pulumi:"publicIPAddress"`
	Subnet                    SubResourcePtrInput   `pulumi:"subnet"`
}

func (FrontendIpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIpConfiguration)(nil)).Elem()
}

func (i FrontendIpConfigurationArgs) ToFrontendIpConfigurationOutput() FrontendIpConfigurationOutput {
	return i.ToFrontendIpConfigurationOutputWithContext(context.Background())
}

func (i FrontendIpConfigurationArgs) ToFrontendIpConfigurationOutputWithContext(ctx context.Context) FrontendIpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendIpConfigurationOutput)
}





type FrontendIpConfigurationArrayInput interface {
	pulumi.Input

	ToFrontendIpConfigurationArrayOutput() FrontendIpConfigurationArrayOutput
	ToFrontendIpConfigurationArrayOutputWithContext(context.Context) FrontendIpConfigurationArrayOutput
}

type FrontendIpConfigurationArray []FrontendIpConfigurationInput

func (FrontendIpConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIpConfiguration)(nil)).Elem()
}

func (i FrontendIpConfigurationArray) ToFrontendIpConfigurationArrayOutput() FrontendIpConfigurationArrayOutput {
	return i.ToFrontendIpConfigurationArrayOutputWithContext(context.Background())
}

func (i FrontendIpConfigurationArray) ToFrontendIpConfigurationArrayOutputWithContext(ctx context.Context) FrontendIpConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendIpConfigurationArrayOutput)
}

type FrontendIpConfigurationOutput struct{ *pulumi.OutputState }

func (FrontendIpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIpConfiguration)(nil)).Elem()
}

func (o FrontendIpConfigurationOutput) ToFrontendIpConfigurationOutput() FrontendIpConfigurationOutput {
	return o
}

func (o FrontendIpConfigurationOutput) ToFrontendIpConfigurationOutputWithContext(ctx context.Context) FrontendIpConfigurationOutput {
	return o
}

func (o FrontendIpConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationOutput) InboundNatPools() SubResourceArrayOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) []SubResource { return v.InboundNatPools }).(SubResourceArrayOutput)
}

func (o FrontendIpConfigurationOutput) InboundNatRules() SubResourceArrayOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) []SubResource { return v.InboundNatRules }).(SubResourceArrayOutput)
}

func (o FrontendIpConfigurationOutput) LoadBalancingRules() SubResourceArrayOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) []SubResource { return v.LoadBalancingRules }).(SubResourceArrayOutput)
}

func (o FrontendIpConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationOutput) OutboundNatRules() SubResourceArrayOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) []SubResource { return v.OutboundNatRules }).(SubResourceArrayOutput)
}

func (o FrontendIpConfigurationOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationOutput) PublicIPAddress() SubResourcePtrOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) *SubResource { return v.PublicIPAddress }).(SubResourcePtrOutput)
}

func (o FrontendIpConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v FrontendIpConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type FrontendIpConfigurationArrayOutput struct{ *pulumi.OutputState }

func (FrontendIpConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIpConfiguration)(nil)).Elem()
}

func (o FrontendIpConfigurationArrayOutput) ToFrontendIpConfigurationArrayOutput() FrontendIpConfigurationArrayOutput {
	return o
}

func (o FrontendIpConfigurationArrayOutput) ToFrontendIpConfigurationArrayOutputWithContext(ctx context.Context) FrontendIpConfigurationArrayOutput {
	return o
}

func (o FrontendIpConfigurationArrayOutput) Index(i pulumi.IntInput) FrontendIpConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendIpConfiguration {
		return vs[0].([]FrontendIpConfiguration)[vs[1].(int)]
	}).(FrontendIpConfigurationOutput)
}

type FrontendIpConfigurationResponse struct {
	Etag                      *string               `pulumi:"etag"`
	Id                        *string               `pulumi:"id"`
	InboundNatPools           []SubResourceResponse `pulumi:"inboundNatPools"`
	InboundNatRules           []SubResourceResponse `pulumi:"inboundNatRules"`
	LoadBalancingRules        []SubResourceResponse `pulumi:"loadBalancingRules"`
	Name                      *string               `pulumi:"name"`
	OutboundNatRules          []SubResourceResponse `pulumi:"outboundNatRules"`
	PrivateIPAddress          *string               `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string               `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string               `pulumi:"provisioningState"`
	PublicIPAddress           *SubResourceResponse  `pulumi:"publicIPAddress"`
	Subnet                    *SubResourceResponse  `pulumi:"subnet"`
}





type FrontendIpConfigurationResponseInput interface {
	pulumi.Input

	ToFrontendIpConfigurationResponseOutput() FrontendIpConfigurationResponseOutput
	ToFrontendIpConfigurationResponseOutputWithContext(context.Context) FrontendIpConfigurationResponseOutput
}

type FrontendIpConfigurationResponseArgs struct {
	Etag                      pulumi.StringPtrInput         `pulumi:"etag"`
	Id                        pulumi.StringPtrInput         `pulumi:"id"`
	InboundNatPools           SubResourceResponseArrayInput `pulumi:"inboundNatPools"`
	InboundNatRules           SubResourceResponseArrayInput `pulumi:"inboundNatRules"`
	LoadBalancingRules        SubResourceResponseArrayInput `pulumi:"loadBalancingRules"`
	Name                      pulumi.StringPtrInput         `pulumi:"name"`
	OutboundNatRules          SubResourceResponseArrayInput `pulumi:"outboundNatRules"`
	PrivateIPAddress          pulumi.StringPtrInput         `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput         `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringPtrInput         `pulumi:"provisioningState"`
	PublicIPAddress           SubResourceResponsePtrInput   `pulumi:"publicIPAddress"`
	Subnet                    SubResourceResponsePtrInput   `pulumi:"subnet"`
}

func (FrontendIpConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIpConfigurationResponse)(nil)).Elem()
}

func (i FrontendIpConfigurationResponseArgs) ToFrontendIpConfigurationResponseOutput() FrontendIpConfigurationResponseOutput {
	return i.ToFrontendIpConfigurationResponseOutputWithContext(context.Background())
}

func (i FrontendIpConfigurationResponseArgs) ToFrontendIpConfigurationResponseOutputWithContext(ctx context.Context) FrontendIpConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendIpConfigurationResponseOutput)
}





type FrontendIpConfigurationResponseArrayInput interface {
	pulumi.Input

	ToFrontendIpConfigurationResponseArrayOutput() FrontendIpConfigurationResponseArrayOutput
	ToFrontendIpConfigurationResponseArrayOutputWithContext(context.Context) FrontendIpConfigurationResponseArrayOutput
}

type FrontendIpConfigurationResponseArray []FrontendIpConfigurationResponseInput

func (FrontendIpConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIpConfigurationResponse)(nil)).Elem()
}

func (i FrontendIpConfigurationResponseArray) ToFrontendIpConfigurationResponseArrayOutput() FrontendIpConfigurationResponseArrayOutput {
	return i.ToFrontendIpConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i FrontendIpConfigurationResponseArray) ToFrontendIpConfigurationResponseArrayOutputWithContext(ctx context.Context) FrontendIpConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendIpConfigurationResponseArrayOutput)
}

type FrontendIpConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FrontendIpConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIpConfigurationResponse)(nil)).Elem()
}

func (o FrontendIpConfigurationResponseOutput) ToFrontendIpConfigurationResponseOutput() FrontendIpConfigurationResponseOutput {
	return o
}

func (o FrontendIpConfigurationResponseOutput) ToFrontendIpConfigurationResponseOutputWithContext(ctx context.Context) FrontendIpConfigurationResponseOutput {
	return o
}

func (o FrontendIpConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationResponseOutput) InboundNatPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) []SubResourceResponse { return v.InboundNatPools }).(SubResourceResponseArrayOutput)
}

func (o FrontendIpConfigurationResponseOutput) InboundNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) []SubResourceResponse { return v.InboundNatRules }).(SubResourceResponseArrayOutput)
}

func (o FrontendIpConfigurationResponseOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) []SubResourceResponse { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

func (o FrontendIpConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationResponseOutput) OutboundNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) []SubResourceResponse { return v.OutboundNatRules }).(SubResourceResponseArrayOutput)
}

func (o FrontendIpConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o FrontendIpConfigurationResponseOutput) PublicIPAddress() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) *SubResourceResponse { return v.PublicIPAddress }).(SubResourceResponsePtrOutput)
}

func (o FrontendIpConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v FrontendIpConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type FrontendIpConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontendIpConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIpConfigurationResponse)(nil)).Elem()
}

func (o FrontendIpConfigurationResponseArrayOutput) ToFrontendIpConfigurationResponseArrayOutput() FrontendIpConfigurationResponseArrayOutput {
	return o
}

func (o FrontendIpConfigurationResponseArrayOutput) ToFrontendIpConfigurationResponseArrayOutputWithContext(ctx context.Context) FrontendIpConfigurationResponseArrayOutput {
	return o
}

func (o FrontendIpConfigurationResponseArrayOutput) Index(i pulumi.IntInput) FrontendIpConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendIpConfigurationResponse {
		return vs[0].([]FrontendIpConfigurationResponse)[vs[1].(int)]
	}).(FrontendIpConfigurationResponseOutput)
}

type InboundNatPool struct {
	BackendPort             int          `pulumi:"backendPort"`
	Etag                    *string      `pulumi:"etag"`
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    int          `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  int          `pulumi:"frontendPortRangeStart"`
	Id                      *string      `pulumi:"id"`
	Name                    *string      `pulumi:"name"`
	Protocol                string       `pulumi:"protocol"`
	ProvisioningState       *string      `pulumi:"provisioningState"`
}





type InboundNatPoolInput interface {
	pulumi.Input

	ToInboundNatPoolOutput() InboundNatPoolOutput
	ToInboundNatPoolOutputWithContext(context.Context) InboundNatPoolOutput
}

type InboundNatPoolArgs struct {
	BackendPort             pulumi.IntInput       `pulumi:"backendPort"`
	Etag                    pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    pulumi.IntInput       `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  pulumi.IntInput       `pulumi:"frontendPortRangeStart"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	Protocol                pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (InboundNatPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatPool)(nil)).Elem()
}

func (i InboundNatPoolArgs) ToInboundNatPoolOutput() InboundNatPoolOutput {
	return i.ToInboundNatPoolOutputWithContext(context.Background())
}

func (i InboundNatPoolArgs) ToInboundNatPoolOutputWithContext(ctx context.Context) InboundNatPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatPoolOutput)
}





type InboundNatPoolArrayInput interface {
	pulumi.Input

	ToInboundNatPoolArrayOutput() InboundNatPoolArrayOutput
	ToInboundNatPoolArrayOutputWithContext(context.Context) InboundNatPoolArrayOutput
}

type InboundNatPoolArray []InboundNatPoolInput

func (InboundNatPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatPool)(nil)).Elem()
}

func (i InboundNatPoolArray) ToInboundNatPoolArrayOutput() InboundNatPoolArrayOutput {
	return i.ToInboundNatPoolArrayOutputWithContext(context.Background())
}

func (i InboundNatPoolArray) ToInboundNatPoolArrayOutputWithContext(ctx context.Context) InboundNatPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatPoolArrayOutput)
}

type InboundNatPoolOutput struct{ *pulumi.OutputState }

func (InboundNatPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatPool)(nil)).Elem()
}

func (o InboundNatPoolOutput) ToInboundNatPoolOutput() InboundNatPoolOutput {
	return o
}

func (o InboundNatPoolOutput) ToInboundNatPoolOutputWithContext(ctx context.Context) InboundNatPoolOutput {
	return o
}

func (o InboundNatPoolOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPool) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o InboundNatPoolOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPool) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolOutput) FrontendIPConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v InboundNatPool) *SubResource { return v.FrontendIPConfiguration }).(SubResourcePtrOutput)
}

func (o InboundNatPoolOutput) FrontendPortRangeEnd() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPool) int { return v.FrontendPortRangeEnd }).(pulumi.IntOutput)
}

func (o InboundNatPoolOutput) FrontendPortRangeStart() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPool) int { return v.FrontendPortRangeStart }).(pulumi.IntOutput)
}

func (o InboundNatPoolOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPool) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPool) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v InboundNatPool) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o InboundNatPoolOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPool) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type InboundNatPoolArrayOutput struct{ *pulumi.OutputState }

func (InboundNatPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatPool)(nil)).Elem()
}

func (o InboundNatPoolArrayOutput) ToInboundNatPoolArrayOutput() InboundNatPoolArrayOutput {
	return o
}

func (o InboundNatPoolArrayOutput) ToInboundNatPoolArrayOutputWithContext(ctx context.Context) InboundNatPoolArrayOutput {
	return o
}

func (o InboundNatPoolArrayOutput) Index(i pulumi.IntInput) InboundNatPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatPool {
		return vs[0].([]InboundNatPool)[vs[1].(int)]
	}).(InboundNatPoolOutput)
}

type InboundNatPoolResponse struct {
	BackendPort             int                  `pulumi:"backendPort"`
	Etag                    *string              `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    int                  `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  int                  `pulumi:"frontendPortRangeStart"`
	Id                      *string              `pulumi:"id"`
	Name                    *string              `pulumi:"name"`
	Protocol                string               `pulumi:"protocol"`
	ProvisioningState       *string              `pulumi:"provisioningState"`
}





type InboundNatPoolResponseInput interface {
	pulumi.Input

	ToInboundNatPoolResponseOutput() InboundNatPoolResponseOutput
	ToInboundNatPoolResponseOutputWithContext(context.Context) InboundNatPoolResponseOutput
}

type InboundNatPoolResponseArgs struct {
	BackendPort             pulumi.IntInput             `pulumi:"backendPort"`
	Etag                    pulumi.StringPtrInput       `pulumi:"etag"`
	FrontendIPConfiguration SubResourceResponsePtrInput `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    pulumi.IntInput             `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  pulumi.IntInput             `pulumi:"frontendPortRangeStart"`
	Id                      pulumi.StringPtrInput       `pulumi:"id"`
	Name                    pulumi.StringPtrInput       `pulumi:"name"`
	Protocol                pulumi.StringInput          `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput       `pulumi:"provisioningState"`
}

func (InboundNatPoolResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatPoolResponse)(nil)).Elem()
}

func (i InboundNatPoolResponseArgs) ToInboundNatPoolResponseOutput() InboundNatPoolResponseOutput {
	return i.ToInboundNatPoolResponseOutputWithContext(context.Background())
}

func (i InboundNatPoolResponseArgs) ToInboundNatPoolResponseOutputWithContext(ctx context.Context) InboundNatPoolResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatPoolResponseOutput)
}





type InboundNatPoolResponseArrayInput interface {
	pulumi.Input

	ToInboundNatPoolResponseArrayOutput() InboundNatPoolResponseArrayOutput
	ToInboundNatPoolResponseArrayOutputWithContext(context.Context) InboundNatPoolResponseArrayOutput
}

type InboundNatPoolResponseArray []InboundNatPoolResponseInput

func (InboundNatPoolResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatPoolResponse)(nil)).Elem()
}

func (i InboundNatPoolResponseArray) ToInboundNatPoolResponseArrayOutput() InboundNatPoolResponseArrayOutput {
	return i.ToInboundNatPoolResponseArrayOutputWithContext(context.Background())
}

func (i InboundNatPoolResponseArray) ToInboundNatPoolResponseArrayOutputWithContext(ctx context.Context) InboundNatPoolResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatPoolResponseArrayOutput)
}

type InboundNatPoolResponseOutput struct{ *pulumi.OutputState }

func (InboundNatPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatPoolResponse)(nil)).Elem()
}

func (o InboundNatPoolResponseOutput) ToInboundNatPoolResponseOutput() InboundNatPoolResponseOutput {
	return o
}

func (o InboundNatPoolResponseOutput) ToInboundNatPoolResponseOutputWithContext(ctx context.Context) InboundNatPoolResponseOutput {
	return o
}

func (o InboundNatPoolResponseOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o InboundNatPoolResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolResponseOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o InboundNatPoolResponseOutput) FrontendPortRangeEnd() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) int { return v.FrontendPortRangeEnd }).(pulumi.IntOutput)
}

func (o InboundNatPoolResponseOutput) FrontendPortRangeStart() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) int { return v.FrontendPortRangeStart }).(pulumi.IntOutput)
}

func (o InboundNatPoolResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o InboundNatPoolResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type InboundNatPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (InboundNatPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatPoolResponse)(nil)).Elem()
}

func (o InboundNatPoolResponseArrayOutput) ToInboundNatPoolResponseArrayOutput() InboundNatPoolResponseArrayOutput {
	return o
}

func (o InboundNatPoolResponseArrayOutput) ToInboundNatPoolResponseArrayOutputWithContext(ctx context.Context) InboundNatPoolResponseArrayOutput {
	return o
}

func (o InboundNatPoolResponseArrayOutput) Index(i pulumi.IntInput) InboundNatPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatPoolResponse {
		return vs[0].([]InboundNatPoolResponse)[vs[1].(int)]
	}).(InboundNatPoolResponseOutput)
}

type InboundNatRule struct {
	BackendIPConfiguration  *SubResource `pulumi:"backendIPConfiguration"`
	BackendPort             *int         `pulumi:"backendPort"`
	EnableFloatingIP        bool         `pulumi:"enableFloatingIP"`
	Etag                    *string      `pulumi:"etag"`
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPort            int          `pulumi:"frontendPort"`
	Id                      *string      `pulumi:"id"`
	IdleTimeoutInMinutes    *int         `pulumi:"idleTimeoutInMinutes"`
	Name                    *string      `pulumi:"name"`
	Protocol                string       `pulumi:"protocol"`
	ProvisioningState       *string      `pulumi:"provisioningState"`
}





type InboundNatRuleInput interface {
	pulumi.Input

	ToInboundNatRuleOutput() InboundNatRuleOutput
	ToInboundNatRuleOutputWithContext(context.Context) InboundNatRuleOutput
}

type InboundNatRuleArgs struct {
	BackendIPConfiguration  SubResourcePtrInput   `pulumi:"backendIPConfiguration"`
	BackendPort             pulumi.IntPtrInput    `pulumi:"backendPort"`
	EnableFloatingIP        pulumi.BoolInput      `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPort            pulumi.IntInput       `pulumi:"frontendPort"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	IdleTimeoutInMinutes    pulumi.IntPtrInput    `pulumi:"idleTimeoutInMinutes"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	Protocol                pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (InboundNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRule)(nil)).Elem()
}

func (i InboundNatRuleArgs) ToInboundNatRuleOutput() InboundNatRuleOutput {
	return i.ToInboundNatRuleOutputWithContext(context.Background())
}

func (i InboundNatRuleArgs) ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleOutput)
}





type InboundNatRuleArrayInput interface {
	pulumi.Input

	ToInboundNatRuleArrayOutput() InboundNatRuleArrayOutput
	ToInboundNatRuleArrayOutputWithContext(context.Context) InboundNatRuleArrayOutput
}

type InboundNatRuleArray []InboundNatRuleInput

func (InboundNatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRule)(nil)).Elem()
}

func (i InboundNatRuleArray) ToInboundNatRuleArrayOutput() InboundNatRuleArrayOutput {
	return i.ToInboundNatRuleArrayOutputWithContext(context.Background())
}

func (i InboundNatRuleArray) ToInboundNatRuleArrayOutputWithContext(ctx context.Context) InboundNatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleArrayOutput)
}

type InboundNatRuleOutput struct{ *pulumi.OutputState }

func (InboundNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRule)(nil)).Elem()
}

func (o InboundNatRuleOutput) ToInboundNatRuleOutput() InboundNatRuleOutput {
	return o
}

func (o InboundNatRuleOutput) ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput {
	return o
}

func (o InboundNatRuleOutput) BackendIPConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v InboundNatRule) *SubResource { return v.BackendIPConfiguration }).(SubResourcePtrOutput)
}

func (o InboundNatRuleOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleOutput) EnableFloatingIP() pulumi.BoolOutput {
	return o.ApplyT(func(v InboundNatRule) bool { return v.EnableFloatingIP }).(pulumi.BoolOutput)
}

func (o InboundNatRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleOutput) FrontendIPConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v InboundNatRule) *SubResource { return v.FrontendIPConfiguration }).(SubResourcePtrOutput)
}

func (o InboundNatRuleOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatRule) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o InboundNatRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v InboundNatRule) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o InboundNatRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type InboundNatRuleArrayOutput struct{ *pulumi.OutputState }

func (InboundNatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRule)(nil)).Elem()
}

func (o InboundNatRuleArrayOutput) ToInboundNatRuleArrayOutput() InboundNatRuleArrayOutput {
	return o
}

func (o InboundNatRuleArrayOutput) ToInboundNatRuleArrayOutputWithContext(ctx context.Context) InboundNatRuleArrayOutput {
	return o
}

func (o InboundNatRuleArrayOutput) Index(i pulumi.IntInput) InboundNatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatRule {
		return vs[0].([]InboundNatRule)[vs[1].(int)]
	}).(InboundNatRuleOutput)
}

type InboundNatRuleResponse struct {
	BackendIPConfiguration  *SubResourceResponse `pulumi:"backendIPConfiguration"`
	BackendPort             *int                 `pulumi:"backendPort"`
	EnableFloatingIP        bool                 `pulumi:"enableFloatingIP"`
	Etag                    *string              `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	FrontendPort            int                  `pulumi:"frontendPort"`
	Id                      *string              `pulumi:"id"`
	IdleTimeoutInMinutes    *int                 `pulumi:"idleTimeoutInMinutes"`
	Name                    *string              `pulumi:"name"`
	Protocol                string               `pulumi:"protocol"`
	ProvisioningState       *string              `pulumi:"provisioningState"`
}





type InboundNatRuleResponseInput interface {
	pulumi.Input

	ToInboundNatRuleResponseOutput() InboundNatRuleResponseOutput
	ToInboundNatRuleResponseOutputWithContext(context.Context) InboundNatRuleResponseOutput
}

type InboundNatRuleResponseArgs struct {
	BackendIPConfiguration  SubResourceResponsePtrInput `pulumi:"backendIPConfiguration"`
	BackendPort             pulumi.IntPtrInput          `pulumi:"backendPort"`
	EnableFloatingIP        pulumi.BoolInput            `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrInput       `pulumi:"etag"`
	FrontendIPConfiguration SubResourceResponsePtrInput `pulumi:"frontendIPConfiguration"`
	FrontendPort            pulumi.IntInput             `pulumi:"frontendPort"`
	Id                      pulumi.StringPtrInput       `pulumi:"id"`
	IdleTimeoutInMinutes    pulumi.IntPtrInput          `pulumi:"idleTimeoutInMinutes"`
	Name                    pulumi.StringPtrInput       `pulumi:"name"`
	Protocol                pulumi.StringInput          `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput       `pulumi:"provisioningState"`
}

func (InboundNatRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRuleResponse)(nil)).Elem()
}

func (i InboundNatRuleResponseArgs) ToInboundNatRuleResponseOutput() InboundNatRuleResponseOutput {
	return i.ToInboundNatRuleResponseOutputWithContext(context.Background())
}

func (i InboundNatRuleResponseArgs) ToInboundNatRuleResponseOutputWithContext(ctx context.Context) InboundNatRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleResponseOutput)
}





type InboundNatRuleResponseArrayInput interface {
	pulumi.Input

	ToInboundNatRuleResponseArrayOutput() InboundNatRuleResponseArrayOutput
	ToInboundNatRuleResponseArrayOutputWithContext(context.Context) InboundNatRuleResponseArrayOutput
}

type InboundNatRuleResponseArray []InboundNatRuleResponseInput

func (InboundNatRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRuleResponse)(nil)).Elem()
}

func (i InboundNatRuleResponseArray) ToInboundNatRuleResponseArrayOutput() InboundNatRuleResponseArrayOutput {
	return i.ToInboundNatRuleResponseArrayOutputWithContext(context.Background())
}

func (i InboundNatRuleResponseArray) ToInboundNatRuleResponseArrayOutputWithContext(ctx context.Context) InboundNatRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleResponseArrayOutput)
}

type InboundNatRuleResponseOutput struct{ *pulumi.OutputState }

func (InboundNatRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRuleResponse)(nil)).Elem()
}

func (o InboundNatRuleResponseOutput) ToInboundNatRuleResponseOutput() InboundNatRuleResponseOutput {
	return o
}

func (o InboundNatRuleResponseOutput) ToInboundNatRuleResponseOutputWithContext(ctx context.Context) InboundNatRuleResponseOutput {
	return o
}

func (o InboundNatRuleResponseOutput) BackendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *SubResourceResponse { return v.BackendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o InboundNatRuleResponseOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleResponseOutput) EnableFloatingIP() pulumi.BoolOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) bool { return v.EnableFloatingIP }).(pulumi.BoolOutput)
}

func (o InboundNatRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleResponseOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o InboundNatRuleResponseOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o InboundNatRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o InboundNatRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type InboundNatRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (InboundNatRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRuleResponse)(nil)).Elem()
}

func (o InboundNatRuleResponseArrayOutput) ToInboundNatRuleResponseArrayOutput() InboundNatRuleResponseArrayOutput {
	return o
}

func (o InboundNatRuleResponseArrayOutput) ToInboundNatRuleResponseArrayOutputWithContext(ctx context.Context) InboundNatRuleResponseArrayOutput {
	return o
}

func (o InboundNatRuleResponseArrayOutput) Index(i pulumi.IntInput) InboundNatRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatRuleResponse {
		return vs[0].([]InboundNatRuleResponse)[vs[1].(int)]
	}).(InboundNatRuleResponseOutput)
}

type LoadBalancingRule struct {
	BackendAddressPool      SubResource  `pulumi:"backendAddressPool"`
	BackendPort             *int         `pulumi:"backendPort"`
	EnableFloatingIP        bool         `pulumi:"enableFloatingIP"`
	Etag                    *string      `pulumi:"etag"`
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPort            int          `pulumi:"frontendPort"`
	Id                      *string      `pulumi:"id"`
	IdleTimeoutInMinutes    *int         `pulumi:"idleTimeoutInMinutes"`
	LoadDistribution        *string      `pulumi:"loadDistribution"`
	Name                    *string      `pulumi:"name"`
	Probe                   *SubResource `pulumi:"probe"`
	Protocol                string       `pulumi:"protocol"`
	ProvisioningState       *string      `pulumi:"provisioningState"`
}





type LoadBalancingRuleInput interface {
	pulumi.Input

	ToLoadBalancingRuleOutput() LoadBalancingRuleOutput
	ToLoadBalancingRuleOutputWithContext(context.Context) LoadBalancingRuleOutput
}

type LoadBalancingRuleArgs struct {
	BackendAddressPool      SubResourceInput      `pulumi:"backendAddressPool"`
	BackendPort             pulumi.IntPtrInput    `pulumi:"backendPort"`
	EnableFloatingIP        pulumi.BoolInput      `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPort            pulumi.IntInput       `pulumi:"frontendPort"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	IdleTimeoutInMinutes    pulumi.IntPtrInput    `pulumi:"idleTimeoutInMinutes"`
	LoadDistribution        pulumi.StringPtrInput `pulumi:"loadDistribution"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	Probe                   SubResourcePtrInput   `pulumi:"probe"`
	Protocol                pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (LoadBalancingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRule)(nil)).Elem()
}

func (i LoadBalancingRuleArgs) ToLoadBalancingRuleOutput() LoadBalancingRuleOutput {
	return i.ToLoadBalancingRuleOutputWithContext(context.Background())
}

func (i LoadBalancingRuleArgs) ToLoadBalancingRuleOutputWithContext(ctx context.Context) LoadBalancingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleOutput)
}





type LoadBalancingRuleArrayInput interface {
	pulumi.Input

	ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput
	ToLoadBalancingRuleArrayOutputWithContext(context.Context) LoadBalancingRuleArrayOutput
}

type LoadBalancingRuleArray []LoadBalancingRuleInput

func (LoadBalancingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRule)(nil)).Elem()
}

func (i LoadBalancingRuleArray) ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput {
	return i.ToLoadBalancingRuleArrayOutputWithContext(context.Background())
}

func (i LoadBalancingRuleArray) ToLoadBalancingRuleArrayOutputWithContext(ctx context.Context) LoadBalancingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleArrayOutput)
}

type LoadBalancingRuleOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRule)(nil)).Elem()
}

func (o LoadBalancingRuleOutput) ToLoadBalancingRuleOutput() LoadBalancingRuleOutput {
	return o
}

func (o LoadBalancingRuleOutput) ToLoadBalancingRuleOutputWithContext(ctx context.Context) LoadBalancingRuleOutput {
	return o
}

func (o LoadBalancingRuleOutput) BackendAddressPool() SubResourceOutput {
	return o.ApplyT(func(v LoadBalancingRule) SubResource { return v.BackendAddressPool }).(SubResourceOutput)
}

func (o LoadBalancingRuleOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingRuleOutput) EnableFloatingIP() pulumi.BoolOutput {
	return o.ApplyT(func(v LoadBalancingRule) bool { return v.EnableFloatingIP }).(pulumi.BoolOutput)
}

func (o LoadBalancingRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) FrontendIPConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *SubResource { return v.FrontendIPConfiguration }).(SubResourcePtrOutput)
}

func (o LoadBalancingRuleOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRule) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingRuleOutput) LoadDistribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.LoadDistribution }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) Probe() SubResourcePtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *SubResource { return v.Probe }).(SubResourcePtrOutput)
}

func (o LoadBalancingRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRule) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LoadBalancingRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type LoadBalancingRuleArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRule)(nil)).Elem()
}

func (o LoadBalancingRuleArrayOutput) ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput {
	return o
}

func (o LoadBalancingRuleArrayOutput) ToLoadBalancingRuleArrayOutputWithContext(ctx context.Context) LoadBalancingRuleArrayOutput {
	return o
}

func (o LoadBalancingRuleArrayOutput) Index(i pulumi.IntInput) LoadBalancingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingRule {
		return vs[0].([]LoadBalancingRule)[vs[1].(int)]
	}).(LoadBalancingRuleOutput)
}

type LoadBalancingRuleResponse struct {
	BackendAddressPool      SubResourceResponse  `pulumi:"backendAddressPool"`
	BackendPort             *int                 `pulumi:"backendPort"`
	EnableFloatingIP        bool                 `pulumi:"enableFloatingIP"`
	Etag                    *string              `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	FrontendPort            int                  `pulumi:"frontendPort"`
	Id                      *string              `pulumi:"id"`
	IdleTimeoutInMinutes    *int                 `pulumi:"idleTimeoutInMinutes"`
	LoadDistribution        *string              `pulumi:"loadDistribution"`
	Name                    *string              `pulumi:"name"`
	Probe                   *SubResourceResponse `pulumi:"probe"`
	Protocol                string               `pulumi:"protocol"`
	ProvisioningState       *string              `pulumi:"provisioningState"`
}





type LoadBalancingRuleResponseInput interface {
	pulumi.Input

	ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput
	ToLoadBalancingRuleResponseOutputWithContext(context.Context) LoadBalancingRuleResponseOutput
}

type LoadBalancingRuleResponseArgs struct {
	BackendAddressPool      SubResourceResponseInput    `pulumi:"backendAddressPool"`
	BackendPort             pulumi.IntPtrInput          `pulumi:"backendPort"`
	EnableFloatingIP        pulumi.BoolInput            `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrInput       `pulumi:"etag"`
	FrontendIPConfiguration SubResourceResponsePtrInput `pulumi:"frontendIPConfiguration"`
	FrontendPort            pulumi.IntInput             `pulumi:"frontendPort"`
	Id                      pulumi.StringPtrInput       `pulumi:"id"`
	IdleTimeoutInMinutes    pulumi.IntPtrInput          `pulumi:"idleTimeoutInMinutes"`
	LoadDistribution        pulumi.StringPtrInput       `pulumi:"loadDistribution"`
	Name                    pulumi.StringPtrInput       `pulumi:"name"`
	Probe                   SubResourceResponsePtrInput `pulumi:"probe"`
	Protocol                pulumi.StringInput          `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput       `pulumi:"provisioningState"`
}

func (LoadBalancingRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRuleResponse)(nil)).Elem()
}

func (i LoadBalancingRuleResponseArgs) ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput {
	return i.ToLoadBalancingRuleResponseOutputWithContext(context.Background())
}

func (i LoadBalancingRuleResponseArgs) ToLoadBalancingRuleResponseOutputWithContext(ctx context.Context) LoadBalancingRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleResponseOutput)
}





type LoadBalancingRuleResponseArrayInput interface {
	pulumi.Input

	ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput
	ToLoadBalancingRuleResponseArrayOutputWithContext(context.Context) LoadBalancingRuleResponseArrayOutput
}

type LoadBalancingRuleResponseArray []LoadBalancingRuleResponseInput

func (LoadBalancingRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRuleResponse)(nil)).Elem()
}

func (i LoadBalancingRuleResponseArray) ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput {
	return i.ToLoadBalancingRuleResponseArrayOutputWithContext(context.Background())
}

func (i LoadBalancingRuleResponseArray) ToLoadBalancingRuleResponseArrayOutputWithContext(ctx context.Context) LoadBalancingRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleResponseArrayOutput)
}

type LoadBalancingRuleResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRuleResponse)(nil)).Elem()
}

func (o LoadBalancingRuleResponseOutput) ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput {
	return o
}

func (o LoadBalancingRuleResponseOutput) ToLoadBalancingRuleResponseOutputWithContext(ctx context.Context) LoadBalancingRuleResponseOutput {
	return o
}

func (o LoadBalancingRuleResponseOutput) BackendAddressPool() SubResourceResponseOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponseOutput)
}

func (o LoadBalancingRuleResponseOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) EnableFloatingIP() pulumi.BoolOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) bool { return v.EnableFloatingIP }).(pulumi.BoolOutput)
}

func (o LoadBalancingRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o LoadBalancingRuleResponseOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) LoadDistribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.LoadDistribution }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) Probe() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *SubResourceResponse { return v.Probe }).(SubResourceResponsePtrOutput)
}

func (o LoadBalancingRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LoadBalancingRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type LoadBalancingRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRuleResponse)(nil)).Elem()
}

func (o LoadBalancingRuleResponseArrayOutput) ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput {
	return o
}

func (o LoadBalancingRuleResponseArrayOutput) ToLoadBalancingRuleResponseArrayOutputWithContext(ctx context.Context) LoadBalancingRuleResponseArrayOutput {
	return o
}

func (o LoadBalancingRuleResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancingRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingRuleResponse {
		return vs[0].([]LoadBalancingRuleResponse)[vs[1].(int)]
	}).(LoadBalancingRuleResponseOutput)
}

type NetworkInterfaceDnsSettings struct {
	AppliedDnsServers    []string `pulumi:"appliedDnsServers"`
	DnsServers           []string `pulumi:"dnsServers"`
	InternalDnsNameLabel *string  `pulumi:"internalDnsNameLabel"`
	InternalFqdn         *string  `pulumi:"internalFqdn"`
}





type NetworkInterfaceDnsSettingsInput interface {
	pulumi.Input

	ToNetworkInterfaceDnsSettingsOutput() NetworkInterfaceDnsSettingsOutput
	ToNetworkInterfaceDnsSettingsOutputWithContext(context.Context) NetworkInterfaceDnsSettingsOutput
}

type NetworkInterfaceDnsSettingsArgs struct {
	AppliedDnsServers    pulumi.StringArrayInput `pulumi:"appliedDnsServers"`
	DnsServers           pulumi.StringArrayInput `pulumi:"dnsServers"`
	InternalDnsNameLabel pulumi.StringPtrInput   `pulumi:"internalDnsNameLabel"`
	InternalFqdn         pulumi.StringPtrInput   `pulumi:"internalFqdn"`
}

func (NetworkInterfaceDnsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceDnsSettings)(nil)).Elem()
}

func (i NetworkInterfaceDnsSettingsArgs) ToNetworkInterfaceDnsSettingsOutput() NetworkInterfaceDnsSettingsOutput {
	return i.ToNetworkInterfaceDnsSettingsOutputWithContext(context.Background())
}

func (i NetworkInterfaceDnsSettingsArgs) ToNetworkInterfaceDnsSettingsOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceDnsSettingsOutput)
}

func (i NetworkInterfaceDnsSettingsArgs) ToNetworkInterfaceDnsSettingsPtrOutput() NetworkInterfaceDnsSettingsPtrOutput {
	return i.ToNetworkInterfaceDnsSettingsPtrOutputWithContext(context.Background())
}

func (i NetworkInterfaceDnsSettingsArgs) ToNetworkInterfaceDnsSettingsPtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceDnsSettingsOutput).ToNetworkInterfaceDnsSettingsPtrOutputWithContext(ctx)
}









type NetworkInterfaceDnsSettingsPtrInput interface {
	pulumi.Input

	ToNetworkInterfaceDnsSettingsPtrOutput() NetworkInterfaceDnsSettingsPtrOutput
	ToNetworkInterfaceDnsSettingsPtrOutputWithContext(context.Context) NetworkInterfaceDnsSettingsPtrOutput
}

type networkInterfaceDnsSettingsPtrType NetworkInterfaceDnsSettingsArgs

func NetworkInterfaceDnsSettingsPtr(v *NetworkInterfaceDnsSettingsArgs) NetworkInterfaceDnsSettingsPtrInput {
	return (*networkInterfaceDnsSettingsPtrType)(v)
}

func (*networkInterfaceDnsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceDnsSettings)(nil)).Elem()
}

func (i *networkInterfaceDnsSettingsPtrType) ToNetworkInterfaceDnsSettingsPtrOutput() NetworkInterfaceDnsSettingsPtrOutput {
	return i.ToNetworkInterfaceDnsSettingsPtrOutputWithContext(context.Background())
}

func (i *networkInterfaceDnsSettingsPtrType) ToNetworkInterfaceDnsSettingsPtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceDnsSettingsPtrOutput)
}

type NetworkInterfaceDnsSettingsOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceDnsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceDnsSettings)(nil)).Elem()
}

func (o NetworkInterfaceDnsSettingsOutput) ToNetworkInterfaceDnsSettingsOutput() NetworkInterfaceDnsSettingsOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsOutput) ToNetworkInterfaceDnsSettingsOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsOutput) ToNetworkInterfaceDnsSettingsPtrOutput() NetworkInterfaceDnsSettingsPtrOutput {
	return o.ToNetworkInterfaceDnsSettingsPtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceDnsSettingsOutput) ToNetworkInterfaceDnsSettingsPtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkInterfaceDnsSettings) *NetworkInterfaceDnsSettings {
		return &v
	}).(NetworkInterfaceDnsSettingsPtrOutput)
}

func (o NetworkInterfaceDnsSettingsOutput) AppliedDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettings) []string { return v.AppliedDnsServers }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettings) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsOutput) InternalDnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettings) *string { return v.InternalDnsNameLabel }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsOutput) InternalFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettings) *string { return v.InternalFqdn }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceDnsSettingsPtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceDnsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceDnsSettings)(nil)).Elem()
}

func (o NetworkInterfaceDnsSettingsPtrOutput) ToNetworkInterfaceDnsSettingsPtrOutput() NetworkInterfaceDnsSettingsPtrOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsPtrOutput) ToNetworkInterfaceDnsSettingsPtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsPtrOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsPtrOutput) Elem() NetworkInterfaceDnsSettingsOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) NetworkInterfaceDnsSettings {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceDnsSettings
		return ret
	}).(NetworkInterfaceDnsSettingsOutput)
}

func (o NetworkInterfaceDnsSettingsPtrOutput) AppliedDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) []string {
		if v == nil {
			return nil
		}
		return v.AppliedDnsServers
	}).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsPtrOutput) InternalDnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.InternalDnsNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsPtrOutput) InternalFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.InternalFqdn
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceDnsSettingsResponse struct {
	AppliedDnsServers    []string `pulumi:"appliedDnsServers"`
	DnsServers           []string `pulumi:"dnsServers"`
	InternalDnsNameLabel *string  `pulumi:"internalDnsNameLabel"`
	InternalFqdn         *string  `pulumi:"internalFqdn"`
}





type NetworkInterfaceDnsSettingsResponseInput interface {
	pulumi.Input

	ToNetworkInterfaceDnsSettingsResponseOutput() NetworkInterfaceDnsSettingsResponseOutput
	ToNetworkInterfaceDnsSettingsResponseOutputWithContext(context.Context) NetworkInterfaceDnsSettingsResponseOutput
}

type NetworkInterfaceDnsSettingsResponseArgs struct {
	AppliedDnsServers    pulumi.StringArrayInput `pulumi:"appliedDnsServers"`
	DnsServers           pulumi.StringArrayInput `pulumi:"dnsServers"`
	InternalDnsNameLabel pulumi.StringPtrInput   `pulumi:"internalDnsNameLabel"`
	InternalFqdn         pulumi.StringPtrInput   `pulumi:"internalFqdn"`
}

func (NetworkInterfaceDnsSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceDnsSettingsResponse)(nil)).Elem()
}

func (i NetworkInterfaceDnsSettingsResponseArgs) ToNetworkInterfaceDnsSettingsResponseOutput() NetworkInterfaceDnsSettingsResponseOutput {
	return i.ToNetworkInterfaceDnsSettingsResponseOutputWithContext(context.Background())
}

func (i NetworkInterfaceDnsSettingsResponseArgs) ToNetworkInterfaceDnsSettingsResponseOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceDnsSettingsResponseOutput)
}

func (i NetworkInterfaceDnsSettingsResponseArgs) ToNetworkInterfaceDnsSettingsResponsePtrOutput() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return i.ToNetworkInterfaceDnsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i NetworkInterfaceDnsSettingsResponseArgs) ToNetworkInterfaceDnsSettingsResponsePtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceDnsSettingsResponseOutput).ToNetworkInterfaceDnsSettingsResponsePtrOutputWithContext(ctx)
}









type NetworkInterfaceDnsSettingsResponsePtrInput interface {
	pulumi.Input

	ToNetworkInterfaceDnsSettingsResponsePtrOutput() NetworkInterfaceDnsSettingsResponsePtrOutput
	ToNetworkInterfaceDnsSettingsResponsePtrOutputWithContext(context.Context) NetworkInterfaceDnsSettingsResponsePtrOutput
}

type networkInterfaceDnsSettingsResponsePtrType NetworkInterfaceDnsSettingsResponseArgs

func NetworkInterfaceDnsSettingsResponsePtr(v *NetworkInterfaceDnsSettingsResponseArgs) NetworkInterfaceDnsSettingsResponsePtrInput {
	return (*networkInterfaceDnsSettingsResponsePtrType)(v)
}

func (*networkInterfaceDnsSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceDnsSettingsResponse)(nil)).Elem()
}

func (i *networkInterfaceDnsSettingsResponsePtrType) ToNetworkInterfaceDnsSettingsResponsePtrOutput() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return i.ToNetworkInterfaceDnsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *networkInterfaceDnsSettingsResponsePtrType) ToNetworkInterfaceDnsSettingsResponsePtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

type NetworkInterfaceDnsSettingsResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceDnsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceDnsSettingsResponse)(nil)).Elem()
}

func (o NetworkInterfaceDnsSettingsResponseOutput) ToNetworkInterfaceDnsSettingsResponseOutput() NetworkInterfaceDnsSettingsResponseOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsResponseOutput) ToNetworkInterfaceDnsSettingsResponseOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsResponseOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsResponseOutput) ToNetworkInterfaceDnsSettingsResponsePtrOutput() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ToNetworkInterfaceDnsSettingsResponsePtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceDnsSettingsResponseOutput) ToNetworkInterfaceDnsSettingsResponsePtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkInterfaceDnsSettingsResponse) *NetworkInterfaceDnsSettingsResponse {
		return &v
	}).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

func (o NetworkInterfaceDnsSettingsResponseOutput) AppliedDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettingsResponse) []string { return v.AppliedDnsServers }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettingsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsResponseOutput) InternalDnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettingsResponse) *string { return v.InternalDnsNameLabel }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsResponseOutput) InternalFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettingsResponse) *string { return v.InternalFqdn }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceDnsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceDnsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceDnsSettingsResponse)(nil)).Elem()
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) ToNetworkInterfaceDnsSettingsResponsePtrOutput() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) ToNetworkInterfaceDnsSettingsResponsePtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) Elem() NetworkInterfaceDnsSettingsResponseOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) NetworkInterfaceDnsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceDnsSettingsResponse
		return ret
	}).(NetworkInterfaceDnsSettingsResponseOutput)
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) AppliedDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.AppliedDnsServers
	}).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) InternalDnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InternalDnsNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) InternalFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InternalFqdn
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceIpConfiguration struct {
	Etag                            *string       `pulumi:"etag"`
	Id                              *string       `pulumi:"id"`
	LoadBalancerBackendAddressPools []SubResource `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules     []SubResource `pulumi:"loadBalancerInboundNatRules"`
	Name                            *string       `pulumi:"name"`
	PrivateIPAddress                *string       `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod       *string       `pulumi:"privateIPAllocationMethod"`
	ProvisioningState               *string       `pulumi:"provisioningState"`
	PublicIPAddress                 *SubResource  `pulumi:"publicIPAddress"`
	Subnet                          *SubResource  `pulumi:"subnet"`
}





type NetworkInterfaceIpConfigurationInput interface {
	pulumi.Input

	ToNetworkInterfaceIpConfigurationOutput() NetworkInterfaceIpConfigurationOutput
	ToNetworkInterfaceIpConfigurationOutputWithContext(context.Context) NetworkInterfaceIpConfigurationOutput
}

type NetworkInterfaceIpConfigurationArgs struct {
	Etag                            pulumi.StringPtrInput `pulumi:"etag"`
	Id                              pulumi.StringPtrInput `pulumi:"id"`
	LoadBalancerBackendAddressPools SubResourceArrayInput `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules     SubResourceArrayInput `pulumi:"loadBalancerInboundNatRules"`
	Name                            pulumi.StringPtrInput `pulumi:"name"`
	PrivateIPAddress                pulumi.StringPtrInput `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod       pulumi.StringPtrInput `pulumi:"privateIPAllocationMethod"`
	ProvisioningState               pulumi.StringPtrInput `pulumi:"provisioningState"`
	PublicIPAddress                 SubResourcePtrInput   `pulumi:"publicIPAddress"`
	Subnet                          SubResourcePtrInput   `pulumi:"subnet"`
}

func (NetworkInterfaceIpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIpConfiguration)(nil)).Elem()
}

func (i NetworkInterfaceIpConfigurationArgs) ToNetworkInterfaceIpConfigurationOutput() NetworkInterfaceIpConfigurationOutput {
	return i.ToNetworkInterfaceIpConfigurationOutputWithContext(context.Background())
}

func (i NetworkInterfaceIpConfigurationArgs) ToNetworkInterfaceIpConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceIpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIpConfigurationOutput)
}





type NetworkInterfaceIpConfigurationArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceIpConfigurationArrayOutput() NetworkInterfaceIpConfigurationArrayOutput
	ToNetworkInterfaceIpConfigurationArrayOutputWithContext(context.Context) NetworkInterfaceIpConfigurationArrayOutput
}

type NetworkInterfaceIpConfigurationArray []NetworkInterfaceIpConfigurationInput

func (NetworkInterfaceIpConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIpConfiguration)(nil)).Elem()
}

func (i NetworkInterfaceIpConfigurationArray) ToNetworkInterfaceIpConfigurationArrayOutput() NetworkInterfaceIpConfigurationArrayOutput {
	return i.ToNetworkInterfaceIpConfigurationArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceIpConfigurationArray) ToNetworkInterfaceIpConfigurationArrayOutputWithContext(ctx context.Context) NetworkInterfaceIpConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIpConfigurationArrayOutput)
}

type NetworkInterfaceIpConfigurationOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIpConfiguration)(nil)).Elem()
}

func (o NetworkInterfaceIpConfigurationOutput) ToNetworkInterfaceIpConfigurationOutput() NetworkInterfaceIpConfigurationOutput {
	return o
}

func (o NetworkInterfaceIpConfigurationOutput) ToNetworkInterfaceIpConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceIpConfigurationOutput {
	return o
}

func (o NetworkInterfaceIpConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationOutput) LoadBalancerBackendAddressPools() SubResourceArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfiguration) []SubResource { return v.LoadBalancerBackendAddressPools }).(SubResourceArrayOutput)
}

func (o NetworkInterfaceIpConfigurationOutput) LoadBalancerInboundNatRules() SubResourceArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfiguration) []SubResource { return v.LoadBalancerInboundNatRules }).(SubResourceArrayOutput)
}

func (o NetworkInterfaceIpConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfiguration) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfiguration) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationOutput) PublicIPAddress() SubResourcePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfiguration) *SubResource { return v.PublicIPAddress }).(SubResourcePtrOutput)
}

func (o NetworkInterfaceIpConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type NetworkInterfaceIpConfigurationArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIpConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIpConfiguration)(nil)).Elem()
}

func (o NetworkInterfaceIpConfigurationArrayOutput) ToNetworkInterfaceIpConfigurationArrayOutput() NetworkInterfaceIpConfigurationArrayOutput {
	return o
}

func (o NetworkInterfaceIpConfigurationArrayOutput) ToNetworkInterfaceIpConfigurationArrayOutputWithContext(ctx context.Context) NetworkInterfaceIpConfigurationArrayOutput {
	return o
}

func (o NetworkInterfaceIpConfigurationArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceIpConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceIpConfiguration {
		return vs[0].([]NetworkInterfaceIpConfiguration)[vs[1].(int)]
	}).(NetworkInterfaceIpConfigurationOutput)
}

type NetworkInterfaceIpConfigurationResponse struct {
	Etag                            *string               `pulumi:"etag"`
	Id                              *string               `pulumi:"id"`
	LoadBalancerBackendAddressPools []SubResourceResponse `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules     []SubResourceResponse `pulumi:"loadBalancerInboundNatRules"`
	Name                            *string               `pulumi:"name"`
	PrivateIPAddress                *string               `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod       *string               `pulumi:"privateIPAllocationMethod"`
	ProvisioningState               *string               `pulumi:"provisioningState"`
	PublicIPAddress                 *SubResourceResponse  `pulumi:"publicIPAddress"`
	Subnet                          *SubResourceResponse  `pulumi:"subnet"`
}





type NetworkInterfaceIpConfigurationResponseInput interface {
	pulumi.Input

	ToNetworkInterfaceIpConfigurationResponseOutput() NetworkInterfaceIpConfigurationResponseOutput
	ToNetworkInterfaceIpConfigurationResponseOutputWithContext(context.Context) NetworkInterfaceIpConfigurationResponseOutput
}

type NetworkInterfaceIpConfigurationResponseArgs struct {
	Etag                            pulumi.StringPtrInput         `pulumi:"etag"`
	Id                              pulumi.StringPtrInput         `pulumi:"id"`
	LoadBalancerBackendAddressPools SubResourceResponseArrayInput `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules     SubResourceResponseArrayInput `pulumi:"loadBalancerInboundNatRules"`
	Name                            pulumi.StringPtrInput         `pulumi:"name"`
	PrivateIPAddress                pulumi.StringPtrInput         `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod       pulumi.StringPtrInput         `pulumi:"privateIPAllocationMethod"`
	ProvisioningState               pulumi.StringPtrInput         `pulumi:"provisioningState"`
	PublicIPAddress                 SubResourceResponsePtrInput   `pulumi:"publicIPAddress"`
	Subnet                          SubResourceResponsePtrInput   `pulumi:"subnet"`
}

func (NetworkInterfaceIpConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIpConfigurationResponse)(nil)).Elem()
}

func (i NetworkInterfaceIpConfigurationResponseArgs) ToNetworkInterfaceIpConfigurationResponseOutput() NetworkInterfaceIpConfigurationResponseOutput {
	return i.ToNetworkInterfaceIpConfigurationResponseOutputWithContext(context.Background())
}

func (i NetworkInterfaceIpConfigurationResponseArgs) ToNetworkInterfaceIpConfigurationResponseOutputWithContext(ctx context.Context) NetworkInterfaceIpConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIpConfigurationResponseOutput)
}





type NetworkInterfaceIpConfigurationResponseArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceIpConfigurationResponseArrayOutput() NetworkInterfaceIpConfigurationResponseArrayOutput
	ToNetworkInterfaceIpConfigurationResponseArrayOutputWithContext(context.Context) NetworkInterfaceIpConfigurationResponseArrayOutput
}

type NetworkInterfaceIpConfigurationResponseArray []NetworkInterfaceIpConfigurationResponseInput

func (NetworkInterfaceIpConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIpConfigurationResponse)(nil)).Elem()
}

func (i NetworkInterfaceIpConfigurationResponseArray) ToNetworkInterfaceIpConfigurationResponseArrayOutput() NetworkInterfaceIpConfigurationResponseArrayOutput {
	return i.ToNetworkInterfaceIpConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceIpConfigurationResponseArray) ToNetworkInterfaceIpConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceIpConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIpConfigurationResponseArrayOutput)
}

type NetworkInterfaceIpConfigurationResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIpConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIpConfigurationResponse)(nil)).Elem()
}

func (o NetworkInterfaceIpConfigurationResponseOutput) ToNetworkInterfaceIpConfigurationResponseOutput() NetworkInterfaceIpConfigurationResponseOutput {
	return o
}

func (o NetworkInterfaceIpConfigurationResponseOutput) ToNetworkInterfaceIpConfigurationResponseOutputWithContext(ctx context.Context) NetworkInterfaceIpConfigurationResponseOutput {
	return o
}

func (o NetworkInterfaceIpConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationResponseOutput) LoadBalancerBackendAddressPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfigurationResponse) []SubResourceResponse {
		return v.LoadBalancerBackendAddressPools
	}).(SubResourceResponseArrayOutput)
}

func (o NetworkInterfaceIpConfigurationResponseOutput) LoadBalancerInboundNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfigurationResponse) []SubResourceResponse {
		return v.LoadBalancerInboundNatRules
	}).(SubResourceResponseArrayOutput)
}

func (o NetworkInterfaceIpConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIpConfigurationResponseOutput) PublicIPAddress() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfigurationResponse) *SubResourceResponse { return v.PublicIPAddress }).(SubResourceResponsePtrOutput)
}

func (o NetworkInterfaceIpConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIpConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type NetworkInterfaceIpConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIpConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIpConfigurationResponse)(nil)).Elem()
}

func (o NetworkInterfaceIpConfigurationResponseArrayOutput) ToNetworkInterfaceIpConfigurationResponseArrayOutput() NetworkInterfaceIpConfigurationResponseArrayOutput {
	return o
}

func (o NetworkInterfaceIpConfigurationResponseArrayOutput) ToNetworkInterfaceIpConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceIpConfigurationResponseArrayOutput {
	return o
}

func (o NetworkInterfaceIpConfigurationResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceIpConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceIpConfigurationResponse {
		return vs[0].([]NetworkInterfaceIpConfigurationResponse)[vs[1].(int)]
	}).(NetworkInterfaceIpConfigurationResponseOutput)
}

type OutboundNatRule struct {
	AllocatedOutboundPorts   int           `pulumi:"allocatedOutboundPorts"`
	BackendAddressPool       SubResource   `pulumi:"backendAddressPool"`
	Etag                     *string       `pulumi:"etag"`
	FrontendIPConfigurations []SubResource `pulumi:"frontendIPConfigurations"`
	Id                       *string       `pulumi:"id"`
	Name                     *string       `pulumi:"name"`
	ProvisioningState        *string       `pulumi:"provisioningState"`
}





type OutboundNatRuleInput interface {
	pulumi.Input

	ToOutboundNatRuleOutput() OutboundNatRuleOutput
	ToOutboundNatRuleOutputWithContext(context.Context) OutboundNatRuleOutput
}

type OutboundNatRuleArgs struct {
	AllocatedOutboundPorts   pulumi.IntInput       `pulumi:"allocatedOutboundPorts"`
	BackendAddressPool       SubResourceInput      `pulumi:"backendAddressPool"`
	Etag                     pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfigurations SubResourceArrayInput `pulumi:"frontendIPConfigurations"`
	Id                       pulumi.StringPtrInput `pulumi:"id"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState        pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (OutboundNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundNatRule)(nil)).Elem()
}

func (i OutboundNatRuleArgs) ToOutboundNatRuleOutput() OutboundNatRuleOutput {
	return i.ToOutboundNatRuleOutputWithContext(context.Background())
}

func (i OutboundNatRuleArgs) ToOutboundNatRuleOutputWithContext(ctx context.Context) OutboundNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundNatRuleOutput)
}





type OutboundNatRuleArrayInput interface {
	pulumi.Input

	ToOutboundNatRuleArrayOutput() OutboundNatRuleArrayOutput
	ToOutboundNatRuleArrayOutputWithContext(context.Context) OutboundNatRuleArrayOutput
}

type OutboundNatRuleArray []OutboundNatRuleInput

func (OutboundNatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutboundNatRule)(nil)).Elem()
}

func (i OutboundNatRuleArray) ToOutboundNatRuleArrayOutput() OutboundNatRuleArrayOutput {
	return i.ToOutboundNatRuleArrayOutputWithContext(context.Background())
}

func (i OutboundNatRuleArray) ToOutboundNatRuleArrayOutputWithContext(ctx context.Context) OutboundNatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundNatRuleArrayOutput)
}

type OutboundNatRuleOutput struct{ *pulumi.OutputState }

func (OutboundNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundNatRule)(nil)).Elem()
}

func (o OutboundNatRuleOutput) ToOutboundNatRuleOutput() OutboundNatRuleOutput {
	return o
}

func (o OutboundNatRuleOutput) ToOutboundNatRuleOutputWithContext(ctx context.Context) OutboundNatRuleOutput {
	return o
}

func (o OutboundNatRuleOutput) AllocatedOutboundPorts() pulumi.IntOutput {
	return o.ApplyT(func(v OutboundNatRule) int { return v.AllocatedOutboundPorts }).(pulumi.IntOutput)
}

func (o OutboundNatRuleOutput) BackendAddressPool() SubResourceOutput {
	return o.ApplyT(func(v OutboundNatRule) SubResource { return v.BackendAddressPool }).(SubResourceOutput)
}

func (o OutboundNatRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRule) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleOutput) FrontendIPConfigurations() SubResourceArrayOutput {
	return o.ApplyT(func(v OutboundNatRule) []SubResource { return v.FrontendIPConfigurations }).(SubResourceArrayOutput)
}

func (o OutboundNatRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRule) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type OutboundNatRuleArrayOutput struct{ *pulumi.OutputState }

func (OutboundNatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutboundNatRule)(nil)).Elem()
}

func (o OutboundNatRuleArrayOutput) ToOutboundNatRuleArrayOutput() OutboundNatRuleArrayOutput {
	return o
}

func (o OutboundNatRuleArrayOutput) ToOutboundNatRuleArrayOutputWithContext(ctx context.Context) OutboundNatRuleArrayOutput {
	return o
}

func (o OutboundNatRuleArrayOutput) Index(i pulumi.IntInput) OutboundNatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutboundNatRule {
		return vs[0].([]OutboundNatRule)[vs[1].(int)]
	}).(OutboundNatRuleOutput)
}

type OutboundNatRuleResponse struct {
	AllocatedOutboundPorts   int                   `pulumi:"allocatedOutboundPorts"`
	BackendAddressPool       SubResourceResponse   `pulumi:"backendAddressPool"`
	Etag                     *string               `pulumi:"etag"`
	FrontendIPConfigurations []SubResourceResponse `pulumi:"frontendIPConfigurations"`
	Id                       *string               `pulumi:"id"`
	Name                     *string               `pulumi:"name"`
	ProvisioningState        *string               `pulumi:"provisioningState"`
}





type OutboundNatRuleResponseInput interface {
	pulumi.Input

	ToOutboundNatRuleResponseOutput() OutboundNatRuleResponseOutput
	ToOutboundNatRuleResponseOutputWithContext(context.Context) OutboundNatRuleResponseOutput
}

type OutboundNatRuleResponseArgs struct {
	AllocatedOutboundPorts   pulumi.IntInput               `pulumi:"allocatedOutboundPorts"`
	BackendAddressPool       SubResourceResponseInput      `pulumi:"backendAddressPool"`
	Etag                     pulumi.StringPtrInput         `pulumi:"etag"`
	FrontendIPConfigurations SubResourceResponseArrayInput `pulumi:"frontendIPConfigurations"`
	Id                       pulumi.StringPtrInput         `pulumi:"id"`
	Name                     pulumi.StringPtrInput         `pulumi:"name"`
	ProvisioningState        pulumi.StringPtrInput         `pulumi:"provisioningState"`
}

func (OutboundNatRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundNatRuleResponse)(nil)).Elem()
}

func (i OutboundNatRuleResponseArgs) ToOutboundNatRuleResponseOutput() OutboundNatRuleResponseOutput {
	return i.ToOutboundNatRuleResponseOutputWithContext(context.Background())
}

func (i OutboundNatRuleResponseArgs) ToOutboundNatRuleResponseOutputWithContext(ctx context.Context) OutboundNatRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundNatRuleResponseOutput)
}





type OutboundNatRuleResponseArrayInput interface {
	pulumi.Input

	ToOutboundNatRuleResponseArrayOutput() OutboundNatRuleResponseArrayOutput
	ToOutboundNatRuleResponseArrayOutputWithContext(context.Context) OutboundNatRuleResponseArrayOutput
}

type OutboundNatRuleResponseArray []OutboundNatRuleResponseInput

func (OutboundNatRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutboundNatRuleResponse)(nil)).Elem()
}

func (i OutboundNatRuleResponseArray) ToOutboundNatRuleResponseArrayOutput() OutboundNatRuleResponseArrayOutput {
	return i.ToOutboundNatRuleResponseArrayOutputWithContext(context.Background())
}

func (i OutboundNatRuleResponseArray) ToOutboundNatRuleResponseArrayOutputWithContext(ctx context.Context) OutboundNatRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundNatRuleResponseArrayOutput)
}

type OutboundNatRuleResponseOutput struct{ *pulumi.OutputState }

func (OutboundNatRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundNatRuleResponse)(nil)).Elem()
}

func (o OutboundNatRuleResponseOutput) ToOutboundNatRuleResponseOutput() OutboundNatRuleResponseOutput {
	return o
}

func (o OutboundNatRuleResponseOutput) ToOutboundNatRuleResponseOutputWithContext(ctx context.Context) OutboundNatRuleResponseOutput {
	return o
}

func (o OutboundNatRuleResponseOutput) AllocatedOutboundPorts() pulumi.IntOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) int { return v.AllocatedOutboundPorts }).(pulumi.IntOutput)
}

func (o OutboundNatRuleResponseOutput) BackendAddressPool() SubResourceResponseOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponseOutput)
}

func (o OutboundNatRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleResponseOutput) FrontendIPConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) []SubResourceResponse { return v.FrontendIPConfigurations }).(SubResourceResponseArrayOutput)
}

func (o OutboundNatRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type OutboundNatRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (OutboundNatRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutboundNatRuleResponse)(nil)).Elem()
}

func (o OutboundNatRuleResponseArrayOutput) ToOutboundNatRuleResponseArrayOutput() OutboundNatRuleResponseArrayOutput {
	return o
}

func (o OutboundNatRuleResponseArrayOutput) ToOutboundNatRuleResponseArrayOutputWithContext(ctx context.Context) OutboundNatRuleResponseArrayOutput {
	return o
}

func (o OutboundNatRuleResponseArrayOutput) Index(i pulumi.IntInput) OutboundNatRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutboundNatRuleResponse {
		return vs[0].([]OutboundNatRuleResponse)[vs[1].(int)]
	}).(OutboundNatRuleResponseOutput)
}

type Probe struct {
	Etag               *string       `pulumi:"etag"`
	Id                 *string       `pulumi:"id"`
	IntervalInSeconds  *int          `pulumi:"intervalInSeconds"`
	LoadBalancingRules []SubResource `pulumi:"loadBalancingRules"`
	Name               *string       `pulumi:"name"`
	NumberOfProbes     *int          `pulumi:"numberOfProbes"`
	Port               int           `pulumi:"port"`
	Protocol           string        `pulumi:"protocol"`
	ProvisioningState  *string       `pulumi:"provisioningState"`
	RequestPath        *string       `pulumi:"requestPath"`
}





type ProbeInput interface {
	pulumi.Input

	ToProbeOutput() ProbeOutput
	ToProbeOutputWithContext(context.Context) ProbeOutput
}

type ProbeArgs struct {
	Etag               pulumi.StringPtrInput `pulumi:"etag"`
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	IntervalInSeconds  pulumi.IntPtrInput    `pulumi:"intervalInSeconds"`
	LoadBalancingRules SubResourceArrayInput `pulumi:"loadBalancingRules"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	NumberOfProbes     pulumi.IntPtrInput    `pulumi:"numberOfProbes"`
	Port               pulumi.IntInput       `pulumi:"port"`
	Protocol           pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState  pulumi.StringPtrInput `pulumi:"provisioningState"`
	RequestPath        pulumi.StringPtrInput `pulumi:"requestPath"`
}

func (ProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Probe)(nil)).Elem()
}

func (i ProbeArgs) ToProbeOutput() ProbeOutput {
	return i.ToProbeOutputWithContext(context.Background())
}

func (i ProbeArgs) ToProbeOutputWithContext(ctx context.Context) ProbeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProbeOutput)
}





type ProbeArrayInput interface {
	pulumi.Input

	ToProbeArrayOutput() ProbeArrayOutput
	ToProbeArrayOutputWithContext(context.Context) ProbeArrayOutput
}

type ProbeArray []ProbeInput

func (ProbeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Probe)(nil)).Elem()
}

func (i ProbeArray) ToProbeArrayOutput() ProbeArrayOutput {
	return i.ToProbeArrayOutputWithContext(context.Background())
}

func (i ProbeArray) ToProbeArrayOutputWithContext(ctx context.Context) ProbeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProbeArrayOutput)
}

type ProbeOutput struct{ *pulumi.OutputState }

func (ProbeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Probe)(nil)).Elem()
}

func (o ProbeOutput) ToProbeOutput() ProbeOutput {
	return o
}

func (o ProbeOutput) ToProbeOutputWithContext(ctx context.Context) ProbeOutput {
	return o
}

func (o ProbeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Probe) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ProbeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Probe) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ProbeOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Probe) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ProbeOutput) LoadBalancingRules() SubResourceArrayOutput {
	return o.ApplyT(func(v Probe) []SubResource { return v.LoadBalancingRules }).(SubResourceArrayOutput)
}

func (o ProbeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Probe) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ProbeOutput) NumberOfProbes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Probe) *int { return v.NumberOfProbes }).(pulumi.IntPtrOutput)
}

func (o ProbeOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v Probe) int { return v.Port }).(pulumi.IntOutput)
}

func (o ProbeOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v Probe) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o ProbeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Probe) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ProbeOutput) RequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Probe) *string { return v.RequestPath }).(pulumi.StringPtrOutput)
}

type ProbeArrayOutput struct{ *pulumi.OutputState }

func (ProbeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Probe)(nil)).Elem()
}

func (o ProbeArrayOutput) ToProbeArrayOutput() ProbeArrayOutput {
	return o
}

func (o ProbeArrayOutput) ToProbeArrayOutputWithContext(ctx context.Context) ProbeArrayOutput {
	return o
}

func (o ProbeArrayOutput) Index(i pulumi.IntInput) ProbeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Probe {
		return vs[0].([]Probe)[vs[1].(int)]
	}).(ProbeOutput)
}

type ProbeResponse struct {
	Etag               *string               `pulumi:"etag"`
	Id                 *string               `pulumi:"id"`
	IntervalInSeconds  *int                  `pulumi:"intervalInSeconds"`
	LoadBalancingRules []SubResourceResponse `pulumi:"loadBalancingRules"`
	Name               *string               `pulumi:"name"`
	NumberOfProbes     *int                  `pulumi:"numberOfProbes"`
	Port               int                   `pulumi:"port"`
	Protocol           string                `pulumi:"protocol"`
	ProvisioningState  *string               `pulumi:"provisioningState"`
	RequestPath        *string               `pulumi:"requestPath"`
}





type ProbeResponseInput interface {
	pulumi.Input

	ToProbeResponseOutput() ProbeResponseOutput
	ToProbeResponseOutputWithContext(context.Context) ProbeResponseOutput
}

type ProbeResponseArgs struct {
	Etag               pulumi.StringPtrInput         `pulumi:"etag"`
	Id                 pulumi.StringPtrInput         `pulumi:"id"`
	IntervalInSeconds  pulumi.IntPtrInput            `pulumi:"intervalInSeconds"`
	LoadBalancingRules SubResourceResponseArrayInput `pulumi:"loadBalancingRules"`
	Name               pulumi.StringPtrInput         `pulumi:"name"`
	NumberOfProbes     pulumi.IntPtrInput            `pulumi:"numberOfProbes"`
	Port               pulumi.IntInput               `pulumi:"port"`
	Protocol           pulumi.StringInput            `pulumi:"protocol"`
	ProvisioningState  pulumi.StringPtrInput         `pulumi:"provisioningState"`
	RequestPath        pulumi.StringPtrInput         `pulumi:"requestPath"`
}

func (ProbeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProbeResponse)(nil)).Elem()
}

func (i ProbeResponseArgs) ToProbeResponseOutput() ProbeResponseOutput {
	return i.ToProbeResponseOutputWithContext(context.Background())
}

func (i ProbeResponseArgs) ToProbeResponseOutputWithContext(ctx context.Context) ProbeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProbeResponseOutput)
}





type ProbeResponseArrayInput interface {
	pulumi.Input

	ToProbeResponseArrayOutput() ProbeResponseArrayOutput
	ToProbeResponseArrayOutputWithContext(context.Context) ProbeResponseArrayOutput
}

type ProbeResponseArray []ProbeResponseInput

func (ProbeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProbeResponse)(nil)).Elem()
}

func (i ProbeResponseArray) ToProbeResponseArrayOutput() ProbeResponseArrayOutput {
	return i.ToProbeResponseArrayOutputWithContext(context.Background())
}

func (i ProbeResponseArray) ToProbeResponseArrayOutputWithContext(ctx context.Context) ProbeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProbeResponseArrayOutput)
}

type ProbeResponseOutput struct{ *pulumi.OutputState }

func (ProbeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProbeResponse)(nil)).Elem()
}

func (o ProbeResponseOutput) ToProbeResponseOutput() ProbeResponseOutput {
	return o
}

func (o ProbeResponseOutput) ToProbeResponseOutputWithContext(ctx context.Context) ProbeResponseOutput {
	return o
}

func (o ProbeResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ProbeResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ProbeResponseOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ProbeResponseOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v ProbeResponse) []SubResourceResponse { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

func (o ProbeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ProbeResponseOutput) NumberOfProbes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *int { return v.NumberOfProbes }).(pulumi.IntPtrOutput)
}

func (o ProbeResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ProbeResponse) int { return v.Port }).(pulumi.IntOutput)
}

func (o ProbeResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v ProbeResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o ProbeResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ProbeResponseOutput) RequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *string { return v.RequestPath }).(pulumi.StringPtrOutput)
}

type ProbeResponseArrayOutput struct{ *pulumi.OutputState }

func (ProbeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProbeResponse)(nil)).Elem()
}

func (o ProbeResponseArrayOutput) ToProbeResponseArrayOutput() ProbeResponseArrayOutput {
	return o
}

func (o ProbeResponseArrayOutput) ToProbeResponseArrayOutputWithContext(ctx context.Context) ProbeResponseArrayOutput {
	return o
}

func (o ProbeResponseArrayOutput) Index(i pulumi.IntInput) ProbeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProbeResponse {
		return vs[0].([]ProbeResponse)[vs[1].(int)]
	}).(ProbeResponseOutput)
}

type PublicIpAddressDnsSettings struct {
	DomainNameLabel *string `pulumi:"domainNameLabel"`
	Fqdn            *string `pulumi:"fqdn"`
	ReverseFqdn     *string `pulumi:"reverseFqdn"`
}





type PublicIpAddressDnsSettingsInput interface {
	pulumi.Input

	ToPublicIpAddressDnsSettingsOutput() PublicIpAddressDnsSettingsOutput
	ToPublicIpAddressDnsSettingsOutputWithContext(context.Context) PublicIpAddressDnsSettingsOutput
}

type PublicIpAddressDnsSettingsArgs struct {
	DomainNameLabel pulumi.StringPtrInput `pulumi:"domainNameLabel"`
	Fqdn            pulumi.StringPtrInput `pulumi:"fqdn"`
	ReverseFqdn     pulumi.StringPtrInput `pulumi:"reverseFqdn"`
}

func (PublicIpAddressDnsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpAddressDnsSettings)(nil)).Elem()
}

func (i PublicIpAddressDnsSettingsArgs) ToPublicIpAddressDnsSettingsOutput() PublicIpAddressDnsSettingsOutput {
	return i.ToPublicIpAddressDnsSettingsOutputWithContext(context.Background())
}

func (i PublicIpAddressDnsSettingsArgs) ToPublicIpAddressDnsSettingsOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpAddressDnsSettingsOutput)
}

func (i PublicIpAddressDnsSettingsArgs) ToPublicIpAddressDnsSettingsPtrOutput() PublicIpAddressDnsSettingsPtrOutput {
	return i.ToPublicIpAddressDnsSettingsPtrOutputWithContext(context.Background())
}

func (i PublicIpAddressDnsSettingsArgs) ToPublicIpAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpAddressDnsSettingsOutput).ToPublicIpAddressDnsSettingsPtrOutputWithContext(ctx)
}









type PublicIpAddressDnsSettingsPtrInput interface {
	pulumi.Input

	ToPublicIpAddressDnsSettingsPtrOutput() PublicIpAddressDnsSettingsPtrOutput
	ToPublicIpAddressDnsSettingsPtrOutputWithContext(context.Context) PublicIpAddressDnsSettingsPtrOutput
}

type publicIpAddressDnsSettingsPtrType PublicIpAddressDnsSettingsArgs

func PublicIpAddressDnsSettingsPtr(v *PublicIpAddressDnsSettingsArgs) PublicIpAddressDnsSettingsPtrInput {
	return (*publicIpAddressDnsSettingsPtrType)(v)
}

func (*publicIpAddressDnsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpAddressDnsSettings)(nil)).Elem()
}

func (i *publicIpAddressDnsSettingsPtrType) ToPublicIpAddressDnsSettingsPtrOutput() PublicIpAddressDnsSettingsPtrOutput {
	return i.ToPublicIpAddressDnsSettingsPtrOutputWithContext(context.Background())
}

func (i *publicIpAddressDnsSettingsPtrType) ToPublicIpAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpAddressDnsSettingsPtrOutput)
}

type PublicIpAddressDnsSettingsOutput struct{ *pulumi.OutputState }

func (PublicIpAddressDnsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpAddressDnsSettings)(nil)).Elem()
}

func (o PublicIpAddressDnsSettingsOutput) ToPublicIpAddressDnsSettingsOutput() PublicIpAddressDnsSettingsOutput {
	return o
}

func (o PublicIpAddressDnsSettingsOutput) ToPublicIpAddressDnsSettingsOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsOutput {
	return o
}

func (o PublicIpAddressDnsSettingsOutput) ToPublicIpAddressDnsSettingsPtrOutput() PublicIpAddressDnsSettingsPtrOutput {
	return o.ToPublicIpAddressDnsSettingsPtrOutputWithContext(context.Background())
}

func (o PublicIpAddressDnsSettingsOutput) ToPublicIpAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIpAddressDnsSettings) *PublicIpAddressDnsSettings {
		return &v
	}).(PublicIpAddressDnsSettingsPtrOutput)
}

func (o PublicIpAddressDnsSettingsOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIpAddressDnsSettings) *string { return v.DomainNameLabel }).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressDnsSettingsOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIpAddressDnsSettings) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressDnsSettingsOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIpAddressDnsSettings) *string { return v.ReverseFqdn }).(pulumi.StringPtrOutput)
}

type PublicIpAddressDnsSettingsPtrOutput struct{ *pulumi.OutputState }

func (PublicIpAddressDnsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpAddressDnsSettings)(nil)).Elem()
}

func (o PublicIpAddressDnsSettingsPtrOutput) ToPublicIpAddressDnsSettingsPtrOutput() PublicIpAddressDnsSettingsPtrOutput {
	return o
}

func (o PublicIpAddressDnsSettingsPtrOutput) ToPublicIpAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsPtrOutput {
	return o
}

func (o PublicIpAddressDnsSettingsPtrOutput) Elem() PublicIpAddressDnsSettingsOutput {
	return o.ApplyT(func(v *PublicIpAddressDnsSettings) PublicIpAddressDnsSettings {
		if v != nil {
			return *v
		}
		var ret PublicIpAddressDnsSettings
		return ret
	}).(PublicIpAddressDnsSettingsOutput)
}

func (o PublicIpAddressDnsSettingsPtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpAddressDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressDnsSettingsPtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpAddressDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressDnsSettingsPtrOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpAddressDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReverseFqdn
	}).(pulumi.StringPtrOutput)
}

type PublicIpAddressDnsSettingsResponse struct {
	DomainNameLabel *string `pulumi:"domainNameLabel"`
	Fqdn            *string `pulumi:"fqdn"`
	ReverseFqdn     *string `pulumi:"reverseFqdn"`
}





type PublicIpAddressDnsSettingsResponseInput interface {
	pulumi.Input

	ToPublicIpAddressDnsSettingsResponseOutput() PublicIpAddressDnsSettingsResponseOutput
	ToPublicIpAddressDnsSettingsResponseOutputWithContext(context.Context) PublicIpAddressDnsSettingsResponseOutput
}

type PublicIpAddressDnsSettingsResponseArgs struct {
	DomainNameLabel pulumi.StringPtrInput `pulumi:"domainNameLabel"`
	Fqdn            pulumi.StringPtrInput `pulumi:"fqdn"`
	ReverseFqdn     pulumi.StringPtrInput `pulumi:"reverseFqdn"`
}

func (PublicIpAddressDnsSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpAddressDnsSettingsResponse)(nil)).Elem()
}

func (i PublicIpAddressDnsSettingsResponseArgs) ToPublicIpAddressDnsSettingsResponseOutput() PublicIpAddressDnsSettingsResponseOutput {
	return i.ToPublicIpAddressDnsSettingsResponseOutputWithContext(context.Background())
}

func (i PublicIpAddressDnsSettingsResponseArgs) ToPublicIpAddressDnsSettingsResponseOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpAddressDnsSettingsResponseOutput)
}

func (i PublicIpAddressDnsSettingsResponseArgs) ToPublicIpAddressDnsSettingsResponsePtrOutput() PublicIpAddressDnsSettingsResponsePtrOutput {
	return i.ToPublicIpAddressDnsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i PublicIpAddressDnsSettingsResponseArgs) ToPublicIpAddressDnsSettingsResponsePtrOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpAddressDnsSettingsResponseOutput).ToPublicIpAddressDnsSettingsResponsePtrOutputWithContext(ctx)
}









type PublicIpAddressDnsSettingsResponsePtrInput interface {
	pulumi.Input

	ToPublicIpAddressDnsSettingsResponsePtrOutput() PublicIpAddressDnsSettingsResponsePtrOutput
	ToPublicIpAddressDnsSettingsResponsePtrOutputWithContext(context.Context) PublicIpAddressDnsSettingsResponsePtrOutput
}

type publicIpAddressDnsSettingsResponsePtrType PublicIpAddressDnsSettingsResponseArgs

func PublicIpAddressDnsSettingsResponsePtr(v *PublicIpAddressDnsSettingsResponseArgs) PublicIpAddressDnsSettingsResponsePtrInput {
	return (*publicIpAddressDnsSettingsResponsePtrType)(v)
}

func (*publicIpAddressDnsSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpAddressDnsSettingsResponse)(nil)).Elem()
}

func (i *publicIpAddressDnsSettingsResponsePtrType) ToPublicIpAddressDnsSettingsResponsePtrOutput() PublicIpAddressDnsSettingsResponsePtrOutput {
	return i.ToPublicIpAddressDnsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *publicIpAddressDnsSettingsResponsePtrType) ToPublicIpAddressDnsSettingsResponsePtrOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpAddressDnsSettingsResponsePtrOutput)
}

type PublicIpAddressDnsSettingsResponseOutput struct{ *pulumi.OutputState }

func (PublicIpAddressDnsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpAddressDnsSettingsResponse)(nil)).Elem()
}

func (o PublicIpAddressDnsSettingsResponseOutput) ToPublicIpAddressDnsSettingsResponseOutput() PublicIpAddressDnsSettingsResponseOutput {
	return o
}

func (o PublicIpAddressDnsSettingsResponseOutput) ToPublicIpAddressDnsSettingsResponseOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsResponseOutput {
	return o
}

func (o PublicIpAddressDnsSettingsResponseOutput) ToPublicIpAddressDnsSettingsResponsePtrOutput() PublicIpAddressDnsSettingsResponsePtrOutput {
	return o.ToPublicIpAddressDnsSettingsResponsePtrOutputWithContext(context.Background())
}

func (o PublicIpAddressDnsSettingsResponseOutput) ToPublicIpAddressDnsSettingsResponsePtrOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIpAddressDnsSettingsResponse) *PublicIpAddressDnsSettingsResponse {
		return &v
	}).(PublicIpAddressDnsSettingsResponsePtrOutput)
}

func (o PublicIpAddressDnsSettingsResponseOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIpAddressDnsSettingsResponse) *string { return v.DomainNameLabel }).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressDnsSettingsResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIpAddressDnsSettingsResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressDnsSettingsResponseOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIpAddressDnsSettingsResponse) *string { return v.ReverseFqdn }).(pulumi.StringPtrOutput)
}

type PublicIpAddressDnsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (PublicIpAddressDnsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpAddressDnsSettingsResponse)(nil)).Elem()
}

func (o PublicIpAddressDnsSettingsResponsePtrOutput) ToPublicIpAddressDnsSettingsResponsePtrOutput() PublicIpAddressDnsSettingsResponsePtrOutput {
	return o
}

func (o PublicIpAddressDnsSettingsResponsePtrOutput) ToPublicIpAddressDnsSettingsResponsePtrOutputWithContext(ctx context.Context) PublicIpAddressDnsSettingsResponsePtrOutput {
	return o
}

func (o PublicIpAddressDnsSettingsResponsePtrOutput) Elem() PublicIpAddressDnsSettingsResponseOutput {
	return o.ApplyT(func(v *PublicIpAddressDnsSettingsResponse) PublicIpAddressDnsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret PublicIpAddressDnsSettingsResponse
		return ret
	}).(PublicIpAddressDnsSettingsResponseOutput)
}

func (o PublicIpAddressDnsSettingsResponsePtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpAddressDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressDnsSettingsResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpAddressDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressDnsSettingsResponsePtrOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpAddressDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReverseFqdn
	}).(pulumi.StringPtrOutput)
}

type RouteType struct {
	AddressPrefix     *string `pulumi:"addressPrefix"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	NextHopIpAddress  *string `pulumi:"nextHopIpAddress"`
	NextHopType       string  `pulumi:"nextHopType"`
	ProvisioningState *string `pulumi:"provisioningState"`
}





type RouteTypeInput interface {
	pulumi.Input

	ToRouteTypeOutput() RouteTypeOutput
	ToRouteTypeOutputWithContext(context.Context) RouteTypeOutput
}

type RouteTypeArgs struct {
	AddressPrefix     pulumi.StringPtrInput `pulumi:"addressPrefix"`
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	NextHopIpAddress  pulumi.StringPtrInput `pulumi:"nextHopIpAddress"`
	NextHopType       pulumi.StringInput    `pulumi:"nextHopType"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (RouteTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteType)(nil)).Elem()
}

func (i RouteTypeArgs) ToRouteTypeOutput() RouteTypeOutput {
	return i.ToRouteTypeOutputWithContext(context.Background())
}

func (i RouteTypeArgs) ToRouteTypeOutputWithContext(ctx context.Context) RouteTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTypeOutput)
}





type RouteTypeArrayInput interface {
	pulumi.Input

	ToRouteTypeArrayOutput() RouteTypeArrayOutput
	ToRouteTypeArrayOutputWithContext(context.Context) RouteTypeArrayOutput
}

type RouteTypeArray []RouteTypeInput

func (RouteTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteType)(nil)).Elem()
}

func (i RouteTypeArray) ToRouteTypeArrayOutput() RouteTypeArrayOutput {
	return i.ToRouteTypeArrayOutputWithContext(context.Background())
}

func (i RouteTypeArray) ToRouteTypeArrayOutputWithContext(ctx context.Context) RouteTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTypeArrayOutput)
}

type RouteTypeOutput struct{ *pulumi.OutputState }

func (RouteTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteType)(nil)).Elem()
}

func (o RouteTypeOutput) ToRouteTypeOutput() RouteTypeOutput {
	return o
}

func (o RouteTypeOutput) ToRouteTypeOutputWithContext(ctx context.Context) RouteTypeOutput {
	return o
}

func (o RouteTypeOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o RouteTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RouteTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RouteTypeOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

func (o RouteTypeOutput) NextHopType() pulumi.StringOutput {
	return o.ApplyT(func(v RouteType) string { return v.NextHopType }).(pulumi.StringOutput)
}

func (o RouteTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type RouteTypeArrayOutput struct{ *pulumi.OutputState }

func (RouteTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteType)(nil)).Elem()
}

func (o RouteTypeArrayOutput) ToRouteTypeArrayOutput() RouteTypeArrayOutput {
	return o
}

func (o RouteTypeArrayOutput) ToRouteTypeArrayOutputWithContext(ctx context.Context) RouteTypeArrayOutput {
	return o
}

func (o RouteTypeArrayOutput) Index(i pulumi.IntInput) RouteTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteType {
		return vs[0].([]RouteType)[vs[1].(int)]
	}).(RouteTypeOutput)
}

type RouteResponse struct {
	AddressPrefix     *string `pulumi:"addressPrefix"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	NextHopIpAddress  *string `pulumi:"nextHopIpAddress"`
	NextHopType       string  `pulumi:"nextHopType"`
	ProvisioningState *string `pulumi:"provisioningState"`
}





type RouteResponseInput interface {
	pulumi.Input

	ToRouteResponseOutput() RouteResponseOutput
	ToRouteResponseOutputWithContext(context.Context) RouteResponseOutput
}

type RouteResponseArgs struct {
	AddressPrefix     pulumi.StringPtrInput `pulumi:"addressPrefix"`
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	NextHopIpAddress  pulumi.StringPtrInput `pulumi:"nextHopIpAddress"`
	NextHopType       pulumi.StringInput    `pulumi:"nextHopType"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (RouteResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteResponse)(nil)).Elem()
}

func (i RouteResponseArgs) ToRouteResponseOutput() RouteResponseOutput {
	return i.ToRouteResponseOutputWithContext(context.Background())
}

func (i RouteResponseArgs) ToRouteResponseOutputWithContext(ctx context.Context) RouteResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteResponseOutput)
}





type RouteResponseArrayInput interface {
	pulumi.Input

	ToRouteResponseArrayOutput() RouteResponseArrayOutput
	ToRouteResponseArrayOutputWithContext(context.Context) RouteResponseArrayOutput
}

type RouteResponseArray []RouteResponseInput

func (RouteResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteResponse)(nil)).Elem()
}

func (i RouteResponseArray) ToRouteResponseArrayOutput() RouteResponseArrayOutput {
	return i.ToRouteResponseArrayOutputWithContext(context.Background())
}

func (i RouteResponseArray) ToRouteResponseArrayOutputWithContext(ctx context.Context) RouteResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteResponseArrayOutput)
}

type RouteResponseOutput struct{ *pulumi.OutputState }

func (RouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteResponse)(nil)).Elem()
}

func (o RouteResponseOutput) ToRouteResponseOutput() RouteResponseOutput {
	return o
}

func (o RouteResponseOutput) ToRouteResponseOutputWithContext(ctx context.Context) RouteResponseOutput {
	return o
}

func (o RouteResponseOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o RouteResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RouteResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RouteResponseOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

func (o RouteResponseOutput) NextHopType() pulumi.StringOutput {
	return o.ApplyT(func(v RouteResponse) string { return v.NextHopType }).(pulumi.StringOutput)
}

func (o RouteResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type RouteResponseArrayOutput struct{ *pulumi.OutputState }

func (RouteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteResponse)(nil)).Elem()
}

func (o RouteResponseArrayOutput) ToRouteResponseArrayOutput() RouteResponseArrayOutput {
	return o
}

func (o RouteResponseArrayOutput) ToRouteResponseArrayOutputWithContext(ctx context.Context) RouteResponseArrayOutput {
	return o
}

func (o RouteResponseArrayOutput) Index(i pulumi.IntInput) RouteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteResponse {
		return vs[0].([]RouteResponse)[vs[1].(int)]
	}).(RouteResponseOutput)
}

type SecurityRuleType struct {
	Access                   string  `pulumi:"access"`
	Description              *string `pulumi:"description"`
	DestinationAddressPrefix string  `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     *string `pulumi:"destinationPortRange"`
	Direction                string  `pulumi:"direction"`
	Etag                     *string `pulumi:"etag"`
	Id                       *string `pulumi:"id"`
	Name                     *string `pulumi:"name"`
	Priority                 *int    `pulumi:"priority"`
	Protocol                 string  `pulumi:"protocol"`
	ProvisioningState        *string `pulumi:"provisioningState"`
	SourceAddressPrefix      string  `pulumi:"sourceAddressPrefix"`
	SourcePortRange          *string `pulumi:"sourcePortRange"`
}





type SecurityRuleTypeInput interface {
	pulumi.Input

	ToSecurityRuleTypeOutput() SecurityRuleTypeOutput
	ToSecurityRuleTypeOutputWithContext(context.Context) SecurityRuleTypeOutput
}

type SecurityRuleTypeArgs struct {
	Access                   pulumi.StringInput    `pulumi:"access"`
	Description              pulumi.StringPtrInput `pulumi:"description"`
	DestinationAddressPrefix pulumi.StringInput    `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     pulumi.StringPtrInput `pulumi:"destinationPortRange"`
	Direction                pulumi.StringInput    `pulumi:"direction"`
	Etag                     pulumi.StringPtrInput `pulumi:"etag"`
	Id                       pulumi.StringPtrInput `pulumi:"id"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	Priority                 pulumi.IntPtrInput    `pulumi:"priority"`
	Protocol                 pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState        pulumi.StringPtrInput `pulumi:"provisioningState"`
	SourceAddressPrefix      pulumi.StringInput    `pulumi:"sourceAddressPrefix"`
	SourcePortRange          pulumi.StringPtrInput `pulumi:"sourcePortRange"`
}

func (SecurityRuleTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleType)(nil)).Elem()
}

func (i SecurityRuleTypeArgs) ToSecurityRuleTypeOutput() SecurityRuleTypeOutput {
	return i.ToSecurityRuleTypeOutputWithContext(context.Background())
}

func (i SecurityRuleTypeArgs) ToSecurityRuleTypeOutputWithContext(ctx context.Context) SecurityRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityRuleTypeOutput)
}





type SecurityRuleTypeArrayInput interface {
	pulumi.Input

	ToSecurityRuleTypeArrayOutput() SecurityRuleTypeArrayOutput
	ToSecurityRuleTypeArrayOutputWithContext(context.Context) SecurityRuleTypeArrayOutput
}

type SecurityRuleTypeArray []SecurityRuleTypeInput

func (SecurityRuleTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityRuleType)(nil)).Elem()
}

func (i SecurityRuleTypeArray) ToSecurityRuleTypeArrayOutput() SecurityRuleTypeArrayOutput {
	return i.ToSecurityRuleTypeArrayOutputWithContext(context.Background())
}

func (i SecurityRuleTypeArray) ToSecurityRuleTypeArrayOutputWithContext(ctx context.Context) SecurityRuleTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityRuleTypeArrayOutput)
}

type SecurityRuleTypeOutput struct{ *pulumi.OutputState }

func (SecurityRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleType)(nil)).Elem()
}

func (o SecurityRuleTypeOutput) ToSecurityRuleTypeOutput() SecurityRuleTypeOutput {
	return o
}

func (o SecurityRuleTypeOutput) ToSecurityRuleTypeOutputWithContext(ctx context.Context) SecurityRuleTypeOutput {
	return o
}

func (o SecurityRuleTypeOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.Access }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) DestinationAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.DestinationAddressPrefix }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.Direction }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SecurityRuleTypeOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) SourceAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.SourceAddressPrefix }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

type SecurityRuleTypeArrayOutput struct{ *pulumi.OutputState }

func (SecurityRuleTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityRuleType)(nil)).Elem()
}

func (o SecurityRuleTypeArrayOutput) ToSecurityRuleTypeArrayOutput() SecurityRuleTypeArrayOutput {
	return o
}

func (o SecurityRuleTypeArrayOutput) ToSecurityRuleTypeArrayOutputWithContext(ctx context.Context) SecurityRuleTypeArrayOutput {
	return o
}

func (o SecurityRuleTypeArrayOutput) Index(i pulumi.IntInput) SecurityRuleTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityRuleType {
		return vs[0].([]SecurityRuleType)[vs[1].(int)]
	}).(SecurityRuleTypeOutput)
}

type SecurityRuleResponse struct {
	Access                   string  `pulumi:"access"`
	Description              *string `pulumi:"description"`
	DestinationAddressPrefix string  `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     *string `pulumi:"destinationPortRange"`
	Direction                string  `pulumi:"direction"`
	Etag                     *string `pulumi:"etag"`
	Id                       *string `pulumi:"id"`
	Name                     *string `pulumi:"name"`
	Priority                 *int    `pulumi:"priority"`
	Protocol                 string  `pulumi:"protocol"`
	ProvisioningState        *string `pulumi:"provisioningState"`
	SourceAddressPrefix      string  `pulumi:"sourceAddressPrefix"`
	SourcePortRange          *string `pulumi:"sourcePortRange"`
}





type SecurityRuleResponseInput interface {
	pulumi.Input

	ToSecurityRuleResponseOutput() SecurityRuleResponseOutput
	ToSecurityRuleResponseOutputWithContext(context.Context) SecurityRuleResponseOutput
}

type SecurityRuleResponseArgs struct {
	Access                   pulumi.StringInput    `pulumi:"access"`
	Description              pulumi.StringPtrInput `pulumi:"description"`
	DestinationAddressPrefix pulumi.StringInput    `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     pulumi.StringPtrInput `pulumi:"destinationPortRange"`
	Direction                pulumi.StringInput    `pulumi:"direction"`
	Etag                     pulumi.StringPtrInput `pulumi:"etag"`
	Id                       pulumi.StringPtrInput `pulumi:"id"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	Priority                 pulumi.IntPtrInput    `pulumi:"priority"`
	Protocol                 pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState        pulumi.StringPtrInput `pulumi:"provisioningState"`
	SourceAddressPrefix      pulumi.StringInput    `pulumi:"sourceAddressPrefix"`
	SourcePortRange          pulumi.StringPtrInput `pulumi:"sourcePortRange"`
}

func (SecurityRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleResponse)(nil)).Elem()
}

func (i SecurityRuleResponseArgs) ToSecurityRuleResponseOutput() SecurityRuleResponseOutput {
	return i.ToSecurityRuleResponseOutputWithContext(context.Background())
}

func (i SecurityRuleResponseArgs) ToSecurityRuleResponseOutputWithContext(ctx context.Context) SecurityRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityRuleResponseOutput)
}





type SecurityRuleResponseArrayInput interface {
	pulumi.Input

	ToSecurityRuleResponseArrayOutput() SecurityRuleResponseArrayOutput
	ToSecurityRuleResponseArrayOutputWithContext(context.Context) SecurityRuleResponseArrayOutput
}

type SecurityRuleResponseArray []SecurityRuleResponseInput

func (SecurityRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityRuleResponse)(nil)).Elem()
}

func (i SecurityRuleResponseArray) ToSecurityRuleResponseArrayOutput() SecurityRuleResponseArrayOutput {
	return i.ToSecurityRuleResponseArrayOutputWithContext(context.Background())
}

func (i SecurityRuleResponseArray) ToSecurityRuleResponseArrayOutputWithContext(ctx context.Context) SecurityRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityRuleResponseArrayOutput)
}

type SecurityRuleResponseOutput struct{ *pulumi.OutputState }

func (SecurityRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleResponse)(nil)).Elem()
}

func (o SecurityRuleResponseOutput) ToSecurityRuleResponseOutput() SecurityRuleResponseOutput {
	return o
}

func (o SecurityRuleResponseOutput) ToSecurityRuleResponseOutputWithContext(ctx context.Context) SecurityRuleResponseOutput {
	return o
}

func (o SecurityRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) DestinationAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.DestinationAddressPrefix }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SecurityRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) SourceAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.SourceAddressPrefix }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

type SecurityRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (SecurityRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityRuleResponse)(nil)).Elem()
}

func (o SecurityRuleResponseArrayOutput) ToSecurityRuleResponseArrayOutput() SecurityRuleResponseArrayOutput {
	return o
}

func (o SecurityRuleResponseArrayOutput) ToSecurityRuleResponseArrayOutputWithContext(ctx context.Context) SecurityRuleResponseArrayOutput {
	return o
}

func (o SecurityRuleResponseArrayOutput) Index(i pulumi.IntInput) SecurityRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityRuleResponse {
		return vs[0].([]SecurityRuleResponse)[vs[1].(int)]
	}).(SecurityRuleResponseOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}





type SubResourceArrayInput interface {
	pulumi.Input

	ToSubResourceArrayOutput() SubResourceArrayOutput
	ToSubResourceArrayOutputWithContext(context.Context) SubResourceArrayOutput
}

type SubResourceArray []SubResourceInput

func (SubResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (i SubResourceArray) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return i.ToSubResourceArrayOutputWithContext(context.Background())
}

func (i SubResourceArray) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceArrayOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceArrayOutput struct{ *pulumi.OutputState }

func (SubResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) Index(i pulumi.IntInput) SubResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResource {
		return vs[0].([]SubResource)[vs[1].(int)]
	}).(SubResourceOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}





type SubResourceResponseInput interface {
	pulumi.Input

	ToSubResourceResponseOutput() SubResourceResponseOutput
	ToSubResourceResponseOutputWithContext(context.Context) SubResourceResponseOutput
}

type SubResourceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return i.ToSubResourceResponseOutputWithContext(context.Background())
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseOutput)
}

func (i SubResourceResponseArgs) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return i.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (i SubResourceResponseArgs) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseOutput).ToSubResourceResponsePtrOutputWithContext(ctx)
}









type SubResourceResponsePtrInput interface {
	pulumi.Input

	ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput
	ToSubResourceResponsePtrOutputWithContext(context.Context) SubResourceResponsePtrOutput
}

type subResourceResponsePtrType SubResourceResponseArgs

func SubResourceResponsePtr(v *SubResourceResponseArgs) SubResourceResponsePtrInput {
	return (*subResourceResponsePtrType)(v)
}

func (*subResourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (i *subResourceResponsePtrType) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return i.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (i *subResourceResponsePtrType) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponsePtrOutput)
}





type SubResourceResponseArrayInput interface {
	pulumi.Input

	ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput
	ToSubResourceResponseArrayOutputWithContext(context.Context) SubResourceResponseArrayOutput
}

type SubResourceResponseArray []SubResourceResponseInput

func (SubResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (i SubResourceResponseArray) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return i.ToSubResourceResponseArrayOutputWithContext(context.Background())
}

func (i SubResourceResponseArray) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseArrayOutput)
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (o SubResourceResponseOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResourceResponse) *SubResourceResponse {
		return &v
	}).(SubResourceResponsePtrOutput)
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
}

type SubnetType struct {
	AddressPrefix        string        `pulumi:"addressPrefix"`
	Etag                 *string       `pulumi:"etag"`
	Id                   *string       `pulumi:"id"`
	IpConfigurations     []SubResource `pulumi:"ipConfigurations"`
	Name                 *string       `pulumi:"name"`
	NetworkSecurityGroup *SubResource  `pulumi:"networkSecurityGroup"`
	ProvisioningState    *string       `pulumi:"provisioningState"`
	RouteTable           *SubResource  `pulumi:"routeTable"`
}





type SubnetTypeInput interface {
	pulumi.Input

	ToSubnetTypeOutput() SubnetTypeOutput
	ToSubnetTypeOutputWithContext(context.Context) SubnetTypeOutput
}

type SubnetTypeArgs struct {
	AddressPrefix        pulumi.StringInput    `pulumi:"addressPrefix"`
	Etag                 pulumi.StringPtrInput `pulumi:"etag"`
	Id                   pulumi.StringPtrInput `pulumi:"id"`
	IpConfigurations     SubResourceArrayInput `pulumi:"ipConfigurations"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
	NetworkSecurityGroup SubResourcePtrInput   `pulumi:"networkSecurityGroup"`
	ProvisioningState    pulumi.StringPtrInput `pulumi:"provisioningState"`
	RouteTable           SubResourcePtrInput   `pulumi:"routeTable"`
}

func (SubnetTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetType)(nil)).Elem()
}

func (i SubnetTypeArgs) ToSubnetTypeOutput() SubnetTypeOutput {
	return i.ToSubnetTypeOutputWithContext(context.Background())
}

func (i SubnetTypeArgs) ToSubnetTypeOutputWithContext(ctx context.Context) SubnetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypeOutput)
}





type SubnetTypeArrayInput interface {
	pulumi.Input

	ToSubnetTypeArrayOutput() SubnetTypeArrayOutput
	ToSubnetTypeArrayOutputWithContext(context.Context) SubnetTypeArrayOutput
}

type SubnetTypeArray []SubnetTypeInput

func (SubnetTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetType)(nil)).Elem()
}

func (i SubnetTypeArray) ToSubnetTypeArrayOutput() SubnetTypeArrayOutput {
	return i.ToSubnetTypeArrayOutputWithContext(context.Background())
}

func (i SubnetTypeArray) ToSubnetTypeArrayOutputWithContext(ctx context.Context) SubnetTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypeArrayOutput)
}

type SubnetTypeOutput struct{ *pulumi.OutputState }

func (SubnetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetType)(nil)).Elem()
}

func (o SubnetTypeOutput) ToSubnetTypeOutput() SubnetTypeOutput {
	return o
}

func (o SubnetTypeOutput) ToSubnetTypeOutputWithContext(ctx context.Context) SubnetTypeOutput {
	return o
}

func (o SubnetTypeOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetType) string { return v.AddressPrefix }).(pulumi.StringOutput)
}

func (o SubnetTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) IpConfigurations() SubResourceArrayOutput {
	return o.ApplyT(func(v SubnetType) []SubResource { return v.IpConfigurations }).(SubResourceArrayOutput)
}

func (o SubnetTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) NetworkSecurityGroup() SubResourcePtrOutput {
	return o.ApplyT(func(v SubnetType) *SubResource { return v.NetworkSecurityGroup }).(SubResourcePtrOutput)
}

func (o SubnetTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) RouteTable() SubResourcePtrOutput {
	return o.ApplyT(func(v SubnetType) *SubResource { return v.RouteTable }).(SubResourcePtrOutput)
}

type SubnetTypeArrayOutput struct{ *pulumi.OutputState }

func (SubnetTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetType)(nil)).Elem()
}

func (o SubnetTypeArrayOutput) ToSubnetTypeArrayOutput() SubnetTypeArrayOutput {
	return o
}

func (o SubnetTypeArrayOutput) ToSubnetTypeArrayOutputWithContext(ctx context.Context) SubnetTypeArrayOutput {
	return o
}

func (o SubnetTypeArrayOutput) Index(i pulumi.IntInput) SubnetTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetType {
		return vs[0].([]SubnetType)[vs[1].(int)]
	}).(SubnetTypeOutput)
}

type SubnetResponse struct {
	AddressPrefix        string                `pulumi:"addressPrefix"`
	Etag                 *string               `pulumi:"etag"`
	Id                   *string               `pulumi:"id"`
	IpConfigurations     []SubResourceResponse `pulumi:"ipConfigurations"`
	Name                 *string               `pulumi:"name"`
	NetworkSecurityGroup *SubResourceResponse  `pulumi:"networkSecurityGroup"`
	ProvisioningState    *string               `pulumi:"provisioningState"`
	RouteTable           *SubResourceResponse  `pulumi:"routeTable"`
}





type SubnetResponseInput interface {
	pulumi.Input

	ToSubnetResponseOutput() SubnetResponseOutput
	ToSubnetResponseOutputWithContext(context.Context) SubnetResponseOutput
}

type SubnetResponseArgs struct {
	AddressPrefix        pulumi.StringInput            `pulumi:"addressPrefix"`
	Etag                 pulumi.StringPtrInput         `pulumi:"etag"`
	Id                   pulumi.StringPtrInput         `pulumi:"id"`
	IpConfigurations     SubResourceResponseArrayInput `pulumi:"ipConfigurations"`
	Name                 pulumi.StringPtrInput         `pulumi:"name"`
	NetworkSecurityGroup SubResourceResponsePtrInput   `pulumi:"networkSecurityGroup"`
	ProvisioningState    pulumi.StringPtrInput         `pulumi:"provisioningState"`
	RouteTable           SubResourceResponsePtrInput   `pulumi:"routeTable"`
}

func (SubnetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (i SubnetResponseArgs) ToSubnetResponseOutput() SubnetResponseOutput {
	return i.ToSubnetResponseOutputWithContext(context.Background())
}

func (i SubnetResponseArgs) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResponseOutput)
}





type SubnetResponseArrayInput interface {
	pulumi.Input

	ToSubnetResponseArrayOutput() SubnetResponseArrayOutput
	ToSubnetResponseArrayOutputWithContext(context.Context) SubnetResponseArrayOutput
}

type SubnetResponseArray []SubnetResponseInput

func (SubnetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResponse)(nil)).Elem()
}

func (i SubnetResponseArray) ToSubnetResponseArrayOutput() SubnetResponseArrayOutput {
	return i.ToSubnetResponseArrayOutputWithContext(context.Background())
}

func (i SubnetResponseArray) ToSubnetResponseArrayOutputWithContext(ctx context.Context) SubnetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResponseArrayOutput)
}

type SubnetResponseOutput struct{ *pulumi.OutputState }

func (SubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseOutput) ToSubnetResponseOutput() SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetResponse) string { return v.AddressPrefix }).(pulumi.StringOutput)
}

func (o SubnetResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) IpConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []SubResourceResponse { return v.IpConfigurations }).(SubResourceResponseArrayOutput)
}

func (o SubnetResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) NetworkSecurityGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v SubnetResponse) *SubResourceResponse { return v.NetworkSecurityGroup }).(SubResourceResponsePtrOutput)
}

func (o SubnetResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) RouteTable() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v SubnetResponse) *SubResourceResponse { return v.RouteTable }).(SubResourceResponsePtrOutput)
}

type SubnetResponseArrayOutput struct{ *pulumi.OutputState }

func (SubnetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutput() SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutputWithContext(ctx context.Context) SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) Index(i pulumi.IntInput) SubnetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetResponse {
		return vs[0].([]SubnetResponse)[vs[1].(int)]
	}).(SubnetResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressSpaceInput)(nil)).Elem(), AddressSpaceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressSpacePtrInput)(nil)).Elem(), AddressSpaceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressSpaceResponseInput)(nil)).Elem(), AddressSpaceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressSpaceResponsePtrInput)(nil)).Elem(), AddressSpaceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendAddressInput)(nil)).Elem(), ApplicationGatewayBackendAddressArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendAddressArrayInput)(nil)).Elem(), ApplicationGatewayBackendAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendAddressPoolInput)(nil)).Elem(), ApplicationGatewayBackendAddressPoolArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendAddressPoolArrayInput)(nil)).Elem(), ApplicationGatewayBackendAddressPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendAddressPoolResponseInput)(nil)).Elem(), ApplicationGatewayBackendAddressPoolResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendAddressPoolResponseArrayInput)(nil)).Elem(), ApplicationGatewayBackendAddressPoolResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendAddressResponseInput)(nil)).Elem(), ApplicationGatewayBackendAddressResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendAddressResponseArrayInput)(nil)).Elem(), ApplicationGatewayBackendAddressResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendHttpSettingsInput)(nil)).Elem(), ApplicationGatewayBackendHttpSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendHttpSettingsArrayInput)(nil)).Elem(), ApplicationGatewayBackendHttpSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendHttpSettingsResponseInput)(nil)).Elem(), ApplicationGatewayBackendHttpSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayBackendHttpSettingsResponseArrayInput)(nil)).Elem(), ApplicationGatewayBackendHttpSettingsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayFrontendIPConfigurationInput)(nil)).Elem(), ApplicationGatewayFrontendIPConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayFrontendIPConfigurationArrayInput)(nil)).Elem(), ApplicationGatewayFrontendIPConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayFrontendIPConfigurationResponseInput)(nil)).Elem(), ApplicationGatewayFrontendIPConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayFrontendIPConfigurationResponseArrayInput)(nil)).Elem(), ApplicationGatewayFrontendIPConfigurationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayFrontendPortInput)(nil)).Elem(), ApplicationGatewayFrontendPortArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayFrontendPortArrayInput)(nil)).Elem(), ApplicationGatewayFrontendPortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayFrontendPortResponseInput)(nil)).Elem(), ApplicationGatewayFrontendPortResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayFrontendPortResponseArrayInput)(nil)).Elem(), ApplicationGatewayFrontendPortResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayHttpListenerInput)(nil)).Elem(), ApplicationGatewayHttpListenerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayHttpListenerArrayInput)(nil)).Elem(), ApplicationGatewayHttpListenerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayHttpListenerResponseInput)(nil)).Elem(), ApplicationGatewayHttpListenerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayHttpListenerResponseArrayInput)(nil)).Elem(), ApplicationGatewayHttpListenerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayIPConfigurationInput)(nil)).Elem(), ApplicationGatewayIPConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayIPConfigurationArrayInput)(nil)).Elem(), ApplicationGatewayIPConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayIPConfigurationResponseInput)(nil)).Elem(), ApplicationGatewayIPConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayIPConfigurationResponseArrayInput)(nil)).Elem(), ApplicationGatewayIPConfigurationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayRequestRoutingRuleInput)(nil)).Elem(), ApplicationGatewayRequestRoutingRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayRequestRoutingRuleArrayInput)(nil)).Elem(), ApplicationGatewayRequestRoutingRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayRequestRoutingRuleResponseInput)(nil)).Elem(), ApplicationGatewayRequestRoutingRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewayRequestRoutingRuleResponseArrayInput)(nil)).Elem(), ApplicationGatewayRequestRoutingRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewaySkuInput)(nil)).Elem(), ApplicationGatewaySkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewaySkuPtrInput)(nil)).Elem(), ApplicationGatewaySkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewaySkuResponseInput)(nil)).Elem(), ApplicationGatewaySkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewaySkuResponsePtrInput)(nil)).Elem(), ApplicationGatewaySkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewaySslCertificateInput)(nil)).Elem(), ApplicationGatewaySslCertificateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewaySslCertificateArrayInput)(nil)).Elem(), ApplicationGatewaySslCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewaySslCertificateResponseInput)(nil)).Elem(), ApplicationGatewaySslCertificateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGatewaySslCertificateResponseArrayInput)(nil)).Elem(), ApplicationGatewaySslCertificateResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendAddressPoolInput)(nil)).Elem(), BackendAddressPoolArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendAddressPoolArrayInput)(nil)).Elem(), BackendAddressPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendAddressPoolResponseInput)(nil)).Elem(), BackendAddressPoolResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendAddressPoolResponseArrayInput)(nil)).Elem(), BackendAddressPoolResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpOptionsInput)(nil)).Elem(), DhcpOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpOptionsPtrInput)(nil)).Elem(), DhcpOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpOptionsResponseInput)(nil)).Elem(), DhcpOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpOptionsResponsePtrInput)(nil)).Elem(), DhcpOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitAuthorizationTypeInput)(nil)).Elem(), ExpressRouteCircuitAuthorizationTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitAuthorizationTypeArrayInput)(nil)).Elem(), ExpressRouteCircuitAuthorizationTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitAuthorizationResponseInput)(nil)).Elem(), ExpressRouteCircuitAuthorizationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitAuthorizationResponseArrayInput)(nil)).Elem(), ExpressRouteCircuitAuthorizationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitPeeringTypeInput)(nil)).Elem(), ExpressRouteCircuitPeeringTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitPeeringTypeArrayInput)(nil)).Elem(), ExpressRouteCircuitPeeringTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitPeeringConfigInput)(nil)).Elem(), ExpressRouteCircuitPeeringConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitPeeringConfigPtrInput)(nil)).Elem(), ExpressRouteCircuitPeeringConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitPeeringConfigResponseInput)(nil)).Elem(), ExpressRouteCircuitPeeringConfigResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitPeeringConfigResponsePtrInput)(nil)).Elem(), ExpressRouteCircuitPeeringConfigResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitPeeringResponseInput)(nil)).Elem(), ExpressRouteCircuitPeeringResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitPeeringResponseArrayInput)(nil)).Elem(), ExpressRouteCircuitPeeringResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitServiceProviderPropertiesInput)(nil)).Elem(), ExpressRouteCircuitServiceProviderPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitServiceProviderPropertiesPtrInput)(nil)).Elem(), ExpressRouteCircuitServiceProviderPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitServiceProviderPropertiesResponseInput)(nil)).Elem(), ExpressRouteCircuitServiceProviderPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitServiceProviderPropertiesResponsePtrInput)(nil)).Elem(), ExpressRouteCircuitServiceProviderPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitSkuInput)(nil)).Elem(), ExpressRouteCircuitSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitSkuPtrInput)(nil)).Elem(), ExpressRouteCircuitSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitSkuResponseInput)(nil)).Elem(), ExpressRouteCircuitSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitSkuResponsePtrInput)(nil)).Elem(), ExpressRouteCircuitSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitStatsInput)(nil)).Elem(), ExpressRouteCircuitStatsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitStatsPtrInput)(nil)).Elem(), ExpressRouteCircuitStatsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitStatsResponseInput)(nil)).Elem(), ExpressRouteCircuitStatsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitStatsResponsePtrInput)(nil)).Elem(), ExpressRouteCircuitStatsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FrontendIpConfigurationInput)(nil)).Elem(), FrontendIpConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FrontendIpConfigurationArrayInput)(nil)).Elem(), FrontendIpConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FrontendIpConfigurationResponseInput)(nil)).Elem(), FrontendIpConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FrontendIpConfigurationResponseArrayInput)(nil)).Elem(), FrontendIpConfigurationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InboundNatPoolInput)(nil)).Elem(), InboundNatPoolArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InboundNatPoolArrayInput)(nil)).Elem(), InboundNatPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InboundNatPoolResponseInput)(nil)).Elem(), InboundNatPoolResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InboundNatPoolResponseArrayInput)(nil)).Elem(), InboundNatPoolResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InboundNatRuleInput)(nil)).Elem(), InboundNatRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InboundNatRuleArrayInput)(nil)).Elem(), InboundNatRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InboundNatRuleResponseInput)(nil)).Elem(), InboundNatRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InboundNatRuleResponseArrayInput)(nil)).Elem(), InboundNatRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancingRuleInput)(nil)).Elem(), LoadBalancingRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancingRuleArrayInput)(nil)).Elem(), LoadBalancingRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancingRuleResponseInput)(nil)).Elem(), LoadBalancingRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancingRuleResponseArrayInput)(nil)).Elem(), LoadBalancingRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceDnsSettingsInput)(nil)).Elem(), NetworkInterfaceDnsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceDnsSettingsPtrInput)(nil)).Elem(), NetworkInterfaceDnsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceDnsSettingsResponseInput)(nil)).Elem(), NetworkInterfaceDnsSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceDnsSettingsResponsePtrInput)(nil)).Elem(), NetworkInterfaceDnsSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceIpConfigurationInput)(nil)).Elem(), NetworkInterfaceIpConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceIpConfigurationArrayInput)(nil)).Elem(), NetworkInterfaceIpConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceIpConfigurationResponseInput)(nil)).Elem(), NetworkInterfaceIpConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceIpConfigurationResponseArrayInput)(nil)).Elem(), NetworkInterfaceIpConfigurationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutboundNatRuleInput)(nil)).Elem(), OutboundNatRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutboundNatRuleArrayInput)(nil)).Elem(), OutboundNatRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutboundNatRuleResponseInput)(nil)).Elem(), OutboundNatRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutboundNatRuleResponseArrayInput)(nil)).Elem(), OutboundNatRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProbeInput)(nil)).Elem(), ProbeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProbeArrayInput)(nil)).Elem(), ProbeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProbeResponseInput)(nil)).Elem(), ProbeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProbeResponseArrayInput)(nil)).Elem(), ProbeResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpAddressDnsSettingsInput)(nil)).Elem(), PublicIpAddressDnsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpAddressDnsSettingsPtrInput)(nil)).Elem(), PublicIpAddressDnsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpAddressDnsSettingsResponseInput)(nil)).Elem(), PublicIpAddressDnsSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpAddressDnsSettingsResponsePtrInput)(nil)).Elem(), PublicIpAddressDnsSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTypeInput)(nil)).Elem(), RouteTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTypeArrayInput)(nil)).Elem(), RouteTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteResponseInput)(nil)).Elem(), RouteResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteResponseArrayInput)(nil)).Elem(), RouteResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityRuleTypeInput)(nil)).Elem(), SecurityRuleTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityRuleTypeArrayInput)(nil)).Elem(), SecurityRuleTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityRuleResponseInput)(nil)).Elem(), SecurityRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityRuleResponseArrayInput)(nil)).Elem(), SecurityRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceInput)(nil)).Elem(), SubResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourcePtrInput)(nil)).Elem(), SubResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceArrayInput)(nil)).Elem(), SubResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceResponseInput)(nil)).Elem(), SubResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceResponsePtrInput)(nil)).Elem(), SubResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceResponseArrayInput)(nil)).Elem(), SubResourceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetTypeInput)(nil)).Elem(), SubnetTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetTypeArrayInput)(nil)).Elem(), SubnetTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetResponseInput)(nil)).Elem(), SubnetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetResponseArrayInput)(nil)).Elem(), SubnetResponseArray{})
	pulumi.RegisterOutputType(AddressSpaceOutput{})
	pulumi.RegisterOutputType(AddressSpacePtrOutput{})
	pulumi.RegisterOutputType(AddressSpaceResponseOutput{})
	pulumi.RegisterOutputType(AddressSpaceResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressPoolOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressPoolArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressPoolResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendHttpSettingsOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendHttpSettingsArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendHttpSettingsResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendHttpSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendIPConfigurationOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendPortOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendPortArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendPortResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendPortResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayHttpListenerOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayHttpListenerArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayHttpListenerResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayHttpListenerResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayIPConfigurationOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRequestRoutingRuleOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRequestRoutingRuleArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRequestRoutingRuleResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRequestRoutingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySkuOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySkuPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySkuResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslCertificateOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslCertificateArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslCertificateResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolArrayOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolResponseOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(DhcpOptionsOutput{})
	pulumi.RegisterOutputType(DhcpOptionsPtrOutput{})
	pulumi.RegisterOutputType(DhcpOptionsResponseOutput{})
	pulumi.RegisterOutputType(DhcpOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitAuthorizationTypeOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitAuthorizationTypeArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitAuthorizationResponseArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringTypeOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringTypeArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringConfigOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringConfigPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringConfigResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringResponseArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitServiceProviderPropertiesOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitServiceProviderPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitServiceProviderPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitStatsOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitStatsPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitStatsResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitStatsResponsePtrOutput{})
	pulumi.RegisterOutputType(FrontendIpConfigurationOutput{})
	pulumi.RegisterOutputType(FrontendIpConfigurationArrayOutput{})
	pulumi.RegisterOutputType(FrontendIpConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FrontendIpConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(InboundNatPoolOutput{})
	pulumi.RegisterOutputType(InboundNatPoolArrayOutput{})
	pulumi.RegisterOutputType(InboundNatPoolResponseOutput{})
	pulumi.RegisterOutputType(InboundNatPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(InboundNatRuleOutput{})
	pulumi.RegisterOutputType(InboundNatRuleArrayOutput{})
	pulumi.RegisterOutputType(InboundNatRuleResponseOutput{})
	pulumi.RegisterOutputType(InboundNatRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIpConfigurationOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIpConfigurationArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIpConfigurationResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIpConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleArrayOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleResponseOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ProbeOutput{})
	pulumi.RegisterOutputType(ProbeArrayOutput{})
	pulumi.RegisterOutputType(ProbeResponseOutput{})
	pulumi.RegisterOutputType(ProbeResponseArrayOutput{})
	pulumi.RegisterOutputType(PublicIpAddressDnsSettingsOutput{})
	pulumi.RegisterOutputType(PublicIpAddressDnsSettingsPtrOutput{})
	pulumi.RegisterOutputType(PublicIpAddressDnsSettingsResponseOutput{})
	pulumi.RegisterOutputType(PublicIpAddressDnsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(RouteTypeOutput{})
	pulumi.RegisterOutputType(RouteTypeArrayOutput{})
	pulumi.RegisterOutputType(RouteResponseOutput{})
	pulumi.RegisterOutputType(RouteResponseArrayOutput{})
	pulumi.RegisterOutputType(SecurityRuleTypeOutput{})
	pulumi.RegisterOutputType(SecurityRuleTypeArrayOutput{})
	pulumi.RegisterOutputType(SecurityRuleResponseOutput{})
	pulumi.RegisterOutputType(SecurityRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SubnetTypeOutput{})
	pulumi.RegisterOutputType(SubnetTypeArrayOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponseArrayOutput{})
}
