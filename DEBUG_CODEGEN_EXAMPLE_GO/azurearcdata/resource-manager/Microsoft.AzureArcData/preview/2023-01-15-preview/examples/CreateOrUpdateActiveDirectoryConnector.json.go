package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurearcdata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurearcdata.NewActiveDirectoryConnector(ctx, "activeDirectoryConnector", &azurearcdata.ActiveDirectoryConnectorArgs{
			ActiveDirectoryConnectorName: pulumi.String("testADConnector"),
			DataControllerName:           pulumi.String("testdataController"),
			Properties: &azurearcdata.ActiveDirectoryConnectorPropertiesArgs{
				Spec: &azurearcdata.ActiveDirectoryConnectorSpecArgs{
					ActiveDirectory: &azurearcdata.ActiveDirectoryConnectorDomainDetailsArgs{
						DomainControllers: &azurearcdata.ActiveDirectoryDomainControllersArgs{
							PrimaryDomainController: &azurearcdata.ActiveDirectoryDomainControllerArgs{
								Hostname: pulumi.String("dc1.contoso.local"),
							},
							SecondaryDomainControllers: azurearcdata.ActiveDirectoryDomainControllerArray{
								&azurearcdata.ActiveDirectoryDomainControllerArgs{
									Hostname: pulumi.String("dc2.contoso.local"),
								},
								&azurearcdata.ActiveDirectoryDomainControllerArgs{
									Hostname: pulumi.String("dc3.contoso.local"),
								},
							},
						},
						Realm:                      pulumi.String("CONTOSO.LOCAL"),
						ServiceAccountProvisioning: pulumi.String("manual"),
					},
					Dns: &azurearcdata.ActiveDirectoryConnectorDNSDetailsArgs{
						NameserverIPAddresses: pulumi.StringArray{
							pulumi.String("11.11.111.111"),
							pulumi.String("22.22.222.222"),
						},
						PreferK8sDnsForPtrLookups: pulumi.Bool(false),
						Replicas:                  pulumi.Float64(1),
					},
				},
			},
			ResourceGroupName: pulumi.String("testrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
