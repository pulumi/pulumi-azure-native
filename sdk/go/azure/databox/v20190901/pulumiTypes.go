


package v20190901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountCredentialDetailsResponse struct {
	AccountConnectionString string                           `pulumi:"accountConnectionString"`
	AccountName             string                           `pulumi:"accountName"`
	DataDestinationType     string                           `pulumi:"dataDestinationType"`
	ShareCredentialDetails  []ShareCredentialDetailsResponse `pulumi:"shareCredentialDetails"`
}





type AccountCredentialDetailsResponseInput interface {
	pulumi.Input

	ToAccountCredentialDetailsResponseOutput() AccountCredentialDetailsResponseOutput
	ToAccountCredentialDetailsResponseOutputWithContext(context.Context) AccountCredentialDetailsResponseOutput
}

type AccountCredentialDetailsResponseArgs struct {
	AccountConnectionString pulumi.StringInput                       `pulumi:"accountConnectionString"`
	AccountName             pulumi.StringInput                       `pulumi:"accountName"`
	DataDestinationType     pulumi.StringInput                       `pulumi:"dataDestinationType"`
	ShareCredentialDetails  ShareCredentialDetailsResponseArrayInput `pulumi:"shareCredentialDetails"`
}

func (AccountCredentialDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountCredentialDetailsResponse)(nil)).Elem()
}

func (i AccountCredentialDetailsResponseArgs) ToAccountCredentialDetailsResponseOutput() AccountCredentialDetailsResponseOutput {
	return i.ToAccountCredentialDetailsResponseOutputWithContext(context.Background())
}

func (i AccountCredentialDetailsResponseArgs) ToAccountCredentialDetailsResponseOutputWithContext(ctx context.Context) AccountCredentialDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountCredentialDetailsResponseOutput)
}





type AccountCredentialDetailsResponseArrayInput interface {
	pulumi.Input

	ToAccountCredentialDetailsResponseArrayOutput() AccountCredentialDetailsResponseArrayOutput
	ToAccountCredentialDetailsResponseArrayOutputWithContext(context.Context) AccountCredentialDetailsResponseArrayOutput
}

type AccountCredentialDetailsResponseArray []AccountCredentialDetailsResponseInput

func (AccountCredentialDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountCredentialDetailsResponse)(nil)).Elem()
}

func (i AccountCredentialDetailsResponseArray) ToAccountCredentialDetailsResponseArrayOutput() AccountCredentialDetailsResponseArrayOutput {
	return i.ToAccountCredentialDetailsResponseArrayOutputWithContext(context.Background())
}

func (i AccountCredentialDetailsResponseArray) ToAccountCredentialDetailsResponseArrayOutputWithContext(ctx context.Context) AccountCredentialDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountCredentialDetailsResponseArrayOutput)
}

type AccountCredentialDetailsResponseOutput struct{ *pulumi.OutputState }

func (AccountCredentialDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountCredentialDetailsResponse)(nil)).Elem()
}

func (o AccountCredentialDetailsResponseOutput) ToAccountCredentialDetailsResponseOutput() AccountCredentialDetailsResponseOutput {
	return o
}

func (o AccountCredentialDetailsResponseOutput) ToAccountCredentialDetailsResponseOutputWithContext(ctx context.Context) AccountCredentialDetailsResponseOutput {
	return o
}

func (o AccountCredentialDetailsResponseOutput) AccountConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v AccountCredentialDetailsResponse) string { return v.AccountConnectionString }).(pulumi.StringOutput)
}

func (o AccountCredentialDetailsResponseOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v AccountCredentialDetailsResponse) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o AccountCredentialDetailsResponseOutput) DataDestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v AccountCredentialDetailsResponse) string { return v.DataDestinationType }).(pulumi.StringOutput)
}

func (o AccountCredentialDetailsResponseOutput) ShareCredentialDetails() ShareCredentialDetailsResponseArrayOutput {
	return o.ApplyT(func(v AccountCredentialDetailsResponse) []ShareCredentialDetailsResponse {
		return v.ShareCredentialDetails
	}).(ShareCredentialDetailsResponseArrayOutput)
}

type AccountCredentialDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (AccountCredentialDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountCredentialDetailsResponse)(nil)).Elem()
}

func (o AccountCredentialDetailsResponseArrayOutput) ToAccountCredentialDetailsResponseArrayOutput() AccountCredentialDetailsResponseArrayOutput {
	return o
}

func (o AccountCredentialDetailsResponseArrayOutput) ToAccountCredentialDetailsResponseArrayOutputWithContext(ctx context.Context) AccountCredentialDetailsResponseArrayOutput {
	return o
}

func (o AccountCredentialDetailsResponseArrayOutput) Index(i pulumi.IntInput) AccountCredentialDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccountCredentialDetailsResponse {
		return vs[0].([]AccountCredentialDetailsResponse)[vs[1].(int)]
	}).(AccountCredentialDetailsResponseOutput)
}

type ApplianceNetworkConfigurationResponse struct {
	MacAddress string `pulumi:"macAddress"`
	Name       string `pulumi:"name"`
}





type ApplianceNetworkConfigurationResponseInput interface {
	pulumi.Input

	ToApplianceNetworkConfigurationResponseOutput() ApplianceNetworkConfigurationResponseOutput
	ToApplianceNetworkConfigurationResponseOutputWithContext(context.Context) ApplianceNetworkConfigurationResponseOutput
}

type ApplianceNetworkConfigurationResponseArgs struct {
	MacAddress pulumi.StringInput `pulumi:"macAddress"`
	Name       pulumi.StringInput `pulumi:"name"`
}

func (ApplianceNetworkConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceNetworkConfigurationResponse)(nil)).Elem()
}

func (i ApplianceNetworkConfigurationResponseArgs) ToApplianceNetworkConfigurationResponseOutput() ApplianceNetworkConfigurationResponseOutput {
	return i.ToApplianceNetworkConfigurationResponseOutputWithContext(context.Background())
}

func (i ApplianceNetworkConfigurationResponseArgs) ToApplianceNetworkConfigurationResponseOutputWithContext(ctx context.Context) ApplianceNetworkConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceNetworkConfigurationResponseOutput)
}





type ApplianceNetworkConfigurationResponseArrayInput interface {
	pulumi.Input

	ToApplianceNetworkConfigurationResponseArrayOutput() ApplianceNetworkConfigurationResponseArrayOutput
	ToApplianceNetworkConfigurationResponseArrayOutputWithContext(context.Context) ApplianceNetworkConfigurationResponseArrayOutput
}

type ApplianceNetworkConfigurationResponseArray []ApplianceNetworkConfigurationResponseInput

func (ApplianceNetworkConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceNetworkConfigurationResponse)(nil)).Elem()
}

func (i ApplianceNetworkConfigurationResponseArray) ToApplianceNetworkConfigurationResponseArrayOutput() ApplianceNetworkConfigurationResponseArrayOutput {
	return i.ToApplianceNetworkConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i ApplianceNetworkConfigurationResponseArray) ToApplianceNetworkConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplianceNetworkConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceNetworkConfigurationResponseArrayOutput)
}

type ApplianceNetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ApplianceNetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceNetworkConfigurationResponse)(nil)).Elem()
}

func (o ApplianceNetworkConfigurationResponseOutput) ToApplianceNetworkConfigurationResponseOutput() ApplianceNetworkConfigurationResponseOutput {
	return o
}

func (o ApplianceNetworkConfigurationResponseOutput) ToApplianceNetworkConfigurationResponseOutputWithContext(ctx context.Context) ApplianceNetworkConfigurationResponseOutput {
	return o
}

func (o ApplianceNetworkConfigurationResponseOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v ApplianceNetworkConfigurationResponse) string { return v.MacAddress }).(pulumi.StringOutput)
}

func (o ApplianceNetworkConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplianceNetworkConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ApplianceNetworkConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplianceNetworkConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceNetworkConfigurationResponse)(nil)).Elem()
}

func (o ApplianceNetworkConfigurationResponseArrayOutput) ToApplianceNetworkConfigurationResponseArrayOutput() ApplianceNetworkConfigurationResponseArrayOutput {
	return o
}

func (o ApplianceNetworkConfigurationResponseArrayOutput) ToApplianceNetworkConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplianceNetworkConfigurationResponseArrayOutput {
	return o
}

func (o ApplianceNetworkConfigurationResponseArrayOutput) Index(i pulumi.IntInput) ApplianceNetworkConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplianceNetworkConfigurationResponse {
		return vs[0].([]ApplianceNetworkConfigurationResponse)[vs[1].(int)]
	}).(ApplianceNetworkConfigurationResponseOutput)
}

type ContactDetails struct {
	ContactName            string                   `pulumi:"contactName"`
	EmailList              []string                 `pulumi:"emailList"`
	Mobile                 *string                  `pulumi:"mobile"`
	NotificationPreference []NotificationPreference `pulumi:"notificationPreference"`
	Phone                  string                   `pulumi:"phone"`
	PhoneExtension         *string                  `pulumi:"phoneExtension"`
}





type ContactDetailsInput interface {
	pulumi.Input

	ToContactDetailsOutput() ContactDetailsOutput
	ToContactDetailsOutputWithContext(context.Context) ContactDetailsOutput
}

type ContactDetailsArgs struct {
	ContactName            pulumi.StringInput               `pulumi:"contactName"`
	EmailList              pulumi.StringArrayInput          `pulumi:"emailList"`
	Mobile                 pulumi.StringPtrInput            `pulumi:"mobile"`
	NotificationPreference NotificationPreferenceArrayInput `pulumi:"notificationPreference"`
	Phone                  pulumi.StringInput               `pulumi:"phone"`
	PhoneExtension         pulumi.StringPtrInput            `pulumi:"phoneExtension"`
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

func (o ContactDetailsOutput) NotificationPreference() NotificationPreferenceArrayOutput {
	return o.ApplyT(func(v ContactDetails) []NotificationPreference { return v.NotificationPreference }).(NotificationPreferenceArrayOutput)
}

func (o ContactDetailsOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetails) string { return v.Phone }).(pulumi.StringOutput)
}

func (o ContactDetailsOutput) PhoneExtension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetails) *string { return v.PhoneExtension }).(pulumi.StringPtrOutput)
}

type ContactDetailsResponse struct {
	ContactName            string                           `pulumi:"contactName"`
	EmailList              []string                         `pulumi:"emailList"`
	Mobile                 *string                          `pulumi:"mobile"`
	NotificationPreference []NotificationPreferenceResponse `pulumi:"notificationPreference"`
	Phone                  string                           `pulumi:"phone"`
	PhoneExtension         *string                          `pulumi:"phoneExtension"`
}





type ContactDetailsResponseInput interface {
	pulumi.Input

	ToContactDetailsResponseOutput() ContactDetailsResponseOutput
	ToContactDetailsResponseOutputWithContext(context.Context) ContactDetailsResponseOutput
}

type ContactDetailsResponseArgs struct {
	ContactName            pulumi.StringInput                       `pulumi:"contactName"`
	EmailList              pulumi.StringArrayInput                  `pulumi:"emailList"`
	Mobile                 pulumi.StringPtrInput                    `pulumi:"mobile"`
	NotificationPreference NotificationPreferenceResponseArrayInput `pulumi:"notificationPreference"`
	Phone                  pulumi.StringInput                       `pulumi:"phone"`
	PhoneExtension         pulumi.StringPtrInput                    `pulumi:"phoneExtension"`
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

func (o ContactDetailsResponseOutput) NotificationPreference() NotificationPreferenceResponseArrayOutput {
	return o.ApplyT(func(v ContactDetailsResponse) []NotificationPreferenceResponse { return v.NotificationPreference }).(NotificationPreferenceResponseArrayOutput)
}

func (o ContactDetailsResponseOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetailsResponse) string { return v.Phone }).(pulumi.StringOutput)
}

func (o ContactDetailsResponseOutput) PhoneExtension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailsResponse) *string { return v.PhoneExtension }).(pulumi.StringPtrOutput)
}

type CopyProgressResponse struct {
	AccountId                string  `pulumi:"accountId"`
	BytesSentToCloud         float64 `pulumi:"bytesSentToCloud"`
	DataDestinationType      string  `pulumi:"dataDestinationType"`
	FilesErroredOut          float64 `pulumi:"filesErroredOut"`
	FilesProcessed           float64 `pulumi:"filesProcessed"`
	InvalidFileBytesUploaded float64 `pulumi:"invalidFileBytesUploaded"`
	InvalidFilesProcessed    float64 `pulumi:"invalidFilesProcessed"`
	RenamedContainerCount    float64 `pulumi:"renamedContainerCount"`
	StorageAccountName       string  `pulumi:"storageAccountName"`
	TotalBytesToProcess      float64 `pulumi:"totalBytesToProcess"`
	TotalFilesToProcess      float64 `pulumi:"totalFilesToProcess"`
}





type CopyProgressResponseInput interface {
	pulumi.Input

	ToCopyProgressResponseOutput() CopyProgressResponseOutput
	ToCopyProgressResponseOutputWithContext(context.Context) CopyProgressResponseOutput
}

type CopyProgressResponseArgs struct {
	AccountId                pulumi.StringInput  `pulumi:"accountId"`
	BytesSentToCloud         pulumi.Float64Input `pulumi:"bytesSentToCloud"`
	DataDestinationType      pulumi.StringInput  `pulumi:"dataDestinationType"`
	FilesErroredOut          pulumi.Float64Input `pulumi:"filesErroredOut"`
	FilesProcessed           pulumi.Float64Input `pulumi:"filesProcessed"`
	InvalidFileBytesUploaded pulumi.Float64Input `pulumi:"invalidFileBytesUploaded"`
	InvalidFilesProcessed    pulumi.Float64Input `pulumi:"invalidFilesProcessed"`
	RenamedContainerCount    pulumi.Float64Input `pulumi:"renamedContainerCount"`
	StorageAccountName       pulumi.StringInput  `pulumi:"storageAccountName"`
	TotalBytesToProcess      pulumi.Float64Input `pulumi:"totalBytesToProcess"`
	TotalFilesToProcess      pulumi.Float64Input `pulumi:"totalFilesToProcess"`
}

func (CopyProgressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyProgressResponse)(nil)).Elem()
}

func (i CopyProgressResponseArgs) ToCopyProgressResponseOutput() CopyProgressResponseOutput {
	return i.ToCopyProgressResponseOutputWithContext(context.Background())
}

func (i CopyProgressResponseArgs) ToCopyProgressResponseOutputWithContext(ctx context.Context) CopyProgressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyProgressResponseOutput)
}





type CopyProgressResponseArrayInput interface {
	pulumi.Input

	ToCopyProgressResponseArrayOutput() CopyProgressResponseArrayOutput
	ToCopyProgressResponseArrayOutputWithContext(context.Context) CopyProgressResponseArrayOutput
}

type CopyProgressResponseArray []CopyProgressResponseInput

func (CopyProgressResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CopyProgressResponse)(nil)).Elem()
}

func (i CopyProgressResponseArray) ToCopyProgressResponseArrayOutput() CopyProgressResponseArrayOutput {
	return i.ToCopyProgressResponseArrayOutputWithContext(context.Background())
}

func (i CopyProgressResponseArray) ToCopyProgressResponseArrayOutputWithContext(ctx context.Context) CopyProgressResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyProgressResponseArrayOutput)
}

type CopyProgressResponseOutput struct{ *pulumi.OutputState }

func (CopyProgressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyProgressResponse)(nil)).Elem()
}

func (o CopyProgressResponseOutput) ToCopyProgressResponseOutput() CopyProgressResponseOutput {
	return o
}

func (o CopyProgressResponseOutput) ToCopyProgressResponseOutputWithContext(ctx context.Context) CopyProgressResponseOutput {
	return o
}

func (o CopyProgressResponseOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressResponse) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o CopyProgressResponseOutput) BytesSentToCloud() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.BytesSentToCloud }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) DataDestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressResponse) string { return v.DataDestinationType }).(pulumi.StringOutput)
}

func (o CopyProgressResponseOutput) FilesErroredOut() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.FilesErroredOut }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) FilesProcessed() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.FilesProcessed }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) InvalidFileBytesUploaded() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.InvalidFileBytesUploaded }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) InvalidFilesProcessed() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.InvalidFilesProcessed }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) RenamedContainerCount() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.RenamedContainerCount }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressResponse) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o CopyProgressResponseOutput) TotalBytesToProcess() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.TotalBytesToProcess }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) TotalFilesToProcess() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.TotalFilesToProcess }).(pulumi.Float64Output)
}

type CopyProgressResponseArrayOutput struct{ *pulumi.OutputState }

func (CopyProgressResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CopyProgressResponse)(nil)).Elem()
}

func (o CopyProgressResponseArrayOutput) ToCopyProgressResponseArrayOutput() CopyProgressResponseArrayOutput {
	return o
}

func (o CopyProgressResponseArrayOutput) ToCopyProgressResponseArrayOutputWithContext(ctx context.Context) CopyProgressResponseArrayOutput {
	return o
}

func (o CopyProgressResponseArrayOutput) Index(i pulumi.IntInput) CopyProgressResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CopyProgressResponse {
		return vs[0].([]CopyProgressResponse)[vs[1].(int)]
	}).(CopyProgressResponseOutput)
}

type DataBoxAccountCopyLogDetailsResponse struct {
	AccountName        string `pulumi:"accountName"`
	CopyLogDetailsType string `pulumi:"copyLogDetailsType"`
	CopyLogLink        string `pulumi:"copyLogLink"`
}





type DataBoxAccountCopyLogDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxAccountCopyLogDetailsResponseOutput() DataBoxAccountCopyLogDetailsResponseOutput
	ToDataBoxAccountCopyLogDetailsResponseOutputWithContext(context.Context) DataBoxAccountCopyLogDetailsResponseOutput
}

type DataBoxAccountCopyLogDetailsResponseArgs struct {
	AccountName        pulumi.StringInput `pulumi:"accountName"`
	CopyLogDetailsType pulumi.StringInput `pulumi:"copyLogDetailsType"`
	CopyLogLink        pulumi.StringInput `pulumi:"copyLogLink"`
}

func (DataBoxAccountCopyLogDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxAccountCopyLogDetailsResponse)(nil)).Elem()
}

func (i DataBoxAccountCopyLogDetailsResponseArgs) ToDataBoxAccountCopyLogDetailsResponseOutput() DataBoxAccountCopyLogDetailsResponseOutput {
	return i.ToDataBoxAccountCopyLogDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxAccountCopyLogDetailsResponseArgs) ToDataBoxAccountCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxAccountCopyLogDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxAccountCopyLogDetailsResponseOutput)
}

type DataBoxAccountCopyLogDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxAccountCopyLogDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxAccountCopyLogDetailsResponse)(nil)).Elem()
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) ToDataBoxAccountCopyLogDetailsResponseOutput() DataBoxAccountCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) ToDataBoxAccountCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxAccountCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxAccountCopyLogDetailsResponse) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) CopyLogDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxAccountCopyLogDetailsResponse) string { return v.CopyLogDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) CopyLogLink() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxAccountCopyLogDetailsResponse) string { return v.CopyLogLink }).(pulumi.StringOutput)
}

type DataBoxDiskCopyLogDetailsResponse struct {
	CopyLogDetailsType string `pulumi:"copyLogDetailsType"`
	DiskSerialNumber   string `pulumi:"diskSerialNumber"`
	ErrorLogLink       string `pulumi:"errorLogLink"`
	VerboseLogLink     string `pulumi:"verboseLogLink"`
}





type DataBoxDiskCopyLogDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxDiskCopyLogDetailsResponseOutput() DataBoxDiskCopyLogDetailsResponseOutput
	ToDataBoxDiskCopyLogDetailsResponseOutputWithContext(context.Context) DataBoxDiskCopyLogDetailsResponseOutput
}

type DataBoxDiskCopyLogDetailsResponseArgs struct {
	CopyLogDetailsType pulumi.StringInput `pulumi:"copyLogDetailsType"`
	DiskSerialNumber   pulumi.StringInput `pulumi:"diskSerialNumber"`
	ErrorLogLink       pulumi.StringInput `pulumi:"errorLogLink"`
	VerboseLogLink     pulumi.StringInput `pulumi:"verboseLogLink"`
}

func (DataBoxDiskCopyLogDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskCopyLogDetailsResponse)(nil)).Elem()
}

func (i DataBoxDiskCopyLogDetailsResponseArgs) ToDataBoxDiskCopyLogDetailsResponseOutput() DataBoxDiskCopyLogDetailsResponseOutput {
	return i.ToDataBoxDiskCopyLogDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxDiskCopyLogDetailsResponseArgs) ToDataBoxDiskCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxDiskCopyLogDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskCopyLogDetailsResponseOutput)
}

type DataBoxDiskCopyLogDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxDiskCopyLogDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskCopyLogDetailsResponse)(nil)).Elem()
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) ToDataBoxDiskCopyLogDetailsResponseOutput() DataBoxDiskCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) ToDataBoxDiskCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxDiskCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) CopyLogDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyLogDetailsResponse) string { return v.CopyLogDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) DiskSerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyLogDetailsResponse) string { return v.DiskSerialNumber }).(pulumi.StringOutput)
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) ErrorLogLink() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyLogDetailsResponse) string { return v.ErrorLogLink }).(pulumi.StringOutput)
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) VerboseLogLink() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyLogDetailsResponse) string { return v.VerboseLogLink }).(pulumi.StringOutput)
}

type DataBoxDiskCopyProgressResponse struct {
	BytesCopied     float64 `pulumi:"bytesCopied"`
	PercentComplete int     `pulumi:"percentComplete"`
	SerialNumber    string  `pulumi:"serialNumber"`
	Status          string  `pulumi:"status"`
}





type DataBoxDiskCopyProgressResponseInput interface {
	pulumi.Input

	ToDataBoxDiskCopyProgressResponseOutput() DataBoxDiskCopyProgressResponseOutput
	ToDataBoxDiskCopyProgressResponseOutputWithContext(context.Context) DataBoxDiskCopyProgressResponseOutput
}

type DataBoxDiskCopyProgressResponseArgs struct {
	BytesCopied     pulumi.Float64Input `pulumi:"bytesCopied"`
	PercentComplete pulumi.IntInput     `pulumi:"percentComplete"`
	SerialNumber    pulumi.StringInput  `pulumi:"serialNumber"`
	Status          pulumi.StringInput  `pulumi:"status"`
}

func (DataBoxDiskCopyProgressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskCopyProgressResponse)(nil)).Elem()
}

func (i DataBoxDiskCopyProgressResponseArgs) ToDataBoxDiskCopyProgressResponseOutput() DataBoxDiskCopyProgressResponseOutput {
	return i.ToDataBoxDiskCopyProgressResponseOutputWithContext(context.Background())
}

func (i DataBoxDiskCopyProgressResponseArgs) ToDataBoxDiskCopyProgressResponseOutputWithContext(ctx context.Context) DataBoxDiskCopyProgressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskCopyProgressResponseOutput)
}





type DataBoxDiskCopyProgressResponseArrayInput interface {
	pulumi.Input

	ToDataBoxDiskCopyProgressResponseArrayOutput() DataBoxDiskCopyProgressResponseArrayOutput
	ToDataBoxDiskCopyProgressResponseArrayOutputWithContext(context.Context) DataBoxDiskCopyProgressResponseArrayOutput
}

type DataBoxDiskCopyProgressResponseArray []DataBoxDiskCopyProgressResponseInput

func (DataBoxDiskCopyProgressResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxDiskCopyProgressResponse)(nil)).Elem()
}

func (i DataBoxDiskCopyProgressResponseArray) ToDataBoxDiskCopyProgressResponseArrayOutput() DataBoxDiskCopyProgressResponseArrayOutput {
	return i.ToDataBoxDiskCopyProgressResponseArrayOutputWithContext(context.Background())
}

func (i DataBoxDiskCopyProgressResponseArray) ToDataBoxDiskCopyProgressResponseArrayOutputWithContext(ctx context.Context) DataBoxDiskCopyProgressResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskCopyProgressResponseArrayOutput)
}

type DataBoxDiskCopyProgressResponseOutput struct{ *pulumi.OutputState }

func (DataBoxDiskCopyProgressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskCopyProgressResponse)(nil)).Elem()
}

func (o DataBoxDiskCopyProgressResponseOutput) ToDataBoxDiskCopyProgressResponseOutput() DataBoxDiskCopyProgressResponseOutput {
	return o
}

func (o DataBoxDiskCopyProgressResponseOutput) ToDataBoxDiskCopyProgressResponseOutputWithContext(ctx context.Context) DataBoxDiskCopyProgressResponseOutput {
	return o
}

func (o DataBoxDiskCopyProgressResponseOutput) BytesCopied() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxDiskCopyProgressResponse) float64 { return v.BytesCopied }).(pulumi.Float64Output)
}

func (o DataBoxDiskCopyProgressResponseOutput) PercentComplete() pulumi.IntOutput {
	return o.ApplyT(func(v DataBoxDiskCopyProgressResponse) int { return v.PercentComplete }).(pulumi.IntOutput)
}

func (o DataBoxDiskCopyProgressResponseOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyProgressResponse) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o DataBoxDiskCopyProgressResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyProgressResponse) string { return v.Status }).(pulumi.StringOutput)
}

type DataBoxDiskCopyProgressResponseArrayOutput struct{ *pulumi.OutputState }

func (DataBoxDiskCopyProgressResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxDiskCopyProgressResponse)(nil)).Elem()
}

func (o DataBoxDiskCopyProgressResponseArrayOutput) ToDataBoxDiskCopyProgressResponseArrayOutput() DataBoxDiskCopyProgressResponseArrayOutput {
	return o
}

func (o DataBoxDiskCopyProgressResponseArrayOutput) ToDataBoxDiskCopyProgressResponseArrayOutputWithContext(ctx context.Context) DataBoxDiskCopyProgressResponseArrayOutput {
	return o
}

func (o DataBoxDiskCopyProgressResponseArrayOutput) Index(i pulumi.IntInput) DataBoxDiskCopyProgressResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataBoxDiskCopyProgressResponse {
		return vs[0].([]DataBoxDiskCopyProgressResponse)[vs[1].(int)]
	}).(DataBoxDiskCopyProgressResponseOutput)
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





type DataBoxDiskJobDetailsInput interface {
	pulumi.Input

	ToDataBoxDiskJobDetailsOutput() DataBoxDiskJobDetailsOutput
	ToDataBoxDiskJobDetailsOutputWithContext(context.Context) DataBoxDiskJobDetailsOutput
}

type DataBoxDiskJobDetailsArgs struct {
	ContactDetails              ContactDetailsInput   `pulumi:"contactDetails"`
	DestinationAccountDetails   pulumi.ArrayInput     `pulumi:"destinationAccountDetails"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput    `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput    `pulumi:"jobDetailsType"`
	Passkey                     pulumi.StringPtrInput `pulumi:"passkey"`
	Preferences                 PreferencesPtrInput   `pulumi:"preferences"`
	PreferredDisks              pulumi.IntMapInput    `pulumi:"preferredDisks"`
	ShippingAddress             ShippingAddressInput  `pulumi:"shippingAddress"`
}

func (DataBoxDiskJobDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobDetails)(nil)).Elem()
}

func (i DataBoxDiskJobDetailsArgs) ToDataBoxDiskJobDetailsOutput() DataBoxDiskJobDetailsOutput {
	return i.ToDataBoxDiskJobDetailsOutputWithContext(context.Background())
}

func (i DataBoxDiskJobDetailsArgs) ToDataBoxDiskJobDetailsOutputWithContext(ctx context.Context) DataBoxDiskJobDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskJobDetailsOutput)
}

type DataBoxDiskJobDetailsOutput struct{ *pulumi.OutputState }

func (DataBoxDiskJobDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobDetails)(nil)).Elem()
}

func (o DataBoxDiskJobDetailsOutput) ToDataBoxDiskJobDetailsOutput() DataBoxDiskJobDetailsOutput {
	return o
}

func (o DataBoxDiskJobDetailsOutput) ToDataBoxDiskJobDetailsOutputWithContext(ctx context.Context) DataBoxDiskJobDetailsOutput {
	return o
}

func (o DataBoxDiskJobDetailsOutput) ContactDetails() ContactDetailsOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) ContactDetails { return v.ContactDetails }).(ContactDetailsOutput)
}

func (o DataBoxDiskJobDetailsOutput) DestinationAccountDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) []interface{} { return v.DestinationAccountDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxDiskJobDetailsOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxDiskJobDetailsOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobDetailsOutput) Passkey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) *string { return v.Passkey }).(pulumi.StringPtrOutput)
}

func (o DataBoxDiskJobDetailsOutput) Preferences() PreferencesPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) *Preferences { return v.Preferences }).(PreferencesPtrOutput)
}

func (o DataBoxDiskJobDetailsOutput) PreferredDisks() pulumi.IntMapOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) map[string]int { return v.PreferredDisks }).(pulumi.IntMapOutput)
}

func (o DataBoxDiskJobDetailsOutput) ShippingAddress() ShippingAddressOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) ShippingAddress { return v.ShippingAddress }).(ShippingAddressOutput)
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





type DataBoxDiskJobDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxDiskJobDetailsResponseOutput() DataBoxDiskJobDetailsResponseOutput
	ToDataBoxDiskJobDetailsResponseOutputWithContext(context.Context) DataBoxDiskJobDetailsResponseOutput
}

type DataBoxDiskJobDetailsResponseArgs struct {
	ChainOfCustodySasKey        pulumi.StringInput                        `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponseInput               `pulumi:"contactDetails"`
	CopyLogDetails              pulumi.ArrayInput                         `pulumi:"copyLogDetails"`
	CopyProgress                DataBoxDiskCopyProgressResponseArrayInput `pulumi:"copyProgress"`
	DeliveryPackage             PackageShippingDetailsResponseInput       `pulumi:"deliveryPackage"`
	DestinationAccountDetails   pulumi.ArrayInput                         `pulumi:"destinationAccountDetails"`
	DisksAndSizeDetails         pulumi.IntMapInput                        `pulumi:"disksAndSizeDetails"`
	ErrorDetails                JobErrorDetailsResponseArrayInput         `pulumi:"errorDetails"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput                        `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput                        `pulumi:"jobDetailsType"`
	JobStages                   JobStagesResponseArrayInput               `pulumi:"jobStages"`
	Passkey                     pulumi.StringPtrInput                     `pulumi:"passkey"`
	Preferences                 PreferencesResponsePtrInput               `pulumi:"preferences"`
	PreferredDisks              pulumi.IntMapInput                        `pulumi:"preferredDisks"`
	ReturnPackage               PackageShippingDetailsResponseInput       `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  pulumi.StringInput                        `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             ShippingAddressResponseInput              `pulumi:"shippingAddress"`
}

func (DataBoxDiskJobDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobDetailsResponse)(nil)).Elem()
}

func (i DataBoxDiskJobDetailsResponseArgs) ToDataBoxDiskJobDetailsResponseOutput() DataBoxDiskJobDetailsResponseOutput {
	return i.ToDataBoxDiskJobDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxDiskJobDetailsResponseArgs) ToDataBoxDiskJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxDiskJobDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskJobDetailsResponseOutput)
}

type DataBoxDiskJobDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxDiskJobDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobDetailsResponse)(nil)).Elem()
}

func (o DataBoxDiskJobDetailsResponseOutput) ToDataBoxDiskJobDetailsResponseOutput() DataBoxDiskJobDetailsResponseOutput {
	return o
}

func (o DataBoxDiskJobDetailsResponseOutput) ToDataBoxDiskJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxDiskJobDetailsResponseOutput {
	return o
}

func (o DataBoxDiskJobDetailsResponseOutput) ChainOfCustodySasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) string { return v.ChainOfCustodySasKey }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) CopyLogDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []interface{} { return v.CopyLogDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) CopyProgress() DataBoxDiskCopyProgressResponseArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []DataBoxDiskCopyProgressResponse { return v.CopyProgress }).(DataBoxDiskCopyProgressResponseArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) DeliveryPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) PackageShippingDetailsResponse { return v.DeliveryPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) DestinationAccountDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []interface{} { return v.DestinationAccountDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) DisksAndSizeDetails() pulumi.IntMapOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) map[string]int { return v.DisksAndSizeDetails }).(pulumi.IntMapOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ErrorDetails() JobErrorDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []JobErrorDetailsResponse { return v.ErrorDetails }).(JobErrorDetailsResponseArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) JobStages() JobStagesResponseArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []JobStagesResponse { return v.JobStages }).(JobStagesResponseArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) Passkey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) *string { return v.Passkey }).(pulumi.StringPtrOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) Preferences() PreferencesResponsePtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) *PreferencesResponse { return v.Preferences }).(PreferencesResponsePtrOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) PreferredDisks() pulumi.IntMapOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) map[string]int { return v.PreferredDisks }).(pulumi.IntMapOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ReturnPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) PackageShippingDetailsResponse { return v.ReturnPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ReverseShipmentLabelSasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) string { return v.ReverseShipmentLabelSasKey }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ShippingAddress() ShippingAddressResponseOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponseOutput)
}

type DataBoxDiskJobSecretsResponse struct {
	DcAccessSecurityCode *DcAccessSecurityCodeResponse `pulumi:"dcAccessSecurityCode"`
	DiskSecrets          []DiskSecretResponse          `pulumi:"diskSecrets"`
	IsPasskeyUserDefined bool                          `pulumi:"isPasskeyUserDefined"`
	JobSecretsType       string                        `pulumi:"jobSecretsType"`
	PassKey              string                        `pulumi:"passKey"`
}





type DataBoxDiskJobSecretsResponseInput interface {
	pulumi.Input

	ToDataBoxDiskJobSecretsResponseOutput() DataBoxDiskJobSecretsResponseOutput
	ToDataBoxDiskJobSecretsResponseOutputWithContext(context.Context) DataBoxDiskJobSecretsResponseOutput
}

type DataBoxDiskJobSecretsResponseArgs struct {
	DcAccessSecurityCode DcAccessSecurityCodeResponsePtrInput `pulumi:"dcAccessSecurityCode"`
	DiskSecrets          DiskSecretResponseArrayInput         `pulumi:"diskSecrets"`
	IsPasskeyUserDefined pulumi.BoolInput                     `pulumi:"isPasskeyUserDefined"`
	JobSecretsType       pulumi.StringInput                   `pulumi:"jobSecretsType"`
	PassKey              pulumi.StringInput                   `pulumi:"passKey"`
}

func (DataBoxDiskJobSecretsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobSecretsResponse)(nil)).Elem()
}

func (i DataBoxDiskJobSecretsResponseArgs) ToDataBoxDiskJobSecretsResponseOutput() DataBoxDiskJobSecretsResponseOutput {
	return i.ToDataBoxDiskJobSecretsResponseOutputWithContext(context.Background())
}

func (i DataBoxDiskJobSecretsResponseArgs) ToDataBoxDiskJobSecretsResponseOutputWithContext(ctx context.Context) DataBoxDiskJobSecretsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskJobSecretsResponseOutput)
}

type DataBoxDiskJobSecretsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxDiskJobSecretsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobSecretsResponse)(nil)).Elem()
}

func (o DataBoxDiskJobSecretsResponseOutput) ToDataBoxDiskJobSecretsResponseOutput() DataBoxDiskJobSecretsResponseOutput {
	return o
}

func (o DataBoxDiskJobSecretsResponseOutput) ToDataBoxDiskJobSecretsResponseOutputWithContext(ctx context.Context) DataBoxDiskJobSecretsResponseOutput {
	return o
}

func (o DataBoxDiskJobSecretsResponseOutput) DcAccessSecurityCode() DcAccessSecurityCodeResponsePtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) *DcAccessSecurityCodeResponse { return v.DcAccessSecurityCode }).(DcAccessSecurityCodeResponsePtrOutput)
}

func (o DataBoxDiskJobSecretsResponseOutput) DiskSecrets() DiskSecretResponseArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) []DiskSecretResponse { return v.DiskSecrets }).(DiskSecretResponseArrayOutput)
}

func (o DataBoxDiskJobSecretsResponseOutput) IsPasskeyUserDefined() pulumi.BoolOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) bool { return v.IsPasskeyUserDefined }).(pulumi.BoolOutput)
}

func (o DataBoxDiskJobSecretsResponseOutput) JobSecretsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) string { return v.JobSecretsType }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobSecretsResponseOutput) PassKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) string { return v.PassKey }).(pulumi.StringOutput)
}

type DataBoxHeavyAccountCopyLogDetailsResponse struct {
	AccountName        string   `pulumi:"accountName"`
	CopyLogDetailsType string   `pulumi:"copyLogDetailsType"`
	CopyLogLink        []string `pulumi:"copyLogLink"`
}





type DataBoxHeavyAccountCopyLogDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxHeavyAccountCopyLogDetailsResponseOutput() DataBoxHeavyAccountCopyLogDetailsResponseOutput
	ToDataBoxHeavyAccountCopyLogDetailsResponseOutputWithContext(context.Context) DataBoxHeavyAccountCopyLogDetailsResponseOutput
}

type DataBoxHeavyAccountCopyLogDetailsResponseArgs struct {
	AccountName        pulumi.StringInput      `pulumi:"accountName"`
	CopyLogDetailsType pulumi.StringInput      `pulumi:"copyLogDetailsType"`
	CopyLogLink        pulumi.StringArrayInput `pulumi:"copyLogLink"`
}

func (DataBoxHeavyAccountCopyLogDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyAccountCopyLogDetailsResponse)(nil)).Elem()
}

func (i DataBoxHeavyAccountCopyLogDetailsResponseArgs) ToDataBoxHeavyAccountCopyLogDetailsResponseOutput() DataBoxHeavyAccountCopyLogDetailsResponseOutput {
	return i.ToDataBoxHeavyAccountCopyLogDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxHeavyAccountCopyLogDetailsResponseArgs) ToDataBoxHeavyAccountCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxHeavyAccountCopyLogDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavyAccountCopyLogDetailsResponseOutput)
}

type DataBoxHeavyAccountCopyLogDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxHeavyAccountCopyLogDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyAccountCopyLogDetailsResponse)(nil)).Elem()
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) ToDataBoxHeavyAccountCopyLogDetailsResponseOutput() DataBoxHeavyAccountCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) ToDataBoxHeavyAccountCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxHeavyAccountCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyAccountCopyLogDetailsResponse) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) CopyLogDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyAccountCopyLogDetailsResponse) string { return v.CopyLogDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) CopyLogLink() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyAccountCopyLogDetailsResponse) []string { return v.CopyLogLink }).(pulumi.StringArrayOutput)
}

type DataBoxHeavyJobDetails struct {
	ContactDetails              ContactDetails  `pulumi:"contactDetails"`
	DestinationAccountDetails   []interface{}   `pulumi:"destinationAccountDetails"`
	DevicePassword              *string         `pulumi:"devicePassword"`
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





type DataBoxHeavyJobDetailsInput interface {
	pulumi.Input

	ToDataBoxHeavyJobDetailsOutput() DataBoxHeavyJobDetailsOutput
	ToDataBoxHeavyJobDetailsOutputWithContext(context.Context) DataBoxHeavyJobDetailsOutput
}

type DataBoxHeavyJobDetailsArgs struct {
	ContactDetails              ContactDetailsInput   `pulumi:"contactDetails"`
	DestinationAccountDetails   pulumi.ArrayInput     `pulumi:"destinationAccountDetails"`
	DevicePassword              pulumi.StringPtrInput `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput    `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput    `pulumi:"jobDetailsType"`
	Preferences                 PreferencesPtrInput   `pulumi:"preferences"`
	ShippingAddress             ShippingAddressInput  `pulumi:"shippingAddress"`
}

func (DataBoxHeavyJobDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobDetails)(nil)).Elem()
}

func (i DataBoxHeavyJobDetailsArgs) ToDataBoxHeavyJobDetailsOutput() DataBoxHeavyJobDetailsOutput {
	return i.ToDataBoxHeavyJobDetailsOutputWithContext(context.Background())
}

func (i DataBoxHeavyJobDetailsArgs) ToDataBoxHeavyJobDetailsOutputWithContext(ctx context.Context) DataBoxHeavyJobDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavyJobDetailsOutput)
}

type DataBoxHeavyJobDetailsOutput struct{ *pulumi.OutputState }

func (DataBoxHeavyJobDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobDetails)(nil)).Elem()
}

func (o DataBoxHeavyJobDetailsOutput) ToDataBoxHeavyJobDetailsOutput() DataBoxHeavyJobDetailsOutput {
	return o
}

func (o DataBoxHeavyJobDetailsOutput) ToDataBoxHeavyJobDetailsOutputWithContext(ctx context.Context) DataBoxHeavyJobDetailsOutput {
	return o
}

func (o DataBoxHeavyJobDetailsOutput) ContactDetails() ContactDetailsOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) ContactDetails { return v.ContactDetails }).(ContactDetailsOutput)
}

func (o DataBoxHeavyJobDetailsOutput) DestinationAccountDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) []interface{} { return v.DestinationAccountDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxHeavyJobDetailsOutput) DevicePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) *string { return v.DevicePassword }).(pulumi.StringPtrOutput)
}

func (o DataBoxHeavyJobDetailsOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxHeavyJobDetailsOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxHeavyJobDetailsOutput) Preferences() PreferencesPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) *Preferences { return v.Preferences }).(PreferencesPtrOutput)
}

func (o DataBoxHeavyJobDetailsOutput) ShippingAddress() ShippingAddressOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) ShippingAddress { return v.ShippingAddress }).(ShippingAddressOutput)
}

type DataBoxHeavyJobDetailsResponse struct {
	ChainOfCustodySasKey        string                         `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse         `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                  `pulumi:"copyLogDetails"`
	CopyProgress                []CopyProgressResponse         `pulumi:"copyProgress"`
	DeliveryPackage             PackageShippingDetailsResponse `pulumi:"deliveryPackage"`
	DestinationAccountDetails   []interface{}                  `pulumi:"destinationAccountDetails"`
	DevicePassword              *string                        `pulumi:"devicePassword"`
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





type DataBoxHeavyJobDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxHeavyJobDetailsResponseOutput() DataBoxHeavyJobDetailsResponseOutput
	ToDataBoxHeavyJobDetailsResponseOutputWithContext(context.Context) DataBoxHeavyJobDetailsResponseOutput
}

type DataBoxHeavyJobDetailsResponseArgs struct {
	ChainOfCustodySasKey        pulumi.StringInput                  `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponseInput         `pulumi:"contactDetails"`
	CopyLogDetails              pulumi.ArrayInput                   `pulumi:"copyLogDetails"`
	CopyProgress                CopyProgressResponseArrayInput      `pulumi:"copyProgress"`
	DeliveryPackage             PackageShippingDetailsResponseInput `pulumi:"deliveryPackage"`
	DestinationAccountDetails   pulumi.ArrayInput                   `pulumi:"destinationAccountDetails"`
	DevicePassword              pulumi.StringPtrInput               `pulumi:"devicePassword"`
	ErrorDetails                JobErrorDetailsResponseArrayInput   `pulumi:"errorDetails"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput                  `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput                  `pulumi:"jobDetailsType"`
	JobStages                   JobStagesResponseArrayInput         `pulumi:"jobStages"`
	Preferences                 PreferencesResponsePtrInput         `pulumi:"preferences"`
	ReturnPackage               PackageShippingDetailsResponseInput `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  pulumi.StringInput                  `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             ShippingAddressResponseInput        `pulumi:"shippingAddress"`
}

func (DataBoxHeavyJobDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobDetailsResponse)(nil)).Elem()
}

func (i DataBoxHeavyJobDetailsResponseArgs) ToDataBoxHeavyJobDetailsResponseOutput() DataBoxHeavyJobDetailsResponseOutput {
	return i.ToDataBoxHeavyJobDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxHeavyJobDetailsResponseArgs) ToDataBoxHeavyJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxHeavyJobDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavyJobDetailsResponseOutput)
}

type DataBoxHeavyJobDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxHeavyJobDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobDetailsResponse)(nil)).Elem()
}

func (o DataBoxHeavyJobDetailsResponseOutput) ToDataBoxHeavyJobDetailsResponseOutput() DataBoxHeavyJobDetailsResponseOutput {
	return o
}

func (o DataBoxHeavyJobDetailsResponseOutput) ToDataBoxHeavyJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxHeavyJobDetailsResponseOutput {
	return o
}

func (o DataBoxHeavyJobDetailsResponseOutput) ChainOfCustodySasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) string { return v.ChainOfCustodySasKey }).(pulumi.StringOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) CopyLogDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []interface{} { return v.CopyLogDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) CopyProgress() CopyProgressResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []CopyProgressResponse { return v.CopyProgress }).(CopyProgressResponseArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) DeliveryPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) PackageShippingDetailsResponse { return v.DeliveryPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) DestinationAccountDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []interface{} { return v.DestinationAccountDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) DevicePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) *string { return v.DevicePassword }).(pulumi.StringPtrOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ErrorDetails() JobErrorDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []JobErrorDetailsResponse { return v.ErrorDetails }).(JobErrorDetailsResponseArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) JobStages() JobStagesResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []JobStagesResponse { return v.JobStages }).(JobStagesResponseArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) Preferences() PreferencesResponsePtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) *PreferencesResponse { return v.Preferences }).(PreferencesResponsePtrOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ReturnPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) PackageShippingDetailsResponse { return v.ReturnPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ReverseShipmentLabelSasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) string { return v.ReverseShipmentLabelSasKey }).(pulumi.StringOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ShippingAddress() ShippingAddressResponseOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponseOutput)
}

type DataBoxHeavyJobSecretsResponse struct {
	CabinetPodSecrets    []DataBoxHeavySecretResponse  `pulumi:"cabinetPodSecrets"`
	DcAccessSecurityCode *DcAccessSecurityCodeResponse `pulumi:"dcAccessSecurityCode"`
	JobSecretsType       string                        `pulumi:"jobSecretsType"`
}





type DataBoxHeavyJobSecretsResponseInput interface {
	pulumi.Input

	ToDataBoxHeavyJobSecretsResponseOutput() DataBoxHeavyJobSecretsResponseOutput
	ToDataBoxHeavyJobSecretsResponseOutputWithContext(context.Context) DataBoxHeavyJobSecretsResponseOutput
}

type DataBoxHeavyJobSecretsResponseArgs struct {
	CabinetPodSecrets    DataBoxHeavySecretResponseArrayInput `pulumi:"cabinetPodSecrets"`
	DcAccessSecurityCode DcAccessSecurityCodeResponsePtrInput `pulumi:"dcAccessSecurityCode"`
	JobSecretsType       pulumi.StringInput                   `pulumi:"jobSecretsType"`
}

func (DataBoxHeavyJobSecretsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobSecretsResponse)(nil)).Elem()
}

func (i DataBoxHeavyJobSecretsResponseArgs) ToDataBoxHeavyJobSecretsResponseOutput() DataBoxHeavyJobSecretsResponseOutput {
	return i.ToDataBoxHeavyJobSecretsResponseOutputWithContext(context.Background())
}

func (i DataBoxHeavyJobSecretsResponseArgs) ToDataBoxHeavyJobSecretsResponseOutputWithContext(ctx context.Context) DataBoxHeavyJobSecretsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavyJobSecretsResponseOutput)
}

type DataBoxHeavyJobSecretsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxHeavyJobSecretsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobSecretsResponse)(nil)).Elem()
}

func (o DataBoxHeavyJobSecretsResponseOutput) ToDataBoxHeavyJobSecretsResponseOutput() DataBoxHeavyJobSecretsResponseOutput {
	return o
}

func (o DataBoxHeavyJobSecretsResponseOutput) ToDataBoxHeavyJobSecretsResponseOutputWithContext(ctx context.Context) DataBoxHeavyJobSecretsResponseOutput {
	return o
}

func (o DataBoxHeavyJobSecretsResponseOutput) CabinetPodSecrets() DataBoxHeavySecretResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobSecretsResponse) []DataBoxHeavySecretResponse { return v.CabinetPodSecrets }).(DataBoxHeavySecretResponseArrayOutput)
}

func (o DataBoxHeavyJobSecretsResponseOutput) DcAccessSecurityCode() DcAccessSecurityCodeResponsePtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobSecretsResponse) *DcAccessSecurityCodeResponse { return v.DcAccessSecurityCode }).(DcAccessSecurityCodeResponsePtrOutput)
}

func (o DataBoxHeavyJobSecretsResponseOutput) JobSecretsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobSecretsResponse) string { return v.JobSecretsType }).(pulumi.StringOutput)
}

type DataBoxHeavySecretResponse struct {
	AccountCredentialDetails    []AccountCredentialDetailsResponse      `pulumi:"accountCredentialDetails"`
	DevicePassword              string                                  `pulumi:"devicePassword"`
	DeviceSerialNumber          string                                  `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey string                                  `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       []ApplianceNetworkConfigurationResponse `pulumi:"networkConfigurations"`
}





type DataBoxHeavySecretResponseInput interface {
	pulumi.Input

	ToDataBoxHeavySecretResponseOutput() DataBoxHeavySecretResponseOutput
	ToDataBoxHeavySecretResponseOutputWithContext(context.Context) DataBoxHeavySecretResponseOutput
}

type DataBoxHeavySecretResponseArgs struct {
	AccountCredentialDetails    AccountCredentialDetailsResponseArrayInput      `pulumi:"accountCredentialDetails"`
	DevicePassword              pulumi.StringInput                              `pulumi:"devicePassword"`
	DeviceSerialNumber          pulumi.StringInput                              `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey pulumi.StringInput                              `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       ApplianceNetworkConfigurationResponseArrayInput `pulumi:"networkConfigurations"`
}

func (DataBoxHeavySecretResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavySecretResponse)(nil)).Elem()
}

func (i DataBoxHeavySecretResponseArgs) ToDataBoxHeavySecretResponseOutput() DataBoxHeavySecretResponseOutput {
	return i.ToDataBoxHeavySecretResponseOutputWithContext(context.Background())
}

func (i DataBoxHeavySecretResponseArgs) ToDataBoxHeavySecretResponseOutputWithContext(ctx context.Context) DataBoxHeavySecretResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavySecretResponseOutput)
}





type DataBoxHeavySecretResponseArrayInput interface {
	pulumi.Input

	ToDataBoxHeavySecretResponseArrayOutput() DataBoxHeavySecretResponseArrayOutput
	ToDataBoxHeavySecretResponseArrayOutputWithContext(context.Context) DataBoxHeavySecretResponseArrayOutput
}

type DataBoxHeavySecretResponseArray []DataBoxHeavySecretResponseInput

func (DataBoxHeavySecretResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxHeavySecretResponse)(nil)).Elem()
}

func (i DataBoxHeavySecretResponseArray) ToDataBoxHeavySecretResponseArrayOutput() DataBoxHeavySecretResponseArrayOutput {
	return i.ToDataBoxHeavySecretResponseArrayOutputWithContext(context.Background())
}

func (i DataBoxHeavySecretResponseArray) ToDataBoxHeavySecretResponseArrayOutputWithContext(ctx context.Context) DataBoxHeavySecretResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavySecretResponseArrayOutput)
}

type DataBoxHeavySecretResponseOutput struct{ *pulumi.OutputState }

func (DataBoxHeavySecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavySecretResponse)(nil)).Elem()
}

func (o DataBoxHeavySecretResponseOutput) ToDataBoxHeavySecretResponseOutput() DataBoxHeavySecretResponseOutput {
	return o
}

func (o DataBoxHeavySecretResponseOutput) ToDataBoxHeavySecretResponseOutputWithContext(ctx context.Context) DataBoxHeavySecretResponseOutput {
	return o
}

func (o DataBoxHeavySecretResponseOutput) AccountCredentialDetails() AccountCredentialDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavySecretResponse) []AccountCredentialDetailsResponse {
		return v.AccountCredentialDetails
	}).(AccountCredentialDetailsResponseArrayOutput)
}

func (o DataBoxHeavySecretResponseOutput) DevicePassword() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavySecretResponse) string { return v.DevicePassword }).(pulumi.StringOutput)
}

func (o DataBoxHeavySecretResponseOutput) DeviceSerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavySecretResponse) string { return v.DeviceSerialNumber }).(pulumi.StringOutput)
}

func (o DataBoxHeavySecretResponseOutput) EncodedValidationCertPubKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavySecretResponse) string { return v.EncodedValidationCertPubKey }).(pulumi.StringOutput)
}

func (o DataBoxHeavySecretResponseOutput) NetworkConfigurations() ApplianceNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavySecretResponse) []ApplianceNetworkConfigurationResponse {
		return v.NetworkConfigurations
	}).(ApplianceNetworkConfigurationResponseArrayOutput)
}

type DataBoxHeavySecretResponseArrayOutput struct{ *pulumi.OutputState }

func (DataBoxHeavySecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxHeavySecretResponse)(nil)).Elem()
}

func (o DataBoxHeavySecretResponseArrayOutput) ToDataBoxHeavySecretResponseArrayOutput() DataBoxHeavySecretResponseArrayOutput {
	return o
}

func (o DataBoxHeavySecretResponseArrayOutput) ToDataBoxHeavySecretResponseArrayOutputWithContext(ctx context.Context) DataBoxHeavySecretResponseArrayOutput {
	return o
}

func (o DataBoxHeavySecretResponseArrayOutput) Index(i pulumi.IntInput) DataBoxHeavySecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataBoxHeavySecretResponse {
		return vs[0].([]DataBoxHeavySecretResponse)[vs[1].(int)]
	}).(DataBoxHeavySecretResponseOutput)
}

type DataBoxJobDetails struct {
	ContactDetails              ContactDetails  `pulumi:"contactDetails"`
	DestinationAccountDetails   []interface{}   `pulumi:"destinationAccountDetails"`
	DevicePassword              *string         `pulumi:"devicePassword"`
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





type DataBoxJobDetailsInput interface {
	pulumi.Input

	ToDataBoxJobDetailsOutput() DataBoxJobDetailsOutput
	ToDataBoxJobDetailsOutputWithContext(context.Context) DataBoxJobDetailsOutput
}

type DataBoxJobDetailsArgs struct {
	ContactDetails              ContactDetailsInput   `pulumi:"contactDetails"`
	DestinationAccountDetails   pulumi.ArrayInput     `pulumi:"destinationAccountDetails"`
	DevicePassword              pulumi.StringPtrInput `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput    `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput    `pulumi:"jobDetailsType"`
	Preferences                 PreferencesPtrInput   `pulumi:"preferences"`
	ShippingAddress             ShippingAddressInput  `pulumi:"shippingAddress"`
}

func (DataBoxJobDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxJobDetails)(nil)).Elem()
}

func (i DataBoxJobDetailsArgs) ToDataBoxJobDetailsOutput() DataBoxJobDetailsOutput {
	return i.ToDataBoxJobDetailsOutputWithContext(context.Background())
}

func (i DataBoxJobDetailsArgs) ToDataBoxJobDetailsOutputWithContext(ctx context.Context) DataBoxJobDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxJobDetailsOutput)
}

type DataBoxJobDetailsOutput struct{ *pulumi.OutputState }

func (DataBoxJobDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxJobDetails)(nil)).Elem()
}

func (o DataBoxJobDetailsOutput) ToDataBoxJobDetailsOutput() DataBoxJobDetailsOutput {
	return o
}

func (o DataBoxJobDetailsOutput) ToDataBoxJobDetailsOutputWithContext(ctx context.Context) DataBoxJobDetailsOutput {
	return o
}

func (o DataBoxJobDetailsOutput) ContactDetails() ContactDetailsOutput {
	return o.ApplyT(func(v DataBoxJobDetails) ContactDetails { return v.ContactDetails }).(ContactDetailsOutput)
}

func (o DataBoxJobDetailsOutput) DestinationAccountDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetails) []interface{} { return v.DestinationAccountDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxJobDetailsOutput) DevicePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetails) *string { return v.DevicePassword }).(pulumi.StringPtrOutput)
}

func (o DataBoxJobDetailsOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetails) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxJobDetailsOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxJobDetails) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxJobDetailsOutput) Preferences() PreferencesPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetails) *Preferences { return v.Preferences }).(PreferencesPtrOutput)
}

func (o DataBoxJobDetailsOutput) ShippingAddress() ShippingAddressOutput {
	return o.ApplyT(func(v DataBoxJobDetails) ShippingAddress { return v.ShippingAddress }).(ShippingAddressOutput)
}

type DataBoxJobDetailsResponse struct {
	ChainOfCustodySasKey        string                         `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse         `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                  `pulumi:"copyLogDetails"`
	CopyProgress                []CopyProgressResponse         `pulumi:"copyProgress"`
	DeliveryPackage             PackageShippingDetailsResponse `pulumi:"deliveryPackage"`
	DestinationAccountDetails   []interface{}                  `pulumi:"destinationAccountDetails"`
	DevicePassword              *string                        `pulumi:"devicePassword"`
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





type DataBoxJobDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxJobDetailsResponseOutput() DataBoxJobDetailsResponseOutput
	ToDataBoxJobDetailsResponseOutputWithContext(context.Context) DataBoxJobDetailsResponseOutput
}

type DataBoxJobDetailsResponseArgs struct {
	ChainOfCustodySasKey        pulumi.StringInput                  `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponseInput         `pulumi:"contactDetails"`
	CopyLogDetails              pulumi.ArrayInput                   `pulumi:"copyLogDetails"`
	CopyProgress                CopyProgressResponseArrayInput      `pulumi:"copyProgress"`
	DeliveryPackage             PackageShippingDetailsResponseInput `pulumi:"deliveryPackage"`
	DestinationAccountDetails   pulumi.ArrayInput                   `pulumi:"destinationAccountDetails"`
	DevicePassword              pulumi.StringPtrInput               `pulumi:"devicePassword"`
	ErrorDetails                JobErrorDetailsResponseArrayInput   `pulumi:"errorDetails"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput                  `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput                  `pulumi:"jobDetailsType"`
	JobStages                   JobStagesResponseArrayInput         `pulumi:"jobStages"`
	Preferences                 PreferencesResponsePtrInput         `pulumi:"preferences"`
	ReturnPackage               PackageShippingDetailsResponseInput `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  pulumi.StringInput                  `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             ShippingAddressResponseInput        `pulumi:"shippingAddress"`
}

func (DataBoxJobDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxJobDetailsResponse)(nil)).Elem()
}

func (i DataBoxJobDetailsResponseArgs) ToDataBoxJobDetailsResponseOutput() DataBoxJobDetailsResponseOutput {
	return i.ToDataBoxJobDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxJobDetailsResponseArgs) ToDataBoxJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxJobDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxJobDetailsResponseOutput)
}

type DataBoxJobDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxJobDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxJobDetailsResponse)(nil)).Elem()
}

func (o DataBoxJobDetailsResponseOutput) ToDataBoxJobDetailsResponseOutput() DataBoxJobDetailsResponseOutput {
	return o
}

func (o DataBoxJobDetailsResponseOutput) ToDataBoxJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxJobDetailsResponseOutput {
	return o
}

func (o DataBoxJobDetailsResponseOutput) ChainOfCustodySasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) string { return v.ChainOfCustodySasKey }).(pulumi.StringOutput)
}

func (o DataBoxJobDetailsResponseOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o DataBoxJobDetailsResponseOutput) CopyLogDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []interface{} { return v.CopyLogDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) CopyProgress() CopyProgressResponseArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []CopyProgressResponse { return v.CopyProgress }).(CopyProgressResponseArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) DeliveryPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) PackageShippingDetailsResponse { return v.DeliveryPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxJobDetailsResponseOutput) DestinationAccountDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []interface{} { return v.DestinationAccountDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) DevicePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) *string { return v.DevicePassword }).(pulumi.StringPtrOutput)
}

func (o DataBoxJobDetailsResponseOutput) ErrorDetails() JobErrorDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []JobErrorDetailsResponse { return v.ErrorDetails }).(JobErrorDetailsResponseArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxJobDetailsResponseOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxJobDetailsResponseOutput) JobStages() JobStagesResponseArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []JobStagesResponse { return v.JobStages }).(JobStagesResponseArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) Preferences() PreferencesResponsePtrOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) *PreferencesResponse { return v.Preferences }).(PreferencesResponsePtrOutput)
}

func (o DataBoxJobDetailsResponseOutput) ReturnPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) PackageShippingDetailsResponse { return v.ReturnPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxJobDetailsResponseOutput) ReverseShipmentLabelSasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) string { return v.ReverseShipmentLabelSasKey }).(pulumi.StringOutput)
}

func (o DataBoxJobDetailsResponseOutput) ShippingAddress() ShippingAddressResponseOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponseOutput)
}

type DataBoxSecretResponse struct {
	AccountCredentialDetails    []AccountCredentialDetailsResponse      `pulumi:"accountCredentialDetails"`
	DevicePassword              string                                  `pulumi:"devicePassword"`
	DeviceSerialNumber          string                                  `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey string                                  `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       []ApplianceNetworkConfigurationResponse `pulumi:"networkConfigurations"`
}





type DataBoxSecretResponseInput interface {
	pulumi.Input

	ToDataBoxSecretResponseOutput() DataBoxSecretResponseOutput
	ToDataBoxSecretResponseOutputWithContext(context.Context) DataBoxSecretResponseOutput
}

type DataBoxSecretResponseArgs struct {
	AccountCredentialDetails    AccountCredentialDetailsResponseArrayInput      `pulumi:"accountCredentialDetails"`
	DevicePassword              pulumi.StringInput                              `pulumi:"devicePassword"`
	DeviceSerialNumber          pulumi.StringInput                              `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey pulumi.StringInput                              `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       ApplianceNetworkConfigurationResponseArrayInput `pulumi:"networkConfigurations"`
}

func (DataBoxSecretResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxSecretResponse)(nil)).Elem()
}

func (i DataBoxSecretResponseArgs) ToDataBoxSecretResponseOutput() DataBoxSecretResponseOutput {
	return i.ToDataBoxSecretResponseOutputWithContext(context.Background())
}

func (i DataBoxSecretResponseArgs) ToDataBoxSecretResponseOutputWithContext(ctx context.Context) DataBoxSecretResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxSecretResponseOutput)
}





type DataBoxSecretResponseArrayInput interface {
	pulumi.Input

	ToDataBoxSecretResponseArrayOutput() DataBoxSecretResponseArrayOutput
	ToDataBoxSecretResponseArrayOutputWithContext(context.Context) DataBoxSecretResponseArrayOutput
}

type DataBoxSecretResponseArray []DataBoxSecretResponseInput

func (DataBoxSecretResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxSecretResponse)(nil)).Elem()
}

func (i DataBoxSecretResponseArray) ToDataBoxSecretResponseArrayOutput() DataBoxSecretResponseArrayOutput {
	return i.ToDataBoxSecretResponseArrayOutputWithContext(context.Background())
}

func (i DataBoxSecretResponseArray) ToDataBoxSecretResponseArrayOutputWithContext(ctx context.Context) DataBoxSecretResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxSecretResponseArrayOutput)
}

type DataBoxSecretResponseOutput struct{ *pulumi.OutputState }

func (DataBoxSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxSecretResponse)(nil)).Elem()
}

func (o DataBoxSecretResponseOutput) ToDataBoxSecretResponseOutput() DataBoxSecretResponseOutput {
	return o
}

func (o DataBoxSecretResponseOutput) ToDataBoxSecretResponseOutputWithContext(ctx context.Context) DataBoxSecretResponseOutput {
	return o
}

func (o DataBoxSecretResponseOutput) AccountCredentialDetails() AccountCredentialDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxSecretResponse) []AccountCredentialDetailsResponse { return v.AccountCredentialDetails }).(AccountCredentialDetailsResponseArrayOutput)
}

func (o DataBoxSecretResponseOutput) DevicePassword() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxSecretResponse) string { return v.DevicePassword }).(pulumi.StringOutput)
}

func (o DataBoxSecretResponseOutput) DeviceSerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxSecretResponse) string { return v.DeviceSerialNumber }).(pulumi.StringOutput)
}

func (o DataBoxSecretResponseOutput) EncodedValidationCertPubKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxSecretResponse) string { return v.EncodedValidationCertPubKey }).(pulumi.StringOutput)
}

func (o DataBoxSecretResponseOutput) NetworkConfigurations() ApplianceNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v DataBoxSecretResponse) []ApplianceNetworkConfigurationResponse { return v.NetworkConfigurations }).(ApplianceNetworkConfigurationResponseArrayOutput)
}

type DataBoxSecretResponseArrayOutput struct{ *pulumi.OutputState }

func (DataBoxSecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxSecretResponse)(nil)).Elem()
}

func (o DataBoxSecretResponseArrayOutput) ToDataBoxSecretResponseArrayOutput() DataBoxSecretResponseArrayOutput {
	return o
}

func (o DataBoxSecretResponseArrayOutput) ToDataBoxSecretResponseArrayOutputWithContext(ctx context.Context) DataBoxSecretResponseArrayOutput {
	return o
}

func (o DataBoxSecretResponseArrayOutput) Index(i pulumi.IntInput) DataBoxSecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataBoxSecretResponse {
		return vs[0].([]DataBoxSecretResponse)[vs[1].(int)]
	}).(DataBoxSecretResponseOutput)
}

type DataboxJobSecretsResponse struct {
	DcAccessSecurityCode *DcAccessSecurityCodeResponse `pulumi:"dcAccessSecurityCode"`
	JobSecretsType       string                        `pulumi:"jobSecretsType"`
	PodSecrets           []DataBoxSecretResponse       `pulumi:"podSecrets"`
}





type DataboxJobSecretsResponseInput interface {
	pulumi.Input

	ToDataboxJobSecretsResponseOutput() DataboxJobSecretsResponseOutput
	ToDataboxJobSecretsResponseOutputWithContext(context.Context) DataboxJobSecretsResponseOutput
}

type DataboxJobSecretsResponseArgs struct {
	DcAccessSecurityCode DcAccessSecurityCodeResponsePtrInput `pulumi:"dcAccessSecurityCode"`
	JobSecretsType       pulumi.StringInput                   `pulumi:"jobSecretsType"`
	PodSecrets           DataBoxSecretResponseArrayInput      `pulumi:"podSecrets"`
}

func (DataboxJobSecretsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataboxJobSecretsResponse)(nil)).Elem()
}

func (i DataboxJobSecretsResponseArgs) ToDataboxJobSecretsResponseOutput() DataboxJobSecretsResponseOutput {
	return i.ToDataboxJobSecretsResponseOutputWithContext(context.Background())
}

func (i DataboxJobSecretsResponseArgs) ToDataboxJobSecretsResponseOutputWithContext(ctx context.Context) DataboxJobSecretsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataboxJobSecretsResponseOutput)
}

type DataboxJobSecretsResponseOutput struct{ *pulumi.OutputState }

func (DataboxJobSecretsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataboxJobSecretsResponse)(nil)).Elem()
}

func (o DataboxJobSecretsResponseOutput) ToDataboxJobSecretsResponseOutput() DataboxJobSecretsResponseOutput {
	return o
}

func (o DataboxJobSecretsResponseOutput) ToDataboxJobSecretsResponseOutputWithContext(ctx context.Context) DataboxJobSecretsResponseOutput {
	return o
}

func (o DataboxJobSecretsResponseOutput) DcAccessSecurityCode() DcAccessSecurityCodeResponsePtrOutput {
	return o.ApplyT(func(v DataboxJobSecretsResponse) *DcAccessSecurityCodeResponse { return v.DcAccessSecurityCode }).(DcAccessSecurityCodeResponsePtrOutput)
}

func (o DataboxJobSecretsResponseOutput) JobSecretsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataboxJobSecretsResponse) string { return v.JobSecretsType }).(pulumi.StringOutput)
}

func (o DataboxJobSecretsResponseOutput) PodSecrets() DataBoxSecretResponseArrayOutput {
	return o.ApplyT(func(v DataboxJobSecretsResponse) []DataBoxSecretResponse { return v.PodSecrets }).(DataBoxSecretResponseArrayOutput)
}

type DcAccessSecurityCodeResponse struct {
	ForwardDCAccessCode *string `pulumi:"forwardDCAccessCode"`
	ReverseDCAccessCode *string `pulumi:"reverseDCAccessCode"`
}





type DcAccessSecurityCodeResponseInput interface {
	pulumi.Input

	ToDcAccessSecurityCodeResponseOutput() DcAccessSecurityCodeResponseOutput
	ToDcAccessSecurityCodeResponseOutputWithContext(context.Context) DcAccessSecurityCodeResponseOutput
}

type DcAccessSecurityCodeResponseArgs struct {
	ForwardDCAccessCode pulumi.StringPtrInput `pulumi:"forwardDCAccessCode"`
	ReverseDCAccessCode pulumi.StringPtrInput `pulumi:"reverseDCAccessCode"`
}

func (DcAccessSecurityCodeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DcAccessSecurityCodeResponse)(nil)).Elem()
}

func (i DcAccessSecurityCodeResponseArgs) ToDcAccessSecurityCodeResponseOutput() DcAccessSecurityCodeResponseOutput {
	return i.ToDcAccessSecurityCodeResponseOutputWithContext(context.Background())
}

func (i DcAccessSecurityCodeResponseArgs) ToDcAccessSecurityCodeResponseOutputWithContext(ctx context.Context) DcAccessSecurityCodeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DcAccessSecurityCodeResponseOutput)
}

func (i DcAccessSecurityCodeResponseArgs) ToDcAccessSecurityCodeResponsePtrOutput() DcAccessSecurityCodeResponsePtrOutput {
	return i.ToDcAccessSecurityCodeResponsePtrOutputWithContext(context.Background())
}

func (i DcAccessSecurityCodeResponseArgs) ToDcAccessSecurityCodeResponsePtrOutputWithContext(ctx context.Context) DcAccessSecurityCodeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DcAccessSecurityCodeResponseOutput).ToDcAccessSecurityCodeResponsePtrOutputWithContext(ctx)
}









type DcAccessSecurityCodeResponsePtrInput interface {
	pulumi.Input

	ToDcAccessSecurityCodeResponsePtrOutput() DcAccessSecurityCodeResponsePtrOutput
	ToDcAccessSecurityCodeResponsePtrOutputWithContext(context.Context) DcAccessSecurityCodeResponsePtrOutput
}

type dcAccessSecurityCodeResponsePtrType DcAccessSecurityCodeResponseArgs

func DcAccessSecurityCodeResponsePtr(v *DcAccessSecurityCodeResponseArgs) DcAccessSecurityCodeResponsePtrInput {
	return (*dcAccessSecurityCodeResponsePtrType)(v)
}

func (*dcAccessSecurityCodeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DcAccessSecurityCodeResponse)(nil)).Elem()
}

func (i *dcAccessSecurityCodeResponsePtrType) ToDcAccessSecurityCodeResponsePtrOutput() DcAccessSecurityCodeResponsePtrOutput {
	return i.ToDcAccessSecurityCodeResponsePtrOutputWithContext(context.Background())
}

func (i *dcAccessSecurityCodeResponsePtrType) ToDcAccessSecurityCodeResponsePtrOutputWithContext(ctx context.Context) DcAccessSecurityCodeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DcAccessSecurityCodeResponsePtrOutput)
}

type DcAccessSecurityCodeResponseOutput struct{ *pulumi.OutputState }

func (DcAccessSecurityCodeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DcAccessSecurityCodeResponse)(nil)).Elem()
}

func (o DcAccessSecurityCodeResponseOutput) ToDcAccessSecurityCodeResponseOutput() DcAccessSecurityCodeResponseOutput {
	return o
}

func (o DcAccessSecurityCodeResponseOutput) ToDcAccessSecurityCodeResponseOutputWithContext(ctx context.Context) DcAccessSecurityCodeResponseOutput {
	return o
}

func (o DcAccessSecurityCodeResponseOutput) ToDcAccessSecurityCodeResponsePtrOutput() DcAccessSecurityCodeResponsePtrOutput {
	return o.ToDcAccessSecurityCodeResponsePtrOutputWithContext(context.Background())
}

func (o DcAccessSecurityCodeResponseOutput) ToDcAccessSecurityCodeResponsePtrOutputWithContext(ctx context.Context) DcAccessSecurityCodeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DcAccessSecurityCodeResponse) *DcAccessSecurityCodeResponse {
		return &v
	}).(DcAccessSecurityCodeResponsePtrOutput)
}

func (o DcAccessSecurityCodeResponseOutput) ForwardDCAccessCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DcAccessSecurityCodeResponse) *string { return v.ForwardDCAccessCode }).(pulumi.StringPtrOutput)
}

func (o DcAccessSecurityCodeResponseOutput) ReverseDCAccessCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DcAccessSecurityCodeResponse) *string { return v.ReverseDCAccessCode }).(pulumi.StringPtrOutput)
}

type DcAccessSecurityCodeResponsePtrOutput struct{ *pulumi.OutputState }

func (DcAccessSecurityCodeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DcAccessSecurityCodeResponse)(nil)).Elem()
}

func (o DcAccessSecurityCodeResponsePtrOutput) ToDcAccessSecurityCodeResponsePtrOutput() DcAccessSecurityCodeResponsePtrOutput {
	return o
}

func (o DcAccessSecurityCodeResponsePtrOutput) ToDcAccessSecurityCodeResponsePtrOutputWithContext(ctx context.Context) DcAccessSecurityCodeResponsePtrOutput {
	return o
}

func (o DcAccessSecurityCodeResponsePtrOutput) Elem() DcAccessSecurityCodeResponseOutput {
	return o.ApplyT(func(v *DcAccessSecurityCodeResponse) DcAccessSecurityCodeResponse {
		if v != nil {
			return *v
		}
		var ret DcAccessSecurityCodeResponse
		return ret
	}).(DcAccessSecurityCodeResponseOutput)
}

func (o DcAccessSecurityCodeResponsePtrOutput) ForwardDCAccessCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DcAccessSecurityCodeResponse) *string {
		if v == nil {
			return nil
		}
		return v.ForwardDCAccessCode
	}).(pulumi.StringPtrOutput)
}

func (o DcAccessSecurityCodeResponsePtrOutput) ReverseDCAccessCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DcAccessSecurityCodeResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReverseDCAccessCode
	}).(pulumi.StringPtrOutput)
}

type DestinationManagedDiskDetails struct {
	AccountId               *string `pulumi:"accountId"`
	DataDestinationType     string  `pulumi:"dataDestinationType"`
	ResourceGroupId         string  `pulumi:"resourceGroupId"`
	SharePassword           *string `pulumi:"sharePassword"`
	StagingStorageAccountId string  `pulumi:"stagingStorageAccountId"`
}





type DestinationManagedDiskDetailsInput interface {
	pulumi.Input

	ToDestinationManagedDiskDetailsOutput() DestinationManagedDiskDetailsOutput
	ToDestinationManagedDiskDetailsOutputWithContext(context.Context) DestinationManagedDiskDetailsOutput
}

type DestinationManagedDiskDetailsArgs struct {
	AccountId               pulumi.StringPtrInput `pulumi:"accountId"`
	DataDestinationType     pulumi.StringInput    `pulumi:"dataDestinationType"`
	ResourceGroupId         pulumi.StringInput    `pulumi:"resourceGroupId"`
	SharePassword           pulumi.StringPtrInput `pulumi:"sharePassword"`
	StagingStorageAccountId pulumi.StringInput    `pulumi:"stagingStorageAccountId"`
}

func (DestinationManagedDiskDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationManagedDiskDetails)(nil)).Elem()
}

func (i DestinationManagedDiskDetailsArgs) ToDestinationManagedDiskDetailsOutput() DestinationManagedDiskDetailsOutput {
	return i.ToDestinationManagedDiskDetailsOutputWithContext(context.Background())
}

func (i DestinationManagedDiskDetailsArgs) ToDestinationManagedDiskDetailsOutputWithContext(ctx context.Context) DestinationManagedDiskDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationManagedDiskDetailsOutput)
}

type DestinationManagedDiskDetailsOutput struct{ *pulumi.OutputState }

func (DestinationManagedDiskDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationManagedDiskDetails)(nil)).Elem()
}

func (o DestinationManagedDiskDetailsOutput) ToDestinationManagedDiskDetailsOutput() DestinationManagedDiskDetailsOutput {
	return o
}

func (o DestinationManagedDiskDetailsOutput) ToDestinationManagedDiskDetailsOutputWithContext(ctx context.Context) DestinationManagedDiskDetailsOutput {
	return o
}

func (o DestinationManagedDiskDetailsOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationManagedDiskDetails) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o DestinationManagedDiskDetailsOutput) DataDestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v DestinationManagedDiskDetails) string { return v.DataDestinationType }).(pulumi.StringOutput)
}

func (o DestinationManagedDiskDetailsOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v DestinationManagedDiskDetails) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

func (o DestinationManagedDiskDetailsOutput) SharePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationManagedDiskDetails) *string { return v.SharePassword }).(pulumi.StringPtrOutput)
}

func (o DestinationManagedDiskDetailsOutput) StagingStorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v DestinationManagedDiskDetails) string { return v.StagingStorageAccountId }).(pulumi.StringOutput)
}

type DestinationManagedDiskDetailsResponse struct {
	AccountId               *string `pulumi:"accountId"`
	DataDestinationType     string  `pulumi:"dataDestinationType"`
	ResourceGroupId         string  `pulumi:"resourceGroupId"`
	SharePassword           *string `pulumi:"sharePassword"`
	StagingStorageAccountId string  `pulumi:"stagingStorageAccountId"`
}





type DestinationManagedDiskDetailsResponseInput interface {
	pulumi.Input

	ToDestinationManagedDiskDetailsResponseOutput() DestinationManagedDiskDetailsResponseOutput
	ToDestinationManagedDiskDetailsResponseOutputWithContext(context.Context) DestinationManagedDiskDetailsResponseOutput
}

type DestinationManagedDiskDetailsResponseArgs struct {
	AccountId               pulumi.StringPtrInput `pulumi:"accountId"`
	DataDestinationType     pulumi.StringInput    `pulumi:"dataDestinationType"`
	ResourceGroupId         pulumi.StringInput    `pulumi:"resourceGroupId"`
	SharePassword           pulumi.StringPtrInput `pulumi:"sharePassword"`
	StagingStorageAccountId pulumi.StringInput    `pulumi:"stagingStorageAccountId"`
}

func (DestinationManagedDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationManagedDiskDetailsResponse)(nil)).Elem()
}

func (i DestinationManagedDiskDetailsResponseArgs) ToDestinationManagedDiskDetailsResponseOutput() DestinationManagedDiskDetailsResponseOutput {
	return i.ToDestinationManagedDiskDetailsResponseOutputWithContext(context.Background())
}

func (i DestinationManagedDiskDetailsResponseArgs) ToDestinationManagedDiskDetailsResponseOutputWithContext(ctx context.Context) DestinationManagedDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationManagedDiskDetailsResponseOutput)
}

type DestinationManagedDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (DestinationManagedDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationManagedDiskDetailsResponse)(nil)).Elem()
}

func (o DestinationManagedDiskDetailsResponseOutput) ToDestinationManagedDiskDetailsResponseOutput() DestinationManagedDiskDetailsResponseOutput {
	return o
}

func (o DestinationManagedDiskDetailsResponseOutput) ToDestinationManagedDiskDetailsResponseOutputWithContext(ctx context.Context) DestinationManagedDiskDetailsResponseOutput {
	return o
}

func (o DestinationManagedDiskDetailsResponseOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationManagedDiskDetailsResponse) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o DestinationManagedDiskDetailsResponseOutput) DataDestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v DestinationManagedDiskDetailsResponse) string { return v.DataDestinationType }).(pulumi.StringOutput)
}

func (o DestinationManagedDiskDetailsResponseOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v DestinationManagedDiskDetailsResponse) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

func (o DestinationManagedDiskDetailsResponseOutput) SharePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationManagedDiskDetailsResponse) *string { return v.SharePassword }).(pulumi.StringPtrOutput)
}

func (o DestinationManagedDiskDetailsResponseOutput) StagingStorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v DestinationManagedDiskDetailsResponse) string { return v.StagingStorageAccountId }).(pulumi.StringOutput)
}

type DestinationStorageAccountDetails struct {
	AccountId           *string `pulumi:"accountId"`
	DataDestinationType string  `pulumi:"dataDestinationType"`
	SharePassword       *string `pulumi:"sharePassword"`
	StorageAccountId    string  `pulumi:"storageAccountId"`
}





type DestinationStorageAccountDetailsInput interface {
	pulumi.Input

	ToDestinationStorageAccountDetailsOutput() DestinationStorageAccountDetailsOutput
	ToDestinationStorageAccountDetailsOutputWithContext(context.Context) DestinationStorageAccountDetailsOutput
}

type DestinationStorageAccountDetailsArgs struct {
	AccountId           pulumi.StringPtrInput `pulumi:"accountId"`
	DataDestinationType pulumi.StringInput    `pulumi:"dataDestinationType"`
	SharePassword       pulumi.StringPtrInput `pulumi:"sharePassword"`
	StorageAccountId    pulumi.StringInput    `pulumi:"storageAccountId"`
}

func (DestinationStorageAccountDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationStorageAccountDetails)(nil)).Elem()
}

func (i DestinationStorageAccountDetailsArgs) ToDestinationStorageAccountDetailsOutput() DestinationStorageAccountDetailsOutput {
	return i.ToDestinationStorageAccountDetailsOutputWithContext(context.Background())
}

func (i DestinationStorageAccountDetailsArgs) ToDestinationStorageAccountDetailsOutputWithContext(ctx context.Context) DestinationStorageAccountDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationStorageAccountDetailsOutput)
}

type DestinationStorageAccountDetailsOutput struct{ *pulumi.OutputState }

func (DestinationStorageAccountDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationStorageAccountDetails)(nil)).Elem()
}

func (o DestinationStorageAccountDetailsOutput) ToDestinationStorageAccountDetailsOutput() DestinationStorageAccountDetailsOutput {
	return o
}

func (o DestinationStorageAccountDetailsOutput) ToDestinationStorageAccountDetailsOutputWithContext(ctx context.Context) DestinationStorageAccountDetailsOutput {
	return o
}

func (o DestinationStorageAccountDetailsOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationStorageAccountDetails) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o DestinationStorageAccountDetailsOutput) DataDestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v DestinationStorageAccountDetails) string { return v.DataDestinationType }).(pulumi.StringOutput)
}

func (o DestinationStorageAccountDetailsOutput) SharePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationStorageAccountDetails) *string { return v.SharePassword }).(pulumi.StringPtrOutput)
}

func (o DestinationStorageAccountDetailsOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v DestinationStorageAccountDetails) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type DestinationStorageAccountDetailsResponse struct {
	AccountId           *string `pulumi:"accountId"`
	DataDestinationType string  `pulumi:"dataDestinationType"`
	SharePassword       *string `pulumi:"sharePassword"`
	StorageAccountId    string  `pulumi:"storageAccountId"`
}





type DestinationStorageAccountDetailsResponseInput interface {
	pulumi.Input

	ToDestinationStorageAccountDetailsResponseOutput() DestinationStorageAccountDetailsResponseOutput
	ToDestinationStorageAccountDetailsResponseOutputWithContext(context.Context) DestinationStorageAccountDetailsResponseOutput
}

type DestinationStorageAccountDetailsResponseArgs struct {
	AccountId           pulumi.StringPtrInput `pulumi:"accountId"`
	DataDestinationType pulumi.StringInput    `pulumi:"dataDestinationType"`
	SharePassword       pulumi.StringPtrInput `pulumi:"sharePassword"`
	StorageAccountId    pulumi.StringInput    `pulumi:"storageAccountId"`
}

func (DestinationStorageAccountDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationStorageAccountDetailsResponse)(nil)).Elem()
}

func (i DestinationStorageAccountDetailsResponseArgs) ToDestinationStorageAccountDetailsResponseOutput() DestinationStorageAccountDetailsResponseOutput {
	return i.ToDestinationStorageAccountDetailsResponseOutputWithContext(context.Background())
}

func (i DestinationStorageAccountDetailsResponseArgs) ToDestinationStorageAccountDetailsResponseOutputWithContext(ctx context.Context) DestinationStorageAccountDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationStorageAccountDetailsResponseOutput)
}

type DestinationStorageAccountDetailsResponseOutput struct{ *pulumi.OutputState }

func (DestinationStorageAccountDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationStorageAccountDetailsResponse)(nil)).Elem()
}

func (o DestinationStorageAccountDetailsResponseOutput) ToDestinationStorageAccountDetailsResponseOutput() DestinationStorageAccountDetailsResponseOutput {
	return o
}

func (o DestinationStorageAccountDetailsResponseOutput) ToDestinationStorageAccountDetailsResponseOutputWithContext(ctx context.Context) DestinationStorageAccountDetailsResponseOutput {
	return o
}

func (o DestinationStorageAccountDetailsResponseOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationStorageAccountDetailsResponse) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o DestinationStorageAccountDetailsResponseOutput) DataDestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v DestinationStorageAccountDetailsResponse) string { return v.DataDestinationType }).(pulumi.StringOutput)
}

func (o DestinationStorageAccountDetailsResponseOutput) SharePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationStorageAccountDetailsResponse) *string { return v.SharePassword }).(pulumi.StringPtrOutput)
}

func (o DestinationStorageAccountDetailsResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v DestinationStorageAccountDetailsResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type DiskSecretResponse struct {
	BitLockerKey     string `pulumi:"bitLockerKey"`
	DiskSerialNumber string `pulumi:"diskSerialNumber"`
}





type DiskSecretResponseInput interface {
	pulumi.Input

	ToDiskSecretResponseOutput() DiskSecretResponseOutput
	ToDiskSecretResponseOutputWithContext(context.Context) DiskSecretResponseOutput
}

type DiskSecretResponseArgs struct {
	BitLockerKey     pulumi.StringInput `pulumi:"bitLockerKey"`
	DiskSerialNumber pulumi.StringInput `pulumi:"diskSerialNumber"`
}

func (DiskSecretResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSecretResponse)(nil)).Elem()
}

func (i DiskSecretResponseArgs) ToDiskSecretResponseOutput() DiskSecretResponseOutput {
	return i.ToDiskSecretResponseOutputWithContext(context.Background())
}

func (i DiskSecretResponseArgs) ToDiskSecretResponseOutputWithContext(ctx context.Context) DiskSecretResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSecretResponseOutput)
}





type DiskSecretResponseArrayInput interface {
	pulumi.Input

	ToDiskSecretResponseArrayOutput() DiskSecretResponseArrayOutput
	ToDiskSecretResponseArrayOutputWithContext(context.Context) DiskSecretResponseArrayOutput
}

type DiskSecretResponseArray []DiskSecretResponseInput

func (DiskSecretResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskSecretResponse)(nil)).Elem()
}

func (i DiskSecretResponseArray) ToDiskSecretResponseArrayOutput() DiskSecretResponseArrayOutput {
	return i.ToDiskSecretResponseArrayOutputWithContext(context.Background())
}

func (i DiskSecretResponseArray) ToDiskSecretResponseArrayOutputWithContext(ctx context.Context) DiskSecretResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSecretResponseArrayOutput)
}

type DiskSecretResponseOutput struct{ *pulumi.OutputState }

func (DiskSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSecretResponse)(nil)).Elem()
}

func (o DiskSecretResponseOutput) ToDiskSecretResponseOutput() DiskSecretResponseOutput {
	return o
}

func (o DiskSecretResponseOutput) ToDiskSecretResponseOutputWithContext(ctx context.Context) DiskSecretResponseOutput {
	return o
}

func (o DiskSecretResponseOutput) BitLockerKey() pulumi.StringOutput {
	return o.ApplyT(func(v DiskSecretResponse) string { return v.BitLockerKey }).(pulumi.StringOutput)
}

func (o DiskSecretResponseOutput) DiskSerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DiskSecretResponse) string { return v.DiskSerialNumber }).(pulumi.StringOutput)
}

type DiskSecretResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskSecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskSecretResponse)(nil)).Elem()
}

func (o DiskSecretResponseArrayOutput) ToDiskSecretResponseArrayOutput() DiskSecretResponseArrayOutput {
	return o
}

func (o DiskSecretResponseArrayOutput) ToDiskSecretResponseArrayOutputWithContext(ctx context.Context) DiskSecretResponseArrayOutput {
	return o
}

func (o DiskSecretResponseArrayOutput) Index(i pulumi.IntInput) DiskSecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskSecretResponse {
		return vs[0].([]DiskSecretResponse)[vs[1].(int)]
	}).(DiskSecretResponseOutput)
}

type ErrorResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}





type ErrorResponseInput interface {
	pulumi.Input

	ToErrorResponseOutput() ErrorResponseOutput
	ToErrorResponseOutputWithContext(context.Context) ErrorResponseOutput
}

type ErrorResponseArgs struct {
	Code    pulumi.StringInput `pulumi:"code"`
	Message pulumi.StringInput `pulumi:"message"`
}

func (ErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponse)(nil)).Elem()
}

func (i ErrorResponseArgs) ToErrorResponseOutput() ErrorResponseOutput {
	return i.ToErrorResponseOutputWithContext(context.Background())
}

func (i ErrorResponseArgs) ToErrorResponseOutputWithContext(ctx context.Context) ErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseOutput)
}

func (i ErrorResponseArgs) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return i.ToErrorResponsePtrOutputWithContext(context.Background())
}

func (i ErrorResponseArgs) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseOutput).ToErrorResponsePtrOutputWithContext(ctx)
}









type ErrorResponsePtrInput interface {
	pulumi.Input

	ToErrorResponsePtrOutput() ErrorResponsePtrOutput
	ToErrorResponsePtrOutputWithContext(context.Context) ErrorResponsePtrOutput
}

type errorResponsePtrType ErrorResponseArgs

func ErrorResponsePtr(v *ErrorResponseArgs) ErrorResponsePtrInput {
	return (*errorResponsePtrType)(v)
}

func (*errorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponse)(nil)).Elem()
}

func (i *errorResponsePtrType) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return i.ToErrorResponsePtrOutputWithContext(context.Background())
}

func (i *errorResponsePtrType) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponsePtrOutput)
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

func (o ErrorResponseOutput) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return o.ToErrorResponsePtrOutputWithContext(context.Background())
}

func (o ErrorResponseOutput) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ErrorResponse) *ErrorResponse {
		return &v
	}).(ErrorResponsePtrOutput)
}

func (o ErrorResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponse) string { return v.Message }).(pulumi.StringOutput)
}

type ErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponse)(nil)).Elem()
}

func (o ErrorResponsePtrOutput) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return o
}

func (o ErrorResponsePtrOutput) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return o
}

func (o ErrorResponsePtrOutput) Elem() ErrorResponseOutput {
	return o.ApplyT(func(v *ErrorResponse) ErrorResponse {
		if v != nil {
			return *v
		}
		var ret ErrorResponse
		return ret
	}).(ErrorResponseOutput)
}

func (o ErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
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





type JobDeliveryInfoResponseInput interface {
	pulumi.Input

	ToJobDeliveryInfoResponseOutput() JobDeliveryInfoResponseOutput
	ToJobDeliveryInfoResponseOutputWithContext(context.Context) JobDeliveryInfoResponseOutput
}

type JobDeliveryInfoResponseArgs struct {
	ScheduledDateTime pulumi.StringPtrInput `pulumi:"scheduledDateTime"`
}

func (JobDeliveryInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDeliveryInfoResponse)(nil)).Elem()
}

func (i JobDeliveryInfoResponseArgs) ToJobDeliveryInfoResponseOutput() JobDeliveryInfoResponseOutput {
	return i.ToJobDeliveryInfoResponseOutputWithContext(context.Background())
}

func (i JobDeliveryInfoResponseArgs) ToJobDeliveryInfoResponseOutputWithContext(ctx context.Context) JobDeliveryInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoResponseOutput)
}

func (i JobDeliveryInfoResponseArgs) ToJobDeliveryInfoResponsePtrOutput() JobDeliveryInfoResponsePtrOutput {
	return i.ToJobDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (i JobDeliveryInfoResponseArgs) ToJobDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) JobDeliveryInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoResponseOutput).ToJobDeliveryInfoResponsePtrOutputWithContext(ctx)
}









type JobDeliveryInfoResponsePtrInput interface {
	pulumi.Input

	ToJobDeliveryInfoResponsePtrOutput() JobDeliveryInfoResponsePtrOutput
	ToJobDeliveryInfoResponsePtrOutputWithContext(context.Context) JobDeliveryInfoResponsePtrOutput
}

type jobDeliveryInfoResponsePtrType JobDeliveryInfoResponseArgs

func JobDeliveryInfoResponsePtr(v *JobDeliveryInfoResponseArgs) JobDeliveryInfoResponsePtrInput {
	return (*jobDeliveryInfoResponsePtrType)(v)
}

func (*jobDeliveryInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDeliveryInfoResponse)(nil)).Elem()
}

func (i *jobDeliveryInfoResponsePtrType) ToJobDeliveryInfoResponsePtrOutput() JobDeliveryInfoResponsePtrOutput {
	return i.ToJobDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (i *jobDeliveryInfoResponsePtrType) ToJobDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) JobDeliveryInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoResponsePtrOutput)
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

func (o JobDeliveryInfoResponseOutput) ToJobDeliveryInfoResponsePtrOutput() JobDeliveryInfoResponsePtrOutput {
	return o.ToJobDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (o JobDeliveryInfoResponseOutput) ToJobDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) JobDeliveryInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobDeliveryInfoResponse) *JobDeliveryInfoResponse {
		return &v
	}).(JobDeliveryInfoResponsePtrOutput)
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

type JobErrorDetailsResponse struct {
	ErrorCode         int    `pulumi:"errorCode"`
	ErrorMessage      string `pulumi:"errorMessage"`
	ExceptionMessage  string `pulumi:"exceptionMessage"`
	RecommendedAction string `pulumi:"recommendedAction"`
}





type JobErrorDetailsResponseInput interface {
	pulumi.Input

	ToJobErrorDetailsResponseOutput() JobErrorDetailsResponseOutput
	ToJobErrorDetailsResponseOutputWithContext(context.Context) JobErrorDetailsResponseOutput
}

type JobErrorDetailsResponseArgs struct {
	ErrorCode         pulumi.IntInput    `pulumi:"errorCode"`
	ErrorMessage      pulumi.StringInput `pulumi:"errorMessage"`
	ExceptionMessage  pulumi.StringInput `pulumi:"exceptionMessage"`
	RecommendedAction pulumi.StringInput `pulumi:"recommendedAction"`
}

func (JobErrorDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorDetailsResponse)(nil)).Elem()
}

func (i JobErrorDetailsResponseArgs) ToJobErrorDetailsResponseOutput() JobErrorDetailsResponseOutput {
	return i.ToJobErrorDetailsResponseOutputWithContext(context.Background())
}

func (i JobErrorDetailsResponseArgs) ToJobErrorDetailsResponseOutputWithContext(ctx context.Context) JobErrorDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorDetailsResponseOutput)
}





type JobErrorDetailsResponseArrayInput interface {
	pulumi.Input

	ToJobErrorDetailsResponseArrayOutput() JobErrorDetailsResponseArrayOutput
	ToJobErrorDetailsResponseArrayOutputWithContext(context.Context) JobErrorDetailsResponseArrayOutput
}

type JobErrorDetailsResponseArray []JobErrorDetailsResponseInput

func (JobErrorDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobErrorDetailsResponse)(nil)).Elem()
}

func (i JobErrorDetailsResponseArray) ToJobErrorDetailsResponseArrayOutput() JobErrorDetailsResponseArrayOutput {
	return i.ToJobErrorDetailsResponseArrayOutputWithContext(context.Background())
}

func (i JobErrorDetailsResponseArray) ToJobErrorDetailsResponseArrayOutputWithContext(ctx context.Context) JobErrorDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorDetailsResponseArrayOutput)
}

type JobErrorDetailsResponseOutput struct{ *pulumi.OutputState }

func (JobErrorDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorDetailsResponse)(nil)).Elem()
}

func (o JobErrorDetailsResponseOutput) ToJobErrorDetailsResponseOutput() JobErrorDetailsResponseOutput {
	return o
}

func (o JobErrorDetailsResponseOutput) ToJobErrorDetailsResponseOutputWithContext(ctx context.Context) JobErrorDetailsResponseOutput {
	return o
}

func (o JobErrorDetailsResponseOutput) ErrorCode() pulumi.IntOutput {
	return o.ApplyT(func(v JobErrorDetailsResponse) int { return v.ErrorCode }).(pulumi.IntOutput)
}

func (o JobErrorDetailsResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorDetailsResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o JobErrorDetailsResponseOutput) ExceptionMessage() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorDetailsResponse) string { return v.ExceptionMessage }).(pulumi.StringOutput)
}

func (o JobErrorDetailsResponseOutput) RecommendedAction() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorDetailsResponse) string { return v.RecommendedAction }).(pulumi.StringOutput)
}

type JobErrorDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (JobErrorDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobErrorDetailsResponse)(nil)).Elem()
}

func (o JobErrorDetailsResponseArrayOutput) ToJobErrorDetailsResponseArrayOutput() JobErrorDetailsResponseArrayOutput {
	return o
}

func (o JobErrorDetailsResponseArrayOutput) ToJobErrorDetailsResponseArrayOutputWithContext(ctx context.Context) JobErrorDetailsResponseArrayOutput {
	return o
}

func (o JobErrorDetailsResponseArrayOutput) Index(i pulumi.IntInput) JobErrorDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobErrorDetailsResponse {
		return vs[0].([]JobErrorDetailsResponse)[vs[1].(int)]
	}).(JobErrorDetailsResponseOutput)
}

type JobStagesResponse struct {
	DisplayName     string                    `pulumi:"displayName"`
	ErrorDetails    []JobErrorDetailsResponse `pulumi:"errorDetails"`
	JobStageDetails interface{}               `pulumi:"jobStageDetails"`
	StageName       string                    `pulumi:"stageName"`
	StageStatus     string                    `pulumi:"stageStatus"`
	StageTime       string                    `pulumi:"stageTime"`
}





type JobStagesResponseInput interface {
	pulumi.Input

	ToJobStagesResponseOutput() JobStagesResponseOutput
	ToJobStagesResponseOutputWithContext(context.Context) JobStagesResponseOutput
}

type JobStagesResponseArgs struct {
	DisplayName     pulumi.StringInput                `pulumi:"displayName"`
	ErrorDetails    JobErrorDetailsResponseArrayInput `pulumi:"errorDetails"`
	JobStageDetails pulumi.Input                      `pulumi:"jobStageDetails"`
	StageName       pulumi.StringInput                `pulumi:"stageName"`
	StageStatus     pulumi.StringInput                `pulumi:"stageStatus"`
	StageTime       pulumi.StringInput                `pulumi:"stageTime"`
}

func (JobStagesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStagesResponse)(nil)).Elem()
}

func (i JobStagesResponseArgs) ToJobStagesResponseOutput() JobStagesResponseOutput {
	return i.ToJobStagesResponseOutputWithContext(context.Background())
}

func (i JobStagesResponseArgs) ToJobStagesResponseOutputWithContext(ctx context.Context) JobStagesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStagesResponseOutput)
}





type JobStagesResponseArrayInput interface {
	pulumi.Input

	ToJobStagesResponseArrayOutput() JobStagesResponseArrayOutput
	ToJobStagesResponseArrayOutputWithContext(context.Context) JobStagesResponseArrayOutput
}

type JobStagesResponseArray []JobStagesResponseInput

func (JobStagesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobStagesResponse)(nil)).Elem()
}

func (i JobStagesResponseArray) ToJobStagesResponseArrayOutput() JobStagesResponseArrayOutput {
	return i.ToJobStagesResponseArrayOutputWithContext(context.Background())
}

func (i JobStagesResponseArray) ToJobStagesResponseArrayOutputWithContext(ctx context.Context) JobStagesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStagesResponseArrayOutput)
}

type JobStagesResponseOutput struct{ *pulumi.OutputState }

func (JobStagesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStagesResponse)(nil)).Elem()
}

func (o JobStagesResponseOutput) ToJobStagesResponseOutput() JobStagesResponseOutput {
	return o
}

func (o JobStagesResponseOutput) ToJobStagesResponseOutputWithContext(ctx context.Context) JobStagesResponseOutput {
	return o
}

func (o JobStagesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStagesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o JobStagesResponseOutput) ErrorDetails() JobErrorDetailsResponseArrayOutput {
	return o.ApplyT(func(v JobStagesResponse) []JobErrorDetailsResponse { return v.ErrorDetails }).(JobErrorDetailsResponseArrayOutput)
}

func (o JobStagesResponseOutput) JobStageDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v JobStagesResponse) interface{} { return v.JobStageDetails }).(pulumi.AnyOutput)
}

func (o JobStagesResponseOutput) StageName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStagesResponse) string { return v.StageName }).(pulumi.StringOutput)
}

func (o JobStagesResponseOutput) StageStatus() pulumi.StringOutput {
	return o.ApplyT(func(v JobStagesResponse) string { return v.StageStatus }).(pulumi.StringOutput)
}

func (o JobStagesResponseOutput) StageTime() pulumi.StringOutput {
	return o.ApplyT(func(v JobStagesResponse) string { return v.StageTime }).(pulumi.StringOutput)
}

type JobStagesResponseArrayOutput struct{ *pulumi.OutputState }

func (JobStagesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobStagesResponse)(nil)).Elem()
}

func (o JobStagesResponseArrayOutput) ToJobStagesResponseArrayOutput() JobStagesResponseArrayOutput {
	return o
}

func (o JobStagesResponseArrayOutput) ToJobStagesResponseArrayOutputWithContext(ctx context.Context) JobStagesResponseArrayOutput {
	return o
}

func (o JobStagesResponseArrayOutput) Index(i pulumi.IntInput) JobStagesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobStagesResponse {
		return vs[0].([]JobStagesResponse)[vs[1].(int)]
	}).(JobStagesResponseOutput)
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

type PackageShippingDetailsResponse struct {
	CarrierName string `pulumi:"carrierName"`
	TrackingId  string `pulumi:"trackingId"`
	TrackingUrl string `pulumi:"trackingUrl"`
}





type PackageShippingDetailsResponseInput interface {
	pulumi.Input

	ToPackageShippingDetailsResponseOutput() PackageShippingDetailsResponseOutput
	ToPackageShippingDetailsResponseOutputWithContext(context.Context) PackageShippingDetailsResponseOutput
}

type PackageShippingDetailsResponseArgs struct {
	CarrierName pulumi.StringInput `pulumi:"carrierName"`
	TrackingId  pulumi.StringInput `pulumi:"trackingId"`
	TrackingUrl pulumi.StringInput `pulumi:"trackingUrl"`
}

func (PackageShippingDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageShippingDetailsResponse)(nil)).Elem()
}

func (i PackageShippingDetailsResponseArgs) ToPackageShippingDetailsResponseOutput() PackageShippingDetailsResponseOutput {
	return i.ToPackageShippingDetailsResponseOutputWithContext(context.Background())
}

func (i PackageShippingDetailsResponseArgs) ToPackageShippingDetailsResponseOutputWithContext(ctx context.Context) PackageShippingDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageShippingDetailsResponseOutput)
}

type PackageShippingDetailsResponseOutput struct{ *pulumi.OutputState }

func (PackageShippingDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageShippingDetailsResponse)(nil)).Elem()
}

func (o PackageShippingDetailsResponseOutput) ToPackageShippingDetailsResponseOutput() PackageShippingDetailsResponseOutput {
	return o
}

func (o PackageShippingDetailsResponseOutput) ToPackageShippingDetailsResponseOutputWithContext(ctx context.Context) PackageShippingDetailsResponseOutput {
	return o
}

func (o PackageShippingDetailsResponseOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v PackageShippingDetailsResponse) string { return v.CarrierName }).(pulumi.StringOutput)
}

func (o PackageShippingDetailsResponseOutput) TrackingId() pulumi.StringOutput {
	return o.ApplyT(func(v PackageShippingDetailsResponse) string { return v.TrackingId }).(pulumi.StringOutput)
}

func (o PackageShippingDetailsResponseOutput) TrackingUrl() pulumi.StringOutput {
	return o.ApplyT(func(v PackageShippingDetailsResponse) string { return v.TrackingUrl }).(pulumi.StringOutput)
}

type Preferences struct {
	PreferredDataCenterRegion []string              `pulumi:"preferredDataCenterRegion"`
	TransportPreferences      *TransportPreferences `pulumi:"transportPreferences"`
}





type PreferencesInput interface {
	pulumi.Input

	ToPreferencesOutput() PreferencesOutput
	ToPreferencesOutputWithContext(context.Context) PreferencesOutput
}

type PreferencesArgs struct {
	PreferredDataCenterRegion pulumi.StringArrayInput      `pulumi:"preferredDataCenterRegion"`
	TransportPreferences      TransportPreferencesPtrInput `pulumi:"transportPreferences"`
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

func (o PreferencesOutput) PreferredDataCenterRegion() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Preferences) []string { return v.PreferredDataCenterRegion }).(pulumi.StringArrayOutput)
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

func (o PreferencesPtrOutput) PreferredDataCenterRegion() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Preferences) []string {
		if v == nil {
			return nil
		}
		return v.PreferredDataCenterRegion
	}).(pulumi.StringArrayOutput)
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
	PreferredDataCenterRegion []string                      `pulumi:"preferredDataCenterRegion"`
	TransportPreferences      *TransportPreferencesResponse `pulumi:"transportPreferences"`
}





type PreferencesResponseInput interface {
	pulumi.Input

	ToPreferencesResponseOutput() PreferencesResponseOutput
	ToPreferencesResponseOutputWithContext(context.Context) PreferencesResponseOutput
}

type PreferencesResponseArgs struct {
	PreferredDataCenterRegion pulumi.StringArrayInput              `pulumi:"preferredDataCenterRegion"`
	TransportPreferences      TransportPreferencesResponsePtrInput `pulumi:"transportPreferences"`
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

func (o PreferencesResponseOutput) PreferredDataCenterRegion() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PreferencesResponse) []string { return v.PreferredDataCenterRegion }).(pulumi.StringArrayOutput)
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

func (o PreferencesResponsePtrOutput) PreferredDataCenterRegion() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PreferencesResponse) []string {
		if v == nil {
			return nil
		}
		return v.PreferredDataCenterRegion
	}).(pulumi.StringArrayOutput)
}

func (o PreferencesResponsePtrOutput) TransportPreferences() TransportPreferencesResponsePtrOutput {
	return o.ApplyT(func(v *PreferencesResponse) *TransportPreferencesResponse {
		if v == nil {
			return nil
		}
		return v.TransportPreferences
	}).(TransportPreferencesResponsePtrOutput)
}

type ShareCredentialDetailsResponse struct {
	Password                 string   `pulumi:"password"`
	ShareName                string   `pulumi:"shareName"`
	ShareType                string   `pulumi:"shareType"`
	SupportedAccessProtocols []string `pulumi:"supportedAccessProtocols"`
	UserName                 string   `pulumi:"userName"`
}





type ShareCredentialDetailsResponseInput interface {
	pulumi.Input

	ToShareCredentialDetailsResponseOutput() ShareCredentialDetailsResponseOutput
	ToShareCredentialDetailsResponseOutputWithContext(context.Context) ShareCredentialDetailsResponseOutput
}

type ShareCredentialDetailsResponseArgs struct {
	Password                 pulumi.StringInput      `pulumi:"password"`
	ShareName                pulumi.StringInput      `pulumi:"shareName"`
	ShareType                pulumi.StringInput      `pulumi:"shareType"`
	SupportedAccessProtocols pulumi.StringArrayInput `pulumi:"supportedAccessProtocols"`
	UserName                 pulumi.StringInput      `pulumi:"userName"`
}

func (ShareCredentialDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareCredentialDetailsResponse)(nil)).Elem()
}

func (i ShareCredentialDetailsResponseArgs) ToShareCredentialDetailsResponseOutput() ShareCredentialDetailsResponseOutput {
	return i.ToShareCredentialDetailsResponseOutputWithContext(context.Background())
}

func (i ShareCredentialDetailsResponseArgs) ToShareCredentialDetailsResponseOutputWithContext(ctx context.Context) ShareCredentialDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareCredentialDetailsResponseOutput)
}





type ShareCredentialDetailsResponseArrayInput interface {
	pulumi.Input

	ToShareCredentialDetailsResponseArrayOutput() ShareCredentialDetailsResponseArrayOutput
	ToShareCredentialDetailsResponseArrayOutputWithContext(context.Context) ShareCredentialDetailsResponseArrayOutput
}

type ShareCredentialDetailsResponseArray []ShareCredentialDetailsResponseInput

func (ShareCredentialDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareCredentialDetailsResponse)(nil)).Elem()
}

func (i ShareCredentialDetailsResponseArray) ToShareCredentialDetailsResponseArrayOutput() ShareCredentialDetailsResponseArrayOutput {
	return i.ToShareCredentialDetailsResponseArrayOutputWithContext(context.Background())
}

func (i ShareCredentialDetailsResponseArray) ToShareCredentialDetailsResponseArrayOutputWithContext(ctx context.Context) ShareCredentialDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareCredentialDetailsResponseArrayOutput)
}

type ShareCredentialDetailsResponseOutput struct{ *pulumi.OutputState }

func (ShareCredentialDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareCredentialDetailsResponse)(nil)).Elem()
}

func (o ShareCredentialDetailsResponseOutput) ToShareCredentialDetailsResponseOutput() ShareCredentialDetailsResponseOutput {
	return o
}

func (o ShareCredentialDetailsResponseOutput) ToShareCredentialDetailsResponseOutputWithContext(ctx context.Context) ShareCredentialDetailsResponseOutput {
	return o
}

func (o ShareCredentialDetailsResponseOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v ShareCredentialDetailsResponse) string { return v.Password }).(pulumi.StringOutput)
}

func (o ShareCredentialDetailsResponseOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v ShareCredentialDetailsResponse) string { return v.ShareName }).(pulumi.StringOutput)
}

func (o ShareCredentialDetailsResponseOutput) ShareType() pulumi.StringOutput {
	return o.ApplyT(func(v ShareCredentialDetailsResponse) string { return v.ShareType }).(pulumi.StringOutput)
}

func (o ShareCredentialDetailsResponseOutput) SupportedAccessProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ShareCredentialDetailsResponse) []string { return v.SupportedAccessProtocols }).(pulumi.StringArrayOutput)
}

func (o ShareCredentialDetailsResponseOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ShareCredentialDetailsResponse) string { return v.UserName }).(pulumi.StringOutput)
}

type ShareCredentialDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (ShareCredentialDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareCredentialDetailsResponse)(nil)).Elem()
}

func (o ShareCredentialDetailsResponseArrayOutput) ToShareCredentialDetailsResponseArrayOutput() ShareCredentialDetailsResponseArrayOutput {
	return o
}

func (o ShareCredentialDetailsResponseArrayOutput) ToShareCredentialDetailsResponseArrayOutputWithContext(ctx context.Context) ShareCredentialDetailsResponseArrayOutput {
	return o
}

func (o ShareCredentialDetailsResponseArrayOutput) Index(i pulumi.IntInput) ShareCredentialDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ShareCredentialDetailsResponse {
		return vs[0].([]ShareCredentialDetailsResponse)[vs[1].(int)]
	}).(ShareCredentialDetailsResponseOutput)
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
	PostalCode      pulumi.StringInput    `pulumi:"postalCode"`
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

func (o ShippingAddressOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingAddress) string { return v.PostalCode }).(pulumi.StringOutput)
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
	PostalCode      pulumi.StringInput    `pulumi:"postalCode"`
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

func (o ShippingAddressResponseOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingAddressResponse) string { return v.PostalCode }).(pulumi.StringOutput)
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

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
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

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
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

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Family      *string `pulumi:"family"`
	Name        string  `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Family      pulumi.StringPtrInput `pulumi:"family"`
	Name        pulumi.StringInput    `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
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

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
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

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
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

type UnencryptedCredentialsResponse struct {
	JobName    string      `pulumi:"jobName"`
	JobSecrets interface{} `pulumi:"jobSecrets"`
}





type UnencryptedCredentialsResponseInput interface {
	pulumi.Input

	ToUnencryptedCredentialsResponseOutput() UnencryptedCredentialsResponseOutput
	ToUnencryptedCredentialsResponseOutputWithContext(context.Context) UnencryptedCredentialsResponseOutput
}

type UnencryptedCredentialsResponseArgs struct {
	JobName    pulumi.StringInput `pulumi:"jobName"`
	JobSecrets pulumi.Input       `pulumi:"jobSecrets"`
}

func (UnencryptedCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UnencryptedCredentialsResponse)(nil)).Elem()
}

func (i UnencryptedCredentialsResponseArgs) ToUnencryptedCredentialsResponseOutput() UnencryptedCredentialsResponseOutput {
	return i.ToUnencryptedCredentialsResponseOutputWithContext(context.Background())
}

func (i UnencryptedCredentialsResponseArgs) ToUnencryptedCredentialsResponseOutputWithContext(ctx context.Context) UnencryptedCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnencryptedCredentialsResponseOutput)
}





type UnencryptedCredentialsResponseArrayInput interface {
	pulumi.Input

	ToUnencryptedCredentialsResponseArrayOutput() UnencryptedCredentialsResponseArrayOutput
	ToUnencryptedCredentialsResponseArrayOutputWithContext(context.Context) UnencryptedCredentialsResponseArrayOutput
}

type UnencryptedCredentialsResponseArray []UnencryptedCredentialsResponseInput

func (UnencryptedCredentialsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UnencryptedCredentialsResponse)(nil)).Elem()
}

func (i UnencryptedCredentialsResponseArray) ToUnencryptedCredentialsResponseArrayOutput() UnencryptedCredentialsResponseArrayOutput {
	return i.ToUnencryptedCredentialsResponseArrayOutputWithContext(context.Background())
}

func (i UnencryptedCredentialsResponseArray) ToUnencryptedCredentialsResponseArrayOutputWithContext(ctx context.Context) UnencryptedCredentialsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnencryptedCredentialsResponseArrayOutput)
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

func init() {
	pulumi.RegisterOutputType(AccountCredentialDetailsResponseOutput{})
	pulumi.RegisterOutputType(AccountCredentialDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplianceNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ApplianceNetworkConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(ContactDetailsOutput{})
	pulumi.RegisterOutputType(ContactDetailsResponseOutput{})
	pulumi.RegisterOutputType(CopyProgressResponseOutput{})
	pulumi.RegisterOutputType(CopyProgressResponseArrayOutput{})
	pulumi.RegisterOutputType(DataBoxAccountCopyLogDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxDiskCopyLogDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxDiskCopyProgressResponseOutput{})
	pulumi.RegisterOutputType(DataBoxDiskCopyProgressResponseArrayOutput{})
	pulumi.RegisterOutputType(DataBoxDiskJobDetailsOutput{})
	pulumi.RegisterOutputType(DataBoxDiskJobDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxDiskJobSecretsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxHeavyAccountCopyLogDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxHeavyJobDetailsOutput{})
	pulumi.RegisterOutputType(DataBoxHeavyJobDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxHeavyJobSecretsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxHeavySecretResponseOutput{})
	pulumi.RegisterOutputType(DataBoxHeavySecretResponseArrayOutput{})
	pulumi.RegisterOutputType(DataBoxJobDetailsOutput{})
	pulumi.RegisterOutputType(DataBoxJobDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxSecretResponseOutput{})
	pulumi.RegisterOutputType(DataBoxSecretResponseArrayOutput{})
	pulumi.RegisterOutputType(DataboxJobSecretsResponseOutput{})
	pulumi.RegisterOutputType(DcAccessSecurityCodeResponseOutput{})
	pulumi.RegisterOutputType(DcAccessSecurityCodeResponsePtrOutput{})
	pulumi.RegisterOutputType(DestinationManagedDiskDetailsOutput{})
	pulumi.RegisterOutputType(DestinationManagedDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(DestinationStorageAccountDetailsOutput{})
	pulumi.RegisterOutputType(DestinationStorageAccountDetailsResponseOutput{})
	pulumi.RegisterOutputType(DiskSecretResponseOutput{})
	pulumi.RegisterOutputType(DiskSecretResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoPtrOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoResponseOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(JobErrorDetailsResponseOutput{})
	pulumi.RegisterOutputType(JobErrorDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(JobStagesResponseOutput{})
	pulumi.RegisterOutputType(JobStagesResponseArrayOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceArrayOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceResponseOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(PackageShippingDetailsResponseOutput{})
	pulumi.RegisterOutputType(PreferencesOutput{})
	pulumi.RegisterOutputType(PreferencesPtrOutput{})
	pulumi.RegisterOutputType(PreferencesResponseOutput{})
	pulumi.RegisterOutputType(PreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(ShareCredentialDetailsResponseOutput{})
	pulumi.RegisterOutputType(ShareCredentialDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(ShippingAddressOutput{})
	pulumi.RegisterOutputType(ShippingAddressResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(TransportPreferencesOutput{})
	pulumi.RegisterOutputType(TransportPreferencesPtrOutput{})
	pulumi.RegisterOutputType(TransportPreferencesResponseOutput{})
	pulumi.RegisterOutputType(TransportPreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(UnencryptedCredentialsResponseOutput{})
	pulumi.RegisterOutputType(UnencryptedCredentialsResponseArrayOutput{})
}
