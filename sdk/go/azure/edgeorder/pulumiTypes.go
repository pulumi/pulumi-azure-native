


package edgeorder

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

type BillingMeterDetailsResponse struct {
	Frequency    string      `pulumi:"frequency"`
	MeterDetails interface{} `pulumi:"meterDetails"`
	MeteringType string      `pulumi:"meteringType"`
	Name         string      `pulumi:"name"`
}

type ConfigurationFilters struct {
	FilterableProperty   []FilterableProperty `pulumi:"filterableProperty"`
	HierarchyInformation HierarchyInformation `pulumi:"hierarchyInformation"`
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

type CustomerSubscriptionDetails struct {
	LocationPlacementId *string                                  `pulumi:"locationPlacementId"`
	QuotaId             string                                   `pulumi:"quotaId"`
	RegisteredFeatures  []CustomerSubscriptionRegisteredFeatures `pulumi:"registeredFeatures"`
}

type CustomerSubscriptionRegisteredFeatures struct {
	Name  *string `pulumi:"name"`
	State *string `pulumi:"state"`
}

type DescriptionResponse struct {
	Attributes       []string       `pulumi:"attributes"`
	DescriptionType  string         `pulumi:"descriptionType"`
	Keywords         []string       `pulumi:"keywords"`
	Links            []LinkResponse `pulumi:"links"`
	LongDescription  string         `pulumi:"longDescription"`
	ShortDescription string         `pulumi:"shortDescription"`
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

type FilterablePropertyResponse struct {
	SupportedValues []string `pulumi:"supportedValues"`
	Type            string   `pulumi:"type"`
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

type LinkResponse struct {
	LinkType string `pulumi:"linkType"`
	LinkUrl  string `pulumi:"linkUrl"`
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
	ManagementRpDetails     ResourceProviderDetailsResponse   `pulumi:"managementRpDetails"`
	ManagementRpDetailsList []ResourceProviderDetailsResponse `pulumi:"managementRpDetailsList"`
	NotificationEmailList   []string                          `pulumi:"notificationEmailList"`
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

func (o OrderItemDetailsResponseOutput) ManagementRpDetails() ResourceProviderDetailsResponseOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) ResourceProviderDetailsResponse { return v.ManagementRpDetails }).(ResourceProviderDetailsResponseOutput)
}

func (o OrderItemDetailsResponseOutput) ManagementRpDetailsList() ResourceProviderDetailsResponseArrayOutput {
	return o.ApplyT(func(v OrderItemDetailsResponse) []ResourceProviderDetailsResponse { return v.ManagementRpDetailsList }).(ResourceProviderDetailsResponseArrayOutput)
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
	HierarchyInformation HierarchyInformation `pulumi:"hierarchyInformation"`
}





type ProductDetailsInput interface {
	pulumi.Input

	ToProductDetailsOutput() ProductDetailsOutput
	ToProductDetailsOutputWithContext(context.Context) ProductDetailsOutput
}

type ProductDetailsArgs struct {
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

type ProductDetailsResponse struct {
	Count                         int                          `pulumi:"count"`
	DeviceDetails                 []DeviceDetailsResponse      `pulumi:"deviceDetails"`
	DisplayInfo                   *DisplayInfoResponse         `pulumi:"displayInfo"`
	HierarchyInformation          HierarchyInformationResponse `pulumi:"hierarchyInformation"`
	ProductDoubleEncryptionStatus string                       `pulumi:"productDoubleEncryptionStatus"`
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

func (o ProductDetailsResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v ProductDetailsResponse) int { return v.Count }).(pulumi.IntOutput)
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

func (o ProductDetailsResponseOutput) ProductDoubleEncryptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ProductDetailsResponse) string { return v.ProductDoubleEncryptionStatus }).(pulumi.StringOutput)
}

type ProductFamilyResponse struct {
	AvailabilityInformation AvailabilityInformationResponse   `pulumi:"availabilityInformation"`
	CostInformation         CostInformationResponse           `pulumi:"costInformation"`
	Description             DescriptionResponse               `pulumi:"description"`
	DisplayName             string                            `pulumi:"displayName"`
	FilterableProperties    []FilterablePropertyResponse      `pulumi:"filterableProperties"`
	HierarchyInformation    HierarchyInformationResponse      `pulumi:"hierarchyInformation"`
	ImageInformation        []ImageInformationResponse        `pulumi:"imageInformation"`
	ProductLines            []ProductLineResponse             `pulumi:"productLines"`
	ResourceProviderDetails []ResourceProviderDetailsResponse `pulumi:"resourceProviderDetails"`
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

func init() {
	pulumi.RegisterOutputType(AddressDetailsOutput{})
	pulumi.RegisterOutputType(AddressDetailsResponseOutput{})
	pulumi.RegisterOutputType(AddressPropertiesOutput{})
	pulumi.RegisterOutputType(AddressPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ContactDetailsOutput{})
	pulumi.RegisterOutputType(ContactDetailsResponseOutput{})
	pulumi.RegisterOutputType(DeviceDetailsResponseOutput{})
	pulumi.RegisterOutputType(DeviceDetailsResponseArrayOutput{})
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
	pulumi.RegisterOutputType(ForwardShippingDetailsResponseOutput{})
	pulumi.RegisterOutputType(HierarchyInformationOutput{})
	pulumi.RegisterOutputType(HierarchyInformationResponseOutput{})
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
	pulumi.RegisterOutputType(ResourceProviderDetailsResponseOutput{})
	pulumi.RegisterOutputType(ResourceProviderDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(ReverseShippingDetailsResponseOutput{})
	pulumi.RegisterOutputType(ShippingAddressOutput{})
	pulumi.RegisterOutputType(ShippingAddressPtrOutput{})
	pulumi.RegisterOutputType(ShippingAddressResponseOutput{})
	pulumi.RegisterOutputType(ShippingAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(StageDetailsResponseOutput{})
	pulumi.RegisterOutputType(StageDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TransportPreferencesOutput{})
	pulumi.RegisterOutputType(TransportPreferencesPtrOutput{})
	pulumi.RegisterOutputType(TransportPreferencesResponseOutput{})
	pulumi.RegisterOutputType(TransportPreferencesResponsePtrOutput{})
}
