package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewBlobServiceProperties(ctx, "blobServiceProperties", &storage.BlobServicePropertiesArgs{
			AccountName:      pulumi.String("sto8607"),
			BlobServicesName: pulumi.String("default"),
			ChangeFeed: &storage.ChangeFeedArgs{
				Enabled:         pulumi.Bool(true),
				RetentionInDays: pulumi.Int(7),
			},
			Cors: &storage.CorsRulesArgs{
				CorsRules: storage.CorsRuleArray{
					&storage.CorsRuleArgs{
						AllowedHeaders: pulumi.StringArray{
							pulumi.String("x-ms-meta-abc"),
							pulumi.String("x-ms-meta-data*"),
							pulumi.String("x-ms-meta-target*"),
						},
						AllowedMethods: pulumi.StringArray{
							pulumi.String("GET"),
							pulumi.String("HEAD"),
							pulumi.String("POST"),
							pulumi.String("OPTIONS"),
							pulumi.String("MERGE"),
							pulumi.String("PUT"),
						},
						AllowedOrigins: pulumi.StringArray{
							pulumi.String("http://www.contoso.com"),
							pulumi.String("http://www.fabrikam.com"),
						},
						ExposedHeaders: pulumi.StringArray{
							pulumi.String("x-ms-meta-*"),
						},
						MaxAgeInSeconds: pulumi.Int(100),
					},
					&storage.CorsRuleArgs{
						AllowedHeaders: pulumi.StringArray{
							pulumi.String("*"),
						},
						AllowedMethods: pulumi.StringArray{
							pulumi.String("GET"),
						},
						AllowedOrigins: pulumi.StringArray{
							pulumi.String("*"),
						},
						ExposedHeaders: pulumi.StringArray{
							pulumi.String("*"),
						},
						MaxAgeInSeconds: pulumi.Int(2),
					},
					&storage.CorsRuleArgs{
						AllowedHeaders: pulumi.StringArray{
							pulumi.String("x-ms-meta-12345675754564*"),
						},
						AllowedMethods: pulumi.StringArray{
							pulumi.String("GET"),
							pulumi.String("PUT"),
						},
						AllowedOrigins: pulumi.StringArray{
							pulumi.String("http://www.abc23.com"),
							pulumi.String("https://www.fabrikam.com/*"),
						},
						ExposedHeaders: pulumi.StringArray{
							pulumi.String("x-ms-meta-abc"),
							pulumi.String("x-ms-meta-data*"),
							pulumi.String("x -ms-meta-target*"),
						},
						MaxAgeInSeconds: pulumi.Int(2000),
					},
				},
			},
			DefaultServiceVersion: pulumi.String("2017-07-29"),
			DeleteRetentionPolicy: &storage.DeleteRetentionPolicyArgs{
				Days:    pulumi.Int(300),
				Enabled: pulumi.Bool(true),
			},
			IsVersioningEnabled: pulumi.Bool(true),
			ResourceGroupName:   pulumi.String("res4410"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
