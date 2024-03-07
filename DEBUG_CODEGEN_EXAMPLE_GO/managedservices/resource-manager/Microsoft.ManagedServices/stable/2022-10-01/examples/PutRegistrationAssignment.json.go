package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managedservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managedservices.NewRegistrationAssignment(ctx, "registrationAssignment", &managedservices.RegistrationAssignmentArgs{
			Properties: &managedservices.RegistrationAssignmentPropertiesArgs{
				RegistrationDefinitionId: pulumi.String("/subscriptions/0afefe50-734e-4610-8a82-a144ahf49dea/providers/Microsoft.ManagedServices/registrationDefinitions/26c128c2-fefa-4340-9bb1-6e081c90ada2"),
			},
			RegistrationAssignmentId: pulumi.String("26c128c2-fefa-4340-9bb1-6e081c90ada2"),
			Scope:                    pulumi.String("subscription/0afefe50-734e-4610-8a82-a144ahf49dea"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
