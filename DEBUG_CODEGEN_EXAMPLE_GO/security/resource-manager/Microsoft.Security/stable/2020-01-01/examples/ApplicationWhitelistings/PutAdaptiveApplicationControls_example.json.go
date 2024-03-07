package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewAdaptiveApplicationControl(ctx, "adaptiveApplicationControl", &security.AdaptiveApplicationControlArgs{
			AscLocation:     pulumi.String("centralus"),
			EnforcementMode: pulumi.String("Audit"),
			GroupName:       pulumi.String("ERELGROUP1"),
			PathRecommendations: security.PathRecommendationArray{
				&security.PathRecommendationArgs{
					Action:              pulumi.String("Recommended"),
					Common:              pulumi.Bool(true),
					ConfigurationStatus: pulumi.String("Configured"),
					FileType:            pulumi.String("Exe"),
					Path:                pulumi.String("[Exe] O=MICROSOFT CORPORATION, L=REDMOND, S=WASHINGTON, C=US\\*\\*\\0.0.0.0"),
					PublisherInfo: &security.PublisherInfoArgs{
						BinaryName:    pulumi.String("*"),
						ProductName:   pulumi.String("*"),
						PublisherName: pulumi.String("O=MICROSOFT CORPORATION, L=REDMOND, S=WASHINGTON, C=US"),
						Version:       pulumi.String("0.0.0.0"),
					},
					Type: pulumi.String("PublisherSignature"),
					UserSids: pulumi.StringArray{
						pulumi.String("S-1-1-0"),
					},
					Usernames: security.UserRecommendationArray{
						&security.UserRecommendationArgs{
							RecommendationAction: pulumi.String("Recommended"),
							Username:             pulumi.String("Everyone"),
						},
					},
				},
				&security.PathRecommendationArgs{
					Action:              pulumi.String("Recommended"),
					Common:              pulumi.Bool(true),
					ConfigurationStatus: pulumi.String("Configured"),
					FileType:            pulumi.String("Exe"),
					Path:                pulumi.String("%OSDRIVE%\\WINDOWSAZURE\\SECAGENT\\WASECAGENTPROV.EXE"),
					PublisherInfo: &security.PublisherInfoArgs{
						BinaryName:    pulumi.String("*"),
						ProductName:   pulumi.String("MICROSOFT® COREXT"),
						PublisherName: pulumi.String("CN=MICROSOFT AZURE DEPENDENCY CODE SIGN"),
						Version:       pulumi.String("0.0.0.0"),
					},
					Type: pulumi.String("ProductSignature"),
					UserSids: pulumi.StringArray{
						pulumi.String("S-1-1-0"),
					},
					Usernames: security.UserRecommendationArray{
						&security.UserRecommendationArgs{
							RecommendationAction: pulumi.String("Recommended"),
							Username:             pulumi.String("NT AUTHORITY\\SYSTEM"),
						},
					},
				},
				&security.PathRecommendationArgs{
					Action:              pulumi.String("Recommended"),
					Common:              pulumi.Bool(true),
					ConfigurationStatus: pulumi.String("Configured"),
					FileType:            pulumi.String("Exe"),
					Path:                pulumi.String("%OSDRIVE%\\WINDOWSAZURE\\PACKAGES_201973_7415\\COLLECTGUESTLOGS.EXE"),
					PublisherInfo: &security.PublisherInfoArgs{
						BinaryName:    pulumi.String("*"),
						ProductName:   pulumi.String("*"),
						PublisherName: pulumi.String("CN=MICROSOFT AZURE DEPENDENCY CODE SIGN"),
						Version:       pulumi.String("0.0.0.0"),
					},
					Type: pulumi.String("PublisherSignature"),
					UserSids: pulumi.StringArray{
						pulumi.String("S-1-1-0"),
					},
					Usernames: security.UserRecommendationArray{
						&security.UserRecommendationArgs{
							RecommendationAction: pulumi.String("Recommended"),
							Username:             pulumi.String("NT AUTHORITY\\SYSTEM"),
						},
					},
				},
				&security.PathRecommendationArgs{
					Action: pulumi.String("Add"),
					Common: pulumi.Bool(true),
					Path:   pulumi.String("C:\\directory\\file.exe"),
					Type:   pulumi.String("File"),
				},
			},
			ProtectionMode: &security.ProtectionModeArgs{
				Exe:    pulumi.String("Audit"),
				Msi:    pulumi.String("None"),
				Script: pulumi.String("None"),
			},
			VmRecommendations: security.VmRecommendationArray{
				&security.VmRecommendationArgs{
					ConfigurationStatus:  pulumi.String("Configured"),
					EnforcementSupport:   pulumi.String("Supported"),
					RecommendationAction: pulumi.String("Recommended"),
					ResourceId:           pulumi.String("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/erelh-stable/providers/microsoft.compute/virtualmachines/erelh-16090"),
				},
				&security.VmRecommendationArgs{
					ConfigurationStatus:  pulumi.String("Configured"),
					EnforcementSupport:   pulumi.String("Supported"),
					RecommendationAction: pulumi.String("Recommended"),
					ResourceId:           pulumi.String("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/matanvs/providers/microsoft.compute/virtualmachines/matanvs19"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
