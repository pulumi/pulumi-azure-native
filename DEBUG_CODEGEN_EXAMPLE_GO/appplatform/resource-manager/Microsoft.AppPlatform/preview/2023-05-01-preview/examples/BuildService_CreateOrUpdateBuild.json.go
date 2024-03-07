package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewBuildServiceBuild(ctx, "buildServiceBuild", &appplatform.BuildServiceBuildArgs{
			BuildName:        pulumi.String("mybuild"),
			BuildServiceName: pulumi.String("default"),
			Properties: &appplatform.BuildPropertiesArgs{
				AgentPool: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/agentPools/default"),
				Apms: appplatform.ApmReferenceArray{
					&appplatform.ApmReferenceArgs{
						ResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apms/myappinsights"),
					},
				},
				Builder: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/builders/default"),
				Certificates: appplatform.CertificateReferenceArray{
					&appplatform.CertificateReferenceArgs{
						ResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert1"),
					},
					&appplatform.CertificateReferenceArgs{
						ResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert2"),
					},
				},
				Env: pulumi.StringMap{
					"environmentVariable": pulumi.String("test"),
				},
				RelativePath: pulumi.String("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855-20210601-3ed9f4a2-986b-4bbd-b833-a42dccb2f777"),
				ResourceRequests: &appplatform.BuildResourceRequestsArgs{
					Cpu:    pulumi.String("1"),
					Memory: pulumi.String("2Gi"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
