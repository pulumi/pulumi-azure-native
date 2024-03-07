package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/healthcareapis/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := healthcareapis.NewIotConnectorFhirDestination(ctx, "iotConnectorFhirDestination", &healthcareapis.IotConnectorFhirDestinationArgs{
			FhirDestinationName: pulumi.String("dest1"),
			FhirMapping: &healthcareapis.IotMappingPropertiesArgs{
				Content: pulumi.Any{
					Template: []map[string]interface{}{
						map[string]interface{}{
							"template": map[string]interface{}{
								"codes": []map[string]interface{}{
									map[string]interface{}{
										"code":    "8867-4",
										"display": "Heart rate",
										"system":  "http://loinc.org",
									},
								},
								"periodInterval": 60,
								"typeName":       "heartrate",
								"value": map[string]interface{}{
									"defaultPeriod": 5000,
									"unit":          "count/min",
									"valueName":     "hr",
									"valueType":     "SampledData",
								},
							},
							"templateType": "CodeValueFhir",
						},
					},
					TemplateType: "CollectionFhirTemplate",
				},
			},
			FhirServiceResourceId:          pulumi.String("subscriptions/11111111-2222-3333-4444-555566667777/resourceGroups/myrg/providers/Microsoft.HealthcareApis/workspaces/myworkspace/fhirservices/myfhirservice"),
			IotConnectorName:               pulumi.String("blue"),
			Location:                       pulumi.String("westus"),
			ResourceGroupName:              pulumi.String("testRG"),
			ResourceIdentityResolutionType: pulumi.String("Create"),
			WorkspaceName:                  pulumi.String("workspace1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
