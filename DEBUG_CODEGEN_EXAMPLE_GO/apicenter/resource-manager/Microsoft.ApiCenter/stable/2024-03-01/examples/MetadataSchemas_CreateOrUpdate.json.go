package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apicenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apicenter.NewMetadataSchema(ctx, "metadataSchema", &apicenter.MetadataSchemaArgs{
			AssignedTo: apicenter.MetadataAssignmentArray{
				&apicenter.MetadataAssignmentArgs{
					Deprecated: pulumi.Bool(true),
					Entity:     pulumi.String("api"),
				},
			},
			MetadataSchemaName: pulumi.String("author"),
			ResourceGroupName:  pulumi.String("contoso-resources"),
			Schema:             pulumi.String("{\"type\":\"string\", \"title\":\"Author\", pattern: \"^[a-zA-Z]+$\"}"),
			ServiceName:        pulumi.String("contoso"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
