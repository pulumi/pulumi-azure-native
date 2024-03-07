package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewFunction(ctx, "function", &streamanalytics.FunctionArgs{
			FunctionName: pulumi.String("function588"),
			JobName:      pulumi.String("sj9093"),
			Properties: streamanalytics.ScalarFunctionProperties{
				Binding: streamanalytics.AzureMachineLearningWebServiceFunctionBinding{
					ApiKey:    "someApiKey==",
					BatchSize: 1000,
					Endpoint:  "someAzureMLEndpointURL",
					Inputs: streamanalytics.AzureMachineLearningWebServiceInputs{
						ColumnNames: []streamanalytics.AzureMachineLearningWebServiceInputColumn{
							{
								DataType: "string",
								MapTo:    0,
								Name:     "tweet",
							},
						},
						Name: "input1",
					},
					Outputs: []streamanalytics.AzureMachineLearningWebServiceOutputColumn{
						{
							DataType: "string",
							Name:     "Sentiment",
						},
					},
					Type: "Microsoft.MachineLearning/WebService",
				},
				Inputs: []streamanalytics.FunctionInputType{
					{
						DataType: "nvarchar(max)",
					},
				},
				Output: streamanalytics.FunctionOutputType{
					DataType: "nvarchar(max)",
				},
				Type: "Scalar",
			},
			ResourceGroupName: pulumi.String("sjrg7"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
