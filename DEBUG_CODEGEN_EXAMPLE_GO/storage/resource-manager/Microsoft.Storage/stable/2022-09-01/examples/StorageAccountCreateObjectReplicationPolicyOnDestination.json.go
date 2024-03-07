package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewObjectReplicationPolicy(ctx, "objectReplicationPolicy", &storage.ObjectReplicationPolicyArgs{
			AccountName:               pulumi.String("dst112"),
			DestinationAccount:        pulumi.String("dst112"),
			ObjectReplicationPolicyId: pulumi.String("default"),
			ResourceGroupName:         pulumi.String("res7687"),
			Rules: storage.ObjectReplicationPolicyRuleArray{
				&storage.ObjectReplicationPolicyRuleArgs{
					DestinationContainer: pulumi.String("dcont139"),
					Filters: &storage.ObjectReplicationPolicyFilterArgs{
						PrefixMatch: pulumi.StringArray{
							pulumi.String("blobA"),
							pulumi.String("blobB"),
						},
					},
					SourceContainer: pulumi.String("scont139"),
				},
			},
			SourceAccount: pulumi.String("src1122"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
