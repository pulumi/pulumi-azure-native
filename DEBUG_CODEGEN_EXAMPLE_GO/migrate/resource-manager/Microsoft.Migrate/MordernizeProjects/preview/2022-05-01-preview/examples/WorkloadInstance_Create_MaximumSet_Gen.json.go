package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := migrate.NewWorkloadInstance(ctx, "workloadInstance", &migrate.WorkloadInstanceArgs{
ModernizeProjectName: pulumi.String("mx8"),
Properties: migrate.WorkloadInstanceModelPropertiesResponse{
CustomProperties: interface{}{
InstanceType: pulumi.String("IISWorkload"),
WebAppArmId: pulumi.String("xseseqsrzdiga"),
WebAppSiteName: pulumi.String("mirgzmy"),
},
DisplayName: pulumi.String("juoorbubchvk"),
MasterSiteName: pulumi.String("ubks"),
MigrateAgentId: pulumi.String("aqgzsxqbk"),
Name: pulumi.String("wonkuhgsafzviuwqerzdmme"),
SourceName: pulumi.String("weuxcqzwpeyzsjhdgqflhxlwjhbz"),
SourcePlatform: pulumi.String("eh"),
},
ResourceGroupName: pulumi.String("rgmigrateEngine"),
Tags: pulumi.StringMap{
"key2836": pulumi.String("biqip"),
},
WorkloadInstanceName: pulumi.String("m"),
})
if err != nil {
return err
}
return nil
})
}
