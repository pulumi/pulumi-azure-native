package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewGeoBackupPolicy(ctx, "geoBackupPolicy", &sql.GeoBackupPolicyArgs{
			DatabaseName:        pulumi.String("testdw"),
			GeoBackupPolicyName: pulumi.String("Default"),
			ResourceGroupName:   pulumi.String("sqlcrudtest-4799"),
			ServerName:          pulumi.String("sqlcrudtest-5961"),
			State:               sql.GeoBackupPolicyStateEnabled,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
