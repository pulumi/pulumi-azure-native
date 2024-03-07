package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datashare/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datashare.NewInvitation(ctx, "invitation", &datashare.InvitationArgs{
			AccountName:       pulumi.String("Account1"),
			ExpirationDate:    pulumi.String("2020-08-26T22:33:24.5785265Z"),
			InvitationName:    pulumi.String("Invitation1"),
			ResourceGroupName: pulumi.String("SampleResourceGroup"),
			ShareName:         pulumi.String("Share1"),
			TargetEmail:       pulumi.String("receiver@microsoft.com"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
