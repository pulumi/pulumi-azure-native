


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountCredentialDetailsResponse struct {
	AccountConnectionString string                           `pulumi:"accountConnectionString"`
	AccountName             string                           `pulumi:"accountName"`
	DataAccountType         string                           `pulumi:"dataAccountType"`
	ShareCredentialDetails  []ShareCredentialDetailsResponse `pulumi:"shareCredentialDetails"`
}

type AdditionalErrorInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type *string     `pulumi:"type"`
}

type AdditionalErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (AdditionalErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalErrorInfoResponse)(nil)).Elem()
}

func (o AdditionalErrorInfoResponseOutput) ToAdditionalErrorInfoResponseOutput() AdditionalErrorInfoResponseOutput {
	return o
}

func (o AdditionalErrorInfoResponseOutput) ToAdditionalErrorInfoResponseOutputWithContext(ctx context.Context) AdditionalErrorInfoResponseOutput {
	return o
}

func (o AdditionalErrorInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v AdditionalErrorInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o AdditionalErrorInfoResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalErrorInfoResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AdditionalErrorInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (AdditionalErrorInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalErrorInfoResponse)(nil)).Elem()
}

func (o AdditionalErrorInfoResponseArrayOutput) ToAdditionalErrorInfoResponseArrayOutput() AdditionalErrorInfoResponseArrayOutput {
	return o
}

func (o AdditionalErrorInfoResponseArrayOutput) ToAdditionalErrorInfoResponseArrayOutputWithContext(ctx context.Context) AdditionalErrorInfoResponseArrayOutput {
	return o
}

func (o AdditionalErrorInfoResponseArrayOutput) Index(i pulumi.IntInput) AdditionalErrorInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalErrorInfoResponse {
		return vs[0].([]AdditionalErrorInfoResponse)[vs[1].(int)]
	}).(AdditionalErrorInfoResponseOutput)
}

type ApplianceNetworkConfigurationResponse struct {
	MacAddress string `pulumi:"macAddress"`
	Name       string `pulumi:"name"`
}

type AzureFileFilterDetails struct {
	FilePathList   []string `pulumi:"filePathList"`
	FilePrefixList []string `pulumi:"filePrefixList"`
	FileShareList  []string `pulumi:"fileShareList"`
}

type AzureFileFilterDetailsResponse struct {
	FilePathList   []string `pulumi:"filePathList"`
	FilePrefixList []string `pulumi:"filePrefixList"`
	FileShareList  []string `pulumi:"fileShareList"`
}

type BlobFilterDetails struct {
	BlobPathList   []string `pulumi:"blobPathList"`
	BlobPrefixList []string `pulumi:"blobPrefixList"`
	ContainerList  []string `pulumi:"containerList"`
}

type BlobFilterDetailsResponse struct {
	BlobPathList   []string `pulumi:"blobPathList"`
	BlobPrefixList []string `pulumi:"blobPrefixList"`
	ContainerList  []string `pulumi:"containerList"`
}

type CloudErrorResponse struct {
	AdditionalInfo []AdditionalErrorInfoResponse `pulumi:"additionalInfo"`
	Code           *string                       `pulumi:"code"`
	Details        []CloudErrorResponse          `pulumi:"details"`
	Message        *string                       `pulumi:"message"`
	Target         *string                       `pulumi:"target"`
}

type CloudErrorResponseOutput struct{ *pulumi.OutputState }

func (CloudErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudErrorResponse)(nil)).Elem()
}

func (o CloudErrorResponseOutput) ToCloudErrorResponseOutput() CloudErrorResponseOutput {
	return o
}

func (o CloudErrorResponseOutput) ToCloudErrorResponseOutputWithContext(ctx context.Context) CloudErrorResponseOutput {
	return o
}

func (o CloudErrorResponseOutput) AdditionalInfo() AdditionalErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v CloudErrorResponse) []AdditionalErrorInfoResponse { return v.AdditionalInfo }).(AdditionalErrorInfoResponseArrayOutput)
}

func (o CloudErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o CloudErrorResponseOutput) Details() CloudErrorResponseArrayOutput {
	return o.ApplyT(func(v CloudErrorResponse) []CloudErrorResponse { return v.Details }).(CloudErrorResponseArrayOutput)
}

func (o CloudErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o CloudErrorResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type CloudErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (CloudErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudErrorResponse)(nil)).Elem()
}

func (o CloudErrorResponseArrayOutput) ToCloudErrorResponseArrayOutput() CloudErrorResponseArrayOutput {
	return o
}

func (o CloudErrorResponseArrayOutput) ToCloudErrorResponseArrayOutputWithContext(ctx context.Context) CloudErrorResponseArrayOutput {
	return o
}

func (o CloudErrorResponseArrayOutput) Index(i pulumi.IntInput) CloudErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudErrorResponse {
		return vs[0].([]CloudErrorResponse)[vs[1].(int)]
	}).(CloudErrorResponseOutput)
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
	AccountId                   string  `pulumi:"accountId"`
	BytesProcessed              float64 `pulumi:"bytesProcessed"`
	DataAccountType             string  `pulumi:"dataAccountType"`
	DirectoriesErroredOut       float64 `pulumi:"directoriesErroredOut"`
	FilesErroredOut             float64 `pulumi:"filesErroredOut"`
	FilesProcessed              float64 `pulumi:"filesProcessed"`
	InvalidDirectoriesProcessed float64 `pulumi:"invalidDirectoriesProcessed"`
	InvalidFileBytesUploaded    float64 `pulumi:"invalidFileBytesUploaded"`
	InvalidFilesProcessed       float64 `pulumi:"invalidFilesProcessed"`
	IsEnumerationInProgress     bool    `pulumi:"isEnumerationInProgress"`
	RenamedContainerCount       float64 `pulumi:"renamedContainerCount"`
	StorageAccountName          string  `pulumi:"storageAccountName"`
	TotalBytesToProcess         float64 `pulumi:"totalBytesToProcess"`
	TotalFilesToProcess         float64 `pulumi:"totalFilesToProcess"`
	TransferType                string  `pulumi:"transferType"`
}

type DataBoxAccountCopyLogDetailsResponse struct {
	AccountName        string `pulumi:"accountName"`
	CopyLogDetailsType string `pulumi:"copyLogDetailsType"`
	CopyLogLink        string `pulumi:"copyLogLink"`
	CopyVerboseLogLink string `pulumi:"copyVerboseLogLink"`
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
	ContactDetails              ContactDetails      `pulumi:"contactDetails"`
	DataExportDetails           []DataExportDetails `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetails `pulumi:"dataImportDetails"`
	ExpectedDataSizeInTeraBytes *int                `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string              `pulumi:"jobDetailsType"`
	KeyEncryptionKey            *KeyEncryptionKey   `pulumi:"keyEncryptionKey"`
	Passkey                     *string             `pulumi:"passkey"`
	Preferences                 *Preferences        `pulumi:"preferences"`
	PreferredDisks              map[string]int      `pulumi:"preferredDisks"`
	ShippingAddress             *ShippingAddress    `pulumi:"shippingAddress"`
}


func (val *DataBoxDiskJobDetails) Defaults() *DataBoxDiskJobDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.KeyEncryptionKey = tmp.KeyEncryptionKey.Defaults()

	tmp.Preferences = tmp.Preferences.Defaults()

	tmp.ShippingAddress = tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxDiskJobDetailsResponse struct {
	Actions                     []string                          `pulumi:"actions"`
	ChainOfCustodySasKey        string                            `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse            `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                     `pulumi:"copyLogDetails"`
	CopyProgress                []DataBoxDiskCopyProgressResponse `pulumi:"copyProgress"`
	DataExportDetails           []DataExportDetailsResponse       `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetailsResponse       `pulumi:"dataImportDetails"`
	DeliveryPackage             PackageShippingDetailsResponse    `pulumi:"deliveryPackage"`
	DisksAndSizeDetails         map[string]int                    `pulumi:"disksAndSizeDetails"`
	ExpectedDataSizeInTeraBytes *int                              `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string                            `pulumi:"jobDetailsType"`
	JobStages                   []JobStagesResponse               `pulumi:"jobStages"`
	KeyEncryptionKey            *KeyEncryptionKeyResponse         `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob   LastMitigationActionOnJobResponse `pulumi:"lastMitigationActionOnJob"`
	Passkey                     *string                           `pulumi:"passkey"`
	Preferences                 *PreferencesResponse              `pulumi:"preferences"`
	PreferredDisks              map[string]int                    `pulumi:"preferredDisks"`
	ReturnPackage               PackageShippingDetailsResponse    `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  string                            `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             *ShippingAddressResponse          `pulumi:"shippingAddress"`
}


func (val *DataBoxDiskJobDetailsResponse) Defaults() *DataBoxDiskJobDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.KeyEncryptionKey = tmp.KeyEncryptionKey.Defaults()

	tmp.Preferences = tmp.Preferences.Defaults()

	tmp.ShippingAddress = tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxDiskJobSecretsResponse struct {
	DcAccessSecurityCode DcAccessSecurityCodeResponse `pulumi:"dcAccessSecurityCode"`
	DiskSecrets          []DiskSecretResponse         `pulumi:"diskSecrets"`
	Error                CloudErrorResponse           `pulumi:"error"`
	IsPasskeyUserDefined bool                         `pulumi:"isPasskeyUserDefined"`
	JobSecretsType       string                       `pulumi:"jobSecretsType"`
	PassKey              string                       `pulumi:"passKey"`
}

type DataBoxHeavyAccountCopyLogDetailsResponse struct {
	AccountName        string   `pulumi:"accountName"`
	CopyLogDetailsType string   `pulumi:"copyLogDetailsType"`
	CopyLogLink        []string `pulumi:"copyLogLink"`
	CopyVerboseLogLink []string `pulumi:"copyVerboseLogLink"`
}

type DataBoxHeavyJobDetails struct {
	ContactDetails              ContactDetails      `pulumi:"contactDetails"`
	DataExportDetails           []DataExportDetails `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetails `pulumi:"dataImportDetails"`
	DevicePassword              *string             `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes *int                `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string              `pulumi:"jobDetailsType"`
	KeyEncryptionKey            *KeyEncryptionKey   `pulumi:"keyEncryptionKey"`
	Preferences                 *Preferences        `pulumi:"preferences"`
	ShippingAddress             *ShippingAddress    `pulumi:"shippingAddress"`
}


func (val *DataBoxHeavyJobDetails) Defaults() *DataBoxHeavyJobDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.KeyEncryptionKey = tmp.KeyEncryptionKey.Defaults()

	tmp.Preferences = tmp.Preferences.Defaults()

	tmp.ShippingAddress = tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxHeavyJobDetailsResponse struct {
	Actions                     []string                          `pulumi:"actions"`
	ChainOfCustodySasKey        string                            `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse            `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                     `pulumi:"copyLogDetails"`
	CopyProgress                []CopyProgressResponse            `pulumi:"copyProgress"`
	DataExportDetails           []DataExportDetailsResponse       `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetailsResponse       `pulumi:"dataImportDetails"`
	DeliveryPackage             PackageShippingDetailsResponse    `pulumi:"deliveryPackage"`
	DevicePassword              *string                           `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes *int                              `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string                            `pulumi:"jobDetailsType"`
	JobStages                   []JobStagesResponse               `pulumi:"jobStages"`
	KeyEncryptionKey            *KeyEncryptionKeyResponse         `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob   LastMitigationActionOnJobResponse `pulumi:"lastMitigationActionOnJob"`
	Preferences                 *PreferencesResponse              `pulumi:"preferences"`
	ReturnPackage               PackageShippingDetailsResponse    `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  string                            `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             *ShippingAddressResponse          `pulumi:"shippingAddress"`
}


func (val *DataBoxHeavyJobDetailsResponse) Defaults() *DataBoxHeavyJobDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.KeyEncryptionKey = tmp.KeyEncryptionKey.Defaults()

	tmp.Preferences = tmp.Preferences.Defaults()

	tmp.ShippingAddress = tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxHeavyJobSecretsResponse struct {
	CabinetPodSecrets    []DataBoxHeavySecretResponse `pulumi:"cabinetPodSecrets"`
	DcAccessSecurityCode DcAccessSecurityCodeResponse `pulumi:"dcAccessSecurityCode"`
	Error                CloudErrorResponse           `pulumi:"error"`
	JobSecretsType       string                       `pulumi:"jobSecretsType"`
}

type DataBoxHeavySecretResponse struct {
	AccountCredentialDetails    []AccountCredentialDetailsResponse      `pulumi:"accountCredentialDetails"`
	DevicePassword              string                                  `pulumi:"devicePassword"`
	DeviceSerialNumber          string                                  `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey string                                  `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       []ApplianceNetworkConfigurationResponse `pulumi:"networkConfigurations"`
}

type DataBoxJobDetails struct {
	ContactDetails              ContactDetails      `pulumi:"contactDetails"`
	DataExportDetails           []DataExportDetails `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetails `pulumi:"dataImportDetails"`
	DevicePassword              *string             `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes *int                `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string              `pulumi:"jobDetailsType"`
	KeyEncryptionKey            *KeyEncryptionKey   `pulumi:"keyEncryptionKey"`
	Preferences                 *Preferences        `pulumi:"preferences"`
	ShippingAddress             *ShippingAddress    `pulumi:"shippingAddress"`
}


func (val *DataBoxJobDetails) Defaults() *DataBoxJobDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.KeyEncryptionKey = tmp.KeyEncryptionKey.Defaults()

	tmp.Preferences = tmp.Preferences.Defaults()

	tmp.ShippingAddress = tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxJobDetailsResponse struct {
	Actions                     []string                          `pulumi:"actions"`
	ChainOfCustodySasKey        string                            `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse            `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                     `pulumi:"copyLogDetails"`
	CopyProgress                []CopyProgressResponse            `pulumi:"copyProgress"`
	DataExportDetails           []DataExportDetailsResponse       `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetailsResponse       `pulumi:"dataImportDetails"`
	DeliveryPackage             PackageShippingDetailsResponse    `pulumi:"deliveryPackage"`
	DevicePassword              *string                           `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes *int                              `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string                            `pulumi:"jobDetailsType"`
	JobStages                   []JobStagesResponse               `pulumi:"jobStages"`
	KeyEncryptionKey            *KeyEncryptionKeyResponse         `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob   LastMitigationActionOnJobResponse `pulumi:"lastMitigationActionOnJob"`
	Preferences                 *PreferencesResponse              `pulumi:"preferences"`
	ReturnPackage               PackageShippingDetailsResponse    `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  string                            `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             *ShippingAddressResponse          `pulumi:"shippingAddress"`
}


func (val *DataBoxJobDetailsResponse) Defaults() *DataBoxJobDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.KeyEncryptionKey = tmp.KeyEncryptionKey.Defaults()

	tmp.Preferences = tmp.Preferences.Defaults()

	tmp.ShippingAddress = tmp.ShippingAddress.Defaults()

	return &tmp
}

type DataBoxSecretResponse struct {
	AccountCredentialDetails    []AccountCredentialDetailsResponse      `pulumi:"accountCredentialDetails"`
	DevicePassword              string                                  `pulumi:"devicePassword"`
	DeviceSerialNumber          string                                  `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey string                                  `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       []ApplianceNetworkConfigurationResponse `pulumi:"networkConfigurations"`
}

type DataExportDetails struct {
	AccountDetails        interface{}           `pulumi:"accountDetails"`
	LogCollectionLevel    *string               `pulumi:"logCollectionLevel"`
	TransferConfiguration TransferConfiguration `pulumi:"transferConfiguration"`
}


func (val *DataExportDetails) Defaults() *DataExportDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LogCollectionLevel) {
		logCollectionLevel_ := "Error"
		tmp.LogCollectionLevel = &logCollectionLevel_
	}
	tmp.TransferConfiguration = *tmp.TransferConfiguration.Defaults()

	return &tmp
}

type DataExportDetailsResponse struct {
	AccountDetails        interface{}                   `pulumi:"accountDetails"`
	LogCollectionLevel    *string                       `pulumi:"logCollectionLevel"`
	TransferConfiguration TransferConfigurationResponse `pulumi:"transferConfiguration"`
}


func (val *DataExportDetailsResponse) Defaults() *DataExportDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LogCollectionLevel) {
		logCollectionLevel_ := "Error"
		tmp.LogCollectionLevel = &logCollectionLevel_
	}
	tmp.TransferConfiguration = *tmp.TransferConfiguration.Defaults()

	return &tmp
}

type DataImportDetails struct {
	AccountDetails interface{} `pulumi:"accountDetails"`
}

type DataImportDetailsResponse struct {
	AccountDetails interface{} `pulumi:"accountDetails"`
}

type DataboxJobSecretsResponse struct {
	DcAccessSecurityCode DcAccessSecurityCodeResponse `pulumi:"dcAccessSecurityCode"`
	Error                CloudErrorResponse           `pulumi:"error"`
	JobSecretsType       string                       `pulumi:"jobSecretsType"`
	PodSecrets           []DataBoxSecretResponse      `pulumi:"podSecrets"`
}

type DcAccessSecurityCodeResponse struct {
	ForwardDCAccessCode *string `pulumi:"forwardDCAccessCode"`
	ReverseDCAccessCode *string `pulumi:"reverseDCAccessCode"`
}

type DiskSecretResponse struct {
	BitLockerKey     string `pulumi:"bitLockerKey"`
	DiskSerialNumber string `pulumi:"diskSerialNumber"`
}

type EncryptionPreferences struct {
	DoubleEncryption *string `pulumi:"doubleEncryption"`
}


func (val *EncryptionPreferences) Defaults() *EncryptionPreferences {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DoubleEncryption) {
		doubleEncryption_ := "Disabled"
		tmp.DoubleEncryption = &doubleEncryption_
	}
	return &tmp
}

type EncryptionPreferencesResponse struct {
	DoubleEncryption *string `pulumi:"doubleEncryption"`
}


func (val *EncryptionPreferencesResponse) Defaults() *EncryptionPreferencesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DoubleEncryption) {
		doubleEncryption_ := "Disabled"
		tmp.DoubleEncryption = &doubleEncryption_
	}
	return &tmp
}

type FilterFileDetails struct {
	FilterFilePath string `pulumi:"filterFilePath"`
	FilterFileType string `pulumi:"filterFileType"`
}

type FilterFileDetailsResponse struct {
	FilterFilePath string `pulumi:"filterFilePath"`
	FilterFileType string `pulumi:"filterFileType"`
}

type IdentityProperties struct {
	Type         *string                 `pulumi:"type"`
	UserAssigned *UserAssignedProperties `pulumi:"userAssigned"`
}

type IdentityPropertiesResponse struct {
	Type         *string                         `pulumi:"type"`
	UserAssigned *UserAssignedPropertiesResponse `pulumi:"userAssigned"`
}

type JobDeliveryInfo struct {
	ScheduledDateTime *string `pulumi:"scheduledDateTime"`
}





type JobDeliveryInfoInput interface {
	pulumi.Input

	ToJobDeliveryInfoOutput() JobDeliveryInfoOutput
	ToJobDeliveryInfoOutputWithContext(context.Context) JobDeliveryInfoOutput
}

type JobDeliveryInfoArgs struct {
	ScheduledDateTime pulumi.StringPtrInput `pulumi:"scheduledDateTime"`
}

func (JobDeliveryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDeliveryInfo)(nil)).Elem()
}

func (i JobDeliveryInfoArgs) ToJobDeliveryInfoOutput() JobDeliveryInfoOutput {
	return i.ToJobDeliveryInfoOutputWithContext(context.Background())
}

func (i JobDeliveryInfoArgs) ToJobDeliveryInfoOutputWithContext(ctx context.Context) JobDeliveryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoOutput)
}

func (i JobDeliveryInfoArgs) ToJobDeliveryInfoPtrOutput() JobDeliveryInfoPtrOutput {
	return i.ToJobDeliveryInfoPtrOutputWithContext(context.Background())
}

func (i JobDeliveryInfoArgs) ToJobDeliveryInfoPtrOutputWithContext(ctx context.Context) JobDeliveryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoOutput).ToJobDeliveryInfoPtrOutputWithContext(ctx)
}









type JobDeliveryInfoPtrInput interface {
	pulumi.Input

	ToJobDeliveryInfoPtrOutput() JobDeliveryInfoPtrOutput
	ToJobDeliveryInfoPtrOutputWithContext(context.Context) JobDeliveryInfoPtrOutput
}

type jobDeliveryInfoPtrType JobDeliveryInfoArgs

func JobDeliveryInfoPtr(v *JobDeliveryInfoArgs) JobDeliveryInfoPtrInput {
	return (*jobDeliveryInfoPtrType)(v)
}

func (*jobDeliveryInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDeliveryInfo)(nil)).Elem()
}

func (i *jobDeliveryInfoPtrType) ToJobDeliveryInfoPtrOutput() JobDeliveryInfoPtrOutput {
	return i.ToJobDeliveryInfoPtrOutputWithContext(context.Background())
}

func (i *jobDeliveryInfoPtrType) ToJobDeliveryInfoPtrOutputWithContext(ctx context.Context) JobDeliveryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoPtrOutput)
}

type JobDeliveryInfoOutput struct{ *pulumi.OutputState }

func (JobDeliveryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDeliveryInfo)(nil)).Elem()
}

func (o JobDeliveryInfoOutput) ToJobDeliveryInfoOutput() JobDeliveryInfoOutput {
	return o
}

func (o JobDeliveryInfoOutput) ToJobDeliveryInfoOutputWithContext(ctx context.Context) JobDeliveryInfoOutput {
	return o
}

func (o JobDeliveryInfoOutput) ToJobDeliveryInfoPtrOutput() JobDeliveryInfoPtrOutput {
	return o.ToJobDeliveryInfoPtrOutputWithContext(context.Background())
}

func (o JobDeliveryInfoOutput) ToJobDeliveryInfoPtrOutputWithContext(ctx context.Context) JobDeliveryInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobDeliveryInfo) *JobDeliveryInfo {
		return &v
	}).(JobDeliveryInfoPtrOutput)
}

func (o JobDeliveryInfoOutput) ScheduledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDeliveryInfo) *string { return v.ScheduledDateTime }).(pulumi.StringPtrOutput)
}

type JobDeliveryInfoPtrOutput struct{ *pulumi.OutputState }

func (JobDeliveryInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDeliveryInfo)(nil)).Elem()
}

func (o JobDeliveryInfoPtrOutput) ToJobDeliveryInfoPtrOutput() JobDeliveryInfoPtrOutput {
	return o
}

func (o JobDeliveryInfoPtrOutput) ToJobDeliveryInfoPtrOutputWithContext(ctx context.Context) JobDeliveryInfoPtrOutput {
	return o
}

func (o JobDeliveryInfoPtrOutput) Elem() JobDeliveryInfoOutput {
	return o.ApplyT(func(v *JobDeliveryInfo) JobDeliveryInfo {
		if v != nil {
			return *v
		}
		var ret JobDeliveryInfo
		return ret
	}).(JobDeliveryInfoOutput)
}

func (o JobDeliveryInfoPtrOutput) ScheduledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDeliveryInfo) *string {
		if v == nil {
			return nil
		}
		return v.ScheduledDateTime
	}).(pulumi.StringPtrOutput)
}

type JobDeliveryInfoResponse struct {
	ScheduledDateTime *string `pulumi:"scheduledDateTime"`
}

type JobDeliveryInfoResponseOutput struct{ *pulumi.OutputState }

func (JobDeliveryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDeliveryInfoResponse)(nil)).Elem()
}

func (o JobDeliveryInfoResponseOutput) ToJobDeliveryInfoResponseOutput() JobDeliveryInfoResponseOutput {
	return o
}

func (o JobDeliveryInfoResponseOutput) ToJobDeliveryInfoResponseOutputWithContext(ctx context.Context) JobDeliveryInfoResponseOutput {
	return o
}

func (o JobDeliveryInfoResponseOutput) ScheduledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDeliveryInfoResponse) *string { return v.ScheduledDateTime }).(pulumi.StringPtrOutput)
}

type JobDeliveryInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (JobDeliveryInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDeliveryInfoResponse)(nil)).Elem()
}

func (o JobDeliveryInfoResponsePtrOutput) ToJobDeliveryInfoResponsePtrOutput() JobDeliveryInfoResponsePtrOutput {
	return o
}

func (o JobDeliveryInfoResponsePtrOutput) ToJobDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) JobDeliveryInfoResponsePtrOutput {
	return o
}

func (o JobDeliveryInfoResponsePtrOutput) Elem() JobDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *JobDeliveryInfoResponse) JobDeliveryInfoResponse {
		if v != nil {
			return *v
		}
		var ret JobDeliveryInfoResponse
		return ret
	}).(JobDeliveryInfoResponseOutput)
}

func (o JobDeliveryInfoResponsePtrOutput) ScheduledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDeliveryInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScheduledDateTime
	}).(pulumi.StringPtrOutput)
}

type JobStagesResponse struct {
	DisplayName     string      `pulumi:"displayName"`
	JobStageDetails interface{} `pulumi:"jobStageDetails"`
	StageName       string      `pulumi:"stageName"`
	StageStatus     string      `pulumi:"stageStatus"`
	StageTime       string      `pulumi:"stageTime"`
}

type KeyEncryptionKey struct {
	IdentityProperties *IdentityProperties `pulumi:"identityProperties"`
	KekType            string              `pulumi:"kekType"`
	KekUrl             *string             `pulumi:"kekUrl"`
	KekVaultResourceID *string             `pulumi:"kekVaultResourceID"`
}


func (val *KeyEncryptionKey) Defaults() *KeyEncryptionKey {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KekType) {
		tmp.KekType = "MicrosoftManaged"
	}
	return &tmp
}

type KeyEncryptionKeyResponse struct {
	IdentityProperties *IdentityPropertiesResponse `pulumi:"identityProperties"`
	KekType            string                      `pulumi:"kekType"`
	KekUrl             *string                     `pulumi:"kekUrl"`
	KekVaultResourceID *string                     `pulumi:"kekVaultResourceID"`
}


func (val *KeyEncryptionKeyResponse) Defaults() *KeyEncryptionKeyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KekType) {
		tmp.KekType = "MicrosoftManaged"
	}
	return &tmp
}

type LastMitigationActionOnJobResponse struct {
	ActionDateTimeInUtc   *string `pulumi:"actionDateTimeInUtc"`
	CustomerResolution    *string `pulumi:"customerResolution"`
	IsPerformedByCustomer *bool   `pulumi:"isPerformedByCustomer"`
}

type ManagedDiskDetails struct {
	DataAccountType         string  `pulumi:"dataAccountType"`
	ResourceGroupId         string  `pulumi:"resourceGroupId"`
	SharePassword           *string `pulumi:"sharePassword"`
	StagingStorageAccountId string  `pulumi:"stagingStorageAccountId"`
}


func (val *ManagedDiskDetails) Defaults() *ManagedDiskDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataAccountType) {
		tmp.DataAccountType = "StorageAccount"
	}
	return &tmp
}

type ManagedDiskDetailsResponse struct {
	DataAccountType         string `pulumi:"dataAccountType"`
	ResourceGroupId         string `pulumi:"resourceGroupId"`
	StagingStorageAccountId string `pulumi:"stagingStorageAccountId"`
}


func (val *ManagedDiskDetailsResponse) Defaults() *ManagedDiskDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataAccountType) {
		tmp.DataAccountType = "StorageAccount"
	}
	return &tmp
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
	EncryptionPreferences     *EncryptionPreferences `pulumi:"encryptionPreferences"`
	PreferredDataCenterRegion []string               `pulumi:"preferredDataCenterRegion"`
	TransportPreferences      *TransportPreferences  `pulumi:"transportPreferences"`
}


func (val *Preferences) Defaults() *Preferences {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EncryptionPreferences = tmp.EncryptionPreferences.Defaults()

	return &tmp
}

type PreferencesResponse struct {
	EncryptionPreferences     *EncryptionPreferencesResponse `pulumi:"encryptionPreferences"`
	PreferredDataCenterRegion []string                       `pulumi:"preferredDataCenterRegion"`
	TransportPreferences      *TransportPreferencesResponse  `pulumi:"transportPreferences"`
}


func (val *PreferencesResponse) Defaults() *PreferencesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EncryptionPreferences = tmp.EncryptionPreferences.Defaults()

	return &tmp
}

type ResourceIdentity struct {
	Type                   *string                `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}


func (val *ResourceIdentity) Defaults() *ResourceIdentity {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "None"
		tmp.Type = &type_
	}
	return &tmp
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput       `pulumi:"userAssignedIdentities"`
}


func (val *ResourceIdentityArgs) Defaults() *ResourceIdentityArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		tmp.Type = pulumi.StringPtr("None")
	}
	return &tmp
}
func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ResourceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ResourceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ResourceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}


func (val *ResourceIdentityResponse) Defaults() *ResourceIdentityResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "None"
		tmp.Type = &type_
	}
	return &tmp
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
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
	PostalCode      *string `pulumi:"postalCode"`
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
	PostalCode      *string `pulumi:"postalCode"`
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

type StorageAccountDetails struct {
	DataAccountType  string  `pulumi:"dataAccountType"`
	SharePassword    *string `pulumi:"sharePassword"`
	StorageAccountId string  `pulumi:"storageAccountId"`
}


func (val *StorageAccountDetails) Defaults() *StorageAccountDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataAccountType) {
		tmp.DataAccountType = "StorageAccount"
	}
	return &tmp
}

type StorageAccountDetailsResponse struct {
	DataAccountType  string `pulumi:"dataAccountType"`
	StorageAccountId string `pulumi:"storageAccountId"`
}


func (val *StorageAccountDetailsResponse) Defaults() *StorageAccountDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataAccountType) {
		tmp.DataAccountType = "StorageAccount"
	}
	return &tmp
}

type SystemDataResponse struct {
	CreatedAt          string `pulumi:"createdAt"`
	CreatedBy          string `pulumi:"createdBy"`
	CreatedByType      string `pulumi:"createdByType"`
	LastModifiedAt     string `pulumi:"lastModifiedAt"`
	LastModifiedBy     string `pulumi:"lastModifiedBy"`
	LastModifiedByType string `pulumi:"lastModifiedByType"`
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

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedByType }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedByType }).(pulumi.StringOutput)
}

type TransferAllDetails struct {
	DataAccountType  string `pulumi:"dataAccountType"`
	TransferAllBlobs *bool  `pulumi:"transferAllBlobs"`
	TransferAllFiles *bool  `pulumi:"transferAllFiles"`
}


func (val *TransferAllDetails) Defaults() *TransferAllDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataAccountType) {
		tmp.DataAccountType = "StorageAccount"
	}
	return &tmp
}

type TransferAllDetailsResponse struct {
	DataAccountType  string `pulumi:"dataAccountType"`
	TransferAllBlobs *bool  `pulumi:"transferAllBlobs"`
	TransferAllFiles *bool  `pulumi:"transferAllFiles"`
}


func (val *TransferAllDetailsResponse) Defaults() *TransferAllDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataAccountType) {
		tmp.DataAccountType = "StorageAccount"
	}
	return &tmp
}

type TransferConfiguration struct {
	TransferAllDetails        *TransferConfigurationTransferAllDetails    `pulumi:"transferAllDetails"`
	TransferConfigurationType string                                      `pulumi:"transferConfigurationType"`
	TransferFilterDetails     *TransferConfigurationTransferFilterDetails `pulumi:"transferFilterDetails"`
}


func (val *TransferConfiguration) Defaults() *TransferConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.TransferAllDetails = tmp.TransferAllDetails.Defaults()

	tmp.TransferFilterDetails = tmp.TransferFilterDetails.Defaults()

	return &tmp
}

type TransferConfigurationResponse struct {
	TransferAllDetails        *TransferConfigurationResponseTransferAllDetails    `pulumi:"transferAllDetails"`
	TransferConfigurationType string                                              `pulumi:"transferConfigurationType"`
	TransferFilterDetails     *TransferConfigurationResponseTransferFilterDetails `pulumi:"transferFilterDetails"`
}


func (val *TransferConfigurationResponse) Defaults() *TransferConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.TransferAllDetails = tmp.TransferAllDetails.Defaults()

	tmp.TransferFilterDetails = tmp.TransferFilterDetails.Defaults()

	return &tmp
}

type TransferConfigurationResponseTransferAllDetails struct {
	Include *TransferAllDetailsResponse `pulumi:"include"`
}


func (val *TransferConfigurationResponseTransferAllDetails) Defaults() *TransferConfigurationResponseTransferAllDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Include = tmp.Include.Defaults()

	return &tmp
}

type TransferConfigurationResponseTransferFilterDetails struct {
	Include *TransferFilterDetailsResponse `pulumi:"include"`
}


func (val *TransferConfigurationResponseTransferFilterDetails) Defaults() *TransferConfigurationResponseTransferFilterDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Include = tmp.Include.Defaults()

	return &tmp
}

type TransferConfigurationTransferAllDetails struct {
	Include *TransferAllDetails `pulumi:"include"`
}


func (val *TransferConfigurationTransferAllDetails) Defaults() *TransferConfigurationTransferAllDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Include = tmp.Include.Defaults()

	return &tmp
}

type TransferConfigurationTransferFilterDetails struct {
	Include *TransferFilterDetails `pulumi:"include"`
}


func (val *TransferConfigurationTransferFilterDetails) Defaults() *TransferConfigurationTransferFilterDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Include = tmp.Include.Defaults()

	return &tmp
}

type TransferFilterDetails struct {
	AzureFileFilterDetails *AzureFileFilterDetails `pulumi:"azureFileFilterDetails"`
	BlobFilterDetails      *BlobFilterDetails      `pulumi:"blobFilterDetails"`
	DataAccountType        string                  `pulumi:"dataAccountType"`
	FilterFileDetails      []FilterFileDetails     `pulumi:"filterFileDetails"`
}


func (val *TransferFilterDetails) Defaults() *TransferFilterDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataAccountType) {
		tmp.DataAccountType = "StorageAccount"
	}
	return &tmp
}

type TransferFilterDetailsResponse struct {
	AzureFileFilterDetails *AzureFileFilterDetailsResponse `pulumi:"azureFileFilterDetails"`
	BlobFilterDetails      *BlobFilterDetailsResponse      `pulumi:"blobFilterDetails"`
	DataAccountType        string                          `pulumi:"dataAccountType"`
	FilterFileDetails      []FilterFileDetailsResponse     `pulumi:"filterFileDetails"`
}


func (val *TransferFilterDetailsResponse) Defaults() *TransferFilterDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataAccountType) {
		tmp.DataAccountType = "StorageAccount"
	}
	return &tmp
}

type TransportPreferences struct {
	PreferredShipmentType string `pulumi:"preferredShipmentType"`
}

type TransportPreferencesResponse struct {
	PreferredShipmentType string `pulumi:"preferredShipmentType"`
}

type UnencryptedCredentialsResponse struct {
	JobName    string      `pulumi:"jobName"`
	JobSecrets interface{} `pulumi:"jobSecrets"`
}

type UnencryptedCredentialsResponseOutput struct{ *pulumi.OutputState }

func (UnencryptedCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnencryptedCredentialsResponse)(nil)).Elem()
}

func (o UnencryptedCredentialsResponseOutput) ToUnencryptedCredentialsResponseOutput() UnencryptedCredentialsResponseOutput {
	return o
}

func (o UnencryptedCredentialsResponseOutput) ToUnencryptedCredentialsResponseOutputWithContext(ctx context.Context) UnencryptedCredentialsResponseOutput {
	return o
}

func (o UnencryptedCredentialsResponseOutput) JobName() pulumi.StringOutput {
	return o.ApplyT(func(v UnencryptedCredentialsResponse) string { return v.JobName }).(pulumi.StringOutput)
}

func (o UnencryptedCredentialsResponseOutput) JobSecrets() pulumi.AnyOutput {
	return o.ApplyT(func(v UnencryptedCredentialsResponse) interface{} { return v.JobSecrets }).(pulumi.AnyOutput)
}

type UnencryptedCredentialsResponseArrayOutput struct{ *pulumi.OutputState }

func (UnencryptedCredentialsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UnencryptedCredentialsResponse)(nil)).Elem()
}

func (o UnencryptedCredentialsResponseArrayOutput) ToUnencryptedCredentialsResponseArrayOutput() UnencryptedCredentialsResponseArrayOutput {
	return o
}

func (o UnencryptedCredentialsResponseArrayOutput) ToUnencryptedCredentialsResponseArrayOutputWithContext(ctx context.Context) UnencryptedCredentialsResponseArrayOutput {
	return o
}

func (o UnencryptedCredentialsResponseArrayOutput) Index(i pulumi.IntInput) UnencryptedCredentialsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UnencryptedCredentialsResponse {
		return vs[0].([]UnencryptedCredentialsResponse)[vs[1].(int)]
	}).(UnencryptedCredentialsResponseOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type UserAssignedProperties struct {
	ResourceId *string `pulumi:"resourceId"`
}

type UserAssignedPropertiesResponse struct {
	ResourceId *string `pulumi:"resourceId"`
}

func init() {
	pulumi.RegisterOutputType(AdditionalErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(AdditionalErrorInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(CloudErrorResponseOutput{})
	pulumi.RegisterOutputType(CloudErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoPtrOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoResponseOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UnencryptedCredentialsResponseOutput{})
	pulumi.RegisterOutputType(UnencryptedCredentialsResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
