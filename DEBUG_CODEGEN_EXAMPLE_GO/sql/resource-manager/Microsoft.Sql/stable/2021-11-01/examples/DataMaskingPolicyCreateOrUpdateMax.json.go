package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDataMaskingPolicy(ctx, "dataMaskingPolicy", &sql.DataMaskingPolicyArgs{
			DataMaskingPolicyName: pulumi.String("Default"),
			DataMaskingState:      sql.DataMaskingStateEnabled,
			DatabaseName:          pulumi.String("sqlcrudtest-331"),
			ExemptPrincipals:      pulumi.String("testuser;"),
			ResourceGroupName:     pulumi.String("sqlcrudtest-6852"),
			ServerName:            pulumi.String("sqlcrudtest-2080"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
