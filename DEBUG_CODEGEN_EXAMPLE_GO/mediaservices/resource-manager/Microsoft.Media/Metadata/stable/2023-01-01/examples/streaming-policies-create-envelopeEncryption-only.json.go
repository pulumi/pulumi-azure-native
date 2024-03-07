package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := media.NewStreamingPolicy(ctx, "streamingPolicy", &media.StreamingPolicyArgs{
AccountName: pulumi.String("contosomedia"),
DefaultContentKeyPolicyName: pulumi.String("PolicyWithClearKeyOptionAndTokenRestriction"),
EnvelopeEncryption: media.EnvelopeEncryptionResponse{
ContentKeys: interface{}{
DefaultKey: &media.DefaultKeyArgs{
Label: pulumi.String("aesDefaultKey"),
},
},
CustomKeyAcquisitionUrlTemplate: pulumi.String("https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}"),
EnabledProtocols: &media.EnabledProtocolsArgs{
Dash: pulumi.Bool(true),
Download: pulumi.Bool(false),
Hls: pulumi.Bool(true),
SmoothStreaming: pulumi.Bool(true),
},
},
ResourceGroupName: pulumi.String("contosorg"),
StreamingPolicyName: pulumi.String("UserCreatedSecureStreamingPolicyWithEnvelopeEncryptionOnly"),
})
if err != nil {
return err
}
return nil
})
}
