package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewSecurityConnectorApplication(ctx, "securityConnectorApplication", &security.SecurityConnectorApplicationArgs{
			ApplicationId:         pulumi.String("ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
			Description:           pulumi.String("An application on critical GCP recommendations"),
			DisplayName:           pulumi.String("GCP Admin's application"),
			ResourceGroupName:     pulumi.String("gcpResourceGroup"),
			SecurityConnectorName: pulumi.String("gcpconnector"),
			SourceResourceType:    pulumi.String("Assessments"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
