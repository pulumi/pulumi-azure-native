


package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddressDetails struct {
	ForwardAddress AddressProperties `pulumi:"forwardAddress"`
}





type AddressDetailsInput interface {
	pulumi.Input

	ToAddressDetailsOutput() AddressDetailsOutput
	ToAddressDetailsOutputWithContext(context.Context) AddressDetailsOutput
}

type AddressDetailsArgs struct {
	ForwardAddress AddressPropertiesInput `pulumi:"forwardAddress"`
}

func (AddressDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressDetails)(nil)).Elem()
}

func (i AddressDetailsArgs) ToAddressDetailsOutput() AddressDetailsOutput {
	return i.ToAddressDetailsOutputWithContext(context.Background())
}

func (i AddressDetailsArgs) ToAddressDetailsOutputWithContext(ctx context.Context) AddressDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressDetailsOutput)
}

func (i AddressDetailsArgs) ToAddressDetailsPtrOutput() AddressDetailsPtrOutput {
	return i.ToAddressDetailsPtrOutputWithContext(context.Background())
}

func (i AddressDetailsArgs) ToAddressDetailsPtrOutputWithContext(ctx context.Context) AddressDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressDetailsOutput).ToAddressDetailsPtrOutputWithContext(ctx)
}









type AddressDetailsPtrInput interface {
	pulumi.Input

	ToAddressDetailsPtrOutput() AddressDetailsPtrOutput
	ToAddressDetailsPtrOutputWithContext(context.Context) AddressDetailsPtrOutput
}

type addressDetailsPtrType AddressDetailsArgs

func AddressDetailsPtr(v *AddressDetailsArgs) AddressDetailsPtrInput {
	return (*addressDetailsPtrType)(v)
}

func (*addressDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressDetails)(nil)).Elem()
}

func (i *addressDetailsPtrType) ToAddressDetailsPtrOutput() AddressDetailsPtrOutput {
	return i.ToAddressDetailsPtrOutputWithContext(context.Background())
}

func (i *addressDetailsPtrType) ToAddressDetailsPtrOutputWithContext(ctx context.Context) AddressDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressDetailsPtrOutput)
}

type AddressDetailsOutput struct{ *pulumi.OutputState }

func (AddressDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressDetails)(nil)).Elem()
}

func (o AddressDetailsOutput) ToAddressDetailsOutput() AddressDetailsOutput {
	return o
}

func (o AddressDetailsOutput) ToAddressDetailsOutputWithContext(ctx context.Context) AddressDetailsOutput {
	return o
}

func (o AddressDetailsOutput) ToAddressDetailsPtrOutput() AddressDetailsPtrOutput {
	return o.ToAddressDetailsPtrOutputWithContext(context.Background())
}

func (o AddressDetailsOutput) ToAddressDetailsPtrOutputWithContext(ctx context.Context) AddressDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressDetails) *AddressDetails {
		return &v
	}).(AddressDetailsPtrOutput)
}

func (o AddressDetailsOutput) ForwardAddress() AddressPropertiesOutput {
	return o.ApplyT(func(v AddressDetails) AddressProperties { return v.ForwardAddress }).(AddressPropertiesOutput)
}

type AddressDetailsPtrOutput struct{ *pulumi.OutputState }

func (AddressDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressDetails)(nil)).Elem()
}

func (o AddressDetailsPtrOutput) ToAddressDetailsPtrOutput() AddressDetailsPtrOutput {
	return o
}

func (o AddressDetailsPtrOutput) ToAddressDetailsPtrOutputWithContext(ctx context.Context) AddressDetailsPtrOutput {
	return o
}

func (o AddressDetailsPtrOutput) Elem() AddressDetailsOutput {
	return o.ApplyT(func(v *AddressDetails) AddressDetails {
		if v != nil {
			return *v
		}
		var ret AddressDetails
		return ret
	}).(AddressDetailsOutput)
}

func (o AddressDetailsPtrOutput) ForwardAddress() AddressPropertiesPtrOutput {
	return o.ApplyT(func(v *AddressDetails) *AddressProperties {
		if v == nil {
			return nil
		}
		return &v.ForwardAddress
	}).(AddressPropertiesPtrOutput)
}

type AddressDetailsResponse struct {
	ForwardAddress AddressPropertiesResponse `pulumi:"forwardAddress"`
	ReturnAddress  AddressPropertiesResponse `pulumi:"returnAddress"`
}





type AddressDetailsResponseInput interface {
	pulumi.Input

	ToAddressDetailsResponseOutput() AddressDetailsResponseOutput
	ToAddressDetailsResponseOutputWithContext(context.Context) AddressDetailsResponseOutput
}

type AddressDetailsResponseArgs struct {
	ForwardAddress AddressPropertiesResponseInput `pulumi:"forwardAddress"`
	ReturnAddress  AddressPropertiesResponseInput `pulumi:"returnAddress"`
}

func (AddressDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressDetailsResponse)(nil)).Elem()
}

func (i AddressDetailsResponseArgs) ToAddressDetailsResponseOutput() AddressDetailsResponseOutput {
	return i.ToAddressDetailsResponseOutputWithContext(context.Background())
}

func (i AddressDetailsResponseArgs) ToAddressDetailsResponseOutputWithContext(ctx context.Context) AddressDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressDetailsResponseOutput)
}

func (i AddressDetailsResponseArgs) ToAddressDetailsResponsePtrOutput() AddressDetailsResponsePtrOutput {
	return i.ToAddressDetailsResponsePtrOutputWithContext(context.Background())
}

func (i AddressDetailsResponseArgs) ToAddressDetailsResponsePtrOutputWithContext(ctx context.Context) AddressDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressDetailsResponseOutput).ToAddressDetailsResponsePtrOutputWithContext(ctx)
}









type AddressDetailsResponsePtrInput interface {
	pulumi.Input

	ToAddressDetailsResponsePtrOutput() AddressDetailsResponsePtrOutput
	ToAddressDetailsResponsePtrOutputWithContext(context.Context) AddressDetailsResponsePtrOutput
}

type addressDetailsResponsePtrType AddressDetailsResponseArgs

func AddressDetailsResponsePtr(v *AddressDetailsResponseArgs) AddressDetailsResponsePtrInput {
	return (*addressDetailsResponsePtrType)(v)
}

func (*addressDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressDetailsResponse)(nil)).Elem()
}

func (i *addressDetailsResponsePtrType) ToAddressDetailsResponsePtrOutput() AddressDetailsResponsePtrOutput {
	return i.ToAddressDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *addressDetailsResponsePtrType) ToAddressDetailsResponsePtrOutputWithContext(ctx context.Context) AddressDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressDetailsResponsePtrOutput)
}

type AddressDetailsResponseOutput struct{ *pulumi.OutputState }

func (AddressDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressDetailsResponse)(nil)).Elem()
}

func (o AddressDetailsResponseOutput) ToAddressDetailsResponseOutput() AddressDetailsResponseOutput {
	return o
}

func (o AddressDetailsResponseOutput) ToAddressDetailsResponseOutputWithContext(ctx context.Context) AddressDetailsResponseOutput {
	return o
}

func (o AddressDetailsResponseOutput) ToAddressDetailsResponsePtrOutput() AddressDetailsResponsePtrOutput {
	return o.ToAddressDetailsResponsePtrOutputWithContext(context.Background())
}

func (o AddressDetailsResponseOutput) ToAddressDetailsResponsePtrOutputWithContext(ctx context.Context) AddressDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressDetailsResponse) *AddressDetailsResponse {
		return &v
	}).(AddressDetailsResponsePtrOutput)
}

func (o AddressDetailsResponseOutput) ForwardAddress() AddressPropertiesResponseOutput {
	return o.ApplyT(func(v AddressDetailsResponse) AddressPropertiesResponse { return v.ForwardAddress }).(AddressPropertiesResponseOutput)
}

func (o AddressDetailsResponseOutput) ReturnAddress() AddressPropertiesResponseOutput {
	return o.ApplyT(func(v AddressDetailsResponse) AddressPropertiesResponse { return v.ReturnAddress }).(AddressPropertiesResponseOutput)
}

type AddressDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (AddressDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressDetailsResponse)(nil)).Elem()
}

func (o AddressDetailsResponsePtrOutput) ToAddressDetailsResponsePtrOutput() AddressDetailsResponsePtrOutput {
	return o
}

func (o AddressDetailsResponsePtrOutput) ToAddressDetailsResponsePtrOutputWithContext(ctx context.Context) AddressDetailsResponsePtrOutput {
	return o
}

func (o AddressDetailsResponsePtrOutput) Elem() AddressDetailsResponseOutput {
	return o.ApplyT(func(v *AddressDetailsResponse) AddressDetailsResponse {
		if v != nil {
			return *v
		}
		var ret AddressDetailsResponse
		return ret
	}).(AddressDetailsResponseOutput)
}

func (o AddressDetailsResponsePtrOutput) ForwardAddress() AddressPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AddressDetailsResponse) *AddressPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.ForwardAddress
	}).(AddressPropertiesResponsePtrOutput)
}

func (o AddressDetailsResponsePtrOutput) ReturnAddress() AddressPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AddressDetailsResponse) *AddressPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.ReturnAddress
	}).(AddressPropertiesResponsePtrOutput)
}

type AddressProperties struct {
	ContactDetails  ContactDetails   `pulumi:"contactDetails"`
	ShippingAddress *ShippingAddress `pulumi:"shippingAddress"`
}





type AddressPropertiesInput interface {
	pulumi.Input

	ToAddressPropertiesOutput() AddressPropertiesOutput
	ToAddressPropertiesOutputWithContext(context.Context) AddressPropertiesOutput
}

type AddressPropertiesArgs struct {
	ContactDetails  ContactDetailsInput     `pulumi:"contactDetails"`
	ShippingAddress ShippingAddressPtrInput `pulumi:"shippingAddress"`
}

func (AddressPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressProperties)(nil)).Elem()
}

func (i AddressPropertiesArgs) ToAddressPropertiesOutput() AddressPropertiesOutput {
	return i.ToAddressPropertiesOutputWithContext(context.Background())
}

func (i AddressPropertiesArgs) ToAddressPropertiesOutputWithContext(ctx context.Context) AddressPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPropertiesOutput)
}

func (i AddressPropertiesArgs) ToAddressPropertiesPtrOutput() AddressPropertiesPtrOutput {
	return i.ToAddressPropertiesPtrOutputWithContext(context.Background())
}

func (i AddressPropertiesArgs) ToAddressPropertiesPtrOutputWithContext(ctx context.Context) AddressPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPropertiesOutput).ToAddressPropertiesPtrOutputWithContext(ctx)
}









type AddressPropertiesPtrInput interface {
	pulumi.Input

	ToAddressPropertiesPtrOutput() AddressPropertiesPtrOutput
	ToAddressPropertiesPtrOutputWithContext(context.Context) AddressPropertiesPtrOutput
}

type addressPropertiesPtrType AddressPropertiesArgs

func AddressPropertiesPtr(v *AddressPropertiesArgs) AddressPropertiesPtrInput {
	return (*addressPropertiesPtrType)(v)
}

func (*addressPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressProperties)(nil)).Elem()
}

func (i *addressPropertiesPtrType) ToAddressPropertiesPtrOutput() AddressPropertiesPtrOutput {
	return i.ToAddressPropertiesPtrOutputWithContext(context.Background())
}

func (i *addressPropertiesPtrType) ToAddressPropertiesPtrOutputWithContext(ctx context.Context) AddressPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPropertiesPtrOutput)
}

type AddressPropertiesOutput struct{ *pulumi.OutputState }

func (AddressPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressProperties)(nil)).Elem()
}

func (o AddressPropertiesOutput) ToAddressPropertiesOutput() AddressPropertiesOutput {
	return o
}

func (o AddressPropertiesOutput) ToAddressPropertiesOutputWithContext(ctx context.Context) AddressPropertiesOutput {
	return o
}

func (o AddressPropertiesOutput) ToAddressPropertiesPtrOutput() AddressPropertiesPtrOutput {
	return o.ToAddressPropertiesPtrOutputWithContext(context.Background())
}

func (o AddressPropertiesOutput) ToAddressPropertiesPtrOutputWithContext(ctx context.Context) AddressPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressProperties) *AddressProperties {
		return &v
	}).(AddressPropertiesPtrOutput)
}

func (o AddressPropertiesOutput) ContactDetails() ContactDetailsOutput {
	return o.ApplyT(func(v AddressProperties) ContactDetails { return v.ContactDetails }).(ContactDetailsOutput)
}

func (o AddressPropertiesOutput) ShippingAddress() ShippingAddressPtrOutput {
	return o.ApplyT(func(v AddressProperties) *ShippingAddress { return v.ShippingAddress }).(ShippingAddressPtrOutput)
}

type AddressPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AddressPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressProperties)(nil)).Elem()
}

func (o AddressPropertiesPtrOutput) ToAddressPropertiesPtrOutput() AddressPropertiesPtrOutput {
	return o
}

func (o AddressPropertiesPtrOutput) ToAddressPropertiesPtrOutputWithContext(ctx context.Context) AddressPropertiesPtrOutput {
	return o
}

func (o AddressPropertiesPtrOutput) Elem() AddressPropertiesOutput {
	return o.ApplyT(func(v *AddressProperties) AddressProperties {
		if v != nil {
			return *v
		}
		var ret AddressProperties
		return ret
	}).(AddressPropertiesOutput)
}

func (o AddressPropertiesPtrOutput) ContactDetails() ContactDetailsPtrOutput {
	return o.ApplyT(func(v *AddressProperties) *ContactDetails {
		if v == nil {
			return nil
		}
		return &v.ContactDetails
	}).(ContactDetailsPtrOutput)
}

func (o AddressPropertiesPtrOutput) ShippingAddress() ShippingAddressPtrOutput {
	return o.ApplyT(func(v *AddressProperties) *ShippingAddress {
		if v == nil {
			return nil
		}
		return v.ShippingAddress
	}).(ShippingAddressPtrOutput)
}

type AddressPropertiesResponse struct {
	ContactDetails  ContactDetailsResponse   `pulumi:"contactDetails"`
	ShippingAddress *ShippingAddressResponse `pulumi:"shippingAddress"`
}





type AddressPropertiesResponseInput interface {
	pulumi.Input

	ToAddressPropertiesResponseOutput() AddressPropertiesResponseOutput
	ToAddressPropertiesResponseOutputWithContext(context.Context) AddressPropertiesResponseOutput
}

type AddressPropertiesResponseArgs struct {
	ContactDetails  ContactDetailsResponseInput     `pulumi:"contactDetails"`
	ShippingAddress ShippingAddressResponsePtrInput `pulumi:"shippingAddress"`
}

func (AddressPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressPropertiesResponse)(nil)).Elem()
}

func (i AddressPropertiesResponseArgs) ToAddressPropertiesResponseOutput() AddressPropertiesResponseOutput {
	return i.ToAddressPropertiesResponseOutputWithContext(context.Background())
}

func (i AddressPropertiesResponseArgs) ToAddressPropertiesResponseOutputWithContext(ctx context.Context) AddressPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPropertiesResponseOutput)
}

func (i AddressPropertiesResponseArgs) ToAddressPropertiesResponsePtrOutput() AddressPropertiesResponsePtrOutput {
	return i.ToAddressPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AddressPropertiesResponseArgs) ToAddressPropertiesResponsePtrOutputWithContext(ctx context.Context) AddressPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPropertiesResponseOutput).ToAddressPropertiesResponsePtrOutputWithContext(ctx)
}









type AddressPropertiesResponsePtrInput interface {
	pulumi.Input

	ToAddressPropertiesResponsePtrOutput() AddressPropertiesResponsePtrOutput
	ToAddressPropertiesResponsePtrOutputWithContext(context.Context) AddressPropertiesResponsePtrOutput
}

type addressPropertiesResponsePtrType AddressPropertiesResponseArgs

func AddressPropertiesResponsePtr(v *AddressPropertiesResponseArgs) AddressPropertiesResponsePtrInput {
	return (*addressPropertiesResponsePtrType)(v)
}

func (*addressPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressPropertiesResponse)(nil)).Elem()
}

func (i *addressPropertiesResponsePtrType) ToAddressPropertiesResponsePtrOutput() AddressPropertiesResponsePtrOutput {
	return i.ToAddressPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *addressPropertiesResponsePtrType) ToAddressPropertiesResponsePtrOutputWithContext(ctx context.Context) AddressPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPropertiesResponsePtrOutput)
}

type AddressPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AddressPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressPropertiesResponse)(nil)).Elem()
}

func (o AddressPropertiesResponseOutput) ToAddressPropertiesResponseOutput() AddressPropertiesResponseOutput {
	return o
}

func (o AddressPropertiesResponseOutput) ToAddressPropertiesResponseOutputWithContext(ctx context.Context) AddressPropertiesResponseOutput {
	return o
}

func (o AddressPropertiesResponseOutput) ToAddressPropertiesResponsePtrOutput() AddressPropertiesResponsePtrOutput {
	return o.ToAddressPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AddressPropertiesResponseOutput) ToAddressPropertiesResponsePtrOutputWithContext(ctx context.Context) AddressPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressPropertiesResponse) *AddressPropertiesResponse {
		return &v
	}).(AddressPropertiesResponsePtrOutput)
}

func (o AddressPropertiesResponseOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v AddressPropertiesResponse) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o AddressPropertiesResponseOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v AddressPropertiesResponse) *ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponsePtrOutput)
}

type AddressPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AddressPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressPropertiesResponse)(nil)).Elem()
}

func (o AddressPropertiesResponsePtrOutput) ToAddressPropertiesResponsePtrOutput() AddressPropertiesResponsePtrOutput {
	return o
}

func (o AddressPropertiesResponsePtrOutput) ToAddressPropertiesResponsePtrOutputWithContext(ctx context.Context) AddressPropertiesResponsePtrOutput {
	return o
}

func (o AddressPropertiesResponsePtrOutput) Elem() AddressPropertiesResponseOutput {
	return o.ApplyT(func(v *AddressPropertiesResponse) AddressPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AddressPropertiesResponse
		return ret
	}).(AddressPropertiesResponseOutput)
}

func (o AddressPropertiesResponsePtrOutput) ContactDetails() ContactDetailsResponsePtrOutput {
	return o.ApplyT(func(v *AddressPropertiesResponse) *ContactDetailsResponse {
		if v == nil {
			return nil
		}
		return &v.ContactDetails
	}).(ContactDetailsResponsePtrOutput)
}

func (o AddressPropertiesResponsePtrOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v *AddressPropertiesResponse) *ShippingAddressResponse {
		if v == nil {
			return nil
		}
		return v.ShippingAddress
	}).(ShippingAddressResponsePtrOutput)
}

type AvailabilityInformationResponse struct {
	AvailabilityStage     string `pulumi:"availabilityStage"`
	DisabledReason        string `pulumi:"disabledReason"`
	DisabledReasonMessage string `pulumi:"disabledReasonMessage"`
}





type AvailabilityInformationResponseInput interface {
	pulumi.Input

	ToAvailabilityInformationResponseOutput() AvailabilityInformationResponseOutput
	ToAvailabilityInformationResponseOutputWithContext(context.Context) AvailabilityInformationResponseOutput
}

type AvailabilityInformationResponseArgs struct {
	AvailabilityStage     pulumi.StringInput `pulumi:"availabilityStage"`
	DisabledReason        pulumi.StringInput `pulumi:"disabledReason"`
	DisabledReasonMessage pulumi.StringInput `pulumi:"disabledReasonMessage"`
}

func (AvailabilityInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityInformationResponse)(nil)).Elem()
}

func (i AvailabilityInformationResponseArgs) ToAvailabilityInformationResponseOutput() AvailabilityInformationResponseOutput {
	return i.ToAvailabilityInformationResponseOutputWithContext(context.Background())
}

func (i AvailabilityInformationResponseArgs) ToAvailabilityInformationResponseOutputWithContext(ctx context.Context) AvailabilityInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityInformationResponseOutput)
}

type AvailabilityInformationResponseOutput struct{ *pulumi.OutputState }

func (AvailabilityInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityInformationResponse)(nil)).Elem()
}

func (o AvailabilityInformationResponseOutput) ToAvailabilityInformationResponseOutput() AvailabilityInformationResponseOutput {
	return o
}

func (o AvailabilityInformationResponseOutput) ToAvailabilityInformationResponseOutputWithContext(ctx context.Context) AvailabilityInformationResponseOutput {
	return o
}

func (o AvailabilityInformationResponseOutput) AvailabilityStage() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilityInformationResponse) string { return v.AvailabilityStage }).(pulumi.StringOutput)
}

func (o AvailabilityInformationResponseOutput) DisabledReason() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilityInformationResponse) string { return v.DisabledReason }).(pulumi.StringOutput)
}

func (o AvailabilityInformationResponseOutput) DisabledReasonMessage() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilityInformationResponse) string { return v.DisabledReasonMessage }).(pulumi.StringOutput)
}

type BillingMeterDetailsResponse struct {
	Frequency    string      `pulumi:"frequency"`
	MeterDetails interface{} `pulumi:"meterDetails"`
	MeteringType string      `pulumi:"meteringType"`
	Name         string      `pulumi:"name"`
}





type BillingMeterDetailsResponseInput interface {
	pulumi.Input

	ToBillingMeterDetailsResponseOutput() BillingMeterDetailsResponseOutput
	ToBillingMeterDetailsResponseOutputWithContext(context.Context) BillingMeterDetailsResponseOutput
}

type BillingMeterDetailsResponseArgs struct {
	Frequency    pulumi.StringInput `pulumi:"frequency"`
	MeterDetails pulumi.Input       `pulumi:"meterDetails"`
	MeteringType pulumi.StringInput `pulumi:"meteringType"`
	Name         pulumi.StringInput `pulumi:"name"`
}

func (BillingMeterDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingMeterDetailsResponse)(nil)).Elem()
}

func (i BillingMeterDetailsResponseArgs) ToBillingMeterDetailsResponseOutput() BillingMeterDetailsResponseOutput {
	return i.ToBillingMeterDetailsResponseOutputWithContext(context.Background())
}

func (i BillingMeterDetailsResponseArgs) ToBillingMeterDetailsResponseOutputWithContext(ctx context.Context) BillingMeterDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingMeterDetailsResponseOutput)
}





type BillingMeterDetailsResponseArrayInput interface {
	pulumi.Input

	ToBillingMeterDetailsResponseArrayOutput() BillingMeterDetailsResponseArrayOutput
	ToBillingMeterDetailsResponseArrayOutputWithContext(context.Context) BillingMeterDetailsResponseArrayOutput
}

type BillingMeterDetailsResponseArray []BillingMeterDetailsResponseInput

func (BillingMeterDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BillingMeterDetailsResponse)(nil)).Elem()
}

func (i BillingMeterDetailsResponseArray) ToBillingMeterDetailsResponseArrayOutput() BillingMeterDetailsResponseArrayOutput {
	return i.ToBillingMeterDetailsResponseArrayOutputWithContext(context.Background())
}

func (i BillingMeterDetailsResponseArray) ToBillingMeterDetailsResponseArrayOutputWithContext(ctx context.Context) BillingMeterDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingMeterDetailsResponseArrayOutput)
}

type BillingMeterDetailsResponseOutput struct{ *pulumi.OutputState }

func (BillingMeterDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingMeterDetailsResponse)(nil)).Elem()
}

func (o BillingMeterDetailsResponseOutput) ToBillingMeterDetailsResponseOutput() BillingMeterDetailsResponseOutput {
	return o
}

func (o BillingMeterDetailsResponseOutput) ToBillingMeterDetailsResponseOutputWithContext(ctx context.Context) BillingMeterDetailsResponseOutput {
	return o
}

func (o BillingMeterDetailsResponseOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v BillingMeterDetailsResponse) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o BillingMeterDetailsResponseOutput) MeterDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v BillingMeterDetailsResponse) interface{} { return v.MeterDetails }).(pulumi.AnyOutput)
}

func (o BillingMeterDetailsResponseOutput) MeteringType() pulumi.StringOutput {
	return o.ApplyT(func(v BillingMeterDetailsResponse) string { return v.MeteringType }).(pulumi.StringOutput)
}

func (o BillingMeterDetailsResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BillingMeterDetailsResponse) string { return v.Name }).(pulumi.StringOutput)
}

type BillingMeterDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (BillingMeterDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BillingMeterDetailsResponse)(nil)).Elem()
}

func (o BillingMeterDetailsResponseArrayOutput) ToBillingMeterDetailsResponseArrayOutput() BillingMeterDetailsResponseArrayOutput {
	return o
}

func (o BillingMeterDetailsResponseArrayOutput) ToBillingMeterDetailsResponseArrayOutputWithContext(ctx context.Context) BillingMeterDetailsResponseArrayOutput {
	return o
}

func (o BillingMeterDetailsResponseArrayOutput) Index(i pulumi.IntInput) BillingMeterDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BillingMeterDetailsResponse {
		return vs[0].([]BillingMeterDetailsResponse)[vs[1].(int)]
	}).(BillingMeterDetailsResponseOutput)
}

type ConfigurationFilters struct {
	FilterableProperty   []FilterableProperty `pulumi:"filterableProperty"`
	HierarchyInformation HierarchyInformation `pulumi:"hierarchyInformation"`
}





type ConfigurationFiltersInput interface {
	pulumi.Input

	ToConfigurationFiltersOutput() ConfigurationFiltersOutput
	ToConfigurationFiltersOutputWithContext(context.Context) ConfigurationFiltersOutput
}

type ConfigurationFiltersArgs struct {
	FilterableProperty   FilterablePropertyArrayInput `pulumi:"filterableProperty"`
	HierarchyInformation HierarchyInformationInput    `pulumi:"hierarchyInformation"`
}

func (ConfigurationFiltersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationFilters)(nil)).Elem()
}

func (i ConfigurationFiltersArgs) ToConfigurationFiltersOutput() ConfigurationFiltersOutput {
	return i.ToConfigurationFiltersOutputWithContext(context.Background())
}

func (i ConfigurationFiltersArgs) ToConfigurationFiltersOutputWithContext(ctx context.Context) ConfigurationFiltersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationFiltersOutput)
}





type ConfigurationFiltersArrayInput interface {
	pulumi.Input

	ToConfigurationFiltersArrayOutput() ConfigurationFiltersArrayOutput
	ToConfigurationFiltersArrayOutputWithContext(context.Context) ConfigurationFiltersArrayOutput
}

type ConfigurationFiltersArray []ConfigurationFiltersInput

func (ConfigurationFiltersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationFilters)(nil)).Elem()
}

func (i ConfigurationFiltersArray) ToConfigurationFiltersArrayOutput() ConfigurationFiltersArrayOutput {
	return i.ToConfigurationFiltersArrayOutputWithContext(context.Background())
}

func (i ConfigurationFiltersArray) ToConfigurationFiltersArrayOutputWithContext(ctx context.Context) ConfigurationFiltersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationFiltersArrayOutput)
}

type ConfigurationFiltersOutput struct{ *pulumi.OutputState }

func (ConfigurationFiltersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationFilters)(nil)).Elem()
}

func (o ConfigurationFiltersOutput) ToConfigurationFiltersOutput() ConfigurationFiltersOutput {
	return o
}

func (o ConfigurationFiltersOutput) ToConfigurationFiltersOutputWithContext(ctx context.Context) ConfigurationFiltersOutput {
	return o
}

func (o ConfigurationFiltersOutput) FilterableProperty() FilterablePropertyArrayOutput {
	return o.ApplyT(func(v ConfigurationFilters) []FilterableProperty { return v.FilterableProperty }).(FilterablePropertyArrayOutput)
}

func (o ConfigurationFiltersOutput) HierarchyInformation() HierarchyInformationOutput {
	return o.ApplyT(func(v ConfigurationFilters) HierarchyInformation { return v.HierarchyInformation }).(HierarchyInformationOutput)
}

type ConfigurationFiltersArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationFiltersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationFilters)(nil)).Elem()
}

func (o ConfigurationFiltersArrayOutput) ToConfigurationFiltersArrayOutput() ConfigurationFiltersArrayOutput {
	return o
}

func (o ConfigurationFiltersArrayOutput) ToConfigurationFiltersArrayOutputWithContext(ctx context.Context) ConfigurationFiltersArrayOutput {
	return o
}

func (o ConfigurationFiltersArrayOutput) Index(i pulumi.IntInput) ConfigurationFiltersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationFilters {
		return vs[0].([]ConfigurationFilters)[vs[1].(int)]
	}).(ConfigurationFiltersOutput)
}

type ConfigurationResponse struct {
	AvailabilityInformation AvailabilityInformationResponse `pulumi:"availabilityInformation"`
	CostInformation         CostInformationResponse         `pulumi:"costInformation"`
	Description             DescriptionResponse             `pulumi:"description"`
	Dimensions              DimensionsResponse              `pulumi:"dimensions"`
	DisplayName             string                          `pulumi:"displayName"`
	FilterableProperties    []FilterablePropertyResponse    `pulumi:"filterableProperties"`
	HierarchyInformation    HierarchyInformationResponse    `pulumi:"hierarchyInformation"`
	ImageInformation        []ImageInformationResponse      `pulumi:"imageInformation"`
	Specifications          []SpecificationResponse         `pulumi:"specifications"`
}





type ConfigurationResponseInput interface {
	pulumi.Input

	ToConfigurationResponseOutput() ConfigurationResponseOutput
	ToConfigurationResponseOutputWithContext(context.Context) ConfigurationResponseOutput
}

type ConfigurationResponseArgs struct {
	AvailabilityInformation AvailabilityInformationResponseInput `pulumi:"availabilityInformation"`
	CostInformation         CostInformationResponseInput         `pulumi:"costInformation"`
	Description             DescriptionResponseInput             `pulumi:"description"`
	Dimensions              DimensionsResponseInput              `pulumi:"dimensions"`
	DisplayName             pulumi.StringInput                   `pulumi:"displayName"`
	FilterableProperties    FilterablePropertyResponseArrayInput `pulumi:"filterableProperties"`
	HierarchyInformation    HierarchyInformationResponseInput    `pulumi:"hierarchyInformation"`
	ImageInformation        ImageInformationResponseArrayInput   `pulumi:"imageInformation"`
	Specifications          SpecificationResponseArrayInput      `pulumi:"specifications"`
}

func (ConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationResponse)(nil)).Elem()
}

func (i ConfigurationResponseArgs) ToConfigurationResponseOutput() ConfigurationResponseOutput {
	return i.ToConfigurationResponseOutputWithContext(context.Background())
}

func (i ConfigurationResponseArgs) ToConfigurationResponseOutputWithContext(ctx context.Context) ConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationResponseOutput)
}





type ConfigurationResponseArrayInput interface {
	pulumi.Input

	ToConfigurationResponseArrayOutput() ConfigurationResponseArrayOutput
	ToConfigurationResponseArrayOutputWithContext(context.Context) ConfigurationResponseArrayOutput
}

type ConfigurationResponseArray []ConfigurationResponseInput

func (ConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationResponse)(nil)).Elem()
}

func (i ConfigurationResponseArray) ToConfigurationResponseArrayOutput() ConfigurationResponseArrayOutput {
	return i.ToConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i ConfigurationResponseArray) ToConfigurationResponseArrayOutputWithContext(ctx context.Context) ConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationResponseArrayOutput)
}

type ConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationResponse)(nil)).Elem()
}

func (o ConfigurationResponseOutput) ToConfigurationResponseOutput() ConfigurationResponseOutput {
	return o
}

func (o ConfigurationResponseOutput) ToConfigurationResponseOutputWithContext(ctx context.Context) ConfigurationResponseOutput {
	return o
}

func (o ConfigurationResponseOutput) AvailabilityInformation() AvailabilityInformationResponseOutput {
	return o.ApplyT(func(v ConfigurationResponse) AvailabilityInformationResponse { return v.AvailabilityInformation }).(AvailabilityInformationResponseOutput)
}

func (o ConfigurationResponseOutput) CostInformation() CostInformationResponseOutput {
	return o.ApplyT(func(v ConfigurationResponse) CostInformationResponse { return v.CostInformation }).(CostInformationResponseOutput)
}

func (o ConfigurationResponseOutput) Description() DescriptionResponseOutput {
	return o.ApplyT(func(v ConfigurationResponse) DescriptionResponse { return v.Description }).(DescriptionResponseOutput)
}

func (o ConfigurationResponseOutput) Dimensions() DimensionsResponseOutput {
	return o.ApplyT(func(v ConfigurationResponse) DimensionsResponse { return v.Dimensions }).(DimensionsResponseOutput)
}

func (o ConfigurationResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ConfigurationResponseOutput) FilterableProperties() FilterablePropertyResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationResponse) []FilterablePropertyResponse { return v.FilterableProperties }).(FilterablePropertyResponseArrayOutput)
}

func (o ConfigurationResponseOutput) HierarchyInformation() HierarchyInformationResponseOutput {
	return o.ApplyT(func(v ConfigurationResponse) HierarchyInformationResponse { return v.HierarchyInformation }).(HierarchyInformationResponseOutput)
}

func (o ConfigurationResponseOutput) ImageInformation() ImageInformationResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationResponse) []ImageInformationResponse { return v.ImageInformation }).(ImageInformationResponseArrayOutput)
}

func (o ConfigurationResponseOutput) Specifications() SpecificationResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationResponse) []SpecificationResponse { return v.Specifications }).(SpecificationResponseArrayOutput)
}

type ConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationResponse)(nil)).Elem()
}

func (o ConfigurationResponseArrayOutput) ToConfigurationResponseArrayOutput() ConfigurationResponseArrayOutput {
	return o
}

func (o ConfigurationResponseArrayOutput) ToConfigurationResponseArrayOutputWithContext(ctx context.Context) ConfigurationResponseArrayOutput {
	return o
}

func (o ConfigurationResponseArrayOutput) Index(i pulumi.IntInput) ConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationResponse {
		return vs[0].([]ConfigurationResponse)[vs[1].(int)]
	}).(ConfigurationResponseOutput)
}

type ContactDetails struct {
	ContactName    string   `pulumi:"contactName"`
	EmailList      []string `pulumi:"emailList"`
	Mobile         *string  `pulumi:"mobile"`
	Phone          string   `pulumi:"phone"`
	PhoneExtension *string  `pulumi:"phoneExtension"`
}





type ContactDetailsInput interface {
	pulumi.Input

	ToContactDetailsOutput() ContactDetailsOutput
	ToContactDetailsOutputWithContext(context.Context) ContactDetailsOutput
}

type ContactDetailsArgs struct {
	ContactName    pulumi.StringInput      `pulumi:"contactName"`
	EmailList      pulumi.StringArrayInput `pulumi:"emailList"`
	Mobile         pulumi.StringPtrInput   `pulumi:"mobile"`
	Phone          pulumi.StringInput      `pulumi:"phone"`
	PhoneExtension pulumi.StringPtrInput   `pulumi:"phoneExtension"`
}

func (ContactDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetails)(nil)).Elem()
}

func (i ContactDetailsArgs) ToContactDetailsOutput() ContactDetailsOutput {
	return i.ToContactDetailsOutputWithContext(context.Background())
}

func (i ContactDetailsArgs) ToContactDetailsOutputWithContext(ctx context.Context) ContactDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsOutput)
}

func (i ContactDetailsArgs) ToContactDetailsPtrOutput() ContactDetailsPtrOutput {
	return i.ToContactDetailsPtrOutputWithContext(context.Background())
}

func (i ContactDetailsArgs) ToContactDetailsPtrOutputWithContext(ctx context.Context) ContactDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsOutput).ToContactDetailsPtrOutputWithContext(ctx)
}









type ContactDetailsPtrInput interface {
	pulumi.Input

	ToContactDetailsPtrOutput() ContactDetailsPtrOutput
	ToContactDetailsPtrOutputWithContext(context.Context) ContactDetailsPtrOutput
}

type contactDetailsPtrType ContactDetailsArgs

func ContactDetailsPtr(v *ContactDetailsArgs) ContactDetailsPtrInput {
	return (*contactDetailsPtrType)(v)
}

func (*contactDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactDetails)(nil)).Elem()
}

func (i *contactDetailsPtrType) ToContactDetailsPtrOutput() ContactDetailsPtrOutput {
	return i.ToContactDetailsPtrOutputWithContext(context.Background())
}

func (i *contactDetailsPtrType) ToContactDetailsPtrOutputWithContext(ctx context.Context) ContactDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsPtrOutput)
}

type ContactDetailsOutput struct{ *pulumi.OutputState }

func (ContactDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetails)(nil)).Elem()
}

func (o ContactDetailsOutput) ToContactDetailsOutput() ContactDetailsOutput {
	return o
}

func (o ContactDetailsOutput) ToContactDetailsOutputWithContext(ctx context.Context) ContactDetailsOutput {
	return o
}

func (o ContactDetailsOutput) ToContactDetailsPtrOutput() ContactDetailsPtrOutput {
	return o.ToContactDetailsPtrOutputWithContext(context.Background())
}

func (o ContactDetailsOutput) ToContactDetailsPtrOutputWithContext(ctx context.Context) ContactDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContactDetails) *ContactDetails {
		return &v
	}).(ContactDetailsPtrOutput)
}

func (o ContactDetailsOutput) ContactName() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetails) string { return v.ContactName }).(pulumi.StringOutput)
}

func (o ContactDetailsOutput) EmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactDetails) []string { return v.EmailList }).(pulumi.StringArrayOutput)
}

func (o ContactDetailsOutput) Mobile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetails) *string { return v.Mobile }).(pulumi.StringPtrOutput)
}

func (o ContactDetailsOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetails) string { return v.Phone }).(pulumi.StringOutput)
}

func (o ContactDetailsOutput) PhoneExtension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetails) *string { return v.PhoneExtension }).(pulumi.StringPtrOutput)
}

type ContactDetailsPtrOutput struct{ *pulumi.OutputState }

func (ContactDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactDetails)(nil)).Elem()
}

func (o ContactDetailsPtrOutput) ToContactDetailsPtrOutput() ContactDetailsPtrOutput {
	return o
}

func (o ContactDetailsPtrOutput) ToContactDetailsPtrOutputWithContext(ctx context.Context) ContactDetailsPtrOutput {
	return o
}

func (o ContactDetailsPtrOutput) Elem() ContactDetailsOutput {
	return o.ApplyT(func(v *ContactDetails) ContactDetails {
		if v != nil {
			return *v
		}
		var ret ContactDetails
		return ret
	}).(ContactDetailsOutput)
}

func (o ContactDetailsPtrOutput) ContactName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetails) *string {
		if v == nil {
			return nil
		}
		return &v.ContactName
	}).(pulumi.StringPtrOutput)
}

func (o ContactDetailsPtrOutput) EmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactDetails) []string {
		if v == nil {
			return nil
		}
		return v.EmailList
	}).(pulumi.StringArrayOutput)
}

func (o ContactDetailsPtrOutput) Mobile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetails) *string {
		if v == nil {
			return nil
		}
		return v.Mobile
	}).(pulumi.StringPtrOutput)
}

func (o ContactDetailsPtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetails) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

func (o ContactDetailsPtrOutput) PhoneExtension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetails) *string {
		if v == nil {
			return nil
		}
		return v.PhoneExtension
	}).(pulumi.StringPtrOutput)
}

type ContactDetailsResponse struct {
	ContactName    string   `pulumi:"contactName"`
	EmailList      []string `pulumi:"emailList"`
	Mobile         *string  `pulumi:"mobile"`
	Phone          string   `pulumi:"phone"`
	PhoneExtension *string  `pulumi:"phoneExtension"`
}





type ContactDetailsResponseInput interface {
	pulumi.Input

	ToContactDetailsResponseOutput() ContactDetailsResponseOutput
	ToContactDetailsResponseOutputWithContext(context.Context) ContactDetailsResponseOutput
}

type ContactDetailsResponseArgs struct {
	ContactName    pulumi.StringInput      `pulumi:"contactName"`
	EmailList      pulumi.StringArrayInput `pulumi:"emailList"`
	Mobile         pulumi.StringPtrInput   `pulumi:"mobile"`
	Phone          pulumi.StringInput      `pulumi:"phone"`
	PhoneExtension pulumi.StringPtrInput   `pulumi:"phoneExtension"`
}

func (ContactDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetailsResponse)(nil)).Elem()
}

func (i ContactDetailsResponseArgs) ToContactDetailsResponseOutput() ContactDetailsResponseOutput {
	return i.ToContactDetailsResponseOutputWithContext(context.Background())
}

func (i ContactDetailsResponseArgs) ToContactDetailsResponseOutputWithContext(ctx context.Context) ContactDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsResponseOutput)
}

func (i ContactDetailsResponseArgs) ToContactDetailsResponsePtrOutput() ContactDetailsResponsePtrOutput {
	return i.ToContactDetailsResponsePtrOutputWithContext(context.Background())
}

func (i ContactDetailsResponseArgs) ToContactDetailsResponsePtrOutputWithContext(ctx context.Context) ContactDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsResponseOutput).ToContactDetailsResponsePtrOutputWithContext(ctx)
}









type ContactDetailsResponsePtrInput interface {
	pulumi.Input

	ToContactDetailsResponsePtrOutput() ContactDetailsResponsePtrOutput
	ToContactDetailsResponsePtrOutputWithContext(context.Context) ContactDetailsResponsePtrOutput
}

type contactDetailsResponsePtrType ContactDetailsResponseArgs

func ContactDetailsResponsePtr(v *ContactDetailsResponseArgs) ContactDetailsResponsePtrInput {
	return (*contactDetailsResponsePtrType)(v)
}

func (*contactDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactDetailsResponse)(nil)).Elem()
}

func (i *contactDetailsResponsePtrType) ToContactDetailsResponsePtrOutput() ContactDetailsResponsePtrOutput {
	return i.ToContactDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *contactDetailsResponsePtrType) ToContactDetailsResponsePtrOutputWithContext(ctx context.Context) ContactDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsResponsePtrOutput)
}

type ContactDetailsResponseOutput struct{ *pulumi.OutputState }

func (ContactDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetailsResponse)(nil)).Elem()
}

func (o ContactDetailsResponseOutput) ToContactDetailsResponseOutput() ContactDetailsResponseOutput {
	return o
}

func (o ContactDetailsResponseOutput) ToContactDetailsResponseOutputWithContext(ctx context.Context) ContactDetailsResponseOutput {
	return o
}

func (o ContactDetailsResponseOutput) ToContactDetailsResponsePtrOutput() ContactDetailsResponsePtrOutput {
	return o.ToContactDetailsResponsePtrOutputWithContext(context.Background())
}

func (o ContactDetailsResponseOutput) ToContactDetailsResponsePtrOutputWithContext(ctx context.Context) ContactDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContactDetailsResponse) *ContactDetailsResponse {
		return &v
	}).(ContactDetailsResponsePtrOutput)
}

func (o ContactDetailsResponseOutput) ContactName() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetailsResponse) string { return v.ContactName }).(pulumi.StringOutput)
}

func (o ContactDetailsResponseOutput) EmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactDetailsResponse) []string { return v.EmailList }).(pulumi.StringArrayOutput)
}

func (o ContactDetailsResponseOutput) Mobile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailsResponse) *string { return v.Mobile }).(pulumi.StringPtrOutput)
}

func (o ContactDetailsResponseOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetailsResponse) string { return v.Phone }).(pulumi.StringOutput)
}

func (o ContactDetailsResponseOutput) PhoneExtension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailsResponse) *string { return v.PhoneExtension }).(pulumi.StringPtrOutput)
}

type ContactDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ContactDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactDetailsResponse)(nil)).Elem()
}

func (o ContactDetailsResponsePtrOutput) ToContactDetailsResponsePtrOutput() ContactDetailsResponsePtrOutput {
	return o
}

func (o ContactDetailsResponsePtrOutput) ToContactDetailsResponsePtrOutputWithContext(ctx context.Context) ContactDetailsResponsePtrOutput {
	return o
}

func (o ContactDetailsResponsePtrOutput) Elem() ContactDetailsResponseOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) ContactDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ContactDetailsResponse
		return ret
	}).(ContactDetailsResponseOutput)
}

func (o ContactDetailsResponsePtrOutput) ContactName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ContactName
	}).(pulumi.StringPtrOutput)
}

func (o ContactDetailsResponsePtrOutput) EmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.EmailList
	}).(pulumi.StringArrayOutput)
}

func (o ContactDetailsResponsePtrOutput) Mobile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mobile
	}).(pulumi.StringPtrOutput)
}

func (o ContactDetailsResponsePtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

func (o ContactDetailsResponsePtrOutput) PhoneExtension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PhoneExtension
	}).(pulumi.StringPtrOutput)
}

type CostInformationResponse struct {
	BillingInfoUrl      string                        `pulumi:"billingInfoUrl"`
	BillingMeterDetails []BillingMeterDetailsResponse `pulumi:"billingMeterDetails"`
}





type CostInformationResponseInput interface {
	pulumi.Input

	ToCostInformationResponseOutput() CostInformationResponseOutput
	ToCostInformationResponseOutputWithContext(context.Context) CostInformationResponseOutput
}

type CostInformationResponseArgs struct {
	BillingInfoUrl      pulumi.StringInput                    `pulumi:"billingInfoUrl"`
	BillingMeterDetails BillingMeterDetailsResponseArrayInput `pulumi:"billingMeterDetails"`
}

func (CostInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CostInformationResponse)(nil)).Elem()
}

func (i CostInformationResponseArgs) ToCostInformationResponseOutput() CostInformationResponseOutput {
	return i.ToCostInformationResponseOutputWithContext(context.Background())
}

func (i CostInformationResponseArgs) ToCostInformationResponseOutputWithContext(ctx context.Context) CostInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostInformationResponseOutput)
}

type CostInformationResponseOutput struct{ *pulumi.OutputState }

func (CostInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostInformationResponse)(nil)).Elem()
}

func (o CostInformationResponseOutput) ToCostInformationResponseOutput() CostInformationResponseOutput {
	return o
}

func (o CostInformationResponseOutput) ToCostInformationResponseOutputWithContext(ctx context.Context) CostInformationResponseOutput {
	return o
}

func (o CostInformationResponseOutput) BillingInfoUrl() pulumi.StringOutput {
	return o.ApplyT(func(v CostInformationResponse) string { return v.BillingInfoUrl }).(pulumi.StringOutput)
}

func (o CostInformationResponseOutput) BillingMeterDetails() BillingMeterDetailsResponseArrayOutput {
	return o.ApplyT(func(v CostInformationResponse) []BillingMeterDetailsResponse { return v.BillingMeterDetails }).(BillingMeterDetailsResponseArrayOutput)
}

type CustomerSubscriptionDetails struct {
	LocationPlacementId *string                                  `pulumi:"locationPlacementId"`
	QuotaId             string                                   `pulumi:"quotaId"`
	RegisteredFeatures  []CustomerSubscriptionRegisteredFeatures `pulumi:"registeredFeatures"`
}





type CustomerSubscriptionDetailsInput interface {
	pulumi.Input

	ToCustomerSubscriptionDetailsOutput() CustomerSubscriptionDetailsOutput
	ToCustomerSubscriptionDetailsOutputWithContext(context.Context) CustomerSubscriptionDetailsOutput
}

type CustomerSubscriptionDetailsArgs struct {
	LocationPlacementId pulumi.StringPtrInput                            `pulumi:"locationPlacementId"`
	QuotaId             pulumi.StringInput                               `pulumi:"quotaId"`
	RegisteredFeatures  CustomerSubscriptionRegisteredFeaturesArrayInput `pulumi:"registeredFeatures"`
}

func (CustomerSubscriptionDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSubscriptionDetails)(nil)).Elem()
}

func (i CustomerSubscriptionDetailsArgs) ToCustomerSubscriptionDetailsOutput() CustomerSubscriptionDetailsOutput {
	return i.ToCustomerSubscriptionDetailsOutputWithContext(context.Background())
}

func (i CustomerSubscriptionDetailsArgs) ToCustomerSubscriptionDetailsOutputWithContext(ctx context.Context) CustomerSubscriptionDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionDetailsOutput)
}

func (i CustomerSubscriptionDetailsArgs) ToCustomerSubscriptionDetailsPtrOutput() CustomerSubscriptionDetailsPtrOutput {
	return i.ToCustomerSubscriptionDetailsPtrOutputWithContext(context.Background())
}

func (i CustomerSubscriptionDetailsArgs) ToCustomerSubscriptionDetailsPtrOutputWithContext(ctx context.Context) CustomerSubscriptionDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionDetailsOutput).ToCustomerSubscriptionDetailsPtrOutputWithContext(ctx)
}









type CustomerSubscriptionDetailsPtrInput interface {
	pulumi.Input

	ToCustomerSubscriptionDetailsPtrOutput() CustomerSubscriptionDetailsPtrOutput
	ToCustomerSubscriptionDetailsPtrOutputWithContext(context.Context) CustomerSubscriptionDetailsPtrOutput
}

type customerSubscriptionDetailsPtrType CustomerSubscriptionDetailsArgs

func CustomerSubscriptionDetailsPtr(v *CustomerSubscriptionDetailsArgs) CustomerSubscriptionDetailsPtrInput {
	return (*customerSubscriptionDetailsPtrType)(v)
}

func (*customerSubscriptionDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerSubscriptionDetails)(nil)).Elem()
}

func (i *customerSubscriptionDetailsPtrType) ToCustomerSubscriptionDetailsPtrOutput() CustomerSubscriptionDetailsPtrOutput {
	return i.ToCustomerSubscriptionDetailsPtrOutputWithContext(context.Background())
}

func (i *customerSubscriptionDetailsPtrType) ToCustomerSubscriptionDetailsPtrOutputWithContext(ctx context.Context) CustomerSubscriptionDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionDetailsPtrOutput)
}

type CustomerSubscriptionDetailsOutput struct{ *pulumi.OutputState }

func (CustomerSubscriptionDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSubscriptionDetails)(nil)).Elem()
}

func (o CustomerSubscriptionDetailsOutput) ToCustomerSubscriptionDetailsOutput() CustomerSubscriptionDetailsOutput {
	return o
}

func (o CustomerSubscriptionDetailsOutput) ToCustomerSubscriptionDetailsOutputWithContext(ctx context.Context) CustomerSubscriptionDetailsOutput {
	return o
}

func (o CustomerSubscriptionDetailsOutput) ToCustomerSubscriptionDetailsPtrOutput() CustomerSubscriptionDetailsPtrOutput {
	return o.ToCustomerSubscriptionDetailsPtrOutputWithContext(context.Background())
}

func (o CustomerSubscriptionDetailsOutput) ToCustomerSubscriptionDetailsPtrOutputWithContext(ctx context.Context) CustomerSubscriptionDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomerSubscriptionDetails) *CustomerSubscriptionDetails {
		return &v
	}).(CustomerSubscriptionDetailsPtrOutput)
}

func (o CustomerSubscriptionDetailsOutput) LocationPlacementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomerSubscriptionDetails) *string { return v.LocationPlacementId }).(pulumi.StringPtrOutput)
}

func (o CustomerSubscriptionDetailsOutput) QuotaId() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSubscriptionDetails) string { return v.QuotaId }).(pulumi.StringOutput)
}

func (o CustomerSubscriptionDetailsOutput) RegisteredFeatures() CustomerSubscriptionRegisteredFeaturesArrayOutput {
	return o.ApplyT(func(v CustomerSubscriptionDetails) []CustomerSubscriptionRegisteredFeatures {
		return v.RegisteredFeatures
	}).(CustomerSubscriptionRegisteredFeaturesArrayOutput)
}

type CustomerSubscriptionDetailsPtrOutput struct{ *pulumi.OutputState }

func (CustomerSubscriptionDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerSubscriptionDetails)(nil)).Elem()
}

func (o CustomerSubscriptionDetailsPtrOutput) ToCustomerSubscriptionDetailsPtrOutput() CustomerSubscriptionDetailsPtrOutput {
	return o
}

func (o CustomerSubscriptionDetailsPtrOutput) ToCustomerSubscriptionDetailsPtrOutputWithContext(ctx context.Context) CustomerSubscriptionDetailsPtrOutput {
	return o
}

func (o CustomerSubscriptionDetailsPtrOutput) Elem() CustomerSubscriptionDetailsOutput {
	return o.ApplyT(func(v *CustomerSubscriptionDetails) CustomerSubscriptionDetails {
		if v != nil {
			return *v
		}
		var ret CustomerSubscriptionDetails
		return ret
	}).(CustomerSubscriptionDetailsOutput)
}

func (o CustomerSubscriptionDetailsPtrOutput) LocationPlacementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerSubscriptionDetails) *string {
		if v == nil {
			return nil
		}
		return v.LocationPlacementId
	}).(pulumi.StringPtrOutput)
}

func (o CustomerSubscriptionDetailsPtrOutput) QuotaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerSubscriptionDetails) *string {
		if v == nil {
			return nil
		}
		return &v.QuotaId
	}).(pulumi.StringPtrOutput)
}

func (o CustomerSubscriptionDetailsPtrOutput) RegisteredFeatures() CustomerSubscriptionRegisteredFeaturesArrayOutput {
	return o.ApplyT(func(v *CustomerSubscriptionDetails) []CustomerSubscriptionRegisteredFeatures {
		if v == nil {
			return nil
		}
		return v.RegisteredFeatures
	}).(CustomerSubscriptionRegisteredFeaturesArrayOutput)
}

type CustomerSubscriptionRegisteredFeatures struct {
	Name  *string `pulumi:"name"`
	State *string `pulumi:"state"`
}





type CustomerSubscriptionRegisteredFeaturesInput interface {
	pulumi.Input

	ToCustomerSubscriptionRegisteredFeaturesOutput() CustomerSubscriptionRegisteredFeaturesOutput
	ToCustomerSubscriptionRegisteredFeaturesOutputWithContext(context.Context) CustomerSubscriptionRegisteredFeaturesOutput
}

type CustomerSubscriptionRegisteredFeaturesArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (CustomerSubscriptionRegisteredFeaturesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSubscriptionRegisteredFeatures)(nil)).Elem()
}

func (i CustomerSubscriptionRegisteredFeaturesArgs) ToCustomerSubscriptionRegisteredFeaturesOutput() CustomerSubscriptionRegisteredFeaturesOutput {
	return i.ToCustomerSubscriptionRegisteredFeaturesOutputWithContext(context.Background())
}

func (i CustomerSubscriptionRegisteredFeaturesArgs) ToCustomerSubscriptionRegisteredFeaturesOutputWithContext(ctx context.Context) CustomerSubscriptionRegisteredFeaturesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionRegisteredFeaturesOutput)
}





type CustomerSubscriptionRegisteredFeaturesArrayInput interface {
	pulumi.Input

	ToCustomerSubscriptionRegisteredFeaturesArrayOutput() CustomerSubscriptionRegisteredFeaturesArrayOutput
	ToCustomerSubscriptionRegisteredFeaturesArrayOutputWithContext(context.Context) CustomerSubscriptionRegisteredFeaturesArrayOutput
}

type CustomerSubscriptionRegisteredFeaturesArray []CustomerSubscriptionRegisteredFeaturesInput

func (CustomerSubscriptionRegisteredFeaturesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomerSubscriptionRegisteredFeatures)(nil)).Elem()
}

func (i CustomerSubscriptionRegisteredFeaturesArray) ToCustomerSubscriptionRegisteredFeaturesArrayOutput() CustomerSubscriptionRegisteredFeaturesArrayOutput {
	return i.ToCustomerSubscriptionRegisteredFeaturesArrayOutputWithContext(context.Background())
}

func (i CustomerSubscriptionRegisteredFeaturesArray) ToCustomerSubscriptionRegisteredFeaturesArrayOutputWithContext(ctx context.Context) CustomerSubscriptionRegisteredFeaturesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionRegisteredFeaturesArrayOutput)
}

type CustomerSubscriptionRegisteredFeaturesOutput struct{ *pulumi.OutputState }

func (CustomerSubscriptionRegisteredFeaturesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSubscriptionRegisteredFeatures)(nil)).Elem()
}

func (o CustomerSubscriptionRegisteredFeaturesOutput) ToCustomerSubscriptionRegisteredFeaturesOutput() CustomerSubscriptionRegisteredFeaturesOutput {
	return o
}

func (o CustomerSubscriptionRegisteredFeaturesOutput) ToCustomerSubscriptionRegisteredFeaturesOutputWithContext(ctx context.Context) CustomerSubscriptionRegisteredFeaturesOutput {
	return o
}

func (o CustomerSubscriptionRegisteredFeaturesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomerSubscriptionRegisteredFeatures) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CustomerSubscriptionRegisteredFeaturesOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomerSubscriptionRegisteredFeatures) *string { return v.State }).(pulumi.StringPtrOutput)
}

type CustomerSubscriptionRegisteredFeaturesArrayOutput struct{ *pulumi.OutputState }

func (CustomerSubscriptionRegisteredFeaturesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomerSubscriptionRegisteredFeatures)(nil)).Elem()
}

func (o CustomerSubscriptionRegisteredFeaturesArrayOutput) ToCustomerSubscriptionRegisteredFeaturesArrayOutput() CustomerSubscriptionRegisteredFeaturesArrayOutput {
	return o
}

func (o CustomerSubscriptionRegisteredFeaturesArrayOutput) ToCustomerSubscriptionRegisteredFeaturesArrayOutputWithContext(ctx context.Context) CustomerSubscriptionRegisteredFeaturesArrayOutput {
	return o
}

func (o CustomerSubscriptionRegisteredFeaturesArrayOutput) Index(i pulumi.IntInput) CustomerSubscriptionRegisteredFeaturesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomerSubscriptionRegisteredFeatures {
		return vs[0].([]CustomerSubscriptionRegisteredFeatures)[vs[1].(int)]
	}).(CustomerSubscriptionRegisteredFeaturesOutput)
}

type DescriptionResponse struct {
	Attributes       []string       `pulumi:"attributes"`
	DescriptionType  string         `pulumi:"descriptionType"`
	Keywords         []string       `pulumi:"keywords"`
	Links            []LinkResponse `pulumi:"links"`
	LongDescription  string         `pulumi:"longDescription"`
	ShortDescription string         `pulumi:"shortDescription"`
}





type DescriptionResponseInput interface {
	pulumi.Input

	ToDescriptionResponseOutput() DescriptionResponseOutput
	ToDescriptionResponseOutputWithContext(context.Context) DescriptionResponseOutput
}

type DescriptionResponseArgs struct {
	Attributes       pulumi.StringArrayInput `pulumi:"attributes"`
	DescriptionType  pulumi.StringInput      `pulumi:"descriptionType"`
	Keywords         pulumi.StringArrayInput `pulumi:"keywords"`
	Links            LinkResponseArrayInput  `pulumi:"links"`
	LongDescription  pulumi.StringInput      `pulumi:"longDescription"`
	ShortDescription pulumi.StringInput      `pulumi:"shortDescription"`
}

func (DescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DescriptionResponse)(nil)).Elem()
}

func (i DescriptionResponseArgs) ToDescriptionResponseOutput() DescriptionResponseOutput {
	return i.ToDescriptionResponseOutputWithContext(context.Background())
}

func (i DescriptionResponseArgs) ToDescriptionResponseOutputWithContext(ctx context.Context) DescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DescriptionResponseOutput)
}

type DescriptionResponseOutput struct{ *pulumi.OutputState }

func (DescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DescriptionResponse)(nil)).Elem()
}

func (o DescriptionResponseOutput) ToDescriptionResponseOutput() DescriptionResponseOutput {
	return o
}

func (o DescriptionResponseOutput) ToDescriptionResponseOutputWithContext(ctx context.Context) DescriptionResponseOutput {
	return o
}

func (o DescriptionResponseOutput) Attributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DescriptionResponse) []string { return v.Attributes }).(pulumi.StringArrayOutput)
}

func (o DescriptionResponseOutput) DescriptionType() pulumi.StringOutput {
	return o.ApplyT(func(v DescriptionResponse) string { return v.DescriptionType }).(pulumi.StringOutput)
}

func (o DescriptionResponseOutput) Keywords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DescriptionResponse) []string { return v.Keywords }).(pulumi.StringArrayOutput)
}

func (o DescriptionResponseOutput) Links() LinkResponseArrayOutput {
	return o.ApplyT(func(v DescriptionResponse) []LinkResponse { return v.Links }).(LinkResponseArrayOutput)
}

func (o DescriptionResponseOutput) LongDescription() pulumi.StringOutput {
	return o.ApplyT(func(v DescriptionResponse) string { return v.LongDescription }).(pulumi.StringOutput)
}

func (o DescriptionResponseOutput) ShortDescription() pulumi.StringOutput {
	return o.ApplyT(func(v DescriptionResponse) string { return v.ShortDescription }).(pulumi.StringOutput)
}

type DeviceDetailsResponse struct {
	ManagementResourceId string `pulumi:"managementResourceId"`
	SerialNumber         string `pulumi:"serialNumber"`
}





type DeviceDetailsResponseInput interface {
	pulumi.Input

	ToDeviceDetailsResponseOutput() DeviceDetailsResponseOutput
	ToDeviceDetailsResponseOutputWithContext(context.Context) DeviceDetailsResponseOutput
}

type DeviceDetailsResponseArgs struct {
	ManagementResourceId pulumi.StringInput `pulumi:"managementResourceId"`
	SerialNumber         pulumi.StringInput `pulumi:"serialNumber"`
}

func (DeviceDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceDetailsResponse)(nil)).Elem()
}

func (i DeviceDetailsResponseArgs) ToDeviceDetailsResponseOutput() DeviceDetailsResponseOutput {
	return i.ToDeviceDetailsResponseOutputWithContext(context.Background())
}

func (i DeviceDetailsResponseArgs) ToDeviceDetailsResponseOutputWithContext(ctx context.Context) DeviceDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceDetailsResponseOutput)
}





type DeviceDetailsResponseArrayInput interface {
	pulumi.Input

	ToDeviceDetailsResponseArrayOutput() DeviceDetailsResponseArrayOutput
	ToDeviceDetailsResponseArrayOutputWithContext(context.Context) DeviceDetailsResponseArrayOutput
}

type DeviceDetailsResponseArray []DeviceDetailsResponseInput

func (DeviceDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeviceDetailsResponse)(nil)).Elem()
}

func (i DeviceDetailsResponseArray) ToDeviceDetailsResponseArrayOutput() DeviceDetailsResponseArrayOutput {
	return i.ToDeviceDetailsResponseArrayOutputWithContext(context.Background())
}

func (i DeviceDetailsResponseArray) ToDeviceDetailsResponseArrayOutputWithContext(ctx context.Context) DeviceDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceDetailsResponseArrayOutput)
}

type DeviceDetailsResponseOutput struct{ *pulumi.OutputState }

func (DeviceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceDetailsResponse)(nil)).Elem()
}

func (o DeviceDetailsResponseOutput) ToDeviceDetailsResponseOutput() DeviceDetailsResponseOutput {
	return o
}

func (o DeviceDetailsResponseOutput) ToDeviceDetailsResponseOutputWithContext(ctx context.Context) DeviceDetailsResponseOutput {
	return o
}

func (o DeviceDetailsResponseOutput) ManagementResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v DeviceDetailsResponse) string { return v.ManagementResourceId }).(pulumi.StringOutput)
}

func (o DeviceDetailsResponseOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DeviceDetailsResponse) string { return v.SerialNumber }).(pulumi.StringOutput)
}

type DeviceDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (DeviceDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeviceDetailsResponse)(nil)).Elem()
}

func (o DeviceDetailsResponseArrayOutput) ToDeviceDetailsResponseArrayOutput() DeviceDetailsResponseArrayOutput {
	return o
}

func (o DeviceDetailsResponseArrayOutput) ToDeviceDetailsResponseArrayOutputWithContext(ctx context.Context) DeviceDetailsResponseArrayOutput {
	return o
}

func (o DeviceDetailsResponseArrayOutput) Index(i pulumi.IntInput) DeviceDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeviceDetailsResponse {
		return vs[0].([]DeviceDetailsResponse)[vs[1].(int)]
	}).(DeviceDetailsResponseOutput)
}

type DimensionsResponse struct {
	Depth            float64 `pulumi:"depth"`
	Height           float64 `pulumi:"height"`
	Length           float64 `pulumi:"length"`
	LengthHeightUnit string  `pulumi:"lengthHeightUnit"`
	Weight           float64 `pulumi:"weight"`
	WeightUnit       string  `pulumi:"weightUnit"`
	Width            float64 `pulumi:"width"`
}





type DimensionsResponseInput interface {
	pulumi.Input

	ToDimensionsResponseOutput() DimensionsResponseOutput
	ToDimensionsResponseOutputWithContext(context.Context) DimensionsResponseOutput
}

type DimensionsResponseArgs struct {
	Depth            pulumi.Float64Input `pulumi:"depth"`
	Height           pulumi.Float64Input `pulumi:"height"`
	Length           pulumi.Float64Input `pulumi:"length"`
	LengthHeightUnit pulumi.StringInput  `pulumi:"lengthHeightUnit"`
	Weight           pulumi.Float64Input `pulumi:"weight"`
	WeightUnit       pulumi.StringInput  `pulumi:"weightUnit"`
	Width            pulumi.Float64Input `pulumi:"width"`
}

func (DimensionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DimensionsResponse)(nil)).Elem()
}

func (i DimensionsResponseArgs) ToDimensionsResponseOutput() DimensionsResponseOutput {
	return i.ToDimensionsResponseOutputWithContext(context.Background())
}

func (i DimensionsResponseArgs) ToDimensionsResponseOutputWithContext(ctx context.Context) DimensionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DimensionsResponseOutput)
}

type DimensionsResponseOutput struct{ *pulumi.OutputState }

func (DimensionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DimensionsResponse)(nil)).Elem()
}

func (o DimensionsResponseOutput) ToDimensionsResponseOutput() DimensionsResponseOutput {
	return o
}

func (o DimensionsResponseOutput) ToDimensionsResponseOutputWithContext(ctx context.Context) DimensionsResponseOutput {
	return o
}

func (o DimensionsResponseOutput) Depth() pulumi.Float64Output {
	return o.ApplyT(func(v DimensionsResponse) float64 { return v.Depth }).(pulumi.Float64Output)
}

func (o DimensionsResponseOutput) Height() pulumi.Float64Output {
	return o.ApplyT(func(v DimensionsResponse) float64 { return v.Height }).(pulumi.Float64Output)
}

func (o DimensionsResponseOutput) Length() pulumi.Float64Output {
	return o.ApplyT(func(v DimensionsResponse) float64 { return v.Length }).(pulumi.Float64Output)
}

func (o DimensionsResponseOutput) LengthHeightUnit() pulumi.StringOutput {
	return o.ApplyT(func(v DimensionsResponse) string { return v.LengthHeightUnit }).(pulumi.StringOutput)
}

func (o DimensionsResponseOutput) Weight() pulumi.Float64Output {
	return o.ApplyT(func(v DimensionsResponse) float64 { return v.Weight }).(pulumi.Float64Output)
}

func (o DimensionsResponseOutput) WeightUnit() pulumi.StringOutput {
	return o.ApplyT(func(v DimensionsResponse) string { return v.WeightUnit }).(pulumi.StringOutput)
}

func (o DimensionsResponseOutput) Width() pulumi.Float64Output {
	return o.ApplyT(func(v DimensionsResponse) float64 { return v.Width }).(pulumi.Float64Output)
}

type DisplayInfoResponse struct {
	ConfigurationDisplayName string `pulumi:"configurationDisplayName"`
	ProductFamilyDisplayName string `pulumi:"productFamilyDisplayName"`
}





type DisplayInfoResponseInput interface {
	pulumi.Input

	ToDisplayInfoResponseOutput() DisplayInfoResponseOutput
	ToDisplayInfoResponseOutputWithContext(context.Context) DisplayInfoResponseOutput
}

type DisplayInfoResponseArgs struct {
	ConfigurationDisplayName pulumi.StringInput `pulumi:"configurationDisplayName"`
	ProductFamilyDisplayName pulumi.StringInput `pulumi:"productFamilyDisplayName"`
}

func (DisplayInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DisplayInfoResponse)(nil)).Elem()
}

func (i DisplayInfoResponseArgs) ToDisplayInfoResponseOutput() DisplayInfoResponseOutput {
	return i.ToDisplayInfoResponseOutputWithContext(context.Background())
}

func (i DisplayInfoResponseArgs) ToDisplayInfoResponseOutputWithContext(ctx context.Context) DisplayInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisplayInfoResponseOutput)
}

func (i DisplayInfoResponseArgs) ToDisplayInfoResponsePtrOutput() DisplayInfoResponsePtrOutput {
	return i.ToDisplayInfoResponsePtrOutputWithContext(context.Background())
}

func (i DisplayInfoResponseArgs) ToDisplayInfoResponsePtrOutputWithContext(ctx context.Context) DisplayInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisplayInfoResponseOutput).ToDisplayInfoResponsePtrOutputWithContext(ctx)
}









type DisplayInfoResponsePtrInput interface {
	pulumi.Input

	ToDisplayInfoResponsePtrOutput() DisplayInfoResponsePtrOutput
	ToDisplayInfoResponsePtrOutputWithContext(context.Context) DisplayInfoResponsePtrOutput
}

type displayInfoResponsePtrType DisplayInfoResponseArgs

func DisplayInfoResponsePtr(v *DisplayInfoResponseArgs) DisplayInfoResponsePtrInput {
	return (*displayInfoResponsePtrType)(v)
}

func (*displayInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DisplayInfoResponse)(nil)).Elem()
}

func (i *displayInfoResponsePtrType) ToDisplayInfoResponsePtrOutput() DisplayInfoResponsePtrOutput {
	return i.ToDisplayInfoResponsePtrOutputWithContext(context.Background())
}

func (i *displayInfoResponsePtrType) ToDisplayInfoResponsePtrOutputWithContext(ctx context.Context) DisplayInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisplayInfoResponsePtrOutput)
}

type DisplayInfoResponseOutput struct{ *pulumi.OutputState }

func (DisplayInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DisplayInfoResponse)(nil)).Elem()
}

func (o DisplayInfoResponseOutput) ToDisplayInfoResponseOutput() DisplayInfoResponseOutput {
	return o
}

func (o DisplayInfoResponseOutput) ToDisplayInfoResponseOutputWithContext(ctx context.Context) DisplayInfoResponseOutput {
	return o
}

func (o DisplayInfoResponseOutput) ToDisplayInfoResponsePtrOutput() DisplayInfoResponsePtrOutput {
	return o.ToDisplayInfoResponsePtrOutputWithContext(context.Background())
}

func (o DisplayInfoResponseOutput) ToDisplayInfoResponsePtrOutputWithContext(ctx context.Context) DisplayInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DisplayInfoResponse) *DisplayInfoResponse {
		return &v
	}).(DisplayInfoResponsePtrOutput)
}

func (o DisplayInfoResponseOutput) ConfigurationDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v DisplayInfoResponse) string { return v.ConfigurationDisplayName }).(pulumi.StringOutput)
}

func (o DisplayInfoResponseOutput) ProductFamilyDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v DisplayInfoResponse) string { return v.ProductFamilyDisplayName }).(pulumi.StringOutput)
}

type DisplayInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (DisplayInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DisplayInfoResponse)(nil)).Elem()
}

func (o DisplayInfoResponsePtrOutput) ToDisplayInfoResponsePtrOutput() DisplayInfoResponsePtrOutput {
	return o
}

func (o DisplayInfoResponsePtrOutput) ToDisplayInfoResponsePtrOutputWithContext(ctx context.Context) DisplayInfoResponsePtrOutput {
	return o
}

func (o DisplayInfoResponsePtrOutput) Elem() DisplayInfoResponseOutput {
	return o.ApplyT(func(v *DisplayInfoResponse) DisplayInfoResponse {
		if v != nil {
			return *v
		}
		var ret DisplayInfoResponse
		return ret
	}).(DisplayInfoResponseOutput)
}

func (o DisplayInfoResponsePtrOutput) ConfigurationDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DisplayInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ConfigurationDisplayName
	}).(pulumi.StringPtrOutput)
}

func (o DisplayInfoResponsePtrOutput) ProductFamilyDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DisplayInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProductFamilyDisplayName
	}).(pulumi.StringPtrOutput)
}

type EncryptionPreferences struct {
	DoubleEncryptionStatus *string `pulumi:"doubleEncryptionStatus"`
}





type EncryptionPreferencesInput interface {
	pulumi.Input

	ToEncryptionPreferencesOutput() EncryptionPreferencesOutput
	ToEncryptionPreferencesOutputWithContext(context.Context) EncryptionPreferencesOutput
}

type EncryptionPreferencesArgs struct {
	DoubleEncryptionStatus pulumi.StringPtrInput `pulumi:"doubleEncryptionStatus"`
}

func (EncryptionPreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPreferences)(nil)).Elem()
}

func (i EncryptionPreferencesArgs) ToEncryptionPreferencesOutput() EncryptionPreferencesOutput {
	return i.ToEncryptionPreferencesOutputWithContext(context.Background())
}

func (i EncryptionPreferencesArgs) ToEncryptionPreferencesOutputWithContext(ctx context.Context) EncryptionPreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesOutput)
}

func (i EncryptionPreferencesArgs) ToEncryptionPreferencesPtrOutput() EncryptionPreferencesPtrOutput {
	return i.ToEncryptionPreferencesPtrOutputWithContext(context.Background())
}

func (i EncryptionPreferencesArgs) ToEncryptionPreferencesPtrOutputWithContext(ctx context.Context) EncryptionPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesOutput).ToEncryptionPreferencesPtrOutputWithContext(ctx)
}









type EncryptionPreferencesPtrInput interface {
	pulumi.Input

	ToEncryptionPreferencesPtrOutput() EncryptionPreferencesPtrOutput
	ToEncryptionPreferencesPtrOutputWithContext(context.Context) EncryptionPreferencesPtrOutput
}

type encryptionPreferencesPtrType EncryptionPreferencesArgs

func EncryptionPreferencesPtr(v *EncryptionPreferencesArgs) EncryptionPreferencesPtrInput {
	return (*encryptionPreferencesPtrType)(v)
}

func (*encryptionPreferencesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPreferences)(nil)).Elem()
}

func (i *encryptionPreferencesPtrType) ToEncryptionPreferencesPtrOutput() EncryptionPreferencesPtrOutput {
	return i.ToEncryptionPreferencesPtrOutputWithContext(context.Background())
}

func (i *encryptionPreferencesPtrType) ToEncryptionPreferencesPtrOutputWithContext(ctx context.Context) EncryptionPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesPtrOutput)
}

type EncryptionPreferencesOutput struct{ *pulumi.OutputState }

func (EncryptionPreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPreferences)(nil)).Elem()
}

func (o EncryptionPreferencesOutput) ToEncryptionPreferencesOutput() EncryptionPreferencesOutput {
	return o
}

func (o EncryptionPreferencesOutput) ToEncryptionPreferencesOutputWithContext(ctx context.Context) EncryptionPreferencesOutput {
	return o
}

func (o EncryptionPreferencesOutput) ToEncryptionPreferencesPtrOutput() EncryptionPreferencesPtrOutput {
	return o.ToEncryptionPreferencesPtrOutputWithContext(context.Background())
}

func (o EncryptionPreferencesOutput) ToEncryptionPreferencesPtrOutputWithContext(ctx context.Context) EncryptionPreferencesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPreferences) *EncryptionPreferences {
		return &v
	}).(EncryptionPreferencesPtrOutput)
}

func (o EncryptionPreferencesOutput) DoubleEncryptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPreferences) *string { return v.DoubleEncryptionStatus }).(pulumi.StringPtrOutput)
}

type EncryptionPreferencesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPreferencesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPreferences)(nil)).Elem()
}

func (o EncryptionPreferencesPtrOutput) ToEncryptionPreferencesPtrOutput() EncryptionPreferencesPtrOutput {
	return o
}

func (o EncryptionPreferencesPtrOutput) ToEncryptionPreferencesPtrOutputWithContext(ctx context.Context) EncryptionPreferencesPtrOutput {
	return o
}

func (o EncryptionPreferencesPtrOutput) Elem() EncryptionPreferencesOutput {
	return o.ApplyT(func(v *EncryptionPreferences) EncryptionPreferences {
		if v != nil {
			return *v
		}
		var ret EncryptionPreferences
		return ret
	}).(EncryptionPreferencesOutput)
}

func (o EncryptionPreferencesPtrOutput) DoubleEncryptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPreferences) *string {
		if v == nil {
			return nil
		}
		return v.DoubleEncryptionStatus
	}).(pulumi.StringPtrOutput)
}

type EncryptionPreferencesResponse struct {
	DoubleEncryptionStatus *string `pulumi:"doubleEncryptionStatus"`
}





type EncryptionPreferencesResponseInput interface {
	pulumi.Input

	ToEncryptionPreferencesResponseOutput() EncryptionPreferencesResponseOutput
	ToEncryptionPreferencesResponseOutputWithContext(context.Context) EncryptionPreferencesResponseOutput
}

type EncryptionPreferencesResponseArgs struct {
	DoubleEncryptionStatus pulumi.StringPtrInput `pulumi:"doubleEncryptionStatus"`
}

func (EncryptionPreferencesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPreferencesResponse)(nil)).Elem()
}

func (i EncryptionPreferencesResponseArgs) ToEncryptionPreferencesResponseOutput() EncryptionPreferencesResponseOutput {
	return i.ToEncryptionPreferencesResponseOutputWithContext(context.Background())
}

func (i EncryptionPreferencesResponseArgs) ToEncryptionPreferencesResponseOutputWithContext(ctx context.Context) EncryptionPreferencesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesResponseOutput)
}

func (i EncryptionPreferencesResponseArgs) ToEncryptionPreferencesResponsePtrOutput() EncryptionPreferencesResponsePtrOutput {
	return i.ToEncryptionPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionPreferencesResponseArgs) ToEncryptionPreferencesResponsePtrOutputWithContext(ctx context.Context) EncryptionPreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesResponseOutput).ToEncryptionPreferencesResponsePtrOutputWithContext(ctx)
}









type EncryptionPreferencesResponsePtrInput interface {
	pulumi.Input

	ToEncryptionPreferencesResponsePtrOutput() EncryptionPreferencesResponsePtrOutput
	ToEncryptionPreferencesResponsePtrOutputWithContext(context.Context) EncryptionPreferencesResponsePtrOutput
}

type encryptionPreferencesResponsePtrType EncryptionPreferencesResponseArgs

func EncryptionPreferencesResponsePtr(v *EncryptionPreferencesResponseArgs) EncryptionPreferencesResponsePtrInput {
	return (*encryptionPreferencesResponsePtrType)(v)
}

func (*encryptionPreferencesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPreferencesResponse)(nil)).Elem()
}

func (i *encryptionPreferencesResponsePtrType) ToEncryptionPreferencesResponsePtrOutput() EncryptionPreferencesResponsePtrOutput {
	return i.ToEncryptionPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionPreferencesResponsePtrType) ToEncryptionPreferencesResponsePtrOutputWithContext(ctx context.Context) EncryptionPreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesResponsePtrOutput)
}

type EncryptionPreferencesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPreferencesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPreferencesResponse)(nil)).Elem()
}

func (o EncryptionPreferencesResponseOutput) ToEncryptionPreferencesResponseOutput() EncryptionPreferencesResponseOutput {
	return o
}

func (o EncryptionPreferencesResponseOutput) ToEncryptionPreferencesResponseOutputWithContext(ctx context.Context) EncryptionPreferencesResponseOutput {
	return o
}

func (o EncryptionPreferencesResponseOutput) ToEncryptionPreferencesResponsePtrOutput() EncryptionPreferencesResponsePtrOutput {
	return o.ToEncryptionPreferencesResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionPreferencesResponseOutput) ToEncryptionPreferencesResponsePtrOutputWithContext(ctx context.Context) EncryptionPreferencesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPreferencesResponse) *EncryptionPreferencesResponse {
		return &v
	}).(EncryptionPreferencesResponsePtrOutput)
}

func (o EncryptionPreferencesResponseOutput) DoubleEncryptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPreferencesResponse) *string { return v.DoubleEncryptionStatus }).(pulumi.StringPtrOutput)
}

type EncryptionPreferencesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPreferencesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPreferencesResponse)(nil)).Elem()
}

func (o EncryptionPreferencesResponsePtrOutput) ToEncryptionPreferencesResponsePtrOutput() EncryptionPreferencesResponsePtrOutput {
	return o
}

func (o EncryptionPreferencesResponsePtrOutput) ToEncryptionPreferencesResponsePtrOutputWithContext(ctx context.Context) EncryptionPreferencesResponsePtrOutput {
	return o
}

func (o EncryptionPreferencesResponsePtrOutput) Elem() EncryptionPreferencesResponseOutput {
	return o.ApplyT(func(v *EncryptionPreferencesResponse) EncryptionPreferencesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPreferencesResponse
		return ret
	}).(EncryptionPreferencesResponseOutput)
}

func (o EncryptionPreferencesResponsePtrOutput) DoubleEncryptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPreferencesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DoubleEncryptionStatus
	}).(pulumi.StringPtrOutput)
}

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}





type ErrorAdditionalInfoResponseInput interface {
	pulumi.Input

	ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput
	ToErrorAdditionalInfoResponseOutputWithContext(context.Context) ErrorAdditionalInfoResponseOutput
}

type ErrorAdditionalInfoResponseArgs struct {
	Info pulumi.Input       `pulumi:"info"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ErrorAdditionalInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (i ErrorAdditionalInfoResponseArgs) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return i.ToErrorAdditionalInfoResponseOutputWithContext(context.Background())
}

func (i ErrorAdditionalInfoResponseArgs) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorAdditionalInfoResponseOutput)
}





type ErrorAdditionalInfoResponseArrayInput interface {
	pulumi.Input

	ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput
	ToErrorAdditionalInfoResponseArrayOutputWithContext(context.Context) ErrorAdditionalInfoResponseArrayOutput
}

type ErrorAdditionalInfoResponseArray []ErrorAdditionalInfoResponseInput

func (ErrorAdditionalInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (i ErrorAdditionalInfoResponseArray) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return i.ToErrorAdditionalInfoResponseArrayOutputWithContext(context.Background())
}

func (i ErrorAdditionalInfoResponseArray) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorAdditionalInfoResponseArrayOutput)
}

type ErrorAdditionalInfoResponseOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o ErrorAdditionalInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ErrorAdditionalInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) Index(i pulumi.IntInput) ErrorAdditionalInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorAdditionalInfoResponse {
		return vs[0].([]ErrorAdditionalInfoResponse)[vs[1].(int)]
	}).(ErrorAdditionalInfoResponseOutput)
}

type ErrorDetailResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorDetailResponse         `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}





type ErrorDetailResponseInput interface {
	pulumi.Input

	ToErrorDetailResponseOutput() ErrorDetailResponseOutput
	ToErrorDetailResponseOutputWithContext(context.Context) ErrorDetailResponseOutput
}

type ErrorDetailResponseArgs struct {
	AdditionalInfo ErrorAdditionalInfoResponseArrayInput `pulumi:"additionalInfo"`
	Code           pulumi.StringInput                    `pulumi:"code"`
	Details        ErrorDetailResponseArrayInput         `pulumi:"details"`
	Message        pulumi.StringInput                    `pulumi:"message"`
	Target         pulumi.StringInput                    `pulumi:"target"`
}

func (ErrorDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return i.ToErrorDetailResponseOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput)
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return i.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput).ToErrorDetailResponsePtrOutputWithContext(ctx)
}









type ErrorDetailResponsePtrInput interface {
	pulumi.Input

	ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput
	ToErrorDetailResponsePtrOutputWithContext(context.Context) ErrorDetailResponsePtrOutput
}

type errorDetailResponsePtrType ErrorDetailResponseArgs

func ErrorDetailResponsePtr(v *ErrorDetailResponseArgs) ErrorDetailResponsePtrInput {
	return (*errorDetailResponsePtrType)(v)
}

func (*errorDetailResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDetailResponse)(nil)).Elem()
}

func (i *errorDetailResponsePtrType) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return i.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (i *errorDetailResponsePtrType) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponsePtrOutput)
}





type ErrorDetailResponseArrayInput interface {
	pulumi.Input

	ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput
	ToErrorDetailResponseArrayOutputWithContext(context.Context) ErrorDetailResponseArrayOutput
}

type ErrorDetailResponseArray []ErrorDetailResponseInput

func (ErrorDetailResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return i.ToErrorDetailResponseArrayOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseArrayOutput)
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return o.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ErrorDetailResponse) *ErrorDetailResponse {
		return &v
	}).(ErrorDetailResponsePtrOutput)
}

func (o ErrorDetailResponseOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorAdditionalInfoResponse { return v.AdditionalInfo }).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorDetailResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponsePtrOutput) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return o
}

func (o ErrorDetailResponsePtrOutput) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return o
}

func (o ErrorDetailResponsePtrOutput) Elem() ErrorDetailResponseOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) ErrorDetailResponse {
		if v != nil {
			return *v
		}
		var ret ErrorDetailResponse
		return ret
	}).(ErrorDetailResponseOutput)
}

func (o ErrorDetailResponsePtrOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) []ErrorAdditionalInfoResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalInfo
	}).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailResponsePtrOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) []ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Target
	}).(pulumi.StringPtrOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
}

type FilterableProperty struct {
	SupportedValues []string `pulumi:"supportedValues"`
	Type            string   `pulumi:"type"`
}





type FilterablePropertyInput interface {
	pulumi.Input

	ToFilterablePropertyOutput() FilterablePropertyOutput
	ToFilterablePropertyOutputWithContext(context.Context) FilterablePropertyOutput
}

type FilterablePropertyArgs struct {
	SupportedValues pulumi.StringArrayInput `pulumi:"supportedValues"`
	Type            pulumi.StringInput      `pulumi:"type"`
}

func (FilterablePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterableProperty)(nil)).Elem()
}

func (i FilterablePropertyArgs) ToFilterablePropertyOutput() FilterablePropertyOutput {
	return i.ToFilterablePropertyOutputWithContext(context.Background())
}

func (i FilterablePropertyArgs) ToFilterablePropertyOutputWithContext(ctx context.Context) FilterablePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterablePropertyOutput)
}





type FilterablePropertyArrayInput interface {
	pulumi.Input

	ToFilterablePropertyArrayOutput() FilterablePropertyArrayOutput
	ToFilterablePropertyArrayOutputWithContext(context.Context) FilterablePropertyArrayOutput
}

type FilterablePropertyArray []FilterablePropertyInput

func (FilterablePropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterableProperty)(nil)).Elem()
}

func (i FilterablePropertyArray) ToFilterablePropertyArrayOutput() FilterablePropertyArrayOutput {
	return i.ToFilterablePropertyArrayOutputWithContext(context.Background())
}

func (i FilterablePropertyArray) ToFilterablePropertyArrayOutputWithContext(ctx context.Context) FilterablePropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterablePropertyArrayOutput)
}

type FilterablePropertyOutput struct{ *pulumi.OutputState }

func (FilterablePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterableProperty)(nil)).Elem()
}

func (o FilterablePropertyOutput) ToFilterablePropertyOutput() FilterablePropertyOutput {
	return o
}

func (o FilterablePropertyOutput) ToFilterablePropertyOutputWithContext(ctx context.Context) FilterablePropertyOutput {
	return o
}

func (o FilterablePropertyOutput) SupportedValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FilterableProperty) []string { return v.SupportedValues }).(pulumi.StringArrayOutput)
}

func (o FilterablePropertyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FilterableProperty) string { return v.Type }).(pulumi.StringOutput)
}

type FilterablePropertyArrayOutput struct{ *pulumi.OutputState }

func (FilterablePropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterableProperty)(nil)).Elem()
}

func (o FilterablePropertyArrayOutput) ToFilterablePropertyArrayOutput() FilterablePropertyArrayOutput {
	return o
}

func (o FilterablePropertyArrayOutput) ToFilterablePropertyArrayOutputWithContext(ctx context.Context) FilterablePropertyArrayOutput {
	return o
}

func (o FilterablePropertyArrayOutput) Index(i pulumi.IntInput) FilterablePropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterableProperty {
		return vs[0].([]FilterableProperty)[vs[1].(int)]
	}).(FilterablePropertyOutput)
}

type FilterablePropertyResponse struct {
	SupportedValues []string `pulumi:"supportedValues"`
	Type            string   `pulumi:"type"`
}





type FilterablePropertyResponseInput interface {
	pulumi.Input

	ToFilterablePropertyResponseOutput() FilterablePropertyResponseOutput
	ToFilterablePropertyResponseOutputWithContext(context.Context) FilterablePropertyResponseOutput
}

type FilterablePropertyResponseArgs struct {
	SupportedValues pulumi.StringArrayInput `pulumi:"supportedValues"`
	Type            pulumi.StringInput      `pulumi:"type"`
}

func (FilterablePropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterablePropertyResponse)(nil)).Elem()
}

func (i FilterablePropertyResponseArgs) ToFilterablePropertyResponseOutput() FilterablePropertyResponseOutput {
	return i.ToFilterablePropertyResponseOutputWithContext(context.Background())
}

func (i FilterablePropertyResponseArgs) ToFilterablePropertyResponseOutputWithContext(ctx context.Context) FilterablePropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterablePropertyResponseOutput)
}





type FilterablePropertyResponseArrayInput interface {
	pulumi.Input

	ToFilterablePropertyResponseArrayOutput() FilterablePropertyResponseArrayOutput
	ToFilterablePropertyResponseArrayOutputWithContext(context.Context) FilterablePropertyResponseArrayOutput
}

type FilterablePropertyResponseArray []FilterablePropertyResponseInput

func (FilterablePropertyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterablePropertyResponse)(nil)).Elem()
}

func (i FilterablePropertyResponseArray) ToFilterablePropertyResponseArrayOutput() FilterablePropertyResponseArrayOutput {
	return i.ToFilterablePropertyResponseArrayOutputWithContext(context.Background())
}

func (i FilterablePropertyResponseArray) ToFilterablePropertyResponseArrayOutputWithContext(ctx context.Context) FilterablePropertyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterablePropertyResponseArrayOutput)
}

type FilterablePropertyResponseOutput struct{ *pulumi.OutputState }

func (FilterablePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterablePropertyResponse)(nil)).Elem()
}

func (o FilterablePropertyResponseOutput) ToFilterablePropertyResponseOutput() FilterablePropertyResponseOutput {
	return o
}

func (o FilterablePropertyResponseOutput) ToFilterablePropertyResponseOutputWithContext(ctx context.Context) FilterablePropertyResponseOutput {
	return o
}

func (o FilterablePropertyResponseOutput) SupportedValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FilterablePropertyResponse) []string { return v.SupportedValues }).(pulumi.StringArrayOutput)
}

func (o FilterablePropertyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FilterablePropertyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type FilterablePropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (FilterablePropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterablePropertyResponse)(nil)).Elem()
}

func (o FilterablePropertyResponseArrayOutput) ToFilterablePropertyResponseArrayOutput() FilterablePropertyResponseArrayOutput {
	return o
}

func (o FilterablePropertyResponseArrayOutput) ToFilterablePropertyResponseArrayOutputWithContext(ctx context.Context) FilterablePropertyResponseArrayOutput {
	return o
}

func (o FilterablePropertyResponseArrayOutput) Index(i pulumi.IntInput) FilterablePropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterablePropertyResponse {
		return vs[0].([]FilterablePropertyResponse)[vs[1].(int)]
	}).(FilterablePropertyResponseOutput)
}

type HierarchyInformation struct {
	ConfigurationName *string `pulumi:"configurationName"`
	ProductFamilyName *string `pulumi:"productFamilyName"`
	ProductLineName   *string `pulumi:"productLineName"`
	ProductName       *string `pulumi:"productName"`
}





type HierarchyInformationInput interface {
	pulumi.Input

	ToHierarchyInformationOutput() HierarchyInformationOutput
	ToHierarchyInformationOutputWithContext(context.Context) HierarchyInformationOutput
}

type HierarchyInformationArgs struct {
	ConfigurationName pulumi.StringPtrInput `pulumi:"configurationName"`
	ProductFamilyName pulumi.StringPtrInput `pulumi:"productFamilyName"`
	ProductLineName   pulumi.StringPtrInput `pulumi:"productLineName"`
	ProductName       pulumi.StringPtrInput `pulumi:"productName"`
}

func (HierarchyInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HierarchyInformation)(nil)).Elem()
}

func (i HierarchyInformationArgs) ToHierarchyInformationOutput() HierarchyInformationOutput {
	return i.ToHierarchyInformationOutputWithContext(context.Background())
}

func (i HierarchyInformationArgs) ToHierarchyInformationOutputWithContext(ctx context.Context) HierarchyInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HierarchyInformationOutput)
}

func (i HierarchyInformationArgs) ToHierarchyInformationPtrOutput() HierarchyInformationPtrOutput {
	return i.ToHierarchyInformationPtrOutputWithContext(context.Background())
}

func (i HierarchyInformationArgs) ToHierarchyInformationPtrOutputWithContext(ctx context.Context) HierarchyInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HierarchyInformationOutput).ToHierarchyInformationPtrOutputWithContext(ctx)
}









type HierarchyInformationPtrInput interface {
	pulumi.Input

	ToHierarchyInformationPtrOutput() HierarchyInformationPtrOutput
	ToHierarchyInformationPtrOutputWithContext(context.Context) HierarchyInformationPtrOutput
}

type hierarchyInformationPtrType HierarchyInformationArgs

func HierarchyInformationPtr(v *HierarchyInformationArgs) HierarchyInformationPtrInput {
	return (*hierarchyInformationPtrType)(v)
}

func (*hierarchyInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HierarchyInformation)(nil)).Elem()
}

func (i *hierarchyInformationPtrType) ToHierarchyInformationPtrOutput() HierarchyInformationPtrOutput {
	return i.ToHierarchyInformationPtrOutputWithContext(context.Background())
}

func (i *hierarchyInformationPtrType) ToHierarchyInformationPtrOutputWithContext(ctx context.Context) HierarchyInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HierarchyInformationPtrOutput)
}

type HierarchyInformationOutput struct{ *pulumi.OutputState }

func (HierarchyInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HierarchyInformation)(nil)).Elem()
}

func (o HierarchyInformationOutput) ToHierarchyInformationOutput() HierarchyInformationOutput {
	return o
}

func (o HierarchyInformationOutput) ToHierarchyInformationOutputWithContext(ctx context.Context) HierarchyInformationOutput {
	return o
}

func (o HierarchyInformationOutput) ToHierarchyInformationPtrOutput() HierarchyInformationPtrOutput {
	return o.ToHierarchyInformationPtrOutputWithContext(context.Background())
}

func (o HierarchyInformationOutput) ToHierarchyInformationPtrOutputWithContext(ctx context.Context) HierarchyInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HierarchyInformation) *HierarchyInformation {
		return &v
	}).(HierarchyInformationPtrOutput)
}

func (o HierarchyInformationOutput) ConfigurationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HierarchyInformation) *string { return v.ConfigurationName }).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationOutput) ProductFamilyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HierarchyInformation) *string { return v.ProductFamilyName }).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationOutput) ProductLineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HierarchyInformation) *string { return v.ProductLineName }).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HierarchyInformation) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

type HierarchyInformationPtrOutput struct{ *pulumi.OutputState }

func (HierarchyInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HierarchyInformation)(nil)).Elem()
}

func (o HierarchyInformationPtrOutput) ToHierarchyInformationPtrOutput() HierarchyInformationPtrOutput {
	return o
}

func (o HierarchyInformationPtrOutput) ToHierarchyInformationPtrOutputWithContext(ctx context.Context) HierarchyInformationPtrOutput {
	return o
}

func (o HierarchyInformationPtrOutput) Elem() HierarchyInformationOutput {
	return o.ApplyT(func(v *HierarchyInformation) HierarchyInformation {
		if v != nil {
			return *v
		}
		var ret HierarchyInformation
		return ret
	}).(HierarchyInformationOutput)
}

func (o HierarchyInformationPtrOutput) ConfigurationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HierarchyInformation) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationName
	}).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationPtrOutput) ProductFamilyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HierarchyInformation) *string {
		if v == nil {
			return nil
		}
		return v.ProductFamilyName
	}).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationPtrOutput) ProductLineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HierarchyInformation) *string {
		if v == nil {
			return nil
		}
		return v.ProductLineName
	}).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationPtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HierarchyInformation) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

type HierarchyInformationResponse struct {
	ConfigurationName *string `pulumi:"configurationName"`
	ProductFamilyName *string `pulumi:"productFamilyName"`
	ProductLineName   *string `pulumi:"productLineName"`
	ProductName       *string `pulumi:"productName"`
}





type HierarchyInformationResponseInput interface {
	pulumi.Input

	ToHierarchyInformationResponseOutput() HierarchyInformationResponseOutput
	ToHierarchyInformationResponseOutputWithContext(context.Context) HierarchyInformationResponseOutput
}

type HierarchyInformationResponseArgs struct {
	ConfigurationName pulumi.StringPtrInput `pulumi:"configurationName"`
	ProductFamilyName pulumi.StringPtrInput `pulumi:"productFamilyName"`
	ProductLineName   pulumi.StringPtrInput `pulumi:"productLineName"`
	ProductName       pulumi.StringPtrInput `pulumi:"productName"`
}

func (HierarchyInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HierarchyInformationResponse)(nil)).Elem()
}

func (i HierarchyInformationResponseArgs) ToHierarchyInformationResponseOutput() HierarchyInformationResponseOutput {
	return i.ToHierarchyInformationResponseOutputWithContext(context.Background())
}

func (i HierarchyInformationResponseArgs) ToHierarchyInformationResponseOutputWithContext(ctx context.Context) HierarchyInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HierarchyInformationResponseOutput)
}

func (i HierarchyInformationResponseArgs) ToHierarchyInformationResponsePtrOutput() HierarchyInformationResponsePtrOutput {
	return i.ToHierarchyInformationResponsePtrOutputWithContext(context.Background())
}

func (i HierarchyInformationResponseArgs) ToHierarchyInformationResponsePtrOutputWithContext(ctx context.Context) HierarchyInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HierarchyInformationResponseOutput).ToHierarchyInformationResponsePtrOutputWithContext(ctx)
}









type HierarchyInformationResponsePtrInput interface {
	pulumi.Input

	ToHierarchyInformationResponsePtrOutput() HierarchyInformationResponsePtrOutput
	ToHierarchyInformationResponsePtrOutputWithContext(context.Context) HierarchyInformationResponsePtrOutput
}

type hierarchyInformationResponsePtrType HierarchyInformationResponseArgs

func HierarchyInformationResponsePtr(v *HierarchyInformationResponseArgs) HierarchyInformationResponsePtrInput {
	return (*hierarchyInformationResponsePtrType)(v)
}

func (*hierarchyInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HierarchyInformationResponse)(nil)).Elem()
}

func (i *hierarchyInformationResponsePtrType) ToHierarchyInformationResponsePtrOutput() HierarchyInformationResponsePtrOutput {
	return i.ToHierarchyInformationResponsePtrOutputWithContext(context.Background())
}

func (i *hierarchyInformationResponsePtrType) ToHierarchyInformationResponsePtrOutputWithContext(ctx context.Context) HierarchyInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HierarchyInformationResponsePtrOutput)
}

type HierarchyInformationResponseOutput struct{ *pulumi.OutputState }

func (HierarchyInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HierarchyInformationResponse)(nil)).Elem()
}

func (o HierarchyInformationResponseOutput) ToHierarchyInformationResponseOutput() HierarchyInformationResponseOutput {
	return o
}

func (o HierarchyInformationResponseOutput) ToHierarchyInformationResponseOutputWithContext(ctx context.Context) HierarchyInformationResponseOutput {
	return o
}

func (o HierarchyInformationResponseOutput) ToHierarchyInformationResponsePtrOutput() HierarchyInformationResponsePtrOutput {
	return o.ToHierarchyInformationResponsePtrOutputWithContext(context.Background())
}

func (o HierarchyInformationResponseOutput) ToHierarchyInformationResponsePtrOutputWithContext(ctx context.Context) HierarchyInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HierarchyInformationResponse) *HierarchyInformationResponse {
		return &v
	}).(HierarchyInformationResponsePtrOutput)
}

func (o HierarchyInformationResponseOutput) ConfigurationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HierarchyInformationResponse) *string { return v.ConfigurationName }).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationResponseOutput) ProductFamilyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HierarchyInformationResponse) *string { return v.ProductFamilyName }).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationResponseOutput) ProductLineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HierarchyInformationResponse) *string { return v.ProductLineName }).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationResponseOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HierarchyInformationResponse) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

type HierarchyInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (HierarchyInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HierarchyInformationResponse)(nil)).Elem()
}

func (o HierarchyInformationResponsePtrOutput) ToHierarchyInformationResponsePtrOutput() HierarchyInformationResponsePtrOutput {
	return o
}

func (o HierarchyInformationResponsePtrOutput) ToHierarchyInformationResponsePtrOutputWithContext(ctx context.Context) HierarchyInformationResponsePtrOutput {
	return o
}

func (o HierarchyInformationResponsePtrOutput) Elem() HierarchyInformationResponseOutput {
	return o.ApplyT(func(v *HierarchyInformationResponse) HierarchyInformationResponse {
		if v != nil {
			return *v
		}
		var ret HierarchyInformationResponse
		return ret
	}).(HierarchyInformationResponseOutput)
}

func (o HierarchyInformationResponsePtrOutput) ConfigurationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HierarchyInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationName
	}).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationResponsePtrOutput) ProductFamilyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HierarchyInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProductFamilyName
	}).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationResponsePtrOutput) ProductLineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HierarchyInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProductLineName
	}).(pulumi.StringPtrOutput)
}

func (o HierarchyInformationResponsePtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HierarchyInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

type ImageInformationResponse struct {
	ImageType string `pulumi:"imageType"`
	ImageUrl  string `pulumi:"imageUrl"`
}





type ImageInformationResponseInput interface {
	pulumi.Input

	ToImageInformationResponseOutput() ImageInformationResponseOutput
	ToImageInformationResponseOutputWithContext(context.Context) ImageInformationResponseOutput
}

type ImageInformationResponseArgs struct {
	ImageType pulumi.StringInput `pulumi:"imageType"`
	ImageUrl  pulumi.StringInput `pulumi:"imageUrl"`
}

func (ImageInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageInformationResponse)(nil)).Elem()
}

func (i ImageInformationResponseArgs) ToImageInformationResponseOutput() ImageInformationResponseOutput {
	return i.ToImageInformationResponseOutputWithContext(context.Background())
}

func (i ImageInformationResponseArgs) ToImageInformationResponseOutputWithContext(ctx context.Context) ImageInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageInformationResponseOutput)
}





type ImageInformationResponseArrayInput interface {
	pulumi.Input

	ToImageInformationResponseArrayOutput() ImageInformationResponseArrayOutput
	ToImageInformationResponseArrayOutputWithContext(context.Context) ImageInformationResponseArrayOutput
}

type ImageInformationResponseArray []ImageInformationResponseInput

func (ImageInformationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageInformationResponse)(nil)).Elem()
}

func (i ImageInformationResponseArray) ToImageInformationResponseArrayOutput() ImageInformationResponseArrayOutput {
	return i.ToImageInformationResponseArrayOutputWithContext(context.Background())
}

func (i ImageInformationResponseArray) ToImageInformationResponseArrayOutputWithContext(ctx context.Context) ImageInformationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageInformationResponseArrayOutput)
}

type ImageInformationResponseOutput struct{ *pulumi.OutputState }

func (ImageInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageInformationResponse)(nil)).Elem()
}

func (o ImageInformationResponseOutput) ToImageInformationResponseOutput() ImageInformationResponseOutput {
	return o
}

func (o ImageInformationResponseOutput) ToImageInformationResponseOutputWithContext(ctx context.Context) ImageInformationResponseOutput {
	return o
}

func (o ImageInformationResponseOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v ImageInformationResponse) string { return v.ImageType }).(pulumi.StringOutput)
}

func (o ImageInformationResponseOutput) ImageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ImageInformationResponse) string { return v.ImageUrl }).(pulumi.StringOutput)
}

type ImageInformationResponseArrayOutput struct{ *pulumi.OutputState }

func (ImageInformationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageInformationResponse)(nil)).Elem()
}

func (o ImageInformationResponseArrayOutput) ToImageInformationResponseArrayOutput() ImageInformationResponseArrayOutput {
	return o
}

func (o ImageInformationResponseArrayOutput) ToImageInformationResponseArrayOutputWithContext(ctx context.Context) ImageInformationResponseArrayOutput {
	return o
}

func (o ImageInformationResponseArrayOutput) Index(i pulumi.IntInput) ImageInformationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageInformationResponse {
		return vs[0].([]ImageInformationResponse)[vs[1].(int)]
	}).(ImageInformationResponseOutput)
}

type LinkResponse struct {
	LinkType string `pulumi:"linkType"`
	LinkUrl  string `pulumi:"linkUrl"`
}





type LinkResponseInput interface {
	pulumi.Input

	ToLinkResponseOutput() LinkResponseOutput
	ToLinkResponseOutputWithContext(context.Context) LinkResponseOutput
}

type LinkResponseArgs struct {
	LinkType pulumi.StringInput `pulumi:"linkType"`
	LinkUrl  pulumi.StringInput `pulumi:"linkUrl"`
}

func (LinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkResponse)(nil)).Elem()
}

func (i LinkResponseArgs) ToLinkResponseOutput() LinkResponseOutput {
	return i.ToLinkResponseOutputWithContext(context.Background())
}

func (i LinkResponseArgs) ToLinkResponseOutputWithContext(ctx context.Context) LinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkResponseOutput)
}





type LinkResponseArrayInput interface {
	pulumi.Input

	ToLinkResponseArrayOutput() LinkResponseArrayOutput
	ToLinkResponseArrayOutputWithContext(context.Context) LinkResponseArrayOutput
}

type LinkResponseArray []LinkResponseInput

func (LinkResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkResponse)(nil)).Elem()
}

func (i LinkResponseArray) ToLinkResponseArrayOutput() LinkResponseArrayOutput {
	return i.ToLinkResponseArrayOutputWithContext(context.Background())
}

func (i LinkResponseArray) ToLinkResponseArrayOutputWithContext(ctx context.Context) LinkResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkResponseArrayOutput)
}

type LinkResponseOutput struct{ *pulumi.OutputState }

func (LinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkResponse)(nil)).Elem()
}

func (o LinkResponseOutput) ToLinkResponseOutput() LinkResponseOutput {
	return o
}

func (o LinkResponseOutput) ToLinkResponseOutputWithContext(ctx context.Context) LinkResponseOutput {
	return o
}

func (o LinkResponseOutput) LinkType() pulumi.StringOutput {
	return o.ApplyT(func(v LinkResponse) string { return v.LinkType }).(pulumi.StringOutput)
}

func (o LinkResponseOutput) LinkUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LinkResponse) string { return v.LinkUrl }).(pulumi.StringOutput)
}

type LinkResponseArrayOutput struct{ *pulumi.OutputState }

func (LinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkResponse)(nil)).Elem()
}

func (o LinkResponseArrayOutput) ToLinkResponseArrayOutput() LinkResponseArrayOutput {
	return o
}

func (o LinkResponseArrayOutput) ToLinkResponseArrayOutputWithContext(ctx context.Context) LinkResponseArrayOutput {
	return o
}

func (o LinkResponseArrayOutput) Index(i pulumi.IntInput) LinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkResponse {
		return vs[0].([]LinkResponse)[vs[1].(int)]
	}).(LinkResponseOutput)
}

type ManagementResourcePreferences struct {
	PreferredManagementResourceId *string `pulumi:"preferredManagementResourceId"`
}





type ManagementResourcePreferencesInput interface {
	pulumi.Input

	ToManagementResourcePreferencesOutput() ManagementResourcePreferencesOutput
	ToManagementResourcePreferencesOutputWithContext(context.Context) ManagementResourcePreferencesOutput
}

type ManagementResourcePreferencesArgs struct {
	PreferredManagementResourceId pulumi.StringPtrInput `pulumi:"preferredManagementResourceId"`
}

func (ManagementResourcePreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementResourcePreferences)(nil)).Elem()
}

func (i ManagementResourcePreferencesArgs) ToManagementResourcePreferencesOutput() ManagementResourcePreferencesOutput {
	return i.ToManagementResourcePreferencesOutputWithContext(context.Background())
}

func (i ManagementResourcePreferencesArgs) ToManagementResourcePreferencesOutputWithContext(ctx context.Context) ManagementResourcePreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementResourcePreferencesOutput)
}

func (i ManagementResourcePreferencesArgs) ToManagementResourcePreferencesPtrOutput() ManagementResourcePreferencesPtrOutput {
	return i.ToManagementResourcePreferencesPtrOutputWithContext(context.Background())
}

func (i ManagementResourcePreferencesArgs) ToManagementResourcePreferencesPtrOutputWithContext(ctx context.Context) ManagementResourcePreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementResourcePreferencesOutput).ToManagementResourcePreferencesPtrOutputWithContext(ctx)
}









type ManagementResourcePreferencesPtrInput interface {
	pulumi.Input

	ToManagementResourcePreferencesPtrOutput() ManagementResourcePreferencesPtrOutput
	ToManagementResourcePreferencesPtrOutputWithContext(context.Context) ManagementResourcePreferencesPtrOutput
}

type managementResourcePreferencesPtrType ManagementResourcePreferencesArgs

func ManagementResourcePreferencesPtr(v *ManagementResourcePreferencesArgs) ManagementResourcePreferencesPtrInput {
	return (*managementResourcePreferencesPtrType)(v)
}

func (*managementResourcePreferencesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementResourcePreferences)(nil)).Elem()
}

func (i *managementResourcePreferencesPtrType) ToManagementResourcePreferencesPtrOutput() ManagementResourcePreferencesPtrOutput {
	return i.ToManagementResourcePreferencesPtrOutputWithContext(context.Background())
}

func (i *managementResourcePreferencesPtrType) ToManagementResourcePreferencesPtrOutputWithContext(ctx context.Context) ManagementResourcePreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementResourcePreferencesPtrOutput)
}

type ManagementResourcePreferencesOutput struct{ *pulumi.OutputState }

func (ManagementResourcePreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementResourcePreferences)(nil)).Elem()
}

func (o ManagementResourcePreferencesOutput) ToManagementResourcePreferencesOutput() ManagementResourcePreferencesOutput {
	return o
}

func (o ManagementResourcePreferencesOutput) ToManagementResourcePreferencesOutputWithContext(ctx context.Context) ManagementResourcePreferencesOutput {
	return o
}

func (o ManagementResourcePreferencesOutput) ToManagementResourcePreferencesPtrOutput() ManagementResourcePreferencesPtrOutput {
	return o.ToManagementResourcePreferencesPtrOutputWithContext(context.Background())
}

func (o ManagementResourcePreferencesOutput) ToManagementResourcePreferencesPtrOutputWithContext(ctx context.Context) ManagementResourcePreferencesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementResourcePreferences) *ManagementResourcePreferences {
		return &v
	}).(ManagementResourcePreferencesPtrOutput)
}

func (o ManagementResourcePreferencesOutput) PreferredManagementResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementResourcePreferences) *string { return v.PreferredManagementResourceId }).(pulumi.StringPtrOutput)
}

type ManagementResourcePreferencesPtrOutput struct{ *pulumi.OutputState }

func (ManagementResourcePreferencesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementResourcePreferences)(nil)).Elem()
}

func (o ManagementResourcePreferencesPtrOutput) ToManagementResourcePreferencesPtrOutput() ManagementResourcePreferencesPtrOutput {
	return o
}

func (o ManagementResourcePreferencesPtrOutput) ToManagementResourcePreferencesPtrOutputWithContext(ctx context.Context) ManagementResourcePreferencesPtrOutput {
	return o
}

func (o ManagementResourcePreferencesPtrOutput) Elem() ManagementResourcePreferencesOutput {
	return o.ApplyT(func(v *ManagementResourcePreferences) ManagementResourcePreferences {
		if v != nil {
			return *v
		}
		var ret ManagementResourcePreferences
		return ret
	}).(ManagementResourcePreferencesOutput)
}

func (o ManagementResourcePreferencesPtrOutput) PreferredManagementResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementResourcePreferences) *string {
		if v == nil {
			return nil
		}
		return v.PreferredManagementResourceId
	}).(pulumi.StringPtrOutput)
}

type ManagementResourcePreferencesResponse struct {
	PreferredManagementResourceId *string `pulumi:"preferredManagementResourceId"`
}





type ManagementResourcePreferencesResponseInput interface {
	pulumi.Input

	ToManagementResourcePreferencesResponseOutput() ManagementResourcePreferencesResponseOutput
	ToManagementResourcePreferencesResponseOutputWithContext(context.Context) ManagementResourcePreferencesResponseOutput
}

type ManagementResourcePreferencesResponseArgs struct {
	PreferredManagementResourceId pulumi.StringPtrInput `pulumi:"preferredManagementResourceId"`
}

func (ManagementResourcePreferencesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementResourcePreferencesResponse)(nil)).Elem()
}

func (i ManagementResourcePreferencesResponseArgs) ToManagementResourcePreferencesResponseOutput() ManagementResourcePreferencesResponseOutput {
	return i.ToManagementResourcePreferencesResponseOutputWithContext(context.Background())
}

func (i ManagementResourcePreferencesResponseArgs) ToManagementResourcePreferencesResponseOutputWithContext(ctx context.Context) ManagementResourcePreferencesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementResourcePreferencesResponseOutput)
}

func (i ManagementResourcePreferencesResponseArgs) ToManagementResourcePreferencesResponsePtrOutput() ManagementResourcePreferencesResponsePtrOutput {
	return i.ToManagementResourcePreferencesResponsePtrOutputWithContext(context.Background())
}

func (i ManagementResourcePreferencesResponseArgs) ToManagementResourcePreferencesResponsePtrOutputWithContext(ctx context.Context) ManagementResourcePreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementResourcePreferencesResponseOutput).ToManagementResourcePreferencesResponsePtrOutputWithContext(ctx)
}









type ManagementResourcePreferencesResponsePtrInput interface {
	pulumi.Input

	ToManagementResourcePreferencesResponsePtrOutput() ManagementResourcePreferencesResponsePtrOutput
	ToManagementResourcePreferencesResponsePtrOutputWithContext(context.Context) ManagementResourcePreferencesResponsePtrOutput
}

type managementResourcePreferencesResponsePtrType ManagementResourcePreferencesResponseArgs

func ManagementResourcePreferencesResponsePtr(v *ManagementResourcePreferencesResponseArgs) ManagementResourcePreferencesResponsePtrInput {
	return (*managementResourcePreferencesResponsePtrType)(v)
}

func (*managementResourcePreferencesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementResourcePreferencesResponse)(nil)).Elem()
}

func (i *managementResourcePreferencesResponsePtrType) ToManagementResourcePreferencesResponsePtrOutput() ManagementResourcePreferencesResponsePtrOutput {
	return i.ToManagementResourcePreferencesResponsePtrOutputWithContext(context.Background())
}

func (i *managementResourcePreferencesResponsePtrType) ToManagementResourcePreferencesResponsePtrOutputWithContext(ctx context.Context) ManagementResourcePreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementResourcePreferencesResponsePtrOutput)
}

type ManagementResourcePreferencesResponseOutput struct{ *pulumi.OutputState }

func (ManagementResourcePreferencesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementResourcePreferencesResponse)(nil)).Elem()
}

func (o ManagementResourcePreferencesResponseOutput) ToManagementResourcePreferencesResponseOutput() ManagementResourcePreferencesResponseOutput {
	return o
}

func (o ManagementResourcePreferencesResponseOutput) ToManagementResourcePreferencesResponseOutputWithContext(ctx context.Context) ManagementResourcePreferencesResponseOutput {
	return o
}

func (o ManagementResourcePreferencesResponseOutput) ToManagementResourcePreferencesResponsePtrOutput() ManagementResourcePreferencesResponsePtrOutput {
	return o.ToManagementResourcePreferencesResponsePtrOutputWithContext(context.Background())
}

func (o ManagementResourcePreferencesResponseOutput) ToManagementResourcePreferencesResponsePtrOutputWithContext(ctx context.Context) ManagementResourcePreferencesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementResourcePreferencesResponse) *ManagementResourcePreferencesResponse {
		return &v
	}).(ManagementResourcePreferencesResponsePtrOutput)
}

func (o ManagementResourcePreferencesResponseOutput) PreferredManagementResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementResourcePreferencesResponse) *string { return v.PreferredManagementResourceId }).(pulumi.StringPtrOutput)
}

type ManagementResourcePreferencesResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementResourcePreferencesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementResourcePreferencesResponse)(nil)).Elem()
}

func (o ManagementResourcePreferencesResponsePtrOutput) ToManagementResourcePreferencesResponsePtrOutput() ManagementResourcePreferencesResponsePtrOutput {
	return o
}

func (o ManagementResourcePreferencesResponsePtrOutput) ToManagementResourcePreferencesResponsePtrOutputWithContext(ctx context.Context) ManagementResourcePreferencesResponsePtrOutput {
	return o
}

func (o ManagementResourcePreferencesResponsePtrOutput) Elem() ManagementResourcePreferencesResponseOutput {
	return o.ApplyT(func(v *ManagementResourcePreferencesResponse) ManagementResourcePreferencesResponse {
		if v != nil {
			return *v
		}
		var ret ManagementResourcePreferencesResponse
		return ret
	}).(ManagementResourcePreferencesResponseOutput)
}

func (o ManagementResourcePreferencesResponsePtrOutput) PreferredManagementResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementResourcePreferencesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreferredManagementResourceId
	}).(pulumi.StringPtrOutput)
}

type NotificationPreference struct {
	SendNotification bool   `pulumi:"sendNotification"`
	StageName        string `pulumi:"stageName"`
}





type NotificationPreferenceInput interface {
	pulumi.Input

	ToNotificationPreferenceOutput() NotificationPreferenceOutput
	ToNotificationPreferenceOutputWithContext(context.Context) NotificationPreferenceOutput
}

type NotificationPreferenceArgs struct {
	SendNotification pulumi.BoolInput   `pulumi:"sendNotification"`
	StageName        pulumi.StringInput `pulumi:"stageName"`
}

func (NotificationPreferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPreference)(nil)).Elem()
}

func (i NotificationPreferenceArgs) ToNotificationPreferenceOutput() NotificationPreferenceOutput {
	return i.ToNotificationPreferenceOutputWithContext(context.Background())
}

func (i NotificationPreferenceArgs) ToNotificationPreferenceOutputWithContext(ctx context.Context) NotificationPreferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPreferenceOutput)
}





type NotificationPreferenceArrayInput interface {
	pulumi.Input

	ToNotificationPreferenceArrayOutput() NotificationPreferenceArrayOutput
	ToNotificationPreferenceArrayOutputWithContext(context.Context) NotificationPreferenceArrayOutput
}

type NotificationPreferenceArray []NotificationPreferenceInput

func (NotificationPreferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationPreference)(nil)).Elem()
}

func (i NotificationPreferenceArray) ToNotificationPreferenceArrayOutput() NotificationPreferenceArrayOutput {
	return i.ToNotificationPreferenceArrayOutputWithContext(context.Background())
}

func (i NotificationPreferenceArray) ToNotificationPreferenceArrayOutputWithContext(ctx context.Context) NotificationPreferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPreferenceArrayOutput)
}

type NotificationPreferenceOutput struct{ *pulumi.OutputState }

func (NotificationPreferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPreference)(nil)).Elem()
}

func (o NotificationPreferenceOutput) ToNotificationPreferenceOutput() NotificationPreferenceOutput {
	return o
}

func (o NotificationPreferenceOutput) ToNotificationPreferenceOutputWithContext(ctx context.Context) NotificationPreferenceOutput {
	return o
}

func (o NotificationPreferenceOutput) SendNotification() pulumi.BoolOutput {
	return o.ApplyT(func(v NotificationPreference) bool { return v.SendNotification }).(pulumi.BoolOutput)
}

func (o NotificationPreferenceOutput) StageName() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationPreference) string { return v.StageName }).(pulumi.StringOutput)
}

type NotificationPreferenceArrayOutput struct{ *pulumi.OutputState }

func (NotificationPreferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationPreference)(nil)).Elem()
}

func (o NotificationPreferenceArrayOutput) ToNotificationPreferenceArrayOutput() NotificationPreferenceArrayOutput {
	return o
}

func (o NotificationPreferenceArrayOutput) ToNotificationPreferenceArrayOutputWithContext(ctx context.Context) NotificationPreferenceArrayOutput {
	return o
}

func (o NotificationPreferenceArrayOutput) Index(i pulumi.IntInput) NotificationPreferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationPreference {
		return vs[0].([]NotificationPreference)[vs[1].(int)]
	}).(NotificationPreferenceOutput)
}

type NotificationPreferenceResponse struct {
	SendNotification bool   `pulumi:"sendNotification"`
	StageName        string `pulumi:"stageName"`
}





type NotificationPreferenceResponseInput interface {
	pulumi.Input

	ToNotificationPreferenceResponseOutput() NotificationPreferenceResponseOutput
	ToNotificationPreferenceResponseOutputWithContext(context.Context) NotificationPreferenceResponseOutput
}

type NotificationPreferenceResponseArgs struct {
	SendNotification pulumi.BoolInput   `pulumi:"sendNotification"`
	StageName        pulumi.StringInput `pulumi:"stageName"`
}

func (NotificationPreferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPreferenceResponse)(nil)).Elem()
}

func (i NotificationPreferenceResponseArgs) ToNotificationPreferenceResponseOutput() NotificationPreferenceResponseOutput {
	return i.ToNotificationPreferenceResponseOutputWithContext(context.Background())
}

func (i NotificationPreferenceResponseArgs) ToNotificationPreferenceResponseOutputWithContext(ctx context.Context) NotificationPreferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPreferenceResponseOutput)
}





type NotificationPreferenceResponseArrayInput interface {
	pulumi.Input

	ToNotificationPreferenceResponseArrayOutput() NotificationPreferenceResponseArrayOutput
	ToNotificationPreferenceResponseArrayOutputWithContext(context.Context) NotificationPreferenceResponseArrayOutput
}

type NotificationPreferenceResponseArray []NotificationPreferenceResponseInput

func (NotificationPreferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationPreferenceResponse)(nil)).Elem()
}

func (i NotificationPreferenceResponseArray) ToNotificationPreferenceResponseArrayOutput() NotificationPreferenceResponseArrayOutput {
	return i.ToNotificationPreferenceResponseArrayOutputWithContext(context.Background())
}

func (i NotificationPreferenceResponseArray) ToNotificationPreferenceResponseArrayOutputWithContext(ctx context.Context) NotificationPreferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPreferenceResponseArrayOutput)
}

type NotificationPreferenceResponseOutput struct{ *pulumi.OutputState }

func (NotificationPreferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPreferenceResponse)(nil)).Elem()
}

func (o NotificationPreferenceResponseOutput) ToNotificationPreferenceResponseOutput() NotificationPreferenceResponseOutput {
	return o
}

func (o NotificationPreferenceResponseOutput) ToNotificationPreferenceResponseOutputWithContext(ctx context.Context) NotificationPreferenceResponseOutput {
	return o
}

func (o NotificationPreferenceResponseOutput) SendNotification() pulumi.BoolOutput {
	return o.ApplyT(func(v NotificationPreferenceResponse) bool { return v.SendNotification }).(pulumi.BoolOutput)
}

func (o NotificationPreferenceResponseOutput) StageName() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationPreferenceResponse) string { return v.StageName }).(pulumi.StringOutput)
}

type NotificationPreferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (NotificationPreferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationPreferenceResponse)(nil)).Elem()
}

func (o NotificationPreferenceResponseArrayOutput) ToNotificationPreferenceResponseArrayOutput() NotificationPreferenceResponseArrayOutput {
	return o
}

func (o NotificationPreferenceResponseArrayOutput) ToNotificationPreferenceResponseArrayOutputWithContext(ctx context.Context) NotificationPreferenceResponseArrayOutput {
	return o
}

func (o NotificationPreferenceResponseArrayOutput) Index(i pulumi.IntInput) NotificationPreferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationPreferenceResponse {
		return vs[0].([]NotificationPreferenceResponse)[vs[1].(int)]
	}).(NotificationPreferenceResponseOutput)
}

type OrderItemDetails struct {
	NotificationEmailList []string       `pulumi:"notificationEmailList"`
	OrderItemType         string         `pulumi:"orderItemType"`
	Preferences           *Preferences   `pulumi:"preferences"`
	ProductDetails        ProductDetails `pulumi:"productDetails"`
}





type OrderItemDetailsInput interface {
	pulumi.Input

	ToOrderItemDetailsOutput() OrderItemDetailsOutput
	ToOrderItemDetailsOutputWithContext(context.Context) OrderItemDetailsOutput
}

type OrderItemDetailsArgs struct {
	NotificationEmailList pulumi.StringArrayInput `pulumi:"notificationEmailList"`
	OrderItemType         pulumi.StringInput      `pulumi:"orderItemType"`
	Preferences           PreferencesPtrInput     `pulumi:"preferences"`
	ProductDetails        ProductDetailsInput     `pulumi:"productDetails"`
}

func (OrderItemDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderItemDetails)(nil)).Elem()
}

func (i OrderItemDetailsArgs) ToOrderItemDetailsOutput() OrderItemDetailsOutput {
	return i.ToOrderItemDetailsOutputWithContext(context.Background())
}

func (i OrderItemDetailsArgs) ToOrderItemDetailsOutputWithContext(ctx context.Context) OrderItemDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderItemDetailsOutput)
}

func (i OrderItemDetailsArgs) ToOrderItemDetailsPtrOutput() OrderItemDetailsPtrOutput {
	return i.ToOrderItemDetailsPtrOutputWithContext(context.Background())
}

func (i OrderItemDetailsArgs) ToOrderItemDetailsPtrOutputWithContext(ctx context.Context) OrderItemDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderItemDetailsOutput).ToOrderItemDetailsPtrOutputWithContext(ctx)
}









type OrderItemDetailsPtrInput interface {
	pulumi.Input

	ToOrderItemDetailsPtrOutput() OrderItemDetailsPtrOutput
	ToOrderItemDetailsPtrOutputWithContext(context.Context) OrderItemDetailsPtrOutput
}

type orderItemDetailsPtrType OrderItemDetailsArgs

func OrderItemDetailsPtr(v *OrderItemDetailsArgs) OrderItemDetailsPtrInput {
	return (*orderItemDetailsPtrType)(v)
}

func (*orderItemDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderItemDetails)(nil)).Elem()
}

func (i *orderItemDetailsPtrType) ToOrderItemDetailsPtrOutput() OrderItemDetailsPtrOutput {
	return i.ToOrderItemDetailsPtrOutputWithContext(context.Background())
}

func (i *orderItemDetailsPtrType) ToOrderItemDetailsPtrOutputWithContext(ctx context.Context) OrderItemDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderItemDetailsPtrOutput)
}

type OrderItemDetailsOutput struct{ *pulumi.OutputState }

func (OrderItemDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderItemDetails)(nil)).Elem()
}

func (o OrderItemDetailsOutput) ToOrderItemDetailsOutput() OrderItemDetailsOutput {
	return o
}

func (o OrderItemDetailsOutput) ToOrderItemDetailsOutputWithContext(ctx context.Context) OrderItemDetailsOutput {
	return o
}

func (o OrderItemDetailsOutput) ToOrderItemDetailsPtrOutput() OrderItemDetailsPtrOutput {
	return o.ToOrderItemDetailsPtrOutputWithContext(context.Background())
}

func (o OrderItemDetailsOutput) ToOrderItemDetailsPtrOutputWithContext(ctx context.Context) OrderItemDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrderItemDetails) *OrderItemDetails {
		return &v
	}).(OrderItemDetailsPtrOutput)
}

func (o OrderItemDetailsOutput) NotificationEmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OrderItemDetails) []string { return v.NotificationEmailList }).(pulumi.StringArrayOutput)
}

func (o OrderItemDetailsOutput) OrderItemType() pulumi.StringOutput {
	return o.ApplyT(func(v OrderItemDetails) string { return v.OrderItemType }).(pulumi.StringOutput)
}

func (o OrderItemDetailsOutput) Preferences() PreferencesPtrOutput {
	return o.ApplyT(func(v OrderItemDetails) *Preferences { return v.Preferences }).(PreferencesPtrOutput)
}

func (o OrderItemDetailsOutput) ProductDetails() ProductDetailsOutput {
	return o.ApplyT(func(v OrderItemDetails) ProductDetails { return v.ProductDetails }).(ProductDetailsOutput)
}

type OrderItemDetailsPtrOutput struct{ *pulumi.OutputState }

func (OrderItemDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderItemDetails)(nil)).Elem()
}

func (o OrderItemDetailsPtrOutput) ToOrderItemDetailsPtrOutput() OrderItemDetailsPtrOutput {
	return o
}

func (o OrderItemDetailsPtrOutput) ToOrderItemDetailsPtrOutputWithContext(ctx context.Context) OrderItemDetailsPtrOutput {
	return o
}

func (o OrderItemDetailsPtrOutput) Elem() OrderItemDetailsOutput {
	return o.ApplyT(func(v *OrderItemDetails) OrderItemDetails {
		if v != nil {
			return *v
		}
		var ret OrderItemDetails
		return ret
	}).(OrderItemDetailsOutput)
}

func (o OrderItemDetailsPtrOutput) NotificationEmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrderItemDetails) []string {
		if v == nil {
			return nil
		}
		return v.NotificationEmailList
	}).(pulumi.StringArrayOutput)
}

func (o OrderItemDetailsPtrOutput) OrderItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderItemDetails) *string {
		if v == nil {
			return nil
		}
		return &v.OrderItemType
	}).(pulumi.StringPtrOutput)
}

func (o OrderItemDetailsPtrOutput) Preferences() PreferencesPtrOutput {
	return o.ApplyT(func(v *OrderItemDetails) *Preferences {
		if v == nil {
			return nil
		}
		return v.Preferences
	}).(PreferencesPtrOutput)
}

func (o OrderItemDetailsPtrOutput) ProductDetails() ProductDetailsPtrOutput {
	return o.ApplyT(func(v *OrderItemDetails) *ProductDetails {
		if v == nil {
			return nil
		}
		return &v.ProductDetails
	}).(ProductDetailsPtrOutput)
}

type OrderItemDetailsResponse struct {
	CancellationReason     string                  `pulumi:"cancellationReason"`
	CancellationStatus     string                  `pulumi:"cancellationStatus"`
	CurrentStage           StageDetailsResponse    `pulumi:"currentStage"`
	DeletionStatus         string                  `pulumi:"deletionStatus"`
	Error                  ErrorDetailResponse     `pulumi:"error"`
	ForwardShippingDetails ShippingDetailsResponse `pulumi:"forwardShippingDetails"`
	ManagementRpDetails    interface{}             `pulumi:"managementRpDetails"`
	NotificationEmailList  []string                `pulumi:"notificationEmailList"`
	OrderItemStageHistory  []StageDetailsResponse  `pulumi:"orderItemStageHistory"`
	OrderItemType          string                  `pulumi:"orderItemType"`
	Preferences            *PreferencesResponse    `pulumi:"preferences"`
	ProductDetails         ProductDetailsResponse  `pulumi:"productDetails"`
	ReturnReason           string                  `pulumi:"returnReason"`
	ReturnStatus           string                  `pulumi:"returnStatus"`
	ReverseShippingDetails ShippingDetailsResponse `pulumi:"reverseShippingDetails"`
}





type OrderItemDetailsResponseInput interface {
	pulumi.Input

	ToOrderItemDetailsResponseOutput() OrderItemDetailsResponseOutput
	ToOrderItemDetailsResponseOutputWithContext(context.Context) OrderItemDetailsResponseOutput
}

type OrderItemDetailsResponseArgs struct {
	CancellationReason     pulumi.StringInput             `pulumi:"cancellationReason"`
	CancellationStatus     pulumi.StringInput             `pulumi:"cancellationStatus"`
	CurrentStage           StageDetailsResponseInput      `pulumi:"currentStage"`
	DeletionStatus         pulumi.StringInput             `pulumi:"deletionStatus"`
	Error                  ErrorDetailResponseInput       `pulumi:"error"`
	ForwardShippingDetails ShippingDetailsResponseInput   `pulumi:"forwardShippingDetails"`
	ManagementRpDetails    pulumi.Input                   `pulumi:"managementRpDetails"`
	NotificationEmailList  pulumi.StringArrayInput        `pulumi:"notificationEmailList"`
	OrderItemStageHistory  StageDetailsResponseArrayInput `pulumi:"orderItemStageHistory"`
	OrderItemType          pulumi.StringInput             `pulumi:"orderItemType"`
	Preferences            PreferencesResponsePtrInput    `pulumi:"preferences"`
	ProductDetails         ProductDetailsResponseInput    `pulumi:"productDetails"`
	ReturnReason           pulumi.StringInput             `pulumi:"returnReason"`
	ReturnStatus           pulumi.StringInput             `pulumi:"returnStatus"`
	ReverseShippingDetails ShippingDetailsResponseInput   `pulumi:"reverseShippingDetails"`
}

func (OrderItemDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderItemDetailsResponse)(nil)).Elem()
}

func (i OrderItemDetailsResponseArgs) ToOrderItemDetailsResponseOutput() OrderItemDetailsResponseOutput {
	return i.ToOrderItemDetailsResponseOutputWithContext(context.Background())
}

func (i OrderItemDetailsResponseArgs) ToOrderItemDetailsResponseOutputWithContext(ctx context.Context) OrderItemDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderItemDetailsResponseOutput)
}

func (i OrderItemDetailsResponseArgs) ToOrderItemDetailsResponsePtrOutput() OrderItemDetailsResponsePtrOutput {
	return i.ToOrderItemDetailsResponsePtrOutputWithContext(context.Background())
}

func (i OrderItemDetailsResponseArgs) ToOrderItemDetailsResponsePtrOutputWithContext(ctx context.Context) OrderItemDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderItemDetailsResponseOutput).ToOrderItemDetailsResponsePtrOutputWithContext(ctx)
}









type OrderItemDetailsResponsePtrInput interface {
	pulumi.Input

	ToOrderItemDetailsResponsePtrOutput() OrderItemDetailsResponsePtrOutput
	ToOrderItemDetailsResponsePtrOutputWithContext(context.Context) OrderItemDetailsResponsePtrOutput
}

type orderItemDetailsResponsePtrType OrderItemDetailsResponseArgs

func OrderItemDetailsResponsePtr(v *OrderItemDetailsResponseArgs) OrderItemDetailsResponsePtrInput {
	return (*orderItemDetailsResponsePtrType)(v)
}

func (*orderItemDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderItemDetailsResponse)(nil)).Elem()
}

func (i *orderItemDetailsResponsePtrType) ToOrderItemDetailsResponsePtrOutput() OrderItemDetailsResponsePtrOutput {
	return i.ToOrderItemDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *orderItemDetailsResponsePtrType) ToOrderItemDetailsResponsePtrOutputWithContext(ctx context.Context) OrderItemDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderItemDetailsResponsePtrOutput)
}

type OrderItemDetailsResponseOutput struct{ *pulumi.OutputState }

func (OrderItemDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderItemDetailsResponse)(nil)).Elem()
}

func (o OrderItemDetailsResponseOutput) ToOrderItemDetailsResponseOutput() OrderItemDetailsResponseOutput {
	return o
}

func (o OrderItemDetailsResponseOutput) ToOrderItemDetailsResponseOutputWithContext(ctx context.Context) OrderItemDetailsResponseOutput {
	return o
}

func (o OrderItemDetailsResponseOutput) ToOrderItemDetailsResponsePtrOutput() OrderItemDetailsResponsePtrOutput {
	return o.ToOrderItemDetailsResponsePtrOutputWithContext(context.Background())
}

func (o OrderItemDetailsResponseOutput) ToOrderItemDetailsResponsePtrOutputWithContext(ctx context.Context) OrderItemDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrderItemDetailsResponse) *OrderItemDetailsResponse {
		return &v
	}).(OrderItemDetailsResponsePtrOutput)
}

func (o OrderItemDetailsResponseOutput) CancellationReason() pulumi.StringOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) string { return v.CancellationReason }).(pulumi.StringOutput)
}

func (o OrderItemDetailsResponseOutput) CancellationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) string { return v.CancellationStatus }).(pulumi.StringOutput)
}

func (o OrderItemDetailsResponseOutput) CurrentStage() StageDetailsResponseOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) StageDetailsResponse { return v.CurrentStage }).(StageDetailsResponseOutput)
}

func (o OrderItemDetailsResponseOutput) DeletionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) string { return v.DeletionStatus }).(pulumi.StringOutput)
}

func (o OrderItemDetailsResponseOutput) Error() ErrorDetailResponseOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) ErrorDetailResponse { return v.Error }).(ErrorDetailResponseOutput)
}

func (o OrderItemDetailsResponseOutput) ForwardShippingDetails() ShippingDetailsResponseOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) ShippingDetailsResponse { return v.ForwardShippingDetails }).(ShippingDetailsResponseOutput)
}

func (o OrderItemDetailsResponseOutput) ManagementRpDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) interface{} { return v.ManagementRpDetails }).(pulumi.AnyOutput)
}

func (o OrderItemDetailsResponseOutput) NotificationEmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) []string { return v.NotificationEmailList }).(pulumi.StringArrayOutput)
}

func (o OrderItemDetailsResponseOutput) OrderItemStageHistory() StageDetailsResponseArrayOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) []StageDetailsResponse { return v.OrderItemStageHistory }).(StageDetailsResponseArrayOutput)
}

func (o OrderItemDetailsResponseOutput) OrderItemType() pulumi.StringOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) string { return v.OrderItemType }).(pulumi.StringOutput)
}

func (o OrderItemDetailsResponseOutput) Preferences() PreferencesResponsePtrOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) *PreferencesResponse { return v.Preferences }).(PreferencesResponsePtrOutput)
}

func (o OrderItemDetailsResponseOutput) ProductDetails() ProductDetailsResponseOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) ProductDetailsResponse { return v.ProductDetails }).(ProductDetailsResponseOutput)
}

func (o OrderItemDetailsResponseOutput) ReturnReason() pulumi.StringOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) string { return v.ReturnReason }).(pulumi.StringOutput)
}

func (o OrderItemDetailsResponseOutput) ReturnStatus() pulumi.StringOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) string { return v.ReturnStatus }).(pulumi.StringOutput)
}

func (o OrderItemDetailsResponseOutput) ReverseShippingDetails() ShippingDetailsResponseOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) ShippingDetailsResponse { return v.ReverseShippingDetails }).(ShippingDetailsResponseOutput)
}

type OrderItemDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (OrderItemDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderItemDetailsResponse)(nil)).Elem()
}

func (o OrderItemDetailsResponsePtrOutput) ToOrderItemDetailsResponsePtrOutput() OrderItemDetailsResponsePtrOutput {
	return o
}

func (o OrderItemDetailsResponsePtrOutput) ToOrderItemDetailsResponsePtrOutputWithContext(ctx context.Context) OrderItemDetailsResponsePtrOutput {
	return o
}

func (o OrderItemDetailsResponsePtrOutput) Elem() OrderItemDetailsResponseOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) OrderItemDetailsResponse {
		if v != nil {
			return *v
		}
		var ret OrderItemDetailsResponse
		return ret
	}).(OrderItemDetailsResponseOutput)
}

func (o OrderItemDetailsResponsePtrOutput) CancellationReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CancellationReason
	}).(pulumi.StringPtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) CancellationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CancellationStatus
	}).(pulumi.StringPtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) CurrentStage() StageDetailsResponsePtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *StageDetailsResponse {
		if v == nil {
			return nil
		}
		return &v.CurrentStage
	}).(StageDetailsResponsePtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) DeletionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DeletionStatus
	}).(pulumi.StringPtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) Error() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return &v.Error
	}).(ErrorDetailResponsePtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) ForwardShippingDetails() ShippingDetailsResponsePtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *ShippingDetailsResponse {
		if v == nil {
			return nil
		}
		return &v.ForwardShippingDetails
	}).(ShippingDetailsResponsePtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) ManagementRpDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.ManagementRpDetails
	}).(pulumi.AnyOutput)
}

func (o OrderItemDetailsResponsePtrOutput) NotificationEmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.NotificationEmailList
	}).(pulumi.StringArrayOutput)
}

func (o OrderItemDetailsResponsePtrOutput) OrderItemStageHistory() StageDetailsResponseArrayOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) []StageDetailsResponse {
		if v == nil {
			return nil
		}
		return v.OrderItemStageHistory
	}).(StageDetailsResponseArrayOutput)
}

func (o OrderItemDetailsResponsePtrOutput) OrderItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OrderItemType
	}).(pulumi.StringPtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) Preferences() PreferencesResponsePtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *PreferencesResponse {
		if v == nil {
			return nil
		}
		return v.Preferences
	}).(PreferencesResponsePtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) ProductDetails() ProductDetailsResponsePtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *ProductDetailsResponse {
		if v == nil {
			return nil
		}
		return &v.ProductDetails
	}).(ProductDetailsResponsePtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) ReturnReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ReturnReason
	}).(pulumi.StringPtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) ReturnStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ReturnStatus
	}).(pulumi.StringPtrOutput)
}

func (o OrderItemDetailsResponsePtrOutput) ReverseShippingDetails() ShippingDetailsResponsePtrOutput {
	return o.ApplyT(func(v *OrderItemDetailsResponse) *ShippingDetailsResponse {
		if v == nil {
			return nil
		}
		return &v.ReverseShippingDetails
	}).(ShippingDetailsResponsePtrOutput)
}

type Pav2MeterDetailsResponse struct {
	BillingType  string  `pulumi:"billingType"`
	ChargingType string  `pulumi:"chargingType"`
	MeterGuid    string  `pulumi:"meterGuid"`
	Multiplier   float64 `pulumi:"multiplier"`
}





type Pav2MeterDetailsResponseInput interface {
	pulumi.Input

	ToPav2MeterDetailsResponseOutput() Pav2MeterDetailsResponseOutput
	ToPav2MeterDetailsResponseOutputWithContext(context.Context) Pav2MeterDetailsResponseOutput
}

type Pav2MeterDetailsResponseArgs struct {
	BillingType  pulumi.StringInput  `pulumi:"billingType"`
	ChargingType pulumi.StringInput  `pulumi:"chargingType"`
	MeterGuid    pulumi.StringInput  `pulumi:"meterGuid"`
	Multiplier   pulumi.Float64Input `pulumi:"multiplier"`
}

func (Pav2MeterDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Pav2MeterDetailsResponse)(nil)).Elem()
}

func (i Pav2MeterDetailsResponseArgs) ToPav2MeterDetailsResponseOutput() Pav2MeterDetailsResponseOutput {
	return i.ToPav2MeterDetailsResponseOutputWithContext(context.Background())
}

func (i Pav2MeterDetailsResponseArgs) ToPav2MeterDetailsResponseOutputWithContext(ctx context.Context) Pav2MeterDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Pav2MeterDetailsResponseOutput)
}

type Pav2MeterDetailsResponseOutput struct{ *pulumi.OutputState }

func (Pav2MeterDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pav2MeterDetailsResponse)(nil)).Elem()
}

func (o Pav2MeterDetailsResponseOutput) ToPav2MeterDetailsResponseOutput() Pav2MeterDetailsResponseOutput {
	return o
}

func (o Pav2MeterDetailsResponseOutput) ToPav2MeterDetailsResponseOutputWithContext(ctx context.Context) Pav2MeterDetailsResponseOutput {
	return o
}

func (o Pav2MeterDetailsResponseOutput) BillingType() pulumi.StringOutput {
	return o.ApplyT(func(v Pav2MeterDetailsResponse) string { return v.BillingType }).(pulumi.StringOutput)
}

func (o Pav2MeterDetailsResponseOutput) ChargingType() pulumi.StringOutput {
	return o.ApplyT(func(v Pav2MeterDetailsResponse) string { return v.ChargingType }).(pulumi.StringOutput)
}

func (o Pav2MeterDetailsResponseOutput) MeterGuid() pulumi.StringOutput {
	return o.ApplyT(func(v Pav2MeterDetailsResponse) string { return v.MeterGuid }).(pulumi.StringOutput)
}

func (o Pav2MeterDetailsResponseOutput) Multiplier() pulumi.Float64Output {
	return o.ApplyT(func(v Pav2MeterDetailsResponse) float64 { return v.Multiplier }).(pulumi.Float64Output)
}

type Preferences struct {
	EncryptionPreferences         *EncryptionPreferences         `pulumi:"encryptionPreferences"`
	ManagementResourcePreferences *ManagementResourcePreferences `pulumi:"managementResourcePreferences"`
	NotificationPreferences       []NotificationPreference       `pulumi:"notificationPreferences"`
	TransportPreferences          *TransportPreferences          `pulumi:"transportPreferences"`
}





type PreferencesInput interface {
	pulumi.Input

	ToPreferencesOutput() PreferencesOutput
	ToPreferencesOutputWithContext(context.Context) PreferencesOutput
}

type PreferencesArgs struct {
	EncryptionPreferences         EncryptionPreferencesPtrInput         `pulumi:"encryptionPreferences"`
	ManagementResourcePreferences ManagementResourcePreferencesPtrInput `pulumi:"managementResourcePreferences"`
	NotificationPreferences       NotificationPreferenceArrayInput      `pulumi:"notificationPreferences"`
	TransportPreferences          TransportPreferencesPtrInput          `pulumi:"transportPreferences"`
}

func (PreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Preferences)(nil)).Elem()
}

func (i PreferencesArgs) ToPreferencesOutput() PreferencesOutput {
	return i.ToPreferencesOutputWithContext(context.Background())
}

func (i PreferencesArgs) ToPreferencesOutputWithContext(ctx context.Context) PreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesOutput)
}

func (i PreferencesArgs) ToPreferencesPtrOutput() PreferencesPtrOutput {
	return i.ToPreferencesPtrOutputWithContext(context.Background())
}

func (i PreferencesArgs) ToPreferencesPtrOutputWithContext(ctx context.Context) PreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesOutput).ToPreferencesPtrOutputWithContext(ctx)
}









type PreferencesPtrInput interface {
	pulumi.Input

	ToPreferencesPtrOutput() PreferencesPtrOutput
	ToPreferencesPtrOutputWithContext(context.Context) PreferencesPtrOutput
}

type preferencesPtrType PreferencesArgs

func PreferencesPtr(v *PreferencesArgs) PreferencesPtrInput {
	return (*preferencesPtrType)(v)
}

func (*preferencesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Preferences)(nil)).Elem()
}

func (i *preferencesPtrType) ToPreferencesPtrOutput() PreferencesPtrOutput {
	return i.ToPreferencesPtrOutputWithContext(context.Background())
}

func (i *preferencesPtrType) ToPreferencesPtrOutputWithContext(ctx context.Context) PreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesPtrOutput)
}

type PreferencesOutput struct{ *pulumi.OutputState }

func (PreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Preferences)(nil)).Elem()
}

func (o PreferencesOutput) ToPreferencesOutput() PreferencesOutput {
	return o
}

func (o PreferencesOutput) ToPreferencesOutputWithContext(ctx context.Context) PreferencesOutput {
	return o
}

func (o PreferencesOutput) ToPreferencesPtrOutput() PreferencesPtrOutput {
	return o.ToPreferencesPtrOutputWithContext(context.Background())
}

func (o PreferencesOutput) ToPreferencesPtrOutputWithContext(ctx context.Context) PreferencesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Preferences) *Preferences {
		return &v
	}).(PreferencesPtrOutput)
}

func (o PreferencesOutput) EncryptionPreferences() EncryptionPreferencesPtrOutput {
	return o.ApplyT(func(v Preferences) *EncryptionPreferences { return v.EncryptionPreferences }).(EncryptionPreferencesPtrOutput)
}

func (o PreferencesOutput) ManagementResourcePreferences() ManagementResourcePreferencesPtrOutput {
	return o.ApplyT(func(v Preferences) *ManagementResourcePreferences { return v.ManagementResourcePreferences }).(ManagementResourcePreferencesPtrOutput)
}

func (o PreferencesOutput) NotificationPreferences() NotificationPreferenceArrayOutput {
	return o.ApplyT(func(v Preferences) []NotificationPreference { return v.NotificationPreferences }).(NotificationPreferenceArrayOutput)
}

func (o PreferencesOutput) TransportPreferences() TransportPreferencesPtrOutput {
	return o.ApplyT(func(v Preferences) *TransportPreferences { return v.TransportPreferences }).(TransportPreferencesPtrOutput)
}

type PreferencesPtrOutput struct{ *pulumi.OutputState }

func (PreferencesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Preferences)(nil)).Elem()
}

func (o PreferencesPtrOutput) ToPreferencesPtrOutput() PreferencesPtrOutput {
	return o
}

func (o PreferencesPtrOutput) ToPreferencesPtrOutputWithContext(ctx context.Context) PreferencesPtrOutput {
	return o
}

func (o PreferencesPtrOutput) Elem() PreferencesOutput {
	return o.ApplyT(func(v *Preferences) Preferences {
		if v != nil {
			return *v
		}
		var ret Preferences
		return ret
	}).(PreferencesOutput)
}

func (o PreferencesPtrOutput) EncryptionPreferences() EncryptionPreferencesPtrOutput {
	return o.ApplyT(func(v *Preferences) *EncryptionPreferences {
		if v == nil {
			return nil
		}
		return v.EncryptionPreferences
	}).(EncryptionPreferencesPtrOutput)
}

func (o PreferencesPtrOutput) ManagementResourcePreferences() ManagementResourcePreferencesPtrOutput {
	return o.ApplyT(func(v *Preferences) *ManagementResourcePreferences {
		if v == nil {
			return nil
		}
		return v.ManagementResourcePreferences
	}).(ManagementResourcePreferencesPtrOutput)
}

func (o PreferencesPtrOutput) NotificationPreferences() NotificationPreferenceArrayOutput {
	return o.ApplyT(func(v *Preferences) []NotificationPreference {
		if v == nil {
			return nil
		}
		return v.NotificationPreferences
	}).(NotificationPreferenceArrayOutput)
}

func (o PreferencesPtrOutput) TransportPreferences() TransportPreferencesPtrOutput {
	return o.ApplyT(func(v *Preferences) *TransportPreferences {
		if v == nil {
			return nil
		}
		return v.TransportPreferences
	}).(TransportPreferencesPtrOutput)
}

type PreferencesResponse struct {
	EncryptionPreferences         *EncryptionPreferencesResponse         `pulumi:"encryptionPreferences"`
	ManagementResourcePreferences *ManagementResourcePreferencesResponse `pulumi:"managementResourcePreferences"`
	NotificationPreferences       []NotificationPreferenceResponse       `pulumi:"notificationPreferences"`
	TransportPreferences          *TransportPreferencesResponse          `pulumi:"transportPreferences"`
}





type PreferencesResponseInput interface {
	pulumi.Input

	ToPreferencesResponseOutput() PreferencesResponseOutput
	ToPreferencesResponseOutputWithContext(context.Context) PreferencesResponseOutput
}

type PreferencesResponseArgs struct {
	EncryptionPreferences         EncryptionPreferencesResponsePtrInput         `pulumi:"encryptionPreferences"`
	ManagementResourcePreferences ManagementResourcePreferencesResponsePtrInput `pulumi:"managementResourcePreferences"`
	NotificationPreferences       NotificationPreferenceResponseArrayInput      `pulumi:"notificationPreferences"`
	TransportPreferences          TransportPreferencesResponsePtrInput          `pulumi:"transportPreferences"`
}

func (PreferencesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PreferencesResponse)(nil)).Elem()
}

func (i PreferencesResponseArgs) ToPreferencesResponseOutput() PreferencesResponseOutput {
	return i.ToPreferencesResponseOutputWithContext(context.Background())
}

func (i PreferencesResponseArgs) ToPreferencesResponseOutputWithContext(ctx context.Context) PreferencesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesResponseOutput)
}

func (i PreferencesResponseArgs) ToPreferencesResponsePtrOutput() PreferencesResponsePtrOutput {
	return i.ToPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i PreferencesResponseArgs) ToPreferencesResponsePtrOutputWithContext(ctx context.Context) PreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesResponseOutput).ToPreferencesResponsePtrOutputWithContext(ctx)
}









type PreferencesResponsePtrInput interface {
	pulumi.Input

	ToPreferencesResponsePtrOutput() PreferencesResponsePtrOutput
	ToPreferencesResponsePtrOutputWithContext(context.Context) PreferencesResponsePtrOutput
}

type preferencesResponsePtrType PreferencesResponseArgs

func PreferencesResponsePtr(v *PreferencesResponseArgs) PreferencesResponsePtrInput {
	return (*preferencesResponsePtrType)(v)
}

func (*preferencesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PreferencesResponse)(nil)).Elem()
}

func (i *preferencesResponsePtrType) ToPreferencesResponsePtrOutput() PreferencesResponsePtrOutput {
	return i.ToPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i *preferencesResponsePtrType) ToPreferencesResponsePtrOutputWithContext(ctx context.Context) PreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesResponsePtrOutput)
}

type PreferencesResponseOutput struct{ *pulumi.OutputState }

func (PreferencesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PreferencesResponse)(nil)).Elem()
}

func (o PreferencesResponseOutput) ToPreferencesResponseOutput() PreferencesResponseOutput {
	return o
}

func (o PreferencesResponseOutput) ToPreferencesResponseOutputWithContext(ctx context.Context) PreferencesResponseOutput {
	return o
}

func (o PreferencesResponseOutput) ToPreferencesResponsePtrOutput() PreferencesResponsePtrOutput {
	return o.ToPreferencesResponsePtrOutputWithContext(context.Background())
}

func (o PreferencesResponseOutput) ToPreferencesResponsePtrOutputWithContext(ctx context.Context) PreferencesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PreferencesResponse) *PreferencesResponse {
		return &v
	}).(PreferencesResponsePtrOutput)
}

func (o PreferencesResponseOutput) EncryptionPreferences() EncryptionPreferencesResponsePtrOutput {
	return o.ApplyT(func(v PreferencesResponse) *EncryptionPreferencesResponse { return v.EncryptionPreferences }).(EncryptionPreferencesResponsePtrOutput)
}

func (o PreferencesResponseOutput) ManagementResourcePreferences() ManagementResourcePreferencesResponsePtrOutput {
	return o.ApplyT(func(v PreferencesResponse) *ManagementResourcePreferencesResponse {
		return v.ManagementResourcePreferences
	}).(ManagementResourcePreferencesResponsePtrOutput)
}

func (o PreferencesResponseOutput) NotificationPreferences() NotificationPreferenceResponseArrayOutput {
	return o.ApplyT(func(v PreferencesResponse) []NotificationPreferenceResponse { return v.NotificationPreferences }).(NotificationPreferenceResponseArrayOutput)
}

func (o PreferencesResponseOutput) TransportPreferences() TransportPreferencesResponsePtrOutput {
	return o.ApplyT(func(v PreferencesResponse) *TransportPreferencesResponse { return v.TransportPreferences }).(TransportPreferencesResponsePtrOutput)
}

type PreferencesResponsePtrOutput struct{ *pulumi.OutputState }

func (PreferencesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PreferencesResponse)(nil)).Elem()
}

func (o PreferencesResponsePtrOutput) ToPreferencesResponsePtrOutput() PreferencesResponsePtrOutput {
	return o
}

func (o PreferencesResponsePtrOutput) ToPreferencesResponsePtrOutputWithContext(ctx context.Context) PreferencesResponsePtrOutput {
	return o
}

func (o PreferencesResponsePtrOutput) Elem() PreferencesResponseOutput {
	return o.ApplyT(func(v *PreferencesResponse) PreferencesResponse {
		if v != nil {
			return *v
		}
		var ret PreferencesResponse
		return ret
	}).(PreferencesResponseOutput)
}

func (o PreferencesResponsePtrOutput) EncryptionPreferences() EncryptionPreferencesResponsePtrOutput {
	return o.ApplyT(func(v *PreferencesResponse) *EncryptionPreferencesResponse {
		if v == nil {
			return nil
		}
		return v.EncryptionPreferences
	}).(EncryptionPreferencesResponsePtrOutput)
}

func (o PreferencesResponsePtrOutput) ManagementResourcePreferences() ManagementResourcePreferencesResponsePtrOutput {
	return o.ApplyT(func(v *PreferencesResponse) *ManagementResourcePreferencesResponse {
		if v == nil {
			return nil
		}
		return v.ManagementResourcePreferences
	}).(ManagementResourcePreferencesResponsePtrOutput)
}

func (o PreferencesResponsePtrOutput) NotificationPreferences() NotificationPreferenceResponseArrayOutput {
	return o.ApplyT(func(v *PreferencesResponse) []NotificationPreferenceResponse {
		if v == nil {
			return nil
		}
		return v.NotificationPreferences
	}).(NotificationPreferenceResponseArrayOutput)
}

func (o PreferencesResponsePtrOutput) TransportPreferences() TransportPreferencesResponsePtrOutput {
	return o.ApplyT(func(v *PreferencesResponse) *TransportPreferencesResponse {
		if v == nil {
			return nil
		}
		return v.TransportPreferences
	}).(TransportPreferencesResponsePtrOutput)
}

type ProductDetails struct {
	Count                *int                 `pulumi:"count"`
	HierarchyInformation HierarchyInformation `pulumi:"hierarchyInformation"`
}





type ProductDetailsInput interface {
	pulumi.Input

	ToProductDetailsOutput() ProductDetailsOutput
	ToProductDetailsOutputWithContext(context.Context) ProductDetailsOutput
}

type ProductDetailsArgs struct {
	Count                pulumi.IntPtrInput        `pulumi:"count"`
	HierarchyInformation HierarchyInformationInput `pulumi:"hierarchyInformation"`
}

func (ProductDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductDetails)(nil)).Elem()
}

func (i ProductDetailsArgs) ToProductDetailsOutput() ProductDetailsOutput {
	return i.ToProductDetailsOutputWithContext(context.Background())
}

func (i ProductDetailsArgs) ToProductDetailsOutputWithContext(ctx context.Context) ProductDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductDetailsOutput)
}

func (i ProductDetailsArgs) ToProductDetailsPtrOutput() ProductDetailsPtrOutput {
	return i.ToProductDetailsPtrOutputWithContext(context.Background())
}

func (i ProductDetailsArgs) ToProductDetailsPtrOutputWithContext(ctx context.Context) ProductDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductDetailsOutput).ToProductDetailsPtrOutputWithContext(ctx)
}









type ProductDetailsPtrInput interface {
	pulumi.Input

	ToProductDetailsPtrOutput() ProductDetailsPtrOutput
	ToProductDetailsPtrOutputWithContext(context.Context) ProductDetailsPtrOutput
}

type productDetailsPtrType ProductDetailsArgs

func ProductDetailsPtr(v *ProductDetailsArgs) ProductDetailsPtrInput {
	return (*productDetailsPtrType)(v)
}

func (*productDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductDetails)(nil)).Elem()
}

func (i *productDetailsPtrType) ToProductDetailsPtrOutput() ProductDetailsPtrOutput {
	return i.ToProductDetailsPtrOutputWithContext(context.Background())
}

func (i *productDetailsPtrType) ToProductDetailsPtrOutputWithContext(ctx context.Context) ProductDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductDetailsPtrOutput)
}

type ProductDetailsOutput struct{ *pulumi.OutputState }

func (ProductDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductDetails)(nil)).Elem()
}

func (o ProductDetailsOutput) ToProductDetailsOutput() ProductDetailsOutput {
	return o
}

func (o ProductDetailsOutput) ToProductDetailsOutputWithContext(ctx context.Context) ProductDetailsOutput {
	return o
}

func (o ProductDetailsOutput) ToProductDetailsPtrOutput() ProductDetailsPtrOutput {
	return o.ToProductDetailsPtrOutputWithContext(context.Background())
}

func (o ProductDetailsOutput) ToProductDetailsPtrOutputWithContext(ctx context.Context) ProductDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProductDetails) *ProductDetails {
		return &v
	}).(ProductDetailsPtrOutput)
}

func (o ProductDetailsOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProductDetails) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ProductDetailsOutput) HierarchyInformation() HierarchyInformationOutput {
	return o.ApplyT(func(v ProductDetails) HierarchyInformation { return v.HierarchyInformation }).(HierarchyInformationOutput)
}

type ProductDetailsPtrOutput struct{ *pulumi.OutputState }

func (ProductDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductDetails)(nil)).Elem()
}

func (o ProductDetailsPtrOutput) ToProductDetailsPtrOutput() ProductDetailsPtrOutput {
	return o
}

func (o ProductDetailsPtrOutput) ToProductDetailsPtrOutputWithContext(ctx context.Context) ProductDetailsPtrOutput {
	return o
}

func (o ProductDetailsPtrOutput) Elem() ProductDetailsOutput {
	return o.ApplyT(func(v *ProductDetails) ProductDetails {
		if v != nil {
			return *v
		}
		var ret ProductDetails
		return ret
	}).(ProductDetailsOutput)
}

func (o ProductDetailsPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProductDetails) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o ProductDetailsPtrOutput) HierarchyInformation() HierarchyInformationPtrOutput {
	return o.ApplyT(func(v *ProductDetails) *HierarchyInformation {
		if v == nil {
			return nil
		}
		return &v.HierarchyInformation
	}).(HierarchyInformationPtrOutput)
}

type ProductDetailsResponse struct {
	Count                *int                         `pulumi:"count"`
	DeviceDetails        []DeviceDetailsResponse      `pulumi:"deviceDetails"`
	DisplayInfo          *DisplayInfoResponse         `pulumi:"displayInfo"`
	HierarchyInformation HierarchyInformationResponse `pulumi:"hierarchyInformation"`
}





type ProductDetailsResponseInput interface {
	pulumi.Input

	ToProductDetailsResponseOutput() ProductDetailsResponseOutput
	ToProductDetailsResponseOutputWithContext(context.Context) ProductDetailsResponseOutput
}

type ProductDetailsResponseArgs struct {
	Count                pulumi.IntPtrInput                `pulumi:"count"`
	DeviceDetails        DeviceDetailsResponseArrayInput   `pulumi:"deviceDetails"`
	DisplayInfo          DisplayInfoResponsePtrInput       `pulumi:"displayInfo"`
	HierarchyInformation HierarchyInformationResponseInput `pulumi:"hierarchyInformation"`
}

func (ProductDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductDetailsResponse)(nil)).Elem()
}

func (i ProductDetailsResponseArgs) ToProductDetailsResponseOutput() ProductDetailsResponseOutput {
	return i.ToProductDetailsResponseOutputWithContext(context.Background())
}

func (i ProductDetailsResponseArgs) ToProductDetailsResponseOutputWithContext(ctx context.Context) ProductDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductDetailsResponseOutput)
}

func (i ProductDetailsResponseArgs) ToProductDetailsResponsePtrOutput() ProductDetailsResponsePtrOutput {
	return i.ToProductDetailsResponsePtrOutputWithContext(context.Background())
}

func (i ProductDetailsResponseArgs) ToProductDetailsResponsePtrOutputWithContext(ctx context.Context) ProductDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductDetailsResponseOutput).ToProductDetailsResponsePtrOutputWithContext(ctx)
}









type ProductDetailsResponsePtrInput interface {
	pulumi.Input

	ToProductDetailsResponsePtrOutput() ProductDetailsResponsePtrOutput
	ToProductDetailsResponsePtrOutputWithContext(context.Context) ProductDetailsResponsePtrOutput
}

type productDetailsResponsePtrType ProductDetailsResponseArgs

func ProductDetailsResponsePtr(v *ProductDetailsResponseArgs) ProductDetailsResponsePtrInput {
	return (*productDetailsResponsePtrType)(v)
}

func (*productDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductDetailsResponse)(nil)).Elem()
}

func (i *productDetailsResponsePtrType) ToProductDetailsResponsePtrOutput() ProductDetailsResponsePtrOutput {
	return i.ToProductDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *productDetailsResponsePtrType) ToProductDetailsResponsePtrOutputWithContext(ctx context.Context) ProductDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductDetailsResponsePtrOutput)
}

type ProductDetailsResponseOutput struct{ *pulumi.OutputState }

func (ProductDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductDetailsResponse)(nil)).Elem()
}

func (o ProductDetailsResponseOutput) ToProductDetailsResponseOutput() ProductDetailsResponseOutput {
	return o
}

func (o ProductDetailsResponseOutput) ToProductDetailsResponseOutputWithContext(ctx context.Context) ProductDetailsResponseOutput {
	return o
}

func (o ProductDetailsResponseOutput) ToProductDetailsResponsePtrOutput() ProductDetailsResponsePtrOutput {
	return o.ToProductDetailsResponsePtrOutputWithContext(context.Background())
}

func (o ProductDetailsResponseOutput) ToProductDetailsResponsePtrOutputWithContext(ctx context.Context) ProductDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProductDetailsResponse) *ProductDetailsResponse {
		return &v
	}).(ProductDetailsResponsePtrOutput)
}

func (o ProductDetailsResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProductDetailsResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ProductDetailsResponseOutput) DeviceDetails() DeviceDetailsResponseArrayOutput {
	return o.ApplyT(func(v ProductDetailsResponse) []DeviceDetailsResponse { return v.DeviceDetails }).(DeviceDetailsResponseArrayOutput)
}

func (o ProductDetailsResponseOutput) DisplayInfo() DisplayInfoResponsePtrOutput {
	return o.ApplyT(func(v ProductDetailsResponse) *DisplayInfoResponse { return v.DisplayInfo }).(DisplayInfoResponsePtrOutput)
}

func (o ProductDetailsResponseOutput) HierarchyInformation() HierarchyInformationResponseOutput {
	return o.ApplyT(func(v ProductDetailsResponse) HierarchyInformationResponse { return v.HierarchyInformation }).(HierarchyInformationResponseOutput)
}

type ProductDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ProductDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductDetailsResponse)(nil)).Elem()
}

func (o ProductDetailsResponsePtrOutput) ToProductDetailsResponsePtrOutput() ProductDetailsResponsePtrOutput {
	return o
}

func (o ProductDetailsResponsePtrOutput) ToProductDetailsResponsePtrOutputWithContext(ctx context.Context) ProductDetailsResponsePtrOutput {
	return o
}

func (o ProductDetailsResponsePtrOutput) Elem() ProductDetailsResponseOutput {
	return o.ApplyT(func(v *ProductDetailsResponse) ProductDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ProductDetailsResponse
		return ret
	}).(ProductDetailsResponseOutput)
}

func (o ProductDetailsResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProductDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o ProductDetailsResponsePtrOutput) DeviceDetails() DeviceDetailsResponseArrayOutput {
	return o.ApplyT(func(v *ProductDetailsResponse) []DeviceDetailsResponse {
		if v == nil {
			return nil
		}
		return v.DeviceDetails
	}).(DeviceDetailsResponseArrayOutput)
}

func (o ProductDetailsResponsePtrOutput) DisplayInfo() DisplayInfoResponsePtrOutput {
	return o.ApplyT(func(v *ProductDetailsResponse) *DisplayInfoResponse {
		if v == nil {
			return nil
		}
		return v.DisplayInfo
	}).(DisplayInfoResponsePtrOutput)
}

func (o ProductDetailsResponsePtrOutput) HierarchyInformation() HierarchyInformationResponsePtrOutput {
	return o.ApplyT(func(v *ProductDetailsResponse) *HierarchyInformationResponse {
		if v == nil {
			return nil
		}
		return &v.HierarchyInformation
	}).(HierarchyInformationResponsePtrOutput)
}

type ProductFamilyResponse struct {
	AvailabilityInformation AvailabilityInformationResponse `pulumi:"availabilityInformation"`
	CostInformation         CostInformationResponse         `pulumi:"costInformation"`
	Description             DescriptionResponse             `pulumi:"description"`
	DisplayName             string                          `pulumi:"displayName"`
	FilterableProperties    []FilterablePropertyResponse    `pulumi:"filterableProperties"`
	HierarchyInformation    HierarchyInformationResponse    `pulumi:"hierarchyInformation"`
	ImageInformation        []ImageInformationResponse      `pulumi:"imageInformation"`
	ProductLines            []ProductLineResponse           `pulumi:"productLines"`
}





type ProductFamilyResponseInput interface {
	pulumi.Input

	ToProductFamilyResponseOutput() ProductFamilyResponseOutput
	ToProductFamilyResponseOutputWithContext(context.Context) ProductFamilyResponseOutput
}

type ProductFamilyResponseArgs struct {
	AvailabilityInformation AvailabilityInformationResponseInput `pulumi:"availabilityInformation"`
	CostInformation         CostInformationResponseInput         `pulumi:"costInformation"`
	Description             DescriptionResponseInput             `pulumi:"description"`
	DisplayName             pulumi.StringInput                   `pulumi:"displayName"`
	FilterableProperties    FilterablePropertyResponseArrayInput `pulumi:"filterableProperties"`
	HierarchyInformation    HierarchyInformationResponseInput    `pulumi:"hierarchyInformation"`
	ImageInformation        ImageInformationResponseArrayInput   `pulumi:"imageInformation"`
	ProductLines            ProductLineResponseArrayInput        `pulumi:"productLines"`
}

func (ProductFamilyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductFamilyResponse)(nil)).Elem()
}

func (i ProductFamilyResponseArgs) ToProductFamilyResponseOutput() ProductFamilyResponseOutput {
	return i.ToProductFamilyResponseOutputWithContext(context.Background())
}

func (i ProductFamilyResponseArgs) ToProductFamilyResponseOutputWithContext(ctx context.Context) ProductFamilyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductFamilyResponseOutput)
}





type ProductFamilyResponseArrayInput interface {
	pulumi.Input

	ToProductFamilyResponseArrayOutput() ProductFamilyResponseArrayOutput
	ToProductFamilyResponseArrayOutputWithContext(context.Context) ProductFamilyResponseArrayOutput
}

type ProductFamilyResponseArray []ProductFamilyResponseInput

func (ProductFamilyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductFamilyResponse)(nil)).Elem()
}

func (i ProductFamilyResponseArray) ToProductFamilyResponseArrayOutput() ProductFamilyResponseArrayOutput {
	return i.ToProductFamilyResponseArrayOutputWithContext(context.Background())
}

func (i ProductFamilyResponseArray) ToProductFamilyResponseArrayOutputWithContext(ctx context.Context) ProductFamilyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductFamilyResponseArrayOutput)
}

type ProductFamilyResponseOutput struct{ *pulumi.OutputState }

func (ProductFamilyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductFamilyResponse)(nil)).Elem()
}

func (o ProductFamilyResponseOutput) ToProductFamilyResponseOutput() ProductFamilyResponseOutput {
	return o
}

func (o ProductFamilyResponseOutput) ToProductFamilyResponseOutputWithContext(ctx context.Context) ProductFamilyResponseOutput {
	return o
}

func (o ProductFamilyResponseOutput) AvailabilityInformation() AvailabilityInformationResponseOutput {
	return o.ApplyT(func(v ProductFamilyResponse) AvailabilityInformationResponse { return v.AvailabilityInformation }).(AvailabilityInformationResponseOutput)
}

func (o ProductFamilyResponseOutput) CostInformation() CostInformationResponseOutput {
	return o.ApplyT(func(v ProductFamilyResponse) CostInformationResponse { return v.CostInformation }).(CostInformationResponseOutput)
}

func (o ProductFamilyResponseOutput) Description() DescriptionResponseOutput {
	return o.ApplyT(func(v ProductFamilyResponse) DescriptionResponse { return v.Description }).(DescriptionResponseOutput)
}

func (o ProductFamilyResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ProductFamilyResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ProductFamilyResponseOutput) FilterableProperties() FilterablePropertyResponseArrayOutput {
	return o.ApplyT(func(v ProductFamilyResponse) []FilterablePropertyResponse { return v.FilterableProperties }).(FilterablePropertyResponseArrayOutput)
}

func (o ProductFamilyResponseOutput) HierarchyInformation() HierarchyInformationResponseOutput {
	return o.ApplyT(func(v ProductFamilyResponse) HierarchyInformationResponse { return v.HierarchyInformation }).(HierarchyInformationResponseOutput)
}

func (o ProductFamilyResponseOutput) ImageInformation() ImageInformationResponseArrayOutput {
	return o.ApplyT(func(v ProductFamilyResponse) []ImageInformationResponse { return v.ImageInformation }).(ImageInformationResponseArrayOutput)
}

func (o ProductFamilyResponseOutput) ProductLines() ProductLineResponseArrayOutput {
	return o.ApplyT(func(v ProductFamilyResponse) []ProductLineResponse { return v.ProductLines }).(ProductLineResponseArrayOutput)
}

type ProductFamilyResponseArrayOutput struct{ *pulumi.OutputState }

func (ProductFamilyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductFamilyResponse)(nil)).Elem()
}

func (o ProductFamilyResponseArrayOutput) ToProductFamilyResponseArrayOutput() ProductFamilyResponseArrayOutput {
	return o
}

func (o ProductFamilyResponseArrayOutput) ToProductFamilyResponseArrayOutputWithContext(ctx context.Context) ProductFamilyResponseArrayOutput {
	return o
}

func (o ProductFamilyResponseArrayOutput) Index(i pulumi.IntInput) ProductFamilyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProductFamilyResponse {
		return vs[0].([]ProductFamilyResponse)[vs[1].(int)]
	}).(ProductFamilyResponseOutput)
}

type ProductLineResponse struct {
	AvailabilityInformation AvailabilityInformationResponse `pulumi:"availabilityInformation"`
	CostInformation         CostInformationResponse         `pulumi:"costInformation"`
	Description             DescriptionResponse             `pulumi:"description"`
	DisplayName             string                          `pulumi:"displayName"`
	FilterableProperties    []FilterablePropertyResponse    `pulumi:"filterableProperties"`
	HierarchyInformation    HierarchyInformationResponse    `pulumi:"hierarchyInformation"`
	ImageInformation        []ImageInformationResponse      `pulumi:"imageInformation"`
	Products                []ProductResponse               `pulumi:"products"`
}





type ProductLineResponseInput interface {
	pulumi.Input

	ToProductLineResponseOutput() ProductLineResponseOutput
	ToProductLineResponseOutputWithContext(context.Context) ProductLineResponseOutput
}

type ProductLineResponseArgs struct {
	AvailabilityInformation AvailabilityInformationResponseInput `pulumi:"availabilityInformation"`
	CostInformation         CostInformationResponseInput         `pulumi:"costInformation"`
	Description             DescriptionResponseInput             `pulumi:"description"`
	DisplayName             pulumi.StringInput                   `pulumi:"displayName"`
	FilterableProperties    FilterablePropertyResponseArrayInput `pulumi:"filterableProperties"`
	HierarchyInformation    HierarchyInformationResponseInput    `pulumi:"hierarchyInformation"`
	ImageInformation        ImageInformationResponseArrayInput   `pulumi:"imageInformation"`
	Products                ProductResponseArrayInput            `pulumi:"products"`
}

func (ProductLineResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductLineResponse)(nil)).Elem()
}

func (i ProductLineResponseArgs) ToProductLineResponseOutput() ProductLineResponseOutput {
	return i.ToProductLineResponseOutputWithContext(context.Background())
}

func (i ProductLineResponseArgs) ToProductLineResponseOutputWithContext(ctx context.Context) ProductLineResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductLineResponseOutput)
}





type ProductLineResponseArrayInput interface {
	pulumi.Input

	ToProductLineResponseArrayOutput() ProductLineResponseArrayOutput
	ToProductLineResponseArrayOutputWithContext(context.Context) ProductLineResponseArrayOutput
}

type ProductLineResponseArray []ProductLineResponseInput

func (ProductLineResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductLineResponse)(nil)).Elem()
}

func (i ProductLineResponseArray) ToProductLineResponseArrayOutput() ProductLineResponseArrayOutput {
	return i.ToProductLineResponseArrayOutputWithContext(context.Background())
}

func (i ProductLineResponseArray) ToProductLineResponseArrayOutputWithContext(ctx context.Context) ProductLineResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductLineResponseArrayOutput)
}

type ProductLineResponseOutput struct{ *pulumi.OutputState }

func (ProductLineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductLineResponse)(nil)).Elem()
}

func (o ProductLineResponseOutput) ToProductLineResponseOutput() ProductLineResponseOutput {
	return o
}

func (o ProductLineResponseOutput) ToProductLineResponseOutputWithContext(ctx context.Context) ProductLineResponseOutput {
	return o
}

func (o ProductLineResponseOutput) AvailabilityInformation() AvailabilityInformationResponseOutput {
	return o.ApplyT(func(v ProductLineResponse) AvailabilityInformationResponse { return v.AvailabilityInformation }).(AvailabilityInformationResponseOutput)
}

func (o ProductLineResponseOutput) CostInformation() CostInformationResponseOutput {
	return o.ApplyT(func(v ProductLineResponse) CostInformationResponse { return v.CostInformation }).(CostInformationResponseOutput)
}

func (o ProductLineResponseOutput) Description() DescriptionResponseOutput {
	return o.ApplyT(func(v ProductLineResponse) DescriptionResponse { return v.Description }).(DescriptionResponseOutput)
}

func (o ProductLineResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ProductLineResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ProductLineResponseOutput) FilterableProperties() FilterablePropertyResponseArrayOutput {
	return o.ApplyT(func(v ProductLineResponse) []FilterablePropertyResponse { return v.FilterableProperties }).(FilterablePropertyResponseArrayOutput)
}

func (o ProductLineResponseOutput) HierarchyInformation() HierarchyInformationResponseOutput {
	return o.ApplyT(func(v ProductLineResponse) HierarchyInformationResponse { return v.HierarchyInformation }).(HierarchyInformationResponseOutput)
}

func (o ProductLineResponseOutput) ImageInformation() ImageInformationResponseArrayOutput {
	return o.ApplyT(func(v ProductLineResponse) []ImageInformationResponse { return v.ImageInformation }).(ImageInformationResponseArrayOutput)
}

func (o ProductLineResponseOutput) Products() ProductResponseArrayOutput {
	return o.ApplyT(func(v ProductLineResponse) []ProductResponse { return v.Products }).(ProductResponseArrayOutput)
}

type ProductLineResponseArrayOutput struct{ *pulumi.OutputState }

func (ProductLineResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductLineResponse)(nil)).Elem()
}

func (o ProductLineResponseArrayOutput) ToProductLineResponseArrayOutput() ProductLineResponseArrayOutput {
	return o
}

func (o ProductLineResponseArrayOutput) ToProductLineResponseArrayOutputWithContext(ctx context.Context) ProductLineResponseArrayOutput {
	return o
}

func (o ProductLineResponseArrayOutput) Index(i pulumi.IntInput) ProductLineResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProductLineResponse {
		return vs[0].([]ProductLineResponse)[vs[1].(int)]
	}).(ProductLineResponseOutput)
}

type ProductResponse struct {
	AvailabilityInformation AvailabilityInformationResponse `pulumi:"availabilityInformation"`
	Configurations          []ConfigurationResponse         `pulumi:"configurations"`
	CostInformation         CostInformationResponse         `pulumi:"costInformation"`
	Description             DescriptionResponse             `pulumi:"description"`
	DisplayName             string                          `pulumi:"displayName"`
	FilterableProperties    []FilterablePropertyResponse    `pulumi:"filterableProperties"`
	HierarchyInformation    HierarchyInformationResponse    `pulumi:"hierarchyInformation"`
	ImageInformation        []ImageInformationResponse      `pulumi:"imageInformation"`
}





type ProductResponseInput interface {
	pulumi.Input

	ToProductResponseOutput() ProductResponseOutput
	ToProductResponseOutputWithContext(context.Context) ProductResponseOutput
}

type ProductResponseArgs struct {
	AvailabilityInformation AvailabilityInformationResponseInput `pulumi:"availabilityInformation"`
	Configurations          ConfigurationResponseArrayInput      `pulumi:"configurations"`
	CostInformation         CostInformationResponseInput         `pulumi:"costInformation"`
	Description             DescriptionResponseInput             `pulumi:"description"`
	DisplayName             pulumi.StringInput                   `pulumi:"displayName"`
	FilterableProperties    FilterablePropertyResponseArrayInput `pulumi:"filterableProperties"`
	HierarchyInformation    HierarchyInformationResponseInput    `pulumi:"hierarchyInformation"`
	ImageInformation        ImageInformationResponseArrayInput   `pulumi:"imageInformation"`
}

func (ProductResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductResponse)(nil)).Elem()
}

func (i ProductResponseArgs) ToProductResponseOutput() ProductResponseOutput {
	return i.ToProductResponseOutputWithContext(context.Background())
}

func (i ProductResponseArgs) ToProductResponseOutputWithContext(ctx context.Context) ProductResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductResponseOutput)
}





type ProductResponseArrayInput interface {
	pulumi.Input

	ToProductResponseArrayOutput() ProductResponseArrayOutput
	ToProductResponseArrayOutputWithContext(context.Context) ProductResponseArrayOutput
}

type ProductResponseArray []ProductResponseInput

func (ProductResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductResponse)(nil)).Elem()
}

func (i ProductResponseArray) ToProductResponseArrayOutput() ProductResponseArrayOutput {
	return i.ToProductResponseArrayOutputWithContext(context.Background())
}

func (i ProductResponseArray) ToProductResponseArrayOutputWithContext(ctx context.Context) ProductResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductResponseArrayOutput)
}

type ProductResponseOutput struct{ *pulumi.OutputState }

func (ProductResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductResponse)(nil)).Elem()
}

func (o ProductResponseOutput) ToProductResponseOutput() ProductResponseOutput {
	return o
}

func (o ProductResponseOutput) ToProductResponseOutputWithContext(ctx context.Context) ProductResponseOutput {
	return o
}

func (o ProductResponseOutput) AvailabilityInformation() AvailabilityInformationResponseOutput {
	return o.ApplyT(func(v ProductResponse) AvailabilityInformationResponse { return v.AvailabilityInformation }).(AvailabilityInformationResponseOutput)
}

func (o ProductResponseOutput) Configurations() ConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ProductResponse) []ConfigurationResponse { return v.Configurations }).(ConfigurationResponseArrayOutput)
}

func (o ProductResponseOutput) CostInformation() CostInformationResponseOutput {
	return o.ApplyT(func(v ProductResponse) CostInformationResponse { return v.CostInformation }).(CostInformationResponseOutput)
}

func (o ProductResponseOutput) Description() DescriptionResponseOutput {
	return o.ApplyT(func(v ProductResponse) DescriptionResponse { return v.Description }).(DescriptionResponseOutput)
}

func (o ProductResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ProductResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ProductResponseOutput) FilterableProperties() FilterablePropertyResponseArrayOutput {
	return o.ApplyT(func(v ProductResponse) []FilterablePropertyResponse { return v.FilterableProperties }).(FilterablePropertyResponseArrayOutput)
}

func (o ProductResponseOutput) HierarchyInformation() HierarchyInformationResponseOutput {
	return o.ApplyT(func(v ProductResponse) HierarchyInformationResponse { return v.HierarchyInformation }).(HierarchyInformationResponseOutput)
}

func (o ProductResponseOutput) ImageInformation() ImageInformationResponseArrayOutput {
	return o.ApplyT(func(v ProductResponse) []ImageInformationResponse { return v.ImageInformation }).(ImageInformationResponseArrayOutput)
}

type ProductResponseArrayOutput struct{ *pulumi.OutputState }

func (ProductResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductResponse)(nil)).Elem()
}

func (o ProductResponseArrayOutput) ToProductResponseArrayOutput() ProductResponseArrayOutput {
	return o
}

func (o ProductResponseArrayOutput) ToProductResponseArrayOutputWithContext(ctx context.Context) ProductResponseArrayOutput {
	return o
}

func (o ProductResponseArrayOutput) Index(i pulumi.IntInput) ProductResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProductResponse {
		return vs[0].([]ProductResponse)[vs[1].(int)]
	}).(ProductResponseOutput)
}

type PurchaseMeterDetailsResponse struct {
	BillingType  string  `pulumi:"billingType"`
	ChargingType string  `pulumi:"chargingType"`
	Multiplier   float64 `pulumi:"multiplier"`
	ProductId    string  `pulumi:"productId"`
	SkuId        string  `pulumi:"skuId"`
	TermId       string  `pulumi:"termId"`
}





type PurchaseMeterDetailsResponseInput interface {
	pulumi.Input

	ToPurchaseMeterDetailsResponseOutput() PurchaseMeterDetailsResponseOutput
	ToPurchaseMeterDetailsResponseOutputWithContext(context.Context) PurchaseMeterDetailsResponseOutput
}

type PurchaseMeterDetailsResponseArgs struct {
	BillingType  pulumi.StringInput  `pulumi:"billingType"`
	ChargingType pulumi.StringInput  `pulumi:"chargingType"`
	Multiplier   pulumi.Float64Input `pulumi:"multiplier"`
	ProductId    pulumi.StringInput  `pulumi:"productId"`
	SkuId        pulumi.StringInput  `pulumi:"skuId"`
	TermId       pulumi.StringInput  `pulumi:"termId"`
}

func (PurchaseMeterDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PurchaseMeterDetailsResponse)(nil)).Elem()
}

func (i PurchaseMeterDetailsResponseArgs) ToPurchaseMeterDetailsResponseOutput() PurchaseMeterDetailsResponseOutput {
	return i.ToPurchaseMeterDetailsResponseOutputWithContext(context.Background())
}

func (i PurchaseMeterDetailsResponseArgs) ToPurchaseMeterDetailsResponseOutputWithContext(ctx context.Context) PurchaseMeterDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurchaseMeterDetailsResponseOutput)
}

type PurchaseMeterDetailsResponseOutput struct{ *pulumi.OutputState }

func (PurchaseMeterDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurchaseMeterDetailsResponse)(nil)).Elem()
}

func (o PurchaseMeterDetailsResponseOutput) ToPurchaseMeterDetailsResponseOutput() PurchaseMeterDetailsResponseOutput {
	return o
}

func (o PurchaseMeterDetailsResponseOutput) ToPurchaseMeterDetailsResponseOutputWithContext(ctx context.Context) PurchaseMeterDetailsResponseOutput {
	return o
}

func (o PurchaseMeterDetailsResponseOutput) BillingType() pulumi.StringOutput {
	return o.ApplyT(func(v PurchaseMeterDetailsResponse) string { return v.BillingType }).(pulumi.StringOutput)
}

func (o PurchaseMeterDetailsResponseOutput) ChargingType() pulumi.StringOutput {
	return o.ApplyT(func(v PurchaseMeterDetailsResponse) string { return v.ChargingType }).(pulumi.StringOutput)
}

func (o PurchaseMeterDetailsResponseOutput) Multiplier() pulumi.Float64Output {
	return o.ApplyT(func(v PurchaseMeterDetailsResponse) float64 { return v.Multiplier }).(pulumi.Float64Output)
}

func (o PurchaseMeterDetailsResponseOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v PurchaseMeterDetailsResponse) string { return v.ProductId }).(pulumi.StringOutput)
}

func (o PurchaseMeterDetailsResponseOutput) SkuId() pulumi.StringOutput {
	return o.ApplyT(func(v PurchaseMeterDetailsResponse) string { return v.SkuId }).(pulumi.StringOutput)
}

func (o PurchaseMeterDetailsResponseOutput) TermId() pulumi.StringOutput {
	return o.ApplyT(func(v PurchaseMeterDetailsResponse) string { return v.TermId }).(pulumi.StringOutput)
}

type ShippingAddress struct {
	AddressType     *string `pulumi:"addressType"`
	City            *string `pulumi:"city"`
	CompanyName     *string `pulumi:"companyName"`
	Country         string  `pulumi:"country"`
	PostalCode      *string `pulumi:"postalCode"`
	StateOrProvince *string `pulumi:"stateOrProvince"`
	StreetAddress1  string  `pulumi:"streetAddress1"`
	StreetAddress2  *string `pulumi:"streetAddress2"`
	StreetAddress3  *string `pulumi:"streetAddress3"`
	ZipExtendedCode *string `pulumi:"zipExtendedCode"`
}





type ShippingAddressInput interface {
	pulumi.Input

	ToShippingAddressOutput() ShippingAddressOutput
	ToShippingAddressOutputWithContext(context.Context) ShippingAddressOutput
}

type ShippingAddressArgs struct {
	AddressType     pulumi.StringPtrInput `pulumi:"addressType"`
	City            pulumi.StringPtrInput `pulumi:"city"`
	CompanyName     pulumi.StringPtrInput `pulumi:"companyName"`
	Country         pulumi.StringInput    `pulumi:"country"`
	PostalCode      pulumi.StringPtrInput `pulumi:"postalCode"`
	StateOrProvince pulumi.StringPtrInput `pulumi:"stateOrProvince"`
	StreetAddress1  pulumi.StringInput    `pulumi:"streetAddress1"`
	StreetAddress2  pulumi.StringPtrInput `pulumi:"streetAddress2"`
	StreetAddress3  pulumi.StringPtrInput `pulumi:"streetAddress3"`
	ZipExtendedCode pulumi.StringPtrInput `pulumi:"zipExtendedCode"`
}

func (ShippingAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingAddress)(nil)).Elem()
}

func (i ShippingAddressArgs) ToShippingAddressOutput() ShippingAddressOutput {
	return i.ToShippingAddressOutputWithContext(context.Background())
}

func (i ShippingAddressArgs) ToShippingAddressOutputWithContext(ctx context.Context) ShippingAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressOutput)
}

func (i ShippingAddressArgs) ToShippingAddressPtrOutput() ShippingAddressPtrOutput {
	return i.ToShippingAddressPtrOutputWithContext(context.Background())
}

func (i ShippingAddressArgs) ToShippingAddressPtrOutputWithContext(ctx context.Context) ShippingAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressOutput).ToShippingAddressPtrOutputWithContext(ctx)
}









type ShippingAddressPtrInput interface {
	pulumi.Input

	ToShippingAddressPtrOutput() ShippingAddressPtrOutput
	ToShippingAddressPtrOutputWithContext(context.Context) ShippingAddressPtrOutput
}

type shippingAddressPtrType ShippingAddressArgs

func ShippingAddressPtr(v *ShippingAddressArgs) ShippingAddressPtrInput {
	return (*shippingAddressPtrType)(v)
}

func (*shippingAddressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingAddress)(nil)).Elem()
}

func (i *shippingAddressPtrType) ToShippingAddressPtrOutput() ShippingAddressPtrOutput {
	return i.ToShippingAddressPtrOutputWithContext(context.Background())
}

func (i *shippingAddressPtrType) ToShippingAddressPtrOutputWithContext(ctx context.Context) ShippingAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressPtrOutput)
}

type ShippingAddressOutput struct{ *pulumi.OutputState }

func (ShippingAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingAddress)(nil)).Elem()
}

func (o ShippingAddressOutput) ToShippingAddressOutput() ShippingAddressOutput {
	return o
}

func (o ShippingAddressOutput) ToShippingAddressOutputWithContext(ctx context.Context) ShippingAddressOutput {
	return o
}

func (o ShippingAddressOutput) ToShippingAddressPtrOutput() ShippingAddressPtrOutput {
	return o.ToShippingAddressPtrOutputWithContext(context.Background())
}

func (o ShippingAddressOutput) ToShippingAddressPtrOutputWithContext(ctx context.Context) ShippingAddressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShippingAddress) *ShippingAddress {
		return &v
	}).(ShippingAddressPtrOutput)
}

func (o ShippingAddressOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.AddressType }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.CompanyName }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingAddress) string { return v.Country }).(pulumi.StringOutput)
}

func (o ShippingAddressOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.StateOrProvince }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) StreetAddress1() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingAddress) string { return v.StreetAddress1 }).(pulumi.StringOutput)
}

func (o ShippingAddressOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.StreetAddress2 }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) StreetAddress3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.StreetAddress3 }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) ZipExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.ZipExtendedCode }).(pulumi.StringPtrOutput)
}

type ShippingAddressPtrOutput struct{ *pulumi.OutputState }

func (ShippingAddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingAddress)(nil)).Elem()
}

func (o ShippingAddressPtrOutput) ToShippingAddressPtrOutput() ShippingAddressPtrOutput {
	return o
}

func (o ShippingAddressPtrOutput) ToShippingAddressPtrOutputWithContext(ctx context.Context) ShippingAddressPtrOutput {
	return o
}

func (o ShippingAddressPtrOutput) Elem() ShippingAddressOutput {
	return o.ApplyT(func(v *ShippingAddress) ShippingAddress {
		if v != nil {
			return *v
		}
		var ret ShippingAddress
		return ret
	}).(ShippingAddressOutput)
}

func (o ShippingAddressPtrOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.AddressType
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.CompanyName
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return &v.Country
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.StateOrProvince
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) StreetAddress1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return &v.StreetAddress1
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress2
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) StreetAddress3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress3
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) ZipExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.ZipExtendedCode
	}).(pulumi.StringPtrOutput)
}

type ShippingAddressResponse struct {
	AddressType     *string `pulumi:"addressType"`
	City            *string `pulumi:"city"`
	CompanyName     *string `pulumi:"companyName"`
	Country         string  `pulumi:"country"`
	PostalCode      *string `pulumi:"postalCode"`
	StateOrProvince *string `pulumi:"stateOrProvince"`
	StreetAddress1  string  `pulumi:"streetAddress1"`
	StreetAddress2  *string `pulumi:"streetAddress2"`
	StreetAddress3  *string `pulumi:"streetAddress3"`
	ZipExtendedCode *string `pulumi:"zipExtendedCode"`
}





type ShippingAddressResponseInput interface {
	pulumi.Input

	ToShippingAddressResponseOutput() ShippingAddressResponseOutput
	ToShippingAddressResponseOutputWithContext(context.Context) ShippingAddressResponseOutput
}

type ShippingAddressResponseArgs struct {
	AddressType     pulumi.StringPtrInput `pulumi:"addressType"`
	City            pulumi.StringPtrInput `pulumi:"city"`
	CompanyName     pulumi.StringPtrInput `pulumi:"companyName"`
	Country         pulumi.StringInput    `pulumi:"country"`
	PostalCode      pulumi.StringPtrInput `pulumi:"postalCode"`
	StateOrProvince pulumi.StringPtrInput `pulumi:"stateOrProvince"`
	StreetAddress1  pulumi.StringInput    `pulumi:"streetAddress1"`
	StreetAddress2  pulumi.StringPtrInput `pulumi:"streetAddress2"`
	StreetAddress3  pulumi.StringPtrInput `pulumi:"streetAddress3"`
	ZipExtendedCode pulumi.StringPtrInput `pulumi:"zipExtendedCode"`
}

func (ShippingAddressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingAddressResponse)(nil)).Elem()
}

func (i ShippingAddressResponseArgs) ToShippingAddressResponseOutput() ShippingAddressResponseOutput {
	return i.ToShippingAddressResponseOutputWithContext(context.Background())
}

func (i ShippingAddressResponseArgs) ToShippingAddressResponseOutputWithContext(ctx context.Context) ShippingAddressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressResponseOutput)
}

func (i ShippingAddressResponseArgs) ToShippingAddressResponsePtrOutput() ShippingAddressResponsePtrOutput {
	return i.ToShippingAddressResponsePtrOutputWithContext(context.Background())
}

func (i ShippingAddressResponseArgs) ToShippingAddressResponsePtrOutputWithContext(ctx context.Context) ShippingAddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressResponseOutput).ToShippingAddressResponsePtrOutputWithContext(ctx)
}









type ShippingAddressResponsePtrInput interface {
	pulumi.Input

	ToShippingAddressResponsePtrOutput() ShippingAddressResponsePtrOutput
	ToShippingAddressResponsePtrOutputWithContext(context.Context) ShippingAddressResponsePtrOutput
}

type shippingAddressResponsePtrType ShippingAddressResponseArgs

func ShippingAddressResponsePtr(v *ShippingAddressResponseArgs) ShippingAddressResponsePtrInput {
	return (*shippingAddressResponsePtrType)(v)
}

func (*shippingAddressResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingAddressResponse)(nil)).Elem()
}

func (i *shippingAddressResponsePtrType) ToShippingAddressResponsePtrOutput() ShippingAddressResponsePtrOutput {
	return i.ToShippingAddressResponsePtrOutputWithContext(context.Background())
}

func (i *shippingAddressResponsePtrType) ToShippingAddressResponsePtrOutputWithContext(ctx context.Context) ShippingAddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressResponsePtrOutput)
}

type ShippingAddressResponseOutput struct{ *pulumi.OutputState }

func (ShippingAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingAddressResponse)(nil)).Elem()
}

func (o ShippingAddressResponseOutput) ToShippingAddressResponseOutput() ShippingAddressResponseOutput {
	return o
}

func (o ShippingAddressResponseOutput) ToShippingAddressResponseOutputWithContext(ctx context.Context) ShippingAddressResponseOutput {
	return o
}

func (o ShippingAddressResponseOutput) ToShippingAddressResponsePtrOutput() ShippingAddressResponsePtrOutput {
	return o.ToShippingAddressResponsePtrOutputWithContext(context.Background())
}

func (o ShippingAddressResponseOutput) ToShippingAddressResponsePtrOutputWithContext(ctx context.Context) ShippingAddressResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShippingAddressResponse) *ShippingAddressResponse {
		return &v
	}).(ShippingAddressResponsePtrOutput)
}

func (o ShippingAddressResponseOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.AddressType }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.CompanyName }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingAddressResponse) string { return v.Country }).(pulumi.StringOutput)
}

func (o ShippingAddressResponseOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.StateOrProvince }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) StreetAddress1() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingAddressResponse) string { return v.StreetAddress1 }).(pulumi.StringOutput)
}

func (o ShippingAddressResponseOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.StreetAddress2 }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) StreetAddress3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.StreetAddress3 }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) ZipExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.ZipExtendedCode }).(pulumi.StringPtrOutput)
}

type ShippingAddressResponsePtrOutput struct{ *pulumi.OutputState }

func (ShippingAddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingAddressResponse)(nil)).Elem()
}

func (o ShippingAddressResponsePtrOutput) ToShippingAddressResponsePtrOutput() ShippingAddressResponsePtrOutput {
	return o
}

func (o ShippingAddressResponsePtrOutput) ToShippingAddressResponsePtrOutputWithContext(ctx context.Context) ShippingAddressResponsePtrOutput {
	return o
}

func (o ShippingAddressResponsePtrOutput) Elem() ShippingAddressResponseOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) ShippingAddressResponse {
		if v != nil {
			return *v
		}
		var ret ShippingAddressResponse
		return ret
	}).(ShippingAddressResponseOutput)
}

func (o ShippingAddressResponsePtrOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.AddressType
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.CompanyName
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Country
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.StateOrProvince
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) StreetAddress1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StreetAddress1
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress2
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) StreetAddress3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress3
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) ZipExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.ZipExtendedCode
	}).(pulumi.StringPtrOutput)
}

type ShippingDetailsResponse struct {
	CarrierDisplayName string `pulumi:"carrierDisplayName"`
	CarrierName        string `pulumi:"carrierName"`
	TrackingId         string `pulumi:"trackingId"`
	TrackingUrl        string `pulumi:"trackingUrl"`
}





type ShippingDetailsResponseInput interface {
	pulumi.Input

	ToShippingDetailsResponseOutput() ShippingDetailsResponseOutput
	ToShippingDetailsResponseOutputWithContext(context.Context) ShippingDetailsResponseOutput
}

type ShippingDetailsResponseArgs struct {
	CarrierDisplayName pulumi.StringInput `pulumi:"carrierDisplayName"`
	CarrierName        pulumi.StringInput `pulumi:"carrierName"`
	TrackingId         pulumi.StringInput `pulumi:"trackingId"`
	TrackingUrl        pulumi.StringInput `pulumi:"trackingUrl"`
}

func (ShippingDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingDetailsResponse)(nil)).Elem()
}

func (i ShippingDetailsResponseArgs) ToShippingDetailsResponseOutput() ShippingDetailsResponseOutput {
	return i.ToShippingDetailsResponseOutputWithContext(context.Background())
}

func (i ShippingDetailsResponseArgs) ToShippingDetailsResponseOutputWithContext(ctx context.Context) ShippingDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingDetailsResponseOutput)
}

func (i ShippingDetailsResponseArgs) ToShippingDetailsResponsePtrOutput() ShippingDetailsResponsePtrOutput {
	return i.ToShippingDetailsResponsePtrOutputWithContext(context.Background())
}

func (i ShippingDetailsResponseArgs) ToShippingDetailsResponsePtrOutputWithContext(ctx context.Context) ShippingDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingDetailsResponseOutput).ToShippingDetailsResponsePtrOutputWithContext(ctx)
}









type ShippingDetailsResponsePtrInput interface {
	pulumi.Input

	ToShippingDetailsResponsePtrOutput() ShippingDetailsResponsePtrOutput
	ToShippingDetailsResponsePtrOutputWithContext(context.Context) ShippingDetailsResponsePtrOutput
}

type shippingDetailsResponsePtrType ShippingDetailsResponseArgs

func ShippingDetailsResponsePtr(v *ShippingDetailsResponseArgs) ShippingDetailsResponsePtrInput {
	return (*shippingDetailsResponsePtrType)(v)
}

func (*shippingDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingDetailsResponse)(nil)).Elem()
}

func (i *shippingDetailsResponsePtrType) ToShippingDetailsResponsePtrOutput() ShippingDetailsResponsePtrOutput {
	return i.ToShippingDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *shippingDetailsResponsePtrType) ToShippingDetailsResponsePtrOutputWithContext(ctx context.Context) ShippingDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingDetailsResponsePtrOutput)
}

type ShippingDetailsResponseOutput struct{ *pulumi.OutputState }

func (ShippingDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingDetailsResponse)(nil)).Elem()
}

func (o ShippingDetailsResponseOutput) ToShippingDetailsResponseOutput() ShippingDetailsResponseOutput {
	return o
}

func (o ShippingDetailsResponseOutput) ToShippingDetailsResponseOutputWithContext(ctx context.Context) ShippingDetailsResponseOutput {
	return o
}

func (o ShippingDetailsResponseOutput) ToShippingDetailsResponsePtrOutput() ShippingDetailsResponsePtrOutput {
	return o.ToShippingDetailsResponsePtrOutputWithContext(context.Background())
}

func (o ShippingDetailsResponseOutput) ToShippingDetailsResponsePtrOutputWithContext(ctx context.Context) ShippingDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShippingDetailsResponse) *ShippingDetailsResponse {
		return &v
	}).(ShippingDetailsResponsePtrOutput)
}

func (o ShippingDetailsResponseOutput) CarrierDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingDetailsResponse) string { return v.CarrierDisplayName }).(pulumi.StringOutput)
}

func (o ShippingDetailsResponseOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingDetailsResponse) string { return v.CarrierName }).(pulumi.StringOutput)
}

func (o ShippingDetailsResponseOutput) TrackingId() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingDetailsResponse) string { return v.TrackingId }).(pulumi.StringOutput)
}

func (o ShippingDetailsResponseOutput) TrackingUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingDetailsResponse) string { return v.TrackingUrl }).(pulumi.StringOutput)
}

type ShippingDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ShippingDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingDetailsResponse)(nil)).Elem()
}

func (o ShippingDetailsResponsePtrOutput) ToShippingDetailsResponsePtrOutput() ShippingDetailsResponsePtrOutput {
	return o
}

func (o ShippingDetailsResponsePtrOutput) ToShippingDetailsResponsePtrOutputWithContext(ctx context.Context) ShippingDetailsResponsePtrOutput {
	return o
}

func (o ShippingDetailsResponsePtrOutput) Elem() ShippingDetailsResponseOutput {
	return o.ApplyT(func(v *ShippingDetailsResponse) ShippingDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ShippingDetailsResponse
		return ret
	}).(ShippingDetailsResponseOutput)
}

func (o ShippingDetailsResponsePtrOutput) CarrierDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CarrierDisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ShippingDetailsResponsePtrOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CarrierName
	}).(pulumi.StringPtrOutput)
}

func (o ShippingDetailsResponsePtrOutput) TrackingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TrackingId
	}).(pulumi.StringPtrOutput)
}

func (o ShippingDetailsResponsePtrOutput) TrackingUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TrackingUrl
	}).(pulumi.StringPtrOutput)
}

type SpecificationResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type SpecificationResponseInput interface {
	pulumi.Input

	ToSpecificationResponseOutput() SpecificationResponseOutput
	ToSpecificationResponseOutputWithContext(context.Context) SpecificationResponseOutput
}

type SpecificationResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (SpecificationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SpecificationResponse)(nil)).Elem()
}

func (i SpecificationResponseArgs) ToSpecificationResponseOutput() SpecificationResponseOutput {
	return i.ToSpecificationResponseOutputWithContext(context.Background())
}

func (i SpecificationResponseArgs) ToSpecificationResponseOutputWithContext(ctx context.Context) SpecificationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpecificationResponseOutput)
}





type SpecificationResponseArrayInput interface {
	pulumi.Input

	ToSpecificationResponseArrayOutput() SpecificationResponseArrayOutput
	ToSpecificationResponseArrayOutputWithContext(context.Context) SpecificationResponseArrayOutput
}

type SpecificationResponseArray []SpecificationResponseInput

func (SpecificationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpecificationResponse)(nil)).Elem()
}

func (i SpecificationResponseArray) ToSpecificationResponseArrayOutput() SpecificationResponseArrayOutput {
	return i.ToSpecificationResponseArrayOutputWithContext(context.Background())
}

func (i SpecificationResponseArray) ToSpecificationResponseArrayOutputWithContext(ctx context.Context) SpecificationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpecificationResponseArrayOutput)
}

type SpecificationResponseOutput struct{ *pulumi.OutputState }

func (SpecificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpecificationResponse)(nil)).Elem()
}

func (o SpecificationResponseOutput) ToSpecificationResponseOutput() SpecificationResponseOutput {
	return o
}

func (o SpecificationResponseOutput) ToSpecificationResponseOutputWithContext(ctx context.Context) SpecificationResponseOutput {
	return o
}

func (o SpecificationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SpecificationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SpecificationResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SpecificationResponse) string { return v.Value }).(pulumi.StringOutput)
}

type SpecificationResponseArrayOutput struct{ *pulumi.OutputState }

func (SpecificationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpecificationResponse)(nil)).Elem()
}

func (o SpecificationResponseArrayOutput) ToSpecificationResponseArrayOutput() SpecificationResponseArrayOutput {
	return o
}

func (o SpecificationResponseArrayOutput) ToSpecificationResponseArrayOutputWithContext(ctx context.Context) SpecificationResponseArrayOutput {
	return o
}

func (o SpecificationResponseArrayOutput) Index(i pulumi.IntInput) SpecificationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpecificationResponse {
		return vs[0].([]SpecificationResponse)[vs[1].(int)]
	}).(SpecificationResponseOutput)
}

type StageDetailsResponse struct {
	DisplayName string `pulumi:"displayName"`
	StageName   string `pulumi:"stageName"`
	StageStatus string `pulumi:"stageStatus"`
	StartTime   string `pulumi:"startTime"`
}





type StageDetailsResponseInput interface {
	pulumi.Input

	ToStageDetailsResponseOutput() StageDetailsResponseOutput
	ToStageDetailsResponseOutputWithContext(context.Context) StageDetailsResponseOutput
}

type StageDetailsResponseArgs struct {
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	StageName   pulumi.StringInput `pulumi:"stageName"`
	StageStatus pulumi.StringInput `pulumi:"stageStatus"`
	StartTime   pulumi.StringInput `pulumi:"startTime"`
}

func (StageDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StageDetailsResponse)(nil)).Elem()
}

func (i StageDetailsResponseArgs) ToStageDetailsResponseOutput() StageDetailsResponseOutput {
	return i.ToStageDetailsResponseOutputWithContext(context.Background())
}

func (i StageDetailsResponseArgs) ToStageDetailsResponseOutputWithContext(ctx context.Context) StageDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageDetailsResponseOutput)
}

func (i StageDetailsResponseArgs) ToStageDetailsResponsePtrOutput() StageDetailsResponsePtrOutput {
	return i.ToStageDetailsResponsePtrOutputWithContext(context.Background())
}

func (i StageDetailsResponseArgs) ToStageDetailsResponsePtrOutputWithContext(ctx context.Context) StageDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageDetailsResponseOutput).ToStageDetailsResponsePtrOutputWithContext(ctx)
}









type StageDetailsResponsePtrInput interface {
	pulumi.Input

	ToStageDetailsResponsePtrOutput() StageDetailsResponsePtrOutput
	ToStageDetailsResponsePtrOutputWithContext(context.Context) StageDetailsResponsePtrOutput
}

type stageDetailsResponsePtrType StageDetailsResponseArgs

func StageDetailsResponsePtr(v *StageDetailsResponseArgs) StageDetailsResponsePtrInput {
	return (*stageDetailsResponsePtrType)(v)
}

func (*stageDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StageDetailsResponse)(nil)).Elem()
}

func (i *stageDetailsResponsePtrType) ToStageDetailsResponsePtrOutput() StageDetailsResponsePtrOutput {
	return i.ToStageDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *stageDetailsResponsePtrType) ToStageDetailsResponsePtrOutputWithContext(ctx context.Context) StageDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageDetailsResponsePtrOutput)
}





type StageDetailsResponseArrayInput interface {
	pulumi.Input

	ToStageDetailsResponseArrayOutput() StageDetailsResponseArrayOutput
	ToStageDetailsResponseArrayOutputWithContext(context.Context) StageDetailsResponseArrayOutput
}

type StageDetailsResponseArray []StageDetailsResponseInput

func (StageDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StageDetailsResponse)(nil)).Elem()
}

func (i StageDetailsResponseArray) ToStageDetailsResponseArrayOutput() StageDetailsResponseArrayOutput {
	return i.ToStageDetailsResponseArrayOutputWithContext(context.Background())
}

func (i StageDetailsResponseArray) ToStageDetailsResponseArrayOutputWithContext(ctx context.Context) StageDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageDetailsResponseArrayOutput)
}

type StageDetailsResponseOutput struct{ *pulumi.OutputState }

func (StageDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StageDetailsResponse)(nil)).Elem()
}

func (o StageDetailsResponseOutput) ToStageDetailsResponseOutput() StageDetailsResponseOutput {
	return o
}

func (o StageDetailsResponseOutput) ToStageDetailsResponseOutputWithContext(ctx context.Context) StageDetailsResponseOutput {
	return o
}

func (o StageDetailsResponseOutput) ToStageDetailsResponsePtrOutput() StageDetailsResponsePtrOutput {
	return o.ToStageDetailsResponsePtrOutputWithContext(context.Background())
}

func (o StageDetailsResponseOutput) ToStageDetailsResponsePtrOutputWithContext(ctx context.Context) StageDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StageDetailsResponse) *StageDetailsResponse {
		return &v
	}).(StageDetailsResponsePtrOutput)
}

func (o StageDetailsResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v StageDetailsResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o StageDetailsResponseOutput) StageName() pulumi.StringOutput {
	return o.ApplyT(func(v StageDetailsResponse) string { return v.StageName }).(pulumi.StringOutput)
}

func (o StageDetailsResponseOutput) StageStatus() pulumi.StringOutput {
	return o.ApplyT(func(v StageDetailsResponse) string { return v.StageStatus }).(pulumi.StringOutput)
}

func (o StageDetailsResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v StageDetailsResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

type StageDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (StageDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageDetailsResponse)(nil)).Elem()
}

func (o StageDetailsResponsePtrOutput) ToStageDetailsResponsePtrOutput() StageDetailsResponsePtrOutput {
	return o
}

func (o StageDetailsResponsePtrOutput) ToStageDetailsResponsePtrOutputWithContext(ctx context.Context) StageDetailsResponsePtrOutput {
	return o
}

func (o StageDetailsResponsePtrOutput) Elem() StageDetailsResponseOutput {
	return o.ApplyT(func(v *StageDetailsResponse) StageDetailsResponse {
		if v != nil {
			return *v
		}
		var ret StageDetailsResponse
		return ret
	}).(StageDetailsResponseOutput)
}

func (o StageDetailsResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o StageDetailsResponsePtrOutput) StageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StageName
	}).(pulumi.StringPtrOutput)
}

func (o StageDetailsResponsePtrOutput) StageStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StageStatus
	}).(pulumi.StringPtrOutput)
}

func (o StageDetailsResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartTime
	}).(pulumi.StringPtrOutput)
}

type StageDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (StageDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StageDetailsResponse)(nil)).Elem()
}

func (o StageDetailsResponseArrayOutput) ToStageDetailsResponseArrayOutput() StageDetailsResponseArrayOutput {
	return o
}

func (o StageDetailsResponseArrayOutput) ToStageDetailsResponseArrayOutputWithContext(ctx context.Context) StageDetailsResponseArrayOutput {
	return o
}

func (o StageDetailsResponseArrayOutput) Index(i pulumi.IntInput) StageDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StageDetailsResponse {
		return vs[0].([]StageDetailsResponse)[vs[1].(int)]
	}).(StageDetailsResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type TransportPreferences struct {
	PreferredShipmentType string `pulumi:"preferredShipmentType"`
}





type TransportPreferencesInput interface {
	pulumi.Input

	ToTransportPreferencesOutput() TransportPreferencesOutput
	ToTransportPreferencesOutputWithContext(context.Context) TransportPreferencesOutput
}

type TransportPreferencesArgs struct {
	PreferredShipmentType pulumi.StringInput `pulumi:"preferredShipmentType"`
}

func (TransportPreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportPreferences)(nil)).Elem()
}

func (i TransportPreferencesArgs) ToTransportPreferencesOutput() TransportPreferencesOutput {
	return i.ToTransportPreferencesOutputWithContext(context.Background())
}

func (i TransportPreferencesArgs) ToTransportPreferencesOutputWithContext(ctx context.Context) TransportPreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesOutput)
}

func (i TransportPreferencesArgs) ToTransportPreferencesPtrOutput() TransportPreferencesPtrOutput {
	return i.ToTransportPreferencesPtrOutputWithContext(context.Background())
}

func (i TransportPreferencesArgs) ToTransportPreferencesPtrOutputWithContext(ctx context.Context) TransportPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesOutput).ToTransportPreferencesPtrOutputWithContext(ctx)
}









type TransportPreferencesPtrInput interface {
	pulumi.Input

	ToTransportPreferencesPtrOutput() TransportPreferencesPtrOutput
	ToTransportPreferencesPtrOutputWithContext(context.Context) TransportPreferencesPtrOutput
}

type transportPreferencesPtrType TransportPreferencesArgs

func TransportPreferencesPtr(v *TransportPreferencesArgs) TransportPreferencesPtrInput {
	return (*transportPreferencesPtrType)(v)
}

func (*transportPreferencesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransportPreferences)(nil)).Elem()
}

func (i *transportPreferencesPtrType) ToTransportPreferencesPtrOutput() TransportPreferencesPtrOutput {
	return i.ToTransportPreferencesPtrOutputWithContext(context.Background())
}

func (i *transportPreferencesPtrType) ToTransportPreferencesPtrOutputWithContext(ctx context.Context) TransportPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesPtrOutput)
}

type TransportPreferencesOutput struct{ *pulumi.OutputState }

func (TransportPreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportPreferences)(nil)).Elem()
}

func (o TransportPreferencesOutput) ToTransportPreferencesOutput() TransportPreferencesOutput {
	return o
}

func (o TransportPreferencesOutput) ToTransportPreferencesOutputWithContext(ctx context.Context) TransportPreferencesOutput {
	return o
}

func (o TransportPreferencesOutput) ToTransportPreferencesPtrOutput() TransportPreferencesPtrOutput {
	return o.ToTransportPreferencesPtrOutputWithContext(context.Background())
}

func (o TransportPreferencesOutput) ToTransportPreferencesPtrOutputWithContext(ctx context.Context) TransportPreferencesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransportPreferences) *TransportPreferences {
		return &v
	}).(TransportPreferencesPtrOutput)
}

func (o TransportPreferencesOutput) PreferredShipmentType() pulumi.StringOutput {
	return o.ApplyT(func(v TransportPreferences) string { return v.PreferredShipmentType }).(pulumi.StringOutput)
}

type TransportPreferencesPtrOutput struct{ *pulumi.OutputState }

func (TransportPreferencesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransportPreferences)(nil)).Elem()
}

func (o TransportPreferencesPtrOutput) ToTransportPreferencesPtrOutput() TransportPreferencesPtrOutput {
	return o
}

func (o TransportPreferencesPtrOutput) ToTransportPreferencesPtrOutputWithContext(ctx context.Context) TransportPreferencesPtrOutput {
	return o
}

func (o TransportPreferencesPtrOutput) Elem() TransportPreferencesOutput {
	return o.ApplyT(func(v *TransportPreferences) TransportPreferences {
		if v != nil {
			return *v
		}
		var ret TransportPreferences
		return ret
	}).(TransportPreferencesOutput)
}

func (o TransportPreferencesPtrOutput) PreferredShipmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransportPreferences) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredShipmentType
	}).(pulumi.StringPtrOutput)
}

type TransportPreferencesResponse struct {
	PreferredShipmentType string `pulumi:"preferredShipmentType"`
}





type TransportPreferencesResponseInput interface {
	pulumi.Input

	ToTransportPreferencesResponseOutput() TransportPreferencesResponseOutput
	ToTransportPreferencesResponseOutputWithContext(context.Context) TransportPreferencesResponseOutput
}

type TransportPreferencesResponseArgs struct {
	PreferredShipmentType pulumi.StringInput `pulumi:"preferredShipmentType"`
}

func (TransportPreferencesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportPreferencesResponse)(nil)).Elem()
}

func (i TransportPreferencesResponseArgs) ToTransportPreferencesResponseOutput() TransportPreferencesResponseOutput {
	return i.ToTransportPreferencesResponseOutputWithContext(context.Background())
}

func (i TransportPreferencesResponseArgs) ToTransportPreferencesResponseOutputWithContext(ctx context.Context) TransportPreferencesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesResponseOutput)
}

func (i TransportPreferencesResponseArgs) ToTransportPreferencesResponsePtrOutput() TransportPreferencesResponsePtrOutput {
	return i.ToTransportPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i TransportPreferencesResponseArgs) ToTransportPreferencesResponsePtrOutputWithContext(ctx context.Context) TransportPreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesResponseOutput).ToTransportPreferencesResponsePtrOutputWithContext(ctx)
}









type TransportPreferencesResponsePtrInput interface {
	pulumi.Input

	ToTransportPreferencesResponsePtrOutput() TransportPreferencesResponsePtrOutput
	ToTransportPreferencesResponsePtrOutputWithContext(context.Context) TransportPreferencesResponsePtrOutput
}

type transportPreferencesResponsePtrType TransportPreferencesResponseArgs

func TransportPreferencesResponsePtr(v *TransportPreferencesResponseArgs) TransportPreferencesResponsePtrInput {
	return (*transportPreferencesResponsePtrType)(v)
}

func (*transportPreferencesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransportPreferencesResponse)(nil)).Elem()
}

func (i *transportPreferencesResponsePtrType) ToTransportPreferencesResponsePtrOutput() TransportPreferencesResponsePtrOutput {
	return i.ToTransportPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i *transportPreferencesResponsePtrType) ToTransportPreferencesResponsePtrOutputWithContext(ctx context.Context) TransportPreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesResponsePtrOutput)
}

type TransportPreferencesResponseOutput struct{ *pulumi.OutputState }

func (TransportPreferencesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportPreferencesResponse)(nil)).Elem()
}

func (o TransportPreferencesResponseOutput) ToTransportPreferencesResponseOutput() TransportPreferencesResponseOutput {
	return o
}

func (o TransportPreferencesResponseOutput) ToTransportPreferencesResponseOutputWithContext(ctx context.Context) TransportPreferencesResponseOutput {
	return o
}

func (o TransportPreferencesResponseOutput) ToTransportPreferencesResponsePtrOutput() TransportPreferencesResponsePtrOutput {
	return o.ToTransportPreferencesResponsePtrOutputWithContext(context.Background())
}

func (o TransportPreferencesResponseOutput) ToTransportPreferencesResponsePtrOutputWithContext(ctx context.Context) TransportPreferencesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransportPreferencesResponse) *TransportPreferencesResponse {
		return &v
	}).(TransportPreferencesResponsePtrOutput)
}

func (o TransportPreferencesResponseOutput) PreferredShipmentType() pulumi.StringOutput {
	return o.ApplyT(func(v TransportPreferencesResponse) string { return v.PreferredShipmentType }).(pulumi.StringOutput)
}

type TransportPreferencesResponsePtrOutput struct{ *pulumi.OutputState }

func (TransportPreferencesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransportPreferencesResponse)(nil)).Elem()
}

func (o TransportPreferencesResponsePtrOutput) ToTransportPreferencesResponsePtrOutput() TransportPreferencesResponsePtrOutput {
	return o
}

func (o TransportPreferencesResponsePtrOutput) ToTransportPreferencesResponsePtrOutputWithContext(ctx context.Context) TransportPreferencesResponsePtrOutput {
	return o
}

func (o TransportPreferencesResponsePtrOutput) Elem() TransportPreferencesResponseOutput {
	return o.ApplyT(func(v *TransportPreferencesResponse) TransportPreferencesResponse {
		if v != nil {
			return *v
		}
		var ret TransportPreferencesResponse
		return ret
	}).(TransportPreferencesResponseOutput)
}

func (o TransportPreferencesResponsePtrOutput) PreferredShipmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransportPreferencesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredShipmentType
	}).(pulumi.StringPtrOutput)
}

type FilterablePropertyArrayMap map[string]FilterablePropertyArrayInput

func (FilterablePropertyArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterablePropertyArray)(nil)).Elem()
}

func (i FilterablePropertyArrayMap) ToFilterablePropertyArrayMapOutput() FilterablePropertyArrayMapOutput {
	return i.ToFilterablePropertyArrayMapOutputWithContext(context.Background())
}

func (i FilterablePropertyArrayMap) ToFilterablePropertyArrayMapOutputWithContext(ctx context.Context) FilterablePropertyArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterablePropertyArrayMapOutput)
}

type FilterablePropertyArrayMapOutput struct{ *pulumi.OutputState }

func (FilterablePropertyArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FilterablePropertyArray)(nil)).Elem()
}

func (o FilterablePropertyArrayMapOutput) ToFilterablePropertyArrayMapOutput() FilterablePropertyArrayMapOutput {
	return o
}

func (o FilterablePropertyArrayMapOutput) ToFilterablePropertyArrayMapOutputWithContext(ctx context.Context) FilterablePropertyArrayMapOutput {
	return o
}

func (o FilterablePropertyArrayMapOutput) MapIndex(k pulumi.StringInput) FilterablePropertyArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FilterablePropertyArray {
		return vs[0].(map[string]FilterablePropertyArray)[vs[1].(string)]
	}).(FilterablePropertyArrayOutput)
}





type FilterablePropertyArrayMapInput interface {
	pulumi.Input

	ToFilterablePropertyArrayMapOutput() FilterablePropertyArrayMapOutput
	ToFilterablePropertyArrayMapOutputWithContext(context.Context) FilterablePropertyArrayMapOutput
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressDetailsInput)(nil)).Elem(), AddressDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressDetailsPtrInput)(nil)).Elem(), AddressDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressDetailsResponseInput)(nil)).Elem(), AddressDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressDetailsResponsePtrInput)(nil)).Elem(), AddressDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPropertiesInput)(nil)).Elem(), AddressPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPropertiesPtrInput)(nil)).Elem(), AddressPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPropertiesResponseInput)(nil)).Elem(), AddressPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPropertiesResponsePtrInput)(nil)).Elem(), AddressPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AvailabilityInformationResponseInput)(nil)).Elem(), AvailabilityInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BillingMeterDetailsResponseInput)(nil)).Elem(), BillingMeterDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BillingMeterDetailsResponseArrayInput)(nil)).Elem(), BillingMeterDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationFiltersInput)(nil)).Elem(), ConfigurationFiltersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationFiltersArrayInput)(nil)).Elem(), ConfigurationFiltersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationResponseInput)(nil)).Elem(), ConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationResponseArrayInput)(nil)).Elem(), ConfigurationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactDetailsInput)(nil)).Elem(), ContactDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactDetailsPtrInput)(nil)).Elem(), ContactDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactDetailsResponseInput)(nil)).Elem(), ContactDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactDetailsResponsePtrInput)(nil)).Elem(), ContactDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CostInformationResponseInput)(nil)).Elem(), CostInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerSubscriptionDetailsInput)(nil)).Elem(), CustomerSubscriptionDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerSubscriptionDetailsPtrInput)(nil)).Elem(), CustomerSubscriptionDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerSubscriptionRegisteredFeaturesInput)(nil)).Elem(), CustomerSubscriptionRegisteredFeaturesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerSubscriptionRegisteredFeaturesArrayInput)(nil)).Elem(), CustomerSubscriptionRegisteredFeaturesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DescriptionResponseInput)(nil)).Elem(), DescriptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceDetailsResponseInput)(nil)).Elem(), DeviceDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceDetailsResponseArrayInput)(nil)).Elem(), DeviceDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DimensionsResponseInput)(nil)).Elem(), DimensionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DisplayInfoResponseInput)(nil)).Elem(), DisplayInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DisplayInfoResponsePtrInput)(nil)).Elem(), DisplayInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPreferencesInput)(nil)).Elem(), EncryptionPreferencesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPreferencesPtrInput)(nil)).Elem(), EncryptionPreferencesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPreferencesResponseInput)(nil)).Elem(), EncryptionPreferencesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPreferencesResponsePtrInput)(nil)).Elem(), EncryptionPreferencesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorAdditionalInfoResponseInput)(nil)).Elem(), ErrorAdditionalInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorAdditionalInfoResponseArrayInput)(nil)).Elem(), ErrorAdditionalInfoResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorDetailResponseInput)(nil)).Elem(), ErrorDetailResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorDetailResponsePtrInput)(nil)).Elem(), ErrorDetailResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorDetailResponseArrayInput)(nil)).Elem(), ErrorDetailResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterablePropertyInput)(nil)).Elem(), FilterablePropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterablePropertyArrayInput)(nil)).Elem(), FilterablePropertyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterablePropertyResponseInput)(nil)).Elem(), FilterablePropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterablePropertyResponseArrayInput)(nil)).Elem(), FilterablePropertyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HierarchyInformationInput)(nil)).Elem(), HierarchyInformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HierarchyInformationPtrInput)(nil)).Elem(), HierarchyInformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HierarchyInformationResponseInput)(nil)).Elem(), HierarchyInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HierarchyInformationResponsePtrInput)(nil)).Elem(), HierarchyInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInformationResponseInput)(nil)).Elem(), ImageInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInformationResponseArrayInput)(nil)).Elem(), ImageInformationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkResponseInput)(nil)).Elem(), LinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkResponseArrayInput)(nil)).Elem(), LinkResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementResourcePreferencesInput)(nil)).Elem(), ManagementResourcePreferencesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementResourcePreferencesPtrInput)(nil)).Elem(), ManagementResourcePreferencesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementResourcePreferencesResponseInput)(nil)).Elem(), ManagementResourcePreferencesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementResourcePreferencesResponsePtrInput)(nil)).Elem(), ManagementResourcePreferencesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationPreferenceInput)(nil)).Elem(), NotificationPreferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationPreferenceArrayInput)(nil)).Elem(), NotificationPreferenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationPreferenceResponseInput)(nil)).Elem(), NotificationPreferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationPreferenceResponseArrayInput)(nil)).Elem(), NotificationPreferenceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderItemDetailsInput)(nil)).Elem(), OrderItemDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderItemDetailsPtrInput)(nil)).Elem(), OrderItemDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderItemDetailsResponseInput)(nil)).Elem(), OrderItemDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderItemDetailsResponsePtrInput)(nil)).Elem(), OrderItemDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Pav2MeterDetailsResponseInput)(nil)).Elem(), Pav2MeterDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PreferencesInput)(nil)).Elem(), PreferencesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PreferencesPtrInput)(nil)).Elem(), PreferencesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PreferencesResponseInput)(nil)).Elem(), PreferencesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PreferencesResponsePtrInput)(nil)).Elem(), PreferencesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductDetailsInput)(nil)).Elem(), ProductDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductDetailsPtrInput)(nil)).Elem(), ProductDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductDetailsResponseInput)(nil)).Elem(), ProductDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductDetailsResponsePtrInput)(nil)).Elem(), ProductDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductFamilyResponseInput)(nil)).Elem(), ProductFamilyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductFamilyResponseArrayInput)(nil)).Elem(), ProductFamilyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductLineResponseInput)(nil)).Elem(), ProductLineResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductLineResponseArrayInput)(nil)).Elem(), ProductLineResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductResponseInput)(nil)).Elem(), ProductResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductResponseArrayInput)(nil)).Elem(), ProductResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PurchaseMeterDetailsResponseInput)(nil)).Elem(), PurchaseMeterDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShippingAddressInput)(nil)).Elem(), ShippingAddressArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShippingAddressPtrInput)(nil)).Elem(), ShippingAddressArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShippingAddressResponseInput)(nil)).Elem(), ShippingAddressResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShippingAddressResponsePtrInput)(nil)).Elem(), ShippingAddressResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShippingDetailsResponseInput)(nil)).Elem(), ShippingDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShippingDetailsResponsePtrInput)(nil)).Elem(), ShippingDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpecificationResponseInput)(nil)).Elem(), SpecificationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpecificationResponseArrayInput)(nil)).Elem(), SpecificationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageDetailsResponseInput)(nil)).Elem(), StageDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageDetailsResponsePtrInput)(nil)).Elem(), StageDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageDetailsResponseArrayInput)(nil)).Elem(), StageDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransportPreferencesInput)(nil)).Elem(), TransportPreferencesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransportPreferencesPtrInput)(nil)).Elem(), TransportPreferencesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransportPreferencesResponseInput)(nil)).Elem(), TransportPreferencesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransportPreferencesResponsePtrInput)(nil)).Elem(), TransportPreferencesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterablePropertyArrayMapInput)(nil)).Elem(), FilterablePropertyArrayMap{})
	pulumi.RegisterOutputType(AddressDetailsOutput{})
	pulumi.RegisterOutputType(AddressDetailsPtrOutput{})
	pulumi.RegisterOutputType(AddressDetailsResponseOutput{})
	pulumi.RegisterOutputType(AddressDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(AddressPropertiesOutput{})
	pulumi.RegisterOutputType(AddressPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AddressPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AddressPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AvailabilityInformationResponseOutput{})
	pulumi.RegisterOutputType(BillingMeterDetailsResponseOutput{})
	pulumi.RegisterOutputType(BillingMeterDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationFiltersOutput{})
	pulumi.RegisterOutputType(ConfigurationFiltersArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(ContactDetailsOutput{})
	pulumi.RegisterOutputType(ContactDetailsPtrOutput{})
	pulumi.RegisterOutputType(ContactDetailsResponseOutput{})
	pulumi.RegisterOutputType(ContactDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(CostInformationResponseOutput{})
	pulumi.RegisterOutputType(CustomerSubscriptionDetailsOutput{})
	pulumi.RegisterOutputType(CustomerSubscriptionDetailsPtrOutput{})
	pulumi.RegisterOutputType(CustomerSubscriptionRegisteredFeaturesOutput{})
	pulumi.RegisterOutputType(CustomerSubscriptionRegisteredFeaturesArrayOutput{})
	pulumi.RegisterOutputType(DescriptionResponseOutput{})
	pulumi.RegisterOutputType(DeviceDetailsResponseOutput{})
	pulumi.RegisterOutputType(DeviceDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DimensionsResponseOutput{})
	pulumi.RegisterOutputType(DisplayInfoResponseOutput{})
	pulumi.RegisterOutputType(DisplayInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionPreferencesOutput{})
	pulumi.RegisterOutputType(EncryptionPreferencesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPreferencesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(FilterablePropertyOutput{})
	pulumi.RegisterOutputType(FilterablePropertyArrayOutput{})
	pulumi.RegisterOutputType(FilterablePropertyResponseOutput{})
	pulumi.RegisterOutputType(FilterablePropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(HierarchyInformationOutput{})
	pulumi.RegisterOutputType(HierarchyInformationPtrOutput{})
	pulumi.RegisterOutputType(HierarchyInformationResponseOutput{})
	pulumi.RegisterOutputType(HierarchyInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageInformationResponseOutput{})
	pulumi.RegisterOutputType(ImageInformationResponseArrayOutput{})
	pulumi.RegisterOutputType(LinkResponseOutput{})
	pulumi.RegisterOutputType(LinkResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementResourcePreferencesOutput{})
	pulumi.RegisterOutputType(ManagementResourcePreferencesPtrOutput{})
	pulumi.RegisterOutputType(ManagementResourcePreferencesResponseOutput{})
	pulumi.RegisterOutputType(ManagementResourcePreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceArrayOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceResponseOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(OrderItemDetailsOutput{})
	pulumi.RegisterOutputType(OrderItemDetailsPtrOutput{})
	pulumi.RegisterOutputType(OrderItemDetailsResponseOutput{})
	pulumi.RegisterOutputType(OrderItemDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(Pav2MeterDetailsResponseOutput{})
	pulumi.RegisterOutputType(PreferencesOutput{})
	pulumi.RegisterOutputType(PreferencesPtrOutput{})
	pulumi.RegisterOutputType(PreferencesResponseOutput{})
	pulumi.RegisterOutputType(PreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(ProductDetailsOutput{})
	pulumi.RegisterOutputType(ProductDetailsPtrOutput{})
	pulumi.RegisterOutputType(ProductDetailsResponseOutput{})
	pulumi.RegisterOutputType(ProductDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ProductFamilyResponseOutput{})
	pulumi.RegisterOutputType(ProductFamilyResponseArrayOutput{})
	pulumi.RegisterOutputType(ProductLineResponseOutput{})
	pulumi.RegisterOutputType(ProductLineResponseArrayOutput{})
	pulumi.RegisterOutputType(ProductResponseOutput{})
	pulumi.RegisterOutputType(ProductResponseArrayOutput{})
	pulumi.RegisterOutputType(PurchaseMeterDetailsResponseOutput{})
	pulumi.RegisterOutputType(ShippingAddressOutput{})
	pulumi.RegisterOutputType(ShippingAddressPtrOutput{})
	pulumi.RegisterOutputType(ShippingAddressResponseOutput{})
	pulumi.RegisterOutputType(ShippingAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(ShippingDetailsResponseOutput{})
	pulumi.RegisterOutputType(ShippingDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(SpecificationResponseOutput{})
	pulumi.RegisterOutputType(SpecificationResponseArrayOutput{})
	pulumi.RegisterOutputType(StageDetailsResponseOutput{})
	pulumi.RegisterOutputType(StageDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(StageDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TransportPreferencesOutput{})
	pulumi.RegisterOutputType(TransportPreferencesPtrOutput{})
	pulumi.RegisterOutputType(TransportPreferencesResponseOutput{})
	pulumi.RegisterOutputType(TransportPreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(FilterablePropertyArrayMapOutput{})
}
