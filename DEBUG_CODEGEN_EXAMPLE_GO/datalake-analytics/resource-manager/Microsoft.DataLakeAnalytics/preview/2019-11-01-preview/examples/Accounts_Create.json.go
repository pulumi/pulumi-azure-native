package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datalakeanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datalakeanalytics.NewAccount(ctx, "account", &datalakeanalytics.AccountArgs{
			AccountName: pulumi.String("contosoadla"),
			ComputePolicies: datalakeanalytics.CreateComputePolicyWithAccountParametersArray{
				&datalakeanalytics.CreateComputePolicyWithAccountParametersArgs{
					MaxDegreeOfParallelismPerJob: pulumi.Int(1),
					MinPriorityPerJob:            pulumi.Int(1),
					Name:                         pulumi.String("test_policy"),
					ObjectId:                     pulumi.String("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
					ObjectType:                   pulumi.String("User"),
				},
			},
			DataLakeStoreAccounts: datalakeanalytics.AddDataLakeStoreWithAccountParametersArray{
				&datalakeanalytics.AddDataLakeStoreWithAccountParametersArgs{
					Name:   pulumi.String("test_adls"),
					Suffix: pulumi.String("test_suffix"),
				},
			},
			DefaultDataLakeStoreAccount: pulumi.String("test_adls"),
			FirewallAllowAzureIps:       datalakeanalytics.FirewallAllowAzureIpsStateEnabled,
			FirewallRules: datalakeanalytics.CreateFirewallRuleWithAccountParametersArray{
				&datalakeanalytics.CreateFirewallRuleWithAccountParametersArgs{
					EndIpAddress:   pulumi.String("2.2.2.2"),
					Name:           pulumi.String("test_rule"),
					StartIpAddress: pulumi.String("1.1.1.1"),
				},
			},
			FirewallState:                datalakeanalytics.FirewallStateEnabled,
			Location:                     pulumi.String("eastus2"),
			MaxDegreeOfParallelism:       pulumi.Int(30),
			MaxDegreeOfParallelismPerJob: pulumi.Int(1),
			MaxJobCount:                  pulumi.Int(3),
			MinPriorityPerJob:            pulumi.Int(1),
			NewTier:                      datalakeanalytics.TierTypeConsumption,
			QueryStoreRetention:          pulumi.Int(30),
			ResourceGroupName:            pulumi.String("contosorg"),
			StorageAccounts: datalakeanalytics.AddStorageAccountWithAccountParametersArray{
				&datalakeanalytics.AddStorageAccountWithAccountParametersArgs{
					AccessKey: pulumi.String("34adfa4f-cedf-4dc0-ba29-b6d1a69ab346"),
					Name:      pulumi.String("test_storage"),
					Suffix:    pulumi.String("test_suffix"),
				},
			},
			Tags: pulumi.StringMap{
				"test_key": pulumi.String("test_value"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
