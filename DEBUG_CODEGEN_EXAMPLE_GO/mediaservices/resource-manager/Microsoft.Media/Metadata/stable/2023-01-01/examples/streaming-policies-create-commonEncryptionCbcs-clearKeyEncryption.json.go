package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := media.NewStreamingPolicy(ctx, "streamingPolicy", &media.StreamingPolicyArgs{
AccountName: pulumi.String("contosomedia"),
CommonEncryptionCbcs: media.CommonEncryptionCbcsResponse{
ClearKeyEncryptionConfiguration: &media.ClearKeyEncryptionConfigurationArgs{
CustomKeysAcquisitionUrlTemplate: pulumi.String("https://contoso.com/{AlternativeMediaId}/clearkey/"),
},
ContentKeys: interface{}{
DefaultKey: &media.DefaultKeyArgs{
Label: pulumi.String("cbcsDefaultKey"),
},
},
EnabledProtocols: &media.EnabledProtocolsArgs{
Dash: pulumi.Bool(false),
Download: pulumi.Bool(false),
Hls: pulumi.Bool(true),
SmoothStreaming: pulumi.Bool(false),
},
},
DefaultContentKeyPolicyName: pulumi.String("PolicyWithMultipleOptions"),
ResourceGroupName: pulumi.String("contosorg"),
StreamingPolicyName: pulumi.String("UserCreatedSecureStreamingPolicyWithCommonEncryptionCbcsOnly"),
})
if err != nil {
return err
}
return nil
})
}
