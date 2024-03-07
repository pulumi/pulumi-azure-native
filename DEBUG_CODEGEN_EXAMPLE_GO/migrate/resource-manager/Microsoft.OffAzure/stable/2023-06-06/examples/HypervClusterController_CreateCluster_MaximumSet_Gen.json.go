package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewHypervClusterControllerCluster(ctx, "hypervClusterControllerCluster", &offazure.HypervClusterControllerClusterArgs{
			ClusterName: pulumi.String("v5285"),
			Fqdn:        pulumi.String("lyqquicmqfagukcwfquemrkrexic"),
			HostFqdnList: pulumi.StringArray{
				pulumi.String("frozqzatdahnwlccznpmw"),
			},
			ProvisioningState: pulumi.String("Created"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			RunAsAccountId:    pulumi.String("valugnnrofauhagzzxksfjbcwuqhue"),
			SiteName:          pulumi.String("224qC-GNR"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
