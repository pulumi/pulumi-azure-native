package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewP2sVpnServerConfiguration(ctx, "p2sVpnServerConfiguration", &network.P2sVpnServerConfigurationArgs{
			P2SVpnServerConfigRadiusClientRootCertificates: []network.P2SVpnServerConfigRadiusClientRootCertificateArgs{
				{
					Name:       pulumi.String("p2sVpnServerConfigRadiusClientRootCert1"),
					Thumbprint: pulumi.String("83FFBFC8848B5A5836C94D0112367E16148A286F"),
				},
			},
			P2SVpnServerConfigRadiusServerRootCertificates: []network.P2SVpnServerConfigRadiusServerRootCertificateArgs{
				{
					Name:           pulumi.String("p2sVpnServerConfigRadiusServerRootCert1"),
					PublicCertData: pulumi.String("MIIC5zCCAc+gAwIBAgIQErQ0Hk4aDJxIA+Q5RagB+jANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtQMlNSb290Q2VydDAeFw0xNzEyMTQyMTA3MzhaFw0xODEyMTQyMTI3MzhaMBYxFDASBgNVBAMMC1AyU1Jvb3RDZXJ0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArP7/NQXmW7cQ/ZR1mv3Y3I29Lt7HTOqzo/1KUOoVH3NItbQIRAQbwKy3UWrOFz4eGNX2GWtNRMdCyWsKeqy9Ltsdfcm1IbKXkl84DFeU/ZacXu4Dl3xX3gV5du4TLZjEowJELyur11Ea2YcjPRQ/FzAF9/hGuboS1HZQEPLx4FdUs9OxCYOtc0MxBCwLfVTTRqarb0Ne+arNYd4kCzIhAke1nOyKAJBda5ZL+VHy3S5S8qGlD46jm8HXugmAkUygS4oIIXOmj/1O9sNAi3LN60zufSzCmP8Rm/iUGX+DHAGGiXxwZOKQLEDaZXKqoHjMPP0XudmSWwOIbyeQVrLhkwIDAQABozEwLzAOBgNVHQ8BAf8EBAMCAgQwHQYDVR0OBBYEFEfeNU2trYxNLF9ONmuJUsT13pKDMA0GCSqGSIb3DQEBCwUAA4IBAQBmM6RJzsGGipxyMhimHKN2xlkejhVsgBoTAhOU0llW9aUSwINJ9zFUGgI8IzUFy1VG776fchHp0LMRmPSIUYk5btEPxbsrPtumPuMH8EQGrS+Rt4pD+78c8H1fEPkq5CmDl/PKu4JoFGv+aFcE+Od0hlILstIF10Qysf++QXDolKfzJa/56bgMeYKFiju73loiRM57ns8ddXpfLl792UVpRkFU62LNns6Y1LKTwapmUF4IvIuAIzd6LZNOQng64LAKXtKnViJ1JQiXwf4CEzhgvAti3/ejpb3U90hsrUcyZi6wBv9bZLcAJRWpz61JNYliM1d1grSwQDKGXNQE4xuM"),
				},
			},
			P2SVpnServerConfigVpnClientRevokedCertificates: []network.P2SVpnServerConfigVpnClientRevokedCertificateArgs{
				{
					Name:       pulumi.String("p2sVpnServerConfigVpnClientRevokedCert1"),
					Thumbprint: pulumi.String("83FFBFC8848B5A5836C94D0112367E16148A286F"),
				},
			},
			P2SVpnServerConfigVpnClientRootCertificates: []network.P2SVpnServerConfigVpnClientRootCertificateArgs{
				{
					Name:           pulumi.String("p2sVpnServerConfigVpnClientRootCert1"),
					PublicCertData: pulumi.String("MIIC5zCCAc+gAwIBAgIQErQ0Hk4aDJxIA+Q5RagB+jANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtQMlNSb290Q2VydDAeFw0xNzEyMTQyMTA3MzhaFw0xODEyMTQyMTI3MzhaMBYxFDASBgNVBAMMC1AyU1Jvb3RDZXJ0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArP7/NQXmW7cQ/ZR1mv3Y3I29Lt7HTOqzo/1KUOoVH3NItbQIRAQbwKy3UWrOFz4eGNX2GWtNRMdCyWsKeqy9Ltsdfcm1IbKXkl84DFeU/ZacXu4Dl3xX3gV5du4TLZjEowJELyur11Ea2YcjPRQ/FzAF9/hGuboS1HZQEPLx4FdUs9OxCYOtc0MxBCwLfVTTRqarb0Ne+arNYd4kCzIhAke1nOyKAJBda5ZL+VHy3S5S8qGlD46jm8HXugmAkUygS4oIIXOmj/1O9sNAi3LN60zufSzCmP8Rm/iUGX+DHAGGiXxwZOKQLEDaZXKqoHjMPP0XudmSWwOIbyeQVrLhkwIDAQABozEwLzAOBgNVHQ8BAf8EBAMCAgQwHQYDVR0OBBYEFEfeNU2trYxNLF9ONmuJUsT13pKDMA0GCSqGSIb3DQEBCwUAA4IBAQBmM6RJzsGGipxyMhimHKN2xlkejhVsgBoTAhOU0llW9aUSwINJ9zFUGgI8IzUFy1VG776fchHp0LMRmPSIUYk5btEPxbsrPtumPuMH8EQGrS+Rt4pD+78c8H1fEPkq5CmDl/PKu4JoFGv+aFcE+Od0hlILstIF10Qysf++QXDolKfzJa/56bgMeYKFiju73loiRM57ns8ddXpfLl792UVpRkFU62LNns6Y1LKTwapmUF4IvIuAIzd6LZNOQng64LAKXtKnViJ1JQiXwf4CEzhgvAti3/ejpb3U90hsrUcyZi6wBv9bZLcAJRWpz61JNYliM1d1grSwQDKGXNQE4xuN"),
				},
			},
			P2SVpnServerConfigurationName: pulumi.String("p2sVpnServerConfiguration1"),
			RadiusServerAddress:           pulumi.String("8.9.9.9"),
			RadiusServerSecret:            pulumi.String("<radiusServerSecret>"),
			ResourceGroupName:             pulumi.String("rg1"),
			VirtualWanName:                pulumi.String("virtualWan1"),
			VpnClientIpsecPolicies: []network.IpsecPolicyArgs{
				{
					DhGroup:             pulumi.String("DHGroup14"),
					IkeEncryption:       pulumi.String("AES256"),
					IkeIntegrity:        pulumi.String("SHA384"),
					IpsecEncryption:     pulumi.String("AES256"),
					IpsecIntegrity:      pulumi.String("SHA256"),
					PfsGroup:            pulumi.String("PFS14"),
					SaDataSizeKilobytes: pulumi.Int(429497),
					SaLifeTimeSeconds:   pulumi.Int(86472),
				},
			},
			VpnProtocols: pulumi.StringArray{
				pulumi.String("IkeV2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
