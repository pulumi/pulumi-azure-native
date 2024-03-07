package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewAccount(ctx, "account", &netapp.AccountArgs{
			AccountName: pulumi.String("account1"),
			ActiveDirectories: netapp.ActiveDirectoryArray{
				&netapp.ActiveDirectoryArgs{
					AesEncryption:      pulumi.Bool(true),
					Dns:                pulumi.String("10.10.10.3"),
					Domain:             pulumi.String("10.10.10.3"),
					LdapOverTLS:        pulumi.Bool(false),
					LdapSigning:        pulumi.Bool(false),
					OrganizationalUnit: pulumi.String("OU=Engineering"),
					Password:           pulumi.String("ad_password"),
					Site:               pulumi.String("SiteName"),
					SmbServerName:      pulumi.String("SMBServer"),
					Username:           pulumi.String("ad_user_name"),
				},
			},
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("myRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
