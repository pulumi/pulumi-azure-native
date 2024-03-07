package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewProfile(ctx, "profile", &customerinsights.ProfileArgs{
			ApiEntitySetName: pulumi.String("TestProfileType396"),
			Fields: customerinsights.PropertyDefinitionArray{
				&customerinsights.PropertyDefinitionArgs{
					FieldName:  pulumi.String("Id"),
					FieldType:  pulumi.String("Edm.String"),
					IsArray:    pulumi.Bool(false),
					IsRequired: pulumi.Bool(true),
				},
				&customerinsights.PropertyDefinitionArgs{
					FieldName:  pulumi.String("ProfileId"),
					FieldType:  pulumi.String("Edm.String"),
					IsArray:    pulumi.Bool(false),
					IsRequired: pulumi.Bool(true),
				},
				&customerinsights.PropertyDefinitionArgs{
					FieldName:  pulumi.String("LastName"),
					FieldType:  pulumi.String("Edm.String"),
					IsArray:    pulumi.Bool(false),
					IsRequired: pulumi.Bool(true),
				},
				&customerinsights.PropertyDefinitionArgs{
					FieldName:  pulumi.String("TestProfileType396"),
					FieldType:  pulumi.String("Edm.String"),
					IsArray:    pulumi.Bool(false),
					IsRequired: pulumi.Bool(true),
				},
				&customerinsights.PropertyDefinitionArgs{
					FieldName:  pulumi.String("SavingAccountBalance"),
					FieldType:  pulumi.String("Edm.Int32"),
					IsArray:    pulumi.Bool(false),
					IsRequired: pulumi.Bool(true),
				},
			},
			HubName:            pulumi.String("sdkTestHub"),
			LargeImage:         pulumi.String("\\\\Images\\\\LargeImage"),
			MediumImage:        pulumi.String("\\\\Images\\\\MediumImage"),
			ProfileName:        pulumi.String("TestProfileType396"),
			ResourceGroupName:  pulumi.String("TestHubRG"),
			SchemaItemTypeLink: pulumi.String("SchemaItemTypeLink"),
			SmallImage:         pulumi.String("\\\\Images\\\\smallImage"),
			StrongIds: customerinsights.StrongIdArray{
				&customerinsights.StrongIdArgs{
					KeyPropertyNames: pulumi.StringArray{
						pulumi.String("Id"),
						pulumi.String("SavingAccountBalance"),
					},
					StrongIdName: pulumi.String("Id"),
				},
				&customerinsights.StrongIdArgs{
					KeyPropertyNames: pulumi.StringArray{
						pulumi.String("ProfileId"),
						pulumi.String("LastName"),
					},
					StrongIdName: pulumi.String("ProfileId"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
