package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewHypervClusterControllerCluster(ctx, "hypervClusterControllerCluster", &offazure.HypervClusterControllerClusterArgs{
			ClusterName:       pulumi.String("-18--O4iS57-729MV9PCErt"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("--xY37--V518"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
