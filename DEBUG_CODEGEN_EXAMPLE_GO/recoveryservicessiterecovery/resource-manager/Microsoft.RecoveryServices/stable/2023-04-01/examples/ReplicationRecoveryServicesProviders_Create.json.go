package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewReplicationRecoveryServicesProvider(ctx, "replicationRecoveryServicesProvider", &recoveryservices.ReplicationRecoveryServicesProviderArgs{
			FabricName: pulumi.String("vmwarefabric1"),
			Properties: recoveryservices.RecoveryServicesProviderPropertiesResponse{
				AuthenticationIdentityInput: &recoveryservices.IdentityProviderInputArgs{
					AadAuthority:  pulumi.String("https://login.microsoftonline.com"),
					ApplicationId: pulumi.String("f66fce08-c0c6-47a1-beeb-0ede5ea94f90"),
					Audience:      pulumi.String("https://microsoft.onmicrosoft.com/cf19e349-644c-4c6a-bcae-9c8f35357874"),
					ObjectId:      pulumi.String("141360b8-5686-4240-a027-5e24e6affeba"),
					TenantId:      pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
				},
				MachineName: pulumi.String("vmwareprovider1"),
				ResourceAccessIdentityInput: &recoveryservices.IdentityProviderInputArgs{
					AadAuthority:  pulumi.String("https://login.microsoftonline.com"),
					ApplicationId: pulumi.String("f66fce08-c0c6-47a1-beeb-0ede5ea94f90"),
					Audience:      pulumi.String("https://microsoft.onmicrosoft.com/cf19e349-644c-4c6a-bcae-9c8f35357874"),
					ObjectId:      pulumi.String("141360b8-5686-4240-a027-5e24e6affeba"),
					TenantId:      pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
				},
			},
			ProviderName:      pulumi.String("vmwareprovider1"),
			ResourceGroupName: pulumi.String("resourcegroup1"),
			ResourceName:      pulumi.String("migrationvault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
