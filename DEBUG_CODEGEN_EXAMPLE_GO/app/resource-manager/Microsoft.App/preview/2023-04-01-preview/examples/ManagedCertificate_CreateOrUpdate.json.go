package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewManagedCertificate(ctx, "managedCertificate", &app.ManagedCertificateArgs{
			EnvironmentName:        pulumi.String("testcontainerenv"),
			Location:               pulumi.String("East US"),
			ManagedCertificateName: pulumi.String("certificate-firendly-name"),
			Properties: &app.ManagedCertificatePropertiesArgs{
				DomainControlValidation: pulumi.String("CNAME"),
				SubjectName:             pulumi.String("my-subject-name.company.country.net"),
			},
			ResourceGroupName: pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
