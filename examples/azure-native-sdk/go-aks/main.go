// Copyright 2021, Pulumi Corporation.  All rights reserved.

package main

import (
	"encoding/base64"

	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"

	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		randomRgName, err := random.NewRandomString(ctx, "randomRgName", &random.RandomStringArgs{
			Length:  pulumi.Int(12),
			Special: pulumi.Bool(false),
			Upper:   pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
		// Create an Azure Resource Group
		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", &resources.ResourceGroupArgs{
			ResourceGroupName: randomRgName.Result,
			Location:          pulumi.String("WestUS"),
		})
		if err != nil {
			return err
		}

		// Create an AD service principal.
		adApp, err := azuread.NewApplication(ctx, "aks", nil)
		if err != nil {
			return err
		}

		adSp, err := azuread.NewServicePrincipal(ctx, "aksSp", &azuread.ServicePrincipalArgs{
			ApplicationId: adApp.ApplicationId,
		})
		if err != nil {
			return err
		}

		// Create the auto-generated Service Principal password.
		adSpPassword, err := azuread.NewServicePrincipalPassword(ctx, "aksSpPassword", &azuread.ServicePrincipalPasswordArgs{
			ServicePrincipalId: adSp.ID(),
			EndDate:            pulumi.String("2099-01-01T00:00:00Z"),
		})
		if err != nil {
			return err
		}

		// Generate an SSH key.
		sshArgs := tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA"),
			RsaBits:   pulumi.Int(4096),
		}
		sshKey, err := tls.NewPrivateKey(ctx, "ssh-key", &sshArgs)
		if err != nil {
			return err
		}

		randomClusterName, err := random.NewRandomString(ctx, "randomClusterName", &random.RandomStringArgs{
			Length:  pulumi.Int(12),
			Special: pulumi.Bool(false),
			Upper:   pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
		cluster, err := containerservice.NewManagedCluster(ctx, "cluster", &containerservice.ManagedClusterArgs{
			ResourceName:      randomClusterName.Result,
			Location:          pulumi.String("WestUS"),
			ResourceGroupName: resourceGroup.Name,
			DnsPrefix:         randomClusterName.Result,
			AgentPoolProfiles: containerservice.ManagedClusterAgentPoolProfileArray{
				&containerservice.ManagedClusterAgentPoolProfileArgs{
					Name:         pulumi.String("agentpool"),
					Mode:         pulumi.String("System"),
					OsDiskSizeGB: pulumi.Int(30),
					Count:        pulumi.Int(3),
					VmSize:       pulumi.String("Standard_DS2_v2"),
					OsType:       pulumi.String("Linux"),
				},
			},
			LinuxProfile: &containerservice.ContainerServiceLinuxProfileArgs{
				AdminUsername: pulumi.String("testuser"),
				Ssh: containerservice.ContainerServiceSshConfigurationArgs{
					PublicKeys: containerservice.ContainerServiceSshPublicKeyArray{
						containerservice.ContainerServiceSshPublicKeyArgs{
							KeyData: sshKey.PublicKeyOpenssh,
						},
					},
				},
			},
			ServicePrincipalProfile: &containerservice.ManagedClusterServicePrincipalProfileArgs{
				ClientId: adApp.ApplicationId,
				Secret:   adSpPassword.Value,
			},
			KubernetesVersion: pulumi.String("1.22.4"),
		})
		if err != nil {
			return err
		}

		ctx.Export("kubeconfig", pulumi.All(cluster.Name, resourceGroup.Name, resourceGroup.ID()).ApplyT(func(args interface{}) (string, error) {
			clusterName := args.([]interface{})[0].(string)
			resourceGroupNAme := args.([]interface{})[1].(string)
			creds, err := containerservice.ListManagedClusterUserCredentials(ctx, &containerservice.ListManagedClusterUserCredentialsArgs{
				ResourceGroupName: resourceGroupNAme,
				ResourceName:      clusterName,
			})
			if err != nil {
				return "", err
			}
			encoded := creds.Kubeconfigs[0].Value
			kubeconfig, err := base64.StdEncoding.DecodeString(encoded)
			if err != nil {
				return "", err
			}
			return string(kubeconfig), nil
		}).(pulumi.StringOutput))

		return nil
	})
}
