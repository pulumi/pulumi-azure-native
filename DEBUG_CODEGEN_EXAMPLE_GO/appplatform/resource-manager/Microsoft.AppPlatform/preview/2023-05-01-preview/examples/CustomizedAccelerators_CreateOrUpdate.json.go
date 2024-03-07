package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewCustomizedAccelerator(ctx, "customizedAccelerator", &appplatform.CustomizedAcceleratorArgs{
			ApplicationAcceleratorName: pulumi.String("default"),
			CustomizedAcceleratorName:  pulumi.String("acc-name"),
			Properties: &appplatform.CustomizedAcceleratorPropertiesArgs{
				AcceleratorTags: pulumi.StringArray{
					pulumi.String("tag-a"),
					pulumi.String("tag-b"),
				},
				Description: pulumi.String("acc-desc"),
				DisplayName: pulumi.String("acc-name"),
				GitRepository: &appplatform.AcceleratorGitRepositoryArgs{
					AuthSetting: appplatform.AcceleratorSshSetting{
						AuthType:         "SSH",
						HostKey:          "git-auth-hostkey",
						HostKeyAlgorithm: "git-auth-algorithm",
						PrivateKey:       "git-auth-privatekey",
					},
					Branch:            pulumi.String("git-branch"),
					Commit:            pulumi.String("12345"),
					GitTag:            pulumi.String("git-tag"),
					IntervalInSeconds: pulumi.Int(70),
					Url:               pulumi.String("git-url"),
				},
				IconUrl: pulumi.String("acc-icon"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
			Sku: &appplatform.SkuArgs{
				Capacity: pulumi.Int(2),
				Name:     pulumi.String("E0"),
				Tier:     pulumi.String("Enterprise"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
