package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewMigrateAgent(ctx, "migrateAgent", &migrate.MigrateAgentArgs{
			AgentName:            pulumi.String("l"),
			ModernizeProjectName: pulumi.String("rq1yec"),
			Properties: &migrate.MigrateAgentModelPropertiesArgs{
				AuthenticationIdentity: &migrate.IdentityModelArgs{
					AadAuthority:  pulumi.String("isbicanvfefdaci"),
					ApplicationId: pulumi.String("dibfqwjrnzikktkwe"),
					Audience:      pulumi.String("yrfxszjhkczoyfi"),
					ObjectId:      pulumi.String("xfhhdosr"),
					TenantId:      pulumi.String("uwceuawplakwjswbvllffbsz"),
				},
				CustomProperties: nil,
				MachineId:        pulumi.String("sihoniqzqfz"),
				MachineName:      pulumi.String("glhejppirkiamgxxro"),
			},
			ResourceGroupName: pulumi.String("rgmigrateEngine"),
			Tags: pulumi.StringMap{
				"key5560": pulumi.String("jgffrfcgjrm"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
