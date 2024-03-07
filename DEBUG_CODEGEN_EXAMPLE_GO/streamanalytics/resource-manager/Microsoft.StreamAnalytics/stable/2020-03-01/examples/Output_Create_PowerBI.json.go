package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.PowerBIOutputDataSource{
				Dataset:                "someDataset",
				GroupId:                "ac40305e-3e8d-43ac-8161-c33799f43e95",
				GroupName:              "MyPowerBIGroup",
				RefreshToken:           "someRefreshToken==",
				Table:                  "someTable",
				TokenUserDisplayName:   "Bob Smith",
				TokenUserPrincipalName: "bobsmith@contoso.com",
				Type:                   "PowerBI",
			},
			JobName:           pulumi.String("sj2331"),
			OutputName:        pulumi.String("output3022"),
			ResourceGroupName: pulumi.String("sjrg7983"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
