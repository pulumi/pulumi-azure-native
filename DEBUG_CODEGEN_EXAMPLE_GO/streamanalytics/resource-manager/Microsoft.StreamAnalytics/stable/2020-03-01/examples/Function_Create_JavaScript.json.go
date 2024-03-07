package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewFunction(ctx, "function", &streamanalytics.FunctionArgs{
			FunctionName: pulumi.String("function8197"),
			JobName:      pulumi.String("sj8653"),
			Properties: streamanalytics.ScalarFunctionProperties{
				Binding: streamanalytics.JavaScriptFunctionBinding{
					Script: "function (x, y) { return x + y; }",
					Type:   "Microsoft.StreamAnalytics/JavascriptUdf",
				},
				Inputs: []streamanalytics.FunctionInputType{
					{
						DataType: "Any",
					},
				},
				Output: streamanalytics.FunctionOutputType{
					DataType: "Any",
				},
				Type: "Scalar",
			},
			ResourceGroupName: pulumi.String("sjrg1637"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
