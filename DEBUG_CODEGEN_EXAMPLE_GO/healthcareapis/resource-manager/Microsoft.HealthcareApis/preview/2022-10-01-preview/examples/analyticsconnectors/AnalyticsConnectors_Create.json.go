package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/healthcareapis/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := healthcareapis.NewAnalyticsConnector(ctx, "analyticsConnector", &healthcareapis.AnalyticsConnectorArgs{
			AnalyticsConnectorName: pulumi.String("exampleconnector"),
			DataDestinationConfiguration: &healthcareapis.AnalyticsConnectorDataLakeDataDestinationArgs{
				DataLakeName: pulumi.String("exampledatalake"),
				Type:         pulumi.String("datalake"),
			},
			DataMappingConfiguration: &healthcareapis.AnalyticsConnectorFhirToParquetMappingArgs{
				ExtensionSchemaReference:     pulumi.String("acrexample.azurecr.io/blah@sha256aaa/Extension"),
				FilterConfigurationReference: pulumi.String("acrexample.azurecr.io/blah@sha256xxx"),
				Type:                         pulumi.String("fhirToParquet"),
			},
			DataSourceConfiguration: &healthcareapis.AnalyticsConnectorFhirServiceDataSourceArgs{
				Kind: pulumi.String("R4"),
				Type: pulumi.String("fhirservice"),
				Url:  pulumi.String("https://workspace-examplefhir.fhir.azurehealthcareapis.com"),
			},
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("testRG"),
			WorkspaceName:     pulumi.String("workspace1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
