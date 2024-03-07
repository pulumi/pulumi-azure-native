package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurearcdata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurearcdata.NewFailoverGroup(ctx, "failoverGroup", &azurearcdata.FailoverGroupArgs{
			FailoverGroupName: pulumi.String("testFailoverGroupName"),
			Properties: azurearcdata.FailoverGroupPropertiesResponse{
				PartnerManagedInstanceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/sqlManagedInstances/partnerMI"),
				Spec: &azurearcdata.FailoverGroupSpecArgs{
					PartnerSyncMode: pulumi.String("async"),
					Role:            pulumi.String("primary"),
				},
			},
			ResourceGroupName:      pulumi.String("testrg"),
			SqlManagedInstanceName: pulumi.String("testSqlManagedInstance"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
