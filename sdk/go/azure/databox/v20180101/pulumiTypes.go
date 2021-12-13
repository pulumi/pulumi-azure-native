


package v20180101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountCredentialDetailsResponse struct {
	AccountConnectionString string                           `pulumi:"accountConnectionString"`
	AccountName             string                           `pulumi:"accountName"`
	ShareCredentialDetails  []ShareCredentialDetailsResponse `pulumi:"shareCredentialDetails"`
}

type ApplianceNetworkConfigurationResponse struct {
	MacAddress string `pulumi:"macAddress"`
	Name       string `pulumi:"name"`
}

type ContactDetails struct {
	ContactName            string                   `pulumi:"contactName"`
	EmailList              []string                 `pulumi:"emailList"`
	Mobile                 *string                  `pulumi:"mobile"`
	NotificationPreference []NotificationPreference `pulumi:"notificationPreference"`
	Phone                  string                   `pulumi:"phone"`
	PhoneExtension         *string                  `pulumi:"phoneExtension"`
}

type ContactDetailsResponse struct {
	ContactName            string                           `pulumi:"contactName"`
	EmailList              []string                         `pulumi:"emailList"`
	Mobile                 *string                          `pulumi:"mobile"`
	NotificationPreference []NotificationPreferenceResponse `pulumi:"notificationPreference"`
	Phone                  string                           `pulumi:"phone"`
	PhoneExtension         *string                          `pulumi:"phoneExtension"`
}

type CopyProgressResponse struct {
	AccountId           string  `pulumi:"accountId"`
	BytesSentToCloud    float64 `pulumi:"bytesSentToCloud"`
	FilesProcessed      float64 `pulumi:"filesProcessed"`
	StorageAccountName  string  `pulumi:"storageAccountName"`
	TotalBytesToProcess float64 `pulumi:"totalBytesToProcess"`
	TotalFilesToProcess float64 `pulumi:"totalFilesToProcess"`
}

type DataBoxAccountCopyLogDetailsResponse struct {
	AccountName        string `pulumi:"accountName"`
	CopyLogDetailsType string `pulumi:"copyLogDetailsType"`
	CopyLogLink        string `pulumi:"copyLogLink"`
}

type DataBoxDiskCopyLogDetailsResponse struct {
	CopyLogDetailsType string `pulumi:"copyLogDetailsType"`
	DiskSerialNumber   string `pulumi:"diskSerialNumber"`
	ErrorLogLink       string `pulumi:"errorLogLink"`
	VerboseLogLink     string `pulumi:"verboseLogLink"`
}

type DataBoxDiskCopyProgressResponse struct {
	BytesCopied     float64 `pulumi:"bytesCopied"`
	PercentComplete int     `pulumi:"percentComplete"`
	SerialNumber    string  `pulumi:"serialNumber"`
	Status          string  `pulumi:"status"`
}

type DataBoxDiskJobDetails struct {
	ContactDetails              ContactDetails  `pulumi:"contactDetails"`
	DestinationAccountDetails   []interface{}   `pulumi:"destinationAccountDetails"`
	ExpectedDataSizeInTeraBytes *int            `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string          `pulumi:"jobDetailsType"`
	Passkey                     *string         `pulumi:"passkey"`
	Preferences                 *Preferences    `pulumi:"preferences"`
	PreferredDisks              map[string]int  `pulumi:"preferredDisks"`
	ShippingAddress             ShippingAddress `pulumi:"shippingAddress"`
}


func (val *DataBoxDiskJobDetails) Defaults() *DataBoxDiskJobDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ShippingAddress = *tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxDiskJobDetailsResponse struct {
	ChainOfCustodySasKey        string                            `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse            `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                     `pulumi:"copyLogDetails"`
	CopyProgress                []DataBoxDiskCopyProgressResponse `pulumi:"copyProgress"`
	DeliveryPackage             PackageShippingDetailsResponse    `pulumi:"deliveryPackage"`
	DestinationAccountDetails   []interface{}                     `pulumi:"destinationAccountDetails"`
	DisksAndSizeDetails         map[string]int                    `pulumi:"disksAndSizeDetails"`
	ErrorDetails                []JobErrorDetailsResponse         `pulumi:"errorDetails"`
	ExpectedDataSizeInTeraBytes *int                              `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string                            `pulumi:"jobDetailsType"`
	JobStages                   []JobStagesResponse               `pulumi:"jobStages"`
	Passkey                     *string                           `pulumi:"passkey"`
	Preferences                 *PreferencesResponse              `pulumi:"preferences"`
	PreferredDisks              map[string]int                    `pulumi:"preferredDisks"`
	ReturnPackage               PackageShippingDetailsResponse    `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  string                            `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             ShippingAddressResponse           `pulumi:"shippingAddress"`
}


func (val *DataBoxDiskJobDetailsResponse) Defaults() *DataBoxDiskJobDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ShippingAddress = *tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxDiskJobSecretsResponse struct {
	DiskSecrets          []DiskSecretResponse `pulumi:"diskSecrets"`
	IsPasskeyUserDefined bool                 `pulumi:"isPasskeyUserDefined"`
	JobSecretsType       string               `pulumi:"jobSecretsType"`
	PassKey              string               `pulumi:"passKey"`
}

type DataBoxHeavyAccountCopyLogDetailsResponse struct {
	AccountName        string   `pulumi:"accountName"`
	CopyLogDetailsType string   `pulumi:"copyLogDetailsType"`
	CopyLogLink        []string `pulumi:"copyLogLink"`
}

type DataBoxHeavyJobDetails struct {
	ContactDetails              ContactDetails  `pulumi:"contactDetails"`
	DestinationAccountDetails   []interface{}   `pulumi:"destinationAccountDetails"`
	ExpectedDataSizeInTeraBytes *int            `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string          `pulumi:"jobDetailsType"`
	Preferences                 *Preferences    `pulumi:"preferences"`
	ShippingAddress             ShippingAddress `pulumi:"shippingAddress"`
}


func (val *DataBoxHeavyJobDetails) Defaults() *DataBoxHeavyJobDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ShippingAddress = *tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxHeavyJobDetailsResponse struct {
	ChainOfCustodySasKey        string                         `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse         `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                  `pulumi:"copyLogDetails"`
	CopyProgress                []CopyProgressResponse         `pulumi:"copyProgress"`
	DeliveryPackage             PackageShippingDetailsResponse `pulumi:"deliveryPackage"`
	DestinationAccountDetails   []interface{}                  `pulumi:"destinationAccountDetails"`
	ErrorDetails                []JobErrorDetailsResponse      `pulumi:"errorDetails"`
	ExpectedDataSizeInTeraBytes *int                           `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string                         `pulumi:"jobDetailsType"`
	JobStages                   []JobStagesResponse            `pulumi:"jobStages"`
	Preferences                 *PreferencesResponse           `pulumi:"preferences"`
	ReturnPackage               PackageShippingDetailsResponse `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  string                         `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             ShippingAddressResponse        `pulumi:"shippingAddress"`
}


func (val *DataBoxHeavyJobDetailsResponse) Defaults() *DataBoxHeavyJobDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ShippingAddress = *tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxHeavyJobSecretsResponse struct {
	CabinetPodSecrets []DataBoxHeavySecretResponse `pulumi:"cabinetPodSecrets"`
	JobSecretsType    string                       `pulumi:"jobSecretsType"`
}

type DataBoxHeavySecretResponse struct {
	AccountCredentialDetails    []AccountCredentialDetailsResponse      `pulumi:"accountCredentialDetails"`
	DevicePassword              string                                  `pulumi:"devicePassword"`
	DeviceSerialNumber          string                                  `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey string                                  `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       []ApplianceNetworkConfigurationResponse `pulumi:"networkConfigurations"`
}

type DataBoxJobDetails struct {
	ContactDetails              ContactDetails  `pulumi:"contactDetails"`
	DestinationAccountDetails   []interface{}   `pulumi:"destinationAccountDetails"`
	ExpectedDataSizeInTeraBytes *int            `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string          `pulumi:"jobDetailsType"`
	Preferences                 *Preferences    `pulumi:"preferences"`
	ShippingAddress             ShippingAddress `pulumi:"shippingAddress"`
}


func (val *DataBoxJobDetails) Defaults() *DataBoxJobDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ShippingAddress = *tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxJobDetailsResponse struct {
	ChainOfCustodySasKey        string                         `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse         `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                  `pulumi:"copyLogDetails"`
	CopyProgress                []CopyProgressResponse         `pulumi:"copyProgress"`
	DeliveryPackage             PackageShippingDetailsResponse `pulumi:"deliveryPackage"`
	DestinationAccountDetails   []interface{}                  `pulumi:"destinationAccountDetails"`
	ErrorDetails                []JobErrorDetailsResponse      `pulumi:"errorDetails"`
	ExpectedDataSizeInTeraBytes *int                           `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string                         `pulumi:"jobDetailsType"`
	JobStages                   []JobStagesResponse            `pulumi:"jobStages"`
	Preferences                 *PreferencesResponse           `pulumi:"preferences"`
	ReturnPackage               PackageShippingDetailsResponse `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  string                         `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             ShippingAddressResponse        `pulumi:"shippingAddress"`
}


func (val *DataBoxJobDetailsResponse) Defaults() *DataBoxJobDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ShippingAddress = *tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxSecretResponse struct {
	AccountCredentialDetails    []AccountCredentialDetailsResponse      `pulumi:"accountCredentialDetails"`
	DevicePassword              string                                  `pulumi:"devicePassword"`
	DeviceSerialNumber          string                                  `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey string                                  `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       []ApplianceNetworkConfigurationResponse `pulumi:"networkConfigurations"`
}

type DataboxJobSecretsResponse struct {
	JobSecretsType string                  `pulumi:"jobSecretsType"`
	PodSecrets     []DataBoxSecretResponse `pulumi:"podSecrets"`
}

type DestinationManagedDiskDetails struct {
	AccountId               *string `pulumi:"accountId"`
	DataDestinationType     string  `pulumi:"dataDestinationType"`
	ResourceGroupId         string  `pulumi:"resourceGroupId"`
	StagingStorageAccountId string  `pulumi:"stagingStorageAccountId"`
}

type DestinationManagedDiskDetailsResponse struct {
	AccountId               *string `pulumi:"accountId"`
	DataDestinationType     string  `pulumi:"dataDestinationType"`
	ResourceGroupId         string  `pulumi:"resourceGroupId"`
	StagingStorageAccountId string  `pulumi:"stagingStorageAccountId"`
}

type DestinationStorageAccountDetails struct {
	AccountId           *string `pulumi:"accountId"`
	DataDestinationType string  `pulumi:"dataDestinationType"`
	StorageAccountId    string  `pulumi:"storageAccountId"`
}

type DestinationStorageAccountDetailsResponse struct {
	AccountId           *string `pulumi:"accountId"`
	DataDestinationType string  `pulumi:"dataDestinationType"`
	StorageAccountId    string  `pulumi:"storageAccountId"`
}

type DiskSecretResponse struct {
	BitLockerKey     string `pulumi:"bitLockerKey"`
	DiskSerialNumber string `pulumi:"diskSerialNumber"`
}

type ErrorResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}

type ErrorResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponse)(nil)).Elem()
}

func (o ErrorResponseOutput) ToErrorResponseOutput() ErrorResponseOutput {
	return o
}

func (o ErrorResponseOutput) ToErrorResponseOutputWithContext(ctx context.Context) ErrorResponseOutput {
	return o
}

func (o ErrorResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponse) string { return v.Message }).(pulumi.StringOutput)
}

type JobErrorDetailsResponse struct {
	ErrorCode         int    `pulumi:"errorCode"`
	ErrorMessage      string `pulumi:"errorMessage"`
	ExceptionMessage  string `pulumi:"exceptionMessage"`
	RecommendedAction string `pulumi:"recommendedAction"`
}

type JobStagesResponse struct {
	DisplayName     string                    `pulumi:"displayName"`
	ErrorDetails    []JobErrorDetailsResponse `pulumi:"errorDetails"`
	JobStageDetails interface{}               `pulumi:"jobStageDetails"`
	StageName       string                    `pulumi:"stageName"`
	StageStatus     string                    `pulumi:"stageStatus"`
	StageTime       string                    `pulumi:"stageTime"`
}

type NotificationPreference struct {
	SendNotification bool   `pulumi:"sendNotification"`
	StageName        string `pulumi:"stageName"`
}


func (val *NotificationPreference) Defaults() *NotificationPreference {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SendNotification) {
		tmp.SendNotification = true
	}
	return &tmp
}

type NotificationPreferenceResponse struct {
	SendNotification bool   `pulumi:"sendNotification"`
	StageName        string `pulumi:"stageName"`
}


func (val *NotificationPreferenceResponse) Defaults() *NotificationPreferenceResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SendNotification) {
		tmp.SendNotification = true
	}
	return &tmp
}

type PackageShippingDetailsResponse struct {
	CarrierName string `pulumi:"carrierName"`
	TrackingId  string `pulumi:"trackingId"`
	TrackingUrl string `pulumi:"trackingUrl"`
}

type Preferences struct {
	PreferredDataCenterRegion []string `pulumi:"preferredDataCenterRegion"`
}

type PreferencesResponse struct {
	PreferredDataCenterRegion []string `pulumi:"preferredDataCenterRegion"`
}

type ShareCredentialDetailsResponse struct {
	Password                 string   `pulumi:"password"`
	ShareName                string   `pulumi:"shareName"`
	ShareType                string   `pulumi:"shareType"`
	SupportedAccessProtocols []string `pulumi:"supportedAccessProtocols"`
	UserName                 string   `pulumi:"userName"`
}

type ShippingAddress struct {
	AddressType     *string `pulumi:"addressType"`
	City            *string `pulumi:"city"`
	CompanyName     *string `pulumi:"companyName"`
	Country         string  `pulumi:"country"`
	PostalCode      string  `pulumi:"postalCode"`
	StateOrProvince *string `pulumi:"stateOrProvince"`
	StreetAddress1  string  `pulumi:"streetAddress1"`
	StreetAddress2  *string `pulumi:"streetAddress2"`
	StreetAddress3  *string `pulumi:"streetAddress3"`
	ZipExtendedCode *string `pulumi:"zipExtendedCode"`
}


func (val *ShippingAddress) Defaults() *ShippingAddress {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AddressType) {
		addressType_ := "None"
		tmp.AddressType = &addressType_
	}
	return &tmp
}

type ShippingAddressResponse struct {
	AddressType     *string `pulumi:"addressType"`
	City            *string `pulumi:"city"`
	CompanyName     *string `pulumi:"companyName"`
	Country         string  `pulumi:"country"`
	PostalCode      string  `pulumi:"postalCode"`
	StateOrProvince *string `pulumi:"stateOrProvince"`
	StreetAddress1  string  `pulumi:"streetAddress1"`
	StreetAddress2  *string `pulumi:"streetAddress2"`
	StreetAddress3  *string `pulumi:"streetAddress3"`
	ZipExtendedCode *string `pulumi:"zipExtendedCode"`
}


func (val *ShippingAddressResponse) Defaults() *ShippingAddressResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AddressType) {
		addressType_ := "None"
		tmp.AddressType = &addressType_
	}
	return &tmp
}

type Sku struct {
	DisplayName *string `pulumi:"displayName"`
	Family      *string `pulumi:"family"`
	Name        string  `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Family      pulumi.StringPtrInput `pulumi:"family"`
	Name        pulumi.StringInput    `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Family      *string `pulumi:"family"`
	Name        string  `pulumi:"name"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type UnencryptedCredentialsResponse struct {
	JobName    string      `pulumi:"jobName"`
	JobSecrets interface{} `pulumi:"jobSecrets"`
}

func init() {
	pulumi.RegisterOutputType(ErrorResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
}
