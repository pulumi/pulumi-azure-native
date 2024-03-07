package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := logic.NewIntegrationAccountAgreement(ctx, "integrationAccountAgreement", &logic.IntegrationAccountAgreementArgs{
AgreementName: pulumi.String("testAgreement"),
AgreementType: logic.AgreementTypeAS2,
Content: logic.AgreementContentResponse{
AS2: interface{}{
ReceiveAgreement: interface{}{
ProtocolSettings: interface{}{
AcknowledgementConnectionSettings: &logic.AS2AcknowledgementConnectionSettingsArgs{
IgnoreCertificateNameMismatch: pulumi.Bool(true),
KeepHttpConnectionAlive: pulumi.Bool(true),
SupportHttpStatusCodeContinue: pulumi.Bool(true),
UnfoldHttpHeaders: pulumi.Bool(true),
},
EnvelopeSettings: &logic.AS2EnvelopeSettingsArgs{
AutogenerateFileName: pulumi.Bool(true),
FileNameTemplate: pulumi.String("Test"),
MessageContentType: pulumi.String("text/plain"),
SuspendMessageOnFileNameGenerationError: pulumi.Bool(true),
TransmitFileNameInMimeHeader: pulumi.Bool(true),
},
ErrorSettings: &logic.AS2ErrorSettingsArgs{
ResendIfMDNNotReceived: pulumi.Bool(true),
SuspendDuplicateMessage: pulumi.Bool(true),
},
MdnSettings: &logic.AS2MdnSettingsArgs{
DispositionNotificationTo: pulumi.String("http://tempuri.org"),
MdnText: pulumi.String("Sample"),
MicHashingAlgorithm: pulumi.String("SHA1"),
NeedMDN: pulumi.Bool(true),
ReceiptDeliveryUrl: pulumi.String("http://tempuri.org"),
SendInboundMDNToMessageBox: pulumi.Bool(true),
SendMDNAsynchronously: pulumi.Bool(true),
SignMDN: pulumi.Bool(true),
SignOutboundMDNIfOptional: pulumi.Bool(true),
},
MessageConnectionSettings: &logic.AS2MessageConnectionSettingsArgs{
IgnoreCertificateNameMismatch: pulumi.Bool(true),
KeepHttpConnectionAlive: pulumi.Bool(true),
SupportHttpStatusCodeContinue: pulumi.Bool(true),
UnfoldHttpHeaders: pulumi.Bool(true),
},
SecuritySettings: &logic.AS2SecuritySettingsArgs{
EnableNRRForInboundDecodedMessages: pulumi.Bool(true),
EnableNRRForInboundEncodedMessages: pulumi.Bool(true),
EnableNRRForInboundMDN: pulumi.Bool(true),
EnableNRRForOutboundDecodedMessages: pulumi.Bool(true),
EnableNRRForOutboundEncodedMessages: pulumi.Bool(true),
EnableNRRForOutboundMDN: pulumi.Bool(true),
OverrideGroupSigningCertificate: pulumi.Bool(false),
},
ValidationSettings: &logic.AS2ValidationSettingsArgs{
CheckCertificateRevocationListOnReceive: pulumi.Bool(true),
CheckCertificateRevocationListOnSend: pulumi.Bool(true),
CheckDuplicateMessage: pulumi.Bool(true),
CompressMessage: pulumi.Bool(true),
EncryptMessage: pulumi.Bool(false),
EncryptionAlgorithm: pulumi.String("AES128"),
InterchangeDuplicatesValidityDays: pulumi.Int(100),
OverrideMessageProperties: pulumi.Bool(true),
SignMessage: pulumi.Bool(false),
},
},
ReceiverBusinessIdentity: &logic.BusinessIdentityArgs{
Qualifier: pulumi.String("ZZ"),
Value: pulumi.String("ZZ"),
},
SenderBusinessIdentity: &logic.BusinessIdentityArgs{
Qualifier: pulumi.String("AA"),
Value: pulumi.String("AA"),
},
},
SendAgreement: interface{}{
ProtocolSettings: interface{}{
AcknowledgementConnectionSettings: &logic.AS2AcknowledgementConnectionSettingsArgs{
IgnoreCertificateNameMismatch: pulumi.Bool(true),
KeepHttpConnectionAlive: pulumi.Bool(true),
SupportHttpStatusCodeContinue: pulumi.Bool(true),
UnfoldHttpHeaders: pulumi.Bool(true),
},
EnvelopeSettings: &logic.AS2EnvelopeSettingsArgs{
AutogenerateFileName: pulumi.Bool(true),
FileNameTemplate: pulumi.String("Test"),
MessageContentType: pulumi.String("text/plain"),
SuspendMessageOnFileNameGenerationError: pulumi.Bool(true),
TransmitFileNameInMimeHeader: pulumi.Bool(true),
},
ErrorSettings: &logic.AS2ErrorSettingsArgs{
ResendIfMDNNotReceived: pulumi.Bool(true),
SuspendDuplicateMessage: pulumi.Bool(true),
},
MdnSettings: &logic.AS2MdnSettingsArgs{
DispositionNotificationTo: pulumi.String("http://tempuri.org"),
MdnText: pulumi.String("Sample"),
MicHashingAlgorithm: pulumi.String("SHA1"),
NeedMDN: pulumi.Bool(true),
ReceiptDeliveryUrl: pulumi.String("http://tempuri.org"),
SendInboundMDNToMessageBox: pulumi.Bool(true),
SendMDNAsynchronously: pulumi.Bool(true),
SignMDN: pulumi.Bool(true),
SignOutboundMDNIfOptional: pulumi.Bool(true),
},
MessageConnectionSettings: &logic.AS2MessageConnectionSettingsArgs{
IgnoreCertificateNameMismatch: pulumi.Bool(true),
KeepHttpConnectionAlive: pulumi.Bool(true),
SupportHttpStatusCodeContinue: pulumi.Bool(true),
UnfoldHttpHeaders: pulumi.Bool(true),
},
SecuritySettings: &logic.AS2SecuritySettingsArgs{
EnableNRRForInboundDecodedMessages: pulumi.Bool(true),
EnableNRRForInboundEncodedMessages: pulumi.Bool(true),
EnableNRRForInboundMDN: pulumi.Bool(true),
EnableNRRForOutboundDecodedMessages: pulumi.Bool(true),
EnableNRRForOutboundEncodedMessages: pulumi.Bool(true),
EnableNRRForOutboundMDN: pulumi.Bool(true),
OverrideGroupSigningCertificate: pulumi.Bool(false),
},
ValidationSettings: &logic.AS2ValidationSettingsArgs{
CheckCertificateRevocationListOnReceive: pulumi.Bool(true),
CheckCertificateRevocationListOnSend: pulumi.Bool(true),
CheckDuplicateMessage: pulumi.Bool(true),
CompressMessage: pulumi.Bool(true),
EncryptMessage: pulumi.Bool(false),
EncryptionAlgorithm: pulumi.String("AES128"),
InterchangeDuplicatesValidityDays: pulumi.Int(100),
OverrideMessageProperties: pulumi.Bool(true),
SignMessage: pulumi.Bool(false),
},
},
ReceiverBusinessIdentity: &logic.BusinessIdentityArgs{
Qualifier: pulumi.String("AA"),
Value: pulumi.String("AA"),
},
SenderBusinessIdentity: &logic.BusinessIdentityArgs{
Qualifier: pulumi.String("ZZ"),
Value: pulumi.String("ZZ"),
},
},
},
},
GuestIdentity: &logic.BusinessIdentityArgs{
Qualifier: pulumi.String("AA"),
Value: pulumi.String("AA"),
},
GuestPartner: pulumi.String("GuestPartner"),
HostIdentity: &logic.BusinessIdentityArgs{
Qualifier: pulumi.String("ZZ"),
Value: pulumi.String("ZZ"),
},
HostPartner: pulumi.String("HostPartner"),
IntegrationAccountName: pulumi.String("testIntegrationAccount"),
Location: pulumi.String("westus"),
Metadata: nil,
ResourceGroupName: pulumi.String("testResourceGroup"),
Tags: pulumi.StringMap{
"IntegrationAccountAgreement": pulumi.String("<IntegrationAccountAgreementName>"),
},
})
if err != nil {
return err
}
return nil
})
}
