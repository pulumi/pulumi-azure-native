package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datafactory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datafactory.NewDataFlow(ctx, "dataFlow", &datafactory.DataFlowArgs{
			DataFlowName: pulumi.String("exampleDataFlow"),
			FactoryName:  pulumi.String("exampleFactoryName"),
			Properties: datafactory.MappingDataFlow{
				Description: "Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation.",
				ScriptLines: []string{
					"source(output(",
					"PreviousConversionRate as double,",
					"Country as string,",
					"DateTime1 as string,",
					"CurrentConversionRate as double",
					"),",
					"allowSchemaDrift: false,",
					"validateSchema: false) ~> USDCurrency",
					"source(output(",
					"PreviousConversionRate as double,",
					"Country as string,",
					"DateTime1 as string,",
					"CurrentConversionRate as double",
					"),",
					"allowSchemaDrift: true,",
					"validateSchema: false) ~> CADSource",
					"USDCurrency, CADSource union(byName: true)~> Union",
					"Union derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn",
					"NewCurrencyColumn split(Country == 'USD',",
					"Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)",
					"ConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink",
					"ConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink",
				},
				Sinks: []datafactory.DataFlowSink{
					{
						Dataset: {
							ReferenceName: "USDOutput",
							Type:          "DatasetReference",
						},
						Name: "USDSink",
					},
					{
						Dataset: {
							ReferenceName: "CADOutput",
							Type:          "DatasetReference",
						},
						Name: "CADSink",
					},
				},
				Sources: []datafactory.DataFlowSource{
					{
						Dataset: {
							ReferenceName: "CurrencyDatasetUSD",
							Type:          "DatasetReference",
						},
						Name: "USDCurrency",
					},
					{
						Dataset: {
							ReferenceName: "CurrencyDatasetCAD",
							Type:          "DatasetReference",
						},
						Name: "CADSource",
					},
				},
				Type: "MappingDataFlow",
			},
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
