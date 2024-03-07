package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costmanagement.NewCloudConnector(ctx, "cloudConnector", &costmanagement.CloudConnectorArgs{
			ConnectorName:     pulumi.String("aws-123456789012"),
			CredentialsKey:    pulumi.String("arn:aws:iam::123456789012:role/AzureCostManagementRole"),
			CredentialsSecret: pulumi.String("external-id"),
			DisplayName:       pulumi.String("AWS-Consolidated-1"),
			ReportId:          pulumi.String("HourlyWithResources"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
