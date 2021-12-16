


package v20171101

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

type ApplicationGatewayAuthenticationCertificate struct {
	Data              *string `pulumi:"data"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
	Type              *string `pulumi:"type"`
}





type ApplicationGatewayAuthenticationCertificateInput interface {
	pulumi.Input

	ToApplicationGatewayAuthenticationCertificateOutput() ApplicationGatewayAuthenticationCertificateOutput
	ToApplicationGatewayAuthenticationCertificateOutputWithContext(context.Context) ApplicationGatewayAuthenticationCertificateOutput
}

type ApplicationGatewayAuthenticationCertificateArgs struct {
	Data              pulumi.StringPtrInput `pulumi:"data"`
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	Type              pulumi.StringPtrInput `pulumi:"type"`
}

func (ApplicationGatewayAuthenticationCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayAuthenticationCertificate)(nil)).Elem()
}

func (i ApplicationGatewayAuthenticationCertificateArgs) ToApplicationGatewayAuthenticationCertificateOutput() ApplicationGatewayAuthenticationCertificateOutput {
	return i.ToApplicationGatewayAuthenticationCertificateOutputWithContext(context.Background())
}

func (i ApplicationGatewayAuthenticationCertificateArgs) ToApplicationGatewayAuthenticationCertificateOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayAuthenticationCertificateOutput)
}





type ApplicationGatewayAuthenticationCertificateArrayInput interface {
	pulumi.Input

	ToApplicationGatewayAuthenticationCertificateArrayOutput() ApplicationGatewayAuthenticationCertificateArrayOutput
	ToApplicationGatewayAuthenticationCertificateArrayOutputWithContext(context.Context) ApplicationGatewayAuthenticationCertificateArrayOutput
}

type ApplicationGatewayAuthenticationCertificateArray []ApplicationGatewayAuthenticationCertificateInput

func (ApplicationGatewayAuthenticationCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayAuthenticationCertificate)(nil)).Elem()
}

func (i ApplicationGatewayAuthenticationCertificateArray) ToApplicationGatewayAuthenticationCertificateArrayOutput() ApplicationGatewayAuthenticationCertificateArrayOutput {
	return i.ToApplicationGatewayAuthenticationCertificateArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayAuthenticationCertificateArray) ToApplicationGatewayAuthenticationCertificateArrayOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayAuthenticationCertificateArrayOutput)
}

type ApplicationGatewayAuthenticationCertificateOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayAuthenticationCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayAuthenticationCertificate)(nil)).Elem()
}

func (o ApplicationGatewayAuthenticationCertificateOutput) ToApplicationGatewayAuthenticationCertificateOutput() ApplicationGatewayAuthenticationCertificateOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateOutput) ToApplicationGatewayAuthenticationCertificateOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayAuthenticationCertificateArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayAuthenticationCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayAuthenticationCertificate)(nil)).Elem()
}

func (o ApplicationGatewayAuthenticationCertificateArrayOutput) ToApplicationGatewayAuthenticationCertificateArrayOutput() ApplicationGatewayAuthenticationCertificateArrayOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateArrayOutput) ToApplicationGatewayAuthenticationCertificateArrayOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateArrayOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayAuthenticationCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayAuthenticationCertificate {
		return vs[0].([]ApplicationGatewayAuthenticationCertificate)[vs[1].(int)]
	}).(ApplicationGatewayAuthenticationCertificateOutput)
}

type ApplicationGatewayAuthenticationCertificateResponse struct {
	Data              *string `pulumi:"data"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
	Type              *string `pulumi:"type"`
}





type ApplicationGatewayAuthenticationCertificateResponseInput interface {
	pulumi.Input

	ToApplicationGatewayAuthenticationCertificateResponseOutput() ApplicationGatewayAuthenticationCertificateResponseOutput
	ToApplicationGatewayAuthenticationCertificateResponseOutputWithContext(context.Context) ApplicationGatewayAuthenticationCertificateResponseOutput
}

type ApplicationGatewayAuthenticationCertificateResponseArgs struct {
	Data              pulumi.StringPtrInput `pulumi:"data"`
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	Type              pulumi.StringPtrInput `pulumi:"type"`
}

func (ApplicationGatewayAuthenticationCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayAuthenticationCertificateResponse)(nil)).Elem()
}

func (i ApplicationGatewayAuthenticationCertificateResponseArgs) ToApplicationGatewayAuthenticationCertificateResponseOutput() ApplicationGatewayAuthenticationCertificateResponseOutput {
	return i.ToApplicationGatewayAuthenticationCertificateResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayAuthenticationCertificateResponseArgs) ToApplicationGatewayAuthenticationCertificateResponseOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayAuthenticationCertificateResponseOutput)
}





type ApplicationGatewayAuthenticationCertificateResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayAuthenticationCertificateResponseArrayOutput() ApplicationGatewayAuthenticationCertificateResponseArrayOutput
	ToApplicationGatewayAuthenticationCertificateResponseArrayOutputWithContext(context.Context) ApplicationGatewayAuthenticationCertificateResponseArrayOutput
}

type ApplicationGatewayAuthenticationCertificateResponseArray []ApplicationGatewayAuthenticationCertificateResponseInput

func (ApplicationGatewayAuthenticationCertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayAuthenticationCertificateResponse)(nil)).Elem()
}

func (i ApplicationGatewayAuthenticationCertificateResponseArray) ToApplicationGatewayAuthenticationCertificateResponseArrayOutput() ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
	return i.ToApplicationGatewayAuthenticationCertificateResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayAuthenticationCertificateResponseArray) ToApplicationGatewayAuthenticationCertificateResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayAuthenticationCertificateResponseArrayOutput)
}

type ApplicationGatewayAuthenticationCertificateResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayAuthenticationCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayAuthenticationCertificateResponse)(nil)).Elem()
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) ToApplicationGatewayAuthenticationCertificateResponseOutput() ApplicationGatewayAuthenticationCertificateResponseOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) ToApplicationGatewayAuthenticationCertificateResponseOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateResponseOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayAuthenticationCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayAuthenticationCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayAuthenticationCertificateResponse)(nil)).Elem()
}

func (o ApplicationGatewayAuthenticationCertificateResponseArrayOutput) ToApplicationGatewayAuthenticationCertificateResponseArrayOutput() ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateResponseArrayOutput) ToApplicationGatewayAuthenticationCertificateResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayAuthenticationCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayAuthenticationCertificateResponse {
		return vs[0].([]ApplicationGatewayAuthenticationCertificateResponse)[vs[1].(int)]
	}).(ApplicationGatewayAuthenticationCertificateResponseOutput)
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
	BackendIPConfigurations []NetworkInterfaceIPConfiguration  `pulumi:"backendIPConfigurations"`
	Etag                    *string                            `pulumi:"etag"`
	Id                      *string                            `pulumi:"id"`
	Name                    *string                            `pulumi:"name"`
	ProvisioningState       *string                            `pulumi:"provisioningState"`
	Type                    *string                            `pulumi:"type"`
}





type ApplicationGatewayBackendAddressPoolInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressPoolOutput() ApplicationGatewayBackendAddressPoolOutput
	ToApplicationGatewayBackendAddressPoolOutputWithContext(context.Context) ApplicationGatewayBackendAddressPoolOutput
}

type ApplicationGatewayBackendAddressPoolArgs struct {
	BackendAddresses        ApplicationGatewayBackendAddressArrayInput `pulumi:"backendAddresses"`
	BackendIPConfigurations NetworkInterfaceIPConfigurationArrayInput  `pulumi:"backendIPConfigurations"`
	Etag                    pulumi.StringPtrInput                      `pulumi:"etag"`
	Id                      pulumi.StringPtrInput                      `pulumi:"id"`
	Name                    pulumi.StringPtrInput                      `pulumi:"name"`
	ProvisioningState       pulumi.StringPtrInput                      `pulumi:"provisioningState"`
	Type                    pulumi.StringPtrInput                      `pulumi:"type"`
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

func (o ApplicationGatewayBackendAddressPoolOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) []NetworkInterfaceIPConfiguration {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationArrayOutput)
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

func (o ApplicationGatewayBackendAddressPoolOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	BackendIPConfigurations []NetworkInterfaceIPConfigurationResponse  `pulumi:"backendIPConfigurations"`
	Etag                    *string                                    `pulumi:"etag"`
	Id                      *string                                    `pulumi:"id"`
	Name                    *string                                    `pulumi:"name"`
	ProvisioningState       *string                                    `pulumi:"provisioningState"`
	Type                    *string                                    `pulumi:"type"`
}





type ApplicationGatewayBackendAddressPoolResponseInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressPoolResponseOutput() ApplicationGatewayBackendAddressPoolResponseOutput
	ToApplicationGatewayBackendAddressPoolResponseOutputWithContext(context.Context) ApplicationGatewayBackendAddressPoolResponseOutput
}

type ApplicationGatewayBackendAddressPoolResponseArgs struct {
	BackendAddresses        ApplicationGatewayBackendAddressResponseArrayInput `pulumi:"backendAddresses"`
	BackendIPConfigurations NetworkInterfaceIPConfigurationResponseArrayInput  `pulumi:"backendIPConfigurations"`
	Etag                    pulumi.StringPtrInput                              `pulumi:"etag"`
	Id                      pulumi.StringPtrInput                              `pulumi:"id"`
	Name                    pulumi.StringPtrInput                              `pulumi:"name"`
	ProvisioningState       pulumi.StringPtrInput                              `pulumi:"provisioningState"`
	Type                    pulumi.StringPtrInput                              `pulumi:"type"`
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

func (o ApplicationGatewayBackendAddressPoolResponseOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) []NetworkInterfaceIPConfigurationResponse {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
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

func (o ApplicationGatewayBackendAddressPoolResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	AffinityCookieName             *string                               `pulumi:"affinityCookieName"`
	AuthenticationCertificates     []SubResource                         `pulumi:"authenticationCertificates"`
	ConnectionDraining             *ApplicationGatewayConnectionDraining `pulumi:"connectionDraining"`
	CookieBasedAffinity            *string                               `pulumi:"cookieBasedAffinity"`
	Etag                           *string                               `pulumi:"etag"`
	HostName                       *string                               `pulumi:"hostName"`
	Id                             *string                               `pulumi:"id"`
	Name                           *string                               `pulumi:"name"`
	Path                           *string                               `pulumi:"path"`
	PickHostNameFromBackendAddress *bool                                 `pulumi:"pickHostNameFromBackendAddress"`
	Port                           *int                                  `pulumi:"port"`
	Probe                          *SubResource                          `pulumi:"probe"`
	ProbeEnabled                   *bool                                 `pulumi:"probeEnabled"`
	Protocol                       *string                               `pulumi:"protocol"`
	ProvisioningState              *string                               `pulumi:"provisioningState"`
	RequestTimeout                 *int                                  `pulumi:"requestTimeout"`
	Type                           *string                               `pulumi:"type"`
}





type ApplicationGatewayBackendHttpSettingsInput interface {
	pulumi.Input

	ToApplicationGatewayBackendHttpSettingsOutput() ApplicationGatewayBackendHttpSettingsOutput
	ToApplicationGatewayBackendHttpSettingsOutputWithContext(context.Context) ApplicationGatewayBackendHttpSettingsOutput
}

type ApplicationGatewayBackendHttpSettingsArgs struct {
	AffinityCookieName             pulumi.StringPtrInput                        `pulumi:"affinityCookieName"`
	AuthenticationCertificates     SubResourceArrayInput                        `pulumi:"authenticationCertificates"`
	ConnectionDraining             ApplicationGatewayConnectionDrainingPtrInput `pulumi:"connectionDraining"`
	CookieBasedAffinity            pulumi.StringPtrInput                        `pulumi:"cookieBasedAffinity"`
	Etag                           pulumi.StringPtrInput                        `pulumi:"etag"`
	HostName                       pulumi.StringPtrInput                        `pulumi:"hostName"`
	Id                             pulumi.StringPtrInput                        `pulumi:"id"`
	Name                           pulumi.StringPtrInput                        `pulumi:"name"`
	Path                           pulumi.StringPtrInput                        `pulumi:"path"`
	PickHostNameFromBackendAddress pulumi.BoolPtrInput                          `pulumi:"pickHostNameFromBackendAddress"`
	Port                           pulumi.IntPtrInput                           `pulumi:"port"`
	Probe                          SubResourcePtrInput                          `pulumi:"probe"`
	ProbeEnabled                   pulumi.BoolPtrInput                          `pulumi:"probeEnabled"`
	Protocol                       pulumi.StringPtrInput                        `pulumi:"protocol"`
	ProvisioningState              pulumi.StringPtrInput                        `pulumi:"provisioningState"`
	RequestTimeout                 pulumi.IntPtrInput                           `pulumi:"requestTimeout"`
	Type                           pulumi.StringPtrInput                        `pulumi:"type"`
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

func (o ApplicationGatewayBackendHttpSettingsOutput) AffinityCookieName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.AffinityCookieName }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) AuthenticationCertificates() SubResourceArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) []SubResource { return v.AuthenticationCertificates }).(SubResourceArrayOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) ConnectionDraining() ApplicationGatewayConnectionDrainingPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *ApplicationGatewayConnectionDraining {
		return v.ConnectionDraining
	}).(ApplicationGatewayConnectionDrainingPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) CookieBasedAffinity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.CookieBasedAffinity }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) PickHostNameFromBackendAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *bool { return v.PickHostNameFromBackendAddress }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Probe() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *SubResource { return v.Probe }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) ProbeEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *bool { return v.ProbeEnabled }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) RequestTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *int { return v.RequestTimeout }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	AffinityCookieName             *string                                       `pulumi:"affinityCookieName"`
	AuthenticationCertificates     []SubResourceResponse                         `pulumi:"authenticationCertificates"`
	ConnectionDraining             *ApplicationGatewayConnectionDrainingResponse `pulumi:"connectionDraining"`
	CookieBasedAffinity            *string                                       `pulumi:"cookieBasedAffinity"`
	Etag                           *string                                       `pulumi:"etag"`
	HostName                       *string                                       `pulumi:"hostName"`
	Id                             *string                                       `pulumi:"id"`
	Name                           *string                                       `pulumi:"name"`
	Path                           *string                                       `pulumi:"path"`
	PickHostNameFromBackendAddress *bool                                         `pulumi:"pickHostNameFromBackendAddress"`
	Port                           *int                                          `pulumi:"port"`
	Probe                          *SubResourceResponse                          `pulumi:"probe"`
	ProbeEnabled                   *bool                                         `pulumi:"probeEnabled"`
	Protocol                       *string                                       `pulumi:"protocol"`
	ProvisioningState              *string                                       `pulumi:"provisioningState"`
	RequestTimeout                 *int                                          `pulumi:"requestTimeout"`
	Type                           *string                                       `pulumi:"type"`
}





type ApplicationGatewayBackendHttpSettingsResponseInput interface {
	pulumi.Input

	ToApplicationGatewayBackendHttpSettingsResponseOutput() ApplicationGatewayBackendHttpSettingsResponseOutput
	ToApplicationGatewayBackendHttpSettingsResponseOutputWithContext(context.Context) ApplicationGatewayBackendHttpSettingsResponseOutput
}

type ApplicationGatewayBackendHttpSettingsResponseArgs struct {
	AffinityCookieName             pulumi.StringPtrInput                                `pulumi:"affinityCookieName"`
	AuthenticationCertificates     SubResourceResponseArrayInput                        `pulumi:"authenticationCertificates"`
	ConnectionDraining             ApplicationGatewayConnectionDrainingResponsePtrInput `pulumi:"connectionDraining"`
	CookieBasedAffinity            pulumi.StringPtrInput                                `pulumi:"cookieBasedAffinity"`
	Etag                           pulumi.StringPtrInput                                `pulumi:"etag"`
	HostName                       pulumi.StringPtrInput                                `pulumi:"hostName"`
	Id                             pulumi.StringPtrInput                                `pulumi:"id"`
	Name                           pulumi.StringPtrInput                                `pulumi:"name"`
	Path                           pulumi.StringPtrInput                                `pulumi:"path"`
	PickHostNameFromBackendAddress pulumi.BoolPtrInput                                  `pulumi:"pickHostNameFromBackendAddress"`
	Port                           pulumi.IntPtrInput                                   `pulumi:"port"`
	Probe                          SubResourceResponsePtrInput                          `pulumi:"probe"`
	ProbeEnabled                   pulumi.BoolPtrInput                                  `pulumi:"probeEnabled"`
	Protocol                       pulumi.StringPtrInput                                `pulumi:"protocol"`
	ProvisioningState              pulumi.StringPtrInput                                `pulumi:"provisioningState"`
	RequestTimeout                 pulumi.IntPtrInput                                   `pulumi:"requestTimeout"`
	Type                           pulumi.StringPtrInput                                `pulumi:"type"`
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

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) AffinityCookieName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.AffinityCookieName }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) AuthenticationCertificates() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) []SubResourceResponse {
		return v.AuthenticationCertificates
	}).(SubResourceResponseArrayOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) ConnectionDraining() ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *ApplicationGatewayConnectionDrainingResponse {
		return v.ConnectionDraining
	}).(ApplicationGatewayConnectionDrainingResponsePtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) CookieBasedAffinity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.CookieBasedAffinity }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) PickHostNameFromBackendAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *bool { return v.PickHostNameFromBackendAddress }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Probe() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *SubResourceResponse { return v.Probe }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) ProbeEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *bool { return v.ProbeEnabled }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) RequestTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *int { return v.RequestTimeout }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
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

type ApplicationGatewayConnectionDraining struct {
	DrainTimeoutInSec int  `pulumi:"drainTimeoutInSec"`
	Enabled           bool `pulumi:"enabled"`
}





type ApplicationGatewayConnectionDrainingInput interface {
	pulumi.Input

	ToApplicationGatewayConnectionDrainingOutput() ApplicationGatewayConnectionDrainingOutput
	ToApplicationGatewayConnectionDrainingOutputWithContext(context.Context) ApplicationGatewayConnectionDrainingOutput
}

type ApplicationGatewayConnectionDrainingArgs struct {
	DrainTimeoutInSec pulumi.IntInput  `pulumi:"drainTimeoutInSec"`
	Enabled           pulumi.BoolInput `pulumi:"enabled"`
}

func (ApplicationGatewayConnectionDrainingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayConnectionDraining)(nil)).Elem()
}

func (i ApplicationGatewayConnectionDrainingArgs) ToApplicationGatewayConnectionDrainingOutput() ApplicationGatewayConnectionDrainingOutput {
	return i.ToApplicationGatewayConnectionDrainingOutputWithContext(context.Background())
}

func (i ApplicationGatewayConnectionDrainingArgs) ToApplicationGatewayConnectionDrainingOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayConnectionDrainingOutput)
}

func (i ApplicationGatewayConnectionDrainingArgs) ToApplicationGatewayConnectionDrainingPtrOutput() ApplicationGatewayConnectionDrainingPtrOutput {
	return i.ToApplicationGatewayConnectionDrainingPtrOutputWithContext(context.Background())
}

func (i ApplicationGatewayConnectionDrainingArgs) ToApplicationGatewayConnectionDrainingPtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayConnectionDrainingOutput).ToApplicationGatewayConnectionDrainingPtrOutputWithContext(ctx)
}









type ApplicationGatewayConnectionDrainingPtrInput interface {
	pulumi.Input

	ToApplicationGatewayConnectionDrainingPtrOutput() ApplicationGatewayConnectionDrainingPtrOutput
	ToApplicationGatewayConnectionDrainingPtrOutputWithContext(context.Context) ApplicationGatewayConnectionDrainingPtrOutput
}

type applicationGatewayConnectionDrainingPtrType ApplicationGatewayConnectionDrainingArgs

func ApplicationGatewayConnectionDrainingPtr(v *ApplicationGatewayConnectionDrainingArgs) ApplicationGatewayConnectionDrainingPtrInput {
	return (*applicationGatewayConnectionDrainingPtrType)(v)
}

func (*applicationGatewayConnectionDrainingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayConnectionDraining)(nil)).Elem()
}

func (i *applicationGatewayConnectionDrainingPtrType) ToApplicationGatewayConnectionDrainingPtrOutput() ApplicationGatewayConnectionDrainingPtrOutput {
	return i.ToApplicationGatewayConnectionDrainingPtrOutputWithContext(context.Background())
}

func (i *applicationGatewayConnectionDrainingPtrType) ToApplicationGatewayConnectionDrainingPtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayConnectionDrainingPtrOutput)
}

type ApplicationGatewayConnectionDrainingOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayConnectionDrainingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayConnectionDraining)(nil)).Elem()
}

func (o ApplicationGatewayConnectionDrainingOutput) ToApplicationGatewayConnectionDrainingOutput() ApplicationGatewayConnectionDrainingOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingOutput) ToApplicationGatewayConnectionDrainingOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingOutput) ToApplicationGatewayConnectionDrainingPtrOutput() ApplicationGatewayConnectionDrainingPtrOutput {
	return o.ToApplicationGatewayConnectionDrainingPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayConnectionDrainingOutput) ToApplicationGatewayConnectionDrainingPtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayConnectionDraining) *ApplicationGatewayConnectionDraining {
		return &v
	}).(ApplicationGatewayConnectionDrainingPtrOutput)
}

func (o ApplicationGatewayConnectionDrainingOutput) DrainTimeoutInSec() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationGatewayConnectionDraining) int { return v.DrainTimeoutInSec }).(pulumi.IntOutput)
}

func (o ApplicationGatewayConnectionDrainingOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationGatewayConnectionDraining) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ApplicationGatewayConnectionDrainingPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayConnectionDrainingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayConnectionDraining)(nil)).Elem()
}

func (o ApplicationGatewayConnectionDrainingPtrOutput) ToApplicationGatewayConnectionDrainingPtrOutput() ApplicationGatewayConnectionDrainingPtrOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingPtrOutput) ToApplicationGatewayConnectionDrainingPtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingPtrOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingPtrOutput) Elem() ApplicationGatewayConnectionDrainingOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDraining) ApplicationGatewayConnectionDraining {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayConnectionDraining
		return ret
	}).(ApplicationGatewayConnectionDrainingOutput)
}

func (o ApplicationGatewayConnectionDrainingPtrOutput) DrainTimeoutInSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDraining) *int {
		if v == nil {
			return nil
		}
		return &v.DrainTimeoutInSec
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayConnectionDrainingPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDraining) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type ApplicationGatewayConnectionDrainingResponse struct {
	DrainTimeoutInSec int  `pulumi:"drainTimeoutInSec"`
	Enabled           bool `pulumi:"enabled"`
}





type ApplicationGatewayConnectionDrainingResponseInput interface {
	pulumi.Input

	ToApplicationGatewayConnectionDrainingResponseOutput() ApplicationGatewayConnectionDrainingResponseOutput
	ToApplicationGatewayConnectionDrainingResponseOutputWithContext(context.Context) ApplicationGatewayConnectionDrainingResponseOutput
}

type ApplicationGatewayConnectionDrainingResponseArgs struct {
	DrainTimeoutInSec pulumi.IntInput  `pulumi:"drainTimeoutInSec"`
	Enabled           pulumi.BoolInput `pulumi:"enabled"`
}

func (ApplicationGatewayConnectionDrainingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayConnectionDrainingResponse)(nil)).Elem()
}

func (i ApplicationGatewayConnectionDrainingResponseArgs) ToApplicationGatewayConnectionDrainingResponseOutput() ApplicationGatewayConnectionDrainingResponseOutput {
	return i.ToApplicationGatewayConnectionDrainingResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayConnectionDrainingResponseArgs) ToApplicationGatewayConnectionDrainingResponseOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayConnectionDrainingResponseOutput)
}

func (i ApplicationGatewayConnectionDrainingResponseArgs) ToApplicationGatewayConnectionDrainingResponsePtrOutput() ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return i.ToApplicationGatewayConnectionDrainingResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationGatewayConnectionDrainingResponseArgs) ToApplicationGatewayConnectionDrainingResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayConnectionDrainingResponseOutput).ToApplicationGatewayConnectionDrainingResponsePtrOutputWithContext(ctx)
}









type ApplicationGatewayConnectionDrainingResponsePtrInput interface {
	pulumi.Input

	ToApplicationGatewayConnectionDrainingResponsePtrOutput() ApplicationGatewayConnectionDrainingResponsePtrOutput
	ToApplicationGatewayConnectionDrainingResponsePtrOutputWithContext(context.Context) ApplicationGatewayConnectionDrainingResponsePtrOutput
}

type applicationGatewayConnectionDrainingResponsePtrType ApplicationGatewayConnectionDrainingResponseArgs

func ApplicationGatewayConnectionDrainingResponsePtr(v *ApplicationGatewayConnectionDrainingResponseArgs) ApplicationGatewayConnectionDrainingResponsePtrInput {
	return (*applicationGatewayConnectionDrainingResponsePtrType)(v)
}

func (*applicationGatewayConnectionDrainingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayConnectionDrainingResponse)(nil)).Elem()
}

func (i *applicationGatewayConnectionDrainingResponsePtrType) ToApplicationGatewayConnectionDrainingResponsePtrOutput() ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return i.ToApplicationGatewayConnectionDrainingResponsePtrOutputWithContext(context.Background())
}

func (i *applicationGatewayConnectionDrainingResponsePtrType) ToApplicationGatewayConnectionDrainingResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayConnectionDrainingResponsePtrOutput)
}

type ApplicationGatewayConnectionDrainingResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayConnectionDrainingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayConnectionDrainingResponse)(nil)).Elem()
}

func (o ApplicationGatewayConnectionDrainingResponseOutput) ToApplicationGatewayConnectionDrainingResponseOutput() ApplicationGatewayConnectionDrainingResponseOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingResponseOutput) ToApplicationGatewayConnectionDrainingResponseOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingResponseOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingResponseOutput) ToApplicationGatewayConnectionDrainingResponsePtrOutput() ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return o.ToApplicationGatewayConnectionDrainingResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayConnectionDrainingResponseOutput) ToApplicationGatewayConnectionDrainingResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayConnectionDrainingResponse) *ApplicationGatewayConnectionDrainingResponse {
		return &v
	}).(ApplicationGatewayConnectionDrainingResponsePtrOutput)
}

func (o ApplicationGatewayConnectionDrainingResponseOutput) DrainTimeoutInSec() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationGatewayConnectionDrainingResponse) int { return v.DrainTimeoutInSec }).(pulumi.IntOutput)
}

func (o ApplicationGatewayConnectionDrainingResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationGatewayConnectionDrainingResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ApplicationGatewayConnectionDrainingResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayConnectionDrainingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayConnectionDrainingResponse)(nil)).Elem()
}

func (o ApplicationGatewayConnectionDrainingResponsePtrOutput) ToApplicationGatewayConnectionDrainingResponsePtrOutput() ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingResponsePtrOutput) ToApplicationGatewayConnectionDrainingResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingResponsePtrOutput) Elem() ApplicationGatewayConnectionDrainingResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDrainingResponse) ApplicationGatewayConnectionDrainingResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayConnectionDrainingResponse
		return ret
	}).(ApplicationGatewayConnectionDrainingResponseOutput)
}

func (o ApplicationGatewayConnectionDrainingResponsePtrOutput) DrainTimeoutInSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDrainingResponse) *int {
		if v == nil {
			return nil
		}
		return &v.DrainTimeoutInSec
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayConnectionDrainingResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDrainingResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroup struct {
	RuleGroupName string `pulumi:"ruleGroupName"`
	Rules         []int  `pulumi:"rules"`
}





type ApplicationGatewayFirewallDisabledRuleGroupInput interface {
	pulumi.Input

	ToApplicationGatewayFirewallDisabledRuleGroupOutput() ApplicationGatewayFirewallDisabledRuleGroupOutput
	ToApplicationGatewayFirewallDisabledRuleGroupOutputWithContext(context.Context) ApplicationGatewayFirewallDisabledRuleGroupOutput
}

type ApplicationGatewayFirewallDisabledRuleGroupArgs struct {
	RuleGroupName pulumi.StringInput   `pulumi:"ruleGroupName"`
	Rules         pulumi.IntArrayInput `pulumi:"rules"`
}

func (ApplicationGatewayFirewallDisabledRuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFirewallDisabledRuleGroup)(nil)).Elem()
}

func (i ApplicationGatewayFirewallDisabledRuleGroupArgs) ToApplicationGatewayFirewallDisabledRuleGroupOutput() ApplicationGatewayFirewallDisabledRuleGroupOutput {
	return i.ToApplicationGatewayFirewallDisabledRuleGroupOutputWithContext(context.Background())
}

func (i ApplicationGatewayFirewallDisabledRuleGroupArgs) ToApplicationGatewayFirewallDisabledRuleGroupOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFirewallDisabledRuleGroupOutput)
}





type ApplicationGatewayFirewallDisabledRuleGroupArrayInput interface {
	pulumi.Input

	ToApplicationGatewayFirewallDisabledRuleGroupArrayOutput() ApplicationGatewayFirewallDisabledRuleGroupArrayOutput
	ToApplicationGatewayFirewallDisabledRuleGroupArrayOutputWithContext(context.Context) ApplicationGatewayFirewallDisabledRuleGroupArrayOutput
}

type ApplicationGatewayFirewallDisabledRuleGroupArray []ApplicationGatewayFirewallDisabledRuleGroupInput

func (ApplicationGatewayFirewallDisabledRuleGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFirewallDisabledRuleGroup)(nil)).Elem()
}

func (i ApplicationGatewayFirewallDisabledRuleGroupArray) ToApplicationGatewayFirewallDisabledRuleGroupArrayOutput() ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return i.ToApplicationGatewayFirewallDisabledRuleGroupArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayFirewallDisabledRuleGroupArray) ToApplicationGatewayFirewallDisabledRuleGroupArrayOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFirewallDisabledRuleGroupArrayOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroupOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFirewallDisabledRuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFirewallDisabledRuleGroup)(nil)).Elem()
}

func (o ApplicationGatewayFirewallDisabledRuleGroupOutput) ToApplicationGatewayFirewallDisabledRuleGroupOutput() ApplicationGatewayFirewallDisabledRuleGroupOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupOutput) ToApplicationGatewayFirewallDisabledRuleGroupOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayFirewallDisabledRuleGroup) string { return v.RuleGroupName }).(pulumi.StringOutput)
}

func (o ApplicationGatewayFirewallDisabledRuleGroupOutput) Rules() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayFirewallDisabledRuleGroup) []int { return v.Rules }).(pulumi.IntArrayOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroupArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFirewallDisabledRuleGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFirewallDisabledRuleGroup)(nil)).Elem()
}

func (o ApplicationGatewayFirewallDisabledRuleGroupArrayOutput) ToApplicationGatewayFirewallDisabledRuleGroupArrayOutput() ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupArrayOutput) ToApplicationGatewayFirewallDisabledRuleGroupArrayOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFirewallDisabledRuleGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFirewallDisabledRuleGroup {
		return vs[0].([]ApplicationGatewayFirewallDisabledRuleGroup)[vs[1].(int)]
	}).(ApplicationGatewayFirewallDisabledRuleGroupOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroupResponse struct {
	RuleGroupName string `pulumi:"ruleGroupName"`
	Rules         []int  `pulumi:"rules"`
}





type ApplicationGatewayFirewallDisabledRuleGroupResponseInput interface {
	pulumi.Input

	ToApplicationGatewayFirewallDisabledRuleGroupResponseOutput() ApplicationGatewayFirewallDisabledRuleGroupResponseOutput
	ToApplicationGatewayFirewallDisabledRuleGroupResponseOutputWithContext(context.Context) ApplicationGatewayFirewallDisabledRuleGroupResponseOutput
}

type ApplicationGatewayFirewallDisabledRuleGroupResponseArgs struct {
	RuleGroupName pulumi.StringInput   `pulumi:"ruleGroupName"`
	Rules         pulumi.IntArrayInput `pulumi:"rules"`
}

func (ApplicationGatewayFirewallDisabledRuleGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFirewallDisabledRuleGroupResponse)(nil)).Elem()
}

func (i ApplicationGatewayFirewallDisabledRuleGroupResponseArgs) ToApplicationGatewayFirewallDisabledRuleGroupResponseOutput() ApplicationGatewayFirewallDisabledRuleGroupResponseOutput {
	return i.ToApplicationGatewayFirewallDisabledRuleGroupResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayFirewallDisabledRuleGroupResponseArgs) ToApplicationGatewayFirewallDisabledRuleGroupResponseOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFirewallDisabledRuleGroupResponseOutput)
}





type ApplicationGatewayFirewallDisabledRuleGroupResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput() ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput
	ToApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutputWithContext(context.Context) ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput
}

type ApplicationGatewayFirewallDisabledRuleGroupResponseArray []ApplicationGatewayFirewallDisabledRuleGroupResponseInput

func (ApplicationGatewayFirewallDisabledRuleGroupResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFirewallDisabledRuleGroupResponse)(nil)).Elem()
}

func (i ApplicationGatewayFirewallDisabledRuleGroupResponseArray) ToApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput() ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput {
	return i.ToApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayFirewallDisabledRuleGroupResponseArray) ToApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroupResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFirewallDisabledRuleGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFirewallDisabledRuleGroupResponse)(nil)).Elem()
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseOutput) ToApplicationGatewayFirewallDisabledRuleGroupResponseOutput() ApplicationGatewayFirewallDisabledRuleGroupResponseOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseOutput) ToApplicationGatewayFirewallDisabledRuleGroupResponseOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupResponseOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayFirewallDisabledRuleGroupResponse) string { return v.RuleGroupName }).(pulumi.StringOutput)
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseOutput) Rules() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayFirewallDisabledRuleGroupResponse) []int { return v.Rules }).(pulumi.IntArrayOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFirewallDisabledRuleGroupResponse)(nil)).Elem()
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput) ToApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput() ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput) ToApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFirewallDisabledRuleGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFirewallDisabledRuleGroupResponse {
		return vs[0].([]ApplicationGatewayFirewallDisabledRuleGroupResponse)[vs[1].(int)]
	}).(ApplicationGatewayFirewallDisabledRuleGroupResponseOutput)
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
	Type                      *string      `pulumi:"type"`
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
	Type                      pulumi.StringPtrInput `pulumi:"type"`
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

func (o ApplicationGatewayFrontendIPConfigurationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	Type                      *string              `pulumi:"type"`
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
	Type                      pulumi.StringPtrInput       `pulumi:"type"`
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

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	Type              *string `pulumi:"type"`
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
	Type              pulumi.StringPtrInput `pulumi:"type"`
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

func (o ApplicationGatewayFrontendPortOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	Type              *string `pulumi:"type"`
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
	Type              pulumi.StringPtrInput `pulumi:"type"`
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

func (o ApplicationGatewayFrontendPortResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	Etag                        *string      `pulumi:"etag"`
	FrontendIPConfiguration     *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPort                *SubResource `pulumi:"frontendPort"`
	HostName                    *string      `pulumi:"hostName"`
	Id                          *string      `pulumi:"id"`
	Name                        *string      `pulumi:"name"`
	Protocol                    *string      `pulumi:"protocol"`
	ProvisioningState           *string      `pulumi:"provisioningState"`
	RequireServerNameIndication *bool        `pulumi:"requireServerNameIndication"`
	SslCertificate              *SubResource `pulumi:"sslCertificate"`
	Type                        *string      `pulumi:"type"`
}





type ApplicationGatewayHttpListenerInput interface {
	pulumi.Input

	ToApplicationGatewayHttpListenerOutput() ApplicationGatewayHttpListenerOutput
	ToApplicationGatewayHttpListenerOutputWithContext(context.Context) ApplicationGatewayHttpListenerOutput
}

type ApplicationGatewayHttpListenerArgs struct {
	Etag                        pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration     SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPort                SubResourcePtrInput   `pulumi:"frontendPort"`
	HostName                    pulumi.StringPtrInput `pulumi:"hostName"`
	Id                          pulumi.StringPtrInput `pulumi:"id"`
	Name                        pulumi.StringPtrInput `pulumi:"name"`
	Protocol                    pulumi.StringPtrInput `pulumi:"protocol"`
	ProvisioningState           pulumi.StringPtrInput `pulumi:"provisioningState"`
	RequireServerNameIndication pulumi.BoolPtrInput   `pulumi:"requireServerNameIndication"`
	SslCertificate              SubResourcePtrInput   `pulumi:"sslCertificate"`
	Type                        pulumi.StringPtrInput `pulumi:"type"`
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

func (o ApplicationGatewayHttpListenerOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.HostName }).(pulumi.StringPtrOutput)
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

func (o ApplicationGatewayHttpListenerOutput) RequireServerNameIndication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *bool { return v.RequireServerNameIndication }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) SslCertificate() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *SubResource { return v.SslCertificate }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	Etag                        *string              `pulumi:"etag"`
	FrontendIPConfiguration     *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	FrontendPort                *SubResourceResponse `pulumi:"frontendPort"`
	HostName                    *string              `pulumi:"hostName"`
	Id                          *string              `pulumi:"id"`
	Name                        *string              `pulumi:"name"`
	Protocol                    *string              `pulumi:"protocol"`
	ProvisioningState           *string              `pulumi:"provisioningState"`
	RequireServerNameIndication *bool                `pulumi:"requireServerNameIndication"`
	SslCertificate              *SubResourceResponse `pulumi:"sslCertificate"`
	Type                        *string              `pulumi:"type"`
}





type ApplicationGatewayHttpListenerResponseInput interface {
	pulumi.Input

	ToApplicationGatewayHttpListenerResponseOutput() ApplicationGatewayHttpListenerResponseOutput
	ToApplicationGatewayHttpListenerResponseOutputWithContext(context.Context) ApplicationGatewayHttpListenerResponseOutput
}

type ApplicationGatewayHttpListenerResponseArgs struct {
	Etag                        pulumi.StringPtrInput       `pulumi:"etag"`
	FrontendIPConfiguration     SubResourceResponsePtrInput `pulumi:"frontendIPConfiguration"`
	FrontendPort                SubResourceResponsePtrInput `pulumi:"frontendPort"`
	HostName                    pulumi.StringPtrInput       `pulumi:"hostName"`
	Id                          pulumi.StringPtrInput       `pulumi:"id"`
	Name                        pulumi.StringPtrInput       `pulumi:"name"`
	Protocol                    pulumi.StringPtrInput       `pulumi:"protocol"`
	ProvisioningState           pulumi.StringPtrInput       `pulumi:"provisioningState"`
	RequireServerNameIndication pulumi.BoolPtrInput         `pulumi:"requireServerNameIndication"`
	SslCertificate              SubResourceResponsePtrInput `pulumi:"sslCertificate"`
	Type                        pulumi.StringPtrInput       `pulumi:"type"`
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

func (o ApplicationGatewayHttpListenerResponseOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.HostName }).(pulumi.StringPtrOutput)
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

func (o ApplicationGatewayHttpListenerResponseOutput) RequireServerNameIndication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *bool { return v.RequireServerNameIndication }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) SslCertificate() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *SubResourceResponse { return v.SslCertificate }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	Type              *string      `pulumi:"type"`
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
	Type              pulumi.StringPtrInput `pulumi:"type"`
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

func (o ApplicationGatewayIPConfigurationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	Type              *string              `pulumi:"type"`
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
	Type              pulumi.StringPtrInput       `pulumi:"type"`
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

func (o ApplicationGatewayIPConfigurationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
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

type ApplicationGatewayPathRule struct {
	BackendAddressPool    *SubResource `pulumi:"backendAddressPool"`
	BackendHttpSettings   *SubResource `pulumi:"backendHttpSettings"`
	Etag                  *string      `pulumi:"etag"`
	Id                    *string      `pulumi:"id"`
	Name                  *string      `pulumi:"name"`
	Paths                 []string     `pulumi:"paths"`
	ProvisioningState     *string      `pulumi:"provisioningState"`
	RedirectConfiguration *SubResource `pulumi:"redirectConfiguration"`
	Type                  *string      `pulumi:"type"`
}





type ApplicationGatewayPathRuleInput interface {
	pulumi.Input

	ToApplicationGatewayPathRuleOutput() ApplicationGatewayPathRuleOutput
	ToApplicationGatewayPathRuleOutputWithContext(context.Context) ApplicationGatewayPathRuleOutput
}

type ApplicationGatewayPathRuleArgs struct {
	BackendAddressPool    SubResourcePtrInput     `pulumi:"backendAddressPool"`
	BackendHttpSettings   SubResourcePtrInput     `pulumi:"backendHttpSettings"`
	Etag                  pulumi.StringPtrInput   `pulumi:"etag"`
	Id                    pulumi.StringPtrInput   `pulumi:"id"`
	Name                  pulumi.StringPtrInput   `pulumi:"name"`
	Paths                 pulumi.StringArrayInput `pulumi:"paths"`
	ProvisioningState     pulumi.StringPtrInput   `pulumi:"provisioningState"`
	RedirectConfiguration SubResourcePtrInput     `pulumi:"redirectConfiguration"`
	Type                  pulumi.StringPtrInput   `pulumi:"type"`
}

func (ApplicationGatewayPathRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayPathRule)(nil)).Elem()
}

func (i ApplicationGatewayPathRuleArgs) ToApplicationGatewayPathRuleOutput() ApplicationGatewayPathRuleOutput {
	return i.ToApplicationGatewayPathRuleOutputWithContext(context.Background())
}

func (i ApplicationGatewayPathRuleArgs) ToApplicationGatewayPathRuleOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayPathRuleOutput)
}





type ApplicationGatewayPathRuleArrayInput interface {
	pulumi.Input

	ToApplicationGatewayPathRuleArrayOutput() ApplicationGatewayPathRuleArrayOutput
	ToApplicationGatewayPathRuleArrayOutputWithContext(context.Context) ApplicationGatewayPathRuleArrayOutput
}

type ApplicationGatewayPathRuleArray []ApplicationGatewayPathRuleInput

func (ApplicationGatewayPathRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayPathRule)(nil)).Elem()
}

func (i ApplicationGatewayPathRuleArray) ToApplicationGatewayPathRuleArrayOutput() ApplicationGatewayPathRuleArrayOutput {
	return i.ToApplicationGatewayPathRuleArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayPathRuleArray) ToApplicationGatewayPathRuleArrayOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayPathRuleArrayOutput)
}

type ApplicationGatewayPathRuleOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPathRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayPathRule)(nil)).Elem()
}

func (o ApplicationGatewayPathRuleOutput) ToApplicationGatewayPathRuleOutput() ApplicationGatewayPathRuleOutput {
	return o
}

func (o ApplicationGatewayPathRuleOutput) ToApplicationGatewayPathRuleOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleOutput {
	return o
}

func (o ApplicationGatewayPathRuleOutput) BackendAddressPool() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *SubResource { return v.BackendAddressPool }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) BackendHttpSettings() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *SubResource { return v.BackendHttpSettings }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewayPathRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) RedirectConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *SubResource { return v.RedirectConfiguration }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayPathRuleArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPathRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayPathRule)(nil)).Elem()
}

func (o ApplicationGatewayPathRuleArrayOutput) ToApplicationGatewayPathRuleArrayOutput() ApplicationGatewayPathRuleArrayOutput {
	return o
}

func (o ApplicationGatewayPathRuleArrayOutput) ToApplicationGatewayPathRuleArrayOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleArrayOutput {
	return o
}

func (o ApplicationGatewayPathRuleArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayPathRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayPathRule {
		return vs[0].([]ApplicationGatewayPathRule)[vs[1].(int)]
	}).(ApplicationGatewayPathRuleOutput)
}

type ApplicationGatewayPathRuleResponse struct {
	BackendAddressPool    *SubResourceResponse `pulumi:"backendAddressPool"`
	BackendHttpSettings   *SubResourceResponse `pulumi:"backendHttpSettings"`
	Etag                  *string              `pulumi:"etag"`
	Id                    *string              `pulumi:"id"`
	Name                  *string              `pulumi:"name"`
	Paths                 []string             `pulumi:"paths"`
	ProvisioningState     *string              `pulumi:"provisioningState"`
	RedirectConfiguration *SubResourceResponse `pulumi:"redirectConfiguration"`
	Type                  *string              `pulumi:"type"`
}





type ApplicationGatewayPathRuleResponseInput interface {
	pulumi.Input

	ToApplicationGatewayPathRuleResponseOutput() ApplicationGatewayPathRuleResponseOutput
	ToApplicationGatewayPathRuleResponseOutputWithContext(context.Context) ApplicationGatewayPathRuleResponseOutput
}

type ApplicationGatewayPathRuleResponseArgs struct {
	BackendAddressPool    SubResourceResponsePtrInput `pulumi:"backendAddressPool"`
	BackendHttpSettings   SubResourceResponsePtrInput `pulumi:"backendHttpSettings"`
	Etag                  pulumi.StringPtrInput       `pulumi:"etag"`
	Id                    pulumi.StringPtrInput       `pulumi:"id"`
	Name                  pulumi.StringPtrInput       `pulumi:"name"`
	Paths                 pulumi.StringArrayInput     `pulumi:"paths"`
	ProvisioningState     pulumi.StringPtrInput       `pulumi:"provisioningState"`
	RedirectConfiguration SubResourceResponsePtrInput `pulumi:"redirectConfiguration"`
	Type                  pulumi.StringPtrInput       `pulumi:"type"`
}

func (ApplicationGatewayPathRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayPathRuleResponse)(nil)).Elem()
}

func (i ApplicationGatewayPathRuleResponseArgs) ToApplicationGatewayPathRuleResponseOutput() ApplicationGatewayPathRuleResponseOutput {
	return i.ToApplicationGatewayPathRuleResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayPathRuleResponseArgs) ToApplicationGatewayPathRuleResponseOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayPathRuleResponseOutput)
}





type ApplicationGatewayPathRuleResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayPathRuleResponseArrayOutput() ApplicationGatewayPathRuleResponseArrayOutput
	ToApplicationGatewayPathRuleResponseArrayOutputWithContext(context.Context) ApplicationGatewayPathRuleResponseArrayOutput
}

type ApplicationGatewayPathRuleResponseArray []ApplicationGatewayPathRuleResponseInput

func (ApplicationGatewayPathRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayPathRuleResponse)(nil)).Elem()
}

func (i ApplicationGatewayPathRuleResponseArray) ToApplicationGatewayPathRuleResponseArrayOutput() ApplicationGatewayPathRuleResponseArrayOutput {
	return i.ToApplicationGatewayPathRuleResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayPathRuleResponseArray) ToApplicationGatewayPathRuleResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayPathRuleResponseArrayOutput)
}

type ApplicationGatewayPathRuleResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPathRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayPathRuleResponse)(nil)).Elem()
}

func (o ApplicationGatewayPathRuleResponseOutput) ToApplicationGatewayPathRuleResponseOutput() ApplicationGatewayPathRuleResponseOutput {
	return o
}

func (o ApplicationGatewayPathRuleResponseOutput) ToApplicationGatewayPathRuleResponseOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleResponseOutput {
	return o
}

func (o ApplicationGatewayPathRuleResponseOutput) BackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) BackendHttpSettings() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *SubResourceResponse { return v.BackendHttpSettings }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) RedirectConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *SubResourceResponse { return v.RedirectConfiguration }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayPathRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPathRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayPathRuleResponse)(nil)).Elem()
}

func (o ApplicationGatewayPathRuleResponseArrayOutput) ToApplicationGatewayPathRuleResponseArrayOutput() ApplicationGatewayPathRuleResponseArrayOutput {
	return o
}

func (o ApplicationGatewayPathRuleResponseArrayOutput) ToApplicationGatewayPathRuleResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleResponseArrayOutput {
	return o
}

func (o ApplicationGatewayPathRuleResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayPathRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayPathRuleResponse {
		return vs[0].([]ApplicationGatewayPathRuleResponse)[vs[1].(int)]
	}).(ApplicationGatewayPathRuleResponseOutput)
}

type ApplicationGatewayProbe struct {
	Etag                                *string                                     `pulumi:"etag"`
	Host                                *string                                     `pulumi:"host"`
	Id                                  *string                                     `pulumi:"id"`
	Interval                            *int                                        `pulumi:"interval"`
	Match                               *ApplicationGatewayProbeHealthResponseMatch `pulumi:"match"`
	MinServers                          *int                                        `pulumi:"minServers"`
	Name                                *string                                     `pulumi:"name"`
	Path                                *string                                     `pulumi:"path"`
	PickHostNameFromBackendHttpSettings *bool                                       `pulumi:"pickHostNameFromBackendHttpSettings"`
	Protocol                            *string                                     `pulumi:"protocol"`
	ProvisioningState                   *string                                     `pulumi:"provisioningState"`
	Timeout                             *int                                        `pulumi:"timeout"`
	Type                                *string                                     `pulumi:"type"`
	UnhealthyThreshold                  *int                                        `pulumi:"unhealthyThreshold"`
}





type ApplicationGatewayProbeInput interface {
	pulumi.Input

	ToApplicationGatewayProbeOutput() ApplicationGatewayProbeOutput
	ToApplicationGatewayProbeOutputWithContext(context.Context) ApplicationGatewayProbeOutput
}

type ApplicationGatewayProbeArgs struct {
	Etag                                pulumi.StringPtrInput                              `pulumi:"etag"`
	Host                                pulumi.StringPtrInput                              `pulumi:"host"`
	Id                                  pulumi.StringPtrInput                              `pulumi:"id"`
	Interval                            pulumi.IntPtrInput                                 `pulumi:"interval"`
	Match                               ApplicationGatewayProbeHealthResponseMatchPtrInput `pulumi:"match"`
	MinServers                          pulumi.IntPtrInput                                 `pulumi:"minServers"`
	Name                                pulumi.StringPtrInput                              `pulumi:"name"`
	Path                                pulumi.StringPtrInput                              `pulumi:"path"`
	PickHostNameFromBackendHttpSettings pulumi.BoolPtrInput                                `pulumi:"pickHostNameFromBackendHttpSettings"`
	Protocol                            pulumi.StringPtrInput                              `pulumi:"protocol"`
	ProvisioningState                   pulumi.StringPtrInput                              `pulumi:"provisioningState"`
	Timeout                             pulumi.IntPtrInput                                 `pulumi:"timeout"`
	Type                                pulumi.StringPtrInput                              `pulumi:"type"`
	UnhealthyThreshold                  pulumi.IntPtrInput                                 `pulumi:"unhealthyThreshold"`
}

func (ApplicationGatewayProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbe)(nil)).Elem()
}

func (i ApplicationGatewayProbeArgs) ToApplicationGatewayProbeOutput() ApplicationGatewayProbeOutput {
	return i.ToApplicationGatewayProbeOutputWithContext(context.Background())
}

func (i ApplicationGatewayProbeArgs) ToApplicationGatewayProbeOutputWithContext(ctx context.Context) ApplicationGatewayProbeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeOutput)
}





type ApplicationGatewayProbeArrayInput interface {
	pulumi.Input

	ToApplicationGatewayProbeArrayOutput() ApplicationGatewayProbeArrayOutput
	ToApplicationGatewayProbeArrayOutputWithContext(context.Context) ApplicationGatewayProbeArrayOutput
}

type ApplicationGatewayProbeArray []ApplicationGatewayProbeInput

func (ApplicationGatewayProbeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayProbe)(nil)).Elem()
}

func (i ApplicationGatewayProbeArray) ToApplicationGatewayProbeArrayOutput() ApplicationGatewayProbeArrayOutput {
	return i.ToApplicationGatewayProbeArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayProbeArray) ToApplicationGatewayProbeArrayOutputWithContext(ctx context.Context) ApplicationGatewayProbeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeArrayOutput)
}

type ApplicationGatewayProbeOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbe)(nil)).Elem()
}

func (o ApplicationGatewayProbeOutput) ToApplicationGatewayProbeOutput() ApplicationGatewayProbeOutput {
	return o
}

func (o ApplicationGatewayProbeOutput) ToApplicationGatewayProbeOutputWithContext(ctx context.Context) ApplicationGatewayProbeOutput {
	return o
}

func (o ApplicationGatewayProbeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Match() ApplicationGatewayProbeHealthResponseMatchPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *ApplicationGatewayProbeHealthResponseMatch { return v.Match }).(ApplicationGatewayProbeHealthResponseMatchPtrOutput)
}

func (o ApplicationGatewayProbeOutput) MinServers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *int { return v.MinServers }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) PickHostNameFromBackendHttpSettings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *bool { return v.PickHostNameFromBackendHttpSettings }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) UnhealthyThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *int { return v.UnhealthyThreshold }).(pulumi.IntPtrOutput)
}

type ApplicationGatewayProbeArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayProbe)(nil)).Elem()
}

func (o ApplicationGatewayProbeArrayOutput) ToApplicationGatewayProbeArrayOutput() ApplicationGatewayProbeArrayOutput {
	return o
}

func (o ApplicationGatewayProbeArrayOutput) ToApplicationGatewayProbeArrayOutputWithContext(ctx context.Context) ApplicationGatewayProbeArrayOutput {
	return o
}

func (o ApplicationGatewayProbeArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayProbeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayProbe {
		return vs[0].([]ApplicationGatewayProbe)[vs[1].(int)]
	}).(ApplicationGatewayProbeOutput)
}

type ApplicationGatewayProbeHealthResponseMatch struct {
	Body        *string  `pulumi:"body"`
	StatusCodes []string `pulumi:"statusCodes"`
}





type ApplicationGatewayProbeHealthResponseMatchInput interface {
	pulumi.Input

	ToApplicationGatewayProbeHealthResponseMatchOutput() ApplicationGatewayProbeHealthResponseMatchOutput
	ToApplicationGatewayProbeHealthResponseMatchOutputWithContext(context.Context) ApplicationGatewayProbeHealthResponseMatchOutput
}

type ApplicationGatewayProbeHealthResponseMatchArgs struct {
	Body        pulumi.StringPtrInput   `pulumi:"body"`
	StatusCodes pulumi.StringArrayInput `pulumi:"statusCodes"`
}

func (ApplicationGatewayProbeHealthResponseMatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbeHealthResponseMatch)(nil)).Elem()
}

func (i ApplicationGatewayProbeHealthResponseMatchArgs) ToApplicationGatewayProbeHealthResponseMatchOutput() ApplicationGatewayProbeHealthResponseMatchOutput {
	return i.ToApplicationGatewayProbeHealthResponseMatchOutputWithContext(context.Background())
}

func (i ApplicationGatewayProbeHealthResponseMatchArgs) ToApplicationGatewayProbeHealthResponseMatchOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeHealthResponseMatchOutput)
}

func (i ApplicationGatewayProbeHealthResponseMatchArgs) ToApplicationGatewayProbeHealthResponseMatchPtrOutput() ApplicationGatewayProbeHealthResponseMatchPtrOutput {
	return i.ToApplicationGatewayProbeHealthResponseMatchPtrOutputWithContext(context.Background())
}

func (i ApplicationGatewayProbeHealthResponseMatchArgs) ToApplicationGatewayProbeHealthResponseMatchPtrOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeHealthResponseMatchOutput).ToApplicationGatewayProbeHealthResponseMatchPtrOutputWithContext(ctx)
}









type ApplicationGatewayProbeHealthResponseMatchPtrInput interface {
	pulumi.Input

	ToApplicationGatewayProbeHealthResponseMatchPtrOutput() ApplicationGatewayProbeHealthResponseMatchPtrOutput
	ToApplicationGatewayProbeHealthResponseMatchPtrOutputWithContext(context.Context) ApplicationGatewayProbeHealthResponseMatchPtrOutput
}

type applicationGatewayProbeHealthResponseMatchPtrType ApplicationGatewayProbeHealthResponseMatchArgs

func ApplicationGatewayProbeHealthResponseMatchPtr(v *ApplicationGatewayProbeHealthResponseMatchArgs) ApplicationGatewayProbeHealthResponseMatchPtrInput {
	return (*applicationGatewayProbeHealthResponseMatchPtrType)(v)
}

func (*applicationGatewayProbeHealthResponseMatchPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayProbeHealthResponseMatch)(nil)).Elem()
}

func (i *applicationGatewayProbeHealthResponseMatchPtrType) ToApplicationGatewayProbeHealthResponseMatchPtrOutput() ApplicationGatewayProbeHealthResponseMatchPtrOutput {
	return i.ToApplicationGatewayProbeHealthResponseMatchPtrOutputWithContext(context.Background())
}

func (i *applicationGatewayProbeHealthResponseMatchPtrType) ToApplicationGatewayProbeHealthResponseMatchPtrOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeHealthResponseMatchPtrOutput)
}

type ApplicationGatewayProbeHealthResponseMatchOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeHealthResponseMatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbeHealthResponseMatch)(nil)).Elem()
}

func (o ApplicationGatewayProbeHealthResponseMatchOutput) ToApplicationGatewayProbeHealthResponseMatchOutput() ApplicationGatewayProbeHealthResponseMatchOutput {
	return o
}

func (o ApplicationGatewayProbeHealthResponseMatchOutput) ToApplicationGatewayProbeHealthResponseMatchOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchOutput {
	return o
}

func (o ApplicationGatewayProbeHealthResponseMatchOutput) ToApplicationGatewayProbeHealthResponseMatchPtrOutput() ApplicationGatewayProbeHealthResponseMatchPtrOutput {
	return o.ToApplicationGatewayProbeHealthResponseMatchPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayProbeHealthResponseMatchOutput) ToApplicationGatewayProbeHealthResponseMatchPtrOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayProbeHealthResponseMatch) *ApplicationGatewayProbeHealthResponseMatch {
		return &v
	}).(ApplicationGatewayProbeHealthResponseMatchPtrOutput)
}

func (o ApplicationGatewayProbeHealthResponseMatchOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeHealthResponseMatch) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeHealthResponseMatchOutput) StatusCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeHealthResponseMatch) []string { return v.StatusCodes }).(pulumi.StringArrayOutput)
}

type ApplicationGatewayProbeHealthResponseMatchPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeHealthResponseMatchPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayProbeHealthResponseMatch)(nil)).Elem()
}

func (o ApplicationGatewayProbeHealthResponseMatchPtrOutput) ToApplicationGatewayProbeHealthResponseMatchPtrOutput() ApplicationGatewayProbeHealthResponseMatchPtrOutput {
	return o
}

func (o ApplicationGatewayProbeHealthResponseMatchPtrOutput) ToApplicationGatewayProbeHealthResponseMatchPtrOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchPtrOutput {
	return o
}

func (o ApplicationGatewayProbeHealthResponseMatchPtrOutput) Elem() ApplicationGatewayProbeHealthResponseMatchOutput {
	return o.ApplyT(func(v *ApplicationGatewayProbeHealthResponseMatch) ApplicationGatewayProbeHealthResponseMatch {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayProbeHealthResponseMatch
		return ret
	}).(ApplicationGatewayProbeHealthResponseMatchOutput)
}

func (o ApplicationGatewayProbeHealthResponseMatchPtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayProbeHealthResponseMatch) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeHealthResponseMatchPtrOutput) StatusCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewayProbeHealthResponseMatch) []string {
		if v == nil {
			return nil
		}
		return v.StatusCodes
	}).(pulumi.StringArrayOutput)
}

type ApplicationGatewayProbeHealthResponseMatchResponse struct {
	Body        *string  `pulumi:"body"`
	StatusCodes []string `pulumi:"statusCodes"`
}





type ApplicationGatewayProbeHealthResponseMatchResponseInput interface {
	pulumi.Input

	ToApplicationGatewayProbeHealthResponseMatchResponseOutput() ApplicationGatewayProbeHealthResponseMatchResponseOutput
	ToApplicationGatewayProbeHealthResponseMatchResponseOutputWithContext(context.Context) ApplicationGatewayProbeHealthResponseMatchResponseOutput
}

type ApplicationGatewayProbeHealthResponseMatchResponseArgs struct {
	Body        pulumi.StringPtrInput   `pulumi:"body"`
	StatusCodes pulumi.StringArrayInput `pulumi:"statusCodes"`
}

func (ApplicationGatewayProbeHealthResponseMatchResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbeHealthResponseMatchResponse)(nil)).Elem()
}

func (i ApplicationGatewayProbeHealthResponseMatchResponseArgs) ToApplicationGatewayProbeHealthResponseMatchResponseOutput() ApplicationGatewayProbeHealthResponseMatchResponseOutput {
	return i.ToApplicationGatewayProbeHealthResponseMatchResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayProbeHealthResponseMatchResponseArgs) ToApplicationGatewayProbeHealthResponseMatchResponseOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeHealthResponseMatchResponseOutput)
}

func (i ApplicationGatewayProbeHealthResponseMatchResponseArgs) ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutput() ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput {
	return i.ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationGatewayProbeHealthResponseMatchResponseArgs) ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeHealthResponseMatchResponseOutput).ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutputWithContext(ctx)
}









type ApplicationGatewayProbeHealthResponseMatchResponsePtrInput interface {
	pulumi.Input

	ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutput() ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput
	ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutputWithContext(context.Context) ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput
}

type applicationGatewayProbeHealthResponseMatchResponsePtrType ApplicationGatewayProbeHealthResponseMatchResponseArgs

func ApplicationGatewayProbeHealthResponseMatchResponsePtr(v *ApplicationGatewayProbeHealthResponseMatchResponseArgs) ApplicationGatewayProbeHealthResponseMatchResponsePtrInput {
	return (*applicationGatewayProbeHealthResponseMatchResponsePtrType)(v)
}

func (*applicationGatewayProbeHealthResponseMatchResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayProbeHealthResponseMatchResponse)(nil)).Elem()
}

func (i *applicationGatewayProbeHealthResponseMatchResponsePtrType) ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutput() ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput {
	return i.ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutputWithContext(context.Background())
}

func (i *applicationGatewayProbeHealthResponseMatchResponsePtrType) ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput)
}

type ApplicationGatewayProbeHealthResponseMatchResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeHealthResponseMatchResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbeHealthResponseMatchResponse)(nil)).Elem()
}

func (o ApplicationGatewayProbeHealthResponseMatchResponseOutput) ToApplicationGatewayProbeHealthResponseMatchResponseOutput() ApplicationGatewayProbeHealthResponseMatchResponseOutput {
	return o
}

func (o ApplicationGatewayProbeHealthResponseMatchResponseOutput) ToApplicationGatewayProbeHealthResponseMatchResponseOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchResponseOutput {
	return o
}

func (o ApplicationGatewayProbeHealthResponseMatchResponseOutput) ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutput() ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput {
	return o.ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayProbeHealthResponseMatchResponseOutput) ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayProbeHealthResponseMatchResponse) *ApplicationGatewayProbeHealthResponseMatchResponse {
		return &v
	}).(ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput)
}

func (o ApplicationGatewayProbeHealthResponseMatchResponseOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeHealthResponseMatchResponse) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeHealthResponseMatchResponseOutput) StatusCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeHealthResponseMatchResponse) []string { return v.StatusCodes }).(pulumi.StringArrayOutput)
}

type ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayProbeHealthResponseMatchResponse)(nil)).Elem()
}

func (o ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput) ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutput() ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput {
	return o
}

func (o ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput) ToApplicationGatewayProbeHealthResponseMatchResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput {
	return o
}

func (o ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput) Elem() ApplicationGatewayProbeHealthResponseMatchResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewayProbeHealthResponseMatchResponse) ApplicationGatewayProbeHealthResponseMatchResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayProbeHealthResponseMatchResponse
		return ret
	}).(ApplicationGatewayProbeHealthResponseMatchResponseOutput)
}

func (o ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayProbeHealthResponseMatchResponse) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput) StatusCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewayProbeHealthResponseMatchResponse) []string {
		if v == nil {
			return nil
		}
		return v.StatusCodes
	}).(pulumi.StringArrayOutput)
}

type ApplicationGatewayProbeResponse struct {
	Etag                                *string                                             `pulumi:"etag"`
	Host                                *string                                             `pulumi:"host"`
	Id                                  *string                                             `pulumi:"id"`
	Interval                            *int                                                `pulumi:"interval"`
	Match                               *ApplicationGatewayProbeHealthResponseMatchResponse `pulumi:"match"`
	MinServers                          *int                                                `pulumi:"minServers"`
	Name                                *string                                             `pulumi:"name"`
	Path                                *string                                             `pulumi:"path"`
	PickHostNameFromBackendHttpSettings *bool                                               `pulumi:"pickHostNameFromBackendHttpSettings"`
	Protocol                            *string                                             `pulumi:"protocol"`
	ProvisioningState                   *string                                             `pulumi:"provisioningState"`
	Timeout                             *int                                                `pulumi:"timeout"`
	Type                                *string                                             `pulumi:"type"`
	UnhealthyThreshold                  *int                                                `pulumi:"unhealthyThreshold"`
}





type ApplicationGatewayProbeResponseInput interface {
	pulumi.Input

	ToApplicationGatewayProbeResponseOutput() ApplicationGatewayProbeResponseOutput
	ToApplicationGatewayProbeResponseOutputWithContext(context.Context) ApplicationGatewayProbeResponseOutput
}

type ApplicationGatewayProbeResponseArgs struct {
	Etag                                pulumi.StringPtrInput                                      `pulumi:"etag"`
	Host                                pulumi.StringPtrInput                                      `pulumi:"host"`
	Id                                  pulumi.StringPtrInput                                      `pulumi:"id"`
	Interval                            pulumi.IntPtrInput                                         `pulumi:"interval"`
	Match                               ApplicationGatewayProbeHealthResponseMatchResponsePtrInput `pulumi:"match"`
	MinServers                          pulumi.IntPtrInput                                         `pulumi:"minServers"`
	Name                                pulumi.StringPtrInput                                      `pulumi:"name"`
	Path                                pulumi.StringPtrInput                                      `pulumi:"path"`
	PickHostNameFromBackendHttpSettings pulumi.BoolPtrInput                                        `pulumi:"pickHostNameFromBackendHttpSettings"`
	Protocol                            pulumi.StringPtrInput                                      `pulumi:"protocol"`
	ProvisioningState                   pulumi.StringPtrInput                                      `pulumi:"provisioningState"`
	Timeout                             pulumi.IntPtrInput                                         `pulumi:"timeout"`
	Type                                pulumi.StringPtrInput                                      `pulumi:"type"`
	UnhealthyThreshold                  pulumi.IntPtrInput                                         `pulumi:"unhealthyThreshold"`
}

func (ApplicationGatewayProbeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbeResponse)(nil)).Elem()
}

func (i ApplicationGatewayProbeResponseArgs) ToApplicationGatewayProbeResponseOutput() ApplicationGatewayProbeResponseOutput {
	return i.ToApplicationGatewayProbeResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayProbeResponseArgs) ToApplicationGatewayProbeResponseOutputWithContext(ctx context.Context) ApplicationGatewayProbeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeResponseOutput)
}





type ApplicationGatewayProbeResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayProbeResponseArrayOutput() ApplicationGatewayProbeResponseArrayOutput
	ToApplicationGatewayProbeResponseArrayOutputWithContext(context.Context) ApplicationGatewayProbeResponseArrayOutput
}

type ApplicationGatewayProbeResponseArray []ApplicationGatewayProbeResponseInput

func (ApplicationGatewayProbeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayProbeResponse)(nil)).Elem()
}

func (i ApplicationGatewayProbeResponseArray) ToApplicationGatewayProbeResponseArrayOutput() ApplicationGatewayProbeResponseArrayOutput {
	return i.ToApplicationGatewayProbeResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayProbeResponseArray) ToApplicationGatewayProbeResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayProbeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeResponseArrayOutput)
}

type ApplicationGatewayProbeResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbeResponse)(nil)).Elem()
}

func (o ApplicationGatewayProbeResponseOutput) ToApplicationGatewayProbeResponseOutput() ApplicationGatewayProbeResponseOutput {
	return o
}

func (o ApplicationGatewayProbeResponseOutput) ToApplicationGatewayProbeResponseOutputWithContext(ctx context.Context) ApplicationGatewayProbeResponseOutput {
	return o
}

func (o ApplicationGatewayProbeResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Match() ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *ApplicationGatewayProbeHealthResponseMatchResponse {
		return v.Match
	}).(ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) MinServers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *int { return v.MinServers }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) PickHostNameFromBackendHttpSettings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *bool { return v.PickHostNameFromBackendHttpSettings }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) UnhealthyThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *int { return v.UnhealthyThreshold }).(pulumi.IntPtrOutput)
}

type ApplicationGatewayProbeResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayProbeResponse)(nil)).Elem()
}

func (o ApplicationGatewayProbeResponseArrayOutput) ToApplicationGatewayProbeResponseArrayOutput() ApplicationGatewayProbeResponseArrayOutput {
	return o
}

func (o ApplicationGatewayProbeResponseArrayOutput) ToApplicationGatewayProbeResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayProbeResponseArrayOutput {
	return o
}

func (o ApplicationGatewayProbeResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayProbeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayProbeResponse {
		return vs[0].([]ApplicationGatewayProbeResponse)[vs[1].(int)]
	}).(ApplicationGatewayProbeResponseOutput)
}

type ApplicationGatewayRedirectConfiguration struct {
	Etag                *string       `pulumi:"etag"`
	Id                  *string       `pulumi:"id"`
	IncludePath         *bool         `pulumi:"includePath"`
	IncludeQueryString  *bool         `pulumi:"includeQueryString"`
	Name                *string       `pulumi:"name"`
	PathRules           []SubResource `pulumi:"pathRules"`
	RedirectType        *string       `pulumi:"redirectType"`
	RequestRoutingRules []SubResource `pulumi:"requestRoutingRules"`
	TargetListener      *SubResource  `pulumi:"targetListener"`
	TargetUrl           *string       `pulumi:"targetUrl"`
	Type                *string       `pulumi:"type"`
	UrlPathMaps         []SubResource `pulumi:"urlPathMaps"`
}





type ApplicationGatewayRedirectConfigurationInput interface {
	pulumi.Input

	ToApplicationGatewayRedirectConfigurationOutput() ApplicationGatewayRedirectConfigurationOutput
	ToApplicationGatewayRedirectConfigurationOutputWithContext(context.Context) ApplicationGatewayRedirectConfigurationOutput
}

type ApplicationGatewayRedirectConfigurationArgs struct {
	Etag                pulumi.StringPtrInput `pulumi:"etag"`
	Id                  pulumi.StringPtrInput `pulumi:"id"`
	IncludePath         pulumi.BoolPtrInput   `pulumi:"includePath"`
	IncludeQueryString  pulumi.BoolPtrInput   `pulumi:"includeQueryString"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	PathRules           SubResourceArrayInput `pulumi:"pathRules"`
	RedirectType        pulumi.StringPtrInput `pulumi:"redirectType"`
	RequestRoutingRules SubResourceArrayInput `pulumi:"requestRoutingRules"`
	TargetListener      SubResourcePtrInput   `pulumi:"targetListener"`
	TargetUrl           pulumi.StringPtrInput `pulumi:"targetUrl"`
	Type                pulumi.StringPtrInput `pulumi:"type"`
	UrlPathMaps         SubResourceArrayInput `pulumi:"urlPathMaps"`
}

func (ApplicationGatewayRedirectConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRedirectConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayRedirectConfigurationArgs) ToApplicationGatewayRedirectConfigurationOutput() ApplicationGatewayRedirectConfigurationOutput {
	return i.ToApplicationGatewayRedirectConfigurationOutputWithContext(context.Background())
}

func (i ApplicationGatewayRedirectConfigurationArgs) ToApplicationGatewayRedirectConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayRedirectConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayRedirectConfigurationOutput)
}





type ApplicationGatewayRedirectConfigurationArrayInput interface {
	pulumi.Input

	ToApplicationGatewayRedirectConfigurationArrayOutput() ApplicationGatewayRedirectConfigurationArrayOutput
	ToApplicationGatewayRedirectConfigurationArrayOutputWithContext(context.Context) ApplicationGatewayRedirectConfigurationArrayOutput
}

type ApplicationGatewayRedirectConfigurationArray []ApplicationGatewayRedirectConfigurationInput

func (ApplicationGatewayRedirectConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRedirectConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayRedirectConfigurationArray) ToApplicationGatewayRedirectConfigurationArrayOutput() ApplicationGatewayRedirectConfigurationArrayOutput {
	return i.ToApplicationGatewayRedirectConfigurationArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayRedirectConfigurationArray) ToApplicationGatewayRedirectConfigurationArrayOutputWithContext(ctx context.Context) ApplicationGatewayRedirectConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayRedirectConfigurationArrayOutput)
}

type ApplicationGatewayRedirectConfigurationOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRedirectConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRedirectConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayRedirectConfigurationOutput) ToApplicationGatewayRedirectConfigurationOutput() ApplicationGatewayRedirectConfigurationOutput {
	return o
}

func (o ApplicationGatewayRedirectConfigurationOutput) ToApplicationGatewayRedirectConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayRedirectConfigurationOutput {
	return o
}

func (o ApplicationGatewayRedirectConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) IncludePath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) *bool { return v.IncludePath }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) IncludeQueryString() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) *bool { return v.IncludeQueryString }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) PathRules() SubResourceArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) []SubResource { return v.PathRules }).(SubResourceArrayOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) RedirectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) *string { return v.RedirectType }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) RequestRoutingRules() SubResourceArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) []SubResource { return v.RequestRoutingRules }).(SubResourceArrayOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) TargetListener() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) *SubResource { return v.TargetListener }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) TargetUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) *string { return v.TargetUrl }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationOutput) UrlPathMaps() SubResourceArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfiguration) []SubResource { return v.UrlPathMaps }).(SubResourceArrayOutput)
}

type ApplicationGatewayRedirectConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRedirectConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRedirectConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayRedirectConfigurationArrayOutput) ToApplicationGatewayRedirectConfigurationArrayOutput() ApplicationGatewayRedirectConfigurationArrayOutput {
	return o
}

func (o ApplicationGatewayRedirectConfigurationArrayOutput) ToApplicationGatewayRedirectConfigurationArrayOutputWithContext(ctx context.Context) ApplicationGatewayRedirectConfigurationArrayOutput {
	return o
}

func (o ApplicationGatewayRedirectConfigurationArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayRedirectConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayRedirectConfiguration {
		return vs[0].([]ApplicationGatewayRedirectConfiguration)[vs[1].(int)]
	}).(ApplicationGatewayRedirectConfigurationOutput)
}

type ApplicationGatewayRedirectConfigurationResponse struct {
	Etag                *string               `pulumi:"etag"`
	Id                  *string               `pulumi:"id"`
	IncludePath         *bool                 `pulumi:"includePath"`
	IncludeQueryString  *bool                 `pulumi:"includeQueryString"`
	Name                *string               `pulumi:"name"`
	PathRules           []SubResourceResponse `pulumi:"pathRules"`
	RedirectType        *string               `pulumi:"redirectType"`
	RequestRoutingRules []SubResourceResponse `pulumi:"requestRoutingRules"`
	TargetListener      *SubResourceResponse  `pulumi:"targetListener"`
	TargetUrl           *string               `pulumi:"targetUrl"`
	Type                *string               `pulumi:"type"`
	UrlPathMaps         []SubResourceResponse `pulumi:"urlPathMaps"`
}





type ApplicationGatewayRedirectConfigurationResponseInput interface {
	pulumi.Input

	ToApplicationGatewayRedirectConfigurationResponseOutput() ApplicationGatewayRedirectConfigurationResponseOutput
	ToApplicationGatewayRedirectConfigurationResponseOutputWithContext(context.Context) ApplicationGatewayRedirectConfigurationResponseOutput
}

type ApplicationGatewayRedirectConfigurationResponseArgs struct {
	Etag                pulumi.StringPtrInput         `pulumi:"etag"`
	Id                  pulumi.StringPtrInput         `pulumi:"id"`
	IncludePath         pulumi.BoolPtrInput           `pulumi:"includePath"`
	IncludeQueryString  pulumi.BoolPtrInput           `pulumi:"includeQueryString"`
	Name                pulumi.StringPtrInput         `pulumi:"name"`
	PathRules           SubResourceResponseArrayInput `pulumi:"pathRules"`
	RedirectType        pulumi.StringPtrInput         `pulumi:"redirectType"`
	RequestRoutingRules SubResourceResponseArrayInput `pulumi:"requestRoutingRules"`
	TargetListener      SubResourceResponsePtrInput   `pulumi:"targetListener"`
	TargetUrl           pulumi.StringPtrInput         `pulumi:"targetUrl"`
	Type                pulumi.StringPtrInput         `pulumi:"type"`
	UrlPathMaps         SubResourceResponseArrayInput `pulumi:"urlPathMaps"`
}

func (ApplicationGatewayRedirectConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRedirectConfigurationResponse)(nil)).Elem()
}

func (i ApplicationGatewayRedirectConfigurationResponseArgs) ToApplicationGatewayRedirectConfigurationResponseOutput() ApplicationGatewayRedirectConfigurationResponseOutput {
	return i.ToApplicationGatewayRedirectConfigurationResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayRedirectConfigurationResponseArgs) ToApplicationGatewayRedirectConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayRedirectConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayRedirectConfigurationResponseOutput)
}





type ApplicationGatewayRedirectConfigurationResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayRedirectConfigurationResponseArrayOutput() ApplicationGatewayRedirectConfigurationResponseArrayOutput
	ToApplicationGatewayRedirectConfigurationResponseArrayOutputWithContext(context.Context) ApplicationGatewayRedirectConfigurationResponseArrayOutput
}

type ApplicationGatewayRedirectConfigurationResponseArray []ApplicationGatewayRedirectConfigurationResponseInput

func (ApplicationGatewayRedirectConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRedirectConfigurationResponse)(nil)).Elem()
}

func (i ApplicationGatewayRedirectConfigurationResponseArray) ToApplicationGatewayRedirectConfigurationResponseArrayOutput() ApplicationGatewayRedirectConfigurationResponseArrayOutput {
	return i.ToApplicationGatewayRedirectConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayRedirectConfigurationResponseArray) ToApplicationGatewayRedirectConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayRedirectConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayRedirectConfigurationResponseArrayOutput)
}

type ApplicationGatewayRedirectConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRedirectConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRedirectConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) ToApplicationGatewayRedirectConfigurationResponseOutput() ApplicationGatewayRedirectConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) ToApplicationGatewayRedirectConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayRedirectConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) IncludePath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) *bool { return v.IncludePath }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) IncludeQueryString() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) *bool { return v.IncludeQueryString }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) PathRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) []SubResourceResponse { return v.PathRules }).(SubResourceResponseArrayOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) RedirectType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) *string { return v.RedirectType }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) RequestRoutingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) []SubResourceResponse {
		return v.RequestRoutingRules
	}).(SubResourceResponseArrayOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) TargetListener() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) *SubResourceResponse { return v.TargetListener }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) TargetUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) *string { return v.TargetUrl }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRedirectConfigurationResponseOutput) UrlPathMaps() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayRedirectConfigurationResponse) []SubResourceResponse { return v.UrlPathMaps }).(SubResourceResponseArrayOutput)
}

type ApplicationGatewayRedirectConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRedirectConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRedirectConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayRedirectConfigurationResponseArrayOutput) ToApplicationGatewayRedirectConfigurationResponseArrayOutput() ApplicationGatewayRedirectConfigurationResponseArrayOutput {
	return o
}

func (o ApplicationGatewayRedirectConfigurationResponseArrayOutput) ToApplicationGatewayRedirectConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayRedirectConfigurationResponseArrayOutput {
	return o
}

func (o ApplicationGatewayRedirectConfigurationResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayRedirectConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayRedirectConfigurationResponse {
		return vs[0].([]ApplicationGatewayRedirectConfigurationResponse)[vs[1].(int)]
	}).(ApplicationGatewayRedirectConfigurationResponseOutput)
}

type ApplicationGatewayRequestRoutingRule struct {
	BackendAddressPool    *SubResource `pulumi:"backendAddressPool"`
	BackendHttpSettings   *SubResource `pulumi:"backendHttpSettings"`
	Etag                  *string      `pulumi:"etag"`
	HttpListener          *SubResource `pulumi:"httpListener"`
	Id                    *string      `pulumi:"id"`
	Name                  *string      `pulumi:"name"`
	ProvisioningState     *string      `pulumi:"provisioningState"`
	RedirectConfiguration *SubResource `pulumi:"redirectConfiguration"`
	RuleType              *string      `pulumi:"ruleType"`
	Type                  *string      `pulumi:"type"`
	UrlPathMap            *SubResource `pulumi:"urlPathMap"`
}





type ApplicationGatewayRequestRoutingRuleInput interface {
	pulumi.Input

	ToApplicationGatewayRequestRoutingRuleOutput() ApplicationGatewayRequestRoutingRuleOutput
	ToApplicationGatewayRequestRoutingRuleOutputWithContext(context.Context) ApplicationGatewayRequestRoutingRuleOutput
}

type ApplicationGatewayRequestRoutingRuleArgs struct {
	BackendAddressPool    SubResourcePtrInput   `pulumi:"backendAddressPool"`
	BackendHttpSettings   SubResourcePtrInput   `pulumi:"backendHttpSettings"`
	Etag                  pulumi.StringPtrInput `pulumi:"etag"`
	HttpListener          SubResourcePtrInput   `pulumi:"httpListener"`
	Id                    pulumi.StringPtrInput `pulumi:"id"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState     pulumi.StringPtrInput `pulumi:"provisioningState"`
	RedirectConfiguration SubResourcePtrInput   `pulumi:"redirectConfiguration"`
	RuleType              pulumi.StringPtrInput `pulumi:"ruleType"`
	Type                  pulumi.StringPtrInput `pulumi:"type"`
	UrlPathMap            SubResourcePtrInput   `pulumi:"urlPathMap"`
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

func (o ApplicationGatewayRequestRoutingRuleOutput) RedirectConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *SubResource { return v.RedirectConfiguration }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) RuleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.RuleType }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) UrlPathMap() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *SubResource { return v.UrlPathMap }).(SubResourcePtrOutput)
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
	BackendAddressPool    *SubResourceResponse `pulumi:"backendAddressPool"`
	BackendHttpSettings   *SubResourceResponse `pulumi:"backendHttpSettings"`
	Etag                  *string              `pulumi:"etag"`
	HttpListener          *SubResourceResponse `pulumi:"httpListener"`
	Id                    *string              `pulumi:"id"`
	Name                  *string              `pulumi:"name"`
	ProvisioningState     *string              `pulumi:"provisioningState"`
	RedirectConfiguration *SubResourceResponse `pulumi:"redirectConfiguration"`
	RuleType              *string              `pulumi:"ruleType"`
	Type                  *string              `pulumi:"type"`
	UrlPathMap            *SubResourceResponse `pulumi:"urlPathMap"`
}





type ApplicationGatewayRequestRoutingRuleResponseInput interface {
	pulumi.Input

	ToApplicationGatewayRequestRoutingRuleResponseOutput() ApplicationGatewayRequestRoutingRuleResponseOutput
	ToApplicationGatewayRequestRoutingRuleResponseOutputWithContext(context.Context) ApplicationGatewayRequestRoutingRuleResponseOutput
}

type ApplicationGatewayRequestRoutingRuleResponseArgs struct {
	BackendAddressPool    SubResourceResponsePtrInput `pulumi:"backendAddressPool"`
	BackendHttpSettings   SubResourceResponsePtrInput `pulumi:"backendHttpSettings"`
	Etag                  pulumi.StringPtrInput       `pulumi:"etag"`
	HttpListener          SubResourceResponsePtrInput `pulumi:"httpListener"`
	Id                    pulumi.StringPtrInput       `pulumi:"id"`
	Name                  pulumi.StringPtrInput       `pulumi:"name"`
	ProvisioningState     pulumi.StringPtrInput       `pulumi:"provisioningState"`
	RedirectConfiguration SubResourceResponsePtrInput `pulumi:"redirectConfiguration"`
	RuleType              pulumi.StringPtrInput       `pulumi:"ruleType"`
	Type                  pulumi.StringPtrInput       `pulumi:"type"`
	UrlPathMap            SubResourceResponsePtrInput `pulumi:"urlPathMap"`
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

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) RedirectConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *SubResourceResponse {
		return v.RedirectConfiguration
	}).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) RuleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.RuleType }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) UrlPathMap() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *SubResourceResponse { return v.UrlPathMap }).(SubResourceResponsePtrOutput)
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
	Type              *string `pulumi:"type"`
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
	Type              pulumi.StringPtrInput `pulumi:"type"`
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

func (o ApplicationGatewaySslCertificateOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Type }).(pulumi.StringPtrOutput)
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
	Type              *string `pulumi:"type"`
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
	Type              pulumi.StringPtrInput `pulumi:"type"`
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

func (o ApplicationGatewaySslCertificateResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
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

type ApplicationGatewaySslPolicy struct {
	CipherSuites         []string `pulumi:"cipherSuites"`
	DisabledSslProtocols []string `pulumi:"disabledSslProtocols"`
	MinProtocolVersion   *string  `pulumi:"minProtocolVersion"`
	PolicyName           *string  `pulumi:"policyName"`
	PolicyType           *string  `pulumi:"policyType"`
}





type ApplicationGatewaySslPolicyInput interface {
	pulumi.Input

	ToApplicationGatewaySslPolicyOutput() ApplicationGatewaySslPolicyOutput
	ToApplicationGatewaySslPolicyOutputWithContext(context.Context) ApplicationGatewaySslPolicyOutput
}

type ApplicationGatewaySslPolicyArgs struct {
	CipherSuites         pulumi.StringArrayInput `pulumi:"cipherSuites"`
	DisabledSslProtocols pulumi.StringArrayInput `pulumi:"disabledSslProtocols"`
	MinProtocolVersion   pulumi.StringPtrInput   `pulumi:"minProtocolVersion"`
	PolicyName           pulumi.StringPtrInput   `pulumi:"policyName"`
	PolicyType           pulumi.StringPtrInput   `pulumi:"policyType"`
}

func (ApplicationGatewaySslPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicy)(nil)).Elem()
}

func (i ApplicationGatewaySslPolicyArgs) ToApplicationGatewaySslPolicyOutput() ApplicationGatewaySslPolicyOutput {
	return i.ToApplicationGatewaySslPolicyOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslPolicyArgs) ToApplicationGatewaySslPolicyOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslPolicyOutput)
}

func (i ApplicationGatewaySslPolicyArgs) ToApplicationGatewaySslPolicyPtrOutput() ApplicationGatewaySslPolicyPtrOutput {
	return i.ToApplicationGatewaySslPolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslPolicyArgs) ToApplicationGatewaySslPolicyPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslPolicyOutput).ToApplicationGatewaySslPolicyPtrOutputWithContext(ctx)
}









type ApplicationGatewaySslPolicyPtrInput interface {
	pulumi.Input

	ToApplicationGatewaySslPolicyPtrOutput() ApplicationGatewaySslPolicyPtrOutput
	ToApplicationGatewaySslPolicyPtrOutputWithContext(context.Context) ApplicationGatewaySslPolicyPtrOutput
}

type applicationGatewaySslPolicyPtrType ApplicationGatewaySslPolicyArgs

func ApplicationGatewaySslPolicyPtr(v *ApplicationGatewaySslPolicyArgs) ApplicationGatewaySslPolicyPtrInput {
	return (*applicationGatewaySslPolicyPtrType)(v)
}

func (*applicationGatewaySslPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslPolicy)(nil)).Elem()
}

func (i *applicationGatewaySslPolicyPtrType) ToApplicationGatewaySslPolicyPtrOutput() ApplicationGatewaySslPolicyPtrOutput {
	return i.ToApplicationGatewaySslPolicyPtrOutputWithContext(context.Background())
}

func (i *applicationGatewaySslPolicyPtrType) ToApplicationGatewaySslPolicyPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslPolicyPtrOutput)
}

type ApplicationGatewaySslPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicy)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyOutput) ToApplicationGatewaySslPolicyOutput() ApplicationGatewaySslPolicyOutput {
	return o
}

func (o ApplicationGatewaySslPolicyOutput) ToApplicationGatewaySslPolicyOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyOutput {
	return o
}

func (o ApplicationGatewaySslPolicyOutput) ToApplicationGatewaySslPolicyPtrOutput() ApplicationGatewaySslPolicyPtrOutput {
	return o.ToApplicationGatewaySslPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyOutput) ToApplicationGatewaySslPolicyPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySslPolicy) *ApplicationGatewaySslPolicy {
		return &v
	}).(ApplicationGatewaySslPolicyPtrOutput)
}

func (o ApplicationGatewaySslPolicyOutput) CipherSuites() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicy) []string { return v.CipherSuites }).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewaySslPolicyOutput) DisabledSslProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicy) []string { return v.DisabledSslProtocols }).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewaySslPolicyOutput) MinProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicy) *string { return v.MinProtocolVersion }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslPolicyOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicy) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslPolicyOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicy) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslPolicy)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyPtrOutput) ToApplicationGatewaySslPolicyPtrOutput() ApplicationGatewaySslPolicyPtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyPtrOutput) ToApplicationGatewaySslPolicyPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyPtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyPtrOutput) Elem() ApplicationGatewaySslPolicyOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicy) ApplicationGatewaySslPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySslPolicy
		return ret
	}).(ApplicationGatewaySslPolicyOutput)
}

func (o ApplicationGatewaySslPolicyPtrOutput) CipherSuites() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicy) []string {
		if v == nil {
			return nil
		}
		return v.CipherSuites
	}).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewaySslPolicyPtrOutput) DisabledSslProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicy) []string {
		if v == nil {
			return nil
		}
		return v.DisabledSslProtocols
	}).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewaySslPolicyPtrOutput) MinProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicy) *string {
		if v == nil {
			return nil
		}
		return v.MinProtocolVersion
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslPolicyPtrOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicy) *string {
		if v == nil {
			return nil
		}
		return v.PolicyName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslPolicyPtrOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicy) *string {
		if v == nil {
			return nil
		}
		return v.PolicyType
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslPolicyResponse struct {
	CipherSuites         []string `pulumi:"cipherSuites"`
	DisabledSslProtocols []string `pulumi:"disabledSslProtocols"`
	MinProtocolVersion   *string  `pulumi:"minProtocolVersion"`
	PolicyName           *string  `pulumi:"policyName"`
	PolicyType           *string  `pulumi:"policyType"`
}





type ApplicationGatewaySslPolicyResponseInput interface {
	pulumi.Input

	ToApplicationGatewaySslPolicyResponseOutput() ApplicationGatewaySslPolicyResponseOutput
	ToApplicationGatewaySslPolicyResponseOutputWithContext(context.Context) ApplicationGatewaySslPolicyResponseOutput
}

type ApplicationGatewaySslPolicyResponseArgs struct {
	CipherSuites         pulumi.StringArrayInput `pulumi:"cipherSuites"`
	DisabledSslProtocols pulumi.StringArrayInput `pulumi:"disabledSslProtocols"`
	MinProtocolVersion   pulumi.StringPtrInput   `pulumi:"minProtocolVersion"`
	PolicyName           pulumi.StringPtrInput   `pulumi:"policyName"`
	PolicyType           pulumi.StringPtrInput   `pulumi:"policyType"`
}

func (ApplicationGatewaySslPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicyResponse)(nil)).Elem()
}

func (i ApplicationGatewaySslPolicyResponseArgs) ToApplicationGatewaySslPolicyResponseOutput() ApplicationGatewaySslPolicyResponseOutput {
	return i.ToApplicationGatewaySslPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslPolicyResponseArgs) ToApplicationGatewaySslPolicyResponseOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslPolicyResponseOutput)
}

func (i ApplicationGatewaySslPolicyResponseArgs) ToApplicationGatewaySslPolicyResponsePtrOutput() ApplicationGatewaySslPolicyResponsePtrOutput {
	return i.ToApplicationGatewaySslPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslPolicyResponseArgs) ToApplicationGatewaySslPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslPolicyResponseOutput).ToApplicationGatewaySslPolicyResponsePtrOutputWithContext(ctx)
}









type ApplicationGatewaySslPolicyResponsePtrInput interface {
	pulumi.Input

	ToApplicationGatewaySslPolicyResponsePtrOutput() ApplicationGatewaySslPolicyResponsePtrOutput
	ToApplicationGatewaySslPolicyResponsePtrOutputWithContext(context.Context) ApplicationGatewaySslPolicyResponsePtrOutput
}

type applicationGatewaySslPolicyResponsePtrType ApplicationGatewaySslPolicyResponseArgs

func ApplicationGatewaySslPolicyResponsePtr(v *ApplicationGatewaySslPolicyResponseArgs) ApplicationGatewaySslPolicyResponsePtrInput {
	return (*applicationGatewaySslPolicyResponsePtrType)(v)
}

func (*applicationGatewaySslPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslPolicyResponse)(nil)).Elem()
}

func (i *applicationGatewaySslPolicyResponsePtrType) ToApplicationGatewaySslPolicyResponsePtrOutput() ApplicationGatewaySslPolicyResponsePtrOutput {
	return i.ToApplicationGatewaySslPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *applicationGatewaySslPolicyResponsePtrType) ToApplicationGatewaySslPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslPolicyResponsePtrOutput)
}

type ApplicationGatewaySslPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicyResponse)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyResponseOutput) ToApplicationGatewaySslPolicyResponseOutput() ApplicationGatewaySslPolicyResponseOutput {
	return o
}

func (o ApplicationGatewaySslPolicyResponseOutput) ToApplicationGatewaySslPolicyResponseOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyResponseOutput {
	return o
}

func (o ApplicationGatewaySslPolicyResponseOutput) ToApplicationGatewaySslPolicyResponsePtrOutput() ApplicationGatewaySslPolicyResponsePtrOutput {
	return o.ToApplicationGatewaySslPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyResponseOutput) ToApplicationGatewaySslPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySslPolicyResponse) *ApplicationGatewaySslPolicyResponse {
		return &v
	}).(ApplicationGatewaySslPolicyResponsePtrOutput)
}

func (o ApplicationGatewaySslPolicyResponseOutput) CipherSuites() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicyResponse) []string { return v.CipherSuites }).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewaySslPolicyResponseOutput) DisabledSslProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicyResponse) []string { return v.DisabledSslProtocols }).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewaySslPolicyResponseOutput) MinProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicyResponse) *string { return v.MinProtocolVersion }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslPolicyResponseOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicyResponse) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslPolicyResponseOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicyResponse) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslPolicyResponse)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) ToApplicationGatewaySslPolicyResponsePtrOutput() ApplicationGatewaySslPolicyResponsePtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) ToApplicationGatewaySslPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyResponsePtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) Elem() ApplicationGatewaySslPolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicyResponse) ApplicationGatewaySslPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySslPolicyResponse
		return ret
	}).(ApplicationGatewaySslPolicyResponseOutput)
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) CipherSuites() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicyResponse) []string {
		if v == nil {
			return nil
		}
		return v.CipherSuites
	}).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) DisabledSslProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicyResponse) []string {
		if v == nil {
			return nil
		}
		return v.DisabledSslProtocols
	}).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) MinProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.MinProtocolVersion
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyType
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayUrlPathMap struct {
	DefaultBackendAddressPool    *SubResource                 `pulumi:"defaultBackendAddressPool"`
	DefaultBackendHttpSettings   *SubResource                 `pulumi:"defaultBackendHttpSettings"`
	DefaultRedirectConfiguration *SubResource                 `pulumi:"defaultRedirectConfiguration"`
	Etag                         *string                      `pulumi:"etag"`
	Id                           *string                      `pulumi:"id"`
	Name                         *string                      `pulumi:"name"`
	PathRules                    []ApplicationGatewayPathRule `pulumi:"pathRules"`
	ProvisioningState            *string                      `pulumi:"provisioningState"`
	Type                         *string                      `pulumi:"type"`
}





type ApplicationGatewayUrlPathMapInput interface {
	pulumi.Input

	ToApplicationGatewayUrlPathMapOutput() ApplicationGatewayUrlPathMapOutput
	ToApplicationGatewayUrlPathMapOutputWithContext(context.Context) ApplicationGatewayUrlPathMapOutput
}

type ApplicationGatewayUrlPathMapArgs struct {
	DefaultBackendAddressPool    SubResourcePtrInput                  `pulumi:"defaultBackendAddressPool"`
	DefaultBackendHttpSettings   SubResourcePtrInput                  `pulumi:"defaultBackendHttpSettings"`
	DefaultRedirectConfiguration SubResourcePtrInput                  `pulumi:"defaultRedirectConfiguration"`
	Etag                         pulumi.StringPtrInput                `pulumi:"etag"`
	Id                           pulumi.StringPtrInput                `pulumi:"id"`
	Name                         pulumi.StringPtrInput                `pulumi:"name"`
	PathRules                    ApplicationGatewayPathRuleArrayInput `pulumi:"pathRules"`
	ProvisioningState            pulumi.StringPtrInput                `pulumi:"provisioningState"`
	Type                         pulumi.StringPtrInput                `pulumi:"type"`
}

func (ApplicationGatewayUrlPathMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayUrlPathMap)(nil)).Elem()
}

func (i ApplicationGatewayUrlPathMapArgs) ToApplicationGatewayUrlPathMapOutput() ApplicationGatewayUrlPathMapOutput {
	return i.ToApplicationGatewayUrlPathMapOutputWithContext(context.Background())
}

func (i ApplicationGatewayUrlPathMapArgs) ToApplicationGatewayUrlPathMapOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayUrlPathMapOutput)
}





type ApplicationGatewayUrlPathMapArrayInput interface {
	pulumi.Input

	ToApplicationGatewayUrlPathMapArrayOutput() ApplicationGatewayUrlPathMapArrayOutput
	ToApplicationGatewayUrlPathMapArrayOutputWithContext(context.Context) ApplicationGatewayUrlPathMapArrayOutput
}

type ApplicationGatewayUrlPathMapArray []ApplicationGatewayUrlPathMapInput

func (ApplicationGatewayUrlPathMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayUrlPathMap)(nil)).Elem()
}

func (i ApplicationGatewayUrlPathMapArray) ToApplicationGatewayUrlPathMapArrayOutput() ApplicationGatewayUrlPathMapArrayOutput {
	return i.ToApplicationGatewayUrlPathMapArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayUrlPathMapArray) ToApplicationGatewayUrlPathMapArrayOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayUrlPathMapArrayOutput)
}

type ApplicationGatewayUrlPathMapOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayUrlPathMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayUrlPathMap)(nil)).Elem()
}

func (o ApplicationGatewayUrlPathMapOutput) ToApplicationGatewayUrlPathMapOutput() ApplicationGatewayUrlPathMapOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapOutput) ToApplicationGatewayUrlPathMapOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapOutput) DefaultBackendAddressPool() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *SubResource { return v.DefaultBackendAddressPool }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) DefaultBackendHttpSettings() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *SubResource { return v.DefaultBackendHttpSettings }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) DefaultRedirectConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *SubResource { return v.DefaultRedirectConfiguration }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) PathRules() ApplicationGatewayPathRuleArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) []ApplicationGatewayPathRule { return v.PathRules }).(ApplicationGatewayPathRuleArrayOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayUrlPathMapArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayUrlPathMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayUrlPathMap)(nil)).Elem()
}

func (o ApplicationGatewayUrlPathMapArrayOutput) ToApplicationGatewayUrlPathMapArrayOutput() ApplicationGatewayUrlPathMapArrayOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapArrayOutput) ToApplicationGatewayUrlPathMapArrayOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapArrayOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayUrlPathMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayUrlPathMap {
		return vs[0].([]ApplicationGatewayUrlPathMap)[vs[1].(int)]
	}).(ApplicationGatewayUrlPathMapOutput)
}

type ApplicationGatewayUrlPathMapResponse struct {
	DefaultBackendAddressPool    *SubResourceResponse                 `pulumi:"defaultBackendAddressPool"`
	DefaultBackendHttpSettings   *SubResourceResponse                 `pulumi:"defaultBackendHttpSettings"`
	DefaultRedirectConfiguration *SubResourceResponse                 `pulumi:"defaultRedirectConfiguration"`
	Etag                         *string                              `pulumi:"etag"`
	Id                           *string                              `pulumi:"id"`
	Name                         *string                              `pulumi:"name"`
	PathRules                    []ApplicationGatewayPathRuleResponse `pulumi:"pathRules"`
	ProvisioningState            *string                              `pulumi:"provisioningState"`
	Type                         *string                              `pulumi:"type"`
}





type ApplicationGatewayUrlPathMapResponseInput interface {
	pulumi.Input

	ToApplicationGatewayUrlPathMapResponseOutput() ApplicationGatewayUrlPathMapResponseOutput
	ToApplicationGatewayUrlPathMapResponseOutputWithContext(context.Context) ApplicationGatewayUrlPathMapResponseOutput
}

type ApplicationGatewayUrlPathMapResponseArgs struct {
	DefaultBackendAddressPool    SubResourceResponsePtrInput                  `pulumi:"defaultBackendAddressPool"`
	DefaultBackendHttpSettings   SubResourceResponsePtrInput                  `pulumi:"defaultBackendHttpSettings"`
	DefaultRedirectConfiguration SubResourceResponsePtrInput                  `pulumi:"defaultRedirectConfiguration"`
	Etag                         pulumi.StringPtrInput                        `pulumi:"etag"`
	Id                           pulumi.StringPtrInput                        `pulumi:"id"`
	Name                         pulumi.StringPtrInput                        `pulumi:"name"`
	PathRules                    ApplicationGatewayPathRuleResponseArrayInput `pulumi:"pathRules"`
	ProvisioningState            pulumi.StringPtrInput                        `pulumi:"provisioningState"`
	Type                         pulumi.StringPtrInput                        `pulumi:"type"`
}

func (ApplicationGatewayUrlPathMapResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayUrlPathMapResponse)(nil)).Elem()
}

func (i ApplicationGatewayUrlPathMapResponseArgs) ToApplicationGatewayUrlPathMapResponseOutput() ApplicationGatewayUrlPathMapResponseOutput {
	return i.ToApplicationGatewayUrlPathMapResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayUrlPathMapResponseArgs) ToApplicationGatewayUrlPathMapResponseOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayUrlPathMapResponseOutput)
}





type ApplicationGatewayUrlPathMapResponseArrayInput interface {
	pulumi.Input

	ToApplicationGatewayUrlPathMapResponseArrayOutput() ApplicationGatewayUrlPathMapResponseArrayOutput
	ToApplicationGatewayUrlPathMapResponseArrayOutputWithContext(context.Context) ApplicationGatewayUrlPathMapResponseArrayOutput
}

type ApplicationGatewayUrlPathMapResponseArray []ApplicationGatewayUrlPathMapResponseInput

func (ApplicationGatewayUrlPathMapResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayUrlPathMapResponse)(nil)).Elem()
}

func (i ApplicationGatewayUrlPathMapResponseArray) ToApplicationGatewayUrlPathMapResponseArrayOutput() ApplicationGatewayUrlPathMapResponseArrayOutput {
	return i.ToApplicationGatewayUrlPathMapResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayUrlPathMapResponseArray) ToApplicationGatewayUrlPathMapResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayUrlPathMapResponseArrayOutput)
}

type ApplicationGatewayUrlPathMapResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayUrlPathMapResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayUrlPathMapResponse)(nil)).Elem()
}

func (o ApplicationGatewayUrlPathMapResponseOutput) ToApplicationGatewayUrlPathMapResponseOutput() ApplicationGatewayUrlPathMapResponseOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapResponseOutput) ToApplicationGatewayUrlPathMapResponseOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapResponseOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapResponseOutput) DefaultBackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *SubResourceResponse { return v.DefaultBackendAddressPool }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) DefaultBackendHttpSettings() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *SubResourceResponse { return v.DefaultBackendHttpSettings }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) DefaultRedirectConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *SubResourceResponse {
		return v.DefaultRedirectConfiguration
	}).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) PathRules() ApplicationGatewayPathRuleResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) []ApplicationGatewayPathRuleResponse { return v.PathRules }).(ApplicationGatewayPathRuleResponseArrayOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayUrlPathMapResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayUrlPathMapResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayUrlPathMapResponse)(nil)).Elem()
}

func (o ApplicationGatewayUrlPathMapResponseArrayOutput) ToApplicationGatewayUrlPathMapResponseArrayOutput() ApplicationGatewayUrlPathMapResponseArrayOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapResponseArrayOutput) ToApplicationGatewayUrlPathMapResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapResponseArrayOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayUrlPathMapResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayUrlPathMapResponse {
		return vs[0].([]ApplicationGatewayUrlPathMapResponse)[vs[1].(int)]
	}).(ApplicationGatewayUrlPathMapResponseOutput)
}

type ApplicationGatewayWebApplicationFirewallConfiguration struct {
	DisabledRuleGroups []ApplicationGatewayFirewallDisabledRuleGroup `pulumi:"disabledRuleGroups"`
	Enabled            bool                                          `pulumi:"enabled"`
	FirewallMode       string                                        `pulumi:"firewallMode"`
	RuleSetType        string                                        `pulumi:"ruleSetType"`
	RuleSetVersion     string                                        `pulumi:"ruleSetVersion"`
}





type ApplicationGatewayWebApplicationFirewallConfigurationInput interface {
	pulumi.Input

	ToApplicationGatewayWebApplicationFirewallConfigurationOutput() ApplicationGatewayWebApplicationFirewallConfigurationOutput
	ToApplicationGatewayWebApplicationFirewallConfigurationOutputWithContext(context.Context) ApplicationGatewayWebApplicationFirewallConfigurationOutput
}

type ApplicationGatewayWebApplicationFirewallConfigurationArgs struct {
	DisabledRuleGroups ApplicationGatewayFirewallDisabledRuleGroupArrayInput `pulumi:"disabledRuleGroups"`
	Enabled            pulumi.BoolInput                                      `pulumi:"enabled"`
	FirewallMode       pulumi.StringInput                                    `pulumi:"firewallMode"`
	RuleSetType        pulumi.StringInput                                    `pulumi:"ruleSetType"`
	RuleSetVersion     pulumi.StringInput                                    `pulumi:"ruleSetVersion"`
}

func (ApplicationGatewayWebApplicationFirewallConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayWebApplicationFirewallConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationArgs) ToApplicationGatewayWebApplicationFirewallConfigurationOutput() ApplicationGatewayWebApplicationFirewallConfigurationOutput {
	return i.ToApplicationGatewayWebApplicationFirewallConfigurationOutputWithContext(context.Background())
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationArgs) ToApplicationGatewayWebApplicationFirewallConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayWebApplicationFirewallConfigurationOutput)
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationArgs) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return i.ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(context.Background())
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationArgs) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayWebApplicationFirewallConfigurationOutput).ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(ctx)
}









type ApplicationGatewayWebApplicationFirewallConfigurationPtrInput interface {
	pulumi.Input

	ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput
	ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(context.Context) ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput
}

type applicationGatewayWebApplicationFirewallConfigurationPtrType ApplicationGatewayWebApplicationFirewallConfigurationArgs

func ApplicationGatewayWebApplicationFirewallConfigurationPtr(v *ApplicationGatewayWebApplicationFirewallConfigurationArgs) ApplicationGatewayWebApplicationFirewallConfigurationPtrInput {
	return (*applicationGatewayWebApplicationFirewallConfigurationPtrType)(v)
}

func (*applicationGatewayWebApplicationFirewallConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayWebApplicationFirewallConfiguration)(nil)).Elem()
}

func (i *applicationGatewayWebApplicationFirewallConfigurationPtrType) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return i.ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(context.Background())
}

func (i *applicationGatewayWebApplicationFirewallConfigurationPtrType) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput)
}

type ApplicationGatewayWebApplicationFirewallConfigurationOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayWebApplicationFirewallConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayWebApplicationFirewallConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) ToApplicationGatewayWebApplicationFirewallConfigurationOutput() ApplicationGatewayWebApplicationFirewallConfigurationOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) ToApplicationGatewayWebApplicationFirewallConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return o.ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayWebApplicationFirewallConfiguration) *ApplicationGatewayWebApplicationFirewallConfiguration {
		return &v
	}).(ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) DisabledRuleGroups() ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfiguration) []ApplicationGatewayFirewallDisabledRuleGroup {
		return v.DisabledRuleGroups
	}).(ApplicationGatewayFirewallDisabledRuleGroupArrayOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfiguration) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) FirewallMode() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfiguration) string { return v.FirewallMode }).(pulumi.StringOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) RuleSetType() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfiguration) string { return v.RuleSetType }).(pulumi.StringOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) RuleSetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfiguration) string { return v.RuleSetVersion }).(pulumi.StringOutput)
}

type ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayWebApplicationFirewallConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) Elem() ApplicationGatewayWebApplicationFirewallConfigurationOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) ApplicationGatewayWebApplicationFirewallConfiguration {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayWebApplicationFirewallConfiguration
		return ret
	}).(ApplicationGatewayWebApplicationFirewallConfigurationOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) DisabledRuleGroups() ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) []ApplicationGatewayFirewallDisabledRuleGroup {
		if v == nil {
			return nil
		}
		return v.DisabledRuleGroups
	}).(ApplicationGatewayFirewallDisabledRuleGroupArrayOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) FirewallMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.FirewallMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) RuleSetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.RuleSetType
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) RuleSetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.RuleSetVersion
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayWebApplicationFirewallConfigurationResponse struct {
	DisabledRuleGroups []ApplicationGatewayFirewallDisabledRuleGroupResponse `pulumi:"disabledRuleGroups"`
	Enabled            bool                                                  `pulumi:"enabled"`
	FirewallMode       string                                                `pulumi:"firewallMode"`
	RuleSetType        string                                                `pulumi:"ruleSetType"`
	RuleSetVersion     string                                                `pulumi:"ruleSetVersion"`
}





type ApplicationGatewayWebApplicationFirewallConfigurationResponseInput interface {
	pulumi.Input

	ToApplicationGatewayWebApplicationFirewallConfigurationResponseOutput() ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput
	ToApplicationGatewayWebApplicationFirewallConfigurationResponseOutputWithContext(context.Context) ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput
}

type ApplicationGatewayWebApplicationFirewallConfigurationResponseArgs struct {
	DisabledRuleGroups ApplicationGatewayFirewallDisabledRuleGroupResponseArrayInput `pulumi:"disabledRuleGroups"`
	Enabled            pulumi.BoolInput                                              `pulumi:"enabled"`
	FirewallMode       pulumi.StringInput                                            `pulumi:"firewallMode"`
	RuleSetType        pulumi.StringInput                                            `pulumi:"ruleSetType"`
	RuleSetVersion     pulumi.StringInput                                            `pulumi:"ruleSetVersion"`
}

func (ApplicationGatewayWebApplicationFirewallConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayWebApplicationFirewallConfigurationResponse)(nil)).Elem()
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationResponseArgs) ToApplicationGatewayWebApplicationFirewallConfigurationResponseOutput() ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput {
	return i.ToApplicationGatewayWebApplicationFirewallConfigurationResponseOutputWithContext(context.Background())
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationResponseArgs) ToApplicationGatewayWebApplicationFirewallConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput)
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationResponseArgs) ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return i.ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationResponseArgs) ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput).ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutputWithContext(ctx)
}









type ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrInput interface {
	pulumi.Input

	ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput
	ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutputWithContext(context.Context) ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput
}

type applicationGatewayWebApplicationFirewallConfigurationResponsePtrType ApplicationGatewayWebApplicationFirewallConfigurationResponseArgs

func ApplicationGatewayWebApplicationFirewallConfigurationResponsePtr(v *ApplicationGatewayWebApplicationFirewallConfigurationResponseArgs) ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrInput {
	return (*applicationGatewayWebApplicationFirewallConfigurationResponsePtrType)(v)
}

func (*applicationGatewayWebApplicationFirewallConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayWebApplicationFirewallConfigurationResponse)(nil)).Elem()
}

func (i *applicationGatewayWebApplicationFirewallConfigurationResponsePtrType) ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return i.ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *applicationGatewayWebApplicationFirewallConfigurationResponsePtrType) ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput)
}

type ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayWebApplicationFirewallConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) ToApplicationGatewayWebApplicationFirewallConfigurationResponseOutput() ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) ToApplicationGatewayWebApplicationFirewallConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return o.ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayWebApplicationFirewallConfigurationResponse) *ApplicationGatewayWebApplicationFirewallConfigurationResponse {
		return &v
	}).(ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) DisabledRuleGroups() ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfigurationResponse) []ApplicationGatewayFirewallDisabledRuleGroupResponse {
		return v.DisabledRuleGroups
	}).(ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfigurationResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) FirewallMode() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfigurationResponse) string { return v.FirewallMode }).(pulumi.StringOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) RuleSetType() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfigurationResponse) string { return v.RuleSetType }).(pulumi.StringOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) RuleSetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfigurationResponse) string { return v.RuleSetVersion }).(pulumi.StringOutput)
}

type ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayWebApplicationFirewallConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) Elem() ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) ApplicationGatewayWebApplicationFirewallConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayWebApplicationFirewallConfigurationResponse
		return ret
	}).(ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) DisabledRuleGroups() ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) []ApplicationGatewayFirewallDisabledRuleGroupResponse {
		if v == nil {
			return nil
		}
		return v.DisabledRuleGroups
	}).(ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) FirewallMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FirewallMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) RuleSetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RuleSetType
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) RuleSetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RuleSetVersion
	}).(pulumi.StringPtrOutput)
}

type ApplicationSecurityGroupType struct {
	Id       *string           `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Tags     map[string]string `pulumi:"tags"`
}





type ApplicationSecurityGroupTypeInput interface {
	pulumi.Input

	ToApplicationSecurityGroupTypeOutput() ApplicationSecurityGroupTypeOutput
	ToApplicationSecurityGroupTypeOutputWithContext(context.Context) ApplicationSecurityGroupTypeOutput
}

type ApplicationSecurityGroupTypeArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	Tags     pulumi.StringMapInput `pulumi:"tags"`
}

func (ApplicationSecurityGroupTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationSecurityGroupType)(nil)).Elem()
}

func (i ApplicationSecurityGroupTypeArgs) ToApplicationSecurityGroupTypeOutput() ApplicationSecurityGroupTypeOutput {
	return i.ToApplicationSecurityGroupTypeOutputWithContext(context.Background())
}

func (i ApplicationSecurityGroupTypeArgs) ToApplicationSecurityGroupTypeOutputWithContext(ctx context.Context) ApplicationSecurityGroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSecurityGroupTypeOutput)
}





type ApplicationSecurityGroupTypeArrayInput interface {
	pulumi.Input

	ToApplicationSecurityGroupTypeArrayOutput() ApplicationSecurityGroupTypeArrayOutput
	ToApplicationSecurityGroupTypeArrayOutputWithContext(context.Context) ApplicationSecurityGroupTypeArrayOutput
}

type ApplicationSecurityGroupTypeArray []ApplicationSecurityGroupTypeInput

func (ApplicationSecurityGroupTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationSecurityGroupType)(nil)).Elem()
}

func (i ApplicationSecurityGroupTypeArray) ToApplicationSecurityGroupTypeArrayOutput() ApplicationSecurityGroupTypeArrayOutput {
	return i.ToApplicationSecurityGroupTypeArrayOutputWithContext(context.Background())
}

func (i ApplicationSecurityGroupTypeArray) ToApplicationSecurityGroupTypeArrayOutputWithContext(ctx context.Context) ApplicationSecurityGroupTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSecurityGroupTypeArrayOutput)
}

type ApplicationSecurityGroupTypeOutput struct{ *pulumi.OutputState }

func (ApplicationSecurityGroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationSecurityGroupType)(nil)).Elem()
}

func (o ApplicationSecurityGroupTypeOutput) ToApplicationSecurityGroupTypeOutput() ApplicationSecurityGroupTypeOutput {
	return o
}

func (o ApplicationSecurityGroupTypeOutput) ToApplicationSecurityGroupTypeOutputWithContext(ctx context.Context) ApplicationSecurityGroupTypeOutput {
	return o
}

func (o ApplicationSecurityGroupTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationSecurityGroupTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationSecurityGroupTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ApplicationSecurityGroupTypeArrayOutput struct{ *pulumi.OutputState }

func (ApplicationSecurityGroupTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationSecurityGroupType)(nil)).Elem()
}

func (o ApplicationSecurityGroupTypeArrayOutput) ToApplicationSecurityGroupTypeArrayOutput() ApplicationSecurityGroupTypeArrayOutput {
	return o
}

func (o ApplicationSecurityGroupTypeArrayOutput) ToApplicationSecurityGroupTypeArrayOutputWithContext(ctx context.Context) ApplicationSecurityGroupTypeArrayOutput {
	return o
}

func (o ApplicationSecurityGroupTypeArrayOutput) Index(i pulumi.IntInput) ApplicationSecurityGroupTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationSecurityGroupType {
		return vs[0].([]ApplicationSecurityGroupType)[vs[1].(int)]
	}).(ApplicationSecurityGroupTypeOutput)
}

type ApplicationSecurityGroupResponse struct {
	Etag              string            `pulumi:"etag"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	ResourceGuid      string            `pulumi:"resourceGuid"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}





type ApplicationSecurityGroupResponseInput interface {
	pulumi.Input

	ToApplicationSecurityGroupResponseOutput() ApplicationSecurityGroupResponseOutput
	ToApplicationSecurityGroupResponseOutputWithContext(context.Context) ApplicationSecurityGroupResponseOutput
}

type ApplicationSecurityGroupResponseArgs struct {
	Etag              pulumi.StringInput    `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Location          pulumi.StringPtrInput `pulumi:"location"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ProvisioningState pulumi.StringInput    `pulumi:"provisioningState"`
	ResourceGuid      pulumi.StringInput    `pulumi:"resourceGuid"`
	Tags              pulumi.StringMapInput `pulumi:"tags"`
	Type              pulumi.StringInput    `pulumi:"type"`
}

func (ApplicationSecurityGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationSecurityGroupResponse)(nil)).Elem()
}

func (i ApplicationSecurityGroupResponseArgs) ToApplicationSecurityGroupResponseOutput() ApplicationSecurityGroupResponseOutput {
	return i.ToApplicationSecurityGroupResponseOutputWithContext(context.Background())
}

func (i ApplicationSecurityGroupResponseArgs) ToApplicationSecurityGroupResponseOutputWithContext(ctx context.Context) ApplicationSecurityGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSecurityGroupResponseOutput)
}





type ApplicationSecurityGroupResponseArrayInput interface {
	pulumi.Input

	ToApplicationSecurityGroupResponseArrayOutput() ApplicationSecurityGroupResponseArrayOutput
	ToApplicationSecurityGroupResponseArrayOutputWithContext(context.Context) ApplicationSecurityGroupResponseArrayOutput
}

type ApplicationSecurityGroupResponseArray []ApplicationSecurityGroupResponseInput

func (ApplicationSecurityGroupResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationSecurityGroupResponse)(nil)).Elem()
}

func (i ApplicationSecurityGroupResponseArray) ToApplicationSecurityGroupResponseArrayOutput() ApplicationSecurityGroupResponseArrayOutput {
	return i.ToApplicationSecurityGroupResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationSecurityGroupResponseArray) ToApplicationSecurityGroupResponseArrayOutputWithContext(ctx context.Context) ApplicationSecurityGroupResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSecurityGroupResponseArrayOutput)
}

type ApplicationSecurityGroupResponseOutput struct{ *pulumi.OutputState }

func (ApplicationSecurityGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationSecurityGroupResponse)(nil)).Elem()
}

func (o ApplicationSecurityGroupResponseOutput) ToApplicationSecurityGroupResponseOutput() ApplicationSecurityGroupResponseOutput {
	return o
}

func (o ApplicationSecurityGroupResponseOutput) ToApplicationSecurityGroupResponseOutputWithContext(ctx context.Context) ApplicationSecurityGroupResponseOutput {
	return o
}

func (o ApplicationSecurityGroupResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ApplicationSecurityGroupResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationSecurityGroupResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationSecurityGroupResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationSecurityGroupResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApplicationSecurityGroupResponseOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupResponse) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o ApplicationSecurityGroupResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationSecurityGroupResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationSecurityGroupResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ApplicationSecurityGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationSecurityGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationSecurityGroupResponse)(nil)).Elem()
}

func (o ApplicationSecurityGroupResponseArrayOutput) ToApplicationSecurityGroupResponseArrayOutput() ApplicationSecurityGroupResponseArrayOutput {
	return o
}

func (o ApplicationSecurityGroupResponseArrayOutput) ToApplicationSecurityGroupResponseArrayOutputWithContext(ctx context.Context) ApplicationSecurityGroupResponseArrayOutput {
	return o
}

func (o ApplicationSecurityGroupResponseArrayOutput) Index(i pulumi.IntInput) ApplicationSecurityGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationSecurityGroupResponse {
		return vs[0].([]ApplicationSecurityGroupResponse)[vs[1].(int)]
	}).(ApplicationSecurityGroupResponseOutput)
}

type BackendAddressPool struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
}





type BackendAddressPoolInput interface {
	pulumi.Input

	ToBackendAddressPoolOutput() BackendAddressPoolOutput
	ToBackendAddressPoolOutputWithContext(context.Context) BackendAddressPoolOutput
}

type BackendAddressPoolArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
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

func (o BackendAddressPoolOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.Name }).(pulumi.StringPtrOutput)
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
	BackendIPConfigurations []NetworkInterfaceIPConfigurationResponse `pulumi:"backendIPConfigurations"`
	Etag                    *string                                   `pulumi:"etag"`
	Id                      *string                                   `pulumi:"id"`
	LoadBalancingRules      []SubResourceResponse                     `pulumi:"loadBalancingRules"`
	Name                    *string                                   `pulumi:"name"`
	OutboundNatRule         SubResourceResponse                       `pulumi:"outboundNatRule"`
	ProvisioningState       *string                                   `pulumi:"provisioningState"`
}





type BackendAddressPoolResponseInput interface {
	pulumi.Input

	ToBackendAddressPoolResponseOutput() BackendAddressPoolResponseOutput
	ToBackendAddressPoolResponseOutputWithContext(context.Context) BackendAddressPoolResponseOutput
}

type BackendAddressPoolResponseArgs struct {
	BackendIPConfigurations NetworkInterfaceIPConfigurationResponseArrayInput `pulumi:"backendIPConfigurations"`
	Etag                    pulumi.StringPtrInput                             `pulumi:"etag"`
	Id                      pulumi.StringPtrInput                             `pulumi:"id"`
	LoadBalancingRules      SubResourceResponseArrayInput                     `pulumi:"loadBalancingRules"`
	Name                    pulumi.StringPtrInput                             `pulumi:"name"`
	OutboundNatRule         SubResourceResponseInput                          `pulumi:"outboundNatRule"`
	ProvisioningState       pulumi.StringPtrInput                             `pulumi:"provisioningState"`
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

func (o BackendAddressPoolResponseOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) []NetworkInterfaceIPConfigurationResponse {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
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

func (o BackendAddressPoolResponseOutput) OutboundNatRule() SubResourceResponseOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) SubResourceResponse { return v.OutboundNatRule }).(SubResourceResponseOutput)
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

type BgpPeerStatusResponse struct {
	Asn               int     `pulumi:"asn"`
	ConnectedDuration string  `pulumi:"connectedDuration"`
	LocalAddress      string  `pulumi:"localAddress"`
	MessagesReceived  float64 `pulumi:"messagesReceived"`
	MessagesSent      float64 `pulumi:"messagesSent"`
	Neighbor          string  `pulumi:"neighbor"`
	RoutesReceived    float64 `pulumi:"routesReceived"`
	State             string  `pulumi:"state"`
}





type BgpPeerStatusResponseInput interface {
	pulumi.Input

	ToBgpPeerStatusResponseOutput() BgpPeerStatusResponseOutput
	ToBgpPeerStatusResponseOutputWithContext(context.Context) BgpPeerStatusResponseOutput
}

type BgpPeerStatusResponseArgs struct {
	Asn               pulumi.IntInput     `pulumi:"asn"`
	ConnectedDuration pulumi.StringInput  `pulumi:"connectedDuration"`
	LocalAddress      pulumi.StringInput  `pulumi:"localAddress"`
	MessagesReceived  pulumi.Float64Input `pulumi:"messagesReceived"`
	MessagesSent      pulumi.Float64Input `pulumi:"messagesSent"`
	Neighbor          pulumi.StringInput  `pulumi:"neighbor"`
	RoutesReceived    pulumi.Float64Input `pulumi:"routesReceived"`
	State             pulumi.StringInput  `pulumi:"state"`
}

func (BgpPeerStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpPeerStatusResponse)(nil)).Elem()
}

func (i BgpPeerStatusResponseArgs) ToBgpPeerStatusResponseOutput() BgpPeerStatusResponseOutput {
	return i.ToBgpPeerStatusResponseOutputWithContext(context.Background())
}

func (i BgpPeerStatusResponseArgs) ToBgpPeerStatusResponseOutputWithContext(ctx context.Context) BgpPeerStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerStatusResponseOutput)
}





type BgpPeerStatusResponseArrayInput interface {
	pulumi.Input

	ToBgpPeerStatusResponseArrayOutput() BgpPeerStatusResponseArrayOutput
	ToBgpPeerStatusResponseArrayOutputWithContext(context.Context) BgpPeerStatusResponseArrayOutput
}

type BgpPeerStatusResponseArray []BgpPeerStatusResponseInput

func (BgpPeerStatusResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BgpPeerStatusResponse)(nil)).Elem()
}

func (i BgpPeerStatusResponseArray) ToBgpPeerStatusResponseArrayOutput() BgpPeerStatusResponseArrayOutput {
	return i.ToBgpPeerStatusResponseArrayOutputWithContext(context.Background())
}

func (i BgpPeerStatusResponseArray) ToBgpPeerStatusResponseArrayOutputWithContext(ctx context.Context) BgpPeerStatusResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerStatusResponseArrayOutput)
}

type BgpPeerStatusResponseOutput struct{ *pulumi.OutputState }

func (BgpPeerStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpPeerStatusResponse)(nil)).Elem()
}

func (o BgpPeerStatusResponseOutput) ToBgpPeerStatusResponseOutput() BgpPeerStatusResponseOutput {
	return o
}

func (o BgpPeerStatusResponseOutput) ToBgpPeerStatusResponseOutputWithContext(ctx context.Context) BgpPeerStatusResponseOutput {
	return o
}

func (o BgpPeerStatusResponseOutput) Asn() pulumi.IntOutput {
	return o.ApplyT(func(v BgpPeerStatusResponse) int { return v.Asn }).(pulumi.IntOutput)
}

func (o BgpPeerStatusResponseOutput) ConnectedDuration() pulumi.StringOutput {
	return o.ApplyT(func(v BgpPeerStatusResponse) string { return v.ConnectedDuration }).(pulumi.StringOutput)
}

func (o BgpPeerStatusResponseOutput) LocalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v BgpPeerStatusResponse) string { return v.LocalAddress }).(pulumi.StringOutput)
}

func (o BgpPeerStatusResponseOutput) MessagesReceived() pulumi.Float64Output {
	return o.ApplyT(func(v BgpPeerStatusResponse) float64 { return v.MessagesReceived }).(pulumi.Float64Output)
}

func (o BgpPeerStatusResponseOutput) MessagesSent() pulumi.Float64Output {
	return o.ApplyT(func(v BgpPeerStatusResponse) float64 { return v.MessagesSent }).(pulumi.Float64Output)
}

func (o BgpPeerStatusResponseOutput) Neighbor() pulumi.StringOutput {
	return o.ApplyT(func(v BgpPeerStatusResponse) string { return v.Neighbor }).(pulumi.StringOutput)
}

func (o BgpPeerStatusResponseOutput) RoutesReceived() pulumi.Float64Output {
	return o.ApplyT(func(v BgpPeerStatusResponse) float64 { return v.RoutesReceived }).(pulumi.Float64Output)
}

func (o BgpPeerStatusResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v BgpPeerStatusResponse) string { return v.State }).(pulumi.StringOutput)
}

type BgpPeerStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (BgpPeerStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BgpPeerStatusResponse)(nil)).Elem()
}

func (o BgpPeerStatusResponseArrayOutput) ToBgpPeerStatusResponseArrayOutput() BgpPeerStatusResponseArrayOutput {
	return o
}

func (o BgpPeerStatusResponseArrayOutput) ToBgpPeerStatusResponseArrayOutputWithContext(ctx context.Context) BgpPeerStatusResponseArrayOutput {
	return o
}

func (o BgpPeerStatusResponseArrayOutput) Index(i pulumi.IntInput) BgpPeerStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BgpPeerStatusResponse {
		return vs[0].([]BgpPeerStatusResponse)[vs[1].(int)]
	}).(BgpPeerStatusResponseOutput)
}

type BgpSettings struct {
	Asn               *float64 `pulumi:"asn"`
	BgpPeeringAddress *string  `pulumi:"bgpPeeringAddress"`
	PeerWeight        *int     `pulumi:"peerWeight"`
}





type BgpSettingsInput interface {
	pulumi.Input

	ToBgpSettingsOutput() BgpSettingsOutput
	ToBgpSettingsOutputWithContext(context.Context) BgpSettingsOutput
}

type BgpSettingsArgs struct {
	Asn               pulumi.Float64PtrInput `pulumi:"asn"`
	BgpPeeringAddress pulumi.StringPtrInput  `pulumi:"bgpPeeringAddress"`
	PeerWeight        pulumi.IntPtrInput     `pulumi:"peerWeight"`
}

func (BgpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSettings)(nil)).Elem()
}

func (i BgpSettingsArgs) ToBgpSettingsOutput() BgpSettingsOutput {
	return i.ToBgpSettingsOutputWithContext(context.Background())
}

func (i BgpSettingsArgs) ToBgpSettingsOutputWithContext(ctx context.Context) BgpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSettingsOutput)
}

func (i BgpSettingsArgs) ToBgpSettingsPtrOutput() BgpSettingsPtrOutput {
	return i.ToBgpSettingsPtrOutputWithContext(context.Background())
}

func (i BgpSettingsArgs) ToBgpSettingsPtrOutputWithContext(ctx context.Context) BgpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSettingsOutput).ToBgpSettingsPtrOutputWithContext(ctx)
}









type BgpSettingsPtrInput interface {
	pulumi.Input

	ToBgpSettingsPtrOutput() BgpSettingsPtrOutput
	ToBgpSettingsPtrOutputWithContext(context.Context) BgpSettingsPtrOutput
}

type bgpSettingsPtrType BgpSettingsArgs

func BgpSettingsPtr(v *BgpSettingsArgs) BgpSettingsPtrInput {
	return (*bgpSettingsPtrType)(v)
}

func (*bgpSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSettings)(nil)).Elem()
}

func (i *bgpSettingsPtrType) ToBgpSettingsPtrOutput() BgpSettingsPtrOutput {
	return i.ToBgpSettingsPtrOutputWithContext(context.Background())
}

func (i *bgpSettingsPtrType) ToBgpSettingsPtrOutputWithContext(ctx context.Context) BgpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSettingsPtrOutput)
}

type BgpSettingsOutput struct{ *pulumi.OutputState }

func (BgpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSettings)(nil)).Elem()
}

func (o BgpSettingsOutput) ToBgpSettingsOutput() BgpSettingsOutput {
	return o
}

func (o BgpSettingsOutput) ToBgpSettingsOutputWithContext(ctx context.Context) BgpSettingsOutput {
	return o
}

func (o BgpSettingsOutput) ToBgpSettingsPtrOutput() BgpSettingsPtrOutput {
	return o.ToBgpSettingsPtrOutputWithContext(context.Background())
}

func (o BgpSettingsOutput) ToBgpSettingsPtrOutputWithContext(ctx context.Context) BgpSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BgpSettings) *BgpSettings {
		return &v
	}).(BgpSettingsPtrOutput)
}

func (o BgpSettingsOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BgpSettings) *float64 { return v.Asn }).(pulumi.Float64PtrOutput)
}

func (o BgpSettingsOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSettings) *string { return v.BgpPeeringAddress }).(pulumi.StringPtrOutput)
}

func (o BgpSettingsOutput) PeerWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpSettings) *int { return v.PeerWeight }).(pulumi.IntPtrOutput)
}

type BgpSettingsPtrOutput struct{ *pulumi.OutputState }

func (BgpSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSettings)(nil)).Elem()
}

func (o BgpSettingsPtrOutput) ToBgpSettingsPtrOutput() BgpSettingsPtrOutput {
	return o
}

func (o BgpSettingsPtrOutput) ToBgpSettingsPtrOutputWithContext(ctx context.Context) BgpSettingsPtrOutput {
	return o
}

func (o BgpSettingsPtrOutput) Elem() BgpSettingsOutput {
	return o.ApplyT(func(v *BgpSettings) BgpSettings {
		if v != nil {
			return *v
		}
		var ret BgpSettings
		return ret
	}).(BgpSettingsOutput)
}

func (o BgpSettingsPtrOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BgpSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.Asn
	}).(pulumi.Float64PtrOutput)
}

func (o BgpSettingsPtrOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSettings) *string {
		if v == nil {
			return nil
		}
		return v.BgpPeeringAddress
	}).(pulumi.StringPtrOutput)
}

func (o BgpSettingsPtrOutput) PeerWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BgpSettings) *int {
		if v == nil {
			return nil
		}
		return v.PeerWeight
	}).(pulumi.IntPtrOutput)
}

type BgpSettingsResponse struct {
	Asn               *float64 `pulumi:"asn"`
	BgpPeeringAddress *string  `pulumi:"bgpPeeringAddress"`
	PeerWeight        *int     `pulumi:"peerWeight"`
}





type BgpSettingsResponseInput interface {
	pulumi.Input

	ToBgpSettingsResponseOutput() BgpSettingsResponseOutput
	ToBgpSettingsResponseOutputWithContext(context.Context) BgpSettingsResponseOutput
}

type BgpSettingsResponseArgs struct {
	Asn               pulumi.Float64PtrInput `pulumi:"asn"`
	BgpPeeringAddress pulumi.StringPtrInput  `pulumi:"bgpPeeringAddress"`
	PeerWeight        pulumi.IntPtrInput     `pulumi:"peerWeight"`
}

func (BgpSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSettingsResponse)(nil)).Elem()
}

func (i BgpSettingsResponseArgs) ToBgpSettingsResponseOutput() BgpSettingsResponseOutput {
	return i.ToBgpSettingsResponseOutputWithContext(context.Background())
}

func (i BgpSettingsResponseArgs) ToBgpSettingsResponseOutputWithContext(ctx context.Context) BgpSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSettingsResponseOutput)
}

func (i BgpSettingsResponseArgs) ToBgpSettingsResponsePtrOutput() BgpSettingsResponsePtrOutput {
	return i.ToBgpSettingsResponsePtrOutputWithContext(context.Background())
}

func (i BgpSettingsResponseArgs) ToBgpSettingsResponsePtrOutputWithContext(ctx context.Context) BgpSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSettingsResponseOutput).ToBgpSettingsResponsePtrOutputWithContext(ctx)
}









type BgpSettingsResponsePtrInput interface {
	pulumi.Input

	ToBgpSettingsResponsePtrOutput() BgpSettingsResponsePtrOutput
	ToBgpSettingsResponsePtrOutputWithContext(context.Context) BgpSettingsResponsePtrOutput
}

type bgpSettingsResponsePtrType BgpSettingsResponseArgs

func BgpSettingsResponsePtr(v *BgpSettingsResponseArgs) BgpSettingsResponsePtrInput {
	return (*bgpSettingsResponsePtrType)(v)
}

func (*bgpSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSettingsResponse)(nil)).Elem()
}

func (i *bgpSettingsResponsePtrType) ToBgpSettingsResponsePtrOutput() BgpSettingsResponsePtrOutput {
	return i.ToBgpSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *bgpSettingsResponsePtrType) ToBgpSettingsResponsePtrOutputWithContext(ctx context.Context) BgpSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSettingsResponsePtrOutput)
}

type BgpSettingsResponseOutput struct{ *pulumi.OutputState }

func (BgpSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSettingsResponse)(nil)).Elem()
}

func (o BgpSettingsResponseOutput) ToBgpSettingsResponseOutput() BgpSettingsResponseOutput {
	return o
}

func (o BgpSettingsResponseOutput) ToBgpSettingsResponseOutputWithContext(ctx context.Context) BgpSettingsResponseOutput {
	return o
}

func (o BgpSettingsResponseOutput) ToBgpSettingsResponsePtrOutput() BgpSettingsResponsePtrOutput {
	return o.ToBgpSettingsResponsePtrOutputWithContext(context.Background())
}

func (o BgpSettingsResponseOutput) ToBgpSettingsResponsePtrOutputWithContext(ctx context.Context) BgpSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BgpSettingsResponse) *BgpSettingsResponse {
		return &v
	}).(BgpSettingsResponsePtrOutput)
}

func (o BgpSettingsResponseOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BgpSettingsResponse) *float64 { return v.Asn }).(pulumi.Float64PtrOutput)
}

func (o BgpSettingsResponseOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSettingsResponse) *string { return v.BgpPeeringAddress }).(pulumi.StringPtrOutput)
}

func (o BgpSettingsResponseOutput) PeerWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpSettingsResponse) *int { return v.PeerWeight }).(pulumi.IntPtrOutput)
}

type BgpSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (BgpSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSettingsResponse)(nil)).Elem()
}

func (o BgpSettingsResponsePtrOutput) ToBgpSettingsResponsePtrOutput() BgpSettingsResponsePtrOutput {
	return o
}

func (o BgpSettingsResponsePtrOutput) ToBgpSettingsResponsePtrOutputWithContext(ctx context.Context) BgpSettingsResponsePtrOutput {
	return o
}

func (o BgpSettingsResponsePtrOutput) Elem() BgpSettingsResponseOutput {
	return o.ApplyT(func(v *BgpSettingsResponse) BgpSettingsResponse {
		if v != nil {
			return *v
		}
		var ret BgpSettingsResponse
		return ret
	}).(BgpSettingsResponseOutput)
}

func (o BgpSettingsResponsePtrOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BgpSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Asn
	}).(pulumi.Float64PtrOutput)
}

func (o BgpSettingsResponsePtrOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.BgpPeeringAddress
	}).(pulumi.StringPtrOutput)
}

func (o BgpSettingsResponsePtrOutput) PeerWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BgpSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.PeerWeight
	}).(pulumi.IntPtrOutput)
}

type ConnectionMonitorDestination struct {
	Address    *string `pulumi:"address"`
	Port       *int    `pulumi:"port"`
	ResourceId *string `pulumi:"resourceId"`
}





type ConnectionMonitorDestinationInput interface {
	pulumi.Input

	ToConnectionMonitorDestinationOutput() ConnectionMonitorDestinationOutput
	ToConnectionMonitorDestinationOutputWithContext(context.Context) ConnectionMonitorDestinationOutput
}

type ConnectionMonitorDestinationArgs struct {
	Address    pulumi.StringPtrInput `pulumi:"address"`
	Port       pulumi.IntPtrInput    `pulumi:"port"`
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ConnectionMonitorDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorDestination)(nil)).Elem()
}

func (i ConnectionMonitorDestinationArgs) ToConnectionMonitorDestinationOutput() ConnectionMonitorDestinationOutput {
	return i.ToConnectionMonitorDestinationOutputWithContext(context.Background())
}

func (i ConnectionMonitorDestinationArgs) ToConnectionMonitorDestinationOutputWithContext(ctx context.Context) ConnectionMonitorDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorDestinationOutput)
}

func (i ConnectionMonitorDestinationArgs) ToConnectionMonitorDestinationPtrOutput() ConnectionMonitorDestinationPtrOutput {
	return i.ToConnectionMonitorDestinationPtrOutputWithContext(context.Background())
}

func (i ConnectionMonitorDestinationArgs) ToConnectionMonitorDestinationPtrOutputWithContext(ctx context.Context) ConnectionMonitorDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorDestinationOutput).ToConnectionMonitorDestinationPtrOutputWithContext(ctx)
}









type ConnectionMonitorDestinationPtrInput interface {
	pulumi.Input

	ToConnectionMonitorDestinationPtrOutput() ConnectionMonitorDestinationPtrOutput
	ToConnectionMonitorDestinationPtrOutputWithContext(context.Context) ConnectionMonitorDestinationPtrOutput
}

type connectionMonitorDestinationPtrType ConnectionMonitorDestinationArgs

func ConnectionMonitorDestinationPtr(v *ConnectionMonitorDestinationArgs) ConnectionMonitorDestinationPtrInput {
	return (*connectionMonitorDestinationPtrType)(v)
}

func (*connectionMonitorDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorDestination)(nil)).Elem()
}

func (i *connectionMonitorDestinationPtrType) ToConnectionMonitorDestinationPtrOutput() ConnectionMonitorDestinationPtrOutput {
	return i.ToConnectionMonitorDestinationPtrOutputWithContext(context.Background())
}

func (i *connectionMonitorDestinationPtrType) ToConnectionMonitorDestinationPtrOutputWithContext(ctx context.Context) ConnectionMonitorDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorDestinationPtrOutput)
}

type ConnectionMonitorDestinationOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorDestination)(nil)).Elem()
}

func (o ConnectionMonitorDestinationOutput) ToConnectionMonitorDestinationOutput() ConnectionMonitorDestinationOutput {
	return o
}

func (o ConnectionMonitorDestinationOutput) ToConnectionMonitorDestinationOutputWithContext(ctx context.Context) ConnectionMonitorDestinationOutput {
	return o
}

func (o ConnectionMonitorDestinationOutput) ToConnectionMonitorDestinationPtrOutput() ConnectionMonitorDestinationPtrOutput {
	return o.ToConnectionMonitorDestinationPtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorDestinationOutput) ToConnectionMonitorDestinationPtrOutputWithContext(ctx context.Context) ConnectionMonitorDestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionMonitorDestination) *ConnectionMonitorDestination {
		return &v
	}).(ConnectionMonitorDestinationPtrOutput)
}

func (o ConnectionMonitorDestinationOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionMonitorDestination) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o ConnectionMonitorDestinationOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConnectionMonitorDestination) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionMonitorDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ConnectionMonitorDestinationPtrOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorDestination)(nil)).Elem()
}

func (o ConnectionMonitorDestinationPtrOutput) ToConnectionMonitorDestinationPtrOutput() ConnectionMonitorDestinationPtrOutput {
	return o
}

func (o ConnectionMonitorDestinationPtrOutput) ToConnectionMonitorDestinationPtrOutputWithContext(ctx context.Context) ConnectionMonitorDestinationPtrOutput {
	return o
}

func (o ConnectionMonitorDestinationPtrOutput) Elem() ConnectionMonitorDestinationOutput {
	return o.ApplyT(func(v *ConnectionMonitorDestination) ConnectionMonitorDestination {
		if v != nil {
			return *v
		}
		var ret ConnectionMonitorDestination
		return ret
	}).(ConnectionMonitorDestinationOutput)
}

func (o ConnectionMonitorDestinationPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorDestination) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionMonitorDestinationPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorDestination) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorDestinationPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorDestination) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type ConnectionMonitorDestinationResponse struct {
	Address    *string `pulumi:"address"`
	Port       *int    `pulumi:"port"`
	ResourceId *string `pulumi:"resourceId"`
}





type ConnectionMonitorDestinationResponseInput interface {
	pulumi.Input

	ToConnectionMonitorDestinationResponseOutput() ConnectionMonitorDestinationResponseOutput
	ToConnectionMonitorDestinationResponseOutputWithContext(context.Context) ConnectionMonitorDestinationResponseOutput
}

type ConnectionMonitorDestinationResponseArgs struct {
	Address    pulumi.StringPtrInput `pulumi:"address"`
	Port       pulumi.IntPtrInput    `pulumi:"port"`
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ConnectionMonitorDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorDestinationResponse)(nil)).Elem()
}

func (i ConnectionMonitorDestinationResponseArgs) ToConnectionMonitorDestinationResponseOutput() ConnectionMonitorDestinationResponseOutput {
	return i.ToConnectionMonitorDestinationResponseOutputWithContext(context.Background())
}

func (i ConnectionMonitorDestinationResponseArgs) ToConnectionMonitorDestinationResponseOutputWithContext(ctx context.Context) ConnectionMonitorDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorDestinationResponseOutput)
}

func (i ConnectionMonitorDestinationResponseArgs) ToConnectionMonitorDestinationResponsePtrOutput() ConnectionMonitorDestinationResponsePtrOutput {
	return i.ToConnectionMonitorDestinationResponsePtrOutputWithContext(context.Background())
}

func (i ConnectionMonitorDestinationResponseArgs) ToConnectionMonitorDestinationResponsePtrOutputWithContext(ctx context.Context) ConnectionMonitorDestinationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorDestinationResponseOutput).ToConnectionMonitorDestinationResponsePtrOutputWithContext(ctx)
}









type ConnectionMonitorDestinationResponsePtrInput interface {
	pulumi.Input

	ToConnectionMonitorDestinationResponsePtrOutput() ConnectionMonitorDestinationResponsePtrOutput
	ToConnectionMonitorDestinationResponsePtrOutputWithContext(context.Context) ConnectionMonitorDestinationResponsePtrOutput
}

type connectionMonitorDestinationResponsePtrType ConnectionMonitorDestinationResponseArgs

func ConnectionMonitorDestinationResponsePtr(v *ConnectionMonitorDestinationResponseArgs) ConnectionMonitorDestinationResponsePtrInput {
	return (*connectionMonitorDestinationResponsePtrType)(v)
}

func (*connectionMonitorDestinationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorDestinationResponse)(nil)).Elem()
}

func (i *connectionMonitorDestinationResponsePtrType) ToConnectionMonitorDestinationResponsePtrOutput() ConnectionMonitorDestinationResponsePtrOutput {
	return i.ToConnectionMonitorDestinationResponsePtrOutputWithContext(context.Background())
}

func (i *connectionMonitorDestinationResponsePtrType) ToConnectionMonitorDestinationResponsePtrOutputWithContext(ctx context.Context) ConnectionMonitorDestinationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorDestinationResponsePtrOutput)
}

type ConnectionMonitorDestinationResponseOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorDestinationResponse)(nil)).Elem()
}

func (o ConnectionMonitorDestinationResponseOutput) ToConnectionMonitorDestinationResponseOutput() ConnectionMonitorDestinationResponseOutput {
	return o
}

func (o ConnectionMonitorDestinationResponseOutput) ToConnectionMonitorDestinationResponseOutputWithContext(ctx context.Context) ConnectionMonitorDestinationResponseOutput {
	return o
}

func (o ConnectionMonitorDestinationResponseOutput) ToConnectionMonitorDestinationResponsePtrOutput() ConnectionMonitorDestinationResponsePtrOutput {
	return o.ToConnectionMonitorDestinationResponsePtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorDestinationResponseOutput) ToConnectionMonitorDestinationResponsePtrOutputWithContext(ctx context.Context) ConnectionMonitorDestinationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionMonitorDestinationResponse) *ConnectionMonitorDestinationResponse {
		return &v
	}).(ConnectionMonitorDestinationResponsePtrOutput)
}

func (o ConnectionMonitorDestinationResponseOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionMonitorDestinationResponse) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o ConnectionMonitorDestinationResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConnectionMonitorDestinationResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionMonitorDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ConnectionMonitorDestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorDestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorDestinationResponse)(nil)).Elem()
}

func (o ConnectionMonitorDestinationResponsePtrOutput) ToConnectionMonitorDestinationResponsePtrOutput() ConnectionMonitorDestinationResponsePtrOutput {
	return o
}

func (o ConnectionMonitorDestinationResponsePtrOutput) ToConnectionMonitorDestinationResponsePtrOutputWithContext(ctx context.Context) ConnectionMonitorDestinationResponsePtrOutput {
	return o
}

func (o ConnectionMonitorDestinationResponsePtrOutput) Elem() ConnectionMonitorDestinationResponseOutput {
	return o.ApplyT(func(v *ConnectionMonitorDestinationResponse) ConnectionMonitorDestinationResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionMonitorDestinationResponse
		return ret
	}).(ConnectionMonitorDestinationResponseOutput)
}

func (o ConnectionMonitorDestinationResponsePtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionMonitorDestinationResponsePtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorDestinationResponse) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorDestinationResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type ConnectionMonitorSource struct {
	Port       *int   `pulumi:"port"`
	ResourceId string `pulumi:"resourceId"`
}





type ConnectionMonitorSourceInput interface {
	pulumi.Input

	ToConnectionMonitorSourceOutput() ConnectionMonitorSourceOutput
	ToConnectionMonitorSourceOutputWithContext(context.Context) ConnectionMonitorSourceOutput
}

type ConnectionMonitorSourceArgs struct {
	Port       pulumi.IntPtrInput `pulumi:"port"`
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
}

func (ConnectionMonitorSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorSource)(nil)).Elem()
}

func (i ConnectionMonitorSourceArgs) ToConnectionMonitorSourceOutput() ConnectionMonitorSourceOutput {
	return i.ToConnectionMonitorSourceOutputWithContext(context.Background())
}

func (i ConnectionMonitorSourceArgs) ToConnectionMonitorSourceOutputWithContext(ctx context.Context) ConnectionMonitorSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorSourceOutput)
}

func (i ConnectionMonitorSourceArgs) ToConnectionMonitorSourcePtrOutput() ConnectionMonitorSourcePtrOutput {
	return i.ToConnectionMonitorSourcePtrOutputWithContext(context.Background())
}

func (i ConnectionMonitorSourceArgs) ToConnectionMonitorSourcePtrOutputWithContext(ctx context.Context) ConnectionMonitorSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorSourceOutput).ToConnectionMonitorSourcePtrOutputWithContext(ctx)
}









type ConnectionMonitorSourcePtrInput interface {
	pulumi.Input

	ToConnectionMonitorSourcePtrOutput() ConnectionMonitorSourcePtrOutput
	ToConnectionMonitorSourcePtrOutputWithContext(context.Context) ConnectionMonitorSourcePtrOutput
}

type connectionMonitorSourcePtrType ConnectionMonitorSourceArgs

func ConnectionMonitorSourcePtr(v *ConnectionMonitorSourceArgs) ConnectionMonitorSourcePtrInput {
	return (*connectionMonitorSourcePtrType)(v)
}

func (*connectionMonitorSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorSource)(nil)).Elem()
}

func (i *connectionMonitorSourcePtrType) ToConnectionMonitorSourcePtrOutput() ConnectionMonitorSourcePtrOutput {
	return i.ToConnectionMonitorSourcePtrOutputWithContext(context.Background())
}

func (i *connectionMonitorSourcePtrType) ToConnectionMonitorSourcePtrOutputWithContext(ctx context.Context) ConnectionMonitorSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorSourcePtrOutput)
}

type ConnectionMonitorSourceOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorSource)(nil)).Elem()
}

func (o ConnectionMonitorSourceOutput) ToConnectionMonitorSourceOutput() ConnectionMonitorSourceOutput {
	return o
}

func (o ConnectionMonitorSourceOutput) ToConnectionMonitorSourceOutputWithContext(ctx context.Context) ConnectionMonitorSourceOutput {
	return o
}

func (o ConnectionMonitorSourceOutput) ToConnectionMonitorSourcePtrOutput() ConnectionMonitorSourcePtrOutput {
	return o.ToConnectionMonitorSourcePtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorSourceOutput) ToConnectionMonitorSourcePtrOutputWithContext(ctx context.Context) ConnectionMonitorSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionMonitorSource) *ConnectionMonitorSource {
		return &v
	}).(ConnectionMonitorSourcePtrOutput)
}

func (o ConnectionMonitorSourceOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConnectionMonitorSource) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorSourceOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionMonitorSource) string { return v.ResourceId }).(pulumi.StringOutput)
}

type ConnectionMonitorSourcePtrOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorSource)(nil)).Elem()
}

func (o ConnectionMonitorSourcePtrOutput) ToConnectionMonitorSourcePtrOutput() ConnectionMonitorSourcePtrOutput {
	return o
}

func (o ConnectionMonitorSourcePtrOutput) ToConnectionMonitorSourcePtrOutputWithContext(ctx context.Context) ConnectionMonitorSourcePtrOutput {
	return o
}

func (o ConnectionMonitorSourcePtrOutput) Elem() ConnectionMonitorSourceOutput {
	return o.ApplyT(func(v *ConnectionMonitorSource) ConnectionMonitorSource {
		if v != nil {
			return *v
		}
		var ret ConnectionMonitorSource
		return ret
	}).(ConnectionMonitorSourceOutput)
}

func (o ConnectionMonitorSourcePtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorSource) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorSourcePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorSource) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type ConnectionMonitorSourceResponse struct {
	Port       *int   `pulumi:"port"`
	ResourceId string `pulumi:"resourceId"`
}





type ConnectionMonitorSourceResponseInput interface {
	pulumi.Input

	ToConnectionMonitorSourceResponseOutput() ConnectionMonitorSourceResponseOutput
	ToConnectionMonitorSourceResponseOutputWithContext(context.Context) ConnectionMonitorSourceResponseOutput
}

type ConnectionMonitorSourceResponseArgs struct {
	Port       pulumi.IntPtrInput `pulumi:"port"`
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
}

func (ConnectionMonitorSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorSourceResponse)(nil)).Elem()
}

func (i ConnectionMonitorSourceResponseArgs) ToConnectionMonitorSourceResponseOutput() ConnectionMonitorSourceResponseOutput {
	return i.ToConnectionMonitorSourceResponseOutputWithContext(context.Background())
}

func (i ConnectionMonitorSourceResponseArgs) ToConnectionMonitorSourceResponseOutputWithContext(ctx context.Context) ConnectionMonitorSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorSourceResponseOutput)
}

func (i ConnectionMonitorSourceResponseArgs) ToConnectionMonitorSourceResponsePtrOutput() ConnectionMonitorSourceResponsePtrOutput {
	return i.ToConnectionMonitorSourceResponsePtrOutputWithContext(context.Background())
}

func (i ConnectionMonitorSourceResponseArgs) ToConnectionMonitorSourceResponsePtrOutputWithContext(ctx context.Context) ConnectionMonitorSourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorSourceResponseOutput).ToConnectionMonitorSourceResponsePtrOutputWithContext(ctx)
}









type ConnectionMonitorSourceResponsePtrInput interface {
	pulumi.Input

	ToConnectionMonitorSourceResponsePtrOutput() ConnectionMonitorSourceResponsePtrOutput
	ToConnectionMonitorSourceResponsePtrOutputWithContext(context.Context) ConnectionMonitorSourceResponsePtrOutput
}

type connectionMonitorSourceResponsePtrType ConnectionMonitorSourceResponseArgs

func ConnectionMonitorSourceResponsePtr(v *ConnectionMonitorSourceResponseArgs) ConnectionMonitorSourceResponsePtrInput {
	return (*connectionMonitorSourceResponsePtrType)(v)
}

func (*connectionMonitorSourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorSourceResponse)(nil)).Elem()
}

func (i *connectionMonitorSourceResponsePtrType) ToConnectionMonitorSourceResponsePtrOutput() ConnectionMonitorSourceResponsePtrOutput {
	return i.ToConnectionMonitorSourceResponsePtrOutputWithContext(context.Background())
}

func (i *connectionMonitorSourceResponsePtrType) ToConnectionMonitorSourceResponsePtrOutputWithContext(ctx context.Context) ConnectionMonitorSourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorSourceResponsePtrOutput)
}

type ConnectionMonitorSourceResponseOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorSourceResponse)(nil)).Elem()
}

func (o ConnectionMonitorSourceResponseOutput) ToConnectionMonitorSourceResponseOutput() ConnectionMonitorSourceResponseOutput {
	return o
}

func (o ConnectionMonitorSourceResponseOutput) ToConnectionMonitorSourceResponseOutputWithContext(ctx context.Context) ConnectionMonitorSourceResponseOutput {
	return o
}

func (o ConnectionMonitorSourceResponseOutput) ToConnectionMonitorSourceResponsePtrOutput() ConnectionMonitorSourceResponsePtrOutput {
	return o.ToConnectionMonitorSourceResponsePtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorSourceResponseOutput) ToConnectionMonitorSourceResponsePtrOutputWithContext(ctx context.Context) ConnectionMonitorSourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionMonitorSourceResponse) *ConnectionMonitorSourceResponse {
		return &v
	}).(ConnectionMonitorSourceResponsePtrOutput)
}

func (o ConnectionMonitorSourceResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConnectionMonitorSourceResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorSourceResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionMonitorSourceResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

type ConnectionMonitorSourceResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorSourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorSourceResponse)(nil)).Elem()
}

func (o ConnectionMonitorSourceResponsePtrOutput) ToConnectionMonitorSourceResponsePtrOutput() ConnectionMonitorSourceResponsePtrOutput {
	return o
}

func (o ConnectionMonitorSourceResponsePtrOutput) ToConnectionMonitorSourceResponsePtrOutputWithContext(ctx context.Context) ConnectionMonitorSourceResponsePtrOutput {
	return o
}

func (o ConnectionMonitorSourceResponsePtrOutput) Elem() ConnectionMonitorSourceResponseOutput {
	return o.ApplyT(func(v *ConnectionMonitorSourceResponse) ConnectionMonitorSourceResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionMonitorSourceResponse
		return ret
	}).(ConnectionMonitorSourceResponseOutput)
}

func (o ConnectionMonitorSourceResponsePtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorSourceResponse) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorSourceResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorSourceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
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
	Etag                   string  `pulumi:"etag"`
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
	Etag                   pulumi.StringInput    `pulumi:"etag"`
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

func (o ExpressRouteCircuitAuthorizationResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) string { return v.Etag }).(pulumi.StringOutput)
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
	AzureASN                   *int                                  `pulumi:"azureASN"`
	GatewayManagerEtag         *string                               `pulumi:"gatewayManagerEtag"`
	Id                         *string                               `pulumi:"id"`
	Ipv6PeeringConfig          *Ipv6ExpressRouteCircuitPeeringConfig `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             *string                               `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfig     `pulumi:"microsoftPeeringConfig"`
	Name                       *string                               `pulumi:"name"`
	PeerASN                    *float64                              `pulumi:"peerASN"`
	PeeringType                *string                               `pulumi:"peeringType"`
	PrimaryAzurePort           *string                               `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   *string                               `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          *string                               `pulumi:"provisioningState"`
	RouteFilter                *RouteFilterType                      `pulumi:"routeFilter"`
	SecondaryAzurePort         *string                               `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix *string                               `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  *string                               `pulumi:"sharedKey"`
	State                      *string                               `pulumi:"state"`
	Stats                      *ExpressRouteCircuitStats             `pulumi:"stats"`
	VlanId                     *int                                  `pulumi:"vlanId"`
}





type ExpressRouteCircuitPeeringTypeInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringTypeOutput() ExpressRouteCircuitPeeringTypeOutput
	ToExpressRouteCircuitPeeringTypeOutputWithContext(context.Context) ExpressRouteCircuitPeeringTypeOutput
}

type ExpressRouteCircuitPeeringTypeArgs struct {
	AzureASN                   pulumi.IntPtrInput                           `pulumi:"azureASN"`
	GatewayManagerEtag         pulumi.StringPtrInput                        `pulumi:"gatewayManagerEtag"`
	Id                         pulumi.StringPtrInput                        `pulumi:"id"`
	Ipv6PeeringConfig          Ipv6ExpressRouteCircuitPeeringConfigPtrInput `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             pulumi.StringPtrInput                        `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigPtrInput     `pulumi:"microsoftPeeringConfig"`
	Name                       pulumi.StringPtrInput                        `pulumi:"name"`
	PeerASN                    pulumi.Float64PtrInput                       `pulumi:"peerASN"`
	PeeringType                pulumi.StringPtrInput                        `pulumi:"peeringType"`
	PrimaryAzurePort           pulumi.StringPtrInput                        `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   pulumi.StringPtrInput                        `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          pulumi.StringPtrInput                        `pulumi:"provisioningState"`
	RouteFilter                RouteFilterTypePtrInput                      `pulumi:"routeFilter"`
	SecondaryAzurePort         pulumi.StringPtrInput                        `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix pulumi.StringPtrInput                        `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  pulumi.StringPtrInput                        `pulumi:"sharedKey"`
	State                      pulumi.StringPtrInput                        `pulumi:"state"`
	Stats                      ExpressRouteCircuitStatsPtrInput             `pulumi:"stats"`
	VlanId                     pulumi.IntPtrInput                           `pulumi:"vlanId"`
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

func (o ExpressRouteCircuitPeeringTypeOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *Ipv6ExpressRouteCircuitPeeringConfig {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *ExpressRouteCircuitPeeringConfig {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) PeerASN() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *float64 { return v.PeerASN }).(pulumi.Float64PtrOutput)
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

func (o ExpressRouteCircuitPeeringTypeOutput) RouteFilter() RouteFilterTypePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *RouteFilterType { return v.RouteFilter }).(RouteFilterTypePtrOutput)
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
	AdvertisedCommunities         []string `pulumi:"advertisedCommunities"`
	AdvertisedPublicPrefixes      []string `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState *string  `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   *int     `pulumi:"customerASN"`
	LegacyMode                    *int     `pulumi:"legacyMode"`
	RoutingRegistryName           *string  `pulumi:"routingRegistryName"`
}





type ExpressRouteCircuitPeeringConfigInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringConfigOutput() ExpressRouteCircuitPeeringConfigOutput
	ToExpressRouteCircuitPeeringConfigOutputWithContext(context.Context) ExpressRouteCircuitPeeringConfigOutput
}

type ExpressRouteCircuitPeeringConfigArgs struct {
	AdvertisedCommunities         pulumi.StringArrayInput `pulumi:"advertisedCommunities"`
	AdvertisedPublicPrefixes      pulumi.StringArrayInput `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState pulumi.StringPtrInput   `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   pulumi.IntPtrInput      `pulumi:"customerASN"`
	LegacyMode                    pulumi.IntPtrInput      `pulumi:"legacyMode"`
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

func (o ExpressRouteCircuitPeeringConfigOutput) AdvertisedCommunities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfig) []string { return v.AdvertisedCommunities }).(pulumi.StringArrayOutput)
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

func (o ExpressRouteCircuitPeeringConfigOutput) LegacyMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfig) *int { return v.LegacyMode }).(pulumi.IntPtrOutput)
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

func (o ExpressRouteCircuitPeeringConfigPtrOutput) AdvertisedCommunities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) []string {
		if v == nil {
			return nil
		}
		return v.AdvertisedCommunities
	}).(pulumi.StringArrayOutput)
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

func (o ExpressRouteCircuitPeeringConfigPtrOutput) LegacyMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) *int {
		if v == nil {
			return nil
		}
		return v.LegacyMode
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
	AdvertisedCommunities         []string `pulumi:"advertisedCommunities"`
	AdvertisedPublicPrefixes      []string `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState *string  `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   *int     `pulumi:"customerASN"`
	LegacyMode                    *int     `pulumi:"legacyMode"`
	RoutingRegistryName           *string  `pulumi:"routingRegistryName"`
}





type ExpressRouteCircuitPeeringConfigResponseInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringConfigResponseOutput() ExpressRouteCircuitPeeringConfigResponseOutput
	ToExpressRouteCircuitPeeringConfigResponseOutputWithContext(context.Context) ExpressRouteCircuitPeeringConfigResponseOutput
}

type ExpressRouteCircuitPeeringConfigResponseArgs struct {
	AdvertisedCommunities         pulumi.StringArrayInput `pulumi:"advertisedCommunities"`
	AdvertisedPublicPrefixes      pulumi.StringArrayInput `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState pulumi.StringPtrInput   `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   pulumi.IntPtrInput      `pulumi:"customerASN"`
	LegacyMode                    pulumi.IntPtrInput      `pulumi:"legacyMode"`
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

func (o ExpressRouteCircuitPeeringConfigResponseOutput) AdvertisedCommunities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfigResponse) []string { return v.AdvertisedCommunities }).(pulumi.StringArrayOutput)
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

func (o ExpressRouteCircuitPeeringConfigResponseOutput) LegacyMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfigResponse) *int { return v.LegacyMode }).(pulumi.IntPtrOutput)
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

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) AdvertisedCommunities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) []string {
		if v == nil {
			return nil
		}
		return v.AdvertisedCommunities
	}).(pulumi.StringArrayOutput)
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

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) LegacyMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.LegacyMode
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
	AzureASN                   *int                                          `pulumi:"azureASN"`
	Etag                       string                                        `pulumi:"etag"`
	GatewayManagerEtag         *string                                       `pulumi:"gatewayManagerEtag"`
	Id                         *string                                       `pulumi:"id"`
	Ipv6PeeringConfig          *Ipv6ExpressRouteCircuitPeeringConfigResponse `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             *string                                       `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfigResponse     `pulumi:"microsoftPeeringConfig"`
	Name                       *string                                       `pulumi:"name"`
	PeerASN                    *float64                                      `pulumi:"peerASN"`
	PeeringType                *string                                       `pulumi:"peeringType"`
	PrimaryAzurePort           *string                                       `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   *string                                       `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          *string                                       `pulumi:"provisioningState"`
	RouteFilter                *RouteFilterResponse                          `pulumi:"routeFilter"`
	SecondaryAzurePort         *string                                       `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix *string                                       `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  *string                                       `pulumi:"sharedKey"`
	State                      *string                                       `pulumi:"state"`
	Stats                      *ExpressRouteCircuitStatsResponse             `pulumi:"stats"`
	VlanId                     *int                                          `pulumi:"vlanId"`
}





type ExpressRouteCircuitPeeringResponseInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringResponseOutput() ExpressRouteCircuitPeeringResponseOutput
	ToExpressRouteCircuitPeeringResponseOutputWithContext(context.Context) ExpressRouteCircuitPeeringResponseOutput
}

type ExpressRouteCircuitPeeringResponseArgs struct {
	AzureASN                   pulumi.IntPtrInput                                   `pulumi:"azureASN"`
	Etag                       pulumi.StringInput                                   `pulumi:"etag"`
	GatewayManagerEtag         pulumi.StringPtrInput                                `pulumi:"gatewayManagerEtag"`
	Id                         pulumi.StringPtrInput                                `pulumi:"id"`
	Ipv6PeeringConfig          Ipv6ExpressRouteCircuitPeeringConfigResponsePtrInput `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             pulumi.StringPtrInput                                `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigResponsePtrInput     `pulumi:"microsoftPeeringConfig"`
	Name                       pulumi.StringPtrInput                                `pulumi:"name"`
	PeerASN                    pulumi.Float64PtrInput                               `pulumi:"peerASN"`
	PeeringType                pulumi.StringPtrInput                                `pulumi:"peeringType"`
	PrimaryAzurePort           pulumi.StringPtrInput                                `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   pulumi.StringPtrInput                                `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          pulumi.StringPtrInput                                `pulumi:"provisioningState"`
	RouteFilter                RouteFilterResponsePtrInput                          `pulumi:"routeFilter"`
	SecondaryAzurePort         pulumi.StringPtrInput                                `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix pulumi.StringPtrInput                                `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  pulumi.StringPtrInput                                `pulumi:"sharedKey"`
	State                      pulumi.StringPtrInput                                `pulumi:"state"`
	Stats                      ExpressRouteCircuitStatsResponsePtrInput             `pulumi:"stats"`
	VlanId                     pulumi.IntPtrInput                                   `pulumi:"vlanId"`
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

func (o ExpressRouteCircuitPeeringResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *Ipv6ExpressRouteCircuitPeeringConfigResponse {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *ExpressRouteCircuitPeeringConfigResponse {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) PeerASN() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *float64 { return v.PeerASN }).(pulumi.Float64PtrOutput)
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

func (o ExpressRouteCircuitPeeringResponseOutput) RouteFilter() RouteFilterResponsePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *RouteFilterResponse { return v.RouteFilter }).(RouteFilterResponsePtrOutput)
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
	PrimarybytesIn    *float64 `pulumi:"primarybytesIn"`
	PrimarybytesOut   *float64 `pulumi:"primarybytesOut"`
	SecondarybytesIn  *float64 `pulumi:"secondarybytesIn"`
	SecondarybytesOut *float64 `pulumi:"secondarybytesOut"`
}





type ExpressRouteCircuitStatsInput interface {
	pulumi.Input

	ToExpressRouteCircuitStatsOutput() ExpressRouteCircuitStatsOutput
	ToExpressRouteCircuitStatsOutputWithContext(context.Context) ExpressRouteCircuitStatsOutput
}

type ExpressRouteCircuitStatsArgs struct {
	PrimarybytesIn    pulumi.Float64PtrInput `pulumi:"primarybytesIn"`
	PrimarybytesOut   pulumi.Float64PtrInput `pulumi:"primarybytesOut"`
	SecondarybytesIn  pulumi.Float64PtrInput `pulumi:"secondarybytesIn"`
	SecondarybytesOut pulumi.Float64PtrInput `pulumi:"secondarybytesOut"`
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

func (o ExpressRouteCircuitStatsOutput) PrimarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStats) *float64 { return v.PrimarybytesIn }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsOutput) PrimarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStats) *float64 { return v.PrimarybytesOut }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsOutput) SecondarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStats) *float64 { return v.SecondarybytesIn }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsOutput) SecondarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStats) *float64 { return v.SecondarybytesOut }).(pulumi.Float64PtrOutput)
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

func (o ExpressRouteCircuitStatsPtrOutput) PrimarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) *float64 {
		if v == nil {
			return nil
		}
		return v.PrimarybytesIn
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsPtrOutput) PrimarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) *float64 {
		if v == nil {
			return nil
		}
		return v.PrimarybytesOut
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsPtrOutput) SecondarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) *float64 {
		if v == nil {
			return nil
		}
		return v.SecondarybytesIn
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsPtrOutput) SecondarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) *float64 {
		if v == nil {
			return nil
		}
		return v.SecondarybytesOut
	}).(pulumi.Float64PtrOutput)
}

type ExpressRouteCircuitStatsResponse struct {
	PrimarybytesIn    *float64 `pulumi:"primarybytesIn"`
	PrimarybytesOut   *float64 `pulumi:"primarybytesOut"`
	SecondarybytesIn  *float64 `pulumi:"secondarybytesIn"`
	SecondarybytesOut *float64 `pulumi:"secondarybytesOut"`
}





type ExpressRouteCircuitStatsResponseInput interface {
	pulumi.Input

	ToExpressRouteCircuitStatsResponseOutput() ExpressRouteCircuitStatsResponseOutput
	ToExpressRouteCircuitStatsResponseOutputWithContext(context.Context) ExpressRouteCircuitStatsResponseOutput
}

type ExpressRouteCircuitStatsResponseArgs struct {
	PrimarybytesIn    pulumi.Float64PtrInput `pulumi:"primarybytesIn"`
	PrimarybytesOut   pulumi.Float64PtrInput `pulumi:"primarybytesOut"`
	SecondarybytesIn  pulumi.Float64PtrInput `pulumi:"secondarybytesIn"`
	SecondarybytesOut pulumi.Float64PtrInput `pulumi:"secondarybytesOut"`
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

func (o ExpressRouteCircuitStatsResponseOutput) PrimarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStatsResponse) *float64 { return v.PrimarybytesIn }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponseOutput) PrimarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStatsResponse) *float64 { return v.PrimarybytesOut }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponseOutput) SecondarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStatsResponse) *float64 { return v.SecondarybytesIn }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponseOutput) SecondarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStatsResponse) *float64 { return v.SecondarybytesOut }).(pulumi.Float64PtrOutput)
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

func (o ExpressRouteCircuitStatsResponsePtrOutput) PrimarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.PrimarybytesIn
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) PrimarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.PrimarybytesOut
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) SecondarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SecondarybytesIn
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) SecondarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SecondarybytesOut
	}).(pulumi.Float64PtrOutput)
}

type FrontendIPConfiguration struct {
	Etag                      *string              `pulumi:"etag"`
	Id                        *string              `pulumi:"id"`
	Name                      *string              `pulumi:"name"`
	PrivateIPAddress          *string              `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string              `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string              `pulumi:"provisioningState"`
	PublicIPAddress           *PublicIPAddressType `pulumi:"publicIPAddress"`
	Subnet                    *SubnetType          `pulumi:"subnet"`
	Zones                     []string             `pulumi:"zones"`
}





type FrontendIPConfigurationInput interface {
	pulumi.Input

	ToFrontendIPConfigurationOutput() FrontendIPConfigurationOutput
	ToFrontendIPConfigurationOutputWithContext(context.Context) FrontendIPConfigurationOutput
}

type FrontendIPConfigurationArgs struct {
	Etag                      pulumi.StringPtrInput       `pulumi:"etag"`
	Id                        pulumi.StringPtrInput       `pulumi:"id"`
	Name                      pulumi.StringPtrInput       `pulumi:"name"`
	PrivateIPAddress          pulumi.StringPtrInput       `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput       `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringPtrInput       `pulumi:"provisioningState"`
	PublicIPAddress           PublicIPAddressTypePtrInput `pulumi:"publicIPAddress"`
	Subnet                    SubnetTypePtrInput          `pulumi:"subnet"`
	Zones                     pulumi.StringArrayInput     `pulumi:"zones"`
}

func (FrontendIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIPConfiguration)(nil)).Elem()
}

func (i FrontendIPConfigurationArgs) ToFrontendIPConfigurationOutput() FrontendIPConfigurationOutput {
	return i.ToFrontendIPConfigurationOutputWithContext(context.Background())
}

func (i FrontendIPConfigurationArgs) ToFrontendIPConfigurationOutputWithContext(ctx context.Context) FrontendIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendIPConfigurationOutput)
}





type FrontendIPConfigurationArrayInput interface {
	pulumi.Input

	ToFrontendIPConfigurationArrayOutput() FrontendIPConfigurationArrayOutput
	ToFrontendIPConfigurationArrayOutputWithContext(context.Context) FrontendIPConfigurationArrayOutput
}

type FrontendIPConfigurationArray []FrontendIPConfigurationInput

func (FrontendIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIPConfiguration)(nil)).Elem()
}

func (i FrontendIPConfigurationArray) ToFrontendIPConfigurationArrayOutput() FrontendIPConfigurationArrayOutput {
	return i.ToFrontendIPConfigurationArrayOutputWithContext(context.Background())
}

func (i FrontendIPConfigurationArray) ToFrontendIPConfigurationArrayOutputWithContext(ctx context.Context) FrontendIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendIPConfigurationArrayOutput)
}

type FrontendIPConfigurationOutput struct{ *pulumi.OutputState }

func (FrontendIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIPConfiguration)(nil)).Elem()
}

func (o FrontendIPConfigurationOutput) ToFrontendIPConfigurationOutput() FrontendIPConfigurationOutput {
	return o
}

func (o FrontendIPConfigurationOutput) ToFrontendIPConfigurationOutputWithContext(ctx context.Context) FrontendIPConfigurationOutput {
	return o
}

func (o FrontendIPConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) PublicIPAddress() PublicIPAddressTypePtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *PublicIPAddressType { return v.PublicIPAddress }).(PublicIPAddressTypePtrOutput)
}

func (o FrontendIPConfigurationOutput) Subnet() SubnetTypePtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *SubnetType { return v.Subnet }).(SubnetTypePtrOutput)
}

func (o FrontendIPConfigurationOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

type FrontendIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (FrontendIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIPConfiguration)(nil)).Elem()
}

func (o FrontendIPConfigurationArrayOutput) ToFrontendIPConfigurationArrayOutput() FrontendIPConfigurationArrayOutput {
	return o
}

func (o FrontendIPConfigurationArrayOutput) ToFrontendIPConfigurationArrayOutputWithContext(ctx context.Context) FrontendIPConfigurationArrayOutput {
	return o
}

func (o FrontendIPConfigurationArrayOutput) Index(i pulumi.IntInput) FrontendIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendIPConfiguration {
		return vs[0].([]FrontendIPConfiguration)[vs[1].(int)]
	}).(FrontendIPConfigurationOutput)
}

type FrontendIPConfigurationResponse struct {
	Etag                      *string                  `pulumi:"etag"`
	Id                        *string                  `pulumi:"id"`
	InboundNatPools           []SubResourceResponse    `pulumi:"inboundNatPools"`
	InboundNatRules           []SubResourceResponse    `pulumi:"inboundNatRules"`
	LoadBalancingRules        []SubResourceResponse    `pulumi:"loadBalancingRules"`
	Name                      *string                  `pulumi:"name"`
	OutboundNatRules          []SubResourceResponse    `pulumi:"outboundNatRules"`
	PrivateIPAddress          *string                  `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string                  `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string                  `pulumi:"provisioningState"`
	PublicIPAddress           *PublicIPAddressResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubnetResponse          `pulumi:"subnet"`
	Zones                     []string                 `pulumi:"zones"`
}





type FrontendIPConfigurationResponseInput interface {
	pulumi.Input

	ToFrontendIPConfigurationResponseOutput() FrontendIPConfigurationResponseOutput
	ToFrontendIPConfigurationResponseOutputWithContext(context.Context) FrontendIPConfigurationResponseOutput
}

type FrontendIPConfigurationResponseArgs struct {
	Etag                      pulumi.StringPtrInput           `pulumi:"etag"`
	Id                        pulumi.StringPtrInput           `pulumi:"id"`
	InboundNatPools           SubResourceResponseArrayInput   `pulumi:"inboundNatPools"`
	InboundNatRules           SubResourceResponseArrayInput   `pulumi:"inboundNatRules"`
	LoadBalancingRules        SubResourceResponseArrayInput   `pulumi:"loadBalancingRules"`
	Name                      pulumi.StringPtrInput           `pulumi:"name"`
	OutboundNatRules          SubResourceResponseArrayInput   `pulumi:"outboundNatRules"`
	PrivateIPAddress          pulumi.StringPtrInput           `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput           `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringPtrInput           `pulumi:"provisioningState"`
	PublicIPAddress           PublicIPAddressResponsePtrInput `pulumi:"publicIPAddress"`
	Subnet                    SubnetResponsePtrInput          `pulumi:"subnet"`
	Zones                     pulumi.StringArrayInput         `pulumi:"zones"`
}

func (FrontendIPConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIPConfigurationResponse)(nil)).Elem()
}

func (i FrontendIPConfigurationResponseArgs) ToFrontendIPConfigurationResponseOutput() FrontendIPConfigurationResponseOutput {
	return i.ToFrontendIPConfigurationResponseOutputWithContext(context.Background())
}

func (i FrontendIPConfigurationResponseArgs) ToFrontendIPConfigurationResponseOutputWithContext(ctx context.Context) FrontendIPConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendIPConfigurationResponseOutput)
}





type FrontendIPConfigurationResponseArrayInput interface {
	pulumi.Input

	ToFrontendIPConfigurationResponseArrayOutput() FrontendIPConfigurationResponseArrayOutput
	ToFrontendIPConfigurationResponseArrayOutputWithContext(context.Context) FrontendIPConfigurationResponseArrayOutput
}

type FrontendIPConfigurationResponseArray []FrontendIPConfigurationResponseInput

func (FrontendIPConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIPConfigurationResponse)(nil)).Elem()
}

func (i FrontendIPConfigurationResponseArray) ToFrontendIPConfigurationResponseArrayOutput() FrontendIPConfigurationResponseArrayOutput {
	return i.ToFrontendIPConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i FrontendIPConfigurationResponseArray) ToFrontendIPConfigurationResponseArrayOutputWithContext(ctx context.Context) FrontendIPConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendIPConfigurationResponseArrayOutput)
}

type FrontendIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FrontendIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIPConfigurationResponse)(nil)).Elem()
}

func (o FrontendIPConfigurationResponseOutput) ToFrontendIPConfigurationResponseOutput() FrontendIPConfigurationResponseOutput {
	return o
}

func (o FrontendIPConfigurationResponseOutput) ToFrontendIPConfigurationResponseOutputWithContext(ctx context.Context) FrontendIPConfigurationResponseOutput {
	return o
}

func (o FrontendIPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) InboundNatPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) []SubResourceResponse { return v.InboundNatPools }).(SubResourceResponseArrayOutput)
}

func (o FrontendIPConfigurationResponseOutput) InboundNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) []SubResourceResponse { return v.InboundNatRules }).(SubResourceResponseArrayOutput)
}

func (o FrontendIPConfigurationResponseOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) []SubResourceResponse { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

func (o FrontendIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) OutboundNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) []SubResourceResponse { return v.OutboundNatRules }).(SubResourceResponseArrayOutput)
}

func (o FrontendIPConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *PublicIPAddressResponse { return v.PublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

type FrontendIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontendIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIPConfigurationResponse)(nil)).Elem()
}

func (o FrontendIPConfigurationResponseArrayOutput) ToFrontendIPConfigurationResponseArrayOutput() FrontendIPConfigurationResponseArrayOutput {
	return o
}

func (o FrontendIPConfigurationResponseArrayOutput) ToFrontendIPConfigurationResponseArrayOutputWithContext(ctx context.Context) FrontendIPConfigurationResponseArrayOutput {
	return o
}

func (o FrontendIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) FrontendIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendIPConfigurationResponse {
		return vs[0].([]FrontendIPConfigurationResponse)[vs[1].(int)]
	}).(FrontendIPConfigurationResponseOutput)
}

type GatewayRouteResponse struct {
	AsPath       string `pulumi:"asPath"`
	LocalAddress string `pulumi:"localAddress"`
	Network      string `pulumi:"network"`
	NextHop      string `pulumi:"nextHop"`
	Origin       string `pulumi:"origin"`
	SourcePeer   string `pulumi:"sourcePeer"`
	Weight       int    `pulumi:"weight"`
}





type GatewayRouteResponseInput interface {
	pulumi.Input

	ToGatewayRouteResponseOutput() GatewayRouteResponseOutput
	ToGatewayRouteResponseOutputWithContext(context.Context) GatewayRouteResponseOutput
}

type GatewayRouteResponseArgs struct {
	AsPath       pulumi.StringInput `pulumi:"asPath"`
	LocalAddress pulumi.StringInput `pulumi:"localAddress"`
	Network      pulumi.StringInput `pulumi:"network"`
	NextHop      pulumi.StringInput `pulumi:"nextHop"`
	Origin       pulumi.StringInput `pulumi:"origin"`
	SourcePeer   pulumi.StringInput `pulumi:"sourcePeer"`
	Weight       pulumi.IntInput    `pulumi:"weight"`
}

func (GatewayRouteResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRouteResponse)(nil)).Elem()
}

func (i GatewayRouteResponseArgs) ToGatewayRouteResponseOutput() GatewayRouteResponseOutput {
	return i.ToGatewayRouteResponseOutputWithContext(context.Background())
}

func (i GatewayRouteResponseArgs) ToGatewayRouteResponseOutputWithContext(ctx context.Context) GatewayRouteResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteResponseOutput)
}





type GatewayRouteResponseArrayInput interface {
	pulumi.Input

	ToGatewayRouteResponseArrayOutput() GatewayRouteResponseArrayOutput
	ToGatewayRouteResponseArrayOutputWithContext(context.Context) GatewayRouteResponseArrayOutput
}

type GatewayRouteResponseArray []GatewayRouteResponseInput

func (GatewayRouteResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayRouteResponse)(nil)).Elem()
}

func (i GatewayRouteResponseArray) ToGatewayRouteResponseArrayOutput() GatewayRouteResponseArrayOutput {
	return i.ToGatewayRouteResponseArrayOutputWithContext(context.Background())
}

func (i GatewayRouteResponseArray) ToGatewayRouteResponseArrayOutputWithContext(ctx context.Context) GatewayRouteResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteResponseArrayOutput)
}

type GatewayRouteResponseOutput struct{ *pulumi.OutputState }

func (GatewayRouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRouteResponse)(nil)).Elem()
}

func (o GatewayRouteResponseOutput) ToGatewayRouteResponseOutput() GatewayRouteResponseOutput {
	return o
}

func (o GatewayRouteResponseOutput) ToGatewayRouteResponseOutputWithContext(ctx context.Context) GatewayRouteResponseOutput {
	return o
}

func (o GatewayRouteResponseOutput) AsPath() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.AsPath }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) LocalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.LocalAddress }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.Network }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) NextHop() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.NextHop }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.Origin }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) SourcePeer() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.SourcePeer }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v GatewayRouteResponse) int { return v.Weight }).(pulumi.IntOutput)
}

type GatewayRouteResponseArrayOutput struct{ *pulumi.OutputState }

func (GatewayRouteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayRouteResponse)(nil)).Elem()
}

func (o GatewayRouteResponseArrayOutput) ToGatewayRouteResponseArrayOutput() GatewayRouteResponseArrayOutput {
	return o
}

func (o GatewayRouteResponseArrayOutput) ToGatewayRouteResponseArrayOutputWithContext(ctx context.Context) GatewayRouteResponseArrayOutput {
	return o
}

func (o GatewayRouteResponseArrayOutput) Index(i pulumi.IntInput) GatewayRouteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GatewayRouteResponse {
		return vs[0].([]GatewayRouteResponse)[vs[1].(int)]
	}).(GatewayRouteResponseOutput)
}

type IPConfigurationResponse struct {
	Etag                      *string                  `pulumi:"etag"`
	Id                        *string                  `pulumi:"id"`
	Name                      *string                  `pulumi:"name"`
	PrivateIPAddress          *string                  `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string                  `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string                  `pulumi:"provisioningState"`
	PublicIPAddress           *PublicIPAddressResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubnetResponse          `pulumi:"subnet"`
}





type IPConfigurationResponseInput interface {
	pulumi.Input

	ToIPConfigurationResponseOutput() IPConfigurationResponseOutput
	ToIPConfigurationResponseOutputWithContext(context.Context) IPConfigurationResponseOutput
}

type IPConfigurationResponseArgs struct {
	Etag                      pulumi.StringPtrInput           `pulumi:"etag"`
	Id                        pulumi.StringPtrInput           `pulumi:"id"`
	Name                      pulumi.StringPtrInput           `pulumi:"name"`
	PrivateIPAddress          pulumi.StringPtrInput           `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput           `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringPtrInput           `pulumi:"provisioningState"`
	PublicIPAddress           PublicIPAddressResponsePtrInput `pulumi:"publicIPAddress"`
	Subnet                    SubnetResponsePtrInput          `pulumi:"subnet"`
}

func (IPConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPConfigurationResponse)(nil)).Elem()
}

func (i IPConfigurationResponseArgs) ToIPConfigurationResponseOutput() IPConfigurationResponseOutput {
	return i.ToIPConfigurationResponseOutputWithContext(context.Background())
}

func (i IPConfigurationResponseArgs) ToIPConfigurationResponseOutputWithContext(ctx context.Context) IPConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPConfigurationResponseOutput)
}

func (i IPConfigurationResponseArgs) ToIPConfigurationResponsePtrOutput() IPConfigurationResponsePtrOutput {
	return i.ToIPConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i IPConfigurationResponseArgs) ToIPConfigurationResponsePtrOutputWithContext(ctx context.Context) IPConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPConfigurationResponseOutput).ToIPConfigurationResponsePtrOutputWithContext(ctx)
}









type IPConfigurationResponsePtrInput interface {
	pulumi.Input

	ToIPConfigurationResponsePtrOutput() IPConfigurationResponsePtrOutput
	ToIPConfigurationResponsePtrOutputWithContext(context.Context) IPConfigurationResponsePtrOutput
}

type ipconfigurationResponsePtrType IPConfigurationResponseArgs

func IPConfigurationResponsePtr(v *IPConfigurationResponseArgs) IPConfigurationResponsePtrInput {
	return (*ipconfigurationResponsePtrType)(v)
}

func (*ipconfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IPConfigurationResponse)(nil)).Elem()
}

func (i *ipconfigurationResponsePtrType) ToIPConfigurationResponsePtrOutput() IPConfigurationResponsePtrOutput {
	return i.ToIPConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *ipconfigurationResponsePtrType) ToIPConfigurationResponsePtrOutputWithContext(ctx context.Context) IPConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPConfigurationResponsePtrOutput)
}





type IPConfigurationResponseArrayInput interface {
	pulumi.Input

	ToIPConfigurationResponseArrayOutput() IPConfigurationResponseArrayOutput
	ToIPConfigurationResponseArrayOutputWithContext(context.Context) IPConfigurationResponseArrayOutput
}

type IPConfigurationResponseArray []IPConfigurationResponseInput

func (IPConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPConfigurationResponse)(nil)).Elem()
}

func (i IPConfigurationResponseArray) ToIPConfigurationResponseArrayOutput() IPConfigurationResponseArrayOutput {
	return i.ToIPConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i IPConfigurationResponseArray) ToIPConfigurationResponseArrayOutputWithContext(ctx context.Context) IPConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPConfigurationResponseArrayOutput)
}

type IPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (IPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPConfigurationResponse)(nil)).Elem()
}

func (o IPConfigurationResponseOutput) ToIPConfigurationResponseOutput() IPConfigurationResponseOutput {
	return o
}

func (o IPConfigurationResponseOutput) ToIPConfigurationResponseOutputWithContext(ctx context.Context) IPConfigurationResponseOutput {
	return o
}

func (o IPConfigurationResponseOutput) ToIPConfigurationResponsePtrOutput() IPConfigurationResponsePtrOutput {
	return o.ToIPConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o IPConfigurationResponseOutput) ToIPConfigurationResponsePtrOutputWithContext(ctx context.Context) IPConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPConfigurationResponse) *IPConfigurationResponse {
		return &v
	}).(IPConfigurationResponsePtrOutput)
}

func (o IPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *PublicIPAddressResponse { return v.PublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

func (o IPConfigurationResponseOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

type IPConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (IPConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPConfigurationResponse)(nil)).Elem()
}

func (o IPConfigurationResponsePtrOutput) ToIPConfigurationResponsePtrOutput() IPConfigurationResponsePtrOutput {
	return o
}

func (o IPConfigurationResponsePtrOutput) ToIPConfigurationResponsePtrOutputWithContext(ctx context.Context) IPConfigurationResponsePtrOutput {
	return o
}

func (o IPConfigurationResponsePtrOutput) Elem() IPConfigurationResponseOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) IPConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret IPConfigurationResponse
		return ret
	}).(IPConfigurationResponseOutput)
}

func (o IPConfigurationResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAddress
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *PublicIPAddressResponse {
		if v == nil {
			return nil
		}
		return v.PublicIPAddress
	}).(PublicIPAddressResponsePtrOutput)
}

func (o IPConfigurationResponsePtrOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *SubnetResponse {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(SubnetResponsePtrOutput)
}

type IPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (IPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPConfigurationResponse)(nil)).Elem()
}

func (o IPConfigurationResponseArrayOutput) ToIPConfigurationResponseArrayOutput() IPConfigurationResponseArrayOutput {
	return o
}

func (o IPConfigurationResponseArrayOutput) ToIPConfigurationResponseArrayOutputWithContext(ctx context.Context) IPConfigurationResponseArrayOutput {
	return o
}

func (o IPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) IPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPConfigurationResponse {
		return vs[0].([]IPConfigurationResponse)[vs[1].(int)]
	}).(IPConfigurationResponseOutput)
}

type InboundNatPool struct {
	BackendPort             int          `pulumi:"backendPort"`
	EnableFloatingIP        *bool        `pulumi:"enableFloatingIP"`
	Etag                    *string      `pulumi:"etag"`
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    int          `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  int          `pulumi:"frontendPortRangeStart"`
	Id                      *string      `pulumi:"id"`
	IdleTimeoutInMinutes    *int         `pulumi:"idleTimeoutInMinutes"`
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
	EnableFloatingIP        pulumi.BoolPtrInput   `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    pulumi.IntInput       `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  pulumi.IntInput       `pulumi:"frontendPortRangeStart"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	IdleTimeoutInMinutes    pulumi.IntPtrInput    `pulumi:"idleTimeoutInMinutes"`
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

func (o InboundNatPoolOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InboundNatPool) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
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

func (o InboundNatPoolOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatPool) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
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
	EnableFloatingIP        *bool                `pulumi:"enableFloatingIP"`
	Etag                    *string              `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    int                  `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  int                  `pulumi:"frontendPortRangeStart"`
	Id                      *string              `pulumi:"id"`
	IdleTimeoutInMinutes    *int                 `pulumi:"idleTimeoutInMinutes"`
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
	EnableFloatingIP        pulumi.BoolPtrInput         `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrInput       `pulumi:"etag"`
	FrontendIPConfiguration SubResourceResponsePtrInput `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    pulumi.IntInput             `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  pulumi.IntInput             `pulumi:"frontendPortRangeStart"`
	Id                      pulumi.StringPtrInput       `pulumi:"id"`
	IdleTimeoutInMinutes    pulumi.IntPtrInput          `pulumi:"idleTimeoutInMinutes"`
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

func (o InboundNatPoolResponseOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
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

func (o InboundNatPoolResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
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

type InboundNatRuleType struct {
	BackendPort             *int         `pulumi:"backendPort"`
	EnableFloatingIP        *bool        `pulumi:"enableFloatingIP"`
	Etag                    *string      `pulumi:"etag"`
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPort            *int         `pulumi:"frontendPort"`
	Id                      *string      `pulumi:"id"`
	IdleTimeoutInMinutes    *int         `pulumi:"idleTimeoutInMinutes"`
	Name                    *string      `pulumi:"name"`
	Protocol                *string      `pulumi:"protocol"`
	ProvisioningState       *string      `pulumi:"provisioningState"`
}





type InboundNatRuleTypeInput interface {
	pulumi.Input

	ToInboundNatRuleTypeOutput() InboundNatRuleTypeOutput
	ToInboundNatRuleTypeOutputWithContext(context.Context) InboundNatRuleTypeOutput
}

type InboundNatRuleTypeArgs struct {
	BackendPort             pulumi.IntPtrInput    `pulumi:"backendPort"`
	EnableFloatingIP        pulumi.BoolPtrInput   `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPort            pulumi.IntPtrInput    `pulumi:"frontendPort"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	IdleTimeoutInMinutes    pulumi.IntPtrInput    `pulumi:"idleTimeoutInMinutes"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	Protocol                pulumi.StringPtrInput `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (InboundNatRuleTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRuleType)(nil)).Elem()
}

func (i InboundNatRuleTypeArgs) ToInboundNatRuleTypeOutput() InboundNatRuleTypeOutput {
	return i.ToInboundNatRuleTypeOutputWithContext(context.Background())
}

func (i InboundNatRuleTypeArgs) ToInboundNatRuleTypeOutputWithContext(ctx context.Context) InboundNatRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleTypeOutput)
}





type InboundNatRuleTypeArrayInput interface {
	pulumi.Input

	ToInboundNatRuleTypeArrayOutput() InboundNatRuleTypeArrayOutput
	ToInboundNatRuleTypeArrayOutputWithContext(context.Context) InboundNatRuleTypeArrayOutput
}

type InboundNatRuleTypeArray []InboundNatRuleTypeInput

func (InboundNatRuleTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRuleType)(nil)).Elem()
}

func (i InboundNatRuleTypeArray) ToInboundNatRuleTypeArrayOutput() InboundNatRuleTypeArrayOutput {
	return i.ToInboundNatRuleTypeArrayOutputWithContext(context.Background())
}

func (i InboundNatRuleTypeArray) ToInboundNatRuleTypeArrayOutputWithContext(ctx context.Context) InboundNatRuleTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleTypeArrayOutput)
}

type InboundNatRuleTypeOutput struct{ *pulumi.OutputState }

func (InboundNatRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRuleType)(nil)).Elem()
}

func (o InboundNatRuleTypeOutput) ToInboundNatRuleTypeOutput() InboundNatRuleTypeOutput {
	return o
}

func (o InboundNatRuleTypeOutput) ToInboundNatRuleTypeOutputWithContext(ctx context.Context) InboundNatRuleTypeOutput {
	return o
}

func (o InboundNatRuleTypeOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleType) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleTypeOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InboundNatRuleType) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
}

func (o InboundNatRuleTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleTypeOutput) FrontendIPConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v InboundNatRuleType) *SubResource { return v.FrontendIPConfiguration }).(SubResourcePtrOutput)
}

func (o InboundNatRuleTypeOutput) FrontendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleType) *int { return v.FrontendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleTypeOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleType) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleTypeOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleType) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type InboundNatRuleTypeArrayOutput struct{ *pulumi.OutputState }

func (InboundNatRuleTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRuleType)(nil)).Elem()
}

func (o InboundNatRuleTypeArrayOutput) ToInboundNatRuleTypeArrayOutput() InboundNatRuleTypeArrayOutput {
	return o
}

func (o InboundNatRuleTypeArrayOutput) ToInboundNatRuleTypeArrayOutputWithContext(ctx context.Context) InboundNatRuleTypeArrayOutput {
	return o
}

func (o InboundNatRuleTypeArrayOutput) Index(i pulumi.IntInput) InboundNatRuleTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatRuleType {
		return vs[0].([]InboundNatRuleType)[vs[1].(int)]
	}).(InboundNatRuleTypeOutput)
}

type InboundNatRuleResponse struct {
	BackendIPConfiguration  NetworkInterfaceIPConfigurationResponse `pulumi:"backendIPConfiguration"`
	BackendPort             *int                                    `pulumi:"backendPort"`
	EnableFloatingIP        *bool                                   `pulumi:"enableFloatingIP"`
	Etag                    *string                                 `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse                    `pulumi:"frontendIPConfiguration"`
	FrontendPort            *int                                    `pulumi:"frontendPort"`
	Id                      *string                                 `pulumi:"id"`
	IdleTimeoutInMinutes    *int                                    `pulumi:"idleTimeoutInMinutes"`
	Name                    *string                                 `pulumi:"name"`
	Protocol                *string                                 `pulumi:"protocol"`
	ProvisioningState       *string                                 `pulumi:"provisioningState"`
}





type InboundNatRuleResponseInput interface {
	pulumi.Input

	ToInboundNatRuleResponseOutput() InboundNatRuleResponseOutput
	ToInboundNatRuleResponseOutputWithContext(context.Context) InboundNatRuleResponseOutput
}

type InboundNatRuleResponseArgs struct {
	BackendIPConfiguration  NetworkInterfaceIPConfigurationResponseInput `pulumi:"backendIPConfiguration"`
	BackendPort             pulumi.IntPtrInput                           `pulumi:"backendPort"`
	EnableFloatingIP        pulumi.BoolPtrInput                          `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrInput                        `pulumi:"etag"`
	FrontendIPConfiguration SubResourceResponsePtrInput                  `pulumi:"frontendIPConfiguration"`
	FrontendPort            pulumi.IntPtrInput                           `pulumi:"frontendPort"`
	Id                      pulumi.StringPtrInput                        `pulumi:"id"`
	IdleTimeoutInMinutes    pulumi.IntPtrInput                           `pulumi:"idleTimeoutInMinutes"`
	Name                    pulumi.StringPtrInput                        `pulumi:"name"`
	Protocol                pulumi.StringPtrInput                        `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput                        `pulumi:"provisioningState"`
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

func (o InboundNatRuleResponseOutput) BackendIPConfiguration() NetworkInterfaceIPConfigurationResponseOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) NetworkInterfaceIPConfigurationResponse {
		return v.BackendIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponseOutput)
}

func (o InboundNatRuleResponseOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleResponseOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
}

func (o InboundNatRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleResponseOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o InboundNatRuleResponseOutput) FrontendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *int { return v.FrontendPort }).(pulumi.IntPtrOutput)
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

func (o InboundNatRuleResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
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

type IpTag struct {
	IpTagType *string `pulumi:"ipTagType"`
	Tag       *string `pulumi:"tag"`
}





type IpTagInput interface {
	pulumi.Input

	ToIpTagOutput() IpTagOutput
	ToIpTagOutputWithContext(context.Context) IpTagOutput
}

type IpTagArgs struct {
	IpTagType pulumi.StringPtrInput `pulumi:"ipTagType"`
	Tag       pulumi.StringPtrInput `pulumi:"tag"`
}

func (IpTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpTag)(nil)).Elem()
}

func (i IpTagArgs) ToIpTagOutput() IpTagOutput {
	return i.ToIpTagOutputWithContext(context.Background())
}

func (i IpTagArgs) ToIpTagOutputWithContext(ctx context.Context) IpTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpTagOutput)
}





type IpTagArrayInput interface {
	pulumi.Input

	ToIpTagArrayOutput() IpTagArrayOutput
	ToIpTagArrayOutputWithContext(context.Context) IpTagArrayOutput
}

type IpTagArray []IpTagInput

func (IpTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpTag)(nil)).Elem()
}

func (i IpTagArray) ToIpTagArrayOutput() IpTagArrayOutput {
	return i.ToIpTagArrayOutputWithContext(context.Background())
}

func (i IpTagArray) ToIpTagArrayOutputWithContext(ctx context.Context) IpTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpTagArrayOutput)
}

type IpTagOutput struct{ *pulumi.OutputState }

func (IpTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpTag)(nil)).Elem()
}

func (o IpTagOutput) ToIpTagOutput() IpTagOutput {
	return o
}

func (o IpTagOutput) ToIpTagOutputWithContext(ctx context.Context) IpTagOutput {
	return o
}

func (o IpTagOutput) IpTagType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpTag) *string { return v.IpTagType }).(pulumi.StringPtrOutput)
}

func (o IpTagOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpTag) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type IpTagArrayOutput struct{ *pulumi.OutputState }

func (IpTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpTag)(nil)).Elem()
}

func (o IpTagArrayOutput) ToIpTagArrayOutput() IpTagArrayOutput {
	return o
}

func (o IpTagArrayOutput) ToIpTagArrayOutputWithContext(ctx context.Context) IpTagArrayOutput {
	return o
}

func (o IpTagArrayOutput) Index(i pulumi.IntInput) IpTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpTag {
		return vs[0].([]IpTag)[vs[1].(int)]
	}).(IpTagOutput)
}

type IpTagResponse struct {
	IpTagType *string `pulumi:"ipTagType"`
	Tag       *string `pulumi:"tag"`
}





type IpTagResponseInput interface {
	pulumi.Input

	ToIpTagResponseOutput() IpTagResponseOutput
	ToIpTagResponseOutputWithContext(context.Context) IpTagResponseOutput
}

type IpTagResponseArgs struct {
	IpTagType pulumi.StringPtrInput `pulumi:"ipTagType"`
	Tag       pulumi.StringPtrInput `pulumi:"tag"`
}

func (IpTagResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpTagResponse)(nil)).Elem()
}

func (i IpTagResponseArgs) ToIpTagResponseOutput() IpTagResponseOutput {
	return i.ToIpTagResponseOutputWithContext(context.Background())
}

func (i IpTagResponseArgs) ToIpTagResponseOutputWithContext(ctx context.Context) IpTagResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpTagResponseOutput)
}





type IpTagResponseArrayInput interface {
	pulumi.Input

	ToIpTagResponseArrayOutput() IpTagResponseArrayOutput
	ToIpTagResponseArrayOutputWithContext(context.Context) IpTagResponseArrayOutput
}

type IpTagResponseArray []IpTagResponseInput

func (IpTagResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpTagResponse)(nil)).Elem()
}

func (i IpTagResponseArray) ToIpTagResponseArrayOutput() IpTagResponseArrayOutput {
	return i.ToIpTagResponseArrayOutputWithContext(context.Background())
}

func (i IpTagResponseArray) ToIpTagResponseArrayOutputWithContext(ctx context.Context) IpTagResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpTagResponseArrayOutput)
}

type IpTagResponseOutput struct{ *pulumi.OutputState }

func (IpTagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpTagResponse)(nil)).Elem()
}

func (o IpTagResponseOutput) ToIpTagResponseOutput() IpTagResponseOutput {
	return o
}

func (o IpTagResponseOutput) ToIpTagResponseOutputWithContext(ctx context.Context) IpTagResponseOutput {
	return o
}

func (o IpTagResponseOutput) IpTagType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpTagResponse) *string { return v.IpTagType }).(pulumi.StringPtrOutput)
}

func (o IpTagResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpTagResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type IpTagResponseArrayOutput struct{ *pulumi.OutputState }

func (IpTagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpTagResponse)(nil)).Elem()
}

func (o IpTagResponseArrayOutput) ToIpTagResponseArrayOutput() IpTagResponseArrayOutput {
	return o
}

func (o IpTagResponseArrayOutput) ToIpTagResponseArrayOutputWithContext(ctx context.Context) IpTagResponseArrayOutput {
	return o
}

func (o IpTagResponseArrayOutput) Index(i pulumi.IntInput) IpTagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpTagResponse {
		return vs[0].([]IpTagResponse)[vs[1].(int)]
	}).(IpTagResponseOutput)
}

type IpsecPolicy struct {
	DhGroup             string `pulumi:"dhGroup"`
	IkeEncryption       string `pulumi:"ikeEncryption"`
	IkeIntegrity        string `pulumi:"ikeIntegrity"`
	IpsecEncryption     string `pulumi:"ipsecEncryption"`
	IpsecIntegrity      string `pulumi:"ipsecIntegrity"`
	PfsGroup            string `pulumi:"pfsGroup"`
	SaDataSizeKilobytes int    `pulumi:"saDataSizeKilobytes"`
	SaLifeTimeSeconds   int    `pulumi:"saLifeTimeSeconds"`
}





type IpsecPolicyInput interface {
	pulumi.Input

	ToIpsecPolicyOutput() IpsecPolicyOutput
	ToIpsecPolicyOutputWithContext(context.Context) IpsecPolicyOutput
}

type IpsecPolicyArgs struct {
	DhGroup             pulumi.StringInput `pulumi:"dhGroup"`
	IkeEncryption       pulumi.StringInput `pulumi:"ikeEncryption"`
	IkeIntegrity        pulumi.StringInput `pulumi:"ikeIntegrity"`
	IpsecEncryption     pulumi.StringInput `pulumi:"ipsecEncryption"`
	IpsecIntegrity      pulumi.StringInput `pulumi:"ipsecIntegrity"`
	PfsGroup            pulumi.StringInput `pulumi:"pfsGroup"`
	SaDataSizeKilobytes pulumi.IntInput    `pulumi:"saDataSizeKilobytes"`
	SaLifeTimeSeconds   pulumi.IntInput    `pulumi:"saLifeTimeSeconds"`
}

func (IpsecPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecPolicy)(nil)).Elem()
}

func (i IpsecPolicyArgs) ToIpsecPolicyOutput() IpsecPolicyOutput {
	return i.ToIpsecPolicyOutputWithContext(context.Background())
}

func (i IpsecPolicyArgs) ToIpsecPolicyOutputWithContext(ctx context.Context) IpsecPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecPolicyOutput)
}





type IpsecPolicyArrayInput interface {
	pulumi.Input

	ToIpsecPolicyArrayOutput() IpsecPolicyArrayOutput
	ToIpsecPolicyArrayOutputWithContext(context.Context) IpsecPolicyArrayOutput
}

type IpsecPolicyArray []IpsecPolicyInput

func (IpsecPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpsecPolicy)(nil)).Elem()
}

func (i IpsecPolicyArray) ToIpsecPolicyArrayOutput() IpsecPolicyArrayOutput {
	return i.ToIpsecPolicyArrayOutputWithContext(context.Background())
}

func (i IpsecPolicyArray) ToIpsecPolicyArrayOutputWithContext(ctx context.Context) IpsecPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecPolicyArrayOutput)
}

type IpsecPolicyOutput struct{ *pulumi.OutputState }

func (IpsecPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecPolicy)(nil)).Elem()
}

func (o IpsecPolicyOutput) ToIpsecPolicyOutput() IpsecPolicyOutput {
	return o
}

func (o IpsecPolicyOutput) ToIpsecPolicyOutputWithContext(ctx context.Context) IpsecPolicyOutput {
	return o
}

func (o IpsecPolicyOutput) DhGroup() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.DhGroup }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) IkeEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.IkeEncryption }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) IkeIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.IkeIntegrity }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) IpsecEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.IpsecEncryption }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) IpsecIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.IpsecIntegrity }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) PfsGroup() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.PfsGroup }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) SaDataSizeKilobytes() pulumi.IntOutput {
	return o.ApplyT(func(v IpsecPolicy) int { return v.SaDataSizeKilobytes }).(pulumi.IntOutput)
}

func (o IpsecPolicyOutput) SaLifeTimeSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v IpsecPolicy) int { return v.SaLifeTimeSeconds }).(pulumi.IntOutput)
}

type IpsecPolicyArrayOutput struct{ *pulumi.OutputState }

func (IpsecPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpsecPolicy)(nil)).Elem()
}

func (o IpsecPolicyArrayOutput) ToIpsecPolicyArrayOutput() IpsecPolicyArrayOutput {
	return o
}

func (o IpsecPolicyArrayOutput) ToIpsecPolicyArrayOutputWithContext(ctx context.Context) IpsecPolicyArrayOutput {
	return o
}

func (o IpsecPolicyArrayOutput) Index(i pulumi.IntInput) IpsecPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpsecPolicy {
		return vs[0].([]IpsecPolicy)[vs[1].(int)]
	}).(IpsecPolicyOutput)
}

type IpsecPolicyResponse struct {
	DhGroup             string `pulumi:"dhGroup"`
	IkeEncryption       string `pulumi:"ikeEncryption"`
	IkeIntegrity        string `pulumi:"ikeIntegrity"`
	IpsecEncryption     string `pulumi:"ipsecEncryption"`
	IpsecIntegrity      string `pulumi:"ipsecIntegrity"`
	PfsGroup            string `pulumi:"pfsGroup"`
	SaDataSizeKilobytes int    `pulumi:"saDataSizeKilobytes"`
	SaLifeTimeSeconds   int    `pulumi:"saLifeTimeSeconds"`
}





type IpsecPolicyResponseInput interface {
	pulumi.Input

	ToIpsecPolicyResponseOutput() IpsecPolicyResponseOutput
	ToIpsecPolicyResponseOutputWithContext(context.Context) IpsecPolicyResponseOutput
}

type IpsecPolicyResponseArgs struct {
	DhGroup             pulumi.StringInput `pulumi:"dhGroup"`
	IkeEncryption       pulumi.StringInput `pulumi:"ikeEncryption"`
	IkeIntegrity        pulumi.StringInput `pulumi:"ikeIntegrity"`
	IpsecEncryption     pulumi.StringInput `pulumi:"ipsecEncryption"`
	IpsecIntegrity      pulumi.StringInput `pulumi:"ipsecIntegrity"`
	PfsGroup            pulumi.StringInput `pulumi:"pfsGroup"`
	SaDataSizeKilobytes pulumi.IntInput    `pulumi:"saDataSizeKilobytes"`
	SaLifeTimeSeconds   pulumi.IntInput    `pulumi:"saLifeTimeSeconds"`
}

func (IpsecPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecPolicyResponse)(nil)).Elem()
}

func (i IpsecPolicyResponseArgs) ToIpsecPolicyResponseOutput() IpsecPolicyResponseOutput {
	return i.ToIpsecPolicyResponseOutputWithContext(context.Background())
}

func (i IpsecPolicyResponseArgs) ToIpsecPolicyResponseOutputWithContext(ctx context.Context) IpsecPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecPolicyResponseOutput)
}





type IpsecPolicyResponseArrayInput interface {
	pulumi.Input

	ToIpsecPolicyResponseArrayOutput() IpsecPolicyResponseArrayOutput
	ToIpsecPolicyResponseArrayOutputWithContext(context.Context) IpsecPolicyResponseArrayOutput
}

type IpsecPolicyResponseArray []IpsecPolicyResponseInput

func (IpsecPolicyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpsecPolicyResponse)(nil)).Elem()
}

func (i IpsecPolicyResponseArray) ToIpsecPolicyResponseArrayOutput() IpsecPolicyResponseArrayOutput {
	return i.ToIpsecPolicyResponseArrayOutputWithContext(context.Background())
}

func (i IpsecPolicyResponseArray) ToIpsecPolicyResponseArrayOutputWithContext(ctx context.Context) IpsecPolicyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecPolicyResponseArrayOutput)
}

type IpsecPolicyResponseOutput struct{ *pulumi.OutputState }

func (IpsecPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecPolicyResponse)(nil)).Elem()
}

func (o IpsecPolicyResponseOutput) ToIpsecPolicyResponseOutput() IpsecPolicyResponseOutput {
	return o
}

func (o IpsecPolicyResponseOutput) ToIpsecPolicyResponseOutputWithContext(ctx context.Context) IpsecPolicyResponseOutput {
	return o
}

func (o IpsecPolicyResponseOutput) DhGroup() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.DhGroup }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) IkeEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.IkeEncryption }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) IkeIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.IkeIntegrity }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) IpsecEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.IpsecEncryption }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) IpsecIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.IpsecIntegrity }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) PfsGroup() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.PfsGroup }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) SaDataSizeKilobytes() pulumi.IntOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) int { return v.SaDataSizeKilobytes }).(pulumi.IntOutput)
}

func (o IpsecPolicyResponseOutput) SaLifeTimeSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) int { return v.SaLifeTimeSeconds }).(pulumi.IntOutput)
}

type IpsecPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (IpsecPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpsecPolicyResponse)(nil)).Elem()
}

func (o IpsecPolicyResponseArrayOutput) ToIpsecPolicyResponseArrayOutput() IpsecPolicyResponseArrayOutput {
	return o
}

func (o IpsecPolicyResponseArrayOutput) ToIpsecPolicyResponseArrayOutputWithContext(ctx context.Context) IpsecPolicyResponseArrayOutput {
	return o
}

func (o IpsecPolicyResponseArrayOutput) Index(i pulumi.IntInput) IpsecPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpsecPolicyResponse {
		return vs[0].([]IpsecPolicyResponse)[vs[1].(int)]
	}).(IpsecPolicyResponseOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfig struct {
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfig `pulumi:"microsoftPeeringConfig"`
	PrimaryPeerAddressPrefix   *string                           `pulumi:"primaryPeerAddressPrefix"`
	RouteFilter                *RouteFilterType                  `pulumi:"routeFilter"`
	SecondaryPeerAddressPrefix *string                           `pulumi:"secondaryPeerAddressPrefix"`
	State                      *string                           `pulumi:"state"`
}





type Ipv6ExpressRouteCircuitPeeringConfigInput interface {
	pulumi.Input

	ToIpv6ExpressRouteCircuitPeeringConfigOutput() Ipv6ExpressRouteCircuitPeeringConfigOutput
	ToIpv6ExpressRouteCircuitPeeringConfigOutputWithContext(context.Context) Ipv6ExpressRouteCircuitPeeringConfigOutput
}

type Ipv6ExpressRouteCircuitPeeringConfigArgs struct {
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigPtrInput `pulumi:"microsoftPeeringConfig"`
	PrimaryPeerAddressPrefix   pulumi.StringPtrInput                    `pulumi:"primaryPeerAddressPrefix"`
	RouteFilter                RouteFilterTypePtrInput                  `pulumi:"routeFilter"`
	SecondaryPeerAddressPrefix pulumi.StringPtrInput                    `pulumi:"secondaryPeerAddressPrefix"`
	State                      pulumi.StringPtrInput                    `pulumi:"state"`
}

func (Ipv6ExpressRouteCircuitPeeringConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (i Ipv6ExpressRouteCircuitPeeringConfigArgs) ToIpv6ExpressRouteCircuitPeeringConfigOutput() Ipv6ExpressRouteCircuitPeeringConfigOutput {
	return i.ToIpv6ExpressRouteCircuitPeeringConfigOutputWithContext(context.Background())
}

func (i Ipv6ExpressRouteCircuitPeeringConfigArgs) ToIpv6ExpressRouteCircuitPeeringConfigOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6ExpressRouteCircuitPeeringConfigOutput)
}

func (i Ipv6ExpressRouteCircuitPeeringConfigArgs) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutput() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return i.ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (i Ipv6ExpressRouteCircuitPeeringConfigArgs) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6ExpressRouteCircuitPeeringConfigOutput).ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx)
}









type Ipv6ExpressRouteCircuitPeeringConfigPtrInput interface {
	pulumi.Input

	ToIpv6ExpressRouteCircuitPeeringConfigPtrOutput() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput
	ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Context) Ipv6ExpressRouteCircuitPeeringConfigPtrOutput
}

type ipv6ExpressRouteCircuitPeeringConfigPtrType Ipv6ExpressRouteCircuitPeeringConfigArgs

func Ipv6ExpressRouteCircuitPeeringConfigPtr(v *Ipv6ExpressRouteCircuitPeeringConfigArgs) Ipv6ExpressRouteCircuitPeeringConfigPtrInput {
	return (*ipv6ExpressRouteCircuitPeeringConfigPtrType)(v)
}

func (*ipv6ExpressRouteCircuitPeeringConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (i *ipv6ExpressRouteCircuitPeeringConfigPtrType) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutput() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return i.ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (i *ipv6ExpressRouteCircuitPeeringConfigPtrType) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6ExpressRouteCircuitPeeringConfigPtrOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfigOutput struct{ *pulumi.OutputState }

func (Ipv6ExpressRouteCircuitPeeringConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) ToIpv6ExpressRouteCircuitPeeringConfigOutput() Ipv6ExpressRouteCircuitPeeringConfigOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) ToIpv6ExpressRouteCircuitPeeringConfigOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutput() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ipv6ExpressRouteCircuitPeeringConfig) *Ipv6ExpressRouteCircuitPeeringConfig {
		return &v
	}).(Ipv6ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfig) *ExpressRouteCircuitPeeringConfig {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfig) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) RouteFilter() RouteFilterTypePtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfig) *RouteFilterType { return v.RouteFilter }).(RouteFilterTypePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfig) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfig) *string { return v.State }).(pulumi.StringPtrOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfigPtrOutput struct{ *pulumi.OutputState }

func (Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutput() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) Elem() Ipv6ExpressRouteCircuitPeeringConfigOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) Ipv6ExpressRouteCircuitPeeringConfig {
		if v != nil {
			return *v
		}
		var ret Ipv6ExpressRouteCircuitPeeringConfig
		return ret
	}).(Ipv6ExpressRouteCircuitPeeringConfigOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) *ExpressRouteCircuitPeeringConfig {
		if v == nil {
			return nil
		}
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryPeerAddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) RouteFilter() RouteFilterTypePtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) *RouteFilterType {
		if v == nil {
			return nil
		}
		return v.RouteFilter
	}).(RouteFilterTypePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryPeerAddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfigResponse struct {
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfigResponse `pulumi:"microsoftPeeringConfig"`
	PrimaryPeerAddressPrefix   *string                                   `pulumi:"primaryPeerAddressPrefix"`
	RouteFilter                *RouteFilterResponse                      `pulumi:"routeFilter"`
	SecondaryPeerAddressPrefix *string                                   `pulumi:"secondaryPeerAddressPrefix"`
	State                      *string                                   `pulumi:"state"`
}





type Ipv6ExpressRouteCircuitPeeringConfigResponseInput interface {
	pulumi.Input

	ToIpv6ExpressRouteCircuitPeeringConfigResponseOutput() Ipv6ExpressRouteCircuitPeeringConfigResponseOutput
	ToIpv6ExpressRouteCircuitPeeringConfigResponseOutputWithContext(context.Context) Ipv6ExpressRouteCircuitPeeringConfigResponseOutput
}

type Ipv6ExpressRouteCircuitPeeringConfigResponseArgs struct {
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigResponsePtrInput `pulumi:"microsoftPeeringConfig"`
	PrimaryPeerAddressPrefix   pulumi.StringPtrInput                            `pulumi:"primaryPeerAddressPrefix"`
	RouteFilter                RouteFilterResponsePtrInput                      `pulumi:"routeFilter"`
	SecondaryPeerAddressPrefix pulumi.StringPtrInput                            `pulumi:"secondaryPeerAddressPrefix"`
	State                      pulumi.StringPtrInput                            `pulumi:"state"`
}

func (Ipv6ExpressRouteCircuitPeeringConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (i Ipv6ExpressRouteCircuitPeeringConfigResponseArgs) ToIpv6ExpressRouteCircuitPeeringConfigResponseOutput() Ipv6ExpressRouteCircuitPeeringConfigResponseOutput {
	return i.ToIpv6ExpressRouteCircuitPeeringConfigResponseOutputWithContext(context.Background())
}

func (i Ipv6ExpressRouteCircuitPeeringConfigResponseArgs) ToIpv6ExpressRouteCircuitPeeringConfigResponseOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6ExpressRouteCircuitPeeringConfigResponseOutput)
}

func (i Ipv6ExpressRouteCircuitPeeringConfigResponseArgs) ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutput() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return i.ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(context.Background())
}

func (i Ipv6ExpressRouteCircuitPeeringConfigResponseArgs) ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6ExpressRouteCircuitPeeringConfigResponseOutput).ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx)
}









type Ipv6ExpressRouteCircuitPeeringConfigResponsePtrInput interface {
	pulumi.Input

	ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutput() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput
	ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(context.Context) Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput
}

type ipv6ExpressRouteCircuitPeeringConfigResponsePtrType Ipv6ExpressRouteCircuitPeeringConfigResponseArgs

func Ipv6ExpressRouteCircuitPeeringConfigResponsePtr(v *Ipv6ExpressRouteCircuitPeeringConfigResponseArgs) Ipv6ExpressRouteCircuitPeeringConfigResponsePtrInput {
	return (*ipv6ExpressRouteCircuitPeeringConfigResponsePtrType)(v)
}

func (*ipv6ExpressRouteCircuitPeeringConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (i *ipv6ExpressRouteCircuitPeeringConfigResponsePtrType) ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutput() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return i.ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(context.Background())
}

func (i *ipv6ExpressRouteCircuitPeeringConfigResponsePtrType) ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfigResponseOutput struct{ *pulumi.OutputState }

func (Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) ToIpv6ExpressRouteCircuitPeeringConfigResponseOutput() Ipv6ExpressRouteCircuitPeeringConfigResponseOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) ToIpv6ExpressRouteCircuitPeeringConfigResponseOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigResponseOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutput() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(context.Background())
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ipv6ExpressRouteCircuitPeeringConfigResponse) *Ipv6ExpressRouteCircuitPeeringConfigResponse {
		return &v
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfigResponse) *ExpressRouteCircuitPeeringConfigResponse {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfigResponse) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) RouteFilter() RouteFilterResponsePtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfigResponse) *RouteFilterResponse { return v.RouteFilter }).(RouteFilterResponsePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfigResponse) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfigResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutput() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) Elem() Ipv6ExpressRouteCircuitPeeringConfigResponseOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) Ipv6ExpressRouteCircuitPeeringConfigResponse {
		if v != nil {
			return *v
		}
		var ret Ipv6ExpressRouteCircuitPeeringConfigResponse
		return ret
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponseOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) *ExpressRouteCircuitPeeringConfigResponse {
		if v == nil {
			return nil
		}
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryPeerAddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) RouteFilter() RouteFilterResponsePtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) *RouteFilterResponse {
		if v == nil {
			return nil
		}
		return v.RouteFilter
	}).(RouteFilterResponsePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryPeerAddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerSku struct {
	Name *string `pulumi:"name"`
}





type LoadBalancerSkuInput interface {
	pulumi.Input

	ToLoadBalancerSkuOutput() LoadBalancerSkuOutput
	ToLoadBalancerSkuOutputWithContext(context.Context) LoadBalancerSkuOutput
}

type LoadBalancerSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LoadBalancerSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSku)(nil)).Elem()
}

func (i LoadBalancerSkuArgs) ToLoadBalancerSkuOutput() LoadBalancerSkuOutput {
	return i.ToLoadBalancerSkuOutputWithContext(context.Background())
}

func (i LoadBalancerSkuArgs) ToLoadBalancerSkuOutputWithContext(ctx context.Context) LoadBalancerSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerSkuOutput)
}

func (i LoadBalancerSkuArgs) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return i.ToLoadBalancerSkuPtrOutputWithContext(context.Background())
}

func (i LoadBalancerSkuArgs) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerSkuOutput).ToLoadBalancerSkuPtrOutputWithContext(ctx)
}









type LoadBalancerSkuPtrInput interface {
	pulumi.Input

	ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput
	ToLoadBalancerSkuPtrOutputWithContext(context.Context) LoadBalancerSkuPtrOutput
}

type loadBalancerSkuPtrType LoadBalancerSkuArgs

func LoadBalancerSkuPtr(v *LoadBalancerSkuArgs) LoadBalancerSkuPtrInput {
	return (*loadBalancerSkuPtrType)(v)
}

func (*loadBalancerSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerSku)(nil)).Elem()
}

func (i *loadBalancerSkuPtrType) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return i.ToLoadBalancerSkuPtrOutputWithContext(context.Background())
}

func (i *loadBalancerSkuPtrType) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerSkuPtrOutput)
}

type LoadBalancerSkuOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSku)(nil)).Elem()
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuOutput() LoadBalancerSkuOutput {
	return o
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuOutputWithContext(ctx context.Context) LoadBalancerSkuOutput {
	return o
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return o.ToLoadBalancerSkuPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerSku) *LoadBalancerSku {
		return &v
	}).(LoadBalancerSkuPtrOutput)
}

func (o LoadBalancerSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type LoadBalancerSkuPtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerSku)(nil)).Elem()
}

func (o LoadBalancerSkuPtrOutput) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return o
}

func (o LoadBalancerSkuPtrOutput) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return o
}

func (o LoadBalancerSkuPtrOutput) Elem() LoadBalancerSkuOutput {
	return o.ApplyT(func(v *LoadBalancerSku) LoadBalancerSku {
		if v != nil {
			return *v
		}
		var ret LoadBalancerSku
		return ret
	}).(LoadBalancerSkuOutput)
}

func (o LoadBalancerSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerSkuResponse struct {
	Name *string `pulumi:"name"`
}





type LoadBalancerSkuResponseInput interface {
	pulumi.Input

	ToLoadBalancerSkuResponseOutput() LoadBalancerSkuResponseOutput
	ToLoadBalancerSkuResponseOutputWithContext(context.Context) LoadBalancerSkuResponseOutput
}

type LoadBalancerSkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LoadBalancerSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSkuResponse)(nil)).Elem()
}

func (i LoadBalancerSkuResponseArgs) ToLoadBalancerSkuResponseOutput() LoadBalancerSkuResponseOutput {
	return i.ToLoadBalancerSkuResponseOutputWithContext(context.Background())
}

func (i LoadBalancerSkuResponseArgs) ToLoadBalancerSkuResponseOutputWithContext(ctx context.Context) LoadBalancerSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerSkuResponseOutput)
}

func (i LoadBalancerSkuResponseArgs) ToLoadBalancerSkuResponsePtrOutput() LoadBalancerSkuResponsePtrOutput {
	return i.ToLoadBalancerSkuResponsePtrOutputWithContext(context.Background())
}

func (i LoadBalancerSkuResponseArgs) ToLoadBalancerSkuResponsePtrOutputWithContext(ctx context.Context) LoadBalancerSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerSkuResponseOutput).ToLoadBalancerSkuResponsePtrOutputWithContext(ctx)
}









type LoadBalancerSkuResponsePtrInput interface {
	pulumi.Input

	ToLoadBalancerSkuResponsePtrOutput() LoadBalancerSkuResponsePtrOutput
	ToLoadBalancerSkuResponsePtrOutputWithContext(context.Context) LoadBalancerSkuResponsePtrOutput
}

type loadBalancerSkuResponsePtrType LoadBalancerSkuResponseArgs

func LoadBalancerSkuResponsePtr(v *LoadBalancerSkuResponseArgs) LoadBalancerSkuResponsePtrInput {
	return (*loadBalancerSkuResponsePtrType)(v)
}

func (*loadBalancerSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerSkuResponse)(nil)).Elem()
}

func (i *loadBalancerSkuResponsePtrType) ToLoadBalancerSkuResponsePtrOutput() LoadBalancerSkuResponsePtrOutput {
	return i.ToLoadBalancerSkuResponsePtrOutputWithContext(context.Background())
}

func (i *loadBalancerSkuResponsePtrType) ToLoadBalancerSkuResponsePtrOutputWithContext(ctx context.Context) LoadBalancerSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerSkuResponsePtrOutput)
}

type LoadBalancerSkuResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSkuResponse)(nil)).Elem()
}

func (o LoadBalancerSkuResponseOutput) ToLoadBalancerSkuResponseOutput() LoadBalancerSkuResponseOutput {
	return o
}

func (o LoadBalancerSkuResponseOutput) ToLoadBalancerSkuResponseOutputWithContext(ctx context.Context) LoadBalancerSkuResponseOutput {
	return o
}

func (o LoadBalancerSkuResponseOutput) ToLoadBalancerSkuResponsePtrOutput() LoadBalancerSkuResponsePtrOutput {
	return o.ToLoadBalancerSkuResponsePtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuResponseOutput) ToLoadBalancerSkuResponsePtrOutputWithContext(ctx context.Context) LoadBalancerSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerSkuResponse) *LoadBalancerSkuResponse {
		return &v
	}).(LoadBalancerSkuResponsePtrOutput)
}

func (o LoadBalancerSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type LoadBalancerSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerSkuResponse)(nil)).Elem()
}

func (o LoadBalancerSkuResponsePtrOutput) ToLoadBalancerSkuResponsePtrOutput() LoadBalancerSkuResponsePtrOutput {
	return o
}

func (o LoadBalancerSkuResponsePtrOutput) ToLoadBalancerSkuResponsePtrOutputWithContext(ctx context.Context) LoadBalancerSkuResponsePtrOutput {
	return o
}

func (o LoadBalancerSkuResponsePtrOutput) Elem() LoadBalancerSkuResponseOutput {
	return o.ApplyT(func(v *LoadBalancerSkuResponse) LoadBalancerSkuResponse {
		if v != nil {
			return *v
		}
		var ret LoadBalancerSkuResponse
		return ret
	}).(LoadBalancerSkuResponseOutput)
}

func (o LoadBalancerSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type LoadBalancingRule struct {
	BackendAddressPool      *SubResource `pulumi:"backendAddressPool"`
	BackendPort             *int         `pulumi:"backendPort"`
	DisableOutboundSnat     *bool        `pulumi:"disableOutboundSnat"`
	EnableFloatingIP        *bool        `pulumi:"enableFloatingIP"`
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
	BackendAddressPool      SubResourcePtrInput   `pulumi:"backendAddressPool"`
	BackendPort             pulumi.IntPtrInput    `pulumi:"backendPort"`
	DisableOutboundSnat     pulumi.BoolPtrInput   `pulumi:"disableOutboundSnat"`
	EnableFloatingIP        pulumi.BoolPtrInput   `pulumi:"enableFloatingIP"`
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

func (o LoadBalancingRuleOutput) BackendAddressPool() SubResourcePtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *SubResource { return v.BackendAddressPool }).(SubResourcePtrOutput)
}

func (o LoadBalancingRuleOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingRuleOutput) DisableOutboundSnat() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *bool { return v.DisableOutboundSnat }).(pulumi.BoolPtrOutput)
}

func (o LoadBalancingRuleOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
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
	BackendAddressPool      *SubResourceResponse `pulumi:"backendAddressPool"`
	BackendPort             *int                 `pulumi:"backendPort"`
	DisableOutboundSnat     *bool                `pulumi:"disableOutboundSnat"`
	EnableFloatingIP        *bool                `pulumi:"enableFloatingIP"`
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
	BackendAddressPool      SubResourceResponsePtrInput `pulumi:"backendAddressPool"`
	BackendPort             pulumi.IntPtrInput          `pulumi:"backendPort"`
	DisableOutboundSnat     pulumi.BoolPtrInput         `pulumi:"disableOutboundSnat"`
	EnableFloatingIP        pulumi.BoolPtrInput         `pulumi:"enableFloatingIP"`
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

func (o LoadBalancingRuleResponseOutput) BackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponsePtrOutput)
}

func (o LoadBalancingRuleResponseOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) DisableOutboundSnat() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *bool { return v.DisableOutboundSnat }).(pulumi.BoolPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
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

type LocalNetworkGatewayType struct {
	BgpSettings              *BgpSettings      `pulumi:"bgpSettings"`
	Etag                     *string           `pulumi:"etag"`
	GatewayIpAddress         *string           `pulumi:"gatewayIpAddress"`
	Id                       *string           `pulumi:"id"`
	LocalNetworkAddressSpace *AddressSpace     `pulumi:"localNetworkAddressSpace"`
	Location                 *string           `pulumi:"location"`
	ResourceGuid             *string           `pulumi:"resourceGuid"`
	Tags                     map[string]string `pulumi:"tags"`
}





type LocalNetworkGatewayTypeInput interface {
	pulumi.Input

	ToLocalNetworkGatewayTypeOutput() LocalNetworkGatewayTypeOutput
	ToLocalNetworkGatewayTypeOutputWithContext(context.Context) LocalNetworkGatewayTypeOutput
}

type LocalNetworkGatewayTypeArgs struct {
	BgpSettings              BgpSettingsPtrInput   `pulumi:"bgpSettings"`
	Etag                     pulumi.StringPtrInput `pulumi:"etag"`
	GatewayIpAddress         pulumi.StringPtrInput `pulumi:"gatewayIpAddress"`
	Id                       pulumi.StringPtrInput `pulumi:"id"`
	LocalNetworkAddressSpace AddressSpacePtrInput  `pulumi:"localNetworkAddressSpace"`
	Location                 pulumi.StringPtrInput `pulumi:"location"`
	ResourceGuid             pulumi.StringPtrInput `pulumi:"resourceGuid"`
	Tags                     pulumi.StringMapInput `pulumi:"tags"`
}

func (LocalNetworkGatewayTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNetworkGatewayType)(nil)).Elem()
}

func (i LocalNetworkGatewayTypeArgs) ToLocalNetworkGatewayTypeOutput() LocalNetworkGatewayTypeOutput {
	return i.ToLocalNetworkGatewayTypeOutputWithContext(context.Background())
}

func (i LocalNetworkGatewayTypeArgs) ToLocalNetworkGatewayTypeOutputWithContext(ctx context.Context) LocalNetworkGatewayTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayTypeOutput)
}

func (i LocalNetworkGatewayTypeArgs) ToLocalNetworkGatewayTypePtrOutput() LocalNetworkGatewayTypePtrOutput {
	return i.ToLocalNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (i LocalNetworkGatewayTypeArgs) ToLocalNetworkGatewayTypePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayTypeOutput).ToLocalNetworkGatewayTypePtrOutputWithContext(ctx)
}









type LocalNetworkGatewayTypePtrInput interface {
	pulumi.Input

	ToLocalNetworkGatewayTypePtrOutput() LocalNetworkGatewayTypePtrOutput
	ToLocalNetworkGatewayTypePtrOutputWithContext(context.Context) LocalNetworkGatewayTypePtrOutput
}

type localNetworkGatewayTypePtrType LocalNetworkGatewayTypeArgs

func LocalNetworkGatewayTypePtr(v *LocalNetworkGatewayTypeArgs) LocalNetworkGatewayTypePtrInput {
	return (*localNetworkGatewayTypePtrType)(v)
}

func (*localNetworkGatewayTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGatewayType)(nil)).Elem()
}

func (i *localNetworkGatewayTypePtrType) ToLocalNetworkGatewayTypePtrOutput() LocalNetworkGatewayTypePtrOutput {
	return i.ToLocalNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (i *localNetworkGatewayTypePtrType) ToLocalNetworkGatewayTypePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayTypePtrOutput)
}

type LocalNetworkGatewayTypeOutput struct{ *pulumi.OutputState }

func (LocalNetworkGatewayTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNetworkGatewayType)(nil)).Elem()
}

func (o LocalNetworkGatewayTypeOutput) ToLocalNetworkGatewayTypeOutput() LocalNetworkGatewayTypeOutput {
	return o
}

func (o LocalNetworkGatewayTypeOutput) ToLocalNetworkGatewayTypeOutputWithContext(ctx context.Context) LocalNetworkGatewayTypeOutput {
	return o
}

func (o LocalNetworkGatewayTypeOutput) ToLocalNetworkGatewayTypePtrOutput() LocalNetworkGatewayTypePtrOutput {
	return o.ToLocalNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (o LocalNetworkGatewayTypeOutput) ToLocalNetworkGatewayTypePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalNetworkGatewayType) *LocalNetworkGatewayType {
		return &v
	}).(LocalNetworkGatewayTypePtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) BgpSettings() BgpSettingsPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *BgpSettings { return v.BgpSettings }).(BgpSettingsPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *string { return v.GatewayIpAddress }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) LocalNetworkAddressSpace() AddressSpacePtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *AddressSpace { return v.LocalNetworkAddressSpace }).(AddressSpacePtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type LocalNetworkGatewayTypePtrOutput struct{ *pulumi.OutputState }

func (LocalNetworkGatewayTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGatewayType)(nil)).Elem()
}

func (o LocalNetworkGatewayTypePtrOutput) ToLocalNetworkGatewayTypePtrOutput() LocalNetworkGatewayTypePtrOutput {
	return o
}

func (o LocalNetworkGatewayTypePtrOutput) ToLocalNetworkGatewayTypePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayTypePtrOutput {
	return o
}

func (o LocalNetworkGatewayTypePtrOutput) Elem() LocalNetworkGatewayTypeOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) LocalNetworkGatewayType {
		if v != nil {
			return *v
		}
		var ret LocalNetworkGatewayType
		return ret
	}).(LocalNetworkGatewayTypeOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) BgpSettings() BgpSettingsPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *BgpSettings {
		if v == nil {
			return nil
		}
		return v.BgpSettings
	}).(BgpSettingsPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.GatewayIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) LocalNetworkAddressSpace() AddressSpacePtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *AddressSpace {
		if v == nil {
			return nil
		}
		return v.LocalNetworkAddressSpace
	}).(AddressSpacePtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type LocalNetworkGatewayResponse struct {
	BgpSettings              *BgpSettingsResponse  `pulumi:"bgpSettings"`
	Etag                     *string               `pulumi:"etag"`
	GatewayIpAddress         *string               `pulumi:"gatewayIpAddress"`
	Id                       *string               `pulumi:"id"`
	LocalNetworkAddressSpace *AddressSpaceResponse `pulumi:"localNetworkAddressSpace"`
	Location                 *string               `pulumi:"location"`
	Name                     string                `pulumi:"name"`
	ProvisioningState        string                `pulumi:"provisioningState"`
	ResourceGuid             *string               `pulumi:"resourceGuid"`
	Tags                     map[string]string     `pulumi:"tags"`
	Type                     string                `pulumi:"type"`
}





type LocalNetworkGatewayResponseInput interface {
	pulumi.Input

	ToLocalNetworkGatewayResponseOutput() LocalNetworkGatewayResponseOutput
	ToLocalNetworkGatewayResponseOutputWithContext(context.Context) LocalNetworkGatewayResponseOutput
}

type LocalNetworkGatewayResponseArgs struct {
	BgpSettings              BgpSettingsResponsePtrInput  `pulumi:"bgpSettings"`
	Etag                     pulumi.StringPtrInput        `pulumi:"etag"`
	GatewayIpAddress         pulumi.StringPtrInput        `pulumi:"gatewayIpAddress"`
	Id                       pulumi.StringPtrInput        `pulumi:"id"`
	LocalNetworkAddressSpace AddressSpaceResponsePtrInput `pulumi:"localNetworkAddressSpace"`
	Location                 pulumi.StringPtrInput        `pulumi:"location"`
	Name                     pulumi.StringInput           `pulumi:"name"`
	ProvisioningState        pulumi.StringInput           `pulumi:"provisioningState"`
	ResourceGuid             pulumi.StringPtrInput        `pulumi:"resourceGuid"`
	Tags                     pulumi.StringMapInput        `pulumi:"tags"`
	Type                     pulumi.StringInput           `pulumi:"type"`
}

func (LocalNetworkGatewayResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNetworkGatewayResponse)(nil)).Elem()
}

func (i LocalNetworkGatewayResponseArgs) ToLocalNetworkGatewayResponseOutput() LocalNetworkGatewayResponseOutput {
	return i.ToLocalNetworkGatewayResponseOutputWithContext(context.Background())
}

func (i LocalNetworkGatewayResponseArgs) ToLocalNetworkGatewayResponseOutputWithContext(ctx context.Context) LocalNetworkGatewayResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayResponseOutput)
}

func (i LocalNetworkGatewayResponseArgs) ToLocalNetworkGatewayResponsePtrOutput() LocalNetworkGatewayResponsePtrOutput {
	return i.ToLocalNetworkGatewayResponsePtrOutputWithContext(context.Background())
}

func (i LocalNetworkGatewayResponseArgs) ToLocalNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayResponseOutput).ToLocalNetworkGatewayResponsePtrOutputWithContext(ctx)
}









type LocalNetworkGatewayResponsePtrInput interface {
	pulumi.Input

	ToLocalNetworkGatewayResponsePtrOutput() LocalNetworkGatewayResponsePtrOutput
	ToLocalNetworkGatewayResponsePtrOutputWithContext(context.Context) LocalNetworkGatewayResponsePtrOutput
}

type localNetworkGatewayResponsePtrType LocalNetworkGatewayResponseArgs

func LocalNetworkGatewayResponsePtr(v *LocalNetworkGatewayResponseArgs) LocalNetworkGatewayResponsePtrInput {
	return (*localNetworkGatewayResponsePtrType)(v)
}

func (*localNetworkGatewayResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGatewayResponse)(nil)).Elem()
}

func (i *localNetworkGatewayResponsePtrType) ToLocalNetworkGatewayResponsePtrOutput() LocalNetworkGatewayResponsePtrOutput {
	return i.ToLocalNetworkGatewayResponsePtrOutputWithContext(context.Background())
}

func (i *localNetworkGatewayResponsePtrType) ToLocalNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayResponsePtrOutput)
}

type LocalNetworkGatewayResponseOutput struct{ *pulumi.OutputState }

func (LocalNetworkGatewayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNetworkGatewayResponse)(nil)).Elem()
}

func (o LocalNetworkGatewayResponseOutput) ToLocalNetworkGatewayResponseOutput() LocalNetworkGatewayResponseOutput {
	return o
}

func (o LocalNetworkGatewayResponseOutput) ToLocalNetworkGatewayResponseOutputWithContext(ctx context.Context) LocalNetworkGatewayResponseOutput {
	return o
}

func (o LocalNetworkGatewayResponseOutput) ToLocalNetworkGatewayResponsePtrOutput() LocalNetworkGatewayResponsePtrOutput {
	return o.ToLocalNetworkGatewayResponsePtrOutputWithContext(context.Background())
}

func (o LocalNetworkGatewayResponseOutput) ToLocalNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalNetworkGatewayResponse) *LocalNetworkGatewayResponse {
		return &v
	}).(LocalNetworkGatewayResponsePtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *BgpSettingsResponse { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *string { return v.GatewayIpAddress }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) LocalNetworkAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *AddressSpaceResponse { return v.LocalNetworkAddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o LocalNetworkGatewayResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LocalNetworkGatewayResponseOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LocalNetworkGatewayResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) string { return v.Type }).(pulumi.StringOutput)
}

type LocalNetworkGatewayResponsePtrOutput struct{ *pulumi.OutputState }

func (LocalNetworkGatewayResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGatewayResponse)(nil)).Elem()
}

func (o LocalNetworkGatewayResponsePtrOutput) ToLocalNetworkGatewayResponsePtrOutput() LocalNetworkGatewayResponsePtrOutput {
	return o
}

func (o LocalNetworkGatewayResponsePtrOutput) ToLocalNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayResponsePtrOutput {
	return o
}

func (o LocalNetworkGatewayResponsePtrOutput) Elem() LocalNetworkGatewayResponseOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) LocalNetworkGatewayResponse {
		if v != nil {
			return *v
		}
		var ret LocalNetworkGatewayResponse
		return ret
	}).(LocalNetworkGatewayResponseOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *BgpSettingsResponse {
		if v == nil {
			return nil
		}
		return v.BgpSettings
	}).(BgpSettingsResponsePtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.GatewayIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) LocalNetworkAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *AddressSpaceResponse {
		if v == nil {
			return nil
		}
		return v.LocalNetworkAddressSpace
	}).(AddressSpaceResponsePtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceDnsSettings struct {
	AppliedDnsServers        []string `pulumi:"appliedDnsServers"`
	DnsServers               []string `pulumi:"dnsServers"`
	InternalDnsNameLabel     *string  `pulumi:"internalDnsNameLabel"`
	InternalDomainNameSuffix *string  `pulumi:"internalDomainNameSuffix"`
	InternalFqdn             *string  `pulumi:"internalFqdn"`
}





type NetworkInterfaceDnsSettingsInput interface {
	pulumi.Input

	ToNetworkInterfaceDnsSettingsOutput() NetworkInterfaceDnsSettingsOutput
	ToNetworkInterfaceDnsSettingsOutputWithContext(context.Context) NetworkInterfaceDnsSettingsOutput
}

type NetworkInterfaceDnsSettingsArgs struct {
	AppliedDnsServers        pulumi.StringArrayInput `pulumi:"appliedDnsServers"`
	DnsServers               pulumi.StringArrayInput `pulumi:"dnsServers"`
	InternalDnsNameLabel     pulumi.StringPtrInput   `pulumi:"internalDnsNameLabel"`
	InternalDomainNameSuffix pulumi.StringPtrInput   `pulumi:"internalDomainNameSuffix"`
	InternalFqdn             pulumi.StringPtrInput   `pulumi:"internalFqdn"`
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

func (o NetworkInterfaceDnsSettingsOutput) InternalDomainNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettings) *string { return v.InternalDomainNameSuffix }).(pulumi.StringPtrOutput)
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

func (o NetworkInterfaceDnsSettingsPtrOutput) InternalDomainNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.InternalDomainNameSuffix
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
	AppliedDnsServers        []string `pulumi:"appliedDnsServers"`
	DnsServers               []string `pulumi:"dnsServers"`
	InternalDnsNameLabel     *string  `pulumi:"internalDnsNameLabel"`
	InternalDomainNameSuffix *string  `pulumi:"internalDomainNameSuffix"`
	InternalFqdn             *string  `pulumi:"internalFqdn"`
}





type NetworkInterfaceDnsSettingsResponseInput interface {
	pulumi.Input

	ToNetworkInterfaceDnsSettingsResponseOutput() NetworkInterfaceDnsSettingsResponseOutput
	ToNetworkInterfaceDnsSettingsResponseOutputWithContext(context.Context) NetworkInterfaceDnsSettingsResponseOutput
}

type NetworkInterfaceDnsSettingsResponseArgs struct {
	AppliedDnsServers        pulumi.StringArrayInput `pulumi:"appliedDnsServers"`
	DnsServers               pulumi.StringArrayInput `pulumi:"dnsServers"`
	InternalDnsNameLabel     pulumi.StringPtrInput   `pulumi:"internalDnsNameLabel"`
	InternalDomainNameSuffix pulumi.StringPtrInput   `pulumi:"internalDomainNameSuffix"`
	InternalFqdn             pulumi.StringPtrInput   `pulumi:"internalFqdn"`
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

func (o NetworkInterfaceDnsSettingsResponseOutput) InternalDomainNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettingsResponse) *string { return v.InternalDomainNameSuffix }).(pulumi.StringPtrOutput)
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

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) InternalDomainNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InternalDomainNameSuffix
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

type NetworkInterfaceIPConfiguration struct {
	ApplicationGatewayBackendAddressPools []ApplicationGatewayBackendAddressPool `pulumi:"applicationGatewayBackendAddressPools"`
	ApplicationSecurityGroups             []ApplicationSecurityGroupType         `pulumi:"applicationSecurityGroups"`
	Etag                                  *string                                `pulumi:"etag"`
	Id                                    *string                                `pulumi:"id"`
	LoadBalancerBackendAddressPools       []BackendAddressPool                   `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules           []InboundNatRuleType                   `pulumi:"loadBalancerInboundNatRules"`
	Name                                  *string                                `pulumi:"name"`
	Primary                               *bool                                  `pulumi:"primary"`
	PrivateIPAddress                      *string                                `pulumi:"privateIPAddress"`
	PrivateIPAddressVersion               *string                                `pulumi:"privateIPAddressVersion"`
	PrivateIPAllocationMethod             *string                                `pulumi:"privateIPAllocationMethod"`
	ProvisioningState                     *string                                `pulumi:"provisioningState"`
	PublicIPAddress                       *PublicIPAddressType                   `pulumi:"publicIPAddress"`
	Subnet                                *SubnetType                            `pulumi:"subnet"`
}





type NetworkInterfaceIPConfigurationInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationOutput() NetworkInterfaceIPConfigurationOutput
	ToNetworkInterfaceIPConfigurationOutputWithContext(context.Context) NetworkInterfaceIPConfigurationOutput
}

type NetworkInterfaceIPConfigurationArgs struct {
	ApplicationGatewayBackendAddressPools ApplicationGatewayBackendAddressPoolArrayInput `pulumi:"applicationGatewayBackendAddressPools"`
	ApplicationSecurityGroups             ApplicationSecurityGroupTypeArrayInput         `pulumi:"applicationSecurityGroups"`
	Etag                                  pulumi.StringPtrInput                          `pulumi:"etag"`
	Id                                    pulumi.StringPtrInput                          `pulumi:"id"`
	LoadBalancerBackendAddressPools       BackendAddressPoolArrayInput                   `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules           InboundNatRuleTypeArrayInput                   `pulumi:"loadBalancerInboundNatRules"`
	Name                                  pulumi.StringPtrInput                          `pulumi:"name"`
	Primary                               pulumi.BoolPtrInput                            `pulumi:"primary"`
	PrivateIPAddress                      pulumi.StringPtrInput                          `pulumi:"privateIPAddress"`
	PrivateIPAddressVersion               pulumi.StringPtrInput                          `pulumi:"privateIPAddressVersion"`
	PrivateIPAllocationMethod             pulumi.StringPtrInput                          `pulumi:"privateIPAllocationMethod"`
	ProvisioningState                     pulumi.StringPtrInput                          `pulumi:"provisioningState"`
	PublicIPAddress                       PublicIPAddressTypePtrInput                    `pulumi:"publicIPAddress"`
	Subnet                                SubnetTypePtrInput                             `pulumi:"subnet"`
}

func (NetworkInterfaceIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (i NetworkInterfaceIPConfigurationArgs) ToNetworkInterfaceIPConfigurationOutput() NetworkInterfaceIPConfigurationOutput {
	return i.ToNetworkInterfaceIPConfigurationOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationArgs) ToNetworkInterfaceIPConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationOutput)
}





type NetworkInterfaceIPConfigurationArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationArrayOutput() NetworkInterfaceIPConfigurationArrayOutput
	ToNetworkInterfaceIPConfigurationArrayOutputWithContext(context.Context) NetworkInterfaceIPConfigurationArrayOutput
}

type NetworkInterfaceIPConfigurationArray []NetworkInterfaceIPConfigurationInput

func (NetworkInterfaceIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (i NetworkInterfaceIPConfigurationArray) ToNetworkInterfaceIPConfigurationArrayOutput() NetworkInterfaceIPConfigurationArrayOutput {
	return i.ToNetworkInterfaceIPConfigurationArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationArray) ToNetworkInterfaceIPConfigurationArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationArrayOutput)
}

type NetworkInterfaceIPConfigurationOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationOutput) ToNetworkInterfaceIPConfigurationOutput() NetworkInterfaceIPConfigurationOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationOutput) ToNetworkInterfaceIPConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationOutput) ApplicationGatewayBackendAddressPools() ApplicationGatewayBackendAddressPoolArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) []ApplicationGatewayBackendAddressPool {
		return v.ApplicationGatewayBackendAddressPools
	}).(ApplicationGatewayBackendAddressPoolArrayOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) ApplicationSecurityGroups() ApplicationSecurityGroupTypeArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) []ApplicationSecurityGroupType {
		return v.ApplicationSecurityGroups
	}).(ApplicationSecurityGroupTypeArrayOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) LoadBalancerBackendAddressPools() BackendAddressPoolArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) []BackendAddressPool { return v.LoadBalancerBackendAddressPools }).(BackendAddressPoolArrayOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) LoadBalancerInboundNatRules() InboundNatRuleTypeArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) []InboundNatRuleType { return v.LoadBalancerInboundNatRules }).(InboundNatRuleTypeArrayOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.PrivateIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) PublicIPAddress() PublicIPAddressTypePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *PublicIPAddressType { return v.PublicIPAddress }).(PublicIPAddressTypePtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Subnet() SubnetTypePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *SubnetType { return v.Subnet }).(SubnetTypePtrOutput)
}

type NetworkInterfaceIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationArrayOutput) ToNetworkInterfaceIPConfigurationArrayOutput() NetworkInterfaceIPConfigurationArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationArrayOutput) ToNetworkInterfaceIPConfigurationArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceIPConfiguration {
		return vs[0].([]NetworkInterfaceIPConfiguration)[vs[1].(int)]
	}).(NetworkInterfaceIPConfigurationOutput)
}

type NetworkInterfaceIPConfigurationResponse struct {
	ApplicationGatewayBackendAddressPools []ApplicationGatewayBackendAddressPoolResponse `pulumi:"applicationGatewayBackendAddressPools"`
	ApplicationSecurityGroups             []ApplicationSecurityGroupResponse             `pulumi:"applicationSecurityGroups"`
	Etag                                  *string                                        `pulumi:"etag"`
	Id                                    *string                                        `pulumi:"id"`
	LoadBalancerBackendAddressPools       []BackendAddressPoolResponse                   `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules           []InboundNatRuleResponse                       `pulumi:"loadBalancerInboundNatRules"`
	Name                                  *string                                        `pulumi:"name"`
	Primary                               *bool                                          `pulumi:"primary"`
	PrivateIPAddress                      *string                                        `pulumi:"privateIPAddress"`
	PrivateIPAddressVersion               *string                                        `pulumi:"privateIPAddressVersion"`
	PrivateIPAllocationMethod             *string                                        `pulumi:"privateIPAllocationMethod"`
	ProvisioningState                     *string                                        `pulumi:"provisioningState"`
	PublicIPAddress                       *PublicIPAddressResponse                       `pulumi:"publicIPAddress"`
	Subnet                                *SubnetResponse                                `pulumi:"subnet"`
}





type NetworkInterfaceIPConfigurationResponseInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationResponseOutput() NetworkInterfaceIPConfigurationResponseOutput
	ToNetworkInterfaceIPConfigurationResponseOutputWithContext(context.Context) NetworkInterfaceIPConfigurationResponseOutput
}

type NetworkInterfaceIPConfigurationResponseArgs struct {
	ApplicationGatewayBackendAddressPools ApplicationGatewayBackendAddressPoolResponseArrayInput `pulumi:"applicationGatewayBackendAddressPools"`
	ApplicationSecurityGroups             ApplicationSecurityGroupResponseArrayInput             `pulumi:"applicationSecurityGroups"`
	Etag                                  pulumi.StringPtrInput                                  `pulumi:"etag"`
	Id                                    pulumi.StringPtrInput                                  `pulumi:"id"`
	LoadBalancerBackendAddressPools       BackendAddressPoolResponseArrayInput                   `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules           InboundNatRuleResponseArrayInput                       `pulumi:"loadBalancerInboundNatRules"`
	Name                                  pulumi.StringPtrInput                                  `pulumi:"name"`
	Primary                               pulumi.BoolPtrInput                                    `pulumi:"primary"`
	PrivateIPAddress                      pulumi.StringPtrInput                                  `pulumi:"privateIPAddress"`
	PrivateIPAddressVersion               pulumi.StringPtrInput                                  `pulumi:"privateIPAddressVersion"`
	PrivateIPAllocationMethod             pulumi.StringPtrInput                                  `pulumi:"privateIPAllocationMethod"`
	ProvisioningState                     pulumi.StringPtrInput                                  `pulumi:"provisioningState"`
	PublicIPAddress                       PublicIPAddressResponsePtrInput                        `pulumi:"publicIPAddress"`
	Subnet                                SubnetResponsePtrInput                                 `pulumi:"subnet"`
}

func (NetworkInterfaceIPConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (i NetworkInterfaceIPConfigurationResponseArgs) ToNetworkInterfaceIPConfigurationResponseOutput() NetworkInterfaceIPConfigurationResponseOutput {
	return i.ToNetworkInterfaceIPConfigurationResponseOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationResponseArgs) ToNetworkInterfaceIPConfigurationResponseOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationResponseOutput)
}

func (i NetworkInterfaceIPConfigurationResponseArgs) ToNetworkInterfaceIPConfigurationResponsePtrOutput() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return i.ToNetworkInterfaceIPConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationResponseArgs) ToNetworkInterfaceIPConfigurationResponsePtrOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationResponseOutput).ToNetworkInterfaceIPConfigurationResponsePtrOutputWithContext(ctx)
}









type NetworkInterfaceIPConfigurationResponsePtrInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationResponsePtrOutput() NetworkInterfaceIPConfigurationResponsePtrOutput
	ToNetworkInterfaceIPConfigurationResponsePtrOutputWithContext(context.Context) NetworkInterfaceIPConfigurationResponsePtrOutput
}

type networkInterfaceIPConfigurationResponsePtrType NetworkInterfaceIPConfigurationResponseArgs

func NetworkInterfaceIPConfigurationResponsePtr(v *NetworkInterfaceIPConfigurationResponseArgs) NetworkInterfaceIPConfigurationResponsePtrInput {
	return (*networkInterfaceIPConfigurationResponsePtrType)(v)
}

func (*networkInterfaceIPConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (i *networkInterfaceIPConfigurationResponsePtrType) ToNetworkInterfaceIPConfigurationResponsePtrOutput() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return i.ToNetworkInterfaceIPConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *networkInterfaceIPConfigurationResponsePtrType) ToNetworkInterfaceIPConfigurationResponsePtrOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationResponsePtrOutput)
}





type NetworkInterfaceIPConfigurationResponseArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationResponseArrayOutput() NetworkInterfaceIPConfigurationResponseArrayOutput
	ToNetworkInterfaceIPConfigurationResponseArrayOutputWithContext(context.Context) NetworkInterfaceIPConfigurationResponseArrayOutput
}

type NetworkInterfaceIPConfigurationResponseArray []NetworkInterfaceIPConfigurationResponseInput

func (NetworkInterfaceIPConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (i NetworkInterfaceIPConfigurationResponseArray) ToNetworkInterfaceIPConfigurationResponseArrayOutput() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return i.ToNetworkInterfaceIPConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationResponseArray) ToNetworkInterfaceIPConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

type NetworkInterfaceIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ToNetworkInterfaceIPConfigurationResponseOutput() NetworkInterfaceIPConfigurationResponseOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ToNetworkInterfaceIPConfigurationResponseOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponseOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ToNetworkInterfaceIPConfigurationResponsePtrOutput() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o.ToNetworkInterfaceIPConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ToNetworkInterfaceIPConfigurationResponsePtrOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkInterfaceIPConfigurationResponse) *NetworkInterfaceIPConfigurationResponse {
		return &v
	}).(NetworkInterfaceIPConfigurationResponsePtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ApplicationGatewayBackendAddressPools() ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) []ApplicationGatewayBackendAddressPoolResponse {
		return v.ApplicationGatewayBackendAddressPools
	}).(ApplicationGatewayBackendAddressPoolResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) []ApplicationSecurityGroupResponse {
		return v.ApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) LoadBalancerBackendAddressPools() BackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) []BackendAddressPoolResponse {
		return v.LoadBalancerBackendAddressPools
	}).(BackendAddressPoolResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) LoadBalancerInboundNatRules() InboundNatRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) []InboundNatRuleResponse {
		return v.LoadBalancerInboundNatRules
	}).(InboundNatRuleResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.PrivateIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *PublicIPAddressResponse { return v.PublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

type NetworkInterfaceIPConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) ToNetworkInterfaceIPConfigurationResponsePtrOutput() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) ToNetworkInterfaceIPConfigurationResponsePtrOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) Elem() NetworkInterfaceIPConfigurationResponseOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) NetworkInterfaceIPConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceIPConfigurationResponse
		return ret
	}).(NetworkInterfaceIPConfigurationResponseOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) ApplicationGatewayBackendAddressPools() ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) []ApplicationGatewayBackendAddressPoolResponse {
		if v == nil {
			return nil
		}
		return v.ApplicationGatewayBackendAddressPools
	}).(ApplicationGatewayBackendAddressPoolResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) ApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) []ApplicationSecurityGroupResponse {
		if v == nil {
			return nil
		}
		return v.ApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) LoadBalancerBackendAddressPools() BackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) []BackendAddressPoolResponse {
		if v == nil {
			return nil
		}
		return v.LoadBalancerBackendAddressPools
	}).(BackendAddressPoolResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) LoadBalancerInboundNatRules() InboundNatRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) []InboundNatRuleResponse {
		if v == nil {
			return nil
		}
		return v.LoadBalancerInboundNatRules
	}).(InboundNatRuleResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Primary
	}).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAddress
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) *PublicIPAddressResponse {
		if v == nil {
			return nil
		}
		return v.PublicIPAddress
	}).(PublicIPAddressResponsePtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponsePtrOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceIPConfigurationResponse) *SubnetResponse {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(SubnetResponsePtrOutput)
}

type NetworkInterfaceIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationResponseArrayOutput) ToNetworkInterfaceIPConfigurationResponseArrayOutput() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseArrayOutput) ToNetworkInterfaceIPConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceIPConfigurationResponse {
		return vs[0].([]NetworkInterfaceIPConfigurationResponse)[vs[1].(int)]
	}).(NetworkInterfaceIPConfigurationResponseOutput)
}

type NetworkInterfaceResponse struct {
	DnsSettings                 *NetworkInterfaceDnsSettingsResponse      `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking *bool                                     `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          *bool                                     `pulumi:"enableIPForwarding"`
	Etag                        *string                                   `pulumi:"etag"`
	Id                          *string                                   `pulumi:"id"`
	IpConfigurations            []NetworkInterfaceIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location                    *string                                   `pulumi:"location"`
	MacAddress                  *string                                   `pulumi:"macAddress"`
	Name                        string                                    `pulumi:"name"`
	NetworkSecurityGroup        *NetworkSecurityGroupResponse             `pulumi:"networkSecurityGroup"`
	Primary                     *bool                                     `pulumi:"primary"`
	ProvisioningState           *string                                   `pulumi:"provisioningState"`
	ResourceGuid                *string                                   `pulumi:"resourceGuid"`
	Tags                        map[string]string                         `pulumi:"tags"`
	Type                        string                                    `pulumi:"type"`
	VirtualMachine              *SubResourceResponse                      `pulumi:"virtualMachine"`
}





type NetworkInterfaceResponseInput interface {
	pulumi.Input

	ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput
	ToNetworkInterfaceResponseOutputWithContext(context.Context) NetworkInterfaceResponseOutput
}

type NetworkInterfaceResponseArgs struct {
	DnsSettings                 NetworkInterfaceDnsSettingsResponsePtrInput       `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking pulumi.BoolPtrInput                               `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          pulumi.BoolPtrInput                               `pulumi:"enableIPForwarding"`
	Etag                        pulumi.StringPtrInput                             `pulumi:"etag"`
	Id                          pulumi.StringPtrInput                             `pulumi:"id"`
	IpConfigurations            NetworkInterfaceIPConfigurationResponseArrayInput `pulumi:"ipConfigurations"`
	Location                    pulumi.StringPtrInput                             `pulumi:"location"`
	MacAddress                  pulumi.StringPtrInput                             `pulumi:"macAddress"`
	Name                        pulumi.StringInput                                `pulumi:"name"`
	NetworkSecurityGroup        NetworkSecurityGroupResponsePtrInput              `pulumi:"networkSecurityGroup"`
	Primary                     pulumi.BoolPtrInput                               `pulumi:"primary"`
	ProvisioningState           pulumi.StringPtrInput                             `pulumi:"provisioningState"`
	ResourceGuid                pulumi.StringPtrInput                             `pulumi:"resourceGuid"`
	Tags                        pulumi.StringMapInput                             `pulumi:"tags"`
	Type                        pulumi.StringInput                                `pulumi:"type"`
	VirtualMachine              SubResourceResponsePtrInput                       `pulumi:"virtualMachine"`
}

func (NetworkInterfaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResponse)(nil)).Elem()
}

func (i NetworkInterfaceResponseArgs) ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput {
	return i.ToNetworkInterfaceResponseOutputWithContext(context.Background())
}

func (i NetworkInterfaceResponseArgs) ToNetworkInterfaceResponseOutputWithContext(ctx context.Context) NetworkInterfaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceResponseOutput)
}





type NetworkInterfaceResponseArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput
	ToNetworkInterfaceResponseArrayOutputWithContext(context.Context) NetworkInterfaceResponseArrayOutput
}

type NetworkInterfaceResponseArray []NetworkInterfaceResponseInput

func (NetworkInterfaceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceResponse)(nil)).Elem()
}

func (i NetworkInterfaceResponseArray) ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput {
	return i.ToNetworkInterfaceResponseArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceResponseArray) ToNetworkInterfaceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceResponseArrayOutput)
}

type NetworkInterfaceResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutputWithContext(ctx context.Context) NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) DnsSettings() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *NetworkInterfaceDnsSettingsResponse { return v.DnsSettings }).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

func (o NetworkInterfaceResponseOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceResponseOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) IpConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) []NetworkInterfaceIPConfigurationResponse { return v.IpConfigurations }).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o NetworkInterfaceResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *NetworkSecurityGroupResponse { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o NetworkInterfaceResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkInterfaceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) VirtualMachine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *SubResourceResponse { return v.VirtualMachine }).(SubResourceResponsePtrOutput)
}

type NetworkInterfaceResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseArrayOutput) ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceResponseArrayOutput) ToNetworkInterfaceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceResponse {
		return vs[0].([]NetworkInterfaceResponse)[vs[1].(int)]
	}).(NetworkInterfaceResponseOutput)
}

type NetworkSecurityGroupType struct {
	DefaultSecurityRules []SecurityRuleType `pulumi:"defaultSecurityRules"`
	Etag                 *string            `pulumi:"etag"`
	Id                   *string            `pulumi:"id"`
	Location             *string            `pulumi:"location"`
	ProvisioningState    *string            `pulumi:"provisioningState"`
	ResourceGuid         *string            `pulumi:"resourceGuid"`
	SecurityRules        []SecurityRuleType `pulumi:"securityRules"`
	Tags                 map[string]string  `pulumi:"tags"`
}





type NetworkSecurityGroupTypeInput interface {
	pulumi.Input

	ToNetworkSecurityGroupTypeOutput() NetworkSecurityGroupTypeOutput
	ToNetworkSecurityGroupTypeOutputWithContext(context.Context) NetworkSecurityGroupTypeOutput
}

type NetworkSecurityGroupTypeArgs struct {
	DefaultSecurityRules SecurityRuleTypeArrayInput `pulumi:"defaultSecurityRules"`
	Etag                 pulumi.StringPtrInput      `pulumi:"etag"`
	Id                   pulumi.StringPtrInput      `pulumi:"id"`
	Location             pulumi.StringPtrInput      `pulumi:"location"`
	ProvisioningState    pulumi.StringPtrInput      `pulumi:"provisioningState"`
	ResourceGuid         pulumi.StringPtrInput      `pulumi:"resourceGuid"`
	SecurityRules        SecurityRuleTypeArrayInput `pulumi:"securityRules"`
	Tags                 pulumi.StringMapInput      `pulumi:"tags"`
}

func (NetworkSecurityGroupTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupType)(nil)).Elem()
}

func (i NetworkSecurityGroupTypeArgs) ToNetworkSecurityGroupTypeOutput() NetworkSecurityGroupTypeOutput {
	return i.ToNetworkSecurityGroupTypeOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupTypeArgs) ToNetworkSecurityGroupTypeOutputWithContext(ctx context.Context) NetworkSecurityGroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupTypeOutput)
}

func (i NetworkSecurityGroupTypeArgs) ToNetworkSecurityGroupTypePtrOutput() NetworkSecurityGroupTypePtrOutput {
	return i.ToNetworkSecurityGroupTypePtrOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupTypeArgs) ToNetworkSecurityGroupTypePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupTypeOutput).ToNetworkSecurityGroupTypePtrOutputWithContext(ctx)
}









type NetworkSecurityGroupTypePtrInput interface {
	pulumi.Input

	ToNetworkSecurityGroupTypePtrOutput() NetworkSecurityGroupTypePtrOutput
	ToNetworkSecurityGroupTypePtrOutputWithContext(context.Context) NetworkSecurityGroupTypePtrOutput
}

type networkSecurityGroupTypePtrType NetworkSecurityGroupTypeArgs

func NetworkSecurityGroupTypePtr(v *NetworkSecurityGroupTypeArgs) NetworkSecurityGroupTypePtrInput {
	return (*networkSecurityGroupTypePtrType)(v)
}

func (*networkSecurityGroupTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroupType)(nil)).Elem()
}

func (i *networkSecurityGroupTypePtrType) ToNetworkSecurityGroupTypePtrOutput() NetworkSecurityGroupTypePtrOutput {
	return i.ToNetworkSecurityGroupTypePtrOutputWithContext(context.Background())
}

func (i *networkSecurityGroupTypePtrType) ToNetworkSecurityGroupTypePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupTypePtrOutput)
}

type NetworkSecurityGroupTypeOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupType)(nil)).Elem()
}

func (o NetworkSecurityGroupTypeOutput) ToNetworkSecurityGroupTypeOutput() NetworkSecurityGroupTypeOutput {
	return o
}

func (o NetworkSecurityGroupTypeOutput) ToNetworkSecurityGroupTypeOutputWithContext(ctx context.Context) NetworkSecurityGroupTypeOutput {
	return o
}

func (o NetworkSecurityGroupTypeOutput) ToNetworkSecurityGroupTypePtrOutput() NetworkSecurityGroupTypePtrOutput {
	return o.ToNetworkSecurityGroupTypePtrOutputWithContext(context.Background())
}

func (o NetworkSecurityGroupTypeOutput) ToNetworkSecurityGroupTypePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkSecurityGroupType) *NetworkSecurityGroupType {
		return &v
	}).(NetworkSecurityGroupTypePtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) DefaultSecurityRules() SecurityRuleTypeArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) []SecurityRuleType { return v.DefaultSecurityRules }).(SecurityRuleTypeArrayOutput)
}

func (o NetworkSecurityGroupTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) SecurityRules() SecurityRuleTypeArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) []SecurityRuleType { return v.SecurityRules }).(SecurityRuleTypeArrayOutput)
}

func (o NetworkSecurityGroupTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type NetworkSecurityGroupTypePtrOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroupType)(nil)).Elem()
}

func (o NetworkSecurityGroupTypePtrOutput) ToNetworkSecurityGroupTypePtrOutput() NetworkSecurityGroupTypePtrOutput {
	return o
}

func (o NetworkSecurityGroupTypePtrOutput) ToNetworkSecurityGroupTypePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupTypePtrOutput {
	return o
}

func (o NetworkSecurityGroupTypePtrOutput) Elem() NetworkSecurityGroupTypeOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) NetworkSecurityGroupType {
		if v != nil {
			return *v
		}
		var ret NetworkSecurityGroupType
		return ret
	}).(NetworkSecurityGroupTypeOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) DefaultSecurityRules() SecurityRuleTypeArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) []SecurityRuleType {
		if v == nil {
			return nil
		}
		return v.DefaultSecurityRules
	}).(SecurityRuleTypeArrayOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) SecurityRules() SecurityRuleTypeArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) []SecurityRuleType {
		if v == nil {
			return nil
		}
		return v.SecurityRules
	}).(SecurityRuleTypeArrayOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type NetworkSecurityGroupResponse struct {
	DefaultSecurityRules []SecurityRuleResponse     `pulumi:"defaultSecurityRules"`
	Etag                 *string                    `pulumi:"etag"`
	Id                   *string                    `pulumi:"id"`
	Location             *string                    `pulumi:"location"`
	Name                 string                     `pulumi:"name"`
	NetworkInterfaces    []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	ProvisioningState    *string                    `pulumi:"provisioningState"`
	ResourceGuid         *string                    `pulumi:"resourceGuid"`
	SecurityRules        []SecurityRuleResponse     `pulumi:"securityRules"`
	Subnets              []SubnetResponse           `pulumi:"subnets"`
	Tags                 map[string]string          `pulumi:"tags"`
	Type                 string                     `pulumi:"type"`
}





type NetworkSecurityGroupResponseInput interface {
	pulumi.Input

	ToNetworkSecurityGroupResponseOutput() NetworkSecurityGroupResponseOutput
	ToNetworkSecurityGroupResponseOutputWithContext(context.Context) NetworkSecurityGroupResponseOutput
}

type NetworkSecurityGroupResponseArgs struct {
	DefaultSecurityRules SecurityRuleResponseArrayInput     `pulumi:"defaultSecurityRules"`
	Etag                 pulumi.StringPtrInput              `pulumi:"etag"`
	Id                   pulumi.StringPtrInput              `pulumi:"id"`
	Location             pulumi.StringPtrInput              `pulumi:"location"`
	Name                 pulumi.StringInput                 `pulumi:"name"`
	NetworkInterfaces    NetworkInterfaceResponseArrayInput `pulumi:"networkInterfaces"`
	ProvisioningState    pulumi.StringPtrInput              `pulumi:"provisioningState"`
	ResourceGuid         pulumi.StringPtrInput              `pulumi:"resourceGuid"`
	SecurityRules        SecurityRuleResponseArrayInput     `pulumi:"securityRules"`
	Subnets              SubnetResponseArrayInput           `pulumi:"subnets"`
	Tags                 pulumi.StringMapInput              `pulumi:"tags"`
	Type                 pulumi.StringInput                 `pulumi:"type"`
}

func (NetworkSecurityGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResponse)(nil)).Elem()
}

func (i NetworkSecurityGroupResponseArgs) ToNetworkSecurityGroupResponseOutput() NetworkSecurityGroupResponseOutput {
	return i.ToNetworkSecurityGroupResponseOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupResponseArgs) ToNetworkSecurityGroupResponseOutputWithContext(ctx context.Context) NetworkSecurityGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupResponseOutput)
}

func (i NetworkSecurityGroupResponseArgs) ToNetworkSecurityGroupResponsePtrOutput() NetworkSecurityGroupResponsePtrOutput {
	return i.ToNetworkSecurityGroupResponsePtrOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupResponseArgs) ToNetworkSecurityGroupResponsePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupResponseOutput).ToNetworkSecurityGroupResponsePtrOutputWithContext(ctx)
}









type NetworkSecurityGroupResponsePtrInput interface {
	pulumi.Input

	ToNetworkSecurityGroupResponsePtrOutput() NetworkSecurityGroupResponsePtrOutput
	ToNetworkSecurityGroupResponsePtrOutputWithContext(context.Context) NetworkSecurityGroupResponsePtrOutput
}

type networkSecurityGroupResponsePtrType NetworkSecurityGroupResponseArgs

func NetworkSecurityGroupResponsePtr(v *NetworkSecurityGroupResponseArgs) NetworkSecurityGroupResponsePtrInput {
	return (*networkSecurityGroupResponsePtrType)(v)
}

func (*networkSecurityGroupResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroupResponse)(nil)).Elem()
}

func (i *networkSecurityGroupResponsePtrType) ToNetworkSecurityGroupResponsePtrOutput() NetworkSecurityGroupResponsePtrOutput {
	return i.ToNetworkSecurityGroupResponsePtrOutputWithContext(context.Background())
}

func (i *networkSecurityGroupResponsePtrType) ToNetworkSecurityGroupResponsePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupResponsePtrOutput)
}

type NetworkSecurityGroupResponseOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResponse)(nil)).Elem()
}

func (o NetworkSecurityGroupResponseOutput) ToNetworkSecurityGroupResponseOutput() NetworkSecurityGroupResponseOutput {
	return o
}

func (o NetworkSecurityGroupResponseOutput) ToNetworkSecurityGroupResponseOutputWithContext(ctx context.Context) NetworkSecurityGroupResponseOutput {
	return o
}

func (o NetworkSecurityGroupResponseOutput) ToNetworkSecurityGroupResponsePtrOutput() NetworkSecurityGroupResponsePtrOutput {
	return o.ToNetworkSecurityGroupResponsePtrOutputWithContext(context.Background())
}

func (o NetworkSecurityGroupResponseOutput) ToNetworkSecurityGroupResponsePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkSecurityGroupResponse) *NetworkSecurityGroupResponse {
		return &v
	}).(NetworkSecurityGroupResponsePtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) DefaultSecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) []SecurityRuleResponse { return v.DefaultSecurityRules }).(SecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkSecurityGroupResponseOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o NetworkSecurityGroupResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) SecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) []SecurityRuleResponse { return v.SecurityRules }).(SecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupResponseOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

func (o NetworkSecurityGroupResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityGroupResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) string { return v.Type }).(pulumi.StringOutput)
}

type NetworkSecurityGroupResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroupResponse)(nil)).Elem()
}

func (o NetworkSecurityGroupResponsePtrOutput) ToNetworkSecurityGroupResponsePtrOutput() NetworkSecurityGroupResponsePtrOutput {
	return o
}

func (o NetworkSecurityGroupResponsePtrOutput) ToNetworkSecurityGroupResponsePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupResponsePtrOutput {
	return o
}

func (o NetworkSecurityGroupResponsePtrOutput) Elem() NetworkSecurityGroupResponseOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) NetworkSecurityGroupResponse {
		if v != nil {
			return *v
		}
		var ret NetworkSecurityGroupResponse
		return ret
	}).(NetworkSecurityGroupResponseOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) DefaultSecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) []SecurityRuleResponse {
		if v == nil {
			return nil
		}
		return v.DefaultSecurityRules
	}).(SecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) []NetworkInterfaceResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceResponseArrayOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) SecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) []SecurityRuleResponse {
		if v == nil {
			return nil
		}
		return v.SecurityRules
	}).(SecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) []SubnetResponse {
		if v == nil {
			return nil
		}
		return v.Subnets
	}).(SubnetResponseArrayOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type OutboundNatRule struct {
	AllocatedOutboundPorts   *int          `pulumi:"allocatedOutboundPorts"`
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
	AllocatedOutboundPorts   pulumi.IntPtrInput    `pulumi:"allocatedOutboundPorts"`
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

func (o OutboundNatRuleOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OutboundNatRule) *int { return v.AllocatedOutboundPorts }).(pulumi.IntPtrOutput)
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
	AllocatedOutboundPorts   *int                  `pulumi:"allocatedOutboundPorts"`
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
	AllocatedOutboundPorts   pulumi.IntPtrInput            `pulumi:"allocatedOutboundPorts"`
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

func (o OutboundNatRuleResponseOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) *int { return v.AllocatedOutboundPorts }).(pulumi.IntPtrOutput)
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

type PacketCaptureFilter struct {
	LocalIPAddress  *string `pulumi:"localIPAddress"`
	LocalPort       *string `pulumi:"localPort"`
	Protocol        *string `pulumi:"protocol"`
	RemoteIPAddress *string `pulumi:"remoteIPAddress"`
	RemotePort      *string `pulumi:"remotePort"`
}


func (val *PacketCaptureFilter) Defaults() *PacketCaptureFilter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Protocol) {
		protocol_ := "Any"
		tmp.Protocol = &protocol_
	}
	return &tmp
}





type PacketCaptureFilterInput interface {
	pulumi.Input

	ToPacketCaptureFilterOutput() PacketCaptureFilterOutput
	ToPacketCaptureFilterOutputWithContext(context.Context) PacketCaptureFilterOutput
}

type PacketCaptureFilterArgs struct {
	LocalIPAddress  pulumi.StringPtrInput `pulumi:"localIPAddress"`
	LocalPort       pulumi.StringPtrInput `pulumi:"localPort"`
	Protocol        pulumi.StringPtrInput `pulumi:"protocol"`
	RemoteIPAddress pulumi.StringPtrInput `pulumi:"remoteIPAddress"`
	RemotePort      pulumi.StringPtrInput `pulumi:"remotePort"`
}

func (PacketCaptureFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureFilter)(nil)).Elem()
}

func (i PacketCaptureFilterArgs) ToPacketCaptureFilterOutput() PacketCaptureFilterOutput {
	return i.ToPacketCaptureFilterOutputWithContext(context.Background())
}

func (i PacketCaptureFilterArgs) ToPacketCaptureFilterOutputWithContext(ctx context.Context) PacketCaptureFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureFilterOutput)
}





type PacketCaptureFilterArrayInput interface {
	pulumi.Input

	ToPacketCaptureFilterArrayOutput() PacketCaptureFilterArrayOutput
	ToPacketCaptureFilterArrayOutputWithContext(context.Context) PacketCaptureFilterArrayOutput
}

type PacketCaptureFilterArray []PacketCaptureFilterInput

func (PacketCaptureFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PacketCaptureFilter)(nil)).Elem()
}

func (i PacketCaptureFilterArray) ToPacketCaptureFilterArrayOutput() PacketCaptureFilterArrayOutput {
	return i.ToPacketCaptureFilterArrayOutputWithContext(context.Background())
}

func (i PacketCaptureFilterArray) ToPacketCaptureFilterArrayOutputWithContext(ctx context.Context) PacketCaptureFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureFilterArrayOutput)
}

type PacketCaptureFilterOutput struct{ *pulumi.OutputState }

func (PacketCaptureFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureFilter)(nil)).Elem()
}

func (o PacketCaptureFilterOutput) ToPacketCaptureFilterOutput() PacketCaptureFilterOutput {
	return o
}

func (o PacketCaptureFilterOutput) ToPacketCaptureFilterOutputWithContext(ctx context.Context) PacketCaptureFilterOutput {
	return o
}

func (o PacketCaptureFilterOutput) LocalIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilter) *string { return v.LocalIPAddress }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterOutput) LocalPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilter) *string { return v.LocalPort }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilter) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterOutput) RemoteIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilter) *string { return v.RemoteIPAddress }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterOutput) RemotePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilter) *string { return v.RemotePort }).(pulumi.StringPtrOutput)
}

type PacketCaptureFilterArrayOutput struct{ *pulumi.OutputState }

func (PacketCaptureFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PacketCaptureFilter)(nil)).Elem()
}

func (o PacketCaptureFilterArrayOutput) ToPacketCaptureFilterArrayOutput() PacketCaptureFilterArrayOutput {
	return o
}

func (o PacketCaptureFilterArrayOutput) ToPacketCaptureFilterArrayOutputWithContext(ctx context.Context) PacketCaptureFilterArrayOutput {
	return o
}

func (o PacketCaptureFilterArrayOutput) Index(i pulumi.IntInput) PacketCaptureFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PacketCaptureFilter {
		return vs[0].([]PacketCaptureFilter)[vs[1].(int)]
	}).(PacketCaptureFilterOutput)
}

type PacketCaptureFilterResponse struct {
	LocalIPAddress  *string `pulumi:"localIPAddress"`
	LocalPort       *string `pulumi:"localPort"`
	Protocol        *string `pulumi:"protocol"`
	RemoteIPAddress *string `pulumi:"remoteIPAddress"`
	RemotePort      *string `pulumi:"remotePort"`
}


func (val *PacketCaptureFilterResponse) Defaults() *PacketCaptureFilterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Protocol) {
		protocol_ := "Any"
		tmp.Protocol = &protocol_
	}
	return &tmp
}





type PacketCaptureFilterResponseInput interface {
	pulumi.Input

	ToPacketCaptureFilterResponseOutput() PacketCaptureFilterResponseOutput
	ToPacketCaptureFilterResponseOutputWithContext(context.Context) PacketCaptureFilterResponseOutput
}

type PacketCaptureFilterResponseArgs struct {
	LocalIPAddress  pulumi.StringPtrInput `pulumi:"localIPAddress"`
	LocalPort       pulumi.StringPtrInput `pulumi:"localPort"`
	Protocol        pulumi.StringPtrInput `pulumi:"protocol"`
	RemoteIPAddress pulumi.StringPtrInput `pulumi:"remoteIPAddress"`
	RemotePort      pulumi.StringPtrInput `pulumi:"remotePort"`
}

func (PacketCaptureFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureFilterResponse)(nil)).Elem()
}

func (i PacketCaptureFilterResponseArgs) ToPacketCaptureFilterResponseOutput() PacketCaptureFilterResponseOutput {
	return i.ToPacketCaptureFilterResponseOutputWithContext(context.Background())
}

func (i PacketCaptureFilterResponseArgs) ToPacketCaptureFilterResponseOutputWithContext(ctx context.Context) PacketCaptureFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureFilterResponseOutput)
}





type PacketCaptureFilterResponseArrayInput interface {
	pulumi.Input

	ToPacketCaptureFilterResponseArrayOutput() PacketCaptureFilterResponseArrayOutput
	ToPacketCaptureFilterResponseArrayOutputWithContext(context.Context) PacketCaptureFilterResponseArrayOutput
}

type PacketCaptureFilterResponseArray []PacketCaptureFilterResponseInput

func (PacketCaptureFilterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PacketCaptureFilterResponse)(nil)).Elem()
}

func (i PacketCaptureFilterResponseArray) ToPacketCaptureFilterResponseArrayOutput() PacketCaptureFilterResponseArrayOutput {
	return i.ToPacketCaptureFilterResponseArrayOutputWithContext(context.Background())
}

func (i PacketCaptureFilterResponseArray) ToPacketCaptureFilterResponseArrayOutputWithContext(ctx context.Context) PacketCaptureFilterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureFilterResponseArrayOutput)
}

type PacketCaptureFilterResponseOutput struct{ *pulumi.OutputState }

func (PacketCaptureFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureFilterResponse)(nil)).Elem()
}

func (o PacketCaptureFilterResponseOutput) ToPacketCaptureFilterResponseOutput() PacketCaptureFilterResponseOutput {
	return o
}

func (o PacketCaptureFilterResponseOutput) ToPacketCaptureFilterResponseOutputWithContext(ctx context.Context) PacketCaptureFilterResponseOutput {
	return o
}

func (o PacketCaptureFilterResponseOutput) LocalIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilterResponse) *string { return v.LocalIPAddress }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterResponseOutput) LocalPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilterResponse) *string { return v.LocalPort }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilterResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterResponseOutput) RemoteIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilterResponse) *string { return v.RemoteIPAddress }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterResponseOutput) RemotePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilterResponse) *string { return v.RemotePort }).(pulumi.StringPtrOutput)
}

type PacketCaptureFilterResponseArrayOutput struct{ *pulumi.OutputState }

func (PacketCaptureFilterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PacketCaptureFilterResponse)(nil)).Elem()
}

func (o PacketCaptureFilterResponseArrayOutput) ToPacketCaptureFilterResponseArrayOutput() PacketCaptureFilterResponseArrayOutput {
	return o
}

func (o PacketCaptureFilterResponseArrayOutput) ToPacketCaptureFilterResponseArrayOutputWithContext(ctx context.Context) PacketCaptureFilterResponseArrayOutput {
	return o
}

func (o PacketCaptureFilterResponseArrayOutput) Index(i pulumi.IntInput) PacketCaptureFilterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PacketCaptureFilterResponse {
		return vs[0].([]PacketCaptureFilterResponse)[vs[1].(int)]
	}).(PacketCaptureFilterResponseOutput)
}

type PacketCaptureStorageLocation struct {
	FilePath    *string `pulumi:"filePath"`
	StorageId   *string `pulumi:"storageId"`
	StoragePath *string `pulumi:"storagePath"`
}





type PacketCaptureStorageLocationInput interface {
	pulumi.Input

	ToPacketCaptureStorageLocationOutput() PacketCaptureStorageLocationOutput
	ToPacketCaptureStorageLocationOutputWithContext(context.Context) PacketCaptureStorageLocationOutput
}

type PacketCaptureStorageLocationArgs struct {
	FilePath    pulumi.StringPtrInput `pulumi:"filePath"`
	StorageId   pulumi.StringPtrInput `pulumi:"storageId"`
	StoragePath pulumi.StringPtrInput `pulumi:"storagePath"`
}

func (PacketCaptureStorageLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureStorageLocation)(nil)).Elem()
}

func (i PacketCaptureStorageLocationArgs) ToPacketCaptureStorageLocationOutput() PacketCaptureStorageLocationOutput {
	return i.ToPacketCaptureStorageLocationOutputWithContext(context.Background())
}

func (i PacketCaptureStorageLocationArgs) ToPacketCaptureStorageLocationOutputWithContext(ctx context.Context) PacketCaptureStorageLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureStorageLocationOutput)
}

func (i PacketCaptureStorageLocationArgs) ToPacketCaptureStorageLocationPtrOutput() PacketCaptureStorageLocationPtrOutput {
	return i.ToPacketCaptureStorageLocationPtrOutputWithContext(context.Background())
}

func (i PacketCaptureStorageLocationArgs) ToPacketCaptureStorageLocationPtrOutputWithContext(ctx context.Context) PacketCaptureStorageLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureStorageLocationOutput).ToPacketCaptureStorageLocationPtrOutputWithContext(ctx)
}









type PacketCaptureStorageLocationPtrInput interface {
	pulumi.Input

	ToPacketCaptureStorageLocationPtrOutput() PacketCaptureStorageLocationPtrOutput
	ToPacketCaptureStorageLocationPtrOutputWithContext(context.Context) PacketCaptureStorageLocationPtrOutput
}

type packetCaptureStorageLocationPtrType PacketCaptureStorageLocationArgs

func PacketCaptureStorageLocationPtr(v *PacketCaptureStorageLocationArgs) PacketCaptureStorageLocationPtrInput {
	return (*packetCaptureStorageLocationPtrType)(v)
}

func (*packetCaptureStorageLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCaptureStorageLocation)(nil)).Elem()
}

func (i *packetCaptureStorageLocationPtrType) ToPacketCaptureStorageLocationPtrOutput() PacketCaptureStorageLocationPtrOutput {
	return i.ToPacketCaptureStorageLocationPtrOutputWithContext(context.Background())
}

func (i *packetCaptureStorageLocationPtrType) ToPacketCaptureStorageLocationPtrOutputWithContext(ctx context.Context) PacketCaptureStorageLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureStorageLocationPtrOutput)
}

type PacketCaptureStorageLocationOutput struct{ *pulumi.OutputState }

func (PacketCaptureStorageLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureStorageLocation)(nil)).Elem()
}

func (o PacketCaptureStorageLocationOutput) ToPacketCaptureStorageLocationOutput() PacketCaptureStorageLocationOutput {
	return o
}

func (o PacketCaptureStorageLocationOutput) ToPacketCaptureStorageLocationOutputWithContext(ctx context.Context) PacketCaptureStorageLocationOutput {
	return o
}

func (o PacketCaptureStorageLocationOutput) ToPacketCaptureStorageLocationPtrOutput() PacketCaptureStorageLocationPtrOutput {
	return o.ToPacketCaptureStorageLocationPtrOutputWithContext(context.Background())
}

func (o PacketCaptureStorageLocationOutput) ToPacketCaptureStorageLocationPtrOutputWithContext(ctx context.Context) PacketCaptureStorageLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PacketCaptureStorageLocation) *PacketCaptureStorageLocation {
		return &v
	}).(PacketCaptureStorageLocationPtrOutput)
}

func (o PacketCaptureStorageLocationOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocation) *string { return v.FilePath }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationOutput) StorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocation) *string { return v.StorageId }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationOutput) StoragePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocation) *string { return v.StoragePath }).(pulumi.StringPtrOutput)
}

type PacketCaptureStorageLocationPtrOutput struct{ *pulumi.OutputState }

func (PacketCaptureStorageLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCaptureStorageLocation)(nil)).Elem()
}

func (o PacketCaptureStorageLocationPtrOutput) ToPacketCaptureStorageLocationPtrOutput() PacketCaptureStorageLocationPtrOutput {
	return o
}

func (o PacketCaptureStorageLocationPtrOutput) ToPacketCaptureStorageLocationPtrOutputWithContext(ctx context.Context) PacketCaptureStorageLocationPtrOutput {
	return o
}

func (o PacketCaptureStorageLocationPtrOutput) Elem() PacketCaptureStorageLocationOutput {
	return o.ApplyT(func(v *PacketCaptureStorageLocation) PacketCaptureStorageLocation {
		if v != nil {
			return *v
		}
		var ret PacketCaptureStorageLocation
		return ret
	}).(PacketCaptureStorageLocationOutput)
}

func (o PacketCaptureStorageLocationPtrOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCaptureStorageLocation) *string {
		if v == nil {
			return nil
		}
		return v.FilePath
	}).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationPtrOutput) StorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCaptureStorageLocation) *string {
		if v == nil {
			return nil
		}
		return v.StorageId
	}).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationPtrOutput) StoragePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCaptureStorageLocation) *string {
		if v == nil {
			return nil
		}
		return v.StoragePath
	}).(pulumi.StringPtrOutput)
}

type PacketCaptureStorageLocationResponse struct {
	FilePath    *string `pulumi:"filePath"`
	StorageId   *string `pulumi:"storageId"`
	StoragePath *string `pulumi:"storagePath"`
}





type PacketCaptureStorageLocationResponseInput interface {
	pulumi.Input

	ToPacketCaptureStorageLocationResponseOutput() PacketCaptureStorageLocationResponseOutput
	ToPacketCaptureStorageLocationResponseOutputWithContext(context.Context) PacketCaptureStorageLocationResponseOutput
}

type PacketCaptureStorageLocationResponseArgs struct {
	FilePath    pulumi.StringPtrInput `pulumi:"filePath"`
	StorageId   pulumi.StringPtrInput `pulumi:"storageId"`
	StoragePath pulumi.StringPtrInput `pulumi:"storagePath"`
}

func (PacketCaptureStorageLocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureStorageLocationResponse)(nil)).Elem()
}

func (i PacketCaptureStorageLocationResponseArgs) ToPacketCaptureStorageLocationResponseOutput() PacketCaptureStorageLocationResponseOutput {
	return i.ToPacketCaptureStorageLocationResponseOutputWithContext(context.Background())
}

func (i PacketCaptureStorageLocationResponseArgs) ToPacketCaptureStorageLocationResponseOutputWithContext(ctx context.Context) PacketCaptureStorageLocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureStorageLocationResponseOutput)
}

func (i PacketCaptureStorageLocationResponseArgs) ToPacketCaptureStorageLocationResponsePtrOutput() PacketCaptureStorageLocationResponsePtrOutput {
	return i.ToPacketCaptureStorageLocationResponsePtrOutputWithContext(context.Background())
}

func (i PacketCaptureStorageLocationResponseArgs) ToPacketCaptureStorageLocationResponsePtrOutputWithContext(ctx context.Context) PacketCaptureStorageLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureStorageLocationResponseOutput).ToPacketCaptureStorageLocationResponsePtrOutputWithContext(ctx)
}









type PacketCaptureStorageLocationResponsePtrInput interface {
	pulumi.Input

	ToPacketCaptureStorageLocationResponsePtrOutput() PacketCaptureStorageLocationResponsePtrOutput
	ToPacketCaptureStorageLocationResponsePtrOutputWithContext(context.Context) PacketCaptureStorageLocationResponsePtrOutput
}

type packetCaptureStorageLocationResponsePtrType PacketCaptureStorageLocationResponseArgs

func PacketCaptureStorageLocationResponsePtr(v *PacketCaptureStorageLocationResponseArgs) PacketCaptureStorageLocationResponsePtrInput {
	return (*packetCaptureStorageLocationResponsePtrType)(v)
}

func (*packetCaptureStorageLocationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCaptureStorageLocationResponse)(nil)).Elem()
}

func (i *packetCaptureStorageLocationResponsePtrType) ToPacketCaptureStorageLocationResponsePtrOutput() PacketCaptureStorageLocationResponsePtrOutput {
	return i.ToPacketCaptureStorageLocationResponsePtrOutputWithContext(context.Background())
}

func (i *packetCaptureStorageLocationResponsePtrType) ToPacketCaptureStorageLocationResponsePtrOutputWithContext(ctx context.Context) PacketCaptureStorageLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureStorageLocationResponsePtrOutput)
}

type PacketCaptureStorageLocationResponseOutput struct{ *pulumi.OutputState }

func (PacketCaptureStorageLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureStorageLocationResponse)(nil)).Elem()
}

func (o PacketCaptureStorageLocationResponseOutput) ToPacketCaptureStorageLocationResponseOutput() PacketCaptureStorageLocationResponseOutput {
	return o
}

func (o PacketCaptureStorageLocationResponseOutput) ToPacketCaptureStorageLocationResponseOutputWithContext(ctx context.Context) PacketCaptureStorageLocationResponseOutput {
	return o
}

func (o PacketCaptureStorageLocationResponseOutput) ToPacketCaptureStorageLocationResponsePtrOutput() PacketCaptureStorageLocationResponsePtrOutput {
	return o.ToPacketCaptureStorageLocationResponsePtrOutputWithContext(context.Background())
}

func (o PacketCaptureStorageLocationResponseOutput) ToPacketCaptureStorageLocationResponsePtrOutputWithContext(ctx context.Context) PacketCaptureStorageLocationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PacketCaptureStorageLocationResponse) *PacketCaptureStorageLocationResponse {
		return &v
	}).(PacketCaptureStorageLocationResponsePtrOutput)
}

func (o PacketCaptureStorageLocationResponseOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocationResponse) *string { return v.FilePath }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationResponseOutput) StorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocationResponse) *string { return v.StorageId }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationResponseOutput) StoragePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocationResponse) *string { return v.StoragePath }).(pulumi.StringPtrOutput)
}

type PacketCaptureStorageLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (PacketCaptureStorageLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCaptureStorageLocationResponse)(nil)).Elem()
}

func (o PacketCaptureStorageLocationResponsePtrOutput) ToPacketCaptureStorageLocationResponsePtrOutput() PacketCaptureStorageLocationResponsePtrOutput {
	return o
}

func (o PacketCaptureStorageLocationResponsePtrOutput) ToPacketCaptureStorageLocationResponsePtrOutputWithContext(ctx context.Context) PacketCaptureStorageLocationResponsePtrOutput {
	return o
}

func (o PacketCaptureStorageLocationResponsePtrOutput) Elem() PacketCaptureStorageLocationResponseOutput {
	return o.ApplyT(func(v *PacketCaptureStorageLocationResponse) PacketCaptureStorageLocationResponse {
		if v != nil {
			return *v
		}
		var ret PacketCaptureStorageLocationResponse
		return ret
	}).(PacketCaptureStorageLocationResponseOutput)
}

func (o PacketCaptureStorageLocationResponsePtrOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCaptureStorageLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.FilePath
	}).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationResponsePtrOutput) StorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCaptureStorageLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageId
	}).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationResponsePtrOutput) StoragePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCaptureStorageLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StoragePath
	}).(pulumi.StringPtrOutput)
}

type Probe struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	IntervalInSeconds *int    `pulumi:"intervalInSeconds"`
	Name              *string `pulumi:"name"`
	NumberOfProbes    *int    `pulumi:"numberOfProbes"`
	Port              int     `pulumi:"port"`
	Protocol          string  `pulumi:"protocol"`
	ProvisioningState *string `pulumi:"provisioningState"`
	RequestPath       *string `pulumi:"requestPath"`
}





type ProbeInput interface {
	pulumi.Input

	ToProbeOutput() ProbeOutput
	ToProbeOutputWithContext(context.Context) ProbeOutput
}

type ProbeArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	IntervalInSeconds pulumi.IntPtrInput    `pulumi:"intervalInSeconds"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	NumberOfProbes    pulumi.IntPtrInput    `pulumi:"numberOfProbes"`
	Port              pulumi.IntInput       `pulumi:"port"`
	Protocol          pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	RequestPath       pulumi.StringPtrInput `pulumi:"requestPath"`
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

type PublicIPAddressType struct {
	DnsSettings              *PublicIPAddressDnsSettings `pulumi:"dnsSettings"`
	Etag                     *string                     `pulumi:"etag"`
	Id                       *string                     `pulumi:"id"`
	IdleTimeoutInMinutes     *int                        `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                     `pulumi:"ipAddress"`
	IpTags                   []IpTag                     `pulumi:"ipTags"`
	Location                 *string                     `pulumi:"location"`
	ProvisioningState        *string                     `pulumi:"provisioningState"`
	PublicIPAddressVersion   *string                     `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                     `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             *string                     `pulumi:"resourceGuid"`
	Sku                      *PublicIPAddressSku         `pulumi:"sku"`
	Tags                     map[string]string           `pulumi:"tags"`
	Zones                    []string                    `pulumi:"zones"`
}





type PublicIPAddressTypeInput interface {
	pulumi.Input

	ToPublicIPAddressTypeOutput() PublicIPAddressTypeOutput
	ToPublicIPAddressTypeOutputWithContext(context.Context) PublicIPAddressTypeOutput
}

type PublicIPAddressTypeArgs struct {
	DnsSettings              PublicIPAddressDnsSettingsPtrInput `pulumi:"dnsSettings"`
	Etag                     pulumi.StringPtrInput              `pulumi:"etag"`
	Id                       pulumi.StringPtrInput              `pulumi:"id"`
	IdleTimeoutInMinutes     pulumi.IntPtrInput                 `pulumi:"idleTimeoutInMinutes"`
	IpAddress                pulumi.StringPtrInput              `pulumi:"ipAddress"`
	IpTags                   IpTagArrayInput                    `pulumi:"ipTags"`
	Location                 pulumi.StringPtrInput              `pulumi:"location"`
	ProvisioningState        pulumi.StringPtrInput              `pulumi:"provisioningState"`
	PublicIPAddressVersion   pulumi.StringPtrInput              `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod pulumi.StringPtrInput              `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             pulumi.StringPtrInput              `pulumi:"resourceGuid"`
	Sku                      PublicIPAddressSkuPtrInput         `pulumi:"sku"`
	Tags                     pulumi.StringMapInput              `pulumi:"tags"`
	Zones                    pulumi.StringArrayInput            `pulumi:"zones"`
}

func (PublicIPAddressTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressType)(nil)).Elem()
}

func (i PublicIPAddressTypeArgs) ToPublicIPAddressTypeOutput() PublicIPAddressTypeOutput {
	return i.ToPublicIPAddressTypeOutputWithContext(context.Background())
}

func (i PublicIPAddressTypeArgs) ToPublicIPAddressTypeOutputWithContext(ctx context.Context) PublicIPAddressTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressTypeOutput)
}

func (i PublicIPAddressTypeArgs) ToPublicIPAddressTypePtrOutput() PublicIPAddressTypePtrOutput {
	return i.ToPublicIPAddressTypePtrOutputWithContext(context.Background())
}

func (i PublicIPAddressTypeArgs) ToPublicIPAddressTypePtrOutputWithContext(ctx context.Context) PublicIPAddressTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressTypeOutput).ToPublicIPAddressTypePtrOutputWithContext(ctx)
}









type PublicIPAddressTypePtrInput interface {
	pulumi.Input

	ToPublicIPAddressTypePtrOutput() PublicIPAddressTypePtrOutput
	ToPublicIPAddressTypePtrOutputWithContext(context.Context) PublicIPAddressTypePtrOutput
}

type publicIPAddressTypePtrType PublicIPAddressTypeArgs

func PublicIPAddressTypePtr(v *PublicIPAddressTypeArgs) PublicIPAddressTypePtrInput {
	return (*publicIPAddressTypePtrType)(v)
}

func (*publicIPAddressTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressType)(nil)).Elem()
}

func (i *publicIPAddressTypePtrType) ToPublicIPAddressTypePtrOutput() PublicIPAddressTypePtrOutput {
	return i.ToPublicIPAddressTypePtrOutputWithContext(context.Background())
}

func (i *publicIPAddressTypePtrType) ToPublicIPAddressTypePtrOutputWithContext(ctx context.Context) PublicIPAddressTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressTypePtrOutput)
}

type PublicIPAddressTypeOutput struct{ *pulumi.OutputState }

func (PublicIPAddressTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressType)(nil)).Elem()
}

func (o PublicIPAddressTypeOutput) ToPublicIPAddressTypeOutput() PublicIPAddressTypeOutput {
	return o
}

func (o PublicIPAddressTypeOutput) ToPublicIPAddressTypeOutputWithContext(ctx context.Context) PublicIPAddressTypeOutput {
	return o
}

func (o PublicIPAddressTypeOutput) ToPublicIPAddressTypePtrOutput() PublicIPAddressTypePtrOutput {
	return o.ToPublicIPAddressTypePtrOutputWithContext(context.Background())
}

func (o PublicIPAddressTypeOutput) ToPublicIPAddressTypePtrOutputWithContext(ctx context.Context) PublicIPAddressTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressType) *PublicIPAddressType {
		return &v
	}).(PublicIPAddressTypePtrOutput)
}

func (o PublicIPAddressTypeOutput) DnsSettings() PublicIPAddressDnsSettingsPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *PublicIPAddressDnsSettings { return v.DnsSettings }).(PublicIPAddressDnsSettingsPtrOutput)
}

func (o PublicIPAddressTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o PublicIPAddressTypeOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) IpTags() IpTagArrayOutput {
	return o.ApplyT(func(v PublicIPAddressType) []IpTag { return v.IpTags }).(IpTagArrayOutput)
}

func (o PublicIPAddressTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) Sku() PublicIPAddressSkuPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *PublicIPAddressSku { return v.Sku }).(PublicIPAddressSkuPtrOutput)
}

func (o PublicIPAddressTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v PublicIPAddressType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PublicIPAddressTypeOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PublicIPAddressType) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

type PublicIPAddressTypePtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressType)(nil)).Elem()
}

func (o PublicIPAddressTypePtrOutput) ToPublicIPAddressTypePtrOutput() PublicIPAddressTypePtrOutput {
	return o
}

func (o PublicIPAddressTypePtrOutput) ToPublicIPAddressTypePtrOutputWithContext(ctx context.Context) PublicIPAddressTypePtrOutput {
	return o
}

func (o PublicIPAddressTypePtrOutput) Elem() PublicIPAddressTypeOutput {
	return o.ApplyT(func(v *PublicIPAddressType) PublicIPAddressType {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressType
		return ret
	}).(PublicIPAddressTypeOutput)
}

func (o PublicIPAddressTypePtrOutput) DnsSettings() PublicIPAddressDnsSettingsPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *PublicIPAddressDnsSettings {
		if v == nil {
			return nil
		}
		return v.DnsSettings
	}).(PublicIPAddressDnsSettingsPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) IpTags() IpTagArrayOutput {
	return o.ApplyT(func(v *PublicIPAddressType) []IpTag {
		if v == nil {
			return nil
		}
		return v.IpTags
	}).(IpTagArrayOutput)
}

func (o PublicIPAddressTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) Sku() PublicIPAddressSkuPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *PublicIPAddressSku {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(PublicIPAddressSkuPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PublicIPAddressType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o PublicIPAddressTypePtrOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PublicIPAddressType) []string {
		if v == nil {
			return nil
		}
		return v.Zones
	}).(pulumi.StringArrayOutput)
}

type PublicIPAddressDnsSettings struct {
	DomainNameLabel *string `pulumi:"domainNameLabel"`
	Fqdn            *string `pulumi:"fqdn"`
	ReverseFqdn     *string `pulumi:"reverseFqdn"`
}





type PublicIPAddressDnsSettingsInput interface {
	pulumi.Input

	ToPublicIPAddressDnsSettingsOutput() PublicIPAddressDnsSettingsOutput
	ToPublicIPAddressDnsSettingsOutputWithContext(context.Context) PublicIPAddressDnsSettingsOutput
}

type PublicIPAddressDnsSettingsArgs struct {
	DomainNameLabel pulumi.StringPtrInput `pulumi:"domainNameLabel"`
	Fqdn            pulumi.StringPtrInput `pulumi:"fqdn"`
	ReverseFqdn     pulumi.StringPtrInput `pulumi:"reverseFqdn"`
}

func (PublicIPAddressDnsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressDnsSettings)(nil)).Elem()
}

func (i PublicIPAddressDnsSettingsArgs) ToPublicIPAddressDnsSettingsOutput() PublicIPAddressDnsSettingsOutput {
	return i.ToPublicIPAddressDnsSettingsOutputWithContext(context.Background())
}

func (i PublicIPAddressDnsSettingsArgs) ToPublicIPAddressDnsSettingsOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressDnsSettingsOutput)
}

func (i PublicIPAddressDnsSettingsArgs) ToPublicIPAddressDnsSettingsPtrOutput() PublicIPAddressDnsSettingsPtrOutput {
	return i.ToPublicIPAddressDnsSettingsPtrOutputWithContext(context.Background())
}

func (i PublicIPAddressDnsSettingsArgs) ToPublicIPAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressDnsSettingsOutput).ToPublicIPAddressDnsSettingsPtrOutputWithContext(ctx)
}









type PublicIPAddressDnsSettingsPtrInput interface {
	pulumi.Input

	ToPublicIPAddressDnsSettingsPtrOutput() PublicIPAddressDnsSettingsPtrOutput
	ToPublicIPAddressDnsSettingsPtrOutputWithContext(context.Context) PublicIPAddressDnsSettingsPtrOutput
}

type publicIPAddressDnsSettingsPtrType PublicIPAddressDnsSettingsArgs

func PublicIPAddressDnsSettingsPtr(v *PublicIPAddressDnsSettingsArgs) PublicIPAddressDnsSettingsPtrInput {
	return (*publicIPAddressDnsSettingsPtrType)(v)
}

func (*publicIPAddressDnsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressDnsSettings)(nil)).Elem()
}

func (i *publicIPAddressDnsSettingsPtrType) ToPublicIPAddressDnsSettingsPtrOutput() PublicIPAddressDnsSettingsPtrOutput {
	return i.ToPublicIPAddressDnsSettingsPtrOutputWithContext(context.Background())
}

func (i *publicIPAddressDnsSettingsPtrType) ToPublicIPAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressDnsSettingsPtrOutput)
}

type PublicIPAddressDnsSettingsOutput struct{ *pulumi.OutputState }

func (PublicIPAddressDnsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressDnsSettings)(nil)).Elem()
}

func (o PublicIPAddressDnsSettingsOutput) ToPublicIPAddressDnsSettingsOutput() PublicIPAddressDnsSettingsOutput {
	return o
}

func (o PublicIPAddressDnsSettingsOutput) ToPublicIPAddressDnsSettingsOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsOutput {
	return o
}

func (o PublicIPAddressDnsSettingsOutput) ToPublicIPAddressDnsSettingsPtrOutput() PublicIPAddressDnsSettingsPtrOutput {
	return o.ToPublicIPAddressDnsSettingsPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressDnsSettingsOutput) ToPublicIPAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressDnsSettings) *PublicIPAddressDnsSettings {
		return &v
	}).(PublicIPAddressDnsSettingsPtrOutput)
}

func (o PublicIPAddressDnsSettingsOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettings) *string { return v.DomainNameLabel }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettings) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettings) *string { return v.ReverseFqdn }).(pulumi.StringPtrOutput)
}

type PublicIPAddressDnsSettingsPtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressDnsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressDnsSettings)(nil)).Elem()
}

func (o PublicIPAddressDnsSettingsPtrOutput) ToPublicIPAddressDnsSettingsPtrOutput() PublicIPAddressDnsSettingsPtrOutput {
	return o
}

func (o PublicIPAddressDnsSettingsPtrOutput) ToPublicIPAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsPtrOutput {
	return o
}

func (o PublicIPAddressDnsSettingsPtrOutput) Elem() PublicIPAddressDnsSettingsOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettings) PublicIPAddressDnsSettings {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressDnsSettings
		return ret
	}).(PublicIPAddressDnsSettingsOutput)
}

func (o PublicIPAddressDnsSettingsPtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsPtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsPtrOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReverseFqdn
	}).(pulumi.StringPtrOutput)
}

type PublicIPAddressDnsSettingsResponse struct {
	DomainNameLabel *string `pulumi:"domainNameLabel"`
	Fqdn            *string `pulumi:"fqdn"`
	ReverseFqdn     *string `pulumi:"reverseFqdn"`
}





type PublicIPAddressDnsSettingsResponseInput interface {
	pulumi.Input

	ToPublicIPAddressDnsSettingsResponseOutput() PublicIPAddressDnsSettingsResponseOutput
	ToPublicIPAddressDnsSettingsResponseOutputWithContext(context.Context) PublicIPAddressDnsSettingsResponseOutput
}

type PublicIPAddressDnsSettingsResponseArgs struct {
	DomainNameLabel pulumi.StringPtrInput `pulumi:"domainNameLabel"`
	Fqdn            pulumi.StringPtrInput `pulumi:"fqdn"`
	ReverseFqdn     pulumi.StringPtrInput `pulumi:"reverseFqdn"`
}

func (PublicIPAddressDnsSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressDnsSettingsResponse)(nil)).Elem()
}

func (i PublicIPAddressDnsSettingsResponseArgs) ToPublicIPAddressDnsSettingsResponseOutput() PublicIPAddressDnsSettingsResponseOutput {
	return i.ToPublicIPAddressDnsSettingsResponseOutputWithContext(context.Background())
}

func (i PublicIPAddressDnsSettingsResponseArgs) ToPublicIPAddressDnsSettingsResponseOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressDnsSettingsResponseOutput)
}

func (i PublicIPAddressDnsSettingsResponseArgs) ToPublicIPAddressDnsSettingsResponsePtrOutput() PublicIPAddressDnsSettingsResponsePtrOutput {
	return i.ToPublicIPAddressDnsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i PublicIPAddressDnsSettingsResponseArgs) ToPublicIPAddressDnsSettingsResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressDnsSettingsResponseOutput).ToPublicIPAddressDnsSettingsResponsePtrOutputWithContext(ctx)
}









type PublicIPAddressDnsSettingsResponsePtrInput interface {
	pulumi.Input

	ToPublicIPAddressDnsSettingsResponsePtrOutput() PublicIPAddressDnsSettingsResponsePtrOutput
	ToPublicIPAddressDnsSettingsResponsePtrOutputWithContext(context.Context) PublicIPAddressDnsSettingsResponsePtrOutput
}

type publicIPAddressDnsSettingsResponsePtrType PublicIPAddressDnsSettingsResponseArgs

func PublicIPAddressDnsSettingsResponsePtr(v *PublicIPAddressDnsSettingsResponseArgs) PublicIPAddressDnsSettingsResponsePtrInput {
	return (*publicIPAddressDnsSettingsResponsePtrType)(v)
}

func (*publicIPAddressDnsSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressDnsSettingsResponse)(nil)).Elem()
}

func (i *publicIPAddressDnsSettingsResponsePtrType) ToPublicIPAddressDnsSettingsResponsePtrOutput() PublicIPAddressDnsSettingsResponsePtrOutput {
	return i.ToPublicIPAddressDnsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *publicIPAddressDnsSettingsResponsePtrType) ToPublicIPAddressDnsSettingsResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

type PublicIPAddressDnsSettingsResponseOutput struct{ *pulumi.OutputState }

func (PublicIPAddressDnsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressDnsSettingsResponse)(nil)).Elem()
}

func (o PublicIPAddressDnsSettingsResponseOutput) ToPublicIPAddressDnsSettingsResponseOutput() PublicIPAddressDnsSettingsResponseOutput {
	return o
}

func (o PublicIPAddressDnsSettingsResponseOutput) ToPublicIPAddressDnsSettingsResponseOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsResponseOutput {
	return o
}

func (o PublicIPAddressDnsSettingsResponseOutput) ToPublicIPAddressDnsSettingsResponsePtrOutput() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ToPublicIPAddressDnsSettingsResponsePtrOutputWithContext(context.Background())
}

func (o PublicIPAddressDnsSettingsResponseOutput) ToPublicIPAddressDnsSettingsResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressDnsSettingsResponse) *PublicIPAddressDnsSettingsResponse {
		return &v
	}).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

func (o PublicIPAddressDnsSettingsResponseOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettingsResponse) *string { return v.DomainNameLabel }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettingsResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsResponseOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettingsResponse) *string { return v.ReverseFqdn }).(pulumi.StringPtrOutput)
}

type PublicIPAddressDnsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressDnsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressDnsSettingsResponse)(nil)).Elem()
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) ToPublicIPAddressDnsSettingsResponsePtrOutput() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) ToPublicIPAddressDnsSettingsResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsResponsePtrOutput {
	return o
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) Elem() PublicIPAddressDnsSettingsResponseOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettingsResponse) PublicIPAddressDnsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressDnsSettingsResponse
		return ret
	}).(PublicIPAddressDnsSettingsResponseOutput)
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReverseFqdn
	}).(pulumi.StringPtrOutput)
}

type PublicIPAddressResponse struct {
	DnsSettings              *PublicIPAddressDnsSettingsResponse `pulumi:"dnsSettings"`
	Etag                     *string                             `pulumi:"etag"`
	Id                       *string                             `pulumi:"id"`
	IdleTimeoutInMinutes     *int                                `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                             `pulumi:"ipAddress"`
	IpConfiguration          IPConfigurationResponse             `pulumi:"ipConfiguration"`
	IpTags                   []IpTagResponse                     `pulumi:"ipTags"`
	Location                 *string                             `pulumi:"location"`
	Name                     string                              `pulumi:"name"`
	ProvisioningState        *string                             `pulumi:"provisioningState"`
	PublicIPAddressVersion   *string                             `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                             `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             *string                             `pulumi:"resourceGuid"`
	Sku                      *PublicIPAddressSkuResponse         `pulumi:"sku"`
	Tags                     map[string]string                   `pulumi:"tags"`
	Type                     string                              `pulumi:"type"`
	Zones                    []string                            `pulumi:"zones"`
}





type PublicIPAddressResponseInput interface {
	pulumi.Input

	ToPublicIPAddressResponseOutput() PublicIPAddressResponseOutput
	ToPublicIPAddressResponseOutputWithContext(context.Context) PublicIPAddressResponseOutput
}

type PublicIPAddressResponseArgs struct {
	DnsSettings              PublicIPAddressDnsSettingsResponsePtrInput `pulumi:"dnsSettings"`
	Etag                     pulumi.StringPtrInput                      `pulumi:"etag"`
	Id                       pulumi.StringPtrInput                      `pulumi:"id"`
	IdleTimeoutInMinutes     pulumi.IntPtrInput                         `pulumi:"idleTimeoutInMinutes"`
	IpAddress                pulumi.StringPtrInput                      `pulumi:"ipAddress"`
	IpConfiguration          IPConfigurationResponseInput               `pulumi:"ipConfiguration"`
	IpTags                   IpTagResponseArrayInput                    `pulumi:"ipTags"`
	Location                 pulumi.StringPtrInput                      `pulumi:"location"`
	Name                     pulumi.StringInput                         `pulumi:"name"`
	ProvisioningState        pulumi.StringPtrInput                      `pulumi:"provisioningState"`
	PublicIPAddressVersion   pulumi.StringPtrInput                      `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod pulumi.StringPtrInput                      `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             pulumi.StringPtrInput                      `pulumi:"resourceGuid"`
	Sku                      PublicIPAddressSkuResponsePtrInput         `pulumi:"sku"`
	Tags                     pulumi.StringMapInput                      `pulumi:"tags"`
	Type                     pulumi.StringInput                         `pulumi:"type"`
	Zones                    pulumi.StringArrayInput                    `pulumi:"zones"`
}

func (PublicIPAddressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResponse)(nil)).Elem()
}

func (i PublicIPAddressResponseArgs) ToPublicIPAddressResponseOutput() PublicIPAddressResponseOutput {
	return i.ToPublicIPAddressResponseOutputWithContext(context.Background())
}

func (i PublicIPAddressResponseArgs) ToPublicIPAddressResponseOutputWithContext(ctx context.Context) PublicIPAddressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressResponseOutput)
}

func (i PublicIPAddressResponseArgs) ToPublicIPAddressResponsePtrOutput() PublicIPAddressResponsePtrOutput {
	return i.ToPublicIPAddressResponsePtrOutputWithContext(context.Background())
}

func (i PublicIPAddressResponseArgs) ToPublicIPAddressResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressResponseOutput).ToPublicIPAddressResponsePtrOutputWithContext(ctx)
}









type PublicIPAddressResponsePtrInput interface {
	pulumi.Input

	ToPublicIPAddressResponsePtrOutput() PublicIPAddressResponsePtrOutput
	ToPublicIPAddressResponsePtrOutputWithContext(context.Context) PublicIPAddressResponsePtrOutput
}

type publicIPAddressResponsePtrType PublicIPAddressResponseArgs

func PublicIPAddressResponsePtr(v *PublicIPAddressResponseArgs) PublicIPAddressResponsePtrInput {
	return (*publicIPAddressResponsePtrType)(v)
}

func (*publicIPAddressResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressResponse)(nil)).Elem()
}

func (i *publicIPAddressResponsePtrType) ToPublicIPAddressResponsePtrOutput() PublicIPAddressResponsePtrOutput {
	return i.ToPublicIPAddressResponsePtrOutputWithContext(context.Background())
}

func (i *publicIPAddressResponsePtrType) ToPublicIPAddressResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressResponsePtrOutput)
}

type PublicIPAddressResponseOutput struct{ *pulumi.OutputState }

func (PublicIPAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResponse)(nil)).Elem()
}

func (o PublicIPAddressResponseOutput) ToPublicIPAddressResponseOutput() PublicIPAddressResponseOutput {
	return o
}

func (o PublicIPAddressResponseOutput) ToPublicIPAddressResponseOutputWithContext(ctx context.Context) PublicIPAddressResponseOutput {
	return o
}

func (o PublicIPAddressResponseOutput) ToPublicIPAddressResponsePtrOutput() PublicIPAddressResponsePtrOutput {
	return o.ToPublicIPAddressResponsePtrOutputWithContext(context.Background())
}

func (o PublicIPAddressResponseOutput) ToPublicIPAddressResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressResponse) *PublicIPAddressResponse {
		return &v
	}).(PublicIPAddressResponsePtrOutput)
}

func (o PublicIPAddressResponseOutput) DnsSettings() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *PublicIPAddressDnsSettingsResponse { return v.DnsSettings }).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

func (o PublicIPAddressResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o PublicIPAddressResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) IpConfiguration() IPConfigurationResponseOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) IPConfigurationResponse { return v.IpConfiguration }).(IPConfigurationResponseOutput)
}

func (o PublicIPAddressResponseOutput) IpTags() IpTagResponseArrayOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) []IpTagResponse { return v.IpTags }).(IpTagResponseArrayOutput)
}

func (o PublicIPAddressResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PublicIPAddressResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) Sku() PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *PublicIPAddressSkuResponse { return v.Sku }).(PublicIPAddressSkuResponsePtrOutput)
}

func (o PublicIPAddressResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PublicIPAddressResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o PublicIPAddressResponseOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

type PublicIPAddressResponsePtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressResponse)(nil)).Elem()
}

func (o PublicIPAddressResponsePtrOutput) ToPublicIPAddressResponsePtrOutput() PublicIPAddressResponsePtrOutput {
	return o
}

func (o PublicIPAddressResponsePtrOutput) ToPublicIPAddressResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressResponsePtrOutput {
	return o
}

func (o PublicIPAddressResponsePtrOutput) Elem() PublicIPAddressResponseOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) PublicIPAddressResponse {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressResponse
		return ret
	}).(PublicIPAddressResponseOutput)
}

func (o PublicIPAddressResponsePtrOutput) DnsSettings() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *PublicIPAddressDnsSettingsResponse {
		if v == nil {
			return nil
		}
		return v.DnsSettings
	}).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) IpConfiguration() IPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *IPConfigurationResponse {
		if v == nil {
			return nil
		}
		return &v.IpConfiguration
	}).(IPConfigurationResponsePtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) IpTags() IpTagResponseArrayOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) []IpTagResponse {
		if v == nil {
			return nil
		}
		return v.IpTags
	}).(IpTagResponseArrayOutput)
}

func (o PublicIPAddressResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Sku() PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *PublicIPAddressSkuResponse {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(PublicIPAddressSkuResponsePtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o PublicIPAddressResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) []string {
		if v == nil {
			return nil
		}
		return v.Zones
	}).(pulumi.StringArrayOutput)
}

type PublicIPAddressSku struct {
	Name *string `pulumi:"name"`
}





type PublicIPAddressSkuInput interface {
	pulumi.Input

	ToPublicIPAddressSkuOutput() PublicIPAddressSkuOutput
	ToPublicIPAddressSkuOutputWithContext(context.Context) PublicIPAddressSkuOutput
}

type PublicIPAddressSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (PublicIPAddressSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSku)(nil)).Elem()
}

func (i PublicIPAddressSkuArgs) ToPublicIPAddressSkuOutput() PublicIPAddressSkuOutput {
	return i.ToPublicIPAddressSkuOutputWithContext(context.Background())
}

func (i PublicIPAddressSkuArgs) ToPublicIPAddressSkuOutputWithContext(ctx context.Context) PublicIPAddressSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressSkuOutput)
}

func (i PublicIPAddressSkuArgs) ToPublicIPAddressSkuPtrOutput() PublicIPAddressSkuPtrOutput {
	return i.ToPublicIPAddressSkuPtrOutputWithContext(context.Background())
}

func (i PublicIPAddressSkuArgs) ToPublicIPAddressSkuPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressSkuOutput).ToPublicIPAddressSkuPtrOutputWithContext(ctx)
}









type PublicIPAddressSkuPtrInput interface {
	pulumi.Input

	ToPublicIPAddressSkuPtrOutput() PublicIPAddressSkuPtrOutput
	ToPublicIPAddressSkuPtrOutputWithContext(context.Context) PublicIPAddressSkuPtrOutput
}

type publicIPAddressSkuPtrType PublicIPAddressSkuArgs

func PublicIPAddressSkuPtr(v *PublicIPAddressSkuArgs) PublicIPAddressSkuPtrInput {
	return (*publicIPAddressSkuPtrType)(v)
}

func (*publicIPAddressSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressSku)(nil)).Elem()
}

func (i *publicIPAddressSkuPtrType) ToPublicIPAddressSkuPtrOutput() PublicIPAddressSkuPtrOutput {
	return i.ToPublicIPAddressSkuPtrOutputWithContext(context.Background())
}

func (i *publicIPAddressSkuPtrType) ToPublicIPAddressSkuPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressSkuPtrOutput)
}

type PublicIPAddressSkuOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSku)(nil)).Elem()
}

func (o PublicIPAddressSkuOutput) ToPublicIPAddressSkuOutput() PublicIPAddressSkuOutput {
	return o
}

func (o PublicIPAddressSkuOutput) ToPublicIPAddressSkuOutputWithContext(ctx context.Context) PublicIPAddressSkuOutput {
	return o
}

func (o PublicIPAddressSkuOutput) ToPublicIPAddressSkuPtrOutput() PublicIPAddressSkuPtrOutput {
	return o.ToPublicIPAddressSkuPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuOutput) ToPublicIPAddressSkuPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressSku) *PublicIPAddressSku {
		return &v
	}).(PublicIPAddressSkuPtrOutput)
}

func (o PublicIPAddressSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type PublicIPAddressSkuPtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressSku)(nil)).Elem()
}

func (o PublicIPAddressSkuPtrOutput) ToPublicIPAddressSkuPtrOutput() PublicIPAddressSkuPtrOutput {
	return o
}

func (o PublicIPAddressSkuPtrOutput) ToPublicIPAddressSkuPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuPtrOutput {
	return o
}

func (o PublicIPAddressSkuPtrOutput) Elem() PublicIPAddressSkuOutput {
	return o.ApplyT(func(v *PublicIPAddressSku) PublicIPAddressSku {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressSku
		return ret
	}).(PublicIPAddressSkuOutput)
}

func (o PublicIPAddressSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type PublicIPAddressSkuResponse struct {
	Name *string `pulumi:"name"`
}





type PublicIPAddressSkuResponseInput interface {
	pulumi.Input

	ToPublicIPAddressSkuResponseOutput() PublicIPAddressSkuResponseOutput
	ToPublicIPAddressSkuResponseOutputWithContext(context.Context) PublicIPAddressSkuResponseOutput
}

type PublicIPAddressSkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (PublicIPAddressSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSkuResponse)(nil)).Elem()
}

func (i PublicIPAddressSkuResponseArgs) ToPublicIPAddressSkuResponseOutput() PublicIPAddressSkuResponseOutput {
	return i.ToPublicIPAddressSkuResponseOutputWithContext(context.Background())
}

func (i PublicIPAddressSkuResponseArgs) ToPublicIPAddressSkuResponseOutputWithContext(ctx context.Context) PublicIPAddressSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressSkuResponseOutput)
}

func (i PublicIPAddressSkuResponseArgs) ToPublicIPAddressSkuResponsePtrOutput() PublicIPAddressSkuResponsePtrOutput {
	return i.ToPublicIPAddressSkuResponsePtrOutputWithContext(context.Background())
}

func (i PublicIPAddressSkuResponseArgs) ToPublicIPAddressSkuResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressSkuResponseOutput).ToPublicIPAddressSkuResponsePtrOutputWithContext(ctx)
}









type PublicIPAddressSkuResponsePtrInput interface {
	pulumi.Input

	ToPublicIPAddressSkuResponsePtrOutput() PublicIPAddressSkuResponsePtrOutput
	ToPublicIPAddressSkuResponsePtrOutputWithContext(context.Context) PublicIPAddressSkuResponsePtrOutput
}

type publicIPAddressSkuResponsePtrType PublicIPAddressSkuResponseArgs

func PublicIPAddressSkuResponsePtr(v *PublicIPAddressSkuResponseArgs) PublicIPAddressSkuResponsePtrInput {
	return (*publicIPAddressSkuResponsePtrType)(v)
}

func (*publicIPAddressSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressSkuResponse)(nil)).Elem()
}

func (i *publicIPAddressSkuResponsePtrType) ToPublicIPAddressSkuResponsePtrOutput() PublicIPAddressSkuResponsePtrOutput {
	return i.ToPublicIPAddressSkuResponsePtrOutputWithContext(context.Background())
}

func (i *publicIPAddressSkuResponsePtrType) ToPublicIPAddressSkuResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressSkuResponsePtrOutput)
}

type PublicIPAddressSkuResponseOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSkuResponse)(nil)).Elem()
}

func (o PublicIPAddressSkuResponseOutput) ToPublicIPAddressSkuResponseOutput() PublicIPAddressSkuResponseOutput {
	return o
}

func (o PublicIPAddressSkuResponseOutput) ToPublicIPAddressSkuResponseOutputWithContext(ctx context.Context) PublicIPAddressSkuResponseOutput {
	return o
}

func (o PublicIPAddressSkuResponseOutput) ToPublicIPAddressSkuResponsePtrOutput() PublicIPAddressSkuResponsePtrOutput {
	return o.ToPublicIPAddressSkuResponsePtrOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuResponseOutput) ToPublicIPAddressSkuResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressSkuResponse) *PublicIPAddressSkuResponse {
		return &v
	}).(PublicIPAddressSkuResponsePtrOutput)
}

func (o PublicIPAddressSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type PublicIPAddressSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressSkuResponse)(nil)).Elem()
}

func (o PublicIPAddressSkuResponsePtrOutput) ToPublicIPAddressSkuResponsePtrOutput() PublicIPAddressSkuResponsePtrOutput {
	return o
}

func (o PublicIPAddressSkuResponsePtrOutput) ToPublicIPAddressSkuResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressSkuResponsePtrOutput {
	return o
}

func (o PublicIPAddressSkuResponsePtrOutput) Elem() PublicIPAddressSkuResponseOutput {
	return o.ApplyT(func(v *PublicIPAddressSkuResponse) PublicIPAddressSkuResponse {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressSkuResponse
		return ret
	}).(PublicIPAddressSkuResponseOutput)
}

func (o PublicIPAddressSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ResourceNavigationLink struct {
	Id                 *string `pulumi:"id"`
	Link               *string `pulumi:"link"`
	LinkedResourceType *string `pulumi:"linkedResourceType"`
	Name               *string `pulumi:"name"`
}





type ResourceNavigationLinkInput interface {
	pulumi.Input

	ToResourceNavigationLinkOutput() ResourceNavigationLinkOutput
	ToResourceNavigationLinkOutputWithContext(context.Context) ResourceNavigationLinkOutput
}

type ResourceNavigationLinkArgs struct {
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	Link               pulumi.StringPtrInput `pulumi:"link"`
	LinkedResourceType pulumi.StringPtrInput `pulumi:"linkedResourceType"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
}

func (ResourceNavigationLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceNavigationLink)(nil)).Elem()
}

func (i ResourceNavigationLinkArgs) ToResourceNavigationLinkOutput() ResourceNavigationLinkOutput {
	return i.ToResourceNavigationLinkOutputWithContext(context.Background())
}

func (i ResourceNavigationLinkArgs) ToResourceNavigationLinkOutputWithContext(ctx context.Context) ResourceNavigationLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceNavigationLinkOutput)
}





type ResourceNavigationLinkArrayInput interface {
	pulumi.Input

	ToResourceNavigationLinkArrayOutput() ResourceNavigationLinkArrayOutput
	ToResourceNavigationLinkArrayOutputWithContext(context.Context) ResourceNavigationLinkArrayOutput
}

type ResourceNavigationLinkArray []ResourceNavigationLinkInput

func (ResourceNavigationLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceNavigationLink)(nil)).Elem()
}

func (i ResourceNavigationLinkArray) ToResourceNavigationLinkArrayOutput() ResourceNavigationLinkArrayOutput {
	return i.ToResourceNavigationLinkArrayOutputWithContext(context.Background())
}

func (i ResourceNavigationLinkArray) ToResourceNavigationLinkArrayOutputWithContext(ctx context.Context) ResourceNavigationLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceNavigationLinkArrayOutput)
}

type ResourceNavigationLinkOutput struct{ *pulumi.OutputState }

func (ResourceNavigationLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceNavigationLink)(nil)).Elem()
}

func (o ResourceNavigationLinkOutput) ToResourceNavigationLinkOutput() ResourceNavigationLinkOutput {
	return o
}

func (o ResourceNavigationLinkOutput) ToResourceNavigationLinkOutputWithContext(ctx context.Context) ResourceNavigationLinkOutput {
	return o
}

func (o ResourceNavigationLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLink) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkOutput) LinkedResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLink) *string { return v.LinkedResourceType }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLink) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ResourceNavigationLinkArrayOutput struct{ *pulumi.OutputState }

func (ResourceNavigationLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceNavigationLink)(nil)).Elem()
}

func (o ResourceNavigationLinkArrayOutput) ToResourceNavigationLinkArrayOutput() ResourceNavigationLinkArrayOutput {
	return o
}

func (o ResourceNavigationLinkArrayOutput) ToResourceNavigationLinkArrayOutputWithContext(ctx context.Context) ResourceNavigationLinkArrayOutput {
	return o
}

func (o ResourceNavigationLinkArrayOutput) Index(i pulumi.IntInput) ResourceNavigationLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceNavigationLink {
		return vs[0].([]ResourceNavigationLink)[vs[1].(int)]
	}).(ResourceNavigationLinkOutput)
}

type ResourceNavigationLinkResponse struct {
	Etag               string  `pulumi:"etag"`
	Id                 *string `pulumi:"id"`
	Link               *string `pulumi:"link"`
	LinkedResourceType *string `pulumi:"linkedResourceType"`
	Name               *string `pulumi:"name"`
	ProvisioningState  string  `pulumi:"provisioningState"`
}





type ResourceNavigationLinkResponseInput interface {
	pulumi.Input

	ToResourceNavigationLinkResponseOutput() ResourceNavigationLinkResponseOutput
	ToResourceNavigationLinkResponseOutputWithContext(context.Context) ResourceNavigationLinkResponseOutput
}

type ResourceNavigationLinkResponseArgs struct {
	Etag               pulumi.StringInput    `pulumi:"etag"`
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	Link               pulumi.StringPtrInput `pulumi:"link"`
	LinkedResourceType pulumi.StringPtrInput `pulumi:"linkedResourceType"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState  pulumi.StringInput    `pulumi:"provisioningState"`
}

func (ResourceNavigationLinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceNavigationLinkResponse)(nil)).Elem()
}

func (i ResourceNavigationLinkResponseArgs) ToResourceNavigationLinkResponseOutput() ResourceNavigationLinkResponseOutput {
	return i.ToResourceNavigationLinkResponseOutputWithContext(context.Background())
}

func (i ResourceNavigationLinkResponseArgs) ToResourceNavigationLinkResponseOutputWithContext(ctx context.Context) ResourceNavigationLinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceNavigationLinkResponseOutput)
}





type ResourceNavigationLinkResponseArrayInput interface {
	pulumi.Input

	ToResourceNavigationLinkResponseArrayOutput() ResourceNavigationLinkResponseArrayOutput
	ToResourceNavigationLinkResponseArrayOutputWithContext(context.Context) ResourceNavigationLinkResponseArrayOutput
}

type ResourceNavigationLinkResponseArray []ResourceNavigationLinkResponseInput

func (ResourceNavigationLinkResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceNavigationLinkResponse)(nil)).Elem()
}

func (i ResourceNavigationLinkResponseArray) ToResourceNavigationLinkResponseArrayOutput() ResourceNavigationLinkResponseArrayOutput {
	return i.ToResourceNavigationLinkResponseArrayOutputWithContext(context.Background())
}

func (i ResourceNavigationLinkResponseArray) ToResourceNavigationLinkResponseArrayOutputWithContext(ctx context.Context) ResourceNavigationLinkResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceNavigationLinkResponseArrayOutput)
}

type ResourceNavigationLinkResponseOutput struct{ *pulumi.OutputState }

func (ResourceNavigationLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceNavigationLinkResponse)(nil)).Elem()
}

func (o ResourceNavigationLinkResponseOutput) ToResourceNavigationLinkResponseOutput() ResourceNavigationLinkResponseOutput {
	return o
}

func (o ResourceNavigationLinkResponseOutput) ToResourceNavigationLinkResponseOutputWithContext(ctx context.Context) ResourceNavigationLinkResponseOutput {
	return o
}

func (o ResourceNavigationLinkResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ResourceNavigationLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkResponseOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkResponseOutput) LinkedResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) *string { return v.LinkedResourceType }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ResourceNavigationLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceNavigationLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceNavigationLinkResponse)(nil)).Elem()
}

func (o ResourceNavigationLinkResponseArrayOutput) ToResourceNavigationLinkResponseArrayOutput() ResourceNavigationLinkResponseArrayOutput {
	return o
}

func (o ResourceNavigationLinkResponseArrayOutput) ToResourceNavigationLinkResponseArrayOutputWithContext(ctx context.Context) ResourceNavigationLinkResponseArrayOutput {
	return o
}

func (o ResourceNavigationLinkResponseArrayOutput) Index(i pulumi.IntInput) ResourceNavigationLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceNavigationLinkResponse {
		return vs[0].([]ResourceNavigationLinkResponse)[vs[1].(int)]
	}).(ResourceNavigationLinkResponseOutput)
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

type RouteFilterType struct {
	Id       *string                          `pulumi:"id"`
	Location string                           `pulumi:"location"`
	Peerings []ExpressRouteCircuitPeeringType `pulumi:"peerings"`
	Rules    []RouteFilterRuleType            `pulumi:"rules"`
	Tags     map[string]string                `pulumi:"tags"`
}





type RouteFilterTypeInput interface {
	pulumi.Input

	ToRouteFilterTypeOutput() RouteFilterTypeOutput
	ToRouteFilterTypeOutputWithContext(context.Context) RouteFilterTypeOutput
}

type RouteFilterTypeArgs struct {
	Id       pulumi.StringPtrInput                    `pulumi:"id"`
	Location pulumi.StringInput                       `pulumi:"location"`
	Peerings ExpressRouteCircuitPeeringTypeArrayInput `pulumi:"peerings"`
	Rules    RouteFilterRuleTypeArrayInput            `pulumi:"rules"`
	Tags     pulumi.StringMapInput                    `pulumi:"tags"`
}

func (RouteFilterTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterType)(nil)).Elem()
}

func (i RouteFilterTypeArgs) ToRouteFilterTypeOutput() RouteFilterTypeOutput {
	return i.ToRouteFilterTypeOutputWithContext(context.Background())
}

func (i RouteFilterTypeArgs) ToRouteFilterTypeOutputWithContext(ctx context.Context) RouteFilterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterTypeOutput)
}

func (i RouteFilterTypeArgs) ToRouteFilterTypePtrOutput() RouteFilterTypePtrOutput {
	return i.ToRouteFilterTypePtrOutputWithContext(context.Background())
}

func (i RouteFilterTypeArgs) ToRouteFilterTypePtrOutputWithContext(ctx context.Context) RouteFilterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterTypeOutput).ToRouteFilterTypePtrOutputWithContext(ctx)
}









type RouteFilterTypePtrInput interface {
	pulumi.Input

	ToRouteFilterTypePtrOutput() RouteFilterTypePtrOutput
	ToRouteFilterTypePtrOutputWithContext(context.Context) RouteFilterTypePtrOutput
}

type routeFilterTypePtrType RouteFilterTypeArgs

func RouteFilterTypePtr(v *RouteFilterTypeArgs) RouteFilterTypePtrInput {
	return (*routeFilterTypePtrType)(v)
}

func (*routeFilterTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilterType)(nil)).Elem()
}

func (i *routeFilterTypePtrType) ToRouteFilterTypePtrOutput() RouteFilterTypePtrOutput {
	return i.ToRouteFilterTypePtrOutputWithContext(context.Background())
}

func (i *routeFilterTypePtrType) ToRouteFilterTypePtrOutputWithContext(ctx context.Context) RouteFilterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterTypePtrOutput)
}

type RouteFilterTypeOutput struct{ *pulumi.OutputState }

func (RouteFilterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterType)(nil)).Elem()
}

func (o RouteFilterTypeOutput) ToRouteFilterTypeOutput() RouteFilterTypeOutput {
	return o
}

func (o RouteFilterTypeOutput) ToRouteFilterTypeOutputWithContext(ctx context.Context) RouteFilterTypeOutput {
	return o
}

func (o RouteFilterTypeOutput) ToRouteFilterTypePtrOutput() RouteFilterTypePtrOutput {
	return o.ToRouteFilterTypePtrOutputWithContext(context.Background())
}

func (o RouteFilterTypeOutput) ToRouteFilterTypePtrOutputWithContext(ctx context.Context) RouteFilterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RouteFilterType) *RouteFilterType {
		return &v
	}).(RouteFilterTypePtrOutput)
}

func (o RouteFilterTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteFilterTypeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterType) string { return v.Location }).(pulumi.StringOutput)
}

func (o RouteFilterTypeOutput) Peerings() ExpressRouteCircuitPeeringTypeArrayOutput {
	return o.ApplyT(func(v RouteFilterType) []ExpressRouteCircuitPeeringType { return v.Peerings }).(ExpressRouteCircuitPeeringTypeArrayOutput)
}

func (o RouteFilterTypeOutput) Rules() RouteFilterRuleTypeArrayOutput {
	return o.ApplyT(func(v RouteFilterType) []RouteFilterRuleType { return v.Rules }).(RouteFilterRuleTypeArrayOutput)
}

func (o RouteFilterTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RouteFilterType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type RouteFilterTypePtrOutput struct{ *pulumi.OutputState }

func (RouteFilterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilterType)(nil)).Elem()
}

func (o RouteFilterTypePtrOutput) ToRouteFilterTypePtrOutput() RouteFilterTypePtrOutput {
	return o
}

func (o RouteFilterTypePtrOutput) ToRouteFilterTypePtrOutputWithContext(ctx context.Context) RouteFilterTypePtrOutput {
	return o
}

func (o RouteFilterTypePtrOutput) Elem() RouteFilterTypeOutput {
	return o.ApplyT(func(v *RouteFilterType) RouteFilterType {
		if v != nil {
			return *v
		}
		var ret RouteFilterType
		return ret
	}).(RouteFilterTypeOutput)
}

func (o RouteFilterTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterType) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterTypePtrOutput) Peerings() ExpressRouteCircuitPeeringTypeArrayOutput {
	return o.ApplyT(func(v *RouteFilterType) []ExpressRouteCircuitPeeringType {
		if v == nil {
			return nil
		}
		return v.Peerings
	}).(ExpressRouteCircuitPeeringTypeArrayOutput)
}

func (o RouteFilterTypePtrOutput) Rules() RouteFilterRuleTypeArrayOutput {
	return o.ApplyT(func(v *RouteFilterType) []RouteFilterRuleType {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(RouteFilterRuleTypeArrayOutput)
}

func (o RouteFilterTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteFilterType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type RouteFilterResponse struct {
	Etag              string                               `pulumi:"etag"`
	Id                *string                              `pulumi:"id"`
	Location          string                               `pulumi:"location"`
	Name              string                               `pulumi:"name"`
	Peerings          []ExpressRouteCircuitPeeringResponse `pulumi:"peerings"`
	ProvisioningState string                               `pulumi:"provisioningState"`
	Rules             []RouteFilterRuleResponse            `pulumi:"rules"`
	Tags              map[string]string                    `pulumi:"tags"`
	Type              string                               `pulumi:"type"`
}





type RouteFilterResponseInput interface {
	pulumi.Input

	ToRouteFilterResponseOutput() RouteFilterResponseOutput
	ToRouteFilterResponseOutputWithContext(context.Context) RouteFilterResponseOutput
}

type RouteFilterResponseArgs struct {
	Etag              pulumi.StringInput                           `pulumi:"etag"`
	Id                pulumi.StringPtrInput                        `pulumi:"id"`
	Location          pulumi.StringInput                           `pulumi:"location"`
	Name              pulumi.StringInput                           `pulumi:"name"`
	Peerings          ExpressRouteCircuitPeeringResponseArrayInput `pulumi:"peerings"`
	ProvisioningState pulumi.StringInput                           `pulumi:"provisioningState"`
	Rules             RouteFilterRuleResponseArrayInput            `pulumi:"rules"`
	Tags              pulumi.StringMapInput                        `pulumi:"tags"`
	Type              pulumi.StringInput                           `pulumi:"type"`
}

func (RouteFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterResponse)(nil)).Elem()
}

func (i RouteFilterResponseArgs) ToRouteFilterResponseOutput() RouteFilterResponseOutput {
	return i.ToRouteFilterResponseOutputWithContext(context.Background())
}

func (i RouteFilterResponseArgs) ToRouteFilterResponseOutputWithContext(ctx context.Context) RouteFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterResponseOutput)
}

func (i RouteFilterResponseArgs) ToRouteFilterResponsePtrOutput() RouteFilterResponsePtrOutput {
	return i.ToRouteFilterResponsePtrOutputWithContext(context.Background())
}

func (i RouteFilterResponseArgs) ToRouteFilterResponsePtrOutputWithContext(ctx context.Context) RouteFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterResponseOutput).ToRouteFilterResponsePtrOutputWithContext(ctx)
}









type RouteFilterResponsePtrInput interface {
	pulumi.Input

	ToRouteFilterResponsePtrOutput() RouteFilterResponsePtrOutput
	ToRouteFilterResponsePtrOutputWithContext(context.Context) RouteFilterResponsePtrOutput
}

type routeFilterResponsePtrType RouteFilterResponseArgs

func RouteFilterResponsePtr(v *RouteFilterResponseArgs) RouteFilterResponsePtrInput {
	return (*routeFilterResponsePtrType)(v)
}

func (*routeFilterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilterResponse)(nil)).Elem()
}

func (i *routeFilterResponsePtrType) ToRouteFilterResponsePtrOutput() RouteFilterResponsePtrOutput {
	return i.ToRouteFilterResponsePtrOutputWithContext(context.Background())
}

func (i *routeFilterResponsePtrType) ToRouteFilterResponsePtrOutputWithContext(ctx context.Context) RouteFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterResponsePtrOutput)
}

type RouteFilterResponseOutput struct{ *pulumi.OutputState }

func (RouteFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterResponse)(nil)).Elem()
}

func (o RouteFilterResponseOutput) ToRouteFilterResponseOutput() RouteFilterResponseOutput {
	return o
}

func (o RouteFilterResponseOutput) ToRouteFilterResponseOutputWithContext(ctx context.Context) RouteFilterResponseOutput {
	return o
}

func (o RouteFilterResponseOutput) ToRouteFilterResponsePtrOutput() RouteFilterResponsePtrOutput {
	return o.ToRouteFilterResponsePtrOutputWithContext(context.Background())
}

func (o RouteFilterResponseOutput) ToRouteFilterResponsePtrOutputWithContext(ctx context.Context) RouteFilterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RouteFilterResponse) *RouteFilterResponse {
		return &v
	}).(RouteFilterResponsePtrOutput)
}

func (o RouteFilterResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o RouteFilterResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o RouteFilterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RouteFilterResponseOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v RouteFilterResponse) []ExpressRouteCircuitPeeringResponse { return v.Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

func (o RouteFilterResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RouteFilterResponseOutput) Rules() RouteFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v RouteFilterResponse) []RouteFilterRuleResponse { return v.Rules }).(RouteFilterRuleResponseArrayOutput)
}

func (o RouteFilterResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RouteFilterResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RouteFilterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RouteFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (RouteFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilterResponse)(nil)).Elem()
}

func (o RouteFilterResponsePtrOutput) ToRouteFilterResponsePtrOutput() RouteFilterResponsePtrOutput {
	return o
}

func (o RouteFilterResponsePtrOutput) ToRouteFilterResponsePtrOutputWithContext(ctx context.Context) RouteFilterResponsePtrOutput {
	return o
}

func (o RouteFilterResponsePtrOutput) Elem() RouteFilterResponseOutput {
	return o.ApplyT(func(v *RouteFilterResponse) RouteFilterResponse {
		if v != nil {
			return *v
		}
		var ret RouteFilterResponse
		return ret
	}).(RouteFilterResponseOutput)
}

func (o RouteFilterResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponsePtrOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v *RouteFilterResponse) []ExpressRouteCircuitPeeringResponse {
		if v == nil {
			return nil
		}
		return v.Peerings
	}).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

func (o RouteFilterResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponsePtrOutput) Rules() RouteFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v *RouteFilterResponse) []RouteFilterRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(RouteFilterRuleResponseArrayOutput)
}

func (o RouteFilterResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteFilterResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o RouteFilterResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type RouteFilterRuleType struct {
	Access              string   `pulumi:"access"`
	Communities         []string `pulumi:"communities"`
	Id                  *string  `pulumi:"id"`
	Location            *string  `pulumi:"location"`
	Name                *string  `pulumi:"name"`
	RouteFilterRuleType string   `pulumi:"routeFilterRuleType"`
}





type RouteFilterRuleTypeInput interface {
	pulumi.Input

	ToRouteFilterRuleTypeOutput() RouteFilterRuleTypeOutput
	ToRouteFilterRuleTypeOutputWithContext(context.Context) RouteFilterRuleTypeOutput
}

type RouteFilterRuleTypeArgs struct {
	Access              pulumi.StringInput      `pulumi:"access"`
	Communities         pulumi.StringArrayInput `pulumi:"communities"`
	Id                  pulumi.StringPtrInput   `pulumi:"id"`
	Location            pulumi.StringPtrInput   `pulumi:"location"`
	Name                pulumi.StringPtrInput   `pulumi:"name"`
	RouteFilterRuleType pulumi.StringInput      `pulumi:"routeFilterRuleType"`
}

func (RouteFilterRuleTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterRuleType)(nil)).Elem()
}

func (i RouteFilterRuleTypeArgs) ToRouteFilterRuleTypeOutput() RouteFilterRuleTypeOutput {
	return i.ToRouteFilterRuleTypeOutputWithContext(context.Background())
}

func (i RouteFilterRuleTypeArgs) ToRouteFilterRuleTypeOutputWithContext(ctx context.Context) RouteFilterRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterRuleTypeOutput)
}





type RouteFilterRuleTypeArrayInput interface {
	pulumi.Input

	ToRouteFilterRuleTypeArrayOutput() RouteFilterRuleTypeArrayOutput
	ToRouteFilterRuleTypeArrayOutputWithContext(context.Context) RouteFilterRuleTypeArrayOutput
}

type RouteFilterRuleTypeArray []RouteFilterRuleTypeInput

func (RouteFilterRuleTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteFilterRuleType)(nil)).Elem()
}

func (i RouteFilterRuleTypeArray) ToRouteFilterRuleTypeArrayOutput() RouteFilterRuleTypeArrayOutput {
	return i.ToRouteFilterRuleTypeArrayOutputWithContext(context.Background())
}

func (i RouteFilterRuleTypeArray) ToRouteFilterRuleTypeArrayOutputWithContext(ctx context.Context) RouteFilterRuleTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterRuleTypeArrayOutput)
}

type RouteFilterRuleTypeOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterRuleType)(nil)).Elem()
}

func (o RouteFilterRuleTypeOutput) ToRouteFilterRuleTypeOutput() RouteFilterRuleTypeOutput {
	return o
}

func (o RouteFilterRuleTypeOutput) ToRouteFilterRuleTypeOutputWithContext(ctx context.Context) RouteFilterRuleTypeOutput {
	return o
}

func (o RouteFilterRuleTypeOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleType) string { return v.Access }).(pulumi.StringOutput)
}

func (o RouteFilterRuleTypeOutput) Communities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RouteFilterRuleType) []string { return v.Communities }).(pulumi.StringArrayOutput)
}

func (o RouteFilterRuleTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterRuleType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterRuleType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterRuleType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleTypeOutput) RouteFilterRuleType() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleType) string { return v.RouteFilterRuleType }).(pulumi.StringOutput)
}

type RouteFilterRuleTypeArrayOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteFilterRuleType)(nil)).Elem()
}

func (o RouteFilterRuleTypeArrayOutput) ToRouteFilterRuleTypeArrayOutput() RouteFilterRuleTypeArrayOutput {
	return o
}

func (o RouteFilterRuleTypeArrayOutput) ToRouteFilterRuleTypeArrayOutputWithContext(ctx context.Context) RouteFilterRuleTypeArrayOutput {
	return o
}

func (o RouteFilterRuleTypeArrayOutput) Index(i pulumi.IntInput) RouteFilterRuleTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteFilterRuleType {
		return vs[0].([]RouteFilterRuleType)[vs[1].(int)]
	}).(RouteFilterRuleTypeOutput)
}

type RouteFilterRuleResponse struct {
	Access              string   `pulumi:"access"`
	Communities         []string `pulumi:"communities"`
	Etag                string   `pulumi:"etag"`
	Id                  *string  `pulumi:"id"`
	Location            *string  `pulumi:"location"`
	Name                *string  `pulumi:"name"`
	ProvisioningState   string   `pulumi:"provisioningState"`
	RouteFilterRuleType string   `pulumi:"routeFilterRuleType"`
}





type RouteFilterRuleResponseInput interface {
	pulumi.Input

	ToRouteFilterRuleResponseOutput() RouteFilterRuleResponseOutput
	ToRouteFilterRuleResponseOutputWithContext(context.Context) RouteFilterRuleResponseOutput
}

type RouteFilterRuleResponseArgs struct {
	Access              pulumi.StringInput      `pulumi:"access"`
	Communities         pulumi.StringArrayInput `pulumi:"communities"`
	Etag                pulumi.StringInput      `pulumi:"etag"`
	Id                  pulumi.StringPtrInput   `pulumi:"id"`
	Location            pulumi.StringPtrInput   `pulumi:"location"`
	Name                pulumi.StringPtrInput   `pulumi:"name"`
	ProvisioningState   pulumi.StringInput      `pulumi:"provisioningState"`
	RouteFilterRuleType pulumi.StringInput      `pulumi:"routeFilterRuleType"`
}

func (RouteFilterRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterRuleResponse)(nil)).Elem()
}

func (i RouteFilterRuleResponseArgs) ToRouteFilterRuleResponseOutput() RouteFilterRuleResponseOutput {
	return i.ToRouteFilterRuleResponseOutputWithContext(context.Background())
}

func (i RouteFilterRuleResponseArgs) ToRouteFilterRuleResponseOutputWithContext(ctx context.Context) RouteFilterRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterRuleResponseOutput)
}





type RouteFilterRuleResponseArrayInput interface {
	pulumi.Input

	ToRouteFilterRuleResponseArrayOutput() RouteFilterRuleResponseArrayOutput
	ToRouteFilterRuleResponseArrayOutputWithContext(context.Context) RouteFilterRuleResponseArrayOutput
}

type RouteFilterRuleResponseArray []RouteFilterRuleResponseInput

func (RouteFilterRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteFilterRuleResponse)(nil)).Elem()
}

func (i RouteFilterRuleResponseArray) ToRouteFilterRuleResponseArrayOutput() RouteFilterRuleResponseArrayOutput {
	return i.ToRouteFilterRuleResponseArrayOutputWithContext(context.Background())
}

func (i RouteFilterRuleResponseArray) ToRouteFilterRuleResponseArrayOutputWithContext(ctx context.Context) RouteFilterRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterRuleResponseArrayOutput)
}

type RouteFilterRuleResponseOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterRuleResponse)(nil)).Elem()
}

func (o RouteFilterRuleResponseOutput) ToRouteFilterRuleResponseOutput() RouteFilterRuleResponseOutput {
	return o
}

func (o RouteFilterRuleResponseOutput) ToRouteFilterRuleResponseOutputWithContext(ctx context.Context) RouteFilterRuleResponseOutput {
	return o
}

func (o RouteFilterRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o RouteFilterRuleResponseOutput) Communities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) []string { return v.Communities }).(pulumi.StringArrayOutput)
}

func (o RouteFilterRuleResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o RouteFilterRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RouteFilterRuleResponseOutput) RouteFilterRuleType() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) string { return v.RouteFilterRuleType }).(pulumi.StringOutput)
}

type RouteFilterRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteFilterRuleResponse)(nil)).Elem()
}

func (o RouteFilterRuleResponseArrayOutput) ToRouteFilterRuleResponseArrayOutput() RouteFilterRuleResponseArrayOutput {
	return o
}

func (o RouteFilterRuleResponseArrayOutput) ToRouteFilterRuleResponseArrayOutputWithContext(ctx context.Context) RouteFilterRuleResponseArrayOutput {
	return o
}

func (o RouteFilterRuleResponseArrayOutput) Index(i pulumi.IntInput) RouteFilterRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteFilterRuleResponse {
		return vs[0].([]RouteFilterRuleResponse)[vs[1].(int)]
	}).(RouteFilterRuleResponseOutput)
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

type RouteTableType struct {
	DisableBgpRoutePropagation *bool             `pulumi:"disableBgpRoutePropagation"`
	Etag                       *string           `pulumi:"etag"`
	Id                         *string           `pulumi:"id"`
	Location                   *string           `pulumi:"location"`
	ProvisioningState          *string           `pulumi:"provisioningState"`
	Routes                     []RouteType       `pulumi:"routes"`
	Tags                       map[string]string `pulumi:"tags"`
}





type RouteTableTypeInput interface {
	pulumi.Input

	ToRouteTableTypeOutput() RouteTableTypeOutput
	ToRouteTableTypeOutputWithContext(context.Context) RouteTableTypeOutput
}

type RouteTableTypeArgs struct {
	DisableBgpRoutePropagation pulumi.BoolPtrInput   `pulumi:"disableBgpRoutePropagation"`
	Etag                       pulumi.StringPtrInput `pulumi:"etag"`
	Id                         pulumi.StringPtrInput `pulumi:"id"`
	Location                   pulumi.StringPtrInput `pulumi:"location"`
	ProvisioningState          pulumi.StringPtrInput `pulumi:"provisioningState"`
	Routes                     RouteTypeArrayInput   `pulumi:"routes"`
	Tags                       pulumi.StringMapInput `pulumi:"tags"`
}

func (RouteTableTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableType)(nil)).Elem()
}

func (i RouteTableTypeArgs) ToRouteTableTypeOutput() RouteTableTypeOutput {
	return i.ToRouteTableTypeOutputWithContext(context.Background())
}

func (i RouteTableTypeArgs) ToRouteTableTypeOutputWithContext(ctx context.Context) RouteTableTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableTypeOutput)
}

func (i RouteTableTypeArgs) ToRouteTableTypePtrOutput() RouteTableTypePtrOutput {
	return i.ToRouteTableTypePtrOutputWithContext(context.Background())
}

func (i RouteTableTypeArgs) ToRouteTableTypePtrOutputWithContext(ctx context.Context) RouteTableTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableTypeOutput).ToRouteTableTypePtrOutputWithContext(ctx)
}









type RouteTableTypePtrInput interface {
	pulumi.Input

	ToRouteTableTypePtrOutput() RouteTableTypePtrOutput
	ToRouteTableTypePtrOutputWithContext(context.Context) RouteTableTypePtrOutput
}

type routeTableTypePtrType RouteTableTypeArgs

func RouteTableTypePtr(v *RouteTableTypeArgs) RouteTableTypePtrInput {
	return (*routeTableTypePtrType)(v)
}

func (*routeTableTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTableType)(nil)).Elem()
}

func (i *routeTableTypePtrType) ToRouteTableTypePtrOutput() RouteTableTypePtrOutput {
	return i.ToRouteTableTypePtrOutputWithContext(context.Background())
}

func (i *routeTableTypePtrType) ToRouteTableTypePtrOutputWithContext(ctx context.Context) RouteTableTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableTypePtrOutput)
}

type RouteTableTypeOutput struct{ *pulumi.OutputState }

func (RouteTableTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableType)(nil)).Elem()
}

func (o RouteTableTypeOutput) ToRouteTableTypeOutput() RouteTableTypeOutput {
	return o
}

func (o RouteTableTypeOutput) ToRouteTableTypeOutputWithContext(ctx context.Context) RouteTableTypeOutput {
	return o
}

func (o RouteTableTypeOutput) ToRouteTableTypePtrOutput() RouteTableTypePtrOutput {
	return o.ToRouteTableTypePtrOutputWithContext(context.Background())
}

func (o RouteTableTypeOutput) ToRouteTableTypePtrOutputWithContext(ctx context.Context) RouteTableTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RouteTableType) *RouteTableType {
		return &v
	}).(RouteTableTypePtrOutput)
}

func (o RouteTableTypeOutput) DisableBgpRoutePropagation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RouteTableType) *bool { return v.DisableBgpRoutePropagation }).(pulumi.BoolPtrOutput)
}

func (o RouteTableTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RouteTableTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteTableTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RouteTableTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o RouteTableTypeOutput) Routes() RouteTypeArrayOutput {
	return o.ApplyT(func(v RouteTableType) []RouteType { return v.Routes }).(RouteTypeArrayOutput)
}

func (o RouteTableTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RouteTableType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type RouteTableTypePtrOutput struct{ *pulumi.OutputState }

func (RouteTableTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTableType)(nil)).Elem()
}

func (o RouteTableTypePtrOutput) ToRouteTableTypePtrOutput() RouteTableTypePtrOutput {
	return o
}

func (o RouteTableTypePtrOutput) ToRouteTableTypePtrOutputWithContext(ctx context.Context) RouteTableTypePtrOutput {
	return o
}

func (o RouteTableTypePtrOutput) Elem() RouteTableTypeOutput {
	return o.ApplyT(func(v *RouteTableType) RouteTableType {
		if v != nil {
			return *v
		}
		var ret RouteTableType
		return ret
	}).(RouteTableTypeOutput)
}

func (o RouteTableTypePtrOutput) DisableBgpRoutePropagation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RouteTableType) *bool {
		if v == nil {
			return nil
		}
		return v.DisableBgpRoutePropagation
	}).(pulumi.BoolPtrOutput)
}

func (o RouteTableTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableTypePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableType) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableTypePtrOutput) Routes() RouteTypeArrayOutput {
	return o.ApplyT(func(v *RouteTableType) []RouteType {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(RouteTypeArrayOutput)
}

func (o RouteTableTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteTableType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type RouteTableResponse struct {
	DisableBgpRoutePropagation *bool             `pulumi:"disableBgpRoutePropagation"`
	Etag                       *string           `pulumi:"etag"`
	Id                         *string           `pulumi:"id"`
	Location                   *string           `pulumi:"location"`
	Name                       string            `pulumi:"name"`
	ProvisioningState          *string           `pulumi:"provisioningState"`
	Routes                     []RouteResponse   `pulumi:"routes"`
	Subnets                    []SubnetResponse  `pulumi:"subnets"`
	Tags                       map[string]string `pulumi:"tags"`
	Type                       string            `pulumi:"type"`
}





type RouteTableResponseInput interface {
	pulumi.Input

	ToRouteTableResponseOutput() RouteTableResponseOutput
	ToRouteTableResponseOutputWithContext(context.Context) RouteTableResponseOutput
}

type RouteTableResponseArgs struct {
	DisableBgpRoutePropagation pulumi.BoolPtrInput      `pulumi:"disableBgpRoutePropagation"`
	Etag                       pulumi.StringPtrInput    `pulumi:"etag"`
	Id                         pulumi.StringPtrInput    `pulumi:"id"`
	Location                   pulumi.StringPtrInput    `pulumi:"location"`
	Name                       pulumi.StringInput       `pulumi:"name"`
	ProvisioningState          pulumi.StringPtrInput    `pulumi:"provisioningState"`
	Routes                     RouteResponseArrayInput  `pulumi:"routes"`
	Subnets                    SubnetResponseArrayInput `pulumi:"subnets"`
	Tags                       pulumi.StringMapInput    `pulumi:"tags"`
	Type                       pulumi.StringInput       `pulumi:"type"`
}

func (RouteTableResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableResponse)(nil)).Elem()
}

func (i RouteTableResponseArgs) ToRouteTableResponseOutput() RouteTableResponseOutput {
	return i.ToRouteTableResponseOutputWithContext(context.Background())
}

func (i RouteTableResponseArgs) ToRouteTableResponseOutputWithContext(ctx context.Context) RouteTableResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableResponseOutput)
}

func (i RouteTableResponseArgs) ToRouteTableResponsePtrOutput() RouteTableResponsePtrOutput {
	return i.ToRouteTableResponsePtrOutputWithContext(context.Background())
}

func (i RouteTableResponseArgs) ToRouteTableResponsePtrOutputWithContext(ctx context.Context) RouteTableResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableResponseOutput).ToRouteTableResponsePtrOutputWithContext(ctx)
}









type RouteTableResponsePtrInput interface {
	pulumi.Input

	ToRouteTableResponsePtrOutput() RouteTableResponsePtrOutput
	ToRouteTableResponsePtrOutputWithContext(context.Context) RouteTableResponsePtrOutput
}

type routeTableResponsePtrType RouteTableResponseArgs

func RouteTableResponsePtr(v *RouteTableResponseArgs) RouteTableResponsePtrInput {
	return (*routeTableResponsePtrType)(v)
}

func (*routeTableResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTableResponse)(nil)).Elem()
}

func (i *routeTableResponsePtrType) ToRouteTableResponsePtrOutput() RouteTableResponsePtrOutput {
	return i.ToRouteTableResponsePtrOutputWithContext(context.Background())
}

func (i *routeTableResponsePtrType) ToRouteTableResponsePtrOutputWithContext(ctx context.Context) RouteTableResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableResponsePtrOutput)
}

type RouteTableResponseOutput struct{ *pulumi.OutputState }

func (RouteTableResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableResponse)(nil)).Elem()
}

func (o RouteTableResponseOutput) ToRouteTableResponseOutput() RouteTableResponseOutput {
	return o
}

func (o RouteTableResponseOutput) ToRouteTableResponseOutputWithContext(ctx context.Context) RouteTableResponseOutput {
	return o
}

func (o RouteTableResponseOutput) ToRouteTableResponsePtrOutput() RouteTableResponsePtrOutput {
	return o.ToRouteTableResponsePtrOutputWithContext(context.Background())
}

func (o RouteTableResponseOutput) ToRouteTableResponsePtrOutputWithContext(ctx context.Context) RouteTableResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RouteTableResponse) *RouteTableResponse {
		return &v
	}).(RouteTableResponsePtrOutput)
}

func (o RouteTableResponseOutput) DisableBgpRoutePropagation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RouteTableResponse) *bool { return v.DisableBgpRoutePropagation }).(pulumi.BoolPtrOutput)
}

func (o RouteTableResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RouteTableResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteTableResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RouteTableResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RouteTableResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RouteTableResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o RouteTableResponseOutput) Routes() RouteResponseArrayOutput {
	return o.ApplyT(func(v RouteTableResponse) []RouteResponse { return v.Routes }).(RouteResponseArrayOutput)
}

func (o RouteTableResponseOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v RouteTableResponse) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

func (o RouteTableResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RouteTableResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RouteTableResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RouteTableResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RouteTableResponsePtrOutput struct{ *pulumi.OutputState }

func (RouteTableResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTableResponse)(nil)).Elem()
}

func (o RouteTableResponsePtrOutput) ToRouteTableResponsePtrOutput() RouteTableResponsePtrOutput {
	return o
}

func (o RouteTableResponsePtrOutput) ToRouteTableResponsePtrOutputWithContext(ctx context.Context) RouteTableResponsePtrOutput {
	return o
}

func (o RouteTableResponsePtrOutput) Elem() RouteTableResponseOutput {
	return o.ApplyT(func(v *RouteTableResponse) RouteTableResponse {
		if v != nil {
			return *v
		}
		var ret RouteTableResponse
		return ret
	}).(RouteTableResponseOutput)
}

func (o RouteTableResponsePtrOutput) DisableBgpRoutePropagation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableBgpRoutePropagation
	}).(pulumi.BoolPtrOutput)
}

func (o RouteTableResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableResponsePtrOutput) Routes() RouteResponseArrayOutput {
	return o.ApplyT(func(v *RouteTableResponse) []RouteResponse {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(RouteResponseArrayOutput)
}

func (o RouteTableResponsePtrOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *RouteTableResponse) []SubnetResponse {
		if v == nil {
			return nil
		}
		return v.Subnets
	}).(SubnetResponseArrayOutput)
}

func (o RouteTableResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteTableResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o RouteTableResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type SecurityRuleType struct {
	Access                               string                         `pulumi:"access"`
	Description                          *string                        `pulumi:"description"`
	DestinationAddressPrefix             *string                        `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes           []string                       `pulumi:"destinationAddressPrefixes"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupType `pulumi:"destinationApplicationSecurityGroups"`
	DestinationPortRange                 *string                        `pulumi:"destinationPortRange"`
	DestinationPortRanges                []string                       `pulumi:"destinationPortRanges"`
	Direction                            string                         `pulumi:"direction"`
	Etag                                 *string                        `pulumi:"etag"`
	Id                                   *string                        `pulumi:"id"`
	Name                                 *string                        `pulumi:"name"`
	Priority                             *int                           `pulumi:"priority"`
	Protocol                             string                         `pulumi:"protocol"`
	ProvisioningState                    *string                        `pulumi:"provisioningState"`
	SourceAddressPrefix                  *string                        `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes                []string                       `pulumi:"sourceAddressPrefixes"`
	SourceApplicationSecurityGroups      []ApplicationSecurityGroupType `pulumi:"sourceApplicationSecurityGroups"`
	SourcePortRange                      *string                        `pulumi:"sourcePortRange"`
	SourcePortRanges                     []string                       `pulumi:"sourcePortRanges"`
}





type SecurityRuleTypeInput interface {
	pulumi.Input

	ToSecurityRuleTypeOutput() SecurityRuleTypeOutput
	ToSecurityRuleTypeOutputWithContext(context.Context) SecurityRuleTypeOutput
}

type SecurityRuleTypeArgs struct {
	Access                               pulumi.StringInput                     `pulumi:"access"`
	Description                          pulumi.StringPtrInput                  `pulumi:"description"`
	DestinationAddressPrefix             pulumi.StringPtrInput                  `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes           pulumi.StringArrayInput                `pulumi:"destinationAddressPrefixes"`
	DestinationApplicationSecurityGroups ApplicationSecurityGroupTypeArrayInput `pulumi:"destinationApplicationSecurityGroups"`
	DestinationPortRange                 pulumi.StringPtrInput                  `pulumi:"destinationPortRange"`
	DestinationPortRanges                pulumi.StringArrayInput                `pulumi:"destinationPortRanges"`
	Direction                            pulumi.StringInput                     `pulumi:"direction"`
	Etag                                 pulumi.StringPtrInput                  `pulumi:"etag"`
	Id                                   pulumi.StringPtrInput                  `pulumi:"id"`
	Name                                 pulumi.StringPtrInput                  `pulumi:"name"`
	Priority                             pulumi.IntPtrInput                     `pulumi:"priority"`
	Protocol                             pulumi.StringInput                     `pulumi:"protocol"`
	ProvisioningState                    pulumi.StringPtrInput                  `pulumi:"provisioningState"`
	SourceAddressPrefix                  pulumi.StringPtrInput                  `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes                pulumi.StringArrayInput                `pulumi:"sourceAddressPrefixes"`
	SourceApplicationSecurityGroups      ApplicationSecurityGroupTypeArrayInput `pulumi:"sourceApplicationSecurityGroups"`
	SourcePortRange                      pulumi.StringPtrInput                  `pulumi:"sourcePortRange"`
	SourcePortRanges                     pulumi.StringArrayInput                `pulumi:"sourcePortRanges"`
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

func (o SecurityRuleTypeOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []string { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleTypeOutput) DestinationApplicationSecurityGroups() ApplicationSecurityGroupTypeArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []ApplicationSecurityGroupType { return v.DestinationApplicationSecurityGroups }).(ApplicationSecurityGroupTypeArrayOutput)
}

func (o SecurityRuleTypeOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
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

func (o SecurityRuleTypeOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []string { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleTypeOutput) SourceApplicationSecurityGroups() ApplicationSecurityGroupTypeArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []ApplicationSecurityGroupType { return v.SourceApplicationSecurityGroups }).(ApplicationSecurityGroupTypeArrayOutput)
}

func (o SecurityRuleTypeOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
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
	Access                               string                             `pulumi:"access"`
	Description                          *string                            `pulumi:"description"`
	DestinationAddressPrefix             *string                            `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes           []string                           `pulumi:"destinationAddressPrefixes"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupResponse `pulumi:"destinationApplicationSecurityGroups"`
	DestinationPortRange                 *string                            `pulumi:"destinationPortRange"`
	DestinationPortRanges                []string                           `pulumi:"destinationPortRanges"`
	Direction                            string                             `pulumi:"direction"`
	Etag                                 *string                            `pulumi:"etag"`
	Id                                   *string                            `pulumi:"id"`
	Name                                 *string                            `pulumi:"name"`
	Priority                             *int                               `pulumi:"priority"`
	Protocol                             string                             `pulumi:"protocol"`
	ProvisioningState                    *string                            `pulumi:"provisioningState"`
	SourceAddressPrefix                  *string                            `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes                []string                           `pulumi:"sourceAddressPrefixes"`
	SourceApplicationSecurityGroups      []ApplicationSecurityGroupResponse `pulumi:"sourceApplicationSecurityGroups"`
	SourcePortRange                      *string                            `pulumi:"sourcePortRange"`
	SourcePortRanges                     []string                           `pulumi:"sourcePortRanges"`
}





type SecurityRuleResponseInput interface {
	pulumi.Input

	ToSecurityRuleResponseOutput() SecurityRuleResponseOutput
	ToSecurityRuleResponseOutputWithContext(context.Context) SecurityRuleResponseOutput
}

type SecurityRuleResponseArgs struct {
	Access                               pulumi.StringInput                         `pulumi:"access"`
	Description                          pulumi.StringPtrInput                      `pulumi:"description"`
	DestinationAddressPrefix             pulumi.StringPtrInput                      `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes           pulumi.StringArrayInput                    `pulumi:"destinationAddressPrefixes"`
	DestinationApplicationSecurityGroups ApplicationSecurityGroupResponseArrayInput `pulumi:"destinationApplicationSecurityGroups"`
	DestinationPortRange                 pulumi.StringPtrInput                      `pulumi:"destinationPortRange"`
	DestinationPortRanges                pulumi.StringArrayInput                    `pulumi:"destinationPortRanges"`
	Direction                            pulumi.StringInput                         `pulumi:"direction"`
	Etag                                 pulumi.StringPtrInput                      `pulumi:"etag"`
	Id                                   pulumi.StringPtrInput                      `pulumi:"id"`
	Name                                 pulumi.StringPtrInput                      `pulumi:"name"`
	Priority                             pulumi.IntPtrInput                         `pulumi:"priority"`
	Protocol                             pulumi.StringInput                         `pulumi:"protocol"`
	ProvisioningState                    pulumi.StringPtrInput                      `pulumi:"provisioningState"`
	SourceAddressPrefix                  pulumi.StringPtrInput                      `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes                pulumi.StringArrayInput                    `pulumi:"sourceAddressPrefixes"`
	SourceApplicationSecurityGroups      ApplicationSecurityGroupResponseArrayInput `pulumi:"sourceApplicationSecurityGroups"`
	SourcePortRange                      pulumi.StringPtrInput                      `pulumi:"sourcePortRange"`
	SourcePortRanges                     pulumi.StringArrayInput                    `pulumi:"sourcePortRanges"`
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

func (o SecurityRuleResponseOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []string { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleResponseOutput) DestinationApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []ApplicationSecurityGroupResponse {
		return v.DestinationApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

func (o SecurityRuleResponseOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
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

func (o SecurityRuleResponseOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []string { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleResponseOutput) SourceApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []ApplicationSecurityGroupResponse {
		return v.SourceApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

func (o SecurityRuleResponseOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
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

type ServiceEndpointPropertiesFormat struct {
	Locations         []string `pulumi:"locations"`
	ProvisioningState *string  `pulumi:"provisioningState"`
	Service           *string  `pulumi:"service"`
}





type ServiceEndpointPropertiesFormatInput interface {
	pulumi.Input

	ToServiceEndpointPropertiesFormatOutput() ServiceEndpointPropertiesFormatOutput
	ToServiceEndpointPropertiesFormatOutputWithContext(context.Context) ServiceEndpointPropertiesFormatOutput
}

type ServiceEndpointPropertiesFormatArgs struct {
	Locations         pulumi.StringArrayInput `pulumi:"locations"`
	ProvisioningState pulumi.StringPtrInput   `pulumi:"provisioningState"`
	Service           pulumi.StringPtrInput   `pulumi:"service"`
}

func (ServiceEndpointPropertiesFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPropertiesFormat)(nil)).Elem()
}

func (i ServiceEndpointPropertiesFormatArgs) ToServiceEndpointPropertiesFormatOutput() ServiceEndpointPropertiesFormatOutput {
	return i.ToServiceEndpointPropertiesFormatOutputWithContext(context.Background())
}

func (i ServiceEndpointPropertiesFormatArgs) ToServiceEndpointPropertiesFormatOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPropertiesFormatOutput)
}





type ServiceEndpointPropertiesFormatArrayInput interface {
	pulumi.Input

	ToServiceEndpointPropertiesFormatArrayOutput() ServiceEndpointPropertiesFormatArrayOutput
	ToServiceEndpointPropertiesFormatArrayOutputWithContext(context.Context) ServiceEndpointPropertiesFormatArrayOutput
}

type ServiceEndpointPropertiesFormatArray []ServiceEndpointPropertiesFormatInput

func (ServiceEndpointPropertiesFormatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPropertiesFormat)(nil)).Elem()
}

func (i ServiceEndpointPropertiesFormatArray) ToServiceEndpointPropertiesFormatArrayOutput() ServiceEndpointPropertiesFormatArrayOutput {
	return i.ToServiceEndpointPropertiesFormatArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointPropertiesFormatArray) ToServiceEndpointPropertiesFormatArrayOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPropertiesFormatArrayOutput)
}

type ServiceEndpointPropertiesFormatOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPropertiesFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPropertiesFormat)(nil)).Elem()
}

func (o ServiceEndpointPropertiesFormatOutput) ToServiceEndpointPropertiesFormatOutput() ServiceEndpointPropertiesFormatOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatOutput) ToServiceEndpointPropertiesFormatOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormat) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o ServiceEndpointPropertiesFormatOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormat) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPropertiesFormatOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormat) *string { return v.Service }).(pulumi.StringPtrOutput)
}

type ServiceEndpointPropertiesFormatArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPropertiesFormatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPropertiesFormat)(nil)).Elem()
}

func (o ServiceEndpointPropertiesFormatArrayOutput) ToServiceEndpointPropertiesFormatArrayOutput() ServiceEndpointPropertiesFormatArrayOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatArrayOutput) ToServiceEndpointPropertiesFormatArrayOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatArrayOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatArrayOutput) Index(i pulumi.IntInput) ServiceEndpointPropertiesFormatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointPropertiesFormat {
		return vs[0].([]ServiceEndpointPropertiesFormat)[vs[1].(int)]
	}).(ServiceEndpointPropertiesFormatOutput)
}

type ServiceEndpointPropertiesFormatResponse struct {
	Locations         []string `pulumi:"locations"`
	ProvisioningState *string  `pulumi:"provisioningState"`
	Service           *string  `pulumi:"service"`
}





type ServiceEndpointPropertiesFormatResponseInput interface {
	pulumi.Input

	ToServiceEndpointPropertiesFormatResponseOutput() ServiceEndpointPropertiesFormatResponseOutput
	ToServiceEndpointPropertiesFormatResponseOutputWithContext(context.Context) ServiceEndpointPropertiesFormatResponseOutput
}

type ServiceEndpointPropertiesFormatResponseArgs struct {
	Locations         pulumi.StringArrayInput `pulumi:"locations"`
	ProvisioningState pulumi.StringPtrInput   `pulumi:"provisioningState"`
	Service           pulumi.StringPtrInput   `pulumi:"service"`
}

func (ServiceEndpointPropertiesFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPropertiesFormatResponse)(nil)).Elem()
}

func (i ServiceEndpointPropertiesFormatResponseArgs) ToServiceEndpointPropertiesFormatResponseOutput() ServiceEndpointPropertiesFormatResponseOutput {
	return i.ToServiceEndpointPropertiesFormatResponseOutputWithContext(context.Background())
}

func (i ServiceEndpointPropertiesFormatResponseArgs) ToServiceEndpointPropertiesFormatResponseOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPropertiesFormatResponseOutput)
}





type ServiceEndpointPropertiesFormatResponseArrayInput interface {
	pulumi.Input

	ToServiceEndpointPropertiesFormatResponseArrayOutput() ServiceEndpointPropertiesFormatResponseArrayOutput
	ToServiceEndpointPropertiesFormatResponseArrayOutputWithContext(context.Context) ServiceEndpointPropertiesFormatResponseArrayOutput
}

type ServiceEndpointPropertiesFormatResponseArray []ServiceEndpointPropertiesFormatResponseInput

func (ServiceEndpointPropertiesFormatResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPropertiesFormatResponse)(nil)).Elem()
}

func (i ServiceEndpointPropertiesFormatResponseArray) ToServiceEndpointPropertiesFormatResponseArrayOutput() ServiceEndpointPropertiesFormatResponseArrayOutput {
	return i.ToServiceEndpointPropertiesFormatResponseArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointPropertiesFormatResponseArray) ToServiceEndpointPropertiesFormatResponseArrayOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPropertiesFormatResponseArrayOutput)
}

type ServiceEndpointPropertiesFormatResponseOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPropertiesFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPropertiesFormatResponse)(nil)).Elem()
}

func (o ServiceEndpointPropertiesFormatResponseOutput) ToServiceEndpointPropertiesFormatResponseOutput() ServiceEndpointPropertiesFormatResponseOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatResponseOutput) ToServiceEndpointPropertiesFormatResponseOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatResponseOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormatResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o ServiceEndpointPropertiesFormatResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormatResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPropertiesFormatResponseOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormatResponse) *string { return v.Service }).(pulumi.StringPtrOutput)
}

type ServiceEndpointPropertiesFormatResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPropertiesFormatResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPropertiesFormatResponse)(nil)).Elem()
}

func (o ServiceEndpointPropertiesFormatResponseArrayOutput) ToServiceEndpointPropertiesFormatResponseArrayOutput() ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatResponseArrayOutput) ToServiceEndpointPropertiesFormatResponseArrayOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatResponseArrayOutput) Index(i pulumi.IntInput) ServiceEndpointPropertiesFormatResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointPropertiesFormatResponse {
		return vs[0].([]ServiceEndpointPropertiesFormatResponse)[vs[1].(int)]
	}).(ServiceEndpointPropertiesFormatResponseOutput)
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
	AddressPrefix           *string                           `pulumi:"addressPrefix"`
	Etag                    *string                           `pulumi:"etag"`
	Id                      *string                           `pulumi:"id"`
	Name                    *string                           `pulumi:"name"`
	NetworkSecurityGroup    *NetworkSecurityGroupType         `pulumi:"networkSecurityGroup"`
	ProvisioningState       *string                           `pulumi:"provisioningState"`
	ResourceNavigationLinks []ResourceNavigationLink          `pulumi:"resourceNavigationLinks"`
	RouteTable              *RouteTableType                   `pulumi:"routeTable"`
	ServiceEndpoints        []ServiceEndpointPropertiesFormat `pulumi:"serviceEndpoints"`
}





type SubnetTypeInput interface {
	pulumi.Input

	ToSubnetTypeOutput() SubnetTypeOutput
	ToSubnetTypeOutputWithContext(context.Context) SubnetTypeOutput
}

type SubnetTypeArgs struct {
	AddressPrefix           pulumi.StringPtrInput                     `pulumi:"addressPrefix"`
	Etag                    pulumi.StringPtrInput                     `pulumi:"etag"`
	Id                      pulumi.StringPtrInput                     `pulumi:"id"`
	Name                    pulumi.StringPtrInput                     `pulumi:"name"`
	NetworkSecurityGroup    NetworkSecurityGroupTypePtrInput          `pulumi:"networkSecurityGroup"`
	ProvisioningState       pulumi.StringPtrInput                     `pulumi:"provisioningState"`
	ResourceNavigationLinks ResourceNavigationLinkArrayInput          `pulumi:"resourceNavigationLinks"`
	RouteTable              RouteTableTypePtrInput                    `pulumi:"routeTable"`
	ServiceEndpoints        ServiceEndpointPropertiesFormatArrayInput `pulumi:"serviceEndpoints"`
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

func (i SubnetTypeArgs) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return i.ToSubnetTypePtrOutputWithContext(context.Background())
}

func (i SubnetTypeArgs) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypeOutput).ToSubnetTypePtrOutputWithContext(ctx)
}









type SubnetTypePtrInput interface {
	pulumi.Input

	ToSubnetTypePtrOutput() SubnetTypePtrOutput
	ToSubnetTypePtrOutputWithContext(context.Context) SubnetTypePtrOutput
}

type subnetTypePtrType SubnetTypeArgs

func SubnetTypePtr(v *SubnetTypeArgs) SubnetTypePtrInput {
	return (*subnetTypePtrType)(v)
}

func (*subnetTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetType)(nil)).Elem()
}

func (i *subnetTypePtrType) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return i.ToSubnetTypePtrOutputWithContext(context.Background())
}

func (i *subnetTypePtrType) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypePtrOutput)
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

func (o SubnetTypeOutput) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return o.ToSubnetTypePtrOutputWithContext(context.Background())
}

func (o SubnetTypeOutput) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetType) *SubnetType {
		return &v
	}).(SubnetTypePtrOutput)
}

func (o SubnetTypeOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) NetworkSecurityGroup() NetworkSecurityGroupTypePtrOutput {
	return o.ApplyT(func(v SubnetType) *NetworkSecurityGroupType { return v.NetworkSecurityGroup }).(NetworkSecurityGroupTypePtrOutput)
}

func (o SubnetTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) ResourceNavigationLinks() ResourceNavigationLinkArrayOutput {
	return o.ApplyT(func(v SubnetType) []ResourceNavigationLink { return v.ResourceNavigationLinks }).(ResourceNavigationLinkArrayOutput)
}

func (o SubnetTypeOutput) RouteTable() RouteTableTypePtrOutput {
	return o.ApplyT(func(v SubnetType) *RouteTableType { return v.RouteTable }).(RouteTableTypePtrOutput)
}

func (o SubnetTypeOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatArrayOutput {
	return o.ApplyT(func(v SubnetType) []ServiceEndpointPropertiesFormat { return v.ServiceEndpoints }).(ServiceEndpointPropertiesFormatArrayOutput)
}

type SubnetTypePtrOutput struct{ *pulumi.OutputState }

func (SubnetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetType)(nil)).Elem()
}

func (o SubnetTypePtrOutput) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return o
}

func (o SubnetTypePtrOutput) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return o
}

func (o SubnetTypePtrOutput) Elem() SubnetTypeOutput {
	return o.ApplyT(func(v *SubnetType) SubnetType {
		if v != nil {
			return *v
		}
		var ret SubnetType
		return ret
	}).(SubnetTypeOutput)
}

func (o SubnetTypePtrOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.AddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) NetworkSecurityGroup() NetworkSecurityGroupTypePtrOutput {
	return o.ApplyT(func(v *SubnetType) *NetworkSecurityGroupType {
		if v == nil {
			return nil
		}
		return v.NetworkSecurityGroup
	}).(NetworkSecurityGroupTypePtrOutput)
}

func (o SubnetTypePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) ResourceNavigationLinks() ResourceNavigationLinkArrayOutput {
	return o.ApplyT(func(v *SubnetType) []ResourceNavigationLink {
		if v == nil {
			return nil
		}
		return v.ResourceNavigationLinks
	}).(ResourceNavigationLinkArrayOutput)
}

func (o SubnetTypePtrOutput) RouteTable() RouteTableTypePtrOutput {
	return o.ApplyT(func(v *SubnetType) *RouteTableType {
		if v == nil {
			return nil
		}
		return v.RouteTable
	}).(RouteTableTypePtrOutput)
}

func (o SubnetTypePtrOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatArrayOutput {
	return o.ApplyT(func(v *SubnetType) []ServiceEndpointPropertiesFormat {
		if v == nil {
			return nil
		}
		return v.ServiceEndpoints
	}).(ServiceEndpointPropertiesFormatArrayOutput)
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
	AddressPrefix           *string                                   `pulumi:"addressPrefix"`
	Etag                    *string                                   `pulumi:"etag"`
	Id                      *string                                   `pulumi:"id"`
	IpConfigurations        []IPConfigurationResponse                 `pulumi:"ipConfigurations"`
	Name                    *string                                   `pulumi:"name"`
	NetworkSecurityGroup    *NetworkSecurityGroupResponse             `pulumi:"networkSecurityGroup"`
	ProvisioningState       *string                                   `pulumi:"provisioningState"`
	ResourceNavigationLinks []ResourceNavigationLinkResponse          `pulumi:"resourceNavigationLinks"`
	RouteTable              *RouteTableResponse                       `pulumi:"routeTable"`
	ServiceEndpoints        []ServiceEndpointPropertiesFormatResponse `pulumi:"serviceEndpoints"`
}





type SubnetResponseInput interface {
	pulumi.Input

	ToSubnetResponseOutput() SubnetResponseOutput
	ToSubnetResponseOutputWithContext(context.Context) SubnetResponseOutput
}

type SubnetResponseArgs struct {
	AddressPrefix           pulumi.StringPtrInput                             `pulumi:"addressPrefix"`
	Etag                    pulumi.StringPtrInput                             `pulumi:"etag"`
	Id                      pulumi.StringPtrInput                             `pulumi:"id"`
	IpConfigurations        IPConfigurationResponseArrayInput                 `pulumi:"ipConfigurations"`
	Name                    pulumi.StringPtrInput                             `pulumi:"name"`
	NetworkSecurityGroup    NetworkSecurityGroupResponsePtrInput              `pulumi:"networkSecurityGroup"`
	ProvisioningState       pulumi.StringPtrInput                             `pulumi:"provisioningState"`
	ResourceNavigationLinks ResourceNavigationLinkResponseArrayInput          `pulumi:"resourceNavigationLinks"`
	RouteTable              RouteTableResponsePtrInput                        `pulumi:"routeTable"`
	ServiceEndpoints        ServiceEndpointPropertiesFormatResponseArrayInput `pulumi:"serviceEndpoints"`
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

func (i SubnetResponseArgs) ToSubnetResponsePtrOutput() SubnetResponsePtrOutput {
	return i.ToSubnetResponsePtrOutputWithContext(context.Background())
}

func (i SubnetResponseArgs) ToSubnetResponsePtrOutputWithContext(ctx context.Context) SubnetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResponseOutput).ToSubnetResponsePtrOutputWithContext(ctx)
}









type SubnetResponsePtrInput interface {
	pulumi.Input

	ToSubnetResponsePtrOutput() SubnetResponsePtrOutput
	ToSubnetResponsePtrOutputWithContext(context.Context) SubnetResponsePtrOutput
}

type subnetResponsePtrType SubnetResponseArgs

func SubnetResponsePtr(v *SubnetResponseArgs) SubnetResponsePtrInput {
	return (*subnetResponsePtrType)(v)
}

func (*subnetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetResponse)(nil)).Elem()
}

func (i *subnetResponsePtrType) ToSubnetResponsePtrOutput() SubnetResponsePtrOutput {
	return i.ToSubnetResponsePtrOutputWithContext(context.Background())
}

func (i *subnetResponsePtrType) ToSubnetResponsePtrOutputWithContext(ctx context.Context) SubnetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResponsePtrOutput)
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

func (o SubnetResponseOutput) ToSubnetResponsePtrOutput() SubnetResponsePtrOutput {
	return o.ToSubnetResponsePtrOutputWithContext(context.Background())
}

func (o SubnetResponseOutput) ToSubnetResponsePtrOutputWithContext(ctx context.Context) SubnetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetResponse) *SubnetResponse {
		return &v
	}).(SubnetResponsePtrOutput)
}

func (o SubnetResponseOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []IPConfigurationResponse { return v.IpConfigurations }).(IPConfigurationResponseArrayOutput)
}

func (o SubnetResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v SubnetResponse) *NetworkSecurityGroupResponse { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o SubnetResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) ResourceNavigationLinks() ResourceNavigationLinkResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []ResourceNavigationLinkResponse { return v.ResourceNavigationLinks }).(ResourceNavigationLinkResponseArrayOutput)
}

func (o SubnetResponseOutput) RouteTable() RouteTableResponsePtrOutput {
	return o.ApplyT(func(v SubnetResponse) *RouteTableResponse { return v.RouteTable }).(RouteTableResponsePtrOutput)
}

func (o SubnetResponseOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []ServiceEndpointPropertiesFormatResponse { return v.ServiceEndpoints }).(ServiceEndpointPropertiesFormatResponseArrayOutput)
}

type SubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetResponse)(nil)).Elem()
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutput() SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutputWithContext(ctx context.Context) SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) Elem() SubnetResponseOutput {
	return o.ApplyT(func(v *SubnetResponse) SubnetResponse {
		if v != nil {
			return *v
		}
		var ret SubnetResponse
		return ret
	}).(SubnetResponseOutput)
}

func (o SubnetResponsePtrOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.AddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []IPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.IpConfigurations
	}).(IPConfigurationResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *NetworkSecurityGroupResponse {
		if v == nil {
			return nil
		}
		return v.NetworkSecurityGroup
	}).(NetworkSecurityGroupResponsePtrOutput)
}

func (o SubnetResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) ResourceNavigationLinks() ResourceNavigationLinkResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []ResourceNavigationLinkResponse {
		if v == nil {
			return nil
		}
		return v.ResourceNavigationLinks
	}).(ResourceNavigationLinkResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) RouteTable() RouteTableResponsePtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *RouteTableResponse {
		if v == nil {
			return nil
		}
		return v.RouteTable
	}).(RouteTableResponsePtrOutput)
}

func (o SubnetResponsePtrOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []ServiceEndpointPropertiesFormatResponse {
		if v == nil {
			return nil
		}
		return v.ServiceEndpoints
	}).(ServiceEndpointPropertiesFormatResponseArrayOutput)
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

type TunnelConnectionHealthResponse struct {
	ConnectionStatus                 string  `pulumi:"connectionStatus"`
	EgressBytesTransferred           float64 `pulumi:"egressBytesTransferred"`
	IngressBytesTransferred          float64 `pulumi:"ingressBytesTransferred"`
	LastConnectionEstablishedUtcTime string  `pulumi:"lastConnectionEstablishedUtcTime"`
	Tunnel                           string  `pulumi:"tunnel"`
}





type TunnelConnectionHealthResponseInput interface {
	pulumi.Input

	ToTunnelConnectionHealthResponseOutput() TunnelConnectionHealthResponseOutput
	ToTunnelConnectionHealthResponseOutputWithContext(context.Context) TunnelConnectionHealthResponseOutput
}

type TunnelConnectionHealthResponseArgs struct {
	ConnectionStatus                 pulumi.StringInput  `pulumi:"connectionStatus"`
	EgressBytesTransferred           pulumi.Float64Input `pulumi:"egressBytesTransferred"`
	IngressBytesTransferred          pulumi.Float64Input `pulumi:"ingressBytesTransferred"`
	LastConnectionEstablishedUtcTime pulumi.StringInput  `pulumi:"lastConnectionEstablishedUtcTime"`
	Tunnel                           pulumi.StringInput  `pulumi:"tunnel"`
}

func (TunnelConnectionHealthResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TunnelConnectionHealthResponse)(nil)).Elem()
}

func (i TunnelConnectionHealthResponseArgs) ToTunnelConnectionHealthResponseOutput() TunnelConnectionHealthResponseOutput {
	return i.ToTunnelConnectionHealthResponseOutputWithContext(context.Background())
}

func (i TunnelConnectionHealthResponseArgs) ToTunnelConnectionHealthResponseOutputWithContext(ctx context.Context) TunnelConnectionHealthResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelConnectionHealthResponseOutput)
}





type TunnelConnectionHealthResponseArrayInput interface {
	pulumi.Input

	ToTunnelConnectionHealthResponseArrayOutput() TunnelConnectionHealthResponseArrayOutput
	ToTunnelConnectionHealthResponseArrayOutputWithContext(context.Context) TunnelConnectionHealthResponseArrayOutput
}

type TunnelConnectionHealthResponseArray []TunnelConnectionHealthResponseInput

func (TunnelConnectionHealthResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TunnelConnectionHealthResponse)(nil)).Elem()
}

func (i TunnelConnectionHealthResponseArray) ToTunnelConnectionHealthResponseArrayOutput() TunnelConnectionHealthResponseArrayOutput {
	return i.ToTunnelConnectionHealthResponseArrayOutputWithContext(context.Background())
}

func (i TunnelConnectionHealthResponseArray) ToTunnelConnectionHealthResponseArrayOutputWithContext(ctx context.Context) TunnelConnectionHealthResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelConnectionHealthResponseArrayOutput)
}

type TunnelConnectionHealthResponseOutput struct{ *pulumi.OutputState }

func (TunnelConnectionHealthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TunnelConnectionHealthResponse)(nil)).Elem()
}

func (o TunnelConnectionHealthResponseOutput) ToTunnelConnectionHealthResponseOutput() TunnelConnectionHealthResponseOutput {
	return o
}

func (o TunnelConnectionHealthResponseOutput) ToTunnelConnectionHealthResponseOutputWithContext(ctx context.Context) TunnelConnectionHealthResponseOutput {
	return o
}

func (o TunnelConnectionHealthResponseOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o TunnelConnectionHealthResponseOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o TunnelConnectionHealthResponseOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o TunnelConnectionHealthResponseOutput) LastConnectionEstablishedUtcTime() pulumi.StringOutput {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) string { return v.LastConnectionEstablishedUtcTime }).(pulumi.StringOutput)
}

func (o TunnelConnectionHealthResponseOutput) Tunnel() pulumi.StringOutput {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) string { return v.Tunnel }).(pulumi.StringOutput)
}

type TunnelConnectionHealthResponseArrayOutput struct{ *pulumi.OutputState }

func (TunnelConnectionHealthResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TunnelConnectionHealthResponse)(nil)).Elem()
}

func (o TunnelConnectionHealthResponseArrayOutput) ToTunnelConnectionHealthResponseArrayOutput() TunnelConnectionHealthResponseArrayOutput {
	return o
}

func (o TunnelConnectionHealthResponseArrayOutput) ToTunnelConnectionHealthResponseArrayOutputWithContext(ctx context.Context) TunnelConnectionHealthResponseArrayOutput {
	return o
}

func (o TunnelConnectionHealthResponseArrayOutput) Index(i pulumi.IntInput) TunnelConnectionHealthResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TunnelConnectionHealthResponse {
		return vs[0].([]TunnelConnectionHealthResponse)[vs[1].(int)]
	}).(TunnelConnectionHealthResponseOutput)
}

type VirtualNetworkGatewayType struct {
	ActiveActive           *bool                                  `pulumi:"activeActive"`
	BgpSettings            *BgpSettings                           `pulumi:"bgpSettings"`
	EnableBgp              *bool                                  `pulumi:"enableBgp"`
	Etag                   *string                                `pulumi:"etag"`
	GatewayDefaultSite     *SubResource                           `pulumi:"gatewayDefaultSite"`
	GatewayType            *string                                `pulumi:"gatewayType"`
	Id                     *string                                `pulumi:"id"`
	IpConfigurations       []VirtualNetworkGatewayIPConfiguration `pulumi:"ipConfigurations"`
	Location               *string                                `pulumi:"location"`
	ResourceGuid           *string                                `pulumi:"resourceGuid"`
	Sku                    *VirtualNetworkGatewaySku              `pulumi:"sku"`
	Tags                   map[string]string                      `pulumi:"tags"`
	VpnClientConfiguration *VpnClientConfiguration                `pulumi:"vpnClientConfiguration"`
	VpnType                *string                                `pulumi:"vpnType"`
}





type VirtualNetworkGatewayTypeInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayTypeOutput() VirtualNetworkGatewayTypeOutput
	ToVirtualNetworkGatewayTypeOutputWithContext(context.Context) VirtualNetworkGatewayTypeOutput
}

type VirtualNetworkGatewayTypeArgs struct {
	ActiveActive           pulumi.BoolPtrInput                            `pulumi:"activeActive"`
	BgpSettings            BgpSettingsPtrInput                            `pulumi:"bgpSettings"`
	EnableBgp              pulumi.BoolPtrInput                            `pulumi:"enableBgp"`
	Etag                   pulumi.StringPtrInput                          `pulumi:"etag"`
	GatewayDefaultSite     SubResourcePtrInput                            `pulumi:"gatewayDefaultSite"`
	GatewayType            pulumi.StringPtrInput                          `pulumi:"gatewayType"`
	Id                     pulumi.StringPtrInput                          `pulumi:"id"`
	IpConfigurations       VirtualNetworkGatewayIPConfigurationArrayInput `pulumi:"ipConfigurations"`
	Location               pulumi.StringPtrInput                          `pulumi:"location"`
	ResourceGuid           pulumi.StringPtrInput                          `pulumi:"resourceGuid"`
	Sku                    VirtualNetworkGatewaySkuPtrInput               `pulumi:"sku"`
	Tags                   pulumi.StringMapInput                          `pulumi:"tags"`
	VpnClientConfiguration VpnClientConfigurationPtrInput                 `pulumi:"vpnClientConfiguration"`
	VpnType                pulumi.StringPtrInput                          `pulumi:"vpnType"`
}

func (VirtualNetworkGatewayTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayType)(nil)).Elem()
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypeOutput() VirtualNetworkGatewayTypeOutput {
	return i.ToVirtualNetworkGatewayTypeOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypeOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayTypeOutput)
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return i.ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayTypeOutput).ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx)
}









type VirtualNetworkGatewayTypePtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput
	ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Context) VirtualNetworkGatewayTypePtrOutput
}

type virtualNetworkGatewayTypePtrType VirtualNetworkGatewayTypeArgs

func VirtualNetworkGatewayTypePtr(v *VirtualNetworkGatewayTypeArgs) VirtualNetworkGatewayTypePtrInput {
	return (*virtualNetworkGatewayTypePtrType)(v)
}

func (*virtualNetworkGatewayTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayType)(nil)).Elem()
}

func (i *virtualNetworkGatewayTypePtrType) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return i.ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkGatewayTypePtrType) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayTypePtrOutput)
}

type VirtualNetworkGatewayTypeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayType)(nil)).Elem()
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypeOutput() VirtualNetworkGatewayTypeOutput {
	return o
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypeOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeOutput {
	return o
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return o.ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewayType) *VirtualNetworkGatewayType {
		return &v
	}).(VirtualNetworkGatewayTypePtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *bool { return v.ActiveActive }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) BgpSettings() BgpSettingsPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *BgpSettings { return v.BgpSettings }).(BgpSettingsPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) GatewayDefaultSite() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *SubResource { return v.GatewayDefaultSite }).(SubResourcePtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.GatewayType }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) []VirtualNetworkGatewayIPConfiguration { return v.IpConfigurations }).(VirtualNetworkGatewayIPConfigurationArrayOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Sku() VirtualNetworkGatewaySkuPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *VirtualNetworkGatewaySku { return v.Sku }).(VirtualNetworkGatewaySkuPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayTypeOutput) VpnClientConfiguration() VpnClientConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *VpnClientConfiguration { return v.VpnClientConfiguration }).(VpnClientConfigurationPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.VpnType }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayTypePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayType)(nil)).Elem()
}

func (o VirtualNetworkGatewayTypePtrOutput) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return o
}

func (o VirtualNetworkGatewayTypePtrOutput) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return o
}

func (o VirtualNetworkGatewayTypePtrOutput) Elem() VirtualNetworkGatewayTypeOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) VirtualNetworkGatewayType {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewayType
		return ret
	}).(VirtualNetworkGatewayTypeOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *bool {
		if v == nil {
			return nil
		}
		return v.ActiveActive
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) BgpSettings() BgpSettingsPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *BgpSettings {
		if v == nil {
			return nil
		}
		return v.BgpSettings
	}).(BgpSettingsPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *bool {
		if v == nil {
			return nil
		}
		return v.EnableBgp
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) GatewayDefaultSite() SubResourcePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *SubResource {
		if v == nil {
			return nil
		}
		return v.GatewayDefaultSite
	}).(SubResourcePtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.GatewayType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) []VirtualNetworkGatewayIPConfiguration {
		if v == nil {
			return nil
		}
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationArrayOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Sku() VirtualNetworkGatewaySkuPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *VirtualNetworkGatewaySku {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(VirtualNetworkGatewaySkuPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) VpnClientConfiguration() VpnClientConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *VpnClientConfiguration {
		if v == nil {
			return nil
		}
		return v.VpnClientConfiguration
	}).(VpnClientConfigurationPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.VpnType
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayIPConfiguration struct {
	Etag                      *string      `pulumi:"etag"`
	Id                        *string      `pulumi:"id"`
	Name                      *string      `pulumi:"name"`
	PrivateIPAllocationMethod *string      `pulumi:"privateIPAllocationMethod"`
	PublicIPAddress           *SubResource `pulumi:"publicIPAddress"`
	Subnet                    *SubResource `pulumi:"subnet"`
}





type VirtualNetworkGatewayIPConfigurationInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayIPConfigurationOutput() VirtualNetworkGatewayIPConfigurationOutput
	ToVirtualNetworkGatewayIPConfigurationOutputWithContext(context.Context) VirtualNetworkGatewayIPConfigurationOutput
}

type VirtualNetworkGatewayIPConfigurationArgs struct {
	Etag                      pulumi.StringPtrInput `pulumi:"etag"`
	Id                        pulumi.StringPtrInput `pulumi:"id"`
	Name                      pulumi.StringPtrInput `pulumi:"name"`
	PrivateIPAllocationMethod pulumi.StringPtrInput `pulumi:"privateIPAllocationMethod"`
	PublicIPAddress           SubResourcePtrInput   `pulumi:"publicIPAddress"`
	Subnet                    SubResourcePtrInput   `pulumi:"subnet"`
}

func (VirtualNetworkGatewayIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (i VirtualNetworkGatewayIPConfigurationArgs) ToVirtualNetworkGatewayIPConfigurationOutput() VirtualNetworkGatewayIPConfigurationOutput {
	return i.ToVirtualNetworkGatewayIPConfigurationOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayIPConfigurationArgs) ToVirtualNetworkGatewayIPConfigurationOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayIPConfigurationOutput)
}





type VirtualNetworkGatewayIPConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayIPConfigurationArrayOutput() VirtualNetworkGatewayIPConfigurationArrayOutput
	ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(context.Context) VirtualNetworkGatewayIPConfigurationArrayOutput
}

type VirtualNetworkGatewayIPConfigurationArray []VirtualNetworkGatewayIPConfigurationInput

func (VirtualNetworkGatewayIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (i VirtualNetworkGatewayIPConfigurationArray) ToVirtualNetworkGatewayIPConfigurationArrayOutput() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return i.ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayIPConfigurationArray) ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayIPConfigurationArrayOutput)
}

type VirtualNetworkGatewayIPConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationOutput) ToVirtualNetworkGatewayIPConfigurationOutput() VirtualNetworkGatewayIPConfigurationOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationOutput) ToVirtualNetworkGatewayIPConfigurationOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) PublicIPAddress() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *SubResource { return v.PublicIPAddress }).(SubResourcePtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type VirtualNetworkGatewayIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationArrayOutput) ToVirtualNetworkGatewayIPConfigurationArrayOutput() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationArrayOutput) ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualNetworkGatewayIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkGatewayIPConfiguration {
		return vs[0].([]VirtualNetworkGatewayIPConfiguration)[vs[1].(int)]
	}).(VirtualNetworkGatewayIPConfigurationOutput)
}

type VirtualNetworkGatewayIPConfigurationResponse struct {
	Etag                      *string              `pulumi:"etag"`
	Id                        *string              `pulumi:"id"`
	Name                      *string              `pulumi:"name"`
	PrivateIPAllocationMethod *string              `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         string               `pulumi:"provisioningState"`
	PublicIPAddress           *SubResourceResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubResourceResponse `pulumi:"subnet"`
}





type VirtualNetworkGatewayIPConfigurationResponseInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayIPConfigurationResponseOutput() VirtualNetworkGatewayIPConfigurationResponseOutput
	ToVirtualNetworkGatewayIPConfigurationResponseOutputWithContext(context.Context) VirtualNetworkGatewayIPConfigurationResponseOutput
}

type VirtualNetworkGatewayIPConfigurationResponseArgs struct {
	Etag                      pulumi.StringPtrInput       `pulumi:"etag"`
	Id                        pulumi.StringPtrInput       `pulumi:"id"`
	Name                      pulumi.StringPtrInput       `pulumi:"name"`
	PrivateIPAllocationMethod pulumi.StringPtrInput       `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringInput          `pulumi:"provisioningState"`
	PublicIPAddress           SubResourceResponsePtrInput `pulumi:"publicIPAddress"`
	Subnet                    SubResourceResponsePtrInput `pulumi:"subnet"`
}

func (VirtualNetworkGatewayIPConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayIPConfigurationResponse)(nil)).Elem()
}

func (i VirtualNetworkGatewayIPConfigurationResponseArgs) ToVirtualNetworkGatewayIPConfigurationResponseOutput() VirtualNetworkGatewayIPConfigurationResponseOutput {
	return i.ToVirtualNetworkGatewayIPConfigurationResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayIPConfigurationResponseArgs) ToVirtualNetworkGatewayIPConfigurationResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayIPConfigurationResponseOutput)
}





type VirtualNetworkGatewayIPConfigurationResponseArrayInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayIPConfigurationResponseArrayOutput() VirtualNetworkGatewayIPConfigurationResponseArrayOutput
	ToVirtualNetworkGatewayIPConfigurationResponseArrayOutputWithContext(context.Context) VirtualNetworkGatewayIPConfigurationResponseArrayOutput
}

type VirtualNetworkGatewayIPConfigurationResponseArray []VirtualNetworkGatewayIPConfigurationResponseInput

func (VirtualNetworkGatewayIPConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkGatewayIPConfigurationResponse)(nil)).Elem()
}

func (i VirtualNetworkGatewayIPConfigurationResponseArray) ToVirtualNetworkGatewayIPConfigurationResponseArrayOutput() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return i.ToVirtualNetworkGatewayIPConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayIPConfigurationResponseArray) ToVirtualNetworkGatewayIPConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayIPConfigurationResponseArrayOutput)
}

type VirtualNetworkGatewayIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) ToVirtualNetworkGatewayIPConfigurationResponseOutput() VirtualNetworkGatewayIPConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) ToVirtualNetworkGatewayIPConfigurationResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) PublicIPAddress() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *SubResourceResponse { return v.PublicIPAddress }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type VirtualNetworkGatewayIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkGatewayIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationResponseArrayOutput) ToVirtualNetworkGatewayIPConfigurationResponseArrayOutput() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseArrayOutput) ToVirtualNetworkGatewayIPConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkGatewayIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkGatewayIPConfigurationResponse {
		return vs[0].([]VirtualNetworkGatewayIPConfigurationResponse)[vs[1].(int)]
	}).(VirtualNetworkGatewayIPConfigurationResponseOutput)
}

type VirtualNetworkGatewayResponse struct {
	ActiveActive           *bool                                          `pulumi:"activeActive"`
	BgpSettings            *BgpSettingsResponse                           `pulumi:"bgpSettings"`
	EnableBgp              *bool                                          `pulumi:"enableBgp"`
	Etag                   *string                                        `pulumi:"etag"`
	GatewayDefaultSite     *SubResourceResponse                           `pulumi:"gatewayDefaultSite"`
	GatewayType            *string                                        `pulumi:"gatewayType"`
	Id                     *string                                        `pulumi:"id"`
	IpConfigurations       []VirtualNetworkGatewayIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location               *string                                        `pulumi:"location"`
	Name                   string                                         `pulumi:"name"`
	ProvisioningState      string                                         `pulumi:"provisioningState"`
	ResourceGuid           *string                                        `pulumi:"resourceGuid"`
	Sku                    *VirtualNetworkGatewaySkuResponse              `pulumi:"sku"`
	Tags                   map[string]string                              `pulumi:"tags"`
	Type                   string                                         `pulumi:"type"`
	VpnClientConfiguration *VpnClientConfigurationResponse                `pulumi:"vpnClientConfiguration"`
	VpnType                *string                                        `pulumi:"vpnType"`
}





type VirtualNetworkGatewayResponseInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayResponseOutput() VirtualNetworkGatewayResponseOutput
	ToVirtualNetworkGatewayResponseOutputWithContext(context.Context) VirtualNetworkGatewayResponseOutput
}

type VirtualNetworkGatewayResponseArgs struct {
	ActiveActive           pulumi.BoolPtrInput                                    `pulumi:"activeActive"`
	BgpSettings            BgpSettingsResponsePtrInput                            `pulumi:"bgpSettings"`
	EnableBgp              pulumi.BoolPtrInput                                    `pulumi:"enableBgp"`
	Etag                   pulumi.StringPtrInput                                  `pulumi:"etag"`
	GatewayDefaultSite     SubResourceResponsePtrInput                            `pulumi:"gatewayDefaultSite"`
	GatewayType            pulumi.StringPtrInput                                  `pulumi:"gatewayType"`
	Id                     pulumi.StringPtrInput                                  `pulumi:"id"`
	IpConfigurations       VirtualNetworkGatewayIPConfigurationResponseArrayInput `pulumi:"ipConfigurations"`
	Location               pulumi.StringPtrInput                                  `pulumi:"location"`
	Name                   pulumi.StringInput                                     `pulumi:"name"`
	ProvisioningState      pulumi.StringInput                                     `pulumi:"provisioningState"`
	ResourceGuid           pulumi.StringPtrInput                                  `pulumi:"resourceGuid"`
	Sku                    VirtualNetworkGatewaySkuResponsePtrInput               `pulumi:"sku"`
	Tags                   pulumi.StringMapInput                                  `pulumi:"tags"`
	Type                   pulumi.StringInput                                     `pulumi:"type"`
	VpnClientConfiguration VpnClientConfigurationResponsePtrInput                 `pulumi:"vpnClientConfiguration"`
	VpnType                pulumi.StringPtrInput                                  `pulumi:"vpnType"`
}

func (VirtualNetworkGatewayResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayResponse)(nil)).Elem()
}

func (i VirtualNetworkGatewayResponseArgs) ToVirtualNetworkGatewayResponseOutput() VirtualNetworkGatewayResponseOutput {
	return i.ToVirtualNetworkGatewayResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayResponseArgs) ToVirtualNetworkGatewayResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewayResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayResponseOutput)
}

func (i VirtualNetworkGatewayResponseArgs) ToVirtualNetworkGatewayResponsePtrOutput() VirtualNetworkGatewayResponsePtrOutput {
	return i.ToVirtualNetworkGatewayResponsePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayResponseArgs) ToVirtualNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayResponseOutput).ToVirtualNetworkGatewayResponsePtrOutputWithContext(ctx)
}









type VirtualNetworkGatewayResponsePtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayResponsePtrOutput() VirtualNetworkGatewayResponsePtrOutput
	ToVirtualNetworkGatewayResponsePtrOutputWithContext(context.Context) VirtualNetworkGatewayResponsePtrOutput
}

type virtualNetworkGatewayResponsePtrType VirtualNetworkGatewayResponseArgs

func VirtualNetworkGatewayResponsePtr(v *VirtualNetworkGatewayResponseArgs) VirtualNetworkGatewayResponsePtrInput {
	return (*virtualNetworkGatewayResponsePtrType)(v)
}

func (*virtualNetworkGatewayResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayResponse)(nil)).Elem()
}

func (i *virtualNetworkGatewayResponsePtrType) ToVirtualNetworkGatewayResponsePtrOutput() VirtualNetworkGatewayResponsePtrOutput {
	return i.ToVirtualNetworkGatewayResponsePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkGatewayResponsePtrType) ToVirtualNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayResponsePtrOutput)
}

type VirtualNetworkGatewayResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayResponseOutput) ToVirtualNetworkGatewayResponseOutput() VirtualNetworkGatewayResponseOutput {
	return o
}

func (o VirtualNetworkGatewayResponseOutput) ToVirtualNetworkGatewayResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewayResponseOutput {
	return o
}

func (o VirtualNetworkGatewayResponseOutput) ToVirtualNetworkGatewayResponsePtrOutput() VirtualNetworkGatewayResponsePtrOutput {
	return o.ToVirtualNetworkGatewayResponsePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayResponseOutput) ToVirtualNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewayResponse) *VirtualNetworkGatewayResponse {
		return &v
	}).(VirtualNetworkGatewayResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *bool { return v.ActiveActive }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *BgpSettingsResponse { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) GatewayDefaultSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *SubResourceResponse { return v.GatewayDefaultSite }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.GatewayType }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) []VirtualNetworkGatewayIPConfigurationResponse {
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationResponseArrayOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Sku() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *VirtualNetworkGatewaySkuResponse { return v.Sku }).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) VpnClientConfiguration() VpnClientConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *VpnClientConfigurationResponse { return v.VpnClientConfiguration }).(VpnClientConfigurationResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.VpnType }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayResponsePtrOutput) ToVirtualNetworkGatewayResponsePtrOutput() VirtualNetworkGatewayResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewayResponsePtrOutput) ToVirtualNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewayResponsePtrOutput) Elem() VirtualNetworkGatewayResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) VirtualNetworkGatewayResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewayResponse
		return ret
	}).(VirtualNetworkGatewayResponseOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ActiveActive
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *BgpSettingsResponse {
		if v == nil {
			return nil
		}
		return v.BgpSettings
	}).(BgpSettingsResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableBgp
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) GatewayDefaultSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.GatewayDefaultSite
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.GatewayType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) []VirtualNetworkGatewayIPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationResponseArrayOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Sku() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *VirtualNetworkGatewaySkuResponse {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) VpnClientConfiguration() VpnClientConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *VpnClientConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientConfiguration
	}).(VpnClientConfigurationResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.VpnType
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type VirtualNetworkGatewaySkuInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuOutput() VirtualNetworkGatewaySkuOutput
	ToVirtualNetworkGatewaySkuOutputWithContext(context.Context) VirtualNetworkGatewaySkuOutput
}

type VirtualNetworkGatewaySkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (VirtualNetworkGatewaySkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySku)(nil)).Elem()
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuOutput() VirtualNetworkGatewaySkuOutput {
	return i.ToVirtualNetworkGatewaySkuOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuOutput)
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return i.ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuOutput).ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx)
}









type VirtualNetworkGatewaySkuPtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput
	ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Context) VirtualNetworkGatewaySkuPtrOutput
}

type virtualNetworkGatewaySkuPtrType VirtualNetworkGatewaySkuArgs

func VirtualNetworkGatewaySkuPtr(v *VirtualNetworkGatewaySkuArgs) VirtualNetworkGatewaySkuPtrInput {
	return (*virtualNetworkGatewaySkuPtrType)(v)
}

func (*virtualNetworkGatewaySkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySku)(nil)).Elem()
}

func (i *virtualNetworkGatewaySkuPtrType) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return i.ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkGatewaySkuPtrType) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuPtrOutput)
}

type VirtualNetworkGatewaySkuOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySku)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuOutput() VirtualNetworkGatewaySkuOutput {
	return o
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuOutput {
	return o
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return o.ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewaySku) *VirtualNetworkGatewaySku {
		return &v
	}).(VirtualNetworkGatewaySkuPtrOutput)
}

func (o VirtualNetworkGatewaySkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewaySkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySku)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuPtrOutput) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuPtrOutput) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuPtrOutput) Elem() VirtualNetworkGatewaySkuOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) VirtualNetworkGatewaySku {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewaySku
		return ret
	}).(VirtualNetworkGatewaySkuOutput)
}

func (o VirtualNetworkGatewaySkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewaySkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type VirtualNetworkGatewaySkuResponseInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuResponseOutput() VirtualNetworkGatewaySkuResponseOutput
	ToVirtualNetworkGatewaySkuResponseOutputWithContext(context.Context) VirtualNetworkGatewaySkuResponseOutput
}

type VirtualNetworkGatewaySkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (VirtualNetworkGatewaySkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySkuResponse)(nil)).Elem()
}

func (i VirtualNetworkGatewaySkuResponseArgs) ToVirtualNetworkGatewaySkuResponseOutput() VirtualNetworkGatewaySkuResponseOutput {
	return i.ToVirtualNetworkGatewaySkuResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewaySkuResponseArgs) ToVirtualNetworkGatewaySkuResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuResponseOutput)
}

func (i VirtualNetworkGatewaySkuResponseArgs) ToVirtualNetworkGatewaySkuResponsePtrOutput() VirtualNetworkGatewaySkuResponsePtrOutput {
	return i.ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewaySkuResponseArgs) ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuResponseOutput).ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(ctx)
}









type VirtualNetworkGatewaySkuResponsePtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuResponsePtrOutput() VirtualNetworkGatewaySkuResponsePtrOutput
	ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(context.Context) VirtualNetworkGatewaySkuResponsePtrOutput
}

type virtualNetworkGatewaySkuResponsePtrType VirtualNetworkGatewaySkuResponseArgs

func VirtualNetworkGatewaySkuResponsePtr(v *VirtualNetworkGatewaySkuResponseArgs) VirtualNetworkGatewaySkuResponsePtrInput {
	return (*virtualNetworkGatewaySkuResponsePtrType)(v)
}

func (*virtualNetworkGatewaySkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySkuResponse)(nil)).Elem()
}

func (i *virtualNetworkGatewaySkuResponsePtrType) ToVirtualNetworkGatewaySkuResponsePtrOutput() VirtualNetworkGatewaySkuResponsePtrOutput {
	return i.ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkGatewaySkuResponsePtrType) ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

type VirtualNetworkGatewaySkuResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySkuResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuResponseOutput) ToVirtualNetworkGatewaySkuResponseOutput() VirtualNetworkGatewaySkuResponseOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponseOutput) ToVirtualNetworkGatewaySkuResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuResponseOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponseOutput) ToVirtualNetworkGatewaySkuResponsePtrOutput() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuResponseOutput) ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewaySkuResponse) *VirtualNetworkGatewaySkuResponse {
		return &v
	}).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

func (o VirtualNetworkGatewaySkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySkuResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) ToVirtualNetworkGatewaySkuResponsePtrOutput() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Elem() VirtualNetworkGatewaySkuResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) VirtualNetworkGatewaySkuResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewaySkuResponse
		return ret
	}).(VirtualNetworkGatewaySkuResponseOutput)
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringType struct {
	AllowForwardedTraffic     *bool         `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool         `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool         `pulumi:"allowVirtualNetworkAccess"`
	Etag                      *string       `pulumi:"etag"`
	Id                        *string       `pulumi:"id"`
	Name                      *string       `pulumi:"name"`
	PeeringState              *string       `pulumi:"peeringState"`
	ProvisioningState         *string       `pulumi:"provisioningState"`
	RemoteAddressSpace        *AddressSpace `pulumi:"remoteAddressSpace"`
	RemoteVirtualNetwork      *SubResource  `pulumi:"remoteVirtualNetwork"`
	UseRemoteGateways         *bool         `pulumi:"useRemoteGateways"`
}





type VirtualNetworkPeeringTypeInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringTypeOutput() VirtualNetworkPeeringTypeOutput
	ToVirtualNetworkPeeringTypeOutputWithContext(context.Context) VirtualNetworkPeeringTypeOutput
}

type VirtualNetworkPeeringTypeArgs struct {
	AllowForwardedTraffic     pulumi.BoolPtrInput   `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       pulumi.BoolPtrInput   `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess pulumi.BoolPtrInput   `pulumi:"allowVirtualNetworkAccess"`
	Etag                      pulumi.StringPtrInput `pulumi:"etag"`
	Id                        pulumi.StringPtrInput `pulumi:"id"`
	Name                      pulumi.StringPtrInput `pulumi:"name"`
	PeeringState              pulumi.StringPtrInput `pulumi:"peeringState"`
	ProvisioningState         pulumi.StringPtrInput `pulumi:"provisioningState"`
	RemoteAddressSpace        AddressSpacePtrInput  `pulumi:"remoteAddressSpace"`
	RemoteVirtualNetwork      SubResourcePtrInput   `pulumi:"remoteVirtualNetwork"`
	UseRemoteGateways         pulumi.BoolPtrInput   `pulumi:"useRemoteGateways"`
}

func (VirtualNetworkPeeringTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringType)(nil)).Elem()
}

func (i VirtualNetworkPeeringTypeArgs) ToVirtualNetworkPeeringTypeOutput() VirtualNetworkPeeringTypeOutput {
	return i.ToVirtualNetworkPeeringTypeOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringTypeArgs) ToVirtualNetworkPeeringTypeOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringTypeOutput)
}





type VirtualNetworkPeeringTypeArrayInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringTypeArrayOutput() VirtualNetworkPeeringTypeArrayOutput
	ToVirtualNetworkPeeringTypeArrayOutputWithContext(context.Context) VirtualNetworkPeeringTypeArrayOutput
}

type VirtualNetworkPeeringTypeArray []VirtualNetworkPeeringTypeInput

func (VirtualNetworkPeeringTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPeeringType)(nil)).Elem()
}

func (i VirtualNetworkPeeringTypeArray) ToVirtualNetworkPeeringTypeArrayOutput() VirtualNetworkPeeringTypeArrayOutput {
	return i.ToVirtualNetworkPeeringTypeArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringTypeArray) ToVirtualNetworkPeeringTypeArrayOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringTypeArrayOutput)
}

type VirtualNetworkPeeringTypeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringType)(nil)).Elem()
}

func (o VirtualNetworkPeeringTypeOutput) ToVirtualNetworkPeeringTypeOutput() VirtualNetworkPeeringTypeOutput {
	return o
}

func (o VirtualNetworkPeeringTypeOutput) ToVirtualNetworkPeeringTypeOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeOutput {
	return o
}

func (o VirtualNetworkPeeringTypeOutput) AllowForwardedTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.AllowForwardedTraffic }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) AllowGatewayTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.AllowGatewayTransit }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) AllowVirtualNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.AllowVirtualNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) PeeringState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.PeeringState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) RemoteAddressSpace() AddressSpacePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *AddressSpace { return v.RemoteAddressSpace }).(AddressSpacePtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) RemoteVirtualNetwork() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *SubResource { return v.RemoteVirtualNetwork }).(SubResourcePtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) UseRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.UseRemoteGateways }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkPeeringTypeArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPeeringType)(nil)).Elem()
}

func (o VirtualNetworkPeeringTypeArrayOutput) ToVirtualNetworkPeeringTypeArrayOutput() VirtualNetworkPeeringTypeArrayOutput {
	return o
}

func (o VirtualNetworkPeeringTypeArrayOutput) ToVirtualNetworkPeeringTypeArrayOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeArrayOutput {
	return o
}

func (o VirtualNetworkPeeringTypeArrayOutput) Index(i pulumi.IntInput) VirtualNetworkPeeringTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkPeeringType {
		return vs[0].([]VirtualNetworkPeeringType)[vs[1].(int)]
	}).(VirtualNetworkPeeringTypeOutput)
}

type VirtualNetworkPeeringResponse struct {
	AllowForwardedTraffic     *bool                 `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool                 `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool                 `pulumi:"allowVirtualNetworkAccess"`
	Etag                      *string               `pulumi:"etag"`
	Id                        *string               `pulumi:"id"`
	Name                      *string               `pulumi:"name"`
	PeeringState              *string               `pulumi:"peeringState"`
	ProvisioningState         *string               `pulumi:"provisioningState"`
	RemoteAddressSpace        *AddressSpaceResponse `pulumi:"remoteAddressSpace"`
	RemoteVirtualNetwork      *SubResourceResponse  `pulumi:"remoteVirtualNetwork"`
	UseRemoteGateways         *bool                 `pulumi:"useRemoteGateways"`
}





type VirtualNetworkPeeringResponseInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringResponseOutput() VirtualNetworkPeeringResponseOutput
	ToVirtualNetworkPeeringResponseOutputWithContext(context.Context) VirtualNetworkPeeringResponseOutput
}

type VirtualNetworkPeeringResponseArgs struct {
	AllowForwardedTraffic     pulumi.BoolPtrInput          `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       pulumi.BoolPtrInput          `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess pulumi.BoolPtrInput          `pulumi:"allowVirtualNetworkAccess"`
	Etag                      pulumi.StringPtrInput        `pulumi:"etag"`
	Id                        pulumi.StringPtrInput        `pulumi:"id"`
	Name                      pulumi.StringPtrInput        `pulumi:"name"`
	PeeringState              pulumi.StringPtrInput        `pulumi:"peeringState"`
	ProvisioningState         pulumi.StringPtrInput        `pulumi:"provisioningState"`
	RemoteAddressSpace        AddressSpaceResponsePtrInput `pulumi:"remoteAddressSpace"`
	RemoteVirtualNetwork      SubResourceResponsePtrInput  `pulumi:"remoteVirtualNetwork"`
	UseRemoteGateways         pulumi.BoolPtrInput          `pulumi:"useRemoteGateways"`
}

func (VirtualNetworkPeeringResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringResponse)(nil)).Elem()
}

func (i VirtualNetworkPeeringResponseArgs) ToVirtualNetworkPeeringResponseOutput() VirtualNetworkPeeringResponseOutput {
	return i.ToVirtualNetworkPeeringResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringResponseArgs) ToVirtualNetworkPeeringResponseOutputWithContext(ctx context.Context) VirtualNetworkPeeringResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringResponseOutput)
}





type VirtualNetworkPeeringResponseArrayInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringResponseArrayOutput() VirtualNetworkPeeringResponseArrayOutput
	ToVirtualNetworkPeeringResponseArrayOutputWithContext(context.Context) VirtualNetworkPeeringResponseArrayOutput
}

type VirtualNetworkPeeringResponseArray []VirtualNetworkPeeringResponseInput

func (VirtualNetworkPeeringResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPeeringResponse)(nil)).Elem()
}

func (i VirtualNetworkPeeringResponseArray) ToVirtualNetworkPeeringResponseArrayOutput() VirtualNetworkPeeringResponseArrayOutput {
	return i.ToVirtualNetworkPeeringResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringResponseArray) ToVirtualNetworkPeeringResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkPeeringResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringResponseArrayOutput)
}

type VirtualNetworkPeeringResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringResponse)(nil)).Elem()
}

func (o VirtualNetworkPeeringResponseOutput) ToVirtualNetworkPeeringResponseOutput() VirtualNetworkPeeringResponseOutput {
	return o
}

func (o VirtualNetworkPeeringResponseOutput) ToVirtualNetworkPeeringResponseOutputWithContext(ctx context.Context) VirtualNetworkPeeringResponseOutput {
	return o
}

func (o VirtualNetworkPeeringResponseOutput) AllowForwardedTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.AllowForwardedTraffic }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) AllowGatewayTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.AllowGatewayTransit }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) AllowVirtualNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.AllowVirtualNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) PeeringState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.PeeringState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) RemoteAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *AddressSpaceResponse { return v.RemoteAddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) RemoteVirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *SubResourceResponse { return v.RemoteVirtualNetwork }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) UseRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.UseRemoteGateways }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkPeeringResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPeeringResponse)(nil)).Elem()
}

func (o VirtualNetworkPeeringResponseArrayOutput) ToVirtualNetworkPeeringResponseArrayOutput() VirtualNetworkPeeringResponseArrayOutput {
	return o
}

func (o VirtualNetworkPeeringResponseArrayOutput) ToVirtualNetworkPeeringResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkPeeringResponseArrayOutput {
	return o
}

func (o VirtualNetworkPeeringResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkPeeringResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkPeeringResponse {
		return vs[0].([]VirtualNetworkPeeringResponse)[vs[1].(int)]
	}).(VirtualNetworkPeeringResponseOutput)
}

type VpnClientConfiguration struct {
	RadiusServerAddress          *string                       `pulumi:"radiusServerAddress"`
	RadiusServerSecret           *string                       `pulumi:"radiusServerSecret"`
	VpnClientAddressPool         *AddressSpace                 `pulumi:"vpnClientAddressPool"`
	VpnClientProtocols           []string                      `pulumi:"vpnClientProtocols"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificate `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    []VpnClientRootCertificate    `pulumi:"vpnClientRootCertificates"`
}





type VpnClientConfigurationInput interface {
	pulumi.Input

	ToVpnClientConfigurationOutput() VpnClientConfigurationOutput
	ToVpnClientConfigurationOutputWithContext(context.Context) VpnClientConfigurationOutput
}

type VpnClientConfigurationArgs struct {
	RadiusServerAddress          pulumi.StringPtrInput                 `pulumi:"radiusServerAddress"`
	RadiusServerSecret           pulumi.StringPtrInput                 `pulumi:"radiusServerSecret"`
	VpnClientAddressPool         AddressSpacePtrInput                  `pulumi:"vpnClientAddressPool"`
	VpnClientProtocols           pulumi.StringArrayInput               `pulumi:"vpnClientProtocols"`
	VpnClientRevokedCertificates VpnClientRevokedCertificateArrayInput `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    VpnClientRootCertificateArrayInput    `pulumi:"vpnClientRootCertificates"`
}

func (VpnClientConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfiguration)(nil)).Elem()
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationOutput() VpnClientConfigurationOutput {
	return i.ToVpnClientConfigurationOutputWithContext(context.Background())
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationOutputWithContext(ctx context.Context) VpnClientConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationOutput)
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return i.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationOutput).ToVpnClientConfigurationPtrOutputWithContext(ctx)
}









type VpnClientConfigurationPtrInput interface {
	pulumi.Input

	ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput
	ToVpnClientConfigurationPtrOutputWithContext(context.Context) VpnClientConfigurationPtrOutput
}

type vpnClientConfigurationPtrType VpnClientConfigurationArgs

func VpnClientConfigurationPtr(v *VpnClientConfigurationArgs) VpnClientConfigurationPtrInput {
	return (*vpnClientConfigurationPtrType)(v)
}

func (*vpnClientConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfiguration)(nil)).Elem()
}

func (i *vpnClientConfigurationPtrType) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return i.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (i *vpnClientConfigurationPtrType) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationPtrOutput)
}

type VpnClientConfigurationOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfiguration)(nil)).Elem()
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationOutput() VpnClientConfigurationOutput {
	return o
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationOutputWithContext(ctx context.Context) VpnClientConfigurationOutput {
	return o
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return o.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnClientConfiguration) *VpnClientConfiguration {
		return &v
	}).(VpnClientConfigurationPtrOutput)
}

func (o VpnClientConfigurationOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) VpnClientAddressPool() AddressSpacePtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *AddressSpace { return v.VpnClientAddressPool }).(AddressSpacePtrOutput)
}

func (o VpnClientConfigurationOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []string { return v.VpnClientProtocols }).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []VpnClientRevokedCertificate { return v.VpnClientRevokedCertificates }).(VpnClientRevokedCertificateArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientRootCertificates() VpnClientRootCertificateArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []VpnClientRootCertificate { return v.VpnClientRootCertificates }).(VpnClientRootCertificateArrayOutput)
}

type VpnClientConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfiguration)(nil)).Elem()
}

func (o VpnClientConfigurationPtrOutput) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return o
}

func (o VpnClientConfigurationPtrOutput) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return o
}

func (o VpnClientConfigurationPtrOutput) Elem() VpnClientConfigurationOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) VpnClientConfiguration {
		if v != nil {
			return *v
		}
		var ret VpnClientConfiguration
		return ret
	}).(VpnClientConfigurationOutput)
}

func (o VpnClientConfigurationPtrOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerAddress
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerSecret
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientAddressPool() AddressSpacePtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *AddressSpace {
		if v == nil {
			return nil
		}
		return v.VpnClientAddressPool
	}).(AddressSpacePtrOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.VpnClientProtocols
	}).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []VpnClientRevokedCertificate {
		if v == nil {
			return nil
		}
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientRootCertificates() VpnClientRootCertificateArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []VpnClientRootCertificate {
		if v == nil {
			return nil
		}
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateArrayOutput)
}

type VpnClientConfigurationResponse struct {
	RadiusServerAddress          *string                               `pulumi:"radiusServerAddress"`
	RadiusServerSecret           *string                               `pulumi:"radiusServerSecret"`
	VpnClientAddressPool         *AddressSpaceResponse                 `pulumi:"vpnClientAddressPool"`
	VpnClientProtocols           []string                              `pulumi:"vpnClientProtocols"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificateResponse `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    []VpnClientRootCertificateResponse    `pulumi:"vpnClientRootCertificates"`
}





type VpnClientConfigurationResponseInput interface {
	pulumi.Input

	ToVpnClientConfigurationResponseOutput() VpnClientConfigurationResponseOutput
	ToVpnClientConfigurationResponseOutputWithContext(context.Context) VpnClientConfigurationResponseOutput
}

type VpnClientConfigurationResponseArgs struct {
	RadiusServerAddress          pulumi.StringPtrInput                         `pulumi:"radiusServerAddress"`
	RadiusServerSecret           pulumi.StringPtrInput                         `pulumi:"radiusServerSecret"`
	VpnClientAddressPool         AddressSpaceResponsePtrInput                  `pulumi:"vpnClientAddressPool"`
	VpnClientProtocols           pulumi.StringArrayInput                       `pulumi:"vpnClientProtocols"`
	VpnClientRevokedCertificates VpnClientRevokedCertificateResponseArrayInput `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    VpnClientRootCertificateResponseArrayInput    `pulumi:"vpnClientRootCertificates"`
}

func (VpnClientConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfigurationResponse)(nil)).Elem()
}

func (i VpnClientConfigurationResponseArgs) ToVpnClientConfigurationResponseOutput() VpnClientConfigurationResponseOutput {
	return i.ToVpnClientConfigurationResponseOutputWithContext(context.Background())
}

func (i VpnClientConfigurationResponseArgs) ToVpnClientConfigurationResponseOutputWithContext(ctx context.Context) VpnClientConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationResponseOutput)
}

func (i VpnClientConfigurationResponseArgs) ToVpnClientConfigurationResponsePtrOutput() VpnClientConfigurationResponsePtrOutput {
	return i.ToVpnClientConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i VpnClientConfigurationResponseArgs) ToVpnClientConfigurationResponsePtrOutputWithContext(ctx context.Context) VpnClientConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationResponseOutput).ToVpnClientConfigurationResponsePtrOutputWithContext(ctx)
}









type VpnClientConfigurationResponsePtrInput interface {
	pulumi.Input

	ToVpnClientConfigurationResponsePtrOutput() VpnClientConfigurationResponsePtrOutput
	ToVpnClientConfigurationResponsePtrOutputWithContext(context.Context) VpnClientConfigurationResponsePtrOutput
}

type vpnClientConfigurationResponsePtrType VpnClientConfigurationResponseArgs

func VpnClientConfigurationResponsePtr(v *VpnClientConfigurationResponseArgs) VpnClientConfigurationResponsePtrInput {
	return (*vpnClientConfigurationResponsePtrType)(v)
}

func (*vpnClientConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfigurationResponse)(nil)).Elem()
}

func (i *vpnClientConfigurationResponsePtrType) ToVpnClientConfigurationResponsePtrOutput() VpnClientConfigurationResponsePtrOutput {
	return i.ToVpnClientConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *vpnClientConfigurationResponsePtrType) ToVpnClientConfigurationResponsePtrOutputWithContext(ctx context.Context) VpnClientConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationResponsePtrOutput)
}

type VpnClientConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfigurationResponse)(nil)).Elem()
}

func (o VpnClientConfigurationResponseOutput) ToVpnClientConfigurationResponseOutput() VpnClientConfigurationResponseOutput {
	return o
}

func (o VpnClientConfigurationResponseOutput) ToVpnClientConfigurationResponseOutputWithContext(ctx context.Context) VpnClientConfigurationResponseOutput {
	return o
}

func (o VpnClientConfigurationResponseOutput) ToVpnClientConfigurationResponsePtrOutput() VpnClientConfigurationResponsePtrOutput {
	return o.ToVpnClientConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o VpnClientConfigurationResponseOutput) ToVpnClientConfigurationResponsePtrOutputWithContext(ctx context.Context) VpnClientConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnClientConfigurationResponse) *VpnClientConfigurationResponse {
		return &v
	}).(VpnClientConfigurationResponsePtrOutput)
}

func (o VpnClientConfigurationResponseOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientAddressPool() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *AddressSpaceResponse { return v.VpnClientAddressPool }).(AddressSpaceResponsePtrOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []string { return v.VpnClientProtocols }).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []VpnClientRevokedCertificateResponse {
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateResponseArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientRootCertificates() VpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []VpnClientRootCertificateResponse {
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateResponseArrayOutput)
}

type VpnClientConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfigurationResponse)(nil)).Elem()
}

func (o VpnClientConfigurationResponsePtrOutput) ToVpnClientConfigurationResponsePtrOutput() VpnClientConfigurationResponsePtrOutput {
	return o
}

func (o VpnClientConfigurationResponsePtrOutput) ToVpnClientConfigurationResponsePtrOutputWithContext(ctx context.Context) VpnClientConfigurationResponsePtrOutput {
	return o
}

func (o VpnClientConfigurationResponsePtrOutput) Elem() VpnClientConfigurationResponseOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) VpnClientConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VpnClientConfigurationResponse
		return ret
	}).(VpnClientConfigurationResponseOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerAddress
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerSecret
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientAddressPool() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *AddressSpaceResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientAddressPool
	}).(AddressSpaceResponsePtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.VpnClientProtocols
	}).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []VpnClientRevokedCertificateResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateResponseArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientRootCertificates() VpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []VpnClientRootCertificateResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateResponseArrayOutput)
}

type VpnClientRevokedCertificate struct {
	Etag       *string `pulumi:"etag"`
	Id         *string `pulumi:"id"`
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type VpnClientRevokedCertificateInput interface {
	pulumi.Input

	ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput
	ToVpnClientRevokedCertificateOutputWithContext(context.Context) VpnClientRevokedCertificateOutput
}

type VpnClientRevokedCertificateArgs struct {
	Etag       pulumi.StringPtrInput `pulumi:"etag"`
	Id         pulumi.StringPtrInput `pulumi:"id"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnClientRevokedCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnClientRevokedCertificateArgs) ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput {
	return i.ToVpnClientRevokedCertificateOutputWithContext(context.Background())
}

func (i VpnClientRevokedCertificateArgs) ToVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnClientRevokedCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRevokedCertificateOutput)
}





type VpnClientRevokedCertificateArrayInput interface {
	pulumi.Input

	ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput
	ToVpnClientRevokedCertificateArrayOutputWithContext(context.Context) VpnClientRevokedCertificateArrayOutput
}

type VpnClientRevokedCertificateArray []VpnClientRevokedCertificateInput

func (VpnClientRevokedCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnClientRevokedCertificateArray) ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput {
	return i.ToVpnClientRevokedCertificateArrayOutputWithContext(context.Background())
}

func (i VpnClientRevokedCertificateArray) ToVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRevokedCertificateArrayOutput)
}

type VpnClientRevokedCertificateOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnClientRevokedCertificateOutput) ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput {
	return o
}

func (o VpnClientRevokedCertificateOutput) ToVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnClientRevokedCertificateOutput {
	return o
}

func (o VpnClientRevokedCertificateOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnClientRevokedCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnClientRevokedCertificateArrayOutput) ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateArrayOutput) ToVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateArrayOutput) Index(i pulumi.IntInput) VpnClientRevokedCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRevokedCertificate {
		return vs[0].([]VpnClientRevokedCertificate)[vs[1].(int)]
	}).(VpnClientRevokedCertificateOutput)
}

type VpnClientRevokedCertificateResponse struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Thumbprint        *string `pulumi:"thumbprint"`
}





type VpnClientRevokedCertificateResponseInput interface {
	pulumi.Input

	ToVpnClientRevokedCertificateResponseOutput() VpnClientRevokedCertificateResponseOutput
	ToVpnClientRevokedCertificateResponseOutputWithContext(context.Context) VpnClientRevokedCertificateResponseOutput
}

type VpnClientRevokedCertificateResponseArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState pulumi.StringInput    `pulumi:"provisioningState"`
	Thumbprint        pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnClientRevokedCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (i VpnClientRevokedCertificateResponseArgs) ToVpnClientRevokedCertificateResponseOutput() VpnClientRevokedCertificateResponseOutput {
	return i.ToVpnClientRevokedCertificateResponseOutputWithContext(context.Background())
}

func (i VpnClientRevokedCertificateResponseArgs) ToVpnClientRevokedCertificateResponseOutputWithContext(ctx context.Context) VpnClientRevokedCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRevokedCertificateResponseOutput)
}





type VpnClientRevokedCertificateResponseArrayInput interface {
	pulumi.Input

	ToVpnClientRevokedCertificateResponseArrayOutput() VpnClientRevokedCertificateResponseArrayOutput
	ToVpnClientRevokedCertificateResponseArrayOutputWithContext(context.Context) VpnClientRevokedCertificateResponseArrayOutput
}

type VpnClientRevokedCertificateResponseArray []VpnClientRevokedCertificateResponseInput

func (VpnClientRevokedCertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (i VpnClientRevokedCertificateResponseArray) ToVpnClientRevokedCertificateResponseArrayOutput() VpnClientRevokedCertificateResponseArrayOutput {
	return i.ToVpnClientRevokedCertificateResponseArrayOutputWithContext(context.Background())
}

func (i VpnClientRevokedCertificateResponseArray) ToVpnClientRevokedCertificateResponseArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRevokedCertificateResponseArrayOutput)
}

type VpnClientRevokedCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnClientRevokedCertificateResponseOutput) ToVpnClientRevokedCertificateResponseOutput() VpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseOutput) ToVpnClientRevokedCertificateResponseOutputWithContext(ctx context.Context) VpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnClientRevokedCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnClientRevokedCertificateResponseArrayOutput) ToVpnClientRevokedCertificateResponseArrayOutput() VpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseArrayOutput) ToVpnClientRevokedCertificateResponseArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnClientRevokedCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRevokedCertificateResponse {
		return vs[0].([]VpnClientRevokedCertificateResponse)[vs[1].(int)]
	}).(VpnClientRevokedCertificateResponseOutput)
}

type VpnClientRootCertificate struct {
	Etag           *string `pulumi:"etag"`
	Id             *string `pulumi:"id"`
	Name           *string `pulumi:"name"`
	PublicCertData string  `pulumi:"publicCertData"`
}





type VpnClientRootCertificateInput interface {
	pulumi.Input

	ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput
	ToVpnClientRootCertificateOutputWithContext(context.Context) VpnClientRootCertificateOutput
}

type VpnClientRootCertificateArgs struct {
	Etag           pulumi.StringPtrInput `pulumi:"etag"`
	Id             pulumi.StringPtrInput `pulumi:"id"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PublicCertData pulumi.StringInput    `pulumi:"publicCertData"`
}

func (VpnClientRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificate)(nil)).Elem()
}

func (i VpnClientRootCertificateArgs) ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput {
	return i.ToVpnClientRootCertificateOutputWithContext(context.Background())
}

func (i VpnClientRootCertificateArgs) ToVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnClientRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRootCertificateOutput)
}





type VpnClientRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput
	ToVpnClientRootCertificateArrayOutputWithContext(context.Context) VpnClientRootCertificateArrayOutput
}

type VpnClientRootCertificateArray []VpnClientRootCertificateInput

func (VpnClientRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificate)(nil)).Elem()
}

func (i VpnClientRootCertificateArray) ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput {
	return i.ToVpnClientRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnClientRootCertificateArray) ToVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRootCertificateArrayOutput)
}

type VpnClientRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificate)(nil)).Elem()
}

func (o VpnClientRootCertificateOutput) ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput {
	return o
}

func (o VpnClientRootCertificateOutput) ToVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnClientRootCertificateOutput {
	return o
}

func (o VpnClientRootCertificateOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateOutput) PublicCertData() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) string { return v.PublicCertData }).(pulumi.StringOutput)
}

type VpnClientRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificate)(nil)).Elem()
}

func (o VpnClientRootCertificateArrayOutput) ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnClientRootCertificateArrayOutput) ToVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnClientRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnClientRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRootCertificate {
		return vs[0].([]VpnClientRootCertificate)[vs[1].(int)]
	}).(VpnClientRootCertificateOutput)
}

type VpnClientRootCertificateResponse struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	PublicCertData    string  `pulumi:"publicCertData"`
}





type VpnClientRootCertificateResponseInput interface {
	pulumi.Input

	ToVpnClientRootCertificateResponseOutput() VpnClientRootCertificateResponseOutput
	ToVpnClientRootCertificateResponseOutputWithContext(context.Context) VpnClientRootCertificateResponseOutput
}

type VpnClientRootCertificateResponseArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState pulumi.StringInput    `pulumi:"provisioningState"`
	PublicCertData    pulumi.StringInput    `pulumi:"publicCertData"`
}

func (VpnClientRootCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificateResponse)(nil)).Elem()
}

func (i VpnClientRootCertificateResponseArgs) ToVpnClientRootCertificateResponseOutput() VpnClientRootCertificateResponseOutput {
	return i.ToVpnClientRootCertificateResponseOutputWithContext(context.Background())
}

func (i VpnClientRootCertificateResponseArgs) ToVpnClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnClientRootCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRootCertificateResponseOutput)
}





type VpnClientRootCertificateResponseArrayInput interface {
	pulumi.Input

	ToVpnClientRootCertificateResponseArrayOutput() VpnClientRootCertificateResponseArrayOutput
	ToVpnClientRootCertificateResponseArrayOutputWithContext(context.Context) VpnClientRootCertificateResponseArrayOutput
}

type VpnClientRootCertificateResponseArray []VpnClientRootCertificateResponseInput

func (VpnClientRootCertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificateResponse)(nil)).Elem()
}

func (i VpnClientRootCertificateResponseArray) ToVpnClientRootCertificateResponseArrayOutput() VpnClientRootCertificateResponseArrayOutput {
	return i.ToVpnClientRootCertificateResponseArrayOutputWithContext(context.Background())
}

func (i VpnClientRootCertificateResponseArray) ToVpnClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRootCertificateResponseArrayOutput)
}

type VpnClientRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnClientRootCertificateResponseOutput) ToVpnClientRootCertificateResponseOutput() VpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnClientRootCertificateResponseOutput) ToVpnClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnClientRootCertificateResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnClientRootCertificateResponseOutput) PublicCertData() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) string { return v.PublicCertData }).(pulumi.StringOutput)
}

type VpnClientRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnClientRootCertificateResponseArrayOutput) ToVpnClientRootCertificateResponseArrayOutput() VpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRootCertificateResponseArrayOutput) ToVpnClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnClientRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRootCertificateResponse {
		return vs[0].([]VpnClientRootCertificateResponse)[vs[1].(int)]
	}).(VpnClientRootCertificateResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AddressSpaceOutput{})
	pulumi.RegisterOutputType(AddressSpacePtrOutput{})
	pulumi.RegisterOutputType(AddressSpaceResponseOutput{})
	pulumi.RegisterOutputType(AddressSpaceResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayAuthenticationCertificateOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayAuthenticationCertificateArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayAuthenticationCertificateResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayAuthenticationCertificateResponseArrayOutput{})
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
	pulumi.RegisterOutputType(ApplicationGatewayConnectionDrainingOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayConnectionDrainingPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayConnectionDrainingResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayConnectionDrainingResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallDisabledRuleGroupOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallDisabledRuleGroupArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallDisabledRuleGroupResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput{})
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
	pulumi.RegisterOutputType(ApplicationGatewayPathRuleOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayPathRuleArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayPathRuleResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayPathRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeHealthResponseMatchOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeHealthResponseMatchPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeHealthResponseMatchResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeHealthResponseMatchResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRedirectConfigurationOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRedirectConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRedirectConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRedirectConfigurationResponseArrayOutput{})
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
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayUrlPathMapOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayUrlPathMapArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayUrlPathMapResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayUrlPathMapResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayWebApplicationFirewallConfigurationOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationSecurityGroupTypeOutput{})
	pulumi.RegisterOutputType(ApplicationSecurityGroupTypeArrayOutput{})
	pulumi.RegisterOutputType(ApplicationSecurityGroupResponseOutput{})
	pulumi.RegisterOutputType(ApplicationSecurityGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolArrayOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolResponseOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(BgpPeerStatusResponseOutput{})
	pulumi.RegisterOutputType(BgpPeerStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(BgpSettingsOutput{})
	pulumi.RegisterOutputType(BgpSettingsPtrOutput{})
	pulumi.RegisterOutputType(BgpSettingsResponseOutput{})
	pulumi.RegisterOutputType(BgpSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorDestinationOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorDestinationPtrOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorDestinationResponseOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorSourceOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorSourcePtrOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorSourceResponseOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorSourceResponsePtrOutput{})
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
	pulumi.RegisterOutputType(FrontendIPConfigurationOutput{})
	pulumi.RegisterOutputType(FrontendIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(FrontendIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FrontendIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(GatewayRouteResponseOutput{})
	pulumi.RegisterOutputType(GatewayRouteResponseArrayOutput{})
	pulumi.RegisterOutputType(IPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(IPConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(IPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(InboundNatPoolOutput{})
	pulumi.RegisterOutputType(InboundNatPoolArrayOutput{})
	pulumi.RegisterOutputType(InboundNatPoolResponseOutput{})
	pulumi.RegisterOutputType(InboundNatPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(InboundNatRuleTypeOutput{})
	pulumi.RegisterOutputType(InboundNatRuleTypeArrayOutput{})
	pulumi.RegisterOutputType(InboundNatRuleResponseOutput{})
	pulumi.RegisterOutputType(InboundNatRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(IpTagOutput{})
	pulumi.RegisterOutputType(IpTagArrayOutput{})
	pulumi.RegisterOutputType(IpTagResponseOutput{})
	pulumi.RegisterOutputType(IpTagResponseArrayOutput{})
	pulumi.RegisterOutputType(IpsecPolicyOutput{})
	pulumi.RegisterOutputType(IpsecPolicyArrayOutput{})
	pulumi.RegisterOutputType(IpsecPolicyResponseOutput{})
	pulumi.RegisterOutputType(IpsecPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(Ipv6ExpressRouteCircuitPeeringConfigOutput{})
	pulumi.RegisterOutputType(Ipv6ExpressRouteCircuitPeeringConfigPtrOutput{})
	pulumi.RegisterOutputType(Ipv6ExpressRouteCircuitPeeringConfigResponseOutput{})
	pulumi.RegisterOutputType(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuPtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(LocalNetworkGatewayTypeOutput{})
	pulumi.RegisterOutputType(LocalNetworkGatewayTypePtrOutput{})
	pulumi.RegisterOutputType(LocalNetworkGatewayResponseOutput{})
	pulumi.RegisterOutputType(LocalNetworkGatewayResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupTypeOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupTypePtrOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupResponseOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupResponsePtrOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleArrayOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleResponseOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(PacketCaptureFilterOutput{})
	pulumi.RegisterOutputType(PacketCaptureFilterArrayOutput{})
	pulumi.RegisterOutputType(PacketCaptureFilterResponseOutput{})
	pulumi.RegisterOutputType(PacketCaptureFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(PacketCaptureStorageLocationOutput{})
	pulumi.RegisterOutputType(PacketCaptureStorageLocationPtrOutput{})
	pulumi.RegisterOutputType(PacketCaptureStorageLocationResponseOutput{})
	pulumi.RegisterOutputType(PacketCaptureStorageLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(ProbeOutput{})
	pulumi.RegisterOutputType(ProbeArrayOutput{})
	pulumi.RegisterOutputType(ProbeResponseOutput{})
	pulumi.RegisterOutputType(ProbeResponseArrayOutput{})
	pulumi.RegisterOutputType(PublicIPAddressTypeOutput{})
	pulumi.RegisterOutputType(PublicIPAddressTypePtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressDnsSettingsOutput{})
	pulumi.RegisterOutputType(PublicIPAddressDnsSettingsPtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressDnsSettingsResponseOutput{})
	pulumi.RegisterOutputType(PublicIPAddressDnsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressResponseOutput{})
	pulumi.RegisterOutputType(PublicIPAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuPtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuResponseOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceNavigationLinkOutput{})
	pulumi.RegisterOutputType(ResourceNavigationLinkArrayOutput{})
	pulumi.RegisterOutputType(ResourceNavigationLinkResponseOutput{})
	pulumi.RegisterOutputType(ResourceNavigationLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(RouteTypeOutput{})
	pulumi.RegisterOutputType(RouteTypeArrayOutput{})
	pulumi.RegisterOutputType(RouteFilterTypeOutput{})
	pulumi.RegisterOutputType(RouteFilterTypePtrOutput{})
	pulumi.RegisterOutputType(RouteFilterResponseOutput{})
	pulumi.RegisterOutputType(RouteFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleTypeOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleTypeArrayOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleResponseOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RouteResponseOutput{})
	pulumi.RegisterOutputType(RouteResponseArrayOutput{})
	pulumi.RegisterOutputType(RouteTableTypeOutput{})
	pulumi.RegisterOutputType(RouteTableTypePtrOutput{})
	pulumi.RegisterOutputType(RouteTableResponseOutput{})
	pulumi.RegisterOutputType(RouteTableResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityRuleTypeOutput{})
	pulumi.RegisterOutputType(SecurityRuleTypeArrayOutput{})
	pulumi.RegisterOutputType(SecurityRuleResponseOutput{})
	pulumi.RegisterOutputType(SecurityRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPropertiesFormatOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPropertiesFormatArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPropertiesFormatResponseOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPropertiesFormatResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SubnetTypeOutput{})
	pulumi.RegisterOutputType(SubnetTypePtrOutput{})
	pulumi.RegisterOutputType(SubnetTypeArrayOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetResponseArrayOutput{})
	pulumi.RegisterOutputType(TunnelConnectionHealthResponseOutput{})
	pulumi.RegisterOutputType(TunnelConnectionHealthResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayTypePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringTypeArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateResponseArrayOutput{})
}
