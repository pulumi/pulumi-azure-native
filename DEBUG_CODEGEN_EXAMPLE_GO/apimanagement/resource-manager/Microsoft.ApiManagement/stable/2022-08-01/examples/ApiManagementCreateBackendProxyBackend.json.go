package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewBackend(ctx, "backend", &apimanagement.BackendArgs{
			BackendId: pulumi.String("proxybackend"),
			Credentials: &apimanagement.BackendCredentialsContractArgs{
				Authorization: &apimanagement.BackendAuthorizationHeaderCredentialsArgs{
					Parameter: pulumi.String("opensesma"),
					Scheme:    pulumi.String("Basic"),
				},
				Header: pulumi.StringArrayMap{
					"x-my-1": pulumi.StringArray{
						pulumi.String("val1"),
						pulumi.String("val2"),
					},
				},
				Query: pulumi.StringArrayMap{
					"sv": pulumi.StringArray{
						pulumi.String("xx"),
						pulumi.String("bb"),
						pulumi.String("cc"),
					},
				},
			},
			Description: pulumi.String("description5308"),
			Protocol:    pulumi.String("http"),
			Proxy: &apimanagement.BackendProxyContractArgs{
				Password: pulumi.String("<password>"),
				Url:      pulumi.String("http://192.168.1.1:8080"),
				Username: pulumi.String("Contoso\\admin"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Tls: &apimanagement.BackendTlsPropertiesArgs{
				ValidateCertificateChain: pulumi.Bool(true),
				ValidateCertificateName:  pulumi.Bool(true),
			},
			Url: pulumi.String("https://backendname2644/"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
