package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewManagementPolicy(ctx, "managementPolicy", &storage.ManagementPolicyArgs{
			AccountName:          pulumi.String("sto9699"),
			ManagementPolicyName: pulumi.String("default"),
			Policy: &storage.ManagementPolicySchemaArgs{
				Rules: storage.ManagementPolicyRuleArray{
					&storage.ManagementPolicyRuleArgs{
						Definition: &storage.ManagementPolicyDefinitionArgs{
							Actions: &storage.ManagementPolicyActionArgs{
								BaseBlob: &storage.ManagementPolicyBaseBlobArgs{
									Delete: &storage.DateAfterModificationArgs{
										DaysAfterModificationGreaterThan: pulumi.Float64(1000),
									},
									TierToArchive: &storage.DateAfterModificationArgs{
										DaysAfterModificationGreaterThan: pulumi.Float64(90),
									},
									TierToCool: &storage.DateAfterModificationArgs{
										DaysAfterModificationGreaterThan: pulumi.Float64(30),
									},
								},
								Snapshot: &storage.ManagementPolicySnapShotArgs{
									Delete: &storage.DateAfterCreationArgs{
										DaysAfterCreationGreaterThan: pulumi.Float64(30),
									},
								},
							},
							Filters: &storage.ManagementPolicyFilterArgs{
								BlobTypes: pulumi.StringArray{
									pulumi.String("blockBlob"),
								},
								PrefixMatch: pulumi.StringArray{
									pulumi.String("olcmtestcontainer1"),
								},
							},
						},
						Enabled: pulumi.Bool(true),
						Name:    pulumi.String("olcmtest1"),
						Type:    pulumi.String("Lifecycle"),
					},
					&storage.ManagementPolicyRuleArgs{
						Definition: &storage.ManagementPolicyDefinitionArgs{
							Actions: &storage.ManagementPolicyActionArgs{
								BaseBlob: &storage.ManagementPolicyBaseBlobArgs{
									Delete: &storage.DateAfterModificationArgs{
										DaysAfterModificationGreaterThan: pulumi.Float64(1000),
									},
									TierToArchive: &storage.DateAfterModificationArgs{
										DaysAfterModificationGreaterThan: pulumi.Float64(90),
									},
									TierToCool: &storage.DateAfterModificationArgs{
										DaysAfterModificationGreaterThan: pulumi.Float64(30),
									},
								},
							},
							Filters: &storage.ManagementPolicyFilterArgs{
								BlobIndexMatch: storage.TagFilterArray{
									&storage.TagFilterArgs{
										Name:  pulumi.String("tag1"),
										Op:    pulumi.String("=="),
										Value: pulumi.String("val1"),
									},
									&storage.TagFilterArgs{
										Name:  pulumi.String("tag2"),
										Op:    pulumi.String("=="),
										Value: pulumi.String("val2"),
									},
								},
								BlobTypes: pulumi.StringArray{
									pulumi.String("blockBlob"),
								},
								PrefixMatch: pulumi.StringArray{
									pulumi.String("olcmtestcontainer2"),
								},
							},
						},
						Enabled: pulumi.Bool(true),
						Name:    pulumi.String("olcmtest2"),
						Type:    pulumi.String("Lifecycle"),
					},
				},
			},
			ResourceGroupName: pulumi.String("res7687"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
