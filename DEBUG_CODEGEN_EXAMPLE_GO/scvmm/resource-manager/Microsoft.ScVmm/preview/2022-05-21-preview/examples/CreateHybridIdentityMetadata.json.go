package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/scvmm/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := scvmm.NewHybridIdentityMetadata(ctx, "hybridIdentityMetadata", &scvmm.HybridIdentityMetadataArgs{
			MetadataName:       pulumi.String("default"),
			PublicKey:          pulumi.String("8ec7d60c-9700-40b1-8e6e-e5b2f6f477f2"),
			ResourceGroupName:  pulumi.String("testrg"),
			ResourceUid:        pulumi.String("f8b82dff-38ef-4220-99ef-d3a3f86ddc6c"),
			VirtualMachineName: pulumi.String("ContosoVm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
