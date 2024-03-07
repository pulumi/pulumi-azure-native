package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.AzureDataLakeStoreOutputDataSource{
				AccountName:            "someaccount",
				DateFormat:             "yyyy/MM/dd",
				FilePathPrefix:         "{date}/{time}",
				RefreshToken:           "someRefreshToken==",
				TenantId:               "cea4e98b-c798-49e7-8c40-4a2b3beb47dd",
				TimeFormat:             "HH",
				TokenUserDisplayName:   "Bob Smith",
				TokenUserPrincipalName: "bobsmith@contoso.com",
				Type:                   "Microsoft.DataLake/Accounts",
			},
			JobName:           pulumi.String("sj3310"),
			OutputName:        pulumi.String("output5195"),
			ResourceGroupName: pulumi.String("sjrg6912"),
			Serialization: streamanalytics.JsonSerialization{
				Encoding: "UTF8",
				Format:   "Array",
				Type:     "Json",
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
