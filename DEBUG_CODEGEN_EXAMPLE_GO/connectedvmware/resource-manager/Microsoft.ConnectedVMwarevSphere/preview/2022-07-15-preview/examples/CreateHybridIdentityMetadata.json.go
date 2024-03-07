package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/connectedvmwarevsphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connectedvmwarevsphere.NewHybridIdentityMetadatum(ctx, "hybridIdentityMetadatum", &connectedvmwarevsphere.HybridIdentityMetadatumArgs{
			MetadataName:       pulumi.String("default"),
			PublicKey:          pulumi.String("8ec7d60c-9700-40b1-8e6e-e5b2f6f477f2"),
			ResourceGroupName:  pulumi.String("testrg"),
			VirtualMachineName: pulumi.String("ContosoVm"),
			VmId:               pulumi.String("f8b82dff-38ef-4220-99ef-d3a3f86ddc6c"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
