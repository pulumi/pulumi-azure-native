package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := migrate.NewWorkloadDeployment(ctx, "workloadDeployment", &migrate.WorkloadDeploymentArgs{
ModernizeProjectName: pulumi.String("l6r8"),
Properties: &migrate.WorkloadDeploymentModelPropertiesArgs{
CustomProperties: interface{}{
InstanceType: pulumi.String("IISAKSWorkloadDeployment"),
},
DisplayName: pulumi.String("wqe"),
TargetPlatform: pulumi.String("AzureKubernetesService"),
WorkloadInstanceProperties: &migrate.WorkloadInstanceModelPropertiesArgs{
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
},
ResourceGroupName: pulumi.String("rgmigrateEngine"),
Tags: pulumi.StringMap{
"key8241": pulumi.String("gcyxztzr"),
},
WorkloadDeploymentName: pulumi.String("l4t"),
})
if err != nil {
return err
}
return nil
})
}
