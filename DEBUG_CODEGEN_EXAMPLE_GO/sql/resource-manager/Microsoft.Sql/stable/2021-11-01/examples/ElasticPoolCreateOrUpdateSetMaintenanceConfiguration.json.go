package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewElasticPool(ctx, "elasticPool", &sql.ElasticPoolArgs{
			ElasticPoolName:            pulumi.String("sqlcrudtest-8102"),
			Location:                   pulumi.String("Japan East"),
			MaintenanceConfigurationId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_1"),
			ResourceGroupName:          pulumi.String("sqlcrudtest-2369"),
			ServerName:                 pulumi.String("sqlcrudtest-8069"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
