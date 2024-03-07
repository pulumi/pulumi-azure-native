package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/chaos/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := chaos.NewTarget(ctx, "target", &chaos.TargetArgs{
			ParentProviderNamespace: pulumi.String("Microsoft.Compute"),
			ParentResourceName:      pulumi.String("exampleVM"),
			ParentResourceType:      pulumi.String("virtualMachines"),
			Properties: pulumi.Any{
				Identities: []map[string]interface{}{
					map[string]interface{}{
						"subject": "CN=example.subject",
						"type":    "CertificateSubjectIssuer",
					},
				},
			},
			ResourceGroupName: pulumi.String("exampleRG"),
			TargetName:        pulumi.String("Microsoft-Agent"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
