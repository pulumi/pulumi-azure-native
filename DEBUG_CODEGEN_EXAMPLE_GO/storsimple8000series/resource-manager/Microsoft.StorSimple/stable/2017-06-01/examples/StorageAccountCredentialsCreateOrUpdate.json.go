package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storsimple.NewStorageAccountCredential(ctx, "storageAccountCredential", &storsimple.StorageAccountCredentialArgs{
			AccessKey: &storsimple.AsymmetricEncryptedSecretArgs{
				EncryptionAlgorithm:      storsimple.EncryptionAlgorithm_RSAES_PKCS1_v_1_5,
				EncryptionCertThumbprint: pulumi.String("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
				Value:                    pulumi.String("ATuJSkmrFk4h8r1jrZ4nd3nthLSddcguEO5QLO/NECUtTuB9kL4dNv3/jC4WOvFkeVr3x1UvfhlIeMmJBF1SMr6hR1JzD0xNU/TtQqUeXN7V3jk7I+2l67P9StuHWR6OMd3XOLwvznxOEQtEWpweDiobZU1ZiY03WafcGZFpV5j6tEoHeopoZ1J/GhPtkYmx+TqxzUN6qnir5rP3NSYiZciImP/qu8U9yUV/xpVRv39KvFc2Yr5SpKpMMRUj55XW10UnPer63M6KovF8X9Wi/fNnrZAs1Esl5XddZETGrW/e5B++VMJ6w0Q/uvPR+UBwrOU0804l0SzwdIe3qVVd0Q=="),
			},
			EndPoint:                     pulumi.String("blob.core.windows.net"),
			ManagerName:                  pulumi.String("ManagerForSDKTest1"),
			ResourceGroupName:            pulumi.String("ResourceGroupForSDKTest"),
			SslStatus:                    storsimple.SslStatusEnabled,
			StorageAccountCredentialName: pulumi.String("SACForTest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
