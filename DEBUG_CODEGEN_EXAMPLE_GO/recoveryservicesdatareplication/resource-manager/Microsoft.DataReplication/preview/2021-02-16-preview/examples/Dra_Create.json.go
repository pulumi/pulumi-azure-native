package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datareplication/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datareplication.NewDra(ctx, "dra", &datareplication.DraArgs{
			FabricAgentName: pulumi.String("M"),
			FabricName:      pulumi.String("wPR"),
			Properties: &datareplication.DraModelPropertiesArgs{
				AuthenticationIdentity: &datareplication.IdentityModelArgs{
					AadAuthority:  pulumi.String("bubwwbowfhdmujrt"),
					ApplicationId: pulumi.String("cwktzrwajuvfyyymfstpey"),
					Audience:      pulumi.String("dkjobanyqgzenivyxhvavottpc"),
					ObjectId:      pulumi.String("khsiaqfbpuhp"),
					TenantId:      pulumi.String("joclkkdovixwapephhxaqtefubhhmq"),
				},
				CustomProperties: nil,
				MachineId:        pulumi.String("envzcoijbqhtrpncbjbhk"),
				MachineName:      pulumi.String("y"),
				ResourceAccessIdentity: &datareplication.IdentityModelArgs{
					AadAuthority:  pulumi.String("bubwwbowfhdmujrt"),
					ApplicationId: pulumi.String("cwktzrwajuvfyyymfstpey"),
					Audience:      pulumi.String("dkjobanyqgzenivyxhvavottpc"),
					ObjectId:      pulumi.String("khsiaqfbpuhp"),
					TenantId:      pulumi.String("joclkkdovixwapephhxaqtefubhhmq"),
				},
			},
			ResourceGroupName: pulumi.String("rgrecoveryservicesdatareplication"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
