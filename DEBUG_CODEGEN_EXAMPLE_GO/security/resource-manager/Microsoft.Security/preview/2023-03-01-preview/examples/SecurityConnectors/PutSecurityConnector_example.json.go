package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewSecurityConnector(ctx, "securityConnector", &security.SecurityConnectorArgs{
			EnvironmentData: security.AwsEnvironmentData{
				EnvironmentType: "AwsAccount",
			},
			EnvironmentName:     pulumi.String("AWS"),
			HierarchyIdentifier: pulumi.String("exampleHierarchyId"),
			Location:            pulumi.String("Central US"),
			Offerings: pulumi.Array{
				security.CspmMonitorAwsOffering{
					NativeCloudConnection: security.CspmMonitorAwsOfferingNativeCloudConnection{
						CloudRoleArn: "arn:aws:iam::00000000:role/ASCMonitor",
					},
					OfferingType: "CspmMonitorAws",
				},
			},
			ResourceGroupName:     pulumi.String("exampleResourceGroup"),
			SecurityConnectorName: pulumi.String("exampleSecurityConnectorName"),
			Tags:                  nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
