package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewVolumeQuotaRule(ctx, "volumeQuotaRule", &netapp.VolumeQuotaRuleArgs{
			AccountName:         pulumi.String("account-9957"),
			Location:            pulumi.String("westus"),
			PoolName:            pulumi.String("pool-5210"),
			QuotaSizeInKiBs:     pulumi.Float64(100005),
			QuotaTarget:         pulumi.String("1821"),
			QuotaType:           pulumi.String("IndividualUserQuota"),
			ResourceGroupName:   pulumi.String("myRG"),
			VolumeName:          pulumi.String("volume-6387"),
			VolumeQuotaRuleName: pulumi.String("rule-0004"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
