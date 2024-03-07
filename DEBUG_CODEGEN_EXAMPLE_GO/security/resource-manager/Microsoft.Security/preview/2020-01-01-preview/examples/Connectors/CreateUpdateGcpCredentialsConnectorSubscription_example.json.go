package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewConnector(ctx, "connector", &security.ConnectorArgs{
			AuthenticationDetails: security.GcpCredentialsDetailsProperties{
				AuthProviderX509CertUrl: "https://www.googleapis.com/oauth2/v1/certs",
				AuthUri:                 "https://accounts.google.com/o/oauth2/auth",
				AuthenticationType:      "gcpCredentials",
				ClientEmail:             "asc-135@asc-project-1234.iam.gserviceaccount.com",
				ClientId:                "105889053725632919854",
				ClientX509CertUrl:       "https://www.googleapis.com/robot/v1/metadata/x509/asc-135%40asc-project-1234.iam.gserviceaccount.com",
				OrganizationId:          "AscDemoOrg",
				PrivateKey:              "******",
				PrivateKeyId:            "6efg587hra2568as34d22326b044cc20dc2af",
				ProjectId:               "asc-project-1234",
				TokenUri:                "https://oauth2.googleapis.com/token",
				Type:                    "service_account",
			},
			ConnectorName: pulumi.String("gcp_dev"),
			HybridComputeSettings: &security.HybridComputeSettingsPropertiesArgs{
				AutoProvision: pulumi.String("Off"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
