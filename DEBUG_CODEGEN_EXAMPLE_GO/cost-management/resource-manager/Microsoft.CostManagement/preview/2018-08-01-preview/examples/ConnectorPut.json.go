package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costmanagement.NewConnector(ctx, "connector", &costmanagement.ConnectorArgs{
			ConnectorName:     pulumi.String("AWSBillingAccount"),
			CredentialsKey:    pulumi.String("arn:aws:iam::123456789012:role/AzureCostManagementRole"),
			CredentialsSecret: pulumi.String("external-id"),
			DisplayName:       pulumi.String("AWS-Consolidated-1"),
			Location:          pulumi.String("westus"),
			ReportId:          pulumi.String("HourlyWithResources"),
			ResourceGroupName: pulumi.String("rg1"),
			Status:            pulumi.String("active"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
