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
			ObjectReplicationPolicyId: pulumi.String("2a20bb73-5717-4635-985a-5d4cf777438f"),
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
					RuleId:          pulumi.String("d5d18a48-8801-4554-aeaa-74faf65f5ef9"),
					SourceContainer: pulumi.String("scont139"),
				},
				&storage.ObjectReplicationPolicyRuleArgs{
					DestinationContainer: pulumi.String("dcont179"),
					SourceContainer:      pulumi.String("scont179"),
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
