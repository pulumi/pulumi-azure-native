


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdditionalConfiguration struct {
	HierarchyInformation HierarchyInformation `pulumi:"hierarchyInformation"`
	Quantity             int                  `pulumi:"quantity"`
}





type AdditionalConfigurationInput interface {
	pulumi.Input

	ToAdditionalConfigurationOutput() AdditionalConfigurationOutput
	ToAdditionalConfigurationOutputWithContext(context.Context) AdditionalConfigurationOutput
}

type AdditionalConfigurationArgs struct {
	HierarchyInformation HierarchyInformationInput `pulumi:"hierarchyInformation"`
	Quantity             pulumi.IntInput           `pulumi:"quantity"`
}

func (AdditionalConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalConfiguration)(nil)).Elem()
}

func (i AdditionalConfigurationArgs) ToAdditionalConfigurationOutput() AdditionalConfigurationOutput {
	return i.ToAdditionalConfigurationOutputWithContext(context.Background())
}

func (i AdditionalConfigurationArgs) ToAdditionalConfigurationOutputWithContext(ctx context.Context) AdditionalConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalConfigurationOutput)
}





type AdditionalConfigurationArrayInput interface {
	pulumi.Input

	ToAdditionalConfigurationArrayOutput() AdditionalConfigurationArrayOutput
	ToAdditionalConfigurationArrayOutputWithContext(context.Context) AdditionalConfigurationArrayOutput
}

type AdditionalConfigurationArray []AdditionalConfigurationInput

func (AdditionalConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalConfiguration)(nil)).Elem()
}

func (i AdditionalConfigurationArray) ToAdditionalConfigurationArrayOutput() AdditionalConfigurationArrayOutput {
	return i.ToAdditionalConfigurationArrayOutputWithContext(context.Background())
}

func (i AdditionalConfigurationArray) ToAdditionalConfigurationArrayOutputWithContext(ctx context.Context) AdditionalConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalConfigurationArrayOutput)
}

type AdditionalConfigurationOutput struct{ *pulumi.OutputState }

func (AdditionalConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalConfiguration)(nil)).Elem()
}

func (o AdditionalConfigurationOutput) ToAdditionalConfigurationOutput() AdditionalConfigurationOutput {
	return o
}

func (o AdditionalConfigurationOutput) ToAdditionalConfigurationOutputWithContext(ctx context.Context) AdditionalConfigurationOutput {
	return o
}

func (o AdditionalConfigurationOutput) HierarchyInformation() HierarchyInformationOutput {
	return o.ApplyT(func(v AdditionalConfiguration) HierarchyInformation { return v.HierarchyInformation }).(HierarchyInformationOutput)
}

func (o AdditionalConfigurationOutput) Quantity() pulumi.IntOutput {
	return o.ApplyT(func(v AdditionalConfiguration) int { return v.Quantity }).(pulumi.IntOutput)
}

type AdditionalConfigurationArrayOutput struct{ *pulumi.OutputState }

func (AdditionalConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalConfiguration)(nil)).Elem()
}

func (o AdditionalConfigurationArrayOutput) ToAdditionalConfigurationArrayOutput() AdditionalConfigurationArrayOutput {
	return o
}

func (o AdditionalConfigurationArrayOutput) ToAdditionalConfigurationArrayOutputWithContext(ctx context.Context) AdditionalConfigurationArrayOutput {
	return o
}

func (o AdditionalConfigurationArrayOutput) Index(i pulumi.IntInput) AdditionalConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalConfiguration {
		return vs[0].([]AdditionalConfiguration)[vs[1].(int)]
	}).(AdditionalConfigurationOutput)
}

type AdditionalConfigurationResponse struct {
	HierarchyInformation HierarchyInformationResponse `pulumi:"hierarchyInformation"`
	Quantity             int                          `pulumi:"quantity"`
}

type AdditionalConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AdditionalConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalConfigurationResponse)(nil)).Elem()
}

func (o AdditionalConfigurationResponseOutput) ToAdditionalConfigurationResponseOutput() AdditionalConfigurationResponseOutput {
	return o
}

func (o AdditionalConfigurationResponseOutput) ToAdditionalConfigurationResponseOutputWithContext(ctx context.Context) AdditionalConfigurationResponseOutput {
	return o
}

func (o AdditionalConfigurationResponseOutput) HierarchyInformation() HierarchyInformationResponseOutput {
	return o.ApplyT(func(v AdditionalConfigurationResponse) HierarchyInformationResponse { return v.HierarchyInformation }).(HierarchyInformationResponseOutput)
}

func (o AdditionalConfigurationResponseOutput) Quantity() pulumi.IntOutput {
	return o.ApplyT(func(v AdditionalConfigurationResponse) int { return v.Quantity }).(pulumi.IntOutput)
}

type AdditionalConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (AdditionalConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalConfigurationResponse)(nil)).Elem()
}

func (o AdditionalConfigurationResponseArrayOutput) ToAdditionalConfigurationResponseArrayOutput() AdditionalConfigurationResponseArrayOutput {
	return o
}

func (o AdditionalConfigurationResponseArrayOutput) ToAdditionalConfigurationResponseArrayOutputWithContext(ctx context.Context) AdditionalConfigurationResponseArrayOutput {
	return o
}

func (o AdditionalConfigurationResponseArrayOutput) Index(i pulumi.IntInput) AdditionalConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalConfigurationResponse {
		return vs[0].([]AdditionalConfigurationResponse)[vs[1].(int)]
	}).(AdditionalConfigurationResponseOutput)
}

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

func (o AddressDetailsOutput) ForwardAddress() AddressPropertiesOutput {
	return o.ApplyT(func(v AddressDetails) AddressProperties { return v.ForwardAddress }).(AddressPropertiesOutput)
}

type AddressDetailsResponse struct {
	ForwardAddress AddressPropertiesResponse `pulumi:"forwardAddress"`
	ReturnAddress  AddressPropertiesResponse `pulumi:"returnAddress"`
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

func (o AddressDetailsResponseOutput) ForwardAddress() AddressPropertiesResponseOutput {
	return o.ApplyT(func(v AddressDetailsResponse) AddressPropertiesResponse { return v.ForwardAddress }).(AddressPropertiesResponseOutput)
}

func (o AddressDetailsResponseOutput) ReturnAddress() AddressPropertiesResponseOutput {
	return o.ApplyT(func(v AddressDetailsResponse) AddressPropertiesResponse { return v.ReturnAddress }).(AddressPropertiesResponseOutput)
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

func (o AddressPropertiesOutput) ContactDetails() ContactDetailsOutput {
	return o.ApplyT(func(v AddressProperties) ContactDetails { return v.ContactDetails }).(ContactDetailsOutput)
}

func (o AddressPropertiesOutput) ShippingAddress() ShippingAddressPtrOutput {
	return o.ApplyT(func(v AddressProperties) *ShippingAddress { return v.ShippingAddress }).(ShippingAddressPtrOutput)
}

type AddressPropertiesResponse struct {
	AddressValidationStatus string                   `pulumi:"addressValidationStatus"`
	ContactDetails          ContactDetailsResponse   `pulumi:"contactDetails"`
	ShippingAddress         *ShippingAddressResponse `pulumi:"shippingAddress"`
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

func (o AddressPropertiesResponseOutput) AddressValidationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v AddressPropertiesResponse) string { return v.AddressValidationStatus }).(pulumi.StringOutput)
}

func (o AddressPropertiesResponseOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v AddressPropertiesResponse) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o AddressPropertiesResponseOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v AddressPropertiesResponse) *ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponsePtrOutput)
}

type AvailabilityInformationResponse struct {
	AvailabilityStage     string `pulumi:"availabilityStage"`
	DisabledReason        string `pulumi:"disabledReason"`
	DisabledReasonMessage string `pulumi:"disabledReasonMessage"`
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

type CategoryInformationResponse struct {
	CategoryDisplayName *string        `pulumi:"categoryDisplayName"`
	CategoryName        *string        `pulumi:"categoryName"`
	Description         *string        `pulumi:"description"`
	Links               []LinkResponse `pulumi:"links"`
}

type CategoryInformationResponseOutput struct{ *pulumi.OutputState }

func (CategoryInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CategoryInformationResponse)(nil)).Elem()
}

func (o CategoryInformationResponseOutput) ToCategoryInformationResponseOutput() CategoryInformationResponseOutput {
	return o
}

func (o CategoryInformationResponseOutput) ToCategoryInformationResponseOutputWithContext(ctx context.Context) CategoryInformationResponseOutput {
	return o
}

func (o CategoryInformationResponseOutput) CategoryDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CategoryInformationResponse) *string { return v.CategoryDisplayName }).(pulumi.StringPtrOutput)
}

func (o CategoryInformationResponseOutput) CategoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CategoryInformationResponse) *string { return v.CategoryName }).(pulumi.StringPtrOutput)
}

func (o CategoryInformationResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CategoryInformationResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CategoryInformationResponseOutput) Links() LinkResponseArrayOutput {
	return o.ApplyT(func(v CategoryInformationResponse) []LinkResponse { return v.Links }).(LinkResponseArrayOutput)
}

type ChildConfigurationFilter struct {
	ChildConfigurationTypes []string               `pulumi:"childConfigurationTypes"`
	HierarchyInformations   []HierarchyInformation `pulumi:"hierarchyInformations"`
}





type ChildConfigurationFilterInput interface {
	pulumi.Input

	ToChildConfigurationFilterOutput() ChildConfigurationFilterOutput
	ToChildConfigurationFilterOutputWithContext(context.Context) ChildConfigurationFilterOutput
}

type ChildConfigurationFilterArgs struct {
	ChildConfigurationTypes pulumi.StringArrayInput        `pulumi:"childConfigurationTypes"`
	HierarchyInformations   HierarchyInformationArrayInput `pulumi:"hierarchyInformations"`
}

func (ChildConfigurationFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ChildConfigurationFilter)(nil)).Elem()
}

func (i ChildConfigurationFilterArgs) ToChildConfigurationFilterOutput() ChildConfigurationFilterOutput {
	return i.ToChildConfigurationFilterOutputWithContext(context.Background())
}

func (i ChildConfigurationFilterArgs) ToChildConfigurationFilterOutputWithContext(ctx context.Context) ChildConfigurationFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChildConfigurationFilterOutput)
}

func (i ChildConfigurationFilterArgs) ToChildConfigurationFilterPtrOutput() ChildConfigurationFilterPtrOutput {
	return i.ToChildConfigurationFilterPtrOutputWithContext(context.Background())
}

func (i ChildConfigurationFilterArgs) ToChildConfigurationFilterPtrOutputWithContext(ctx context.Context) ChildConfigurationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChildConfigurationFilterOutput).ToChildConfigurationFilterPtrOutputWithContext(ctx)
}









type ChildConfigurationFilterPtrInput interface {
	pulumi.Input

	ToChildConfigurationFilterPtrOutput() ChildConfigurationFilterPtrOutput
	ToChildConfigurationFilterPtrOutputWithContext(context.Context) ChildConfigurationFilterPtrOutput
}

type childConfigurationFilterPtrType ChildConfigurationFilterArgs

func ChildConfigurationFilterPtr(v *ChildConfigurationFilterArgs) ChildConfigurationFilterPtrInput {
	return (*childConfigurationFilterPtrType)(v)
}

func (*childConfigurationFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ChildConfigurationFilter)(nil)).Elem()
}

func (i *childConfigurationFilterPtrType) ToChildConfigurationFilterPtrOutput() ChildConfigurationFilterPtrOutput {
	return i.ToChildConfigurationFilterPtrOutputWithContext(context.Background())
}

func (i *childConfigurationFilterPtrType) ToChildConfigurationFilterPtrOutputWithContext(ctx context.Context) ChildConfigurationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChildConfigurationFilterPtrOutput)
}

type ChildConfigurationFilterOutput struct{ *pulumi.OutputState }

func (ChildConfigurationFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChildConfigurationFilter)(nil)).Elem()
}

func (o ChildConfigurationFilterOutput) ToChildConfigurationFilterOutput() ChildConfigurationFilterOutput {
	return o
}

func (o ChildConfigurationFilterOutput) ToChildConfigurationFilterOutputWithContext(ctx context.Context) ChildConfigurationFilterOutput {
	return o
}

func (o ChildConfigurationFilterOutput) ToChildConfigurationFilterPtrOutput() ChildConfigurationFilterPtrOutput {
	return o.ToChildConfigurationFilterPtrOutputWithContext(context.Background())
}

func (o ChildConfigurationFilterOutput) ToChildConfigurationFilterPtrOutputWithContext(ctx context.Context) ChildConfigurationFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ChildConfigurationFilter) *ChildConfigurationFilter {
		return &v
	}).(ChildConfigurationFilterPtrOutput)
}

func (o ChildConfigurationFilterOutput) ChildConfigurationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ChildConfigurationFilter) []string { return v.ChildConfigurationTypes }).(pulumi.StringArrayOutput)
}

func (o ChildConfigurationFilterOutput) HierarchyInformations() HierarchyInformationArrayOutput {
	return o.ApplyT(func(v ChildConfigurationFilter) []HierarchyInformation { return v.HierarchyInformations }).(HierarchyInformationArrayOutput)
}

type ChildConfigurationFilterPtrOutput struct{ *pulumi.OutputState }

func (ChildConfigurationFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChildConfigurationFilter)(nil)).Elem()
}

func (o ChildConfigurationFilterPtrOutput) ToChildConfigurationFilterPtrOutput() ChildConfigurationFilterPtrOutput {
	return o
}

func (o ChildConfigurationFilterPtrOutput) ToChildConfigurationFilterPtrOutputWithContext(ctx context.Context) ChildConfigurationFilterPtrOutput {
	return o
}

func (o ChildConfigurationFilterPtrOutput) Elem() ChildConfigurationFilterOutput {
	return o.ApplyT(func(v *ChildConfigurationFilter) ChildConfigurationFilter {
		if v != nil {
			return *v
		}
		var ret ChildConfigurationFilter
		return ret
	}).(ChildConfigurationFilterOutput)
}

func (o ChildConfigurationFilterPtrOutput) ChildConfigurationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ChildConfigurationFilter) []string {
		if v == nil {
			return nil
		}
		return v.ChildConfigurationTypes
	}).(pulumi.StringArrayOutput)
}

func (o ChildConfigurationFilterPtrOutput) HierarchyInformations() HierarchyInformationArrayOutput {
	return o.ApplyT(func(v *ChildConfigurationFilter) []HierarchyInformation {
		if v == nil {
			return nil
		}
		return v.HierarchyInformations
	}).(HierarchyInformationArrayOutput)
}

type ChildConfigurationResponse struct {
	AvailabilityInformation    AvailabilityInformationResponse      `pulumi:"availabilityInformation"`
	ChildConfigurationType     string                               `pulumi:"childConfigurationType"`
	ChildConfigurationTypes    []string                             `pulumi:"childConfigurationTypes"`
	CostInformation            CostInformationResponse              `pulumi:"costInformation"`
	Description                DescriptionResponse                  `pulumi:"description"`
	Dimensions                 DimensionsResponse                   `pulumi:"dimensions"`
	DisplayName                string                               `pulumi:"displayName"`
	FilterableProperties       []FilterablePropertyResponse         `pulumi:"filterableProperties"`
	FulfilledBy                string                               `pulumi:"fulfilledBy"`
	GroupedChildConfigurations []GroupedChildConfigurationsResponse `pulumi:"groupedChildConfigurations"`
	HierarchyInformation       HierarchyInformationResponse         `pulumi:"hierarchyInformation"`
	ImageInformation           []ImageInformationResponse           `pulumi:"imageInformation"`
	IsPartOfBaseConfiguration  bool                                 `pulumi:"isPartOfBaseConfiguration"`
	MaximumQuantity            int                                  `pulumi:"maximumQuantity"`
	MinimumQuantity            int                                  `pulumi:"minimumQuantity"`
	Specifications             []SpecificationResponse              `pulumi:"specifications"`
}

type ChildConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ChildConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChildConfigurationResponse)(nil)).Elem()
}

func (o ChildConfigurationResponseOutput) ToChildConfigurationResponseOutput() ChildConfigurationResponseOutput {
	return o
}

func (o ChildConfigurationResponseOutput) ToChildConfigurationResponseOutputWithContext(ctx context.Context) ChildConfigurationResponseOutput {
	return o
}

func (o ChildConfigurationResponseOutput) AvailabilityInformation() AvailabilityInformationResponseOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) AvailabilityInformationResponse { return v.AvailabilityInformation }).(AvailabilityInformationResponseOutput)
}

func (o ChildConfigurationResponseOutput) ChildConfigurationType() pulumi.StringOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) string { return v.ChildConfigurationType }).(pulumi.StringOutput)
}

func (o ChildConfigurationResponseOutput) ChildConfigurationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) []string { return v.ChildConfigurationTypes }).(pulumi.StringArrayOutput)
}

func (o ChildConfigurationResponseOutput) CostInformation() CostInformationResponseOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) CostInformationResponse { return v.CostInformation }).(CostInformationResponseOutput)
}

func (o ChildConfigurationResponseOutput) Description() DescriptionResponseOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) DescriptionResponse { return v.Description }).(DescriptionResponseOutput)
}

func (o ChildConfigurationResponseOutput) Dimensions() DimensionsResponseOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) DimensionsResponse { return v.Dimensions }).(DimensionsResponseOutput)
}

func (o ChildConfigurationResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ChildConfigurationResponseOutput) FilterableProperties() FilterablePropertyResponseArrayOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) []FilterablePropertyResponse { return v.FilterableProperties }).(FilterablePropertyResponseArrayOutput)
}

func (o ChildConfigurationResponseOutput) FulfilledBy() pulumi.StringOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) string { return v.FulfilledBy }).(pulumi.StringOutput)
}

func (o ChildConfigurationResponseOutput) GroupedChildConfigurations() GroupedChildConfigurationsResponseArrayOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) []GroupedChildConfigurationsResponse {
		return v.GroupedChildConfigurations
	}).(GroupedChildConfigurationsResponseArrayOutput)
}

func (o ChildConfigurationResponseOutput) HierarchyInformation() HierarchyInformationResponseOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) HierarchyInformationResponse { return v.HierarchyInformation }).(HierarchyInformationResponseOutput)
}

func (o ChildConfigurationResponseOutput) ImageInformation() ImageInformationResponseArrayOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) []ImageInformationResponse { return v.ImageInformation }).(ImageInformationResponseArrayOutput)
}

func (o ChildConfigurationResponseOutput) IsPartOfBaseConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) bool { return v.IsPartOfBaseConfiguration }).(pulumi.BoolOutput)
}

func (o ChildConfigurationResponseOutput) MaximumQuantity() pulumi.IntOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) int { return v.MaximumQuantity }).(pulumi.IntOutput)
}

func (o ChildConfigurationResponseOutput) MinimumQuantity() pulumi.IntOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) int { return v.MinimumQuantity }).(pulumi.IntOutput)
}

func (o ChildConfigurationResponseOutput) Specifications() SpecificationResponseArrayOutput {
	return o.ApplyT(func(v ChildConfigurationResponse) []SpecificationResponse { return v.Specifications }).(SpecificationResponseArrayOutput)
}

type ChildConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (ChildConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChildConfigurationResponse)(nil)).Elem()
}

func (o ChildConfigurationResponseArrayOutput) ToChildConfigurationResponseArrayOutput() ChildConfigurationResponseArrayOutput {
	return o
}

func (o ChildConfigurationResponseArrayOutput) ToChildConfigurationResponseArrayOutputWithContext(ctx context.Context) ChildConfigurationResponseArrayOutput {
	return o
}

func (o ChildConfigurationResponseArrayOutput) Index(i pulumi.IntInput) ChildConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChildConfigurationResponse {
		return vs[0].([]ChildConfigurationResponse)[vs[1].(int)]
	}).(ChildConfigurationResponseOutput)
}

type ConfigurationDeviceDetailsResponse struct {
	DeviceDetails        []DeviceDetailsResponse      `pulumi:"deviceDetails"`
	DisplayInfo          *DisplayInfoResponse         `pulumi:"displayInfo"`
	HierarchyInformation HierarchyInformationResponse `pulumi:"hierarchyInformation"`
	IdentificationType   string                       `pulumi:"identificationType"`
	Quantity             int                          `pulumi:"quantity"`
}

type ConfigurationDeviceDetailsResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationDeviceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationDeviceDetailsResponse)(nil)).Elem()
}

func (o ConfigurationDeviceDetailsResponseOutput) ToConfigurationDeviceDetailsResponseOutput() ConfigurationDeviceDetailsResponseOutput {
	return o
}

func (o ConfigurationDeviceDetailsResponseOutput) ToConfigurationDeviceDetailsResponseOutputWithContext(ctx context.Context) ConfigurationDeviceDetailsResponseOutput {
	return o
}

func (o ConfigurationDeviceDetailsResponseOutput) DeviceDetails() DeviceDetailsResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationDeviceDetailsResponse) []DeviceDetailsResponse { return v.DeviceDetails }).(DeviceDetailsResponseArrayOutput)
}

func (o ConfigurationDeviceDetailsResponseOutput) DisplayInfo() DisplayInfoResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationDeviceDetailsResponse) *DisplayInfoResponse { return v.DisplayInfo }).(DisplayInfoResponsePtrOutput)
}

func (o ConfigurationDeviceDetailsResponseOutput) HierarchyInformation() HierarchyInformationResponseOutput {
	return o.ApplyT(func(v ConfigurationDeviceDetailsResponse) HierarchyInformationResponse { return v.HierarchyInformation }).(HierarchyInformationResponseOutput)
}

func (o ConfigurationDeviceDetailsResponseOutput) IdentificationType() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationDeviceDetailsResponse) string { return v.IdentificationType }).(pulumi.StringOutput)
}

func (o ConfigurationDeviceDetailsResponseOutput) Quantity() pulumi.IntOutput {
	return o.ApplyT(func(v ConfigurationDeviceDetailsResponse) int { return v.Quantity }).(pulumi.IntOutput)
}

type ConfigurationDeviceDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationDeviceDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationDeviceDetailsResponse)(nil)).Elem()
}

func (o ConfigurationDeviceDetailsResponseArrayOutput) ToConfigurationDeviceDetailsResponseArrayOutput() ConfigurationDeviceDetailsResponseArrayOutput {
	return o
}

func (o ConfigurationDeviceDetailsResponseArrayOutput) ToConfigurationDeviceDetailsResponseArrayOutputWithContext(ctx context.Context) ConfigurationDeviceDetailsResponseArrayOutput {
	return o
}

func (o ConfigurationDeviceDetailsResponseArrayOutput) Index(i pulumi.IntInput) ConfigurationDeviceDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationDeviceDetailsResponse {
		return vs[0].([]ConfigurationDeviceDetailsResponse)[vs[1].(int)]
	}).(ConfigurationDeviceDetailsResponseOutput)
}

type ConfigurationFilter struct {
	ChildConfigurationFilter *ChildConfigurationFilter `pulumi:"childConfigurationFilter"`
	FilterableProperty       []FilterableProperty      `pulumi:"filterableProperty"`
	HierarchyInformation     HierarchyInformation      `pulumi:"hierarchyInformation"`
}





type ConfigurationFilterInput interface {
	pulumi.Input

	ToConfigurationFilterOutput() ConfigurationFilterOutput
	ToConfigurationFilterOutputWithContext(context.Context) ConfigurationFilterOutput
}

type ConfigurationFilterArgs struct {
	ChildConfigurationFilter ChildConfigurationFilterPtrInput `pulumi:"childConfigurationFilter"`
	FilterableProperty       FilterablePropertyArrayInput     `pulumi:"filterableProperty"`
	HierarchyInformation     HierarchyInformationInput        `pulumi:"hierarchyInformation"`
}

func (ConfigurationFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationFilter)(nil)).Elem()
}

func (i ConfigurationFilterArgs) ToConfigurationFilterOutput() ConfigurationFilterOutput {
	return i.ToConfigurationFilterOutputWithContext(context.Background())
}

func (i ConfigurationFilterArgs) ToConfigurationFilterOutputWithContext(ctx context.Context) ConfigurationFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationFilterOutput)
}

func (i ConfigurationFilterArgs) ToConfigurationFilterPtrOutput() ConfigurationFilterPtrOutput {
	return i.ToConfigurationFilterPtrOutputWithContext(context.Background())
}

func (i ConfigurationFilterArgs) ToConfigurationFilterPtrOutputWithContext(ctx context.Context) ConfigurationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationFilterOutput).ToConfigurationFilterPtrOutputWithContext(ctx)
}









type ConfigurationFilterPtrInput interface {
	pulumi.Input

	ToConfigurationFilterPtrOutput() ConfigurationFilterPtrOutput
	ToConfigurationFilterPtrOutputWithContext(context.Context) ConfigurationFilterPtrOutput
}

type configurationFilterPtrType ConfigurationFilterArgs

func ConfigurationFilterPtr(v *ConfigurationFilterArgs) ConfigurationFilterPtrInput {
	return (*configurationFilterPtrType)(v)
}

func (*configurationFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationFilter)(nil)).Elem()
}

func (i *configurationFilterPtrType) ToConfigurationFilterPtrOutput() ConfigurationFilterPtrOutput {
	return i.ToConfigurationFilterPtrOutputWithContext(context.Background())
}

func (i *configurationFilterPtrType) ToConfigurationFilterPtrOutputWithContext(ctx context.Context) ConfigurationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationFilterPtrOutput)
}

type ConfigurationFilterOutput struct{ *pulumi.OutputState }

func (ConfigurationFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationFilter)(nil)).Elem()
}

func (o ConfigurationFilterOutput) ToConfigurationFilterOutput() ConfigurationFilterOutput {
	return o
}

func (o ConfigurationFilterOutput) ToConfigurationFilterOutputWithContext(ctx context.Context) ConfigurationFilterOutput {
	return o
}

func (o ConfigurationFilterOutput) ToConfigurationFilterPtrOutput() ConfigurationFilterPtrOutput {
	return o.ToConfigurationFilterPtrOutputWithContext(context.Background())
}

func (o ConfigurationFilterOutput) ToConfigurationFilterPtrOutputWithContext(ctx context.Context) ConfigurationFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationFilter) *ConfigurationFilter {
		return &v
	}).(ConfigurationFilterPtrOutput)
}

func (o ConfigurationFilterOutput) ChildConfigurationFilter() ChildConfigurationFilterPtrOutput {
	return o.ApplyT(func(v ConfigurationFilter) *ChildConfigurationFilter { return v.ChildConfigurationFilter }).(ChildConfigurationFilterPtrOutput)
}

func (o ConfigurationFilterOutput) FilterableProperty() FilterablePropertyArrayOutput {
	return o.ApplyT(func(v ConfigurationFilter) []FilterableProperty { return v.FilterableProperty }).(FilterablePropertyArrayOutput)
}

func (o ConfigurationFilterOutput) HierarchyInformation() HierarchyInformationOutput {
	return o.ApplyT(func(v ConfigurationFilter) HierarchyInformation { return v.HierarchyInformation }).(HierarchyInformationOutput)
}

type ConfigurationFilterPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationFilter)(nil)).Elem()
}

func (o ConfigurationFilterPtrOutput) ToConfigurationFilterPtrOutput() ConfigurationFilterPtrOutput {
	return o
}

func (o ConfigurationFilterPtrOutput) ToConfigurationFilterPtrOutputWithContext(ctx context.Context) ConfigurationFilterPtrOutput {
	return o
}

func (o ConfigurationFilterPtrOutput) Elem() ConfigurationFilterOutput {
	return o.ApplyT(func(v *ConfigurationFilter) ConfigurationFilter {
		if v != nil {
			return *v
		}
		var ret ConfigurationFilter
		return ret
	}).(ConfigurationFilterOutput)
}

func (o ConfigurationFilterPtrOutput) ChildConfigurationFilter() ChildConfigurationFilterPtrOutput {
	return o.ApplyT(func(v *ConfigurationFilter) *ChildConfigurationFilter {
		if v == nil {
			return nil
		}
		return v.ChildConfigurationFilter
	}).(ChildConfigurationFilterPtrOutput)
}

func (o ConfigurationFilterPtrOutput) FilterableProperty() FilterablePropertyArrayOutput {
	return o.ApplyT(func(v *ConfigurationFilter) []FilterableProperty {
		if v == nil {
			return nil
		}
		return v.FilterableProperty
	}).(FilterablePropertyArrayOutput)
}

func (o ConfigurationFilterPtrOutput) HierarchyInformation() HierarchyInformationPtrOutput {
	return o.ApplyT(func(v *ConfigurationFilter) *HierarchyInformation {
		if v == nil {
			return nil
		}
		return &v.HierarchyInformation
	}).(HierarchyInformationPtrOutput)
}

type ConfigurationResponse struct {
	AvailabilityInformation    AvailabilityInformationResponse      `pulumi:"availabilityInformation"`
	ChildConfigurationTypes    []string                             `pulumi:"childConfigurationTypes"`
	CostInformation            CostInformationResponse              `pulumi:"costInformation"`
	Description                DescriptionResponse                  `pulumi:"description"`
	Dimensions                 DimensionsResponse                   `pulumi:"dimensions"`
	DisplayName                string                               `pulumi:"displayName"`
	FilterableProperties       []FilterablePropertyResponse         `pulumi:"filterableProperties"`
	FulfilledBy                string                               `pulumi:"fulfilledBy"`
	GroupedChildConfigurations []GroupedChildConfigurationsResponse `pulumi:"groupedChildConfigurations"`
	HierarchyInformation       HierarchyInformationResponse         `pulumi:"hierarchyInformation"`
	ImageInformation           []ImageInformationResponse           `pulumi:"imageInformation"`
	Specifications             []SpecificationResponse              `pulumi:"specifications"`
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

func (o ConfigurationResponseOutput) ChildConfigurationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConfigurationResponse) []string { return v.ChildConfigurationTypes }).(pulumi.StringArrayOutput)
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

func (o ConfigurationResponseOutput) FulfilledBy() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationResponse) string { return v.FulfilledBy }).(pulumi.StringOutput)
}

func (o ConfigurationResponseOutput) GroupedChildConfigurations() GroupedChildConfigurationsResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationResponse) []GroupedChildConfigurationsResponse {
		return v.GroupedChildConfigurations
	}).(GroupedChildConfigurationsResponseArrayOutput)
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

type ContactDetailsResponse struct {
	ContactName    string   `pulumi:"contactName"`
	EmailList      []string `pulumi:"emailList"`
	Mobile         *string  `pulumi:"mobile"`
	Phone          string   `pulumi:"phone"`
	PhoneExtension *string  `pulumi:"phoneExtension"`
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

type CostInformationResponse struct {
	BillingInfoUrl      string                        `pulumi:"billingInfoUrl"`
	BillingMeterDetails []BillingMeterDetailsResponse `pulumi:"billingMeterDetails"`
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
	ManagementResourceId       string `pulumi:"managementResourceId"`
	ManagementResourceTenantId string `pulumi:"managementResourceTenantId"`
	SerialNumber               string `pulumi:"serialNumber"`
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

func (o DeviceDetailsResponseOutput) ManagementResourceTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v DeviceDetailsResponse) string { return v.ManagementResourceTenantId }).(pulumi.StringOutput)
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

type ForwardShippingDetailsResponse struct {
	CarrierDisplayName string `pulumi:"carrierDisplayName"`
	CarrierName        string `pulumi:"carrierName"`
	TrackingId         string `pulumi:"trackingId"`
	TrackingUrl        string `pulumi:"trackingUrl"`
}

type ForwardShippingDetailsResponseOutput struct{ *pulumi.OutputState }

func (ForwardShippingDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardShippingDetailsResponse)(nil)).Elem()
}

func (o ForwardShippingDetailsResponseOutput) ToForwardShippingDetailsResponseOutput() ForwardShippingDetailsResponseOutput {
	return o
}

func (o ForwardShippingDetailsResponseOutput) ToForwardShippingDetailsResponseOutputWithContext(ctx context.Context) ForwardShippingDetailsResponseOutput {
	return o
}

func (o ForwardShippingDetailsResponseOutput) CarrierDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ForwardShippingDetailsResponse) string { return v.CarrierDisplayName }).(pulumi.StringOutput)
}

func (o ForwardShippingDetailsResponseOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v ForwardShippingDetailsResponse) string { return v.CarrierName }).(pulumi.StringOutput)
}

func (o ForwardShippingDetailsResponseOutput) TrackingId() pulumi.StringOutput {
	return o.ApplyT(func(v ForwardShippingDetailsResponse) string { return v.TrackingId }).(pulumi.StringOutput)
}

func (o ForwardShippingDetailsResponseOutput) TrackingUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ForwardShippingDetailsResponse) string { return v.TrackingUrl }).(pulumi.StringOutput)
}

type GroupedChildConfigurationsResponse struct {
	CategoryInformation CategoryInformationResponse  `pulumi:"categoryInformation"`
	ChildConfigurations []ChildConfigurationResponse `pulumi:"childConfigurations"`
}

type GroupedChildConfigurationsResponseOutput struct{ *pulumi.OutputState }

func (GroupedChildConfigurationsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupedChildConfigurationsResponse)(nil)).Elem()
}

func (o GroupedChildConfigurationsResponseOutput) ToGroupedChildConfigurationsResponseOutput() GroupedChildConfigurationsResponseOutput {
	return o
}

func (o GroupedChildConfigurationsResponseOutput) ToGroupedChildConfigurationsResponseOutputWithContext(ctx context.Context) GroupedChildConfigurationsResponseOutput {
	return o
}

func (o GroupedChildConfigurationsResponseOutput) CategoryInformation() CategoryInformationResponseOutput {
	return o.ApplyT(func(v GroupedChildConfigurationsResponse) CategoryInformationResponse { return v.CategoryInformation }).(CategoryInformationResponseOutput)
}

func (o GroupedChildConfigurationsResponseOutput) ChildConfigurations() ChildConfigurationResponseArrayOutput {
	return o.ApplyT(func(v GroupedChildConfigurationsResponse) []ChildConfigurationResponse { return v.ChildConfigurations }).(ChildConfigurationResponseArrayOutput)
}

type GroupedChildConfigurationsResponseArrayOutput struct{ *pulumi.OutputState }

func (GroupedChildConfigurationsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupedChildConfigurationsResponse)(nil)).Elem()
}

func (o GroupedChildConfigurationsResponseArrayOutput) ToGroupedChildConfigurationsResponseArrayOutput() GroupedChildConfigurationsResponseArrayOutput {
	return o
}

func (o GroupedChildConfigurationsResponseArrayOutput) ToGroupedChildConfigurationsResponseArrayOutputWithContext(ctx context.Context) GroupedChildConfigurationsResponseArrayOutput {
	return o
}

func (o GroupedChildConfigurationsResponseArrayOutput) Index(i pulumi.IntInput) GroupedChildConfigurationsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupedChildConfigurationsResponse {
		return vs[0].([]GroupedChildConfigurationsResponse)[vs[1].(int)]
	}).(GroupedChildConfigurationsResponseOutput)
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





type HierarchyInformationArrayInput interface {
	pulumi.Input

	ToHierarchyInformationArrayOutput() HierarchyInformationArrayOutput
	ToHierarchyInformationArrayOutputWithContext(context.Context) HierarchyInformationArrayOutput
}

type HierarchyInformationArray []HierarchyInformationInput

func (HierarchyInformationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HierarchyInformation)(nil)).Elem()
}

func (i HierarchyInformationArray) ToHierarchyInformationArrayOutput() HierarchyInformationArrayOutput {
	return i.ToHierarchyInformationArrayOutputWithContext(context.Background())
}

func (i HierarchyInformationArray) ToHierarchyInformationArrayOutputWithContext(ctx context.Context) HierarchyInformationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HierarchyInformationArrayOutput)
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

type HierarchyInformationArrayOutput struct{ *pulumi.OutputState }

func (HierarchyInformationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HierarchyInformation)(nil)).Elem()
}

func (o HierarchyInformationArrayOutput) ToHierarchyInformationArrayOutput() HierarchyInformationArrayOutput {
	return o
}

func (o HierarchyInformationArrayOutput) ToHierarchyInformationArrayOutputWithContext(ctx context.Context) HierarchyInformationArrayOutput {
	return o
}

func (o HierarchyInformationArrayOutput) Index(i pulumi.IntInput) HierarchyInformationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HierarchyInformation {
		return vs[0].([]HierarchyInformation)[vs[1].(int)]
	}).(HierarchyInformationOutput)
}

type HierarchyInformationResponse struct {
	ConfigurationName *string `pulumi:"configurationName"`
	ProductFamilyName *string `pulumi:"productFamilyName"`
	ProductLineName   *string `pulumi:"productLineName"`
	ProductName       *string `pulumi:"productName"`
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

type ImageInformationResponse struct {
	ImageType string `pulumi:"imageType"`
	ImageUrl  string `pulumi:"imageUrl"`
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
	OrderItemMode         *string        `pulumi:"orderItemMode"`
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
	OrderItemMode         pulumi.StringPtrInput   `pulumi:"orderItemMode"`
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

func (o OrderItemDetailsOutput) NotificationEmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OrderItemDetails) []string { return v.NotificationEmailList }).(pulumi.StringArrayOutput)
}

func (o OrderItemDetailsOutput) OrderItemMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrderItemDetails) *string { return v.OrderItemMode }).(pulumi.StringPtrOutput)
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

type OrderItemDetailsResponse struct {
	CancellationReason      string                            `pulumi:"cancellationReason"`
	CancellationStatus      string                            `pulumi:"cancellationStatus"`
	CurrentStage            StageDetailsResponse              `pulumi:"currentStage"`
	DeletionStatus          string                            `pulumi:"deletionStatus"`
	Error                   ErrorDetailResponse               `pulumi:"error"`
	ForwardShippingDetails  ForwardShippingDetailsResponse    `pulumi:"forwardShippingDetails"`
	ManagementRpDetailsList []ResourceProviderDetailsResponse `pulumi:"managementRpDetailsList"`
	NotificationEmailList   []string                          `pulumi:"notificationEmailList"`
	OrderItemMode           *string                           `pulumi:"orderItemMode"`
	OrderItemStageHistory   []StageDetailsResponse            `pulumi:"orderItemStageHistory"`
	OrderItemType           string                            `pulumi:"orderItemType"`
	Preferences             *PreferencesResponse              `pulumi:"preferences"`
	ProductDetails          ProductDetailsResponse            `pulumi:"productDetails"`
	ReturnReason            string                            `pulumi:"returnReason"`
	ReturnStatus            string                            `pulumi:"returnStatus"`
	ReverseShippingDetails  ReverseShippingDetailsResponse    `pulumi:"reverseShippingDetails"`
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

func (o OrderItemDetailsResponseOutput) ForwardShippingDetails() ForwardShippingDetailsResponseOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) ForwardShippingDetailsResponse { return v.ForwardShippingDetails }).(ForwardShippingDetailsResponseOutput)
}

func (o OrderItemDetailsResponseOutput) ManagementRpDetailsList() ResourceProviderDetailsResponseArrayOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) []ResourceProviderDetailsResponse { return v.ManagementRpDetailsList }).(ResourceProviderDetailsResponseArrayOutput)
}

func (o OrderItemDetailsResponseOutput) NotificationEmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) []string { return v.NotificationEmailList }).(pulumi.StringArrayOutput)
}

func (o OrderItemDetailsResponseOutput) OrderItemMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) *string { return v.OrderItemMode }).(pulumi.StringPtrOutput)
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

func (o OrderItemDetailsResponseOutput) ReverseShippingDetails() ReverseShippingDetailsResponseOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) ReverseShippingDetailsResponse { return v.ReverseShippingDetails }).(ReverseShippingDetailsResponseOutput)
}

type Pav2MeterDetailsResponse struct {
	BillingType  string  `pulumi:"billingType"`
	ChargingType string  `pulumi:"chargingType"`
	MeterGuid    string  `pulumi:"meterGuid"`
	Multiplier   float64 `pulumi:"multiplier"`
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
	HierarchyInformation          HierarchyInformation      `pulumi:"hierarchyInformation"`
	OptInAdditionalConfigurations []AdditionalConfiguration `pulumi:"optInAdditionalConfigurations"`
}





type ProductDetailsInput interface {
	pulumi.Input

	ToProductDetailsOutput() ProductDetailsOutput
	ToProductDetailsOutputWithContext(context.Context) ProductDetailsOutput
}

type ProductDetailsArgs struct {
	HierarchyInformation          HierarchyInformationInput         `pulumi:"hierarchyInformation"`
	OptInAdditionalConfigurations AdditionalConfigurationArrayInput `pulumi:"optInAdditionalConfigurations"`
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

func (o ProductDetailsOutput) HierarchyInformation() HierarchyInformationOutput {
	return o.ApplyT(func(v ProductDetails) HierarchyInformation { return v.HierarchyInformation }).(HierarchyInformationOutput)
}

func (o ProductDetailsOutput) OptInAdditionalConfigurations() AdditionalConfigurationArrayOutput {
	return o.ApplyT(func(v ProductDetails) []AdditionalConfiguration { return v.OptInAdditionalConfigurations }).(AdditionalConfigurationArrayOutput)
}

type ProductDetailsResponse struct {
	ChildConfigurationDeviceDetails []ConfigurationDeviceDetailsResponse `pulumi:"childConfigurationDeviceDetails"`
	DisplayInfo                     *DisplayInfoResponse                 `pulumi:"displayInfo"`
	HierarchyInformation            HierarchyInformationResponse         `pulumi:"hierarchyInformation"`
	IdentificationType              string                               `pulumi:"identificationType"`
	OptInAdditionalConfigurations   []AdditionalConfigurationResponse    `pulumi:"optInAdditionalConfigurations"`
	ParentDeviceDetails             DeviceDetailsResponse                `pulumi:"parentDeviceDetails"`
	ProductDoubleEncryptionStatus   string                               `pulumi:"productDoubleEncryptionStatus"`
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

func (o ProductDetailsResponseOutput) ChildConfigurationDeviceDetails() ConfigurationDeviceDetailsResponseArrayOutput {
	return o.ApplyT(func(v ProductDetailsResponse) []ConfigurationDeviceDetailsResponse {
		return v.ChildConfigurationDeviceDetails
	}).(ConfigurationDeviceDetailsResponseArrayOutput)
}

func (o ProductDetailsResponseOutput) DisplayInfo() DisplayInfoResponsePtrOutput {
	return o.ApplyT(func(v ProductDetailsResponse) *DisplayInfoResponse { return v.DisplayInfo }).(DisplayInfoResponsePtrOutput)
}

func (o ProductDetailsResponseOutput) HierarchyInformation() HierarchyInformationResponseOutput {
	return o.ApplyT(func(v ProductDetailsResponse) HierarchyInformationResponse { return v.HierarchyInformation }).(HierarchyInformationResponseOutput)
}

func (o ProductDetailsResponseOutput) IdentificationType() pulumi.StringOutput {
	return o.ApplyT(func(v ProductDetailsResponse) string { return v.IdentificationType }).(pulumi.StringOutput)
}

func (o ProductDetailsResponseOutput) OptInAdditionalConfigurations() AdditionalConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ProductDetailsResponse) []AdditionalConfigurationResponse {
		return v.OptInAdditionalConfigurations
	}).(AdditionalConfigurationResponseArrayOutput)
}

func (o ProductDetailsResponseOutput) ParentDeviceDetails() DeviceDetailsResponseOutput {
	return o.ApplyT(func(v ProductDetailsResponse) DeviceDetailsResponse { return v.ParentDeviceDetails }).(DeviceDetailsResponseOutput)
}

func (o ProductDetailsResponseOutput) ProductDoubleEncryptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ProductDetailsResponse) string { return v.ProductDoubleEncryptionStatus }).(pulumi.StringOutput)
}

type ProductFamilyResponse struct {
	AvailabilityInformation AvailabilityInformationResponse   `pulumi:"availabilityInformation"`
	CostInformation         CostInformationResponse           `pulumi:"costInformation"`
	Description             DescriptionResponse               `pulumi:"description"`
	DisplayName             string                            `pulumi:"displayName"`
	FilterableProperties    []FilterablePropertyResponse      `pulumi:"filterableProperties"`
	FulfilledBy             string                            `pulumi:"fulfilledBy"`
	HierarchyInformation    HierarchyInformationResponse      `pulumi:"hierarchyInformation"`
	ImageInformation        []ImageInformationResponse        `pulumi:"imageInformation"`
	ProductLines            []ProductLineResponse             `pulumi:"productLines"`
	ResourceProviderDetails []ResourceProviderDetailsResponse `pulumi:"resourceProviderDetails"`
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

func (o ProductFamilyResponseOutput) FulfilledBy() pulumi.StringOutput {
	return o.ApplyT(func(v ProductFamilyResponse) string { return v.FulfilledBy }).(pulumi.StringOutput)
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

func (o ProductFamilyResponseOutput) ResourceProviderDetails() ResourceProviderDetailsResponseArrayOutput {
	return o.ApplyT(func(v ProductFamilyResponse) []ResourceProviderDetailsResponse { return v.ResourceProviderDetails }).(ResourceProviderDetailsResponseArrayOutput)
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
	FulfilledBy             string                          `pulumi:"fulfilledBy"`
	HierarchyInformation    HierarchyInformationResponse    `pulumi:"hierarchyInformation"`
	ImageInformation        []ImageInformationResponse      `pulumi:"imageInformation"`
	Products                []ProductResponse               `pulumi:"products"`
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

func (o ProductLineResponseOutput) FulfilledBy() pulumi.StringOutput {
	return o.ApplyT(func(v ProductLineResponse) string { return v.FulfilledBy }).(pulumi.StringOutput)
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
	FulfilledBy             string                          `pulumi:"fulfilledBy"`
	HierarchyInformation    HierarchyInformationResponse    `pulumi:"hierarchyInformation"`
	ImageInformation        []ImageInformationResponse      `pulumi:"imageInformation"`
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

func (o ProductResponseOutput) FulfilledBy() pulumi.StringOutput {
	return o.ApplyT(func(v ProductResponse) string { return v.FulfilledBy }).(pulumi.StringOutput)
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

type ResourceProviderDetailsResponse struct {
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
}

type ResourceProviderDetailsResponseOutput struct{ *pulumi.OutputState }

func (ResourceProviderDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceProviderDetailsResponse)(nil)).Elem()
}

func (o ResourceProviderDetailsResponseOutput) ToResourceProviderDetailsResponseOutput() ResourceProviderDetailsResponseOutput {
	return o
}

func (o ResourceProviderDetailsResponseOutput) ToResourceProviderDetailsResponseOutputWithContext(ctx context.Context) ResourceProviderDetailsResponseOutput {
	return o
}

func (o ResourceProviderDetailsResponseOutput) ResourceProviderNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceProviderDetailsResponse) string { return v.ResourceProviderNamespace }).(pulumi.StringOutput)
}

type ResourceProviderDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceProviderDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceProviderDetailsResponse)(nil)).Elem()
}

func (o ResourceProviderDetailsResponseArrayOutput) ToResourceProviderDetailsResponseArrayOutput() ResourceProviderDetailsResponseArrayOutput {
	return o
}

func (o ResourceProviderDetailsResponseArrayOutput) ToResourceProviderDetailsResponseArrayOutputWithContext(ctx context.Context) ResourceProviderDetailsResponseArrayOutput {
	return o
}

func (o ResourceProviderDetailsResponseArrayOutput) Index(i pulumi.IntInput) ResourceProviderDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceProviderDetailsResponse {
		return vs[0].([]ResourceProviderDetailsResponse)[vs[1].(int)]
	}).(ResourceProviderDetailsResponseOutput)
}

type ReverseShippingDetailsResponse struct {
	CarrierDisplayName string `pulumi:"carrierDisplayName"`
	CarrierName        string `pulumi:"carrierName"`
	SasKeyForLabel     string `pulumi:"sasKeyForLabel"`
	TrackingId         string `pulumi:"trackingId"`
	TrackingUrl        string `pulumi:"trackingUrl"`
}

type ReverseShippingDetailsResponseOutput struct{ *pulumi.OutputState }

func (ReverseShippingDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReverseShippingDetailsResponse)(nil)).Elem()
}

func (o ReverseShippingDetailsResponseOutput) ToReverseShippingDetailsResponseOutput() ReverseShippingDetailsResponseOutput {
	return o
}

func (o ReverseShippingDetailsResponseOutput) ToReverseShippingDetailsResponseOutputWithContext(ctx context.Context) ReverseShippingDetailsResponseOutput {
	return o
}

func (o ReverseShippingDetailsResponseOutput) CarrierDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseShippingDetailsResponse) string { return v.CarrierDisplayName }).(pulumi.StringOutput)
}

func (o ReverseShippingDetailsResponseOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseShippingDetailsResponse) string { return v.CarrierName }).(pulumi.StringOutput)
}

func (o ReverseShippingDetailsResponseOutput) SasKeyForLabel() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseShippingDetailsResponse) string { return v.SasKeyForLabel }).(pulumi.StringOutput)
}

func (o ReverseShippingDetailsResponseOutput) TrackingId() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseShippingDetailsResponse) string { return v.TrackingId }).(pulumi.StringOutput)
}

func (o ReverseShippingDetailsResponseOutput) TrackingUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseShippingDetailsResponse) string { return v.TrackingUrl }).(pulumi.StringOutput)
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

type SpecificationResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
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
	return reflect.TypeOf((*map[string][]FilterableProperty)(nil)).Elem()
}

func (i FilterablePropertyArrayMap) ToFilterablePropertyArrayMapOutput() FilterablePropertyArrayMapOutput {
	return i.ToFilterablePropertyArrayMapOutputWithContext(context.Background())
}

func (i FilterablePropertyArrayMap) ToFilterablePropertyArrayMapOutputWithContext(ctx context.Context) FilterablePropertyArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterablePropertyArrayMapOutput)
}





type FilterablePropertyArrayMapInput interface {
	pulumi.Input

	ToFilterablePropertyArrayMapOutput() FilterablePropertyArrayMapOutput
	ToFilterablePropertyArrayMapOutputWithContext(context.Context) FilterablePropertyArrayMapOutput
}

type FilterablePropertyArrayMapOutput struct{ *pulumi.OutputState }

func (FilterablePropertyArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string][]FilterableProperty)(nil)).Elem()
}

func (o FilterablePropertyArrayMapOutput) ToFilterablePropertyArrayMapOutput() FilterablePropertyArrayMapOutput {
	return o
}

func (o FilterablePropertyArrayMapOutput) ToFilterablePropertyArrayMapOutputWithContext(ctx context.Context) FilterablePropertyArrayMapOutput {
	return o
}

func (o FilterablePropertyArrayMapOutput) MapIndex(k pulumi.StringInput) FilterablePropertyArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) []FilterableProperty {
		return vs[0].(map[string][]FilterableProperty)[vs[1].(string)]
	}).(FilterablePropertyArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AdditionalConfigurationOutput{})
	pulumi.RegisterOutputType(AdditionalConfigurationArrayOutput{})
	pulumi.RegisterOutputType(AdditionalConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AdditionalConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(AddressDetailsOutput{})
	pulumi.RegisterOutputType(AddressDetailsResponseOutput{})
	pulumi.RegisterOutputType(AddressPropertiesOutput{})
	pulumi.RegisterOutputType(AddressPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AvailabilityInformationResponseOutput{})
	pulumi.RegisterOutputType(BillingMeterDetailsResponseOutput{})
	pulumi.RegisterOutputType(BillingMeterDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(CategoryInformationResponseOutput{})
	pulumi.RegisterOutputType(ChildConfigurationFilterOutput{})
	pulumi.RegisterOutputType(ChildConfigurationFilterPtrOutput{})
	pulumi.RegisterOutputType(ChildConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ChildConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationDeviceDetailsResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationDeviceDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationFilterOutput{})
	pulumi.RegisterOutputType(ConfigurationFilterPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(ContactDetailsOutput{})
	pulumi.RegisterOutputType(ContactDetailsResponseOutput{})
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
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(FilterablePropertyOutput{})
	pulumi.RegisterOutputType(FilterablePropertyArrayOutput{})
	pulumi.RegisterOutputType(FilterablePropertyResponseOutput{})
	pulumi.RegisterOutputType(FilterablePropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(ForwardShippingDetailsResponseOutput{})
	pulumi.RegisterOutputType(GroupedChildConfigurationsResponseOutput{})
	pulumi.RegisterOutputType(GroupedChildConfigurationsResponseArrayOutput{})
	pulumi.RegisterOutputType(HierarchyInformationOutput{})
	pulumi.RegisterOutputType(HierarchyInformationPtrOutput{})
	pulumi.RegisterOutputType(HierarchyInformationArrayOutput{})
	pulumi.RegisterOutputType(HierarchyInformationResponseOutput{})
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
	pulumi.RegisterOutputType(OrderItemDetailsResponseOutput{})
	pulumi.RegisterOutputType(PreferencesOutput{})
	pulumi.RegisterOutputType(PreferencesPtrOutput{})
	pulumi.RegisterOutputType(PreferencesResponseOutput{})
	pulumi.RegisterOutputType(PreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(ProductDetailsOutput{})
	pulumi.RegisterOutputType(ProductDetailsResponseOutput{})
	pulumi.RegisterOutputType(ProductFamilyResponseOutput{})
	pulumi.RegisterOutputType(ProductFamilyResponseArrayOutput{})
	pulumi.RegisterOutputType(ProductLineResponseOutput{})
	pulumi.RegisterOutputType(ProductLineResponseArrayOutput{})
	pulumi.RegisterOutputType(ProductResponseOutput{})
	pulumi.RegisterOutputType(ProductResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceProviderDetailsResponseOutput{})
	pulumi.RegisterOutputType(ResourceProviderDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(ReverseShippingDetailsResponseOutput{})
	pulumi.RegisterOutputType(ShippingAddressOutput{})
	pulumi.RegisterOutputType(ShippingAddressPtrOutput{})
	pulumi.RegisterOutputType(ShippingAddressResponseOutput{})
	pulumi.RegisterOutputType(ShippingAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(SpecificationResponseOutput{})
	pulumi.RegisterOutputType(SpecificationResponseArrayOutput{})
	pulumi.RegisterOutputType(StageDetailsResponseOutput{})
	pulumi.RegisterOutputType(StageDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TransportPreferencesOutput{})
	pulumi.RegisterOutputType(TransportPreferencesPtrOutput{})
	pulumi.RegisterOutputType(TransportPreferencesResponseOutput{})
	pulumi.RegisterOutputType(TransportPreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(FilterablePropertyArrayMapOutput{})
}
