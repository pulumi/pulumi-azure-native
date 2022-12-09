


package servicenetworking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssociationSubnet struct {
	Id string `pulumi:"id"`
}





type AssociationSubnetInput interface {
	pulumi.Input

	ToAssociationSubnetOutput() AssociationSubnetOutput
	ToAssociationSubnetOutputWithContext(context.Context) AssociationSubnetOutput
}

type AssociationSubnetArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (AssociationSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationSubnet)(nil)).Elem()
}

func (i AssociationSubnetArgs) ToAssociationSubnetOutput() AssociationSubnetOutput {
	return i.ToAssociationSubnetOutputWithContext(context.Background())
}

func (i AssociationSubnetArgs) ToAssociationSubnetOutputWithContext(ctx context.Context) AssociationSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationSubnetOutput)
}

func (i AssociationSubnetArgs) ToAssociationSubnetPtrOutput() AssociationSubnetPtrOutput {
	return i.ToAssociationSubnetPtrOutputWithContext(context.Background())
}

func (i AssociationSubnetArgs) ToAssociationSubnetPtrOutputWithContext(ctx context.Context) AssociationSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationSubnetOutput).ToAssociationSubnetPtrOutputWithContext(ctx)
}









type AssociationSubnetPtrInput interface {
	pulumi.Input

	ToAssociationSubnetPtrOutput() AssociationSubnetPtrOutput
	ToAssociationSubnetPtrOutputWithContext(context.Context) AssociationSubnetPtrOutput
}

type associationSubnetPtrType AssociationSubnetArgs

func AssociationSubnetPtr(v *AssociationSubnetArgs) AssociationSubnetPtrInput {
	return (*associationSubnetPtrType)(v)
}

func (*associationSubnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationSubnet)(nil)).Elem()
}

func (i *associationSubnetPtrType) ToAssociationSubnetPtrOutput() AssociationSubnetPtrOutput {
	return i.ToAssociationSubnetPtrOutputWithContext(context.Background())
}

func (i *associationSubnetPtrType) ToAssociationSubnetPtrOutputWithContext(ctx context.Context) AssociationSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationSubnetPtrOutput)
}

type AssociationSubnetOutput struct{ *pulumi.OutputState }

func (AssociationSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationSubnet)(nil)).Elem()
}

func (o AssociationSubnetOutput) ToAssociationSubnetOutput() AssociationSubnetOutput {
	return o
}

func (o AssociationSubnetOutput) ToAssociationSubnetOutputWithContext(ctx context.Context) AssociationSubnetOutput {
	return o
}

func (o AssociationSubnetOutput) ToAssociationSubnetPtrOutput() AssociationSubnetPtrOutput {
	return o.ToAssociationSubnetPtrOutputWithContext(context.Background())
}

func (o AssociationSubnetOutput) ToAssociationSubnetPtrOutputWithContext(ctx context.Context) AssociationSubnetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssociationSubnet) *AssociationSubnet {
		return &v
	}).(AssociationSubnetPtrOutput)
}

func (o AssociationSubnetOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AssociationSubnet) string { return v.Id }).(pulumi.StringOutput)
}

type AssociationSubnetPtrOutput struct{ *pulumi.OutputState }

func (AssociationSubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationSubnet)(nil)).Elem()
}

func (o AssociationSubnetPtrOutput) ToAssociationSubnetPtrOutput() AssociationSubnetPtrOutput {
	return o
}

func (o AssociationSubnetPtrOutput) ToAssociationSubnetPtrOutputWithContext(ctx context.Context) AssociationSubnetPtrOutput {
	return o
}

func (o AssociationSubnetPtrOutput) Elem() AssociationSubnetOutput {
	return o.ApplyT(func(v *AssociationSubnet) AssociationSubnet {
		if v != nil {
			return *v
		}
		var ret AssociationSubnet
		return ret
	}).(AssociationSubnetOutput)
}

func (o AssociationSubnetPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssociationSubnet) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type AssociationSubnetResponse struct {
	Id string `pulumi:"id"`
}

type AssociationSubnetResponseOutput struct{ *pulumi.OutputState }

func (AssociationSubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationSubnetResponse)(nil)).Elem()
}

func (o AssociationSubnetResponseOutput) ToAssociationSubnetResponseOutput() AssociationSubnetResponseOutput {
	return o
}

func (o AssociationSubnetResponseOutput) ToAssociationSubnetResponseOutputWithContext(ctx context.Context) AssociationSubnetResponseOutput {
	return o
}

func (o AssociationSubnetResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AssociationSubnetResponse) string { return v.Id }).(pulumi.StringOutput)
}

type AssociationSubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (AssociationSubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationSubnetResponse)(nil)).Elem()
}

func (o AssociationSubnetResponsePtrOutput) ToAssociationSubnetResponsePtrOutput() AssociationSubnetResponsePtrOutput {
	return o
}

func (o AssociationSubnetResponsePtrOutput) ToAssociationSubnetResponsePtrOutputWithContext(ctx context.Context) AssociationSubnetResponsePtrOutput {
	return o
}

func (o AssociationSubnetResponsePtrOutput) Elem() AssociationSubnetResponseOutput {
	return o.ApplyT(func(v *AssociationSubnetResponse) AssociationSubnetResponse {
		if v != nil {
			return *v
		}
		var ret AssociationSubnetResponse
		return ret
	}).(AssociationSubnetResponseOutput)
}

func (o AssociationSubnetResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssociationSubnetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type FrontendPropertiesIPAddress struct {
	Id string `pulumi:"id"`
}





type FrontendPropertiesIPAddressInput interface {
	pulumi.Input

	ToFrontendPropertiesIPAddressOutput() FrontendPropertiesIPAddressOutput
	ToFrontendPropertiesIPAddressOutputWithContext(context.Context) FrontendPropertiesIPAddressOutput
}

type FrontendPropertiesIPAddressArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (FrontendPropertiesIPAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendPropertiesIPAddress)(nil)).Elem()
}

func (i FrontendPropertiesIPAddressArgs) ToFrontendPropertiesIPAddressOutput() FrontendPropertiesIPAddressOutput {
	return i.ToFrontendPropertiesIPAddressOutputWithContext(context.Background())
}

func (i FrontendPropertiesIPAddressArgs) ToFrontendPropertiesIPAddressOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendPropertiesIPAddressOutput)
}

func (i FrontendPropertiesIPAddressArgs) ToFrontendPropertiesIPAddressPtrOutput() FrontendPropertiesIPAddressPtrOutput {
	return i.ToFrontendPropertiesIPAddressPtrOutputWithContext(context.Background())
}

func (i FrontendPropertiesIPAddressArgs) ToFrontendPropertiesIPAddressPtrOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendPropertiesIPAddressOutput).ToFrontendPropertiesIPAddressPtrOutputWithContext(ctx)
}









type FrontendPropertiesIPAddressPtrInput interface {
	pulumi.Input

	ToFrontendPropertiesIPAddressPtrOutput() FrontendPropertiesIPAddressPtrOutput
	ToFrontendPropertiesIPAddressPtrOutputWithContext(context.Context) FrontendPropertiesIPAddressPtrOutput
}

type frontendPropertiesIPAddressPtrType FrontendPropertiesIPAddressArgs

func FrontendPropertiesIPAddressPtr(v *FrontendPropertiesIPAddressArgs) FrontendPropertiesIPAddressPtrInput {
	return (*frontendPropertiesIPAddressPtrType)(v)
}

func (*frontendPropertiesIPAddressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendPropertiesIPAddress)(nil)).Elem()
}

func (i *frontendPropertiesIPAddressPtrType) ToFrontendPropertiesIPAddressPtrOutput() FrontendPropertiesIPAddressPtrOutput {
	return i.ToFrontendPropertiesIPAddressPtrOutputWithContext(context.Background())
}

func (i *frontendPropertiesIPAddressPtrType) ToFrontendPropertiesIPAddressPtrOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendPropertiesIPAddressPtrOutput)
}

type FrontendPropertiesIPAddressOutput struct{ *pulumi.OutputState }

func (FrontendPropertiesIPAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendPropertiesIPAddress)(nil)).Elem()
}

func (o FrontendPropertiesIPAddressOutput) ToFrontendPropertiesIPAddressOutput() FrontendPropertiesIPAddressOutput {
	return o
}

func (o FrontendPropertiesIPAddressOutput) ToFrontendPropertiesIPAddressOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressOutput {
	return o
}

func (o FrontendPropertiesIPAddressOutput) ToFrontendPropertiesIPAddressPtrOutput() FrontendPropertiesIPAddressPtrOutput {
	return o.ToFrontendPropertiesIPAddressPtrOutputWithContext(context.Background())
}

func (o FrontendPropertiesIPAddressOutput) ToFrontendPropertiesIPAddressPtrOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontendPropertiesIPAddress) *FrontendPropertiesIPAddress {
		return &v
	}).(FrontendPropertiesIPAddressPtrOutput)
}

func (o FrontendPropertiesIPAddressOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendPropertiesIPAddress) string { return v.Id }).(pulumi.StringOutput)
}

type FrontendPropertiesIPAddressPtrOutput struct{ *pulumi.OutputState }

func (FrontendPropertiesIPAddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendPropertiesIPAddress)(nil)).Elem()
}

func (o FrontendPropertiesIPAddressPtrOutput) ToFrontendPropertiesIPAddressPtrOutput() FrontendPropertiesIPAddressPtrOutput {
	return o
}

func (o FrontendPropertiesIPAddressPtrOutput) ToFrontendPropertiesIPAddressPtrOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressPtrOutput {
	return o
}

func (o FrontendPropertiesIPAddressPtrOutput) Elem() FrontendPropertiesIPAddressOutput {
	return o.ApplyT(func(v *FrontendPropertiesIPAddress) FrontendPropertiesIPAddress {
		if v != nil {
			return *v
		}
		var ret FrontendPropertiesIPAddress
		return ret
	}).(FrontendPropertiesIPAddressOutput)
}

func (o FrontendPropertiesIPAddressPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendPropertiesIPAddress) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type FrontendPropertiesIPAddressResponse struct {
	Id string `pulumi:"id"`
}

type FrontendPropertiesIPAddressResponseOutput struct{ *pulumi.OutputState }

func (FrontendPropertiesIPAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendPropertiesIPAddressResponse)(nil)).Elem()
}

func (o FrontendPropertiesIPAddressResponseOutput) ToFrontendPropertiesIPAddressResponseOutput() FrontendPropertiesIPAddressResponseOutput {
	return o
}

func (o FrontendPropertiesIPAddressResponseOutput) ToFrontendPropertiesIPAddressResponseOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressResponseOutput {
	return o
}

func (o FrontendPropertiesIPAddressResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendPropertiesIPAddressResponse) string { return v.Id }).(pulumi.StringOutput)
}

type FrontendPropertiesIPAddressResponsePtrOutput struct{ *pulumi.OutputState }

func (FrontendPropertiesIPAddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendPropertiesIPAddressResponse)(nil)).Elem()
}

func (o FrontendPropertiesIPAddressResponsePtrOutput) ToFrontendPropertiesIPAddressResponsePtrOutput() FrontendPropertiesIPAddressResponsePtrOutput {
	return o
}

func (o FrontendPropertiesIPAddressResponsePtrOutput) ToFrontendPropertiesIPAddressResponsePtrOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressResponsePtrOutput {
	return o
}

func (o FrontendPropertiesIPAddressResponsePtrOutput) Elem() FrontendPropertiesIPAddressResponseOutput {
	return o.ApplyT(func(v *FrontendPropertiesIPAddressResponse) FrontendPropertiesIPAddressResponse {
		if v != nil {
			return *v
		}
		var ret FrontendPropertiesIPAddressResponse
		return ret
	}).(FrontendPropertiesIPAddressResponseOutput)
}

func (o FrontendPropertiesIPAddressResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendPropertiesIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceIDResponse struct {
	Id string `pulumi:"id"`
}

type ResourceIDResponseOutput struct{ *pulumi.OutputState }

func (ResourceIDResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIDResponse)(nil)).Elem()
}

func (o ResourceIDResponseOutput) ToResourceIDResponseOutput() ResourceIDResponseOutput {
	return o
}

func (o ResourceIDResponseOutput) ToResourceIDResponseOutputWithContext(ctx context.Context) ResourceIDResponseOutput {
	return o
}

func (o ResourceIDResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIDResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ResourceIDResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceIDResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceIDResponse)(nil)).Elem()
}

func (o ResourceIDResponseArrayOutput) ToResourceIDResponseArrayOutput() ResourceIDResponseArrayOutput {
	return o
}

func (o ResourceIDResponseArrayOutput) ToResourceIDResponseArrayOutputWithContext(ctx context.Context) ResourceIDResponseArrayOutput {
	return o
}

func (o ResourceIDResponseArrayOutput) Index(i pulumi.IntInput) ResourceIDResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceIDResponse {
		return vs[0].([]ResourceIDResponse)[vs[1].(int)]
	}).(ResourceIDResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssociationSubnetOutput{})
	pulumi.RegisterOutputType(AssociationSubnetPtrOutput{})
	pulumi.RegisterOutputType(AssociationSubnetResponseOutput{})
	pulumi.RegisterOutputType(AssociationSubnetResponsePtrOutput{})
	pulumi.RegisterOutputType(FrontendPropertiesIPAddressOutput{})
	pulumi.RegisterOutputType(FrontendPropertiesIPAddressPtrOutput{})
	pulumi.RegisterOutputType(FrontendPropertiesIPAddressResponseOutput{})
	pulumi.RegisterOutputType(FrontendPropertiesIPAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIDResponseOutput{})
	pulumi.RegisterOutputType(ResourceIDResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
